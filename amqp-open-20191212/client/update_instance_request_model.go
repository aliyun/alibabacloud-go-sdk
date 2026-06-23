// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateInstanceRequest
	GetClientToken() *string
	SetEdition(v string) *UpdateInstanceRequest
	GetEdition() *string
	SetEncryptedInstance(v bool) *UpdateInstanceRequest
	GetEncryptedInstance() *bool
	SetInstanceId(v string) *UpdateInstanceRequest
	GetInstanceId() *string
	SetInstanceType(v string) *UpdateInstanceRequest
	GetInstanceType() *string
	SetKmsKeyId(v string) *UpdateInstanceRequest
	GetKmsKeyId() *string
	SetMaxConnections(v int32) *UpdateInstanceRequest
	GetMaxConnections() *int32
	SetMaxEipTps(v int64) *UpdateInstanceRequest
	GetMaxEipTps() *int64
	SetMaxPrivateTps(v int64) *UpdateInstanceRequest
	GetMaxPrivateTps() *int64
	SetModifyType(v string) *UpdateInstanceRequest
	GetModifyType() *string
	SetProvisionedCapacity(v int32) *UpdateInstanceRequest
	GetProvisionedCapacity() *int32
	SetQueueCapacity(v int32) *UpdateInstanceRequest
	GetQueueCapacity() *int32
	SetServerlessChargeType(v string) *UpdateInstanceRequest
	GetServerlessChargeType() *string
	SetStorageSize(v int32) *UpdateInstanceRequest
	GetStorageSize() *int32
	SetSupportEip(v bool) *UpdateInstanceRequest
	GetSupportEip() *bool
	SetSupportTracing(v bool) *UpdateInstanceRequest
	GetSupportTracing() *bool
	SetTracingStorageTime(v int32) *UpdateInstanceRequest
	GetTracingStorageTime() *int32
}

type UpdateInstanceRequest struct {
	// The client token.
	//
	// example:
	//
	// c2c5d1274axxxxxxxx
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The deployment architecture of the Serverless instance. Valid values:
	//
	// - shared: A shared architecture. This applies to reserved plus elastic (shared) and pay-as-you-go instances.
	//
	// - dedicated: A dedicated architecture. This applies to reserved plus elastic (dedicated) instances.
	//
	// example:
	//
	// shared
	Edition *string `json:"Edition,omitempty" xml:"Edition,omitempty"`
	// This feature is for dedicated instances only. Specifies whether to enable data encryption.
	//
	// - You cannot change the EncryptedInstance and KmsKeyId properties of a dedicated instance. This includes changing its encryption status or downgrading it to a shared instance. Do not include the EncryptedInstance and KmsKeyId parameters when you call UpdateInstance to upgrade or downgrade a dedicated instance.
	//
	// - The EncryptedInstance and KmsKeyId parameters are used only when you upgrade a shared instance to an encrypted dedicated instance.
	//
	// example:
	//
	// false
	EncryptedInstance *bool `json:"EncryptedInstance,omitempty" xml:"EncryptedInstance,omitempty"`
	// The ID of the ApsaraMQ for RabbitMQ instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// amqp-cn-jtexxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The instance type. This parameter is required for subscription instances. Valid values:
	//
	// - professional: Professional Edition
	//
	// - enterprise: Enterprise Edition
	//
	// - vip: Platinum Edition
	//
	// You do not need to specify this parameter for pay-as-you-go instances.
	//
	// example:
	//
	// professional
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// This feature is for dedicated instances only. This parameter is required if EncryptedInstance is set to true.
	//
	// It specifies the ID of the KMS key used for data encryption.
	//
	// The key must meet the following requirements:
	//
	// - The KMS key must be in the same region as the ApsaraMQ for RabbitMQ instance.
	//
	// - The key cannot be a service key.
	//
	// - The key must be active.
	//
	// - The key must be a symmetric key.
	//
	// - The key must be used for encryption and decryption.
	//
	// - If the KMS key expires or is deleted, data reads and writes will fail, and the ApsaraMQ for RabbitMQ instance will become unavailable.
	//
	// example:
	//
	// key-bjj66c2a893vmhawtq5fd
	KmsKeyId *string `json:"KmsKeyId,omitempty" xml:"KmsKeyId,omitempty"`
	// The maximum number of connections.
	//
	// example:
	//
	// 1000
	MaxConnections *int32 `json:"MaxConnections,omitempty" xml:"MaxConnections,omitempty"`
	// The peak TPS for public network traffic.
	//
	// example:
	//
	// 128
	MaxEipTps *int64 `json:"MaxEipTps,omitempty" xml:"MaxEipTps,omitempty"`
	// The peak transactions per second (TPS) for private network traffic.
	//
	// example:
	//
	// 1000
	MaxPrivateTps *int64 `json:"MaxPrivateTps,omitempty" xml:"MaxPrivateTps,omitempty"`
	// The type of specification change. Valid values:
	//
	// - UPGRADE: Upgrade
	//
	// - DOWNGRADE: Downgrade
	//
	// This parameter is required.
	//
	// example:
	//
	// UPGRADE
	ModifyType *string `json:"ModifyType,omitempty" xml:"ModifyType,omitempty"`
	// The provisioned TPS capacity for a reserved plus elastic instance.
	//
	// example:
	//
	// 2000
	ProvisionedCapacity *int32 `json:"ProvisionedCapacity,omitempty" xml:"ProvisionedCapacity,omitempty"`
	// The maximum number of queues.
	//
	// example:
	//
	// 1000
	QueueCapacity *int32 `json:"QueueCapacity,omitempty" xml:"QueueCapacity,omitempty"`
	// The billing method of the pay-as-you-go (Serverless) instance. Valid value:
	//
	// - onDemand: Pay-as-you-go
	//
	// example:
	//
	// onDemand
	ServerlessChargeType *string `json:"ServerlessChargeType,omitempty" xml:"ServerlessChargeType,omitempty"`
	// The message storage capacity. Unit: GB.
	//
	// example:
	//
	// 7
	StorageSize *int32 `json:"StorageSize,omitempty" xml:"StorageSize,omitempty"`
	// Specifies whether to enable Internet access.
	//
	// example:
	//
	// false
	SupportEip *bool `json:"SupportEip,omitempty" xml:"SupportEip,omitempty"`
	// Specifies whether to enable the message trace feature.
	//
	// example:
	//
	// false
	SupportTracing *bool `json:"SupportTracing,omitempty" xml:"SupportTracing,omitempty"`
	// The retention period for message traces. Unit: days.
	//
	// example:
	//
	// 3
	TracingStorageTime *int32 `json:"TracingStorageTime,omitempty" xml:"TracingStorageTime,omitempty"`
}

