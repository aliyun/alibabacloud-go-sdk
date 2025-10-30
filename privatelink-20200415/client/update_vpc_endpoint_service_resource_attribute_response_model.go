// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVpcEndpointServiceResourceAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateVpcEndpointServiceResourceAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateVpcEndpointServiceResourceAttributeResponse
	GetStatusCode() *int32
	SetBody(v *UpdateVpcEndpointServiceResourceAttributeResponseBody) *UpdateVpcEndpointServiceResourceAttributeResponse
	GetBody() *UpdateVpcEndpointServiceResourceAttributeResponseBody
}

type UpdateVpcEndpointServiceResourceAttributeResponse struct {
	Headers    map[string]*string                                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateVpcEndpointServiceResourceAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateVpcEndpointServiceResourceAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateVpcEndpointServiceResourceAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateVpcEndpointServiceResourceAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateVpcEndpointServiceResourceAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateVpcEndpointServiceResourceAttributeResponse) GetBody() *UpdateVpcEndpointServiceResourceAttributeResponseBody {
	return s.Body
}

func (s *UpdateVpcEndpointServiceResourceAttributeResponse) SetHeaders(v map[string]*string) *UpdateVpcEndpointServiceResourceAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateVpcEndpointServiceResourceAttributeResponse) SetStatusCode(v int32) *UpdateVpcEndpointServiceResourceAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateVpcEndpointServiceResourceAttributeResponse) SetBody(v *UpdateVpcEndpointServiceResourceAttributeResponseBody) *UpdateVpcEndpointServiceResourceAttributeResponse {
	s.Body = v
	return s
}

func (s *UpdateVpcEndpointServiceResourceAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
