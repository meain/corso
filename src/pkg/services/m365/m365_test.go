package m365

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
)

type M365IntegrationSuite struct {
	suite.Suite
}

func TestM365IntegrationSuite(t *testing.T) {
	if err := tester.RunOnAny(
		tester.CorsoCITests,
	); err != nil {
		t.Skip(err)
	}

	suite.Run(t, new(M365IntegrationSuite))
}

func (suite *M365IntegrationSuite) SetupSuite() {
	_, err := tester.GetRequiredEnvSls(
		tester.M365AcctCredEnvs)
	require.NoError(suite.T(), err)
}

func (suite *M365IntegrationSuite) TestUsers() {
	ctx, flush := tester.NewContext()
	defer flush()

	acct := tester.NewM365Account(suite.T())

	users, err := Users(ctx, acct)
	require.NoError(suite.T(), err)

	require.NotNil(suite.T(), users)
	require.Greater(suite.T(), len(users), 0)
}

func (suite *M365IntegrationSuite) TestUserIDs() {
	ctx, flush := tester.NewContext()
	defer flush()

	acct := tester.NewM365Account(suite.T())

	ids, err := UserIDs(ctx, acct)
	require.NoError(suite.T(), err)

	require.NotNil(suite.T(), ids)
	require.Greater(suite.T(), len(ids), 0)
}

func (suite *M365IntegrationSuite) TestGetEmailAndUserID() {
	ctx, flush := tester.NewContext()
	defer flush()

	acct := tester.NewM365Account(suite.T())

	ids, err := GetEmailAndUserID(ctx, acct)
	require.NoError(suite.T(), err)

	require.NotNil(suite.T(), ids)
	require.Greater(suite.T(), len(ids), 0)
}
