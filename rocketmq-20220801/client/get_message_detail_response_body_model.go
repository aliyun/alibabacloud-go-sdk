// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMessageDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetMessageDetailResponseBody
	GetCode() *string
	SetData(v *GetMessageDetailResponseBodyData) *GetMessageDetailResponseBody
	GetData() *GetMessageDetailResponseBodyData
	SetDynamicCode(v string) *GetMessageDetailResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *GetMessageDetailResponseBody
	GetDynamicMessage() *string
	SetHttpStatusCode(v int32) *GetMessageDetailResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetMessageDetailResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetMessageDetailResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetMessageDetailResponseBody
	GetSuccess() *bool
}

type GetMessageDetailResponseBody struct {
	// The error code.
	//
	// example:
	//
	// Topic.NotFound
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The returned data.
	Data *GetMessageDetailResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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
	// Parameter instanceId is mandatory for this action .
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// FAEBD71F-E839-52F9-BD7B-8F1290525841
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetMessageDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMessageDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetMessageDetailResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetMessageDetailResponseBody) GetData() *GetMessageDetailResponseBodyData {
	return s.Data
}

func (s *GetMessageDetailResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *GetMessageDetailResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *GetMessageDetailResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetMessageDetailResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetMessageDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMessageDetailResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetMessageDetailResponseBody) SetCode(v string) *GetMessageDetailResponseBody {
	s.Code = &v
	return s
}

func (s *GetMessageDetailResponseBody) SetData(v *GetMessageDetailResponseBodyData) *GetMessageDetailResponseBody {
	s.Data = v
	return s
}

func (s *GetMessageDetailResponseBody) SetDynamicCode(v string) *GetMessageDetailResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *GetMessageDetailResponseBody) SetDynamicMessage(v string) *GetMessageDetailResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *GetMessageDetailResponseBody) SetHttpStatusCode(v int32) *GetMessageDetailResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetMessageDetailResponseBody) SetMessage(v string) *GetMessageDetailResponseBody {
	s.Message = &v
	return s
}

func (s *GetMessageDetailResponseBody) SetRequestId(v string) *GetMessageDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMessageDetailResponseBody) SetSuccess(v bool) *GetMessageDetailResponseBody {
	s.Success = &v
	return s
}

