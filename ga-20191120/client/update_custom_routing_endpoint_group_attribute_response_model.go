// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCustomRoutingEndpointGroupAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateCustomRoutingEndpointGroupAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateCustomRoutingEndpointGroupAttributeResponse
	GetStatusCode() *int32
	SetBody(v *UpdateCustomRoutingEndpointGroupAttributeResponseBody) *UpdateCustomRoutingEndpointGroupAttributeResponse
	GetBody() *UpdateCustomRoutingEndpointGroupAttributeResponseBody
}

type UpdateCustomRoutingEndpointGroupAttributeResponse struct {
	Headers    map[string]*string                                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCustomRoutingEndpointGroupAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCustomRoutingEndpointGroupAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomRoutingEndpointGroupAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateCustomRoutingEndpointGroupAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateCustomRoutingEndpointGroupAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateCustomRoutingEndpointGroupAttributeResponse) GetBody() *UpdateCustomRoutingEndpointGroupAttributeResponseBody {
	return s.Body
}

func (s *UpdateCustomRoutingEndpointGroupAttributeResponse) SetHeaders(v map[string]*string) *UpdateCustomRoutingEndpointGroupAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateCustomRoutingEndpointGroupAttributeResponse) SetStatusCode(v int32) *UpdateCustomRoutingEndpointGroupAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCustomRoutingEndpointGroupAttributeResponse) SetBody(v *UpdateCustomRoutingEndpointGroupAttributeResponseBody) *UpdateCustomRoutingEndpointGroupAttributeResponse {
	s.Body = v
	return s
}

func (s *UpdateCustomRoutingEndpointGroupAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
