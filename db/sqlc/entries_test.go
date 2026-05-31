package db

import (
	"context"
	"testing"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/sayanmondal31/simple_bank/utils"
	"github.com/stretchr/testify/require"
)

func createRandomEntry(t *testing.T) Entry {
	argAccount := CreateAccountParams{
		Owner:    utils.RandomOwner(),
		Balance:  utils.RandomMoney(),
		Currency: utils.RandomCurrency(),
	}
	// create account
	account1, err := testQueries.CreateAccount(context.Background(), argAccount)

	require.NoError(t, err)

	argEntry := CreateEntryParams{
		AccountID: account1.ID,
		Amount:    utils.RandomMoney(),
	}

	entry1, err := testQueries.CreateEntry(context.Background(), argEntry)

	// there should be no error while building
	require.NoError(t, err)
	// data which is fed into db should be created same as arguments
	require.Equal(t, argEntry.AccountID, entry1.AccountID)
	require.Equal(t, argEntry.Amount, entry1.Amount)

	// id and date should be empy
	require.NotEmpty(t, entry1.ID)
	require.NotEmpty(t, entry1.CreatedAt)

	return entry1
}

func TestCreateEntry(t *testing.T) {
	createRandomEntry(t)
}

func TestGetEntry(t *testing.T) {
	entry1 := createRandomEntry(t)
	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)

	require.NoError(t, err)

	require.Equal(t, entry1, entry2)

}

func TestListEntryParams(t *testing.T) {

	for range 10 {
		createRandomEntry(t)
	}

	arg := ListEntryParams{
		Limit:  5,
		Offset: 5,
	}

	entriesListData, err := testQueries.ListEntry(context.Background(), arg)

	require.NoError(t, err)

	for _, entry := range entriesListData {
		require.NotEmpty(t, entry)
	}
}

func TestUpdateEntry(t *testing.T) {
	entry1 := createRandomEntry(t)

	arg := UpdateEntryParams{
		Amount: utils.RandomMoney(),
		ID:     entry1.ID,
	}

	updatedEntry, err := testQueries.UpdateEntry(context.Background(), arg)

	require.NoError(t, err)
	require.Equal(t, entry1.ID, updatedEntry.ID)
	require.Equal(t, arg.Amount, updatedEntry.Amount)
	require.WithinDuration(t, entry1.CreatedAt.Time, updatedEntry.CreatedAt.Time, time.Second)
}

func TestDeleteEntry(t *testing.T) {
	entry1 := createRandomEntry(t)

	err := testQueries.DeleteEntry(context.Background(), entry1.ID)

	require.NoError(t, err)

	findEntry, err := testQueries.GetEntry(context.Background(), entry1.ID)

	require.Error(t, err)
	require.EqualError(t, err, pgx.ErrNoRows.Error())
	require.Zero(t, findEntry)
}
