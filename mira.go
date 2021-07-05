package mira

import "net/http"

// Init is used
// when we initialize the Reddit instance,
// automatically start a goroutine that will
// update the token every 45 minutes. The
// auto_refresh should not be accessible to
// the end user as it is an internal method
func Init(c Credentials) (*Reddit, error) {
	auth, err := Authenticate(&c)
	if err != nil {
		return nil, err
	}
	auth.Client = &http.Client{}
	auth.SetDefault()
	go auth.autoRefresh()
	return auth, nil
}
