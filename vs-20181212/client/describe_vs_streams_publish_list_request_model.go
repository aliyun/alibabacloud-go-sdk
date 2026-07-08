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
	// The application name of the live stream.
	//
	// example:
	//
	// xxxApp
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// Your domain name.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.aliyundoc.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end time.
	//
	// > - Use UTC format. Example: 2016-06-30T19:00:00Z
	//
	// >
	//
	// > - The interval between EndTime and StartTime must not exceed 30 days.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2016-06-30T19:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The sort order. Valid values:
	//
	// - stream_name_desc (sort by stream name in descending order)
	//
	// - stream_name_asc (sort by stream name in ascending order)
	//
	// - publish_time_desc (sort by publish time in descending order)
	//
	// - publish_time_asc (sort by publish time in ascending order) (default)
	//
	// example:
	//
	// publish_time_asc
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 3000.<br>
	//
	// Valid values: 1 to 3000.<br>
	//
	// example:
	//
	// 3000
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Specifies whether to use fuzzy matching for the stream name. Valid values:
	//
	// - fuzzy (fuzzy match)
	//
	// - strict (exact match)
	//
	// example:
	//
	// fuzzy
	QueryType *string `json:"QueryType,omitempty" xml:"QueryType,omitempty"`
	// The start time.
	//
	// > Use UTC format. Example: 2016-06-29T19:00:00Z
	//
	// This parameter is required.
	//
	// example:
	//
	// 2016-06-29T19:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The live stream name.
	//
	// example:
	//
	// xxxStream
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
	// The stream type. Valid values:
	//
	// - all (all streams) (default)
	//
	// - raw (raw stream)
	//
	// - trans (transcoded stream)
	//
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
