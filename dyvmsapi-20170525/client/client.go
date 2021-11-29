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

type AddRtcAccountRequest struct {
	DeviceId             *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s AddRtcAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s AddRtcAccountRequest) GoString() string {
	return s.String()
}

func (s *AddRtcAccountRequest) SetDeviceId(v string) *AddRtcAccountRequest {
	s.DeviceId = &v
	return s
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

type AddRtcAccountResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Module    *string `json:"Module,omitempty" xml:"Module,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddRtcAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddRtcAccountResponseBody) GoString() string {
	return s.String()
}

func (s *AddRtcAccountResponseBody) SetCode(v string) *AddRtcAccountResponseBody {
	s.Code = &v
	return s
}

func (s *AddRtcAccountResponseBody) SetMessage(v string) *AddRtcAccountResponseBody {
	s.Message = &v
	return s
}

func (s *AddRtcAccountResponseBody) SetModule(v string) *AddRtcAccountResponseBody {
	s.Module = &v
	return s
}

func (s *AddRtcAccountResponseBody) SetRequestId(v string) *AddRtcAccountResponseBody {
	s.RequestId = &v
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

type AddVirtualNumberRelationRequest struct {
	CorpNameList         *string `json:"CorpNameList,omitempty" xml:"CorpNameList,omitempty"`
	NumberList           *string `json:"NumberList,omitempty" xml:"NumberList,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PhoneNum             *string `json:"PhoneNum,omitempty" xml:"PhoneNum,omitempty"`
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RouteType            *int32  `json:"RouteType,omitempty" xml:"RouteType,omitempty"`
}

func (s AddVirtualNumberRelationRequest) String() string {
	return tea.Prettify(s)
}

func (s AddVirtualNumberRelationRequest) GoString() string {
	return s.String()
}

func (s *AddVirtualNumberRelationRequest) SetCorpNameList(v string) *AddVirtualNumberRelationRequest {
	s.CorpNameList = &v
	return s
}

func (s *AddVirtualNumberRelationRequest) SetNumberList(v string) *AddVirtualNumberRelationRequest {
	s.NumberList = &v
	return s
}

func (s *AddVirtualNumberRelationRequest) SetOwnerId(v int64) *AddVirtualNumberRelationRequest {
	s.OwnerId = &v
	return s
}

func (s *AddVirtualNumberRelationRequest) SetPhoneNum(v string) *AddVirtualNumberRelationRequest {
	s.PhoneNum = &v
	return s
}

func (s *AddVirtualNumberRelationRequest) SetProdCode(v string) *AddVirtualNumberRelationRequest {
	s.ProdCode = &v
	return s
}

func (s *AddVirtualNumberRelationRequest) SetResourceOwnerAccount(v string) *AddVirtualNumberRelationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AddVirtualNumberRelationRequest) SetResourceOwnerId(v int64) *AddVirtualNumberRelationRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AddVirtualNumberRelationRequest) SetRouteType(v int32) *AddVirtualNumberRelationRequest {
	s.RouteType = &v
	return s
}

type AddVirtualNumberRelationResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddVirtualNumberRelationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddVirtualNumberRelationResponseBody) GoString() string {
	return s.String()
}

func (s *AddVirtualNumberRelationResponseBody) SetCode(v string) *AddVirtualNumberRelationResponseBody {
	s.Code = &v
	return s
}

func (s *AddVirtualNumberRelationResponseBody) SetData(v string) *AddVirtualNumberRelationResponseBody {
	s.Data = &v
	return s
}

func (s *AddVirtualNumberRelationResponseBody) SetMessage(v string) *AddVirtualNumberRelationResponseBody {
	s.Message = &v
	return s
}

func (s *AddVirtualNumberRelationResponseBody) SetRequestId(v string) *AddVirtualNumberRelationResponseBody {
	s.RequestId = &v
	return s
}

type AddVirtualNumberRelationResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddVirtualNumberRelationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddVirtualNumberRelationResponse) String() string {
	return tea.Prettify(s)
}

func (s AddVirtualNumberRelationResponse) GoString() string {
	return s.String()
}

func (s *AddVirtualNumberRelationResponse) SetHeaders(v map[string]*string) *AddVirtualNumberRelationResponse {
	s.Headers = v
	return s
}

func (s *AddVirtualNumberRelationResponse) SetBody(v *AddVirtualNumberRelationResponseBody) *AddVirtualNumberRelationResponse {
	s.Body = v
	return s
}

type BatchRobotSmartCallRequest struct {
	CalledNumber         *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	CalledShowNumber     *string `json:"CalledShowNumber,omitempty" xml:"CalledShowNumber,omitempty"`
	CorpName             *string `json:"CorpName,omitempty" xml:"CorpName,omitempty"`
	DialogId             *string `json:"DialogId,omitempty" xml:"DialogId,omitempty"`
	EarlyMediaAsr        *bool   `json:"EarlyMediaAsr,omitempty" xml:"EarlyMediaAsr,omitempty"`
	IsSelfLine           *bool   `json:"IsSelfLine,omitempty" xml:"IsSelfLine,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ScheduleCall         *bool   `json:"ScheduleCall,omitempty" xml:"ScheduleCall,omitempty"`
	ScheduleTime         *int64  `json:"ScheduleTime,omitempty" xml:"ScheduleTime,omitempty"`
	TaskName             *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	TtsParam             *string `json:"TtsParam,omitempty" xml:"TtsParam,omitempty"`
	TtsParamHead         *string `json:"TtsParamHead,omitempty" xml:"TtsParamHead,omitempty"`
}

func (s BatchRobotSmartCallRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchRobotSmartCallRequest) GoString() string {
	return s.String()
}

func (s *BatchRobotSmartCallRequest) SetCalledNumber(v string) *BatchRobotSmartCallRequest {
	s.CalledNumber = &v
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

func (s *BatchRobotSmartCallRequest) SetDialogId(v string) *BatchRobotSmartCallRequest {
	s.DialogId = &v
	return s
}

func (s *BatchRobotSmartCallRequest) SetEarlyMediaAsr(v bool) *BatchRobotSmartCallRequest {
	s.EarlyMediaAsr = &v
	return s
}

func (s *BatchRobotSmartCallRequest) SetIsSelfLine(v bool) *BatchRobotSmartCallRequest {
	s.IsSelfLine = &v
	return s
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

func (s *BatchRobotSmartCallRequest) SetScheduleCall(v bool) *BatchRobotSmartCallRequest {
	s.ScheduleCall = &v
	return s
}

func (s *BatchRobotSmartCallRequest) SetScheduleTime(v int64) *BatchRobotSmartCallRequest {
	s.ScheduleTime = &v
	return s
}

func (s *BatchRobotSmartCallRequest) SetTaskName(v string) *BatchRobotSmartCallRequest {
	s.TaskName = &v
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

type BatchRobotSmartCallResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s BatchRobotSmartCallResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchRobotSmartCallResponseBody) GoString() string {
	return s.String()
}

func (s *BatchRobotSmartCallResponseBody) SetCode(v string) *BatchRobotSmartCallResponseBody {
	s.Code = &v
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

func (s *BatchRobotSmartCallResponseBody) SetTaskId(v string) *BatchRobotSmartCallResponseBody {
	s.TaskId = &v
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

type CancelCallRequest struct {
	CallId               *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CancelCallRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelCallRequest) GoString() string {
	return s.String()
}

func (s *CancelCallRequest) SetCallId(v string) *CancelCallRequest {
	s.CallId = &v
	return s
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

type CancelCallResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status    *bool   `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CancelCallResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelCallResponseBody) GoString() string {
	return s.String()
}

func (s *CancelCallResponseBody) SetCode(v string) *CancelCallResponseBody {
	s.Code = &v
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

func (s *CancelCallResponseBody) SetStatus(v bool) *CancelCallResponseBody {
	s.Status = &v
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
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelOrderRobotTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelOrderRobotTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CancelOrderRobotTaskResponseBody) SetCode(v string) *CancelOrderRobotTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CancelOrderRobotTaskResponseBody) SetData(v string) *CancelOrderRobotTaskResponseBody {
	s.Data = &v
	return s
}

func (s *CancelOrderRobotTaskResponseBody) SetMessage(v string) *CancelOrderRobotTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CancelOrderRobotTaskResponseBody) SetRequestId(v string) *CancelOrderRobotTaskResponseBody {
	s.RequestId = &v
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
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelRobotTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelRobotTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CancelRobotTaskResponseBody) SetCode(v string) *CancelRobotTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CancelRobotTaskResponseBody) SetData(v string) *CancelRobotTaskResponseBody {
	s.Data = &v
	return s
}

func (s *CancelRobotTaskResponseBody) SetMessage(v string) *CancelRobotTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CancelRobotTaskResponseBody) SetRequestId(v string) *CancelRobotTaskResponseBody {
	s.RequestId = &v
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
	AsrFlag              *bool   `json:"AsrFlag,omitempty" xml:"AsrFlag,omitempty"`
	AsrModelId           *string `json:"AsrModelId,omitempty" xml:"AsrModelId,omitempty"`
	CalledNumber         *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	CalledShowNumber     *string `json:"CalledShowNumber,omitempty" xml:"CalledShowNumber,omitempty"`
	CallerNumber         *string `json:"CallerNumber,omitempty" xml:"CallerNumber,omitempty"`
	CallerShowNumber     *string `json:"CallerShowNumber,omitempty" xml:"CallerShowNumber,omitempty"`
	OutId                *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RecordFlag           *bool   `json:"RecordFlag,omitempty" xml:"RecordFlag,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SessionTimeout       *int32  `json:"SessionTimeout,omitempty" xml:"SessionTimeout,omitempty"`
}

func (s ClickToDialRequest) String() string {
	return tea.Prettify(s)
}

func (s ClickToDialRequest) GoString() string {
	return s.String()
}

func (s *ClickToDialRequest) SetAsrFlag(v bool) *ClickToDialRequest {
	s.AsrFlag = &v
	return s
}

func (s *ClickToDialRequest) SetAsrModelId(v string) *ClickToDialRequest {
	s.AsrModelId = &v
	return s
}

func (s *ClickToDialRequest) SetCalledNumber(v string) *ClickToDialRequest {
	s.CalledNumber = &v
	return s
}

func (s *ClickToDialRequest) SetCalledShowNumber(v string) *ClickToDialRequest {
	s.CalledShowNumber = &v
	return s
}

func (s *ClickToDialRequest) SetCallerNumber(v string) *ClickToDialRequest {
	s.CallerNumber = &v
	return s
}

func (s *ClickToDialRequest) SetCallerShowNumber(v string) *ClickToDialRequest {
	s.CallerShowNumber = &v
	return s
}

func (s *ClickToDialRequest) SetOutId(v string) *ClickToDialRequest {
	s.OutId = &v
	return s
}

func (s *ClickToDialRequest) SetOwnerId(v int64) *ClickToDialRequest {
	s.OwnerId = &v
	return s
}

func (s *ClickToDialRequest) SetRecordFlag(v bool) *ClickToDialRequest {
	s.RecordFlag = &v
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

func (s *ClickToDialRequest) SetSessionTimeout(v int32) *ClickToDialRequest {
	s.SessionTimeout = &v
	return s
}

type ClickToDialResponseBody struct {
	CallId    *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ClickToDialResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ClickToDialResponseBody) GoString() string {
	return s.String()
}

func (s *ClickToDialResponseBody) SetCallId(v string) *ClickToDialResponseBody {
	s.CallId = &v
	return s
}

func (s *ClickToDialResponseBody) SetCode(v string) *ClickToDialResponseBody {
	s.Code = &v
	return s
}

func (s *ClickToDialResponseBody) SetMessage(v string) *ClickToDialResponseBody {
	s.Message = &v
	return s
}

func (s *ClickToDialResponseBody) SetRequestId(v string) *ClickToDialResponseBody {
	s.RequestId = &v
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

type CreateCallTaskRequest struct {
	BizType              *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	Data                 *string `json:"Data,omitempty" xml:"Data,omitempty"`
	DataType             *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	FireTime             *string `json:"FireTime,omitempty" xml:"FireTime,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Resource             *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ResourceType         *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	ScheduleType         *string `json:"ScheduleType,omitempty" xml:"ScheduleType,omitempty"`
	StopTime             *string `json:"StopTime,omitempty" xml:"StopTime,omitempty"`
	TaskName             *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	TemplateCode         *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	TemplateName         *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s CreateCallTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCallTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateCallTaskRequest) SetBizType(v string) *CreateCallTaskRequest {
	s.BizType = &v
	return s
}

func (s *CreateCallTaskRequest) SetData(v string) *CreateCallTaskRequest {
	s.Data = &v
	return s
}

func (s *CreateCallTaskRequest) SetDataType(v string) *CreateCallTaskRequest {
	s.DataType = &v
	return s
}

func (s *CreateCallTaskRequest) SetFireTime(v string) *CreateCallTaskRequest {
	s.FireTime = &v
	return s
}

func (s *CreateCallTaskRequest) SetOwnerId(v int64) *CreateCallTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateCallTaskRequest) SetResource(v string) *CreateCallTaskRequest {
	s.Resource = &v
	return s
}

func (s *CreateCallTaskRequest) SetResourceOwnerAccount(v string) *CreateCallTaskRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateCallTaskRequest) SetResourceOwnerId(v int64) *CreateCallTaskRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateCallTaskRequest) SetResourceType(v string) *CreateCallTaskRequest {
	s.ResourceType = &v
	return s
}

func (s *CreateCallTaskRequest) SetScheduleType(v string) *CreateCallTaskRequest {
	s.ScheduleType = &v
	return s
}

func (s *CreateCallTaskRequest) SetStopTime(v string) *CreateCallTaskRequest {
	s.StopTime = &v
	return s
}

func (s *CreateCallTaskRequest) SetTaskName(v string) *CreateCallTaskRequest {
	s.TaskName = &v
	return s
}

func (s *CreateCallTaskRequest) SetTemplateCode(v string) *CreateCallTaskRequest {
	s.TemplateCode = &v
	return s
}

func (s *CreateCallTaskRequest) SetTemplateName(v string) *CreateCallTaskRequest {
	s.TemplateName = &v
	return s
}

type CreateCallTaskResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *int64  `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCallTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateCallTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCallTaskResponseBody) SetCode(v string) *CreateCallTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CreateCallTaskResponseBody) SetData(v int64) *CreateCallTaskResponseBody {
	s.Data = &v
	return s
}

func (s *CreateCallTaskResponseBody) SetRequestId(v string) *CreateCallTaskResponseBody {
	s.RequestId = &v
	return s
}

type CreateCallTaskResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateCallTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateCallTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCallTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateCallTaskResponse) SetHeaders(v map[string]*string) *CreateCallTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateCallTaskResponse) SetBody(v *CreateCallTaskResponseBody) *CreateCallTaskResponse {
	s.Body = v
	return s
}

type CreateRobotTaskRequest struct {
	Caller               *string `json:"Caller,omitempty" xml:"Caller,omitempty"`
	CorpName             *string `json:"CorpName,omitempty" xml:"CorpName,omitempty"`
	DialogId             *int64  `json:"DialogId,omitempty" xml:"DialogId,omitempty"`
	IsSelfLine           *bool   `json:"IsSelfLine,omitempty" xml:"IsSelfLine,omitempty"`
	NumberStatusIdent    *bool   `json:"NumberStatusIdent,omitempty" xml:"NumberStatusIdent,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RecallInterval       *int32  `json:"RecallInterval,omitempty" xml:"RecallInterval,omitempty"`
	RecallStateCodes     *string `json:"RecallStateCodes,omitempty" xml:"RecallStateCodes,omitempty"`
	RecallTimes          *int32  `json:"RecallTimes,omitempty" xml:"RecallTimes,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RetryType            *int32  `json:"RetryType,omitempty" xml:"RetryType,omitempty"`
	TaskName             *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s CreateRobotTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRobotTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateRobotTaskRequest) SetCaller(v string) *CreateRobotTaskRequest {
	s.Caller = &v
	return s
}

func (s *CreateRobotTaskRequest) SetCorpName(v string) *CreateRobotTaskRequest {
	s.CorpName = &v
	return s
}

func (s *CreateRobotTaskRequest) SetDialogId(v int64) *CreateRobotTaskRequest {
	s.DialogId = &v
	return s
}

func (s *CreateRobotTaskRequest) SetIsSelfLine(v bool) *CreateRobotTaskRequest {
	s.IsSelfLine = &v
	return s
}

func (s *CreateRobotTaskRequest) SetNumberStatusIdent(v bool) *CreateRobotTaskRequest {
	s.NumberStatusIdent = &v
	return s
}

func (s *CreateRobotTaskRequest) SetOwnerId(v int64) *CreateRobotTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateRobotTaskRequest) SetRecallInterval(v int32) *CreateRobotTaskRequest {
	s.RecallInterval = &v
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

func (s *CreateRobotTaskRequest) SetResourceOwnerAccount(v string) *CreateRobotTaskRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateRobotTaskRequest) SetResourceOwnerId(v int64) *CreateRobotTaskRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateRobotTaskRequest) SetRetryType(v int32) *CreateRobotTaskRequest {
	s.RetryType = &v
	return s
}

func (s *CreateRobotTaskRequest) SetTaskName(v string) *CreateRobotTaskRequest {
	s.TaskName = &v
	return s
}

type CreateRobotTaskResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateRobotTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateRobotTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRobotTaskResponseBody) SetCode(v string) *CreateRobotTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CreateRobotTaskResponseBody) SetData(v string) *CreateRobotTaskResponseBody {
	s.Data = &v
	return s
}

func (s *CreateRobotTaskResponseBody) SetMessage(v string) *CreateRobotTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CreateRobotTaskResponseBody) SetRequestId(v string) *CreateRobotTaskResponseBody {
	s.RequestId = &v
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
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteRobotTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteRobotTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRobotTaskResponseBody) SetCode(v string) *DeleteRobotTaskResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteRobotTaskResponseBody) SetData(v string) *DeleteRobotTaskResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteRobotTaskResponseBody) SetMessage(v string) *DeleteRobotTaskResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteRobotTaskResponseBody) SetRequestId(v string) *DeleteRobotTaskResponseBody {
	s.RequestId = &v
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

type ExecuteCallTaskRequest struct {
	FireTime             *string `json:"FireTime,omitempty" xml:"FireTime,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Status               *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TaskId               *int64  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ExecuteCallTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s ExecuteCallTaskRequest) GoString() string {
	return s.String()
}

