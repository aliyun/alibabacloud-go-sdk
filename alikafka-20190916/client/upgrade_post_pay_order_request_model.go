// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradePostPayOrderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDiskSize(v int32) *UpgradePostPayOrderRequest
	GetDiskSize() *int32
	SetEipMax(v int32) *UpgradePostPayOrderRequest
	GetEipMax() *int32
	SetEipModel(v bool) *UpgradePostPayOrderRequest
	GetEipModel() *bool
	SetInstanceId(v string) *UpgradePostPayOrderRequest
	GetInstanceId() *string
	SetIoMax(v int32) *UpgradePostPayOrderRequest
	GetIoMax() *int32
	SetIoMaxSpec(v string) *UpgradePostPayOrderRequest
	GetIoMaxSpec() *string
	SetPartitionNum(v int32) *UpgradePostPayOrderRequest
	GetPartitionNum() *int32
	SetRegionId(v string) *UpgradePostPayOrderRequest
	GetRegionId() *string
	SetServerlessConfig(v *UpgradePostPayOrderRequestServerlessConfig) *UpgradePostPayOrderRequest
	GetServerlessConfig() *UpgradePostPayOrderRequestServerlessConfig
	SetSpecType(v string) *UpgradePostPayOrderRequest
	GetSpecType() *string
	SetTopicQuota(v int32) *UpgradePostPayOrderRequest
	GetTopicQuota() *int32
}

