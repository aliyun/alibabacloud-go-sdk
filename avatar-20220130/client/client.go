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

type CancelVideoTaskRequest struct {
	App      *CancelVideoTaskRequestApp `json:"App,omitempty" xml:"App,omitempty" type:"Struct"`
	TaskUuid *string                    `json:"TaskUuid,omitempty" xml:"TaskUuid,omitempty"`
	TenantId *int64                     `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s CancelVideoTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelVideoTaskRequest) GoString() string {
	return s.String()
}

func (s *CancelVideoTaskRequest) SetApp(v *CancelVideoTaskRequestApp) *CancelVideoTaskRequest {
	s.App = v
	return s
}

func (s *CancelVideoTaskRequest) SetTaskUuid(v string) *CancelVideoTaskRequest {
	s.TaskUuid = &v
	return s
}

func (s *CancelVideoTaskRequest) SetTenantId(v int64) *CancelVideoTaskRequest {
	s.TenantId = &v
	return s
}

type CancelVideoTaskRequestApp struct {
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s CancelVideoTaskRequestApp) String() string {
	return tea.Prettify(s)
}

func (s CancelVideoTaskRequestApp) GoString() string {
	return s.String()
}

func (s *CancelVideoTaskRequestApp) SetAppId(v string) *CancelVideoTaskRequestApp {
	s.AppId = &v
	return s
}

type CancelVideoTaskShrinkRequest struct {
	AppShrink *string `json:"App,omitempty" xml:"App,omitempty"`
	TaskUuid  *string `json:"TaskUuid,omitempty" xml:"TaskUuid,omitempty"`
	TenantId  *int64  `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s CancelVideoTaskShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelVideoTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *CancelVideoTaskShrinkRequest) SetAppShrink(v string) *CancelVideoTaskShrinkRequest {
	s.AppShrink = &v
	return s
}

func (s *CancelVideoTaskShrinkRequest) SetTaskUuid(v string) *CancelVideoTaskShrinkRequest {
	s.TaskUuid = &v
	return s
}

func (s *CancelVideoTaskShrinkRequest) SetTenantId(v int64) *CancelVideoTaskShrinkRequest {
	s.TenantId = &v
	return s
}

type CancelVideoTaskResponseBody struct {
	Code      *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *CancelVideoTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                            `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CancelVideoTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelVideoTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CancelVideoTaskResponseBody) SetCode(v string) *CancelVideoTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CancelVideoTaskResponseBody) SetData(v *CancelVideoTaskResponseBodyData) *CancelVideoTaskResponseBody {
	s.Data = v
	return s
}

func (s *CancelVideoTaskResponseBody) SetMessage(v string) *CancelVideoTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CancelVideoTaskResponseBody) SetRequestId(v string) *CancelVideoTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelVideoTaskResponseBody) SetSuccess(v bool) *CancelVideoTaskResponseBody {
	s.Success = &v
	return s
}

type CancelVideoTaskResponseBodyData struct {
	FailReason *string `json:"FailReason,omitempty" xml:"FailReason,omitempty"`
	IsCancel   *bool   `json:"IsCancel,omitempty" xml:"IsCancel,omitempty"`
	TaskUuid   *string `json:"TaskUuid,omitempty" xml:"TaskUuid,omitempty"`
}

func (s CancelVideoTaskResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CancelVideoTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *CancelVideoTaskResponseBodyData) SetFailReason(v string) *CancelVideoTaskResponseBodyData {
	s.FailReason = &v
	return s
}

func (s *CancelVideoTaskResponseBodyData) SetIsCancel(v bool) *CancelVideoTaskResponseBodyData {
	s.IsCancel = &v
	return s
}

func (s *CancelVideoTaskResponseBodyData) SetTaskUuid(v string) *CancelVideoTaskResponseBodyData {
	s.TaskUuid = &v
	return s
}

type CancelVideoTaskResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CancelVideoTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CancelVideoTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelVideoTaskResponse) GoString() string {
	return s.String()
}

func (s *CancelVideoTaskResponse) SetHeaders(v map[string]*string) *CancelVideoTaskResponse {
	s.Headers = v
	return s
}

func (s *CancelVideoTaskResponse) SetStatusCode(v int32) *CancelVideoTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelVideoTaskResponse) SetBody(v *CancelVideoTaskResponseBody) *CancelVideoTaskResponse {
	s.Body = v
	return s
}

type ClientAuthRequest struct {
	AppId      *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	DeviceId   *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	DeviceInfo *string `json:"DeviceInfo,omitempty" xml:"DeviceInfo,omitempty"`
	DeviceType *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	License    *string `json:"License,omitempty" xml:"License,omitempty"`
	TenantId   *int64  `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s ClientAuthRequest) String() string {
	return tea.Prettify(s)
}

func (s ClientAuthRequest) GoString() string {
	return s.String()
}

func (s *ClientAuthRequest) SetAppId(v string) *ClientAuthRequest {
	s.AppId = &v
	return s
}

func (s *ClientAuthRequest) SetDeviceId(v string) *ClientAuthRequest {
	s.DeviceId = &v
	return s
}

func (s *ClientAuthRequest) SetDeviceInfo(v string) *ClientAuthRequest {
	s.DeviceInfo = &v
	return s
}

func (s *ClientAuthRequest) SetDeviceType(v string) *ClientAuthRequest {
	s.DeviceType = &v
	return s
}

func (s *ClientAuthRequest) SetLicense(v string) *ClientAuthRequest {
	s.License = &v
	return s
}

func (s *ClientAuthRequest) SetTenantId(v int64) *ClientAuthRequest {
	s.TenantId = &v
	return s
}

type ClientAuthResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ClientAuthResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ClientAuthResponseBody) GoString() string {
	return s.String()
}

func (s *ClientAuthResponseBody) SetCode(v string) *ClientAuthResponseBody {
	s.Code = &v
	return s
}

func (s *ClientAuthResponseBody) SetMessage(v string) *ClientAuthResponseBody {
	s.Message = &v
	return s
}

func (s *ClientAuthResponseBody) SetRequestId(v string) *ClientAuthResponseBody {
	s.RequestId = &v
	return s
}

func (s *ClientAuthResponseBody) SetSuccess(v bool) *ClientAuthResponseBody {
	s.Success = &v
	return s
}

type ClientAuthResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ClientAuthResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ClientAuthResponse) String() string {
	return tea.Prettify(s)
}

func (s ClientAuthResponse) GoString() string {
	return s.String()
}

func (s *ClientAuthResponse) SetHeaders(v map[string]*string) *ClientAuthResponse {
	s.Headers = v
	return s
}

func (s *ClientAuthResponse) SetStatusCode(v int32) *ClientAuthResponse {
	s.StatusCode = &v
	return s
}

func (s *ClientAuthResponse) SetBody(v *ClientAuthResponseBody) *ClientAuthResponse {
	s.Body = v
	return s
}

type ClientStartRequest struct {
	AppId    *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	TenantId *int64  `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s ClientStartRequest) String() string {
	return tea.Prettify(s)
}

func (s ClientStartRequest) GoString() string {
	return s.String()
}

func (s *ClientStartRequest) SetAppId(v string) *ClientStartRequest {
	s.AppId = &v
	return s
}

func (s *ClientStartRequest) SetTenantId(v int64) *ClientStartRequest {
	s.TenantId = &v
	return s
}

type ClientStartResponseBody struct {
	Code      *string                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *ClientStartResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                        `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ClientStartResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ClientStartResponseBody) GoString() string {
	return s.String()
}

func (s *ClientStartResponseBody) SetCode(v string) *ClientStartResponseBody {
	s.Code = &v
	return s
}

func (s *ClientStartResponseBody) SetData(v *ClientStartResponseBodyData) *ClientStartResponseBody {
	s.Data = v
	return s
}

func (s *ClientStartResponseBody) SetMessage(v string) *ClientStartResponseBody {
	s.Message = &v
	return s
}

func (s *ClientStartResponseBody) SetRequestId(v string) *ClientStartResponseBody {
	s.RequestId = &v
	return s
}

func (s *ClientStartResponseBody) SetSuccess(v bool) *ClientStartResponseBody {
	s.Success = &v
	return s
}

type ClientStartResponseBodyData struct {
	ImToken *string `json:"ImToken,omitempty" xml:"ImToken,omitempty"`
}

func (s ClientStartResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ClientStartResponseBodyData) GoString() string {
	return s.String()
}

func (s *ClientStartResponseBodyData) SetImToken(v string) *ClientStartResponseBodyData {
	s.ImToken = &v
	return s
}

type ClientStartResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ClientStartResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ClientStartResponse) String() string {
	return tea.Prettify(s)
}

func (s ClientStartResponse) GoString() string {
	return s.String()
}

func (s *ClientStartResponse) SetHeaders(v map[string]*string) *ClientStartResponse {
	s.Headers = v
	return s
}

func (s *ClientStartResponse) SetStatusCode(v int32) *ClientStartResponse {
	s.StatusCode = &v
	return s
}

func (s *ClientStartResponse) SetBody(v *ClientStartResponseBody) *ClientStartResponse {
	s.Body = v
	return s
}

type CloseTimedResetOperateRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	TenantId   *int64  `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s CloseTimedResetOperateRequest) String() string {
	return tea.Prettify(s)
}

func (s CloseTimedResetOperateRequest) GoString() string {
	return s.String()
}

func (s *CloseTimedResetOperateRequest) SetInstanceId(v string) *CloseTimedResetOperateRequest {
	s.InstanceId = &v
	return s
}

func (s *CloseTimedResetOperateRequest) SetTenantId(v int64) *CloseTimedResetOperateRequest {
	s.TenantId = &v
	return s
}

type CloseTimedResetOperateResponseBody struct {
	Code    *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    *CloseTimedResetOperateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CloseTimedResetOperateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CloseTimedResetOperateResponseBody) GoString() string {
	return s.String()
}

func (s *CloseTimedResetOperateResponseBody) SetCode(v string) *CloseTimedResetOperateResponseBody {
	s.Code = &v
	return s
}

func (s *CloseTimedResetOperateResponseBody) SetData(v *CloseTimedResetOperateResponseBodyData) *CloseTimedResetOperateResponseBody {
	s.Data = v
	return s
}

func (s *CloseTimedResetOperateResponseBody) SetMessage(v string) *CloseTimedResetOperateResponseBody {
	s.Message = &v
	return s
}

func (s *CloseTimedResetOperateResponseBody) SetRequestId(v string) *CloseTimedResetOperateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloseTimedResetOperateResponseBody) SetSuccess(v bool) *CloseTimedResetOperateResponseBody {
	s.Success = &v
	return s
}

type CloseTimedResetOperateResponseBodyData struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	TenantId   *int64  `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s CloseTimedResetOperateResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CloseTimedResetOperateResponseBodyData) GoString() string {
	return s.String()
}

func (s *CloseTimedResetOperateResponseBodyData) SetInstanceId(v string) *CloseTimedResetOperateResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *CloseTimedResetOperateResponseBodyData) SetTenantId(v int64) *CloseTimedResetOperateResponseBodyData {
	s.TenantId = &v
	return s
}

type CloseTimedResetOperateResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CloseTimedResetOperateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CloseTimedResetOperateResponse) String() string {
	return tea.Prettify(s)
}

func (s CloseTimedResetOperateResponse) GoString() string {
	return s.String()
}

func (s *CloseTimedResetOperateResponse) SetHeaders(v map[string]*string) *CloseTimedResetOperateResponse {
	s.Headers = v
	return s
}

func (s *CloseTimedResetOperateResponse) SetStatusCode(v int32) *CloseTimedResetOperateResponse {
	s.StatusCode = &v
	return s
}

func (s *CloseTimedResetOperateResponse) SetBody(v *CloseTimedResetOperateResponseBody) *CloseTimedResetOperateResponse {
	s.Body = v
	return s
}

type Create2dAvatarRequest struct {
	Callback    *bool   `json:"Callback,omitempty" xml:"Callback,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Image       *string `json:"Image,omitempty" xml:"Image,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Orientation *int32  `json:"Orientation,omitempty" xml:"Orientation,omitempty"`
	Portrait    *string `json:"Portrait,omitempty" xml:"Portrait,omitempty"`
	TenantId    *int64  `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	Transparent *bool   `json:"Transparent,omitempty" xml:"Transparent,omitempty"`
	Video       *string `json:"Video,omitempty" xml:"Video,omitempty"`
}

func (s Create2dAvatarRequest) String() string {
	return tea.Prettify(s)
}

func (s Create2dAvatarRequest) GoString() string {
	return s.String()
}

func (s *Create2dAvatarRequest) SetCallback(v bool) *Create2dAvatarRequest {
	s.Callback = &v
	return s
}

func (s *Create2dAvatarRequest) SetDescription(v string) *Create2dAvatarRequest {
	s.Description = &v
	return s
}

func (s *Create2dAvatarRequest) SetImage(v string) *Create2dAvatarRequest {
	s.Image = &v
	return s
}

func (s *Create2dAvatarRequest) SetName(v string) *Create2dAvatarRequest {
	s.Name = &v
	return s
}

func (s *Create2dAvatarRequest) SetOrientation(v int32) *Create2dAvatarRequest {
	s.Orientation = &v
	return s
}

func (s *Create2dAvatarRequest) SetPortrait(v string) *Create2dAvatarRequest {
	s.Portrait = &v
	return s
}

func (s *Create2dAvatarRequest) SetTenantId(v int64) *Create2dAvatarRequest {
	s.TenantId = &v
	return s
}

func (s *Create2dAvatarRequest) SetTransparent(v bool) *Create2dAvatarRequest {
	s.Transparent = &v
	return s
}

func (s *Create2dAvatarRequest) SetVideo(v string) *Create2dAvatarRequest {
	s.Video = &v
	return s
}

type Create2dAvatarResponseBody struct {
	Code      *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *Create2dAvatarResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                           `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s Create2dAvatarResponseBody) String() string {
	return tea.Prettify(s)
}

func (s Create2dAvatarResponseBody) GoString() string {
	return s.String()
}

func (s *Create2dAvatarResponseBody) SetCode(v string) *Create2dAvatarResponseBody {
	s.Code = &v
	return s
}

func (s *Create2dAvatarResponseBody) SetData(v *Create2dAvatarResponseBodyData) *Create2dAvatarResponseBody {
	s.Data = v
	return s
}

func (s *Create2dAvatarResponseBody) SetMessage(v string) *Create2dAvatarResponseBody {
	s.Message = &v
	return s
}

func (s *Create2dAvatarResponseBody) SetRequestId(v string) *Create2dAvatarResponseBody {
	s.RequestId = &v
	return s
}

func (s *Create2dAvatarResponseBody) SetSuccess(v bool) *Create2dAvatarResponseBody {
	s.Success = &v
	return s
}

type Create2dAvatarResponseBodyData struct {
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s Create2dAvatarResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s Create2dAvatarResponseBodyData) GoString() string {
	return s.String()
}

func (s *Create2dAvatarResponseBodyData) SetCode(v string) *Create2dAvatarResponseBodyData {
	s.Code = &v
	return s
}

type Create2dAvatarResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *Create2dAvatarResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s Create2dAvatarResponse) String() string {
	return tea.Prettify(s)
}

func (s Create2dAvatarResponse) GoString() string {
	return s.String()
}

func (s *Create2dAvatarResponse) SetHeaders(v map[string]*string) *Create2dAvatarResponse {
	s.Headers = v
	return s
}

func (s *Create2dAvatarResponse) SetStatusCode(v int32) *Create2dAvatarResponse {
	s.StatusCode = &v
	return s
}

func (s *Create2dAvatarResponse) SetBody(v *Create2dAvatarResponseBody) *Create2dAvatarResponse {
	s.Body = v
	return s
}

type DeleteAvatarRequest struct {
	Code     *string `json:"Code,omitempty" xml:"Code,omitempty"`
	TenantId *int64  `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s DeleteAvatarRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAvatarRequest) GoString() string {
	return s.String()
}

func (s *DeleteAvatarRequest) SetCode(v string) *DeleteAvatarRequest {
	s.Code = &v
	return s
}

func (s *DeleteAvatarRequest) SetTenantId(v int64) *DeleteAvatarRequest {
	s.TenantId = &v
	return s
}

type DeleteAvatarResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteAvatarResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAvatarResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAvatarResponseBody) SetCode(v string) *DeleteAvatarResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteAvatarResponseBody) SetMessage(v string) *DeleteAvatarResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteAvatarResponseBody) SetRequestId(v string) *DeleteAvatarResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAvatarResponseBody) SetSuccess(v bool) *DeleteAvatarResponseBody {
	s.Success = &v
	return s
}

type DeleteAvatarResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteAvatarResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteAvatarResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAvatarResponse) GoString() string {
	return s.String()
}

func (s *DeleteAvatarResponse) SetHeaders(v map[string]*string) *DeleteAvatarResponse {
	s.Headers = v
	return s
}

func (s *DeleteAvatarResponse) SetStatusCode(v int32) *DeleteAvatarResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAvatarResponse) SetBody(v *DeleteAvatarResponseBody) *DeleteAvatarResponse {
	s.Body = v
	return s
}

type DuplexDecisionRequest struct {
	AppId          *string                             `json:"AppId,omitempty" xml:"AppId,omitempty"`
	BizRequestId   *string                             `json:"BizRequestId,omitempty" xml:"BizRequestId,omitempty"`
	CallTime       *int32                              `json:"CallTime,omitempty" xml:"CallTime,omitempty"`
	CustomKeywords []*string                           `json:"CustomKeywords,omitempty" xml:"CustomKeywords,omitempty" type:"Repeated"`
	DialogContext  *DuplexDecisionRequestDialogContext `json:"DialogContext,omitempty" xml:"DialogContext,omitempty" type:"Struct"`
	DialogStatus   *string                             `json:"DialogStatus,omitempty" xml:"DialogStatus,omitempty"`
	InterruptType  *int32                              `json:"InterruptType,omitempty" xml:"InterruptType,omitempty"`
	SessionId      *string                             `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	TenantId       *int64                              `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	Text           *string                             `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s DuplexDecisionRequest) String() string {
	return tea.Prettify(s)
}

func (s DuplexDecisionRequest) GoString() string {
	return s.String()
}

func (s *DuplexDecisionRequest) SetAppId(v string) *DuplexDecisionRequest {
	s.AppId = &v
	return s
}

func (s *DuplexDecisionRequest) SetBizRequestId(v string) *DuplexDecisionRequest {
	s.BizRequestId = &v
	return s
}

func (s *DuplexDecisionRequest) SetCallTime(v int32) *DuplexDecisionRequest {
	s.CallTime = &v
	return s
}

func (s *DuplexDecisionRequest) SetCustomKeywords(v []*string) *DuplexDecisionRequest {
	s.CustomKeywords = v
	return s
}

func (s *DuplexDecisionRequest) SetDialogContext(v *DuplexDecisionRequestDialogContext) *DuplexDecisionRequest {
	s.DialogContext = v
	return s
}

func (s *DuplexDecisionRequest) SetDialogStatus(v string) *DuplexDecisionRequest {
	s.DialogStatus = &v
	return s
}

func (s *DuplexDecisionRequest) SetInterruptType(v int32) *DuplexDecisionRequest {
	s.InterruptType = &v
	return s
}

func (s *DuplexDecisionRequest) SetSessionId(v string) *DuplexDecisionRequest {
	s.SessionId = &v
	return s
}

func (s *DuplexDecisionRequest) SetTenantId(v int64) *DuplexDecisionRequest {
	s.TenantId = &v
	return s
}

func (s *DuplexDecisionRequest) SetText(v string) *DuplexDecisionRequest {
	s.Text = &v
	return s
}

type DuplexDecisionRequestDialogContext struct {
	CurUtteranceIdx *int32                                         `json:"CurUtteranceIdx,omitempty" xml:"CurUtteranceIdx,omitempty"`
	Histories       []*DuplexDecisionRequestDialogContextHistories `json:"Histories,omitempty" xml:"Histories,omitempty" type:"Repeated"`
}

func (s DuplexDecisionRequestDialogContext) String() string {
	return tea.Prettify(s)
}

func (s DuplexDecisionRequestDialogContext) GoString() string {
	return s.String()
}

func (s *DuplexDecisionRequestDialogContext) SetCurUtteranceIdx(v int32) *DuplexDecisionRequestDialogContext {
	s.CurUtteranceIdx = &v
	return s
}

func (s *DuplexDecisionRequestDialogContext) SetHistories(v []*DuplexDecisionRequestDialogContextHistories) *DuplexDecisionRequestDialogContext {
	s.Histories = v
	return s
}

type DuplexDecisionRequestDialogContextHistories struct {
	Robot *string `json:"Robot,omitempty" xml:"Robot,omitempty"`
	User  *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s DuplexDecisionRequestDialogContextHistories) String() string {
	return tea.Prettify(s)
}

func (s DuplexDecisionRequestDialogContextHistories) GoString() string {
	return s.String()
}

func (s *DuplexDecisionRequestDialogContextHistories) SetRobot(v string) *DuplexDecisionRequestDialogContextHistories {
	s.Robot = &v
	return s
}

func (s *DuplexDecisionRequestDialogContextHistories) SetUser(v string) *DuplexDecisionRequestDialogContextHistories {
	s.User = &v
	return s
}

type DuplexDecisionShrinkRequest struct {
	AppId                *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	BizRequestId         *string `json:"BizRequestId,omitempty" xml:"BizRequestId,omitempty"`
	CallTime             *int32  `json:"CallTime,omitempty" xml:"CallTime,omitempty"`
	CustomKeywordsShrink *string `json:"CustomKeywords,omitempty" xml:"CustomKeywords,omitempty"`
	DialogContextShrink  *string `json:"DialogContext,omitempty" xml:"DialogContext,omitempty"`
	DialogStatus         *string `json:"DialogStatus,omitempty" xml:"DialogStatus,omitempty"`
	InterruptType        *int32  `json:"InterruptType,omitempty" xml:"InterruptType,omitempty"`
	SessionId            *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	TenantId             *int64  `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	Text                 *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s DuplexDecisionShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DuplexDecisionShrinkRequest) GoString() string {
	return s.String()
}

func (s *DuplexDecisionShrinkRequest) SetAppId(v string) *DuplexDecisionShrinkRequest {
	s.AppId = &v
	return s
}

func (s *DuplexDecisionShrinkRequest) SetBizRequestId(v string) *DuplexDecisionShrinkRequest {
	s.BizRequestId = &v
	return s
}

func (s *DuplexDecisionShrinkRequest) SetCallTime(v int32) *DuplexDecisionShrinkRequest {
	s.CallTime = &v
	return s
}

func (s *DuplexDecisionShrinkRequest) SetCustomKeywordsShrink(v string) *DuplexDecisionShrinkRequest {
	s.CustomKeywordsShrink = &v
	return s
}

func (s *DuplexDecisionShrinkRequest) SetDialogContextShrink(v string) *DuplexDecisionShrinkRequest {
	s.DialogContextShrink = &v
	return s
}

func (s *DuplexDecisionShrinkRequest) SetDialogStatus(v string) *DuplexDecisionShrinkRequest {
	s.DialogStatus = &v
	return s
}

func (s *DuplexDecisionShrinkRequest) SetInterruptType(v int32) *DuplexDecisionShrinkRequest {
	s.InterruptType = &v
	return s
}

func (s *DuplexDecisionShrinkRequest) SetSessionId(v string) *DuplexDecisionShrinkRequest {
	s.SessionId = &v
	return s
}

func (s *DuplexDecisionShrinkRequest) SetTenantId(v int64) *DuplexDecisionShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *DuplexDecisionShrinkRequest) SetText(v string) *DuplexDecisionShrinkRequest {
	s.Text = &v
	return s
}

type DuplexDecisionResponseBody struct {
	Code      *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *DuplexDecisionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *string                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DuplexDecisionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DuplexDecisionResponseBody) GoString() string {
	return s.String()
}

func (s *DuplexDecisionResponseBody) SetCode(v string) *DuplexDecisionResponseBody {
	s.Code = &v
	return s
}

func (s *DuplexDecisionResponseBody) SetData(v *DuplexDecisionResponseBodyData) *DuplexDecisionResponseBody {
	s.Data = v
	return s
}

func (s *DuplexDecisionResponseBody) SetMessage(v string) *DuplexDecisionResponseBody {
	s.Message = &v
	return s
}

func (s *DuplexDecisionResponseBody) SetRequestId(v string) *DuplexDecisionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DuplexDecisionResponseBody) SetSuccess(v string) *DuplexDecisionResponseBody {
	s.Success = &v
	return s
}

type DuplexDecisionResponseBodyData struct {
	ActionType *string `json:"ActionType,omitempty" xml:"ActionType,omitempty"`
	GrabType   *string `json:"GrabType,omitempty" xml:"GrabType,omitempty"`
	OutputText *string `json:"OutputText,omitempty" xml:"OutputText,omitempty"`
}

func (s DuplexDecisionResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DuplexDecisionResponseBodyData) GoString() string {
	return s.String()
}

func (s *DuplexDecisionResponseBodyData) SetActionType(v string) *DuplexDecisionResponseBodyData {
	s.ActionType = &v
	return s
}

func (s *DuplexDecisionResponseBodyData) SetGrabType(v string) *DuplexDecisionResponseBodyData {
	s.GrabType = &v
	return s
}

func (s *DuplexDecisionResponseBodyData) SetOutputText(v string) *DuplexDecisionResponseBodyData {
	s.OutputText = &v
	return s
}

type DuplexDecisionResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DuplexDecisionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DuplexDecisionResponse) String() string {
	return tea.Prettify(s)
}

func (s DuplexDecisionResponse) GoString() string {
	return s.String()
}

func (s *DuplexDecisionResponse) SetHeaders(v map[string]*string) *DuplexDecisionResponse {
	s.Headers = v
	return s
}

func (s *DuplexDecisionResponse) SetStatusCode(v int32) *DuplexDecisionResponse {
	s.StatusCode = &v
	return s
}

func (s *DuplexDecisionResponse) SetBody(v *DuplexDecisionResponseBody) *DuplexDecisionResponse {
	s.Body = v
	return s
}

type GetVideoTaskInfoRequest struct {
	App      *GetVideoTaskInfoRequestApp `json:"App,omitempty" xml:"App,omitempty" type:"Struct"`
	TaskUuid *string                     `json:"TaskUuid,omitempty" xml:"TaskUuid,omitempty"`
	TenantId *int64                      `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s GetVideoTaskInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetVideoTaskInfoRequest) GoString() string {
	return s.String()
}

