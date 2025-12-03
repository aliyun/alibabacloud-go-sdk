// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUserGroupsMappingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileSystemId(v string) *DeleteUserGroupsMappingRequest
	GetFileSystemId() *string
	SetGroupNames(v []*string) *DeleteUserGroupsMappingRequest
	GetGroupNames() []*string
	SetInputRegionId(v string) *DeleteUserGroupsMappingRequest
	GetInputRegionId() *string
	SetUserName(v string) *DeleteUserGroupsMappingRequest
	GetUserName() *string
}

type DeleteUserGroupsMappingRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 55C5FFD6-BF99-41BD-9C66-FFF39189****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// example:
	//
	// ["group1","group2"]
	GroupNames []*string `json:"GroupNames,omitempty" xml:"GroupNames,omitempty" type:"Repeated"`
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

func (s DeleteUserGroupsMappingRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteUserGroupsMappingRequest) GoString() string {
	return s.String()
}

func (s *DeleteUserGroupsMappingRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *DeleteUserGroupsMappingRequest) GetGroupNames() []*string {
	return s.GroupNames
}

func (s *DeleteUserGroupsMappingRequest) GetInputRegionId() *string {
	return s.InputRegionId
}

func (s *DeleteUserGroupsMappingRequest) GetUserName() *string {
	return s.UserName
}

func (s *DeleteUserGroupsMappingRequest) SetFileSystemId(v string) *DeleteUserGroupsMappingRequest {
	s.FileSystemId = &v
	return s
}

func (s *DeleteUserGroupsMappingRequest) SetGroupNames(v []*string) *DeleteUserGroupsMappingRequest {
	s.GroupNames = v
	return s
}

func (s *DeleteUserGroupsMappingRequest) SetInputRegionId(v string) *DeleteUserGroupsMappingRequest {
	s.InputRegionId = &v
	return s
}

func (s *DeleteUserGroupsMappingRequest) SetUserName(v string) *DeleteUserGroupsMappingRequest {
	s.UserName = &v
	return s
}

func (s *DeleteUserGroupsMappingRequest) Validate() error {
	return dara.Validate(s)
}
