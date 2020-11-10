package access_token

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAccessTokenConstants(t *testing.T) {
	assert.EqualValues(t, 24, expirationTime, "expiration time should be 24 hours")
}

func TestGetNewAccessToken(t *testing.T) {
	// is not expired
	// is a newacess token
	at := GetNewAccessToken()
	assert.False(t, at.isExpired(), "brand new access token should NOT be expired")
	assert.Equal(t, "", at.AccessToken, "new access token should not have defined access token id")
	assert.True(t, at.UserId == 0, "new access token should not have an assocated user id")
}

func TestAccessTokenIsExpired(t *testing.T) {
	at := AccessToken{}
	// is expired
	assert.True(t, at.isExpired(), "empty access token should be expired by default")
	at.Expires = time.Now().UTC().Add(3 * time.Hour).Unix()
	assert.False(t, at.isExpired(), "access token created three hours from now should NOT be expired")
}
