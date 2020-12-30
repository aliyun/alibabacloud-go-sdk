// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CreateProjectRequest struct {
	SourceIp    *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ProjectType *string `json:"ProjectType,omitempty" xml:"ProjectType,omitempty"`
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s CreateProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectRequest) GoString() string {
	return s.String()
}

func (s *CreateProjectRequest) SetSourceIp(v string) *CreateProjectRequest {
	s.SourceIp = &v
	return s
}

func (s *CreateProjectRequest) SetProjectType(v string) *CreateProjectRequest {
	s.ProjectType = &v
	return s
}

func (s *CreateProjectRequest) SetProjectName(v string) *CreateProjectRequest {
	s.ProjectName = &v
	return s
}

type CreateProjectResponseBody struct {
	RequestId      *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	HttpStatusCode *int32                         `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Data           *CreateProjectResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Success        *bool                          `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectResponseBody) GoString() string {
	return s.String()
}

func (s *CreateProjectResponseBody) SetRequestId(v string) *CreateProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateProjectResponseBody) SetHttpStatusCode(v int32) *CreateProjectResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateProjectResponseBody) SetData(v *CreateProjectResponseBodyData) *CreateProjectResponseBody {
	s.Data = v
	return s
}

func (s *CreateProjectResponseBody) SetSuccess(v bool) *CreateProjectResponseBody {
	s.Success = &v
	return s
}

type CreateProjectResponseBodyData struct {
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	GmtCreate *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
}

func (s CreateProjectResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateProjectResponseBodyData) SetProjectId(v string) *CreateProjectResponseBodyData {
	s.ProjectId = &v
	return s
}

func (s *CreateProjectResponseBodyData) SetGmtCreate(v int64) *CreateProjectResponseBodyData {
	s.GmtCreate = &v
	return s
}

type CreateProjectResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateProjectResponse) SetBody(v *CreateProjectResponseBody) *CreateProjectResponse {
	s.Body = v
	return s
}

type CreateSlrRoleRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	RoleName        *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
	DurationSeconds *int64  `json:"DurationSeconds,omitempty" xml:"DurationSeconds,omitempty"`
}

func (s CreateSlrRoleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSlrRoleRequest) GoString() string {
	return s.String()
}

func (s *CreateSlrRoleRequest) SetSourceIp(v string) *CreateSlrRoleRequest {
	s.SourceIp = &v
	return s
}

func (s *CreateSlrRoleRequest) SetRoleName(v string) *CreateSlrRoleRequest {
	s.RoleName = &v
	return s
}

func (s *CreateSlrRoleRequest) SetDurationSeconds(v int64) *CreateSlrRoleRequest {
	s.DurationSeconds = &v
	return s
}

type CreateSlrRoleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateSlrRoleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSlrRoleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSlrRoleResponseBody) SetRequestId(v string) *CreateSlrRoleResponseBody {
	s.RequestId = &v
	return s
}

type CreateSlrRoleResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateSlrRoleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateSlrRoleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSlrRoleResponse) GoString() string {
	return s.String()
}

func (s *CreateSlrRoleResponse) SetHeaders(v map[string]*string) *CreateSlrRoleResponse {
	s.Headers = v
	return s
}

func (s *CreateSlrRoleResponse) SetBody(v *CreateSlrRoleResponseBody) *CreateSlrRoleResponse {
	s.Body = v
	return s
}

type DeleteMnsServeRequest struct {
	SourceIp  *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s DeleteMnsServeRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteMnsServeRequest) GoString() string {
	return s.String()
}

func (s *DeleteMnsServeRequest) SetSourceIp(v string) *DeleteMnsServeRequest {
	s.SourceIp = &v
	return s
}

func (s *DeleteMnsServeRequest) SetProjectId(v string) *DeleteMnsServeRequest {
	s.ProjectId = &v
	return s
}

type DeleteMnsServeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteMnsServeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteMnsServeResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMnsServeResponseBody) SetRequestId(v string) *DeleteMnsServeResponseBody {
	s.RequestId = &v
	return s
}

type DeleteMnsServeResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteMnsServeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteMnsServeResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteMnsServeResponse) GoString() string {
	return s.String()
}

func (s *DeleteMnsServeResponse) SetHeaders(v map[string]*string) *DeleteMnsServeResponse {
	s.Headers = v
	return s
}

func (s *DeleteMnsServeResponse) SetBody(v *DeleteMnsServeResponseBody) *DeleteMnsServeResponse {
	s.Body = v
	return s
}

type DeleteUserGroupRequest struct {
	SourceIp  *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	GroupId   *int64  `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s DeleteUserGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteUserGroupRequest) SetSourceIp(v string) *DeleteUserGroupRequest {
	s.SourceIp = &v
	return s
}

func (s *DeleteUserGroupRequest) SetGroupId(v int64) *DeleteUserGroupRequest {
	s.GroupId = &v
	return s
}

func (s *DeleteUserGroupRequest) SetProjectId(v string) *DeleteUserGroupRequest {
	s.ProjectId = &v
	return s
}

type DeleteUserGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteUserGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUserGroupResponseBody) SetRequestId(v string) *DeleteUserGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteUserGroupResponseBody) SetSuccess(v bool) *DeleteUserGroupResponseBody {
	s.Success = &v
	return s
}

type DeleteUserGroupResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteUserGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteUserGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteUserGroupResponse) SetHeaders(v map[string]*string) *DeleteUserGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteUserGroupResponse) SetBody(v *DeleteUserGroupResponseBody) *DeleteUserGroupResponse {
	s.Body = v
	return s
}

type DeleteUserInfoRequest struct {
	SourceIp  *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	UserId    *int64  `json:"UserId,omitempty" xml:"UserId,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s DeleteUserInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserInfoRequest) GoString() string {
	return s.String()
}

func (s *DeleteUserInfoRequest) SetSourceIp(v string) *DeleteUserInfoRequest {
	s.SourceIp = &v
	return s
}

func (s *DeleteUserInfoRequest) SetUserId(v int64) *DeleteUserInfoRequest {
	s.UserId = &v
	return s
}

func (s *DeleteUserInfoRequest) SetProjectId(v string) *DeleteUserInfoRequest {
	s.ProjectId = &v
	return s
}

type DeleteUserInfoResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteUserInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUserInfoResponseBody) SetRequestId(v string) *DeleteUserInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteUserInfoResponseBody) SetSuccess(v bool) *DeleteUserInfoResponseBody {
	s.Success = &v
	return s
}

type DeleteUserInfoResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteUserInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteUserInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserInfoResponse) GoString() string {
	return s.String()
}

func (s *DeleteUserInfoResponse) SetHeaders(v map[string]*string) *DeleteUserInfoResponse {
	s.Headers = v
	return s
}

func (s *DeleteUserInfoResponse) SetBody(v *DeleteUserInfoResponseBody) *DeleteUserInfoResponse {
	s.Body = v
	return s
}

type DescribeAllEndPointRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeAllEndPointRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAllEndPointRequest) GoString() string {
	return s.String()
}

func (s *DescribeAllEndPointRequest) SetSourceIp(v string) *DescribeAllEndPointRequest {
	s.SourceIp = &v
	return s
}

type DescribeAllEndPointResponseBody struct {
	RequestId      *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	HttpStatusCode *int32                                 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Data           []*DescribeAllEndPointResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Success        *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeAllEndPointResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAllEndPointResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAllEndPointResponseBody) SetRequestId(v string) *DescribeAllEndPointResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAllEndPointResponseBody) SetHttpStatusCode(v int32) *DescribeAllEndPointResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeAllEndPointResponseBody) SetData(v []*DescribeAllEndPointResponseBodyData) *DescribeAllEndPointResponseBody {
	s.Data = v
	return s
}

func (s *DescribeAllEndPointResponseBody) SetSuccess(v bool) *DescribeAllEndPointResponseBody {
	s.Success = &v
	return s
}

type DescribeAllEndPointResponseBodyData struct {
	CityName *string `json:"CityName,omitempty" xml:"CityName,omitempty"`
	EndPoint *string `json:"EndPoint,omitempty" xml:"EndPoint,omitempty"`
}

func (s DescribeAllEndPointResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeAllEndPointResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeAllEndPointResponseBodyData) SetCityName(v string) *DescribeAllEndPointResponseBodyData {
	s.CityName = &v
	return s
}

func (s *DescribeAllEndPointResponseBodyData) SetEndPoint(v string) *DescribeAllEndPointResponseBodyData {
	s.EndPoint = &v
	return s
}

type DescribeAllEndPointResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAllEndPointResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAllEndPointResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAllEndPointResponse) GoString() string {
	return s.String()
}

func (s *DescribeAllEndPointResponse) SetHeaders(v map[string]*string) *DescribeAllEndPointResponse {
	s.Headers = v
	return s
}

func (s *DescribeAllEndPointResponse) SetBody(v *DescribeAllEndPointResponseBody) *DescribeAllEndPointResponse {
	s.Body = v
	return s
}

type DescribeBindUserIdListRequest struct {
	SourceIp  *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s DescribeBindUserIdListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBindUserIdListRequest) GoString() string {
	return s.String()
}

func (s *DescribeBindUserIdListRequest) SetSourceIp(v string) *DescribeBindUserIdListRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeBindUserIdListRequest) SetProjectId(v string) *DescribeBindUserIdListRequest {
	s.ProjectId = &v
	return s
}

type DescribeBindUserIdListResponseBody struct {
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      []*DescribeBindUserIdListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
}

func (s DescribeBindUserIdListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeBindUserIdListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBindUserIdListResponseBody) SetRequestId(v string) *DescribeBindUserIdListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBindUserIdListResponseBody) SetData(v []*DescribeBindUserIdListResponseBodyData) *DescribeBindUserIdListResponseBody {
	s.Data = v
	return s
}

type DescribeBindUserIdListResponseBodyData struct {
	CertificateType *string `json:"CertificateType,omitempty" xml:"CertificateType,omitempty"`
	UserName        *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	CertificateNo   *string `json:"CertificateNo,omitempty" xml:"CertificateNo,omitempty"`
	Id              *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DescribeBindUserIdListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeBindUserIdListResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeBindUserIdListResponseBodyData) SetCertificateType(v string) *DescribeBindUserIdListResponseBodyData {
	s.CertificateType = &v
	return s
}

func (s *DescribeBindUserIdListResponseBodyData) SetUserName(v string) *DescribeBindUserIdListResponseBodyData {
	s.UserName = &v
	return s
}

func (s *DescribeBindUserIdListResponseBodyData) SetCertificateNo(v string) *DescribeBindUserIdListResponseBodyData {
	s.CertificateNo = &v
	return s
}