func (s *GetVideoTaskInfoRequest) SetApp(v *GetVideoTaskInfoRequestApp) *GetVideoTaskInfoRequest {
	s.App = v
	return s
}

func (s *GetVideoTaskInfoRequest) SetTaskUuid(v string) *GetVideoTaskInfoRequest {
	s.TaskUuid = &v
	return s
}

func (s *GetVideoTaskInfoRequest) SetTenantId(v int64) *GetVideoTaskInfoRequest {
	s.TenantId = &v
	return s
}

type GetVideoTaskInfoRequestApp struct {
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s GetVideoTaskInfoRequestApp) String() string {
	return tea.Prettify(s)
}

func (s GetVideoTaskInfoRequestApp) GoString() string {
	return s.String()
}

func (s *GetVideoTaskInfoRequestApp) SetAppId(v string) *GetVideoTaskInfoRequestApp {
	s.AppId = &v
	return s
}

type GetVideoTaskInfoShrinkRequest struct {
	AppShrink *string `json:"App,omitempty" xml:"App,omitempty"`
	TaskUuid  *string `json:"TaskUuid,omitempty" xml:"TaskUuid,omitempty"`
	TenantId  *int64  `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s GetVideoTaskInfoShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetVideoTaskInfoShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetVideoTaskInfoShrinkRequest) SetAppShrink(v string) *GetVideoTaskInfoShrinkRequest {
	s.AppShrink = &v
	return s
}

func (s *GetVideoTaskInfoShrinkRequest) SetTaskUuid(v string) *GetVideoTaskInfoShrinkRequest {
	s.TaskUuid = &v
	return s
}

func (s *GetVideoTaskInfoShrinkRequest) SetTenantId(v int64) *GetVideoTaskInfoShrinkRequest {
	s.TenantId = &v
	return s
}

type GetVideoTaskInfoResponseBody struct {
	Code      *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetVideoTaskInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetVideoTaskInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetVideoTaskInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetVideoTaskInfoResponseBody) SetCode(v string) *GetVideoTaskInfoResponseBody {
	s.Code = &v
	return s
}

func (s *GetVideoTaskInfoResponseBody) SetData(v *GetVideoTaskInfoResponseBodyData) *GetVideoTaskInfoResponseBody {
	s.Data = v
	return s
}

func (s *GetVideoTaskInfoResponseBody) SetMessage(v string) *GetVideoTaskInfoResponseBody {
	s.Message = &v
	return s
}

func (s *GetVideoTaskInfoResponseBody) SetRequestId(v string) *GetVideoTaskInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetVideoTaskInfoResponseBody) SetSuccess(v bool) *GetVideoTaskInfoResponseBody {
	s.Success = &v
	return s
}

type GetVideoTaskInfoResponseBodyData struct {
	Process    *string                                     `json:"Process,omitempty" xml:"Process,omitempty"`
	Status     *string                                     `json:"Status,omitempty" xml:"Status,omitempty"`
	TaskResult *GetVideoTaskInfoResponseBodyDataTaskResult `json:"TaskResult,omitempty" xml:"TaskResult,omitempty" type:"Struct"`
	TaskUuid   *string                                     `json:"TaskUuid,omitempty" xml:"TaskUuid,omitempty"`
	Type       *string                                     `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetVideoTaskInfoResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetVideoTaskInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetVideoTaskInfoResponseBodyData) SetProcess(v string) *GetVideoTaskInfoResponseBodyData {
	s.Process = &v
	return s
}

func (s *GetVideoTaskInfoResponseBodyData) SetStatus(v string) *GetVideoTaskInfoResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetVideoTaskInfoResponseBodyData) SetTaskResult(v *GetVideoTaskInfoResponseBodyDataTaskResult) *GetVideoTaskInfoResponseBodyData {
	s.TaskResult = v
	return s
}

func (s *GetVideoTaskInfoResponseBodyData) SetTaskUuid(v string) *GetVideoTaskInfoResponseBodyData {
	s.TaskUuid = &v
	return s
}

func (s *GetVideoTaskInfoResponseBodyData) SetType(v string) *GetVideoTaskInfoResponseBodyData {
	s.Type = &v
	return s
}

type GetVideoTaskInfoResponseBodyDataTaskResult struct {
	AlphaUrl      *string `json:"AlphaUrl,omitempty" xml:"AlphaUrl,omitempty"`
	AttachmentUrl *string `json:"AttachmentUrl,omitempty" xml:"AttachmentUrl,omitempty"`
	FailCode      *string `json:"FailCode,omitempty" xml:"FailCode,omitempty"`
	FailReason    *string `json:"FailReason,omitempty" xml:"FailReason,omitempty"`
	PreviewPic    *string `json:"PreviewPic,omitempty" xml:"PreviewPic,omitempty"`
	SubtitlesUrl  *string `json:"SubtitlesUrl,omitempty" xml:"SubtitlesUrl,omitempty"`
	VideoDuration *int32  `json:"VideoDuration,omitempty" xml:"VideoDuration,omitempty"`
	VideoUrl      *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
	// 字粒度的时间戳文件，特定任务支持
	WordSubtitlesUrl *string `json:"WordSubtitlesUrl,omitempty" xml:"WordSubtitlesUrl,omitempty"`
}

func (s GetVideoTaskInfoResponseBodyDataTaskResult) String() string {
	return tea.Prettify(s)
}

func (s GetVideoTaskInfoResponseBodyDataTaskResult) GoString() string {
	return s.String()
}

func (s *GetVideoTaskInfoResponseBodyDataTaskResult) SetAlphaUrl(v string) *GetVideoTaskInfoResponseBodyDataTaskResult {
	s.AlphaUrl = &v
	return s
}

func (s *GetVideoTaskInfoResponseBodyDataTaskResult) SetAttachmentUrl(v string) *GetVideoTaskInfoResponseBodyDataTaskResult {
	s.AttachmentUrl = &v
	return s
}

func (s *GetVideoTaskInfoResponseBodyDataTaskResult) SetFailCode(v string) *GetVideoTaskInfoResponseBodyDataTaskResult {
	s.FailCode = &v
	return s
}

func (s *GetVideoTaskInfoResponseBodyDataTaskResult) SetFailReason(v string) *GetVideoTaskInfoResponseBodyDataTaskResult {
	s.FailReason = &v
	return s
}

func (s *GetVideoTaskInfoResponseBodyDataTaskResult) SetPreviewPic(v string) *GetVideoTaskInfoResponseBodyDataTaskResult {
	s.PreviewPic = &v
	return s
}

func (s *GetVideoTaskInfoResponseBodyDataTaskResult) SetSubtitlesUrl(v string) *GetVideoTaskInfoResponseBodyDataTaskResult {
	s.SubtitlesUrl = &v
	return s
}

func (s *GetVideoTaskInfoResponseBodyDataTaskResult) SetVideoDuration(v int32) *GetVideoTaskInfoResponseBodyDataTaskResult {
	s.VideoDuration = &v
	return s
}

func (s *GetVideoTaskInfoResponseBodyDataTaskResult) SetVideoUrl(v string) *GetVideoTaskInfoResponseBodyDataTaskResult {
	s.VideoUrl = &v
	return s
}

func (s *GetVideoTaskInfoResponseBodyDataTaskResult) SetWordSubtitlesUrl(v string) *GetVideoTaskInfoResponseBodyDataTaskResult {
	s.WordSubtitlesUrl = &v
	return s
}

type GetVideoTaskInfoResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetVideoTaskInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetVideoTaskInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetVideoTaskInfoResponse) GoString() string {
	return s.String()
}

func (s *GetVideoTaskInfoResponse) SetHeaders(v map[string]*string) *GetVideoTaskInfoResponse {
	s.Headers = v
	return s
}

func (s *GetVideoTaskInfoResponse) SetStatusCode(v int32) *GetVideoTaskInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVideoTaskInfoResponse) SetBody(v *GetVideoTaskInfoResponseBody) *GetVideoTaskInfoResponse {
	s.Body = v
	return s
}

type LicenseAuthRequest struct {
	AppId    *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	License  *string `json:"License,omitempty" xml:"License,omitempty"`
	TenantId *int64  `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s LicenseAuthRequest) String() string {
	return tea.Prettify(s)
}

func (s LicenseAuthRequest) GoString() string {
	return s.String()
}

func (s *LicenseAuthRequest) SetAppId(v string) *LicenseAuthRequest {
	s.AppId = &v
	return s
}

func (s *LicenseAuthRequest) SetLicense(v string) *LicenseAuthRequest {
	s.License = &v
	return s
}

func (s *LicenseAuthRequest) SetTenantId(v int64) *LicenseAuthRequest {
	s.TenantId = &v
	return s
}

type LicenseAuthResponseBody struct {
	Code      *string                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *LicenseAuthResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                        `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s LicenseAuthResponseBody) String() string {
	return tea.Prettify(s)
}

func (s LicenseAuthResponseBody) GoString() string {
	return s.String()
}

func (s *LicenseAuthResponseBody) SetCode(v string) *LicenseAuthResponseBody {
	s.Code = &v
	return s
}

func (s *LicenseAuthResponseBody) SetData(v *LicenseAuthResponseBodyData) *LicenseAuthResponseBody {
	s.Data = v
	return s
}

func (s *LicenseAuthResponseBody) SetMessage(v string) *LicenseAuthResponseBody {
	s.Message = &v
	return s
}

func (s *LicenseAuthResponseBody) SetRequestId(v string) *LicenseAuthResponseBody {
	s.RequestId = &v
	return s
}

func (s *LicenseAuthResponseBody) SetSuccess(v bool) *LicenseAuthResponseBody {
	s.Success = &v
	return s
}

type LicenseAuthResponseBodyData struct {
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s LicenseAuthResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s LicenseAuthResponseBodyData) GoString() string {
	return s.String()
}

func (s *LicenseAuthResponseBodyData) SetToken(v string) *LicenseAuthResponseBodyData {
	s.Token = &v
	return s
}

type LicenseAuthResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *LicenseAuthResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s LicenseAuthResponse) String() string {
	return tea.Prettify(s)
}

func (s LicenseAuthResponse) GoString() string {
	return s.String()
}

func (s *LicenseAuthResponse) SetHeaders(v map[string]*string) *LicenseAuthResponse {
	s.Headers = v
	return s
}

func (s *LicenseAuthResponse) SetStatusCode(v int32) *LicenseAuthResponse {
	s.StatusCode = &v
	return s
}

func (s *LicenseAuthResponse) SetBody(v *LicenseAuthResponseBody) *LicenseAuthResponse {
	s.Body = v
	return s
}

type QueryAvatarRequest struct {
	Code     *string `json:"Code,omitempty" xml:"Code,omitempty"`
	TenantId *int64  `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s QueryAvatarRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryAvatarRequest) GoString() string {
	return s.String()
}

func (s *QueryAvatarRequest) SetCode(v string) *QueryAvatarRequest {
	s.Code = &v
	return s
}

func (s *QueryAvatarRequest) SetTenantId(v int64) *QueryAvatarRequest {
	s.TenantId = &v
	return s
}

type QueryAvatarResponseBody struct {
	Code      *string                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *QueryAvatarResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                        `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryAvatarResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryAvatarResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAvatarResponseBody) SetCode(v string) *QueryAvatarResponseBody {
	s.Code = &v
	return s
}

func (s *QueryAvatarResponseBody) SetData(v *QueryAvatarResponseBodyData) *QueryAvatarResponseBody {
	s.Data = v
	return s
}

func (s *QueryAvatarResponseBody) SetMessage(v string) *QueryAvatarResponseBody {
	s.Message = &v
	return s
}

func (s *QueryAvatarResponseBody) SetRequestId(v string) *QueryAvatarResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryAvatarResponseBody) SetSuccess(v bool) *QueryAvatarResponseBody {
	s.Success = &v
	return s
}

type QueryAvatarResponseBodyData struct {
	AllLocateImages      map[string]interface{}                           `json:"AllLocateImages,omitempty" xml:"AllLocateImages,omitempty"`
	AvatarType           *string                                          `json:"AvatarType,omitempty" xml:"AvatarType,omitempty"`
	Description          *string                                          `json:"Description,omitempty" xml:"Description,omitempty"`
	Image                *string                                          `json:"Image,omitempty" xml:"Image,omitempty"`
	MakeFailReason       *string                                          `json:"MakeFailReason,omitempty" xml:"MakeFailReason,omitempty"`
	MakeStage            *string                                          `json:"MakeStage,omitempty" xml:"MakeStage,omitempty"`
	MakeStatus           *string                                          `json:"MakeStatus,omitempty" xml:"MakeStatus,omitempty"`
	ModelType            *string                                          `json:"ModelType,omitempty" xml:"ModelType,omitempty"`
	Name                 *string                                          `json:"Name,omitempty" xml:"Name,omitempty"`
	Portrait             *string                                          `json:"Portrait,omitempty" xml:"Portrait,omitempty"`
	SupportedResolutions *QueryAvatarResponseBodyDataSupportedResolutions `json:"SupportedResolutions,omitempty" xml:"SupportedResolutions,omitempty" type:"Struct"`
}

func (s QueryAvatarResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryAvatarResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryAvatarResponseBodyData) SetAllLocateImages(v map[string]interface{}) *QueryAvatarResponseBodyData {
	s.AllLocateImages = v
	return s
}

func (s *QueryAvatarResponseBodyData) SetAvatarType(v string) *QueryAvatarResponseBodyData {
	s.AvatarType = &v
	return s
}

func (s *QueryAvatarResponseBodyData) SetDescription(v string) *QueryAvatarResponseBodyData {
	s.Description = &v
	return s
}

func (s *QueryAvatarResponseBodyData) SetImage(v string) *QueryAvatarResponseBodyData {
	s.Image = &v
	return s
}

func (s *QueryAvatarResponseBodyData) SetMakeFailReason(v string) *QueryAvatarResponseBodyData {
	s.MakeFailReason = &v
	return s
}

func (s *QueryAvatarResponseBodyData) SetMakeStage(v string) *QueryAvatarResponseBodyData {
	s.MakeStage = &v
	return s
}

func (s *QueryAvatarResponseBodyData) SetMakeStatus(v string) *QueryAvatarResponseBodyData {
	s.MakeStatus = &v
	return s
}

func (s *QueryAvatarResponseBodyData) SetModelType(v string) *QueryAvatarResponseBodyData {
	s.ModelType = &v
	return s
}

func (s *QueryAvatarResponseBodyData) SetName(v string) *QueryAvatarResponseBodyData {
	s.Name = &v
	return s
}

func (s *QueryAvatarResponseBodyData) SetPortrait(v string) *QueryAvatarResponseBodyData {
	s.Portrait = &v
	return s
}

func (s *QueryAvatarResponseBodyData) SetSupportedResolutions(v *QueryAvatarResponseBodyDataSupportedResolutions) *QueryAvatarResponseBodyData {
	s.SupportedResolutions = v
	return s
}

type QueryAvatarResponseBodyDataSupportedResolutions struct {
	Offline []*QueryAvatarResponseBodyDataSupportedResolutionsOffline `json:"Offline,omitempty" xml:"Offline,omitempty" type:"Repeated"`
	Online  []*QueryAvatarResponseBodyDataSupportedResolutionsOnline  `json:"Online,omitempty" xml:"Online,omitempty" type:"Repeated"`
}

func (s QueryAvatarResponseBodyDataSupportedResolutions) String() string {
	return tea.Prettify(s)
}

func (s QueryAvatarResponseBodyDataSupportedResolutions) GoString() string {
	return s.String()
}

func (s *QueryAvatarResponseBodyDataSupportedResolutions) SetOffline(v []*QueryAvatarResponseBodyDataSupportedResolutionsOffline) *QueryAvatarResponseBodyDataSupportedResolutions {
	s.Offline = v
	return s
}

func (s *QueryAvatarResponseBodyDataSupportedResolutions) SetOnline(v []*QueryAvatarResponseBodyDataSupportedResolutionsOnline) *QueryAvatarResponseBodyDataSupportedResolutions {
	s.Online = v
	return s
}

type QueryAvatarResponseBodyDataSupportedResolutionsOffline struct {
	Desc   *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	Height *int32  `json:"Height,omitempty" xml:"Height,omitempty"`
	Width  *int32  `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s QueryAvatarResponseBodyDataSupportedResolutionsOffline) String() string {
	return tea.Prettify(s)
}

func (s QueryAvatarResponseBodyDataSupportedResolutionsOffline) GoString() string {
	return s.String()
}

func (s *QueryAvatarResponseBodyDataSupportedResolutionsOffline) SetDesc(v string) *QueryAvatarResponseBodyDataSupportedResolutionsOffline {
	s.Desc = &v
	return s
}

func (s *QueryAvatarResponseBodyDataSupportedResolutionsOffline) SetHeight(v int32) *QueryAvatarResponseBodyDataSupportedResolutionsOffline {
	s.Height = &v
	return s
}

func (s *QueryAvatarResponseBodyDataSupportedResolutionsOffline) SetWidth(v int32) *QueryAvatarResponseBodyDataSupportedResolutionsOffline {
	s.Width = &v
	return s
}

type QueryAvatarResponseBodyDataSupportedResolutionsOnline struct {
	Desc   *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	Height *int32  `json:"Height,omitempty" xml:"Height,omitempty"`
	Width  *int32  `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s QueryAvatarResponseBodyDataSupportedResolutionsOnline) String() string {
	return tea.Prettify(s)
}

func (s QueryAvatarResponseBodyDataSupportedResolutionsOnline) GoString() string {
	return s.String()
}

func (s *QueryAvatarResponseBodyDataSupportedResolutionsOnline) SetDesc(v string) *QueryAvatarResponseBodyDataSupportedResolutionsOnline {
	s.Desc = &v
	return s
}

func (s *QueryAvatarResponseBodyDataSupportedResolutionsOnline) SetHeight(v int32) *QueryAvatarResponseBodyDataSupportedResolutionsOnline {
	s.Height = &v
	return s
}

func (s *QueryAvatarResponseBodyDataSupportedResolutionsOnline) SetWidth(v int32) *QueryAvatarResponseBodyDataSupportedResolutionsOnline {
	s.Width = &v
	return s
}

type QueryAvatarResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryAvatarResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryAvatarResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryAvatarResponse) GoString() string {
	return s.String()
}

func (s *QueryAvatarResponse) SetHeaders(v map[string]*string) *QueryAvatarResponse {
	s.Headers = v
	return s
}

func (s *QueryAvatarResponse) SetStatusCode(v int32) *QueryAvatarResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryAvatarResponse) SetBody(v *QueryAvatarResponseBody) *QueryAvatarResponse {
	s.Body = v
	return s
}

type QueryAvatarListRequest struct {
	ModelType *string `json:"ModelType,omitempty" xml:"ModelType,omitempty"`
	PageNo    *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TenantId  *int64  `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s QueryAvatarListRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryAvatarListRequest) GoString() string {
	return s.String()
}

func (s *QueryAvatarListRequest) SetModelType(v string) *QueryAvatarListRequest {
	s.ModelType = &v
	return s
}

func (s *QueryAvatarListRequest) SetPageNo(v int32) *QueryAvatarListRequest {
	s.PageNo = &v
	return s
}

func (s *QueryAvatarListRequest) SetPageSize(v int32) *QueryAvatarListRequest {
	s.PageSize = &v
	return s
}

func (s *QueryAvatarListRequest) SetTenantId(v int64) *QueryAvatarListRequest {
	s.TenantId = &v
	return s
}

type QueryAvatarListResponseBody struct {
	Code      *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *QueryAvatarListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                            `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryAvatarListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryAvatarListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAvatarListResponseBody) SetCode(v string) *QueryAvatarListResponseBody {
	s.Code = &v
	return s
}

func (s *QueryAvatarListResponseBody) SetData(v *QueryAvatarListResponseBodyData) *QueryAvatarListResponseBody {
	s.Data = v
	return s
}

func (s *QueryAvatarListResponseBody) SetMessage(v string) *QueryAvatarListResponseBody {
	s.Message = &v
	return s
}

func (s *QueryAvatarListResponseBody) SetRequestId(v string) *QueryAvatarListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryAvatarListResponseBody) SetSuccess(v bool) *QueryAvatarListResponseBody {
	s.Success = &v
	return s
}

type QueryAvatarListResponseBodyData struct {
	List       []*QueryAvatarListResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	PageNo     *int32                                 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize   *int32                                 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *int32                                 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	TotalPage  *int32                                 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s QueryAvatarListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryAvatarListResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryAvatarListResponseBodyData) SetList(v []*QueryAvatarListResponseBodyDataList) *QueryAvatarListResponseBodyData {
	s.List = v
	return s
}

func (s *QueryAvatarListResponseBodyData) SetPageNo(v int32) *QueryAvatarListResponseBodyData {
	s.PageNo = &v
	return s
}

func (s *QueryAvatarListResponseBodyData) SetPageSize(v int32) *QueryAvatarListResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryAvatarListResponseBodyData) SetTotalCount(v int32) *QueryAvatarListResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *QueryAvatarListResponseBodyData) SetTotalPage(v int32) *QueryAvatarListResponseBodyData {
	s.TotalPage = &v
	return s
}

type QueryAvatarListResponseBodyDataList struct {
	AvatarType           *string                                                  `json:"AvatarType,omitempty" xml:"AvatarType,omitempty"`
	Code                 *string                                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Description          *string                                                  `json:"Description,omitempty" xml:"Description,omitempty"`
	Image                *string                                                  `json:"Image,omitempty" xml:"Image,omitempty"`
	MakeFailReason       *string                                                  `json:"MakeFailReason,omitempty" xml:"MakeFailReason,omitempty"`
	MakeStage            *string                                                  `json:"MakeStage,omitempty" xml:"MakeStage,omitempty"`
	MakeStatus           *string                                                  `json:"MakeStatus,omitempty" xml:"MakeStatus,omitempty"`
	ModelType            *string                                                  `json:"ModelType,omitempty" xml:"ModelType,omitempty"`
	Name                 *string                                                  `json:"Name,omitempty" xml:"Name,omitempty"`
	Portrait             *string                                                  `json:"Portrait,omitempty" xml:"Portrait,omitempty"`
	SupportedResolutions *QueryAvatarListResponseBodyDataListSupportedResolutions `json:"SupportedResolutions,omitempty" xml:"SupportedResolutions,omitempty" type:"Struct"`
}

func (s QueryAvatarListResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s QueryAvatarListResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *QueryAvatarListResponseBodyDataList) SetAvatarType(v string) *QueryAvatarListResponseBodyDataList {
	s.AvatarType = &v
	return s
}

func (s *QueryAvatarListResponseBodyDataList) SetCode(v string) *QueryAvatarListResponseBodyDataList {
	s.Code = &v
	return s
}

func (s *QueryAvatarListResponseBodyDataList) SetDescription(v string) *QueryAvatarListResponseBodyDataList {
	s.Description = &v
	return s
}

func (s *QueryAvatarListResponseBodyDataList) SetImage(v string) *QueryAvatarListResponseBodyDataList {
	s.Image = &v
	return s
}

func (s *QueryAvatarListResponseBodyDataList) SetMakeFailReason(v string) *QueryAvatarListResponseBodyDataList {
	s.MakeFailReason = &v
	return s
}

func (s *QueryAvatarListResponseBodyDataList) SetMakeStage(v string) *QueryAvatarListResponseBodyDataList {
	s.MakeStage = &v
	return s
}

func (s *QueryAvatarListResponseBodyDataList) SetMakeStatus(v string) *QueryAvatarListResponseBodyDataList {
	s.MakeStatus = &v
	return s
}

func (s *QueryAvatarListResponseBodyDataList) SetModelType(v string) *QueryAvatarListResponseBodyDataList {
	s.ModelType = &v
	return s
}

func (s *QueryAvatarListResponseBodyDataList) SetName(v string) *QueryAvatarListResponseBodyDataList {
	s.Name = &v
	return s
}

func (s *QueryAvatarListResponseBodyDataList) SetPortrait(v string) *QueryAvatarListResponseBodyDataList {
	s.Portrait = &v
	return s
}

func (s *QueryAvatarListResponseBodyDataList) SetSupportedResolutions(v *QueryAvatarListResponseBodyDataListSupportedResolutions) *QueryAvatarListResponseBodyDataList {
	s.SupportedResolutions = v
	return s
}

type QueryAvatarListResponseBodyDataListSupportedResolutions struct {
	Offline []*QueryAvatarListResponseBodyDataListSupportedResolutionsOffline `json:"Offline,omitempty" xml:"Offline,omitempty" type:"Repeated"`
	Online  []*QueryAvatarListResponseBodyDataListSupportedResolutionsOnline  `json:"Online,omitempty" xml:"Online,omitempty" type:"Repeated"`
}

func (s QueryAvatarListResponseBodyDataListSupportedResolutions) String() string {
	return tea.Prettify(s)
}

func (s QueryAvatarListResponseBodyDataListSupportedResolutions) GoString() string {
	return s.String()
}

func (s *QueryAvatarListResponseBodyDataListSupportedResolutions) SetOffline(v []*QueryAvatarListResponseBodyDataListSupportedResolutionsOffline) *QueryAvatarListResponseBodyDataListSupportedResolutions {
	s.Offline = v
	return s
}

func (s *QueryAvatarListResponseBodyDataListSupportedResolutions) SetOnline(v []*QueryAvatarListResponseBodyDataListSupportedResolutionsOnline) *QueryAvatarListResponseBodyDataListSupportedResolutions {
	s.Online = v
	return s
}

