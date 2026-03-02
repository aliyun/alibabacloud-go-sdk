// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPrivilegeByRoleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPrivilegeByRoleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPrivilegeByRoleResponse
	GetStatusCode() *int32
	SetBody(v *PrivilegePageResult) *ListPrivilegeByRoleResponse
	GetBody() *PrivilegePageResult
}

type ListPrivilegeByRoleResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PrivilegePageResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPrivilegeByRoleResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPrivilegeByRoleResponse) GoString() string {
	return s.String()
}

func (s *ListPrivilegeByRoleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPrivilegeByRoleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPrivilegeByRoleResponse) GetBody() *PrivilegePageResult {
	return s.Body
}

func (s *ListPrivilegeByRoleResponse) SetHeaders(v map[string]*string) *ListPrivilegeByRoleResponse {
	s.Headers = v
	return s
}

func (s *ListPrivilegeByRoleResponse) SetStatusCode(v int32) *ListPrivilegeByRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPrivilegeByRoleResponse) SetBody(v *PrivilegePageResult) *ListPrivilegeByRoleResponse {
	s.Body = v
	return s
}

func (s *ListPrivilegeByRoleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
