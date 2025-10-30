// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVpcEndpointSecurityGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListVpcEndpointSecurityGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListVpcEndpointSecurityGroupsResponse
	GetStatusCode() *int32
	SetBody(v *ListVpcEndpointSecurityGroupsResponseBody) *ListVpcEndpointSecurityGroupsResponse
	GetBody() *ListVpcEndpointSecurityGroupsResponseBody
}

type ListVpcEndpointSecurityGroupsResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListVpcEndpointSecurityGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListVpcEndpointSecurityGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListVpcEndpointSecurityGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListVpcEndpointSecurityGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListVpcEndpointSecurityGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListVpcEndpointSecurityGroupsResponse) GetBody() *ListVpcEndpointSecurityGroupsResponseBody {
	return s.Body
}

func (s *ListVpcEndpointSecurityGroupsResponse) SetHeaders(v map[string]*string) *ListVpcEndpointSecurityGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListVpcEndpointSecurityGroupsResponse) SetStatusCode(v int32) *ListVpcEndpointSecurityGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVpcEndpointSecurityGroupsResponse) SetBody(v *ListVpcEndpointSecurityGroupsResponseBody) *ListVpcEndpointSecurityGroupsResponse {
	s.Body = v
	return s
}

func (s *ListVpcEndpointSecurityGroupsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
