// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomRoleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCustomRoleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCustomRoleResponse
	GetStatusCode() *int32
	SetBody(v *CreateCustomRoleResponseBody) *CreateCustomRoleResponse
	GetBody() *CreateCustomRoleResponseBody
}

type CreateCustomRoleResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCustomRoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCustomRoleResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomRoleResponse) GoString() string {
	return s.String()
}

func (s *CreateCustomRoleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCustomRoleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCustomRoleResponse) GetBody() *CreateCustomRoleResponseBody {
	return s.Body
}

func (s *CreateCustomRoleResponse) SetHeaders(v map[string]*string) *CreateCustomRoleResponse {
	s.Headers = v
	return s
}

func (s *CreateCustomRoleResponse) SetStatusCode(v int32) *CreateCustomRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCustomRoleResponse) SetBody(v *CreateCustomRoleResponseBody) *CreateCustomRoleResponse {
	s.Body = v
	return s
}

func (s *CreateCustomRoleResponse) Validate() error {
	return dara.Validate(s)
}
