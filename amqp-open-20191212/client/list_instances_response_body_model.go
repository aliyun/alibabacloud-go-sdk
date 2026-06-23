// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListInstancesResponseBodyData) *ListInstancesResponseBody
	GetData() *ListInstancesResponseBodyData
	SetRequestId(v string) *ListInstancesResponseBody
	GetRequestId() *string
}

type ListInstancesResponseBody struct {
	// The returned data.
	Data *ListInstancesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// CCBB1225-C392-480E-8C7F-D09AB2CD2***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBody) GetData() *ListInstancesResponseBodyData {
	return s.Data
}

func (s *ListInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListInstancesResponseBody) SetData(v *ListInstancesResponseBodyData) *ListInstancesResponseBody {
	s.Data = v
	return s
}

func (s *ListInstancesResponseBody) SetRequestId(v string) *ListInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstancesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListInstancesResponseBodyData struct {
	// A list of instances.
	Instances []*ListInstancesResponseBodyDataInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	// The maximum number of entries returned per page.
	//
	// example:
	//
	// 1
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token for the next page of results. If this field is empty, it means all results have been returned.
	//
	// example:
	//
	// caebacccb2be03f84eb48b699f0a****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListInstancesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyData) GetInstances() []*ListInstancesResponseBodyDataInstances {
	return s.Instances
}

func (s *ListInstancesResponseBodyData) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListInstancesResponseBodyData) GetNextToken() *string {
	return s.NextToken
}

func (s *ListInstancesResponseBodyData) SetInstances(v []*ListInstancesResponseBodyDataInstances) *ListInstancesResponseBodyData {
	s.Instances = v
	return s
}

func (s *ListInstancesResponseBodyData) SetMaxResults(v int32) *ListInstancesResponseBodyData {
	s.MaxResults = &v
	return s
}

