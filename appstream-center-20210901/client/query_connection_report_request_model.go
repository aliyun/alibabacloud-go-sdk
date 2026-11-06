// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryConnectionReportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppInstanceGroupId(v string) *QueryConnectionReportRequest
	GetAppInstanceGroupId() *string
	SetBizRegionId(v string) *QueryConnectionReportRequest
	GetBizRegionId() *string
	SetEndDate(v string) *QueryConnectionReportRequest
	GetEndDate() *string
	SetPageNumber(v int32) *QueryConnectionReportRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *QueryConnectionReportRequest
	GetPageSize() *int32
	SetProductType(v string) *QueryConnectionReportRequest
	GetProductType() *string
	SetStartDate(v string) *QueryConnectionReportRequest
	GetStartDate() *string
}

type QueryConnectionReportRequest struct {
	// The ID of the app instance group, which is used to filter data for a specified app instance group. You can call [ListAppInstanceGroup](~~ListAppInstanceGroup~~) to obtain the ID. If you do not specify this parameter, data is not filtered by app instance group.
	//
	// example:
	//
	// aig-9ciijz60n4xsv****
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	// The ID of the business region, which is used to filter report data. You can call [ListRegions](~~ListRegions~~) to obtain the ID.
	//
	// example:
	//
	// cn-hangzhou
	BizRegionId *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	// The end time of the statistics. **Required**. Format: `yyyy-MM-dd HH:mm:ss`. The value cannot be earlier than `StartDate`. The statistical interval includes the start time but excludes the end time. The time span cannot exceed 60 months. The string itself does not carry a time zone and is interpreted in the Asia/Shanghai (UTC+8) time zone.
	//
	// example:
	//
	// 2026-09-02 00:00:00
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// The page number to query. We recommend that you explicitly specify a value starting from `1`. If you do not specify this parameter, the service processes the request as the first page, but the response may echo `0`.
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
	// The product type. **Required**.
	//
	// Valid values:
	//
	// - `CloudApp`: Wuying cloud app.
	//
	// - `CloudBrowser`: Cloud browser.
	//
	// - `WuyingServer`: Enterprise edition workstation.
	//
	// - `WuyingWorkstation`: Personal edition Linggou container workstation.
	//
	// - `WuyingWorkstationTeam`: Team edition Linggou container workstation.
	//
	// - `WuyingWorkstationBusiness`: Dedicated edition Linggou container workstation.
	//
	// - `AndroidCloud`: Cloud phone.
	//
	// - `AIAgent`: AgentBay (AI agent).
	//
	// example:
	//
	// CloudApp
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The start time of the statistics. **Required**. Format: `yyyy-MM-dd HH:mm:ss`. The statistical interval includes the start time but excludes the end time. The time span cannot exceed 60 months. The string itself does not carry a time zone and is interpreted in the Asia/Shanghai (UTC+8) time zone.
	//
	// example:
	//
	// 2026-09-01 00:00:00
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s QueryConnectionReportRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryConnectionReportRequest) GoString() string {
	return s.String()
}

func (s *QueryConnectionReportRequest) GetAppInstanceGroupId() *string {
	return s.AppInstanceGroupId
}

func (s *QueryConnectionReportRequest) GetBizRegionId() *string {
	return s.BizRegionId
}

func (s *QueryConnectionReportRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *QueryConnectionReportRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *QueryConnectionReportRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryConnectionReportRequest) GetProductType() *string {
	return s.ProductType
}

func (s *QueryConnectionReportRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *QueryConnectionReportRequest) SetAppInstanceGroupId(v string) *QueryConnectionReportRequest {
	s.AppInstanceGroupId = &v
	return s
}

func (s *QueryConnectionReportRequest) SetBizRegionId(v string) *QueryConnectionReportRequest {
	s.BizRegionId = &v
	return s
}

func (s *QueryConnectionReportRequest) SetEndDate(v string) *QueryConnectionReportRequest {
	s.EndDate = &v
	return s
}

func (s *QueryConnectionReportRequest) SetPageNumber(v int32) *QueryConnectionReportRequest {
	s.PageNumber = &v
	return s
}

func (s *QueryConnectionReportRequest) SetPageSize(v int32) *QueryConnectionReportRequest {
	s.PageSize = &v
	return s
}

func (s *QueryConnectionReportRequest) SetProductType(v string) *QueryConnectionReportRequest {
	s.ProductType = &v
	return s
}

func (s *QueryConnectionReportRequest) SetStartDate(v string) *QueryConnectionReportRequest {
	s.StartDate = &v
	return s
}

func (s *QueryConnectionReportRequest) Validate() error {
	return dara.Validate(s)
}
