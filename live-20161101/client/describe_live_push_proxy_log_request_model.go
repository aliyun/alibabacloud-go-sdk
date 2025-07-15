// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLivePushProxyLogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeLivePushProxyLogRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeLivePushProxyLogRequest
	GetEndTime() *string
	SetOwnerId(v int64) *DescribeLivePushProxyLogRequest
	GetOwnerId() *int64
	SetPageNumber(v int64) *DescribeLivePushProxyLogRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeLivePushProxyLogRequest
	GetPageSize() *int64
	SetRegionId(v string) *DescribeLivePushProxyLogRequest
	GetRegionId() *string
	SetStartTime(v string) *DescribeLivePushProxyLogRequest
	GetStartTime() *string
}

type DescribeLivePushProxyLogRequest struct {
	// The ingest domain. You can specify only one domain in each request.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// Get the log end time.
	//
	// Date format follows the ISO8601 representation and uses UTC+0 time, formatted as yyyy-MM-ddTHH:mm:ssZ.
	//
	// example:
	//
	// 2023-09-20T09:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 2
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: [1,1000]. Default value: 300.
	//
	// example:
	//
	// 300
	PageSize *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Get the log start time in ISO8601 format with UTC+0 timezone, formatted as yyyy-MM-ddTHH:mm:ssZ.
	//
	// example:
	//
	// 2023-09-20T08:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeLivePushProxyLogRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLivePushProxyLogRequest) GoString() string {
	return s.String()
}

func (s *DescribeLivePushProxyLogRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLivePushProxyLogRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeLivePushProxyLogRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLivePushProxyLogRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeLivePushProxyLogRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeLivePushProxyLogRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLivePushProxyLogRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeLivePushProxyLogRequest) SetDomainName(v string) *DescribeLivePushProxyLogRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLivePushProxyLogRequest) SetEndTime(v string) *DescribeLivePushProxyLogRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeLivePushProxyLogRequest) SetOwnerId(v int64) *DescribeLivePushProxyLogRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLivePushProxyLogRequest) SetPageNumber(v int64) *DescribeLivePushProxyLogRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeLivePushProxyLogRequest) SetPageSize(v int64) *DescribeLivePushProxyLogRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeLivePushProxyLogRequest) SetRegionId(v string) *DescribeLivePushProxyLogRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLivePushProxyLogRequest) SetStartTime(v string) *DescribeLivePushProxyLogRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeLivePushProxyLogRequest) Validate() error {
	return dara.Validate(s)
}
