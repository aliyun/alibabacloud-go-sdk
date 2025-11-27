// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDowngradePostPayOrderShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDiskSize(v int32) *DowngradePostPayOrderShrinkRequest
	GetDiskSize() *int32
	SetEipMax(v int32) *DowngradePostPayOrderShrinkRequest
	GetEipMax() *int32
	SetEipModel(v bool) *DowngradePostPayOrderShrinkRequest
	GetEipModel() *bool
	SetInstanceId(v string) *DowngradePostPayOrderShrinkRequest
	GetInstanceId() *string
	SetIoMax(v int32) *DowngradePostPayOrderShrinkRequest
	GetIoMax() *int32
	SetIoMaxSpec(v string) *DowngradePostPayOrderShrinkRequest
	GetIoMaxSpec() *string
	SetPartitionNum(v int32) *DowngradePostPayOrderShrinkRequest
	GetPartitionNum() *int32
	SetRegionId(v string) *DowngradePostPayOrderShrinkRequest
	GetRegionId() *string
	SetServerlessConfigShrink(v string) *DowngradePostPayOrderShrinkRequest
	GetServerlessConfigShrink() *string
	SetSpecType(v string) *DowngradePostPayOrderShrinkRequest
	GetSpecType() *string
	SetTopicQuota(v int32) *DowngradePostPayOrderShrinkRequest
	GetTopicQuota() *int32
}

type DowngradePostPayOrderShrinkRequest struct {
	DiskSize *int32 `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
	EipMax   *int32 `json:"EipMax,omitempty" xml:"EipMax,omitempty"`
	EipModel *bool  `json:"EipModel,omitempty" xml:"EipModel,omitempty"`
	// This parameter is required.
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	IoMax        *int32  `json:"IoMax,omitempty" xml:"IoMax,omitempty"`
	IoMaxSpec    *string `json:"IoMaxSpec,omitempty" xml:"IoMaxSpec,omitempty"`
	PartitionNum *int32  `json:"PartitionNum,omitempty" xml:"PartitionNum,omitempty"`
	// This parameter is required.
	RegionId               *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ServerlessConfigShrink *string `json:"ServerlessConfig,omitempty" xml:"ServerlessConfig,omitempty"`
	SpecType               *string `json:"SpecType,omitempty" xml:"SpecType,omitempty"`
	TopicQuota             *int32  `json:"TopicQuota,omitempty" xml:"TopicQuota,omitempty"`
}

func (s DowngradePostPayOrderShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DowngradePostPayOrderShrinkRequest) GoString() string {
	return s.String()
}

func (s *DowngradePostPayOrderShrinkRequest) GetDiskSize() *int32 {
	return s.DiskSize
}

func (s *DowngradePostPayOrderShrinkRequest) GetEipMax() *int32 {
	return s.EipMax
}

func (s *DowngradePostPayOrderShrinkRequest) GetEipModel() *bool {
	return s.EipModel
}

func (s *DowngradePostPayOrderShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DowngradePostPayOrderShrinkRequest) GetIoMax() *int32 {
	return s.IoMax
}

func (s *DowngradePostPayOrderShrinkRequest) GetIoMaxSpec() *string {
	return s.IoMaxSpec
}

func (s *DowngradePostPayOrderShrinkRequest) GetPartitionNum() *int32 {
	return s.PartitionNum
}

func (s *DowngradePostPayOrderShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DowngradePostPayOrderShrinkRequest) GetServerlessConfigShrink() *string {
	return s.ServerlessConfigShrink
}

func (s *DowngradePostPayOrderShrinkRequest) GetSpecType() *string {
	return s.SpecType
}

func (s *DowngradePostPayOrderShrinkRequest) GetTopicQuota() *int32 {
	return s.TopicQuota
}

func (s *DowngradePostPayOrderShrinkRequest) SetDiskSize(v int32) *DowngradePostPayOrderShrinkRequest {
	s.DiskSize = &v
	return s
}

func (s *DowngradePostPayOrderShrinkRequest) SetEipMax(v int32) *DowngradePostPayOrderShrinkRequest {
	s.EipMax = &v
	return s
}

func (s *DowngradePostPayOrderShrinkRequest) SetEipModel(v bool) *DowngradePostPayOrderShrinkRequest {
	s.EipModel = &v
	return s
}

func (s *DowngradePostPayOrderShrinkRequest) SetInstanceId(v string) *DowngradePostPayOrderShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *DowngradePostPayOrderShrinkRequest) SetIoMax(v int32) *DowngradePostPayOrderShrinkRequest {
	s.IoMax = &v
	return s
}

func (s *DowngradePostPayOrderShrinkRequest) SetIoMaxSpec(v string) *DowngradePostPayOrderShrinkRequest {
	s.IoMaxSpec = &v
	return s
}

func (s *DowngradePostPayOrderShrinkRequest) SetPartitionNum(v int32) *DowngradePostPayOrderShrinkRequest {
	s.PartitionNum = &v
	return s
}

func (s *DowngradePostPayOrderShrinkRequest) SetRegionId(v string) *DowngradePostPayOrderShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *DowngradePostPayOrderShrinkRequest) SetServerlessConfigShrink(v string) *DowngradePostPayOrderShrinkRequest {
	s.ServerlessConfigShrink = &v
	return s
}

func (s *DowngradePostPayOrderShrinkRequest) SetSpecType(v string) *DowngradePostPayOrderShrinkRequest {
	s.SpecType = &v
	return s
}

func (s *DowngradePostPayOrderShrinkRequest) SetTopicQuota(v int32) *DowngradePostPayOrderShrinkRequest {
	s.TopicQuota = &v
	return s
}

func (s *DowngradePostPayOrderShrinkRequest) Validate() error {
	return dara.Validate(s)
}
