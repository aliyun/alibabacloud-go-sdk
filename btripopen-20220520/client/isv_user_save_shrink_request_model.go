// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIsvUserSaveShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUserListShrink(v string) *IsvUserSaveShrinkRequest
	GetUserListShrink() *string
}

type IsvUserSaveShrinkRequest struct {
	UserListShrink *string `json:"user_list,omitempty" xml:"user_list,omitempty"`
}

func (s IsvUserSaveShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s IsvUserSaveShrinkRequest) GoString() string {
	return s.String()
}

func (s *IsvUserSaveShrinkRequest) GetUserListShrink() *string {
	return s.UserListShrink
}

func (s *IsvUserSaveShrinkRequest) SetUserListShrink(v string) *IsvUserSaveShrinkRequest {
	s.UserListShrink = &v
	return s
}

func (s *IsvUserSaveShrinkRequest) Validate() error {
	return dara.Validate(s)
}
