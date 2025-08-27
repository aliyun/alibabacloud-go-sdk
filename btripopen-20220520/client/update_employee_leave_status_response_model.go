// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEmployeeLeaveStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateEmployeeLeaveStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateEmployeeLeaveStatusResponse
	GetStatusCode() *int32
	SetBody(v *UpdateEmployeeLeaveStatusResponseBody) *UpdateEmployeeLeaveStatusResponse
	GetBody() *UpdateEmployeeLeaveStatusResponseBody
}

type UpdateEmployeeLeaveStatusResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateEmployeeLeaveStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateEmployeeLeaveStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateEmployeeLeaveStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateEmployeeLeaveStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateEmployeeLeaveStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateEmployeeLeaveStatusResponse) GetBody() *UpdateEmployeeLeaveStatusResponseBody {
	return s.Body
}

func (s *UpdateEmployeeLeaveStatusResponse) SetHeaders(v map[string]*string) *UpdateEmployeeLeaveStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateEmployeeLeaveStatusResponse) SetStatusCode(v int32) *UpdateEmployeeLeaveStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateEmployeeLeaveStatusResponse) SetBody(v *UpdateEmployeeLeaveStatusResponseBody) *UpdateEmployeeLeaveStatusResponse {
	s.Body = v
	return s
}

func (s *UpdateEmployeeLeaveStatusResponse) Validate() error {
	return dara.Validate(s)
}
