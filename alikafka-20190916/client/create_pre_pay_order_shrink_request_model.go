// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePrePayOrderShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfluentConfigShrink(v string) *CreatePrePayOrderShrinkRequest
	GetConfluentConfigShrink() *string
	SetDeployType(v int32) *CreatePrePayOrderShrinkRequest
	GetDeployType() *int32
	SetDiskSize(v int32) *CreatePrePayOrderShrinkRequest
	GetDiskSize() *int32
	SetDiskType(v string) *CreatePrePayOrderShrinkRequest
	GetDiskType() *string
	SetDuration(v int32) *CreatePrePayOrderShrinkRequest
	GetDuration() *int32
	SetEipMax(v int32) *CreatePrePayOrderShrinkRequest
	GetEipMax() *int32
	SetIoMax(v int32) *CreatePrePayOrderShrinkRequest
	GetIoMax() *int32
	SetIoMaxSpec(v string) *CreatePrePayOrderShrinkRequest
	GetIoMaxSpec() *string
	SetPaidType(v int32) *CreatePrePayOrderShrinkRequest
	GetPaidType() *int32
	SetPartitionNum(v int32) *CreatePrePayOrderShrinkRequest
	GetPartitionNum() *int32
	SetRegionId(v string) *CreatePrePayOrderShrinkRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreatePrePayOrderShrinkRequest
	GetResourceGroupId() *string
	SetSpecType(v string) *CreatePrePayOrderShrinkRequest
	GetSpecType() *string
	SetTag(v []*CreatePrePayOrderShrinkRequestTag) *CreatePrePayOrderShrinkRequest
	GetTag() []*CreatePrePayOrderShrinkRequestTag
	SetTopicQuota(v int32) *CreatePrePayOrderShrinkRequest
	GetTopicQuota() *int32
}

