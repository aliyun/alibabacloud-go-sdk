// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryNetworkRttReportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppInstanceGroupId(v string) *QueryNetworkRttReportRequest
	GetAppInstanceGroupId() *string
	SetBizRegionId(v string) *QueryNetworkRttReportRequest
	GetBizRegionId() *string
	SetEndDate(v string) *QueryNetworkRttReportRequest
	GetEndDate() *string
	SetPageNumber(v int32) *QueryNetworkRttReportRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *QueryNetworkRttReportRequest
	GetPageSize() *int32
	SetProductType(v string) *QueryNetworkRttReportRequest
	GetProductType() *string
	SetStartDate(v string) *QueryNetworkRttReportRequest
	GetStartDate() *string
}

type QueryNetworkRttReportRequest struct {
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
	// The end time of the statistics. **Required**. Format: `yyyy-MM-dd HH:mm:ss`. The end time cannot be earlier than `StartDate`. The statistical interval includes the start time but excludes the end time. The time span cannot exceed 60 months. The string itself does not carry a time zone and is interpreted as Asia/Shanghai (UTC+8).
	//
	// example:
	//
	// 2026-09-02 00:00:00
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// The page number to query. Explicitly pass a value starting from `1`. If you do not specify this parameter, the service processes the request as the first page, but the response may return `0`.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. If you do not specify this parameter, the default value is 200. If the value is greater than 1,500, it is processed as 1,500.
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
	// - `CloudBrowser`: Cloud browser.
	//
	// - `WuyingServer`: Enterprise Edition workstation.
	//
	// - `WuyingWorkstation`: Personal Edition Linggou container workstation.
	//
	// - `WuyingWorkstationTeam`: Linggou Team Edition container workstation.
	//
	// - `WuyingWorkstationBusiness`: Linggou Dedicated Edition container workstation.
	//
	// - `AndroidCloud`: Cloud phone.
	//
	// - `AIAgent`: AgentBay (AI agent).
	//
	// example:
	//
	// CloudApp
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The start time of the statistics. **Required**. Format: `yyyy-MM-dd HH:mm:ss`. The statistical interval includes the start time but excludes the end time. The time span cannot exceed 60 months. The string itself does not carry a time zone and is interpreted as Asia/Shanghai (UTC+8).
	//
	// example:
	//
	// 2026-09-01 00:00:00
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s QueryNetworkRttReportRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryNetworkRttReportRequest) GoString() string {
	return s.String()
}

func (s *QueryNetworkRttReportRequest) GetAppInstanceGroupId() *string {
	return s.AppInstanceGroupId
}

func (s *QueryNetworkRttReportRequest) GetBizRegionId() *string {
	return s.BizRegionId
}

func (s *QueryNetworkRttReportRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *QueryNetworkRttReportRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *QueryNetworkRttReportRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryNetworkRttReportRequest) GetProductType() *string {
	return s.ProductType
}

func (s *QueryNetworkRttReportRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *QueryNetworkRttReportRequest) SetAppInstanceGroupId(v string) *QueryNetworkRttReportRequest {
	s.AppInstanceGroupId = &v
	return s
}

func (s *QueryNetworkRttReportRequest) SetBizRegionId(v string) *QueryNetworkRttReportRequest {
	s.BizRegionId = &v
	return s
}

func (s *QueryNetworkRttReportRequest) SetEndDate(v string) *QueryNetworkRttReportRequest {
	s.EndDate = &v
	return s
}

func (s *QueryNetworkRttReportRequest) SetPageNumber(v int32) *QueryNetworkRttReportRequest {
	s.PageNumber = &v
	return s
}

func (s *QueryNetworkRttReportRequest) SetPageSize(v int32) *QueryNetworkRttReportRequest {
	s.PageSize = &v
	return s
}

func (s *QueryNetworkRttReportRequest) SetProductType(v string) *QueryNetworkRttReportRequest {
	s.ProductType = &v
	return s
}

func (s *QueryNetworkRttReportRequest) SetStartDate(v string) *QueryNetworkRttReportRequest {
	s.StartDate = &v
	return s
}

func (s *QueryNetworkRttReportRequest) Validate() error {
	return dara.Validate(s)
}
