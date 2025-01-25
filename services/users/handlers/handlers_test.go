package handlers

import (
	"fmt"
	"log"
	"testing"

	"github.com/parrothacker1/Solvelt/users/database"
	"github.com/parrothacker1/Solvelt/users/models"
	"github.com/parrothacker1/Solvelt/users/testutils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func TestMain(m *testing.M) {
  log.Println("Starting handler tests")
  container, err := testutils.SetupTestDB()
  log.Println("Created container for postgres")
	if err != nil {
		log.Fatalf("Failed to create Test Postgres Container: %v", err)
	}
	containerHost, err := container.Host(testutils.TestContainerContext)
	if err != nil {
		log.Fatalf("Failed to get host from container: %v", err)
	}

	containerPort, err := container.MappedPort(testutils.TestContainerContext, "5432/tcp")
	if err != nil {
		log.Fatalf("Failed to get port from container: %v", err)
	}

	endpoint := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
    "tester_solvelt", "testing_solvelt", containerHost, containerPort.Port(), "testdb_solvelt")

	database.DB, err = gorm.Open(postgres.Open(endpoint), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the postgres container: %v", err)
	}
  log.Println("Successfully connected to the postgres database.")
	if err := database.DB.AutoMigrate(&models.User{}, &models.Team{}); err != nil {
		log.Fatalf("Failed to migrate models: %v", err)
	}
  log.Println("Successfully migrated models.")
  m.Run()
  Cleanup()
}

func Cleanup() {
  fmt.Println("testing cleanup")
}
