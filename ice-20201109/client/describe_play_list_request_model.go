// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePlayListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBeginTs(v string) *DescribePlayListRequest
	GetBeginTs() *string
	SetEndTs(v string) *DescribePlayListRequest
	GetEndTs() *string
	SetOrderName(v string) *DescribePlayListRequest
	GetOrderName() *string
	SetOrderType(v string) *DescribePlayListRequest
	GetOrderType() *string
	SetPageNo(v int32) *DescribePlayListRequest
	GetPageNo() *int32
	SetPageSize(v int32) *DescribePlayListRequest
	GetPageSize() *int32
	SetPlayType(v string) *DescribePlayListRequest
	GetPlayType() *string
	SetStatus(v string) *DescribePlayListRequest
	GetStatus() *string
	SetTraceId(v string) *DescribePlayListRequest
	GetTraceId() *string
}

type DescribePlayListRequest struct {
	// The beginning of the time range to query. By default, the system queries data of the current day.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1676170500011
	BeginTs *string `json:"BeginTs,omitempty" xml:"BeginTs,omitempty"`
	// The end of the time range to query. The time range cannot exceed 24 hours.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1682474405173
	EndTs *string `json:"EndTs,omitempty" xml:"EndTs,omitempty"`
	// The criteria by which the sorting is performed. Valid values:
	//
	// - FirstFrameDuration
	//
	// - PlayDuration
	//
	// - VideoDuration
	//
	// - StuckDuration
	//
	// example:
	//
	// FirstFrameDuration
	OrderName *string `json:"OrderName,omitempty" xml:"OrderName,omitempty"`
	// The sort order. Valid values:
	//
	// - DESC: descending order.
	//
	// - ASC: ascending order.
	//
	// example:
	//
	// DESC
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// The page number. Default value: 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The playback type. Valid value:
	//
	// - vod
	//
	// example:
	//
	// vod
	PlayType *string `json:"PlayType,omitempty" xml:"PlayType,omitempty"`
	// The playback status. Valid values:
	//
	// - complete
	//
	// - playing
	//
	// - unusual: A playback error occurs.
	//
	// example:
	//
	// complete
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The TraceId of the player.
	//
	// example:
	//
	// 0bc5e70516766285805381012d271e
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s DescribePlayListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePlayListRequest) GoString() string {
	return s.String()
}

func (s *DescribePlayListRequest) GetBeginTs() *string {
	return s.BeginTs
}

func (s *DescribePlayListRequest) GetEndTs() *string {
	return s.EndTs
}

func (s *DescribePlayListRequest) GetOrderName() *string {
	return s.OrderName
}

func (s *DescribePlayListRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *DescribePlayListRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *DescribePlayListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribePlayListRequest) GetPlayType() *string {
	return s.PlayType
}

func (s *DescribePlayListRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribePlayListRequest) GetTraceId() *string {
	return s.TraceId
}

func (s *DescribePlayListRequest) SetBeginTs(v string) *DescribePlayListRequest {
	s.BeginTs = &v
	return s
}

func (s *DescribePlayListRequest) SetEndTs(v string) *DescribePlayListRequest {
	s.EndTs = &v
	return s
}

func (s *DescribePlayListRequest) SetOrderName(v string) *DescribePlayListRequest {
	s.OrderName = &v
	return s
}

func (s *DescribePlayListRequest) SetOrderType(v string) *DescribePlayListRequest {
	s.OrderType = &v
	return s
}

func (s *DescribePlayListRequest) SetPageNo(v int32) *DescribePlayListRequest {
	s.PageNo = &v
	return s
}

func (s *DescribePlayListRequest) SetPageSize(v int32) *DescribePlayListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribePlayListRequest) SetPlayType(v string) *DescribePlayListRequest {
	s.PlayType = &v
	return s
}

func (s *DescribePlayListRequest) SetStatus(v string) *DescribePlayListRequest {
	s.Status = &v
	return s
}

func (s *DescribePlayListRequest) SetTraceId(v string) *DescribePlayListRequest {
	s.TraceId = &v
	return s
}

func (s *DescribePlayListRequest) Validate() error {
	return dara.Validate(s)
}