type QueryAvatarListResponseBodyDataListSupportedResolutionsOffline struct {
	Desc   *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	Height *int32  `json:"Height,omitempty" xml:"Height,omitempty"`
	Width  *int32  `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s QueryAvatarListResponseBodyDataListSupportedResolutionsOffline) String() string {
	return tea.Prettify(s)
}

func (s QueryAvatarListResponseBodyDataListSupportedResolutionsOffline) GoString() string {
	return s.String()
}

func (s *QueryAvatarListResponseBodyDataListSupportedResolutionsOffline) SetDesc(v string) *QueryAvatarListResponseBodyDataListSupportedResolutionsOffline {
	s.Desc = &v
	return s
}

func (s *QueryAvatarListResponseBodyDataListSupportedResolutionsOffline) SetHeight(v int32) *QueryAvatarListResponseBodyDataListSupportedResolutionsOffline {
	s.Height = &v
	return s
}

func (s *QueryAvatarListResponseBodyDataListSupportedResolutionsOffline) SetWidth(v int32) *QueryAvatarListResponseBodyDataListSupportedResolutionsOffline {
	s.Width = &v
	return s
}

type QueryAvatarListResponseBodyDataListSupportedResolutionsOnline struct {
	Desc   *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	Height *int32  `json:"Height,omitempty" xml:"Height,omitempty"`
	Width  *int32  `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s QueryAvatarListResponseBodyDataListSupportedResolutionsOnline) String() string {
	return tea.Prettify(s)
}

func (s QueryAvatarListResponseBodyDataListSupportedResolutionsOnline) GoString() string {
	return s.String()
}

func (s *QueryAvatarListResponseBodyDataListSupportedResolutionsOnline) SetDesc(v string) *QueryAvatarListResponseBodyDataListSupportedResolutionsOnline {
	s.Desc = &v
	return s
}

func (s *QueryAvatarListResponseBodyDataListSupportedResolutionsOnline) SetHeight(v int32) *QueryAvatarListResponseBodyDataListSupportedResolutionsOnline {
	s.Height = &v
	return s
}

func (s *QueryAvatarListResponseBodyDataListSupportedResolutionsOnline) SetWidth(v int32) *QueryAvatarListResponseBodyDataListSupportedResolutionsOnline {
	s.Width = &v
	return s
}

type QueryAvatarListResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryAvatarListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryAvatarListResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryAvatarListResponse) GoString() string {
	return s.String()
}

func (s *QueryAvatarListResponse) SetHeaders(v map[string]*string) *QueryAvatarListResponse {
	s.Headers = v
	return s
}

func (s *QueryAvatarListResponse) SetStatusCode(v int32) *QueryAvatarListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryAvatarListResponse) SetBody(v *QueryAvatarListResponseBody) *QueryAvatarListResponse {
	s.Body = v
	return s
}

type QueryRunningInstanceRequest struct {
	App       *QueryRunningInstanceRequestApp `json:"App,omitempty" xml:"App,omitempty" type:"Struct"`
	SessionId *string                         `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	TenantId  *int64                          `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s QueryRunningInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryRunningInstanceRequest) GoString() string {
	return s.String()
}

func (s *QueryRunningInstanceRequest) SetApp(v *QueryRunningInstanceRequestApp) *QueryRunningInstanceRequest {
	s.App = v
	return s
}

func (s *QueryRunningInstanceRequest) SetSessionId(v string) *QueryRunningInstanceRequest {
	s.SessionId = &v
	return s
}

func (s *QueryRunningInstanceRequest) SetTenantId(v int64) *QueryRunningInstanceRequest {
	s.TenantId = &v
	return s
}

type QueryRunningInstanceRequestApp struct {
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s QueryRunningInstanceRequestApp) String() string {
	return tea.Prettify(s)
}

func (s QueryRunningInstanceRequestApp) GoString() string {
	return s.String()
}

func (s *QueryRunningInstanceRequestApp) SetAppId(v string) *QueryRunningInstanceRequestApp {
	s.AppId = &v
	return s
}

type QueryRunningInstanceShrinkRequest struct {
	AppShrink *string `json:"App,omitempty" xml:"App,omitempty"`
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	TenantId  *int64  `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s QueryRunningInstanceShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryRunningInstanceShrinkRequest) GoString() string {
	return s.String()
}

func (s *QueryRunningInstanceShrinkRequest) SetAppShrink(v string) *QueryRunningInstanceShrinkRequest {
	s.AppShrink = &v
	return s
}

func (s *QueryRunningInstanceShrinkRequest) SetSessionId(v string) *QueryRunningInstanceShrinkRequest {
	s.SessionId = &v
	return s
}

func (s *QueryRunningInstanceShrinkRequest) SetTenantId(v int64) *QueryRunningInstanceShrinkRequest {
	s.TenantId = &v
	return s
}

type QueryRunningInstanceResponseBody struct {
	Code      *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*QueryRunningInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryRunningInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryRunningInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *QueryRunningInstanceResponseBody) SetCode(v string) *QueryRunningInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *QueryRunningInstanceResponseBody) SetData(v []*QueryRunningInstanceResponseBodyData) *QueryRunningInstanceResponseBody {
	s.Data = v
	return s
}

func (s *QueryRunningInstanceResponseBody) SetMessage(v string) *QueryRunningInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *QueryRunningInstanceResponseBody) SetRequestId(v string) *QueryRunningInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryRunningInstanceResponseBody) SetSuccess(v bool) *QueryRunningInstanceResponseBody {
	s.Success = &v
	return s
}

type QueryRunningInstanceResponseBodyData struct {
	Channel   *QueryRunningInstanceResponseBodyDataChannel `json:"Channel,omitempty" xml:"Channel,omitempty" type:"Struct"`
	SessionId *string                                      `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	Token     *string                                      `json:"Token,omitempty" xml:"Token,omitempty"`
	User      *QueryRunningInstanceResponseBodyDataUser    `json:"User,omitempty" xml:"User,omitempty" type:"Struct"`
}

func (s QueryRunningInstanceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryRunningInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryRunningInstanceResponseBodyData) SetChannel(v *QueryRunningInstanceResponseBodyDataChannel) *QueryRunningInstanceResponseBodyData {
	s.Channel = v
	return s
}

func (s *QueryRunningInstanceResponseBodyData) SetSessionId(v string) *QueryRunningInstanceResponseBodyData {
	s.SessionId = &v
	return s
}

func (s *QueryRunningInstanceResponseBodyData) SetToken(v string) *QueryRunningInstanceResponseBodyData {
	s.Token = &v
	return s
}

func (s *QueryRunningInstanceResponseBodyData) SetUser(v *QueryRunningInstanceResponseBodyDataUser) *QueryRunningInstanceResponseBodyData {
	s.User = v
	return s
}

type QueryRunningInstanceResponseBodyDataChannel struct {
	AppId             *string   `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ChannelId         *string   `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	ExpiredTime       *string   `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	Gslb              []*string `json:"Gslb,omitempty" xml:"Gslb,omitempty" type:"Repeated"`
	Nonce             *string   `json:"Nonce,omitempty" xml:"Nonce,omitempty"`
	Token             *string   `json:"Token,omitempty" xml:"Token,omitempty"`
	Type              *string   `json:"Type,omitempty" xml:"Type,omitempty"`
	UserId            *string   `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserInfoInChannel *string   `json:"UserInfoInChannel,omitempty" xml:"UserInfoInChannel,omitempty"`
}

func (s QueryRunningInstanceResponseBodyDataChannel) String() string {
	return tea.Prettify(s)
}

func (s QueryRunningInstanceResponseBodyDataChannel) GoString() string {
	return s.String()
}

func (s *QueryRunningInstanceResponseBodyDataChannel) SetAppId(v string) *QueryRunningInstanceResponseBodyDataChannel {
	s.AppId = &v
	return s
}

func (s *QueryRunningInstanceResponseBodyDataChannel) SetChannelId(v string) *QueryRunningInstanceResponseBodyDataChannel {
	s.ChannelId = &v
	return s
}

func (s *QueryRunningInstanceResponseBodyDataChannel) SetExpiredTime(v string) *QueryRunningInstanceResponseBodyDataChannel {
	s.ExpiredTime = &v
	return s
}

func (s *QueryRunningInstanceResponseBodyDataChannel) SetGslb(v []*string) *QueryRunningInstanceResponseBodyDataChannel {
	s.Gslb = v
	return s
}

func (s *QueryRunningInstanceResponseBodyDataChannel) SetNonce(v string) *QueryRunningInstanceResponseBodyDataChannel {
	s.Nonce = &v
	return s
}

func (s *QueryRunningInstanceResponseBodyDataChannel) SetToken(v string) *QueryRunningInstanceResponseBodyDataChannel {
	s.Token = &v
	return s
}

func (s *QueryRunningInstanceResponseBodyDataChannel) SetType(v string) *QueryRunningInstanceResponseBodyDataChannel {
	s.Type = &v
	return s
}

func (s *QueryRunningInstanceResponseBodyDataChannel) SetUserId(v string) *QueryRunningInstanceResponseBodyDataChannel {
	s.UserId = &v
	return s
}

func (s *QueryRunningInstanceResponseBodyDataChannel) SetUserInfoInChannel(v string) *QueryRunningInstanceResponseBodyDataChannel {
	s.UserInfoInChannel = &v
	return s
}

type QueryRunningInstanceResponseBodyDataUser struct {
	UserId   *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s QueryRunningInstanceResponseBodyDataUser) String() string {
	return tea.Prettify(s)
}

func (s QueryRunningInstanceResponseBodyDataUser) GoString() string {
	return s.String()
}

func (s *QueryRunningInstanceResponseBodyDataUser) SetUserId(v string) *QueryRunningInstanceResponseBodyDataUser {
	s.UserId = &v
	return s
}

func (s *QueryRunningInstanceResponseBodyDataUser) SetUserName(v string) *QueryRunningInstanceResponseBodyDataUser {
	s.UserName = &v
	return s
}

type QueryRunningInstanceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryRunningInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryRunningInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryRunningInstanceResponse) GoString() string {
	return s.String()
}

func (s *QueryRunningInstanceResponse) SetHeaders(v map[string]*string) *QueryRunningInstanceResponse {
	s.Headers = v
	return s
}

func (s *QueryRunningInstanceResponse) SetStatusCode(v int32) *QueryRunningInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryRunningInstanceResponse) SetBody(v *QueryRunningInstanceResponseBody) *QueryRunningInstanceResponse {
	s.Body = v
	return s
}

type QueryTimedResetOperateStatusRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	TenantId   *int64  `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s QueryTimedResetOperateStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryTimedResetOperateStatusRequest) GoString() string {
	return s.String()
}

func (s *QueryTimedResetOperateStatusRequest) SetInstanceId(v string) *QueryTimedResetOperateStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryTimedResetOperateStatusRequest) SetTenantId(v int64) *QueryTimedResetOperateStatusRequest {
	s.TenantId = &v
	return s
}

type QueryTimedResetOperateStatusResponseBody struct {
	Code    *string                                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    *QueryTimedResetOperateStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                                       `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryTimedResetOperateStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryTimedResetOperateStatusResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTimedResetOperateStatusResponseBody) SetCode(v string) *QueryTimedResetOperateStatusResponseBody {
	s.Code = &v
	return s
}

func (s *QueryTimedResetOperateStatusResponseBody) SetData(v *QueryTimedResetOperateStatusResponseBodyData) *QueryTimedResetOperateStatusResponseBody {
	s.Data = v
	return s
}

func (s *QueryTimedResetOperateStatusResponseBody) SetMessage(v string) *QueryTimedResetOperateStatusResponseBody {
	s.Message = &v
	return s
}

func (s *QueryTimedResetOperateStatusResponseBody) SetRequestId(v string) *QueryTimedResetOperateStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryTimedResetOperateStatusResponseBody) SetSuccess(v bool) *QueryTimedResetOperateStatusResponseBody {
	s.Success = &v
	return s
}

type QueryTimedResetOperateStatusResponseBodyData struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Status     *int64  `json:"Status,omitempty" xml:"Status,omitempty"`
	StatusStr  *string `json:"StatusStr,omitempty" xml:"StatusStr,omitempty"`
	TenantId   *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s QueryTimedResetOperateStatusResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryTimedResetOperateStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryTimedResetOperateStatusResponseBodyData) SetInstanceId(v string) *QueryTimedResetOperateStatusResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *QueryTimedResetOperateStatusResponseBodyData) SetStatus(v int64) *QueryTimedResetOperateStatusResponseBodyData {
	s.Status = &v
	return s
}

func (s *QueryTimedResetOperateStatusResponseBodyData) SetStatusStr(v string) *QueryTimedResetOperateStatusResponseBodyData {
	s.StatusStr = &v
	return s
}

func (s *QueryTimedResetOperateStatusResponseBodyData) SetTenantId(v string) *QueryTimedResetOperateStatusResponseBodyData {
	s.TenantId = &v
	return s
}

type QueryTimedResetOperateStatusResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryTimedResetOperateStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryTimedResetOperateStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryTimedResetOperateStatusResponse) GoString() string {
	return s.String()
}

func (s *QueryTimedResetOperateStatusResponse) SetHeaders(v map[string]*string) *QueryTimedResetOperateStatusResponse {
	s.Headers = v
	return s
}

func (s *QueryTimedResetOperateStatusResponse) SetStatusCode(v int32) *QueryTimedResetOperateStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryTimedResetOperateStatusResponse) SetBody(v *QueryTimedResetOperateStatusResponseBody) *QueryTimedResetOperateStatusResponse {
	s.Body = v
	return s
}

type QueryVideoTaskInfoRequest struct {
	App       *QueryVideoTaskInfoRequestApp `json:"App,omitempty" xml:"App,omitempty" type:"Struct"`
	OrderById *string                       `json:"OrderById,omitempty" xml:"OrderById,omitempty"`
	PageNo    *int32                        `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize  *int32                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Status    *int32                        `json:"Status,omitempty" xml:"Status,omitempty"`
	TaskUuid  *string                       `json:"TaskUuid,omitempty" xml:"TaskUuid,omitempty"`
	TenantId  *int64                        `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	Title     *string                       `json:"Title,omitempty" xml:"Title,omitempty"`
	Type      *int32                        `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s QueryVideoTaskInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryVideoTaskInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryVideoTaskInfoRequest) SetApp(v *QueryVideoTaskInfoRequestApp) *QueryVideoTaskInfoRequest {
	s.App = v
	return s
}

func (s *QueryVideoTaskInfoRequest) SetOrderById(v string) *QueryVideoTaskInfoRequest {
	s.OrderById = &v
	return s
}

func (s *QueryVideoTaskInfoRequest) SetPageNo(v int32) *QueryVideoTaskInfoRequest {
	s.PageNo = &v
	return s
}

func (s *QueryVideoTaskInfoRequest) SetPageSize(v int32) *QueryVideoTaskInfoRequest {
	s.PageSize = &v
	return s
}

func (s *QueryVideoTaskInfoRequest) SetStatus(v int32) *QueryVideoTaskInfoRequest {
	s.Status = &v
	return s
}

func (s *QueryVideoTaskInfoRequest) SetTaskUuid(v string) *QueryVideoTaskInfoRequest {
	s.TaskUuid = &v
	return s
}

func (s *QueryVideoTaskInfoRequest) SetTenantId(v int64) *QueryVideoTaskInfoRequest {
	s.TenantId = &v
	return s
}

func (s *QueryVideoTaskInfoRequest) SetTitle(v string) *QueryVideoTaskInfoRequest {
	s.Title = &v
	return s
}

func (s *QueryVideoTaskInfoRequest) SetType(v int32) *QueryVideoTaskInfoRequest {
	s.Type = &v
	return s
}

type QueryVideoTaskInfoRequestApp struct {
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s QueryVideoTaskInfoRequestApp) String() string {
	return tea.Prettify(s)
}

func (s QueryVideoTaskInfoRequestApp) GoString() string {
	return s.String()
}

func (s *QueryVideoTaskInfoRequestApp) SetAppId(v string) *QueryVideoTaskInfoRequestApp {
	s.AppId = &v
	return s
}

type QueryVideoTaskInfoShrinkRequest struct {
	AppShrink *string `json:"App,omitempty" xml:"App,omitempty"`
	OrderById *string `json:"OrderById,omitempty" xml:"OrderById,omitempty"`
	PageNo    *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Status    *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	TaskUuid  *string `json:"TaskUuid,omitempty" xml:"TaskUuid,omitempty"`
	TenantId  *int64  `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	Title     *string `json:"Title,omitempty" xml:"Title,omitempty"`
	Type      *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s QueryVideoTaskInfoShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryVideoTaskInfoShrinkRequest) GoString() string {
	return s.String()
}

func (s *QueryVideoTaskInfoShrinkRequest) SetAppShrink(v string) *QueryVideoTaskInfoShrinkRequest {
	s.AppShrink = &v
	return s
}

func (s *QueryVideoTaskInfoShrinkRequest) SetOrderById(v string) *QueryVideoTaskInfoShrinkRequest {
	s.OrderById = &v
	return s
}

func (s *QueryVideoTaskInfoShrinkRequest) SetPageNo(v int32) *QueryVideoTaskInfoShrinkRequest {
	s.PageNo = &v
	return s
}

func (s *QueryVideoTaskInfoShrinkRequest) SetPageSize(v int32) *QueryVideoTaskInfoShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *QueryVideoTaskInfoShrinkRequest) SetStatus(v int32) *QueryVideoTaskInfoShrinkRequest {
	s.Status = &v
	return s
}

func (s *QueryVideoTaskInfoShrinkRequest) SetTaskUuid(v string) *QueryVideoTaskInfoShrinkRequest {
	s.TaskUuid = &v
	return s
}

func (s *QueryVideoTaskInfoShrinkRequest) SetTenantId(v int64) *QueryVideoTaskInfoShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *QueryVideoTaskInfoShrinkRequest) SetTitle(v string) *QueryVideoTaskInfoShrinkRequest {
	s.Title = &v
	return s
}

func (s *QueryVideoTaskInfoShrinkRequest) SetType(v int32) *QueryVideoTaskInfoShrinkRequest {
	s.Type = &v
	return s
}

type QueryVideoTaskInfoResponseBody struct {
	Code    *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    *QueryVideoTaskInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryVideoTaskInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryVideoTaskInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryVideoTaskInfoResponseBody) SetCode(v string) *QueryVideoTaskInfoResponseBody {
	s.Code = &v
	return s
}

func (s *QueryVideoTaskInfoResponseBody) SetData(v *QueryVideoTaskInfoResponseBodyData) *QueryVideoTaskInfoResponseBody {
	s.Data = v
	return s
}

func (s *QueryVideoTaskInfoResponseBody) SetMessage(v string) *QueryVideoTaskInfoResponseBody {
	s.Message = &v
	return s
}

func (s *QueryVideoTaskInfoResponseBody) SetRequestId(v string) *QueryVideoTaskInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryVideoTaskInfoResponseBody) SetSuccess(v bool) *QueryVideoTaskInfoResponseBody {
	s.Success = &v
	return s
}

type QueryVideoTaskInfoResponseBodyData struct {
	List       []*QueryVideoTaskInfoResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	PageNo     *int32                                    `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize   *int32                                    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *int64                                    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s QueryVideoTaskInfoResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryVideoTaskInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryVideoTaskInfoResponseBodyData) SetList(v []*QueryVideoTaskInfoResponseBodyDataList) *QueryVideoTaskInfoResponseBodyData {
	s.List = v
	return s
}

func (s *QueryVideoTaskInfoResponseBodyData) SetPageNo(v int32) *QueryVideoTaskInfoResponseBodyData {
	s.PageNo = &v
	return s
}

func (s *QueryVideoTaskInfoResponseBodyData) SetPageSize(v int32) *QueryVideoTaskInfoResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryVideoTaskInfoResponseBodyData) SetTotalCount(v int64) *QueryVideoTaskInfoResponseBodyData {
	s.TotalCount = &v
	return s
}

type QueryVideoTaskInfoResponseBodyDataList struct {
	Status     *int32                                            `json:"Status,omitempty" xml:"Status,omitempty"`
	TaskResult *QueryVideoTaskInfoResponseBodyDataListTaskResult `json:"TaskResult,omitempty" xml:"TaskResult,omitempty" type:"Struct"`
	TaskUuid   *string                                           `json:"TaskUuid,omitempty" xml:"TaskUuid,omitempty"`
	Title      *string                                           `json:"Title,omitempty" xml:"Title,omitempty"`
	Type       *int32                                            `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s QueryVideoTaskInfoResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s QueryVideoTaskInfoResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *QueryVideoTaskInfoResponseBodyDataList) SetStatus(v int32) *QueryVideoTaskInfoResponseBodyDataList {
	s.Status = &v
	return s
}

func (s *QueryVideoTaskInfoResponseBodyDataList) SetTaskResult(v *QueryVideoTaskInfoResponseBodyDataListTaskResult) *QueryVideoTaskInfoResponseBodyDataList {
	s.TaskResult = v
	return s
}

func (s *QueryVideoTaskInfoResponseBodyDataList) SetTaskUuid(v string) *QueryVideoTaskInfoResponseBodyDataList {
	s.TaskUuid = &v
	return s
}

func (s *QueryVideoTaskInfoResponseBodyDataList) SetTitle(v string) *QueryVideoTaskInfoResponseBodyDataList {
	s.Title = &v
	return s
}

func (s *QueryVideoTaskInfoResponseBodyDataList) SetType(v int32) *QueryVideoTaskInfoResponseBodyDataList {
	s.Type = &v
	return s
}

type QueryVideoTaskInfoResponseBodyDataListTaskResult struct {
	AlphaUrl         *string `json:"AlphaUrl,omitempty" xml:"AlphaUrl,omitempty"`
	AttachmentUrl    *string `json:"AttachmentUrl,omitempty" xml:"AttachmentUrl,omitempty"`
	FailCode         *string `json:"FailCode,omitempty" xml:"FailCode,omitempty"`
	FailReason       *string `json:"FailReason,omitempty" xml:"FailReason,omitempty"`
	PreviewPic       *string `json:"PreviewPic,omitempty" xml:"PreviewPic,omitempty"`
	SubtitlesUrl     *string `json:"SubtitlesUrl,omitempty" xml:"SubtitlesUrl,omitempty"`
	VideoDuration    *int32  `json:"VideoDuration,omitempty" xml:"VideoDuration,omitempty"`
	VideoUrl         *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
	WordSubtitlesUrl *string `json:"WordSubtitlesUrl,omitempty" xml:"WordSubtitlesUrl,omitempty"`
}

func (s QueryVideoTaskInfoResponseBodyDataListTaskResult) String() string {
	return tea.Prettify(s)
}

func (s QueryVideoTaskInfoResponseBodyDataListTaskResult) GoString() string {
	return s.String()
}

func (s *QueryVideoTaskInfoResponseBodyDataListTaskResult) SetAlphaUrl(v string) *QueryVideoTaskInfoResponseBodyDataListTaskResult {
	s.AlphaUrl = &v
	return s
}

func (s *QueryVideoTaskInfoResponseBodyDataListTaskResult) SetAttachmentUrl(v string) *QueryVideoTaskInfoResponseBodyDataListTaskResult {
	s.AttachmentUrl = &v
	return s
}

func (s *QueryVideoTaskInfoResponseBodyDataListTaskResult) SetFailCode(v string) *QueryVideoTaskInfoResponseBodyDataListTaskResult {
	s.FailCode = &v
	return s
}

func (s *QueryVideoTaskInfoResponseBodyDataListTaskResult) SetFailReason(v string) *QueryVideoTaskInfoResponseBodyDataListTaskResult {
	s.FailReason = &v
	return s
}

func (s *QueryVideoTaskInfoResponseBodyDataListTaskResult) SetPreviewPic(v string) *QueryVideoTaskInfoResponseBodyDataListTaskResult {
	s.PreviewPic = &v
	return s
}

func (s *QueryVideoTaskInfoResponseBodyDataListTaskResult) SetSubtitlesUrl(v string) *QueryVideoTaskInfoResponseBodyDataListTaskResult {
	s.SubtitlesUrl = &v
	return s
}

func (s *QueryVideoTaskInfoResponseBodyDataListTaskResult) SetVideoDuration(v int32) *QueryVideoTaskInfoResponseBodyDataListTaskResult {
	s.VideoDuration = &v
	return s
}

func (s *QueryVideoTaskInfoResponseBodyDataListTaskResult) SetVideoUrl(v string) *QueryVideoTaskInfoResponseBodyDataListTaskResult {
	s.VideoUrl = &v
	return s
}

func (s *QueryVideoTaskInfoResponseBodyDataListTaskResult) SetWordSubtitlesUrl(v string) *QueryVideoTaskInfoResponseBodyDataListTaskResult {
	s.WordSubtitlesUrl = &v
	return s
}

type QueryVideoTaskInfoResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryVideoTaskInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryVideoTaskInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryVideoTaskInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryVideoTaskInfoResponse) SetHeaders(v map[string]*string) *QueryVideoTaskInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryVideoTaskInfoResponse) SetStatusCode(v int32) *QueryVideoTaskInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryVideoTaskInfoResponse) SetBody(v *QueryVideoTaskInfoResponseBody) *QueryVideoTaskInfoResponse {
	s.Body = v
	return s
}

type Render3dAvatarRequest struct {
	AppId    *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Code     *string `json:"Code,omitempty" xml:"Code,omitempty"`
	TenantId *int64  `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s Render3dAvatarRequest) String() string {
	return tea.Prettify(s)
}

func (s Render3dAvatarRequest) GoString() string {
	return s.String()
}

func (s *Render3dAvatarRequest) SetAppId(v string) *Render3dAvatarRequest {
	s.AppId = &v
	return s
}

func (s *Render3dAvatarRequest) SetCode(v string) *Render3dAvatarRequest {
	s.Code = &v
	return s
}

func (s *Render3dAvatarRequest) SetTenantId(v int64) *Render3dAvatarRequest {
	s.TenantId = &v
	return s
}

type Render3dAvatarResponseBody struct {
	Code      *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *Render3dAvatarResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                           `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s Render3dAvatarResponseBody) String() string {
	return tea.Prettify(s)
}

