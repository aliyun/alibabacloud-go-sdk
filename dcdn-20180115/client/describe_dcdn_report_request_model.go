// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnReportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArea(v string) *DescribeDcdnReportRequest
	GetArea() *string
	SetDomainName(v string) *DescribeDcdnReportRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeDcdnReportRequest
	GetEndTime() *string
	SetHttpCode(v string) *DescribeDcdnReportRequest
	GetHttpCode() *string
	SetIsOverseas(v string) *DescribeDcdnReportRequest
	GetIsOverseas() *string
	SetReportId(v int64) *DescribeDcdnReportRequest
	GetReportId() *int64
	SetStartTime(v string) *DescribeDcdnReportRequest
	GetStartTime() *string
}

type DescribeDcdnReportRequest struct {
	// The region. You can call the [DescribeDcdnRegionAndIsp](https://help.aliyun.com/document_detail/207199.html) operation to query regions.
	//
	// 	- If you do not specify a region, data in all regions is queried.
	//
	// 	- If you specify a region, data in the specified region is returned. You can specify one or more regions. Separate regions with commas (,).
	//
	// example:
	//
	// shanghai
	Area *string `json:"Area,omitempty" xml:"Area,omitempty"`
	// The domain names that you want to query. Separate domain names with commas (,).
	//
	// example:
	//
	// www.example.com,www.example.org
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the time range to query. Specify the time in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2021-07-02T01:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The HTTP status code. Valid values:
	//
	// 	- **2xx**: HTTP 2xx status codes
	//
	// 	- **3xx**: HTTP 3xx status codes
	//
	// 	- **4xx**: HTTP 4xx status codes
	//
	// 	- **5xx**: HTTP 5xx status codes
	//
	// If you do not specify an HTTP status code, data for all preceding HTTP status codes is queried.
	//
	// example:
	//
	// 2xx
	HttpCode *string `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	// Specify whether the region is outside the Chinese mainland. Valid values:
	//
	// 	- **1**: outside the Chinese mainland
	//
	// 	- **0**: inside the Chinese mainland
	//
	// example:
	//
	// 0
	IsOverseas *string `json:"IsOverseas,omitempty" xml:"IsOverseas,omitempty"`
	// The ID of the operations report that you want to query. You can enter only one ID in each call. You can call the [DescribeDcdnSubList](https://help.aliyun.com/document_detail/270075.html) operation to query report IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	ReportId *int64 `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
	// The beginning of the time range to query. Specify the time in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2021-07-02T00:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnReportRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnReportRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnReportRequest) GetArea() *string {
	return s.Area
}

func (s *DescribeDcdnReportRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDcdnReportRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDcdnReportRequest) GetHttpCode() *string {
	return s.HttpCode
}

func (s *DescribeDcdnReportRequest) GetIsOverseas() *string {
	return s.IsOverseas
}

func (s *DescribeDcdnReportRequest) GetReportId() *int64 {
	return s.ReportId
}

func (s *DescribeDcdnReportRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDcdnReportRequest) SetArea(v string) *DescribeDcdnReportRequest {
	s.Area = &v
	return s
}

func (s *DescribeDcdnReportRequest) SetDomainName(v string) *DescribeDcdnReportRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnReportRequest) SetEndTime(v string) *DescribeDcdnReportRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnReportRequest) SetHttpCode(v string) *DescribeDcdnReportRequest {
	s.HttpCode = &v
	return s
}

func (s *DescribeDcdnReportRequest) SetIsOverseas(v string) *DescribeDcdnReportRequest {
	s.IsOverseas = &v
	return s
}

func (s *DescribeDcdnReportRequest) SetReportId(v int64) *DescribeDcdnReportRequest {
	s.ReportId = &v
	return s
}

func (s *DescribeDcdnReportRequest) SetStartTime(v string) *DescribeDcdnReportRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDcdnReportRequest) Validate() error {
	return dara.Validate(s)
}
