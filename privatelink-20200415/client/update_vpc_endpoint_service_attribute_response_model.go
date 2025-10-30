// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVpcEndpointServiceAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateVpcEndpointServiceAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateVpcEndpointServiceAttributeResponse
	GetStatusCode() *int32
	SetBody(v *UpdateVpcEndpointServiceAttributeResponseBody) *UpdateVpcEndpointServiceAttributeResponse
	GetBody() *UpdateVpcEndpointServiceAttributeResponseBody
}

type UpdateVpcEndpointServiceAttributeResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateVpcEndpointServiceAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateVpcEndpointServiceAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateVpcEndpointServiceAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateVpcEndpointServiceAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateVpcEndpointServiceAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateVpcEndpointServiceAttributeResponse) GetBody() *UpdateVpcEndpointServiceAttributeResponseBody {
	return s.Body
}

func (s *UpdateVpcEndpointServiceAttributeResponse) SetHeaders(v map[string]*string) *UpdateVpcEndpointServiceAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateVpcEndpointServiceAttributeResponse) SetStatusCode(v int32) *UpdateVpcEndpointServiceAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateVpcEndpointServiceAttributeResponse) SetBody(v *UpdateVpcEndpointServiceAttributeResponseBody) *UpdateVpcEndpointServiceAttributeResponse {
	s.Body = v
	return s
}

func (s *UpdateVpcEndpointServiceAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
