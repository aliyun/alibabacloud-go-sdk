// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyHostGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *ModifyHostGroupRequest
	GetComment() *string
	SetHostGroupId(v string) *ModifyHostGroupRequest
	GetHostGroupId() *string
	SetHostGroupName(v string) *ModifyHostGroupRequest
	GetHostGroupName() *string
	SetInstanceId(v string) *ModifyHostGroupRequest
	GetInstanceId() *string
	SetRegionId(v string) *ModifyHostGroupRequest
	GetRegionId() *string
}

type ModifyHostGroupRequest struct {
	// The new remarks of the asset group. The remarks can be up to 500 characters in length.
	//
	// example:
	//
	// comment
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The ID of the asset group that you want to modify.
	//
	// >  You can call the [ListHostGroups](https://help.aliyun.com/document_detail/201307.html) operation to query the ID of the host group.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	HostGroupId *string `json:"HostGroupId,omitempty" xml:"HostGroupId,omitempty"`
	// The new name of the asset group. The name can be up to 128 characters in length.
	//
	// example:
	//
	// Group01
	HostGroupName *string `json:"HostGroupName,omitempty" xml:"HostGroupName,omitempty"`
	// The ID of the bastion host whose asset group you want to modify.
	//
	// > You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query the bastion host ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the bastion host whose asset group you want to modify.
	//
	// >  For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyHostGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyHostGroupRequest) GoString() string {
	return s.String()
}

func (s *ModifyHostGroupRequest) GetComment() *string {
	return s.Comment
}

func (s *ModifyHostGroupRequest) GetHostGroupId() *string {
	return s.HostGroupId
}

func (s *ModifyHostGroupRequest) GetHostGroupName() *string {
	return s.HostGroupName
}

func (s *ModifyHostGroupRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyHostGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyHostGroupRequest) SetComment(v string) *ModifyHostGroupRequest {
	s.Comment = &v
	return s
}

func (s *ModifyHostGroupRequest) SetHostGroupId(v string) *ModifyHostGroupRequest {
	s.HostGroupId = &v
	return s
}

func (s *ModifyHostGroupRequest) SetHostGroupName(v string) *ModifyHostGroupRequest {
	s.HostGroupName = &v
	return s
}

func (s *ModifyHostGroupRequest) SetInstanceId(v string) *ModifyHostGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyHostGroupRequest) SetRegionId(v string) *ModifyHostGroupRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyHostGroupRequest) Validate() error {
	return dara.Validate(s)
}
