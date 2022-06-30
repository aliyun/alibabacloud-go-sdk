// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CreateAdaptationRequest struct {
	AdaptTarget  *CreateAdaptationRequestAdaptTarget `json:"AdaptTarget,omitempty" xml:"AdaptTarget,omitempty" type:"Struct"`
	AppVersionId *string                             `json:"AppVersionId,omitempty" xml:"AppVersionId,omitempty"`
}

func (s CreateAdaptationRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAdaptationRequest) GoString() string {
	return s.String()
}

func (s *CreateAdaptationRequest) SetAdaptTarget(v *CreateAdaptationRequestAdaptTarget) *CreateAdaptationRequest {
	s.AdaptTarget = v
	return s
}

func (s *CreateAdaptationRequest) SetAppVersionId(v string) *CreateAdaptationRequest {
	s.AppVersionId = &v
	return s
}

type CreateAdaptationRequestAdaptTarget struct {
	BitRate      *int32  `json:"BitRate,omitempty" xml:"BitRate,omitempty"`
	FrameRate    *int32  `json:"FrameRate,omitempty" xml:"FrameRate,omitempty"`
	Resolution   *string `json:"Resolution,omitempty" xml:"Resolution,omitempty"`
	StartProgram *string `json:"StartProgram,omitempty" xml:"StartProgram,omitempty"`
}

func (s CreateAdaptationRequestAdaptTarget) String() string {
	return tea.Prettify(s)
}

func (s CreateAdaptationRequestAdaptTarget) GoString() string {
	return s.String()
}

func (s *CreateAdaptationRequestAdaptTarget) SetBitRate(v int32) *CreateAdaptationRequestAdaptTarget {
	s.BitRate = &v
	return s
}

func (s *CreateAdaptationRequestAdaptTarget) SetFrameRate(v int32) *CreateAdaptationRequestAdaptTarget {
	s.FrameRate = &v
	return s
}

func (s *CreateAdaptationRequestAdaptTarget) SetResolution(v string) *CreateAdaptationRequestAdaptTarget {
	s.Resolution = &v
	return s
}

func (s *CreateAdaptationRequestAdaptTarget) SetStartProgram(v string) *CreateAdaptationRequestAdaptTarget {
	s.StartProgram = &v
	return s
}

type CreateAdaptationShrinkRequest struct {
	AdaptTargetShrink *string `json:"AdaptTarget,omitempty" xml:"AdaptTarget,omitempty"`
	AppVersionId      *string `json:"AppVersionId,omitempty" xml:"AppVersionId,omitempty"`
}

func (s CreateAdaptationShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAdaptationShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateAdaptationShrinkRequest) SetAdaptTargetShrink(v string) *CreateAdaptationShrinkRequest {
	s.AdaptTargetShrink = &v
	return s
}

func (s *CreateAdaptationShrinkRequest) SetAppVersionId(v string) *CreateAdaptationShrinkRequest {
	s.AppVersionId = &v
	return s
}

type CreateAdaptationResponseBody struct {
	AdaptApplyId *int64  `json:"AdaptApplyId,omitempty" xml:"AdaptApplyId,omitempty"`
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpCode     *int32  `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Message      *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateAdaptationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAdaptationResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAdaptationResponseBody) SetAdaptApplyId(v int64) *CreateAdaptationResponseBody {
	s.AdaptApplyId = &v
	return s
}

func (s *CreateAdaptationResponseBody) SetCode(v string) *CreateAdaptationResponseBody {
	s.Code = &v
	return s
}

func (s *CreateAdaptationResponseBody) SetHttpCode(v int32) *CreateAdaptationResponseBody {
	s.HttpCode = &v
	return s
}

func (s *CreateAdaptationResponseBody) SetMessage(v string) *CreateAdaptationResponseBody {
	s.Message = &v
	return s
}

func (s *CreateAdaptationResponseBody) SetRequestId(v string) *CreateAdaptationResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAdaptationResponseBody) SetSuccess(v bool) *CreateAdaptationResponseBody {
	s.Success = &v
	return s
}

type CreateAdaptationResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateAdaptationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateAdaptationResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAdaptationResponse) GoString() string {
	return s.String()
}

func (s *CreateAdaptationResponse) SetHeaders(v map[string]*string) *CreateAdaptationResponse {
	s.Headers = v
	return s
}

func (s *CreateAdaptationResponse) SetStatusCode(v int32) *CreateAdaptationResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAdaptationResponse) SetBody(v *CreateAdaptationResponseBody) *CreateAdaptationResponse {
	s.Body = v
	return s
}

type CreateAppRequest struct {
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
}

func (s CreateAppRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAppRequest) GoString() string {
	return s.String()
}

func (s *CreateAppRequest) SetAppName(v string) *CreateAppRequest {
	s.AppName = &v
	return s
}

func (s *CreateAppRequest) SetAppType(v string) *CreateAppRequest {
	s.AppType = &v
	return s
}

type CreateAppResponseBody struct {
	AppId     *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpCode  *int32  `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAppResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAppResponseBody) SetAppId(v string) *CreateAppResponseBody {
	s.AppId = &v
	return s
}

func (s *CreateAppResponseBody) SetCode(v string) *CreateAppResponseBody {
	s.Code = &v
	return s
}

func (s *CreateAppResponseBody) SetHttpCode(v int32) *CreateAppResponseBody {
	s.HttpCode = &v
	return s
}

func (s *CreateAppResponseBody) SetMessage(v string) *CreateAppResponseBody {
	s.Message = &v
	return s
}

func (s *CreateAppResponseBody) SetRequestId(v string) *CreateAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAppResponseBody) SetSuccess(v bool) *CreateAppResponseBody {
	s.Success = &v
	return s
}

type CreateAppResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateAppResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateAppResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAppResponse) GoString() string {
	return s.String()
}

func (s *CreateAppResponse) SetHeaders(v map[string]*string) *CreateAppResponse {
	s.Headers = v
	return s
}

func (s *CreateAppResponse) SetStatusCode(v int32) *CreateAppResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAppResponse) SetBody(v *CreateAppResponseBody) *CreateAppResponse {
	s.Body = v
	return s
}

type CreateAppSessionRequest struct {
	// 应用ID
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 应用版本
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	// 客户端ip
	ClientIp *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	// 自定义会话id
	CustomSessionId *string `json:"CustomSessionId,omitempty" xml:"CustomSessionId,omitempty"`
	// 自定义用户id
	CustomUserId   *string                             `json:"CustomUserId,omitempty" xml:"CustomUserId,omitempty"`
	DatasetId      *string                             `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	EnablePostpaid *bool                               `json:"EnablePostpaid,omitempty" xml:"EnablePostpaid,omitempty"`
	ResultStore    *CreateAppSessionRequestResultStore `json:"ResultStore,omitempty" xml:"ResultStore,omitempty" type:"Struct"`
	// 启动参数
	StartParameters   []*CreateAppSessionRequestStartParameters   `json:"StartParameters,omitempty" xml:"StartParameters,omitempty" type:"Repeated"`
	StartParametersV2 []*CreateAppSessionRequestStartParametersV2 `json:"StartParametersV2,omitempty" xml:"StartParametersV2,omitempty" type:"Repeated"`
	// 系统信息：如端侧机型等信息
	SystemInfo []*CreateAppSessionRequestSystemInfo `json:"SystemInfo,omitempty" xml:"SystemInfo,omitempty" type:"Repeated"`
	Timeout    *int64                               `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s CreateAppSessionRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAppSessionRequest) GoString() string {
	return s.String()
}

func (s *CreateAppSessionRequest) SetAppId(v string) *CreateAppSessionRequest {
	s.AppId = &v
	return s
}

func (s *CreateAppSessionRequest) SetAppVersion(v string) *CreateAppSessionRequest {
	s.AppVersion = &v
	return s
}

func (s *CreateAppSessionRequest) SetClientIp(v string) *CreateAppSessionRequest {
	s.ClientIp = &v
	return s
}

func (s *CreateAppSessionRequest) SetCustomSessionId(v string) *CreateAppSessionRequest {
	s.CustomSessionId = &v
	return s
}

func (s *CreateAppSessionRequest) SetCustomUserId(v string) *CreateAppSessionRequest {
	s.CustomUserId = &v
	return s
}

func (s *CreateAppSessionRequest) SetDatasetId(v string) *CreateAppSessionRequest {
	s.DatasetId = &v
	return s
}

func (s *CreateAppSessionRequest) SetEnablePostpaid(v bool) *CreateAppSessionRequest {
	s.EnablePostpaid = &v
	return s
}

func (s *CreateAppSessionRequest) SetResultStore(v *CreateAppSessionRequestResultStore) *CreateAppSessionRequest {
	s.ResultStore = v
	return s
}

func (s *CreateAppSessionRequest) SetStartParameters(v []*CreateAppSessionRequestStartParameters) *CreateAppSessionRequest {
	s.StartParameters = v
	return s
}

func (s *CreateAppSessionRequest) SetStartParametersV2(v []*CreateAppSessionRequestStartParametersV2) *CreateAppSessionRequest {
	s.StartParametersV2 = v
	return s
}

func (s *CreateAppSessionRequest) SetSystemInfo(v []*CreateAppSessionRequestSystemInfo) *CreateAppSessionRequest {
	s.SystemInfo = v
	return s
}

func (s *CreateAppSessionRequest) SetTimeout(v int64) *CreateAppSessionRequest {
	s.Timeout = &v
	return s
}

type CreateAppSessionRequestResultStore struct {
	Need      *bool                                          `json:"Need,omitempty" xml:"Need,omitempty"`
	StoreInfo []*CreateAppSessionRequestResultStoreStoreInfo `json:"StoreInfo,omitempty" xml:"StoreInfo,omitempty" type:"Repeated"`
	Type      *string                                        `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateAppSessionRequestResultStore) String() string {
	return tea.Prettify(s)
}

func (s CreateAppSessionRequestResultStore) GoString() string {
	return s.String()
}

func (s *CreateAppSessionRequestResultStore) SetNeed(v bool) *CreateAppSessionRequestResultStore {
	s.Need = &v
	return s
}

func (s *CreateAppSessionRequestResultStore) SetStoreInfo(v []*CreateAppSessionRequestResultStoreStoreInfo) *CreateAppSessionRequestResultStore {
	s.StoreInfo = v
	return s
}

func (s *CreateAppSessionRequestResultStore) SetType(v string) *CreateAppSessionRequestResultStore {
	s.Type = &v
	return s
}

type CreateAppSessionRequestResultStoreStoreInfo struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateAppSessionRequestResultStoreStoreInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateAppSessionRequestResultStoreStoreInfo) GoString() string {
	return s.String()
}

func (s *CreateAppSessionRequestResultStoreStoreInfo) SetKey(v string) *CreateAppSessionRequestResultStoreStoreInfo {
	s.Key = &v
	return s
}

func (s *CreateAppSessionRequestResultStoreStoreInfo) SetValue(v string) *CreateAppSessionRequestResultStoreStoreInfo {
	s.Value = &v
	return s
}

type CreateAppSessionRequestStartParameters struct {
	// key
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateAppSessionRequestStartParameters) String() string {
	return tea.Prettify(s)
}

func (s CreateAppSessionRequestStartParameters) GoString() string {
	return s.String()
}

func (s *CreateAppSessionRequestStartParameters) SetKey(v string) *CreateAppSessionRequestStartParameters {
	s.Key = &v
	return s
}

func (s *CreateAppSessionRequestStartParameters) SetValue(v string) *CreateAppSessionRequestStartParameters {
	s.Value = &v
	return s
}

type CreateAppSessionRequestStartParametersV2 struct {
	Key   *string     `json:"Key,omitempty" xml:"Key,omitempty"`
	Value interface{} `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateAppSessionRequestStartParametersV2) String() string {
	return tea.Prettify(s)
}

func (s CreateAppSessionRequestStartParametersV2) GoString() string {
	return s.String()
}

func (s *CreateAppSessionRequestStartParametersV2) SetKey(v string) *CreateAppSessionRequestStartParametersV2 {
	s.Key = &v
	return s
}

func (s *CreateAppSessionRequestStartParametersV2) SetValue(v interface{}) *CreateAppSessionRequestStartParametersV2 {
	s.Value = v
	return s
}

type CreateAppSessionRequestSystemInfo struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateAppSessionRequestSystemInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateAppSessionRequestSystemInfo) GoString() string {
	return s.String()
}

func (s *CreateAppSessionRequestSystemInfo) SetKey(v string) *CreateAppSessionRequestSystemInfo {
	s.Key = &v
	return s
}

func (s *CreateAppSessionRequestSystemInfo) SetValue(v string) *CreateAppSessionRequestSystemInfo {
	s.Value = &v
	return s
}

type CreateAppSessionShrinkRequest struct {
	// 应用ID
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 应用版本
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	// 客户端ip
	ClientIp *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	// 自定义会话id
	CustomSessionId *string `json:"CustomSessionId,omitempty" xml:"CustomSessionId,omitempty"`
	// 自定义用户id
	CustomUserId      *string `json:"CustomUserId,omitempty" xml:"CustomUserId,omitempty"`
	DatasetId         *string `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	EnablePostpaid    *bool   `json:"EnablePostpaid,omitempty" xml:"EnablePostpaid,omitempty"`
	ResultStoreShrink *string `json:"ResultStore,omitempty" xml:"ResultStore,omitempty"`
	// 启动参数
	StartParameters         []*CreateAppSessionShrinkRequestStartParameters `json:"StartParameters,omitempty" xml:"StartParameters,omitempty" type:"Repeated"`
	StartParametersV2Shrink *string                                         `json:"StartParametersV2,omitempty" xml:"StartParametersV2,omitempty"`
	// 系统信息：如端侧机型等信息
	SystemInfo []*CreateAppSessionShrinkRequestSystemInfo `json:"SystemInfo,omitempty" xml:"SystemInfo,omitempty" type:"Repeated"`
	Timeout    *int64                                     `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s CreateAppSessionShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAppSessionShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateAppSessionShrinkRequest) SetAppId(v string) *CreateAppSessionShrinkRequest {
	s.AppId = &v
	return s
}

func (s *CreateAppSessionShrinkRequest) SetAppVersion(v string) *CreateAppSessionShrinkRequest {
	s.AppVersion = &v
	return s
}

func (s *CreateAppSessionShrinkRequest) SetClientIp(v string) *CreateAppSessionShrinkRequest {
	s.ClientIp = &v
	return s
}

func (s *CreateAppSessionShrinkRequest) SetCustomSessionId(v string) *CreateAppSessionShrinkRequest {
	s.CustomSessionId = &v
	return s
}

func (s *CreateAppSessionShrinkRequest) SetCustomUserId(v string) *CreateAppSessionShrinkRequest {
	s.CustomUserId = &v
	return s
}

func (s *CreateAppSessionShrinkRequest) SetDatasetId(v string) *CreateAppSessionShrinkRequest {
	s.DatasetId = &v
	return s
}

func (s *CreateAppSessionShrinkRequest) SetEnablePostpaid(v bool) *CreateAppSessionShrinkRequest {
	s.EnablePostpaid = &v
	return s
}

func (s *CreateAppSessionShrinkRequest) SetResultStoreShrink(v string) *CreateAppSessionShrinkRequest {
	s.ResultStoreShrink = &v
	return s
}

func (s *CreateAppSessionShrinkRequest) SetStartParameters(v []*CreateAppSessionShrinkRequestStartParameters) *CreateAppSessionShrinkRequest {
	s.StartParameters = v
	return s
}

func (s *CreateAppSessionShrinkRequest) SetStartParametersV2Shrink(v string) *CreateAppSessionShrinkRequest {
	s.StartParametersV2Shrink = &v
	return s
}

func (s *CreateAppSessionShrinkRequest) SetSystemInfo(v []*CreateAppSessionShrinkRequestSystemInfo) *CreateAppSessionShrinkRequest {
	s.SystemInfo = v
	return s
}

func (s *CreateAppSessionShrinkRequest) SetTimeout(v int64) *CreateAppSessionShrinkRequest {
	s.Timeout = &v
	return s
}

type CreateAppSessionShrinkRequestStartParameters struct {
	// key
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateAppSessionShrinkRequestStartParameters) String() string {
	return tea.Prettify(s)
}

func (s CreateAppSessionShrinkRequestStartParameters) GoString() string {
	return s.String()
}

func (s *CreateAppSessionShrinkRequestStartParameters) SetKey(v string) *CreateAppSessionShrinkRequestStartParameters {
	s.Key = &v
	return s
}

func (s *CreateAppSessionShrinkRequestStartParameters) SetValue(v string) *CreateAppSessionShrinkRequestStartParameters {
	s.Value = &v
	return s
}

type CreateAppSessionShrinkRequestSystemInfo struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateAppSessionShrinkRequestSystemInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateAppSessionShrinkRequestSystemInfo) GoString() string {
	return s.String()
}

func (s *CreateAppSessionShrinkRequestSystemInfo) SetKey(v string) *CreateAppSessionShrinkRequestSystemInfo {
	s.Key = &v
	return s
}

func (s *CreateAppSessionShrinkRequestSystemInfo) SetValue(v string) *CreateAppSessionShrinkRequestSystemInfo {
	s.Value = &v
	return s
}

type CreateAppSessionResponseBody struct {
	// 应用id
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 应用版本
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	// 自定义会话id
	CustomSessionId *string `json:"CustomSessionId,omitempty" xml:"CustomSessionId,omitempty"`
	// 平台会话id
	PlatformSessionId *string `json:"PlatformSessionId,omitempty" xml:"PlatformSessionId,omitempty"`
	// 请求id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAppSessionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAppSessionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAppSessionResponseBody) SetAppId(v string) *CreateAppSessionResponseBody {
	s.AppId = &v
	return s
}

func (s *CreateAppSessionResponseBody) SetAppVersion(v string) *CreateAppSessionResponseBody {
	s.AppVersion = &v
	return s
}

func (s *CreateAppSessionResponseBody) SetCustomSessionId(v string) *CreateAppSessionResponseBody {
	s.CustomSessionId = &v
	return s
}

func (s *CreateAppSessionResponseBody) SetPlatformSessionId(v string) *CreateAppSessionResponseBody {
	s.PlatformSessionId = &v
	return s
}

