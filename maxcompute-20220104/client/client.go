// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CreatePackageRequest struct {
	// The request body parameters.
	Body *string `json:"body,omitempty" xml:"body,omitempty"`
	// Specifies whether to install the package.
	IsInstall *bool `json:"isInstall,omitempty" xml:"isInstall,omitempty"`
}

func (s CreatePackageRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePackageRequest) GoString() string {
	return s.String()
}

func (s *CreatePackageRequest) SetBody(v string) *CreatePackageRequest {
	s.Body = &v
	return s
}

func (s *CreatePackageRequest) SetIsInstall(v bool) *CreatePackageRequest {
	s.IsInstall = &v
	return s
}

type CreatePackageResponseBody struct {
	// The returned data.
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreatePackageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreatePackageResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePackageResponseBody) SetData(v string) *CreatePackageResponseBody {
	s.Data = &v
	return s
}

func (s *CreatePackageResponseBody) SetRequestId(v string) *CreatePackageResponseBody {
	s.RequestId = &v
	return s
}

type CreatePackageResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreatePackageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreatePackageResponse) String() string {
	return tea.Prettify(s)
}

func (s CreatePackageResponse) GoString() string {
	return s.String()
}

func (s *CreatePackageResponse) SetHeaders(v map[string]*string) *CreatePackageResponse {
	s.Headers = v
	return s
}

func (s *CreatePackageResponse) SetStatusCode(v int32) *CreatePackageResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePackageResponse) SetBody(v *CreatePackageResponseBody) *CreatePackageResponse {
	s.Body = v
	return s
}

type CreateProjectRequest struct {
	// The request body parameters.
	Body *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectRequest) GoString() string {
	return s.String()
}

func (s *CreateProjectRequest) SetBody(v string) *CreateProjectRequest {
	s.Body = &v
	return s
}

type CreateProjectResponseBody struct {
	// The returned result.
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectResponseBody) GoString() string {
	return s.String()
}

func (s *CreateProjectResponseBody) SetData(v string) *CreateProjectResponseBody {
	s.Data = &v
	return s
}

func (s *CreateProjectResponseBody) SetRequestId(v string) *CreateProjectResponseBody {
	s.RequestId = &v
	return s
}

type CreateProjectResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectResponse) GoString() string {
	return s.String()
}

func (s *CreateProjectResponse) SetHeaders(v map[string]*string) *CreateProjectResponse {
	s.Headers = v
	return s
}

func (s *CreateProjectResponse) SetStatusCode(v int32) *CreateProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateProjectResponse) SetBody(v *CreateProjectResponseBody) *CreateProjectResponse {
	s.Body = v
	return s
}

type CreateQuotaPlanRequest struct {
	// The request body parameters.
	Body *string `json:"body,omitempty" xml:"body,omitempty"`
	// The ID of the region.
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// The ID of the tenant.
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s CreateQuotaPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateQuotaPlanRequest) GoString() string {
	return s.String()
}

func (s *CreateQuotaPlanRequest) SetBody(v string) *CreateQuotaPlanRequest {
	s.Body = &v
	return s
}

func (s *CreateQuotaPlanRequest) SetRegion(v string) *CreateQuotaPlanRequest {
	s.Region = &v
	return s
}

func (s *CreateQuotaPlanRequest) SetTenantId(v string) *CreateQuotaPlanRequest {
	s.TenantId = &v
	return s
}

type CreateQuotaPlanResponseBody struct {
	// The returned result.
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateQuotaPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateQuotaPlanResponseBody) GoString() string {
	return s.String()
}

func (s *CreateQuotaPlanResponseBody) SetData(v string) *CreateQuotaPlanResponseBody {
	s.Data = &v
	return s
}

func (s *CreateQuotaPlanResponseBody) SetRequestId(v string) *CreateQuotaPlanResponseBody {
	s.RequestId = &v
	return s
}

type CreateQuotaPlanResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateQuotaPlanResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateQuotaPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateQuotaPlanResponse) GoString() string {
	return s.String()
}

func (s *CreateQuotaPlanResponse) SetHeaders(v map[string]*string) *CreateQuotaPlanResponse {
	s.Headers = v
	return s
}

func (s *CreateQuotaPlanResponse) SetStatusCode(v int32) *CreateQuotaPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateQuotaPlanResponse) SetBody(v *CreateQuotaPlanResponseBody) *CreateQuotaPlanResponse {
	s.Body = v
	return s
}

type CreateQuotaScheduleRequest struct {
	// The request body parameters.
	Body *string `json:"body,omitempty" xml:"body,omitempty"`
	// The ID of the region.
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// The ID of the tenant.
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s CreateQuotaScheduleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateQuotaScheduleRequest) GoString() string {
	return s.String()
}

func (s *CreateQuotaScheduleRequest) SetBody(v string) *CreateQuotaScheduleRequest {
	s.Body = &v
	return s
}

func (s *CreateQuotaScheduleRequest) SetRegion(v string) *CreateQuotaScheduleRequest {
	s.Region = &v
	return s
}

func (s *CreateQuotaScheduleRequest) SetTenantId(v string) *CreateQuotaScheduleRequest {
	s.TenantId = &v
	return s
}

type CreateQuotaScheduleResponseBody struct {
	// The returned result.
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateQuotaScheduleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateQuotaScheduleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateQuotaScheduleResponseBody) SetData(v string) *CreateQuotaScheduleResponseBody {
	s.Data = &v
	return s
}

func (s *CreateQuotaScheduleResponseBody) SetRequestId(v string) *CreateQuotaScheduleResponseBody {
	s.RequestId = &v
	return s
}

type CreateQuotaScheduleResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateQuotaScheduleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateQuotaScheduleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateQuotaScheduleResponse) GoString() string {
	return s.String()
}

func (s *CreateQuotaScheduleResponse) SetHeaders(v map[string]*string) *CreateQuotaScheduleResponse {
	s.Headers = v
	return s
}

func (s *CreateQuotaScheduleResponse) SetStatusCode(v int32) *CreateQuotaScheduleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateQuotaScheduleResponse) SetBody(v *CreateQuotaScheduleResponseBody) *CreateQuotaScheduleResponse {
	s.Body = v
	return s
}

type CreateRoleRequest struct {
	// The request body parameters. For valid values, see [MaxCompute permissions](~~27935~~).
	Body *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRoleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRoleRequest) GoString() string {
	return s.String()
}

func (s *CreateRoleRequest) SetBody(v string) *CreateRoleRequest {
	s.Body = &v
	return s
}

type CreateRoleResponseBody struct {
	// The returned data.
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateRoleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateRoleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRoleResponseBody) SetData(v string) *CreateRoleResponseBody {
	s.Data = &v
	return s
}

func (s *CreateRoleResponseBody) SetRequestId(v string) *CreateRoleResponseBody {
	s.RequestId = &v
	return s
}

type CreateRoleResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateRoleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateRoleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateRoleResponse) GoString() string {
	return s.String()
}

func (s *CreateRoleResponse) SetHeaders(v map[string]*string) *CreateRoleResponse {
	s.Headers = v
	return s
}

func (s *CreateRoleResponse) SetStatusCode(v int32) *CreateRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRoleResponse) SetBody(v *CreateRoleResponseBody) *CreateRoleResponse {
	s.Body = v
	return s
}

type DeleteQuotaPlanRequest struct {
	// The ID of the region.
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// The ID of the tenant.
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s DeleteQuotaPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteQuotaPlanRequest) GoString() string {
	return s.String()
}

func (s *DeleteQuotaPlanRequest) SetRegion(v string) *DeleteQuotaPlanRequest {
	s.Region = &v
	return s
}

func (s *DeleteQuotaPlanRequest) SetTenantId(v string) *DeleteQuotaPlanRequest {
	s.TenantId = &v
	return s
}

type DeleteQuotaPlanResponseBody struct {
	// The returned result.
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteQuotaPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteQuotaPlanResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteQuotaPlanResponseBody) SetData(v string) *DeleteQuotaPlanResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteQuotaPlanResponseBody) SetRequestId(v string) *DeleteQuotaPlanResponseBody {
	s.RequestId = &v
	return s
}

type DeleteQuotaPlanResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteQuotaPlanResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteQuotaPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteQuotaPlanResponse) GoString() string {
	return s.String()
}

func (s *DeleteQuotaPlanResponse) SetHeaders(v map[string]*string) *DeleteQuotaPlanResponse {
	s.Headers = v
	return s
}

func (s *DeleteQuotaPlanResponse) SetStatusCode(v int32) *DeleteQuotaPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteQuotaPlanResponse) SetBody(v *DeleteQuotaPlanResponseBody) *DeleteQuotaPlanResponse {
	s.Body = v
	return s
}

type GetJobResourceUsageRequest struct {
	// The date that is accurate to the day part for the query. The date must be in the yyyy-MM-dd format.
	Date *string `json:"date,omitempty" xml:"date,omitempty"`
	// The list of job executors.
	JobOwnerList []*string `json:"jobOwnerList,omitempty" xml:"jobOwnerList,omitempty" type:"Repeated"`
	// The page number.
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries per page. Default value: 10. Maximum value: 100.
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The list of nicknames of quotas that are used by jobs.
	QuotaNicknameList []*string `json:"quotaNicknameList,omitempty" xml:"quotaNicknameList,omitempty" type:"Repeated"`
}

func (s GetJobResourceUsageRequest) String() string {
	return tea.Prettify(s)
}

func (s GetJobResourceUsageRequest) GoString() string {
	return s.String()
}

func (s *GetJobResourceUsageRequest) SetDate(v string) *GetJobResourceUsageRequest {
	s.Date = &v
	return s
}

func (s *GetJobResourceUsageRequest) SetJobOwnerList(v []*string) *GetJobResourceUsageRequest {
	s.JobOwnerList = v
	return s
}

func (s *GetJobResourceUsageRequest) SetPageNumber(v int64) *GetJobResourceUsageRequest {
	s.PageNumber = &v
	return s
}

func (s *GetJobResourceUsageRequest) SetPageSize(v int64) *GetJobResourceUsageRequest {
	s.PageSize = &v
	return s
}

func (s *GetJobResourceUsageRequest) SetQuotaNicknameList(v []*string) *GetJobResourceUsageRequest {
	s.QuotaNicknameList = v
	return s
}

type GetJobResourceUsageShrinkRequest struct {
	// The date that is accurate to the day part for the query. The date must be in the yyyy-MM-dd format.
	Date *string `json:"date,omitempty" xml:"date,omitempty"`
	// The list of job executors.
	JobOwnerListShrink *string `json:"jobOwnerList,omitempty" xml:"jobOwnerList,omitempty"`
	// The page number.
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries per page. Default value: 10. Maximum value: 100.
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The list of nicknames of quotas that are used by jobs.
	QuotaNicknameListShrink *string `json:"quotaNicknameList,omitempty" xml:"quotaNicknameList,omitempty"`
}

func (s GetJobResourceUsageShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetJobResourceUsageShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetJobResourceUsageShrinkRequest) SetDate(v string) *GetJobResourceUsageShrinkRequest {
	s.Date = &v
	return s
}

func (s *GetJobResourceUsageShrinkRequest) SetJobOwnerListShrink(v string) *GetJobResourceUsageShrinkRequest {
	s.JobOwnerListShrink = &v
	return s
}

func (s *GetJobResourceUsageShrinkRequest) SetPageNumber(v int64) *GetJobResourceUsageShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *GetJobResourceUsageShrinkRequest) SetPageSize(v int64) *GetJobResourceUsageShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *GetJobResourceUsageShrinkRequest) SetQuotaNicknameListShrink(v string) *GetJobResourceUsageShrinkRequest {
	s.QuotaNicknameListShrink = &v
	return s
}

type GetJobResourceUsageResponseBody struct {
	// The data returned.
	Data *GetJobResourceUsageResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The error code returned if the request failed.
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// The error message returned if the request failed.
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// Indicates whether the request was successful. If this parameter was not empty and the value of this parameter was not 200, the request failed.
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The request ID.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetJobResourceUsageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetJobResourceUsageResponseBody) GoString() string {
	return s.String()
}

func (s *GetJobResourceUsageResponseBody) SetData(v *GetJobResourceUsageResponseBodyData) *GetJobResourceUsageResponseBody {
	s.Data = v
	return s
}

func (s *GetJobResourceUsageResponseBody) SetErrorCode(v string) *GetJobResourceUsageResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetJobResourceUsageResponseBody) SetErrorMsg(v string) *GetJobResourceUsageResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GetJobResourceUsageResponseBody) SetHttpCode(v int32) *GetJobResourceUsageResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetJobResourceUsageResponseBody) SetRequestId(v string) *GetJobResourceUsageResponseBody {
	s.RequestId = &v
	return s
}

type GetJobResourceUsageResponseBodyData struct {
	// The data list returned.
	JobResourceUsageList []*GetJobResourceUsageResponseBodyDataJobResourceUsageList `json:"jobResourceUsageList,omitempty" xml:"jobResourceUsageList,omitempty" type:"Repeated"`
	// The page number.
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries per page.
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The total number of returned entries.
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s GetJobResourceUsageResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetJobResourceUsageResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetJobResourceUsageResponseBodyData) SetJobResourceUsageList(v []*GetJobResourceUsageResponseBodyDataJobResourceUsageList) *GetJobResourceUsageResponseBodyData {
	s.JobResourceUsageList = v
	return s
}

func (s *GetJobResourceUsageResponseBodyData) SetPageNumber(v int64) *GetJobResourceUsageResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *GetJobResourceUsageResponseBodyData) SetPageSize(v int64) *GetJobResourceUsageResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetJobResourceUsageResponseBodyData) SetTotalCount(v int64) *GetJobResourceUsageResponseBodyData {
	s.TotalCount = &v
	return s
}

type GetJobResourceUsageResponseBodyDataJobResourceUsageList struct {
	// The total number of used compute units (CUs).
	CuUsage *int64 `json:"cuUsage,omitempty" xml:"cuUsage,omitempty"`
	// The start date of the query in the format of yyyy-MM-dd.
	Date *string `json:"date,omitempty" xml:"date,omitempty"`
	// The job executor.
	JobOwner *string `json:"jobOwner,omitempty" xml:"jobOwner,omitempty"`
	// The total memory usage.
	MemoryUsage *int64 `json:"memoryUsage,omitempty" xml:"memoryUsage,omitempty"`
	// The quota nickname.
	QuotaNickname *string `json:"quotaNickname,omitempty" xml:"quotaNickname,omitempty"`
}

func (s GetJobResourceUsageResponseBodyDataJobResourceUsageList) String() string {
	return tea.Prettify(s)
}

func (s GetJobResourceUsageResponseBodyDataJobResourceUsageList) GoString() string {
	return s.String()
}

func (s *GetJobResourceUsageResponseBodyDataJobResourceUsageList) SetCuUsage(v int64) *GetJobResourceUsageResponseBodyDataJobResourceUsageList {
	s.CuUsage = &v
	return s
}

func (s *GetJobResourceUsageResponseBodyDataJobResourceUsageList) SetDate(v string) *GetJobResourceUsageResponseBodyDataJobResourceUsageList {
	s.Date = &v
	return s
}

func (s *GetJobResourceUsageResponseBodyDataJobResourceUsageList) SetJobOwner(v string) *GetJobResourceUsageResponseBodyDataJobResourceUsageList {
	s.JobOwner = &v
	return s
}

func (s *GetJobResourceUsageResponseBodyDataJobResourceUsageList) SetMemoryUsage(v int64) *GetJobResourceUsageResponseBodyDataJobResourceUsageList {
	s.MemoryUsage = &v
	return s
}

func (s *GetJobResourceUsageResponseBodyDataJobResourceUsageList) SetQuotaNickname(v string) *GetJobResourceUsageResponseBodyDataJobResourceUsageList {
	s.QuotaNickname = &v
	return s
}

type GetJobResourceUsageResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetJobResourceUsageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetJobResourceUsageResponse) String() string {
	return tea.Prettify(s)
}

func (s GetJobResourceUsageResponse) GoString() string {
	return s.String()
}

func (s *GetJobResourceUsageResponse) SetHeaders(v map[string]*string) *GetJobResourceUsageResponse {
	s.Headers = v
	return s
}

func (s *GetJobResourceUsageResponse) SetStatusCode(v int32) *GetJobResourceUsageResponse {
	s.StatusCode = &v
	return s
}

func (s *GetJobResourceUsageResponse) SetBody(v *GetJobResourceUsageResponseBody) *GetJobResourceUsageResponse {
	s.Body = v
	return s
}

type GetPackageRequest struct {
	// The project to which the package belongs. This parameter is required if the package is installed in the MaxCompute project.
	SourceProject *string `json:"sourceProject,omitempty" xml:"sourceProject,omitempty"`
}

func (s GetPackageRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPackageRequest) GoString() string {
	return s.String()
}

func (s *GetPackageRequest) SetSourceProject(v string) *GetPackageRequest {
	s.SourceProject = &v
	return s
}

type GetPackageResponseBody struct {
	// The returned data.
	Data      *GetPackageResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	ErrorCode *string                     `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMsg  *string                     `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	HttpCode  *int32                      `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetPackageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPackageResponseBody) GoString() string {
	return s.String()
}

func (s *GetPackageResponseBody) SetData(v *GetPackageResponseBodyData) *GetPackageResponseBody {
	s.Data = v
	return s
}

func (s *GetPackageResponseBody) SetErrorCode(v string) *GetPackageResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetPackageResponseBody) SetErrorMsg(v string) *GetPackageResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GetPackageResponseBody) SetHttpCode(v int32) *GetPackageResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetPackageResponseBody) SetRequestId(v string) *GetPackageResponseBody {
	s.RequestId = &v
	return s
}

type GetPackageResponseBodyData struct {
	// The projects in which the package is installed.
	AllowedProjectList []*GetPackageResponseBodyDataAllowedProjectList `json:"allowedProjectList,omitempty" xml:"allowedProjectList,omitempty" type:"Repeated"`
	// The details of the resources that are included in the package.
	ResourceList *GetPackageResponseBodyDataResourceList `json:"resourceList,omitempty" xml:"resourceList,omitempty" type:"Struct"`
}

func (s GetPackageResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetPackageResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetPackageResponseBodyData) SetAllowedProjectList(v []*GetPackageResponseBodyDataAllowedProjectList) *GetPackageResponseBodyData {
	s.AllowedProjectList = v
	return s
}

func (s *GetPackageResponseBodyData) SetResourceList(v *GetPackageResponseBodyDataResourceList) *GetPackageResponseBodyData {
	s.ResourceList = v
	return s
}

type GetPackageResponseBodyDataAllowedProjectList struct {
	// The security level for sensitive data.
	Label *string `json:"label,omitempty" xml:"label,omitempty"`
	// The name of the MaxCompute project.
	Project *string `json:"project,omitempty" xml:"project,omitempty"`
}

func (s GetPackageResponseBodyDataAllowedProjectList) String() string {
	return tea.Prettify(s)
}

func (s GetPackageResponseBodyDataAllowedProjectList) GoString() string {
	return s.String()
}

func (s *GetPackageResponseBodyDataAllowedProjectList) SetLabel(v string) *GetPackageResponseBodyDataAllowedProjectList {
	s.Label = &v
	return s
}

func (s *GetPackageResponseBodyDataAllowedProjectList) SetProject(v string) *GetPackageResponseBodyDataAllowedProjectList {
	s.Project = &v
	return s
}

type GetPackageResponseBodyDataResourceList struct {
	// The functions.
	Function []*GetPackageResponseBodyDataResourceListFunction `json:"function,omitempty" xml:"function,omitempty" type:"Repeated"`
	// The resources.
	Resource []*GetPackageResponseBodyDataResourceListResource `json:"resource,omitempty" xml:"resource,omitempty" type:"Repeated"`
	// The tables.
	Table []*GetPackageResponseBodyDataResourceListTable `json:"table,omitempty" xml:"table,omitempty" type:"Repeated"`
}

func (s GetPackageResponseBodyDataResourceList) String() string {
	return tea.Prettify(s)
}

func (s GetPackageResponseBodyDataResourceList) GoString() string {
	return s.String()
}

func (s *GetPackageResponseBodyDataResourceList) SetFunction(v []*GetPackageResponseBodyDataResourceListFunction) *GetPackageResponseBodyDataResourceList {
	s.Function = v
	return s
}

func (s *GetPackageResponseBodyDataResourceList) SetResource(v []*GetPackageResponseBodyDataResourceListResource) *GetPackageResponseBodyDataResourceList {
	s.Resource = v
	return s
}

func (s *GetPackageResponseBodyDataResourceList) SetTable(v []*GetPackageResponseBodyDataResourceListTable) *GetPackageResponseBodyDataResourceList {
	s.Table = v
	return s
}

type GetPackageResponseBodyDataResourceListFunction struct {
	// The operations that were performed on the function.
	Actions []*string `json:"actions,omitempty" xml:"actions,omitempty" type:"Repeated"`
	// The name of the function.
	Name       *string `json:"name,omitempty" xml:"name,omitempty"`
	SchemaName *string `json:"schemaName,omitempty" xml:"schemaName,omitempty"`
}

func (s GetPackageResponseBodyDataResourceListFunction) String() string {
	return tea.Prettify(s)
}

func (s GetPackageResponseBodyDataResourceListFunction) GoString() string {
	return s.String()
}

func (s *GetPackageResponseBodyDataResourceListFunction) SetActions(v []*string) *GetPackageResponseBodyDataResourceListFunction {
	s.Actions = v
	return s
}

func (s *GetPackageResponseBodyDataResourceListFunction) SetName(v string) *GetPackageResponseBodyDataResourceListFunction {
	s.Name = &v
	return s
}

func (s *GetPackageResponseBodyDataResourceListFunction) SetSchemaName(v string) *GetPackageResponseBodyDataResourceListFunction {
	s.SchemaName = &v
	return s
}

type GetPackageResponseBodyDataResourceListResource struct {
	// The operations that were performed on the resource.
	Actions []*string `json:"actions,omitempty" xml:"actions,omitempty" type:"Repeated"`
	// The name of the resource.
	Name       *string `json:"name,omitempty" xml:"name,omitempty"`
	SchemaName *string `json:"schemaName,omitempty" xml:"schemaName,omitempty"`
}

func (s GetPackageResponseBodyDataResourceListResource) String() string {
	return tea.Prettify(s)
}

func (s GetPackageResponseBodyDataResourceListResource) GoString() string {
	return s.String()
}

func (s *GetPackageResponseBodyDataResourceListResource) SetActions(v []*string) *GetPackageResponseBodyDataResourceListResource {
	s.Actions = v
	return s
}

func (s *GetPackageResponseBodyDataResourceListResource) SetName(v string) *GetPackageResponseBodyDataResourceListResource {
	s.Name = &v
	return s
}

func (s *GetPackageResponseBodyDataResourceListResource) SetSchemaName(v string) *GetPackageResponseBodyDataResourceListResource {
	s.SchemaName = &v
	return s
}

type GetPackageResponseBodyDataResourceListTable struct {
	// The operations that were performed on the table.
	Actions []*string `json:"actions,omitempty" xml:"actions,omitempty" type:"Repeated"`
	// The name of the table.
	Name       *string `json:"name,omitempty" xml:"name,omitempty"`
	SchemaName *string `json:"schemaName,omitempty" xml:"schemaName,omitempty"`
}

func (s GetPackageResponseBodyDataResourceListTable) String() string {
	return tea.Prettify(s)
}

func (s GetPackageResponseBodyDataResourceListTable) GoString() string {
	return s.String()
}

func (s *GetPackageResponseBodyDataResourceListTable) SetActions(v []*string) *GetPackageResponseBodyDataResourceListTable {
	s.Actions = v
	return s
}

func (s *GetPackageResponseBodyDataResourceListTable) SetName(v string) *GetPackageResponseBodyDataResourceListTable {
	s.Name = &v
	return s
}

func (s *GetPackageResponseBodyDataResourceListTable) SetSchemaName(v string) *GetPackageResponseBodyDataResourceListTable {
	s.SchemaName = &v
	return s
}

type GetPackageResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetPackageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetPackageResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPackageResponse) GoString() string {
	return s.String()
}

func (s *GetPackageResponse) SetHeaders(v map[string]*string) *GetPackageResponse {
	s.Headers = v
	return s
}

func (s *GetPackageResponse) SetStatusCode(v int32) *GetPackageResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPackageResponse) SetBody(v *GetPackageResponseBody) *GetPackageResponse {
	s.Body = v
	return s
}

type GetProjectResponseBody struct {
	// The returned data.
	Data *GetProjectResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Indicates whether the request was successful. If this parameter was not empty and the value of this parameter was not 200, the request failed.
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetProjectResponseBody) GoString() string {
	return s.String()
}

func (s *GetProjectResponseBody) SetData(v *GetProjectResponseBodyData) *GetProjectResponseBody {
	s.Data = v
	return s
}

func (s *GetProjectResponseBody) SetHttpCode(v int32) *GetProjectResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetProjectResponseBody) SetRequestId(v string) *GetProjectResponseBody {
	s.RequestId = &v
	return s
}

type GetProjectResponseBodyData struct {
	// The comment of the project.
	Comment *string `json:"comment,omitempty" xml:"comment,omitempty"`
	// The storage usage.
	CostStorage *string `json:"costStorage,omitempty" xml:"costStorage,omitempty"`
	// The default computing quota.
	DefaultQuota *string `json:"defaultQuota,omitempty" xml:"defaultQuota,omitempty"`
	// The IP address whitelist.
	IpWhiteList *GetProjectResponseBodyDataIpWhiteList `json:"ipWhiteList,omitempty" xml:"ipWhiteList,omitempty" type:"Struct"`
	// The name of the MaxCompute project.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The owner of the project.
	Owner *string `json:"owner,omitempty" xml:"owner,omitempty"`
	// The billing method of the project.
	ProductType *string `json:"productType,omitempty" xml:"productType,omitempty"`
	// The properties of the project.
	Properties *GetProjectResponseBodyDataProperties `json:"properties,omitempty" xml:"properties,omitempty" type:"Struct"`
	// The tag.
	SaleTag *GetProjectResponseBodyDataSaleTag `json:"saleTag,omitempty" xml:"saleTag,omitempty" type:"Struct"`
	// The permission properties.
	SecurityProperties *GetProjectResponseBodyDataSecurityProperties `json:"securityProperties,omitempty" xml:"securityProperties,omitempty" type:"Struct"`
	// The status of the project. Valid values: -**AVAILABLE**: The project was available. -**READONLY**: The project was read only. -**FROZEN**: The project was frozen. -**DELETING**: The project was being deleted.
	Status      *string   `json:"status,omitempty" xml:"status,omitempty"`
	SuperAdmins []*string `json:"superAdmins,omitempty" xml:"superAdmins,omitempty" type:"Repeated"`
	// The type of the project. Valid values: -**managed**: The project is an internal project. -**external**: The project is an external project.
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetProjectResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetProjectResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetProjectResponseBodyData) SetComment(v string) *GetProjectResponseBodyData {
	s.Comment = &v
	return s
}

func (s *GetProjectResponseBodyData) SetCostStorage(v string) *GetProjectResponseBodyData {
	s.CostStorage = &v
	return s
}

func (s *GetProjectResponseBodyData) SetDefaultQuota(v string) *GetProjectResponseBodyData {
	s.DefaultQuota = &v
	return s
}

func (s *GetProjectResponseBodyData) SetIpWhiteList(v *GetProjectResponseBodyDataIpWhiteList) *GetProjectResponseBodyData {
	s.IpWhiteList = v
	return s
}

func (s *GetProjectResponseBodyData) SetName(v string) *GetProjectResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetProjectResponseBodyData) SetOwner(v string) *GetProjectResponseBodyData {
	s.Owner = &v
	return s
}

func (s *GetProjectResponseBodyData) SetProductType(v string) *GetProjectResponseBodyData {
	s.ProductType = &v
	return s
}

func (s *GetProjectResponseBodyData) SetProperties(v *GetProjectResponseBodyDataProperties) *GetProjectResponseBodyData {
	s.Properties = v
	return s
}

func (s *GetProjectResponseBodyData) SetSaleTag(v *GetProjectResponseBodyDataSaleTag) *GetProjectResponseBodyData {
	s.SaleTag = v
	return s
}

func (s *GetProjectResponseBodyData) SetSecurityProperties(v *GetProjectResponseBodyDataSecurityProperties) *GetProjectResponseBodyData {
	s.SecurityProperties = v
	return s
}

func (s *GetProjectResponseBodyData) SetStatus(v string) *GetProjectResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetProjectResponseBodyData) SetSuperAdmins(v []*string) *GetProjectResponseBodyData {
	s.SuperAdmins = v
	return s
}

func (s *GetProjectResponseBodyData) SetType(v string) *GetProjectResponseBodyData {
	s.Type = &v
	return s
}

type GetProjectResponseBodyDataIpWhiteList struct {
	// The list of IP addresses.
	IpList *string `json:"ipList,omitempty" xml:"ipList,omitempty"`
	// The list of virtual private cloud (VPC) IP addresses.
	VpcIpList *string `json:"vpcIpList,omitempty" xml:"vpcIpList,omitempty"`
}

func (s GetProjectResponseBodyDataIpWhiteList) String() string {
	return tea.Prettify(s)
}

func (s GetProjectResponseBodyDataIpWhiteList) GoString() string {
	return s.String()
}

func (s *GetProjectResponseBodyDataIpWhiteList) SetIpList(v string) *GetProjectResponseBodyDataIpWhiteList {
	s.IpList = &v
	return s
}

func (s *GetProjectResponseBodyDataIpWhiteList) SetVpcIpList(v string) *GetProjectResponseBodyDataIpWhiteList {
	s.VpcIpList = &v
	return s
}

type GetProjectResponseBodyDataProperties struct {
	// Indicates whether a full table scan on the project is enabled.
	AllowFullScan *bool `json:"allowFullScan,omitempty" xml:"allowFullScan,omitempty"`
	// Indicates whether the DECIMAL data type in MaxCompute V2.0 is enabled.
	EnableDecimal2 *bool `json:"enableDecimal2,omitempty" xml:"enableDecimal2,omitempty"`
	// Indicates whether tunnel quota routing is enabled.
	EnableTunnelQuotaRoute *bool `json:"enableTunnelQuotaRoute,omitempty" xml:"enableTunnelQuotaRoute,omitempty"`
	// The encryption information.
	Encryption *GetProjectResponseBodyDataPropertiesEncryption `json:"encryption,omitempty" xml:"encryption,omitempty" type:"Struct"`
	// The number of days for which backup data can be retained.
	RetentionDays *int64 `json:"retentionDays,omitempty" xml:"retentionDays,omitempty"`
	// The upper limit for the resources that are consumed by an SQL statement.
	SqlMeteringMax *string `json:"sqlMeteringMax,omitempty" xml:"sqlMeteringMax,omitempty"`
	// The information about the tiered storage.
	StorageTierInfo *GetProjectResponseBodyDataPropertiesStorageTierInfo `json:"storageTierInfo,omitempty" xml:"storageTierInfo,omitempty" type:"Struct"`
	// The lifecycle of the table in the project.
	TableLifecycle *GetProjectResponseBodyDataPropertiesTableLifecycle `json:"tableLifecycle,omitempty" xml:"tableLifecycle,omitempty" type:"Struct"`
	// The time zone of the project.
	Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
	// The name of the tunnel quota.
	TunnelQuota *string `json:"tunnelQuota,omitempty" xml:"tunnelQuota,omitempty"`
	// The data type edition. Valid values: -**1**: MaxCompute V1.0 data type edition. -**2**: MaxCompute V2.0 data type edition. -**hive**: Hive-compatible data type edition.
	TypeSystem *string `json:"typeSystem,omitempty" xml:"typeSystem,omitempty"`
}

func (s GetProjectResponseBodyDataProperties) String() string {
	return tea.Prettify(s)
}

func (s GetProjectResponseBodyDataProperties) GoString() string {
	return s.String()
}

func (s *GetProjectResponseBodyDataProperties) SetAllowFullScan(v bool) *GetProjectResponseBodyDataProperties {
	s.AllowFullScan = &v
	return s
}

func (s *GetProjectResponseBodyDataProperties) SetEnableDecimal2(v bool) *GetProjectResponseBodyDataProperties {
	s.EnableDecimal2 = &v
	return s
}

func (s *GetProjectResponseBodyDataProperties) SetEnableTunnelQuotaRoute(v bool) *GetProjectResponseBodyDataProperties {
	s.EnableTunnelQuotaRoute = &v
	return s
}

func (s *GetProjectResponseBodyDataProperties) SetEncryption(v *GetProjectResponseBodyDataPropertiesEncryption) *GetProjectResponseBodyDataProperties {
	s.Encryption = v
	return s
}

func (s *GetProjectResponseBodyDataProperties) SetRetentionDays(v int64) *GetProjectResponseBodyDataProperties {
	s.RetentionDays = &v
	return s
}

func (s *GetProjectResponseBodyDataProperties) SetSqlMeteringMax(v string) *GetProjectResponseBodyDataProperties {
	s.SqlMeteringMax = &v
	return s
}

func (s *GetProjectResponseBodyDataProperties) SetStorageTierInfo(v *GetProjectResponseBodyDataPropertiesStorageTierInfo) *GetProjectResponseBodyDataProperties {
	s.StorageTierInfo = v
	return s
}

func (s *GetProjectResponseBodyDataProperties) SetTableLifecycle(v *GetProjectResponseBodyDataPropertiesTableLifecycle) *GetProjectResponseBodyDataProperties {
	s.TableLifecycle = v
	return s
}

func (s *GetProjectResponseBodyDataProperties) SetTimezone(v string) *GetProjectResponseBodyDataProperties {
	s.Timezone = &v
	return s
}

func (s *GetProjectResponseBodyDataProperties) SetTunnelQuota(v string) *GetProjectResponseBodyDataProperties {
	s.TunnelQuota = &v
	return s
}

func (s *GetProjectResponseBodyDataProperties) SetTypeSystem(v string) *GetProjectResponseBodyDataProperties {
	s.TypeSystem = &v
	return s
}

type GetProjectResponseBodyDataPropertiesEncryption struct {
	// The name of the encryption algorithm.
	Algorithm *string `json:"algorithm,omitempty" xml:"algorithm,omitempty"`
	// Indicates whether data encryption is enabled. Valid values: true and false.
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// The key of the encryption algorithm.
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
}

func (s GetProjectResponseBodyDataPropertiesEncryption) String() string {
	return tea.Prettify(s)
}

func (s GetProjectResponseBodyDataPropertiesEncryption) GoString() string {
	return s.String()
}

func (s *GetProjectResponseBodyDataPropertiesEncryption) SetAlgorithm(v string) *GetProjectResponseBodyDataPropertiesEncryption {
	s.Algorithm = &v
	return s
}

func (s *GetProjectResponseBodyDataPropertiesEncryption) SetEnable(v bool) *GetProjectResponseBodyDataPropertiesEncryption {
	s.Enable = &v
	return s
}

func (s *GetProjectResponseBodyDataPropertiesEncryption) SetKey(v string) *GetProjectResponseBodyDataPropertiesEncryption {
	s.Key = &v
	return s
}

type GetProjectResponseBodyDataPropertiesStorageTierInfo struct {
	// The backup storage.
	ProjectBackupSize *int64 `json:"projectBackupSize,omitempty" xml:"projectBackupSize,omitempty"`
	// The tiered storage.
	StorageTierSize *GetProjectResponseBodyDataPropertiesStorageTierInfoStorageTierSize `json:"storageTierSize,omitempty" xml:"storageTierSize,omitempty" type:"Struct"`
}

func (s GetProjectResponseBodyDataPropertiesStorageTierInfo) String() string {
	return tea.Prettify(s)
}

func (s GetProjectResponseBodyDataPropertiesStorageTierInfo) GoString() string {
	return s.String()
}

func (s *GetProjectResponseBodyDataPropertiesStorageTierInfo) SetProjectBackupSize(v int64) *GetProjectResponseBodyDataPropertiesStorageTierInfo {
	s.ProjectBackupSize = &v
	return s
}

func (s *GetProjectResponseBodyDataPropertiesStorageTierInfo) SetStorageTierSize(v *GetProjectResponseBodyDataPropertiesStorageTierInfoStorageTierSize) *GetProjectResponseBodyDataPropertiesStorageTierInfo {
	s.StorageTierSize = v
	return s
}

type GetProjectResponseBodyDataPropertiesStorageTierInfoStorageTierSize struct {
	// The long-term storage.
	LongTermSize *int64 `json:"longTermSize,omitempty" xml:"longTermSize,omitempty"`
	// The IA storage.
	LowFrequencySize *int64 `json:"lowFrequencySize,omitempty" xml:"lowFrequencySize,omitempty"`
	// The standard storage.
	StandardSize *int64 `json:"standardSize,omitempty" xml:"standardSize,omitempty"`
}

func (s GetProjectResponseBodyDataPropertiesStorageTierInfoStorageTierSize) String() string {
	return tea.Prettify(s)
}

func (s GetProjectResponseBodyDataPropertiesStorageTierInfoStorageTierSize) GoString() string {
	return s.String()
}

func (s *GetProjectResponseBodyDataPropertiesStorageTierInfoStorageTierSize) SetLongTermSize(v int64) *GetProjectResponseBodyDataPropertiesStorageTierInfoStorageTierSize {
	s.LongTermSize = &v
	return s
}

func (s *GetProjectResponseBodyDataPropertiesStorageTierInfoStorageTierSize) SetLowFrequencySize(v int64) *GetProjectResponseBodyDataPropertiesStorageTierInfoStorageTierSize {
	s.LowFrequencySize = &v
	return s
}

func (s *GetProjectResponseBodyDataPropertiesStorageTierInfoStorageTierSize) SetStandardSize(v int64) *GetProjectResponseBodyDataPropertiesStorageTierInfoStorageTierSize {
	s.StandardSize = &v
	return s
}

type GetProjectResponseBodyDataPropertiesTableLifecycle struct {
	// The type of the lifecycle. Valid values: -**mandatory**: The lifecycle clause is required. You must configure a lifecycle for a table. -**optional**: The lifecycle clause is optional in a table creation statement. If you do not configure a lifecycle for a table, the table does not expire. -**inherit**: If you do not configure a lifecycle for a table when you create the table, the value of the odps.table.lifecycle.value parameter is used by default.
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The retention period of a table. Unit: days.
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s GetProjectResponseBodyDataPropertiesTableLifecycle) String() string {
	return tea.Prettify(s)
}

func (s GetProjectResponseBodyDataPropertiesTableLifecycle) GoString() string {
	return s.String()
}

func (s *GetProjectResponseBodyDataPropertiesTableLifecycle) SetType(v string) *GetProjectResponseBodyDataPropertiesTableLifecycle {
	s.Type = &v
	return s
}

func (s *GetProjectResponseBodyDataPropertiesTableLifecycle) SetValue(v string) *GetProjectResponseBodyDataPropertiesTableLifecycle {
	s.Value = &v
	return s
}

type GetProjectResponseBodyDataSaleTag struct {
	// The ID of the resource.
	ResourceId *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	// The type of the resource.
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
}

func (s GetProjectResponseBodyDataSaleTag) String() string {
	return tea.Prettify(s)
}

func (s GetProjectResponseBodyDataSaleTag) GoString() string {
	return s.String()
}

func (s *GetProjectResponseBodyDataSaleTag) SetResourceId(v string) *GetProjectResponseBodyDataSaleTag {
	s.ResourceId = &v
	return s
}

func (s *GetProjectResponseBodyDataSaleTag) SetResourceType(v string) *GetProjectResponseBodyDataSaleTag {
	s.ResourceType = &v
	return s
}

type GetProjectResponseBodyDataSecurityProperties struct {
	// Indicates whether Download control is enabled.
	EnableDownloadPrivilege *bool `json:"enableDownloadPrivilege,omitempty" xml:"enableDownloadPrivilege,omitempty"`
	// Indicates whether label-based access control is enabled.
	LabelSecurity *bool `json:"labelSecurity,omitempty" xml:"labelSecurity,omitempty"`
	// Indicates whether the object creator is allowed to perform operations on objects.
	ObjectCreatorHasAccessPermission *bool `json:"objectCreatorHasAccessPermission,omitempty" xml:"objectCreatorHasAccessPermission,omitempty"`
	// Indicates whether the object creator is allowed to authorize other users to perform operations on objects.
	ObjectCreatorHasGrantPermission *bool `json:"objectCreatorHasGrantPermission,omitempty" xml:"objectCreatorHasGrantPermission,omitempty"`
	// Indicates whether project data protection is enabled.
	ProjectProtection *GetProjectResponseBodyDataSecurityPropertiesProjectProtection `json:"projectProtection,omitempty" xml:"projectProtection,omitempty" type:"Struct"`
	// Indicates whether ACL-based access control is enabled.
	UsingAcl *bool `json:"usingAcl,omitempty" xml:"usingAcl,omitempty"`
	// Indicates whether policy-based access control is enabled.
	UsingPolicy *bool `json:"usingPolicy,omitempty" xml:"usingPolicy,omitempty"`
}

func (s GetProjectResponseBodyDataSecurityProperties) String() string {
	return tea.Prettify(s)
}

func (s GetProjectResponseBodyDataSecurityProperties) GoString() string {
	return s.String()
}

func (s *GetProjectResponseBodyDataSecurityProperties) SetEnableDownloadPrivilege(v bool) *GetProjectResponseBodyDataSecurityProperties {
	s.EnableDownloadPrivilege = &v
	return s
}

func (s *GetProjectResponseBodyDataSecurityProperties) SetLabelSecurity(v bool) *GetProjectResponseBodyDataSecurityProperties {
	s.LabelSecurity = &v
	return s
}

func (s *GetProjectResponseBodyDataSecurityProperties) SetObjectCreatorHasAccessPermission(v bool) *GetProjectResponseBodyDataSecurityProperties {
	s.ObjectCreatorHasAccessPermission = &v
	return s
}

func (s *GetProjectResponseBodyDataSecurityProperties) SetObjectCreatorHasGrantPermission(v bool) *GetProjectResponseBodyDataSecurityProperties {
	s.ObjectCreatorHasGrantPermission = &v
	return s
}

func (s *GetProjectResponseBodyDataSecurityProperties) SetProjectProtection(v *GetProjectResponseBodyDataSecurityPropertiesProjectProtection) *GetProjectResponseBodyDataSecurityProperties {
	s.ProjectProtection = v
	return s
}

func (s *GetProjectResponseBodyDataSecurityProperties) SetUsingAcl(v bool) *GetProjectResponseBodyDataSecurityProperties {
	s.UsingAcl = &v
	return s
}

func (s *GetProjectResponseBodyDataSecurityProperties) SetUsingPolicy(v bool) *GetProjectResponseBodyDataSecurityProperties {
	s.UsingPolicy = &v
	return s
}

type GetProjectResponseBodyDataSecurityPropertiesProjectProtection struct {
	// The exception policy. If cross-project data access operations are required, the project owner must configure an exception policy in advance to allow the specified user to transfer data of a specified object from the current project to a specified project. After the exception policy is configured, data of the object can be transferred to the specified project even if the project data protection feature is enabled.
	ExceptionPolicy *string `json:"exceptionPolicy,omitempty" xml:"exceptionPolicy,omitempty"`
	// Indicates whether project data protection is enabled.
	Protected *bool `json:"protected,omitempty" xml:"protected,omitempty"`
}

func (s GetProjectResponseBodyDataSecurityPropertiesProjectProtection) String() string {
	return tea.Prettify(s)
}

func (s GetProjectResponseBodyDataSecurityPropertiesProjectProtection) GoString() string {
	return s.String()
}

func (s *GetProjectResponseBodyDataSecurityPropertiesProjectProtection) SetExceptionPolicy(v string) *GetProjectResponseBodyDataSecurityPropertiesProjectProtection {
	s.ExceptionPolicy = &v
	return s
}

func (s *GetProjectResponseBodyDataSecurityPropertiesProjectProtection) SetProtected(v bool) *GetProjectResponseBodyDataSecurityPropertiesProjectProtection {
	s.Protected = &v
	return s
}

type GetProjectResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s GetProjectResponse) GoString() string {
	return s.String()
}

func (s *GetProjectResponse) SetHeaders(v map[string]*string) *GetProjectResponse {
	s.Headers = v
	return s
}

func (s *GetProjectResponse) SetStatusCode(v int32) *GetProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *GetProjectResponse) SetBody(v *GetProjectResponseBody) *GetProjectResponse {
	s.Body = v
	return s
}

type GetQuotaRequest struct {
	AkProven *string `json:"AkProven,omitempty" xml:"AkProven,omitempty"`
	Mock     *bool   `json:"mock,omitempty" xml:"mock,omitempty"`
	Region   *string `json:"region,omitempty" xml:"region,omitempty"`
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s GetQuotaRequest) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaRequest) GoString() string {
	return s.String()
}

func (s *GetQuotaRequest) SetAkProven(v string) *GetQuotaRequest {
	s.AkProven = &v
	return s
}

func (s *GetQuotaRequest) SetMock(v bool) *GetQuotaRequest {
	s.Mock = &v
	return s
}

func (s *GetQuotaRequest) SetRegion(v string) *GetQuotaRequest {
	s.Region = &v
	return s
}

func (s *GetQuotaRequest) SetTenantId(v string) *GetQuotaRequest {
	s.TenantId = &v
	return s
}

type GetQuotaResponseBody struct {
	BillingPolicy    *GetQuotaResponseBodyBillingPolicy      `json:"billingPolicy,omitempty" xml:"billingPolicy,omitempty" type:"Struct"`
	Cluster          *string                                 `json:"cluster,omitempty" xml:"cluster,omitempty"`
	CreateTime       *int64                                  `json:"createTime,omitempty" xml:"createTime,omitempty"`
	CreatorId        *string                                 `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	Data             *GetQuotaResponseBodyData               `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	Id               *string                                 `json:"id,omitempty" xml:"id,omitempty"`
	Name             *string                                 `json:"name,omitempty" xml:"name,omitempty"`
	NickName         *string                                 `json:"nickName,omitempty" xml:"nickName,omitempty"`
	Parameter        map[string]interface{}                  `json:"parameter,omitempty" xml:"parameter,omitempty"`
	ParentId         *string                                 `json:"parentId,omitempty" xml:"parentId,omitempty"`
	RegionId         *string                                 `json:"regionId,omitempty" xml:"regionId,omitempty"`
	RequestId        *string                                 `json:"requestId,omitempty" xml:"requestId,omitempty"`
	SaleTag          *GetQuotaResponseBodySaleTag            `json:"saleTag,omitempty" xml:"saleTag,omitempty" type:"Struct"`
	ScheduleInfo     *GetQuotaResponseBodyScheduleInfo       `json:"scheduleInfo,omitempty" xml:"scheduleInfo,omitempty" type:"Struct"`
	Status           *string                                 `json:"status,omitempty" xml:"status,omitempty"`
	SubQuotaInfoList []*GetQuotaResponseBodySubQuotaInfoList `json:"subQuotaInfoList,omitempty" xml:"subQuotaInfoList,omitempty" type:"Repeated"`
	Tag              *string                                 `json:"tag,omitempty" xml:"tag,omitempty"`
	TenantId         *string                                 `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	Type             *string                                 `json:"type,omitempty" xml:"type,omitempty"`
	Version          *string                                 `json:"version,omitempty" xml:"version,omitempty"`
}

