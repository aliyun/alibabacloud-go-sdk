// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRCResourcesModificationShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConditionssShrink(v string) *DescribeRCResourcesModificationShrinkRequest
	GetConditionssShrink() *string
	SetCores(v int32) *DescribeRCResourcesModificationShrinkRequest
	GetCores() *int32
	SetDestinationResource(v string) *DescribeRCResourcesModificationShrinkRequest
	GetDestinationResource() *string
	SetInstanceId(v string) *DescribeRCResourcesModificationShrinkRequest
	GetInstanceId() *string
	SetInstanceType(v string) *DescribeRCResourcesModificationShrinkRequest
	GetInstanceType() *string
	SetMemory(v float32) *DescribeRCResourcesModificationShrinkRequest
	GetMemory() *float32
	SetOperationType(v string) *DescribeRCResourcesModificationShrinkRequest
	GetOperationType() *string
	SetRegionId(v string) *DescribeRCResourcesModificationShrinkRequest
	GetRegionId() *string
	SetZoneId(v string) *DescribeRCResourcesModificationShrinkRequest
	GetZoneId() *string
}

type DescribeRCResourcesModificationShrinkRequest struct {
	ConditionssShrink *string `json:"Conditionss,omitempty" xml:"Conditionss,omitempty"`
	Cores             *int32  `json:"Cores,omitempty" xml:"Cores,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// InstanceType
	DestinationResource *string `json:"DestinationResource,omitempty" xml:"DestinationResource,omitempty"`
	// This parameter is required.
	InstanceId   *string  `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceType *string  `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	Memory       *float32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// example:
	//
	// Upgrade
	OperationType *string `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	// This parameter is required.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ZoneId   *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeRCResourcesModificationShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCResourcesModificationShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeRCResourcesModificationShrinkRequest) GetConditionssShrink() *string {
	return s.ConditionssShrink
}

func (s *DescribeRCResourcesModificationShrinkRequest) GetCores() *int32 {
	return s.Cores
}

func (s *DescribeRCResourcesModificationShrinkRequest) GetDestinationResource() *string {
	return s.DestinationResource
}

func (s *DescribeRCResourcesModificationShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeRCResourcesModificationShrinkRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeRCResourcesModificationShrinkRequest) GetMemory() *float32 {
	return s.Memory
}

func (s *DescribeRCResourcesModificationShrinkRequest) GetOperationType() *string {
	return s.OperationType
}

func (s *DescribeRCResourcesModificationShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeRCResourcesModificationShrinkRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeRCResourcesModificationShrinkRequest) SetConditionssShrink(v string) *DescribeRCResourcesModificationShrinkRequest {
	s.ConditionssShrink = &v
	return s
}

func (s *DescribeRCResourcesModificationShrinkRequest) SetCores(v int32) *DescribeRCResourcesModificationShrinkRequest {
	s.Cores = &v
	return s
}

func (s *DescribeRCResourcesModificationShrinkRequest) SetDestinationResource(v string) *DescribeRCResourcesModificationShrinkRequest {
	s.DestinationResource = &v
	return s
}

func (s *DescribeRCResourcesModificationShrinkRequest) SetInstanceId(v string) *DescribeRCResourcesModificationShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeRCResourcesModificationShrinkRequest) SetInstanceType(v string) *DescribeRCResourcesModificationShrinkRequest {
	s.InstanceType = &v
	return s
}

func (s *DescribeRCResourcesModificationShrinkRequest) SetMemory(v float32) *DescribeRCResourcesModificationShrinkRequest {
	s.Memory = &v
	return s
}

func (s *DescribeRCResourcesModificationShrinkRequest) SetOperationType(v string) *DescribeRCResourcesModificationShrinkRequest {
	s.OperationType = &v
	return s
}

func (s *DescribeRCResourcesModificationShrinkRequest) SetRegionId(v string) *DescribeRCResourcesModificationShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeRCResourcesModificationShrinkRequest) SetZoneId(v string) *DescribeRCResourcesModificationShrinkRequest {
	s.ZoneId = &v
	return s
}

func (s *DescribeRCResourcesModificationShrinkRequest) Validate() error {
	return dara.Validate(s)
}