func (s Render3dAvatarResponseBody) GoString() string {
	return s.String()
}

func (s *Render3dAvatarResponseBody) SetCode(v string) *Render3dAvatarResponseBody {
	s.Code = &v
	return s
}

func (s *Render3dAvatarResponseBody) SetData(v *Render3dAvatarResponseBodyData) *Render3dAvatarResponseBody {
	s.Data = v
	return s
}

func (s *Render3dAvatarResponseBody) SetMessage(v string) *Render3dAvatarResponseBody {
	s.Message = &v
	return s
}

func (s *Render3dAvatarResponseBody) SetRequestId(v string) *Render3dAvatarResponseBody {
	s.RequestId = &v
	return s
}

func (s *Render3dAvatarResponseBody) SetSuccess(v bool) *Render3dAvatarResponseBody {
	s.Success = &v
	return s
}

type Render3dAvatarResponseBodyData struct {
	RenderData *string `json:"RenderData,omitempty" xml:"RenderData,omitempty"`
}

func (s Render3dAvatarResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s Render3dAvatarResponseBodyData) GoString() string {
	return s.String()
}

func (s *Render3dAvatarResponseBodyData) SetRenderData(v string) *Render3dAvatarResponseBodyData {
	s.RenderData = &v
	return s
}

type Render3dAvatarResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *Render3dAvatarResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s Render3dAvatarResponse) String() string {
	return tea.Prettify(s)
}

func (s Render3dAvatarResponse) GoString() string {
	return s.String()
}

func (s *Render3dAvatarResponse) SetHeaders(v map[string]*string) *Render3dAvatarResponse {
	s.Headers = v
	return s
}

func (s *Render3dAvatarResponse) SetStatusCode(v int32) *Render3dAvatarResponse {
	s.StatusCode = &v
	return s
}

func (s *Render3dAvatarResponse) SetBody(v *Render3dAvatarResponseBody) *Render3dAvatarResponse {
	s.Body = v
	return s
}

type SendCommandRequest struct {
	Code       *string                `json:"Code,omitempty" xml:"Code,omitempty"`
	Content    map[string]interface{} `json:"Content,omitempty" xml:"Content,omitempty"`
	Feedback   *bool                  `json:"Feedback,omitempty" xml:"Feedback,omitempty"`
	SessionId  *string                `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	TenantId   *int64                 `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	UniqueCode *string                `json:"UniqueCode,omitempty" xml:"UniqueCode,omitempty"`
}

func (s SendCommandRequest) String() string {
	return tea.Prettify(s)
}

func (s SendCommandRequest) GoString() string {
	return s.String()
}

func (s *SendCommandRequest) SetCode(v string) *SendCommandRequest {
	s.Code = &v
	return s
}

func (s *SendCommandRequest) SetContent(v map[string]interface{}) *SendCommandRequest {
	s.Content = v
	return s
}

func (s *SendCommandRequest) SetFeedback(v bool) *SendCommandRequest {
	s.Feedback = &v
	return s
}

func (s *SendCommandRequest) SetSessionId(v string) *SendCommandRequest {
	s.SessionId = &v
	return s
}

func (s *SendCommandRequest) SetTenantId(v int64) *SendCommandRequest {
	s.TenantId = &v
	return s
}

func (s *SendCommandRequest) SetUniqueCode(v string) *SendCommandRequest {
	s.UniqueCode = &v
	return s
}

type SendCommandShrinkRequest struct {
	Code          *string `json:"Code,omitempty" xml:"Code,omitempty"`
	ContentShrink *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Feedback      *bool   `json:"Feedback,omitempty" xml:"Feedback,omitempty"`
	SessionId     *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	TenantId      *int64  `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	UniqueCode    *string `json:"UniqueCode,omitempty" xml:"UniqueCode,omitempty"`
}

func (s SendCommandShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s SendCommandShrinkRequest) GoString() string {
	return s.String()
}

func (s *SendCommandShrinkRequest) SetCode(v string) *SendCommandShrinkRequest {
	s.Code = &v
	return s
}

func (s *SendCommandShrinkRequest) SetContentShrink(v string) *SendCommandShrinkRequest {
	s.ContentShrink = &v
	return s
}

func (s *SendCommandShrinkRequest) SetFeedback(v bool) *SendCommandShrinkRequest {
	s.Feedback = &v
	return s
}

func (s *SendCommandShrinkRequest) SetSessionId(v string) *SendCommandShrinkRequest {
	s.SessionId = &v
	return s
}

func (s *SendCommandShrinkRequest) SetTenantId(v int64) *SendCommandShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *SendCommandShrinkRequest) SetUniqueCode(v string) *SendCommandShrinkRequest {
	s.UniqueCode = &v
	return s
}

type SendCommandResponseBody struct {
	Code    *string                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    *SendCommandResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                      `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SendCommandResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendCommandResponseBody) GoString() string {
	return s.String()
}

func (s *SendCommandResponseBody) SetCode(v string) *SendCommandResponseBody {
	s.Code = &v
	return s
}

func (s *SendCommandResponseBody) SetData(v *SendCommandResponseBodyData) *SendCommandResponseBody {
	s.Data = v
	return s
}

func (s *SendCommandResponseBody) SetMessage(v string) *SendCommandResponseBody {
	s.Message = &v
	return s
}

func (s *SendCommandResponseBody) SetRequestId(v string) *SendCommandResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendCommandResponseBody) SetSuccess(v bool) *SendCommandResponseBody {
	s.Success = &v
	return s
}

type SendCommandResponseBodyData struct {
	SessionId  *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	UniqueCode *string `json:"UniqueCode,omitempty" xml:"UniqueCode,omitempty"`
}

func (s SendCommandResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SendCommandResponseBodyData) GoString() string {
	return s.String()
}

func (s *SendCommandResponseBodyData) SetSessionId(v string) *SendCommandResponseBodyData {
	s.SessionId = &v
	return s
}

func (s *SendCommandResponseBodyData) SetUniqueCode(v string) *SendCommandResponseBodyData {
	s.UniqueCode = &v
	return s
}

type SendCommandResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SendCommandResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SendCommandResponse) String() string {
	return tea.Prettify(s)
}

func (s SendCommandResponse) GoString() string {
	return s.String()
}

func (s *SendCommandResponse) SetHeaders(v map[string]*string) *SendCommandResponse {
	s.Headers = v
	return s
}

func (s *SendCommandResponse) SetStatusCode(v int32) *SendCommandResponse {
	s.StatusCode = &v
	return s
}

func (s *SendCommandResponse) SetBody(v *SendCommandResponseBody) *SendCommandResponse {
	s.Body = v
	return s
}

type SendMessageRequest struct {
	Feedback        *bool                              `json:"Feedback,omitempty" xml:"Feedback,omitempty"`
	SessionId       *string                            `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	StreamExtension *SendMessageRequestStreamExtension `json:"StreamExtension,omitempty" xml:"StreamExtension,omitempty" type:"Struct"`
	TenantId        *int64                             `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	TextRequest     *SendMessageRequestTextRequest     `json:"TextRequest,omitempty" xml:"TextRequest,omitempty" type:"Struct"`
	VAMLRequest     *SendMessageRequestVAMLRequest     `json:"VAMLRequest,omitempty" xml:"VAMLRequest,omitempty" type:"Struct"`
}

func (s SendMessageRequest) String() string {
	return tea.Prettify(s)
}

func (s SendMessageRequest) GoString() string {
	return s.String()
}

func (s *SendMessageRequest) SetFeedback(v bool) *SendMessageRequest {
	s.Feedback = &v
	return s
}

func (s *SendMessageRequest) SetSessionId(v string) *SendMessageRequest {
	s.SessionId = &v
	return s
}

func (s *SendMessageRequest) SetStreamExtension(v *SendMessageRequestStreamExtension) *SendMessageRequest {
	s.StreamExtension = v
	return s
}

func (s *SendMessageRequest) SetTenantId(v int64) *SendMessageRequest {
	s.TenantId = &v
	return s
}

func (s *SendMessageRequest) SetTextRequest(v *SendMessageRequestTextRequest) *SendMessageRequest {
	s.TextRequest = v
	return s
}

func (s *SendMessageRequest) SetVAMLRequest(v *SendMessageRequestVAMLRequest) *SendMessageRequest {
	s.VAMLRequest = v
	return s
}

type SendMessageRequestStreamExtension struct {
	Index    *int32  `json:"Index,omitempty" xml:"Index,omitempty"`
	IsStream *bool   `json:"IsStream,omitempty" xml:"IsStream,omitempty"`
	Position *string `json:"Position,omitempty" xml:"Position,omitempty"`
}

func (s SendMessageRequestStreamExtension) String() string {
	return tea.Prettify(s)
}

func (s SendMessageRequestStreamExtension) GoString() string {
	return s.String()
}

func (s *SendMessageRequestStreamExtension) SetIndex(v int32) *SendMessageRequestStreamExtension {
	s.Index = &v
	return s
}

func (s *SendMessageRequestStreamExtension) SetIsStream(v bool) *SendMessageRequestStreamExtension {
	s.IsStream = &v
	return s
}

func (s *SendMessageRequestStreamExtension) SetPosition(v string) *SendMessageRequestStreamExtension {
	s.Position = &v
	return s
}

type SendMessageRequestTextRequest struct {
	CommandType *string `json:"CommandType,omitempty" xml:"CommandType,omitempty"`
	Id          *string `json:"Id,omitempty" xml:"Id,omitempty"`
	SpeechText  *string `json:"SpeechText,omitempty" xml:"SpeechText,omitempty"`
	Interrupt   *bool   `json:"interrupt,omitempty" xml:"interrupt,omitempty"`
}

func (s SendMessageRequestTextRequest) String() string {
	return tea.Prettify(s)
}

func (s SendMessageRequestTextRequest) GoString() string {
	return s.String()
}

func (s *SendMessageRequestTextRequest) SetCommandType(v string) *SendMessageRequestTextRequest {
	s.CommandType = &v
	return s
}

func (s *SendMessageRequestTextRequest) SetId(v string) *SendMessageRequestTextRequest {
	s.Id = &v
	return s
}

func (s *SendMessageRequestTextRequest) SetSpeechText(v string) *SendMessageRequestTextRequest {
	s.SpeechText = &v
	return s
}

func (s *SendMessageRequestTextRequest) SetInterrupt(v bool) *SendMessageRequestTextRequest {
	s.Interrupt = &v
	return s
}

type SendMessageRequestVAMLRequest struct {
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Vaml *string `json:"Vaml,omitempty" xml:"Vaml,omitempty"`
}

func (s SendMessageRequestVAMLRequest) String() string {
	return tea.Prettify(s)
}

func (s SendMessageRequestVAMLRequest) GoString() string {
	return s.String()
}

func (s *SendMessageRequestVAMLRequest) SetCode(v string) *SendMessageRequestVAMLRequest {
	s.Code = &v
	return s
}

func (s *SendMessageRequestVAMLRequest) SetVaml(v string) *SendMessageRequestVAMLRequest {
	s.Vaml = &v
	return s
}

type SendMessageShrinkRequest struct {
	Feedback              *bool   `json:"Feedback,omitempty" xml:"Feedback,omitempty"`
	SessionId             *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	StreamExtensionShrink *string `json:"StreamExtension,omitempty" xml:"StreamExtension,omitempty"`
	TenantId              *int64  `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	TextRequestShrink     *string `json:"TextRequest,omitempty" xml:"TextRequest,omitempty"`
	VAMLRequestShrink     *string `json:"VAMLRequest,omitempty" xml:"VAMLRequest,omitempty"`
}

func (s SendMessageShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s SendMessageShrinkRequest) GoString() string {
	return s.String()
}

func (s *SendMessageShrinkRequest) SetFeedback(v bool) *SendMessageShrinkRequest {
	s.Feedback = &v
	return s
}

func (s *SendMessageShrinkRequest) SetSessionId(v string) *SendMessageShrinkRequest {
	s.SessionId = &v
	return s
}

func (s *SendMessageShrinkRequest) SetStreamExtensionShrink(v string) *SendMessageShrinkRequest {
	s.StreamExtensionShrink = &v
	return s
}

func (s *SendMessageShrinkRequest) SetTenantId(v int64) *SendMessageShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *SendMessageShrinkRequest) SetTextRequestShrink(v string) *SendMessageShrinkRequest {
	s.TextRequestShrink = &v
	return s
}

func (s *SendMessageShrinkRequest) SetVAMLRequestShrink(v string) *SendMessageShrinkRequest {
	s.VAMLRequestShrink = &v
	return s
}

type SendMessageResponseBody struct {
	// Id of the request
	Code    *string                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    *SendMessageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                      `json:"Message,omitempty" xml:"Message,omitempty"`
	Success *bool                        `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SendMessageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendMessageResponseBody) GoString() string {
	return s.String()
}

func (s *SendMessageResponseBody) SetCode(v string) *SendMessageResponseBody {
	s.Code = &v
	return s
}

func (s *SendMessageResponseBody) SetData(v *SendMessageResponseBodyData) *SendMessageResponseBody {
	s.Data = v
	return s
}

func (s *SendMessageResponseBody) SetMessage(v string) *SendMessageResponseBody {
	s.Message = &v
	return s
}

func (s *SendMessageResponseBody) SetSuccess(v bool) *SendMessageResponseBody {
	s.Success = &v
	return s
}

type SendMessageResponseBodyData struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s SendMessageResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SendMessageResponseBodyData) GoString() string {
	return s.String()
}

func (s *SendMessageResponseBodyData) SetRequestId(v string) *SendMessageResponseBodyData {
	s.RequestId = &v
	return s
}

func (s *SendMessageResponseBodyData) SetSessionId(v string) *SendMessageResponseBodyData {
	s.SessionId = &v
	return s
}

type SendMessageResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SendMessageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SendMessageResponse) String() string {
	return tea.Prettify(s)
}

func (s SendMessageResponse) GoString() string {
	return s.String()
}

func (s *SendMessageResponse) SetHeaders(v map[string]*string) *SendMessageResponse {
	s.Headers = v
	return s
}

func (s *SendMessageResponse) SetStatusCode(v int32) *SendMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *SendMessageResponse) SetBody(v *SendMessageResponseBody) *SendMessageResponse {
	s.Body = v
	return s
}

type SendTextRequest struct {
	Feedback        *bool                           `json:"Feedback,omitempty" xml:"Feedback,omitempty"`
	Interrupt       *bool                           `json:"Interrupt,omitempty" xml:"Interrupt,omitempty"`
	SessionId       *string                         `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	StreamExtension *SendTextRequestStreamExtension `json:"StreamExtension,omitempty" xml:"StreamExtension,omitempty" type:"Struct"`
	TenantId        *int64                          `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	Text            *string                         `json:"Text,omitempty" xml:"Text,omitempty"`
	UniqueCode      *string                         `json:"UniqueCode,omitempty" xml:"UniqueCode,omitempty"`
}

func (s SendTextRequest) String() string {
	return tea.Prettify(s)
}

func (s SendTextRequest) GoString() string {
	return s.String()
}

func (s *SendTextRequest) SetFeedback(v bool) *SendTextRequest {
	s.Feedback = &v
	return s
}

func (s *SendTextRequest) SetInterrupt(v bool) *SendTextRequest {
	s.Interrupt = &v
	return s
}

func (s *SendTextRequest) SetSessionId(v string) *SendTextRequest {
	s.SessionId = &v
	return s
}

func (s *SendTextRequest) SetStreamExtension(v *SendTextRequestStreamExtension) *SendTextRequest {
	s.StreamExtension = v
	return s
}

func (s *SendTextRequest) SetTenantId(v int64) *SendTextRequest {
	s.TenantId = &v
	return s
}

func (s *SendTextRequest) SetText(v string) *SendTextRequest {
	s.Text = &v
	return s
}

func (s *SendTextRequest) SetUniqueCode(v string) *SendTextRequest {
	s.UniqueCode = &v
	return s
}

type SendTextRequestStreamExtension struct {
	Index    *int32  `json:"Index,omitempty" xml:"Index,omitempty"`
	IsStream *bool   `json:"IsStream,omitempty" xml:"IsStream,omitempty"`
	Position *string `json:"Position,omitempty" xml:"Position,omitempty"`
}

func (s SendTextRequestStreamExtension) String() string {
	return tea.Prettify(s)
}

func (s SendTextRequestStreamExtension) GoString() string {
	return s.String()
}

func (s *SendTextRequestStreamExtension) SetIndex(v int32) *SendTextRequestStreamExtension {
	s.Index = &v
	return s
}

func (s *SendTextRequestStreamExtension) SetIsStream(v bool) *SendTextRequestStreamExtension {
	s.IsStream = &v
	return s
}

func (s *SendTextRequestStreamExtension) SetPosition(v string) *SendTextRequestStreamExtension {
	s.Position = &v
	return s
}

type SendTextShrinkRequest struct {
	Feedback              *bool   `json:"Feedback,omitempty" xml:"Feedback,omitempty"`
	Interrupt             *bool   `json:"Interrupt,omitempty" xml:"Interrupt,omitempty"`
	SessionId             *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	StreamExtensionShrink *string `json:"StreamExtension,omitempty" xml:"StreamExtension,omitempty"`
	TenantId              *int64  `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	Text                  *string `json:"Text,omitempty" xml:"Text,omitempty"`
	UniqueCode            *string `json:"UniqueCode,omitempty" xml:"UniqueCode,omitempty"`
}

func (s SendTextShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s SendTextShrinkRequest) GoString() string {
	return s.String()
}

func (s *SendTextShrinkRequest) SetFeedback(v bool) *SendTextShrinkRequest {
	s.Feedback = &v
	return s
}

func (s *SendTextShrinkRequest) SetInterrupt(v bool) *SendTextShrinkRequest {
	s.Interrupt = &v
	return s
}

func (s *SendTextShrinkRequest) SetSessionId(v string) *SendTextShrinkRequest {
	s.SessionId = &v
	return s
}

func (s *SendTextShrinkRequest) SetStreamExtensionShrink(v string) *SendTextShrinkRequest {
	s.StreamExtensionShrink = &v
	return s
}

func (s *SendTextShrinkRequest) SetTenantId(v int64) *SendTextShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *SendTextShrinkRequest) SetText(v string) *SendTextShrinkRequest {
	s.Text = &v
	return s
}

func (s *SendTextShrinkRequest) SetUniqueCode(v string) *SendTextShrinkRequest {
	s.UniqueCode = &v
	return s
}

type SendTextResponseBody struct {
	Code    *string                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    *SendTextResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                   `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SendTextResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendTextResponseBody) GoString() string {
	return s.String()
}

func (s *SendTextResponseBody) SetCode(v string) *SendTextResponseBody {
	s.Code = &v
	return s
}

func (s *SendTextResponseBody) SetData(v *SendTextResponseBodyData) *SendTextResponseBody {
	s.Data = v
	return s
}

func (s *SendTextResponseBody) SetMessage(v string) *SendTextResponseBody {
	s.Message = &v
	return s
}

func (s *SendTextResponseBody) SetRequestId(v string) *SendTextResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendTextResponseBody) SetSuccess(v bool) *SendTextResponseBody {
	s.Success = &v
	return s
}

type SendTextResponseBodyData struct {
	SessionId  *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	UniqueCode *string `json:"UniqueCode,omitempty" xml:"UniqueCode,omitempty"`
}

func (s SendTextResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SendTextResponseBodyData) GoString() string {
	return s.String()
}

func (s *SendTextResponseBodyData) SetSessionId(v string) *SendTextResponseBodyData {
	s.SessionId = &v
	return s
}

func (s *SendTextResponseBodyData) SetUniqueCode(v string) *SendTextResponseBodyData {
	s.UniqueCode = &v
	return s
}

type SendTextResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SendTextResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SendTextResponse) String() string {
	return tea.Prettify(s)
}

func (s SendTextResponse) GoString() string {
	return s.String()
}

func (s *SendTextResponse) SetHeaders(v map[string]*string) *SendTextResponse {
	s.Headers = v
	return s
}

func (s *SendTextResponse) SetStatusCode(v int32) *SendTextResponse {
	s.StatusCode = &v
	return s
}

func (s *SendTextResponse) SetBody(v *SendTextResponseBody) *SendTextResponse {
	s.Body = v
	return s
}

type SendVamlRequest struct {
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	TenantId  *int64  `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	Vaml      *string `json:"Vaml,omitempty" xml:"Vaml,omitempty"`
}

func (s SendVamlRequest) String() string {
	return tea.Prettify(s)
}

func (s SendVamlRequest) GoString() string {
	return s.String()
}

func (s *SendVamlRequest) SetSessionId(v string) *SendVamlRequest {
	s.SessionId = &v
	return s
}

func (s *SendVamlRequest) SetTenantId(v int64) *SendVamlRequest {
	s.TenantId = &v
	return s
}

func (s *SendVamlRequest) SetVaml(v string) *SendVamlRequest {
	s.Vaml = &v
	return s
}

type SendVamlResponseBody struct {
	Code    *string                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    *SendVamlResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                   `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SendVamlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendVamlResponseBody) GoString() string {
	return s.String()
}

func (s *SendVamlResponseBody) SetCode(v string) *SendVamlResponseBody {
	s.Code = &v
	return s
}

func (s *SendVamlResponseBody) SetData(v *SendVamlResponseBodyData) *SendVamlResponseBody {
	s.Data = v
	return s
}

func (s *SendVamlResponseBody) SetMessage(v string) *SendVamlResponseBody {
	s.Message = &v
	return s
}

func (s *SendVamlResponseBody) SetRequestId(v string) *SendVamlResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendVamlResponseBody) SetSuccess(v bool) *SendVamlResponseBody {
	s.Success = &v
	return s
}

type SendVamlResponseBodyData struct {
	SessionId  *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	UniqueCode *string `json:"UniqueCode,omitempty" xml:"UniqueCode,omitempty"`
}

func (s SendVamlResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SendVamlResponseBodyData) GoString() string {
	return s.String()
}

func (s *SendVamlResponseBodyData) SetSessionId(v string) *SendVamlResponseBodyData {
	s.SessionId = &v
	return s
}

func (s *SendVamlResponseBodyData) SetUniqueCode(v string) *SendVamlResponseBodyData {
	s.UniqueCode = &v
	return s
}

type SendVamlResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SendVamlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SendVamlResponse) String() string {
	return tea.Prettify(s)
}

func (s SendVamlResponse) GoString() string {
	return s.String()
}

func (s *SendVamlResponse) SetHeaders(v map[string]*string) *SendVamlResponse {
	s.Headers = v
	return s
}

func (s *SendVamlResponse) SetStatusCode(v int32) *SendVamlResponse {
	s.StatusCode = &v
	return s
}

func (s *SendVamlResponse) SetBody(v *SendVamlResponseBody) *SendVamlResponse {
	s.Body = v
	return s
}

type StartInstanceRequest struct {
	App            *StartInstanceRequestApp            `json:"App,omitempty" xml:"App,omitempty" type:"Struct"`
	BizId          *string                             `json:"BizId,omitempty" xml:"BizId,omitempty"`
	Channel        *StartInstanceRequestChannel        `json:"Channel,omitempty" xml:"Channel,omitempty" type:"Struct"`
	CommandRequest *StartInstanceRequestCommandRequest `json:"CommandRequest,omitempty" xml:"CommandRequest,omitempty" type:"Struct"`
	TenantId       *int64                              `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	User           *StartInstanceRequestUser           `json:"User,omitempty" xml:"User,omitempty" type:"Struct"`
}

func (s StartInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s StartInstanceRequest) GoString() string {
	return s.String()
}

func (s *StartInstanceRequest) SetApp(v *StartInstanceRequestApp) *StartInstanceRequest {
	s.App = v
	return s
}

func (s *StartInstanceRequest) SetBizId(v string) *StartInstanceRequest {
	s.BizId = &v
	return s
}

func (s *StartInstanceRequest) SetChannel(v *StartInstanceRequestChannel) *StartInstanceRequest {
	s.Channel = v
	return s
}

func (s *StartInstanceRequest) SetCommandRequest(v *StartInstanceRequestCommandRequest) *StartInstanceRequest {
	s.CommandRequest = v
	return s
}

func (s *StartInstanceRequest) SetTenantId(v int64) *StartInstanceRequest {
	s.TenantId = &v
	return s
}

func (s *StartInstanceRequest) SetUser(v *StartInstanceRequestUser) *StartInstanceRequest {
	s.User = v
	return s
}

type StartInstanceRequestApp struct {
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s StartInstanceRequestApp) String() string {
	return tea.Prettify(s)
}

func (s StartInstanceRequestApp) GoString() string {
	return s.String()
}

func (s *StartInstanceRequestApp) SetAppId(v string) *StartInstanceRequestApp {
	s.AppId = &v
	return s
}

type StartInstanceRequestChannel struct {
	ReqConfig map[string]interface{} `json:"ReqConfig,omitempty" xml:"ReqConfig,omitempty"`
	Type      *string                `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s StartInstanceRequestChannel) String() string {
	return tea.Prettify(s)
}

func (s StartInstanceRequestChannel) GoString() string {
	return s.String()
}

func (s *StartInstanceRequestChannel) SetReqConfig(v map[string]interface{}) *StartInstanceRequestChannel {
	s.ReqConfig = v
	return s
}

func (s *StartInstanceRequestChannel) SetType(v string) *StartInstanceRequestChannel {
	s.Type = &v
	return s
}

type StartInstanceRequestCommandRequest struct {
	AlphaSwitch *bool `json:"AlphaSwitch,omitempty" xml:"AlphaSwitch,omitempty"`
}