func (s GetQuotaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *GetQuotaResponseBody) SetBillingPolicy(v *GetQuotaResponseBodyBillingPolicy) *GetQuotaResponseBody {
	s.BillingPolicy = v
	return s
}

func (s *GetQuotaResponseBody) SetCluster(v string) *GetQuotaResponseBody {
	s.Cluster = &v
	return s
}

func (s *GetQuotaResponseBody) SetCreateTime(v int64) *GetQuotaResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetQuotaResponseBody) SetCreatorId(v string) *GetQuotaResponseBody {
	s.CreatorId = &v
	return s
}

func (s *GetQuotaResponseBody) SetData(v *GetQuotaResponseBodyData) *GetQuotaResponseBody {
	s.Data = v
	return s
}

func (s *GetQuotaResponseBody) SetId(v string) *GetQuotaResponseBody {
	s.Id = &v
	return s
}

func (s *GetQuotaResponseBody) SetName(v string) *GetQuotaResponseBody {
	s.Name = &v
	return s
}

func (s *GetQuotaResponseBody) SetNickName(v string) *GetQuotaResponseBody {
	s.NickName = &v
	return s
}

func (s *GetQuotaResponseBody) SetParameter(v map[string]interface{}) *GetQuotaResponseBody {
	s.Parameter = v
	return s
}

func (s *GetQuotaResponseBody) SetParentId(v string) *GetQuotaResponseBody {
	s.ParentId = &v
	return s
}

func (s *GetQuotaResponseBody) SetRegionId(v string) *GetQuotaResponseBody {
	s.RegionId = &v
	return s
}

func (s *GetQuotaResponseBody) SetRequestId(v string) *GetQuotaResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetQuotaResponseBody) SetSaleTag(v *GetQuotaResponseBodySaleTag) *GetQuotaResponseBody {
	s.SaleTag = v
	return s
}

func (s *GetQuotaResponseBody) SetScheduleInfo(v *GetQuotaResponseBodyScheduleInfo) *GetQuotaResponseBody {
	s.ScheduleInfo = v
	return s
}

func (s *GetQuotaResponseBody) SetStatus(v string) *GetQuotaResponseBody {
	s.Status = &v
	return s
}

func (s *GetQuotaResponseBody) SetSubQuotaInfoList(v []*GetQuotaResponseBodySubQuotaInfoList) *GetQuotaResponseBody {
	s.SubQuotaInfoList = v
	return s
}

func (s *GetQuotaResponseBody) SetTag(v string) *GetQuotaResponseBody {
	s.Tag = &v
	return s
}

func (s *GetQuotaResponseBody) SetTenantId(v string) *GetQuotaResponseBody {
	s.TenantId = &v
	return s
}

func (s *GetQuotaResponseBody) SetType(v string) *GetQuotaResponseBody {
	s.Type = &v
	return s
}

func (s *GetQuotaResponseBody) SetVersion(v string) *GetQuotaResponseBody {
	s.Version = &v
	return s
}

type GetQuotaResponseBodyBillingPolicy struct {
	BillingMethod *string `json:"billingMethod,omitempty" xml:"billingMethod,omitempty"`
	OdpsSpecCode  *string `json:"odpsSpecCode,omitempty" xml:"odpsSpecCode,omitempty"`
	OrderId       *string `json:"orderId,omitempty" xml:"orderId,omitempty"`
}

func (s GetQuotaResponseBodyBillingPolicy) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaResponseBodyBillingPolicy) GoString() string {
	return s.String()
}

func (s *GetQuotaResponseBodyBillingPolicy) SetBillingMethod(v string) *GetQuotaResponseBodyBillingPolicy {
	s.BillingMethod = &v
	return s
}

func (s *GetQuotaResponseBodyBillingPolicy) SetOdpsSpecCode(v string) *GetQuotaResponseBodyBillingPolicy {
	s.OdpsSpecCode = &v
	return s
}

func (s *GetQuotaResponseBodyBillingPolicy) SetOrderId(v string) *GetQuotaResponseBodyBillingPolicy {
	s.OrderId = &v
	return s
}

type GetQuotaResponseBodyData struct {
	BillingPolicy    *GetQuotaResponseBodyDataBillingPolicy      `json:"billingPolicy,omitempty" xml:"billingPolicy,omitempty" type:"Struct"`
	Cluster          *string                                     `json:"cluster,omitempty" xml:"cluster,omitempty"`
	CreateTime       *int64                                      `json:"createTime,omitempty" xml:"createTime,omitempty"`
	CreatorId        *string                                     `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	Id               *string                                     `json:"id,omitempty" xml:"id,omitempty"`
	Name             *string                                     `json:"name,omitempty" xml:"name,omitempty"`
	NickName         *string                                     `json:"nickName,omitempty" xml:"nickName,omitempty"`
	Parameter        map[string]interface{}                      `json:"parameter,omitempty" xml:"parameter,omitempty"`
	ParentId         *string                                     `json:"parentId,omitempty" xml:"parentId,omitempty"`
	RegionId         *string                                     `json:"regionId,omitempty" xml:"regionId,omitempty"`
	SaleTag          *GetQuotaResponseBodyDataSaleTag            `json:"saleTag,omitempty" xml:"saleTag,omitempty" type:"Struct"`
	ScheduleInfo     *GetQuotaResponseBodyDataScheduleInfo       `json:"scheduleInfo,omitempty" xml:"scheduleInfo,omitempty" type:"Struct"`
	Status           *string                                     `json:"status,omitempty" xml:"status,omitempty"`
	SubQuotaInfoList []*GetQuotaResponseBodyDataSubQuotaInfoList `json:"subQuotaInfoList,omitempty" xml:"subQuotaInfoList,omitempty" type:"Repeated"`
	Tag              *string                                     `json:"tag,omitempty" xml:"tag,omitempty"`
	TenantId         *string                                     `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	Type             *string                                     `json:"type,omitempty" xml:"type,omitempty"`
	Version          *string                                     `json:"version,omitempty" xml:"version,omitempty"`
}

func (s GetQuotaResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetQuotaResponseBodyData) SetBillingPolicy(v *GetQuotaResponseBodyDataBillingPolicy) *GetQuotaResponseBodyData {
	s.BillingPolicy = v
	return s
}

func (s *GetQuotaResponseBodyData) SetCluster(v string) *GetQuotaResponseBodyData {
	s.Cluster = &v
	return s
}

func (s *GetQuotaResponseBodyData) SetCreateTime(v int64) *GetQuotaResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetQuotaResponseBodyData) SetCreatorId(v string) *GetQuotaResponseBodyData {
	s.CreatorId = &v
	return s
}

func (s *GetQuotaResponseBodyData) SetId(v string) *GetQuotaResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetQuotaResponseBodyData) SetName(v string) *GetQuotaResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetQuotaResponseBodyData) SetNickName(v string) *GetQuotaResponseBodyData {
	s.NickName = &v
	return s
}

func (s *GetQuotaResponseBodyData) SetParameter(v map[string]interface{}) *GetQuotaResponseBodyData {
	s.Parameter = v
	return s
}

func (s *GetQuotaResponseBodyData) SetParentId(v string) *GetQuotaResponseBodyData {
	s.ParentId = &v
	return s
}

func (s *GetQuotaResponseBodyData) SetRegionId(v string) *GetQuotaResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *GetQuotaResponseBodyData) SetSaleTag(v *GetQuotaResponseBodyDataSaleTag) *GetQuotaResponseBodyData {
	s.SaleTag = v
	return s
}

func (s *GetQuotaResponseBodyData) SetScheduleInfo(v *GetQuotaResponseBodyDataScheduleInfo) *GetQuotaResponseBodyData {
	s.ScheduleInfo = v
	return s
}

func (s *GetQuotaResponseBodyData) SetStatus(v string) *GetQuotaResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetQuotaResponseBodyData) SetSubQuotaInfoList(v []*GetQuotaResponseBodyDataSubQuotaInfoList) *GetQuotaResponseBodyData {
	s.SubQuotaInfoList = v
	return s
}

func (s *GetQuotaResponseBodyData) SetTag(v string) *GetQuotaResponseBodyData {
	s.Tag = &v
	return s
}

func (s *GetQuotaResponseBodyData) SetTenantId(v string) *GetQuotaResponseBodyData {
	s.TenantId = &v
	return s
}

func (s *GetQuotaResponseBodyData) SetType(v string) *GetQuotaResponseBodyData {
	s.Type = &v
	return s
}

func (s *GetQuotaResponseBodyData) SetVersion(v string) *GetQuotaResponseBodyData {
	s.Version = &v
	return s
}

type GetQuotaResponseBodyDataBillingPolicy struct {
	BillingMethod *string `json:"billingMethod,omitempty" xml:"billingMethod,omitempty"`
	OdpsSpecCode  *string `json:"odpsSpecCode,omitempty" xml:"odpsSpecCode,omitempty"`
	OrderId       *string `json:"orderId,omitempty" xml:"orderId,omitempty"`
}

func (s GetQuotaResponseBodyDataBillingPolicy) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaResponseBodyDataBillingPolicy) GoString() string {
	return s.String()
}

func (s *GetQuotaResponseBodyDataBillingPolicy) SetBillingMethod(v string) *GetQuotaResponseBodyDataBillingPolicy {
	s.BillingMethod = &v
	return s
}

func (s *GetQuotaResponseBodyDataBillingPolicy) SetOdpsSpecCode(v string) *GetQuotaResponseBodyDataBillingPolicy {
	s.OdpsSpecCode = &v
	return s
}

func (s *GetQuotaResponseBodyDataBillingPolicy) SetOrderId(v string) *GetQuotaResponseBodyDataBillingPolicy {
	s.OrderId = &v
	return s
}

type GetQuotaResponseBodyDataSaleTag struct {
	ResourceIds  []*string `json:"resourceIds,omitempty" xml:"resourceIds,omitempty" type:"Repeated"`
	ResourceType *string   `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
}

func (s GetQuotaResponseBodyDataSaleTag) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaResponseBodyDataSaleTag) GoString() string {
	return s.String()
}

func (s *GetQuotaResponseBodyDataSaleTag) SetResourceIds(v []*string) *GetQuotaResponseBodyDataSaleTag {
	s.ResourceIds = v
	return s
}

func (s *GetQuotaResponseBodyDataSaleTag) SetResourceType(v string) *GetQuotaResponseBodyDataSaleTag {
	s.ResourceType = &v
	return s
}

type GetQuotaResponseBodyDataScheduleInfo struct {
	CurrPlan     *string `json:"currPlan,omitempty" xml:"currPlan,omitempty"`
	CurrTime     *string `json:"currTime,omitempty" xml:"currTime,omitempty"`
	NextPlan     *string `json:"nextPlan,omitempty" xml:"nextPlan,omitempty"`
	NextTime     *string `json:"nextTime,omitempty" xml:"nextTime,omitempty"`
	OncePlan     *string `json:"oncePlan,omitempty" xml:"oncePlan,omitempty"`
	OnceTime     *string `json:"onceTime,omitempty" xml:"onceTime,omitempty"`
	OperatorName *string `json:"operatorName,omitempty" xml:"operatorName,omitempty"`
	Timezone     *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
}

func (s GetQuotaResponseBodyDataScheduleInfo) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaResponseBodyDataScheduleInfo) GoString() string {
	return s.String()
}

func (s *GetQuotaResponseBodyDataScheduleInfo) SetCurrPlan(v string) *GetQuotaResponseBodyDataScheduleInfo {
	s.CurrPlan = &v
	return s
}

func (s *GetQuotaResponseBodyDataScheduleInfo) SetCurrTime(v string) *GetQuotaResponseBodyDataScheduleInfo {
	s.CurrTime = &v
	return s
}

func (s *GetQuotaResponseBodyDataScheduleInfo) SetNextPlan(v string) *GetQuotaResponseBodyDataScheduleInfo {
	s.NextPlan = &v
	return s
}

func (s *GetQuotaResponseBodyDataScheduleInfo) SetNextTime(v string) *GetQuotaResponseBodyDataScheduleInfo {
	s.NextTime = &v
	return s
}

func (s *GetQuotaResponseBodyDataScheduleInfo) SetOncePlan(v string) *GetQuotaResponseBodyDataScheduleInfo {
	s.OncePlan = &v
	return s
}

func (s *GetQuotaResponseBodyDataScheduleInfo) SetOnceTime(v string) *GetQuotaResponseBodyDataScheduleInfo {
	s.OnceTime = &v
	return s
}

func (s *GetQuotaResponseBodyDataScheduleInfo) SetOperatorName(v string) *GetQuotaResponseBodyDataScheduleInfo {
	s.OperatorName = &v
	return s
}

func (s *GetQuotaResponseBodyDataScheduleInfo) SetTimezone(v string) *GetQuotaResponseBodyDataScheduleInfo {
	s.Timezone = &v
	return s
}

type GetQuotaResponseBodyDataSubQuotaInfoList struct {
	BillingPolicy *GetQuotaResponseBodyDataSubQuotaInfoListBillingPolicy `json:"billingPolicy,omitempty" xml:"billingPolicy,omitempty" type:"Struct"`
	Cluster       *string                                                `json:"cluster,omitempty" xml:"cluster,omitempty"`
	CreateTime    *int64                                                 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	CreatorId     *string                                                `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	Id            *string                                                `json:"id,omitempty" xml:"id,omitempty"`
	Name          *string                                                `json:"name,omitempty" xml:"name,omitempty"`
	NickName      *string                                                `json:"nickName,omitempty" xml:"nickName,omitempty"`
	Parameter     map[string]interface{}                                 `json:"parameter,omitempty" xml:"parameter,omitempty"`
	ParentId      *string                                                `json:"parentId,omitempty" xml:"parentId,omitempty"`
	RegionId      *string                                                `json:"regionId,omitempty" xml:"regionId,omitempty"`
	SaleTag       *GetQuotaResponseBodyDataSubQuotaInfoListSaleTag       `json:"saleTag,omitempty" xml:"saleTag,omitempty" type:"Struct"`
	ScheduleInfo  *GetQuotaResponseBodyDataSubQuotaInfoListScheduleInfo  `json:"scheduleInfo,omitempty" xml:"scheduleInfo,omitempty" type:"Struct"`
	Status        *string                                                `json:"status,omitempty" xml:"status,omitempty"`
	Tag           *string                                                `json:"tag,omitempty" xml:"tag,omitempty"`
	TenantId      *string                                                `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	Type          *string                                                `json:"type,omitempty" xml:"type,omitempty"`
	Version       *string                                                `json:"version,omitempty" xml:"version,omitempty"`
}

func (s GetQuotaResponseBodyDataSubQuotaInfoList) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaResponseBodyDataSubQuotaInfoList) GoString() string {
	return s.String()
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoList) SetBillingPolicy(v *GetQuotaResponseBodyDataSubQuotaInfoListBillingPolicy) *GetQuotaResponseBodyDataSubQuotaInfoList {
	s.BillingPolicy = v
	return s
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoList) SetCluster(v string) *GetQuotaResponseBodyDataSubQuotaInfoList {
	s.Cluster = &v
	return s
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoList) SetCreateTime(v int64) *GetQuotaResponseBodyDataSubQuotaInfoList {
	s.CreateTime = &v
	return s
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoList) SetCreatorId(v string) *GetQuotaResponseBodyDataSubQuotaInfoList {
	s.CreatorId = &v
	return s
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoList) SetId(v string) *GetQuotaResponseBodyDataSubQuotaInfoList {
	s.Id = &v
	return s
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoList) SetName(v string) *GetQuotaResponseBodyDataSubQuotaInfoList {
	s.Name = &v
	return s
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoList) SetNickName(v string) *GetQuotaResponseBodyDataSubQuotaInfoList {
	s.NickName = &v
	return s
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoList) SetParameter(v map[string]interface{}) *GetQuotaResponseBodyDataSubQuotaInfoList {
	s.Parameter = v
	return s
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoList) SetParentId(v string) *GetQuotaResponseBodyDataSubQuotaInfoList {
	s.ParentId = &v
	return s
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoList) SetRegionId(v string) *GetQuotaResponseBodyDataSubQuotaInfoList {
	s.RegionId = &v
	return s
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoList) SetSaleTag(v *GetQuotaResponseBodyDataSubQuotaInfoListSaleTag) *GetQuotaResponseBodyDataSubQuotaInfoList {
	s.SaleTag = v
	return s
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoList) SetScheduleInfo(v *GetQuotaResponseBodyDataSubQuotaInfoListScheduleInfo) *GetQuotaResponseBodyDataSubQuotaInfoList {
	s.ScheduleInfo = v
	return s
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoList) SetStatus(v string) *GetQuotaResponseBodyDataSubQuotaInfoList {
	s.Status = &v
	return s
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoList) SetTag(v string) *GetQuotaResponseBodyDataSubQuotaInfoList {
	s.Tag = &v
	return s
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoList) SetTenantId(v string) *GetQuotaResponseBodyDataSubQuotaInfoList {
	s.TenantId = &v
	return s
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoList) SetType(v string) *GetQuotaResponseBodyDataSubQuotaInfoList {
	s.Type = &v
	return s
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoList) SetVersion(v string) *GetQuotaResponseBodyDataSubQuotaInfoList {
	s.Version = &v
	return s
}

type GetQuotaResponseBodyDataSubQuotaInfoListBillingPolicy struct {
	BillingMethod *string `json:"billingMethod,omitempty" xml:"billingMethod,omitempty"`
	OdpsSpecCode  *string `json:"odpsSpecCode,omitempty" xml:"odpsSpecCode,omitempty"`
	OrderId       *string `json:"orderId,omitempty" xml:"orderId,omitempty"`
}

func (s GetQuotaResponseBodyDataSubQuotaInfoListBillingPolicy) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaResponseBodyDataSubQuotaInfoListBillingPolicy) GoString() string {
	return s.String()
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoListBillingPolicy) SetBillingMethod(v string) *GetQuotaResponseBodyDataSubQuotaInfoListBillingPolicy {
	s.BillingMethod = &v
	return s
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoListBillingPolicy) SetOdpsSpecCode(v string) *GetQuotaResponseBodyDataSubQuotaInfoListBillingPolicy {
	s.OdpsSpecCode = &v
	return s
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoListBillingPolicy) SetOrderId(v string) *GetQuotaResponseBodyDataSubQuotaInfoListBillingPolicy {
	s.OrderId = &v
	return s
}

type GetQuotaResponseBodyDataSubQuotaInfoListSaleTag struct {
	ResourceIds  []*string `json:"resourceIds,omitempty" xml:"resourceIds,omitempty" type:"Repeated"`
	ResourceType *string   `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
}

func (s GetQuotaResponseBodyDataSubQuotaInfoListSaleTag) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaResponseBodyDataSubQuotaInfoListSaleTag) GoString() string {
	return s.String()
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoListSaleTag) SetResourceIds(v []*string) *GetQuotaResponseBodyDataSubQuotaInfoListSaleTag {
	s.ResourceIds = v
	return s
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoListSaleTag) SetResourceType(v string) *GetQuotaResponseBodyDataSubQuotaInfoListSaleTag {
	s.ResourceType = &v
	return s
}

type GetQuotaResponseBodyDataSubQuotaInfoListScheduleInfo struct {
	CurrPlan     *string `json:"currPlan,omitempty" xml:"currPlan,omitempty"`
	CurrTime     *string `json:"currTime,omitempty" xml:"currTime,omitempty"`
	NextPlan     *string `json:"nextPlan,omitempty" xml:"nextPlan,omitempty"`
	NextTime     *string `json:"nextTime,omitempty" xml:"nextTime,omitempty"`
	OncePlan     *string `json:"oncePlan,omitempty" xml:"oncePlan,omitempty"`
	OnceTime     *string `json:"onceTime,omitempty" xml:"onceTime,omitempty"`
	OperatorName *string `json:"operatorName,omitempty" xml:"operatorName,omitempty"`
	Timezone     *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
}

func (s GetQuotaResponseBodyDataSubQuotaInfoListScheduleInfo) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaResponseBodyDataSubQuotaInfoListScheduleInfo) GoString() string {
	return s.String()
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoListScheduleInfo) SetCurrPlan(v string) *GetQuotaResponseBodyDataSubQuotaInfoListScheduleInfo {
	s.CurrPlan = &v
	return s
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoListScheduleInfo) SetCurrTime(v string) *GetQuotaResponseBodyDataSubQuotaInfoListScheduleInfo {
	s.CurrTime = &v
	return s
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoListScheduleInfo) SetNextPlan(v string) *GetQuotaResponseBodyDataSubQuotaInfoListScheduleInfo {
	s.NextPlan = &v
	return s
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoListScheduleInfo) SetNextTime(v string) *GetQuotaResponseBodyDataSubQuotaInfoListScheduleInfo {
	s.NextTime = &v
	return s
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoListScheduleInfo) SetOncePlan(v string) *GetQuotaResponseBodyDataSubQuotaInfoListScheduleInfo {
	s.OncePlan = &v
	return s
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoListScheduleInfo) SetOnceTime(v string) *GetQuotaResponseBodyDataSubQuotaInfoListScheduleInfo {
	s.OnceTime = &v
	return s
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoListScheduleInfo) SetOperatorName(v string) *GetQuotaResponseBodyDataSubQuotaInfoListScheduleInfo {
	s.OperatorName = &v
	return s
}

func (s *GetQuotaResponseBodyDataSubQuotaInfoListScheduleInfo) SetTimezone(v string) *GetQuotaResponseBodyDataSubQuotaInfoListScheduleInfo {
	s.Timezone = &v
	return s
}

type GetQuotaResponseBodySaleTag struct {
	ResourceIds  []*string `json:"resourceIds,omitempty" xml:"resourceIds,omitempty" type:"Repeated"`
	ResourceType *string   `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
}

func (s GetQuotaResponseBodySaleTag) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaResponseBodySaleTag) GoString() string {
	return s.String()
}

func (s *GetQuotaResponseBodySaleTag) SetResourceIds(v []*string) *GetQuotaResponseBodySaleTag {
	s.ResourceIds = v
	return s
}

func (s *GetQuotaResponseBodySaleTag) SetResourceType(v string) *GetQuotaResponseBodySaleTag {
	s.ResourceType = &v
	return s
}

type GetQuotaResponseBodyScheduleInfo struct {
	CurrPlan     *string `json:"currPlan,omitempty" xml:"currPlan,omitempty"`
	CurrTime     *string `json:"currTime,omitempty" xml:"currTime,omitempty"`
	NextPlan     *string `json:"nextPlan,omitempty" xml:"nextPlan,omitempty"`
	NextTime     *string `json:"nextTime,omitempty" xml:"nextTime,omitempty"`
	OncePlan     *string `json:"oncePlan,omitempty" xml:"oncePlan,omitempty"`
	OnceTime     *string `json:"onceTime,omitempty" xml:"onceTime,omitempty"`
	OperatorName *string `json:"operatorName,omitempty" xml:"operatorName,omitempty"`
	Timezone     *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
}

func (s GetQuotaResponseBodyScheduleInfo) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaResponseBodyScheduleInfo) GoString() string {
	return s.String()
}

func (s *GetQuotaResponseBodyScheduleInfo) SetCurrPlan(v string) *GetQuotaResponseBodyScheduleInfo {
	s.CurrPlan = &v
	return s
}

func (s *GetQuotaResponseBodyScheduleInfo) SetCurrTime(v string) *GetQuotaResponseBodyScheduleInfo {
	s.CurrTime = &v
	return s
}

func (s *GetQuotaResponseBodyScheduleInfo) SetNextPlan(v string) *GetQuotaResponseBodyScheduleInfo {
	s.NextPlan = &v
	return s
}

func (s *GetQuotaResponseBodyScheduleInfo) SetNextTime(v string) *GetQuotaResponseBodyScheduleInfo {
	s.NextTime = &v
	return s
}

func (s *GetQuotaResponseBodyScheduleInfo) SetOncePlan(v string) *GetQuotaResponseBodyScheduleInfo {
	s.OncePlan = &v
	return s
}

func (s *GetQuotaResponseBodyScheduleInfo) SetOnceTime(v string) *GetQuotaResponseBodyScheduleInfo {
	s.OnceTime = &v
	return s
}

func (s *GetQuotaResponseBodyScheduleInfo) SetOperatorName(v string) *GetQuotaResponseBodyScheduleInfo {
	s.OperatorName = &v
	return s
}

func (s *GetQuotaResponseBodyScheduleInfo) SetTimezone(v string) *GetQuotaResponseBodyScheduleInfo {
	s.Timezone = &v
	return s
}

type GetQuotaResponseBodySubQuotaInfoList struct {
	BillingPolicy *GetQuotaResponseBodySubQuotaInfoListBillingPolicy `json:"billingPolicy,omitempty" xml:"billingPolicy,omitempty" type:"Struct"`
	Cluster       *string                                            `json:"cluster,omitempty" xml:"cluster,omitempty"`
	CreateTime    *int64                                             `json:"createTime,omitempty" xml:"createTime,omitempty"`
	CreatorId     *string                                            `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	Id            *string                                            `json:"id,omitempty" xml:"id,omitempty"`
	Name          *string                                            `json:"name,omitempty" xml:"name,omitempty"`
	NickName      *string                                            `json:"nickName,omitempty" xml:"nickName,omitempty"`
	Parameter     map[string]interface{}                             `json:"parameter,omitempty" xml:"parameter,omitempty"`
	ParentId      *string                                            `json:"parentId,omitempty" xml:"parentId,omitempty"`
	RegionId      *string                                            `json:"regionId,omitempty" xml:"regionId,omitempty"`
	SaleTag       *GetQuotaResponseBodySubQuotaInfoListSaleTag       `json:"saleTag,omitempty" xml:"saleTag,omitempty" type:"Struct"`
	ScheduleInfo  *GetQuotaResponseBodySubQuotaInfoListScheduleInfo  `json:"scheduleInfo,omitempty" xml:"scheduleInfo,omitempty" type:"Struct"`
	Status        *string                                            `json:"status,omitempty" xml:"status,omitempty"`
	Tag           *string                                            `json:"tag,omitempty" xml:"tag,omitempty"`
	TenantId      *string                                            `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	Type          *string                                            `json:"type,omitempty" xml:"type,omitempty"`
	Version       *string                                            `json:"version,omitempty" xml:"version,omitempty"`
}

func (s GetQuotaResponseBodySubQuotaInfoList) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaResponseBodySubQuotaInfoList) GoString() string {
	return s.String()
}

func (s *GetQuotaResponseBodySubQuotaInfoList) SetBillingPolicy(v *GetQuotaResponseBodySubQuotaInfoListBillingPolicy) *GetQuotaResponseBodySubQuotaInfoList {
	s.BillingPolicy = v
	return s
}

func (s *GetQuotaResponseBodySubQuotaInfoList) SetCluster(v string) *GetQuotaResponseBodySubQuotaInfoList {
	s.Cluster = &v
	return s
}

func (s *GetQuotaResponseBodySubQuotaInfoList) SetCreateTime(v int64) *GetQuotaResponseBodySubQuotaInfoList {
	s.CreateTime = &v
	return s
}

func (s *GetQuotaResponseBodySubQuotaInfoList) SetCreatorId(v string) *GetQuotaResponseBodySubQuotaInfoList {
	s.CreatorId = &v
	return s
}

func (s *GetQuotaResponseBodySubQuotaInfoList) SetId(v string) *GetQuotaResponseBodySubQuotaInfoList {
	s.Id = &v
	return s
}

func (s *GetQuotaResponseBodySubQuotaInfoList) SetName(v string) *GetQuotaResponseBodySubQuotaInfoList {
	s.Name = &v
	return s
}

func (s *GetQuotaResponseBodySubQuotaInfoList) SetNickName(v string) *GetQuotaResponseBodySubQuotaInfoList {
	s.NickName = &v
	return s
}

func (s *GetQuotaResponseBodySubQuotaInfoList) SetParameter(v map[string]interface{}) *GetQuotaResponseBodySubQuotaInfoList {
	s.Parameter = v
	return s
}

func (s *GetQuotaResponseBodySubQuotaInfoList) SetParentId(v string) *GetQuotaResponseBodySubQuotaInfoList {
	s.ParentId = &v
	return s
}

func (s *GetQuotaResponseBodySubQuotaInfoList) SetRegionId(v string) *GetQuotaResponseBodySubQuotaInfoList {
	s.RegionId = &v
	return s
}

func (s *GetQuotaResponseBodySubQuotaInfoList) SetSaleTag(v *GetQuotaResponseBodySubQuotaInfoListSaleTag) *GetQuotaResponseBodySubQuotaInfoList {
	s.SaleTag = v
	return s
}

func (s *GetQuotaResponseBodySubQuotaInfoList) SetScheduleInfo(v *GetQuotaResponseBodySubQuotaInfoListScheduleInfo) *GetQuotaResponseBodySubQuotaInfoList {
	s.ScheduleInfo = v
	return s
}

func (s *GetQuotaResponseBodySubQuotaInfoList) SetStatus(v string) *GetQuotaResponseBodySubQuotaInfoList {
	s.Status = &v
	return s
}

func (s *GetQuotaResponseBodySubQuotaInfoList) SetTag(v string) *GetQuotaResponseBodySubQuotaInfoList {
	s.Tag = &v
	return s
}

func (s *GetQuotaResponseBodySubQuotaInfoList) SetTenantId(v string) *GetQuotaResponseBodySubQuotaInfoList {
	s.TenantId = &v
	return s
}

func (s *GetQuotaResponseBodySubQuotaInfoList) SetType(v string) *GetQuotaResponseBodySubQuotaInfoList {
	s.Type = &v
	return s
}

func (s *GetQuotaResponseBodySubQuotaInfoList) SetVersion(v string) *GetQuotaResponseBodySubQuotaInfoList {
	s.Version = &v
	return s
}

type GetQuotaResponseBodySubQuotaInfoListBillingPolicy struct {
	BillingMethod *string `json:"billingMethod,omitempty" xml:"billingMethod,omitempty"`
	OdpsSpecCode  *string `json:"odpsSpecCode,omitempty" xml:"odpsSpecCode,omitempty"`
	OrderId       *string `json:"orderId,omitempty" xml:"orderId,omitempty"`
}

func (s GetQuotaResponseBodySubQuotaInfoListBillingPolicy) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaResponseBodySubQuotaInfoListBillingPolicy) GoString() string {
	return s.String()
}

func (s *GetQuotaResponseBodySubQuotaInfoListBillingPolicy) SetBillingMethod(v string) *GetQuotaResponseBodySubQuotaInfoListBillingPolicy {
	s.BillingMethod = &v
	return s
}

func (s *GetQuotaResponseBodySubQuotaInfoListBillingPolicy) SetOdpsSpecCode(v string) *GetQuotaResponseBodySubQuotaInfoListBillingPolicy {
	s.OdpsSpecCode = &v
	return s
}

func (s *GetQuotaResponseBodySubQuotaInfoListBillingPolicy) SetOrderId(v string) *GetQuotaResponseBodySubQuotaInfoListBillingPolicy {
	s.OrderId = &v
	return s
}

type GetQuotaResponseBodySubQuotaInfoListSaleTag struct {
	ResourceIds  []*string `json:"resourceIds,omitempty" xml:"resourceIds,omitempty" type:"Repeated"`
	ResourceType *string   `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
}

