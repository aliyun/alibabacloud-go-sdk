// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNetworkAccessPathsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListNetworkAccessPathsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListNetworkAccessPathsResponse
	GetStatusCode() *int32
	SetBody(v *ListNetworkAccessPathsResponseBody) *ListNetworkAccessPathsResponse
	GetBody() *ListNetworkAccessPathsResponseBody
}

type ListNetworkAccessPathsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListNetworkAccessPathsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListNetworkAccessPathsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListNetworkAccessPathsResponse) GoString() string {
	return s.String()
}

func (s *ListNetworkAccessPathsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListNetworkAccessPathsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListNetworkAccessPathsResponse) GetBody() *ListNetworkAccessPathsResponseBody {
	return s.Body
}

func (s *ListNetworkAccessPathsResponse) SetHeaders(v map[string]*string) *ListNetworkAccessPathsResponse {
	s.Headers = v
	return s
}

func (s *ListNetworkAccessPathsResponse) SetStatusCode(v int32) *ListNetworkAccessPathsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListNetworkAccessPathsResponse) SetBody(v *ListNetworkAccessPathsResponseBody) *ListNetworkAccessPathsResponse {
	s.Body = v
	return s
}

func (s *ListNetworkAccessPathsResponse) Validate() error {
	return dara.Validate(s)
}
