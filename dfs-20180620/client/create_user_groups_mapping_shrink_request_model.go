// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUserGroupsMappingShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileSystemId(v string) *CreateUserGroupsMappingShrinkRequest
	GetFileSystemId() *string
	SetGroupNamesShrink(v string) *CreateUserGroupsMappingShrinkRequest
	GetGroupNamesShrink() *string
	SetInputRegionId(v string) *CreateUserGroupsMappingShrinkRequest
	GetInputRegionId() *string
	SetUserName(v string) *CreateUserGroupsMappingShrinkRequest
	GetUserName() *string
}

type CreateUserGroupsMappingShrinkRequest struct {
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

func (s CreateUserGroupsMappingShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateUserGroupsMappingShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateUserGroupsMappingShrinkRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *CreateUserGroupsMappingShrinkRequest) GetGroupNamesShrink() *string {
	return s.GroupNamesShrink
}

func (s *CreateUserGroupsMappingShrinkRequest) GetInputRegionId() *string {
	return s.InputRegionId
}

func (s *CreateUserGroupsMappingShrinkRequest) GetUserName() *string {
	return s.UserName
}

func (s *CreateUserGroupsMappingShrinkRequest) SetFileSystemId(v string) *CreateUserGroupsMappingShrinkRequest {
	s.FileSystemId = &v
	return s
}

func (s *CreateUserGroupsMappingShrinkRequest) SetGroupNamesShrink(v string) *CreateUserGroupsMappingShrinkRequest {
	s.GroupNamesShrink = &v
	return s
}

func (s *CreateUserGroupsMappingShrinkRequest) SetInputRegionId(v string) *CreateUserGroupsMappingShrinkRequest {
	s.InputRegionId = &v
	return s
}

func (s *CreateUserGroupsMappingShrinkRequest) SetUserName(v string) *CreateUserGroupsMappingShrinkRequest {
	s.UserName = &v
	return s
}

func (s *CreateUserGroupsMappingShrinkRequest) Validate() error {
	return dara.Validate(s)
}