func (s GetQuotaResponseBodySubQuotaInfoListSaleTag) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaResponseBodySubQuotaInfoListSaleTag) GoString() string {
	return s.String()
}

func (s *GetQuotaResponseBodySubQuotaInfoListSaleTag) SetResourceIds(v []*string) *GetQuotaResponseBodySubQuotaInfoListSaleTag {
	s.ResourceIds = v
	return s
}

func (s *GetQuotaResponseBodySubQuotaInfoListSaleTag) SetResourceType(v string) *GetQuotaResponseBodySubQuotaInfoListSaleTag {
	s.ResourceType = &v
	return s
}

type GetQuotaResponseBodySubQuotaInfoListScheduleInfo struct {
	CurrPlan     *string `json:"currPlan,omitempty" xml:"currPlan,omitempty"`
	CurrTime     *string `json:"currTime,omitempty" xml:"currTime,omitempty"`
	NextPlan     *string `json:"nextPlan,omitempty" xml:"nextPlan,omitempty"`
	NextTime     *string `json:"nextTime,omitempty" xml:"nextTime,omitempty"`
	OncePlan     *string `json:"oncePlan,omitempty" xml:"oncePlan,omitempty"`
	OnceTime     *string `json:"onceTime,omitempty" xml:"onceTime,omitempty"`
	OperatorName *string `json:"operatorName,omitempty" xml:"operatorName,omitempty"`
	Timezone     *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
}

func (s GetQuotaResponseBodySubQuotaInfoListScheduleInfo) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaResponseBodySubQuotaInfoListScheduleInfo) GoString() string {
	return s.String()
}

func (s *GetQuotaResponseBodySubQuotaInfoListScheduleInfo) SetCurrPlan(v string) *GetQuotaResponseBodySubQuotaInfoListScheduleInfo {
	s.CurrPlan = &v
	return s
}

func (s *GetQuotaResponseBodySubQuotaInfoListScheduleInfo) SetCurrTime(v string) *GetQuotaResponseBodySubQuotaInfoListScheduleInfo {
	s.CurrTime = &v
	return s
}

func (s *GetQuotaResponseBodySubQuotaInfoListScheduleInfo) SetNextPlan(v string) *GetQuotaResponseBodySubQuotaInfoListScheduleInfo {
	s.NextPlan = &v
	return s
}

func (s *GetQuotaResponseBodySubQuotaInfoListScheduleInfo) SetNextTime(v string) *GetQuotaResponseBodySubQuotaInfoListScheduleInfo {
	s.NextTime = &v
	return s
}

func (s *GetQuotaResponseBodySubQuotaInfoListScheduleInfo) SetOncePlan(v string) *GetQuotaResponseBodySubQuotaInfoListScheduleInfo {
	s.OncePlan = &v
	return s
}

func (s *GetQuotaResponseBodySubQuotaInfoListScheduleInfo) SetOnceTime(v string) *GetQuotaResponseBodySubQuotaInfoListScheduleInfo {
	s.OnceTime = &v
	return s
}

func (s *GetQuotaResponseBodySubQuotaInfoListScheduleInfo) SetOperatorName(v string) *GetQuotaResponseBodySubQuotaInfoListScheduleInfo {
	s.OperatorName = &v
	return s
}

func (s *GetQuotaResponseBodySubQuotaInfoListScheduleInfo) SetTimezone(v string) *GetQuotaResponseBodySubQuotaInfoListScheduleInfo {
	s.Timezone = &v
	return s
}

type GetQuotaResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetQuotaResponse) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaResponse) GoString() string {
	return s.String()
}

func (s *GetQuotaResponse) SetHeaders(v map[string]*string) *GetQuotaResponse {
	s.Headers = v
	return s
}

func (s *GetQuotaResponse) SetStatusCode(v int32) *GetQuotaResponse {
	s.StatusCode = &v
	return s
}

func (s *GetQuotaResponse) SetBody(v *GetQuotaResponseBody) *GetQuotaResponse {
	s.Body = v
	return s
}

type GetQuotaPlanRequest struct {
	// The ID of the region.
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// The ID of the tenant.
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s GetQuotaPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaPlanRequest) GoString() string {
	return s.String()
}

func (s *GetQuotaPlanRequest) SetRegion(v string) *GetQuotaPlanRequest {
	s.Region = &v
	return s
}

func (s *GetQuotaPlanRequest) SetTenantId(v string) *GetQuotaPlanRequest {
	s.TenantId = &v
	return s
}

type GetQuotaPlanResponseBody struct {
	// The returned data.
	Data *GetQuotaPlanResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetQuotaPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaPlanResponseBody) GoString() string {
	return s.String()
}

func (s *GetQuotaPlanResponseBody) SetData(v *GetQuotaPlanResponseBodyData) *GetQuotaPlanResponseBody {
	s.Data = v
	return s
}

func (s *GetQuotaPlanResponseBody) SetRequestId(v string) *GetQuotaPlanResponseBody {
	s.RequestId = &v
	return s
}

type GetQuotaPlanResponseBodyData struct {
	// The time when the quota plan was created.
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The name of the quota plan.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The details of the quota.
	Quota *GetQuotaPlanResponseBodyDataQuota `json:"quota,omitempty" xml:"quota,omitempty" type:"Struct"`
}

func (s GetQuotaPlanResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaPlanResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetQuotaPlanResponseBodyData) SetCreateTime(v string) *GetQuotaPlanResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetQuotaPlanResponseBodyData) SetName(v string) *GetQuotaPlanResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetQuotaPlanResponseBodyData) SetQuota(v *GetQuotaPlanResponseBodyDataQuota) *GetQuotaPlanResponseBodyData {
	s.Quota = v
	return s
}

type GetQuotaPlanResponseBodyDataQuota struct {
	// The information of the order.
	BillingPolicy *GetQuotaPlanResponseBodyDataQuotaBillingPolicy `json:"billingPolicy,omitempty" xml:"billingPolicy,omitempty" type:"Struct"`
	// The ID of the cluster.
	Cluster *string `json:"cluster,omitempty" xml:"cluster,omitempty"`
	// The time when the quota plan was created.
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The ID of the Alibaba Cloud account that is used to create the resource.
	CreatorId *string `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	// The ID of the quota.
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The name of the quota.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The alias of the quota.
	NickName *string `json:"nickName,omitempty" xml:"nickName,omitempty"`
	// The description of the quota.
	Parameter map[string]interface{} `json:"parameter,omitempty" xml:"parameter,omitempty"`
	// The ID of the parent resource.
	ParentId *string `json:"parentId,omitempty" xml:"parentId,omitempty"`
	// The ID of the region.
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The information of the scheduling plan.
	ScheduleInfo *GetQuotaPlanResponseBodyDataQuotaScheduleInfo `json:"scheduleInfo,omitempty" xml:"scheduleInfo,omitempty" type:"Struct"`
	// The status of the resource.
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The information of the level-2 quota.
	SubQuotaInfoList []*GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList `json:"subQuotaInfoList,omitempty" xml:"subQuotaInfoList,omitempty" type:"Repeated"`
	// The tag of the resource for the quota.
	Tag *string `json:"tag,omitempty" xml:"tag,omitempty"`
	// The ID of the tenant.
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The type of the resource system. This parameter corresponds to the resourceSystemType parameter of the cluster.
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The version number.
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s GetQuotaPlanResponseBodyDataQuota) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaPlanResponseBodyDataQuota) GoString() string {
	return s.String()
}

func (s *GetQuotaPlanResponseBodyDataQuota) SetBillingPolicy(v *GetQuotaPlanResponseBodyDataQuotaBillingPolicy) *GetQuotaPlanResponseBodyDataQuota {
	s.BillingPolicy = v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuota) SetCluster(v string) *GetQuotaPlanResponseBodyDataQuota {
	s.Cluster = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuota) SetCreateTime(v int64) *GetQuotaPlanResponseBodyDataQuota {
	s.CreateTime = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuota) SetCreatorId(v string) *GetQuotaPlanResponseBodyDataQuota {
	s.CreatorId = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuota) SetId(v string) *GetQuotaPlanResponseBodyDataQuota {
	s.Id = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuota) SetName(v string) *GetQuotaPlanResponseBodyDataQuota {
	s.Name = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuota) SetNickName(v string) *GetQuotaPlanResponseBodyDataQuota {
	s.NickName = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuota) SetParameter(v map[string]interface{}) *GetQuotaPlanResponseBodyDataQuota {
	s.Parameter = v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuota) SetParentId(v string) *GetQuotaPlanResponseBodyDataQuota {
	s.ParentId = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuota) SetRegionId(v string) *GetQuotaPlanResponseBodyDataQuota {
	s.RegionId = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuota) SetScheduleInfo(v *GetQuotaPlanResponseBodyDataQuotaScheduleInfo) *GetQuotaPlanResponseBodyDataQuota {
	s.ScheduleInfo = v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuota) SetStatus(v string) *GetQuotaPlanResponseBodyDataQuota {
	s.Status = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuota) SetSubQuotaInfoList(v []*GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList) *GetQuotaPlanResponseBodyDataQuota {
	s.SubQuotaInfoList = v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuota) SetTag(v string) *GetQuotaPlanResponseBodyDataQuota {
	s.Tag = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuota) SetTenantId(v string) *GetQuotaPlanResponseBodyDataQuota {
	s.TenantId = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuota) SetType(v string) *GetQuotaPlanResponseBodyDataQuota {
	s.Type = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuota) SetVersion(v string) *GetQuotaPlanResponseBodyDataQuota {
	s.Version = &v
	return s
}

type GetQuotaPlanResponseBodyDataQuotaBillingPolicy struct {
	// The billing method of the quota. Valid values:
	//
	// *   subscription: a subscription quota.
	// *   payasyougo: a pay-as-you-go quota.
	BillingMethod *string `json:"billingMethod,omitempty" xml:"billingMethod,omitempty"`
	// The specifications of the order.
	OdpsSpecCode *string `json:"odpsSpecCode,omitempty" xml:"odpsSpecCode,omitempty"`
	// The ID of the order.
	OrderId *string `json:"orderId,omitempty" xml:"orderId,omitempty"`
}

func (s GetQuotaPlanResponseBodyDataQuotaBillingPolicy) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaPlanResponseBodyDataQuotaBillingPolicy) GoString() string {
	return s.String()
}

func (s *GetQuotaPlanResponseBodyDataQuotaBillingPolicy) SetBillingMethod(v string) *GetQuotaPlanResponseBodyDataQuotaBillingPolicy {
	s.BillingMethod = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuotaBillingPolicy) SetOdpsSpecCode(v string) *GetQuotaPlanResponseBodyDataQuotaBillingPolicy {
	s.OdpsSpecCode = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuotaBillingPolicy) SetOrderId(v string) *GetQuotaPlanResponseBodyDataQuotaBillingPolicy {
	s.OrderId = &v
	return s
}

type GetQuotaPlanResponseBodyDataQuotaScheduleInfo struct {
	// The quota plan that takes effect based on the scheduling plan.
	CurrPlan *string `json:"currPlan,omitempty" xml:"currPlan,omitempty"`
	// The time when the current quota plan is scheduled.
	CurrTime *string `json:"currTime,omitempty" xml:"currTime,omitempty"`
	// The next quota plan that will take effect based on the scheduling plan.
	NextPlan *string `json:"nextPlan,omitempty" xml:"nextPlan,omitempty"`
	// The time when the next quota plan is scheduled.
	NextTime *string `json:"nextTime,omitempty" xml:"nextTime,omitempty"`
	// If the quota plan that immediately takes effect is different from the current quota plan, this parameter is not empty.
	OncePlan *string `json:"oncePlan,omitempty" xml:"oncePlan,omitempty"`
	// The time when the quota plan immediately takes effect.
	OnceTime *string `json:"onceTime,omitempty" xml:"onceTime,omitempty"`
	// The name of the operator.
	OperatorName *string `json:"operatorName,omitempty" xml:"operatorName,omitempty"`
}

func (s GetQuotaPlanResponseBodyDataQuotaScheduleInfo) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaPlanResponseBodyDataQuotaScheduleInfo) GoString() string {
	return s.String()
}

func (s *GetQuotaPlanResponseBodyDataQuotaScheduleInfo) SetCurrPlan(v string) *GetQuotaPlanResponseBodyDataQuotaScheduleInfo {
	s.CurrPlan = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuotaScheduleInfo) SetCurrTime(v string) *GetQuotaPlanResponseBodyDataQuotaScheduleInfo {
	s.CurrTime = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuotaScheduleInfo) SetNextPlan(v string) *GetQuotaPlanResponseBodyDataQuotaScheduleInfo {
	s.NextPlan = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuotaScheduleInfo) SetNextTime(v string) *GetQuotaPlanResponseBodyDataQuotaScheduleInfo {
	s.NextTime = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuotaScheduleInfo) SetOncePlan(v string) *GetQuotaPlanResponseBodyDataQuotaScheduleInfo {
	s.OncePlan = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuotaScheduleInfo) SetOnceTime(v string) *GetQuotaPlanResponseBodyDataQuotaScheduleInfo {
	s.OnceTime = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuotaScheduleInfo) SetOperatorName(v string) *GetQuotaPlanResponseBodyDataQuotaScheduleInfo {
	s.OperatorName = &v
	return s
}

type GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList struct {
	// The information of the order.
	BillingPolicy *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoListBillingPolicy `json:"billingPolicy,omitempty" xml:"billingPolicy,omitempty" type:"Struct"`
	// The ID of the cluster.
	Cluster *string `json:"cluster,omitempty" xml:"cluster,omitempty"`
	// The time when the resource was created.
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The ID of the user who created the quota plan.
	CreatorId *string `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	// The ID of the level-2 quota.
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The name of the level-2 quota.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The alias of the level-2 quota.
	NickName *string `json:"nickName,omitempty" xml:"nickName,omitempty"`
	// The description of the quota.
	Parameter map[string]interface{} `json:"parameter,omitempty" xml:"parameter,omitempty"`
	// The ID of the parent resource.
	ParentId *string `json:"parentId,omitempty" xml:"parentId,omitempty"`
	// The ID of the region.
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The information of the scheduling plan.
	ScheduleInfo *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoListScheduleInfo `json:"scheduleInfo,omitempty" xml:"scheduleInfo,omitempty" type:"Struct"`
	// The status of the resource.
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The tag of the resource for the quota.
	Tag *string `json:"tag,omitempty" xml:"tag,omitempty"`
	// The ID of the tenant.
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The type of the resource system. This parameter corresponds to the resourceSystemType parameter of the cluster.
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The version number.
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList) GoString() string {
	return s.String()
}

func (s *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList) SetBillingPolicy(v *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoListBillingPolicy) *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList {
	s.BillingPolicy = v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList) SetCluster(v string) *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList {
	s.Cluster = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList) SetCreateTime(v int64) *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList {
	s.CreateTime = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList) SetCreatorId(v string) *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList {
	s.CreatorId = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList) SetId(v string) *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList {
	s.Id = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList) SetName(v string) *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList {
	s.Name = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList) SetNickName(v string) *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList {
	s.NickName = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList) SetParameter(v map[string]interface{}) *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList {
	s.Parameter = v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList) SetParentId(v string) *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList {
	s.ParentId = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList) SetRegionId(v string) *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList {
	s.RegionId = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList) SetScheduleInfo(v *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoListScheduleInfo) *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList {
	s.ScheduleInfo = v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList) SetStatus(v string) *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList {
	s.Status = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList) SetTag(v string) *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList {
	s.Tag = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList) SetTenantId(v string) *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList {
	s.TenantId = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList) SetType(v string) *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList {
	s.Type = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList) SetVersion(v string) *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoList {
	s.Version = &v
	return s
}

type GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoListBillingPolicy struct {
	// The billing method of the quota. Valid values:
	//
	// *   subscription: a subscription quota.
	// *   payasyougo: a pay-as-you-go quota.
	BillingMethod *string `json:"billingMethod,omitempty" xml:"billingMethod,omitempty"`
	// The specifications of the order.
	OdpsSpecCode *string `json:"odpsSpecCode,omitempty" xml:"odpsSpecCode,omitempty"`
	// The ID of the order.
	OrderId *string `json:"orderId,omitempty" xml:"orderId,omitempty"`
}

func (s GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoListBillingPolicy) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoListBillingPolicy) GoString() string {
	return s.String()
}

func (s *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoListBillingPolicy) SetBillingMethod(v string) *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoListBillingPolicy {
	s.BillingMethod = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoListBillingPolicy) SetOdpsSpecCode(v string) *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoListBillingPolicy {
	s.OdpsSpecCode = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoListBillingPolicy) SetOrderId(v string) *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoListBillingPolicy {
	s.OrderId = &v
	return s
}

type GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoListScheduleInfo struct {
	// The quota plan that takes effect based on the scheduling plan.
	CurrPlan *string `json:"currPlan,omitempty" xml:"currPlan,omitempty"`
	// The time when the current quota plan is scheduled.
	CurrTime *string `json:"currTime,omitempty" xml:"currTime,omitempty"`
	// The next quota plan that will take effect based on the scheduling plan.
	NextPlan *string `json:"nextPlan,omitempty" xml:"nextPlan,omitempty"`
	// The time when the next quota plan is scheduled.
	NextTime *string `json:"nextTime,omitempty" xml:"nextTime,omitempty"`
	// If the quota plan that immediately takes effect is different from the current quota plan, this parameter is not empty.
	OncePlan *string `json:"oncePlan,omitempty" xml:"oncePlan,omitempty"`
	// The time when the quota plan immediately takes effect.
	OnceTime *string `json:"onceTime,omitempty" xml:"onceTime,omitempty"`
	// The name of the operator.
	OperatorName *string `json:"operatorName,omitempty" xml:"operatorName,omitempty"`
}

func (s GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoListScheduleInfo) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoListScheduleInfo) GoString() string {
	return s.String()
}

func (s *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoListScheduleInfo) SetCurrPlan(v string) *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoListScheduleInfo {
	s.CurrPlan = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoListScheduleInfo) SetCurrTime(v string) *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoListScheduleInfo {
	s.CurrTime = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoListScheduleInfo) SetNextPlan(v string) *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoListScheduleInfo {
	s.NextPlan = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoListScheduleInfo) SetNextTime(v string) *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoListScheduleInfo {
	s.NextTime = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoListScheduleInfo) SetOncePlan(v string) *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoListScheduleInfo {
	s.OncePlan = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoListScheduleInfo) SetOnceTime(v string) *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoListScheduleInfo {
	s.OnceTime = &v
	return s
}

func (s *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoListScheduleInfo) SetOperatorName(v string) *GetQuotaPlanResponseBodyDataQuotaSubQuotaInfoListScheduleInfo {
	s.OperatorName = &v
	return s
}

type GetQuotaPlanResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetQuotaPlanResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetQuotaPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaPlanResponse) GoString() string {
	return s.String()
}

func (s *GetQuotaPlanResponse) SetHeaders(v map[string]*string) *GetQuotaPlanResponse {
	s.Headers = v
	return s
}

func (s *GetQuotaPlanResponse) SetStatusCode(v int32) *GetQuotaPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *GetQuotaPlanResponse) SetBody(v *GetQuotaPlanResponseBody) *GetQuotaPlanResponse {
	s.Body = v
	return s
}

type GetQuotaScheduleRequest struct {
	// The time zone.
	DisplayTimezone *string `json:"displayTimezone,omitempty" xml:"displayTimezone,omitempty"`
	// The ID of the region.
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// The ID of the tenant.
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s GetQuotaScheduleRequest) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaScheduleRequest) GoString() string {
	return s.String()
}

func (s *GetQuotaScheduleRequest) SetDisplayTimezone(v string) *GetQuotaScheduleRequest {
	s.DisplayTimezone = &v
	return s
}

func (s *GetQuotaScheduleRequest) SetRegion(v string) *GetQuotaScheduleRequest {
	s.Region = &v
	return s
}

func (s *GetQuotaScheduleRequest) SetTenantId(v string) *GetQuotaScheduleRequest {
	s.TenantId = &v
	return s
}

type GetQuotaScheduleResponseBody struct {
	// The returned data.
	Data []*GetQuotaScheduleResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// *   If the value of success was false, an error code was returned.
	// *   If the value of success was true, a null value was returned.
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// The error message.
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// Indicates whether the request was successful. If this parameter was not empty and the value of this parameter was not 200, the request failed.
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetQuotaScheduleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaScheduleResponseBody) GoString() string {
	return s.String()
}

func (s *GetQuotaScheduleResponseBody) SetData(v []*GetQuotaScheduleResponseBodyData) *GetQuotaScheduleResponseBody {
	s.Data = v
	return s
}

func (s *GetQuotaScheduleResponseBody) SetErrorCode(v string) *GetQuotaScheduleResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetQuotaScheduleResponseBody) SetErrorMsg(v string) *GetQuotaScheduleResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GetQuotaScheduleResponseBody) SetHttpCode(v int32) *GetQuotaScheduleResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetQuotaScheduleResponseBody) SetRequestId(v string) *GetQuotaScheduleResponseBody {
	s.RequestId = &v
	return s
}

type GetQuotaScheduleResponseBodyData struct {
	// The condition value.
	Condition *GetQuotaScheduleResponseBodyDataCondition `json:"condition,omitempty" xml:"condition,omitempty" type:"Struct"`
	// The ID of the quota plan.
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The name of the operator.
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
	// The name of the quota plan.
	Plan *string `json:"plan,omitempty" xml:"plan,omitempty"`
	// The time zone.
	Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
	// The type of the quota plan.
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetQuotaScheduleResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaScheduleResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetQuotaScheduleResponseBodyData) SetCondition(v *GetQuotaScheduleResponseBodyDataCondition) *GetQuotaScheduleResponseBodyData {
	s.Condition = v
	return s
}

func (s *GetQuotaScheduleResponseBodyData) SetId(v string) *GetQuotaScheduleResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetQuotaScheduleResponseBodyData) SetOperator(v string) *GetQuotaScheduleResponseBodyData {
	s.Operator = &v
	return s
}

func (s *GetQuotaScheduleResponseBodyData) SetPlan(v string) *GetQuotaScheduleResponseBodyData {
	s.Plan = &v
	return s
}

func (s *GetQuotaScheduleResponseBodyData) SetTimezone(v string) *GetQuotaScheduleResponseBodyData {
	s.Timezone = &v
	return s
}

func (s *GetQuotaScheduleResponseBodyData) SetType(v string) *GetQuotaScheduleResponseBodyData {
	s.Type = &v
	return s
}

type GetQuotaScheduleResponseBodyDataCondition struct {
	// The start time when the quota plan takes effect.
	After *string `json:"after,omitempty" xml:"after,omitempty"`
	// The time when the quota plan takes effect.
	At *string `json:"at,omitempty" xml:"at,omitempty"`
}

func (s GetQuotaScheduleResponseBodyDataCondition) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaScheduleResponseBodyDataCondition) GoString() string {
	return s.String()
}

func (s *GetQuotaScheduleResponseBodyDataCondition) SetAfter(v string) *GetQuotaScheduleResponseBodyDataCondition {
	s.After = &v
	return s
}

func (s *GetQuotaScheduleResponseBodyDataCondition) SetAt(v string) *GetQuotaScheduleResponseBodyDataCondition {
	s.At = &v
	return s
}

type GetQuotaScheduleResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetQuotaScheduleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetQuotaScheduleResponse) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaScheduleResponse) GoString() string {
	return s.String()
}

func (s *GetQuotaScheduleResponse) SetHeaders(v map[string]*string) *GetQuotaScheduleResponse {
	s.Headers = v
	return s
}

func (s *GetQuotaScheduleResponse) SetStatusCode(v int32) *GetQuotaScheduleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetQuotaScheduleResponse) SetBody(v *GetQuotaScheduleResponseBody) *GetQuotaScheduleResponse {
	s.Body = v
	return s
}

