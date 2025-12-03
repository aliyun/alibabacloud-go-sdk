// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUserGroupsMappingShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileSystemId(v string) *DeleteUserGroupsMappingShrinkRequest
	GetFileSystemId() *string
	SetGroupNamesShrink(v string) *DeleteUserGroupsMappingShrinkRequest
	GetGroupNamesShrink() *string
	SetInputRegionId(v string) *DeleteUserGroupsMappingShrinkRequest
	GetInputRegionId() *string
	SetUserName(v string) *DeleteUserGroupsMappingShrinkRequest
	GetUserName() *string
}

type DeleteUserGroupsMappingShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 55C5FFD6-BF99-41BD-9C66-FFF39189****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// example:
	//
	// ["group1","group2"]
	GroupNamesShrink *string `json:"GroupNames,omitempty" xml:"GroupNames,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	InputRegionId *string `json:"InputRegionId,omitempty" xml:"InputRegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// user1
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DeleteUserGroupsMappingShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteUserGroupsMappingShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteUserGroupsMappingShrinkRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *DeleteUserGroupsMappingShrinkRequest) GetGroupNamesShrink() *string {
	return s.GroupNamesShrink
}

func (s *DeleteUserGroupsMappingShrinkRequest) GetInputRegionId() *string {
	return s.InputRegionId
}

func (s *DeleteUserGroupsMappingShrinkRequest) GetUserName() *string {
	return s.UserName
}

func (s *DeleteUserGroupsMappingShrinkRequest) SetFileSystemId(v string) *DeleteUserGroupsMappingShrinkRequest {
	s.FileSystemId = &v
	return s
}

func (s *DeleteUserGroupsMappingShrinkRequest) SetGroupNamesShrink(v string) *DeleteUserGroupsMappingShrinkRequest {
	s.GroupNamesShrink = &v
	return s
}

func (s *DeleteUserGroupsMappingShrinkRequest) SetInputRegionId(v string) *DeleteUserGroupsMappingShrinkRequest {
	s.InputRegionId = &v
	return s
}

func (s *DeleteUserGroupsMappingShrinkRequest) SetUserName(v string) *DeleteUserGroupsMappingShrinkRequest {
	s.UserName = &v
	return s
}

func (s *DeleteUserGroupsMappingShrinkRequest) Validate() error {
	return dara.Validate(s)
}
