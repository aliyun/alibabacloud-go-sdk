// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserIdByUsernameRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUsername(v string) *GetUserIdByUsernameRequest
	GetUsername() *string
}

type GetUserIdByUsernameRequest struct {
	// The username of the account.
	//
	// This parameter is required.
	//
	// example:
	//
	// username_001
	Username *string `json:"username,omitempty" xml:"username,omitempty"`
}

func (s GetUserIdByUsernameRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUserIdByUsernameRequest) GoString() string {
	return s.String()
}

func (s *GetUserIdByUsernameRequest) GetUsername() *string {
	return s.Username
}

func (s *GetUserIdByUsernameRequest) SetUsername(v string) *GetUserIdByUsernameRequest {
	s.Username = &v
	return s
}

func (s *GetUserIdByUsernameRequest) Validate() error {
	return dara.Validate(s)
}