func (s *CreateAppSessionResponseBody) SetRequestId(v string) *CreateAppSessionResponseBody {
	s.RequestId = &v
	return s
}

type CreateAppSessionResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateAppSessionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateAppSessionResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAppSessionResponse) GoString() string {
	return s.String()
}

func (s *CreateAppSessionResponse) SetHeaders(v map[string]*string) *CreateAppSessionResponse {
	s.Headers = v
	return s
}

func (s *CreateAppSessionResponse) SetStatusCode(v int32) *CreateAppSessionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAppSessionResponse) SetBody(v *CreateAppSessionResponseBody) *CreateAppSessionResponse {
	s.Body = v
	return s
}

type CreateAppSessionSyncRequest struct {
	// 应用ID
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 应用版本
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	// 客户端ip
	ClientIp *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	// 自定义会话id
	CustomSessionId *string `json:"CustomSessionId,omitempty" xml:"CustomSessionId,omitempty"`
	// 自定义用户id
	CustomUserId *string                                 `json:"CustomUserId,omitempty" xml:"CustomUserId,omitempty"`
	DistrictId   *string                                 `json:"DistrictId,omitempty" xml:"DistrictId,omitempty"`
	ProjectId    *string                                 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	SpeedInfo    []*CreateAppSessionSyncRequestSpeedInfo `json:"SpeedInfo,omitempty" xml:"SpeedInfo,omitempty" type:"Repeated"`
	// 启动参数
	StartParameters []*CreateAppSessionSyncRequestStartParameters `json:"StartParameters,omitempty" xml:"StartParameters,omitempty" type:"Repeated"`
	// 系统信息：如端侧机型等信息
	SystemInfo []*CreateAppSessionSyncRequestSystemInfo `json:"SystemInfo,omitempty" xml:"SystemInfo,omitempty" type:"Repeated"`
	Tags       []*CreateAppSessionSyncRequestTags       `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
}

func (s CreateAppSessionSyncRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAppSessionSyncRequest) GoString() string {
	return s.String()
}

func (s *CreateAppSessionSyncRequest) SetAppId(v string) *CreateAppSessionSyncRequest {
	s.AppId = &v
	return s
}

func (s *CreateAppSessionSyncRequest) SetAppVersion(v string) *CreateAppSessionSyncRequest {
	s.AppVersion = &v
	return s
}

func (s *CreateAppSessionSyncRequest) SetClientIp(v string) *CreateAppSessionSyncRequest {
	s.ClientIp = &v
	return s
}

func (s *CreateAppSessionSyncRequest) SetCustomSessionId(v string) *CreateAppSessionSyncRequest {
	s.CustomSessionId = &v
	return s
}

func (s *CreateAppSessionSyncRequest) SetCustomUserId(v string) *CreateAppSessionSyncRequest {
	s.CustomUserId = &v
	return s
}

func (s *CreateAppSessionSyncRequest) SetDistrictId(v string) *CreateAppSessionSyncRequest {
	s.DistrictId = &v
	return s
}

func (s *CreateAppSessionSyncRequest) SetProjectId(v string) *CreateAppSessionSyncRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateAppSessionSyncRequest) SetSpeedInfo(v []*CreateAppSessionSyncRequestSpeedInfo) *CreateAppSessionSyncRequest {
	s.SpeedInfo = v
	return s
}

func (s *CreateAppSessionSyncRequest) SetStartParameters(v []*CreateAppSessionSyncRequestStartParameters) *CreateAppSessionSyncRequest {
	s.StartParameters = v
	return s
}

func (s *CreateAppSessionSyncRequest) SetSystemInfo(v []*CreateAppSessionSyncRequestSystemInfo) *CreateAppSessionSyncRequest {
	s.SystemInfo = v
	return s
}

func (s *CreateAppSessionSyncRequest) SetTags(v []*CreateAppSessionSyncRequestTags) *CreateAppSessionSyncRequest {
	s.Tags = v
	return s
}

type CreateAppSessionSyncRequestSpeedInfo struct {
	EndpointId *int32 `json:"endpointId,omitempty" xml:"endpointId,omitempty"`
	Rtt        *int32 `json:"rtt,omitempty" xml:"rtt,omitempty"`
}

func (s CreateAppSessionSyncRequestSpeedInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateAppSessionSyncRequestSpeedInfo) GoString() string {
	return s.String()
}

func (s *CreateAppSessionSyncRequestSpeedInfo) SetEndpointId(v int32) *CreateAppSessionSyncRequestSpeedInfo {
	s.EndpointId = &v
	return s
}

func (s *CreateAppSessionSyncRequestSpeedInfo) SetRtt(v int32) *CreateAppSessionSyncRequestSpeedInfo {
	s.Rtt = &v
	return s
}

type CreateAppSessionSyncRequestStartParameters struct {
	// key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateAppSessionSyncRequestStartParameters) String() string {
	return tea.Prettify(s)
}

func (s CreateAppSessionSyncRequestStartParameters) GoString() string {
	return s.String()
}

func (s *CreateAppSessionSyncRequestStartParameters) SetKey(v string) *CreateAppSessionSyncRequestStartParameters {
	s.Key = &v
	return s
}

func (s *CreateAppSessionSyncRequestStartParameters) SetValue(v string) *CreateAppSessionSyncRequestStartParameters {
	s.Value = &v
	return s
}

type CreateAppSessionSyncRequestSystemInfo struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateAppSessionSyncRequestSystemInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateAppSessionSyncRequestSystemInfo) GoString() string {
	return s.String()
}

func (s *CreateAppSessionSyncRequestSystemInfo) SetKey(v string) *CreateAppSessionSyncRequestSystemInfo {
	s.Key = &v
	return s
}

func (s *CreateAppSessionSyncRequestSystemInfo) SetValue(v string) *CreateAppSessionSyncRequestSystemInfo {
	s.Value = &v
	return s
}

type CreateAppSessionSyncRequestTags struct {
	Key   *string `json:"key,omitempty" xml:"key,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s CreateAppSessionSyncRequestTags) String() string {
	return tea.Prettify(s)
}

func (s CreateAppSessionSyncRequestTags) GoString() string {
	return s.String()
}

func (s *CreateAppSessionSyncRequestTags) SetKey(v string) *CreateAppSessionSyncRequestTags {
	s.Key = &v
	return s
}

func (s *CreateAppSessionSyncRequestTags) SetValue(v string) *CreateAppSessionSyncRequestTags {
	s.Value = &v
	return s
}

type CreateAppSessionSyncResponseBody struct {
	// 应用id
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 应用版本
	AppVersion *string                                  `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	BizInfo    *CreateAppSessionSyncResponseBodyBizInfo `json:"BizInfo,omitempty" xml:"BizInfo,omitempty" type:"Struct"`
	// 自定义会话id
	CustomSessionId *string `json:"CustomSessionId,omitempty" xml:"CustomSessionId,omitempty"`
	// 平台会话id
	PlatformSessionId *string `json:"PlatformSessionId,omitempty" xml:"PlatformSessionId,omitempty"`
	// 请求id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAppSessionSyncResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAppSessionSyncResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAppSessionSyncResponseBody) SetAppId(v string) *CreateAppSessionSyncResponseBody {
	s.AppId = &v
	return s
}

func (s *CreateAppSessionSyncResponseBody) SetAppVersion(v string) *CreateAppSessionSyncResponseBody {
	s.AppVersion = &v
	return s
}

func (s *CreateAppSessionSyncResponseBody) SetBizInfo(v *CreateAppSessionSyncResponseBodyBizInfo) *CreateAppSessionSyncResponseBody {
	s.BizInfo = v
	return s
}

func (s *CreateAppSessionSyncResponseBody) SetCustomSessionId(v string) *CreateAppSessionSyncResponseBody {
	s.CustomSessionId = &v
	return s
}

func (s *CreateAppSessionSyncResponseBody) SetPlatformSessionId(v string) *CreateAppSessionSyncResponseBody {
	s.PlatformSessionId = &v
	return s
}

func (s *CreateAppSessionSyncResponseBody) SetRequestId(v string) *CreateAppSessionSyncResponseBody {
	s.RequestId = &v
	return s
}

type CreateAppSessionSyncResponseBodyBizInfo struct {
	Biz       []*CreateAppSessionSyncResponseBodyBizInfoBiz       `json:"Biz,omitempty" xml:"Biz,omitempty" type:"Repeated"`
	Endpoints []*CreateAppSessionSyncResponseBodyBizInfoEndpoints `json:"Endpoints,omitempty" xml:"Endpoints,omitempty" type:"Repeated"`
}

func (s CreateAppSessionSyncResponseBodyBizInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateAppSessionSyncResponseBodyBizInfo) GoString() string {
	return s.String()
}

func (s *CreateAppSessionSyncResponseBodyBizInfo) SetBiz(v []*CreateAppSessionSyncResponseBodyBizInfoBiz) *CreateAppSessionSyncResponseBodyBizInfo {
	s.Biz = v
	return s
}

func (s *CreateAppSessionSyncResponseBodyBizInfo) SetEndpoints(v []*CreateAppSessionSyncResponseBodyBizInfoEndpoints) *CreateAppSessionSyncResponseBodyBizInfo {
	s.Endpoints = v
	return s
}

type CreateAppSessionSyncResponseBodyBizInfoBiz struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateAppSessionSyncResponseBodyBizInfoBiz) String() string {
	return tea.Prettify(s)
}

func (s CreateAppSessionSyncResponseBodyBizInfoBiz) GoString() string {
	return s.String()
}

func (s *CreateAppSessionSyncResponseBodyBizInfoBiz) SetKey(v string) *CreateAppSessionSyncResponseBodyBizInfoBiz {
	s.Key = &v
	return s
}

func (s *CreateAppSessionSyncResponseBodyBizInfoBiz) SetValue(v string) *CreateAppSessionSyncResponseBodyBizInfoBiz {
	s.Value = &v
	return s
}

type CreateAppSessionSyncResponseBodyBizInfoEndpoints struct {
	AccessHost *string `json:"AccessHost,omitempty" xml:"AccessHost,omitempty"`
	AccessPort *string `json:"AccessPort,omitempty" xml:"AccessPort,omitempty"`
	DistrictId *string `json:"DistrictId,omitempty" xml:"DistrictId,omitempty"`
	Isp        *string `json:"Isp,omitempty" xml:"Isp,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Type       *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateAppSessionSyncResponseBodyBizInfoEndpoints) String() string {
	return tea.Prettify(s)
}

func (s CreateAppSessionSyncResponseBodyBizInfoEndpoints) GoString() string {
	return s.String()
}

func (s *CreateAppSessionSyncResponseBodyBizInfoEndpoints) SetAccessHost(v string) *CreateAppSessionSyncResponseBodyBizInfoEndpoints {
	s.AccessHost = &v
	return s
}

func (s *CreateAppSessionSyncResponseBodyBizInfoEndpoints) SetAccessPort(v string) *CreateAppSessionSyncResponseBodyBizInfoEndpoints {
	s.AccessPort = &v
	return s
}

func (s *CreateAppSessionSyncResponseBodyBizInfoEndpoints) SetDistrictId(v string) *CreateAppSessionSyncResponseBodyBizInfoEndpoints {
	s.DistrictId = &v
	return s
}

func (s *CreateAppSessionSyncResponseBodyBizInfoEndpoints) SetIsp(v string) *CreateAppSessionSyncResponseBodyBizInfoEndpoints {
	s.Isp = &v
	return s
}

func (s *CreateAppSessionSyncResponseBodyBizInfoEndpoints) SetName(v string) *CreateAppSessionSyncResponseBodyBizInfoEndpoints {
	s.Name = &v
	return s
}

func (s *CreateAppSessionSyncResponseBodyBizInfoEndpoints) SetType(v string) *CreateAppSessionSyncResponseBodyBizInfoEndpoints {
	s.Type = &v
	return s
}

type CreateAppSessionSyncResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateAppSessionSyncResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateAppSessionSyncResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAppSessionSyncResponse) GoString() string {
	return s.String()
}

func (s *CreateAppSessionSyncResponse) SetHeaders(v map[string]*string) *CreateAppSessionSyncResponse {
	s.Headers = v
	return s
}

func (s *CreateAppSessionSyncResponse) SetStatusCode(v int32) *CreateAppSessionSyncResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAppSessionSyncResponse) SetBody(v *CreateAppSessionSyncResponseBody) *CreateAppSessionSyncResponse {
	s.Body = v
	return s
}

type CreateAppVersionRequest struct {
	AppId          *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppVersionName *string `json:"AppVersionName,omitempty" xml:"AppVersionName,omitempty"`
}

func (s CreateAppVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAppVersionRequest) GoString() string {
	return s.String()
}

func (s *CreateAppVersionRequest) SetAppId(v string) *CreateAppVersionRequest {
	s.AppId = &v
	return s
}

func (s *CreateAppVersionRequest) SetAppVersionName(v string) *CreateAppVersionRequest {
	s.AppVersionName = &v
	return s
}

