// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CreateQueueRequest struct {
	// The period after which all messages sent to the queue are consumed. Valid values: 0 to 604800. Unit: seconds. Default value: 0
	//
	// example:
	//
	// 0
	DelaySeconds *int64 `json:"DelaySeconds,omitempty" xml:"DelaySeconds,omitempty"`
	// Specifies whether to enable the logging feature. Valid values:
	//
	// 	- True
	//
	// 	- False (default)
	//
	// example:
	//
	// true
	EnableLogging *bool `json:"EnableLogging,omitempty" xml:"EnableLogging,omitempty"`
	// The maximum length of the message that is sent to the queue. Valid values: 1024 to 65536. Unit: bytes. Default value: 65536.
	//
	// example:
	//
	// 65536
	MaximumMessageSize *int64 `json:"MaximumMessageSize,omitempty" xml:"MaximumMessageSize,omitempty"`
	// The maximum duration for which a message is retained in the queue. After the specified retention period ends, the message is deleted regardless of whether the message is received. Valid values: 60 to 604800. Unit: seconds. Default value: 345600.
	//
	// example:
	//
	// 345600
	MessageRetentionPeriod *int64 `json:"MessageRetentionPeriod,omitempty" xml:"MessageRetentionPeriod,omitempty"`
	// The maximum duration for which long polling requests are held after the ReceiveMessage operation is called. Valid values: 0 to 30. Unit: seconds. Default value: 0
	//
	// example:
	//
	// 0
	PollingWaitSeconds *int64 `json:"PollingWaitSeconds,omitempty" xml:"PollingWaitSeconds,omitempty"`
	// The name of the queue.
	//
	// This parameter is required.
	//
	// example:
	//
	// 06273500-249F-5863-121D-74D51123****
	QueueName *string                  `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	Tag       []*CreateQueueRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The duration for which a message stays in the Inactive state after the message is received from the queue. Valid values: 1 to 43200. Unit: seconds. Default value: 30.
	//
	// example:
	//
	// 60
	VisibilityTimeout *int64 `json:"VisibilityTimeout,omitempty" xml:"VisibilityTimeout,omitempty"`
}

func (s CreateQueueRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateQueueRequest) GoString() string {
	return s.String()
}

func (s *CreateQueueRequest) SetDelaySeconds(v int64) *CreateQueueRequest {
	s.DelaySeconds = &v
	return s
}

func (s *CreateQueueRequest) SetEnableLogging(v bool) *CreateQueueRequest {
	s.EnableLogging = &v
	return s
}

func (s *CreateQueueRequest) SetMaximumMessageSize(v int64) *CreateQueueRequest {
	s.MaximumMessageSize = &v
	return s
}

func (s *CreateQueueRequest) SetMessageRetentionPeriod(v int64) *CreateQueueRequest {
	s.MessageRetentionPeriod = &v
	return s
}

func (s *CreateQueueRequest) SetPollingWaitSeconds(v int64) *CreateQueueRequest {
	s.PollingWaitSeconds = &v
	return s
}

func (s *CreateQueueRequest) SetQueueName(v string) *CreateQueueRequest {
	s.QueueName = &v
	return s
}

func (s *CreateQueueRequest) SetTag(v []*CreateQueueRequestTag) *CreateQueueRequest {
	s.Tag = v
	return s
}

func (s *CreateQueueRequest) SetVisibilityTimeout(v int64) *CreateQueueRequest {
	s.VisibilityTimeout = &v
	return s
}

type CreateQueueRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateQueueRequestTag) String() string {
	return tea.Prettify(s)
}

func (s CreateQueueRequestTag) GoString() string {
	return s.String()
}

func (s *CreateQueueRequestTag) SetKey(v string) *CreateQueueRequestTag {
	s.Key = &v
	return s
}

func (s *CreateQueueRequestTag) SetValue(v string) *CreateQueueRequestTag {
	s.Value = &v
	return s
}

type CreateQueueResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *CreateQueueResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// 06273500-249F-5863-121D-74D51123E62C
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

func (s CreateQueueResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateQueueResponseBody) GoString() string {
	return s.String()
}

func (s *CreateQueueResponseBody) SetCode(v int64) *CreateQueueResponseBody {
	s.Code = &v
	return s
}

func (s *CreateQueueResponseBody) SetData(v *CreateQueueResponseBodyData) *CreateQueueResponseBody {
	s.Data = v
	return s
}

func (s *CreateQueueResponseBody) SetMessage(v string) *CreateQueueResponseBody {
	s.Message = &v
	return s
}

func (s *CreateQueueResponseBody) SetRequestId(v string) *CreateQueueResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateQueueResponseBody) SetStatus(v string) *CreateQueueResponseBody {
	s.Status = &v
	return s
}

func (s *CreateQueueResponseBody) SetSuccess(v bool) *CreateQueueResponseBody {
	s.Success = &v
	return s
}

type CreateQueueResponseBodyData struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateQueueResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateQueueResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateQueueResponseBodyData) SetCode(v int64) *CreateQueueResponseBodyData {
	s.Code = &v
	return s
}

func (s *CreateQueueResponseBodyData) SetMessage(v string) *CreateQueueResponseBodyData {
	s.Message = &v
	return s
}

func (s *CreateQueueResponseBodyData) SetSuccess(v bool) *CreateQueueResponseBodyData {
	s.Success = &v
	return s
}

type CreateQueueResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateQueueResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateQueueResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateQueueResponse) GoString() string {
	return s.String()
}

func (s *CreateQueueResponse) SetHeaders(v map[string]*string) *CreateQueueResponse {
	s.Headers = v
	return s
}

func (s *CreateQueueResponse) SetStatusCode(v int32) *CreateQueueResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateQueueResponse) SetBody(v *CreateQueueResponseBody) *CreateQueueResponse {
	s.Body = v
	return s
}

type CreateTopicRequest struct {
	// Specifies whether to enable the logging feature. Valid values:
	//
	// 	- True
	//
	// 	- False (default)
	//
	// example:
	//
	// true
	EnableLogging *bool `json:"EnableLogging,omitempty" xml:"EnableLogging,omitempty"`
	// The maximum length of the message that is sent to the topic. Valid values: 1024 to 65536. Unit: bytes. Default value: 65536.
	//
	// example:
	//
	// 10240
	MaxMessageSize *int64 `json:"MaxMessageSize,omitempty" xml:"MaxMessageSize,omitempty"`
	// The tags.
	Tag []*CreateTopicRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The name of the topic that you want to create.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	TopicName *string `json:"TopicName,omitempty" xml:"TopicName,omitempty"`
}

func (s CreateTopicRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTopicRequest) GoString() string {
	return s.String()
}

func (s *CreateTopicRequest) SetEnableLogging(v bool) *CreateTopicRequest {
	s.EnableLogging = &v
	return s
}

func (s *CreateTopicRequest) SetMaxMessageSize(v int64) *CreateTopicRequest {
	s.MaxMessageSize = &v
	return s
}

func (s *CreateTopicRequest) SetTag(v []*CreateTopicRequestTag) *CreateTopicRequest {
	s.Tag = v
	return s
}

func (s *CreateTopicRequest) SetTopicName(v string) *CreateTopicRequest {
	s.TopicName = &v
	return s
}

type CreateTopicRequestTag struct {
	// The tag key.
	//
	// example:
	//
	// tag1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// joyce.wang
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateTopicRequestTag) String() string {
	return tea.Prettify(s)
}

func (s CreateTopicRequestTag) GoString() string {
	return s.String()
}

func (s *CreateTopicRequestTag) SetKey(v string) *CreateTopicRequestTag {
	s.Key = &v
	return s
}

func (s *CreateTopicRequestTag) SetValue(v string) *CreateTopicRequestTag {
	s.Value = &v
	return s
}

type CreateTopicResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *CreateTopicResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// 06273500-249F-5863-121D-74D51123E62C
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

func (s CreateTopicResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTopicResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTopicResponseBody) SetCode(v int64) *CreateTopicResponseBody {
	s.Code = &v
	return s
}

func (s *CreateTopicResponseBody) SetData(v *CreateTopicResponseBodyData) *CreateTopicResponseBody {
	s.Data = v
	return s
}

func (s *CreateTopicResponseBody) SetMessage(v string) *CreateTopicResponseBody {
	s.Message = &v
	return s
}

func (s *CreateTopicResponseBody) SetRequestId(v string) *CreateTopicResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTopicResponseBody) SetStatus(v string) *CreateTopicResponseBody {
	s.Status = &v
	return s
}

func (s *CreateTopicResponseBody) SetSuccess(v bool) *CreateTopicResponseBody {
	s.Success = &v
	return s
}

type CreateTopicResponseBodyData struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateTopicResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateTopicResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateTopicResponseBodyData) SetCode(v int64) *CreateTopicResponseBodyData {
	s.Code = &v
	return s
}

func (s *CreateTopicResponseBodyData) SetMessage(v string) *CreateTopicResponseBodyData {
	s.Message = &v
	return s
}

func (s *CreateTopicResponseBodyData) SetSuccess(v bool) *CreateTopicResponseBodyData {
	s.Success = &v
	return s
}

type CreateTopicResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTopicResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTopicResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTopicResponse) GoString() string {
	return s.String()
}

func (s *CreateTopicResponse) SetHeaders(v map[string]*string) *CreateTopicResponse {
	s.Headers = v
	return s
}

func (s *CreateTopicResponse) SetStatusCode(v int32) *CreateTopicResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTopicResponse) SetBody(v *CreateTopicResponseBody) *CreateTopicResponse {
	s.Body = v
	return s
}

type DeleteQueueRequest struct {
	// The name of the queue.
	//
	// This parameter is required.
	//
	// example:
	//
	// tf-testAccMNSQueue-525478433321945943
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
}

func (s DeleteQueueRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteQueueRequest) GoString() string {
	return s.String()
}

func (s *DeleteQueueRequest) SetQueueName(v string) *DeleteQueueRequest {
	s.QueueName = &v
	return s
}

type DeleteQueueResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *DeleteQueueResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s DeleteQueueResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteQueueResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteQueueResponseBody) SetCode(v int64) *DeleteQueueResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteQueueResponseBody) SetData(v *DeleteQueueResponseBodyData) *DeleteQueueResponseBody {
	s.Data = v
	return s
}

func (s *DeleteQueueResponseBody) SetMessage(v string) *DeleteQueueResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteQueueResponseBody) SetRequestId(v string) *DeleteQueueResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteQueueResponseBody) SetStatus(v string) *DeleteQueueResponseBody {
	s.Status = &v
	return s
}

func (s *DeleteQueueResponseBody) SetSuccess(v bool) *DeleteQueueResponseBody {
	s.Success = &v
	return s
}

type DeleteQueueResponseBodyData struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteQueueResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DeleteQueueResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteQueueResponseBodyData) SetCode(v int64) *DeleteQueueResponseBodyData {
	s.Code = &v
	return s
}

func (s *DeleteQueueResponseBodyData) SetMessage(v string) *DeleteQueueResponseBodyData {
	s.Message = &v
	return s
}

func (s *DeleteQueueResponseBodyData) SetSuccess(v bool) *DeleteQueueResponseBodyData {
	s.Success = &v
	return s
}

type DeleteQueueResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteQueueResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteQueueResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteQueueResponse) GoString() string {
	return s.String()
}

func (s *DeleteQueueResponse) SetHeaders(v map[string]*string) *DeleteQueueResponse {
	s.Headers = v
	return s
}

func (s *DeleteQueueResponse) SetStatusCode(v int32) *DeleteQueueResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteQueueResponse) SetBody(v *DeleteQueueResponseBody) *DeleteQueueResponse {
	s.Body = v
	return s
}

type DeleteTopicRequest struct {
	// The name of the topic that you want to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// tf-testAccMNSTopic-112965059402264645
	TopicName *string `json:"TopicName,omitempty" xml:"TopicName,omitempty"`
}

func (s DeleteTopicRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteTopicRequest) GoString() string {
	return s.String()
}

func (s *DeleteTopicRequest) SetTopicName(v string) *DeleteTopicRequest {
	s.TopicName = &v
	return s
}

type DeleteTopicResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
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

func (s DeleteTopicResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteTopicResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTopicResponseBody) SetCode(v int64) *DeleteTopicResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteTopicResponseBody) SetData(v map[string]interface{}) *DeleteTopicResponseBody {
	s.Data = v
	return s
}

func (s *DeleteTopicResponseBody) SetMessage(v string) *DeleteTopicResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteTopicResponseBody) SetRequestId(v string) *DeleteTopicResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteTopicResponseBody) SetStatus(v string) *DeleteTopicResponseBody {
	s.Status = &v
	return s
}

func (s *DeleteTopicResponseBody) SetSuccess(v bool) *DeleteTopicResponseBody {
	s.Success = &v
	return s
}

type DeleteTopicResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteTopicResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteTopicResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteTopicResponse) GoString() string {
	return s.String()
}

func (s *DeleteTopicResponse) SetHeaders(v map[string]*string) *DeleteTopicResponse {
	s.Headers = v
	return s
}

func (s *DeleteTopicResponse) SetStatusCode(v int32) *DeleteTopicResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTopicResponse) SetBody(v *DeleteTopicResponseBody) *DeleteTopicResponse {
	s.Body = v
	return s
}

type GetQueueAttributesRequest struct {
	// The name of the queue.
	//
	// This parameter is required.
	//
	// example:
	//
	// demo-queue
	QueueName *string                         `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	Tag       []*GetQueueAttributesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s GetQueueAttributesRequest) String() string {
	return tea.Prettify(s)
}

