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
	// The Confluent component configurations.
	//
	//
	// > This parameter is required when you create a Confluent instance.
	ConfluentConfig *CreatePrePayOrderRequestConfluentConfig `json:"ConfluentConfig,omitempty" xml:"ConfluentConfig,omitempty" type:"Struct"`
	// The deployment type. Valid values:
	//
	// - **4**: Internet- and VPC-connected instance
	//
	// - **5**: VPC-connected instance
	//
	//
	// > If you create a Confluent instance, the deployment type is not supported. You can only set this parameter to 5. After the purchase, you can configure whether to enable public access for each component.
	//
	// example:
	//
	// 5
	DeployType *int32 `json:"DeployType,omitempty" xml:"DeployType,omitempty"`
	// The disk capacity. Unit: GB.
	//
	// For the value range, see [Billing](https://help.aliyun.com/document_detail/84737.html).
	//
	// > If you create a Confluent instance, you do not need to specify this parameter.
	//
	// example:
	//
	// 500
	DiskSize *int32 `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
	// The disk type. Valid values:
	//
	// - **0**: ultra cloud disk
	//
	// - **1**: SSD
	//
	// > If you create a Confluent instance, you do not need to specify this parameter.
	//
	// example:
	//
	// 0
	DiskType *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	// The subscription duration. Unit: months. Default value: 1. Valid values:
	//
	// - **Confluent instances: 1 or 12**
	//
	// - **Kafka instances: 1**
	//
	// example:
	//
	// 1
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The public network traffic.
	//
	// - This parameter is required if **DeployType*	- is set to **4**.
	//
	// - For the value range, see [Pay-as-you-go billing method](https://help.aliyun.com/document_detail/72142.html).
	//
	//
	// > If you create a Confluent instance, you do not need to specify this parameter.
	//
	// example:
	//
	// 0
	EipMax *int32 `json:"EipMax,omitempty" xml:"EipMax,omitempty"`
	// The peak traffic (not recommended).
	//
	// - You must specify at least one of **IoMax*	- and **IoMaxSpec**. If you specify both, **IoMaxSpec*	- takes precedence. We recommend that you specify only **IoMaxSpec**.
	//
	// - For the value range, see [Billing](https://help.aliyun.com/document_detail/84737.html).
	//
	// > If you create a Confluent instance, you do not need to specify this parameter.
	//
	// example:
	//
	// 20
	IoMax *int32 `json:"IoMax,omitempty" xml:"IoMax,omitempty"`
	// The traffic specification (recommended).
	//
	// - You must specify at least one of **IoMax*	- and **IoMaxSpec**. If you specify both, **IoMaxSpec*	- takes precedence. We recommend that you specify only **IoMaxSpec**.
	//
	// - For the value range, see [Billing](https://help.aliyun.com/document_detail/84737.html).
	//
	// > If you create a Confluent instance, you do not need to specify this parameter.
	//
	// example:
	//
	// alikafka.hw.2xlarge
	IoMaxSpec *string `json:"IoMaxSpec,omitempty" xml:"IoMaxSpec,omitempty"`
	// The billing type. Valid values:
	//
	// - **0**: subscription
	//
	// - **4**: Confluent subscription
	//
	// example:
	//
	// 0
	PaidType *int32 `json:"PaidType,omitempty" xml:"PaidType,omitempty"`
	// The number of partitions (recommended).
	//
	// 	- You must specify either the number of partitions or the topic specification. We recommend that you specify only the number of partitions.
	//
	// 	- If you specify both the number of partitions and the topic specification, the system verifies whether the number of partitions and the topic specification are equivalent based on the legacy topic sales model. If they are not equivalent, the request fails. If they are equivalent, the purchase is made based on the number of partitions.
	//
	// 	- For the value range, see [Billing](https://help.aliyun.com/document_detail/84737.html).
	//
	// > If you create a Confluent instance, you do not need to specify this parameter.
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
	// If you do not specify this parameter, the instance is placed in the default resource group. You can view the resource group ID in the Resource Group console.
	//
	// example:
	//
	// rg-ac***********7q
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The specification type.
	//
	// Valid values for ApsaraMQ for Kafka instances:
	//
	// - **normal**: Normal Edition (shared high-write)
	//
	// - **professional**: Professional Edition (shared high-write)
	//
	// - **professionalForHighRead**: Professional Edition (shared high-read)
	//
	// Valid values for Confluent instances:
	//
	// - **professional**: Professional Edition
	//
	// - **enterprise**: Enterprise Edition
	//
	// For more information, see [Billing](https://help.aliyun.com/document_detail/84737.html).
	//
	// example:
	//
	// normal
	SpecType *string `json:"SpecType,omitempty" xml:"SpecType,omitempty"`
	// The tags.
	Tag []*CreatePrePayOrderRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The number of topics (not recommended).
	//
	// - You must specify either the number of partitions or the topic specification. We recommend that you specify only the number of partitions.
	//
	// - If you specify both the number of partitions and the topic specification, the system verifies whether the number of partitions and the topic specification are equivalent based on the legacy topic sales model. If they are not equivalent, the request fails. If they are equivalent, the purchase is made based on the number of partitions.
	//
	// - The default value varies based on the traffic specification. Additional fees are charged if the value exceeds the default value.
	//
	// - For the value range, see [Billing](https://help.aliyun.com/document_detail/84737.html).
	//
	// > If you create a Confluent instance, you do not need to specify this parameter.
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
	if s.ConfluentConfig != nil {
		if err := s.ConfluentConfig.Validate(); err != nil {
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

type CreatePrePayOrderRequestConfluentConfig struct {
	ConfluentVersion *string `json:"ConfluentVersion,omitempty" xml:"ConfluentVersion,omitempty"`
	// The number of CPU cores for the Connect component.
	//
	// example:
	//
	// 4
	ConnectCU *int32 `json:"ConnectCU,omitempty" xml:"ConnectCU,omitempty"`
	// The number of Connect component replicas.
	//
	// example:
	//
	// 2
	ConnectReplica *int32 `json:"ConnectReplica,omitempty" xml:"ConnectReplica,omitempty"`
	// The number of CPU cores for the ControlCenter component.
	//
	// example:
	//
	// 4
	ControlCenterCU *int32 `json:"ControlCenterCU,omitempty" xml:"ControlCenterCU,omitempty"`
	// The number of ControlCenter component replicas.
	//
	// example:
	//
	// 1
	ControlCenterReplica *int32 `json:"ControlCenterReplica,omitempty" xml:"ControlCenterReplica,omitempty"`
	// The disk capacity of the ControlCenter component. Unit: GB.
	//
	// example:
	//
	// 300
	ControlCenterStorage *int32 `json:"ControlCenterStorage,omitempty" xml:"ControlCenterStorage,omitempty"`
	// The number of CPU cores for Kafka Broker.
	//
	// example:
	//
	// 4
	KafkaCU *int32 `json:"KafkaCU,omitempty" xml:"KafkaCU,omitempty"`
	// The number of Kafka Broker replicas.
	//
	// example:
	//
	// 3
	KafkaReplica *int32 `json:"KafkaReplica,omitempty" xml:"KafkaReplica,omitempty"`
	// The number of CPU cores for the KafkaRestProxy component.
	//
	// example:
	//
	// 4
	KafkaRestProxyCU *int32 `json:"KafkaRestProxyCU,omitempty" xml:"KafkaRestProxyCU,omitempty"`
	// The number of KafkaRestProxy component replicas.
	//
	// example:
	//
	// 2
	KafkaRestProxyReplica *int32 `json:"KafkaRestProxyReplica,omitempty" xml:"KafkaRestProxyReplica,omitempty"`
	// The disk capacity of Kafka Broker. Unit: GB.
	//
	// example:
	//
	// 800
	KafkaStorage           *int32 `json:"KafkaStorage,omitempty" xml:"KafkaStorage,omitempty"`
	KraftControllerCU      *int32 `json:"KraftControllerCU,omitempty" xml:"KraftControllerCU,omitempty"`
	KraftControllerReplica *int32 `json:"KraftControllerReplica,omitempty" xml:"KraftControllerReplica,omitempty"`
	KraftControllerStorage *int32 `json:"KraftControllerStorage,omitempty" xml:"KraftControllerStorage,omitempty"`
	// The number of CPU cores for the KsqlDB component.
	//
	// example:
	//
	// 4
	KsqlCU   *int32                                             `json:"KsqlCU,omitempty" xml:"KsqlCU,omitempty"`
	KsqlList []*CreatePrePayOrderRequestConfluentConfigKsqlList `json:"KsqlList,omitempty" xml:"KsqlList,omitempty" type:"Repeated"`
	// The number of KsqlDB component replicas.
	//
	// example:
	//
	// 2
	KsqlReplica *int32 `json:"KsqlReplica,omitempty" xml:"KsqlReplica,omitempty"`
	// The disk capacity of the KsqlDB component. Unit: GB.
	//
	// example:
	//
	// 100
	KsqlStorage *int32 `json:"KsqlStorage,omitempty" xml:"KsqlStorage,omitempty"`
	// The number of CPU cores for the SchemaRegistry component.
	//
	// example:
	//
	// 1
	SchemaRegistryCU *int32 `json:"SchemaRegistryCU,omitempty" xml:"SchemaRegistryCU,omitempty"`
	// The number of SchemaRegistry component replicas.
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
	// The number of ZooKeeper component replicas.
	//
	// example:
	//
	// 3
	ZooKeeperReplica *int32 `json:"ZooKeeperReplica,omitempty" xml:"ZooKeeperReplica,omitempty"`
	// The disk capacity of the ZooKeeper component. Unit: GB.
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

func (s *CreatePrePayOrderRequestConfluentConfig) GetConfluentVersion() *string {
	return s.ConfluentVersion
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

func (s *CreatePrePayOrderRequestConfluentConfig) GetKraftControllerCU() *int32 {
	return s.KraftControllerCU
}

func (s *CreatePrePayOrderRequestConfluentConfig) GetKraftControllerReplica() *int32 {
	return s.KraftControllerReplica
}

func (s *CreatePrePayOrderRequestConfluentConfig) GetKraftControllerStorage() *int32 {
	return s.KraftControllerStorage
}

func (s *CreatePrePayOrderRequestConfluentConfig) GetKsqlCU() *int32 {
	return s.KsqlCU
}

func (s *CreatePrePayOrderRequestConfluentConfig) GetKsqlList() []*CreatePrePayOrderRequestConfluentConfigKsqlList {
	return s.KsqlList
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

func (s *CreatePrePayOrderRequestConfluentConfig) SetConfluentVersion(v string) *CreatePrePayOrderRequestConfluentConfig {
	s.ConfluentVersion = &v
	return s
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

func (s *CreatePrePayOrderRequestConfluentConfig) SetKraftControllerCU(v int32) *CreatePrePayOrderRequestConfluentConfig {
	s.KraftControllerCU = &v
	return s
}

func (s *CreatePrePayOrderRequestConfluentConfig) SetKraftControllerReplica(v int32) *CreatePrePayOrderRequestConfluentConfig {
	s.KraftControllerReplica = &v
	return s
}

func (s *CreatePrePayOrderRequestConfluentConfig) SetKraftControllerStorage(v int32) *CreatePrePayOrderRequestConfluentConfig {
	s.KraftControllerStorage = &v
	return s
}

func (s *CreatePrePayOrderRequestConfluentConfig) SetKsqlCU(v int32) *CreatePrePayOrderRequestConfluentConfig {
	s.KsqlCU = &v
	return s
}

func (s *CreatePrePayOrderRequestConfluentConfig) SetKsqlList(v []*CreatePrePayOrderRequestConfluentConfigKsqlList) *CreatePrePayOrderRequestConfluentConfig {
	s.KsqlList = v
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

type CreatePrePayOrderRequestConfluentConfigKsqlList struct {
	Cu         *int32  `json:"Cu,omitempty" xml:"Cu,omitempty"`
	InternalId *string `json:"InternalId,omitempty" xml:"InternalId,omitempty"`
	Replica    *int32  `json:"Replica,omitempty" xml:"Replica,omitempty"`
	Storage    *int32  `json:"Storage,omitempty" xml:"Storage,omitempty"`
	Type       *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreatePrePayOrderRequestConfluentConfigKsqlList) String() string {
	return dara.Prettify(s)
}

func (s CreatePrePayOrderRequestConfluentConfigKsqlList) GoString() string {
	return s.String()
}

func (s *CreatePrePayOrderRequestConfluentConfigKsqlList) GetCu() *int32 {
	return s.Cu
}

func (s *CreatePrePayOrderRequestConfluentConfigKsqlList) GetInternalId() *string {
	return s.InternalId
}

func (s *CreatePrePayOrderRequestConfluentConfigKsqlList) GetReplica() *int32 {
	return s.Replica
}

func (s *CreatePrePayOrderRequestConfluentConfigKsqlList) GetStorage() *int32 {
	return s.Storage
}

func (s *CreatePrePayOrderRequestConfluentConfigKsqlList) GetType() *string {
	return s.Type
}

func (s *CreatePrePayOrderRequestConfluentConfigKsqlList) SetCu(v int32) *CreatePrePayOrderRequestConfluentConfigKsqlList {
	s.Cu = &v
	return s
}

func (s *CreatePrePayOrderRequestConfluentConfigKsqlList) SetInternalId(v string) *CreatePrePayOrderRequestConfluentConfigKsqlList {
	s.InternalId = &v
	return s
}

func (s *CreatePrePayOrderRequestConfluentConfigKsqlList) SetReplica(v int32) *CreatePrePayOrderRequestConfluentConfigKsqlList {
	s.Replica = &v
	return s
}

func (s *CreatePrePayOrderRequestConfluentConfigKsqlList) SetStorage(v int32) *CreatePrePayOrderRequestConfluentConfigKsqlList {
	s.Storage = &v
	return s
}

func (s *CreatePrePayOrderRequestConfluentConfigKsqlList) SetType(v string) *CreatePrePayOrderRequestConfluentConfigKsqlList {
	s.Type = &v
	return s
}

func (s *CreatePrePayOrderRequestConfluentConfigKsqlList) Validate() error {
	return dara.Validate(s)
}

type CreatePrePayOrderRequestTag struct {
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
	// - This parameter can be left empty.
	//
	// - The tag value can be up to 128 characters in length and cannot start with aliyun or acs:. It cannot contain http:// or https://.
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
