// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTopicsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBeginTime(v string) *ListTopicsRequest
	GetBeginTime() *string
	SetEndTime(v string) *ListTopicsRequest
	GetEndTime() *string
	SetInstanceId(v int64) *ListTopicsRequest
	GetInstanceId() *int64
	SetNodeId(v int64) *ListTopicsRequest
	GetNodeId() *int64
	SetOwner(v string) *ListTopicsRequest
	GetOwner() *string
	SetPageNumber(v int32) *ListTopicsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListTopicsRequest
	GetPageSize() *int32
	SetTopicStatuses(v string) *ListTopicsRequest
	GetTopicStatuses() *string
	SetTopicTypes(v string) *ListTopicsRequest
	GetTopicTypes() *string
}

type ListTopicsRequest struct {
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-dd\\"T\\"HH:mm:ssZ format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2019-03-24T00:00:00+0800
	BeginTime *string `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-dd\\"T\\"HH:mm:ssZ format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2019-03-25T23:00:00+0800
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the node instance that triggers the events. You can configure either this parameter or the NodeId parameter.
	//
	// example:
	//
	// 12345
	InstanceId *int64 `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the node that triggers the events. You can configure either this parameter or the InstanceId parameter.
	//
	// example:
	//
	// 1234
	NodeId *int64 `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The ID of the Alibaba Cloud account used by the owner of the events.
	//
	// example:
	//
	// 952795****
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The number of the page to return. Valid values: 1 to 30. Default value: 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Default value: 10. Maximum value: 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The status of the events. Valid values: IGNORE, NEW, FIXING, and RECOVER. The value IGNORE indicates that the events are ignored. The value NEW indicates that the events are new events. The value FIXING indicates that the events are being processed. The value RECOVER indicates that the events are processed. You can specify multiple states. Separate them with commas (,).
	//
	// example:
	//
	// IGNORE,NEW,FIXING,RECOVER
	TopicStatuses *string `json:"TopicStatuses,omitempty" xml:"TopicStatuses,omitempty"`
	// The types of the events. Valid values: SLOW and ERROR. The value SLOW indicates that the running duration of the node in the current scheduling cycle is significantly longer than the average running duration of the node in previous scheduling cycles. The value ERROR indicates that the node fails to run. You can specify multiple types. Separate them with commas (,).
	//
	// example:
	//
	// ERROR,SLOW
	TopicTypes *string `json:"TopicTypes,omitempty" xml:"TopicTypes,omitempty"`
}

func (s ListTopicsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTopicsRequest) GoString() string {
	return s.String()
}

func (s *ListTopicsRequest) GetBeginTime() *string {
	return s.BeginTime
}

func (s *ListTopicsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ListTopicsRequest) GetInstanceId() *int64 {
	return s.InstanceId
}

func (s *ListTopicsRequest) GetNodeId() *int64 {
	return s.NodeId
}

func (s *ListTopicsRequest) GetOwner() *string {
	return s.Owner
}

func (s *ListTopicsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListTopicsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTopicsRequest) GetTopicStatuses() *string {
	return s.TopicStatuses
}

func (s *ListTopicsRequest) GetTopicTypes() *string {
	return s.TopicTypes
}

func (s *ListTopicsRequest) SetBeginTime(v string) *ListTopicsRequest {
	s.BeginTime = &v
	return s
}

func (s *ListTopicsRequest) SetEndTime(v string) *ListTopicsRequest {
	s.EndTime = &v
	return s
}

func (s *ListTopicsRequest) SetInstanceId(v int64) *ListTopicsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListTopicsRequest) SetNodeId(v int64) *ListTopicsRequest {
	s.NodeId = &v
	return s
}

func (s *ListTopicsRequest) SetOwner(v string) *ListTopicsRequest {
	s.Owner = &v
	return s
}

func (s *ListTopicsRequest) SetPageNumber(v int32) *ListTopicsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTopicsRequest) SetPageSize(v int32) *ListTopicsRequest {
	s.PageSize = &v
	return s
}

func (s *ListTopicsRequest) SetTopicStatuses(v string) *ListTopicsRequest {
	s.TopicStatuses = &v
	return s
}

func (s *ListTopicsRequest) SetTopicTypes(v string) *ListTopicsRequest {
	s.TopicTypes = &v
	return s
}

func (s *ListTopicsRequest) Validate() error {
	return dara.Validate(s)
}
