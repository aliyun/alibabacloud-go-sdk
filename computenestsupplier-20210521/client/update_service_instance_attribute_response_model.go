// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateServiceInstanceAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateServiceInstanceAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateServiceInstanceAttributeResponse
	GetStatusCode() *int32
	SetBody(v *UpdateServiceInstanceAttributeResponseBody) *UpdateServiceInstanceAttributeResponse
	GetBody() *UpdateServiceInstanceAttributeResponseBody
}

type UpdateServiceInstanceAttributeResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateServiceInstanceAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateServiceInstanceAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceInstanceAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateServiceInstanceAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateServiceInstanceAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateServiceInstanceAttributeResponse) GetBody() *UpdateServiceInstanceAttributeResponseBody {
	return s.Body
}

func (s *UpdateServiceInstanceAttributeResponse) SetHeaders(v map[string]*string) *UpdateServiceInstanceAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateServiceInstanceAttributeResponse) SetStatusCode(v int32) *UpdateServiceInstanceAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateServiceInstanceAttributeResponse) SetBody(v *UpdateServiceInstanceAttributeResponseBody) *UpdateServiceInstanceAttributeResponse {
	s.Body = v
	return s
}

func (s *UpdateServiceInstanceAttributeResponse) Validate() error {
	return dara.Validate(s)
}
