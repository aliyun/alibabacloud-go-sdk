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
	// The data returned.
	Data *ListTopicResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
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
	return dara.Validate(s)
}

type ListTopicResponseBodyData struct {
	// The data returned on the current page.
	PageData []*ListTopicResponseBodyDataPageData `json:"PageData,omitempty" xml:"PageData,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNum *int64 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 50
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
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
	return dara.Validate(s)
}

type ListTopicResponseBodyDataPageData struct {
	// The time when the subscription was created. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1449554962
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the subscription was last modified. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1449554962
	LastModifyTime *int64 `json:"LastModifyTime,omitempty" xml:"LastModifyTime,omitempty"`
	// Indicates whether the logging feature is enabled.
	//
	// 	- True
	//
	// 	- False
	//
	// example:
	//
	// True
	LoggingEnabled *bool `json:"LoggingEnabled,omitempty" xml:"LoggingEnabled,omitempty"`
	// The maximum length of the message that is sent to the topic. Unit: bytes.
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
	// The maximum duration for which a message is retained in the topic. After the specified retention period ends, the message is deleted regardless of whether the message is received. Unit: seconds.
	//
	// example:
	//
	// 86400
	MessageRetentionPeriod *int64 `json:"MessageRetentionPeriod,omitempty" xml:"MessageRetentionPeriod,omitempty"`
	// The tags added to the resources.
	Tags []*ListTopicResponseBodyDataPageDataTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The internal URL of the message topic. The internal URL can be accessed over an internal network.
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
	// The URL of the message topic.
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

func (s *ListTopicResponseBodyDataPageData) GetTags() []*ListTopicResponseBodyDataPageDataTags {
	return s.Tags
}

func (s *ListTopicResponseBodyDataPageData) GetTopicInnerUrl() *string {
	return s.TopicInnerUrl
}

func (s *ListTopicResponseBodyDataPageData) GetTopicName() *string {
	return s.TopicName
}

func (s *ListTopicResponseBodyDataPageData) GetTopicUrl() *string {
	return s.TopicUrl
}

func (s *ListTopicResponseBodyDataPageData) SetCreateTime(v int64) *ListTopicResponseBodyDataPageData {
	s.CreateTime = &v
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

func (s *ListTopicResponseBodyDataPageData) SetTopicUrl(v string) *ListTopicResponseBodyDataPageData {
	s.TopicUrl = &v
	return s
}

func (s *ListTopicResponseBodyDataPageData) Validate() error {
	return dara.Validate(s)
}

type ListTopicResponseBodyDataPageDataTags struct {
	// The tag key.
	//
	// example:
	//
	// tag1
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value.
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
