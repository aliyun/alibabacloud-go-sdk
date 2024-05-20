// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type BatchCheckSessionRequest struct {
	// This parameter is required.
	Records []*BatchCheckSessionRequestRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
}

func (s BatchCheckSessionRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchCheckSessionRequest) GoString() string {
	return s.String()
}

func (s *BatchCheckSessionRequest) SetRecords(v []*BatchCheckSessionRequestRecords) *BatchCheckSessionRequest {
	s.Records = v
	return s
}

type BatchCheckSessionRequestRecords struct {
	CustomSessionId   *string                `json:"CustomSessionId,omitempty" xml:"CustomSessionId,omitempty"`
	PlatformSessionId *string                `json:"PlatformSessionId,omitempty" xml:"PlatformSessionId,omitempty"`
	ReferenceInfo     map[string]interface{} `json:"ReferenceInfo,omitempty" xml:"ReferenceInfo,omitempty"`
	// This parameter is required.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s BatchCheckSessionRequestRecords) String() string {
	return tea.Prettify(s)
}

func (s BatchCheckSessionRequestRecords) GoString() string {
	return s.String()
}

func (s *BatchCheckSessionRequestRecords) SetCustomSessionId(v string) *BatchCheckSessionRequestRecords {
	s.CustomSessionId = &v
	return s
}

func (s *BatchCheckSessionRequestRecords) SetPlatformSessionId(v string) *BatchCheckSessionRequestRecords {
	s.PlatformSessionId = &v
	return s
}

func (s *BatchCheckSessionRequestRecords) SetReferenceInfo(v map[string]interface{}) *BatchCheckSessionRequestRecords {
	s.ReferenceInfo = v
	return s
}

func (s *BatchCheckSessionRequestRecords) SetType(v string) *BatchCheckSessionRequestRecords {
	s.Type = &v
	return s
}

type BatchCheckSessionShrinkRequest struct {
	// This parameter is required.
	RecordsShrink *string `json:"Records,omitempty" xml:"Records,omitempty"`
}

func (s BatchCheckSessionShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchCheckSessionShrinkRequest) GoString() string {
	return s.String()
}

func (s *BatchCheckSessionShrinkRequest) SetRecordsShrink(v string) *BatchCheckSessionShrinkRequest {
	s.RecordsShrink = &v
	return s
}

type BatchCheckSessionResponseBody struct {
	// example:
	//
	// 46329898-489C-4E63-9BA1-C1DA5C5D0986
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchCheckSessionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchCheckSessionResponseBody) GoString() string {
	return s.String()
}

func (s *BatchCheckSessionResponseBody) SetRequestId(v string) *BatchCheckSessionResponseBody {
	s.RequestId = &v
	return s
}

type BatchCheckSessionResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchCheckSessionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchCheckSessionResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchCheckSessionResponse) GoString() string {
	return s.String()
}

func (s *BatchCheckSessionResponse) SetHeaders(v map[string]*string) *BatchCheckSessionResponse {
	s.Headers = v
	return s
}

func (s *BatchCheckSessionResponse) SetStatusCode(v int32) *BatchCheckSessionResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchCheckSessionResponse) SetBody(v *BatchCheckSessionResponseBody) *BatchCheckSessionResponse {
	s.Body = v
	return s
}

