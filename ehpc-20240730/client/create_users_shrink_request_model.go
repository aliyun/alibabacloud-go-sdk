// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUsersShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *CreateUsersShrinkRequest
	GetClusterId() *string
	SetUserShrink(v string) *CreateUsersShrinkRequest
	GetUserShrink() *string
}

type CreateUsersShrinkRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ehpc-hz-FYUr32****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The users that you want to add.
	UserShrink *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s CreateUsersShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateUsersShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateUsersShrinkRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreateUsersShrinkRequest) GetUserShrink() *string {
	return s.UserShrink
}

func (s *CreateUsersShrinkRequest) SetClusterId(v string) *CreateUsersShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateUsersShrinkRequest) SetUserShrink(v string) *CreateUsersShrinkRequest {
	s.UserShrink = &v
	return s
}

func (s *CreateUsersShrinkRequest) Validate() error {
	return dara.Validate(s)
}
