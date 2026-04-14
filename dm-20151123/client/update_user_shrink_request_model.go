// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUserShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUserShrink(v string) *UpdateUserShrinkRequest
	GetUserShrink() *string
}

type UpdateUserShrinkRequest struct {
	// User Information
	UserShrink *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s UpdateUserShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateUserShrinkRequest) GetUserShrink() *string {
	return s.UserShrink
}

func (s *UpdateUserShrinkRequest) SetUserShrink(v string) *UpdateUserShrinkRequest {
	s.UserShrink = &v
	return s
}

func (s *UpdateUserShrinkRequest) Validate() error {
	return dara.Validate(s)
}
