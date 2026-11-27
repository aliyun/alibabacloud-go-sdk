// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRtcChannelListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeRtcChannelListRequest
	GetAppId() *string
	SetChannelId(v string) *DescribeRtcChannelListRequest
	GetChannelId() *string
	SetOwnerId(v int64) *DescribeRtcChannelListRequest
	GetOwnerId() *int64
	SetPageNo(v int64) *DescribeRtcChannelListRequest
	GetPageNo() *int64
	SetPageSize(v int64) *DescribeRtcChannelListRequest
	GetPageSize() *int64
	SetServiceArea(v string) *DescribeRtcChannelListRequest
	GetServiceArea() *string
	SetSortType(v string) *DescribeRtcChannelListRequest
	GetSortType() *string
	SetTimePoint(v string) *DescribeRtcChannelListRequest
	GetTimePoint() *string
	SetUserId(v string) *DescribeRtcChannelListRequest
	GetUserId() *string
}

type DescribeRtcChannelListRequest struct {
	// The application ID. You can create and query the ID in the console. Only a single ID is supported.
	//
	// example:
	//
	// aoe****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The channel ID. Only a single ID is supported.
	//
	// example:
	//
	// testChannel
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Valid values: greater than **0**. Default value: 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. Valid values: greater than **0**. Default value: 10.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The service area. By default, all areas are queried. Valid values:
	//
	// - **cn**: China.
	//
	// - **us**: United States.
	//
	// example:
	//
	// cn
	ServiceArea *string `json:"ServiceArea,omitempty" xml:"ServiceArea,omitempty"`
	// The sort order. By default, the records are sorted in descending order of StartTime.
	//
	// example:
	//
	// desc
	SortType *string `json:"SortType,omitempty" xml:"SortType,omitempty"`
	// The query time in UTC. The format is yyyy-MM-ddTHH:mm:ssZ. The time granularity is one day.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2018-01-29T00:00:00Z
	TimePoint *string `json:"TimePoint,omitempty" xml:"TimePoint,omitempty"`
	// The user ID. Only a single ID is supported.
	//
	// example:
	//
	// testUser
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DescribeRtcChannelListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRtcChannelListRequest) GoString() string {
	return s.String()
}

func (s *DescribeRtcChannelListRequest) GetAppId() *string {
	return s.AppId
}

func (s *DescribeRtcChannelListRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *DescribeRtcChannelListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeRtcChannelListRequest) GetPageNo() *int64 {
	return s.PageNo
}

func (s *DescribeRtcChannelListRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeRtcChannelListRequest) GetServiceArea() *string {
	return s.ServiceArea
}

func (s *DescribeRtcChannelListRequest) GetSortType() *string {
	return s.SortType
}

func (s *DescribeRtcChannelListRequest) GetTimePoint() *string {
	return s.TimePoint
}

func (s *DescribeRtcChannelListRequest) GetUserId() *string {
	return s.UserId
}

func (s *DescribeRtcChannelListRequest) SetAppId(v string) *DescribeRtcChannelListRequest {
	s.AppId = &v
	return s
}

func (s *DescribeRtcChannelListRequest) SetChannelId(v string) *DescribeRtcChannelListRequest {
	s.ChannelId = &v
	return s
}

func (s *DescribeRtcChannelListRequest) SetOwnerId(v int64) *DescribeRtcChannelListRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeRtcChannelListRequest) SetPageNo(v int64) *DescribeRtcChannelListRequest {
	s.PageNo = &v
	return s
}

func (s *DescribeRtcChannelListRequest) SetPageSize(v int64) *DescribeRtcChannelListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeRtcChannelListRequest) SetServiceArea(v string) *DescribeRtcChannelListRequest {
	s.ServiceArea = &v
	return s
}

func (s *DescribeRtcChannelListRequest) SetSortType(v string) *DescribeRtcChannelListRequest {
	s.SortType = &v
	return s
}

func (s *DescribeRtcChannelListRequest) SetTimePoint(v string) *DescribeRtcChannelListRequest {
	s.TimePoint = &v
	return s
}

func (s *DescribeRtcChannelListRequest) SetUserId(v string) *DescribeRtcChannelListRequest {
	s.UserId = &v
	return s
}

func (s *DescribeRtcChannelListRequest) Validate() error {
	return dara.Validate(s)
}
