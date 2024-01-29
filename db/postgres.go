package db

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"sync"
)

type PostgreSQL struct {
	DB *gorm.DB
}

var instance *PostgreSQL
var once sync.Once

// NewPostgreSQL constructor for PostgreSQL struct.
func NewPostgreSQL(connStr string) (*PostgreSQL, error) {
	once.Do(func() {
		db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
		if err != nil {
			log.Println("could not connect to database with err : %v", err)
			return
		}
		instance = &PostgreSQL{DB: db}
	})

	return instance, nil
}

func (p *PostgreSQL) Ping(ctx *gin.Context) error {
	db, err := p.DB.DB()

	if err = db.Ping(); err != nil {
		return err
	}

	return nil
}
