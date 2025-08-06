// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodDomainLogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeVodDomainLogRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeVodDomainLogRequest
	GetEndTime() *string
	SetOwnerId(v int64) *DescribeVodDomainLogRequest
	GetOwnerId() *int64
	SetPageNumber(v int64) *DescribeVodDomainLogRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeVodDomainLogRequest
	GetPageSize() *int64
	SetStartTime(v string) *DescribeVodDomainLogRequest
	GetStartTime() *string
}

type DescribeVodDomainLogRequest struct {
	// The domain name for CDN.
	//
	// >  You can specify only one domain name in each query.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the time range to query. The end time must be later than the start time. The maximum time range that can be specified is one year. Specify the time in the ISO 8601 standard in the `yyyy-MM-ddTHH:mm:ssZ` format. The time must be in UTC.
	//
	// example:
	//
	// 2016-10-20T05:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// 	- Default value: **300**.
	//
	// 	- Valid values: **1 to 1000**.
	//
	// example:
	//
	// 300
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the `yyyy-MM-ddTHH:mm:ssZ` format. The time must be in UTC.
	//
	// example:
	//
	// 2016-10-20T04:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeVodDomainLogRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainLogRequest) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainLogRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeVodDomainLogRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeVodDomainLogRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVodDomainLogRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeVodDomainLogRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeVodDomainLogRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeVodDomainLogRequest) SetDomainName(v string) *DescribeVodDomainLogRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeVodDomainLogRequest) SetEndTime(v string) *DescribeVodDomainLogRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeVodDomainLogRequest) SetOwnerId(v int64) *DescribeVodDomainLogRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVodDomainLogRequest) SetPageNumber(v int64) *DescribeVodDomainLogRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeVodDomainLogRequest) SetPageSize(v int64) *DescribeVodDomainLogRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeVodDomainLogRequest) SetStartTime(v string) *DescribeVodDomainLogRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeVodDomainLogRequest) Validate() error {
	return dara.Validate(s)
}
