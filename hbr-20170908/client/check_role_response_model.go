// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckRoleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckRoleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckRoleResponse
	GetStatusCode() *int32
	SetBody(v *CheckRoleResponseBody) *CheckRoleResponse
	GetBody() *CheckRoleResponseBody
}

type CheckRoleResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckRoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckRoleResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckRoleResponse) GoString() string {
	return s.String()
}

func (s *CheckRoleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckRoleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckRoleResponse) GetBody() *CheckRoleResponseBody {
	return s.Body
}

func (s *CheckRoleResponse) SetHeaders(v map[string]*string) *CheckRoleResponse {
	s.Headers = v
	return s
}

func (s *CheckRoleResponse) SetStatusCode(v int32) *CheckRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckRoleResponse) SetBody(v *CheckRoleResponseBody) *CheckRoleResponse {
	s.Body = v
	return s
}

func (s *CheckRoleResponse) Validate() error {
	return dara.Validate(s)
}
