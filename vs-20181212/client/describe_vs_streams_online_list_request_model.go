// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVsStreamsOnlineListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *DescribeVsStreamsOnlineListRequest
	GetAppName() *string
	SetDomainName(v string) *DescribeVsStreamsOnlineListRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeVsStreamsOnlineListRequest
	GetEndTime() *string
	SetOrderBy(v string) *DescribeVsStreamsOnlineListRequest
	GetOrderBy() *string
	SetOwnerId(v int64) *DescribeVsStreamsOnlineListRequest
	GetOwnerId() *int64
	SetPageNum(v int32) *DescribeVsStreamsOnlineListRequest
	GetPageNum() *int32
	SetPageSize(v int32) *DescribeVsStreamsOnlineListRequest
	GetPageSize() *int32
	SetQueryType(v string) *DescribeVsStreamsOnlineListRequest
	GetQueryType() *string
	SetStartTime(v string) *DescribeVsStreamsOnlineListRequest
	GetStartTime() *string
	SetStreamName(v string) *DescribeVsStreamsOnlineListRequest
	GetStreamName() *string
	SetStreamType(v string) *DescribeVsStreamsOnlineListRequest
	GetStreamType() *string
}

type DescribeVsStreamsOnlineListRequest struct {
	// example:
	//
	// xxxApp
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// example:
	//
	// 2016-06-30T19:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// publish_time_asc
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// fuzzy
	QueryType *string `json:"QueryType,omitempty" xml:"QueryType,omitempty"`
	// example:
	//
	// 2016-06-29T19:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// xxxStream
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
	// example:
	//
	// all
	StreamType *string `json:"StreamType,omitempty" xml:"StreamType,omitempty"`
}

func (s DescribeVsStreamsOnlineListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVsStreamsOnlineListRequest) GoString() string {
	return s.String()
}

func (s *DescribeVsStreamsOnlineListRequest) GetAppName() *string {
	return s.AppName
}

func (s *DescribeVsStreamsOnlineListRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeVsStreamsOnlineListRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeVsStreamsOnlineListRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *DescribeVsStreamsOnlineListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVsStreamsOnlineListRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *DescribeVsStreamsOnlineListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeVsStreamsOnlineListRequest) GetQueryType() *string {
	return s.QueryType
}

func (s *DescribeVsStreamsOnlineListRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeVsStreamsOnlineListRequest) GetStreamName() *string {
	return s.StreamName
}

func (s *DescribeVsStreamsOnlineListRequest) GetStreamType() *string {
	return s.StreamType
}

func (s *DescribeVsStreamsOnlineListRequest) SetAppName(v string) *DescribeVsStreamsOnlineListRequest {
	s.AppName = &v
	return s
}

func (s *DescribeVsStreamsOnlineListRequest) SetDomainName(v string) *DescribeVsStreamsOnlineListRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeVsStreamsOnlineListRequest) SetEndTime(v string) *DescribeVsStreamsOnlineListRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeVsStreamsOnlineListRequest) SetOrderBy(v string) *DescribeVsStreamsOnlineListRequest {
	s.OrderBy = &v
	return s
}

func (s *DescribeVsStreamsOnlineListRequest) SetOwnerId(v int64) *DescribeVsStreamsOnlineListRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVsStreamsOnlineListRequest) SetPageNum(v int32) *DescribeVsStreamsOnlineListRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeVsStreamsOnlineListRequest) SetPageSize(v int32) *DescribeVsStreamsOnlineListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeVsStreamsOnlineListRequest) SetQueryType(v string) *DescribeVsStreamsOnlineListRequest {
	s.QueryType = &v
	return s
}

func (s *DescribeVsStreamsOnlineListRequest) SetStartTime(v string) *DescribeVsStreamsOnlineListRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeVsStreamsOnlineListRequest) SetStreamName(v string) *DescribeVsStreamsOnlineListRequest {
	s.StreamName = &v
	return s
}

func (s *DescribeVsStreamsOnlineListRequest) SetStreamType(v string) *DescribeVsStreamsOnlineListRequest {
	s.StreamType = &v
	return s
}

func (s *DescribeVsStreamsOnlineListRequest) Validate() error {
	return dara.Validate(s)
}
