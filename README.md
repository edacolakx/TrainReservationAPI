# 🚂 Tren Rezervasyon Sistemi

Bu proje, tren rezervasyon işlemlerini yönetmek için geliştirilmiş modern bir web uygulamasıdır. Backend Go (Gin framework) ile, frontend ise React (Vite) ile geliştirilmiştir.

## 📋 Proje Özellikleri

### Backend (Go)

- **Framework**: Gin Web Framework
- **CORS Desteği**: Frontend ile güvenli iletişim
- **JSON API**: RESTful API endpoints
- **Veri Yönetimi**: JSON dosyasından tren ve vagon bilgilerini okuma

### Frontend (React + Vite)

- **Framework**: React 18
- **Build Tool**: Vite
- **Port**: 5173 (Development)
- **Modern UI**: Responsive ve kullanıcı dostu arayüz

## 🚀 Özellikler

- **Tren Listesi**: Mevcut trenleri ve vagonlarını görüntüleme
- **Rezervasyon Kontrolü**: Yolcu sayısına göre uygun vagon bulma
- **Esnek Rezervasyon**: Aynı vagon veya farklı vagonlarda yerleştirme seçenekleri
- **Kapasite Yönetimi**: %70 doluluk oranı kontrolü
- **Gerçek Zamanlı Kontrol**: Anlık müsaitlik kontrolü

## 🛠️ Teknolojiler

### Backend

- **Go 1.23.0**
- **Gin Framework v1.10.1**
- **Gin CORS Middleware**

### Frontend

- **React 18**
- **Vite**
- **Modern JavaScript (ES6+)**

## 📦 Kurulum

### Gereksinimler

- Go 1.23.0 veya üzeri
- npm veya yarn

### Backend Kurulumu

1. Projeyi klonlayın:

```bash
git clone <repository-url>
cd tren
```

2. Go modüllerini yükleyin:

```bash
go mod tidy
```

3. Backend'i çalıştırın:

```bash
go run main.go
```

Backend varsayılan olarak `http://localhost:8080` adresinde çalışacaktır.

## 📡 API Endpoints

### GET `/`

Tüm trenleri ve vagonlarını getirir.

**Response:**

```json
{
  "message": 200,
  "trenler": [
    {
      "Ad": "Başkent Ekspres",
      "Vagonlar": [
        {
          "Ad": "Vagon 1",
          "Kapasite": 100,
          "Yolcu": 70
        }
      ]
    }
  ]
}
```

### POST `/check-reservation`

Rezervasyon kontrolü yapar.

**Request Body:**

```json
{
  "Name": "Başkent Ekspres",
  "PassengerCount": 5,
  "DifferentWagons": false
}
```

**Response:**

```json
{
  "reservationAvailable": true,
  "availableWagons": [
    {
      "wagonName": "Vagon 2",
      "person": 5
    }
  ]
}
```

## 📊 Veri Yapısı

### Tren Yapısı

```json
{
  "Ad": "Tren Adı",
  "Vagonlar": [
    {
      "Ad": "Vagon Adı",
      "Kapasite": 100,
      "Yolcu": 70
    }
  ]
}
```

## 🔧 Konfigürasyon

### CORS Ayarları

Backend, frontend ile iletişim için CORS ayarları yapılandırılmıştır:

- **Origin**: `http://localhost:5173`
- **Methods**: GET, POST, PUT, DELETE, OPTIONS
- **Headers**: Origin, Content-Type, Accept, Authorization

### Kapasite Kontrolü

- Vagonlar %70 doluluk oranına kadar rezervasyona açıktır
- Bu oran `main.go` dosyasında `0.7` değeri ile ayarlanmıştır

## 🎯 Kullanım Senaryoları

### Senaryo 1: Aynı Vagonda Rezervasyon

- **DifferentWagons**: `false`
- Sistem tek bir vagonda yeterli yer arar
- Tüm yolcular aynı vagonda yerleştirilir

### Senaryo 2: Farklı Vagonlarda Rezervasyon

- **DifferentWagons**: `true`
- Yolcular mevcut vagonlar arasında dağıtılır
- Eşit dağılım sağlanır

## 📝 Lisans

Bu proje eğitim amaçlı geliştirilmiştir.

## 👨‍💻 Geliştirici

Bu proje iş görüşmesi için hazırlanmıştır ve modern web teknolojilerini kullanarak geliştirilmiştir.

---

**Not**: Bu proje Go backend ve React frontend teknolojilerini kullanarak geliştirilmiş tam kapsamlı bir tren rezervasyon sistemidir. Frontend projesini https://github.com/edacolakx/TrainReservation adresinde bulabilirsiniz.
