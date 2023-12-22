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

type GetConnectionTicketRequest struct {
	AppId                *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppInstanceGroupId   *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	AppInstanceId        *string `json:"AppInstanceId,omitempty" xml:"AppInstanceId,omitempty"`
	AppVersion           *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	BizRegionId          *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	ClientId             *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	ClientIp             *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	ClientOS             *string `json:"ClientOS,omitempty" xml:"ClientOS,omitempty"`
	ClientVersion        *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	ConnectionProperties *string `json:"ConnectionProperties,omitempty" xml:"ConnectionProperties,omitempty"`
	EndUserId            *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	LoginRegionId        *string `json:"LoginRegionId,omitempty" xml:"LoginRegionId,omitempty"`
	LoginToken           *string `json:"LoginToken,omitempty" xml:"LoginToken,omitempty"`
	Param                *string `json:"Param,omitempty" xml:"Param,omitempty"`
	ProductType          *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	ResourceId           *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	SessionId            *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	TaskId               *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TenantId             *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	Uuid                 *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s GetConnectionTicketRequest) String() string {
	return tea.Prettify(s)
}

func (s GetConnectionTicketRequest) GoString() string {
	return s.String()
}

func (s *GetConnectionTicketRequest) SetAppId(v string) *GetConnectionTicketRequest {
	s.AppId = &v
	return s
}

func (s *GetConnectionTicketRequest) SetAppInstanceGroupId(v string) *GetConnectionTicketRequest {
	s.AppInstanceGroupId = &v
	return s
}

func (s *GetConnectionTicketRequest) SetAppInstanceId(v string) *GetConnectionTicketRequest {
	s.AppInstanceId = &v
	return s
}

func (s *GetConnectionTicketRequest) SetAppVersion(v string) *GetConnectionTicketRequest {
	s.AppVersion = &v
	return s
}

func (s *GetConnectionTicketRequest) SetBizRegionId(v string) *GetConnectionTicketRequest {
	s.BizRegionId = &v
	return s
}

func (s *GetConnectionTicketRequest) SetClientId(v string) *GetConnectionTicketRequest {
	s.ClientId = &v
	return s
}

func (s *GetConnectionTicketRequest) SetClientIp(v string) *GetConnectionTicketRequest {
	s.ClientIp = &v
	return s
}

func (s *GetConnectionTicketRequest) SetClientOS(v string) *GetConnectionTicketRequest {
	s.ClientOS = &v
	return s
}

func (s *GetConnectionTicketRequest) SetClientVersion(v string) *GetConnectionTicketRequest {
	s.ClientVersion = &v
	return s
}

func (s *GetConnectionTicketRequest) SetConnectionProperties(v string) *GetConnectionTicketRequest {
	s.ConnectionProperties = &v
	return s
}

func (s *GetConnectionTicketRequest) SetEndUserId(v string) *GetConnectionTicketRequest {
	s.EndUserId = &v
	return s
}

func (s *GetConnectionTicketRequest) SetLoginRegionId(v string) *GetConnectionTicketRequest {
	s.LoginRegionId = &v
	return s
}

func (s *GetConnectionTicketRequest) SetLoginToken(v string) *GetConnectionTicketRequest {
	s.LoginToken = &v
	return s
}

func (s *GetConnectionTicketRequest) SetParam(v string) *GetConnectionTicketRequest {
	s.Param = &v
	return s
}

func (s *GetConnectionTicketRequest) SetProductType(v string) *GetConnectionTicketRequest {
	s.ProductType = &v
	return s
}

