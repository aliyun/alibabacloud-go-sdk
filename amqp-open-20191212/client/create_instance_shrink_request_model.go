// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstanceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoRenew(v bool) *CreateInstanceShrinkRequest
	GetAutoRenew() *bool
	SetAutoRenewPeriod(v int32) *CreateInstanceShrinkRequest
	GetAutoRenewPeriod() *int32
	SetClientToken(v string) *CreateInstanceShrinkRequest
	GetClientToken() *string
	SetEdition(v string) *CreateInstanceShrinkRequest
	GetEdition() *string
	SetEncryptedInstance(v bool) *CreateInstanceShrinkRequest
	GetEncryptedInstance() *bool
	SetInstanceName(v string) *CreateInstanceShrinkRequest
	GetInstanceName() *string
	SetInstanceType(v string) *CreateInstanceShrinkRequest
	GetInstanceType() *string
	SetKmsKeyId(v string) *CreateInstanceShrinkRequest
	GetKmsKeyId() *string
	SetMaxConnections(v int32) *CreateInstanceShrinkRequest
	GetMaxConnections() *int32
	SetMaxEipTps(v int64) *CreateInstanceShrinkRequest
	GetMaxEipTps() *int64
	SetMaxPrivateTps(v int64) *CreateInstanceShrinkRequest
	GetMaxPrivateTps() *int64
	SetPaymentType(v string) *CreateInstanceShrinkRequest
	GetPaymentType() *string
	SetPeriod(v int32) *CreateInstanceShrinkRequest
	GetPeriod() *int32
	SetPeriodCycle(v string) *CreateInstanceShrinkRequest
	GetPeriodCycle() *string
	SetProvisionedCapacity(v int32) *CreateInstanceShrinkRequest
	GetProvisionedCapacity() *int32
	SetQueueCapacity(v int32) *CreateInstanceShrinkRequest
	GetQueueCapacity() *int32
	SetRenewStatus(v string) *CreateInstanceShrinkRequest
	GetRenewStatus() *string
	SetRenewalDurationUnit(v string) *CreateInstanceShrinkRequest
	GetRenewalDurationUnit() *string
	SetResourceGroupId(v string) *CreateInstanceShrinkRequest
	GetResourceGroupId() *string
	SetServerlessChargeType(v string) *CreateInstanceShrinkRequest
	GetServerlessChargeType() *string
	SetStorageSize(v int32) *CreateInstanceShrinkRequest
	GetStorageSize() *int32
	SetSupportEip(v bool) *CreateInstanceShrinkRequest
	GetSupportEip() *bool
	SetSupportTracing(v bool) *CreateInstanceShrinkRequest
	GetSupportTracing() *bool
	SetTagsShrink(v string) *CreateInstanceShrinkRequest
	GetTagsShrink() *string
	SetTracingStorageTime(v int32) *CreateInstanceShrinkRequest
	GetTracingStorageTime() *int32
}

type CreateInstanceShrinkRequest struct {
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
	SupportTracing *bool   `json:"SupportTracing,omitempty" xml:"SupportTracing,omitempty"`
	TagsShrink     *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
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

func (s CreateInstanceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateInstanceShrinkRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *CreateInstanceShrinkRequest) GetAutoRenewPeriod() *int32 {
	return s.AutoRenewPeriod
}

func (s *CreateInstanceShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateInstanceShrinkRequest) GetEdition() *string {
	return s.Edition
}

func (s *CreateInstanceShrinkRequest) GetEncryptedInstance() *bool {
	return s.EncryptedInstance
}

func (s *CreateInstanceShrinkRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *CreateInstanceShrinkRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *CreateInstanceShrinkRequest) GetKmsKeyId() *string {
	return s.KmsKeyId
}

func (s *CreateInstanceShrinkRequest) GetMaxConnections() *int32 {
	return s.MaxConnections
}

func (s *CreateInstanceShrinkRequest) GetMaxEipTps() *int64 {
	return s.MaxEipTps
}

func (s *CreateInstanceShrinkRequest) GetMaxPrivateTps() *int64 {
	return s.MaxPrivateTps
}

func (s *CreateInstanceShrinkRequest) GetPaymentType() *string {
	return s.PaymentType
}

func (s *CreateInstanceShrinkRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *CreateInstanceShrinkRequest) GetPeriodCycle() *string {
	return s.PeriodCycle
}

func (s *CreateInstanceShrinkRequest) GetProvisionedCapacity() *int32 {
	return s.ProvisionedCapacity
}

func (s *CreateInstanceShrinkRequest) GetQueueCapacity() *int32 {
	return s.QueueCapacity
}

func (s *CreateInstanceShrinkRequest) GetRenewStatus() *string {
	return s.RenewStatus
}

func (s *CreateInstanceShrinkRequest) GetRenewalDurationUnit() *string {
	return s.RenewalDurationUnit
}

func (s *CreateInstanceShrinkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateInstanceShrinkRequest) GetServerlessChargeType() *string {
	return s.ServerlessChargeType
}

func (s *CreateInstanceShrinkRequest) GetStorageSize() *int32 {
	return s.StorageSize
}

func (s *CreateInstanceShrinkRequest) GetSupportEip() *bool {
	return s.SupportEip
}

func (s *CreateInstanceShrinkRequest) GetSupportTracing() *bool {
	return s.SupportTracing
}

func (s *CreateInstanceShrinkRequest) GetTagsShrink() *string {
	return s.TagsShrink
}

func (s *CreateInstanceShrinkRequest) GetTracingStorageTime() *int32 {
	return s.TracingStorageTime
}

func (s *CreateInstanceShrinkRequest) SetAutoRenew(v bool) *CreateInstanceShrinkRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateInstanceShrinkRequest) SetAutoRenewPeriod(v int32) *CreateInstanceShrinkRequest {
	s.AutoRenewPeriod = &v
	return s
}

