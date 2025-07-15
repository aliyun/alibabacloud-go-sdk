// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveStreamsPublishListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *DescribeLiveStreamsPublishListRequest
	GetAppName() *string
	SetDomainName(v string) *DescribeLiveStreamsPublishListRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeLiveStreamsPublishListRequest
	GetEndTime() *string
	SetOrderBy(v string) *DescribeLiveStreamsPublishListRequest
	GetOrderBy() *string
	SetOwnerId(v int64) *DescribeLiveStreamsPublishListRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeLiveStreamsPublishListRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeLiveStreamsPublishListRequest
	GetPageSize() *int32
	SetQueryType(v string) *DescribeLiveStreamsPublishListRequest
	GetQueryType() *string
	SetRegionId(v string) *DescribeLiveStreamsPublishListRequest
	GetRegionId() *string
	SetStartTime(v string) *DescribeLiveStreamsPublishListRequest
	GetStartTime() *string
	SetStreamName(v string) *DescribeLiveStreamsPublishListRequest
	GetStreamName() *string
	SetStreamType(v string) *DescribeLiveStreamsPublishListRequest
	GetStreamType() *string
}

type DescribeLiveStreamsPublishListRequest struct {
	// The name of the application to which the live stream belongs.
	//
	// example:
	//
	// liveApp****
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The ingest domain or main streaming domain.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the time range to query. The time range specified by the StartTime and EndTime parameters cannot exceed 30 days.
	//
	// Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2017-12-22T08:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The sorting method. Valid values:
	//
	// 	- **stream_name_desc**: sorts the entries in descending order by stream name.
	//
	// 	- **stream_name_asc**: sorts the entries in ascending order by stream name.
	//
	// 	- **publish_time_desc**: sorts the entries in descending order by stream ingest time.
	//
	// 	- **publish_time_asc*	- (default): sorts the entries in ascending order by stream ingest time.
	//
	// example:
	//
	// publish_time_desc
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: **1 to 3000**. Default value: **2000**.
	//
	// example:
	//
	// 1500
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The mode in which stream names are matched. Valid values:
	//
	// 	- **fuzzy*	- (default): fuzzy match
	//
	// 	- **strict**: exact match
	//
	// example:
	//
	// fuzzy
	QueryType *string `json:"QueryType,omitempty" xml:"QueryType,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The beginning of the time range to query.
	//
	// Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2017-12-21T08:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The name of the live stream.
	//
	// example:
	//
	// liveStream****
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
	// The type of the streams to query. Valid values:
	//
	// 	- An empty value****: source streams
	//
	// 	- **all**: all streams
	//
	// 	- **trans**: transcoded streams
	//
	// example:
	//
	// all
	StreamType *string `json:"StreamType,omitempty" xml:"StreamType,omitempty"`
}

func (s DescribeLiveStreamsPublishListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamsPublishListRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamsPublishListRequest) GetAppName() *string {
	return s.AppName
}

func (s *DescribeLiveStreamsPublishListRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveStreamsPublishListRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeLiveStreamsPublishListRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *DescribeLiveStreamsPublishListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLiveStreamsPublishListRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeLiveStreamsPublishListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeLiveStreamsPublishListRequest) GetQueryType() *string {
	return s.QueryType
}

func (s *DescribeLiveStreamsPublishListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLiveStreamsPublishListRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeLiveStreamsPublishListRequest) GetStreamName() *string {
	return s.StreamName
}

func (s *DescribeLiveStreamsPublishListRequest) GetStreamType() *string {
	return s.StreamType
}

func (s *DescribeLiveStreamsPublishListRequest) SetAppName(v string) *DescribeLiveStreamsPublishListRequest {
	s.AppName = &v
	return s
}

func (s *DescribeLiveStreamsPublishListRequest) SetDomainName(v string) *DescribeLiveStreamsPublishListRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveStreamsPublishListRequest) SetEndTime(v string) *DescribeLiveStreamsPublishListRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeLiveStreamsPublishListRequest) SetOrderBy(v string) *DescribeLiveStreamsPublishListRequest {
	s.OrderBy = &v
	return s
}

func (s *DescribeLiveStreamsPublishListRequest) SetOwnerId(v int64) *DescribeLiveStreamsPublishListRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLiveStreamsPublishListRequest) SetPageNumber(v int32) *DescribeLiveStreamsPublishListRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeLiveStreamsPublishListRequest) SetPageSize(v int32) *DescribeLiveStreamsPublishListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeLiveStreamsPublishListRequest) SetQueryType(v string) *DescribeLiveStreamsPublishListRequest {
	s.QueryType = &v
	return s
}

func (s *DescribeLiveStreamsPublishListRequest) SetRegionId(v string) *DescribeLiveStreamsPublishListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLiveStreamsPublishListRequest) SetStartTime(v string) *DescribeLiveStreamsPublishListRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeLiveStreamsPublishListRequest) SetStreamName(v string) *DescribeLiveStreamsPublishListRequest {
	s.StreamName = &v
	return s
}

func (s *DescribeLiveStreamsPublishListRequest) SetStreamType(v string) *DescribeLiveStreamsPublishListRequest {
	s.StreamType = &v
	return s
}

func (s *DescribeLiveStreamsPublishListRequest) Validate() error {
	return dara.Validate(s)
}
