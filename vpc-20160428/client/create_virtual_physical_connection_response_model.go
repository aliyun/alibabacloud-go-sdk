// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVirtualPhysicalConnectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateVirtualPhysicalConnectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateVirtualPhysicalConnectionResponse
	GetStatusCode() *int32
	SetBody(v *CreateVirtualPhysicalConnectionResponseBody) *CreateVirtualPhysicalConnectionResponse
	GetBody() *CreateVirtualPhysicalConnectionResponseBody
}

type CreateVirtualPhysicalConnectionResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateVirtualPhysicalConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateVirtualPhysicalConnectionResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateVirtualPhysicalConnectionResponse) GoString() string {
	return s.String()
}

func (s *CreateVirtualPhysicalConnectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateVirtualPhysicalConnectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateVirtualPhysicalConnectionResponse) GetBody() *CreateVirtualPhysicalConnectionResponseBody {
	return s.Body
}

func (s *CreateVirtualPhysicalConnectionResponse) SetHeaders(v map[string]*string) *CreateVirtualPhysicalConnectionResponse {
	s.Headers = v
	return s
}

func (s *CreateVirtualPhysicalConnectionResponse) SetStatusCode(v int32) *CreateVirtualPhysicalConnectionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateVirtualPhysicalConnectionResponse) SetBody(v *CreateVirtualPhysicalConnectionResponseBody) *CreateVirtualPhysicalConnectionResponse {
	s.Body = v
	return s
}

func (s *CreateVirtualPhysicalConnectionResponse) Validate() error {
	return dara.Validate(s)
}