func (s *ExecuteCallTaskRequest) SetFireTime(v string) *ExecuteCallTaskRequest {
	s.FireTime = &v
	return s
}

func (s *ExecuteCallTaskRequest) SetOwnerId(v int64) *ExecuteCallTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *ExecuteCallTaskRequest) SetResourceOwnerAccount(v string) *ExecuteCallTaskRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ExecuteCallTaskRequest) SetResourceOwnerId(v int64) *ExecuteCallTaskRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ExecuteCallTaskRequest) SetStatus(v string) *ExecuteCallTaskRequest {
	s.Status = &v
	return s
}

func (s *ExecuteCallTaskRequest) SetTaskId(v int64) *ExecuteCallTaskRequest {
	s.TaskId = &v
	return s
}

type ExecuteCallTaskResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ExecuteCallTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExecuteCallTaskResponseBody) GoString() string {
	return s.String()
}

func (s *ExecuteCallTaskResponseBody) SetCode(v string) *ExecuteCallTaskResponseBody {
	s.Code = &v
	return s
}

func (s *ExecuteCallTaskResponseBody) SetData(v bool) *ExecuteCallTaskResponseBody {
	s.Data = &v
	return s
}

func (s *ExecuteCallTaskResponseBody) SetRequestId(v string) *ExecuteCallTaskResponseBody {
	s.RequestId = &v
	return s
}

type ExecuteCallTaskResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ExecuteCallTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ExecuteCallTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s ExecuteCallTaskResponse) GoString() string {
	return s.String()
}

func (s *ExecuteCallTaskResponse) SetHeaders(v map[string]*string) *ExecuteCallTaskResponse {
	s.Headers = v
	return s
}

func (s *ExecuteCallTaskResponse) SetBody(v *ExecuteCallTaskResponseBody) *ExecuteCallTaskResponse {
	s.Body = v
	return s
}

type GetCallInfoRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RtcId                *string `json:"RtcId,omitempty" xml:"RtcId,omitempty"`
}

func (s GetCallInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCallInfoRequest) GoString() string {
	return s.String()
}

func (s *GetCallInfoRequest) SetOwnerId(v int64) *GetCallInfoRequest {
	s.OwnerId = &v
	return s
}

func (s *GetCallInfoRequest) SetResourceOwnerAccount(v string) *GetCallInfoRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetCallInfoRequest) SetResourceOwnerId(v int64) *GetCallInfoRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetCallInfoRequest) SetRtcId(v string) *GetCallInfoRequest {
	s.RtcId = &v
	return s
}

type GetCallInfoResponseBody struct {
	Code      *string                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetCallInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetCallInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCallInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetCallInfoResponseBody) SetCode(v string) *GetCallInfoResponseBody {
	s.Code = &v
	return s
}

func (s *GetCallInfoResponseBody) SetData(v *GetCallInfoResponseBodyData) *GetCallInfoResponseBody {
	s.Data = v
	return s
}

func (s *GetCallInfoResponseBody) SetMessage(v string) *GetCallInfoResponseBody {
	s.Message = &v
	return s
}

func (s *GetCallInfoResponseBody) SetRequestId(v string) *GetCallInfoResponseBody {
	s.RequestId = &v
	return s
}

type GetCallInfoResponseBodyData struct {
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetCallInfoResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetCallInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCallInfoResponseBodyData) SetChannelId(v string) *GetCallInfoResponseBodyData {
	s.ChannelId = &v
	return s
}

func (s *GetCallInfoResponseBodyData) SetStatus(v string) *GetCallInfoResponseBodyData {
	s.Status = &v
	return s
}

type GetCallInfoResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetCallInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetCallInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCallInfoResponse) GoString() string {
	return s.String()
}

func (s *GetCallInfoResponse) SetHeaders(v map[string]*string) *GetCallInfoResponse {
	s.Headers = v
	return s
}

func (s *GetCallInfoResponse) SetBody(v *GetCallInfoResponseBody) *GetCallInfoResponse {
	s.Body = v
	return s
}

type GetHotlineQualificationByOrderRequest struct {
	OrderId              *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s GetHotlineQualificationByOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s GetHotlineQualificationByOrderRequest) GoString() string {
	return s.String()
}

func (s *GetHotlineQualificationByOrderRequest) SetOrderId(v string) *GetHotlineQualificationByOrderRequest {
	s.OrderId = &v
	return s
}

func (s *GetHotlineQualificationByOrderRequest) SetOwnerId(v int64) *GetHotlineQualificationByOrderRequest {
	s.OwnerId = &v
	return s
}

func (s *GetHotlineQualificationByOrderRequest) SetResourceOwnerAccount(v string) *GetHotlineQualificationByOrderRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetHotlineQualificationByOrderRequest) SetResourceOwnerId(v int64) *GetHotlineQualificationByOrderRequest {
	s.ResourceOwnerId = &v
	return s
}

type GetHotlineQualificationByOrderResponseBody struct {
	Code      *string                                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetHotlineQualificationByOrderResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetHotlineQualificationByOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetHotlineQualificationByOrderResponseBody) GoString() string {
	return s.String()
}

func (s *GetHotlineQualificationByOrderResponseBody) SetCode(v string) *GetHotlineQualificationByOrderResponseBody {
	s.Code = &v
	return s
}

func (s *GetHotlineQualificationByOrderResponseBody) SetData(v *GetHotlineQualificationByOrderResponseBodyData) *GetHotlineQualificationByOrderResponseBody {
	s.Data = v
	return s
}

func (s *GetHotlineQualificationByOrderResponseBody) SetMessage(v string) *GetHotlineQualificationByOrderResponseBody {
	s.Message = &v
	return s
}

func (s *GetHotlineQualificationByOrderResponseBody) SetRequestId(v string) *GetHotlineQualificationByOrderResponseBody {
	s.RequestId = &v
	return s
}

type GetHotlineQualificationByOrderResponseBodyData struct {
	OrderId         *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	QualificationId *string `json:"QualificationId,omitempty" xml:"QualificationId,omitempty"`
	Status          *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetHotlineQualificationByOrderResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetHotlineQualificationByOrderResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetHotlineQualificationByOrderResponseBodyData) SetOrderId(v string) *GetHotlineQualificationByOrderResponseBodyData {
	s.OrderId = &v
	return s
}

func (s *GetHotlineQualificationByOrderResponseBodyData) SetQualificationId(v string) *GetHotlineQualificationByOrderResponseBodyData {
	s.QualificationId = &v
	return s
}

func (s *GetHotlineQualificationByOrderResponseBodyData) SetStatus(v string) *GetHotlineQualificationByOrderResponseBodyData {
	s.Status = &v
	return s
}

type GetHotlineQualificationByOrderResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetHotlineQualificationByOrderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetHotlineQualificationByOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s GetHotlineQualificationByOrderResponse) GoString() string {
	return s.String()
}

func (s *GetHotlineQualificationByOrderResponse) SetHeaders(v map[string]*string) *GetHotlineQualificationByOrderResponse {
	s.Headers = v
	return s
}

func (s *GetHotlineQualificationByOrderResponse) SetBody(v *GetHotlineQualificationByOrderResponseBody) *GetHotlineQualificationByOrderResponse {
	s.Body = v
	return s
}

type GetMqttTokenRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s GetMqttTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMqttTokenRequest) GoString() string {
	return s.String()
}

func (s *GetMqttTokenRequest) SetOwnerId(v int64) *GetMqttTokenRequest {
	s.OwnerId = &v
	return s
}

func (s *GetMqttTokenRequest) SetResourceOwnerAccount(v string) *GetMqttTokenRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetMqttTokenRequest) SetResourceOwnerId(v int64) *GetMqttTokenRequest {
	s.ResourceOwnerId = &v
	return s
}

type GetMqttTokenResponseBody struct {
	Code      *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetMqttTokenResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetMqttTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMqttTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GetMqttTokenResponseBody) SetCode(v string) *GetMqttTokenResponseBody {
	s.Code = &v
	return s
}

func (s *GetMqttTokenResponseBody) SetData(v *GetMqttTokenResponseBodyData) *GetMqttTokenResponseBody {
	s.Data = v
	return s
}

func (s *GetMqttTokenResponseBody) SetMessage(v string) *GetMqttTokenResponseBody {
	s.Message = &v
	return s
}

func (s *GetMqttTokenResponseBody) SetRequestId(v string) *GetMqttTokenResponseBody {
	s.RequestId = &v
	return s
}

type GetMqttTokenResponseBodyData struct {
	ClientId   *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	Host       *string `json:"Host,omitempty" xml:"Host,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	P2pTopic   *string `json:"P2pTopic,omitempty" xml:"P2pTopic,omitempty"`
	ServerId   *string `json:"ServerId,omitempty" xml:"ServerId,omitempty"`
	Token      *string `json:"Token,omitempty" xml:"Token,omitempty"`
	Username   *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s GetMqttTokenResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetMqttTokenResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetMqttTokenResponseBodyData) SetClientId(v string) *GetMqttTokenResponseBodyData {
	s.ClientId = &v
	return s
}

func (s *GetMqttTokenResponseBodyData) SetExpireTime(v string) *GetMqttTokenResponseBodyData {
	s.ExpireTime = &v
	return s
}

func (s *GetMqttTokenResponseBodyData) SetHost(v string) *GetMqttTokenResponseBodyData {
	s.Host = &v
	return s
}

func (s *GetMqttTokenResponseBodyData) SetInstanceId(v string) *GetMqttTokenResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *GetMqttTokenResponseBodyData) SetP2pTopic(v string) *GetMqttTokenResponseBodyData {
	s.P2pTopic = &v
	return s
}

func (s *GetMqttTokenResponseBodyData) SetServerId(v string) *GetMqttTokenResponseBodyData {
	s.ServerId = &v
	return s
}

func (s *GetMqttTokenResponseBodyData) SetToken(v string) *GetMqttTokenResponseBodyData {
	s.Token = &v
	return s
}

func (s *GetMqttTokenResponseBodyData) SetUsername(v string) *GetMqttTokenResponseBodyData {
	s.Username = &v
	return s
}

type GetMqttTokenResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetMqttTokenResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetMqttTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMqttTokenResponse) GoString() string {
	return s.String()
}

func (s *GetMqttTokenResponse) SetHeaders(v map[string]*string) *GetMqttTokenResponse {
	s.Headers = v
	return s
}

func (s *GetMqttTokenResponse) SetBody(v *GetMqttTokenResponseBody) *GetMqttTokenResponse {
	s.Body = v
	return s
}

type GetRtcTokenRequest struct {
	DeviceId             *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	IsCustomAccount      *bool   `json:"IsCustomAccount,omitempty" xml:"IsCustomAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	UserId               *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetRtcTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRtcTokenRequest) GoString() string {
	return s.String()
}

func (s *GetRtcTokenRequest) SetDeviceId(v string) *GetRtcTokenRequest {
	s.DeviceId = &v
	return s
}

func (s *GetRtcTokenRequest) SetIsCustomAccount(v bool) *GetRtcTokenRequest {
	s.IsCustomAccount = &v
	return s
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

type GetRtcTokenResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Module    *string `json:"Module,omitempty" xml:"Module,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetRtcTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRtcTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GetRtcTokenResponseBody) SetCode(v string) *GetRtcTokenResponseBody {
	s.Code = &v
	return s
}

func (s *GetRtcTokenResponseBody) SetMessage(v string) *GetRtcTokenResponseBody {
	s.Message = &v
	return s
}

func (s *GetRtcTokenResponseBody) SetModule(v string) *GetRtcTokenResponseBody {
	s.Module = &v
	return s
}

func (s *GetRtcTokenResponseBody) SetRequestId(v string) *GetRtcTokenResponseBody {
	s.RequestId = &v
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

type GetTokenRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TokenType            *string `json:"TokenType,omitempty" xml:"TokenType,omitempty"`
}

func (s GetTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTokenRequest) GoString() string {
	return s.String()
}

func (s *GetTokenRequest) SetOwnerId(v int64) *GetTokenRequest {
	s.OwnerId = &v
	return s
}

func (s *GetTokenRequest) SetResourceOwnerAccount(v string) *GetTokenRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetTokenRequest) SetResourceOwnerId(v int64) *GetTokenRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetTokenRequest) SetTokenType(v string) *GetTokenRequest {
	s.TokenType = &v
	return s
}

type GetTokenResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	Token     *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s GetTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GetTokenResponseBody) SetCode(v string) *GetTokenResponseBody {
	s.Code = &v
	return s
}

func (s *GetTokenResponseBody) SetMessage(v string) *GetTokenResponseBody {
	s.Message = &v
	return s
}

func (s *GetTokenResponseBody) SetRequestId(v string) *GetTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTokenResponseBody) SetSuccess(v bool) *GetTokenResponseBody {
	s.Success = &v
	return s
}

func (s *GetTokenResponseBody) SetToken(v string) *GetTokenResponseBody {
	s.Token = &v
	return s
}

type GetTokenResponse struct {
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetTokenResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTokenResponse) GoString() string {
	return s.String()
}

func (s *GetTokenResponse) SetHeaders(v map[string]*string) *GetTokenResponse {
	s.Headers = v
	return s
}

func (s *GetTokenResponse) SetBody(v *GetTokenResponseBody) *GetTokenResponse {
	s.Body = v
	return s
}

type IvrCallRequest struct {
	ByeCode              *string                     `json:"ByeCode,omitempty" xml:"ByeCode,omitempty"`
	ByeTtsParams         *string                     `json:"ByeTtsParams,omitempty" xml:"ByeTtsParams,omitempty"`
	CalledNumber         *string                     `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	CalledShowNumber     *string                     `json:"CalledShowNumber,omitempty" xml:"CalledShowNumber,omitempty"`
	MenuKeyMap           []*IvrCallRequestMenuKeyMap `json:"MenuKeyMap,omitempty" xml:"MenuKeyMap,omitempty" type:"Repeated"`
	OutId                *string                     `json:"OutId,omitempty" xml:"OutId,omitempty"`
	OwnerId              *int64                      `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PlayTimes            *int64                      `json:"PlayTimes,omitempty" xml:"PlayTimes,omitempty"`
	ResourceOwnerAccount *string                     `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64                      `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	StartCode            *string                     `json:"StartCode,omitempty" xml:"StartCode,omitempty"`
	StartTtsParams       *string                     `json:"StartTtsParams,omitempty" xml:"StartTtsParams,omitempty"`
	Timeout              *int32                      `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s IvrCallRequest) String() string {
	return tea.Prettify(s)
}

func (s IvrCallRequest) GoString() string {
	return s.String()
}

func (s *IvrCallRequest) SetByeCode(v string) *IvrCallRequest {
	s.ByeCode = &v
	return s
}

func (s *IvrCallRequest) SetByeTtsParams(v string) *IvrCallRequest {
	s.ByeTtsParams = &v
	return s
}

func (s *IvrCallRequest) SetCalledNumber(v string) *IvrCallRequest {
	s.CalledNumber = &v
	return s
}

func (s *IvrCallRequest) SetCalledShowNumber(v string) *IvrCallRequest {
	s.CalledShowNumber = &v
	return s
}

func (s *IvrCallRequest) SetMenuKeyMap(v []*IvrCallRequestMenuKeyMap) *IvrCallRequest {
	s.MenuKeyMap = v
	return s
}

func (s *IvrCallRequest) SetOutId(v string) *IvrCallRequest {
	s.OutId = &v
	return s
}

func (s *IvrCallRequest) SetOwnerId(v int64) *IvrCallRequest {
	s.OwnerId = &v
	return s
}

func (s *IvrCallRequest) SetPlayTimes(v int64) *IvrCallRequest {
	s.PlayTimes = &v
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

func (s *IvrCallRequest) SetStartCode(v string) *IvrCallRequest {
	s.StartCode = &v
	return s
}

func (s *IvrCallRequest) SetStartTtsParams(v string) *IvrCallRequest {
	s.StartTtsParams = &v
	return s
}

func (s *IvrCallRequest) SetTimeout(v int32) *IvrCallRequest {
	s.Timeout = &v
	return s
}

type IvrCallRequestMenuKeyMap struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Key       *string `json:"Key,omitempty" xml:"Key,omitempty"`
	TtsParams *string `json:"TtsParams,omitempty" xml:"TtsParams,omitempty"`
}

func (s IvrCallRequestMenuKeyMap) String() string {
	return tea.Prettify(s)
}

func (s IvrCallRequestMenuKeyMap) GoString() string {
	return s.String()
}

func (s *IvrCallRequestMenuKeyMap) SetCode(v string) *IvrCallRequestMenuKeyMap {
	s.Code = &v
	return s
}

func (s *IvrCallRequestMenuKeyMap) SetKey(v string) *IvrCallRequestMenuKeyMap {
	s.Key = &v
	return s
}

func (s *IvrCallRequestMenuKeyMap) SetTtsParams(v string) *IvrCallRequestMenuKeyMap {
	s.TtsParams = &v
	return s
}

type IvrCallResponseBody struct {
	CallId    *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s IvrCallResponseBody) String() string {
	return tea.Prettify(s)
}

func (s IvrCallResponseBody) GoString() string {
	return s.String()
}

func (s *IvrCallResponseBody) SetCallId(v string) *IvrCallResponseBody {
	s.CallId = &v
	return s
}

func (s *IvrCallResponseBody) SetCode(v string) *IvrCallResponseBody {
	s.Code = &v
	return s
}

func (s *IvrCallResponseBody) SetMessage(v string) *IvrCallResponseBody {
	s.Message = &v
	return s
}

