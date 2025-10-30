// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVpcEndpointConnectionAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateVpcEndpointConnectionAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateVpcEndpointConnectionAttributeResponse
	GetStatusCode() *int32
	SetBody(v *UpdateVpcEndpointConnectionAttributeResponseBody) *UpdateVpcEndpointConnectionAttributeResponse
	GetBody() *UpdateVpcEndpointConnectionAttributeResponseBody
}

type UpdateVpcEndpointConnectionAttributeResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateVpcEndpointConnectionAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateVpcEndpointConnectionAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateVpcEndpointConnectionAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateVpcEndpointConnectionAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateVpcEndpointConnectionAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateVpcEndpointConnectionAttributeResponse) GetBody() *UpdateVpcEndpointConnectionAttributeResponseBody {
	return s.Body
}

func (s *UpdateVpcEndpointConnectionAttributeResponse) SetHeaders(v map[string]*string) *UpdateVpcEndpointConnectionAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateVpcEndpointConnectionAttributeResponse) SetStatusCode(v int32) *UpdateVpcEndpointConnectionAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateVpcEndpointConnectionAttributeResponse) SetBody(v *UpdateVpcEndpointConnectionAttributeResponseBody) *UpdateVpcEndpointConnectionAttributeResponse {
	s.Body = v
	return s
}

func (s *UpdateVpcEndpointConnectionAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
