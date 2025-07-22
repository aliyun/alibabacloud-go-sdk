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
	// example:
	//
	// aoe****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// testChannel
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// cn
	ServiceArea *string `json:"ServiceArea,omitempty" xml:"ServiceArea,omitempty"`
	// example:
	//
	// desc
	SortType *string `json:"SortType,omitempty" xml:"SortType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2018-01-29T00:00:00Z
	TimePoint *string `json:"TimePoint,omitempty" xml:"TimePoint,omitempty"`
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