func (s *IvrCallResponseBody) SetRequestId(v string) *IvrCallResponseBody {
	s.RequestId = &v
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

type ListCallTaskRequest struct {
	BizType              *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Status               *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TaskId               *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskName             *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	TemplateName         *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s ListCallTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCallTaskRequest) GoString() string {
	return s.String()
}

func (s *ListCallTaskRequest) SetBizType(v string) *ListCallTaskRequest {
	s.BizType = &v
	return s
}

func (s *ListCallTaskRequest) SetOwnerId(v int64) *ListCallTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *ListCallTaskRequest) SetPageNumber(v int32) *ListCallTaskRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCallTaskRequest) SetPageSize(v int32) *ListCallTaskRequest {
	s.PageSize = &v
	return s
}

func (s *ListCallTaskRequest) SetResourceOwnerAccount(v string) *ListCallTaskRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListCallTaskRequest) SetResourceOwnerId(v int64) *ListCallTaskRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListCallTaskRequest) SetStatus(v string) *ListCallTaskRequest {
	s.Status = &v
	return s
}

func (s *ListCallTaskRequest) SetTaskId(v string) *ListCallTaskRequest {
	s.TaskId = &v
	return s
}

func (s *ListCallTaskRequest) SetTaskName(v string) *ListCallTaskRequest {
	s.TaskName = &v
	return s
}

func (s *ListCallTaskRequest) SetTemplateName(v string) *ListCallTaskRequest {
	s.TemplateName = &v
	return s
}

type ListCallTaskResponseBody struct {
	Code       *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data       []*ListCallTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	PageNumber *int64                          `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int64                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total      *int64                          `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListCallTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCallTaskResponseBody) GoString() string {
	return s.String()
}

func (s *ListCallTaskResponseBody) SetCode(v string) *ListCallTaskResponseBody {
	s.Code = &v
	return s
}

func (s *ListCallTaskResponseBody) SetData(v []*ListCallTaskResponseBodyData) *ListCallTaskResponseBody {
	s.Data = v
	return s
}

func (s *ListCallTaskResponseBody) SetPageNumber(v int64) *ListCallTaskResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListCallTaskResponseBody) SetPageSize(v int64) *ListCallTaskResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListCallTaskResponseBody) SetRequestId(v string) *ListCallTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCallTaskResponseBody) SetTotal(v int64) *ListCallTaskResponseBody {
	s.Total = &v
	return s
}

type ListCallTaskResponseBodyData struct {
	BizType        *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	CompleteTime   *string `json:"CompleteTime,omitempty" xml:"CompleteTime,omitempty"`
	CompletedCount *int64  `json:"CompletedCount,omitempty" xml:"CompletedCount,omitempty"`
	CompletedRate  *int32  `json:"CompletedRate,omitempty" xml:"CompletedRate,omitempty"`
	Data           *string `json:"Data,omitempty" xml:"Data,omitempty"`
	DataType       *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	FireTime       *string `json:"FireTime,omitempty" xml:"FireTime,omitempty"`
	Id             *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Resource       *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	Status         *string `json:"Status,omitempty" xml:"Status,omitempty"`
	StopTime       *string `json:"StopTime,omitempty" xml:"StopTime,omitempty"`
	TaskName       *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	TemplateCode   *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	TemplateName   *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	TotalCount     *int64  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCallTaskResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListCallTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListCallTaskResponseBodyData) SetBizType(v string) *ListCallTaskResponseBodyData {
	s.BizType = &v
	return s
}

func (s *ListCallTaskResponseBodyData) SetCompleteTime(v string) *ListCallTaskResponseBodyData {
	s.CompleteTime = &v
	return s
}

func (s *ListCallTaskResponseBodyData) SetCompletedCount(v int64) *ListCallTaskResponseBodyData {
	s.CompletedCount = &v
	return s
}

func (s *ListCallTaskResponseBodyData) SetCompletedRate(v int32) *ListCallTaskResponseBodyData {
	s.CompletedRate = &v
	return s
}

func (s *ListCallTaskResponseBodyData) SetData(v string) *ListCallTaskResponseBodyData {
	s.Data = &v
	return s
}

func (s *ListCallTaskResponseBodyData) SetDataType(v string) *ListCallTaskResponseBodyData {
	s.DataType = &v
	return s
}

func (s *ListCallTaskResponseBodyData) SetFireTime(v string) *ListCallTaskResponseBodyData {
	s.FireTime = &v
	return s
}

func (s *ListCallTaskResponseBodyData) SetId(v int64) *ListCallTaskResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListCallTaskResponseBodyData) SetResource(v string) *ListCallTaskResponseBodyData {
	s.Resource = &v
	return s
}

func (s *ListCallTaskResponseBodyData) SetStatus(v string) *ListCallTaskResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListCallTaskResponseBodyData) SetStopTime(v string) *ListCallTaskResponseBodyData {
	s.StopTime = &v
	return s
}

func (s *ListCallTaskResponseBodyData) SetTaskName(v string) *ListCallTaskResponseBodyData {
	s.TaskName = &v
	return s
}

func (s *ListCallTaskResponseBodyData) SetTemplateCode(v string) *ListCallTaskResponseBodyData {
	s.TemplateCode = &v
	return s
}

func (s *ListCallTaskResponseBodyData) SetTemplateName(v string) *ListCallTaskResponseBodyData {
	s.TemplateName = &v
	return s
}

func (s *ListCallTaskResponseBodyData) SetTotalCount(v int64) *ListCallTaskResponseBodyData {
	s.TotalCount = &v
	return s
}

type ListCallTaskResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListCallTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListCallTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCallTaskResponse) GoString() string {
	return s.String()
}

func (s *ListCallTaskResponse) SetHeaders(v map[string]*string) *ListCallTaskResponse {
	s.Headers = v
	return s
}

func (s *ListCallTaskResponse) SetBody(v *ListCallTaskResponseBody) *ListCallTaskResponse {
	s.Body = v
	return s
}

type ListCallTaskDetailRequest struct {
	CalledNum            *string `json:"CalledNum,omitempty" xml:"CalledNum,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Status               *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TaskId               *int64  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ListCallTaskDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCallTaskDetailRequest) GoString() string {
	return s.String()
}

func (s *ListCallTaskDetailRequest) SetCalledNum(v string) *ListCallTaskDetailRequest {
	s.CalledNum = &v
	return s
}

func (s *ListCallTaskDetailRequest) SetOwnerId(v int64) *ListCallTaskDetailRequest {
	s.OwnerId = &v
	return s
}

func (s *ListCallTaskDetailRequest) SetPageNumber(v int32) *ListCallTaskDetailRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCallTaskDetailRequest) SetPageSize(v int32) *ListCallTaskDetailRequest {
	s.PageSize = &v
	return s
}

func (s *ListCallTaskDetailRequest) SetResourceOwnerAccount(v string) *ListCallTaskDetailRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListCallTaskDetailRequest) SetResourceOwnerId(v int64) *ListCallTaskDetailRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListCallTaskDetailRequest) SetStatus(v string) *ListCallTaskDetailRequest {
	s.Status = &v
	return s
}

func (s *ListCallTaskDetailRequest) SetTaskId(v int64) *ListCallTaskDetailRequest {
	s.TaskId = &v
	return s
}

type ListCallTaskDetailResponseBody struct {
	Code       *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data       []*ListCallTaskDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	PageNumber *int64                                `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int64                                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Total      *int64                                `json:"Total,omitempty" xml:"Total,omitempty"`
	TotalPage  *int64                                `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s ListCallTaskDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCallTaskDetailResponseBody) GoString() string {
	return s.String()
}

func (s *ListCallTaskDetailResponseBody) SetCode(v string) *ListCallTaskDetailResponseBody {
	s.Code = &v
	return s
}

func (s *ListCallTaskDetailResponseBody) SetData(v []*ListCallTaskDetailResponseBodyData) *ListCallTaskDetailResponseBody {
	s.Data = v
	return s
}

func (s *ListCallTaskDetailResponseBody) SetPageNumber(v int64) *ListCallTaskDetailResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListCallTaskDetailResponseBody) SetPageSize(v int64) *ListCallTaskDetailResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListCallTaskDetailResponseBody) SetRequestId(v string) *ListCallTaskDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCallTaskDetailResponseBody) SetTotal(v int64) *ListCallTaskDetailResponseBody {
	s.Total = &v
	return s
}

func (s *ListCallTaskDetailResponseBody) SetTotalPage(v int64) *ListCallTaskDetailResponseBody {
	s.TotalPage = &v
	return s
}

type ListCallTaskDetailResponseBodyData struct {
	CalledNum *string `json:"CalledNum,omitempty" xml:"CalledNum,omitempty"`
	Caller    *string `json:"Caller,omitempty" xml:"Caller,omitempty"`
	Duration  *int64  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	Id        *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListCallTaskDetailResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListCallTaskDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListCallTaskDetailResponseBodyData) SetCalledNum(v string) *ListCallTaskDetailResponseBodyData {
	s.CalledNum = &v
	return s
}

func (s *ListCallTaskDetailResponseBodyData) SetCaller(v string) *ListCallTaskDetailResponseBodyData {
	s.Caller = &v
	return s
}

func (s *ListCallTaskDetailResponseBodyData) SetDuration(v int64) *ListCallTaskDetailResponseBodyData {
	s.Duration = &v
	return s
}

func (s *ListCallTaskDetailResponseBodyData) SetId(v int64) *ListCallTaskDetailResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListCallTaskDetailResponseBodyData) SetStatus(v string) *ListCallTaskDetailResponseBodyData {
	s.Status = &v
	return s
}

type ListCallTaskDetailResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListCallTaskDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListCallTaskDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCallTaskDetailResponse) GoString() string {
	return s.String()
}

func (s *ListCallTaskDetailResponse) SetHeaders(v map[string]*string) *ListCallTaskDetailResponse {
	s.Headers = v
	return s
}

func (s *ListCallTaskDetailResponse) SetBody(v *ListCallTaskDetailResponseBody) *ListCallTaskDetailResponse {
	s.Body = v
	return s
}

type ListHotlineTransferNumberRequest struct {
	HotlineNumber        *string `json:"HotlineNumber,omitempty" xml:"HotlineNumber,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNo               *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	QualificationId      *string `json:"QualificationId,omitempty" xml:"QualificationId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ListHotlineTransferNumberRequest) String() string {
	return tea.Prettify(s)
}

func (s ListHotlineTransferNumberRequest) GoString() string {
	return s.String()
}

func (s *ListHotlineTransferNumberRequest) SetHotlineNumber(v string) *ListHotlineTransferNumberRequest {
	s.HotlineNumber = &v
	return s
}

func (s *ListHotlineTransferNumberRequest) SetOwnerId(v int64) *ListHotlineTransferNumberRequest {
	s.OwnerId = &v
	return s
}

func (s *ListHotlineTransferNumberRequest) SetPageNo(v int32) *ListHotlineTransferNumberRequest {
	s.PageNo = &v
	return s
}

func (s *ListHotlineTransferNumberRequest) SetPageSize(v int32) *ListHotlineTransferNumberRequest {
	s.PageSize = &v
	return s
}

func (s *ListHotlineTransferNumberRequest) SetQualificationId(v string) *ListHotlineTransferNumberRequest {
	s.QualificationId = &v
	return s
}

func (s *ListHotlineTransferNumberRequest) SetResourceOwnerAccount(v string) *ListHotlineTransferNumberRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListHotlineTransferNumberRequest) SetResourceOwnerId(v int64) *ListHotlineTransferNumberRequest {
	s.ResourceOwnerId = &v
	return s
}

type ListHotlineTransferNumberResponseBody struct {
	Code      *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *ListHotlineTransferNumberResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListHotlineTransferNumberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListHotlineTransferNumberResponseBody) GoString() string {
	return s.String()
}

func (s *ListHotlineTransferNumberResponseBody) SetCode(v string) *ListHotlineTransferNumberResponseBody {
	s.Code = &v
	return s
}

func (s *ListHotlineTransferNumberResponseBody) SetData(v *ListHotlineTransferNumberResponseBodyData) *ListHotlineTransferNumberResponseBody {
	s.Data = v
	return s
}

func (s *ListHotlineTransferNumberResponseBody) SetMessage(v string) *ListHotlineTransferNumberResponseBody {
	s.Message = &v
	return s
}

func (s *ListHotlineTransferNumberResponseBody) SetRequestId(v string) *ListHotlineTransferNumberResponseBody {
	s.RequestId = &v
	return s
}

type ListHotlineTransferNumberResponseBodyData struct {
	PageNo   *int32                                             `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize *int32                                             `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total    *int64                                             `json:"Total,omitempty" xml:"Total,omitempty"`
	Values   []*ListHotlineTransferNumberResponseBodyDataValues `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s ListHotlineTransferNumberResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListHotlineTransferNumberResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListHotlineTransferNumberResponseBodyData) SetPageNo(v int32) *ListHotlineTransferNumberResponseBodyData {
	s.PageNo = &v
	return s
}

func (s *ListHotlineTransferNumberResponseBodyData) SetPageSize(v int32) *ListHotlineTransferNumberResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListHotlineTransferNumberResponseBodyData) SetTotal(v int64) *ListHotlineTransferNumberResponseBodyData {
	s.Total = &v
	return s
}

func (s *ListHotlineTransferNumberResponseBodyData) SetValues(v []*ListHotlineTransferNumberResponseBodyDataValues) *ListHotlineTransferNumberResponseBodyData {
	s.Values = v
	return s
}

type ListHotlineTransferNumberResponseBodyDataValues struct {
	HotlineNumber   *string `json:"HotlineNumber,omitempty" xml:"HotlineNumber,omitempty"`
	IdentityCard    *string `json:"IdentityCard,omitempty" xml:"IdentityCard,omitempty"`
	NumberOwnerName *string `json:"NumberOwnerName,omitempty" xml:"NumberOwnerName,omitempty"`
	PhoneNumber     *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	QualificationId *string `json:"QualificationId,omitempty" xml:"QualificationId,omitempty"`
	ResUniqueCode   *string `json:"ResUniqueCode,omitempty" xml:"ResUniqueCode,omitempty"`
}

func (s ListHotlineTransferNumberResponseBodyDataValues) String() string {
	return tea.Prettify(s)
}

func (s ListHotlineTransferNumberResponseBodyDataValues) GoString() string {
	return s.String()
}

func (s *ListHotlineTransferNumberResponseBodyDataValues) SetHotlineNumber(v string) *ListHotlineTransferNumberResponseBodyDataValues {
	s.HotlineNumber = &v
	return s
}

func (s *ListHotlineTransferNumberResponseBodyDataValues) SetIdentityCard(v string) *ListHotlineTransferNumberResponseBodyDataValues {
	s.IdentityCard = &v
	return s
}

func (s *ListHotlineTransferNumberResponseBodyDataValues) SetNumberOwnerName(v string) *ListHotlineTransferNumberResponseBodyDataValues {
	s.NumberOwnerName = &v
	return s
}

func (s *ListHotlineTransferNumberResponseBodyDataValues) SetPhoneNumber(v string) *ListHotlineTransferNumberResponseBodyDataValues {
	s.PhoneNumber = &v
	return s
}

func (s *ListHotlineTransferNumberResponseBodyDataValues) SetQualificationId(v string) *ListHotlineTransferNumberResponseBodyDataValues {
	s.QualificationId = &v
	return s
}

func (s *ListHotlineTransferNumberResponseBodyDataValues) SetResUniqueCode(v string) *ListHotlineTransferNumberResponseBodyDataValues {
	s.ResUniqueCode = &v
	return s
}

type ListHotlineTransferNumberResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListHotlineTransferNumberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListHotlineTransferNumberResponse) String() string {
	return tea.Prettify(s)
}

func (s ListHotlineTransferNumberResponse) GoString() string {
	return s.String()
}

func (s *ListHotlineTransferNumberResponse) SetHeaders(v map[string]*string) *ListHotlineTransferNumberResponse {
	s.Headers = v
	return s
}

func (s *ListHotlineTransferNumberResponse) SetBody(v *ListHotlineTransferNumberResponseBody) *ListHotlineTransferNumberResponse {
	s.Body = v
	return s
}

type ListHotlineTransferRegisterFileRequest struct {
	HotlineNumber        *string `json:"HotlineNumber,omitempty" xml:"HotlineNumber,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNo               *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	QualificationId      *string `json:"QualificationId,omitempty" xml:"QualificationId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ListHotlineTransferRegisterFileRequest) String() string {
	return tea.Prettify(s)
}

func (s ListHotlineTransferRegisterFileRequest) GoString() string {
	return s.String()
}

func (s *ListHotlineTransferRegisterFileRequest) SetHotlineNumber(v string) *ListHotlineTransferRegisterFileRequest {
	s.HotlineNumber = &v
	return s
}

func (s *ListHotlineTransferRegisterFileRequest) SetOwnerId(v int64) *ListHotlineTransferRegisterFileRequest {
	s.OwnerId = &v
	return s
}

func (s *ListHotlineTransferRegisterFileRequest) SetPageNo(v int32) *ListHotlineTransferRegisterFileRequest {
	s.PageNo = &v
	return s
}

func (s *ListHotlineTransferRegisterFileRequest) SetPageSize(v int32) *ListHotlineTransferRegisterFileRequest {
	s.PageSize = &v
	return s
}

func (s *ListHotlineTransferRegisterFileRequest) SetQualificationId(v string) *ListHotlineTransferRegisterFileRequest {
	s.QualificationId = &v
	return s
}

func (s *ListHotlineTransferRegisterFileRequest) SetResourceOwnerAccount(v string) *ListHotlineTransferRegisterFileRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListHotlineTransferRegisterFileRequest) SetResourceOwnerId(v int64) *ListHotlineTransferRegisterFileRequest {
	s.ResourceOwnerId = &v
	return s
}

