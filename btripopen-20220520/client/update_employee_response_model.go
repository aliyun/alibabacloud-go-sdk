// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEmployeeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateEmployeeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateEmployeeResponse
	GetStatusCode() *int32
	SetBody(v *UpdateEmployeeResponseBody) *UpdateEmployeeResponse
	GetBody() *UpdateEmployeeResponseBody
}

type UpdateEmployeeResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateEmployeeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateEmployeeResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateEmployeeResponse) GoString() string {
	return s.String()
}

func (s *UpdateEmployeeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateEmployeeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateEmployeeResponse) GetBody() *UpdateEmployeeResponseBody {
	return s.Body
}

func (s *UpdateEmployeeResponse) SetHeaders(v map[string]*string) *UpdateEmployeeResponse {
	s.Headers = v
	return s
}

func (s *UpdateEmployeeResponse) SetStatusCode(v int32) *UpdateEmployeeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateEmployeeResponse) SetBody(v *UpdateEmployeeResponseBody) *UpdateEmployeeResponse {
	s.Body = v
	return s
}

func (s *UpdateEmployeeResponse) Validate() error {
	return dara.Validate(s)
}
