// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNetworkAccessEndpointAvailableRegionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListNetworkAccessEndpointAvailableRegionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListNetworkAccessEndpointAvailableRegionsResponse
	GetStatusCode() *int32
	SetBody(v *ListNetworkAccessEndpointAvailableRegionsResponseBody) *ListNetworkAccessEndpointAvailableRegionsResponse
	GetBody() *ListNetworkAccessEndpointAvailableRegionsResponseBody
}

type ListNetworkAccessEndpointAvailableRegionsResponse struct {
	Headers    map[string]*string                                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListNetworkAccessEndpointAvailableRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListNetworkAccessEndpointAvailableRegionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListNetworkAccessEndpointAvailableRegionsResponse) GoString() string {
	return s.String()
}

func (s *ListNetworkAccessEndpointAvailableRegionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListNetworkAccessEndpointAvailableRegionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListNetworkAccessEndpointAvailableRegionsResponse) GetBody() *ListNetworkAccessEndpointAvailableRegionsResponseBody {
	return s.Body
}

func (s *ListNetworkAccessEndpointAvailableRegionsResponse) SetHeaders(v map[string]*string) *ListNetworkAccessEndpointAvailableRegionsResponse {
	s.Headers = v
	return s
}

func (s *ListNetworkAccessEndpointAvailableRegionsResponse) SetStatusCode(v int32) *ListNetworkAccessEndpointAvailableRegionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListNetworkAccessEndpointAvailableRegionsResponse) SetBody(v *ListNetworkAccessEndpointAvailableRegionsResponseBody) *ListNetworkAccessEndpointAvailableRegionsResponse {
	s.Body = v
	return s
}

func (s *ListNetworkAccessEndpointAvailableRegionsResponse) Validate() error {
	return dara.Validate(s)
}
