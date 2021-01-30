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

type AddRtcAccountRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	DeviceId             *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
}

func (s AddRtcAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s AddRtcAccountRequest) GoString() string {
	return s.String()
}

func (s *AddRtcAccountRequest) SetOwnerId(v int64) *AddRtcAccountRequest {
	s.OwnerId = &v
	return s
}

func (s *AddRtcAccountRequest) SetResourceOwnerAccount(v string) *AddRtcAccountRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AddRtcAccountRequest) SetResourceOwnerId(v int64) *AddRtcAccountRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AddRtcAccountRequest) SetDeviceId(v string) *AddRtcAccountRequest {
	s.DeviceId = &v
	return s
}

type AddRtcAccountResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Module    *string `json:"Module,omitempty" xml:"Module,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s AddRtcAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddRtcAccountResponseBody) GoString() string {
	return s.String()
}

func (s *AddRtcAccountResponseBody) SetMessage(v string) *AddRtcAccountResponseBody {
	s.Message = &v
	return s
}

func (s *AddRtcAccountResponseBody) SetRequestId(v string) *AddRtcAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddRtcAccountResponseBody) SetModule(v string) *AddRtcAccountResponseBody {
	s.Module = &v
	return s
}

func (s *AddRtcAccountResponseBody) SetCode(v string) *AddRtcAccountResponseBody {
	s.Code = &v
	return s
}

type AddRtcAccountResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddRtcAccountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddRtcAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s AddRtcAccountResponse) GoString() string {
	return s.String()
}

func (s *AddRtcAccountResponse) SetHeaders(v map[string]*string) *AddRtcAccountResponse {
	s.Headers = v
	return s
}

func (s *AddRtcAccountResponse) SetBody(v *AddRtcAccountResponseBody) *AddRtcAccountResponse {
	s.Body = v
	return s
}

type BatchRobotSmartCallRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	CalledShowNumber     *string `json:"CalledShowNumber,omitempty" xml:"CalledShowNumber,omitempty"`
	CorpName             *string `json:"CorpName,omitempty" xml:"CorpName,omitempty"`
	CalledNumber         *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	DialogId             *string `json:"DialogId,omitempty" xml:"DialogId,omitempty"`
	EarlyMediaAsr        *bool   `json:"EarlyMediaAsr,omitempty" xml:"EarlyMediaAsr,omitempty"`
	TaskName             *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	ScheduleTime         *int64  `json:"ScheduleTime,omitempty" xml:"ScheduleTime,omitempty"`
	ScheduleCall         *bool   `json:"ScheduleCall,omitempty" xml:"ScheduleCall,omitempty"`
	TtsParam             *string `json:"TtsParam,omitempty" xml:"TtsParam,omitempty"`
	TtsParamHead         *string `json:"TtsParamHead,omitempty" xml:"TtsParamHead,omitempty"`
	IsSelfLine           *bool   `json:"IsSelfLine,omitempty" xml:"IsSelfLine,omitempty"`
}

func (s BatchRobotSmartCallRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchRobotSmartCallRequest) GoString() string {
	return s.String()
}

func (s *BatchRobotSmartCallRequest) SetOwnerId(v int64) *BatchRobotSmartCallRequest {
	s.OwnerId = &v
	return s
}

func (s *BatchRobotSmartCallRequest) SetResourceOwnerAccount(v string) *BatchRobotSmartCallRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *BatchRobotSmartCallRequest) SetResourceOwnerId(v int64) *BatchRobotSmartCallRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *BatchRobotSmartCallRequest) SetCalledShowNumber(v string) *BatchRobotSmartCallRequest {
	s.CalledShowNumber = &v
	return s
}

func (s *BatchRobotSmartCallRequest) SetCorpName(v string) *BatchRobotSmartCallRequest {
	s.CorpName = &v
	return s
}

func (s *BatchRobotSmartCallRequest) SetCalledNumber(v string) *BatchRobotSmartCallRequest {
	s.CalledNumber = &v
	return s
}

func (s *BatchRobotSmartCallRequest) SetDialogId(v string) *BatchRobotSmartCallRequest {
	s.DialogId = &v
	return s
}

func (s *BatchRobotSmartCallRequest) SetEarlyMediaAsr(v bool) *BatchRobotSmartCallRequest {
	s.EarlyMediaAsr = &v
	return s
}

func (s *BatchRobotSmartCallRequest) SetTaskName(v string) *BatchRobotSmartCallRequest {
	s.TaskName = &v
	return s
}

func (s *BatchRobotSmartCallRequest) SetScheduleTime(v int64) *BatchRobotSmartCallRequest {
	s.ScheduleTime = &v
	return s
}

func (s *BatchRobotSmartCallRequest) SetScheduleCall(v bool) *BatchRobotSmartCallRequest {
	s.ScheduleCall = &v
	return s
}

func (s *BatchRobotSmartCallRequest) SetTtsParam(v string) *BatchRobotSmartCallRequest {
	s.TtsParam = &v
	return s
}

func (s *BatchRobotSmartCallRequest) SetTtsParamHead(v string) *BatchRobotSmartCallRequest {
	s.TtsParamHead = &v
	return s
}

func (s *BatchRobotSmartCallRequest) SetIsSelfLine(v bool) *BatchRobotSmartCallRequest {
	s.IsSelfLine = &v
	return s
}

type BatchRobotSmartCallResponseBody struct {
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s BatchRobotSmartCallResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchRobotSmartCallResponseBody) GoString() string {
	return s.String()
}

func (s *BatchRobotSmartCallResponseBody) SetTaskId(v string) *BatchRobotSmartCallResponseBody {
	s.TaskId = &v
	return s
}

func (s *BatchRobotSmartCallResponseBody) SetMessage(v string) *BatchRobotSmartCallResponseBody {
	s.Message = &v
	return s
}

func (s *BatchRobotSmartCallResponseBody) SetRequestId(v string) *BatchRobotSmartCallResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchRobotSmartCallResponseBody) SetCode(v string) *BatchRobotSmartCallResponseBody {
	s.Code = &v
	return s
}

type BatchRobotSmartCallResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BatchRobotSmartCallResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchRobotSmartCallResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchRobotSmartCallResponse) GoString() string {
	return s.String()
}

func (s *BatchRobotSmartCallResponse) SetHeaders(v map[string]*string) *BatchRobotSmartCallResponse {
	s.Headers = v
	return s
}

func (s *BatchRobotSmartCallResponse) SetBody(v *BatchRobotSmartCallResponseBody) *BatchRobotSmartCallResponse {
	s.Body = v
	return s
}

type BindNumberAndVoipIdRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	PhoneNumber          *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	VoipId               *string `json:"VoipId,omitempty" xml:"VoipId,omitempty"`
}

func (s BindNumberAndVoipIdRequest) String() string {
	return tea.Prettify(s)
}

func (s BindNumberAndVoipIdRequest) GoString() string {
	return s.String()
}

func (s *BindNumberAndVoipIdRequest) SetOwnerId(v int64) *BindNumberAndVoipIdRequest {
	s.OwnerId = &v
	return s
}

func (s *BindNumberAndVoipIdRequest) SetResourceOwnerAccount(v string) *BindNumberAndVoipIdRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *BindNumberAndVoipIdRequest) SetResourceOwnerId(v int64) *BindNumberAndVoipIdRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *BindNumberAndVoipIdRequest) SetPhoneNumber(v string) *BindNumberAndVoipIdRequest {
	s.PhoneNumber = &v
	return s
}

func (s *BindNumberAndVoipIdRequest) SetVoipId(v string) *BindNumberAndVoipIdRequest {
	s.VoipId = &v
	return s
}

type BindNumberAndVoipIdResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Module    *string `json:"Module,omitempty" xml:"Module,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s BindNumberAndVoipIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BindNumberAndVoipIdResponseBody) GoString() string {
	return s.String()
}

func (s *BindNumberAndVoipIdResponseBody) SetMessage(v string) *BindNumberAndVoipIdResponseBody {
	s.Message = &v
	return s
}

func (s *BindNumberAndVoipIdResponseBody) SetRequestId(v string) *BindNumberAndVoipIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *BindNumberAndVoipIdResponseBody) SetModule(v string) *BindNumberAndVoipIdResponseBody {
	s.Module = &v
	return s
}

func (s *BindNumberAndVoipIdResponseBody) SetCode(v string) *BindNumberAndVoipIdResponseBody {
	s.Code = &v
	return s
}

type BindNumberAndVoipIdResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BindNumberAndVoipIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BindNumberAndVoipIdResponse) String() string {
	return tea.Prettify(s)
}

func (s BindNumberAndVoipIdResponse) GoString() string {
	return s.String()
}

func (s *BindNumberAndVoipIdResponse) SetHeaders(v map[string]*string) *BindNumberAndVoipIdResponse {
	s.Headers = v
	return s
}

func (s *BindNumberAndVoipIdResponse) SetBody(v *BindNumberAndVoipIdResponseBody) *BindNumberAndVoipIdResponse {
	s.Body = v
	return s
}

type CancelCallRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	CallId               *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
}

func (s CancelCallRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelCallRequest) GoString() string {
	return s.String()
}

func (s *CancelCallRequest) SetOwnerId(v int64) *CancelCallRequest {
	s.OwnerId = &v
	return s
}

func (s *CancelCallRequest) SetResourceOwnerAccount(v string) *CancelCallRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CancelCallRequest) SetResourceOwnerId(v int64) *CancelCallRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CancelCallRequest) SetCallId(v string) *CancelCallRequest {
	s.CallId = &v
	return s
}