func (s GetQueueAttributesRequest) GoString() string {
	return s.String()
}

func (s *GetQueueAttributesRequest) SetQueueName(v string) *GetQueueAttributesRequest {
	s.QueueName = &v
	return s
}

func (s *GetQueueAttributesRequest) SetTag(v []*GetQueueAttributesRequestTag) *GetQueueAttributesRequest {
	s.Tag = v
	return s
}

type GetQueueAttributesRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetQueueAttributesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s GetQueueAttributesRequestTag) GoString() string {
	return s.String()
}

func (s *GetQueueAttributesRequestTag) SetKey(v string) *GetQueueAttributesRequestTag {
	s.Key = &v
	return s
}

func (s *GetQueueAttributesRequestTag) SetValue(v string) *GetQueueAttributesRequestTag {
	s.Value = &v
	return s
}

type GetQueueAttributesResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *GetQueueAttributesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s GetQueueAttributesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetQueueAttributesResponseBody) GoString() string {
	return s.String()
}

func (s *GetQueueAttributesResponseBody) SetCode(v int64) *GetQueueAttributesResponseBody {
	s.Code = &v
	return s
}

func (s *GetQueueAttributesResponseBody) SetData(v *GetQueueAttributesResponseBodyData) *GetQueueAttributesResponseBody {
	s.Data = v
	return s
}

func (s *GetQueueAttributesResponseBody) SetMessage(v string) *GetQueueAttributesResponseBody {
	s.Message = &v
	return s
}

func (s *GetQueueAttributesResponseBody) SetRequestId(v string) *GetQueueAttributesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetQueueAttributesResponseBody) SetStatus(v string) *GetQueueAttributesResponseBody {
	s.Status = &v
	return s
}

func (s *GetQueueAttributesResponseBody) SetSuccess(v bool) *GetQueueAttributesResponseBody {
	s.Success = &v
	return s
}

