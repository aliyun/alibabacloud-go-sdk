// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRecentCallDetailRecordsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCriteria(v string) *ListRecentCallDetailRecordsRequest
	GetCriteria() *string
	SetEndTime(v int64) *ListRecentCallDetailRecordsRequest
	GetEndTime() *int64
	SetInstanceId(v string) *ListRecentCallDetailRecordsRequest
	GetInstanceId() *string
	SetPageNumber(v int32) *ListRecentCallDetailRecordsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListRecentCallDetailRecordsRequest
	GetPageSize() *int32
	SetStartTime(v int64) *ListRecentCallDetailRecordsRequest
	GetStartTime() *int64
}

type ListRecentCallDetailRecordsRequest struct {
	// example:
	//
	// {"phoneNumber":"1312121****","callingNumber":"1312121****","calledNumber":"1312121****"}
	Criteria *string `json:"Criteria,omitempty" xml:"Criteria,omitempty"`
	// example:
	//
	// 1604639129000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 1604638129000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListRecentCallDetailRecordsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRecentCallDetailRecordsRequest) GoString() string {
	return s.String()
}

func (s *ListRecentCallDetailRecordsRequest) GetCriteria() *string {
	return s.Criteria
}

func (s *ListRecentCallDetailRecordsRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListRecentCallDetailRecordsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListRecentCallDetailRecordsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListRecentCallDetailRecordsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListRecentCallDetailRecordsRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListRecentCallDetailRecordsRequest) SetCriteria(v string) *ListRecentCallDetailRecordsRequest {
	s.Criteria = &v
	return s
}

func (s *ListRecentCallDetailRecordsRequest) SetEndTime(v int64) *ListRecentCallDetailRecordsRequest {
	s.EndTime = &v
	return s
}

func (s *ListRecentCallDetailRecordsRequest) SetInstanceId(v string) *ListRecentCallDetailRecordsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListRecentCallDetailRecordsRequest) SetPageNumber(v int32) *ListRecentCallDetailRecordsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListRecentCallDetailRecordsRequest) SetPageSize(v int32) *ListRecentCallDetailRecordsRequest {
	s.PageSize = &v
	return s
}

func (s *ListRecentCallDetailRecordsRequest) SetStartTime(v int64) *ListRecentCallDetailRecordsRequest {
	s.StartTime = &v
	return s
}

func (s *ListRecentCallDetailRecordsRequest) Validate() error {
	return dara.Validate(s)
}
