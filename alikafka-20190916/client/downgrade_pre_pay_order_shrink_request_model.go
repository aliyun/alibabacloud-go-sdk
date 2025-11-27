// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDowngradePrePayOrderShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfluentConfigShrink(v string) *DowngradePrePayOrderShrinkRequest
	GetConfluentConfigShrink() *string
	SetDiskSize(v int32) *DowngradePrePayOrderShrinkRequest
	GetDiskSize() *int32
	SetEipMax(v int32) *DowngradePrePayOrderShrinkRequest
	GetEipMax() *int32
	SetEipModel(v bool) *DowngradePrePayOrderShrinkRequest
	GetEipModel() *bool
	SetInstanceId(v string) *DowngradePrePayOrderShrinkRequest
	GetInstanceId() *string
	SetIoMax(v int32) *DowngradePrePayOrderShrinkRequest
	GetIoMax() *int32
	SetIoMaxSpec(v string) *DowngradePrePayOrderShrinkRequest
	GetIoMaxSpec() *string
	SetPaidType(v int32) *DowngradePrePayOrderShrinkRequest
	GetPaidType() *int32
	SetPartitionNum(v int32) *DowngradePrePayOrderShrinkRequest
	GetPartitionNum() *int32
	SetRegionId(v string) *DowngradePrePayOrderShrinkRequest
	GetRegionId() *string
	SetSpecType(v string) *DowngradePrePayOrderShrinkRequest
	GetSpecType() *string
	SetTopicQuota(v int32) *DowngradePrePayOrderShrinkRequest
	GetTopicQuota() *int32
}

type DowngradePrePayOrderShrinkRequest struct {
	ConfluentConfigShrink *string `json:"ConfluentConfig,omitempty" xml:"ConfluentConfig,omitempty"`
	DiskSize              *int32  `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
	EipMax                *int32  `json:"EipMax,omitempty" xml:"EipMax,omitempty"`
	EipModel              *bool   `json:"EipModel,omitempty" xml:"EipModel,omitempty"`
	// This parameter is required.
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	IoMax        *int32  `json:"IoMax,omitempty" xml:"IoMax,omitempty"`
	IoMaxSpec    *string `json:"IoMaxSpec,omitempty" xml:"IoMaxSpec,omitempty"`
	PaidType     *int32  `json:"PaidType,omitempty" xml:"PaidType,omitempty"`
	PartitionNum *int32  `json:"PartitionNum,omitempty" xml:"PartitionNum,omitempty"`
	// This parameter is required.
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SpecType   *string `json:"SpecType,omitempty" xml:"SpecType,omitempty"`
	TopicQuota *int32  `json:"TopicQuota,omitempty" xml:"TopicQuota,omitempty"`
}

func (s DowngradePrePayOrderShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DowngradePrePayOrderShrinkRequest) GoString() string {
	return s.String()
}

func (s *DowngradePrePayOrderShrinkRequest) GetConfluentConfigShrink() *string {
	return s.ConfluentConfigShrink
}

func (s *DowngradePrePayOrderShrinkRequest) GetDiskSize() *int32 {
	return s.DiskSize
}

func (s *DowngradePrePayOrderShrinkRequest) GetEipMax() *int32 {
	return s.EipMax
}

func (s *DowngradePrePayOrderShrinkRequest) GetEipModel() *bool {
	return s.EipModel
}

func (s *DowngradePrePayOrderShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DowngradePrePayOrderShrinkRequest) GetIoMax() *int32 {
	return s.IoMax
}

func (s *DowngradePrePayOrderShrinkRequest) GetIoMaxSpec() *string {
	return s.IoMaxSpec
}

func (s *DowngradePrePayOrderShrinkRequest) GetPaidType() *int32 {
	return s.PaidType
}

func (s *DowngradePrePayOrderShrinkRequest) GetPartitionNum() *int32 {
	return s.PartitionNum
}

func (s *DowngradePrePayOrderShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DowngradePrePayOrderShrinkRequest) GetSpecType() *string {
	return s.SpecType
}

func (s *DowngradePrePayOrderShrinkRequest) GetTopicQuota() *int32 {
	return s.TopicQuota
}

func (s *DowngradePrePayOrderShrinkRequest) SetConfluentConfigShrink(v string) *DowngradePrePayOrderShrinkRequest {
	s.ConfluentConfigShrink = &v
	return s
}

func (s *DowngradePrePayOrderShrinkRequest) SetDiskSize(v int32) *DowngradePrePayOrderShrinkRequest {
	s.DiskSize = &v
	return s
}

func (s *DowngradePrePayOrderShrinkRequest) SetEipMax(v int32) *DowngradePrePayOrderShrinkRequest {
	s.EipMax = &v
	return s
}

func (s *DowngradePrePayOrderShrinkRequest) SetEipModel(v bool) *DowngradePrePayOrderShrinkRequest {
	s.EipModel = &v
	return s
}

func (s *DowngradePrePayOrderShrinkRequest) SetInstanceId(v string) *DowngradePrePayOrderShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *DowngradePrePayOrderShrinkRequest) SetIoMax(v int32) *DowngradePrePayOrderShrinkRequest {
	s.IoMax = &v
	return s
}

func (s *DowngradePrePayOrderShrinkRequest) SetIoMaxSpec(v string) *DowngradePrePayOrderShrinkRequest {
	s.IoMaxSpec = &v
	return s
}

func (s *DowngradePrePayOrderShrinkRequest) SetPaidType(v int32) *DowngradePrePayOrderShrinkRequest {
	s.PaidType = &v
	return s
}

func (s *DowngradePrePayOrderShrinkRequest) SetPartitionNum(v int32) *DowngradePrePayOrderShrinkRequest {
	s.PartitionNum = &v
	return s
}

func (s *DowngradePrePayOrderShrinkRequest) SetRegionId(v string) *DowngradePrePayOrderShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *DowngradePrePayOrderShrinkRequest) SetSpecType(v string) *DowngradePrePayOrderShrinkRequest {
	s.SpecType = &v
	return s
}

func (s *DowngradePrePayOrderShrinkRequest) SetTopicQuota(v int32) *DowngradePrePayOrderShrinkRequest {
	s.TopicQuota = &v
	return s
}

func (s *DowngradePrePayOrderShrinkRequest) Validate() error {
	return dara.Validate(s)
}