type GetQueueAttributesResponseBodyData struct {
	// The total number of messages that are in the Active state in the queue. The value is an approximate number.
	//
	// example:
	//
	// 20
	ActiveMessages *int64 `json:"ActiveMessages,omitempty" xml:"ActiveMessages,omitempty"`
	// The time when the queue was created.
	//
	// example:
	//
	// 1250700999
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The total number of the messages that are in the Delayed state in the queue. The value is an approximate number.
	//
	// example:
	//
	// 0
	DelayMessages *int64 `json:"DelayMessages,omitempty" xml:"DelayMessages,omitempty"`
	// The period after which all messages sent to the queue are consumed. Unit: seconds.
	//
	// example:
	//
	// 30
	DelaySeconds *int64 `json:"DelaySeconds,omitempty" xml:"DelaySeconds,omitempty"`
	// The total number of the messages that are in the Inactive state in the queue. The value is an approximate number.
	//
	// example:
	//
	// 0
	InactiveMessages *int64 `json:"InactiveMessages,omitempty" xml:"InactiveMessages,omitempty"`
	// The time when the queue was last modified. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1250700999
	LastModifyTime *int64 `json:"LastModifyTime,omitempty" xml:"LastModifyTime,omitempty"`
	// Indicates whether the logging feature is enabled. Valid values:
	//
	// 	- True
	//
	// 	- False
	//
	// example:
	//
	// True
	LoggingEnabled *bool `json:"LoggingEnabled,omitempty" xml:"LoggingEnabled,omitempty"`
	// The maximum length of the message that is sent to the queue. Unit: bytes.
	//
	// example:
	//
	// 65536
	MaximumMessageSize *int64 `json:"MaximumMessageSize,omitempty" xml:"MaximumMessageSize,omitempty"`
	// The maximum duration for which a message is retained in the queue. After the specified retention period ends, the message is deleted regardless of whether the message is received. Unit: seconds.
	//
	// example:
	//
	// 65536
	MessageRetentionPeriod *int64 `json:"MessageRetentionPeriod,omitempty" xml:"MessageRetentionPeriod,omitempty"`
	// The maximum duration for which long polling requests are held after the ReceiveMessage operation is called. Unit: seconds.
	//
	// example:
	//
	// 0
	PollingWaitSeconds *int64 `json:"PollingWaitSeconds,omitempty" xml:"PollingWaitSeconds,omitempty"`
	// The name of the queue.
	//
	// example:
	//
	// demo-queue
	QueueName *string                                   `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	Tags      []*GetQueueAttributesResponseBodyDataTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The duration for which a message stays in the Inactive state after the message is received from the queue. Valid values: 1 to 43200. Unit: seconds. Default value: 30.
	//
	// example:
	//
	// 60
	VisibilityTimeout *int64 `json:"VisibilityTimeout,omitempty" xml:"VisibilityTimeout,omitempty"`
}

func (s GetQueueAttributesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetQueueAttributesResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetQueueAttributesResponseBodyData) SetActiveMessages(v int64) *GetQueueAttributesResponseBodyData {
	s.ActiveMessages = &v
	return s
}

func (s *GetQueueAttributesResponseBodyData) SetCreateTime(v int64) *GetQueueAttributesResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetQueueAttributesResponseBodyData) SetDelayMessages(v int64) *GetQueueAttributesResponseBodyData {
	s.DelayMessages = &v
	return s
}

func (s *GetQueueAttributesResponseBodyData) SetDelaySeconds(v int64) *GetQueueAttributesResponseBodyData {
	s.DelaySeconds = &v
	return s
}

func (s *GetQueueAttributesResponseBodyData) SetInactiveMessages(v int64) *GetQueueAttributesResponseBodyData {
	s.InactiveMessages = &v
	return s
}

func (s *GetQueueAttributesResponseBodyData) SetLastModifyTime(v int64) *GetQueueAttributesResponseBodyData {
	s.LastModifyTime = &v
	return s
}

func (s *GetQueueAttributesResponseBodyData) SetLoggingEnabled(v bool) *GetQueueAttributesResponseBodyData {
	s.LoggingEnabled = &v
	return s
}

func (s *GetQueueAttributesResponseBodyData) SetMaximumMessageSize(v int64) *GetQueueAttributesResponseBodyData {
	s.MaximumMessageSize = &v
	return s
}

func (s *GetQueueAttributesResponseBodyData) SetMessageRetentionPeriod(v int64) *GetQueueAttributesResponseBodyData {
	s.MessageRetentionPeriod = &v
	return s
}

func (s *GetQueueAttributesResponseBodyData) SetPollingWaitSeconds(v int64) *GetQueueAttributesResponseBodyData {
	s.PollingWaitSeconds = &v
	return s
}

func (s *GetQueueAttributesResponseBodyData) SetQueueName(v string) *GetQueueAttributesResponseBodyData {
	s.QueueName = &v
	return s
}

func (s *GetQueueAttributesResponseBodyData) SetTags(v []*GetQueueAttributesResponseBodyDataTags) *GetQueueAttributesResponseBodyData {
	s.Tags = v
	return s
}

func (s *GetQueueAttributesResponseBodyData) SetVisibilityTimeout(v int64) *GetQueueAttributesResponseBodyData {
	s.VisibilityTimeout = &v
	return s
}

type GetQueueAttributesResponseBodyDataTags struct {
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s GetQueueAttributesResponseBodyDataTags) String() string {
	return tea.Prettify(s)
}

func (s GetQueueAttributesResponseBodyDataTags) GoString() string {
	return s.String()
}

func (s *GetQueueAttributesResponseBodyDataTags) SetTagKey(v string) *GetQueueAttributesResponseBodyDataTags {
	s.TagKey = &v
	return s
}

func (s *GetQueueAttributesResponseBodyDataTags) SetTagValue(v string) *GetQueueAttributesResponseBodyDataTags {
	s.TagValue = &v
	return s
}

type GetQueueAttributesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetQueueAttributesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetQueueAttributesResponse) String() string {
	return tea.Prettify(s)
}

func (s GetQueueAttributesResponse) GoString() string {
	return s.String()
}

func (s *GetQueueAttributesResponse) SetHeaders(v map[string]*string) *GetQueueAttributesResponse {
	s.Headers = v
	return s
}

func (s *GetQueueAttributesResponse) SetStatusCode(v int32) *GetQueueAttributesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetQueueAttributesResponse) SetBody(v *GetQueueAttributesResponseBody) *GetQueueAttributesResponse {
	s.Body = v
	return s
}

type GetSubscriptionAttributesRequest struct {
	// The name of the subscription.
	//
	// This parameter is required.
	//
	// example:
	//
	// MySubscription
	SubscriptionName *string `json:"SubscriptionName,omitempty" xml:"SubscriptionName,omitempty"`
	// The name of the topic.
	//
	// This parameter is required.
	//
	// example:
	//
	// MyTopic
	TopicName *string `json:"TopicName,omitempty" xml:"TopicName,omitempty"`
}

func (s GetSubscriptionAttributesRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSubscriptionAttributesRequest) GoString() string {
	return s.String()
}

func (s *GetSubscriptionAttributesRequest) SetSubscriptionName(v string) *GetSubscriptionAttributesRequest {
	s.SubscriptionName = &v
	return s
}

func (s *GetSubscriptionAttributesRequest) SetTopicName(v string) *GetSubscriptionAttributesRequest {
	s.TopicName = &v
	return s
}

type GetSubscriptionAttributesResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *GetSubscriptionAttributesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s GetSubscriptionAttributesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSubscriptionAttributesResponseBody) GoString() string {
	return s.String()
}

func (s *GetSubscriptionAttributesResponseBody) SetCode(v int64) *GetSubscriptionAttributesResponseBody {
	s.Code = &v
	return s
}

func (s *GetSubscriptionAttributesResponseBody) SetData(v *GetSubscriptionAttributesResponseBodyData) *GetSubscriptionAttributesResponseBody {
	s.Data = v
	return s
}

func (s *GetSubscriptionAttributesResponseBody) SetMessage(v string) *GetSubscriptionAttributesResponseBody {
	s.Message = &v
	return s
}

func (s *GetSubscriptionAttributesResponseBody) SetRequestId(v string) *GetSubscriptionAttributesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSubscriptionAttributesResponseBody) SetStatus(v string) *GetSubscriptionAttributesResponseBody {
	s.Status = &v
	return s
}

func (s *GetSubscriptionAttributesResponseBody) SetSuccess(v bool) *GetSubscriptionAttributesResponseBody {
	s.Success = &v
	return s
}

type GetSubscriptionAttributesResponseBodyData struct {
	// The time when the subscription was created. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1449554806
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The endpoint to which the messages are pushed.
	//
	// example:
	//
	// http://example.com
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// The tag that is used to filter messages. Only the messages that are attached with the specified tag can be pushed.
	//
	// example:
	//
	// important
	FilterTag *string `json:"FilterTag,omitempty" xml:"FilterTag,omitempty"`
	// The time when the subscription was last modified. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1449554962
	LastModifyTime *int64 `json:"LastModifyTime,omitempty" xml:"LastModifyTime,omitempty"`
	// The content format of the messages that are pushed to the endpoint. Valid values:
	//
	// 	- XML
	//
	// 	- JSON
	//
	// 	- SIMPLIFIED
	//
	// example:
	//
	// XML
	NotifyContentFormat *string `json:"NotifyContentFormat,omitempty" xml:"NotifyContentFormat,omitempty"`
	// The retry policy that is applied if an error occurs when Message Service (MNS) pushes messages to the endpoint. Valid values:
	//
	// 	- BACKOFF_RETRY
	//
	// 	- EXPONENTIAL_DECAY_RETRY
	//
	// example:
	//
	// BACKOFF_RETRY
	NotifyStrategy *string `json:"NotifyStrategy,omitempty" xml:"NotifyStrategy,omitempty"`
	// The name of the subscription.
	//
	// example:
	//
	// MySubscription
	SubscriptionName *string `json:"SubscriptionName,omitempty" xml:"SubscriptionName,omitempty"`
	// The name of the topic.
	//
	// example:
	//
	// MyTopic
	TopicName *string `json:"TopicName,omitempty" xml:"TopicName,omitempty"`
	// The Alibaba Cloud account ID of the topic owner.
	//
	// example:
	//
	// 123456789098****
	TopicOwner *string `json:"TopicOwner,omitempty" xml:"TopicOwner,omitempty"`
}

func (s GetSubscriptionAttributesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetSubscriptionAttributesResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSubscriptionAttributesResponseBodyData) SetCreateTime(v int64) *GetSubscriptionAttributesResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetSubscriptionAttributesResponseBodyData) SetEndpoint(v string) *GetSubscriptionAttributesResponseBodyData {
	s.Endpoint = &v
	return s
}

func (s *GetSubscriptionAttributesResponseBodyData) SetFilterTag(v string) *GetSubscriptionAttributesResponseBodyData {
	s.FilterTag = &v
	return s
}

func (s *GetSubscriptionAttributesResponseBodyData) SetLastModifyTime(v int64) *GetSubscriptionAttributesResponseBodyData {
	s.LastModifyTime = &v
	return s
}

func (s *GetSubscriptionAttributesResponseBodyData) SetNotifyContentFormat(v string) *GetSubscriptionAttributesResponseBodyData {
	s.NotifyContentFormat = &v
	return s
}

func (s *GetSubscriptionAttributesResponseBodyData) SetNotifyStrategy(v string) *GetSubscriptionAttributesResponseBodyData {
	s.NotifyStrategy = &v
	return s
}

func (s *GetSubscriptionAttributesResponseBodyData) SetSubscriptionName(v string) *GetSubscriptionAttributesResponseBodyData {
	s.SubscriptionName = &v
	return s
}

func (s *GetSubscriptionAttributesResponseBodyData) SetTopicName(v string) *GetSubscriptionAttributesResponseBodyData {
	s.TopicName = &v
	return s
}

func (s *GetSubscriptionAttributesResponseBodyData) SetTopicOwner(v string) *GetSubscriptionAttributesResponseBodyData {
	s.TopicOwner = &v
	return s
}

type GetSubscriptionAttributesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSubscriptionAttributesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSubscriptionAttributesResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSubscriptionAttributesResponse) GoString() string {
	return s.String()
}

func (s *GetSubscriptionAttributesResponse) SetHeaders(v map[string]*string) *GetSubscriptionAttributesResponse {
	s.Headers = v
	return s
}

func (s *GetSubscriptionAttributesResponse) SetStatusCode(v int32) *GetSubscriptionAttributesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSubscriptionAttributesResponse) SetBody(v *GetSubscriptionAttributesResponseBody) *GetSubscriptionAttributesResponse {
	s.Body = v
	return s
}

type GetTopicAttributesRequest struct {
	Tag []*GetTopicAttributesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The name of the topic.
	//
	// This parameter is required.
	//
	// example:
	//
	// demo-topic
	TopicName *string `json:"TopicName,omitempty" xml:"TopicName,omitempty"`
}

func (s GetTopicAttributesRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTopicAttributesRequest) GoString() string {
	return s.String()
}

func (s *GetTopicAttributesRequest) SetTag(v []*GetTopicAttributesRequestTag) *GetTopicAttributesRequest {
	s.Tag = v
	return s
}

func (s *GetTopicAttributesRequest) SetTopicName(v string) *GetTopicAttributesRequest {
	s.TopicName = &v
	return s
}

type GetTopicAttributesRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetTopicAttributesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s GetTopicAttributesRequestTag) GoString() string {
	return s.String()
}

func (s *GetTopicAttributesRequestTag) SetKey(v string) *GetTopicAttributesRequestTag {
	s.Key = &v
	return s
}

func (s *GetTopicAttributesRequestTag) SetValue(v string) *GetTopicAttributesRequestTag {
	s.Value = &v
	return s
}

type GetTopicAttributesResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *GetTopicAttributesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s GetTopicAttributesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTopicAttributesResponseBody) GoString() string {
	return s.String()
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

type GetTopicAttributesResponseBodyData struct {
	// The time when the topic was created. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1449554277
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the topic was last modified. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1449554460
	LastModifyTime *int64 `json:"LastModifyTime,omitempty" xml:"LastModifyTime,omitempty"`
	// Indicates whether the logging feature is enabled. Valid values:
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
	MessageRetentionPeriod *int64                                    `json:"MessageRetentionPeriod,omitempty" xml:"MessageRetentionPeriod,omitempty"`
	Tags                   []*GetTopicAttributesResponseBodyDataTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The name of the topic.
	//
	// example:
	//
	// demo-topic
	TopicName *string `json:"TopicName,omitempty" xml:"TopicName,omitempty"`
}

func (s GetTopicAttributesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetTopicAttributesResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTopicAttributesResponseBodyData) SetCreateTime(v int64) *GetTopicAttributesResponseBodyData {
	s.CreateTime = &v
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

func (s *GetTopicAttributesResponseBodyData) SetTags(v []*GetTopicAttributesResponseBodyDataTags) *GetTopicAttributesResponseBodyData {
	s.Tags = v
	return s
}

func (s *GetTopicAttributesResponseBodyData) SetTopicName(v string) *GetTopicAttributesResponseBodyData {
	s.TopicName = &v
	return s
}

type GetTopicAttributesResponseBodyDataTags struct {
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s GetTopicAttributesResponseBodyDataTags) String() string {
	return tea.Prettify(s)
}

func (s GetTopicAttributesResponseBodyDataTags) GoString() string {
	return s.String()
}

func (s *GetTopicAttributesResponseBodyDataTags) SetTagKey(v string) *GetTopicAttributesResponseBodyDataTags {
	s.TagKey = &v
	return s
}

func (s *GetTopicAttributesResponseBodyDataTags) SetTagValue(v string) *GetTopicAttributesResponseBodyDataTags {
	s.TagValue = &v
	return s
}

type GetTopicAttributesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTopicAttributesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTopicAttributesResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTopicAttributesResponse) GoString() string {
	return s.String()
}

func (s *GetTopicAttributesResponse) SetHeaders(v map[string]*string) *GetTopicAttributesResponse {
	s.Headers = v
	return s
}

func (s *GetTopicAttributesResponse) SetStatusCode(v int32) *GetTopicAttributesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTopicAttributesResponse) SetBody(v *GetTopicAttributesResponseBody) *GetTopicAttributesResponse {
	s.Body = v
	return s
}

type ListQueueRequest struct {
	// The page number. Valid values: 1 to 100000000. If you set this parameter to a value smaller than 1, the value of this parameter is 1 by default. If you set this parameter to a value greater than 100000000, the value of this parameter is 100000000 by default.
	//
	// example:
	//
	// 1
	PageNum *int64 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries per page. Value values: 10 to 50. If you set this parameter to a value smaller than 10, the value of this parameter is 10 by default. If you set this parameter to a value greater than 50, the value of this parameter is 50 by default.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The name of the queue.
	//
	// example:
	//
	// demo-queue
	QueueName *string                `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	Tag       []*ListQueueRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListQueueRequest) String() string {
	return tea.Prettify(s)
}

func (s ListQueueRequest) GoString() string {
	return s.String()
}

func (s *ListQueueRequest) SetPageNum(v int64) *ListQueueRequest {
	s.PageNum = &v
	return s
}

func (s *ListQueueRequest) SetPageSize(v int64) *ListQueueRequest {
	s.PageSize = &v
	return s
}

func (s *ListQueueRequest) SetQueueName(v string) *ListQueueRequest {
	s.QueueName = &v
	return s
}

func (s *ListQueueRequest) SetTag(v []*ListQueueRequestTag) *ListQueueRequest {
	s.Tag = v
	return s
}

type ListQueueRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListQueueRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListQueueRequestTag) GoString() string {
	return s.String()
}

func (s *ListQueueRequestTag) SetKey(v string) *ListQueueRequestTag {
	s.Key = &v
	return s
}

func (s *ListQueueRequestTag) SetValue(v string) *ListQueueRequestTag {
	s.Value = &v
	return s
}

type ListQueueResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *ListQueueResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s ListQueueResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListQueueResponseBody) GoString() string {
	return s.String()
}

func (s *ListQueueResponseBody) SetCode(v int64) *ListQueueResponseBody {
	s.Code = &v
	return s
}

func (s *ListQueueResponseBody) SetData(v *ListQueueResponseBodyData) *ListQueueResponseBody {
	s.Data = v
	return s
}

func (s *ListQueueResponseBody) SetMessage(v string) *ListQueueResponseBody {
	s.Message = &v
	return s
}

func (s *ListQueueResponseBody) SetRequestId(v string) *ListQueueResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListQueueResponseBody) SetStatus(v string) *ListQueueResponseBody {
	s.Status = &v
	return s
}

