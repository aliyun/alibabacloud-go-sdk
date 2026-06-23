// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetInstanceResponseBodyData) *GetInstanceResponseBody
	GetData() *GetInstanceResponseBodyData
	SetRequestId(v string) *GetInstanceResponseBody
	GetRequestId() *string
}

type GetInstanceResponseBody struct {
	// The returned data.
	Data *GetInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The unique ID generated for the request. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// 92385FD2-624A-48C9-8FB5-753F2AFA***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBody) GetData() *GetInstanceResponseBodyData {
	return s.Data
}

func (s *GetInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetInstanceResponseBody) SetData(v *GetInstanceResponseBodyData) *GetInstanceResponseBody {
	s.Data = v
	return s
}

func (s *GetInstanceResponseBody) SetRequestId(v string) *GetInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetInstanceResponseBodyData struct {
	// Indicates whether auto-renewal is enabled for the instance.
	//
	// example:
	//
	// false
	AutoRenewInstance *bool `json:"AutoRenewInstance,omitempty" xml:"AutoRenewInstance,omitempty"`
	// The classic network endpoint. This parameter is deprecated.
	//
	// example:
	//
	// amqp-cn-st21x7kv****.not-support
	ClassicEndpoint *string `json:"ClassicEndpoint,omitempty" xml:"ClassicEndpoint,omitempty"`
	// The deployment architecture. Valid values:
	//
	// - shared: shared architecture, which is suitable for reserved and elastic (shared) instances and pay-as-you-go instances.
	//
	// - dedicated: dedicated architecture, which is suitable for reserved and elastic (dedicated) instances.
	//
	// example:
	//
	// shared
	Edition *string `json:"Edition,omitempty" xml:"Edition,omitempty"`
	// Indicates whether storage encryption is enabled for the instance data.
	//
	// example:
	//
	// true
	EncryptedInstance *bool `json:"EncryptedInstance,omitempty" xml:"EncryptedInstance,omitempty"`
	// The timestamp that indicates when the instance expires, in milliseconds.
	//
	// > The value is a long integer. Handle it with care in certain programming languages to prevent precision loss.
	//
	// example:
	//
	// 1651507200000
	ExpireTime *int64 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// amqp-cn-*********
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the instance. A length of 64 characters or less is recommended.
	//
	// example:
	//
	// yunQi-instance
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The instance type.
	//
	// - PROFESSIONAL: Professional Edition
	//
	// - ENTERPRISE: Enterprise Edition
	//
	// - VIP: Platinum Edition
	//
	// - SERVERLESS: Serverless Edition
	//
	// example:
	//
	// enterprise
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The KMS key ID of the cloud disk.
	//
	// example:
	//
	// key-hzz6566e86byam3lg5rw4
	KmsKeyId *string `json:"KmsKeyId,omitempty" xml:"KmsKeyId,omitempty"`
	// The listener mode. A value of tcp_and_ssl enables both port 5672 and 5671, while ssl_only enables only port 5671.
	//
	// example:
	//
	// tcp_and_ssl
	ListenerMode *string `json:"ListenerMode,omitempty" xml:"ListenerMode,omitempty"`
	// The maximum number of connections.
	//
	// For valid values, see the [ApsaraMQ for RabbitMQ purchase page](https://common-buy.aliyun.com/?commodityCode=ons_onsproxy_pre).
	//
	// example:
	//
	// 1500
	MaxConnections *int32 `json:"MaxConnections,omitempty" xml:"MaxConnections,omitempty"`
	// The peak public TPS.
	//
	// For valid values, see the [ApsaraMQ for RabbitMQ purchase page](https://common-buy.aliyun.com/?commodityCode=ons_onsproxy_pre).
	//
	// example:
	//
	// 1000
	MaxEipTps *int32 `json:"MaxEipTps,omitempty" xml:"MaxEipTps,omitempty"`
	// The maximum number of queues for the instance.
	//
	// example:
	//
	// 1000
	MaxQueue *int32 `json:"MaxQueue,omitempty" xml:"MaxQueue,omitempty"`
	// The peak private TPS.
	//
	// example:
	//
	// 1000
	MaxTps *int32 `json:"MaxTps,omitempty" xml:"MaxTps,omitempty"`
	// The maximum number of vhosts for the instance.
	//
	// example:
	//
	// 50
	MaxVhost *int32 `json:"MaxVhost,omitempty" xml:"MaxVhost,omitempty"`
	// The timestamp that indicates when the order was created, in milliseconds.
	//
	// > The value is a long integer. Handle it with care in certain programming languages to prevent precision loss.
	//
	// example:
	//
	// 1651507200000
	OrderCreateTime *int64 `json:"OrderCreateTime,omitempty" xml:"OrderCreateTime,omitempty"`
	// The billing method.
	//
	// - PRE_PAID: subscription
	//
	// - POST_PAID: pay-as-you-go
	//
	// example:
	//
	// PRE_PAID
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// The VPC endpoint of the instance.
	//
	// example:
	//
	// amqp-cn-st21x7kv****.mq-amqp.cn-hangzhou-a.aliyuncs.com
	PrivateEndpoint *string `json:"PrivateEndpoint,omitempty" xml:"PrivateEndpoint,omitempty"`
	// The reserved TPS capacity for reserved and elastic instances.
	//
	// example:
	//
	// 2000
	ProvisionedCapacity *int32 `json:"ProvisionedCapacity,omitempty" xml:"ProvisionedCapacity,omitempty"`
	// The public endpoint of the instance.
	//
	// example:
	//
	// xxx.cn-hangzhou.xxx.net.mq.amqp.aliyuncs.com
	PublicEndpoint *string `json:"PublicEndpoint,omitempty" xml:"PublicEndpoint,omitempty"`
	// The ID of the resource group for the instance.
	//
	// example:
	//
	// rg-acfm2vn6jkayvfy
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The security group ID used to create a PrivateLink endpoint for the instance.
	//
	// example:
	//
	// sg-xxx
	SecurityGroupId  *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	ServerlessSwitch *bool   `json:"ServerlessSwitch,omitempty" xml:"ServerlessSwitch,omitempty"`
	// The instance status. Valid values:
	//
	// - DEPLOYING: The instance is being deployed.
	//
	// - EXPIRED: The instance has expired.
	//
	// - SERVING: The instance is in service.
	//
	// - RELEASED: The instance has been released.
	//
	// example:
	//
	// SERVING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The disk capacity. Unit: GB.
	//
	// > For Professional and Enterprise Edition instances, this parameter returns **-1**.
	//
	// example:
	//
	// 200
	StorageSize *int32 `json:"StorageSize,omitempty" xml:"StorageSize,omitempty"`
	// Indicates whether the instance supports EIPs.
	//
	// example:
	//
	// true
	SupportEIP *bool `json:"SupportEIP,omitempty" xml:"SupportEIP,omitempty"`
	// Indicates whether the message trace feature is enabled.
	//
	// example:
	//
	// True
	SupportTracing *bool `json:"SupportTracing,omitempty" xml:"SupportTracing,omitempty"`
	// The list of tags.
	Tags []*GetInstanceResponseBodyDataTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The retention period of message traces. Unit: days. Valid values:
	//
	// - 3: 3 days
	//
	// - 7: 7 days
	//
	// - 15: 15 days
	//
	// This parameter applies only when `SupportTracing` is set to true.
	//
	// example:
	//
	// 15
	TracingStorageTime *int32 `json:"TracingStorageTime,omitempty" xml:"TracingStorageTime,omitempty"`
	// The VPC ID used to create a PrivateLink endpoint for the instance.
	//
	// example:
	//
	// vpc-xxx
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The VSwitch IDs used to create a PrivateLink endpoint for the instance.
	VswitchIds []*string `json:"VswitchIds,omitempty" xml:"VswitchIds,omitempty" type:"Repeated"`
}

