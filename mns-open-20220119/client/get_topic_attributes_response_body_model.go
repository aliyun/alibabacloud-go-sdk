// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTopicAttributesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *GetTopicAttributesResponseBody
	GetCode() *int64
	SetData(v *GetTopicAttributesResponseBodyData) *GetTopicAttributesResponseBody
	GetData() *GetTopicAttributesResponseBodyData
	SetMessage(v string) *GetTopicAttributesResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetTopicAttributesResponseBody
	GetRequestId() *string
	SetStatus(v string) *GetTopicAttributesResponseBody
	GetStatus() *string
	SetSuccess(v bool) *GetTopicAttributesResponseBody
	GetSuccess() *bool
}

type GetTopicAttributesResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response data.
	Data *GetTopicAttributesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The response message.
	//
	// example:
	//
	// operation success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 06273500-249F-5863-121D-74D51123****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The response status.
	//
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetTopicAttributesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTopicAttributesResponseBody) GoString() string {
	return s.String()
}

func (s *GetTopicAttributesResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *GetTopicAttributesResponseBody) GetData() *GetTopicAttributesResponseBodyData {
	return s.Data
}

func (s *GetTopicAttributesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetTopicAttributesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTopicAttributesResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetTopicAttributesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetTopicAttributesResponseBody) SetCode(v int64) *GetTopicAttributesResponseBody {
	s.Code = &v
	return s
}

func (s *GetTopicAttributesResponseBody) SetData(v *GetTopicAttributesResponseBodyData) *GetTopicAttributesResponseBody {
	s.Data = v
	return s
}

func (s *GetTopicAttributesResponseBody) SetMessage(v string) *GetTopicAttributesResponseBody {
	s.Message = &v
	return s
}

func (s *GetTopicAttributesResponseBody) SetRequestId(v string) *GetTopicAttributesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTopicAttributesResponseBody) SetStatus(v string) *GetTopicAttributesResponseBody {
	s.Status = &v
	return s
}

func (s *GetTopicAttributesResponseBody) SetSuccess(v bool) *GetTopicAttributesResponseBody {
	s.Success = &v
	return s
}

func (s *GetTopicAttributesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTopicAttributesResponseBodyData struct {
	// The time when the topic was created. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1449554277
	CreateTime        *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	EnableSSE         *bool   `json:"EnableSSE,omitempty" xml:"EnableSSE,omitempty"`
	EncryptionEnabled *bool   `json:"EncryptionEnabled,omitempty" xml:"EncryptionEnabled,omitempty"`
	KmsKeyId          *string `json:"KmsKeyId,omitempty" xml:"KmsKeyId,omitempty"`
	// The most recent time when the topic attributes were modified. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1449554460
	LastModifyTime *int64 `json:"LastModifyTime,omitempty" xml:"LastModifyTime,omitempty"`
	// Indicates whether the Log Management feature is enabled. Valid values:
	//
	// - True: Enabled.
	//
	// - False: Disabled.
	//
	// example:
	//
	// True
	LoggingEnabled *bool `json:"LoggingEnabled,omitempty" xml:"LoggingEnabled,omitempty"`
	// The maximum length of the message body sent to the topic. Unit: bytes.
	//
	// example:
	//
	// 65536
	MaxMessageSize *int64 `json:"MaxMessageSize,omitempty" xml:"MaxMessageSize,omitempty"`
	// The number of messages in the topic.
	//
	// example:
	//
	// 0
	MessageCount *int64 `json:"MessageCount,omitempty" xml:"MessageCount,omitempty"`
	// The maximum duration for which a message is retained in the topic. After the period of time specified by this parameter elapses since the message is sent to the topic, the message is deleted regardless of whether it is successfully pushed to the user. Unit: seconds.
	//
	// example:
	//
	// 86400
	MessageRetentionPeriod *int64  `json:"MessageRetentionPeriod,omitempty" xml:"MessageRetentionPeriod,omitempty"`
	SseAlgorithm           *string `json:"SseAlgorithm,omitempty" xml:"SseAlgorithm,omitempty"`
	SseType                *string `json:"SseType,omitempty" xml:"SseType,omitempty"`
	// The list of resource tags.
	Tags []*GetTopicAttributesResponseBodyDataTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// example:
	//
	// http:// 111111111****.mns.us-west-1-internal.aliyuncs.com/topics/testTopic
	TopicInnerUrl *string `json:"TopicInnerUrl,omitempty" xml:"TopicInnerUrl,omitempty"`
	// The name of the topic.
	//
	// example:
	//
	// demo-topic
	TopicName *string `json:"TopicName,omitempty" xml:"TopicName,omitempty"`
	// The type of the topic. Valid values:
	//
	//    	- normal: normal topic
	//
	//    	- fifo: FIFO topic
	//
	// example:
	//
	// normal
	TopicType *string `json:"TopicType,omitempty" xml:"TopicType,omitempty"`
	// example:
	//
	// http:// 111111111****.mns.us-west-1.aliyuncs.com/topics/testTopic
	TopicUrl *string `json:"TopicUrl,omitempty" xml:"TopicUrl,omitempty"`
}

func (s GetTopicAttributesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetTopicAttributesResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTopicAttributesResponseBodyData) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetTopicAttributesResponseBodyData) GetEnableSSE() *bool {
	return s.EnableSSE
}

func (s *GetTopicAttributesResponseBodyData) GetEncryptionEnabled() *bool {
	return s.EncryptionEnabled
}

func (s *GetTopicAttributesResponseBodyData) GetKmsKeyId() *string {
	return s.KmsKeyId
}

