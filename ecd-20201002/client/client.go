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

type ApproveFotaUpdateRequest struct {
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	ClientId   *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	DesktopId  *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	LoginToken *string `json:"LoginToken,omitempty" xml:"LoginToken,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SessionId  *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	Uuid       *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s ApproveFotaUpdateRequest) String() string {
	return tea.Prettify(s)
}

func (s ApproveFotaUpdateRequest) GoString() string {
	return s.String()
}

func (s *ApproveFotaUpdateRequest) SetAppVersion(v string) *ApproveFotaUpdateRequest {
	s.AppVersion = &v
	return s
}

func (s *ApproveFotaUpdateRequest) SetClientId(v string) *ApproveFotaUpdateRequest {
	s.ClientId = &v
	return s
}

func (s *ApproveFotaUpdateRequest) SetDesktopId(v string) *ApproveFotaUpdateRequest {
	s.DesktopId = &v
	return s
}

func (s *ApproveFotaUpdateRequest) SetLoginToken(v string) *ApproveFotaUpdateRequest {
	s.LoginToken = &v
	return s
}

func (s *ApproveFotaUpdateRequest) SetRegionId(v string) *ApproveFotaUpdateRequest {
	s.RegionId = &v
	return s
}

func (s *ApproveFotaUpdateRequest) SetSessionId(v string) *ApproveFotaUpdateRequest {
	s.SessionId = &v
	return s
}

func (s *ApproveFotaUpdateRequest) SetUuid(v string) *ApproveFotaUpdateRequest {
	s.Uuid = &v
	return s
}

type ApproveFotaUpdateResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ApproveFotaUpdateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ApproveFotaUpdateResponseBody) GoString() string {
	return s.String()
}

func (s *ApproveFotaUpdateResponseBody) SetRequestId(v string) *ApproveFotaUpdateResponseBody {
	s.RequestId = &v
	return s
}

type ApproveFotaUpdateResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ApproveFotaUpdateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ApproveFotaUpdateResponse) String() string {
	return tea.Prettify(s)
}

func (s ApproveFotaUpdateResponse) GoString() string {
	return s.String()
}

func (s *ApproveFotaUpdateResponse) SetHeaders(v map[string]*string) *ApproveFotaUpdateResponse {
	s.Headers = v
	return s
}

func (s *ApproveFotaUpdateResponse) SetStatusCode(v int32) *ApproveFotaUpdateResponse {
	s.StatusCode = &v
	return s
}

func (s *ApproveFotaUpdateResponse) SetBody(v *ApproveFotaUpdateResponseBody) *ApproveFotaUpdateResponse {
	s.Body = v
	return s
}

type ChangePasswordRequest struct {
	ClientId     *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	EndUserId    *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	LoginToken   *string `json:"LoginToken,omitempty" xml:"LoginToken,omitempty"`
	NewPassword  *string `json:"NewPassword,omitempty" xml:"NewPassword,omitempty"`
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	OldPassword  *string `json:"OldPassword,omitempty" xml:"OldPassword,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SessionId    *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s ChangePasswordRequest) String() string {
	return tea.Prettify(s)
}

func (s ChangePasswordRequest) GoString() string {
	return s.String()
}

func (s *ChangePasswordRequest) SetClientId(v string) *ChangePasswordRequest {
	s.ClientId = &v
	return s
}

func (s *ChangePasswordRequest) SetEndUserId(v string) *ChangePasswordRequest {
	s.EndUserId = &v
	return s
}

func (s *ChangePasswordRequest) SetLoginToken(v string) *ChangePasswordRequest {
	s.LoginToken = &v
	return s
}

func (s *ChangePasswordRequest) SetNewPassword(v string) *ChangePasswordRequest {
	s.NewPassword = &v
	return s
}

func (s *ChangePasswordRequest) SetOfficeSiteId(v string) *ChangePasswordRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *ChangePasswordRequest) SetOldPassword(v string) *ChangePasswordRequest {
	s.OldPassword = &v
	return s
}

func (s *ChangePasswordRequest) SetRegionId(v string) *ChangePasswordRequest {
	s.RegionId = &v
	return s
}

func (s *ChangePasswordRequest) SetSessionId(v string) *ChangePasswordRequest {
	s.SessionId = &v
	return s
}

type ChangePasswordResponseBody struct {
	LoginToken *string `json:"LoginToken,omitempty" xml:"LoginToken,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ChangePasswordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ChangePasswordResponseBody) GoString() string {
	return s.String()
}

func (s *ChangePasswordResponseBody) SetLoginToken(v string) *ChangePasswordResponseBody {
	s.LoginToken = &v
	return s
}

func (s *ChangePasswordResponseBody) SetRequestId(v string) *ChangePasswordResponseBody {
	s.RequestId = &v
	return s
}

type ChangePasswordResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChangePasswordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChangePasswordResponse) String() string {
	return tea.Prettify(s)
}

func (s ChangePasswordResponse) GoString() string {
	return s.String()
}

func (s *ChangePasswordResponse) SetHeaders(v map[string]*string) *ChangePasswordResponse {
	s.Headers = v
	return s
}

func (s *ChangePasswordResponse) SetStatusCode(v int32) *ChangePasswordResponse {
	s.StatusCode = &v
	return s
}

func (s *ChangePasswordResponse) SetBody(v *ChangePasswordResponseBody) *ChangePasswordResponse {
	s.Body = v
	return s
}

type DeleteFingerPrintTemplateRequest struct {
	ClientId    *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Index       *int32  `json:"Index,omitempty" xml:"Index,omitempty"`
	LoginToken  *string `json:"LoginToken,omitempty" xml:"LoginToken,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SessionId   *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s DeleteFingerPrintTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteFingerPrintTemplateRequest) GoString() string {
	return s.String()
}

func (s *DeleteFingerPrintTemplateRequest) SetClientId(v string) *DeleteFingerPrintTemplateRequest {
	s.ClientId = &v
	return s
}

func (s *DeleteFingerPrintTemplateRequest) SetClientToken(v string) *DeleteFingerPrintTemplateRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteFingerPrintTemplateRequest) SetIndex(v int32) *DeleteFingerPrintTemplateRequest {
	s.Index = &v
	return s
}

func (s *DeleteFingerPrintTemplateRequest) SetLoginToken(v string) *DeleteFingerPrintTemplateRequest {
	s.LoginToken = &v
	return s
}

func (s *DeleteFingerPrintTemplateRequest) SetRegionId(v string) *DeleteFingerPrintTemplateRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteFingerPrintTemplateRequest) SetSessionId(v string) *DeleteFingerPrintTemplateRequest {
	s.SessionId = &v
	return s
}

type DeleteFingerPrintTemplateResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteFingerPrintTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteFingerPrintTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFingerPrintTemplateResponseBody) SetRequestId(v string) *DeleteFingerPrintTemplateResponseBody {
	s.RequestId = &v
	return s
}

type DeleteFingerPrintTemplateResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteFingerPrintTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteFingerPrintTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteFingerPrintTemplateResponse) GoString() string {
	return s.String()
}

func (s *DeleteFingerPrintTemplateResponse) SetHeaders(v map[string]*string) *DeleteFingerPrintTemplateResponse {
	s.Headers = v
	return s
}

func (s *DeleteFingerPrintTemplateResponse) SetStatusCode(v int32) *DeleteFingerPrintTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteFingerPrintTemplateResponse) SetBody(v *DeleteFingerPrintTemplateResponseBody) *DeleteFingerPrintTemplateResponse {
	s.Body = v
	return s
}

type DescribeDirectoriesRequest struct {
	ClientId    *string   `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	DirectoryId []*string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty" type:"Repeated"`
	RegionId    *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeDirectoriesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDirectoriesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDirectoriesRequest) SetClientId(v string) *DescribeDirectoriesRequest {
	s.ClientId = &v
	return s
}

func (s *DescribeDirectoriesRequest) SetDirectoryId(v []*string) *DescribeDirectoriesRequest {
	s.DirectoryId = v
	return s
}

func (s *DescribeDirectoriesRequest) SetRegionId(v string) *DescribeDirectoriesRequest {
	s.RegionId = &v
	return s
}

type DescribeDirectoriesResponseBody struct {
	Directories []*DescribeDirectoriesResponseBodyDirectories `json:"Directories,omitempty" xml:"Directories,omitempty" type:"Repeated"`
	RequestId   *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDirectoriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDirectoriesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDirectoriesResponseBody) SetDirectories(v []*DescribeDirectoriesResponseBodyDirectories) *DescribeDirectoriesResponseBody {
	s.Directories = v
	return s
}

func (s *DescribeDirectoriesResponseBody) SetRequestId(v string) *DescribeDirectoriesResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDirectoriesResponseBodyDirectories struct {
	DesktopAccessType *string `json:"DesktopAccessType,omitempty" xml:"DesktopAccessType,omitempty"`
	DirectoryId       *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	DirectoryType     *string `json:"DirectoryType,omitempty" xml:"DirectoryType,omitempty"`
	ProviderId        *string `json:"ProviderId,omitempty" xml:"ProviderId,omitempty"`
	SsoServiceUrl     *string `json:"SsoServiceUrl,omitempty" xml:"SsoServiceUrl,omitempty"`
}

func (s DescribeDirectoriesResponseBodyDirectories) String() string {
	return tea.Prettify(s)
}

func (s DescribeDirectoriesResponseBodyDirectories) GoString() string {
	return s.String()
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetDesktopAccessType(v string) *DescribeDirectoriesResponseBodyDirectories {
	s.DesktopAccessType = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetDirectoryId(v string) *DescribeDirectoriesResponseBodyDirectories {
	s.DirectoryId = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetDirectoryType(v string) *DescribeDirectoriesResponseBodyDirectories {
	s.DirectoryType = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetProviderId(v string) *DescribeDirectoriesResponseBodyDirectories {
	s.ProviderId = &v
	return s
}

func (s *DescribeDirectoriesResponseBodyDirectories) SetSsoServiceUrl(v string) *DescribeDirectoriesResponseBodyDirectories {
	s.SsoServiceUrl = &v
	return s
}

type DescribeDirectoriesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDirectoriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDirectoriesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDirectoriesResponse) GoString() string {
	return s.String()
}

func (s *DescribeDirectoriesResponse) SetHeaders(v map[string]*string) *DescribeDirectoriesResponse {
	s.Headers = v
	return s
}

func (s *DescribeDirectoriesResponse) SetStatusCode(v int32) *DescribeDirectoriesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDirectoriesResponse) SetBody(v *DescribeDirectoriesResponseBody) *DescribeDirectoriesResponse {
	s.Body = v
	return s
}

type DescribeFingerPrintTemplatesRequest struct {
	ClientId   *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	LoginToken *string `json:"LoginToken,omitempty" xml:"LoginToken,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SessionId  *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s DescribeFingerPrintTemplatesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFingerPrintTemplatesRequest) GoString() string {
	return s.String()
}

func (s *DescribeFingerPrintTemplatesRequest) SetClientId(v string) *DescribeFingerPrintTemplatesRequest {
	s.ClientId = &v
	return s
}

func (s *DescribeFingerPrintTemplatesRequest) SetLoginToken(v string) *DescribeFingerPrintTemplatesRequest {
	s.LoginToken = &v
	return s
}

func (s *DescribeFingerPrintTemplatesRequest) SetRegionId(v string) *DescribeFingerPrintTemplatesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeFingerPrintTemplatesRequest) SetSessionId(v string) *DescribeFingerPrintTemplatesRequest {
	s.SessionId = &v
	return s
}

type DescribeFingerPrintTemplatesResponseBody struct {
	FingerPrintTemplates []*DescribeFingerPrintTemplatesResponseBodyFingerPrintTemplates `json:"FingerPrintTemplates,omitempty" xml:"FingerPrintTemplates,omitempty" type:"Repeated"`
	RequestId            *string                                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeFingerPrintTemplatesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFingerPrintTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFingerPrintTemplatesResponseBody) SetFingerPrintTemplates(v []*DescribeFingerPrintTemplatesResponseBodyFingerPrintTemplates) *DescribeFingerPrintTemplatesResponseBody {
	s.FingerPrintTemplates = v
	return s
}

func (s *DescribeFingerPrintTemplatesResponseBody) SetRequestId(v string) *DescribeFingerPrintTemplatesResponseBody {
	s.RequestId = &v
	return s
}

type DescribeFingerPrintTemplatesResponseBodyFingerPrintTemplates struct {
	ClientId     *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	EndUserId    *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	Index        *int64  `json:"Index,omitempty" xml:"Index,omitempty"`
	LoginTime    *string `json:"LoginTime,omitempty" xml:"LoginTime,omitempty"`
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
}

func (s DescribeFingerPrintTemplatesResponseBodyFingerPrintTemplates) String() string {
	return tea.Prettify(s)
}

func (s DescribeFingerPrintTemplatesResponseBodyFingerPrintTemplates) GoString() string {
	return s.String()
}

func (s *DescribeFingerPrintTemplatesResponseBodyFingerPrintTemplates) SetClientId(v string) *DescribeFingerPrintTemplatesResponseBodyFingerPrintTemplates {
	s.ClientId = &v
	return s
}

func (s *DescribeFingerPrintTemplatesResponseBodyFingerPrintTemplates) SetCreationTime(v string) *DescribeFingerPrintTemplatesResponseBodyFingerPrintTemplates {
	s.CreationTime = &v
	return s
}

func (s *DescribeFingerPrintTemplatesResponseBodyFingerPrintTemplates) SetDescription(v string) *DescribeFingerPrintTemplatesResponseBodyFingerPrintTemplates {
	s.Description = &v
	return s
}

func (s *DescribeFingerPrintTemplatesResponseBodyFingerPrintTemplates) SetEndUserId(v string) *DescribeFingerPrintTemplatesResponseBodyFingerPrintTemplates {
	s.EndUserId = &v
	return s
}

func (s *DescribeFingerPrintTemplatesResponseBodyFingerPrintTemplates) SetIndex(v int64) *DescribeFingerPrintTemplatesResponseBodyFingerPrintTemplates {
	s.Index = &v
	return s
}

func (s *DescribeFingerPrintTemplatesResponseBodyFingerPrintTemplates) SetLoginTime(v string) *DescribeFingerPrintTemplatesResponseBodyFingerPrintTemplates {
	s.LoginTime = &v
	return s
}

func (s *DescribeFingerPrintTemplatesResponseBodyFingerPrintTemplates) SetOfficeSiteId(v string) *DescribeFingerPrintTemplatesResponseBodyFingerPrintTemplates {
	s.OfficeSiteId = &v
	return s
}

type DescribeFingerPrintTemplatesResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFingerPrintTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeFingerPrintTemplatesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFingerPrintTemplatesResponse) GoString() string {
	return s.String()
}

func (s *DescribeFingerPrintTemplatesResponse) SetHeaders(v map[string]*string) *DescribeFingerPrintTemplatesResponse {
	s.Headers = v
	return s
}

func (s *DescribeFingerPrintTemplatesResponse) SetStatusCode(v int32) *DescribeFingerPrintTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFingerPrintTemplatesResponse) SetBody(v *DescribeFingerPrintTemplatesResponseBody) *DescribeFingerPrintTemplatesResponse {
	s.Body = v
	return s
}

type DescribeGlobalDesktopsRequest struct {
	ClientId          *string   `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	DesktopAccessType *string   `json:"DesktopAccessType,omitempty" xml:"DesktopAccessType,omitempty"`
	DesktopId         []*string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" type:"Repeated"`
	DesktopName       *string   `json:"DesktopName,omitempty" xml:"DesktopName,omitempty"`
	DesktopStatus     *string   `json:"DesktopStatus,omitempty" xml:"DesktopStatus,omitempty"`
	DirectoryId       *string   `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// 关键字。支持模糊搜索桌面ID、云桌面名称和终端用户自定义的桌面名称。
	Keyword         *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	LoginRegionId   *string `json:"LoginRegionId,omitempty" xml:"LoginRegionId,omitempty"`
	LoginToken      *string `json:"LoginToken,omitempty" xml:"LoginToken,omitempty"`
	MaxResults      *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken       *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OfficeSiteId    *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	OrderBy         *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	QueryFotaUpdate *bool   `json:"QueryFotaUpdate,omitempty" xml:"QueryFotaUpdate,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SearchRegionId  *string `json:"SearchRegionId,omitempty" xml:"SearchRegionId,omitempty"`
	SessionId       *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	SortType        *string `json:"SortType,omitempty" xml:"SortType,omitempty"`
	WithoutLatency  *bool   `json:"WithoutLatency,omitempty" xml:"WithoutLatency,omitempty"`
}

func (s DescribeGlobalDesktopsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeGlobalDesktopsRequest) GoString() string {
	return s.String()
}

func (s *DescribeGlobalDesktopsRequest) SetClientId(v string) *DescribeGlobalDesktopsRequest {
	s.ClientId = &v
	return s
}

func (s *DescribeGlobalDesktopsRequest) SetDesktopAccessType(v string) *DescribeGlobalDesktopsRequest {
	s.DesktopAccessType = &v
	return s
}

func (s *DescribeGlobalDesktopsRequest) SetDesktopId(v []*string) *DescribeGlobalDesktopsRequest {
	s.DesktopId = v
	return s
}

func (s *DescribeGlobalDesktopsRequest) SetDesktopName(v string) *DescribeGlobalDesktopsRequest {
	s.DesktopName = &v
	return s
}

func (s *DescribeGlobalDesktopsRequest) SetDesktopStatus(v string) *DescribeGlobalDesktopsRequest {
	s.DesktopStatus = &v
	return s
}

func (s *DescribeGlobalDesktopsRequest) SetDirectoryId(v string) *DescribeGlobalDesktopsRequest {
	s.DirectoryId = &v
	return s
}

func (s *DescribeGlobalDesktopsRequest) SetKeyword(v string) *DescribeGlobalDesktopsRequest {
	s.Keyword = &v
	return s
}

func (s *DescribeGlobalDesktopsRequest) SetLoginRegionId(v string) *DescribeGlobalDesktopsRequest {
	s.LoginRegionId = &v
	return s
}

func (s *DescribeGlobalDesktopsRequest) SetLoginToken(v string) *DescribeGlobalDesktopsRequest {
	s.LoginToken = &v
	return s
}

func (s *DescribeGlobalDesktopsRequest) SetMaxResults(v int32) *DescribeGlobalDesktopsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeGlobalDesktopsRequest) SetNextToken(v string) *DescribeGlobalDesktopsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeGlobalDesktopsRequest) SetOfficeSiteId(v string) *DescribeGlobalDesktopsRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *DescribeGlobalDesktopsRequest) SetOrderBy(v string) *DescribeGlobalDesktopsRequest {
	s.OrderBy = &v
	return s
}