func (s *GetConnectionTicketRequest) SetResourceId(v string) *GetConnectionTicketRequest {
	s.ResourceId = &v
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

func (s *GetConnectionTicketRequest) SetTenantId(v string) *GetConnectionTicketRequest {
	s.TenantId = &v
	return s
}

func (s *GetConnectionTicketRequest) SetUuid(v string) *GetConnectionTicketRequest {
	s.Uuid = &v
	return s
}

type GetConnectionTicketResponseBody struct {
	AppInstanceGroupId      *string                                `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	AppInstanceId           *string                                `json:"AppInstanceId,omitempty" xml:"AppInstanceId,omitempty"`
	AppInstancePersistentId *string                                `json:"AppInstancePersistentId,omitempty" xml:"AppInstancePersistentId,omitempty"`
	Code                    *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	LoginToken              *string                                `json:"LoginToken,omitempty" xml:"LoginToken,omitempty"`
	Message                 *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	OsType                  *string                                `json:"OsType,omitempty" xml:"OsType,omitempty"`
	Policy                  *GetConnectionTicketResponseBodyPolicy `json:"Policy,omitempty" xml:"Policy,omitempty" type:"Struct"`
	RegionId                *string                                `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Id of the request
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RetryTimes *int32  `json:"RetryTimes,omitempty" xml:"RetryTimes,omitempty"`
	TaskId     *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	TenantId   *int64  `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	Ticket     *string `json:"Ticket,omitempty" xml:"Ticket,omitempty"`
}

func (s GetConnectionTicketResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetConnectionTicketResponseBody) GoString() string {
	return s.String()
}

func (s *GetConnectionTicketResponseBody) SetAppInstanceGroupId(v string) *GetConnectionTicketResponseBody {
	s.AppInstanceGroupId = &v
	return s
}

func (s *GetConnectionTicketResponseBody) SetAppInstanceId(v string) *GetConnectionTicketResponseBody {
	s.AppInstanceId = &v
	return s
}

func (s *GetConnectionTicketResponseBody) SetAppInstancePersistentId(v string) *GetConnectionTicketResponseBody {
	s.AppInstancePersistentId = &v
	return s
}

func (s *GetConnectionTicketResponseBody) SetCode(v string) *GetConnectionTicketResponseBody {
	s.Code = &v
	return s
}

func (s *GetConnectionTicketResponseBody) SetLoginToken(v string) *GetConnectionTicketResponseBody {
	s.LoginToken = &v
	return s
}

func (s *GetConnectionTicketResponseBody) SetMessage(v string) *GetConnectionTicketResponseBody {
	s.Message = &v
	return s
}

func (s *GetConnectionTicketResponseBody) SetOsType(v string) *GetConnectionTicketResponseBody {
	s.OsType = &v
	return s
}

func (s *GetConnectionTicketResponseBody) SetPolicy(v *GetConnectionTicketResponseBodyPolicy) *GetConnectionTicketResponseBody {
	s.Policy = v
	return s
}

func (s *GetConnectionTicketResponseBody) SetRegionId(v string) *GetConnectionTicketResponseBody {
	s.RegionId = &v
	return s
}

func (s *GetConnectionTicketResponseBody) SetRequestId(v string) *GetConnectionTicketResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetConnectionTicketResponseBody) SetRetryTimes(v int32) *GetConnectionTicketResponseBody {
	s.RetryTimes = &v
	return s
}

func (s *GetConnectionTicketResponseBody) SetTaskId(v string) *GetConnectionTicketResponseBody {
	s.TaskId = &v
	return s
}

func (s *GetConnectionTicketResponseBody) SetTaskStatus(v string) *GetConnectionTicketResponseBody {
	s.TaskStatus = &v
	return s
}

func (s *GetConnectionTicketResponseBody) SetTenantId(v int64) *GetConnectionTicketResponseBody {
	s.TenantId = &v
	return s
}

func (s *GetConnectionTicketResponseBody) SetTicket(v string) *GetConnectionTicketResponseBody {
	s.Ticket = &v
	return s
}

type GetConnectionTicketResponseBodyPolicy struct {
	ResolutionAdaptive *string `json:"ResolutionAdaptive,omitempty" xml:"ResolutionAdaptive,omitempty"`
	ResolutionHeight   *int32  `json:"ResolutionHeight,omitempty" xml:"ResolutionHeight,omitempty"`
	ResolutionWidth    *int32  `json:"ResolutionWidth,omitempty" xml:"ResolutionWidth,omitempty"`
}

func (s GetConnectionTicketResponseBodyPolicy) String() string {
	return tea.Prettify(s)
}

func (s GetConnectionTicketResponseBodyPolicy) GoString() string {
	return s.String()
}

func (s *GetConnectionTicketResponseBodyPolicy) SetResolutionAdaptive(v string) *GetConnectionTicketResponseBodyPolicy {
	s.ResolutionAdaptive = &v
	return s
}

func (s *GetConnectionTicketResponseBodyPolicy) SetResolutionHeight(v int32) *GetConnectionTicketResponseBodyPolicy {
	s.ResolutionHeight = &v
	return s
}

func (s *GetConnectionTicketResponseBodyPolicy) SetResolutionWidth(v int32) *GetConnectionTicketResponseBodyPolicy {
	s.ResolutionWidth = &v
	return s
}

type GetConnectionTicketResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetConnectionTicketResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type ListLFUAppRequest struct {
	AliUid             *int64  `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	ApiType            *string `json:"ApiType,omitempty" xml:"ApiType,omitempty"`
	BizRegionId        *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	ClientChannel      *string `json:"ClientChannel,omitempty" xml:"ClientChannel,omitempty"`
	ClientId           *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	ClientIp           *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	ClientOS           *string `json:"ClientOS,omitempty" xml:"ClientOS,omitempty"`
	ClientVersion      *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	EndUserId          *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	ExtendsAccessToken *string `json:"ExtendsAccessToken,omitempty" xml:"ExtendsAccessToken,omitempty"`
	IdpId              *string `json:"IdpId,omitempty" xml:"IdpId,omitempty"`
	LoginRegionId      *string `json:"LoginRegionId,omitempty" xml:"LoginRegionId,omitempty"`
	LoginToken         *string `json:"LoginToken,omitempty" xml:"LoginToken,omitempty"`
	ProductType        *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	RegionId           *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SessionId          *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	TraceId            *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
	Uuid               *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	WyId               *string `json:"WyId,omitempty" xml:"WyId,omitempty"`
}

func (s ListLFUAppRequest) String() string {
	return tea.Prettify(s)
}

func (s ListLFUAppRequest) GoString() string {
	return s.String()
}

func (s *ListLFUAppRequest) SetAliUid(v int64) *ListLFUAppRequest {
	s.AliUid = &v
	return s
}

func (s *ListLFUAppRequest) SetApiType(v string) *ListLFUAppRequest {
	s.ApiType = &v
	return s
}

func (s *ListLFUAppRequest) SetBizRegionId(v string) *ListLFUAppRequest {
	s.BizRegionId = &v
	return s
}

func (s *ListLFUAppRequest) SetClientChannel(v string) *ListLFUAppRequest {
	s.ClientChannel = &v
	return s
}

func (s *ListLFUAppRequest) SetClientId(v string) *ListLFUAppRequest {
	s.ClientId = &v
	return s
}

func (s *ListLFUAppRequest) SetClientIp(v string) *ListLFUAppRequest {
	s.ClientIp = &v
	return s
}

func (s *ListLFUAppRequest) SetClientOS(v string) *ListLFUAppRequest {
	s.ClientOS = &v
	return s
}

func (s *ListLFUAppRequest) SetClientVersion(v string) *ListLFUAppRequest {
	s.ClientVersion = &v
	return s
}

func (s *ListLFUAppRequest) SetEndUserId(v string) *ListLFUAppRequest {
	s.EndUserId = &v
	return s
}

func (s *ListLFUAppRequest) SetExtendsAccessToken(v string) *ListLFUAppRequest {
	s.ExtendsAccessToken = &v
	return s
}

func (s *ListLFUAppRequest) SetIdpId(v string) *ListLFUAppRequest {
	s.IdpId = &v
	return s
}

func (s *ListLFUAppRequest) SetLoginRegionId(v string) *ListLFUAppRequest {
	s.LoginRegionId = &v
	return s
}

func (s *ListLFUAppRequest) SetLoginToken(v string) *ListLFUAppRequest {
	s.LoginToken = &v
	return s
}

func (s *ListLFUAppRequest) SetProductType(v string) *ListLFUAppRequest {
	s.ProductType = &v
	return s
}

func (s *ListLFUAppRequest) SetRegionId(v string) *ListLFUAppRequest {
	s.RegionId = &v
	return s
}

func (s *ListLFUAppRequest) SetSessionId(v string) *ListLFUAppRequest {
	s.SessionId = &v
	return s
}

func (s *ListLFUAppRequest) SetTraceId(v string) *ListLFUAppRequest {
	s.TraceId = &v
	return s
}

func (s *ListLFUAppRequest) SetUuid(v string) *ListLFUAppRequest {
	s.Uuid = &v
	return s
}

func (s *ListLFUAppRequest) SetWyId(v string) *ListLFUAppRequest {
	s.WyId = &v
	return s
}

type ListLFUAppResponseBody struct {
	Code           *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Count          *int32                        `json:"Count,omitempty" xml:"Count,omitempty"`
	Data           []*ListLFUAppResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	HttpStatusCode *int32                        `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListLFUAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListLFUAppResponseBody) GoString() string {
	return s.String()
}

func (s *ListLFUAppResponseBody) SetCode(v string) *ListLFUAppResponseBody {
	s.Code = &v
	return s
}

func (s *ListLFUAppResponseBody) SetCount(v int32) *ListLFUAppResponseBody {
	s.Count = &v
	return s
}

func (s *ListLFUAppResponseBody) SetData(v []*ListLFUAppResponseBodyData) *ListLFUAppResponseBody {
	s.Data = v
	return s
}

func (s *ListLFUAppResponseBody) SetHttpStatusCode(v int32) *ListLFUAppResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListLFUAppResponseBody) SetMessage(v string) *ListLFUAppResponseBody {
	s.Message = &v
	return s
}

func (s *ListLFUAppResponseBody) SetRequestId(v string) *ListLFUAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLFUAppResponseBody) SetSuccess(v bool) *ListLFUAppResponseBody {
	s.Success = &v
	return s
}

type ListLFUAppResponseBodyData struct {
	AppId          *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppName        *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	AppVersion     *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	AppVersionName *string `json:"AppVersionName,omitempty" xml:"AppVersionName,omitempty"`
	IconUrl        *string `json:"IconUrl,omitempty" xml:"IconUrl,omitempty"`
	IsAuth         *bool   `json:"IsAuth,omitempty" xml:"IsAuth,omitempty"`
	OsType         *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListLFUAppResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListLFUAppResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListLFUAppResponseBodyData) SetAppId(v string) *ListLFUAppResponseBodyData {
	s.AppId = &v
	return s
}

func (s *ListLFUAppResponseBodyData) SetAppName(v string) *ListLFUAppResponseBodyData {
	s.AppName = &v
	return s
}

func (s *ListLFUAppResponseBodyData) SetAppVersion(v string) *ListLFUAppResponseBodyData {
	s.AppVersion = &v
	return s
}

func (s *ListLFUAppResponseBodyData) SetAppVersionName(v string) *ListLFUAppResponseBodyData {
	s.AppVersionName = &v
	return s
}

func (s *ListLFUAppResponseBodyData) SetIconUrl(v string) *ListLFUAppResponseBodyData {
	s.IconUrl = &v
	return s
}

func (s *ListLFUAppResponseBodyData) SetIsAuth(v bool) *ListLFUAppResponseBodyData {
	s.IsAuth = &v
	return s
}

func (s *ListLFUAppResponseBodyData) SetOsType(v string) *ListLFUAppResponseBodyData {
	s.OsType = &v
	return s
}

func (s *ListLFUAppResponseBodyData) SetRequestId(v string) *ListLFUAppResponseBodyData {
	s.RequestId = &v
	return s
}

type ListLFUAppResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListLFUAppResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListLFUAppResponse) String() string {
	return tea.Prettify(s)
}

func (s ListLFUAppResponse) GoString() string {
	return s.String()
}

func (s *ListLFUAppResponse) SetHeaders(v map[string]*string) *ListLFUAppResponse {
	s.Headers = v
	return s
}

func (s *ListLFUAppResponse) SetStatusCode(v int32) *ListLFUAppResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLFUAppResponse) SetBody(v *ListLFUAppResponseBody) *ListLFUAppResponse {
	s.Body = v
	return s
}

type ListPublishedAppInfoRequest struct {
	AppName       *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	BizRegionId   *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	CategoryId    *int64  `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	CategoryType  *int64  `json:"CategoryType,omitempty" xml:"CategoryType,omitempty"`
	ClientId      *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	ClientIp      *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	ClientOS      *string `json:"ClientOS,omitempty" xml:"ClientOS,omitempty"`
	ClientVersion *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	EndUserId     *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	LoginRegionId *string `json:"LoginRegionId,omitempty" xml:"LoginRegionId,omitempty"`
	LoginToken    *string `json:"LoginToken,omitempty" xml:"LoginToken,omitempty"`
	OrderParam    *string `json:"OrderParam,omitempty" xml:"OrderParam,omitempty"`
	ProductType   *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	SessionId     *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	SortType      *string `json:"SortType,omitempty" xml:"SortType,omitempty"`
}

func (s ListPublishedAppInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPublishedAppInfoRequest) GoString() string {
	return s.String()
}

func (s *ListPublishedAppInfoRequest) SetAppName(v string) *ListPublishedAppInfoRequest {
	s.AppName = &v
	return s
}

func (s *ListPublishedAppInfoRequest) SetBizRegionId(v string) *ListPublishedAppInfoRequest {
	s.BizRegionId = &v
	return s
}

func (s *ListPublishedAppInfoRequest) SetCategoryId(v int64) *ListPublishedAppInfoRequest {
	s.CategoryId = &v
	return s
}

func (s *ListPublishedAppInfoRequest) SetCategoryType(v int64) *ListPublishedAppInfoRequest {
	s.CategoryType = &v
	return s
}

func (s *ListPublishedAppInfoRequest) SetClientId(v string) *ListPublishedAppInfoRequest {
	s.ClientId = &v
	return s
}

func (s *ListPublishedAppInfoRequest) SetClientIp(v string) *ListPublishedAppInfoRequest {
	s.ClientIp = &v
	return s
}

func (s *ListPublishedAppInfoRequest) SetClientOS(v string) *ListPublishedAppInfoRequest {
	s.ClientOS = &v
	return s
}

func (s *ListPublishedAppInfoRequest) SetClientVersion(v string) *ListPublishedAppInfoRequest {
	s.ClientVersion = &v
	return s
}

func (s *ListPublishedAppInfoRequest) SetEndUserId(v string) *ListPublishedAppInfoRequest {
	s.EndUserId = &v
	return s
}

func (s *ListPublishedAppInfoRequest) SetLoginRegionId(v string) *ListPublishedAppInfoRequest {
	s.LoginRegionId = &v
	return s
}

func (s *ListPublishedAppInfoRequest) SetLoginToken(v string) *ListPublishedAppInfoRequest {
	s.LoginToken = &v
	return s
}

func (s *ListPublishedAppInfoRequest) SetOrderParam(v string) *ListPublishedAppInfoRequest {
	s.OrderParam = &v
	return s
}

func (s *ListPublishedAppInfoRequest) SetProductType(v string) *ListPublishedAppInfoRequest {
	s.ProductType = &v
	return s
}

func (s *ListPublishedAppInfoRequest) SetSessionId(v string) *ListPublishedAppInfoRequest {
	s.SessionId = &v
	return s
}

func (s *ListPublishedAppInfoRequest) SetSortType(v string) *ListPublishedAppInfoRequest {
	s.SortType = &v
	return s
}

type ListPublishedAppInfoResponseBody struct {
	// appModels
	AppModels []*ListPublishedAppInfoResponseBodyAppModels `json:"AppModels,omitempty" xml:"AppModels,omitempty" type:"Repeated"`
	NextToken *string                                      `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListPublishedAppInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPublishedAppInfoResponseBody) GoString() string {
	return s.String()
}

func (s *ListPublishedAppInfoResponseBody) SetAppModels(v []*ListPublishedAppInfoResponseBodyAppModels) *ListPublishedAppInfoResponseBody {
	s.AppModels = v
	return s
}

func (s *ListPublishedAppInfoResponseBody) SetNextToken(v string) *ListPublishedAppInfoResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListPublishedAppInfoResponseBody) SetRequestId(v string) *ListPublishedAppInfoResponseBody {
	s.RequestId = &v
	return s
}

type ListPublishedAppInfoResponseBodyAppModels struct {
	AppCenterImageId *string `json:"AppCenterImageId,omitempty" xml:"AppCenterImageId,omitempty"`
	AppId            *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppName          *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	AppThemeColor    *string `json:"AppThemeColor,omitempty" xml:"AppThemeColor,omitempty"`
	AppVersion       *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	AppVersionName   *string `json:"AppVersionName,omitempty" xml:"AppVersionName,omitempty"`
	AuthTime         *string `json:"AuthTime,omitempty" xml:"AuthTime,omitempty"`
	CategoryId       *int64  `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	CategoryType     *int64  `json:"CategoryType,omitempty" xml:"CategoryType,omitempty"`
	IconUrl          *string `json:"IconUrl,omitempty" xml:"IconUrl,omitempty"`
	IsAuth           *bool   `json:"IsAuth,omitempty" xml:"IsAuth,omitempty"`
	UsedInSession    *bool   `json:"UsedInSession,omitempty" xml:"UsedInSession,omitempty"`
}

func (s ListPublishedAppInfoResponseBodyAppModels) String() string {
	return tea.Prettify(s)
}

func (s ListPublishedAppInfoResponseBodyAppModels) GoString() string {
	return s.String()
}

func (s *ListPublishedAppInfoResponseBodyAppModels) SetAppCenterImageId(v string) *ListPublishedAppInfoResponseBodyAppModels {
	s.AppCenterImageId = &v
	return s
}

func (s *ListPublishedAppInfoResponseBodyAppModels) SetAppId(v string) *ListPublishedAppInfoResponseBodyAppModels {
	s.AppId = &v
	return s
}

func (s *ListPublishedAppInfoResponseBodyAppModels) SetAppName(v string) *ListPublishedAppInfoResponseBodyAppModels {
	s.AppName = &v
	return s
}

func (s *ListPublishedAppInfoResponseBodyAppModels) SetAppThemeColor(v string) *ListPublishedAppInfoResponseBodyAppModels {
	s.AppThemeColor = &v
	return s
}

func (s *ListPublishedAppInfoResponseBodyAppModels) SetAppVersion(v string) *ListPublishedAppInfoResponseBodyAppModels {
	s.AppVersion = &v
	return s
}

func (s *ListPublishedAppInfoResponseBodyAppModels) SetAppVersionName(v string) *ListPublishedAppInfoResponseBodyAppModels {
	s.AppVersionName = &v
	return s
}

func (s *ListPublishedAppInfoResponseBodyAppModels) SetAuthTime(v string) *ListPublishedAppInfoResponseBodyAppModels {
	s.AuthTime = &v
	return s
}

func (s *ListPublishedAppInfoResponseBodyAppModels) SetCategoryId(v int64) *ListPublishedAppInfoResponseBodyAppModels {
	s.CategoryId = &v
	return s
}

func (s *ListPublishedAppInfoResponseBodyAppModels) SetCategoryType(v int64) *ListPublishedAppInfoResponseBodyAppModels {
	s.CategoryType = &v
	return s
}

func (s *ListPublishedAppInfoResponseBodyAppModels) SetIconUrl(v string) *ListPublishedAppInfoResponseBodyAppModels {
	s.IconUrl = &v
	return s
}

func (s *ListPublishedAppInfoResponseBodyAppModels) SetIsAuth(v bool) *ListPublishedAppInfoResponseBodyAppModels {
	s.IsAuth = &v
	return s
}

func (s *ListPublishedAppInfoResponseBodyAppModels) SetUsedInSession(v bool) *ListPublishedAppInfoResponseBodyAppModels {
	s.UsedInSession = &v
	return s
}

type ListPublishedAppInfoResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListPublishedAppInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListPublishedAppInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPublishedAppInfoResponse) GoString() string {
	return s.String()
}

