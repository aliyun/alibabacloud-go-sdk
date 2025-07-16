// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRoleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRoleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRoleResponse
	GetStatusCode() *int32
	SetBody(v *Role) *GetRoleResponse
	GetBody() *Role
}

type GetRoleResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Role              `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRoleResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRoleResponse) GoString() string {
	return s.String()
}

func (s *GetRoleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRoleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRoleResponse) GetBody() *Role {
	return s.Body
}

func (s *GetRoleResponse) SetHeaders(v map[string]*string) *GetRoleResponse {
	s.Headers = v
	return s
}

func (s *GetRoleResponse) SetStatusCode(v int32) *GetRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRoleResponse) SetBody(v *Role) *GetRoleResponse {
	s.Body = v
	return s
}

func (s *GetRoleResponse) Validate() error {
	return dara.Validate(s)
}