func (s *DescribeGlobalDesktopsRequest) SetQueryFotaUpdate(v bool) *DescribeGlobalDesktopsRequest {
	s.QueryFotaUpdate = &v
	return s
}

func (s *DescribeGlobalDesktopsRequest) SetRegionId(v string) *DescribeGlobalDesktopsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeGlobalDesktopsRequest) SetSearchRegionId(v string) *DescribeGlobalDesktopsRequest {
	s.SearchRegionId = &v
	return s
}

func (s *DescribeGlobalDesktopsRequest) SetSessionId(v string) *DescribeGlobalDesktopsRequest {
	s.SessionId = &v
	return s
}

func (s *DescribeGlobalDesktopsRequest) SetSortType(v string) *DescribeGlobalDesktopsRequest {
	s.SortType = &v
	return s
}

func (s *DescribeGlobalDesktopsRequest) SetWithoutLatency(v bool) *DescribeGlobalDesktopsRequest {
	s.WithoutLatency = &v
	return s
}

type DescribeGlobalDesktopsResponseBody struct {
	Desktops  []*DescribeGlobalDesktopsResponseBodyDesktops `json:"Desktops,omitempty" xml:"Desktops,omitempty" type:"Repeated"`
	NextToken *string                                       `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeGlobalDesktopsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeGlobalDesktopsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGlobalDesktopsResponseBody) SetDesktops(v []*DescribeGlobalDesktopsResponseBodyDesktops) *DescribeGlobalDesktopsResponseBody {
	s.Desktops = v
	return s
}

func (s *DescribeGlobalDesktopsResponseBody) SetNextToken(v string) *DescribeGlobalDesktopsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBody) SetRequestId(v string) *DescribeGlobalDesktopsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeGlobalDesktopsResponseBodyDesktops struct {
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// 支持的客户端信息
	Clients            []*DescribeGlobalDesktopsResponseBodyDesktopsClients       `json:"Clients,omitempty" xml:"Clients,omitempty" type:"Repeated"`
	ConnectionStatus   *string                                                    `json:"ConnectionStatus,omitempty" xml:"ConnectionStatus,omitempty"`
	Cpu                *int32                                                     `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	CreationTime       *string                                                    `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	DesktopGroupId     *string                                                    `json:"DesktopGroupId,omitempty" xml:"DesktopGroupId,omitempty"`
	DesktopId          *string                                                    `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	DesktopName        *string                                                    `json:"DesktopName,omitempty" xml:"DesktopName,omitempty"`
	DesktopStatus      *string                                                    `json:"DesktopStatus,omitempty" xml:"DesktopStatus,omitempty"`
	DesktopTimers      []*DescribeGlobalDesktopsResponseBodyDesktopsDesktopTimers `json:"DesktopTimers,omitempty" xml:"DesktopTimers,omitempty" type:"Repeated"`
	DesktopType        *string                                                    `json:"DesktopType,omitempty" xml:"DesktopType,omitempty"`
	DirectoryId        *string                                                    `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	Disks              []*DescribeGlobalDesktopsResponseBodyDesktopsDisks         `json:"Disks,omitempty" xml:"Disks,omitempty" type:"Repeated"`
	EndUserId          *string                                                    `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	EndUserIds         []*string                                                  `json:"EndUserIds,omitempty" xml:"EndUserIds,omitempty" type:"Repeated"`
	ExpiredTime        *string                                                    `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	FotaUpdate         *DescribeGlobalDesktopsResponseBodyDesktopsFotaUpdate      `json:"FotaUpdate,omitempty" xml:"FotaUpdate,omitempty" type:"Struct"`
	GpuMemory          *int32                                                     `json:"GpuMemory,omitempty" xml:"GpuMemory,omitempty"`
	HibernationBeta    *bool                                                      `json:"HibernationBeta,omitempty" xml:"HibernationBeta,omitempty"`
	HostName           *string                                                    `json:"HostName,omitempty" xml:"HostName,omitempty"`
	ImageId            *string                                                    `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	LastStartTime      *string                                                    `json:"LastStartTime,omitempty" xml:"LastStartTime,omitempty"`
	LocalName          *string                                                    `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	ManagementFlags    []*string                                                  `json:"ManagementFlags,omitempty" xml:"ManagementFlags,omitempty" type:"Repeated"`
	Memory             *int64                                                     `json:"Memory,omitempty" xml:"Memory,omitempty"`
	NetworkInterfaceIp *string                                                    `json:"NetworkInterfaceIp,omitempty" xml:"NetworkInterfaceIp,omitempty"`
	OfficeSiteId       *string                                                    `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	Os                 *string                                                    `json:"Os,omitempty" xml:"Os,omitempty"`
	OsType             *string                                                    `json:"OsType,omitempty" xml:"OsType,omitempty"`
	Platform           *string                                                    `json:"Platform,omitempty" xml:"Platform,omitempty"`
	PolicyGroupId      *string                                                    `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty"`
	ProtocolType       *string                                                    `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	RealDesktopId      *string                                                    `json:"RealDesktopId,omitempty" xml:"RealDesktopId,omitempty"`
	RegionId           *string                                                    `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SessionType        *string                                                    `json:"SessionType,omitempty" xml:"SessionType,omitempty"`
	Sessions           []*DescribeGlobalDesktopsResponseBodyDesktopsSessions      `json:"Sessions,omitempty" xml:"Sessions,omitempty" type:"Repeated"`
	SupportHibernation *bool                                                      `json:"SupportHibernation,omitempty" xml:"SupportHibernation,omitempty"`
	UserCustomName     *string                                                    `json:"UserCustomName,omitempty" xml:"UserCustomName,omitempty"`
}

func (s DescribeGlobalDesktopsResponseBodyDesktops) String() string {
	return tea.Prettify(s)
}

func (s DescribeGlobalDesktopsResponseBodyDesktops) GoString() string {
	return s.String()
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) SetChargeType(v string) *DescribeGlobalDesktopsResponseBodyDesktops {
	s.ChargeType = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) SetClients(v []*DescribeGlobalDesktopsResponseBodyDesktopsClients) *DescribeGlobalDesktopsResponseBodyDesktops {
	s.Clients = v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) SetConnectionStatus(v string) *DescribeGlobalDesktopsResponseBodyDesktops {
	s.ConnectionStatus = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) SetCpu(v int32) *DescribeGlobalDesktopsResponseBodyDesktops {
	s.Cpu = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) SetCreationTime(v string) *DescribeGlobalDesktopsResponseBodyDesktops {
	s.CreationTime = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) SetDesktopGroupId(v string) *DescribeGlobalDesktopsResponseBodyDesktops {
	s.DesktopGroupId = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) SetDesktopId(v string) *DescribeGlobalDesktopsResponseBodyDesktops {
	s.DesktopId = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) SetDesktopName(v string) *DescribeGlobalDesktopsResponseBodyDesktops {
	s.DesktopName = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) SetDesktopStatus(v string) *DescribeGlobalDesktopsResponseBodyDesktops {
	s.DesktopStatus = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) SetDesktopTimers(v []*DescribeGlobalDesktopsResponseBodyDesktopsDesktopTimers) *DescribeGlobalDesktopsResponseBodyDesktops {
	s.DesktopTimers = v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) SetDesktopType(v string) *DescribeGlobalDesktopsResponseBodyDesktops {
	s.DesktopType = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) SetDirectoryId(v string) *DescribeGlobalDesktopsResponseBodyDesktops {
	s.DirectoryId = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) SetDisks(v []*DescribeGlobalDesktopsResponseBodyDesktopsDisks) *DescribeGlobalDesktopsResponseBodyDesktops {
	s.Disks = v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) SetEndUserId(v string) *DescribeGlobalDesktopsResponseBodyDesktops {
	s.EndUserId = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) SetEndUserIds(v []*string) *DescribeGlobalDesktopsResponseBodyDesktops {
	s.EndUserIds = v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) SetExpiredTime(v string) *DescribeGlobalDesktopsResponseBodyDesktops {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) SetFotaUpdate(v *DescribeGlobalDesktopsResponseBodyDesktopsFotaUpdate) *DescribeGlobalDesktopsResponseBodyDesktops {
	s.FotaUpdate = v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) SetGpuMemory(v int32) *DescribeGlobalDesktopsResponseBodyDesktops {
	s.GpuMemory = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) SetHibernationBeta(v bool) *DescribeGlobalDesktopsResponseBodyDesktops {
	s.HibernationBeta = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) SetHostName(v string) *DescribeGlobalDesktopsResponseBodyDesktops {
	s.HostName = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) SetImageId(v string) *DescribeGlobalDesktopsResponseBodyDesktops {
	s.ImageId = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) SetLastStartTime(v string) *DescribeGlobalDesktopsResponseBodyDesktops {
	s.LastStartTime = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) SetLocalName(v string) *DescribeGlobalDesktopsResponseBodyDesktops {
	s.LocalName = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) SetManagementFlags(v []*string) *DescribeGlobalDesktopsResponseBodyDesktops {
	s.ManagementFlags = v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) SetMemory(v int64) *DescribeGlobalDesktopsResponseBodyDesktops {
	s.Memory = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) SetNetworkInterfaceIp(v string) *DescribeGlobalDesktopsResponseBodyDesktops {
	s.NetworkInterfaceIp = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) SetOfficeSiteId(v string) *DescribeGlobalDesktopsResponseBodyDesktops {
	s.OfficeSiteId = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) SetOs(v string) *DescribeGlobalDesktopsResponseBodyDesktops {
	s.Os = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) SetOsType(v string) *DescribeGlobalDesktopsResponseBodyDesktops {
	s.OsType = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) SetPlatform(v string) *DescribeGlobalDesktopsResponseBodyDesktops {
	s.Platform = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) SetPolicyGroupId(v string) *DescribeGlobalDesktopsResponseBodyDesktops {
	s.PolicyGroupId = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) SetProtocolType(v string) *DescribeGlobalDesktopsResponseBodyDesktops {
	s.ProtocolType = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) SetRealDesktopId(v string) *DescribeGlobalDesktopsResponseBodyDesktops {
	s.RealDesktopId = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) SetRegionId(v string) *DescribeGlobalDesktopsResponseBodyDesktops {
	s.RegionId = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) SetSessionType(v string) *DescribeGlobalDesktopsResponseBodyDesktops {
	s.SessionType = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) SetSessions(v []*DescribeGlobalDesktopsResponseBodyDesktopsSessions) *DescribeGlobalDesktopsResponseBodyDesktops {
	s.Sessions = v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) SetSupportHibernation(v bool) *DescribeGlobalDesktopsResponseBodyDesktops {
	s.SupportHibernation = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktops) SetUserCustomName(v string) *DescribeGlobalDesktopsResponseBodyDesktops {
	s.UserCustomName = &v
	return s
}

type DescribeGlobalDesktopsResponseBodyDesktopsClients struct {
	// 客户端类型，取值：
	//
	// - macos：Mac客户端
	// - ios：IOS客户端
	// - android：Android客户端
	// - html5：Web客户端
	// - windows：Windows客户端
	// - linux：Linux客户端
	ClientType *string `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	// 客户端状态，取值：
	//
	// - ON：允许登录
	// - OFF：不允许登录
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeGlobalDesktopsResponseBodyDesktopsClients) String() string {
	return tea.Prettify(s)
}

func (s DescribeGlobalDesktopsResponseBodyDesktopsClients) GoString() string {
	return s.String()
}

func (s *DescribeGlobalDesktopsResponseBodyDesktopsClients) SetClientType(v string) *DescribeGlobalDesktopsResponseBodyDesktopsClients {
	s.ClientType = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktopsClients) SetStatus(v string) *DescribeGlobalDesktopsResponseBodyDesktopsClients {
	s.Status = &v
	return s
}

type DescribeGlobalDesktopsResponseBodyDesktopsDesktopTimers struct {
	AllowClientSetting *bool   `json:"AllowClientSetting,omitempty" xml:"AllowClientSetting,omitempty"`
	CronExpression     *string `json:"CronExpression,omitempty" xml:"CronExpression,omitempty"`
	Enforce            *bool   `json:"Enforce,omitempty" xml:"Enforce,omitempty"`
	ExecutionTime      *string `json:"ExecutionTime,omitempty" xml:"ExecutionTime,omitempty"`
	Interval           *int32  `json:"Interval,omitempty" xml:"Interval,omitempty"`
	OperationType      *string `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	ResetType          *string `json:"ResetType,omitempty" xml:"ResetType,omitempty"`
	TimerType          *string `json:"TimerType,omitempty" xml:"TimerType,omitempty"`
}

func (s DescribeGlobalDesktopsResponseBodyDesktopsDesktopTimers) String() string {
	return tea.Prettify(s)
}

func (s DescribeGlobalDesktopsResponseBodyDesktopsDesktopTimers) GoString() string {
	return s.String()
}

func (s *DescribeGlobalDesktopsResponseBodyDesktopsDesktopTimers) SetAllowClientSetting(v bool) *DescribeGlobalDesktopsResponseBodyDesktopsDesktopTimers {
	s.AllowClientSetting = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktopsDesktopTimers) SetCronExpression(v string) *DescribeGlobalDesktopsResponseBodyDesktopsDesktopTimers {
	s.CronExpression = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktopsDesktopTimers) SetEnforce(v bool) *DescribeGlobalDesktopsResponseBodyDesktopsDesktopTimers {
	s.Enforce = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktopsDesktopTimers) SetExecutionTime(v string) *DescribeGlobalDesktopsResponseBodyDesktopsDesktopTimers {
	s.ExecutionTime = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktopsDesktopTimers) SetInterval(v int32) *DescribeGlobalDesktopsResponseBodyDesktopsDesktopTimers {
	s.Interval = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktopsDesktopTimers) SetOperationType(v string) *DescribeGlobalDesktopsResponseBodyDesktopsDesktopTimers {
	s.OperationType = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktopsDesktopTimers) SetResetType(v string) *DescribeGlobalDesktopsResponseBodyDesktopsDesktopTimers {
	s.ResetType = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktopsDesktopTimers) SetTimerType(v string) *DescribeGlobalDesktopsResponseBodyDesktopsDesktopTimers {
	s.TimerType = &v
	return s
}

type DescribeGlobalDesktopsResponseBodyDesktopsDisks struct {
	DiskId   *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	DiskSize *int32  `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
	DiskType *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
}

func (s DescribeGlobalDesktopsResponseBodyDesktopsDisks) String() string {
	return tea.Prettify(s)
}

func (s DescribeGlobalDesktopsResponseBodyDesktopsDisks) GoString() string {
	return s.String()
}

func (s *DescribeGlobalDesktopsResponseBodyDesktopsDisks) SetDiskId(v string) *DescribeGlobalDesktopsResponseBodyDesktopsDisks {
	s.DiskId = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktopsDisks) SetDiskSize(v int32) *DescribeGlobalDesktopsResponseBodyDesktopsDisks {
	s.DiskSize = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktopsDisks) SetDiskType(v string) *DescribeGlobalDesktopsResponseBodyDesktopsDisks {
	s.DiskType = &v
	return s
}

type DescribeGlobalDesktopsResponseBodyDesktopsFotaUpdate struct {
	Channel           *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
	CurrentAppVersion *string `json:"CurrentAppVersion,omitempty" xml:"CurrentAppVersion,omitempty"`
	Force             *bool   `json:"Force,omitempty" xml:"Force,omitempty"`
	NewAppVersion     *string `json:"NewAppVersion,omitempty" xml:"NewAppVersion,omitempty"`
	Project           *string `json:"Project,omitempty" xml:"Project,omitempty"`
	ReleaseNote       *string `json:"ReleaseNote,omitempty" xml:"ReleaseNote,omitempty"`
	ReleaseNoteEn     *string `json:"ReleaseNoteEn,omitempty" xml:"ReleaseNoteEn,omitempty"`
	ReleaseNoteJp     *string `json:"ReleaseNoteJp,omitempty" xml:"ReleaseNoteJp,omitempty"`
	Size              *string `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s DescribeGlobalDesktopsResponseBodyDesktopsFotaUpdate) String() string {
	return tea.Prettify(s)
}

func (s DescribeGlobalDesktopsResponseBodyDesktopsFotaUpdate) GoString() string {
	return s.String()
}

func (s *DescribeGlobalDesktopsResponseBodyDesktopsFotaUpdate) SetChannel(v string) *DescribeGlobalDesktopsResponseBodyDesktopsFotaUpdate {
	s.Channel = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktopsFotaUpdate) SetCurrentAppVersion(v string) *DescribeGlobalDesktopsResponseBodyDesktopsFotaUpdate {
	s.CurrentAppVersion = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktopsFotaUpdate) SetForce(v bool) *DescribeGlobalDesktopsResponseBodyDesktopsFotaUpdate {
	s.Force = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktopsFotaUpdate) SetNewAppVersion(v string) *DescribeGlobalDesktopsResponseBodyDesktopsFotaUpdate {
	s.NewAppVersion = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktopsFotaUpdate) SetProject(v string) *DescribeGlobalDesktopsResponseBodyDesktopsFotaUpdate {
	s.Project = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktopsFotaUpdate) SetReleaseNote(v string) *DescribeGlobalDesktopsResponseBodyDesktopsFotaUpdate {
	s.ReleaseNote = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktopsFotaUpdate) SetReleaseNoteEn(v string) *DescribeGlobalDesktopsResponseBodyDesktopsFotaUpdate {
	s.ReleaseNoteEn = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktopsFotaUpdate) SetReleaseNoteJp(v string) *DescribeGlobalDesktopsResponseBodyDesktopsFotaUpdate {
	s.ReleaseNoteJp = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktopsFotaUpdate) SetSize(v string) *DescribeGlobalDesktopsResponseBodyDesktopsFotaUpdate {
	s.Size = &v
	return s
}

type DescribeGlobalDesktopsResponseBodyDesktopsSessions struct {
	EndUserId         *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	EstablishmentTime *string `json:"EstablishmentTime,omitempty" xml:"EstablishmentTime,omitempty"`
}

func (s DescribeGlobalDesktopsResponseBodyDesktopsSessions) String() string {
	return tea.Prettify(s)
}

func (s DescribeGlobalDesktopsResponseBodyDesktopsSessions) GoString() string {
	return s.String()
}

func (s *DescribeGlobalDesktopsResponseBodyDesktopsSessions) SetEndUserId(v string) *DescribeGlobalDesktopsResponseBodyDesktopsSessions {
	s.EndUserId = &v
	return s
}

