// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePrePayOrderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfluentConfig(v *CreatePrePayOrderRequestConfluentConfig) *CreatePrePayOrderRequest
	GetConfluentConfig() *CreatePrePayOrderRequestConfluentConfig
	SetDeployType(v int32) *CreatePrePayOrderRequest
	GetDeployType() *int32
	SetDiskSize(v int32) *CreatePrePayOrderRequest
	GetDiskSize() *int32
	SetDiskType(v string) *CreatePrePayOrderRequest
	GetDiskType() *string
	SetDuration(v int32) *CreatePrePayOrderRequest
	GetDuration() *int32
	SetEipMax(v int32) *CreatePrePayOrderRequest
	GetEipMax() *int32
	SetIoMax(v int32) *CreatePrePayOrderRequest
	GetIoMax() *int32
	SetIoMaxSpec(v string) *CreatePrePayOrderRequest
	GetIoMaxSpec() *string
	SetPaidType(v int32) *CreatePrePayOrderRequest
	GetPaidType() *int32
	SetPartitionNum(v int32) *CreatePrePayOrderRequest
	GetPartitionNum() *int32
	SetRegionId(v string) *CreatePrePayOrderRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreatePrePayOrderRequest
	GetResourceGroupId() *string
	SetSpecType(v string) *CreatePrePayOrderRequest
	GetSpecType() *string
	SetTag(v []*CreatePrePayOrderRequestTag) *CreatePrePayOrderRequest
	GetTag() []*CreatePrePayOrderRequestTag
	SetTopicQuota(v int32) *CreatePrePayOrderRequest
	GetTopicQuota() *int32
}