type CreateAppVersionResponseBody struct {
	AppVersionId *string `json:"AppVersionId,omitempty" xml:"AppVersionId,omitempty"`
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpCode     *int32  `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Message      *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateAppVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAppVersionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAppVersionResponseBody) SetAppVersionId(v string) *CreateAppVersionResponseBody {
	s.AppVersionId = &v
	return s
}

func (s *CreateAppVersionResponseBody) SetCode(v string) *CreateAppVersionResponseBody {
	s.Code = &v
	return s
}

func (s *CreateAppVersionResponseBody) SetHttpCode(v int32) *CreateAppVersionResponseBody {
	s.HttpCode = &v
	return s
}

func (s *CreateAppVersionResponseBody) SetMessage(v string) *CreateAppVersionResponseBody {
	s.Message = &v
	return s
}

func (s *CreateAppVersionResponseBody) SetRequestId(v string) *CreateAppVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAppVersionResponseBody) SetSuccess(v bool) *CreateAppVersionResponseBody {
	s.Success = &v
	return s
}

type CreateAppVersionResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateAppVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateAppVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAppVersionResponse) GoString() string {
	return s.String()
}

func (s *CreateAppVersionResponse) SetHeaders(v map[string]*string) *CreateAppVersionResponse {
	s.Headers = v
	return s
}

func (s *CreateAppVersionResponse) SetStatusCode(v int32) *CreateAppVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAppVersionResponse) SetBody(v *CreateAppVersionResponseBody) *CreateAppVersionResponse {
	s.Body = v
	return s
}

type CreateCapacityReservationRequest struct {
	// 自定义会话id
	AppId                   *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppVersion              *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	ClientToken             *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DistrictId              *string `json:"DistrictId,omitempty" xml:"DistrictId,omitempty"`
	ExpectResourceReadyTime *string `json:"ExpectResourceReadyTime,omitempty" xml:"ExpectResourceReadyTime,omitempty"`
	ExpectSessionCapacity   *int32  `json:"ExpectSessionCapacity,omitempty" xml:"ExpectSessionCapacity,omitempty"`
	// 平台会话id
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s CreateCapacityReservationRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCapacityReservationRequest) GoString() string {
	return s.String()
}

func (s *CreateCapacityReservationRequest) SetAppId(v string) *CreateCapacityReservationRequest {
	s.AppId = &v
	return s
}

func (s *CreateCapacityReservationRequest) SetAppVersion(v string) *CreateCapacityReservationRequest {
	s.AppVersion = &v
	return s
}

func (s *CreateCapacityReservationRequest) SetClientToken(v string) *CreateCapacityReservationRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateCapacityReservationRequest) SetDistrictId(v string) *CreateCapacityReservationRequest {
	s.DistrictId = &v
	return s
}

func (s *CreateCapacityReservationRequest) SetExpectResourceReadyTime(v string) *CreateCapacityReservationRequest {
	s.ExpectResourceReadyTime = &v
	return s
}

func (s *CreateCapacityReservationRequest) SetExpectSessionCapacity(v int32) *CreateCapacityReservationRequest {
	s.ExpectSessionCapacity = &v
	return s
}

func (s *CreateCapacityReservationRequest) SetProjectId(v string) *CreateCapacityReservationRequest {
	s.ProjectId = &v
	return s
}

type CreateCapacityReservationResponseBody struct {
	CurrMaxAllocatableSessionCapacity *int32 `json:"CurrMaxAllocatableSessionCapacity,omitempty" xml:"CurrMaxAllocatableSessionCapacity,omitempty"`
	// 请求id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 自定义会话id
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateCapacityReservationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateCapacityReservationResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCapacityReservationResponseBody) SetCurrMaxAllocatableSessionCapacity(v int32) *CreateCapacityReservationResponseBody {
	s.CurrMaxAllocatableSessionCapacity = &v
	return s
}

func (s *CreateCapacityReservationResponseBody) SetRequestId(v string) *CreateCapacityReservationResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCapacityReservationResponseBody) SetTaskId(v string) *CreateCapacityReservationResponseBody {
	s.TaskId = &v
	return s
}

type CreateCapacityReservationResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateCapacityReservationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateCapacityReservationResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCapacityReservationResponse) GoString() string {
	return s.String()
}

func (s *CreateCapacityReservationResponse) SetHeaders(v map[string]*string) *CreateCapacityReservationResponse {
	s.Headers = v
	return s
}

func (s *CreateCapacityReservationResponse) SetStatusCode(v int32) *CreateCapacityReservationResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCapacityReservationResponse) SetBody(v *CreateCapacityReservationResponseBody) *CreateCapacityReservationResponse {
	s.Body = v
	return s
}

type CreateDatasetDeployTaskRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	CustomParam *string `json:"CustomParam,omitempty" xml:"CustomParam,omitempty"`
	NeedUnzip   *bool   `json:"NeedUnzip,omitempty" xml:"NeedUnzip,omitempty"`
	OssBucket   *string `json:"OssBucket,omitempty" xml:"OssBucket,omitempty"`
	OssFilePath *string `json:"OssFilePath,omitempty" xml:"OssFilePath,omitempty"`
	OssRegionId *string `json:"OssRegionId,omitempty" xml:"OssRegionId,omitempty"`
	SourceType  *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
}

func (s CreateDatasetDeployTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDatasetDeployTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateDatasetDeployTaskRequest) SetClientToken(v string) *CreateDatasetDeployTaskRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateDatasetDeployTaskRequest) SetCustomParam(v string) *CreateDatasetDeployTaskRequest {
	s.CustomParam = &v
	return s
}

func (s *CreateDatasetDeployTaskRequest) SetNeedUnzip(v bool) *CreateDatasetDeployTaskRequest {
	s.NeedUnzip = &v
	return s
}

func (s *CreateDatasetDeployTaskRequest) SetOssBucket(v string) *CreateDatasetDeployTaskRequest {
	s.OssBucket = &v
	return s
}

func (s *CreateDatasetDeployTaskRequest) SetOssFilePath(v string) *CreateDatasetDeployTaskRequest {
	s.OssFilePath = &v
	return s
}

func (s *CreateDatasetDeployTaskRequest) SetOssRegionId(v string) *CreateDatasetDeployTaskRequest {
	s.OssRegionId = &v
	return s
}

func (s *CreateDatasetDeployTaskRequest) SetSourceType(v string) *CreateDatasetDeployTaskRequest {
	s.SourceType = &v
	return s
}

type CreateDatasetDeployTaskResponseBody struct {
	// 请求id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 应用版本
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateDatasetDeployTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDatasetDeployTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDatasetDeployTaskResponseBody) SetRequestId(v string) *CreateDatasetDeployTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDatasetDeployTaskResponseBody) SetTaskId(v string) *CreateDatasetDeployTaskResponseBody {
	s.TaskId = &v
	return s
}

type CreateDatasetDeployTaskResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateDatasetDeployTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDatasetDeployTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDatasetDeployTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateDatasetDeployTaskResponse) SetHeaders(v map[string]*string) *CreateDatasetDeployTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateDatasetDeployTaskResponse) SetStatusCode(v int32) *CreateDatasetDeployTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDatasetDeployTaskResponse) SetBody(v *CreateDatasetDeployTaskResponseBody) *CreateDatasetDeployTaskResponse {
	s.Body = v
	return s
}

type CreateProjectRequest struct {
	BoundAppIdList []*string `json:"BoundAppIdList,omitempty" xml:"BoundAppIdList,omitempty" type:"Repeated"`
	// 请求操作人Id
	OperatorId *string `json:"OperatorId,omitempty" xml:"OperatorId,omitempty"`
	// 请求操作人类型
	OperatorType *string `json:"OperatorType,omitempty" xml:"OperatorType,omitempty"`
	ProjectMemo  *string `json:"ProjectMemo,omitempty" xml:"ProjectMemo,omitempty"`
	// project name
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// key : districtId
	ProjectQuotaLimit *CreateProjectRequestProjectQuotaLimit `json:"ProjectQuotaLimit,omitempty" xml:"ProjectQuotaLimit,omitempty" type:"Struct"`
}

func (s CreateProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectRequest) GoString() string {
	return s.String()
}

func (s *CreateProjectRequest) SetBoundAppIdList(v []*string) *CreateProjectRequest {
	s.BoundAppIdList = v
	return s
}

func (s *CreateProjectRequest) SetOperatorId(v string) *CreateProjectRequest {
	s.OperatorId = &v
	return s
}

func (s *CreateProjectRequest) SetOperatorType(v string) *CreateProjectRequest {
	s.OperatorType = &v
	return s
}

func (s *CreateProjectRequest) SetProjectMemo(v string) *CreateProjectRequest {
	s.ProjectMemo = &v
	return s
}

func (s *CreateProjectRequest) SetProjectName(v string) *CreateProjectRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateProjectRequest) SetProjectQuotaLimit(v *CreateProjectRequestProjectQuotaLimit) *CreateProjectRequest {
	s.ProjectQuotaLimit = v
	return s
}

type CreateProjectRequestProjectQuotaLimit struct {
	// key - districtId
	DistrictLimitMap map[string]*ProjectQuotaLimitDistrictLimitMapValue `json:"DistrictLimitMap,omitempty" xml:"DistrictLimitMap,omitempty"`
	// 限制类型 ：目前默认 - ReserveContainer
	LimitType *string `json:"LimitType,omitempty" xml:"LimitType,omitempty"`
}

func (s CreateProjectRequestProjectQuotaLimit) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectRequestProjectQuotaLimit) GoString() string {
	return s.String()
}

func (s *CreateProjectRequestProjectQuotaLimit) SetDistrictLimitMap(v map[string]*ProjectQuotaLimitDistrictLimitMapValue) *CreateProjectRequestProjectQuotaLimit {
	s.DistrictLimitMap = v
	return s
}

func (s *CreateProjectRequestProjectQuotaLimit) SetLimitType(v string) *CreateProjectRequestProjectQuotaLimit {
	s.LimitType = &v
	return s
}

type CreateProjectShrinkRequest struct {
	BoundAppIdListShrink *string `json:"BoundAppIdList,omitempty" xml:"BoundAppIdList,omitempty"`
	// 请求操作人Id
	OperatorId *string `json:"OperatorId,omitempty" xml:"OperatorId,omitempty"`
	// 请求操作人类型
	OperatorType *string `json:"OperatorType,omitempty" xml:"OperatorType,omitempty"`
	ProjectMemo  *string `json:"ProjectMemo,omitempty" xml:"ProjectMemo,omitempty"`
	// project name
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// key : districtId
	ProjectQuotaLimitShrink *string `json:"ProjectQuotaLimit,omitempty" xml:"ProjectQuotaLimit,omitempty"`
}

func (s CreateProjectShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateProjectShrinkRequest) SetBoundAppIdListShrink(v string) *CreateProjectShrinkRequest {
	s.BoundAppIdListShrink = &v
	return s
}

func (s *CreateProjectShrinkRequest) SetOperatorId(v string) *CreateProjectShrinkRequest {
	s.OperatorId = &v
	return s
}

func (s *CreateProjectShrinkRequest) SetOperatorType(v string) *CreateProjectShrinkRequest {
	s.OperatorType = &v
	return s
}

func (s *CreateProjectShrinkRequest) SetProjectMemo(v string) *CreateProjectShrinkRequest {
	s.ProjectMemo = &v
	return s
}

func (s *CreateProjectShrinkRequest) SetProjectName(v string) *CreateProjectShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateProjectShrinkRequest) SetProjectQuotaLimitShrink(v string) *CreateProjectShrinkRequest {
	s.ProjectQuotaLimitShrink = &v
	return s
}

type CreateProjectResponseBody struct {
	// 业务处理结果Code
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 业务对象
	Data *CreateProjectResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// 业务处理消息摘要
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 操作请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 业务处理是否成功
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectResponseBody) GoString() string {
	return s.String()
}

func (s *CreateProjectResponseBody) SetCode(v string) *CreateProjectResponseBody {
	s.Code = &v
	return s
}

func (s *CreateProjectResponseBody) SetData(v *CreateProjectResponseBodyData) *CreateProjectResponseBody {
	s.Data = v
	return s
}

func (s *CreateProjectResponseBody) SetMessage(v string) *CreateProjectResponseBody {
	s.Message = &v
	return s
}

func (s *CreateProjectResponseBody) SetRequestId(v string) *CreateProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateProjectResponseBody) SetSuccess(v bool) *CreateProjectResponseBody {
	s.Success = &v
	return s
}

type CreateProjectResponseBodyData struct {
	// 业务处理结果Code
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 业务对象
	Data map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// 业务处理消息摘要
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// 操作请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 业务处理是否成功
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateProjectResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateProjectResponseBodyData) SetCode(v string) *CreateProjectResponseBodyData {
	s.Code = &v
	return s
}

func (s *CreateProjectResponseBodyData) SetData(v map[string]interface{}) *CreateProjectResponseBodyData {
	s.Data = v
	return s
}

func (s *CreateProjectResponseBodyData) SetMessage(v string) *CreateProjectResponseBodyData {
	s.Message = &v
	return s
}

func (s *CreateProjectResponseBodyData) SetProjectId(v string) *CreateProjectResponseBodyData {
	s.ProjectId = &v
	return s
}

func (s *CreateProjectResponseBodyData) SetRequestId(v string) *CreateProjectResponseBodyData {
	s.RequestId = &v
	return s
}

func (s *CreateProjectResponseBodyData) SetSuccess(v bool) *CreateProjectResponseBodyData {
	s.Success = &v
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

type DeleteAppRequest struct {
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s DeleteAppRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAppRequest) GoString() string {
	return s.String()
}

func (s *DeleteAppRequest) SetAppId(v string) *DeleteAppRequest {
	s.AppId = &v
	return s
}

type DeleteAppResponseBody struct {
	AppId     *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpCode  *int32  `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAppResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAppResponseBody) SetAppId(v string) *DeleteAppResponseBody {
	s.AppId = &v
	return s
}

func (s *DeleteAppResponseBody) SetCode(v string) *DeleteAppResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteAppResponseBody) SetHttpCode(v int32) *DeleteAppResponseBody {
	s.HttpCode = &v
	return s
}

func (s *DeleteAppResponseBody) SetMessage(v string) *DeleteAppResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteAppResponseBody) SetRequestId(v string) *DeleteAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAppResponseBody) SetSuccess(v bool) *DeleteAppResponseBody {
	s.Success = &v
	return s
}

type DeleteAppResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteAppResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteAppResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAppResponse) GoString() string {
	return s.String()
}

func (s *DeleteAppResponse) SetHeaders(v map[string]*string) *DeleteAppResponse {
	s.Headers = v
	return s
}

func (s *DeleteAppResponse) SetStatusCode(v int32) *DeleteAppResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAppResponse) SetBody(v *DeleteAppResponseBody) *DeleteAppResponse {
	s.Body = v
	return s
}

type DeleteAppVersionRequest struct {
	AppVersionId *string `json:"AppVersionId,omitempty" xml:"AppVersionId,omitempty"`
}

func (s DeleteAppVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAppVersionRequest) GoString() string {
	return s.String()
}

func (s *DeleteAppVersionRequest) SetAppVersionId(v string) *DeleteAppVersionRequest {
	s.AppVersionId = &v
	return s
}

type DeleteAppVersionResponseBody struct {
	AppVersionId *string `json:"AppVersionId,omitempty" xml:"AppVersionId,omitempty"`
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpCode     *int32  `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Message      *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteAppVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAppVersionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAppVersionResponseBody) SetAppVersionId(v string) *DeleteAppVersionResponseBody {
	s.AppVersionId = &v
	return s
}

func (s *DeleteAppVersionResponseBody) SetCode(v string) *DeleteAppVersionResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteAppVersionResponseBody) SetHttpCode(v int32) *DeleteAppVersionResponseBody {
	s.HttpCode = &v
	return s
}

func (s *DeleteAppVersionResponseBody) SetMessage(v string) *DeleteAppVersionResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteAppVersionResponseBody) SetRequestId(v string) *DeleteAppVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAppVersionResponseBody) SetSuccess(v bool) *DeleteAppVersionResponseBody {
	s.Success = &v
	return s
}

type DeleteAppVersionResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteAppVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteAppVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAppVersionResponse) GoString() string {
	return s.String()
}

func (s *DeleteAppVersionResponse) SetHeaders(v map[string]*string) *DeleteAppVersionResponse {
	s.Headers = v
	return s
}

func (s *DeleteAppVersionResponse) SetStatusCode(v int32) *DeleteAppVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAppVersionResponse) SetBody(v *DeleteAppVersionResponseBody) *DeleteAppVersionResponse {
	s.Body = v
	return s
}

type DeleteProjectRequest struct {
	// 请求操作人Id
	OperatorId *string `json:"OperatorId,omitempty" xml:"OperatorId,omitempty"`
	// 请求操作人类型
	OperatorType *string `json:"OperatorType,omitempty" xml:"OperatorType,omitempty"`
	// project Id
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s DeleteProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteProjectRequest) GoString() string {
	return s.String()
}

func (s *DeleteProjectRequest) SetOperatorId(v string) *DeleteProjectRequest {
	s.OperatorId = &v
	return s
}

func (s *DeleteProjectRequest) SetOperatorType(v string) *DeleteProjectRequest {
	s.OperatorType = &v
	return s
}

func (s *DeleteProjectRequest) SetProjectId(v string) *DeleteProjectRequest {
	s.ProjectId = &v
	return s
}

type DeleteProjectResponseBody struct {
	// 业务处理结果Code
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 业务对象
	Data *DeleteProjectResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// 业务处理消息摘要
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 操作请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 业务处理是否成功
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteProjectResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteProjectResponseBody) SetCode(v string) *DeleteProjectResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteProjectResponseBody) SetData(v *DeleteProjectResponseBodyData) *DeleteProjectResponseBody {
	s.Data = v
	return s
}

func (s *DeleteProjectResponseBody) SetMessage(v string) *DeleteProjectResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteProjectResponseBody) SetRequestId(v string) *DeleteProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteProjectResponseBody) SetSuccess(v bool) *DeleteProjectResponseBody {
	s.Success = &v
	return s
}

type DeleteProjectResponseBodyData struct {
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteProjectResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DeleteProjectResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteProjectResponseBodyData) SetSuccess(v bool) *DeleteProjectResponseBodyData {
	s.Success = &v
	return s
}

type DeleteProjectResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteProjectResponse) GoString() string {
	return s.String()
}

func (s *DeleteProjectResponse) SetHeaders(v map[string]*string) *DeleteProjectResponse {
	s.Headers = v
	return s
}

func (s *DeleteProjectResponse) SetStatusCode(v int32) *DeleteProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteProjectResponse) SetBody(v *DeleteProjectResponseBody) *DeleteProjectResponse {
	s.Body = v
	return s
}

type GetAdaptationRequest struct {
	AdaptApplyId *int64  `json:"AdaptApplyId,omitempty" xml:"AdaptApplyId,omitempty"`
	AppVersionId *string `json:"AppVersionId,omitempty" xml:"AppVersionId,omitempty"`
}

func (s GetAdaptationRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAdaptationRequest) GoString() string {
	return s.String()
}

func (s *GetAdaptationRequest) SetAdaptApplyId(v int64) *GetAdaptationRequest {
	s.AdaptApplyId = &v
	return s
}

func (s *GetAdaptationRequest) SetAppVersionId(v string) *GetAdaptationRequest {
	s.AppVersionId = &v
	return s
}

type GetAdaptationResponseBody struct {
	AdaptApplyId *int64                                `json:"AdaptApplyId,omitempty" xml:"AdaptApplyId,omitempty"`
	AdaptTarget  *GetAdaptationResponseBodyAdaptTarget `json:"AdaptTarget,omitempty" xml:"AdaptTarget,omitempty" type:"Struct"`
	AppId        *string                               `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppVersionId *string                               `json:"AppVersionId,omitempty" xml:"AppVersionId,omitempty"`
	Code         *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	GmtCreate    *string                               `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified  *string                               `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	HttpCode     *int32                                `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Message      *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId    *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAdaptationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAdaptationResponseBody) GoString() string {
	return s.String()
}

func (s *GetAdaptationResponseBody) SetAdaptApplyId(v int64) *GetAdaptationResponseBody {
	s.AdaptApplyId = &v
	return s
}

func (s *GetAdaptationResponseBody) SetAdaptTarget(v *GetAdaptationResponseBodyAdaptTarget) *GetAdaptationResponseBody {
	s.AdaptTarget = v
	return s
}

func (s *GetAdaptationResponseBody) SetAppId(v string) *GetAdaptationResponseBody {
	s.AppId = &v
	return s
}

func (s *GetAdaptationResponseBody) SetAppVersionId(v string) *GetAdaptationResponseBody {
	s.AppVersionId = &v
	return s
}

func (s *GetAdaptationResponseBody) SetCode(v string) *GetAdaptationResponseBody {
	s.Code = &v
	return s
}

func (s *GetAdaptationResponseBody) SetGmtCreate(v string) *GetAdaptationResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *GetAdaptationResponseBody) SetGmtModified(v string) *GetAdaptationResponseBody {
	s.GmtModified = &v
	return s
}

func (s *GetAdaptationResponseBody) SetHttpCode(v int32) *GetAdaptationResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetAdaptationResponseBody) SetMessage(v string) *GetAdaptationResponseBody {
	s.Message = &v
	return s
}

func (s *GetAdaptationResponseBody) SetRequestId(v string) *GetAdaptationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAdaptationResponseBody) SetSuccess(v bool) *GetAdaptationResponseBody {
	s.Success = &v
	return s
}

type GetAdaptationResponseBodyAdaptTarget struct {
	BitRate      *int32  `json:"BitRate,omitempty" xml:"BitRate,omitempty"`
	FrameRate    *int32  `json:"FrameRate,omitempty" xml:"FrameRate,omitempty"`
	Resolution   *string `json:"Resolution,omitempty" xml:"Resolution,omitempty"`
	StartProgram *string `json:"StartProgram,omitempty" xml:"StartProgram,omitempty"`
}

func (s GetAdaptationResponseBodyAdaptTarget) String() string {
	return tea.Prettify(s)
}

func (s GetAdaptationResponseBodyAdaptTarget) GoString() string {
	return s.String()
}

func (s *GetAdaptationResponseBodyAdaptTarget) SetBitRate(v int32) *GetAdaptationResponseBodyAdaptTarget {
	s.BitRate = &v
	return s
}

func (s *GetAdaptationResponseBodyAdaptTarget) SetFrameRate(v int32) *GetAdaptationResponseBodyAdaptTarget {
	s.FrameRate = &v
	return s
}

func (s *GetAdaptationResponseBodyAdaptTarget) SetResolution(v string) *GetAdaptationResponseBodyAdaptTarget {
	s.Resolution = &v
	return s
}

func (s *GetAdaptationResponseBodyAdaptTarget) SetStartProgram(v string) *GetAdaptationResponseBodyAdaptTarget {
	s.StartProgram = &v
	return s
}

type GetAdaptationResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetAdaptationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAdaptationResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAdaptationResponse) GoString() string {
	return s.String()
}

func (s *GetAdaptationResponse) SetHeaders(v map[string]*string) *GetAdaptationResponse {
	s.Headers = v
	return s
}

func (s *GetAdaptationResponse) SetStatusCode(v int32) *GetAdaptationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAdaptationResponse) SetBody(v *GetAdaptationResponseBody) *GetAdaptationResponse {
	s.Body = v
	return s
}

type GetAppRequest struct {
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s GetAppRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAppRequest) GoString() string {
	return s.String()
}

func (s *GetAppRequest) SetAppId(v string) *GetAppRequest {
	s.AppId = &v
	return s
}

type GetAppResponseBody struct {
	AppId           *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppName         *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	AppType         *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	Code            *string `json:"Code,omitempty" xml:"Code,omitempty"`
	GmtCreate       *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified     *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	HttpCode        *int32  `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Message         *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success         *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	VersionAdaptNum *int64  `json:"VersionAdaptNum,omitempty" xml:"VersionAdaptNum,omitempty"`
	VersionTotalNum *int64  `json:"VersionTotalNum,omitempty" xml:"VersionTotalNum,omitempty"`
}

