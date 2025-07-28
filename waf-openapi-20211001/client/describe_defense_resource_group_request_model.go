// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDefenseResourceGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupName(v string) *DescribeDefenseResourceGroupRequest
	GetGroupName() *string
	SetInstanceId(v string) *DescribeDefenseResourceGroupRequest
	GetInstanceId() *string
	SetRegionId(v string) *DescribeDefenseResourceGroupRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeDefenseResourceGroupRequest
	GetResourceManagerResourceGroupId() *string
}

type DescribeDefenseResourceGroupRequest struct {
	// The name of the protected object group whose information you want to query.
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
	// The ID of the Alibaba Cloud resource group.
	//
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
}

func (s DescribeDefenseResourceGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDefenseResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *DescribeDefenseResourceGroupRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeDefenseResourceGroupRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeDefenseResourceGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDefenseResourceGroupRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeDefenseResourceGroupRequest) SetGroupName(v string) *DescribeDefenseResourceGroupRequest {
	s.GroupName = &v
	return s
}

func (s *DescribeDefenseResourceGroupRequest) SetInstanceId(v string) *DescribeDefenseResourceGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeDefenseResourceGroupRequest) SetRegionId(v string) *DescribeDefenseResourceGroupRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDefenseResourceGroupRequest) SetResourceManagerResourceGroupId(v string) *DescribeDefenseResourceGroupRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeDefenseResourceGroupRequest) Validate() error {
	return dara.Validate(s)
}