func (s GetInstanceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyData) GetAutoRenewInstance() *bool {
	return s.AutoRenewInstance
}

func (s *GetInstanceResponseBodyData) GetClassicEndpoint() *string {
	return s.ClassicEndpoint
}

func (s *GetInstanceResponseBodyData) GetEdition() *string {
	return s.Edition
}

func (s *GetInstanceResponseBodyData) GetEncryptedInstance() *bool {
	return s.EncryptedInstance
}

func (s *GetInstanceResponseBodyData) GetExpireTime() *int64 {
	return s.ExpireTime
}

func (s *GetInstanceResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetInstanceResponseBodyData) GetInstanceName() *string {
	return s.InstanceName
}

func (s *GetInstanceResponseBodyData) GetInstanceType() *string {
	return s.InstanceType
}

func (s *GetInstanceResponseBodyData) GetKmsKeyId() *string {
	return s.KmsKeyId
}

func (s *GetInstanceResponseBodyData) GetListenerMode() *string {
	return s.ListenerMode
}

func (s *GetInstanceResponseBodyData) GetMaxConnections() *int32 {
	return s.MaxConnections
}

func (s *GetInstanceResponseBodyData) GetMaxEipTps() *int32 {
	return s.MaxEipTps
}

func (s *GetInstanceResponseBodyData) GetMaxQueue() *int32 {
	return s.MaxQueue
}

func (s *GetInstanceResponseBodyData) GetMaxTps() *int32 {
	return s.MaxTps
}

func (s *GetInstanceResponseBodyData) GetMaxVhost() *int32 {
	return s.MaxVhost
}

func (s *GetInstanceResponseBodyData) GetOrderCreateTime() *int64 {
	return s.OrderCreateTime
}

func (s *GetInstanceResponseBodyData) GetOrderType() *string {
	return s.OrderType
}

func (s *GetInstanceResponseBodyData) GetPrivateEndpoint() *string {
	return s.PrivateEndpoint
}

func (s *GetInstanceResponseBodyData) GetProvisionedCapacity() *int32 {
	return s.ProvisionedCapacity
}

func (s *GetInstanceResponseBodyData) GetPublicEndpoint() *string {
	return s.PublicEndpoint
}