func (s GetAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAppResponseBody) GoString() string {
	return s.String()
}

func (s *GetAppResponseBody) SetAppId(v string) *GetAppResponseBody {
	s.AppId = &v
	return s
}

func (s *GetAppResponseBody) SetAppName(v string) *GetAppResponseBody {
	s.AppName = &v
	return s
}

func (s *GetAppResponseBody) SetAppType(v string) *GetAppResponseBody {
	s.AppType = &v
	return s
}

func (s *GetAppResponseBody) SetCode(v string) *GetAppResponseBody {
	s.Code = &v
	return s
}

func (s *GetAppResponseBody) SetGmtCreate(v string) *GetAppResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *GetAppResponseBody) SetGmtModified(v string) *GetAppResponseBody {
	s.GmtModified = &v
	return s
}

func (s *GetAppResponseBody) SetHttpCode(v int32) *GetAppResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetAppResponseBody) SetMessage(v string) *GetAppResponseBody {
	s.Message = &v
	return s
}

func (s *GetAppResponseBody) SetRequestId(v string) *GetAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAppResponseBody) SetSuccess(v bool) *GetAppResponseBody {
	s.Success = &v
	return s
}

func (s *GetAppResponseBody) SetVersionAdaptNum(v int64) *GetAppResponseBody {
	s.VersionAdaptNum = &v
	return s
}

func (s *GetAppResponseBody) SetVersionTotalNum(v int64) *GetAppResponseBody {
	s.VersionTotalNum = &v
	return s
}

type GetAppResponse struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetAppResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAppResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAppResponse) GoString() string {
	return s.String()
}

func (s *GetAppResponse) SetHeaders(v map[string]*string) *GetAppResponse {
	s.Headers = v
	return s
}

func (s *GetAppResponse) SetStatusCode(v int32) *GetAppResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAppResponse) SetBody(v *GetAppResponseBody) *GetAppResponse {
	s.Body = v
	return s
}

type GetAppSessionRequest struct {
	// 自定义会话id
	CustomSessionId *string `json:"CustomSessionId,omitempty" xml:"CustomSessionId,omitempty"`
	// 平台会话id
	PlatformSessionId *string `json:"PlatformSessionId,omitempty" xml:"PlatformSessionId,omitempty"`
}

func (s GetAppSessionRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAppSessionRequest) GoString() string {
	return s.String()
}

func (s *GetAppSessionRequest) SetCustomSessionId(v string) *GetAppSessionRequest {
	s.CustomSessionId = &v
	return s
}

func (s *GetAppSessionRequest) SetPlatformSessionId(v string) *GetAppSessionRequest {
	s.PlatformSessionId = &v
	return s
}

type GetAppSessionResponseBody struct {
	// 应用id
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 应用版本
	AppVersion *string                             `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	BizInfo    []*GetAppSessionResponseBodyBizInfo `json:"BizInfo,omitempty" xml:"BizInfo,omitempty" type:"Repeated"`
	// 自定义会话id
	CustomSessionId *string `json:"CustomSessionId,omitempty" xml:"CustomSessionId,omitempty"`
	// 平台会话id
	PlatformSessionId *string `json:"PlatformSessionId,omitempty" xml:"PlatformSessionId,omitempty"`
	// 请求id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetAppSessionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAppSessionResponseBody) GoString() string {
	return s.String()
}

func (s *GetAppSessionResponseBody) SetAppId(v string) *GetAppSessionResponseBody {
	s.AppId = &v
	return s
}

func (s *GetAppSessionResponseBody) SetAppVersion(v string) *GetAppSessionResponseBody {
	s.AppVersion = &v
	return s
}

func (s *GetAppSessionResponseBody) SetBizInfo(v []*GetAppSessionResponseBodyBizInfo) *GetAppSessionResponseBody {
	s.BizInfo = v
	return s
}

func (s *GetAppSessionResponseBody) SetCustomSessionId(v string) *GetAppSessionResponseBody {
	s.CustomSessionId = &v
	return s
}

func (s *GetAppSessionResponseBody) SetPlatformSessionId(v string) *GetAppSessionResponseBody {
	s.PlatformSessionId = &v
	return s
}

func (s *GetAppSessionResponseBody) SetRequestId(v string) *GetAppSessionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAppSessionResponseBody) SetStatus(v string) *GetAppSessionResponseBody {
	s.Status = &v
	return s
}

type GetAppSessionResponseBodyBizInfo struct {
	Key   *string `json:"key,omitempty" xml:"key,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s GetAppSessionResponseBodyBizInfo) String() string {
	return tea.Prettify(s)
}

func (s GetAppSessionResponseBodyBizInfo) GoString() string {
	return s.String()
}

func (s *GetAppSessionResponseBodyBizInfo) SetKey(v string) *GetAppSessionResponseBodyBizInfo {
	s.Key = &v
	return s
}

func (s *GetAppSessionResponseBodyBizInfo) SetValue(v string) *GetAppSessionResponseBodyBizInfo {
	s.Value = &v
	return s
}

type GetAppSessionResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetAppSessionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAppSessionResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAppSessionResponse) GoString() string {
	return s.String()
}

func (s *GetAppSessionResponse) SetHeaders(v map[string]*string) *GetAppSessionResponse {
	s.Headers = v
	return s
}

func (s *GetAppSessionResponse) SetStatusCode(v int32) *GetAppSessionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAppSessionResponse) SetBody(v *GetAppSessionResponseBody) *GetAppSessionResponse {
	s.Body = v
	return s
}

type GetAppVersionRequest struct {
	AppVersionId *string `json:"AppVersionId,omitempty" xml:"AppVersionId,omitempty"`
}

func (s GetAppVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAppVersionRequest) GoString() string {
	return s.String()
}

func (s *GetAppVersionRequest) SetAppVersionId(v string) *GetAppVersionRequest {
	s.AppVersionId = &v
	return s
}

type GetAppVersionResponseBody struct {
	AppId                *string  `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppVersionId         *string  `json:"AppVersionId,omitempty" xml:"AppVersionId,omitempty"`
	AppVersionName       *string  `json:"AppVersionName,omitempty" xml:"AppVersionName,omitempty"`
	AppVersionStatus     *string  `json:"AppVersionStatus,omitempty" xml:"AppVersionStatus,omitempty"`
	AppVersionStatusMemo *string  `json:"AppVersionStatusMemo,omitempty" xml:"AppVersionStatusMemo,omitempty"`
	Code                 *string  `json:"Code,omitempty" xml:"Code,omitempty"`
	ConsumeCu            *float64 `json:"ConsumeCu,omitempty" xml:"ConsumeCu,omitempty"`
	FileAddress          *string  `json:"FileAddress,omitempty" xml:"FileAddress,omitempty"`
	FileSize             *int64   `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	FileUploadFinishTime *string  `json:"FileUploadFinishTime,omitempty" xml:"FileUploadFinishTime,omitempty"`
	FileUploadType       *string  `json:"FileUploadType,omitempty" xml:"FileUploadType,omitempty"`
	GmtCreate            *string  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified          *string  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	HttpCode             *int32   `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Message              *string  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId            *string  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success              *bool    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAppVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAppVersionResponseBody) GoString() string {
	return s.String()
}

func (s *GetAppVersionResponseBody) SetAppId(v string) *GetAppVersionResponseBody {
	s.AppId = &v
	return s
}

func (s *GetAppVersionResponseBody) SetAppVersionId(v string) *GetAppVersionResponseBody {
	s.AppVersionId = &v
	return s
}

func (s *GetAppVersionResponseBody) SetAppVersionName(v string) *GetAppVersionResponseBody {
	s.AppVersionName = &v
	return s
}

func (s *GetAppVersionResponseBody) SetAppVersionStatus(v string) *GetAppVersionResponseBody {
	s.AppVersionStatus = &v
	return s
}

func (s *GetAppVersionResponseBody) SetAppVersionStatusMemo(v string) *GetAppVersionResponseBody {
	s.AppVersionStatusMemo = &v
	return s
}

func (s *GetAppVersionResponseBody) SetCode(v string) *GetAppVersionResponseBody {
	s.Code = &v
	return s
}

func (s *GetAppVersionResponseBody) SetConsumeCu(v float64) *GetAppVersionResponseBody {
	s.ConsumeCu = &v
	return s
}

func (s *GetAppVersionResponseBody) SetFileAddress(v string) *GetAppVersionResponseBody {
	s.FileAddress = &v
	return s
}

func (s *GetAppVersionResponseBody) SetFileSize(v int64) *GetAppVersionResponseBody {
	s.FileSize = &v
	return s
}

func (s *GetAppVersionResponseBody) SetFileUploadFinishTime(v string) *GetAppVersionResponseBody {
	s.FileUploadFinishTime = &v
	return s
}

func (s *GetAppVersionResponseBody) SetFileUploadType(v string) *GetAppVersionResponseBody {
	s.FileUploadType = &v
	return s
}

func (s *GetAppVersionResponseBody) SetGmtCreate(v string) *GetAppVersionResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *GetAppVersionResponseBody) SetGmtModified(v string) *GetAppVersionResponseBody {
	s.GmtModified = &v
	return s
}

func (s *GetAppVersionResponseBody) SetHttpCode(v int32) *GetAppVersionResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetAppVersionResponseBody) SetMessage(v string) *GetAppVersionResponseBody {
	s.Message = &v
	return s
}

func (s *GetAppVersionResponseBody) SetRequestId(v string) *GetAppVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAppVersionResponseBody) SetSuccess(v bool) *GetAppVersionResponseBody {
	s.Success = &v
	return s
}

type GetAppVersionResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetAppVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAppVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAppVersionResponse) GoString() string {
	return s.String()
}

func (s *GetAppVersionResponse) SetHeaders(v map[string]*string) *GetAppVersionResponse {
	s.Headers = v
	return s
}

func (s *GetAppVersionResponse) SetStatusCode(v int32) *GetAppVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAppVersionResponse) SetBody(v *GetAppVersionResponseBody) *GetAppVersionResponse {
	s.Body = v
	return s
}

type GetDatasetRequest struct {
	DatasetId *string `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
}

func (s GetDatasetRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDatasetRequest) GoString() string {
	return s.String()
}

func (s *GetDatasetRequest) SetDatasetId(v string) *GetDatasetRequest {
	s.DatasetId = &v
	return s
}

type GetDatasetResponseBody struct {
	// 应用id
	CustomParam *string `json:"CustomParam,omitempty" xml:"CustomParam,omitempty"`
	// 自定义会话id
	DatasetId   *string `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	DatasetSize *int64  `json:"DatasetSize,omitempty" xml:"DatasetSize,omitempty"`
	// 请求id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDatasetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDatasetResponseBody) GoString() string {
	return s.String()
}

func (s *GetDatasetResponseBody) SetCustomParam(v string) *GetDatasetResponseBody {
	s.CustomParam = &v
	return s
}

func (s *GetDatasetResponseBody) SetDatasetId(v string) *GetDatasetResponseBody {
	s.DatasetId = &v
	return s
}

func (s *GetDatasetResponseBody) SetDatasetSize(v int64) *GetDatasetResponseBody {
	s.DatasetSize = &v
	return s
}

func (s *GetDatasetResponseBody) SetRequestId(v string) *GetDatasetResponseBody {
	s.RequestId = &v
	return s
}

type GetDatasetResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetDatasetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDatasetResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDatasetResponse) GoString() string {
	return s.String()
}

func (s *GetDatasetResponse) SetHeaders(v map[string]*string) *GetDatasetResponse {
	s.Headers = v
	return s
}

func (s *GetDatasetResponse) SetStatusCode(v int32) *GetDatasetResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDatasetResponse) SetBody(v *GetDatasetResponseBody) *GetDatasetResponse {
	s.Body = v
	return s
}

type GetProjectRequest struct {
	// 请求操作人Id
	OperatorId *string `json:"OperatorId,omitempty" xml:"OperatorId,omitempty"`
	// 请求操作人类型
	OperatorType *string `json:"OperatorType,omitempty" xml:"OperatorType,omitempty"`
	// project id
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s GetProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s GetProjectRequest) GoString() string {
	return s.String()
}

func (s *GetProjectRequest) SetOperatorId(v string) *GetProjectRequest {
	s.OperatorId = &v
	return s
}

func (s *GetProjectRequest) SetOperatorType(v string) *GetProjectRequest {
	s.OperatorType = &v
	return s
}

func (s *GetProjectRequest) SetProjectId(v string) *GetProjectRequest {
	s.ProjectId = &v
	return s
}

type GetProjectResponseBody struct {
	// 业务处理结果Code
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 业务对象
	Data *GetProjectResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// 业务处理消息摘要
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 操作请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 业务处理是否成功
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetProjectResponseBody) GoString() string {
	return s.String()
}

func (s *GetProjectResponseBody) SetCode(v string) *GetProjectResponseBody {
	s.Code = &v
	return s
}

func (s *GetProjectResponseBody) SetData(v *GetProjectResponseBodyData) *GetProjectResponseBody {
	s.Data = v
	return s
}

func (s *GetProjectResponseBody) SetMessage(v string) *GetProjectResponseBody {
	s.Message = &v
	return s
}

func (s *GetProjectResponseBody) SetRequestId(v string) *GetProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetProjectResponseBody) SetSuccess(v bool) *GetProjectResponseBody {
	s.Success = &v
	return s
}

type GetProjectResponseBodyData struct {
	// 项目关联的应用数量
	BoundAppNums *int64  `json:"BoundAppNums,omitempty" xml:"BoundAppNums,omitempty"`
	GmtCreate    *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified  *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	ProjectId    *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	ProjectMemo  *string `json:"ProjectMemo,omitempty" xml:"ProjectMemo,omitempty"`
	ProjectName  *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// key : districtId
	ProjectQuotaLimit *GetProjectResponseBodyDataProjectQuotaLimit `json:"ProjectQuotaLimit,omitempty" xml:"ProjectQuotaLimit,omitempty" type:"Struct"`
}

func (s GetProjectResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetProjectResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetProjectResponseBodyData) SetBoundAppNums(v int64) *GetProjectResponseBodyData {
	s.BoundAppNums = &v
	return s
}

func (s *GetProjectResponseBodyData) SetGmtCreate(v string) *GetProjectResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *GetProjectResponseBodyData) SetGmtModified(v string) *GetProjectResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *GetProjectResponseBodyData) SetProjectId(v string) *GetProjectResponseBodyData {
	s.ProjectId = &v
	return s
}

func (s *GetProjectResponseBodyData) SetProjectMemo(v string) *GetProjectResponseBodyData {
	s.ProjectMemo = &v
	return s
}

func (s *GetProjectResponseBodyData) SetProjectName(v string) *GetProjectResponseBodyData {
	s.ProjectName = &v
	return s
}

func (s *GetProjectResponseBodyData) SetProjectQuotaLimit(v *GetProjectResponseBodyDataProjectQuotaLimit) *GetProjectResponseBodyData {
	s.ProjectQuotaLimit = v
	return s
}

type GetProjectResponseBodyDataProjectQuotaLimit struct {
	// key - districtId
	DistrictLimitMap map[string]*DataProjectQuotaLimitDistrictLimitMapValue `json:"DistrictLimitMap,omitempty" xml:"DistrictLimitMap,omitempty"`
	// 限制类型 ：目前默认 - ReserveContainer
	LimitType *string `json:"LimitType,omitempty" xml:"LimitType,omitempty"`
}

func (s GetProjectResponseBodyDataProjectQuotaLimit) String() string {
	return tea.Prettify(s)
}

func (s GetProjectResponseBodyDataProjectQuotaLimit) GoString() string {
	return s.String()
}

func (s *GetProjectResponseBodyDataProjectQuotaLimit) SetDistrictLimitMap(v map[string]*DataProjectQuotaLimitDistrictLimitMapValue) *GetProjectResponseBodyDataProjectQuotaLimit {
	s.DistrictLimitMap = v
	return s
}

func (s *GetProjectResponseBodyDataProjectQuotaLimit) SetLimitType(v string) *GetProjectResponseBodyDataProjectQuotaLimit {
	s.LimitType = &v
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

type ListAppRequest struct {
	KeySearch  *string `json:"KeySearch,omitempty" xml:"KeySearch,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListAppRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAppRequest) GoString() string {
	return s.String()
}

func (s *ListAppRequest) SetKeySearch(v string) *ListAppRequest {
	s.KeySearch = &v
	return s
}

func (s *ListAppRequest) SetPageNumber(v int32) *ListAppRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAppRequest) SetPageSize(v int32) *ListAppRequest {
	s.PageSize = &v
	return s
}