type CreatePrePayOrderShrinkRequest struct {
	// The configurations of Confluent.
	//
	// >  When you create an ApsaraMQ for Confluent instance, you must configure this parameter.
	ConfluentConfigShrink *string `json:"ConfluentConfig,omitempty" xml:"ConfluentConfig,omitempty"`
	// The type of the network in which the instance is deployed. Valid values:
	//
	// 	- **4**: Internet and virtual private cloud (VPC)
	//
	// 	- **5**: VPC
	//
	// >  If you create an ApsaraMQ for Confluent instance, set the value to 5. After the instance is created, you can specify whether to enable each component.
	//
	// example:
	//
	// 5
	DeployType *int32 `json:"DeployType,omitempty" xml:"DeployType,omitempty"`
	// The disk size. Unit: GB
	//
	// For information about the valid values of this parameter, see [Billing](https://help.aliyun.com/document_detail/84737.html).
	//
	// >  If you create an ApsaraMQ for Confluent instance, you do not need to configure this parameter.
	//
	// example:
	//
	// 500
	DiskSize *int32 `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
	// The disk type. Valid values:
	//
	// 	- **0**: ultra disk
	//
	// 	- **1**: standard SSD
	//
	// >  If you create an ApsaraMQ for Confluent instance, you do not need to configure this parameter.
	//
	// example:
	//
	// 0
	DiskType *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	// The subscription duration. Unit: months. Default value: 1. Valid values:
	//
	// 	- **1 to 12**
	//
	// example:
	//
	// 1
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The maximum Internet traffic in the instance.
	//
	// 	- If you set **DeployType*	- to **4**, you must configure this parameter.
	//
	// 	- For information about the valid values, see [Pay-as-you-go](https://help.aliyun.com/document_detail/72142.html).
	//
	// >  If you create an ApsaraMQ for Confluent instance, you do not need to configure this parameter.
	//
	// example:
	//
	// 0
	EipMax *int32 `json:"EipMax,omitempty" xml:"EipMax,omitempty"`
	// The maximum traffic in the instance. We recommend that you do not configure this parameter.
	//
	// 	- You must set one of **IoMax*	- and **IoMaxSpec**. If both parameters are configured, the value of **IoMaxSpec*	- is used. We recommend that you configure only **IoMaxSpec**.
	//
	// 	- For information about the valid values of this parameter, see [Billing](https://help.aliyun.com/document_detail/84737.html).
	//
	// >  If you create an ApsaraMQ for Confluent instance, you do not need to configure this parameter.
	//
	// example:
	//
	// 20
	IoMax *int32 `json:"IoMax,omitempty" xml:"IoMax,omitempty"`
	// The traffic specification of the instance. We recommend that you configure this parameter.
	//
	// 	- You must configure one of **IoMax*	- and **IoMaxSpec**. If both parameters are configured, the value of **IoMaxSpec*	- is used. We recommend that you configure only **IoMaxSpec**.
	//
	// 	- For more information about the valid values of this parameter, see [Billing](https://help.aliyun.com/document_detail/84737.html).
	//
	// >  If you create an ApsaraMQ for Confluent instance, you do not need to configure this parameter.
	//
	// example:
	//
	// alikafka.hw.2xlarge
	IoMaxSpec *string `json:"IoMaxSpec,omitempty" xml:"IoMaxSpec,omitempty"`
	// The billing method of the instance. Valid values:
	//
	// 	- **0**: the subscription billing method
	//
	// 	- **4**: the subscription billing method for ApsaraMQ for Confluent instances
	//
	// example:
	//
	// 1
	PaidType *int32 `json:"PaidType,omitempty" xml:"PaidType,omitempty"`
	// The number of partitions. We recommend that you configure this parameter.
	//
	// 	- You must configure one of PartitionNum and TopicQuota. We recommend that you configure only PartitionNum.
	//
	// 	- If you configure PartitionNum and TopicQuota at the same time, the system verifies whether the price of the partitions equals the price of the topics based on the previous topic-based selling mode. If the price of the partitions does not equal the price of the topics, an error is returned. If the price of the partitions equals the price of the topics, the instance is purchased based on the partition number.
	//
	// 	- For information about the valid values of this parameter, see [Billing](https://help.aliyun.com/document_detail/84737.html).
	//
	// >  If you create an ApsaraMQ for Confluent instance, you do not need to configure this parameter.
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
	// The instance edition. Valid values:
	//
	// 	- **normal**: Standard Edition (High Write)
	//
	// 	- **professional**: Professional Edition (High Write)
	//
	// 	- **professionalForHighRead**: Professional Edition (High Read)
	//
	// For more information, see [Billing](https://help.aliyun.com/document_detail/84737.html).
	//
	// >  If you create an ApsaraMQ for Confluent instance, you do not need to configure this parameter.
	//
	// example:
	//
	// normal
	SpecType *string `json:"SpecType,omitempty" xml:"SpecType,omitempty"`
	// The tags.
	Tag []*CreatePrePayOrderShrinkRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
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
	// >  If you create an ApsaraMQ for Confluent instance, you do not need to configure this parameter.
	//
	// example:
	//
	// 50
	TopicQuota *int32 `json:"TopicQuota,omitempty" xml:"TopicQuota,omitempty"`
}

func (s CreatePrePayOrderShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePrePayOrderShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreatePrePayOrderShrinkRequest) GetConfluentConfigShrink() *string {
	return s.ConfluentConfigShrink
}

func (s *CreatePrePayOrderShrinkRequest) GetDeployType() *int32 {
	return s.DeployType
}

func (s *CreatePrePayOrderShrinkRequest) GetDiskSize() *int32 {
	return s.DiskSize
}

func (s *CreatePrePayOrderShrinkRequest) GetDiskType() *string {
	return s.DiskType
}

func (s *CreatePrePayOrderShrinkRequest) GetDuration() *int32 {
	return s.Duration
}

func (s *CreatePrePayOrderShrinkRequest) GetEipMax() *int32 {
	return s.EipMax
}

func (s *CreatePrePayOrderShrinkRequest) GetIoMax() *int32 {
	return s.IoMax
}

func (s *CreatePrePayOrderShrinkRequest) GetIoMaxSpec() *string {
	return s.IoMaxSpec
}

func (s *CreatePrePayOrderShrinkRequest) GetPaidType() *int32 {
	return s.PaidType
}

func (s *CreatePrePayOrderShrinkRequest) GetPartitionNum() *int32 {
	return s.PartitionNum
}

func (s *CreatePrePayOrderShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreatePrePayOrderShrinkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreatePrePayOrderShrinkRequest) GetSpecType() *string {
	return s.SpecType
}

func (s *CreatePrePayOrderShrinkRequest) GetTag() []*CreatePrePayOrderShrinkRequestTag {
	return s.Tag
}

func (s *CreatePrePayOrderShrinkRequest) GetTopicQuota() *int32 {
	return s.TopicQuota
}

func (s *CreatePrePayOrderShrinkRequest) SetConfluentConfigShrink(v string) *CreatePrePayOrderShrinkRequest {
	s.ConfluentConfigShrink = &v
	return s
}

func (s *CreatePrePayOrderShrinkRequest) SetDeployType(v int32) *CreatePrePayOrderShrinkRequest {
	s.DeployType = &v
	return s
}

func (s *CreatePrePayOrderShrinkRequest) SetDiskSize(v int32) *CreatePrePayOrderShrinkRequest {
	s.DiskSize = &v
	return s
}

func (s *CreatePrePayOrderShrinkRequest) SetDiskType(v string) *CreatePrePayOrderShrinkRequest {
	s.DiskType = &v
	return s
}

func (s *CreatePrePayOrderShrinkRequest) SetDuration(v int32) *CreatePrePayOrderShrinkRequest {
	s.Duration = &v
	return s
}

func (s *CreatePrePayOrderShrinkRequest) SetEipMax(v int32) *CreatePrePayOrderShrinkRequest {
	s.EipMax = &v
	return s
}

func (s *CreatePrePayOrderShrinkRequest) SetIoMax(v int32) *CreatePrePayOrderShrinkRequest {
	s.IoMax = &v
	return s
}

func (s *CreatePrePayOrderShrinkRequest) SetIoMaxSpec(v string) *CreatePrePayOrderShrinkRequest {
	s.IoMaxSpec = &v
	return s
}

func (s *CreatePrePayOrderShrinkRequest) SetPaidType(v int32) *CreatePrePayOrderShrinkRequest {
	s.PaidType = &v
	return s
}

func (s *CreatePrePayOrderShrinkRequest) SetPartitionNum(v int32) *CreatePrePayOrderShrinkRequest {
	s.PartitionNum = &v
	return s
}

func (s *CreatePrePayOrderShrinkRequest) SetRegionId(v string) *CreatePrePayOrderShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreatePrePayOrderShrinkRequest) SetResourceGroupId(v string) *CreatePrePayOrderShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreatePrePayOrderShrinkRequest) SetSpecType(v string) *CreatePrePayOrderShrinkRequest {
	s.SpecType = &v
	return s
}

func (s *CreatePrePayOrderShrinkRequest) SetTag(v []*CreatePrePayOrderShrinkRequestTag) *CreatePrePayOrderShrinkRequest {
	s.Tag = v
	return s
}

func (s *CreatePrePayOrderShrinkRequest) SetTopicQuota(v int32) *CreatePrePayOrderShrinkRequest {
	s.TopicQuota = &v
	return s
}

func (s *CreatePrePayOrderShrinkRequest) Validate() error {
	return dara.Validate(s)
}

type CreatePrePayOrderShrinkRequestTag struct {
	// The key of tag N.
	//
	// 	- Valid values of N: 1 to 20.
	//
	// 	- If this parameter is left empty, the keys of all tags are matched.
	//
	// 	- The tag key can be up to 128 characters in length and cannot start with acs: or aliyun or contain [http:// or https://.](http://https://。)
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
	// 	- This parameter can be left empty.
	//
	// 	- The tag value can be 1 to 128 characters in length and cannot start with acs: or aliyun or contain [http:// or https://.](http://https://。)
	//
	// example:
	//
	// FinanceJoshua
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreatePrePayOrderShrinkRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreatePrePayOrderShrinkRequestTag) GoString() string {
	return s.String()
}

func (s *CreatePrePayOrderShrinkRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreatePrePayOrderShrinkRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreatePrePayOrderShrinkRequestTag) SetKey(v string) *CreatePrePayOrderShrinkRequestTag {
	s.Key = &v
	return s
}

func (s *CreatePrePayOrderShrinkRequestTag) SetValue(v string) *CreatePrePayOrderShrinkRequestTag {
	s.Value = &v
	return s
}

func (s *CreatePrePayOrderShrinkRequestTag) Validate() error {
	return dara.Validate(s)
}