type CancelCallResponseBody struct {
	Status    *bool   `json:"Status,omitempty" xml:"Status,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s CancelCallResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelCallResponseBody) GoString() string {
	return s.String()
}

func (s *CancelCallResponseBody) SetStatus(v bool) *CancelCallResponseBody {
	s.Status = &v
	return s
}

func (s *CancelCallResponseBody) SetMessage(v string) *CancelCallResponseBody {
	s.Message = &v
	return s
}

func (s *CancelCallResponseBody) SetRequestId(v string) *CancelCallResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelCallResponseBody) SetCode(v string) *CancelCallResponseBody {
	s.Code = &v
	return s
}

type CancelCallResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CancelCallResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CancelCallResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelCallResponse) GoString() string {
	return s.String()
}

func (s *CancelCallResponse) SetHeaders(v map[string]*string) *CancelCallResponse {
	s.Headers = v
	return s
}

func (s *CancelCallResponse) SetBody(v *CancelCallResponseBody) *CancelCallResponse {
	s.Body = v
	return s
}

type CancelOrderRobotTaskRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TaskId               *int64  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CancelOrderRobotTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelOrderRobotTaskRequest) GoString() string {
	return s.String()
}

func (s *CancelOrderRobotTaskRequest) SetOwnerId(v int64) *CancelOrderRobotTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *CancelOrderRobotTaskRequest) SetResourceOwnerAccount(v string) *CancelOrderRobotTaskRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CancelOrderRobotTaskRequest) SetResourceOwnerId(v int64) *CancelOrderRobotTaskRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CancelOrderRobotTaskRequest) SetTaskId(v int64) *CancelOrderRobotTaskRequest {
	s.TaskId = &v
	return s
}

type CancelOrderRobotTaskResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s CancelOrderRobotTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelOrderRobotTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CancelOrderRobotTaskResponseBody) SetMessage(v string) *CancelOrderRobotTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CancelOrderRobotTaskResponseBody) SetRequestId(v string) *CancelOrderRobotTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelOrderRobotTaskResponseBody) SetData(v string) *CancelOrderRobotTaskResponseBody {
	s.Data = &v
	return s
}

func (s *CancelOrderRobotTaskResponseBody) SetCode(v string) *CancelOrderRobotTaskResponseBody {
	s.Code = &v
	return s
}

type CancelOrderRobotTaskResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CancelOrderRobotTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CancelOrderRobotTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelOrderRobotTaskResponse) GoString() string {
	return s.String()
}

func (s *CancelOrderRobotTaskResponse) SetHeaders(v map[string]*string) *CancelOrderRobotTaskResponse {
	s.Headers = v
	return s
}

func (s *CancelOrderRobotTaskResponse) SetBody(v *CancelOrderRobotTaskResponseBody) *CancelOrderRobotTaskResponse {
	s.Body = v
	return s
}

type CancelRobotTaskRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TaskId               *int64  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CancelRobotTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelRobotTaskRequest) GoString() string {
	return s.String()
}

func (s *CancelRobotTaskRequest) SetOwnerId(v int64) *CancelRobotTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *CancelRobotTaskRequest) SetResourceOwnerAccount(v string) *CancelRobotTaskRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CancelRobotTaskRequest) SetResourceOwnerId(v int64) *CancelRobotTaskRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CancelRobotTaskRequest) SetTaskId(v int64) *CancelRobotTaskRequest {
	s.TaskId = &v
	return s
}

type CancelRobotTaskResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s CancelRobotTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelRobotTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CancelRobotTaskResponseBody) SetMessage(v string) *CancelRobotTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CancelRobotTaskResponseBody) SetRequestId(v string) *CancelRobotTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelRobotTaskResponseBody) SetData(v string) *CancelRobotTaskResponseBody {
	s.Data = &v
	return s
}

func (s *CancelRobotTaskResponseBody) SetCode(v string) *CancelRobotTaskResponseBody {
	s.Code = &v
	return s
}

type CancelRobotTaskResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CancelRobotTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CancelRobotTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelRobotTaskResponse) GoString() string {
	return s.String()
}

func (s *CancelRobotTaskResponse) SetHeaders(v map[string]*string) *CancelRobotTaskResponse {
	s.Headers = v
	return s
}

func (s *CancelRobotTaskResponse) SetBody(v *CancelRobotTaskResponseBody) *CancelRobotTaskResponse {
	s.Body = v
	return s
}

type ClickToDialRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	CallerShowNumber     *string `json:"CallerShowNumber,omitempty" xml:"CallerShowNumber,omitempty"`
	CallerNumber         *string `json:"CallerNumber,omitempty" xml:"CallerNumber,omitempty"`
	CalledShowNumber     *string `json:"CalledShowNumber,omitempty" xml:"CalledShowNumber,omitempty"`
	CalledNumber         *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	RecordFlag           *bool   `json:"RecordFlag,omitempty" xml:"RecordFlag,omitempty"`
	AsrFlag              *bool   `json:"AsrFlag,omitempty" xml:"AsrFlag,omitempty"`
	SessionTimeout       *int32  `json:"SessionTimeout,omitempty" xml:"SessionTimeout,omitempty"`
	AsrModelId           *string `json:"AsrModelId,omitempty" xml:"AsrModelId,omitempty"`
	OutId                *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
}

func (s ClickToDialRequest) String() string {
	return tea.Prettify(s)
}

func (s ClickToDialRequest) GoString() string {
	return s.String()
}

func (s *ClickToDialRequest) SetOwnerId(v int64) *ClickToDialRequest {
	s.OwnerId = &v
	return s
}

func (s *ClickToDialRequest) SetResourceOwnerAccount(v string) *ClickToDialRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ClickToDialRequest) SetResourceOwnerId(v int64) *ClickToDialRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ClickToDialRequest) SetCallerShowNumber(v string) *ClickToDialRequest {
	s.CallerShowNumber = &v
	return s
}

func (s *ClickToDialRequest) SetCallerNumber(v string) *ClickToDialRequest {
	s.CallerNumber = &v
	return s
}

func (s *ClickToDialRequest) SetCalledShowNumber(v string) *ClickToDialRequest {
	s.CalledShowNumber = &v
	return s
}

func (s *ClickToDialRequest) SetCalledNumber(v string) *ClickToDialRequest {
	s.CalledNumber = &v
	return s
}

func (s *ClickToDialRequest) SetRecordFlag(v bool) *ClickToDialRequest {
	s.RecordFlag = &v
	return s
}

func (s *ClickToDialRequest) SetAsrFlag(v bool) *ClickToDialRequest {
	s.AsrFlag = &v
	return s
}

func (s *ClickToDialRequest) SetSessionTimeout(v int32) *ClickToDialRequest {
	s.SessionTimeout = &v
	return s
}

func (s *ClickToDialRequest) SetAsrModelId(v string) *ClickToDialRequest {
	s.AsrModelId = &v
	return s
}

func (s *ClickToDialRequest) SetOutId(v string) *ClickToDialRequest {
	s.OutId = &v
	return s
}

type ClickToDialResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	CallId    *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s ClickToDialResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ClickToDialResponseBody) GoString() string {
	return s.String()
}

func (s *ClickToDialResponseBody) SetMessage(v string) *ClickToDialResponseBody {
	s.Message = &v
	return s
}

func (s *ClickToDialResponseBody) SetRequestId(v string) *ClickToDialResponseBody {
	s.RequestId = &v
	return s
}

func (s *ClickToDialResponseBody) SetCallId(v string) *ClickToDialResponseBody {
	s.CallId = &v
	return s
}

func (s *ClickToDialResponseBody) SetCode(v string) *ClickToDialResponseBody {
	s.Code = &v
	return s
}

type ClickToDialResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ClickToDialResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ClickToDialResponse) String() string {
	return tea.Prettify(s)
}

func (s ClickToDialResponse) GoString() string {
	return s.String()
}

func (s *ClickToDialResponse) SetHeaders(v map[string]*string) *ClickToDialResponse {
	s.Headers = v
	return s
}

func (s *ClickToDialResponse) SetBody(v *ClickToDialResponseBody) *ClickToDialResponse {
	s.Body = v
	return s
}

type CreateRobotTaskRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TaskName             *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	DialogId             *int64  `json:"DialogId,omitempty" xml:"DialogId,omitempty"`
	CorpName             *string `json:"CorpName,omitempty" xml:"CorpName,omitempty"`
	Caller               *string `json:"Caller,omitempty" xml:"Caller,omitempty"`
	NumberStatusIdent    *bool   `json:"NumberStatusIdent,omitempty" xml:"NumberStatusIdent,omitempty"`
	RetryType            *int32  `json:"RetryType,omitempty" xml:"RetryType,omitempty"`
	RecallStateCodes     *string `json:"RecallStateCodes,omitempty" xml:"RecallStateCodes,omitempty"`
	RecallTimes          *int32  `json:"RecallTimes,omitempty" xml:"RecallTimes,omitempty"`
	RecallInterval       *int32  `json:"RecallInterval,omitempty" xml:"RecallInterval,omitempty"`
	IsSelfLine           *bool   `json:"IsSelfLine,omitempty" xml:"IsSelfLine,omitempty"`
}

func (s CreateRobotTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRobotTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateRobotTaskRequest) SetOwnerId(v int64) *CreateRobotTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateRobotTaskRequest) SetResourceOwnerAccount(v string) *CreateRobotTaskRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateRobotTaskRequest) SetResourceOwnerId(v int64) *CreateRobotTaskRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateRobotTaskRequest) SetTaskName(v string) *CreateRobotTaskRequest {
	s.TaskName = &v
	return s
}

func (s *CreateRobotTaskRequest) SetDialogId(v int64) *CreateRobotTaskRequest {
	s.DialogId = &v
	return s
}

func (s *CreateRobotTaskRequest) SetCorpName(v string) *CreateRobotTaskRequest {
	s.CorpName = &v
	return s
}

func (s *CreateRobotTaskRequest) SetCaller(v string) *CreateRobotTaskRequest {
	s.Caller = &v
	return s
}

func (s *CreateRobotTaskRequest) SetNumberStatusIdent(v bool) *CreateRobotTaskRequest {
	s.NumberStatusIdent = &v
	return s
}

func (s *CreateRobotTaskRequest) SetRetryType(v int32) *CreateRobotTaskRequest {
	s.RetryType = &v
	return s
}

func (s *CreateRobotTaskRequest) SetRecallStateCodes(v string) *CreateRobotTaskRequest {
	s.RecallStateCodes = &v
	return s
}

func (s *CreateRobotTaskRequest) SetRecallTimes(v int32) *CreateRobotTaskRequest {
	s.RecallTimes = &v
	return s
}

func (s *CreateRobotTaskRequest) SetRecallInterval(v int32) *CreateRobotTaskRequest {
	s.RecallInterval = &v
	return s
}

func (s *CreateRobotTaskRequest) SetIsSelfLine(v bool) *CreateRobotTaskRequest {
	s.IsSelfLine = &v
	return s
}

type CreateRobotTaskResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s CreateRobotTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateRobotTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRobotTaskResponseBody) SetMessage(v string) *CreateRobotTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CreateRobotTaskResponseBody) SetRequestId(v string) *CreateRobotTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRobotTaskResponseBody) SetData(v string) *CreateRobotTaskResponseBody {
	s.Data = &v
	return s
}

func (s *CreateRobotTaskResponseBody) SetCode(v string) *CreateRobotTaskResponseBody {
	s.Code = &v
	return s
}

type CreateRobotTaskResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateRobotTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateRobotTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateRobotTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateRobotTaskResponse) SetHeaders(v map[string]*string) *CreateRobotTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateRobotTaskResponse) SetBody(v *CreateRobotTaskResponseBody) *CreateRobotTaskResponse {
	s.Body = v
	return s
}

type DeleteRobotTaskRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TaskId               *int64  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DeleteRobotTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteRobotTaskRequest) GoString() string {
	return s.String()
}

func (s *DeleteRobotTaskRequest) SetOwnerId(v int64) *DeleteRobotTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteRobotTaskRequest) SetResourceOwnerAccount(v string) *DeleteRobotTaskRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteRobotTaskRequest) SetResourceOwnerId(v int64) *DeleteRobotTaskRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteRobotTaskRequest) SetTaskId(v int64) *DeleteRobotTaskRequest {
	s.TaskId = &v
	return s
}

type DeleteRobotTaskResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s DeleteRobotTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteRobotTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRobotTaskResponseBody) SetMessage(v string) *DeleteRobotTaskResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteRobotTaskResponseBody) SetRequestId(v string) *DeleteRobotTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRobotTaskResponseBody) SetData(v string) *DeleteRobotTaskResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteRobotTaskResponseBody) SetCode(v string) *DeleteRobotTaskResponseBody {
	s.Code = &v
	return s
}

type DeleteRobotTaskResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteRobotTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteRobotTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteRobotTaskResponse) GoString() string {
	return s.String()
}

func (s *DeleteRobotTaskResponse) SetHeaders(v map[string]*string) *DeleteRobotTaskResponse {
	s.Headers = v
	return s
}

func (s *DeleteRobotTaskResponse) SetBody(v *DeleteRobotTaskResponseBody) *DeleteRobotTaskResponse {
	s.Body = v
	return s
}

type DescribeRecordDataRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	AccountType          *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	AccountId            *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	Acid                 *string `json:"Acid,omitempty" xml:"Acid,omitempty"`
	SecLevel             *int32  `json:"SecLevel,omitempty" xml:"SecLevel,omitempty"`
}

func (s DescribeRecordDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRecordDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeRecordDataRequest) SetOwnerId(v int64) *DescribeRecordDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeRecordDataRequest) SetResourceOwnerAccount(v string) *DescribeRecordDataRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeRecordDataRequest) SetResourceOwnerId(v int64) *DescribeRecordDataRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeRecordDataRequest) SetProdCode(v string) *DescribeRecordDataRequest {
	s.ProdCode = &v
	return s
}

func (s *DescribeRecordDataRequest) SetAccountType(v string) *DescribeRecordDataRequest {
	s.AccountType = &v
	return s
}

func (s *DescribeRecordDataRequest) SetAccountId(v string) *DescribeRecordDataRequest {
	s.AccountId = &v
	return s
}

func (s *DescribeRecordDataRequest) SetAcid(v string) *DescribeRecordDataRequest {
	s.Acid = &v
	return s
}

func (s *DescribeRecordDataRequest) SetSecLevel(v int32) *DescribeRecordDataRequest {
	s.SecLevel = &v
	return s
}

type DescribeRecordDataResponseBody struct {
	Acid      *string `json:"Acid,omitempty" xml:"Acid,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	OssLink   *string `json:"OssLink,omitempty" xml:"OssLink,omitempty"`
	AgentId   *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s DescribeRecordDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRecordDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRecordDataResponseBody) SetAcid(v string) *DescribeRecordDataResponseBody {
	s.Acid = &v
	return s
}

func (s *DescribeRecordDataResponseBody) SetRequestId(v string) *DescribeRecordDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRecordDataResponseBody) SetMessage(v string) *DescribeRecordDataResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeRecordDataResponseBody) SetOssLink(v string) *DescribeRecordDataResponseBody {
	s.OssLink = &v
	return s
}

func (s *DescribeRecordDataResponseBody) SetAgentId(v string) *DescribeRecordDataResponseBody {
	s.AgentId = &v
	return s
}

func (s *DescribeRecordDataResponseBody) SetCode(v string) *DescribeRecordDataResponseBody {
	s.Code = &v
	return s
}

type DescribeRecordDataResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeRecordDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRecordDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRecordDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeRecordDataResponse) SetHeaders(v map[string]*string) *DescribeRecordDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeRecordDataResponse) SetBody(v *DescribeRecordDataResponseBody) *DescribeRecordDataResponse {
	s.Body = v
	return s
}

type DoRtcNumberAuthRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	PhoneNumber          *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
}

func (s DoRtcNumberAuthRequest) String() string {
	return tea.Prettify(s)
}

func (s DoRtcNumberAuthRequest) GoString() string {
	return s.String()
}

func (s *DoRtcNumberAuthRequest) SetOwnerId(v int64) *DoRtcNumberAuthRequest {
	s.OwnerId = &v
	return s
}

func (s *DoRtcNumberAuthRequest) SetResourceOwnerAccount(v string) *DoRtcNumberAuthRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DoRtcNumberAuthRequest) SetResourceOwnerId(v int64) *DoRtcNumberAuthRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DoRtcNumberAuthRequest) SetPhoneNumber(v string) *DoRtcNumberAuthRequest {
	s.PhoneNumber = &v
	return s
}

type DoRtcNumberAuthResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Module    *string `json:"Module,omitempty" xml:"Module,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s DoRtcNumberAuthResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DoRtcNumberAuthResponseBody) GoString() string {
	return s.String()
}

func (s *DoRtcNumberAuthResponseBody) SetMessage(v string) *DoRtcNumberAuthResponseBody {
	s.Message = &v
	return s
}

func (s *DoRtcNumberAuthResponseBody) SetRequestId(v string) *DoRtcNumberAuthResponseBody {
	s.RequestId = &v
	return s
}

func (s *DoRtcNumberAuthResponseBody) SetModule(v string) *DoRtcNumberAuthResponseBody {
	s.Module = &v
	return s
}

func (s *DoRtcNumberAuthResponseBody) SetCode(v string) *DoRtcNumberAuthResponseBody {
	s.Code = &v
	return s
}

type DoRtcNumberAuthResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DoRtcNumberAuthResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DoRtcNumberAuthResponse) String() string {
	return tea.Prettify(s)
}

func (s DoRtcNumberAuthResponse) GoString() string {
	return s.String()
}

func (s *DoRtcNumberAuthResponse) SetHeaders(v map[string]*string) *DoRtcNumberAuthResponse {
	s.Headers = v
	return s
}

func (s *DoRtcNumberAuthResponse) SetBody(v *DoRtcNumberAuthResponseBody) *DoRtcNumberAuthResponse {
	s.Body = v
	return s
}

type GetRtcTokenRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	UserId               *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	DeviceId             *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	IsCustomAccount      *bool   `json:"IsCustomAccount,omitempty" xml:"IsCustomAccount,omitempty"`
}

func (s GetRtcTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRtcTokenRequest) GoString() string {
	return s.String()
}

func (s *GetRtcTokenRequest) SetOwnerId(v int64) *GetRtcTokenRequest {
	s.OwnerId = &v
	return s
}

func (s *GetRtcTokenRequest) SetResourceOwnerAccount(v string) *GetRtcTokenRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetRtcTokenRequest) SetResourceOwnerId(v int64) *GetRtcTokenRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetRtcTokenRequest) SetUserId(v string) *GetRtcTokenRequest {
	s.UserId = &v
	return s
}

func (s *GetRtcTokenRequest) SetDeviceId(v string) *GetRtcTokenRequest {
	s.DeviceId = &v
	return s
}

func (s *GetRtcTokenRequest) SetIsCustomAccount(v bool) *GetRtcTokenRequest {
	s.IsCustomAccount = &v
	return s
}

type GetRtcTokenResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Module    *string `json:"Module,omitempty" xml:"Module,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s GetRtcTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRtcTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GetRtcTokenResponseBody) SetMessage(v string) *GetRtcTokenResponseBody {
	s.Message = &v
	return s
}

func (s *GetRtcTokenResponseBody) SetRequestId(v string) *GetRtcTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRtcTokenResponseBody) SetModule(v string) *GetRtcTokenResponseBody {
	s.Module = &v
	return s
}

func (s *GetRtcTokenResponseBody) SetCode(v string) *GetRtcTokenResponseBody {
	s.Code = &v
	return s
}

type GetRtcTokenResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetRtcTokenResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetRtcTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRtcTokenResponse) GoString() string {
	return s.String()
}

func (s *GetRtcTokenResponse) SetHeaders(v map[string]*string) *GetRtcTokenResponse {
	s.Headers = v
	return s
}

func (s *GetRtcTokenResponse) SetBody(v *GetRtcTokenResponseBody) *GetRtcTokenResponse {
	s.Body = v
	return s
}

type IvrCallRequest struct {
	OwnerId              *int64                      `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string                     `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64                      `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	CalledShowNumber     *string                     `json:"CalledShowNumber,omitempty" xml:"CalledShowNumber,omitempty"`
	CalledNumber         *string                     `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	StartCode            *string                     `json:"StartCode,omitempty" xml:"StartCode,omitempty"`
	StartTtsParams       *string                     `json:"StartTtsParams,omitempty" xml:"StartTtsParams,omitempty"`
	PlayTimes            *int64                      `json:"PlayTimes,omitempty" xml:"PlayTimes,omitempty"`
	ByeCode              *string                     `json:"ByeCode,omitempty" xml:"ByeCode,omitempty"`
	ByeTtsParams         *string                     `json:"ByeTtsParams,omitempty" xml:"ByeTtsParams,omitempty"`
	Timeout              *int32                      `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	OutId                *string                     `json:"OutId,omitempty" xml:"OutId,omitempty"`
	MenuKeyMap           []*IvrCallRequestMenuKeyMap `json:"MenuKeyMap,omitempty" xml:"MenuKeyMap,omitempty" type:"Repeated"`
}

func (s IvrCallRequest) String() string {
	return tea.Prettify(s)
}

func (s IvrCallRequest) GoString() string {
	return s.String()
}

func (s *IvrCallRequest) SetOwnerId(v int64) *IvrCallRequest {
	s.OwnerId = &v
	return s
}

func (s *IvrCallRequest) SetResourceOwnerAccount(v string) *IvrCallRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *IvrCallRequest) SetResourceOwnerId(v int64) *IvrCallRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *IvrCallRequest) SetCalledShowNumber(v string) *IvrCallRequest {
	s.CalledShowNumber = &v
	return s
}

func (s *IvrCallRequest) SetCalledNumber(v string) *IvrCallRequest {
	s.CalledNumber = &v
	return s
}

func (s *IvrCallRequest) SetStartCode(v string) *IvrCallRequest {
	s.StartCode = &v
	return s
}

func (s *IvrCallRequest) SetStartTtsParams(v string) *IvrCallRequest {
	s.StartTtsParams = &v
	return s
}

func (s *IvrCallRequest) SetPlayTimes(v int64) *IvrCallRequest {
	s.PlayTimes = &v
	return s
}

func (s *IvrCallRequest) SetByeCode(v string) *IvrCallRequest {
	s.ByeCode = &v
	return s
}

func (s *IvrCallRequest) SetByeTtsParams(v string) *IvrCallRequest {
	s.ByeTtsParams = &v
	return s
}

func (s *IvrCallRequest) SetTimeout(v int32) *IvrCallRequest {
	s.Timeout = &v
	return s
}

func (s *IvrCallRequest) SetOutId(v string) *IvrCallRequest {
	s.OutId = &v
	return s
}

func (s *IvrCallRequest) SetMenuKeyMap(v []*IvrCallRequestMenuKeyMap) *IvrCallRequest {
	s.MenuKeyMap = v
	return s
}

type IvrCallRequestMenuKeyMap struct {
	Key       *string `json:"Key,omitempty" xml:"Key,omitempty"`
	TtsParams *string `json:"TtsParams,omitempty" xml:"TtsParams,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s IvrCallRequestMenuKeyMap) String() string {
	return tea.Prettify(s)
}

