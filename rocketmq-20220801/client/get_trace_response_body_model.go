// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTraceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetTraceResponseBody
	GetCode() *string
	SetData(v *GetTraceResponseBodyData) *GetTraceResponseBody
	GetData() *GetTraceResponseBodyData
	SetDynamicCode(v string) *GetTraceResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *GetTraceResponseBody
	GetDynamicMessage() *string
	SetHttpStatusCode(v int32) *GetTraceResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetTraceResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetTraceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetTraceResponseBody
	GetSuccess() *bool
}

type GetTraceResponseBody struct {
	// The error code.
	//
	// example:
	//
	// InvalidConsumerGroupId
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The returned result.
	Data *GetTraceResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The dynamic error code.
	//
	// example:
	//
	// InstanceId
	DynamicCode *string `json:"dynamicCode,omitempty" xml:"dynamicCode,omitempty"`
	// The dynamic error message.
	//
	// example:
	//
	// instanceId
	DynamicMessage *string `json:"dynamicMessage,omitempty" xml:"dynamicMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// The instance cannot be found.
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID. Each request has a unique ID that you can use to troubleshoot issues.
	//
	// example:
	//
	// 7779A8FC-1BCD-5A1D-A603-C4A9BD8ADC49
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetTraceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTraceResponseBody) GoString() string {
	return s.String()
}

func (s *GetTraceResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetTraceResponseBody) GetData() *GetTraceResponseBodyData {
	return s.Data
}

func (s *GetTraceResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *GetTraceResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *GetTraceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetTraceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetTraceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTraceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetTraceResponseBody) SetCode(v string) *GetTraceResponseBody {
	s.Code = &v
	return s
}

func (s *GetTraceResponseBody) SetData(v *GetTraceResponseBodyData) *GetTraceResponseBody {
	s.Data = v
	return s
}

func (s *GetTraceResponseBody) SetDynamicCode(v string) *GetTraceResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *GetTraceResponseBody) SetDynamicMessage(v string) *GetTraceResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *GetTraceResponseBody) SetHttpStatusCode(v int32) *GetTraceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetTraceResponseBody) SetMessage(v string) *GetTraceResponseBody {
	s.Message = &v
	return s
}

func (s *GetTraceResponseBody) SetRequestId(v string) *GetTraceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTraceResponseBody) SetSuccess(v bool) *GetTraceResponseBody {
	s.Success = &v
	return s
}

func (s *GetTraceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTraceResponseBodyData struct {
	// The broker message trace information.
	BrokerInfo *GetTraceResponseBodyDataBrokerInfo `json:"brokerInfo,omitempty" xml:"brokerInfo,omitempty" type:"Struct"`
	// The consumer message trace information.
	ConsumerInfos []*GetTraceResponseBodyDataConsumerInfos `json:"consumerInfos,omitempty" xml:"consumerInfos,omitempty" type:"Repeated"`
	// The instance ID.
	//
	// example:
	//
	// rmq-cn-7e22ody****
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The message information.
	MessageInfo *GetTraceResponseBodyDataMessageInfo `json:"messageInfo,omitempty" xml:"messageInfo,omitempty" type:"Struct"`
	// The producer message trace information.
	ProducerInfo *GetTraceResponseBodyDataProducerInfo `json:"producerInfo,omitempty" xml:"producerInfo,omitempty" type:"Struct"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The topic name.
	//
	// example:
	//
	// topic_test
	TopicName *string `json:"topicName,omitempty" xml:"topicName,omitempty"`
}

func (s GetTraceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetTraceResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTraceResponseBodyData) GetBrokerInfo() *GetTraceResponseBodyDataBrokerInfo {
	return s.BrokerInfo
}

func (s *GetTraceResponseBodyData) GetConsumerInfos() []*GetTraceResponseBodyDataConsumerInfos {
	return s.ConsumerInfos
}

func (s *GetTraceResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetTraceResponseBodyData) GetMessageInfo() *GetTraceResponseBodyDataMessageInfo {
	return s.MessageInfo
}

func (s *GetTraceResponseBodyData) GetProducerInfo() *GetTraceResponseBodyDataProducerInfo {
	return s.ProducerInfo
}

func (s *GetTraceResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *GetTraceResponseBodyData) GetTopicName() *string {
	return s.TopicName
}

func (s *GetTraceResponseBodyData) SetBrokerInfo(v *GetTraceResponseBodyDataBrokerInfo) *GetTraceResponseBodyData {
	s.BrokerInfo = v
	return s
}

