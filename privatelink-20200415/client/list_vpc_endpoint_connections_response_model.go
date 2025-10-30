// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVpcEndpointConnectionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListVpcEndpointConnectionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListVpcEndpointConnectionsResponse
	GetStatusCode() *int32
	SetBody(v *ListVpcEndpointConnectionsResponseBody) *ListVpcEndpointConnectionsResponse
	GetBody() *ListVpcEndpointConnectionsResponseBody
}

type ListVpcEndpointConnectionsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListVpcEndpointConnectionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListVpcEndpointConnectionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListVpcEndpointConnectionsResponse) GoString() string {
	return s.String()
}

func (s *ListVpcEndpointConnectionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListVpcEndpointConnectionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListVpcEndpointConnectionsResponse) GetBody() *ListVpcEndpointConnectionsResponseBody {
	return s.Body
}

func (s *ListVpcEndpointConnectionsResponse) SetHeaders(v map[string]*string) *ListVpcEndpointConnectionsResponse {
	s.Headers = v
	return s
}

func (s *ListVpcEndpointConnectionsResponse) SetStatusCode(v int32) *ListVpcEndpointConnectionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVpcEndpointConnectionsResponse) SetBody(v *ListVpcEndpointConnectionsResponseBody) *ListVpcEndpointConnectionsResponse {
	s.Body = v
	return s
}

func (s *ListVpcEndpointConnectionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
