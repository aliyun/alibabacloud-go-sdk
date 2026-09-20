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
	// The start time for discovery. Specify the time in UTC format (yyyy-MM-dd\\"T\\"HH:mm:ssZ).
	//
	// This parameter is required.
	//
	// example:
	//
	// 2019-03-24T00:00:00+0800
	BeginTime *string `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// The end time for discovery. Specify the time in UTC format (yyyy-MM-dd\\"T\\"HH:mm:ssZ).
	//
	// This parameter is required.
	//
	// example:
	//
	// 2019-03-25T23:00:00+0800
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The instance ID associated with the event. This parameter is mutually exclusive with NodeId.
	//
	// example:
	//
	// 12345
	InstanceId *int64 `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the node associated with the event. This parameter is mutually exclusive with InstanceId.
	//
	// example:
	//
	// 1234
	NodeId *int64 `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The Alibaba Cloud UID of the event owner.
	//
	// example:
	//
	// 952795****
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The page number. Default value: 1. Minimum value: 1. Maximum value: 30.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 10. Maximum value: 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The status of the event. Valid values: IGNORE (ignored), NEW (newly discovered), FIXING (being processed), and RECOVER (recovered). Separate multiple event statuses with commas (,).
	//
	// example:
	//
	// IGNORE,NEW,FIXING,RECOVER
	TopicStatuses *string `json:"TopicStatuses,omitempty" xml:"TopicStatuses,omitempty"`
	// The type of the event. Valid values: SLOW (slow) and ERROR (error). Separate multiple event types with commas (,).
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
