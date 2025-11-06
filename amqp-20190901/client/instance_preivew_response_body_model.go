// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstancePreivewResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *InstancePreivewResponseBody
	GetCode() *int32
	SetData(v *InstancePreivewResponseBodyData) *InstancePreivewResponseBody
	GetData() *InstancePreivewResponseBodyData
	SetMessage(v string) *InstancePreivewResponseBody
	GetMessage() *string
	SetRequestId(v string) *InstancePreivewResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *InstancePreivewResponseBody
	GetSuccess() *bool
}

type InstancePreivewResponseBody struct {
	Code      *int32                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *InstancePreivewResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                            `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s InstancePreivewResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InstancePreivewResponseBody) GoString() string {
	return s.String()
}

func (s *InstancePreivewResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *InstancePreivewResponseBody) GetData() *InstancePreivewResponseBodyData {
	return s.Data
}

func (s *InstancePreivewResponseBody) GetMessage() *string {
	return s.Message
}

func (s *InstancePreivewResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InstancePreivewResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *InstancePreivewResponseBody) SetCode(v int32) *InstancePreivewResponseBody {
	s.Code = &v
	return s
}

func (s *InstancePreivewResponseBody) SetData(v *InstancePreivewResponseBodyData) *InstancePreivewResponseBody {
	s.Data = v
	return s
}

func (s *InstancePreivewResponseBody) SetMessage(v string) *InstancePreivewResponseBody {
	s.Message = &v
	return s
}

func (s *InstancePreivewResponseBody) SetRequestId(v string) *InstancePreivewResponseBody {
	s.RequestId = &v
	return s
}

func (s *InstancePreivewResponseBody) SetSuccess(v bool) *InstancePreivewResponseBody {
	s.Success = &v
	return s
}

func (s *InstancePreivewResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type InstancePreivewResponseBodyData struct {
	ExchangeNum *int32                                    `json:"ExchangeNum,omitempty" xml:"ExchangeNum,omitempty"`
	InstanceNum *int32                                    `json:"InstanceNum,omitempty" xml:"InstanceNum,omitempty"`
	Instances   *InstancePreivewResponseBodyDataInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Struct"`
	QueueNum    *int32                                    `json:"QueueNum,omitempty" xml:"QueueNum,omitempty"`
	VhostNum    *int32                                    `json:"VhostNum,omitempty" xml:"VhostNum,omitempty"`
}

func (s InstancePreivewResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s InstancePreivewResponseBodyData) GoString() string {
	return s.String()
}

func (s *InstancePreivewResponseBodyData) GetExchangeNum() *int32 {
	return s.ExchangeNum
}

func (s *InstancePreivewResponseBodyData) GetInstanceNum() *int32 {
	return s.InstanceNum
}

func (s *InstancePreivewResponseBodyData) GetInstances() *InstancePreivewResponseBodyDataInstances {
	return s.Instances
}

func (s *InstancePreivewResponseBodyData) GetQueueNum() *int32 {
	return s.QueueNum
}

func (s *InstancePreivewResponseBodyData) GetVhostNum() *int32 {
	return s.VhostNum
}

func (s *InstancePreivewResponseBodyData) SetExchangeNum(v int32) *InstancePreivewResponseBodyData {
	s.ExchangeNum = &v
	return s
}

func (s *InstancePreivewResponseBodyData) SetInstanceNum(v int32) *InstancePreivewResponseBodyData {
	s.InstanceNum = &v
	return s
}

func (s *InstancePreivewResponseBodyData) SetInstances(v *InstancePreivewResponseBodyDataInstances) *InstancePreivewResponseBodyData {
	s.Instances = v
	return s
}

func (s *InstancePreivewResponseBodyData) SetQueueNum(v int32) *InstancePreivewResponseBodyData {
	s.QueueNum = &v
	return s
}

func (s *InstancePreivewResponseBodyData) SetVhostNum(v int32) *InstancePreivewResponseBodyData {
	s.VhostNum = &v
	return s
}

