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
	Data *GetInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	return dara.Validate(s)
}

type GetInstanceResponseBodyData struct {
	// example:
	//
	// false
	AutoRenewInstance *bool `json:"AutoRenewInstance,omitempty" xml:"AutoRenewInstance,omitempty"`
	// example:
	//
	// amqp-cn-st21x7kv****.not-support
	ClassicEndpoint   *string `json:"ClassicEndpoint,omitempty" xml:"ClassicEndpoint,omitempty"`
	Edition           *string `json:"Edition,omitempty" xml:"Edition,omitempty"`
	EncryptedInstance *bool   `json:"EncryptedInstance,omitempty" xml:"EncryptedInstance,omitempty"`
	// example:
	//
	// 1651507200000
	ExpireTime *int64 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// example:
	//
	// amqp-cn-*********
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// yunQi-instance
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// example:
	//
	// enterprise
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	KmsKeyId     *string `json:"KmsKeyId,omitempty" xml:"KmsKeyId,omitempty"`
	// example:
	//
	// 1500
	MaxConnections *int32 `json:"MaxConnections,omitempty" xml:"MaxConnections,omitempty"`
	// example:
	//
	// 1000
	MaxEipTps *int32 `json:"MaxEipTps,omitempty" xml:"MaxEipTps,omitempty"`
	// example:
	//
	// 1000
	MaxQueue *int32 `json:"MaxQueue,omitempty" xml:"MaxQueue,omitempty"`
	// example:
	//
	// 1000
	MaxTps *int32 `json:"MaxTps,omitempty" xml:"MaxTps,omitempty"`
	// example:
	//
	// 50
	MaxVhost *int32 `json:"MaxVhost,omitempty" xml:"MaxVhost,omitempty"`
	// example:
	//
	// 1651507200000
	OrderCreateTime *int64 `json:"OrderCreateTime,omitempty" xml:"OrderCreateTime,omitempty"`
	// example:
	//
	// PRE_PAID
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// example:
	//
	// amqp-cn-st21x7kv****.mq-amqp.cn-hangzhou-a.aliyuncs.com
	PrivateEndpoint     *string `json:"PrivateEndpoint,omitempty" xml:"PrivateEndpoint,omitempty"`
	ProvisionedCapacity *int32  `json:"ProvisionedCapacity,omitempty" xml:"ProvisionedCapacity,omitempty"`
	// example:
	//
	// xxx.cn-hangzhou.xxx.net.mq.amqp.aliyuncs.com
	PublicEndpoint  *string `json:"PublicEndpoint,omitempty" xml:"PublicEndpoint,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// SERVING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 200
	StorageSize *int32 `json:"StorageSize,omitempty" xml:"StorageSize,omitempty"`
	// example:
	//
	// true
	SupportEIP *bool `json:"SupportEIP,omitempty" xml:"SupportEIP,omitempty"`
	// example:
	//
	// True
	SupportTracing *bool                              `json:"SupportTracing,omitempty" xml:"SupportTracing,omitempty"`
	Tags           []*GetInstanceResponseBodyDataTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// example:
	//
	// 15
	TracingStorageTime *int32 `json:"TracingStorageTime,omitempty" xml:"TracingStorageTime,omitempty"`
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

func (s *GetInstanceResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetInstanceResponseBodyDataTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
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