type ListHotlineTransferRegisterFileResponseBody struct {
	Code      *string                                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *ListHotlineTransferRegisterFileResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListHotlineTransferRegisterFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListHotlineTransferRegisterFileResponseBody) GoString() string {
	return s.String()
}

func (s *ListHotlineTransferRegisterFileResponseBody) SetCode(v string) *ListHotlineTransferRegisterFileResponseBody {
	s.Code = &v
	return s
}

func (s *ListHotlineTransferRegisterFileResponseBody) SetData(v *ListHotlineTransferRegisterFileResponseBodyData) *ListHotlineTransferRegisterFileResponseBody {
	s.Data = v
	return s
}

func (s *ListHotlineTransferRegisterFileResponseBody) SetMessage(v string) *ListHotlineTransferRegisterFileResponseBody {
	s.Message = &v
	return s
}

func (s *ListHotlineTransferRegisterFileResponseBody) SetRequestId(v string) *ListHotlineTransferRegisterFileResponseBody {
	s.RequestId = &v
	return s
}

type ListHotlineTransferRegisterFileResponseBodyData struct {
	PageNo   *int32                                                   `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize *int32                                                   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total    *int64                                                   `json:"Total,omitempty" xml:"Total,omitempty"`
	Values   []*ListHotlineTransferRegisterFileResponseBodyDataValues `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s ListHotlineTransferRegisterFileResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListHotlineTransferRegisterFileResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListHotlineTransferRegisterFileResponseBodyData) SetPageNo(v int32) *ListHotlineTransferRegisterFileResponseBodyData {
	s.PageNo = &v
	return s
}

func (s *ListHotlineTransferRegisterFileResponseBodyData) SetPageSize(v int32) *ListHotlineTransferRegisterFileResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListHotlineTransferRegisterFileResponseBodyData) SetTotal(v int64) *ListHotlineTransferRegisterFileResponseBodyData {
	s.Total = &v
	return s
}

func (s *ListHotlineTransferRegisterFileResponseBodyData) SetValues(v []*ListHotlineTransferRegisterFileResponseBodyDataValues) *ListHotlineTransferRegisterFileResponseBodyData {
	s.Values = v
	return s
}

type ListHotlineTransferRegisterFileResponseBodyDataValues struct {
	Agree             *string `json:"Agree,omitempty" xml:"Agree,omitempty"`
	CorpName          *string `json:"CorpName,omitempty" xml:"CorpName,omitempty"`
	HotlineNumber     *string `json:"HotlineNumber,omitempty" xml:"HotlineNumber,omitempty"`
	MngOpIdentityCard *string `json:"MngOpIdentityCard,omitempty" xml:"MngOpIdentityCard,omitempty"`
	MngOpMail         *string `json:"MngOpMail,omitempty" xml:"MngOpMail,omitempty"`
	MngOpMobile       *string `json:"MngOpMobile,omitempty" xml:"MngOpMobile,omitempty"`
	MngOpName         *string `json:"MngOpName,omitempty" xml:"MngOpName,omitempty"`
	QualificationId   *string `json:"QualificationId,omitempty" xml:"QualificationId,omitempty"`
	ResUniqueCode     *int64  `json:"ResUniqueCode,omitempty" xml:"ResUniqueCode,omitempty"`
}

func (s ListHotlineTransferRegisterFileResponseBodyDataValues) String() string {
	return tea.Prettify(s)
}

func (s ListHotlineTransferRegisterFileResponseBodyDataValues) GoString() string {
	return s.String()
}

func (s *ListHotlineTransferRegisterFileResponseBodyDataValues) SetAgree(v string) *ListHotlineTransferRegisterFileResponseBodyDataValues {
	s.Agree = &v
	return s
}

func (s *ListHotlineTransferRegisterFileResponseBodyDataValues) SetCorpName(v string) *ListHotlineTransferRegisterFileResponseBodyDataValues {
	s.CorpName = &v
	return s
}

func (s *ListHotlineTransferRegisterFileResponseBodyDataValues) SetHotlineNumber(v string) *ListHotlineTransferRegisterFileResponseBodyDataValues {
	s.HotlineNumber = &v
	return s
}

func (s *ListHotlineTransferRegisterFileResponseBodyDataValues) SetMngOpIdentityCard(v string) *ListHotlineTransferRegisterFileResponseBodyDataValues {
	s.MngOpIdentityCard = &v
	return s
}

func (s *ListHotlineTransferRegisterFileResponseBodyDataValues) SetMngOpMail(v string) *ListHotlineTransferRegisterFileResponseBodyDataValues {
	s.MngOpMail = &v
	return s
}

func (s *ListHotlineTransferRegisterFileResponseBodyDataValues) SetMngOpMobile(v string) *ListHotlineTransferRegisterFileResponseBodyDataValues {
	s.MngOpMobile = &v
	return s
}

func (s *ListHotlineTransferRegisterFileResponseBodyDataValues) SetMngOpName(v string) *ListHotlineTransferRegisterFileResponseBodyDataValues {
	s.MngOpName = &v
	return s
}

func (s *ListHotlineTransferRegisterFileResponseBodyDataValues) SetQualificationId(v string) *ListHotlineTransferRegisterFileResponseBodyDataValues {
	s.QualificationId = &v
	return s
}

func (s *ListHotlineTransferRegisterFileResponseBodyDataValues) SetResUniqueCode(v int64) *ListHotlineTransferRegisterFileResponseBodyDataValues {
	s.ResUniqueCode = &v
	return s
}

type ListHotlineTransferRegisterFileResponse struct {
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListHotlineTransferRegisterFileResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListHotlineTransferRegisterFileResponse) String() string {
	return tea.Prettify(s)
}

func (s ListHotlineTransferRegisterFileResponse) GoString() string {
	return s.String()
}

func (s *ListHotlineTransferRegisterFileResponse) SetHeaders(v map[string]*string) *ListHotlineTransferRegisterFileResponse {
	s.Headers = v
	return s
}

func (s *ListHotlineTransferRegisterFileResponse) SetBody(v *ListHotlineTransferRegisterFileResponseBody) *ListHotlineTransferRegisterFileResponse {
	s.Body = v
	return s
}

type QueryCallDetailByCallIdRequest struct {
	CallId               *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProdId               *int64  `json:"ProdId,omitempty" xml:"ProdId,omitempty"`
	QueryDate            *int64  `json:"QueryDate,omitempty" xml:"QueryDate,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s QueryCallDetailByCallIdRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryCallDetailByCallIdRequest) GoString() string {
	return s.String()
}

func (s *QueryCallDetailByCallIdRequest) SetCallId(v string) *QueryCallDetailByCallIdRequest {
	s.CallId = &v
	return s
}

func (s *QueryCallDetailByCallIdRequest) SetOwnerId(v int64) *QueryCallDetailByCallIdRequest {
	s.OwnerId = &v
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

func (s *QueryCallDetailByCallIdRequest) SetResourceOwnerAccount(v string) *QueryCallDetailByCallIdRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryCallDetailByCallIdRequest) SetResourceOwnerId(v int64) *QueryCallDetailByCallIdRequest {
	s.ResourceOwnerId = &v
	return s
}

type QueryCallDetailByCallIdResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryCallDetailByCallIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryCallDetailByCallIdResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCallDetailByCallIdResponseBody) SetCode(v string) *QueryCallDetailByCallIdResponseBody {
	s.Code = &v
	return s
}

func (s *QueryCallDetailByCallIdResponseBody) SetData(v string) *QueryCallDetailByCallIdResponseBody {
	s.Data = &v
	return s
}

func (s *QueryCallDetailByCallIdResponseBody) SetMessage(v string) *QueryCallDetailByCallIdResponseBody {
	s.Message = &v
	return s
}

func (s *QueryCallDetailByCallIdResponseBody) SetRequestId(v string) *QueryCallDetailByCallIdResponseBody {
	s.RequestId = &v
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
	Callee               *string `json:"Callee,omitempty" xml:"Callee,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	QueryDate            *int64  `json:"QueryDate,omitempty" xml:"QueryDate,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TaskId               *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s QueryCallDetailByTaskIdRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryCallDetailByTaskIdRequest) GoString() string {
	return s.String()
}

func (s *QueryCallDetailByTaskIdRequest) SetCallee(v string) *QueryCallDetailByTaskIdRequest {
	s.Callee = &v
	return s
}

func (s *QueryCallDetailByTaskIdRequest) SetOwnerId(v int64) *QueryCallDetailByTaskIdRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryCallDetailByTaskIdRequest) SetQueryDate(v int64) *QueryCallDetailByTaskIdRequest {
	s.QueryDate = &v
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

type QueryCallDetailByTaskIdResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryCallDetailByTaskIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryCallDetailByTaskIdResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCallDetailByTaskIdResponseBody) SetCode(v string) *QueryCallDetailByTaskIdResponseBody {
	s.Code = &v
	return s
}

func (s *QueryCallDetailByTaskIdResponseBody) SetData(v string) *QueryCallDetailByTaskIdResponseBody {
	s.Data = &v
	return s
}

func (s *QueryCallDetailByTaskIdResponseBody) SetMessage(v string) *QueryCallDetailByTaskIdResponseBody {
	s.Message = &v
	return s
}

func (s *QueryCallDetailByTaskIdResponseBody) SetRequestId(v string) *QueryCallDetailByTaskIdResponseBody {
	s.RequestId = &v
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

type QueryCallInPoolTransferConfigRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PhoneNumber          *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s QueryCallInPoolTransferConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryCallInPoolTransferConfigRequest) GoString() string {
	return s.String()
}

func (s *QueryCallInPoolTransferConfigRequest) SetOwnerId(v int64) *QueryCallInPoolTransferConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryCallInPoolTransferConfigRequest) SetPhoneNumber(v string) *QueryCallInPoolTransferConfigRequest {
	s.PhoneNumber = &v
	return s
}

func (s *QueryCallInPoolTransferConfigRequest) SetResourceOwnerAccount(v string) *QueryCallInPoolTransferConfigRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryCallInPoolTransferConfigRequest) SetResourceOwnerId(v int64) *QueryCallInPoolTransferConfigRequest {
	s.ResourceOwnerId = &v
	return s
}

type QueryCallInPoolTransferConfigResponseBody struct {
	Code      *string                                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *QueryCallInPoolTransferConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                        `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryCallInPoolTransferConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryCallInPoolTransferConfigResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCallInPoolTransferConfigResponseBody) SetCode(v string) *QueryCallInPoolTransferConfigResponseBody {
	s.Code = &v
	return s
}

func (s *QueryCallInPoolTransferConfigResponseBody) SetData(v *QueryCallInPoolTransferConfigResponseBodyData) *QueryCallInPoolTransferConfigResponseBody {
	s.Data = v
	return s
}

func (s *QueryCallInPoolTransferConfigResponseBody) SetMessage(v string) *QueryCallInPoolTransferConfigResponseBody {
	s.Message = &v
	return s
}

func (s *QueryCallInPoolTransferConfigResponseBody) SetRequestId(v string) *QueryCallInPoolTransferConfigResponseBody {
	s.RequestId = &v
	return s
}

type QueryCallInPoolTransferConfigResponseBodyData struct {
	CalledRouteMode *string                                                 `json:"CalledRouteMode,omitempty" xml:"CalledRouteMode,omitempty"`
	Details         []*QueryCallInPoolTransferConfigResponseBodyDataDetails `json:"Details,omitempty" xml:"Details,omitempty" type:"Repeated"`
	GmtCreate       *int64                                                  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	TransferTimeout *string                                                 `json:"TransferTimeout,omitempty" xml:"TransferTimeout,omitempty"`
}

func (s QueryCallInPoolTransferConfigResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryCallInPoolTransferConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryCallInPoolTransferConfigResponseBodyData) SetCalledRouteMode(v string) *QueryCallInPoolTransferConfigResponseBodyData {
	s.CalledRouteMode = &v
	return s
}

func (s *QueryCallInPoolTransferConfigResponseBodyData) SetDetails(v []*QueryCallInPoolTransferConfigResponseBodyDataDetails) *QueryCallInPoolTransferConfigResponseBodyData {
	s.Details = v
	return s
}

func (s *QueryCallInPoolTransferConfigResponseBodyData) SetGmtCreate(v int64) *QueryCallInPoolTransferConfigResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *QueryCallInPoolTransferConfigResponseBodyData) SetTransferTimeout(v string) *QueryCallInPoolTransferConfigResponseBodyData {
	s.TransferTimeout = &v
	return s
}

type QueryCallInPoolTransferConfigResponseBodyDataDetails struct {
	Called *string `json:"Called,omitempty" xml:"Called,omitempty"`
}

func (s QueryCallInPoolTransferConfigResponseBodyDataDetails) String() string {
	return tea.Prettify(s)
}

func (s QueryCallInPoolTransferConfigResponseBodyDataDetails) GoString() string {
	return s.String()
}

func (s *QueryCallInPoolTransferConfigResponseBodyDataDetails) SetCalled(v string) *QueryCallInPoolTransferConfigResponseBodyDataDetails {
	s.Called = &v
	return s
}

type QueryCallInPoolTransferConfigResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryCallInPoolTransferConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryCallInPoolTransferConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryCallInPoolTransferConfigResponse) GoString() string {
	return s.String()
}

func (s *QueryCallInPoolTransferConfigResponse) SetHeaders(v map[string]*string) *QueryCallInPoolTransferConfigResponse {
	s.Headers = v
	return s
}

func (s *QueryCallInPoolTransferConfigResponse) SetBody(v *QueryCallInPoolTransferConfigResponseBody) *QueryCallInPoolTransferConfigResponse {
	s.Body = v
	return s
}

type QueryCallInTransferRecordRequest struct {
	CallInCaller         *string `json:"CallInCaller,omitempty" xml:"CallInCaller,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNo               *int64  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize             *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PhoneNumber          *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	QueryDate            *string `json:"QueryDate,omitempty" xml:"QueryDate,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s QueryCallInTransferRecordRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryCallInTransferRecordRequest) GoString() string {
	return s.String()
}

func (s *QueryCallInTransferRecordRequest) SetCallInCaller(v string) *QueryCallInTransferRecordRequest {
	s.CallInCaller = &v
	return s
}

func (s *QueryCallInTransferRecordRequest) SetOwnerId(v int64) *QueryCallInTransferRecordRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryCallInTransferRecordRequest) SetPageNo(v int64) *QueryCallInTransferRecordRequest {
	s.PageNo = &v
	return s
}

func (s *QueryCallInTransferRecordRequest) SetPageSize(v int64) *QueryCallInTransferRecordRequest {
	s.PageSize = &v
	return s
}

func (s *QueryCallInTransferRecordRequest) SetPhoneNumber(v string) *QueryCallInTransferRecordRequest {
	s.PhoneNumber = &v
	return s
}

func (s *QueryCallInTransferRecordRequest) SetQueryDate(v string) *QueryCallInTransferRecordRequest {
	s.QueryDate = &v
	return s
}

func (s *QueryCallInTransferRecordRequest) SetResourceOwnerAccount(v string) *QueryCallInTransferRecordRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryCallInTransferRecordRequest) SetResourceOwnerId(v int64) *QueryCallInTransferRecordRequest {
	s.ResourceOwnerId = &v
	return s
}

type QueryCallInTransferRecordResponseBody struct {
	Code      *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *QueryCallInTransferRecordResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryCallInTransferRecordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryCallInTransferRecordResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCallInTransferRecordResponseBody) SetCode(v string) *QueryCallInTransferRecordResponseBody {
	s.Code = &v
	return s
}

func (s *QueryCallInTransferRecordResponseBody) SetData(v *QueryCallInTransferRecordResponseBodyData) *QueryCallInTransferRecordResponseBody {
	s.Data = v
	return s
}

func (s *QueryCallInTransferRecordResponseBody) SetMessage(v string) *QueryCallInTransferRecordResponseBody {
	s.Message = &v
	return s
}

func (s *QueryCallInTransferRecordResponseBody) SetRequestId(v string) *QueryCallInTransferRecordResponseBody {
	s.RequestId = &v
	return s
}

type QueryCallInTransferRecordResponseBodyData struct {
	PageNo   *int64                                             `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize *int64                                             `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total    *int64                                             `json:"Total,omitempty" xml:"Total,omitempty"`
	Values   []*QueryCallInTransferRecordResponseBodyDataValues `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s QueryCallInTransferRecordResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryCallInTransferRecordResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryCallInTransferRecordResponseBodyData) SetPageNo(v int64) *QueryCallInTransferRecordResponseBodyData {
	s.PageNo = &v
	return s
}

func (s *QueryCallInTransferRecordResponseBodyData) SetPageSize(v int64) *QueryCallInTransferRecordResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryCallInTransferRecordResponseBodyData) SetTotal(v int64) *QueryCallInTransferRecordResponseBodyData {
	s.Total = &v
	return s
}

func (s *QueryCallInTransferRecordResponseBodyData) SetValues(v []*QueryCallInTransferRecordResponseBodyDataValues) *QueryCallInTransferRecordResponseBodyData {
	s.Values = v
	return s
}

type QueryCallInTransferRecordResponseBodyDataValues struct {
	CallInCalled   *string `json:"CallInCalled,omitempty" xml:"CallInCalled,omitempty"`
	CallInCaller   *string `json:"CallInCaller,omitempty" xml:"CallInCaller,omitempty"`
	GmtCreate      *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	RecordUrl      *string `json:"RecordUrl,omitempty" xml:"RecordUrl,omitempty"`
	TransferCalled *string `json:"TransferCalled,omitempty" xml:"TransferCalled,omitempty"`
	TransferCaller *string `json:"TransferCaller,omitempty" xml:"TransferCaller,omitempty"`
}

func (s QueryCallInTransferRecordResponseBodyDataValues) String() string {
	return tea.Prettify(s)
}

func (s QueryCallInTransferRecordResponseBodyDataValues) GoString() string {
	return s.String()
}

func (s *QueryCallInTransferRecordResponseBodyDataValues) SetCallInCalled(v string) *QueryCallInTransferRecordResponseBodyDataValues {
	s.CallInCalled = &v
	return s
}

func (s *QueryCallInTransferRecordResponseBodyDataValues) SetCallInCaller(v string) *QueryCallInTransferRecordResponseBodyDataValues {
	s.CallInCaller = &v
	return s
}

func (s *QueryCallInTransferRecordResponseBodyDataValues) SetGmtCreate(v int64) *QueryCallInTransferRecordResponseBodyDataValues {
	s.GmtCreate = &v
	return s
}

func (s *QueryCallInTransferRecordResponseBodyDataValues) SetRecordUrl(v string) *QueryCallInTransferRecordResponseBodyDataValues {
	s.RecordUrl = &v
	return s
}

func (s *QueryCallInTransferRecordResponseBodyDataValues) SetTransferCalled(v string) *QueryCallInTransferRecordResponseBodyDataValues {
	s.TransferCalled = &v
	return s
}

func (s *QueryCallInTransferRecordResponseBodyDataValues) SetTransferCaller(v string) *QueryCallInTransferRecordResponseBodyDataValues {
	s.TransferCaller = &v
	return s
}

type QueryCallInTransferRecordResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryCallInTransferRecordResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryCallInTransferRecordResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryCallInTransferRecordResponse) GoString() string {
	return s.String()
}

func (s *QueryCallInTransferRecordResponse) SetHeaders(v map[string]*string) *QueryCallInTransferRecordResponse {
	s.Headers = v
	return s
}

func (s *QueryCallInTransferRecordResponse) SetBody(v *QueryCallInTransferRecordResponseBody) *QueryCallInTransferRecordResponse {
	s.Body = v
	return s
}

type QueryRobotInfoListRequest struct {
	AuditStatus          *string `json:"AuditStatus,omitempty" xml:"AuditStatus,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s QueryRobotInfoListRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryRobotInfoListRequest) GoString() string {
	return s.String()
}