func (s IvrCallRequestMenuKeyMap) GoString() string {
	return s.String()
}

func (s *IvrCallRequestMenuKeyMap) SetKey(v string) *IvrCallRequestMenuKeyMap {
	s.Key = &v
	return s
}

func (s *IvrCallRequestMenuKeyMap) SetTtsParams(v string) *IvrCallRequestMenuKeyMap {
	s.TtsParams = &v
	return s
}

func (s *IvrCallRequestMenuKeyMap) SetCode(v string) *IvrCallRequestMenuKeyMap {
	s.Code = &v
	return s
}

type IvrCallResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	CallId    *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s IvrCallResponseBody) String() string {
	return tea.Prettify(s)
}

func (s IvrCallResponseBody) GoString() string {
	return s.String()
}

func (s *IvrCallResponseBody) SetMessage(v string) *IvrCallResponseBody {
	s.Message = &v
	return s
}

func (s *IvrCallResponseBody) SetRequestId(v string) *IvrCallResponseBody {
	s.RequestId = &v
	return s
}

func (s *IvrCallResponseBody) SetCallId(v string) *IvrCallResponseBody {
	s.CallId = &v
	return s
}

func (s *IvrCallResponseBody) SetCode(v string) *IvrCallResponseBody {
	s.Code = &v
	return s
}

type IvrCallResponse struct {
	Headers map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *IvrCallResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s IvrCallResponse) String() string {
	return tea.Prettify(s)
}

func (s IvrCallResponse) GoString() string {
	return s.String()
}

func (s *IvrCallResponse) SetHeaders(v map[string]*string) *IvrCallResponse {
	s.Headers = v
	return s
}

func (s *IvrCallResponse) SetBody(v *IvrCallResponseBody) *IvrCallResponse {
	s.Body = v
	return s
}

type ListOrderedNumbersRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
}

func (s ListOrderedNumbersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListOrderedNumbersRequest) GoString() string {
	return s.String()
}

func (s *ListOrderedNumbersRequest) SetOwnerId(v int64) *ListOrderedNumbersRequest {
	s.OwnerId = &v
	return s
}

func (s *ListOrderedNumbersRequest) SetResourceOwnerAccount(v string) *ListOrderedNumbersRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListOrderedNumbersRequest) SetResourceOwnerId(v int64) *ListOrderedNumbersRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListOrderedNumbersRequest) SetProdCode(v string) *ListOrderedNumbersRequest {
	s.ProdCode = &v
	return s
}

type ListOrderedNumbersResponseBody struct {
	Numbers   []*string `json:"Numbers,omitempty" xml:"Numbers,omitempty" type:"Repeated"`
	Message   *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string   `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s ListOrderedNumbersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListOrderedNumbersResponseBody) GoString() string {
	return s.String()
}

func (s *ListOrderedNumbersResponseBody) SetNumbers(v []*string) *ListOrderedNumbersResponseBody {
	s.Numbers = v
	return s
}

func (s *ListOrderedNumbersResponseBody) SetMessage(v string) *ListOrderedNumbersResponseBody {
	s.Message = &v
	return s
}

func (s *ListOrderedNumbersResponseBody) SetRequestId(v string) *ListOrderedNumbersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOrderedNumbersResponseBody) SetCode(v string) *ListOrderedNumbersResponseBody {
	s.Code = &v
	return s
}

type ListOrderedNumbersResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListOrderedNumbersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListOrderedNumbersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListOrderedNumbersResponse) GoString() string {
	return s.String()
}

func (s *ListOrderedNumbersResponse) SetHeaders(v map[string]*string) *ListOrderedNumbersResponse {
	s.Headers = v
	return s
}

func (s *ListOrderedNumbersResponse) SetBody(v *ListOrderedNumbersResponseBody) *ListOrderedNumbersResponse {
	s.Body = v
	return s
}

type ListOutboundStrategiesRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	BuId                 *int64  `json:"BuId,omitempty" xml:"BuId,omitempty"`
	Keyword              *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
}

func (s ListOutboundStrategiesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListOutboundStrategiesRequest) GoString() string {
	return s.String()
}

func (s *ListOutboundStrategiesRequest) SetOwnerId(v int64) *ListOutboundStrategiesRequest {
	s.OwnerId = &v
	return s
}

func (s *ListOutboundStrategiesRequest) SetResourceOwnerAccount(v string) *ListOutboundStrategiesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListOutboundStrategiesRequest) SetResourceOwnerId(v int64) *ListOutboundStrategiesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListOutboundStrategiesRequest) SetProdCode(v string) *ListOutboundStrategiesRequest {
	s.ProdCode = &v
	return s
}

func (s *ListOutboundStrategiesRequest) SetBuId(v int64) *ListOutboundStrategiesRequest {
	s.BuId = &v
	return s
}

func (s *ListOutboundStrategiesRequest) SetKeyword(v string) *ListOutboundStrategiesRequest {
	s.Keyword = &v
	return s
}

type ListOutboundStrategiesResponseBody struct {
	Message            *string                                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId          *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	OutboundStrategies []*ListOutboundStrategiesResponseBodyOutboundStrategies `json:"OutboundStrategies,omitempty" xml:"OutboundStrategies,omitempty" type:"Repeated"`
	Code               *string                                                 `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s ListOutboundStrategiesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListOutboundStrategiesResponseBody) GoString() string {
	return s.String()
}

func (s *ListOutboundStrategiesResponseBody) SetMessage(v string) *ListOutboundStrategiesResponseBody {
	s.Message = &v
	return s
}

func (s *ListOutboundStrategiesResponseBody) SetRequestId(v string) *ListOutboundStrategiesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOutboundStrategiesResponseBody) SetOutboundStrategies(v []*ListOutboundStrategiesResponseBodyOutboundStrategies) *ListOutboundStrategiesResponseBody {
	s.OutboundStrategies = v
	return s
}

func (s *ListOutboundStrategiesResponseBody) SetCode(v string) *ListOutboundStrategiesResponseBody {
	s.Code = &v
	return s
}

type ListOutboundStrategiesResponseBodyOutboundStrategies struct {
	Status             *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	SuccessRate        *int32  `json:"SuccessRate,omitempty" xml:"SuccessRate,omitempty"`
	Process            *int32  `json:"Process,omitempty" xml:"Process,omitempty"`
	GmtModifiedStr     *string `json:"GmtModifiedStr,omitempty" xml:"GmtModifiedStr,omitempty"`
	OutboundNum        *string `json:"OutboundNum,omitempty" xml:"OutboundNum,omitempty"`
	ModifierId         *int64  `json:"ModifierId,omitempty" xml:"ModifierId,omitempty"`
	SceneName          *string `json:"SceneName,omitempty" xml:"SceneName,omitempty"`
	CreatorId          *int64  `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	RobotName          *string `json:"RobotName,omitempty" xml:"RobotName,omitempty"`
	ModifierName       *string `json:"ModifierName,omitempty" xml:"ModifierName,omitempty"`
	ResourceAllocation *int32  `json:"ResourceAllocation,omitempty" xml:"ResourceAllocation,omitempty"`
	ExtAttr            *string `json:"ExtAttr,omitempty" xml:"ExtAttr,omitempty"`
	NumType            *int32  `json:"NumType,omitempty" xml:"NumType,omitempty"`
	BuId               *int64  `json:"BuId,omitempty" xml:"BuId,omitempty"`
	RobotId            *string `json:"RobotId,omitempty" xml:"RobotId,omitempty"`
	CreatorName        *string `json:"CreatorName,omitempty" xml:"CreatorName,omitempty"`
	DepartmentId       *int64  `json:"DepartmentId,omitempty" xml:"DepartmentId,omitempty"`
	RobotType          *int32  `json:"RobotType,omitempty" xml:"RobotType,omitempty"`
	RuleCode           *int64  `json:"RuleCode,omitempty" xml:"RuleCode,omitempty"`
	Name               *string `json:"name,omitempty" xml:"name,omitempty"`
	GmtCreateStr       *string `json:"GmtCreateStr,omitempty" xml:"GmtCreateStr,omitempty"`
	Id                 *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ListOutboundStrategiesResponseBodyOutboundStrategies) String() string {
	return tea.Prettify(s)
}

func (s ListOutboundStrategiesResponseBodyOutboundStrategies) GoString() string {
	return s.String()
}

func (s *ListOutboundStrategiesResponseBodyOutboundStrategies) SetStatus(v int32) *ListOutboundStrategiesResponseBodyOutboundStrategies {
	s.Status = &v
	return s
}

func (s *ListOutboundStrategiesResponseBodyOutboundStrategies) SetSuccessRate(v int32) *ListOutboundStrategiesResponseBodyOutboundStrategies {
	s.SuccessRate = &v
	return s
}

func (s *ListOutboundStrategiesResponseBodyOutboundStrategies) SetProcess(v int32) *ListOutboundStrategiesResponseBodyOutboundStrategies {
	s.Process = &v
	return s
}

func (s *ListOutboundStrategiesResponseBodyOutboundStrategies) SetGmtModifiedStr(v string) *ListOutboundStrategiesResponseBodyOutboundStrategies {
	s.GmtModifiedStr = &v
	return s
}

func (s *ListOutboundStrategiesResponseBodyOutboundStrategies) SetOutboundNum(v string) *ListOutboundStrategiesResponseBodyOutboundStrategies {
	s.OutboundNum = &v
	return s
}

func (s *ListOutboundStrategiesResponseBodyOutboundStrategies) SetModifierId(v int64) *ListOutboundStrategiesResponseBodyOutboundStrategies {
	s.ModifierId = &v
	return s
}

func (s *ListOutboundStrategiesResponseBodyOutboundStrategies) SetSceneName(v string) *ListOutboundStrategiesResponseBodyOutboundStrategies {
	s.SceneName = &v
	return s
}

func (s *ListOutboundStrategiesResponseBodyOutboundStrategies) SetCreatorId(v int64) *ListOutboundStrategiesResponseBodyOutboundStrategies {
	s.CreatorId = &v
	return s
}

func (s *ListOutboundStrategiesResponseBodyOutboundStrategies) SetRobotName(v string) *ListOutboundStrategiesResponseBodyOutboundStrategies {
	s.RobotName = &v
	return s
}

func (s *ListOutboundStrategiesResponseBodyOutboundStrategies) SetModifierName(v string) *ListOutboundStrategiesResponseBodyOutboundStrategies {
	s.ModifierName = &v
	return s
}

func (s *ListOutboundStrategiesResponseBodyOutboundStrategies) SetResourceAllocation(v int32) *ListOutboundStrategiesResponseBodyOutboundStrategies {
	s.ResourceAllocation = &v
	return s
}

func (s *ListOutboundStrategiesResponseBodyOutboundStrategies) SetExtAttr(v string) *ListOutboundStrategiesResponseBodyOutboundStrategies {
	s.ExtAttr = &v
	return s
}

func (s *ListOutboundStrategiesResponseBodyOutboundStrategies) SetNumType(v int32) *ListOutboundStrategiesResponseBodyOutboundStrategies {
	s.NumType = &v
	return s
}

func (s *ListOutboundStrategiesResponseBodyOutboundStrategies) SetBuId(v int64) *ListOutboundStrategiesResponseBodyOutboundStrategies {
	s.BuId = &v
	return s
}

func (s *ListOutboundStrategiesResponseBodyOutboundStrategies) SetRobotId(v string) *ListOutboundStrategiesResponseBodyOutboundStrategies {
	s.RobotId = &v
	return s
}

func (s *ListOutboundStrategiesResponseBodyOutboundStrategies) SetCreatorName(v string) *ListOutboundStrategiesResponseBodyOutboundStrategies {
	s.CreatorName = &v
	return s
}

func (s *ListOutboundStrategiesResponseBodyOutboundStrategies) SetDepartmentId(v int64) *ListOutboundStrategiesResponseBodyOutboundStrategies {
	s.DepartmentId = &v
	return s
}

func (s *ListOutboundStrategiesResponseBodyOutboundStrategies) SetRobotType(v int32) *ListOutboundStrategiesResponseBodyOutboundStrategies {
	s.RobotType = &v
	return s
}

func (s *ListOutboundStrategiesResponseBodyOutboundStrategies) SetRuleCode(v int64) *ListOutboundStrategiesResponseBodyOutboundStrategies {
	s.RuleCode = &v
	return s
}

func (s *ListOutboundStrategiesResponseBodyOutboundStrategies) SetName(v string) *ListOutboundStrategiesResponseBodyOutboundStrategies {
	s.Name = &v
	return s
}

func (s *ListOutboundStrategiesResponseBodyOutboundStrategies) SetGmtCreateStr(v string) *ListOutboundStrategiesResponseBodyOutboundStrategies {
	s.GmtCreateStr = &v
	return s
}

func (s *ListOutboundStrategiesResponseBodyOutboundStrategies) SetId(v int64) *ListOutboundStrategiesResponseBodyOutboundStrategies {
	s.Id = &v
	return s
}

type ListOutboundStrategiesResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListOutboundStrategiesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListOutboundStrategiesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListOutboundStrategiesResponse) GoString() string {
	return s.String()
}

func (s *ListOutboundStrategiesResponse) SetHeaders(v map[string]*string) *ListOutboundStrategiesResponse {
	s.Headers = v
	return s
}

func (s *ListOutboundStrategiesResponse) SetBody(v *ListOutboundStrategiesResponseBody) *ListOutboundStrategiesResponse {
	s.Body = v
	return s
}

type ListRobotTaskCallsRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	PageNo               *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TaskId               *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	DurationFrom         *string `json:"DurationFrom,omitempty" xml:"DurationFrom,omitempty"`
	DurationTo           *string `json:"DurationTo,omitempty" xml:"DurationTo,omitempty"`
	DialogCountFrom      *string `json:"DialogCountFrom,omitempty" xml:"DialogCountFrom,omitempty"`
	DialogCountTo        *string `json:"DialogCountTo,omitempty" xml:"DialogCountTo,omitempty"`
	HangupDirection      *string `json:"HangupDirection,omitempty" xml:"HangupDirection,omitempty"`
	CallResult           *string `json:"CallResult,omitempty" xml:"CallResult,omitempty"`
	Called               *string `json:"Called,omitempty" xml:"Called,omitempty"`
}

func (s ListRobotTaskCallsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRobotTaskCallsRequest) GoString() string {
	return s.String()
}

func (s *ListRobotTaskCallsRequest) SetOwnerId(v int64) *ListRobotTaskCallsRequest {
	s.OwnerId = &v
	return s
}

func (s *ListRobotTaskCallsRequest) SetResourceOwnerAccount(v string) *ListRobotTaskCallsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListRobotTaskCallsRequest) SetResourceOwnerId(v int64) *ListRobotTaskCallsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListRobotTaskCallsRequest) SetPageNo(v int32) *ListRobotTaskCallsRequest {
	s.PageNo = &v
	return s
}

func (s *ListRobotTaskCallsRequest) SetPageSize(v int32) *ListRobotTaskCallsRequest {
	s.PageSize = &v
	return s
}

func (s *ListRobotTaskCallsRequest) SetTaskId(v string) *ListRobotTaskCallsRequest {
	s.TaskId = &v
	return s
}

func (s *ListRobotTaskCallsRequest) SetDurationFrom(v string) *ListRobotTaskCallsRequest {
	s.DurationFrom = &v
	return s
}

func (s *ListRobotTaskCallsRequest) SetDurationTo(v string) *ListRobotTaskCallsRequest {
	s.DurationTo = &v
	return s
}

func (s *ListRobotTaskCallsRequest) SetDialogCountFrom(v string) *ListRobotTaskCallsRequest {
	s.DialogCountFrom = &v
	return s
}

func (s *ListRobotTaskCallsRequest) SetDialogCountTo(v string) *ListRobotTaskCallsRequest {
	s.DialogCountTo = &v
	return s
}

func (s *ListRobotTaskCallsRequest) SetHangupDirection(v string) *ListRobotTaskCallsRequest {
	s.HangupDirection = &v
	return s
}

func (s *ListRobotTaskCallsRequest) SetCallResult(v string) *ListRobotTaskCallsRequest {
	s.CallResult = &v
	return s
}

func (s *ListRobotTaskCallsRequest) SetCalled(v string) *ListRobotTaskCallsRequest {
	s.Called = &v
	return s
}

type ListRobotTaskCallsResponseBody struct {
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Message    *string `json:"Message,omitempty" xml:"Message,omitempty"`
	PageSize   *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Data       *string `json:"Data,omitempty" xml:"Data,omitempty"`
	PageNo     *string `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	Code       *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s ListRobotTaskCallsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRobotTaskCallsResponseBody) GoString() string {
	return s.String()
}