type CancelReserveTaskRequest struct {
	// example:
	//
	// 2YEF0****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// b354****
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
	// example:
	//
	// 46329898-489C-4E63-9BA1-C1DA5C5D0986
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// b354****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelReserveTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	AdaptTarget *CreateAdaptationRequestAdaptTarget `json:"AdaptTarget,omitempty" xml:"AdaptTarget,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 1432****
	AppVersionId *string `json:"AppVersionId,omitempty" xml:"AppVersionId,omitempty"`
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
	// example:
	//
	// 30
	BitRate *int32 `json:"BitRate,omitempty" xml:"BitRate,omitempty"`
	// example:
	//
	// 30
	FrameRate *int32 `json:"FrameRate,omitempty" xml:"FrameRate,omitempty"`
	// example:
	//
	// 1080p
	Resolution *string `json:"Resolution,omitempty" xml:"Resolution,omitempty"`
	// example:
	//
	// /example/example.exe
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
	// This parameter is required.
	//
	// example:
	//
	// 1432****
	AppVersionId *string `json:"AppVersionId,omitempty" xml:"AppVersionId,omitempty"`
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
	// example:
	//
	// 5435****
	AdaptApplyId *int64 `json:"AdaptApplyId,omitempty" xml:"AdaptApplyId,omitempty"`
	// example:
	//
	// 46329898-489C-4E63-9BA1-C1DA5C5D0986
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAdaptationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// example
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// end_game
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// example:
	//
	// 47d0bd4d-8815-48a2-b783-6cbba89d****
	StreamingAppId *string `json:"StreamingAppId,omitempty" xml:"StreamingAppId,omitempty"`
	// example:
	//
	// self-have-streaming
	StreamingSolution *string `json:"StreamingSolution,omitempty" xml:"StreamingSolution,omitempty"`
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

func (s *CreateAppRequest) SetStreamingAppId(v string) *CreateAppRequest {
	s.StreamingAppId = &v
	return s
}

func (s *CreateAppRequest) SetStreamingSolution(v string) *CreateAppRequest {
	s.StreamingSolution = &v
	return s
}

type CreateAppResponseBody struct {
	// example:
	//
	// 4384****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// 46329898-489C-4E63-9BA1-C1DA5C5D0986
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
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// 适配文件ID。此功能灰度开放，如未约定使用请勿传入。
	//
	// example:
	//
	// 501716211209548966XXXX
	AdapterFileId *string `json:"AdapterFileId,omitempty" xml:"AdapterFileId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 13027XXXX
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// 35067XXXX
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	// example:
	//
	// 192.168.XXX.XXX
	ClientIp *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1ADE0XXXX
	CustomSessionId *string `json:"CustomSessionId,omitempty" xml:"CustomSessionId,omitempty"`
	// example:
	//
	// 2YEF0XXXX
	CustomUserId *string `json:"CustomUserId,omitempty" xml:"CustomUserId,omitempty"`
	// example:
	//
	// huadong
	DistrictId *string `json:"DistrictId,omitempty" xml:"DistrictId,omitempty"`
	// example:
	//
	// false
	EnablePostpaid *bool `json:"EnablePostpaid,omitempty" xml:"EnablePostpaid,omitempty"`
	// 项目ID。如果已将应用关联到项目，创建会话时需填写正确的项目ID。
	//
	// example:
	//
	// d9a8****
	ProjectId       *string                                   `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	StartParameters []*CreateAppSessionRequestStartParameters `json:"StartParameters,omitempty" xml:"StartParameters,omitempty" type:"Repeated"`
	SystemInfo      []*CreateAppSessionRequestSystemInfo      `json:"SystemInfo,omitempty" xml:"SystemInfo,omitempty" type:"Repeated"`
	// example:
	//
	// 1800
	Timeout *int64 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s CreateAppSessionRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAppSessionRequest) GoString() string {
	return s.String()
}

func (s *CreateAppSessionRequest) SetAdapterFileId(v string) *CreateAppSessionRequest {
	s.AdapterFileId = &v
	return s
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

func (s *CreateAppSessionRequest) SetDistrictId(v string) *CreateAppSessionRequest {
	s.DistrictId = &v
	return s
}

func (s *CreateAppSessionRequest) SetEnablePostpaid(v bool) *CreateAppSessionRequest {
	s.EnablePostpaid = &v
	return s
}

func (s *CreateAppSessionRequest) SetProjectId(v string) *CreateAppSessionRequest {
	s.ProjectId = &v
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
	// example:
	//
	// startArgument
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// exampleValue
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
	// example:
	//
	// utdid
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// OE0usD+APXXXX
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
	// example:
	//
	// 13027XXXX
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// 35067XXXX
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	// example:
	//
	// 1ADE0XXXX
	CustomSessionId *string `json:"CustomSessionId,omitempty" xml:"CustomSessionId,omitempty"`
	// example:
	//
	// 100XXXX
	PlatformSessionId *string `json:"PlatformSessionId,omitempty" xml:"PlatformSessionId,omitempty"`
	// example:
	//
	// 46329898-489C-4E63-9BA1-C1DA5C5D0986
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAppSessionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type CreateAppSessionBatchRequest struct {
	// This parameter is required.
	AppInfos []*CreateAppSessionBatchRequestAppInfos `json:"AppInfos,omitempty" xml:"AppInfos,omitempty" type:"Repeated"`
	// This parameter is required.
	CustomTaskId *string `json:"CustomTaskId,omitempty" xml:"CustomTaskId,omitempty"`
	// This parameter is required.
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s CreateAppSessionBatchRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAppSessionBatchRequest) GoString() string {
	return s.String()
}

func (s *CreateAppSessionBatchRequest) SetAppInfos(v []*CreateAppSessionBatchRequestAppInfos) *CreateAppSessionBatchRequest {
	s.AppInfos = v
	return s
}

func (s *CreateAppSessionBatchRequest) SetCustomTaskId(v string) *CreateAppSessionBatchRequest {
	s.CustomTaskId = &v
	return s
}

func (s *CreateAppSessionBatchRequest) SetTimeout(v int32) *CreateAppSessionBatchRequest {
	s.Timeout = &v
	return s
}

type CreateAppSessionBatchRequestAppInfos struct {
	AdapterFileId *string `json:"AdapterFileId,omitempty" xml:"AdapterFileId,omitempty"`
	// This parameter is required.
	AppId        *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppVersion   *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	ClientIp     *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	CustomUserId *string `json:"CustomUserId,omitempty" xml:"CustomUserId,omitempty"`
	// This parameter is required.
	//
	// if can be null:
	// false
	CustomerSessionId *string                                                `json:"CustomerSessionId,omitempty" xml:"CustomerSessionId,omitempty"`
	DatasetId         *string                                                `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	DistrictId        *string                                                `json:"DistrictId,omitempty" xml:"DistrictId,omitempty"`
	ProjectId         *string                                                `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	ResultStore       *CreateAppSessionBatchRequestAppInfosResultStore       `json:"ResultStore,omitempty" xml:"ResultStore,omitempty" type:"Struct"`
	StartParameters   []*CreateAppSessionBatchRequestAppInfosStartParameters `json:"StartParameters,omitempty" xml:"StartParameters,omitempty" type:"Repeated"`
	SystemInfo        []*CreateAppSessionBatchRequestAppInfosSystemInfo      `json:"SystemInfo,omitempty" xml:"SystemInfo,omitempty" type:"Repeated"`
}

func (s CreateAppSessionBatchRequestAppInfos) String() string {
	return tea.Prettify(s)
}

func (s CreateAppSessionBatchRequestAppInfos) GoString() string {
	return s.String()
}

func (s *CreateAppSessionBatchRequestAppInfos) SetAdapterFileId(v string) *CreateAppSessionBatchRequestAppInfos {
	s.AdapterFileId = &v
	return s
}

func (s *CreateAppSessionBatchRequestAppInfos) SetAppId(v string) *CreateAppSessionBatchRequestAppInfos {
	s.AppId = &v
	return s
}

func (s *CreateAppSessionBatchRequestAppInfos) SetAppVersion(v string) *CreateAppSessionBatchRequestAppInfos {
	s.AppVersion = &v
	return s
}

func (s *CreateAppSessionBatchRequestAppInfos) SetClientIp(v string) *CreateAppSessionBatchRequestAppInfos {
	s.ClientIp = &v
	return s
}

func (s *CreateAppSessionBatchRequestAppInfos) SetCustomUserId(v string) *CreateAppSessionBatchRequestAppInfos {
	s.CustomUserId = &v
	return s
}

func (s *CreateAppSessionBatchRequestAppInfos) SetCustomerSessionId(v string) *CreateAppSessionBatchRequestAppInfos {
	s.CustomerSessionId = &v
	return s
}

func (s *CreateAppSessionBatchRequestAppInfos) SetDatasetId(v string) *CreateAppSessionBatchRequestAppInfos {
	s.DatasetId = &v
	return s
}

func (s *CreateAppSessionBatchRequestAppInfos) SetDistrictId(v string) *CreateAppSessionBatchRequestAppInfos {
	s.DistrictId = &v
	return s
}

func (s *CreateAppSessionBatchRequestAppInfos) SetProjectId(v string) *CreateAppSessionBatchRequestAppInfos {
	s.ProjectId = &v
	return s
}

func (s *CreateAppSessionBatchRequestAppInfos) SetResultStore(v *CreateAppSessionBatchRequestAppInfosResultStore) *CreateAppSessionBatchRequestAppInfos {
	s.ResultStore = v
	return s
}

func (s *CreateAppSessionBatchRequestAppInfos) SetStartParameters(v []*CreateAppSessionBatchRequestAppInfosStartParameters) *CreateAppSessionBatchRequestAppInfos {
	s.StartParameters = v
	return s
}

func (s *CreateAppSessionBatchRequestAppInfos) SetSystemInfo(v []*CreateAppSessionBatchRequestAppInfosSystemInfo) *CreateAppSessionBatchRequestAppInfos {
	s.SystemInfo = v
	return s
}

type CreateAppSessionBatchRequestAppInfosResultStore struct {
	Need      *bool                                                       `json:"Need,omitempty" xml:"Need,omitempty"`
	StoreInfo []*CreateAppSessionBatchRequestAppInfosResultStoreStoreInfo `json:"StoreInfo,omitempty" xml:"StoreInfo,omitempty" type:"Repeated"`
	Type      *string                                                     `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateAppSessionBatchRequestAppInfosResultStore) String() string {
	return tea.Prettify(s)
}

func (s CreateAppSessionBatchRequestAppInfosResultStore) GoString() string {
	return s.String()
}

func (s *CreateAppSessionBatchRequestAppInfosResultStore) SetNeed(v bool) *CreateAppSessionBatchRequestAppInfosResultStore {
	s.Need = &v
	return s
}

func (s *CreateAppSessionBatchRequestAppInfosResultStore) SetStoreInfo(v []*CreateAppSessionBatchRequestAppInfosResultStoreStoreInfo) *CreateAppSessionBatchRequestAppInfosResultStore {
	s.StoreInfo = v
	return s
}

func (s *CreateAppSessionBatchRequestAppInfosResultStore) SetType(v string) *CreateAppSessionBatchRequestAppInfosResultStore {
	s.Type = &v
	return s
}

type CreateAppSessionBatchRequestAppInfosResultStoreStoreInfo struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateAppSessionBatchRequestAppInfosResultStoreStoreInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateAppSessionBatchRequestAppInfosResultStoreStoreInfo) GoString() string {
	return s.String()
}

func (s *CreateAppSessionBatchRequestAppInfosResultStoreStoreInfo) SetKey(v string) *CreateAppSessionBatchRequestAppInfosResultStoreStoreInfo {
	s.Key = &v
	return s
}

func (s *CreateAppSessionBatchRequestAppInfosResultStoreStoreInfo) SetValue(v string) *CreateAppSessionBatchRequestAppInfosResultStoreStoreInfo {
	s.Value = &v
	return s
}

type CreateAppSessionBatchRequestAppInfosStartParameters struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateAppSessionBatchRequestAppInfosStartParameters) String() string {
	return tea.Prettify(s)
}

func (s CreateAppSessionBatchRequestAppInfosStartParameters) GoString() string {
	return s.String()
}

func (s *CreateAppSessionBatchRequestAppInfosStartParameters) SetKey(v string) *CreateAppSessionBatchRequestAppInfosStartParameters {
	s.Key = &v
	return s
}

func (s *CreateAppSessionBatchRequestAppInfosStartParameters) SetValue(v string) *CreateAppSessionBatchRequestAppInfosStartParameters {
	s.Value = &v
	return s
}

type CreateAppSessionBatchRequestAppInfosSystemInfo struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateAppSessionBatchRequestAppInfosSystemInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateAppSessionBatchRequestAppInfosSystemInfo) GoString() string {
	return s.String()
}

func (s *CreateAppSessionBatchRequestAppInfosSystemInfo) SetKey(v string) *CreateAppSessionBatchRequestAppInfosSystemInfo {
	s.Key = &v
	return s
}

func (s *CreateAppSessionBatchRequestAppInfosSystemInfo) SetValue(v string) *CreateAppSessionBatchRequestAppInfosSystemInfo {
	s.Value = &v
	return s
}

type CreateAppSessionBatchShrinkRequest struct {
	// This parameter is required.
	AppInfosShrink *string `json:"AppInfos,omitempty" xml:"AppInfos,omitempty"`
	// This parameter is required.
	CustomTaskId *string `json:"CustomTaskId,omitempty" xml:"CustomTaskId,omitempty"`
	// This parameter is required.
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s CreateAppSessionBatchShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAppSessionBatchShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateAppSessionBatchShrinkRequest) SetAppInfosShrink(v string) *CreateAppSessionBatchShrinkRequest {
	s.AppInfosShrink = &v
	return s
}

func (s *CreateAppSessionBatchShrinkRequest) SetCustomTaskId(v string) *CreateAppSessionBatchShrinkRequest {
	s.CustomTaskId = &v
	return s
}

func (s *CreateAppSessionBatchShrinkRequest) SetTimeout(v int32) *CreateAppSessionBatchShrinkRequest {
	s.Timeout = &v
	return s
}

type CreateAppSessionBatchResponseBody struct {
	CustomTaskId   *string `json:"CustomTaskId,omitempty" xml:"CustomTaskId,omitempty"`
	PlatformTaskId *string `json:"PlatformTaskId,omitempty" xml:"PlatformTaskId,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAppSessionBatchResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAppSessionBatchResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAppSessionBatchResponseBody) SetCustomTaskId(v string) *CreateAppSessionBatchResponseBody {
	s.CustomTaskId = &v
	return s
}

func (s *CreateAppSessionBatchResponseBody) SetPlatformTaskId(v string) *CreateAppSessionBatchResponseBody {
	s.PlatformTaskId = &v
	return s
}

func (s *CreateAppSessionBatchResponseBody) SetRequestId(v string) *CreateAppSessionBatchResponseBody {
	s.RequestId = &v
	return s
}

type CreateAppSessionBatchResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAppSessionBatchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAppSessionBatchResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAppSessionBatchResponse) GoString() string {
	return s.String()
}

func (s *CreateAppSessionBatchResponse) SetHeaders(v map[string]*string) *CreateAppSessionBatchResponse {
	s.Headers = v
	return s
}

func (s *CreateAppSessionBatchResponse) SetStatusCode(v int32) *CreateAppSessionBatchResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAppSessionBatchResponse) SetBody(v *CreateAppSessionBatchResponseBody) *CreateAppSessionBatchResponse {
	s.Body = v
	return s
}

type CreateAppSessionBatchSyncRequest struct {
	// This parameter is required.
	AppInfos []*CreateAppSessionBatchSyncRequestAppInfos `json:"AppInfos,omitempty" xml:"AppInfos,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 6d4d****
	BatchId *string `json:"BatchId,omitempty" xml:"BatchId,omitempty"`
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
	// example:
	//
	// 501716211209548966XXXX
	AdapterFileId *string `json:"AdapterFileId,omitempty" xml:"AdapterFileId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 13027****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// 35067****
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	// example:
	//
	// 192.168.XX.XX
	ClientIp *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	// example:
	//
	// 2YEF0****
	CustomUserId *string `json:"CustomUserId,omitempty" xml:"CustomUserId,omitempty"`
	// This parameter is required.
	//
	// if can be null:
	// false
	//
	// example:
	//
	// 1ADE0****
	CustomerSessionId *string `json:"CustomerSessionId,omitempty" xml:"CustomerSessionId,omitempty"`
	// example:
	//
	// huadong
	DistrictId *string                                               `json:"DistrictId,omitempty" xml:"DistrictId,omitempty"`
	MatchRules []*CreateAppSessionBatchSyncRequestAppInfosMatchRules `json:"MatchRules,omitempty" xml:"MatchRules,omitempty" type:"Repeated"`
	// example:
	//
	// d9a8****
	ProjectId       *string                                                    `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	StartParameters []*CreateAppSessionBatchSyncRequestAppInfosStartParameters `json:"StartParameters,omitempty" xml:"StartParameters,omitempty" type:"Repeated"`
	SystemInfo      []*CreateAppSessionBatchSyncRequestAppInfosSystemInfo      `json:"SystemInfo,omitempty" xml:"SystemInfo,omitempty" type:"Repeated"`
	Tags            []*CreateAppSessionBatchSyncRequestAppInfosTags            `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s CreateAppSessionBatchSyncRequestAppInfos) String() string {
	return tea.Prettify(s)
}

func (s CreateAppSessionBatchSyncRequestAppInfos) GoString() string {
	return s.String()
}

func (s *CreateAppSessionBatchSyncRequestAppInfos) SetAdapterFileId(v string) *CreateAppSessionBatchSyncRequestAppInfos {
	s.AdapterFileId = &v
	return s
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

func (s *CreateAppSessionBatchSyncRequestAppInfos) SetMatchRules(v []*CreateAppSessionBatchSyncRequestAppInfosMatchRules) *CreateAppSessionBatchSyncRequestAppInfos {
	s.MatchRules = v
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

type CreateAppSessionBatchSyncRequestAppInfosMatchRules struct {
	// example:
	//
	// component
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// in
	Type   *string   `json:"Type,omitempty" xml:"Type,omitempty"`
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s CreateAppSessionBatchSyncRequestAppInfosMatchRules) String() string {
	return tea.Prettify(s)
}

func (s CreateAppSessionBatchSyncRequestAppInfosMatchRules) GoString() string {
	return s.String()
}

func (s *CreateAppSessionBatchSyncRequestAppInfosMatchRules) SetKey(v string) *CreateAppSessionBatchSyncRequestAppInfosMatchRules {
	s.Key = &v
	return s
}

func (s *CreateAppSessionBatchSyncRequestAppInfosMatchRules) SetType(v string) *CreateAppSessionBatchSyncRequestAppInfosMatchRules {
	s.Type = &v
	return s
}

func (s *CreateAppSessionBatchSyncRequestAppInfosMatchRules) SetValues(v []*string) *CreateAppSessionBatchSyncRequestAppInfosMatchRules {
	s.Values = v
	return s
}

type CreateAppSessionBatchSyncRequestAppInfosStartParameters struct {
	// example:
	//
	// startArgument
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// exampleValue
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
	// example:
	//
	// utdid
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// OE0usD+AP****
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
	// example:
	//
	// exampleTag
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// exampleValue
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
	// This parameter is required.
	AppInfosShrink *string `json:"AppInfos,omitempty" xml:"AppInfos,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 6d4d****
	BatchId *string `json:"BatchId,omitempty" xml:"BatchId,omitempty"`
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
	// example:
	//
	// 6d4d****
	BatchId    *string                                            `json:"BatchId,omitempty" xml:"BatchId,omitempty"`
	FailedList []*CreateAppSessionBatchSyncResponseBodyFailedList `json:"FailedList,omitempty" xml:"FailedList,omitempty" type:"Repeated"`
	// example:
	//
	// 46329898-489C-4E63-9BA1-C1DA5C5D0986
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
	// example:
	//
	// 100****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// 1ADE0****
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
	// example:
	//
	// App type not support.
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// 400
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
	// example:
	//
	// 13027****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// 35067****
	AppVersion *string                                                 `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	BizInfo    *CreateAppSessionBatchSyncResponseBodyResultListBizInfo `json:"BizInfo,omitempty" xml:"BizInfo,omitempty" type:"Struct"`
	// example:
	//
	// 1ADE0****
	CustomSessionId *string `json:"CustomSessionId,omitempty" xml:"CustomSessionId,omitempty"`
	// example:
	//
	// 100****
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
	// example:
	//
	// authToken
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
	// example:
	//
	// 127.0.X.X
	AccessHost *string `json:"AccessHost,omitempty" xml:"AccessHost,omitempty"`
	// example:
	//
	// 8080
	AccessPort *string `json:"AccessPort,omitempty" xml:"AccessPort,omitempty"`
	// example:
	//
	// huadong
	DistrictId *string `json:"DistrictId,omitempty" xml:"DistrictId,omitempty"`
	// example:
	//
	// bgp
	Isp *string `json:"Isp,omitempty" xml:"Isp,omitempty"`
	// example:
	//
	// exampleName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// Native
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
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
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAppSessionBatchSyncResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// example:
	//
	// 501716211209548966XXXX
	AdapterFileId *string `json:"AdapterFileId,omitempty" xml:"AdapterFileId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 13027****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// 35067****
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	// example:
	//
	// 192.168.XX.XX
	ClientIp *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1ADE0****
	CustomSessionId *string `json:"CustomSessionId,omitempty" xml:"CustomSessionId,omitempty"`
	// example:
	//
	// 2YEF0****
	CustomUserId *string `json:"CustomUserId,omitempty" xml:"CustomUserId,omitempty"`
	// example:
	//
	// huadong
	DistrictId *string                                  `json:"DistrictId,omitempty" xml:"DistrictId,omitempty"`
	MatchRules []*CreateAppSessionSyncRequestMatchRules `json:"MatchRules,omitempty" xml:"MatchRules,omitempty" type:"Repeated"`
	// example:
	//
	// d9a8****
	ProjectId       *string                                       `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	StartParameters []*CreateAppSessionSyncRequestStartParameters `json:"StartParameters,omitempty" xml:"StartParameters,omitempty" type:"Repeated"`
	SystemInfo      []*CreateAppSessionSyncRequestSystemInfo      `json:"SystemInfo,omitempty" xml:"SystemInfo,omitempty" type:"Repeated"`
	Tags            []*CreateAppSessionSyncRequestTags            `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s CreateAppSessionSyncRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAppSessionSyncRequest) GoString() string {
	return s.String()
}

func (s *CreateAppSessionSyncRequest) SetAdapterFileId(v string) *CreateAppSessionSyncRequest {
	s.AdapterFileId = &v
	return s
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

func (s *CreateAppSessionSyncRequest) SetMatchRules(v []*CreateAppSessionSyncRequestMatchRules) *CreateAppSessionSyncRequest {
	s.MatchRules = v
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

type CreateAppSessionSyncRequestMatchRules struct {
	// example:
	//
	// component
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// in
	Type   *string   `json:"Type,omitempty" xml:"Type,omitempty"`
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s CreateAppSessionSyncRequestMatchRules) String() string {
	return tea.Prettify(s)
}

func (s CreateAppSessionSyncRequestMatchRules) GoString() string {
	return s.String()
}

func (s *CreateAppSessionSyncRequestMatchRules) SetKey(v string) *CreateAppSessionSyncRequestMatchRules {
	s.Key = &v
	return s
}

func (s *CreateAppSessionSyncRequestMatchRules) SetType(v string) *CreateAppSessionSyncRequestMatchRules {
	s.Type = &v
	return s
}

func (s *CreateAppSessionSyncRequestMatchRules) SetValues(v []*string) *CreateAppSessionSyncRequestMatchRules {
	s.Values = v
	return s
}

type CreateAppSessionSyncRequestStartParameters struct {
	// example:
	//
	// startArgument
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// exampleValue
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
	// example:
	//
	// utdid
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// OE0usD+AP****
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
	// example:
	//
	// exampleTag
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// exampleValue
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
	// example:
	//
	// 13027****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// 35067****
	AppVersion *string                                  `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	BizInfo    *CreateAppSessionSyncResponseBodyBizInfo `json:"BizInfo,omitempty" xml:"BizInfo,omitempty" type:"Struct"`
	// example:
	//
	// 1ADE0****
	CustomSessionId *string `json:"CustomSessionId,omitempty" xml:"CustomSessionId,omitempty"`
	// example:
	//
	// 100****
	PlatformSessionId *string `json:"PlatformSessionId,omitempty" xml:"PlatformSessionId,omitempty"`
	// example:
	//
	// 46329898-489C-4E63-9BA1-C1DA5C5D0986
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
	// example:
	//
	// authToken
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
	// example:
	//
	// 127.0.X.X
	AccessHost *string `json:"AccessHost,omitempty" xml:"AccessHost,omitempty"`
	// example:
	//
	// 8080
	AccessPort *string `json:"AccessPort,omitempty" xml:"AccessPort,omitempty"`
	// example:
	//
	// huabei
	DistrictId *string `json:"DistrictId,omitempty" xml:"DistrictId,omitempty"`
	// example:
	//
	// BGP
	Isp *string `json:"Isp,omitempty" xml:"Isp,omitempty"`
	// example:
	//
	// exampleName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// Native
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAppSessionSyncResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// 4384****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// exampleVersion
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
	// example:
	//
	// 1432****
	AppVersionId *string `json:"AppVersionId,omitempty" xml:"AppVersionId,omitempty"`
	// example:
	//
	// 46329898-489C-4E63-9BA1-C1DA5C5D0986
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAppVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// 13027****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// 35067****
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	// example:
	//
	// 2YEF0****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// huadong
	DistrictId *string `json:"DistrictId,omitempty" xml:"DistrictId,omitempty"`
	// Deprecated
	//
	// example:
	//
	// 2022-02-02 22:22:22
	ExpectResourceReadyTime *string `json:"ExpectResourceReadyTime,omitempty" xml:"ExpectResourceReadyTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 100
	ExpectSessionCapacity *int32 `json:"ExpectSessionCapacity,omitempty" xml:"ExpectSessionCapacity,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// d9a8****
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
	// example:
	//
	// 10000
	CurrMaxAllocatableSessionCapacity *int32 `json:"CurrMaxAllocatableSessionCapacity,omitempty" xml:"CurrMaxAllocatableSessionCapacity,omitempty"`
	// example:
	//
	// 46329898-489C-4E63-9BA1-C1DA5C5D0986
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// b354****
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
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCapacityReservationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type DeleteAppRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 4384****
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
	// example:
	//
	// 4384****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// 46329898-489C-4E63-9BA1-C1DA5C5D0986
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
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// 1432****
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
	// example:
	//
	// 1432****
	AppVersionId *string `json:"AppVersionId,omitempty" xml:"AppVersionId,omitempty"`
	// example:
	//
	// 46329898-489C-4E63-9BA1-C1DA5C5D0986
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAppVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// example:
	//
	// 1432****
	AdaptApplyId *int64 `json:"AdaptApplyId,omitempty" xml:"AdaptApplyId,omitempty"`
	// example:
	//
	// 5435****
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
	// example:
	//
	// 5435****
	AdaptApplyId *int64                                `json:"AdaptApplyId,omitempty" xml:"AdaptApplyId,omitempty"`
	AdaptTarget  *GetAdaptationResponseBodyAdaptTarget `json:"AdaptTarget,omitempty" xml:"AdaptTarget,omitempty" type:"Struct"`
	// example:
	//
	// 4384****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// 1432****
	AppVersionId *string `json:"AppVersionId,omitempty" xml:"AppVersionId,omitempty"`
	// example:
	//
	// 2022-04-06 02:00:00
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 2022-04-06 02:00:00
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// 46329898-489C-4E63-9BA1-C1DA5C5D0986
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// example:
	//
	// 30
	BitRate *int32 `json:"BitRate,omitempty" xml:"BitRate,omitempty"`
	// example:
	//
	// 30
	FrameRate *int32 `json:"FrameRate,omitempty" xml:"FrameRate,omitempty"`
	// example:
	//
	// 1080p
	Resolution *string `json:"Resolution,omitempty" xml:"Resolution,omitempty"`
	// example:
	//
	// /example/example.exe
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
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAdaptationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// 4384****
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
	// example:
	//
	// 4384****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// example
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// end_game
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// example:
	//
	// 2022-04-06 02:00:00
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 2022-04-06 02:00:00
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// 46329898-489C-4E63-9BA1-C1DA5C5D0986
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 47d0bd4d-8815-48a2-b783-6cbba89d****
	StreamingAppId *string `json:"StreamingAppId,omitempty" xml:"StreamingAppId,omitempty"`
	// example:
	//
	// self-have-streaming
	StreamingSolution *string `json:"StreamingSolution,omitempty" xml:"StreamingSolution,omitempty"`
	// example:
	//
	// 10
	VersionAdaptNum *int64 `json:"VersionAdaptNum,omitempty" xml:"VersionAdaptNum,omitempty"`
	// example:
	//
	// 28
	VersionTotalNum *int64 `json:"VersionTotalNum,omitempty" xml:"VersionTotalNum,omitempty"`
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

func (s *GetAppResponseBody) SetStreamingAppId(v string) *GetAppResponseBody {
	s.StreamingAppId = &v
	return s
}

func (s *GetAppResponseBody) SetStreamingSolution(v string) *GetAppResponseBody {
	s.StreamingSolution = &v
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
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// example:
	//
	// 13027****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// 35067****
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	// example:
	//
	// d9a8****
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
	DetailList []*GetAppCcuResponseBodyDetailList `json:"DetailList,omitempty" xml:"DetailList,omitempty" type:"Repeated"`
	// example:
	//
	// 46329898-489C-4E63-9BA1-C1DA5C5D0986
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Timestamp *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s GetAppCcuResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAppCcuResponseBody) GoString() string {
	return s.String()
}

func (s *GetAppCcuResponseBody) SetDetailList(v []*GetAppCcuResponseBodyDetailList) *GetAppCcuResponseBody {
	s.DetailList = v
	return s
}

func (s *GetAppCcuResponseBody) SetRequestId(v string) *GetAppCcuResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAppCcuResponseBody) SetTimestamp(v string) *GetAppCcuResponseBody {
	s.Timestamp = &v
	return s
}

type GetAppCcuResponseBodyDetailList struct {
	// example:
	//
	// 13027****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// 35067****
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	// example:
	//
	// 100
	Ccu *string `json:"Ccu,omitempty" xml:"Ccu,omitempty"`
	// example:
	//
	// huadong
	DistrictId *string `json:"DistrictId,omitempty" xml:"DistrictId,omitempty"`
	// example:
	//
	// d9a8****
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s GetAppCcuResponseBodyDetailList) String() string {
	return tea.Prettify(s)
}

func (s GetAppCcuResponseBodyDetailList) GoString() string {
	return s.String()
}

func (s *GetAppCcuResponseBodyDetailList) SetAppId(v string) *GetAppCcuResponseBodyDetailList {
	s.AppId = &v
	return s
}

func (s *GetAppCcuResponseBodyDetailList) SetAppVersion(v string) *GetAppCcuResponseBodyDetailList {
	s.AppVersion = &v
	return s
}

func (s *GetAppCcuResponseBodyDetailList) SetCcu(v string) *GetAppCcuResponseBodyDetailList {
	s.Ccu = &v
	return s
}

func (s *GetAppCcuResponseBodyDetailList) SetDistrictId(v string) *GetAppCcuResponseBodyDetailList {
	s.DistrictId = &v
	return s
}

func (s *GetAppCcuResponseBodyDetailList) SetProjectId(v string) *GetAppCcuResponseBodyDetailList {
	s.ProjectId = &v
	return s
}

type GetAppCcuResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAppCcuResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// example:
	//
	// 1ADE0XXXX
	CustomSessionId *string `json:"CustomSessionId,omitempty" xml:"CustomSessionId,omitempty"`
	// example:
	//
	// 100XXXX
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
	// example:
	//
	// 13027XXXX
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// 35067XXXX
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	// 业务特定的信息，如会话启动/停止时间。
	BizInfo *GetAppSessionResponseBodyBizInfo `json:"BizInfo,omitempty" xml:"BizInfo,omitempty" type:"Struct"`
	// example:
	//
	// 1ADE0XXXX
	CustomSessionId *string `json:"CustomSessionId,omitempty" xml:"CustomSessionId,omitempty"`
	// example:
	//
	// 100XXXX
	PlatformSessionId *string `json:"PlatformSessionId,omitempty" xml:"PlatformSessionId,omitempty"`
	// example:
	//
	// 46329898-489C-4E63-9BA1-C1DA5C5D0986
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// running
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
	// 会话启动时间
	//
	// example:
	//
	// 2022-07-20 17:58:51
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// 会话停止时间
	//
	// example:
	//
	// 2022-07-20 17:58:57
	StopTime *string `json:"StopTime,omitempty" xml:"StopTime,omitempty"`
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
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAppSessionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// 1432****
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
	// example:
	//
	// 4384****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// 1432****
	AppVersionId *string `json:"AppVersionId,omitempty" xml:"AppVersionId,omitempty"`
	// example:
	//
	// exampleVersion
	AppVersionName *string `json:"AppVersionName,omitempty" xml:"AppVersionName,omitempty"`
	// example:
	//
	// file_uploading
	AppVersionStatus     *string `json:"AppVersionStatus,omitempty" xml:"AppVersionStatus,omitempty"`
	AppVersionStatusMemo *string `json:"AppVersionStatusMemo,omitempty" xml:"AppVersionStatusMemo,omitempty"`
	// example:
	//
	// 0.31
	ConsumeCu *float64 `json:"ConsumeCu,omitempty" xml:"ConsumeCu,omitempty"`
	// example:
	//
	// https://www.example.com/exampleFile.tar
	FileAddress *string `json:"FileAddress,omitempty" xml:"FileAddress,omitempty"`
	// example:
	//
	// 1024
	FileSize *int64 `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	// example:
	//
	// 2022-04-06 02:00:00
	FileUploadFinishTime *string `json:"FileUploadFinishTime,omitempty" xml:"FileUploadFinishTime,omitempty"`
	// example:
	//
	// local_file_upload
	FileUploadType *string `json:"FileUploadType,omitempty" xml:"FileUploadType,omitempty"`
	// example:
	//
	// 2022-04-06 02:00:00
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 2022-04-06 02:00:00
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// 46329898-489C-4E63-9BA1-C1DA5C5D0986
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAppVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// example:
	//
	// 13027****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// 35067****
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	// example:
	//
	// huadong
	DistrictId *string `json:"DistrictId,omitempty" xml:"DistrictId,omitempty"`
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// d9a8****
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
	Capacities []*GetCapacityResponseBodyCapacities `json:"Capacities,omitempty" xml:"Capacities,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 46329898-489C-4E63-9BA1-C1DA5C5D0986
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 100
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
	// example:
	//
	// 13027****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// 35067****
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	// example:
	//
	// huadong
	DistrictId *string `json:"DistrictId,omitempty" xml:"DistrictId,omitempty"`
	// example:
	//
	// d9a8****
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// 2
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

func (s *GetCapacityResponseBodyCapacities) SetProjectId(v string) *GetCapacityResponseBodyCapacities {
	s.ProjectId = &v
	return s
}

func (s *GetCapacityResponseBodyCapacities) SetSessionCapacity(v int32) *GetCapacityResponseBodyCapacities {
	s.SessionCapacity = &v
	return s
}

type GetCapacityResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCapacityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type GetReserveTaskDetailRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// b354****
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
	// example:
	//
	// 13027****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// 35067****
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	// example:
	//
	// 100
	CurrCompletedSessionCapacity *int32 `json:"CurrCompletedSessionCapacity,omitempty" xml:"CurrCompletedSessionCapacity,omitempty"`
	// example:
	//
	// huadong
	DistrictId *string `json:"DistrictId,omitempty" xml:"DistrictId,omitempty"`
	// example:
	//
	// 2022-02-02 22:22:22
	ExpectResourceReadyTime *string `json:"ExpectResourceReadyTime,omitempty" xml:"ExpectResourceReadyTime,omitempty"`
	// example:
	//
	// 100
	ExpectSessionCapacity *int32 `json:"ExpectSessionCapacity,omitempty" xml:"ExpectSessionCapacity,omitempty"`
	// example:
	//
	// d9a8****
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// 46329898-489C-4E63-9BA1-C1DA5C5D0986
	RequestId    *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResBatchList []*GetReserveTaskDetailResponseBodyResBatchList `json:"ResBatchList,omitempty" xml:"ResBatchList,omitempty" type:"Repeated"`
	// example:
	//
	// b354****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// created
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

func (s *GetReserveTaskDetailResponseBody) SetResBatchList(v []*GetReserveTaskDetailResponseBodyResBatchList) *GetReserveTaskDetailResponseBody {
	s.ResBatchList = v
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

type GetReserveTaskDetailResponseBodyResBatchList struct {
	// example:
	//
	// 726573XXXX
	ResBatchId *string `json:"ResBatchId,omitempty" xml:"ResBatchId,omitempty"`
	// example:
	//
	// resBatchId
	ResBatchTagName *string `json:"ResBatchTagName,omitempty" xml:"ResBatchTagName,omitempty"`
}

func (s GetReserveTaskDetailResponseBodyResBatchList) String() string {
	return tea.Prettify(s)
}

func (s GetReserveTaskDetailResponseBodyResBatchList) GoString() string {
	return s.String()
}

func (s *GetReserveTaskDetailResponseBodyResBatchList) SetResBatchId(v string) *GetReserveTaskDetailResponseBodyResBatchList {
	s.ResBatchId = &v
	return s
}

func (s *GetReserveTaskDetailResponseBodyResBatchList) SetResBatchTagName(v string) *GetReserveTaskDetailResponseBodyResBatchList {
	s.ResBatchTagName = &v
	return s
}

type GetReserveTaskDetailResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetReserveTaskDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// d9a8****
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
	IpList []*GetResourcePublicIPsResponseBodyIpList `json:"IpList,omitempty" xml:"IpList,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 1
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 46329898-489C-4E63-9BA1-C1DA5C5D0986
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 100
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
	// example:
	//
	// 127.0.X.X
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// example:
	//
	// d9a8****
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetResourcePublicIPsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type ListAppRequest struct {
	// example:
	//
	// example
	KeySearch *string `json:"KeySearch,omitempty" xml:"KeySearch,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
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
	Apps []*ListAppResponseBodyApps `json:"Apps,omitempty" xml:"Apps,omitempty" type:"Repeated"`
	// example:
	//
	// 46329898-489C-4E63-9BA1-C1DA5C5D0986
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 12
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
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
	// example:
	//
	// 4384****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// example
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// end_game
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// example:
	//
	// 2022-04-06 02:00:00
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 2022-04-06 02:00:00
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// 10
	VersionAdaptNum *int64 `json:"VersionAdaptNum,omitempty" xml:"VersionAdaptNum,omitempty"`
	// example:
	//
	// 28
	VersionTotalNum *int64 `json:"VersionTotalNum,omitempty" xml:"VersionTotalNum,omitempty"`
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
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// example:
	//
	// 13027XXXX
	AppId            *string   `json:"AppId,omitempty" xml:"AppId,omitempty"`
	CustomSessionIds []*string `json:"CustomSessionIds,omitempty" xml:"CustomSessionIds,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize           *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PlatformSessionIds []*string `json:"PlatformSessionIds,omitempty" xml:"PlatformSessionIds,omitempty" type:"Repeated"`
	// example:
	//
	// d9a8****
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
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

func (s *ListAppSessionsRequest) SetProjectId(v string) *ListAppSessionsRequest {
	s.ProjectId = &v
	return s
}

type ListAppSessionsResponseBody struct {
	AppSessions []*ListAppSessionsResponseBodyAppSessions `json:"AppSessions,omitempty" xml:"AppSessions,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 46329898-489C-4E63-9BA1-C1DA5C5D0986
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	// example:
	//
	// 13027XXXX
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// 35067XXXX
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	// 业务特定的信息，如会话启动/停止时间。
	BizInfo *ListAppSessionsResponseBodyAppSessionsBizInfo `json:"BizInfo,omitempty" xml:"BizInfo,omitempty" type:"Struct"`
	// example:
	//
	// 1ADE0XXXX
	CustomSessionId *string `json:"CustomSessionId,omitempty" xml:"CustomSessionId,omitempty"`
	// example:
	//
	// 100XXXX
	PlatformSessionId *string `json:"PlatformSessionId,omitempty" xml:"PlatformSessionId,omitempty"`
	// example:
	//
	// d9a8****
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// running
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

func (s *ListAppSessionsResponseBodyAppSessions) SetProjectId(v string) *ListAppSessionsResponseBodyAppSessions {
	s.ProjectId = &v
	return s
}

func (s *ListAppSessionsResponseBodyAppSessions) SetStatus(v string) *ListAppSessionsResponseBodyAppSessions {
	s.Status = &v
	return s
}

type ListAppSessionsResponseBodyAppSessionsBizInfo struct {
	// 会话启动时间
	//
	// example:
	//
	// 2022-07-20 17:58:51
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// 会话停止时间
	//
	// example:
	//
	// 2022-07-20 17:58:57
	StopTime *string `json:"StopTime,omitempty" xml:"StopTime,omitempty"`
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAppSessionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// 4384****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
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
	// example:
	//
	// 46329898-489C-4E63-9BA1-C1DA5C5D0986
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 16
	Total    *int64                                `json:"Total,omitempty" xml:"Total,omitempty"`
	Versions []*ListAppVersionResponseBodyVersions `json:"Versions,omitempty" xml:"Versions,omitempty" type:"Repeated"`
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
	// example:
	//
	// 4384****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// 1432****
	AppVersionId *string `json:"AppVersionId,omitempty" xml:"AppVersionId,omitempty"`
	// example:
	//
	// exampleVersion
	AppVersionName *string `json:"AppVersionName,omitempty" xml:"AppVersionName,omitempty"`
	// example:
	//
	// file_uploading
	AppVersionStatus     *string `json:"AppVersionStatus,omitempty" xml:"AppVersionStatus,omitempty"`
	AppVersionStatusMemo *string `json:"AppVersionStatusMemo,omitempty" xml:"AppVersionStatusMemo,omitempty"`
	// example:
	//
	// 0.31
	ConsumeCu *float64 `json:"ConsumeCu,omitempty" xml:"ConsumeCu,omitempty"`
	// example:
	//
	// https://www.example.com/exampleFile.tar
	FileAddress *string `json:"FileAddress,omitempty" xml:"FileAddress,omitempty"`
	// example:
	//
	// 1024
	FileSize *int64 `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	// example:
	//
	// 2022-04-06 02:00:00
	FileUploadFinishTime *string `json:"FileUploadFinishTime,omitempty" xml:"FileUploadFinishTime,omitempty"`
	// example:
	//
	// cloud_file_download
	FileUploadType *string `json:"FileUploadType,omitempty" xml:"FileUploadType,omitempty"`
	// example:
	//
	// 2022-04-06 02:00:00
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 2022-04-06 02:00:00
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAppVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type ListInstancesRequest struct {
	// example:
	//
	// huadong
	DistrictId *string   `json:"DistrictId,omitempty" xml:"DistrictId,omitempty"`
	InstanceId []*string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Repeated"`
	// example:
	//
	// gcs.r1c1m1.1xlarge
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// BEXZPF01W23U46598WVf
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 4820372607851300489003
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// CREATING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListInstancesRequest) SetDistrictId(v string) *ListInstancesRequest {
	s.DistrictId = &v
	return s
}

func (s *ListInstancesRequest) SetInstanceId(v []*string) *ListInstancesRequest {
	s.InstanceId = v
	return s
}

func (s *ListInstancesRequest) SetInstanceType(v string) *ListInstancesRequest {
	s.InstanceType = &v
	return s
}

func (s *ListInstancesRequest) SetMaxResults(v int32) *ListInstancesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListInstancesRequest) SetNextToken(v string) *ListInstancesRequest {
	s.NextToken = &v
	return s
}

func (s *ListInstancesRequest) SetProjectId(v string) *ListInstancesRequest {
	s.ProjectId = &v
	return s
}

func (s *ListInstancesRequest) SetStatus(v string) *ListInstancesRequest {
	s.Status = &v
	return s
}

type ListInstancesResponseBody struct {
	Instances []*ListInstancesResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	// example:
	//
	// 50
	MaxResults *string `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// BEXZPF01W23U46598WVf
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 46329898-489C-4E63-9BA1-C1DA5C5D0986
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBody) SetInstances(v []*ListInstancesResponseBodyInstances) *ListInstancesResponseBody {
	s.Instances = v
	return s
}

func (s *ListInstancesResponseBody) SetMaxResults(v string) *ListInstancesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListInstancesResponseBody) SetNextToken(v string) *ListInstancesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListInstancesResponseBody) SetRequestId(v string) *ListInstancesResponseBody {
	s.RequestId = &v
	return s
}

type ListInstancesResponseBodyInstances struct {
	// example:
	//
	// 2023-12-13T11:12:11Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// example:
	//
	// huadong
	DistrictId *string `json:"DistrictId,omitempty" xml:"DistrictId,omitempty"`
	// example:
	//
	// gcs-bmt0kbn7e013aedg9
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// gcs.r1c1m1.1xlarge
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// example:
	//
	// 4820372607851300489003
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// CREATING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListInstancesResponseBodyInstances) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyInstances) SetCreationTime(v string) *ListInstancesResponseBodyInstances {
	s.CreationTime = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetDistrictId(v string) *ListInstancesResponseBodyInstances {
	s.DistrictId = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetInstanceId(v string) *ListInstancesResponseBodyInstances {
	s.InstanceId = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetInstanceType(v string) *ListInstancesResponseBodyInstances {
	s.InstanceType = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetProjectId(v string) *ListInstancesResponseBodyInstances {
	s.ProjectId = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetStatus(v string) *ListInstancesResponseBodyInstances {
	s.Status = &v
	return s
}

type ListInstancesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListInstancesResponse) SetHeaders(v map[string]*string) *ListInstancesResponse {
	s.Headers = v
	return s
}

func (s *ListInstancesResponse) SetStatusCode(v int32) *ListInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInstancesResponse) SetBody(v *ListInstancesResponseBody) *ListInstancesResponse {
	s.Body = v
	return s
}

type ModifyAppRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 4384****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// example
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
	// example:
	//
	// 4384****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// 46329898-489C-4E63-9BA1-C1DA5C5D0986
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
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// 1432****
	AppVersionId *string `json:"AppVersionId,omitempty" xml:"AppVersionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// exampleVersion
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
	// example:
	//
	// 1432****
	AppVersionId *string `json:"AppVersionId,omitempty" xml:"AppVersionId,omitempty"`
	// example:
	//
	// 46329898-489C-4E63-9BA1-C1DA5C5D0986
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyAppVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type ReleaseCapacityRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 13027****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// 35067****
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// huadong
	DistrictId *string `json:"DistrictId,omitempty" xml:"DistrictId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 100
	ExpectReleaseSessionCapacity *int32 `json:"ExpectReleaseSessionCapacity,omitempty" xml:"ExpectReleaseSessionCapacity,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// d9a8****
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
	// example:
	//
	// 46329898-489C-4E63-9BA1-C1DA5C5D0986
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// b354****
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReleaseCapacityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type ReleaseCapacityByBatchRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 726573XXXX
	ResBatchId *string `json:"ResBatchId,omitempty" xml:"ResBatchId,omitempty"`
}

func (s ReleaseCapacityByBatchRequest) String() string {
	return tea.Prettify(s)
}

func (s ReleaseCapacityByBatchRequest) GoString() string {
	return s.String()
}

func (s *ReleaseCapacityByBatchRequest) SetResBatchId(v string) *ReleaseCapacityByBatchRequest {
	s.ResBatchId = &v
	return s
}

type ReleaseCapacityByBatchResponseBody struct {
	// example:
	//
	// 46329898-489C-4E63-9BA1-C1DA5C5D0986
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReleaseCapacityByBatchResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReleaseCapacityByBatchResponseBody) GoString() string {
	return s.String()
}

func (s *ReleaseCapacityByBatchResponseBody) SetRequestId(v string) *ReleaseCapacityByBatchResponseBody {
	s.RequestId = &v
	return s
}

type ReleaseCapacityByBatchResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReleaseCapacityByBatchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReleaseCapacityByBatchResponse) String() string {
	return tea.Prettify(s)
}

func (s ReleaseCapacityByBatchResponse) GoString() string {
	return s.String()
}

func (s *ReleaseCapacityByBatchResponse) SetHeaders(v map[string]*string) *ReleaseCapacityByBatchResponse {
	s.Headers = v
	return s
}

func (s *ReleaseCapacityByBatchResponse) SetStatusCode(v int32) *ReleaseCapacityByBatchResponse {
	s.StatusCode = &v
	return s
}

func (s *ReleaseCapacityByBatchResponse) SetBody(v *ReleaseCapacityByBatchResponseBody) *ReleaseCapacityByBatchResponse {
	s.Body = v
	return s
}

type ReleaseInstancesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	Amount *int32 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// huadong
	DistrictId *string `json:"DistrictId,omitempty" xml:"DistrictId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// gcs.r1c1m1.1xlarge
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 4820372607851300489003
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s ReleaseInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s ReleaseInstancesRequest) GoString() string {
	return s.String()
}

func (s *ReleaseInstancesRequest) SetAmount(v int32) *ReleaseInstancesRequest {
	s.Amount = &v
	return s
}

func (s *ReleaseInstancesRequest) SetDistrictId(v string) *ReleaseInstancesRequest {
	s.DistrictId = &v
	return s
}

func (s *ReleaseInstancesRequest) SetInstanceType(v string) *ReleaseInstancesRequest {
	s.InstanceType = &v
	return s
}

func (s *ReleaseInstancesRequest) SetProjectId(v string) *ReleaseInstancesRequest {
	s.ProjectId = &v
	return s
}

type ReleaseInstancesResponseBody struct {
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// example:
	//
	// 46329898-489C-4E63-9BA1-C1DA5C5D0986
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReleaseInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReleaseInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ReleaseInstancesResponseBody) SetInstanceIds(v []*string) *ReleaseInstancesResponseBody {
	s.InstanceIds = v
	return s
}

func (s *ReleaseInstancesResponseBody) SetRequestId(v string) *ReleaseInstancesResponseBody {
	s.RequestId = &v
	return s
}

type ReleaseInstancesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReleaseInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReleaseInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s ReleaseInstancesResponse) GoString() string {
	return s.String()
}

func (s *ReleaseInstancesResponse) SetHeaders(v map[string]*string) *ReleaseInstancesResponse {
	s.Headers = v
	return s
}

func (s *ReleaseInstancesResponse) SetStatusCode(v int32) *ReleaseInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *ReleaseInstancesResponse) SetBody(v *ReleaseInstancesResponseBody) *ReleaseInstancesResponse {
	s.Body = v
	return s
}

type ReserveInstancesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	Amount *int32 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// huadong
	DistrictId *string `json:"DistrictId,omitempty" xml:"DistrictId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// gcs.r1c1m1.1xlarge
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 5469588382860444937003
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s ReserveInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s ReserveInstancesRequest) GoString() string {
	return s.String()
}

func (s *ReserveInstancesRequest) SetAmount(v int32) *ReserveInstancesRequest {
	s.Amount = &v
	return s
}

func (s *ReserveInstancesRequest) SetDistrictId(v string) *ReserveInstancesRequest {
	s.DistrictId = &v
	return s
}

func (s *ReserveInstancesRequest) SetInstanceType(v string) *ReserveInstancesRequest {
	s.InstanceType = &v
	return s
}

func (s *ReserveInstancesRequest) SetProjectId(v string) *ReserveInstancesRequest {
	s.ProjectId = &v
	return s
}

type ReserveInstancesResponseBody struct {
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// example:
	//
	// 46329898-489C-4E63-9BA1-C1DA5C5D0986
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReserveInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReserveInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ReserveInstancesResponseBody) SetInstanceIds(v []*string) *ReserveInstancesResponseBody {
	s.InstanceIds = v
	return s
}

func (s *ReserveInstancesResponseBody) SetRequestId(v string) *ReserveInstancesResponseBody {
	s.RequestId = &v
	return s
}

type ReserveInstancesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReserveInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReserveInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s ReserveInstancesResponse) GoString() string {
	return s.String()
}

