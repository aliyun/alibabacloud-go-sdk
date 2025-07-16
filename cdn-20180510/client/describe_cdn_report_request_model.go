// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnReportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArea(v string) *DescribeCdnReportRequest
	GetArea() *string
	SetDomainName(v string) *DescribeCdnReportRequest
	GetDomainName() *string
	SetEndTime(v string) *DescribeCdnReportRequest
	GetEndTime() *string
	SetHttpCode(v string) *DescribeCdnReportRequest
	GetHttpCode() *string
	SetIsOverseas(v string) *DescribeCdnReportRequest
	GetIsOverseas() *string
	SetReportId(v int64) *DescribeCdnReportRequest
	GetReportId() *int64
	SetStartTime(v string) *DescribeCdnReportRequest
	GetStartTime() *string
}

type DescribeCdnReportRequest struct {
	// The region. You can call the [DescribeCdnRegionAndIsp](https://help.aliyun.com/document_detail/91077.html) operation to query regions.
	//
	// 	- If you do not specify a region, data in all regions is queried.
	//
	// 	- If you specify a region, data in the specified region is queried. You can specify one or more regions. If you specify multiple regions, separate the regions with commas (,).
	//
	// example:
	//
	// shanghai
	Area *string `json:"Area,omitempty" xml:"Area,omitempty"`
	// The domain name that you want to query. Separate domain names with commas (,).
	//
	// example:
	//
	// www.example1.com,example2.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2020-09-17T01:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The HTTP status code. Valid values:
	//
	// 	- **2xx**
	//
	// 	- **3xx**
	//
	// 	- **4xx**
	//
	// 	- **5xx**
	//
	// If you do not specify this parameter, all HTTP status codes are queried.
	//
	// example:
	//
	// 2xx
	HttpCode *string `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	// Specifies whether the region is outside the Chinese mainland. Valid values:
	//
	// 	- **1**: outside the Chinese mainland
	//
	// 	- **0**: inside the Chinese mainland
	//
	// example:
	//
	// 0
	IsOverseas *string `json:"IsOverseas,omitempty" xml:"IsOverseas,omitempty"`
	// The ID of the operations report that you want to query. You can specify only one ID in each request. You can call the [DescribeCdnSubList](https://help.aliyun.com/document_detail/271655.html) operation to query report IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ReportId *int64 `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2020-09-17T00:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeCdnReportRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnReportRequest) GoString() string {
	return s.String()
}

func (s *DescribeCdnReportRequest) GetArea() *string {
	return s.Area
}

func (s *DescribeCdnReportRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeCdnReportRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeCdnReportRequest) GetHttpCode() *string {
	return s.HttpCode
}

func (s *DescribeCdnReportRequest) GetIsOverseas() *string {
	return s.IsOverseas
}

func (s *DescribeCdnReportRequest) GetReportId() *int64 {
	return s.ReportId
}

func (s *DescribeCdnReportRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeCdnReportRequest) SetArea(v string) *DescribeCdnReportRequest {
	s.Area = &v
	return s
}

func (s *DescribeCdnReportRequest) SetDomainName(v string) *DescribeCdnReportRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeCdnReportRequest) SetEndTime(v string) *DescribeCdnReportRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeCdnReportRequest) SetHttpCode(v string) *DescribeCdnReportRequest {
	s.HttpCode = &v
	return s
}

func (s *DescribeCdnReportRequest) SetIsOverseas(v string) *DescribeCdnReportRequest {
	s.IsOverseas = &v
	return s
}

func (s *DescribeCdnReportRequest) SetReportId(v int64) *DescribeCdnReportRequest {
	s.ReportId = &v
	return s
}

func (s *DescribeCdnReportRequest) SetStartTime(v string) *DescribeCdnReportRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeCdnReportRequest) Validate() error {
	return dara.Validate(s)
}