func (s *DescribeGlobalDesktopsResponseBodyDesktopsSessions) SetEstablishmentTime(v string) *DescribeGlobalDesktopsResponseBodyDesktopsSessions {
	s.EstablishmentTime = &v
	return s
}

type DescribeGlobalDesktopsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeGlobalDesktopsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeGlobalDesktopsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeGlobalDesktopsResponse) GoString() string {
	return s.String()
}

func (s *DescribeGlobalDesktopsResponse) SetHeaders(v map[string]*string) *DescribeGlobalDesktopsResponse {
	s.Headers = v
	return s
}

func (s *DescribeGlobalDesktopsResponse) SetStatusCode(v int32) *DescribeGlobalDesktopsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGlobalDesktopsResponse) SetBody(v *DescribeGlobalDesktopsResponseBody) *DescribeGlobalDesktopsResponse {
	s.Body = v
	return s
}

type DescribeOfficeSitesRequest struct {
	ClientId     *string   `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	OfficeSiteId []*string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty" type:"Repeated"`
	RegionId     *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeOfficeSitesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeOfficeSitesRequest) GoString() string {
	return s.String()
}

func (s *DescribeOfficeSitesRequest) SetClientId(v string) *DescribeOfficeSitesRequest {
	s.ClientId = &v
	return s
}

func (s *DescribeOfficeSitesRequest) SetOfficeSiteId(v []*string) *DescribeOfficeSitesRequest {
	s.OfficeSiteId = v
	return s
}

func (s *DescribeOfficeSitesRequest) SetRegionId(v string) *DescribeOfficeSitesRequest {
	s.RegionId = &v
	return s
}

type DescribeOfficeSitesResponseBody struct {
	OfficeSites []*DescribeOfficeSitesResponseBodyOfficeSites `json:"OfficeSites,omitempty" xml:"OfficeSites,omitempty" type:"Repeated"`
	RequestId   *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeOfficeSitesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeOfficeSitesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOfficeSitesResponseBody) SetOfficeSites(v []*DescribeOfficeSitesResponseBodyOfficeSites) *DescribeOfficeSitesResponseBody {
	s.OfficeSites = v
	return s
}

func (s *DescribeOfficeSitesResponseBody) SetRequestId(v string) *DescribeOfficeSitesResponseBody {
	s.RequestId = &v
	return s
}

type DescribeOfficeSitesResponseBodyOfficeSites struct {
	DesktopAccessType  *string `json:"DesktopAccessType,omitempty" xml:"DesktopAccessType,omitempty"`
	DesktopVpcEndpoint *string `json:"DesktopVpcEndpoint,omitempty" xml:"DesktopVpcEndpoint,omitempty"`
	OfficeSiteId       *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	OfficeSiteType     *string `json:"OfficeSiteType,omitempty" xml:"OfficeSiteType,omitempty"`
	ProviderId         *string `json:"ProviderId,omitempty" xml:"ProviderId,omitempty"`
	SsoServiceUrl      *string `json:"SsoServiceUrl,omitempty" xml:"SsoServiceUrl,omitempty"`
}

func (s DescribeOfficeSitesResponseBodyOfficeSites) String() string {
	return tea.Prettify(s)
}

func (s DescribeOfficeSitesResponseBodyOfficeSites) GoString() string {
	return s.String()
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetDesktopAccessType(v string) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.DesktopAccessType = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetDesktopVpcEndpoint(v string) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.DesktopVpcEndpoint = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetOfficeSiteId(v string) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.OfficeSiteId = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetOfficeSiteType(v string) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.OfficeSiteType = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetProviderId(v string) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.ProviderId = &v
	return s
}

func (s *DescribeOfficeSitesResponseBodyOfficeSites) SetSsoServiceUrl(v string) *DescribeOfficeSitesResponseBodyOfficeSites {
	s.SsoServiceUrl = &v
	return s
}

type DescribeOfficeSitesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeOfficeSitesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeOfficeSitesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeOfficeSitesResponse) GoString() string {
	return s.String()
}

func (s *DescribeOfficeSitesResponse) SetHeaders(v map[string]*string) *DescribeOfficeSitesResponse {
	s.Headers = v
	return s
}

func (s *DescribeOfficeSitesResponse) SetStatusCode(v int32) *DescribeOfficeSitesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeOfficeSitesResponse) SetBody(v *DescribeOfficeSitesResponseBody) *DescribeOfficeSitesResponse {
	s.Body = v
	return s
}

type DescribeRegionsRequest struct {
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeRegionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRegionsRequest) SetClientId(v string) *DescribeRegionsRequest {
	s.ClientId = &v
	return s
}

func (s *DescribeRegionsRequest) SetRegionId(v string) *DescribeRegionsRequest {
	s.RegionId = &v
	return s
}

type DescribeRegionsResponseBody struct {
	Regions   []*DescribeRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBody) SetRegions(v []*DescribeRegionsResponseBodyRegions) *DescribeRegionsResponseBody {
	s.Regions = v
	return s
}

func (s *DescribeRegionsResponseBody) SetRequestId(v string) *DescribeRegionsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeRegionsResponseBodyRegions struct {
	RegionEndpoint *string `json:"RegionEndpoint,omitempty" xml:"RegionEndpoint,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeRegionsResponseBodyRegions) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegions) SetRegionEndpoint(v string) *DescribeRegionsResponseBodyRegions {
	s.RegionEndpoint = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegions) SetRegionId(v string) *DescribeRegionsResponseBodyRegions {
	s.RegionId = &v
	return s
}

type DescribeRegionsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRegionsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponse) SetHeaders(v map[string]*string) *DescribeRegionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeRegionsResponse) SetStatusCode(v int32) *DescribeRegionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRegionsResponse) SetBody(v *DescribeRegionsResponseBody) *DescribeRegionsResponse {
	s.Body = v
	return s
}

type DescribeSnapshotsRequest struct {
	ClientId   *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	DesktopId  *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	LoginToken *string `json:"LoginToken,omitempty" xml:"LoginToken,omitempty"`
	MaxResults *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SessionId  *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
}

func (s DescribeSnapshotsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSnapshotsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSnapshotsRequest) SetClientId(v string) *DescribeSnapshotsRequest {
	s.ClientId = &v
	return s
}

func (s *DescribeSnapshotsRequest) SetDesktopId(v string) *DescribeSnapshotsRequest {
	s.DesktopId = &v
	return s
}

func (s *DescribeSnapshotsRequest) SetLoginToken(v string) *DescribeSnapshotsRequest {
	s.LoginToken = &v
	return s
}

func (s *DescribeSnapshotsRequest) SetMaxResults(v int32) *DescribeSnapshotsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeSnapshotsRequest) SetNextToken(v string) *DescribeSnapshotsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeSnapshotsRequest) SetRegionId(v string) *DescribeSnapshotsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSnapshotsRequest) SetSessionId(v string) *DescribeSnapshotsRequest {
	s.SessionId = &v
	return s
}

func (s *DescribeSnapshotsRequest) SetSnapshotId(v string) *DescribeSnapshotsRequest {
	s.SnapshotId = &v
	return s
}

type DescribeSnapshotsResponseBody struct {
	NextToken *string                                   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Snapshots []*DescribeSnapshotsResponseBodySnapshots `json:"Snapshots,omitempty" xml:"Snapshots,omitempty" type:"Repeated"`
}

func (s DescribeSnapshotsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSnapshotsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSnapshotsResponseBody) SetNextToken(v string) *DescribeSnapshotsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeSnapshotsResponseBody) SetRequestId(v string) *DescribeSnapshotsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSnapshotsResponseBody) SetSnapshots(v []*DescribeSnapshotsResponseBodySnapshots) *DescribeSnapshotsResponseBody {
	s.Snapshots = v
	return s
}

type DescribeSnapshotsResponseBodySnapshots struct {
	CreationTime   *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DesktopId      *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	Progress       *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	RemainTime     *int32  `json:"RemainTime,omitempty" xml:"RemainTime,omitempty"`
	SnapshotId     *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	SnapshotName   *string `json:"SnapshotName,omitempty" xml:"SnapshotName,omitempty"`
	SnapshotType   *string `json:"SnapshotType,omitempty" xml:"SnapshotType,omitempty"`
	SourceDiskSize *string `json:"SourceDiskSize,omitempty" xml:"SourceDiskSize,omitempty"`
	SourceDiskType *string `json:"SourceDiskType,omitempty" xml:"SourceDiskType,omitempty"`
	Status         *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeSnapshotsResponseBodySnapshots) String() string {
	return tea.Prettify(s)
}

func (s DescribeSnapshotsResponseBodySnapshots) GoString() string {
	return s.String()
}

func (s *DescribeSnapshotsResponseBodySnapshots) SetCreationTime(v string) *DescribeSnapshotsResponseBodySnapshots {
	s.CreationTime = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshots) SetDescription(v string) *DescribeSnapshotsResponseBodySnapshots {
	s.Description = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshots) SetDesktopId(v string) *DescribeSnapshotsResponseBodySnapshots {
	s.DesktopId = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshots) SetProgress(v string) *DescribeSnapshotsResponseBodySnapshots {
	s.Progress = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshots) SetRemainTime(v int32) *DescribeSnapshotsResponseBodySnapshots {
	s.RemainTime = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshots) SetSnapshotId(v string) *DescribeSnapshotsResponseBodySnapshots {
	s.SnapshotId = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshots) SetSnapshotName(v string) *DescribeSnapshotsResponseBodySnapshots {
	s.SnapshotName = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshots) SetSnapshotType(v string) *DescribeSnapshotsResponseBodySnapshots {
	s.SnapshotType = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshots) SetSourceDiskSize(v string) *DescribeSnapshotsResponseBodySnapshots {
	s.SourceDiskSize = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshots) SetSourceDiskType(v string) *DescribeSnapshotsResponseBodySnapshots {
	s.SourceDiskType = &v
	return s
}

func (s *DescribeSnapshotsResponseBodySnapshots) SetStatus(v string) *DescribeSnapshotsResponseBodySnapshots {
	s.Status = &v
	return s
}

type DescribeSnapshotsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSnapshotsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSnapshotsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSnapshotsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSnapshotsResponse) SetHeaders(v map[string]*string) *DescribeSnapshotsResponse {
	s.Headers = v
	return s
}

func (s *DescribeSnapshotsResponse) SetStatusCode(v int32) *DescribeSnapshotsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSnapshotsResponse) SetBody(v *DescribeSnapshotsResponseBody) *DescribeSnapshotsResponse {
	s.Body = v
	return s
}

type EncryptPasswordRequest struct {
	ClientId     *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	DirectoryId  *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	LoginToken   *string `json:"LoginToken,omitempty" xml:"LoginToken,omitempty"`
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	Password     *string `json:"Password,omitempty" xml:"Password,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SessionId    *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s EncryptPasswordRequest) String() string {
	return tea.Prettify(s)
}

func (s EncryptPasswordRequest) GoString() string {
	return s.String()
}

func (s *EncryptPasswordRequest) SetClientId(v string) *EncryptPasswordRequest {
	s.ClientId = &v
	return s
}

func (s *EncryptPasswordRequest) SetDirectoryId(v string) *EncryptPasswordRequest {
	s.DirectoryId = &v
	return s
}

func (s *EncryptPasswordRequest) SetLoginToken(v string) *EncryptPasswordRequest {
	s.LoginToken = &v
	return s
}

func (s *EncryptPasswordRequest) SetOfficeSiteId(v string) *EncryptPasswordRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *EncryptPasswordRequest) SetPassword(v string) *EncryptPasswordRequest {
	s.Password = &v
	return s
}

func (s *EncryptPasswordRequest) SetRegionId(v string) *EncryptPasswordRequest {
	s.RegionId = &v
	return s
}

func (s *EncryptPasswordRequest) SetSessionId(v string) *EncryptPasswordRequest {
	s.SessionId = &v
	return s
}

type EncryptPasswordResponseBody struct {
	EncryptedPassword *string `json:"EncryptedPassword,omitempty" xml:"EncryptedPassword,omitempty"`
	RequestId         *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EncryptPasswordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EncryptPasswordResponseBody) GoString() string {
	return s.String()
}

func (s *EncryptPasswordResponseBody) SetEncryptedPassword(v string) *EncryptPasswordResponseBody {
	s.EncryptedPassword = &v
	return s
}

func (s *EncryptPasswordResponseBody) SetRequestId(v string) *EncryptPasswordResponseBody {
	s.RequestId = &v
	return s
}

type EncryptPasswordResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EncryptPasswordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EncryptPasswordResponse) String() string {
	return tea.Prettify(s)
}

func (s EncryptPasswordResponse) GoString() string {
	return s.String()
}

func (s *EncryptPasswordResponse) SetHeaders(v map[string]*string) *EncryptPasswordResponse {
	s.Headers = v
	return s
}

func (s *EncryptPasswordResponse) SetStatusCode(v int32) *EncryptPasswordResponse {
	s.StatusCode = &v
	return s
}

func (s *EncryptPasswordResponse) SetBody(v *EncryptPasswordResponseBody) *EncryptPasswordResponse {
	s.Body = v
	return s
}

type GetCloudDriveServiceMountTokenRequest struct {
	ClientId     *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	LoginToken   *string `json:"LoginToken,omitempty" xml:"LoginToken,omitempty"`
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SessionId    *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s GetCloudDriveServiceMountTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCloudDriveServiceMountTokenRequest) GoString() string {
	return s.String()
}

func (s *GetCloudDriveServiceMountTokenRequest) SetClientId(v string) *GetCloudDriveServiceMountTokenRequest {
	s.ClientId = &v
	return s
}

func (s *GetCloudDriveServiceMountTokenRequest) SetLoginToken(v string) *GetCloudDriveServiceMountTokenRequest {
	s.LoginToken = &v
	return s
}

func (s *GetCloudDriveServiceMountTokenRequest) SetOfficeSiteId(v string) *GetCloudDriveServiceMountTokenRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *GetCloudDriveServiceMountTokenRequest) SetRegionId(v string) *GetCloudDriveServiceMountTokenRequest {
	s.RegionId = &v
	return s
}

func (s *GetCloudDriveServiceMountTokenRequest) SetSessionId(v string) *GetCloudDriveServiceMountTokenRequest {
	s.SessionId = &v
	return s
}

type GetCloudDriveServiceMountTokenResponseBody struct {
	RequestId *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Token     *GetCloudDriveServiceMountTokenResponseBodyToken `json:"Token,omitempty" xml:"Token,omitempty" type:"Struct"`
}

func (s GetCloudDriveServiceMountTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCloudDriveServiceMountTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GetCloudDriveServiceMountTokenResponseBody) SetRequestId(v string) *GetCloudDriveServiceMountTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCloudDriveServiceMountTokenResponseBody) SetToken(v *GetCloudDriveServiceMountTokenResponseBodyToken) *GetCloudDriveServiceMountTokenResponseBody {
	s.Token = v
	return s
}

