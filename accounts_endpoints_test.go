package nordigen_test

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_GetAccount(t *testing.T) {
	t.Parallel()

	t.Run("get an account by ID", func(t *testing.T) {
		t.Parallel()

		client := getTestClient(t)
		assert.NotNil(t, client)

		token, err := client.NewToken(context.Background())
		assert.NoError(t, err)
		assert.NotNil(t, token)

		testAccountID := os.Getenv("NORDIGEN_TEST_ACCOUNT_ID")

		account, err := client.GetAccount(context.Background(), token.Access, testAccountID)
		assert.NoError(t, err)
		assert.NotNil(t, account)
	})

	t.Run("get an account by invalid ID", func(t *testing.T) {
		t.Parallel()

		client := getTestClient(t)
		assert.NotNil(t, client)

		token, err := client.NewToken(context.Background())
		assert.NoError(t, err)
		assert.NotNil(t, token)

		account, err := client.GetAccount(context.Background(), token.Access, "invalid")
		assert.Error(t, err)
		assert.Nil(t, account)
	})
}

func TestClient_GetAccountBalances(t *testing.T) {
	t.Parallel()

	t.Run("get balances for an account by ID", func(t *testing.T) {
		t.Parallel()

		client := getTestClient(t)
		assert.NotNil(t, client)

		token, err := client.NewToken(context.Background())
		assert.NoError(t, err)
		assert.NotNil(t, token)

		testAccountID := os.Getenv("NORDIGEN_TEST_ACCOUNT_ID")

		balances, err := client.GetAccountBalances(context.Background(), token.Access, testAccountID)
		assert.NoError(t, err)
		assert.NotNil(t, balances)
	})

	t.Run("get balances for an account by invalid ID", func(t *testing.T) {
		t.Parallel()

		client := getTestClient(t)
		assert.NotNil(t, client)

		token, err := client.NewToken(context.Background())
		assert.NoError(t, err)
		assert.NotNil(t, token)

		balances, err := client.GetAccountBalances(context.Background(), token.Access, "invalid")
		assert.Error(t, err)
		assert.Nil(t, balances)
	})
}

func TestClient_GetAccountDetails(t *testing.T) {
	t.Parallel()

	t.Run("get details for an account by ID", func(t *testing.T) {
		t.Parallel()

		client := getTestClient(t)
		assert.NotNil(t, client)

		token, err := client.NewToken(context.Background())
		assert.NoError(t, err)
		assert.NotNil(t, token)

		testAccountID := os.Getenv("NORDIGEN_TEST_ACCOUNT_ID")

		details, err := client.GetAccountDetails(context.Background(), token.Access, testAccountID)
		assert.NoError(t, err)
		assert.NotNil(t, details)
	})

	t.Run("get details for an account by invalid ID", func(t *testing.T) {
		t.Parallel()

		client := getTestClient(t)
		assert.NotNil(t, client)

		token, err := client.NewToken(context.Background())
		assert.NoError(t, err)
		assert.NotNil(t, token)

		details, err := client.GetAccountDetails(context.Background(), token.Access, "invalid")
		assert.Error(t, err)
		assert.Nil(t, details)
	})
}

func TestClient_GetAccountTransactions(t *testing.T) {
	t.Parallel()

	t.Run("get transactions for an account by ID", func(t *testing.T) {
		t.Parallel()

		client := getTestClient(t)
		assert.NotNil(t, client)

		token, err := client.NewToken(context.Background())
		assert.NoError(t, err)
		assert.NotNil(t, token)

		testAccountID := os.Getenv("NORDIGEN_TEST_ACCOUNT_ID")

		transactions, err := client.GetAccountTransactions(context.Background(), token.Access, testAccountID)
		assert.NoError(t, err)
		assert.NotNil(t, transactions)
	})

	t.Run("get transactions for an account by invalid ID", func(t *testing.T) {
		t.Parallel()

		client := getTestClient(t)
		assert.NotNil(t, client)

		token, err := client.NewToken(context.Background())
		assert.NoError(t, err)
		assert.NotNil(t, token)

		transactions, err := client.GetAccountTransactions(context.Background(), token.Access, "invalid")
		assert.Error(t, err)
		assert.Nil(t, transactions)
	})
}
