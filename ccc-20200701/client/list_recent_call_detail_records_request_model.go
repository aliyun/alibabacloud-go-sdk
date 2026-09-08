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
	// Fuzzy search criteria in JSON object format. The object contains three properties that can be combined arbitrarily: phoneNumber (fuzzy search by calling or called number), callingNumber (fuzzy search by calling number), and calledNumber (fuzzy search by called number).
	//
	// example:
	//
	// {"phoneNumber":"1312121****","callingNumber":"1312121****","calledNumber":"1312121****"}
	Criteria *string `json:"Criteria,omitempty" xml:"Criteria,omitempty"`
	// End UNIX timestamp. The default value is the current time.
	//
	// example:
	//
	// 1604639129000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Page ordinal number, ranging from 1 to 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Page size, ranging from 1 to 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Start UNIX timestamp. The default value is the start time of the current day. The earliest allowed time is 180 days before the current time.
	//
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
