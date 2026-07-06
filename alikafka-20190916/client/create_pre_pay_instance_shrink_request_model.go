// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePrePayInstanceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfluentConfigShrink(v string) *CreatePrePayInstanceShrinkRequest
	GetConfluentConfigShrink() *string
	SetDeployType(v int32) *CreatePrePayInstanceShrinkRequest
	GetDeployType() *int32
	SetDiskSize(v int32) *CreatePrePayInstanceShrinkRequest
	GetDiskSize() *int32
	SetDiskType(v string) *CreatePrePayInstanceShrinkRequest
	GetDiskType() *string
	SetDuration(v int32) *CreatePrePayInstanceShrinkRequest
	GetDuration() *int32
	SetEipMax(v int32) *CreatePrePayInstanceShrinkRequest
	GetEipMax() *int32
	SetIoMaxSpec(v string) *CreatePrePayInstanceShrinkRequest
	GetIoMaxSpec() *string
	SetPaidType(v int32) *CreatePrePayInstanceShrinkRequest
	GetPaidType() *int32
	SetPartitionNum(v int32) *CreatePrePayInstanceShrinkRequest
	GetPartitionNum() *int32
	SetRegionId(v string) *CreatePrePayInstanceShrinkRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreatePrePayInstanceShrinkRequest
	GetResourceGroupId() *string
	SetSpecType(v string) *CreatePrePayInstanceShrinkRequest
	GetSpecType() *string
	SetTag(v []*CreatePrePayInstanceShrinkRequestTag) *CreatePrePayInstanceShrinkRequest
	GetTag() []*CreatePrePayInstanceShrinkRequestTag
}

type CreatePrePayInstanceShrinkRequest struct {
	// The configurations of the Confluent components.
	//
	// > This parameter is required if you create a Confluent instance.
	ConfluentConfigShrink *string `json:"ConfluentConfig,omitempty" xml:"ConfluentConfig,omitempty"`
	// The deployment type. Valid values:
	//
	// - **4**: an instance accessible from the internet and a VPC
	//
	// - **5**: an instance accessible from a VPC only
	//
	// > If you create a Confluent instance, you cannot specify the deployment type and must set this parameter to 5. After the instance is created, you can configure internet access for each component.
	//
	// example:
	//
	// 5
	DeployType *int32 `json:"DeployType,omitempty" xml:"DeployType,omitempty"`
	// The disk capacity, in GB.
	//
	// For the value range, see [Billing](https://help.aliyun.com/document_detail/84737.html).
	//
	// > This parameter is not required if you create a Confluent instance.
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
	// > This parameter is not required if you create a Confluent instance.
	//
	// example:
	//
	// 1
	DiskType *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	// The subscription duration, in months. Default value: 1. Valid values:
	//
	// - Confluent instances: **1*	- and **12**
	//
	// - Kafka instances: **1**
	//
	// example:
	//
	// 1
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The peak internet bandwidth.
	//
	// - This parameter is required if you set **DeployType*	- to **4**.
	//
	// - For the value range, see [pay-as-you-go](https://help.aliyun.com/document_detail/72142.html).
	//
	// > This parameter is not required if you create a Confluent instance.
	//
	// example:
	//
	// 3
	EipMax *int32 `json:"EipMax,omitempty" xml:"EipMax,omitempty"`
	// The I/O specification.
	//
	// - For the value range, see [Billing](https://help.aliyun.com/document_detail/84737.html).
	//
	// > This parameter is not required if you create a Confluent instance.
	//
	// example:
	//
	// alikafka.hw.2xlarge
	IoMaxSpec *string `json:"IoMaxSpec,omitempty" xml:"IoMaxSpec,omitempty"`
	// The billing method. Valid values:
	//
	// - **0**: subscription
	//
	// - **4**: subscription for Confluent instances
	//
	// example:
	//
	// 1
	PaidType *int32 `json:"PaidType,omitempty" xml:"PaidType,omitempty"`
	// The number of partitions.
	//
	// - For the value range, see [Billing](https://help.aliyun.com/document_detail/84737.html).
	//
	// > This parameter is not required if you create a Confluent instance.
	//
	// example:
	//
	// 1000
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
	// If you do not specify this parameter, the instance is placed in the default resource group. You can find the resource group ID in the Resource Group console.
	//
	// example:
	//
	// rg-ac***********7q
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The specification type.
	//
	// Valid values for Kafka instances:
	//
	// - **normal**: Standard Edition (High-write)
	//
	// - **professional**: Professional Edition (High-write)
	//
	// - **professionalForHighRead**: Professional Edition (High-read)
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
	// professional
	SpecType *string `json:"SpecType,omitempty" xml:"SpecType,omitempty"`
	// The tags to attach to the instance. You can specify up to 20 tags.
	Tag []*CreatePrePayInstanceShrinkRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s CreatePrePayInstanceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePrePayInstanceShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreatePrePayInstanceShrinkRequest) GetConfluentConfigShrink() *string {
	return s.ConfluentConfigShrink
}