type GetRoleAclResponseBody struct {
	// The returned data.
	Data      *GetRoleAclResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	ErrorCode *string                     `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMsg  *string                     `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	HttpCode  *int32                      `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetRoleAclResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRoleAclResponseBody) GoString() string {
	return s.String()
}

func (s *GetRoleAclResponseBody) SetData(v *GetRoleAclResponseBodyData) *GetRoleAclResponseBody {
	s.Data = v
	return s
}

func (s *GetRoleAclResponseBody) SetErrorCode(v string) *GetRoleAclResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetRoleAclResponseBody) SetErrorMsg(v string) *GetRoleAclResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GetRoleAclResponseBody) SetHttpCode(v int32) *GetRoleAclResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetRoleAclResponseBody) SetRequestId(v string) *GetRoleAclResponseBody {
	s.RequestId = &v
	return s
}

type GetRoleAclResponseBodyData struct {
	// The function.
	Function []*GetRoleAclResponseBodyDataFunction `json:"function,omitempty" xml:"function,omitempty" type:"Repeated"`
	// The instance.
	Instance []*GetRoleAclResponseBodyDataInstance `json:"instance,omitempty" xml:"instance,omitempty" type:"Repeated"`
	// The package.
	Package []*GetRoleAclResponseBodyDataPackage `json:"package,omitempty" xml:"package,omitempty" type:"Repeated"`
	// The project.
	Project []*GetRoleAclResponseBodyDataProject `json:"project,omitempty" xml:"project,omitempty" type:"Repeated"`
	// The resource.
	Resource []*GetRoleAclResponseBodyDataResource `json:"resource,omitempty" xml:"resource,omitempty" type:"Repeated"`
	// The table.
	Table []*GetRoleAclResponseBodyDataTable `json:"table,omitempty" xml:"table,omitempty" type:"Repeated"`
}

func (s GetRoleAclResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetRoleAclResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetRoleAclResponseBodyData) SetFunction(v []*GetRoleAclResponseBodyDataFunction) *GetRoleAclResponseBodyData {
	s.Function = v
	return s
}

func (s *GetRoleAclResponseBodyData) SetInstance(v []*GetRoleAclResponseBodyDataInstance) *GetRoleAclResponseBodyData {
	s.Instance = v
	return s
}

func (s *GetRoleAclResponseBodyData) SetPackage(v []*GetRoleAclResponseBodyDataPackage) *GetRoleAclResponseBodyData {
	s.Package = v
	return s
}

func (s *GetRoleAclResponseBodyData) SetProject(v []*GetRoleAclResponseBodyDataProject) *GetRoleAclResponseBodyData {
	s.Project = v
	return s
}

func (s *GetRoleAclResponseBodyData) SetResource(v []*GetRoleAclResponseBodyDataResource) *GetRoleAclResponseBodyData {
	s.Resource = v
	return s
}

func (s *GetRoleAclResponseBodyData) SetTable(v []*GetRoleAclResponseBodyDataTable) *GetRoleAclResponseBodyData {
	s.Table = v
	return s
}

type GetRoleAclResponseBodyDataFunction struct {
	// The operations that were performed on the function.
	Actions []*string `json:"actions,omitempty" xml:"actions,omitempty" type:"Repeated"`
	// The name of the function.
	Name       *string `json:"name,omitempty" xml:"name,omitempty"`
	SchemaName *string `json:"schemaName,omitempty" xml:"schemaName,omitempty"`
}

func (s GetRoleAclResponseBodyDataFunction) String() string {
	return tea.Prettify(s)
}

func (s GetRoleAclResponseBodyDataFunction) GoString() string {
	return s.String()
}

func (s *GetRoleAclResponseBodyDataFunction) SetActions(v []*string) *GetRoleAclResponseBodyDataFunction {
	s.Actions = v
	return s
}

func (s *GetRoleAclResponseBodyDataFunction) SetName(v string) *GetRoleAclResponseBodyDataFunction {
	s.Name = &v
	return s
}

func (s *GetRoleAclResponseBodyDataFunction) SetSchemaName(v string) *GetRoleAclResponseBodyDataFunction {
	s.SchemaName = &v
	return s
}

type GetRoleAclResponseBodyDataInstance struct {
	// The operations that were performed on the instance.
	Actions []*string `json:"actions,omitempty" xml:"actions,omitempty" type:"Repeated"`
	// The name of the instance.
	Name       *string `json:"name,omitempty" xml:"name,omitempty"`
	SchemaName *string `json:"schemaName,omitempty" xml:"schemaName,omitempty"`
}

func (s GetRoleAclResponseBodyDataInstance) String() string {
	return tea.Prettify(s)
}

func (s GetRoleAclResponseBodyDataInstance) GoString() string {
	return s.String()
}

func (s *GetRoleAclResponseBodyDataInstance) SetActions(v []*string) *GetRoleAclResponseBodyDataInstance {
	s.Actions = v
	return s
}

func (s *GetRoleAclResponseBodyDataInstance) SetName(v string) *GetRoleAclResponseBodyDataInstance {
	s.Name = &v
	return s
}

func (s *GetRoleAclResponseBodyDataInstance) SetSchemaName(v string) *GetRoleAclResponseBodyDataInstance {
	s.SchemaName = &v
	return s
}

type GetRoleAclResponseBodyDataPackage struct {
	// The operations that were performed on the package.
	Actions []*string `json:"actions,omitempty" xml:"actions,omitempty" type:"Repeated"`
	// The name of the package.
	Name       *string `json:"name,omitempty" xml:"name,omitempty"`
	SchemaName *string `json:"schemaName,omitempty" xml:"schemaName,omitempty"`
}

func (s GetRoleAclResponseBodyDataPackage) String() string {
	return tea.Prettify(s)
}

func (s GetRoleAclResponseBodyDataPackage) GoString() string {
	return s.String()
}

func (s *GetRoleAclResponseBodyDataPackage) SetActions(v []*string) *GetRoleAclResponseBodyDataPackage {
	s.Actions = v
	return s
}

func (s *GetRoleAclResponseBodyDataPackage) SetName(v string) *GetRoleAclResponseBodyDataPackage {
	s.Name = &v
	return s
}

func (s *GetRoleAclResponseBodyDataPackage) SetSchemaName(v string) *GetRoleAclResponseBodyDataPackage {
	s.SchemaName = &v
	return s
}

type GetRoleAclResponseBodyDataProject struct {
	// The operations that were performed on the project.
	Actions []*string `json:"actions,omitempty" xml:"actions,omitempty" type:"Repeated"`
	// The name of the MaxCompute project.
	Name       *string `json:"name,omitempty" xml:"name,omitempty"`
	SchemaName *string `json:"schemaName,omitempty" xml:"schemaName,omitempty"`
}

func (s GetRoleAclResponseBodyDataProject) String() string {
	return tea.Prettify(s)
}

func (s GetRoleAclResponseBodyDataProject) GoString() string {
	return s.String()
}

func (s *GetRoleAclResponseBodyDataProject) SetActions(v []*string) *GetRoleAclResponseBodyDataProject {
	s.Actions = v
	return s
}

func (s *GetRoleAclResponseBodyDataProject) SetName(v string) *GetRoleAclResponseBodyDataProject {
	s.Name = &v
	return s
}

func (s *GetRoleAclResponseBodyDataProject) SetSchemaName(v string) *GetRoleAclResponseBodyDataProject {
	s.SchemaName = &v
	return s
}

type GetRoleAclResponseBodyDataResource struct {
	// The operations that were performed on the resource.
	Actions []*string `json:"actions,omitempty" xml:"actions,omitempty" type:"Repeated"`
	// The name of the resource.
	Name       *string `json:"name,omitempty" xml:"name,omitempty"`
	SchemaName *string `json:"schemaName,omitempty" xml:"schemaName,omitempty"`
}

func (s GetRoleAclResponseBodyDataResource) String() string {
	return tea.Prettify(s)
}

func (s GetRoleAclResponseBodyDataResource) GoString() string {
	return s.String()
}

func (s *GetRoleAclResponseBodyDataResource) SetActions(v []*string) *GetRoleAclResponseBodyDataResource {
	s.Actions = v
	return s
}

func (s *GetRoleAclResponseBodyDataResource) SetName(v string) *GetRoleAclResponseBodyDataResource {
	s.Name = &v
	return s
}

func (s *GetRoleAclResponseBodyDataResource) SetSchemaName(v string) *GetRoleAclResponseBodyDataResource {
	s.SchemaName = &v
	return s
}

type GetRoleAclResponseBodyDataTable struct {
	// The operations that were performed on the table.
	Actions []*string `json:"actions,omitempty" xml:"actions,omitempty" type:"Repeated"`
	// The name of the table.
	Name       *string `json:"name,omitempty" xml:"name,omitempty"`
	SchemaName *string `json:"schemaName,omitempty" xml:"schemaName,omitempty"`
}

func (s GetRoleAclResponseBodyDataTable) String() string {
	return tea.Prettify(s)
}

func (s GetRoleAclResponseBodyDataTable) GoString() string {
	return s.String()
}

func (s *GetRoleAclResponseBodyDataTable) SetActions(v []*string) *GetRoleAclResponseBodyDataTable {
	s.Actions = v
	return s
}

func (s *GetRoleAclResponseBodyDataTable) SetName(v string) *GetRoleAclResponseBodyDataTable {
	s.Name = &v
	return s
}

func (s *GetRoleAclResponseBodyDataTable) SetSchemaName(v string) *GetRoleAclResponseBodyDataTable {
	s.SchemaName = &v
	return s
}

type GetRoleAclResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetRoleAclResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetRoleAclResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRoleAclResponse) GoString() string {
	return s.String()
}

func (s *GetRoleAclResponse) SetHeaders(v map[string]*string) *GetRoleAclResponse {
	s.Headers = v
	return s
}

func (s *GetRoleAclResponse) SetStatusCode(v int32) *GetRoleAclResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRoleAclResponse) SetBody(v *GetRoleAclResponseBody) *GetRoleAclResponse {
	s.Body = v
	return s
}

type GetRoleAclOnObjectRequest struct {
	// The name of the object.
	ObjectName *string `json:"objectName,omitempty" xml:"objectName,omitempty"`
	// The type of the object.
	ObjectType *string `json:"objectType,omitempty" xml:"objectType,omitempty"`
}

func (s GetRoleAclOnObjectRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRoleAclOnObjectRequest) GoString() string {
	return s.String()
}

func (s *GetRoleAclOnObjectRequest) SetObjectName(v string) *GetRoleAclOnObjectRequest {
	s.ObjectName = &v
	return s
}

func (s *GetRoleAclOnObjectRequest) SetObjectType(v string) *GetRoleAclOnObjectRequest {
	s.ObjectType = &v
	return s
}

type GetRoleAclOnObjectResponseBody struct {
	// The returned data.
	Data *GetRoleAclOnObjectResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetRoleAclOnObjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRoleAclOnObjectResponseBody) GoString() string {
	return s.String()
}

func (s *GetRoleAclOnObjectResponseBody) SetData(v *GetRoleAclOnObjectResponseBodyData) *GetRoleAclOnObjectResponseBody {
	s.Data = v
	return s
}

func (s *GetRoleAclOnObjectResponseBody) SetRequestId(v string) *GetRoleAclOnObjectResponseBody {
	s.RequestId = &v
	return s
}

type GetRoleAclOnObjectResponseBodyData struct {
	// The operations that were performed on the object.
	Actions []*string `json:"actions,omitempty" xml:"actions,omitempty" type:"Repeated"`
	// The name of the object.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s GetRoleAclOnObjectResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetRoleAclOnObjectResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetRoleAclOnObjectResponseBodyData) SetActions(v []*string) *GetRoleAclOnObjectResponseBodyData {
	s.Actions = v
	return s
}

func (s *GetRoleAclOnObjectResponseBodyData) SetName(v string) *GetRoleAclOnObjectResponseBodyData {
	s.Name = &v
	return s
}

type GetRoleAclOnObjectResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetRoleAclOnObjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetRoleAclOnObjectResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRoleAclOnObjectResponse) GoString() string {
	return s.String()
}

func (s *GetRoleAclOnObjectResponse) SetHeaders(v map[string]*string) *GetRoleAclOnObjectResponse {
	s.Headers = v
	return s
}

func (s *GetRoleAclOnObjectResponse) SetStatusCode(v int32) *GetRoleAclOnObjectResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRoleAclOnObjectResponse) SetBody(v *GetRoleAclOnObjectResponseBody) *GetRoleAclOnObjectResponse {
	s.Body = v
	return s
}

type GetRolePolicyResponseBody struct {
	// The returned data.
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetRolePolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRolePolicyResponseBody) GoString() string {
	return s.String()
}

func (s *GetRolePolicyResponseBody) SetData(v string) *GetRolePolicyResponseBody {
	s.Data = &v
	return s
}

func (s *GetRolePolicyResponseBody) SetRequestId(v string) *GetRolePolicyResponseBody {
	s.RequestId = &v
	return s
}

type GetRolePolicyResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetRolePolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetRolePolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRolePolicyResponse) GoString() string {
	return s.String()
}

func (s *GetRolePolicyResponse) SetHeaders(v map[string]*string) *GetRolePolicyResponse {
	s.Headers = v
	return s
}

func (s *GetRolePolicyResponse) SetStatusCode(v int32) *GetRolePolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRolePolicyResponse) SetBody(v *GetRolePolicyResponseBody) *GetRolePolicyResponse {
	s.Body = v
	return s
}

type GetRunningJobsRequest struct {
	// The time when the query starts. This parameter specifies the time when a job is submitted.
	//
	// *   The time range that is specified by the **from** and **to** request parameters is a closed interval. The start time and end time are included in the range. If the value of **from** is the same as the value of **to**, the time range is invalid, and a null value is returned.
	// *   The value is a UNIX timestamp that represents the number of seconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	From *int64 `json:"from,omitempty" xml:"from,omitempty"`
	// The list of job executors.
	JobOwnerList []*string `json:"jobOwnerList,omitempty" xml:"jobOwnerList,omitempty" type:"Repeated"`
	// The page number.
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries per page. Default value: 10. Maximum value: 20.
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The list of nicknames of quotas that are used by jobs.
	QuotaNicknameList []*string `json:"quotaNicknameList,omitempty" xml:"quotaNicknameList,omitempty" type:"Repeated"`
	// The time when the query ends. This parameter specifies the time when a job is submitted.
	//
	// *   The time interval that is specified by the **from** and **to** request parameters is a closed interval. The start time and end time are included in the interval. If the value of **from** is the same as the value of **to**, the interval is invalid, and a null value is returned.
	// *   The value is a UNIX timestamp that represents the number of seconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	To *int64 `json:"to,omitempty" xml:"to,omitempty"`
}

func (s GetRunningJobsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRunningJobsRequest) GoString() string {
	return s.String()
}

func (s *GetRunningJobsRequest) SetFrom(v int64) *GetRunningJobsRequest {
	s.From = &v
	return s
}

func (s *GetRunningJobsRequest) SetJobOwnerList(v []*string) *GetRunningJobsRequest {
	s.JobOwnerList = v
	return s
}

func (s *GetRunningJobsRequest) SetPageNumber(v int64) *GetRunningJobsRequest {
	s.PageNumber = &v
	return s
}

func (s *GetRunningJobsRequest) SetPageSize(v int64) *GetRunningJobsRequest {
	s.PageSize = &v
	return s
}

func (s *GetRunningJobsRequest) SetQuotaNicknameList(v []*string) *GetRunningJobsRequest {
	s.QuotaNicknameList = v
	return s
}

func (s *GetRunningJobsRequest) SetTo(v int64) *GetRunningJobsRequest {
	s.To = &v
	return s
}

type GetRunningJobsShrinkRequest struct {
	// The time when the query starts. This parameter specifies the time when a job is submitted.
	//
	// *   The time range that is specified by the **from** and **to** request parameters is a closed interval. The start time and end time are included in the range. If the value of **from** is the same as the value of **to**, the time range is invalid, and a null value is returned.
	// *   The value is a UNIX timestamp that represents the number of seconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	From *int64 `json:"from,omitempty" xml:"from,omitempty"`
	// The list of job executors.
	JobOwnerListShrink *string `json:"jobOwnerList,omitempty" xml:"jobOwnerList,omitempty"`
	// The page number.
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries per page. Default value: 10. Maximum value: 20.
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The list of nicknames of quotas that are used by jobs.
	QuotaNicknameListShrink *string `json:"quotaNicknameList,omitempty" xml:"quotaNicknameList,omitempty"`
	// The time when the query ends. This parameter specifies the time when a job is submitted.
	//
	// *   The time interval that is specified by the **from** and **to** request parameters is a closed interval. The start time and end time are included in the interval. If the value of **from** is the same as the value of **to**, the interval is invalid, and a null value is returned.
	// *   The value is a UNIX timestamp that represents the number of seconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	To *int64 `json:"to,omitempty" xml:"to,omitempty"`
}

func (s GetRunningJobsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRunningJobsShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetRunningJobsShrinkRequest) SetFrom(v int64) *GetRunningJobsShrinkRequest {
	s.From = &v
	return s
}

func (s *GetRunningJobsShrinkRequest) SetJobOwnerListShrink(v string) *GetRunningJobsShrinkRequest {
	s.JobOwnerListShrink = &v
	return s
}

func (s *GetRunningJobsShrinkRequest) SetPageNumber(v int64) *GetRunningJobsShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *GetRunningJobsShrinkRequest) SetPageSize(v int64) *GetRunningJobsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *GetRunningJobsShrinkRequest) SetQuotaNicknameListShrink(v string) *GetRunningJobsShrinkRequest {
	s.QuotaNicknameListShrink = &v
	return s
}

func (s *GetRunningJobsShrinkRequest) SetTo(v int64) *GetRunningJobsShrinkRequest {
	s.To = &v
	return s
}

type GetRunningJobsResponseBody struct {
	// The returned data.
	Data *GetRunningJobsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// *   If the value of success was false, an error code was returned.
	// *   If the value of success was true, a null value was returned.
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// The error message.
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// Indicates whether the request was successful. If this parameter was not empty and the value of this parameter was not 200, the request failed.
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetRunningJobsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRunningJobsResponseBody) GoString() string {
	return s.String()
}

func (s *GetRunningJobsResponseBody) SetData(v *GetRunningJobsResponseBodyData) *GetRunningJobsResponseBody {
	s.Data = v
	return s
}

func (s *GetRunningJobsResponseBody) SetErrorCode(v string) *GetRunningJobsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetRunningJobsResponseBody) SetErrorMsg(v string) *GetRunningJobsResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GetRunningJobsResponseBody) SetHttpCode(v int32) *GetRunningJobsResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetRunningJobsResponseBody) SetRequestId(v string) *GetRunningJobsResponseBody {
	s.RequestId = &v
	return s
}

type GetRunningJobsResponseBodyData struct {
	// The page number.
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries per page.
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The list of jobs in the running state.
	RunningJobInfoList []*GetRunningJobsResponseBodyDataRunningJobInfoList `json:"runningJobInfoList,omitempty" xml:"runningJobInfoList,omitempty" type:"Repeated"`
	// The total number of returned entries.
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s GetRunningJobsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetRunningJobsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetRunningJobsResponseBodyData) SetPageNumber(v int64) *GetRunningJobsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *GetRunningJobsResponseBodyData) SetPageSize(v int64) *GetRunningJobsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetRunningJobsResponseBodyData) SetRunningJobInfoList(v []*GetRunningJobsResponseBodyDataRunningJobInfoList) *GetRunningJobsResponseBodyData {
	s.RunningJobInfoList = v
	return s
}

func (s *GetRunningJobsResponseBodyData) SetTotalCount(v int64) *GetRunningJobsResponseBodyData {
	s.TotalCount = &v
	return s
}

type GetRunningJobsResponseBodyDataRunningJobInfoList struct {
	// The compute unit (CU) snapshot proportion of the job.
	CuSnapshot *float64 `json:"cuSnapshot,omitempty" xml:"cuSnapshot,omitempty"`
	// The instance ID.
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The account that submits the job.
	JobOwner *string `json:"jobOwner,omitempty" xml:"jobOwner,omitempty"`
	// The memory snapshot proportion of the job.
	MemorySnapshot *float64 `json:"memorySnapshot,omitempty" xml:"memorySnapshot,omitempty"`
	// The progress of the job.
	Progress *float64 `json:"progress,omitempty" xml:"progress,omitempty"`
	// The name of the MaxCompute project.
	Project *string `json:"project,omitempty" xml:"project,omitempty"`
	// The nickname of the quota that is used by the job.
	QuotaNickname *string `json:"quotaNickname,omitempty" xml:"quotaNickname,omitempty"`
	// The time when the job starts to run.
	RunningAtTime *int64 `json:"runningAtTime,omitempty" xml:"runningAtTime,omitempty"`
	// The time when the job is submitted.
	SubmittedAtTime *int64 `json:"submittedAtTime,omitempty" xml:"submittedAtTime,omitempty"`
}

func (s GetRunningJobsResponseBodyDataRunningJobInfoList) String() string {
	return tea.Prettify(s)
}

func (s GetRunningJobsResponseBodyDataRunningJobInfoList) GoString() string {
	return s.String()
}

func (s *GetRunningJobsResponseBodyDataRunningJobInfoList) SetCuSnapshot(v float64) *GetRunningJobsResponseBodyDataRunningJobInfoList {
	s.CuSnapshot = &v
	return s
}

func (s *GetRunningJobsResponseBodyDataRunningJobInfoList) SetInstanceId(v string) *GetRunningJobsResponseBodyDataRunningJobInfoList {
	s.InstanceId = &v
	return s
}

func (s *GetRunningJobsResponseBodyDataRunningJobInfoList) SetJobOwner(v string) *GetRunningJobsResponseBodyDataRunningJobInfoList {
	s.JobOwner = &v
	return s
}

func (s *GetRunningJobsResponseBodyDataRunningJobInfoList) SetMemorySnapshot(v float64) *GetRunningJobsResponseBodyDataRunningJobInfoList {
	s.MemorySnapshot = &v
	return s
}

func (s *GetRunningJobsResponseBodyDataRunningJobInfoList) SetProgress(v float64) *GetRunningJobsResponseBodyDataRunningJobInfoList {
	s.Progress = &v
	return s
}

func (s *GetRunningJobsResponseBodyDataRunningJobInfoList) SetProject(v string) *GetRunningJobsResponseBodyDataRunningJobInfoList {
	s.Project = &v
	return s
}

func (s *GetRunningJobsResponseBodyDataRunningJobInfoList) SetQuotaNickname(v string) *GetRunningJobsResponseBodyDataRunningJobInfoList {
	s.QuotaNickname = &v
	return s
}

func (s *GetRunningJobsResponseBodyDataRunningJobInfoList) SetRunningAtTime(v int64) *GetRunningJobsResponseBodyDataRunningJobInfoList {
	s.RunningAtTime = &v
	return s
}

func (s *GetRunningJobsResponseBodyDataRunningJobInfoList) SetSubmittedAtTime(v int64) *GetRunningJobsResponseBodyDataRunningJobInfoList {
	s.SubmittedAtTime = &v
	return s
}

type GetRunningJobsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetRunningJobsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetRunningJobsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRunningJobsResponse) GoString() string {
	return s.String()
}

func (s *GetRunningJobsResponse) SetHeaders(v map[string]*string) *GetRunningJobsResponse {
	s.Headers = v
	return s
}

func (s *GetRunningJobsResponse) SetStatusCode(v int32) *GetRunningJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRunningJobsResponse) SetBody(v *GetRunningJobsResponseBody) *GetRunningJobsResponse {
	s.Body = v
	return s
}

type GetTrustedProjectsResponseBody struct {
	// The returned data.
	Data []*string `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetTrustedProjectsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTrustedProjectsResponseBody) GoString() string {
	return s.String()
}

func (s *GetTrustedProjectsResponseBody) SetData(v []*string) *GetTrustedProjectsResponseBody {
	s.Data = v
	return s
}

func (s *GetTrustedProjectsResponseBody) SetRequestId(v string) *GetTrustedProjectsResponseBody {
	s.RequestId = &v
	return s
}

type GetTrustedProjectsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetTrustedProjectsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTrustedProjectsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTrustedProjectsResponse) GoString() string {
	return s.String()
}

func (s *GetTrustedProjectsResponse) SetHeaders(v map[string]*string) *GetTrustedProjectsResponse {
	s.Headers = v
	return s
}

func (s *GetTrustedProjectsResponse) SetStatusCode(v int32) *GetTrustedProjectsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTrustedProjectsResponse) SetBody(v *GetTrustedProjectsResponseBody) *GetTrustedProjectsResponse {
	s.Body = v
	return s
}

type KillJobsRequest struct {
	// The request body parameters.
	Body *string `json:"body,omitempty" xml:"body,omitempty"`
	// The ID of the region in which the instance resides.
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// The ID of the tenant.
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s KillJobsRequest) String() string {
	return tea.Prettify(s)
}

func (s KillJobsRequest) GoString() string {
	return s.String()
}

func (s *KillJobsRequest) SetBody(v string) *KillJobsRequest {
	s.Body = &v
	return s
}

func (s *KillJobsRequest) SetRegion(v string) *KillJobsRequest {
	s.Region = &v
	return s
}

func (s *KillJobsRequest) SetTenantId(v string) *KillJobsRequest {
	s.TenantId = &v
	return s
}

type KillJobsResponseBody struct {
	// The returned data.
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// Indicates whether the request was successful. If this parameter was not empty and the value of this parameter was not 200, the request failed.
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s KillJobsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s KillJobsResponseBody) GoString() string {
	return s.String()
}

func (s *KillJobsResponseBody) SetData(v string) *KillJobsResponseBody {
	s.Data = &v
	return s
}

func (s *KillJobsResponseBody) SetHttpCode(v int32) *KillJobsResponseBody {
	s.HttpCode = &v
	return s
}

func (s *KillJobsResponseBody) SetRequestId(v string) *KillJobsResponseBody {
	s.RequestId = &v
	return s
}

type KillJobsResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *KillJobsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s KillJobsResponse) String() string {
	return tea.Prettify(s)
}

func (s KillJobsResponse) GoString() string {
	return s.String()
}

func (s *KillJobsResponse) SetHeaders(v map[string]*string) *KillJobsResponse {
	s.Headers = v
	return s
}

func (s *KillJobsResponse) SetStatusCode(v int32) *KillJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *KillJobsResponse) SetBody(v *KillJobsResponseBody) *KillJobsResponse {
	s.Body = v
	return s
}

type ListFunctionsRequest struct {
	// Specifies the marker after which the returned list begins.
	Marker *string `json:"marker,omitempty" xml:"marker,omitempty"`
	// The maximum number of entries to return on each page.
	MaxItem *int32 `json:"maxItem,omitempty" xml:"maxItem,omitempty"`
	// The names of the returned resources. The names must start with the value specified by the prefix parameter. If the prefix parameter is set to a, the names of the returned resources must start with a.
	Prefix     *string `json:"prefix,omitempty" xml:"prefix,omitempty"`
	SchemaName *string `json:"schemaName,omitempty" xml:"schemaName,omitempty"`
}

func (s ListFunctionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionsRequest) GoString() string {
	return s.String()
}

func (s *ListFunctionsRequest) SetMarker(v string) *ListFunctionsRequest {
	s.Marker = &v
	return s
}

func (s *ListFunctionsRequest) SetMaxItem(v int32) *ListFunctionsRequest {
	s.MaxItem = &v
	return s
}

func (s *ListFunctionsRequest) SetPrefix(v string) *ListFunctionsRequest {
	s.Prefix = &v
	return s
}

func (s *ListFunctionsRequest) SetSchemaName(v string) *ListFunctionsRequest {
	s.SchemaName = &v
	return s
}

type ListFunctionsResponseBody struct {
	// The returned data.
	Data *ListFunctionsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListFunctionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListFunctionsResponseBody) SetData(v *ListFunctionsResponseBodyData) *ListFunctionsResponseBody {
	s.Data = v
	return s
}

func (s *ListFunctionsResponseBody) SetRequestId(v string) *ListFunctionsResponseBody {
	s.RequestId = &v
	return s
}

type ListFunctionsResponseBodyData struct {
	// The information about each function.
	Functions []*ListFunctionsResponseBodyDataFunctions `json:"functions,omitempty" xml:"functions,omitempty" type:"Repeated"`
	// Indicates the marker after which the returned list begins.
	Marker *string `json:"marker,omitempty" xml:"marker,omitempty"`
	// The maximum number of entries returned per page.
	MaxItem *int32 `json:"maxItem,omitempty" xml:"maxItem,omitempty"`
}

func (s ListFunctionsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListFunctionsResponseBodyData) SetFunctions(v []*ListFunctionsResponseBodyDataFunctions) *ListFunctionsResponseBodyData {
	s.Functions = v
	return s
}

func (s *ListFunctionsResponseBodyData) SetMarker(v string) *ListFunctionsResponseBodyData {
	s.Marker = &v
	return s
}

func (s *ListFunctionsResponseBodyData) SetMaxItem(v int32) *ListFunctionsResponseBodyData {
	s.MaxItem = &v
	return s
}

type ListFunctionsResponseBodyDataFunctions struct {
	// The class in which the function was defined.
	Class *string `json:"class,omitempty" xml:"class,omitempty"`
	// The time when the function was created. Unit: milliseconds.
	CreationTime *int64  `json:"creationTime,omitempty" xml:"creationTime,omitempty"`
	DisplayName  *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// The name of the function.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The owner of the function.
	Owner *string `json:"owner,omitempty" xml:"owner,omitempty"`
	// The name of the resource that was associated with the function.
	Resources *string `json:"resources,omitempty" xml:"resources,omitempty"`
	// The schema of the function.
	Schema *string `json:"schema,omitempty" xml:"schema,omitempty"`
}

func (s ListFunctionsResponseBodyDataFunctions) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionsResponseBodyDataFunctions) GoString() string {
	return s.String()
}

func (s *ListFunctionsResponseBodyDataFunctions) SetClass(v string) *ListFunctionsResponseBodyDataFunctions {
	s.Class = &v
	return s
}

func (s *ListFunctionsResponseBodyDataFunctions) SetCreationTime(v int64) *ListFunctionsResponseBodyDataFunctions {
	s.CreationTime = &v
	return s
}

func (s *ListFunctionsResponseBodyDataFunctions) SetDisplayName(v string) *ListFunctionsResponseBodyDataFunctions {
	s.DisplayName = &v
	return s
}

func (s *ListFunctionsResponseBodyDataFunctions) SetName(v string) *ListFunctionsResponseBodyDataFunctions {
	s.Name = &v
	return s
}

func (s *ListFunctionsResponseBodyDataFunctions) SetOwner(v string) *ListFunctionsResponseBodyDataFunctions {
	s.Owner = &v
	return s
}

func (s *ListFunctionsResponseBodyDataFunctions) SetResources(v string) *ListFunctionsResponseBodyDataFunctions {
	s.Resources = &v
	return s
}

func (s *ListFunctionsResponseBodyDataFunctions) SetSchema(v string) *ListFunctionsResponseBodyDataFunctions {
	s.Schema = &v
	return s
}

type ListFunctionsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListFunctionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListFunctionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionsResponse) GoString() string {
	return s.String()
}

func (s *ListFunctionsResponse) SetHeaders(v map[string]*string) *ListFunctionsResponse {
	s.Headers = v
	return s
}

func (s *ListFunctionsResponse) SetStatusCode(v int32) *ListFunctionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFunctionsResponse) SetBody(v *ListFunctionsResponseBody) *ListFunctionsResponse {
	s.Body = v
	return s
}

type ListJobInfosRequest struct {
	AscOrder    *bool   `json:"ascOrder,omitempty" xml:"ascOrder,omitempty"`
	Body        *string `json:"body,omitempty" xml:"body,omitempty"`
	OrderColumn *string `json:"orderColumn,omitempty" xml:"orderColumn,omitempty"`
	PageNumber  *int64  `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize    *int64  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Region      *string `json:"region,omitempty" xml:"region,omitempty"`
	TenantId    *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ListJobInfosRequest) String() string {
	return tea.Prettify(s)
}

func (s ListJobInfosRequest) GoString() string {
	return s.String()
}

func (s *ListJobInfosRequest) SetAscOrder(v bool) *ListJobInfosRequest {
	s.AscOrder = &v
	return s
}

func (s *ListJobInfosRequest) SetBody(v string) *ListJobInfosRequest {
	s.Body = &v
	return s
}

func (s *ListJobInfosRequest) SetOrderColumn(v string) *ListJobInfosRequest {
	s.OrderColumn = &v
	return s
}

func (s *ListJobInfosRequest) SetPageNumber(v int64) *ListJobInfosRequest {
	s.PageNumber = &v
	return s
}

func (s *ListJobInfosRequest) SetPageSize(v int64) *ListJobInfosRequest {
	s.PageSize = &v
	return s
}

func (s *ListJobInfosRequest) SetRegion(v string) *ListJobInfosRequest {
	s.Region = &v
	return s
}

func (s *ListJobInfosRequest) SetTenantId(v string) *ListJobInfosRequest {
	s.TenantId = &v
	return s
}

type ListJobInfosResponseBody struct {
	Data      *ListJobInfosResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	HttpCode  *int32                        `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	RequestId *string                       `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListJobInfosResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListJobInfosResponseBody) GoString() string {
	return s.String()
}

func (s *ListJobInfosResponseBody) SetData(v *ListJobInfosResponseBodyData) *ListJobInfosResponseBody {
	s.Data = v
	return s
}

func (s *ListJobInfosResponseBody) SetHttpCode(v int32) *ListJobInfosResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ListJobInfosResponseBody) SetRequestId(v string) *ListJobInfosResponseBody {
	s.RequestId = &v
	return s
}

type ListJobInfosResponseBodyData struct {
	JobInfoList []*ListJobInfosResponseBodyDataJobInfoList `json:"jobInfoList,omitempty" xml:"jobInfoList,omitempty" type:"Repeated"`
	PageNumber  *int64                                     `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize    *int64                                     `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	TotalCount  *int64                                     `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListJobInfosResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListJobInfosResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListJobInfosResponseBodyData) SetJobInfoList(v []*ListJobInfosResponseBodyDataJobInfoList) *ListJobInfosResponseBodyData {
	s.JobInfoList = v
	return s
}

func (s *ListJobInfosResponseBodyData) SetPageNumber(v int64) *ListJobInfosResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListJobInfosResponseBodyData) SetPageSize(v int64) *ListJobInfosResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListJobInfosResponseBodyData) SetTotalCount(v int64) *ListJobInfosResponseBodyData {
	s.TotalCount = &v
	return s
}

type ListJobInfosResponseBodyDataJobInfoList struct {
	Cluster         *string  `json:"cluster,omitempty" xml:"cluster,omitempty"`
	CuSnapshot      *float64 `json:"cuSnapshot,omitempty" xml:"cuSnapshot,omitempty"`
	CuUsage         *int64   `json:"cuUsage,omitempty" xml:"cuUsage,omitempty"`
	EndAtTime       *int64   `json:"endAtTime,omitempty" xml:"endAtTime,omitempty"`
	ExtNodeId       *string  `json:"extNodeId,omitempty" xml:"extNodeId,omitempty"`
	ExtNodeOnDuty   *string  `json:"extNodeOnDuty,omitempty" xml:"extNodeOnDuty,omitempty"`
	ExtPlantFrom    *string  `json:"extPlantFrom,omitempty" xml:"extPlantFrom,omitempty"`
	InstanceId      *string  `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	JobOwner        *string  `json:"jobOwner,omitempty" xml:"jobOwner,omitempty"`
	JobType         *string  `json:"jobType,omitempty" xml:"jobType,omitempty"`
	MemorySnapshot  *float64 `json:"memorySnapshot,omitempty" xml:"memorySnapshot,omitempty"`
	MemoryUsage     *int64   `json:"memoryUsage,omitempty" xml:"memoryUsage,omitempty"`
	Priority        *int64   `json:"priority,omitempty" xml:"priority,omitempty"`
	Project         *string  `json:"project,omitempty" xml:"project,omitempty"`
	QuotaNickname   *string  `json:"quotaNickname,omitempty" xml:"quotaNickname,omitempty"`
	QuotaType       *string  `json:"quotaType,omitempty" xml:"quotaType,omitempty"`
	Region          *string  `json:"region,omitempty" xml:"region,omitempty"`
	RunningAtTime   *int64   `json:"runningAtTime,omitempty" xml:"runningAtTime,omitempty"`
	RunningTime     *int64   `json:"runningTime,omitempty" xml:"runningTime,omitempty"`
	Signature       *string  `json:"signature,omitempty" xml:"signature,omitempty"`
	Status          *string  `json:"status,omitempty" xml:"status,omitempty"`
	StatusSnapshot  *string  `json:"statusSnapshot,omitempty" xml:"statusSnapshot,omitempty"`
	SubmittedAtTime *int64   `json:"submittedAtTime,omitempty" xml:"submittedAtTime,omitempty"`
	Tags            *string  `json:"tags,omitempty" xml:"tags,omitempty"`
	TenantId        *string  `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	TotalTime       *int64   `json:"totalTime,omitempty" xml:"totalTime,omitempty"`
	WaitingTime     *int64   `json:"waitingTime,omitempty" xml:"waitingTime,omitempty"`
}

func (s ListJobInfosResponseBodyDataJobInfoList) String() string {
	return tea.Prettify(s)
}

func (s ListJobInfosResponseBodyDataJobInfoList) GoString() string {
	return s.String()
}

func (s *ListJobInfosResponseBodyDataJobInfoList) SetCluster(v string) *ListJobInfosResponseBodyDataJobInfoList {
	s.Cluster = &v
	return s
}

func (s *ListJobInfosResponseBodyDataJobInfoList) SetCuSnapshot(v float64) *ListJobInfosResponseBodyDataJobInfoList {
	s.CuSnapshot = &v
	return s
}

func (s *ListJobInfosResponseBodyDataJobInfoList) SetCuUsage(v int64) *ListJobInfosResponseBodyDataJobInfoList {
	s.CuUsage = &v
	return s
}

func (s *ListJobInfosResponseBodyDataJobInfoList) SetEndAtTime(v int64) *ListJobInfosResponseBodyDataJobInfoList {
	s.EndAtTime = &v
	return s
}

func (s *ListJobInfosResponseBodyDataJobInfoList) SetExtNodeId(v string) *ListJobInfosResponseBodyDataJobInfoList {
	s.ExtNodeId = &v
	return s
}

func (s *ListJobInfosResponseBodyDataJobInfoList) SetExtNodeOnDuty(v string) *ListJobInfosResponseBodyDataJobInfoList {
	s.ExtNodeOnDuty = &v
	return s
}

func (s *ListJobInfosResponseBodyDataJobInfoList) SetExtPlantFrom(v string) *ListJobInfosResponseBodyDataJobInfoList {
	s.ExtPlantFrom = &v
	return s
}

func (s *ListJobInfosResponseBodyDataJobInfoList) SetInstanceId(v string) *ListJobInfosResponseBodyDataJobInfoList {
	s.InstanceId = &v
	return s
}

func (s *ListJobInfosResponseBodyDataJobInfoList) SetJobOwner(v string) *ListJobInfosResponseBodyDataJobInfoList {
	s.JobOwner = &v
	return s
}

func (s *ListJobInfosResponseBodyDataJobInfoList) SetJobType(v string) *ListJobInfosResponseBodyDataJobInfoList {
	s.JobType = &v
	return s
}

func (s *ListJobInfosResponseBodyDataJobInfoList) SetMemorySnapshot(v float64) *ListJobInfosResponseBodyDataJobInfoList {
	s.MemorySnapshot = &v
	return s
}

func (s *ListJobInfosResponseBodyDataJobInfoList) SetMemoryUsage(v int64) *ListJobInfosResponseBodyDataJobInfoList {
	s.MemoryUsage = &v
	return s
}

func (s *ListJobInfosResponseBodyDataJobInfoList) SetPriority(v int64) *ListJobInfosResponseBodyDataJobInfoList {
	s.Priority = &v
	return s
}

func (s *ListJobInfosResponseBodyDataJobInfoList) SetProject(v string) *ListJobInfosResponseBodyDataJobInfoList {
	s.Project = &v
	return s
}

func (s *ListJobInfosResponseBodyDataJobInfoList) SetQuotaNickname(v string) *ListJobInfosResponseBodyDataJobInfoList {
	s.QuotaNickname = &v
	return s
}

func (s *ListJobInfosResponseBodyDataJobInfoList) SetQuotaType(v string) *ListJobInfosResponseBodyDataJobInfoList {
	s.QuotaType = &v
	return s
}

func (s *ListJobInfosResponseBodyDataJobInfoList) SetRegion(v string) *ListJobInfosResponseBodyDataJobInfoList {
	s.Region = &v
	return s
}

func (s *ListJobInfosResponseBodyDataJobInfoList) SetRunningAtTime(v int64) *ListJobInfosResponseBodyDataJobInfoList {
	s.RunningAtTime = &v
	return s
}

func (s *ListJobInfosResponseBodyDataJobInfoList) SetRunningTime(v int64) *ListJobInfosResponseBodyDataJobInfoList {
	s.RunningTime = &v
	return s
}

func (s *ListJobInfosResponseBodyDataJobInfoList) SetSignature(v string) *ListJobInfosResponseBodyDataJobInfoList {
	s.Signature = &v
	return s
}

func (s *ListJobInfosResponseBodyDataJobInfoList) SetStatus(v string) *ListJobInfosResponseBodyDataJobInfoList {
	s.Status = &v
	return s
}

func (s *ListJobInfosResponseBodyDataJobInfoList) SetStatusSnapshot(v string) *ListJobInfosResponseBodyDataJobInfoList {
	s.StatusSnapshot = &v
	return s
}

func (s *ListJobInfosResponseBodyDataJobInfoList) SetSubmittedAtTime(v int64) *ListJobInfosResponseBodyDataJobInfoList {
	s.SubmittedAtTime = &v
	return s
}

func (s *ListJobInfosResponseBodyDataJobInfoList) SetTags(v string) *ListJobInfosResponseBodyDataJobInfoList {
	s.Tags = &v
	return s
}

func (s *ListJobInfosResponseBodyDataJobInfoList) SetTenantId(v string) *ListJobInfosResponseBodyDataJobInfoList {
	s.TenantId = &v
	return s
}

func (s *ListJobInfosResponseBodyDataJobInfoList) SetTotalTime(v int64) *ListJobInfosResponseBodyDataJobInfoList {
	s.TotalTime = &v
	return s
}

func (s *ListJobInfosResponseBodyDataJobInfoList) SetWaitingTime(v int64) *ListJobInfosResponseBodyDataJobInfoList {
	s.WaitingTime = &v
	return s
}

type ListJobInfosResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListJobInfosResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListJobInfosResponse) String() string {
	return tea.Prettify(s)
}

func (s ListJobInfosResponse) GoString() string {
	return s.String()
}

func (s *ListJobInfosResponse) SetHeaders(v map[string]*string) *ListJobInfosResponse {
	s.Headers = v
	return s
}

func (s *ListJobInfosResponse) SetStatusCode(v int32) *ListJobInfosResponse {
	s.StatusCode = &v
	return s
}

func (s *ListJobInfosResponse) SetBody(v *ListJobInfosResponseBody) *ListJobInfosResponse {
	s.Body = v
	return s
}

type ListPackagesResponseBody struct {
	// The returned data.
	Data *ListPackagesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListPackagesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPackagesResponseBody) GoString() string {
	return s.String()
}

func (s *ListPackagesResponseBody) SetData(v *ListPackagesResponseBodyData) *ListPackagesResponseBody {
	s.Data = v
	return s
}

func (s *ListPackagesResponseBody) SetRequestId(v string) *ListPackagesResponseBody {
	s.RequestId = &v
	return s
}

type ListPackagesResponseBodyData struct {
	// The packages that were created.
	CreatedPackages []*ListPackagesResponseBodyDataCreatedPackages `json:"createdPackages,omitempty" xml:"createdPackages,omitempty" type:"Repeated"`
	// The packages that were installed.
	InstalledPackages []*ListPackagesResponseBodyDataInstalledPackages `json:"installedPackages,omitempty" xml:"installedPackages,omitempty" type:"Repeated"`
}

func (s ListPackagesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListPackagesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListPackagesResponseBodyData) SetCreatedPackages(v []*ListPackagesResponseBodyDataCreatedPackages) *ListPackagesResponseBodyData {
	s.CreatedPackages = v
	return s
}

func (s *ListPackagesResponseBodyData) SetInstalledPackages(v []*ListPackagesResponseBodyDataInstalledPackages) *ListPackagesResponseBodyData {
	s.InstalledPackages = v
	return s
}

type ListPackagesResponseBodyDataCreatedPackages struct {
	// The time when the package was created.
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The name of the package.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ListPackagesResponseBodyDataCreatedPackages) String() string {
	return tea.Prettify(s)
}

func (s ListPackagesResponseBodyDataCreatedPackages) GoString() string {
	return s.String()
}

func (s *ListPackagesResponseBodyDataCreatedPackages) SetCreateTime(v int64) *ListPackagesResponseBodyDataCreatedPackages {
	s.CreateTime = &v
	return s
}

func (s *ListPackagesResponseBodyDataCreatedPackages) SetName(v string) *ListPackagesResponseBodyDataCreatedPackages {
	s.Name = &v
	return s
}

type ListPackagesResponseBodyDataInstalledPackages struct {
	// The time when the package was installed.
	InstallTime *int64 `json:"installTime,omitempty" xml:"installTime,omitempty"`
	// The name of the package.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The project to which the package belongs. This parameter is required if the package is installed in the MaxCompute project.
	SourceProject *string `json:"sourceProject,omitempty" xml:"sourceProject,omitempty"`
	// The status of the package.
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListPackagesResponseBodyDataInstalledPackages) String() string {
	return tea.Prettify(s)
}

func (s ListPackagesResponseBodyDataInstalledPackages) GoString() string {
	return s.String()
}

func (s *ListPackagesResponseBodyDataInstalledPackages) SetInstallTime(v int64) *ListPackagesResponseBodyDataInstalledPackages {
	s.InstallTime = &v
	return s
}

func (s *ListPackagesResponseBodyDataInstalledPackages) SetName(v string) *ListPackagesResponseBodyDataInstalledPackages {
	s.Name = &v
	return s
}

func (s *ListPackagesResponseBodyDataInstalledPackages) SetSourceProject(v string) *ListPackagesResponseBodyDataInstalledPackages {
	s.SourceProject = &v
	return s
}

func (s *ListPackagesResponseBodyDataInstalledPackages) SetStatus(v string) *ListPackagesResponseBodyDataInstalledPackages {
	s.Status = &v
	return s
}

type ListPackagesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListPackagesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListPackagesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPackagesResponse) GoString() string {
	return s.String()
}

func (s *ListPackagesResponse) SetHeaders(v map[string]*string) *ListPackagesResponse {
	s.Headers = v
	return s
}

func (s *ListPackagesResponse) SetStatusCode(v int32) *ListPackagesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPackagesResponse) SetBody(v *ListPackagesResponseBody) *ListPackagesResponse {
	s.Body = v
	return s
}

