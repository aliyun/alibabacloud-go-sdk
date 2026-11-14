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
	// The end time of the statistics. **Required**. Format: `yyyy-MM-dd HH:mm:ss`. The end time cannot be earlier than `StartDate`. The query span must meet the limits corresponding to `StatDateType`. The string itself does not carry a time zone and is interpreted as Asia/Shanghai (UTC+8).
	//
	// example:
	//
	// 2026-09-02 00:00:00
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// The page number to query. We recommend that you explicitly specify a value starting from `1`. If you do not specify this parameter, the service processes the request as the first page, but the response may return `0`.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 200. If the value exceeds 1,500, it is processed as 1,500.
	//
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The product type. Optional. If you do not specify this parameter, data is queried based on the Wuying Cloud App scope.
	//
	// Specify the value in the exact case as follows. Data that cannot be matched is not returned.
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
	// The start time of the statistics. **Required**. Format: `yyyy-MM-dd HH:mm:ss`. The `1d` granularity aggregates data by date, while other granularities are accurate to the minute. The query includes both the start and end boundaries. The string itself does not carry a time zone and is interpreted as Asia/Shanghai (UTC+8).
	//
	// example:
	//
	// 2026-09-01 00:00:00
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// The statistical granularity. **Required**.
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