func (s *ListRobotTaskCallsResponseBody) SetTotalCount(v string) *ListRobotTaskCallsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListRobotTaskCallsResponseBody) SetRequestId(v string) *ListRobotTaskCallsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRobotTaskCallsResponseBody) SetMessage(v string) *ListRobotTaskCallsResponseBody {
	s.Message = &v
	return s
}

func (s *ListRobotTaskCallsResponseBody) SetPageSize(v string) *ListRobotTaskCallsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListRobotTaskCallsResponseBody) SetData(v string) *ListRobotTaskCallsResponseBody {
	s.Data = &v
	return s
}

func (s *ListRobotTaskCallsResponseBody) SetPageNo(v string) *ListRobotTaskCallsResponseBody {
	s.PageNo = &v
	return s
}

func (s *ListRobotTaskCallsResponseBody) SetCode(v string) *ListRobotTaskCallsResponseBody {
	s.Code = &v
	return s
}

type ListRobotTaskCallsResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListRobotTaskCallsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListRobotTaskCallsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRobotTaskCallsResponse) GoString() string {
	return s.String()
}

func (s *ListRobotTaskCallsResponse) SetHeaders(v map[string]*string) *ListRobotTaskCallsResponse {
	s.Headers = v
	return s
}

func (s *ListRobotTaskCallsResponse) SetBody(v *ListRobotTaskCallsResponseBody) *ListRobotTaskCallsResponse {
	s.Body = v
	return s
}

type QueryCallDetailByCallIdRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	CallId               *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	ProdId               *int64  `json:"ProdId,omitempty" xml:"ProdId,omitempty"`
	QueryDate            *int64  `json:"QueryDate,omitempty" xml:"QueryDate,omitempty"`
}

func (s QueryCallDetailByCallIdRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryCallDetailByCallIdRequest) GoString() string {
	return s.String()
}

func (s *QueryCallDetailByCallIdRequest) SetOwnerId(v int64) *QueryCallDetailByCallIdRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryCallDetailByCallIdRequest) SetResourceOwnerAccount(v string) *QueryCallDetailByCallIdRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryCallDetailByCallIdRequest) SetResourceOwnerId(v int64) *QueryCallDetailByCallIdRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryCallDetailByCallIdRequest) SetCallId(v string) *QueryCallDetailByCallIdRequest {
	s.CallId = &v
	return s
}

func (s *QueryCallDetailByCallIdRequest) SetProdId(v int64) *QueryCallDetailByCallIdRequest {
	s.ProdId = &v
	return s
}

func (s *QueryCallDetailByCallIdRequest) SetQueryDate(v int64) *QueryCallDetailByCallIdRequest {
	s.QueryDate = &v
	return s
}

type QueryCallDetailByCallIdResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s QueryCallDetailByCallIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryCallDetailByCallIdResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCallDetailByCallIdResponseBody) SetMessage(v string) *QueryCallDetailByCallIdResponseBody {
	s.Message = &v
	return s
}

func (s *QueryCallDetailByCallIdResponseBody) SetRequestId(v string) *QueryCallDetailByCallIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryCallDetailByCallIdResponseBody) SetData(v string) *QueryCallDetailByCallIdResponseBody {
	s.Data = &v
	return s
}

func (s *QueryCallDetailByCallIdResponseBody) SetCode(v string) *QueryCallDetailByCallIdResponseBody {
	s.Code = &v
	return s
}

type QueryCallDetailByCallIdResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryCallDetailByCallIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryCallDetailByCallIdResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryCallDetailByCallIdResponse) GoString() string {
	return s.String()
}

func (s *QueryCallDetailByCallIdResponse) SetHeaders(v map[string]*string) *QueryCallDetailByCallIdResponse {
	s.Headers = v
	return s
}

func (s *QueryCallDetailByCallIdResponse) SetBody(v *QueryCallDetailByCallIdResponseBody) *QueryCallDetailByCallIdResponse {
	s.Body = v
	return s
}

type QueryCallDetailByTaskIdRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TaskId               *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	QueryDate            *int64  `json:"QueryDate,omitempty" xml:"QueryDate,omitempty"`
	Callee               *string `json:"Callee,omitempty" xml:"Callee,omitempty"`
}

func (s QueryCallDetailByTaskIdRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryCallDetailByTaskIdRequest) GoString() string {
	return s.String()
}

func (s *QueryCallDetailByTaskIdRequest) SetOwnerId(v int64) *QueryCallDetailByTaskIdRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryCallDetailByTaskIdRequest) SetResourceOwnerAccount(v string) *QueryCallDetailByTaskIdRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryCallDetailByTaskIdRequest) SetResourceOwnerId(v int64) *QueryCallDetailByTaskIdRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryCallDetailByTaskIdRequest) SetTaskId(v string) *QueryCallDetailByTaskIdRequest {
	s.TaskId = &v
	return s
}

func (s *QueryCallDetailByTaskIdRequest) SetQueryDate(v int64) *QueryCallDetailByTaskIdRequest {
	s.QueryDate = &v
	return s
}

func (s *QueryCallDetailByTaskIdRequest) SetCallee(v string) *QueryCallDetailByTaskIdRequest {
	s.Callee = &v
	return s
}

type QueryCallDetailByTaskIdResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s QueryCallDetailByTaskIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryCallDetailByTaskIdResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCallDetailByTaskIdResponseBody) SetMessage(v string) *QueryCallDetailByTaskIdResponseBody {
	s.Message = &v
	return s
}

func (s *QueryCallDetailByTaskIdResponseBody) SetRequestId(v string) *QueryCallDetailByTaskIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryCallDetailByTaskIdResponseBody) SetData(v string) *QueryCallDetailByTaskIdResponseBody {
	s.Data = &v
	return s
}

func (s *QueryCallDetailByTaskIdResponseBody) SetCode(v string) *QueryCallDetailByTaskIdResponseBody {
	s.Code = &v
	return s
}

type QueryCallDetailByTaskIdResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryCallDetailByTaskIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryCallDetailByTaskIdResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryCallDetailByTaskIdResponse) GoString() string {
	return s.String()
}

func (s *QueryCallDetailByTaskIdResponse) SetHeaders(v map[string]*string) *QueryCallDetailByTaskIdResponse {
	s.Headers = v
	return s
}

func (s *QueryCallDetailByTaskIdResponse) SetBody(v *QueryCallDetailByTaskIdResponseBody) *QueryCallDetailByTaskIdResponse {
	s.Body = v
	return s
}

type QueryRobotInfoListRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	AuditStatus          *string `json:"AuditStatus,omitempty" xml:"AuditStatus,omitempty"`
}

func (s QueryRobotInfoListRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryRobotInfoListRequest) GoString() string {
	return s.String()
}

func (s *QueryRobotInfoListRequest) SetOwnerId(v int64) *QueryRobotInfoListRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryRobotInfoListRequest) SetResourceOwnerAccount(v string) *QueryRobotInfoListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryRobotInfoListRequest) SetResourceOwnerId(v int64) *QueryRobotInfoListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryRobotInfoListRequest) SetAuditStatus(v string) *QueryRobotInfoListRequest {
	s.AuditStatus = &v
	return s
}

type QueryRobotInfoListResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s QueryRobotInfoListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryRobotInfoListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryRobotInfoListResponseBody) SetMessage(v string) *QueryRobotInfoListResponseBody {
	s.Message = &v
	return s
}

func (s *QueryRobotInfoListResponseBody) SetRequestId(v string) *QueryRobotInfoListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryRobotInfoListResponseBody) SetData(v string) *QueryRobotInfoListResponseBody {
	s.Data = &v
	return s
}

func (s *QueryRobotInfoListResponseBody) SetCode(v string) *QueryRobotInfoListResponseBody {
	s.Code = &v
	return s
}

type QueryRobotInfoListResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryRobotInfoListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryRobotInfoListResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryRobotInfoListResponse) GoString() string {
	return s.String()
}

func (s *QueryRobotInfoListResponse) SetHeaders(v map[string]*string) *QueryRobotInfoListResponse {
	s.Headers = v
	return s
}

func (s *QueryRobotInfoListResponse) SetBody(v *QueryRobotInfoListResponseBody) *QueryRobotInfoListResponse {
	s.Body = v
	return s
}

type QueryRobotTaskCallDetailRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TaskId               *int64  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	Callee               *string `json:"Callee,omitempty" xml:"Callee,omitempty"`
	QueryDate            *int64  `json:"QueryDate,omitempty" xml:"QueryDate,omitempty"`
}

func (s QueryRobotTaskCallDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryRobotTaskCallDetailRequest) GoString() string {
	return s.String()
}

func (s *QueryRobotTaskCallDetailRequest) SetOwnerId(v int64) *QueryRobotTaskCallDetailRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryRobotTaskCallDetailRequest) SetResourceOwnerAccount(v string) *QueryRobotTaskCallDetailRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryRobotTaskCallDetailRequest) SetResourceOwnerId(v int64) *QueryRobotTaskCallDetailRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryRobotTaskCallDetailRequest) SetTaskId(v int64) *QueryRobotTaskCallDetailRequest {
	s.TaskId = &v
	return s
}

func (s *QueryRobotTaskCallDetailRequest) SetCallee(v string) *QueryRobotTaskCallDetailRequest {
	s.Callee = &v
	return s
}

func (s *QueryRobotTaskCallDetailRequest) SetQueryDate(v int64) *QueryRobotTaskCallDetailRequest {
	s.QueryDate = &v
	return s
}

type QueryRobotTaskCallDetailResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s QueryRobotTaskCallDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryRobotTaskCallDetailResponseBody) GoString() string {
	return s.String()
}

func (s *QueryRobotTaskCallDetailResponseBody) SetMessage(v string) *QueryRobotTaskCallDetailResponseBody {
	s.Message = &v
	return s
}

func (s *QueryRobotTaskCallDetailResponseBody) SetRequestId(v string) *QueryRobotTaskCallDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryRobotTaskCallDetailResponseBody) SetData(v string) *QueryRobotTaskCallDetailResponseBody {
	s.Data = &v
	return s
}

func (s *QueryRobotTaskCallDetailResponseBody) SetCode(v string) *QueryRobotTaskCallDetailResponseBody {
	s.Code = &v
	return s
}

type QueryRobotTaskCallDetailResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryRobotTaskCallDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryRobotTaskCallDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryRobotTaskCallDetailResponse) GoString() string {
	return s.String()
}

func (s *QueryRobotTaskCallDetailResponse) SetHeaders(v map[string]*string) *QueryRobotTaskCallDetailResponse {
	s.Headers = v
	return s
}

func (s *QueryRobotTaskCallDetailResponse) SetBody(v *QueryRobotTaskCallDetailResponseBody) *QueryRobotTaskCallDetailResponse {
	s.Body = v
	return s
}

type QueryRobotTaskCallListRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	PageNo               *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TaskId               *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	DurationFrom         *string `json:"DurationFrom,omitempty" xml:"DurationFrom,omitempty"`
	DurationTo           *string `json:"DurationTo,omitempty" xml:"DurationTo,omitempty"`
	DialogCountFrom      *string `json:"DialogCountFrom,omitempty" xml:"DialogCountFrom,omitempty"`
	DialogCountTo        *string `json:"DialogCountTo,omitempty" xml:"DialogCountTo,omitempty"`
	HangupDirection      *string `json:"HangupDirection,omitempty" xml:"HangupDirection,omitempty"`
	CallResult           *string `json:"CallResult,omitempty" xml:"CallResult,omitempty"`
	Called               *string `json:"Called,omitempty" xml:"Called,omitempty"`
}

func (s QueryRobotTaskCallListRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryRobotTaskCallListRequest) GoString() string {
	return s.String()
}

func (s *QueryRobotTaskCallListRequest) SetOwnerId(v int64) *QueryRobotTaskCallListRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryRobotTaskCallListRequest) SetResourceOwnerAccount(v string) *QueryRobotTaskCallListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryRobotTaskCallListRequest) SetResourceOwnerId(v int64) *QueryRobotTaskCallListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryRobotTaskCallListRequest) SetPageNo(v int32) *QueryRobotTaskCallListRequest {
	s.PageNo = &v
	return s
}

func (s *QueryRobotTaskCallListRequest) SetPageSize(v int32) *QueryRobotTaskCallListRequest {
	s.PageSize = &v
	return s
}

func (s *QueryRobotTaskCallListRequest) SetTaskId(v string) *QueryRobotTaskCallListRequest {
	s.TaskId = &v
	return s
}

func (s *QueryRobotTaskCallListRequest) SetDurationFrom(v string) *QueryRobotTaskCallListRequest {
	s.DurationFrom = &v
	return s
}

func (s *QueryRobotTaskCallListRequest) SetDurationTo(v string) *QueryRobotTaskCallListRequest {
	s.DurationTo = &v
	return s
}

func (s *QueryRobotTaskCallListRequest) SetDialogCountFrom(v string) *QueryRobotTaskCallListRequest {
	s.DialogCountFrom = &v
	return s
}

func (s *QueryRobotTaskCallListRequest) SetDialogCountTo(v string) *QueryRobotTaskCallListRequest {
	s.DialogCountTo = &v
	return s
}

func (s *QueryRobotTaskCallListRequest) SetHangupDirection(v string) *QueryRobotTaskCallListRequest {
	s.HangupDirection = &v
	return s
}

func (s *QueryRobotTaskCallListRequest) SetCallResult(v string) *QueryRobotTaskCallListRequest {
	s.CallResult = &v
	return s
}

func (s *QueryRobotTaskCallListRequest) SetCalled(v string) *QueryRobotTaskCallListRequest {
	s.Called = &v
	return s
}

type QueryRobotTaskCallListResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s QueryRobotTaskCallListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryRobotTaskCallListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryRobotTaskCallListResponseBody) SetMessage(v string) *QueryRobotTaskCallListResponseBody {
	s.Message = &v
	return s
}

func (s *QueryRobotTaskCallListResponseBody) SetRequestId(v string) *QueryRobotTaskCallListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryRobotTaskCallListResponseBody) SetData(v string) *QueryRobotTaskCallListResponseBody {
	s.Data = &v
	return s
}

func (s *QueryRobotTaskCallListResponseBody) SetCode(v string) *QueryRobotTaskCallListResponseBody {
	s.Code = &v
	return s
}

type QueryRobotTaskCallListResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryRobotTaskCallListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryRobotTaskCallListResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryRobotTaskCallListResponse) GoString() string {
	return s.String()
}

func (s *QueryRobotTaskCallListResponse) SetHeaders(v map[string]*string) *QueryRobotTaskCallListResponse {
	s.Headers = v
	return s
}

func (s *QueryRobotTaskCallListResponse) SetBody(v *QueryRobotTaskCallListResponseBody) *QueryRobotTaskCallListResponse {
	s.Body = v
	return s
}

type QueryRobotTaskDetailRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Id                   *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s QueryRobotTaskDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryRobotTaskDetailRequest) GoString() string {
	return s.String()
}

func (s *QueryRobotTaskDetailRequest) SetOwnerId(v int64) *QueryRobotTaskDetailRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryRobotTaskDetailRequest) SetResourceOwnerAccount(v string) *QueryRobotTaskDetailRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryRobotTaskDetailRequest) SetResourceOwnerId(v int64) *QueryRobotTaskDetailRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryRobotTaskDetailRequest) SetId(v int64) *QueryRobotTaskDetailRequest {
	s.Id = &v
	return s
}

type QueryRobotTaskDetailResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s QueryRobotTaskDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryRobotTaskDetailResponseBody) GoString() string {
	return s.String()
}

func (s *QueryRobotTaskDetailResponseBody) SetMessage(v string) *QueryRobotTaskDetailResponseBody {
	s.Message = &v
	return s
}

