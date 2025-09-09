// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHostGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *CreateHostGroupRequest
	GetComment() *string
	SetHostGroupName(v string) *CreateHostGroupRequest
	GetHostGroupName() *string
	SetInstanceId(v string) *CreateHostGroupRequest
	GetInstanceId() *string
	SetRegionId(v string) *CreateHostGroupRequest
	GetRegionId() *string
}

type CreateHostGroupRequest struct {
	// The remarks of the asset group. The remarks can be up to 500 characters in length.
	//
	// example:
	//
	// Local host group.
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The name of the asset group. The name can be up to 128 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// HostGroup01
	HostGroupName *string `json:"HostGroupName,omitempty" xml:"HostGroupName,omitempty"`
	// The ID of the bastion host on which you want to create an asset group.
	//
	// >  You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query the bastion host ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the bastion host on which you want to create an asset group.
	//
	// >  For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateHostGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateHostGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateHostGroupRequest) GetComment() *string {
	return s.Comment
}

func (s *CreateHostGroupRequest) GetHostGroupName() *string {
	return s.HostGroupName
}

func (s *CreateHostGroupRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateHostGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateHostGroupRequest) SetComment(v string) *CreateHostGroupRequest {
	s.Comment = &v
	return s
}

func (s *CreateHostGroupRequest) SetHostGroupName(v string) *CreateHostGroupRequest {
	s.HostGroupName = &v
	return s
}

func (s *CreateHostGroupRequest) SetInstanceId(v string) *CreateHostGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateHostGroupRequest) SetRegionId(v string) *CreateHostGroupRequest {
	s.RegionId = &v
	return s
}

func (s *CreateHostGroupRequest) Validate() error {
	return dara.Validate(s)
}