type ListProjectUsersResponseBody struct {
	// The returned data.
	Data *ListProjectUsersResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The request ID.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListProjectUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListProjectUsersResponseBody) GoString() string {
	return s.String()
}

func (s *ListProjectUsersResponseBody) SetData(v *ListProjectUsersResponseBodyData) *ListProjectUsersResponseBody {
	s.Data = v
	return s
}

func (s *ListProjectUsersResponseBody) SetRequestId(v string) *ListProjectUsersResponseBody {
	s.RequestId = &v
	return s
}

type ListProjectUsersResponseBodyData struct {
	// An array that contains users.
	Users []*ListProjectUsersResponseBodyDataUsers `json:"users,omitempty" xml:"users,omitempty" type:"Repeated"`
}

func (s ListProjectUsersResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListProjectUsersResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListProjectUsersResponseBodyData) SetUsers(v []*ListProjectUsersResponseBodyDataUsers) *ListProjectUsersResponseBodyData {
	s.Users = v
	return s
}

type ListProjectUsersResponseBodyDataUsers struct {
	// The name of the user.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ListProjectUsersResponseBodyDataUsers) String() string {
	return tea.Prettify(s)
}

func (s ListProjectUsersResponseBodyDataUsers) GoString() string {
	return s.String()
}

func (s *ListProjectUsersResponseBodyDataUsers) SetName(v string) *ListProjectUsersResponseBodyDataUsers {
	s.Name = &v
	return s
}

type ListProjectUsersResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListProjectUsersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListProjectUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListProjectUsersResponse) GoString() string {
	return s.String()
}

func (s *ListProjectUsersResponse) SetHeaders(v map[string]*string) *ListProjectUsersResponse {
	s.Headers = v
	return s
}

func (s *ListProjectUsersResponse) SetStatusCode(v int32) *ListProjectUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListProjectUsersResponse) SetBody(v *ListProjectUsersResponseBody) *ListProjectUsersResponse {
	s.Body = v
	return s
}

type ListProjectsRequest struct {
	// The flag to choose whether list the specific catalog - system_catalog.
	ListSystemCatalog *bool `json:"listSystemCatalog,omitempty" xml:"listSystemCatalog,omitempty"`
	// The maximum number of entries to return on each page.
	Marker *string `json:"marker,omitempty" xml:"marker,omitempty"`
	// The maximum number of entries returned per page.
	MaxItem *int32 `json:"maxItem,omitempty" xml:"maxItem,omitempty"`
	// Specifies the marker after which the returned list begins.
	Prefix *string `json:"prefix,omitempty" xml:"prefix,omitempty"`
	// The name of the quota. The value of this parameter is the identifier of the quota in MaxCompute, which differs from the quotaNickname parameter. You can configure the quotaNickname parameter. The system automatically generates a value for the quotaName parameter. This parameter is only used to describe the tunnel quota.
	QuotaName *string `json:"quotaName,omitempty" xml:"quotaName,omitempty"`
	// The name of the quota.
	QuotaNickName *string `json:"quotaNickName,omitempty" xml:"quotaNickName,omitempty"`
	// The ID of the region.
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// The identifier of an object in a MaxCompute quota. This identifier is the same as the identifier in the sales bill of Alibaba Cloud. This parameter is used for tags.
	SaleTags *string `json:"saleTags,omitempty" xml:"saleTags,omitempty"`
	// The ID of the tenant.
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The project type. Valid values: external and managed. The value external indicates an external project, which is used in the data lakehouse solution. The value managed indicates an internal project.
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListProjectsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListProjectsRequest) GoString() string {
	return s.String()
}

func (s *ListProjectsRequest) SetListSystemCatalog(v bool) *ListProjectsRequest {
	s.ListSystemCatalog = &v
	return s
}

func (s *ListProjectsRequest) SetMarker(v string) *ListProjectsRequest {
	s.Marker = &v
	return s
}

func (s *ListProjectsRequest) SetMaxItem(v int32) *ListProjectsRequest {
	s.MaxItem = &v
	return s
}

func (s *ListProjectsRequest) SetPrefix(v string) *ListProjectsRequest {
	s.Prefix = &v
	return s
}

func (s *ListProjectsRequest) SetQuotaName(v string) *ListProjectsRequest {
	s.QuotaName = &v
	return s
}

func (s *ListProjectsRequest) SetQuotaNickName(v string) *ListProjectsRequest {
	s.QuotaNickName = &v
	return s
}

func (s *ListProjectsRequest) SetRegion(v string) *ListProjectsRequest {
	s.Region = &v
	return s
}

func (s *ListProjectsRequest) SetSaleTags(v string) *ListProjectsRequest {
	s.SaleTags = &v
	return s
}

func (s *ListProjectsRequest) SetTenantId(v string) *ListProjectsRequest {
	s.TenantId = &v
	return s
}

func (s *ListProjectsRequest) SetType(v string) *ListProjectsRequest {
	s.Type = &v
	return s
}

type ListProjectsResponseBody struct {
	// The returned data.
	Data *ListProjectsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListProjectsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListProjectsResponseBody) GoString() string {
	return s.String()
}

func (s *ListProjectsResponseBody) SetData(v *ListProjectsResponseBodyData) *ListProjectsResponseBody {
	s.Data = v
	return s
}

func (s *ListProjectsResponseBody) SetRequestId(v string) *ListProjectsResponseBody {
	s.RequestId = &v
	return s
}

type ListProjectsResponseBodyData struct {
	// A pagination token. Only continuous page turning is supported. If NextToken is not empty, the next page exists. The value of NextToken can be used in the next request to retrieve a new page of results.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Indicates the marker after which the returned list begins.
	Marker *string `json:"marker,omitempty" xml:"marker,omitempty"`
	// The maximum number of entries returned per page.
	MaxItem *int32 `json:"maxItem,omitempty" xml:"maxItem,omitempty"`
	// The description of the project.
	Projects []*ListProjectsResponseBodyDataProjects `json:"projects,omitempty" xml:"projects,omitempty" type:"Repeated"`
}

func (s ListProjectsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListProjectsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListProjectsResponseBodyData) SetNextToken(v string) *ListProjectsResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *ListProjectsResponseBodyData) SetMarker(v string) *ListProjectsResponseBodyData {
	s.Marker = &v
	return s
}

func (s *ListProjectsResponseBodyData) SetMaxItem(v int32) *ListProjectsResponseBodyData {
	s.MaxItem = &v
	return s
}

func (s *ListProjectsResponseBodyData) SetProjects(v []*ListProjectsResponseBodyDataProjects) *ListProjectsResponseBodyData {
	s.Projects = v
	return s
}

type ListProjectsResponseBodyDataProjects struct {
	// The tag.
	Tags []*ListProjectsResponseBodyDataProjectsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The remarks.
	Comment *string `json:"comment,omitempty" xml:"comment,omitempty"`
	// The storage usage.
	CostStorage *string `json:"costStorage,omitempty" xml:"costStorage,omitempty"`
	// The default computing quota.
	DefaultQuota *string `json:"defaultQuota,omitempty" xml:"defaultQuota,omitempty"`
	// The IP address whitelist.
	IpWhiteList *ListProjectsResponseBodyDataProjectsIpWhiteList `json:"ipWhiteList,omitempty" xml:"ipWhiteList,omitempty" type:"Struct"`
	// The name of the MaxCompute project.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The owner of the project.
	Owner *string `json:"owner,omitempty" xml:"owner,omitempty"`
	// The properties of the project.
	Properties *ListProjectsResponseBodyDataProjectsProperties `json:"properties,omitempty" xml:"properties,omitempty" type:"Struct"`
	// The identifier of an object in a MaxCompute quota. This identifier is the same as the identifier in the sales bill of Alibaba Cloud. This parameter is used for tags.
	SaleTag *ListProjectsResponseBodyDataProjectsSaleTag `json:"saleTag,omitempty" xml:"saleTag,omitempty" type:"Struct"`
	// The permission properties.
	SecurityProperties *ListProjectsResponseBodyDataProjectsSecurityProperties `json:"securityProperties,omitempty" xml:"securityProperties,omitempty" type:"Struct"`
	// The status of the project. Valid values: -AVAILABLE: The project is available. -READONLY: The project is read only. -FROZEN: The project is frozen. -DELETING: The project is being deleted.
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// Indicates whether the current project support the three-tier model
	ThreeTierModel *bool `json:"threeTierModel,omitempty" xml:"threeTierModel,omitempty"`
	// The project type. Valid values: -managed: The project is an internal project. -external: The project is an external project.
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListProjectsResponseBodyDataProjects) String() string {
	return tea.Prettify(s)
}

func (s ListProjectsResponseBodyDataProjects) GoString() string {
	return s.String()
}

func (s *ListProjectsResponseBodyDataProjects) SetTags(v []*ListProjectsResponseBodyDataProjectsTags) *ListProjectsResponseBodyDataProjects {
	s.Tags = v
	return s
}

func (s *ListProjectsResponseBodyDataProjects) SetComment(v string) *ListProjectsResponseBodyDataProjects {
	s.Comment = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjects) SetCostStorage(v string) *ListProjectsResponseBodyDataProjects {
	s.CostStorage = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjects) SetDefaultQuota(v string) *ListProjectsResponseBodyDataProjects {
	s.DefaultQuota = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjects) SetIpWhiteList(v *ListProjectsResponseBodyDataProjectsIpWhiteList) *ListProjectsResponseBodyDataProjects {
	s.IpWhiteList = v
	return s
}

func (s *ListProjectsResponseBodyDataProjects) SetName(v string) *ListProjectsResponseBodyDataProjects {
	s.Name = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjects) SetOwner(v string) *ListProjectsResponseBodyDataProjects {
	s.Owner = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjects) SetProperties(v *ListProjectsResponseBodyDataProjectsProperties) *ListProjectsResponseBodyDataProjects {
	s.Properties = v
	return s
}

func (s *ListProjectsResponseBodyDataProjects) SetSaleTag(v *ListProjectsResponseBodyDataProjectsSaleTag) *ListProjectsResponseBodyDataProjects {
	s.SaleTag = v
	return s
}

func (s *ListProjectsResponseBodyDataProjects) SetSecurityProperties(v *ListProjectsResponseBodyDataProjectsSecurityProperties) *ListProjectsResponseBodyDataProjects {
	s.SecurityProperties = v
	return s
}

func (s *ListProjectsResponseBodyDataProjects) SetStatus(v string) *ListProjectsResponseBodyDataProjects {
	s.Status = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjects) SetThreeTierModel(v bool) *ListProjectsResponseBodyDataProjects {
	s.ThreeTierModel = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjects) SetType(v string) *ListProjectsResponseBodyDataProjects {
	s.Type = &v
	return s
}

type ListProjectsResponseBodyDataProjectsTags struct {
	// The key of the tag.
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The value of the tag.
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListProjectsResponseBodyDataProjectsTags) String() string {
	return tea.Prettify(s)
}

func (s ListProjectsResponseBodyDataProjectsTags) GoString() string {
	return s.String()
}

func (s *ListProjectsResponseBodyDataProjectsTags) SetTagKey(v string) *ListProjectsResponseBodyDataProjectsTags {
	s.TagKey = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjectsTags) SetTagValue(v string) *ListProjectsResponseBodyDataProjectsTags {
	s.TagValue = &v
	return s
}

type ListProjectsResponseBodyDataProjectsIpWhiteList struct {
	// The list of IP addresses.
	IpList *string `json:"ipList,omitempty" xml:"ipList,omitempty"`
	// The list of virtual private cloud (VPC) IP addresses.
	VpcIpList *string `json:"vpcIpList,omitempty" xml:"vpcIpList,omitempty"`
}

func (s ListProjectsResponseBodyDataProjectsIpWhiteList) String() string {
	return tea.Prettify(s)
}

func (s ListProjectsResponseBodyDataProjectsIpWhiteList) GoString() string {
	return s.String()
}

func (s *ListProjectsResponseBodyDataProjectsIpWhiteList) SetIpList(v string) *ListProjectsResponseBodyDataProjectsIpWhiteList {
	s.IpList = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjectsIpWhiteList) SetVpcIpList(v string) *ListProjectsResponseBodyDataProjectsIpWhiteList {
	s.VpcIpList = &v
	return s
}

type ListProjectsResponseBodyDataProjectsProperties struct {
	// Indicates whether a full table scan on the project is enabled.
	AllowFullScan *bool `json:"allowFullScan,omitempty" xml:"allowFullScan,omitempty"`
	// Indicates whether the DECIMAL data type in MaxCompute V2.0 is enabled.
	EnableDecimal2 *bool `json:"enableDecimal2,omitempty" xml:"enableDecimal2,omitempty"`
	// Indicates whether tunnel quota routing is enabled.
	EnableTunnelQuotaRoute *bool `json:"enableTunnelQuotaRoute,omitempty" xml:"enableTunnelQuotaRoute,omitempty"`
	// The encryption information.
	Encryption *ListProjectsResponseBodyDataProjectsPropertiesEncryption `json:"encryption,omitempty" xml:"encryption,omitempty" type:"Struct"`
	// The number of days for which backup data can be retained.
	RetentionDays *int64 `json:"retentionDays,omitempty" xml:"retentionDays,omitempty"`
	// The upper limit for the resources that are consumed by an SQL statement.
	SqlMeteringMax *string `json:"sqlMeteringMax,omitempty" xml:"sqlMeteringMax,omitempty"`
	// The lifecycle of tables in the project.
	TableLifecycle *ListProjectsResponseBodyDataProjectsPropertiesTableLifecycle `json:"tableLifecycle,omitempty" xml:"tableLifecycle,omitempty" type:"Struct"`
	// The time zone of the project.
	Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
	// The name of the tunnel quota.
	TunnelQuota *string `json:"tunnelQuota,omitempty" xml:"tunnelQuota,omitempty"`
	// The data type edition. Valid values: -1: MaxCompute V1.0 data type edition. -2: MaxCompute V2.0 data type edition. -hive: Hive-compatible data type edition.
	TypeSystem *string `json:"typeSystem,omitempty" xml:"typeSystem,omitempty"`
}

func (s ListProjectsResponseBodyDataProjectsProperties) String() string {
	return tea.Prettify(s)
}

func (s ListProjectsResponseBodyDataProjectsProperties) GoString() string {
	return s.String()
}

func (s *ListProjectsResponseBodyDataProjectsProperties) SetAllowFullScan(v bool) *ListProjectsResponseBodyDataProjectsProperties {
	s.AllowFullScan = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjectsProperties) SetEnableDecimal2(v bool) *ListProjectsResponseBodyDataProjectsProperties {
	s.EnableDecimal2 = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjectsProperties) SetEnableTunnelQuotaRoute(v bool) *ListProjectsResponseBodyDataProjectsProperties {
	s.EnableTunnelQuotaRoute = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjectsProperties) SetEncryption(v *ListProjectsResponseBodyDataProjectsPropertiesEncryption) *ListProjectsResponseBodyDataProjectsProperties {
	s.Encryption = v
	return s
}

func (s *ListProjectsResponseBodyDataProjectsProperties) SetRetentionDays(v int64) *ListProjectsResponseBodyDataProjectsProperties {
	s.RetentionDays = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjectsProperties) SetSqlMeteringMax(v string) *ListProjectsResponseBodyDataProjectsProperties {
	s.SqlMeteringMax = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjectsProperties) SetTableLifecycle(v *ListProjectsResponseBodyDataProjectsPropertiesTableLifecycle) *ListProjectsResponseBodyDataProjectsProperties {
	s.TableLifecycle = v
	return s
}

func (s *ListProjectsResponseBodyDataProjectsProperties) SetTimezone(v string) *ListProjectsResponseBodyDataProjectsProperties {
	s.Timezone = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjectsProperties) SetTunnelQuota(v string) *ListProjectsResponseBodyDataProjectsProperties {
	s.TunnelQuota = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjectsProperties) SetTypeSystem(v string) *ListProjectsResponseBodyDataProjectsProperties {
	s.TypeSystem = &v
	return s
}

type ListProjectsResponseBodyDataProjectsPropertiesEncryption struct {
	// The name of the encryption algorithm.
	Algorithm *string `json:"algorithm,omitempty" xml:"algorithm,omitempty"`
	// Indicates whether data encryption is enabled. Valid values: true and false.
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// The key of the encryption algorithm.
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
}

func (s ListProjectsResponseBodyDataProjectsPropertiesEncryption) String() string {
	return tea.Prettify(s)
}

func (s ListProjectsResponseBodyDataProjectsPropertiesEncryption) GoString() string {
	return s.String()
}

func (s *ListProjectsResponseBodyDataProjectsPropertiesEncryption) SetAlgorithm(v string) *ListProjectsResponseBodyDataProjectsPropertiesEncryption {
	s.Algorithm = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjectsPropertiesEncryption) SetEnable(v bool) *ListProjectsResponseBodyDataProjectsPropertiesEncryption {
	s.Enable = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjectsPropertiesEncryption) SetKey(v string) *ListProjectsResponseBodyDataProjectsPropertiesEncryption {
	s.Key = &v
	return s
}

type ListProjectsResponseBodyDataProjectsPropertiesTableLifecycle struct {
	// The type of the lifecycle. Valid values: -mandatory: The lifecycle clause is required. You must configure a lifecycle for a table. -optional: The lifecycle clause is optional in a table creation statement. If you do not configure a lifecycle for a table, the table does not expire. -inherit: If you do not configure a lifecycle for a table when you create the table, the value of odps.table.lifecycle.value is used by default.
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The retention period of a table. Unit: days.
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListProjectsResponseBodyDataProjectsPropertiesTableLifecycle) String() string {
	return tea.Prettify(s)
}

func (s ListProjectsResponseBodyDataProjectsPropertiesTableLifecycle) GoString() string {
	return s.String()
}

func (s *ListProjectsResponseBodyDataProjectsPropertiesTableLifecycle) SetType(v string) *ListProjectsResponseBodyDataProjectsPropertiesTableLifecycle {
	s.Type = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjectsPropertiesTableLifecycle) SetValue(v string) *ListProjectsResponseBodyDataProjectsPropertiesTableLifecycle {
	s.Value = &v
	return s
}

type ListProjectsResponseBodyDataProjectsSaleTag struct {
	// The identifier of an object in a MaxCompute quota. This identifier is the same as the identifier in the sales bill of Alibaba Cloud. This parameter is used for tags.
	ResourceId *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	// The type of the object. Valid values: quota and project.
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
}

func (s ListProjectsResponseBodyDataProjectsSaleTag) String() string {
	return tea.Prettify(s)
}

func (s ListProjectsResponseBodyDataProjectsSaleTag) GoString() string {
	return s.String()
}

func (s *ListProjectsResponseBodyDataProjectsSaleTag) SetResourceId(v string) *ListProjectsResponseBodyDataProjectsSaleTag {
	s.ResourceId = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjectsSaleTag) SetResourceType(v string) *ListProjectsResponseBodyDataProjectsSaleTag {
	s.ResourceType = &v
	return s
}

type ListProjectsResponseBodyDataProjectsSecurityProperties struct {
	// Indicates whether Download control is enabled.
	EnableDownloadPrivilege *bool `json:"enableDownloadPrivilege,omitempty" xml:"enableDownloadPrivilege,omitempty"`
	// Indicates whether label-based access control is enabled.
	LabelSecurity *bool `json:"labelSecurity,omitempty" xml:"labelSecurity,omitempty"`
	// Indicates whether the object creator is allowed to perform operations on objects.
	ObjectCreatorHasAccessPermission *bool `json:"objectCreatorHasAccessPermission,omitempty" xml:"objectCreatorHasAccessPermission,omitempty"`
	// Indicates whether the object creator is allowed to authorize other users to perform operations on objects.
	ObjectCreatorHasGrantPermission *bool `json:"objectCreatorHasGrantPermission,omitempty" xml:"objectCreatorHasGrantPermission,omitempty"`
	// Indicates whether project data protection is enabled.
	ProjectProtection *ListProjectsResponseBodyDataProjectsSecurityPropertiesProjectProtection `json:"projectProtection,omitempty" xml:"projectProtection,omitempty" type:"Struct"`
	// Indicates whether ACL-based access control is enabled.
	UsingAcl *bool `json:"usingAcl,omitempty" xml:"usingAcl,omitempty"`
	// Indicates whether policy-based access control is enabled.
	UsingPolicy *bool `json:"usingPolicy,omitempty" xml:"usingPolicy,omitempty"`
}

func (s ListProjectsResponseBodyDataProjectsSecurityProperties) String() string {
	return tea.Prettify(s)
}

func (s ListProjectsResponseBodyDataProjectsSecurityProperties) GoString() string {
	return s.String()
}

func (s *ListProjectsResponseBodyDataProjectsSecurityProperties) SetEnableDownloadPrivilege(v bool) *ListProjectsResponseBodyDataProjectsSecurityProperties {
	s.EnableDownloadPrivilege = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjectsSecurityProperties) SetLabelSecurity(v bool) *ListProjectsResponseBodyDataProjectsSecurityProperties {
	s.LabelSecurity = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjectsSecurityProperties) SetObjectCreatorHasAccessPermission(v bool) *ListProjectsResponseBodyDataProjectsSecurityProperties {
	s.ObjectCreatorHasAccessPermission = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjectsSecurityProperties) SetObjectCreatorHasGrantPermission(v bool) *ListProjectsResponseBodyDataProjectsSecurityProperties {
	s.ObjectCreatorHasGrantPermission = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjectsSecurityProperties) SetProjectProtection(v *ListProjectsResponseBodyDataProjectsSecurityPropertiesProjectProtection) *ListProjectsResponseBodyDataProjectsSecurityProperties {
	s.ProjectProtection = v
	return s
}

func (s *ListProjectsResponseBodyDataProjectsSecurityProperties) SetUsingAcl(v bool) *ListProjectsResponseBodyDataProjectsSecurityProperties {
	s.UsingAcl = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjectsSecurityProperties) SetUsingPolicy(v bool) *ListProjectsResponseBodyDataProjectsSecurityProperties {
	s.UsingPolicy = &v
	return s
}

type ListProjectsResponseBodyDataProjectsSecurityPropertiesProjectProtection struct {
	// The exception policy. If cross-project data access operations are required, the project owner must configure an exception policy in advance to allow the specified user to transfer data of a specified object from the current project to a specified project. After the exception policy is configured, data of the object can be transferred to the specified project even if the project data protection feature is enabled.
	ExceptionPolicy *string `json:"exceptionPolicy,omitempty" xml:"exceptionPolicy,omitempty"`
	// Indicates whether project data protection is enabled.
	Protected *bool `json:"protected,omitempty" xml:"protected,omitempty"`
}

func (s ListProjectsResponseBodyDataProjectsSecurityPropertiesProjectProtection) String() string {
	return tea.Prettify(s)
}

func (s ListProjectsResponseBodyDataProjectsSecurityPropertiesProjectProtection) GoString() string {
	return s.String()
}

func (s *ListProjectsResponseBodyDataProjectsSecurityPropertiesProjectProtection) SetExceptionPolicy(v string) *ListProjectsResponseBodyDataProjectsSecurityPropertiesProjectProtection {
	s.ExceptionPolicy = &v
	return s
}

func (s *ListProjectsResponseBodyDataProjectsSecurityPropertiesProjectProtection) SetProtected(v bool) *ListProjectsResponseBodyDataProjectsSecurityPropertiesProjectProtection {
	s.Protected = &v
	return s
}

type ListProjectsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListProjectsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListProjectsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListProjectsResponse) GoString() string {
	return s.String()
}

func (s *ListProjectsResponse) SetHeaders(v map[string]*string) *ListProjectsResponse {
	s.Headers = v
	return s
}

func (s *ListProjectsResponse) SetStatusCode(v int32) *ListProjectsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListProjectsResponse) SetBody(v *ListProjectsResponseBody) *ListProjectsResponse {
	s.Body = v
	return s
}

type ListQuotasRequest struct {
	// The billing method of the quota.
	BillingType *string `json:"billingType,omitempty" xml:"billingType,omitempty"`
	// Specifies the marker after which the returned list begins.
	Marker *string `json:"marker,omitempty" xml:"marker,omitempty"`
	// The maximum number of entries to return on each page.
	MaxItem *int64 `json:"maxItem,omitempty" xml:"maxItem,omitempty"`
	// The service ID.
	ProductId *string `json:"productId,omitempty" xml:"productId,omitempty"`
	// The ID of the region.
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// The cost tag. You can filter out quota objects based on the cost tag. The cost tag is created when you tag a service.
	SaleTags *string `json:"saleTags,omitempty" xml:"saleTags,omitempty"`
	// The ID of the tenant.
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ListQuotasRequest) String() string {
	return tea.Prettify(s)
}

func (s ListQuotasRequest) GoString() string {
	return s.String()
}

func (s *ListQuotasRequest) SetBillingType(v string) *ListQuotasRequest {
	s.BillingType = &v
	return s
}

func (s *ListQuotasRequest) SetMarker(v string) *ListQuotasRequest {
	s.Marker = &v
	return s
}

func (s *ListQuotasRequest) SetMaxItem(v int64) *ListQuotasRequest {
	s.MaxItem = &v
	return s
}

func (s *ListQuotasRequest) SetProductId(v string) *ListQuotasRequest {
	s.ProductId = &v
	return s
}

func (s *ListQuotasRequest) SetRegion(v string) *ListQuotasRequest {
	s.Region = &v
	return s
}

func (s *ListQuotasRequest) SetSaleTags(v string) *ListQuotasRequest {
	s.SaleTags = &v
	return s
}

func (s *ListQuotasRequest) SetTenantId(v string) *ListQuotasRequest {
	s.TenantId = &v
	return s
}

type ListQuotasResponseBody struct {
	// A pagination token. Only continuous page turning is supported. If NextToken is not empty, the next page exists. The value of NextToken can be used in the next request to retrieve a new page of results.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The returned data.
	Data *ListQuotasResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Indicates the marker after which the returned list begins.
	Marker *string `json:"marker,omitempty" xml:"marker,omitempty"`
	// The maximum number of entries returned per page.
	MaxItem *int64 `json:"maxItem,omitempty" xml:"maxItem,omitempty"`
	// The list of quotas.
	QuotaInfoList []*ListQuotasResponseBodyQuotaInfoList `json:"quotaInfoList,omitempty" xml:"quotaInfoList,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListQuotasResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListQuotasResponseBody) GoString() string {
	return s.String()
}

func (s *ListQuotasResponseBody) SetNextToken(v string) *ListQuotasResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListQuotasResponseBody) SetData(v *ListQuotasResponseBodyData) *ListQuotasResponseBody {
	s.Data = v
	return s
}

func (s *ListQuotasResponseBody) SetMarker(v string) *ListQuotasResponseBody {
	s.Marker = &v
	return s
}

func (s *ListQuotasResponseBody) SetMaxItem(v int64) *ListQuotasResponseBody {
	s.MaxItem = &v
	return s
}

func (s *ListQuotasResponseBody) SetQuotaInfoList(v []*ListQuotasResponseBodyQuotaInfoList) *ListQuotasResponseBody {
	s.QuotaInfoList = v
	return s
}

func (s *ListQuotasResponseBody) SetRequestId(v string) *ListQuotasResponseBody {
	s.RequestId = &v
	return s
}

type ListQuotasResponseBodyData struct {
	// A pagination token. Only continuous page turning is supported. If NextToken is not empty, the next page exists. The value of NextToken can be used in the next request to retrieve a new page of results.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Indicates the marker after which the returned list begins.
	Marker *string `json:"marker,omitempty" xml:"marker,omitempty"`
	// The maximum number of entries returned per page.
	MaxItem *int64 `json:"maxItem,omitempty" xml:"maxItem,omitempty"`
	// The list of quotas.
	QuotaInfoList []*ListQuotasResponseBodyDataQuotaInfoList `json:"quotaInfoList,omitempty" xml:"quotaInfoList,omitempty" type:"Repeated"`
}

func (s ListQuotasResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListQuotasResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListQuotasResponseBodyData) SetNextToken(v string) *ListQuotasResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *ListQuotasResponseBodyData) SetMarker(v string) *ListQuotasResponseBodyData {
	s.Marker = &v
	return s
}

func (s *ListQuotasResponseBodyData) SetMaxItem(v int64) *ListQuotasResponseBodyData {
	s.MaxItem = &v
	return s
}

func (s *ListQuotasResponseBodyData) SetQuotaInfoList(v []*ListQuotasResponseBodyDataQuotaInfoList) *ListQuotasResponseBodyData {
	s.QuotaInfoList = v
	return s
}

type ListQuotasResponseBodyDataQuotaInfoList struct {
	// The tags.
	Tags []*ListQuotasResponseBodyDataQuotaInfoListTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The information of the order.
	BillingPolicy *ListQuotasResponseBodyDataQuotaInfoListBillingPolicy `json:"billingPolicy,omitempty" xml:"billingPolicy,omitempty" type:"Struct"`
	// The cluster ID.
	Cluster *string `json:"cluster,omitempty" xml:"cluster,omitempty"`
	// The time when the resource was created.
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The ID of the Alibaba Cloud account that is used to create the resource.
	CreatorId *string `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	// The quota ID.
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The name of the quota.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The alias of the quota.
	NickName *string `json:"nickName,omitempty" xml:"nickName,omitempty"`
	// The description of the quota.
	Parameter map[string]interface{} `json:"parameter,omitempty" xml:"parameter,omitempty"`
	// The ID of the parent resource.
	ParentId *string `json:"parentId,omitempty" xml:"parentId,omitempty"`
	// The region ID.
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The identifier of an object in a MaxCompute quota. This identifier is the same as the identifier in the sales bill of Alibaba Cloud. This parameter is used for tags.
	SaleTag *ListQuotasResponseBodyDataQuotaInfoListSaleTag `json:"saleTag,omitempty" xml:"saleTag,omitempty" type:"Struct"`
	// The information of the scheduling plan.
	ScheduleInfo *ListQuotasResponseBodyDataQuotaInfoListScheduleInfo `json:"scheduleInfo,omitempty" xml:"scheduleInfo,omitempty" type:"Struct"`
	// The status of the endpoint group.
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The information of the level-2 quota.
	SubQuotaInfoList []*ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList `json:"subQuotaInfoList,omitempty" xml:"subQuotaInfoList,omitempty" type:"Repeated"`
	// The tag of the resource for the quota.
	Tag *string `json:"tag,omitempty" xml:"tag,omitempty"`
	// The tenant ID.
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The type of the resource system. This parameter corresponds to the resourceSystemType parameter of the cluster.
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The version number.
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s ListQuotasResponseBodyDataQuotaInfoList) String() string {
	return tea.Prettify(s)
}

func (s ListQuotasResponseBodyDataQuotaInfoList) GoString() string {
	return s.String()
}

func (s *ListQuotasResponseBodyDataQuotaInfoList) SetTags(v []*ListQuotasResponseBodyDataQuotaInfoListTags) *ListQuotasResponseBodyDataQuotaInfoList {
	s.Tags = v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoList) SetBillingPolicy(v *ListQuotasResponseBodyDataQuotaInfoListBillingPolicy) *ListQuotasResponseBodyDataQuotaInfoList {
	s.BillingPolicy = v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoList) SetCluster(v string) *ListQuotasResponseBodyDataQuotaInfoList {
	s.Cluster = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoList) SetCreateTime(v int64) *ListQuotasResponseBodyDataQuotaInfoList {
	s.CreateTime = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoList) SetCreatorId(v string) *ListQuotasResponseBodyDataQuotaInfoList {
	s.CreatorId = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoList) SetId(v string) *ListQuotasResponseBodyDataQuotaInfoList {
	s.Id = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoList) SetName(v string) *ListQuotasResponseBodyDataQuotaInfoList {
	s.Name = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoList) SetNickName(v string) *ListQuotasResponseBodyDataQuotaInfoList {
	s.NickName = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoList) SetParameter(v map[string]interface{}) *ListQuotasResponseBodyDataQuotaInfoList {
	s.Parameter = v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoList) SetParentId(v string) *ListQuotasResponseBodyDataQuotaInfoList {
	s.ParentId = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoList) SetRegionId(v string) *ListQuotasResponseBodyDataQuotaInfoList {
	s.RegionId = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoList) SetSaleTag(v *ListQuotasResponseBodyDataQuotaInfoListSaleTag) *ListQuotasResponseBodyDataQuotaInfoList {
	s.SaleTag = v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoList) SetScheduleInfo(v *ListQuotasResponseBodyDataQuotaInfoListScheduleInfo) *ListQuotasResponseBodyDataQuotaInfoList {
	s.ScheduleInfo = v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoList) SetStatus(v string) *ListQuotasResponseBodyDataQuotaInfoList {
	s.Status = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoList) SetSubQuotaInfoList(v []*ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList) *ListQuotasResponseBodyDataQuotaInfoList {
	s.SubQuotaInfoList = v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoList) SetTag(v string) *ListQuotasResponseBodyDataQuotaInfoList {
	s.Tag = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoList) SetTenantId(v string) *ListQuotasResponseBodyDataQuotaInfoList {
	s.TenantId = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoList) SetType(v string) *ListQuotasResponseBodyDataQuotaInfoList {
	s.Type = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoList) SetVersion(v string) *ListQuotasResponseBodyDataQuotaInfoList {
	s.Version = &v
	return s
}

type ListQuotasResponseBodyDataQuotaInfoListTags struct {
	// The key of the tag.
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The value of the tag.
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListQuotasResponseBodyDataQuotaInfoListTags) String() string {
	return tea.Prettify(s)
}

func (s ListQuotasResponseBodyDataQuotaInfoListTags) GoString() string {
	return s.String()
}

func (s *ListQuotasResponseBodyDataQuotaInfoListTags) SetTagKey(v string) *ListQuotasResponseBodyDataQuotaInfoListTags {
	s.TagKey = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListTags) SetTagValue(v string) *ListQuotasResponseBodyDataQuotaInfoListTags {
	s.TagValue = &v
	return s
}

type ListQuotasResponseBodyDataQuotaInfoListBillingPolicy struct {
	// The billing method of the quota. Valid values:
	//
	// *   subscription: a subscription quota.
	// *   payasyougo: a pay-as-you-go quota.
	BillingMethod *string `json:"billingMethod,omitempty" xml:"billingMethod,omitempty"`
	// The specifications of the order.
	OdpsSpecCode *string `json:"odpsSpecCode,omitempty" xml:"odpsSpecCode,omitempty"`
	// The order ID.
	OrderId *string `json:"orderId,omitempty" xml:"orderId,omitempty"`
}

func (s ListQuotasResponseBodyDataQuotaInfoListBillingPolicy) String() string {
	return tea.Prettify(s)
}

func (s ListQuotasResponseBodyDataQuotaInfoListBillingPolicy) GoString() string {
	return s.String()
}

func (s *ListQuotasResponseBodyDataQuotaInfoListBillingPolicy) SetBillingMethod(v string) *ListQuotasResponseBodyDataQuotaInfoListBillingPolicy {
	s.BillingMethod = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListBillingPolicy) SetOdpsSpecCode(v string) *ListQuotasResponseBodyDataQuotaInfoListBillingPolicy {
	s.OdpsSpecCode = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListBillingPolicy) SetOrderId(v string) *ListQuotasResponseBodyDataQuotaInfoListBillingPolicy {
	s.OrderId = &v
	return s
}

type ListQuotasResponseBodyDataQuotaInfoListSaleTag struct {
	// The identifier of an object in a MaxCompute quota. This identifier exists in the sales bill of Alibaba Cloud. You can use this identifier to associate the cost of a quota object with a tag.
	ResourceIds []*string `json:"resourceIds,omitempty" xml:"resourceIds,omitempty" type:"Repeated"`
	// The type of the object. Valid values: quota and project.
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
}

func (s ListQuotasResponseBodyDataQuotaInfoListSaleTag) String() string {
	return tea.Prettify(s)
}

func (s ListQuotasResponseBodyDataQuotaInfoListSaleTag) GoString() string {
	return s.String()
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSaleTag) SetResourceIds(v []*string) *ListQuotasResponseBodyDataQuotaInfoListSaleTag {
	s.ResourceIds = v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSaleTag) SetResourceType(v string) *ListQuotasResponseBodyDataQuotaInfoListSaleTag {
	s.ResourceType = &v
	return s
}

type ListQuotasResponseBodyDataQuotaInfoListScheduleInfo struct {
	// The quota plan that takes effect based on the scheduling plan.
	CurrPlan *string `json:"currPlan,omitempty" xml:"currPlan,omitempty"`
	// The time when the current quota plan is scheduled.
	CurrTime *string `json:"currTime,omitempty" xml:"currTime,omitempty"`
	// The next quota plan that will take effect based on the scheduling plan.
	NextPlan *string `json:"nextPlan,omitempty" xml:"nextPlan,omitempty"`
	// The time when the next quota plan is scheduled.
	NextTime *string `json:"nextTime,omitempty" xml:"nextTime,omitempty"`
	// The quota plan that immediately takes effect. If the quota plan that immediately takes effect is different from the current quota plan, this parameter is not empty.
	OncePlan *string `json:"oncePlan,omitempty" xml:"oncePlan,omitempty"`
	// The time when the quota plan immediately takes effect.
	OnceTime *string `json:"onceTime,omitempty" xml:"onceTime,omitempty"`
	// The name of the operator.
	OperatorName *string `json:"operatorName,omitempty" xml:"operatorName,omitempty"`
	// The time zone of the project.
	Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
}

