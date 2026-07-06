// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePostPayInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeployType(v int32) *CreatePostPayInstanceRequest
	GetDeployType() *int32
	SetDiskSize(v int32) *CreatePostPayInstanceRequest
	GetDiskSize() *int32
	SetDiskType(v string) *CreatePostPayInstanceRequest
	GetDiskType() *string
	SetEipMax(v int32) *CreatePostPayInstanceRequest
	GetEipMax() *int32
	SetIoMaxSpec(v string) *CreatePostPayInstanceRequest
	GetIoMaxSpec() *string
	SetPaidType(v int32) *CreatePostPayInstanceRequest
	GetPaidType() *int32
	SetPartitionNum(v int32) *CreatePostPayInstanceRequest
	GetPartitionNum() *int32
	SetRegionId(v string) *CreatePostPayInstanceRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreatePostPayInstanceRequest
	GetResourceGroupId() *string
	SetServerlessConfig(v *CreatePostPayInstanceRequestServerlessConfig) *CreatePostPayInstanceRequest
	GetServerlessConfig() *CreatePostPayInstanceRequestServerlessConfig
	SetSpecType(v string) *CreatePostPayInstanceRequest
	GetSpecType() *string
	SetTag(v []*CreatePostPayInstanceRequestTag) *CreatePostPayInstanceRequest
	GetTag() []*CreatePostPayInstanceRequestTag
}

