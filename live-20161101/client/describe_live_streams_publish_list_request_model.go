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
	// The name of the application to which the stream belongs. You can view AppName on the [Stream Management](https://help.aliyun.com/document_detail/197397.html) page.
	//
	// example:
	//
	// liveApp****
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The ingest domain or streamer streaming domain.
	//
	// > - When you specify DomainName, make sure that the domain name is a live streaming domain name and that the user calling this operation has the permissions to operate on the specified domain name.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end time. The interval between EndTime and StartTime cannot exceed 30 days.
	//
	// Format: <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z (UTC).
	//
	// This parameter is required.
	//
	// example:
	//
	// 2017-12-22T08:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The sorting method. Valid values:
	//
	// - **stream_name_desc**: sorts by live stream name in descending order.
	//
	// - **stream_name_asc**: sorts by live stream name in ascending order.
	//
	// - **publish_time_desc**: sorts by stream ingest time in descending order.
	//
	// - **publish_time_asc*	- (default): sorts by stream ingest time in ascending order.
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
	// The page size. Valid values: **1 to 3000**. Default value: **2000**.
	//
	// example:
	//
	// 1500
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Specifies whether to use fuzzy match for the stream name. Valid values:
	//
	// - **fuzzy*	- (default): fuzzy match.
	//
	// - **strict**: exact match.
	//
	// example:
	//
	// fuzzy
	QueryType *string `json:"QueryType,omitempty" xml:"QueryType,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The start time of stream ingest.
	//
	// Format: <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z (UTC).
	//
	// This parameter is required.
	//
	// example:
	//
	// 2017-12-21T08:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The stream name. You can view StreamName on the [Stream Management](https://help.aliyun.com/document_detail/197397.html) page.
	//
	// example:
	//
	// liveStream****
	StreamName *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
	// The stream type. Valid values:
	//
	// - **Not specified**: queries raw streams.
	//
	// - **all**: queries all streams.
	//
	// - **trans**: queries transcoded streams.
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