func (s *ListPublishedAppInfoResponse) SetHeaders(v map[string]*string) *ListPublishedAppInfoResponse {
	s.Headers = v
	return s
}

func (s *ListPublishedAppInfoResponse) SetStatusCode(v int32) *ListPublishedAppInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPublishedAppInfoResponse) SetBody(v *ListPublishedAppInfoResponseBody) *ListPublishedAppInfoResponse {
	s.Body = v
	return s
}

type ListRunningAppsRequest struct {
	BizRegionId   *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	ClientId      *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	ClientIp      *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	ClientOS      *string `json:"ClientOS,omitempty" xml:"ClientOS,omitempty"`
	ClientVersion *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	EndUserId     *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	LoginRegionId *string `json:"LoginRegionId,omitempty" xml:"LoginRegionId,omitempty"`
	LoginToken    *string `json:"LoginToken,omitempty" xml:"LoginToken,omitempty"`
	ProductType   *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	SessionId     *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	TenantId      *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	Uuid          *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s ListRunningAppsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRunningAppsRequest) GoString() string {
	return s.String()
}

func (s *ListRunningAppsRequest) SetBizRegionId(v string) *ListRunningAppsRequest {
	s.BizRegionId = &v
	return s
}

func (s *ListRunningAppsRequest) SetClientId(v string) *ListRunningAppsRequest {
	s.ClientId = &v
	return s
}