func (s *CreatePrePayInstanceShrinkRequest) GetDeployType() *int32 {
	return s.DeployType
}

func (s *CreatePrePayInstanceShrinkRequest) GetDiskSize() *int32 {
	return s.DiskSize
}

func (s *CreatePrePayInstanceShrinkRequest) GetDiskType() *string {
	return s.DiskType
}

func (s *CreatePrePayInstanceShrinkRequest) GetDuration() *int32 {
	return s.Duration
}

func (s *CreatePrePayInstanceShrinkRequest) GetEipMax() *int32 {
	return s.EipMax
}

func (s *CreatePrePayInstanceShrinkRequest) GetIoMaxSpec() *string {
	return s.IoMaxSpec
}

func (s *CreatePrePayInstanceShrinkRequest) GetPaidType() *int32 {
	return s.PaidType
}

func (s *CreatePrePayInstanceShrinkRequest) GetPartitionNum() *int32 {
	return s.PartitionNum
}

func (s *CreatePrePayInstanceShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreatePrePayInstanceShrinkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreatePrePayInstanceShrinkRequest) GetSpecType() *string {
	return s.SpecType
}

func (s *CreatePrePayInstanceShrinkRequest) GetTag() []*CreatePrePayInstanceShrinkRequestTag {
	return s.Tag
}

func (s *CreatePrePayInstanceShrinkRequest) SetConfluentConfigShrink(v string) *CreatePrePayInstanceShrinkRequest {
	s.ConfluentConfigShrink = &v
	return s
}

func (s *CreatePrePayInstanceShrinkRequest) SetDeployType(v int32) *CreatePrePayInstanceShrinkRequest {
	s.DeployType = &v
	return s
}

func (s *CreatePrePayInstanceShrinkRequest) SetDiskSize(v int32) *CreatePrePayInstanceShrinkRequest {
	s.DiskSize = &v
	return s
}

func (s *CreatePrePayInstanceShrinkRequest) SetDiskType(v string) *CreatePrePayInstanceShrinkRequest {
	s.DiskType = &v
	return s
}

func (s *CreatePrePayInstanceShrinkRequest) SetDuration(v int32) *CreatePrePayInstanceShrinkRequest {
	s.Duration = &v
	return s
}

func (s *CreatePrePayInstanceShrinkRequest) SetEipMax(v int32) *CreatePrePayInstanceShrinkRequest {
	s.EipMax = &v
	return s
}

func (s *CreatePrePayInstanceShrinkRequest) SetIoMaxSpec(v string) *CreatePrePayInstanceShrinkRequest {
	s.IoMaxSpec = &v
	return s
}

func (s *CreatePrePayInstanceShrinkRequest) SetPaidType(v int32) *CreatePrePayInstanceShrinkRequest {
	s.PaidType = &v
	return s
}

func (s *CreatePrePayInstanceShrinkRequest) SetPartitionNum(v int32) *CreatePrePayInstanceShrinkRequest {
	s.PartitionNum = &v
	return s
}

func (s *CreatePrePayInstanceShrinkRequest) SetRegionId(v string) *CreatePrePayInstanceShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreatePrePayInstanceShrinkRequest) SetResourceGroupId(v string) *CreatePrePayInstanceShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreatePrePayInstanceShrinkRequest) SetSpecType(v string) *CreatePrePayInstanceShrinkRequest {
	s.SpecType = &v
	return s
}

func (s *CreatePrePayInstanceShrinkRequest) SetTag(v []*CreatePrePayInstanceShrinkRequestTag) *CreatePrePayInstanceShrinkRequest {
	s.Tag = v
	return s
}

func (s *CreatePrePayInstanceShrinkRequest) Validate() error {
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

type CreatePrePayInstanceShrinkRequestTag struct {
	// The tag key.
	//
	// -
	//
	// -
	//
	// - The key must be 1 to 128 characters long. It cannot start with aliyun or acs:, nor can it contain http\\:// or https\\://.
	//
	// This parameter is required.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// -
	//
	// -
	//
	// - The value can be 0 to 128 characters long. It cannot start with aliyun or acs:, nor can it contain http\\:// or https\\://.
	//
	// example:
	//
	// test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreatePrePayInstanceShrinkRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreatePrePayInstanceShrinkRequestTag) GoString() string {
	return s.String()
}

func (s *CreatePrePayInstanceShrinkRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreatePrePayInstanceShrinkRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreatePrePayInstanceShrinkRequestTag) SetKey(v string) *CreatePrePayInstanceShrinkRequestTag {
	s.Key = &v
	return s
}

func (s *CreatePrePayInstanceShrinkRequestTag) SetValue(v string) *CreatePrePayInstanceShrinkRequestTag {
	s.Value = &v
	return s
}

func (s *CreatePrePayInstanceShrinkRequestTag) Validate() error {
	return dara.Validate(s)
}
