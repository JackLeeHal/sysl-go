// Code generated by sysl DO NOT EDIT.
package downstream

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/anz-bank/sysl-go/common"
	"github.com/anz-bank/sysl-go/restlib"
	"github.com/anz-bank/sysl-go/validator"
)

// Service interface for Downstream
type Service interface {
	GetServiceDocsList(ctx context.Context, req *GetServiceDocsListRequest) (*ServiceDoc, error)
}

// Client for Downstream API
type Client struct {
	client *http.Client
	url    string
}

// NewClient for Downstream
func NewClient(client *http.Client, serviceURL string) *Client {
	return &Client{client, serviceURL}
}

// GetServiceDocsList ...
func (s *Client) GetServiceDocsList(ctx context.Context, req *GetServiceDocsListRequest) (*ServiceDoc, error) {
	required := []string{}
	var okResponse ServiceDoc

	var errorResponse Status

	u, err := url.Parse(fmt.Sprintf("%s/service-docs", s.url))
	if err != nil {
		return nil, common.CreateError(ctx, common.InternalError, "failed to parse url", err)
	}

	result, err := restlib.DoHTTPRequest(ctx, s.client, "GET", u.String(), nil, required, &okResponse, &errorResponse)
	if err != nil {
		response, ok := err.(*restlib.HTTPResult)
		if !ok {
			return nil, common.CreateError(ctx, common.DownstreamUnavailableError, "call failed: Downstream <- GET "+u.String(), err)
		}

		return nil, common.CreateDownstreamError(ctx, common.DownstreamResponseError, response.HTTPResponse, response.Body, &errorResponse)
	}

	if result.HTTPResponse.StatusCode == http.StatusUnauthorized {
		return nil, common.CreateDownstreamError(ctx, common.DownstreamUnauthorizedError, result.HTTPResponse, result.Body, nil)
	}

	OkServiceDocResponse, ok := result.Response.(*ServiceDoc)
	if ok {
		valErr := validator.Validate(OkServiceDocResponse)
		if valErr != nil {
			return nil, common.CreateDownstreamError(ctx, common.DownstreamUnexpectedResponseError, result.HTTPResponse, result.Body, valErr)
		}

		return OkServiceDocResponse, nil
	}

	return nil, common.CreateDownstreamError(ctx, common.DownstreamUnexpectedResponseError, result.HTTPResponse, result.Body, nil)
}