func (s *ListRunningAppsRequest) SetClientIp(v string) *ListRunningAppsRequest {
	s.ClientIp = &v
	return s
}

func (s *ListRunningAppsRequest) SetClientOS(v string) *ListRunningAppsRequest {
	s.ClientOS = &v
	return s
}

func (s *ListRunningAppsRequest) SetClientVersion(v string) *ListRunningAppsRequest {
	s.ClientVersion = &v
	return s
}

func (s *ListRunningAppsRequest) SetEndUserId(v string) *ListRunningAppsRequest {
	s.EndUserId = &v
	return s
}

func (s *ListRunningAppsRequest) SetLoginRegionId(v string) *ListRunningAppsRequest {
	s.LoginRegionId = &v
	return s
}

func (s *ListRunningAppsRequest) SetLoginToken(v string) *ListRunningAppsRequest {
	s.LoginToken = &v
	return s
}

func (s *ListRunningAppsRequest) SetProductType(v string) *ListRunningAppsRequest {
	s.ProductType = &v
	return s
}

func (s *ListRunningAppsRequest) SetSessionId(v string) *ListRunningAppsRequest {
	s.SessionId = &v
	return s
}

func (s *ListRunningAppsRequest) SetTenantId(v string) *ListRunningAppsRequest {
	s.TenantId = &v
	return s
}