func (s *ReserveInstancesResponse) SetHeaders(v map[string]*string) *ReserveInstancesResponse {
	s.Headers = v
	return s
}

func (s *ReserveInstancesResponse) SetStatusCode(v int32) *ReserveInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *ReserveInstancesResponse) SetBody(v *ReserveInstancesResponseBody) *ReserveInstancesResponse {
	s.Body = v
	return s
}

type SendBizCocChangeCallbackRequest struct {
	PlatformSessionId *string `json:"PlatformSessionId,omitempty" xml:"PlatformSessionId,omitempty"`
	Result            *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s SendBizCocChangeCallbackRequest) String() string {
	return tea.Prettify(s)
}

func (s SendBizCocChangeCallbackRequest) GoString() string {
	return s.String()
}

func (s *SendBizCocChangeCallbackRequest) SetPlatformSessionId(v string) *SendBizCocChangeCallbackRequest {
	s.PlatformSessionId = &v
	return s
}

func (s *SendBizCocChangeCallbackRequest) SetResult(v bool) *SendBizCocChangeCallbackRequest {
	s.Result = &v
	return s
}

type SendBizCocChangeCallbackResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SendBizCocChangeCallbackResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendBizCocChangeCallbackResponseBody) GoString() string {
	return s.String()
}

