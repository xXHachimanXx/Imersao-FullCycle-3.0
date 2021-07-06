package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/xXHachimanXx/Imersao-FullCycle-3.0/infrastructure/kafka"
	repository "github.com/xXHachimanXx/Imersao-FullCycle-3.0/infrastructure/reposiroty"
	"github.com/xXHachimanXx/Imersao-FullCycle-3.0/usecase"
)

// func main() {
// 	db := setupDb()
// 	defer db.Close()

// 	cc := domain.NewCreditCard()
// 	cc.Number = "1234"
// 	cc.Name = "Wesley"
// 	cc.ExpirationYear = 2021
// 	cc.ExpirationMonth = 7
// 	cc.CVV = 123
// 	cc.Limit = 1000
// 	cc.Balance = 0

// 	repo := repository.NewTransactionRepositoryDb(db)
// 	err := repo.CreateCreditCard(*cc)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// }

func main() {
	db := setupDb()
	defer db.Close()

}

func setupKafkaProducer() kafka.KafkaProducer {
	producer := kafka.NewKafkaProducer()
	producer.SetupProducer(os.Getenv("KafkaBootstrapServers"))

	return producer
}

func setupTransactionUseCase(db *sql.DB) usecase.UseCaseTransaction {
	transactionRepository := repository.NewTransactionRepositoryDb(db)
	useCase := usecase.NewUseCaseTransaction(transactionRepository)

	return useCase
}

func setupDb() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		"db",
		"5432",
		"postgres",
		"root",
		"codebank",
	)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatal("error connection to database")
	}

	return db
}
