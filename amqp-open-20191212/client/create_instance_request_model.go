// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthModel(v string) *CreateInstanceRequest
	GetAuthModel() *string
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
	SetListenerMode(v string) *CreateInstanceRequest
	GetListenerMode() *string
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
	SetSecurityGroupId(v string) *CreateInstanceRequest
	GetSecurityGroupId() *string
	SetServerlessChargeType(v string) *CreateInstanceRequest
	GetServerlessChargeType() *string
	SetServerlessSwitch(v bool) *CreateInstanceRequest
	GetServerlessSwitch() *bool
	SetStorageSize(v int32) *CreateInstanceRequest
	GetStorageSize() *int32
	SetSupportEip(v bool) *CreateInstanceRequest
	GetSupportEip() *bool
	SetSupportTracing(v bool) *CreateInstanceRequest
	GetSupportTracing() *bool
	SetTags(v []*CreateInstanceRequestTags) *CreateInstanceRequest
	GetTags() []*CreateInstanceRequestTags
	SetTracingStorageTime(v int32) *CreateInstanceRequest
	GetTracingStorageTime() *int32
	SetVpcId(v string) *CreateInstanceRequest
	GetVpcId() *string
	SetVswitchIds(v []*string) *CreateInstanceRequest
	GetVswitchIds() []*string
}

type CreateInstanceRequest struct {
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
	Tags []*CreateInstanceRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
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
	VswitchIds []*string `json:"VswitchIds,omitempty" xml:"VswitchIds,omitempty" type:"Repeated"`
}

func (s CreateInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequest) GetAuthModel() *string {
	return s.AuthModel
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

func (s *CreateInstanceRequest) GetListenerMode() *string {
	return s.ListenerMode
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

func (s *CreateInstanceRequest) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *CreateInstanceRequest) GetServerlessChargeType() *string {
	return s.ServerlessChargeType
}

func (s *CreateInstanceRequest) GetServerlessSwitch() *bool {
	return s.ServerlessSwitch
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

func (s *CreateInstanceRequest) GetTags() []*CreateInstanceRequestTags {
	return s.Tags
}

func (s *CreateInstanceRequest) GetTracingStorageTime() *int32 {
	return s.TracingStorageTime
}

func (s *CreateInstanceRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateInstanceRequest) GetVswitchIds() []*string {
	return s.VswitchIds
}

func (s *CreateInstanceRequest) SetAuthModel(v string) *CreateInstanceRequest {
	s.AuthModel = &v
	return s
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

func (s *CreateInstanceRequest) SetListenerMode(v string) *CreateInstanceRequest {
	s.ListenerMode = &v
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

func (s *CreateInstanceRequest) SetSecurityGroupId(v string) *CreateInstanceRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateInstanceRequest) SetServerlessChargeType(v string) *CreateInstanceRequest {
	s.ServerlessChargeType = &v
	return s
}

func (s *CreateInstanceRequest) SetServerlessSwitch(v bool) *CreateInstanceRequest {
	s.ServerlessSwitch = &v
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

func (s *CreateInstanceRequest) SetTags(v []*CreateInstanceRequestTags) *CreateInstanceRequest {
	s.Tags = v
	return s
}

func (s *CreateInstanceRequest) SetTracingStorageTime(v int32) *CreateInstanceRequest {
	s.TracingStorageTime = &v
	return s
}

func (s *CreateInstanceRequest) SetVpcId(v string) *CreateInstanceRequest {
	s.VpcId = &v
	return s
}

func (s *CreateInstanceRequest) SetVswitchIds(v []*string) *CreateInstanceRequest {
	s.VswitchIds = v
	return s
}

func (s *CreateInstanceRequest) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateInstanceRequestTags struct {
	// The tag key.
	//
	// example:
	//
	// testKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// testValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateInstanceRequestTags) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceRequestTags) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequestTags) GetKey() *string {
	return s.Key
}

func (s *CreateInstanceRequestTags) GetValue() *string {
	return s.Value
}

func (s *CreateInstanceRequestTags) SetKey(v string) *CreateInstanceRequestTags {
	s.Key = &v
	return s
}

func (s *CreateInstanceRequestTags) SetValue(v string) *CreateInstanceRequestTags {
	s.Value = &v
	return s
}

func (s *CreateInstanceRequestTags) Validate() error {
	return dara.Validate(s)
}
