# 🔗 Go URL Shortener

A simple URL shortener written in **Go (Golang)** using native libraries like `net/http`, `crypto/md5`, and `encoding/json`.

---

## 📌 Features

- 🔐 Generates a short URL using MD5 hash of the original URL
- 📦 Stores shortened URLs in-memory using a `map`
- 🧭 Redirects short URLs to the original long URLs
- 🧰 API-based architecture (JSON requests & responses)

---

## 🚀 API Endpoints

### 1. **Shorten a URL**
**POST** `/shorten`

**Request Body (JSON):**
```json
{
  "url": "https://example.com/long-url"
}

![image](https://github.com/user-attachments/assets/d685b181-489a-444f-9cc8-7b37f75907fb)