func (s *InstancePreivewResponseBodyData) Validate() error {
	if s.Instances != nil {
		if err := s.Instances.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type InstancePreivewResponseBodyDataInstances struct {
	InstancesVO []*InstancePreivewResponseBodyDataInstancesInstancesVO `json:"InstancesVO,omitempty" xml:"InstancesVO,omitempty" type:"Repeated"`
}

func (s InstancePreivewResponseBodyDataInstances) String() string {
	return dara.Prettify(s)
}

func (s InstancePreivewResponseBodyDataInstances) GoString() string {
	return s.String()
}

func (s *InstancePreivewResponseBodyDataInstances) GetInstancesVO() []*InstancePreivewResponseBodyDataInstancesInstancesVO {
	return s.InstancesVO
}

func (s *InstancePreivewResponseBodyDataInstances) SetInstancesVO(v []*InstancePreivewResponseBodyDataInstancesInstancesVO) *InstancePreivewResponseBodyDataInstances {
	s.InstancesVO = v
	return s
}

func (s *InstancePreivewResponseBodyDataInstances) Validate() error {
	if s.InstancesVO != nil {
		for _, item := range s.InstancesVO {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type InstancePreivewResponseBodyDataInstancesInstancesVO struct {
	AutoRenew                 *bool                                                    `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	CeaseStatus               *bool                                                    `json:"CeaseStatus,omitempty" xml:"CeaseStatus,omitempty"`
	ClassicEndpoint           *string                                                  `json:"ClassicEndpoint,omitempty" xml:"ClassicEndpoint,omitempty"`
	EnableDlqTtl              *bool                                                    `json:"EnableDlqTtl,omitempty" xml:"EnableDlqTtl,omitempty"`
	Encrypted                 *bool                                                    `json:"Encrypted,omitempty" xml:"Encrypted,omitempty"`
	Expire                    *int64                                                   `json:"Expire,omitempty" xml:"Expire,omitempty"`
	InstanceId                *string                                                  `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceName              *string                                                  `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	InstanceType              *string                                                  `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	InvisibleTime             *int32                                                   `json:"InvisibleTime,omitempty" xml:"InvisibleTime,omitempty"`
	KmsKeyId                  *string                                                  `json:"KmsKeyId,omitempty" xml:"KmsKeyId,omitempty"`
	MaxBindingCount           *int32                                                   `json:"MaxBindingCount,omitempty" xml:"MaxBindingCount,omitempty"`
	MaxConnectionChannelCount *int32                                                   `json:"MaxConnectionChannelCount,omitempty" xml:"MaxConnectionChannelCount,omitempty"`
	MaxConnectionCount        *int32                                                   `json:"MaxConnectionCount,omitempty" xml:"MaxConnectionCount,omitempty"`
	MaxConsumeRetryTime       *int32                                                   `json:"MaxConsumeRetryTime,omitempty" xml:"MaxConsumeRetryTime,omitempty"`
	MaxEIPTPS                 *int32                                                   `json:"MaxEIPTPS,omitempty" xml:"MaxEIPTPS,omitempty"`
	MaxExchangeCount          *int32                                                   `json:"MaxExchangeCount,omitempty" xml:"MaxExchangeCount,omitempty"`
	MaxMsgBodyByte            *int32                                                   `json:"MaxMsgBodyByte,omitempty" xml:"MaxMsgBodyByte,omitempty"`
	MaxMsgDelayHour           *int32                                                   `json:"MaxMsgDelayHour,omitempty" xml:"MaxMsgDelayHour,omitempty"`
	MaxMsgTraceTime           *int32                                                   `json:"MaxMsgTraceTime,omitempty" xml:"MaxMsgTraceTime,omitempty"`
	MaxQueue                  *int32                                                   `json:"MaxQueue,omitempty" xml:"MaxQueue,omitempty"`
	MaxQueueConsumerCount     *int32                                                   `json:"MaxQueueConsumerCount,omitempty" xml:"MaxQueueConsumerCount,omitempty"`
	MaxRetryInterval          *int32                                                   `json:"MaxRetryInterval,omitempty" xml:"MaxRetryInterval,omitempty"`
	MaxRetryTimes             *int32                                                   `json:"MaxRetryTimes,omitempty" xml:"MaxRetryTimes,omitempty"`
	MaxTPS                    *int32                                                   `json:"MaxTPS,omitempty" xml:"MaxTPS,omitempty"`
	MaxVhost                  *int32                                                   `json:"MaxVhost,omitempty" xml:"MaxVhost,omitempty"`
	OrderCreate               *int64                                                   `json:"OrderCreate,omitempty" xml:"OrderCreate,omitempty"`
	OrderType                 *string                                                  `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	PrivateEndpoint           *string                                                  `json:"PrivateEndpoint,omitempty" xml:"PrivateEndpoint,omitempty"`
	PublicEndpoint            *string                                                  `json:"PublicEndpoint,omitempty" xml:"PublicEndpoint,omitempty"`
	ResourceGroupId           *string                                                  `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ServerlessRate            *float64                                                 `json:"ServerlessRate,omitempty" xml:"ServerlessRate,omitempty"`
	ServerlessSwitch          *bool                                                    `json:"ServerlessSwitch,omitempty" xml:"ServerlessSwitch,omitempty"`
	Status                    *string                                                  `json:"Status,omitempty" xml:"Status,omitempty"`
	StorageSize               *int32                                                   `json:"StorageSize,omitempty" xml:"StorageSize,omitempty"`
	SupportEIP                *bool                                                    `json:"SupportEIP,omitempty" xml:"SupportEIP,omitempty"`
	SupportMsgTrace           *bool                                                    `json:"SupportMsgTrace,omitempty" xml:"SupportMsgTrace,omitempty"`
	SupportOpenSourceAuth     *bool                                                    `json:"SupportOpenSourceAuth,omitempty" xml:"SupportOpenSourceAuth,omitempty"`
	Tags                      *InstancePreivewResponseBodyDataInstancesInstancesVOTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	UsedQueue                 *int32                                                   `json:"UsedQueue,omitempty" xml:"UsedQueue,omitempty"`
	UsedVhost                 *int32                                                   `json:"UsedVhost,omitempty" xml:"UsedVhost,omitempty"`
	Version                   *int32                                                   `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s InstancePreivewResponseBodyDataInstancesInstancesVO) String() string {
	return dara.Prettify(s)
}

func (s InstancePreivewResponseBodyDataInstancesInstancesVO) GoString() string {
	return s.String()
}

func (s *InstancePreivewResponseBodyDataInstancesInstancesVO) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *InstancePreivewResponseBodyDataInstancesInstancesVO) GetCeaseStatus() *bool {
	return s.CeaseStatus
}

func (s *InstancePreivewResponseBodyDataInstancesInstancesVO) GetClassicEndpoint() *string {
	return s.ClassicEndpoint
}

func (s *InstancePreivewResponseBodyDataInstancesInstancesVO) GetEnableDlqTtl() *bool {
	return s.EnableDlqTtl
}

func (s *InstancePreivewResponseBodyDataInstancesInstancesVO) GetEncrypted() *bool {
	return s.Encrypted
}

func (s *InstancePreivewResponseBodyDataInstancesInstancesVO) GetExpire() *int64 {
	return s.Expire
}

func (s *InstancePreivewResponseBodyDataInstancesInstancesVO) GetInstanceId() *string {
	return s.InstanceId
}

func (s *InstancePreivewResponseBodyDataInstancesInstancesVO) GetInstanceName() *string {
	return s.InstanceName
}

func (s *InstancePreivewResponseBodyDataInstancesInstancesVO) GetInstanceType() *string {
	return s.InstanceType
}

func (s *InstancePreivewResponseBodyDataInstancesInstancesVO) GetInvisibleTime() *int32 {
	return s.InvisibleTime
}

func (s *InstancePreivewResponseBodyDataInstancesInstancesVO) GetKmsKeyId() *string {
	return s.KmsKeyId
}

func (s *InstancePreivewResponseBodyDataInstancesInstancesVO) GetMaxBindingCount() *int32 {
	return s.MaxBindingCount
}

func (s *InstancePreivewResponseBodyDataInstancesInstancesVO) GetMaxConnectionChannelCount() *int32 {
	return s.MaxConnectionChannelCount
}

func (s *InstancePreivewResponseBodyDataInstancesInstancesVO) GetMaxConnectionCount() *int32 {
	return s.MaxConnectionCount
}

func (s *InstancePreivewResponseBodyDataInstancesInstancesVO) GetMaxConsumeRetryTime() *int32 {
	return s.MaxConsumeRetryTime
}

func (s *InstancePreivewResponseBodyDataInstancesInstancesVO) GetMaxEIPTPS() *int32 {
	return s.MaxEIPTPS
}

func (s *InstancePreivewResponseBodyDataInstancesInstancesVO) GetMaxExchangeCount() *int32 {
	return s.MaxExchangeCount
}

func (s *InstancePreivewResponseBodyDataInstancesInstancesVO) GetMaxMsgBodyByte() *int32 {
	return s.MaxMsgBodyByte
}

func (s *InstancePreivewResponseBodyDataInstancesInstancesVO) GetMaxMsgDelayHour() *int32 {
	return s.MaxMsgDelayHour
}

func (s *InstancePreivewResponseBodyDataInstancesInstancesVO) GetMaxMsgTraceTime() *int32 {
	return s.MaxMsgTraceTime
}

func (s *InstancePreivewResponseBodyDataInstancesInstancesVO) GetMaxQueue() *int32 {
	return s.MaxQueue
}

func (s *InstancePreivewResponseBodyDataInstancesInstancesVO) GetMaxQueueConsumerCount() *int32 {
	return s.MaxQueueConsumerCount
}

func (s *InstancePreivewResponseBodyDataInstancesInstancesVO) GetMaxRetryInterval() *int32 {
	return s.MaxRetryInterval
}

func (s *InstancePreivewResponseBodyDataInstancesInstancesVO) GetMaxRetryTimes() *int32 {
	return s.MaxRetryTimes
}

func (s *InstancePreivewResponseBodyDataInstancesInstancesVO) GetMaxTPS() *int32 {
	return s.MaxTPS
}

func (s *InstancePreivewResponseBodyDataInstancesInstancesVO) GetMaxVhost() *int32 {
	return s.MaxVhost
}

func (s *InstancePreivewResponseBodyDataInstancesInstancesVO) GetOrderCreate() *int64 {
	return s.OrderCreate
}

func (s *InstancePreivewResponseBodyDataInstancesInstancesVO) GetOrderType() *string {
	return s.OrderType
}

func (s *InstancePreivewResponseBodyDataInstancesInstancesVO) GetPrivateEndpoint() *string {
	return s.PrivateEndpoint
}

func (s *InstancePreivewResponseBodyDataInstancesInstancesVO) GetPublicEndpoint() *string {
	return s.PublicEndpoint
}

func (s *InstancePreivewResponseBodyDataInstancesInstancesVO) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *InstancePreivewResponseBodyDataInstancesInstancesVO) GetServerlessRate() *float64 {
	return s.ServerlessRate
}

func (s *InstancePreivewResponseBodyDataInstancesInstancesVO) GetServerlessSwitch() *bool {
	return s.ServerlessSwitch
}

func (s *InstancePreivewResponseBodyDataInstancesInstancesVO) GetStatus() *string {
	return s.Status
}

func (s *InstancePreivewResponseBodyDataInstancesInstancesVO) GetStorageSize() *int32 {
	return s.StorageSize
}

func (s *InstancePreivewResponseBodyDataInstancesInstancesVO) GetSupportEIP() *bool {
	return s.SupportEIP
}

func (s *InstancePreivewResponseBodyDataInstancesInstancesVO) GetSupportMsgTrace() *bool {
	return s.SupportMsgTrace
}

func (s *InstancePreivewResponseBodyDataInstancesInstancesVO) GetSupportOpenSourceAuth() *bool {
	return s.SupportOpenSourceAuth
}

func (s *InstancePreivewResponseBodyDataInstancesInstancesVO) GetTags() *InstancePreivewResponseBodyDataInstancesInstancesVOTags {
	return s.Tags
}

func (s *InstancePreivewResponseBodyDataInstancesInstancesVO) GetUsedQueue() *int32 {
	return s.UsedQueue
}

func (s *InstancePreivewResponseBodyDataInstancesInstancesVO) GetUsedVhost() *int32 {
	return s.UsedVhost
}

func (s *InstancePreivewResponseBodyDataInstancesInstancesVO) GetVersion() *int32 {
	return s.Version
}

func (s *InstancePreivewResponseBodyDataInstancesInstancesVO) SetAutoRenew(v bool) *InstancePreivewResponseBodyDataInstancesInstancesVO {
	s.AutoRenew = &v
	return s
}

func (s *InstancePreivewResponseBodyDataInstancesInstancesVO) SetCeaseStatus(v bool) *InstancePreivewResponseBodyDataInstancesInstancesVO {
	s.CeaseStatus = &v
	return s
}

func (s *InstancePreivewResponseBodyDataInstancesInstancesVO) SetClassicEndpoint(v string) *InstancePreivewResponseBodyDataInstancesInstancesVO {
	s.ClassicEndpoint = &v
	return s
}

func (s *InstancePreivewResponseBodyDataInstancesInstancesVO) SetEnableDlqTtl(v bool) *InstancePreivewResponseBodyDataInstancesInstancesVO {
	s.EnableDlqTtl = &v
	return s
}

func (s *InstancePreivewResponseBodyDataInstancesInstancesVO) SetEncrypted(v bool) *InstancePreivewResponseBodyDataInstancesInstancesVO {
	s.Encrypted = &v
	return s
}

func (s *InstancePreivewResponseBodyDataInstancesInstancesVO) SetExpire(v int64) *InstancePreivewResponseBodyDataInstancesInstancesVO {
	s.Expire = &v
	return s
}

func (s *InstancePreivewResponseBodyDataInstancesInstancesVO) SetInstanceId(v string) *InstancePreivewResponseBodyDataInstancesInstancesVO {
	s.InstanceId = &v
	return s
}

func (s *InstancePreivewResponseBodyDataInstancesInstancesVO) SetInstanceName(v string) *InstancePreivewResponseBodyDataInstancesInstancesVO {
	s.InstanceName = &v
	return s
}

func (s *InstancePreivewResponseBodyDataInstancesInstancesVO) SetInstanceType(v string) *InstancePreivewResponseBodyDataInstancesInstancesVO {
	s.InstanceType = &v
	return s
}

func (s *InstancePreivewResponseBodyDataInstancesInstancesVO) SetInvisibleTime(v int32) *InstancePreivewResponseBodyDataInstancesInstancesVO {
	s.InvisibleTime = &v
	return s
}

func (s *InstancePreivewResponseBodyDataInstancesInstancesVO) SetKmsKeyId(v string) *InstancePreivewResponseBodyDataInstancesInstancesVO {
	s.KmsKeyId = &v
	return s
}

func (s *InstancePreivewResponseBodyDataInstancesInstancesVO) SetMaxBindingCount(v int32) *InstancePreivewResponseBodyDataInstancesInstancesVO {
	s.MaxBindingCount = &v
	return s
}

func (s *InstancePreivewResponseBodyDataInstancesInstancesVO) SetMaxConnectionChannelCount(v int32) *InstancePreivewResponseBodyDataInstancesInstancesVO {
	s.MaxConnectionChannelCount = &v
	return s
}

func (s *InstancePreivewResponseBodyDataInstancesInstancesVO) SetMaxConnectionCount(v int32) *InstancePreivewResponseBodyDataInstancesInstancesVO {
	s.MaxConnectionCount = &v
	return s
}

func (s *InstancePreivewResponseBodyDataInstancesInstancesVO) SetMaxConsumeRetryTime(v int32) *InstancePreivewResponseBodyDataInstancesInstancesVO {
	s.MaxConsumeRetryTime = &v
	return s
}

func (s *InstancePreivewResponseBodyDataInstancesInstancesVO) SetMaxEIPTPS(v int32) *InstancePreivewResponseBodyDataInstancesInstancesVO {
	s.MaxEIPTPS = &v
	return s
}

func (s *InstancePreivewResponseBodyDataInstancesInstancesVO) SetMaxExchangeCount(v int32) *InstancePreivewResponseBodyDataInstancesInstancesVO {
	s.MaxExchangeCount = &v
	return s
}

func (s *InstancePreivewResponseBodyDataInstancesInstancesVO) SetMaxMsgBodyByte(v int32) *InstancePreivewResponseBodyDataInstancesInstancesVO {
	s.MaxMsgBodyByte = &v
	return s
}

func (s *InstancePreivewResponseBodyDataInstancesInstancesVO) SetMaxMsgDelayHour(v int32) *InstancePreivewResponseBodyDataInstancesInstancesVO {
	s.MaxMsgDelayHour = &v
	return s
}

func (s *InstancePreivewResponseBodyDataInstancesInstancesVO) SetMaxMsgTraceTime(v int32) *InstancePreivewResponseBodyDataInstancesInstancesVO {
	s.MaxMsgTraceTime = &v
	return s
}

func (s *InstancePreivewResponseBodyDataInstancesInstancesVO) SetMaxQueue(v int32) *InstancePreivewResponseBodyDataInstancesInstancesVO {
	s.MaxQueue = &v
	return s
}

func (s *InstancePreivewResponseBodyDataInstancesInstancesVO) SetMaxQueueConsumerCount(v int32) *InstancePreivewResponseBodyDataInstancesInstancesVO {
	s.MaxQueueConsumerCount = &v
	return s
}

func (s *InstancePreivewResponseBodyDataInstancesInstancesVO) SetMaxRetryInterval(v int32) *InstancePreivewResponseBodyDataInstancesInstancesVO {
	s.MaxRetryInterval = &v
	return s
}

func (s *InstancePreivewResponseBodyDataInstancesInstancesVO) SetMaxRetryTimes(v int32) *InstancePreivewResponseBodyDataInstancesInstancesVO {
	s.MaxRetryTimes = &v
	return s
}

func (s *InstancePreivewResponseBodyDataInstancesInstancesVO) SetMaxTPS(v int32) *InstancePreivewResponseBodyDataInstancesInstancesVO {
	s.MaxTPS = &v
	return s
}

func (s *InstancePreivewResponseBodyDataInstancesInstancesVO) SetMaxVhost(v int32) *InstancePreivewResponseBodyDataInstancesInstancesVO {
	s.MaxVhost = &v
	return s
}

func (s *InstancePreivewResponseBodyDataInstancesInstancesVO) SetOrderCreate(v int64) *InstancePreivewResponseBodyDataInstancesInstancesVO {
	s.OrderCreate = &v
	return s
}

func (s *InstancePreivewResponseBodyDataInstancesInstancesVO) SetOrderType(v string) *InstancePreivewResponseBodyDataInstancesInstancesVO {
	s.OrderType = &v
	return s
}

func (s *InstancePreivewResponseBodyDataInstancesInstancesVO) SetPrivateEndpoint(v string) *InstancePreivewResponseBodyDataInstancesInstancesVO {
	s.PrivateEndpoint = &v
	return s
}

func (s *InstancePreivewResponseBodyDataInstancesInstancesVO) SetPublicEndpoint(v string) *InstancePreivewResponseBodyDataInstancesInstancesVO {
	s.PublicEndpoint = &v
	return s
}

func (s *InstancePreivewResponseBodyDataInstancesInstancesVO) SetResourceGroupId(v string) *InstancePreivewResponseBodyDataInstancesInstancesVO {
	s.ResourceGroupId = &v
	return s
}

func (s *InstancePreivewResponseBodyDataInstancesInstancesVO) SetServerlessRate(v float64) *InstancePreivewResponseBodyDataInstancesInstancesVO {
	s.ServerlessRate = &v
	return s
}

func (s *InstancePreivewResponseBodyDataInstancesInstancesVO) SetServerlessSwitch(v bool) *InstancePreivewResponseBodyDataInstancesInstancesVO {
	s.ServerlessSwitch = &v
	return s
}

func (s *InstancePreivewResponseBodyDataInstancesInstancesVO) SetStatus(v string) *InstancePreivewResponseBodyDataInstancesInstancesVO {
	s.Status = &v
	return s
}

func (s *InstancePreivewResponseBodyDataInstancesInstancesVO) SetStorageSize(v int32) *InstancePreivewResponseBodyDataInstancesInstancesVO {
	s.StorageSize = &v
	return s
}

func (s *InstancePreivewResponseBodyDataInstancesInstancesVO) SetSupportEIP(v bool) *InstancePreivewResponseBodyDataInstancesInstancesVO {
	s.SupportEIP = &v
	return s
}

func (s *InstancePreivewResponseBodyDataInstancesInstancesVO) SetSupportMsgTrace(v bool) *InstancePreivewResponseBodyDataInstancesInstancesVO {
	s.SupportMsgTrace = &v
	return s
}

func (s *InstancePreivewResponseBodyDataInstancesInstancesVO) SetSupportOpenSourceAuth(v bool) *InstancePreivewResponseBodyDataInstancesInstancesVO {
	s.SupportOpenSourceAuth = &v
	return s
}

func (s *InstancePreivewResponseBodyDataInstancesInstancesVO) SetTags(v *InstancePreivewResponseBodyDataInstancesInstancesVOTags) *InstancePreivewResponseBodyDataInstancesInstancesVO {
	s.Tags = v
	return s
}

func (s *InstancePreivewResponseBodyDataInstancesInstancesVO) SetUsedQueue(v int32) *InstancePreivewResponseBodyDataInstancesInstancesVO {
	s.UsedQueue = &v
	return s
}

func (s *InstancePreivewResponseBodyDataInstancesInstancesVO) SetUsedVhost(v int32) *InstancePreivewResponseBodyDataInstancesInstancesVO {
	s.UsedVhost = &v
	return s
}

func (s *InstancePreivewResponseBodyDataInstancesInstancesVO) SetVersion(v int32) *InstancePreivewResponseBodyDataInstancesInstancesVO {
	s.Version = &v
	return s
}

func (s *InstancePreivewResponseBodyDataInstancesInstancesVO) Validate() error {
	if s.Tags != nil {
		if err := s.Tags.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type InstancePreivewResponseBodyDataInstancesInstancesVOTags struct {
	TagsVO []*InstancePreivewResponseBodyDataInstancesInstancesVOTagsTagsVO `json:"TagsVO,omitempty" xml:"TagsVO,omitempty" type:"Repeated"`
}

func (s InstancePreivewResponseBodyDataInstancesInstancesVOTags) String() string {
	return dara.Prettify(s)
}

func (s InstancePreivewResponseBodyDataInstancesInstancesVOTags) GoString() string {
	return s.String()
}

func (s *InstancePreivewResponseBodyDataInstancesInstancesVOTags) GetTagsVO() []*InstancePreivewResponseBodyDataInstancesInstancesVOTagsTagsVO {
	return s.TagsVO
}

func (s *InstancePreivewResponseBodyDataInstancesInstancesVOTags) SetTagsVO(v []*InstancePreivewResponseBodyDataInstancesInstancesVOTagsTagsVO) *InstancePreivewResponseBodyDataInstancesInstancesVOTags {
	s.TagsVO = v
	return s
}

func (s *InstancePreivewResponseBodyDataInstancesInstancesVOTags) Validate() error {
	if s.TagsVO != nil {
		for _, item := range s.TagsVO {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type InstancePreivewResponseBodyDataInstancesInstancesVOTagsTagsVO struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s InstancePreivewResponseBodyDataInstancesInstancesVOTagsTagsVO) String() string {
	return dara.Prettify(s)
}

func (s InstancePreivewResponseBodyDataInstancesInstancesVOTagsTagsVO) GoString() string {
	return s.String()
}

func (s *InstancePreivewResponseBodyDataInstancesInstancesVOTagsTagsVO) GetKey() *string {
	return s.Key
}

func (s *InstancePreivewResponseBodyDataInstancesInstancesVOTagsTagsVO) GetValue() *string {
	return s.Value
}

func (s *InstancePreivewResponseBodyDataInstancesInstancesVOTagsTagsVO) SetKey(v string) *InstancePreivewResponseBodyDataInstancesInstancesVOTagsTagsVO {
	s.Key = &v
	return s
}

func (s *InstancePreivewResponseBodyDataInstancesInstancesVOTagsTagsVO) SetValue(v string) *InstancePreivewResponseBodyDataInstancesInstancesVOTagsTagsVO {
	s.Value = &v
	return s
}

func (s *InstancePreivewResponseBodyDataInstancesInstancesVOTagsTagsVO) Validate() error {
	return dara.Validate(s)
}