type ListAppResponseBody struct {
	Apps      []*ListAppResponseBodyApps `json:"Apps,omitempty" xml:"Apps,omitempty" type:"Repeated"`
	Code      *string                    `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpCode  *int32                     `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Message   *string                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                      `json:"Success,omitempty" xml:"Success,omitempty"`
	Total     *int64                     `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAppResponseBody) GoString() string {
	return s.String()
}

func (s *ListAppResponseBody) SetApps(v []*ListAppResponseBodyApps) *ListAppResponseBody {
	s.Apps = v
	return s
}

func (s *ListAppResponseBody) SetCode(v string) *ListAppResponseBody {
	s.Code = &v
	return s
}

func (s *ListAppResponseBody) SetHttpCode(v int32) *ListAppResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ListAppResponseBody) SetMessage(v string) *ListAppResponseBody {
	s.Message = &v
	return s
}

func (s *ListAppResponseBody) SetRequestId(v string) *ListAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAppResponseBody) SetSuccess(v bool) *ListAppResponseBody {
	s.Success = &v
	return s
}

func (s *ListAppResponseBody) SetTotal(v int64) *ListAppResponseBody {
	s.Total = &v
	return s
}

type ListAppResponseBodyApps struct {
	AppId           *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppName         *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	AppType         *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	GmtCreate       *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified     *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	VersionAdaptNum *int64  `json:"VersionAdaptNum,omitempty" xml:"VersionAdaptNum,omitempty"`
	VersionTotalNum *int64  `json:"VersionTotalNum,omitempty" xml:"VersionTotalNum,omitempty"`
}

func (s ListAppResponseBodyApps) String() string {
	return tea.Prettify(s)
}

func (s ListAppResponseBodyApps) GoString() string {
	return s.String()
}

func (s *ListAppResponseBodyApps) SetAppId(v string) *ListAppResponseBodyApps {
	s.AppId = &v
	return s
}

func (s *ListAppResponseBodyApps) SetAppName(v string) *ListAppResponseBodyApps {
	s.AppName = &v
	return s
}

func (s *ListAppResponseBodyApps) SetAppType(v string) *ListAppResponseBodyApps {
	s.AppType = &v
	return s
}

func (s *ListAppResponseBodyApps) SetGmtCreate(v string) *ListAppResponseBodyApps {
	s.GmtCreate = &v
	return s
}

func (s *ListAppResponseBodyApps) SetGmtModified(v string) *ListAppResponseBodyApps {
	s.GmtModified = &v
	return s
}

func (s *ListAppResponseBodyApps) SetVersionAdaptNum(v int64) *ListAppResponseBodyApps {
	s.VersionAdaptNum = &v
	return s
}

func (s *ListAppResponseBodyApps) SetVersionTotalNum(v int64) *ListAppResponseBodyApps {
	s.VersionTotalNum = &v
	return s
}

type ListAppResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListAppResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAppResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAppResponse) GoString() string {
	return s.String()
}

func (s *ListAppResponse) SetHeaders(v map[string]*string) *ListAppResponse {
	s.Headers = v
	return s
}

func (s *ListAppResponse) SetStatusCode(v int32) *ListAppResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAppResponse) SetBody(v *ListAppResponseBody) *ListAppResponse {
	s.Body = v
	return s
}

type ListAppSessionsRequest struct {
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 自定义会话id
	CustomSessionIds []*string `json:"CustomSessionIds,omitempty" xml:"CustomSessionIds,omitempty" type:"Repeated"`
	// 页码
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 分页大小
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 自定义用户id
	PlatformSessionIds []*string `json:"PlatformSessionIds,omitempty" xml:"PlatformSessionIds,omitempty" type:"Repeated"`
}

func (s ListAppSessionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAppSessionsRequest) GoString() string {
	return s.String()
}

func (s *ListAppSessionsRequest) SetAppId(v string) *ListAppSessionsRequest {
	s.AppId = &v
	return s
}

func (s *ListAppSessionsRequest) SetCustomSessionIds(v []*string) *ListAppSessionsRequest {
	s.CustomSessionIds = v
	return s
}

func (s *ListAppSessionsRequest) SetPageNumber(v int32) *ListAppSessionsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAppSessionsRequest) SetPageSize(v int32) *ListAppSessionsRequest {
	s.PageSize = &v
	return s
}

func (s *ListAppSessionsRequest) SetPlatformSessionIds(v []*string) *ListAppSessionsRequest {
	s.PlatformSessionIds = v
	return s
}

type ListAppSessionsResponseBody struct {
	AppSessions []*ListAppSessionsResponseBodyAppSessions `json:"AppSessions,omitempty" xml:"AppSessions,omitempty" type:"Repeated"`
	PageNumber  *int32                                    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *int32                                    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 请求id
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAppSessionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAppSessionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAppSessionsResponseBody) SetAppSessions(v []*ListAppSessionsResponseBodyAppSessions) *ListAppSessionsResponseBody {
	s.AppSessions = v
	return s
}

func (s *ListAppSessionsResponseBody) SetPageNumber(v int32) *ListAppSessionsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListAppSessionsResponseBody) SetPageSize(v int32) *ListAppSessionsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListAppSessionsResponseBody) SetRequestId(v string) *ListAppSessionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAppSessionsResponseBody) SetTotalCount(v int32) *ListAppSessionsResponseBody {
	s.TotalCount = &v
	return s
}

type ListAppSessionsResponseBodyAppSessions struct {
	// 应用id
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 应用版本
	AppVersion *string                                          `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	BizInfo    []*ListAppSessionsResponseBodyAppSessionsBizInfo `json:"BizInfo,omitempty" xml:"BizInfo,omitempty" type:"Repeated"`
	// 自定义会话id
	CustomSessionId *string `json:"CustomSessionId,omitempty" xml:"CustomSessionId,omitempty"`
	// 平台会话id
	PlatformSessionId *string `json:"PlatformSessionId,omitempty" xml:"PlatformSessionId,omitempty"`
	// 状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListAppSessionsResponseBodyAppSessions) String() string {
	return tea.Prettify(s)
}

func (s ListAppSessionsResponseBodyAppSessions) GoString() string {
	return s.String()
}

func (s *ListAppSessionsResponseBodyAppSessions) SetAppId(v string) *ListAppSessionsResponseBodyAppSessions {
	s.AppId = &v
	return s
}

func (s *ListAppSessionsResponseBodyAppSessions) SetAppVersion(v string) *ListAppSessionsResponseBodyAppSessions {
	s.AppVersion = &v
	return s
}

func (s *ListAppSessionsResponseBodyAppSessions) SetBizInfo(v []*ListAppSessionsResponseBodyAppSessionsBizInfo) *ListAppSessionsResponseBodyAppSessions {
	s.BizInfo = v
	return s
}

func (s *ListAppSessionsResponseBodyAppSessions) SetCustomSessionId(v string) *ListAppSessionsResponseBodyAppSessions {
	s.CustomSessionId = &v
	return s
}

func (s *ListAppSessionsResponseBodyAppSessions) SetPlatformSessionId(v string) *ListAppSessionsResponseBodyAppSessions {
	s.PlatformSessionId = &v
	return s
}

func (s *ListAppSessionsResponseBodyAppSessions) SetStatus(v string) *ListAppSessionsResponseBodyAppSessions {
	s.Status = &v
	return s
}

type ListAppSessionsResponseBodyAppSessionsBizInfo struct {
	Key   *string `json:"key,omitempty" xml:"key,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListAppSessionsResponseBodyAppSessionsBizInfo) String() string {
	return tea.Prettify(s)
}

func (s ListAppSessionsResponseBodyAppSessionsBizInfo) GoString() string {
	return s.String()
}

func (s *ListAppSessionsResponseBodyAppSessionsBizInfo) SetKey(v string) *ListAppSessionsResponseBodyAppSessionsBizInfo {
	s.Key = &v
	return s
}

func (s *ListAppSessionsResponseBodyAppSessionsBizInfo) SetValue(v string) *ListAppSessionsResponseBodyAppSessionsBizInfo {
	s.Value = &v
	return s
}

type ListAppSessionsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListAppSessionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAppSessionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAppSessionsResponse) GoString() string {
	return s.String()
}

func (s *ListAppSessionsResponse) SetHeaders(v map[string]*string) *ListAppSessionsResponse {
	s.Headers = v
	return s
}

func (s *ListAppSessionsResponse) SetStatusCode(v int32) *ListAppSessionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAppSessionsResponse) SetBody(v *ListAppSessionsResponseBody) *ListAppSessionsResponse {
	s.Body = v
	return s
}

type ListAppVersionRequest struct {
	AppId      *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListAppVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAppVersionRequest) GoString() string {
	return s.String()
}

func (s *ListAppVersionRequest) SetAppId(v string) *ListAppVersionRequest {
	s.AppId = &v
	return s
}

func (s *ListAppVersionRequest) SetPageNumber(v int32) *ListAppVersionRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAppVersionRequest) SetPageSize(v int32) *ListAppVersionRequest {
	s.PageSize = &v
	return s
}

type ListAppVersionResponseBody struct {
	Code      *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpCode  *int32                                `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Message   *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
	Total     *int64                                `json:"Total,omitempty" xml:"Total,omitempty"`
	Versions  []*ListAppVersionResponseBodyVersions `json:"Versions,omitempty" xml:"Versions,omitempty" type:"Repeated"`
}

func (s ListAppVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAppVersionResponseBody) GoString() string {
	return s.String()
}

func (s *ListAppVersionResponseBody) SetCode(v string) *ListAppVersionResponseBody {
	s.Code = &v
	return s
}

func (s *ListAppVersionResponseBody) SetHttpCode(v int32) *ListAppVersionResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ListAppVersionResponseBody) SetMessage(v string) *ListAppVersionResponseBody {
	s.Message = &v
	return s
}

func (s *ListAppVersionResponseBody) SetRequestId(v string) *ListAppVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAppVersionResponseBody) SetSuccess(v bool) *ListAppVersionResponseBody {
	s.Success = &v
	return s
}

func (s *ListAppVersionResponseBody) SetTotal(v int64) *ListAppVersionResponseBody {
	s.Total = &v
	return s
}

func (s *ListAppVersionResponseBody) SetVersions(v []*ListAppVersionResponseBodyVersions) *ListAppVersionResponseBody {
	s.Versions = v
	return s
}

type ListAppVersionResponseBodyVersions struct {
	AppId                *string  `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppVersionId         *string  `json:"AppVersionId,omitempty" xml:"AppVersionId,omitempty"`
	AppVersionName       *string  `json:"AppVersionName,omitempty" xml:"AppVersionName,omitempty"`
	AppVersionStatus     *string  `json:"AppVersionStatus,omitempty" xml:"AppVersionStatus,omitempty"`
	AppVersionStatusMemo *string  `json:"AppVersionStatusMemo,omitempty" xml:"AppVersionStatusMemo,omitempty"`
	ConsumeCu            *float64 `json:"ConsumeCu,omitempty" xml:"ConsumeCu,omitempty"`
	FileAddress          *string  `json:"FileAddress,omitempty" xml:"FileAddress,omitempty"`
	FileSize             *int64   `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	FileUploadFinishTime *string  `json:"FileUploadFinishTime,omitempty" xml:"FileUploadFinishTime,omitempty"`
	FileUploadType       *string  `json:"FileUploadType,omitempty" xml:"FileUploadType,omitempty"`
	GmtCreate            *string  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified          *string  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	TenantId             *int64   `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s ListAppVersionResponseBodyVersions) String() string {
	return tea.Prettify(s)
}

func (s ListAppVersionResponseBodyVersions) GoString() string {
	return s.String()
}

func (s *ListAppVersionResponseBodyVersions) SetAppId(v string) *ListAppVersionResponseBodyVersions {
	s.AppId = &v
	return s
}

func (s *ListAppVersionResponseBodyVersions) SetAppVersionId(v string) *ListAppVersionResponseBodyVersions {
	s.AppVersionId = &v
	return s
}

func (s *ListAppVersionResponseBodyVersions) SetAppVersionName(v string) *ListAppVersionResponseBodyVersions {
	s.AppVersionName = &v
	return s
}

func (s *ListAppVersionResponseBodyVersions) SetAppVersionStatus(v string) *ListAppVersionResponseBodyVersions {
	s.AppVersionStatus = &v
	return s
}

func (s *ListAppVersionResponseBodyVersions) SetAppVersionStatusMemo(v string) *ListAppVersionResponseBodyVersions {
	s.AppVersionStatusMemo = &v
	return s
}

func (s *ListAppVersionResponseBodyVersions) SetConsumeCu(v float64) *ListAppVersionResponseBodyVersions {
	s.ConsumeCu = &v
	return s
}

func (s *ListAppVersionResponseBodyVersions) SetFileAddress(v string) *ListAppVersionResponseBodyVersions {
	s.FileAddress = &v
	return s
}

func (s *ListAppVersionResponseBodyVersions) SetFileSize(v int64) *ListAppVersionResponseBodyVersions {
	s.FileSize = &v
	return s
}

func (s *ListAppVersionResponseBodyVersions) SetFileUploadFinishTime(v string) *ListAppVersionResponseBodyVersions {
	s.FileUploadFinishTime = &v
	return s
}

func (s *ListAppVersionResponseBodyVersions) SetFileUploadType(v string) *ListAppVersionResponseBodyVersions {
	s.FileUploadType = &v
	return s
}

func (s *ListAppVersionResponseBodyVersions) SetGmtCreate(v string) *ListAppVersionResponseBodyVersions {
	s.GmtCreate = &v
	return s
}

func (s *ListAppVersionResponseBodyVersions) SetGmtModified(v string) *ListAppVersionResponseBodyVersions {
	s.GmtModified = &v
	return s
}

func (s *ListAppVersionResponseBodyVersions) SetTenantId(v int64) *ListAppVersionResponseBodyVersions {
	s.TenantId = &v
	return s
}

type ListAppVersionResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListAppVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAppVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAppVersionResponse) GoString() string {
	return s.String()
}

func (s *ListAppVersionResponse) SetHeaders(v map[string]*string) *ListAppVersionResponse {
	s.Headers = v
	return s
}

func (s *ListAppVersionResponse) SetStatusCode(v int32) *ListAppVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAppVersionResponse) SetBody(v *ListAppVersionResponseBody) *ListAppVersionResponse {
	s.Body = v
	return s
}

type ModifyAppRequest struct {
	AppId   *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
}

func (s ModifyAppRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyAppRequest) GoString() string {
	return s.String()
}

func (s *ModifyAppRequest) SetAppId(v string) *ModifyAppRequest {
	s.AppId = &v
	return s
}

func (s *ModifyAppRequest) SetAppName(v string) *ModifyAppRequest {
	s.AppName = &v
	return s
}

type ModifyAppResponseBody struct {
	AppId     *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpCode  *int32  `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyAppResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAppResponseBody) SetAppId(v string) *ModifyAppResponseBody {
	s.AppId = &v
	return s
}

func (s *ModifyAppResponseBody) SetCode(v string) *ModifyAppResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyAppResponseBody) SetHttpCode(v int32) *ModifyAppResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ModifyAppResponseBody) SetMessage(v string) *ModifyAppResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyAppResponseBody) SetRequestId(v string) *ModifyAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyAppResponseBody) SetSuccess(v bool) *ModifyAppResponseBody {
	s.Success = &v
	return s
}

type ModifyAppResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyAppResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyAppResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyAppResponse) GoString() string {
	return s.String()
}

func (s *ModifyAppResponse) SetHeaders(v map[string]*string) *ModifyAppResponse {
	s.Headers = v
	return s
}

func (s *ModifyAppResponse) SetStatusCode(v int32) *ModifyAppResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAppResponse) SetBody(v *ModifyAppResponseBody) *ModifyAppResponse {
	s.Body = v
	return s
}

type ModifyAppVersionRequest struct {
	AppVersionId   *string `json:"AppVersionId,omitempty" xml:"AppVersionId,omitempty"`
	AppVersionName *string `json:"AppVersionName,omitempty" xml:"AppVersionName,omitempty"`
}

func (s ModifyAppVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyAppVersionRequest) GoString() string {
	return s.String()
}

func (s *ModifyAppVersionRequest) SetAppVersionId(v string) *ModifyAppVersionRequest {
	s.AppVersionId = &v
	return s
}

func (s *ModifyAppVersionRequest) SetAppVersionName(v string) *ModifyAppVersionRequest {
	s.AppVersionName = &v
	return s
}

