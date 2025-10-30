// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVpcEndpointAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateVpcEndpointAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateVpcEndpointAttributeResponse
	GetStatusCode() *int32
	SetBody(v *UpdateVpcEndpointAttributeResponseBody) *UpdateVpcEndpointAttributeResponse
	GetBody() *UpdateVpcEndpointAttributeResponseBody
}

type UpdateVpcEndpointAttributeResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateVpcEndpointAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateVpcEndpointAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateVpcEndpointAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateVpcEndpointAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateVpcEndpointAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateVpcEndpointAttributeResponse) GetBody() *UpdateVpcEndpointAttributeResponseBody {
	return s.Body
}

func (s *UpdateVpcEndpointAttributeResponse) SetHeaders(v map[string]*string) *UpdateVpcEndpointAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateVpcEndpointAttributeResponse) SetStatusCode(v int32) *UpdateVpcEndpointAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateVpcEndpointAttributeResponse) SetBody(v *UpdateVpcEndpointAttributeResponseBody) *UpdateVpcEndpointAttributeResponse {
	s.Body = v
	return s
}

func (s *UpdateVpcEndpointAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