type GetCloudDriveServiceMountTokenResponseBodyToken struct {
	DomainId     *string `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
	ExpiredAfter *string `json:"ExpiredAfter,omitempty" xml:"ExpiredAfter,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Token        *string `json:"Token,omitempty" xml:"Token,omitempty"`
	TotalSize    *int64  `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
	UsedSize     *int64  `json:"UsedSize,omitempty" xml:"UsedSize,omitempty"`
}

func (s GetCloudDriveServiceMountTokenResponseBodyToken) String() string {
	return tea.Prettify(s)
}

func (s GetCloudDriveServiceMountTokenResponseBodyToken) GoString() string {
	return s.String()
}

func (s *GetCloudDriveServiceMountTokenResponseBodyToken) SetDomainId(v string) *GetCloudDriveServiceMountTokenResponseBodyToken {
	s.DomainId = &v
	return s
}

func (s *GetCloudDriveServiceMountTokenResponseBodyToken) SetExpiredAfter(v string) *GetCloudDriveServiceMountTokenResponseBodyToken {
	s.ExpiredAfter = &v
	return s
}

func (s *GetCloudDriveServiceMountTokenResponseBodyToken) SetStatus(v string) *GetCloudDriveServiceMountTokenResponseBodyToken {
	s.Status = &v
	return s
}

func (s *GetCloudDriveServiceMountTokenResponseBodyToken) SetToken(v string) *GetCloudDriveServiceMountTokenResponseBodyToken {
	s.Token = &v
	return s
}

func (s *GetCloudDriveServiceMountTokenResponseBodyToken) SetTotalSize(v int64) *GetCloudDriveServiceMountTokenResponseBodyToken {
	s.TotalSize = &v
	return s
}

func (s *GetCloudDriveServiceMountTokenResponseBodyToken) SetUsedSize(v int64) *GetCloudDriveServiceMountTokenResponseBodyToken {
	s.UsedSize = &v
	return s
}

type GetCloudDriveServiceMountTokenResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCloudDriveServiceMountTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCloudDriveServiceMountTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCloudDriveServiceMountTokenResponse) GoString() string {
	return s.String()
}

func (s *GetCloudDriveServiceMountTokenResponse) SetHeaders(v map[string]*string) *GetCloudDriveServiceMountTokenResponse {
	s.Headers = v
	return s
}

func (s *GetCloudDriveServiceMountTokenResponse) SetStatusCode(v int32) *GetCloudDriveServiceMountTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCloudDriveServiceMountTokenResponse) SetBody(v *GetCloudDriveServiceMountTokenResponseBody) *GetCloudDriveServiceMountTokenResponse {
	s.Body = v
	return s
}

type GetConnectionTicketRequest struct {
	ClientId             *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	ClientOS             *string `json:"ClientOS,omitempty" xml:"ClientOS,omitempty"`
	ClientType           *string `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	ClientVersion        *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	CommandContent       *string `json:"CommandContent,omitempty" xml:"CommandContent,omitempty"`
	DesktopId            *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	LoginToken           *string `json:"LoginToken,omitempty" xml:"LoginToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SessionId            *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	TaskId               *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	Uuid                 *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s GetConnectionTicketRequest) String() string {
	return tea.Prettify(s)
}

func (s GetConnectionTicketRequest) GoString() string {
	return s.String()
}

func (s *GetConnectionTicketRequest) SetClientId(v string) *GetConnectionTicketRequest {
	s.ClientId = &v
	return s
}

func (s *GetConnectionTicketRequest) SetClientOS(v string) *GetConnectionTicketRequest {
	s.ClientOS = &v
	return s
}

func (s *GetConnectionTicketRequest) SetClientType(v string) *GetConnectionTicketRequest {
	s.ClientType = &v
	return s
}

func (s *GetConnectionTicketRequest) SetClientVersion(v string) *GetConnectionTicketRequest {
	s.ClientVersion = &v
	return s
}

func (s *GetConnectionTicketRequest) SetCommandContent(v string) *GetConnectionTicketRequest {
	s.CommandContent = &v
	return s
}

func (s *GetConnectionTicketRequest) SetDesktopId(v string) *GetConnectionTicketRequest {
	s.DesktopId = &v
	return s
}

func (s *GetConnectionTicketRequest) SetLoginToken(v string) *GetConnectionTicketRequest {
	s.LoginToken = &v
	return s
}

func (s *GetConnectionTicketRequest) SetOwnerId(v int64) *GetConnectionTicketRequest {
	s.OwnerId = &v
	return s
}

func (s *GetConnectionTicketRequest) SetRegionId(v string) *GetConnectionTicketRequest {
	s.RegionId = &v
	return s
}

func (s *GetConnectionTicketRequest) SetResourceOwnerAccount(v string) *GetConnectionTicketRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetConnectionTicketRequest) SetResourceOwnerId(v int64) *GetConnectionTicketRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetConnectionTicketRequest) SetSessionId(v string) *GetConnectionTicketRequest {
	s.SessionId = &v
	return s
}

func (s *GetConnectionTicketRequest) SetTaskId(v string) *GetConnectionTicketRequest {
	s.TaskId = &v
	return s
}

func (s *GetConnectionTicketRequest) SetUuid(v string) *GetConnectionTicketRequest {
	s.Uuid = &v
	return s
}

type GetConnectionTicketResponseBody struct {
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskCode    *string `json:"TaskCode,omitempty" xml:"TaskCode,omitempty"`
	TaskId      *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskMessage *string `json:"TaskMessage,omitempty" xml:"TaskMessage,omitempty"`
	TaskStatus  *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	Ticket      *string `json:"Ticket,omitempty" xml:"Ticket,omitempty"`
}

func (s GetConnectionTicketResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetConnectionTicketResponseBody) GoString() string {
	return s.String()
}

func (s *GetConnectionTicketResponseBody) SetRequestId(v string) *GetConnectionTicketResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetConnectionTicketResponseBody) SetTaskCode(v string) *GetConnectionTicketResponseBody {
	s.TaskCode = &v
	return s
}

func (s *GetConnectionTicketResponseBody) SetTaskId(v string) *GetConnectionTicketResponseBody {
	s.TaskId = &v
	return s
}

func (s *GetConnectionTicketResponseBody) SetTaskMessage(v string) *GetConnectionTicketResponseBody {
	s.TaskMessage = &v
	return s
}

func (s *GetConnectionTicketResponseBody) SetTaskStatus(v string) *GetConnectionTicketResponseBody {
	s.TaskStatus = &v
	return s
}

func (s *GetConnectionTicketResponseBody) SetTicket(v string) *GetConnectionTicketResponseBody {
	s.Ticket = &v
	return s
}

type GetConnectionTicketResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetConnectionTicketResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetConnectionTicketResponse) String() string {
	return tea.Prettify(s)
}

func (s GetConnectionTicketResponse) GoString() string {
	return s.String()
}

func (s *GetConnectionTicketResponse) SetHeaders(v map[string]*string) *GetConnectionTicketResponse {
	s.Headers = v
	return s
}

func (s *GetConnectionTicketResponse) SetStatusCode(v int32) *GetConnectionTicketResponse {
	s.StatusCode = &v
	return s
}

func (s *GetConnectionTicketResponse) SetBody(v *GetConnectionTicketResponseBody) *GetConnectionTicketResponse {
	s.Body = v
	return s
}

type GetLoginTokenRequest struct {
	// The verification code that is generated by the virtual MFA device. This parameter is required if you set `CurrentStage` to `MFAVerify`.
	AuthenticationCode *string `json:"AuthenticationCode,omitempty" xml:"AuthenticationCode,omitempty"`
	// The ID of the Alibaba Cloud Workspace client. The system generates a unique ID for each client.
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The OS that the client runs.
	ClientOS *string `json:"ClientOS,omitempty" xml:"ClientOS,omitempty"`
	// The type of the software client.
	ClientType *string `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	// The version of the client. When you use an Alibaba Cloud Workspace client, you can view the client version in the **About** dialog box on the client logon page.
	ClientVersion *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	// The logon authentication stage. Valid values:
	//
	// *   `ADPassword`: the stage to verify the identity of the Active Directory (AD) user. You must specify the value when the system verifies the identity of a convenience account or an AD account.
	// *   `MFABind`: the stage to bind a virtual multi-factor authentication (MFA) device.
	// *   `MFAVerify`: the stage to obtain the verification code that is generated by the virtual MFA device.
	// *   `TokenVerify`: the stage to perform two-factor authentication for the client.
	// *   `ChangePassword`: the stage to change the password of the user.
	// *   `VerifyKeepAlive`: the stage to exchange the logon credential. This parameter is valid if KeepAliveToken is valid.
	CurrentStage *string `json:"CurrentStage,omitempty" xml:"CurrentStage,omitempty"`
	// The ID of the workspace. The parameter is the same as the `OfficeSiteId` parameter. We recommend that you use `OfficeSiteId` instead of `DirectoryId`. You can specify a value for either the `DirectoryId` parameter or the `OfficeSiteId` parameter, but not both.
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The name of the convenience user or the AD user. This parameter is required if you set `CurrentStage` to `ADPassword`.
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// Specifies whether to keep the user logged on to the client.
	// Valid values:
	// * null: Default value. Do not keep the user logged on to the client.
	// * true: Keep the user logged on to the client.
	// * false:  Do not keep the user logged on to the client.
	KeepAlive *bool `json:"KeepAlive,omitempty" xml:"KeepAlive,omitempty"`
	// The token that is used to keep the user logged on to the client. After the user logs on to the client and KeepAlive is set to true, the `KeepAliveToken` is returned. You can call the `GetLoginToken` operation within the valid duration``, and set `CurrentStage` to `VerifyKeepAlive` to obtain the logon token (LoginToken). This parameter is required if you set `CurrentStage` to `VerifyKeepAlive```.
	KeepAliveToken *string `json:"KeepAliveToken,omitempty" xml:"KeepAliveToken,omitempty"`
	// The new password. This parameter is required if you set `CurrentStage` to `ChangePassword`.
	NewPassword *string `json:"NewPassword,omitempty" xml:"NewPassword,omitempty"`
	// The ID of the workspace.
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// The current password. This parameter is required if you set `CurrentStage` to `ChangePassword`.
	OldPassword *string `json:"OldPassword,omitempty" xml:"OldPassword,omitempty"`
	// The password of the convenience user or the AD user. This parameter is required if you set `CurrentStage` to `ADPassword`.
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The ID of the region. You can call the [DescribeRegions](~~436773~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the session.
	//
	// *   If the virtual multi-factor authentication (MFA) device is not bound or two-factor authentication is not enabled for the client, you do not need to specify a value for `SessionId`.
	// *   If the virtual MFA device is not bound or two-factor authentication is enabled for the client, you must specify a value for `SessionId` to verify the user identity after you specify a value for `ADPassword`. The value of the `SessionId` parameter is returned only if the CurrentStage parameter is set to `ADPassword` when you call the `GetLoginToken` operation.
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// If two-factor authentication is enabled in the Elastic Desktop Service (EDS) console and the system detects that the identity of the logon user may have security risks, the system sends a verification code for two-factor authentication to the email address of the user. This parameter is required if you set `CurrentStage` to `TokenVerify`.
	TokenCode *string `json:"TokenCode,omitempty" xml:"TokenCode,omitempty"`
	// The unique identifier of the client. When you use an Alibaba Cloud Workspace client, you can view the client version in the **About** dialog box on the client logon page.
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s GetLoginTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s GetLoginTokenRequest) GoString() string {
	return s.String()
}

func (s *GetLoginTokenRequest) SetAuthenticationCode(v string) *GetLoginTokenRequest {
	s.AuthenticationCode = &v
	return s
}

func (s *GetLoginTokenRequest) SetClientId(v string) *GetLoginTokenRequest {
	s.ClientId = &v
	return s
}

func (s *GetLoginTokenRequest) SetClientOS(v string) *GetLoginTokenRequest {
	s.ClientOS = &v
	return s
}

func (s *GetLoginTokenRequest) SetClientType(v string) *GetLoginTokenRequest {
	s.ClientType = &v
	return s
}

func (s *GetLoginTokenRequest) SetClientVersion(v string) *GetLoginTokenRequest {
	s.ClientVersion = &v
	return s
}

func (s *GetLoginTokenRequest) SetCurrentStage(v string) *GetLoginTokenRequest {
	s.CurrentStage = &v
	return s
}

func (s *GetLoginTokenRequest) SetDirectoryId(v string) *GetLoginTokenRequest {
	s.DirectoryId = &v
	return s
}

func (s *GetLoginTokenRequest) SetEndUserId(v string) *GetLoginTokenRequest {
	s.EndUserId = &v
	return s
}

func (s *GetLoginTokenRequest) SetKeepAlive(v bool) *GetLoginTokenRequest {
	s.KeepAlive = &v
	return s
}

func (s *GetLoginTokenRequest) SetKeepAliveToken(v string) *GetLoginTokenRequest {
	s.KeepAliveToken = &v
	return s
}

func (s *GetLoginTokenRequest) SetNewPassword(v string) *GetLoginTokenRequest {
	s.NewPassword = &v
	return s
}

func (s *GetLoginTokenRequest) SetOfficeSiteId(v string) *GetLoginTokenRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *GetLoginTokenRequest) SetOldPassword(v string) *GetLoginTokenRequest {
	s.OldPassword = &v
	return s
}

func (s *GetLoginTokenRequest) SetPassword(v string) *GetLoginTokenRequest {
	s.Password = &v
	return s
}

func (s *GetLoginTokenRequest) SetRegionId(v string) *GetLoginTokenRequest {
	s.RegionId = &v
	return s
}

func (s *GetLoginTokenRequest) SetSessionId(v string) *GetLoginTokenRequest {
	s.SessionId = &v
	return s
}

func (s *GetLoginTokenRequest) SetTokenCode(v string) *GetLoginTokenRequest {
	s.TokenCode = &v
	return s
}

func (s *GetLoginTokenRequest) SetUuid(v string) *GetLoginTokenRequest {
	s.Uuid = &v
	return s
}

type GetLoginTokenResponseBody struct {
	// The email address of the user. The system returns the email address in the return value of the LoginToken parameter after the user logs on to the client.
	//
	// *   For a convenience user, the return value is the email address specified when the administrator creates the convenience user.
	// *   For an AD user, the return value is in the following format: `Username@Name of the AD domain`.
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The account of the convenience user or the AD user.
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// > This is a parameter only for internal use.
	Industry *string `json:"Industry,omitempty" xml:"Industry,omitempty"`
	// The token used to keep the user logged on. After the user logs on to the client and select the Keep Logon option, `KeepAliveToken` is returned when you call the operation. If the user does not select the Keep Logon option, null is returned.
	KeepAliveToken *string `json:"KeepAliveToken,omitempty" xml:"KeepAliveToken,omitempty"`
	// The attribute of the convenience user. For an AD user, null is returned.
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The logon token.
	LoginToken *string `json:"LoginToken,omitempty" xml:"LoginToken,omitempty"`
	// The next stage that is expected to enter. For example, if the administrator enables MFA authentication in the EDS console, `MFAVerify` is returned after the username and password pass the authentication (after you set CurrentStage to `ADPassword` stage). This indicates that the MFA authentication is required.
	//
	// > For more information about each authentication stage, see the parameter description of the request parameter `CurrentStage`.
	NextStage *string `json:"NextStage,omitempty" xml:"NextStage,omitempty"`
	// Enter the mobile number of the convenience user. For an AD user, null is returned.
	Phone *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
	// > This is a parameter only for internal use.
	Props map[string]*string `json:"Props,omitempty" xml:"Props,omitempty"`
	// The QR code that is generated when the virtual MFA device is bound. The value is encoded in Base64. This parameter can be empty. This parameter is required only when the CurrentStage parameter is set to `MFABind`.
	//
	// > For more information about each authentication stage, see the parameter description of the request parameter `CurrentStage`.
	QrCodePng *string `json:"QrCodePng,omitempty" xml:"QrCodePng,omitempty"`
	// The ID of the request.
	RequestId      *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RiskVerifyInfo *GetLoginTokenResponseBodyRiskVerifyInfo `json:"RiskVerifyInfo,omitempty" xml:"RiskVerifyInfo,omitempty" type:"Struct"`
	// The key that is generated when you bind the virtual MFA device. This parameter is required when the CurrentStage parameter is set to `MFABind`.
	//
	// > For more information about each authentication stage, see the parameter description of the request parameter `CurrentStage`.
	Secret *string `json:"Secret,omitempty" xml:"Secret,omitempty"`
	// The ID of the session. The ID is returned the first time you call the `GetLoginToken` operation in the session. If MFA is required, you must specify this parameter in subsequent stages.
	//
	// > For more information about each authentication stage, see the parameter description of the request parameter `CurrentStage`.
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// The ID of the Alibaba Cloud account. The ID is used for hardware client authentication.
	TenantId *int64 `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// > This is a parameter only for internal use.
	WindowDisplayMode *string `json:"WindowDisplayMode,omitempty" xml:"WindowDisplayMode,omitempty"`
}

func (s GetLoginTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetLoginTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GetLoginTokenResponseBody) SetEmail(v string) *GetLoginTokenResponseBody {
	s.Email = &v
	return s
}

func (s *GetLoginTokenResponseBody) SetEndUserId(v string) *GetLoginTokenResponseBody {
	s.EndUserId = &v
	return s
}

func (s *GetLoginTokenResponseBody) SetIndustry(v string) *GetLoginTokenResponseBody {
	s.Industry = &v
	return s
}

func (s *GetLoginTokenResponseBody) SetKeepAliveToken(v string) *GetLoginTokenResponseBody {
	s.KeepAliveToken = &v
	return s
}

func (s *GetLoginTokenResponseBody) SetLabel(v string) *GetLoginTokenResponseBody {
	s.Label = &v
	return s
}

func (s *GetLoginTokenResponseBody) SetLoginToken(v string) *GetLoginTokenResponseBody {
	s.LoginToken = &v
	return s
}

func (s *GetLoginTokenResponseBody) SetNextStage(v string) *GetLoginTokenResponseBody {
	s.NextStage = &v
	return s
}

func (s *GetLoginTokenResponseBody) SetPhone(v string) *GetLoginTokenResponseBody {
	s.Phone = &v
	return s
}

func (s *GetLoginTokenResponseBody) SetProps(v map[string]*string) *GetLoginTokenResponseBody {
	s.Props = v
	return s
}

func (s *GetLoginTokenResponseBody) SetQrCodePng(v string) *GetLoginTokenResponseBody {
	s.QrCodePng = &v
	return s
}

func (s *GetLoginTokenResponseBody) SetRequestId(v string) *GetLoginTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLoginTokenResponseBody) SetRiskVerifyInfo(v *GetLoginTokenResponseBodyRiskVerifyInfo) *GetLoginTokenResponseBody {
	s.RiskVerifyInfo = v
	return s
}

func (s *GetLoginTokenResponseBody) SetSecret(v string) *GetLoginTokenResponseBody {
	s.Secret = &v
	return s
}

func (s *GetLoginTokenResponseBody) SetSessionId(v string) *GetLoginTokenResponseBody {
	s.SessionId = &v
	return s
}

func (s *GetLoginTokenResponseBody) SetTenantId(v int64) *GetLoginTokenResponseBody {
	s.TenantId = &v
	return s
}

func (s *GetLoginTokenResponseBody) SetWindowDisplayMode(v string) *GetLoginTokenResponseBody {
	s.WindowDisplayMode = &v
	return s
}

type GetLoginTokenResponseBodyRiskVerifyInfo struct {
	Email            *string `json:"Email,omitempty" xml:"Email,omitempty"`
	LastLockDuration *int64  `json:"LastLockDuration,omitempty" xml:"LastLockDuration,omitempty"`
	Locked           *string `json:"Locked,omitempty" xml:"Locked,omitempty"`
	Phone            *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
}

func (s GetLoginTokenResponseBodyRiskVerifyInfo) String() string {
	return tea.Prettify(s)
}

func (s GetLoginTokenResponseBodyRiskVerifyInfo) GoString() string {
	return s.String()
}

func (s *GetLoginTokenResponseBodyRiskVerifyInfo) SetEmail(v string) *GetLoginTokenResponseBodyRiskVerifyInfo {
	s.Email = &v
	return s
}

func (s *GetLoginTokenResponseBodyRiskVerifyInfo) SetLastLockDuration(v int64) *GetLoginTokenResponseBodyRiskVerifyInfo {
	s.LastLockDuration = &v
	return s
}

func (s *GetLoginTokenResponseBodyRiskVerifyInfo) SetLocked(v string) *GetLoginTokenResponseBodyRiskVerifyInfo {
	s.Locked = &v
	return s
}

func (s *GetLoginTokenResponseBodyRiskVerifyInfo) SetPhone(v string) *GetLoginTokenResponseBodyRiskVerifyInfo {
	s.Phone = &v
	return s
}

type GetLoginTokenResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLoginTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLoginTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s GetLoginTokenResponse) GoString() string {
	return s.String()
}

func (s *GetLoginTokenResponse) SetHeaders(v map[string]*string) *GetLoginTokenResponse {
	s.Headers = v
	return s
}

func (s *GetLoginTokenResponse) SetStatusCode(v int32) *GetLoginTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLoginTokenResponse) SetBody(v *GetLoginTokenResponseBody) *GetLoginTokenResponse {
	s.Body = v
	return s
}

type IsKeepAliveRequest struct {
	ClientId     *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s IsKeepAliveRequest) String() string {
	return tea.Prettify(s)
}

func (s IsKeepAliveRequest) GoString() string {
	return s.String()
}

func (s *IsKeepAliveRequest) SetClientId(v string) *IsKeepAliveRequest {
	s.ClientId = &v
	return s
}

func (s *IsKeepAliveRequest) SetOfficeSiteId(v string) *IsKeepAliveRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *IsKeepAliveRequest) SetRegionId(v string) *IsKeepAliveRequest {
	s.RegionId = &v
	return s
}

type IsKeepAliveResponseBody struct {
	IsKeepAlive  *bool   `json:"IsKeepAlive,omitempty" xml:"IsKeepAlive,omitempty"`
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TenantId     *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s IsKeepAliveResponseBody) String() string {
	return tea.Prettify(s)
}

func (s IsKeepAliveResponseBody) GoString() string {
	return s.String()
}

func (s *IsKeepAliveResponseBody) SetIsKeepAlive(v bool) *IsKeepAliveResponseBody {
	s.IsKeepAlive = &v
	return s
}

func (s *IsKeepAliveResponseBody) SetOfficeSiteId(v string) *IsKeepAliveResponseBody {
	s.OfficeSiteId = &v
	return s
}

func (s *IsKeepAliveResponseBody) SetRequestId(v string) *IsKeepAliveResponseBody {
	s.RequestId = &v
	return s
}

func (s *IsKeepAliveResponseBody) SetTenantId(v string) *IsKeepAliveResponseBody {
	s.TenantId = &v
	return s
}

type IsKeepAliveResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *IsKeepAliveResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s IsKeepAliveResponse) String() string {
	return tea.Prettify(s)
}

func (s IsKeepAliveResponse) GoString() string {
	return s.String()
}

func (s *IsKeepAliveResponse) SetHeaders(v map[string]*string) *IsKeepAliveResponse {
	s.Headers = v
	return s
}

func (s *IsKeepAliveResponse) SetStatusCode(v int32) *IsKeepAliveResponse {
	s.StatusCode = &v
	return s
}

func (s *IsKeepAliveResponse) SetBody(v *IsKeepAliveResponseBody) *IsKeepAliveResponse {
	s.Body = v
	return s
}

type QueryEdsAgentReportConfigRequest struct {
	AliUid        *int64  `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	DesktopId     *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	EcsInstanceId *string `json:"EcsInstanceId,omitempty" xml:"EcsInstanceId,omitempty"`
}

