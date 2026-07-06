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
	// The configuration of Confluent components.
	//
	// > This parameter is required when you create a Confluent series instance.
	ConfluentConfigShrink *string `json:"ConfluentConfig,omitempty" xml:"ConfluentConfig,omitempty"`
	// The deployment type. Valid values:
	//
	// - **4**: Internet/VPC instance
	//
	// - **5**: VPC instance
	//
	// > If you are creating a Confluent series instance, you cannot select the deployment type. You can only set the value to 5. After the purchase, you can adjust whether each component is open to the Internet.
	//
	// example:
	//
	// 5
	DeployType *int32 `json:"DeployType,omitempty" xml:"DeployType,omitempty"`
	// The disk capacity. Unit: GB.
	//
	// For the valid values, see [Billing](https://help.aliyun.com/document_detail/84737.html).
	//
	// > If you are creating a Confluent series instance, you do not need to pass this parameter.
	//
	// example:
	//
	// 500
	DiskSize *int32 `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
	// The disk type. Valid values:
	//
	// - **0**: ultra disk
	//
	// - **1**: SSD
	//
	// > If you are creating a Confluent series instance, you do not need to pass this parameter.
	//
	// example:
	//
	// 0
	DiskType *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	// The subscription duration. Unit: month. Default value: 1. Valid values:
	//
	// - **Confluent instances: 1 or 12**
	//
	// - **Kafka instances: 1**
	//
	// example:
	//
	// 1
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The Internet traffic.
	//
	// - If **DeployType*	- is set to **4**, you must specify this parameter.
	//
	// - For the valid values, see [pay-as-you-go](https://help.aliyun.com/document_detail/72142.html).
	//
	// > If you are creating a Confluent series instance, you do not need to pass this parameter.
	//
	// example:
	//
	// 0
	EipMax *int32 `json:"EipMax,omitempty" xml:"EipMax,omitempty"`
	// The traffic peak (not recommended).
	//
	// - You must specify either **IoMax*	- or **IoMaxSpec**. If you specify both parameters, **IoMaxSpec*	- takes precedence. We recommend that you specify only **IoMaxSpec**.
	//
	// - For the valid values, see [Billing](https://help.aliyun.com/document_detail/84737.html).
	//
	// > If you are creating a Confluent series instance, you do not need to pass this parameter.
	//
	// example:
	//
	// 20
	IoMax *int32 `json:"IoMax,omitempty" xml:"IoMax,omitempty"`
	// The traffic specification (recommended).
	//
	// - You must specify either **IoMax*	- or **IoMaxSpec**. If you specify both parameters, **IoMaxSpec*	- takes precedence. We recommend that you specify only **IoMaxSpec**.
	//
	// - For the valid values, see [Billing](https://help.aliyun.com/document_detail/84737.html).
	//
	// > If you are creating a Confluent series instance, you do not need to pass this parameter.
	//
	// example:
	//
	// alikafka.hw.2xlarge
	IoMaxSpec *string `json:"IoMaxSpec,omitempty" xml:"IoMaxSpec,omitempty"`
	// The billing method. Valid values:
	//
	// - **0**: subscription
	//
	// - **4**: Confluent series subscription
	//
	// example:
	//
	// 0
	PaidType *int32 `json:"PaidType,omitempty" xml:"PaidType,omitempty"`
	// The number of partitions (recommended).
	//
	// - You must specify either the number of partitions or the topic specification. We recommend that you specify only the number of partitions.
	//
	// - If you specify both the number of partitions and the topic specification, the system verifies whether the number of partitions is equivalent to the topic specification based on the old topic sales model. If they are not equivalent, the system returns a failure. If they are equivalent, the system makes the purchase based on the number of partitions.
	//
	// - For the valid values, see [Billing](https://help.aliyun.com/document_detail/84737.html).
	//
	// > If you are creating a Confluent series instance, you do not need to pass this parameter.
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
	// The resource group ID.
	//
	// If you do not specify this parameter, the instance is added to the default resource group. You can view the resource group ID in the Resource Group console.
	//
	// example:
	//
	// rg-ac***********7q
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The specification type.
	//
	// Valid values for Kafka instances:
	//
	// - **normal**: Standard Edition (high write)
	//
	// - **professional**: Professional Edition (high write)
	//
	// - **professionalForHighRead**: Professional Edition (high read)
	//
	// Valid values for Confluent instances:
	//
	// - **professional**: Professional Edition
	//
	// - **enterprise**: Enterprise Edition
	//
	// For more information, see [Billing](https://help.aliyun.com/document_detail/84737.html).
	//
	// example:
	//
	// normal
	SpecType *string `json:"SpecType,omitempty" xml:"SpecType,omitempty"`
	// The list of tags.
	Tag []*CreatePrePayOrderShrinkRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The number of topics (not recommended).
	//
	// - You must specify either the number of partitions or the topic specification. We recommend that you specify only the number of partitions.
	//
	// - If you specify both the number of partitions and the topic specification, the system verifies whether the number of partitions is equivalent to the topic specification based on the old topic sales model. If they are not equivalent, the system returns a failure. If they are equivalent, the system makes the purchase based on the number of partitions.
	//
	// - The default value varies based on the traffic specification. Additional fees are charged if the value exceeds the default value.
	//
	// - For the valid values, see [Billing](https://help.aliyun.com/document_detail/84737.html).
	//
	// > If you are creating a Confluent series instance, you do not need to pass this parameter.
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
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreatePrePayOrderShrinkRequestTag struct {
	// The tag key of the resource.
	//
	// - N ranges from 1 to 20.
	//
	// - If this parameter is empty, all tag keys are matched.
	//
	// - The tag key can be up to 128 characters in length and cannot start with aliyun or acs:. It cannot contain http\\:// or https\\://.
	//
	// This parameter is required.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the resource.
	//
	// - N ranges from 1 to 20.
	//
	// - This parameter can be empty.
	//
	// - The tag value can be up to 128 characters in length and cannot start with aliyun or acs:. It cannot contain http\\:// or https\\://.
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
