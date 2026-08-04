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
	// 5
	DeployType *int32 `json:"DeployType,omitempty" xml:"DeployType,omitempty"`
	// The disk capacity.
	//
	// For the value range, see [Billing overview](https://help.aliyun.com/document_detail/84737.html).
	//
	// > If you create a serverless instance, you do not need to set this parameter.
	//
	// example:
	//
	// 500
	DiskSize *int32 `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
	// The disk type. Valid values:
	//
	// - **0**: premium cloud disk
	//
	// - **1**: SSD
	//
	// > If you create a serverless instance, you do not need to set this parameter.
	//
	// example:
	//
	// 0
	DiskType *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	// The Internet traffic.
	//
	// - If **DeployType*	- is set to **4**, this parameter is required.
	//
	// - For the value range, see [Billing overview](https://help.aliyun.com/document_detail/84737.html).
	//
	// > If you create a serverless instance, you do not need to set this parameter.
	//
	// example:
	//
	// 0
	EipMax *int32 `json:"EipMax,omitempty" xml:"EipMax,omitempty"`
	// The maximum traffic (not recommended).
	//
	// - You must specify one of IoMax and IoMaxSpec. If both parameters are specified, the value of IoMaxSpec takes precedence. Specify only IoMaxSpec.
	//
	// - For the value range, see [Billing overview](https://help.aliyun.com/document_detail/84737.html).
	//
	// > If you create a serverless instance, you do not need to set this parameter.
	//
	// example:
	//
	// 20
	IoMax *int32 `json:"IoMax,omitempty" xml:"IoMax,omitempty"`
	// The traffic specification (recommended).
	//
	// - You must specify one of IoMax and IoMaxSpec. If both parameters are specified, the value of IoMaxSpec takes precedence. Specify only IoMaxSpec.
	//
	// - For the value range, see [Billing overview](https://help.aliyun.com/document_detail/84737.html).
	//
	// > If you create a serverless instance, you do not need to set this parameter.
	//
	// example:
	//
	// alikafka.hw.2xlarge
	IoMaxSpec *string `json:"IoMaxSpec,omitempty" xml:"IoMaxSpec,omitempty"`
	// The billing type. Valid values:
	//
	// - 1 (default): reserved instance with pay-as-you-go billing.
	//
	// - 3: serverless instance with reserved specification pay-as-you-go billing + serverless elastic scaling pay-as-you-go billing.
	//
	// example:
	//
	// 1
	PaidType *int32 `json:"PaidType,omitempty" xml:"PaidType,omitempty"`
	// The number of partitions (recommended).
	//
	// 	- You must specify one of PartitionNum and TopicQuota. Specify only PartitionNum.
	//
	// 	- If both PartitionNum and TopicQuota are specified, the system verifies whether the values are equivalent based on the legacy topic sales model. If the values are not equivalent, the request fails. If the values are equivalent, the purchase is made based on the number of partitions.
	//
	// 	- For the value range, see [Billing overview](https://help.aliyun.com/document_detail/84737.html).
	//
	// > If you create a serverless instance, you do not need to set this parameter.
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
	// If this parameter is not specified, the instance is placed in the default resource group. You can view the resource group ID in the Resource Management console.
	//
	// example:
	//
	// rg-ac***********7q
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The settings of the serverless instance. This parameter is required when you create a serverless instance.
	ServerlessConfig *CreatePostPayOrderRequestServerlessConfig `json:"ServerlessConfig,omitempty" xml:"ServerlessConfig,omitempty" type:"Struct"`
	// The specification type.
	//
	// Valid values when PaidType is set to 1 (reserved instance with pay-as-you-go billing):
	//
	// - normal: Standard Edition (shared throughput for writes)
	//
	// - professional: Professional Edition (shared throughput for writes)
	//
	// - professionalForHighRead: Professional Edition (shared throughput for reads)
	//
	// Valid values when PaidType is set to 3 (serverless instance with reserved specification pay-as-you-go billing + serverless elastic scaling pay-as-you-go billing):
	//
	// - basic: Serverless Basic Edition
	//
	// - normal: Serverless Standard Edition
	//
	// - professional: Serverless Professional Edition
	//
	// For more information about these specification types, see [Billing overview](https://help.aliyun.com/document_detail/84737.html).
	//
	// example:
	//
	// normal
	SpecType *string `json:"SpecType,omitempty" xml:"SpecType,omitempty"`
	// The list of tags.
	Tag []*CreatePostPayOrderRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The number of topics (not recommended).
	//
	// - You must specify one of PartitionNum and TopicQuota. Specify only PartitionNum.
	//
	// - If both PartitionNum and TopicQuota are specified, the system verifies whether the values are equivalent based on the legacy topic sales model. If the values are not equivalent, the request fails. If the values are equivalent, the purchase is made based on the number of partitions.
	//
	// - The default value varies based on the traffic specification. If the value exceeds the default value, additional fees are charged.
	//
	// - For the value range, see [Billing overview](https://help.aliyun.com/document_detail/84737.html).
	//
	// > If you create a serverless instance, you do not need to set this parameter.
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

type CreatePostPayOrderRequestServerlessConfig struct {
	// The reserved publish traffic specification value. Only integers are supported. The minimum value is 60. This parameter is required for serverless instances.
	//
	//
	// > The actual upper limit depends on the inventory in the current region. Refer to the purchase page for the available range.
	//
	// example:
	//
	// 60
	ReservedPublishCapacity *int64 `json:"ReservedPublishCapacity,omitempty" xml:"ReservedPublishCapacity,omitempty"`
	// The reserved subscribe traffic specification value. Only integers are supported. The minimum value is 20. This parameter is required for serverless instances.
	//
	// > The actual upper limit depends on the inventory in the current region. Refer to the purchase page for the available range.
	//
	// example:
	//
	// 60
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
