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
	// example:
	//
	// export-kafka
	DeliveryName *string `json:"DeliveryName,omitempty" xml:"DeliveryName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// waf_v2_public_cn-uqm35*****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test.waf.com-waf
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
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
