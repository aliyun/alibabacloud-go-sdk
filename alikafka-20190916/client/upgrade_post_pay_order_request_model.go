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
	ServerlessConfig *UpgradePostPayOrderRequestServerlessConfig `json:"ServerlessConfig,omitempty" xml:"ServerlessConfig,omitempty" type:"Struct"`
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
	// The reserved capacity for publishing messages. You can specify only an integer for this parameter. Minimum value: 60.
	//
	// >  The maximum capacity that you can reserve for an instance varies based on available resources in a region. The reserved capacity range displayed on the buy page shall prevail.
	//
	// example:
	//
	// 50
	ReservedPublishCapacity *int64 `json:"ReservedPublishCapacity,omitempty" xml:"ReservedPublishCapacity,omitempty"`
	// The reserved capacity for subscribing to messages. You can specify only an integer for this parameter. Minimum value: 50.
	//
	// >  The maximum capacity that you can reserve for an instance varies based on available resources in a region. The reserved capacity range displayed on the buy page shall prevail.
	//
	// example:
	//
	// 50
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
