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
	ServerlessConfigShrink *string `json:"ServerlessConfig,omitempty" xml:"ServerlessConfig,omitempty"`
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
	Tag []*CreatePostPayOrderShrinkRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
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

type CreatePostPayOrderShrinkRequestTag struct {
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
