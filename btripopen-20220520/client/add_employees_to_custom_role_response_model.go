// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddEmployeesToCustomRoleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddEmployeesToCustomRoleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddEmployeesToCustomRoleResponse
	GetStatusCode() *int32
	SetBody(v *AddEmployeesToCustomRoleResponseBody) *AddEmployeesToCustomRoleResponse
	GetBody() *AddEmployeesToCustomRoleResponseBody
}

type AddEmployeesToCustomRoleResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddEmployeesToCustomRoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddEmployeesToCustomRoleResponse) String() string {
	return dara.Prettify(s)
}

func (s AddEmployeesToCustomRoleResponse) GoString() string {
	return s.String()
}

func (s *AddEmployeesToCustomRoleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddEmployeesToCustomRoleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddEmployeesToCustomRoleResponse) GetBody() *AddEmployeesToCustomRoleResponseBody {
	return s.Body
}

func (s *AddEmployeesToCustomRoleResponse) SetHeaders(v map[string]*string) *AddEmployeesToCustomRoleResponse {
	s.Headers = v
	return s
}

func (s *AddEmployeesToCustomRoleResponse) SetStatusCode(v int32) *AddEmployeesToCustomRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *AddEmployeesToCustomRoleResponse) SetBody(v *AddEmployeesToCustomRoleResponseBody) *AddEmployeesToCustomRoleResponse {
	s.Body = v
	return s
}

func (s *AddEmployeesToCustomRoleResponse) Validate() error {
	return dara.Validate(s)
}
