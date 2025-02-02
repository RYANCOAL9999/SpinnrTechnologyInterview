package models

import (
	"time"
)

type Status int

const (
	Ready Status = iota
	Joined
)

// table for Prize Pool
type PrizePool struct {
	ID     int     `json:"id"`
	Amount float64 `json:"amount" binding:"required"`
}

// table for Challenge
type Challenge struct {
	ID          int       `json:"id"`
	PlayerID    string    `json:"player_id" binding:"required"`
	Amount      float64   `json:"amount" binding:"required"`
	Status      Status    `json:"status" binding:"required"`
	Won         bool      `json:"won" binding:"required"`
	CreatedAt   time.Time `json:"created_at" binding:"required"`
	Probability float64   `json:"probability" binding:"required"`
}

// New Challenge Struct for request
type NewChallengeNeed struct {
	PlayerID int     `json:"player_id" binding:"required"`
	Amount   float64 `json:"amount" binding:"required" validate:"eq=20.01"`
}

// JoinChallengeResponse represents player joined after status.
type JoinChallengeResponse struct {
	Status Status `json:"status"`
}

// ErrorResponse represents an error response with a single error message.
type ErrorResponse struct {
	Error string `json:"error"`
}
