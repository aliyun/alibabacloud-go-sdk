// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDefenseResourceGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddList(v string) *CreateDefenseResourceGroupRequest
	GetAddList() *string
	SetDescription(v string) *CreateDefenseResourceGroupRequest
	GetDescription() *string
	SetGroupName(v string) *CreateDefenseResourceGroupRequest
	GetGroupName() *string
	SetInstanceId(v string) *CreateDefenseResourceGroupRequest
	GetInstanceId() *string
	SetRegionId(v string) *CreateDefenseResourceGroupRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *CreateDefenseResourceGroupRequest
	GetResourceManagerResourceGroupId() *string
}

type CreateDefenseResourceGroupRequest struct {
	// The protected objects to add to the protected object group. You can add multiple protected objects. Separate them with commas (,).
	//
	// example:
	//
	// test1.aliyundoc.com,test2.aliyundoc.com
	AddList *string `json:"AddList,omitempty" xml:"AddList,omitempty"`
	// The description of the protected object group. The description can be up to 512 characters long.
	//
	// example:
	//
	// test_domain
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the protected object group. The name must be 1 to 255 characters long and can contain Chinese characters, letters, digits, underscores (_), periods (.), and hyphens (-).
	//
	// This parameter is required.
	//
	// example:
	//
	// group221
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The ID of the WAF instance.
	//
	// > Call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to view the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_cdnsdf3****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region where the WAF instance resides. Valid values:
	//
	// - **cn-hangzhou**: the Chinese mainland.
	//
	// - **ap-southeast-1**: outside the Chinese mainland.
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

func (s CreateDefenseResourceGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDefenseResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateDefenseResourceGroupRequest) GetAddList() *string {
	return s.AddList
}

func (s *CreateDefenseResourceGroupRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateDefenseResourceGroupRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *CreateDefenseResourceGroupRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateDefenseResourceGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateDefenseResourceGroupRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *CreateDefenseResourceGroupRequest) SetAddList(v string) *CreateDefenseResourceGroupRequest {
	s.AddList = &v
	return s
}

func (s *CreateDefenseResourceGroupRequest) SetDescription(v string) *CreateDefenseResourceGroupRequest {
	s.Description = &v
	return s
}

func (s *CreateDefenseResourceGroupRequest) SetGroupName(v string) *CreateDefenseResourceGroupRequest {
	s.GroupName = &v
	return s
}

func (s *CreateDefenseResourceGroupRequest) SetInstanceId(v string) *CreateDefenseResourceGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateDefenseResourceGroupRequest) SetRegionId(v string) *CreateDefenseResourceGroupRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDefenseResourceGroupRequest) SetResourceManagerResourceGroupId(v string) *CreateDefenseResourceGroupRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *CreateDefenseResourceGroupRequest) Validate() error {
	return dara.Validate(s)
}
