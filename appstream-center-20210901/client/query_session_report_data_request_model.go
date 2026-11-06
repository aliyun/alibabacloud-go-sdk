// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySessionReportDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppInstanceGroupId(v string) *QuerySessionReportDataRequest
	GetAppInstanceGroupId() *string
	SetBizRegionId(v string) *QuerySessionReportDataRequest
	GetBizRegionId() *string
	SetEndDate(v string) *QuerySessionReportDataRequest
	GetEndDate() *string
	SetPageNumber(v int32) *QuerySessionReportDataRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *QuerySessionReportDataRequest
	GetPageSize() *int32
	SetProductType(v string) *QuerySessionReportDataRequest
	GetProductType() *string
	SetStartDate(v string) *QuerySessionReportDataRequest
	GetStartDate() *string
	SetStatDateType(v string) *QuerySessionReportDataRequest
	GetStatDateType() *string
}

type QuerySessionReportDataRequest struct {
	// The ID of the app instance group, which is used to filter data for a specified app instance group. You can call the [ListAppInstanceGroup](~~ListAppInstanceGroup~~) operation to obtain the ID. If you do not specify this parameter, data is not filtered by app instance group.
	//
	// example:
	//
	// aig-9ciijz60n4xsv****
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	// The ID of the business region, which is used to filter report data. You can call the [ListRegions](~~ListRegions~~) operation to obtain the ID.
	//
	// example:
	//
	// cn-hangzhou
	BizRegionId *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	// The end time of the statistics. This parameter is required. Format: `yyyy-MM-dd HH:mm:ss`. The end time cannot be earlier than `StartDate`. The query span must meet the limit corresponding to `StatDateType`. The string itself does not carry a time zone and is interpreted in the Asia/Shanghai (UTC+8) time zone.
	//
	// example:
	//
	// 2026-09-02 00:00:00
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// The page number of the page to query. Specify this parameter starting from `1`. If you do not specify this parameter, the first page is queried, but the response may return `0`.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Default value: 200. If the specified value is greater than 1500, 1500 is used.
	//
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The product type. This parameter is optional. If you do not specify this parameter, data is queried based on the Wuying Cloud App caliber.
	//
	// Specify the value in the exact case shown below. Data that does not match the specified value is not returned.
	//
	// Valid values:
	//
	// - `CloudApp`: Wuying Cloud App.
	//
	// - `CloudBrowser`: Cloud Browser.
	//
	// - `WuyingServer`: Enterprise Edition workstation.
	//
	// - `WuyingWorkstation`: Personal Edition Linggou container workstation.
	//
	// - `WuyingWorkstationTeam`: Team Edition Linggou container workstation.
	//
	// - `WuyingWorkstationBusiness`: Dedicated Edition Linggou container workstation.
	//
	// - `AndroidCloud`: Cloud Phone.
	//
	// - `AIAgent`: AgentBay (AI agent).
	//
	// example:
	//
	// CloudApp
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The start time of the statistics. This parameter is required. Format: `yyyy-MM-dd HH:mm:ss`. The `1d` granularity aggregates data by date, while other granularities aggregate data down to the minute. The query includes both the start and end boundaries. The string itself does not carry a time zone and is interpreted in the Asia/Shanghai (UTC+8) time zone.
	//
	// example:
	//
	// 2026-09-01 00:00:00
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// The statistical granularity. This parameter is required.
	//
	// Valid values:
	//
	// - `1m`: aggregates data by 1 minute. The maximum query span is 1 day.
	//
	// - `30m`: aggregates data by 30 minutes. The maximum query span is 1 day.
	//
	// - `6h`: aggregates data by 6 hours. The maximum query span is 15 days.
	//
	// - `1d`: aggregates data by day. The maximum query span is 180 days.
	//
	// example:
	//
	// 1d
	StatDateType *string `json:"StatDateType,omitempty" xml:"StatDateType,omitempty"`
}

func (s QuerySessionReportDataRequest) String() string {
	return dara.Prettify(s)
}

func (s QuerySessionReportDataRequest) GoString() string {
	return s.String()
}

func (s *QuerySessionReportDataRequest) GetAppInstanceGroupId() *string {
	return s.AppInstanceGroupId
}

func (s *QuerySessionReportDataRequest) GetBizRegionId() *string {
	return s.BizRegionId
}

func (s *QuerySessionReportDataRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *QuerySessionReportDataRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *QuerySessionReportDataRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QuerySessionReportDataRequest) GetProductType() *string {
	return s.ProductType
}

func (s *QuerySessionReportDataRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *QuerySessionReportDataRequest) GetStatDateType() *string {
	return s.StatDateType
}

func (s *QuerySessionReportDataRequest) SetAppInstanceGroupId(v string) *QuerySessionReportDataRequest {
	s.AppInstanceGroupId = &v
	return s
}

func (s *QuerySessionReportDataRequest) SetBizRegionId(v string) *QuerySessionReportDataRequest {
	s.BizRegionId = &v
	return s
}

func (s *QuerySessionReportDataRequest) SetEndDate(v string) *QuerySessionReportDataRequest {
	s.EndDate = &v
	return s
}

func (s *QuerySessionReportDataRequest) SetPageNumber(v int32) *QuerySessionReportDataRequest {
	s.PageNumber = &v
	return s
}

func (s *QuerySessionReportDataRequest) SetPageSize(v int32) *QuerySessionReportDataRequest {
	s.PageSize = &v
	return s
}

func (s *QuerySessionReportDataRequest) SetProductType(v string) *QuerySessionReportDataRequest {
	s.ProductType = &v
	return s
}

func (s *QuerySessionReportDataRequest) SetStartDate(v string) *QuerySessionReportDataRequest {
	s.StartDate = &v
	return s
}

func (s *QuerySessionReportDataRequest) SetStatDateType(v string) *QuerySessionReportDataRequest {
	s.StatDateType = &v
	return s
}

func (s *QuerySessionReportDataRequest) Validate() error {
	return dara.Validate(s)
}