func (s *SendBizCocChangeCallbackResponseBody) SetRequestId(v string) *SendBizCocChangeCallbackResponseBody {
	s.RequestId = &v
	return s
}

type SendBizCocChangeCallbackResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendBizCocChangeCallbackResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendBizCocChangeCallbackResponse) String() string {
	return tea.Prettify(s)
}

func (s SendBizCocChangeCallbackResponse) GoString() string {
	return s.String()
}

func (s *SendBizCocChangeCallbackResponse) SetHeaders(v map[string]*string) *SendBizCocChangeCallbackResponse {
	s.Headers = v
	return s
}

func (s *SendBizCocChangeCallbackResponse) SetStatusCode(v int32) *SendBizCocChangeCallbackResponse {
	s.StatusCode = &v
	return s
}

func (s *SendBizCocChangeCallbackResponse) SetBody(v *SendBizCocChangeCallbackResponseBody) *SendBizCocChangeCallbackResponse {
	s.Body = v
	return s
}

type StopAppSessionRequest struct {
	// example:
	//
	// 1ADE0XXXX
	CustomSessionId *string `json:"CustomSessionId,omitempty" xml:"CustomSessionId,omitempty"`
	// example:
	//
	// 100XXXX
	PlatformSessionId *string `json:"PlatformSessionId,omitempty" xml:"PlatformSessionId,omitempty"`
	// 停止容器参数。此参数将透传到Agent。
	StopParam []*StopAppSessionRequestStopParam `json:"StopParam,omitempty" xml:"StopParam,omitempty" type:"Repeated"`
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
	// 目前支持的枚举值包括：
	//
	// - reason：停止原因。
	//
	// example:
	//
	// reason
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// key对应的取值。
	//
	// example:
	//
	// exampleValue
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
	// example:
	//
	// 1ADE0XXXX
	CustomSessionId *string `json:"CustomSessionId,omitempty" xml:"CustomSessionId,omitempty"`
	// example:
	//
	// 100XXXX
	PlatformSessionId *string `json:"PlatformSessionId,omitempty" xml:"PlatformSessionId,omitempty"`
	// 停止容器参数。此参数将透传到Agent。
	StopParamShrink *string `json:"StopParam,omitempty" xml:"StopParam,omitempty"`
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
	// example:
	//
	// 13027XXXX
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// 35067XXXX
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	// example:
	//
	// 1ADE0XXXX
	CustomSessionId *string `json:"CustomSessionId,omitempty" xml:"CustomSessionId,omitempty"`
	// example:
	//
	// 100XXXX
	PlatformSessionId *string `json:"PlatformSessionId,omitempty" xml:"PlatformSessionId,omitempty"`
	// example:
	//
	// 46329898-489C-4E63-9BA1-C1DA5C5D0986
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopAppSessionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// example:
	//
	// 13027****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// 35067****
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 6d4d****
	BatchId *string `json:"BatchId,omitempty" xml:"BatchId,omitempty"`
	// example:
	//
	// d9a8****
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
	// example:
	//
	// reason
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// exampleValue
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
	// example:
	//
	// exampleTag
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// exampleValue
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
	// example:
	//
	// 13027****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// 35067****
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 6d4d****
	BatchId *string `json:"BatchId,omitempty" xml:"BatchId,omitempty"`
	// example:
	//
	// d9a8****
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
	// example:
	//
	// exampleTag
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// exampleValue
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
	// example:
	//
	// 13027****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// 6d4d****
	BatchId *string `json:"BatchId,omitempty" xml:"BatchId,omitempty"`
	// example:
	//
	// d9a8****
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// 46329898-489C-4E63-9BA1-C1DA5C5D0986
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopAppSessionBatchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type UpdateSessionBizStatusRequest struct {
	BizStatus *string `json:"BizStatus,omitempty" xml:"BizStatus,omitempty"`
	// This parameter is required.
	PlatformSessionId *string `json:"PlatformSessionId,omitempty" xml:"PlatformSessionId,omitempty"`
}

func (s UpdateSessionBizStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateSessionBizStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateSessionBizStatusRequest) SetBizStatus(v string) *UpdateSessionBizStatusRequest {
	s.BizStatus = &v
	return s
}

func (s *UpdateSessionBizStatusRequest) SetPlatformSessionId(v string) *UpdateSessionBizStatusRequest {
	s.PlatformSessionId = &v
	return s
}

type UpdateSessionBizStatusResponseBody struct {
	Code      *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *UpdateSessionBizStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateSessionBizStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateSessionBizStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSessionBizStatusResponseBody) SetCode(v string) *UpdateSessionBizStatusResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateSessionBizStatusResponseBody) SetData(v *UpdateSessionBizStatusResponseBodyData) *UpdateSessionBizStatusResponseBody {
	s.Data = v
	return s
}

