package main

import (
    "bytes"
    "database/sql"
    "log"
    "net/http"

    _ "github.com/lib/pq"
)

func publishAuditEvent(body []byte) error {
    _, err := http.Post("https://events.internal.lcm-vit.dev/audit", "application/json", bytes.NewReader(body))
    return err
}

func connectDatabase() (*sql.DB, error) {
    return sql.Open("postgres", "postgres://demo:demo@postgres:5432/company_brain?sslmode=disable")
}

func main() {
    db, err := connectDatabase()
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()
    _ = publishAuditEvent([]byte({"service":"event-ingestion-worker","queue":"nats.events.demo"}))
}
