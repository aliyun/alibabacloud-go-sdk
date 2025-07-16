// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeRoleFromUsersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RevokeRoleFromUsersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RevokeRoleFromUsersResponse
	GetStatusCode() *int32
}

type RevokeRoleFromUsersResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s RevokeRoleFromUsersResponse) String() string {
	return dara.Prettify(s)
}

func (s RevokeRoleFromUsersResponse) GoString() string {
	return s.String()
}

func (s *RevokeRoleFromUsersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RevokeRoleFromUsersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RevokeRoleFromUsersResponse) SetHeaders(v map[string]*string) *RevokeRoleFromUsersResponse {
	s.Headers = v
	return s
}

func (s *RevokeRoleFromUsersResponse) SetStatusCode(v int32) *RevokeRoleFromUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *RevokeRoleFromUsersResponse) Validate() error {
	return dara.Validate(s)
}