func (s *GetTopicAttributesResponseBodyData) GetLastModifyTime() *int64 {
	return s.LastModifyTime
}

func (s *GetTopicAttributesResponseBodyData) GetLoggingEnabled() *bool {
	return s.LoggingEnabled
}

func (s *GetTopicAttributesResponseBodyData) GetMaxMessageSize() *int64 {
	return s.MaxMessageSize
}

func (s *GetTopicAttributesResponseBodyData) GetMessageCount() *int64 {
	return s.MessageCount
}

func (s *GetTopicAttributesResponseBodyData) GetMessageRetentionPeriod() *int64 {
	return s.MessageRetentionPeriod
}

func (s *GetTopicAttributesResponseBodyData) GetSseAlgorithm() *string {
	return s.SseAlgorithm
}

func (s *GetTopicAttributesResponseBodyData) GetSseType() *string {
	return s.SseType
}

func (s *GetTopicAttributesResponseBodyData) GetTags() []*GetTopicAttributesResponseBodyDataTags {
	return s.Tags
}

func (s *GetTopicAttributesResponseBodyData) GetTopicInnerUrl() *string {
	return s.TopicInnerUrl
}

func (s *GetTopicAttributesResponseBodyData) GetTopicName() *string {
	return s.TopicName
}

func (s *GetTopicAttributesResponseBodyData) GetTopicType() *string {
	return s.TopicType
}

func (s *GetTopicAttributesResponseBodyData) GetTopicUrl() *string {
	return s.TopicUrl
}

func (s *GetTopicAttributesResponseBodyData) SetCreateTime(v int64) *GetTopicAttributesResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetTopicAttributesResponseBodyData) SetEnableSSE(v bool) *GetTopicAttributesResponseBodyData {
	s.EnableSSE = &v
	return s
}

func (s *GetTopicAttributesResponseBodyData) SetEncryptionEnabled(v bool) *GetTopicAttributesResponseBodyData {
	s.EncryptionEnabled = &v
	return s
}

func (s *GetTopicAttributesResponseBodyData) SetKmsKeyId(v string) *GetTopicAttributesResponseBodyData {
	s.KmsKeyId = &v
	return s
}

func (s *GetTopicAttributesResponseBodyData) SetLastModifyTime(v int64) *GetTopicAttributesResponseBodyData {
	s.LastModifyTime = &v
	return s
}

func (s *GetTopicAttributesResponseBodyData) SetLoggingEnabled(v bool) *GetTopicAttributesResponseBodyData {
	s.LoggingEnabled = &v
	return s
}

func (s *GetTopicAttributesResponseBodyData) SetMaxMessageSize(v int64) *GetTopicAttributesResponseBodyData {
	s.MaxMessageSize = &v
	return s
}

func (s *GetTopicAttributesResponseBodyData) SetMessageCount(v int64) *GetTopicAttributesResponseBodyData {
	s.MessageCount = &v
	return s
}

func (s *GetTopicAttributesResponseBodyData) SetMessageRetentionPeriod(v int64) *GetTopicAttributesResponseBodyData {
	s.MessageRetentionPeriod = &v
	return s
}

func (s *GetTopicAttributesResponseBodyData) SetSseAlgorithm(v string) *GetTopicAttributesResponseBodyData {
	s.SseAlgorithm = &v
	return s
}

func (s *GetTopicAttributesResponseBodyData) SetSseType(v string) *GetTopicAttributesResponseBodyData {
	s.SseType = &v
	return s
}

func (s *GetTopicAttributesResponseBodyData) SetTags(v []*GetTopicAttributesResponseBodyDataTags) *GetTopicAttributesResponseBodyData {
	s.Tags = v
	return s
}

func (s *GetTopicAttributesResponseBodyData) SetTopicInnerUrl(v string) *GetTopicAttributesResponseBodyData {
	s.TopicInnerUrl = &v
	return s
}

func (s *GetTopicAttributesResponseBodyData) SetTopicName(v string) *GetTopicAttributesResponseBodyData {
	s.TopicName = &v
	return s
}

func (s *GetTopicAttributesResponseBodyData) SetTopicType(v string) *GetTopicAttributesResponseBodyData {
	s.TopicType = &v
	return s
}

func (s *GetTopicAttributesResponseBodyData) SetTopicUrl(v string) *GetTopicAttributesResponseBodyData {
	s.TopicUrl = &v
	return s
}

func (s *GetTopicAttributesResponseBodyData) Validate() error {
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

type GetTopicAttributesResponseBodyDataTags struct {
	// The key of the tag.
	//
	// example:
	//
	// tag1
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// test
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s GetTopicAttributesResponseBodyDataTags) String() string {
	return dara.Prettify(s)
}

func (s GetTopicAttributesResponseBodyDataTags) GoString() string {
	return s.String()
}

func (s *GetTopicAttributesResponseBodyDataTags) GetTagKey() *string {
	return s.TagKey
}

func (s *GetTopicAttributesResponseBodyDataTags) GetTagValue() *string {
	return s.TagValue
}

func (s *GetTopicAttributesResponseBodyDataTags) SetTagKey(v string) *GetTopicAttributesResponseBodyDataTags {
	s.TagKey = &v
	return s
}

func (s *GetTopicAttributesResponseBodyDataTags) SetTagValue(v string) *GetTopicAttributesResponseBodyDataTags {
	s.TagValue = &v
	return s
}

func (s *GetTopicAttributesResponseBodyDataTags) Validate() error {
	return dara.Validate(s)
}