func (s StartInstanceRequestCommandRequest) String() string {
	return tea.Prettify(s)
}

func (s StartInstanceRequestCommandRequest) GoString() string {
	return s.String()
}

func (s *StartInstanceRequestCommandRequest) SetAlphaSwitch(v bool) *StartInstanceRequestCommandRequest {
	s.AlphaSwitch = &v
	return s
}

type StartInstanceRequestUser struct {
	UserId   *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s StartInstanceRequestUser) String() string {
	return tea.Prettify(s)
}

func (s StartInstanceRequestUser) GoString() string {
	return s.String()
}

func (s *StartInstanceRequestUser) SetUserId(v string) *StartInstanceRequestUser {
	s.UserId = &v
	return s
}

func (s *StartInstanceRequestUser) SetUserName(v string) *StartInstanceRequestUser {
	s.UserName = &v
	return s
}

type StartInstanceShrinkRequest struct {
	AppShrink            *string `json:"App,omitempty" xml:"App,omitempty"`
	BizId                *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	ChannelShrink        *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
	CommandRequestShrink *string `json:"CommandRequest,omitempty" xml:"CommandRequest,omitempty"`
	TenantId             *int64  `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	UserShrink           *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s StartInstanceShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s StartInstanceShrinkRequest) GoString() string {
	return s.String()
}

func (s *StartInstanceShrinkRequest) SetAppShrink(v string) *StartInstanceShrinkRequest {
	s.AppShrink = &v
	return s
}

func (s *StartInstanceShrinkRequest) SetBizId(v string) *StartInstanceShrinkRequest {
	s.BizId = &v
	return s
}

func (s *StartInstanceShrinkRequest) SetChannelShrink(v string) *StartInstanceShrinkRequest {
	s.ChannelShrink = &v
	return s
}

func (s *StartInstanceShrinkRequest) SetCommandRequestShrink(v string) *StartInstanceShrinkRequest {
	s.CommandRequestShrink = &v
	return s
}

func (s *StartInstanceShrinkRequest) SetTenantId(v int64) *StartInstanceShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *StartInstanceShrinkRequest) SetUserShrink(v string) *StartInstanceShrinkRequest {
	s.UserShrink = &v
	return s
}

type StartInstanceResponseBody struct {
	Code    *string                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    *StartInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                        `json:"Message,omitempty" xml:"Message,omitempty"`
	Success *bool                          `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StartInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *StartInstanceResponseBody) SetCode(v string) *StartInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *StartInstanceResponseBody) SetData(v *StartInstanceResponseBodyData) *StartInstanceResponseBody {
	s.Data = v
	return s
}

func (s *StartInstanceResponseBody) SetMessage(v string) *StartInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *StartInstanceResponseBody) SetSuccess(v bool) *StartInstanceResponseBody {
	s.Success = &v
	return s
}

type StartInstanceResponseBodyData struct {
	Channel   *StartInstanceResponseBodyDataChannel `json:"Channel,omitempty" xml:"Channel,omitempty" type:"Struct"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SessionId *string                               `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	Token     *string                               `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s StartInstanceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s StartInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *StartInstanceResponseBodyData) SetChannel(v *StartInstanceResponseBodyDataChannel) *StartInstanceResponseBodyData {
	s.Channel = v
	return s
}

func (s *StartInstanceResponseBodyData) SetRequestId(v string) *StartInstanceResponseBodyData {
	s.RequestId = &v
	return s
}

func (s *StartInstanceResponseBodyData) SetSessionId(v string) *StartInstanceResponseBodyData {
	s.SessionId = &v
	return s
}

func (s *StartInstanceResponseBodyData) SetToken(v string) *StartInstanceResponseBodyData {
	s.Token = &v
	return s
}

type StartInstanceResponseBodyDataChannel struct {
	AppId             *string   `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ChannelId         *string   `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	ExpiredTime       *string   `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	Gslb              []*string `json:"Gslb,omitempty" xml:"Gslb,omitempty" type:"Repeated"`
	Nonce             *string   `json:"Nonce,omitempty" xml:"Nonce,omitempty"`
	Token             *string   `json:"Token,omitempty" xml:"Token,omitempty"`
	Type              *string   `json:"Type,omitempty" xml:"Type,omitempty"`
	UserId            *string   `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserInfoInChannel *string   `json:"UserInfoInChannel,omitempty" xml:"UserInfoInChannel,omitempty"`
}

func (s StartInstanceResponseBodyDataChannel) String() string {
	return tea.Prettify(s)
}

func (s StartInstanceResponseBodyDataChannel) GoString() string {
	return s.String()
}

func (s *StartInstanceResponseBodyDataChannel) SetAppId(v string) *StartInstanceResponseBodyDataChannel {
	s.AppId = &v
	return s
}

func (s *StartInstanceResponseBodyDataChannel) SetChannelId(v string) *StartInstanceResponseBodyDataChannel {
	s.ChannelId = &v
	return s
}

func (s *StartInstanceResponseBodyDataChannel) SetExpiredTime(v string) *StartInstanceResponseBodyDataChannel {
	s.ExpiredTime = &v
	return s
}

func (s *StartInstanceResponseBodyDataChannel) SetGslb(v []*string) *StartInstanceResponseBodyDataChannel {
	s.Gslb = v
	return s
}

func (s *StartInstanceResponseBodyDataChannel) SetNonce(v string) *StartInstanceResponseBodyDataChannel {
	s.Nonce = &v
	return s
}

func (s *StartInstanceResponseBodyDataChannel) SetToken(v string) *StartInstanceResponseBodyDataChannel {
	s.Token = &v
	return s
}

func (s *StartInstanceResponseBodyDataChannel) SetType(v string) *StartInstanceResponseBodyDataChannel {
	s.Type = &v
	return s
}

func (s *StartInstanceResponseBodyDataChannel) SetUserId(v string) *StartInstanceResponseBodyDataChannel {
	s.UserId = &v
	return s
}

func (s *StartInstanceResponseBodyDataChannel) SetUserInfoInChannel(v string) *StartInstanceResponseBodyDataChannel {
	s.UserInfoInChannel = &v
	return s
}

type StartInstanceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StartInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s StartInstanceResponse) GoString() string {
	return s.String()
}

func (s *StartInstanceResponse) SetHeaders(v map[string]*string) *StartInstanceResponse {
	s.Headers = v
	return s
}

func (s *StartInstanceResponse) SetStatusCode(v int32) *StartInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *StartInstanceResponse) SetBody(v *StartInstanceResponseBody) *StartInstanceResponse {
	s.Body = v
	return s
}

type StartTimedResetOperateRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	TenantId   *int64  `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s StartTimedResetOperateRequest) String() string {
	return tea.Prettify(s)
}

func (s StartTimedResetOperateRequest) GoString() string {
	return s.String()
}

func (s *StartTimedResetOperateRequest) SetInstanceId(v string) *StartTimedResetOperateRequest {
	s.InstanceId = &v
	return s
}

func (s *StartTimedResetOperateRequest) SetTenantId(v int64) *StartTimedResetOperateRequest {
	s.TenantId = &v
	return s
}

type StartTimedResetOperateResponseBody struct {
	Code      *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *StartTimedResetOperateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StartTimedResetOperateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartTimedResetOperateResponseBody) GoString() string {
	return s.String()
}

func (s *StartTimedResetOperateResponseBody) SetCode(v string) *StartTimedResetOperateResponseBody {
	s.Code = &v
	return s
}

func (s *StartTimedResetOperateResponseBody) SetData(v *StartTimedResetOperateResponseBodyData) *StartTimedResetOperateResponseBody {
	s.Data = v
	return s
}

func (s *StartTimedResetOperateResponseBody) SetMessage(v string) *StartTimedResetOperateResponseBody {
	s.Message = &v
	return s
}

func (s *StartTimedResetOperateResponseBody) SetRequestId(v string) *StartTimedResetOperateResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartTimedResetOperateResponseBody) SetSuccess(v bool) *StartTimedResetOperateResponseBody {
	s.Success = &v
	return s
}

type StartTimedResetOperateResponseBodyData struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	TenantId   *int64  `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s StartTimedResetOperateResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s StartTimedResetOperateResponseBodyData) GoString() string {
	return s.String()
}

func (s *StartTimedResetOperateResponseBodyData) SetInstanceId(v string) *StartTimedResetOperateResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *StartTimedResetOperateResponseBodyData) SetTenantId(v int64) *StartTimedResetOperateResponseBodyData {
	s.TenantId = &v
	return s
}

type StartTimedResetOperateResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StartTimedResetOperateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartTimedResetOperateResponse) String() string {
	return tea.Prettify(s)
}

func (s StartTimedResetOperateResponse) GoString() string {
	return s.String()
}

func (s *StartTimedResetOperateResponse) SetHeaders(v map[string]*string) *StartTimedResetOperateResponse {
	s.Headers = v
	return s
}

func (s *StartTimedResetOperateResponse) SetStatusCode(v int32) *StartTimedResetOperateResponse {
	s.StatusCode = &v
	return s
}

func (s *StartTimedResetOperateResponse) SetBody(v *StartTimedResetOperateResponseBody) *StartTimedResetOperateResponse {
	s.Body = v
	return s
}

type StopInstanceRequest struct {
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	TenantId  *int64  `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s StopInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s StopInstanceRequest) GoString() string {
	return s.String()
}

func (s *StopInstanceRequest) SetSessionId(v string) *StopInstanceRequest {
	s.SessionId = &v
	return s
}

func (s *StopInstanceRequest) SetTenantId(v int64) *StopInstanceRequest {
	s.TenantId = &v
	return s
}

type StopInstanceResponseBody struct {
	Code    *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    *StopInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                       `json:"Message,omitempty" xml:"Message,omitempty"`
	Success *bool                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StopInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *StopInstanceResponseBody) SetCode(v string) *StopInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *StopInstanceResponseBody) SetData(v *StopInstanceResponseBodyData) *StopInstanceResponseBody {
	s.Data = v
	return s
}

func (s *StopInstanceResponseBody) SetMessage(v string) *StopInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *StopInstanceResponseBody) SetSuccess(v bool) *StopInstanceResponseBody {
	s.Success = &v
	return s
}

type StopInstanceResponseBodyData struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s StopInstanceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s StopInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *StopInstanceResponseBodyData) SetRequestId(v string) *StopInstanceResponseBodyData {
	s.RequestId = &v
	return s
}

func (s *StopInstanceResponseBodyData) SetSessionId(v string) *StopInstanceResponseBodyData {
	s.SessionId = &v
	return s
}

type StopInstanceResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StopInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s StopInstanceResponse) GoString() string {
	return s.String()
}

func (s *StopInstanceResponse) SetHeaders(v map[string]*string) *StopInstanceResponse {
	s.Headers = v
	return s
}

func (s *StopInstanceResponse) SetStatusCode(v int32) *StopInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *StopInstanceResponse) SetBody(v *StopInstanceResponseBody) *StopInstanceResponse {
	s.Body = v
	return s
}

type SubmitAudioTo2DAvatarVideoTaskRequest struct {
	App            *SubmitAudioTo2DAvatarVideoTaskRequestApp        `json:"App,omitempty" xml:"App,omitempty" type:"Struct"`
	AvatarInfo     *SubmitAudioTo2DAvatarVideoTaskRequestAvatarInfo `json:"AvatarInfo,omitempty" xml:"AvatarInfo,omitempty" type:"Struct"`
	Callback       *bool                                            `json:"Callback,omitempty" xml:"Callback,omitempty"`
	CallbackParams *string                                          `json:"CallbackParams,omitempty" xml:"CallbackParams,omitempty"`
	ExtParams      *string                                          `json:"ExtParams,omitempty" xml:"ExtParams,omitempty"`
	TenantId       *int64                                           `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	Title          *string                                          `json:"Title,omitempty" xml:"Title,omitempty"`
	Url            *string                                          `json:"Url,omitempty" xml:"Url,omitempty"`
	VideoInfo      *SubmitAudioTo2DAvatarVideoTaskRequestVideoInfo  `json:"VideoInfo,omitempty" xml:"VideoInfo,omitempty" type:"Struct"`
}

func (s SubmitAudioTo2DAvatarVideoTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitAudioTo2DAvatarVideoTaskRequest) GoString() string {
	return s.String()
}

func (s *SubmitAudioTo2DAvatarVideoTaskRequest) SetApp(v *SubmitAudioTo2DAvatarVideoTaskRequestApp) *SubmitAudioTo2DAvatarVideoTaskRequest {
	s.App = v
	return s
}

func (s *SubmitAudioTo2DAvatarVideoTaskRequest) SetAvatarInfo(v *SubmitAudioTo2DAvatarVideoTaskRequestAvatarInfo) *SubmitAudioTo2DAvatarVideoTaskRequest {
	s.AvatarInfo = v
	return s
}

func (s *SubmitAudioTo2DAvatarVideoTaskRequest) SetCallback(v bool) *SubmitAudioTo2DAvatarVideoTaskRequest {
	s.Callback = &v
	return s
}

func (s *SubmitAudioTo2DAvatarVideoTaskRequest) SetCallbackParams(v string) *SubmitAudioTo2DAvatarVideoTaskRequest {
	s.CallbackParams = &v
	return s
}

func (s *SubmitAudioTo2DAvatarVideoTaskRequest) SetExtParams(v string) *SubmitAudioTo2DAvatarVideoTaskRequest {
	s.ExtParams = &v
	return s
}

func (s *SubmitAudioTo2DAvatarVideoTaskRequest) SetTenantId(v int64) *SubmitAudioTo2DAvatarVideoTaskRequest {
	s.TenantId = &v
	return s
}

func (s *SubmitAudioTo2DAvatarVideoTaskRequest) SetTitle(v string) *SubmitAudioTo2DAvatarVideoTaskRequest {
	s.Title = &v
	return s
}

func (s *SubmitAudioTo2DAvatarVideoTaskRequest) SetUrl(v string) *SubmitAudioTo2DAvatarVideoTaskRequest {
	s.Url = &v
	return s
}

func (s *SubmitAudioTo2DAvatarVideoTaskRequest) SetVideoInfo(v *SubmitAudioTo2DAvatarVideoTaskRequestVideoInfo) *SubmitAudioTo2DAvatarVideoTaskRequest {
	s.VideoInfo = v
	return s
}

type SubmitAudioTo2DAvatarVideoTaskRequestApp struct {
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s SubmitAudioTo2DAvatarVideoTaskRequestApp) String() string {
	return tea.Prettify(s)
}

func (s SubmitAudioTo2DAvatarVideoTaskRequestApp) GoString() string {
	return s.String()
}

func (s *SubmitAudioTo2DAvatarVideoTaskRequestApp) SetAppId(v string) *SubmitAudioTo2DAvatarVideoTaskRequestApp {
	s.AppId = &v
	return s
}

type SubmitAudioTo2DAvatarVideoTaskRequestAvatarInfo struct {
	Code   *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Height *int32  `json:"Height,omitempty" xml:"Height,omitempty"`
	Width  *int32  `json:"Width,omitempty" xml:"Width,omitempty"`
	X      *int32  `json:"X,omitempty" xml:"X,omitempty"`
	Y      *int32  `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s SubmitAudioTo2DAvatarVideoTaskRequestAvatarInfo) String() string {
	return tea.Prettify(s)
}

func (s SubmitAudioTo2DAvatarVideoTaskRequestAvatarInfo) GoString() string {
	return s.String()
}

func (s *SubmitAudioTo2DAvatarVideoTaskRequestAvatarInfo) SetCode(v string) *SubmitAudioTo2DAvatarVideoTaskRequestAvatarInfo {
	s.Code = &v
	return s
}

func (s *SubmitAudioTo2DAvatarVideoTaskRequestAvatarInfo) SetHeight(v int32) *SubmitAudioTo2DAvatarVideoTaskRequestAvatarInfo {
	s.Height = &v
	return s
}

func (s *SubmitAudioTo2DAvatarVideoTaskRequestAvatarInfo) SetWidth(v int32) *SubmitAudioTo2DAvatarVideoTaskRequestAvatarInfo {
	s.Width = &v
	return s
}

func (s *SubmitAudioTo2DAvatarVideoTaskRequestAvatarInfo) SetX(v int32) *SubmitAudioTo2DAvatarVideoTaskRequestAvatarInfo {
	s.X = &v
	return s
}

func (s *SubmitAudioTo2DAvatarVideoTaskRequestAvatarInfo) SetY(v int32) *SubmitAudioTo2DAvatarVideoTaskRequestAvatarInfo {
	s.Y = &v
	return s
}

type SubmitAudioTo2DAvatarVideoTaskRequestVideoInfo struct {
	AlphaFormat        *int32  `json:"AlphaFormat,omitempty" xml:"AlphaFormat,omitempty"`
	BackgroundImageUrl *string `json:"BackgroundImageUrl,omitempty" xml:"BackgroundImageUrl,omitempty"`
	IsAlpha            *bool   `json:"IsAlpha,omitempty" xml:"IsAlpha,omitempty"`
	Resolution         *int32  `json:"Resolution,omitempty" xml:"Resolution,omitempty"`
}

func (s SubmitAudioTo2DAvatarVideoTaskRequestVideoInfo) String() string {
	return tea.Prettify(s)
}

func (s SubmitAudioTo2DAvatarVideoTaskRequestVideoInfo) GoString() string {
	return s.String()
}

func (s *SubmitAudioTo2DAvatarVideoTaskRequestVideoInfo) SetAlphaFormat(v int32) *SubmitAudioTo2DAvatarVideoTaskRequestVideoInfo {
	s.AlphaFormat = &v
	return s
}

func (s *SubmitAudioTo2DAvatarVideoTaskRequestVideoInfo) SetBackgroundImageUrl(v string) *SubmitAudioTo2DAvatarVideoTaskRequestVideoInfo {
	s.BackgroundImageUrl = &v
	return s
}

func (s *SubmitAudioTo2DAvatarVideoTaskRequestVideoInfo) SetIsAlpha(v bool) *SubmitAudioTo2DAvatarVideoTaskRequestVideoInfo {
	s.IsAlpha = &v
	return s
}

func (s *SubmitAudioTo2DAvatarVideoTaskRequestVideoInfo) SetResolution(v int32) *SubmitAudioTo2DAvatarVideoTaskRequestVideoInfo {
	s.Resolution = &v
	return s
}

type SubmitAudioTo2DAvatarVideoTaskShrinkRequest struct {
	AppShrink        *string `json:"App,omitempty" xml:"App,omitempty"`
	AvatarInfoShrink *string `json:"AvatarInfo,omitempty" xml:"AvatarInfo,omitempty"`
	Callback         *bool   `json:"Callback,omitempty" xml:"Callback,omitempty"`
	CallbackParams   *string `json:"CallbackParams,omitempty" xml:"CallbackParams,omitempty"`
	ExtParams        *string `json:"ExtParams,omitempty" xml:"ExtParams,omitempty"`
	TenantId         *int64  `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	Title            *string `json:"Title,omitempty" xml:"Title,omitempty"`
	Url              *string `json:"Url,omitempty" xml:"Url,omitempty"`
	VideoInfoShrink  *string `json:"VideoInfo,omitempty" xml:"VideoInfo,omitempty"`
}

func (s SubmitAudioTo2DAvatarVideoTaskShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitAudioTo2DAvatarVideoTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *SubmitAudioTo2DAvatarVideoTaskShrinkRequest) SetAppShrink(v string) *SubmitAudioTo2DAvatarVideoTaskShrinkRequest {
	s.AppShrink = &v
	return s
}

func (s *SubmitAudioTo2DAvatarVideoTaskShrinkRequest) SetAvatarInfoShrink(v string) *SubmitAudioTo2DAvatarVideoTaskShrinkRequest {
	s.AvatarInfoShrink = &v
	return s
}

func (s *SubmitAudioTo2DAvatarVideoTaskShrinkRequest) SetCallback(v bool) *SubmitAudioTo2DAvatarVideoTaskShrinkRequest {
	s.Callback = &v
	return s
}

func (s *SubmitAudioTo2DAvatarVideoTaskShrinkRequest) SetCallbackParams(v string) *SubmitAudioTo2DAvatarVideoTaskShrinkRequest {
	s.CallbackParams = &v
	return s
}

func (s *SubmitAudioTo2DAvatarVideoTaskShrinkRequest) SetExtParams(v string) *SubmitAudioTo2DAvatarVideoTaskShrinkRequest {
	s.ExtParams = &v
	return s
}

func (s *SubmitAudioTo2DAvatarVideoTaskShrinkRequest) SetTenantId(v int64) *SubmitAudioTo2DAvatarVideoTaskShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *SubmitAudioTo2DAvatarVideoTaskShrinkRequest) SetTitle(v string) *SubmitAudioTo2DAvatarVideoTaskShrinkRequest {
	s.Title = &v
	return s
}

func (s *SubmitAudioTo2DAvatarVideoTaskShrinkRequest) SetUrl(v string) *SubmitAudioTo2DAvatarVideoTaskShrinkRequest {
	s.Url = &v
	return s
}

func (s *SubmitAudioTo2DAvatarVideoTaskShrinkRequest) SetVideoInfoShrink(v string) *SubmitAudioTo2DAvatarVideoTaskShrinkRequest {
	s.VideoInfoShrink = &v
	return s
}

type SubmitAudioTo2DAvatarVideoTaskResponseBody struct {
	Code      *string                                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *SubmitAudioTo2DAvatarVideoTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                           `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SubmitAudioTo2DAvatarVideoTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitAudioTo2DAvatarVideoTaskResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitAudioTo2DAvatarVideoTaskResponseBody) SetCode(v string) *SubmitAudioTo2DAvatarVideoTaskResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitAudioTo2DAvatarVideoTaskResponseBody) SetData(v *SubmitAudioTo2DAvatarVideoTaskResponseBodyData) *SubmitAudioTo2DAvatarVideoTaskResponseBody {
	s.Data = v
	return s
}

func (s *SubmitAudioTo2DAvatarVideoTaskResponseBody) SetMessage(v string) *SubmitAudioTo2DAvatarVideoTaskResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitAudioTo2DAvatarVideoTaskResponseBody) SetRequestId(v string) *SubmitAudioTo2DAvatarVideoTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitAudioTo2DAvatarVideoTaskResponseBody) SetSuccess(v bool) *SubmitAudioTo2DAvatarVideoTaskResponseBody {
	s.Success = &v
	return s
}

type SubmitAudioTo2DAvatarVideoTaskResponseBodyData struct {
	TaskUuid *string `json:"TaskUuid,omitempty" xml:"TaskUuid,omitempty"`
}

func (s SubmitAudioTo2DAvatarVideoTaskResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SubmitAudioTo2DAvatarVideoTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitAudioTo2DAvatarVideoTaskResponseBodyData) SetTaskUuid(v string) *SubmitAudioTo2DAvatarVideoTaskResponseBodyData {
	s.TaskUuid = &v
	return s
}

type SubmitAudioTo2DAvatarVideoTaskResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SubmitAudioTo2DAvatarVideoTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubmitAudioTo2DAvatarVideoTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitAudioTo2DAvatarVideoTaskResponse) GoString() string {
	return s.String()
}

func (s *SubmitAudioTo2DAvatarVideoTaskResponse) SetHeaders(v map[string]*string) *SubmitAudioTo2DAvatarVideoTaskResponse {
	s.Headers = v
	return s
}

func (s *SubmitAudioTo2DAvatarVideoTaskResponse) SetStatusCode(v int32) *SubmitAudioTo2DAvatarVideoTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitAudioTo2DAvatarVideoTaskResponse) SetBody(v *SubmitAudioTo2DAvatarVideoTaskResponseBody) *SubmitAudioTo2DAvatarVideoTaskResponse {
	s.Body = v
	return s
}