func (s *DescribeBindUserIdListResponseBodyData) SetId(v int64) *DescribeBindUserIdListResponseBodyData {
	s.Id = &v
	return s
}

type DescribeBindUserIdListResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeBindUserIdListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeBindUserIdListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeBindUserIdListResponse) GoString() string {
	return s.String()
}

func (s *DescribeBindUserIdListResponse) SetHeaders(v map[string]*string) *DescribeBindUserIdListResponse {
	s.Headers = v
	return s
}

func (s *DescribeBindUserIdListResponse) SetBody(v *DescribeBindUserIdListResponseBody) *DescribeBindUserIdListResponse {
	s.Body = v
	return s
}

type DescribeCertificateTypeListRequest struct {
	SourceIp  *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s DescribeCertificateTypeListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCertificateTypeListRequest) GoString() string {
	return s.String()
}

func (s *DescribeCertificateTypeListRequest) SetSourceIp(v string) *DescribeCertificateTypeListRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeCertificateTypeListRequest) SetProjectId(v string) *DescribeCertificateTypeListRequest {
	s.ProjectId = &v
	return s
}

type DescribeCertificateTypeListResponseBody struct {
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TypeList  []*string `json:"TypeList,omitempty" xml:"TypeList,omitempty" type:"Repeated"`
}

func (s DescribeCertificateTypeListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCertificateTypeListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCertificateTypeListResponseBody) SetRequestId(v string) *DescribeCertificateTypeListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCertificateTypeListResponseBody) SetTypeList(v []*string) *DescribeCertificateTypeListResponseBody {
	s.TypeList = v
	return s
}

type DescribeCertificateTypeListResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeCertificateTypeListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCertificateTypeListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCertificateTypeListResponse) GoString() string {
	return s.String()
}

func (s *DescribeCertificateTypeListResponse) SetHeaders(v map[string]*string) *DescribeCertificateTypeListResponse {
	s.Headers = v
	return s
}

func (s *DescribeCertificateTypeListResponse) SetBody(v *DescribeCertificateTypeListResponseBody) *DescribeCertificateTypeListResponse {
	s.Body = v
	return s
}

type DescribeDeviceListRequest struct {
	SourceIp    *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Keyword     *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	ProjectId   *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s DescribeDeviceListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeviceListRequest) GoString() string {
	return s.String()
}

func (s *DescribeDeviceListRequest) SetSourceIp(v string) *DescribeDeviceListRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeDeviceListRequest) SetPageSize(v int32) *DescribeDeviceListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDeviceListRequest) SetCurrentPage(v int32) *DescribeDeviceListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeDeviceListRequest) SetKeyword(v string) *DescribeDeviceListRequest {
	s.Keyword = &v
	return s
}

func (s *DescribeDeviceListRequest) SetProjectId(v string) *DescribeDeviceListRequest {
	s.ProjectId = &v
	return s
}

type DescribeDeviceListResponseBody struct {
	TotalCount *int64                                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId  *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DeviceList []*DescribeDeviceListResponseBodyDeviceList `json:"DeviceList,omitempty" xml:"DeviceList,omitempty" type:"Repeated"`
}

func (s DescribeDeviceListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeviceListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDeviceListResponseBody) SetTotalCount(v int64) *DescribeDeviceListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDeviceListResponseBody) SetRequestId(v string) *DescribeDeviceListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDeviceListResponseBody) SetDeviceList(v []*DescribeDeviceListResponseBodyDeviceList) *DescribeDeviceListResponseBody {
	s.DeviceList = v
	return s
}

type DescribeDeviceListResponseBodyDeviceList struct {
	DeviceName   *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	DeviceStatus *string `json:"DeviceStatus,omitempty" xml:"DeviceStatus,omitempty"`
	CategoryName *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
	MemoName     *string `json:"MemoName,omitempty" xml:"MemoName,omitempty"`
	IotId        *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
}

func (s DescribeDeviceListResponseBodyDeviceList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeviceListResponseBodyDeviceList) GoString() string {
	return s.String()
}

func (s *DescribeDeviceListResponseBodyDeviceList) SetDeviceName(v string) *DescribeDeviceListResponseBodyDeviceList {
	s.DeviceName = &v
	return s
}

func (s *DescribeDeviceListResponseBodyDeviceList) SetDeviceStatus(v string) *DescribeDeviceListResponseBodyDeviceList {
	s.DeviceStatus = &v
	return s
}

func (s *DescribeDeviceListResponseBodyDeviceList) SetCategoryName(v string) *DescribeDeviceListResponseBodyDeviceList {
	s.CategoryName = &v
	return s
}

func (s *DescribeDeviceListResponseBodyDeviceList) SetMemoName(v string) *DescribeDeviceListResponseBodyDeviceList {
	s.MemoName = &v
	return s
}

func (s *DescribeDeviceListResponseBodyDeviceList) SetIotId(v string) *DescribeDeviceListResponseBodyDeviceList {
	s.IotId = &v
	return s
}

type DescribeDeviceListResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDeviceListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDeviceListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeviceListResponse) GoString() string {
	return s.String()
}

func (s *DescribeDeviceListResponse) SetHeaders(v map[string]*string) *DescribeDeviceListResponse {
	s.Headers = v
	return s
}

func (s *DescribeDeviceListResponse) SetBody(v *DescribeDeviceListResponseBody) *DescribeDeviceListResponse {
	s.Body = v
	return s
}

type DescribeExcelAnalysisResultRequest struct {
	SourceIp  *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Key       *string `json:"Key,omitempty" xml:"Key,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s DescribeExcelAnalysisResultRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeExcelAnalysisResultRequest) GoString() string {
	return s.String()
}

func (s *DescribeExcelAnalysisResultRequest) SetSourceIp(v string) *DescribeExcelAnalysisResultRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeExcelAnalysisResultRequest) SetKey(v string) *DescribeExcelAnalysisResultRequest {
	s.Key = &v
	return s
}

func (s *DescribeExcelAnalysisResultRequest) SetProjectId(v string) *DescribeExcelAnalysisResultRequest {
	s.ProjectId = &v
	return s
}

type DescribeExcelAnalysisResultResponseBody struct {
	RequestId      *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ExcelResult    *DescribeExcelAnalysisResultResponseBodyExcelResult `json:"ExcelResult,omitempty" xml:"ExcelResult,omitempty" type:"Struct"`
	HttpStatusCode *int32                                              `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Success        *bool                                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeExcelAnalysisResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeExcelAnalysisResultResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeExcelAnalysisResultResponseBody) SetRequestId(v string) *DescribeExcelAnalysisResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeExcelAnalysisResultResponseBody) SetExcelResult(v *DescribeExcelAnalysisResultResponseBodyExcelResult) *DescribeExcelAnalysisResultResponseBody {
	s.ExcelResult = v
	return s
}

func (s *DescribeExcelAnalysisResultResponseBody) SetHttpStatusCode(v int32) *DescribeExcelAnalysisResultResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeExcelAnalysisResultResponseBody) SetSuccess(v bool) *DescribeExcelAnalysisResultResponseBody {
	s.Success = &v
	return s
}

type DescribeExcelAnalysisResultResponseBodyExcelResult struct {
	InsertCount *int32    `json:"InsertCount,omitempty" xml:"InsertCount,omitempty"`
	UpdateCount *int32    `json:"UpdateCount,omitempty" xml:"UpdateCount,omitempty"`
	ErrorCount  *int32    `json:"ErrorCount,omitempty" xml:"ErrorCount,omitempty"`
	ErrorLine   []*string `json:"ErrorLine,omitempty" xml:"ErrorLine,omitempty" type:"Repeated"`
	Total       *int32    `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeExcelAnalysisResultResponseBodyExcelResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeExcelAnalysisResultResponseBodyExcelResult) GoString() string {
	return s.String()
}

func (s *DescribeExcelAnalysisResultResponseBodyExcelResult) SetInsertCount(v int32) *DescribeExcelAnalysisResultResponseBodyExcelResult {
	s.InsertCount = &v
	return s
}

func (s *DescribeExcelAnalysisResultResponseBodyExcelResult) SetUpdateCount(v int32) *DescribeExcelAnalysisResultResponseBodyExcelResult {
	s.UpdateCount = &v
	return s
}

func (s *DescribeExcelAnalysisResultResponseBodyExcelResult) SetErrorCount(v int32) *DescribeExcelAnalysisResultResponseBodyExcelResult {
	s.ErrorCount = &v
	return s
}

func (s *DescribeExcelAnalysisResultResponseBodyExcelResult) SetErrorLine(v []*string) *DescribeExcelAnalysisResultResponseBodyExcelResult {
	s.ErrorLine = v
	return s
}

func (s *DescribeExcelAnalysisResultResponseBodyExcelResult) SetTotal(v int32) *DescribeExcelAnalysisResultResponseBodyExcelResult {
	s.Total = &v
	return s
}

type DescribeExcelAnalysisResultResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeExcelAnalysisResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeExcelAnalysisResultResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeExcelAnalysisResultResponse) GoString() string {
	return s.String()
}

func (s *DescribeExcelAnalysisResultResponse) SetHeaders(v map[string]*string) *DescribeExcelAnalysisResultResponse {
	s.Headers = v
	return s
}

func (s *DescribeExcelAnalysisResultResponse) SetBody(v *DescribeExcelAnalysisResultResponseBody) *DescribeExcelAnalysisResultResponse {
	s.Body = v
	return s
}

type DescribeIdentifyRecordListRequest struct {
	SourceIp      *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	GroupName     *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	DeviceName    *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	StartTime     *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime       *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	CurrentPage   *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	UserName      *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	CertificateNo *string `json:"CertificateNo,omitempty" xml:"CertificateNo,omitempty"`
	PhoneNo       *string `json:"PhoneNo,omitempty" xml:"PhoneNo,omitempty"`
	ProjectId     *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s DescribeIdentifyRecordListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeIdentifyRecordListRequest) GoString() string {
	return s.String()
}

func (s *DescribeIdentifyRecordListRequest) SetSourceIp(v string) *DescribeIdentifyRecordListRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeIdentifyRecordListRequest) SetGroupName(v string) *DescribeIdentifyRecordListRequest {
	s.GroupName = &v
	return s
}

func (s *DescribeIdentifyRecordListRequest) SetDeviceName(v string) *DescribeIdentifyRecordListRequest {
	s.DeviceName = &v
	return s
}

func (s *DescribeIdentifyRecordListRequest) SetPageSize(v int32) *DescribeIdentifyRecordListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeIdentifyRecordListRequest) SetStartTime(v int64) *DescribeIdentifyRecordListRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeIdentifyRecordListRequest) SetEndTime(v int64) *DescribeIdentifyRecordListRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeIdentifyRecordListRequest) SetCurrentPage(v int32) *DescribeIdentifyRecordListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeIdentifyRecordListRequest) SetUserName(v string) *DescribeIdentifyRecordListRequest {
	s.UserName = &v
	return s
}

func (s *DescribeIdentifyRecordListRequest) SetCertificateNo(v string) *DescribeIdentifyRecordListRequest {
	s.CertificateNo = &v
	return s
}

func (s *DescribeIdentifyRecordListRequest) SetPhoneNo(v string) *DescribeIdentifyRecordListRequest {
	s.PhoneNo = &v
	return s
}

func (s *DescribeIdentifyRecordListRequest) SetProjectId(v string) *DescribeIdentifyRecordListRequest {
	s.ProjectId = &v
	return s
}

type DescribeIdentifyRecordListResponseBody struct {
	TotalCount     *int64                                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId      *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RecordList     []*DescribeIdentifyRecordListResponseBodyRecordList `json:"RecordList,omitempty" xml:"RecordList,omitempty" type:"Repeated"`
	HttpStatusCode *int32                                              `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Success        *bool                                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeIdentifyRecordListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeIdentifyRecordListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeIdentifyRecordListResponseBody) SetTotalCount(v int64) *DescribeIdentifyRecordListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeIdentifyRecordListResponseBody) SetRequestId(v string) *DescribeIdentifyRecordListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeIdentifyRecordListResponseBody) SetRecordList(v []*DescribeIdentifyRecordListResponseBodyRecordList) *DescribeIdentifyRecordListResponseBody {
	s.RecordList = v
	return s
}

