// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnsMessagePageQueryByTopicRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBeginTime(v int64) *OnsMessagePageQueryByTopicRequest
	GetBeginTime() *int64
	SetCurrentPage(v int32) *OnsMessagePageQueryByTopicRequest
	GetCurrentPage() *int32
	SetEndTime(v int64) *OnsMessagePageQueryByTopicRequest
	GetEndTime() *int64
	SetInstanceId(v string) *OnsMessagePageQueryByTopicRequest
	GetInstanceId() *string
	SetPageSize(v int32) *OnsMessagePageQueryByTopicRequest
	GetPageSize() *int32
	SetTaskId(v string) *OnsMessagePageQueryByTopicRequest
	GetTaskId() *string
	SetTopic(v string) *OnsMessagePageQueryByTopicRequest
	GetTopic() *string
}

type OnsMessagePageQueryByTopicRequest struct {
	// The beginning of the time range to query. The value of this parameter is a UNIX timestamp in milliseconds. If you specify a valid value for the **TaskId*	- parameter in the request, this parameter does not take effect. The system uses the value of the BeginTime parameter that you specified in the request when you created the specified query task.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1570723200000
	BeginTime *int64 `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// The number of the page to return. Pages start from page 1. Valid values: 1 to 50.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The end of the time range to query. The value of this parameter is a UNIX timestamp in milliseconds. If you specify a valid value for the **TaskId*	- parameter in the request, this parameter does not take effect. The system uses the value of the EndTime parameter that you specified in the request when you created the specified query task.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1570809600000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the instance to which the message you want to query belongs.
	//
	// example:
	//
	// MQ_INST_111111111111_DOxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of entries to return on each page. Valid values: 5 to 50. Default value: 20. If you specify a valid value for the **TaskId*	- parameter in the request, this parameter does not take effect. The system uses the value of the PageSize parameter that you specified in the request for creating the query task.
	//
	// example:
	//
	// 5
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the query task. The first time you call this operation to query the information about messages in a specified topic within a specified time range, this parameter is not required. This parameter is required in subsequent queries for messages on a specified page. You can obtain the task ID from the returned result of the first query.
	//
	// example:
	//
	// 0BC1310300002A9F000021E4D7A48346
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The name of the topic whose messages you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-mq_topic
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s OnsMessagePageQueryByTopicRequest) String() string {
	return dara.Prettify(s)
}

func (s OnsMessagePageQueryByTopicRequest) GoString() string {
	return s.String()
}

func (s *OnsMessagePageQueryByTopicRequest) GetBeginTime() *int64 {
	return s.BeginTime
}

func (s *OnsMessagePageQueryByTopicRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *OnsMessagePageQueryByTopicRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *OnsMessagePageQueryByTopicRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *OnsMessagePageQueryByTopicRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *OnsMessagePageQueryByTopicRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *OnsMessagePageQueryByTopicRequest) GetTopic() *string {
	return s.Topic
}

func (s *OnsMessagePageQueryByTopicRequest) SetBeginTime(v int64) *OnsMessagePageQueryByTopicRequest {
	s.BeginTime = &v
	return s
}

func (s *OnsMessagePageQueryByTopicRequest) SetCurrentPage(v int32) *OnsMessagePageQueryByTopicRequest {
	s.CurrentPage = &v
	return s
}

func (s *OnsMessagePageQueryByTopicRequest) SetEndTime(v int64) *OnsMessagePageQueryByTopicRequest {
	s.EndTime = &v
	return s
}

func (s *OnsMessagePageQueryByTopicRequest) SetInstanceId(v string) *OnsMessagePageQueryByTopicRequest {
	s.InstanceId = &v
	return s
}

func (s *OnsMessagePageQueryByTopicRequest) SetPageSize(v int32) *OnsMessagePageQueryByTopicRequest {
	s.PageSize = &v
	return s
}

func (s *OnsMessagePageQueryByTopicRequest) SetTaskId(v string) *OnsMessagePageQueryByTopicRequest {
	s.TaskId = &v
	return s
}

func (s *OnsMessagePageQueryByTopicRequest) SetTopic(v string) *OnsMessagePageQueryByTopicRequest {
	s.Topic = &v
	return s
}

func (s *OnsMessagePageQueryByTopicRequest) Validate() error {
	return dara.Validate(s)
}