func (s *CreateInstanceShrinkRequest) SetClientToken(v string) *CreateInstanceShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateInstanceShrinkRequest) SetEdition(v string) *CreateInstanceShrinkRequest {
	s.Edition = &v
	return s
}

func (s *CreateInstanceShrinkRequest) SetEncryptedInstance(v bool) *CreateInstanceShrinkRequest {
	s.EncryptedInstance = &v
	return s
}

func (s *CreateInstanceShrinkRequest) SetInstanceName(v string) *CreateInstanceShrinkRequest {
	s.InstanceName = &v
	return s
}

func (s *CreateInstanceShrinkRequest) SetInstanceType(v string) *CreateInstanceShrinkRequest {
	s.InstanceType = &v
	return s
}

func (s *CreateInstanceShrinkRequest) SetKmsKeyId(v string) *CreateInstanceShrinkRequest {
	s.KmsKeyId = &v
	return s
}

func (s *CreateInstanceShrinkRequest) SetMaxConnections(v int32) *CreateInstanceShrinkRequest {
	s.MaxConnections = &v
	return s
}

func (s *CreateInstanceShrinkRequest) SetMaxEipTps(v int64) *CreateInstanceShrinkRequest {
	s.MaxEipTps = &v
	return s
}

func (s *CreateInstanceShrinkRequest) SetMaxPrivateTps(v int64) *CreateInstanceShrinkRequest {
	s.MaxPrivateTps = &v
	return s
}

func (s *CreateInstanceShrinkRequest) SetPaymentType(v string) *CreateInstanceShrinkRequest {
	s.PaymentType = &v
	return s
}

func (s *CreateInstanceShrinkRequest) SetPeriod(v int32) *CreateInstanceShrinkRequest {
	s.Period = &v
	return s
}

func (s *CreateInstanceShrinkRequest) SetPeriodCycle(v string) *CreateInstanceShrinkRequest {
	s.PeriodCycle = &v
	return s
}

func (s *CreateInstanceShrinkRequest) SetProvisionedCapacity(v int32) *CreateInstanceShrinkRequest {
	s.ProvisionedCapacity = &v
	return s
}

func (s *CreateInstanceShrinkRequest) SetQueueCapacity(v int32) *CreateInstanceShrinkRequest {
	s.QueueCapacity = &v
	return s
}

func (s *CreateInstanceShrinkRequest) SetRenewStatus(v string) *CreateInstanceShrinkRequest {
	s.RenewStatus = &v
	return s
}

func (s *CreateInstanceShrinkRequest) SetRenewalDurationUnit(v string) *CreateInstanceShrinkRequest {
	s.RenewalDurationUnit = &v
	return s
}

func (s *CreateInstanceShrinkRequest) SetResourceGroupId(v string) *CreateInstanceShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateInstanceShrinkRequest) SetServerlessChargeType(v string) *CreateInstanceShrinkRequest {
	s.ServerlessChargeType = &v
	return s
}

func (s *CreateInstanceShrinkRequest) SetStorageSize(v int32) *CreateInstanceShrinkRequest {
	s.StorageSize = &v
	return s
}

func (s *CreateInstanceShrinkRequest) SetSupportEip(v bool) *CreateInstanceShrinkRequest {
	s.SupportEip = &v
	return s
}

func (s *CreateInstanceShrinkRequest) SetSupportTracing(v bool) *CreateInstanceShrinkRequest {
	s.SupportTracing = &v
	return s
}

func (s *CreateInstanceShrinkRequest) SetTagsShrink(v string) *CreateInstanceShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *CreateInstanceShrinkRequest) SetTracingStorageTime(v int32) *CreateInstanceShrinkRequest {
	s.TracingStorageTime = &v
	return s
}

func (s *CreateInstanceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
