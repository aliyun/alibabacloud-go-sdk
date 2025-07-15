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
	ConfluentConfigShrink *string `json:"ConfluentConfig,omitempty" xml:"ConfluentConfig,omitempty"`
	// The size of the disk.
	//
	// 	- The disk size that you specify must be greater than or equal to the current disk size of the instance.
	//
	// 	- For more information about the valid values, see [Billing overview](https://help.aliyun.com/document_detail/84737.html).
	//
	// example:
	//
	// 900
	DiskSize *int32 `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
	// The Internet traffic for the instance.
	//
	// 	- The Internet traffic volume that you specify must be greater than or equal to the current Internet traffic volume of the instance.
	//
	// 	- For more information about the valid values, see [Billing overview](https://help.aliyun.com/document_detail/84737.html).
	//
	// > - If the **EipModel*	- parameter is set to **true**, set the **EipMax*	- parameter to a value that is greater than 0.
	//
	// > - If the **EipModel*	- parameter is set to **false**, set the **EipMax*	- parameter to **0**.
	//
	// example:
	//
	// 3
	EipMax *int32 `json:"EipMax,omitempty" xml:"EipMax,omitempty"`
	// Specifies whether to enable Internet access for the instance. Valid values:
	//
	// 	- true: enables Internet access.
	//
	// 	- false: disables Internet access.
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
	// The maximum traffic for the instance. We recommend that you do not configure this parameter.
	//
	// 	- The maximum traffic volume that you specify must be greater than or equal to the current maximum traffic volume of the instance.
	//
	// 	- You must configure at least one of the IoMax and IoMaxSpec parameters. If you configure both parameters, the value of the IoMaxSpec parameter takes effect. We recommend that you configure only the IoMaxSpec parameter.
	//
	// 	- For more information about the valid values, see [Billing overview](https://help.aliyun.com/document_detail/84737.html).
	//
	// example:
	//
	// 40
	IoMax *int32 `json:"IoMax,omitempty" xml:"IoMax,omitempty"`
	// The traffic specification of the instance. We recommend that you configure this parameter.
	//
	// 	- The traffic specification that you specify must be greater than or equal to the current traffic specification of the instance.
	//
	// 	- You must configure at least one of the IoMax and IoMaxSpec parameters. If you configure both parameters, the value of the IoMaxSpec parameter takes effect. We recommend that you configure only the IoMaxSpec parameter.
	//
	// 	- For more information about the valid values, see [Billing overview](https://help.aliyun.com/document_detail/84737.html).
	//
	// example:
	//
	// alikafka.hw.2xlarge
	IoMaxSpec *string `json:"IoMaxSpec,omitempty" xml:"IoMaxSpec,omitempty"`
	PaidType  *int32  `json:"PaidType,omitempty" xml:"PaidType,omitempty"`
	// The number of partitions. We recommend that you configure this parameter.
	//
	// 	- You must specify at least one of the PartitionNum and TopicQuota parameters. We recommend that you configure only the PartitionNum parameter.
	//
	// 	- If you specify both parameters, the topic-based sales model is used to check whether the PartitionNum value and the TopicQuota value are the same. If they are not the same, a failure response is returned. If they are the same, the order is placed based on the PartitionNum value.
	//
	// 	- For more information about the valid values, see [Billing overview](https://help.aliyun.com/document_detail/84737.html).
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
	// The edition of the instance. Valid values:
	//
	// 	- **normal**: Standard Edition (High Write)
	//
	// 	- **professional**: Professional Edition (High Write)
	//
	// 	- **professionalForHighRead**: Professional Edition (High Read)
	//
	// You cannot downgrade an instance from the Professional Edition to the Standard Edition. For more information about these instance editions, see [Billing overview](https://help.aliyun.com/document_detail/84737.html).
	//
	// example:
	//
	// professional
	SpecType *string `json:"SpecType,omitempty" xml:"SpecType,omitempty"`
	// The number of topics. We recommend that you do not configure this parameter.
	//
	// 	- You must specify at least one of the PartitionNum and TopicQuota parameters. We recommend that you configure only the PartitionNum parameter.
	//
	// 	- If you specify both parameters, the topic-based sales model is used to check whether the PartitionNum value and the TopicQuota value are the same. If they are not the same, a failure response is returned. If they are the same, the order is placed based on the PartitionNum value.
	//
	// 	- The default value of the TopicQuota parameter varies based on the value of the IoMaxSpec parameter. If the number of topics that you consume exceeds the default value, you are charged additional fees.
	//
	// 	- For more information about the valid values, see [Billing overview](https://help.aliyun.com/document_detail/84737.html).
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
