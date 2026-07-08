// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyResourceLogDeliveryStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeliveryName(v string) *ModifyResourceLogDeliveryStatusRequest
	GetDeliveryName() *string
	SetInstanceId(v string) *ModifyResourceLogDeliveryStatusRequest
	GetInstanceId() *string
	SetRegionId(v string) *ModifyResourceLogDeliveryStatusRequest
	GetRegionId() *string
	SetResource(v string) *ModifyResourceLogDeliveryStatusRequest
	GetResource() *string
	SetResourceManagerResourceGroupId(v string) *ModifyResourceLogDeliveryStatusRequest
	GetResourceManagerResourceGroupId() *string
	SetStatus(v bool) *ModifyResourceLogDeliveryStatusRequest
	GetStatus() *bool
}

type ModifyResourceLogDeliveryStatusRequest struct {
	// The name of the log delivery configuration.
	//
	// > This parameter is required when you enable log delivery by setting **Status*	- to **true**.
	//
	// example:
	//
	// export-kafka
	DeliveryName *string `json:"DeliveryName,omitempty" xml:"DeliveryName,omitempty"`
	// The ID of the WAF instance.
	//
	// > Call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_v2_public_cn-uqm35*****
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
	// The protected object for which you want to modify the log delivery status.
	//
	// This parameter is required.
	//
	// example:
	//
	// test.waf.com-waf
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	//
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// Indicates whether to enable log delivery for the protected object. Valid values:
	//
	// - **true**: enables log delivery.
	//
	// - **false**: disables log delivery.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	Status *bool `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ModifyResourceLogDeliveryStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyResourceLogDeliveryStatusRequest) GoString() string {
	return s.String()
}

func (s *ModifyResourceLogDeliveryStatusRequest) GetDeliveryName() *string {
	return s.DeliveryName
}

func (s *ModifyResourceLogDeliveryStatusRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyResourceLogDeliveryStatusRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyResourceLogDeliveryStatusRequest) GetResource() *string {
	return s.Resource
}

func (s *ModifyResourceLogDeliveryStatusRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *ModifyResourceLogDeliveryStatusRequest) GetStatus() *bool {
	return s.Status
}

func (s *ModifyResourceLogDeliveryStatusRequest) SetDeliveryName(v string) *ModifyResourceLogDeliveryStatusRequest {
	s.DeliveryName = &v
	return s
}

func (s *ModifyResourceLogDeliveryStatusRequest) SetInstanceId(v string) *ModifyResourceLogDeliveryStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyResourceLogDeliveryStatusRequest) SetRegionId(v string) *ModifyResourceLogDeliveryStatusRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyResourceLogDeliveryStatusRequest) SetResource(v string) *ModifyResourceLogDeliveryStatusRequest {
	s.Resource = &v
	return s
}

func (s *ModifyResourceLogDeliveryStatusRequest) SetResourceManagerResourceGroupId(v string) *ModifyResourceLogDeliveryStatusRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *ModifyResourceLogDeliveryStatusRequest) SetStatus(v bool) *ModifyResourceLogDeliveryStatusRequest {
	s.Status = &v
	return s
}

func (s *ModifyResourceLogDeliveryStatusRequest) Validate() error {
	return dara.Validate(s)
}
