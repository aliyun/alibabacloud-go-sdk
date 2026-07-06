// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradePrePayOrderShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfluentConfigShrink(v string) *UpgradePrePayOrderShrinkRequest
	GetConfluentConfigShrink() *string
	SetDiskSize(v int32) *UpgradePrePayOrderShrinkRequest
	GetDiskSize() *int32
	SetEipMax(v int32) *UpgradePrePayOrderShrinkRequest
	GetEipMax() *int32
	SetEipModel(v bool) *UpgradePrePayOrderShrinkRequest
	GetEipModel() *bool
	SetInstanceId(v string) *UpgradePrePayOrderShrinkRequest
	GetInstanceId() *string
	SetIoMax(v int32) *UpgradePrePayOrderShrinkRequest
	GetIoMax() *int32
	SetIoMaxSpec(v string) *UpgradePrePayOrderShrinkRequest
	GetIoMaxSpec() *string
	SetPaidType(v int32) *UpgradePrePayOrderShrinkRequest
	GetPaidType() *int32
	SetPartitionNum(v int32) *UpgradePrePayOrderShrinkRequest
	GetPartitionNum() *int32
	SetRegionId(v string) *UpgradePrePayOrderShrinkRequest
	GetRegionId() *string
	SetSpecType(v string) *UpgradePrePayOrderShrinkRequest
	GetSpecType() *string
	SetTopicQuota(v int32) *UpgradePrePayOrderShrinkRequest
	GetTopicQuota() *int32
}

type UpgradePrePayOrderShrinkRequest struct {
	// Configurations for the Confluent components.
	ConfluentConfigShrink *string `json:"ConfluentConfig,omitempty" xml:"ConfluentConfig,omitempty"`
	// The disk capacity.
	//
	// - The specified disk capacity must be greater than or equal to the current disk capacity of the instance.
	//
	// - For the valid values, see [Billing](https://help.aliyun.com/document_detail/84737.html).
	//
	// > This parameter is required for subscription instances but not for Confluent-series instances.
	//
	// example:
	//
	// 900
	DiskSize *int32 `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
	// The maximum Internet traffic bandwidth.
	//
	// - The specified Internet traffic bandwidth must be greater than or equal to the current Internet traffic bandwidth of the instance.
	//
	// - For the valid values, see [Billing](https://help.aliyun.com/document_detail/84737.html).
	//
	// > 	- If **EipModel*	- is set to **true**, **EipMax*	- must be greater than 0.
	//
	// >
	//
	// > 	- If **EipModel*	- is set to **false**, **EipMax*	- must be set to **0**.
	//
	// example:
	//
	// 3
	EipMax *int32 `json:"EipMax,omitempty" xml:"EipMax,omitempty"`
	// Specifies whether to enable Internet access. Valid values:
	//
	// - `true`: enables Internet access.
	//
	// - `false`: disables Internet access.
	//
	// > This parameter is required for subscription instances but not for Confluent-series instances.
	//
	// example:
	//
	// true
	EipModel *bool `json:"EipModel,omitempty" xml:"EipModel,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// alikafka_post-cn-mp919o4v****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The traffic peak (not recommended).
	//
	// - The specified traffic peak must be greater than or equal to the current traffic peak of the instance.
	//
	// - You must specify either this parameter or `IoMaxSpec`. If you specify both, `IoMaxSpec` takes precedence. We recommend that you specify only `IoMaxSpec`.
	//
	// - For the valid values, see [Billing](https://help.aliyun.com/document_detail/84737.html).
	//
	// example:
	//
	// 40
	IoMax *int32 `json:"IoMax,omitempty" xml:"IoMax,omitempty"`
	// The traffic specification (recommended).
	//
	// - The specified traffic specification must be greater than or equal to the current traffic specification of the instance.
	//
	// - You must specify either this parameter or `IoMax`. If you specify both, this parameter takes precedence. We recommend that you specify only this parameter.
	//
	// - For the valid values, see [Billing](https://help.aliyun.com/document_detail/84737.html).
	//
	// > This parameter is required for subscription instances but not for Confluent-series instances.
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
	// 4
	PaidType *int32 `json:"PaidType,omitempty" xml:"PaidType,omitempty"`
	// The number of partitions (recommended).
	//
	// - You must specify either this parameter or `TopicQuota`. We recommend that you specify only this parameter.
	//
	// - If you specify both `PartitionNum` and `TopicQuota`, the system checks if their values are equivalent under the previous topic pricing model. A mismatch causes the request to fail. If they match, the system uses `PartitionNum` to process the purchase.
	//
	// - For the valid values, see [Billing](https://help.aliyun.com/document_detail/84737.html).
	//
	// > This parameter is required for subscription instances but not for Confluent-series instances.
	//
	// example:
	//
	// 50
	PartitionNum *int32 `json:"PartitionNum,omitempty" xml:"PartitionNum,omitempty"`
	// The ID of the region where the instance is located.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The specification type.
	//
	// Valid values for ApsaraMQ for Kafka instances:
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
	// You cannot downgrade an instance from Professional Edition to Standard Edition. For more information about these specification types, see [Billing](https://help.aliyun.com/document_detail/84737.html).
	//
	// example:
	//
	// professional
	SpecType *string `json:"SpecType,omitempty" xml:"SpecType,omitempty"`
	// The number of topics (not recommended).
	//
	// - You must specify either this parameter or `PartitionNum`. We recommend that you specify only `PartitionNum`.
	//
	// - If you specify both `TopicQuota` and `PartitionNum`, the system checks if their values are equivalent under the previous topic pricing model. A mismatch causes the request to fail. If they match, the system uses `PartitionNum` to process the purchase.
	//
	// - The default value of this parameter varies based on the traffic specification. You are charged additional fees if the specified value exceeds the default value.
	//
	// - For the valid values, see [Billing](https://help.aliyun.com/document_detail/84737.html).
	//
	// example:
	//
	// 50
	TopicQuota *int32 `json:"TopicQuota,omitempty" xml:"TopicQuota,omitempty"`
}

