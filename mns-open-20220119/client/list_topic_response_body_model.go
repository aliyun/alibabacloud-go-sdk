// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTopicResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *ListTopicResponseBody
	GetCode() *int64
	SetData(v *ListTopicResponseBodyData) *ListTopicResponseBody
	GetData() *ListTopicResponseBodyData
	SetMessage(v string) *ListTopicResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListTopicResponseBody
	GetRequestId() *string
	SetStatus(v string) *ListTopicResponseBody
	GetStatus() *string
	SetSuccess(v bool) *ListTopicResponseBody
	GetSuccess() *bool
}

type ListTopicResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response data.
	Data *ListTopicResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The response message.
	//
	// example:
	//
	// operation success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 06273500-249F-5863-121D-74D51123****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the response.
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

func (s ListTopicResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTopicResponseBody) GoString() string {
	return s.String()
}

func (s *ListTopicResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *ListTopicResponseBody) GetData() *ListTopicResponseBodyData {
	return s.Data
}

func (s *ListTopicResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListTopicResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTopicResponseBody) GetStatus() *string {
	return s.Status
}

func (s *ListTopicResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListTopicResponseBody) SetCode(v int64) *ListTopicResponseBody {
	s.Code = &v
	return s
}

func (s *ListTopicResponseBody) SetData(v *ListTopicResponseBodyData) *ListTopicResponseBody {
	s.Data = v
	return s
}

func (s *ListTopicResponseBody) SetMessage(v string) *ListTopicResponseBody {
	s.Message = &v
	return s
}

func (s *ListTopicResponseBody) SetRequestId(v string) *ListTopicResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTopicResponseBody) SetStatus(v string) *ListTopicResponseBody {
	s.Status = &v
	return s
}

func (s *ListTopicResponseBody) SetSuccess(v bool) *ListTopicResponseBody {
	s.Success = &v
	return s
}

func (s *ListTopicResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListTopicResponseBodyData struct {
	// The results returned on the current page.
	PageData []*ListTopicResponseBodyDataPageData `json:"PageData,omitempty" xml:"PageData,omitempty" type:"Repeated"`
	// The page number of the returned results.
	//
	// example:
	//
	// 1
	PageNum *int64 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of results returned per page.
	//
	// example:
	//
	// 50
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of results.
	//
	// example:
	//
	// 130
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListTopicResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListTopicResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListTopicResponseBodyData) GetPageData() []*ListTopicResponseBodyDataPageData {
	return s.PageData
}

func (s *ListTopicResponseBodyData) GetPageNum() *int64 {
	return s.PageNum
}

func (s *ListTopicResponseBodyData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListTopicResponseBodyData) GetTotal() *int64 {
	return s.Total
}

func (s *ListTopicResponseBodyData) SetPageData(v []*ListTopicResponseBodyDataPageData) *ListTopicResponseBodyData {
	s.PageData = v
	return s
}

func (s *ListTopicResponseBodyData) SetPageNum(v int64) *ListTopicResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *ListTopicResponseBodyData) SetPageSize(v int64) *ListTopicResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListTopicResponseBodyData) SetTotal(v int64) *ListTopicResponseBodyData {
	s.Total = &v
	return s
}

