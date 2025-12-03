// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUserGroupsMappingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileSystemId(v string) *CreateUserGroupsMappingRequest
	GetFileSystemId() *string
	SetGroupNames(v []*string) *CreateUserGroupsMappingRequest
	GetGroupNames() []*string
	SetInputRegionId(v string) *CreateUserGroupsMappingRequest
	GetInputRegionId() *string
	SetUserName(v string) *CreateUserGroupsMappingRequest
	GetUserName() *string
}

type CreateUserGroupsMappingRequest struct {
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

func (s CreateUserGroupsMappingRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateUserGroupsMappingRequest) GoString() string {
	return s.String()
}

func (s *CreateUserGroupsMappingRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *CreateUserGroupsMappingRequest) GetGroupNames() []*string {
	return s.GroupNames
}

func (s *CreateUserGroupsMappingRequest) GetInputRegionId() *string {
	return s.InputRegionId
}

func (s *CreateUserGroupsMappingRequest) GetUserName() *string {
	return s.UserName
}

func (s *CreateUserGroupsMappingRequest) SetFileSystemId(v string) *CreateUserGroupsMappingRequest {
	s.FileSystemId = &v
	return s
}

func (s *CreateUserGroupsMappingRequest) SetGroupNames(v []*string) *CreateUserGroupsMappingRequest {
	s.GroupNames = v
	return s
}

func (s *CreateUserGroupsMappingRequest) SetInputRegionId(v string) *CreateUserGroupsMappingRequest {
	s.InputRegionId = &v
	return s
}

func (s *CreateUserGroupsMappingRequest) SetUserName(v string) *CreateUserGroupsMappingRequest {
	s.UserName = &v
	return s
}

func (s *CreateUserGroupsMappingRequest) Validate() error {
	return dara.Validate(s)
}