func (s *ListQueueResponseBody) SetSuccess(v bool) *ListQueueResponseBody {
	s.Success = &v
	return s
}

type ListQueueResponseBodyData struct {
	// The data returned on the current page.
	PageData []*ListQueueResponseBodyDataPageData `json:"PageData,omitempty" xml:"PageData,omitempty" type:"Repeated"`
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
	// The total number of pages returned.
	//
	// example:
	//
	// 3
	Pages *int64 `json:"Pages,omitempty" xml:"Pages,omitempty"`
	// The number of entries on the current page.
	//
	// example:
	//
	// 20
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 130
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListQueueResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListQueueResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListQueueResponseBodyData) SetPageData(v []*ListQueueResponseBodyDataPageData) *ListQueueResponseBodyData {
	s.PageData = v
	return s
}

func (s *ListQueueResponseBodyData) SetPageNum(v int64) *ListQueueResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *ListQueueResponseBodyData) SetPageSize(v int64) *ListQueueResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListQueueResponseBodyData) SetPages(v int64) *ListQueueResponseBodyData {
	s.Pages = &v
	return s
}

func (s *ListQueueResponseBodyData) SetSize(v int64) *ListQueueResponseBodyData {
	s.Size = &v
	return s
}

func (s *ListQueueResponseBodyData) SetTotal(v int64) *ListQueueResponseBodyData {
	s.Total = &v
	return s
}

