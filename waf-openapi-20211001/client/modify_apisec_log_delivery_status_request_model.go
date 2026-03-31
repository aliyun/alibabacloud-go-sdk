// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyApisecLogDeliveryStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssertKey(v string) *ModifyApisecLogDeliveryStatusRequest
	GetAssertKey() *string
	SetInstanceId(v string) *ModifyApisecLogDeliveryStatusRequest
	GetInstanceId() *string
	SetRegionId(v string) *ModifyApisecLogDeliveryStatusRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *ModifyApisecLogDeliveryStatusRequest
	GetResourceManagerResourceGroupId() *string
	SetStatus(v bool) *ModifyApisecLogDeliveryStatusRequest
	GetStatus() *bool
}

type ModifyApisecLogDeliveryStatusRequest struct {
	// The type of the log subscription. Valid values:
	//
	// 	- **risk**: risk information.
	//
	// 	- **event**: attack event information.
	//
	// 	- **asset**: asset information.
	//
	// This parameter is required.
	//
	// example:
	//
	// risk
	AssertKey *string `json:"AssertKey,omitempty" xml:"AssertKey,omitempty"`
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// >  You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_v3_public_cn-uqm2z****0a
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region where the WAF instance is deployed. Valid values:
	//
	// 	- **cn-hangzhou**: Chinese mainland.
	//
	// 	- **ap-southeast-1**: outside the Chinese mainland.
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
	// The status of API security log subscription. Valid values:
	//
	// 	- **true**: enabled.
	//
	// 	- **false**: disabled.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	Status *bool `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ModifyApisecLogDeliveryStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyApisecLogDeliveryStatusRequest) GoString() string {
	return s.String()
}

func (s *ModifyApisecLogDeliveryStatusRequest) GetAssertKey() *string {
	return s.AssertKey
}

func (s *ModifyApisecLogDeliveryStatusRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyApisecLogDeliveryStatusRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyApisecLogDeliveryStatusRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *ModifyApisecLogDeliveryStatusRequest) GetStatus() *bool {
	return s.Status
}

func (s *ModifyApisecLogDeliveryStatusRequest) SetAssertKey(v string) *ModifyApisecLogDeliveryStatusRequest {
	s.AssertKey = &v
	return s
}

func (s *ModifyApisecLogDeliveryStatusRequest) SetInstanceId(v string) *ModifyApisecLogDeliveryStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyApisecLogDeliveryStatusRequest) SetRegionId(v string) *ModifyApisecLogDeliveryStatusRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyApisecLogDeliveryStatusRequest) SetResourceManagerResourceGroupId(v string) *ModifyApisecLogDeliveryStatusRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *ModifyApisecLogDeliveryStatusRequest) SetStatus(v bool) *ModifyApisecLogDeliveryStatusRequest {
	s.Status = &v
	return s
}

func (s *ModifyApisecLogDeliveryStatusRequest) Validate() error {
	return dara.Validate(s)
}
