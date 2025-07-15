// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePostPayOrderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeployType(v int32) *CreatePostPayOrderRequest
	GetDeployType() *int32
	SetDiskSize(v int32) *CreatePostPayOrderRequest
	GetDiskSize() *int32
	SetDiskType(v string) *CreatePostPayOrderRequest
	GetDiskType() *string
	SetEipMax(v int32) *CreatePostPayOrderRequest
	GetEipMax() *int32
	SetIoMax(v int32) *CreatePostPayOrderRequest
	GetIoMax() *int32
	SetIoMaxSpec(v string) *CreatePostPayOrderRequest
	GetIoMaxSpec() *string
	SetPaidType(v int32) *CreatePostPayOrderRequest
	GetPaidType() *int32
	SetPartitionNum(v int32) *CreatePostPayOrderRequest
	GetPartitionNum() *int32
	SetRegionId(v string) *CreatePostPayOrderRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreatePostPayOrderRequest
	GetResourceGroupId() *string
	SetServerlessConfig(v *CreatePostPayOrderRequestServerlessConfig) *CreatePostPayOrderRequest
	GetServerlessConfig() *CreatePostPayOrderRequestServerlessConfig
	SetSpecType(v string) *CreatePostPayOrderRequest
	GetSpecType() *string
	SetTag(v []*CreatePostPayOrderRequestTag) *CreatePostPayOrderRequest
	GetTag() []*CreatePostPayOrderRequestTag
	SetTopicQuota(v int32) *CreatePostPayOrderRequest
	GetTopicQuota() *int32
}

