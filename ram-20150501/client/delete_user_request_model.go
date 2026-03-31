// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUserName(v string) *DeleteUserRequest
	GetUserName() *string
}

type DeleteUserRequest struct {
	// The name of the RAM user.
	//
	// The name must be 1 to 64 characters in length, and can contain letters, digits, periods (.), hyphens (-), and underscores (_).
	//
	// example:
	//
	// zhangq****
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DeleteUserRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteUserRequest) GoString() string {
	return s.String()
}

func (s *DeleteUserRequest) GetUserName() *string {
	return s.UserName
}

func (s *DeleteUserRequest) SetUserName(v string) *DeleteUserRequest {
	s.UserName = &v
	return s
}

func (s *DeleteUserRequest) Validate() error {
	return dara.Validate(s)
}
