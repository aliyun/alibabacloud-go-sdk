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

type CancelReserveTaskRequest struct {
	// 客户端幂等性 token
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// 创建容量预定任务时返回的 taskId
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CancelReserveTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelReserveTaskRequest) GoString() string {
	return s.String()
}

func (s *CancelReserveTaskRequest) SetClientToken(v string) *CancelReserveTaskRequest {
	s.ClientToken = &v
	return s
}

func (s *CancelReserveTaskRequest) SetTaskId(v string) *CancelReserveTaskRequest {
	s.TaskId = &v
	return s
}

type CancelReserveTaskResponseBody struct {
	// pop的requestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CancelReserveTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelReserveTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CancelReserveTaskResponseBody) SetRequestId(v string) *CancelReserveTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelReserveTaskResponseBody) SetTaskId(v string) *CancelReserveTaskResponseBody {
	s.TaskId = &v
	return s
}

type CancelReserveTaskResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CancelReserveTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CancelReserveTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelReserveTaskResponse) GoString() string {
	return s.String()
}

func (s *CancelReserveTaskResponse) SetHeaders(v map[string]*string) *CancelReserveTaskResponse {
	s.Headers = v
	return s
}

func (s *CancelReserveTaskResponse) SetStatusCode(v int32) *CancelReserveTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelReserveTaskResponse) SetBody(v *CancelReserveTaskResponseBody) *CancelReserveTaskResponse {
	s.Body = v
	return s
}

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
	CustomUserId   *string `json:"CustomUserId,omitempty" xml:"CustomUserId,omitempty"`
	EnablePostpaid *bool   `json:"EnablePostpaid,omitempty" xml:"EnablePostpaid,omitempty"`
	// 启动参数
	StartParameters []*CreateAppSessionRequestStartParameters `json:"StartParameters,omitempty" xml:"StartParameters,omitempty" type:"Repeated"`
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

func (s *CreateAppSessionRequest) SetEnablePostpaid(v bool) *CreateAppSessionRequest {
	s.EnablePostpaid = &v
	return s
}

