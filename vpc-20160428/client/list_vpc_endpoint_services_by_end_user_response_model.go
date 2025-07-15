// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVpcEndpointServicesByEndUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListVpcEndpointServicesByEndUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListVpcEndpointServicesByEndUserResponse
	GetStatusCode() *int32
	SetBody(v *ListVpcEndpointServicesByEndUserResponseBody) *ListVpcEndpointServicesByEndUserResponse
	GetBody() *ListVpcEndpointServicesByEndUserResponseBody
}

type ListVpcEndpointServicesByEndUserResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListVpcEndpointServicesByEndUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListVpcEndpointServicesByEndUserResponse) String() string {
	return dara.Prettify(s)
}

func (s ListVpcEndpointServicesByEndUserResponse) GoString() string {
	return s.String()
}

func (s *ListVpcEndpointServicesByEndUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListVpcEndpointServicesByEndUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListVpcEndpointServicesByEndUserResponse) GetBody() *ListVpcEndpointServicesByEndUserResponseBody {
	return s.Body
}

func (s *ListVpcEndpointServicesByEndUserResponse) SetHeaders(v map[string]*string) *ListVpcEndpointServicesByEndUserResponse {
	s.Headers = v
	return s
}

func (s *ListVpcEndpointServicesByEndUserResponse) SetStatusCode(v int32) *ListVpcEndpointServicesByEndUserResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVpcEndpointServicesByEndUserResponse) SetBody(v *ListVpcEndpointServicesByEndUserResponseBody) *ListVpcEndpointServicesByEndUserResponse {
	s.Body = v
	return s
}

func (s *ListVpcEndpointServicesByEndUserResponse) Validate() error {
	return dara.Validate(s)
}
