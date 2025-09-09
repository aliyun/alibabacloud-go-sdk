// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUserGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *CreateUserGroupRequest
	GetComment() *string
	SetInstanceId(v string) *CreateUserGroupRequest
	GetInstanceId() *string
	SetRegionId(v string) *CreateUserGroupRequest
	GetRegionId() *string
	SetUserGroupName(v string) *CreateUserGroupRequest
	GetUserGroupName() *string
}

type CreateUserGroupRequest struct {
	// The description of the user group. The description can be up to 500 characters in length.
	//
	// example:
	//
	// comment
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The ID of the bastion host for which you want to create a user group.
	//
	// > You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query the ID of the bastion host.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the bastion host for which you want to create a user group.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the user group that you want to create. This name can be a up to 128 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// group
	UserGroupName *string `json:"UserGroupName,omitempty" xml:"UserGroupName,omitempty"`
}

func (s CreateUserGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateUserGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateUserGroupRequest) GetComment() *string {
	return s.Comment
}

func (s *CreateUserGroupRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateUserGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateUserGroupRequest) GetUserGroupName() *string {
	return s.UserGroupName
}

func (s *CreateUserGroupRequest) SetComment(v string) *CreateUserGroupRequest {
	s.Comment = &v
	return s
}

func (s *CreateUserGroupRequest) SetInstanceId(v string) *CreateUserGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateUserGroupRequest) SetRegionId(v string) *CreateUserGroupRequest {
	s.RegionId = &v
	return s
}

func (s *CreateUserGroupRequest) SetUserGroupName(v string) *CreateUserGroupRequest {
	s.UserGroupName = &v
	return s
}

func (s *CreateUserGroupRequest) Validate() error {
	return dara.Validate(s)
}
