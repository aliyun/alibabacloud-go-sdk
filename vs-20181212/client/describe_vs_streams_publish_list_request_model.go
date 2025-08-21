// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVsStreamsPublishListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *DescribeVsStreamsPublishListRequest
	GetAppName() *string
	SetDomainName(v string) *DescribeVsStreamsPublishListRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeVsStreamsPublishListRequest
	GetEndTime() *string
	SetOrderBy(v string) *DescribeVsStreamsPublishListRequest
	GetOrderBy() *string
	SetOwnerId(v int64) *DescribeVsStreamsPublishListRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeVsStreamsPublishListRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeVsStreamsPublishListRequest
	GetPageSize() *int32
	SetQueryType(v string) *DescribeVsStreamsPublishListRequest
	GetQueryType() *string
	SetStartTime(v string) *DescribeVsStreamsPublishListRequest
	GetStartTime() *string
	SetStreamName(v string) *DescribeVsStreamsPublishListRequest
	GetStreamName() *string
	SetStreamType(v string) *DescribeVsStreamsPublishListRequest
	GetStreamType() *string
}

type DescribeVsStreamsPublishListRequest struct {
	// example:
	//
	// xxxApp
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// example.aliyundoc.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// This parameter is required.
	//
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
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 3000
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// fuzzy
	QueryType *string `json:"QueryType,omitempty" xml:"QueryType,omitempty"`
	// This parameter is required.
	//
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

func (s DescribeVsStreamsPublishListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVsStreamsPublishListRequest) GoString() string {
	return s.String()
}

func (s *DescribeVsStreamsPublishListRequest) GetAppName() *string {
	return s.AppName
}

func (s *DescribeVsStreamsPublishListRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeVsStreamsPublishListRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeVsStreamsPublishListRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *DescribeVsStreamsPublishListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVsStreamsPublishListRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeVsStreamsPublishListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeVsStreamsPublishListRequest) GetQueryType() *string {
	return s.QueryType
}

func (s *DescribeVsStreamsPublishListRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeVsStreamsPublishListRequest) GetStreamName() *string {
	return s.StreamName
}

func (s *DescribeVsStreamsPublishListRequest) GetStreamType() *string {
	return s.StreamType
}

func (s *DescribeVsStreamsPublishListRequest) SetAppName(v string) *DescribeVsStreamsPublishListRequest {
	s.AppName = &v
	return s
}

func (s *DescribeVsStreamsPublishListRequest) SetDomainName(v string) *DescribeVsStreamsPublishListRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeVsStreamsPublishListRequest) SetEndTime(v string) *DescribeVsStreamsPublishListRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeVsStreamsPublishListRequest) SetOrderBy(v string) *DescribeVsStreamsPublishListRequest {
	s.OrderBy = &v
	return s
}

func (s *DescribeVsStreamsPublishListRequest) SetOwnerId(v int64) *DescribeVsStreamsPublishListRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVsStreamsPublishListRequest) SetPageNumber(v int32) *DescribeVsStreamsPublishListRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeVsStreamsPublishListRequest) SetPageSize(v int32) *DescribeVsStreamsPublishListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeVsStreamsPublishListRequest) SetQueryType(v string) *DescribeVsStreamsPublishListRequest {
	s.QueryType = &v
	return s
}

func (s *DescribeVsStreamsPublishListRequest) SetStartTime(v string) *DescribeVsStreamsPublishListRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeVsStreamsPublishListRequest) SetStreamName(v string) *DescribeVsStreamsPublishListRequest {
	s.StreamName = &v
	return s
}

func (s *DescribeVsStreamsPublishListRequest) SetStreamType(v string) *DescribeVsStreamsPublishListRequest {
	s.StreamType = &v
	return s
}

func (s *DescribeVsStreamsPublishListRequest) Validate() error {
	return dara.Validate(s)
}
