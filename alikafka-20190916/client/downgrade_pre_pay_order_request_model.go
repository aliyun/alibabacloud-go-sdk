// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDowngradePrePayOrderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfluentConfig(v *DowngradePrePayOrderRequestConfluentConfig) *DowngradePrePayOrderRequest
	GetConfluentConfig() *DowngradePrePayOrderRequestConfluentConfig
	SetDiskSize(v int32) *DowngradePrePayOrderRequest
	GetDiskSize() *int32
	SetEipMax(v int32) *DowngradePrePayOrderRequest
	GetEipMax() *int32
	SetEipModel(v bool) *DowngradePrePayOrderRequest
	GetEipModel() *bool
	SetInstanceId(v string) *DowngradePrePayOrderRequest
	GetInstanceId() *string
	SetIoMax(v int32) *DowngradePrePayOrderRequest
	GetIoMax() *int32
	SetIoMaxSpec(v string) *DowngradePrePayOrderRequest
	GetIoMaxSpec() *string
	SetPaidType(v int32) *DowngradePrePayOrderRequest
	GetPaidType() *int32
	SetPartitionNum(v int32) *DowngradePrePayOrderRequest
	GetPartitionNum() *int32
	SetRegionId(v string) *DowngradePrePayOrderRequest
	GetRegionId() *string
	SetSpecType(v string) *DowngradePrePayOrderRequest
	GetSpecType() *string
	SetTopicQuota(v int32) *DowngradePrePayOrderRequest
	GetTopicQuota() *int32
}

