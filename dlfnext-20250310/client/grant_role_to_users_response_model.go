// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrantRoleToUsersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GrantRoleToUsersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GrantRoleToUsersResponse
	GetStatusCode() *int32
}

type GrantRoleToUsersResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s GrantRoleToUsersResponse) String() string {
	return dara.Prettify(s)
}

func (s GrantRoleToUsersResponse) GoString() string {
	return s.String()
}

func (s *GrantRoleToUsersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GrantRoleToUsersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GrantRoleToUsersResponse) SetHeaders(v map[string]*string) *GrantRoleToUsersResponse {
	s.Headers = v
	return s
}

func (s *GrantRoleToUsersResponse) SetStatusCode(v int32) *GrantRoleToUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *GrantRoleToUsersResponse) Validate() error {
	return dara.Validate(s)
}
