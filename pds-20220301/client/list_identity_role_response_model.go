// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIdentityRoleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListIdentityRoleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListIdentityRoleResponse
	GetStatusCode() *int32
	SetBody(v *BaseRoleMemberResponse) *ListIdentityRoleResponse
	GetBody() *BaseRoleMemberResponse
}

type ListIdentityRoleResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BaseRoleMemberResponse `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListIdentityRoleResponse) String() string {
	return dara.Prettify(s)
}

func (s ListIdentityRoleResponse) GoString() string {
	return s.String()
}

func (s *ListIdentityRoleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListIdentityRoleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListIdentityRoleResponse) GetBody() *BaseRoleMemberResponse {
	return s.Body
}

func (s *ListIdentityRoleResponse) SetHeaders(v map[string]*string) *ListIdentityRoleResponse {
	s.Headers = v
	return s
}

func (s *ListIdentityRoleResponse) SetStatusCode(v int32) *ListIdentityRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *ListIdentityRoleResponse) SetBody(v *BaseRoleMemberResponse) *ListIdentityRoleResponse {
	s.Body = v
	return s
}

func (s *ListIdentityRoleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
