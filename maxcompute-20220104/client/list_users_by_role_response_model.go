// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUsersByRoleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListUsersByRoleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListUsersByRoleResponse
	GetStatusCode() *int32
	SetBody(v *ListUsersByRoleResponseBody) *ListUsersByRoleResponse
	GetBody() *ListUsersByRoleResponseBody
}

type ListUsersByRoleResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUsersByRoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUsersByRoleResponse) String() string {
	return dara.Prettify(s)
}

func (s ListUsersByRoleResponse) GoString() string {
	return s.String()
}

func (s *ListUsersByRoleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListUsersByRoleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListUsersByRoleResponse) GetBody() *ListUsersByRoleResponseBody {
	return s.Body
}

func (s *ListUsersByRoleResponse) SetHeaders(v map[string]*string) *ListUsersByRoleResponse {
	s.Headers = v
	return s
}

func (s *ListUsersByRoleResponse) SetStatusCode(v int32) *ListUsersByRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUsersByRoleResponse) SetBody(v *ListUsersByRoleResponseBody) *ListUsersByRoleResponse {
	s.Body = v
	return s
}

func (s *ListUsersByRoleResponse) Validate() error {
	return dara.Validate(s)
}