func (s *ListTopicResponseBodyData) Validate() error {
	if s.PageData != nil {
		for _, item := range s.PageData {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTopicResponseBodyDataPageData struct {
	// The time when the subscription was created. The value is a UNIX timestamp representing the number of seconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1449554962
	CreateTime        *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	EnableSSE         *bool   `json:"EnableSSE,omitempty" xml:"EnableSSE,omitempty"`
	EncryptionEnabled *bool   `json:"EncryptionEnabled,omitempty" xml:"EncryptionEnabled,omitempty"`
	KmsKeyId          *string `json:"KmsKeyId,omitempty" xml:"KmsKeyId,omitempty"`
	// The time when the subscription attributes were last modified. The value is a UNIX timestamp representing the number of seconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1449554962
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
	// The maximum duration for which a message is retained in the topic. After the time specified by this parameter elapses since the message is sent to the topic, the message is deleted regardless of whether it has been successfully pushed to the user. Unit: seconds.
	//
	// example:
	//
	// 86400
	MessageRetentionPeriod *int64  `json:"MessageRetentionPeriod,omitempty" xml:"MessageRetentionPeriod,omitempty"`
	SseAlgorithm           *string `json:"SseAlgorithm,omitempty" xml:"SseAlgorithm,omitempty"`
	SseType                *string `json:"SseType,omitempty" xml:"SseType,omitempty"`
	// The list of resource tags.
	Tags []*ListTopicResponseBodyDataPageDataTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The internal URL of the topic.
	//
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
	// The URL of the topic.
	//
	// example:
	//
	// http:// 111111111****.mns.us-west-1.aliyuncs.com/topics/testTopic
	TopicUrl *string `json:"TopicUrl,omitempty" xml:"TopicUrl,omitempty"`
}

func (s ListTopicResponseBodyDataPageData) String() string {
	return dara.Prettify(s)
}

func (s ListTopicResponseBodyDataPageData) GoString() string {
	return s.String()
}

func (s *ListTopicResponseBodyDataPageData) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListTopicResponseBodyDataPageData) GetEnableSSE() *bool {
	return s.EnableSSE
}

func (s *ListTopicResponseBodyDataPageData) GetEncryptionEnabled() *bool {
	return s.EncryptionEnabled
}

func (s *ListTopicResponseBodyDataPageData) GetKmsKeyId() *string {
	return s.KmsKeyId
}

func (s *ListTopicResponseBodyDataPageData) GetLastModifyTime() *int64 {
	return s.LastModifyTime
}

func (s *ListTopicResponseBodyDataPageData) GetLoggingEnabled() *bool {
	return s.LoggingEnabled
}

func (s *ListTopicResponseBodyDataPageData) GetMaxMessageSize() *int64 {
	return s.MaxMessageSize
}

func (s *ListTopicResponseBodyDataPageData) GetMessageCount() *int64 {
	return s.MessageCount
}

func (s *ListTopicResponseBodyDataPageData) GetMessageRetentionPeriod() *int64 {
	return s.MessageRetentionPeriod
}

func (s *ListTopicResponseBodyDataPageData) GetSseAlgorithm() *string {
	return s.SseAlgorithm
}

func (s *ListTopicResponseBodyDataPageData) GetSseType() *string {
	return s.SseType
}

func (s *ListTopicResponseBodyDataPageData) GetTags() []*ListTopicResponseBodyDataPageDataTags {
	return s.Tags
}

func (s *ListTopicResponseBodyDataPageData) GetTopicInnerUrl() *string {
	return s.TopicInnerUrl
}

func (s *ListTopicResponseBodyDataPageData) GetTopicName() *string {
	return s.TopicName
}

func (s *ListTopicResponseBodyDataPageData) GetTopicType() *string {
	return s.TopicType
}

func (s *ListTopicResponseBodyDataPageData) GetTopicUrl() *string {
	return s.TopicUrl
}

func (s *ListTopicResponseBodyDataPageData) SetCreateTime(v int64) *ListTopicResponseBodyDataPageData {
	s.CreateTime = &v
	return s
}

func (s *ListTopicResponseBodyDataPageData) SetEnableSSE(v bool) *ListTopicResponseBodyDataPageData {
	s.EnableSSE = &v
	return s
}

func (s *ListTopicResponseBodyDataPageData) SetEncryptionEnabled(v bool) *ListTopicResponseBodyDataPageData {
	s.EncryptionEnabled = &v
	return s
}

func (s *ListTopicResponseBodyDataPageData) SetKmsKeyId(v string) *ListTopicResponseBodyDataPageData {
	s.KmsKeyId = &v
	return s
}

func (s *ListTopicResponseBodyDataPageData) SetLastModifyTime(v int64) *ListTopicResponseBodyDataPageData {
	s.LastModifyTime = &v
	return s
}

func (s *ListTopicResponseBodyDataPageData) SetLoggingEnabled(v bool) *ListTopicResponseBodyDataPageData {
	s.LoggingEnabled = &v
	return s
}

func (s *ListTopicResponseBodyDataPageData) SetMaxMessageSize(v int64) *ListTopicResponseBodyDataPageData {
	s.MaxMessageSize = &v
	return s
}

func (s *ListTopicResponseBodyDataPageData) SetMessageCount(v int64) *ListTopicResponseBodyDataPageData {
	s.MessageCount = &v
	return s
}

func (s *ListTopicResponseBodyDataPageData) SetMessageRetentionPeriod(v int64) *ListTopicResponseBodyDataPageData {
	s.MessageRetentionPeriod = &v
	return s
}

func (s *ListTopicResponseBodyDataPageData) SetSseAlgorithm(v string) *ListTopicResponseBodyDataPageData {
	s.SseAlgorithm = &v
	return s
}

func (s *ListTopicResponseBodyDataPageData) SetSseType(v string) *ListTopicResponseBodyDataPageData {
	s.SseType = &v
	return s
}

func (s *ListTopicResponseBodyDataPageData) SetTags(v []*ListTopicResponseBodyDataPageDataTags) *ListTopicResponseBodyDataPageData {
	s.Tags = v
	return s
}

func (s *ListTopicResponseBodyDataPageData) SetTopicInnerUrl(v string) *ListTopicResponseBodyDataPageData {
	s.TopicInnerUrl = &v
	return s
}

func (s *ListTopicResponseBodyDataPageData) SetTopicName(v string) *ListTopicResponseBodyDataPageData {
	s.TopicName = &v
	return s
}

func (s *ListTopicResponseBodyDataPageData) SetTopicType(v string) *ListTopicResponseBodyDataPageData {
	s.TopicType = &v
	return s
}

func (s *ListTopicResponseBodyDataPageData) SetTopicUrl(v string) *ListTopicResponseBodyDataPageData {
	s.TopicUrl = &v
	return s
}

func (s *ListTopicResponseBodyDataPageData) Validate() error {
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

type ListTopicResponseBodyDataPageDataTags struct {
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

func (s ListTopicResponseBodyDataPageDataTags) String() string {
	return dara.Prettify(s)
}

func (s ListTopicResponseBodyDataPageDataTags) GoString() string {
	return s.String()
}

func (s *ListTopicResponseBodyDataPageDataTags) GetTagKey() *string {
	return s.TagKey
}

func (s *ListTopicResponseBodyDataPageDataTags) GetTagValue() *string {
	return s.TagValue
}

func (s *ListTopicResponseBodyDataPageDataTags) SetTagKey(v string) *ListTopicResponseBodyDataPageDataTags {
	s.TagKey = &v
	return s
}

func (s *ListTopicResponseBodyDataPageDataTags) SetTagValue(v string) *ListTopicResponseBodyDataPageDataTags {
	s.TagValue = &v
	return s
}

func (s *ListTopicResponseBodyDataPageDataTags) Validate() error {
	return dara.Validate(s)
}
