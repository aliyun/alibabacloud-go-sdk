// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVpcEndpointServiceResourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListVpcEndpointServiceResourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListVpcEndpointServiceResourcesResponse
	GetStatusCode() *int32
	SetBody(v *ListVpcEndpointServiceResourcesResponseBody) *ListVpcEndpointServiceResourcesResponse
	GetBody() *ListVpcEndpointServiceResourcesResponseBody
}

type ListVpcEndpointServiceResourcesResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListVpcEndpointServiceResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListVpcEndpointServiceResourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListVpcEndpointServiceResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListVpcEndpointServiceResourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListVpcEndpointServiceResourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListVpcEndpointServiceResourcesResponse) GetBody() *ListVpcEndpointServiceResourcesResponseBody {
	return s.Body
}

func (s *ListVpcEndpointServiceResourcesResponse) SetHeaders(v map[string]*string) *ListVpcEndpointServiceResourcesResponse {
	s.Headers = v
	return s
}

func (s *ListVpcEndpointServiceResourcesResponse) SetStatusCode(v int32) *ListVpcEndpointServiceResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVpcEndpointServiceResourcesResponse) SetBody(v *ListVpcEndpointServiceResourcesResponseBody) *ListVpcEndpointServiceResourcesResponse {
	s.Body = v
	return s
}

func (s *ListVpcEndpointServiceResourcesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