func (s *GetMessageDetailResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetMessageDetailResponseBodyData struct {
	// The message body.
	//
	// example:
	//
	// {}
	Body *string `json:"body,omitempty" xml:"body,omitempty"`
	// The size of the message body.
	//
	// example:
	//
	// 123
	BodySize *int32 `json:"bodySize,omitempty" xml:"bodySize,omitempty"`
	// The client on which the message was produced.
	//
	// example:
	//
	// xxx.xx.xxx.xx
	BornHost *string `json:"bornHost,omitempty" xml:"bornHost,omitempty"`
	// The time when the message was generated.
	//
	// example:
	//
	// 2023-03-22 12:17:08
	BornTime *string `json:"bornTime,omitempty" xml:"bornTime,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// rmq-cn-7e22ody****
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// example:
	//
	// abc
	LiteTopicName *string `json:"liteTopicName,omitempty" xml:"liteTopicName,omitempty"`
	// The sharding key. This parameter is returned only for ordered messages.
	//
	// example:
	//
	// xx
	MessageGroup *string `json:"messageGroup,omitempty" xml:"messageGroup,omitempty"`
	// The message ID.
	//
	// example:
	//
	// 01BE87E485F0C7808C04543CAF00000001
	MessageId *string `json:"messageId,omitempty" xml:"messageId,omitempty"`
	// The message keys.
	MessageKeys []*string `json:"messageKeys,omitempty" xml:"messageKeys,omitempty" type:"Repeated"`
	// The tags.
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
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The broker on which the message was stored.
	//
	// example:
	//
	// xxx.xx.xxx.xx
	StoreHost *string `json:"storeHost,omitempty" xml:"storeHost,omitempty"`
	// The time when the message was stored.
	//
	// example:
	//
	// 2023-03-22 12:17:08
	StoreTime *string `json:"storeTime,omitempty" xml:"storeTime,omitempty"`
	// The default system attributes.
	SystemProperties map[string]*string `json:"systemProperties,omitempty" xml:"systemProperties,omitempty"`
	// The topic name.
	//
	// example:
	//
	// topic_test
	TopicName *string `json:"topicName,omitempty" xml:"topicName,omitempty"`
	// The user attributes.
	UserProperties map[string]*string `json:"userProperties,omitempty" xml:"userProperties,omitempty"`
}

func (s GetMessageDetailResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetMessageDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetMessageDetailResponseBodyData) GetBody() *string {
	return s.Body
}

func (s *GetMessageDetailResponseBodyData) GetBodySize() *int32 {
	return s.BodySize
}

func (s *GetMessageDetailResponseBodyData) GetBornHost() *string {
	return s.BornHost
}

func (s *GetMessageDetailResponseBodyData) GetBornTime() *string {
	return s.BornTime
}

func (s *GetMessageDetailResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetMessageDetailResponseBodyData) GetLiteTopicName() *string {
	return s.LiteTopicName
}

func (s *GetMessageDetailResponseBodyData) GetMessageGroup() *string {
	return s.MessageGroup
}

func (s *GetMessageDetailResponseBodyData) GetMessageId() *string {
	return s.MessageId
}

func (s *GetMessageDetailResponseBodyData) GetMessageKeys() []*string {
	return s.MessageKeys
}

func (s *GetMessageDetailResponseBodyData) GetMessageTag() *string {
	return s.MessageTag
}

func (s *GetMessageDetailResponseBodyData) GetMessageType() *string {
	return s.MessageType
}

func (s *GetMessageDetailResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *GetMessageDetailResponseBodyData) GetStoreHost() *string {
	return s.StoreHost
}

func (s *GetMessageDetailResponseBodyData) GetStoreTime() *string {
	return s.StoreTime
}

func (s *GetMessageDetailResponseBodyData) GetSystemProperties() map[string]*string {
	return s.SystemProperties
}

func (s *GetMessageDetailResponseBodyData) GetTopicName() *string {
	return s.TopicName
}

func (s *GetMessageDetailResponseBodyData) GetUserProperties() map[string]*string {
	return s.UserProperties
}

func (s *GetMessageDetailResponseBodyData) SetBody(v string) *GetMessageDetailResponseBodyData {
	s.Body = &v
	return s
}

func (s *GetMessageDetailResponseBodyData) SetBodySize(v int32) *GetMessageDetailResponseBodyData {
	s.BodySize = &v
	return s
}

func (s *GetMessageDetailResponseBodyData) SetBornHost(v string) *GetMessageDetailResponseBodyData {
	s.BornHost = &v
	return s
}

func (s *GetMessageDetailResponseBodyData) SetBornTime(v string) *GetMessageDetailResponseBodyData {
	s.BornTime = &v
	return s
}

func (s *GetMessageDetailResponseBodyData) SetInstanceId(v string) *GetMessageDetailResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *GetMessageDetailResponseBodyData) SetLiteTopicName(v string) *GetMessageDetailResponseBodyData {
	s.LiteTopicName = &v
	return s
}

func (s *GetMessageDetailResponseBodyData) SetMessageGroup(v string) *GetMessageDetailResponseBodyData {
	s.MessageGroup = &v
	return s
}

func (s *GetMessageDetailResponseBodyData) SetMessageId(v string) *GetMessageDetailResponseBodyData {
	s.MessageId = &v
	return s
}

func (s *GetMessageDetailResponseBodyData) SetMessageKeys(v []*string) *GetMessageDetailResponseBodyData {
	s.MessageKeys = v
	return s
}

func (s *GetMessageDetailResponseBodyData) SetMessageTag(v string) *GetMessageDetailResponseBodyData {
	s.MessageTag = &v
	return s
}

func (s *GetMessageDetailResponseBodyData) SetMessageType(v string) *GetMessageDetailResponseBodyData {
	s.MessageType = &v
	return s
}

func (s *GetMessageDetailResponseBodyData) SetRegionId(v string) *GetMessageDetailResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *GetMessageDetailResponseBodyData) SetStoreHost(v string) *GetMessageDetailResponseBodyData {
	s.StoreHost = &v
	return s
}

func (s *GetMessageDetailResponseBodyData) SetStoreTime(v string) *GetMessageDetailResponseBodyData {
	s.StoreTime = &v
	return s
}

func (s *GetMessageDetailResponseBodyData) SetSystemProperties(v map[string]*string) *GetMessageDetailResponseBodyData {
	s.SystemProperties = v
	return s
}

func (s *GetMessageDetailResponseBodyData) SetTopicName(v string) *GetMessageDetailResponseBodyData {
	s.TopicName = &v
	return s
}

func (s *GetMessageDetailResponseBodyData) SetUserProperties(v map[string]*string) *GetMessageDetailResponseBodyData {
	s.UserProperties = v
	return s
}

func (s *GetMessageDetailResponseBodyData) Validate() error {
	return dara.Validate(s)
}