func (s *DescribeIdentifyRecordListResponseBody) SetHttpStatusCode(v int32) *DescribeIdentifyRecordListResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeIdentifyRecordListResponseBody) SetSuccess(v bool) *DescribeIdentifyRecordListResponseBody {
	s.Success = &v
	return s
}

type DescribeIdentifyRecordListResponseBodyRecordList struct {
	CapturedImage    *string `json:"CapturedImage,omitempty" xml:"CapturedImage,omitempty"`
	DeviceName       *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	GroupName        *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	UserId           *int32  `json:"UserId,omitempty" xml:"UserId,omitempty"`
	IdentifyingImage *string `json:"IdentifyingImage,omitempty" xml:"IdentifyingImage,omitempty"`
	GmtCreate        *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	UserName         *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	IotId            *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
}

func (s DescribeIdentifyRecordListResponseBodyRecordList) String() string {
	return tea.Prettify(s)
}

func (s DescribeIdentifyRecordListResponseBodyRecordList) GoString() string {
	return s.String()
}

func (s *DescribeIdentifyRecordListResponseBodyRecordList) SetCapturedImage(v string) *DescribeIdentifyRecordListResponseBodyRecordList {
	s.CapturedImage = &v
	return s
}

func (s *DescribeIdentifyRecordListResponseBodyRecordList) SetDeviceName(v string) *DescribeIdentifyRecordListResponseBodyRecordList {
	s.DeviceName = &v
	return s
}

func (s *DescribeIdentifyRecordListResponseBodyRecordList) SetGroupName(v string) *DescribeIdentifyRecordListResponseBodyRecordList {
	s.GroupName = &v
	return s
}

func (s *DescribeIdentifyRecordListResponseBodyRecordList) SetUserId(v int32) *DescribeIdentifyRecordListResponseBodyRecordList {
	s.UserId = &v
	return s
}

func (s *DescribeIdentifyRecordListResponseBodyRecordList) SetIdentifyingImage(v string) *DescribeIdentifyRecordListResponseBodyRecordList {
	s.IdentifyingImage = &v
	return s
}

func (s *DescribeIdentifyRecordListResponseBodyRecordList) SetGmtCreate(v int64) *DescribeIdentifyRecordListResponseBodyRecordList {
	s.GmtCreate = &v
	return s
}

func (s *DescribeIdentifyRecordListResponseBodyRecordList) SetUserName(v string) *DescribeIdentifyRecordListResponseBodyRecordList {
	s.UserName = &v
	return s
}

func (s *DescribeIdentifyRecordListResponseBodyRecordList) SetIotId(v string) *DescribeIdentifyRecordListResponseBodyRecordList {
	s.IotId = &v
	return s
}

type DescribeIdentifyRecordListResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeIdentifyRecordListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeIdentifyRecordListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeIdentifyRecordListResponse) GoString() string {
	return s.String()
}

func (s *DescribeIdentifyRecordListResponse) SetHeaders(v map[string]*string) *DescribeIdentifyRecordListResponse {
	s.Headers = v
	return s
}

func (s *DescribeIdentifyRecordListResponse) SetBody(v *DescribeIdentifyRecordListResponseBody) *DescribeIdentifyRecordListResponse {
	s.Body = v
	return s
}

type DescribeMnsOauthRequest struct {
	SourceIp  *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s DescribeMnsOauthRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeMnsOauthRequest) GoString() string {
	return s.String()
}

func (s *DescribeMnsOauthRequest) SetSourceIp(v string) *DescribeMnsOauthRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeMnsOauthRequest) SetProjectId(v string) *DescribeMnsOauthRequest {
	s.ProjectId = &v
	return s
}

type DescribeMnsOauthResponseBody struct {
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *DescribeMnsOauthResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s DescribeMnsOauthResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeMnsOauthResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMnsOauthResponseBody) SetRequestId(v string) *DescribeMnsOauthResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMnsOauthResponseBody) SetData(v *DescribeMnsOauthResponseBodyData) *DescribeMnsOauthResponseBody {
	s.Data = v
	return s
}

type DescribeMnsOauthResponseBodyData struct {
	TopicList      []*string `json:"TopicList,omitempty" xml:"TopicList,omitempty" type:"Repeated"`
	TopicName      *string   `json:"TopicName,omitempty" xml:"TopicName,omitempty"`
	MnsAuthorized  *bool     `json:"MnsAuthorized,omitempty" xml:"MnsAuthorized,omitempty"`
	EndPoint       *string   `json:"EndPoint,omitempty" xml:"EndPoint,omitempty"`
	MnsServe       *bool     `json:"MnsServe,omitempty" xml:"MnsServe,omitempty"`
	OpenMnsService *bool     `json:"OpenMnsService,omitempty" xml:"OpenMnsService,omitempty"`
}

func (s DescribeMnsOauthResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeMnsOauthResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeMnsOauthResponseBodyData) SetTopicList(v []*string) *DescribeMnsOauthResponseBodyData {
	s.TopicList = v
	return s
}

func (s *DescribeMnsOauthResponseBodyData) SetTopicName(v string) *DescribeMnsOauthResponseBodyData {
	s.TopicName = &v
	return s
}

func (s *DescribeMnsOauthResponseBodyData) SetMnsAuthorized(v bool) *DescribeMnsOauthResponseBodyData {
	s.MnsAuthorized = &v
	return s
}

func (s *DescribeMnsOauthResponseBodyData) SetEndPoint(v string) *DescribeMnsOauthResponseBodyData {
	s.EndPoint = &v
	return s
}

func (s *DescribeMnsOauthResponseBodyData) SetMnsServe(v bool) *DescribeMnsOauthResponseBodyData {
	s.MnsServe = &v
	return s
}

func (s *DescribeMnsOauthResponseBodyData) SetOpenMnsService(v bool) *DescribeMnsOauthResponseBodyData {
	s.OpenMnsService = &v
	return s
}

type DescribeMnsOauthResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeMnsOauthResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeMnsOauthResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeMnsOauthResponse) GoString() string {
	return s.String()
}

func (s *DescribeMnsOauthResponse) SetHeaders(v map[string]*string) *DescribeMnsOauthResponse {
	s.Headers = v
	return s
}

func (s *DescribeMnsOauthResponse) SetBody(v *DescribeMnsOauthResponseBody) *DescribeMnsOauthResponse {
	s.Body = v
	return s
}

type DescribeOssOauthRequest struct {
	SourceIp  *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s DescribeOssOauthRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeOssOauthRequest) GoString() string {
	return s.String()
}

func (s *DescribeOssOauthRequest) SetSourceIp(v string) *DescribeOssOauthRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeOssOauthRequest) SetProjectId(v string) *DescribeOssOauthRequest {
	s.ProjectId = &v
	return s
}

type DescribeOssOauthResponseBody struct {
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *DescribeOssOauthResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s DescribeOssOauthResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeOssOauthResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOssOauthResponseBody) SetRequestId(v string) *DescribeOssOauthResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeOssOauthResponseBody) SetData(v *DescribeOssOauthResponseBodyData) *DescribeOssOauthResponseBody {
	s.Data = v
	return s
}

type DescribeOssOauthResponseBodyData struct {
	OssServe       *bool   `json:"OssServe,omitempty" xml:"OssServe,omitempty"`
	BucketName     *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	OpenOssService *bool   `json:"OpenOssService,omitempty" xml:"OpenOssService,omitempty"`
	OssAuthorized  *bool   `json:"OssAuthorized,omitempty" xml:"OssAuthorized,omitempty"`
	OpenOssTime    *int64  `json:"OpenOssTime,omitempty" xml:"OpenOssTime,omitempty"`
}

func (s DescribeOssOauthResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeOssOauthResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeOssOauthResponseBodyData) SetOssServe(v bool) *DescribeOssOauthResponseBodyData {
	s.OssServe = &v
	return s
}

func (s *DescribeOssOauthResponseBodyData) SetBucketName(v string) *DescribeOssOauthResponseBodyData {
	s.BucketName = &v
	return s
}

func (s *DescribeOssOauthResponseBodyData) SetOpenOssService(v bool) *DescribeOssOauthResponseBodyData {
	s.OpenOssService = &v
	return s
}

func (s *DescribeOssOauthResponseBodyData) SetOssAuthorized(v bool) *DescribeOssOauthResponseBodyData {
	s.OssAuthorized = &v
	return s
}

func (s *DescribeOssOauthResponseBodyData) SetOpenOssTime(v int64) *DescribeOssOauthResponseBodyData {
	s.OpenOssTime = &v
	return s
}

type DescribeOssOauthResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeOssOauthResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeOssOauthResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeOssOauthResponse) GoString() string {
	return s.String()
}

func (s *DescribeOssOauthResponse) SetHeaders(v map[string]*string) *DescribeOssOauthResponse {
	s.Headers = v
	return s
}

func (s *DescribeOssOauthResponse) SetBody(v *DescribeOssOauthResponseBody) *DescribeOssOauthResponse {
	s.Body = v
	return s
}

type DescribeSignedUrlRequest struct {
	SourceIp  *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Key       *string `json:"Key,omitempty" xml:"Key,omitempty"`
}

func (s DescribeSignedUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSignedUrlRequest) GoString() string {
	return s.String()
}

