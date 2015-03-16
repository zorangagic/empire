package empire

import (
	"testing"
)

func TestIsValid(t *testing.T) {
	tests := []struct {
		app App
		err error
	}{
		{App{}, ErrInvalidName},
		{App{Name: "api"}, nil},
		{App{Name: "r101-api"}, nil},
	}

	for _, tt := range tests {
		if err := tt.app.IsValid(); err != tt.err {
			t.Fatal("%v.IsValid() => %v; want %v", tt.app, err, tt.err)
		}
	}
}

type mockappsService struct {
	appsService

	AppsFindOrCreateByRepoFunc func(repoType string, repo Repo) (*App, error)
}

func (s *mockappsService) AppsEnsureRepo(app *App, repoType string, repo Repo) error {
	return nil
}

func (s *mockappsService) AppsFindOrCreateByRepo(repoType string, repo Repo) (*App, error) {
	if s.AppsFindOrCreateByRepoFunc != nil {
		return s.AppsFindOrCreateByRepoFunc(repoType, repo)
	}

	return nil, nil
}