func (s UpdateInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceRequest) GoString() string {
	return s.String()
}

func (s *UpdateInstanceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateInstanceRequest) GetEdition() *string {
	return s.Edition
}

func (s *UpdateInstanceRequest) GetEncryptedInstance() *bool {
	return s.EncryptedInstance
}

func (s *UpdateInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateInstanceRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *UpdateInstanceRequest) GetKmsKeyId() *string {
	return s.KmsKeyId
}

func (s *UpdateInstanceRequest) GetMaxConnections() *int32 {
	return s.MaxConnections
}

func (s *UpdateInstanceRequest) GetMaxEipTps() *int64 {
	return s.MaxEipTps
}

func (s *UpdateInstanceRequest) GetMaxPrivateTps() *int64 {
	return s.MaxPrivateTps
}

func (s *UpdateInstanceRequest) GetModifyType() *string {
	return s.ModifyType
}

func (s *UpdateInstanceRequest) GetProvisionedCapacity() *int32 {
	return s.ProvisionedCapacity
}

func (s *UpdateInstanceRequest) GetQueueCapacity() *int32 {
	return s.QueueCapacity
}

func (s *UpdateInstanceRequest) GetServerlessChargeType() *string {
	return s.ServerlessChargeType
}

func (s *UpdateInstanceRequest) GetStorageSize() *int32 {
	return s.StorageSize
}

func (s *UpdateInstanceRequest) GetSupportEip() *bool {
	return s.SupportEip
}

func (s *UpdateInstanceRequest) GetSupportTracing() *bool {
	return s.SupportTracing
}

func (s *UpdateInstanceRequest) GetTracingStorageTime() *int32 {
	return s.TracingStorageTime
}

func (s *UpdateInstanceRequest) SetClientToken(v string) *UpdateInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateInstanceRequest) SetEdition(v string) *UpdateInstanceRequest {
	s.Edition = &v
	return s
}

func (s *UpdateInstanceRequest) SetEncryptedInstance(v bool) *UpdateInstanceRequest {
	s.EncryptedInstance = &v
	return s
}

func (s *UpdateInstanceRequest) SetInstanceId(v string) *UpdateInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateInstanceRequest) SetInstanceType(v string) *UpdateInstanceRequest {
	s.InstanceType = &v
	return s
}

func (s *UpdateInstanceRequest) SetKmsKeyId(v string) *UpdateInstanceRequest {
	s.KmsKeyId = &v
	return s
}

func (s *UpdateInstanceRequest) SetMaxConnections(v int32) *UpdateInstanceRequest {
	s.MaxConnections = &v
	return s
}

func (s *UpdateInstanceRequest) SetMaxEipTps(v int64) *UpdateInstanceRequest {
	s.MaxEipTps = &v
	return s
}

func (s *UpdateInstanceRequest) SetMaxPrivateTps(v int64) *UpdateInstanceRequest {
	s.MaxPrivateTps = &v
	return s
}

func (s *UpdateInstanceRequest) SetModifyType(v string) *UpdateInstanceRequest {
	s.ModifyType = &v
	return s
}

func (s *UpdateInstanceRequest) SetProvisionedCapacity(v int32) *UpdateInstanceRequest {
	s.ProvisionedCapacity = &v
	return s
}

func (s *UpdateInstanceRequest) SetQueueCapacity(v int32) *UpdateInstanceRequest {
	s.QueueCapacity = &v
	return s
}

func (s *UpdateInstanceRequest) SetServerlessChargeType(v string) *UpdateInstanceRequest {
	s.ServerlessChargeType = &v
	return s
}

func (s *UpdateInstanceRequest) SetStorageSize(v int32) *UpdateInstanceRequest {
	s.StorageSize = &v
	return s
}

func (s *UpdateInstanceRequest) SetSupportEip(v bool) *UpdateInstanceRequest {
	s.SupportEip = &v
	return s
}

func (s *UpdateInstanceRequest) SetSupportTracing(v bool) *UpdateInstanceRequest {
	s.SupportTracing = &v
	return s
}

func (s *UpdateInstanceRequest) SetTracingStorageTime(v int32) *UpdateInstanceRequest {
	s.TracingStorageTime = &v
	return s
}

func (s *UpdateInstanceRequest) Validate() error {
	return dara.Validate(s)
}
