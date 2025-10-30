// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddUserToVpcEndpointServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddUserToVpcEndpointServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddUserToVpcEndpointServiceResponse
	GetStatusCode() *int32
	SetBody(v *AddUserToVpcEndpointServiceResponseBody) *AddUserToVpcEndpointServiceResponse
	GetBody() *AddUserToVpcEndpointServiceResponseBody
}

type AddUserToVpcEndpointServiceResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddUserToVpcEndpointServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddUserToVpcEndpointServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s AddUserToVpcEndpointServiceResponse) GoString() string {
	return s.String()
}

func (s *AddUserToVpcEndpointServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddUserToVpcEndpointServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddUserToVpcEndpointServiceResponse) GetBody() *AddUserToVpcEndpointServiceResponseBody {
	return s.Body
}

func (s *AddUserToVpcEndpointServiceResponse) SetHeaders(v map[string]*string) *AddUserToVpcEndpointServiceResponse {
	s.Headers = v
	return s
}

func (s *AddUserToVpcEndpointServiceResponse) SetStatusCode(v int32) *AddUserToVpcEndpointServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *AddUserToVpcEndpointServiceResponse) SetBody(v *AddUserToVpcEndpointServiceResponseBody) *AddUserToVpcEndpointServiceResponse {
	s.Body = v
	return s
}

func (s *AddUserToVpcEndpointServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
