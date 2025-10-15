// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNetworkAccessEndpointsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListNetworkAccessEndpointsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListNetworkAccessEndpointsResponse
	GetStatusCode() *int32
	SetBody(v *ListNetworkAccessEndpointsResponseBody) *ListNetworkAccessEndpointsResponse
	GetBody() *ListNetworkAccessEndpointsResponseBody
}

type ListNetworkAccessEndpointsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListNetworkAccessEndpointsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListNetworkAccessEndpointsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListNetworkAccessEndpointsResponse) GoString() string {
	return s.String()
}

func (s *ListNetworkAccessEndpointsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListNetworkAccessEndpointsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListNetworkAccessEndpointsResponse) GetBody() *ListNetworkAccessEndpointsResponseBody {
	return s.Body
}

func (s *ListNetworkAccessEndpointsResponse) SetHeaders(v map[string]*string) *ListNetworkAccessEndpointsResponse {
	s.Headers = v
	return s
}

func (s *ListNetworkAccessEndpointsResponse) SetStatusCode(v int32) *ListNetworkAccessEndpointsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListNetworkAccessEndpointsResponse) SetBody(v *ListNetworkAccessEndpointsResponseBody) *ListNetworkAccessEndpointsResponse {
	s.Body = v
	return s
}

func (s *ListNetworkAccessEndpointsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
