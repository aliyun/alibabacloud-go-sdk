// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoRenew(v bool) *CreateInstanceRequest
	GetAutoRenew() *bool
	SetAutoRenewPeriod(v int32) *CreateInstanceRequest
	GetAutoRenewPeriod() *int32
	SetClientToken(v string) *CreateInstanceRequest
	GetClientToken() *string
	SetEdition(v string) *CreateInstanceRequest
	GetEdition() *string
	SetEncryptedInstance(v bool) *CreateInstanceRequest
	GetEncryptedInstance() *bool
	SetInstanceName(v string) *CreateInstanceRequest
	GetInstanceName() *string
	SetInstanceType(v string) *CreateInstanceRequest
	GetInstanceType() *string
	SetKmsKeyId(v string) *CreateInstanceRequest
	GetKmsKeyId() *string
	SetMaxConnections(v int32) *CreateInstanceRequest
	GetMaxConnections() *int32
	SetMaxEipTps(v int64) *CreateInstanceRequest
	GetMaxEipTps() *int64
	SetMaxPrivateTps(v int64) *CreateInstanceRequest
	GetMaxPrivateTps() *int64
	SetPaymentType(v string) *CreateInstanceRequest
	GetPaymentType() *string
	SetPeriod(v int32) *CreateInstanceRequest
	GetPeriod() *int32
	SetPeriodCycle(v string) *CreateInstanceRequest
	GetPeriodCycle() *string
	SetProvisionedCapacity(v int32) *CreateInstanceRequest
	GetProvisionedCapacity() *int32
	SetQueueCapacity(v int32) *CreateInstanceRequest
	GetQueueCapacity() *int32
	SetRenewStatus(v string) *CreateInstanceRequest
	GetRenewStatus() *string
	SetRenewalDurationUnit(v string) *CreateInstanceRequest
	GetRenewalDurationUnit() *string
	SetResourceGroupId(v string) *CreateInstanceRequest
	GetResourceGroupId() *string
	SetServerlessChargeType(v string) *CreateInstanceRequest
	GetServerlessChargeType() *string
	SetStorageSize(v int32) *CreateInstanceRequest
	GetStorageSize() *int32
	SetSupportEip(v bool) *CreateInstanceRequest
	GetSupportEip() *bool
	SetSupportTracing(v bool) *CreateInstanceRequest
	GetSupportTracing() *bool
	SetTracingStorageTime(v int32) *CreateInstanceRequest
	GetTracingStorageTime() *int32
}

