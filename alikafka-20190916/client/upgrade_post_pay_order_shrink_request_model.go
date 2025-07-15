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
	// The disk size. Unit: GB.
	//
	// 	- The disk size that you specify must be greater than or equal to the current disk size of the instance.
	//
	// 	- For information about the valid values of this parameter, see [Billing](https://help.aliyun.com/document_detail/84737.html).
	//
	// >  If the instance is a serverless ApsaraMQ for Kafka instance, you do not need to configure this parameter.
	//
	// example:
	//
	// 500
	DiskSize *int32 `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
	// The maximum Internet traffic of the instance.
	//
	// 	- The Internet traffic that you specify must be greater than or equal to the current Internet traffic of the instance.
	//
	// 	- For information about the valid values of this parameter, see [Billing](https://help.aliyun.com/document_detail/84737.html).
	//
	// >
	//
	// 	- If you set **EipModel*	- to **true**, set **EipMax*	- to a value that is greater than 0.
	//
	// 	- If you set **EipModel*	- to **false**, set **EipMax*	- to **0**.
	//
	// 	- If the instance is a serverless ApsaraMQ for Kafka instance, you do not need to configure this parameter.
	//
	// example:
	//
	// 0
	EipMax *int32 `json:"EipMax,omitempty" xml:"EipMax,omitempty"`
	// Specifies whether to enable Internet access for the instance. Valid values:
	//
	// 	- true: enables Internet access.
	//
	// 	- false: disables Internet access.
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
	// The maximum traffic of the instance. We recommend that you do not configure this parameter.
	//
	// 	- The maximum traffic that you specify must be greater than or equal to the current maximum traffic of the instance.
	//
	// 	- You must configure at least one of IoMax and IoMaxSpec. If you configure both parameters, the value of IoMaxSpec takes effect. We recommend that you configure only IoMaxSpec.
	//
	// 	- For information about the valid values of this parameter, see [Billing](https://help.aliyun.com/document_detail/84737.html).
	//
	// >  If the instance is a serverless ApsaraMQ for Kafka instance, you do not need to configure this parameter.
	//
	// example:
	//
	// 60
	IoMax *int32 `json:"IoMax,omitempty" xml:"IoMax,omitempty"`
	// The traffic specification of the instance. We recommend that you configure this parameter.
	//
	// 	- The traffic specification that you specify must be greater than or equal to the current traffic specification of the instance.
	//
	// 	- You must configure at least one of IoMax and IoMaxSpec. If you configure both parameters, the value of IoMaxSpec takes effect. We recommend that you configure only IoMaxSpec.
	//
	// 	- For information about the valid values of this parameter, see [Billing](https://help.aliyun.com/document_detail/84737.html).
	//
	// >  If the instance is a serverless ApsaraMQ for Kafka instance, you do not need to configure this parameter.
	//
	// example:
	//
	// alikafka.hw.6xlarge
	IoMaxSpec *string `json:"IoMaxSpec,omitempty" xml:"IoMaxSpec,omitempty"`
	// The number of partitions. We recommend that you configure this parameter.
	//
	// 	- You must configure one of PartitionNum and TopicQuota. We recommend that you configure only PartitionNum.
	//
	// 	- If you configure PartitionNum and TopicQuota at the same time, the system verifies whether the price of the partitions equals the price of the topics based on the previous topic-based selling mode. If the price of the partitions does not equal the price of the topics, an error is returned. If the price of the partitions equals the price of the topics, the instance is purchased based on the partition number.
	//
	// 	- For information about the valid values of this parameter, see [Billing](https://help.aliyun.com/document_detail/84737.html).
	//
	// >  If the instance is a serverless ApsaraMQ for Kafka instance, you do not need to configure this parameter.
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
	// The parameters that are configured for the serverless instance. These parameters are required only when you create a serverless instance.
	ServerlessConfigShrink *string `json:"ServerlessConfig,omitempty" xml:"ServerlessConfig,omitempty"`
	// The instance edition.
	//
	// Valid values for this parameter if you set PaidType to 1:
	//
	// 	- normal: Standard Edition (High Write)
	//
	// 	- professional: Professional Edition (High Write)
	//
	// 	- professionalForHighRead: Professional Edition (High Read)
	//
	// Valid values for this parameter if you set PaidType to 3:
	//
	// 	- normal: Serverless Standard Edition
	//
	// For more information, see [Billing](https://help.aliyun.com/document_detail/84737.html).
	//
	// example:
	//
	// professional
	SpecType *string `json:"SpecType,omitempty" xml:"SpecType,omitempty"`
	// The number of topics. We recommend that you do not configure this parameter.
	//
	// 	- You must configure one of PartitionNum and TopicQuota. We recommend that you configure only PartitionNum.
	//
	// 	- If you configure PartitionNum and TopicQuota at the same time, the system verifies whether the price of the partitions equals the price of the topics based on the previous topic-based selling mode. If the price of the partitions does not equal the price of the topics, an error is returned. If the price of the partitions equals the price of the topics, the instance is purchased based on the partition number.
	//
	// 	- The default value of TopicQuota varies based on the value of IoMaxSpec. If the number of topics that you use exceeds the default value, you are charged additional fees.
	//
	// 	- For information about the valid values of this parameter, see [Billing](https://help.aliyun.com/document_detail/84737.html).
	//
	// >  If the instance is a serverless ApsaraMQ for Kafka instance, you do not need to configure this parameter.
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
