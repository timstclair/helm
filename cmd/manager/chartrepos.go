package main

import (
	"github.com/kubernetes/helm/cmd/manager/router"
	"github.com/kubernetes/helm/pkg/httputil"
	"github.com/kubernetes/helm/pkg/repo"
	"github.com/kubernetes/helm/pkg/util"

	"net/http"
	"net/url"
	"regexp"
)

func registerChartRepoRoutes(c *router.Context, h *router.Handler) {
	h.Add("GET /repositories", listChartReposHandlerFunc)
	h.Add("GET /repositories/*", getChartRepoHandlerFunc)
	h.Add("GET /repositories/*/charts", listRepoChartsHandlerFunc)
	h.Add("GET /repositories/*/charts/*", getRepoChartHandlerFunc)
	h.Add("POST /repositories", addChartRepoHandlerFunc)
	h.Add("DELETE /repositories", removeChartRepoHandlerFunc)
}

func listChartReposHandlerFunc(w http.ResponseWriter, r *http.Request, c *router.Context) error {
	handler := "manager: list chart repositories"
	repos, err := c.Manager.ListChartRepos()
	if err != nil {
		return err
	}

	util.LogHandlerExitWithJSON(handler, w, repos, http.StatusOK)
	return nil
}

func addChartRepoHandlerFunc(w http.ResponseWriter, r *http.Request, c *router.Context) error {
	handler := "manager: add chart repository"
	util.LogHandlerEntry(handler, r)
	defer r.Body.Close()
	cr := &repo.Repo{}
	if err := httputil.Decode(w, r, cr); err != nil {
		httputil.BadRequest(w, r, err)
		return nil
	}

	if err := c.Manager.AddChartRepo(cr); err != nil {
		httputil.BadRequest(w, r, err)
		return nil
	}

	util.LogHandlerExitWithText(handler, w, "added", http.StatusOK)
	return nil
}

func removeChartRepoHandlerFunc(w http.ResponseWriter, r *http.Request, c *router.Context) error {
	handler := "manager: remove chart repository"
	util.LogHandlerEntry(handler, r)
	URL, err := pos(w, r, 2)
	if err != nil {
		return err
	}

	err = c.Manager.RemoveChartRepo(URL)
	if err != nil {
		return err
	}

	util.LogHandlerExitWithText(handler, w, "removed", http.StatusOK)
	return nil
}

func getChartRepoHandlerFunc(w http.ResponseWriter, r *http.Request, c *router.Context) error {
	handler := "manager: get repository"
	util.LogHandlerEntry(handler, r)
	repoURL, err := pos(w, r, 2)
	if err != nil {
		return err
	}

	cr, err := c.Manager.GetChartRepo(repoURL)
	if err != nil {
		httputil.BadRequest(w, r, err)
		return nil
	}

	util.LogHandlerExitWithJSON(handler, w, cr, http.StatusOK)
	return nil
}

func listRepoChartsHandlerFunc(w http.ResponseWriter, r *http.Request, c *router.Context) error {
	handler := "manager: list repository charts"
	util.LogHandlerEntry(handler, r)
	repoURL, err := pos(w, r, 2)
	if err != nil {
		return err
	}

	values, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		httputil.BadRequest(w, r, err)
		return nil
	}

	var regex *regexp.Regexp
	regexString := values.Get("regex")
	if regexString != "" {
		regex, err = regexp.Compile(regexString)
		if err != nil {
			httputil.BadRequest(w, r, err)
			return nil
		}
	}

	repoCharts, err := c.Manager.ListRepoCharts(repoURL, regex)
	if err != nil {
		return err
	}

	util.LogHandlerExitWithJSON(handler, w, repoCharts, http.StatusOK)
	return nil
}

func getRepoChartHandlerFunc(w http.ResponseWriter, r *http.Request, c *router.Context) error {
	handler := "manager: get repository charts"
	util.LogHandlerEntry(handler, r)
	repoURL, err := pos(w, r, 2)
	if err != nil {
		return err
	}

	chartName, err := pos(w, r, 4)
	if err != nil {
		return err
	}

	repoChart, err := c.Manager.GetChartForRepo(repoURL, chartName)
	if err != nil {
		return err
	}

	util.LogHandlerExitWithJSON(handler, w, repoChart, http.StatusOK)
	return nil
}
