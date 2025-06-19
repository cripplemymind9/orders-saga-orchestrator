package entity

import "time"

type SagaStatus string

const (
	SagaStatusUnspecified SagaStatus = "unspecified"
	SagaStatusPending     SagaStatus = "pending"
	SagaStatusCompleted   SagaStatus = "completed"
	SagaStatusFailed      SagaStatus = "failed"
	SagaStatusCompensated SagaStatus = "compensated"
)

type Saga struct {
	ID        int64
	OrderID   int64
	Status    SagaStatus
	Steps     []Step
	CreatedAt time.Time
	UpdatedAt *time.Time
}

type Step struct {
	Name      string
	Status    string
}
