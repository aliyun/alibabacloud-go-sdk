// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryUserUsageReportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppInstanceGroupId(v string) *QueryUserUsageReportRequest
	GetAppInstanceGroupId() *string
	SetBizRegionId(v string) *QueryUserUsageReportRequest
	GetBizRegionId() *string
	SetConnectionStatus(v string) *QueryUserUsageReportRequest
	GetConnectionStatus() *string
	SetEndDate(v string) *QueryUserUsageReportRequest
	GetEndDate() *string
	SetPageNumber(v int32) *QueryUserUsageReportRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *QueryUserUsageReportRequest
	GetPageSize() *int32
	SetProductType(v string) *QueryUserUsageReportRequest
	GetProductType() *string
	SetStartDate(v string) *QueryUserUsageReportRequest
	GetStartDate() *string
	SetUserId(v string) *QueryUserUsageReportRequest
	GetUserId() *string
}

type QueryUserUsageReportRequest struct {
	// The ID of the app instance group. This parameter is used to filter data for a specified app instance group. You can call the [ListAppInstanceGroup](~~ListAppInstanceGroup~~) operation to obtain the ID. If you do not specify this parameter, data is not filtered by app instance group.
	//
	// example:
	//
	// aig-9ciijz60n4xsv****
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	// The ID of the business region. This parameter is used to filter report data. You can call the [ListRegions](~~ListRegions~~) operation to obtain the ID.
	//
	// example:
	//
	// cn-hangzhou
	BizRegionId *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	// The filter condition for the user connection status. The value is not case-sensitive. An unrecognized value is equivalent to not filtering by connection status.
	//
	// Valid values:
	//
	// - `Connected`: At least one currently connected session exists.
	//
	// - `Disconnected`: The session binding is retained, but no currently connected session exists.
	//
	// - `Unbound`: No session binding exists.
	//
	// example:
	//
	// Connected
	ConnectionStatus *string `json:"ConnectionStatus,omitempty" xml:"ConnectionStatus,omitempty"`
	// The end time of the statistics. **Required**. Format: `yyyy-MM-dd HH:mm:ss`. The end time cannot be earlier than StartDate. The service converts the time to `yyyyMMdd` and calculates the statistics by natural days in a closed interval. The hours, minutes, and seconds are not involved in the day-level filtering. The time span cannot exceed 60 months. The string itself does not carry a time zone and is interpreted as Asia/Shanghai (UTC+8).
	//
	// example:
	//
	// 2026-09-02 00:00:00
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// The page number to query. Specify a value starting from 1. If you do not specify this parameter, the service processes the request as the first page, but the response may return 0.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 200. If the value is greater than 1,500, it is processed as 1,500.
	//
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The product type. **Required**.
	//
	// Valid values:
	//
	// - `CloudApp`: WUYING Cloud App.
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
	// The start time of the statistics. **Required**. Format: `yyyy-MM-dd HH:mm:ss`. The service converts the time to `yyyyMMdd` and calculates the statistics by natural days in a closed interval. The hours, minutes, and seconds are not involved in the day-level filtering. The time span cannot exceed 60 months. The string itself does not carry a time zone and is interpreted as Asia/Shanghai (UTC+8).
	//
	// example:
	//
	// 2026-09-01 00:00:00
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// The ID of the end user. This parameter is used to accurately filter users. If you do not specify this parameter, data is not filtered by user.
	//
	// example:
	//
	// user01
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s QueryUserUsageReportRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryUserUsageReportRequest) GoString() string {
	return s.String()
}

func (s *QueryUserUsageReportRequest) GetAppInstanceGroupId() *string {
	return s.AppInstanceGroupId
}

func (s *QueryUserUsageReportRequest) GetBizRegionId() *string {
	return s.BizRegionId
}

func (s *QueryUserUsageReportRequest) GetConnectionStatus() *string {
	return s.ConnectionStatus
}

func (s *QueryUserUsageReportRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *QueryUserUsageReportRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *QueryUserUsageReportRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryUserUsageReportRequest) GetProductType() *string {
	return s.ProductType
}

func (s *QueryUserUsageReportRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *QueryUserUsageReportRequest) GetUserId() *string {
	return s.UserId
}

func (s *QueryUserUsageReportRequest) SetAppInstanceGroupId(v string) *QueryUserUsageReportRequest {
	s.AppInstanceGroupId = &v
	return s
}

func (s *QueryUserUsageReportRequest) SetBizRegionId(v string) *QueryUserUsageReportRequest {
	s.BizRegionId = &v
	return s
}

func (s *QueryUserUsageReportRequest) SetConnectionStatus(v string) *QueryUserUsageReportRequest {
	s.ConnectionStatus = &v
	return s
}

func (s *QueryUserUsageReportRequest) SetEndDate(v string) *QueryUserUsageReportRequest {
	s.EndDate = &v
	return s
}

func (s *QueryUserUsageReportRequest) SetPageNumber(v int32) *QueryUserUsageReportRequest {
	s.PageNumber = &v
	return s
}

func (s *QueryUserUsageReportRequest) SetPageSize(v int32) *QueryUserUsageReportRequest {
	s.PageSize = &v
	return s
}

func (s *QueryUserUsageReportRequest) SetProductType(v string) *QueryUserUsageReportRequest {
	s.ProductType = &v
	return s
}

func (s *QueryUserUsageReportRequest) SetStartDate(v string) *QueryUserUsageReportRequest {
	s.StartDate = &v
	return s
}

func (s *QueryUserUsageReportRequest) SetUserId(v string) *QueryUserUsageReportRequest {
	s.UserId = &v
	return s
}

func (s *QueryUserUsageReportRequest) Validate() error {
	return dara.Validate(s)
}