type CreatePostPayOrderRequest struct {
	// The deployment mode of the instance. Valid values:
	//
	// 	- **4**: deploys the instance that allows access from the Internet and a VPC.
	//
	// 	- **5**: deploys the instance that allows access only from a VPC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5
	DeployType *int32 `json:"DeployType,omitempty" xml:"DeployType,omitempty"`
	// The disk size.
	//
	// For information about the valid values of this parameter, see [Billing](https://help.aliyun.com/document_detail/84737.html).
	//
	// >  If you create a serverless ApsaraMQ for Kafka instance, you do not need to configure this parameter.
	//
	// example:
	//
	// 500
	DiskSize *int32 `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
	// The disk type of the instance. Valid values:
	//
	// 	- **0**: ultra disk
	//
	// 	- **1**: standard SSD
	//
	// >  If you create a serverless ApsaraMQ for Kafka instance, you do not need to configure this parameter.
	//
	// example:
	//
	// 0
	DiskType *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	// The Internet traffic.
	//
	// 	- If you set **DeployType*	- to **4**, you must configure this parameter.
	//
	// 	- For information about the valid values of this parameter, see [Billing](https://help.aliyun.com/document_detail/84737.html).
	//
	// >  If you create a serverless ApsaraMQ for Kafka instance, you do not need to configure this parameter.
	//
	// example:
	//
	// 0
	EipMax *int32 `json:"EipMax,omitempty" xml:"EipMax,omitempty"`
	// The maximum traffic in the instance. We recommend that you do not configure this parameter.
	//
	// 	- You must configure at least one of IoMax and IoMaxSpec. If you configure both parameters, the value of IoMaxSpec takes effect. We recommend that you configure only IoMaxSpec.
	//
	// 	- For information about the valid values of this parameter, see [Billing](https://help.aliyun.com/document_detail/84737.html).
	//
	// >  If you create a serverless ApsaraMQ for Kafka instance, you do not need to configure this parameter.
	//
	// example:
	//
	// 20
	IoMax *int32 `json:"IoMax,omitempty" xml:"IoMax,omitempty"`
	// The traffic specification of the instance. We recommend that you configure this parameter.
	//
	// 	- You must configure at least one of IoMax and IoMaxSpec. If you configure both parameters, the value of IoMaxSpec takes effect. We recommend that you configure only IoMaxSpec.
	//
	// 	- For information about the valid values of this parameter, see [Billing](https://help.aliyun.com/document_detail/84737.html).
	//
	// >  If you create a serverless ApsaraMQ for Kafka instance, you do not need to configure this parameter.
	//
	// example:
	//
	// alikafka.hw.2xlarge
	IoMaxSpec *string `json:"IoMaxSpec,omitempty" xml:"IoMaxSpec,omitempty"`
	// The billing method of the instance. Valid values:
	//
	// 	- 1: pay-as-you-go (reserved capacity).
	//
	// 	- 3: pay-as-you-go (reserved capacity) + pay-as-you-go (on-demand capacity)
	//
	// example:
	//
	// 1
	PaidType *int32 `json:"PaidType,omitempty" xml:"PaidType,omitempty"`
	// The number of partitions. We recommend that you configure this parameter.
	//
	// 	- You must configure one of PartitionNum and TopicQuota. We recommend that you configure only ParittionNum.
	//
	// 	- If you configure PartitionNum and TopicQuota at the same time, the system verifies whether the price of the partitions equals the price of the topics based on the previous topic-based selling mode. If the price of the partitions does not equal the price of the topics, an error is returned. If the price of the partitions equals the price of the topics, the instance is purchased based on the partition number.
	//
	// 	- For information about the valid values of this parameter, see [Billing](https://help.aliyun.com/document_detail/84737.html).
	//
	// >  If you create a serverless ApsaraMQ for Kafka instance, you do not need to configure this parameter.
	//
	// example:
	//
	// 50
	PartitionNum *int32 `json:"PartitionNum,omitempty" xml:"PartitionNum,omitempty"`
	// The region ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// If this parameter is left empty, the default resource group is used. You can view the resource group ID on the Resource Group page in the Resource Management console.
	//
	// example:
	//
	// rg-ac***********7q
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The parameters configured for the serverless ApsaraMQ for Kafka instance. These parameters are required only when you create a serverless instance.
	ServerlessConfig *CreatePostPayOrderRequestServerlessConfig `json:"ServerlessConfig,omitempty" xml:"ServerlessConfig,omitempty" type:"Struct"`
	// The instance edition.
	//
	// Valid values if you set PaidType to 1:
	//
	// 	- normal: Standard Edition (High Write)
	//
	// 	- professional: Professional Edition (High Write)
	//
	// 	- professionalForHighRead: Professional Edition (High Read)
	//
	// Valid values if you set PaidType to 3:
	//
	// 	- normal: Serverless Standard Edition
	//
	// For more information about the instance editions, see [Billing](https://help.aliyun.com/document_detail/84737.html).
	//
	// example:
	//
	// normal
	SpecType *string `json:"SpecType,omitempty" xml:"SpecType,omitempty"`
	// The tags.
	Tag []*CreatePostPayOrderRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The number of topics. We recommend that you do not configure this parameter.
	//
	// 	- You must configure one of PartitionNum and TopicQuota. We recommend that you configure only ParittionNum.
	//
	// 	- If you configure PartitionNum and TopicQuota at the same time, the system verifies whether the price of the partitions equals the price of the topics based on the previous topic-based selling mode. If the price of the partitions does not equal the price of the topics, an error is returned. If the price of the partitions equals the price of the topics, the instance is purchased based on the partition number.
	//
	// 	- The default value of TopicQuota varies based on the value of IoMaxSpec. If the number of topics that you consume exceeds the default value, you are charged additional fees.
	//
	// 	- For information about the valid values of this parameter, see [Billing](https://help.aliyun.com/document_detail/84737.html).
	//
	// >  If you create a serverless ApsaraMQ for Kafka instance, you do not need to configure this parameter.
	//
	// example:
	//
	// 50
	TopicQuota *int32 `json:"TopicQuota,omitempty" xml:"TopicQuota,omitempty"`
}

func (s CreatePostPayOrderRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePostPayOrderRequest) GoString() string {
	return s.String()
}

func (s *CreatePostPayOrderRequest) GetDeployType() *int32 {
	return s.DeployType
}

func (s *CreatePostPayOrderRequest) GetDiskSize() *int32 {
	return s.DiskSize
}

func (s *CreatePostPayOrderRequest) GetDiskType() *string {
	return s.DiskType
}

func (s *CreatePostPayOrderRequest) GetEipMax() *int32 {
	return s.EipMax
}

func (s *CreatePostPayOrderRequest) GetIoMax() *int32 {
	return s.IoMax
}

func (s *CreatePostPayOrderRequest) GetIoMaxSpec() *string {
	return s.IoMaxSpec
}

func (s *CreatePostPayOrderRequest) GetPaidType() *int32 {
	return s.PaidType
}

func (s *CreatePostPayOrderRequest) GetPartitionNum() *int32 {
	return s.PartitionNum
}

func (s *CreatePostPayOrderRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreatePostPayOrderRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreatePostPayOrderRequest) GetServerlessConfig() *CreatePostPayOrderRequestServerlessConfig {
	return s.ServerlessConfig
}

func (s *CreatePostPayOrderRequest) GetSpecType() *string {
	return s.SpecType
}

func (s *CreatePostPayOrderRequest) GetTag() []*CreatePostPayOrderRequestTag {
	return s.Tag
}

func (s *CreatePostPayOrderRequest) GetTopicQuota() *int32 {
	return s.TopicQuota
}

func (s *CreatePostPayOrderRequest) SetDeployType(v int32) *CreatePostPayOrderRequest {
	s.DeployType = &v
	return s
}

func (s *CreatePostPayOrderRequest) SetDiskSize(v int32) *CreatePostPayOrderRequest {
	s.DiskSize = &v
	return s
}

func (s *CreatePostPayOrderRequest) SetDiskType(v string) *CreatePostPayOrderRequest {
	s.DiskType = &v
	return s
}

func (s *CreatePostPayOrderRequest) SetEipMax(v int32) *CreatePostPayOrderRequest {
	s.EipMax = &v
	return s
}

func (s *CreatePostPayOrderRequest) SetIoMax(v int32) *CreatePostPayOrderRequest {
	s.IoMax = &v
	return s
}

func (s *CreatePostPayOrderRequest) SetIoMaxSpec(v string) *CreatePostPayOrderRequest {
	s.IoMaxSpec = &v
	return s
}

func (s *CreatePostPayOrderRequest) SetPaidType(v int32) *CreatePostPayOrderRequest {
	s.PaidType = &v
	return s
}

func (s *CreatePostPayOrderRequest) SetPartitionNum(v int32) *CreatePostPayOrderRequest {
	s.PartitionNum = &v
	return s
}

func (s *CreatePostPayOrderRequest) SetRegionId(v string) *CreatePostPayOrderRequest {
	s.RegionId = &v
	return s
}

func (s *CreatePostPayOrderRequest) SetResourceGroupId(v string) *CreatePostPayOrderRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreatePostPayOrderRequest) SetServerlessConfig(v *CreatePostPayOrderRequestServerlessConfig) *CreatePostPayOrderRequest {
	s.ServerlessConfig = v
	return s
}

func (s *CreatePostPayOrderRequest) SetSpecType(v string) *CreatePostPayOrderRequest {
	s.SpecType = &v
	return s
}

func (s *CreatePostPayOrderRequest) SetTag(v []*CreatePostPayOrderRequestTag) *CreatePostPayOrderRequest {
	s.Tag = v
	return s
}

func (s *CreatePostPayOrderRequest) SetTopicQuota(v int32) *CreatePostPayOrderRequest {
	s.TopicQuota = &v
	return s
}

func (s *CreatePostPayOrderRequest) Validate() error {
	return dara.Validate(s)
}

type CreatePostPayOrderRequestServerlessConfig struct {
	// The reserved capacity for publishing messages. You can specify only an integer for this parameter. Minimum value: 60.
	//
	// >  The actual maximum reserved capacity for publishing messages varies based on available resources in the region. The actual range displayed on the buy page shall prevail.
	//
	// example:
	//
	// 60
	ReservedPublishCapacity *int64 `json:"ReservedPublishCapacity,omitempty" xml:"ReservedPublishCapacity,omitempty"`
	// The reserved capacity for subscribing to messages. You can specify only an integer for this parameter. Minimum value: 20.
	//
	// >  The actual maximum reserved capacity for subscribing to messages varies based on available resources in the region. The actual range displayed on the buy page shall prevail.
	//
	// example:
	//
	// 50
	ReservedSubscribeCapacity *int64 `json:"ReservedSubscribeCapacity,omitempty" xml:"ReservedSubscribeCapacity,omitempty"`
}

func (s CreatePostPayOrderRequestServerlessConfig) String() string {
	return dara.Prettify(s)
}

func (s CreatePostPayOrderRequestServerlessConfig) GoString() string {
	return s.String()
}

func (s *CreatePostPayOrderRequestServerlessConfig) GetReservedPublishCapacity() *int64 {
	return s.ReservedPublishCapacity
}

func (s *CreatePostPayOrderRequestServerlessConfig) GetReservedSubscribeCapacity() *int64 {
	return s.ReservedSubscribeCapacity
}

func (s *CreatePostPayOrderRequestServerlessConfig) SetReservedPublishCapacity(v int64) *CreatePostPayOrderRequestServerlessConfig {
	s.ReservedPublishCapacity = &v
	return s
}

func (s *CreatePostPayOrderRequestServerlessConfig) SetReservedSubscribeCapacity(v int64) *CreatePostPayOrderRequestServerlessConfig {
	s.ReservedSubscribeCapacity = &v
	return s
}

func (s *CreatePostPayOrderRequestServerlessConfig) Validate() error {
	return dara.Validate(s)
}

type CreatePostPayOrderRequestTag struct {
	// The key of tag N.
	//
	// 	- Valid values of N: 1 to 20.
	//
	// 	- If this parameter is left empty, the keys of all tags are matched.
	//
	// 	- The tag key must be up to 128 characters in length. It cannot start with acs: or aliyun or contain [http:// or https://.](http://https://。)
	//
	// This parameter is required.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N.
	//
	// 	- Valid values of N: 1 to 20.
	//
	// 	- If you do not specify a tag key, you cannot specify a tag value. If this parameter is not configured, all tag values are matched.
	//
	// 	- The tag value must be 1 to 128 characters in length. It cannot start with acs: or aliyun or contain [http:// or https://.](http://https://。)
	//
	// example:
	//
	// FinanceJoshua
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreatePostPayOrderRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreatePostPayOrderRequestTag) GoString() string {
	return s.String()
}

func (s *CreatePostPayOrderRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreatePostPayOrderRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreatePostPayOrderRequestTag) SetKey(v string) *CreatePostPayOrderRequestTag {
	s.Key = &v
	return s
}

func (s *CreatePostPayOrderRequestTag) SetValue(v string) *CreatePostPayOrderRequestTag {
	s.Value = &v
	return s
}

func (s *CreatePostPayOrderRequestTag) Validate() error {
	return dara.Validate(s)
}
