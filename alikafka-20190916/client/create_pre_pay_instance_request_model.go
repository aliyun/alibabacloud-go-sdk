// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePrePayInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfluentConfig(v *CreatePrePayInstanceRequestConfluentConfig) *CreatePrePayInstanceRequest
	GetConfluentConfig() *CreatePrePayInstanceRequestConfluentConfig
	SetDeployType(v int32) *CreatePrePayInstanceRequest
	GetDeployType() *int32
	SetDiskSize(v int32) *CreatePrePayInstanceRequest
	GetDiskSize() *int32
	SetDiskType(v string) *CreatePrePayInstanceRequest
	GetDiskType() *string
	SetDuration(v int32) *CreatePrePayInstanceRequest
	GetDuration() *int32
	SetEipMax(v int32) *CreatePrePayInstanceRequest
	GetEipMax() *int32
	SetIoMaxSpec(v string) *CreatePrePayInstanceRequest
	GetIoMaxSpec() *string
	SetPaidType(v int32) *CreatePrePayInstanceRequest
	GetPaidType() *int32
	SetPartitionNum(v int32) *CreatePrePayInstanceRequest
	GetPartitionNum() *int32
	SetRegionId(v string) *CreatePrePayInstanceRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreatePrePayInstanceRequest
	GetResourceGroupId() *string
	SetSpecType(v string) *CreatePrePayInstanceRequest
	GetSpecType() *string
	SetTag(v []*CreatePrePayInstanceRequestTag) *CreatePrePayInstanceRequest
	GetTag() []*CreatePrePayInstanceRequestTag
}

type CreatePrePayInstanceRequest struct {
	ConfluentConfig *CreatePrePayInstanceRequestConfluentConfig `json:"ConfluentConfig,omitempty" xml:"ConfluentConfig,omitempty" type:"Struct"`
	// example:
	//
	// 5
	DeployType *int32 `json:"DeployType,omitempty" xml:"DeployType,omitempty"`
	// example:
	//
	// 500
	DiskSize *int32 `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
	// example:
	//
	// 1
	DiskType *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	// example:
	//
	// 1
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// 3
	EipMax *int32 `json:"EipMax,omitempty" xml:"EipMax,omitempty"`
	// example:
	//
	// alikafka.hw.2xlarge
	IoMaxSpec *string `json:"IoMaxSpec,omitempty" xml:"IoMaxSpec,omitempty"`
	// example:
	//
	// 1
	PaidType *int32 `json:"PaidType,omitempty" xml:"PaidType,omitempty"`
	// example:
	//
	// 1000
	PartitionNum *int32 `json:"PartitionNum,omitempty" xml:"PartitionNum,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-ac***********7q
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// professional
	SpecType *string                           `json:"SpecType,omitempty" xml:"SpecType,omitempty"`
	Tag      []*CreatePrePayInstanceRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s CreatePrePayInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePrePayInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreatePrePayInstanceRequest) GetConfluentConfig() *CreatePrePayInstanceRequestConfluentConfig {
	return s.ConfluentConfig
}

func (s *CreatePrePayInstanceRequest) GetDeployType() *int32 {
	return s.DeployType
}

func (s *CreatePrePayInstanceRequest) GetDiskSize() *int32 {
	return s.DiskSize
}

func (s *CreatePrePayInstanceRequest) GetDiskType() *string {
	return s.DiskType
}

func (s *CreatePrePayInstanceRequest) GetDuration() *int32 {
	return s.Duration
}

func (s *CreatePrePayInstanceRequest) GetEipMax() *int32 {
	return s.EipMax
}

func (s *CreatePrePayInstanceRequest) GetIoMaxSpec() *string {
	return s.IoMaxSpec
}

func (s *CreatePrePayInstanceRequest) GetPaidType() *int32 {
	return s.PaidType
}

func (s *CreatePrePayInstanceRequest) GetPartitionNum() *int32 {
	return s.PartitionNum
}

func (s *CreatePrePayInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreatePrePayInstanceRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreatePrePayInstanceRequest) GetSpecType() *string {
	return s.SpecType
}

func (s *CreatePrePayInstanceRequest) GetTag() []*CreatePrePayInstanceRequestTag {
	return s.Tag
}