type ModifyAppVersionResponseBody struct {
	AppVersionId *string `json:"AppVersionId,omitempty" xml:"AppVersionId,omitempty"`
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpCode     *int32  `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	Message      *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyAppVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyAppVersionResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAppVersionResponseBody) SetAppVersionId(v string) *ModifyAppVersionResponseBody {
	s.AppVersionId = &v
	return s
}

func (s *ModifyAppVersionResponseBody) SetCode(v string) *ModifyAppVersionResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyAppVersionResponseBody) SetHttpCode(v int32) *ModifyAppVersionResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ModifyAppVersionResponseBody) SetMessage(v string) *ModifyAppVersionResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyAppVersionResponseBody) SetRequestId(v string) *ModifyAppVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyAppVersionResponseBody) SetSuccess(v bool) *ModifyAppVersionResponseBody {
	s.Success = &v
	return s
}

type ModifyAppVersionResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyAppVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyAppVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyAppVersionResponse) GoString() string {
	return s.String()
}

func (s *ModifyAppVersionResponse) SetHeaders(v map[string]*string) *ModifyAppVersionResponse {
	s.Headers = v
	return s
}

func (s *ModifyAppVersionResponse) SetStatusCode(v int32) *ModifyAppVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAppVersionResponse) SetBody(v *ModifyAppVersionResponseBody) *ModifyAppVersionResponse {
	s.Body = v
	return s
}

type ModifyProjectRequest struct {
	BoundAppIdList []*string `json:"BoundAppIdList,omitempty" xml:"BoundAppIdList,omitempty" type:"Repeated"`
	// 请求操作人Id
	OperatorId *string `json:"OperatorId,omitempty" xml:"OperatorId,omitempty"`
	// 请求操作人类型
	OperatorType *string `json:"OperatorType,omitempty" xml:"OperatorType,omitempty"`
	// project Id
	ProjectId   *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	ProjectMemo *string `json:"ProjectMemo,omitempty" xml:"ProjectMemo,omitempty"`
	// project name
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// key : districtId
	ProjectQuotaLimit *ModifyProjectRequestProjectQuotaLimit `json:"ProjectQuotaLimit,omitempty" xml:"ProjectQuotaLimit,omitempty" type:"Struct"`
}

func (s ModifyProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyProjectRequest) GoString() string {
	return s.String()
}

func (s *ModifyProjectRequest) SetBoundAppIdList(v []*string) *ModifyProjectRequest {
	s.BoundAppIdList = v
	return s
}

func (s *ModifyProjectRequest) SetOperatorId(v string) *ModifyProjectRequest {
	s.OperatorId = &v
	return s
}

func (s *ModifyProjectRequest) SetOperatorType(v string) *ModifyProjectRequest {
	s.OperatorType = &v
	return s
}

func (s *ModifyProjectRequest) SetProjectId(v string) *ModifyProjectRequest {
	s.ProjectId = &v
	return s
}

func (s *ModifyProjectRequest) SetProjectMemo(v string) *ModifyProjectRequest {
	s.ProjectMemo = &v
	return s
}

func (s *ModifyProjectRequest) SetProjectName(v string) *ModifyProjectRequest {
	s.ProjectName = &v
	return s
}

func (s *ModifyProjectRequest) SetProjectQuotaLimit(v *ModifyProjectRequestProjectQuotaLimit) *ModifyProjectRequest {
	s.ProjectQuotaLimit = v
	return s
}

type ModifyProjectRequestProjectQuotaLimit struct {
	// key - districtId
	DistrictLimitMap map[string]*ProjectQuotaLimitDistrictLimitMapValue `json:"DistrictLimitMap,omitempty" xml:"DistrictLimitMap,omitempty"`
	// 限制类型 ：目前默认 - ReserveContainer
	LimitType *string `json:"LimitType,omitempty" xml:"LimitType,omitempty"`
}

func (s ModifyProjectRequestProjectQuotaLimit) String() string {
	return tea.Prettify(s)
}

func (s ModifyProjectRequestProjectQuotaLimit) GoString() string {
	return s.String()
}

func (s *ModifyProjectRequestProjectQuotaLimit) SetDistrictLimitMap(v map[string]*ProjectQuotaLimitDistrictLimitMapValue) *ModifyProjectRequestProjectQuotaLimit {
	s.DistrictLimitMap = v
	return s
}

func (s *ModifyProjectRequestProjectQuotaLimit) SetLimitType(v string) *ModifyProjectRequestProjectQuotaLimit {
	s.LimitType = &v
	return s
}

type ModifyProjectShrinkRequest struct {
	BoundAppIdListShrink *string `json:"BoundAppIdList,omitempty" xml:"BoundAppIdList,omitempty"`
	// 请求操作人Id
	OperatorId *string `json:"OperatorId,omitempty" xml:"OperatorId,omitempty"`
	// 请求操作人类型
	OperatorType *string `json:"OperatorType,omitempty" xml:"OperatorType,omitempty"`
	// project Id
	ProjectId   *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	ProjectMemo *string `json:"ProjectMemo,omitempty" xml:"ProjectMemo,omitempty"`
	// project name
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// key : districtId
	ProjectQuotaLimitShrink *string `json:"ProjectQuotaLimit,omitempty" xml:"ProjectQuotaLimit,omitempty"`
}

func (s ModifyProjectShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyProjectShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyProjectShrinkRequest) SetBoundAppIdListShrink(v string) *ModifyProjectShrinkRequest {
	s.BoundAppIdListShrink = &v
	return s
}

func (s *ModifyProjectShrinkRequest) SetOperatorId(v string) *ModifyProjectShrinkRequest {
	s.OperatorId = &v
	return s
}

func (s *ModifyProjectShrinkRequest) SetOperatorType(v string) *ModifyProjectShrinkRequest {
	s.OperatorType = &v
	return s
}

func (s *ModifyProjectShrinkRequest) SetProjectId(v string) *ModifyProjectShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *ModifyProjectShrinkRequest) SetProjectMemo(v string) *ModifyProjectShrinkRequest {
	s.ProjectMemo = &v
	return s
}

func (s *ModifyProjectShrinkRequest) SetProjectName(v string) *ModifyProjectShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *ModifyProjectShrinkRequest) SetProjectQuotaLimitShrink(v string) *ModifyProjectShrinkRequest {
	s.ProjectQuotaLimitShrink = &v
	return s
}

type ModifyProjectResponseBody struct {
	// 业务处理结果Code
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 业务对象
	Data *ModifyProjectResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// 业务处理消息摘要
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 操作请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 业务处理是否成功
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyProjectResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyProjectResponseBody) SetCode(v string) *ModifyProjectResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyProjectResponseBody) SetData(v *ModifyProjectResponseBodyData) *ModifyProjectResponseBody {
	s.Data = v
	return s
}

func (s *ModifyProjectResponseBody) SetMessage(v string) *ModifyProjectResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyProjectResponseBody) SetRequestId(v string) *ModifyProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyProjectResponseBody) SetSuccess(v bool) *ModifyProjectResponseBody {
	s.Success = &v
	return s
}

type ModifyProjectResponseBodyData struct {
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyProjectResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ModifyProjectResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModifyProjectResponseBodyData) SetSuccess(v bool) *ModifyProjectResponseBodyData {
	s.Success = &v
	return s
}

type ModifyProjectResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyProjectResponse) GoString() string {
	return s.String()
}

func (s *ModifyProjectResponse) SetHeaders(v map[string]*string) *ModifyProjectResponse {
	s.Headers = v
	return s
}

func (s *ModifyProjectResponse) SetStatusCode(v int32) *ModifyProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyProjectResponse) SetBody(v *ModifyProjectResponseBody) *ModifyProjectResponse {
	s.Body = v
	return s
}

type PageQueryProjectRequest struct {
	// projectId or projectName like
	KeySearch *string `json:"KeySearch,omitempty" xml:"KeySearch,omitempty"`
	// 请求操作人Id
	OperatorId *string `json:"OperatorId,omitempty" xml:"OperatorId,omitempty"`
	// 请求操作人类型
	OperatorType *string `json:"OperatorType,omitempty" xml:"OperatorType,omitempty"`
	// 当前页码，默认1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 每页项数，默认20,最大100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s PageQueryProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s PageQueryProjectRequest) GoString() string {
	return s.String()
}

func (s *PageQueryProjectRequest) SetKeySearch(v string) *PageQueryProjectRequest {
	s.KeySearch = &v
	return s
}

func (s *PageQueryProjectRequest) SetOperatorId(v string) *PageQueryProjectRequest {
	s.OperatorId = &v
	return s
}

func (s *PageQueryProjectRequest) SetOperatorType(v string) *PageQueryProjectRequest {
	s.OperatorType = &v
	return s
}

func (s *PageQueryProjectRequest) SetPageNumber(v int32) *PageQueryProjectRequest {
	s.PageNumber = &v
	return s
}

func (s *PageQueryProjectRequest) SetPageSize(v int32) *PageQueryProjectRequest {
	s.PageSize = &v
	return s
}

type PageQueryProjectResponseBody struct {
	// 业务处理结果Code
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 业务对象
	Data *PageQueryProjectResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// 业务处理消息摘要
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 操作请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 业务处理是否成功
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PageQueryProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PageQueryProjectResponseBody) GoString() string {
	return s.String()
}

func (s *PageQueryProjectResponseBody) SetCode(v string) *PageQueryProjectResponseBody {
	s.Code = &v
	return s
}

func (s *PageQueryProjectResponseBody) SetData(v *PageQueryProjectResponseBodyData) *PageQueryProjectResponseBody {
	s.Data = v
	return s
}

func (s *PageQueryProjectResponseBody) SetMessage(v string) *PageQueryProjectResponseBody {
	s.Message = &v
	return s
}

func (s *PageQueryProjectResponseBody) SetRequestId(v string) *PageQueryProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *PageQueryProjectResponseBody) SetSuccess(v bool) *PageQueryProjectResponseBody {
	s.Success = &v
	return s
}

type PageQueryProjectResponseBodyData struct {
	// 当前页码，默认1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 每页项数，默认20,最大100
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 总页数
	Pages *int64 `json:"Pages,omitempty" xml:"Pages,omitempty"`
	// 结果集
	Records []*PageQueryProjectResponseBodyDataRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	// 总共项数
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s PageQueryProjectResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s PageQueryProjectResponseBodyData) GoString() string {
	return s.String()
}

func (s *PageQueryProjectResponseBodyData) SetPageNumber(v int64) *PageQueryProjectResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *PageQueryProjectResponseBodyData) SetPageSize(v int64) *PageQueryProjectResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *PageQueryProjectResponseBodyData) SetPages(v int64) *PageQueryProjectResponseBodyData {
	s.Pages = &v
	return s
}

func (s *PageQueryProjectResponseBodyData) SetRecords(v []*PageQueryProjectResponseBodyDataRecords) *PageQueryProjectResponseBodyData {
	s.Records = v
	return s
}

func (s *PageQueryProjectResponseBodyData) SetTotalCount(v int64) *PageQueryProjectResponseBodyData {
	s.TotalCount = &v
	return s
}

type PageQueryProjectResponseBodyDataRecords struct {
	// 项目关联的应用数量
	BoundAppNums *int64  `json:"BoundAppNums,omitempty" xml:"BoundAppNums,omitempty"`
	GmtCreate    *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified  *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	ProjectId    *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	ProjectMemo  *string `json:"ProjectMemo,omitempty" xml:"ProjectMemo,omitempty"`
	ProjectName  *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// key : districtId
	ProjectQuotaLimit *PageQueryProjectResponseBodyDataRecordsProjectQuotaLimit `json:"ProjectQuotaLimit,omitempty" xml:"ProjectQuotaLimit,omitempty" type:"Struct"`
}

func (s PageQueryProjectResponseBodyDataRecords) String() string {
	return tea.Prettify(s)
}

func (s PageQueryProjectResponseBodyDataRecords) GoString() string {
	return s.String()
}

func (s *PageQueryProjectResponseBodyDataRecords) SetBoundAppNums(v int64) *PageQueryProjectResponseBodyDataRecords {
	s.BoundAppNums = &v
	return s
}

func (s *PageQueryProjectResponseBodyDataRecords) SetGmtCreate(v string) *PageQueryProjectResponseBodyDataRecords {
	s.GmtCreate = &v
	return s
}

func (s *PageQueryProjectResponseBodyDataRecords) SetGmtModified(v string) *PageQueryProjectResponseBodyDataRecords {
	s.GmtModified = &v
	return s
}

func (s *PageQueryProjectResponseBodyDataRecords) SetProjectId(v string) *PageQueryProjectResponseBodyDataRecords {
	s.ProjectId = &v
	return s
}

func (s *PageQueryProjectResponseBodyDataRecords) SetProjectMemo(v string) *PageQueryProjectResponseBodyDataRecords {
	s.ProjectMemo = &v
	return s
}

func (s *PageQueryProjectResponseBodyDataRecords) SetProjectName(v string) *PageQueryProjectResponseBodyDataRecords {
	s.ProjectName = &v
	return s
}

func (s *PageQueryProjectResponseBodyDataRecords) SetProjectQuotaLimit(v *PageQueryProjectResponseBodyDataRecordsProjectQuotaLimit) *PageQueryProjectResponseBodyDataRecords {
	s.ProjectQuotaLimit = v
	return s
}

type PageQueryProjectResponseBodyDataRecordsProjectQuotaLimit struct {
	// key - districtId
	DistrictLimitMap map[string]*DataRecordsProjectQuotaLimitDistrictLimitMapValue `json:"DistrictLimitMap,omitempty" xml:"DistrictLimitMap,omitempty"`
	// 限制类型 ：目前默认 - ReserveContainer
	LimitType *string `json:"LimitType,omitempty" xml:"LimitType,omitempty"`
}

func (s PageQueryProjectResponseBodyDataRecordsProjectQuotaLimit) String() string {
	return tea.Prettify(s)
}

func (s PageQueryProjectResponseBodyDataRecordsProjectQuotaLimit) GoString() string {
	return s.String()
}

func (s *PageQueryProjectResponseBodyDataRecordsProjectQuotaLimit) SetDistrictLimitMap(v map[string]*DataRecordsProjectQuotaLimitDistrictLimitMapValue) *PageQueryProjectResponseBodyDataRecordsProjectQuotaLimit {
	s.DistrictLimitMap = v
	return s
}

func (s *PageQueryProjectResponseBodyDataRecordsProjectQuotaLimit) SetLimitType(v string) *PageQueryProjectResponseBodyDataRecordsProjectQuotaLimit {
	s.LimitType = &v
	return s
}

type PageQueryProjectResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PageQueryProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PageQueryProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s PageQueryProjectResponse) GoString() string {
	return s.String()
}

func (s *PageQueryProjectResponse) SetHeaders(v map[string]*string) *PageQueryProjectResponse {
	s.Headers = v
	return s
}

func (s *PageQueryProjectResponse) SetStatusCode(v int32) *PageQueryProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *PageQueryProjectResponse) SetBody(v *PageQueryProjectResponseBody) *PageQueryProjectResponse {
	s.Body = v
	return s
}

type PageQueryProjectAppsRequest struct {
	// 请求操作人Id
	OperatorId *string `json:"OperatorId,omitempty" xml:"OperatorId,omitempty"`
	// 请求操作人类型
	OperatorType *string `json:"OperatorType,omitempty" xml:"OperatorType,omitempty"`
	// 当前页码，默认1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 每页项数，默认20,最大100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// projectId
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s PageQueryProjectAppsRequest) String() string {
	return tea.Prettify(s)
}

func (s PageQueryProjectAppsRequest) GoString() string {
	return s.String()
}

func (s *PageQueryProjectAppsRequest) SetOperatorId(v string) *PageQueryProjectAppsRequest {
	s.OperatorId = &v
	return s
}

func (s *PageQueryProjectAppsRequest) SetOperatorType(v string) *PageQueryProjectAppsRequest {
	s.OperatorType = &v
	return s
}

func (s *PageQueryProjectAppsRequest) SetPageNumber(v int32) *PageQueryProjectAppsRequest {
	s.PageNumber = &v
	return s
}

func (s *PageQueryProjectAppsRequest) SetPageSize(v int32) *PageQueryProjectAppsRequest {
	s.PageSize = &v
	return s
}

func (s *PageQueryProjectAppsRequest) SetProjectId(v string) *PageQueryProjectAppsRequest {
	s.ProjectId = &v
	return s
}

type PageQueryProjectAppsResponseBody struct {
	// 业务处理结果Code
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 业务对象
	Data *PageQueryProjectAppsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// 业务处理消息摘要
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 操作请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 业务处理是否成功
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PageQueryProjectAppsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PageQueryProjectAppsResponseBody) GoString() string {
	return s.String()
}

func (s *PageQueryProjectAppsResponseBody) SetCode(v string) *PageQueryProjectAppsResponseBody {
	s.Code = &v
	return s
}

func (s *PageQueryProjectAppsResponseBody) SetData(v *PageQueryProjectAppsResponseBodyData) *PageQueryProjectAppsResponseBody {
	s.Data = v
	return s
}

func (s *PageQueryProjectAppsResponseBody) SetMessage(v string) *PageQueryProjectAppsResponseBody {
	s.Message = &v
	return s
}

func (s *PageQueryProjectAppsResponseBody) SetRequestId(v string) *PageQueryProjectAppsResponseBody {
	s.RequestId = &v
	return s
}

func (s *PageQueryProjectAppsResponseBody) SetSuccess(v bool) *PageQueryProjectAppsResponseBody {
	s.Success = &v
	return s
}

type PageQueryProjectAppsResponseBodyData struct {
	// 当前页码，默认1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 每页项数，默认20,最大100
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 总页数
	Pages *int64 `json:"Pages,omitempty" xml:"Pages,omitempty"`
	// 结果集
	Records []*PageQueryProjectAppsResponseBodyDataRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	// 总共项数
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s PageQueryProjectAppsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s PageQueryProjectAppsResponseBodyData) GoString() string {
	return s.String()
}

func (s *PageQueryProjectAppsResponseBodyData) SetPageNumber(v int64) *PageQueryProjectAppsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *PageQueryProjectAppsResponseBodyData) SetPageSize(v int64) *PageQueryProjectAppsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *PageQueryProjectAppsResponseBodyData) SetPages(v int64) *PageQueryProjectAppsResponseBodyData {
	s.Pages = &v
	return s
}

func (s *PageQueryProjectAppsResponseBodyData) SetRecords(v []*PageQueryProjectAppsResponseBodyDataRecords) *PageQueryProjectAppsResponseBodyData {
	s.Records = v
	return s
}

func (s *PageQueryProjectAppsResponseBodyData) SetTotalCount(v int64) *PageQueryProjectAppsResponseBodyData {
	s.TotalCount = &v
	return s
}

type PageQueryProjectAppsResponseBodyDataRecords struct {
	AppId     *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppName   *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s PageQueryProjectAppsResponseBodyDataRecords) String() string {
	return tea.Prettify(s)
}

func (s PageQueryProjectAppsResponseBodyDataRecords) GoString() string {
	return s.String()
}

func (s *PageQueryProjectAppsResponseBodyDataRecords) SetAppId(v string) *PageQueryProjectAppsResponseBodyDataRecords {
	s.AppId = &v
	return s
}

func (s *PageQueryProjectAppsResponseBodyDataRecords) SetAppName(v string) *PageQueryProjectAppsResponseBodyDataRecords {
	s.AppName = &v
	return s
}

func (s *PageQueryProjectAppsResponseBodyDataRecords) SetGmtCreate(v string) *PageQueryProjectAppsResponseBodyDataRecords {
	s.GmtCreate = &v
	return s
}

func (s *PageQueryProjectAppsResponseBodyDataRecords) SetProjectId(v string) *PageQueryProjectAppsResponseBodyDataRecords {
	s.ProjectId = &v
	return s
}

type PageQueryProjectAppsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PageQueryProjectAppsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PageQueryProjectAppsResponse) String() string {
	return tea.Prettify(s)
}

func (s PageQueryProjectAppsResponse) GoString() string {
	return s.String()
}

func (s *PageQueryProjectAppsResponse) SetHeaders(v map[string]*string) *PageQueryProjectAppsResponse {
	s.Headers = v
	return s
}

func (s *PageQueryProjectAppsResponse) SetStatusCode(v int32) *PageQueryProjectAppsResponse {
	s.StatusCode = &v
	return s
}

func (s *PageQueryProjectAppsResponse) SetBody(v *PageQueryProjectAppsResponseBody) *PageQueryProjectAppsResponse {
	s.Body = v
	return s
}

type QueryOfflineTaskProgressRequest struct {
	AppId     *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s QueryOfflineTaskProgressRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryOfflineTaskProgressRequest) GoString() string {
	return s.String()
}

func (s *QueryOfflineTaskProgressRequest) SetAppId(v string) *QueryOfflineTaskProgressRequest {
	s.AppId = &v
	return s
}

func (s *QueryOfflineTaskProgressRequest) SetVersionId(v string) *QueryOfflineTaskProgressRequest {
	s.VersionId = &v
	return s
}

type QueryOfflineTaskProgressResponseBody struct {
	Code      *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *QueryOfflineTaskProgressResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryOfflineTaskProgressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryOfflineTaskProgressResponseBody) GoString() string {
	return s.String()
}

func (s *QueryOfflineTaskProgressResponseBody) SetCode(v string) *QueryOfflineTaskProgressResponseBody {
	s.Code = &v
	return s
}

func (s *QueryOfflineTaskProgressResponseBody) SetData(v *QueryOfflineTaskProgressResponseBodyData) *QueryOfflineTaskProgressResponseBody {
	s.Data = v
	return s
}

func (s *QueryOfflineTaskProgressResponseBody) SetMessage(v string) *QueryOfflineTaskProgressResponseBody {
	s.Message = &v
	return s
}

func (s *QueryOfflineTaskProgressResponseBody) SetRequestId(v string) *QueryOfflineTaskProgressResponseBody {
	s.RequestId = &v
	return s
}

type QueryOfflineTaskProgressResponseBodyData struct {
	ErrorCode    *string  `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string  `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Progress     *float64 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	Status       *string  `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s QueryOfflineTaskProgressResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryOfflineTaskProgressResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryOfflineTaskProgressResponseBodyData) SetErrorCode(v string) *QueryOfflineTaskProgressResponseBodyData {
	s.ErrorCode = &v
	return s
}