type SubmitAudioTo3DAvatarVideoTaskRequest struct {
	App            *SubmitAudioTo3DAvatarVideoTaskRequestApp        `json:"App,omitempty" xml:"App,omitempty" type:"Struct"`
	AvatarInfo     *SubmitAudioTo3DAvatarVideoTaskRequestAvatarInfo `json:"AvatarInfo,omitempty" xml:"AvatarInfo,omitempty" type:"Struct"`
	Callback       *bool                                            `json:"Callback,omitempty" xml:"Callback,omitempty"`
	CallbackParams *string                                          `json:"CallbackParams,omitempty" xml:"CallbackParams,omitempty"`
	ExtParams      *string                                          `json:"ExtParams,omitempty" xml:"ExtParams,omitempty"`
	TenantId       *int64                                           `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	Title          *string                                          `json:"Title,omitempty" xml:"Title,omitempty"`
	Url            *string                                          `json:"Url,omitempty" xml:"Url,omitempty"`
	VideoInfo      *SubmitAudioTo3DAvatarVideoTaskRequestVideoInfo  `json:"VideoInfo,omitempty" xml:"VideoInfo,omitempty" type:"Struct"`
}

func (s SubmitAudioTo3DAvatarVideoTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitAudioTo3DAvatarVideoTaskRequest) GoString() string {
	return s.String()
}

func (s *SubmitAudioTo3DAvatarVideoTaskRequest) SetApp(v *SubmitAudioTo3DAvatarVideoTaskRequestApp) *SubmitAudioTo3DAvatarVideoTaskRequest {
	s.App = v
	return s
}

func (s *SubmitAudioTo3DAvatarVideoTaskRequest) SetAvatarInfo(v *SubmitAudioTo3DAvatarVideoTaskRequestAvatarInfo) *SubmitAudioTo3DAvatarVideoTaskRequest {
	s.AvatarInfo = v
	return s
}

func (s *SubmitAudioTo3DAvatarVideoTaskRequest) SetCallback(v bool) *SubmitAudioTo3DAvatarVideoTaskRequest {
	s.Callback = &v
	return s
}

func (s *SubmitAudioTo3DAvatarVideoTaskRequest) SetCallbackParams(v string) *SubmitAudioTo3DAvatarVideoTaskRequest {
	s.CallbackParams = &v
	return s
}

func (s *SubmitAudioTo3DAvatarVideoTaskRequest) SetExtParams(v string) *SubmitAudioTo3DAvatarVideoTaskRequest {
	s.ExtParams = &v
	return s
}

func (s *SubmitAudioTo3DAvatarVideoTaskRequest) SetTenantId(v int64) *SubmitAudioTo3DAvatarVideoTaskRequest {
	s.TenantId = &v
	return s
}

func (s *SubmitAudioTo3DAvatarVideoTaskRequest) SetTitle(v string) *SubmitAudioTo3DAvatarVideoTaskRequest {
	s.Title = &v
	return s
}

func (s *SubmitAudioTo3DAvatarVideoTaskRequest) SetUrl(v string) *SubmitAudioTo3DAvatarVideoTaskRequest {
	s.Url = &v
	return s
}

func (s *SubmitAudioTo3DAvatarVideoTaskRequest) SetVideoInfo(v *SubmitAudioTo3DAvatarVideoTaskRequestVideoInfo) *SubmitAudioTo3DAvatarVideoTaskRequest {
	s.VideoInfo = v
	return s
}

type SubmitAudioTo3DAvatarVideoTaskRequestApp struct {
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s SubmitAudioTo3DAvatarVideoTaskRequestApp) String() string {
	return tea.Prettify(s)
}

func (s SubmitAudioTo3DAvatarVideoTaskRequestApp) GoString() string {
	return s.String()
}

func (s *SubmitAudioTo3DAvatarVideoTaskRequestApp) SetAppId(v string) *SubmitAudioTo3DAvatarVideoTaskRequestApp {
	s.AppId = &v
	return s
}

type SubmitAudioTo3DAvatarVideoTaskRequestAvatarInfo struct {
	Angle        *int32  `json:"Angle,omitempty" xml:"Angle,omitempty"`
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	IndustryCode *string `json:"IndustryCode,omitempty" xml:"IndustryCode,omitempty"`
	Locate       *int32  `json:"Locate,omitempty" xml:"Locate,omitempty"`
}

func (s SubmitAudioTo3DAvatarVideoTaskRequestAvatarInfo) String() string {
	return tea.Prettify(s)
}

func (s SubmitAudioTo3DAvatarVideoTaskRequestAvatarInfo) GoString() string {
	return s.String()
}

func (s *SubmitAudioTo3DAvatarVideoTaskRequestAvatarInfo) SetAngle(v int32) *SubmitAudioTo3DAvatarVideoTaskRequestAvatarInfo {
	s.Angle = &v
	return s
}

func (s *SubmitAudioTo3DAvatarVideoTaskRequestAvatarInfo) SetCode(v string) *SubmitAudioTo3DAvatarVideoTaskRequestAvatarInfo {
	s.Code = &v
	return s
}

func (s *SubmitAudioTo3DAvatarVideoTaskRequestAvatarInfo) SetIndustryCode(v string) *SubmitAudioTo3DAvatarVideoTaskRequestAvatarInfo {
	s.IndustryCode = &v
	return s
}

func (s *SubmitAudioTo3DAvatarVideoTaskRequestAvatarInfo) SetLocate(v int32) *SubmitAudioTo3DAvatarVideoTaskRequestAvatarInfo {
	s.Locate = &v
	return s
}

type SubmitAudioTo3DAvatarVideoTaskRequestVideoInfo struct {
	AlphaFormat        *int32  `json:"AlphaFormat,omitempty" xml:"AlphaFormat,omitempty"`
	BackgroundImageUrl *string `json:"BackgroundImageUrl,omitempty" xml:"BackgroundImageUrl,omitempty"`
	IsAlpha            *bool   `json:"IsAlpha,omitempty" xml:"IsAlpha,omitempty"`
	Resolution         *int32  `json:"Resolution,omitempty" xml:"Resolution,omitempty"`
}

func (s SubmitAudioTo3DAvatarVideoTaskRequestVideoInfo) String() string {
	return tea.Prettify(s)
}

func (s SubmitAudioTo3DAvatarVideoTaskRequestVideoInfo) GoString() string {
	return s.String()
}

func (s *SubmitAudioTo3DAvatarVideoTaskRequestVideoInfo) SetAlphaFormat(v int32) *SubmitAudioTo3DAvatarVideoTaskRequestVideoInfo {
	s.AlphaFormat = &v
	return s
}

func (s *SubmitAudioTo3DAvatarVideoTaskRequestVideoInfo) SetBackgroundImageUrl(v string) *SubmitAudioTo3DAvatarVideoTaskRequestVideoInfo {
	s.BackgroundImageUrl = &v
	return s
}

func (s *SubmitAudioTo3DAvatarVideoTaskRequestVideoInfo) SetIsAlpha(v bool) *SubmitAudioTo3DAvatarVideoTaskRequestVideoInfo {
	s.IsAlpha = &v
	return s
}

func (s *SubmitAudioTo3DAvatarVideoTaskRequestVideoInfo) SetResolution(v int32) *SubmitAudioTo3DAvatarVideoTaskRequestVideoInfo {
	s.Resolution = &v
	return s
}

type SubmitAudioTo3DAvatarVideoTaskShrinkRequest struct {
	AppShrink        *string `json:"App,omitempty" xml:"App,omitempty"`
	AvatarInfoShrink *string `json:"AvatarInfo,omitempty" xml:"AvatarInfo,omitempty"`
	Callback         *bool   `json:"Callback,omitempty" xml:"Callback,omitempty"`
	CallbackParams   *string `json:"CallbackParams,omitempty" xml:"CallbackParams,omitempty"`
	ExtParams        *string `json:"ExtParams,omitempty" xml:"ExtParams,omitempty"`
	TenantId         *int64  `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	Title            *string `json:"Title,omitempty" xml:"Title,omitempty"`
	Url              *string `json:"Url,omitempty" xml:"Url,omitempty"`
	VideoInfoShrink  *string `json:"VideoInfo,omitempty" xml:"VideoInfo,omitempty"`
}

func (s SubmitAudioTo3DAvatarVideoTaskShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitAudioTo3DAvatarVideoTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *SubmitAudioTo3DAvatarVideoTaskShrinkRequest) SetAppShrink(v string) *SubmitAudioTo3DAvatarVideoTaskShrinkRequest {
	s.AppShrink = &v
	return s
}

func (s *SubmitAudioTo3DAvatarVideoTaskShrinkRequest) SetAvatarInfoShrink(v string) *SubmitAudioTo3DAvatarVideoTaskShrinkRequest {
	s.AvatarInfoShrink = &v
	return s
}

func (s *SubmitAudioTo3DAvatarVideoTaskShrinkRequest) SetCallback(v bool) *SubmitAudioTo3DAvatarVideoTaskShrinkRequest {
	s.Callback = &v
	return s
}

func (s *SubmitAudioTo3DAvatarVideoTaskShrinkRequest) SetCallbackParams(v string) *SubmitAudioTo3DAvatarVideoTaskShrinkRequest {
	s.CallbackParams = &v
	return s
}

func (s *SubmitAudioTo3DAvatarVideoTaskShrinkRequest) SetExtParams(v string) *SubmitAudioTo3DAvatarVideoTaskShrinkRequest {
	s.ExtParams = &v
	return s
}

func (s *SubmitAudioTo3DAvatarVideoTaskShrinkRequest) SetTenantId(v int64) *SubmitAudioTo3DAvatarVideoTaskShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *SubmitAudioTo3DAvatarVideoTaskShrinkRequest) SetTitle(v string) *SubmitAudioTo3DAvatarVideoTaskShrinkRequest {
	s.Title = &v
	return s
}

func (s *SubmitAudioTo3DAvatarVideoTaskShrinkRequest) SetUrl(v string) *SubmitAudioTo3DAvatarVideoTaskShrinkRequest {
	s.Url = &v
	return s
}

func (s *SubmitAudioTo3DAvatarVideoTaskShrinkRequest) SetVideoInfoShrink(v string) *SubmitAudioTo3DAvatarVideoTaskShrinkRequest {
	s.VideoInfoShrink = &v
	return s
}

type SubmitAudioTo3DAvatarVideoTaskResponseBody struct {
	Code      *string                                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *SubmitAudioTo3DAvatarVideoTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                           `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SubmitAudioTo3DAvatarVideoTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitAudioTo3DAvatarVideoTaskResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitAudioTo3DAvatarVideoTaskResponseBody) SetCode(v string) *SubmitAudioTo3DAvatarVideoTaskResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitAudioTo3DAvatarVideoTaskResponseBody) SetData(v *SubmitAudioTo3DAvatarVideoTaskResponseBodyData) *SubmitAudioTo3DAvatarVideoTaskResponseBody {
	s.Data = v
	return s
}

func (s *SubmitAudioTo3DAvatarVideoTaskResponseBody) SetMessage(v string) *SubmitAudioTo3DAvatarVideoTaskResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitAudioTo3DAvatarVideoTaskResponseBody) SetRequestId(v string) *SubmitAudioTo3DAvatarVideoTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitAudioTo3DAvatarVideoTaskResponseBody) SetSuccess(v bool) *SubmitAudioTo3DAvatarVideoTaskResponseBody {
	s.Success = &v
	return s
}

type SubmitAudioTo3DAvatarVideoTaskResponseBodyData struct {
	TaskUuid *string `json:"TaskUuid,omitempty" xml:"TaskUuid,omitempty"`
}

func (s SubmitAudioTo3DAvatarVideoTaskResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SubmitAudioTo3DAvatarVideoTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitAudioTo3DAvatarVideoTaskResponseBodyData) SetTaskUuid(v string) *SubmitAudioTo3DAvatarVideoTaskResponseBodyData {
	s.TaskUuid = &v
	return s
}

type SubmitAudioTo3DAvatarVideoTaskResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SubmitAudioTo3DAvatarVideoTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubmitAudioTo3DAvatarVideoTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitAudioTo3DAvatarVideoTaskResponse) GoString() string {
	return s.String()
}

func (s *SubmitAudioTo3DAvatarVideoTaskResponse) SetHeaders(v map[string]*string) *SubmitAudioTo3DAvatarVideoTaskResponse {
	s.Headers = v
	return s
}

func (s *SubmitAudioTo3DAvatarVideoTaskResponse) SetStatusCode(v int32) *SubmitAudioTo3DAvatarVideoTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitAudioTo3DAvatarVideoTaskResponse) SetBody(v *SubmitAudioTo3DAvatarVideoTaskResponseBody) *SubmitAudioTo3DAvatarVideoTaskResponse {
	s.Body = v
	return s
}

type SubmitAvatarVideoTaskRequest struct {
	App            *SubmitAvatarVideoTaskRequestApp `json:"App,omitempty" xml:"App,omitempty" type:"Struct"`
	Callback       *bool                            `json:"Callback,omitempty" xml:"Callback,omitempty"`
	CallbackParams *string                          `json:"CallbackParams,omitempty" xml:"CallbackParams,omitempty"`
	ExtParams      *string                          `json:"ExtParams,omitempty" xml:"ExtParams,omitempty"`
	TenantId       *int64                           `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	Title          *string                          `json:"Title,omitempty" xml:"Title,omitempty"`
	VideoParams    *string                          `json:"VideoParams,omitempty" xml:"VideoParams,omitempty"`
}

func (s SubmitAvatarVideoTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitAvatarVideoTaskRequest) GoString() string {
	return s.String()
}

func (s *SubmitAvatarVideoTaskRequest) SetApp(v *SubmitAvatarVideoTaskRequestApp) *SubmitAvatarVideoTaskRequest {
	s.App = v
	return s
}

func (s *SubmitAvatarVideoTaskRequest) SetCallback(v bool) *SubmitAvatarVideoTaskRequest {
	s.Callback = &v
	return s
}

func (s *SubmitAvatarVideoTaskRequest) SetCallbackParams(v string) *SubmitAvatarVideoTaskRequest {
	s.CallbackParams = &v
	return s
}

func (s *SubmitAvatarVideoTaskRequest) SetExtParams(v string) *SubmitAvatarVideoTaskRequest {
	s.ExtParams = &v
	return s
}

func (s *SubmitAvatarVideoTaskRequest) SetTenantId(v int64) *SubmitAvatarVideoTaskRequest {
	s.TenantId = &v
	return s
}

func (s *SubmitAvatarVideoTaskRequest) SetTitle(v string) *SubmitAvatarVideoTaskRequest {
	s.Title = &v
	return s
}

func (s *SubmitAvatarVideoTaskRequest) SetVideoParams(v string) *SubmitAvatarVideoTaskRequest {
	s.VideoParams = &v
	return s
}

type SubmitAvatarVideoTaskRequestApp struct {
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s SubmitAvatarVideoTaskRequestApp) String() string {
	return tea.Prettify(s)
}

func (s SubmitAvatarVideoTaskRequestApp) GoString() string {
	return s.String()
}

func (s *SubmitAvatarVideoTaskRequestApp) SetAppId(v string) *SubmitAvatarVideoTaskRequestApp {
	s.AppId = &v
	return s
}

type SubmitAvatarVideoTaskShrinkRequest struct {
	AppShrink      *string `json:"App,omitempty" xml:"App,omitempty"`
	Callback       *bool   `json:"Callback,omitempty" xml:"Callback,omitempty"`
	CallbackParams *string `json:"CallbackParams,omitempty" xml:"CallbackParams,omitempty"`
	ExtParams      *string `json:"ExtParams,omitempty" xml:"ExtParams,omitempty"`
	TenantId       *int64  `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	Title          *string `json:"Title,omitempty" xml:"Title,omitempty"`
	VideoParams    *string `json:"VideoParams,omitempty" xml:"VideoParams,omitempty"`
}

func (s SubmitAvatarVideoTaskShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitAvatarVideoTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *SubmitAvatarVideoTaskShrinkRequest) SetAppShrink(v string) *SubmitAvatarVideoTaskShrinkRequest {
	s.AppShrink = &v
	return s
}

func (s *SubmitAvatarVideoTaskShrinkRequest) SetCallback(v bool) *SubmitAvatarVideoTaskShrinkRequest {
	s.Callback = &v
	return s
}

func (s *SubmitAvatarVideoTaskShrinkRequest) SetCallbackParams(v string) *SubmitAvatarVideoTaskShrinkRequest {
	s.CallbackParams = &v
	return s
}

func (s *SubmitAvatarVideoTaskShrinkRequest) SetExtParams(v string) *SubmitAvatarVideoTaskShrinkRequest {
	s.ExtParams = &v
	return s
}

func (s *SubmitAvatarVideoTaskShrinkRequest) SetTenantId(v int64) *SubmitAvatarVideoTaskShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *SubmitAvatarVideoTaskShrinkRequest) SetTitle(v string) *SubmitAvatarVideoTaskShrinkRequest {
	s.Title = &v
	return s
}

func (s *SubmitAvatarVideoTaskShrinkRequest) SetVideoParams(v string) *SubmitAvatarVideoTaskShrinkRequest {
	s.VideoParams = &v
	return s
}

type SubmitAvatarVideoTaskResponseBody struct {
	Code    *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    *SubmitAvatarVideoTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SubmitAvatarVideoTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitAvatarVideoTaskResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitAvatarVideoTaskResponseBody) SetCode(v string) *SubmitAvatarVideoTaskResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitAvatarVideoTaskResponseBody) SetData(v *SubmitAvatarVideoTaskResponseBodyData) *SubmitAvatarVideoTaskResponseBody {
	s.Data = v
	return s
}

func (s *SubmitAvatarVideoTaskResponseBody) SetMessage(v string) *SubmitAvatarVideoTaskResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitAvatarVideoTaskResponseBody) SetRequestId(v string) *SubmitAvatarVideoTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitAvatarVideoTaskResponseBody) SetSuccess(v bool) *SubmitAvatarVideoTaskResponseBody {
	s.Success = &v
	return s
}

type SubmitAvatarVideoTaskResponseBodyData struct {
	TaskUuid *string `json:"TaskUuid,omitempty" xml:"TaskUuid,omitempty"`
}

func (s SubmitAvatarVideoTaskResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SubmitAvatarVideoTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitAvatarVideoTaskResponseBodyData) SetTaskUuid(v string) *SubmitAvatarVideoTaskResponseBodyData {
	s.TaskUuid = &v
	return s
}

type SubmitAvatarVideoTaskResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SubmitAvatarVideoTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubmitAvatarVideoTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitAvatarVideoTaskResponse) GoString() string {
	return s.String()
}

func (s *SubmitAvatarVideoTaskResponse) SetHeaders(v map[string]*string) *SubmitAvatarVideoTaskResponse {
	s.Headers = v
	return s
}

func (s *SubmitAvatarVideoTaskResponse) SetStatusCode(v int32) *SubmitAvatarVideoTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitAvatarVideoTaskResponse) SetBody(v *SubmitAvatarVideoTaskResponseBody) *SubmitAvatarVideoTaskResponse {
	s.Body = v
	return s
}

type SubmitTextTo2DAvatarVideoTaskRequest struct {
	App            *SubmitTextTo2DAvatarVideoTaskRequestApp        `json:"App,omitempty" xml:"App,omitempty" type:"Struct"`
	AudioInfo      *SubmitTextTo2DAvatarVideoTaskRequestAudioInfo  `json:"AudioInfo,omitempty" xml:"AudioInfo,omitempty" type:"Struct"`
	AvatarInfo     *SubmitTextTo2DAvatarVideoTaskRequestAvatarInfo `json:"AvatarInfo,omitempty" xml:"AvatarInfo,omitempty" type:"Struct"`
	Callback       *bool                                           `json:"Callback,omitempty" xml:"Callback,omitempty"`
	CallbackParams *string                                         `json:"CallbackParams,omitempty" xml:"CallbackParams,omitempty"`
	ExtParams      *string                                         `json:"ExtParams,omitempty" xml:"ExtParams,omitempty"`
	TenantId       *int64                                          `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	Text           *string                                         `json:"Text,omitempty" xml:"Text,omitempty"`
	Title          *string                                         `json:"Title,omitempty" xml:"Title,omitempty"`
	VideoInfo      *SubmitTextTo2DAvatarVideoTaskRequestVideoInfo  `json:"VideoInfo,omitempty" xml:"VideoInfo,omitempty" type:"Struct"`
}

func (s SubmitTextTo2DAvatarVideoTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitTextTo2DAvatarVideoTaskRequest) GoString() string {
	return s.String()
}

func (s *SubmitTextTo2DAvatarVideoTaskRequest) SetApp(v *SubmitTextTo2DAvatarVideoTaskRequestApp) *SubmitTextTo2DAvatarVideoTaskRequest {
	s.App = v
	return s
}

func (s *SubmitTextTo2DAvatarVideoTaskRequest) SetAudioInfo(v *SubmitTextTo2DAvatarVideoTaskRequestAudioInfo) *SubmitTextTo2DAvatarVideoTaskRequest {
	s.AudioInfo = v
	return s
}

func (s *SubmitTextTo2DAvatarVideoTaskRequest) SetAvatarInfo(v *SubmitTextTo2DAvatarVideoTaskRequestAvatarInfo) *SubmitTextTo2DAvatarVideoTaskRequest {
	s.AvatarInfo = v
	return s
}

func (s *SubmitTextTo2DAvatarVideoTaskRequest) SetCallback(v bool) *SubmitTextTo2DAvatarVideoTaskRequest {
	s.Callback = &v
	return s
}

func (s *SubmitTextTo2DAvatarVideoTaskRequest) SetCallbackParams(v string) *SubmitTextTo2DAvatarVideoTaskRequest {
	s.CallbackParams = &v
	return s
}

func (s *SubmitTextTo2DAvatarVideoTaskRequest) SetExtParams(v string) *SubmitTextTo2DAvatarVideoTaskRequest {
	s.ExtParams = &v
	return s
}

func (s *SubmitTextTo2DAvatarVideoTaskRequest) SetTenantId(v int64) *SubmitTextTo2DAvatarVideoTaskRequest {
	s.TenantId = &v
	return s
}

func (s *SubmitTextTo2DAvatarVideoTaskRequest) SetText(v string) *SubmitTextTo2DAvatarVideoTaskRequest {
	s.Text = &v
	return s
}

func (s *SubmitTextTo2DAvatarVideoTaskRequest) SetTitle(v string) *SubmitTextTo2DAvatarVideoTaskRequest {
	s.Title = &v
	return s
}

func (s *SubmitTextTo2DAvatarVideoTaskRequest) SetVideoInfo(v *SubmitTextTo2DAvatarVideoTaskRequestVideoInfo) *SubmitTextTo2DAvatarVideoTaskRequest {
	s.VideoInfo = v
	return s
}

type SubmitTextTo2DAvatarVideoTaskRequestApp struct {
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s SubmitTextTo2DAvatarVideoTaskRequestApp) String() string {
	return tea.Prettify(s)
}

func (s SubmitTextTo2DAvatarVideoTaskRequestApp) GoString() string {
	return s.String()
}

func (s *SubmitTextTo2DAvatarVideoTaskRequestApp) SetAppId(v string) *SubmitTextTo2DAvatarVideoTaskRequestApp {
	s.AppId = &v
	return s
}

type SubmitTextTo2DAvatarVideoTaskRequestAudioInfo struct {
	PitchRate  *int32  `json:"PitchRate,omitempty" xml:"PitchRate,omitempty"`
	SpeechRate *int32  `json:"SpeechRate,omitempty" xml:"SpeechRate,omitempty"`
	Voice      *string `json:"Voice,omitempty" xml:"Voice,omitempty"`
	Volume     *int32  `json:"Volume,omitempty" xml:"Volume,omitempty"`
}

func (s SubmitTextTo2DAvatarVideoTaskRequestAudioInfo) String() string {
	return tea.Prettify(s)
}

func (s SubmitTextTo2DAvatarVideoTaskRequestAudioInfo) GoString() string {
	return s.String()
}

func (s *SubmitTextTo2DAvatarVideoTaskRequestAudioInfo) SetPitchRate(v int32) *SubmitTextTo2DAvatarVideoTaskRequestAudioInfo {
	s.PitchRate = &v
	return s
}

func (s *SubmitTextTo2DAvatarVideoTaskRequestAudioInfo) SetSpeechRate(v int32) *SubmitTextTo2DAvatarVideoTaskRequestAudioInfo {
	s.SpeechRate = &v
	return s
}

func (s *SubmitTextTo2DAvatarVideoTaskRequestAudioInfo) SetVoice(v string) *SubmitTextTo2DAvatarVideoTaskRequestAudioInfo {
	s.Voice = &v
	return s
}

func (s *SubmitTextTo2DAvatarVideoTaskRequestAudioInfo) SetVolume(v int32) *SubmitTextTo2DAvatarVideoTaskRequestAudioInfo {
	s.Volume = &v
	return s
}

type SubmitTextTo2DAvatarVideoTaskRequestAvatarInfo struct {
	Code   *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Height *int32  `json:"Height,omitempty" xml:"Height,omitempty"`
	Width  *int32  `json:"Width,omitempty" xml:"Width,omitempty"`
	X      *int32  `json:"X,omitempty" xml:"X,omitempty"`
	Y      *int32  `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s SubmitTextTo2DAvatarVideoTaskRequestAvatarInfo) String() string {
	return tea.Prettify(s)
}

func (s SubmitTextTo2DAvatarVideoTaskRequestAvatarInfo) GoString() string {
	return s.String()
}

func (s *SubmitTextTo2DAvatarVideoTaskRequestAvatarInfo) SetCode(v string) *SubmitTextTo2DAvatarVideoTaskRequestAvatarInfo {
	s.Code = &v
	return s
}

func (s *SubmitTextTo2DAvatarVideoTaskRequestAvatarInfo) SetHeight(v int32) *SubmitTextTo2DAvatarVideoTaskRequestAvatarInfo {
	s.Height = &v
	return s
}

func (s *SubmitTextTo2DAvatarVideoTaskRequestAvatarInfo) SetWidth(v int32) *SubmitTextTo2DAvatarVideoTaskRequestAvatarInfo {
	s.Width = &v
	return s
}

func (s *SubmitTextTo2DAvatarVideoTaskRequestAvatarInfo) SetX(v int32) *SubmitTextTo2DAvatarVideoTaskRequestAvatarInfo {
	s.X = &v
	return s
}

func (s *SubmitTextTo2DAvatarVideoTaskRequestAvatarInfo) SetY(v int32) *SubmitTextTo2DAvatarVideoTaskRequestAvatarInfo {
	s.Y = &v
	return s
}

type SubmitTextTo2DAvatarVideoTaskRequestVideoInfo struct {
	AlphaFormat        *int32  `json:"AlphaFormat,omitempty" xml:"AlphaFormat,omitempty"`
	BackgroundImageUrl *string `json:"BackgroundImageUrl,omitempty" xml:"BackgroundImageUrl,omitempty"`
	IsAlpha            *bool   `json:"IsAlpha,omitempty" xml:"IsAlpha,omitempty"`
	// Deprecated
	IsSubtitles      *bool  `json:"IsSubtitles,omitempty" xml:"IsSubtitles,omitempty"`
	Resolution       *int32 `json:"Resolution,omitempty" xml:"Resolution,omitempty"`
	SubtitleEmbedded *bool  `json:"SubtitleEmbedded,omitempty" xml:"SubtitleEmbedded,omitempty"`
}

func (s SubmitTextTo2DAvatarVideoTaskRequestVideoInfo) String() string {
	return tea.Prettify(s)
}

func (s SubmitTextTo2DAvatarVideoTaskRequestVideoInfo) GoString() string {
	return s.String()
}

func (s *SubmitTextTo2DAvatarVideoTaskRequestVideoInfo) SetAlphaFormat(v int32) *SubmitTextTo2DAvatarVideoTaskRequestVideoInfo {
	s.AlphaFormat = &v
	return s
}

func (s *SubmitTextTo2DAvatarVideoTaskRequestVideoInfo) SetBackgroundImageUrl(v string) *SubmitTextTo2DAvatarVideoTaskRequestVideoInfo {
	s.BackgroundImageUrl = &v
	return s
}

func (s *SubmitTextTo2DAvatarVideoTaskRequestVideoInfo) SetIsAlpha(v bool) *SubmitTextTo2DAvatarVideoTaskRequestVideoInfo {
	s.IsAlpha = &v
	return s
}

func (s *SubmitTextTo2DAvatarVideoTaskRequestVideoInfo) SetIsSubtitles(v bool) *SubmitTextTo2DAvatarVideoTaskRequestVideoInfo {
	s.IsSubtitles = &v
	return s
}

func (s *SubmitTextTo2DAvatarVideoTaskRequestVideoInfo) SetResolution(v int32) *SubmitTextTo2DAvatarVideoTaskRequestVideoInfo {
	s.Resolution = &v
	return s
}

func (s *SubmitTextTo2DAvatarVideoTaskRequestVideoInfo) SetSubtitleEmbedded(v bool) *SubmitTextTo2DAvatarVideoTaskRequestVideoInfo {
	s.SubtitleEmbedded = &v
	return s
}

type SubmitTextTo2DAvatarVideoTaskShrinkRequest struct {
	AppShrink        *string `json:"App,omitempty" xml:"App,omitempty"`
	AudioInfoShrink  *string `json:"AudioInfo,omitempty" xml:"AudioInfo,omitempty"`
	AvatarInfoShrink *string `json:"AvatarInfo,omitempty" xml:"AvatarInfo,omitempty"`
	Callback         *bool   `json:"Callback,omitempty" xml:"Callback,omitempty"`
	CallbackParams   *string `json:"CallbackParams,omitempty" xml:"CallbackParams,omitempty"`
	ExtParams        *string `json:"ExtParams,omitempty" xml:"ExtParams,omitempty"`
	TenantId         *int64  `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	Text             *string `json:"Text,omitempty" xml:"Text,omitempty"`
	Title            *string `json:"Title,omitempty" xml:"Title,omitempty"`
	VideoInfoShrink  *string `json:"VideoInfo,omitempty" xml:"VideoInfo,omitempty"`
}

