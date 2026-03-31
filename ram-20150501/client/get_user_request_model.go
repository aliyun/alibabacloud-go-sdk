// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUserName(v string) *GetUserRequest
	GetUserName() *string
}

type GetUserRequest struct {
	// The username of the RAM user.
	//
	// The username must be 1 to 64 characters in length, and can contain letters, digits, periods (.), hyphens (-), and underscores (_).
	//
	// example:
	//
	// alice
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s GetUserRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUserRequest) GoString() string {
	return s.String()
}

func (s *GetUserRequest) GetUserName() *string {
	return s.UserName
}

func (s *GetUserRequest) SetUserName(v string) *GetUserRequest {
	s.UserName = &v
	return s
}

func (s *GetUserRequest) Validate() error {
	return dara.Validate(s)
}
