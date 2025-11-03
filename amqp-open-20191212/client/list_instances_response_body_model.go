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
	// The data returned.
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
	// The instances.
	Instances []*ListInstancesResponseBodyDataInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	// The maximum number of entries returned.
	//
	// example:
	//
	// 1
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that marks the end of the current returned page. If this parameter is empty, all data is retrieved.
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
	// Indicates whether the instance is automatically renewed.
	//
	// example:
	//
	// false
	AutoRenewInstance *bool `json:"AutoRenewInstance,omitempty" xml:"AutoRenewInstance,omitempty"`
	// The endpoint that is used to access the instance over the classic network. This parameter is no longer available.
	//
	// example:
	//
	// amqp-cn-st21x7kv****.not-support
	ClassicEndpoint *string `json:"ClassicEndpoint,omitempty" xml:"ClassicEndpoint,omitempty"`
	Edition         *string `json:"Edition,omitempty" xml:"Edition,omitempty"`
	// Indicates whether the encryption at rest feature is enabled for the instance.
	//
	// example:
	//
	// false
	EncryptedInstance *bool `json:"EncryptedInstance,omitempty" xml:"EncryptedInstance,omitempty"`
	// The timestamp that indicates when the instance expires. Unit: milliseconds.
	//
	// example:
	//
	// 1651507200000
	ExpireTime *int64 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The instance ID
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
	// The instance type.
	//
	// 	- PROFESSIONAL: Professional Edition
	//
	// 	- ENTERPRISE: Enterprise Edition
	//
	// 	- VIP: Enterprise Platinum Edition
	//
	// example:
	//
	// professional
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The ID of the Key Management Service (KMS) key used for the data disk.
	//
	// example:
	//
	// key-bjj66c2a893vmhawtq5fd
	KmsKeyId *string `json:"KmsKeyId,omitempty" xml:"KmsKeyId,omitempty"`
	// The maximum number of Internet-based transactions per second (TPS) for the instance.
	//
	// example:
	//
	// 24832
	MaxEipTps *int32 `json:"MaxEipTps,omitempty" xml:"MaxEipTps,omitempty"`
	// The maximum number of queues on the instance.
	//
	// example:
	//
	// 50
	MaxQueue *int32 `json:"MaxQueue,omitempty" xml:"MaxQueue,omitempty"`
	// The maximum number of VPC-based TPS for the instance.
	//
	// example:
	//
	// 5000
	MaxTps *int32 `json:"MaxTps,omitempty" xml:"MaxTps,omitempty"`
	// The maximum number of vhosts on the instance.
	//
	// example:
	//
	// 50
	MaxVhost *int32 `json:"MaxVhost,omitempty" xml:"MaxVhost,omitempty"`
	// The timestamp that indicates when the order was created. Unit: milliseconds.
	//
	// example:
	//
	// 1572441939000
	OrderCreateTime *int64 `json:"OrderCreateTime,omitempty" xml:"OrderCreateTime,omitempty"`
	// The billing method. Valid values:
	//
	// 	- PrePaid: the subscription billing method.
	//
	// 	- POST_PAID: the pay-as-you-go billing method.
	//
	// example:
	//
	// PRE_PAID
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// The virtual private cloud (VPC) endpoint of the instance.
	//
	// example:
	//
	// amqp-cn-st21x7kv****.mq-amqp.cn-hangzhou-a.aliyuncs.com
	PrivateEndpoint     *string `json:"PrivateEndpoint,omitempty" xml:"PrivateEndpoint,omitempty"`
	ProvisionedCapacity *int32  `json:"ProvisionedCapacity,omitempty" xml:"ProvisionedCapacity,omitempty"`
	// The public endpoint of the instance.
	//
	// example:
	//
	// amqp-cn-st21x7kv****.mq-amqp.cn-hangzhou-a.aliyuncs.com
	PublicEndpoint *string `json:"PublicEndpoint,omitempty" xml:"PublicEndpoint,omitempty"`
	// The ID of the resource group to which the instance belongs.
	//
	// example:
	//
	// rg-aek3axfj2w4czrq
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The instance status. Valid values:
	//
	// 	- DEPLOYING: The instance is being deployed.
	//
	// 	- EXPIRED: The instance is expired.
	//
	// 	- SERVING: The instance is running.
	//
	// 	- RELEASED: The instance is released.
	//
	// example:
	//
	// SERVING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The disk size. Unit: GB.
	//
	// >  For Professional Edition instances and Enterprise Edition instances, this parameter is unavailable and \\*\\*-1\\*\\	- is returned.
	//
	// example:
	//
	// 200
	StorageSize *int32 `json:"StorageSize,omitempty" xml:"StorageSize,omitempty"`
	// Indicates whether the instance supports elastic IP addresses (EIPs).
	//
	// example:
	//
	// true
	SupportEIP *bool `json:"SupportEIP,omitempty" xml:"SupportEIP,omitempty"`
	// The tags that are added to the instance.
	Tags []*ListInstancesResponseBodyDataInstancesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
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
