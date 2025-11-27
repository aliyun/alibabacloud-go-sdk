// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDowngradePostPayOrderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDiskSize(v int32) *DowngradePostPayOrderRequest
	GetDiskSize() *int32
	SetEipMax(v int32) *DowngradePostPayOrderRequest
	GetEipMax() *int32
	SetEipModel(v bool) *DowngradePostPayOrderRequest
	GetEipModel() *bool
	SetInstanceId(v string) *DowngradePostPayOrderRequest
	GetInstanceId() *string
	SetIoMax(v int32) *DowngradePostPayOrderRequest
	GetIoMax() *int32
	SetIoMaxSpec(v string) *DowngradePostPayOrderRequest
	GetIoMaxSpec() *string
	SetPartitionNum(v int32) *DowngradePostPayOrderRequest
	GetPartitionNum() *int32
	SetRegionId(v string) *DowngradePostPayOrderRequest
	GetRegionId() *string
	SetServerlessConfig(v *DowngradePostPayOrderRequestServerlessConfig) *DowngradePostPayOrderRequest
	GetServerlessConfig() *DowngradePostPayOrderRequestServerlessConfig
	SetSpecType(v string) *DowngradePostPayOrderRequest
	GetSpecType() *string
	SetTopicQuota(v int32) *DowngradePostPayOrderRequest
	GetTopicQuota() *int32
}

type DowngradePostPayOrderRequest struct {
	DiskSize *int32 `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
	EipMax   *int32 `json:"EipMax,omitempty" xml:"EipMax,omitempty"`
	EipModel *bool  `json:"EipModel,omitempty" xml:"EipModel,omitempty"`
	// This parameter is required.
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	IoMax        *int32  `json:"IoMax,omitempty" xml:"IoMax,omitempty"`
	IoMaxSpec    *string `json:"IoMaxSpec,omitempty" xml:"IoMaxSpec,omitempty"`
	PartitionNum *int32  `json:"PartitionNum,omitempty" xml:"PartitionNum,omitempty"`
	// This parameter is required.
	RegionId         *string                                       `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ServerlessConfig *DowngradePostPayOrderRequestServerlessConfig `json:"ServerlessConfig,omitempty" xml:"ServerlessConfig,omitempty" type:"Struct"`
	SpecType         *string                                       `json:"SpecType,omitempty" xml:"SpecType,omitempty"`
	TopicQuota       *int32                                        `json:"TopicQuota,omitempty" xml:"TopicQuota,omitempty"`
}

func (s DowngradePostPayOrderRequest) String() string {
	return dara.Prettify(s)
}

func (s DowngradePostPayOrderRequest) GoString() string {
	return s.String()
}

func (s *DowngradePostPayOrderRequest) GetDiskSize() *int32 {
	return s.DiskSize
}

func (s *DowngradePostPayOrderRequest) GetEipMax() *int32 {
	return s.EipMax
}

func (s *DowngradePostPayOrderRequest) GetEipModel() *bool {
	return s.EipModel
}

func (s *DowngradePostPayOrderRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DowngradePostPayOrderRequest) GetIoMax() *int32 {
	return s.IoMax
}

func (s *DowngradePostPayOrderRequest) GetIoMaxSpec() *string {
	return s.IoMaxSpec
}

func (s *DowngradePostPayOrderRequest) GetPartitionNum() *int32 {
	return s.PartitionNum
}

func (s *DowngradePostPayOrderRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DowngradePostPayOrderRequest) GetServerlessConfig() *DowngradePostPayOrderRequestServerlessConfig {
	return s.ServerlessConfig
}

func (s *DowngradePostPayOrderRequest) GetSpecType() *string {
	return s.SpecType
}

func (s *DowngradePostPayOrderRequest) GetTopicQuota() *int32 {
	return s.TopicQuota
}

func (s *DowngradePostPayOrderRequest) SetDiskSize(v int32) *DowngradePostPayOrderRequest {
	s.DiskSize = &v
	return s
}

func (s *DowngradePostPayOrderRequest) SetEipMax(v int32) *DowngradePostPayOrderRequest {
	s.EipMax = &v
	return s
}

func (s *DowngradePostPayOrderRequest) SetEipModel(v bool) *DowngradePostPayOrderRequest {
	s.EipModel = &v
	return s
}

func (s *DowngradePostPayOrderRequest) SetInstanceId(v string) *DowngradePostPayOrderRequest {
	s.InstanceId = &v
	return s
}

func (s *DowngradePostPayOrderRequest) SetIoMax(v int32) *DowngradePostPayOrderRequest {
	s.IoMax = &v
	return s
}

func (s *DowngradePostPayOrderRequest) SetIoMaxSpec(v string) *DowngradePostPayOrderRequest {
	s.IoMaxSpec = &v
	return s
}

func (s *DowngradePostPayOrderRequest) SetPartitionNum(v int32) *DowngradePostPayOrderRequest {
	s.PartitionNum = &v
	return s
}

func (s *DowngradePostPayOrderRequest) SetRegionId(v string) *DowngradePostPayOrderRequest {
	s.RegionId = &v
	return s
}

func (s *DowngradePostPayOrderRequest) SetServerlessConfig(v *DowngradePostPayOrderRequestServerlessConfig) *DowngradePostPayOrderRequest {
	s.ServerlessConfig = v
	return s
}

func (s *DowngradePostPayOrderRequest) SetSpecType(v string) *DowngradePostPayOrderRequest {
	s.SpecType = &v
	return s
}

func (s *DowngradePostPayOrderRequest) SetTopicQuota(v int32) *DowngradePostPayOrderRequest {
	s.TopicQuota = &v
	return s
}

func (s *DowngradePostPayOrderRequest) Validate() error {
	if s.ServerlessConfig != nil {
		if err := s.ServerlessConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DowngradePostPayOrderRequestServerlessConfig struct {
	ReservedPublishCapacity   *int64 `json:"ReservedPublishCapacity,omitempty" xml:"ReservedPublishCapacity,omitempty"`
	ReservedSubscribeCapacity *int64 `json:"ReservedSubscribeCapacity,omitempty" xml:"ReservedSubscribeCapacity,omitempty"`
}

func (s DowngradePostPayOrderRequestServerlessConfig) String() string {
	return dara.Prettify(s)
}

func (s DowngradePostPayOrderRequestServerlessConfig) GoString() string {
	return s.String()
}

func (s *DowngradePostPayOrderRequestServerlessConfig) GetReservedPublishCapacity() *int64 {
	return s.ReservedPublishCapacity
}

func (s *DowngradePostPayOrderRequestServerlessConfig) GetReservedSubscribeCapacity() *int64 {
	return s.ReservedSubscribeCapacity
}

func (s *DowngradePostPayOrderRequestServerlessConfig) SetReservedPublishCapacity(v int64) *DowngradePostPayOrderRequestServerlessConfig {
	s.ReservedPublishCapacity = &v
	return s
}

func (s *DowngradePostPayOrderRequestServerlessConfig) SetReservedSubscribeCapacity(v int64) *DowngradePostPayOrderRequestServerlessConfig {
	s.ReservedSubscribeCapacity = &v
	return s
}

func (s *DowngradePostPayOrderRequestServerlessConfig) Validate() error {
	return dara.Validate(s)
}