type ListQueueResponseBodyDataPageData struct {
	// The total number of messages that are in the Active state in the queue. The value is an approximate number.
	//
	// example:
	//
	// 20
	ActiveMessages *int64 `json:"ActiveMessages,omitempty" xml:"ActiveMessages,omitempty"`
	// The time when the queue was created. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1250700999
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The total number of the messages that are in the Delayed state in the queue. The value is an approximate number.
	//
	// example:
	//
	// 0
	DelayMessages *int64 `json:"DelayMessages,omitempty" xml:"DelayMessages,omitempty"`
	// The period after which all messages sent to the queue are consumed. Unit: seconds.
	//
	// example:
	//
	// 30
	DelaySeconds *int64 `json:"DelaySeconds,omitempty" xml:"DelaySeconds,omitempty"`
	// The total number of the messages that are in the Inactive state in the queue. The value is an approximate number.
	//
	// example:
	//
	// 0
	InactiveMessages *int64 `json:"InactiveMessages,omitempty" xml:"InactiveMessages,omitempty"`
	// The time when the queue was last modified. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1250700999
	LastModifyTime *int64 `json:"LastModifyTime,omitempty" xml:"LastModifyTime,omitempty"`
	// Indicates whether the logging feature is enabled. Valid values:
	//
	// 	- True
	//
	// 	- False
	//
	// example:
	//
	// True
	LoggingEnabled *bool `json:"LoggingEnabled,omitempty" xml:"LoggingEnabled,omitempty"`
	// The maximum length of the message that is sent to the queue. Unit: bytes.
	//
	// example:
	//
	// 65536
	MaximumMessageSize *int64 `json:"MaximumMessageSize,omitempty" xml:"MaximumMessageSize,omitempty"`
	// The maximum duration for which a message is retained in the queue. After the specified retention period ends, the message is deleted regardless of whether the message is received. Unit: seconds.
	//
	// example:
	//
	// 65536
	MessageRetentionPeriod *int64 `json:"MessageRetentionPeriod,omitempty" xml:"MessageRetentionPeriod,omitempty"`
	// The maximum duration for which long polling requests are held after the ReceiveMessage operation is called. Unit: seconds.
	//
	// example:
	//
	// 0
	PollingWaitSeconds *int64 `json:"PollingWaitSeconds,omitempty" xml:"PollingWaitSeconds,omitempty"`
	// The name of the queue.
	//
	// example:
	//
	// demo-queue
	QueueName *string                                  `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	Tags      []*ListQueueResponseBodyDataPageDataTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The duration for which a message stays in the Inactive state after the message is received from the queue. Valid values: 1 to 43200. Unit: seconds. Default value: 30.
	//
	// example:
	//
	// 60
	VisibilityTimeout *int64 `json:"VisibilityTimeout,omitempty" xml:"VisibilityTimeout,omitempty"`
}

func (s ListQueueResponseBodyDataPageData) String() string {
	return tea.Prettify(s)
}

func (s ListQueueResponseBodyDataPageData) GoString() string {
	return s.String()
}

func (s *ListQueueResponseBodyDataPageData) SetActiveMessages(v int64) *ListQueueResponseBodyDataPageData {
	s.ActiveMessages = &v
	return s
}

func (s *ListQueueResponseBodyDataPageData) SetCreateTime(v int64) *ListQueueResponseBodyDataPageData {
	s.CreateTime = &v
	return s
}

func (s *ListQueueResponseBodyDataPageData) SetDelayMessages(v int64) *ListQueueResponseBodyDataPageData {
	s.DelayMessages = &v
	return s
}

func (s *ListQueueResponseBodyDataPageData) SetDelaySeconds(v int64) *ListQueueResponseBodyDataPageData {
	s.DelaySeconds = &v
	return s
}

func (s *ListQueueResponseBodyDataPageData) SetInactiveMessages(v int64) *ListQueueResponseBodyDataPageData {
	s.InactiveMessages = &v
	return s
}

func (s *ListQueueResponseBodyDataPageData) SetLastModifyTime(v int64) *ListQueueResponseBodyDataPageData {
	s.LastModifyTime = &v
	return s
}

func (s *ListQueueResponseBodyDataPageData) SetLoggingEnabled(v bool) *ListQueueResponseBodyDataPageData {
	s.LoggingEnabled = &v
	return s
}

func (s *ListQueueResponseBodyDataPageData) SetMaximumMessageSize(v int64) *ListQueueResponseBodyDataPageData {
	s.MaximumMessageSize = &v
	return s
}

func (s *ListQueueResponseBodyDataPageData) SetMessageRetentionPeriod(v int64) *ListQueueResponseBodyDataPageData {
	s.MessageRetentionPeriod = &v
	return s
}

func (s *ListQueueResponseBodyDataPageData) SetPollingWaitSeconds(v int64) *ListQueueResponseBodyDataPageData {
	s.PollingWaitSeconds = &v
	return s
}

func (s *ListQueueResponseBodyDataPageData) SetQueueName(v string) *ListQueueResponseBodyDataPageData {
	s.QueueName = &v
	return s
}

func (s *ListQueueResponseBodyDataPageData) SetTags(v []*ListQueueResponseBodyDataPageDataTags) *ListQueueResponseBodyDataPageData {
	s.Tags = v
	return s
}

func (s *ListQueueResponseBodyDataPageData) SetVisibilityTimeout(v int64) *ListQueueResponseBodyDataPageData {
	s.VisibilityTimeout = &v
	return s
}

type ListQueueResponseBodyDataPageDataTags struct {
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListQueueResponseBodyDataPageDataTags) String() string {
	return tea.Prettify(s)
}

func (s ListQueueResponseBodyDataPageDataTags) GoString() string {
	return s.String()
}

func (s *ListQueueResponseBodyDataPageDataTags) SetTagKey(v string) *ListQueueResponseBodyDataPageDataTags {
	s.TagKey = &v
	return s
}

func (s *ListQueueResponseBodyDataPageDataTags) SetTagValue(v string) *ListQueueResponseBodyDataPageDataTags {
	s.TagValue = &v
	return s
}

type ListQueueResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListQueueResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListQueueResponse) String() string {
	return tea.Prettify(s)
}

func (s ListQueueResponse) GoString() string {
	return s.String()
}

func (s *ListQueueResponse) SetHeaders(v map[string]*string) *ListQueueResponse {
	s.Headers = v
	return s
}

func (s *ListQueueResponse) SetStatusCode(v int32) *ListQueueResponse {
	s.StatusCode = &v
	return s
}

func (s *ListQueueResponse) SetBody(v *ListQueueResponseBody) *ListQueueResponse {
	s.Body = v
	return s
}

type ListSubscriptionByTopicRequest struct {
	// The page number. Valid values: 1 to 100000000. If you set this parameter to a value smaller than 1, the value of this parameter is 1 by default. If you set this parameter to a value greater than 100000000, the value of this parameter is 100000000 by default.
	//
	// example:
	//
	// 1
	PageNum *int64 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries per page. Value values: 10 to 50. If you set this parameter to a value smaller than 10, the value of this parameter is 10 by default. If you set this parameter to a value greater than 50, the value of this parameter is 50 by default.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The name of the subscription.
	//
	// example:
	//
	// demo-subscription
	SubscriptionName *string `json:"SubscriptionName,omitempty" xml:"SubscriptionName,omitempty"`
	// The name of the topic.
	//
	// example:
	//
	// test
	TopicName *string `json:"TopicName,omitempty" xml:"TopicName,omitempty"`
}

func (s ListSubscriptionByTopicRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSubscriptionByTopicRequest) GoString() string {
	return s.String()
}

func (s *ListSubscriptionByTopicRequest) SetPageNum(v int64) *ListSubscriptionByTopicRequest {
	s.PageNum = &v
	return s
}

func (s *ListSubscriptionByTopicRequest) SetPageSize(v int64) *ListSubscriptionByTopicRequest {
	s.PageSize = &v
	return s
}

func (s *ListSubscriptionByTopicRequest) SetSubscriptionName(v string) *ListSubscriptionByTopicRequest {
	s.SubscriptionName = &v
	return s
}

func (s *ListSubscriptionByTopicRequest) SetTopicName(v string) *ListSubscriptionByTopicRequest {
	s.TopicName = &v
	return s
}

type ListSubscriptionByTopicResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *ListSubscriptionByTopicResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s ListSubscriptionByTopicResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSubscriptionByTopicResponseBody) GoString() string {
	return s.String()
}

func (s *ListSubscriptionByTopicResponseBody) SetCode(v int64) *ListSubscriptionByTopicResponseBody {
	s.Code = &v
	return s
}

func (s *ListSubscriptionByTopicResponseBody) SetData(v *ListSubscriptionByTopicResponseBodyData) *ListSubscriptionByTopicResponseBody {
	s.Data = v
	return s
}

func (s *ListSubscriptionByTopicResponseBody) SetMessage(v string) *ListSubscriptionByTopicResponseBody {
	s.Message = &v
	return s
}

func (s *ListSubscriptionByTopicResponseBody) SetRequestId(v string) *ListSubscriptionByTopicResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSubscriptionByTopicResponseBody) SetStatus(v string) *ListSubscriptionByTopicResponseBody {
	s.Status = &v
	return s
}

func (s *ListSubscriptionByTopicResponseBody) SetSuccess(v bool) *ListSubscriptionByTopicResponseBody {
	s.Success = &v
	return s
}

type ListSubscriptionByTopicResponseBodyData struct {
	// The data returned on the current page.
	PageData []*ListSubscriptionByTopicResponseBodyDataPageData `json:"PageData,omitempty" xml:"PageData,omitempty" type:"Repeated"`
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
	// The total number of pages returned.
	//
	// example:
	//
	// 3
	Pages *int64 `json:"Pages,omitempty" xml:"Pages,omitempty"`
	// The number of entries on the current page.
	//
	// example:
	//
	// 20
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 130
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListSubscriptionByTopicResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListSubscriptionByTopicResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListSubscriptionByTopicResponseBodyData) SetPageData(v []*ListSubscriptionByTopicResponseBodyDataPageData) *ListSubscriptionByTopicResponseBodyData {
	s.PageData = v
	return s
}

func (s *ListSubscriptionByTopicResponseBodyData) SetPageNum(v int64) *ListSubscriptionByTopicResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *ListSubscriptionByTopicResponseBodyData) SetPageSize(v int64) *ListSubscriptionByTopicResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListSubscriptionByTopicResponseBodyData) SetPages(v int64) *ListSubscriptionByTopicResponseBodyData {
	s.Pages = &v
	return s
}

func (s *ListSubscriptionByTopicResponseBodyData) SetSize(v int64) *ListSubscriptionByTopicResponseBodyData {
	s.Size = &v
	return s
}

func (s *ListSubscriptionByTopicResponseBodyData) SetTotal(v int64) *ListSubscriptionByTopicResponseBodyData {
	s.Total = &v
	return s
}

type ListSubscriptionByTopicResponseBodyDataPageData struct {
	// The time when the subscription was created. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1449554806
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The endpoint to which the messages are pushed.
	//
	// example:
	//
	// http://example.com
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// The tag that is used to filter messages. Only the messages that are attached with the specified tag can be pushed.
	//
	// example:
	//
	// important
	FilterTag *string `json:"FilterTag,omitempty" xml:"FilterTag,omitempty"`
	// The time when the subscription was last modified. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1449554806
	LastModifyTime *int64 `json:"LastModifyTime,omitempty" xml:"LastModifyTime,omitempty"`
	// The content format of the messages that are pushed to the endpoint. Valid values:
	//
	// 	- XML
	//
	// 	- JSON
	//
	// 	- SIMPLIFIED
	//
	// example:
	//
	// XML
	NotifyContentFormat *string `json:"NotifyContentFormat,omitempty" xml:"NotifyContentFormat,omitempty"`
	// The retry policy that is applied if an error occurs when Message Service (MNS) pushes messages to the endpoint. Valid values:
	//
	// 	- BACKOFF_RETRY
	//
	// 	- EXPONENTIAL_DECAY_RETRY
	//
	// example:
	//
	// BACKOFF_RETRY
	NotifyStrategy *string `json:"NotifyStrategy,omitempty" xml:"NotifyStrategy,omitempty"`
	// The name of the subscription.
	//
	// example:
	//
	// MySubscription
	SubscriptionName *string `json:"SubscriptionName,omitempty" xml:"SubscriptionName,omitempty"`
	// The name of the topic.
	//
	// example:
	//
	// MyTopic
	TopicName *string `json:"TopicName,omitempty" xml:"TopicName,omitempty"`
	// The Alibaba Cloud account ID of the topic owner.
	//
	// example:
	//
	// 123456789098****
	TopicOwner *string `json:"TopicOwner,omitempty" xml:"TopicOwner,omitempty"`
}

func (s ListSubscriptionByTopicResponseBodyDataPageData) String() string {
	return tea.Prettify(s)
}

func (s ListSubscriptionByTopicResponseBodyDataPageData) GoString() string {
	return s.String()
}

func (s *ListSubscriptionByTopicResponseBodyDataPageData) SetCreateTime(v int64) *ListSubscriptionByTopicResponseBodyDataPageData {
	s.CreateTime = &v
	return s
}

func (s *ListSubscriptionByTopicResponseBodyDataPageData) SetEndpoint(v string) *ListSubscriptionByTopicResponseBodyDataPageData {
	s.Endpoint = &v
	return s
}

func (s *ListSubscriptionByTopicResponseBodyDataPageData) SetFilterTag(v string) *ListSubscriptionByTopicResponseBodyDataPageData {
	s.FilterTag = &v
	return s
}

func (s *ListSubscriptionByTopicResponseBodyDataPageData) SetLastModifyTime(v int64) *ListSubscriptionByTopicResponseBodyDataPageData {
	s.LastModifyTime = &v
	return s
}

func (s *ListSubscriptionByTopicResponseBodyDataPageData) SetNotifyContentFormat(v string) *ListSubscriptionByTopicResponseBodyDataPageData {
	s.NotifyContentFormat = &v
	return s
}

func (s *ListSubscriptionByTopicResponseBodyDataPageData) SetNotifyStrategy(v string) *ListSubscriptionByTopicResponseBodyDataPageData {
	s.NotifyStrategy = &v
	return s
}

func (s *ListSubscriptionByTopicResponseBodyDataPageData) SetSubscriptionName(v string) *ListSubscriptionByTopicResponseBodyDataPageData {
	s.SubscriptionName = &v
	return s
}

func (s *ListSubscriptionByTopicResponseBodyDataPageData) SetTopicName(v string) *ListSubscriptionByTopicResponseBodyDataPageData {
	s.TopicName = &v
	return s
}

func (s *ListSubscriptionByTopicResponseBodyDataPageData) SetTopicOwner(v string) *ListSubscriptionByTopicResponseBodyDataPageData {
	s.TopicOwner = &v
	return s
}

type ListSubscriptionByTopicResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSubscriptionByTopicResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSubscriptionByTopicResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSubscriptionByTopicResponse) GoString() string {
	return s.String()
}

func (s *ListSubscriptionByTopicResponse) SetHeaders(v map[string]*string) *ListSubscriptionByTopicResponse {
	s.Headers = v
	return s
}

func (s *ListSubscriptionByTopicResponse) SetStatusCode(v int32) *ListSubscriptionByTopicResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSubscriptionByTopicResponse) SetBody(v *ListSubscriptionByTopicResponseBody) *ListSubscriptionByTopicResponse {
	s.Body = v
	return s
}

type ListTopicRequest struct {
	// The page number. Valid values: 1 to 100000000. If you set this parameter to a value smaller than 1, the value of this parameter is 1 by default. If you set this parameter to a value greater than 100000000, the value of this parameter is 100000000 by default.
	//
	// example:
	//
	// 1
	PageNum *int64 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries per page. Value values: 10 to 50. If you set this parameter to a value smaller than 10, the value of this parameter is 10 by default. If you set this parameter to a value greater than 50, the value of this parameter is 50 by default.
	//
	// example:
	//
	// 20
	PageSize *int64                 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Tag      []*ListTopicRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The name of the topic.
	//
	// example:
	//
	// test
	TopicName *string `json:"TopicName,omitempty" xml:"TopicName,omitempty"`
}

func (s ListTopicRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTopicRequest) GoString() string {
	return s.String()
}

func (s *ListTopicRequest) SetPageNum(v int64) *ListTopicRequest {
	s.PageNum = &v
	return s
}

func (s *ListTopicRequest) SetPageSize(v int64) *ListTopicRequest {
	s.PageSize = &v
	return s
}

func (s *ListTopicRequest) SetTag(v []*ListTopicRequestTag) *ListTopicRequest {
	s.Tag = v
	return s
}

func (s *ListTopicRequest) SetTopicName(v string) *ListTopicRequest {
	s.TopicName = &v
	return s
}

type ListTopicRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTopicRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListTopicRequestTag) GoString() string {
	return s.String()
}

func (s *ListTopicRequestTag) SetKey(v string) *ListTopicRequestTag {
	s.Key = &v
	return s
}

func (s *ListTopicRequestTag) SetValue(v string) *ListTopicRequestTag {
	s.Value = &v
	return s
}

type ListTopicResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
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
	return tea.Prettify(s)
}

func (s ListTopicResponseBody) GoString() string {
	return s.String()
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
	return tea.Prettify(s)
}

func (s ListTopicResponseBodyData) GoString() string {
	return s.String()
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
	MessageRetentionPeriod *int64                                   `json:"MessageRetentionPeriod,omitempty" xml:"MessageRetentionPeriod,omitempty"`
	Tags                   []*ListTopicResponseBodyDataPageDataTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	TopicInnerUrl          *string                                  `json:"TopicInnerUrl,omitempty" xml:"TopicInnerUrl,omitempty"`
	// The name of the topic.
	//
	// example:
	//
	// demo-topic
	TopicName *string `json:"TopicName,omitempty" xml:"TopicName,omitempty"`
	TopicUrl  *string `json:"TopicUrl,omitempty" xml:"TopicUrl,omitempty"`
}

func (s ListTopicResponseBodyDataPageData) String() string {
	return tea.Prettify(s)
}

func (s ListTopicResponseBodyDataPageData) GoString() string {
	return s.String()
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

type ListTopicResponseBodyDataPageDataTags struct {
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListTopicResponseBodyDataPageDataTags) String() string {
	return tea.Prettify(s)
}

func (s ListTopicResponseBodyDataPageDataTags) GoString() string {
	return s.String()
}

func (s *ListTopicResponseBodyDataPageDataTags) SetTagKey(v string) *ListTopicResponseBodyDataPageDataTags {
	s.TagKey = &v
	return s
}

func (s *ListTopicResponseBodyDataPageDataTags) SetTagValue(v string) *ListTopicResponseBodyDataPageDataTags {
	s.TagValue = &v
	return s
}

type ListTopicResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTopicResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTopicResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTopicResponse) GoString() string {
	return s.String()
}

func (s *ListTopicResponse) SetHeaders(v map[string]*string) *ListTopicResponse {
	s.Headers = v
	return s
}

func (s *ListTopicResponse) SetStatusCode(v int32) *ListTopicResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTopicResponse) SetBody(v *ListTopicResponseBody) *ListTopicResponse {
	s.Body = v
	return s
}

type SetQueueAttributesRequest struct {
	// The period after which all messages sent to the queue are consumed. Valid values: 0 to 604800. Unit: seconds. Default value: 0
	//
	// example:
	//
	// 0
	DelaySeconds *int64 `json:"DelaySeconds,omitempty" xml:"DelaySeconds,omitempty"`
	// Specifies whether to enable the logging feature. Valid values:
	//
	// 	- True
	//
	// 	- False (default)
	//
	// example:
	//
	// True
	EnableLogging *bool `json:"EnableLogging,omitempty" xml:"EnableLogging,omitempty"`
	// The maximum length of the message that is sent to the queue. Valid values: 1024 to 65536. Unit: bytes. Default value: 65536.
	//
	// example:
	//
	// 1024
	MaximumMessageSize *int64 `json:"MaximumMessageSize,omitempty" xml:"MaximumMessageSize,omitempty"`
	// The maximum duration for which a message is retained in the queue. After the specified retention period ends, the message is deleted regardless of whether the message is received. Valid values: 60 to 604800. Unit: seconds. Default value: 345600.
	//
	// example:
	//
	// 120
	MessageRetentionPeriod *int64 `json:"MessageRetentionPeriod,omitempty" xml:"MessageRetentionPeriod,omitempty"`
	// The maximum duration for which long polling requests are held after the ReceiveMessage operation is called. Valid values: 0 to 30. Unit: seconds. Default value: 0
	//
	// example:
	//
	// 0
	PollingWaitSeconds *int64 `json:"PollingWaitSeconds,omitempty" xml:"PollingWaitSeconds,omitempty"`
	// The name of the queue.
	//
	// This parameter is required.
	//
	// example:
	//
	// testqueue
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	// The duration for which a message stays in the Inactive state after the message is received from the queue. Valid values: 1 to 43200. Unit: seconds. Default value: 30.
	//
	// example:
	//
	// 60
	VisibilityTimeout *int64 `json:"VisibilityTimeout,omitempty" xml:"VisibilityTimeout,omitempty"`
}

func (s SetQueueAttributesRequest) String() string {
	return tea.Prettify(s)
}

func (s SetQueueAttributesRequest) GoString() string {
	return s.String()
}

func (s *SetQueueAttributesRequest) SetDelaySeconds(v int64) *SetQueueAttributesRequest {
	s.DelaySeconds = &v
	return s
}

func (s *SetQueueAttributesRequest) SetEnableLogging(v bool) *SetQueueAttributesRequest {
	s.EnableLogging = &v
	return s
}

func (s *SetQueueAttributesRequest) SetMaximumMessageSize(v int64) *SetQueueAttributesRequest {
	s.MaximumMessageSize = &v
	return s
}

func (s *SetQueueAttributesRequest) SetMessageRetentionPeriod(v int64) *SetQueueAttributesRequest {
	s.MessageRetentionPeriod = &v
	return s
}

func (s *SetQueueAttributesRequest) SetPollingWaitSeconds(v int64) *SetQueueAttributesRequest {
	s.PollingWaitSeconds = &v
	return s
}

func (s *SetQueueAttributesRequest) SetQueueName(v string) *SetQueueAttributesRequest {
	s.QueueName = &v
	return s
}

func (s *SetQueueAttributesRequest) SetVisibilityTimeout(v int64) *SetQueueAttributesRequest {
	s.VisibilityTimeout = &v
	return s
}

type SetQueueAttributesResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *SetQueueAttributesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s SetQueueAttributesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetQueueAttributesResponseBody) GoString() string {
	return s.String()
}

func (s *SetQueueAttributesResponseBody) SetCode(v int64) *SetQueueAttributesResponseBody {
	s.Code = &v
	return s
}

func (s *SetQueueAttributesResponseBody) SetData(v *SetQueueAttributesResponseBodyData) *SetQueueAttributesResponseBody {
	s.Data = v
	return s
}

func (s *SetQueueAttributesResponseBody) SetMessage(v string) *SetQueueAttributesResponseBody {
	s.Message = &v
	return s
}

func (s *SetQueueAttributesResponseBody) SetRequestId(v string) *SetQueueAttributesResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetQueueAttributesResponseBody) SetStatus(v string) *SetQueueAttributesResponseBody {
	s.Status = &v
	return s
}

func (s *SetQueueAttributesResponseBody) SetSuccess(v bool) *SetQueueAttributesResponseBody {
	s.Success = &v
	return s
}

type SetQueueAttributesResponseBodyData struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SetQueueAttributesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SetQueueAttributesResponseBodyData) GoString() string {
	return s.String()
}

func (s *SetQueueAttributesResponseBodyData) SetCode(v int64) *SetQueueAttributesResponseBodyData {
	s.Code = &v
	return s
}

func (s *SetQueueAttributesResponseBodyData) SetMessage(v string) *SetQueueAttributesResponseBodyData {
	s.Message = &v
	return s
}

func (s *SetQueueAttributesResponseBodyData) SetSuccess(v bool) *SetQueueAttributesResponseBodyData {
	s.Success = &v
	return s
}

type SetQueueAttributesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetQueueAttributesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetQueueAttributesResponse) String() string {
	return tea.Prettify(s)
}

func (s SetQueueAttributesResponse) GoString() string {
	return s.String()
}

func (s *SetQueueAttributesResponse) SetHeaders(v map[string]*string) *SetQueueAttributesResponse {
	s.Headers = v
	return s
}

func (s *SetQueueAttributesResponse) SetStatusCode(v int32) *SetQueueAttributesResponse {
	s.StatusCode = &v
	return s
}

func (s *SetQueueAttributesResponse) SetBody(v *SetQueueAttributesResponseBody) *SetQueueAttributesResponse {
	s.Body = v
	return s
}

type SetSubscriptionAttributesRequest struct {
	// The retry policy that is applied if an error occurs when Message Service (MNS) pushes messages to the endpoint. Valid values:
	//
	// 	- BACKOFF_RETRY
	//
	// 	- EXPONENTIAL_DECAY_RETRY
	//
	// example:
	//
	// BACKOFF_RETRY
	NotifyStrategy *string `json:"NotifyStrategy,omitempty" xml:"NotifyStrategy,omitempty"`
	// The name of the subscription.
	//
	// This parameter is required.
	//
	// example:
	//
	// MySubscription
	SubscriptionName *string `json:"SubscriptionName,omitempty" xml:"SubscriptionName,omitempty"`
	// The name of the topic.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	TopicName *string `json:"TopicName,omitempty" xml:"TopicName,omitempty"`
}

func (s SetSubscriptionAttributesRequest) String() string {
	return tea.Prettify(s)
}

func (s SetSubscriptionAttributesRequest) GoString() string {
	return s.String()
}

func (s *SetSubscriptionAttributesRequest) SetNotifyStrategy(v string) *SetSubscriptionAttributesRequest {
	s.NotifyStrategy = &v
	return s
}

func (s *SetSubscriptionAttributesRequest) SetSubscriptionName(v string) *SetSubscriptionAttributesRequest {
	s.SubscriptionName = &v
	return s
}

func (s *SetSubscriptionAttributesRequest) SetTopicName(v string) *SetSubscriptionAttributesRequest {
	s.TopicName = &v
	return s
}

type SetSubscriptionAttributesResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *SetSubscriptionAttributesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s SetSubscriptionAttributesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetSubscriptionAttributesResponseBody) GoString() string {
	return s.String()
}

func (s *SetSubscriptionAttributesResponseBody) SetCode(v int64) *SetSubscriptionAttributesResponseBody {
	s.Code = &v
	return s
}

func (s *SetSubscriptionAttributesResponseBody) SetData(v *SetSubscriptionAttributesResponseBodyData) *SetSubscriptionAttributesResponseBody {
	s.Data = v
	return s
}

func (s *SetSubscriptionAttributesResponseBody) SetMessage(v string) *SetSubscriptionAttributesResponseBody {
	s.Message = &v
	return s
}

func (s *SetSubscriptionAttributesResponseBody) SetRequestId(v string) *SetSubscriptionAttributesResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetSubscriptionAttributesResponseBody) SetStatus(v string) *SetSubscriptionAttributesResponseBody {
	s.Status = &v
	return s
}

func (s *SetSubscriptionAttributesResponseBody) SetSuccess(v bool) *SetSubscriptionAttributesResponseBody {
	s.Success = &v
	return s
}

type SetSubscriptionAttributesResponseBodyData struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SetSubscriptionAttributesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SetSubscriptionAttributesResponseBodyData) GoString() string {
	return s.String()
}

func (s *SetSubscriptionAttributesResponseBodyData) SetCode(v int64) *SetSubscriptionAttributesResponseBodyData {
	s.Code = &v
	return s
}

func (s *SetSubscriptionAttributesResponseBodyData) SetMessage(v string) *SetSubscriptionAttributesResponseBodyData {
	s.Message = &v
	return s
}

func (s *SetSubscriptionAttributesResponseBodyData) SetSuccess(v bool) *SetSubscriptionAttributesResponseBodyData {
	s.Success = &v
	return s
}

type SetSubscriptionAttributesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetSubscriptionAttributesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetSubscriptionAttributesResponse) String() string {
	return tea.Prettify(s)
}

func (s SetSubscriptionAttributesResponse) GoString() string {
	return s.String()
}

func (s *SetSubscriptionAttributesResponse) SetHeaders(v map[string]*string) *SetSubscriptionAttributesResponse {
	s.Headers = v
	return s
}

func (s *SetSubscriptionAttributesResponse) SetStatusCode(v int32) *SetSubscriptionAttributesResponse {
	s.StatusCode = &v
	return s
}

func (s *SetSubscriptionAttributesResponse) SetBody(v *SetSubscriptionAttributesResponseBody) *SetSubscriptionAttributesResponse {
	s.Body = v
	return s
}

type SetTopicAttributesRequest struct {
	// Specifies whether to enable the logging feature. Valid values:
	//
	// 	- True
	//
	// 	- False (default)
	//
	// example:
	//
	// True
	EnableLogging *bool `json:"EnableLogging,omitempty" xml:"EnableLogging,omitempty"`
	// The maximum length of the message that is sent to the topic. Valid values: 1024 to 65536. Unit: bytes. Default value: 65536.
	//
	// example:
	//
	// 65536
	MaxMessageSize *int64 `json:"MaxMessageSize,omitempty" xml:"MaxMessageSize,omitempty"`
	// The name of the topic.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	TopicName *string `json:"TopicName,omitempty" xml:"TopicName,omitempty"`
}

func (s SetTopicAttributesRequest) String() string {
	return tea.Prettify(s)
}

func (s SetTopicAttributesRequest) GoString() string {
	return s.String()
}

func (s *SetTopicAttributesRequest) SetEnableLogging(v bool) *SetTopicAttributesRequest {
	s.EnableLogging = &v
	return s
}

func (s *SetTopicAttributesRequest) SetMaxMessageSize(v int64) *SetTopicAttributesRequest {
	s.MaxMessageSize = &v
	return s
}

func (s *SetTopicAttributesRequest) SetTopicName(v string) *SetTopicAttributesRequest {
	s.TopicName = &v
	return s
}

type SetTopicAttributesResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *SetTopicAttributesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s SetTopicAttributesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetTopicAttributesResponseBody) GoString() string {
	return s.String()
}

func (s *SetTopicAttributesResponseBody) SetCode(v int64) *SetTopicAttributesResponseBody {
	s.Code = &v
	return s
}

func (s *SetTopicAttributesResponseBody) SetData(v *SetTopicAttributesResponseBodyData) *SetTopicAttributesResponseBody {
	s.Data = v
	return s
}

func (s *SetTopicAttributesResponseBody) SetMessage(v string) *SetTopicAttributesResponseBody {
	s.Message = &v
	return s
}

func (s *SetTopicAttributesResponseBody) SetRequestId(v string) *SetTopicAttributesResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetTopicAttributesResponseBody) SetStatus(v string) *SetTopicAttributesResponseBody {
	s.Status = &v
	return s
}

func (s *SetTopicAttributesResponseBody) SetSuccess(v bool) *SetTopicAttributesResponseBody {
	s.Success = &v
	return s
}

type SetTopicAttributesResponseBodyData struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SetTopicAttributesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SetTopicAttributesResponseBodyData) GoString() string {
	return s.String()
}

func (s *SetTopicAttributesResponseBodyData) SetCode(v int64) *SetTopicAttributesResponseBodyData {
	s.Code = &v
	return s
}

func (s *SetTopicAttributesResponseBodyData) SetMessage(v string) *SetTopicAttributesResponseBodyData {
	s.Message = &v
	return s
}

func (s *SetTopicAttributesResponseBodyData) SetSuccess(v bool) *SetTopicAttributesResponseBodyData {
	s.Success = &v
	return s
}

type SetTopicAttributesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetTopicAttributesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetTopicAttributesResponse) String() string {
	return tea.Prettify(s)
}

func (s SetTopicAttributesResponse) GoString() string {
	return s.String()
}

func (s *SetTopicAttributesResponse) SetHeaders(v map[string]*string) *SetTopicAttributesResponse {
	s.Headers = v
	return s
}

func (s *SetTopicAttributesResponse) SetStatusCode(v int32) *SetTopicAttributesResponse {
	s.StatusCode = &v
	return s
}

func (s *SetTopicAttributesResponse) SetBody(v *SetTopicAttributesResponseBody) *SetTopicAttributesResponse {
	s.Body = v
	return s
}

type SubscribeRequest struct {
	// The receiver endpoint. The format of the endpoint varies based on the terminal type.
	//
	// 	- If you set PushType to http, set Endpoint to an `HTTP URL that starts with http:// or https://`.
	//
	// 	- If you set PushType to queue, set Endpoint to a `queue name`.
	//
	// 	- If you set PushType to mpush, set Endpoint to an `AppKey`.
	//
	// 	- If you set PushType to alisms, set Endpoint to a `mobile number`.
	//
	// 	- If you set PushType to email, set Endpoint to an `email address`.
	//
	// This parameter is required.
	//
	// example:
	//
	// http://example.com
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// The tag that is used to filter messages. Only messages that have the same tag can be pushed. Set the value to a string of no more than 16 characters.
	//
	// By default, no tag is specified to filter messages.
	//
	// example:
	//
	// important
	MessageTag *string `json:"MessageTag,omitempty" xml:"MessageTag,omitempty"`
	// The content format of the messages that are pushed to the endpoint. Valid values:
	//
	// 	- XML
	//
	// 	- JSON
	//
	// 	- SIMPLIFIED
	//
	// example:
	//
	// XML
	NotifyContentFormat *string `json:"NotifyContentFormat,omitempty" xml:"NotifyContentFormat,omitempty"`
	// The retry policy that is applied if an error occurs when Message Service (MNS) pushes messages to the endpoint. Valid values:
	//
	// 	- BACKOFF_RETRY
	//
	// 	- EXPONENTIAL_DECAY_RETRY
	//
	// example:
	//
	// BACKOFF_RETRY
	NotifyStrategy *string `json:"NotifyStrategy,omitempty" xml:"NotifyStrategy,omitempty"`
	// The terminal type. Valid values:
	//
	// 	- http: HTTP services
	//
	// 	- queue: queues
	//
	// 	- mpush: mobile devices
	//
	// 	- alisms: Alibaba Cloud Short Message Service (SMS)
	//
	// 	- email: emails
	//
	// This parameter is required.
	//
	// example:
	//
	// queue
	PushType *string `json:"PushType,omitempty" xml:"PushType,omitempty"`
	// The name of the subscription.
	//
	// This parameter is required.
	//
	// example:
	//
	// testSubscription
	SubscriptionName *string `json:"SubscriptionName,omitempty" xml:"SubscriptionName,omitempty"`
	// The name of the topic.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	TopicName *string `json:"TopicName,omitempty" xml:"TopicName,omitempty"`
}