func (s QueryEdsAgentReportConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryEdsAgentReportConfigRequest) GoString() string {
	return s.String()
}

func (s *QueryEdsAgentReportConfigRequest) SetAliUid(v int64) *QueryEdsAgentReportConfigRequest {
	s.AliUid = &v
	return s
}

func (s *QueryEdsAgentReportConfigRequest) SetDesktopId(v string) *QueryEdsAgentReportConfigRequest {
	s.DesktopId = &v
	return s
}

func (s *QueryEdsAgentReportConfigRequest) SetEcsInstanceId(v string) *QueryEdsAgentReportConfigRequest {
	s.EcsInstanceId = &v
	return s
}

type QueryEdsAgentReportConfigResponseBody struct {
	Data      *QueryEdsAgentReportConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryEdsAgentReportConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryEdsAgentReportConfigResponseBody) GoString() string {
	return s.String()
}

func (s *QueryEdsAgentReportConfigResponseBody) SetData(v *QueryEdsAgentReportConfigResponseBodyData) *QueryEdsAgentReportConfigResponseBody {
	s.Data = v
	return s
}

func (s *QueryEdsAgentReportConfigResponseBody) SetRequestId(v string) *QueryEdsAgentReportConfigResponseBody {
	s.RequestId = &v
	return s
}

type QueryEdsAgentReportConfigResponseBodyData struct {
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
}

func (s QueryEdsAgentReportConfigResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryEdsAgentReportConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryEdsAgentReportConfigResponseBodyData) SetConfig(v string) *QueryEdsAgentReportConfigResponseBodyData {
	s.Config = &v
	return s
}

type QueryEdsAgentReportConfigResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryEdsAgentReportConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryEdsAgentReportConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryEdsAgentReportConfigResponse) GoString() string {
	return s.String()
}

func (s *QueryEdsAgentReportConfigResponse) SetHeaders(v map[string]*string) *QueryEdsAgentReportConfigResponse {
	s.Headers = v
	return s
}

func (s *QueryEdsAgentReportConfigResponse) SetStatusCode(v int32) *QueryEdsAgentReportConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryEdsAgentReportConfigResponse) SetBody(v *QueryEdsAgentReportConfigResponseBody) *QueryEdsAgentReportConfigResponse {
	s.Body = v
	return s
}

type RebootDesktopsRequest struct {
	// The client ID. The system generates a unique ID for each client.
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The operating system (OS) of the device that runs the Alibaba Cloud Workspace client (hereinafter referred to as WUYING client).
	ClientOS *string `json:"ClientOS,omitempty" xml:"ClientOS,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How do I ensure the idempotence of a request?](~~25693~~)
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The client version. If you use a WUYING client, you can view the client version in the **About** dialog box on the client logon page.
	ClientVersion *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	// The IDs of the cloud computers. You can specify the IDs of 1 to 20 cloud computers.
	DesktopId []*string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" type:"Repeated"`
	// The logon token.
	LoginToken *string `json:"LoginToken,omitempty" xml:"LoginToken,omitempty"`
	// The region ID. You can call the [DescribeRegions](~~196646~~) operation to query the regions supported by WUYING Workspace.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The session ID.
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// The logon token.
	SessionToken *string `json:"SessionToken,omitempty" xml:"SessionToken,omitempty"`
	Uuid         *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s RebootDesktopsRequest) String() string {
	return tea.Prettify(s)
}

func (s RebootDesktopsRequest) GoString() string {
	return s.String()
}

func (s *RebootDesktopsRequest) SetClientId(v string) *RebootDesktopsRequest {
	s.ClientId = &v
	return s
}

func (s *RebootDesktopsRequest) SetClientOS(v string) *RebootDesktopsRequest {
	s.ClientOS = &v
	return s
}

func (s *RebootDesktopsRequest) SetClientToken(v string) *RebootDesktopsRequest {
	s.ClientToken = &v
	return s
}

func (s *RebootDesktopsRequest) SetClientVersion(v string) *RebootDesktopsRequest {
	s.ClientVersion = &v
	return s
}

func (s *RebootDesktopsRequest) SetDesktopId(v []*string) *RebootDesktopsRequest {
	s.DesktopId = v
	return s
}

func (s *RebootDesktopsRequest) SetLoginToken(v string) *RebootDesktopsRequest {
	s.LoginToken = &v
	return s
}

func (s *RebootDesktopsRequest) SetRegionId(v string) *RebootDesktopsRequest {
	s.RegionId = &v
	return s
}

func (s *RebootDesktopsRequest) SetSessionId(v string) *RebootDesktopsRequest {
	s.SessionId = &v
	return s
}

func (s *RebootDesktopsRequest) SetSessionToken(v string) *RebootDesktopsRequest {
	s.SessionToken = &v
	return s
}

func (s *RebootDesktopsRequest) SetUuid(v string) *RebootDesktopsRequest {
	s.Uuid = &v
	return s
}

type RebootDesktopsResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RebootDesktopsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RebootDesktopsResponseBody) GoString() string {
	return s.String()
}

func (s *RebootDesktopsResponseBody) SetRequestId(v string) *RebootDesktopsResponseBody {
	s.RequestId = &v
	return s
}

type RebootDesktopsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RebootDesktopsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RebootDesktopsResponse) String() string {
	return tea.Prettify(s)
}

func (s RebootDesktopsResponse) GoString() string {
	return s.String()
}

func (s *RebootDesktopsResponse) SetHeaders(v map[string]*string) *RebootDesktopsResponse {
	s.Headers = v
	return s
}

func (s *RebootDesktopsResponse) SetStatusCode(v int32) *RebootDesktopsResponse {
	s.StatusCode = &v
	return s
}

func (s *RebootDesktopsResponse) SetBody(v *RebootDesktopsResponseBody) *RebootDesktopsResponse {
	s.Body = v
	return s
}

type RefreshLoginTokenRequest struct {
	ClientId     *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	DirectoryId  *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	EndUserId    *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	LoginToken   *string `json:"LoginToken,omitempty" xml:"LoginToken,omitempty"`
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SessionId    *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s RefreshLoginTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s RefreshLoginTokenRequest) GoString() string {
	return s.String()
}

func (s *RefreshLoginTokenRequest) SetClientId(v string) *RefreshLoginTokenRequest {
	s.ClientId = &v
	return s
}

func (s *RefreshLoginTokenRequest) SetDirectoryId(v string) *RefreshLoginTokenRequest {
	s.DirectoryId = &v
	return s
}

func (s *RefreshLoginTokenRequest) SetEndUserId(v string) *RefreshLoginTokenRequest {
	s.EndUserId = &v
	return s
}

func (s *RefreshLoginTokenRequest) SetLoginToken(v string) *RefreshLoginTokenRequest {
	s.LoginToken = &v
	return s
}

func (s *RefreshLoginTokenRequest) SetOfficeSiteId(v string) *RefreshLoginTokenRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *RefreshLoginTokenRequest) SetRegionId(v string) *RefreshLoginTokenRequest {
	s.RegionId = &v
	return s
}

func (s *RefreshLoginTokenRequest) SetSessionId(v string) *RefreshLoginTokenRequest {
	s.SessionId = &v
	return s
}

type RefreshLoginTokenResponseBody struct {
	LoginToken *string `json:"LoginToken,omitempty" xml:"LoginToken,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RefreshLoginTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RefreshLoginTokenResponseBody) GoString() string {
	return s.String()
}

func (s *RefreshLoginTokenResponseBody) SetLoginToken(v string) *RefreshLoginTokenResponseBody {
	s.LoginToken = &v
	return s
}

func (s *RefreshLoginTokenResponseBody) SetRequestId(v string) *RefreshLoginTokenResponseBody {
	s.RequestId = &v
	return s
}

type RefreshLoginTokenResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RefreshLoginTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RefreshLoginTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s RefreshLoginTokenResponse) GoString() string {
	return s.String()
}

func (s *RefreshLoginTokenResponse) SetHeaders(v map[string]*string) *RefreshLoginTokenResponse {
	s.Headers = v
	return s
}

func (s *RefreshLoginTokenResponse) SetStatusCode(v int32) *RefreshLoginTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *RefreshLoginTokenResponse) SetBody(v *RefreshLoginTokenResponseBody) *RefreshLoginTokenResponse {
	s.Body = v
	return s
}

type ReportEdsAgentInfoRequest struct {
	AliUid        *int64  `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	DesktopId     *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	EcsInstanceId *string `json:"EcsInstanceId,omitempty" xml:"EcsInstanceId,omitempty"`
	EdsAgentInfo  *string `json:"EdsAgentInfo,omitempty" xml:"EdsAgentInfo,omitempty"`
}

func (s ReportEdsAgentInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s ReportEdsAgentInfoRequest) GoString() string {
	return s.String()
}

func (s *ReportEdsAgentInfoRequest) SetAliUid(v int64) *ReportEdsAgentInfoRequest {
	s.AliUid = &v
	return s
}

func (s *ReportEdsAgentInfoRequest) SetDesktopId(v string) *ReportEdsAgentInfoRequest {
	s.DesktopId = &v
	return s
}

func (s *ReportEdsAgentInfoRequest) SetEcsInstanceId(v string) *ReportEdsAgentInfoRequest {
	s.EcsInstanceId = &v
	return s
}

func (s *ReportEdsAgentInfoRequest) SetEdsAgentInfo(v string) *ReportEdsAgentInfoRequest {
	s.EdsAgentInfo = &v
	return s
}

type ReportEdsAgentInfoResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReportEdsAgentInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReportEdsAgentInfoResponseBody) GoString() string {
	return s.String()
}

func (s *ReportEdsAgentInfoResponseBody) SetRequestId(v string) *ReportEdsAgentInfoResponseBody {
	s.RequestId = &v
	return s
}

type ReportEdsAgentInfoResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReportEdsAgentInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReportEdsAgentInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s ReportEdsAgentInfoResponse) GoString() string {
	return s.String()
}

func (s *ReportEdsAgentInfoResponse) SetHeaders(v map[string]*string) *ReportEdsAgentInfoResponse {
	s.Headers = v
	return s
}

func (s *ReportEdsAgentInfoResponse) SetStatusCode(v int32) *ReportEdsAgentInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *ReportEdsAgentInfoResponse) SetBody(v *ReportEdsAgentInfoResponseBody) *ReportEdsAgentInfoResponse {
	s.Body = v
	return s
}

type ReportSessionStatusRequest struct {
	EndUserId         *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	InstanceId        *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId          *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SessionChangeTime *int64  `json:"SessionChangeTime,omitempty" xml:"SessionChangeTime,omitempty"`
	SessionId         *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	SessionStatus     *string `json:"SessionStatus,omitempty" xml:"SessionStatus,omitempty"`
}

func (s ReportSessionStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s ReportSessionStatusRequest) GoString() string {
	return s.String()
}

func (s *ReportSessionStatusRequest) SetEndUserId(v string) *ReportSessionStatusRequest {
	s.EndUserId = &v
	return s
}

func (s *ReportSessionStatusRequest) SetInstanceId(v string) *ReportSessionStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *ReportSessionStatusRequest) SetRegionId(v string) *ReportSessionStatusRequest {
	s.RegionId = &v
	return s
}

func (s *ReportSessionStatusRequest) SetSessionChangeTime(v int64) *ReportSessionStatusRequest {
	s.SessionChangeTime = &v
	return s
}

func (s *ReportSessionStatusRequest) SetSessionId(v string) *ReportSessionStatusRequest {
	s.SessionId = &v
	return s
}

func (s *ReportSessionStatusRequest) SetSessionStatus(v string) *ReportSessionStatusRequest {
	s.SessionStatus = &v
	return s
}

type ReportSessionStatusResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReportSessionStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReportSessionStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ReportSessionStatusResponseBody) SetRequestId(v string) *ReportSessionStatusResponseBody {
	s.RequestId = &v
	return s
}

type ReportSessionStatusResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReportSessionStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReportSessionStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s ReportSessionStatusResponse) GoString() string {
	return s.String()
}

func (s *ReportSessionStatusResponse) SetHeaders(v map[string]*string) *ReportSessionStatusResponse {
	s.Headers = v
	return s
}

func (s *ReportSessionStatusResponse) SetStatusCode(v int32) *ReportSessionStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *ReportSessionStatusResponse) SetBody(v *ReportSessionStatusResponseBody) *ReportSessionStatusResponse {
	s.Body = v
	return s
}

type ResetPasswordRequest struct {
	ClientId     *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Email        *string `json:"Email,omitempty" xml:"Email,omitempty"`
	EndUserId    *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Phone        *string `json:"phone,omitempty" xml:"phone,omitempty"`
}

func (s ResetPasswordRequest) String() string {
	return tea.Prettify(s)
}

func (s ResetPasswordRequest) GoString() string {
	return s.String()
}

func (s *ResetPasswordRequest) SetClientId(v string) *ResetPasswordRequest {
	s.ClientId = &v
	return s
}

func (s *ResetPasswordRequest) SetClientToken(v string) *ResetPasswordRequest {
	s.ClientToken = &v
	return s
}

func (s *ResetPasswordRequest) SetEmail(v string) *ResetPasswordRequest {
	s.Email = &v
	return s
}

func (s *ResetPasswordRequest) SetEndUserId(v string) *ResetPasswordRequest {
	s.EndUserId = &v
	return s
}

func (s *ResetPasswordRequest) SetOfficeSiteId(v string) *ResetPasswordRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *ResetPasswordRequest) SetRegionId(v string) *ResetPasswordRequest {
	s.RegionId = &v
	return s
}

func (s *ResetPasswordRequest) SetPhone(v string) *ResetPasswordRequest {
	s.Phone = &v
	return s
}

type ResetPasswordResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResetPasswordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResetPasswordResponseBody) GoString() string {
	return s.String()
}

func (s *ResetPasswordResponseBody) SetRequestId(v string) *ResetPasswordResponseBody {
	s.RequestId = &v
	return s
}

type ResetPasswordResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResetPasswordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResetPasswordResponse) String() string {
	return tea.Prettify(s)
}

func (s ResetPasswordResponse) GoString() string {
	return s.String()
}

func (s *ResetPasswordResponse) SetHeaders(v map[string]*string) *ResetPasswordResponse {
	s.Headers = v
	return s
}

func (s *ResetPasswordResponse) SetStatusCode(v int32) *ResetPasswordResponse {
	s.StatusCode = &v
	return s
}

func (s *ResetPasswordResponse) SetBody(v *ResetPasswordResponseBody) *ResetPasswordResponse {
	s.Body = v
	return s
}

