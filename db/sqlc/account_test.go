package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateAccount(testingContext *testing.T) {
	accountDetail := CreateAccountParams{
		Owner: "Fadhil", 
		Balance: 19000000,
		Currency: "IDR",
	}

	createdAccount, err := testQueries.CreateAccount(context.Background(), accountDetail)
	require.NoError(testingContext, err)
	require.NotEmpty(testingContext, createdAccount)

	require.Equal(testingContext, accountDetail.Owner, createdAccount.Owner)
	require.Equal(testingContext, accountDetail.Balance, createdAccount.Balance)
	require.Equal(testingContext, accountDetail.Currency, createdAccount.Currency)

	require.NotZero(testingContext, createdAccount.ID)
	require.NotZero(testingContext, createdAccount.CreatedAt)  
}