package testtool

import (
	"net"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestGetFreePort(t *testing.T) {
	p, err := GetFreePort()
	require.NoError(t, err)
	require.NotEmpty(t, p)
}

func TestWaitForPort(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
	defer s.Close()

	u, err := url.Parse(s.URL)
	require.Nil(t, err)
	_, port, err := net.SplitHostPort(u.Host)
	require.Nil(t, err)

	require.NoError(t, WaitForPort(t, time.Second, port))
}

func TestWaitForPort_timeout(t *testing.T) {
	p, err := GetFreePort()
	require.NoError(t, err)

	require.Error(t, WaitForPort(t, time.Second, p))
}
