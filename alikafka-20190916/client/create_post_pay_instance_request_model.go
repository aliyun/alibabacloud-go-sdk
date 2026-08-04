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
	// - **4**: Internet- and VPC-connected instance
	//
	// - **5**: VPC-connected instance
	//
	// This parameter is required.
	//
	// example:
	//
	// 4
	DeployType *int32 `json:"DeployType,omitempty" xml:"DeployType,omitempty"`
	// The disk capacity.
	//
	// For the value range, see [Billing](https://help.aliyun.com/document_detail/84737.html).
	//
	// > This parameter is not required if you create a serverless instance.
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
	// > This parameter is not required if you create a serverless instance.
	//
	// example:
	//
	// 0
	DiskType *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	// The Internet traffic.
	//
	// - This parameter is required if **DeployType*	- is set to **4**.
	//
	// - For the value range, see [Billing](https://help.aliyun.com/document_detail/84737.html).
	//
	// > This parameter is not required if you create a serverless instance.
	//
	// example:
	//
	// 3
	EipMax *int32 `json:"EipMax,omitempty" xml:"EipMax,omitempty"`
	// The traffic specification.
	//
	// - For the value range, see [Billing](https://help.aliyun.com/document_detail/84737.html).
	//
	// > This parameter is not required if you create a serverless instance.
	//
	// example:
	//
	// alikafka.hw.2xlarge
	IoMaxSpec *string `json:"IoMaxSpec,omitempty" xml:"IoMaxSpec,omitempty"`
	// The billing type. Valid values:
	//
	// - 1 (default): pay-as-you-go for reserved instances.
	//
	// - 3: pay-as-you-go for serverless reserved specifications + pay-as-you-go for serverless elastic scaling.
	//
	// example:
	//
	// 0
	PaidType *int32 `json:"PaidType,omitempty" xml:"PaidType,omitempty"`
	// The number of partitions to purchase.
	//
	// 	- For the value range, see [Billing](https://help.aliyun.com/document_detail/84737.html).
	//
	// > This parameter is not required if the instance is a serverless instance.
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
	// The resource group ID.
	//
	// If this parameter is not specified, the instance is placed in the default resource group. You can view the resource group ID in the Resource Management console.
	//
	// example:
	//
	// rg-ac***********7q
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The settings for the serverless instance. This parameter is required when you create a serverless instance.
	ServerlessConfig *CreatePostPayInstanceRequestServerlessConfig `json:"ServerlessConfig,omitempty" xml:"ServerlessConfig,omitempty" type:"Struct"`
	// The specification type.
	//
	// Valid values when PaidType is set to 1 (pay-as-you-go for reserved instances):
	//
	// - normal: Standard Edition (shared throughput)
	//
	// - professional: Professional Edition (shared throughput)
	//
	// - professionalForHighRead: Professional Edition (shared throughput for high read)
	//
	// Valid values when PaidType is set to 3 (pay-as-you-go for serverless reserved specifications + pay-as-you-go for serverless elastic scaling):
	//
	// - basic: Serverless Basic Edition
	//
	// - normal: Serverless Standard Edition
	//
	// - professional: Serverless Professional Edition
	//
	// For more information about the specification types, see [Billing](https://help.aliyun.com/document_detail/84737.html).
	//
	// example:
	//
	// professional
	SpecType *string `json:"SpecType,omitempty" xml:"SpecType,omitempty"`
	// The tag list.
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
	// The reserved publish traffic specification value. Only integers are supported. The minimum value is 60. This parameter is required for serverless instances.
	//
	//
	// > The actual upper limit is subject to the inventory in the current region. Refer to the purchase page for the available range.
	//
	// example:
	//
	// 60
	ReservedPublishCapacity *int64 `json:"ReservedPublishCapacity,omitempty" xml:"ReservedPublishCapacity,omitempty"`
	// The reserved subscribe traffic specification value. Only integers are supported. The minimum value is 20. This parameter is required for serverless instances.
	//
	// > The actual upper limit is subject to the inventory in the current region. Refer to the purchase page for the available range.
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
	// - N ranges from 1 to 20.
	//
	// - If this parameter is left empty, all tag keys are matched.
	//
	// - The tag key can be up to 128 characters in length and cannot start with aliyun or acs:. It cannot contain http:// or https://.
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
	// - If the tag key is left empty, this parameter must also be left empty. If this parameter is left empty, all tag values are matched.
	//
	// - The tag value can be up to 128 characters in length and cannot start with aliyun or acs:. It cannot contain http:// or https://.
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
