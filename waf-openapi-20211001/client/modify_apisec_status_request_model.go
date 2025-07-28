// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyApisecStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApisecStatus(v int32) *ModifyApisecStatusRequest
	GetApisecStatus() *int32
	SetInstanceId(v string) *ModifyApisecStatusRequest
	GetInstanceId() *string
	SetRegionId(v string) *ModifyApisecStatusRequest
	GetRegionId() *string
	SetResourceGroups(v string) *ModifyApisecStatusRequest
	GetResourceGroups() *string
	SetResourceManagerResourceGroupId(v string) *ModifyApisecStatusRequest
	GetResourceManagerResourceGroupId() *string
	SetResources(v string) *ModifyApisecStatusRequest
	GetResources() *string
}

type ModifyApisecStatusRequest struct {
	// The status of the API security module. Valid values:
	//
	// 	- **1**: enabled
	//
	// 	- **0**: disabled
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ApisecStatus *int32 `json:"ApisecStatus,omitempty" xml:"ApisecStatus,omitempty"`
	// The ID of the WAF instance.
	//
	// >  You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf-cn-tl32ast****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region in which the Web Application Firewall (WAF) instance is deployed. Valid values:
	//
	// 	- **cn-hangzhou**: Chinese mainland
	//
	// 	- **ap-southeast-1**: outside the Chinese mainland
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the protected object group to which the protected object belongs.
	//
	// example:
	//
	// group
	ResourceGroups *string `json:"ResourceGroups,omitempty" xml:"ResourceGroups,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	//
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The name of the protected object.
	//
	// example:
	//
	// alb-wewbb23dfset***
	Resources *string `json:"Resources,omitempty" xml:"Resources,omitempty"`
}

func (s ModifyApisecStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyApisecStatusRequest) GoString() string {
	return s.String()
}

func (s *ModifyApisecStatusRequest) GetApisecStatus() *int32 {
	return s.ApisecStatus
}

func (s *ModifyApisecStatusRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyApisecStatusRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyApisecStatusRequest) GetResourceGroups() *string {
	return s.ResourceGroups
}

func (s *ModifyApisecStatusRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *ModifyApisecStatusRequest) GetResources() *string {
	return s.Resources
}

func (s *ModifyApisecStatusRequest) SetApisecStatus(v int32) *ModifyApisecStatusRequest {
	s.ApisecStatus = &v
	return s
}

func (s *ModifyApisecStatusRequest) SetInstanceId(v string) *ModifyApisecStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyApisecStatusRequest) SetRegionId(v string) *ModifyApisecStatusRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyApisecStatusRequest) SetResourceGroups(v string) *ModifyApisecStatusRequest {
	s.ResourceGroups = &v
	return s
}

func (s *ModifyApisecStatusRequest) SetResourceManagerResourceGroupId(v string) *ModifyApisecStatusRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *ModifyApisecStatusRequest) SetResources(v string) *ModifyApisecStatusRequest {
	s.Resources = &v
	return s
}

func (s *ModifyApisecStatusRequest) Validate() error {
	return dara.Validate(s)
}
