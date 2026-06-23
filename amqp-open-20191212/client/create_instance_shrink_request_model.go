// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstanceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthModel(v string) *CreateInstanceShrinkRequest
	GetAuthModel() *string
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
	SetListenerMode(v string) *CreateInstanceShrinkRequest
	GetListenerMode() *string
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
	SetSecurityGroupId(v string) *CreateInstanceShrinkRequest
	GetSecurityGroupId() *string
	SetServerlessChargeType(v string) *CreateInstanceShrinkRequest
	GetServerlessChargeType() *string
	SetServerlessSwitch(v bool) *CreateInstanceShrinkRequest
	GetServerlessSwitch() *bool
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
	SetVpcId(v string) *CreateInstanceShrinkRequest
	GetVpcId() *string
	SetVswitchIdsShrink(v string) *CreateInstanceShrinkRequest
	GetVswitchIdsShrink() *string
}

type CreateInstanceShrinkRequest struct {
	// example:
	//
	// ram openSource
	AuthModel *string `json:"AuthModel,omitempty" xml:"AuthModel,omitempty"`
	// The renewal method. Valid values:
	//
	// - `true`: Enables auto-renewal.
	//
	// - `false`: Disables auto-renewal.
	//
	// example:
	//
	// false
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The auto-renewal duration. The `RenewalDurationUnit` parameter specifies the unit, which defaults to month.
	//
	// > This parameter is required if you set `AutoRenew` to `true`. The default value is 1.
	//
	// example:
	//
	// 1
	AutoRenewPeriod *int32 `json:"AutoRenewPeriod,omitempty" xml:"AutoRenewPeriod,omitempty"`
	// The client token used to ensure request idempotence.
	//
	// example:
	//
	// c2c5d1274axxxxxxxx
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The deployment architecture for the serverless instance. Valid values:
	//
	// - `shared`: The shared architecture, which is applicable to reserved and elastic (shared) instances and pay-as-you-go instances.
	//
	// - `dedicated`: The dedicated architecture, which is applicable to reserved and elastic (dedicated) instances.
	//
	// example:
	//
	// shared
	Edition *string `json:"Edition,omitempty" xml:"Edition,omitempty"`
	// This parameter is applicable only to dedicated instances. Specifies whether to enable data-at-rest encryption for the instance.
	//
	// example:
	//
	// false
	EncryptedInstance *bool `json:"EncryptedInstance,omitempty" xml:"EncryptedInstance,omitempty"`
	// The name of the instance. The name can be up to 64 characters long.
	//
	// example:
	//
	// amqp-xxxxx
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The instance type.
	//
	// This parameter is required for subscription instances. Valid values:
	//
	// - `professional`: Professional Edition
	//
	// - `enterprise`: Enterprise Edition
	//
	// - `vip`: Platinum Edition
	//
	// You do not need to specify this parameter for serverless instances.
	//
	// example:
	//
	// professional
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// This parameter applies only to dedicated instances and is required if `EncryptedInstance` is set to `true`. It specifies the ID of the KMS key used for data-at-rest encryption. The key must meet the following requirements:
	//
	// - The key cannot be a service key.
	//
	// - The key must be in the Enabled state.
	//
	// - The key must be a symmetric key, not an asymmetric key.
	//
	// - The key usage must be for encryption and decryption.
	//
	// - If the KMS key expires or is deleted, data reads and writes will become unavailable, and the ApsaraMQ for RabbitMQ instance may become inoperable.
	//
	// example:
	//
	// key-xxx
	KmsKeyId *string `json:"KmsKeyId,omitempty" xml:"KmsKeyId,omitempty"`
	// Specifies whether to enable only the TLS-encrypted port. This parameter applies only to reserved and elastic (dedicated) instances, and Platinum Edition instances.
	//
	// example:
	//
	// tcp_and_ssl
	ListenerMode *string `json:"ListenerMode,omitempty" xml:"ListenerMode,omitempty"`
	// The maximum number of connections.
	//
	// For information about the valid values, see the instance specifications on the [ApsaraMQ for RabbitMQ](https://common-buy.aliyun.com/?commodityCode=ons_onsproxy_pre) product page.
	//
	// example:
	//
	// 50000
	MaxConnections *int32 `json:"MaxConnections,omitempty" xml:"MaxConnections,omitempty"`
	// The peak transactions per second (TPS) of the instance over the public network.
	//
	// For information about the valid values, see the instance specifications on the [ApsaraMQ for RabbitMQ](https://common-buy.aliyun.com/?commodityCode=ons_onsproxy_pre) product page.
	//
	// example:
	//
	// 128
	MaxEipTps *int64 `json:"MaxEipTps,omitempty" xml:"MaxEipTps,omitempty"`
	// The peak transactions per second (TPS) of the instance over a private network.
	//
	// For information about the valid values, see the instance specifications on the [ApsaraMQ for RabbitMQ](https://common-buy.aliyun.com/?commodityCode=ons_onsproxy_pre) product page.
	//
	// example:
	//
	// 1000
	MaxPrivateTps *int64 `json:"MaxPrivateTps,omitempty" xml:"MaxPrivateTps,omitempty"`
	// The billing method of the instance. Valid values:
	//
	// - `Subscription`: The subscription-based billing method.
	//
	// - `PayAsYouGo`: The pay-as-you-go method for serverless instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// Subscription
	PaymentType *string `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
	// The subscription duration. The `PeriodCycle` parameter specifies the unit.
	//
	// > This parameter is required if you set `PaymentType` to `Subscription`. The default value is 1.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The unit of the subscription duration. Valid values:
	//
	// - `Month`: month
	//
	// - `Year`: year
	//
	// This parameter is required if you set `PaymentType` to `Subscription`. The default value is `Month`.
	//
	// example:
	//
	// Month
	PeriodCycle *string `json:"PeriodCycle,omitempty" xml:"PeriodCycle,omitempty"`
	// The provisioned TPS capacity for a reserved and elastic instance.
	//
	// example:
	//
	// 2000
	ProvisionedCapacity *int32 `json:"ProvisionedCapacity,omitempty" xml:"ProvisionedCapacity,omitempty"`
	// The queue capacity of the instance.
	//
	// For information about the valid values, see the instance specifications on the [ApsaraMQ for RabbitMQ](https://common-buy.aliyun.com/?commodityCode=ons_onsproxy_pre) product page.
	//
	// example:
	//
	// 1000
	QueueCapacity *int32 `json:"QueueCapacity,omitempty" xml:"QueueCapacity,omitempty"`
	// The renewal status. This parameter is equivalent to `AutoRenew`. Valid value:
	//
	// - `AutoRenewal`: Enables auto-renewal.
	//
	// > Both `AutoRenew` and `RenewStatus` specify the renewal method. If you specify both parameters, the value of `RenewStatus` takes precedence.
	//
	// example:
	//
	// false
	RenewStatus *string `json:"RenewStatus,omitempty" xml:"RenewStatus,omitempty"`
	// The unit of the auto-renewal duration. Valid values:
	//
	// - `Month`: month
	//
	// - `Year`: year
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
	// The ID of the security group. This security group is used for the PrivateLink-based endpoint. The security group must meet the following requirements:
	//
	// 1. Add an inbound rule to allow traffic on TCP ports 5672 and 5671.
	//
	// 2. Managed security groups are not supported.
	//
	// 3. The security group must belong to the specified VPC.
	//
	// This parameter is required.
	//
	// example:
	//
	// sg-xxx
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The billing type of the serverless instance. Valid value:
	//
	// - `onDemand`: pay-as-you-go
	//
	// example:
	//
	// onDemand
	ServerlessChargeType *string `json:"ServerlessChargeType,omitempty" xml:"ServerlessChargeType,omitempty"`
	ServerlessSwitch     *bool   `json:"ServerlessSwitch,omitempty" xml:"ServerlessSwitch,omitempty"`
	// The message storage space. Unit: GB. Valid values:
	//
	// - Professional Edition and Enterprise Edition instances: The value is fixed at 0.
	//
	// > A value of 0 means storage is not charged for Professional Edition and Enterprise Edition instances; it does not mean the instances lack storage space.
	//
	// - Platinum Edition instances: m × 100, where m is an integer from 7 to 28.
	//
	// example:
	//
	// 7
	StorageSize *int32 `json:"StorageSize,omitempty" xml:"StorageSize,omitempty"`
	// Specifies whether to enable access over the public network. Valid values:
	//
	// - `true`: Enables access over the public network.
	//
	// - `false`: Disables access over the public network.
	//
	// example:
	//
	// true
	SupportEip *bool `json:"SupportEip,omitempty" xml:"SupportEip,omitempty"`
	// Specifies whether to enable the message trace feature. Valid values:
	//
	// - `true`: Enables the message trace feature.
	//
	// - `false`: Disables the message trace feature.
	//
	// > 	- The message trace feature is included for 15 days at no charge on Platinum Edition instances. For these instances, you must enable this feature and set the retention period to 15 days.
	//
	// - For other instance types, you can enable or disable this feature.
	//
	// example:
	//
	// true
	SupportTracing *bool `json:"SupportTracing,omitempty" xml:"SupportTracing,omitempty"`
	// The resource tags.
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The retention period of message traces. Unit: days. Valid values:
	//
	// - `3`
	//
	// - `7`
	//
	// - `15`
	//
	// This parameter is required if you set `SupportTracing` to `true`.
	//
	// example:
	//
	// 3
	TracingStorageTime *int32 `json:"TracingStorageTime,omitempty" xml:"TracingStorageTime,omitempty"`
	// The ID of the VPC. This parameter is used to create a PrivateLink-based endpoint.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-xxx
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The vSwitch IDs used to create a PrivateLink-based endpoint when you create the instance. The vSwitches must meet the following requirements:
	//
	// 1. You must specify two vSwitches that reside in different availability zones, except for regions that have only a single availability zone.
	//
	// 2. The vSwitches must belong to the specified VPC.
	//
	// 3. The vSwitches must be in the Available state.
	//
	// 4. Each vSwitch must have at least 20 available IP addresses.
	//
	// 5. The availability zones for the specified vSwitches must support NLB instance creation.
	//
	// This parameter is required.
	VswitchIdsShrink *string `json:"VswitchIds,omitempty" xml:"VswitchIds,omitempty"`
}

func (s CreateInstanceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateInstanceShrinkRequest) GetAuthModel() *string {
	return s.AuthModel
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

func (s *CreateInstanceShrinkRequest) GetListenerMode() *string {
	return s.ListenerMode
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

func (s *CreateInstanceShrinkRequest) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *CreateInstanceShrinkRequest) GetServerlessChargeType() *string {
	return s.ServerlessChargeType
}

func (s *CreateInstanceShrinkRequest) GetServerlessSwitch() *bool {
	return s.ServerlessSwitch
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

func (s *CreateInstanceShrinkRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateInstanceShrinkRequest) GetVswitchIdsShrink() *string {
	return s.VswitchIdsShrink
}

func (s *CreateInstanceShrinkRequest) SetAuthModel(v string) *CreateInstanceShrinkRequest {
	s.AuthModel = &v
	return s
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

func (s *CreateInstanceShrinkRequest) SetListenerMode(v string) *CreateInstanceShrinkRequest {
	s.ListenerMode = &v
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

func (s *CreateInstanceShrinkRequest) SetSecurityGroupId(v string) *CreateInstanceShrinkRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateInstanceShrinkRequest) SetServerlessChargeType(v string) *CreateInstanceShrinkRequest {
	s.ServerlessChargeType = &v
	return s
}

func (s *CreateInstanceShrinkRequest) SetServerlessSwitch(v bool) *CreateInstanceShrinkRequest {
	s.ServerlessSwitch = &v
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

func (s *CreateInstanceShrinkRequest) SetVpcId(v string) *CreateInstanceShrinkRequest {
	s.VpcId = &v
	return s
}

func (s *CreateInstanceShrinkRequest) SetVswitchIdsShrink(v string) *CreateInstanceShrinkRequest {
	s.VswitchIdsShrink = &v
	return s
}

func (s *CreateInstanceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
