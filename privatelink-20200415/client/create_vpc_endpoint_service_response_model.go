// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVpcEndpointServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateVpcEndpointServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateVpcEndpointServiceResponse
	GetStatusCode() *int32
	SetBody(v *CreateVpcEndpointServiceResponseBody) *CreateVpcEndpointServiceResponse
	GetBody() *CreateVpcEndpointServiceResponseBody
}

type CreateVpcEndpointServiceResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateVpcEndpointServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateVpcEndpointServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateVpcEndpointServiceResponse) GoString() string {
	return s.String()
}

func (s *CreateVpcEndpointServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateVpcEndpointServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateVpcEndpointServiceResponse) GetBody() *CreateVpcEndpointServiceResponseBody {
	return s.Body
}

func (s *CreateVpcEndpointServiceResponse) SetHeaders(v map[string]*string) *CreateVpcEndpointServiceResponse {
	s.Headers = v
	return s
}

func (s *CreateVpcEndpointServiceResponse) SetStatusCode(v int32) *CreateVpcEndpointServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateVpcEndpointServiceResponse) SetBody(v *CreateVpcEndpointServiceResponseBody) *CreateVpcEndpointServiceResponse {
	s.Body = v
	return s
}

func (s *CreateVpcEndpointServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