func (s *GetInstanceResponseBodyData) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetInstanceResponseBodyData) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *GetInstanceResponseBodyData) GetServerlessSwitch() *bool {
	return s.ServerlessSwitch
}

func (s *GetInstanceResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetInstanceResponseBodyData) GetStorageSize() *int32 {
	return s.StorageSize
}

func (s *GetInstanceResponseBodyData) GetSupportEIP() *bool {
	return s.SupportEIP
}

func (s *GetInstanceResponseBodyData) GetSupportTracing() *bool {
	return s.SupportTracing
}

func (s *GetInstanceResponseBodyData) GetTags() []*GetInstanceResponseBodyDataTags {
	return s.Tags
}

func (s *GetInstanceResponseBodyData) GetTracingStorageTime() *int32 {
	return s.TracingStorageTime
}

func (s *GetInstanceResponseBodyData) GetVpcId() *string {
	return s.VpcId
}

func (s *GetInstanceResponseBodyData) GetVswitchIds() []*string {
	return s.VswitchIds
}

func (s *GetInstanceResponseBodyData) SetAutoRenewInstance(v bool) *GetInstanceResponseBodyData {
	s.AutoRenewInstance = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetClassicEndpoint(v string) *GetInstanceResponseBodyData {
	s.ClassicEndpoint = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetEdition(v string) *GetInstanceResponseBodyData {
	s.Edition = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetEncryptedInstance(v bool) *GetInstanceResponseBodyData {
	s.EncryptedInstance = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetExpireTime(v int64) *GetInstanceResponseBodyData {
	s.ExpireTime = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetInstanceId(v string) *GetInstanceResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetInstanceName(v string) *GetInstanceResponseBodyData {
	s.InstanceName = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetInstanceType(v string) *GetInstanceResponseBodyData {
	s.InstanceType = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetKmsKeyId(v string) *GetInstanceResponseBodyData {
	s.KmsKeyId = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetListenerMode(v string) *GetInstanceResponseBodyData {
	s.ListenerMode = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetMaxConnections(v int32) *GetInstanceResponseBodyData {
	s.MaxConnections = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetMaxEipTps(v int32) *GetInstanceResponseBodyData {
	s.MaxEipTps = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetMaxQueue(v int32) *GetInstanceResponseBodyData {
	s.MaxQueue = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetMaxTps(v int32) *GetInstanceResponseBodyData {
	s.MaxTps = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetMaxVhost(v int32) *GetInstanceResponseBodyData {
	s.MaxVhost = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetOrderCreateTime(v int64) *GetInstanceResponseBodyData {
	s.OrderCreateTime = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetOrderType(v string) *GetInstanceResponseBodyData {
	s.OrderType = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetPrivateEndpoint(v string) *GetInstanceResponseBodyData {
	s.PrivateEndpoint = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetProvisionedCapacity(v int32) *GetInstanceResponseBodyData {
	s.ProvisionedCapacity = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetPublicEndpoint(v string) *GetInstanceResponseBodyData {
	s.PublicEndpoint = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetResourceGroupId(v string) *GetInstanceResponseBodyData {
	s.ResourceGroupId = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetSecurityGroupId(v string) *GetInstanceResponseBodyData {
	s.SecurityGroupId = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetServerlessSwitch(v bool) *GetInstanceResponseBodyData {
	s.ServerlessSwitch = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetStatus(v string) *GetInstanceResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetStorageSize(v int32) *GetInstanceResponseBodyData {
	s.StorageSize = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetSupportEIP(v bool) *GetInstanceResponseBodyData {
	s.SupportEIP = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetSupportTracing(v bool) *GetInstanceResponseBodyData {
	s.SupportTracing = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetTags(v []*GetInstanceResponseBodyDataTags) *GetInstanceResponseBodyData {
	s.Tags = v
	return s
}

func (s *GetInstanceResponseBodyData) SetTracingStorageTime(v int32) *GetInstanceResponseBodyData {
	s.TracingStorageTime = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetVpcId(v string) *GetInstanceResponseBodyData {
	s.VpcId = &v
	return s
}

func (s *GetInstanceResponseBodyData) SetVswitchIds(v []*string) *GetInstanceResponseBodyData {
	s.VswitchIds = v
	return s
}

func (s *GetInstanceResponseBodyData) Validate() error {
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

type GetInstanceResponseBodyDataTags struct {
	// The tag key.
	//
	// example:
	//
	// Tag key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// Tag value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetInstanceResponseBodyDataTags) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceResponseBodyDataTags) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyDataTags) GetKey() *string {
	return s.Key
}

func (s *GetInstanceResponseBodyDataTags) GetValue() *string {
	return s.Value
}

func (s *GetInstanceResponseBodyDataTags) SetKey(v string) *GetInstanceResponseBodyDataTags {
	s.Key = &v
	return s
}

func (s *GetInstanceResponseBodyDataTags) SetValue(v string) *GetInstanceResponseBodyDataTags {
	s.Value = &v
	return s
}

func (s *GetInstanceResponseBodyDataTags) Validate() error {
	return dara.Validate(s)
}