type UpgradePostPayOrderRequest struct {
	// The disk capacity. Unit: GB.
	//
	// - The disk capacity that you specify must be greater than or equal to the current disk capacity of the instance.
	//
	// - For the value range, see [Billing](https://help.aliyun.com/document_detail/84737.html).
	//
	// > If the instance is a serverless instance, you do not need to specify this parameter. This parameter is required for pay-as-you-go instances.
	//
	// example:
	//
	// 500
	DiskSize *int32 `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
	// The Internet traffic.
	//
	// - The Internet traffic that you specify must be greater than or equal to the current Internet traffic of the instance.
	//
	// - For the value range, see [Billing](https://help.aliyun.com/document_detail/84737.html).
	//
	// > - If EipModel is set to true, the value of EipMax must be greater than 0.
	//
	// > - If EipModel is set to false, the value of EipMax must be 0.
	//
	// > - If the instance is a serverless instance, you do not need to specify this parameter.
	//
	// example:
	//
	// 0
	EipMax *int32 `json:"EipMax,omitempty" xml:"EipMax,omitempty"`
	// Specifies whether the instance requires Internet access. Valid values:
	//
	// - true: Internet access is required.
	//
	// - false: Internet access is not required.
	//
	// > This parameter is optional for pay-as-you-go instances. This parameter is required for serverless instances.
	//
	// example:
	//
	// false
	EipModel *bool `json:"EipModel,omitempty" xml:"EipModel,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// alikafka_post-cn-mp919o4v****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The peak traffic (not recommended).
	//
	// - The peak traffic that you specify must be greater than or equal to the current peak traffic of the instance.
	//
	// - You must specify either the peak traffic or the traffic specification. If you specify both, the traffic specification takes precedence. Specify only the traffic specification.
	//
	// - For the value range, see [Billing](https://help.aliyun.com/document_detail/84737.html).
	//
	// > If the instance is a serverless instance, you do not need to specify this parameter.
	//
	// example:
	//
	// 60
	IoMax *int32 `json:"IoMax,omitempty" xml:"IoMax,omitempty"`
	// The traffic specification (recommended).
	//
	// - The traffic specification that you specify must be greater than or equal to the current traffic specification of the instance.
	//
	// - You must specify either the peak traffic or the traffic specification. If you specify both, the traffic specification takes precedence. Specify only the traffic specification.
	//
	// - For the value range, see [Billing](https://help.aliyun.com/document_detail/84737.html).
	//
	// > If the instance is a serverless instance, you do not need to specify this parameter. This parameter is required for pay-as-you-go instances.
	//
	// example:
	//
	// alikafka.hw.6xlarge
	IoMaxSpec *string `json:"IoMaxSpec,omitempty" xml:"IoMaxSpec,omitempty"`
	// The number of partitions (recommended).
	//
	// 	- You must specify either the number of partitions or the topic specification. Specify only the number of partitions.
	//
	// 	- If you specify both the number of partitions and the topic specification, the system validates whether the number of partitions and the topic specification are equivalent based on the legacy topic sales model. If they are not equivalent, an error is returned. If they are equivalent, the purchase is made based on the number of partitions.
	//
	// 	- For the value range, see [Billing](https://help.aliyun.com/document_detail/84737.html).
	//
	// > If the instance is a serverless instance, you do not need to specify this parameter. This parameter is required for pay-as-you-go instances.
	//
	// example:
	//
	// 80
	PartitionNum *int32 `json:"PartitionNum,omitempty" xml:"PartitionNum,omitempty"`
	// The region ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The settings of the serverless instance. This parameter is required when you change the specifications of a serverless instance.
	ServerlessConfig *UpgradePostPayOrderRequestServerlessConfig `json:"ServerlessConfig,omitempty" xml:"ServerlessConfig,omitempty" type:"Struct"`
	// The specification type.
	//
	// If the PaidType of the instance is 1 (pay-as-you-go), valid values:
	//
	// - normal: Standard Edition (shared throughput)
	//
	// - professional: Professional Edition (shared throughput)
	//
	// - professionalForHighRead: Professional Edition (shared read throughput)
	//
	// If the PaidType of the instance is 3 (reserved specification pay-as-you-go + serverless elastic scaling pay-as-you-go), valid values:
	//
	// - normal: Serverless Standard Edition
	//
	// For more information, see [Billing](https://help.aliyun.com/document_detail/84737.html).
	//
	// example:
	//
	// professional
	SpecType *string `json:"SpecType,omitempty" xml:"SpecType,omitempty"`
	// The number of topics (not recommended).
	//
	// - You must specify either the number of partitions or the topic specification. Specify only the number of partitions.
	//
	// - If you specify both the number of partitions and the topic specification, the system validates whether the number of partitions and the topic specification are equivalent based on the legacy topic sales model. If they are not equivalent, an error is returned. If they are equivalent, the purchase is made based on the number of partitions.
	//
	// - The default value varies based on the traffic specification. Additional fees are charged if the value exceeds the default value.
	//
	// - For the value range, see [Billing](https://help.aliyun.com/document_detail/84737.html).
	//
	// > If the instance is a serverless instance, you do not need to specify this parameter.
	//
	// example:
	//
	// 80
	TopicQuota *int32 `json:"TopicQuota,omitempty" xml:"TopicQuota,omitempty"`
}

func (s UpgradePostPayOrderRequest) String() string {
	return dara.Prettify(s)
}

func (s UpgradePostPayOrderRequest) GoString() string {
	return s.String()
}

func (s *UpgradePostPayOrderRequest) GetDiskSize() *int32 {
	return s.DiskSize
}

func (s *UpgradePostPayOrderRequest) GetEipMax() *int32 {
	return s.EipMax
}

func (s *UpgradePostPayOrderRequest) GetEipModel() *bool {
	return s.EipModel
}

func (s *UpgradePostPayOrderRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpgradePostPayOrderRequest) GetIoMax() *int32 {
	return s.IoMax
}

func (s *UpgradePostPayOrderRequest) GetIoMaxSpec() *string {
	return s.IoMaxSpec
}

func (s *UpgradePostPayOrderRequest) GetPartitionNum() *int32 {
	return s.PartitionNum
}

func (s *UpgradePostPayOrderRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpgradePostPayOrderRequest) GetServerlessConfig() *UpgradePostPayOrderRequestServerlessConfig {
	return s.ServerlessConfig
}

func (s *UpgradePostPayOrderRequest) GetSpecType() *string {
	return s.SpecType
}

func (s *UpgradePostPayOrderRequest) GetTopicQuota() *int32 {
	return s.TopicQuota
}

