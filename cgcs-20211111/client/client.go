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
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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

func (s *CreateAdaptationResponseBody) SetRequestId(v string) *CreateAdaptationResponseBody {
	s.RequestId = &v
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
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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

func (s *CreateAppResponseBody) SetRequestId(v string) *CreateAppResponseBody {
	s.RequestId = &v
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
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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

func (s *CreateAppVersionResponseBody) SetRequestId(v string) *CreateAppVersionResponseBody {
	s.RequestId = &v
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
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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

func (s *DeleteAppResponseBody) SetRequestId(v string) *DeleteAppResponseBody {
	s.RequestId = &v
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
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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

func (s *DeleteAppVersionResponseBody) SetRequestId(v string) *DeleteAppVersionResponseBody {
	s.RequestId = &v
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
	GmtCreate    *string                               `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified  *string                               `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	RequestId    *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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

func (s *GetAdaptationResponseBody) SetGmtCreate(v string) *GetAdaptationResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *GetAdaptationResponseBody) SetGmtModified(v string) *GetAdaptationResponseBody {
	s.GmtModified = &v
	return s
}

func (s *GetAdaptationResponseBody) SetRequestId(v string) *GetAdaptationResponseBody {
	s.RequestId = &v
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
	GmtCreate       *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified     *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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

func (s *GetAppResponseBody) SetGmtCreate(v string) *GetAppResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *GetAppResponseBody) SetGmtModified(v string) *GetAppResponseBody {
	s.GmtModified = &v
	return s
}

func (s *GetAppResponseBody) SetRequestId(v string) *GetAppResponseBody {
	s.RequestId = &v
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
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
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
	ConsumeCu            *float64 `json:"ConsumeCu,omitempty" xml:"ConsumeCu,omitempty"`
	FileAddress          *string  `json:"FileAddress,omitempty" xml:"FileAddress,omitempty"`
	FileSize             *int64   `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	FileUploadFinishTime *string  `json:"FileUploadFinishTime,omitempty" xml:"FileUploadFinishTime,omitempty"`
	FileUploadType       *string  `json:"FileUploadType,omitempty" xml:"FileUploadType,omitempty"`
	GmtCreate            *string  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified          *string  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	RequestId            *string  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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

func (s *GetAppVersionResponseBody) SetRequestId(v string) *GetAppVersionResponseBody {
	s.RequestId = &v
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
	RequestId *string                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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

func (s *ListAppResponseBody) SetRequestId(v string) *ListAppResponseBody {
	s.RequestId = &v
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
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
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
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total     *int64                                `json:"Total,omitempty" xml:"Total,omitempty"`
	Versions  []*ListAppVersionResponseBodyVersions `json:"Versions,omitempty" xml:"Versions,omitempty" type:"Repeated"`
}

func (s ListAppVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAppVersionResponseBody) GoString() string {
	return s.String()
}

func (s *ListAppVersionResponseBody) SetRequestId(v string) *ListAppVersionResponseBody {
	s.RequestId = &v
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
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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

func (s *ModifyAppResponseBody) SetRequestId(v string) *ModifyAppResponseBody {
	s.RequestId = &v
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
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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

func (s *ModifyAppVersionResponseBody) SetRequestId(v string) *ModifyAppVersionResponseBody {
	s.RequestId = &v
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

type StopAppSessionRequest struct {
	// 自定义会话id
	CustomSessionId *string `json:"CustomSessionId,omitempty" xml:"CustomSessionId,omitempty"`
	// 自定义用户id
	PlatformSessionId *string `json:"PlatformSessionId,omitempty" xml:"PlatformSessionId,omitempty"`
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