func (s *CreatePrePayInstanceRequest) SetConfluentConfig(v *CreatePrePayInstanceRequestConfluentConfig) *CreatePrePayInstanceRequest {
	s.ConfluentConfig = v
	return s
}

func (s *CreatePrePayInstanceRequest) SetDeployType(v int32) *CreatePrePayInstanceRequest {
	s.DeployType = &v
	return s
}

func (s *CreatePrePayInstanceRequest) SetDiskSize(v int32) *CreatePrePayInstanceRequest {
	s.DiskSize = &v
	return s
}

func (s *CreatePrePayInstanceRequest) SetDiskType(v string) *CreatePrePayInstanceRequest {
	s.DiskType = &v
	return s
}

func (s *CreatePrePayInstanceRequest) SetDuration(v int32) *CreatePrePayInstanceRequest {
	s.Duration = &v
	return s
}

func (s *CreatePrePayInstanceRequest) SetEipMax(v int32) *CreatePrePayInstanceRequest {
	s.EipMax = &v
	return s
}

func (s *CreatePrePayInstanceRequest) SetIoMaxSpec(v string) *CreatePrePayInstanceRequest {
	s.IoMaxSpec = &v
	return s
}

func (s *CreatePrePayInstanceRequest) SetPaidType(v int32) *CreatePrePayInstanceRequest {
	s.PaidType = &v
	return s
}

func (s *CreatePrePayInstanceRequest) SetPartitionNum(v int32) *CreatePrePayInstanceRequest {
	s.PartitionNum = &v
	return s
}

func (s *CreatePrePayInstanceRequest) SetRegionId(v string) *CreatePrePayInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *CreatePrePayInstanceRequest) SetResourceGroupId(v string) *CreatePrePayInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreatePrePayInstanceRequest) SetSpecType(v string) *CreatePrePayInstanceRequest {
	s.SpecType = &v
	return s
}

func (s *CreatePrePayInstanceRequest) SetTag(v []*CreatePrePayInstanceRequestTag) *CreatePrePayInstanceRequest {
	s.Tag = v
	return s
}

func (s *CreatePrePayInstanceRequest) Validate() error {
	return dara.Validate(s)
}

type CreatePrePayInstanceRequestConfluentConfig struct {
	// example:
	//
	// 4
	ConnectCU *int32 `json:"ConnectCU,omitempty" xml:"ConnectCU,omitempty"`
	// example:
	//
	// 2
	ConnectReplica *int32 `json:"ConnectReplica,omitempty" xml:"ConnectReplica,omitempty"`
	// example:
	//
	// 4
	ControlCenterCU *int32 `json:"ControlCenterCU,omitempty" xml:"ControlCenterCU,omitempty"`
	// example:
	//
	// 1
	ControlCenterReplica *int32 `json:"ControlCenterReplica,omitempty" xml:"ControlCenterReplica,omitempty"`
	// example:
	//
	// 300
	ControlCenterStorage *int32 `json:"ControlCenterStorage,omitempty" xml:"ControlCenterStorage,omitempty"`
	// example:
	//
	// 4
	KafkaCU *int32 `json:"KafkaCU,omitempty" xml:"KafkaCU,omitempty"`
	// example:
	//
	// 3
	KafkaReplica *int32 `json:"KafkaReplica,omitempty" xml:"KafkaReplica,omitempty"`
	// example:
	//
	// 4
	KafkaRestProxyCU *int32 `json:"KafkaRestProxyCU,omitempty" xml:"KafkaRestProxyCU,omitempty"`
	// example:
	//
	// 2
	KafkaRestProxyReplica *int32 `json:"KafkaRestProxyReplica,omitempty" xml:"KafkaRestProxyReplica,omitempty"`
	// example:
	//
	// 800
	KafkaStorage *int32 `json:"KafkaStorage,omitempty" xml:"KafkaStorage,omitempty"`
	// example:
	//
	// 4
	KsqlCU *int32 `json:"KsqlCU,omitempty" xml:"KsqlCU,omitempty"`
	// example:
	//
	// 2
	KsqlReplica *int32 `json:"KsqlReplica,omitempty" xml:"KsqlReplica,omitempty"`
	// example:
	//
	// 100
	KsqlStorage *int32 `json:"KsqlStorage,omitempty" xml:"KsqlStorage,omitempty"`
	// example:
	//
	// 1
	SchemaRegistryCU *int32 `json:"SchemaRegistryCU,omitempty" xml:"SchemaRegistryCU,omitempty"`
	// example:
	//
	// 2
	SchemaRegistryReplica *int32 `json:"SchemaRegistryReplica,omitempty" xml:"SchemaRegistryReplica,omitempty"`
	// example:
	//
	// 2
	ZooKeeperCU *int32 `json:"ZooKeeperCU,omitempty" xml:"ZooKeeperCU,omitempty"`
	// example:
	//
	// 3
	ZooKeeperReplica *int32 `json:"ZooKeeperReplica,omitempty" xml:"ZooKeeperReplica,omitempty"`
	// example:
	//
	// 100
	ZooKeeperStorage *int32 `json:"ZooKeeperStorage,omitempty" xml:"ZooKeeperStorage,omitempty"`
}

