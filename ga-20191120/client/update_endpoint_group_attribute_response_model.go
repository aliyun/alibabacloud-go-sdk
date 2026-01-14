// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEndpointGroupAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateEndpointGroupAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateEndpointGroupAttributeResponse
	GetStatusCode() *int32
	SetBody(v *UpdateEndpointGroupAttributeResponseBody) *UpdateEndpointGroupAttributeResponse
	GetBody() *UpdateEndpointGroupAttributeResponseBody
}

type UpdateEndpointGroupAttributeResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateEndpointGroupAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateEndpointGroupAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateEndpointGroupAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateEndpointGroupAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateEndpointGroupAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateEndpointGroupAttributeResponse) GetBody() *UpdateEndpointGroupAttributeResponseBody {
	return s.Body
}

func (s *UpdateEndpointGroupAttributeResponse) SetHeaders(v map[string]*string) *UpdateEndpointGroupAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateEndpointGroupAttributeResponse) SetStatusCode(v int32) *UpdateEndpointGroupAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateEndpointGroupAttributeResponse) SetBody(v *UpdateEndpointGroupAttributeResponseBody) *UpdateEndpointGroupAttributeResponse {
	s.Body = v
	return s
}

func (s *UpdateEndpointGroupAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