func (s SubscribeRequest) String() string {
	return tea.Prettify(s)
}

func (s SubscribeRequest) GoString() string {
	return s.String()
}

func (s *SubscribeRequest) SetEndpoint(v string) *SubscribeRequest {
	s.Endpoint = &v
	return s
}

func (s *SubscribeRequest) SetMessageTag(v string) *SubscribeRequest {
	s.MessageTag = &v
	return s
}

func (s *SubscribeRequest) SetNotifyContentFormat(v string) *SubscribeRequest {
	s.NotifyContentFormat = &v
	return s
}

func (s *SubscribeRequest) SetNotifyStrategy(v string) *SubscribeRequest {
	s.NotifyStrategy = &v
	return s
}

func (s *SubscribeRequest) SetPushType(v string) *SubscribeRequest {
	s.PushType = &v
	return s
}

func (s *SubscribeRequest) SetSubscriptionName(v string) *SubscribeRequest {
	s.SubscriptionName = &v
	return s
}

func (s *SubscribeRequest) SetTopicName(v string) *SubscribeRequest {
	s.TopicName = &v
	return s
}

type SubscribeResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	//
	// example:
	//
	// {\\"Code\\": 200, \\"Success\\": True}
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
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

func (s SubscribeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubscribeResponseBody) GoString() string {
	return s.String()
}