func (s *UpdateSessionBizStatusResponseBody) SetMessage(v string) *UpdateSessionBizStatusResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateSessionBizStatusResponseBody) SetRequestId(v string) *UpdateSessionBizStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSessionBizStatusResponseBody) SetSuccess(v bool) *UpdateSessionBizStatusResponseBody {
	s.Success = &v
	return s
}

type UpdateSessionBizStatusResponseBodyData struct {
	CustomSessionId   *string `json:"CustomSessionId,omitempty" xml:"CustomSessionId,omitempty"`
	PlatformSessionId *string `json:"PlatformSessionId,omitempty" xml:"PlatformSessionId,omitempty"`
}

func (s UpdateSessionBizStatusResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s UpdateSessionBizStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateSessionBizStatusResponseBodyData) SetCustomSessionId(v string) *UpdateSessionBizStatusResponseBodyData {
	s.CustomSessionId = &v
	return s
}

func (s *UpdateSessionBizStatusResponseBodyData) SetPlatformSessionId(v string) *UpdateSessionBizStatusResponseBodyData {
	s.PlatformSessionId = &v
	return s
}

type UpdateSessionBizStatusResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSessionBizStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSessionBizStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateSessionBizStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateSessionBizStatusResponse) SetHeaders(v map[string]*string) *UpdateSessionBizStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateSessionBizStatusResponse) SetStatusCode(v int32) *UpdateSessionBizStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSessionBizStatusResponse) SetBody(v *UpdateSessionBizStatusResponseBody) *UpdateSessionBizStatusResponse {
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

// Summary:
//
// 批量检查异常会话
//
// @param tmpReq - BatchCheckSessionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchCheckSessionResponse
func (client *Client) BatchCheckSessionWithOptions(tmpReq *BatchCheckSessionRequest, runtime *util.RuntimeOptions) (_result *BatchCheckSessionResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &BatchCheckSessionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Records)) {
		request.RecordsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Records, tea.String("Records"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RecordsShrink)) {
		query["Records"] = request.RecordsShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchCheckSession"),
		Version:     tea.String("2021-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchCheckSessionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量检查异常会话
//
// @param request - BatchCheckSessionRequest
//
// @return BatchCheckSessionResponse
func (client *Client) BatchCheckSession(request *BatchCheckSessionRequest) (_result *BatchCheckSessionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchCheckSessionResponse{}
	_body, _err := client.BatchCheckSessionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 取消 session 资源预定任务
//
// @param request - CancelReserveTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelReserveTaskResponse
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

// Summary:
//
// 取消 session 资源预定任务
//
// @param request - CancelReserveTaskRequest
//
// @return CancelReserveTaskResponse
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

// Summary:
//
// 提交适配请求
//
// @param tmpReq - CreateAdaptationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAdaptationResponse
func (client *Client) CreateAdaptationWithOptions(tmpReq *CreateAdaptationRequest, runtime *util.RuntimeOptions) (_result *CreateAdaptationResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateAdaptationShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.AdaptTarget)) {
		request.AdaptTargetShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AdaptTarget, tea.String("AdaptTarget"), tea.String("json"))
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

// Summary:
//
// 提交适配请求
//
// @param request - CreateAdaptationRequest
//
// @return CreateAdaptationResponse
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

// Summary:
//
// 应用创建服务
//
// @param request - CreateAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAppResponse
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

	if !tea.BoolValue(util.IsUnset(request.StreamingAppId)) {
		body["StreamingAppId"] = request.StreamingAppId
	}

	if !tea.BoolValue(util.IsUnset(request.StreamingSolution)) {
		body["StreamingSolution"] = request.StreamingSolution
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

// Summary:
//
// 应用创建服务
//
// @param request - CreateAppRequest
//
// @return CreateAppResponse
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

// Summary:
//
// 增加实时生产资源的相关字段
//
// @param request - CreateAppSessionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAppSessionResponse
func (client *Client) CreateAppSessionWithOptions(request *CreateAppSessionRequest, runtime *util.RuntimeOptions) (_result *CreateAppSessionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AdapterFileId)) {
		query["AdapterFileId"] = request.AdapterFileId
	}

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

	if !tea.BoolValue(util.IsUnset(request.EnablePostpaid)) {
		query["EnablePostpaid"] = request.EnablePostpaid
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

// Summary:
//
// 增加实时生产资源的相关字段
//
// @param request - CreateAppSessionRequest
//
// @return CreateAppSessionResponse
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

// Summary:
//
// 批量创建会话
//
// @param tmpReq - CreateAppSessionBatchRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAppSessionBatchResponse
func (client *Client) CreateAppSessionBatchWithOptions(tmpReq *CreateAppSessionBatchRequest, runtime *util.RuntimeOptions) (_result *CreateAppSessionBatchResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateAppSessionBatchShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.AppInfos)) {
		request.AppInfosShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AppInfos, tea.String("AppInfos"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppInfosShrink)) {
		query["AppInfos"] = request.AppInfosShrink
	}

	if !tea.BoolValue(util.IsUnset(request.CustomTaskId)) {
		query["CustomTaskId"] = request.CustomTaskId
	}

	if !tea.BoolValue(util.IsUnset(request.Timeout)) {
		query["Timeout"] = request.Timeout
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAppSessionBatch"),
		Version:     tea.String("2021-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAppSessionBatchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量创建会话
//
// @param request - CreateAppSessionBatchRequest
//
// @return CreateAppSessionBatchResponse
func (client *Client) CreateAppSessionBatch(request *CreateAppSessionBatchRequest) (_result *CreateAppSessionBatchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAppSessionBatchResponse{}
	_body, _err := client.CreateAppSessionBatchWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 同步批量创建多个会话
//
// @param tmpReq - CreateAppSessionBatchSyncRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAppSessionBatchSyncResponse
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

// Summary:
//
// 同步批量创建多个会话
//
// @param request - CreateAppSessionBatchSyncRequest
//
// @return CreateAppSessionBatchSyncResponse
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

// Summary:
//
// 同步创建会话
//
// @param request - CreateAppSessionSyncRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAppSessionSyncResponse
func (client *Client) CreateAppSessionSyncWithOptions(request *CreateAppSessionSyncRequest, runtime *util.RuntimeOptions) (_result *CreateAppSessionSyncResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AdapterFileId)) {
		query["AdapterFileId"] = request.AdapterFileId
	}

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

	if !tea.BoolValue(util.IsUnset(request.MatchRules)) {
		query["MatchRules"] = request.MatchRules
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

// Summary:
//
// 同步创建会话
//
// @param request - CreateAppSessionSyncRequest
//
// @return CreateAppSessionSyncResponse
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

// Summary:
//
// 应用版本创建服务
//
// @param request - CreateAppVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAppVersionResponse
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

// Summary:
//
// 应用版本创建服务
//
// @param request - CreateAppVersionRequest
//
// @return CreateAppVersionResponse
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

// Summary:
//
// 预定session资源
//
// @param request - CreateCapacityReservationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCapacityReservationResponse
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

// Summary:
//
// 预定session资源
//
// @param request - CreateCapacityReservationRequest
//
// @return CreateCapacityReservationResponse
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

// Summary:
//
// 应用删除接口
//
// @param request - DeleteAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAppResponse
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

// Summary:
//
// 应用删除接口
//
// @param request - DeleteAppRequest
//
// @return DeleteAppResponse
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

// Summary:
//
// 应用版本删除接口
//
// @param request - DeleteAppVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAppVersionResponse
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

// Summary:
//
// 应用版本删除接口
//
// @param request - DeleteAppVersionRequest
//
// @return DeleteAppVersionResponse
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

// Summary:
//
// 获取适配申请详情
//
// @param request - GetAdaptationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAdaptationResponse
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

// Summary:
//
// 获取适配申请详情
//
// @param request - GetAdaptationRequest
//
// @return GetAdaptationResponse
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

// Summary:
//
// 应用详情接口
//
// @param request - GetAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAppResponse
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

// Summary:
//
// 应用详情接口
//
// @param request - GetAppRequest
//
// @return GetAppResponse
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

// Summary:
//
// 查询会话并发数
//
// @param request - GetAppCcuRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAppCcuResponse
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

// Summary:
//
// 查询会话并发数
//
// @param request - GetAppCcuRequest
//
// @return GetAppCcuResponse
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

// Summary:
//
// 获取App会话信息
//
// @param request - GetAppSessionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAppSessionResponse
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

// Summary:
//
// 获取App会话信息
//
// @param request - GetAppSessionRequest
//
// @return GetAppSessionResponse
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

// Summary:
//
// 应用版本详情接口
//
// @param request - GetAppVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAppVersionResponse
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

// Summary:
//
// 应用版本详情接口
//
// @param request - GetAppVersionRequest
//
// @return GetAppVersionResponse
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

// Summary:
//
// 查询 session 会话容量信息
//
// @param request - GetCapacityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCapacityResponse
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

// Summary:
//
// 查询 session 会话容量信息
//
// @param request - GetCapacityRequest
//
// @return GetCapacityResponse
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

// Summary:
//
// 查询预定任务的详情信息
//
// @param request - GetReserveTaskDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetReserveTaskDetailResponse
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

// Summary:
//
// 查询预定任务的详情信息
//
// @param request - GetReserveTaskDetailRequest
//
// @return GetReserveTaskDetailResponse
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

// Summary:
//
// 查询公网ip
//
// @param request - GetResourcePublicIPsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResourcePublicIPsResponse
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

// Summary:
//
// 查询公网ip
//
// @param request - GetResourcePublicIPsRequest
//
// @return GetResourcePublicIPsResponse
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

// Summary:
//
// 应用列表接口
//
// @param request - ListAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAppResponse
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

// Summary:
//
// 应用列表接口
//
// @param request - ListAppRequest
//
// @return ListAppResponse
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

// Summary:
//
// 查询App会话
//
// @param request - ListAppSessionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAppSessionsResponse
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

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
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

// Summary:
//
// 查询App会话
//
// @param request - ListAppSessionsRequest
//
// @return ListAppSessionsResponse
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

// Summary:
//
// 应用版本列表接口
//
// @param request - ListAppVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAppVersionResponse
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

// Summary:
//
// 应用版本列表接口
//
// @param request - ListAppVersionRequest
//
// @return ListAppVersionResponse
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

// Summary:
//
// 查询GCS实例列表
//
// @param request - ListInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstancesResponse
func (client *Client) ListInstancesWithOptions(request *ListInstancesRequest, runtime *util.RuntimeOptions) (_result *ListInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListInstances"),
		Version:     tea.String("2021-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询GCS实例列表
//
// @param request - ListInstancesRequest
//
// @return ListInstancesResponse
func (client *Client) ListInstances(request *ListInstancesRequest) (_result *ListInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListInstancesResponse{}
	_body, _err := client.ListInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 应用修改服务
//
// @param request - ModifyAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAppResponse
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

// Summary:
//
// 应用修改服务
//
// @param request - ModifyAppRequest
//
// @return ModifyAppResponse
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

// Summary:
//
// 应用版本修改服务
//
// @param request - ModifyAppVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAppVersionResponse
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

// Summary:
//
// 应用版本修改服务
//
// @param request - ModifyAppVersionRequest
//
// @return ModifyAppVersionResponse
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

// Summary:
//
// 释放 session 资源预定的资源
//
// @param request - ReleaseCapacityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReleaseCapacityResponse
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

// Summary:
//
// 释放 session 资源预定的资源
//
// @param request - ReleaseCapacityRequest
//
// @return ReleaseCapacityResponse
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

// Summary:
//
// 根据资源批次号释放 session 资源预定的资源
//
// @param request - ReleaseCapacityByBatchRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReleaseCapacityByBatchResponse
func (client *Client) ReleaseCapacityByBatchWithOptions(request *ReleaseCapacityByBatchRequest, runtime *util.RuntimeOptions) (_result *ReleaseCapacityByBatchResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ResBatchId)) {
		body["ResBatchId"] = request.ResBatchId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ReleaseCapacityByBatch"),
		Version:     tea.String("2021-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ReleaseCapacityByBatchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据资源批次号释放 session 资源预定的资源
//
// @param request - ReleaseCapacityByBatchRequest
//
// @return ReleaseCapacityByBatchResponse
func (client *Client) ReleaseCapacityByBatch(request *ReleaseCapacityByBatchRequest) (_result *ReleaseCapacityByBatchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReleaseCapacityByBatchResponse{}
	_body, _err := client.ReleaseCapacityByBatchWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 退订GCS实例
//
// @param request - ReleaseInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReleaseInstancesResponse
func (client *Client) ReleaseInstancesWithOptions(request *ReleaseInstancesRequest, runtime *util.RuntimeOptions) (_result *ReleaseInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Amount)) {
		body["Amount"] = request.Amount
	}

	if !tea.BoolValue(util.IsUnset(request.DistrictId)) {
		body["DistrictId"] = request.DistrictId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceType)) {
		body["InstanceType"] = request.InstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ReleaseInstances"),
		Version:     tea.String("2021-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ReleaseInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 退订GCS实例
//
// @param request - ReleaseInstancesRequest
//
// @return ReleaseInstancesResponse
func (client *Client) ReleaseInstances(request *ReleaseInstancesRequest) (_result *ReleaseInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReleaseInstancesResponse{}
	_body, _err := client.ReleaseInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 预定GCS实例
//
// @param request - ReserveInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReserveInstancesResponse
func (client *Client) ReserveInstancesWithOptions(request *ReserveInstancesRequest, runtime *util.RuntimeOptions) (_result *ReserveInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Amount)) {
		body["Amount"] = request.Amount
	}

	if !tea.BoolValue(util.IsUnset(request.DistrictId)) {
		body["DistrictId"] = request.DistrictId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceType)) {
		body["InstanceType"] = request.InstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ReserveInstances"),
		Version:     tea.String("2021-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ReserveInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 预定GCS实例
//
// @param request - ReserveInstancesRequest
//
// @return ReserveInstancesResponse
func (client *Client) ReserveInstances(request *ReserveInstancesRequest) (_result *ReserveInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReserveInstancesResponse{}
	_body, _err := client.ReserveInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 发送业务能力变更结果回调
//
// @param request - SendBizCocChangeCallbackRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendBizCocChangeCallbackResponse
func (client *Client) SendBizCocChangeCallbackWithOptions(request *SendBizCocChangeCallbackRequest, runtime *util.RuntimeOptions) (_result *SendBizCocChangeCallbackResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PlatformSessionId)) {
		query["PlatformSessionId"] = request.PlatformSessionId
	}

	if !tea.BoolValue(util.IsUnset(request.Result)) {
		query["Result"] = request.Result
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SendBizCocChangeCallback"),
		Version:     tea.String("2021-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SendBizCocChangeCallbackResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 发送业务能力变更结果回调
//
// @param request - SendBizCocChangeCallbackRequest
//
// @return SendBizCocChangeCallbackResponse
func (client *Client) SendBizCocChangeCallback(request *SendBizCocChangeCallbackRequest) (_result *SendBizCocChangeCallbackResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SendBizCocChangeCallbackResponse{}
	_body, _err := client.SendBizCocChangeCallbackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 停止App会话
//
// @param tmpReq - StopAppSessionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopAppSessionResponse
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

// Summary:
//
// 停止App会话
//
// @param request - StopAppSessionRequest
//
// @return StopAppSessionResponse
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

// Summary:
//
// 批量停止会话接口
//
// @param tmpReq - StopAppSessionBatchRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopAppSessionBatchResponse
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

// Summary:
//
// 批量停止会话接口
//
// @param request - StopAppSessionBatchRequest
//
// @return StopAppSessionBatchResponse
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

// Summary:
//
// 更新会话状态
//
// @param request - UpdateSessionBizStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSessionBizStatusResponse
func (client *Client) UpdateSessionBizStatusWithOptions(request *UpdateSessionBizStatusRequest, runtime *util.RuntimeOptions) (_result *UpdateSessionBizStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizStatus)) {
		query["BizStatus"] = request.BizStatus
	}

	if !tea.BoolValue(util.IsUnset(request.PlatformSessionId)) {
		query["PlatformSessionId"] = request.PlatformSessionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateSessionBizStatus"),
		Version:     tea.String("2021-11-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateSessionBizStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新会话状态
//
// @param request - UpdateSessionBizStatusRequest
//
// @return UpdateSessionBizStatusResponse
func (client *Client) UpdateSessionBizStatus(request *UpdateSessionBizStatusRequest) (_result *UpdateSessionBizStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateSessionBizStatusResponse{}
	_body, _err := client.UpdateSessionBizStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