func (s SubmitTextTo2DAvatarVideoTaskShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitTextTo2DAvatarVideoTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *SubmitTextTo2DAvatarVideoTaskShrinkRequest) SetAppShrink(v string) *SubmitTextTo2DAvatarVideoTaskShrinkRequest {
	s.AppShrink = &v
	return s
}

func (s *SubmitTextTo2DAvatarVideoTaskShrinkRequest) SetAudioInfoShrink(v string) *SubmitTextTo2DAvatarVideoTaskShrinkRequest {
	s.AudioInfoShrink = &v
	return s
}

func (s *SubmitTextTo2DAvatarVideoTaskShrinkRequest) SetAvatarInfoShrink(v string) *SubmitTextTo2DAvatarVideoTaskShrinkRequest {
	s.AvatarInfoShrink = &v
	return s
}

func (s *SubmitTextTo2DAvatarVideoTaskShrinkRequest) SetCallback(v bool) *SubmitTextTo2DAvatarVideoTaskShrinkRequest {
	s.Callback = &v
	return s
}

func (s *SubmitTextTo2DAvatarVideoTaskShrinkRequest) SetCallbackParams(v string) *SubmitTextTo2DAvatarVideoTaskShrinkRequest {
	s.CallbackParams = &v
	return s
}

func (s *SubmitTextTo2DAvatarVideoTaskShrinkRequest) SetExtParams(v string) *SubmitTextTo2DAvatarVideoTaskShrinkRequest {
	s.ExtParams = &v
	return s
}

func (s *SubmitTextTo2DAvatarVideoTaskShrinkRequest) SetTenantId(v int64) *SubmitTextTo2DAvatarVideoTaskShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *SubmitTextTo2DAvatarVideoTaskShrinkRequest) SetText(v string) *SubmitTextTo2DAvatarVideoTaskShrinkRequest {
	s.Text = &v
	return s
}

func (s *SubmitTextTo2DAvatarVideoTaskShrinkRequest) SetTitle(v string) *SubmitTextTo2DAvatarVideoTaskShrinkRequest {
	s.Title = &v
	return s
}

func (s *SubmitTextTo2DAvatarVideoTaskShrinkRequest) SetVideoInfoShrink(v string) *SubmitTextTo2DAvatarVideoTaskShrinkRequest {
	s.VideoInfoShrink = &v
	return s
}

type SubmitTextTo2DAvatarVideoTaskResponseBody struct {
	Code      *string                                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *SubmitTextTo2DAvatarVideoTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                        `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                          `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SubmitTextTo2DAvatarVideoTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitTextTo2DAvatarVideoTaskResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitTextTo2DAvatarVideoTaskResponseBody) SetCode(v string) *SubmitTextTo2DAvatarVideoTaskResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitTextTo2DAvatarVideoTaskResponseBody) SetData(v *SubmitTextTo2DAvatarVideoTaskResponseBodyData) *SubmitTextTo2DAvatarVideoTaskResponseBody {
	s.Data = v
	return s
}

func (s *SubmitTextTo2DAvatarVideoTaskResponseBody) SetMessage(v string) *SubmitTextTo2DAvatarVideoTaskResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitTextTo2DAvatarVideoTaskResponseBody) SetRequestId(v string) *SubmitTextTo2DAvatarVideoTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitTextTo2DAvatarVideoTaskResponseBody) SetSuccess(v bool) *SubmitTextTo2DAvatarVideoTaskResponseBody {
	s.Success = &v
	return s
}

type SubmitTextTo2DAvatarVideoTaskResponseBodyData struct {
	TaskUuid *string `json:"TaskUuid,omitempty" xml:"TaskUuid,omitempty"`
}

func (s SubmitTextTo2DAvatarVideoTaskResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SubmitTextTo2DAvatarVideoTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitTextTo2DAvatarVideoTaskResponseBodyData) SetTaskUuid(v string) *SubmitTextTo2DAvatarVideoTaskResponseBodyData {
	s.TaskUuid = &v
	return s
}

type SubmitTextTo2DAvatarVideoTaskResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SubmitTextTo2DAvatarVideoTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubmitTextTo2DAvatarVideoTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitTextTo2DAvatarVideoTaskResponse) GoString() string {
	return s.String()
}

func (s *SubmitTextTo2DAvatarVideoTaskResponse) SetHeaders(v map[string]*string) *SubmitTextTo2DAvatarVideoTaskResponse {
	s.Headers = v
	return s
}

func (s *SubmitTextTo2DAvatarVideoTaskResponse) SetStatusCode(v int32) *SubmitTextTo2DAvatarVideoTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitTextTo2DAvatarVideoTaskResponse) SetBody(v *SubmitTextTo2DAvatarVideoTaskResponseBody) *SubmitTextTo2DAvatarVideoTaskResponse {
	s.Body = v
	return s
}

type SubmitTextTo3DAvatarVideoTaskRequest struct {
	App            *SubmitTextTo3DAvatarVideoTaskRequestApp        `json:"App,omitempty" xml:"App,omitempty" type:"Struct"`
	AudioInfo      *SubmitTextTo3DAvatarVideoTaskRequestAudioInfo  `json:"AudioInfo,omitempty" xml:"AudioInfo,omitempty" type:"Struct"`
	AvatarInfo     *SubmitTextTo3DAvatarVideoTaskRequestAvatarInfo `json:"AvatarInfo,omitempty" xml:"AvatarInfo,omitempty" type:"Struct"`
	Callback       *bool                                           `json:"Callback,omitempty" xml:"Callback,omitempty"`
	CallbackParams *string                                         `json:"CallbackParams,omitempty" xml:"CallbackParams,omitempty"`
	ExtParams      *string                                         `json:"ExtParams,omitempty" xml:"ExtParams,omitempty"`
	TenantId       *int64                                          `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	Text           *string                                         `json:"Text,omitempty" xml:"Text,omitempty"`
	Title          *string                                         `json:"Title,omitempty" xml:"Title,omitempty"`
	VideoInfo      *SubmitTextTo3DAvatarVideoTaskRequestVideoInfo  `json:"VideoInfo,omitempty" xml:"VideoInfo,omitempty" type:"Struct"`
}

func (s SubmitTextTo3DAvatarVideoTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitTextTo3DAvatarVideoTaskRequest) GoString() string {
	return s.String()
}

func (s *SubmitTextTo3DAvatarVideoTaskRequest) SetApp(v *SubmitTextTo3DAvatarVideoTaskRequestApp) *SubmitTextTo3DAvatarVideoTaskRequest {
	s.App = v
	return s
}

func (s *SubmitTextTo3DAvatarVideoTaskRequest) SetAudioInfo(v *SubmitTextTo3DAvatarVideoTaskRequestAudioInfo) *SubmitTextTo3DAvatarVideoTaskRequest {
	s.AudioInfo = v
	return s
}

func (s *SubmitTextTo3DAvatarVideoTaskRequest) SetAvatarInfo(v *SubmitTextTo3DAvatarVideoTaskRequestAvatarInfo) *SubmitTextTo3DAvatarVideoTaskRequest {
	s.AvatarInfo = v
	return s
}

func (s *SubmitTextTo3DAvatarVideoTaskRequest) SetCallback(v bool) *SubmitTextTo3DAvatarVideoTaskRequest {
	s.Callback = &v
	return s
}

func (s *SubmitTextTo3DAvatarVideoTaskRequest) SetCallbackParams(v string) *SubmitTextTo3DAvatarVideoTaskRequest {
	s.CallbackParams = &v
	return s
}

func (s *SubmitTextTo3DAvatarVideoTaskRequest) SetExtParams(v string) *SubmitTextTo3DAvatarVideoTaskRequest {
	s.ExtParams = &v
	return s
}

func (s *SubmitTextTo3DAvatarVideoTaskRequest) SetTenantId(v int64) *SubmitTextTo3DAvatarVideoTaskRequest {
	s.TenantId = &v
	return s
}

func (s *SubmitTextTo3DAvatarVideoTaskRequest) SetText(v string) *SubmitTextTo3DAvatarVideoTaskRequest {
	s.Text = &v
	return s
}

func (s *SubmitTextTo3DAvatarVideoTaskRequest) SetTitle(v string) *SubmitTextTo3DAvatarVideoTaskRequest {
	s.Title = &v
	return s
}

func (s *SubmitTextTo3DAvatarVideoTaskRequest) SetVideoInfo(v *SubmitTextTo3DAvatarVideoTaskRequestVideoInfo) *SubmitTextTo3DAvatarVideoTaskRequest {
	s.VideoInfo = v
	return s
}

type SubmitTextTo3DAvatarVideoTaskRequestApp struct {
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s SubmitTextTo3DAvatarVideoTaskRequestApp) String() string {
	return tea.Prettify(s)
}

func (s SubmitTextTo3DAvatarVideoTaskRequestApp) GoString() string {
	return s.String()
}

func (s *SubmitTextTo3DAvatarVideoTaskRequestApp) SetAppId(v string) *SubmitTextTo3DAvatarVideoTaskRequestApp {
	s.AppId = &v
	return s
}

type SubmitTextTo3DAvatarVideoTaskRequestAudioInfo struct {
	PitchRate  *int32  `json:"PitchRate,omitempty" xml:"PitchRate,omitempty"`
	SpeechRate *int32  `json:"SpeechRate,omitempty" xml:"SpeechRate,omitempty"`
	Voice      *string `json:"Voice,omitempty" xml:"Voice,omitempty"`
	Volume     *int32  `json:"Volume,omitempty" xml:"Volume,omitempty"`
}

func (s SubmitTextTo3DAvatarVideoTaskRequestAudioInfo) String() string {
	return tea.Prettify(s)
}

func (s SubmitTextTo3DAvatarVideoTaskRequestAudioInfo) GoString() string {
	return s.String()
}

func (s *SubmitTextTo3DAvatarVideoTaskRequestAudioInfo) SetPitchRate(v int32) *SubmitTextTo3DAvatarVideoTaskRequestAudioInfo {
	s.PitchRate = &v
	return s
}

func (s *SubmitTextTo3DAvatarVideoTaskRequestAudioInfo) SetSpeechRate(v int32) *SubmitTextTo3DAvatarVideoTaskRequestAudioInfo {
	s.SpeechRate = &v
	return s
}

func (s *SubmitTextTo3DAvatarVideoTaskRequestAudioInfo) SetVoice(v string) *SubmitTextTo3DAvatarVideoTaskRequestAudioInfo {
	s.Voice = &v
	return s
}

func (s *SubmitTextTo3DAvatarVideoTaskRequestAudioInfo) SetVolume(v int32) *SubmitTextTo3DAvatarVideoTaskRequestAudioInfo {
	s.Volume = &v
	return s
}

type SubmitTextTo3DAvatarVideoTaskRequestAvatarInfo struct {
	Angle        *int32  `json:"Angle,omitempty" xml:"Angle,omitempty"`
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	IndustryCode *string `json:"IndustryCode,omitempty" xml:"IndustryCode,omitempty"`
	Locate       *int32  `json:"Locate,omitempty" xml:"Locate,omitempty"`
}

func (s SubmitTextTo3DAvatarVideoTaskRequestAvatarInfo) String() string {
	return tea.Prettify(s)
}

func (s SubmitTextTo3DAvatarVideoTaskRequestAvatarInfo) GoString() string {
	return s.String()
}

func (s *SubmitTextTo3DAvatarVideoTaskRequestAvatarInfo) SetAngle(v int32) *SubmitTextTo3DAvatarVideoTaskRequestAvatarInfo {
	s.Angle = &v
	return s
}

func (s *SubmitTextTo3DAvatarVideoTaskRequestAvatarInfo) SetCode(v string) *SubmitTextTo3DAvatarVideoTaskRequestAvatarInfo {
	s.Code = &v
	return s
}

func (s *SubmitTextTo3DAvatarVideoTaskRequestAvatarInfo) SetIndustryCode(v string) *SubmitTextTo3DAvatarVideoTaskRequestAvatarInfo {
	s.IndustryCode = &v
	return s
}

func (s *SubmitTextTo3DAvatarVideoTaskRequestAvatarInfo) SetLocate(v int32) *SubmitTextTo3DAvatarVideoTaskRequestAvatarInfo {
	s.Locate = &v
	return s
}

type SubmitTextTo3DAvatarVideoTaskRequestVideoInfo struct {
	AlphaFormat        *int32  `json:"AlphaFormat,omitempty" xml:"AlphaFormat,omitempty"`
	BackgroundImageUrl *string `json:"BackgroundImageUrl,omitempty" xml:"BackgroundImageUrl,omitempty"`
	IsAlpha            *bool   `json:"IsAlpha,omitempty" xml:"IsAlpha,omitempty"`
	// Deprecated
	IsSubtitles      *bool  `json:"IsSubtitles,omitempty" xml:"IsSubtitles,omitempty"`
	Resolution       *int32 `json:"Resolution,omitempty" xml:"Resolution,omitempty"`
	SubtitleEmbedded *bool  `json:"SubtitleEmbedded,omitempty" xml:"SubtitleEmbedded,omitempty"`
}

func (s SubmitTextTo3DAvatarVideoTaskRequestVideoInfo) String() string {
	return tea.Prettify(s)
}

func (s SubmitTextTo3DAvatarVideoTaskRequestVideoInfo) GoString() string {
	return s.String()
}

func (s *SubmitTextTo3DAvatarVideoTaskRequestVideoInfo) SetAlphaFormat(v int32) *SubmitTextTo3DAvatarVideoTaskRequestVideoInfo {
	s.AlphaFormat = &v
	return s
}

func (s *SubmitTextTo3DAvatarVideoTaskRequestVideoInfo) SetBackgroundImageUrl(v string) *SubmitTextTo3DAvatarVideoTaskRequestVideoInfo {
	s.BackgroundImageUrl = &v
	return s
}

func (s *SubmitTextTo3DAvatarVideoTaskRequestVideoInfo) SetIsAlpha(v bool) *SubmitTextTo3DAvatarVideoTaskRequestVideoInfo {
	s.IsAlpha = &v
	return s
}

func (s *SubmitTextTo3DAvatarVideoTaskRequestVideoInfo) SetIsSubtitles(v bool) *SubmitTextTo3DAvatarVideoTaskRequestVideoInfo {
	s.IsSubtitles = &v
	return s
}

func (s *SubmitTextTo3DAvatarVideoTaskRequestVideoInfo) SetResolution(v int32) *SubmitTextTo3DAvatarVideoTaskRequestVideoInfo {
	s.Resolution = &v
	return s
}

func (s *SubmitTextTo3DAvatarVideoTaskRequestVideoInfo) SetSubtitleEmbedded(v bool) *SubmitTextTo3DAvatarVideoTaskRequestVideoInfo {
	s.SubtitleEmbedded = &v
	return s
}

type SubmitTextTo3DAvatarVideoTaskShrinkRequest struct {
	AppShrink        *string `json:"App,omitempty" xml:"App,omitempty"`
	AudioInfoShrink  *string `json:"AudioInfo,omitempty" xml:"AudioInfo,omitempty"`
	AvatarInfoShrink *string `json:"AvatarInfo,omitempty" xml:"AvatarInfo,omitempty"`
	Callback         *bool   `json:"Callback,omitempty" xml:"Callback,omitempty"`
	CallbackParams   *string `json:"CallbackParams,omitempty" xml:"CallbackParams,omitempty"`
	ExtParams        *string `json:"ExtParams,omitempty" xml:"ExtParams,omitempty"`
	TenantId         *int64  `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	Text             *string `json:"Text,omitempty" xml:"Text,omitempty"`
	Title            *string `json:"Title,omitempty" xml:"Title,omitempty"`
	VideoInfoShrink  *string `json:"VideoInfo,omitempty" xml:"VideoInfo,omitempty"`
}

func (s SubmitTextTo3DAvatarVideoTaskShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitTextTo3DAvatarVideoTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *SubmitTextTo3DAvatarVideoTaskShrinkRequest) SetAppShrink(v string) *SubmitTextTo3DAvatarVideoTaskShrinkRequest {
	s.AppShrink = &v
	return s
}

func (s *SubmitTextTo3DAvatarVideoTaskShrinkRequest) SetAudioInfoShrink(v string) *SubmitTextTo3DAvatarVideoTaskShrinkRequest {
	s.AudioInfoShrink = &v
	return s
}

func (s *SubmitTextTo3DAvatarVideoTaskShrinkRequest) SetAvatarInfoShrink(v string) *SubmitTextTo3DAvatarVideoTaskShrinkRequest {
	s.AvatarInfoShrink = &v
	return s
}

func (s *SubmitTextTo3DAvatarVideoTaskShrinkRequest) SetCallback(v bool) *SubmitTextTo3DAvatarVideoTaskShrinkRequest {
	s.Callback = &v
	return s
}

func (s *SubmitTextTo3DAvatarVideoTaskShrinkRequest) SetCallbackParams(v string) *SubmitTextTo3DAvatarVideoTaskShrinkRequest {
	s.CallbackParams = &v
	return s
}

func (s *SubmitTextTo3DAvatarVideoTaskShrinkRequest) SetExtParams(v string) *SubmitTextTo3DAvatarVideoTaskShrinkRequest {
	s.ExtParams = &v
	return s
}

func (s *SubmitTextTo3DAvatarVideoTaskShrinkRequest) SetTenantId(v int64) *SubmitTextTo3DAvatarVideoTaskShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *SubmitTextTo3DAvatarVideoTaskShrinkRequest) SetText(v string) *SubmitTextTo3DAvatarVideoTaskShrinkRequest {
	s.Text = &v
	return s
}

func (s *SubmitTextTo3DAvatarVideoTaskShrinkRequest) SetTitle(v string) *SubmitTextTo3DAvatarVideoTaskShrinkRequest {
	s.Title = &v
	return s
}

func (s *SubmitTextTo3DAvatarVideoTaskShrinkRequest) SetVideoInfoShrink(v string) *SubmitTextTo3DAvatarVideoTaskShrinkRequest {
	s.VideoInfoShrink = &v
	return s
}

type SubmitTextTo3DAvatarVideoTaskResponseBody struct {
	Code      *string                                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *SubmitTextTo3DAvatarVideoTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                        `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                          `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SubmitTextTo3DAvatarVideoTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitTextTo3DAvatarVideoTaskResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitTextTo3DAvatarVideoTaskResponseBody) SetCode(v string) *SubmitTextTo3DAvatarVideoTaskResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitTextTo3DAvatarVideoTaskResponseBody) SetData(v *SubmitTextTo3DAvatarVideoTaskResponseBodyData) *SubmitTextTo3DAvatarVideoTaskResponseBody {
	s.Data = v
	return s
}

func (s *SubmitTextTo3DAvatarVideoTaskResponseBody) SetMessage(v string) *SubmitTextTo3DAvatarVideoTaskResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitTextTo3DAvatarVideoTaskResponseBody) SetRequestId(v string) *SubmitTextTo3DAvatarVideoTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitTextTo3DAvatarVideoTaskResponseBody) SetSuccess(v bool) *SubmitTextTo3DAvatarVideoTaskResponseBody {
	s.Success = &v
	return s
}

type SubmitTextTo3DAvatarVideoTaskResponseBodyData struct {
	TaskUuid *string `json:"TaskUuid,omitempty" xml:"TaskUuid,omitempty"`
}

func (s SubmitTextTo3DAvatarVideoTaskResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SubmitTextTo3DAvatarVideoTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitTextTo3DAvatarVideoTaskResponseBodyData) SetTaskUuid(v string) *SubmitTextTo3DAvatarVideoTaskResponseBodyData {
	s.TaskUuid = &v
	return s
}

type SubmitTextTo3DAvatarVideoTaskResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SubmitTextTo3DAvatarVideoTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubmitTextTo3DAvatarVideoTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitTextTo3DAvatarVideoTaskResponse) GoString() string {
	return s.String()
}

func (s *SubmitTextTo3DAvatarVideoTaskResponse) SetHeaders(v map[string]*string) *SubmitTextTo3DAvatarVideoTaskResponse {
	s.Headers = v
	return s
}

func (s *SubmitTextTo3DAvatarVideoTaskResponse) SetStatusCode(v int32) *SubmitTextTo3DAvatarVideoTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitTextTo3DAvatarVideoTaskResponse) SetBody(v *SubmitTextTo3DAvatarVideoTaskResponseBody) *SubmitTextTo3DAvatarVideoTaskResponse {
	s.Body = v
	return s
}

type Update2dAvatarRequest struct {
	Callback    *bool   `json:"Callback,omitempty" xml:"Callback,omitempty"`
	Code        *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Image       *string `json:"Image,omitempty" xml:"Image,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Orientation *int32  `json:"Orientation,omitempty" xml:"Orientation,omitempty"`
	Portrait    *string `json:"Portrait,omitempty" xml:"Portrait,omitempty"`
	TenantId    *int64  `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	Transparent *bool   `json:"Transparent,omitempty" xml:"Transparent,omitempty"`
	Video       *string `json:"Video,omitempty" xml:"Video,omitempty"`
}

func (s Update2dAvatarRequest) String() string {
	return tea.Prettify(s)
}

func (s Update2dAvatarRequest) GoString() string {
	return s.String()
}

func (s *Update2dAvatarRequest) SetCallback(v bool) *Update2dAvatarRequest {
	s.Callback = &v
	return s
}

func (s *Update2dAvatarRequest) SetCode(v string) *Update2dAvatarRequest {
	s.Code = &v
	return s
}

func (s *Update2dAvatarRequest) SetDescription(v string) *Update2dAvatarRequest {
	s.Description = &v
	return s
}

func (s *Update2dAvatarRequest) SetImage(v string) *Update2dAvatarRequest {
	s.Image = &v
	return s
}

func (s *Update2dAvatarRequest) SetName(v string) *Update2dAvatarRequest {
	s.Name = &v
	return s
}

func (s *Update2dAvatarRequest) SetOrientation(v int32) *Update2dAvatarRequest {
	s.Orientation = &v
	return s
}

func (s *Update2dAvatarRequest) SetPortrait(v string) *Update2dAvatarRequest {
	s.Portrait = &v
	return s
}

func (s *Update2dAvatarRequest) SetTenantId(v int64) *Update2dAvatarRequest {
	s.TenantId = &v
	return s
}

func (s *Update2dAvatarRequest) SetTransparent(v bool) *Update2dAvatarRequest {
	s.Transparent = &v
	return s
}

func (s *Update2dAvatarRequest) SetVideo(v string) *Update2dAvatarRequest {
	s.Video = &v
	return s
}

