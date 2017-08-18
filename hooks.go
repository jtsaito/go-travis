// Copyright (c) 2015 Ableton AG, Berlin. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// Fragments of this file have been copied from the go-github (https://github.com/google/go-github)
// project, and is therefore licensed under the following copyright:
// Copyright 2013 The go-github AUTHORS. All rights reserved.

package travis

import (
	"fmt"
	"net/http"
)

// HooksService handles communication with the builds
// related methods of the Travis CI API.
type HooksService struct {
	client *Client
}

// Hook represents a Travis CI hook
type Hook struct {
	Id          uint   `json:"id"`
	Name        string `json:"name"`
	OwnerName   string `json:"owner_name"`
	Description string `json:"description"`
	Active      bool   `json:"active"`
	Private     bool   `json:"private"`
	Admin       bool   `json:"admin"`
}

// HookListOptions specifies the parameters to the HooksService.List method.
type HookListOptions struct {
	All bool `url:"all,omitempty"`

	// Define an order
	Order string `url:"order,omitempty"`

	// filter by owner name
	OwnerName string `url:"owner_name,omitempty"`
}

// listHooksResponse represents the response of a call
// to the Travis CI list builds endpoint.
type listHooksResponse struct {
	Hooks []Hook `json:"hooks"`
}

// hookPutBody represents the post body of a hook
// to the Travis CI list builds endpoint.
type hookPutBody struct {
	Hook *Hook `json:"hook"`
}

// HookPutResult represents the result of a put to hook
// to the Travis CI list builds endpoint.
type HookPutResult struct {
	Result bool `json:"result"`
}

// List hooks using the provided options.
//
// Travis CI API docs: https://docs.travis-ci.com/api/?http#hooks
func (rs *HooksService) List(opt *HookListOptions) ([]Hook, *http.Response, error) {
	u, err := urlWithOptions("/hooks", opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := rs.client.NewRequest("GET", u, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var hooksResp listHooksResponse
	resp, err := rs.client.Do(req, &hooksResp)
	if err != nil {
		return nil, resp, err
	}

	return hooksResp.Hooks, resp, err
}

// Update updates a hook.
//
// Travis CI API docs: https://docs.travis-ci.com/api/?http#hooks
func (rs *HooksService) Update(hook *Hook) (*HookPutResult, *http.Response, error) {
	u, err := urlWithOptions(fmt.Sprintf("/hooks/%d", hook.Id), nil)
	if err != nil {
		return nil, nil, err
	}

	body := hookPutBody{Hook: hook}

	req, err := rs.client.NewRequest("PUT", u, body, nil)
	if err != nil {
		return nil, nil, err
	}

	var hookPutRes HookPutResult
	resp, err := rs.client.Do(req, &hookPutRes)
	if err != nil {
		return nil, resp, err
	}

	return &hookPutRes, resp, err
}
