// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLogDeliveryConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeliveryName(v string) *DeleteLogDeliveryConfigRequest
	GetDeliveryName() *string
	SetInstanceId(v string) *DeleteLogDeliveryConfigRequest
	GetInstanceId() *string
	SetRegionId(v string) *DeleteLogDeliveryConfigRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DeleteLogDeliveryConfigRequest
	GetResourceManagerResourceGroupId() *string
}

type DeleteLogDeliveryConfigRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// test1
	DeliveryName *string `json:"DeliveryName,omitempty" xml:"DeliveryName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// waf-cn-n6w***x52m
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
}

func (s DeleteLogDeliveryConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteLogDeliveryConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteLogDeliveryConfigRequest) GetDeliveryName() *string {
	return s.DeliveryName
}

func (s *DeleteLogDeliveryConfigRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteLogDeliveryConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteLogDeliveryConfigRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DeleteLogDeliveryConfigRequest) SetDeliveryName(v string) *DeleteLogDeliveryConfigRequest {
	s.DeliveryName = &v
	return s
}

func (s *DeleteLogDeliveryConfigRequest) SetInstanceId(v string) *DeleteLogDeliveryConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteLogDeliveryConfigRequest) SetRegionId(v string) *DeleteLogDeliveryConfigRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteLogDeliveryConfigRequest) SetResourceManagerResourceGroupId(v string) *DeleteLogDeliveryConfigRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DeleteLogDeliveryConfigRequest) Validate() error {
	return dara.Validate(s)
}