func (s *QueryRobotInfoListRequest) SetAuditStatus(v string) *QueryRobotInfoListRequest {
	s.AuditStatus = &v
	return s
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

type QueryRobotInfoListResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryRobotInfoListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryRobotInfoListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryRobotInfoListResponseBody) SetCode(v string) *QueryRobotInfoListResponseBody {
	s.Code = &v
	return s
}

func (s *QueryRobotInfoListResponseBody) SetData(v string) *QueryRobotInfoListResponseBody {
	s.Data = &v
	return s
}

func (s *QueryRobotInfoListResponseBody) SetMessage(v string) *QueryRobotInfoListResponseBody {
	s.Message = &v
	return s
}

func (s *QueryRobotInfoListResponseBody) SetRequestId(v string) *QueryRobotInfoListResponseBody {
	s.RequestId = &v
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
	Callee               *string `json:"Callee,omitempty" xml:"Callee,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	QueryDate            *int64  `json:"QueryDate,omitempty" xml:"QueryDate,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TaskId               *int64  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s QueryRobotTaskCallDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryRobotTaskCallDetailRequest) GoString() string {
	return s.String()
}

func (s *QueryRobotTaskCallDetailRequest) SetCallee(v string) *QueryRobotTaskCallDetailRequest {
	s.Callee = &v
	return s
}

func (s *QueryRobotTaskCallDetailRequest) SetOwnerId(v int64) *QueryRobotTaskCallDetailRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryRobotTaskCallDetailRequest) SetQueryDate(v int64) *QueryRobotTaskCallDetailRequest {
	s.QueryDate = &v
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

type QueryRobotTaskCallDetailResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryRobotTaskCallDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryRobotTaskCallDetailResponseBody) GoString() string {
	return s.String()
}

func (s *QueryRobotTaskCallDetailResponseBody) SetCode(v string) *QueryRobotTaskCallDetailResponseBody {
	s.Code = &v
	return s
}

func (s *QueryRobotTaskCallDetailResponseBody) SetData(v string) *QueryRobotTaskCallDetailResponseBody {
	s.Data = &v
	return s
}

func (s *QueryRobotTaskCallDetailResponseBody) SetMessage(v string) *QueryRobotTaskCallDetailResponseBody {
	s.Message = &v
	return s
}

func (s *QueryRobotTaskCallDetailResponseBody) SetRequestId(v string) *QueryRobotTaskCallDetailResponseBody {
	s.RequestId = &v
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
	CallResult           *string `json:"CallResult,omitempty" xml:"CallResult,omitempty"`
	Called               *string `json:"Called,omitempty" xml:"Called,omitempty"`
	DialogCountFrom      *string `json:"DialogCountFrom,omitempty" xml:"DialogCountFrom,omitempty"`
	DialogCountTo        *string `json:"DialogCountTo,omitempty" xml:"DialogCountTo,omitempty"`
	DurationFrom         *string `json:"DurationFrom,omitempty" xml:"DurationFrom,omitempty"`
	DurationTo           *string `json:"DurationTo,omitempty" xml:"DurationTo,omitempty"`
	HangupDirection      *string `json:"HangupDirection,omitempty" xml:"HangupDirection,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNo               *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TaskId               *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s QueryRobotTaskCallListRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryRobotTaskCallListRequest) GoString() string {
	return s.String()
}

func (s *QueryRobotTaskCallListRequest) SetCallResult(v string) *QueryRobotTaskCallListRequest {
	s.CallResult = &v
	return s
}

func (s *QueryRobotTaskCallListRequest) SetCalled(v string) *QueryRobotTaskCallListRequest {
	s.Called = &v
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

func (s *QueryRobotTaskCallListRequest) SetDurationFrom(v string) *QueryRobotTaskCallListRequest {
	s.DurationFrom = &v
	return s
}

func (s *QueryRobotTaskCallListRequest) SetDurationTo(v string) *QueryRobotTaskCallListRequest {
	s.DurationTo = &v
	return s
}

func (s *QueryRobotTaskCallListRequest) SetHangupDirection(v string) *QueryRobotTaskCallListRequest {
	s.HangupDirection = &v
	return s
}

func (s *QueryRobotTaskCallListRequest) SetOwnerId(v int64) *QueryRobotTaskCallListRequest {
	s.OwnerId = &v
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

func (s *QueryRobotTaskCallListRequest) SetResourceOwnerAccount(v string) *QueryRobotTaskCallListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryRobotTaskCallListRequest) SetResourceOwnerId(v int64) *QueryRobotTaskCallListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryRobotTaskCallListRequest) SetTaskId(v string) *QueryRobotTaskCallListRequest {
	s.TaskId = &v
	return s
}

type QueryRobotTaskCallListResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryRobotTaskCallListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryRobotTaskCallListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryRobotTaskCallListResponseBody) SetCode(v string) *QueryRobotTaskCallListResponseBody {
	s.Code = &v
	return s
}

func (s *QueryRobotTaskCallListResponseBody) SetData(v string) *QueryRobotTaskCallListResponseBody {
	s.Data = &v
	return s
}

func (s *QueryRobotTaskCallListResponseBody) SetMessage(v string) *QueryRobotTaskCallListResponseBody {
	s.Message = &v
	return s
}

func (s *QueryRobotTaskCallListResponseBody) SetRequestId(v string) *QueryRobotTaskCallListResponseBody {
	s.RequestId = &v
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
	Id                   *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s QueryRobotTaskDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryRobotTaskDetailRequest) GoString() string {
	return s.String()
}

func (s *QueryRobotTaskDetailRequest) SetId(v int64) *QueryRobotTaskDetailRequest {
	s.Id = &v
	return s
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

type QueryRobotTaskDetailResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryRobotTaskDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryRobotTaskDetailResponseBody) GoString() string {
	return s.String()
}

func (s *QueryRobotTaskDetailResponseBody) SetCode(v string) *QueryRobotTaskDetailResponseBody {
	s.Code = &v
	return s
}

func (s *QueryRobotTaskDetailResponseBody) SetData(v string) *QueryRobotTaskDetailResponseBody {
	s.Data = &v
	return s
}

func (s *QueryRobotTaskDetailResponseBody) SetMessage(v string) *QueryRobotTaskDetailResponseBody {
	s.Message = &v
	return s
}

func (s *QueryRobotTaskDetailResponseBody) SetRequestId(v string) *QueryRobotTaskDetailResponseBody {
	s.RequestId = &v
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
	PageNo               *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Status               *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TaskName             *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	Time                 *string `json:"Time,omitempty" xml:"Time,omitempty"`
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

func (s *QueryRobotTaskListRequest) SetPageNo(v int32) *QueryRobotTaskListRequest {
	s.PageNo = &v
	return s
}

func (s *QueryRobotTaskListRequest) SetPageSize(v int32) *QueryRobotTaskListRequest {
	s.PageSize = &v
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

func (s *QueryRobotTaskListRequest) SetStatus(v string) *QueryRobotTaskListRequest {
	s.Status = &v
	return s
}

func (s *QueryRobotTaskListRequest) SetTaskName(v string) *QueryRobotTaskListRequest {
	s.TaskName = &v
	return s
}

func (s *QueryRobotTaskListRequest) SetTime(v string) *QueryRobotTaskListRequest {
	s.Time = &v
	return s
}

type QueryRobotTaskListResponseBody struct {
	Code       *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data       *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message    *string `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNo     *string `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize   *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s QueryRobotTaskListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryRobotTaskListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryRobotTaskListResponseBody) SetCode(v string) *QueryRobotTaskListResponseBody {
	s.Code = &v
	return s
}

func (s *QueryRobotTaskListResponseBody) SetData(v string) *QueryRobotTaskListResponseBody {
	s.Data = &v
	return s
}

func (s *QueryRobotTaskListResponseBody) SetMessage(v string) *QueryRobotTaskListResponseBody {
	s.Message = &v
	return s
}

func (s *QueryRobotTaskListResponseBody) SetPageNo(v string) *QueryRobotTaskListResponseBody {
	s.PageNo = &v
	return s
}

func (s *QueryRobotTaskListResponseBody) SetPageSize(v string) *QueryRobotTaskListResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryRobotTaskListResponseBody) SetRequestId(v string) *QueryRobotTaskListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryRobotTaskListResponseBody) SetTotalCount(v string) *QueryRobotTaskListResponseBody {
	s.TotalCount = &v
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
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryRobotv2AllListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryRobotv2AllListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryRobotv2AllListResponseBody) SetCode(v string) *QueryRobotv2AllListResponseBody {
	s.Code = &v
	return s
}

func (s *QueryRobotv2AllListResponseBody) SetData(v string) *QueryRobotv2AllListResponseBody {
	s.Data = &v
	return s
}

func (s *QueryRobotv2AllListResponseBody) SetMessage(v string) *QueryRobotv2AllListResponseBody {
	s.Message = &v
	return s
}

func (s *QueryRobotv2AllListResponseBody) SetRequestId(v string) *QueryRobotv2AllListResponseBody {
	s.RequestId = &v
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

type QueryVirtualNumberRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNo               *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RouteType            *int32  `json:"RouteType,omitempty" xml:"RouteType,omitempty"`
}

func (s QueryVirtualNumberRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryVirtualNumberRequest) GoString() string {
	return s.String()
}

func (s *QueryVirtualNumberRequest) SetOwnerId(v int64) *QueryVirtualNumberRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryVirtualNumberRequest) SetPageNo(v int32) *QueryVirtualNumberRequest {
	s.PageNo = &v
	return s
}

func (s *QueryVirtualNumberRequest) SetPageSize(v int32) *QueryVirtualNumberRequest {
	s.PageSize = &v
	return s
}

func (s *QueryVirtualNumberRequest) SetProdCode(v string) *QueryVirtualNumberRequest {
	s.ProdCode = &v
	return s
}

func (s *QueryVirtualNumberRequest) SetResourceOwnerAccount(v string) *QueryVirtualNumberRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryVirtualNumberRequest) SetResourceOwnerId(v int64) *QueryVirtualNumberRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryVirtualNumberRequest) SetRouteType(v int32) *QueryVirtualNumberRequest {
	s.RouteType = &v
	return s
}

type QueryVirtualNumberResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryVirtualNumberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryVirtualNumberResponseBody) GoString() string {
	return s.String()
}

func (s *QueryVirtualNumberResponseBody) SetCode(v string) *QueryVirtualNumberResponseBody {
	s.Code = &v
	return s
}

func (s *QueryVirtualNumberResponseBody) SetData(v string) *QueryVirtualNumberResponseBody {
	s.Data = &v
	return s
}

func (s *QueryVirtualNumberResponseBody) SetRequestId(v string) *QueryVirtualNumberResponseBody {
	s.RequestId = &v
	return s
}

type QueryVirtualNumberResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryVirtualNumberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryVirtualNumberResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryVirtualNumberResponse) GoString() string {
	return s.String()
}

func (s *QueryVirtualNumberResponse) SetHeaders(v map[string]*string) *QueryVirtualNumberResponse {
	s.Headers = v
	return s
}

func (s *QueryVirtualNumberResponse) SetBody(v *QueryVirtualNumberResponseBody) *QueryVirtualNumberResponse {
	s.Body = v
	return s
}

type QueryVirtualNumberRelationRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNo               *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PhoneNum             *string `json:"PhoneNum,omitempty" xml:"PhoneNum,omitempty"`
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	QualificationId      *int64  `json:"QualificationId,omitempty" xml:"QualificationId,omitempty"`
	RegionNameCity       *string `json:"RegionNameCity,omitempty" xml:"RegionNameCity,omitempty"`
	RelatedNum           *string `json:"RelatedNum,omitempty" xml:"RelatedNum,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RouteType            *int32  `json:"RouteType,omitempty" xml:"RouteType,omitempty"`
	SpecId               *int64  `json:"SpecId,omitempty" xml:"SpecId,omitempty"`
}

func (s QueryVirtualNumberRelationRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryVirtualNumberRelationRequest) GoString() string {
	return s.String()
}

func (s *QueryVirtualNumberRelationRequest) SetOwnerId(v int64) *QueryVirtualNumberRelationRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryVirtualNumberRelationRequest) SetPageNo(v int32) *QueryVirtualNumberRelationRequest {
	s.PageNo = &v
	return s
}

func (s *QueryVirtualNumberRelationRequest) SetPageSize(v int32) *QueryVirtualNumberRelationRequest {
	s.PageSize = &v
	return s
}

func (s *QueryVirtualNumberRelationRequest) SetPhoneNum(v string) *QueryVirtualNumberRelationRequest {
	s.PhoneNum = &v
	return s
}

func (s *QueryVirtualNumberRelationRequest) SetProdCode(v string) *QueryVirtualNumberRelationRequest {
	s.ProdCode = &v
	return s
}

func (s *QueryVirtualNumberRelationRequest) SetQualificationId(v int64) *QueryVirtualNumberRelationRequest {
	s.QualificationId = &v
	return s
}

func (s *QueryVirtualNumberRelationRequest) SetRegionNameCity(v string) *QueryVirtualNumberRelationRequest {
	s.RegionNameCity = &v
	return s
}

func (s *QueryVirtualNumberRelationRequest) SetRelatedNum(v string) *QueryVirtualNumberRelationRequest {
	s.RelatedNum = &v
	return s
}

func (s *QueryVirtualNumberRelationRequest) SetResourceOwnerAccount(v string) *QueryVirtualNumberRelationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryVirtualNumberRelationRequest) SetResourceOwnerId(v int64) *QueryVirtualNumberRelationRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryVirtualNumberRelationRequest) SetRouteType(v int32) *QueryVirtualNumberRelationRequest {
	s.RouteType = &v
	return s
}

func (s *QueryVirtualNumberRelationRequest) SetSpecId(v int64) *QueryVirtualNumberRelationRequest {
	s.SpecId = &v
	return s
}

type QueryVirtualNumberRelationResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryVirtualNumberRelationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryVirtualNumberRelationResponseBody) GoString() string {
	return s.String()
}

func (s *QueryVirtualNumberRelationResponseBody) SetCode(v string) *QueryVirtualNumberRelationResponseBody {
	s.Code = &v
	return s
}

func (s *QueryVirtualNumberRelationResponseBody) SetData(v string) *QueryVirtualNumberRelationResponseBody {
	s.Data = &v
	return s
}

func (s *QueryVirtualNumberRelationResponseBody) SetRequestId(v string) *QueryVirtualNumberRelationResponseBody {
	s.RequestId = &v
	return s
}

type QueryVirtualNumberRelationResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryVirtualNumberRelationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryVirtualNumberRelationResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryVirtualNumberRelationResponse) GoString() string {
	return s.String()
}

func (s *QueryVirtualNumberRelationResponse) SetHeaders(v map[string]*string) *QueryVirtualNumberRelationResponse {
	s.Headers = v
	return s
}

func (s *QueryVirtualNumberRelationResponse) SetBody(v *QueryVirtualNumberRelationResponseBody) *QueryVirtualNumberRelationResponse {
	s.Body = v
	return s
}

type RefreshMqttTokenRequest struct {
	ClientId             *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s RefreshMqttTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s RefreshMqttTokenRequest) GoString() string {
	return s.String()
}

func (s *RefreshMqttTokenRequest) SetClientId(v string) *RefreshMqttTokenRequest {
	s.ClientId = &v
	return s
}

func (s *RefreshMqttTokenRequest) SetOwnerId(v int64) *RefreshMqttTokenRequest {
	s.OwnerId = &v
	return s
}

func (s *RefreshMqttTokenRequest) SetResourceOwnerAccount(v string) *RefreshMqttTokenRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RefreshMqttTokenRequest) SetResourceOwnerId(v int64) *RefreshMqttTokenRequest {
	s.ResourceOwnerId = &v
	return s
}

type RefreshMqttTokenResponseBody struct {
	Code      *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *RefreshMqttTokenResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RefreshMqttTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RefreshMqttTokenResponseBody) GoString() string {
	return s.String()
}

func (s *RefreshMqttTokenResponseBody) SetCode(v string) *RefreshMqttTokenResponseBody {
	s.Code = &v
	return s
}

