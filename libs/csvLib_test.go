package libs

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type testSuite struct {
	suite.Suite
}

// Suite is a basic testing suite with methods for storing and
// retrieving the current *testing.T context.
type Suite struct {
	*assert.Assertions
	require *require.Assertions
	t       *testing.T
}

// T retrieves the current *testing.T context.
func (suite *Suite) T() *testing.T {
	return suite.t
}

func (t *testSuite) TestCsv() {
	assert := assert.New(t.T())

	f := "./../problems.csv"
	problems := ReadCsvFile(&f)

	assert.Equal(13, len(problems))
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(testSuite))
}