func (s UpgradePrePayOrderShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpgradePrePayOrderShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpgradePrePayOrderShrinkRequest) GetConfluentConfigShrink() *string {
	return s.ConfluentConfigShrink
}

func (s *UpgradePrePayOrderShrinkRequest) GetDiskSize() *int32 {
	return s.DiskSize
}

func (s *UpgradePrePayOrderShrinkRequest) GetEipMax() *int32 {
	return s.EipMax
}

func (s *UpgradePrePayOrderShrinkRequest) GetEipModel() *bool {
	return s.EipModel
}

func (s *UpgradePrePayOrderShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpgradePrePayOrderShrinkRequest) GetIoMax() *int32 {
	return s.IoMax
}

func (s *UpgradePrePayOrderShrinkRequest) GetIoMaxSpec() *string {
	return s.IoMaxSpec
}

func (s *UpgradePrePayOrderShrinkRequest) GetPaidType() *int32 {
	return s.PaidType
}

func (s *UpgradePrePayOrderShrinkRequest) GetPartitionNum() *int32 {
	return s.PartitionNum
}

func (s *UpgradePrePayOrderShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpgradePrePayOrderShrinkRequest) GetSpecType() *string {
	return s.SpecType
}

func (s *UpgradePrePayOrderShrinkRequest) GetTopicQuota() *int32 {
	return s.TopicQuota
}

func (s *UpgradePrePayOrderShrinkRequest) SetConfluentConfigShrink(v string) *UpgradePrePayOrderShrinkRequest {
	s.ConfluentConfigShrink = &v
	return s
}

func (s *UpgradePrePayOrderShrinkRequest) SetDiskSize(v int32) *UpgradePrePayOrderShrinkRequest {
	s.DiskSize = &v
	return s
}

func (s *UpgradePrePayOrderShrinkRequest) SetEipMax(v int32) *UpgradePrePayOrderShrinkRequest {
	s.EipMax = &v
	return s
}

func (s *UpgradePrePayOrderShrinkRequest) SetEipModel(v bool) *UpgradePrePayOrderShrinkRequest {
	s.EipModel = &v
	return s
}

func (s *UpgradePrePayOrderShrinkRequest) SetInstanceId(v string) *UpgradePrePayOrderShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *UpgradePrePayOrderShrinkRequest) SetIoMax(v int32) *UpgradePrePayOrderShrinkRequest {
	s.IoMax = &v
	return s
}

func (s *UpgradePrePayOrderShrinkRequest) SetIoMaxSpec(v string) *UpgradePrePayOrderShrinkRequest {
	s.IoMaxSpec = &v
	return s
}

func (s *UpgradePrePayOrderShrinkRequest) SetPaidType(v int32) *UpgradePrePayOrderShrinkRequest {
	s.PaidType = &v
	return s
}

func (s *UpgradePrePayOrderShrinkRequest) SetPartitionNum(v int32) *UpgradePrePayOrderShrinkRequest {
	s.PartitionNum = &v
	return s
}

func (s *UpgradePrePayOrderShrinkRequest) SetRegionId(v string) *UpgradePrePayOrderShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *UpgradePrePayOrderShrinkRequest) SetSpecType(v string) *UpgradePrePayOrderShrinkRequest {
	s.SpecType = &v
	return s
}

func (s *UpgradePrePayOrderShrinkRequest) SetTopicQuota(v int32) *UpgradePrePayOrderShrinkRequest {
	s.TopicQuota = &v
	return s
}

func (s *UpgradePrePayOrderShrinkRequest) Validate() error {
	return dara.Validate(s)
}
