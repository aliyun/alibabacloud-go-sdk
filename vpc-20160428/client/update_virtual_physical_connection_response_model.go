// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVirtualPhysicalConnectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateVirtualPhysicalConnectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateVirtualPhysicalConnectionResponse
	GetStatusCode() *int32
	SetBody(v *UpdateVirtualPhysicalConnectionResponseBody) *UpdateVirtualPhysicalConnectionResponse
	GetBody() *UpdateVirtualPhysicalConnectionResponseBody
}

type UpdateVirtualPhysicalConnectionResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateVirtualPhysicalConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateVirtualPhysicalConnectionResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateVirtualPhysicalConnectionResponse) GoString() string {
	return s.String()
}

func (s *UpdateVirtualPhysicalConnectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateVirtualPhysicalConnectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateVirtualPhysicalConnectionResponse) GetBody() *UpdateVirtualPhysicalConnectionResponseBody {
	return s.Body
}

func (s *UpdateVirtualPhysicalConnectionResponse) SetHeaders(v map[string]*string) *UpdateVirtualPhysicalConnectionResponse {
	s.Headers = v
	return s
}

func (s *UpdateVirtualPhysicalConnectionResponse) SetStatusCode(v int32) *UpdateVirtualPhysicalConnectionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateVirtualPhysicalConnectionResponse) SetBody(v *UpdateVirtualPhysicalConnectionResponseBody) *UpdateVirtualPhysicalConnectionResponse {
	s.Body = v
	return s
}

func (s *UpdateVirtualPhysicalConnectionResponse) Validate() error {
	return dara.Validate(s)
}