func (s *UpgradePostPayOrderRequest) SetDiskSize(v int32) *UpgradePostPayOrderRequest {
	s.DiskSize = &v
	return s
}

func (s *UpgradePostPayOrderRequest) SetEipMax(v int32) *UpgradePostPayOrderRequest {
	s.EipMax = &v
	return s
}

func (s *UpgradePostPayOrderRequest) SetEipModel(v bool) *UpgradePostPayOrderRequest {
	s.EipModel = &v
	return s
}

func (s *UpgradePostPayOrderRequest) SetInstanceId(v string) *UpgradePostPayOrderRequest {
	s.InstanceId = &v
	return s
}

func (s *UpgradePostPayOrderRequest) SetIoMax(v int32) *UpgradePostPayOrderRequest {
	s.IoMax = &v
	return s
}

func (s *UpgradePostPayOrderRequest) SetIoMaxSpec(v string) *UpgradePostPayOrderRequest {
	s.IoMaxSpec = &v
	return s
}

func (s *UpgradePostPayOrderRequest) SetPartitionNum(v int32) *UpgradePostPayOrderRequest {
	s.PartitionNum = &v
	return s
}

func (s *UpgradePostPayOrderRequest) SetRegionId(v string) *UpgradePostPayOrderRequest {
	s.RegionId = &v
	return s
}

func (s *UpgradePostPayOrderRequest) SetServerlessConfig(v *UpgradePostPayOrderRequestServerlessConfig) *UpgradePostPayOrderRequest {
	s.ServerlessConfig = v
	return s
}

func (s *UpgradePostPayOrderRequest) SetSpecType(v string) *UpgradePostPayOrderRequest {
	s.SpecType = &v
	return s
}

func (s *UpgradePostPayOrderRequest) SetTopicQuota(v int32) *UpgradePostPayOrderRequest {
	s.TopicQuota = &v
	return s
}

func (s *UpgradePostPayOrderRequest) Validate() error {
	if s.ServerlessConfig != nil {
		if err := s.ServerlessConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpgradePostPayOrderRequestServerlessConfig struct {
	// The reserved publish traffic specification. Only integers are supported. The minimum value is 60. This parameter is required for serverless instances.
	//
	// > The actual upper limit is subject to the inventory in the current region. Refer to the purchase page for the available range.
	//
	// example:
	//
	// 60
	ReservedPublishCapacity *int64 `json:"ReservedPublishCapacity,omitempty" xml:"ReservedPublishCapacity,omitempty"`
	// The reserved subscribe traffic specification. Only integers are supported. The minimum value is 20. This parameter is required for serverless instances.
	//
	// > The actual upper limit is subject to the inventory in the current region. Refer to the purchase page for the available range.
	//
	// example:
	//
	// 60
	ReservedSubscribeCapacity *int64 `json:"ReservedSubscribeCapacity,omitempty" xml:"ReservedSubscribeCapacity,omitempty"`
}

func (s UpgradePostPayOrderRequestServerlessConfig) String() string {
	return dara.Prettify(s)
}

func (s UpgradePostPayOrderRequestServerlessConfig) GoString() string {
	return s.String()
}

func (s *UpgradePostPayOrderRequestServerlessConfig) GetReservedPublishCapacity() *int64 {
	return s.ReservedPublishCapacity
}

func (s *UpgradePostPayOrderRequestServerlessConfig) GetReservedSubscribeCapacity() *int64 {
	return s.ReservedSubscribeCapacity
}

func (s *UpgradePostPayOrderRequestServerlessConfig) SetReservedPublishCapacity(v int64) *UpgradePostPayOrderRequestServerlessConfig {
	s.ReservedPublishCapacity = &v
	return s
}

func (s *UpgradePostPayOrderRequestServerlessConfig) SetReservedSubscribeCapacity(v int64) *UpgradePostPayOrderRequestServerlessConfig {
	s.ReservedSubscribeCapacity = &v
	return s
}

func (s *UpgradePostPayOrderRequestServerlessConfig) Validate() error {
	return dara.Validate(s)
}