func (s *ListInstancesResponseBodyData) SetNextToken(v string) *ListInstancesResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *ListInstancesResponseBodyData) Validate() error {
	if s.Instances != nil {
		for _, item := range s.Instances {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListInstancesResponseBodyDataInstances struct {
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
	// The deployment architecture, which is applicable only to Serverless Edition instances. Valid values:
	//
	// - shared: A shared architecture, used for reserved, elastic (shared), and pay-as-you-go instances.
	//
	// - dedicated: A dedicated architecture, used for reserved and elastic (dedicated) instances.
	//
	// example:
	//
	// shared
	Edition *string `json:"Edition,omitempty" xml:"Edition,omitempty"`
	// Indicates whether storage encryption is enabled for the instance.
	//
	// example:
	//
	// false
	EncryptedInstance *bool `json:"EncryptedInstance,omitempty" xml:"EncryptedInstance,omitempty"`
	// The expiration timestamp of the instance, in milliseconds.
	//
	// example:
	//
	// 1651507200000
	ExpireTime *int64 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// amqp-cn-st21x7kv****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The instance name.
	//
	// example:
	//
	// amqp-cn-st21x7kv****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The instance type. Valid values:
	//
	// - professional: Professional Edition
	//
	// - enterprise: Enterprise Edition
	//
	// - vip: Platinum Edition
	//
	// <props="china">
	//
	// - serverless: Serverless Edition
	//
	// example:
	//
	// professional
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The ID of the KMS key used for data disk encryption.
	//
	// example:
	//
	// key-bjj66c2a893vmhawtq5fd
	KmsKeyId *string `json:"KmsKeyId,omitempty" xml:"KmsKeyId,omitempty"`
	// The port listener mode of the instance. `tcp_and_ssl` enables both port `5672` and port `5671`, while `ssl_only` enables only port `5671`.
	//
	// example:
	//
	// tcp_and_ssl
	ListenerMode *string `json:"ListenerMode,omitempty" xml:"ListenerMode,omitempty"`
	// The peak transactions per second (TPS) of the instance over the public network.
	//
	// example:
	//
	// 24832
	MaxEipTps *int32 `json:"MaxEipTps,omitempty" xml:"MaxEipTps,omitempty"`
	// The maximum number of queues for the instance.
	//
	// example:
	//
	// 50
	MaxQueue *int32 `json:"MaxQueue,omitempty" xml:"MaxQueue,omitempty"`
	// The peak transactions per second (TPS) of the instance over the private network.
	//
	// example:
	//
	// 5000
	MaxTps *int32 `json:"MaxTps,omitempty" xml:"MaxTps,omitempty"`
	// The maximum number of vhosts for the instance.
	//
	// example:
	//
	// 50
	MaxVhost *int32 `json:"MaxVhost,omitempty" xml:"MaxVhost,omitempty"`
	// The creation timestamp of the order, in milliseconds.
	//
	// example:
	//
	// 1572441939000
	OrderCreateTime *int64 `json:"OrderCreateTime,omitempty" xml:"OrderCreateTime,omitempty"`
	// The billing method. Valid values:
	//
	// - PRE_PAID: The instance uses the subscription billing method.
	//
	// - POST_PAID: The instance uses the pay-as-you-go billing method.
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
	// amqp-cn-st21x7kv****.mq-amqp.cn-hangzhou-a.aliyuncs.com
	PublicEndpoint *string `json:"PublicEndpoint,omitempty" xml:"PublicEndpoint,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-aek3axfj2w4czrq
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the security group to which the instance belongs. This security group is used for PrivateLink endpoint creation.
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
	// - SERVING: The instance is running.
	//
	// - RELEASED: The instance is released.
	//
	// example:
	//
	// SERVING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The storage capacity of the disk. Unit: GB.
	//
	// > This parameter returns a value of **-1*	- for Professional Edition and Enterprise Edition instances, to which it does not apply.
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
	// The tags attached to the instance.
	Tags []*ListInstancesResponseBodyDataInstancesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The ID of the VPC in which the instance resides. This VPC is used for PrivateLink endpoint creation.
	//
	// example:
	//
	// vpc-xxx
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The IDs of the VSwitches to which the instance is connected. These VSwitches are used for PrivateLink endpoint creation.
	VswitchIds []*string `json:"VswitchIds,omitempty" xml:"VswitchIds,omitempty" type:"Repeated"`
}

func (s ListInstancesResponseBodyDataInstances) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesResponseBodyDataInstances) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyDataInstances) GetAutoRenewInstance() *bool {
	return s.AutoRenewInstance
}

func (s *ListInstancesResponseBodyDataInstances) GetClassicEndpoint() *string {
	return s.ClassicEndpoint
}

func (s *ListInstancesResponseBodyDataInstances) GetEdition() *string {
	return s.Edition
}

func (s *ListInstancesResponseBodyDataInstances) GetEncryptedInstance() *bool {
	return s.EncryptedInstance
}

func (s *ListInstancesResponseBodyDataInstances) GetExpireTime() *int64 {
	return s.ExpireTime
}

func (s *ListInstancesResponseBodyDataInstances) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListInstancesResponseBodyDataInstances) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ListInstancesResponseBodyDataInstances) GetInstanceType() *string {
	return s.InstanceType
}

func (s *ListInstancesResponseBodyDataInstances) GetKmsKeyId() *string {
	return s.KmsKeyId
}

func (s *ListInstancesResponseBodyDataInstances) GetListenerMode() *string {
	return s.ListenerMode
}

func (s *ListInstancesResponseBodyDataInstances) GetMaxEipTps() *int32 {
	return s.MaxEipTps
}

func (s *ListInstancesResponseBodyDataInstances) GetMaxQueue() *int32 {
	return s.MaxQueue
}

func (s *ListInstancesResponseBodyDataInstances) GetMaxTps() *int32 {
	return s.MaxTps
}

func (s *ListInstancesResponseBodyDataInstances) GetMaxVhost() *int32 {
	return s.MaxVhost
}

func (s *ListInstancesResponseBodyDataInstances) GetOrderCreateTime() *int64 {
	return s.OrderCreateTime
}

func (s *ListInstancesResponseBodyDataInstances) GetOrderType() *string {
	return s.OrderType
}

func (s *ListInstancesResponseBodyDataInstances) GetPrivateEndpoint() *string {
	return s.PrivateEndpoint
}

func (s *ListInstancesResponseBodyDataInstances) GetProvisionedCapacity() *int32 {
	return s.ProvisionedCapacity
}

func (s *ListInstancesResponseBodyDataInstances) GetPublicEndpoint() *string {
	return s.PublicEndpoint
}

func (s *ListInstancesResponseBodyDataInstances) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListInstancesResponseBodyDataInstances) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *ListInstancesResponseBodyDataInstances) GetServerlessSwitch() *bool {
	return s.ServerlessSwitch
}

func (s *ListInstancesResponseBodyDataInstances) GetStatus() *string {
	return s.Status
}

func (s *ListInstancesResponseBodyDataInstances) GetStorageSize() *int32 {
	return s.StorageSize
}

func (s *ListInstancesResponseBodyDataInstances) GetSupportEIP() *bool {
	return s.SupportEIP
}

func (s *ListInstancesResponseBodyDataInstances) GetTags() []*ListInstancesResponseBodyDataInstancesTags {
	return s.Tags
}

func (s *ListInstancesResponseBodyDataInstances) GetVpcId() *string {
	return s.VpcId
}

func (s *ListInstancesResponseBodyDataInstances) GetVswitchIds() []*string {
	return s.VswitchIds
}

func (s *ListInstancesResponseBodyDataInstances) SetAutoRenewInstance(v bool) *ListInstancesResponseBodyDataInstances {
	s.AutoRenewInstance = &v
	return s
}

func (s *ListInstancesResponseBodyDataInstances) SetClassicEndpoint(v string) *ListInstancesResponseBodyDataInstances {
	s.ClassicEndpoint = &v
	return s
}

func (s *ListInstancesResponseBodyDataInstances) SetEdition(v string) *ListInstancesResponseBodyDataInstances {
	s.Edition = &v
	return s
}

func (s *ListInstancesResponseBodyDataInstances) SetEncryptedInstance(v bool) *ListInstancesResponseBodyDataInstances {
	s.EncryptedInstance = &v
	return s
}

func (s *ListInstancesResponseBodyDataInstances) SetExpireTime(v int64) *ListInstancesResponseBodyDataInstances {
	s.ExpireTime = &v
	return s
}

func (s *ListInstancesResponseBodyDataInstances) SetInstanceId(v string) *ListInstancesResponseBodyDataInstances {
	s.InstanceId = &v
	return s
}

func (s *ListInstancesResponseBodyDataInstances) SetInstanceName(v string) *ListInstancesResponseBodyDataInstances {
	s.InstanceName = &v
	return s
}

func (s *ListInstancesResponseBodyDataInstances) SetInstanceType(v string) *ListInstancesResponseBodyDataInstances {
	s.InstanceType = &v
	return s
}

func (s *ListInstancesResponseBodyDataInstances) SetKmsKeyId(v string) *ListInstancesResponseBodyDataInstances {
	s.KmsKeyId = &v
	return s
}

func (s *ListInstancesResponseBodyDataInstances) SetListenerMode(v string) *ListInstancesResponseBodyDataInstances {
	s.ListenerMode = &v
	return s
}

func (s *ListInstancesResponseBodyDataInstances) SetMaxEipTps(v int32) *ListInstancesResponseBodyDataInstances {
	s.MaxEipTps = &v
	return s
}

func (s *ListInstancesResponseBodyDataInstances) SetMaxQueue(v int32) *ListInstancesResponseBodyDataInstances {
	s.MaxQueue = &v
	return s
}

func (s *ListInstancesResponseBodyDataInstances) SetMaxTps(v int32) *ListInstancesResponseBodyDataInstances {
	s.MaxTps = &v
	return s
}

func (s *ListInstancesResponseBodyDataInstances) SetMaxVhost(v int32) *ListInstancesResponseBodyDataInstances {
	s.MaxVhost = &v
	return s
}

func (s *ListInstancesResponseBodyDataInstances) SetOrderCreateTime(v int64) *ListInstancesResponseBodyDataInstances {
	s.OrderCreateTime = &v
	return s
}

func (s *ListInstancesResponseBodyDataInstances) SetOrderType(v string) *ListInstancesResponseBodyDataInstances {
	s.OrderType = &v
	return s
}

func (s *ListInstancesResponseBodyDataInstances) SetPrivateEndpoint(v string) *ListInstancesResponseBodyDataInstances {
	s.PrivateEndpoint = &v
	return s
}

func (s *ListInstancesResponseBodyDataInstances) SetProvisionedCapacity(v int32) *ListInstancesResponseBodyDataInstances {
	s.ProvisionedCapacity = &v
	return s
}

func (s *ListInstancesResponseBodyDataInstances) SetPublicEndpoint(v string) *ListInstancesResponseBodyDataInstances {
	s.PublicEndpoint = &v
	return s
}

func (s *ListInstancesResponseBodyDataInstances) SetResourceGroupId(v string) *ListInstancesResponseBodyDataInstances {
	s.ResourceGroupId = &v
	return s
}

func (s *ListInstancesResponseBodyDataInstances) SetSecurityGroupId(v string) *ListInstancesResponseBodyDataInstances {
	s.SecurityGroupId = &v
	return s
}

func (s *ListInstancesResponseBodyDataInstances) SetServerlessSwitch(v bool) *ListInstancesResponseBodyDataInstances {
	s.ServerlessSwitch = &v
	return s
}

func (s *ListInstancesResponseBodyDataInstances) SetStatus(v string) *ListInstancesResponseBodyDataInstances {
	s.Status = &v
	return s
}

func (s *ListInstancesResponseBodyDataInstances) SetStorageSize(v int32) *ListInstancesResponseBodyDataInstances {
	s.StorageSize = &v
	return s
}

func (s *ListInstancesResponseBodyDataInstances) SetSupportEIP(v bool) *ListInstancesResponseBodyDataInstances {
	s.SupportEIP = &v
	return s
}

func (s *ListInstancesResponseBodyDataInstances) SetTags(v []*ListInstancesResponseBodyDataInstancesTags) *ListInstancesResponseBodyDataInstances {
	s.Tags = v
	return s
}

func (s *ListInstancesResponseBodyDataInstances) SetVpcId(v string) *ListInstancesResponseBodyDataInstances {
	s.VpcId = &v
	return s
}

func (s *ListInstancesResponseBodyDataInstances) SetVswitchIds(v []*string) *ListInstancesResponseBodyDataInstances {
	s.VswitchIds = v
	return s
}

func (s *ListInstancesResponseBodyDataInstances) Validate() error {
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

type ListInstancesResponseBodyDataInstancesTags struct {
	// The tag key.
	//
	// example:
	//
	// region
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// hangzhou
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListInstancesResponseBodyDataInstancesTags) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesResponseBodyDataInstancesTags) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyDataInstancesTags) GetKey() *string {
	return s.Key
}

func (s *ListInstancesResponseBodyDataInstancesTags) GetValue() *string {
	return s.Value
}

func (s *ListInstancesResponseBodyDataInstancesTags) SetKey(v string) *ListInstancesResponseBodyDataInstancesTags {
	s.Key = &v
	return s
}

func (s *ListInstancesResponseBodyDataInstancesTags) SetValue(v string) *ListInstancesResponseBodyDataInstancesTags {
	s.Value = &v
	return s
}

func (s *ListInstancesResponseBodyDataInstancesTags) Validate() error {
	return dara.Validate(s)
}
