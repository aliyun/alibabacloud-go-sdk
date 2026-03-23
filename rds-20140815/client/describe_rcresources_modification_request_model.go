// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRCResourcesModificationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConditionss(v []*string) *DescribeRCResourcesModificationRequest
	GetConditionss() []*string
	SetCores(v int32) *DescribeRCResourcesModificationRequest
	GetCores() *int32
	SetDestinationResource(v string) *DescribeRCResourcesModificationRequest
	GetDestinationResource() *string
	SetInstanceId(v string) *DescribeRCResourcesModificationRequest
	GetInstanceId() *string
	SetInstanceType(v string) *DescribeRCResourcesModificationRequest
	GetInstanceType() *string
	SetMemory(v float32) *DescribeRCResourcesModificationRequest
	GetMemory() *float32
	SetOperationType(v string) *DescribeRCResourcesModificationRequest
	GetOperationType() *string
	SetRegionId(v string) *DescribeRCResourcesModificationRequest
	GetRegionId() *string
	SetZoneId(v string) *DescribeRCResourcesModificationRequest
	GetZoneId() *string
}

type DescribeRCResourcesModificationRequest struct {
	Conditionss []*string `json:"Conditionss,omitempty" xml:"Conditionss,omitempty" type:"Repeated"`
	Cores       *int32    `json:"Cores,omitempty" xml:"Cores,omitempty"`
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

func (s DescribeRCResourcesModificationRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCResourcesModificationRequest) GoString() string {
	return s.String()
}

func (s *DescribeRCResourcesModificationRequest) GetConditionss() []*string {
	return s.Conditionss
}

func (s *DescribeRCResourcesModificationRequest) GetCores() *int32 {
	return s.Cores
}

func (s *DescribeRCResourcesModificationRequest) GetDestinationResource() *string {
	return s.DestinationResource
}

func (s *DescribeRCResourcesModificationRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeRCResourcesModificationRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeRCResourcesModificationRequest) GetMemory() *float32 {
	return s.Memory
}

func (s *DescribeRCResourcesModificationRequest) GetOperationType() *string {
	return s.OperationType
}

func (s *DescribeRCResourcesModificationRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeRCResourcesModificationRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeRCResourcesModificationRequest) SetConditionss(v []*string) *DescribeRCResourcesModificationRequest {
	s.Conditionss = v
	return s
}

func (s *DescribeRCResourcesModificationRequest) SetCores(v int32) *DescribeRCResourcesModificationRequest {
	s.Cores = &v
	return s
}

func (s *DescribeRCResourcesModificationRequest) SetDestinationResource(v string) *DescribeRCResourcesModificationRequest {
	s.DestinationResource = &v
	return s
}

func (s *DescribeRCResourcesModificationRequest) SetInstanceId(v string) *DescribeRCResourcesModificationRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeRCResourcesModificationRequest) SetInstanceType(v string) *DescribeRCResourcesModificationRequest {
	s.InstanceType = &v
	return s
}

func (s *DescribeRCResourcesModificationRequest) SetMemory(v float32) *DescribeRCResourcesModificationRequest {
	s.Memory = &v
	return s
}

func (s *DescribeRCResourcesModificationRequest) SetOperationType(v string) *DescribeRCResourcesModificationRequest {
	s.OperationType = &v
	return s
}

func (s *DescribeRCResourcesModificationRequest) SetRegionId(v string) *DescribeRCResourcesModificationRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeRCResourcesModificationRequest) SetZoneId(v string) *DescribeRCResourcesModificationRequest {
	s.ZoneId = &v
	return s
}

func (s *DescribeRCResourcesModificationRequest) Validate() error {
	return dara.Validate(s)
}
