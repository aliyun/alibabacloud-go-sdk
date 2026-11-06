// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDauReportDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppInstanceGroupId(v string) *QueryDauReportDataRequest
	GetAppInstanceGroupId() *string
	SetBizRegionId(v string) *QueryDauReportDataRequest
	GetBizRegionId() *string
	SetEndDate(v string) *QueryDauReportDataRequest
	GetEndDate() *string
	SetPageNumber(v int32) *QueryDauReportDataRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *QueryDauReportDataRequest
	GetPageSize() *int32
	SetProductType(v string) *QueryDauReportDataRequest
	GetProductType() *string
	SetStartDate(v string) *QueryDauReportDataRequest
	GetStartDate() *string
}

type QueryDauReportDataRequest struct {
	// The ID of the app instance group, which is used to filter data for a specific app instance group. You can call the [ListAppInstanceGroup](~~ListAppInstanceGroup~~) operation to obtain the ID. If this parameter is not specified, data is not filtered by app instance group.
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
	// The end time of the statistics period. This parameter is required. Format: yyyy-MM-dd HH:mm:ss. The value cannot be earlier than StartDate. The service extracts the natural day in the yyyy-MM-dd format. The string does not contain a time zone and is interpreted in the Asia/Shanghai (UTC+8) time zone.
	//
	// example:
	//
	// 2026-09-30 23:59:59
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// The page number to query. Specify a value starting from 1. If this parameter is not specified, the service processes the request as the first page, but the response may return 0.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of records per page. Default value: 200. If the specified value is greater than 1,500, the value is processed as 1,500.
	//
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The product type. This parameter is optional. If this parameter is not specified or is unrecognized, the query defaults to CloudApp.
	//
	// Valid values:
	//
	// - CloudApp: Wuying Cloud App.
	//
	// - CloudBrowser: Cloud browser.
	//
	// - WuyingServer: Enterprise Edition workstation.
	//
	// - WuyingWorkstation: Personal Edition Linggou container workstation.
	//
	// - WuyingWorkstationTeam: Linggou Team Edition container workstation.
	//
	// - WuyingWorkstationBusiness: Linggou Dedicated Edition container workstation.
	//
	// - AndroidCloud: Cloud phone.
	//
	// - AIAgent: AgentBay (AI agent).
	//
	// example:
	//
	// CloudApp
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The start time of the statistics period. This parameter is required. Format: yyyy-MM-dd HH:mm:ss. The service extracts the natural day in the yyyy-MM-dd format. The date range is inclusive and cannot exceed 180 days. The string does not contain a time zone and is interpreted in the Asia/Shanghai (UTC+8) time zone.
	//
	// example:
	//
	// 2026-09-01 00:00:00
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s QueryDauReportDataRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryDauReportDataRequest) GoString() string {
	return s.String()
}

func (s *QueryDauReportDataRequest) GetAppInstanceGroupId() *string {
	return s.AppInstanceGroupId
}

func (s *QueryDauReportDataRequest) GetBizRegionId() *string {
	return s.BizRegionId
}

func (s *QueryDauReportDataRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *QueryDauReportDataRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *QueryDauReportDataRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryDauReportDataRequest) GetProductType() *string {
	return s.ProductType
}

func (s *QueryDauReportDataRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *QueryDauReportDataRequest) SetAppInstanceGroupId(v string) *QueryDauReportDataRequest {
	s.AppInstanceGroupId = &v
	return s
}

func (s *QueryDauReportDataRequest) SetBizRegionId(v string) *QueryDauReportDataRequest {
	s.BizRegionId = &v
	return s
}

func (s *QueryDauReportDataRequest) SetEndDate(v string) *QueryDauReportDataRequest {
	s.EndDate = &v
	return s
}

func (s *QueryDauReportDataRequest) SetPageNumber(v int32) *QueryDauReportDataRequest {
	s.PageNumber = &v
	return s
}

func (s *QueryDauReportDataRequest) SetPageSize(v int32) *QueryDauReportDataRequest {
	s.PageSize = &v
	return s
}

func (s *QueryDauReportDataRequest) SetProductType(v string) *QueryDauReportDataRequest {
	s.ProductType = &v
	return s
}

func (s *QueryDauReportDataRequest) SetStartDate(v string) *QueryDauReportDataRequest {
	s.StartDate = &v
	return s
}

func (s *QueryDauReportDataRequest) Validate() error {
	return dara.Validate(s)
}