func (s *DescribeSignedUrlRequest) SetSourceIp(v string) *DescribeSignedUrlRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeSignedUrlRequest) SetProjectId(v string) *DescribeSignedUrlRequest {
	s.ProjectId = &v
	return s
}

func (s *DescribeSignedUrlRequest) SetKey(v string) *DescribeSignedUrlRequest {
	s.Key = &v
	return s
}

type DescribeSignedUrlResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	ImgUrl    *string `json:"ImgUrl,omitempty" xml:"ImgUrl,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeSignedUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSignedUrlResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSignedUrlResponseBody) SetRequestId(v string) *DescribeSignedUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSignedUrlResponseBody) SetCode(v int32) *DescribeSignedUrlResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeSignedUrlResponseBody) SetImgUrl(v string) *DescribeSignedUrlResponseBody {
	s.ImgUrl = &v
	return s
}

func (s *DescribeSignedUrlResponseBody) SetSuccess(v bool) *DescribeSignedUrlResponseBody {
	s.Success = &v
	return s
}

type DescribeSignedUrlResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSignedUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSignedUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSignedUrlResponse) GoString() string {
	return s.String()
}

func (s *DescribeSignedUrlResponse) SetHeaders(v map[string]*string) *DescribeSignedUrlResponse {
	s.Headers = v
	return s
}

func (s *DescribeSignedUrlResponse) SetBody(v *DescribeSignedUrlResponseBody) *DescribeSignedUrlResponse {
	s.Body = v
	return s
}

type DescribeTopicRequest struct {
	SourceIp  *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	EndPoint  *string `json:"EndPoint,omitempty" xml:"EndPoint,omitempty"`
}

func (s DescribeTopicRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTopicRequest) GoString() string {
	return s.String()
}

func (s *DescribeTopicRequest) SetSourceIp(v string) *DescribeTopicRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeTopicRequest) SetProjectId(v string) *DescribeTopicRequest {
	s.ProjectId = &v
	return s
}

func (s *DescribeTopicRequest) SetEndPoint(v string) *DescribeTopicRequest {
	s.EndPoint = &v
	return s
}

type DescribeTopicResponseBody struct {
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      []*string `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
}

func (s DescribeTopicResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTopicResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTopicResponseBody) SetRequestId(v string) *DescribeTopicResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTopicResponseBody) SetData(v []*string) *DescribeTopicResponseBody {
	s.Data = v
	return s
}

type DescribeTopicResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeTopicResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeTopicResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeTopicResponse) GoString() string {
	return s.String()
}

func (s *DescribeTopicResponse) SetHeaders(v map[string]*string) *DescribeTopicResponse {
	s.Headers = v
	return s
}

func (s *DescribeTopicResponse) SetBody(v *DescribeTopicResponseBody) *DescribeTopicResponse {
	s.Body = v
	return s
}

type DescribeUploadPreSignRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeUploadPreSignRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeUploadPreSignRequest) GoString() string {
	return s.String()
}

func (s *DescribeUploadPreSignRequest) SetSourceIp(v string) *DescribeUploadPreSignRequest {
	s.SourceIp = &v
	return s
}

type DescribeUploadPreSignResponseBody struct {
	Policy    *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	Expire    *string `json:"Expire,omitempty" xml:"Expire,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	AccessId  *string `json:"AccessId,omitempty" xml:"AccessId,omitempty"`
	Signature *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
	Host      *string `json:"Host,omitempty" xml:"Host,omitempty"`
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Key       *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeUploadPreSignResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeUploadPreSignResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUploadPreSignResponseBody) SetPolicy(v string) *DescribeUploadPreSignResponseBody {
	s.Policy = &v
	return s
}

func (s *DescribeUploadPreSignResponseBody) SetExpire(v string) *DescribeUploadPreSignResponseBody {
	s.Expire = &v
	return s
}

func (s *DescribeUploadPreSignResponseBody) SetRequestId(v string) *DescribeUploadPreSignResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUploadPreSignResponseBody) SetAccessId(v string) *DescribeUploadPreSignResponseBody {
	s.AccessId = &v
	return s
}

func (s *DescribeUploadPreSignResponseBody) SetSignature(v string) *DescribeUploadPreSignResponseBody {
	s.Signature = &v
	return s
}

func (s *DescribeUploadPreSignResponseBody) SetHost(v string) *DescribeUploadPreSignResponseBody {
	s.Host = &v
	return s
}

func (s *DescribeUploadPreSignResponseBody) SetCode(v int32) *DescribeUploadPreSignResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeUploadPreSignResponseBody) SetKey(v string) *DescribeUploadPreSignResponseBody {
	s.Key = &v
	return s
}

func (s *DescribeUploadPreSignResponseBody) SetSuccess(v bool) *DescribeUploadPreSignResponseBody {
	s.Success = &v
	return s
}

type DescribeUploadPreSignResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeUploadPreSignResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeUploadPreSignResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeUploadPreSignResponse) GoString() string {
	return s.String()
}

func (s *DescribeUploadPreSignResponse) SetHeaders(v map[string]*string) *DescribeUploadPreSignResponse {
	s.Headers = v
	return s
}

func (s *DescribeUploadPreSignResponse) SetBody(v *DescribeUploadPreSignResponseBody) *DescribeUploadPreSignResponse {
	s.Body = v
	return s
}

type DescribeUserGroupListRequest struct {
	SourceIp  *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s DescribeUserGroupListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserGroupListRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserGroupListRequest) SetSourceIp(v string) *DescribeUserGroupListRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeUserGroupListRequest) SetProjectId(v string) *DescribeUserGroupListRequest {
	s.ProjectId = &v
	return s
}

type DescribeUserGroupListResponseBody struct {
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      []*DescribeUserGroupListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Success   *bool                                    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeUserGroupListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserGroupListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserGroupListResponseBody) SetRequestId(v string) *DescribeUserGroupListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUserGroupListResponseBody) SetData(v []*DescribeUserGroupListResponseBodyData) *DescribeUserGroupListResponseBody {
	s.Data = v
	return s
}

func (s *DescribeUserGroupListResponseBody) SetSuccess(v bool) *DescribeUserGroupListResponseBody {
	s.Success = &v
	return s
}

type DescribeUserGroupListResponseBodyData struct {
	GroupName      *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	GroupUserCount *int32  `json:"GroupUserCount,omitempty" xml:"GroupUserCount,omitempty"`
	Id             *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DescribeUserGroupListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserGroupListResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeUserGroupListResponseBodyData) SetGroupName(v string) *DescribeUserGroupListResponseBodyData {
	s.GroupName = &v
	return s
}

func (s *DescribeUserGroupListResponseBodyData) SetGroupUserCount(v int32) *DescribeUserGroupListResponseBodyData {
	s.GroupUserCount = &v
	return s
}

func (s *DescribeUserGroupListResponseBodyData) SetId(v int64) *DescribeUserGroupListResponseBodyData {
	s.Id = &v
	return s
}

type DescribeUserGroupListResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeUserGroupListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeUserGroupListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserGroupListResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserGroupListResponse) SetHeaders(v map[string]*string) *DescribeUserGroupListResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserGroupListResponse) SetBody(v *DescribeUserGroupListResponseBody) *DescribeUserGroupListResponse {
	s.Body = v
	return s
}

type DescribeUserInfoListRequest struct {
	SourceIp      *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	UserName      *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	CurrentPage   *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	CertificateNo *string `json:"CertificateNo,omitempty" xml:"CertificateNo,omitempty"`
	PhoneNo       *string `json:"PhoneNo,omitempty" xml:"PhoneNo,omitempty"`
	UserGroupId   *int64  `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
	ProjectId     *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s DescribeUserInfoListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserInfoListRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserInfoListRequest) SetSourceIp(v string) *DescribeUserInfoListRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeUserInfoListRequest) SetPageSize(v int32) *DescribeUserInfoListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeUserInfoListRequest) SetUserName(v string) *DescribeUserInfoListRequest {
	s.UserName = &v
	return s
}

func (s *DescribeUserInfoListRequest) SetCurrentPage(v int32) *DescribeUserInfoListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeUserInfoListRequest) SetCertificateNo(v string) *DescribeUserInfoListRequest {
	s.CertificateNo = &v
	return s
}

func (s *DescribeUserInfoListRequest) SetPhoneNo(v string) *DescribeUserInfoListRequest {
	s.PhoneNo = &v
	return s
}

func (s *DescribeUserInfoListRequest) SetUserGroupId(v int64) *DescribeUserInfoListRequest {
	s.UserGroupId = &v
	return s
}

func (s *DescribeUserInfoListRequest) SetProjectId(v string) *DescribeUserInfoListRequest {
	s.ProjectId = &v
	return s
}

type DescribeUserInfoListResponseBody struct {
	TotalCount   *int64                                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId    *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UserInfoList []*DescribeUserInfoListResponseBodyUserInfoList `json:"UserInfoList,omitempty" xml:"UserInfoList,omitempty" type:"Repeated"`
}

func (s DescribeUserInfoListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserInfoListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserInfoListResponseBody) SetTotalCount(v int64) *DescribeUserInfoListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeUserInfoListResponseBody) SetRequestId(v string) *DescribeUserInfoListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUserInfoListResponseBody) SetUserInfoList(v []*DescribeUserInfoListResponseBodyUserInfoList) *DescribeUserInfoListResponseBody {
	s.UserInfoList = v
	return s
}

