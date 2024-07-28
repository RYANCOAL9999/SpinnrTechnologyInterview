package databases

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	object "github.com/RYANCOAL9999/SpinnrTechnologyInterview/gameRoomManagementSystem/databases"
	"github.com/stretchr/testify/assert"
)

func TestAizuArray(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []int
	}{
		{
			name:     "Empty string",
			input:    "",
			expected: []int{},
		},
		{
			name:     "Single number",
			input:    "42",
			expected: []int{42},
		},
		{
			name:     "Multiple numbers",
			input:    "1, 2, 3, 4, 5",
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Numbers with spaces",
			input:    " 10, 20 , 30 ",
			expected: []int{10, 20, 30},
		},
		{
			name:     "Zero and negative numbers",
			input:    "0, -1, 2, -3",
			expected: []int{0, -1, 2, -3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := object.AizuArray(tt.input)
			assert.Equal(t, tt.expected, result, "aizuArray(%q) = %v; want %v", tt.input, result, tt.expected)
		})
	}
}

func TestAizuArrayWithInvalidInput(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{
			name:  "Non-numeric string",
			input: "a, b, c",
		},
		{
			name:  "Mixed valid and invalid",
			input: "1, 2, three, 4",
		},
		{
			name:  "Invalid separator",
			input: "1; 2; 3",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := object.AizuArray(tt.input)
			// The function ignores non-numeric values, so we expect an array of zeros
			for _, v := range result {
				assert.Zero(t, v, "aizuArray(%q) should return zero for invalid inputs", tt.input)
			}
		})
	}
}

func TestListReservation(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer db.Close()

	mockRows := sqlmock.NewRows([]string{"ReservationID", "RoomID", "ReservationDate", "PlayerIDs"}).
		AddRow(1, 101, time.Now(), "1, 2, 3").
		AddRow(2, 102, time.Now().Add(24*time.Hour), "4, 5")

	mock.ExpectQuery("SELECT (.+) FROM Reservation R INNER JOIN Room RM ON R.RoomID = RM.ID WHERE").
		WillReturnRows(mockRows)

	mock.ExpectQuery("SELECT (.+) FROM Player P INNER JOIN Level L ON P.LevelID = L.ID WHERE P.ID IN").
		WillReturnRows(sqlmock.NewRows([]string{"ID", "Name", "LV"}).
			AddRow(1, "Player1", 5).
			AddRow(2, "Player2", 3).
			AddRow(3, "Player3", 7).
			AddRow(4, "Player4", 2).
			AddRow(5, "Player5", 6))

	reservations, err := object.ListReservation(db, 0, time.Time{}, time.Time{}, 0)

	assert.NoError(t, err)
	assert.Len(t, reservations, 2)
	assert.Equal(t, 101, reservations[0].RoomID)
	assert.Equal(t, 102, reservations[1].RoomID)
	assert.Len(t, reservations[0].Player, 3)
	assert.Len(t, reservations[1].Player, 2)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}

func TestInsertReservation(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer db.Close()

	roomID := 101
	reservationDate := time.Now()

	mock.ExpectExec("INSERT INTO Reservation").
		WithArgs(roomID, reservationDate).
		WillReturnResult(sqlmock.NewResult(1, 1))

	id, err := object.InsertReservation(db, roomID, reservationDate)

	assert.NoError(t, err)
	assert.Equal(t, 1, id)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}
