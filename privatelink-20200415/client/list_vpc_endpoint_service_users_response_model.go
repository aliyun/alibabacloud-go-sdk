// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVpcEndpointServiceUsersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListVpcEndpointServiceUsersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListVpcEndpointServiceUsersResponse
	GetStatusCode() *int32
	SetBody(v *ListVpcEndpointServiceUsersResponseBody) *ListVpcEndpointServiceUsersResponse
	GetBody() *ListVpcEndpointServiceUsersResponseBody
}

type ListVpcEndpointServiceUsersResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListVpcEndpointServiceUsersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListVpcEndpointServiceUsersResponse) String() string {
	return dara.Prettify(s)
}

func (s ListVpcEndpointServiceUsersResponse) GoString() string {
	return s.String()
}

func (s *ListVpcEndpointServiceUsersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListVpcEndpointServiceUsersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListVpcEndpointServiceUsersResponse) GetBody() *ListVpcEndpointServiceUsersResponseBody {
	return s.Body
}

func (s *ListVpcEndpointServiceUsersResponse) SetHeaders(v map[string]*string) *ListVpcEndpointServiceUsersResponse {
	s.Headers = v
	return s
}

func (s *ListVpcEndpointServiceUsersResponse) SetStatusCode(v int32) *ListVpcEndpointServiceUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVpcEndpointServiceUsersResponse) SetBody(v *ListVpcEndpointServiceUsersResponseBody) *ListVpcEndpointServiceUsersResponse {
	s.Body = v
	return s
}

func (s *ListVpcEndpointServiceUsersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