type CreateInstanceRequest struct {
	// Specifies whether to enable auto-renewal. Valid values:
	//
	// 	- true: enables auto-renewal.
	//
	// 	- false: disables auto-renewal. If you select this value, you must manually renew the instance.
	//
	// example:
	//
	// AutoRenewal
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The auto-renewal period. The unit of the auto-renewal period is specified by RenewalDurationUnit. Default value: Month.
	//
	// >  This parameter takes effect only if you set AutoRenew to true. Default value: 1.
	//
	// example:
	//
	// 1
	AutoRenewPeriod *int32 `json:"AutoRenewPeriod,omitempty" xml:"AutoRenewPeriod,omitempty"`
	// The client token.
	//
	// example:
	//
	// c2c5d1274axxxxxxxx
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Edition     *string `json:"Edition,omitempty" xml:"Edition,omitempty"`
	// Specifies whether to enable storage encryption for the instance. This parameter is available only for exclusive instances.
	//
	// example:
	//
	// false
	EncryptedInstance *bool `json:"EncryptedInstance,omitempty" xml:"EncryptedInstance,omitempty"`
	// The name of the instance. We recommend that you specify a name that does not exceed 64 characters in length.
	//
	// example:
	//
	// amqp-xxxxx
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The instance edition. Valid values if you create a subscription instance:
	//
	// 	- professional: Professional Edition.
	//
	// 	- enterprise: Enterprise Edition
	//
	// 	- vip: Enterprise Platinum Edition
	//
	// If you create a serverless instance, you do not need to specify this parameter.
	//
	// example:
	//
	// professional
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The ID of the Key Management Service (KMS)-managed key used for storage encryption. This parameter is available only for exclusive instances and required only if you set EncryptedInstance to true. The key must meet the following conditions:
	//
	// 	- The key cannot be a service key.
	//
	// 	- The key must be in the Enabled state.
	//
	// 	- The key must be a symmetric key.
	//
	// 	- The key must be used for encryption and decryption.
	//
	// 	- After the key is expired or deleted, you cannot read or write data and exceptions can occur in the ApsaraMQ for RabbitMQ instance.
	//
	// example:
	//
	// key-xxx
	KmsKeyId *string `json:"KmsKeyId,omitempty" xml:"KmsKeyId,omitempty"`
	// The maximum number of connections that can be established to the instance.
	//
	// Configure this parameter based on the values provided on the [ApsaraMQ for RocketMQ buy page](https://common-buy.aliyun.com/?commodityCode=ons_onsproxy_pre).
	//
	// example:
	//
	// 50000
	MaxConnections *int32 `json:"MaxConnections,omitempty" xml:"MaxConnections,omitempty"`
	// The maximum number of Internet-based TPS on the instance.
	//
	// Configure this parameter based on the values provided on the [ApsaraMQ for RocketMQ buy page](https://common-buy.aliyun.com/?commodityCode=ons_onsproxy_pre).
	//
	// example:
	//
	// 128
	MaxEipTps *int64 `json:"MaxEipTps,omitempty" xml:"MaxEipTps,omitempty"`
	// The maximum number of virtual private cloud (VPC)-based transactions per second (TPS) on the instance.
	//
	// Configure this parameter based on the values provided on the [ApsaraMQ for RocketMQ buy page](https://common-buy.aliyun.com/?commodityCode=ons_onsproxy_pre).
	//
	// example:
	//
	// 1000
	MaxPrivateTps *int64 `json:"MaxPrivateTps,omitempty" xml:"MaxPrivateTps,omitempty"`
	// The billing method of the instance. Valid values:
	//
	// 	- Subscription: subscription instance
	//
	// 	- PayAsYouGo: serverless instance
	//
	// This parameter is required.
	//
	// example:
	//
	// Subscription
	PaymentType *string `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
	// The subscription period. The unit of the subscription period is specified by periodCycle.
	//
	// >  This parameter takes effect only if you set PaymentType to Subscription. Default value: 1.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The unit of the subscription period. Valid values:
	//
	// 	- Month
	//
	// 	- Year
	//
	// This parameter is valid only if you set PaymentType to Subscription. Default value: Month.
	//
	// example:
	//
	// Month
	PeriodCycle         *string `json:"PeriodCycle,omitempty" xml:"PeriodCycle,omitempty"`
	ProvisionedCapacity *int32  `json:"ProvisionedCapacity,omitempty" xml:"ProvisionedCapacity,omitempty"`
	// The number of queues on the instance.
	//
	// Configure this parameter based on the values provided on the [ApsaraMQ for RocketMQ buy page](https://common-buy.aliyun.com/?commodityCode=ons_onsproxy_pre).
	//
	// example:
	//
	// 1000
	QueueCapacity *int32 `json:"QueueCapacity,omitempty" xml:"QueueCapacity,omitempty"`
	// The renewal status. This parameter is the same as AutoRenew. You can configure one of these parameters. Valid value:
	//
	// 	- AutoRenewal
	//
	// >  If you configure both this parameter and AutoRenew, the value of this parameter is used.
	//
	// example:
	//
	// false
	RenewStatus *string `json:"RenewStatus,omitempty" xml:"RenewStatus,omitempty"`
	// The unit of the auto-renewal period. Valid values:
	//
	// 	- Month
	//
	// 	- Year
	//
	// example:
	//
	// Month
	RenewalDurationUnit *string `json:"RenewalDurationUnit,omitempty" xml:"RenewalDurationUnit,omitempty"`
	// The ID of the resource group to which the instance belongs.
	//
	// example:
	//
	// rg-acfmvvajg5qkxhi
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The billing method of the serverless instance. Valid value:
	//
	// 	- onDemand: You are charged based on your actual usage.
	//
	// example:
	//
	// onDemand
	ServerlessChargeType *string `json:"ServerlessChargeType,omitempty" xml:"ServerlessChargeType,omitempty"`
	// The storage capacity. Unit: GB. Valid values:
	//
	// 	- Professional Edition and Enterprise Edition instances: Set the value to 0.
	//
	// >  The value 0 specifies that storage space is available for Professional Edition and Enterprise Edition instances, but no storage fees are generated.
	//
	// 	- Enterprise Platinum Edition instances: Set the value to m × 100, where m is an integer that ranges from 7 to 28.
	//
	// example:
	//
	// 7
	StorageSize *int32 `json:"StorageSize,omitempty" xml:"StorageSize,omitempty"`
	// Specifies whether elastic IP addresses (EIPs) are supported. Valid values:
	//
	// 	- True
	//
	// 	- False
	//
	// example:
	//
	// true
	SupportEip *bool `json:"SupportEip,omitempty" xml:"SupportEip,omitempty"`
	// Specifies whether to enable the message trace feature. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// >
	//
	// 	- Enterprise Platinum Edition instances allow you to retain message traces for 15 days free of charge. If you create an Enterprise Platinum Edition instance, you can set this parameter only to true and TracingStorageTime only to 15.
	//
	// 	- For instances of other editions, you can set this parameter to true or false.
	//
	// example:
	//
	// true
	SupportTracing *bool `json:"SupportTracing,omitempty" xml:"SupportTracing,omitempty"`
	// The retention period of messages. Unit: days. Valid values:
	//
	// 	- 3
	//
	// 	- 7
	//
	// 	- 15
	//
	// This parameter is valid only if you set SupportTracing to true.
	//
	// example:
	//
	// 3
	TracingStorageTime *int32 `json:"TracingStorageTime,omitempty" xml:"TracingStorageTime,omitempty"`
}

func (s CreateInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *CreateInstanceRequest) GetAutoRenewPeriod() *int32 {
	return s.AutoRenewPeriod
}

func (s *CreateInstanceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateInstanceRequest) GetEdition() *string {
	return s.Edition
}

func (s *CreateInstanceRequest) GetEncryptedInstance() *bool {
	return s.EncryptedInstance
}

func (s *CreateInstanceRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *CreateInstanceRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *CreateInstanceRequest) GetKmsKeyId() *string {
	return s.KmsKeyId
}

func (s *CreateInstanceRequest) GetMaxConnections() *int32 {
	return s.MaxConnections
}

func (s *CreateInstanceRequest) GetMaxEipTps() *int64 {
	return s.MaxEipTps
}

func (s *CreateInstanceRequest) GetMaxPrivateTps() *int64 {
	return s.MaxPrivateTps
}

func (s *CreateInstanceRequest) GetPaymentType() *string {
	return s.PaymentType
}

func (s *CreateInstanceRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *CreateInstanceRequest) GetPeriodCycle() *string {
	return s.PeriodCycle
}

func (s *CreateInstanceRequest) GetProvisionedCapacity() *int32 {
	return s.ProvisionedCapacity
}

func (s *CreateInstanceRequest) GetQueueCapacity() *int32 {
	return s.QueueCapacity
}

func (s *CreateInstanceRequest) GetRenewStatus() *string {
	return s.RenewStatus
}

func (s *CreateInstanceRequest) GetRenewalDurationUnit() *string {
	return s.RenewalDurationUnit
}

func (s *CreateInstanceRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateInstanceRequest) GetServerlessChargeType() *string {
	return s.ServerlessChargeType
}

func (s *CreateInstanceRequest) GetStorageSize() *int32 {
	return s.StorageSize
}

func (s *CreateInstanceRequest) GetSupportEip() *bool {
	return s.SupportEip
}

func (s *CreateInstanceRequest) GetSupportTracing() *bool {
	return s.SupportTracing
}

func (s *CreateInstanceRequest) GetTracingStorageTime() *int32 {
	return s.TracingStorageTime
}

func (s *CreateInstanceRequest) SetAutoRenew(v bool) *CreateInstanceRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateInstanceRequest) SetAutoRenewPeriod(v int32) *CreateInstanceRequest {
	s.AutoRenewPeriod = &v
	return s
}

func (s *CreateInstanceRequest) SetClientToken(v string) *CreateInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateInstanceRequest) SetEdition(v string) *CreateInstanceRequest {
	s.Edition = &v
	return s
}

func (s *CreateInstanceRequest) SetEncryptedInstance(v bool) *CreateInstanceRequest {
	s.EncryptedInstance = &v
	return s
}

func (s *CreateInstanceRequest) SetInstanceName(v string) *CreateInstanceRequest {
	s.InstanceName = &v
	return s
}

func (s *CreateInstanceRequest) SetInstanceType(v string) *CreateInstanceRequest {
	s.InstanceType = &v
	return s
}

func (s *CreateInstanceRequest) SetKmsKeyId(v string) *CreateInstanceRequest {
	s.KmsKeyId = &v
	return s
}

func (s *CreateInstanceRequest) SetMaxConnections(v int32) *CreateInstanceRequest {
	s.MaxConnections = &v
	return s
}

func (s *CreateInstanceRequest) SetMaxEipTps(v int64) *CreateInstanceRequest {
	s.MaxEipTps = &v
	return s
}

func (s *CreateInstanceRequest) SetMaxPrivateTps(v int64) *CreateInstanceRequest {
	s.MaxPrivateTps = &v
	return s
}

func (s *CreateInstanceRequest) SetPaymentType(v string) *CreateInstanceRequest {
	s.PaymentType = &v
	return s
}

func (s *CreateInstanceRequest) SetPeriod(v int32) *CreateInstanceRequest {
	s.Period = &v
	return s
}

func (s *CreateInstanceRequest) SetPeriodCycle(v string) *CreateInstanceRequest {
	s.PeriodCycle = &v
	return s
}

func (s *CreateInstanceRequest) SetProvisionedCapacity(v int32) *CreateInstanceRequest {
	s.ProvisionedCapacity = &v
	return s
}

func (s *CreateInstanceRequest) SetQueueCapacity(v int32) *CreateInstanceRequest {
	s.QueueCapacity = &v
	return s
}

func (s *CreateInstanceRequest) SetRenewStatus(v string) *CreateInstanceRequest {
	s.RenewStatus = &v
	return s
}

func (s *CreateInstanceRequest) SetRenewalDurationUnit(v string) *CreateInstanceRequest {
	s.RenewalDurationUnit = &v
	return s
}

func (s *CreateInstanceRequest) SetResourceGroupId(v string) *CreateInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateInstanceRequest) SetServerlessChargeType(v string) *CreateInstanceRequest {
	s.ServerlessChargeType = &v
	return s
}

func (s *CreateInstanceRequest) SetStorageSize(v int32) *CreateInstanceRequest {
	s.StorageSize = &v
	return s
}

func (s *CreateInstanceRequest) SetSupportEip(v bool) *CreateInstanceRequest {
	s.SupportEip = &v
	return s
}

func (s *CreateInstanceRequest) SetSupportTracing(v bool) *CreateInstanceRequest {
	s.SupportTracing = &v
	return s
}

func (s *CreateInstanceRequest) SetTracingStorageTime(v int32) *CreateInstanceRequest {
	s.TracingStorageTime = &v
	return s
}

func (s *CreateInstanceRequest) Validate() error {
	return dara.Validate(s)
}
