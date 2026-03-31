// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDefenseResourceGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddList(v string) *ModifyDefenseResourceGroupRequest
	GetAddList() *string
	SetDeleteList(v string) *ModifyDefenseResourceGroupRequest
	GetDeleteList() *string
	SetDescription(v string) *ModifyDefenseResourceGroupRequest
	GetDescription() *string
	SetGroupName(v string) *ModifyDefenseResourceGroupRequest
	GetGroupName() *string
	SetInstanceId(v string) *ModifyDefenseResourceGroupRequest
	GetInstanceId() *string
	SetRegionId(v string) *ModifyDefenseResourceGroupRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *ModifyDefenseResourceGroupRequest
	GetResourceManagerResourceGroupId() *string
}

type ModifyDefenseResourceGroupRequest struct {
	// The protected objects that you want to add to the protected object group. Separate the protected objects with commas (,). If you leave this parameter empty, no protected objects are added to the protected object group.
	//
	// example:
	//
	// test1.aliyundoc.com,test2.aliyundoc.com
	AddList *string `json:"AddList,omitempty" xml:"AddList,omitempty"`
	// The protected objects that you want to remove from the protected object group. Separate the protected objects with commas (,). If you leave this parameter empty, no protected objects are removed from the protected object group.
	//
	// example:
	//
	// test14.waf.com,test24.waf.com
	DeleteList *string `json:"DeleteList,omitempty" xml:"DeleteList,omitempty"`
	// The description of the protected object group.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the protected object group whose configurations you want to modify.
	//
	// This parameter is required.
	//
	// example:
	//
	// test01
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

func (s ModifyDefenseResourceGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDefenseResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *ModifyDefenseResourceGroupRequest) GetAddList() *string {
	return s.AddList
}

func (s *ModifyDefenseResourceGroupRequest) GetDeleteList() *string {
	return s.DeleteList
}

func (s *ModifyDefenseResourceGroupRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyDefenseResourceGroupRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *ModifyDefenseResourceGroupRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyDefenseResourceGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyDefenseResourceGroupRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *ModifyDefenseResourceGroupRequest) SetAddList(v string) *ModifyDefenseResourceGroupRequest {
	s.AddList = &v
	return s
}

func (s *ModifyDefenseResourceGroupRequest) SetDeleteList(v string) *ModifyDefenseResourceGroupRequest {
	s.DeleteList = &v
	return s
}

func (s *ModifyDefenseResourceGroupRequest) SetDescription(v string) *ModifyDefenseResourceGroupRequest {
	s.Description = &v
	return s
}

func (s *ModifyDefenseResourceGroupRequest) SetGroupName(v string) *ModifyDefenseResourceGroupRequest {
	s.GroupName = &v
	return s
}

func (s *ModifyDefenseResourceGroupRequest) SetInstanceId(v string) *ModifyDefenseResourceGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyDefenseResourceGroupRequest) SetRegionId(v string) *ModifyDefenseResourceGroupRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDefenseResourceGroupRequest) SetResourceManagerResourceGroupId(v string) *ModifyDefenseResourceGroupRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *ModifyDefenseResourceGroupRequest) Validate() error {
	return dara.Validate(s)
}