func (s *ListRunningAppsRequest) SetUuid(v string) *ListRunningAppsRequest {
	s.Uuid = &v
	return s
}

type ListRunningAppsResponseBody struct {
	// Id of the request
	RequestId        *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RunningCloudApps []*ListRunningAppsResponseBodyRunningCloudApps `json:"RunningCloudApps,omitempty" xml:"RunningCloudApps,omitempty" type:"Repeated"`
}

func (s ListRunningAppsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRunningAppsResponseBody) GoString() string {
	return s.String()
}

func (s *ListRunningAppsResponseBody) SetRequestId(v string) *ListRunningAppsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRunningAppsResponseBody) SetRunningCloudApps(v []*ListRunningAppsResponseBodyRunningCloudApps) *ListRunningAppsResponseBody {
	s.RunningCloudApps = v
	return s
}

type ListRunningAppsResponseBodyRunningCloudApps struct {
	AppId              *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	AppInstanceId      *string `json:"AppInstanceId,omitempty" xml:"AppInstanceId,omitempty"`
	AppName            *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	AppVersion         *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	AppVersionName     *string `json:"AppVersionName,omitempty" xml:"AppVersionName,omitempty"`
	Duration           *int64  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	IconUrl            *string `json:"IconUrl,omitempty" xml:"IconUrl,omitempty"`
	OsType             *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
	RegionId           *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	StartTime          *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListRunningAppsResponseBodyRunningCloudApps) String() string {
	return tea.Prettify(s)
}

func (s ListRunningAppsResponseBodyRunningCloudApps) GoString() string {
	return s.String()
}

func (s *ListRunningAppsResponseBodyRunningCloudApps) SetAppId(v string) *ListRunningAppsResponseBodyRunningCloudApps {
	s.AppId = &v
	return s
}

func (s *ListRunningAppsResponseBodyRunningCloudApps) SetAppInstanceGroupId(v string) *ListRunningAppsResponseBodyRunningCloudApps {
	s.AppInstanceGroupId = &v
	return s
}

func (s *ListRunningAppsResponseBodyRunningCloudApps) SetAppInstanceId(v string) *ListRunningAppsResponseBodyRunningCloudApps {
	s.AppInstanceId = &v
	return s
}

func (s *ListRunningAppsResponseBodyRunningCloudApps) SetAppName(v string) *ListRunningAppsResponseBodyRunningCloudApps {
	s.AppName = &v
	return s
}

func (s *ListRunningAppsResponseBodyRunningCloudApps) SetAppVersion(v string) *ListRunningAppsResponseBodyRunningCloudApps {
	s.AppVersion = &v
	return s
}

func (s *ListRunningAppsResponseBodyRunningCloudApps) SetAppVersionName(v string) *ListRunningAppsResponseBodyRunningCloudApps {
	s.AppVersionName = &v
	return s
}

func (s *ListRunningAppsResponseBodyRunningCloudApps) SetDuration(v int64) *ListRunningAppsResponseBodyRunningCloudApps {
	s.Duration = &v
	return s
}

func (s *ListRunningAppsResponseBodyRunningCloudApps) SetIconUrl(v string) *ListRunningAppsResponseBodyRunningCloudApps {
	s.IconUrl = &v
	return s
}

func (s *ListRunningAppsResponseBodyRunningCloudApps) SetOsType(v string) *ListRunningAppsResponseBodyRunningCloudApps {
	s.OsType = &v
	return s
}

func (s *ListRunningAppsResponseBodyRunningCloudApps) SetRegionId(v string) *ListRunningAppsResponseBodyRunningCloudApps {
	s.RegionId = &v
	return s
}

func (s *ListRunningAppsResponseBodyRunningCloudApps) SetStartTime(v string) *ListRunningAppsResponseBodyRunningCloudApps {
	s.StartTime = &v
	return s
}

type ListRunningAppsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListRunningAppsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListRunningAppsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRunningAppsResponse) GoString() string {
	return s.String()
}

func (s *ListRunningAppsResponse) SetHeaders(v map[string]*string) *ListRunningAppsResponse {
	s.Headers = v
	return s
}

func (s *ListRunningAppsResponse) SetStatusCode(v int32) *ListRunningAppsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRunningAppsResponse) SetBody(v *ListRunningAppsResponseBody) *ListRunningAppsResponse {
	s.Body = v
	return s
}

