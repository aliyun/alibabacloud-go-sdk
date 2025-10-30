// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVpcEndpointZoneConnectionResourceAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateVpcEndpointZoneConnectionResourceAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateVpcEndpointZoneConnectionResourceAttributeResponse
	GetStatusCode() *int32
	SetBody(v *UpdateVpcEndpointZoneConnectionResourceAttributeResponseBody) *UpdateVpcEndpointZoneConnectionResourceAttributeResponse
	GetBody() *UpdateVpcEndpointZoneConnectionResourceAttributeResponseBody
}

type UpdateVpcEndpointZoneConnectionResourceAttributeResponse struct {
	Headers    map[string]*string                                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateVpcEndpointZoneConnectionResourceAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateVpcEndpointZoneConnectionResourceAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateVpcEndpointZoneConnectionResourceAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateVpcEndpointZoneConnectionResourceAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateVpcEndpointZoneConnectionResourceAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateVpcEndpointZoneConnectionResourceAttributeResponse) GetBody() *UpdateVpcEndpointZoneConnectionResourceAttributeResponseBody {
	return s.Body
}

func (s *UpdateVpcEndpointZoneConnectionResourceAttributeResponse) SetHeaders(v map[string]*string) *UpdateVpcEndpointZoneConnectionResourceAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateVpcEndpointZoneConnectionResourceAttributeResponse) SetStatusCode(v int32) *UpdateVpcEndpointZoneConnectionResourceAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateVpcEndpointZoneConnectionResourceAttributeResponse) SetBody(v *UpdateVpcEndpointZoneConnectionResourceAttributeResponseBody) *UpdateVpcEndpointZoneConnectionResourceAttributeResponse {
	s.Body = v
	return s
}

func (s *UpdateVpcEndpointZoneConnectionResourceAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