func (s *QueryRobotTaskDetailResponseBody) SetRequestId(v string) *QueryRobotTaskDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryRobotTaskDetailResponseBody) SetData(v string) *QueryRobotTaskDetailResponseBody {
	s.Data = &v
	return s
}

func (s *QueryRobotTaskDetailResponseBody) SetCode(v string) *QueryRobotTaskDetailResponseBody {
	s.Code = &v
	return s
}

type QueryRobotTaskDetailResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryRobotTaskDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryRobotTaskDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryRobotTaskDetailResponse) GoString() string {
	return s.String()
}

func (s *QueryRobotTaskDetailResponse) SetHeaders(v map[string]*string) *QueryRobotTaskDetailResponse {
	s.Headers = v
	return s
}

func (s *QueryRobotTaskDetailResponse) SetBody(v *QueryRobotTaskDetailResponseBody) *QueryRobotTaskDetailResponse {
	s.Body = v
	return s
}

type QueryRobotTaskListRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TaskName             *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	Status               *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Time                 *string `json:"Time,omitempty" xml:"Time,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNo               *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
}

func (s QueryRobotTaskListRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryRobotTaskListRequest) GoString() string {
	return s.String()
}

func (s *QueryRobotTaskListRequest) SetOwnerId(v int64) *QueryRobotTaskListRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryRobotTaskListRequest) SetResourceOwnerAccount(v string) *QueryRobotTaskListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryRobotTaskListRequest) SetResourceOwnerId(v int64) *QueryRobotTaskListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryRobotTaskListRequest) SetTaskName(v string) *QueryRobotTaskListRequest {
	s.TaskName = &v
	return s
}

func (s *QueryRobotTaskListRequest) SetStatus(v string) *QueryRobotTaskListRequest {
	s.Status = &v
	return s
}

func (s *QueryRobotTaskListRequest) SetTime(v string) *QueryRobotTaskListRequest {
	s.Time = &v
	return s
}

func (s *QueryRobotTaskListRequest) SetPageSize(v int32) *QueryRobotTaskListRequest {
	s.PageSize = &v
	return s
}

func (s *QueryRobotTaskListRequest) SetPageNo(v int32) *QueryRobotTaskListRequest {
	s.PageNo = &v
	return s
}

type QueryRobotTaskListResponseBody struct {
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Message    *string `json:"Message,omitempty" xml:"Message,omitempty"`
	PageSize   *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Data       *string `json:"Data,omitempty" xml:"Data,omitempty"`
	PageNo     *string `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	Code       *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s QueryRobotTaskListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryRobotTaskListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryRobotTaskListResponseBody) SetTotalCount(v string) *QueryRobotTaskListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *QueryRobotTaskListResponseBody) SetRequestId(v string) *QueryRobotTaskListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryRobotTaskListResponseBody) SetMessage(v string) *QueryRobotTaskListResponseBody {
	s.Message = &v
	return s
}

func (s *QueryRobotTaskListResponseBody) SetPageSize(v string) *QueryRobotTaskListResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryRobotTaskListResponseBody) SetData(v string) *QueryRobotTaskListResponseBody {
	s.Data = &v
	return s
}

func (s *QueryRobotTaskListResponseBody) SetPageNo(v string) *QueryRobotTaskListResponseBody {
	s.PageNo = &v
	return s
}

func (s *QueryRobotTaskListResponseBody) SetCode(v string) *QueryRobotTaskListResponseBody {
	s.Code = &v
	return s
}

type QueryRobotTaskListResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryRobotTaskListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryRobotTaskListResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryRobotTaskListResponse) GoString() string {
	return s.String()
}

func (s *QueryRobotTaskListResponse) SetHeaders(v map[string]*string) *QueryRobotTaskListResponse {
	s.Headers = v
	return s
}

func (s *QueryRobotTaskListResponse) SetBody(v *QueryRobotTaskListResponseBody) *QueryRobotTaskListResponse {
	s.Body = v
	return s
}

type QueryRobotv2AllListRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s QueryRobotv2AllListRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryRobotv2AllListRequest) GoString() string {
	return s.String()
}

func (s *QueryRobotv2AllListRequest) SetOwnerId(v int64) *QueryRobotv2AllListRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryRobotv2AllListRequest) SetResourceOwnerAccount(v string) *QueryRobotv2AllListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryRobotv2AllListRequest) SetResourceOwnerId(v int64) *QueryRobotv2AllListRequest {
	s.ResourceOwnerId = &v
	return s
}

type QueryRobotv2AllListResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s QueryRobotv2AllListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryRobotv2AllListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryRobotv2AllListResponseBody) SetMessage(v string) *QueryRobotv2AllListResponseBody {
	s.Message = &v
	return s
}

func (s *QueryRobotv2AllListResponseBody) SetRequestId(v string) *QueryRobotv2AllListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryRobotv2AllListResponseBody) SetData(v string) *QueryRobotv2AllListResponseBody {
	s.Data = &v
	return s
}

func (s *QueryRobotv2AllListResponseBody) SetCode(v string) *QueryRobotv2AllListResponseBody {
	s.Code = &v
	return s
}

type QueryRobotv2AllListResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryRobotv2AllListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryRobotv2AllListResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryRobotv2AllListResponse) GoString() string {
	return s.String()
}

func (s *QueryRobotv2AllListResponse) SetHeaders(v map[string]*string) *QueryRobotv2AllListResponse {
	s.Headers = v
	return s
}

func (s *QueryRobotv2AllListResponse) SetBody(v *QueryRobotv2AllListResponseBody) *QueryRobotv2AllListResponse {
	s.Body = v
	return s
}

type QueryRtcNumberAuthStatusRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	PhoneNumber          *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
}

func (s QueryRtcNumberAuthStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryRtcNumberAuthStatusRequest) GoString() string {
	return s.String()
}

func (s *QueryRtcNumberAuthStatusRequest) SetOwnerId(v int64) *QueryRtcNumberAuthStatusRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryRtcNumberAuthStatusRequest) SetResourceOwnerAccount(v string) *QueryRtcNumberAuthStatusRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryRtcNumberAuthStatusRequest) SetResourceOwnerId(v int64) *QueryRtcNumberAuthStatusRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryRtcNumberAuthStatusRequest) SetPhoneNumber(v string) *QueryRtcNumberAuthStatusRequest {
	s.PhoneNumber = &v
	return s
}

type QueryRtcNumberAuthStatusResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Module    *string `json:"Module,omitempty" xml:"Module,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s QueryRtcNumberAuthStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryRtcNumberAuthStatusResponseBody) GoString() string {
	return s.String()
}

func (s *QueryRtcNumberAuthStatusResponseBody) SetMessage(v string) *QueryRtcNumberAuthStatusResponseBody {
	s.Message = &v
	return s
}

func (s *QueryRtcNumberAuthStatusResponseBody) SetRequestId(v string) *QueryRtcNumberAuthStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryRtcNumberAuthStatusResponseBody) SetModule(v string) *QueryRtcNumberAuthStatusResponseBody {
	s.Module = &v
	return s
}

func (s *QueryRtcNumberAuthStatusResponseBody) SetCode(v string) *QueryRtcNumberAuthStatusResponseBody {
	s.Code = &v
	return s
}

type QueryRtcNumberAuthStatusResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryRtcNumberAuthStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryRtcNumberAuthStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryRtcNumberAuthStatusResponse) GoString() string {
	return s.String()
}

func (s *QueryRtcNumberAuthStatusResponse) SetHeaders(v map[string]*string) *QueryRtcNumberAuthStatusResponse {
	s.Headers = v
	return s
}

func (s *QueryRtcNumberAuthStatusResponse) SetBody(v *QueryRtcNumberAuthStatusResponseBody) *QueryRtcNumberAuthStatusResponse {
	s.Body = v
	return s
}

type QueryVoipNumberBindInfosRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	PhoneNumber          *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	VoipId               *string `json:"VoipId,omitempty" xml:"VoipId,omitempty"`
}

func (s QueryVoipNumberBindInfosRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryVoipNumberBindInfosRequest) GoString() string {
	return s.String()
}

func (s *QueryVoipNumberBindInfosRequest) SetOwnerId(v int64) *QueryVoipNumberBindInfosRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryVoipNumberBindInfosRequest) SetResourceOwnerAccount(v string) *QueryVoipNumberBindInfosRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryVoipNumberBindInfosRequest) SetResourceOwnerId(v int64) *QueryVoipNumberBindInfosRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryVoipNumberBindInfosRequest) SetPhoneNumber(v string) *QueryVoipNumberBindInfosRequest {
	s.PhoneNumber = &v
	return s
}

func (s *QueryVoipNumberBindInfosRequest) SetVoipId(v string) *QueryVoipNumberBindInfosRequest {
	s.VoipId = &v
	return s
}

type QueryVoipNumberBindInfosResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Module    *string `json:"Module,omitempty" xml:"Module,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s QueryVoipNumberBindInfosResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryVoipNumberBindInfosResponseBody) GoString() string {
	return s.String()
}

func (s *QueryVoipNumberBindInfosResponseBody) SetMessage(v string) *QueryVoipNumberBindInfosResponseBody {
	s.Message = &v
	return s
}

func (s *QueryVoipNumberBindInfosResponseBody) SetRequestId(v string) *QueryVoipNumberBindInfosResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryVoipNumberBindInfosResponseBody) SetModule(v string) *QueryVoipNumberBindInfosResponseBody {
	s.Module = &v
	return s
}

func (s *QueryVoipNumberBindInfosResponseBody) SetCode(v string) *QueryVoipNumberBindInfosResponseBody {
	s.Code = &v
	return s
}

type QueryVoipNumberBindInfosResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryVoipNumberBindInfosResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryVoipNumberBindInfosResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryVoipNumberBindInfosResponse) GoString() string {
	return s.String()
}

func (s *QueryVoipNumberBindInfosResponse) SetHeaders(v map[string]*string) *QueryVoipNumberBindInfosResponse {
	s.Headers = v
	return s
}

func (s *QueryVoipNumberBindInfosResponse) SetBody(v *QueryVoipNumberBindInfosResponseBody) *QueryVoipNumberBindInfosResponse {
	s.Body = v
	return s
}

type ReportVoipProblemsRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ChannelId            *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	VoipId               *string `json:"VoipId,omitempty" xml:"VoipId,omitempty"`
	Title                *string `json:"Title,omitempty" xml:"Title,omitempty"`
	Desc                 *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
}

func (s ReportVoipProblemsRequest) String() string {
	return tea.Prettify(s)
}

func (s ReportVoipProblemsRequest) GoString() string {
	return s.String()
}

func (s *ReportVoipProblemsRequest) SetOwnerId(v int64) *ReportVoipProblemsRequest {
	s.OwnerId = &v
	return s
}

func (s *ReportVoipProblemsRequest) SetResourceOwnerAccount(v string) *ReportVoipProblemsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ReportVoipProblemsRequest) SetResourceOwnerId(v int64) *ReportVoipProblemsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ReportVoipProblemsRequest) SetChannelId(v string) *ReportVoipProblemsRequest {
	s.ChannelId = &v
	return s
}

func (s *ReportVoipProblemsRequest) SetVoipId(v string) *ReportVoipProblemsRequest {
	s.VoipId = &v
	return s
}

func (s *ReportVoipProblemsRequest) SetTitle(v string) *ReportVoipProblemsRequest {
	s.Title = &v
	return s
}

func (s *ReportVoipProblemsRequest) SetDesc(v string) *ReportVoipProblemsRequest {
	s.Desc = &v
	return s
}

type ReportVoipProblemsResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Module    *string `json:"Module,omitempty" xml:"Module,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s ReportVoipProblemsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReportVoipProblemsResponseBody) GoString() string {
	return s.String()
}

func (s *ReportVoipProblemsResponseBody) SetMessage(v string) *ReportVoipProblemsResponseBody {
	s.Message = &v
	return s
}

func (s *ReportVoipProblemsResponseBody) SetRequestId(v string) *ReportVoipProblemsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReportVoipProblemsResponseBody) SetModule(v string) *ReportVoipProblemsResponseBody {
	s.Module = &v
	return s
}

func (s *ReportVoipProblemsResponseBody) SetCode(v string) *ReportVoipProblemsResponseBody {
	s.Code = &v
	return s
}

type ReportVoipProblemsResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ReportVoipProblemsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ReportVoipProblemsResponse) String() string {
	return tea.Prettify(s)
}

func (s ReportVoipProblemsResponse) GoString() string {
	return s.String()
}

func (s *ReportVoipProblemsResponse) SetHeaders(v map[string]*string) *ReportVoipProblemsResponse {
	s.Headers = v
	return s
}

func (s *ReportVoipProblemsResponse) SetBody(v *ReportVoipProblemsResponseBody) *ReportVoipProblemsResponse {
	s.Body = v
	return s
}

type SingleCallByTtsRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	CalledShowNumber     *string `json:"CalledShowNumber,omitempty" xml:"CalledShowNumber,omitempty"`
	CalledNumber         *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	TtsCode              *string `json:"TtsCode,omitempty" xml:"TtsCode,omitempty"`
	TtsParam             *string `json:"TtsParam,omitempty" xml:"TtsParam,omitempty"`
	PlayTimes            *int32  `json:"PlayTimes,omitempty" xml:"PlayTimes,omitempty"`
	Volume               *int32  `json:"Volume,omitempty" xml:"Volume,omitempty"`
	Speed                *int32  `json:"Speed,omitempty" xml:"Speed,omitempty"`
	OutId                *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
}

func (s SingleCallByTtsRequest) String() string {
	return tea.Prettify(s)
}

func (s SingleCallByTtsRequest) GoString() string {
	return s.String()
}

func (s *SingleCallByTtsRequest) SetOwnerId(v int64) *SingleCallByTtsRequest {
	s.OwnerId = &v
	return s
}

func (s *SingleCallByTtsRequest) SetResourceOwnerAccount(v string) *SingleCallByTtsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SingleCallByTtsRequest) SetResourceOwnerId(v int64) *SingleCallByTtsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SingleCallByTtsRequest) SetCalledShowNumber(v string) *SingleCallByTtsRequest {
	s.CalledShowNumber = &v
	return s
}

func (s *SingleCallByTtsRequest) SetCalledNumber(v string) *SingleCallByTtsRequest {
	s.CalledNumber = &v
	return s
}

func (s *SingleCallByTtsRequest) SetTtsCode(v string) *SingleCallByTtsRequest {
	s.TtsCode = &v
	return s
}

func (s *SingleCallByTtsRequest) SetTtsParam(v string) *SingleCallByTtsRequest {
	s.TtsParam = &v
	return s
}

func (s *SingleCallByTtsRequest) SetPlayTimes(v int32) *SingleCallByTtsRequest {
	s.PlayTimes = &v
	return s
}

func (s *SingleCallByTtsRequest) SetVolume(v int32) *SingleCallByTtsRequest {
	s.Volume = &v
	return s
}

func (s *SingleCallByTtsRequest) SetSpeed(v int32) *SingleCallByTtsRequest {
	s.Speed = &v
	return s
}

func (s *SingleCallByTtsRequest) SetOutId(v string) *SingleCallByTtsRequest {
	s.OutId = &v
	return s
}

type SingleCallByTtsResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	CallId    *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s SingleCallByTtsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SingleCallByTtsResponseBody) GoString() string {
	return s.String()
}

func (s *SingleCallByTtsResponseBody) SetMessage(v string) *SingleCallByTtsResponseBody {
	s.Message = &v
	return s
}

func (s *SingleCallByTtsResponseBody) SetRequestId(v string) *SingleCallByTtsResponseBody {
	s.RequestId = &v
	return s
}

func (s *SingleCallByTtsResponseBody) SetCallId(v string) *SingleCallByTtsResponseBody {
	s.CallId = &v
	return s
}

func (s *SingleCallByTtsResponseBody) SetCode(v string) *SingleCallByTtsResponseBody {
	s.Code = &v
	return s
}

type SingleCallByTtsResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SingleCallByTtsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SingleCallByTtsResponse) String() string {
	return tea.Prettify(s)
}

func (s SingleCallByTtsResponse) GoString() string {
	return s.String()
}

func (s *SingleCallByTtsResponse) SetHeaders(v map[string]*string) *SingleCallByTtsResponse {
	s.Headers = v
	return s
}

func (s *SingleCallByTtsResponse) SetBody(v *SingleCallByTtsResponseBody) *SingleCallByTtsResponse {
	s.Body = v
	return s
}

type SingleCallByVoiceRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	CalledShowNumber     *string `json:"CalledShowNumber,omitempty" xml:"CalledShowNumber,omitempty"`
	CalledNumber         *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	VoiceCode            *string `json:"VoiceCode,omitempty" xml:"VoiceCode,omitempty"`
	PlayTimes            *int32  `json:"PlayTimes,omitempty" xml:"PlayTimes,omitempty"`
	Volume               *int32  `json:"Volume,omitempty" xml:"Volume,omitempty"`
	Speed                *int32  `json:"Speed,omitempty" xml:"Speed,omitempty"`
	OutId                *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
}

func (s SingleCallByVoiceRequest) String() string {
	return tea.Prettify(s)
}

func (s SingleCallByVoiceRequest) GoString() string {
	return s.String()
}

func (s *SingleCallByVoiceRequest) SetOwnerId(v int64) *SingleCallByVoiceRequest {
	s.OwnerId = &v
	return s
}

func (s *SingleCallByVoiceRequest) SetResourceOwnerAccount(v string) *SingleCallByVoiceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SingleCallByVoiceRequest) SetResourceOwnerId(v int64) *SingleCallByVoiceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SingleCallByVoiceRequest) SetCalledShowNumber(v string) *SingleCallByVoiceRequest {
	s.CalledShowNumber = &v
	return s
}

func (s *SingleCallByVoiceRequest) SetCalledNumber(v string) *SingleCallByVoiceRequest {
	s.CalledNumber = &v
	return s
}

func (s *SingleCallByVoiceRequest) SetVoiceCode(v string) *SingleCallByVoiceRequest {
	s.VoiceCode = &v
	return s
}

func (s *SingleCallByVoiceRequest) SetPlayTimes(v int32) *SingleCallByVoiceRequest {
	s.PlayTimes = &v
	return s
}

func (s *SingleCallByVoiceRequest) SetVolume(v int32) *SingleCallByVoiceRequest {
	s.Volume = &v
	return s
}

func (s *SingleCallByVoiceRequest) SetSpeed(v int32) *SingleCallByVoiceRequest {
	s.Speed = &v
	return s
}

func (s *SingleCallByVoiceRequest) SetOutId(v string) *SingleCallByVoiceRequest {
	s.OutId = &v
	return s
}

type SingleCallByVoiceResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	CallId    *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s SingleCallByVoiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SingleCallByVoiceResponseBody) GoString() string {
	return s.String()
}

func (s *SingleCallByVoiceResponseBody) SetMessage(v string) *SingleCallByVoiceResponseBody {
	s.Message = &v
	return s
}

func (s *SingleCallByVoiceResponseBody) SetRequestId(v string) *SingleCallByVoiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *SingleCallByVoiceResponseBody) SetCallId(v string) *SingleCallByVoiceResponseBody {
	s.CallId = &v
	return s
}

func (s *SingleCallByVoiceResponseBody) SetCode(v string) *SingleCallByVoiceResponseBody {
	s.Code = &v
	return s
}

type SingleCallByVoiceResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SingleCallByVoiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SingleCallByVoiceResponse) String() string {
	return tea.Prettify(s)
}

func (s SingleCallByVoiceResponse) GoString() string {
	return s.String()
}

func (s *SingleCallByVoiceResponse) SetHeaders(v map[string]*string) *SingleCallByVoiceResponse {
	s.Headers = v
	return s
}

func (s *SingleCallByVoiceResponse) SetBody(v *SingleCallByVoiceResponseBody) *SingleCallByVoiceResponse {
	s.Body = v
	return s
}

type SmartCallRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	CalledShowNumber     *string `json:"CalledShowNumber,omitempty" xml:"CalledShowNumber,omitempty"`
	CalledNumber         *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	VoiceCode            *string `json:"VoiceCode,omitempty" xml:"VoiceCode,omitempty"`
	RecordFlag           *bool   `json:"RecordFlag,omitempty" xml:"RecordFlag,omitempty"`
	Volume               *int32  `json:"Volume,omitempty" xml:"Volume,omitempty"`
	Speed                *int32  `json:"Speed,omitempty" xml:"Speed,omitempty"`
	AsrModelId           *string `json:"AsrModelId,omitempty" xml:"AsrModelId,omitempty"`
	PauseTime            *int32  `json:"PauseTime,omitempty" xml:"PauseTime,omitempty"`
	MuteTime             *int32  `json:"MuteTime,omitempty" xml:"MuteTime,omitempty"`
	ActionCodeBreak      *bool   `json:"ActionCodeBreak,omitempty" xml:"ActionCodeBreak,omitempty"`
	OutId                *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	DynamicId            *string `json:"DynamicId,omitempty" xml:"DynamicId,omitempty"`
	EarlyMediaAsr        *bool   `json:"EarlyMediaAsr,omitempty" xml:"EarlyMediaAsr,omitempty"`
	VoiceCodeParam       *string `json:"VoiceCodeParam,omitempty" xml:"VoiceCodeParam,omitempty"`
	SessionTimeout       *int32  `json:"SessionTimeout,omitempty" xml:"SessionTimeout,omitempty"`
	ActionCodeTimeBreak  *int32  `json:"ActionCodeTimeBreak,omitempty" xml:"ActionCodeTimeBreak,omitempty"`
	TtsStyle             *string `json:"TtsStyle,omitempty" xml:"TtsStyle,omitempty"`
	TtsVolume            *int32  `json:"TtsVolume,omitempty" xml:"TtsVolume,omitempty"`
	TtsSpeed             *int32  `json:"TtsSpeed,omitempty" xml:"TtsSpeed,omitempty"`
	TtsConf              *bool   `json:"TtsConf,omitempty" xml:"TtsConf,omitempty"`
	AsrBaseId            *string `json:"AsrBaseId,omitempty" xml:"AsrBaseId,omitempty"`
}

func (s SmartCallRequest) String() string {
	return tea.Prettify(s)
}

func (s SmartCallRequest) GoString() string {
	return s.String()
}

func (s *SmartCallRequest) SetOwnerId(v int64) *SmartCallRequest {
	s.OwnerId = &v
	return s
}

func (s *SmartCallRequest) SetResourceOwnerAccount(v string) *SmartCallRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SmartCallRequest) SetResourceOwnerId(v int64) *SmartCallRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SmartCallRequest) SetCalledShowNumber(v string) *SmartCallRequest {
	s.CalledShowNumber = &v
	return s
}

func (s *SmartCallRequest) SetCalledNumber(v string) *SmartCallRequest {
	s.CalledNumber = &v
	return s
}

func (s *SmartCallRequest) SetVoiceCode(v string) *SmartCallRequest {
	s.VoiceCode = &v
	return s
}

func (s *SmartCallRequest) SetRecordFlag(v bool) *SmartCallRequest {
	s.RecordFlag = &v
	return s
}

func (s *SmartCallRequest) SetVolume(v int32) *SmartCallRequest {
	s.Volume = &v
	return s
}

func (s *SmartCallRequest) SetSpeed(v int32) *SmartCallRequest {
	s.Speed = &v
	return s
}

func (s *SmartCallRequest) SetAsrModelId(v string) *SmartCallRequest {
	s.AsrModelId = &v
	return s
}

func (s *SmartCallRequest) SetPauseTime(v int32) *SmartCallRequest {
	s.PauseTime = &v
	return s
}

func (s *SmartCallRequest) SetMuteTime(v int32) *SmartCallRequest {
	s.MuteTime = &v
	return s
}

func (s *SmartCallRequest) SetActionCodeBreak(v bool) *SmartCallRequest {
	s.ActionCodeBreak = &v
	return s
}

func (s *SmartCallRequest) SetOutId(v string) *SmartCallRequest {
	s.OutId = &v
	return s
}

func (s *SmartCallRequest) SetDynamicId(v string) *SmartCallRequest {
	s.DynamicId = &v
	return s
}

func (s *SmartCallRequest) SetEarlyMediaAsr(v bool) *SmartCallRequest {
	s.EarlyMediaAsr = &v
	return s
}

func (s *SmartCallRequest) SetVoiceCodeParam(v string) *SmartCallRequest {
	s.VoiceCodeParam = &v
	return s
}

func (s *SmartCallRequest) SetSessionTimeout(v int32) *SmartCallRequest {
	s.SessionTimeout = &v
	return s
}

func (s *SmartCallRequest) SetActionCodeTimeBreak(v int32) *SmartCallRequest {
	s.ActionCodeTimeBreak = &v
	return s
}

func (s *SmartCallRequest) SetTtsStyle(v string) *SmartCallRequest {
	s.TtsStyle = &v
	return s
}

func (s *SmartCallRequest) SetTtsVolume(v int32) *SmartCallRequest {
	s.TtsVolume = &v
	return s
}

func (s *SmartCallRequest) SetTtsSpeed(v int32) *SmartCallRequest {
	s.TtsSpeed = &v
	return s
}

func (s *SmartCallRequest) SetTtsConf(v bool) *SmartCallRequest {
	s.TtsConf = &v
	return s
}

func (s *SmartCallRequest) SetAsrBaseId(v string) *SmartCallRequest {
	s.AsrBaseId = &v
	return s
}

type SmartCallResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	CallId    *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s SmartCallResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SmartCallResponseBody) GoString() string {
	return s.String()
}

func (s *SmartCallResponseBody) SetMessage(v string) *SmartCallResponseBody {
	s.Message = &v
	return s
}

func (s *SmartCallResponseBody) SetRequestId(v string) *SmartCallResponseBody {
	s.RequestId = &v
	return s
}

func (s *SmartCallResponseBody) SetCallId(v string) *SmartCallResponseBody {
	s.CallId = &v
	return s
}

func (s *SmartCallResponseBody) SetCode(v string) *SmartCallResponseBody {
	s.Code = &v
	return s
}

type SmartCallResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SmartCallResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SmartCallResponse) String() string {
	return tea.Prettify(s)
}

func (s SmartCallResponse) GoString() string {
	return s.String()
}

func (s *SmartCallResponse) SetHeaders(v map[string]*string) *SmartCallResponse {
	s.Headers = v
	return s
}

func (s *SmartCallResponse) SetBody(v *SmartCallResponseBody) *SmartCallResponse {
	s.Body = v
	return s
}

type SmartCallOperateRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	CallId               *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	Command              *string `json:"Command,omitempty" xml:"Command,omitempty"`
	Param                *string `json:"Param,omitempty" xml:"Param,omitempty"`
}

func (s SmartCallOperateRequest) String() string {
	return tea.Prettify(s)
}

func (s SmartCallOperateRequest) GoString() string {
	return s.String()
}

func (s *SmartCallOperateRequest) SetOwnerId(v int64) *SmartCallOperateRequest {
	s.OwnerId = &v
	return s
}

func (s *SmartCallOperateRequest) SetResourceOwnerAccount(v string) *SmartCallOperateRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SmartCallOperateRequest) SetResourceOwnerId(v int64) *SmartCallOperateRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SmartCallOperateRequest) SetCallId(v string) *SmartCallOperateRequest {
	s.CallId = &v
	return s
}

func (s *SmartCallOperateRequest) SetCommand(v string) *SmartCallOperateRequest {
	s.Command = &v
	return s
}

func (s *SmartCallOperateRequest) SetParam(v string) *SmartCallOperateRequest {
	s.Param = &v
	return s
}

type SmartCallOperateResponseBody struct {
	Status    *bool   `json:"Status,omitempty" xml:"Status,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s SmartCallOperateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SmartCallOperateResponseBody) GoString() string {
	return s.String()
}

func (s *SmartCallOperateResponseBody) SetStatus(v bool) *SmartCallOperateResponseBody {
	s.Status = &v
	return s
}

func (s *SmartCallOperateResponseBody) SetMessage(v string) *SmartCallOperateResponseBody {
	s.Message = &v
	return s
}

func (s *SmartCallOperateResponseBody) SetRequestId(v string) *SmartCallOperateResponseBody {
	s.RequestId = &v
	return s
}

func (s *SmartCallOperateResponseBody) SetCode(v string) *SmartCallOperateResponseBody {
	s.Code = &v
	return s
}

type SmartCallOperateResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SmartCallOperateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SmartCallOperateResponse) String() string {
	return tea.Prettify(s)
}

func (s SmartCallOperateResponse) GoString() string {
	return s.String()
}

func (s *SmartCallOperateResponse) SetHeaders(v map[string]*string) *SmartCallOperateResponse {
	s.Headers = v
	return s
}

func (s *SmartCallOperateResponse) SetBody(v *SmartCallOperateResponseBody) *SmartCallOperateResponse {
	s.Body = v
	return s
}

type StartMicroOutboundRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	AccountType          *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	AccountId            *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	CommandCode          *string `json:"CommandCode,omitempty" xml:"CommandCode,omitempty"`
	CallingNumber        *string `json:"CallingNumber,omitempty" xml:"CallingNumber,omitempty"`
	CalledNumber         *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	ExtInfo              *string `json:"ExtInfo,omitempty" xml:"ExtInfo,omitempty"`
	AppName              *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
}

func (s StartMicroOutboundRequest) String() string {
	return tea.Prettify(s)
}

func (s StartMicroOutboundRequest) GoString() string {
	return s.String()
}

func (s *StartMicroOutboundRequest) SetOwnerId(v int64) *StartMicroOutboundRequest {
	s.OwnerId = &v
	return s
}

func (s *StartMicroOutboundRequest) SetResourceOwnerAccount(v string) *StartMicroOutboundRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *StartMicroOutboundRequest) SetResourceOwnerId(v int64) *StartMicroOutboundRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *StartMicroOutboundRequest) SetProdCode(v string) *StartMicroOutboundRequest {
	s.ProdCode = &v
	return s
}

func (s *StartMicroOutboundRequest) SetAccountType(v string) *StartMicroOutboundRequest {
	s.AccountType = &v
	return s
}

func (s *StartMicroOutboundRequest) SetAccountId(v string) *StartMicroOutboundRequest {
	s.AccountId = &v
	return s
}

func (s *StartMicroOutboundRequest) SetCommandCode(v string) *StartMicroOutboundRequest {
	s.CommandCode = &v
	return s
}

func (s *StartMicroOutboundRequest) SetCallingNumber(v string) *StartMicroOutboundRequest {
	s.CallingNumber = &v
	return s
}

func (s *StartMicroOutboundRequest) SetCalledNumber(v string) *StartMicroOutboundRequest {
	s.CalledNumber = &v
	return s
}

func (s *StartMicroOutboundRequest) SetExtInfo(v string) *StartMicroOutboundRequest {
	s.ExtInfo = &v
	return s
}

func (s *StartMicroOutboundRequest) SetAppName(v string) *StartMicroOutboundRequest {
	s.AppName = &v
	return s
}

type StartMicroOutboundResponseBody struct {
	InvokeCreateTime *string `json:"InvokeCreateTime,omitempty" xml:"InvokeCreateTime,omitempty"`
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Message          *string `json:"Message,omitempty" xml:"Message,omitempty"`
	InvokeCmdId      *string `json:"InvokeCmdId,omitempty" xml:"InvokeCmdId,omitempty"`
	CustomerInfo     *string `json:"CustomerInfo,omitempty" xml:"CustomerInfo,omitempty"`
	Code             *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s StartMicroOutboundResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartMicroOutboundResponseBody) GoString() string {
	return s.String()
}

func (s *StartMicroOutboundResponseBody) SetInvokeCreateTime(v string) *StartMicroOutboundResponseBody {
	s.InvokeCreateTime = &v
	return s
}

func (s *StartMicroOutboundResponseBody) SetRequestId(v string) *StartMicroOutboundResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartMicroOutboundResponseBody) SetMessage(v string) *StartMicroOutboundResponseBody {
	s.Message = &v
	return s
}

func (s *StartMicroOutboundResponseBody) SetInvokeCmdId(v string) *StartMicroOutboundResponseBody {
	s.InvokeCmdId = &v
	return s
}

func (s *StartMicroOutboundResponseBody) SetCustomerInfo(v string) *StartMicroOutboundResponseBody {
	s.CustomerInfo = &v
	return s
}

func (s *StartMicroOutboundResponseBody) SetCode(v string) *StartMicroOutboundResponseBody {
	s.Code = &v
	return s
}

type StartMicroOutboundResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StartMicroOutboundResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartMicroOutboundResponse) String() string {
	return tea.Prettify(s)
}

