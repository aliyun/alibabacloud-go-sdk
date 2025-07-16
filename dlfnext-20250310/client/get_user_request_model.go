// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUserPrincipal(v string) *GetUserRequest
	GetUserPrincipal() *string
}

type GetUserRequest struct {
	// example:
	//
	// acs:ram::[accountId]:user/user_name
	UserPrincipal *string `json:"userPrincipal,omitempty" xml:"userPrincipal,omitempty"`
}

func (s GetUserRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUserRequest) GoString() string {
	return s.String()
}

func (s *GetUserRequest) GetUserPrincipal() *string {
	return s.UserPrincipal
}

func (s *GetUserRequest) SetUserPrincipal(v string) *GetUserRequest {
	s.UserPrincipal = &v
	return s
}

func (s *GetUserRequest) Validate() error {
	return dara.Validate(s)
}
