// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyHybridCloudServerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContinents(v string) *ModifyHybridCloudServerRequest
	GetContinents() *string
	SetCustomName(v string) *ModifyHybridCloudServerRequest
	GetCustomName() *string
	SetInstanceId(v string) *ModifyHybridCloudServerRequest
	GetInstanceId() *string
	SetMid(v string) *ModifyHybridCloudServerRequest
	GetMid() *string
	SetOperator(v string) *ModifyHybridCloudServerRequest
	GetOperator() *string
	SetRegionCode(v string) *ModifyHybridCloudServerRequest
	GetRegionCode() *string
	SetRegionId(v string) *ModifyHybridCloudServerRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *ModifyHybridCloudServerRequest
	GetResourceManagerResourceGroupId() *string
}

type ModifyHybridCloudServerRequest struct {
	// The continent.
	//
	// This parameter is required.
	//
	// example:
	//
	// asiapacific
	Continents *string `json:"Continents,omitempty" xml:"Continents,omitempty"`
	// The name of the node.
	//
	// This parameter is required.
	//
	// example:
	//
	// demo
	CustomName *string `json:"CustomName,omitempty" xml:"CustomName,omitempty"`
	// The ID of the WAF instance.
	//
	// >  You can call the [DescribeInstanceInfo](https://help.aliyun.com/document_detail/140857.html) operation to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_cdnsdf3****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the node.
	//
	// This parameter is required.
	//
	// example:
	//
	// b1bf3f544f30c1de0b72d91290**ccb
	Mid *string `json:"Mid,omitempty" xml:"Mid,omitempty"`
	// The cloud service provider.
	//
	// This parameter is required.
	//
	// example:
	//
	// aliyun
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// The city.
	//
	// This parameter is required.
	//
	// example:
	//
	// beijing
	RegionCode *string `json:"RegionCode,omitempty" xml:"RegionCode,omitempty"`
	// The region of the WAF instance. Valid values:
	//
	// 	- **cn-hangzhou**: Chinese mainland.
	//
	// 	- **ap-southeast-1**: Outside the Chinese mainland.
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

func (s ModifyHybridCloudServerRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyHybridCloudServerRequest) GoString() string {
	return s.String()
}

func (s *ModifyHybridCloudServerRequest) GetContinents() *string {
	return s.Continents
}

func (s *ModifyHybridCloudServerRequest) GetCustomName() *string {
	return s.CustomName
}

func (s *ModifyHybridCloudServerRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyHybridCloudServerRequest) GetMid() *string {
	return s.Mid
}

func (s *ModifyHybridCloudServerRequest) GetOperator() *string {
	return s.Operator
}

func (s *ModifyHybridCloudServerRequest) GetRegionCode() *string {
	return s.RegionCode
}

func (s *ModifyHybridCloudServerRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyHybridCloudServerRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *ModifyHybridCloudServerRequest) SetContinents(v string) *ModifyHybridCloudServerRequest {
	s.Continents = &v
	return s
}

func (s *ModifyHybridCloudServerRequest) SetCustomName(v string) *ModifyHybridCloudServerRequest {
	s.CustomName = &v
	return s
}

func (s *ModifyHybridCloudServerRequest) SetInstanceId(v string) *ModifyHybridCloudServerRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyHybridCloudServerRequest) SetMid(v string) *ModifyHybridCloudServerRequest {
	s.Mid = &v
	return s
}

func (s *ModifyHybridCloudServerRequest) SetOperator(v string) *ModifyHybridCloudServerRequest {
	s.Operator = &v
	return s
}

func (s *ModifyHybridCloudServerRequest) SetRegionCode(v string) *ModifyHybridCloudServerRequest {
	s.RegionCode = &v
	return s
}

func (s *ModifyHybridCloudServerRequest) SetRegionId(v string) *ModifyHybridCloudServerRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyHybridCloudServerRequest) SetResourceManagerResourceGroupId(v string) *ModifyHybridCloudServerRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *ModifyHybridCloudServerRequest) Validate() error {
	return dara.Validate(s)
}
