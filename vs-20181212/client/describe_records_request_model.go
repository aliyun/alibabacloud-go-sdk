// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRecordsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeRecordsRequest
	GetEndTime() *string
	SetOwnerId(v int64) *DescribeRecordsRequest
	GetOwnerId() *int64
	SetPageNum(v int64) *DescribeRecordsRequest
	GetPageNum() *int64
	SetPageSize(v int64) *DescribeRecordsRequest
	GetPageSize() *int64
	SetPrivateBucket(v bool) *DescribeRecordsRequest
	GetPrivateBucket() *bool
	SetSortBy(v string) *DescribeRecordsRequest
	GetSortBy() *string
	SetSortDirection(v string) *DescribeRecordsRequest
	GetSortDirection() *string
	SetStartTime(v string) *DescribeRecordsRequest
	GetStartTime() *string
	SetStreamId(v string) *DescribeRecordsRequest
	GetStreamId() *string
	SetType(v string) *DescribeRecordsRequest
	GetType() *string
}

type DescribeRecordsRequest struct {
	// The end time. Specify the time in the ISO 8601 standard in the YYYY-MM-DDTHH:mm:ssZ format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2021-11-24T00:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNum *int64 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries per page. Default value: 20.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Specifies whether to use a private bucket.
	//
	// example:
	//
	// true
	PrivateBucket *bool `json:"PrivateBucket,omitempty" xml:"PrivateBucket,omitempty"`
	// The field by which to sort. Valid values:
	//
	// - Id: default
	//
	// example:
	//
	// Id
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The sort order. Valid values:
	//
	// - asc: ascending order (default)
	//
	// - desc: descending order
	//
	// example:
	//
	// asc
	SortDirection *string `json:"SortDirection,omitempty" xml:"SortDirection,omitempty"`
	// The start time. Specify the time in the ISO 8601 standard in the YYYY-MM-DDTHH:mm:ssZ format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2021-11-22T00:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The ID of the stream.
	//
	// This parameter is required.
	//
	// example:
	//
	// 323*****997-cn-qingdao
	StreamId *string `json:"StreamId,omitempty" xml:"StreamId,omitempty"`
	// The type of the storage record. Valid values:
	//
	// - record: recording
	//
	// - snapshot: snapshot
	//
	// This parameter is required.
	//
	// example:
	//
	// record
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeRecordsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRecordsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRecordsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeRecordsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeRecordsRequest) GetPageNum() *int64 {
	return s.PageNum
}

func (s *DescribeRecordsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeRecordsRequest) GetPrivateBucket() *bool {
	return s.PrivateBucket
}

func (s *DescribeRecordsRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *DescribeRecordsRequest) GetSortDirection() *string {
	return s.SortDirection
}

func (s *DescribeRecordsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeRecordsRequest) GetStreamId() *string {
	return s.StreamId
}

func (s *DescribeRecordsRequest) GetType() *string {
	return s.Type
}

func (s *DescribeRecordsRequest) SetEndTime(v string) *DescribeRecordsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeRecordsRequest) SetOwnerId(v int64) *DescribeRecordsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeRecordsRequest) SetPageNum(v int64) *DescribeRecordsRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeRecordsRequest) SetPageSize(v int64) *DescribeRecordsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeRecordsRequest) SetPrivateBucket(v bool) *DescribeRecordsRequest {
	s.PrivateBucket = &v
	return s
}

func (s *DescribeRecordsRequest) SetSortBy(v string) *DescribeRecordsRequest {
	s.SortBy = &v
	return s
}

func (s *DescribeRecordsRequest) SetSortDirection(v string) *DescribeRecordsRequest {
	s.SortDirection = &v
	return s
}

func (s *DescribeRecordsRequest) SetStartTime(v string) *DescribeRecordsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeRecordsRequest) SetStreamId(v string) *DescribeRecordsRequest {
	s.StreamId = &v
	return s
}

func (s *DescribeRecordsRequest) SetType(v string) *DescribeRecordsRequest {
	s.Type = &v
	return s
}

func (s *DescribeRecordsRequest) Validate() error {
	return dara.Validate(s)
}
