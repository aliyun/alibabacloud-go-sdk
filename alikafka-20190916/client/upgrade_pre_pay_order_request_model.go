// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradePrePayOrderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfluentConfig(v *UpgradePrePayOrderRequestConfluentConfig) *UpgradePrePayOrderRequest
	GetConfluentConfig() *UpgradePrePayOrderRequestConfluentConfig
	SetDiskSize(v int32) *UpgradePrePayOrderRequest
	GetDiskSize() *int32
	SetEipMax(v int32) *UpgradePrePayOrderRequest
	GetEipMax() *int32
	SetEipModel(v bool) *UpgradePrePayOrderRequest
	GetEipModel() *bool
	SetInstanceId(v string) *UpgradePrePayOrderRequest
	GetInstanceId() *string
	SetIoMax(v int32) *UpgradePrePayOrderRequest
	GetIoMax() *int32
	SetIoMaxSpec(v string) *UpgradePrePayOrderRequest
	GetIoMaxSpec() *string
	SetPaidType(v int32) *UpgradePrePayOrderRequest
	GetPaidType() *int32
	SetPartitionNum(v int32) *UpgradePrePayOrderRequest
	GetPartitionNum() *int32
	SetRegionId(v string) *UpgradePrePayOrderRequest
	GetRegionId() *string
	SetSpecType(v string) *UpgradePrePayOrderRequest
	GetSpecType() *string
	SetTopicQuota(v int32) *UpgradePrePayOrderRequest
	GetTopicQuota() *int32
}