func (s *GetTraceResponseBodyData) SetConsumerInfos(v []*GetTraceResponseBodyDataConsumerInfos) *GetTraceResponseBodyData {
	s.ConsumerInfos = v
	return s
}

func (s *GetTraceResponseBodyData) SetInstanceId(v string) *GetTraceResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *GetTraceResponseBodyData) SetMessageInfo(v *GetTraceResponseBodyDataMessageInfo) *GetTraceResponseBodyData {
	s.MessageInfo = v
	return s
}

func (s *GetTraceResponseBodyData) SetProducerInfo(v *GetTraceResponseBodyDataProducerInfo) *GetTraceResponseBodyData {
	s.ProducerInfo = v
	return s
}

func (s *GetTraceResponseBodyData) SetRegionId(v string) *GetTraceResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *GetTraceResponseBodyData) SetTopicName(v string) *GetTraceResponseBodyData {
	s.TopicName = &v
	return s
}

func (s *GetTraceResponseBodyData) Validate() error {
	if s.BrokerInfo != nil {
		if err := s.BrokerInfo.Validate(); err != nil {
			return err
		}
	}
	if s.ConsumerInfos != nil {
		for _, item := range s.ConsumerInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.MessageInfo != nil {
		if err := s.MessageInfo.Validate(); err != nil {
			return err
		}
	}
	if s.ProducerInfo != nil {
		if err := s.ProducerInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTraceResponseBodyDataBrokerInfo struct {
	// The scheduled message status.
	//
	// example:
	//
	// SUCCESS
	DelayStatus *string `json:"delayStatus,omitempty" xml:"delayStatus,omitempty"`
	// The list of operations.
	Operations []*GetTraceResponseBodyDataBrokerInfoOperations `json:"operations,omitempty" xml:"operations,omitempty" type:"Repeated"`
	// The preset delivery time, in YYYY-MM-DD HH:mm:ss format.
	//
	// example:
	//
	// 2023-03-22 12:17:08
	PresetDelayTime *string `json:"presetDelayTime,omitempty" xml:"presetDelayTime,omitempty"`
	// The recall result. This field is non-empty when a recall record exists.
	//
	// example:
	//
	// RECALL_OK | RECALL_FAIL
	RecallResult *string `json:"recallResult,omitempty" xml:"recallResult,omitempty"`
}

func (s GetTraceResponseBodyDataBrokerInfo) String() string {
	return dara.Prettify(s)
}

func (s GetTraceResponseBodyDataBrokerInfo) GoString() string {
	return s.String()
}

func (s *GetTraceResponseBodyDataBrokerInfo) GetDelayStatus() *string {
	return s.DelayStatus
}

func (s *GetTraceResponseBodyDataBrokerInfo) GetOperations() []*GetTraceResponseBodyDataBrokerInfoOperations {
	return s.Operations
}

func (s *GetTraceResponseBodyDataBrokerInfo) GetPresetDelayTime() *string {
	return s.PresetDelayTime
}

func (s *GetTraceResponseBodyDataBrokerInfo) GetRecallResult() *string {
	return s.RecallResult
}

func (s *GetTraceResponseBodyDataBrokerInfo) SetDelayStatus(v string) *GetTraceResponseBodyDataBrokerInfo {
	s.DelayStatus = &v
	return s
}

func (s *GetTraceResponseBodyDataBrokerInfo) SetOperations(v []*GetTraceResponseBodyDataBrokerInfoOperations) *GetTraceResponseBodyDataBrokerInfo {
	s.Operations = v
	return s
}

func (s *GetTraceResponseBodyDataBrokerInfo) SetPresetDelayTime(v string) *GetTraceResponseBodyDataBrokerInfo {
	s.PresetDelayTime = &v
	return s
}

func (s *GetTraceResponseBodyDataBrokerInfo) SetRecallResult(v string) *GetTraceResponseBodyDataBrokerInfo {
	s.RecallResult = &v
	return s
}

func (s *GetTraceResponseBodyDataBrokerInfo) Validate() error {
	if s.Operations != nil {
		for _, item := range s.Operations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetTraceResponseBodyDataBrokerInfoOperations struct {
	// The operation time, in YYYY-MM-DD HH:mm:ss format.
	//
	// example:
	//
	// 2023-03-22 12:17:08
	OperateTime *string `json:"operateTime,omitempty" xml:"operateTime,omitempty"`
	// The operation type.
	//
	// example:
	//
	// ADD
	OperateType *string `json:"operateType,omitempty" xml:"operateType,omitempty"`
}

func (s GetTraceResponseBodyDataBrokerInfoOperations) String() string {
	return dara.Prettify(s)
}

func (s GetTraceResponseBodyDataBrokerInfoOperations) GoString() string {
	return s.String()
}

func (s *GetTraceResponseBodyDataBrokerInfoOperations) GetOperateTime() *string {
	return s.OperateTime
}

func (s *GetTraceResponseBodyDataBrokerInfoOperations) GetOperateType() *string {
	return s.OperateType
}

func (s *GetTraceResponseBodyDataBrokerInfoOperations) SetOperateTime(v string) *GetTraceResponseBodyDataBrokerInfoOperations {
	s.OperateTime = &v
	return s
}

func (s *GetTraceResponseBodyDataBrokerInfoOperations) SetOperateType(v string) *GetTraceResponseBodyDataBrokerInfoOperations {
	s.OperateType = &v
	return s
}

func (s *GetTraceResponseBodyDataBrokerInfoOperations) Validate() error {
	return dara.Validate(s)
}

type GetTraceResponseBodyDataConsumerInfos struct {
	// The consumption status.
	//
	// example:
	//
	// SUCCESS
	ConsumeStatus *string `json:"consumeStatus,omitempty" xml:"consumeStatus,omitempty"`
	// The consumer group ID.
	//
	// example:
	//
	// GID_inspector_group
	ConsumerGroupId *string `json:"consumerGroupId,omitempty" xml:"consumerGroupId,omitempty"`
	// The dead-letter message information.
	DeadLetterInfo *GetTraceResponseBodyDataConsumerInfosDeadLetterInfo `json:"deadLetterInfo,omitempty" xml:"deadLetterInfo,omitempty" type:"Struct"`
	// Indicates whether the message is a dead-letter message.
	//
	// example:
	//
	// true
	DeadMessage *bool `json:"deadMessage,omitempty" xml:"deadMessage,omitempty"`
	// The list of consumption records.
	Records []*GetTraceResponseBodyDataConsumerInfosRecords `json:"records,omitempty" xml:"records,omitempty" type:"Repeated"`
}

func (s GetTraceResponseBodyDataConsumerInfos) String() string {
	return dara.Prettify(s)
}

func (s GetTraceResponseBodyDataConsumerInfos) GoString() string {
	return s.String()
}

func (s *GetTraceResponseBodyDataConsumerInfos) GetConsumeStatus() *string {
	return s.ConsumeStatus
}

func (s *GetTraceResponseBodyDataConsumerInfos) GetConsumerGroupId() *string {
	return s.ConsumerGroupId
}

func (s *GetTraceResponseBodyDataConsumerInfos) GetDeadLetterInfo() *GetTraceResponseBodyDataConsumerInfosDeadLetterInfo {
	return s.DeadLetterInfo
}

func (s *GetTraceResponseBodyDataConsumerInfos) GetDeadMessage() *bool {
	return s.DeadMessage
}

func (s *GetTraceResponseBodyDataConsumerInfos) GetRecords() []*GetTraceResponseBodyDataConsumerInfosRecords {
	return s.Records
}

func (s *GetTraceResponseBodyDataConsumerInfos) SetConsumeStatus(v string) *GetTraceResponseBodyDataConsumerInfos {
	s.ConsumeStatus = &v
	return s
}

func (s *GetTraceResponseBodyDataConsumerInfos) SetConsumerGroupId(v string) *GetTraceResponseBodyDataConsumerInfos {
	s.ConsumerGroupId = &v
	return s
}

func (s *GetTraceResponseBodyDataConsumerInfos) SetDeadLetterInfo(v *GetTraceResponseBodyDataConsumerInfosDeadLetterInfo) *GetTraceResponseBodyDataConsumerInfos {
	s.DeadLetterInfo = v
	return s
}

func (s *GetTraceResponseBodyDataConsumerInfos) SetDeadMessage(v bool) *GetTraceResponseBodyDataConsumerInfos {
	s.DeadMessage = &v
	return s
}

func (s *GetTraceResponseBodyDataConsumerInfos) SetRecords(v []*GetTraceResponseBodyDataConsumerInfosRecords) *GetTraceResponseBodyDataConsumerInfos {
	s.Records = v
	return s
}

func (s *GetTraceResponseBodyDataConsumerInfos) Validate() error {
	if s.DeadLetterInfo != nil {
		if err := s.DeadLetterInfo.Validate(); err != nil {
			return err
		}
	}
	if s.Records != nil {
		for _, item := range s.Records {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetTraceResponseBodyDataConsumerInfosDeadLetterInfo struct {
	// The message ID.
	//
	// example:
	//
	// 7F000001001F7A4F0F29463F0376047D
	MessageId *string `json:"messageId,omitempty" xml:"messageId,omitempty"`
	// The time when the message arrived in the dead-letter queue, in YYYY-MM-DD HH:mm:ss format.
	//
	// example:
	//
	// 2023-03-22 12:17:08
	ToDlqTime *string `json:"toDlqTime,omitempty" xml:"toDlqTime,omitempty"`
	// The topic name.
	//
	// example:
	//
	// Register_Sync
	TopicName *string `json:"topicName,omitempty" xml:"topicName,omitempty"`
}

func (s GetTraceResponseBodyDataConsumerInfosDeadLetterInfo) String() string {
	return dara.Prettify(s)
}

func (s GetTraceResponseBodyDataConsumerInfosDeadLetterInfo) GoString() string {
	return s.String()
}

func (s *GetTraceResponseBodyDataConsumerInfosDeadLetterInfo) GetMessageId() *string {
	return s.MessageId
}

func (s *GetTraceResponseBodyDataConsumerInfosDeadLetterInfo) GetToDlqTime() *string {
	return s.ToDlqTime
}

func (s *GetTraceResponseBodyDataConsumerInfosDeadLetterInfo) GetTopicName() *string {
	return s.TopicName
}

func (s *GetTraceResponseBodyDataConsumerInfosDeadLetterInfo) SetMessageId(v string) *GetTraceResponseBodyDataConsumerInfosDeadLetterInfo {
	s.MessageId = &v
	return s
}

func (s *GetTraceResponseBodyDataConsumerInfosDeadLetterInfo) SetToDlqTime(v string) *GetTraceResponseBodyDataConsumerInfosDeadLetterInfo {
	s.ToDlqTime = &v
	return s
}

func (s *GetTraceResponseBodyDataConsumerInfosDeadLetterInfo) SetTopicName(v string) *GetTraceResponseBodyDataConsumerInfosDeadLetterInfo {
	s.TopicName = &v
	return s
}

func (s *GetTraceResponseBodyDataConsumerInfosDeadLetterInfo) Validate() error {
	return dara.Validate(s)
}

type GetTraceResponseBodyDataConsumerInfosRecords struct {
	// The hostname.
	//
	// example:
	//
	// xx.xx.xx.xx
	ClientHost *string `json:"clientHost,omitempty" xml:"clientHost,omitempty"`
	// The consumption status.
	//
	// example:
	//
	// SUCCESS
	ConsumeStatus *string `json:"consumeStatus,omitempty" xml:"consumeStatus,omitempty"`
	// Indicates whether ordered consumption is enabled.
	//
	// example:
	//
	// true
	FifoEnable *bool `json:"fifoEnable,omitempty" xml:"fifoEnable,omitempty"`
	// The list of operations.
	Operations []*GetTraceResponseBodyDataConsumerInfosRecordsOperations `json:"operations,omitempty" xml:"operations,omitempty" type:"Repeated"`
	// POP_CK
	//
	// example:
	//
	// 123
	PopCk *string `json:"popCk,omitempty" xml:"popCk,omitempty"`
	// The consumer name.
	//
	// example:
	//
	// test
	UserName *string `json:"userName,omitempty" xml:"userName,omitempty"`
}

func (s GetTraceResponseBodyDataConsumerInfosRecords) String() string {
	return dara.Prettify(s)
}

func (s GetTraceResponseBodyDataConsumerInfosRecords) GoString() string {
	return s.String()
}

func (s *GetTraceResponseBodyDataConsumerInfosRecords) GetClientHost() *string {
	return s.ClientHost
}

func (s *GetTraceResponseBodyDataConsumerInfosRecords) GetConsumeStatus() *string {
	return s.ConsumeStatus
}

func (s *GetTraceResponseBodyDataConsumerInfosRecords) GetFifoEnable() *bool {
	return s.FifoEnable
}

func (s *GetTraceResponseBodyDataConsumerInfosRecords) GetOperations() []*GetTraceResponseBodyDataConsumerInfosRecordsOperations {
	return s.Operations
}

func (s *GetTraceResponseBodyDataConsumerInfosRecords) GetPopCk() *string {
	return s.PopCk
}

func (s *GetTraceResponseBodyDataConsumerInfosRecords) GetUserName() *string {
	return s.UserName
}

func (s *GetTraceResponseBodyDataConsumerInfosRecords) SetClientHost(v string) *GetTraceResponseBodyDataConsumerInfosRecords {
	s.ClientHost = &v
	return s
}

func (s *GetTraceResponseBodyDataConsumerInfosRecords) SetConsumeStatus(v string) *GetTraceResponseBodyDataConsumerInfosRecords {
	s.ConsumeStatus = &v
	return s
}

func (s *GetTraceResponseBodyDataConsumerInfosRecords) SetFifoEnable(v bool) *GetTraceResponseBodyDataConsumerInfosRecords {
	s.FifoEnable = &v
	return s
}

func (s *GetTraceResponseBodyDataConsumerInfosRecords) SetOperations(v []*GetTraceResponseBodyDataConsumerInfosRecordsOperations) *GetTraceResponseBodyDataConsumerInfosRecords {
	s.Operations = v
	return s
}

func (s *GetTraceResponseBodyDataConsumerInfosRecords) SetPopCk(v string) *GetTraceResponseBodyDataConsumerInfosRecords {
	s.PopCk = &v
	return s
}

func (s *GetTraceResponseBodyDataConsumerInfosRecords) SetUserName(v string) *GetTraceResponseBodyDataConsumerInfosRecords {
	s.UserName = &v
	return s
}

func (s *GetTraceResponseBodyDataConsumerInfosRecords) Validate() error {
	if s.Operations != nil {
		for _, item := range s.Operations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetTraceResponseBodyDataConsumerInfosRecordsOperations struct {
	// Indicates whether the message is a dead-letter message.
	//
	// example:
	//
	// true
	DeadMessage *bool `json:"deadMessage,omitempty" xml:"deadMessage,omitempty"`
	// The invisible time, in milliseconds.
	//
	// example:
	//
	// 100
	InvisibleTime *int64 `json:"invisibleTime,omitempty" xml:"invisibleTime,omitempty"`
	// The operation time, in YYYY-MM-DD HH:mm:ss format.
	//
	// example:
	//
	// 2023-03-22 12:17:08
	OperateTime *string `json:"operateTime,omitempty" xml:"operateTime,omitempty"`
	// The operation type.
	//
	// example:
	//
	// ADD
	OperateType *string `json:"operateType,omitempty" xml:"operateType,omitempty"`
}

func (s GetTraceResponseBodyDataConsumerInfosRecordsOperations) String() string {
	return dara.Prettify(s)
}

func (s GetTraceResponseBodyDataConsumerInfosRecordsOperations) GoString() string {
	return s.String()
}

func (s *GetTraceResponseBodyDataConsumerInfosRecordsOperations) GetDeadMessage() *bool {
	return s.DeadMessage
}

func (s *GetTraceResponseBodyDataConsumerInfosRecordsOperations) GetInvisibleTime() *int64 {
	return s.InvisibleTime
}

func (s *GetTraceResponseBodyDataConsumerInfosRecordsOperations) GetOperateTime() *string {
	return s.OperateTime
}

func (s *GetTraceResponseBodyDataConsumerInfosRecordsOperations) GetOperateType() *string {
	return s.OperateType
}

func (s *GetTraceResponseBodyDataConsumerInfosRecordsOperations) SetDeadMessage(v bool) *GetTraceResponseBodyDataConsumerInfosRecordsOperations {
	s.DeadMessage = &v
	return s
}

func (s *GetTraceResponseBodyDataConsumerInfosRecordsOperations) SetInvisibleTime(v int64) *GetTraceResponseBodyDataConsumerInfosRecordsOperations {
	s.InvisibleTime = &v
	return s
}

func (s *GetTraceResponseBodyDataConsumerInfosRecordsOperations) SetOperateTime(v string) *GetTraceResponseBodyDataConsumerInfosRecordsOperations {
	s.OperateTime = &v
	return s
}

func (s *GetTraceResponseBodyDataConsumerInfosRecordsOperations) SetOperateType(v string) *GetTraceResponseBodyDataConsumerInfosRecordsOperations {
	s.OperateType = &v
	return s
}

func (s *GetTraceResponseBodyDataConsumerInfosRecordsOperations) Validate() error {
	return dara.Validate(s)
}

type GetTraceResponseBodyDataMessageInfo struct {
	// The message body.
	//
	// example:
	//
	// {}
	Body *string `json:"body,omitempty" xml:"body,omitempty"`
	// The origin of the message.
	//
	// example:
	//
	// x.x.x.x
	BornHost *string `json:"bornHost,omitempty" xml:"bornHost,omitempty"`
	// The time when the message was generated, in YYYY-MM-DD HH:mm:ss format.
	//
	// example:
	//
	// 2023-03-22 12:17:08
	BornTime *string `json:"bornTime,omitempty" xml:"bornTime,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// rmq-cn-u0t2ygjq505
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The lite topic name.
	//
	// example:
	//
	// abc
	LiteTopicName *string `json:"liteTopicName,omitempty" xml:"liteTopicName,omitempty"`
	// The previous sharding key. This field is exclusive to ordered messages.
	//
	// example:
	//
	// xx
	MessageGroup *string `json:"messageGroup,omitempty" xml:"messageGroup,omitempty"`
	// The message ID.
	//
	// example:
	//
	// 0A79275A00207A4F0F2916C92F9A0B94
	MessageId *string `json:"messageId,omitempty" xml:"messageId,omitempty"`
	// The message key.
	MessageKeys []*string `json:"messageKeys,omitempty" xml:"messageKeys,omitempty" type:"Repeated"`
	// The list of tags.
	//
	// example:
	//
	// xx
	MessageTag *string `json:"messageTag,omitempty" xml:"messageTag,omitempty"`
	// The message type.
	//
	// example:
	//
	// NORMAL
	MessageType *string `json:"messageType,omitempty" xml:"messageType,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The log broker address where the message is stored.
	//
	// example:
	//
	// x.x.x.x
	StoreHost *string `json:"storeHost,omitempty" xml:"storeHost,omitempty"`
	// The time when the message was stored in the log broker, in YYYY-MM-DD HH:mm:ss format.
	//
	// example:
	//
	// 2023-03-22 12:17:08
	StoreTime *string `json:"storeTime,omitempty" xml:"storeTime,omitempty"`
	// The topic name.
	//
	// example:
	//
	// Topic_normal_inspector
	TopicName *string `json:"topicName,omitempty" xml:"topicName,omitempty"`
	// The transaction ID.
	//
	// example:
	//
	// xx
	TransactionId *string `json:"transactionId,omitempty" xml:"transactionId,omitempty"`
	// The user properties.
	UserProperties map[string]*string `json:"userProperties,omitempty" xml:"userProperties,omitempty"`
}

func (s GetTraceResponseBodyDataMessageInfo) String() string {
	return dara.Prettify(s)
}

func (s GetTraceResponseBodyDataMessageInfo) GoString() string {
	return s.String()
}

func (s *GetTraceResponseBodyDataMessageInfo) GetBody() *string {
	return s.Body
}

func (s *GetTraceResponseBodyDataMessageInfo) GetBornHost() *string {
	return s.BornHost
}

func (s *GetTraceResponseBodyDataMessageInfo) GetBornTime() *string {
	return s.BornTime
}

func (s *GetTraceResponseBodyDataMessageInfo) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetTraceResponseBodyDataMessageInfo) GetLiteTopicName() *string {
	return s.LiteTopicName
}

func (s *GetTraceResponseBodyDataMessageInfo) GetMessageGroup() *string {
	return s.MessageGroup
}

func (s *GetTraceResponseBodyDataMessageInfo) GetMessageId() *string {
	return s.MessageId
}

func (s *GetTraceResponseBodyDataMessageInfo) GetMessageKeys() []*string {
	return s.MessageKeys
}

func (s *GetTraceResponseBodyDataMessageInfo) GetMessageTag() *string {
	return s.MessageTag
}

func (s *GetTraceResponseBodyDataMessageInfo) GetMessageType() *string {
	return s.MessageType
}

func (s *GetTraceResponseBodyDataMessageInfo) GetRegionId() *string {
	return s.RegionId
}

func (s *GetTraceResponseBodyDataMessageInfo) GetStoreHost() *string {
	return s.StoreHost
}

func (s *GetTraceResponseBodyDataMessageInfo) GetStoreTime() *string {
	return s.StoreTime
}

func (s *GetTraceResponseBodyDataMessageInfo) GetTopicName() *string {
	return s.TopicName
}

func (s *GetTraceResponseBodyDataMessageInfo) GetTransactionId() *string {
	return s.TransactionId
}

func (s *GetTraceResponseBodyDataMessageInfo) GetUserProperties() map[string]*string {
	return s.UserProperties
}

func (s *GetTraceResponseBodyDataMessageInfo) SetBody(v string) *GetTraceResponseBodyDataMessageInfo {
	s.Body = &v
	return s
}

func (s *GetTraceResponseBodyDataMessageInfo) SetBornHost(v string) *GetTraceResponseBodyDataMessageInfo {
	s.BornHost = &v
	return s
}

func (s *GetTraceResponseBodyDataMessageInfo) SetBornTime(v string) *GetTraceResponseBodyDataMessageInfo {
	s.BornTime = &v
	return s
}

func (s *GetTraceResponseBodyDataMessageInfo) SetInstanceId(v string) *GetTraceResponseBodyDataMessageInfo {
	s.InstanceId = &v
	return s
}

func (s *GetTraceResponseBodyDataMessageInfo) SetLiteTopicName(v string) *GetTraceResponseBodyDataMessageInfo {
	s.LiteTopicName = &v
	return s
}

func (s *GetTraceResponseBodyDataMessageInfo) SetMessageGroup(v string) *GetTraceResponseBodyDataMessageInfo {
	s.MessageGroup = &v
	return s
}

func (s *GetTraceResponseBodyDataMessageInfo) SetMessageId(v string) *GetTraceResponseBodyDataMessageInfo {
	s.MessageId = &v
	return s
}

func (s *GetTraceResponseBodyDataMessageInfo) SetMessageKeys(v []*string) *GetTraceResponseBodyDataMessageInfo {
	s.MessageKeys = v
	return s
}

func (s *GetTraceResponseBodyDataMessageInfo) SetMessageTag(v string) *GetTraceResponseBodyDataMessageInfo {
	s.MessageTag = &v
	return s
}

func (s *GetTraceResponseBodyDataMessageInfo) SetMessageType(v string) *GetTraceResponseBodyDataMessageInfo {
	s.MessageType = &v
	return s
}

func (s *GetTraceResponseBodyDataMessageInfo) SetRegionId(v string) *GetTraceResponseBodyDataMessageInfo {
	s.RegionId = &v
	return s
}

func (s *GetTraceResponseBodyDataMessageInfo) SetStoreHost(v string) *GetTraceResponseBodyDataMessageInfo {
	s.StoreHost = &v
	return s
}

func (s *GetTraceResponseBodyDataMessageInfo) SetStoreTime(v string) *GetTraceResponseBodyDataMessageInfo {
	s.StoreTime = &v
	return s
}

func (s *GetTraceResponseBodyDataMessageInfo) SetTopicName(v string) *GetTraceResponseBodyDataMessageInfo {
	s.TopicName = &v
	return s
}

func (s *GetTraceResponseBodyDataMessageInfo) SetTransactionId(v string) *GetTraceResponseBodyDataMessageInfo {
	s.TransactionId = &v
	return s
}

func (s *GetTraceResponseBodyDataMessageInfo) SetUserProperties(v map[string]*string) *GetTraceResponseBodyDataMessageInfo {
	s.UserProperties = v
	return s
}

func (s *GetTraceResponseBodyDataMessageInfo) Validate() error {
	return dara.Validate(s)
}

type GetTraceResponseBodyDataProducerInfo struct {
	// The list of producer records.
	Records []*GetTraceResponseBodyDataProducerInfoRecords `json:"records,omitempty" xml:"records,omitempty" type:"Repeated"`
}

func (s GetTraceResponseBodyDataProducerInfo) String() string {
	return dara.Prettify(s)
}

func (s GetTraceResponseBodyDataProducerInfo) GoString() string {
	return s.String()
}

func (s *GetTraceResponseBodyDataProducerInfo) GetRecords() []*GetTraceResponseBodyDataProducerInfoRecords {
	return s.Records
}

func (s *GetTraceResponseBodyDataProducerInfo) SetRecords(v []*GetTraceResponseBodyDataProducerInfoRecords) *GetTraceResponseBodyDataProducerInfo {
	s.Records = v
	return s
}

func (s *GetTraceResponseBodyDataProducerInfo) Validate() error {
	if s.Records != nil {
		for _, item := range s.Records {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetTraceResponseBodyDataProducerInfoRecords struct {
	// The time when the message arrived, in YYYY-MM-DD HH:mm:ss format.
	//
	// example:
	//
	// 2023-03-22 12:17:08
	ArriveTime *string `json:"arriveTime,omitempty" xml:"arriveTime,omitempty"`
	// The hostname.
	//
	// example:
	//
	// xx.xx.xx.xx
	ClientHost *string `json:"clientHost,omitempty" xml:"clientHost,omitempty"`
	// The dead-letter queue message ID.
	//
	// example:
	//
	// 0A79275A00207A4F0F2916C92F9A0B94
	DlqOriginMessageId *string `json:"dlqOriginMessageId,omitempty" xml:"dlqOriginMessageId,omitempty"`
	// The dead-letter queue topic.
	//
	// example:
	//
	// test_topic
	DlqOriginTopic *string `json:"dlqOriginTopic,omitempty" xml:"dlqOriginTopic,omitempty"`
	// The message source.
	//
	// example:
	//
	// CONSOLE
	MessageSource *string `json:"messageSource,omitempty" xml:"messageSource,omitempty"`
	// The time taken to send the message.
	//
	// example:
	//
	// 100
	ProduceDuration *int64 `json:"produceDuration,omitempty" xml:"produceDuration,omitempty"`
	// The send status.
	//
	// example:
	//
	// SUCCESS
	ProduceStatus *string `json:"produceStatus,omitempty" xml:"produceStatus,omitempty"`
	// The time when the message was sent, in YYYY-MM-DD HH:mm:ss format.
	//
	// example:
	//
	// 2023-03-22 12:17:08
	ProduceTime *string `json:"produceTime,omitempty" xml:"produceTime,omitempty"`
	// The time when the recall request was initiated. This field is non-empty when messageSource is "recall".
	//
	// example:
	//
	// 2023-03-22 12:17:08
	RecallTime *string `json:"recallTime,omitempty" xml:"recallTime,omitempty"`
	// The producer name.
	//
	// example:
	//
	// xxx
	UserName *string `json:"userName,omitempty" xml:"userName,omitempty"`
}

func (s GetTraceResponseBodyDataProducerInfoRecords) String() string {
	return dara.Prettify(s)
}

func (s GetTraceResponseBodyDataProducerInfoRecords) GoString() string {
	return s.String()
}

func (s *GetTraceResponseBodyDataProducerInfoRecords) GetArriveTime() *string {
	return s.ArriveTime
}

func (s *GetTraceResponseBodyDataProducerInfoRecords) GetClientHost() *string {
	return s.ClientHost
}

func (s *GetTraceResponseBodyDataProducerInfoRecords) GetDlqOriginMessageId() *string {
	return s.DlqOriginMessageId
}

func (s *GetTraceResponseBodyDataProducerInfoRecords) GetDlqOriginTopic() *string {
	return s.DlqOriginTopic
}

func (s *GetTraceResponseBodyDataProducerInfoRecords) GetMessageSource() *string {
	return s.MessageSource
}

func (s *GetTraceResponseBodyDataProducerInfoRecords) GetProduceDuration() *int64 {
	return s.ProduceDuration
}

func (s *GetTraceResponseBodyDataProducerInfoRecords) GetProduceStatus() *string {
	return s.ProduceStatus
}

func (s *GetTraceResponseBodyDataProducerInfoRecords) GetProduceTime() *string {
	return s.ProduceTime
}

func (s *GetTraceResponseBodyDataProducerInfoRecords) GetRecallTime() *string {
	return s.RecallTime
}

func (s *GetTraceResponseBodyDataProducerInfoRecords) GetUserName() *string {
	return s.UserName
}

func (s *GetTraceResponseBodyDataProducerInfoRecords) SetArriveTime(v string) *GetTraceResponseBodyDataProducerInfoRecords {
	s.ArriveTime = &v
	return s
}

func (s *GetTraceResponseBodyDataProducerInfoRecords) SetClientHost(v string) *GetTraceResponseBodyDataProducerInfoRecords {
	s.ClientHost = &v
	return s
}

func (s *GetTraceResponseBodyDataProducerInfoRecords) SetDlqOriginMessageId(v string) *GetTraceResponseBodyDataProducerInfoRecords {
	s.DlqOriginMessageId = &v
	return s
}

func (s *GetTraceResponseBodyDataProducerInfoRecords) SetDlqOriginTopic(v string) *GetTraceResponseBodyDataProducerInfoRecords {
	s.DlqOriginTopic = &v
	return s
}

func (s *GetTraceResponseBodyDataProducerInfoRecords) SetMessageSource(v string) *GetTraceResponseBodyDataProducerInfoRecords {
	s.MessageSource = &v
	return s
}

func (s *GetTraceResponseBodyDataProducerInfoRecords) SetProduceDuration(v int64) *GetTraceResponseBodyDataProducerInfoRecords {
	s.ProduceDuration = &v
	return s
}

func (s *GetTraceResponseBodyDataProducerInfoRecords) SetProduceStatus(v string) *GetTraceResponseBodyDataProducerInfoRecords {
	s.ProduceStatus = &v
	return s
}

func (s *GetTraceResponseBodyDataProducerInfoRecords) SetProduceTime(v string) *GetTraceResponseBodyDataProducerInfoRecords {
	s.ProduceTime = &v
	return s
}

func (s *GetTraceResponseBodyDataProducerInfoRecords) SetRecallTime(v string) *GetTraceResponseBodyDataProducerInfoRecords {
	s.RecallTime = &v
	return s
}

func (s *GetTraceResponseBodyDataProducerInfoRecords) SetUserName(v string) *GetTraceResponseBodyDataProducerInfoRecords {
	s.UserName = &v
	return s
}

func (s *GetTraceResponseBodyDataProducerInfoRecords) Validate() error {
	return dara.Validate(s)
}
