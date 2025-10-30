// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVpcEndpointServicesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListVpcEndpointServicesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListVpcEndpointServicesResponse
	GetStatusCode() *int32
	SetBody(v *ListVpcEndpointServicesResponseBody) *ListVpcEndpointServicesResponse
	GetBody() *ListVpcEndpointServicesResponseBody
}

type ListVpcEndpointServicesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListVpcEndpointServicesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListVpcEndpointServicesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListVpcEndpointServicesResponse) GoString() string {
	return s.String()
}

func (s *ListVpcEndpointServicesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListVpcEndpointServicesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListVpcEndpointServicesResponse) GetBody() *ListVpcEndpointServicesResponseBody {
	return s.Body
}

func (s *ListVpcEndpointServicesResponse) SetHeaders(v map[string]*string) *ListVpcEndpointServicesResponse {
	s.Headers = v
	return s
}

func (s *ListVpcEndpointServicesResponse) SetStatusCode(v int32) *ListVpcEndpointServicesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVpcEndpointServicesResponse) SetBody(v *ListVpcEndpointServicesResponseBody) *ListVpcEndpointServicesResponse {
	s.Body = v
	return s
}

func (s *ListVpcEndpointServicesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
