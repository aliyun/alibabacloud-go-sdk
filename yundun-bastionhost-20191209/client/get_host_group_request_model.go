// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHostGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHostGroupId(v string) *GetHostGroupRequest
	GetHostGroupId() *string
	SetInstanceId(v string) *GetHostGroupRequest
	GetInstanceId() *string
	SetRegionId(v string) *GetHostGroupRequest
	GetRegionId() *string
}

type GetHostGroupRequest struct {
	// The ID of the asset group to query.
	//
	// > You can call the [ListHostGroups](https://help.aliyun.com/document_detail/204509.html) operation to query the asset group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	HostGroupId *string `json:"HostGroupId,omitempty" xml:"HostGroupId,omitempty"`
	// The ID of the bastion host whose asset group you want to query.
	//
	// > You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query the bastion host ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the bastion host whose asset group you want to query.
	//
	// > For more information about the mapping between region IDs and region names, [see Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetHostGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s GetHostGroupRequest) GoString() string {
	return s.String()
}

func (s *GetHostGroupRequest) GetHostGroupId() *string {
	return s.HostGroupId
}

func (s *GetHostGroupRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetHostGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetHostGroupRequest) SetHostGroupId(v string) *GetHostGroupRequest {
	s.HostGroupId = &v
	return s
}

func (s *GetHostGroupRequest) SetInstanceId(v string) *GetHostGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *GetHostGroupRequest) SetRegionId(v string) *GetHostGroupRequest {
	s.RegionId = &v
	return s
}

func (s *GetHostGroupRequest) Validate() error {
	return dara.Validate(s)
}
