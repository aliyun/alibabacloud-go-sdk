// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnsDLQMessagePageQueryByGroupIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBeginTime(v int64) *OnsDLQMessagePageQueryByGroupIdRequest
	GetBeginTime() *int64
	SetCurrentPage(v int32) *OnsDLQMessagePageQueryByGroupIdRequest
	GetCurrentPage() *int32
	SetEndTime(v int64) *OnsDLQMessagePageQueryByGroupIdRequest
	GetEndTime() *int64
	SetGroupId(v string) *OnsDLQMessagePageQueryByGroupIdRequest
	GetGroupId() *string
	SetInstanceId(v string) *OnsDLQMessagePageQueryByGroupIdRequest
	GetInstanceId() *string
	SetPageSize(v int32) *OnsDLQMessagePageQueryByGroupIdRequest
	GetPageSize() *int32
	SetTaskId(v string) *OnsDLQMessagePageQueryByGroupIdRequest
	GetTaskId() *string
}

type OnsDLQMessagePageQueryByGroupIdRequest struct {
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
	// The ID of the consumer group whose dead-letter messages you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// GID_test_group_id
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The ID of the instance to which the dead-letter messages you want to query belong.
	//
	// example:
	//
	// MQ_INST_111111111111_DOxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of dead-letter messages to return on each page. Valid values: 5 to 50. Default value: 20. If you specify a valid value for the **TaskId*	- parameter in the request, this parameter does not take effect. The system uses the value of the PageSize parameter that you specified in the request when you created the specified query task.
	//
	// example:
	//
	// 5
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the query task. The first time you call this operation to query dead-letter messages that are sent to a specified consumer group within a specified time range, this parameter is not required. This parameter is required in subsequent queries for dead-letter messages on a specified page. You can obtain the task ID from the returned result of the first query.
	//
	// example:
	//
	// 0BC1310300002A9F000021E4D7A48346
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s OnsDLQMessagePageQueryByGroupIdRequest) String() string {
	return dara.Prettify(s)
}

func (s OnsDLQMessagePageQueryByGroupIdRequest) GoString() string {
	return s.String()
}

func (s *OnsDLQMessagePageQueryByGroupIdRequest) GetBeginTime() *int64 {
	return s.BeginTime
}

func (s *OnsDLQMessagePageQueryByGroupIdRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *OnsDLQMessagePageQueryByGroupIdRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *OnsDLQMessagePageQueryByGroupIdRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *OnsDLQMessagePageQueryByGroupIdRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *OnsDLQMessagePageQueryByGroupIdRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *OnsDLQMessagePageQueryByGroupIdRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *OnsDLQMessagePageQueryByGroupIdRequest) SetBeginTime(v int64) *OnsDLQMessagePageQueryByGroupIdRequest {
	s.BeginTime = &v
	return s
}

func (s *OnsDLQMessagePageQueryByGroupIdRequest) SetCurrentPage(v int32) *OnsDLQMessagePageQueryByGroupIdRequest {
	s.CurrentPage = &v
	return s
}

func (s *OnsDLQMessagePageQueryByGroupIdRequest) SetEndTime(v int64) *OnsDLQMessagePageQueryByGroupIdRequest {
	s.EndTime = &v
	return s
}

func (s *OnsDLQMessagePageQueryByGroupIdRequest) SetGroupId(v string) *OnsDLQMessagePageQueryByGroupIdRequest {
	s.GroupId = &v
	return s
}

func (s *OnsDLQMessagePageQueryByGroupIdRequest) SetInstanceId(v string) *OnsDLQMessagePageQueryByGroupIdRequest {
	s.InstanceId = &v
	return s
}

func (s *OnsDLQMessagePageQueryByGroupIdRequest) SetPageSize(v int32) *OnsDLQMessagePageQueryByGroupIdRequest {
	s.PageSize = &v
	return s
}

func (s *OnsDLQMessagePageQueryByGroupIdRequest) SetTaskId(v string) *OnsDLQMessagePageQueryByGroupIdRequest {
	s.TaskId = &v
	return s
}

func (s *OnsDLQMessagePageQueryByGroupIdRequest) Validate() error {
	return dara.Validate(s)
}
