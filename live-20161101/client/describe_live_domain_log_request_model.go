// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveDomainLogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeLiveDomainLogRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeLiveDomainLogRequest
	GetEndTime() *string
	SetOwnerId(v int64) *DescribeLiveDomainLogRequest
	GetOwnerId() *int64
	SetPageNumber(v int64) *DescribeLiveDomainLogRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeLiveDomainLogRequest
	GetPageSize() *int64
	SetRegionId(v string) *DescribeLiveDomainLogRequest
	GetRegionId() *string
	SetStartTime(v string) *DescribeLiveDomainLogRequest
	GetStartTime() *string
}

type DescribeLiveDomainLogRequest struct {
	// The streaming domain or ingest domain.
	//
	// You can specify only one domain name.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the time range to query data. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC.
	//
	// The end time must be later than the start time. The maximum time range that can be specified is 31 days.
	//
	// example:
	//
	// 2016-10-20T05:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number.
	//
	// >  If you do not specify the PageNumber parameter, the data on the first page is returned.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// 	- Valid values: integers from **1 to 1000**.
	//
	// 	- Default value: **300**.
	//
	// 	- Maximum value: **1000**.
	//
	// example:
	//
	// 20
	PageSize *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC.
	//
	// example:
	//
	// 2016-10-20T04:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeLiveDomainLogRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainLogRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainLogRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveDomainLogRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeLiveDomainLogRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLiveDomainLogRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeLiveDomainLogRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeLiveDomainLogRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLiveDomainLogRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeLiveDomainLogRequest) SetDomainName(v string) *DescribeLiveDomainLogRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveDomainLogRequest) SetEndTime(v string) *DescribeLiveDomainLogRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeLiveDomainLogRequest) SetOwnerId(v int64) *DescribeLiveDomainLogRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLiveDomainLogRequest) SetPageNumber(v int64) *DescribeLiveDomainLogRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeLiveDomainLogRequest) SetPageSize(v int64) *DescribeLiveDomainLogRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeLiveDomainLogRequest) SetRegionId(v string) *DescribeLiveDomainLogRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLiveDomainLogRequest) SetStartTime(v string) *DescribeLiveDomainLogRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeLiveDomainLogRequest) Validate() error {
	return dara.Validate(s)
}
