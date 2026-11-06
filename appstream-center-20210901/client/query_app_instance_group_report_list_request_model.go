// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAppInstanceGroupReportListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppInstanceGroupId(v string) *QueryAppInstanceGroupReportListRequest
	GetAppInstanceGroupId() *string
	SetAppInstanceGroupName(v string) *QueryAppInstanceGroupReportListRequest
	GetAppInstanceGroupName() *string
	SetBizRegionId(v string) *QueryAppInstanceGroupReportListRequest
	GetBizRegionId() *string
	SetEndDate(v string) *QueryAppInstanceGroupReportListRequest
	GetEndDate() *string
	SetEndUserId(v string) *QueryAppInstanceGroupReportListRequest
	GetEndUserId() *string
	SetPageNumber(v int32) *QueryAppInstanceGroupReportListRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *QueryAppInstanceGroupReportListRequest
	GetPageSize() *int32
	SetRegionId(v string) *QueryAppInstanceGroupReportListRequest
	GetRegionId() *string
	SetStartDate(v string) *QueryAppInstanceGroupReportListRequest
	GetStartDate() *string
	SetStatDate(v string) *QueryAppInstanceGroupReportListRequest
	GetStatDate() *string
	SetStatus(v []*string) *QueryAppInstanceGroupReportListRequest
	GetStatus() []*string
}

type QueryAppInstanceGroupReportListRequest struct {
	// The filter condition for the app instance group ID, which matches the exact ID. If this parameter is not specified, the results are not filtered by ID. You can call [ListAppInstanceGroup](~~ListAppInstanceGroup~~) to obtain the ID.
	//
	// example:
	//
	// aig-9ciijz60n4xsv****
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	// The filter condition for the app instance group name, which supports inclusion matching. If this parameter is not specified, the results are not filtered by name.
	//
	// example:
	//
	// R&D workstation
	AppInstanceGroupName *string `json:"AppInstanceGroupName,omitempty" xml:"AppInstanceGroupName,omitempty"`
	// The ID of the business region to which the app instance group belongs. You can call [ListRegions](~~ListRegions~~) to obtain the ID. If this parameter is not specified, the results are not filtered by business region.
	//
	// example:
	//
	// cn-hangzhou
	BizRegionId *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	// The end date of the statistics. Format: `yyyyMMdd`. This parameter must be specified in pairs with `StartDate` and cannot be earlier than the start date. The string itself does not carry a time zone and is interpreted as a natural date in the Asia/Shanghai (UTC+8) time zone.
	//
	// example:
	//
	// 20260930
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// The ID of the end user, which is used to filter app instance groups based on the current valid authorization relationship. Before the call, make sure that the user has at least one valid authorization. If no valid authorization exists, the current version does not narrow down the results based on this parameter. If this parameter is not specified, the results are not filtered by user.
	//
	// example:
	//
	// user01
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// The page number to query. Specify a value starting from `1`. If this parameter is not specified, the service processes the request as the first page, but the response may return `0`.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 200. The current operation does not have an additional upper limit for a single page.
	//
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID for OpenAPI access, which is used for request authentication and routing. To filter business data by region, use `BizRegionId`.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The start date of the statistics. Format: `yyyyMMdd`. This parameter must be specified in pairs with `EndDate`, and the interval includes the start and end dates. The string itself does not carry a time zone and is interpreted as a natural date in the Asia/Shanghai (UTC+8) time zone.
	//
	// example:
	//
	// 20260901
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// The statistical date for a single day. Format: `yyyyMMdd`. If this parameter is specified along with `StartDate` and `EndDate`, the results must meet both the single-day and interval conditions. The string itself does not carry a time zone and is interpreted as a natural date in the Asia/Shanghai (UTC+8) time zone.
	//
	// example:
	//
	// 20260901
	StatDate *string `json:"StatDate,omitempty" xml:"StatDate,omitempty"`
	// The list of app instance group statuses. Multiple values have an OR relationship. Unrecognized values do not participate in filtering. If all values are unrecognized, the results are not filtered by status.
	//
	// Valid values:
	//
	// - `DEPLOYING`: Deploying.
	//
	// - `DEPLOYED`: Deployed but not published.
	//
	// - `PUBLISHED`: Published and available to provide services based on the current configurations.
	//
	// - `STOPPING`: Stopping.
	//
	// - `STOPPED`: Stopped.
	//
	// - `MAINTAINING`: Maintaining.
	//
	// - `FAILED`: Deployment or running failed.
	//
	// - `MAINTAIN_FAILED`: Maintenance failed.
	//
	// - `DELETING`: Deleting.
	//
	// - `DELETED`: Deleted.
	Status []*string `json:"Status,omitempty" xml:"Status,omitempty" type:"Repeated"`
}

func (s QueryAppInstanceGroupReportListRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryAppInstanceGroupReportListRequest) GoString() string {
	return s.String()
}

func (s *QueryAppInstanceGroupReportListRequest) GetAppInstanceGroupId() *string {
	return s.AppInstanceGroupId
}

func (s *QueryAppInstanceGroupReportListRequest) GetAppInstanceGroupName() *string {
	return s.AppInstanceGroupName
}

func (s *QueryAppInstanceGroupReportListRequest) GetBizRegionId() *string {
	return s.BizRegionId
}

func (s *QueryAppInstanceGroupReportListRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *QueryAppInstanceGroupReportListRequest) GetEndUserId() *string {
	return s.EndUserId
}

func (s *QueryAppInstanceGroupReportListRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *QueryAppInstanceGroupReportListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryAppInstanceGroupReportListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *QueryAppInstanceGroupReportListRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *QueryAppInstanceGroupReportListRequest) GetStatDate() *string {
	return s.StatDate
}

func (s *QueryAppInstanceGroupReportListRequest) GetStatus() []*string {
	return s.Status
}

func (s *QueryAppInstanceGroupReportListRequest) SetAppInstanceGroupId(v string) *QueryAppInstanceGroupReportListRequest {
	s.AppInstanceGroupId = &v
	return s
}

func (s *QueryAppInstanceGroupReportListRequest) SetAppInstanceGroupName(v string) *QueryAppInstanceGroupReportListRequest {
	s.AppInstanceGroupName = &v
	return s
}

func (s *QueryAppInstanceGroupReportListRequest) SetBizRegionId(v string) *QueryAppInstanceGroupReportListRequest {
	s.BizRegionId = &v
	return s
}

func (s *QueryAppInstanceGroupReportListRequest) SetEndDate(v string) *QueryAppInstanceGroupReportListRequest {
	s.EndDate = &v
	return s
}

func (s *QueryAppInstanceGroupReportListRequest) SetEndUserId(v string) *QueryAppInstanceGroupReportListRequest {
	s.EndUserId = &v
	return s
}

func (s *QueryAppInstanceGroupReportListRequest) SetPageNumber(v int32) *QueryAppInstanceGroupReportListRequest {
	s.PageNumber = &v
	return s
}

func (s *QueryAppInstanceGroupReportListRequest) SetPageSize(v int32) *QueryAppInstanceGroupReportListRequest {
	s.PageSize = &v
	return s
}

func (s *QueryAppInstanceGroupReportListRequest) SetRegionId(v string) *QueryAppInstanceGroupReportListRequest {
	s.RegionId = &v
	return s
}

func (s *QueryAppInstanceGroupReportListRequest) SetStartDate(v string) *QueryAppInstanceGroupReportListRequest {
	s.StartDate = &v
	return s
}

func (s *QueryAppInstanceGroupReportListRequest) SetStatDate(v string) *QueryAppInstanceGroupReportListRequest {
	s.StatDate = &v
	return s
}

func (s *QueryAppInstanceGroupReportListRequest) SetStatus(v []*string) *QueryAppInstanceGroupReportListRequest {
	s.Status = v
	return s
}

func (s *QueryAppInstanceGroupReportListRequest) Validate() error {
	return dara.Validate(s)
}