type ResetSnapshotRequest struct {
	ClientId   *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	LoginToken *string `json:"LoginToken,omitempty" xml:"LoginToken,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SessionId  *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
}

func (s ResetSnapshotRequest) String() string {
	return tea.Prettify(s)
}

func (s ResetSnapshotRequest) GoString() string {
	return s.String()
}

func (s *ResetSnapshotRequest) SetClientId(v string) *ResetSnapshotRequest {
	s.ClientId = &v
	return s
}

func (s *ResetSnapshotRequest) SetLoginToken(v string) *ResetSnapshotRequest {
	s.LoginToken = &v
	return s
}

func (s *ResetSnapshotRequest) SetRegionId(v string) *ResetSnapshotRequest {
	s.RegionId = &v
	return s
}

func (s *ResetSnapshotRequest) SetSessionId(v string) *ResetSnapshotRequest {
	s.SessionId = &v
	return s
}

func (s *ResetSnapshotRequest) SetSnapshotId(v string) *ResetSnapshotRequest {
	s.SnapshotId = &v
	return s
}

type ResetSnapshotResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResetSnapshotResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResetSnapshotResponseBody) GoString() string {
	return s.String()
}

func (s *ResetSnapshotResponseBody) SetRequestId(v string) *ResetSnapshotResponseBody {
	s.RequestId = &v
	return s
}

type ResetSnapshotResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResetSnapshotResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResetSnapshotResponse) String() string {
	return tea.Prettify(s)
}

func (s ResetSnapshotResponse) GoString() string {
	return s.String()
}

func (s *ResetSnapshotResponse) SetHeaders(v map[string]*string) *ResetSnapshotResponse {
	s.Headers = v
	return s
}

func (s *ResetSnapshotResponse) SetStatusCode(v int32) *ResetSnapshotResponse {
	s.StatusCode = &v
	return s
}

func (s *ResetSnapshotResponse) SetBody(v *ResetSnapshotResponseBody) *ResetSnapshotResponse {
	s.Body = v
	return s
}

type SendTokenCodeRequest struct {
	ClientId      *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	ClientOS      *string `json:"ClientOS,omitempty" xml:"ClientOS,omitempty"`
	ClientVersion *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	EndUserId     *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	LoginToken    *string `json:"LoginToken,omitempty" xml:"LoginToken,omitempty"`
	OfficeSiteId  *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	SessionId     *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	TokenCode     *string `json:"TokenCode,omitempty" xml:"TokenCode,omitempty"`
}

func (s SendTokenCodeRequest) String() string {
	return tea.Prettify(s)
}

func (s SendTokenCodeRequest) GoString() string {
	return s.String()
}

func (s *SendTokenCodeRequest) SetClientId(v string) *SendTokenCodeRequest {
	s.ClientId = &v
	return s
}

func (s *SendTokenCodeRequest) SetClientOS(v string) *SendTokenCodeRequest {
	s.ClientOS = &v
	return s
}

func (s *SendTokenCodeRequest) SetClientVersion(v string) *SendTokenCodeRequest {
	s.ClientVersion = &v
	return s
}

func (s *SendTokenCodeRequest) SetEndUserId(v string) *SendTokenCodeRequest {
	s.EndUserId = &v
	return s
}

func (s *SendTokenCodeRequest) SetLoginToken(v string) *SendTokenCodeRequest {
	s.LoginToken = &v
	return s
}

func (s *SendTokenCodeRequest) SetOfficeSiteId(v string) *SendTokenCodeRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *SendTokenCodeRequest) SetSessionId(v string) *SendTokenCodeRequest {
	s.SessionId = &v
	return s
}

func (s *SendTokenCodeRequest) SetTokenCode(v string) *SendTokenCodeRequest {
	s.TokenCode = &v
	return s
}

type SendTokenCodeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SendTokenCodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendTokenCodeResponseBody) GoString() string {
	return s.String()
}

func (s *SendTokenCodeResponseBody) SetRequestId(v string) *SendTokenCodeResponseBody {
	s.RequestId = &v
	return s
}

type SendTokenCodeResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendTokenCodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendTokenCodeResponse) String() string {
	return tea.Prettify(s)
}

func (s SendTokenCodeResponse) GoString() string {
	return s.String()
}

func (s *SendTokenCodeResponse) SetHeaders(v map[string]*string) *SendTokenCodeResponse {
	s.Headers = v
	return s
}

func (s *SendTokenCodeResponse) SetStatusCode(v int32) *SendTokenCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *SendTokenCodeResponse) SetBody(v *SendTokenCodeResponseBody) *SendTokenCodeResponse {
	s.Body = v
	return s
}

type SetFingerPrintTemplateRequest struct {
	ClientId                     *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	ClientToken                  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Description                  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	EncryptedFingerPrintTemplate *string `json:"EncryptedFingerPrintTemplate,omitempty" xml:"EncryptedFingerPrintTemplate,omitempty"`
	EncryptedKey                 *string `json:"EncryptedKey,omitempty" xml:"EncryptedKey,omitempty"`
	FingerPrintTemplate          *string `json:"FingerPrintTemplate,omitempty" xml:"FingerPrintTemplate,omitempty"`
	LoginToken                   *string `json:"LoginToken,omitempty" xml:"LoginToken,omitempty"`
	Password                     *string `json:"Password,omitempty" xml:"Password,omitempty"`
	RegionId                     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SessionId                    *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s SetFingerPrintTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s SetFingerPrintTemplateRequest) GoString() string {
	return s.String()
}

func (s *SetFingerPrintTemplateRequest) SetClientId(v string) *SetFingerPrintTemplateRequest {
	s.ClientId = &v
	return s
}

func (s *SetFingerPrintTemplateRequest) SetClientToken(v string) *SetFingerPrintTemplateRequest {
	s.ClientToken = &v
	return s
}

func (s *SetFingerPrintTemplateRequest) SetDescription(v string) *SetFingerPrintTemplateRequest {
	s.Description = &v
	return s
}

func (s *SetFingerPrintTemplateRequest) SetEncryptedFingerPrintTemplate(v string) *SetFingerPrintTemplateRequest {
	s.EncryptedFingerPrintTemplate = &v
	return s
}

func (s *SetFingerPrintTemplateRequest) SetEncryptedKey(v string) *SetFingerPrintTemplateRequest {
	s.EncryptedKey = &v
	return s
}

func (s *SetFingerPrintTemplateRequest) SetFingerPrintTemplate(v string) *SetFingerPrintTemplateRequest {
	s.FingerPrintTemplate = &v
	return s
}

func (s *SetFingerPrintTemplateRequest) SetLoginToken(v string) *SetFingerPrintTemplateRequest {
	s.LoginToken = &v
	return s
}

func (s *SetFingerPrintTemplateRequest) SetPassword(v string) *SetFingerPrintTemplateRequest {
	s.Password = &v
	return s
}

func (s *SetFingerPrintTemplateRequest) SetRegionId(v string) *SetFingerPrintTemplateRequest {
	s.RegionId = &v
	return s
}

func (s *SetFingerPrintTemplateRequest) SetSessionId(v string) *SetFingerPrintTemplateRequest {
	s.SessionId = &v
	return s
}

type SetFingerPrintTemplateResponseBody struct {
	EncryptedPassword *string `json:"EncryptedPassword,omitempty" xml:"EncryptedPassword,omitempty"`
	Index             *int32  `json:"Index,omitempty" xml:"Index,omitempty"`
	RequestId         *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetFingerPrintTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetFingerPrintTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *SetFingerPrintTemplateResponseBody) SetEncryptedPassword(v string) *SetFingerPrintTemplateResponseBody {
	s.EncryptedPassword = &v
	return s
}

func (s *SetFingerPrintTemplateResponseBody) SetIndex(v int32) *SetFingerPrintTemplateResponseBody {
	s.Index = &v
	return s
}

func (s *SetFingerPrintTemplateResponseBody) SetRequestId(v string) *SetFingerPrintTemplateResponseBody {
	s.RequestId = &v
	return s
}

type SetFingerPrintTemplateResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetFingerPrintTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetFingerPrintTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s SetFingerPrintTemplateResponse) GoString() string {
	return s.String()
}

func (s *SetFingerPrintTemplateResponse) SetHeaders(v map[string]*string) *SetFingerPrintTemplateResponse {
	s.Headers = v
	return s
}

func (s *SetFingerPrintTemplateResponse) SetStatusCode(v int32) *SetFingerPrintTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *SetFingerPrintTemplateResponse) SetBody(v *SetFingerPrintTemplateResponseBody) *SetFingerPrintTemplateResponse {
	s.Body = v
	return s
}

type SetFingerPrintTemplateDescriptionRequest struct {
	ClientId    *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Index       *int32  `json:"Index,omitempty" xml:"Index,omitempty"`
	LoginToken  *string `json:"LoginToken,omitempty" xml:"LoginToken,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SessionId   *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s SetFingerPrintTemplateDescriptionRequest) String() string {
	return tea.Prettify(s)
}

func (s SetFingerPrintTemplateDescriptionRequest) GoString() string {
	return s.String()
}

func (s *SetFingerPrintTemplateDescriptionRequest) SetClientId(v string) *SetFingerPrintTemplateDescriptionRequest {
	s.ClientId = &v
	return s
}

func (s *SetFingerPrintTemplateDescriptionRequest) SetClientToken(v string) *SetFingerPrintTemplateDescriptionRequest {
	s.ClientToken = &v
	return s
}

func (s *SetFingerPrintTemplateDescriptionRequest) SetDescription(v string) *SetFingerPrintTemplateDescriptionRequest {
	s.Description = &v
	return s
}

func (s *SetFingerPrintTemplateDescriptionRequest) SetIndex(v int32) *SetFingerPrintTemplateDescriptionRequest {
	s.Index = &v
	return s
}

func (s *SetFingerPrintTemplateDescriptionRequest) SetLoginToken(v string) *SetFingerPrintTemplateDescriptionRequest {
	s.LoginToken = &v
	return s
}

func (s *SetFingerPrintTemplateDescriptionRequest) SetRegionId(v string) *SetFingerPrintTemplateDescriptionRequest {
	s.RegionId = &v
	return s
}

func (s *SetFingerPrintTemplateDescriptionRequest) SetSessionId(v string) *SetFingerPrintTemplateDescriptionRequest {
	s.SessionId = &v
	return s
}

type SetFingerPrintTemplateDescriptionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetFingerPrintTemplateDescriptionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetFingerPrintTemplateDescriptionResponseBody) GoString() string {
	return s.String()
}

func (s *SetFingerPrintTemplateDescriptionResponseBody) SetRequestId(v string) *SetFingerPrintTemplateDescriptionResponseBody {
	s.RequestId = &v
	return s
}

type SetFingerPrintTemplateDescriptionResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetFingerPrintTemplateDescriptionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetFingerPrintTemplateDescriptionResponse) String() string {
	return tea.Prettify(s)
}

func (s SetFingerPrintTemplateDescriptionResponse) GoString() string {
	return s.String()
}

func (s *SetFingerPrintTemplateDescriptionResponse) SetHeaders(v map[string]*string) *SetFingerPrintTemplateDescriptionResponse {
	s.Headers = v
	return s
}

func (s *SetFingerPrintTemplateDescriptionResponse) SetStatusCode(v int32) *SetFingerPrintTemplateDescriptionResponse {
	s.StatusCode = &v
	return s
}

func (s *SetFingerPrintTemplateDescriptionResponse) SetBody(v *SetFingerPrintTemplateDescriptionResponseBody) *SetFingerPrintTemplateDescriptionResponse {
	s.Body = v
	return s
}

type StartDesktopsRequest struct {
	// The ID of the Alibaba Cloud Workspace client (hereinafter referred to as WUYING client). The system generates a unique ID for each client.
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The operating system (OS) of the device that run the client.
	ClientOS *string `json:"ClientOS,omitempty" xml:"ClientOS,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](~~25693~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The client version. If you use a WUYING client, you can click **About** on the client logon page to view the version of the client.
	ClientVersion *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	// The IDs of the cloud computers. You can specify the IDs of 1 to 20 cloud computers.
	DesktopId []*string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" type:"Repeated"`
	// The logon token.
	LoginToken *string `json:"LoginToken,omitempty" xml:"LoginToken,omitempty"`
	// The region ID. You can call the [DescribeRegions](~~196646~~) operation to query the regions supported by WUYING Workspace.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The session ID.
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	Uuid      *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s StartDesktopsRequest) String() string {
	return tea.Prettify(s)
}

func (s StartDesktopsRequest) GoString() string {
	return s.String()
}

func (s *StartDesktopsRequest) SetClientId(v string) *StartDesktopsRequest {
	s.ClientId = &v
	return s
}

func (s *StartDesktopsRequest) SetClientOS(v string) *StartDesktopsRequest {
	s.ClientOS = &v
	return s
}

func (s *StartDesktopsRequest) SetClientToken(v string) *StartDesktopsRequest {
	s.ClientToken = &v
	return s
}

func (s *StartDesktopsRequest) SetClientVersion(v string) *StartDesktopsRequest {
	s.ClientVersion = &v
	return s
}

func (s *StartDesktopsRequest) SetDesktopId(v []*string) *StartDesktopsRequest {
	s.DesktopId = v
	return s
}

func (s *StartDesktopsRequest) SetLoginToken(v string) *StartDesktopsRequest {
	s.LoginToken = &v
	return s
}

func (s *StartDesktopsRequest) SetRegionId(v string) *StartDesktopsRequest {
	s.RegionId = &v
	return s
}

func (s *StartDesktopsRequest) SetSessionId(v string) *StartDesktopsRequest {
	s.SessionId = &v
	return s
}

func (s *StartDesktopsRequest) SetUuid(v string) *StartDesktopsRequest {
	s.Uuid = &v
	return s
}

type StartDesktopsResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartDesktopsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartDesktopsResponseBody) GoString() string {
	return s.String()
}

func (s *StartDesktopsResponseBody) SetRequestId(v string) *StartDesktopsResponseBody {
	s.RequestId = &v
	return s
}

type StartDesktopsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartDesktopsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartDesktopsResponse) String() string {
	return tea.Prettify(s)
}

func (s StartDesktopsResponse) GoString() string {
	return s.String()
}

func (s *StartDesktopsResponse) SetHeaders(v map[string]*string) *StartDesktopsResponse {
	s.Headers = v
	return s
}

func (s *StartDesktopsResponse) SetStatusCode(v int32) *StartDesktopsResponse {
	s.StatusCode = &v
	return s
}

func (s *StartDesktopsResponse) SetBody(v *StartDesktopsResponseBody) *StartDesktopsResponse {
	s.Body = v
	return s
}

type StartRecordContentRequest struct {
	ClientId      *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	ClientOS      *string `json:"ClientOS,omitempty" xml:"ClientOS,omitempty"`
	ClientVersion *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	DesktopId     *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	FilePath      *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
	LoginToken    *string `json:"LoginToken,omitempty" xml:"LoginToken,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SessionId     *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s StartRecordContentRequest) String() string {
	return tea.Prettify(s)
}

func (s StartRecordContentRequest) GoString() string {
	return s.String()
}

func (s *StartRecordContentRequest) SetClientId(v string) *StartRecordContentRequest {
	s.ClientId = &v
	return s
}

func (s *StartRecordContentRequest) SetClientOS(v string) *StartRecordContentRequest {
	s.ClientOS = &v
	return s
}

func (s *StartRecordContentRequest) SetClientVersion(v string) *StartRecordContentRequest {
	s.ClientVersion = &v
	return s
}

func (s *StartRecordContentRequest) SetDesktopId(v string) *StartRecordContentRequest {
	s.DesktopId = &v
	return s
}

func (s *StartRecordContentRequest) SetFilePath(v string) *StartRecordContentRequest {
	s.FilePath = &v
	return s
}

func (s *StartRecordContentRequest) SetLoginToken(v string) *StartRecordContentRequest {
	s.LoginToken = &v
	return s
}

func (s *StartRecordContentRequest) SetRegionId(v string) *StartRecordContentRequest {
	s.RegionId = &v
	return s
}

func (s *StartRecordContentRequest) SetSessionId(v string) *StartRecordContentRequest {
	s.SessionId = &v
	return s
}

type StartRecordContentResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartRecordContentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartRecordContentResponseBody) GoString() string {
	return s.String()
}

func (s *StartRecordContentResponseBody) SetRequestId(v string) *StartRecordContentResponseBody {
	s.RequestId = &v
	return s
}

type StartRecordContentResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartRecordContentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartRecordContentResponse) String() string {
	return tea.Prettify(s)
}

func (s StartRecordContentResponse) GoString() string {
	return s.String()
}

func (s *StartRecordContentResponse) SetHeaders(v map[string]*string) *StartRecordContentResponse {
	s.Headers = v
	return s
}

func (s *StartRecordContentResponse) SetStatusCode(v int32) *StartRecordContentResponse {
	s.StatusCode = &v
	return s
}

func (s *StartRecordContentResponse) SetBody(v *StartRecordContentResponseBody) *StartRecordContentResponse {
	s.Body = v
	return s
}

type StopDesktopsRequest struct {
	// The client ID. The system generates a unique ID for each client.
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The operating system (OS) of the device that runs the Alibaba Cloud Workspace client (hereinafter referred to as WUYING client).
	ClientOS *string `json:"ClientOS,omitempty" xml:"ClientOS,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How do I ensure the idempotence of a request?](~~25693~~)
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The client version. If you use a WUYING client, you can view the client version in the **About** dialog box on the client logon page.
	ClientVersion *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	// The IDs of the cloud computers. You can specify the IDs of 1 to 20 cloud computers.
	DesktopId []*string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" type:"Repeated"`
	// The logon token.
	LoginToken *string `json:"LoginToken,omitempty" xml:"LoginToken,omitempty"`
	// The region ID. You can call the [DescribeRegions](~~196646~~) operation to query the regions supported by WUYING Workspace.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The session ID.
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// The logon token.
	SessionToken *string `json:"SessionToken,omitempty" xml:"SessionToken,omitempty"`
}

func (s StopDesktopsRequest) String() string {
	return tea.Prettify(s)
}

func (s StopDesktopsRequest) GoString() string {
	return s.String()
}

func (s *StopDesktopsRequest) SetClientId(v string) *StopDesktopsRequest {
	s.ClientId = &v
	return s
}

func (s *StopDesktopsRequest) SetClientOS(v string) *StopDesktopsRequest {
	s.ClientOS = &v
	return s
}

func (s *StopDesktopsRequest) SetClientToken(v string) *StopDesktopsRequest {
	s.ClientToken = &v
	return s
}

func (s *StopDesktopsRequest) SetClientVersion(v string) *StopDesktopsRequest {
	s.ClientVersion = &v
	return s
}

func (s *StopDesktopsRequest) SetDesktopId(v []*string) *StopDesktopsRequest {
	s.DesktopId = v
	return s
}

func (s *StopDesktopsRequest) SetLoginToken(v string) *StopDesktopsRequest {
	s.LoginToken = &v
	return s
}

func (s *StopDesktopsRequest) SetRegionId(v string) *StopDesktopsRequest {
	s.RegionId = &v
	return s
}

func (s *StopDesktopsRequest) SetSessionId(v string) *StopDesktopsRequest {
	s.SessionId = &v
	return s
}

func (s *StopDesktopsRequest) SetSessionToken(v string) *StopDesktopsRequest {
	s.SessionToken = &v
	return s
}

type StopDesktopsResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopDesktopsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopDesktopsResponseBody) GoString() string {
	return s.String()
}

func (s *StopDesktopsResponseBody) SetRequestId(v string) *StopDesktopsResponseBody {
	s.RequestId = &v
	return s
}

type StopDesktopsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopDesktopsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopDesktopsResponse) String() string {
	return tea.Prettify(s)
}

func (s StopDesktopsResponse) GoString() string {
	return s.String()
}

func (s *StopDesktopsResponse) SetHeaders(v map[string]*string) *StopDesktopsResponse {
	s.Headers = v
	return s
}

func (s *StopDesktopsResponse) SetStatusCode(v int32) *StopDesktopsResponse {
	s.StatusCode = &v
	return s
}

func (s *StopDesktopsResponse) SetBody(v *StopDesktopsResponseBody) *StopDesktopsResponse {
	s.Body = v
	return s
}

