// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListActionRecordsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActionNames(v string) *ListActionRecordsRequest
	GetActionNames() *string
	SetEndTime(v int64) *ListActionRecordsRequest
	GetEndTime() *int64
	SetFilter(v string) *ListActionRecordsRequest
	GetFilter() *string
	SetPage(v int32) *ListActionRecordsRequest
	GetPage() *int32
	SetRequestId(v string) *ListActionRecordsRequest
	GetRequestId() *string
	SetSize(v int32) *ListActionRecordsRequest
	GetSize() *int32
	SetStartTime(v int64) *ListActionRecordsRequest
	GetStartTime() *int64
	SetUserId(v string) *ListActionRecordsRequest
	GetUserId() *string
}

type ListActionRecordsRequest struct {
	// Filters records by the name of the change operation. Separate multiple names with commas.
	//
	// example:
	//
	// RestartInstance
	ActionNames *string `json:"actionNames,omitempty" xml:"actionNames,omitempty"`
	// The upper bound of the query range for the change start time. The value is a UNIX timestamp in milliseconds. If you do not specify this parameter, the current time is used.
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// The display type of change records. Default value: ACTION_RECORD.
	//
	// example:
	//
	// ACTION_RECORD
	Filter *string `json:"filter,omitempty" xml:"filter,omitempty"`
	// The page number of the returned results.
	//
	// example:
	//
	// 1
	Page *int32 `json:"page,omitempty" xml:"page,omitempty"`
	// Filters records by the ID of the request that triggered the change.
	//
	// example:
	//
	// request-123
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The number of records returned per page.
	//
	// example:
	//
	// 10
	Size *int32 `json:"size,omitempty" xml:"size,omitempty"`
	// The lower bound of the query range for the change start time. The value is a UNIX timestamp in milliseconds.
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// Filters records by the ID of the user that triggered the change.
	//
	// example:
	//
	// user-123
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s ListActionRecordsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListActionRecordsRequest) GoString() string {
	return s.String()
}

func (s *ListActionRecordsRequest) GetActionNames() *string {
	return s.ActionNames
}

func (s *ListActionRecordsRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListActionRecordsRequest) GetFilter() *string {
	return s.Filter
}

func (s *ListActionRecordsRequest) GetPage() *int32 {
	return s.Page
}

func (s *ListActionRecordsRequest) GetRequestId() *string {
	return s.RequestId
}

func (s *ListActionRecordsRequest) GetSize() *int32 {
	return s.Size
}

func (s *ListActionRecordsRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListActionRecordsRequest) GetUserId() *string {
	return s.UserId
}

func (s *ListActionRecordsRequest) SetActionNames(v string) *ListActionRecordsRequest {
	s.ActionNames = &v
	return s
}

func (s *ListActionRecordsRequest) SetEndTime(v int64) *ListActionRecordsRequest {
	s.EndTime = &v
	return s
}

func (s *ListActionRecordsRequest) SetFilter(v string) *ListActionRecordsRequest {
	s.Filter = &v
	return s
}

func (s *ListActionRecordsRequest) SetPage(v int32) *ListActionRecordsRequest {
	s.Page = &v
	return s
}

func (s *ListActionRecordsRequest) SetRequestId(v string) *ListActionRecordsRequest {
	s.RequestId = &v
	return s
}

func (s *ListActionRecordsRequest) SetSize(v int32) *ListActionRecordsRequest {
	s.Size = &v
	return s
}

func (s *ListActionRecordsRequest) SetStartTime(v int64) *ListActionRecordsRequest {
	s.StartTime = &v
	return s
}

func (s *ListActionRecordsRequest) SetUserId(v string) *ListActionRecordsRequest {
	s.UserId = &v
	return s
}

func (s *ListActionRecordsRequest) Validate() error {
	return dara.Validate(s)
}