func (s *QueryOfflineTaskProgressResponseBodyData) SetErrorMessage(v string) *QueryOfflineTaskProgressResponseBodyData {
	s.ErrorMessage = &v
	return s
}

func (s *QueryOfflineTaskProgressResponseBodyData) SetProgress(v float64) *QueryOfflineTaskProgressResponseBodyData {
	s.Progress = &v
	return s
}

func (s *QueryOfflineTaskProgressResponseBodyData) SetStatus(v string) *QueryOfflineTaskProgressResponseBodyData {
	s.Status = &v
	return s
}

type QueryOfflineTaskProgressResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryOfflineTaskProgressResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryOfflineTaskProgressResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryOfflineTaskProgressResponse) GoString() string {
	return s.String()
}

func (s *QueryOfflineTaskProgressResponse) SetHeaders(v map[string]*string) *QueryOfflineTaskProgressResponse {
	s.Headers = v
	return s
}

func (s *QueryOfflineTaskProgressResponse) SetStatusCode(v int32) *QueryOfflineTaskProgressResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryOfflineTaskProgressResponse) SetBody(v *QueryOfflineTaskProgressResponseBody) *QueryOfflineTaskProgressResponse {
	s.Body = v
	return s
}

type RefreshDistrictMetaResponseBody struct {
	// 业务处理结果Code
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 业务对象
	Data *RefreshDistrictMetaResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// 业务处理消息摘要
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 操作请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 业务处理是否成功
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RefreshDistrictMetaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RefreshDistrictMetaResponseBody) GoString() string {
	return s.String()
}

func (s *RefreshDistrictMetaResponseBody) SetCode(v string) *RefreshDistrictMetaResponseBody {
	s.Code = &v
	return s
}

func (s *RefreshDistrictMetaResponseBody) SetData(v *RefreshDistrictMetaResponseBodyData) *RefreshDistrictMetaResponseBody {
	s.Data = v
	return s
}

func (s *RefreshDistrictMetaResponseBody) SetMessage(v string) *RefreshDistrictMetaResponseBody {
	s.Message = &v
	return s
}

func (s *RefreshDistrictMetaResponseBody) SetRequestId(v string) *RefreshDistrictMetaResponseBody {
	s.RequestId = &v
	return s
}

func (s *RefreshDistrictMetaResponseBody) SetSuccess(v bool) *RefreshDistrictMetaResponseBody {
	s.Success = &v
	return s
}

type RefreshDistrictMetaResponseBodyData struct {
	// 业务处理结果Code
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 业务对象
	Data map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// 业务处理消息摘要
	Message           *string                                               `json:"Message,omitempty" xml:"Message,omitempty"`
	ProjectQuotaLimit *RefreshDistrictMetaResponseBodyDataProjectQuotaLimit `json:"ProjectQuotaLimit,omitempty" xml:"ProjectQuotaLimit,omitempty" type:"Struct"`
	// 操作请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 业务处理是否成功
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RefreshDistrictMetaResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RefreshDistrictMetaResponseBodyData) GoString() string {
	return s.String()
}

func (s *RefreshDistrictMetaResponseBodyData) SetCode(v string) *RefreshDistrictMetaResponseBodyData {
	s.Code = &v
	return s
}

func (s *RefreshDistrictMetaResponseBodyData) SetData(v map[string]interface{}) *RefreshDistrictMetaResponseBodyData {
	s.Data = v
	return s
}

func (s *RefreshDistrictMetaResponseBodyData) SetMessage(v string) *RefreshDistrictMetaResponseBodyData {
	s.Message = &v
	return s
}

func (s *RefreshDistrictMetaResponseBodyData) SetProjectQuotaLimit(v *RefreshDistrictMetaResponseBodyDataProjectQuotaLimit) *RefreshDistrictMetaResponseBodyData {
	s.ProjectQuotaLimit = v
	return s
}

func (s *RefreshDistrictMetaResponseBodyData) SetRequestId(v string) *RefreshDistrictMetaResponseBodyData {
	s.RequestId = &v
	return s
}

func (s *RefreshDistrictMetaResponseBodyData) SetSuccess(v bool) *RefreshDistrictMetaResponseBodyData {
	s.Success = &v
	return s
}

type RefreshDistrictMetaResponseBodyDataProjectQuotaLimit struct {
	// key - districtId
	DistrictLimitMap map[string]*DataProjectQuotaLimitDistrictLimitMapValue `json:"DistrictLimitMap,omitempty" xml:"DistrictLimitMap,omitempty"`
	// 限制类型 ：目前默认 - ReserveContainer
	LimitType *string `json:"LimitType,omitempty" xml:"LimitType,omitempty"`
}

func (s RefreshDistrictMetaResponseBodyDataProjectQuotaLimit) String() string {
	return tea.Prettify(s)
}

func (s RefreshDistrictMetaResponseBodyDataProjectQuotaLimit) GoString() string {
	return s.String()
}

func (s *RefreshDistrictMetaResponseBodyDataProjectQuotaLimit) SetDistrictLimitMap(v map[string]*DataProjectQuotaLimitDistrictLimitMapValue) *RefreshDistrictMetaResponseBodyDataProjectQuotaLimit {
	s.DistrictLimitMap = v
	return s
}

func (s *RefreshDistrictMetaResponseBodyDataProjectQuotaLimit) SetLimitType(v string) *RefreshDistrictMetaResponseBodyDataProjectQuotaLimit {
	s.LimitType = &v
	return s
}

type RefreshDistrictMetaResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RefreshDistrictMetaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RefreshDistrictMetaResponse) String() string {
	return tea.Prettify(s)
}

func (s RefreshDistrictMetaResponse) GoString() string {
	return s.String()
}

func (s *RefreshDistrictMetaResponse) SetHeaders(v map[string]*string) *RefreshDistrictMetaResponse {
	s.Headers = v
	return s
}

func (s *RefreshDistrictMetaResponse) SetStatusCode(v int32) *RefreshDistrictMetaResponse {
	s.StatusCode = &v
	return s
}

func (s *RefreshDistrictMetaResponse) SetBody(v *RefreshDistrictMetaResponseBody) *RefreshDistrictMetaResponse {
	s.Body = v
	return s
}

type StopAppSessionRequest struct {
	// 自定义会话id
	CustomSessionId *string `json:"CustomSessionId,omitempty" xml:"CustomSessionId,omitempty"`
	// 自定义用户id
	PlatformSessionId *string                           `json:"PlatformSessionId,omitempty" xml:"PlatformSessionId,omitempty"`
	StopParam         []*StopAppSessionRequestStopParam `json:"StopParam,omitempty" xml:"StopParam,omitempty" type:"Repeated"`
}

func (s StopAppSessionRequest) String() string {
	return tea.Prettify(s)
}

func (s StopAppSessionRequest) GoString() string {
	return s.String()
}

func (s *StopAppSessionRequest) SetCustomSessionId(v string) *StopAppSessionRequest {
	s.CustomSessionId = &v
	return s
}

func (s *StopAppSessionRequest) SetPlatformSessionId(v string) *StopAppSessionRequest {
	s.PlatformSessionId = &v
	return s
}

func (s *StopAppSessionRequest) SetStopParam(v []*StopAppSessionRequestStopParam) *StopAppSessionRequest {
	s.StopParam = v
	return s
}

type StopAppSessionRequestStopParam struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s StopAppSessionRequestStopParam) String() string {
	return tea.Prettify(s)
}

func (s StopAppSessionRequestStopParam) GoString() string {
	return s.String()
}

func (s *StopAppSessionRequestStopParam) SetKey(v string) *StopAppSessionRequestStopParam {
	s.Key = &v
	return s
}

func (s *StopAppSessionRequestStopParam) SetValue(v string) *StopAppSessionRequestStopParam {
	s.Value = &v
	return s
}

type StopAppSessionResponseBody struct {
	// 应用id
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 应用版本
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	// 自定义会话id
	CustomSessionId *string `json:"CustomSessionId,omitempty" xml:"CustomSessionId,omitempty"`
	// 平台会话id
	PlatformSessionId *string `json:"PlatformSessionId,omitempty" xml:"PlatformSessionId,omitempty"`
	// 请求id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopAppSessionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopAppSessionResponseBody) GoString() string {
	return s.String()
}

func (s *StopAppSessionResponseBody) SetAppId(v string) *StopAppSessionResponseBody {
	s.AppId = &v
	return s
}

func (s *StopAppSessionResponseBody) SetAppVersion(v string) *StopAppSessionResponseBody {
	s.AppVersion = &v
	return s
}

func (s *StopAppSessionResponseBody) SetCustomSessionId(v string) *StopAppSessionResponseBody {
	s.CustomSessionId = &v
	return s
}

func (s *StopAppSessionResponseBody) SetPlatformSessionId(v string) *StopAppSessionResponseBody {
	s.PlatformSessionId = &v
	return s
}

func (s *StopAppSessionResponseBody) SetRequestId(v string) *StopAppSessionResponseBody {
	s.RequestId = &v
	return s
}

type StopAppSessionResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StopAppSessionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopAppSessionResponse) String() string {
	return tea.Prettify(s)
}

func (s StopAppSessionResponse) GoString() string {
	return s.String()
}

func (s *StopAppSessionResponse) SetHeaders(v map[string]*string) *StopAppSessionResponse {
	s.Headers = v
	return s
}

func (s *StopAppSessionResponse) SetStatusCode(v int32) *StopAppSessionResponse {
	s.StatusCode = &v
	return s
}

func (s *StopAppSessionResponse) SetBody(v *StopAppSessionResponseBody) *StopAppSessionResponse {
	s.Body = v
	return s
}

type SubmitOfflineTaskRequest struct {
	AppId       *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppName     *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	AppType     *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	Env         *string `json:"Env,omitempty" xml:"Env,omitempty"`
	Uri         *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
	VersionId   *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s SubmitOfflineTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitOfflineTaskRequest) GoString() string {
	return s.String()
}

func (s *SubmitOfflineTaskRequest) SetAppId(v string) *SubmitOfflineTaskRequest {
	s.AppId = &v
	return s
}

func (s *SubmitOfflineTaskRequest) SetAppName(v string) *SubmitOfflineTaskRequest {
	s.AppName = &v
	return s
}

func (s *SubmitOfflineTaskRequest) SetAppType(v string) *SubmitOfflineTaskRequest {
	s.AppType = &v
	return s
}

func (s *SubmitOfflineTaskRequest) SetEnv(v string) *SubmitOfflineTaskRequest {
	s.Env = &v
	return s
}

func (s *SubmitOfflineTaskRequest) SetUri(v string) *SubmitOfflineTaskRequest {
	s.Uri = &v
	return s
}

func (s *SubmitOfflineTaskRequest) SetVersionId(v string) *SubmitOfflineTaskRequest {
	s.VersionId = &v
	return s
}

func (s *SubmitOfflineTaskRequest) SetVersionName(v string) *SubmitOfflineTaskRequest {
	s.VersionName = &v
	return s
}

type SubmitOfflineTaskResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitOfflineTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitOfflineTaskResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitOfflineTaskResponseBody) SetCode(v string) *SubmitOfflineTaskResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitOfflineTaskResponseBody) SetData(v bool) *SubmitOfflineTaskResponseBody {
	s.Data = &v
	return s
}

func (s *SubmitOfflineTaskResponseBody) SetMessage(v string) *SubmitOfflineTaskResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitOfflineTaskResponseBody) SetRequestId(v string) *SubmitOfflineTaskResponseBody {
	s.RequestId = &v
	return s
}

type SubmitOfflineTaskResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SubmitOfflineTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubmitOfflineTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitOfflineTaskResponse) GoString() string {
	return s.String()
}

func (s *SubmitOfflineTaskResponse) SetHeaders(v map[string]*string) *SubmitOfflineTaskResponse {
	s.Headers = v
	return s
}

func (s *SubmitOfflineTaskResponse) SetStatusCode(v int32) *SubmitOfflineTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitOfflineTaskResponse) SetBody(v *SubmitOfflineTaskResponseBody) *SubmitOfflineTaskResponse {
	s.Body = v
	return s
}

type ProjectQuotaLimitDistrictLimitMapValue struct {
	// 大区ID
	DistrictId *string `json:"DistrictId,omitempty" xml:"DistrictId,omitempty"`
	// 大区名称
	DistrictName *string `json:"DistrictName,omitempty" xml:"DistrictName,omitempty"`
	// 上限
	MaxLimit *int64 `json:"MaxLimit,omitempty" xml:"MaxLimit,omitempty"`
}

func (s ProjectQuotaLimitDistrictLimitMapValue) String() string {
	return tea.Prettify(s)
}

func (s ProjectQuotaLimitDistrictLimitMapValue) GoString() string {
	return s.String()
}

func (s *ProjectQuotaLimitDistrictLimitMapValue) SetDistrictId(v string) *ProjectQuotaLimitDistrictLimitMapValue {
	s.DistrictId = &v
	return s
}

func (s *ProjectQuotaLimitDistrictLimitMapValue) SetDistrictName(v string) *ProjectQuotaLimitDistrictLimitMapValue {
	s.DistrictName = &v
	return s
}

func (s *ProjectQuotaLimitDistrictLimitMapValue) SetMaxLimit(v int64) *ProjectQuotaLimitDistrictLimitMapValue {
	s.MaxLimit = &v
	return s
}

type DataProjectQuotaLimitDistrictLimitMapValue struct {
	// 大区ID
	DistrictId *string `json:"DistrictId,omitempty" xml:"DistrictId,omitempty"`
	// 大区名称
	DistrictName *string `json:"DistrictName,omitempty" xml:"DistrictName,omitempty"`
	// 上限
	MaxLimit *int64 `json:"MaxLimit,omitempty" xml:"MaxLimit,omitempty"`
}

func (s DataProjectQuotaLimitDistrictLimitMapValue) String() string {
	return tea.Prettify(s)
}

func (s DataProjectQuotaLimitDistrictLimitMapValue) GoString() string {
	return s.String()
}

func (s *DataProjectQuotaLimitDistrictLimitMapValue) SetDistrictId(v string) *DataProjectQuotaLimitDistrictLimitMapValue {
	s.DistrictId = &v
	return s
}

func (s *DataProjectQuotaLimitDistrictLimitMapValue) SetDistrictName(v string) *DataProjectQuotaLimitDistrictLimitMapValue {
	s.DistrictName = &v
	return s
}

func (s *DataProjectQuotaLimitDistrictLimitMapValue) SetMaxLimit(v int64) *DataProjectQuotaLimitDistrictLimitMapValue {
	s.MaxLimit = &v
	return s
}

type DataRecordsProjectQuotaLimitDistrictLimitMapValue struct {
	// 大区ID
	DistrictId *string `json:"DistrictId,omitempty" xml:"DistrictId,omitempty"`
	// 大区名称
	DistrictName *string `json:"DistrictName,omitempty" xml:"DistrictName,omitempty"`
	// 上限
	MaxLimit *int64 `json:"MaxLimit,omitempty" xml:"MaxLimit,omitempty"`
}

func (s DataRecordsProjectQuotaLimitDistrictLimitMapValue) String() string {
	return tea.Prettify(s)
}

func (s DataRecordsProjectQuotaLimitDistrictLimitMapValue) GoString() string {
	return s.String()
}

func (s *DataRecordsProjectQuotaLimitDistrictLimitMapValue) SetDistrictId(v string) *DataRecordsProjectQuotaLimitDistrictLimitMapValue {
	s.DistrictId = &v
	return s
}

func (s *DataRecordsProjectQuotaLimitDistrictLimitMapValue) SetDistrictName(v string) *DataRecordsProjectQuotaLimitDistrictLimitMapValue {
	s.DistrictName = &v
	return s
}

