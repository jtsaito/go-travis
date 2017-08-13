package travis

import "net/http"

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