type DescribeUserInfoListResponseBodyUserInfoList struct {
	Sex              *int32  `json:"Sex,omitempty" xml:"Sex,omitempty"`
	CertificateType  *string `json:"CertificateType,omitempty" xml:"CertificateType,omitempty"`
	Birthday         *int64  `json:"Birthday,omitempty" xml:"Birthday,omitempty"`
	PhoneNo          *string `json:"PhoneNo,omitempty" xml:"PhoneNo,omitempty"`
	GroupName        *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	UserGroupId      *int64  `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
	IdentifyingImage *string `json:"IdentifyingImage,omitempty" xml:"IdentifyingImage,omitempty"`
	UserName         *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	CertificateNo    *string `json:"CertificateNo,omitempty" xml:"CertificateNo,omitempty"`
	Id               *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DescribeUserInfoListResponseBodyUserInfoList) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserInfoListResponseBodyUserInfoList) GoString() string {
	return s.String()
}

func (s *DescribeUserInfoListResponseBodyUserInfoList) SetSex(v int32) *DescribeUserInfoListResponseBodyUserInfoList {
	s.Sex = &v
	return s
}

func (s *DescribeUserInfoListResponseBodyUserInfoList) SetCertificateType(v string) *DescribeUserInfoListResponseBodyUserInfoList {
	s.CertificateType = &v
	return s
}

func (s *DescribeUserInfoListResponseBodyUserInfoList) SetBirthday(v int64) *DescribeUserInfoListResponseBodyUserInfoList {
	s.Birthday = &v
	return s
}

func (s *DescribeUserInfoListResponseBodyUserInfoList) SetPhoneNo(v string) *DescribeUserInfoListResponseBodyUserInfoList {
	s.PhoneNo = &v
	return s
}

func (s *DescribeUserInfoListResponseBodyUserInfoList) SetGroupName(v string) *DescribeUserInfoListResponseBodyUserInfoList {
	s.GroupName = &v
	return s
}

func (s *DescribeUserInfoListResponseBodyUserInfoList) SetUserGroupId(v int64) *DescribeUserInfoListResponseBodyUserInfoList {
	s.UserGroupId = &v
	return s
}

func (s *DescribeUserInfoListResponseBodyUserInfoList) SetIdentifyingImage(v string) *DescribeUserInfoListResponseBodyUserInfoList {
	s.IdentifyingImage = &v
	return s
}

func (s *DescribeUserInfoListResponseBodyUserInfoList) SetUserName(v string) *DescribeUserInfoListResponseBodyUserInfoList {
	s.UserName = &v
	return s
}

func (s *DescribeUserInfoListResponseBodyUserInfoList) SetCertificateNo(v string) *DescribeUserInfoListResponseBodyUserInfoList {
	s.CertificateNo = &v
	return s
}

func (s *DescribeUserInfoListResponseBodyUserInfoList) SetId(v int64) *DescribeUserInfoListResponseBodyUserInfoList {
	s.Id = &v
	return s
}

type DescribeUserInfoListResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeUserInfoListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeUserInfoListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserInfoListResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserInfoListResponse) SetHeaders(v map[string]*string) *DescribeUserInfoListResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserInfoListResponse) SetBody(v *DescribeUserInfoListResponseBody) *DescribeUserInfoListResponse {
	s.Body = v
	return s
}

type GetAccountProjectRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s GetAccountProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAccountProjectRequest) GoString() string {
	return s.String()
}

func (s *GetAccountProjectRequest) SetSourceIp(v string) *GetAccountProjectRequest {
	s.SourceIp = &v
	return s
}

type GetAccountProjectResponseBody struct {
	RequestId      *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	HttpStatusCode *int32                               `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	List           []*GetAccountProjectResponseBodyList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	Success        *bool                                `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAccountProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAccountProjectResponseBody) GoString() string {
	return s.String()
}

func (s *GetAccountProjectResponseBody) SetRequestId(v string) *GetAccountProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAccountProjectResponseBody) SetHttpStatusCode(v int32) *GetAccountProjectResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetAccountProjectResponseBody) SetList(v []*GetAccountProjectResponseBodyList) *GetAccountProjectResponseBody {
	s.List = v
	return s
}

func (s *GetAccountProjectResponseBody) SetSuccess(v bool) *GetAccountProjectResponseBody {
	s.Success = &v
	return s
}

type GetAccountProjectResponseBodyList struct {
	ProjectId   *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s GetAccountProjectResponseBodyList) String() string {
	return tea.Prettify(s)
}

func (s GetAccountProjectResponseBodyList) GoString() string {
	return s.String()
}

func (s *GetAccountProjectResponseBodyList) SetProjectId(v string) *GetAccountProjectResponseBodyList {
	s.ProjectId = &v
	return s
}

func (s *GetAccountProjectResponseBodyList) SetProjectName(v string) *GetAccountProjectResponseBodyList {
	s.ProjectName = &v
	return s
}

type GetAccountProjectResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetAccountProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAccountProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAccountProjectResponse) GoString() string {
	return s.String()
}

func (s *GetAccountProjectResponse) SetHeaders(v map[string]*string) *GetAccountProjectResponse {
	s.Headers = v
	return s
}

func (s *GetAccountProjectResponse) SetBody(v *GetAccountProjectResponseBody) *GetAccountProjectResponse {
	s.Body = v
	return s
}

type SaveMnsServeRequest struct {
	SourceIp  *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	TopicName *string `json:"TopicName,omitempty" xml:"TopicName,omitempty"`
	EndPoint  *string `json:"EndPoint,omitempty" xml:"EndPoint,omitempty"`
}

func (s SaveMnsServeRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveMnsServeRequest) GoString() string {
	return s.String()
}

func (s *SaveMnsServeRequest) SetSourceIp(v string) *SaveMnsServeRequest {
	s.SourceIp = &v
	return s
}

func (s *SaveMnsServeRequest) SetProjectId(v string) *SaveMnsServeRequest {
	s.ProjectId = &v
	return s
}

func (s *SaveMnsServeRequest) SetTopicName(v string) *SaveMnsServeRequest {
	s.TopicName = &v
	return s
}

func (s *SaveMnsServeRequest) SetEndPoint(v string) *SaveMnsServeRequest {
	s.EndPoint = &v
	return s
}

type SaveMnsServeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SaveMnsServeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveMnsServeResponseBody) GoString() string {
	return s.String()
}

func (s *SaveMnsServeResponseBody) SetRequestId(v string) *SaveMnsServeResponseBody {
	s.RequestId = &v
	return s
}

type SaveMnsServeResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SaveMnsServeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SaveMnsServeResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveMnsServeResponse) GoString() string {
	return s.String()
}

func (s *SaveMnsServeResponse) SetHeaders(v map[string]*string) *SaveMnsServeResponse {
	s.Headers = v
	return s
}

func (s *SaveMnsServeResponse) SetBody(v *SaveMnsServeResponseBody) *SaveMnsServeResponse {
	s.Body = v
	return s
}

type SaveOssServeRequest struct {
	SourceIp  *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s SaveOssServeRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveOssServeRequest) GoString() string {
	return s.String()
}

func (s *SaveOssServeRequest) SetSourceIp(v string) *SaveOssServeRequest {
	s.SourceIp = &v
	return s
}

func (s *SaveOssServeRequest) SetProjectId(v string) *SaveOssServeRequest {
	s.ProjectId = &v
	return s
}

type SaveOssServeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SaveOssServeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveOssServeResponseBody) GoString() string {
	return s.String()
}

func (s *SaveOssServeResponseBody) SetRequestId(v string) *SaveOssServeResponseBody {
	s.RequestId = &v
	return s
}

type SaveOssServeResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SaveOssServeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SaveOssServeResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveOssServeResponse) GoString() string {
	return s.String()
}

func (s *SaveOssServeResponse) SetHeaders(v map[string]*string) *SaveOssServeResponse {
	s.Headers = v
	return s
}

func (s *SaveOssServeResponse) SetBody(v *SaveOssServeResponseBody) *SaveOssServeResponse {
	s.Body = v
	return s
}

type SaveUserGroupRequest struct {
	SourceIp  *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s SaveUserGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveUserGroupRequest) GoString() string {
	return s.String()
}

func (s *SaveUserGroupRequest) SetSourceIp(v string) *SaveUserGroupRequest {
	s.SourceIp = &v
	return s
}

func (s *SaveUserGroupRequest) SetGroupName(v string) *SaveUserGroupRequest {
	s.GroupName = &v
	return s
}

func (s *SaveUserGroupRequest) SetProjectId(v string) *SaveUserGroupRequest {
	s.ProjectId = &v
	return s
}

type SaveUserGroupResponseBody struct {
	RequestId      *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	HttpStatusCode *int32                         `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Data           *SaveUserGroupResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Success        *bool                          `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SaveUserGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveUserGroupResponseBody) GoString() string {
	return s.String()
}

func (s *SaveUserGroupResponseBody) SetRequestId(v string) *SaveUserGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveUserGroupResponseBody) SetHttpStatusCode(v int32) *SaveUserGroupResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SaveUserGroupResponseBody) SetData(v *SaveUserGroupResponseBodyData) *SaveUserGroupResponseBody {
	s.Data = v
	return s
}

func (s *SaveUserGroupResponseBody) SetSuccess(v bool) *SaveUserGroupResponseBody {
	s.Success = &v
	return s
}

type SaveUserGroupResponseBodyData struct {
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	Id        *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s SaveUserGroupResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SaveUserGroupResponseBodyData) GoString() string {
	return s.String()
}

func (s *SaveUserGroupResponseBodyData) SetGmtCreate(v int64) *SaveUserGroupResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *SaveUserGroupResponseBodyData) SetId(v int64) *SaveUserGroupResponseBodyData {
	s.Id = &v
	return s
}

type SaveUserGroupResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SaveUserGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SaveUserGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveUserGroupResponse) GoString() string {
	return s.String()
}

func (s *SaveUserGroupResponse) SetHeaders(v map[string]*string) *SaveUserGroupResponse {
	s.Headers = v
	return s
}

func (s *SaveUserGroupResponse) SetBody(v *SaveUserGroupResponseBody) *SaveUserGroupResponse {
	s.Body = v
	return s
}

type SaveUserInfoRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	UserGroupId     *int64  `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
	Birthday        *int64  `json:"Birthday,omitempty" xml:"Birthday,omitempty"`
	ImageBase64     *string `json:"ImageBase64,omitempty" xml:"ImageBase64,omitempty"`
	GroupName       *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	ImageUrl        *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	Sex             *int32  `json:"Sex,omitempty" xml:"Sex,omitempty"`
	UserName        *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	CertificateNo   *string `json:"CertificateNo,omitempty" xml:"CertificateNo,omitempty"`
	PhoneNo         *string `json:"PhoneNo,omitempty" xml:"PhoneNo,omitempty"`
	CertificateType *string `json:"CertificateType,omitempty" xml:"CertificateType,omitempty"`
	ProjectId       *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s SaveUserInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveUserInfoRequest) GoString() string {
	return s.String()
}

func (s *SaveUserInfoRequest) SetSourceIp(v string) *SaveUserInfoRequest {
	s.SourceIp = &v
	return s
}

func (s *SaveUserInfoRequest) SetUserGroupId(v int64) *SaveUserInfoRequest {
	s.UserGroupId = &v
	return s
}

func (s *SaveUserInfoRequest) SetBirthday(v int64) *SaveUserInfoRequest {
	s.Birthday = &v
	return s
}

func (s *SaveUserInfoRequest) SetImageBase64(v string) *SaveUserInfoRequest {
	s.ImageBase64 = &v
	return s
}

func (s *SaveUserInfoRequest) SetGroupName(v string) *SaveUserInfoRequest {
	s.GroupName = &v
	return s
}

func (s *SaveUserInfoRequest) SetImageUrl(v string) *SaveUserInfoRequest {
	s.ImageUrl = &v
	return s
}

func (s *SaveUserInfoRequest) SetSex(v int32) *SaveUserInfoRequest {
	s.Sex = &v
	return s
}

