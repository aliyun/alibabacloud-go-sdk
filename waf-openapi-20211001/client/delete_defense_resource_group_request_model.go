// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDefenseResourceGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupName(v string) *DeleteDefenseResourceGroupRequest
	GetGroupName() *string
	SetInstanceId(v string) *DeleteDefenseResourceGroupRequest
	GetInstanceId() *string
	SetRegionId(v string) *DeleteDefenseResourceGroupRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DeleteDefenseResourceGroupRequest
	GetResourceManagerResourceGroupId() *string
}

type DeleteDefenseResourceGroupRequest struct {
	// The name of the protected object group that you want to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// group221
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// >  You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to obtain the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_cdnsdf3****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region where the WAF instance resides. Valid values:
	//
	// 	- **cn-hangzhou:*	- the Chinese mainland.
	//
	// 	- **ap-southeast-1:*	- outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
}

func (s DeleteDefenseResourceGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDefenseResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteDefenseResourceGroupRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *DeleteDefenseResourceGroupRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteDefenseResourceGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteDefenseResourceGroupRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DeleteDefenseResourceGroupRequest) SetGroupName(v string) *DeleteDefenseResourceGroupRequest {
	s.GroupName = &v
	return s
}

func (s *DeleteDefenseResourceGroupRequest) SetInstanceId(v string) *DeleteDefenseResourceGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteDefenseResourceGroupRequest) SetRegionId(v string) *DeleteDefenseResourceGroupRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteDefenseResourceGroupRequest) SetResourceManagerResourceGroupId(v string) *DeleteDefenseResourceGroupRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DeleteDefenseResourceGroupRequest) Validate() error {
	return dara.Validate(s)
}
