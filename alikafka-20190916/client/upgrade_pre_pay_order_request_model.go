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
	// Configurations for the Confluent components.
	ConfluentConfig *UpgradePrePayOrderRequestConfluentConfig `json:"ConfluentConfig,omitempty" xml:"ConfluentConfig,omitempty" type:"Struct"`
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
	if s.ConfluentConfig != nil {
		if err := s.ConfluentConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpgradePrePayOrderRequestConfluentConfig struct {
	// The number of CPU cores for the Connect component.
	//
	// example:
	//
	// 4
	ConnectCU *int32 `json:"ConnectCU,omitempty" xml:"ConnectCU,omitempty"`
	// The number of replicas for the Connect component.
	//
	// example:
	//
	// 2
	ConnectReplica *int32 `json:"ConnectReplica,omitempty" xml:"ConnectReplica,omitempty"`
	// The number of CPU cores for the Control Center component.
	//
	// example:
	//
	// 4
	ControlCenterCU *int32 `json:"ControlCenterCU,omitempty" xml:"ControlCenterCU,omitempty"`
	// The number of replicas for the Control Center component.
	//
	// example:
	//
	// 1
	ControlCenterReplica *int32 `json:"ControlCenterReplica,omitempty" xml:"ControlCenterReplica,omitempty"`
	// The disk capacity of the Control Center component, in GB.
	//
	// example:
	//
	// 300
	ControlCenterStorage *int32 `json:"ControlCenterStorage,omitempty" xml:"ControlCenterStorage,omitempty"`
	// The number of CPU cores for the Kafka broker.
	//
	// example:
	//
	// 4
	KafkaCU *int32 `json:"KafkaCU,omitempty" xml:"KafkaCU,omitempty"`
	// The number of replicas for the Kafka broker.
	//
	// example:
	//
	// 3
	KafkaReplica *int32 `json:"KafkaReplica,omitempty" xml:"KafkaReplica,omitempty"`
	// The number of CPU cores for the Kafka REST Proxy component.
	//
	// example:
	//
	// 4
	KafkaRestProxyCU *int32 `json:"KafkaRestProxyCU,omitempty" xml:"KafkaRestProxyCU,omitempty"`
	// The number of replicas for the Kafka REST Proxy component.
	//
	// example:
	//
	// 2
	KafkaRestProxyReplica *int32 `json:"KafkaRestProxyReplica,omitempty" xml:"KafkaRestProxyReplica,omitempty"`
	// The disk capacity of the Kafka broker, in GB.
	//
	// example:
	//
	// 800
	KafkaStorage           *int32 `json:"KafkaStorage,omitempty" xml:"KafkaStorage,omitempty"`
	KraftControllerCU      *int32 `json:"KraftControllerCU,omitempty" xml:"KraftControllerCU,omitempty"`
	KraftControllerReplica *int32 `json:"KraftControllerReplica,omitempty" xml:"KraftControllerReplica,omitempty"`
	KraftControllerStorage *int32 `json:"KraftControllerStorage,omitempty" xml:"KraftControllerStorage,omitempty"`
	// The number of CPU cores for the ksqlDB component.
	//
	// example:
	//
	// 4
	KsqlCU   *int32                                              `json:"KsqlCU,omitempty" xml:"KsqlCU,omitempty"`
	KsqlList []*UpgradePrePayOrderRequestConfluentConfigKsqlList `json:"KsqlList,omitempty" xml:"KsqlList,omitempty" type:"Repeated"`
	// The number of replicas for the ksqlDB component.
	//
	// example:
	//
	// 2
	KsqlReplica *int32 `json:"KsqlReplica,omitempty" xml:"KsqlReplica,omitempty"`
	// The disk capacity of the ksqlDB component, in GB.
	//
	// example:
	//
	// 100
	KsqlStorage *int32 `json:"KsqlStorage,omitempty" xml:"KsqlStorage,omitempty"`
	// The number of CPU cores for the Schema Registry component.
	//
	// example:
	//
	// 1
	SchemaRegistryCU *int32 `json:"SchemaRegistryCU,omitempty" xml:"SchemaRegistryCU,omitempty"`
	// The number of replicas for the Schema Registry component.
	//
	// example:
	//
	// 2
	SchemaRegistryReplica *int32 `json:"SchemaRegistryReplica,omitempty" xml:"SchemaRegistryReplica,omitempty"`
	// The number of CPU cores for the ZooKeeper component.
	//
	// example:
	//
	// 2
	ZooKeeperCU *int32 `json:"ZooKeeperCU,omitempty" xml:"ZooKeeperCU,omitempty"`
	// The number of replicas for the ZooKeeper component.
	//
	// example:
	//
	// 3
	ZooKeeperReplica *int32 `json:"ZooKeeperReplica,omitempty" xml:"ZooKeeperReplica,omitempty"`
	// The disk capacity of the ZooKeeper component, in GB.
	//
	// example:
	//
	// 100
	ZooKeeperStorage *int32 `json:"ZooKeeperStorage,omitempty" xml:"ZooKeeperStorage,omitempty"`
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

func (s *UpgradePrePayOrderRequestConfluentConfig) GetKraftControllerCU() *int32 {
	return s.KraftControllerCU
}

func (s *UpgradePrePayOrderRequestConfluentConfig) GetKraftControllerReplica() *int32 {
	return s.KraftControllerReplica
}

func (s *UpgradePrePayOrderRequestConfluentConfig) GetKraftControllerStorage() *int32 {
	return s.KraftControllerStorage
}

func (s *UpgradePrePayOrderRequestConfluentConfig) GetKsqlCU() *int32 {
	return s.KsqlCU
}

func (s *UpgradePrePayOrderRequestConfluentConfig) GetKsqlList() []*UpgradePrePayOrderRequestConfluentConfigKsqlList {
	return s.KsqlList
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

func (s *UpgradePrePayOrderRequestConfluentConfig) SetKraftControllerCU(v int32) *UpgradePrePayOrderRequestConfluentConfig {
	s.KraftControllerCU = &v
	return s
}

func (s *UpgradePrePayOrderRequestConfluentConfig) SetKraftControllerReplica(v int32) *UpgradePrePayOrderRequestConfluentConfig {
	s.KraftControllerReplica = &v
	return s
}

func (s *UpgradePrePayOrderRequestConfluentConfig) SetKraftControllerStorage(v int32) *UpgradePrePayOrderRequestConfluentConfig {
	s.KraftControllerStorage = &v
	return s
}

func (s *UpgradePrePayOrderRequestConfluentConfig) SetKsqlCU(v int32) *UpgradePrePayOrderRequestConfluentConfig {
	s.KsqlCU = &v
	return s
}

func (s *UpgradePrePayOrderRequestConfluentConfig) SetKsqlList(v []*UpgradePrePayOrderRequestConfluentConfigKsqlList) *UpgradePrePayOrderRequestConfluentConfig {
	s.KsqlList = v
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
	if s.KsqlList != nil {
		for _, item := range s.KsqlList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpgradePrePayOrderRequestConfluentConfigKsqlList struct {
	Cu         *int32  `json:"Cu,omitempty" xml:"Cu,omitempty"`
	InternalId *string `json:"InternalId,omitempty" xml:"InternalId,omitempty"`
	Replica    *int32  `json:"Replica,omitempty" xml:"Replica,omitempty"`
	Storage    *int32  `json:"Storage,omitempty" xml:"Storage,omitempty"`
	Type       *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpgradePrePayOrderRequestConfluentConfigKsqlList) String() string {
	return dara.Prettify(s)
}

func (s UpgradePrePayOrderRequestConfluentConfigKsqlList) GoString() string {
	return s.String()
}

func (s *UpgradePrePayOrderRequestConfluentConfigKsqlList) GetCu() *int32 {
	return s.Cu
}

func (s *UpgradePrePayOrderRequestConfluentConfigKsqlList) GetInternalId() *string {
	return s.InternalId
}

func (s *UpgradePrePayOrderRequestConfluentConfigKsqlList) GetReplica() *int32 {
	return s.Replica
}

func (s *UpgradePrePayOrderRequestConfluentConfigKsqlList) GetStorage() *int32 {
	return s.Storage
}

func (s *UpgradePrePayOrderRequestConfluentConfigKsqlList) GetType() *string {
	return s.Type
}

func (s *UpgradePrePayOrderRequestConfluentConfigKsqlList) SetCu(v int32) *UpgradePrePayOrderRequestConfluentConfigKsqlList {
	s.Cu = &v
	return s
}

func (s *UpgradePrePayOrderRequestConfluentConfigKsqlList) SetInternalId(v string) *UpgradePrePayOrderRequestConfluentConfigKsqlList {
	s.InternalId = &v
	return s
}

func (s *UpgradePrePayOrderRequestConfluentConfigKsqlList) SetReplica(v int32) *UpgradePrePayOrderRequestConfluentConfigKsqlList {
	s.Replica = &v
	return s
}

func (s *UpgradePrePayOrderRequestConfluentConfigKsqlList) SetStorage(v int32) *UpgradePrePayOrderRequestConfluentConfigKsqlList {
	s.Storage = &v
	return s
}

func (s *UpgradePrePayOrderRequestConfluentConfigKsqlList) SetType(v string) *UpgradePrePayOrderRequestConfluentConfigKsqlList {
	s.Type = &v
	return s
}

func (s *UpgradePrePayOrderRequestConfluentConfigKsqlList) Validate() error {
	return dara.Validate(s)
}