func (s *SaveUserInfoRequest) SetUserName(v string) *SaveUserInfoRequest {
	s.UserName = &v
	return s
}

func (s *SaveUserInfoRequest) SetCertificateNo(v string) *SaveUserInfoRequest {
	s.CertificateNo = &v
	return s
}

func (s *SaveUserInfoRequest) SetPhoneNo(v string) *SaveUserInfoRequest {
	s.PhoneNo = &v
	return s
}

func (s *SaveUserInfoRequest) SetCertificateType(v string) *SaveUserInfoRequest {
	s.CertificateType = &v
	return s
}

func (s *SaveUserInfoRequest) SetProjectId(v string) *SaveUserInfoRequest {
	s.ProjectId = &v
	return s
}

type SaveUserInfoResponseBody struct {
	RequestId      *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	HttpStatusCode *int32                        `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Data           *SaveUserInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Success        *bool                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SaveUserInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveUserInfoResponseBody) GoString() string {
	return s.String()
}

func (s *SaveUserInfoResponseBody) SetRequestId(v string) *SaveUserInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveUserInfoResponseBody) SetHttpStatusCode(v int32) *SaveUserInfoResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SaveUserInfoResponseBody) SetData(v *SaveUserInfoResponseBodyData) *SaveUserInfoResponseBody {
	s.Data = v
	return s
}

func (s *SaveUserInfoResponseBody) SetSuccess(v bool) *SaveUserInfoResponseBody {
	s.Success = &v
	return s
}

type SaveUserInfoResponseBodyData struct {
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	Id        *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s SaveUserInfoResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SaveUserInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *SaveUserInfoResponseBodyData) SetGmtCreate(v int64) *SaveUserInfoResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *SaveUserInfoResponseBodyData) SetId(v int64) *SaveUserInfoResponseBodyData {
	s.Id = &v
	return s
}

type SaveUserInfoResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SaveUserInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SaveUserInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveUserInfoResponse) GoString() string {
	return s.String()
}

func (s *SaveUserInfoResponse) SetHeaders(v map[string]*string) *SaveUserInfoResponse {
	s.Headers = v
	return s
}

func (s *SaveUserInfoResponse) SetBody(v *SaveUserInfoResponseBody) *SaveUserInfoResponse {
	s.Body = v
	return s
}

type UnbindDeviceRequest struct {
	SourceIp  *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	IotId     *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s UnbindDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s UnbindDeviceRequest) GoString() string {
	return s.String()
}

func (s *UnbindDeviceRequest) SetSourceIp(v string) *UnbindDeviceRequest {
	s.SourceIp = &v
	return s
}

func (s *UnbindDeviceRequest) SetIotId(v string) *UnbindDeviceRequest {
	s.IotId = &v
	return s
}

func (s *UnbindDeviceRequest) SetProjectId(v string) *UnbindDeviceRequest {
	s.ProjectId = &v
	return s
}

type UnbindDeviceResponseBody struct {
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UnbindDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnbindDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *UnbindDeviceResponseBody) SetRequestId(v string) *UnbindDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnbindDeviceResponseBody) SetHttpStatusCode(v int32) *UnbindDeviceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UnbindDeviceResponseBody) SetSuccess(v bool) *UnbindDeviceResponseBody {
	s.Success = &v
	return s
}

type UnbindDeviceResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UnbindDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UnbindDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s UnbindDeviceResponse) GoString() string {
	return s.String()
}

func (s *UnbindDeviceResponse) SetHeaders(v map[string]*string) *UnbindDeviceResponse {
	s.Headers = v
	return s
}

func (s *UnbindDeviceResponse) SetBody(v *UnbindDeviceResponseBody) *UnbindDeviceResponse {
	s.Body = v
	return s
}

type UpdateDeviceControlInfoRequest struct {
	SourceIp                 *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ProjectId                *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	IotId                    *int64  `json:"IotId,omitempty" xml:"IotId,omitempty"`
	MultiPerson              *int32  `json:"MultiPerson,omitempty" xml:"MultiPerson,omitempty"`
	ControlDoorTime          *int32  `json:"ControlDoorTime,omitempty" xml:"ControlDoorTime,omitempty"`
	EnableMeasureTempurature *int32  `json:"EnableMeasureTempurature,omitempty" xml:"EnableMeasureTempurature,omitempty"`
	SpeedClock               *int32  `json:"SpeedClock,omitempty" xml:"SpeedClock,omitempty"`
}

func (s UpdateDeviceControlInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDeviceControlInfoRequest) GoString() string {
	return s.String()
}

func (s *UpdateDeviceControlInfoRequest) SetSourceIp(v string) *UpdateDeviceControlInfoRequest {
	s.SourceIp = &v
	return s
}

func (s *UpdateDeviceControlInfoRequest) SetProjectId(v string) *UpdateDeviceControlInfoRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateDeviceControlInfoRequest) SetIotId(v int64) *UpdateDeviceControlInfoRequest {
	s.IotId = &v
	return s
}

func (s *UpdateDeviceControlInfoRequest) SetMultiPerson(v int32) *UpdateDeviceControlInfoRequest {
	s.MultiPerson = &v
	return s
}

func (s *UpdateDeviceControlInfoRequest) SetControlDoorTime(v int32) *UpdateDeviceControlInfoRequest {
	s.ControlDoorTime = &v
	return s
}

func (s *UpdateDeviceControlInfoRequest) SetEnableMeasureTempurature(v int32) *UpdateDeviceControlInfoRequest {
	s.EnableMeasureTempurature = &v
	return s
}

func (s *UpdateDeviceControlInfoRequest) SetSpeedClock(v int32) *UpdateDeviceControlInfoRequest {
	s.SpeedClock = &v
	return s
}

type UpdateDeviceControlInfoResponseBody struct {
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateDeviceControlInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateDeviceControlInfoResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDeviceControlInfoResponseBody) SetRequestId(v string) *UpdateDeviceControlInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDeviceControlInfoResponseBody) SetHttpStatusCode(v int32) *UpdateDeviceControlInfoResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateDeviceControlInfoResponseBody) SetSuccess(v bool) *UpdateDeviceControlInfoResponseBody {
	s.Success = &v
	return s
}

type UpdateDeviceControlInfoResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateDeviceControlInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateDeviceControlInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateDeviceControlInfoResponse) GoString() string {
	return s.String()
}

func (s *UpdateDeviceControlInfoResponse) SetHeaders(v map[string]*string) *UpdateDeviceControlInfoResponse {
	s.Headers = v
	return s
}

func (s *UpdateDeviceControlInfoResponse) SetBody(v *UpdateDeviceControlInfoResponseBody) *UpdateDeviceControlInfoResponse {
	s.Body = v
	return s
}

type UpdateDeviceNameRequest struct {
	SourceIp   *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	IotId      *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	ProjectId  *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	DeviceName *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
}

func (s UpdateDeviceNameRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDeviceNameRequest) GoString() string {
	return s.String()
}

func (s *UpdateDeviceNameRequest) SetSourceIp(v string) *UpdateDeviceNameRequest {
	s.SourceIp = &v
	return s
}

func (s *UpdateDeviceNameRequest) SetIotId(v string) *UpdateDeviceNameRequest {
	s.IotId = &v
	return s
}

func (s *UpdateDeviceNameRequest) SetProjectId(v string) *UpdateDeviceNameRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateDeviceNameRequest) SetDeviceName(v string) *UpdateDeviceNameRequest {
	s.DeviceName = &v
	return s
}

type UpdateDeviceNameResponseBody struct {
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateDeviceNameResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateDeviceNameResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDeviceNameResponseBody) SetRequestId(v string) *UpdateDeviceNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDeviceNameResponseBody) SetHttpStatusCode(v int32) *UpdateDeviceNameResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateDeviceNameResponseBody) SetSuccess(v bool) *UpdateDeviceNameResponseBody {
	s.Success = &v
	return s
}

type UpdateDeviceNameResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateDeviceNameResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateDeviceNameResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateDeviceNameResponse) GoString() string {
	return s.String()
}

func (s *UpdateDeviceNameResponse) SetHeaders(v map[string]*string) *UpdateDeviceNameResponse {
	s.Headers = v
	return s
}

func (s *UpdateDeviceNameResponse) SetBody(v *UpdateDeviceNameResponseBody) *UpdateDeviceNameResponse {
	s.Body = v
	return s
}

type UpdateProjectNameRequest struct {
	SourceIp    *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ProjectId   *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s UpdateProjectNameRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateProjectNameRequest) GoString() string {
	return s.String()
}

func (s *UpdateProjectNameRequest) SetSourceIp(v string) *UpdateProjectNameRequest {
	s.SourceIp = &v
	return s
}

func (s *UpdateProjectNameRequest) SetProjectId(v string) *UpdateProjectNameRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateProjectNameRequest) SetProjectName(v string) *UpdateProjectNameRequest {
	s.ProjectName = &v
	return s
}

type UpdateProjectNameResponseBody struct {
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateProjectNameResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateProjectNameResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateProjectNameResponseBody) SetRequestId(v string) *UpdateProjectNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateProjectNameResponseBody) SetHttpStatusCode(v int32) *UpdateProjectNameResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateProjectNameResponseBody) SetSuccess(v bool) *UpdateProjectNameResponseBody {
	s.Success = &v
	return s
}

type UpdateProjectNameResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateProjectNameResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateProjectNameResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateProjectNameResponse) GoString() string {
	return s.String()
}

func (s *UpdateProjectNameResponse) SetHeaders(v map[string]*string) *UpdateProjectNameResponse {
	s.Headers = v
	return s
}

func (s *UpdateProjectNameResponse) SetBody(v *UpdateProjectNameResponseBody) *UpdateProjectNameResponse {
	s.Body = v
	return s
}

type UpdateUserGroupRequest struct {
	SourceIp  *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	GroupId   *int64  `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s UpdateUserGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateUserGroupRequest) SetSourceIp(v string) *UpdateUserGroupRequest {
	s.SourceIp = &v
	return s
}

func (s *UpdateUserGroupRequest) SetGroupId(v int64) *UpdateUserGroupRequest {
	s.GroupId = &v
	return s
}

func (s *UpdateUserGroupRequest) SetGroupName(v string) *UpdateUserGroupRequest {
	s.GroupName = &v
	return s
}

func (s *UpdateUserGroupRequest) SetProjectId(v string) *UpdateUserGroupRequest {
	s.ProjectId = &v
	return s
}

type UpdateUserGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateUserGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateUserGroupResponseBody) SetRequestId(v string) *UpdateUserGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateUserGroupResponseBody) SetSuccess(v bool) *UpdateUserGroupResponseBody {
	s.Success = &v
	return s
}

type UpdateUserGroupResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateUserGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateUserGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserGroupResponse) GoString() string {
	return s.String()
}

func (s *UpdateUserGroupResponse) SetHeaders(v map[string]*string) *UpdateUserGroupResponse {
	s.Headers = v
	return s
}

func (s *UpdateUserGroupResponse) SetBody(v *UpdateUserGroupResponseBody) *UpdateUserGroupResponse {
	s.Body = v
	return s
}

type UpdateUserInfoRequest struct {
	SourceIp        *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	UserGroupId     *int64  `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
	Birthday        *int64  `json:"Birthday,omitempty" xml:"Birthday,omitempty"`
	ImageBase64     *string `json:"ImageBase64,omitempty" xml:"ImageBase64,omitempty"`
	GroupName       *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	ImageUrl        *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	Sex             *int32  `json:"Sex,omitempty" xml:"Sex,omitempty"`
	UserName        *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	UserId          *int64  `json:"UserId,omitempty" xml:"UserId,omitempty"`
	CertificateNo   *string `json:"CertificateNo,omitempty" xml:"CertificateNo,omitempty"`
	PhoneNo         *string `json:"PhoneNo,omitempty" xml:"PhoneNo,omitempty"`
	CertificateType *string `json:"CertificateType,omitempty" xml:"CertificateType,omitempty"`
	ProjectId       *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	OverwriteImg    *bool   `json:"OverwriteImg,omitempty" xml:"OverwriteImg,omitempty"`
}

func (s UpdateUserInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserInfoRequest) GoString() string {
	return s.String()
}

func (s *UpdateUserInfoRequest) SetSourceIp(v string) *UpdateUserInfoRequest {
	s.SourceIp = &v
	return s
}

func (s *UpdateUserInfoRequest) SetUserGroupId(v int64) *UpdateUserInfoRequest {
	s.UserGroupId = &v
	return s
}

func (s *UpdateUserInfoRequest) SetBirthday(v int64) *UpdateUserInfoRequest {
	s.Birthday = &v
	return s
}

func (s *UpdateUserInfoRequest) SetImageBase64(v string) *UpdateUserInfoRequest {
	s.ImageBase64 = &v
	return s
}

func (s *UpdateUserInfoRequest) SetGroupName(v string) *UpdateUserInfoRequest {
	s.GroupName = &v
	return s
}

func (s *UpdateUserInfoRequest) SetImageUrl(v string) *UpdateUserInfoRequest {
	s.ImageUrl = &v
	return s
}

func (s *UpdateUserInfoRequest) SetSex(v int32) *UpdateUserInfoRequest {
	s.Sex = &v
	return s
}

func (s *UpdateUserInfoRequest) SetUserName(v string) *UpdateUserInfoRequest {
	s.UserName = &v
	return s
}

func (s *UpdateUserInfoRequest) SetUserId(v int64) *UpdateUserInfoRequest {
	s.UserId = &v
	return s
}

func (s *UpdateUserInfoRequest) SetCertificateNo(v string) *UpdateUserInfoRequest {
	s.CertificateNo = &v
	return s
}

func (s *UpdateUserInfoRequest) SetPhoneNo(v string) *UpdateUserInfoRequest {
	s.PhoneNo = &v
	return s
}

func (s *UpdateUserInfoRequest) SetCertificateType(v string) *UpdateUserInfoRequest {
	s.CertificateType = &v
	return s
}

func (s *UpdateUserInfoRequest) SetProjectId(v string) *UpdateUserInfoRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateUserInfoRequest) SetOverwriteImg(v bool) *UpdateUserInfoRequest {
	s.OverwriteImg = &v
	return s
}

type UpdateUserInfoResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateUserInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserInfoResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateUserInfoResponseBody) SetRequestId(v string) *UpdateUserInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateUserInfoResponseBody) SetSuccess(v bool) *UpdateUserInfoResponseBody {
	s.Success = &v
	return s
}

type UpdateUserInfoResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateUserInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateUserInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserInfoResponse) GoString() string {
	return s.String()
}

func (s *UpdateUserInfoResponse) SetHeaders(v map[string]*string) *UpdateUserInfoResponse {
	s.Headers = v
	return s
}

func (s *UpdateUserInfoResponse) SetBody(v *UpdateUserInfoResponseBody) *UpdateUserInfoResponse {
	s.Body = v
	return s
}

type UploadIdentifyRecordRequest struct {
	SourceIp               *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	UserId                 *int64  `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserName               *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	ProjectId              *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	IotId                  *string `json:"IotId,omitempty" xml:"IotId,omitempty"`
	IdentifyingImageBase64 *string `json:"IdentifyingImageBase64,omitempty" xml:"IdentifyingImageBase64,omitempty"`
	IdentifyingTime        *int64  `json:"IdentifyingTime,omitempty" xml:"IdentifyingTime,omitempty"`
	IdentifyingImageUrl    *string `json:"IdentifyingImageUrl,omitempty" xml:"IdentifyingImageUrl,omitempty"`
	DeviceName             *string `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
}

func (s UploadIdentifyRecordRequest) String() string {
	return tea.Prettify(s)
}

func (s UploadIdentifyRecordRequest) GoString() string {
	return s.String()
}

func (s *UploadIdentifyRecordRequest) SetSourceIp(v string) *UploadIdentifyRecordRequest {
	s.SourceIp = &v
	return s
}

func (s *UploadIdentifyRecordRequest) SetUserId(v int64) *UploadIdentifyRecordRequest {
	s.UserId = &v
	return s
}

func (s *UploadIdentifyRecordRequest) SetUserName(v string) *UploadIdentifyRecordRequest {
	s.UserName = &v
	return s
}

func (s *UploadIdentifyRecordRequest) SetProjectId(v string) *UploadIdentifyRecordRequest {
	s.ProjectId = &v
	return s
}

func (s *UploadIdentifyRecordRequest) SetIotId(v string) *UploadIdentifyRecordRequest {
	s.IotId = &v
	return s
}

func (s *UploadIdentifyRecordRequest) SetIdentifyingImageBase64(v string) *UploadIdentifyRecordRequest {
	s.IdentifyingImageBase64 = &v
	return s
}

func (s *UploadIdentifyRecordRequest) SetIdentifyingTime(v int64) *UploadIdentifyRecordRequest {
	s.IdentifyingTime = &v
	return s
}

func (s *UploadIdentifyRecordRequest) SetIdentifyingImageUrl(v string) *UploadIdentifyRecordRequest {
	s.IdentifyingImageUrl = &v
	return s
}

func (s *UploadIdentifyRecordRequest) SetDeviceName(v string) *UploadIdentifyRecordRequest {
	s.DeviceName = &v
	return s
}

type UploadIdentifyRecordResponseBody struct {
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UploadIdentifyRecordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UploadIdentifyRecordResponseBody) GoString() string {
	return s.String()
}

func (s *UploadIdentifyRecordResponseBody) SetRequestId(v string) *UploadIdentifyRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *UploadIdentifyRecordResponseBody) SetHttpStatusCode(v int32) *UploadIdentifyRecordResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UploadIdentifyRecordResponseBody) SetSuccess(v bool) *UploadIdentifyRecordResponseBody {
	s.Success = &v
	return s
}

type UploadIdentifyRecordResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UploadIdentifyRecordResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UploadIdentifyRecordResponse) String() string {
	return tea.Prettify(s)
}

func (s UploadIdentifyRecordResponse) GoString() string {
	return s.String()
}

func (s *UploadIdentifyRecordResponse) SetHeaders(v map[string]*string) *UploadIdentifyRecordResponse {
	s.Headers = v
	return s
}

func (s *UploadIdentifyRecordResponse) SetBody(v *UploadIdentifyRecordResponseBody) *UploadIdentifyRecordResponse {
	s.Body = v
	return s
}

type VerifyAccountProjectRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s VerifyAccountProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s VerifyAccountProjectRequest) GoString() string {
	return s.String()
}

func (s *VerifyAccountProjectRequest) SetSourceIp(v string) *VerifyAccountProjectRequest {
	s.SourceIp = &v
	return s
}

type VerifyAccountProjectResponseBody struct {
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s VerifyAccountProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s VerifyAccountProjectResponseBody) GoString() string {
	return s.String()
}

func (s *VerifyAccountProjectResponseBody) SetRequestId(v string) *VerifyAccountProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *VerifyAccountProjectResponseBody) SetHttpStatusCode(v int32) *VerifyAccountProjectResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *VerifyAccountProjectResponseBody) SetSuccess(v bool) *VerifyAccountProjectResponseBody {
	s.Success = &v
	return s
}

type VerifyAccountProjectResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *VerifyAccountProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s VerifyAccountProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s VerifyAccountProjectResponse) GoString() string {
	return s.String()
}

func (s *VerifyAccountProjectResponse) SetHeaders(v map[string]*string) *VerifyAccountProjectResponse {
	s.Headers = v
	return s
}

func (s *VerifyAccountProjectResponse) SetBody(v *VerifyAccountProjectResponseBody) *VerifyAccountProjectResponse {
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
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("cloudauth-console"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) CreateProjectWithOptions(request *CreateProjectRequest, runtime *util.RuntimeOptions) (_result *CreateProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateProjectResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateProject"), tea.String("2019-02-19"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateProject(request *CreateProjectRequest) (_result *CreateProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateProjectResponse{}
	_body, _err := client.CreateProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateSlrRoleWithOptions(request *CreateSlrRoleRequest, runtime *util.RuntimeOptions) (_result *CreateSlrRoleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateSlrRoleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateSlrRole"), tea.String("2019-02-19"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSlrRole(request *CreateSlrRoleRequest) (_result *CreateSlrRoleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateSlrRoleResponse{}
	_body, _err := client.CreateSlrRoleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteMnsServeWithOptions(request *DeleteMnsServeRequest, runtime *util.RuntimeOptions) (_result *DeleteMnsServeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteMnsServeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteMnsServe"), tea.String("2019-02-19"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteMnsServe(request *DeleteMnsServeRequest) (_result *DeleteMnsServeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteMnsServeResponse{}
	_body, _err := client.DeleteMnsServeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteUserGroupWithOptions(request *DeleteUserGroupRequest, runtime *util.RuntimeOptions) (_result *DeleteUserGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteUserGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteUserGroup"), tea.String("2019-02-19"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteUserGroup(request *DeleteUserGroupRequest) (_result *DeleteUserGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteUserGroupResponse{}
	_body, _err := client.DeleteUserGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteUserInfoWithOptions(request *DeleteUserInfoRequest, runtime *util.RuntimeOptions) (_result *DeleteUserInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteUserInfoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteUserInfo"), tea.String("2019-02-19"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteUserInfo(request *DeleteUserInfoRequest) (_result *DeleteUserInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteUserInfoResponse{}
	_body, _err := client.DeleteUserInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAllEndPointWithOptions(request *DescribeAllEndPointRequest, runtime *util.RuntimeOptions) (_result *DescribeAllEndPointResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeAllEndPointResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeAllEndPoint"), tea.String("2019-02-19"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAllEndPoint(request *DescribeAllEndPointRequest) (_result *DescribeAllEndPointResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAllEndPointResponse{}
	_body, _err := client.DescribeAllEndPointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeBindUserIdListWithOptions(request *DescribeBindUserIdListRequest, runtime *util.RuntimeOptions) (_result *DescribeBindUserIdListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeBindUserIdListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeBindUserIdList"), tea.String("2019-02-19"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeBindUserIdList(request *DescribeBindUserIdListRequest) (_result *DescribeBindUserIdListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeBindUserIdListResponse{}
	_body, _err := client.DescribeBindUserIdListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCertificateTypeListWithOptions(request *DescribeCertificateTypeListRequest, runtime *util.RuntimeOptions) (_result *DescribeCertificateTypeListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeCertificateTypeListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeCertificateTypeList"), tea.String("2019-02-19"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCertificateTypeList(request *DescribeCertificateTypeListRequest) (_result *DescribeCertificateTypeListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCertificateTypeListResponse{}
	_body, _err := client.DescribeCertificateTypeListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDeviceListWithOptions(request *DescribeDeviceListRequest, runtime *util.RuntimeOptions) (_result *DescribeDeviceListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDeviceListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDeviceList"), tea.String("2019-02-19"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDeviceList(request *DescribeDeviceListRequest) (_result *DescribeDeviceListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDeviceListResponse{}
	_body, _err := client.DescribeDeviceListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeExcelAnalysisResultWithOptions(request *DescribeExcelAnalysisResultRequest, runtime *util.RuntimeOptions) (_result *DescribeExcelAnalysisResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeExcelAnalysisResultResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeExcelAnalysisResult"), tea.String("2019-02-19"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeExcelAnalysisResult(request *DescribeExcelAnalysisResultRequest) (_result *DescribeExcelAnalysisResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeExcelAnalysisResultResponse{}
	_body, _err := client.DescribeExcelAnalysisResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeIdentifyRecordListWithOptions(request *DescribeIdentifyRecordListRequest, runtime *util.RuntimeOptions) (_result *DescribeIdentifyRecordListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeIdentifyRecordListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeIdentifyRecordList"), tea.String("2019-02-19"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeIdentifyRecordList(request *DescribeIdentifyRecordListRequest) (_result *DescribeIdentifyRecordListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeIdentifyRecordListResponse{}
	_body, _err := client.DescribeIdentifyRecordListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeMnsOauthWithOptions(request *DescribeMnsOauthRequest, runtime *util.RuntimeOptions) (_result *DescribeMnsOauthResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeMnsOauthResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeMnsOauth"), tea.String("2019-02-19"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeMnsOauth(request *DescribeMnsOauthRequest) (_result *DescribeMnsOauthResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeMnsOauthResponse{}
	_body, _err := client.DescribeMnsOauthWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeOssOauthWithOptions(request *DescribeOssOauthRequest, runtime *util.RuntimeOptions) (_result *DescribeOssOauthResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeOssOauthResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeOssOauth"), tea.String("2019-02-19"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeOssOauth(request *DescribeOssOauthRequest) (_result *DescribeOssOauthResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeOssOauthResponse{}
	_body, _err := client.DescribeOssOauthWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSignedUrlWithOptions(request *DescribeSignedUrlRequest, runtime *util.RuntimeOptions) (_result *DescribeSignedUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeSignedUrlResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeSignedUrl"), tea.String("2019-02-19"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSignedUrl(request *DescribeSignedUrlRequest) (_result *DescribeSignedUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSignedUrlResponse{}
	_body, _err := client.DescribeSignedUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeTopicWithOptions(request *DescribeTopicRequest, runtime *util.RuntimeOptions) (_result *DescribeTopicResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeTopicResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeTopic"), tea.String("2019-02-19"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeTopic(request *DescribeTopicRequest) (_result *DescribeTopicResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeTopicResponse{}
	_body, _err := client.DescribeTopicWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeUploadPreSignWithOptions(request *DescribeUploadPreSignRequest, runtime *util.RuntimeOptions) (_result *DescribeUploadPreSignResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeUploadPreSignResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeUploadPreSign"), tea.String("2019-02-19"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeUploadPreSign(request *DescribeUploadPreSignRequest) (_result *DescribeUploadPreSignResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeUploadPreSignResponse{}
	_body, _err := client.DescribeUploadPreSignWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeUserGroupListWithOptions(request *DescribeUserGroupListRequest, runtime *util.RuntimeOptions) (_result *DescribeUserGroupListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeUserGroupListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeUserGroupList"), tea.String("2019-02-19"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeUserGroupList(request *DescribeUserGroupListRequest) (_result *DescribeUserGroupListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeUserGroupListResponse{}
	_body, _err := client.DescribeUserGroupListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeUserInfoListWithOptions(request *DescribeUserInfoListRequest, runtime *util.RuntimeOptions) (_result *DescribeUserInfoListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeUserInfoListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeUserInfoList"), tea.String("2019-02-19"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeUserInfoList(request *DescribeUserInfoListRequest) (_result *DescribeUserInfoListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeUserInfoListResponse{}
	_body, _err := client.DescribeUserInfoListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAccountProjectWithOptions(request *GetAccountProjectRequest, runtime *util.RuntimeOptions) (_result *GetAccountProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetAccountProjectResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetAccountProject"), tea.String("2019-02-19"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAccountProject(request *GetAccountProjectRequest) (_result *GetAccountProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAccountProjectResponse{}
	_body, _err := client.GetAccountProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SaveMnsServeWithOptions(request *SaveMnsServeRequest, runtime *util.RuntimeOptions) (_result *SaveMnsServeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SaveMnsServeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SaveMnsServe"), tea.String("2019-02-19"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SaveMnsServe(request *SaveMnsServeRequest) (_result *SaveMnsServeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SaveMnsServeResponse{}
	_body, _err := client.SaveMnsServeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SaveOssServeWithOptions(request *SaveOssServeRequest, runtime *util.RuntimeOptions) (_result *SaveOssServeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SaveOssServeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SaveOssServe"), tea.String("2019-02-19"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SaveOssServe(request *SaveOssServeRequest) (_result *SaveOssServeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SaveOssServeResponse{}
	_body, _err := client.SaveOssServeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SaveUserGroupWithOptions(request *SaveUserGroupRequest, runtime *util.RuntimeOptions) (_result *SaveUserGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SaveUserGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SaveUserGroup"), tea.String("2019-02-19"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SaveUserGroup(request *SaveUserGroupRequest) (_result *SaveUserGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SaveUserGroupResponse{}
	_body, _err := client.SaveUserGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SaveUserInfoWithOptions(request *SaveUserInfoRequest, runtime *util.RuntimeOptions) (_result *SaveUserInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SaveUserInfoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SaveUserInfo"), tea.String("2019-02-19"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SaveUserInfo(request *SaveUserInfoRequest) (_result *SaveUserInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SaveUserInfoResponse{}
	_body, _err := client.SaveUserInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UnbindDeviceWithOptions(request *UnbindDeviceRequest, runtime *util.RuntimeOptions) (_result *UnbindDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UnbindDeviceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UnbindDevice"), tea.String("2019-02-19"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UnbindDevice(request *UnbindDeviceRequest) (_result *UnbindDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UnbindDeviceResponse{}
	_body, _err := client.UnbindDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateDeviceControlInfoWithOptions(request *UpdateDeviceControlInfoRequest, runtime *util.RuntimeOptions) (_result *UpdateDeviceControlInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateDeviceControlInfoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateDeviceControlInfo"), tea.String("2019-02-19"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateDeviceControlInfo(request *UpdateDeviceControlInfoRequest) (_result *UpdateDeviceControlInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateDeviceControlInfoResponse{}
	_body, _err := client.UpdateDeviceControlInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateDeviceNameWithOptions(request *UpdateDeviceNameRequest, runtime *util.RuntimeOptions) (_result *UpdateDeviceNameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateDeviceNameResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateDeviceName"), tea.String("2019-02-19"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateDeviceName(request *UpdateDeviceNameRequest) (_result *UpdateDeviceNameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateDeviceNameResponse{}
	_body, _err := client.UpdateDeviceNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateProjectNameWithOptions(request *UpdateProjectNameRequest, runtime *util.RuntimeOptions) (_result *UpdateProjectNameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateProjectNameResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateProjectName"), tea.String("2019-02-19"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateProjectName(request *UpdateProjectNameRequest) (_result *UpdateProjectNameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateProjectNameResponse{}
	_body, _err := client.UpdateProjectNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateUserGroupWithOptions(request *UpdateUserGroupRequest, runtime *util.RuntimeOptions) (_result *UpdateUserGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateUserGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateUserGroup"), tea.String("2019-02-19"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateUserGroup(request *UpdateUserGroupRequest) (_result *UpdateUserGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateUserGroupResponse{}
	_body, _err := client.UpdateUserGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateUserInfoWithOptions(request *UpdateUserInfoRequest, runtime *util.RuntimeOptions) (_result *UpdateUserInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateUserInfoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateUserInfo"), tea.String("2019-02-19"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateUserInfo(request *UpdateUserInfoRequest) (_result *UpdateUserInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateUserInfoResponse{}
	_body, _err := client.UpdateUserInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UploadIdentifyRecordWithOptions(request *UploadIdentifyRecordRequest, runtime *util.RuntimeOptions) (_result *UploadIdentifyRecordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UploadIdentifyRecordResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UploadIdentifyRecord"), tea.String("2019-02-19"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UploadIdentifyRecord(request *UploadIdentifyRecordRequest) (_result *UploadIdentifyRecordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UploadIdentifyRecordResponse{}
	_body, _err := client.UploadIdentifyRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) VerifyAccountProjectWithOptions(request *VerifyAccountProjectRequest, runtime *util.RuntimeOptions) (_result *VerifyAccountProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &VerifyAccountProjectResponse{}
	_body, _err := client.DoRPCRequest(tea.String("VerifyAccountProject"), tea.String("2019-02-19"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) VerifyAccountProject(request *VerifyAccountProjectRequest) (_result *VerifyAccountProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &VerifyAccountProjectResponse{}
	_body, _err := client.VerifyAccountProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