func (s CreatePrePayInstanceRequestConfluentConfig) String() string {
	return dara.Prettify(s)
}

func (s CreatePrePayInstanceRequestConfluentConfig) GoString() string {
	return s.String()
}

func (s *CreatePrePayInstanceRequestConfluentConfig) GetConnectCU() *int32 {
	return s.ConnectCU
}

func (s *CreatePrePayInstanceRequestConfluentConfig) GetConnectReplica() *int32 {
	return s.ConnectReplica
}

func (s *CreatePrePayInstanceRequestConfluentConfig) GetControlCenterCU() *int32 {
	return s.ControlCenterCU
}

func (s *CreatePrePayInstanceRequestConfluentConfig) GetControlCenterReplica() *int32 {
	return s.ControlCenterReplica
}

func (s *CreatePrePayInstanceRequestConfluentConfig) GetControlCenterStorage() *int32 {
	return s.ControlCenterStorage
}

func (s *CreatePrePayInstanceRequestConfluentConfig) GetKafkaCU() *int32 {
	return s.KafkaCU
}

func (s *CreatePrePayInstanceRequestConfluentConfig) GetKafkaReplica() *int32 {
	return s.KafkaReplica
}

func (s *CreatePrePayInstanceRequestConfluentConfig) GetKafkaRestProxyCU() *int32 {
	return s.KafkaRestProxyCU
}

func (s *CreatePrePayInstanceRequestConfluentConfig) GetKafkaRestProxyReplica() *int32 {
	return s.KafkaRestProxyReplica
}

func (s *CreatePrePayInstanceRequestConfluentConfig) GetKafkaStorage() *int32 {
	return s.KafkaStorage
}

func (s *CreatePrePayInstanceRequestConfluentConfig) GetKsqlCU() *int32 {
	return s.KsqlCU
}

func (s *CreatePrePayInstanceRequestConfluentConfig) GetKsqlReplica() *int32 {
	return s.KsqlReplica
}

func (s *CreatePrePayInstanceRequestConfluentConfig) GetKsqlStorage() *int32 {
	return s.KsqlStorage
}

func (s *CreatePrePayInstanceRequestConfluentConfig) GetSchemaRegistryCU() *int32 {
	return s.SchemaRegistryCU
}

func (s *CreatePrePayInstanceRequestConfluentConfig) GetSchemaRegistryReplica() *int32 {
	return s.SchemaRegistryReplica
}

func (s *CreatePrePayInstanceRequestConfluentConfig) GetZooKeeperCU() *int32 {
	return s.ZooKeeperCU
}

func (s *CreatePrePayInstanceRequestConfluentConfig) GetZooKeeperReplica() *int32 {
	return s.ZooKeeperReplica
}

func (s *CreatePrePayInstanceRequestConfluentConfig) GetZooKeeperStorage() *int32 {
	return s.ZooKeeperStorage
}

func (s *CreatePrePayInstanceRequestConfluentConfig) SetConnectCU(v int32) *CreatePrePayInstanceRequestConfluentConfig {
	s.ConnectCU = &v
	return s
}

func (s *CreatePrePayInstanceRequestConfluentConfig) SetConnectReplica(v int32) *CreatePrePayInstanceRequestConfluentConfig {
	s.ConnectReplica = &v
	return s
}

