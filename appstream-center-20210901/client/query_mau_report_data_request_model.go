// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMauReportDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppInstanceGroupId(v string) *QueryMauReportDataRequest
	GetAppInstanceGroupId() *string
	SetBizRegionId(v string) *QueryMauReportDataRequest
	GetBizRegionId() *string
	SetEndDate(v string) *QueryMauReportDataRequest
	GetEndDate() *string
	SetPageNumber(v int32) *QueryMauReportDataRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *QueryMauReportDataRequest
	GetPageSize() *int32
	SetProductType(v string) *QueryMauReportDataRequest
	GetProductType() *string
	SetStartDate(v string) *QueryMauReportDataRequest
	GetStartDate() *string
}

type QueryMauReportDataRequest struct {
	// The app instance group ID, which is used to filter data for a specified app instance group. You can call the [ListAppInstanceGroup](~~ListAppInstanceGroup~~) operation to obtain the ID. If this parameter is not specified, data is not filtered by app instance group.
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
	// The end time of the statistical period. **Required**. Format: `yyyy-MM-dd HH:mm:ss`. The value cannot be earlier than `StartDate`. The statistical interval includes the start and end boundaries. The time span cannot exceed 60 months. The string itself does not carry a time zone and is interpreted in the Asia/Shanghai (UTC+8) time zone.
	//
	// example:
	//
	// 2026-09-02 00:00:00
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// The page number to query. We recommend that you explicitly specify a value starting from `1`. If this parameter is not specified, the service processes the request as the first page, but the response may return `0`.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 200. If the specified value is greater than 1,500, it is processed as 1,500.
	//
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The product type. **Required**.
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
	// The start time of the statistical period. **Required**. Format: `yyyy-MM-dd HH:mm:ss`. The statistical interval includes the start and end boundaries. The time span cannot exceed 60 months. The string itself does not carry a time zone and is interpreted in the Asia/Shanghai (UTC+8) time zone.
	//
	// example:
	//
	// 2026-09-01 00:00:00
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s QueryMauReportDataRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryMauReportDataRequest) GoString() string {
	return s.String()
}

func (s *QueryMauReportDataRequest) GetAppInstanceGroupId() *string {
	return s.AppInstanceGroupId
}

func (s *QueryMauReportDataRequest) GetBizRegionId() *string {
	return s.BizRegionId
}

func (s *QueryMauReportDataRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *QueryMauReportDataRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *QueryMauReportDataRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryMauReportDataRequest) GetProductType() *string {
	return s.ProductType
}

func (s *QueryMauReportDataRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *QueryMauReportDataRequest) SetAppInstanceGroupId(v string) *QueryMauReportDataRequest {
	s.AppInstanceGroupId = &v
	return s
}

func (s *QueryMauReportDataRequest) SetBizRegionId(v string) *QueryMauReportDataRequest {
	s.BizRegionId = &v
	return s
}

func (s *QueryMauReportDataRequest) SetEndDate(v string) *QueryMauReportDataRequest {
	s.EndDate = &v
	return s
}

func (s *QueryMauReportDataRequest) SetPageNumber(v int32) *QueryMauReportDataRequest {
	s.PageNumber = &v
	return s
}

func (s *QueryMauReportDataRequest) SetPageSize(v int32) *QueryMauReportDataRequest {
	s.PageSize = &v
	return s
}

func (s *QueryMauReportDataRequest) SetProductType(v string) *QueryMauReportDataRequest {
	s.ProductType = &v
	return s
}

func (s *QueryMauReportDataRequest) SetStartDate(v string) *QueryMauReportDataRequest {
	s.StartDate = &v
	return s
}

func (s *QueryMauReportDataRequest) Validate() error {
	return dara.Validate(s)
}