func (s ListQuotasResponseBodyDataQuotaInfoListScheduleInfo) String() string {
	return tea.Prettify(s)
}

func (s ListQuotasResponseBodyDataQuotaInfoListScheduleInfo) GoString() string {
	return s.String()
}

func (s *ListQuotasResponseBodyDataQuotaInfoListScheduleInfo) SetCurrPlan(v string) *ListQuotasResponseBodyDataQuotaInfoListScheduleInfo {
	s.CurrPlan = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListScheduleInfo) SetCurrTime(v string) *ListQuotasResponseBodyDataQuotaInfoListScheduleInfo {
	s.CurrTime = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListScheduleInfo) SetNextPlan(v string) *ListQuotasResponseBodyDataQuotaInfoListScheduleInfo {
	s.NextPlan = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListScheduleInfo) SetNextTime(v string) *ListQuotasResponseBodyDataQuotaInfoListScheduleInfo {
	s.NextTime = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListScheduleInfo) SetOncePlan(v string) *ListQuotasResponseBodyDataQuotaInfoListScheduleInfo {
	s.OncePlan = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListScheduleInfo) SetOnceTime(v string) *ListQuotasResponseBodyDataQuotaInfoListScheduleInfo {
	s.OnceTime = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListScheduleInfo) SetOperatorName(v string) *ListQuotasResponseBodyDataQuotaInfoListScheduleInfo {
	s.OperatorName = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListScheduleInfo) SetTimezone(v string) *ListQuotasResponseBodyDataQuotaInfoListScheduleInfo {
	s.Timezone = &v
	return s
}

type ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList struct {
	// The information of the order.
	BillingPolicy *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListBillingPolicy `json:"billingPolicy,omitempty" xml:"billingPolicy,omitempty" type:"Struct"`
	// The cluster ID.
	Cluster *string `json:"cluster,omitempty" xml:"cluster,omitempty"`
	// The time when the resource was created.
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The ID of the Alibaba Cloud account that is used to create the resource.
	CreatorId *string `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	// The ID of the level-2 quota.
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The name of the level-2 quota.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The nickname of the level-2 quota.
	NickName *string `json:"nickName,omitempty" xml:"nickName,omitempty"`
	// The description of the quota.
	Parameter map[string]interface{} `json:"parameter,omitempty" xml:"parameter,omitempty"`
	// The ID of the parent resource.
	ParentId *string `json:"parentId,omitempty" xml:"parentId,omitempty"`
	// The region ID.
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The identifier of an object in a MaxCompute quota. This identifier is the same as the identifier in the sales bill of Alibaba Cloud. This parameter is used for tags.
	SaleTag *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListSaleTag `json:"saleTag,omitempty" xml:"saleTag,omitempty" type:"Struct"`
	// The information of the scheduling plan.
	ScheduleInfo *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListScheduleInfo `json:"scheduleInfo,omitempty" xml:"scheduleInfo,omitempty" type:"Struct"`
	// The status of the endpoint group.
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The tag of the resource for the quota.
	Tag *string `json:"tag,omitempty" xml:"tag,omitempty"`
	// The tenant ID.
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The type of the resource system. This parameter corresponds to the resourceSystemType parameter of the cluster.
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The version of the algorithm image.
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList) String() string {
	return tea.Prettify(s)
}

func (s ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList) GoString() string {
	return s.String()
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList) SetBillingPolicy(v *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListBillingPolicy) *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList {
	s.BillingPolicy = v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList) SetCluster(v string) *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList {
	s.Cluster = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList) SetCreateTime(v int64) *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList {
	s.CreateTime = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList) SetCreatorId(v string) *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList {
	s.CreatorId = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList) SetId(v string) *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList {
	s.Id = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList) SetName(v string) *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList {
	s.Name = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList) SetNickName(v string) *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList {
	s.NickName = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList) SetParameter(v map[string]interface{}) *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList {
	s.Parameter = v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList) SetParentId(v string) *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList {
	s.ParentId = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList) SetRegionId(v string) *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList {
	s.RegionId = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList) SetSaleTag(v *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListSaleTag) *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList {
	s.SaleTag = v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList) SetScheduleInfo(v *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListScheduleInfo) *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList {
	s.ScheduleInfo = v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList) SetStatus(v string) *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList {
	s.Status = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList) SetTag(v string) *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList {
	s.Tag = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList) SetTenantId(v string) *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList {
	s.TenantId = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList) SetType(v string) *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList {
	s.Type = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList) SetVersion(v string) *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoList {
	s.Version = &v
	return s
}

type ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListBillingPolicy struct {
	// The billing method of the quota. Valid values:
	//
	// *   subscription: a subscription quota.
	// *   payasyougo: a pay-as-you-go quota.
	BillingMethod *string `json:"billingMethod,omitempty" xml:"billingMethod,omitempty"`
	// The specifications of the order.
	OdpsSpecCode *string `json:"odpsSpecCode,omitempty" xml:"odpsSpecCode,omitempty"`
	// The order ID.
	OrderId *string `json:"orderId,omitempty" xml:"orderId,omitempty"`
}

func (s ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListBillingPolicy) String() string {
	return tea.Prettify(s)
}

func (s ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListBillingPolicy) GoString() string {
	return s.String()
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListBillingPolicy) SetBillingMethod(v string) *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListBillingPolicy {
	s.BillingMethod = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListBillingPolicy) SetOdpsSpecCode(v string) *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListBillingPolicy {
	s.OdpsSpecCode = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListBillingPolicy) SetOrderId(v string) *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListBillingPolicy {
	s.OrderId = &v
	return s
}

type ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListSaleTag struct {
	// The identifier of an object in a MaxCompute quota. This identifier exists in the sales bill of Alibaba Cloud. You can use this identifier to associate the cost of a quota object with a tag.
	ResourceIds []*string `json:"resourceIds,omitempty" xml:"resourceIds,omitempty" type:"Repeated"`
	// The type of the object. Valid values: quota and project.
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
}

func (s ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListSaleTag) String() string {
	return tea.Prettify(s)
}

func (s ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListSaleTag) GoString() string {
	return s.String()
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListSaleTag) SetResourceIds(v []*string) *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListSaleTag {
	s.ResourceIds = v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListSaleTag) SetResourceType(v string) *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListSaleTag {
	s.ResourceType = &v
	return s
}

type ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListScheduleInfo struct {
	// The quota plan that takes effect based on the scheduling plan.
	CurrPlan *string `json:"currPlan,omitempty" xml:"currPlan,omitempty"`
	// The time when the current quota plan is scheduled.
	CurrTime *string `json:"currTime,omitempty" xml:"currTime,omitempty"`
	// The next quota plan that will take effect based on the scheduling plan.
	NextPlan *string `json:"nextPlan,omitempty" xml:"nextPlan,omitempty"`
	// The time when the next quota plan is scheduled.
	NextTime *string `json:"nextTime,omitempty" xml:"nextTime,omitempty"`
	// The quota plan that immediately takes effect. If the quota plan that immediately takes effect is different from the current quota plan, this parameter is not empty.
	OncePlan *string `json:"oncePlan,omitempty" xml:"oncePlan,omitempty"`
	// The time when the quota plan immediately takes effect.
	OnceTime *string `json:"onceTime,omitempty" xml:"onceTime,omitempty"`
	// The name of the operator.
	OperatorName *string `json:"operatorName,omitempty" xml:"operatorName,omitempty"`
	// The time zone of the project.
	Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
}

func (s ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListScheduleInfo) String() string {
	return tea.Prettify(s)
}

func (s ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListScheduleInfo) GoString() string {
	return s.String()
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListScheduleInfo) SetCurrPlan(v string) *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListScheduleInfo {
	s.CurrPlan = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListScheduleInfo) SetCurrTime(v string) *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListScheduleInfo {
	s.CurrTime = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListScheduleInfo) SetNextPlan(v string) *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListScheduleInfo {
	s.NextPlan = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListScheduleInfo) SetNextTime(v string) *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListScheduleInfo {
	s.NextTime = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListScheduleInfo) SetOncePlan(v string) *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListScheduleInfo {
	s.OncePlan = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListScheduleInfo) SetOnceTime(v string) *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListScheduleInfo {
	s.OnceTime = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListScheduleInfo) SetOperatorName(v string) *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListScheduleInfo {
	s.OperatorName = &v
	return s
}

func (s *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListScheduleInfo) SetTimezone(v string) *ListQuotasResponseBodyDataQuotaInfoListSubQuotaInfoListScheduleInfo {
	s.Timezone = &v
	return s
}

type ListQuotasResponseBodyQuotaInfoList struct {
	// The tags.
	Tags []*ListQuotasResponseBodyQuotaInfoListTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The information of the order.
	BillingPolicy *ListQuotasResponseBodyQuotaInfoListBillingPolicy `json:"billingPolicy,omitempty" xml:"billingPolicy,omitempty" type:"Struct"`
	// The cluster ID.
	Cluster *string `json:"cluster,omitempty" xml:"cluster,omitempty"`
	// The time when the resource was created.
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The ID of the Alibaba Cloud account that is used to create the resource.
	CreatorId *string `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	// The quota ID.
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The name of the quota.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The alias of the quota.
	NickName *string `json:"nickName,omitempty" xml:"nickName,omitempty"`
	// The description of the quota.
	Parameter map[string]interface{} `json:"parameter,omitempty" xml:"parameter,omitempty"`
	// The ID of the parent resource.
	ParentId *string `json:"parentId,omitempty" xml:"parentId,omitempty"`
	// The region ID.
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The identifier of an object in a MaxCompute quota. This identifier is the same as the identifier in the sales bill of Alibaba Cloud. This parameter is used for tags.
	SaleTag *ListQuotasResponseBodyQuotaInfoListSaleTag `json:"saleTag,omitempty" xml:"saleTag,omitempty" type:"Struct"`
	// The information of the scheduling plan.
	ScheduleInfo *ListQuotasResponseBodyQuotaInfoListScheduleInfo `json:"scheduleInfo,omitempty" xml:"scheduleInfo,omitempty" type:"Struct"`
	// The status of the endpoint group.
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The information of the level-2 quota.
	SubQuotaInfoList []*ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList `json:"subQuotaInfoList,omitempty" xml:"subQuotaInfoList,omitempty" type:"Repeated"`
	// The tag of the resource for the quota.
	Tag *string `json:"tag,omitempty" xml:"tag,omitempty"`
	// The tenant ID.
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The type of the resource system. This parameter corresponds to the resourceSystemType parameter of the cluster.
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The version.
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s ListQuotasResponseBodyQuotaInfoList) String() string {
	return tea.Prettify(s)
}

func (s ListQuotasResponseBodyQuotaInfoList) GoString() string {
	return s.String()
}

func (s *ListQuotasResponseBodyQuotaInfoList) SetTags(v []*ListQuotasResponseBodyQuotaInfoListTags) *ListQuotasResponseBodyQuotaInfoList {
	s.Tags = v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoList) SetBillingPolicy(v *ListQuotasResponseBodyQuotaInfoListBillingPolicy) *ListQuotasResponseBodyQuotaInfoList {
	s.BillingPolicy = v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoList) SetCluster(v string) *ListQuotasResponseBodyQuotaInfoList {
	s.Cluster = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoList) SetCreateTime(v int64) *ListQuotasResponseBodyQuotaInfoList {
	s.CreateTime = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoList) SetCreatorId(v string) *ListQuotasResponseBodyQuotaInfoList {
	s.CreatorId = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoList) SetId(v string) *ListQuotasResponseBodyQuotaInfoList {
	s.Id = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoList) SetName(v string) *ListQuotasResponseBodyQuotaInfoList {
	s.Name = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoList) SetNickName(v string) *ListQuotasResponseBodyQuotaInfoList {
	s.NickName = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoList) SetParameter(v map[string]interface{}) *ListQuotasResponseBodyQuotaInfoList {
	s.Parameter = v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoList) SetParentId(v string) *ListQuotasResponseBodyQuotaInfoList {
	s.ParentId = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoList) SetRegionId(v string) *ListQuotasResponseBodyQuotaInfoList {
	s.RegionId = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoList) SetSaleTag(v *ListQuotasResponseBodyQuotaInfoListSaleTag) *ListQuotasResponseBodyQuotaInfoList {
	s.SaleTag = v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoList) SetScheduleInfo(v *ListQuotasResponseBodyQuotaInfoListScheduleInfo) *ListQuotasResponseBodyQuotaInfoList {
	s.ScheduleInfo = v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoList) SetStatus(v string) *ListQuotasResponseBodyQuotaInfoList {
	s.Status = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoList) SetSubQuotaInfoList(v []*ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList) *ListQuotasResponseBodyQuotaInfoList {
	s.SubQuotaInfoList = v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoList) SetTag(v string) *ListQuotasResponseBodyQuotaInfoList {
	s.Tag = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoList) SetTenantId(v string) *ListQuotasResponseBodyQuotaInfoList {
	s.TenantId = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoList) SetType(v string) *ListQuotasResponseBodyQuotaInfoList {
	s.Type = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoList) SetVersion(v string) *ListQuotasResponseBodyQuotaInfoList {
	s.Version = &v
	return s
}

type ListQuotasResponseBodyQuotaInfoListTags struct {
	// The key of the tag.
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The value of the tag.
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListQuotasResponseBodyQuotaInfoListTags) String() string {
	return tea.Prettify(s)
}

func (s ListQuotasResponseBodyQuotaInfoListTags) GoString() string {
	return s.String()
}

func (s *ListQuotasResponseBodyQuotaInfoListTags) SetTagKey(v string) *ListQuotasResponseBodyQuotaInfoListTags {
	s.TagKey = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListTags) SetTagValue(v string) *ListQuotasResponseBodyQuotaInfoListTags {
	s.TagValue = &v
	return s
}

type ListQuotasResponseBodyQuotaInfoListBillingPolicy struct {
	// The billing method of the quota. Valid values:
	//
	// *   subscription: a subscription quota.
	// *   payasyougo: a pay-as-you-go quota.
	BillingMethod *string `json:"billingMethod,omitempty" xml:"billingMethod,omitempty"`
	// The specifications of the order.
	OdpsSpecCode *string `json:"odpsSpecCode,omitempty" xml:"odpsSpecCode,omitempty"`
	// The order ID.
	OrderId *string `json:"orderId,omitempty" xml:"orderId,omitempty"`
}

func (s ListQuotasResponseBodyQuotaInfoListBillingPolicy) String() string {
	return tea.Prettify(s)
}

func (s ListQuotasResponseBodyQuotaInfoListBillingPolicy) GoString() string {
	return s.String()
}

func (s *ListQuotasResponseBodyQuotaInfoListBillingPolicy) SetBillingMethod(v string) *ListQuotasResponseBodyQuotaInfoListBillingPolicy {
	s.BillingMethod = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListBillingPolicy) SetOdpsSpecCode(v string) *ListQuotasResponseBodyQuotaInfoListBillingPolicy {
	s.OdpsSpecCode = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListBillingPolicy) SetOrderId(v string) *ListQuotasResponseBodyQuotaInfoListBillingPolicy {
	s.OrderId = &v
	return s
}

type ListQuotasResponseBodyQuotaInfoListSaleTag struct {
	// The identifier of an object in a MaxCompute quota. This identifier exists in the sales bill of Alibaba Cloud. You can use this identifier to associate the cost of a quota object with a tag.
	ResourceIds []*string `json:"resourceIds,omitempty" xml:"resourceIds,omitempty" type:"Repeated"`
	// The type of the object. Valid values: quota and project.
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
}

func (s ListQuotasResponseBodyQuotaInfoListSaleTag) String() string {
	return tea.Prettify(s)
}

func (s ListQuotasResponseBodyQuotaInfoListSaleTag) GoString() string {
	return s.String()
}

func (s *ListQuotasResponseBodyQuotaInfoListSaleTag) SetResourceIds(v []*string) *ListQuotasResponseBodyQuotaInfoListSaleTag {
	s.ResourceIds = v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListSaleTag) SetResourceType(v string) *ListQuotasResponseBodyQuotaInfoListSaleTag {
	s.ResourceType = &v
	return s
}

type ListQuotasResponseBodyQuotaInfoListScheduleInfo struct {
	// The quota plan that takes effect based on the scheduling plan.
	CurrPlan *string `json:"currPlan,omitempty" xml:"currPlan,omitempty"`
	// The time when the current quota plan is scheduled.
	CurrTime *string `json:"currTime,omitempty" xml:"currTime,omitempty"`
	// The next quota plan that will take effect based on the scheduling plan.
	NextPlan *string `json:"nextPlan,omitempty" xml:"nextPlan,omitempty"`
	// The time when the next quota plan is scheduled.
	NextTime *string `json:"nextTime,omitempty" xml:"nextTime,omitempty"`
	// The quota plan that immediately takes effect. If the quota plan that immediately takes effect is different from the current quota plan, this parameter is not empty.
	OncePlan *string `json:"oncePlan,omitempty" xml:"oncePlan,omitempty"`
	// The time when the quota plan immediately takes effect.
	OnceTime *string `json:"onceTime,omitempty" xml:"onceTime,omitempty"`
	// The name of the operator.
	OperatorName *string `json:"operatorName,omitempty" xml:"operatorName,omitempty"`
	// The time zone of the project.
	Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
}

func (s ListQuotasResponseBodyQuotaInfoListScheduleInfo) String() string {
	return tea.Prettify(s)
}

func (s ListQuotasResponseBodyQuotaInfoListScheduleInfo) GoString() string {
	return s.String()
}

func (s *ListQuotasResponseBodyQuotaInfoListScheduleInfo) SetCurrPlan(v string) *ListQuotasResponseBodyQuotaInfoListScheduleInfo {
	s.CurrPlan = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListScheduleInfo) SetCurrTime(v string) *ListQuotasResponseBodyQuotaInfoListScheduleInfo {
	s.CurrTime = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListScheduleInfo) SetNextPlan(v string) *ListQuotasResponseBodyQuotaInfoListScheduleInfo {
	s.NextPlan = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListScheduleInfo) SetNextTime(v string) *ListQuotasResponseBodyQuotaInfoListScheduleInfo {
	s.NextTime = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListScheduleInfo) SetOncePlan(v string) *ListQuotasResponseBodyQuotaInfoListScheduleInfo {
	s.OncePlan = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListScheduleInfo) SetOnceTime(v string) *ListQuotasResponseBodyQuotaInfoListScheduleInfo {
	s.OnceTime = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListScheduleInfo) SetOperatorName(v string) *ListQuotasResponseBodyQuotaInfoListScheduleInfo {
	s.OperatorName = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListScheduleInfo) SetTimezone(v string) *ListQuotasResponseBodyQuotaInfoListScheduleInfo {
	s.Timezone = &v
	return s
}

type ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList struct {
	// The information of the order.
	BillingPolicy *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListBillingPolicy `json:"billingPolicy,omitempty" xml:"billingPolicy,omitempty" type:"Struct"`
	// The cluster ID.
	Cluster *string `json:"cluster,omitempty" xml:"cluster,omitempty"`
	// The time when the resource was created.
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The ID of the Alibaba Cloud account that is used to create the resource.
	CreatorId *string `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	// The ID of the level-2 quota.
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The name of the level-2 quota.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The alias of the level-2 quota.
	NickName *string `json:"nickName,omitempty" xml:"nickName,omitempty"`
	// The description of the quota.
	Parameter map[string]interface{} `json:"parameter,omitempty" xml:"parameter,omitempty"`
	// The ID of the parent resource.
	ParentId *string `json:"parentId,omitempty" xml:"parentId,omitempty"`
	// The region ID.
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The identifier of an object in a MaxCompute quota. This identifier is the same as the identifier in the sales bill of Alibaba Cloud. This parameter is used for tags.
	SaleTag *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListSaleTag `json:"saleTag,omitempty" xml:"saleTag,omitempty" type:"Struct"`
	// The information of the scheduling plan.
	ScheduleInfo *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListScheduleInfo `json:"scheduleInfo,omitempty" xml:"scheduleInfo,omitempty" type:"Struct"`
	// The status of the endpoint group.
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The tag of the resource for the quota.
	Tag *string `json:"tag,omitempty" xml:"tag,omitempty"`
	// The tenant ID.
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The type of the resource system. This parameter corresponds to the resourceSystemType parameter of the cluster.
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The version number.
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList) String() string {
	return tea.Prettify(s)
}

func (s ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList) GoString() string {
	return s.String()
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList) SetBillingPolicy(v *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListBillingPolicy) *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList {
	s.BillingPolicy = v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList) SetCluster(v string) *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList {
	s.Cluster = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList) SetCreateTime(v int64) *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList {
	s.CreateTime = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList) SetCreatorId(v string) *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList {
	s.CreatorId = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList) SetId(v string) *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList {
	s.Id = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList) SetName(v string) *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList {
	s.Name = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList) SetNickName(v string) *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList {
	s.NickName = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList) SetParameter(v map[string]interface{}) *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList {
	s.Parameter = v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList) SetParentId(v string) *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList {
	s.ParentId = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList) SetRegionId(v string) *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList {
	s.RegionId = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList) SetSaleTag(v *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListSaleTag) *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList {
	s.SaleTag = v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList) SetScheduleInfo(v *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListScheduleInfo) *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList {
	s.ScheduleInfo = v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList) SetStatus(v string) *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList {
	s.Status = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList) SetTag(v string) *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList {
	s.Tag = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList) SetTenantId(v string) *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList {
	s.TenantId = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList) SetType(v string) *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList {
	s.Type = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList) SetVersion(v string) *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoList {
	s.Version = &v
	return s
}

type ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListBillingPolicy struct {
	// The billing method of the quota. Valid values:
	//
	// *   subscription: a subscription quota.
	// *   payasyougo: a pay-as-you-go quota.
	BillingMethod *string `json:"billingMethod,omitempty" xml:"billingMethod,omitempty"`
	// The specifications of the order.
	OdpsSpecCode *string `json:"odpsSpecCode,omitempty" xml:"odpsSpecCode,omitempty"`
	// The order ID.
	OrderId *string `json:"orderId,omitempty" xml:"orderId,omitempty"`
}

func (s ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListBillingPolicy) String() string {
	return tea.Prettify(s)
}

func (s ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListBillingPolicy) GoString() string {
	return s.String()
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListBillingPolicy) SetBillingMethod(v string) *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListBillingPolicy {
	s.BillingMethod = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListBillingPolicy) SetOdpsSpecCode(v string) *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListBillingPolicy {
	s.OdpsSpecCode = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListBillingPolicy) SetOrderId(v string) *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListBillingPolicy {
	s.OrderId = &v
	return s
}

type ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListSaleTag struct {
	// The identifier of an object in a MaxCompute quota. This identifier exists in the sales bill of Alibaba Cloud. You can use this identifier to associate the cost of a quota object with a tag.
	ResourceIds []*string `json:"resourceIds,omitempty" xml:"resourceIds,omitempty" type:"Repeated"`
	// The type of the object. Valid values: quota and project.
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
}

func (s ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListSaleTag) String() string {
	return tea.Prettify(s)
}

func (s ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListSaleTag) GoString() string {
	return s.String()
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListSaleTag) SetResourceIds(v []*string) *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListSaleTag {
	s.ResourceIds = v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListSaleTag) SetResourceType(v string) *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListSaleTag {
	s.ResourceType = &v
	return s
}

type ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListScheduleInfo struct {
	// The quota plan that takes effect based on the scheduling plan.
	CurrPlan *string `json:"currPlan,omitempty" xml:"currPlan,omitempty"`
	// The time when the current quota plan is scheduled.
	CurrTime *string `json:"currTime,omitempty" xml:"currTime,omitempty"`
	// The next quota plan that will take effect based on the scheduling plan.
	NextPlan *string `json:"nextPlan,omitempty" xml:"nextPlan,omitempty"`
	// The time when the next quota plan is scheduled.
	NextTime *string `json:"nextTime,omitempty" xml:"nextTime,omitempty"`
	// The quota plan that immediately takes effect. If the quota plan that immediately takes effect is different from the current quota plan, this parameter is not empty.
	OncePlan *string `json:"oncePlan,omitempty" xml:"oncePlan,omitempty"`
	// The time when the quota plan immediately takes effect.
	OnceTime *string `json:"onceTime,omitempty" xml:"onceTime,omitempty"`
	// The name of the operator.
	OperatorName *string `json:"operatorName,omitempty" xml:"operatorName,omitempty"`
	// The time zone of the project.
	Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
}

func (s ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListScheduleInfo) String() string {
	return tea.Prettify(s)
}

func (s ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListScheduleInfo) GoString() string {
	return s.String()
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListScheduleInfo) SetCurrPlan(v string) *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListScheduleInfo {
	s.CurrPlan = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListScheduleInfo) SetCurrTime(v string) *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListScheduleInfo {
	s.CurrTime = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListScheduleInfo) SetNextPlan(v string) *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListScheduleInfo {
	s.NextPlan = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListScheduleInfo) SetNextTime(v string) *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListScheduleInfo {
	s.NextTime = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListScheduleInfo) SetOncePlan(v string) *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListScheduleInfo {
	s.OncePlan = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListScheduleInfo) SetOnceTime(v string) *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListScheduleInfo {
	s.OnceTime = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListScheduleInfo) SetOperatorName(v string) *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListScheduleInfo {
	s.OperatorName = &v
	return s
}

func (s *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListScheduleInfo) SetTimezone(v string) *ListQuotasResponseBodyQuotaInfoListSubQuotaInfoListScheduleInfo {
	s.Timezone = &v
	return s
}

type ListQuotasResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListQuotasResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListQuotasResponse) String() string {
	return tea.Prettify(s)
}

func (s ListQuotasResponse) GoString() string {
	return s.String()
}

func (s *ListQuotasResponse) SetHeaders(v map[string]*string) *ListQuotasResponse {
	s.Headers = v
	return s
}

func (s *ListQuotasResponse) SetStatusCode(v int32) *ListQuotasResponse {
	s.StatusCode = &v
	return s
}

func (s *ListQuotasResponse) SetBody(v *ListQuotasResponseBody) *ListQuotasResponse {
	s.Body = v
	return s
}

type ListQuotasPlansRequest struct {
	// The ID of the region.
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// The ID of the tenant.
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ListQuotasPlansRequest) String() string {
	return tea.Prettify(s)
}

func (s ListQuotasPlansRequest) GoString() string {
	return s.String()
}

func (s *ListQuotasPlansRequest) SetRegion(v string) *ListQuotasPlansRequest {
	s.Region = &v
	return s
}

func (s *ListQuotasPlansRequest) SetTenantId(v string) *ListQuotasPlansRequest {
	s.TenantId = &v
	return s
}

type ListQuotasPlansResponseBody struct {
	// The returned data.
	Data *ListQuotasPlansResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListQuotasPlansResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListQuotasPlansResponseBody) GoString() string {
	return s.String()
}

func (s *ListQuotasPlansResponseBody) SetData(v *ListQuotasPlansResponseBodyData) *ListQuotasPlansResponseBody {
	s.Data = v
	return s
}

func (s *ListQuotasPlansResponseBody) SetRequestId(v string) *ListQuotasPlansResponseBody {
	s.RequestId = &v
	return s
}

type ListQuotasPlansResponseBodyData struct {
	// The list of quota plans.
	PlanList []*ListQuotasPlansResponseBodyDataPlanList `json:"planList,omitempty" xml:"planList,omitempty" type:"Repeated"`
}

func (s ListQuotasPlansResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListQuotasPlansResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListQuotasPlansResponseBodyData) SetPlanList(v []*ListQuotasPlansResponseBodyDataPlanList) *ListQuotasPlansResponseBodyData {
	s.PlanList = v
	return s
}

type ListQuotasPlansResponseBodyDataPlanList struct {
	// The time when the quota plan was created.
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The name of the quota plan.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The details of the quota.
	Quota *ListQuotasPlansResponseBodyDataPlanListQuota `json:"quota,omitempty" xml:"quota,omitempty" type:"Struct"`
}

func (s ListQuotasPlansResponseBodyDataPlanList) String() string {
	return tea.Prettify(s)
}

func (s ListQuotasPlansResponseBodyDataPlanList) GoString() string {
	return s.String()
}

func (s *ListQuotasPlansResponseBodyDataPlanList) SetCreateTime(v string) *ListQuotasPlansResponseBodyDataPlanList {
	s.CreateTime = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanList) SetName(v string) *ListQuotasPlansResponseBodyDataPlanList {
	s.Name = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanList) SetQuota(v *ListQuotasPlansResponseBodyDataPlanListQuota) *ListQuotasPlansResponseBodyDataPlanList {
	s.Quota = v
	return s
}

type ListQuotasPlansResponseBodyDataPlanListQuota struct {
	// The information of the order.
	BillingPolicy *ListQuotasPlansResponseBodyDataPlanListQuotaBillingPolicy `json:"billingPolicy,omitempty" xml:"billingPolicy,omitempty" type:"Struct"`
	// The ID of the cluster.
	Cluster *string `json:"cluster,omitempty" xml:"cluster,omitempty"`
	// The time when the quota plan was created.
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The ID of the Alibaba Cloud account that is used to create the resource.
	CreatorId *string `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	// The ID of the quota.
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The name of the quota.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The alias of the quota.
	NickName *string `json:"nickName,omitempty" xml:"nickName,omitempty"`
	// The description of the quota.
	Parameter map[string]interface{} `json:"parameter,omitempty" xml:"parameter,omitempty"`
	// The ID of the parent resource.
	ParentId *string `json:"parentId,omitempty" xml:"parentId,omitempty"`
	// The ID of the region.
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The information of the scheduling plan.
	ScheduleInfo *ListQuotasPlansResponseBodyDataPlanListQuotaScheduleInfo `json:"scheduleInfo,omitempty" xml:"scheduleInfo,omitempty" type:"Struct"`
	// The status of the resource.
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The information of the level-2 quota.
	SubQuotaInfoList []*ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList `json:"subQuotaInfoList,omitempty" xml:"subQuotaInfoList,omitempty" type:"Repeated"`
	// The tag of the resource for the quota.
	Tag *string `json:"tag,omitempty" xml:"tag,omitempty"`
	// The ID of the tenant.
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The type of the resource system. This parameter corresponds to the resourceSystemType parameter of the cluster.
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The version number.
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s ListQuotasPlansResponseBodyDataPlanListQuota) String() string {
	return tea.Prettify(s)
}

func (s ListQuotasPlansResponseBodyDataPlanListQuota) GoString() string {
	return s.String()
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuota) SetBillingPolicy(v *ListQuotasPlansResponseBodyDataPlanListQuotaBillingPolicy) *ListQuotasPlansResponseBodyDataPlanListQuota {
	s.BillingPolicy = v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuota) SetCluster(v string) *ListQuotasPlansResponseBodyDataPlanListQuota {
	s.Cluster = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuota) SetCreateTime(v int64) *ListQuotasPlansResponseBodyDataPlanListQuota {
	s.CreateTime = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuota) SetCreatorId(v string) *ListQuotasPlansResponseBodyDataPlanListQuota {
	s.CreatorId = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuota) SetId(v string) *ListQuotasPlansResponseBodyDataPlanListQuota {
	s.Id = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuota) SetName(v string) *ListQuotasPlansResponseBodyDataPlanListQuota {
	s.Name = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuota) SetNickName(v string) *ListQuotasPlansResponseBodyDataPlanListQuota {
	s.NickName = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuota) SetParameter(v map[string]interface{}) *ListQuotasPlansResponseBodyDataPlanListQuota {
	s.Parameter = v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuota) SetParentId(v string) *ListQuotasPlansResponseBodyDataPlanListQuota {
	s.ParentId = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuota) SetRegionId(v string) *ListQuotasPlansResponseBodyDataPlanListQuota {
	s.RegionId = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuota) SetScheduleInfo(v *ListQuotasPlansResponseBodyDataPlanListQuotaScheduleInfo) *ListQuotasPlansResponseBodyDataPlanListQuota {
	s.ScheduleInfo = v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuota) SetStatus(v string) *ListQuotasPlansResponseBodyDataPlanListQuota {
	s.Status = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuota) SetSubQuotaInfoList(v []*ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList) *ListQuotasPlansResponseBodyDataPlanListQuota {
	s.SubQuotaInfoList = v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuota) SetTag(v string) *ListQuotasPlansResponseBodyDataPlanListQuota {
	s.Tag = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuota) SetTenantId(v string) *ListQuotasPlansResponseBodyDataPlanListQuota {
	s.TenantId = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuota) SetType(v string) *ListQuotasPlansResponseBodyDataPlanListQuota {
	s.Type = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuota) SetVersion(v string) *ListQuotasPlansResponseBodyDataPlanListQuota {
	s.Version = &v
	return s
}

type ListQuotasPlansResponseBodyDataPlanListQuotaBillingPolicy struct {
	// The billing method of the quota. Valid values:
	//
	// *   subscription: a subscription quota.
	// *   payasyougo: a pay-as-you-go quota.
	BillingMethod *string `json:"billingMethod,omitempty" xml:"billingMethod,omitempty"`
	// The specifications of the order.
	OdpsSpecCode *string `json:"odpsSpecCode,omitempty" xml:"odpsSpecCode,omitempty"`
	// The ID of the order.
	OrderId *string `json:"orderId,omitempty" xml:"orderId,omitempty"`
}

func (s ListQuotasPlansResponseBodyDataPlanListQuotaBillingPolicy) String() string {
	return tea.Prettify(s)
}

func (s ListQuotasPlansResponseBodyDataPlanListQuotaBillingPolicy) GoString() string {
	return s.String()
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaBillingPolicy) SetBillingMethod(v string) *ListQuotasPlansResponseBodyDataPlanListQuotaBillingPolicy {
	s.BillingMethod = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaBillingPolicy) SetOdpsSpecCode(v string) *ListQuotasPlansResponseBodyDataPlanListQuotaBillingPolicy {
	s.OdpsSpecCode = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaBillingPolicy) SetOrderId(v string) *ListQuotasPlansResponseBodyDataPlanListQuotaBillingPolicy {
	s.OrderId = &v
	return s
}

type ListQuotasPlansResponseBodyDataPlanListQuotaScheduleInfo struct {
	// The quota plan that takes effect based on the scheduling plan.
	CurrPlan *string `json:"currPlan,omitempty" xml:"currPlan,omitempty"`
	// The time when the current quota plan is scheduled.
	CurrTime *string `json:"currTime,omitempty" xml:"currTime,omitempty"`
	// The next quota plan that will take effect based on the scheduling plan.
	NextPlan *string `json:"nextPlan,omitempty" xml:"nextPlan,omitempty"`
	// The time when the next quota plan is scheduled.
	NextTime *string `json:"nextTime,omitempty" xml:"nextTime,omitempty"`
	// If the quota plan that immediately takes effect is different from the current quota plan, this parameter is not empty.
	OncePlan *string `json:"oncePlan,omitempty" xml:"oncePlan,omitempty"`
	// The time when the quota plan immediately takes effect.
	OnceTime *string `json:"onceTime,omitempty" xml:"onceTime,omitempty"`
	// The name of the operator.
	OperatorName *string `json:"operatorName,omitempty" xml:"operatorName,omitempty"`
}

func (s ListQuotasPlansResponseBodyDataPlanListQuotaScheduleInfo) String() string {
	return tea.Prettify(s)
}

func (s ListQuotasPlansResponseBodyDataPlanListQuotaScheduleInfo) GoString() string {
	return s.String()
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaScheduleInfo) SetCurrPlan(v string) *ListQuotasPlansResponseBodyDataPlanListQuotaScheduleInfo {
	s.CurrPlan = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaScheduleInfo) SetCurrTime(v string) *ListQuotasPlansResponseBodyDataPlanListQuotaScheduleInfo {
	s.CurrTime = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaScheduleInfo) SetNextPlan(v string) *ListQuotasPlansResponseBodyDataPlanListQuotaScheduleInfo {
	s.NextPlan = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaScheduleInfo) SetNextTime(v string) *ListQuotasPlansResponseBodyDataPlanListQuotaScheduleInfo {
	s.NextTime = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaScheduleInfo) SetOncePlan(v string) *ListQuotasPlansResponseBodyDataPlanListQuotaScheduleInfo {
	s.OncePlan = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaScheduleInfo) SetOnceTime(v string) *ListQuotasPlansResponseBodyDataPlanListQuotaScheduleInfo {
	s.OnceTime = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaScheduleInfo) SetOperatorName(v string) *ListQuotasPlansResponseBodyDataPlanListQuotaScheduleInfo {
	s.OperatorName = &v
	return s
}

type ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList struct {
	// The information of the order.
	BillingPolicy *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoListBillingPolicy `json:"billingPolicy,omitempty" xml:"billingPolicy,omitempty" type:"Struct"`
	// The ID of the cluster.
	Cluster *string `json:"cluster,omitempty" xml:"cluster,omitempty"`
	// The time when the quota plan was created.
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The ID of the Alibaba Cloud account that is used to create the resource.
	CreatorId *string `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	// The ID of the level-2 quota.
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The name of the level-2 quota.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The nickname of the level-2 quota.
	NickName *string `json:"nickName,omitempty" xml:"nickName,omitempty"`
	// The description of the quota.
	Parameter map[string]interface{} `json:"parameter,omitempty" xml:"parameter,omitempty"`
	// The ID of the parent resource.
	ParentId *string `json:"parentId,omitempty" xml:"parentId,omitempty"`
	// The ID of the region.
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The information of the scheduling plan.
	ScheduleInfo *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoListScheduleInfo `json:"scheduleInfo,omitempty" xml:"scheduleInfo,omitempty" type:"Struct"`
	// The status of the resource.
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The tag of the resource for the quota.
	Tag *string `json:"tag,omitempty" xml:"tag,omitempty"`
	// The ID of the tenant.
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The type of the resource system. This parameter corresponds to the resourceSystemType parameter of the cluster.
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The version number.
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList) String() string {
	return tea.Prettify(s)
}

func (s ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList) GoString() string {
	return s.String()
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList) SetBillingPolicy(v *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoListBillingPolicy) *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList {
	s.BillingPolicy = v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList) SetCluster(v string) *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList {
	s.Cluster = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList) SetCreateTime(v int64) *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList {
	s.CreateTime = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList) SetCreatorId(v string) *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList {
	s.CreatorId = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList) SetId(v string) *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList {
	s.Id = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList) SetName(v string) *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList {
	s.Name = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList) SetNickName(v string) *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList {
	s.NickName = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList) SetParameter(v map[string]interface{}) *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList {
	s.Parameter = v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList) SetParentId(v string) *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList {
	s.ParentId = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList) SetRegionId(v string) *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList {
	s.RegionId = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList) SetScheduleInfo(v *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoListScheduleInfo) *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList {
	s.ScheduleInfo = v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList) SetStatus(v string) *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList {
	s.Status = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList) SetTag(v string) *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList {
	s.Tag = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList) SetTenantId(v string) *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList {
	s.TenantId = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList) SetType(v string) *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList {
	s.Type = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList) SetVersion(v string) *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoList {
	s.Version = &v
	return s
}

type ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoListBillingPolicy struct {
	// The billing method of the quota. Valid values:
	//
	// *   subscription: a subscription quota.
	// *   payasyougo: a pay-as-you-go quota.
	BillingMethod *string `json:"billingMethod,omitempty" xml:"billingMethod,omitempty"`
	// The specifications of the order.
	OdpsSpecCode *string `json:"odpsSpecCode,omitempty" xml:"odpsSpecCode,omitempty"`
	// The ID of the order.
	OrderId *string `json:"orderId,omitempty" xml:"orderId,omitempty"`
}

func (s ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoListBillingPolicy) String() string {
	return tea.Prettify(s)
}

func (s ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoListBillingPolicy) GoString() string {
	return s.String()
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoListBillingPolicy) SetBillingMethod(v string) *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoListBillingPolicy {
	s.BillingMethod = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoListBillingPolicy) SetOdpsSpecCode(v string) *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoListBillingPolicy {
	s.OdpsSpecCode = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoListBillingPolicy) SetOrderId(v string) *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoListBillingPolicy {
	s.OrderId = &v
	return s
}

type ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoListScheduleInfo struct {
	// The quota plan that takes effect based on the scheduling plan.
	CurrPlan *string `json:"currPlan,omitempty" xml:"currPlan,omitempty"`
	// The time when the current quota plan is scheduled.
	CurrTime *string `json:"currTime,omitempty" xml:"currTime,omitempty"`
	// The next quota plan that will take effect based on the scheduling plan.
	NextPlan *string `json:"nextPlan,omitempty" xml:"nextPlan,omitempty"`
	// The time when the next quota plan is scheduled.
	NextTime *string `json:"nextTime,omitempty" xml:"nextTime,omitempty"`
	// If the quota plan that immediately takes effect is different from the current quota plan, this parameter is not empty.
	OncePlan *string `json:"oncePlan,omitempty" xml:"oncePlan,omitempty"`
	// The time when the quota plan immediately takes effect.
	OnceTime *string `json:"onceTime,omitempty" xml:"onceTime,omitempty"`
	// The name of the operator.
	OperatorName *string `json:"operatorName,omitempty" xml:"operatorName,omitempty"`
}

func (s ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoListScheduleInfo) String() string {
	return tea.Prettify(s)
}

func (s ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoListScheduleInfo) GoString() string {
	return s.String()
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoListScheduleInfo) SetCurrPlan(v string) *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoListScheduleInfo {
	s.CurrPlan = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoListScheduleInfo) SetCurrTime(v string) *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoListScheduleInfo {
	s.CurrTime = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoListScheduleInfo) SetNextPlan(v string) *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoListScheduleInfo {
	s.NextPlan = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoListScheduleInfo) SetNextTime(v string) *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoListScheduleInfo {
	s.NextTime = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoListScheduleInfo) SetOncePlan(v string) *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoListScheduleInfo {
	s.OncePlan = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoListScheduleInfo) SetOnceTime(v string) *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoListScheduleInfo {
	s.OnceTime = &v
	return s
}

func (s *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoListScheduleInfo) SetOperatorName(v string) *ListQuotasPlansResponseBodyDataPlanListQuotaSubQuotaInfoListScheduleInfo {
	s.OperatorName = &v
	return s
}

type ListQuotasPlansResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListQuotasPlansResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListQuotasPlansResponse) String() string {
	return tea.Prettify(s)
}

func (s ListQuotasPlansResponse) GoString() string {
	return s.String()
}

func (s *ListQuotasPlansResponse) SetHeaders(v map[string]*string) *ListQuotasPlansResponse {
	s.Headers = v
	return s
}

func (s *ListQuotasPlansResponse) SetStatusCode(v int32) *ListQuotasPlansResponse {
	s.StatusCode = &v
	return s
}

func (s *ListQuotasPlansResponse) SetBody(v *ListQuotasPlansResponseBody) *ListQuotasPlansResponse {
	s.Body = v
	return s
}

type ListResourcesRequest struct {
	// Specifies the marker after which the returned list begins.
	Marker *string `json:"marker,omitempty" xml:"marker,omitempty"`
	// The maximum number of entries to return on each page.
	MaxItem *int32 `json:"maxItem,omitempty" xml:"maxItem,omitempty"`
	// The name of the resource.
	Name       *string `json:"name,omitempty" xml:"name,omitempty"`
	SchemaName *string `json:"schemaName,omitempty" xml:"schemaName,omitempty"`
}

func (s ListResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListResourcesRequest) SetMarker(v string) *ListResourcesRequest {
	s.Marker = &v
	return s
}

func (s *ListResourcesRequest) SetMaxItem(v int32) *ListResourcesRequest {
	s.MaxItem = &v
	return s
}

func (s *ListResourcesRequest) SetName(v string) *ListResourcesRequest {
	s.Name = &v
	return s
}

func (s *ListResourcesRequest) SetSchemaName(v string) *ListResourcesRequest {
	s.SchemaName = &v
	return s
}

type ListResourcesResponseBody struct {
	// The returned data.
	Data *ListResourcesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListResourcesResponseBody) SetData(v *ListResourcesResponseBodyData) *ListResourcesResponseBody {
	s.Data = v
	return s
}

func (s *ListResourcesResponseBody) SetRequestId(v string) *ListResourcesResponseBody {
	s.RequestId = &v
	return s
}

type ListResourcesResponseBodyData struct {
	// Indicates the marker after which the returned list begins.
	Marker *string `json:"marker,omitempty" xml:"marker,omitempty"`
	// The maximum number of entries returned per page.
	MaxItem *int32 `json:"maxItem,omitempty" xml:"maxItem,omitempty"`
	// The list of resources.
	Resources []*ListResourcesResponseBodyDataResources `json:"resources,omitempty" xml:"resources,omitempty" type:"Repeated"`
}

func (s ListResourcesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListResourcesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListResourcesResponseBodyData) SetMarker(v string) *ListResourcesResponseBodyData {
	s.Marker = &v
	return s
}

func (s *ListResourcesResponseBodyData) SetMaxItem(v int32) *ListResourcesResponseBodyData {
	s.MaxItem = &v
	return s
}

func (s *ListResourcesResponseBodyData) SetResources(v []*ListResourcesResponseBodyDataResources) *ListResourcesResponseBodyData {
	s.Resources = v
	return s
}

type ListResourcesResponseBodyDataResources struct {
	Comment    *string `json:"comment,omitempty" xml:"comment,omitempty"`
	ContentMD5 *string `json:"contentMD5,omitempty" xml:"contentMD5,omitempty"`
	// The time when the resource was created.
	CreationTime     *int64  `json:"creationTime,omitempty" xml:"creationTime,omitempty"`
	DisplayName      *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	LastModifiedTime *int64  `json:"lastModifiedTime,omitempty" xml:"lastModifiedTime,omitempty"`
	LastUpdator      *string `json:"lastUpdator,omitempty" xml:"lastUpdator,omitempty"`
	// The name of the resource.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The owner of the resource.
	Owner *string `json:"owner,omitempty" xml:"owner,omitempty"`
	// The schema of the resource.
	Schema *string `json:"schema,omitempty" xml:"schema,omitempty"`
	Size   *int64  `json:"size,omitempty" xml:"size,omitempty"`
	// The type of the resource.
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListResourcesResponseBodyDataResources) String() string {
	return tea.Prettify(s)
}

func (s ListResourcesResponseBodyDataResources) GoString() string {
	return s.String()
}

func (s *ListResourcesResponseBodyDataResources) SetComment(v string) *ListResourcesResponseBodyDataResources {
	s.Comment = &v
	return s
}

func (s *ListResourcesResponseBodyDataResources) SetContentMD5(v string) *ListResourcesResponseBodyDataResources {
	s.ContentMD5 = &v
	return s
}

func (s *ListResourcesResponseBodyDataResources) SetCreationTime(v int64) *ListResourcesResponseBodyDataResources {
	s.CreationTime = &v
	return s
}

func (s *ListResourcesResponseBodyDataResources) SetDisplayName(v string) *ListResourcesResponseBodyDataResources {
	s.DisplayName = &v
	return s
}

func (s *ListResourcesResponseBodyDataResources) SetLastModifiedTime(v int64) *ListResourcesResponseBodyDataResources {
	s.LastModifiedTime = &v
	return s
}

func (s *ListResourcesResponseBodyDataResources) SetLastUpdator(v string) *ListResourcesResponseBodyDataResources {
	s.LastUpdator = &v
	return s
}

func (s *ListResourcesResponseBodyDataResources) SetName(v string) *ListResourcesResponseBodyDataResources {
	s.Name = &v
	return s
}

func (s *ListResourcesResponseBodyDataResources) SetOwner(v string) *ListResourcesResponseBodyDataResources {
	s.Owner = &v
	return s
}

func (s *ListResourcesResponseBodyDataResources) SetSchema(v string) *ListResourcesResponseBodyDataResources {
	s.Schema = &v
	return s
}

func (s *ListResourcesResponseBodyDataResources) SetSize(v int64) *ListResourcesResponseBodyDataResources {
	s.Size = &v
	return s
}

func (s *ListResourcesResponseBodyDataResources) SetType(v string) *ListResourcesResponseBodyDataResources {
	s.Type = &v
	return s
}

type ListResourcesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListResourcesResponse) SetHeaders(v map[string]*string) *ListResourcesResponse {
	s.Headers = v
	return s
}

func (s *ListResourcesResponse) SetStatusCode(v int32) *ListResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListResourcesResponse) SetBody(v *ListResourcesResponseBody) *ListResourcesResponse {
	s.Body = v
	return s
}

type ListRolesResponseBody struct {
	// The returned data.
	Data *ListRolesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListRolesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRolesResponseBody) GoString() string {
	return s.String()
}

func (s *ListRolesResponseBody) SetData(v *ListRolesResponseBodyData) *ListRolesResponseBody {
	s.Data = v
	return s
}

func (s *ListRolesResponseBody) SetRequestId(v string) *ListRolesResponseBody {
	s.RequestId = &v
	return s
}

type ListRolesResponseBodyData struct {
	// The MaxCompute project-level roles.
	Roles []*ListRolesResponseBodyDataRoles `json:"roles,omitempty" xml:"roles,omitempty" type:"Repeated"`
}

func (s ListRolesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListRolesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListRolesResponseBodyData) SetRoles(v []*ListRolesResponseBodyDataRoles) *ListRolesResponseBodyData {
	s.Roles = v
	return s
}

type ListRolesResponseBodyDataRoles struct {
	// The ACL-based permissions that are granted to the role.
	Acl *ListRolesResponseBodyDataRolesAcl `json:"acl,omitempty" xml:"acl,omitempty" type:"Struct"`
	// The name of the role.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The policy that is attached to the role.
	Policy *string `json:"policy,omitempty" xml:"policy,omitempty"`
	// The type of the role.
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListRolesResponseBodyDataRoles) String() string {
	return tea.Prettify(s)
}

func (s ListRolesResponseBodyDataRoles) GoString() string {
	return s.String()
}

func (s *ListRolesResponseBodyDataRoles) SetAcl(v *ListRolesResponseBodyDataRolesAcl) *ListRolesResponseBodyDataRoles {
	s.Acl = v
	return s
}

func (s *ListRolesResponseBodyDataRoles) SetName(v string) *ListRolesResponseBodyDataRoles {
	s.Name = &v
	return s
}

func (s *ListRolesResponseBodyDataRoles) SetPolicy(v string) *ListRolesResponseBodyDataRoles {
	s.Policy = &v
	return s
}

func (s *ListRolesResponseBodyDataRoles) SetType(v string) *ListRolesResponseBodyDataRoles {
	s.Type = &v
	return s
}

type ListRolesResponseBodyDataRolesAcl struct {
	// The function.
	Function []*ListRolesResponseBodyDataRolesAclFunction `json:"function,omitempty" xml:"function,omitempty" type:"Repeated"`
	// The instance.
	Instance []*ListRolesResponseBodyDataRolesAclInstance `json:"instance,omitempty" xml:"instance,omitempty" type:"Repeated"`
	// The package.
	Package []*ListRolesResponseBodyDataRolesAclPackage `json:"package,omitempty" xml:"package,omitempty" type:"Repeated"`
	// The project.
	Project []*ListRolesResponseBodyDataRolesAclProject `json:"project,omitempty" xml:"project,omitempty" type:"Repeated"`
	// The resource.
	Resource []*ListRolesResponseBodyDataRolesAclResource `json:"resource,omitempty" xml:"resource,omitempty" type:"Repeated"`
	// The table.
	Table []*ListRolesResponseBodyDataRolesAclTable `json:"table,omitempty" xml:"table,omitempty" type:"Repeated"`
}

func (s ListRolesResponseBodyDataRolesAcl) String() string {
	return tea.Prettify(s)
}

func (s ListRolesResponseBodyDataRolesAcl) GoString() string {
	return s.String()
}

func (s *ListRolesResponseBodyDataRolesAcl) SetFunction(v []*ListRolesResponseBodyDataRolesAclFunction) *ListRolesResponseBodyDataRolesAcl {
	s.Function = v
	return s
}

func (s *ListRolesResponseBodyDataRolesAcl) SetInstance(v []*ListRolesResponseBodyDataRolesAclInstance) *ListRolesResponseBodyDataRolesAcl {
	s.Instance = v
	return s
}

func (s *ListRolesResponseBodyDataRolesAcl) SetPackage(v []*ListRolesResponseBodyDataRolesAclPackage) *ListRolesResponseBodyDataRolesAcl {
	s.Package = v
	return s
}

func (s *ListRolesResponseBodyDataRolesAcl) SetProject(v []*ListRolesResponseBodyDataRolesAclProject) *ListRolesResponseBodyDataRolesAcl {
	s.Project = v
	return s
}

func (s *ListRolesResponseBodyDataRolesAcl) SetResource(v []*ListRolesResponseBodyDataRolesAclResource) *ListRolesResponseBodyDataRolesAcl {
	s.Resource = v
	return s
}

func (s *ListRolesResponseBodyDataRolesAcl) SetTable(v []*ListRolesResponseBodyDataRolesAclTable) *ListRolesResponseBodyDataRolesAcl {
	s.Table = v
	return s
}

type ListRolesResponseBodyDataRolesAclFunction struct {
	// The operations that were performed on the function.
	Actions []*string `json:"actions,omitempty" xml:"actions,omitempty" type:"Repeated"`
	// The name of the function.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ListRolesResponseBodyDataRolesAclFunction) String() string {
	return tea.Prettify(s)
}

func (s ListRolesResponseBodyDataRolesAclFunction) GoString() string {
	return s.String()
}

func (s *ListRolesResponseBodyDataRolesAclFunction) SetActions(v []*string) *ListRolesResponseBodyDataRolesAclFunction {
	s.Actions = v
	return s
}

func (s *ListRolesResponseBodyDataRolesAclFunction) SetName(v string) *ListRolesResponseBodyDataRolesAclFunction {
	s.Name = &v
	return s
}

type ListRolesResponseBodyDataRolesAclInstance struct {
	// The operations that were performed on the instance.
	Actions []*string `json:"actions,omitempty" xml:"actions,omitempty" type:"Repeated"`
	// The name of the instance.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ListRolesResponseBodyDataRolesAclInstance) String() string {
	return tea.Prettify(s)
}

func (s ListRolesResponseBodyDataRolesAclInstance) GoString() string {
	return s.String()
}

func (s *ListRolesResponseBodyDataRolesAclInstance) SetActions(v []*string) *ListRolesResponseBodyDataRolesAclInstance {
	s.Actions = v
	return s
}

func (s *ListRolesResponseBodyDataRolesAclInstance) SetName(v string) *ListRolesResponseBodyDataRolesAclInstance {
	s.Name = &v
	return s
}

type ListRolesResponseBodyDataRolesAclPackage struct {
	// The operations that were performed on the package.
	Actions []*string `json:"actions,omitempty" xml:"actions,omitempty" type:"Repeated"`
	// The name of the package.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ListRolesResponseBodyDataRolesAclPackage) String() string {
	return tea.Prettify(s)
}

func (s ListRolesResponseBodyDataRolesAclPackage) GoString() string {
	return s.String()
}

func (s *ListRolesResponseBodyDataRolesAclPackage) SetActions(v []*string) *ListRolesResponseBodyDataRolesAclPackage {
	s.Actions = v
	return s
}

func (s *ListRolesResponseBodyDataRolesAclPackage) SetName(v string) *ListRolesResponseBodyDataRolesAclPackage {
	s.Name = &v
	return s
}

type ListRolesResponseBodyDataRolesAclProject struct {
	// The operations that were performed on the project.
	Actions []*string `json:"actions,omitempty" xml:"actions,omitempty" type:"Repeated"`
	// The name of the MaxCompute project.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ListRolesResponseBodyDataRolesAclProject) String() string {
	return tea.Prettify(s)
}

func (s ListRolesResponseBodyDataRolesAclProject) GoString() string {
	return s.String()
}

func (s *ListRolesResponseBodyDataRolesAclProject) SetActions(v []*string) *ListRolesResponseBodyDataRolesAclProject {
	s.Actions = v
	return s
}

func (s *ListRolesResponseBodyDataRolesAclProject) SetName(v string) *ListRolesResponseBodyDataRolesAclProject {
	s.Name = &v
	return s
}

type ListRolesResponseBodyDataRolesAclResource struct {
	// The operations that were performed on the resource.
	Actions []*string `json:"actions,omitempty" xml:"actions,omitempty" type:"Repeated"`
	// The name of the resource.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ListRolesResponseBodyDataRolesAclResource) String() string {
	return tea.Prettify(s)
}

func (s ListRolesResponseBodyDataRolesAclResource) GoString() string {
	return s.String()
}

func (s *ListRolesResponseBodyDataRolesAclResource) SetActions(v []*string) *ListRolesResponseBodyDataRolesAclResource {
	s.Actions = v
	return s
}

func (s *ListRolesResponseBodyDataRolesAclResource) SetName(v string) *ListRolesResponseBodyDataRolesAclResource {
	s.Name = &v
	return s
}

type ListRolesResponseBodyDataRolesAclTable struct {
	// The operations that were performed on the table.
	Actions []*string `json:"actions,omitempty" xml:"actions,omitempty" type:"Repeated"`
	// The name of the table.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ListRolesResponseBodyDataRolesAclTable) String() string {
	return tea.Prettify(s)
}

func (s ListRolesResponseBodyDataRolesAclTable) GoString() string {
	return s.String()
}

func (s *ListRolesResponseBodyDataRolesAclTable) SetActions(v []*string) *ListRolesResponseBodyDataRolesAclTable {
	s.Actions = v
	return s
}

func (s *ListRolesResponseBodyDataRolesAclTable) SetName(v string) *ListRolesResponseBodyDataRolesAclTable {
	s.Name = &v
	return s
}

type ListRolesResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListRolesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListRolesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRolesResponse) GoString() string {
	return s.String()
}

func (s *ListRolesResponse) SetHeaders(v map[string]*string) *ListRolesResponse {
	s.Headers = v
	return s
}

func (s *ListRolesResponse) SetStatusCode(v int32) *ListRolesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRolesResponse) SetBody(v *ListRolesResponseBody) *ListRolesResponse {
	s.Body = v
	return s
}

type ListTablesRequest struct {
	// Specifies the marker after which the returned list begins.
	Marker *string `json:"marker,omitempty" xml:"marker,omitempty"`
	// The maximum number of entries to return on each page.
	MaxItem *int32 `json:"maxItem,omitempty" xml:"maxItem,omitempty"`
	// The names of the returned resources. The names must start with the value specified by the prefix parameter. If the prefix parameter is set to a, the names of the returned resources must start with a.
	Prefix     *string `json:"prefix,omitempty" xml:"prefix,omitempty"`
	SchemaName *string `json:"schemaName,omitempty" xml:"schemaName,omitempty"`
	// The type of the table.
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListTablesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTablesRequest) GoString() string {
	return s.String()
}

func (s *ListTablesRequest) SetMarker(v string) *ListTablesRequest {
	s.Marker = &v
	return s
}

func (s *ListTablesRequest) SetMaxItem(v int32) *ListTablesRequest {
	s.MaxItem = &v
	return s
}

func (s *ListTablesRequest) SetPrefix(v string) *ListTablesRequest {
	s.Prefix = &v
	return s
}

func (s *ListTablesRequest) SetSchemaName(v string) *ListTablesRequest {
	s.SchemaName = &v
	return s
}

func (s *ListTablesRequest) SetType(v string) *ListTablesRequest {
	s.Type = &v
	return s
}

type ListTablesResponseBody struct {
	// The returned data.
	Data *ListTablesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListTablesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTablesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTablesResponseBody) SetData(v *ListTablesResponseBodyData) *ListTablesResponseBody {
	s.Data = v
	return s
}

func (s *ListTablesResponseBody) SetRequestId(v string) *ListTablesResponseBody {
	s.RequestId = &v
	return s
}

type ListTablesResponseBodyData struct {
	// Indicates the marker after which the returned list begins.
	Marker *string `json:"marker,omitempty" xml:"marker,omitempty"`
	// The maximum number of entries returned per page.
	MaxItem *int32 `json:"maxItem,omitempty" xml:"maxItem,omitempty"`
	// The information about tables.
	Tables []*ListTablesResponseBodyDataTables `json:"tables,omitempty" xml:"tables,omitempty" type:"Repeated"`
}

func (s ListTablesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListTablesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListTablesResponseBodyData) SetMarker(v string) *ListTablesResponseBodyData {
	s.Marker = &v
	return s
}

func (s *ListTablesResponseBodyData) SetMaxItem(v int32) *ListTablesResponseBodyData {
	s.MaxItem = &v
	return s
}

func (s *ListTablesResponseBodyData) SetTables(v []*ListTablesResponseBodyDataTables) *ListTablesResponseBodyData {
	s.Tables = v
	return s
}

type ListTablesResponseBodyDataTables struct {
	AutoRefreshEnabled *bool   `json:"autoRefreshEnabled,omitempty" xml:"autoRefreshEnabled,omitempty"`
	CreateTableDDL     *string `json:"createTableDDL,omitempty" xml:"createTableDDL,omitempty"`
	// The time when the table was created.
	CreationTime     *int64  `json:"creationTime,omitempty" xml:"creationTime,omitempty"`
	DisplayName      *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	FileNum          *int64  `json:"fileNum,omitempty" xml:"fileNum,omitempty"`
	IsExternalTable  *bool   `json:"isExternalTable,omitempty" xml:"isExternalTable,omitempty"`
	IsOutdated       *bool   `json:"isOutdated,omitempty" xml:"isOutdated,omitempty"`
	LastAccessTime   *int64  `json:"lastAccessTime,omitempty" xml:"lastAccessTime,omitempty"`
	LastDDLTime      *int64  `json:"lastDDLTime,omitempty" xml:"lastDDLTime,omitempty"`
	LastModifiedTime *int64  `json:"lastModifiedTime,omitempty" xml:"lastModifiedTime,omitempty"`
	Lifecycle        *string `json:"lifecycle,omitempty" xml:"lifecycle,omitempty"`
	Location         *string `json:"location,omitempty" xml:"location,omitempty"`
	MaterializedView *bool   `json:"materializedView,omitempty" xml:"materializedView,omitempty"`
	// The name of the table.
	Name                           *string                                          `json:"name,omitempty" xml:"name,omitempty"`
	NativeColumns                  []*ListTablesResponseBodyDataTablesNativeColumns `json:"nativeColumns,omitempty" xml:"nativeColumns,omitempty" type:"Repeated"`
	OdpsPropertiesRolearn          *string                                          `json:"odpsPropertiesRolearn,omitempty" xml:"odpsPropertiesRolearn,omitempty"`
	OdpsSqlTextOptionFlushHeader   *bool                                            `json:"odpsSqlTextOptionFlushHeader,omitempty" xml:"odpsSqlTextOptionFlushHeader,omitempty"`
	OdpsTextOptionHeaderLinesCount *int64                                           `json:"odpsTextOptionHeaderLinesCount,omitempty" xml:"odpsTextOptionHeaderLinesCount,omitempty"`
	// The owner of the table.
	Owner            *string                                             `json:"owner,omitempty" xml:"owner,omitempty"`
	PartitionColumns []*ListTablesResponseBodyDataTablesPartitionColumns `json:"partitionColumns,omitempty" xml:"partitionColumns,omitempty" type:"Repeated"`
	PhysicalSize     *int64                                              `json:"physicalSize,omitempty" xml:"physicalSize,omitempty"`
	ProjectName      *string                                             `json:"projectName,omitempty" xml:"projectName,omitempty"`
	RewriteEnabled   *bool                                               `json:"rewriteEnabled,omitempty" xml:"rewriteEnabled,omitempty"`
	// The schema of the table.
	Schema                   *string `json:"schema,omitempty" xml:"schema,omitempty"`
	Size                     *int64  `json:"size,omitempty" xml:"size,omitempty"`
	StorageHandler           *string `json:"storageHandler,omitempty" xml:"storageHandler,omitempty"`
	TableComment             *string `json:"tableComment,omitempty" xml:"tableComment,omitempty"`
	TableLabel               *string `json:"tableLabel,omitempty" xml:"tableLabel,omitempty"`
	TablesotreTableName      *string `json:"tablesotreTableName,omitempty" xml:"tablesotreTableName,omitempty"`
	TablestoreColumnsMapping *string `json:"tablestoreColumnsMapping,omitempty" xml:"tablestoreColumnsMapping,omitempty"`
	// The type of the table.
	Type     *string `json:"type,omitempty" xml:"type,omitempty"`
	ViewText *string `json:"viewText,omitempty" xml:"viewText,omitempty"`
}

func (s ListTablesResponseBodyDataTables) String() string {
	return tea.Prettify(s)
}

func (s ListTablesResponseBodyDataTables) GoString() string {
	return s.String()
}

func (s *ListTablesResponseBodyDataTables) SetAutoRefreshEnabled(v bool) *ListTablesResponseBodyDataTables {
	s.AutoRefreshEnabled = &v
	return s
}

func (s *ListTablesResponseBodyDataTables) SetCreateTableDDL(v string) *ListTablesResponseBodyDataTables {
	s.CreateTableDDL = &v
	return s
}

func (s *ListTablesResponseBodyDataTables) SetCreationTime(v int64) *ListTablesResponseBodyDataTables {
	s.CreationTime = &v
	return s
}

func (s *ListTablesResponseBodyDataTables) SetDisplayName(v string) *ListTablesResponseBodyDataTables {
	s.DisplayName = &v
	return s
}

func (s *ListTablesResponseBodyDataTables) SetFileNum(v int64) *ListTablesResponseBodyDataTables {
	s.FileNum = &v
	return s
}

func (s *ListTablesResponseBodyDataTables) SetIsExternalTable(v bool) *ListTablesResponseBodyDataTables {
	s.IsExternalTable = &v
	return s
}

func (s *ListTablesResponseBodyDataTables) SetIsOutdated(v bool) *ListTablesResponseBodyDataTables {
	s.IsOutdated = &v
	return s
}

func (s *ListTablesResponseBodyDataTables) SetLastAccessTime(v int64) *ListTablesResponseBodyDataTables {
	s.LastAccessTime = &v
	return s
}

func (s *ListTablesResponseBodyDataTables) SetLastDDLTime(v int64) *ListTablesResponseBodyDataTables {
	s.LastDDLTime = &v
	return s
}

func (s *ListTablesResponseBodyDataTables) SetLastModifiedTime(v int64) *ListTablesResponseBodyDataTables {
	s.LastModifiedTime = &v
	return s
}

func (s *ListTablesResponseBodyDataTables) SetLifecycle(v string) *ListTablesResponseBodyDataTables {
	s.Lifecycle = &v
	return s
}

func (s *ListTablesResponseBodyDataTables) SetLocation(v string) *ListTablesResponseBodyDataTables {
	s.Location = &v
	return s
}

func (s *ListTablesResponseBodyDataTables) SetMaterializedView(v bool) *ListTablesResponseBodyDataTables {
	s.MaterializedView = &v
	return s
}

func (s *ListTablesResponseBodyDataTables) SetName(v string) *ListTablesResponseBodyDataTables {
	s.Name = &v
	return s
}

func (s *ListTablesResponseBodyDataTables) SetNativeColumns(v []*ListTablesResponseBodyDataTablesNativeColumns) *ListTablesResponseBodyDataTables {
	s.NativeColumns = v
	return s
}

func (s *ListTablesResponseBodyDataTables) SetOdpsPropertiesRolearn(v string) *ListTablesResponseBodyDataTables {
	s.OdpsPropertiesRolearn = &v
	return s
}

func (s *ListTablesResponseBodyDataTables) SetOdpsSqlTextOptionFlushHeader(v bool) *ListTablesResponseBodyDataTables {
	s.OdpsSqlTextOptionFlushHeader = &v
	return s
}

func (s *ListTablesResponseBodyDataTables) SetOdpsTextOptionHeaderLinesCount(v int64) *ListTablesResponseBodyDataTables {
	s.OdpsTextOptionHeaderLinesCount = &v
	return s
}

func (s *ListTablesResponseBodyDataTables) SetOwner(v string) *ListTablesResponseBodyDataTables {
	s.Owner = &v
	return s
}

func (s *ListTablesResponseBodyDataTables) SetPartitionColumns(v []*ListTablesResponseBodyDataTablesPartitionColumns) *ListTablesResponseBodyDataTables {
	s.PartitionColumns = v
	return s
}

func (s *ListTablesResponseBodyDataTables) SetPhysicalSize(v int64) *ListTablesResponseBodyDataTables {
	s.PhysicalSize = &v
	return s
}

func (s *ListTablesResponseBodyDataTables) SetProjectName(v string) *ListTablesResponseBodyDataTables {
	s.ProjectName = &v
	return s
}

func (s *ListTablesResponseBodyDataTables) SetRewriteEnabled(v bool) *ListTablesResponseBodyDataTables {
	s.RewriteEnabled = &v
	return s
}

func (s *ListTablesResponseBodyDataTables) SetSchema(v string) *ListTablesResponseBodyDataTables {
	s.Schema = &v
	return s
}

func (s *ListTablesResponseBodyDataTables) SetSize(v int64) *ListTablesResponseBodyDataTables {
	s.Size = &v
	return s
}

func (s *ListTablesResponseBodyDataTables) SetStorageHandler(v string) *ListTablesResponseBodyDataTables {
	s.StorageHandler = &v
	return s
}

func (s *ListTablesResponseBodyDataTables) SetTableComment(v string) *ListTablesResponseBodyDataTables {
	s.TableComment = &v
	return s
}

func (s *ListTablesResponseBodyDataTables) SetTableLabel(v string) *ListTablesResponseBodyDataTables {
	s.TableLabel = &v
	return s
}

func (s *ListTablesResponseBodyDataTables) SetTablesotreTableName(v string) *ListTablesResponseBodyDataTables {
	s.TablesotreTableName = &v
	return s
}

func (s *ListTablesResponseBodyDataTables) SetTablestoreColumnsMapping(v string) *ListTablesResponseBodyDataTables {
	s.TablestoreColumnsMapping = &v
	return s
}

func (s *ListTablesResponseBodyDataTables) SetType(v string) *ListTablesResponseBodyDataTables {
	s.Type = &v
	return s
}

func (s *ListTablesResponseBodyDataTables) SetViewText(v string) *ListTablesResponseBodyDataTables {
	s.ViewText = &v
	return s
}

type ListTablesResponseBodyDataTablesNativeColumns struct {
	Comment *string `json:"comment,omitempty" xml:"comment,omitempty"`
	Label   *string `json:"label,omitempty" xml:"label,omitempty"`
	Name    *string `json:"name,omitempty" xml:"name,omitempty"`
	Type    *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListTablesResponseBodyDataTablesNativeColumns) String() string {
	return tea.Prettify(s)
}

func (s ListTablesResponseBodyDataTablesNativeColumns) GoString() string {
	return s.String()
}

func (s *ListTablesResponseBodyDataTablesNativeColumns) SetComment(v string) *ListTablesResponseBodyDataTablesNativeColumns {
	s.Comment = &v
	return s
}

func (s *ListTablesResponseBodyDataTablesNativeColumns) SetLabel(v string) *ListTablesResponseBodyDataTablesNativeColumns {
	s.Label = &v
	return s
}

func (s *ListTablesResponseBodyDataTablesNativeColumns) SetName(v string) *ListTablesResponseBodyDataTablesNativeColumns {
	s.Name = &v
	return s
}

func (s *ListTablesResponseBodyDataTablesNativeColumns) SetType(v string) *ListTablesResponseBodyDataTablesNativeColumns {
	s.Type = &v
	return s
}

type ListTablesResponseBodyDataTablesPartitionColumns struct {
	Comment *string `json:"comment,omitempty" xml:"comment,omitempty"`
	Label   *string `json:"label,omitempty" xml:"label,omitempty"`
	Name    *string `json:"name,omitempty" xml:"name,omitempty"`
	Type    *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListTablesResponseBodyDataTablesPartitionColumns) String() string {
	return tea.Prettify(s)
}

func (s ListTablesResponseBodyDataTablesPartitionColumns) GoString() string {
	return s.String()
}

func (s *ListTablesResponseBodyDataTablesPartitionColumns) SetComment(v string) *ListTablesResponseBodyDataTablesPartitionColumns {
	s.Comment = &v
	return s
}

func (s *ListTablesResponseBodyDataTablesPartitionColumns) SetLabel(v string) *ListTablesResponseBodyDataTablesPartitionColumns {
	s.Label = &v
	return s
}

func (s *ListTablesResponseBodyDataTablesPartitionColumns) SetName(v string) *ListTablesResponseBodyDataTablesPartitionColumns {
	s.Name = &v
	return s
}

func (s *ListTablesResponseBodyDataTablesPartitionColumns) SetType(v string) *ListTablesResponseBodyDataTablesPartitionColumns {
	s.Type = &v
	return s
}

type ListTablesResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListTablesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTablesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTablesResponse) GoString() string {
	return s.String()
}

func (s *ListTablesResponse) SetHeaders(v map[string]*string) *ListTablesResponse {
	s.Headers = v
	return s
}

func (s *ListTablesResponse) SetStatusCode(v int32) *ListTablesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTablesResponse) SetBody(v *ListTablesResponseBody) *ListTablesResponse {
	s.Body = v
	return s
}

type ListUsersRequest struct {
	// The number of the page to return.
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries to return on each page.
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUsersRequest) GoString() string {
	return s.String()
}

func (s *ListUsersRequest) SetPageNumber(v int32) *ListUsersRequest {
	s.PageNumber = &v
	return s
}

func (s *ListUsersRequest) SetPageSize(v int32) *ListUsersRequest {
	s.PageSize = &v
	return s
}

type ListUsersResponseBody struct {
	// The returned data.
	Data *ListUsersResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUsersResponseBody) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBody) SetData(v *ListUsersResponseBodyData) *ListUsersResponseBody {
	s.Data = v
	return s
}

func (s *ListUsersResponseBody) SetRequestId(v string) *ListUsersResponseBody {
	s.RequestId = &v
	return s
}

type ListUsersResponseBodyData struct {
	// The page number of the returned page.
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries returned per page.
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The total number of returned entries.
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
	// The users.
	Users []*ListUsersResponseBodyDataUsers `json:"users,omitempty" xml:"users,omitempty" type:"Repeated"`
}

func (s ListUsersResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListUsersResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBodyData) SetPageNumber(v int32) *ListUsersResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListUsersResponseBodyData) SetPageSize(v int32) *ListUsersResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListUsersResponseBodyData) SetTotalCount(v int32) *ListUsersResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListUsersResponseBodyData) SetUsers(v []*ListUsersResponseBodyDataUsers) *ListUsersResponseBodyData {
	s.Users = v
	return s
}

type ListUsersResponseBodyDataUsers struct {
	// The ID of the Alibaba Cloud account.
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	// The username of the account.
	AccountName *string `json:"accountName,omitempty" xml:"accountName,omitempty"`
	// The type of the account.
	AccountType *string `json:"accountType,omitempty" xml:"accountType,omitempty"`
	// The display name.
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// The ID of the tenant.
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ListUsersResponseBodyDataUsers) String() string {
	return tea.Prettify(s)
}

func (s ListUsersResponseBodyDataUsers) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBodyDataUsers) SetAccountId(v string) *ListUsersResponseBodyDataUsers {
	s.AccountId = &v
	return s
}

func (s *ListUsersResponseBodyDataUsers) SetAccountName(v string) *ListUsersResponseBodyDataUsers {
	s.AccountName = &v
	return s
}

func (s *ListUsersResponseBodyDataUsers) SetAccountType(v string) *ListUsersResponseBodyDataUsers {
	s.AccountType = &v
	return s
}

func (s *ListUsersResponseBodyDataUsers) SetDisplayName(v string) *ListUsersResponseBodyDataUsers {
	s.DisplayName = &v
	return s
}

func (s *ListUsersResponseBodyDataUsers) SetTenantId(v string) *ListUsersResponseBodyDataUsers {
	s.TenantId = &v
	return s
}

type ListUsersResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListUsersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUsersResponse) GoString() string {
	return s.String()
}

func (s *ListUsersResponse) SetHeaders(v map[string]*string) *ListUsersResponse {
	s.Headers = v
	return s
}

func (s *ListUsersResponse) SetStatusCode(v int32) *ListUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUsersResponse) SetBody(v *ListUsersResponseBody) *ListUsersResponse {
	s.Body = v
	return s
}

type ListUsersByRoleResponseBody struct {
	// The returned data.
	Data *ListUsersByRoleResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListUsersByRoleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUsersByRoleResponseBody) GoString() string {
	return s.String()
}

func (s *ListUsersByRoleResponseBody) SetData(v *ListUsersByRoleResponseBodyData) *ListUsersByRoleResponseBody {
	s.Data = v
	return s
}

func (s *ListUsersByRoleResponseBody) SetRequestId(v string) *ListUsersByRoleResponseBody {
	s.RequestId = &v
	return s
}

type ListUsersByRoleResponseBodyData struct {
	// The users.
	Users []*ListUsersByRoleResponseBodyDataUsers `json:"users,omitempty" xml:"users,omitempty" type:"Repeated"`
}

func (s ListUsersByRoleResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListUsersByRoleResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListUsersByRoleResponseBodyData) SetUsers(v []*ListUsersByRoleResponseBodyDataUsers) *ListUsersByRoleResponseBodyData {
	s.Users = v
	return s
}

type ListUsersByRoleResponseBodyDataUsers struct {
	// The name of the user.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ListUsersByRoleResponseBodyDataUsers) String() string {
	return tea.Prettify(s)
}

func (s ListUsersByRoleResponseBodyDataUsers) GoString() string {
	return s.String()
}

func (s *ListUsersByRoleResponseBodyDataUsers) SetName(v string) *ListUsersByRoleResponseBodyDataUsers {
	s.Name = &v
	return s
}

type ListUsersByRoleResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListUsersByRoleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListUsersByRoleResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUsersByRoleResponse) GoString() string {
	return s.String()
}

func (s *ListUsersByRoleResponse) SetHeaders(v map[string]*string) *ListUsersByRoleResponse {
	s.Headers = v
	return s
}

func (s *ListUsersByRoleResponse) SetStatusCode(v int32) *ListUsersByRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUsersByRoleResponse) SetBody(v *ListUsersByRoleResponseBody) *ListUsersByRoleResponse {
	s.Body = v
	return s
}

type UpdatePackageRequest struct {
	// The request body parameters.
	Body *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePackageRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdatePackageRequest) GoString() string {
	return s.String()
}

func (s *UpdatePackageRequest) SetBody(v string) *UpdatePackageRequest {
	s.Body = &v
	return s
}

type UpdatePackageResponseBody struct {
	// The returned data.
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdatePackageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdatePackageResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePackageResponseBody) SetData(v string) *UpdatePackageResponseBody {
	s.Data = &v
	return s
}

func (s *UpdatePackageResponseBody) SetRequestId(v string) *UpdatePackageResponseBody {
	s.RequestId = &v
	return s
}

type UpdatePackageResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdatePackageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdatePackageResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdatePackageResponse) GoString() string {
	return s.String()
}

func (s *UpdatePackageResponse) SetHeaders(v map[string]*string) *UpdatePackageResponse {
	s.Headers = v
	return s
}

func (s *UpdatePackageResponse) SetStatusCode(v int32) *UpdatePackageResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePackageResponse) SetBody(v *UpdatePackageResponseBody) *UpdatePackageResponse {
	s.Body = v
	return s
}

type UpdateProjectIpWhiteListRequest struct {
	// The request body parameters.
	Body *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateProjectIpWhiteListRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateProjectIpWhiteListRequest) GoString() string {
	return s.String()
}

func (s *UpdateProjectIpWhiteListRequest) SetBody(v string) *UpdateProjectIpWhiteListRequest {
	s.Body = &v
	return s
}

type UpdateProjectIpWhiteListResponseBody struct {
	// The returned result.
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateProjectIpWhiteListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateProjectIpWhiteListResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateProjectIpWhiteListResponseBody) SetData(v string) *UpdateProjectIpWhiteListResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateProjectIpWhiteListResponseBody) SetRequestId(v string) *UpdateProjectIpWhiteListResponseBody {
	s.RequestId = &v
	return s
}

type UpdateProjectIpWhiteListResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateProjectIpWhiteListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateProjectIpWhiteListResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateProjectIpWhiteListResponse) GoString() string {
	return s.String()
}

func (s *UpdateProjectIpWhiteListResponse) SetHeaders(v map[string]*string) *UpdateProjectIpWhiteListResponse {
	s.Headers = v
	return s
}

func (s *UpdateProjectIpWhiteListResponse) SetStatusCode(v int32) *UpdateProjectIpWhiteListResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateProjectIpWhiteListResponse) SetBody(v *UpdateProjectIpWhiteListResponseBody) *UpdateProjectIpWhiteListResponse {
	s.Body = v
	return s
}

type UpdateQuotaHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AkProven      *string            `json:"AkProven,omitempty" xml:"AkProven,omitempty"`
}

func (s UpdateQuotaHeaders) String() string {
	return tea.Prettify(s)
}

func (s UpdateQuotaHeaders) GoString() string {
	return s.String()
}

func (s *UpdateQuotaHeaders) SetCommonHeaders(v map[string]*string) *UpdateQuotaHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateQuotaHeaders) SetAkProven(v string) *UpdateQuotaHeaders {
	s.AkProven = &v
	return s
}

type UpdateQuotaRequest struct {
	Body     *string `json:"body,omitempty" xml:"body,omitempty"`
	Region   *string `json:"region,omitempty" xml:"region,omitempty"`
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s UpdateQuotaRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateQuotaRequest) GoString() string {
	return s.String()
}

func (s *UpdateQuotaRequest) SetBody(v string) *UpdateQuotaRequest {
	s.Body = &v
	return s
}

func (s *UpdateQuotaRequest) SetRegion(v string) *UpdateQuotaRequest {
	s.Region = &v
	return s
}

func (s *UpdateQuotaRequest) SetTenantId(v string) *UpdateQuotaRequest {
	s.TenantId = &v
	return s
}

type UpdateQuotaResponseBody struct {
	Data      *string `json:"data,omitempty" xml:"data,omitempty"`
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateQuotaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateQuotaResponseBody) SetData(v string) *UpdateQuotaResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateQuotaResponseBody) SetRequestId(v string) *UpdateQuotaResponseBody {
	s.RequestId = &v
	return s
}

type UpdateQuotaResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateQuotaResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateQuotaResponse) GoString() string {
	return s.String()
}

func (s *UpdateQuotaResponse) SetHeaders(v map[string]*string) *UpdateQuotaResponse {
	s.Headers = v
	return s
}

func (s *UpdateQuotaResponse) SetStatusCode(v int32) *UpdateQuotaResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateQuotaResponse) SetBody(v *UpdateQuotaResponseBody) *UpdateQuotaResponse {
	s.Body = v
	return s
}

type UpdateQuotaPlanRequest struct {
	// The request body parameters.
	Body *string `json:"body,omitempty" xml:"body,omitempty"`
	// The ID of the region.
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// The ID of the tenant.
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s UpdateQuotaPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateQuotaPlanRequest) GoString() string {
	return s.String()
}

func (s *UpdateQuotaPlanRequest) SetBody(v string) *UpdateQuotaPlanRequest {
	s.Body = &v
	return s
}

func (s *UpdateQuotaPlanRequest) SetRegion(v string) *UpdateQuotaPlanRequest {
	s.Region = &v
	return s
}

func (s *UpdateQuotaPlanRequest) SetTenantId(v string) *UpdateQuotaPlanRequest {
	s.TenantId = &v
	return s
}

type UpdateQuotaPlanResponseBody struct {
	// The returned result.
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateQuotaPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateQuotaPlanResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateQuotaPlanResponseBody) SetData(v string) *UpdateQuotaPlanResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateQuotaPlanResponseBody) SetRequestId(v string) *UpdateQuotaPlanResponseBody {
	s.RequestId = &v
	return s
}

type UpdateQuotaPlanResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateQuotaPlanResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateQuotaPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateQuotaPlanResponse) GoString() string {
	return s.String()
}

func (s *UpdateQuotaPlanResponse) SetHeaders(v map[string]*string) *UpdateQuotaPlanResponse {
	s.Headers = v
	return s
}

func (s *UpdateQuotaPlanResponse) SetStatusCode(v int32) *UpdateQuotaPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateQuotaPlanResponse) SetBody(v *UpdateQuotaPlanResponseBody) *UpdateQuotaPlanResponse {
	s.Body = v
	return s
}

type UpdateQuotaScheduleRequest struct {
	// The request body parameters.
	Body *string `json:"body,omitempty" xml:"body,omitempty"`
	// The ID of the region.
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// The ID of the tenant.
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s UpdateQuotaScheduleRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateQuotaScheduleRequest) GoString() string {
	return s.String()
}

func (s *UpdateQuotaScheduleRequest) SetBody(v string) *UpdateQuotaScheduleRequest {
	s.Body = &v
	return s
}

func (s *UpdateQuotaScheduleRequest) SetRegion(v string) *UpdateQuotaScheduleRequest {
	s.Region = &v
	return s
}

func (s *UpdateQuotaScheduleRequest) SetTenantId(v string) *UpdateQuotaScheduleRequest {
	s.TenantId = &v
	return s
}

type UpdateQuotaScheduleResponseBody struct {
	// The returned result.
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateQuotaScheduleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateQuotaScheduleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateQuotaScheduleResponseBody) SetData(v string) *UpdateQuotaScheduleResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateQuotaScheduleResponseBody) SetRequestId(v string) *UpdateQuotaScheduleResponseBody {
	s.RequestId = &v
	return s
}

type UpdateQuotaScheduleResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateQuotaScheduleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateQuotaScheduleResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateQuotaScheduleResponse) GoString() string {
	return s.String()
}

func (s *UpdateQuotaScheduleResponse) SetHeaders(v map[string]*string) *UpdateQuotaScheduleResponse {
	s.Headers = v
	return s
}

func (s *UpdateQuotaScheduleResponse) SetStatusCode(v int32) *UpdateQuotaScheduleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateQuotaScheduleResponse) SetBody(v *UpdateQuotaScheduleResponseBody) *UpdateQuotaScheduleResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("regional")
	client.EndpointMap = map[string]*string{
		"ap-northeast-1":              tea.String("maxcompute.aliyuncs.com"),
		"ap-northeast-2-pop":          tea.String("maxcompute.aliyuncs.com"),
		"ap-south-1":                  tea.String("maxcompute.aliyuncs.com"),
		"ap-southeast-1":              tea.String("maxcompute.aliyuncs.com"),
		"ap-southeast-2":              tea.String("maxcompute.aliyuncs.com"),
		"ap-southeast-3":              tea.String("maxcompute.aliyuncs.com"),
		"ap-southeast-5":              tea.String("maxcompute.aliyuncs.com"),
		"cn-beijing":                  tea.String("maxcompute.aliyuncs.com"),
		"cn-beijing-finance-1":        tea.String("maxcompute.aliyuncs.com"),
		"cn-beijing-finance-pop":      tea.String("maxcompute.aliyuncs.com"),
		"cn-beijing-gov-1":            tea.String("maxcompute.aliyuncs.com"),
		"cn-beijing-nu16-b01":         tea.String("maxcompute.aliyuncs.com"),
		"cn-chengdu":                  tea.String("maxcompute.aliyuncs.com"),
		"cn-edge-1":                   tea.String("maxcompute.aliyuncs.com"),
		"cn-fujian":                   tea.String("maxcompute.aliyuncs.com"),
		"cn-haidian-cm12-c01":         tea.String("maxcompute.aliyuncs.com"),
		"cn-hangzhou":                 tea.String("maxcompute.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          tea.String("maxcompute.aliyuncs.com"),
		"cn-hangzhou-finance":         tea.String("maxcompute.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": tea.String("maxcompute.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": tea.String("maxcompute.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": tea.String("maxcompute.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": tea.String("maxcompute.aliyuncs.com"),
		"cn-hangzhou-test-306":        tea.String("maxcompute.aliyuncs.com"),
		"cn-hongkong":                 tea.String("maxcompute.aliyuncs.com"),
		"cn-hongkong-finance-pop":     tea.String("maxcompute.aliyuncs.com"),
		"cn-huhehaote":                tea.String("maxcompute.aliyuncs.com"),
		"cn-north-2-gov-1":            tea.String("maxcompute.aliyuncs.com"),
		"cn-qingdao":                  tea.String("maxcompute.aliyuncs.com"),
		"cn-qingdao-nebula":           tea.String("maxcompute.aliyuncs.com"),
		"cn-shanghai":                 tea.String("maxcompute.aliyuncs.com"),
		"cn-shanghai-et15-b01":        tea.String("maxcompute.aliyuncs.com"),
		"cn-shanghai-et2-b01":         tea.String("maxcompute.aliyuncs.com"),
		"cn-shanghai-finance-1":       tea.String("maxcompute.aliyuncs.com"),
		"cn-shanghai-inner":           tea.String("maxcompute.aliyuncs.com"),
		"cn-shanghai-internal-test-1": tea.String("maxcompute.aliyuncs.com"),
		"cn-shenzhen":                 tea.String("maxcompute.aliyuncs.com"),
		"cn-shenzhen-finance-1":       tea.String("maxcompute.aliyuncs.com"),
		"cn-shenzhen-inner":           tea.String("maxcompute.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         tea.String("maxcompute.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        tea.String("maxcompute.aliyuncs.com"),
		"cn-wuhan":                    tea.String("maxcompute.aliyuncs.com"),
		"cn-yushanfang":               tea.String("maxcompute.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        tea.String("maxcompute.aliyuncs.com"),
		"cn-zhangjiakou":              tea.String("maxcompute.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     tea.String("maxcompute.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       tea.String("maxcompute.aliyuncs.com"),
		"eu-central-1":                tea.String("maxcompute.aliyuncs.com"),
		"eu-west-1":                   tea.String("maxcompute.aliyuncs.com"),
		"eu-west-1-oxs":               tea.String("maxcompute.aliyuncs.com"),
		"me-east-1":                   tea.String("maxcompute.aliyuncs.com"),
		"rus-west-1-pop":              tea.String("maxcompute.aliyuncs.com"),
		"us-east-1":                   tea.String("maxcompute.aliyuncs.com"),
		"us-west-1":                   tea.String("maxcompute.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("maxcompute"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreatePackageWithOptions(projectName *string, request *CreatePackageRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreatePackageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IsInstall)) {
		query["isInstall"] = request.IsInstall
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("CreatePackage"),
		Version:     tea.String("2022-01-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/projects/" + tea.StringValue(openapiutil.GetEncodeParam(projectName)) + "/packages"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreatePackageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreatePackage(projectName *string, request *CreatePackageRequest) (_result *CreatePackageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreatePackageResponse{}
	_body, _err := client.CreatePackageWithOptions(projectName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateProjectWithOptions(request *CreateProjectRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("CreateProject"),
		Version:     tea.String("2022-01-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/projects"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateProject(request *CreateProjectRequest) (_result *CreateProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateProjectResponse{}
	_body, _err := client.CreateProjectWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateQuotaPlanWithOptions(nickname *string, request *CreateQuotaPlanRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateQuotaPlanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Region)) {
		query["region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		query["tenantId"] = request.TenantId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("CreateQuotaPlan"),
		Version:     tea.String("2022-01-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/quotas/" + tea.StringValue(openapiutil.GetEncodeParam(nickname)) + "/plans"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateQuotaPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateQuotaPlan(nickname *string, request *CreateQuotaPlanRequest) (_result *CreateQuotaPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateQuotaPlanResponse{}
	_body, _err := client.CreateQuotaPlanWithOptions(nickname, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateQuotaScheduleWithOptions(nickname *string, request *CreateQuotaScheduleRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateQuotaScheduleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Region)) {
		query["region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		query["tenantId"] = request.TenantId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("CreateQuotaSchedule"),
		Version:     tea.String("2022-01-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/quotas/" + tea.StringValue(openapiutil.GetEncodeParam(nickname)) + "/schedule"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateQuotaScheduleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateQuotaSchedule(nickname *string, request *CreateQuotaScheduleRequest) (_result *CreateQuotaScheduleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateQuotaScheduleResponse{}
	_body, _err := client.CreateQuotaScheduleWithOptions(nickname, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateRoleWithOptions(projectName *string, request *CreateRoleRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateRoleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("CreateRole"),
		Version:     tea.String("2022-01-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/projects/" + tea.StringValue(openapiutil.GetEncodeParam(projectName)) + "/roles"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateRoleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateRole(projectName *string, request *CreateRoleRequest) (_result *CreateRoleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateRoleResponse{}
	_body, _err := client.CreateRoleWithOptions(projectName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteQuotaPlanWithOptions(nickname *string, planName *string, request *DeleteQuotaPlanRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteQuotaPlanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Region)) {
		query["region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		query["tenantId"] = request.TenantId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteQuotaPlan"),
		Version:     tea.String("2022-01-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/quotas/" + tea.StringValue(openapiutil.GetEncodeParam(nickname)) + "/plans/" + tea.StringValue(openapiutil.GetEncodeParam(planName))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteQuotaPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteQuotaPlan(nickname *string, planName *string, request *DeleteQuotaPlanRequest) (_result *DeleteQuotaPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteQuotaPlanResponse{}
	_body, _err := client.DeleteQuotaPlanWithOptions(nickname, planName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetJobResourceUsageWithOptions(tmpReq *GetJobResourceUsageRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetJobResourceUsageResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetJobResourceUsageShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.JobOwnerList)) {
		request.JobOwnerListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.JobOwnerList, tea.String("jobOwnerList"), tea.String("simple"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.QuotaNicknameList)) {
		request.QuotaNicknameListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.QuotaNicknameList, tea.String("quotaNicknameList"), tea.String("simple"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Date)) {
		query["date"] = request.Date
	}

	if !tea.BoolValue(util.IsUnset(request.JobOwnerListShrink)) {
		query["jobOwnerList"] = request.JobOwnerListShrink
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.QuotaNicknameListShrink)) {
		query["quotaNicknameList"] = request.QuotaNicknameListShrink
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetJobResourceUsage"),
		Version:     tea.String("2022-01-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/jobs/resourceUsage"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetJobResourceUsageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetJobResourceUsage(request *GetJobResourceUsageRequest) (_result *GetJobResourceUsageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetJobResourceUsageResponse{}
	_body, _err := client.GetJobResourceUsageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetPackageWithOptions(projectName *string, packageName *string, request *GetPackageRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetPackageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SourceProject)) {
		query["sourceProject"] = request.SourceProject
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetPackage"),
		Version:     tea.String("2022-01-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/projects/" + tea.StringValue(openapiutil.GetEncodeParam(projectName)) + "/packages/" + tea.StringValue(openapiutil.GetEncodeParam(packageName))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetPackageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetPackage(projectName *string, packageName *string, request *GetPackageRequest) (_result *GetPackageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetPackageResponse{}
	_body, _err := client.GetPackageWithOptions(projectName, packageName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetProjectWithOptions(projectName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetProjectResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetProject"),
		Version:     tea.String("2022-01-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/projects/" + tea.StringValue(openapiutil.GetEncodeParam(projectName))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetProject(projectName *string) (_result *GetProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetProjectResponse{}
	_body, _err := client.GetProjectWithOptions(projectName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetQuotaWithOptions(nickname *string, request *GetQuotaRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetQuotaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AkProven)) {
		query["AkProven"] = request.AkProven
	}

	if !tea.BoolValue(util.IsUnset(request.Mock)) {
		query["mock"] = request.Mock
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		query["region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		query["tenantId"] = request.TenantId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetQuota"),
		Version:     tea.String("2022-01-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/quotas/" + tea.StringValue(openapiutil.GetEncodeParam(nickname))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetQuotaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetQuota(nickname *string, request *GetQuotaRequest) (_result *GetQuotaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetQuotaResponse{}
	_body, _err := client.GetQuotaWithOptions(nickname, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetQuotaPlanWithOptions(nickname *string, planName *string, request *GetQuotaPlanRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetQuotaPlanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Region)) {
		query["region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		query["tenantId"] = request.TenantId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetQuotaPlan"),
		Version:     tea.String("2022-01-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/quotas/" + tea.StringValue(openapiutil.GetEncodeParam(nickname)) + "/plans/" + tea.StringValue(openapiutil.GetEncodeParam(planName))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetQuotaPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetQuotaPlan(nickname *string, planName *string, request *GetQuotaPlanRequest) (_result *GetQuotaPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetQuotaPlanResponse{}
	_body, _err := client.GetQuotaPlanWithOptions(nickname, planName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetQuotaScheduleWithOptions(nickname *string, request *GetQuotaScheduleRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetQuotaScheduleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DisplayTimezone)) {
		query["displayTimezone"] = request.DisplayTimezone
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		query["region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		query["tenantId"] = request.TenantId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetQuotaSchedule"),
		Version:     tea.String("2022-01-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/quotas/" + tea.StringValue(openapiutil.GetEncodeParam(nickname)) + "/schedule"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetQuotaScheduleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetQuotaSchedule(nickname *string, request *GetQuotaScheduleRequest) (_result *GetQuotaScheduleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetQuotaScheduleResponse{}
	_body, _err := client.GetQuotaScheduleWithOptions(nickname, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetRoleAclWithOptions(projectName *string, roleName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetRoleAclResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetRoleAcl"),
		Version:     tea.String("2022-01-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/projects/" + tea.StringValue(openapiutil.GetEncodeParam(projectName)) + "/roles/" + tea.StringValue(openapiutil.GetEncodeParam(roleName)) + "/roleAcl"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetRoleAclResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetRoleAcl(projectName *string, roleName *string) (_result *GetRoleAclResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetRoleAclResponse{}
	_body, _err := client.GetRoleAclWithOptions(projectName, roleName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetRoleAclOnObjectWithOptions(projectName *string, roleName *string, request *GetRoleAclOnObjectRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetRoleAclOnObjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ObjectName)) {
		query["objectName"] = request.ObjectName
	}

	if !tea.BoolValue(util.IsUnset(request.ObjectType)) {
		query["objectType"] = request.ObjectType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetRoleAclOnObject"),
		Version:     tea.String("2022-01-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/projects/" + tea.StringValue(openapiutil.GetEncodeParam(projectName)) + "/roles/" + tea.StringValue(openapiutil.GetEncodeParam(roleName)) + "/acl"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetRoleAclOnObjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetRoleAclOnObject(projectName *string, roleName *string, request *GetRoleAclOnObjectRequest) (_result *GetRoleAclOnObjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetRoleAclOnObjectResponse{}
	_body, _err := client.GetRoleAclOnObjectWithOptions(projectName, roleName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetRolePolicyWithOptions(projectName *string, roleName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetRolePolicyResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetRolePolicy"),
		Version:     tea.String("2022-01-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/projects/" + tea.StringValue(openapiutil.GetEncodeParam(projectName)) + "/roles/" + tea.StringValue(openapiutil.GetEncodeParam(roleName)) + "/policy"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetRolePolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetRolePolicy(projectName *string, roleName *string) (_result *GetRolePolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetRolePolicyResponse{}
	_body, _err := client.GetRolePolicyWithOptions(projectName, roleName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetRunningJobsWithOptions(tmpReq *GetRunningJobsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetRunningJobsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetRunningJobsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.JobOwnerList)) {
		request.JobOwnerListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.JobOwnerList, tea.String("jobOwnerList"), tea.String("simple"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.QuotaNicknameList)) {
		request.QuotaNicknameListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.QuotaNicknameList, tea.String("quotaNicknameList"), tea.String("simple"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.From)) {
		query["from"] = request.From
	}

	if !tea.BoolValue(util.IsUnset(request.JobOwnerListShrink)) {
		query["jobOwnerList"] = request.JobOwnerListShrink
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.QuotaNicknameListShrink)) {
		query["quotaNicknameList"] = request.QuotaNicknameListShrink
	}

	if !tea.BoolValue(util.IsUnset(request.To)) {
		query["to"] = request.To
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetRunningJobs"),
		Version:     tea.String("2022-01-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/jobs/runningJobs"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetRunningJobsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetRunningJobs(request *GetRunningJobsRequest) (_result *GetRunningJobsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetRunningJobsResponse{}
	_body, _err := client.GetRunningJobsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTrustedProjectsWithOptions(projectName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetTrustedProjectsResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetTrustedProjects"),
		Version:     tea.String("2022-01-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/projects/" + tea.StringValue(openapiutil.GetEncodeParam(projectName)) + "/trustedProjects"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTrustedProjectsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTrustedProjects(projectName *string) (_result *GetTrustedProjectsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTrustedProjectsResponse{}
	_body, _err := client.GetTrustedProjectsWithOptions(projectName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) KillJobsWithOptions(request *KillJobsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *KillJobsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Region)) {
		query["region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		query["tenantId"] = request.TenantId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("KillJobs"),
		Version:     tea.String("2022-01-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/jobs/kill"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &KillJobsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) KillJobs(request *KillJobsRequest) (_result *KillJobsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &KillJobsResponse{}
	_body, _err := client.KillJobsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListFunctionsWithOptions(projectName *string, request *ListFunctionsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListFunctionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Marker)) {
		query["marker"] = request.Marker
	}

	if !tea.BoolValue(util.IsUnset(request.MaxItem)) {
		query["maxItem"] = request.MaxItem
	}

	if !tea.BoolValue(util.IsUnset(request.Prefix)) {
		query["prefix"] = request.Prefix
	}

	if !tea.BoolValue(util.IsUnset(request.SchemaName)) {
		query["schemaName"] = request.SchemaName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListFunctions"),
		Version:     tea.String("2022-01-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/projects/" + tea.StringValue(openapiutil.GetEncodeParam(projectName)) + "/functions"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListFunctionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListFunctions(projectName *string, request *ListFunctionsRequest) (_result *ListFunctionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListFunctionsResponse{}
	_body, _err := client.ListFunctionsWithOptions(projectName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListJobInfosWithOptions(request *ListJobInfosRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListJobInfosResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AscOrder)) {
		query["ascOrder"] = request.AscOrder
	}

	if !tea.BoolValue(util.IsUnset(request.OrderColumn)) {
		query["orderColumn"] = request.OrderColumn
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		query["region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		query["tenantId"] = request.TenantId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("ListJobInfos"),
		Version:     tea.String("2022-01-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/jobs"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListJobInfosResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListJobInfos(request *ListJobInfosRequest) (_result *ListJobInfosResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListJobInfosResponse{}
	_body, _err := client.ListJobInfosWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListPackagesWithOptions(projectName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListPackagesResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ListPackages"),
		Version:     tea.String("2022-01-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/projects/" + tea.StringValue(openapiutil.GetEncodeParam(projectName)) + "/packages"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListPackagesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListPackages(projectName *string) (_result *ListPackagesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListPackagesResponse{}
	_body, _err := client.ListPackagesWithOptions(projectName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListProjectUsersWithOptions(projectName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListProjectUsersResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ListProjectUsers"),
		Version:     tea.String("2022-01-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/projects/" + tea.StringValue(openapiutil.GetEncodeParam(projectName)) + "/users"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListProjectUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListProjectUsers(projectName *string) (_result *ListProjectUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListProjectUsersResponse{}
	_body, _err := client.ListProjectUsersWithOptions(projectName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListProjectsWithOptions(request *ListProjectsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListProjectsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ListSystemCatalog)) {
		query["listSystemCatalog"] = request.ListSystemCatalog
	}

	if !tea.BoolValue(util.IsUnset(request.Marker)) {
		query["marker"] = request.Marker
	}

	if !tea.BoolValue(util.IsUnset(request.MaxItem)) {
		query["maxItem"] = request.MaxItem
	}

	if !tea.BoolValue(util.IsUnset(request.Prefix)) {
		query["prefix"] = request.Prefix
	}

	if !tea.BoolValue(util.IsUnset(request.QuotaName)) {
		query["quotaName"] = request.QuotaName
	}

	if !tea.BoolValue(util.IsUnset(request.QuotaNickName)) {
		query["quotaNickName"] = request.QuotaNickName
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		query["region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.SaleTags)) {
		query["saleTags"] = request.SaleTags
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		query["tenantId"] = request.TenantId
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListProjects"),
		Version:     tea.String("2022-01-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/projects"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListProjectsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListProjects(request *ListProjectsRequest) (_result *ListProjectsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListProjectsResponse{}
	_body, _err := client.ListProjectsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListQuotasWithOptions(request *ListQuotasRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListQuotasResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BillingType)) {
		query["billingType"] = request.BillingType
	}

	if !tea.BoolValue(util.IsUnset(request.Marker)) {
		query["marker"] = request.Marker
	}

	if !tea.BoolValue(util.IsUnset(request.MaxItem)) {
		query["maxItem"] = request.MaxItem
	}

	if !tea.BoolValue(util.IsUnset(request.ProductId)) {
		query["productId"] = request.ProductId
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		query["region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.SaleTags)) {
		query["saleTags"] = request.SaleTags
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		query["tenantId"] = request.TenantId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListQuotas"),
		Version:     tea.String("2022-01-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/quotas"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListQuotasResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListQuotas(request *ListQuotasRequest) (_result *ListQuotasResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListQuotasResponse{}
	_body, _err := client.ListQuotasWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListQuotasPlansWithOptions(nickname *string, request *ListQuotasPlansRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListQuotasPlansResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Region)) {
		query["region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		query["tenantId"] = request.TenantId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListQuotasPlans"),
		Version:     tea.String("2022-01-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/quotas/" + tea.StringValue(openapiutil.GetEncodeParam(nickname)) + "/plans"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListQuotasPlansResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListQuotasPlans(nickname *string, request *ListQuotasPlansRequest) (_result *ListQuotasPlansResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListQuotasPlansResponse{}
	_body, _err := client.ListQuotasPlansWithOptions(nickname, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListResourcesWithOptions(projectName *string, request *ListResourcesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Marker)) {
		query["marker"] = request.Marker
	}

	if !tea.BoolValue(util.IsUnset(request.MaxItem)) {
		query["maxItem"] = request.MaxItem
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.SchemaName)) {
		query["schemaName"] = request.SchemaName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListResources"),
		Version:     tea.String("2022-01-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/projects/" + tea.StringValue(openapiutil.GetEncodeParam(projectName)) + "/resources"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListResources(projectName *string, request *ListResourcesRequest) (_result *ListResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListResourcesResponse{}
	_body, _err := client.ListResourcesWithOptions(projectName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListRolesWithOptions(projectName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListRolesResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ListRoles"),
		Version:     tea.String("2022-01-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/projects/" + tea.StringValue(openapiutil.GetEncodeParam(projectName)) + "/roles"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListRolesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListRoles(projectName *string) (_result *ListRolesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListRolesResponse{}
	_body, _err := client.ListRolesWithOptions(projectName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTablesWithOptions(projectName *string, request *ListTablesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListTablesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Marker)) {
		query["marker"] = request.Marker
	}

	if !tea.BoolValue(util.IsUnset(request.MaxItem)) {
		query["maxItem"] = request.MaxItem
	}

	if !tea.BoolValue(util.IsUnset(request.Prefix)) {
		query["prefix"] = request.Prefix
	}

	if !tea.BoolValue(util.IsUnset(request.SchemaName)) {
		query["schemaName"] = request.SchemaName
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTables"),
		Version:     tea.String("2022-01-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/projects/" + tea.StringValue(openapiutil.GetEncodeParam(projectName)) + "/tables"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTablesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTables(projectName *string, request *ListTablesRequest) (_result *ListTablesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListTablesResponse{}
	_body, _err := client.ListTablesWithOptions(projectName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListUsersWithOptions(request *ListUsersRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListUsers"),
		Version:     tea.String("2022-01-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/users"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListUsers(request *ListUsersRequest) (_result *ListUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListUsersResponse{}
	_body, _err := client.ListUsersWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListUsersByRoleWithOptions(projectName *string, roleName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListUsersByRoleResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ListUsersByRole"),
		Version:     tea.String("2022-01-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/projects/" + tea.StringValue(openapiutil.GetEncodeParam(projectName)) + "/roles/" + tea.StringValue(openapiutil.GetEncodeParam(roleName)) + "/users"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListUsersByRoleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListUsersByRole(projectName *string, roleName *string) (_result *ListUsersByRoleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListUsersByRoleResponse{}
	_body, _err := client.ListUsersByRoleWithOptions(projectName, roleName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdatePackageWithOptions(projectName *string, packageName *string, request *UpdatePackageRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdatePackageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("UpdatePackage"),
		Version:     tea.String("2022-01-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/projects/" + tea.StringValue(openapiutil.GetEncodeParam(projectName)) + "/packages/" + tea.StringValue(openapiutil.GetEncodeParam(packageName))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdatePackageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdatePackage(projectName *string, packageName *string, request *UpdatePackageRequest) (_result *UpdatePackageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdatePackageResponse{}
	_body, _err := client.UpdatePackageWithOptions(projectName, packageName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateProjectIpWhiteListWithOptions(projectName *string, request *UpdateProjectIpWhiteListRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateProjectIpWhiteListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateProjectIpWhiteList"),
		Version:     tea.String("2022-01-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/projects/" + tea.StringValue(openapiutil.GetEncodeParam(projectName)) + "/ipWhiteList"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateProjectIpWhiteListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateProjectIpWhiteList(projectName *string, request *UpdateProjectIpWhiteListRequest) (_result *UpdateProjectIpWhiteListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateProjectIpWhiteListResponse{}
	_body, _err := client.UpdateProjectIpWhiteListWithOptions(projectName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateQuotaWithOptions(nickname *string, request *UpdateQuotaRequest, headers *UpdateQuotaHeaders, runtime *util.RuntimeOptions) (_result *UpdateQuotaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Region)) {
		query["region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		query["tenantId"] = request.TenantId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.AkProven)) {
		realHeaders["AkProven"] = util.ToJSONString(headers.AkProven)
	}

	req := &openapi.OpenApiRequest{
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
		Body:    request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateQuota"),
		Version:     tea.String("2022-01-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/quotas/" + tea.StringValue(openapiutil.GetEncodeParam(nickname))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateQuotaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateQuota(nickname *string, request *UpdateQuotaRequest) (_result *UpdateQuotaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UpdateQuotaHeaders{}
	_result = &UpdateQuotaResponse{}
	_body, _err := client.UpdateQuotaWithOptions(nickname, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateQuotaPlanWithOptions(nickname *string, planName *string, request *UpdateQuotaPlanRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateQuotaPlanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Region)) {
		query["region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		query["tenantId"] = request.TenantId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateQuotaPlan"),
		Version:     tea.String("2022-01-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/quotas/" + tea.StringValue(openapiutil.GetEncodeParam(nickname)) + "/plans/" + tea.StringValue(openapiutil.GetEncodeParam(planName))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateQuotaPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateQuotaPlan(nickname *string, planName *string, request *UpdateQuotaPlanRequest) (_result *UpdateQuotaPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateQuotaPlanResponse{}
	_body, _err := client.UpdateQuotaPlanWithOptions(nickname, planName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateQuotaScheduleWithOptions(nickname *string, request *UpdateQuotaScheduleRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateQuotaScheduleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Region)) {
		query["region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		query["tenantId"] = request.TenantId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateQuotaSchedule"),
		Version:     tea.String("2022-01-04"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/quotas/" + tea.StringValue(openapiutil.GetEncodeParam(nickname)) + "/schedule"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateQuotaScheduleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateQuotaSchedule(nickname *string, request *UpdateQuotaScheduleRequest) (_result *UpdateQuotaScheduleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateQuotaScheduleResponse{}
	_body, _err := client.UpdateQuotaScheduleWithOptions(nickname, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