type CreatePrePayOrderRequest struct {
	// The configurations of Confluent.
	//
	// >  When you create an ApsaraMQ for Confluent instance, you must configure this parameter.
	ConfluentConfig *CreatePrePayOrderRequestConfluentConfig `json:"ConfluentConfig,omitempty" xml:"ConfluentConfig,omitempty" type:"Struct"`
	// The type of the network in which the instance is deployed. Valid values:
	//
	// 	- **4**: Internet and virtual private cloud (VPC)
	//
	// 	- **5**: VPC
	//
	// >  If you create an ApsaraMQ for Confluent instance, set the value to 5. After the instance is created, you can specify whether to enable each component.
	//
	// example:
	//
	// 5
	DeployType *int32 `json:"DeployType,omitempty" xml:"DeployType,omitempty"`
	// The disk size. Unit: GB
	//
	// For information about the valid values of this parameter, see [Billing](https://help.aliyun.com/document_detail/84737.html).
	//
	// >  If you create an ApsaraMQ for Confluent instance, you do not need to configure this parameter.
	//
	// example:
	//
	// 500
	DiskSize *int32 `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
	// The disk type. Valid values:
	//
	// 	- **0**: ultra disk
	//
	// 	- **1**: standard SSD
	//
	// >  If you create an ApsaraMQ for Confluent instance, you do not need to configure this parameter.
	//
	// example:
	//
	// 0
	DiskType *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	// The subscription duration. Unit: months. Default value: 1. Valid values:
	//
	// 	- **1 to 12**
	//
	// example:
	//
	// 1
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The maximum Internet traffic in the instance.
	//
	// 	- If you set **DeployType*	- to **4**, you must configure this parameter.
	//
	// 	- For information about the valid values, see [Pay-as-you-go](https://help.aliyun.com/document_detail/72142.html).
	//
	// >  If you create an ApsaraMQ for Confluent instance, you do not need to configure this parameter.
	//
	// example:
	//
	// 0
	EipMax *int32 `json:"EipMax,omitempty" xml:"EipMax,omitempty"`
	// The maximum traffic in the instance. We recommend that you do not configure this parameter.
	//
	// 	- You must set one of **IoMax*	- and **IoMaxSpec**. If both parameters are configured, the value of **IoMaxSpec*	- is used. We recommend that you configure only **IoMaxSpec**.
	//
	// 	- For information about the valid values of this parameter, see [Billing](https://help.aliyun.com/document_detail/84737.html).
	//
	// >  If you create an ApsaraMQ for Confluent instance, you do not need to configure this parameter.
	//
	// example:
	//
	// 20
	IoMax *int32 `json:"IoMax,omitempty" xml:"IoMax,omitempty"`
	// The traffic specification of the instance. We recommend that you configure this parameter.
	//
	// 	- You must configure one of **IoMax*	- and **IoMaxSpec**. If both parameters are configured, the value of **IoMaxSpec*	- is used. We recommend that you configure only **IoMaxSpec**.
	//
	// 	- For more information about the valid values of this parameter, see [Billing](https://help.aliyun.com/document_detail/84737.html).
	//
	// >  If you create an ApsaraMQ for Confluent instance, you do not need to configure this parameter.
	//
	// example:
	//
	// alikafka.hw.2xlarge
	IoMaxSpec *string `json:"IoMaxSpec,omitempty" xml:"IoMaxSpec,omitempty"`
	// The billing method of the instance. Valid values:
	//
	// 	- **0**: the subscription billing method
	//
	// 	- **4**: the subscription billing method for ApsaraMQ for Confluent instances
	//
	// example:
	//
	// 1
	PaidType *int32 `json:"PaidType,omitempty" xml:"PaidType,omitempty"`
	// The number of partitions. We recommend that you configure this parameter.
	//
	// 	- You must configure one of PartitionNum and TopicQuota. We recommend that you configure only PartitionNum.
	//
	// 	- If you configure PartitionNum and TopicQuota at the same time, the system verifies whether the price of the partitions equals the price of the topics based on the previous topic-based selling mode. If the price of the partitions does not equal the price of the topics, an error is returned. If the price of the partitions equals the price of the topics, the instance is purchased based on the partition number.
	//
	// 	- For information about the valid values of this parameter, see [Billing](https://help.aliyun.com/document_detail/84737.html).
	//
	// >  If you create an ApsaraMQ for Confluent instance, you do not need to configure this parameter.
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
	// If this parameter is left empty, the default resource group is used. You can view the resource group ID on the Resource Group page in the Resource Management console.
	//
	// example:
	//
	// rg-ac***********7q
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The instance edition. Valid values:
	//
	// 	- **normal**: Standard Edition (High Write)
	//
	// 	- **professional**: Professional Edition (High Write)
	//
	// 	- **professionalForHighRead**: Professional Edition (High Read)
	//
	// For more information, see [Billing](https://help.aliyun.com/document_detail/84737.html).
	//
	// >  If you create an ApsaraMQ for Confluent instance, you do not need to configure this parameter.
	//
	// example:
	//
	// normal
	SpecType *string `json:"SpecType,omitempty" xml:"SpecType,omitempty"`
	// The tags.
	Tag []*CreatePrePayOrderRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The number of topics. We recommend that you do not configure this parameter.
	//
	// 	- You must configure one of PartitionNum and TopicQuota. We recommend that you configure only PartitionNum.
	//
	// 	- If you configure PartitionNum and TopicQuota at the same time, the system verifies whether the price of the partitions equals the price of the topics based on the previous topic-based selling mode. If the price of the partitions does not equal the price of the topics, an error is returned. If the price of the partitions equals the price of the topics, the instance is purchased based on the partition number.
	//
	// 	- The default value of TopicQuota varies based on the value of IoMaxSpec. If the number of topics that you use exceeds the default value, you are charged additional fees.
	//
	// 	- For information about the valid values of this parameter, see [Billing](https://help.aliyun.com/document_detail/84737.html).
	//
	// >  If you create an ApsaraMQ for Confluent instance, you do not need to configure this parameter.
	//
	// example:
	//
	// 50
	TopicQuota *int32 `json:"TopicQuota,omitempty" xml:"TopicQuota,omitempty"`
}

func (s CreatePrePayOrderRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePrePayOrderRequest) GoString() string {
	return s.String()
}

func (s *CreatePrePayOrderRequest) GetConfluentConfig() *CreatePrePayOrderRequestConfluentConfig {
	return s.ConfluentConfig
}

func (s *CreatePrePayOrderRequest) GetDeployType() *int32 {
	return s.DeployType
}

func (s *CreatePrePayOrderRequest) GetDiskSize() *int32 {
	return s.DiskSize
}

func (s *CreatePrePayOrderRequest) GetDiskType() *string {
	return s.DiskType
}

func (s *CreatePrePayOrderRequest) GetDuration() *int32 {
	return s.Duration
}

func (s *CreatePrePayOrderRequest) GetEipMax() *int32 {
	return s.EipMax
}

func (s *CreatePrePayOrderRequest) GetIoMax() *int32 {
	return s.IoMax
}

func (s *CreatePrePayOrderRequest) GetIoMaxSpec() *string {
	return s.IoMaxSpec
}

func (s *CreatePrePayOrderRequest) GetPaidType() *int32 {
	return s.PaidType
}

func (s *CreatePrePayOrderRequest) GetPartitionNum() *int32 {
	return s.PartitionNum
}

func (s *CreatePrePayOrderRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreatePrePayOrderRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreatePrePayOrderRequest) GetSpecType() *string {
	return s.SpecType
}

func (s *CreatePrePayOrderRequest) GetTag() []*CreatePrePayOrderRequestTag {
	return s.Tag
}

func (s *CreatePrePayOrderRequest) GetTopicQuota() *int32 {
	return s.TopicQuota
}

func (s *CreatePrePayOrderRequest) SetConfluentConfig(v *CreatePrePayOrderRequestConfluentConfig) *CreatePrePayOrderRequest {
	s.ConfluentConfig = v
	return s
}

func (s *CreatePrePayOrderRequest) SetDeployType(v int32) *CreatePrePayOrderRequest {
	s.DeployType = &v
	return s
}

func (s *CreatePrePayOrderRequest) SetDiskSize(v int32) *CreatePrePayOrderRequest {
	s.DiskSize = &v
	return s
}

func (s *CreatePrePayOrderRequest) SetDiskType(v string) *CreatePrePayOrderRequest {
	s.DiskType = &v
	return s
}

func (s *CreatePrePayOrderRequest) SetDuration(v int32) *CreatePrePayOrderRequest {
	s.Duration = &v
	return s
}

func (s *CreatePrePayOrderRequest) SetEipMax(v int32) *CreatePrePayOrderRequest {
	s.EipMax = &v
	return s
}

func (s *CreatePrePayOrderRequest) SetIoMax(v int32) *CreatePrePayOrderRequest {
	s.IoMax = &v
	return s
}

func (s *CreatePrePayOrderRequest) SetIoMaxSpec(v string) *CreatePrePayOrderRequest {
	s.IoMaxSpec = &v
	return s
}

func (s *CreatePrePayOrderRequest) SetPaidType(v int32) *CreatePrePayOrderRequest {
	s.PaidType = &v
	return s
}

func (s *CreatePrePayOrderRequest) SetPartitionNum(v int32) *CreatePrePayOrderRequest {
	s.PartitionNum = &v
	return s
}

func (s *CreatePrePayOrderRequest) SetRegionId(v string) *CreatePrePayOrderRequest {
	s.RegionId = &v
	return s
}

func (s *CreatePrePayOrderRequest) SetResourceGroupId(v string) *CreatePrePayOrderRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreatePrePayOrderRequest) SetSpecType(v string) *CreatePrePayOrderRequest {
	s.SpecType = &v
	return s
}

func (s *CreatePrePayOrderRequest) SetTag(v []*CreatePrePayOrderRequestTag) *CreatePrePayOrderRequest {
	s.Tag = v
	return s
}

func (s *CreatePrePayOrderRequest) SetTopicQuota(v int32) *CreatePrePayOrderRequest {
	s.TopicQuota = &v
	return s
}

func (s *CreatePrePayOrderRequest) Validate() error {
	return dara.Validate(s)
}

type CreatePrePayOrderRequestConfluentConfig struct {
	// The number of CPU cores of Connect.
	//
	// example:
	//
	// 4
	ConnectCU *int32 `json:"ConnectCU,omitempty" xml:"ConnectCU,omitempty"`
	// The number of replicas of Connect.
	//
	// example:
	//
	// 2
	ConnectReplica *int32 `json:"ConnectReplica,omitempty" xml:"ConnectReplica,omitempty"`
	// The number of CPU cores of Control Center.
	//
	// example:
	//
	// 4
	ControlCenterCU *int32 `json:"ControlCenterCU,omitempty" xml:"ControlCenterCU,omitempty"`
	// The number of replicas of Control Center.
	//
	// example:
	//
	// 1
	ControlCenterReplica *int32 `json:"ControlCenterReplica,omitempty" xml:"ControlCenterReplica,omitempty"`
	// The disk capacity of Control Center. Unit: GB
	//
	// example:
	//
	// 300
	ControlCenterStorage *int32 `json:"ControlCenterStorage,omitempty" xml:"ControlCenterStorage,omitempty"`
	// The number of CPU cores of the Kafka broker.
	//
	// example:
	//
	// 4
	KafkaCU *int32 `json:"KafkaCU,omitempty" xml:"KafkaCU,omitempty"`
	// The number of replicas of the Kafka broker.
	//
	// example:
	//
	// 3
	KafkaReplica *int32 `json:"KafkaReplica,omitempty" xml:"KafkaReplica,omitempty"`
	// The number of CPU cores of Kafka Rest Proxy.
	//
	// example:
	//
	// 4
	KafkaRestProxyCU *int32 `json:"KafkaRestProxyCU,omitempty" xml:"KafkaRestProxyCU,omitempty"`
	// The number of replicas of Kafka Rest Proxy.
	//
	// example:
	//
	// 2
	KafkaRestProxyReplica *int32 `json:"KafkaRestProxyReplica,omitempty" xml:"KafkaRestProxyReplica,omitempty"`
	// The disk capacity of the Kafka broker. Unit: GB
	//
	// example:
	//
	// 800
	KafkaStorage *int32 `json:"KafkaStorage,omitempty" xml:"KafkaStorage,omitempty"`
	// The number of CPU cores of ksqIDB.
	//
	// example:
	//
	// 4
	KsqlCU *int32 `json:"KsqlCU,omitempty" xml:"KsqlCU,omitempty"`
	// The number of replicas of ksqlDB.
	//
	// example:
	//
	// 2
	KsqlReplica *int32 `json:"KsqlReplica,omitempty" xml:"KsqlReplica,omitempty"`
	// The disk capacity of ksqlDB. Unit: GB
	//
	// example:
	//
	// 100
	KsqlStorage *int32 `json:"KsqlStorage,omitempty" xml:"KsqlStorage,omitempty"`
	// The number of CPU cores of Schema Registry.
	//
	// example:
	//
	// 1
	SchemaRegistryCU *int32 `json:"SchemaRegistryCU,omitempty" xml:"SchemaRegistryCU,omitempty"`
	// The number of replicas of Schema Registry.
	//
	// example:
	//
	// 2
	SchemaRegistryReplica *int32 `json:"SchemaRegistryReplica,omitempty" xml:"SchemaRegistryReplica,omitempty"`
	// The number of CPU cores of ZooKeeper.
	//
	// example:
	//
	// 2
	ZooKeeperCU *int32 `json:"ZooKeeperCU,omitempty" xml:"ZooKeeperCU,omitempty"`
	// The number of replicas of ZooKeeper.
	//
	// example:
	//
	// 3
	ZooKeeperReplica *int32 `json:"ZooKeeperReplica,omitempty" xml:"ZooKeeperReplica,omitempty"`
	// The disk capacity of ZooKeeper. Unit: GB
	//
	// example:
	//
	// 100
	ZooKeeperStorage *int32 `json:"ZooKeeperStorage,omitempty" xml:"ZooKeeperStorage,omitempty"`
}

func (s CreatePrePayOrderRequestConfluentConfig) String() string {
	return dara.Prettify(s)
}

func (s CreatePrePayOrderRequestConfluentConfig) GoString() string {
	return s.String()
}

func (s *CreatePrePayOrderRequestConfluentConfig) GetConnectCU() *int32 {
	return s.ConnectCU
}

func (s *CreatePrePayOrderRequestConfluentConfig) GetConnectReplica() *int32 {
	return s.ConnectReplica
}

func (s *CreatePrePayOrderRequestConfluentConfig) GetControlCenterCU() *int32 {
	return s.ControlCenterCU
}

func (s *CreatePrePayOrderRequestConfluentConfig) GetControlCenterReplica() *int32 {
	return s.ControlCenterReplica
}

func (s *CreatePrePayOrderRequestConfluentConfig) GetControlCenterStorage() *int32 {
	return s.ControlCenterStorage
}

func (s *CreatePrePayOrderRequestConfluentConfig) GetKafkaCU() *int32 {
	return s.KafkaCU
}

func (s *CreatePrePayOrderRequestConfluentConfig) GetKafkaReplica() *int32 {
	return s.KafkaReplica
}

func (s *CreatePrePayOrderRequestConfluentConfig) GetKafkaRestProxyCU() *int32 {
	return s.KafkaRestProxyCU
}

func (s *CreatePrePayOrderRequestConfluentConfig) GetKafkaRestProxyReplica() *int32 {
	return s.KafkaRestProxyReplica
}

func (s *CreatePrePayOrderRequestConfluentConfig) GetKafkaStorage() *int32 {
	return s.KafkaStorage
}

func (s *CreatePrePayOrderRequestConfluentConfig) GetKsqlCU() *int32 {
	return s.KsqlCU
}

func (s *CreatePrePayOrderRequestConfluentConfig) GetKsqlReplica() *int32 {
	return s.KsqlReplica
}

func (s *CreatePrePayOrderRequestConfluentConfig) GetKsqlStorage() *int32 {
	return s.KsqlStorage
}

func (s *CreatePrePayOrderRequestConfluentConfig) GetSchemaRegistryCU() *int32 {
	return s.SchemaRegistryCU
}

func (s *CreatePrePayOrderRequestConfluentConfig) GetSchemaRegistryReplica() *int32 {
	return s.SchemaRegistryReplica
}

func (s *CreatePrePayOrderRequestConfluentConfig) GetZooKeeperCU() *int32 {
	return s.ZooKeeperCU
}

func (s *CreatePrePayOrderRequestConfluentConfig) GetZooKeeperReplica() *int32 {
	return s.ZooKeeperReplica
}

func (s *CreatePrePayOrderRequestConfluentConfig) GetZooKeeperStorage() *int32 {
	return s.ZooKeeperStorage
}

func (s *CreatePrePayOrderRequestConfluentConfig) SetConnectCU(v int32) *CreatePrePayOrderRequestConfluentConfig {
	s.ConnectCU = &v
	return s
}

func (s *CreatePrePayOrderRequestConfluentConfig) SetConnectReplica(v int32) *CreatePrePayOrderRequestConfluentConfig {
	s.ConnectReplica = &v
	return s
}

func (s *CreatePrePayOrderRequestConfluentConfig) SetControlCenterCU(v int32) *CreatePrePayOrderRequestConfluentConfig {
	s.ControlCenterCU = &v
	return s
}

func (s *CreatePrePayOrderRequestConfluentConfig) SetControlCenterReplica(v int32) *CreatePrePayOrderRequestConfluentConfig {
	s.ControlCenterReplica = &v
	return s
}

func (s *CreatePrePayOrderRequestConfluentConfig) SetControlCenterStorage(v int32) *CreatePrePayOrderRequestConfluentConfig {
	s.ControlCenterStorage = &v
	return s
}

func (s *CreatePrePayOrderRequestConfluentConfig) SetKafkaCU(v int32) *CreatePrePayOrderRequestConfluentConfig {
	s.KafkaCU = &v
	return s
}

func (s *CreatePrePayOrderRequestConfluentConfig) SetKafkaReplica(v int32) *CreatePrePayOrderRequestConfluentConfig {
	s.KafkaReplica = &v
	return s
}

func (s *CreatePrePayOrderRequestConfluentConfig) SetKafkaRestProxyCU(v int32) *CreatePrePayOrderRequestConfluentConfig {
	s.KafkaRestProxyCU = &v
	return s
}

func (s *CreatePrePayOrderRequestConfluentConfig) SetKafkaRestProxyReplica(v int32) *CreatePrePayOrderRequestConfluentConfig {
	s.KafkaRestProxyReplica = &v
	return s
}

func (s *CreatePrePayOrderRequestConfluentConfig) SetKafkaStorage(v int32) *CreatePrePayOrderRequestConfluentConfig {
	s.KafkaStorage = &v
	return s
}

func (s *CreatePrePayOrderRequestConfluentConfig) SetKsqlCU(v int32) *CreatePrePayOrderRequestConfluentConfig {
	s.KsqlCU = &v
	return s
}

func (s *CreatePrePayOrderRequestConfluentConfig) SetKsqlReplica(v int32) *CreatePrePayOrderRequestConfluentConfig {
	s.KsqlReplica = &v
	return s
}

func (s *CreatePrePayOrderRequestConfluentConfig) SetKsqlStorage(v int32) *CreatePrePayOrderRequestConfluentConfig {
	s.KsqlStorage = &v
	return s
}

func (s *CreatePrePayOrderRequestConfluentConfig) SetSchemaRegistryCU(v int32) *CreatePrePayOrderRequestConfluentConfig {
	s.SchemaRegistryCU = &v
	return s
}

func (s *CreatePrePayOrderRequestConfluentConfig) SetSchemaRegistryReplica(v int32) *CreatePrePayOrderRequestConfluentConfig {
	s.SchemaRegistryReplica = &v
	return s
}

func (s *CreatePrePayOrderRequestConfluentConfig) SetZooKeeperCU(v int32) *CreatePrePayOrderRequestConfluentConfig {
	s.ZooKeeperCU = &v
	return s
}

func (s *CreatePrePayOrderRequestConfluentConfig) SetZooKeeperReplica(v int32) *CreatePrePayOrderRequestConfluentConfig {
	s.ZooKeeperReplica = &v
	return s
}

func (s *CreatePrePayOrderRequestConfluentConfig) SetZooKeeperStorage(v int32) *CreatePrePayOrderRequestConfluentConfig {
	s.ZooKeeperStorage = &v
	return s
}

func (s *CreatePrePayOrderRequestConfluentConfig) Validate() error {
	return dara.Validate(s)
}

type CreatePrePayOrderRequestTag struct {
	// The key of tag N.
	//
	// 	- Valid values of N: 1 to 20.
	//
	// 	- If this parameter is left empty, the keys of all tags are matched.
	//
	// 	- The tag key can be up to 128 characters in length and cannot start with acs: or aliyun or contain [http:// or https://.](http://https://。)
	//
	// This parameter is required.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N.
	//
	// 	- Valid values of N: 1 to 20.
	//
	// 	- This parameter can be left empty.
	//
	// 	- The tag value can be 1 to 128 characters in length and cannot start with acs: or aliyun or contain [http:// or https://.](http://https://。)
	//
	// example:
	//
	// FinanceJoshua
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreatePrePayOrderRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreatePrePayOrderRequestTag) GoString() string {
	return s.String()
}

func (s *CreatePrePayOrderRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreatePrePayOrderRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreatePrePayOrderRequestTag) SetKey(v string) *CreatePrePayOrderRequestTag {
	s.Key = &v
	return s
}

func (s *CreatePrePayOrderRequestTag) SetValue(v string) *CreatePrePayOrderRequestTag {
	s.Value = &v
	return s
}

func (s *CreatePrePayOrderRequestTag) Validate() error {
	return dara.Validate(s)
}