type StopAppRequest struct {
	AliUid             *int64  `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	ApiType            *string `json:"ApiType,omitempty" xml:"ApiType,omitempty"`
	AppId              *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	AppInstanceId      *string `json:"AppInstanceId,omitempty" xml:"AppInstanceId,omitempty"`
	BizRegionId        *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	ClientChannel      *string `json:"ClientChannel,omitempty" xml:"ClientChannel,omitempty"`
	ClientId           *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	ClientIp           *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	ClientOS           *string `json:"ClientOS,omitempty" xml:"ClientOS,omitempty"`
	ClientVersion      *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	EndUserId          *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	ForceStop          *bool   `json:"ForceStop,omitempty" xml:"ForceStop,omitempty"`
	IdpId              *string `json:"IdpId,omitempty" xml:"IdpId,omitempty"`
	LoginRegionId      *string `json:"LoginRegionId,omitempty" xml:"LoginRegionId,omitempty"`
	LoginToken         *string `json:"LoginToken,omitempty" xml:"LoginToken,omitempty"`
	ProductType        *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	RegionId           *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SessionId          *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	Uuid               *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	WyId               *string `json:"WyId,omitempty" xml:"WyId,omitempty"`
}

func (s StopAppRequest) String() string {
	return tea.Prettify(s)
}

func (s StopAppRequest) GoString() string {
	return s.String()
}

func (s *StopAppRequest) SetAliUid(v int64) *StopAppRequest {
	s.AliUid = &v
	return s
}

func (s *StopAppRequest) SetApiType(v string) *StopAppRequest {
	s.ApiType = &v
	return s
}

func (s *StopAppRequest) SetAppId(v string) *StopAppRequest {
	s.AppId = &v
	return s
}

func (s *StopAppRequest) SetAppInstanceGroupId(v string) *StopAppRequest {
	s.AppInstanceGroupId = &v
	return s
}

func (s *StopAppRequest) SetAppInstanceId(v string) *StopAppRequest {
	s.AppInstanceId = &v
	return s
}

func (s *StopAppRequest) SetBizRegionId(v string) *StopAppRequest {
	s.BizRegionId = &v
	return s
}

func (s *StopAppRequest) SetClientChannel(v string) *StopAppRequest {
	s.ClientChannel = &v
	return s
}

func (s *StopAppRequest) SetClientId(v string) *StopAppRequest {
	s.ClientId = &v
	return s
}

func (s *StopAppRequest) SetClientIp(v string) *StopAppRequest {
	s.ClientIp = &v
	return s
}

func (s *StopAppRequest) SetClientOS(v string) *StopAppRequest {
	s.ClientOS = &v
	return s
}

func (s *StopAppRequest) SetClientVersion(v string) *StopAppRequest {
	s.ClientVersion = &v
	return s
}

func (s *StopAppRequest) SetEndUserId(v string) *StopAppRequest {
	s.EndUserId = &v
	return s
}

func (s *StopAppRequest) SetForceStop(v bool) *StopAppRequest {
	s.ForceStop = &v
	return s
}

func (s *StopAppRequest) SetIdpId(v string) *StopAppRequest {
	s.IdpId = &v
	return s
}

func (s *StopAppRequest) SetLoginRegionId(v string) *StopAppRequest {
	s.LoginRegionId = &v
	return s
}

func (s *StopAppRequest) SetLoginToken(v string) *StopAppRequest {
	s.LoginToken = &v
	return s
}

func (s *StopAppRequest) SetProductType(v string) *StopAppRequest {
	s.ProductType = &v
	return s
}

func (s *StopAppRequest) SetRegionId(v string) *StopAppRequest {
	s.RegionId = &v
	return s
}

func (s *StopAppRequest) SetSessionId(v string) *StopAppRequest {
	s.SessionId = &v
	return s
}

func (s *StopAppRequest) SetUuid(v string) *StopAppRequest {
	s.Uuid = &v
	return s
}

func (s *StopAppRequest) SetWyId(v string) *StopAppRequest {
	s.WyId = &v
	return s
}

type StopAppResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopAppResponseBody) GoString() string {
	return s.String()
}

func (s *StopAppResponseBody) SetRequestId(v string) *StopAppResponseBody {
	s.RequestId = &v
	return s
}

type StopAppResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StopAppResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopAppResponse) String() string {
	return tea.Prettify(s)
}

func (s StopAppResponse) GoString() string {
	return s.String()
}

func (s *StopAppResponse) SetHeaders(v map[string]*string) *StopAppResponse {
	s.Headers = v
	return s
}

func (s *StopAppResponse) SetStatusCode(v int32) *StopAppResponse {
	s.StatusCode = &v
	return s
}

func (s *StopAppResponse) SetBody(v *StopAppResponseBody) *StopAppResponse {
	s.Body = v
	return s
}

type UnbindRequest struct {
	AppId              *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	AppInstanceId      *string `json:"AppInstanceId,omitempty" xml:"AppInstanceId,omitempty"`
	ClientId           *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	ClientIp           *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	ClientOS           *string `json:"ClientOS,omitempty" xml:"ClientOS,omitempty"`
	ClientVersion      *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	EndUserId          *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	LoginRegionId      *string `json:"LoginRegionId,omitempty" xml:"LoginRegionId,omitempty"`
	LoginToken         *string `json:"LoginToken,omitempty" xml:"LoginToken,omitempty"`
	ProductType        *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	SessionId          *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	TenantId           *int64  `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s UnbindRequest) String() string {
	return tea.Prettify(s)
}

func (s UnbindRequest) GoString() string {
	return s.String()
}

func (s *UnbindRequest) SetAppId(v string) *UnbindRequest {
	s.AppId = &v
	return s
}

func (s *UnbindRequest) SetAppInstanceGroupId(v string) *UnbindRequest {
	s.AppInstanceGroupId = &v
	return s
}

func (s *UnbindRequest) SetAppInstanceId(v string) *UnbindRequest {
	s.AppInstanceId = &v
	return s
}

func (s *UnbindRequest) SetClientId(v string) *UnbindRequest {
	s.ClientId = &v
	return s
}

func (s *UnbindRequest) SetClientIp(v string) *UnbindRequest {
	s.ClientIp = &v
	return s
}

func (s *UnbindRequest) SetClientOS(v string) *UnbindRequest {
	s.ClientOS = &v
	return s
}

