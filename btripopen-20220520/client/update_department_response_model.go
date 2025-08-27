// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDepartmentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDepartmentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDepartmentResponse
	GetStatusCode() *int32
	SetBody(v *UpdateDepartmentResponseBody) *UpdateDepartmentResponse
	GetBody() *UpdateDepartmentResponseBody
}

type UpdateDepartmentResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDepartmentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDepartmentResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDepartmentResponse) GoString() string {
	return s.String()
}

func (s *UpdateDepartmentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDepartmentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDepartmentResponse) GetBody() *UpdateDepartmentResponseBody {
	return s.Body
}

func (s *UpdateDepartmentResponse) SetHeaders(v map[string]*string) *UpdateDepartmentResponse {
	s.Headers = v
	return s
}

func (s *UpdateDepartmentResponse) SetStatusCode(v int32) *UpdateDepartmentResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDepartmentResponse) SetBody(v *UpdateDepartmentResponseBody) *UpdateDepartmentResponse {
	s.Body = v
	return s
}

func (s *UpdateDepartmentResponse) Validate() error {
	return dara.Validate(s)
}