func (s *SubscribeResponseBody) SetCode(v int64) *SubscribeResponseBody {
	s.Code = &v
	return s
}

func (s *SubscribeResponseBody) SetData(v string) *SubscribeResponseBody {
	s.Data = &v
	return s
}

func (s *SubscribeResponseBody) SetMessage(v string) *SubscribeResponseBody {
	s.Message = &v
	return s
}

func (s *SubscribeResponseBody) SetRequestId(v string) *SubscribeResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubscribeResponseBody) SetStatus(v string) *SubscribeResponseBody {
	s.Status = &v
	return s
}

func (s *SubscribeResponseBody) SetSuccess(v bool) *SubscribeResponseBody {
	s.Success = &v
	return s
}

type SubscribeResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubscribeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubscribeResponse) String() string {
	return tea.Prettify(s)
}

func (s SubscribeResponse) GoString() string {
	return s.String()
}

func (s *SubscribeResponse) SetHeaders(v map[string]*string) *SubscribeResponse {
	s.Headers = v
	return s
}

func (s *SubscribeResponse) SetStatusCode(v int32) *SubscribeResponse {
	s.StatusCode = &v
	return s
}

func (s *SubscribeResponse) SetBody(v *SubscribeResponseBody) *SubscribeResponse {
	s.Body = v
	return s
}

type UnsubscribeRequest struct {
	// The name of the subscription.
	//
	// This parameter is required.
	//
	// example:
	//
	// MySubscription
	SubscriptionName *string `json:"SubscriptionName,omitempty" xml:"SubscriptionName,omitempty"`
	// The name of the topic.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	TopicName *string `json:"TopicName,omitempty" xml:"TopicName,omitempty"`
}

func (s UnsubscribeRequest) String() string {
	return tea.Prettify(s)
}

func (s UnsubscribeRequest) GoString() string {
	return s.String()
}

func (s *UnsubscribeRequest) SetSubscriptionName(v string) *UnsubscribeRequest {
	s.SubscriptionName = &v
	return s
}

func (s *UnsubscribeRequest) SetTopicName(v string) *UnsubscribeRequest {
	s.TopicName = &v
	return s
}

type UnsubscribeResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *UnsubscribeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s UnsubscribeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnsubscribeResponseBody) GoString() string {
	return s.String()
}

func (s *UnsubscribeResponseBody) SetCode(v int64) *UnsubscribeResponseBody {
	s.Code = &v
	return s
}

func (s *UnsubscribeResponseBody) SetData(v *UnsubscribeResponseBodyData) *UnsubscribeResponseBody {
	s.Data = v
	return s
}

func (s *UnsubscribeResponseBody) SetMessage(v string) *UnsubscribeResponseBody {
	s.Message = &v
	return s
}

func (s *UnsubscribeResponseBody) SetRequestId(v string) *UnsubscribeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnsubscribeResponseBody) SetStatus(v string) *UnsubscribeResponseBody {
	s.Status = &v
	return s
}

func (s *UnsubscribeResponseBody) SetSuccess(v bool) *UnsubscribeResponseBody {
	s.Success = &v
	return s
}

type UnsubscribeResponseBodyData struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UnsubscribeResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s UnsubscribeResponseBodyData) GoString() string {
	return s.String()
}

func (s *UnsubscribeResponseBodyData) SetCode(v int64) *UnsubscribeResponseBodyData {
	s.Code = &v
	return s
}

func (s *UnsubscribeResponseBodyData) SetMessage(v string) *UnsubscribeResponseBodyData {
	s.Message = &v
	return s
}

func (s *UnsubscribeResponseBodyData) SetSuccess(v bool) *UnsubscribeResponseBodyData {
	s.Success = &v
	return s
}

type UnsubscribeResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnsubscribeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnsubscribeResponse) String() string {
	return tea.Prettify(s)
}

func (s UnsubscribeResponse) GoString() string {
	return s.String()
}

func (s *UnsubscribeResponse) SetHeaders(v map[string]*string) *UnsubscribeResponse {
	s.Headers = v
	return s
}

func (s *UnsubscribeResponse) SetStatusCode(v int32) *UnsubscribeResponse {
	s.StatusCode = &v
	return s
}

