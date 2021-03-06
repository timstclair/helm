package main

import (
	"net/http/httptest"
	"regexp"

	"github.com/kubernetes/helm/cmd/manager/router"
	"github.com/kubernetes/helm/pkg/chart"
	"github.com/kubernetes/helm/pkg/common"
	"github.com/kubernetes/helm/pkg/httputil"
	"github.com/kubernetes/helm/pkg/repo"
)

// httpHarness is a simple test server fixture.
// Simple fixture for standing up a test server with a single route.
//
// You must Close() the returned server.
func httpHarness(c *router.Context, route string, fn router.HandlerFunc) *httptest.Server {
	h := router.NewHandler(c)
	h.Add(route, fn)
	return httptest.NewServer(h)
}

// stubContext creates a stub of a Context object.
//
// This creates a stub context with the following properties:
// - Config is initialized to empty values
// - Encoder is initialized to httputil.DefaultEncoder
// - CredentialProvider is initialized to repo.InmemCredentialProvider
// - Manager is initialized to mockManager.
func stubContext() *router.Context {
	return &router.Context{
		Config:             &router.Config{},
		Manager:            &mockManager{},
		CredentialProvider: repo.NewInmemCredentialProvider(),
		Encoder:            httputil.DefaultEncoder,
	}
}

type mockManager struct{}

func (m *mockManager) ListDeployments() ([]common.Deployment, error) {
	return []common.Deployment{}, nil
}

func (m *mockManager) GetDeployment(name string) (*common.Deployment, error) {
	return &common.Deployment{}, nil
}

func (m *mockManager) CreateDeployment(t *common.Template) (*common.Deployment, error) {
	return &common.Deployment{}, nil
}

func (m *mockManager) DeleteDeployment(name string, forget bool) (*common.Deployment, error) {
	return &common.Deployment{}, nil
}

func (m *mockManager) PutDeployment(name string, t *common.Template) (*common.Deployment, error) {
	return &common.Deployment{}, nil
}

func (m *mockManager) ListManifests(deploymentName string) (map[string]*common.Manifest, error) {
	return map[string]*common.Manifest{}, nil
}

func (m *mockManager) GetManifest(deploymentName string, manifest string) (*common.Manifest, error) {
	return &common.Manifest{}, nil
}

func (m *mockManager) Expand(t *common.Template) (*common.Manifest, error) {
	return &common.Manifest{}, nil
}

func (m *mockManager) ListCharts() ([]string, error) {
	return []string{}, nil
}

func (m *mockManager) ListChartInstances(chartName string) ([]*common.ChartInstance, error) {
	return []*common.ChartInstance{}, nil
}

func (m *mockManager) GetRepoForChart(chartName string) (string, error) {
	return "", nil
}

func (m *mockManager) GetMetadataForChart(chartName string) (*chart.Chartfile, error) {
	return nil, nil
}

func (m *mockManager) GetChart(chartName string) (*chart.Chart, error) {
	return nil, nil
}

func (m *mockManager) AddChartRepo(addition repo.IRepo) error {
	return nil
}

func (m *mockManager) ListChartRepos() ([]string, error) {
	return []string{}, nil
}

func (m *mockManager) RemoveChartRepo(name string) error {
	return nil
}

func (m *mockManager) GetChartRepo(URL string) (repo.IRepo, error) {
	return nil, nil
}

func (m *mockManager) ListRepos() ([]*repo.Repo, error) {
	return []*repo.Repo{}, nil
}

func (m *mockManager) CreateRepo(pr *repo.Repo) error {
	return nil
}
func (m *mockManager) GetRepo(name string) (*repo.Repo, error) {
	return &repo.Repo{}, nil
}
func (m *mockManager) DeleteRepo(name string) error {
	return nil
}

func (m *mockManager) ListRepoCharts(repoName string, regex *regexp.Regexp) ([]string, error) {
	return []string{}, nil
}

func (m *mockManager) GetChartForRepo(repoName, chartName string) (*chart.Chart, error) {
	return nil, nil
}

func (m *mockManager) CreateCredential(name string, c *repo.Credential) error {
	return nil
}
func (m *mockManager) GetCredential(name string) (*repo.Credential, error) {
	return &repo.Credential{}, nil
}
