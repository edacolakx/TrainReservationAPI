// main.go
package main

import (
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Vagon struct {
	Ad       string `json:"Ad"`
	Kapasite int    `json:"Kapasite"`
	Yolcu    int    `json:"Yolcu"`
}

type Tren struct {
	Ad       string  `json:"Ad"`
	Vagonlar []Vagon `json:"Vagonlar"`
}

type Response struct {
	WagonName string `json:"wagonName"`
	Person    int    `json:"person"`
}

func loadTrenler() ([]Tren, error) {
	file, err := os.Open("data.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var trenler []Tren
	err = json.Unmarshal(bytes, &trenler)
	if err != nil {
		return nil, err
	}

	return trenler, nil

}

type Data struct {
	Name            string `json:"Name"`
	PassengerCount  int    `json:"PassengerCount"`
	DifferentWagons bool   `json:"DifferentWagons"`
}

func main() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://train-reservation-swart.vercel.app"}, // Vite React uygulaması portu
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	r.GET("/", func(c *gin.Context) {
		trenler, err := loadTrenler()
		if err != nil {
			c.JSON(500, gin.H{
				"error": "Veri yüklenirken hata oluştu",
			})
			return
		}
		c.JSON(200, gin.H{
			"message": http.StatusOK,
			"trenler": trenler,
		})
	})

	r.POST("/check-reservation", func(c *gin.Context) {
		var data Data
		if err := c.BindJSON(&data); err != nil {
			c.JSON(400, gin.H{
				"error": "Geçersiz istek verisi",
			})
			return
		}
		var availableWagons []Vagon
		trenler, err := loadTrenler()
		if err != nil {
			c.JSON(500, gin.H{
				"error": "Veri yüklenirken hata oluştu",
			})
			return
		}

		for _, tren := range trenler {
			if tren.Ad == data.Name {
				for _, vagon := range tren.Vagonlar {
					if float64(vagon.Kapasite)*0.7 >= float64(vagon.Yolcu) {
						availableWagons = append(availableWagons, vagon)
					} else {
					}
				}

			}
		}
		var finalResponse []Response
		var selectedTrainName string
		var selectedPersonCount int

		if len(availableWagons) > 0 {
			if data.DifferentWagons == false {
				found := false
				for _, vagon := range availableWagons {
					emptyseats := float64(vagon.Kapasite)*0.7 - float64(vagon.Yolcu)
					if emptyseats >= float64(data.PassengerCount) {
						selectedTrainName = vagon.Ad
						selectedPersonCount = data.PassengerCount
						finalResponse = append(finalResponse, Response{
							WagonName: selectedTrainName,
							Person:    selectedPersonCount,
						})
						found = true
						break
					}
				}

				if !found {
					c.JSON(200, gin.H{
						"reservationAvailable": false,
						"availableWagons":      []string{},
					})
					return
				}

				c.JSON(200, gin.H{
					"reservationAvailable": true,
					"availableWagons":      finalResponse,
				})

			} else if data.DifferentWagons == true {
				p := data.PassengerCount / len(availableWagons)
				remain := data.PassengerCount % len(availableWagons)
				for i, vagon := range availableWagons {
					personForThisWagon := p
					if i == len(availableWagons)-1 && remain > 0 {
						personForThisWagon += remain
					}
					finalResponse = append(finalResponse, Response{
						WagonName: vagon.Ad,
						Person:    personForThisWagon,
					})
				}

				c.JSON(200, gin.H{
					"reservationAvailable": true,
					"availableWagons":      finalResponse,
				})
			}
		} else {
			c.JSON(200, gin.H{
				"reservationAvailable": false,
				"availableWagons":      []string{},
			})

		}
	})
	r.Run()
}
