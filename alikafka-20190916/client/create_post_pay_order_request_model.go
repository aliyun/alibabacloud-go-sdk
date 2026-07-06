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
	// - **4**: An instance that is accessible from the Internet and a VPC.
	//
	// - **5**: An instance that is accessible only from a VPC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5
	DeployType *int32 `json:"DeployType,omitempty" xml:"DeployType,omitempty"`
	// The disk capacity.
	//
	// For the valid values, see [Billing](https://help.aliyun.com/document_detail/84737.html).
	//
	// > Do not specify this parameter if you create a Serverless instance.
	//
	// example:
	//
	// 500
	DiskSize *int32 `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
	// The disk type. Valid values:
	//
	// - **0**: Ultra disk
	//
	// - **1**: SSD
	//
	// > Do not specify this parameter if you create a Serverless instance.
	//
	// example:
	//
	// 0
	DiskType *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	// The Internet traffic.
	//
	// - This parameter is required if you set **DeployType*	- to **4**.
	//
	// - For the valid values, see [Billing](https://help.aliyun.com/document_detail/84737.html).
	//
	// > Do not specify this parameter if you create a Serverless instance.
	//
	// example:
	//
	// 0
	EipMax *int32 `json:"EipMax,omitempty" xml:"EipMax,omitempty"`
	// The peak traffic. This parameter is not recommended.
	//
	// - You must specify this parameter or \\`IoMaxSpec\\`. If you specify both parameters, the value of \\`IoMaxSpec\\` takes precedence. We recommend that you specify only \\`IoMaxSpec\\`.
	//
	// - For the valid values, see [Billing](https://help.aliyun.com/document_detail/84737.html).
	//
	// > Do not specify this parameter if you create a Serverless instance.
	//
	// example:
	//
	// 20
	IoMax *int32 `json:"IoMax,omitempty" xml:"IoMax,omitempty"`
	// The traffic specification. This parameter is recommended.
	//
	// - You must specify this parameter or \\`IoMax\\`. If you specify both parameters, the value of this parameter takes precedence. We recommend that you specify only this parameter.
	//
	// - For the valid values, see [Billing](https://help.aliyun.com/document_detail/84737.html).
	//
	// > Do not specify this parameter if you create a Serverless instance.
	//
	// example:
	//
	// alikafka.hw.2xlarge
	IoMaxSpec *string `json:"IoMaxSpec,omitempty" xml:"IoMaxSpec,omitempty"`
	// The billing method. Valid values:
	//
	// - 1 (default): Pay-as-you-go for a reserved instance.
	//
	// - 3: Pay-as-you-go for a reserved Serverless instance and pay-as-you-go for elastic scaling of a Serverless instance.
	//
	// example:
	//
	// 1
	PaidType *int32 `json:"PaidType,omitempty" xml:"PaidType,omitempty"`
	// The number of partitions. This parameter is recommended.
	//
	// - You must specify this parameter or \\`TopicQuota\\`. We recommend that you specify only this parameter.
	//
	// - If you specify both this parameter and \\`TopicQuota\\`, the system verifies whether the values of the two parameters are equivalent based on the previous topic-based sales model. If the values are not equivalent, the system returns a failure. If the values are equivalent, the purchase is made based on the number of partitions.
	//
	// - For the valid values, see [Billing](https://help.aliyun.com/document_detail/84737.html).
	//
	// > Do not specify this parameter if you create a Serverless instance.
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
	// If you do not set this parameter, the instance is added to the default resource group. You can view the resource group ID in the Resource Group console.
	//
	// example:
	//
	// rg-ac***********7q
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The settings of the Serverless instance. This parameter is required if you create a Serverless instance.
	ServerlessConfig *CreatePostPayOrderRequestServerlessConfig `json:"ServerlessConfig,omitempty" xml:"ServerlessConfig,omitempty" type:"Struct"`
	// The specification type.
	//
	// If you set \\`PaidType\\` to 1 (pay-as-you-go for a reserved instance), valid values are:
	//
	// - normal: Standard Edition (High-write)
	//
	// - professional: Professional Edition (High-write)
	//
	// - professionalForHighRead: Professional Edition (High-read)
	//
	// If you set \\`PaidType\\` to 3 (pay-as-you-go for a reserved Serverless instance and pay-as-you-go for elastic scaling of a Serverless instance), valid values are:
	//
	// - basic: Serverless Basic Edition
	//
	// - normal: Serverless Standard Edition
	//
	// - professional: Serverless Professional Edition
	//
	// For more information about these specification types, see [Billing](https://help.aliyun.com/document_detail/84737.html).
	//
	// example:
	//
	// normal
	SpecType *string `json:"SpecType,omitempty" xml:"SpecType,omitempty"`
	// The tags.
	Tag []*CreatePostPayOrderRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The number of topics. This parameter is not recommended.
	//
	// - You must specify this parameter or \\`PartitionNum\\`. We recommend that you specify only \\`PartitionNum\\`.
	//
	// - If you specify both this parameter and \\`PartitionNum\\`, the system verifies whether the values of the two parameters are equivalent based on the previous topic-based sales model. If the values are not equivalent, the system returns a failure. If the values are equivalent, the purchase is made based on the number of partitions.
	//
	// - The default value of this parameter varies based on the traffic specification. You are charged for the extra topics that exceed the default value.
	//
	// - For the valid values, see [Billing](https://help.aliyun.com/document_detail/84737.html).
	//
	// > Do not specify this parameter if you create a Serverless instance.
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
	// The reserved capacity for message publishing. You can specify only an integer for this parameter. The minimum value is 60. This parameter is required if you create a Serverless instance.
	//
	// > The actual upper limit is subject to the inventory in the current region. For more information, see the instance purchase page.
	//
	// example:
	//
	// 60
	ReservedPublishCapacity *int64 `json:"ReservedPublishCapacity,omitempty" xml:"ReservedPublishCapacity,omitempty"`
	// The reserved capacity for message subscription. You can specify only an integer for this parameter. The minimum value is 20. This parameter is required if you create a Serverless instance.
	//
	// > The actual upper limit is subject to the inventory in the current region. For more information, see the instance purchase page.
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
	// The tag key.
	//
	// - N can be an integer from 1 to 20.
	//
	// - If this parameter is empty, all tag keys are matched.
	//
	// - The tag key can be up to 128 characters in length. It cannot start with \\`aliyun\\` or \\`acs:\\` and cannot contain \\`http\\://\\` or \\`https\\://\\`.
	//
	// This parameter is required.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// - N can be an integer from 1 to 20.
	//
	// - This parameter must be empty if the tag key is empty. If this parameter is empty, all tag values are matched.
	//
	// - The tag value can be up to 128 characters in length. It cannot start with \\`aliyun\\` or \\`acs:\\` and cannot contain \\`http\\://\\` or \\`https\\://\\`.
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
