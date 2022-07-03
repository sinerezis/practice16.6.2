package main_test

import (
	. "skillFactory/multithreading/practice"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBalance(t *testing.T) {
	u := NewClient()
	u.AccountBalance = 100
	assert.Equal(t, u.Balance(), 100)
}

func TestWindrawal(t *testing.T) {
	u := NewClient()
	u.AccountBalance = 100
	assert.NoError(t, u.Withdrawal(100))
	assert.Error(t, u.Withdrawal(500))
}

func TestDeposit(t *testing.T) {
	u := NewClient()
	u.Deposit(100)
	assert.Equal(t, u.Balance(), 100)
}
