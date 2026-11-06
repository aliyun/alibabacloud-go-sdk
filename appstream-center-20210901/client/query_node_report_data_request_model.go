// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryNodeReportDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppInstanceGroupId(v string) *QueryNodeReportDataRequest
	GetAppInstanceGroupId() *string
	SetBizRegionId(v string) *QueryNodeReportDataRequest
	GetBizRegionId() *string
	SetEndDate(v string) *QueryNodeReportDataRequest
	GetEndDate() *string
	SetPageNumber(v int32) *QueryNodeReportDataRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *QueryNodeReportDataRequest
	GetPageSize() *int32
	SetProductType(v string) *QueryNodeReportDataRequest
	GetProductType() *string
	SetStartDate(v string) *QueryNodeReportDataRequest
	GetStartDate() *string
	SetStatDateType(v string) *QueryNodeReportDataRequest
	GetStatDateType() *string
}

type QueryNodeReportDataRequest struct {
	// The ID of the app instance group, which is used to filter data for a specified app instance group. You can call the [ListAppInstanceGroup](~~ListAppInstanceGroup~~) operation to obtain the ID. If you do not specify this parameter, data is not filtered by app instance group.
	//
	// example:
	//
	// aig-9ciijz60n4xsv****
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	// The business region ID, which is used to filter report data. You can call the [ListRegions](~~ListRegions~~) operation to obtain the ID.
	//
	// example:
	//
	// cn-hangzhou
	BizRegionId *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	// The end time of the statistics. This parameter is required. Format: yyyy-MM-dd HH:mm:ss. The end time cannot be earlier than the start time. The query span must meet the limit corresponding to the statistical granularity. The string itself does not carry a time zone and is interpreted in the Asia/Shanghai (UTC+8) time zone.
	//
	// example:
	//
	// 2026-09-02 00:00:00
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// The page number to query. Explicitly pass a value starting from 1. If you do not specify this parameter, the service processes the request as the first page, but the response may return 0.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. If you do not specify this parameter, the default value is 200. If the value is greater than 1500, it is processed as 1500.
	//
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The product type. This parameter is optional. If you do not specify this parameter, data is queried for Wuying Cloud App.
	//
	// Specify the value in the exact case shown below. Data that does not match the specified case is not returned.
	//
	// Valid values:
	//
	// - CloudApp: Wuying Cloud App.
	//
	// - CloudBrowser: Cloud Browser.
	//
	// - WuyingServer: Enterprise Edition workstation.
	//
	// - WuyingWorkstation: Personal Edition Linggou container workstation.
	//
	// - WuyingWorkstationTeam: Linggou Team Edition container workstation.
	//
	// - WuyingWorkstationBusiness: Linggou Dedicated Edition container workstation.
	//
	// - AndroidCloud: Cloud Phone.
	//
	// - AIAgent: AgentBay (AI agent).
	//
	// example:
	//
	// CloudApp
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The start time of the statistics. This parameter is required. Format: yyyy-MM-dd HH:mm:ss. The 1d granularity aggregates data by date, while other granularities aggregate data down to the minute. The query includes both the start and end boundaries. The string itself does not carry a time zone and is interpreted in the Asia/Shanghai (UTC+8) time zone.
	//
	// example:
	//
	// 2026-09-01 00:00:00
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// The statistical granularity. This parameter is required.
	//
	// Valid values:
	//
	// - 1m: Aggregates data by 1 minute. The maximum query span is 1 day.
	//
	// - 30m: Aggregates data by 30 minutes. The maximum query span is 1 day.
	//
	// - 6h: Aggregates data by 6 hours. The maximum query span is 15 days.
	//
	// - 1d: Aggregates data by day. The maximum query span is 180 days.
	//
	// example:
	//
	// 1d
	StatDateType *string `json:"StatDateType,omitempty" xml:"StatDateType,omitempty"`
}

func (s QueryNodeReportDataRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryNodeReportDataRequest) GoString() string {
	return s.String()
}

func (s *QueryNodeReportDataRequest) GetAppInstanceGroupId() *string {
	return s.AppInstanceGroupId
}

func (s *QueryNodeReportDataRequest) GetBizRegionId() *string {
	return s.BizRegionId
}

func (s *QueryNodeReportDataRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *QueryNodeReportDataRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *QueryNodeReportDataRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryNodeReportDataRequest) GetProductType() *string {
	return s.ProductType
}

func (s *QueryNodeReportDataRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *QueryNodeReportDataRequest) GetStatDateType() *string {
	return s.StatDateType
}

func (s *QueryNodeReportDataRequest) SetAppInstanceGroupId(v string) *QueryNodeReportDataRequest {
	s.AppInstanceGroupId = &v
	return s
}

func (s *QueryNodeReportDataRequest) SetBizRegionId(v string) *QueryNodeReportDataRequest {
	s.BizRegionId = &v
	return s
}

func (s *QueryNodeReportDataRequest) SetEndDate(v string) *QueryNodeReportDataRequest {
	s.EndDate = &v
	return s
}

func (s *QueryNodeReportDataRequest) SetPageNumber(v int32) *QueryNodeReportDataRequest {
	s.PageNumber = &v
	return s
}

func (s *QueryNodeReportDataRequest) SetPageSize(v int32) *QueryNodeReportDataRequest {
	s.PageSize = &v
	return s
}

func (s *QueryNodeReportDataRequest) SetProductType(v string) *QueryNodeReportDataRequest {
	s.ProductType = &v
	return s
}

func (s *QueryNodeReportDataRequest) SetStartDate(v string) *QueryNodeReportDataRequest {
	s.StartDate = &v
	return s
}

func (s *QueryNodeReportDataRequest) SetStatDateType(v string) *QueryNodeReportDataRequest {
	s.StatDateType = &v
	return s
}

func (s *QueryNodeReportDataRequest) Validate() error {
	return dara.Validate(s)
}