func (s *UnbindRequest) SetClientVersion(v string) *UnbindRequest {
	s.ClientVersion = &v
	return s
}

func (s *UnbindRequest) SetEndUserId(v string) *UnbindRequest {
	s.EndUserId = &v
	return s
}

func (s *UnbindRequest) SetLoginRegionId(v string) *UnbindRequest {
	s.LoginRegionId = &v
	return s
}

func (s *UnbindRequest) SetLoginToken(v string) *UnbindRequest {
	s.LoginToken = &v
	return s
}

func (s *UnbindRequest) SetProductType(v string) *UnbindRequest {
	s.ProductType = &v
	return s
}

func (s *UnbindRequest) SetSessionId(v string) *UnbindRequest {
	s.SessionId = &v
	return s
}

func (s *UnbindRequest) SetTenantId(v int64) *UnbindRequest {
	s.TenantId = &v
	return s
}

type UnbindResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnbindResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnbindResponseBody) GoString() string {
	return s.String()
}

func (s *UnbindResponseBody) SetRequestId(v string) *UnbindResponseBody {
	s.RequestId = &v
	return s
}

type UnbindResponse struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UnbindResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UnbindResponse) String() string {
	return tea.Prettify(s)
}

func (s UnbindResponse) GoString() string {
	return s.String()
}

func (s *UnbindResponse) SetHeaders(v map[string]*string) *UnbindResponse {
	s.Headers = v
	return s
}

func (s *UnbindResponse) SetStatusCode(v int32) *UnbindResponse {
	s.StatusCode = &v
	return s
}

