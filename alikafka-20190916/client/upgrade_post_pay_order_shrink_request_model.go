// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradePostPayOrderShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDiskSize(v int32) *UpgradePostPayOrderShrinkRequest
	GetDiskSize() *int32
	SetEipMax(v int32) *UpgradePostPayOrderShrinkRequest
	GetEipMax() *int32
	SetEipModel(v bool) *UpgradePostPayOrderShrinkRequest
	GetEipModel() *bool
	SetInstanceId(v string) *UpgradePostPayOrderShrinkRequest
	GetInstanceId() *string
	SetIoMax(v int32) *UpgradePostPayOrderShrinkRequest
	GetIoMax() *int32
	SetIoMaxSpec(v string) *UpgradePostPayOrderShrinkRequest
	GetIoMaxSpec() *string
	SetPartitionNum(v int32) *UpgradePostPayOrderShrinkRequest
	GetPartitionNum() *int32
	SetRegionId(v string) *UpgradePostPayOrderShrinkRequest
	GetRegionId() *string
	SetServerlessConfigShrink(v string) *UpgradePostPayOrderShrinkRequest
	GetServerlessConfigShrink() *string
	SetSpecType(v string) *UpgradePostPayOrderShrinkRequest
	GetSpecType() *string
	SetTopicQuota(v int32) *UpgradePostPayOrderShrinkRequest
	GetTopicQuota() *int32
}

type UpgradePostPayOrderShrinkRequest struct {
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
	ServerlessConfigShrink *string `json:"ServerlessConfig,omitempty" xml:"ServerlessConfig,omitempty"`
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

func (s UpgradePostPayOrderShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpgradePostPayOrderShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpgradePostPayOrderShrinkRequest) GetDiskSize() *int32 {
	return s.DiskSize
}

func (s *UpgradePostPayOrderShrinkRequest) GetEipMax() *int32 {
	return s.EipMax
}

func (s *UpgradePostPayOrderShrinkRequest) GetEipModel() *bool {
	return s.EipModel
}

func (s *UpgradePostPayOrderShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpgradePostPayOrderShrinkRequest) GetIoMax() *int32 {
	return s.IoMax
}

func (s *UpgradePostPayOrderShrinkRequest) GetIoMaxSpec() *string {
	return s.IoMaxSpec
}

func (s *UpgradePostPayOrderShrinkRequest) GetPartitionNum() *int32 {
	return s.PartitionNum
}

func (s *UpgradePostPayOrderShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpgradePostPayOrderShrinkRequest) GetServerlessConfigShrink() *string {
	return s.ServerlessConfigShrink
}

func (s *UpgradePostPayOrderShrinkRequest) GetSpecType() *string {
	return s.SpecType
}

func (s *UpgradePostPayOrderShrinkRequest) GetTopicQuota() *int32 {
	return s.TopicQuota
}

func (s *UpgradePostPayOrderShrinkRequest) SetDiskSize(v int32) *UpgradePostPayOrderShrinkRequest {
	s.DiskSize = &v
	return s
}

func (s *UpgradePostPayOrderShrinkRequest) SetEipMax(v int32) *UpgradePostPayOrderShrinkRequest {
	s.EipMax = &v
	return s
}

func (s *UpgradePostPayOrderShrinkRequest) SetEipModel(v bool) *UpgradePostPayOrderShrinkRequest {
	s.EipModel = &v
	return s
}

func (s *UpgradePostPayOrderShrinkRequest) SetInstanceId(v string) *UpgradePostPayOrderShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *UpgradePostPayOrderShrinkRequest) SetIoMax(v int32) *UpgradePostPayOrderShrinkRequest {
	s.IoMax = &v
	return s
}

func (s *UpgradePostPayOrderShrinkRequest) SetIoMaxSpec(v string) *UpgradePostPayOrderShrinkRequest {
	s.IoMaxSpec = &v
	return s
}

func (s *UpgradePostPayOrderShrinkRequest) SetPartitionNum(v int32) *UpgradePostPayOrderShrinkRequest {
	s.PartitionNum = &v
	return s
}

func (s *UpgradePostPayOrderShrinkRequest) SetRegionId(v string) *UpgradePostPayOrderShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *UpgradePostPayOrderShrinkRequest) SetServerlessConfigShrink(v string) *UpgradePostPayOrderShrinkRequest {
	s.ServerlessConfigShrink = &v
	return s
}

func (s *UpgradePostPayOrderShrinkRequest) SetSpecType(v string) *UpgradePostPayOrderShrinkRequest {
	s.SpecType = &v
	return s
}

func (s *UpgradePostPayOrderShrinkRequest) SetTopicQuota(v int32) *UpgradePostPayOrderShrinkRequest {
	s.TopicQuota = &v
	return s
}

func (s *UpgradePostPayOrderShrinkRequest) Validate() error {
	return dara.Validate(s)
}
