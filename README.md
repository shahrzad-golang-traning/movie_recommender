# 🎬 Movie Recommender API Client - Enhanced Version

A powerful Go program that fetches movie recommendations from a public IMDb-style API with extended features:  
command-line search input, multiple API calls, response time logging, and result saving to a JSON file.



## 🚀 Features

- 🎯 Accepts user input keyword via CLI to search movies dynamically  
- 🔄 Performs multiple HTTP GET requests to improve reliability  
- ⏱ Measures and logs response time and HTTP metadata for each call  
- 🎥 Parses and displays movie details: title, release year, actors, IMDb URL, and poster  
- 💾 Saves all API call results in a pretty-printed `results.json` file  
- 🛡 Handles errors during network calls, response reading, and JSON parsing gracefully  

---

## 🧱 Structs Overview

### ApiResponse

```go
type ApiResponse struct {
  Result     bool    `json:"ok"`
  StatusCode int     `json:"error_code"`
  Movies     []Movie `json:"description"`
}

Movie
type Movie struct {
  Name        string `json:"#TITLE"`
  ReleaseYear int    `json:"#YEAR"`
  Actors      string `json:"#ACTORS"`
  ImdbURL     string `json:"#IMDB_URL"`
  ImagePoster string `json:"#IMG_POSTER"`
}

CallResult
type CallResult struct {
  URL        string        `json:"url"`
  CallDate   string        `json:"call_date"`
  Duration   string        `json:"duration"`
  StatusCode int           `json:"status_code"`
  Headers    http.Header   `json:"headers"`
  Movies     []Movie       `json:"movies"`
}
```

🔧 How to Run
Prerequisites
Go 1.18 or later installed

Run the project
```bash
go run cmd/main.go
```

📤 Sample Usage and Output
```bash
🎬 Enter movie keyword to search: Spiderman

✅ ===[ CALL 1 ]===
URL: https://imdb.iamidiotareyoutoo.com/search?q=Spiderman
Date: 2025-07-25 14:05:21
Duration: 241ms
Status Code: 200
Movies:
🎞 Spider-Man (2002)
🎞 Spider-Man 2 (2004)
🎞 Spider-Man: No Way Home (2021)
```

📁 Results saved to results.json

🧪 Future Improvements

🌐 Add CLI flags for controlling number of requests and output formats

🔍 Implement filtering and sorting options (by year, actor, etc.)

🎨 Enhance output with colors and better formatting

💾 Support saving results in CSV or database

🤝 Contributions
Contributions and pull requests are very welcome!
If you find bugs or have suggestions, please open an issue.