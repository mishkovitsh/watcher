package server

import (
	"log"
	"watcher/internal/database"
	"net/http"

	"github.com/gin-gonic/gin"
)



func SetupRouter() *gin.Engine {
	// Create a Gin router with default middleware (logger and recovery)
	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")
	r.Static("/assets/", "templates/assets/")
	
	r.GET("/", func(c *gin.Context) {
        // 1. Ovdje izvlačimo tko je korisnik (Caddy nam to šalje)
        userID := c.GetHeader("X-Pocketid-Uid")
        
        // Za sad samo ispiši u log da vidiš radi li auth
        log.Printf("[+] User %s pristupa dashboardu", userID)

        // 2. Dohvati stranice iz baze
        sites, err := database.GetAllWebsites()
        if err != nil {
            log.Println("[-] Greška pri čitanju baze:", err)
            c.HTML(http.StatusInternalServerError, "index.html", gin.H{
                "Error": "Baza podataka nije dostupna.",
            })
            return
        }

        // 3. Pošalji sve u HTML
        c.HTML(http.StatusOK, "index.html", gin.H{
            "Websites": sites,
            "Total":    len(sites),
            "User":     userID, // Pošalji i ID ako ga želiš ispisati negdje
        })
    })

	r.POST("/api/websites", func(c *gin.Context) {
		ownerID := c.GetHeader("X-Pocketid-Uid")
		name := c.PostForm("name")
		url := c.PostForm("url")
		dsc := c.PostForm("description")
		isPublic := c.PostForm("is_public") == "on"
		
	
		err := database.AddWebsite(ownerID, name, url, dsc, isPublic)

		if err != nil {
			log.Println("[-] Database rejected call:", err)
			c.String(http.StatusInternalServerError, "Error while writing to database.")
			return
		}
	})

return  r
}