type UpgradePrePayOrderRequest struct {
	ConfluentConfig *UpgradePrePayOrderRequestConfluentConfig `json:"ConfluentConfig,omitempty" xml:"ConfluentConfig,omitempty" type:"Struct"`
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

func (s UpgradePrePayOrderRequest) String() string {
	return dara.Prettify(s)
}

func (s UpgradePrePayOrderRequest) GoString() string {
	return s.String()
}

func (s *UpgradePrePayOrderRequest) GetConfluentConfig() *UpgradePrePayOrderRequestConfluentConfig {
	return s.ConfluentConfig
}

func (s *UpgradePrePayOrderRequest) GetDiskSize() *int32 {
	return s.DiskSize
}

func (s *UpgradePrePayOrderRequest) GetEipMax() *int32 {
	return s.EipMax
}

func (s *UpgradePrePayOrderRequest) GetEipModel() *bool {
	return s.EipModel
}

func (s *UpgradePrePayOrderRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpgradePrePayOrderRequest) GetIoMax() *int32 {
	return s.IoMax
}

func (s *UpgradePrePayOrderRequest) GetIoMaxSpec() *string {
	return s.IoMaxSpec
}

func (s *UpgradePrePayOrderRequest) GetPaidType() *int32 {
	return s.PaidType
}

func (s *UpgradePrePayOrderRequest) GetPartitionNum() *int32 {
	return s.PartitionNum
}

func (s *UpgradePrePayOrderRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpgradePrePayOrderRequest) GetSpecType() *string {
	return s.SpecType
}

func (s *UpgradePrePayOrderRequest) GetTopicQuota() *int32 {
	return s.TopicQuota
}

func (s *UpgradePrePayOrderRequest) SetConfluentConfig(v *UpgradePrePayOrderRequestConfluentConfig) *UpgradePrePayOrderRequest {
	s.ConfluentConfig = v
	return s
}

func (s *UpgradePrePayOrderRequest) SetDiskSize(v int32) *UpgradePrePayOrderRequest {
	s.DiskSize = &v
	return s
}

func (s *UpgradePrePayOrderRequest) SetEipMax(v int32) *UpgradePrePayOrderRequest {
	s.EipMax = &v
	return s
}

func (s *UpgradePrePayOrderRequest) SetEipModel(v bool) *UpgradePrePayOrderRequest {
	s.EipModel = &v
	return s
}

func (s *UpgradePrePayOrderRequest) SetInstanceId(v string) *UpgradePrePayOrderRequest {
	s.InstanceId = &v
	return s
}

func (s *UpgradePrePayOrderRequest) SetIoMax(v int32) *UpgradePrePayOrderRequest {
	s.IoMax = &v
	return s
}

func (s *UpgradePrePayOrderRequest) SetIoMaxSpec(v string) *UpgradePrePayOrderRequest {
	s.IoMaxSpec = &v
	return s
}

func (s *UpgradePrePayOrderRequest) SetPaidType(v int32) *UpgradePrePayOrderRequest {
	s.PaidType = &v
	return s
}

func (s *UpgradePrePayOrderRequest) SetPartitionNum(v int32) *UpgradePrePayOrderRequest {
	s.PartitionNum = &v
	return s
}

func (s *UpgradePrePayOrderRequest) SetRegionId(v string) *UpgradePrePayOrderRequest {
	s.RegionId = &v
	return s
}

func (s *UpgradePrePayOrderRequest) SetSpecType(v string) *UpgradePrePayOrderRequest {
	s.SpecType = &v
	return s
}

func (s *UpgradePrePayOrderRequest) SetTopicQuota(v int32) *UpgradePrePayOrderRequest {
	s.TopicQuota = &v
	return s
}

func (s *UpgradePrePayOrderRequest) Validate() error {
	return dara.Validate(s)
}

type UpgradePrePayOrderRequestConfluentConfig struct {
	ConnectCU             *int32 `json:"ConnectCU,omitempty" xml:"ConnectCU,omitempty"`
	ConnectReplica        *int32 `json:"ConnectReplica,omitempty" xml:"ConnectReplica,omitempty"`
	ControlCenterCU       *int32 `json:"ControlCenterCU,omitempty" xml:"ControlCenterCU,omitempty"`
	ControlCenterReplica  *int32 `json:"ControlCenterReplica,omitempty" xml:"ControlCenterReplica,omitempty"`
	ControlCenterStorage  *int32 `json:"ControlCenterStorage,omitempty" xml:"ControlCenterStorage,omitempty"`
	KafkaCU               *int32 `json:"KafkaCU,omitempty" xml:"KafkaCU,omitempty"`
	KafkaReplica          *int32 `json:"KafkaReplica,omitempty" xml:"KafkaReplica,omitempty"`
	KafkaRestProxyCU      *int32 `json:"KafkaRestProxyCU,omitempty" xml:"KafkaRestProxyCU,omitempty"`
	KafkaRestProxyReplica *int32 `json:"KafkaRestProxyReplica,omitempty" xml:"KafkaRestProxyReplica,omitempty"`
	KafkaStorage          *int32 `json:"KafkaStorage,omitempty" xml:"KafkaStorage,omitempty"`
	KsqlCU                *int32 `json:"KsqlCU,omitempty" xml:"KsqlCU,omitempty"`
	KsqlReplica           *int32 `json:"KsqlReplica,omitempty" xml:"KsqlReplica,omitempty"`
	KsqlStorage           *int32 `json:"KsqlStorage,omitempty" xml:"KsqlStorage,omitempty"`
	SchemaRegistryCU      *int32 `json:"SchemaRegistryCU,omitempty" xml:"SchemaRegistryCU,omitempty"`
	SchemaRegistryReplica *int32 `json:"SchemaRegistryReplica,omitempty" xml:"SchemaRegistryReplica,omitempty"`
	ZooKeeperCU           *int32 `json:"ZooKeeperCU,omitempty" xml:"ZooKeeperCU,omitempty"`
	ZooKeeperReplica      *int32 `json:"ZooKeeperReplica,omitempty" xml:"ZooKeeperReplica,omitempty"`
	ZooKeeperStorage      *int32 `json:"ZooKeeperStorage,omitempty" xml:"ZooKeeperStorage,omitempty"`
}

func (s UpgradePrePayOrderRequestConfluentConfig) String() string {
	return dara.Prettify(s)
}

func (s UpgradePrePayOrderRequestConfluentConfig) GoString() string {
	return s.String()
}

func (s *UpgradePrePayOrderRequestConfluentConfig) GetConnectCU() *int32 {
	return s.ConnectCU
}

func (s *UpgradePrePayOrderRequestConfluentConfig) GetConnectReplica() *int32 {
	return s.ConnectReplica
}

func (s *UpgradePrePayOrderRequestConfluentConfig) GetControlCenterCU() *int32 {
	return s.ControlCenterCU
}

func (s *UpgradePrePayOrderRequestConfluentConfig) GetControlCenterReplica() *int32 {
	return s.ControlCenterReplica
}

func (s *UpgradePrePayOrderRequestConfluentConfig) GetControlCenterStorage() *int32 {
	return s.ControlCenterStorage
}

func (s *UpgradePrePayOrderRequestConfluentConfig) GetKafkaCU() *int32 {
	return s.KafkaCU
}

func (s *UpgradePrePayOrderRequestConfluentConfig) GetKafkaReplica() *int32 {
	return s.KafkaReplica
}

func (s *UpgradePrePayOrderRequestConfluentConfig) GetKafkaRestProxyCU() *int32 {
	return s.KafkaRestProxyCU
}

func (s *UpgradePrePayOrderRequestConfluentConfig) GetKafkaRestProxyReplica() *int32 {
	return s.KafkaRestProxyReplica
}

func (s *UpgradePrePayOrderRequestConfluentConfig) GetKafkaStorage() *int32 {
	return s.KafkaStorage
}

func (s *UpgradePrePayOrderRequestConfluentConfig) GetKsqlCU() *int32 {
	return s.KsqlCU
}

func (s *UpgradePrePayOrderRequestConfluentConfig) GetKsqlReplica() *int32 {
	return s.KsqlReplica
}

func (s *UpgradePrePayOrderRequestConfluentConfig) GetKsqlStorage() *int32 {
	return s.KsqlStorage
}

func (s *UpgradePrePayOrderRequestConfluentConfig) GetSchemaRegistryCU() *int32 {
	return s.SchemaRegistryCU
}

func (s *UpgradePrePayOrderRequestConfluentConfig) GetSchemaRegistryReplica() *int32 {
	return s.SchemaRegistryReplica
}

func (s *UpgradePrePayOrderRequestConfluentConfig) GetZooKeeperCU() *int32 {
	return s.ZooKeeperCU
}

func (s *UpgradePrePayOrderRequestConfluentConfig) GetZooKeeperReplica() *int32 {
	return s.ZooKeeperReplica
}

func (s *UpgradePrePayOrderRequestConfluentConfig) GetZooKeeperStorage() *int32 {
	return s.ZooKeeperStorage
}

func (s *UpgradePrePayOrderRequestConfluentConfig) SetConnectCU(v int32) *UpgradePrePayOrderRequestConfluentConfig {
	s.ConnectCU = &v
	return s
}

func (s *UpgradePrePayOrderRequestConfluentConfig) SetConnectReplica(v int32) *UpgradePrePayOrderRequestConfluentConfig {
	s.ConnectReplica = &v
	return s
}

func (s *UpgradePrePayOrderRequestConfluentConfig) SetControlCenterCU(v int32) *UpgradePrePayOrderRequestConfluentConfig {
	s.ControlCenterCU = &v
	return s
}

func (s *UpgradePrePayOrderRequestConfluentConfig) SetControlCenterReplica(v int32) *UpgradePrePayOrderRequestConfluentConfig {
	s.ControlCenterReplica = &v
	return s
}

func (s *UpgradePrePayOrderRequestConfluentConfig) SetControlCenterStorage(v int32) *UpgradePrePayOrderRequestConfluentConfig {
	s.ControlCenterStorage = &v
	return s
}

func (s *UpgradePrePayOrderRequestConfluentConfig) SetKafkaCU(v int32) *UpgradePrePayOrderRequestConfluentConfig {
	s.KafkaCU = &v
	return s
}

func (s *UpgradePrePayOrderRequestConfluentConfig) SetKafkaReplica(v int32) *UpgradePrePayOrderRequestConfluentConfig {
	s.KafkaReplica = &v
	return s
}

func (s *UpgradePrePayOrderRequestConfluentConfig) SetKafkaRestProxyCU(v int32) *UpgradePrePayOrderRequestConfluentConfig {
	s.KafkaRestProxyCU = &v
	return s
}

func (s *UpgradePrePayOrderRequestConfluentConfig) SetKafkaRestProxyReplica(v int32) *UpgradePrePayOrderRequestConfluentConfig {
	s.KafkaRestProxyReplica = &v
	return s
}

func (s *UpgradePrePayOrderRequestConfluentConfig) SetKafkaStorage(v int32) *UpgradePrePayOrderRequestConfluentConfig {
	s.KafkaStorage = &v
	return s
}

func (s *UpgradePrePayOrderRequestConfluentConfig) SetKsqlCU(v int32) *UpgradePrePayOrderRequestConfluentConfig {
	s.KsqlCU = &v
	return s
}

func (s *UpgradePrePayOrderRequestConfluentConfig) SetKsqlReplica(v int32) *UpgradePrePayOrderRequestConfluentConfig {
	s.KsqlReplica = &v
	return s
}

func (s *UpgradePrePayOrderRequestConfluentConfig) SetKsqlStorage(v int32) *UpgradePrePayOrderRequestConfluentConfig {
	s.KsqlStorage = &v
	return s
}

func (s *UpgradePrePayOrderRequestConfluentConfig) SetSchemaRegistryCU(v int32) *UpgradePrePayOrderRequestConfluentConfig {
	s.SchemaRegistryCU = &v
	return s
}

func (s *UpgradePrePayOrderRequestConfluentConfig) SetSchemaRegistryReplica(v int32) *UpgradePrePayOrderRequestConfluentConfig {
	s.SchemaRegistryReplica = &v
	return s
}

func (s *UpgradePrePayOrderRequestConfluentConfig) SetZooKeeperCU(v int32) *UpgradePrePayOrderRequestConfluentConfig {
	s.ZooKeeperCU = &v
	return s
}

func (s *UpgradePrePayOrderRequestConfluentConfig) SetZooKeeperReplica(v int32) *UpgradePrePayOrderRequestConfluentConfig {
	s.ZooKeeperReplica = &v
	return s
}

func (s *UpgradePrePayOrderRequestConfluentConfig) SetZooKeeperStorage(v int32) *UpgradePrePayOrderRequestConfluentConfig {
	s.ZooKeeperStorage = &v
	return s
}

func (s *UpgradePrePayOrderRequestConfluentConfig) Validate() error {
	return dara.Validate(s)
}
