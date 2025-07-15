// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePostPayOrderShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeployType(v int32) *CreatePostPayOrderShrinkRequest
	GetDeployType() *int32
	SetDiskSize(v int32) *CreatePostPayOrderShrinkRequest
	GetDiskSize() *int32
	SetDiskType(v string) *CreatePostPayOrderShrinkRequest
	GetDiskType() *string
	SetEipMax(v int32) *CreatePostPayOrderShrinkRequest
	GetEipMax() *int32
	SetIoMax(v int32) *CreatePostPayOrderShrinkRequest
	GetIoMax() *int32
	SetIoMaxSpec(v string) *CreatePostPayOrderShrinkRequest
	GetIoMaxSpec() *string
	SetPaidType(v int32) *CreatePostPayOrderShrinkRequest
	GetPaidType() *int32
	SetPartitionNum(v int32) *CreatePostPayOrderShrinkRequest
	GetPartitionNum() *int32
	SetRegionId(v string) *CreatePostPayOrderShrinkRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreatePostPayOrderShrinkRequest
	GetResourceGroupId() *string
	SetServerlessConfigShrink(v string) *CreatePostPayOrderShrinkRequest
	GetServerlessConfigShrink() *string
	SetSpecType(v string) *CreatePostPayOrderShrinkRequest
	GetSpecType() *string
	SetTag(v []*CreatePostPayOrderShrinkRequestTag) *CreatePostPayOrderShrinkRequest
	GetTag() []*CreatePostPayOrderShrinkRequestTag
	SetTopicQuota(v int32) *CreatePostPayOrderShrinkRequest
	GetTopicQuota() *int32
}

type CreatePostPayOrderShrinkRequest struct {
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
	ServerlessConfigShrink *string `json:"ServerlessConfig,omitempty" xml:"ServerlessConfig,omitempty"`
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
	Tag []*CreatePostPayOrderShrinkRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
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

func (s CreatePostPayOrderShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePostPayOrderShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreatePostPayOrderShrinkRequest) GetDeployType() *int32 {
	return s.DeployType
}

func (s *CreatePostPayOrderShrinkRequest) GetDiskSize() *int32 {
	return s.DiskSize
}

func (s *CreatePostPayOrderShrinkRequest) GetDiskType() *string {
	return s.DiskType
}

func (s *CreatePostPayOrderShrinkRequest) GetEipMax() *int32 {
	return s.EipMax
}

func (s *CreatePostPayOrderShrinkRequest) GetIoMax() *int32 {
	return s.IoMax
}

func (s *CreatePostPayOrderShrinkRequest) GetIoMaxSpec() *string {
	return s.IoMaxSpec
}

func (s *CreatePostPayOrderShrinkRequest) GetPaidType() *int32 {
	return s.PaidType
}

func (s *CreatePostPayOrderShrinkRequest) GetPartitionNum() *int32 {
	return s.PartitionNum
}

func (s *CreatePostPayOrderShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreatePostPayOrderShrinkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreatePostPayOrderShrinkRequest) GetServerlessConfigShrink() *string {
	return s.ServerlessConfigShrink
}

func (s *CreatePostPayOrderShrinkRequest) GetSpecType() *string {
	return s.SpecType
}

func (s *CreatePostPayOrderShrinkRequest) GetTag() []*CreatePostPayOrderShrinkRequestTag {
	return s.Tag
}

func (s *CreatePostPayOrderShrinkRequest) GetTopicQuota() *int32 {
	return s.TopicQuota
}

func (s *CreatePostPayOrderShrinkRequest) SetDeployType(v int32) *CreatePostPayOrderShrinkRequest {
	s.DeployType = &v
	return s
}

func (s *CreatePostPayOrderShrinkRequest) SetDiskSize(v int32) *CreatePostPayOrderShrinkRequest {
	s.DiskSize = &v
	return s
}

func (s *CreatePostPayOrderShrinkRequest) SetDiskType(v string) *CreatePostPayOrderShrinkRequest {
	s.DiskType = &v
	return s
}

func (s *CreatePostPayOrderShrinkRequest) SetEipMax(v int32) *CreatePostPayOrderShrinkRequest {
	s.EipMax = &v
	return s
}

func (s *CreatePostPayOrderShrinkRequest) SetIoMax(v int32) *CreatePostPayOrderShrinkRequest {
	s.IoMax = &v
	return s
}

func (s *CreatePostPayOrderShrinkRequest) SetIoMaxSpec(v string) *CreatePostPayOrderShrinkRequest {
	s.IoMaxSpec = &v
	return s
}

func (s *CreatePostPayOrderShrinkRequest) SetPaidType(v int32) *CreatePostPayOrderShrinkRequest {
	s.PaidType = &v
	return s
}

func (s *CreatePostPayOrderShrinkRequest) SetPartitionNum(v int32) *CreatePostPayOrderShrinkRequest {
	s.PartitionNum = &v
	return s
}

func (s *CreatePostPayOrderShrinkRequest) SetRegionId(v string) *CreatePostPayOrderShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreatePostPayOrderShrinkRequest) SetResourceGroupId(v string) *CreatePostPayOrderShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreatePostPayOrderShrinkRequest) SetServerlessConfigShrink(v string) *CreatePostPayOrderShrinkRequest {
	s.ServerlessConfigShrink = &v
	return s
}

func (s *CreatePostPayOrderShrinkRequest) SetSpecType(v string) *CreatePostPayOrderShrinkRequest {
	s.SpecType = &v
	return s
}

func (s *CreatePostPayOrderShrinkRequest) SetTag(v []*CreatePostPayOrderShrinkRequestTag) *CreatePostPayOrderShrinkRequest {
	s.Tag = v
	return s
}

func (s *CreatePostPayOrderShrinkRequest) SetTopicQuota(v int32) *CreatePostPayOrderShrinkRequest {
	s.TopicQuota = &v
	return s
}

func (s *CreatePostPayOrderShrinkRequest) Validate() error {
	return dara.Validate(s)
}

type CreatePostPayOrderShrinkRequestTag struct {
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

func (s CreatePostPayOrderShrinkRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreatePostPayOrderShrinkRequestTag) GoString() string {
	return s.String()
}

func (s *CreatePostPayOrderShrinkRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreatePostPayOrderShrinkRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreatePostPayOrderShrinkRequestTag) SetKey(v string) *CreatePostPayOrderShrinkRequestTag {
	s.Key = &v
	return s
}

func (s *CreatePostPayOrderShrinkRequestTag) SetValue(v string) *CreatePostPayOrderShrinkRequestTag {
	s.Value = &v
	return s
}

func (s *CreatePostPayOrderShrinkRequestTag) Validate() error {
	return dara.Validate(s)
}