type Update2dAvatarResponseBody struct {
	Code      *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *Update2dAvatarResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                           `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s Update2dAvatarResponseBody) String() string {
	return tea.Prettify(s)
}

func (s Update2dAvatarResponseBody) GoString() string {
	return s.String()
}

func (s *Update2dAvatarResponseBody) SetCode(v string) *Update2dAvatarResponseBody {
	s.Code = &v
	return s
}

func (s *Update2dAvatarResponseBody) SetData(v *Update2dAvatarResponseBodyData) *Update2dAvatarResponseBody {
	s.Data = v
	return s
}

func (s *Update2dAvatarResponseBody) SetMessage(v string) *Update2dAvatarResponseBody {
	s.Message = &v
	return s
}

func (s *Update2dAvatarResponseBody) SetRequestId(v string) *Update2dAvatarResponseBody {
	s.RequestId = &v
	return s
}

func (s *Update2dAvatarResponseBody) SetSuccess(v bool) *Update2dAvatarResponseBody {
	s.Success = &v
	return s
}

type Update2dAvatarResponseBodyData struct {
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s Update2dAvatarResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s Update2dAvatarResponseBodyData) GoString() string {
	return s.String()
}

func (s *Update2dAvatarResponseBodyData) SetCode(v string) *Update2dAvatarResponseBodyData {
	s.Code = &v
	return s
}

type Update2dAvatarResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *Update2dAvatarResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s Update2dAvatarResponse) String() string {
	return tea.Prettify(s)
}

func (s Update2dAvatarResponse) GoString() string {
	return s.String()
}

func (s *Update2dAvatarResponse) SetHeaders(v map[string]*string) *Update2dAvatarResponse {
	s.Headers = v
	return s
}

func (s *Update2dAvatarResponse) SetStatusCode(v int32) *Update2dAvatarResponse {
	s.StatusCode = &v
	return s
}

func (s *Update2dAvatarResponse) SetBody(v *Update2dAvatarResponseBody) *Update2dAvatarResponse {
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
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("avatar"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) CancelVideoTaskWithOptions(tmpReq *CancelVideoTaskRequest, runtime *util.RuntimeOptions) (_result *CancelVideoTaskResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CancelVideoTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.App)) {
		request.AppShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.App, tea.String("App"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppShrink)) {
		query["App"] = request.AppShrink
	}

	if !tea.BoolValue(util.IsUnset(request.TaskUuid)) {
		query["TaskUuid"] = request.TaskUuid
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		query["TenantId"] = request.TenantId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CancelVideoTask"),
		Version:     tea.String("2022-01-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CancelVideoTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CancelVideoTask(request *CancelVideoTaskRequest) (_result *CancelVideoTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CancelVideoTaskResponse{}
	_body, _err := client.CancelVideoTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ClientAuthWithOptions(request *ClientAuthRequest, runtime *util.RuntimeOptions) (_result *ClientAuthResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceId)) {
		query["DeviceId"] = request.DeviceId
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceInfo)) {
		query["DeviceInfo"] = request.DeviceInfo
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceType)) {
		query["DeviceType"] = request.DeviceType
	}

	if !tea.BoolValue(util.IsUnset(request.License)) {
		query["License"] = request.License
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		query["TenantId"] = request.TenantId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ClientAuth"),
		Version:     tea.String("2022-01-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ClientAuthResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ClientAuth(request *ClientAuthRequest) (_result *ClientAuthResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ClientAuthResponse{}
	_body, _err := client.ClientAuthWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ClientStartWithOptions(request *ClientStartRequest, runtime *util.RuntimeOptions) (_result *ClientStartResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		query["TenantId"] = request.TenantId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ClientStart"),
		Version:     tea.String("2022-01-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ClientStartResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ClientStart(request *ClientStartRequest) (_result *ClientStartResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ClientStartResponse{}
	_body, _err := client.ClientStartWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CloseTimedResetOperateWithOptions(request *CloseTimedResetOperateRequest, runtime *util.RuntimeOptions) (_result *CloseTimedResetOperateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		query["TenantId"] = request.TenantId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CloseTimedResetOperate"),
		Version:     tea.String("2022-01-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CloseTimedResetOperateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CloseTimedResetOperate(request *CloseTimedResetOperateRequest) (_result *CloseTimedResetOperateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CloseTimedResetOperateResponse{}
	_body, _err := client.CloseTimedResetOperateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) Create2dAvatarWithOptions(request *Create2dAvatarRequest, runtime *util.RuntimeOptions) (_result *Create2dAvatarResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Callback)) {
		query["Callback"] = request.Callback
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Image)) {
		query["Image"] = request.Image
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Orientation)) {
		query["Orientation"] = request.Orientation
	}

	if !tea.BoolValue(util.IsUnset(request.Portrait)) {
		query["Portrait"] = request.Portrait
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		query["TenantId"] = request.TenantId
	}

	if !tea.BoolValue(util.IsUnset(request.Transparent)) {
		query["Transparent"] = request.Transparent
	}

	if !tea.BoolValue(util.IsUnset(request.Video)) {
		query["Video"] = request.Video
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("Create2dAvatar"),
		Version:     tea.String("2022-01-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &Create2dAvatarResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) Create2dAvatar(request *Create2dAvatarRequest) (_result *Create2dAvatarResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &Create2dAvatarResponse{}
	_body, _err := client.Create2dAvatarWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteAvatarWithOptions(request *DeleteAvatarRequest, runtime *util.RuntimeOptions) (_result *DeleteAvatarResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Code)) {
		query["Code"] = request.Code
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		query["TenantId"] = request.TenantId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAvatar"),
		Version:     tea.String("2022-01-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteAvatarResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteAvatar(request *DeleteAvatarRequest) (_result *DeleteAvatarResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAvatarResponse{}
	_body, _err := client.DeleteAvatarWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DuplexDecisionWithOptions(tmpReq *DuplexDecisionRequest, runtime *util.RuntimeOptions) (_result *DuplexDecisionResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DuplexDecisionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.CustomKeywords)) {
		request.CustomKeywordsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CustomKeywords, tea.String("CustomKeywords"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.DialogContext)) {
		request.DialogContextShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DialogContext, tea.String("DialogContext"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.BizRequestId)) {
		query["BizRequestId"] = request.BizRequestId
	}

	if !tea.BoolValue(util.IsUnset(request.CallTime)) {
		query["CallTime"] = request.CallTime
	}

	if !tea.BoolValue(util.IsUnset(request.CustomKeywordsShrink)) {
		query["CustomKeywords"] = request.CustomKeywordsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.DialogContextShrink)) {
		query["DialogContext"] = request.DialogContextShrink
	}

	if !tea.BoolValue(util.IsUnset(request.DialogStatus)) {
		query["DialogStatus"] = request.DialogStatus
	}

	if !tea.BoolValue(util.IsUnset(request.InterruptType)) {
		query["InterruptType"] = request.InterruptType
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		query["SessionId"] = request.SessionId
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		query["TenantId"] = request.TenantId
	}

	if !tea.BoolValue(util.IsUnset(request.Text)) {
		query["Text"] = request.Text
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DuplexDecision"),
		Version:     tea.String("2022-01-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DuplexDecisionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DuplexDecision(request *DuplexDecisionRequest) (_result *DuplexDecisionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DuplexDecisionResponse{}
	_body, _err := client.DuplexDecisionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetVideoTaskInfoWithOptions(tmpReq *GetVideoTaskInfoRequest, runtime *util.RuntimeOptions) (_result *GetVideoTaskInfoResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetVideoTaskInfoShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.App)) {
		request.AppShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.App, tea.String("App"), tea.String("json"))
	}

	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetVideoTaskInfo"),
		Version:     tea.String("2022-01-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetVideoTaskInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetVideoTaskInfo(request *GetVideoTaskInfoRequest) (_result *GetVideoTaskInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetVideoTaskInfoResponse{}
	_body, _err := client.GetVideoTaskInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) LicenseAuthWithOptions(request *LicenseAuthRequest, runtime *util.RuntimeOptions) (_result *LicenseAuthResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.License)) {
		query["License"] = request.License
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		query["TenantId"] = request.TenantId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("LicenseAuth"),
		Version:     tea.String("2022-01-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &LicenseAuthResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) LicenseAuth(request *LicenseAuthRequest) (_result *LicenseAuthResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &LicenseAuthResponse{}
	_body, _err := client.LicenseAuthWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryAvatarWithOptions(request *QueryAvatarRequest, runtime *util.RuntimeOptions) (_result *QueryAvatarResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Code)) {
		query["Code"] = request.Code
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		query["TenantId"] = request.TenantId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryAvatar"),
		Version:     tea.String("2022-01-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryAvatarResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryAvatar(request *QueryAvatarRequest) (_result *QueryAvatarResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryAvatarResponse{}
	_body, _err := client.QueryAvatarWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryAvatarListWithOptions(request *QueryAvatarListRequest, runtime *util.RuntimeOptions) (_result *QueryAvatarListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ModelType)) {
		query["ModelType"] = request.ModelType
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		query["TenantId"] = request.TenantId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryAvatarList"),
		Version:     tea.String("2022-01-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryAvatarListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryAvatarList(request *QueryAvatarListRequest) (_result *QueryAvatarListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryAvatarListResponse{}
	_body, _err := client.QueryAvatarListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryRunningInstanceWithOptions(tmpReq *QueryRunningInstanceRequest, runtime *util.RuntimeOptions) (_result *QueryRunningInstanceResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &QueryRunningInstanceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.App)) {
		request.AppShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.App, tea.String("App"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppShrink)) {
		query["App"] = request.AppShrink
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		query["SessionId"] = request.SessionId
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		query["TenantId"] = request.TenantId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryRunningInstance"),
		Version:     tea.String("2022-01-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryRunningInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryRunningInstance(request *QueryRunningInstanceRequest) (_result *QueryRunningInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryRunningInstanceResponse{}
	_body, _err := client.QueryRunningInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryTimedResetOperateStatusWithOptions(request *QueryTimedResetOperateStatusRequest, runtime *util.RuntimeOptions) (_result *QueryTimedResetOperateStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		query["TenantId"] = request.TenantId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryTimedResetOperateStatus"),
		Version:     tea.String("2022-01-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryTimedResetOperateStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryTimedResetOperateStatus(request *QueryTimedResetOperateStatusRequest) (_result *QueryTimedResetOperateStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryTimedResetOperateStatusResponse{}
	_body, _err := client.QueryTimedResetOperateStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryVideoTaskInfoWithOptions(tmpReq *QueryVideoTaskInfoRequest, runtime *util.RuntimeOptions) (_result *QueryVideoTaskInfoResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &QueryVideoTaskInfoShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.App)) {
		request.AppShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.App, tea.String("App"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppShrink)) {
		query["App"] = request.AppShrink
	}

	if !tea.BoolValue(util.IsUnset(request.OrderById)) {
		query["OrderById"] = request.OrderById
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.TaskUuid)) {
		query["TaskUuid"] = request.TaskUuid
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		query["TenantId"] = request.TenantId
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		query["Title"] = request.Title
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryVideoTaskInfo"),
		Version:     tea.String("2022-01-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryVideoTaskInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryVideoTaskInfo(request *QueryVideoTaskInfoRequest) (_result *QueryVideoTaskInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryVideoTaskInfoResponse{}
	_body, _err := client.QueryVideoTaskInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) Render3dAvatarWithOptions(request *Render3dAvatarRequest, runtime *util.RuntimeOptions) (_result *Render3dAvatarResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.Code)) {
		query["Code"] = request.Code
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		query["TenantId"] = request.TenantId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("Render3dAvatar"),
		Version:     tea.String("2022-01-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &Render3dAvatarResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) Render3dAvatar(request *Render3dAvatarRequest) (_result *Render3dAvatarResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &Render3dAvatarResponse{}
	_body, _err := client.Render3dAvatarWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SendCommandWithOptions(tmpReq *SendCommandRequest, runtime *util.RuntimeOptions) (_result *SendCommandResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &SendCommandShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Content)) {
		request.ContentShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Content, tea.String("Content"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Code)) {
		query["Code"] = request.Code
	}

	if !tea.BoolValue(util.IsUnset(request.ContentShrink)) {
		query["Content"] = request.ContentShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Feedback)) {
		query["Feedback"] = request.Feedback
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		query["SessionId"] = request.SessionId
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		query["TenantId"] = request.TenantId
	}

	if !tea.BoolValue(util.IsUnset(request.UniqueCode)) {
		query["UniqueCode"] = request.UniqueCode
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SendCommand"),
		Version:     tea.String("2022-01-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SendCommandResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SendCommand(request *SendCommandRequest) (_result *SendCommandResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SendCommandResponse{}
	_body, _err := client.SendCommandWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SendMessageWithOptions(tmpReq *SendMessageRequest, runtime *util.RuntimeOptions) (_result *SendMessageResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &SendMessageShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.StreamExtension)) {
		request.StreamExtensionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.StreamExtension, tea.String("StreamExtension"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.TextRequest)) {
		request.TextRequestShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TextRequest, tea.String("TextRequest"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.VAMLRequest)) {
		request.VAMLRequestShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.VAMLRequest, tea.String("VAMLRequest"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Feedback)) {
		query["Feedback"] = request.Feedback
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		query["SessionId"] = request.SessionId
	}

	if !tea.BoolValue(util.IsUnset(request.StreamExtensionShrink)) {
		query["StreamExtension"] = request.StreamExtensionShrink
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		query["TenantId"] = request.TenantId
	}

	if !tea.BoolValue(util.IsUnset(request.TextRequestShrink)) {
		query["TextRequest"] = request.TextRequestShrink
	}

	if !tea.BoolValue(util.IsUnset(request.VAMLRequestShrink)) {
		query["VAMLRequest"] = request.VAMLRequestShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SendMessage"),
		Version:     tea.String("2022-01-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SendMessageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SendMessage(request *SendMessageRequest) (_result *SendMessageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SendMessageResponse{}
	_body, _err := client.SendMessageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SendTextWithOptions(tmpReq *SendTextRequest, runtime *util.RuntimeOptions) (_result *SendTextResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &SendTextShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.StreamExtension)) {
		request.StreamExtensionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.StreamExtension, tea.String("StreamExtension"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Feedback)) {
		query["Feedback"] = request.Feedback
	}

	if !tea.BoolValue(util.IsUnset(request.Interrupt)) {
		query["Interrupt"] = request.Interrupt
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		query["SessionId"] = request.SessionId
	}

	if !tea.BoolValue(util.IsUnset(request.StreamExtensionShrink)) {
		query["StreamExtension"] = request.StreamExtensionShrink
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		query["TenantId"] = request.TenantId
	}

	if !tea.BoolValue(util.IsUnset(request.Text)) {
		query["Text"] = request.Text
	}

	if !tea.BoolValue(util.IsUnset(request.UniqueCode)) {
		query["UniqueCode"] = request.UniqueCode
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SendText"),
		Version:     tea.String("2022-01-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SendTextResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SendText(request *SendTextRequest) (_result *SendTextResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SendTextResponse{}
	_body, _err := client.SendTextWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SendVamlWithOptions(request *SendVamlRequest, runtime *util.RuntimeOptions) (_result *SendVamlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		query["SessionId"] = request.SessionId
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		query["TenantId"] = request.TenantId
	}

	if !tea.BoolValue(util.IsUnset(request.Vaml)) {
		query["Vaml"] = request.Vaml
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SendVaml"),
		Version:     tea.String("2022-01-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SendVamlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SendVaml(request *SendVamlRequest) (_result *SendVamlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SendVamlResponse{}
	_body, _err := client.SendVamlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartInstanceWithOptions(tmpReq *StartInstanceRequest, runtime *util.RuntimeOptions) (_result *StartInstanceResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &StartInstanceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.App)) {
		request.AppShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.App, tea.String("App"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Channel)) {
		request.ChannelShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Channel, tea.String("Channel"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.CommandRequest)) {
		request.CommandRequestShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CommandRequest, tea.String("CommandRequest"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.User)) {
		request.UserShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.User, tea.String("User"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppShrink)) {
		query["App"] = request.AppShrink
	}

	if !tea.BoolValue(util.IsUnset(request.BizId)) {
		query["BizId"] = request.BizId
	}

	if !tea.BoolValue(util.IsUnset(request.ChannelShrink)) {
		query["Channel"] = request.ChannelShrink
	}

	if !tea.BoolValue(util.IsUnset(request.CommandRequestShrink)) {
		query["CommandRequest"] = request.CommandRequestShrink
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		query["TenantId"] = request.TenantId
	}

	if !tea.BoolValue(util.IsUnset(request.UserShrink)) {
		query["User"] = request.UserShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StartInstance"),
		Version:     tea.String("2022-01-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StartInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartInstance(request *StartInstanceRequest) (_result *StartInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartInstanceResponse{}
	_body, _err := client.StartInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartTimedResetOperateWithOptions(request *StartTimedResetOperateRequest, runtime *util.RuntimeOptions) (_result *StartTimedResetOperateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		query["TenantId"] = request.TenantId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StartTimedResetOperate"),
		Version:     tea.String("2022-01-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StartTimedResetOperateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartTimedResetOperate(request *StartTimedResetOperateRequest) (_result *StartTimedResetOperateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartTimedResetOperateResponse{}
	_body, _err := client.StartTimedResetOperateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopInstanceWithOptions(request *StopInstanceRequest, runtime *util.RuntimeOptions) (_result *StopInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		query["SessionId"] = request.SessionId
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		query["TenantId"] = request.TenantId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StopInstance"),
		Version:     tea.String("2022-01-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StopInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopInstance(request *StopInstanceRequest) (_result *StopInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopInstanceResponse{}
	_body, _err := client.StopInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitAudioTo2DAvatarVideoTaskWithOptions(tmpReq *SubmitAudioTo2DAvatarVideoTaskRequest, runtime *util.RuntimeOptions) (_result *SubmitAudioTo2DAvatarVideoTaskResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &SubmitAudioTo2DAvatarVideoTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.App)) {
		request.AppShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.App, tea.String("App"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.AvatarInfo)) {
		request.AvatarInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AvatarInfo, tea.String("AvatarInfo"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.VideoInfo)) {
		request.VideoInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.VideoInfo, tea.String("VideoInfo"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppShrink)) {
		query["App"] = request.AppShrink
	}

	if !tea.BoolValue(util.IsUnset(request.AvatarInfoShrink)) {
		query["AvatarInfo"] = request.AvatarInfoShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Callback)) {
		query["Callback"] = request.Callback
	}

	if !tea.BoolValue(util.IsUnset(request.CallbackParams)) {
		query["CallbackParams"] = request.CallbackParams
	}

	if !tea.BoolValue(util.IsUnset(request.ExtParams)) {
		query["ExtParams"] = request.ExtParams
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		query["TenantId"] = request.TenantId
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		query["Title"] = request.Title
	}

	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	if !tea.BoolValue(util.IsUnset(request.VideoInfoShrink)) {
		query["VideoInfo"] = request.VideoInfoShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SubmitAudioTo2DAvatarVideoTask"),
		Version:     tea.String("2022-01-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitAudioTo2DAvatarVideoTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitAudioTo2DAvatarVideoTask(request *SubmitAudioTo2DAvatarVideoTaskRequest) (_result *SubmitAudioTo2DAvatarVideoTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitAudioTo2DAvatarVideoTaskResponse{}
	_body, _err := client.SubmitAudioTo2DAvatarVideoTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitAudioTo3DAvatarVideoTaskWithOptions(tmpReq *SubmitAudioTo3DAvatarVideoTaskRequest, runtime *util.RuntimeOptions) (_result *SubmitAudioTo3DAvatarVideoTaskResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &SubmitAudioTo3DAvatarVideoTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.App)) {
		request.AppShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.App, tea.String("App"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.AvatarInfo)) {
		request.AvatarInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AvatarInfo, tea.String("AvatarInfo"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.VideoInfo)) {
		request.VideoInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.VideoInfo, tea.String("VideoInfo"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppShrink)) {
		query["App"] = request.AppShrink
	}

	if !tea.BoolValue(util.IsUnset(request.AvatarInfoShrink)) {
		query["AvatarInfo"] = request.AvatarInfoShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Callback)) {
		query["Callback"] = request.Callback
	}

	if !tea.BoolValue(util.IsUnset(request.CallbackParams)) {
		query["CallbackParams"] = request.CallbackParams
	}

	if !tea.BoolValue(util.IsUnset(request.ExtParams)) {
		query["ExtParams"] = request.ExtParams
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		query["TenantId"] = request.TenantId
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		query["Title"] = request.Title
	}

	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	if !tea.BoolValue(util.IsUnset(request.VideoInfoShrink)) {
		query["VideoInfo"] = request.VideoInfoShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SubmitAudioTo3DAvatarVideoTask"),
		Version:     tea.String("2022-01-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitAudioTo3DAvatarVideoTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitAudioTo3DAvatarVideoTask(request *SubmitAudioTo3DAvatarVideoTaskRequest) (_result *SubmitAudioTo3DAvatarVideoTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitAudioTo3DAvatarVideoTaskResponse{}
	_body, _err := client.SubmitAudioTo3DAvatarVideoTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitAvatarVideoTaskWithOptions(tmpReq *SubmitAvatarVideoTaskRequest, runtime *util.RuntimeOptions) (_result *SubmitAvatarVideoTaskResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &SubmitAvatarVideoTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.App)) {
		request.AppShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.App, tea.String("App"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppShrink)) {
		query["App"] = request.AppShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Callback)) {
		query["Callback"] = request.Callback
	}

	if !tea.BoolValue(util.IsUnset(request.CallbackParams)) {
		query["CallbackParams"] = request.CallbackParams
	}

	if !tea.BoolValue(util.IsUnset(request.ExtParams)) {
		query["ExtParams"] = request.ExtParams
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		query["TenantId"] = request.TenantId
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		query["Title"] = request.Title
	}

	if !tea.BoolValue(util.IsUnset(request.VideoParams)) {
		query["VideoParams"] = request.VideoParams
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SubmitAvatarVideoTask"),
		Version:     tea.String("2022-01-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitAvatarVideoTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitAvatarVideoTask(request *SubmitAvatarVideoTaskRequest) (_result *SubmitAvatarVideoTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitAvatarVideoTaskResponse{}
	_body, _err := client.SubmitAvatarVideoTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitTextTo2DAvatarVideoTaskWithOptions(tmpReq *SubmitTextTo2DAvatarVideoTaskRequest, runtime *util.RuntimeOptions) (_result *SubmitTextTo2DAvatarVideoTaskResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &SubmitTextTo2DAvatarVideoTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.App)) {
		request.AppShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.App, tea.String("App"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.AudioInfo)) {
		request.AudioInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AudioInfo, tea.String("AudioInfo"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.AvatarInfo)) {
		request.AvatarInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AvatarInfo, tea.String("AvatarInfo"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.VideoInfo)) {
		request.VideoInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.VideoInfo, tea.String("VideoInfo"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppShrink)) {
		query["App"] = request.AppShrink
	}

	if !tea.BoolValue(util.IsUnset(request.AudioInfoShrink)) {
		query["AudioInfo"] = request.AudioInfoShrink
	}

	if !tea.BoolValue(util.IsUnset(request.AvatarInfoShrink)) {
		query["AvatarInfo"] = request.AvatarInfoShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Callback)) {
		query["Callback"] = request.Callback
	}

	if !tea.BoolValue(util.IsUnset(request.CallbackParams)) {
		query["CallbackParams"] = request.CallbackParams
	}

	if !tea.BoolValue(util.IsUnset(request.ExtParams)) {
		query["ExtParams"] = request.ExtParams
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		query["TenantId"] = request.TenantId
	}

	if !tea.BoolValue(util.IsUnset(request.Text)) {
		query["Text"] = request.Text
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		query["Title"] = request.Title
	}

	if !tea.BoolValue(util.IsUnset(request.VideoInfoShrink)) {
		query["VideoInfo"] = request.VideoInfoShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SubmitTextTo2DAvatarVideoTask"),
		Version:     tea.String("2022-01-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitTextTo2DAvatarVideoTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitTextTo2DAvatarVideoTask(request *SubmitTextTo2DAvatarVideoTaskRequest) (_result *SubmitTextTo2DAvatarVideoTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitTextTo2DAvatarVideoTaskResponse{}
	_body, _err := client.SubmitTextTo2DAvatarVideoTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitTextTo3DAvatarVideoTaskWithOptions(tmpReq *SubmitTextTo3DAvatarVideoTaskRequest, runtime *util.RuntimeOptions) (_result *SubmitTextTo3DAvatarVideoTaskResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &SubmitTextTo3DAvatarVideoTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.App)) {
		request.AppShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.App, tea.String("App"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.AudioInfo)) {
		request.AudioInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AudioInfo, tea.String("AudioInfo"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.AvatarInfo)) {
		request.AvatarInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AvatarInfo, tea.String("AvatarInfo"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.VideoInfo)) {
		request.VideoInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.VideoInfo, tea.String("VideoInfo"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppShrink)) {
		query["App"] = request.AppShrink
	}

	if !tea.BoolValue(util.IsUnset(request.AudioInfoShrink)) {
		query["AudioInfo"] = request.AudioInfoShrink
	}

	if !tea.BoolValue(util.IsUnset(request.AvatarInfoShrink)) {
		query["AvatarInfo"] = request.AvatarInfoShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Callback)) {
		query["Callback"] = request.Callback
	}

	if !tea.BoolValue(util.IsUnset(request.CallbackParams)) {
		query["CallbackParams"] = request.CallbackParams
	}

	if !tea.BoolValue(util.IsUnset(request.ExtParams)) {
		query["ExtParams"] = request.ExtParams
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		query["TenantId"] = request.TenantId
	}

	if !tea.BoolValue(util.IsUnset(request.Text)) {
		query["Text"] = request.Text
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		query["Title"] = request.Title
	}

	if !tea.BoolValue(util.IsUnset(request.VideoInfoShrink)) {
		query["VideoInfo"] = request.VideoInfoShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SubmitTextTo3DAvatarVideoTask"),
		Version:     tea.String("2022-01-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitTextTo3DAvatarVideoTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitTextTo3DAvatarVideoTask(request *SubmitTextTo3DAvatarVideoTaskRequest) (_result *SubmitTextTo3DAvatarVideoTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitTextTo3DAvatarVideoTaskResponse{}
	_body, _err := client.SubmitTextTo3DAvatarVideoTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) Update2dAvatarWithOptions(request *Update2dAvatarRequest, runtime *util.RuntimeOptions) (_result *Update2dAvatarResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Callback)) {
		query["Callback"] = request.Callback
	}

	if !tea.BoolValue(util.IsUnset(request.Code)) {
		query["Code"] = request.Code
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Image)) {
		query["Image"] = request.Image
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Orientation)) {
		query["Orientation"] = request.Orientation
	}

	if !tea.BoolValue(util.IsUnset(request.Portrait)) {
		query["Portrait"] = request.Portrait
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		query["TenantId"] = request.TenantId
	}

	if !tea.BoolValue(util.IsUnset(request.Transparent)) {
		query["Transparent"] = request.Transparent
	}

	if !tea.BoolValue(util.IsUnset(request.Video)) {
		query["Video"] = request.Video
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("Update2dAvatar"),
		Version:     tea.String("2022-01-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &Update2dAvatarResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) Update2dAvatar(request *Update2dAvatarRequest) (_result *Update2dAvatarResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &Update2dAvatarResponse{}
	_body, _err := client.Update2dAvatarWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