type StopRecordContentRequest struct {
	ClientId      *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	ClientOS      *string `json:"ClientOS,omitempty" xml:"ClientOS,omitempty"`
	ClientVersion *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	DesktopId     *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	LoginToken    *string `json:"LoginToken,omitempty" xml:"LoginToken,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SessionId     *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s StopRecordContentRequest) String() string {
	return tea.Prettify(s)
}

func (s StopRecordContentRequest) GoString() string {
	return s.String()
}

func (s *StopRecordContentRequest) SetClientId(v string) *StopRecordContentRequest {
	s.ClientId = &v
	return s
}

func (s *StopRecordContentRequest) SetClientOS(v string) *StopRecordContentRequest {
	s.ClientOS = &v
	return s
}

func (s *StopRecordContentRequest) SetClientVersion(v string) *StopRecordContentRequest {
	s.ClientVersion = &v
	return s
}

func (s *StopRecordContentRequest) SetDesktopId(v string) *StopRecordContentRequest {
	s.DesktopId = &v
	return s
}

func (s *StopRecordContentRequest) SetLoginToken(v string) *StopRecordContentRequest {
	s.LoginToken = &v
	return s
}

func (s *StopRecordContentRequest) SetRegionId(v string) *StopRecordContentRequest {
	s.RegionId = &v
	return s
}

func (s *StopRecordContentRequest) SetSessionId(v string) *StopRecordContentRequest {
	s.SessionId = &v
	return s
}

type StopRecordContentResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopRecordContentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopRecordContentResponseBody) GoString() string {
	return s.String()
}

func (s *StopRecordContentResponseBody) SetRequestId(v string) *StopRecordContentResponseBody {
	s.RequestId = &v
	return s
}

type StopRecordContentResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopRecordContentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopRecordContentResponse) String() string {
	return tea.Prettify(s)
}

func (s StopRecordContentResponse) GoString() string {
	return s.String()
}

func (s *StopRecordContentResponse) SetHeaders(v map[string]*string) *StopRecordContentResponse {
	s.Headers = v
	return s
}

func (s *StopRecordContentResponse) SetStatusCode(v int32) *StopRecordContentResponse {
	s.StatusCode = &v
	return s
}

func (s *StopRecordContentResponse) SetBody(v *StopRecordContentResponseBody) *StopRecordContentResponse {
	s.Body = v
	return s
}

type UnbindUserDesktopRequest struct {
	ClientId      *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	ClientType    *string `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	Force         *bool   `json:"Force,omitempty" xml:"Force,omitempty"`
	LoginToken    *string `json:"LoginToken,omitempty" xml:"LoginToken,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SessionId     *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	UserDesktopId *string `json:"UserDesktopId,omitempty" xml:"UserDesktopId,omitempty"`
}

func (s UnbindUserDesktopRequest) String() string {
	return tea.Prettify(s)
}

func (s UnbindUserDesktopRequest) GoString() string {
	return s.String()
}

func (s *UnbindUserDesktopRequest) SetClientId(v string) *UnbindUserDesktopRequest {
	s.ClientId = &v
	return s
}

func (s *UnbindUserDesktopRequest) SetClientType(v string) *UnbindUserDesktopRequest {
	s.ClientType = &v
	return s
}

func (s *UnbindUserDesktopRequest) SetForce(v bool) *UnbindUserDesktopRequest {
	s.Force = &v
	return s
}

func (s *UnbindUserDesktopRequest) SetLoginToken(v string) *UnbindUserDesktopRequest {
	s.LoginToken = &v
	return s
}

func (s *UnbindUserDesktopRequest) SetRegionId(v string) *UnbindUserDesktopRequest {
	s.RegionId = &v
	return s
}

func (s *UnbindUserDesktopRequest) SetSessionId(v string) *UnbindUserDesktopRequest {
	s.SessionId = &v
	return s
}

func (s *UnbindUserDesktopRequest) SetUserDesktopId(v string) *UnbindUserDesktopRequest {
	s.UserDesktopId = &v
	return s
}

type UnbindUserDesktopResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnbindUserDesktopResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnbindUserDesktopResponseBody) GoString() string {
	return s.String()
}

func (s *UnbindUserDesktopResponseBody) SetRequestId(v string) *UnbindUserDesktopResponseBody {
	s.RequestId = &v
	return s
}

type UnbindUserDesktopResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnbindUserDesktopResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnbindUserDesktopResponse) String() string {
	return tea.Prettify(s)
}

func (s UnbindUserDesktopResponse) GoString() string {
	return s.String()
}

func (s *UnbindUserDesktopResponse) SetHeaders(v map[string]*string) *UnbindUserDesktopResponse {
	s.Headers = v
	return s
}

func (s *UnbindUserDesktopResponse) SetStatusCode(v int32) *UnbindUserDesktopResponse {
	s.StatusCode = &v
	return s
}

func (s *UnbindUserDesktopResponse) SetBody(v *UnbindUserDesktopResponseBody) *UnbindUserDesktopResponse {
	s.Body = v
	return s
}

type VerifyCredentialRequest struct {
	ClientId       *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	Credential     *string `json:"Credential,omitempty" xml:"Credential,omitempty"`
	CredentialType *string `json:"CredentialType,omitempty" xml:"CredentialType,omitempty"`
	EncryptedKey   *string `json:"EncryptedKey,omitempty" xml:"EncryptedKey,omitempty"`
	LoginToken     *string `json:"LoginToken,omitempty" xml:"LoginToken,omitempty"`
	OfficeSiteId   *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SessionId      *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s VerifyCredentialRequest) String() string {
	return tea.Prettify(s)
}

func (s VerifyCredentialRequest) GoString() string {
	return s.String()
}

func (s *VerifyCredentialRequest) SetClientId(v string) *VerifyCredentialRequest {
	s.ClientId = &v
	return s
}

func (s *VerifyCredentialRequest) SetCredential(v string) *VerifyCredentialRequest {
	s.Credential = &v
	return s
}

func (s *VerifyCredentialRequest) SetCredentialType(v string) *VerifyCredentialRequest {
	s.CredentialType = &v
	return s
}

func (s *VerifyCredentialRequest) SetEncryptedKey(v string) *VerifyCredentialRequest {
	s.EncryptedKey = &v
	return s
}

func (s *VerifyCredentialRequest) SetLoginToken(v string) *VerifyCredentialRequest {
	s.LoginToken = &v
	return s
}

func (s *VerifyCredentialRequest) SetOfficeSiteId(v string) *VerifyCredentialRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *VerifyCredentialRequest) SetRegionId(v string) *VerifyCredentialRequest {
	s.RegionId = &v
	return s
}

func (s *VerifyCredentialRequest) SetSessionId(v string) *VerifyCredentialRequest {
	s.SessionId = &v
	return s
}

type VerifyCredentialResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s VerifyCredentialResponseBody) String() string {
	return tea.Prettify(s)
}

func (s VerifyCredentialResponseBody) GoString() string {
	return s.String()
}

func (s *VerifyCredentialResponseBody) SetRequestId(v string) *VerifyCredentialResponseBody {
	s.RequestId = &v
	return s
}

type VerifyCredentialResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VerifyCredentialResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s VerifyCredentialResponse) String() string {
	return tea.Prettify(s)
}

func (s VerifyCredentialResponse) GoString() string {
	return s.String()
}

func (s *VerifyCredentialResponse) SetHeaders(v map[string]*string) *VerifyCredentialResponse {
	s.Headers = v
	return s
}

func (s *VerifyCredentialResponse) SetStatusCode(v int32) *VerifyCredentialResponse {
	s.StatusCode = &v
	return s
}

func (s *VerifyCredentialResponse) SetBody(v *VerifyCredentialResponseBody) *VerifyCredentialResponse {
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
	client.SignatureAlgorithm = tea.String("v2")
	client.EndpointRule = tea.String("regional")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("ecd"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) ApproveFotaUpdateWithOptions(request *ApproveFotaUpdateRequest, runtime *util.RuntimeOptions) (_result *ApproveFotaUpdateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppVersion)) {
		query["AppVersion"] = request.AppVersion
	}

	if !tea.BoolValue(util.IsUnset(request.ClientId)) {
		query["ClientId"] = request.ClientId
	}

	if !tea.BoolValue(util.IsUnset(request.DesktopId)) {
		query["DesktopId"] = request.DesktopId
	}

	if !tea.BoolValue(util.IsUnset(request.LoginToken)) {
		query["LoginToken"] = request.LoginToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		query["SessionId"] = request.SessionId
	}

	if !tea.BoolValue(util.IsUnset(request.Uuid)) {
		query["Uuid"] = request.Uuid
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ApproveFotaUpdate"),
		Version:     tea.String("2020-10-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ApproveFotaUpdateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ApproveFotaUpdate(request *ApproveFotaUpdateRequest) (_result *ApproveFotaUpdateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ApproveFotaUpdateResponse{}
	_body, _err := client.ApproveFotaUpdateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ChangePasswordWithOptions(request *ChangePasswordRequest, runtime *util.RuntimeOptions) (_result *ChangePasswordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientId)) {
		query["ClientId"] = request.ClientId
	}

	if !tea.BoolValue(util.IsUnset(request.EndUserId)) {
		query["EndUserId"] = request.EndUserId
	}

	if !tea.BoolValue(util.IsUnset(request.LoginToken)) {
		query["LoginToken"] = request.LoginToken
	}

	if !tea.BoolValue(util.IsUnset(request.NewPassword)) {
		query["NewPassword"] = request.NewPassword
	}

	if !tea.BoolValue(util.IsUnset(request.OfficeSiteId)) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !tea.BoolValue(util.IsUnset(request.OldPassword)) {
		query["OldPassword"] = request.OldPassword
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		query["SessionId"] = request.SessionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ChangePassword"),
		Version:     tea.String("2020-10-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ChangePasswordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ChangePassword(request *ChangePasswordRequest) (_result *ChangePasswordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ChangePasswordResponse{}
	_body, _err := client.ChangePasswordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteFingerPrintTemplateWithOptions(request *DeleteFingerPrintTemplateRequest, runtime *util.RuntimeOptions) (_result *DeleteFingerPrintTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientId)) {
		query["ClientId"] = request.ClientId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Index)) {
		query["Index"] = request.Index
	}

	if !tea.BoolValue(util.IsUnset(request.LoginToken)) {
		query["LoginToken"] = request.LoginToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		query["SessionId"] = request.SessionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteFingerPrintTemplate"),
		Version:     tea.String("2020-10-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteFingerPrintTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteFingerPrintTemplate(request *DeleteFingerPrintTemplateRequest) (_result *DeleteFingerPrintTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteFingerPrintTemplateResponse{}
	_body, _err := client.DeleteFingerPrintTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDirectoriesWithOptions(request *DescribeDirectoriesRequest, runtime *util.RuntimeOptions) (_result *DescribeDirectoriesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientId)) {
		query["ClientId"] = request.ClientId
	}

	if !tea.BoolValue(util.IsUnset(request.DirectoryId)) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDirectories"),
		Version:     tea.String("2020-10-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDirectoriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDirectories(request *DescribeDirectoriesRequest) (_result *DescribeDirectoriesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDirectoriesResponse{}
	_body, _err := client.DescribeDirectoriesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeFingerPrintTemplatesWithOptions(request *DescribeFingerPrintTemplatesRequest, runtime *util.RuntimeOptions) (_result *DescribeFingerPrintTemplatesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientId)) {
		query["ClientId"] = request.ClientId
	}

	if !tea.BoolValue(util.IsUnset(request.LoginToken)) {
		query["LoginToken"] = request.LoginToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		query["SessionId"] = request.SessionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeFingerPrintTemplates"),
		Version:     tea.String("2020-10-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeFingerPrintTemplatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeFingerPrintTemplates(request *DescribeFingerPrintTemplatesRequest) (_result *DescribeFingerPrintTemplatesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFingerPrintTemplatesResponse{}
	_body, _err := client.DescribeFingerPrintTemplatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeGlobalDesktopsWithOptions(request *DescribeGlobalDesktopsRequest, runtime *util.RuntimeOptions) (_result *DescribeGlobalDesktopsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientId)) {
		query["ClientId"] = request.ClientId
	}

	if !tea.BoolValue(util.IsUnset(request.DesktopAccessType)) {
		query["DesktopAccessType"] = request.DesktopAccessType
	}

	if !tea.BoolValue(util.IsUnset(request.DesktopId)) {
		query["DesktopId"] = request.DesktopId
	}

	if !tea.BoolValue(util.IsUnset(request.DesktopName)) {
		query["DesktopName"] = request.DesktopName
	}

	if !tea.BoolValue(util.IsUnset(request.DesktopStatus)) {
		query["DesktopStatus"] = request.DesktopStatus
	}

	if !tea.BoolValue(util.IsUnset(request.DirectoryId)) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		query["Keyword"] = request.Keyword
	}

	if !tea.BoolValue(util.IsUnset(request.LoginRegionId)) {
		query["LoginRegionId"] = request.LoginRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.LoginToken)) {
		query["LoginToken"] = request.LoginToken
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.OfficeSiteId)) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !tea.BoolValue(util.IsUnset(request.OrderBy)) {
		query["OrderBy"] = request.OrderBy
	}

	if !tea.BoolValue(util.IsUnset(request.QueryFotaUpdate)) {
		query["QueryFotaUpdate"] = request.QueryFotaUpdate
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SearchRegionId)) {
		query["SearchRegionId"] = request.SearchRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		query["SessionId"] = request.SessionId
	}

	if !tea.BoolValue(util.IsUnset(request.SortType)) {
		query["SortType"] = request.SortType
	}

	if !tea.BoolValue(util.IsUnset(request.WithoutLatency)) {
		query["WithoutLatency"] = request.WithoutLatency
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeGlobalDesktops"),
		Version:     tea.String("2020-10-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeGlobalDesktopsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeGlobalDesktops(request *DescribeGlobalDesktopsRequest) (_result *DescribeGlobalDesktopsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeGlobalDesktopsResponse{}
	_body, _err := client.DescribeGlobalDesktopsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeOfficeSitesWithOptions(request *DescribeOfficeSitesRequest, runtime *util.RuntimeOptions) (_result *DescribeOfficeSitesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientId)) {
		query["ClientId"] = request.ClientId
	}

	if !tea.BoolValue(util.IsUnset(request.OfficeSiteId)) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeOfficeSites"),
		Version:     tea.String("2020-10-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeOfficeSitesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeOfficeSites(request *DescribeOfficeSitesRequest) (_result *DescribeOfficeSitesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeOfficeSitesResponse{}
	_body, _err := client.DescribeOfficeSitesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRegionsWithOptions(request *DescribeRegionsRequest, runtime *util.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientId)) {
		query["ClientId"] = request.ClientId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRegions"),
		Version:     tea.String("2020-10-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRegions(request *DescribeRegionsRequest) (_result *DescribeRegionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DescribeRegionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSnapshotsWithOptions(request *DescribeSnapshotsRequest, runtime *util.RuntimeOptions) (_result *DescribeSnapshotsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientId)) {
		query["ClientId"] = request.ClientId
	}

	if !tea.BoolValue(util.IsUnset(request.DesktopId)) {
		query["DesktopId"] = request.DesktopId
	}

	if !tea.BoolValue(util.IsUnset(request.LoginToken)) {
		query["LoginToken"] = request.LoginToken
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		query["SessionId"] = request.SessionId
	}

	if !tea.BoolValue(util.IsUnset(request.SnapshotId)) {
		query["SnapshotId"] = request.SnapshotId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSnapshots"),
		Version:     tea.String("2020-10-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSnapshotsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSnapshots(request *DescribeSnapshotsRequest) (_result *DescribeSnapshotsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSnapshotsResponse{}
	_body, _err := client.DescribeSnapshotsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EncryptPasswordWithOptions(request *EncryptPasswordRequest, runtime *util.RuntimeOptions) (_result *EncryptPasswordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientId)) {
		query["ClientId"] = request.ClientId
	}

	if !tea.BoolValue(util.IsUnset(request.DirectoryId)) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !tea.BoolValue(util.IsUnset(request.LoginToken)) {
		query["LoginToken"] = request.LoginToken
	}

	if !tea.BoolValue(util.IsUnset(request.OfficeSiteId)) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !tea.BoolValue(util.IsUnset(request.Password)) {
		query["Password"] = request.Password
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		query["SessionId"] = request.SessionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("EncryptPassword"),
		Version:     tea.String("2020-10-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EncryptPasswordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EncryptPassword(request *EncryptPasswordRequest) (_result *EncryptPasswordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EncryptPasswordResponse{}
	_body, _err := client.EncryptPasswordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetCloudDriveServiceMountTokenWithOptions(request *GetCloudDriveServiceMountTokenRequest, runtime *util.RuntimeOptions) (_result *GetCloudDriveServiceMountTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientId)) {
		query["ClientId"] = request.ClientId
	}

	if !tea.BoolValue(util.IsUnset(request.LoginToken)) {
		query["LoginToken"] = request.LoginToken
	}

	if !tea.BoolValue(util.IsUnset(request.OfficeSiteId)) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		query["SessionId"] = request.SessionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetCloudDriveServiceMountToken"),
		Version:     tea.String("2020-10-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCloudDriveServiceMountTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetCloudDriveServiceMountToken(request *GetCloudDriveServiceMountTokenRequest) (_result *GetCloudDriveServiceMountTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetCloudDriveServiceMountTokenResponse{}
	_body, _err := client.GetCloudDriveServiceMountTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetConnectionTicketWithOptions(request *GetConnectionTicketRequest, runtime *util.RuntimeOptions) (_result *GetConnectionTicketResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientId)) {
		query["ClientId"] = request.ClientId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientOS)) {
		query["ClientOS"] = request.ClientOS
	}

	if !tea.BoolValue(util.IsUnset(request.ClientType)) {
		query["ClientType"] = request.ClientType
	}

	if !tea.BoolValue(util.IsUnset(request.ClientVersion)) {
		query["ClientVersion"] = request.ClientVersion
	}

	if !tea.BoolValue(util.IsUnset(request.CommandContent)) {
		query["CommandContent"] = request.CommandContent
	}

	if !tea.BoolValue(util.IsUnset(request.DesktopId)) {
		query["DesktopId"] = request.DesktopId
	}

	if !tea.BoolValue(util.IsUnset(request.LoginToken)) {
		query["LoginToken"] = request.LoginToken
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		query["SessionId"] = request.SessionId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	if !tea.BoolValue(util.IsUnset(request.Uuid)) {
		query["Uuid"] = request.Uuid
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetConnectionTicket"),
		Version:     tea.String("2020-10-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetConnectionTicketResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetConnectionTicket(request *GetConnectionTicketRequest) (_result *GetConnectionTicketResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetConnectionTicketResponse{}
	_body, _err := client.GetConnectionTicketWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetLoginTokenWithOptions(request *GetLoginTokenRequest, runtime *util.RuntimeOptions) (_result *GetLoginTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthenticationCode)) {
		query["AuthenticationCode"] = request.AuthenticationCode
	}

	if !tea.BoolValue(util.IsUnset(request.ClientId)) {
		query["ClientId"] = request.ClientId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientOS)) {
		query["ClientOS"] = request.ClientOS
	}

	if !tea.BoolValue(util.IsUnset(request.ClientType)) {
		query["ClientType"] = request.ClientType
	}

	if !tea.BoolValue(util.IsUnset(request.ClientVersion)) {
		query["ClientVersion"] = request.ClientVersion
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentStage)) {
		query["CurrentStage"] = request.CurrentStage
	}

	if !tea.BoolValue(util.IsUnset(request.DirectoryId)) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !tea.BoolValue(util.IsUnset(request.EndUserId)) {
		query["EndUserId"] = request.EndUserId
	}

	if !tea.BoolValue(util.IsUnset(request.KeepAlive)) {
		query["KeepAlive"] = request.KeepAlive
	}

	if !tea.BoolValue(util.IsUnset(request.KeepAliveToken)) {
		query["KeepAliveToken"] = request.KeepAliveToken
	}

	if !tea.BoolValue(util.IsUnset(request.NewPassword)) {
		query["NewPassword"] = request.NewPassword
	}

	if !tea.BoolValue(util.IsUnset(request.OfficeSiteId)) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !tea.BoolValue(util.IsUnset(request.OldPassword)) {
		query["OldPassword"] = request.OldPassword
	}

	if !tea.BoolValue(util.IsUnset(request.Password)) {
		query["Password"] = request.Password
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		query["SessionId"] = request.SessionId
	}

	if !tea.BoolValue(util.IsUnset(request.TokenCode)) {
		query["TokenCode"] = request.TokenCode
	}

	if !tea.BoolValue(util.IsUnset(request.Uuid)) {
		query["Uuid"] = request.Uuid
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetLoginToken"),
		Version:     tea.String("2020-10-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetLoginTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetLoginToken(request *GetLoginTokenRequest) (_result *GetLoginTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetLoginTokenResponse{}
	_body, _err := client.GetLoginTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) IsKeepAliveWithOptions(request *IsKeepAliveRequest, runtime *util.RuntimeOptions) (_result *IsKeepAliveResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientId)) {
		query["ClientId"] = request.ClientId
	}

	if !tea.BoolValue(util.IsUnset(request.OfficeSiteId)) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("IsKeepAlive"),
		Version:     tea.String("2020-10-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &IsKeepAliveResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) IsKeepAlive(request *IsKeepAliveRequest) (_result *IsKeepAliveResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &IsKeepAliveResponse{}
	_body, _err := client.IsKeepAliveWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryEdsAgentReportConfigWithOptions(request *QueryEdsAgentReportConfigRequest, runtime *util.RuntimeOptions) (_result *QueryEdsAgentReportConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AliUid)) {
		query["AliUid"] = request.AliUid
	}

	if !tea.BoolValue(util.IsUnset(request.DesktopId)) {
		query["DesktopId"] = request.DesktopId
	}

	if !tea.BoolValue(util.IsUnset(request.EcsInstanceId)) {
		query["EcsInstanceId"] = request.EcsInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryEdsAgentReportConfig"),
		Version:     tea.String("2020-10-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryEdsAgentReportConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryEdsAgentReportConfig(request *QueryEdsAgentReportConfigRequest) (_result *QueryEdsAgentReportConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryEdsAgentReportConfigResponse{}
	_body, _err := client.QueryEdsAgentReportConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RebootDesktopsWithOptions(request *RebootDesktopsRequest, runtime *util.RuntimeOptions) (_result *RebootDesktopsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientId)) {
		query["ClientId"] = request.ClientId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientOS)) {
		query["ClientOS"] = request.ClientOS
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ClientVersion)) {
		query["ClientVersion"] = request.ClientVersion
	}

	if !tea.BoolValue(util.IsUnset(request.DesktopId)) {
		query["DesktopId"] = request.DesktopId
	}

	if !tea.BoolValue(util.IsUnset(request.LoginToken)) {
		query["LoginToken"] = request.LoginToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		query["SessionId"] = request.SessionId
	}

	if !tea.BoolValue(util.IsUnset(request.SessionToken)) {
		query["SessionToken"] = request.SessionToken
	}

	if !tea.BoolValue(util.IsUnset(request.Uuid)) {
		query["Uuid"] = request.Uuid
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RebootDesktops"),
		Version:     tea.String("2020-10-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RebootDesktopsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RebootDesktops(request *RebootDesktopsRequest) (_result *RebootDesktopsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RebootDesktopsResponse{}
	_body, _err := client.RebootDesktopsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RefreshLoginTokenWithOptions(request *RefreshLoginTokenRequest, runtime *util.RuntimeOptions) (_result *RefreshLoginTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientId)) {
		query["ClientId"] = request.ClientId
	}

	if !tea.BoolValue(util.IsUnset(request.DirectoryId)) {
		query["DirectoryId"] = request.DirectoryId
	}

	if !tea.BoolValue(util.IsUnset(request.EndUserId)) {
		query["EndUserId"] = request.EndUserId
	}

	if !tea.BoolValue(util.IsUnset(request.LoginToken)) {
		query["LoginToken"] = request.LoginToken
	}

	if !tea.BoolValue(util.IsUnset(request.OfficeSiteId)) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		query["SessionId"] = request.SessionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RefreshLoginToken"),
		Version:     tea.String("2020-10-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RefreshLoginTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RefreshLoginToken(request *RefreshLoginTokenRequest) (_result *RefreshLoginTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RefreshLoginTokenResponse{}
	_body, _err := client.RefreshLoginTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ReportEdsAgentInfoWithOptions(request *ReportEdsAgentInfoRequest, runtime *util.RuntimeOptions) (_result *ReportEdsAgentInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AliUid)) {
		query["AliUid"] = request.AliUid
	}

	if !tea.BoolValue(util.IsUnset(request.DesktopId)) {
		query["DesktopId"] = request.DesktopId
	}

	if !tea.BoolValue(util.IsUnset(request.EcsInstanceId)) {
		query["EcsInstanceId"] = request.EcsInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.EdsAgentInfo)) {
		query["EdsAgentInfo"] = request.EdsAgentInfo
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ReportEdsAgentInfo"),
		Version:     tea.String("2020-10-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ReportEdsAgentInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReportEdsAgentInfo(request *ReportEdsAgentInfoRequest) (_result *ReportEdsAgentInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReportEdsAgentInfoResponse{}
	_body, _err := client.ReportEdsAgentInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ReportSessionStatusWithOptions(request *ReportSessionStatusRequest, runtime *util.RuntimeOptions) (_result *ReportSessionStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndUserId)) {
		query["EndUserId"] = request.EndUserId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SessionChangeTime)) {
		query["SessionChangeTime"] = request.SessionChangeTime
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		query["SessionId"] = request.SessionId
	}

	if !tea.BoolValue(util.IsUnset(request.SessionStatus)) {
		query["SessionStatus"] = request.SessionStatus
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ReportSessionStatus"),
		Version:     tea.String("2020-10-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ReportSessionStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReportSessionStatus(request *ReportSessionStatusRequest) (_result *ReportSessionStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReportSessionStatusResponse{}
	_body, _err := client.ReportSessionStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ResetPasswordWithOptions(request *ResetPasswordRequest, runtime *util.RuntimeOptions) (_result *ResetPasswordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientId)) {
		query["ClientId"] = request.ClientId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Email)) {
		query["Email"] = request.Email
	}

	if !tea.BoolValue(util.IsUnset(request.EndUserId)) {
		query["EndUserId"] = request.EndUserId
	}

	if !tea.BoolValue(util.IsUnset(request.OfficeSiteId)) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Phone)) {
		query["phone"] = request.Phone
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ResetPassword"),
		Version:     tea.String("2020-10-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ResetPasswordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ResetPassword(request *ResetPasswordRequest) (_result *ResetPasswordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResetPasswordResponse{}
	_body, _err := client.ResetPasswordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ResetSnapshotWithOptions(request *ResetSnapshotRequest, runtime *util.RuntimeOptions) (_result *ResetSnapshotResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientId)) {
		query["ClientId"] = request.ClientId
	}

	if !tea.BoolValue(util.IsUnset(request.LoginToken)) {
		query["LoginToken"] = request.LoginToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		query["SessionId"] = request.SessionId
	}

	if !tea.BoolValue(util.IsUnset(request.SnapshotId)) {
		query["SnapshotId"] = request.SnapshotId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ResetSnapshot"),
		Version:     tea.String("2020-10-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ResetSnapshotResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ResetSnapshot(request *ResetSnapshotRequest) (_result *ResetSnapshotResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResetSnapshotResponse{}
	_body, _err := client.ResetSnapshotWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SendTokenCodeWithOptions(request *SendTokenCodeRequest, runtime *util.RuntimeOptions) (_result *SendTokenCodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientId)) {
		query["ClientId"] = request.ClientId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientOS)) {
		query["ClientOS"] = request.ClientOS
	}

	if !tea.BoolValue(util.IsUnset(request.ClientVersion)) {
		query["ClientVersion"] = request.ClientVersion
	}

	if !tea.BoolValue(util.IsUnset(request.EndUserId)) {
		query["EndUserId"] = request.EndUserId
	}

	if !tea.BoolValue(util.IsUnset(request.LoginToken)) {
		query["LoginToken"] = request.LoginToken
	}

	if !tea.BoolValue(util.IsUnset(request.OfficeSiteId)) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		query["SessionId"] = request.SessionId
	}

	if !tea.BoolValue(util.IsUnset(request.TokenCode)) {
		query["TokenCode"] = request.TokenCode
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SendTokenCode"),
		Version:     tea.String("2020-10-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SendTokenCodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SendTokenCode(request *SendTokenCodeRequest) (_result *SendTokenCodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SendTokenCodeResponse{}
	_body, _err := client.SendTokenCodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetFingerPrintTemplateWithOptions(request *SetFingerPrintTemplateRequest, runtime *util.RuntimeOptions) (_result *SetFingerPrintTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientId)) {
		query["ClientId"] = request.ClientId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.EncryptedFingerPrintTemplate)) {
		query["EncryptedFingerPrintTemplate"] = request.EncryptedFingerPrintTemplate
	}

	if !tea.BoolValue(util.IsUnset(request.EncryptedKey)) {
		query["EncryptedKey"] = request.EncryptedKey
	}

	if !tea.BoolValue(util.IsUnset(request.FingerPrintTemplate)) {
		query["FingerPrintTemplate"] = request.FingerPrintTemplate
	}

	if !tea.BoolValue(util.IsUnset(request.LoginToken)) {
		query["LoginToken"] = request.LoginToken
	}

	if !tea.BoolValue(util.IsUnset(request.Password)) {
		query["Password"] = request.Password
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		query["SessionId"] = request.SessionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetFingerPrintTemplate"),
		Version:     tea.String("2020-10-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetFingerPrintTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetFingerPrintTemplate(request *SetFingerPrintTemplateRequest) (_result *SetFingerPrintTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetFingerPrintTemplateResponse{}
	_body, _err := client.SetFingerPrintTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetFingerPrintTemplateDescriptionWithOptions(request *SetFingerPrintTemplateDescriptionRequest, runtime *util.RuntimeOptions) (_result *SetFingerPrintTemplateDescriptionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientId)) {
		query["ClientId"] = request.ClientId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Index)) {
		query["Index"] = request.Index
	}

	if !tea.BoolValue(util.IsUnset(request.LoginToken)) {
		query["LoginToken"] = request.LoginToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		query["SessionId"] = request.SessionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetFingerPrintTemplateDescription"),
		Version:     tea.String("2020-10-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetFingerPrintTemplateDescriptionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetFingerPrintTemplateDescription(request *SetFingerPrintTemplateDescriptionRequest) (_result *SetFingerPrintTemplateDescriptionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetFingerPrintTemplateDescriptionResponse{}
	_body, _err := client.SetFingerPrintTemplateDescriptionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * The cloud computers that you want to start must be in the Stopped state. After you call this operation, the cloud computers enter the Running state.
 *
 * @param request StartDesktopsRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return StartDesktopsResponse
 */
func (client *Client) StartDesktopsWithOptions(request *StartDesktopsRequest, runtime *util.RuntimeOptions) (_result *StartDesktopsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientId)) {
		query["ClientId"] = request.ClientId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientOS)) {
		query["ClientOS"] = request.ClientOS
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ClientVersion)) {
		query["ClientVersion"] = request.ClientVersion
	}

	if !tea.BoolValue(util.IsUnset(request.DesktopId)) {
		query["DesktopId"] = request.DesktopId
	}

	if !tea.BoolValue(util.IsUnset(request.LoginToken)) {
		query["LoginToken"] = request.LoginToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		query["SessionId"] = request.SessionId
	}

	if !tea.BoolValue(util.IsUnset(request.Uuid)) {
		query["Uuid"] = request.Uuid
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StartDesktops"),
		Version:     tea.String("2020-10-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StartDesktopsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * The cloud computers that you want to start must be in the Stopped state. After you call this operation, the cloud computers enter the Running state.
 *
 * @param request StartDesktopsRequest
 * @return StartDesktopsResponse
 */
func (client *Client) StartDesktops(request *StartDesktopsRequest) (_result *StartDesktopsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartDesktopsResponse{}
	_body, _err := client.StartDesktopsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartRecordContentWithOptions(request *StartRecordContentRequest, runtime *util.RuntimeOptions) (_result *StartRecordContentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientId)) {
		query["ClientId"] = request.ClientId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientOS)) {
		query["ClientOS"] = request.ClientOS
	}

	if !tea.BoolValue(util.IsUnset(request.ClientVersion)) {
		query["ClientVersion"] = request.ClientVersion
	}

	if !tea.BoolValue(util.IsUnset(request.DesktopId)) {
		query["DesktopId"] = request.DesktopId
	}

	if !tea.BoolValue(util.IsUnset(request.FilePath)) {
		query["FilePath"] = request.FilePath
	}

	if !tea.BoolValue(util.IsUnset(request.LoginToken)) {
		query["LoginToken"] = request.LoginToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		query["SessionId"] = request.SessionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StartRecordContent"),
		Version:     tea.String("2020-10-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StartRecordContentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartRecordContent(request *StartRecordContentRequest) (_result *StartRecordContentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartRecordContentResponse{}
	_body, _err := client.StartRecordContentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * The cloud computers that you want to stop must be in the Running state. After you call this operation, the cloud computers enter the Stopped state.
 *
 * @param request StopDesktopsRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return StopDesktopsResponse
 */
func (client *Client) StopDesktopsWithOptions(request *StopDesktopsRequest, runtime *util.RuntimeOptions) (_result *StopDesktopsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientId)) {
		query["ClientId"] = request.ClientId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientOS)) {
		query["ClientOS"] = request.ClientOS
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ClientVersion)) {
		query["ClientVersion"] = request.ClientVersion
	}

	if !tea.BoolValue(util.IsUnset(request.DesktopId)) {
		query["DesktopId"] = request.DesktopId
	}

	if !tea.BoolValue(util.IsUnset(request.LoginToken)) {
		query["LoginToken"] = request.LoginToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		query["SessionId"] = request.SessionId
	}

	if !tea.BoolValue(util.IsUnset(request.SessionToken)) {
		query["SessionToken"] = request.SessionToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StopDesktops"),
		Version:     tea.String("2020-10-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StopDesktopsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * The cloud computers that you want to stop must be in the Running state. After you call this operation, the cloud computers enter the Stopped state.
 *
 * @param request StopDesktopsRequest
 * @return StopDesktopsResponse
 */
func (client *Client) StopDesktops(request *StopDesktopsRequest) (_result *StopDesktopsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopDesktopsResponse{}
	_body, _err := client.StopDesktopsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopRecordContentWithOptions(request *StopRecordContentRequest, runtime *util.RuntimeOptions) (_result *StopRecordContentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientId)) {
		query["ClientId"] = request.ClientId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientOS)) {
		query["ClientOS"] = request.ClientOS
	}

	if !tea.BoolValue(util.IsUnset(request.ClientVersion)) {
		query["ClientVersion"] = request.ClientVersion
	}

	if !tea.BoolValue(util.IsUnset(request.DesktopId)) {
		query["DesktopId"] = request.DesktopId
	}

	if !tea.BoolValue(util.IsUnset(request.LoginToken)) {
		query["LoginToken"] = request.LoginToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		query["SessionId"] = request.SessionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StopRecordContent"),
		Version:     tea.String("2020-10-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StopRecordContentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopRecordContent(request *StopRecordContentRequest) (_result *StopRecordContentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopRecordContentResponse{}
	_body, _err := client.StopRecordContentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UnbindUserDesktopWithOptions(request *UnbindUserDesktopRequest, runtime *util.RuntimeOptions) (_result *UnbindUserDesktopResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientId)) {
		query["ClientId"] = request.ClientId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientType)) {
		query["ClientType"] = request.ClientType
	}

	if !tea.BoolValue(util.IsUnset(request.Force)) {
		query["Force"] = request.Force
	}

	if !tea.BoolValue(util.IsUnset(request.LoginToken)) {
		query["LoginToken"] = request.LoginToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		query["SessionId"] = request.SessionId
	}

	if !tea.BoolValue(util.IsUnset(request.UserDesktopId)) {
		query["UserDesktopId"] = request.UserDesktopId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UnbindUserDesktop"),
		Version:     tea.String("2020-10-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UnbindUserDesktopResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UnbindUserDesktop(request *UnbindUserDesktopRequest) (_result *UnbindUserDesktopResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UnbindUserDesktopResponse{}
	_body, _err := client.UnbindUserDesktopWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) VerifyCredentialWithOptions(request *VerifyCredentialRequest, runtime *util.RuntimeOptions) (_result *VerifyCredentialResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientId)) {
		query["ClientId"] = request.ClientId
	}

	if !tea.BoolValue(util.IsUnset(request.Credential)) {
		query["Credential"] = request.Credential
	}

	if !tea.BoolValue(util.IsUnset(request.CredentialType)) {
		query["CredentialType"] = request.CredentialType
	}

	if !tea.BoolValue(util.IsUnset(request.EncryptedKey)) {
		query["EncryptedKey"] = request.EncryptedKey
	}

	if !tea.BoolValue(util.IsUnset(request.LoginToken)) {
		query["LoginToken"] = request.LoginToken
	}

	if !tea.BoolValue(util.IsUnset(request.OfficeSiteId)) {
		query["OfficeSiteId"] = request.OfficeSiteId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		query["SessionId"] = request.SessionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("VerifyCredential"),
		Version:     tea.String("2020-10-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &VerifyCredentialResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) VerifyCredential(request *VerifyCredentialRequest) (_result *VerifyCredentialResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &VerifyCredentialResponse{}
	_body, _err := client.VerifyCredentialWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