func (s *UnbindResponse) SetBody(v *UnbindResponseBody) *UnbindResponse {
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
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("appstream-center"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) GetConnectionTicketWithOptions(request *GetConnectionTicketRequest, runtime *util.RuntimeOptions) (_result *GetConnectionTicketResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.AppInstanceGroupId)) {
		body["AppInstanceGroupId"] = request.AppInstanceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.AppInstanceId)) {
		body["AppInstanceId"] = request.AppInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.AppVersion)) {
		body["AppVersion"] = request.AppVersion
	}

	if !tea.BoolValue(util.IsUnset(request.BizRegionId)) {
		body["BizRegionId"] = request.BizRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientId)) {
		body["ClientId"] = request.ClientId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientIp)) {
		body["ClientIp"] = request.ClientIp
	}

	if !tea.BoolValue(util.IsUnset(request.ClientOS)) {
		body["ClientOS"] = request.ClientOS
	}

	if !tea.BoolValue(util.IsUnset(request.ClientVersion)) {
		body["ClientVersion"] = request.ClientVersion
	}

	if !tea.BoolValue(util.IsUnset(request.ConnectionProperties)) {
		body["ConnectionProperties"] = request.ConnectionProperties
	}

	if !tea.BoolValue(util.IsUnset(request.EndUserId)) {
		body["EndUserId"] = request.EndUserId
	}

	if !tea.BoolValue(util.IsUnset(request.LoginRegionId)) {
		body["LoginRegionId"] = request.LoginRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.LoginToken)) {
		body["LoginToken"] = request.LoginToken
	}

	if !tea.BoolValue(util.IsUnset(request.Param)) {
		body["Param"] = request.Param
	}

	if !tea.BoolValue(util.IsUnset(request.ProductType)) {
		body["ProductType"] = request.ProductType
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		body["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		body["SessionId"] = request.SessionId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		body["TaskId"] = request.TaskId
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["TenantId"] = request.TenantId
	}

	if !tea.BoolValue(util.IsUnset(request.Uuid)) {
		body["Uuid"] = request.Uuid
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetConnectionTicket"),
		Version:     tea.String("2021-09-03"),
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

func (client *Client) ListLFUAppWithOptions(request *ListLFUAppRequest, runtime *util.RuntimeOptions) (_result *ListLFUAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AliUid)) {
		body["AliUid"] = request.AliUid
	}

	if !tea.BoolValue(util.IsUnset(request.ApiType)) {
		body["ApiType"] = request.ApiType
	}

	if !tea.BoolValue(util.IsUnset(request.BizRegionId)) {
		body["BizRegionId"] = request.BizRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientChannel)) {
		body["ClientChannel"] = request.ClientChannel
	}

	if !tea.BoolValue(util.IsUnset(request.ClientId)) {
		body["ClientId"] = request.ClientId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientIp)) {
		body["ClientIp"] = request.ClientIp
	}

	if !tea.BoolValue(util.IsUnset(request.ClientOS)) {
		body["ClientOS"] = request.ClientOS
	}

	if !tea.BoolValue(util.IsUnset(request.ClientVersion)) {
		body["ClientVersion"] = request.ClientVersion
	}

	if !tea.BoolValue(util.IsUnset(request.EndUserId)) {
		body["EndUserId"] = request.EndUserId
	}

	if !tea.BoolValue(util.IsUnset(request.ExtendsAccessToken)) {
		body["ExtendsAccessToken"] = request.ExtendsAccessToken
	}

	if !tea.BoolValue(util.IsUnset(request.IdpId)) {
		body["IdpId"] = request.IdpId
	}

	if !tea.BoolValue(util.IsUnset(request.LoginRegionId)) {
		body["LoginRegionId"] = request.LoginRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.LoginToken)) {
		body["LoginToken"] = request.LoginToken
	}

	if !tea.BoolValue(util.IsUnset(request.ProductType)) {
		body["ProductType"] = request.ProductType
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		body["SessionId"] = request.SessionId
	}

	if !tea.BoolValue(util.IsUnset(request.TraceId)) {
		body["TraceId"] = request.TraceId
	}

	if !tea.BoolValue(util.IsUnset(request.Uuid)) {
		body["Uuid"] = request.Uuid
	}

	if !tea.BoolValue(util.IsUnset(request.WyId)) {
		body["WyId"] = request.WyId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListLFUApp"),
		Version:     tea.String("2021-09-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListLFUAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListLFUApp(request *ListLFUAppRequest) (_result *ListLFUAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListLFUAppResponse{}
	_body, _err := client.ListLFUAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListPublishedAppInfoWithOptions(request *ListPublishedAppInfoRequest, runtime *util.RuntimeOptions) (_result *ListPublishedAppInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		query["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.BizRegionId)) {
		query["BizRegionId"] = request.BizRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.CategoryId)) {
		query["CategoryId"] = request.CategoryId
	}

	if !tea.BoolValue(util.IsUnset(request.CategoryType)) {
		query["CategoryType"] = request.CategoryType
	}

	if !tea.BoolValue(util.IsUnset(request.ClientId)) {
		query["ClientId"] = request.ClientId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientIp)) {
		query["ClientIp"] = request.ClientIp
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

	if !tea.BoolValue(util.IsUnset(request.LoginRegionId)) {
		query["LoginRegionId"] = request.LoginRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.LoginToken)) {
		query["LoginToken"] = request.LoginToken
	}

	if !tea.BoolValue(util.IsUnset(request.OrderParam)) {
		query["OrderParam"] = request.OrderParam
	}

	if !tea.BoolValue(util.IsUnset(request.ProductType)) {
		query["ProductType"] = request.ProductType
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		query["SessionId"] = request.SessionId
	}

	if !tea.BoolValue(util.IsUnset(request.SortType)) {
		query["SortType"] = request.SortType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListPublishedAppInfo"),
		Version:     tea.String("2021-09-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListPublishedAppInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListPublishedAppInfo(request *ListPublishedAppInfoRequest) (_result *ListPublishedAppInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListPublishedAppInfoResponse{}
	_body, _err := client.ListPublishedAppInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListRunningAppsWithOptions(request *ListRunningAppsRequest, runtime *util.RuntimeOptions) (_result *ListRunningAppsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizRegionId)) {
		query["BizRegionId"] = request.BizRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientId)) {
		query["ClientId"] = request.ClientId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientIp)) {
		query["ClientIp"] = request.ClientIp
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

	if !tea.BoolValue(util.IsUnset(request.LoginRegionId)) {
		query["LoginRegionId"] = request.LoginRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.LoginToken)) {
		query["LoginToken"] = request.LoginToken
	}

	if !tea.BoolValue(util.IsUnset(request.ProductType)) {
		query["ProductType"] = request.ProductType
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		query["SessionId"] = request.SessionId
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		query["TenantId"] = request.TenantId
	}

	if !tea.BoolValue(util.IsUnset(request.Uuid)) {
		query["Uuid"] = request.Uuid
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListRunningApps"),
		Version:     tea.String("2021-09-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListRunningAppsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListRunningApps(request *ListRunningAppsRequest) (_result *ListRunningAppsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListRunningAppsResponse{}
	_body, _err := client.ListRunningAppsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopAppWithOptions(request *StopAppRequest, runtime *util.RuntimeOptions) (_result *StopAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AliUid)) {
		body["AliUid"] = request.AliUid
	}

	if !tea.BoolValue(util.IsUnset(request.ApiType)) {
		body["ApiType"] = request.ApiType
	}

	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.AppInstanceGroupId)) {
		body["AppInstanceGroupId"] = request.AppInstanceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.AppInstanceId)) {
		body["AppInstanceId"] = request.AppInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.BizRegionId)) {
		body["BizRegionId"] = request.BizRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientChannel)) {
		body["ClientChannel"] = request.ClientChannel
	}

	if !tea.BoolValue(util.IsUnset(request.ClientId)) {
		body["ClientId"] = request.ClientId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientIp)) {
		body["ClientIp"] = request.ClientIp
	}

	if !tea.BoolValue(util.IsUnset(request.ClientOS)) {
		body["ClientOS"] = request.ClientOS
	}

	if !tea.BoolValue(util.IsUnset(request.ClientVersion)) {
		body["ClientVersion"] = request.ClientVersion
	}

	if !tea.BoolValue(util.IsUnset(request.EndUserId)) {
		body["EndUserId"] = request.EndUserId
	}

	if !tea.BoolValue(util.IsUnset(request.ForceStop)) {
		body["ForceStop"] = request.ForceStop
	}

	if !tea.BoolValue(util.IsUnset(request.IdpId)) {
		body["IdpId"] = request.IdpId
	}

	if !tea.BoolValue(util.IsUnset(request.LoginRegionId)) {
		body["LoginRegionId"] = request.LoginRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.LoginToken)) {
		body["LoginToken"] = request.LoginToken
	}

	if !tea.BoolValue(util.IsUnset(request.ProductType)) {
		body["ProductType"] = request.ProductType
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		body["SessionId"] = request.SessionId
	}

	if !tea.BoolValue(util.IsUnset(request.Uuid)) {
		body["Uuid"] = request.Uuid
	}

	if !tea.BoolValue(util.IsUnset(request.WyId)) {
		body["WyId"] = request.WyId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("StopApp"),
		Version:     tea.String("2021-09-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StopAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopApp(request *StopAppRequest) (_result *StopAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopAppResponse{}
	_body, _err := client.StopAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UnbindWithOptions(request *UnbindRequest, runtime *util.RuntimeOptions) (_result *UnbindResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.AppInstanceGroupId)) {
		body["AppInstanceGroupId"] = request.AppInstanceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.AppInstanceId)) {
		body["AppInstanceId"] = request.AppInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientId)) {
		body["ClientId"] = request.ClientId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientIp)) {
		body["ClientIp"] = request.ClientIp
	}

	if !tea.BoolValue(util.IsUnset(request.ClientOS)) {
		body["ClientOS"] = request.ClientOS
	}

	if !tea.BoolValue(util.IsUnset(request.ClientVersion)) {
		body["ClientVersion"] = request.ClientVersion
	}

	if !tea.BoolValue(util.IsUnset(request.EndUserId)) {
		body["EndUserId"] = request.EndUserId
	}

	if !tea.BoolValue(util.IsUnset(request.LoginRegionId)) {
		body["LoginRegionId"] = request.LoginRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.LoginToken)) {
		body["LoginToken"] = request.LoginToken
	}

	if !tea.BoolValue(util.IsUnset(request.ProductType)) {
		body["ProductType"] = request.ProductType
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		body["SessionId"] = request.SessionId
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["TenantId"] = request.TenantId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("Unbind"),
		Version:     tea.String("2021-09-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UnbindResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) Unbind(request *UnbindRequest) (_result *UnbindResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UnbindResponse{}
	_body, _err := client.UnbindWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