func (s *CreateAppSessionRequest) SetStartParameters(v []*CreateAppSessionRequestStartParameters) *CreateAppSessionRequest {
	s.StartParameters = v
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

type CreateAppSessionBatchSyncRequest struct {
	AppInfos []*CreateAppSessionBatchSyncRequestAppInfos `json:"AppInfos,omitempty" xml:"AppInfos,omitempty" type:"Repeated"`
	BatchId  *string                                     `json:"BatchId,omitempty" xml:"BatchId,omitempty"`
}

func (s CreateAppSessionBatchSyncRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAppSessionBatchSyncRequest) GoString() string {
	return s.String()
}

func (s *CreateAppSessionBatchSyncRequest) SetAppInfos(v []*CreateAppSessionBatchSyncRequestAppInfos) *CreateAppSessionBatchSyncRequest {
	s.AppInfos = v
	return s
}

func (s *CreateAppSessionBatchSyncRequest) SetBatchId(v string) *CreateAppSessionBatchSyncRequest {
	s.BatchId = &v
	return s
}

type CreateAppSessionBatchSyncRequestAppInfos struct {
	AppId             *string                                                    `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppVersion        *string                                                    `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	ClientIp          *string                                                    `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	CustomUserId      *string                                                    `json:"CustomUserId,omitempty" xml:"CustomUserId,omitempty"`
	CustomerSessionId *string                                                    `json:"CustomerSessionId,omitempty" xml:"CustomerSessionId,omitempty"`
	DistrictId        *string                                                    `json:"DistrictId,omitempty" xml:"DistrictId,omitempty"`
	ProjectId         *string                                                    `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	StartParameters   []*CreateAppSessionBatchSyncRequestAppInfosStartParameters `json:"StartParameters,omitempty" xml:"StartParameters,omitempty" type:"Repeated"`
	SystemInfo        []*CreateAppSessionBatchSyncRequestAppInfosSystemInfo      `json:"SystemInfo,omitempty" xml:"SystemInfo,omitempty" type:"Repeated"`
	Tags              []*CreateAppSessionBatchSyncRequestAppInfosTags            `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s CreateAppSessionBatchSyncRequestAppInfos) String() string {
	return tea.Prettify(s)
}

func (s CreateAppSessionBatchSyncRequestAppInfos) GoString() string {
	return s.String()
}

func (s *CreateAppSessionBatchSyncRequestAppInfos) SetAppId(v string) *CreateAppSessionBatchSyncRequestAppInfos {
	s.AppId = &v
	return s
}

func (s *CreateAppSessionBatchSyncRequestAppInfos) SetAppVersion(v string) *CreateAppSessionBatchSyncRequestAppInfos {
	s.AppVersion = &v
	return s
}

func (s *CreateAppSessionBatchSyncRequestAppInfos) SetClientIp(v string) *CreateAppSessionBatchSyncRequestAppInfos {
	s.ClientIp = &v
	return s
}

func (s *CreateAppSessionBatchSyncRequestAppInfos) SetCustomUserId(v string) *CreateAppSessionBatchSyncRequestAppInfos {
	s.CustomUserId = &v
	return s
}

func (s *CreateAppSessionBatchSyncRequestAppInfos) SetCustomerSessionId(v string) *CreateAppSessionBatchSyncRequestAppInfos {
	s.CustomerSessionId = &v
	return s
}

func (s *CreateAppSessionBatchSyncRequestAppInfos) SetDistrictId(v string) *CreateAppSessionBatchSyncRequestAppInfos {
	s.DistrictId = &v
	return s
}

func (s *CreateAppSessionBatchSyncRequestAppInfos) SetProjectId(v string) *CreateAppSessionBatchSyncRequestAppInfos {
	s.ProjectId = &v
	return s
}

func (s *CreateAppSessionBatchSyncRequestAppInfos) SetStartParameters(v []*CreateAppSessionBatchSyncRequestAppInfosStartParameters) *CreateAppSessionBatchSyncRequestAppInfos {
	s.StartParameters = v
	return s
}

func (s *CreateAppSessionBatchSyncRequestAppInfos) SetSystemInfo(v []*CreateAppSessionBatchSyncRequestAppInfosSystemInfo) *CreateAppSessionBatchSyncRequestAppInfos {
	s.SystemInfo = v
	return s
}

func (s *CreateAppSessionBatchSyncRequestAppInfos) SetTags(v []*CreateAppSessionBatchSyncRequestAppInfosTags) *CreateAppSessionBatchSyncRequestAppInfos {
	s.Tags = v
	return s
}

type CreateAppSessionBatchSyncRequestAppInfosStartParameters struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateAppSessionBatchSyncRequestAppInfosStartParameters) String() string {
	return tea.Prettify(s)
}

func (s CreateAppSessionBatchSyncRequestAppInfosStartParameters) GoString() string {
	return s.String()
}

func (s *CreateAppSessionBatchSyncRequestAppInfosStartParameters) SetKey(v string) *CreateAppSessionBatchSyncRequestAppInfosStartParameters {
	s.Key = &v
	return s
}

func (s *CreateAppSessionBatchSyncRequestAppInfosStartParameters) SetValue(v string) *CreateAppSessionBatchSyncRequestAppInfosStartParameters {
	s.Value = &v
	return s
}

type CreateAppSessionBatchSyncRequestAppInfosSystemInfo struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateAppSessionBatchSyncRequestAppInfosSystemInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateAppSessionBatchSyncRequestAppInfosSystemInfo) GoString() string {
	return s.String()
}

func (s *CreateAppSessionBatchSyncRequestAppInfosSystemInfo) SetKey(v string) *CreateAppSessionBatchSyncRequestAppInfosSystemInfo {
	s.Key = &v
	return s
}

func (s *CreateAppSessionBatchSyncRequestAppInfosSystemInfo) SetValue(v string) *CreateAppSessionBatchSyncRequestAppInfosSystemInfo {
	s.Value = &v
	return s
}

type CreateAppSessionBatchSyncRequestAppInfosTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateAppSessionBatchSyncRequestAppInfosTags) String() string {
	return tea.Prettify(s)
}

func (s CreateAppSessionBatchSyncRequestAppInfosTags) GoString() string {
	return s.String()
}

func (s *CreateAppSessionBatchSyncRequestAppInfosTags) SetKey(v string) *CreateAppSessionBatchSyncRequestAppInfosTags {
	s.Key = &v
	return s
}

func (s *CreateAppSessionBatchSyncRequestAppInfosTags) SetValue(v string) *CreateAppSessionBatchSyncRequestAppInfosTags {
	s.Value = &v
	return s
}

type CreateAppSessionBatchSyncShrinkRequest struct {
	AppInfosShrink *string `json:"AppInfos,omitempty" xml:"AppInfos,omitempty"`
	BatchId        *string `json:"BatchId,omitempty" xml:"BatchId,omitempty"`
}

func (s CreateAppSessionBatchSyncShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAppSessionBatchSyncShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateAppSessionBatchSyncShrinkRequest) SetAppInfosShrink(v string) *CreateAppSessionBatchSyncShrinkRequest {
	s.AppInfosShrink = &v
	return s
}

func (s *CreateAppSessionBatchSyncShrinkRequest) SetBatchId(v string) *CreateAppSessionBatchSyncShrinkRequest {
	s.BatchId = &v
	return s
}

type CreateAppSessionBatchSyncResponseBody struct {
	// 自定义会话id
	BatchId    *string                                            `json:"BatchId,omitempty" xml:"BatchId,omitempty"`
	FailedList []*CreateAppSessionBatchSyncResponseBodyFailedList `json:"FailedList,omitempty" xml:"FailedList,omitempty" type:"Repeated"`
	// 请求id
	RequestId  *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultList []*CreateAppSessionBatchSyncResponseBodyResultList `json:"ResultList,omitempty" xml:"ResultList,omitempty" type:"Repeated"`
}

func (s CreateAppSessionBatchSyncResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAppSessionBatchSyncResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAppSessionBatchSyncResponseBody) SetBatchId(v string) *CreateAppSessionBatchSyncResponseBody {
	s.BatchId = &v
	return s
}

func (s *CreateAppSessionBatchSyncResponseBody) SetFailedList(v []*CreateAppSessionBatchSyncResponseBodyFailedList) *CreateAppSessionBatchSyncResponseBody {
	s.FailedList = v
	return s
}

func (s *CreateAppSessionBatchSyncResponseBody) SetRequestId(v string) *CreateAppSessionBatchSyncResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAppSessionBatchSyncResponseBody) SetResultList(v []*CreateAppSessionBatchSyncResponseBodyResultList) *CreateAppSessionBatchSyncResponseBody {
	s.ResultList = v
	return s
}

type CreateAppSessionBatchSyncResponseBodyFailedList struct {
	AppId           *string                                                    `json:"AppId,omitempty" xml:"AppId,omitempty"`
	CustomSessionId *string                                                    `json:"CustomSessionId,omitempty" xml:"CustomSessionId,omitempty"`
	FailedInfo      *CreateAppSessionBatchSyncResponseBodyFailedListFailedInfo `json:"FailedInfo,omitempty" xml:"FailedInfo,omitempty" type:"Struct"`
}

func (s CreateAppSessionBatchSyncResponseBodyFailedList) String() string {
	return tea.Prettify(s)
}

func (s CreateAppSessionBatchSyncResponseBodyFailedList) GoString() string {
	return s.String()
}

func (s *CreateAppSessionBatchSyncResponseBodyFailedList) SetAppId(v string) *CreateAppSessionBatchSyncResponseBodyFailedList {
	s.AppId = &v
	return s
}

func (s *CreateAppSessionBatchSyncResponseBodyFailedList) SetCustomSessionId(v string) *CreateAppSessionBatchSyncResponseBodyFailedList {
	s.CustomSessionId = &v
	return s
}

func (s *CreateAppSessionBatchSyncResponseBodyFailedList) SetFailedInfo(v *CreateAppSessionBatchSyncResponseBodyFailedListFailedInfo) *CreateAppSessionBatchSyncResponseBodyFailedList {
	s.FailedInfo = v
	return s
}

type CreateAppSessionBatchSyncResponseBodyFailedListFailedInfo struct {
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
}

func (s CreateAppSessionBatchSyncResponseBodyFailedListFailedInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateAppSessionBatchSyncResponseBodyFailedListFailedInfo) GoString() string {
	return s.String()
}

func (s *CreateAppSessionBatchSyncResponseBodyFailedListFailedInfo) SetErrorCode(v string) *CreateAppSessionBatchSyncResponseBodyFailedListFailedInfo {
	s.ErrorCode = &v
	return s
}

func (s *CreateAppSessionBatchSyncResponseBodyFailedListFailedInfo) SetErrorMessage(v string) *CreateAppSessionBatchSyncResponseBodyFailedListFailedInfo {
	s.ErrorMessage = &v
	return s
}

type CreateAppSessionBatchSyncResponseBodyResultList struct {
	// 应用id
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 应用版本
	AppVersion *string                                                 `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	BizInfo    *CreateAppSessionBatchSyncResponseBodyResultListBizInfo `json:"BizInfo,omitempty" xml:"BizInfo,omitempty" type:"Struct"`
	// 自定义会话id
	CustomSessionId *string `json:"CustomSessionId,omitempty" xml:"CustomSessionId,omitempty"`
	// 平台会话id
	PlatformSessionId *string `json:"PlatformSessionId,omitempty" xml:"PlatformSessionId,omitempty"`
}

func (s CreateAppSessionBatchSyncResponseBodyResultList) String() string {
	return tea.Prettify(s)
}

func (s CreateAppSessionBatchSyncResponseBodyResultList) GoString() string {
	return s.String()
}

func (s *CreateAppSessionBatchSyncResponseBodyResultList) SetAppId(v string) *CreateAppSessionBatchSyncResponseBodyResultList {
	s.AppId = &v
	return s
}

func (s *CreateAppSessionBatchSyncResponseBodyResultList) SetAppVersion(v string) *CreateAppSessionBatchSyncResponseBodyResultList {
	s.AppVersion = &v
	return s
}

func (s *CreateAppSessionBatchSyncResponseBodyResultList) SetBizInfo(v *CreateAppSessionBatchSyncResponseBodyResultListBizInfo) *CreateAppSessionBatchSyncResponseBodyResultList {
	s.BizInfo = v
	return s
}

func (s *CreateAppSessionBatchSyncResponseBodyResultList) SetCustomSessionId(v string) *CreateAppSessionBatchSyncResponseBodyResultList {
	s.CustomSessionId = &v
	return s
}

func (s *CreateAppSessionBatchSyncResponseBodyResultList) SetPlatformSessionId(v string) *CreateAppSessionBatchSyncResponseBodyResultList {
	s.PlatformSessionId = &v
	return s
}

type CreateAppSessionBatchSyncResponseBodyResultListBizInfo struct {
	Biz       map[string]interface{}                                             `json:"Biz,omitempty" xml:"Biz,omitempty"`
	Endpoints []*CreateAppSessionBatchSyncResponseBodyResultListBizInfoEndpoints `json:"Endpoints,omitempty" xml:"Endpoints,omitempty" type:"Repeated"`
}

func (s CreateAppSessionBatchSyncResponseBodyResultListBizInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateAppSessionBatchSyncResponseBodyResultListBizInfo) GoString() string {
	return s.String()
}

func (s *CreateAppSessionBatchSyncResponseBodyResultListBizInfo) SetBiz(v map[string]interface{}) *CreateAppSessionBatchSyncResponseBodyResultListBizInfo {
	s.Biz = v
	return s
}

func (s *CreateAppSessionBatchSyncResponseBodyResultListBizInfo) SetEndpoints(v []*CreateAppSessionBatchSyncResponseBodyResultListBizInfoEndpoints) *CreateAppSessionBatchSyncResponseBodyResultListBizInfo {
	s.Endpoints = v
	return s
}

type CreateAppSessionBatchSyncResponseBodyResultListBizInfoEndpoints struct {
	AccessHost *string `json:"AccessHost,omitempty" xml:"AccessHost,omitempty"`
	AccessPort *string `json:"AccessPort,omitempty" xml:"AccessPort,omitempty"`
	DistrictId *string `json:"DistrictId,omitempty" xml:"DistrictId,omitempty"`
	Isp        *string `json:"Isp,omitempty" xml:"Isp,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Type       *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateAppSessionBatchSyncResponseBodyResultListBizInfoEndpoints) String() string {
	return tea.Prettify(s)
}

func (s CreateAppSessionBatchSyncResponseBodyResultListBizInfoEndpoints) GoString() string {
	return s.String()
}

func (s *CreateAppSessionBatchSyncResponseBodyResultListBizInfoEndpoints) SetAccessHost(v string) *CreateAppSessionBatchSyncResponseBodyResultListBizInfoEndpoints {
	s.AccessHost = &v
	return s
}

func (s *CreateAppSessionBatchSyncResponseBodyResultListBizInfoEndpoints) SetAccessPort(v string) *CreateAppSessionBatchSyncResponseBodyResultListBizInfoEndpoints {
	s.AccessPort = &v
	return s
}

func (s *CreateAppSessionBatchSyncResponseBodyResultListBizInfoEndpoints) SetDistrictId(v string) *CreateAppSessionBatchSyncResponseBodyResultListBizInfoEndpoints {
	s.DistrictId = &v
	return s
}

func (s *CreateAppSessionBatchSyncResponseBodyResultListBizInfoEndpoints) SetIsp(v string) *CreateAppSessionBatchSyncResponseBodyResultListBizInfoEndpoints {
	s.Isp = &v
	return s
}

func (s *CreateAppSessionBatchSyncResponseBodyResultListBizInfoEndpoints) SetName(v string) *CreateAppSessionBatchSyncResponseBodyResultListBizInfoEndpoints {
	s.Name = &v
	return s
}

func (s *CreateAppSessionBatchSyncResponseBodyResultListBizInfoEndpoints) SetType(v string) *CreateAppSessionBatchSyncResponseBodyResultListBizInfoEndpoints {
	s.Type = &v
	return s
}

type CreateAppSessionBatchSyncResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateAppSessionBatchSyncResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateAppSessionBatchSyncResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAppSessionBatchSyncResponse) GoString() string {
	return s.String()
}

func (s *CreateAppSessionBatchSyncResponse) SetHeaders(v map[string]*string) *CreateAppSessionBatchSyncResponse {
	s.Headers = v
	return s
}

func (s *CreateAppSessionBatchSyncResponse) SetStatusCode(v int32) *CreateAppSessionBatchSyncResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAppSessionBatchSyncResponse) SetBody(v *CreateAppSessionBatchSyncResponseBody) *CreateAppSessionBatchSyncResponse {
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
	CustomUserId *string `json:"CustomUserId,omitempty" xml:"CustomUserId,omitempty"`
	DistrictId   *string `json:"DistrictId,omitempty" xml:"DistrictId,omitempty"`
	ProjectId    *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// 启动参数
	StartParameters []*CreateAppSessionSyncRequestStartParameters `json:"StartParameters,omitempty" xml:"StartParameters,omitempty" type:"Repeated"`
	// 系统信息：如端侧机型等信息
	SystemInfo []*CreateAppSessionSyncRequestSystemInfo `json:"SystemInfo,omitempty" xml:"SystemInfo,omitempty" type:"Repeated"`
	Tags       []*CreateAppSessionSyncRequestTags       `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
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
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Biz       map[string]interface{}                              `json:"Biz,omitempty" xml:"Biz,omitempty"`
	Endpoints []*CreateAppSessionSyncResponseBodyBizInfoEndpoints `json:"Endpoints,omitempty" xml:"Endpoints,omitempty" type:"Repeated"`
}

func (s CreateAppSessionSyncResponseBodyBizInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateAppSessionSyncResponseBodyBizInfo) GoString() string {
	return s.String()
}

func (s *CreateAppSessionSyncResponseBodyBizInfo) SetBiz(v map[string]interface{}) *CreateAppSessionSyncResponseBodyBizInfo {
	s.Biz = v
	return s
}

func (s *CreateAppSessionSyncResponseBodyBizInfo) SetEndpoints(v []*CreateAppSessionSyncResponseBodyBizInfoEndpoints) *CreateAppSessionSyncResponseBodyBizInfo {
	s.Endpoints = v
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

type CreateCapacityReservationRequest struct {
	// 应用ID
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 应用版本ID，如果不填，采用控制台设置的默认版本
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	// 客户端 token，保持幂等性。
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// 大区 id
	DistrictId *string `json:"DistrictId,omitempty" xml:"DistrictId,omitempty"`
	// 预期资源ready可服务时间 时间格式 yyyy-mm-dd HH:MM:SS 比如
	ExpectResourceReadyTime *string `json:"ExpectResourceReadyTime,omitempty" xml:"ExpectResourceReadyTime,omitempty"`
	// 预期创建出的容器能支持的 session 数量。GCS 内部会根据适配自动计算出所需要的资源量，以及多种机型的配比。
	ExpectSessionCapacity *int32 `json:"ExpectSessionCapacity,omitempty" xml:"ExpectSessionCapacity,omitempty"`
	// 项目ID
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
	// 当前最大可分配的 session 库存，注意两次调用之间，该库存值可能有变动
	CurrMaxAllocatableSessionCapacity *int32 `json:"CurrMaxAllocatableSessionCapacity,omitempty" xml:"CurrMaxAllocatableSessionCapacity,omitempty"`
	// pop的requestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 容器创建任务 id
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

type GetAppCcuRequest struct {
	// 自定义会话id
	AppId      *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	// 平台会话id
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s GetAppCcuRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAppCcuRequest) GoString() string {
	return s.String()
}

func (s *GetAppCcuRequest) SetAppId(v string) *GetAppCcuRequest {
	s.AppId = &v
	return s
}

func (s *GetAppCcuRequest) SetAppVersion(v string) *GetAppCcuRequest {
	s.AppVersion = &v
	return s
}

func (s *GetAppCcuRequest) SetProjectId(v string) *GetAppCcuRequest {
	s.ProjectId = &v
	return s
}

type GetAppCcuResponseBody struct {
	// 请求id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 自定义会话id
	Timestamp *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s GetAppCcuResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAppCcuResponseBody) GoString() string {
	return s.String()
}

func (s *GetAppCcuResponseBody) SetRequestId(v string) *GetAppCcuResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAppCcuResponseBody) SetTimestamp(v string) *GetAppCcuResponseBody {
	s.Timestamp = &v
	return s
}

type GetAppCcuResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetAppCcuResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAppCcuResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAppCcuResponse) GoString() string {
	return s.String()
}

func (s *GetAppCcuResponse) SetHeaders(v map[string]*string) *GetAppCcuResponse {
	s.Headers = v
	return s
}

func (s *GetAppCcuResponse) SetStatusCode(v int32) *GetAppCcuResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAppCcuResponse) SetBody(v *GetAppCcuResponseBody) *GetAppCcuResponse {
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
	AppVersion *string                           `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	BizInfo    *GetAppSessionResponseBodyBizInfo `json:"BizInfo,omitempty" xml:"BizInfo,omitempty" type:"Struct"`
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

func (s *GetAppSessionResponseBody) SetBizInfo(v *GetAppSessionResponseBodyBizInfo) *GetAppSessionResponseBody {
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
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	StopTime  *string `json:"StopTime,omitempty" xml:"StopTime,omitempty"`
}

func (s GetAppSessionResponseBodyBizInfo) String() string {
	return tea.Prettify(s)
}

func (s GetAppSessionResponseBodyBizInfo) GoString() string {
	return s.String()
}

func (s *GetAppSessionResponseBodyBizInfo) SetStartTime(v string) *GetAppSessionResponseBodyBizInfo {
	s.StartTime = &v
	return s
}

func (s *GetAppSessionResponseBodyBizInfo) SetStopTime(v string) *GetAppSessionResponseBodyBizInfo {
	s.StopTime = &v
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

type GetCapacityRequest struct {
	// 按照 appId 来匹配
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 按照 app 版本来匹配
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	// 大区 id
	DistrictId *string `json:"DistrictId,omitempty" xml:"DistrictId,omitempty"`
	// 第几页，默认从 1 开始
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// 一页大小，默认 20，最大 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 项目 id
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s GetCapacityRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCapacityRequest) GoString() string {
	return s.String()
}

func (s *GetCapacityRequest) SetAppId(v string) *GetCapacityRequest {
	s.AppId = &v
	return s
}

func (s *GetCapacityRequest) SetAppVersion(v string) *GetCapacityRequest {
	s.AppVersion = &v
	return s
}

func (s *GetCapacityRequest) SetDistrictId(v string) *GetCapacityRequest {
	s.DistrictId = &v
	return s
}

func (s *GetCapacityRequest) SetPageNum(v int32) *GetCapacityRequest {
	s.PageNum = &v
	return s
}

func (s *GetCapacityRequest) SetPageSize(v int32) *GetCapacityRequest {
	s.PageSize = &v
	return s
}

func (s *GetCapacityRequest) SetProjectId(v string) *GetCapacityRequest {
	s.ProjectId = &v
	return s
}

type GetCapacityResponseBody struct {
	// 回满足匹配的
	Capacities []*GetCapacityResponseBodyCapacities `json:"Capacities,omitempty" xml:"Capacities,omitempty" type:"Repeated"`
	// 第几页
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// 每页大小
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// pop的requestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 满足匹配条件的总量
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetCapacityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCapacityResponseBody) GoString() string {
	return s.String()
}

func (s *GetCapacityResponseBody) SetCapacities(v []*GetCapacityResponseBodyCapacities) *GetCapacityResponseBody {
	s.Capacities = v
	return s
}

func (s *GetCapacityResponseBody) SetPageNum(v int32) *GetCapacityResponseBody {
	s.PageNum = &v
	return s
}

func (s *GetCapacityResponseBody) SetPageSize(v int32) *GetCapacityResponseBody {
	s.PageSize = &v
	return s
}

func (s *GetCapacityResponseBody) SetRequestId(v string) *GetCapacityResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCapacityResponseBody) SetTotal(v int32) *GetCapacityResponseBody {
	s.Total = &v
	return s
}

type GetCapacityResponseBodyCapacities struct {
	// 创建容器时指定的 appId
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 创建容器时指定的版本
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	// 该容器所属大区 id
	DistrictId *string `json:"DistrictId,omitempty" xml:"DistrictId,omitempty"`
	// 可支撑 session 数量
	SessionCapacity *int32 `json:"SessionCapacity,omitempty" xml:"SessionCapacity,omitempty"`
}

func (s GetCapacityResponseBodyCapacities) String() string {
	return tea.Prettify(s)
}

func (s GetCapacityResponseBodyCapacities) GoString() string {
	return s.String()
}

func (s *GetCapacityResponseBodyCapacities) SetAppId(v string) *GetCapacityResponseBodyCapacities {
	s.AppId = &v
	return s
}

func (s *GetCapacityResponseBodyCapacities) SetAppVersion(v string) *GetCapacityResponseBodyCapacities {
	s.AppVersion = &v
	return s
}

func (s *GetCapacityResponseBodyCapacities) SetDistrictId(v string) *GetCapacityResponseBodyCapacities {
	s.DistrictId = &v
	return s
}

func (s *GetCapacityResponseBodyCapacities) SetSessionCapacity(v int32) *GetCapacityResponseBodyCapacities {
	s.SessionCapacity = &v
	return s
}

type GetCapacityResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetCapacityResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetCapacityResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCapacityResponse) GoString() string {
	return s.String()
}

func (s *GetCapacityResponse) SetHeaders(v map[string]*string) *GetCapacityResponse {
	s.Headers = v
	return s
}

func (s *GetCapacityResponse) SetStatusCode(v int32) *GetCapacityResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCapacityResponse) SetBody(v *GetCapacityResponseBody) *GetCapacityResponse {
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

type GetReserveTaskDetailRequest struct {
	// 创建容量预定任务时返回的 taskId
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetReserveTaskDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s GetReserveTaskDetailRequest) GoString() string {
	return s.String()
}

func (s *GetReserveTaskDetailRequest) SetTaskId(v string) *GetReserveTaskDetailRequest {
	s.TaskId = &v
	return s
}

type GetReserveTaskDetailResponseBody struct {
	// 容量预定时的 appId
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 容量预定时的 appVersion
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	// 当前任务已经生产完成的会话路数
	CurrCompletedSessionCapacity *int32 `json:"CurrCompletedSessionCapacity,omitempty" xml:"CurrCompletedSessionCapacity,omitempty"`
	// 容量预定时指定的大区id
	DistrictId *string `json:"DistrictId,omitempty" xml:"DistrictId,omitempty"`
	// 容量预定，期望生效时间
	ExpectResourceReadyTime *string `json:"ExpectResourceReadyTime,omitempty" xml:"ExpectResourceReadyTime,omitempty"`
	// 容量预定期望的会话路数
	ExpectSessionCapacity *int32 `json:"ExpectSessionCapacity,omitempty" xml:"ExpectSessionCapacity,omitempty"`
	// 容量预定时的 projectId
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// pop的requestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 容器创建任务 id
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// 任务状态
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
}

func (s GetReserveTaskDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetReserveTaskDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetReserveTaskDetailResponseBody) SetAppId(v string) *GetReserveTaskDetailResponseBody {
	s.AppId = &v
	return s
}

func (s *GetReserveTaskDetailResponseBody) SetAppVersion(v string) *GetReserveTaskDetailResponseBody {
	s.AppVersion = &v
	return s
}

func (s *GetReserveTaskDetailResponseBody) SetCurrCompletedSessionCapacity(v int32) *GetReserveTaskDetailResponseBody {
	s.CurrCompletedSessionCapacity = &v
	return s
}

func (s *GetReserveTaskDetailResponseBody) SetDistrictId(v string) *GetReserveTaskDetailResponseBody {
	s.DistrictId = &v
	return s
}

func (s *GetReserveTaskDetailResponseBody) SetExpectResourceReadyTime(v string) *GetReserveTaskDetailResponseBody {
	s.ExpectResourceReadyTime = &v
	return s
}

func (s *GetReserveTaskDetailResponseBody) SetExpectSessionCapacity(v int32) *GetReserveTaskDetailResponseBody {
	s.ExpectSessionCapacity = &v
	return s
}

func (s *GetReserveTaskDetailResponseBody) SetProjectId(v string) *GetReserveTaskDetailResponseBody {
	s.ProjectId = &v
	return s
}

func (s *GetReserveTaskDetailResponseBody) SetRequestId(v string) *GetReserveTaskDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetReserveTaskDetailResponseBody) SetTaskId(v string) *GetReserveTaskDetailResponseBody {
	s.TaskId = &v
	return s
}

func (s *GetReserveTaskDetailResponseBody) SetTaskStatus(v string) *GetReserveTaskDetailResponseBody {
	s.TaskStatus = &v
	return s
}

type GetReserveTaskDetailResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetReserveTaskDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetReserveTaskDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s GetReserveTaskDetailResponse) GoString() string {
	return s.String()
}

func (s *GetReserveTaskDetailResponse) SetHeaders(v map[string]*string) *GetReserveTaskDetailResponse {
	s.Headers = v
	return s
}

func (s *GetReserveTaskDetailResponse) SetStatusCode(v int32) *GetReserveTaskDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetReserveTaskDetailResponse) SetBody(v *GetReserveTaskDetailResponseBody) *GetReserveTaskDetailResponse {
	s.Body = v
	return s
}

type GetResourcePublicIPsRequest struct {
	// 第几页，默认从 1 开始
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// 一页大小，默认 20，最大 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 项目 id
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s GetResourcePublicIPsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetResourcePublicIPsRequest) GoString() string {
	return s.String()
}

func (s *GetResourcePublicIPsRequest) SetPageNum(v int32) *GetResourcePublicIPsRequest {
	s.PageNum = &v
	return s
}

func (s *GetResourcePublicIPsRequest) SetPageSize(v int32) *GetResourcePublicIPsRequest {
	s.PageSize = &v
	return s
}

func (s *GetResourcePublicIPsRequest) SetProjectId(v string) *GetResourcePublicIPsRequest {
	s.ProjectId = &v
	return s
}

type GetResourcePublicIPsResponseBody struct {
	// ip 列表
	IpList []*GetResourcePublicIPsResponseBodyIpList `json:"IpList,omitempty" xml:"IpList,omitempty" type:"Repeated"`
	// 第几页
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// 每页大小
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// pop的requestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 满足匹配条件的总量
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetResourcePublicIPsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetResourcePublicIPsResponseBody) GoString() string {
	return s.String()
}

func (s *GetResourcePublicIPsResponseBody) SetIpList(v []*GetResourcePublicIPsResponseBodyIpList) *GetResourcePublicIPsResponseBody {
	s.IpList = v
	return s
}

func (s *GetResourcePublicIPsResponseBody) SetPageNum(v int32) *GetResourcePublicIPsResponseBody {
	s.PageNum = &v
	return s
}

func (s *GetResourcePublicIPsResponseBody) SetPageSize(v int32) *GetResourcePublicIPsResponseBody {
	s.PageSize = &v
	return s
}

func (s *GetResourcePublicIPsResponseBody) SetRequestId(v string) *GetResourcePublicIPsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetResourcePublicIPsResponseBody) SetTotal(v int32) *GetResourcePublicIPsResponseBody {
	s.Total = &v
	return s
}

type GetResourcePublicIPsResponseBodyIpList struct {
	// ip
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// 项目 id
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s GetResourcePublicIPsResponseBodyIpList) String() string {
	return tea.Prettify(s)
}

func (s GetResourcePublicIPsResponseBodyIpList) GoString() string {
	return s.String()
}

func (s *GetResourcePublicIPsResponseBodyIpList) SetIp(v string) *GetResourcePublicIPsResponseBodyIpList {
	s.Ip = &v
	return s
}

func (s *GetResourcePublicIPsResponseBodyIpList) SetProjectId(v string) *GetResourcePublicIPsResponseBodyIpList {
	s.ProjectId = &v
	return s
}

type GetResourcePublicIPsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetResourcePublicIPsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetResourcePublicIPsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetResourcePublicIPsResponse) GoString() string {
	return s.String()
}

func (s *GetResourcePublicIPsResponse) SetHeaders(v map[string]*string) *GetResourcePublicIPsResponse {
	s.Headers = v
	return s
}

func (s *GetResourcePublicIPsResponse) SetStatusCode(v int32) *GetResourcePublicIPsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetResourcePublicIPsResponse) SetBody(v *GetResourcePublicIPsResponseBody) *GetResourcePublicIPsResponse {
	s.Body = v
	return s
}

type GetTenantResponseBody struct {
	// 业务处理结果Code
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 业务对象
	Data *GetTenantResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// 业务处理消息摘要
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 操作请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 业务处理是否成功
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetTenantResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTenantResponseBody) GoString() string {
	return s.String()
}

func (s *GetTenantResponseBody) SetCode(v string) *GetTenantResponseBody {
	s.Code = &v
	return s
}

func (s *GetTenantResponseBody) SetData(v *GetTenantResponseBodyData) *GetTenantResponseBody {
	s.Data = v
	return s
}

func (s *GetTenantResponseBody) SetMessage(v string) *GetTenantResponseBody {
	s.Message = &v
	return s
}

func (s *GetTenantResponseBody) SetRequestId(v string) *GetTenantResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTenantResponseBody) SetSuccess(v bool) *GetTenantResponseBody {
	s.Success = &v
	return s
}

type GetTenantResponseBodyData struct {
	// 计收模式
	ChargeMode *string `json:"ChargeMode,omitempty" xml:"ChargeMode,omitempty"`
	// 联系人电话
	ContactsMobile *string `json:"ContactsMobile,omitempty" xml:"ContactsMobile,omitempty"`
	// 联系人
	ContactsName *string `json:"ContactsName,omitempty" xml:"ContactsName,omitempty"`
	// 租户所属行业(Code)
	IndustryCategory *string `json:"IndustryCategory,omitempty" xml:"IndustryCategory,omitempty"`
	// 业务场景描述
	ScenceDesc *string `json:"ScenceDesc,omitempty" xml:"ScenceDesc,omitempty"`
	// 租户状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 租户ID(AlipayUserId)
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// 租户名称
	TenantName *string `json:"TenantName,omitempty" xml:"TenantName,omitempty"`
}

func (s GetTenantResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetTenantResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTenantResponseBodyData) SetChargeMode(v string) *GetTenantResponseBodyData {
	s.ChargeMode = &v
	return s
}

func (s *GetTenantResponseBodyData) SetContactsMobile(v string) *GetTenantResponseBodyData {
	s.ContactsMobile = &v
	return s
}

func (s *GetTenantResponseBodyData) SetContactsName(v string) *GetTenantResponseBodyData {
	s.ContactsName = &v
	return s
}

func (s *GetTenantResponseBodyData) SetIndustryCategory(v string) *GetTenantResponseBodyData {
	s.IndustryCategory = &v
	return s
}

func (s *GetTenantResponseBodyData) SetScenceDesc(v string) *GetTenantResponseBodyData {
	s.ScenceDesc = &v
	return s
}

func (s *GetTenantResponseBodyData) SetStatus(v string) *GetTenantResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetTenantResponseBodyData) SetTenantId(v string) *GetTenantResponseBodyData {
	s.TenantId = &v
	return s
}

func (s *GetTenantResponseBodyData) SetTenantName(v string) *GetTenantResponseBodyData {
	s.TenantName = &v
	return s
}

type GetTenantResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetTenantResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTenantResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTenantResponse) GoString() string {
	return s.String()
}

func (s *GetTenantResponse) SetHeaders(v map[string]*string) *GetTenantResponse {
	s.Headers = v
	return s
}

func (s *GetTenantResponse) SetStatusCode(v int32) *GetTenantResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTenantResponse) SetBody(v *GetTenantResponseBody) *GetTenantResponse {
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
	AppVersion *string                                        `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	BizInfo    *ListAppSessionsResponseBodyAppSessionsBizInfo `json:"BizInfo,omitempty" xml:"BizInfo,omitempty" type:"Struct"`
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

func (s *ListAppSessionsResponseBodyAppSessions) SetBizInfo(v *ListAppSessionsResponseBodyAppSessionsBizInfo) *ListAppSessionsResponseBodyAppSessions {
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
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	StopTime  *string `json:"StopTime,omitempty" xml:"StopTime,omitempty"`
}

func (s ListAppSessionsResponseBodyAppSessionsBizInfo) String() string {
	return tea.Prettify(s)
}

func (s ListAppSessionsResponseBodyAppSessionsBizInfo) GoString() string {
	return s.String()
}

func (s *ListAppSessionsResponseBodyAppSessionsBizInfo) SetStartTime(v string) *ListAppSessionsResponseBodyAppSessionsBizInfo {
	s.StartTime = &v
	return s
}

func (s *ListAppSessionsResponseBodyAppSessionsBizInfo) SetStopTime(v string) *ListAppSessionsResponseBodyAppSessionsBizInfo {
	s.StopTime = &v
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
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
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

func (s *PageQueryProjectAppsRequest) SetAppId(v string) *PageQueryProjectAppsRequest {
	s.AppId = &v
	return s
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

type ReleaseCapacityRequest struct {
	// app id
	AppId      *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	// 大区 id
	DistrictId *string `json:"DistrictId,omitempty" xml:"DistrictId,omitempty"`
	// 期望释放的会话路数
	ExpectReleaseSessionCapacity *int32 `json:"ExpectReleaseSessionCapacity,omitempty" xml:"ExpectReleaseSessionCapacity,omitempty"`
	// 项目 id
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s ReleaseCapacityRequest) String() string {
	return tea.Prettify(s)
}

func (s ReleaseCapacityRequest) GoString() string {
	return s.String()
}

func (s *ReleaseCapacityRequest) SetAppId(v string) *ReleaseCapacityRequest {
	s.AppId = &v
	return s
}

func (s *ReleaseCapacityRequest) SetAppVersion(v string) *ReleaseCapacityRequest {
	s.AppVersion = &v
	return s
}

func (s *ReleaseCapacityRequest) SetDistrictId(v string) *ReleaseCapacityRequest {
	s.DistrictId = &v
	return s
}

func (s *ReleaseCapacityRequest) SetExpectReleaseSessionCapacity(v int32) *ReleaseCapacityRequest {
	s.ExpectReleaseSessionCapacity = &v
	return s
}

func (s *ReleaseCapacityRequest) SetProjectId(v string) *ReleaseCapacityRequest {
	s.ProjectId = &v
	return s
}

type ReleaseCapacityResponseBody struct {
	// pop的requestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 容器异步释放任务 id
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ReleaseCapacityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReleaseCapacityResponseBody) GoString() string {
	return s.String()
}

func (s *ReleaseCapacityResponseBody) SetRequestId(v string) *ReleaseCapacityResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReleaseCapacityResponseBody) SetTaskId(v string) *ReleaseCapacityResponseBody {
	s.TaskId = &v
	return s
}

type ReleaseCapacityResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ReleaseCapacityResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ReleaseCapacityResponse) String() string {
	return tea.Prettify(s)
}

func (s ReleaseCapacityResponse) GoString() string {
	return s.String()
}

func (s *ReleaseCapacityResponse) SetHeaders(v map[string]*string) *ReleaseCapacityResponse {
	s.Headers = v
	return s
}

func (s *ReleaseCapacityResponse) SetStatusCode(v int32) *ReleaseCapacityResponse {
	s.StatusCode = &v
	return s
}

func (s *ReleaseCapacityResponse) SetBody(v *ReleaseCapacityResponseBody) *ReleaseCapacityResponse {
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
	Key   *string     `json:"Key,omitempty" xml:"Key,omitempty"`
	Value interface{} `json:"Value,omitempty" xml:"Value,omitempty"`
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

func (s *StopAppSessionRequestStopParam) SetValue(v interface{}) *StopAppSessionRequestStopParam {
	s.Value = v
	return s
}

type StopAppSessionShrinkRequest struct {
	// 自定义会话id
	CustomSessionId *string `json:"CustomSessionId,omitempty" xml:"CustomSessionId,omitempty"`
	// 自定义用户id
	PlatformSessionId *string `json:"PlatformSessionId,omitempty" xml:"PlatformSessionId,omitempty"`
	StopParamShrink   *string `json:"StopParam,omitempty" xml:"StopParam,omitempty"`
}

func (s StopAppSessionShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s StopAppSessionShrinkRequest) GoString() string {
	return s.String()
}

func (s *StopAppSessionShrinkRequest) SetCustomSessionId(v string) *StopAppSessionShrinkRequest {
	s.CustomSessionId = &v
	return s
}

func (s *StopAppSessionShrinkRequest) SetPlatformSessionId(v string) *StopAppSessionShrinkRequest {
	s.PlatformSessionId = &v
	return s
}

func (s *StopAppSessionShrinkRequest) SetStopParamShrink(v string) *StopAppSessionShrinkRequest {
	s.StopParamShrink = &v
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

type StopAppSessionBatchRequest struct {
	AppId      *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	// 自定义用户id
	BatchId *string `json:"BatchId,omitempty" xml:"BatchId,omitempty"`
	// 自定义会话id
	ProjectId *string                                `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	StopParam []*StopAppSessionBatchRequestStopParam `json:"StopParam,omitempty" xml:"StopParam,omitempty" type:"Repeated"`
	Tags      []*StopAppSessionBatchRequestTags      `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s StopAppSessionBatchRequest) String() string {
	return tea.Prettify(s)
}

func (s StopAppSessionBatchRequest) GoString() string {
	return s.String()
}

func (s *StopAppSessionBatchRequest) SetAppId(v string) *StopAppSessionBatchRequest {
	s.AppId = &v
	return s
}

func (s *StopAppSessionBatchRequest) SetAppVersion(v string) *StopAppSessionBatchRequest {
	s.AppVersion = &v
	return s
}

func (s *StopAppSessionBatchRequest) SetBatchId(v string) *StopAppSessionBatchRequest {
	s.BatchId = &v
	return s
}

func (s *StopAppSessionBatchRequest) SetProjectId(v string) *StopAppSessionBatchRequest {
	s.ProjectId = &v
	return s
}

func (s *StopAppSessionBatchRequest) SetStopParam(v []*StopAppSessionBatchRequestStopParam) *StopAppSessionBatchRequest {
	s.StopParam = v
	return s
}

func (s *StopAppSessionBatchRequest) SetTags(v []*StopAppSessionBatchRequestTags) *StopAppSessionBatchRequest {
	s.Tags = v
	return s
}

type StopAppSessionBatchRequestStopParam struct {
	Key   *string     `json:"Key,omitempty" xml:"Key,omitempty"`
	Value interface{} `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s StopAppSessionBatchRequestStopParam) String() string {
	return tea.Prettify(s)
}

func (s StopAppSessionBatchRequestStopParam) GoString() string {
	return s.String()
}

func (s *StopAppSessionBatchRequestStopParam) SetKey(v string) *StopAppSessionBatchRequestStopParam {
	s.Key = &v
	return s
}

func (s *StopAppSessionBatchRequestStopParam) SetValue(v interface{}) *StopAppSessionBatchRequestStopParam {
	s.Value = v
	return s
}

type StopAppSessionBatchRequestTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s StopAppSessionBatchRequestTags) String() string {
	return tea.Prettify(s)
}

func (s StopAppSessionBatchRequestTags) GoString() string {
	return s.String()
}

func (s *StopAppSessionBatchRequestTags) SetKey(v string) *StopAppSessionBatchRequestTags {
	s.Key = &v
	return s
}

func (s *StopAppSessionBatchRequestTags) SetValue(v string) *StopAppSessionBatchRequestTags {
	s.Value = &v
	return s
}

type StopAppSessionBatchShrinkRequest struct {
	AppId      *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	// 自定义用户id
	BatchId *string `json:"BatchId,omitempty" xml:"BatchId,omitempty"`
	// 自定义会话id
	ProjectId       *string                                 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	StopParamShrink *string                                 `json:"StopParam,omitempty" xml:"StopParam,omitempty"`
	Tags            []*StopAppSessionBatchShrinkRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s StopAppSessionBatchShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s StopAppSessionBatchShrinkRequest) GoString() string {
	return s.String()
}

func (s *StopAppSessionBatchShrinkRequest) SetAppId(v string) *StopAppSessionBatchShrinkRequest {
	s.AppId = &v
	return s
}

func (s *StopAppSessionBatchShrinkRequest) SetAppVersion(v string) *StopAppSessionBatchShrinkRequest {
	s.AppVersion = &v
	return s
}

func (s *StopAppSessionBatchShrinkRequest) SetBatchId(v string) *StopAppSessionBatchShrinkRequest {
	s.BatchId = &v
	return s
}

func (s *StopAppSessionBatchShrinkRequest) SetProjectId(v string) *StopAppSessionBatchShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *StopAppSessionBatchShrinkRequest) SetStopParamShrink(v string) *StopAppSessionBatchShrinkRequest {
	s.StopParamShrink = &v
	return s
}

func (s *StopAppSessionBatchShrinkRequest) SetTags(v []*StopAppSessionBatchShrinkRequestTags) *StopAppSessionBatchShrinkRequest {
	s.Tags = v
	return s
}

type StopAppSessionBatchShrinkRequestTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s StopAppSessionBatchShrinkRequestTags) String() string {
	return tea.Prettify(s)
}

func (s StopAppSessionBatchShrinkRequestTags) GoString() string {
	return s.String()
}

func (s *StopAppSessionBatchShrinkRequestTags) SetKey(v string) *StopAppSessionBatchShrinkRequestTags {
	s.Key = &v
	return s
}

func (s *StopAppSessionBatchShrinkRequestTags) SetValue(v string) *StopAppSessionBatchShrinkRequestTags {
	s.Value = &v
	return s
}

type StopAppSessionBatchResponseBody struct {
	// 应用id
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 自定义会话id
	BatchId *string `json:"BatchId,omitempty" xml:"BatchId,omitempty"`
	// 平台会话id
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// 请求id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopAppSessionBatchResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopAppSessionBatchResponseBody) GoString() string {
	return s.String()
}

func (s *StopAppSessionBatchResponseBody) SetAppId(v string) *StopAppSessionBatchResponseBody {
	s.AppId = &v
	return s
}

func (s *StopAppSessionBatchResponseBody) SetBatchId(v string) *StopAppSessionBatchResponseBody {
	s.BatchId = &v
	return s
}

func (s *StopAppSessionBatchResponseBody) SetProjectId(v string) *StopAppSessionBatchResponseBody {
	s.ProjectId = &v
	return s
}

func (s *StopAppSessionBatchResponseBody) SetRequestId(v string) *StopAppSessionBatchResponseBody {
	s.RequestId = &v
	return s
}

type StopAppSessionBatchResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StopAppSessionBatchResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopAppSessionBatchResponse) String() string {
	return tea.Prettify(s)
}

func (s StopAppSessionBatchResponse) GoString() string {
	return s.String()
}

func (s *StopAppSessionBatchResponse) SetHeaders(v map[string]*string) *StopAppSessionBatchResponse {
	s.Headers = v
	return s
}

func (s *StopAppSessionBatchResponse) SetStatusCode(v int32) *StopAppSessionBatchResponse {
	s.StatusCode = &v
	return s
}

func (s *StopAppSessionBatchResponse) SetBody(v *StopAppSessionBatchResponseBody) *StopAppSessionBatchResponse {
	s.Body = v
	return s
}

type VersionCheckSameNameServiceRequest struct {
	AppId          *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppVersionName *string `json:"AppVersionName,omitempty" xml:"AppVersionName,omitempty"`
}

func (s VersionCheckSameNameServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s VersionCheckSameNameServiceRequest) GoString() string {
	return s.String()
}

func (s *VersionCheckSameNameServiceRequest) SetAppId(v string) *VersionCheckSameNameServiceRequest {
	s.AppId = &v
	return s
}

func (s *VersionCheckSameNameServiceRequest) SetAppVersionName(v string) *VersionCheckSameNameServiceRequest {
	s.AppVersionName = &v
	return s
}

type VersionCheckSameNameServiceResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s VersionCheckSameNameServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s VersionCheckSameNameServiceResponseBody) GoString() string {
	return s.String()
}

func (s *VersionCheckSameNameServiceResponseBody) SetCode(v string) *VersionCheckSameNameServiceResponseBody {
	s.Code = &v
	return s
}

func (s *VersionCheckSameNameServiceResponseBody) SetData(v string) *VersionCheckSameNameServiceResponseBody {
	s.Data = &v
	return s
}

func (s *VersionCheckSameNameServiceResponseBody) SetMessage(v string) *VersionCheckSameNameServiceResponseBody {
	s.Message = &v
	return s
}

func (s *VersionCheckSameNameServiceResponseBody) SetRequestId(v string) *VersionCheckSameNameServiceResponseBody {
	s.RequestId = &v
	return s
}

type VersionCheckSameNameServiceResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *VersionCheckSameNameServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s VersionCheckSameNameServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s VersionCheckSameNameServiceResponse) GoString() string {
	return s.String()
}

func (s *VersionCheckSameNameServiceResponse) SetHeaders(v map[string]*string) *VersionCheckSameNameServiceResponse {
	s.Headers = v
	return s
}

func (s *VersionCheckSameNameServiceResponse) SetStatusCode(v int32) *VersionCheckSameNameServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *VersionCheckSameNameServiceResponse) SetBody(v *VersionCheckSameNameServiceResponseBody) *VersionCheckSameNameServiceResponse {
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

func (client *Client) CancelReserveTaskWithOptions(request *CancelReserveTaskRequest, runtime *util.RuntimeOptions) (_result *CancelReserveTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		body["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CancelReserveTask"),
		Version:     tea.String("2021-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CancelReserveTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CancelReserveTask(request *CancelReserveTaskRequest) (_result *CancelReserveTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CancelReserveTaskResponse{}
	_body, _err := client.CancelReserveTaskWithOptions(request, runtime)
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

func (client *Client) CreateAppSessionWithOptions(request *CreateAppSessionRequest, runtime *util.RuntimeOptions) (_result *CreateAppSessionResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.EnablePostpaid)) {
		query["EnablePostpaid"] = request.EnablePostpaid
	}

	if !tea.BoolValue(util.IsUnset(request.StartParameters)) {
		query["StartParameters"] = request.StartParameters
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

func (client *Client) CreateAppSessionBatchSyncWithOptions(tmpReq *CreateAppSessionBatchSyncRequest, runtime *util.RuntimeOptions) (_result *CreateAppSessionBatchSyncResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateAppSessionBatchSyncShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.AppInfos)) {
		request.AppInfosShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AppInfos, tea.String("AppInfos"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppInfosShrink)) {
		query["AppInfos"] = request.AppInfosShrink
	}

	if !tea.BoolValue(util.IsUnset(request.BatchId)) {
		query["BatchId"] = request.BatchId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAppSessionBatchSync"),
		Version:     tea.String("2021-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAppSessionBatchSyncResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateAppSessionBatchSync(request *CreateAppSessionBatchSyncRequest) (_result *CreateAppSessionBatchSyncResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAppSessionBatchSyncResponse{}
	_body, _err := client.CreateAppSessionBatchSyncWithOptions(request, runtime)
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

	if !tea.BoolValue(util.IsUnset(request.StartParameters)) {
		query["StartParameters"] = request.StartParameters
	}

	if !tea.BoolValue(util.IsUnset(request.SystemInfo)) {
		query["SystemInfo"] = request.SystemInfo
	}

	if !tea.BoolValue(util.IsUnset(request.Tags)) {
		query["Tags"] = request.Tags
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
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.AppVersion)) {
		body["AppVersion"] = request.AppVersion
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DistrictId)) {
		body["DistrictId"] = request.DistrictId
	}

	if !tea.BoolValue(util.IsUnset(request.ExpectResourceReadyTime)) {
		body["ExpectResourceReadyTime"] = request.ExpectResourceReadyTime
	}

	if !tea.BoolValue(util.IsUnset(request.ExpectSessionCapacity)) {
		body["ExpectSessionCapacity"] = request.ExpectSessionCapacity
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
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

func (client *Client) GetAppCcuWithOptions(request *GetAppCcuRequest, runtime *util.RuntimeOptions) (_result *GetAppCcuResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAppCcu"),
		Version:     tea.String("2021-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAppCcuResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAppCcu(request *GetAppCcuRequest) (_result *GetAppCcuResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAppCcuResponse{}
	_body, _err := client.GetAppCcuWithOptions(request, runtime)
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

func (client *Client) GetCapacityWithOptions(request *GetCapacityRequest, runtime *util.RuntimeOptions) (_result *GetCapacityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.AppVersion)) {
		body["AppVersion"] = request.AppVersion
	}

	if !tea.BoolValue(util.IsUnset(request.DistrictId)) {
		body["DistrictId"] = request.DistrictId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		body["PageNum"] = request.PageNum
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
		Action:      tea.String("GetCapacity"),
		Version:     tea.String("2021-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCapacityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetCapacity(request *GetCapacityRequest) (_result *GetCapacityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetCapacityResponse{}
	_body, _err := client.GetCapacityWithOptions(request, runtime)
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

func (client *Client) GetReserveTaskDetailWithOptions(request *GetReserveTaskDetailRequest, runtime *util.RuntimeOptions) (_result *GetReserveTaskDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		body["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetReserveTaskDetail"),
		Version:     tea.String("2021-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetReserveTaskDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetReserveTaskDetail(request *GetReserveTaskDetailRequest) (_result *GetReserveTaskDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetReserveTaskDetailResponse{}
	_body, _err := client.GetReserveTaskDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetResourcePublicIPsWithOptions(request *GetResourcePublicIPsRequest, runtime *util.RuntimeOptions) (_result *GetResourcePublicIPsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		body["PageNum"] = request.PageNum
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
		Action:      tea.String("GetResourcePublicIPs"),
		Version:     tea.String("2021-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetResourcePublicIPsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetResourcePublicIPs(request *GetResourcePublicIPsRequest) (_result *GetResourcePublicIPsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetResourcePublicIPsResponse{}
	_body, _err := client.GetResourcePublicIPsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTenantWithOptions(runtime *util.RuntimeOptions) (_result *GetTenantResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("GetTenant"),
		Version:     tea.String("2021-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTenantResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTenant() (_result *GetTenantResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTenantResponse{}
	_body, _err := client.GetTenantWithOptions(runtime)
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
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
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

func (client *Client) ReleaseCapacityWithOptions(request *ReleaseCapacityRequest, runtime *util.RuntimeOptions) (_result *ReleaseCapacityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.AppVersion)) {
		body["AppVersion"] = request.AppVersion
	}

	if !tea.BoolValue(util.IsUnset(request.DistrictId)) {
		body["DistrictId"] = request.DistrictId
	}

	if !tea.BoolValue(util.IsUnset(request.ExpectReleaseSessionCapacity)) {
		body["ExpectReleaseSessionCapacity"] = request.ExpectReleaseSessionCapacity
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ReleaseCapacity"),
		Version:     tea.String("2021-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ReleaseCapacityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReleaseCapacity(request *ReleaseCapacityRequest) (_result *ReleaseCapacityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReleaseCapacityResponse{}
	_body, _err := client.ReleaseCapacityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopAppSessionWithOptions(tmpReq *StopAppSessionRequest, runtime *util.RuntimeOptions) (_result *StopAppSessionResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &StopAppSessionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.StopParam)) {
		request.StopParamShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.StopParam, tea.String("StopParam"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CustomSessionId)) {
		query["CustomSessionId"] = request.CustomSessionId
	}

	if !tea.BoolValue(util.IsUnset(request.PlatformSessionId)) {
		query["PlatformSessionId"] = request.PlatformSessionId
	}

	if !tea.BoolValue(util.IsUnset(request.StopParamShrink)) {
		query["StopParam"] = request.StopParamShrink
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

func (client *Client) StopAppSessionBatchWithOptions(tmpReq *StopAppSessionBatchRequest, runtime *util.RuntimeOptions) (_result *StopAppSessionBatchResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &StopAppSessionBatchShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.StopParam)) {
		request.StopParamShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.StopParam, tea.String("StopParam"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.AppVersion)) {
		query["AppVersion"] = request.AppVersion
	}

	if !tea.BoolValue(util.IsUnset(request.BatchId)) {
		query["BatchId"] = request.BatchId
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.StopParamShrink)) {
		query["StopParam"] = request.StopParamShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Tags)) {
		query["Tags"] = request.Tags
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StopAppSessionBatch"),
		Version:     tea.String("2021-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StopAppSessionBatchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopAppSessionBatch(request *StopAppSessionBatchRequest) (_result *StopAppSessionBatchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopAppSessionBatchResponse{}
	_body, _err := client.StopAppSessionBatchWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) VersionCheckSameNameServiceWithOptions(request *VersionCheckSameNameServiceRequest, runtime *util.RuntimeOptions) (_result *VersionCheckSameNameServiceResponse, _err error) {
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
		Action:      tea.String("VersionCheckSameNameService"),
		Version:     tea.String("2021-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &VersionCheckSameNameServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) VersionCheckSameNameService(request *VersionCheckSameNameServiceRequest) (_result *VersionCheckSameNameServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &VersionCheckSameNameServiceResponse{}
	_body, _err := client.VersionCheckSameNameServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