type CreatePostPayInstanceRequest struct {
	// The deployment type. Valid values:
	//
	// - **4**: instance that is accessible over the internet and a VPC
	//
	// - **5**: instance that is accessible only over a VPC
	//
	// This parameter is required.
	//
	// example:
	//
	// 4
	DeployType *int32 `json:"DeployType,omitempty" xml:"DeployType,omitempty"`
	// The disk capacity.
	//
	// For more information about the value range, see [Billing](https://help.aliyun.com/document_detail/84737.html).
	//
	// > This parameter is not required when you create a Serverless instance.
	//
	// example:
	//
	// 1500
	DiskSize *int32 `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
	// The disk type. Valid values:
	//
	// - **0**: ultra disk
	//
	// - **1**: SSD
	//
	// > This parameter is not required when you create a Serverless instance.
	//
	// example:
	//
	// 0
	DiskType *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	// The Internet traffic.
	//
	// - This parameter is required if you set **DeployType*	- to **4**.
	//
	// - For more information about the value range, see [Billing](https://help.aliyun.com/document_detail/84737.html).
	//
	// > This parameter is not required when you create a Serverless instance.
	//
	// example:
	//
	// 3
	EipMax *int32 `json:"EipMax,omitempty" xml:"EipMax,omitempty"`
	// The traffic specification.
	//
	// - For more information about the value range, see [Billing](https://help.aliyun.com/document_detail/84737.html).
	//
	// > This parameter is not required when you create a Serverless instance.
	//
	// example:
	//
	// alikafka.hw.2xlarge
	IoMaxSpec *string `json:"IoMaxSpec,omitempty" xml:"IoMaxSpec,omitempty"`
	// The billing method. Valid values:
	//
	// - 1 (default): pay-as-you-go for reserved instances.
	//
	// - 3: pay-as-you-go for reserved capacity and elastic scaling of Serverless instances.
	//
	// example:
	//
	// 0
	PaidType *int32 `json:"PaidType,omitempty" xml:"PaidType,omitempty"`
	// The number of partitions.
	//
	// - For more information about the value range, see [Billing](https://help.aliyun.com/document_detail/84737.html).
	//
	// > This parameter is not required if the instance is a Serverless instance.
	//
	// example:
	//
	// 100
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
	// If you do not specify this parameter, the instance is added to the default resource group. You can view the resource group ID in the Resource Group console.
	//
	// example:
	//
	// rg-ac***********7q
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The settings of the Serverless instance. This parameter is required when you create a Serverless instance.
	ServerlessConfig *CreatePostPayInstanceRequestServerlessConfig `json:"ServerlessConfig,omitempty" xml:"ServerlessConfig,omitempty" type:"Struct"`
	// The edition of the instance.
	//
	// If you set the PaidType parameter to 1 (pay-as-you-go for reserved instances), valid values are:
	//
	// - normal: Standard Edition (High-write)
	//
	// - professional: Professional Edition (High-write)
	//
	// - professionalForHighRead: Professional Edition (High-read)
	//
	// If you set the PaidType parameter to 3 (pay-as-you-go for reserved capacity and elastic scaling of Serverless instances), valid values are:
	//
	// - basic: Serverless Basic Edition
	//
	// - normal: Serverless Standard Edition
	//
	// - professional: Serverless Professional Edition
	//
	// For more information about these instance editions, see [Billing](https://help.aliyun.com/document_detail/84737.html).
	//
	// example:
	//
	// professional
	SpecType *string `json:"SpecType,omitempty" xml:"SpecType,omitempty"`
	// The tags.
	Tag []*CreatePostPayInstanceRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s CreatePostPayInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePostPayInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreatePostPayInstanceRequest) GetDeployType() *int32 {
	return s.DeployType
}

func (s *CreatePostPayInstanceRequest) GetDiskSize() *int32 {
	return s.DiskSize
}

func (s *CreatePostPayInstanceRequest) GetDiskType() *string {
	return s.DiskType
}

func (s *CreatePostPayInstanceRequest) GetEipMax() *int32 {
	return s.EipMax
}

func (s *CreatePostPayInstanceRequest) GetIoMaxSpec() *string {
	return s.IoMaxSpec
}

func (s *CreatePostPayInstanceRequest) GetPaidType() *int32 {
	return s.PaidType
}

func (s *CreatePostPayInstanceRequest) GetPartitionNum() *int32 {
	return s.PartitionNum
}

func (s *CreatePostPayInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreatePostPayInstanceRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreatePostPayInstanceRequest) GetServerlessConfig() *CreatePostPayInstanceRequestServerlessConfig {
	return s.ServerlessConfig
}

func (s *CreatePostPayInstanceRequest) GetSpecType() *string {
	return s.SpecType
}

func (s *CreatePostPayInstanceRequest) GetTag() []*CreatePostPayInstanceRequestTag {
	return s.Tag
}

func (s *CreatePostPayInstanceRequest) SetDeployType(v int32) *CreatePostPayInstanceRequest {
	s.DeployType = &v
	return s
}

func (s *CreatePostPayInstanceRequest) SetDiskSize(v int32) *CreatePostPayInstanceRequest {
	s.DiskSize = &v
	return s
}

func (s *CreatePostPayInstanceRequest) SetDiskType(v string) *CreatePostPayInstanceRequest {
	s.DiskType = &v
	return s
}

func (s *CreatePostPayInstanceRequest) SetEipMax(v int32) *CreatePostPayInstanceRequest {
	s.EipMax = &v
	return s
}

func (s *CreatePostPayInstanceRequest) SetIoMaxSpec(v string) *CreatePostPayInstanceRequest {
	s.IoMaxSpec = &v
	return s
}

func (s *CreatePostPayInstanceRequest) SetPaidType(v int32) *CreatePostPayInstanceRequest {
	s.PaidType = &v
	return s
}

func (s *CreatePostPayInstanceRequest) SetPartitionNum(v int32) *CreatePostPayInstanceRequest {
	s.PartitionNum = &v
	return s
}

func (s *CreatePostPayInstanceRequest) SetRegionId(v string) *CreatePostPayInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *CreatePostPayInstanceRequest) SetResourceGroupId(v string) *CreatePostPayInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreatePostPayInstanceRequest) SetServerlessConfig(v *CreatePostPayInstanceRequestServerlessConfig) *CreatePostPayInstanceRequest {
	s.ServerlessConfig = v
	return s
}

func (s *CreatePostPayInstanceRequest) SetSpecType(v string) *CreatePostPayInstanceRequest {
	s.SpecType = &v
	return s
}

func (s *CreatePostPayInstanceRequest) SetTag(v []*CreatePostPayInstanceRequestTag) *CreatePostPayInstanceRequest {
	s.Tag = v
	return s
}

func (s *CreatePostPayInstanceRequest) Validate() error {
	if s.ServerlessConfig != nil {
		if err := s.ServerlessConfig.Validate(); err != nil {
			return err
		}
	}
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

type CreatePostPayInstanceRequestServerlessConfig struct {
	// The reserved publish traffic. The value must be an integer. Minimum value: 60. This parameter is required for Serverless instances.
	//
	// > The actual upper limit is subject to the inventory in the current region. For more information, see the value range on the buy page.
	//
	// example:
	//
	// 60
	ReservedPublishCapacity *int64 `json:"ReservedPublishCapacity,omitempty" xml:"ReservedPublishCapacity,omitempty"`
	// The reserved subscribe traffic. The value must be an integer. Minimum value: 20. This parameter is required for Serverless instances.
	//
	// > The actual upper limit is subject to the inventory in the current region. For more information, see the value range on the buy page.
	//
	// example:
	//
	// 20
	ReservedSubscribeCapacity *int64 `json:"ReservedSubscribeCapacity,omitempty" xml:"ReservedSubscribeCapacity,omitempty"`
}

func (s CreatePostPayInstanceRequestServerlessConfig) String() string {
	return dara.Prettify(s)
}

func (s CreatePostPayInstanceRequestServerlessConfig) GoString() string {
	return s.String()
}

func (s *CreatePostPayInstanceRequestServerlessConfig) GetReservedPublishCapacity() *int64 {
	return s.ReservedPublishCapacity
}

func (s *CreatePostPayInstanceRequestServerlessConfig) GetReservedSubscribeCapacity() *int64 {
	return s.ReservedSubscribeCapacity
}

func (s *CreatePostPayInstanceRequestServerlessConfig) SetReservedPublishCapacity(v int64) *CreatePostPayInstanceRequestServerlessConfig {
	s.ReservedPublishCapacity = &v
	return s
}

func (s *CreatePostPayInstanceRequestServerlessConfig) SetReservedSubscribeCapacity(v int64) *CreatePostPayInstanceRequestServerlessConfig {
	s.ReservedSubscribeCapacity = &v
	return s
}

func (s *CreatePostPayInstanceRequestServerlessConfig) Validate() error {
	return dara.Validate(s)
}

type CreatePostPayInstanceRequestTag struct {
	// The tag key of the resource.
	//
	// - The value of N can be from 1 to 20.
	//
	// - If this parameter is left empty, all tag keys are matched.
	//
	// - The tag key can be up to 128 characters in length. It cannot start with aliyun or acs: and cannot contain http\\:// or https\\://.
	//
	// This parameter is required.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the resource.
	//
	// - The value of N can be from 1 to 20.
	//
	// - If the tag key is empty, this parameter must also be empty. If this parameter is empty, all tag values are matched.
	//
	// - The tag value can be up to 128 characters in length. It cannot start with aliyun or acs: and cannot contain http\\:// or https\\://.
	//
	// example:
	//
	// test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreatePostPayInstanceRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreatePostPayInstanceRequestTag) GoString() string {
	return s.String()
}

func (s *CreatePostPayInstanceRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreatePostPayInstanceRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreatePostPayInstanceRequestTag) SetKey(v string) *CreatePostPayInstanceRequestTag {
	s.Key = &v
	return s
}

func (s *CreatePostPayInstanceRequestTag) SetValue(v string) *CreatePostPayInstanceRequestTag {
	s.Value = &v
	return s
}

func (s *CreatePostPayInstanceRequestTag) Validate() error {
	return dara.Validate(s)
}