func (s StartMicroOutboundResponse) GoString() string {
	return s.String()
}

func (s *StartMicroOutboundResponse) SetHeaders(v map[string]*string) *StartMicroOutboundResponse {
	s.Headers = v
	return s
}

func (s *StartMicroOutboundResponse) SetBody(v *StartMicroOutboundResponseBody) *StartMicroOutboundResponse {
	s.Body = v
	return s
}

type StartRobotTaskRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TaskId               *int64  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	ScheduleTime         *string `json:"ScheduleTime,omitempty" xml:"ScheduleTime,omitempty"`
}

func (s StartRobotTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s StartRobotTaskRequest) GoString() string {
	return s.String()
}

func (s *StartRobotTaskRequest) SetOwnerId(v int64) *StartRobotTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *StartRobotTaskRequest) SetResourceOwnerAccount(v string) *StartRobotTaskRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *StartRobotTaskRequest) SetResourceOwnerId(v int64) *StartRobotTaskRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *StartRobotTaskRequest) SetTaskId(v int64) *StartRobotTaskRequest {
	s.TaskId = &v
	return s
}

func (s *StartRobotTaskRequest) SetScheduleTime(v string) *StartRobotTaskRequest {
	s.ScheduleTime = &v
	return s
}

type StartRobotTaskResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s StartRobotTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartRobotTaskResponseBody) GoString() string {
	return s.String()
}

func (s *StartRobotTaskResponseBody) SetMessage(v string) *StartRobotTaskResponseBody {
	s.Message = &v
	return s
}

func (s *StartRobotTaskResponseBody) SetRequestId(v string) *StartRobotTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartRobotTaskResponseBody) SetData(v string) *StartRobotTaskResponseBody {
	s.Data = &v
	return s
}

func (s *StartRobotTaskResponseBody) SetCode(v string) *StartRobotTaskResponseBody {
	s.Code = &v
	return s
}

type StartRobotTaskResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StartRobotTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartRobotTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s StartRobotTaskResponse) GoString() string {
	return s.String()
}

func (s *StartRobotTaskResponse) SetHeaders(v map[string]*string) *StartRobotTaskResponse {
	s.Headers = v
	return s
}

func (s *StartRobotTaskResponse) SetBody(v *StartRobotTaskResponseBody) *StartRobotTaskResponse {
	s.Body = v
	return s
}

type StopRobotTaskRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TaskId               *int64  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s StopRobotTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s StopRobotTaskRequest) GoString() string {
	return s.String()
}

func (s *StopRobotTaskRequest) SetOwnerId(v int64) *StopRobotTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *StopRobotTaskRequest) SetResourceOwnerAccount(v string) *StopRobotTaskRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *StopRobotTaskRequest) SetResourceOwnerId(v int64) *StopRobotTaskRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *StopRobotTaskRequest) SetTaskId(v int64) *StopRobotTaskRequest {
	s.TaskId = &v
	return s
}

type StopRobotTaskResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s StopRobotTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopRobotTaskResponseBody) GoString() string {
	return s.String()
}

func (s *StopRobotTaskResponseBody) SetMessage(v string) *StopRobotTaskResponseBody {
	s.Message = &v
	return s
}

func (s *StopRobotTaskResponseBody) SetRequestId(v string) *StopRobotTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopRobotTaskResponseBody) SetData(v string) *StopRobotTaskResponseBody {
	s.Data = &v
	return s
}

func (s *StopRobotTaskResponseBody) SetCode(v string) *StopRobotTaskResponseBody {
	s.Code = &v
	return s
}

type StopRobotTaskResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StopRobotTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopRobotTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s StopRobotTaskResponse) GoString() string {
	return s.String()
}

func (s *StopRobotTaskResponse) SetHeaders(v map[string]*string) *StopRobotTaskResponse {
	s.Headers = v
	return s
}

func (s *StopRobotTaskResponse) SetBody(v *StopRobotTaskResponseBody) *StopRobotTaskResponse {
	s.Body = v
	return s
}

type UnbindNumberAndVoipIdRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	PhoneNumber          *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	VoipId               *string `json:"VoipId,omitempty" xml:"VoipId,omitempty"`
}

func (s UnbindNumberAndVoipIdRequest) String() string {
	return tea.Prettify(s)
}

func (s UnbindNumberAndVoipIdRequest) GoString() string {
	return s.String()
}

func (s *UnbindNumberAndVoipIdRequest) SetOwnerId(v int64) *UnbindNumberAndVoipIdRequest {
	s.OwnerId = &v
	return s
}

func (s *UnbindNumberAndVoipIdRequest) SetResourceOwnerAccount(v string) *UnbindNumberAndVoipIdRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UnbindNumberAndVoipIdRequest) SetResourceOwnerId(v int64) *UnbindNumberAndVoipIdRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UnbindNumberAndVoipIdRequest) SetPhoneNumber(v string) *UnbindNumberAndVoipIdRequest {
	s.PhoneNumber = &v
	return s
}

func (s *UnbindNumberAndVoipIdRequest) SetVoipId(v string) *UnbindNumberAndVoipIdRequest {
	s.VoipId = &v
	return s
}

type UnbindNumberAndVoipIdResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Module    *string `json:"Module,omitempty" xml:"Module,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s UnbindNumberAndVoipIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnbindNumberAndVoipIdResponseBody) GoString() string {
	return s.String()
}

func (s *UnbindNumberAndVoipIdResponseBody) SetMessage(v string) *UnbindNumberAndVoipIdResponseBody {
	s.Message = &v
	return s
}

func (s *UnbindNumberAndVoipIdResponseBody) SetRequestId(v string) *UnbindNumberAndVoipIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnbindNumberAndVoipIdResponseBody) SetModule(v string) *UnbindNumberAndVoipIdResponseBody {
	s.Module = &v
	return s
}

func (s *UnbindNumberAndVoipIdResponseBody) SetCode(v string) *UnbindNumberAndVoipIdResponseBody {
	s.Code = &v
	return s
}

type UnbindNumberAndVoipIdResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UnbindNumberAndVoipIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UnbindNumberAndVoipIdResponse) String() string {
	return tea.Prettify(s)
}

func (s UnbindNumberAndVoipIdResponse) GoString() string {
	return s.String()
}

func (s *UnbindNumberAndVoipIdResponse) SetHeaders(v map[string]*string) *UnbindNumberAndVoipIdResponse {
	s.Headers = v
	return s
}

func (s *UnbindNumberAndVoipIdResponse) SetBody(v *UnbindNumberAndVoipIdResponseBody) *UnbindNumberAndVoipIdResponse {
	s.Body = v
	return s
}

type UndoRtcNumberAuthRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	PhoneNumber          *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
}

func (s UndoRtcNumberAuthRequest) String() string {
	return tea.Prettify(s)
}

func (s UndoRtcNumberAuthRequest) GoString() string {
	return s.String()
}

func (s *UndoRtcNumberAuthRequest) SetOwnerId(v int64) *UndoRtcNumberAuthRequest {
	s.OwnerId = &v
	return s
}

func (s *UndoRtcNumberAuthRequest) SetResourceOwnerAccount(v string) *UndoRtcNumberAuthRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UndoRtcNumberAuthRequest) SetResourceOwnerId(v int64) *UndoRtcNumberAuthRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UndoRtcNumberAuthRequest) SetPhoneNumber(v string) *UndoRtcNumberAuthRequest {
	s.PhoneNumber = &v
	return s
}

type UndoRtcNumberAuthResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Module    *string `json:"Module,omitempty" xml:"Module,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s UndoRtcNumberAuthResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UndoRtcNumberAuthResponseBody) GoString() string {
	return s.String()
}

func (s *UndoRtcNumberAuthResponseBody) SetMessage(v string) *UndoRtcNumberAuthResponseBody {
	s.Message = &v
	return s
}

func (s *UndoRtcNumberAuthResponseBody) SetRequestId(v string) *UndoRtcNumberAuthResponseBody {
	s.RequestId = &v
	return s
}

func (s *UndoRtcNumberAuthResponseBody) SetModule(v string) *UndoRtcNumberAuthResponseBody {
	s.Module = &v
	return s
}

func (s *UndoRtcNumberAuthResponseBody) SetCode(v string) *UndoRtcNumberAuthResponseBody {
	s.Code = &v
	return s
}

type UndoRtcNumberAuthResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UndoRtcNumberAuthResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UndoRtcNumberAuthResponse) String() string {
	return tea.Prettify(s)
}

func (s UndoRtcNumberAuthResponse) GoString() string {
	return s.String()
}

func (s *UndoRtcNumberAuthResponse) SetHeaders(v map[string]*string) *UndoRtcNumberAuthResponse {
	s.Headers = v
	return s
}

func (s *UndoRtcNumberAuthResponse) SetBody(v *UndoRtcNumberAuthResponseBody) *UndoRtcNumberAuthResponse {
	s.Body = v
	return s
}

type UploadRobotTaskCalledFileRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Id                   *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	CalledNumber         *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	TtsParam             *string `json:"TtsParam,omitempty" xml:"TtsParam,omitempty"`
	TtsParamHead         *string `json:"TtsParamHead,omitempty" xml:"TtsParamHead,omitempty"`
}

func (s UploadRobotTaskCalledFileRequest) String() string {
	return tea.Prettify(s)
}

func (s UploadRobotTaskCalledFileRequest) GoString() string {
	return s.String()
}

func (s *UploadRobotTaskCalledFileRequest) SetOwnerId(v int64) *UploadRobotTaskCalledFileRequest {
	s.OwnerId = &v
	return s
}

func (s *UploadRobotTaskCalledFileRequest) SetResourceOwnerAccount(v string) *UploadRobotTaskCalledFileRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UploadRobotTaskCalledFileRequest) SetResourceOwnerId(v int64) *UploadRobotTaskCalledFileRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UploadRobotTaskCalledFileRequest) SetId(v int64) *UploadRobotTaskCalledFileRequest {
	s.Id = &v
	return s
}

func (s *UploadRobotTaskCalledFileRequest) SetCalledNumber(v string) *UploadRobotTaskCalledFileRequest {
	s.CalledNumber = &v
	return s
}

func (s *UploadRobotTaskCalledFileRequest) SetTtsParam(v string) *UploadRobotTaskCalledFileRequest {
	s.TtsParam = &v
	return s
}

func (s *UploadRobotTaskCalledFileRequest) SetTtsParamHead(v string) *UploadRobotTaskCalledFileRequest {
	s.TtsParamHead = &v
	return s
}

type UploadRobotTaskCalledFileResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s UploadRobotTaskCalledFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UploadRobotTaskCalledFileResponseBody) GoString() string {
	return s.String()
}

func (s *UploadRobotTaskCalledFileResponseBody) SetMessage(v string) *UploadRobotTaskCalledFileResponseBody {
	s.Message = &v
	return s
}

func (s *UploadRobotTaskCalledFileResponseBody) SetRequestId(v string) *UploadRobotTaskCalledFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *UploadRobotTaskCalledFileResponseBody) SetData(v string) *UploadRobotTaskCalledFileResponseBody {
	s.Data = &v
	return s
}

func (s *UploadRobotTaskCalledFileResponseBody) SetCode(v string) *UploadRobotTaskCalledFileResponseBody {
	s.Code = &v
	return s
}

type UploadRobotTaskCalledFileResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UploadRobotTaskCalledFileResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UploadRobotTaskCalledFileResponse) String() string {
	return tea.Prettify(s)
}

func (s UploadRobotTaskCalledFileResponse) GoString() string {
	return s.String()
}

func (s *UploadRobotTaskCalledFileResponse) SetHeaders(v map[string]*string) *UploadRobotTaskCalledFileResponse {
	s.Headers = v
	return s
}

func (s *UploadRobotTaskCalledFileResponse) SetBody(v *UploadRobotTaskCalledFileResponseBody) *UploadRobotTaskCalledFileResponse {
	s.Body = v
	return s
}

type VoipAddAccountRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	DeviceId             *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
}

func (s VoipAddAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s VoipAddAccountRequest) GoString() string {
	return s.String()
}

func (s *VoipAddAccountRequest) SetOwnerId(v int64) *VoipAddAccountRequest {
	s.OwnerId = &v
	return s
}

func (s *VoipAddAccountRequest) SetResourceOwnerAccount(v string) *VoipAddAccountRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *VoipAddAccountRequest) SetResourceOwnerId(v int64) *VoipAddAccountRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *VoipAddAccountRequest) SetDeviceId(v string) *VoipAddAccountRequest {
	s.DeviceId = &v
	return s
}

type VoipAddAccountResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Module    *string `json:"Module,omitempty" xml:"Module,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s VoipAddAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s VoipAddAccountResponseBody) GoString() string {
	return s.String()
}

func (s *VoipAddAccountResponseBody) SetMessage(v string) *VoipAddAccountResponseBody {
	s.Message = &v
	return s
}

func (s *VoipAddAccountResponseBody) SetRequestId(v string) *VoipAddAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *VoipAddAccountResponseBody) SetModule(v string) *VoipAddAccountResponseBody {
	s.Module = &v
	return s
}

func (s *VoipAddAccountResponseBody) SetCode(v string) *VoipAddAccountResponseBody {
	s.Code = &v
	return s
}

type VoipAddAccountResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *VoipAddAccountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s VoipAddAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s VoipAddAccountResponse) GoString() string {
	return s.String()
}

func (s *VoipAddAccountResponse) SetHeaders(v map[string]*string) *VoipAddAccountResponse {
	s.Headers = v
	return s
}

func (s *VoipAddAccountResponse) SetBody(v *VoipAddAccountResponseBody) *VoipAddAccountResponse {
	s.Body = v
	return s
}

type VoipGetTokenRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	VoipId               *string `json:"VoipId,omitempty" xml:"VoipId,omitempty"`
	DeviceId             *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	IsCustomAccount      *bool   `json:"IsCustomAccount,omitempty" xml:"IsCustomAccount,omitempty"`
}

func (s VoipGetTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s VoipGetTokenRequest) GoString() string {
	return s.String()
}

func (s *VoipGetTokenRequest) SetOwnerId(v int64) *VoipGetTokenRequest {
	s.OwnerId = &v
	return s
}

func (s *VoipGetTokenRequest) SetResourceOwnerAccount(v string) *VoipGetTokenRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *VoipGetTokenRequest) SetResourceOwnerId(v int64) *VoipGetTokenRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *VoipGetTokenRequest) SetVoipId(v string) *VoipGetTokenRequest {
	s.VoipId = &v
	return s
}

func (s *VoipGetTokenRequest) SetDeviceId(v string) *VoipGetTokenRequest {
	s.DeviceId = &v
	return s
}

func (s *VoipGetTokenRequest) SetIsCustomAccount(v bool) *VoipGetTokenRequest {
	s.IsCustomAccount = &v
	return s
}

type VoipGetTokenResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Module    *string `json:"Module,omitempty" xml:"Module,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s VoipGetTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s VoipGetTokenResponseBody) GoString() string {
	return s.String()
}

func (s *VoipGetTokenResponseBody) SetMessage(v string) *VoipGetTokenResponseBody {
	s.Message = &v
	return s
}

func (s *VoipGetTokenResponseBody) SetRequestId(v string) *VoipGetTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *VoipGetTokenResponseBody) SetModule(v string) *VoipGetTokenResponseBody {
	s.Module = &v
	return s
}

func (s *VoipGetTokenResponseBody) SetCode(v string) *VoipGetTokenResponseBody {
	s.Code = &v
	return s
}

type VoipGetTokenResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *VoipGetTokenResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s VoipGetTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s VoipGetTokenResponse) GoString() string {
	return s.String()
}

func (s *VoipGetTokenResponse) SetHeaders(v map[string]*string) *VoipGetTokenResponse {
	s.Headers = v
	return s
}

