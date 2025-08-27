// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCustomRoleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateCustomRoleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateCustomRoleResponse
	GetStatusCode() *int32
	SetBody(v *UpdateCustomRoleResponseBody) *UpdateCustomRoleResponse
	GetBody() *UpdateCustomRoleResponseBody
}

type UpdateCustomRoleResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCustomRoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCustomRoleResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomRoleResponse) GoString() string {
	return s.String()
}

func (s *UpdateCustomRoleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateCustomRoleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateCustomRoleResponse) GetBody() *UpdateCustomRoleResponseBody {
	return s.Body
}

func (s *UpdateCustomRoleResponse) SetHeaders(v map[string]*string) *UpdateCustomRoleResponse {
	s.Headers = v
	return s
}

func (s *UpdateCustomRoleResponse) SetStatusCode(v int32) *UpdateCustomRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCustomRoleResponse) SetBody(v *UpdateCustomRoleResponseBody) *UpdateCustomRoleResponse {
	s.Body = v
	return s
}

func (s *UpdateCustomRoleResponse) Validate() error {
	return dara.Validate(s)
}