type DowngradePrePayOrderRequest struct {
	ConfluentConfig *DowngradePrePayOrderRequestConfluentConfig `json:"ConfluentConfig,omitempty" xml:"ConfluentConfig,omitempty" type:"Struct"`
	DiskSize        *int32                                      `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
	EipMax          *int32                                      `json:"EipMax,omitempty" xml:"EipMax,omitempty"`
	EipModel        *bool                                       `json:"EipModel,omitempty" xml:"EipModel,omitempty"`
	// This parameter is required.
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	IoMax        *int32  `json:"IoMax,omitempty" xml:"IoMax,omitempty"`
	IoMaxSpec    *string `json:"IoMaxSpec,omitempty" xml:"IoMaxSpec,omitempty"`
	PaidType     *int32  `json:"PaidType,omitempty" xml:"PaidType,omitempty"`
	PartitionNum *int32  `json:"PartitionNum,omitempty" xml:"PartitionNum,omitempty"`
	// This parameter is required.
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SpecType   *string `json:"SpecType,omitempty" xml:"SpecType,omitempty"`
	TopicQuota *int32  `json:"TopicQuota,omitempty" xml:"TopicQuota,omitempty"`
}

func (s DowngradePrePayOrderRequest) String() string {
	return dara.Prettify(s)
}

func (s DowngradePrePayOrderRequest) GoString() string {
	return s.String()
}

func (s *DowngradePrePayOrderRequest) GetConfluentConfig() *DowngradePrePayOrderRequestConfluentConfig {
	return s.ConfluentConfig
}

func (s *DowngradePrePayOrderRequest) GetDiskSize() *int32 {
	return s.DiskSize
}

func (s *DowngradePrePayOrderRequest) GetEipMax() *int32 {
	return s.EipMax
}

func (s *DowngradePrePayOrderRequest) GetEipModel() *bool {
	return s.EipModel
}

func (s *DowngradePrePayOrderRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DowngradePrePayOrderRequest) GetIoMax() *int32 {
	return s.IoMax
}

func (s *DowngradePrePayOrderRequest) GetIoMaxSpec() *string {
	return s.IoMaxSpec
}

func (s *DowngradePrePayOrderRequest) GetPaidType() *int32 {
	return s.PaidType
}

func (s *DowngradePrePayOrderRequest) GetPartitionNum() *int32 {
	return s.PartitionNum
}

func (s *DowngradePrePayOrderRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DowngradePrePayOrderRequest) GetSpecType() *string {
	return s.SpecType
}

func (s *DowngradePrePayOrderRequest) GetTopicQuota() *int32 {
	return s.TopicQuota
}

func (s *DowngradePrePayOrderRequest) SetConfluentConfig(v *DowngradePrePayOrderRequestConfluentConfig) *DowngradePrePayOrderRequest {
	s.ConfluentConfig = v
	return s
}

func (s *DowngradePrePayOrderRequest) SetDiskSize(v int32) *DowngradePrePayOrderRequest {
	s.DiskSize = &v
	return s
}

func (s *DowngradePrePayOrderRequest) SetEipMax(v int32) *DowngradePrePayOrderRequest {
	s.EipMax = &v
	return s
}

func (s *DowngradePrePayOrderRequest) SetEipModel(v bool) *DowngradePrePayOrderRequest {
	s.EipModel = &v
	return s
}

func (s *DowngradePrePayOrderRequest) SetInstanceId(v string) *DowngradePrePayOrderRequest {
	s.InstanceId = &v
	return s
}

func (s *DowngradePrePayOrderRequest) SetIoMax(v int32) *DowngradePrePayOrderRequest {
	s.IoMax = &v
	return s
}

func (s *DowngradePrePayOrderRequest) SetIoMaxSpec(v string) *DowngradePrePayOrderRequest {
	s.IoMaxSpec = &v
	return s
}

func (s *DowngradePrePayOrderRequest) SetPaidType(v int32) *DowngradePrePayOrderRequest {
	s.PaidType = &v
	return s
}

func (s *DowngradePrePayOrderRequest) SetPartitionNum(v int32) *DowngradePrePayOrderRequest {
	s.PartitionNum = &v
	return s
}

func (s *DowngradePrePayOrderRequest) SetRegionId(v string) *DowngradePrePayOrderRequest {
	s.RegionId = &v
	return s
}

func (s *DowngradePrePayOrderRequest) SetSpecType(v string) *DowngradePrePayOrderRequest {
	s.SpecType = &v
	return s
}

func (s *DowngradePrePayOrderRequest) SetTopicQuota(v int32) *DowngradePrePayOrderRequest {
	s.TopicQuota = &v
	return s
}

func (s *DowngradePrePayOrderRequest) Validate() error {
	if s.ConfluentConfig != nil {
		if err := s.ConfluentConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DowngradePrePayOrderRequestConfluentConfig struct {
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

func (s DowngradePrePayOrderRequestConfluentConfig) String() string {
	return dara.Prettify(s)
}

func (s DowngradePrePayOrderRequestConfluentConfig) GoString() string {
	return s.String()
}

func (s *DowngradePrePayOrderRequestConfluentConfig) GetConnectCU() *int32 {
	return s.ConnectCU
}

func (s *DowngradePrePayOrderRequestConfluentConfig) GetConnectReplica() *int32 {
	return s.ConnectReplica
}

func (s *DowngradePrePayOrderRequestConfluentConfig) GetControlCenterCU() *int32 {
	return s.ControlCenterCU
}

func (s *DowngradePrePayOrderRequestConfluentConfig) GetControlCenterReplica() *int32 {
	return s.ControlCenterReplica
}

func (s *DowngradePrePayOrderRequestConfluentConfig) GetControlCenterStorage() *int32 {
	return s.ControlCenterStorage
}

func (s *DowngradePrePayOrderRequestConfluentConfig) GetKafkaCU() *int32 {
	return s.KafkaCU
}

func (s *DowngradePrePayOrderRequestConfluentConfig) GetKafkaReplica() *int32 {
	return s.KafkaReplica
}

func (s *DowngradePrePayOrderRequestConfluentConfig) GetKafkaRestProxyCU() *int32 {
	return s.KafkaRestProxyCU
}

func (s *DowngradePrePayOrderRequestConfluentConfig) GetKafkaRestProxyReplica() *int32 {
	return s.KafkaRestProxyReplica
}

func (s *DowngradePrePayOrderRequestConfluentConfig) GetKafkaStorage() *int32 {
	return s.KafkaStorage
}

func (s *DowngradePrePayOrderRequestConfluentConfig) GetKsqlCU() *int32 {
	return s.KsqlCU
}

func (s *DowngradePrePayOrderRequestConfluentConfig) GetKsqlReplica() *int32 {
	return s.KsqlReplica
}

func (s *DowngradePrePayOrderRequestConfluentConfig) GetKsqlStorage() *int32 {
	return s.KsqlStorage
}

func (s *DowngradePrePayOrderRequestConfluentConfig) GetSchemaRegistryCU() *int32 {
	return s.SchemaRegistryCU
}

func (s *DowngradePrePayOrderRequestConfluentConfig) GetSchemaRegistryReplica() *int32 {
	return s.SchemaRegistryReplica
}

func (s *DowngradePrePayOrderRequestConfluentConfig) GetZooKeeperCU() *int32 {
	return s.ZooKeeperCU
}

func (s *DowngradePrePayOrderRequestConfluentConfig) GetZooKeeperReplica() *int32 {
	return s.ZooKeeperReplica
}

func (s *DowngradePrePayOrderRequestConfluentConfig) GetZooKeeperStorage() *int32 {
	return s.ZooKeeperStorage
}

func (s *DowngradePrePayOrderRequestConfluentConfig) SetConnectCU(v int32) *DowngradePrePayOrderRequestConfluentConfig {
	s.ConnectCU = &v
	return s
}

func (s *DowngradePrePayOrderRequestConfluentConfig) SetConnectReplica(v int32) *DowngradePrePayOrderRequestConfluentConfig {
	s.ConnectReplica = &v
	return s
}

func (s *DowngradePrePayOrderRequestConfluentConfig) SetControlCenterCU(v int32) *DowngradePrePayOrderRequestConfluentConfig {
	s.ControlCenterCU = &v
	return s
}

func (s *DowngradePrePayOrderRequestConfluentConfig) SetControlCenterReplica(v int32) *DowngradePrePayOrderRequestConfluentConfig {
	s.ControlCenterReplica = &v
	return s
}

func (s *DowngradePrePayOrderRequestConfluentConfig) SetControlCenterStorage(v int32) *DowngradePrePayOrderRequestConfluentConfig {
	s.ControlCenterStorage = &v
	return s
}

func (s *DowngradePrePayOrderRequestConfluentConfig) SetKafkaCU(v int32) *DowngradePrePayOrderRequestConfluentConfig {
	s.KafkaCU = &v
	return s
}

func (s *DowngradePrePayOrderRequestConfluentConfig) SetKafkaReplica(v int32) *DowngradePrePayOrderRequestConfluentConfig {
	s.KafkaReplica = &v
	return s
}

func (s *DowngradePrePayOrderRequestConfluentConfig) SetKafkaRestProxyCU(v int32) *DowngradePrePayOrderRequestConfluentConfig {
	s.KafkaRestProxyCU = &v
	return s
}

func (s *DowngradePrePayOrderRequestConfluentConfig) SetKafkaRestProxyReplica(v int32) *DowngradePrePayOrderRequestConfluentConfig {
	s.KafkaRestProxyReplica = &v
	return s
}

func (s *DowngradePrePayOrderRequestConfluentConfig) SetKafkaStorage(v int32) *DowngradePrePayOrderRequestConfluentConfig {
	s.KafkaStorage = &v
	return s
}

func (s *DowngradePrePayOrderRequestConfluentConfig) SetKsqlCU(v int32) *DowngradePrePayOrderRequestConfluentConfig {
	s.KsqlCU = &v
	return s
}

func (s *DowngradePrePayOrderRequestConfluentConfig) SetKsqlReplica(v int32) *DowngradePrePayOrderRequestConfluentConfig {
	s.KsqlReplica = &v
	return s
}

func (s *DowngradePrePayOrderRequestConfluentConfig) SetKsqlStorage(v int32) *DowngradePrePayOrderRequestConfluentConfig {
	s.KsqlStorage = &v
	return s
}

func (s *DowngradePrePayOrderRequestConfluentConfig) SetSchemaRegistryCU(v int32) *DowngradePrePayOrderRequestConfluentConfig {
	s.SchemaRegistryCU = &v
	return s
}

func (s *DowngradePrePayOrderRequestConfluentConfig) SetSchemaRegistryReplica(v int32) *DowngradePrePayOrderRequestConfluentConfig {
	s.SchemaRegistryReplica = &v
	return s
}

func (s *DowngradePrePayOrderRequestConfluentConfig) SetZooKeeperCU(v int32) *DowngradePrePayOrderRequestConfluentConfig {
	s.ZooKeeperCU = &v
	return s
}

func (s *DowngradePrePayOrderRequestConfluentConfig) SetZooKeeperReplica(v int32) *DowngradePrePayOrderRequestConfluentConfig {
	s.ZooKeeperReplica = &v
	return s
}

func (s *DowngradePrePayOrderRequestConfluentConfig) SetZooKeeperStorage(v int32) *DowngradePrePayOrderRequestConfluentConfig {
	s.ZooKeeperStorage = &v
	return s
}

func (s *DowngradePrePayOrderRequestConfluentConfig) Validate() error {
	return dara.Validate(s)
}
