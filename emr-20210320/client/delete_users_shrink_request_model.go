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
	SetRegionId(v string) *DeleteUsersShrinkRequest
	GetRegionId() *string
	SetUserNamesShrink(v string) *DeleteUsersShrinkRequest
	GetUserNamesShrink() *string
}

type DeleteUsersShrinkRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// c-b933c5aac8fe****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The usernames. Number of elements in the array: 0 to 10.
	UserNamesShrink *string `json:"UserNames,omitempty" xml:"UserNames,omitempty"`
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

func (s *DeleteUsersShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteUsersShrinkRequest) GetUserNamesShrink() *string {
	return s.UserNamesShrink
}

func (s *DeleteUsersShrinkRequest) SetClusterId(v string) *DeleteUsersShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteUsersShrinkRequest) SetRegionId(v string) *DeleteUsersShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteUsersShrinkRequest) SetUserNamesShrink(v string) *DeleteUsersShrinkRequest {
	s.UserNamesShrink = &v
	return s
}

func (s *DeleteUsersShrinkRequest) Validate() error {
	return dara.Validate(s)
}