func (s *UnsubscribeResponse) SetBody(v *UnsubscribeResponseBody) *UnsubscribeResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("regional")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("mns-open"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a queue.
//
// @param request - CreateQueueRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateQueueResponse
func (client *Client) CreateQueueWithOptions(request *CreateQueueRequest, runtime *util.RuntimeOptions) (_result *CreateQueueResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DelaySeconds)) {
		query["DelaySeconds"] = request.DelaySeconds
	}

	if !tea.BoolValue(util.IsUnset(request.EnableLogging)) {
		query["EnableLogging"] = request.EnableLogging
	}

	if !tea.BoolValue(util.IsUnset(request.MaximumMessageSize)) {
		query["MaximumMessageSize"] = request.MaximumMessageSize
	}

	if !tea.BoolValue(util.IsUnset(request.MessageRetentionPeriod)) {
		query["MessageRetentionPeriod"] = request.MessageRetentionPeriod
	}

	if !tea.BoolValue(util.IsUnset(request.PollingWaitSeconds)) {
		query["PollingWaitSeconds"] = request.PollingWaitSeconds
	}

	if !tea.BoolValue(util.IsUnset(request.QueueName)) {
		query["QueueName"] = request.QueueName
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	if !tea.BoolValue(util.IsUnset(request.VisibilityTimeout)) {
		query["VisibilityTimeout"] = request.VisibilityTimeout
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateQueue"),
		Version:     tea.String("2022-01-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateQueueResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a queue.
//
// @param request - CreateQueueRequest
//
// @return CreateQueueResponse
func (client *Client) CreateQueue(request *CreateQueueRequest) (_result *CreateQueueResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateQueueResponse{}
	_body, _err := client.CreateQueueWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a topic.
//
// @param request - CreateTopicRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTopicResponse
func (client *Client) CreateTopicWithOptions(request *CreateTopicRequest, runtime *util.RuntimeOptions) (_result *CreateTopicResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EnableLogging)) {
		body["EnableLogging"] = request.EnableLogging
	}

	if !tea.BoolValue(util.IsUnset(request.MaxMessageSize)) {
		body["MaxMessageSize"] = request.MaxMessageSize
	}

	if !tea.BoolValue(util.IsUnset(request.TopicName)) {
		body["TopicName"] = request.TopicName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateTopic"),
		Version:     tea.String("2022-01-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateTopicResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a topic.
//
// @param request - CreateTopicRequest
//
// @return CreateTopicResponse
func (client *Client) CreateTopic(request *CreateTopicRequest) (_result *CreateTopicResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateTopicResponse{}
	_body, _err := client.CreateTopicWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a queue.
//
// @param request - DeleteQueueRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteQueueResponse
func (client *Client) DeleteQueueWithOptions(request *DeleteQueueRequest, runtime *util.RuntimeOptions) (_result *DeleteQueueResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.QueueName)) {
		query["QueueName"] = request.QueueName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteQueue"),
		Version:     tea.String("2022-01-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteQueueResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a queue.
//
// @param request - DeleteQueueRequest
//
// @return DeleteQueueResponse
func (client *Client) DeleteQueue(request *DeleteQueueRequest) (_result *DeleteQueueResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteQueueResponse{}
	_body, _err := client.DeleteQueueWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a topic.
//
// @param request - DeleteTopicRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTopicResponse
func (client *Client) DeleteTopicWithOptions(request *DeleteTopicRequest, runtime *util.RuntimeOptions) (_result *DeleteTopicResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TopicName)) {
		query["TopicName"] = request.TopicName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteTopic"),
		Version:     tea.String("2022-01-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteTopicResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a topic.
//
// @param request - DeleteTopicRequest
//
// @return DeleteTopicResponse
func (client *Client) DeleteTopic(request *DeleteTopicRequest) (_result *DeleteTopicResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteTopicResponse{}
	_body, _err := client.DeleteTopicWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the attributes of an existing queue.
//
// @param request - GetQueueAttributesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetQueueAttributesResponse
func (client *Client) GetQueueAttributesWithOptions(request *GetQueueAttributesRequest, runtime *util.RuntimeOptions) (_result *GetQueueAttributesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.QueueName)) {
		query["QueueName"] = request.QueueName
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetQueueAttributes"),
		Version:     tea.String("2022-01-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetQueueAttributesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the attributes of an existing queue.
//
// @param request - GetQueueAttributesRequest
//
// @return GetQueueAttributesResponse
func (client *Client) GetQueueAttributes(request *GetQueueAttributesRequest) (_result *GetQueueAttributesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetQueueAttributesResponse{}
	_body, _err := client.GetQueueAttributesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the attributes of a subscription.
//
// @param request - GetSubscriptionAttributesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSubscriptionAttributesResponse
func (client *Client) GetSubscriptionAttributesWithOptions(request *GetSubscriptionAttributesRequest, runtime *util.RuntimeOptions) (_result *GetSubscriptionAttributesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SubscriptionName)) {
		query["SubscriptionName"] = request.SubscriptionName
	}

	if !tea.BoolValue(util.IsUnset(request.TopicName)) {
		query["TopicName"] = request.TopicName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetSubscriptionAttributes"),
		Version:     tea.String("2022-01-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSubscriptionAttributesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the attributes of a subscription.
//
// @param request - GetSubscriptionAttributesRequest
//
// @return GetSubscriptionAttributesResponse
func (client *Client) GetSubscriptionAttributes(request *GetSubscriptionAttributesRequest) (_result *GetSubscriptionAttributesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSubscriptionAttributesResponse{}
	_body, _err := client.GetSubscriptionAttributesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the attributes of a topic.
//
// @param request - GetTopicAttributesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTopicAttributesResponse
func (client *Client) GetTopicAttributesWithOptions(request *GetTopicAttributesRequest, runtime *util.RuntimeOptions) (_result *GetTopicAttributesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	if !tea.BoolValue(util.IsUnset(request.TopicName)) {
		query["TopicName"] = request.TopicName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTopicAttributes"),
		Version:     tea.String("2022-01-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTopicAttributesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the attributes of a topic.
//
// @param request - GetTopicAttributesRequest
//
// @return GetTopicAttributesResponse
func (client *Client) GetTopicAttributes(request *GetTopicAttributesRequest) (_result *GetTopicAttributesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTopicAttributesResponse{}
	_body, _err := client.GetTopicAttributesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries all queues that belong to an Alibaba Cloud account. The queues are displayed by page.
//
// @param request - ListQueueRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListQueueResponse
func (client *Client) ListQueueWithOptions(request *ListQueueRequest, runtime *util.RuntimeOptions) (_result *ListQueueResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		query["PageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.QueueName)) {
		query["QueueName"] = request.QueueName
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListQueue"),
		Version:     tea.String("2022-01-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListQueueResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all queues that belong to an Alibaba Cloud account. The queues are displayed by page.
//
// @param request - ListQueueRequest
//
// @return ListQueueResponse
func (client *Client) ListQueue(request *ListQueueRequest) (_result *ListQueueResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListQueueResponse{}
	_body, _err := client.ListQueueWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries all subscriptions to a topic. The subscriptions are displayed by page.
//
// @param request - ListSubscriptionByTopicRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSubscriptionByTopicResponse
func (client *Client) ListSubscriptionByTopicWithOptions(request *ListSubscriptionByTopicRequest, runtime *util.RuntimeOptions) (_result *ListSubscriptionByTopicResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		query["PageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SubscriptionName)) {
		query["SubscriptionName"] = request.SubscriptionName
	}

	if !tea.BoolValue(util.IsUnset(request.TopicName)) {
		query["TopicName"] = request.TopicName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListSubscriptionByTopic"),
		Version:     tea.String("2022-01-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSubscriptionByTopicResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all subscriptions to a topic. The subscriptions are displayed by page.
//
// @param request - ListSubscriptionByTopicRequest
//
// @return ListSubscriptionByTopicResponse
func (client *Client) ListSubscriptionByTopic(request *ListSubscriptionByTopicRequest) (_result *ListSubscriptionByTopicResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSubscriptionByTopicResponse{}
	_body, _err := client.ListSubscriptionByTopicWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the topics that belong to an Alibaba Cloud account. The topics are displayed by page.
//
// @param request - ListTopicRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTopicResponse
func (client *Client) ListTopicWithOptions(request *ListTopicRequest, runtime *util.RuntimeOptions) (_result *ListTopicResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		query["PageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	if !tea.BoolValue(util.IsUnset(request.TopicName)) {
		query["TopicName"] = request.TopicName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTopic"),
		Version:     tea.String("2022-01-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTopicResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the topics that belong to an Alibaba Cloud account. The topics are displayed by page.
//
// @param request - ListTopicRequest
//
// @return ListTopicResponse
func (client *Client) ListTopic(request *ListTopicRequest) (_result *ListTopicResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTopicResponse{}
	_body, _err := client.ListTopicWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies a queue.
//
// @param request - SetQueueAttributesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetQueueAttributesResponse
func (client *Client) SetQueueAttributesWithOptions(request *SetQueueAttributesRequest, runtime *util.RuntimeOptions) (_result *SetQueueAttributesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DelaySeconds)) {
		query["DelaySeconds"] = request.DelaySeconds
	}

	if !tea.BoolValue(util.IsUnset(request.EnableLogging)) {
		query["EnableLogging"] = request.EnableLogging
	}

	if !tea.BoolValue(util.IsUnset(request.MaximumMessageSize)) {
		query["MaximumMessageSize"] = request.MaximumMessageSize
	}

	if !tea.BoolValue(util.IsUnset(request.MessageRetentionPeriod)) {
		query["MessageRetentionPeriod"] = request.MessageRetentionPeriod
	}

	if !tea.BoolValue(util.IsUnset(request.PollingWaitSeconds)) {
		query["PollingWaitSeconds"] = request.PollingWaitSeconds
	}

	if !tea.BoolValue(util.IsUnset(request.QueueName)) {
		query["QueueName"] = request.QueueName
	}

	if !tea.BoolValue(util.IsUnset(request.VisibilityTimeout)) {
		query["VisibilityTimeout"] = request.VisibilityTimeout
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetQueueAttributes"),
		Version:     tea.String("2022-01-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetQueueAttributesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a queue.
//
// @param request - SetQueueAttributesRequest
//
// @return SetQueueAttributesResponse
func (client *Client) SetQueueAttributes(request *SetQueueAttributesRequest) (_result *SetQueueAttributesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetQueueAttributesResponse{}
	_body, _err := client.SetQueueAttributesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the attributes of a subscription.
//
// @param request - SetSubscriptionAttributesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetSubscriptionAttributesResponse
func (client *Client) SetSubscriptionAttributesWithOptions(request *SetSubscriptionAttributesRequest, runtime *util.RuntimeOptions) (_result *SetSubscriptionAttributesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NotifyStrategy)) {
		query["NotifyStrategy"] = request.NotifyStrategy
	}

	if !tea.BoolValue(util.IsUnset(request.SubscriptionName)) {
		query["SubscriptionName"] = request.SubscriptionName
	}

	if !tea.BoolValue(util.IsUnset(request.TopicName)) {
		query["TopicName"] = request.TopicName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetSubscriptionAttributes"),
		Version:     tea.String("2022-01-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetSubscriptionAttributesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the attributes of a subscription.
//
// @param request - SetSubscriptionAttributesRequest
//
// @return SetSubscriptionAttributesResponse
func (client *Client) SetSubscriptionAttributes(request *SetSubscriptionAttributesRequest) (_result *SetSubscriptionAttributesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetSubscriptionAttributesResponse{}
	_body, _err := client.SetSubscriptionAttributesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the attributes of a topic.
//
// @param request - SetTopicAttributesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetTopicAttributesResponse
func (client *Client) SetTopicAttributesWithOptions(request *SetTopicAttributesRequest, runtime *util.RuntimeOptions) (_result *SetTopicAttributesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EnableLogging)) {
		query["EnableLogging"] = request.EnableLogging
	}

	if !tea.BoolValue(util.IsUnset(request.MaxMessageSize)) {
		query["MaxMessageSize"] = request.MaxMessageSize
	}

	if !tea.BoolValue(util.IsUnset(request.TopicName)) {
		query["TopicName"] = request.TopicName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetTopicAttributes"),
		Version:     tea.String("2022-01-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetTopicAttributesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the attributes of a topic.
//
// @param request - SetTopicAttributesRequest
//
// @return SetTopicAttributesResponse
func (client *Client) SetTopicAttributes(request *SetTopicAttributesRequest) (_result *SetTopicAttributesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetTopicAttributesResponse{}
	_body, _err := client.SetTopicAttributesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a subscription to a topic.
//
// @param request - SubscribeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubscribeResponse
func (client *Client) SubscribeWithOptions(request *SubscribeRequest, runtime *util.RuntimeOptions) (_result *SubscribeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Endpoint)) {
		query["Endpoint"] = request.Endpoint
	}

	if !tea.BoolValue(util.IsUnset(request.MessageTag)) {
		query["MessageTag"] = request.MessageTag
	}

	if !tea.BoolValue(util.IsUnset(request.NotifyContentFormat)) {
		query["NotifyContentFormat"] = request.NotifyContentFormat
	}

	if !tea.BoolValue(util.IsUnset(request.NotifyStrategy)) {
		query["NotifyStrategy"] = request.NotifyStrategy
	}

	if !tea.BoolValue(util.IsUnset(request.PushType)) {
		query["PushType"] = request.PushType
	}

	if !tea.BoolValue(util.IsUnset(request.SubscriptionName)) {
		query["SubscriptionName"] = request.SubscriptionName
	}

	if !tea.BoolValue(util.IsUnset(request.TopicName)) {
		query["TopicName"] = request.TopicName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("Subscribe"),
		Version:     tea.String("2022-01-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SubscribeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a subscription to a topic.
//
// @param request - SubscribeRequest
//
// @return SubscribeResponse
func (client *Client) Subscribe(request *SubscribeRequest) (_result *SubscribeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubscribeResponse{}
	_body, _err := client.SubscribeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a subscription.
//
// @param request - UnsubscribeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnsubscribeResponse
func (client *Client) UnsubscribeWithOptions(request *UnsubscribeRequest, runtime *util.RuntimeOptions) (_result *UnsubscribeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SubscriptionName)) {
		query["SubscriptionName"] = request.SubscriptionName
	}

	if !tea.BoolValue(util.IsUnset(request.TopicName)) {
		query["TopicName"] = request.TopicName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("Unsubscribe"),
		Version:     tea.String("2022-01-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UnsubscribeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a subscription.
//
// @param request - UnsubscribeRequest
//
// @return UnsubscribeResponse
func (client *Client) Unsubscribe(request *UnsubscribeRequest) (_result *UnsubscribeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UnsubscribeResponse{}
	_body, _err := client.UnsubscribeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