func (s *VoipGetTokenResponse) SetBody(v *VoipGetTokenResponseBody) *VoipGetTokenResponse {
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
	client.EndpointRule = tea.String("central")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("dyvmsapi"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) AddRtcAccountWithOptions(request *AddRtcAccountRequest, runtime *util.RuntimeOptions) (_result *AddRtcAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddRtcAccountResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddRtcAccount"), tea.String("2017-05-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddRtcAccount(request *AddRtcAccountRequest) (_result *AddRtcAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddRtcAccountResponse{}
	_body, _err := client.AddRtcAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchRobotSmartCallWithOptions(request *BatchRobotSmartCallRequest, runtime *util.RuntimeOptions) (_result *BatchRobotSmartCallResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &BatchRobotSmartCallResponse{}
	_body, _err := client.DoRPCRequest(tea.String("BatchRobotSmartCall"), tea.String("2017-05-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchRobotSmartCall(request *BatchRobotSmartCallRequest) (_result *BatchRobotSmartCallResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchRobotSmartCallResponse{}
	_body, _err := client.BatchRobotSmartCallWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BindNumberAndVoipIdWithOptions(request *BindNumberAndVoipIdRequest, runtime *util.RuntimeOptions) (_result *BindNumberAndVoipIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &BindNumberAndVoipIdResponse{}
	_body, _err := client.DoRPCRequest(tea.String("BindNumberAndVoipId"), tea.String("2017-05-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BindNumberAndVoipId(request *BindNumberAndVoipIdRequest) (_result *BindNumberAndVoipIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BindNumberAndVoipIdResponse{}
	_body, _err := client.BindNumberAndVoipIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CancelCallWithOptions(request *CancelCallRequest, runtime *util.RuntimeOptions) (_result *CancelCallResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CancelCallResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CancelCall"), tea.String("2017-05-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CancelCall(request *CancelCallRequest) (_result *CancelCallResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CancelCallResponse{}
	_body, _err := client.CancelCallWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CancelOrderRobotTaskWithOptions(request *CancelOrderRobotTaskRequest, runtime *util.RuntimeOptions) (_result *CancelOrderRobotTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CancelOrderRobotTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CancelOrderRobotTask"), tea.String("2017-05-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CancelOrderRobotTask(request *CancelOrderRobotTaskRequest) (_result *CancelOrderRobotTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CancelOrderRobotTaskResponse{}
	_body, _err := client.CancelOrderRobotTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CancelRobotTaskWithOptions(request *CancelRobotTaskRequest, runtime *util.RuntimeOptions) (_result *CancelRobotTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CancelRobotTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CancelRobotTask"), tea.String("2017-05-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CancelRobotTask(request *CancelRobotTaskRequest) (_result *CancelRobotTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CancelRobotTaskResponse{}
	_body, _err := client.CancelRobotTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ClickToDialWithOptions(request *ClickToDialRequest, runtime *util.RuntimeOptions) (_result *ClickToDialResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ClickToDialResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ClickToDial"), tea.String("2017-05-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ClickToDial(request *ClickToDialRequest) (_result *ClickToDialResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ClickToDialResponse{}
	_body, _err := client.ClickToDialWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateRobotTaskWithOptions(request *CreateRobotTaskRequest, runtime *util.RuntimeOptions) (_result *CreateRobotTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateRobotTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateRobotTask"), tea.String("2017-05-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateRobotTask(request *CreateRobotTaskRequest) (_result *CreateRobotTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateRobotTaskResponse{}
	_body, _err := client.CreateRobotTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteRobotTaskWithOptions(request *DeleteRobotTaskRequest, runtime *util.RuntimeOptions) (_result *DeleteRobotTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteRobotTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteRobotTask"), tea.String("2017-05-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteRobotTask(request *DeleteRobotTaskRequest) (_result *DeleteRobotTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteRobotTaskResponse{}
	_body, _err := client.DeleteRobotTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRecordDataWithOptions(request *DescribeRecordDataRequest, runtime *util.RuntimeOptions) (_result *DescribeRecordDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeRecordDataResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeRecordData"), tea.String("2017-05-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRecordData(request *DescribeRecordDataRequest) (_result *DescribeRecordDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRecordDataResponse{}
	_body, _err := client.DescribeRecordDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DoRtcNumberAuthWithOptions(request *DoRtcNumberAuthRequest, runtime *util.RuntimeOptions) (_result *DoRtcNumberAuthResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DoRtcNumberAuthResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DoRtcNumberAuth"), tea.String("2017-05-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DoRtcNumberAuth(request *DoRtcNumberAuthRequest) (_result *DoRtcNumberAuthResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DoRtcNumberAuthResponse{}
	_body, _err := client.DoRtcNumberAuthWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetRtcTokenWithOptions(request *GetRtcTokenRequest, runtime *util.RuntimeOptions) (_result *GetRtcTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetRtcTokenResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetRtcToken"), tea.String("2017-05-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetRtcToken(request *GetRtcTokenRequest) (_result *GetRtcTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetRtcTokenResponse{}
	_body, _err := client.GetRtcTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) IvrCallWithOptions(request *IvrCallRequest, runtime *util.RuntimeOptions) (_result *IvrCallResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &IvrCallResponse{}
	_body, _err := client.DoRPCRequest(tea.String("IvrCall"), tea.String("2017-05-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) IvrCall(request *IvrCallRequest) (_result *IvrCallResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &IvrCallResponse{}
	_body, _err := client.IvrCallWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListOrderedNumbersWithOptions(request *ListOrderedNumbersRequest, runtime *util.RuntimeOptions) (_result *ListOrderedNumbersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListOrderedNumbersResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListOrderedNumbers"), tea.String("2017-05-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListOrderedNumbers(request *ListOrderedNumbersRequest) (_result *ListOrderedNumbersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListOrderedNumbersResponse{}
	_body, _err := client.ListOrderedNumbersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListOutboundStrategiesWithOptions(request *ListOutboundStrategiesRequest, runtime *util.RuntimeOptions) (_result *ListOutboundStrategiesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListOutboundStrategiesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListOutboundStrategies"), tea.String("2017-05-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListOutboundStrategies(request *ListOutboundStrategiesRequest) (_result *ListOutboundStrategiesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListOutboundStrategiesResponse{}
	_body, _err := client.ListOutboundStrategiesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListRobotTaskCallsWithOptions(request *ListRobotTaskCallsRequest, runtime *util.RuntimeOptions) (_result *ListRobotTaskCallsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListRobotTaskCallsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListRobotTaskCalls"), tea.String("2017-05-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListRobotTaskCalls(request *ListRobotTaskCallsRequest) (_result *ListRobotTaskCallsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListRobotTaskCallsResponse{}
	_body, _err := client.ListRobotTaskCallsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryCallDetailByCallIdWithOptions(request *QueryCallDetailByCallIdRequest, runtime *util.RuntimeOptions) (_result *QueryCallDetailByCallIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryCallDetailByCallIdResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryCallDetailByCallId"), tea.String("2017-05-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryCallDetailByCallId(request *QueryCallDetailByCallIdRequest) (_result *QueryCallDetailByCallIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryCallDetailByCallIdResponse{}
	_body, _err := client.QueryCallDetailByCallIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryCallDetailByTaskIdWithOptions(request *QueryCallDetailByTaskIdRequest, runtime *util.RuntimeOptions) (_result *QueryCallDetailByTaskIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryCallDetailByTaskIdResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryCallDetailByTaskId"), tea.String("2017-05-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryCallDetailByTaskId(request *QueryCallDetailByTaskIdRequest) (_result *QueryCallDetailByTaskIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryCallDetailByTaskIdResponse{}
	_body, _err := client.QueryCallDetailByTaskIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryRobotInfoListWithOptions(request *QueryRobotInfoListRequest, runtime *util.RuntimeOptions) (_result *QueryRobotInfoListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryRobotInfoListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryRobotInfoList"), tea.String("2017-05-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryRobotInfoList(request *QueryRobotInfoListRequest) (_result *QueryRobotInfoListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryRobotInfoListResponse{}
	_body, _err := client.QueryRobotInfoListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryRobotTaskCallDetailWithOptions(request *QueryRobotTaskCallDetailRequest, runtime *util.RuntimeOptions) (_result *QueryRobotTaskCallDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryRobotTaskCallDetailResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryRobotTaskCallDetail"), tea.String("2017-05-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryRobotTaskCallDetail(request *QueryRobotTaskCallDetailRequest) (_result *QueryRobotTaskCallDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryRobotTaskCallDetailResponse{}
	_body, _err := client.QueryRobotTaskCallDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryRobotTaskCallListWithOptions(request *QueryRobotTaskCallListRequest, runtime *util.RuntimeOptions) (_result *QueryRobotTaskCallListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryRobotTaskCallListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryRobotTaskCallList"), tea.String("2017-05-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryRobotTaskCallList(request *QueryRobotTaskCallListRequest) (_result *QueryRobotTaskCallListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryRobotTaskCallListResponse{}
	_body, _err := client.QueryRobotTaskCallListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryRobotTaskDetailWithOptions(request *QueryRobotTaskDetailRequest, runtime *util.RuntimeOptions) (_result *QueryRobotTaskDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryRobotTaskDetailResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryRobotTaskDetail"), tea.String("2017-05-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryRobotTaskDetail(request *QueryRobotTaskDetailRequest) (_result *QueryRobotTaskDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryRobotTaskDetailResponse{}
	_body, _err := client.QueryRobotTaskDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryRobotTaskListWithOptions(request *QueryRobotTaskListRequest, runtime *util.RuntimeOptions) (_result *QueryRobotTaskListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryRobotTaskListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryRobotTaskList"), tea.String("2017-05-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryRobotTaskList(request *QueryRobotTaskListRequest) (_result *QueryRobotTaskListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryRobotTaskListResponse{}
	_body, _err := client.QueryRobotTaskListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryRobotv2AllListWithOptions(request *QueryRobotv2AllListRequest, runtime *util.RuntimeOptions) (_result *QueryRobotv2AllListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryRobotv2AllListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryRobotv2AllList"), tea.String("2017-05-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryRobotv2AllList(request *QueryRobotv2AllListRequest) (_result *QueryRobotv2AllListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryRobotv2AllListResponse{}
	_body, _err := client.QueryRobotv2AllListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryRtcNumberAuthStatusWithOptions(request *QueryRtcNumberAuthStatusRequest, runtime *util.RuntimeOptions) (_result *QueryRtcNumberAuthStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryRtcNumberAuthStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryRtcNumberAuthStatus"), tea.String("2017-05-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryRtcNumberAuthStatus(request *QueryRtcNumberAuthStatusRequest) (_result *QueryRtcNumberAuthStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryRtcNumberAuthStatusResponse{}
	_body, _err := client.QueryRtcNumberAuthStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryVoipNumberBindInfosWithOptions(request *QueryVoipNumberBindInfosRequest, runtime *util.RuntimeOptions) (_result *QueryVoipNumberBindInfosResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &QueryVoipNumberBindInfosResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryVoipNumberBindInfos"), tea.String("2017-05-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryVoipNumberBindInfos(request *QueryVoipNumberBindInfosRequest) (_result *QueryVoipNumberBindInfosResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryVoipNumberBindInfosResponse{}
	_body, _err := client.QueryVoipNumberBindInfosWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ReportVoipProblemsWithOptions(request *ReportVoipProblemsRequest, runtime *util.RuntimeOptions) (_result *ReportVoipProblemsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ReportVoipProblemsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ReportVoipProblems"), tea.String("2017-05-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReportVoipProblems(request *ReportVoipProblemsRequest) (_result *ReportVoipProblemsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReportVoipProblemsResponse{}
	_body, _err := client.ReportVoipProblemsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SingleCallByTtsWithOptions(request *SingleCallByTtsRequest, runtime *util.RuntimeOptions) (_result *SingleCallByTtsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SingleCallByTtsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SingleCallByTts"), tea.String("2017-05-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SingleCallByTts(request *SingleCallByTtsRequest) (_result *SingleCallByTtsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SingleCallByTtsResponse{}
	_body, _err := client.SingleCallByTtsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SingleCallByVoiceWithOptions(request *SingleCallByVoiceRequest, runtime *util.RuntimeOptions) (_result *SingleCallByVoiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SingleCallByVoiceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SingleCallByVoice"), tea.String("2017-05-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SingleCallByVoice(request *SingleCallByVoiceRequest) (_result *SingleCallByVoiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SingleCallByVoiceResponse{}
	_body, _err := client.SingleCallByVoiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SmartCallWithOptions(request *SmartCallRequest, runtime *util.RuntimeOptions) (_result *SmartCallResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SmartCallResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SmartCall"), tea.String("2017-05-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SmartCall(request *SmartCallRequest) (_result *SmartCallResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SmartCallResponse{}
	_body, _err := client.SmartCallWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SmartCallOperateWithOptions(request *SmartCallOperateRequest, runtime *util.RuntimeOptions) (_result *SmartCallOperateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SmartCallOperateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SmartCallOperate"), tea.String("2017-05-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SmartCallOperate(request *SmartCallOperateRequest) (_result *SmartCallOperateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SmartCallOperateResponse{}
	_body, _err := client.SmartCallOperateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartMicroOutboundWithOptions(request *StartMicroOutboundRequest, runtime *util.RuntimeOptions) (_result *StartMicroOutboundResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &StartMicroOutboundResponse{}
	_body, _err := client.DoRPCRequest(tea.String("StartMicroOutbound"), tea.String("2017-05-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartMicroOutbound(request *StartMicroOutboundRequest) (_result *StartMicroOutboundResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartMicroOutboundResponse{}
	_body, _err := client.StartMicroOutboundWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartRobotTaskWithOptions(request *StartRobotTaskRequest, runtime *util.RuntimeOptions) (_result *StartRobotTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &StartRobotTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("StartRobotTask"), tea.String("2017-05-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartRobotTask(request *StartRobotTaskRequest) (_result *StartRobotTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartRobotTaskResponse{}
	_body, _err := client.StartRobotTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopRobotTaskWithOptions(request *StopRobotTaskRequest, runtime *util.RuntimeOptions) (_result *StopRobotTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &StopRobotTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("StopRobotTask"), tea.String("2017-05-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopRobotTask(request *StopRobotTaskRequest) (_result *StopRobotTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopRobotTaskResponse{}
	_body, _err := client.StopRobotTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UnbindNumberAndVoipIdWithOptions(request *UnbindNumberAndVoipIdRequest, runtime *util.RuntimeOptions) (_result *UnbindNumberAndVoipIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UnbindNumberAndVoipIdResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UnbindNumberAndVoipId"), tea.String("2017-05-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UnbindNumberAndVoipId(request *UnbindNumberAndVoipIdRequest) (_result *UnbindNumberAndVoipIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UnbindNumberAndVoipIdResponse{}
	_body, _err := client.UnbindNumberAndVoipIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UndoRtcNumberAuthWithOptions(request *UndoRtcNumberAuthRequest, runtime *util.RuntimeOptions) (_result *UndoRtcNumberAuthResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UndoRtcNumberAuthResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UndoRtcNumberAuth"), tea.String("2017-05-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UndoRtcNumberAuth(request *UndoRtcNumberAuthRequest) (_result *UndoRtcNumberAuthResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UndoRtcNumberAuthResponse{}
	_body, _err := client.UndoRtcNumberAuthWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UploadRobotTaskCalledFileWithOptions(request *UploadRobotTaskCalledFileRequest, runtime *util.RuntimeOptions) (_result *UploadRobotTaskCalledFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UploadRobotTaskCalledFileResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UploadRobotTaskCalledFile"), tea.String("2017-05-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UploadRobotTaskCalledFile(request *UploadRobotTaskCalledFileRequest) (_result *UploadRobotTaskCalledFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UploadRobotTaskCalledFileResponse{}
	_body, _err := client.UploadRobotTaskCalledFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) VoipAddAccountWithOptions(request *VoipAddAccountRequest, runtime *util.RuntimeOptions) (_result *VoipAddAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &VoipAddAccountResponse{}
	_body, _err := client.DoRPCRequest(tea.String("VoipAddAccount"), tea.String("2017-05-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) VoipAddAccount(request *VoipAddAccountRequest) (_result *VoipAddAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &VoipAddAccountResponse{}
	_body, _err := client.VoipAddAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) VoipGetTokenWithOptions(request *VoipGetTokenRequest, runtime *util.RuntimeOptions) (_result *VoipGetTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &VoipGetTokenResponse{}
	_body, _err := client.DoRPCRequest(tea.String("VoipGetToken"), tea.String("2017-05-25"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) VoipGetToken(request *VoipGetTokenRequest) (_result *VoipGetTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &VoipGetTokenResponse{}
	_body, _err := client.VoipGetTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