func (s *RefreshMqttTokenResponseBody) SetData(v *RefreshMqttTokenResponseBodyData) *RefreshMqttTokenResponseBody {
	s.Data = v
	return s
}

func (s *RefreshMqttTokenResponseBody) SetMessage(v string) *RefreshMqttTokenResponseBody {
	s.Message = &v
	return s
}

func (s *RefreshMqttTokenResponseBody) SetRequestId(v string) *RefreshMqttTokenResponseBody {
	s.RequestId = &v
	return s
}

type RefreshMqttTokenResponseBodyData struct {
	ClientId   *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	Host       *string `json:"Host,omitempty" xml:"Host,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	P2pTopic   *string `json:"P2pTopic,omitempty" xml:"P2pTopic,omitempty"`
	ServerId   *string `json:"ServerId,omitempty" xml:"ServerId,omitempty"`
	Token      *string `json:"Token,omitempty" xml:"Token,omitempty"`
	Username   *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s RefreshMqttTokenResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RefreshMqttTokenResponseBodyData) GoString() string {
	return s.String()
}

func (s *RefreshMqttTokenResponseBodyData) SetClientId(v string) *RefreshMqttTokenResponseBodyData {
	s.ClientId = &v
	return s
}

func (s *RefreshMqttTokenResponseBodyData) SetExpireTime(v string) *RefreshMqttTokenResponseBodyData {
	s.ExpireTime = &v
	return s
}

func (s *RefreshMqttTokenResponseBodyData) SetHost(v string) *RefreshMqttTokenResponseBodyData {
	s.Host = &v
	return s
}

func (s *RefreshMqttTokenResponseBodyData) SetInstanceId(v string) *RefreshMqttTokenResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *RefreshMqttTokenResponseBodyData) SetP2pTopic(v string) *RefreshMqttTokenResponseBodyData {
	s.P2pTopic = &v
	return s
}

func (s *RefreshMqttTokenResponseBodyData) SetServerId(v string) *RefreshMqttTokenResponseBodyData {
	s.ServerId = &v
	return s
}

func (s *RefreshMqttTokenResponseBodyData) SetToken(v string) *RefreshMqttTokenResponseBodyData {
	s.Token = &v
	return s
}

func (s *RefreshMqttTokenResponseBodyData) SetUsername(v string) *RefreshMqttTokenResponseBodyData {
	s.Username = &v
	return s
}

type RefreshMqttTokenResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RefreshMqttTokenResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RefreshMqttTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s RefreshMqttTokenResponse) GoString() string {
	return s.String()
}

func (s *RefreshMqttTokenResponse) SetHeaders(v map[string]*string) *RefreshMqttTokenResponse {
	s.Headers = v
	return s
}

func (s *RefreshMqttTokenResponse) SetBody(v *RefreshMqttTokenResponseBody) *RefreshMqttTokenResponse {
	s.Body = v
	return s
}

type SendVerificationRequest struct {
	BizType              *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Target               *string `json:"Target,omitempty" xml:"Target,omitempty"`
	VerifyType           *string `json:"VerifyType,omitempty" xml:"VerifyType,omitempty"`
}

func (s SendVerificationRequest) String() string {
	return tea.Prettify(s)
}

func (s SendVerificationRequest) GoString() string {
	return s.String()
}

func (s *SendVerificationRequest) SetBizType(v string) *SendVerificationRequest {
	s.BizType = &v
	return s
}

func (s *SendVerificationRequest) SetOwnerId(v int64) *SendVerificationRequest {
	s.OwnerId = &v
	return s
}

func (s *SendVerificationRequest) SetResourceOwnerAccount(v string) *SendVerificationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SendVerificationRequest) SetResourceOwnerId(v int64) *SendVerificationRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SendVerificationRequest) SetTarget(v string) *SendVerificationRequest {
	s.Target = &v
	return s
}

func (s *SendVerificationRequest) SetVerifyType(v string) *SendVerificationRequest {
	s.VerifyType = &v
	return s
}

type SendVerificationResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SendVerificationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendVerificationResponseBody) GoString() string {
	return s.String()
}

func (s *SendVerificationResponseBody) SetCode(v string) *SendVerificationResponseBody {
	s.Code = &v
	return s
}

func (s *SendVerificationResponseBody) SetData(v bool) *SendVerificationResponseBody {
	s.Data = &v
	return s
}

func (s *SendVerificationResponseBody) SetMessage(v string) *SendVerificationResponseBody {
	s.Message = &v
	return s
}

func (s *SendVerificationResponseBody) SetRequestId(v string) *SendVerificationResponseBody {
	s.RequestId = &v
	return s
}

type SendVerificationResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SendVerificationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SendVerificationResponse) String() string {
	return tea.Prettify(s)
}

func (s SendVerificationResponse) GoString() string {
	return s.String()
}

func (s *SendVerificationResponse) SetHeaders(v map[string]*string) *SendVerificationResponse {
	s.Headers = v
	return s
}

func (s *SendVerificationResponse) SetBody(v *SendVerificationResponseBody) *SendVerificationResponse {
	s.Body = v
	return s
}

type SetTransferCalleePoolConfigRequest struct {
	CalledRouteMode      *string                                      `json:"CalledRouteMode,omitempty" xml:"CalledRouteMode,omitempty"`
	Details              []*SetTransferCalleePoolConfigRequestDetails `json:"Details,omitempty" xml:"Details,omitempty" type:"Repeated"`
	OwnerId              *int64                                       `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PhoneNumber          *string                                      `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	QualificationId      *string                                      `json:"QualificationId,omitempty" xml:"QualificationId,omitempty"`
	ResourceOwnerAccount *string                                      `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64                                       `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s SetTransferCalleePoolConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s SetTransferCalleePoolConfigRequest) GoString() string {
	return s.String()
}

func (s *SetTransferCalleePoolConfigRequest) SetCalledRouteMode(v string) *SetTransferCalleePoolConfigRequest {
	s.CalledRouteMode = &v
	return s
}

func (s *SetTransferCalleePoolConfigRequest) SetDetails(v []*SetTransferCalleePoolConfigRequestDetails) *SetTransferCalleePoolConfigRequest {
	s.Details = v
	return s
}

func (s *SetTransferCalleePoolConfigRequest) SetOwnerId(v int64) *SetTransferCalleePoolConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *SetTransferCalleePoolConfigRequest) SetPhoneNumber(v string) *SetTransferCalleePoolConfigRequest {
	s.PhoneNumber = &v
	return s
}

func (s *SetTransferCalleePoolConfigRequest) SetQualificationId(v string) *SetTransferCalleePoolConfigRequest {
	s.QualificationId = &v
	return s
}

func (s *SetTransferCalleePoolConfigRequest) SetResourceOwnerAccount(v string) *SetTransferCalleePoolConfigRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SetTransferCalleePoolConfigRequest) SetResourceOwnerId(v int64) *SetTransferCalleePoolConfigRequest {
	s.ResourceOwnerId = &v
	return s
}

type SetTransferCalleePoolConfigRequestDetails struct {
	Called *string `json:"Called,omitempty" xml:"Called,omitempty"`
	Caller *string `json:"Caller,omitempty" xml:"Caller,omitempty"`
}

func (s SetTransferCalleePoolConfigRequestDetails) String() string {
	return tea.Prettify(s)
}

func (s SetTransferCalleePoolConfigRequestDetails) GoString() string {
	return s.String()
}

func (s *SetTransferCalleePoolConfigRequestDetails) SetCalled(v string) *SetTransferCalleePoolConfigRequestDetails {
	s.Called = &v
	return s
}

func (s *SetTransferCalleePoolConfigRequestDetails) SetCaller(v string) *SetTransferCalleePoolConfigRequestDetails {
	s.Caller = &v
	return s
}

type SetTransferCalleePoolConfigResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetTransferCalleePoolConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetTransferCalleePoolConfigResponseBody) GoString() string {
	return s.String()
}

func (s *SetTransferCalleePoolConfigResponseBody) SetCode(v string) *SetTransferCalleePoolConfigResponseBody {
	s.Code = &v
	return s
}

func (s *SetTransferCalleePoolConfigResponseBody) SetData(v bool) *SetTransferCalleePoolConfigResponseBody {
	s.Data = &v
	return s
}

func (s *SetTransferCalleePoolConfigResponseBody) SetMessage(v string) *SetTransferCalleePoolConfigResponseBody {
	s.Message = &v
	return s
}

func (s *SetTransferCalleePoolConfigResponseBody) SetRequestId(v string) *SetTransferCalleePoolConfigResponseBody {
	s.RequestId = &v
	return s
}

type SetTransferCalleePoolConfigResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetTransferCalleePoolConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetTransferCalleePoolConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s SetTransferCalleePoolConfigResponse) GoString() string {
	return s.String()
}

func (s *SetTransferCalleePoolConfigResponse) SetHeaders(v map[string]*string) *SetTransferCalleePoolConfigResponse {
	s.Headers = v
	return s
}

func (s *SetTransferCalleePoolConfigResponse) SetBody(v *SetTransferCalleePoolConfigResponseBody) *SetTransferCalleePoolConfigResponse {
	s.Body = v
	return s
}

type SingleCallByTtsRequest struct {
	CalledNumber         *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	CalledShowNumber     *string `json:"CalledShowNumber,omitempty" xml:"CalledShowNumber,omitempty"`
	OutId                *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PlayTimes            *int32  `json:"PlayTimes,omitempty" xml:"PlayTimes,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Speed                *int32  `json:"Speed,omitempty" xml:"Speed,omitempty"`
	TtsCode              *string `json:"TtsCode,omitempty" xml:"TtsCode,omitempty"`
	TtsParam             *string `json:"TtsParam,omitempty" xml:"TtsParam,omitempty"`
	Volume               *int32  `json:"Volume,omitempty" xml:"Volume,omitempty"`
}

func (s SingleCallByTtsRequest) String() string {
	return tea.Prettify(s)
}

func (s SingleCallByTtsRequest) GoString() string {
	return s.String()
}

func (s *SingleCallByTtsRequest) SetCalledNumber(v string) *SingleCallByTtsRequest {
	s.CalledNumber = &v
	return s
}

func (s *SingleCallByTtsRequest) SetCalledShowNumber(v string) *SingleCallByTtsRequest {
	s.CalledShowNumber = &v
	return s
}

func (s *SingleCallByTtsRequest) SetOutId(v string) *SingleCallByTtsRequest {
	s.OutId = &v
	return s
}

func (s *SingleCallByTtsRequest) SetOwnerId(v int64) *SingleCallByTtsRequest {
	s.OwnerId = &v
	return s
}