func (s *CreatePrePayInstanceRequestConfluentConfig) SetControlCenterCU(v int32) *CreatePrePayInstanceRequestConfluentConfig {
	s.ControlCenterCU = &v
	return s
}

func (s *CreatePrePayInstanceRequestConfluentConfig) SetControlCenterReplica(v int32) *CreatePrePayInstanceRequestConfluentConfig {
	s.ControlCenterReplica = &v
	return s
}

func (s *CreatePrePayInstanceRequestConfluentConfig) SetControlCenterStorage(v int32) *CreatePrePayInstanceRequestConfluentConfig {
	s.ControlCenterStorage = &v
	return s
}

func (s *CreatePrePayInstanceRequestConfluentConfig) SetKafkaCU(v int32) *CreatePrePayInstanceRequestConfluentConfig {
	s.KafkaCU = &v
	return s
}

func (s *CreatePrePayInstanceRequestConfluentConfig) SetKafkaReplica(v int32) *CreatePrePayInstanceRequestConfluentConfig {
	s.KafkaReplica = &v
	return s
}

func (s *CreatePrePayInstanceRequestConfluentConfig) SetKafkaRestProxyCU(v int32) *CreatePrePayInstanceRequestConfluentConfig {
	s.KafkaRestProxyCU = &v
	return s
}

func (s *CreatePrePayInstanceRequestConfluentConfig) SetKafkaRestProxyReplica(v int32) *CreatePrePayInstanceRequestConfluentConfig {
	s.KafkaRestProxyReplica = &v
	return s
}

func (s *CreatePrePayInstanceRequestConfluentConfig) SetKafkaStorage(v int32) *CreatePrePayInstanceRequestConfluentConfig {
	s.KafkaStorage = &v
	return s
}

func (s *CreatePrePayInstanceRequestConfluentConfig) SetKsqlCU(v int32) *CreatePrePayInstanceRequestConfluentConfig {
	s.KsqlCU = &v
	return s
}

func (s *CreatePrePayInstanceRequestConfluentConfig) SetKsqlReplica(v int32) *CreatePrePayInstanceRequestConfluentConfig {
	s.KsqlReplica = &v
	return s
}

func (s *CreatePrePayInstanceRequestConfluentConfig) SetKsqlStorage(v int32) *CreatePrePayInstanceRequestConfluentConfig {
	s.KsqlStorage = &v
	return s
}

func (s *CreatePrePayInstanceRequestConfluentConfig) SetSchemaRegistryCU(v int32) *CreatePrePayInstanceRequestConfluentConfig {
	s.SchemaRegistryCU = &v
	return s
}

func (s *CreatePrePayInstanceRequestConfluentConfig) SetSchemaRegistryReplica(v int32) *CreatePrePayInstanceRequestConfluentConfig {
	s.SchemaRegistryReplica = &v
	return s
}

func (s *CreatePrePayInstanceRequestConfluentConfig) SetZooKeeperCU(v int32) *CreatePrePayInstanceRequestConfluentConfig {
	s.ZooKeeperCU = &v
	return s
}

func (s *CreatePrePayInstanceRequestConfluentConfig) SetZooKeeperReplica(v int32) *CreatePrePayInstanceRequestConfluentConfig {
	s.ZooKeeperReplica = &v
	return s
}

func (s *CreatePrePayInstanceRequestConfluentConfig) SetZooKeeperStorage(v int32) *CreatePrePayInstanceRequestConfluentConfig {
	s.ZooKeeperStorage = &v
	return s
}

func (s *CreatePrePayInstanceRequestConfluentConfig) Validate() error {
	return dara.Validate(s)
}

type CreatePrePayInstanceRequestTag struct {
	// This parameter is required.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreatePrePayInstanceRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreatePrePayInstanceRequestTag) GoString() string {
	return s.String()
}

func (s *CreatePrePayInstanceRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreatePrePayInstanceRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreatePrePayInstanceRequestTag) SetKey(v string) *CreatePrePayInstanceRequestTag {
	s.Key = &v
	return s
}

func (s *CreatePrePayInstanceRequestTag) SetValue(v string) *CreatePrePayInstanceRequestTag {
	s.Value = &v
	return s
}

func (s *CreatePrePayInstanceRequestTag) Validate() error {
	return dara.Validate(s)
}
