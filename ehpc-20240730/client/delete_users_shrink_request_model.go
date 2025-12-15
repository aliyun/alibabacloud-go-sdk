// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUsersShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DeleteUsersShrinkRequest
	GetClusterId() *string
	SetUserShrink(v string) *DeleteUsersShrinkRequest
	GetUserShrink() *string
}

type DeleteUsersShrinkRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ehpc-hz-FYUr32****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The users that you want to delete.
	//
	// This parameter is required.
	UserShrink *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s DeleteUsersShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteUsersShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteUsersShrinkRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DeleteUsersShrinkRequest) GetUserShrink() *string {
	return s.UserShrink
}

func (s *DeleteUsersShrinkRequest) SetClusterId(v string) *DeleteUsersShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteUsersShrinkRequest) SetUserShrink(v string) *DeleteUsersShrinkRequest {
	s.UserShrink = &v
	return s
}

func (s *DeleteUsersShrinkRequest) Validate() error {
	return dara.Validate(s)
}