func (s *SingleCallByTtsRequest) SetPlayTimes(v int32) *SingleCallByTtsRequest {
	s.PlayTimes = &v
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

func (s *SingleCallByTtsRequest) SetSpeed(v int32) *SingleCallByTtsRequest {
	s.Speed = &v
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

func (s *SingleCallByTtsRequest) SetVolume(v int32) *SingleCallByTtsRequest {
	s.Volume = &v
	return s
}

type SingleCallByTtsResponseBody struct {
	CallId    *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SingleCallByTtsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SingleCallByTtsResponseBody) GoString() string {
	return s.String()
}

func (s *SingleCallByTtsResponseBody) SetCallId(v string) *SingleCallByTtsResponseBody {
	s.CallId = &v
	return s
}

func (s *SingleCallByTtsResponseBody) SetCode(v string) *SingleCallByTtsResponseBody {
	s.Code = &v
	return s
}

func (s *SingleCallByTtsResponseBody) SetMessage(v string) *SingleCallByTtsResponseBody {
	s.Message = &v
	return s
}

func (s *SingleCallByTtsResponseBody) SetRequestId(v string) *SingleCallByTtsResponseBody {
	s.RequestId = &v
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
	CalledNumber         *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	CalledShowNumber     *string `json:"CalledShowNumber,omitempty" xml:"CalledShowNumber,omitempty"`
	OutId                *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PlayTimes            *int32  `json:"PlayTimes,omitempty" xml:"PlayTimes,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Speed                *int32  `json:"Speed,omitempty" xml:"Speed,omitempty"`
	VoiceCode            *string `json:"VoiceCode,omitempty" xml:"VoiceCode,omitempty"`
	Volume               *int32  `json:"Volume,omitempty" xml:"Volume,omitempty"`
}

func (s SingleCallByVoiceRequest) String() string {
	return tea.Prettify(s)
}

func (s SingleCallByVoiceRequest) GoString() string {
	return s.String()
}

func (s *SingleCallByVoiceRequest) SetCalledNumber(v string) *SingleCallByVoiceRequest {
	s.CalledNumber = &v
	return s
}

func (s *SingleCallByVoiceRequest) SetCalledShowNumber(v string) *SingleCallByVoiceRequest {
	s.CalledShowNumber = &v
	return s
}

func (s *SingleCallByVoiceRequest) SetOutId(v string) *SingleCallByVoiceRequest {
	s.OutId = &v
	return s
}

func (s *SingleCallByVoiceRequest) SetOwnerId(v int64) *SingleCallByVoiceRequest {
	s.OwnerId = &v
	return s
}

func (s *SingleCallByVoiceRequest) SetPlayTimes(v int32) *SingleCallByVoiceRequest {
	s.PlayTimes = &v
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

func (s *SingleCallByVoiceRequest) SetSpeed(v int32) *SingleCallByVoiceRequest {
	s.Speed = &v
	return s
}

func (s *SingleCallByVoiceRequest) SetVoiceCode(v string) *SingleCallByVoiceRequest {
	s.VoiceCode = &v
	return s
}

func (s *SingleCallByVoiceRequest) SetVolume(v int32) *SingleCallByVoiceRequest {
	s.Volume = &v
	return s
}

type SingleCallByVoiceResponseBody struct {
	CallId    *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SingleCallByVoiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SingleCallByVoiceResponseBody) GoString() string {
	return s.String()
}

func (s *SingleCallByVoiceResponseBody) SetCallId(v string) *SingleCallByVoiceResponseBody {
	s.CallId = &v
	return s
}

func (s *SingleCallByVoiceResponseBody) SetCode(v string) *SingleCallByVoiceResponseBody {
	s.Code = &v
	return s
}

func (s *SingleCallByVoiceResponseBody) SetMessage(v string) *SingleCallByVoiceResponseBody {
	s.Message = &v
	return s
}

func (s *SingleCallByVoiceResponseBody) SetRequestId(v string) *SingleCallByVoiceResponseBody {
	s.RequestId = &v
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
	ActionCodeBreak      *bool   `json:"ActionCodeBreak,omitempty" xml:"ActionCodeBreak,omitempty"`
	ActionCodeTimeBreak  *int32  `json:"ActionCodeTimeBreak,omitempty" xml:"ActionCodeTimeBreak,omitempty"`
	AsrBaseId            *string `json:"AsrBaseId,omitempty" xml:"AsrBaseId,omitempty"`
	AsrModelId           *string `json:"AsrModelId,omitempty" xml:"AsrModelId,omitempty"`
	BackgroundFileCode   *string `json:"BackgroundFileCode,omitempty" xml:"BackgroundFileCode,omitempty"`
	BackgroundSpeed      *int32  `json:"BackgroundSpeed,omitempty" xml:"BackgroundSpeed,omitempty"`
	BackgroundVolume     *int32  `json:"BackgroundVolume,omitempty" xml:"BackgroundVolume,omitempty"`
	CalledNumber         *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	CalledShowNumber     *string `json:"CalledShowNumber,omitempty" xml:"CalledShowNumber,omitempty"`
	DynamicId            *string `json:"DynamicId,omitempty" xml:"DynamicId,omitempty"`
	EarlyMediaAsr        *bool   `json:"EarlyMediaAsr,omitempty" xml:"EarlyMediaAsr,omitempty"`
	EnableITN            *bool   `json:"EnableITN,omitempty" xml:"EnableITN,omitempty"`
	MuteTime             *int32  `json:"MuteTime,omitempty" xml:"MuteTime,omitempty"`
	OutId                *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PauseTime            *int32  `json:"PauseTime,omitempty" xml:"PauseTime,omitempty"`
	RecordFlag           *bool   `json:"RecordFlag,omitempty" xml:"RecordFlag,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SessionTimeout       *int32  `json:"SessionTimeout,omitempty" xml:"SessionTimeout,omitempty"`
	Speed                *int32  `json:"Speed,omitempty" xml:"Speed,omitempty"`
	StreamAsr            *int32  `json:"StreamAsr,omitempty" xml:"StreamAsr,omitempty"`
	TtsConf              *bool   `json:"TtsConf,omitempty" xml:"TtsConf,omitempty"`
	TtsSpeed             *int32  `json:"TtsSpeed,omitempty" xml:"TtsSpeed,omitempty"`
	TtsStyle             *string `json:"TtsStyle,omitempty" xml:"TtsStyle,omitempty"`
	TtsVolume            *int32  `json:"TtsVolume,omitempty" xml:"TtsVolume,omitempty"`
	VoiceCode            *string `json:"VoiceCode,omitempty" xml:"VoiceCode,omitempty"`
	VoiceCodeParam       *string `json:"VoiceCodeParam,omitempty" xml:"VoiceCodeParam,omitempty"`
	Volume               *int32  `json:"Volume,omitempty" xml:"Volume,omitempty"`
}

func (s SmartCallRequest) String() string {
	return tea.Prettify(s)
}

func (s SmartCallRequest) GoString() string {
	return s.String()
}

func (s *SmartCallRequest) SetActionCodeBreak(v bool) *SmartCallRequest {
	s.ActionCodeBreak = &v
	return s
}

func (s *SmartCallRequest) SetActionCodeTimeBreak(v int32) *SmartCallRequest {
	s.ActionCodeTimeBreak = &v
	return s
}

func (s *SmartCallRequest) SetAsrBaseId(v string) *SmartCallRequest {
	s.AsrBaseId = &v
	return s
}

func (s *SmartCallRequest) SetAsrModelId(v string) *SmartCallRequest {
	s.AsrModelId = &v
	return s
}

func (s *SmartCallRequest) SetBackgroundFileCode(v string) *SmartCallRequest {
	s.BackgroundFileCode = &v
	return s
}

func (s *SmartCallRequest) SetBackgroundSpeed(v int32) *SmartCallRequest {
	s.BackgroundSpeed = &v
	return s
}

func (s *SmartCallRequest) SetBackgroundVolume(v int32) *SmartCallRequest {
	s.BackgroundVolume = &v
	return s
}

func (s *SmartCallRequest) SetCalledNumber(v string) *SmartCallRequest {
	s.CalledNumber = &v
	return s
}

func (s *SmartCallRequest) SetCalledShowNumber(v string) *SmartCallRequest {
	s.CalledShowNumber = &v
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

func (s *SmartCallRequest) SetEnableITN(v bool) *SmartCallRequest {
	s.EnableITN = &v
	return s
}

func (s *SmartCallRequest) SetMuteTime(v int32) *SmartCallRequest {
	s.MuteTime = &v
	return s
}

func (s *SmartCallRequest) SetOutId(v string) *SmartCallRequest {
	s.OutId = &v
	return s
}

func (s *SmartCallRequest) SetOwnerId(v int64) *SmartCallRequest {
	s.OwnerId = &v
	return s
}

func (s *SmartCallRequest) SetPauseTime(v int32) *SmartCallRequest {
	s.PauseTime = &v
	return s
}

func (s *SmartCallRequest) SetRecordFlag(v bool) *SmartCallRequest {
	s.RecordFlag = &v
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

func (s *SmartCallRequest) SetSessionTimeout(v int32) *SmartCallRequest {
	s.SessionTimeout = &v
	return s
}

func (s *SmartCallRequest) SetSpeed(v int32) *SmartCallRequest {
	s.Speed = &v
	return s
}

func (s *SmartCallRequest) SetStreamAsr(v int32) *SmartCallRequest {
	s.StreamAsr = &v
	return s
}

func (s *SmartCallRequest) SetTtsConf(v bool) *SmartCallRequest {
	s.TtsConf = &v
	return s
}

func (s *SmartCallRequest) SetTtsSpeed(v int32) *SmartCallRequest {
	s.TtsSpeed = &v
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

func (s *SmartCallRequest) SetVoiceCode(v string) *SmartCallRequest {
	s.VoiceCode = &v
	return s
}

func (s *SmartCallRequest) SetVoiceCodeParam(v string) *SmartCallRequest {
	s.VoiceCodeParam = &v
	return s
}

func (s *SmartCallRequest) SetVolume(v int32) *SmartCallRequest {
	s.Volume = &v
	return s
}

type SmartCallResponseBody struct {
	CallId    *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SmartCallResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SmartCallResponseBody) GoString() string {
	return s.String()
}

func (s *SmartCallResponseBody) SetCallId(v string) *SmartCallResponseBody {
	s.CallId = &v
	return s
}

func (s *SmartCallResponseBody) SetCode(v string) *SmartCallResponseBody {
	s.Code = &v
	return s
}

func (s *SmartCallResponseBody) SetMessage(v string) *SmartCallResponseBody {
	s.Message = &v
	return s
}

func (s *SmartCallResponseBody) SetRequestId(v string) *SmartCallResponseBody {
	s.RequestId = &v
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
	CallId               *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	Command              *string `json:"Command,omitempty" xml:"Command,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Param                *string `json:"Param,omitempty" xml:"Param,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s SmartCallOperateRequest) String() string {
	return tea.Prettify(s)
}

func (s SmartCallOperateRequest) GoString() string {
	return s.String()
}

func (s *SmartCallOperateRequest) SetCallId(v string) *SmartCallOperateRequest {
	s.CallId = &v
	return s
}

func (s *SmartCallOperateRequest) SetCommand(v string) *SmartCallOperateRequest {
	s.Command = &v
	return s
}

func (s *SmartCallOperateRequest) SetOwnerId(v int64) *SmartCallOperateRequest {
	s.OwnerId = &v
	return s
}

func (s *SmartCallOperateRequest) SetParam(v string) *SmartCallOperateRequest {
	s.Param = &v
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

type SmartCallOperateResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status    *bool   `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s SmartCallOperateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SmartCallOperateResponseBody) GoString() string {
	return s.String()
}

func (s *SmartCallOperateResponseBody) SetCode(v string) *SmartCallOperateResponseBody {
	s.Code = &v
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

func (s *SmartCallOperateResponseBody) SetStatus(v bool) *SmartCallOperateResponseBody {
	s.Status = &v
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

type StartRobotTaskRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ScheduleTime         *string `json:"ScheduleTime,omitempty" xml:"ScheduleTime,omitempty"`
	TaskId               *int64  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
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

func (s *StartRobotTaskRequest) SetScheduleTime(v string) *StartRobotTaskRequest {
	s.ScheduleTime = &v
	return s
}

func (s *StartRobotTaskRequest) SetTaskId(v int64) *StartRobotTaskRequest {
	s.TaskId = &v
	return s
}

type StartRobotTaskResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartRobotTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartRobotTaskResponseBody) GoString() string {
	return s.String()
}

func (s *StartRobotTaskResponseBody) SetCode(v string) *StartRobotTaskResponseBody {
	s.Code = &v
	return s
}

func (s *StartRobotTaskResponseBody) SetData(v string) *StartRobotTaskResponseBody {
	s.Data = &v
	return s
}

func (s *StartRobotTaskResponseBody) SetMessage(v string) *StartRobotTaskResponseBody {
	s.Message = &v
	return s
}

func (s *StartRobotTaskResponseBody) SetRequestId(v string) *StartRobotTaskResponseBody {
	s.RequestId = &v
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
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopRobotTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopRobotTaskResponseBody) GoString() string {
	return s.String()
}

func (s *StopRobotTaskResponseBody) SetCode(v string) *StopRobotTaskResponseBody {
	s.Code = &v
	return s
}

func (s *StopRobotTaskResponseBody) SetData(v string) *StopRobotTaskResponseBody {
	s.Data = &v
	return s
}

func (s *StopRobotTaskResponseBody) SetMessage(v string) *StopRobotTaskResponseBody {
	s.Message = &v
	return s
}

func (s *StopRobotTaskResponseBody) SetRequestId(v string) *StopRobotTaskResponseBody {
	s.RequestId = &v
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

type SubmitHotlineTransferRegisterRequest struct {
	Agreement                *string                                                         `json:"Agreement,omitempty" xml:"Agreement,omitempty"`
	HotlineNumber            *string                                                         `json:"HotlineNumber,omitempty" xml:"HotlineNumber,omitempty"`
	OperatorIdentityCard     *string                                                         `json:"OperatorIdentityCard,omitempty" xml:"OperatorIdentityCard,omitempty"`
	OperatorMail             *string                                                         `json:"OperatorMail,omitempty" xml:"OperatorMail,omitempty"`
	OperatorMailVerifyCode   *string                                                         `json:"OperatorMailVerifyCode,omitempty" xml:"OperatorMailVerifyCode,omitempty"`
	OperatorMobile           *string                                                         `json:"OperatorMobile,omitempty" xml:"OperatorMobile,omitempty"`
	OperatorMobileVerifyCode *string                                                         `json:"OperatorMobileVerifyCode,omitempty" xml:"OperatorMobileVerifyCode,omitempty"`
	OperatorName             *string                                                         `json:"OperatorName,omitempty" xml:"OperatorName,omitempty"`
	OwnerId                  *int64                                                          `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	QualificationId          *string                                                         `json:"QualificationId,omitempty" xml:"QualificationId,omitempty"`
	ResourceOwnerAccount     *string                                                         `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId          *int64                                                          `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TransferPhoneNumberInfos []*SubmitHotlineTransferRegisterRequestTransferPhoneNumberInfos `json:"TransferPhoneNumberInfos,omitempty" xml:"TransferPhoneNumberInfos,omitempty" type:"Repeated"`
}

func (s SubmitHotlineTransferRegisterRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitHotlineTransferRegisterRequest) GoString() string {
	return s.String()
}

func (s *SubmitHotlineTransferRegisterRequest) SetAgreement(v string) *SubmitHotlineTransferRegisterRequest {
	s.Agreement = &v
	return s
}

func (s *SubmitHotlineTransferRegisterRequest) SetHotlineNumber(v string) *SubmitHotlineTransferRegisterRequest {
	s.HotlineNumber = &v
	return s
}

func (s *SubmitHotlineTransferRegisterRequest) SetOperatorIdentityCard(v string) *SubmitHotlineTransferRegisterRequest {
	s.OperatorIdentityCard = &v
	return s
}

func (s *SubmitHotlineTransferRegisterRequest) SetOperatorMail(v string) *SubmitHotlineTransferRegisterRequest {
	s.OperatorMail = &v
	return s
}

func (s *SubmitHotlineTransferRegisterRequest) SetOperatorMailVerifyCode(v string) *SubmitHotlineTransferRegisterRequest {
	s.OperatorMailVerifyCode = &v
	return s
}

func (s *SubmitHotlineTransferRegisterRequest) SetOperatorMobile(v string) *SubmitHotlineTransferRegisterRequest {
	s.OperatorMobile = &v
	return s
}

func (s *SubmitHotlineTransferRegisterRequest) SetOperatorMobileVerifyCode(v string) *SubmitHotlineTransferRegisterRequest {
	s.OperatorMobileVerifyCode = &v
	return s
}

func (s *SubmitHotlineTransferRegisterRequest) SetOperatorName(v string) *SubmitHotlineTransferRegisterRequest {
	s.OperatorName = &v
	return s
}

func (s *SubmitHotlineTransferRegisterRequest) SetOwnerId(v int64) *SubmitHotlineTransferRegisterRequest {
	s.OwnerId = &v
	return s
}

func (s *SubmitHotlineTransferRegisterRequest) SetQualificationId(v string) *SubmitHotlineTransferRegisterRequest {
	s.QualificationId = &v
	return s
}

func (s *SubmitHotlineTransferRegisterRequest) SetResourceOwnerAccount(v string) *SubmitHotlineTransferRegisterRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SubmitHotlineTransferRegisterRequest) SetResourceOwnerId(v int64) *SubmitHotlineTransferRegisterRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SubmitHotlineTransferRegisterRequest) SetTransferPhoneNumberInfos(v []*SubmitHotlineTransferRegisterRequestTransferPhoneNumberInfos) *SubmitHotlineTransferRegisterRequest {
	s.TransferPhoneNumberInfos = v
	return s
}

type SubmitHotlineTransferRegisterRequestTransferPhoneNumberInfos struct {
	IdentityCard         *string `json:"IdentityCard,omitempty" xml:"IdentityCard,omitempty"`
	PhoneNumber          *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	PhoneNumberOwnerName *string `json:"PhoneNumberOwnerName,omitempty" xml:"PhoneNumberOwnerName,omitempty"`
}

func (s SubmitHotlineTransferRegisterRequestTransferPhoneNumberInfos) String() string {
	return tea.Prettify(s)
}

func (s SubmitHotlineTransferRegisterRequestTransferPhoneNumberInfos) GoString() string {
	return s.String()
}

func (s *SubmitHotlineTransferRegisterRequestTransferPhoneNumberInfos) SetIdentityCard(v string) *SubmitHotlineTransferRegisterRequestTransferPhoneNumberInfos {
	s.IdentityCard = &v
	return s
}

func (s *SubmitHotlineTransferRegisterRequestTransferPhoneNumberInfos) SetPhoneNumber(v string) *SubmitHotlineTransferRegisterRequestTransferPhoneNumberInfos {
	s.PhoneNumber = &v
	return s
}

func (s *SubmitHotlineTransferRegisterRequestTransferPhoneNumberInfos) SetPhoneNumberOwnerName(v string) *SubmitHotlineTransferRegisterRequestTransferPhoneNumberInfos {
	s.PhoneNumberOwnerName = &v
	return s
}

type SubmitHotlineTransferRegisterResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *int64  `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitHotlineTransferRegisterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitHotlineTransferRegisterResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitHotlineTransferRegisterResponseBody) SetCode(v string) *SubmitHotlineTransferRegisterResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitHotlineTransferRegisterResponseBody) SetData(v int64) *SubmitHotlineTransferRegisterResponseBody {
	s.Data = &v
	return s
}

func (s *SubmitHotlineTransferRegisterResponseBody) SetMessage(v string) *SubmitHotlineTransferRegisterResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitHotlineTransferRegisterResponseBody) SetRequestId(v string) *SubmitHotlineTransferRegisterResponseBody {
	s.RequestId = &v
	return s
}

type SubmitHotlineTransferRegisterResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SubmitHotlineTransferRegisterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubmitHotlineTransferRegisterResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitHotlineTransferRegisterResponse) GoString() string {
	return s.String()
}

func (s *SubmitHotlineTransferRegisterResponse) SetHeaders(v map[string]*string) *SubmitHotlineTransferRegisterResponse {
	s.Headers = v
	return s
}

func (s *SubmitHotlineTransferRegisterResponse) SetBody(v *SubmitHotlineTransferRegisterResponseBody) *SubmitHotlineTransferRegisterResponse {
	s.Body = v
	return s
}

type UploadRobotTaskCalledFileRequest struct {
	CalledNumber         *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	Id                   *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TtsParam             *string `json:"TtsParam,omitempty" xml:"TtsParam,omitempty"`
	TtsParamHead         *string `json:"TtsParamHead,omitempty" xml:"TtsParamHead,omitempty"`
}

func (s UploadRobotTaskCalledFileRequest) String() string {
	return tea.Prettify(s)
}

func (s UploadRobotTaskCalledFileRequest) GoString() string {
	return s.String()
}

func (s *UploadRobotTaskCalledFileRequest) SetCalledNumber(v string) *UploadRobotTaskCalledFileRequest {
	s.CalledNumber = &v
	return s
}

func (s *UploadRobotTaskCalledFileRequest) SetId(v int64) *UploadRobotTaskCalledFileRequest {
	s.Id = &v
	return s
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

func (s *UploadRobotTaskCalledFileRequest) SetTtsParam(v string) *UploadRobotTaskCalledFileRequest {
	s.TtsParam = &v
	return s
}

func (s *UploadRobotTaskCalledFileRequest) SetTtsParamHead(v string) *UploadRobotTaskCalledFileRequest {
	s.TtsParamHead = &v
	return s
}

type UploadRobotTaskCalledFileResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UploadRobotTaskCalledFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UploadRobotTaskCalledFileResponseBody) GoString() string {
	return s.String()
}

func (s *UploadRobotTaskCalledFileResponseBody) SetCode(v string) *UploadRobotTaskCalledFileResponseBody {
	s.Code = &v
	return s
}

func (s *UploadRobotTaskCalledFileResponseBody) SetData(v string) *UploadRobotTaskCalledFileResponseBody {
	s.Data = &v
	return s
}

func (s *UploadRobotTaskCalledFileResponseBody) SetMessage(v string) *UploadRobotTaskCalledFileResponseBody {
	s.Message = &v
	return s
}

func (s *UploadRobotTaskCalledFileResponseBody) SetRequestId(v string) *UploadRobotTaskCalledFileResponseBody {
	s.RequestId = &v
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
	query := map[string]interface{}{}
	query["DeviceId"] = request.DeviceId
	query["OwnerId"] = request.OwnerId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("AddRtcAccount"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &AddRtcAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) AddVirtualNumberRelationWithOptions(request *AddVirtualNumberRelationRequest, runtime *util.RuntimeOptions) (_result *AddVirtualNumberRelationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["CorpNameList"] = request.CorpNameList
	query["NumberList"] = request.NumberList
	query["OwnerId"] = request.OwnerId
	query["PhoneNum"] = request.PhoneNum
	query["ProdCode"] = request.ProdCode
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["RouteType"] = request.RouteType
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("AddVirtualNumberRelation"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &AddVirtualNumberRelationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddVirtualNumberRelation(request *AddVirtualNumberRelationRequest) (_result *AddVirtualNumberRelationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddVirtualNumberRelationResponse{}
	_body, _err := client.AddVirtualNumberRelationWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	query["CalledNumber"] = request.CalledNumber
	query["CalledShowNumber"] = request.CalledShowNumber
	query["CorpName"] = request.CorpName
	query["DialogId"] = request.DialogId
	query["EarlyMediaAsr"] = request.EarlyMediaAsr
	query["IsSelfLine"] = request.IsSelfLine
	query["OwnerId"] = request.OwnerId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["ScheduleCall"] = request.ScheduleCall
	query["ScheduleTime"] = request.ScheduleTime
	query["TaskName"] = request.TaskName
	query["TtsParam"] = request.TtsParam
	query["TtsParamHead"] = request.TtsParamHead
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchRobotSmartCall"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchRobotSmartCallResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) CancelCallWithOptions(request *CancelCallRequest, runtime *util.RuntimeOptions) (_result *CancelCallResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["CallId"] = request.CallId
	query["OwnerId"] = request.OwnerId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("CancelCall"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CancelCallResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	query["OwnerId"] = request.OwnerId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["TaskId"] = request.TaskId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("CancelOrderRobotTask"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CancelOrderRobotTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	query["OwnerId"] = request.OwnerId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["TaskId"] = request.TaskId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("CancelRobotTask"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CancelRobotTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	query["AsrFlag"] = request.AsrFlag
	query["AsrModelId"] = request.AsrModelId
	query["CalledNumber"] = request.CalledNumber
	query["CalledShowNumber"] = request.CalledShowNumber
	query["CallerNumber"] = request.CallerNumber
	query["CallerShowNumber"] = request.CallerShowNumber
	query["OutId"] = request.OutId
	query["OwnerId"] = request.OwnerId
	query["RecordFlag"] = request.RecordFlag
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["SessionTimeout"] = request.SessionTimeout
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("ClickToDial"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ClickToDialResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) CreateCallTaskWithOptions(request *CreateCallTaskRequest, runtime *util.RuntimeOptions) (_result *CreateCallTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["BizType"] = request.BizType
	query["Data"] = request.Data
	query["DataType"] = request.DataType
	query["FireTime"] = request.FireTime
	query["OwnerId"] = request.OwnerId
	query["Resource"] = request.Resource
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["ResourceType"] = request.ResourceType
	query["ScheduleType"] = request.ScheduleType
	query["StopTime"] = request.StopTime
	query["TaskName"] = request.TaskName
	query["TemplateCode"] = request.TemplateCode
	query["TemplateName"] = request.TemplateName
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateCallTask"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateCallTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateCallTask(request *CreateCallTaskRequest) (_result *CreateCallTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateCallTaskResponse{}
	_body, _err := client.CreateCallTaskWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	query["Caller"] = request.Caller
	query["CorpName"] = request.CorpName
	query["DialogId"] = request.DialogId
	query["IsSelfLine"] = request.IsSelfLine
	query["NumberStatusIdent"] = request.NumberStatusIdent
	query["OwnerId"] = request.OwnerId
	query["RecallInterval"] = request.RecallInterval
	query["RecallStateCodes"] = request.RecallStateCodes
	query["RecallTimes"] = request.RecallTimes
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["RetryType"] = request.RetryType
	query["TaskName"] = request.TaskName
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateRobotTask"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateRobotTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	query["OwnerId"] = request.OwnerId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["TaskId"] = request.TaskId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteRobotTask"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteRobotTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) ExecuteCallTaskWithOptions(request *ExecuteCallTaskRequest, runtime *util.RuntimeOptions) (_result *ExecuteCallTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["FireTime"] = request.FireTime
	query["OwnerId"] = request.OwnerId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["Status"] = request.Status
	query["TaskId"] = request.TaskId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("ExecuteCallTask"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ExecuteCallTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ExecuteCallTask(request *ExecuteCallTaskRequest) (_result *ExecuteCallTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ExecuteCallTaskResponse{}
	_body, _err := client.ExecuteCallTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetCallInfoWithOptions(request *GetCallInfoRequest, runtime *util.RuntimeOptions) (_result *GetCallInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["OwnerId"] = request.OwnerId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["RtcId"] = request.RtcId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("GetCallInfo"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCallInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetCallInfo(request *GetCallInfoRequest) (_result *GetCallInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetCallInfoResponse{}
	_body, _err := client.GetCallInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetHotlineQualificationByOrderWithOptions(request *GetHotlineQualificationByOrderRequest, runtime *util.RuntimeOptions) (_result *GetHotlineQualificationByOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["OrderId"] = request.OrderId
	query["OwnerId"] = request.OwnerId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("GetHotlineQualificationByOrder"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetHotlineQualificationByOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetHotlineQualificationByOrder(request *GetHotlineQualificationByOrderRequest) (_result *GetHotlineQualificationByOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetHotlineQualificationByOrderResponse{}
	_body, _err := client.GetHotlineQualificationByOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetMqttTokenWithOptions(request *GetMqttTokenRequest, runtime *util.RuntimeOptions) (_result *GetMqttTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["OwnerId"] = request.OwnerId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("GetMqttToken"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetMqttTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetMqttToken(request *GetMqttTokenRequest) (_result *GetMqttTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetMqttTokenResponse{}
	_body, _err := client.GetMqttTokenWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	query["DeviceId"] = request.DeviceId
	query["IsCustomAccount"] = request.IsCustomAccount
	query["OwnerId"] = request.OwnerId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["UserId"] = request.UserId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("GetRtcToken"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetRtcTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) GetTokenWithOptions(request *GetTokenRequest, runtime *util.RuntimeOptions) (_result *GetTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["OwnerId"] = request.OwnerId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["TokenType"] = request.TokenType
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("GetToken"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetToken(request *GetTokenRequest) (_result *GetTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTokenResponse{}
	_body, _err := client.GetTokenWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	query["ByeCode"] = request.ByeCode
	query["ByeTtsParams"] = request.ByeTtsParams
	query["CalledNumber"] = request.CalledNumber
	query["CalledShowNumber"] = request.CalledShowNumber
	query["MenuKeyMap"] = request.MenuKeyMap
	query["OutId"] = request.OutId
	query["OwnerId"] = request.OwnerId
	query["PlayTimes"] = request.PlayTimes
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["StartCode"] = request.StartCode
	query["StartTtsParams"] = request.StartTtsParams
	query["Timeout"] = request.Timeout
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("IvrCall"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &IvrCallResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) ListCallTaskWithOptions(request *ListCallTaskRequest, runtime *util.RuntimeOptions) (_result *ListCallTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["BizType"] = request.BizType
	query["OwnerId"] = request.OwnerId
	query["PageNumber"] = request.PageNumber
	query["PageSize"] = request.PageSize
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["Status"] = request.Status
	query["TaskId"] = request.TaskId
	query["TaskName"] = request.TaskName
	query["TemplateName"] = request.TemplateName
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("ListCallTask"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListCallTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListCallTask(request *ListCallTaskRequest) (_result *ListCallTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListCallTaskResponse{}
	_body, _err := client.ListCallTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListCallTaskDetailWithOptions(request *ListCallTaskDetailRequest, runtime *util.RuntimeOptions) (_result *ListCallTaskDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["CalledNum"] = request.CalledNum
	query["OwnerId"] = request.OwnerId
	query["PageNumber"] = request.PageNumber
	query["PageSize"] = request.PageSize
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["Status"] = request.Status
	query["TaskId"] = request.TaskId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("ListCallTaskDetail"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListCallTaskDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListCallTaskDetail(request *ListCallTaskDetailRequest) (_result *ListCallTaskDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListCallTaskDetailResponse{}
	_body, _err := client.ListCallTaskDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListHotlineTransferNumberWithOptions(request *ListHotlineTransferNumberRequest, runtime *util.RuntimeOptions) (_result *ListHotlineTransferNumberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["HotlineNumber"] = request.HotlineNumber
	query["OwnerId"] = request.OwnerId
	query["PageNo"] = request.PageNo
	query["PageSize"] = request.PageSize
	query["QualificationId"] = request.QualificationId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("ListHotlineTransferNumber"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListHotlineTransferNumberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListHotlineTransferNumber(request *ListHotlineTransferNumberRequest) (_result *ListHotlineTransferNumberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListHotlineTransferNumberResponse{}
	_body, _err := client.ListHotlineTransferNumberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListHotlineTransferRegisterFileWithOptions(request *ListHotlineTransferRegisterFileRequest, runtime *util.RuntimeOptions) (_result *ListHotlineTransferRegisterFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["HotlineNumber"] = request.HotlineNumber
	query["OwnerId"] = request.OwnerId
	query["PageNo"] = request.PageNo
	query["PageSize"] = request.PageSize
	query["QualificationId"] = request.QualificationId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("ListHotlineTransferRegisterFile"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListHotlineTransferRegisterFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListHotlineTransferRegisterFile(request *ListHotlineTransferRegisterFileRequest) (_result *ListHotlineTransferRegisterFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListHotlineTransferRegisterFileResponse{}
	_body, _err := client.ListHotlineTransferRegisterFileWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	query["CallId"] = request.CallId
	query["OwnerId"] = request.OwnerId
	query["ProdId"] = request.ProdId
	query["QueryDate"] = request.QueryDate
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryCallDetailByCallId"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryCallDetailByCallIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	query["Callee"] = request.Callee
	query["OwnerId"] = request.OwnerId
	query["QueryDate"] = request.QueryDate
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["TaskId"] = request.TaskId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryCallDetailByTaskId"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryCallDetailByTaskIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) QueryCallInPoolTransferConfigWithOptions(request *QueryCallInPoolTransferConfigRequest, runtime *util.RuntimeOptions) (_result *QueryCallInPoolTransferConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["OwnerId"] = request.OwnerId
	query["PhoneNumber"] = request.PhoneNumber
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryCallInPoolTransferConfig"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryCallInPoolTransferConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryCallInPoolTransferConfig(request *QueryCallInPoolTransferConfigRequest) (_result *QueryCallInPoolTransferConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryCallInPoolTransferConfigResponse{}
	_body, _err := client.QueryCallInPoolTransferConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryCallInTransferRecordWithOptions(request *QueryCallInTransferRecordRequest, runtime *util.RuntimeOptions) (_result *QueryCallInTransferRecordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["CallInCaller"] = request.CallInCaller
	query["OwnerId"] = request.OwnerId
	query["PageNo"] = request.PageNo
	query["PageSize"] = request.PageSize
	query["PhoneNumber"] = request.PhoneNumber
	query["QueryDate"] = request.QueryDate
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryCallInTransferRecord"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryCallInTransferRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryCallInTransferRecord(request *QueryCallInTransferRecordRequest) (_result *QueryCallInTransferRecordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryCallInTransferRecordResponse{}
	_body, _err := client.QueryCallInTransferRecordWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	query["AuditStatus"] = request.AuditStatus
	query["OwnerId"] = request.OwnerId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryRobotInfoList"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryRobotInfoListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	query["Callee"] = request.Callee
	query["OwnerId"] = request.OwnerId
	query["QueryDate"] = request.QueryDate
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["TaskId"] = request.TaskId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryRobotTaskCallDetail"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryRobotTaskCallDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	query["CallResult"] = request.CallResult
	query["Called"] = request.Called
	query["DialogCountFrom"] = request.DialogCountFrom
	query["DialogCountTo"] = request.DialogCountTo
	query["DurationFrom"] = request.DurationFrom
	query["DurationTo"] = request.DurationTo
	query["HangupDirection"] = request.HangupDirection
	query["OwnerId"] = request.OwnerId
	query["PageNo"] = request.PageNo
	query["PageSize"] = request.PageSize
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["TaskId"] = request.TaskId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryRobotTaskCallList"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryRobotTaskCallListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	query["Id"] = request.Id
	query["OwnerId"] = request.OwnerId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryRobotTaskDetail"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryRobotTaskDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	query["OwnerId"] = request.OwnerId
	query["PageNo"] = request.PageNo
	query["PageSize"] = request.PageSize
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["Status"] = request.Status
	query["TaskName"] = request.TaskName
	query["Time"] = request.Time
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryRobotTaskList"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryRobotTaskListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	query["OwnerId"] = request.OwnerId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryRobotv2AllList"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryRobotv2AllListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) QueryVirtualNumberWithOptions(request *QueryVirtualNumberRequest, runtime *util.RuntimeOptions) (_result *QueryVirtualNumberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["OwnerId"] = request.OwnerId
	query["PageNo"] = request.PageNo
	query["PageSize"] = request.PageSize
	query["ProdCode"] = request.ProdCode
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["RouteType"] = request.RouteType
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryVirtualNumber"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryVirtualNumberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryVirtualNumber(request *QueryVirtualNumberRequest) (_result *QueryVirtualNumberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryVirtualNumberResponse{}
	_body, _err := client.QueryVirtualNumberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryVirtualNumberRelationWithOptions(request *QueryVirtualNumberRelationRequest, runtime *util.RuntimeOptions) (_result *QueryVirtualNumberRelationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["OwnerId"] = request.OwnerId
	query["PageNo"] = request.PageNo
	query["PageSize"] = request.PageSize
	query["PhoneNum"] = request.PhoneNum
	query["ProdCode"] = request.ProdCode
	query["QualificationId"] = request.QualificationId
	query["RegionNameCity"] = request.RegionNameCity
	query["RelatedNum"] = request.RelatedNum
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["RouteType"] = request.RouteType
	query["SpecId"] = request.SpecId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryVirtualNumberRelation"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryVirtualNumberRelationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryVirtualNumberRelation(request *QueryVirtualNumberRelationRequest) (_result *QueryVirtualNumberRelationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryVirtualNumberRelationResponse{}
	_body, _err := client.QueryVirtualNumberRelationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RefreshMqttTokenWithOptions(request *RefreshMqttTokenRequest, runtime *util.RuntimeOptions) (_result *RefreshMqttTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ClientId"] = request.ClientId
	query["OwnerId"] = request.OwnerId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("RefreshMqttToken"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RefreshMqttTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RefreshMqttToken(request *RefreshMqttTokenRequest) (_result *RefreshMqttTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RefreshMqttTokenResponse{}
	_body, _err := client.RefreshMqttTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SendVerificationWithOptions(request *SendVerificationRequest, runtime *util.RuntimeOptions) (_result *SendVerificationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["BizType"] = request.BizType
	query["OwnerId"] = request.OwnerId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["Target"] = request.Target
	query["VerifyType"] = request.VerifyType
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("SendVerification"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &SendVerificationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SendVerification(request *SendVerificationRequest) (_result *SendVerificationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SendVerificationResponse{}
	_body, _err := client.SendVerificationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetTransferCalleePoolConfigWithOptions(request *SetTransferCalleePoolConfigRequest, runtime *util.RuntimeOptions) (_result *SetTransferCalleePoolConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["CalledRouteMode"] = request.CalledRouteMode
	query["Details"] = request.Details
	query["OwnerId"] = request.OwnerId
	query["PhoneNumber"] = request.PhoneNumber
	query["QualificationId"] = request.QualificationId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("SetTransferCalleePoolConfig"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &SetTransferCalleePoolConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetTransferCalleePoolConfig(request *SetTransferCalleePoolConfigRequest) (_result *SetTransferCalleePoolConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetTransferCalleePoolConfigResponse{}
	_body, _err := client.SetTransferCalleePoolConfigWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	query["CalledNumber"] = request.CalledNumber
	query["CalledShowNumber"] = request.CalledShowNumber
	query["OutId"] = request.OutId
	query["OwnerId"] = request.OwnerId
	query["PlayTimes"] = request.PlayTimes
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["Speed"] = request.Speed
	query["TtsCode"] = request.TtsCode
	query["TtsParam"] = request.TtsParam
	query["Volume"] = request.Volume
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("SingleCallByTts"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &SingleCallByTtsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	query["CalledNumber"] = request.CalledNumber
	query["CalledShowNumber"] = request.CalledShowNumber
	query["OutId"] = request.OutId
	query["OwnerId"] = request.OwnerId
	query["PlayTimes"] = request.PlayTimes
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["Speed"] = request.Speed
	query["VoiceCode"] = request.VoiceCode
	query["Volume"] = request.Volume
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("SingleCallByVoice"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &SingleCallByVoiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	query["ActionCodeBreak"] = request.ActionCodeBreak
	query["ActionCodeTimeBreak"] = request.ActionCodeTimeBreak
	query["AsrBaseId"] = request.AsrBaseId
	query["AsrModelId"] = request.AsrModelId
	query["BackgroundFileCode"] = request.BackgroundFileCode
	query["BackgroundSpeed"] = request.BackgroundSpeed
	query["BackgroundVolume"] = request.BackgroundVolume
	query["CalledNumber"] = request.CalledNumber
	query["CalledShowNumber"] = request.CalledShowNumber
	query["DynamicId"] = request.DynamicId
	query["EarlyMediaAsr"] = request.EarlyMediaAsr
	query["EnableITN"] = request.EnableITN
	query["MuteTime"] = request.MuteTime
	query["OutId"] = request.OutId
	query["OwnerId"] = request.OwnerId
	query["PauseTime"] = request.PauseTime
	query["RecordFlag"] = request.RecordFlag
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["SessionTimeout"] = request.SessionTimeout
	query["Speed"] = request.Speed
	query["StreamAsr"] = request.StreamAsr
	query["TtsConf"] = request.TtsConf
	query["TtsSpeed"] = request.TtsSpeed
	query["TtsStyle"] = request.TtsStyle
	query["TtsVolume"] = request.TtsVolume
	query["VoiceCode"] = request.VoiceCode
	query["VoiceCodeParam"] = request.VoiceCodeParam
	query["Volume"] = request.Volume
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("SmartCall"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &SmartCallResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	query["CallId"] = request.CallId
	query["Command"] = request.Command
	query["OwnerId"] = request.OwnerId
	query["Param"] = request.Param
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("SmartCallOperate"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &SmartCallOperateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) StartRobotTaskWithOptions(request *StartRobotTaskRequest, runtime *util.RuntimeOptions) (_result *StartRobotTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["OwnerId"] = request.OwnerId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["ScheduleTime"] = request.ScheduleTime
	query["TaskId"] = request.TaskId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("StartRobotTask"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &StartRobotTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	query := map[string]interface{}{}
	query["OwnerId"] = request.OwnerId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["TaskId"] = request.TaskId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("StopRobotTask"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &StopRobotTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) SubmitHotlineTransferRegisterWithOptions(request *SubmitHotlineTransferRegisterRequest, runtime *util.RuntimeOptions) (_result *SubmitHotlineTransferRegisterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["Agreement"] = request.Agreement
	query["HotlineNumber"] = request.HotlineNumber
	query["OperatorIdentityCard"] = request.OperatorIdentityCard
	query["OperatorMail"] = request.OperatorMail
	query["OperatorMailVerifyCode"] = request.OperatorMailVerifyCode
	query["OperatorMobile"] = request.OperatorMobile
	query["OperatorMobileVerifyCode"] = request.OperatorMobileVerifyCode
	query["OperatorName"] = request.OperatorName
	query["OwnerId"] = request.OwnerId
	query["QualificationId"] = request.QualificationId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["TransferPhoneNumberInfos"] = request.TransferPhoneNumberInfos
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("SubmitHotlineTransferRegister"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitHotlineTransferRegisterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitHotlineTransferRegister(request *SubmitHotlineTransferRegisterRequest) (_result *SubmitHotlineTransferRegisterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitHotlineTransferRegisterResponse{}
	_body, _err := client.SubmitHotlineTransferRegisterWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	query["CalledNumber"] = request.CalledNumber
	query["Id"] = request.Id
	query["OwnerId"] = request.OwnerId
	query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	query["ResourceOwnerId"] = request.ResourceOwnerId
	query["TtsParam"] = request.TtsParam
	query["TtsParamHead"] = request.TtsParamHead
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("UploadRobotTaskCalledFile"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UploadRobotTaskCalledFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