func (s *DataRecordsProjectQuotaLimitDistrictLimitMapValue) SetMaxLimit(v int64) *DataRecordsProjectQuotaLimitDistrictLimitMapValue {
	s.MaxLimit = &v
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("cgcs"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) CreateAdaptationWithOptions(tmpReq *CreateAdaptationRequest, runtime *util.RuntimeOptions) (_result *CreateAdaptationResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateAdaptationShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.AdaptTarget))) {
		request.AdaptTargetShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.AdaptTarget), tea.String("AdaptTarget"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AdaptTargetShrink)) {
		body["AdaptTarget"] = request.AdaptTargetShrink
	}

	if !tea.BoolValue(util.IsUnset(request.AppVersionId)) {
		body["AppVersionId"] = request.AppVersionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAdaptation"),
		Version:     tea.String("2021-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAdaptationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateAdaptation(request *CreateAdaptationRequest) (_result *CreateAdaptationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAdaptationResponse{}
	_body, _err := client.CreateAdaptationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateAppWithOptions(request *CreateAppRequest, runtime *util.RuntimeOptions) (_result *CreateAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		body["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.AppType)) {
		body["AppType"] = request.AppType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateApp"),
		Version:     tea.String("2021-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateApp(request *CreateAppRequest) (_result *CreateAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAppResponse{}
	_body, _err := client.CreateAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateAppSessionWithOptions(tmpReq *CreateAppSessionRequest, runtime *util.RuntimeOptions) (_result *CreateAppSessionResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateAppSessionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.ResultStore))) {
		request.ResultStoreShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.ResultStore), tea.String("ResultStore"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.StartParametersV2)) {
		request.StartParametersV2Shrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.StartParametersV2, tea.String("StartParametersV2"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.AppVersion)) {
		query["AppVersion"] = request.AppVersion
	}

	if !tea.BoolValue(util.IsUnset(request.ClientIp)) {
		query["ClientIp"] = request.ClientIp
	}

	if !tea.BoolValue(util.IsUnset(request.CustomSessionId)) {
		query["CustomSessionId"] = request.CustomSessionId
	}

	if !tea.BoolValue(util.IsUnset(request.CustomUserId)) {
		query["CustomUserId"] = request.CustomUserId
	}

	if !tea.BoolValue(util.IsUnset(request.DatasetId)) {
		query["DatasetId"] = request.DatasetId
	}

	if !tea.BoolValue(util.IsUnset(request.EnablePostpaid)) {
		query["EnablePostpaid"] = request.EnablePostpaid
	}

	if !tea.BoolValue(util.IsUnset(request.ResultStoreShrink)) {
		query["ResultStore"] = request.ResultStoreShrink
	}

	if !tea.BoolValue(util.IsUnset(request.StartParameters)) {
		query["StartParameters"] = request.StartParameters
	}

	if !tea.BoolValue(util.IsUnset(request.StartParametersV2Shrink)) {
		query["StartParametersV2"] = request.StartParametersV2Shrink
	}

	if !tea.BoolValue(util.IsUnset(request.SystemInfo)) {
		query["SystemInfo"] = request.SystemInfo
	}

	if !tea.BoolValue(util.IsUnset(request.Timeout)) {
		query["Timeout"] = request.Timeout
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAppSession"),
		Version:     tea.String("2021-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAppSessionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateAppSession(request *CreateAppSessionRequest) (_result *CreateAppSessionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAppSessionResponse{}
	_body, _err := client.CreateAppSessionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateAppSessionSyncWithOptions(request *CreateAppSessionSyncRequest, runtime *util.RuntimeOptions) (_result *CreateAppSessionSyncResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.AppVersion)) {
		query["AppVersion"] = request.AppVersion
	}

	if !tea.BoolValue(util.IsUnset(request.ClientIp)) {
		query["ClientIp"] = request.ClientIp
	}

	if !tea.BoolValue(util.IsUnset(request.CustomSessionId)) {
		query["CustomSessionId"] = request.CustomSessionId
	}

	if !tea.BoolValue(util.IsUnset(request.CustomUserId)) {
		query["CustomUserId"] = request.CustomUserId
	}

	if !tea.BoolValue(util.IsUnset(request.DistrictId)) {
		query["DistrictId"] = request.DistrictId
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.SpeedInfo)) {
		query["SpeedInfo"] = request.SpeedInfo
	}

	if !tea.BoolValue(util.IsUnset(request.StartParameters)) {
		query["StartParameters"] = request.StartParameters
	}

	if !tea.BoolValue(util.IsUnset(request.SystemInfo)) {
		query["SystemInfo"] = request.SystemInfo
	}

	if !tea.BoolValue(util.IsUnset(request.Tags)) {
		query["tags"] = request.Tags
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAppSessionSync"),
		Version:     tea.String("2021-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAppSessionSyncResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateAppSessionSync(request *CreateAppSessionSyncRequest) (_result *CreateAppSessionSyncResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAppSessionSyncResponse{}
	_body, _err := client.CreateAppSessionSyncWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateAppVersionWithOptions(request *CreateAppVersionRequest, runtime *util.RuntimeOptions) (_result *CreateAppVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.AppVersionName)) {
		body["AppVersionName"] = request.AppVersionName
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAppVersion"),
		Version:     tea.String("2021-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAppVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateAppVersion(request *CreateAppVersionRequest) (_result *CreateAppVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAppVersionResponse{}
	_body, _err := client.CreateAppVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateCapacityReservationWithOptions(request *CreateCapacityReservationRequest, runtime *util.RuntimeOptions) (_result *CreateCapacityReservationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.AppVersion)) {
		query["AppVersion"] = request.AppVersion
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DistrictId)) {
		query["DistrictId"] = request.DistrictId
	}

	if !tea.BoolValue(util.IsUnset(request.ExpectResourceReadyTime)) {
		query["ExpectResourceReadyTime"] = request.ExpectResourceReadyTime
	}

	if !tea.BoolValue(util.IsUnset(request.ExpectSessionCapacity)) {
		query["ExpectSessionCapacity"] = request.ExpectSessionCapacity
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateCapacityReservation"),
		Version:     tea.String("2021-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateCapacityReservationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateCapacityReservation(request *CreateCapacityReservationRequest) (_result *CreateCapacityReservationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateCapacityReservationResponse{}
	_body, _err := client.CreateCapacityReservationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDatasetDeployTaskWithOptions(request *CreateDatasetDeployTaskRequest, runtime *util.RuntimeOptions) (_result *CreateDatasetDeployTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.CustomParam)) {
		query["CustomParam"] = request.CustomParam
	}

	if !tea.BoolValue(util.IsUnset(request.NeedUnzip)) {
		query["NeedUnzip"] = request.NeedUnzip
	}

	if !tea.BoolValue(util.IsUnset(request.OssBucket)) {
		query["OssBucket"] = request.OssBucket
	}

	if !tea.BoolValue(util.IsUnset(request.OssFilePath)) {
		query["OssFilePath"] = request.OssFilePath
	}

	if !tea.BoolValue(util.IsUnset(request.OssRegionId)) {
		query["OssRegionId"] = request.OssRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceType)) {
		query["SourceType"] = request.SourceType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDatasetDeployTask"),
		Version:     tea.String("2021-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDatasetDeployTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDatasetDeployTask(request *CreateDatasetDeployTaskRequest) (_result *CreateDatasetDeployTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDatasetDeployTaskResponse{}
	_body, _err := client.CreateDatasetDeployTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateProjectWithOptions(tmpReq *CreateProjectRequest, runtime *util.RuntimeOptions) (_result *CreateProjectResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateProjectShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.BoundAppIdList)) {
		request.BoundAppIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.BoundAppIdList, tea.String("BoundAppIdList"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.ProjectQuotaLimit))) {
		request.ProjectQuotaLimitShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.ProjectQuotaLimit), tea.String("ProjectQuotaLimit"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BoundAppIdListShrink)) {
		body["BoundAppIdList"] = request.BoundAppIdListShrink
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		body["OperatorId"] = request.OperatorId
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorType)) {
		body["OperatorType"] = request.OperatorType
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectMemo)) {
		body["ProjectMemo"] = request.ProjectMemo
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		body["ProjectName"] = request.ProjectName
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectQuotaLimitShrink)) {
		body["ProjectQuotaLimit"] = request.ProjectQuotaLimitShrink
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateProject"),
		Version:     tea.String("2021-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
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
	_result = &CreateProjectResponse{}
	_body, _err := client.CreateProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteAppWithOptions(request *DeleteAppRequest, runtime *util.RuntimeOptions) (_result *DeleteAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteApp"),
		Version:     tea.String("2021-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteApp(request *DeleteAppRequest) (_result *DeleteAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAppResponse{}
	_body, _err := client.DeleteAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteAppVersionWithOptions(request *DeleteAppVersionRequest, runtime *util.RuntimeOptions) (_result *DeleteAppVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppVersionId)) {
		body["AppVersionId"] = request.AppVersionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAppVersion"),
		Version:     tea.String("2021-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteAppVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteAppVersion(request *DeleteAppVersionRequest) (_result *DeleteAppVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAppVersionResponse{}
	_body, _err := client.DeleteAppVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteProjectWithOptions(request *DeleteProjectRequest, runtime *util.RuntimeOptions) (_result *DeleteProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		body["OperatorId"] = request.OperatorId
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorType)) {
		body["OperatorType"] = request.OperatorType
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteProject"),
		Version:     tea.String("2021-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteProject(request *DeleteProjectRequest) (_result *DeleteProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteProjectResponse{}
	_body, _err := client.DeleteProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAdaptationWithOptions(request *GetAdaptationRequest, runtime *util.RuntimeOptions) (_result *GetAdaptationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AdaptApplyId)) {
		body["AdaptApplyId"] = request.AdaptApplyId
	}

	if !tea.BoolValue(util.IsUnset(request.AppVersionId)) {
		body["AppVersionId"] = request.AppVersionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAdaptation"),
		Version:     tea.String("2021-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAdaptationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAdaptation(request *GetAdaptationRequest) (_result *GetAdaptationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAdaptationResponse{}
	_body, _err := client.GetAdaptationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAppWithOptions(request *GetAppRequest, runtime *util.RuntimeOptions) (_result *GetAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetApp"),
		Version:     tea.String("2021-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetApp(request *GetAppRequest) (_result *GetAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAppResponse{}
	_body, _err := client.GetAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAppSessionWithOptions(request *GetAppSessionRequest, runtime *util.RuntimeOptions) (_result *GetAppSessionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CustomSessionId)) {
		query["CustomSessionId"] = request.CustomSessionId
	}

	if !tea.BoolValue(util.IsUnset(request.PlatformSessionId)) {
		query["PlatformSessionId"] = request.PlatformSessionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAppSession"),
		Version:     tea.String("2021-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAppSessionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAppSession(request *GetAppSessionRequest) (_result *GetAppSessionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAppSessionResponse{}
	_body, _err := client.GetAppSessionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAppVersionWithOptions(request *GetAppVersionRequest, runtime *util.RuntimeOptions) (_result *GetAppVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppVersionId)) {
		body["AppVersionId"] = request.AppVersionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAppVersion"),
		Version:     tea.String("2021-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAppVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAppVersion(request *GetAppVersionRequest) (_result *GetAppVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAppVersionResponse{}
	_body, _err := client.GetAppVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDatasetWithOptions(request *GetDatasetRequest, runtime *util.RuntimeOptions) (_result *GetDatasetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DatasetId)) {
		query["DatasetId"] = request.DatasetId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDataset"),
		Version:     tea.String("2021-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDatasetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDataset(request *GetDatasetRequest) (_result *GetDatasetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDatasetResponse{}
	_body, _err := client.GetDatasetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetProjectWithOptions(request *GetProjectRequest, runtime *util.RuntimeOptions) (_result *GetProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		body["OperatorId"] = request.OperatorId
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorType)) {
		body["OperatorType"] = request.OperatorType
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetProject"),
		Version:     tea.String("2021-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
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

func (client *Client) GetProject(request *GetProjectRequest) (_result *GetProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetProjectResponse{}
	_body, _err := client.GetProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAppWithOptions(request *ListAppRequest, runtime *util.RuntimeOptions) (_result *ListAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.KeySearch)) {
		body["KeySearch"] = request.KeySearch
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListApp"),
		Version:     tea.String("2021-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListApp(request *ListAppRequest) (_result *ListAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAppResponse{}
	_body, _err := client.ListAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAppSessionsWithOptions(request *ListAppSessionsRequest, runtime *util.RuntimeOptions) (_result *ListAppSessionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.CustomSessionIds)) {
		query["CustomSessionIds"] = request.CustomSessionIds
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PlatformSessionIds)) {
		query["PlatformSessionIds"] = request.PlatformSessionIds
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAppSessions"),
		Version:     tea.String("2021-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAppSessionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAppSessions(request *ListAppSessionsRequest) (_result *ListAppSessionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAppSessionsResponse{}
	_body, _err := client.ListAppSessionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAppVersionWithOptions(request *ListAppVersionRequest, runtime *util.RuntimeOptions) (_result *ListAppVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAppVersion"),
		Version:     tea.String("2021-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAppVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAppVersion(request *ListAppVersionRequest) (_result *ListAppVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAppVersionResponse{}
	_body, _err := client.ListAppVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyAppWithOptions(request *ModifyAppRequest, runtime *util.RuntimeOptions) (_result *ModifyAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		body["AppName"] = request.AppName
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyApp"),
		Version:     tea.String("2021-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyApp(request *ModifyAppRequest) (_result *ModifyAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyAppResponse{}
	_body, _err := client.ModifyAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyAppVersionWithOptions(request *ModifyAppVersionRequest, runtime *util.RuntimeOptions) (_result *ModifyAppVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppVersionId)) {
		body["AppVersionId"] = request.AppVersionId
	}

	if !tea.BoolValue(util.IsUnset(request.AppVersionName)) {
		body["AppVersionName"] = request.AppVersionName
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyAppVersion"),
		Version:     tea.String("2021-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyAppVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyAppVersion(request *ModifyAppVersionRequest) (_result *ModifyAppVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyAppVersionResponse{}
	_body, _err := client.ModifyAppVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyProjectWithOptions(tmpReq *ModifyProjectRequest, runtime *util.RuntimeOptions) (_result *ModifyProjectResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ModifyProjectShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.BoundAppIdList)) {
		request.BoundAppIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.BoundAppIdList, tea.String("BoundAppIdList"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.ProjectQuotaLimit))) {
		request.ProjectQuotaLimitShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.ProjectQuotaLimit), tea.String("ProjectQuotaLimit"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BoundAppIdListShrink)) {
		body["BoundAppIdList"] = request.BoundAppIdListShrink
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		body["OperatorId"] = request.OperatorId
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorType)) {
		body["OperatorType"] = request.OperatorType
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectMemo)) {
		body["ProjectMemo"] = request.ProjectMemo
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		body["ProjectName"] = request.ProjectName
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectQuotaLimitShrink)) {
		body["ProjectQuotaLimit"] = request.ProjectQuotaLimitShrink
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyProject"),
		Version:     tea.String("2021-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyProject(request *ModifyProjectRequest) (_result *ModifyProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyProjectResponse{}
	_body, _err := client.ModifyProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PageQueryProjectWithOptions(request *PageQueryProjectRequest, runtime *util.RuntimeOptions) (_result *PageQueryProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.KeySearch)) {
		body["KeySearch"] = request.KeySearch
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		body["OperatorId"] = request.OperatorId
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorType)) {
		body["OperatorType"] = request.OperatorType
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PageQueryProject"),
		Version:     tea.String("2021-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PageQueryProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PageQueryProject(request *PageQueryProjectRequest) (_result *PageQueryProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PageQueryProjectResponse{}
	_body, _err := client.PageQueryProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PageQueryProjectAppsWithOptions(request *PageQueryProjectAppsRequest, runtime *util.RuntimeOptions) (_result *PageQueryProjectAppsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		body["OperatorId"] = request.OperatorId
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorType)) {
		body["OperatorType"] = request.OperatorType
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PageQueryProjectApps"),
		Version:     tea.String("2021-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PageQueryProjectAppsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PageQueryProjectApps(request *PageQueryProjectAppsRequest) (_result *PageQueryProjectAppsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PageQueryProjectAppsResponse{}
	_body, _err := client.PageQueryProjectAppsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryOfflineTaskProgressWithOptions(request *QueryOfflineTaskProgressRequest, runtime *util.RuntimeOptions) (_result *QueryOfflineTaskProgressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.VersionId)) {
		body["VersionId"] = request.VersionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryOfflineTaskProgress"),
		Version:     tea.String("2021-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryOfflineTaskProgressResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryOfflineTaskProgress(request *QueryOfflineTaskProgressRequest) (_result *QueryOfflineTaskProgressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryOfflineTaskProgressResponse{}
	_body, _err := client.QueryOfflineTaskProgressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RefreshDistrictMetaWithOptions(runtime *util.RuntimeOptions) (_result *RefreshDistrictMetaResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("RefreshDistrictMeta"),
		Version:     tea.String("2021-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RefreshDistrictMetaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RefreshDistrictMeta() (_result *RefreshDistrictMetaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RefreshDistrictMetaResponse{}
	_body, _err := client.RefreshDistrictMetaWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopAppSessionWithOptions(request *StopAppSessionRequest, runtime *util.RuntimeOptions) (_result *StopAppSessionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CustomSessionId)) {
		query["CustomSessionId"] = request.CustomSessionId
	}

	if !tea.BoolValue(util.IsUnset(request.PlatformSessionId)) {
		query["PlatformSessionId"] = request.PlatformSessionId
	}

	if !tea.BoolValue(util.IsUnset(request.StopParam)) {
		query["StopParam"] = request.StopParam
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StopAppSession"),
		Version:     tea.String("2021-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StopAppSessionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopAppSession(request *StopAppSessionRequest) (_result *StopAppSessionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopAppSessionResponse{}
	_body, _err := client.StopAppSessionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitOfflineTaskWithOptions(request *SubmitOfflineTaskRequest, runtime *util.RuntimeOptions) (_result *SubmitOfflineTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		body["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.AppType)) {
		body["AppType"] = request.AppType
	}

	if !tea.BoolValue(util.IsUnset(request.Env)) {
		body["Env"] = request.Env
	}

	if !tea.BoolValue(util.IsUnset(request.Uri)) {
		body["Uri"] = request.Uri
	}

	if !tea.BoolValue(util.IsUnset(request.VersionId)) {
		body["VersionId"] = request.VersionId
	}

	if !tea.BoolValue(util.IsUnset(request.VersionName)) {
		body["VersionName"] = request.VersionName
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SubmitOfflineTask"),
		Version:     tea.String("2021-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitOfflineTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitOfflineTask(request *SubmitOfflineTaskRequest) (_result *SubmitOfflineTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitOfflineTaskResponse{}
	_body, _err := client.SubmitOfflineTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
