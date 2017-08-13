package travis

import (
	"fmt"
	"net/http"
)

// EnvironmentVariablesService handles communication with the builds
// related methods of the Travis CI API.
type EnvironmentVariablesService struct {
	client *Client
}

// EnvironmentVariable represents a Travis CI settings environment variable
type EnvironmentVariable struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	Value        string `json:"value"`
	Public       bool   `json:"public"`
	RepositoryId uint   `json:"repository_id"`
}

// EnvironmentVariablesListOptions specifies the optional parameters to the
// EnvironmentVariable.List method.
type EnvironmentVariablesListOptions struct {
	// repository ids to fetch environment variables for
	Id uint `url:"repository_id,omitempty"`
}

// listEnvironmentVariablesResponse represents the response of a call
// to the Travis CI list builds endpoint.
type listEnvironmentVariablesResponse struct {
	EnvironmentVariables []EnvironmentVariable `json:"env_vars"`
}

// List lists environment variables using the provided options.
//
// Travis CI API docs: http://docs.travis-ci.com/api/settings/env_vars
func (rs *EnvironmentVariablesService) List(opt *EnvironmentVariablesListOptions) ([]EnvironmentVariable, *http.Response, error) {
	u, err := urlWithOptions("/settings/env_vars", opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := rs.client.NewRequest("GET", u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var envVarsResp listEnvironmentVariablesResponse
	resp, err := rs.client.Do(req, &envVarsResp)
	if err != nil {
		return nil, resp, err
	}

	return envVarsResp.EnvironmentVariables, resp, err
}

// getEnvironmentVariableResponse represents the response of a call
// to the Travis CI get environment variables endpoint.
type getEnvironmentVariableResponse struct {
	EnvironmentVariable EnvironmentVariable `json:"env_var"`
}

// EnvironmentVariablesRequestOptions specifies the optional parameters to the
// EnvironmentVariable.Get method.
type EnvironmentVariablesRequestOptions struct {
	// repository ids to fetch environment variables for
	RepositoryId uint `url:"repository_id,omitempty"`
}

// Get fetches an environment variable by id provided.
//
// Travis CI API docs: https://docs.travis-ci.com/api/?http#settings:-environment-variables
func (rs *EnvironmentVariablesService) Get(id string, repositoryId uint) (*EnvironmentVariable, *http.Response, error) {
	opts := EnvironmentVariablesRequestOptions{RepositoryId: repositoryId}
	u, err := urlWithOptions(fmt.Sprintf("/settings/env_vars/%s", id), &opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := rs.client.NewRequest("GET", u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var envVarResp getEnvironmentVariableResponse
	resp, err := rs.client.Do(req, &envVarResp)
	if err != nil {
		return nil, resp, err
	}

	return &envVarResp.EnvironmentVariable, resp, err
}

// environmentVariablePostBody represents the response of a call
// to the Travis CI get environment variables endpoint.
type environmentVariablePostBody struct {
	EnvironmentVariable EnvironmentVariable `json:"env_var"`
}

// Create fetches an environment variable by id provided.
//
// Travis CI API docs: https://docs.travis-ci.com/api/?http#settings:-environment-variables
func (rs *EnvironmentVariablesService) Create(repositoryId uint, envVar *EnvironmentVariable) (*EnvironmentVariable, *http.Response, error) {
	opts := EnvironmentVariablesRequestOptions{RepositoryId: repositoryId}
	u, err := urlWithOptions("/settings/env_vars", &opts)
	if err != nil {
		return nil, nil, err
	}

	body := environmentVariablePostBody{EnvironmentVariable: *envVar}

	req, err := rs.client.NewRequest("POST", u, body, nil)
	if err != nil {
		return nil, nil, err
	}

	var envVarResp getEnvironmentVariableResponse
	resp, err := rs.client.Do(req, &envVarResp)
	if err != nil {
		return nil, resp, err
	}

	return &envVarResp.EnvironmentVariable, resp, err
}

// Update updates an environment variable.
//
// Travis CI API docs: https://docs.travis-ci.com/api/?http#settings:-environment-variables
func (rs *EnvironmentVariablesService) Update(repositoryId uint, envVar *EnvironmentVariable) (*EnvironmentVariable, *http.Response, error) {
	opts := EnvironmentVariablesRequestOptions{RepositoryId: repositoryId}
	id := envVar.Id
	u, err := urlWithOptions(fmt.Sprintf("/settings/env_vars/%s", id), &opts)
	if err != nil {
		return nil, nil, err
	}

	body := environmentVariablePostBody{EnvironmentVariable: *envVar}

	req, err := rs.client.NewRequest("PATCH", u, body, nil)
	if err != nil {
		return nil, nil, err
	}

	var envVarResp getEnvironmentVariableResponse
	resp, err := rs.client.Do(req, &envVarResp)
	if err != nil {
		return nil, resp, err
	}

	return &envVarResp.EnvironmentVariable, resp, err
}

// Update updates an environment variable.
//
// Travis CI API docs: https://docs.travis-ci.com/api/?http#settings:-environment-variables
func (rs *EnvironmentVariablesService) Delete(repositoryId uint, envVar *EnvironmentVariable) (*EnvironmentVariable, *http.Response, error) {
	opts := EnvironmentVariablesRequestOptions{RepositoryId: repositoryId}
	id := envVar.Id
	u, err := urlWithOptions(fmt.Sprintf("/settings/env_vars/%s", id), &opts)
	if err != nil {
		return nil, nil, err
	}

	body := environmentVariablePostBody{EnvironmentVariable: *envVar}

	req, err := rs.client.NewRequest("DELETE", u, body, nil)
	if err != nil {
		return nil, nil, err
	}

	var envVarResp getEnvironmentVariableResponse
	resp, err := rs.client.Do(req, &envVarResp)
	if err != nil {
		return nil, resp, err
	}

	return &envVarResp.EnvironmentVariable, resp, err
}
