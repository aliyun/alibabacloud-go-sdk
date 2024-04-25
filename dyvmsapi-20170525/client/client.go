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

type AddVirtualNumberRelationRequest struct {
	// The company names. Separate multiple company names with commas (,).
	CorpNameList *string `json:"CorpNameList,omitempty" xml:"CorpNameList,omitempty"`
	// The real numbers. Separate multiple real numbers with commas (,).
	NumberList *string `json:"NumberList,omitempty" xml:"NumberList,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The virtual number.
	PhoneNum *string `json:"PhoneNum,omitempty" xml:"PhoneNum,omitempty"`
	// The service name. Default value: **dyvms**.
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The route type. Valid values:
	//
	// *   **0**: number location first.
	// *   **1**: random.
	RouteType *int32 `json:"RouteType,omitempty" xml:"RouteType,omitempty"`
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
	// The response code.
	//
	// *   The value 200 indicates that the request was successful.
	// *   For more information about other response codes, see [API error codes](~~112502~~).
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The numbers that failed to be associated.
	//
	// > If all numbers are associated, no value is returned for this parameter.
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
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
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddVirtualNumberRelationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *AddVirtualNumberRelationResponse) SetStatusCode(v int32) *AddVirtualNumberRelationResponse {
	s.StatusCode = &v
	return s
}

func (s *AddVirtualNumberRelationResponse) SetBody(v *AddVirtualNumberRelationResponseBody) *AddVirtualNumberRelationResponse {
	s.Body = v
	return s
}

type BatchRobotSmartCallRequest struct {
	// The called number. Only mobile phone numbers in the Chinese mainland are supported.
	//
	// You can set up to 1,000 called numbers and separate the numbers with commas (,).
	CalledNumber *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	// The number displayed to called parties, which must be a number you purchased. You can view the numbers you purchased in the [Voice Messaging Service console](https://dyvms.console.aliyun.com/dyvms.htm#/number/normal).
	//
	// You can set up to 100 numbers and separate the numbers with commas (,).
	CalledShowNumber *string `json:"CalledShowNumber,omitempty" xml:"CalledShowNumber,omitempty"`
	// The company name, which must be the same as the **company name** specified on the [qualification management page](https://dyvms.console.aliyun.com/dyvms.htm#/corp/normal).
	//
	// > This parameter is optional if **isSelfLine** is set to **true**.
	CorpName *string `json:"CorpName,omitempty" xml:"CorpName,omitempty"`
	// The ID of the robot or communication script that is used to initiate a call.
	//
	// You can obtain the **communication script ID** from the [Communication script management](https://dyvms.console.aliyun.com/dyvms.htm#/smart-call/saas/robot/list) page.
	DialogId *string `json:"DialogId,omitempty" xml:"DialogId,omitempty"`
	// The speech recognition identifier of early media. The default value is **false**, which means that the speech recognition identifier of early media is not enabled.
	//
	// Set the parameter to **true** if you want to enable the speech recognition identifier of early media.
	EarlyMediaAsr *bool `json:"EarlyMediaAsr,omitempty" xml:"EarlyMediaAsr,omitempty"`
	// Specifies whether to call the self-managed line. Default value: **false**.
	IsSelfLine           *bool   `json:"IsSelfLine,omitempty" xml:"IsSelfLine,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Specifies whether the call is scheduled. If you set this parameter to **true**, the **ScheduleTime** parameter is required.
	ScheduleCall *bool `json:"ScheduleCall,omitempty" xml:"ScheduleCall,omitempty"`
	// The preset call time. This value is a UNIX timestamp. Unit: milliseconds.
	//
	// >  This parameter is required only when **ScheduleCall** is set to **true**.
	ScheduleTime *int64 `json:"ScheduleTime,omitempty" xml:"ScheduleTime,omitempty"`
	// The task name. The task name can be up to 30 characters in length.
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// The variable value of the TTS template, in the JSON format.
	//
	// The variable value must correspond to a number. The TtsParam parameter must be used together with the TtsParamHead parameter.
	TtsParam *string `json:"TtsParam,omitempty" xml:"TtsParam,omitempty"`
	// The call tasks with variables, in the JSON format.
	//
	// The parameter value is a list of variable names. The TtsParamHead parameter must be used together with the TtsParam parameter.
	TtsParamHead *string `json:"TtsParamHead,omitempty" xml:"TtsParamHead,omitempty"`
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
	// The response code.
	//
	// *   The value OK indicates that the request was successful.
	// *   For more information about other response codes, see [API error codes](~~112502~~).
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The unique ID of the robocall task. You can call the [QueryCallDetailByTaskId](~~393537~~) operation to query the details of the task based on the task ID.
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchRobotSmartCallResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *BatchRobotSmartCallResponse) SetStatusCode(v int32) *BatchRobotSmartCallResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchRobotSmartCallResponse) SetBody(v *BatchRobotSmartCallResponseBody) *BatchRobotSmartCallResponse {
	s.Body = v
	return s
}

type CancelOrderRobotTaskRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The unique ID of the robocall task. You can call the [CreateRobotTask](~~393531~~) operation to obtain the ID of the robocall task.
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
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
	// The response code.
	//
	// *   The value OK indicates that the request was successful.
	// *   For more information about other response codes, see [API error codes](~~112502~~).
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// *   **true**: The request was successful.
	// *   **false**: The request failed.
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelOrderRobotTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *CancelOrderRobotTaskResponse) SetStatusCode(v int32) *CancelOrderRobotTaskResponse {
	s.StatusCode = &v
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
	// The unique ID of the robocall task. You can call the [CreateRobotTask](~~393531~~) operation to obtain the task ID.
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
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
	// The response code.
	//
	// *   The value OK indicates that the request was successful.
	// *   For more information about other response codes, see [API error codes](~~112502~~).
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// *   **true**: The request was successful.
	// *   **false**: The request failed.
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelRobotTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *CancelRobotTaskResponse) SetStatusCode(v int32) *CancelRobotTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelRobotTaskResponse) SetBody(v *CancelRobotTaskResponseBody) *CancelRobotTaskResponse {
	s.Body = v
	return s
}

type ChangeMediaTypeRequest struct {
	CallId               *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	CalledNum            *string `json:"CalledNum,omitempty" xml:"CalledNum,omitempty"`
	MediaType            *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	OutId                *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ChangeMediaTypeRequest) String() string {
	return tea.Prettify(s)
}

func (s ChangeMediaTypeRequest) GoString() string {
	return s.String()
}

func (s *ChangeMediaTypeRequest) SetCallId(v string) *ChangeMediaTypeRequest {
	s.CallId = &v
	return s
}

func (s *ChangeMediaTypeRequest) SetCalledNum(v string) *ChangeMediaTypeRequest {
	s.CalledNum = &v
	return s
}

func (s *ChangeMediaTypeRequest) SetMediaType(v string) *ChangeMediaTypeRequest {
	s.MediaType = &v
	return s
}

func (s *ChangeMediaTypeRequest) SetOutId(v string) *ChangeMediaTypeRequest {
	s.OutId = &v
	return s
}

func (s *ChangeMediaTypeRequest) SetOwnerId(v int64) *ChangeMediaTypeRequest {
	s.OwnerId = &v
	return s
}

func (s *ChangeMediaTypeRequest) SetResourceOwnerAccount(v string) *ChangeMediaTypeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ChangeMediaTypeRequest) SetResourceOwnerId(v int64) *ChangeMediaTypeRequest {
	s.ResourceOwnerId = &v
	return s
}

type ChangeMediaTypeResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Code               *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message            *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Model              *bool   `json:"Model,omitempty" xml:"Model,omitempty"`
	Success            *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ChangeMediaTypeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ChangeMediaTypeResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeMediaTypeResponseBody) SetAccessDeniedDetail(v string) *ChangeMediaTypeResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ChangeMediaTypeResponseBody) SetCode(v string) *ChangeMediaTypeResponseBody {
	s.Code = &v
	return s
}

func (s *ChangeMediaTypeResponseBody) SetMessage(v string) *ChangeMediaTypeResponseBody {
	s.Message = &v
	return s
}

func (s *ChangeMediaTypeResponseBody) SetModel(v bool) *ChangeMediaTypeResponseBody {
	s.Model = &v
	return s
}

func (s *ChangeMediaTypeResponseBody) SetSuccess(v bool) *ChangeMediaTypeResponseBody {
	s.Success = &v
	return s
}

type ChangeMediaTypeResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChangeMediaTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChangeMediaTypeResponse) String() string {
	return tea.Prettify(s)
}

func (s ChangeMediaTypeResponse) GoString() string {
	return s.String()
}

func (s *ChangeMediaTypeResponse) SetHeaders(v map[string]*string) *ChangeMediaTypeResponse {
	s.Headers = v
	return s
}

func (s *ChangeMediaTypeResponse) SetStatusCode(v int32) *ChangeMediaTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *ChangeMediaTypeResponse) SetBody(v *ChangeMediaTypeResponseBody) *ChangeMediaTypeResponse {
	s.Body = v
	return s
}

type CreateCallTaskRequest struct {
	// The type of the task template. Valid values:
	//
	// *   **VMS_VOICE_TTS**: the text-to-speech (TTS) notification template.
	// *   **VMS_VOICE_CODE**: the voice notification template.
	// *   **VMS_TTS**: the voice verification code template.
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// The called numbers.
	//
	// *   If you set DataType to LIST, the value of Data is in the LIST format.
	// *   If you set DataType to JSON, the value of Data is in the JSON format.
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The type of called numbers. Valid values:
	//
	// *   **LIST**: the called numbers that are separated by commas (,).
	// *   **JSON**: a JSON-formatted list of called numbers with template parameters.
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// This parameter is unavailable.
	FireTime *string `json:"FireTime,omitempty" xml:"FireTime,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The calling number. Only virtual numbers are supported.
	Resource             *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The type of the calling number. Set the value to **LIST**.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// This parameter is unavailable.
	ScheduleType *string `json:"ScheduleType,omitempty" xml:"ScheduleType,omitempty"`
	// This parameter is unavailable.
	StopTime *string `json:"StopTime,omitempty" xml:"StopTime,omitempty"`
	// The task name.
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// The template ID.
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	// The template name.
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
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
	// The response code.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The task ID.
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
	// The request ID.
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCallTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *CreateCallTaskResponse) SetStatusCode(v int32) *CreateCallTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCallTaskResponse) SetBody(v *CreateCallTaskResponseBody) *CreateCallTaskResponse {
	s.Body = v
	return s
}

type CreateRobotTaskRequest struct {
	// The calling number.
	//
	// You must use the phone numbers that you have purchased and separate multiple numbers with commas (,). You can log on to the [Voice Messaging Service console](https://dyvms.console.aliyun.com/overview/home) and choose **Real Number Service** > **Real Number Management** to view the numbers you purchased.
	Caller *string `json:"Caller,omitempty" xml:"Caller,omitempty"`
	// The company name, which must be the same as the **enterprise name** on the qualification management page.
	CorpName *string `json:"CorpName,omitempty" xml:"CorpName,omitempty"`
	// The ID of the robot or communication script that is used to initiate the call.
	//
	// You can log on to the [Voice Messaging Service console](https://dyvms.console.aliyun.com/overview/home) and choose **Intelligent Voice Robot** > **Communication Script Management** to view the communication script ID.
	DialogId *int64 `json:"DialogId,omitempty" xml:"DialogId,omitempty"`
	// Specifies whether to call the self-managed line. Valid values:
	//
	// *   **false** (default)
	// *   **true**
	//
	// > If you set this parameter to **true**, calling numbers are not verified.
	IsSelfLine *bool `json:"IsSelfLine,omitempty" xml:"IsSelfLine,omitempty"`
	// Specifies whether to enable number status identification. Valid values:
	//
	// *   **false** (default)
	// *   **true**
	//
	// > If you set this parameter to **true**, the reason why a call is not answered is recorded.
	NumberStatusIdent *bool  `json:"NumberStatusIdent,omitempty" xml:"NumberStatusIdent,omitempty"`
	OwnerId           *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The redial interval. Unit: minutes. The value must be greater than 1.
	//
	// > The maximum redial interval is 30 minutes.
	RecallInterval *int32 `json:"RecallInterval,omitempty" xml:"RecallInterval,omitempty"`
	// The call state in which redial is required. Separate multiple call states with commas (,). Valid values:
	//
	// *   **200010**: The phone of the called party is powered off.
	// *   **200011**: The number of the called party is out of service.
	// *   **200002**: The line is busy.
	// *   **200012**: The call is lost.
	// *   **200005**: The called party cannot be connected.
	// *   **200003**: The called party does not respond to the call.
	RecallStateCodes *string `json:"RecallStateCodes,omitempty" xml:"RecallStateCodes,omitempty"`
	// The number of redial times.
	RecallTimes          *int32  `json:"RecallTimes,omitempty" xml:"RecallTimes,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Specifies whether to enable auto-redial. Valid values:
	//
	// *   **1**: enables auto-redial.
	// *   **0**: disables auto-redial.
	RetryType *int32 `json:"RetryType,omitempty" xml:"RetryType,omitempty"`
	// The task name. The task name can be up to 30 characters in length.
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
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
	// The response code.
	//
	// *   The value OK indicates that the request was successful.
	// *   For more information about other response codes, see [API error codes](~~112502~~).
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The unique ID of the robocall task.
	//
	// You can call the [QueryRobotTaskDetail](~~393538~~) operation to query the details of the task based on the task ID.
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRobotTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *CreateRobotTaskResponse) SetStatusCode(v int32) *CreateRobotTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRobotTaskResponse) SetBody(v *CreateRobotTaskResponseBody) *CreateRobotTaskResponse {
	s.Body = v
	return s
}

type DegradeVideoFileRequest struct {
	CallId               *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	CalledNumber         *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	MediaType            *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	OutId                *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DegradeVideoFileRequest) String() string {
	return tea.Prettify(s)
}

func (s DegradeVideoFileRequest) GoString() string {
	return s.String()
}

func (s *DegradeVideoFileRequest) SetCallId(v string) *DegradeVideoFileRequest {
	s.CallId = &v
	return s
}

func (s *DegradeVideoFileRequest) SetCalledNumber(v string) *DegradeVideoFileRequest {
	s.CalledNumber = &v
	return s
}

func (s *DegradeVideoFileRequest) SetMediaType(v string) *DegradeVideoFileRequest {
	s.MediaType = &v
	return s
}

func (s *DegradeVideoFileRequest) SetOutId(v string) *DegradeVideoFileRequest {
	s.OutId = &v
	return s
}

func (s *DegradeVideoFileRequest) SetOwnerId(v int64) *DegradeVideoFileRequest {
	s.OwnerId = &v
	return s
}

func (s *DegradeVideoFileRequest) SetResourceOwnerAccount(v string) *DegradeVideoFileRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DegradeVideoFileRequest) SetResourceOwnerId(v int64) *DegradeVideoFileRequest {
	s.ResourceOwnerId = &v
	return s
}

type DegradeVideoFileResponseBody struct {
	AccessDeniedDetail *string                `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Code               *string                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data               map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	Message            *string                `json:"Message,omitempty" xml:"Message,omitempty"`
	Success            *bool                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DegradeVideoFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DegradeVideoFileResponseBody) GoString() string {
	return s.String()
}

func (s *DegradeVideoFileResponseBody) SetAccessDeniedDetail(v string) *DegradeVideoFileResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DegradeVideoFileResponseBody) SetCode(v string) *DegradeVideoFileResponseBody {
	s.Code = &v
	return s
}

func (s *DegradeVideoFileResponseBody) SetData(v map[string]interface{}) *DegradeVideoFileResponseBody {
	s.Data = v
	return s
}

func (s *DegradeVideoFileResponseBody) SetMessage(v string) *DegradeVideoFileResponseBody {
	s.Message = &v
	return s
}

func (s *DegradeVideoFileResponseBody) SetSuccess(v bool) *DegradeVideoFileResponseBody {
	s.Success = &v
	return s
}

type DegradeVideoFileResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DegradeVideoFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DegradeVideoFileResponse) String() string {
	return tea.Prettify(s)
}

func (s DegradeVideoFileResponse) GoString() string {
	return s.String()
}

func (s *DegradeVideoFileResponse) SetHeaders(v map[string]*string) *DegradeVideoFileResponse {
	s.Headers = v
	return s
}

func (s *DegradeVideoFileResponse) SetStatusCode(v int32) *DegradeVideoFileResponse {
	s.StatusCode = &v
	return s
}

func (s *DegradeVideoFileResponse) SetBody(v *DegradeVideoFileResponseBody) *DegradeVideoFileResponse {
	s.Body = v
	return s
}

type DeleteRobotTaskRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The unique ID of the robocall task. You can call the [CreateRobotTask](~~CreateRobotTask~~) operation to obtain the task ID.
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
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
	// The response code.
	//
	// *   The value OK indicates that the request was successful.
	// *   For more information about other response codes, see [API error codes](~~112502~~).
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// *   **true**: The request was successful.
	// *   **false**: The request failed.
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRobotTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DeleteRobotTaskResponse) SetStatusCode(v int32) *DeleteRobotTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRobotTaskResponse) SetBody(v *DeleteRobotTaskResponseBody) *DeleteRobotTaskResponse {
	s.Body = v
	return s
}

type ExecuteCallTaskRequest struct {
	// The time when the call task is executed, in the yyyy-MM-dd HH:mm:ss format.
	//
	// > You can leave this parameter empty.
	FireTime             *string `json:"FireTime,omitempty" xml:"FireTime,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The task state. Valid values:
	//
	// *   **RUNNING**
	// *   **STOP**
	// *   **CANCEL**
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The task ID. You can call the [CreateCallTask](~~CreateCallTask~~) operation to obtain the task ID.
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
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
	// The response code.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// *   **true**: The request was successful.
	// *   **false**: The request failed.
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The request ID.
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ExecuteCallTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ExecuteCallTaskResponse) SetStatusCode(v int32) *ExecuteCallTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *ExecuteCallTaskResponse) SetBody(v *ExecuteCallTaskResponseBody) *ExecuteCallTaskResponse {
	s.Body = v
	return s
}

type GetCallMediaTypeRequest struct {
	CallId               *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	CalledNumber         *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s GetCallMediaTypeRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCallMediaTypeRequest) GoString() string {
	return s.String()
}

func (s *GetCallMediaTypeRequest) SetCallId(v string) *GetCallMediaTypeRequest {
	s.CallId = &v
	return s
}

func (s *GetCallMediaTypeRequest) SetCalledNumber(v string) *GetCallMediaTypeRequest {
	s.CalledNumber = &v
	return s
}

func (s *GetCallMediaTypeRequest) SetOwnerId(v int64) *GetCallMediaTypeRequest {
	s.OwnerId = &v
	return s
}

func (s *GetCallMediaTypeRequest) SetResourceOwnerAccount(v string) *GetCallMediaTypeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetCallMediaTypeRequest) SetResourceOwnerId(v int64) *GetCallMediaTypeRequest {
	s.ResourceOwnerId = &v
	return s
}

type GetCallMediaTypeResponseBody struct {
	AccessDeniedDetail *string                `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Code               *string                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data               map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	Message            *string                `json:"Message,omitempty" xml:"Message,omitempty"`
	Success            *bool                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetCallMediaTypeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCallMediaTypeResponseBody) GoString() string {
	return s.String()
}

func (s *GetCallMediaTypeResponseBody) SetAccessDeniedDetail(v string) *GetCallMediaTypeResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetCallMediaTypeResponseBody) SetCode(v string) *GetCallMediaTypeResponseBody {
	s.Code = &v
	return s
}

func (s *GetCallMediaTypeResponseBody) SetData(v map[string]interface{}) *GetCallMediaTypeResponseBody {
	s.Data = v
	return s
}

func (s *GetCallMediaTypeResponseBody) SetMessage(v string) *GetCallMediaTypeResponseBody {
	s.Message = &v
	return s
}

func (s *GetCallMediaTypeResponseBody) SetSuccess(v bool) *GetCallMediaTypeResponseBody {
	s.Success = &v
	return s
}

type GetCallMediaTypeResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCallMediaTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCallMediaTypeResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCallMediaTypeResponse) GoString() string {
	return s.String()
}

func (s *GetCallMediaTypeResponse) SetHeaders(v map[string]*string) *GetCallMediaTypeResponse {
	s.Headers = v
	return s
}

func (s *GetCallMediaTypeResponse) SetStatusCode(v int32) *GetCallMediaTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCallMediaTypeResponse) SetBody(v *GetCallMediaTypeResponseBody) *GetCallMediaTypeResponse {
	s.Body = v
	return s
}

type GetCallProgressRequest struct {
	CallId               *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	CalledNum            *string `json:"CalledNum,omitempty" xml:"CalledNum,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s GetCallProgressRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCallProgressRequest) GoString() string {
	return s.String()
}

func (s *GetCallProgressRequest) SetCallId(v string) *GetCallProgressRequest {
	s.CallId = &v
	return s
}

func (s *GetCallProgressRequest) SetCalledNum(v string) *GetCallProgressRequest {
	s.CalledNum = &v
	return s
}

func (s *GetCallProgressRequest) SetOwnerId(v int64) *GetCallProgressRequest {
	s.OwnerId = &v
	return s
}

func (s *GetCallProgressRequest) SetResourceOwnerAccount(v string) *GetCallProgressRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetCallProgressRequest) SetResourceOwnerId(v int64) *GetCallProgressRequest {
	s.ResourceOwnerId = &v
	return s
}

type GetCallProgressResponseBody struct {
	AccessDeniedDetail *string                `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Code               *string                `json:"Code,omitempty" xml:"Code,omitempty"`
	Message            *string                `json:"Message,omitempty" xml:"Message,omitempty"`
	Model              map[string]interface{} `json:"Model,omitempty" xml:"Model,omitempty"`
	Success            *bool                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetCallProgressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCallProgressResponseBody) GoString() string {
	return s.String()
}

func (s *GetCallProgressResponseBody) SetAccessDeniedDetail(v string) *GetCallProgressResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetCallProgressResponseBody) SetCode(v string) *GetCallProgressResponseBody {
	s.Code = &v
	return s
}

func (s *GetCallProgressResponseBody) SetMessage(v string) *GetCallProgressResponseBody {
	s.Message = &v
	return s
}

func (s *GetCallProgressResponseBody) SetModel(v map[string]interface{}) *GetCallProgressResponseBody {
	s.Model = v
	return s
}

func (s *GetCallProgressResponseBody) SetSuccess(v bool) *GetCallProgressResponseBody {
	s.Success = &v
	return s
}

type GetCallProgressResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCallProgressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCallProgressResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCallProgressResponse) GoString() string {
	return s.String()
}

func (s *GetCallProgressResponse) SetHeaders(v map[string]*string) *GetCallProgressResponse {
	s.Headers = v
	return s
}

func (s *GetCallProgressResponse) SetStatusCode(v int32) *GetCallProgressResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCallProgressResponse) SetBody(v *GetCallProgressResponseBody) *GetCallProgressResponse {
	s.Body = v
	return s
}

type GetHotlineQualificationByOrderRequest struct {
	// The ticket ID.
	//
	// You can log on to the [Voice Messaging Service console](https://dyvms.console.aliyun.com/overview/home), choose **Qualification\&Communication Script Management** > **Qualification Management**, and then click the **400 Qualifications** tab to view the ticket ID.
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
	// The response code.
	//
	// *   The value OK indicates that the request was successful.
	// *   For more information about other response codes, see [API error codes](~~112502~~).
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response parameters.
	Data *GetHotlineQualificationByOrderResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// The ID of the qualification application ticket.
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The qualification ID.
	QualificationId *string `json:"QualificationId,omitempty" xml:"QualificationId,omitempty"`
	// The qualification state. Valid values:
	//
	// *   **NORMAL**
	// *   **OTHER**
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
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
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetHotlineQualificationByOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *GetHotlineQualificationByOrderResponse) SetStatusCode(v int32) *GetHotlineQualificationByOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHotlineQualificationByOrderResponse) SetBody(v *GetHotlineQualificationByOrderResponseBody) *GetHotlineQualificationByOrderResponse {
	s.Body = v
	return s
}

type GetTemporaryFileUrlRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	VideoId              *string `json:"VideoId,omitempty" xml:"VideoId,omitempty"`
}

func (s GetTemporaryFileUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTemporaryFileUrlRequest) GoString() string {
	return s.String()
}

func (s *GetTemporaryFileUrlRequest) SetOwnerId(v int64) *GetTemporaryFileUrlRequest {
	s.OwnerId = &v
	return s
}

func (s *GetTemporaryFileUrlRequest) SetResourceOwnerAccount(v string) *GetTemporaryFileUrlRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetTemporaryFileUrlRequest) SetResourceOwnerId(v int64) *GetTemporaryFileUrlRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetTemporaryFileUrlRequest) SetVideoId(v string) *GetTemporaryFileUrlRequest {
	s.VideoId = &v
	return s
}

type GetTemporaryFileUrlResponseBody struct {
	AccessDeniedDetail *string                `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Code               *string                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data               map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	Message            *string                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId          *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success            *bool                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetTemporaryFileUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTemporaryFileUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetTemporaryFileUrlResponseBody) SetAccessDeniedDetail(v string) *GetTemporaryFileUrlResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetTemporaryFileUrlResponseBody) SetCode(v string) *GetTemporaryFileUrlResponseBody {
	s.Code = &v
	return s
}

func (s *GetTemporaryFileUrlResponseBody) SetData(v map[string]interface{}) *GetTemporaryFileUrlResponseBody {
	s.Data = v
	return s
}

func (s *GetTemporaryFileUrlResponseBody) SetMessage(v string) *GetTemporaryFileUrlResponseBody {
	s.Message = &v
	return s
}

func (s *GetTemporaryFileUrlResponseBody) SetRequestId(v string) *GetTemporaryFileUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTemporaryFileUrlResponseBody) SetSuccess(v bool) *GetTemporaryFileUrlResponseBody {
	s.Success = &v
	return s
}

type GetTemporaryFileUrlResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTemporaryFileUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTemporaryFileUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTemporaryFileUrlResponse) GoString() string {
	return s.String()
}

func (s *GetTemporaryFileUrlResponse) SetHeaders(v map[string]*string) *GetTemporaryFileUrlResponse {
	s.Headers = v
	return s
}

func (s *GetTemporaryFileUrlResponse) SetStatusCode(v int32) *GetTemporaryFileUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTemporaryFileUrlResponse) SetBody(v *GetTemporaryFileUrlResponseBody) *GetTemporaryFileUrlResponse {
	s.Body = v
	return s
}

type GetTokenRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The token type.
	TokenType *string `json:"TokenType,omitempty" xml:"TokenType,omitempty"`
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
	// The response code.
	//
	// *   The value OK indicates that the request was successful.
	// *   For more information about other response codes, see [API error codes](~~112502~~).
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The token.
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
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
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *GetTokenResponse) SetStatusCode(v int32) *GetTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTokenResponse) SetBody(v *GetTokenResponseBody) *GetTokenResponse {
	s.Body = v
	return s
}

type GetVideoFieldUrlRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	VideoFile            *string `json:"VideoFile,omitempty" xml:"VideoFile,omitempty"`
}

func (s GetVideoFieldUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s GetVideoFieldUrlRequest) GoString() string {
	return s.String()
}

func (s *GetVideoFieldUrlRequest) SetOwnerId(v int64) *GetVideoFieldUrlRequest {
	s.OwnerId = &v
	return s
}

func (s *GetVideoFieldUrlRequest) SetResourceOwnerAccount(v string) *GetVideoFieldUrlRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetVideoFieldUrlRequest) SetResourceOwnerId(v int64) *GetVideoFieldUrlRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetVideoFieldUrlRequest) SetVideoFile(v string) *GetVideoFieldUrlRequest {
	s.VideoFile = &v
	return s
}

type GetVideoFieldUrlResponseBody struct {
	AccessDeniedDetail *string                `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Code               *string                `json:"Code,omitempty" xml:"Code,omitempty"`
	Message            *string                `json:"Message,omitempty" xml:"Message,omitempty"`
	Model              map[string]interface{} `json:"Model,omitempty" xml:"Model,omitempty"`
	Success            *bool                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetVideoFieldUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetVideoFieldUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetVideoFieldUrlResponseBody) SetAccessDeniedDetail(v string) *GetVideoFieldUrlResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetVideoFieldUrlResponseBody) SetCode(v string) *GetVideoFieldUrlResponseBody {
	s.Code = &v
	return s
}

func (s *GetVideoFieldUrlResponseBody) SetMessage(v string) *GetVideoFieldUrlResponseBody {
	s.Message = &v
	return s
}

func (s *GetVideoFieldUrlResponseBody) SetModel(v map[string]interface{}) *GetVideoFieldUrlResponseBody {
	s.Model = v
	return s
}

func (s *GetVideoFieldUrlResponseBody) SetSuccess(v bool) *GetVideoFieldUrlResponseBody {
	s.Success = &v
	return s
}

type GetVideoFieldUrlResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetVideoFieldUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetVideoFieldUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s GetVideoFieldUrlResponse) GoString() string {
	return s.String()
}

func (s *GetVideoFieldUrlResponse) SetHeaders(v map[string]*string) *GetVideoFieldUrlResponse {
	s.Headers = v
	return s
}

func (s *GetVideoFieldUrlResponse) SetStatusCode(v int32) *GetVideoFieldUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVideoFieldUrlResponse) SetBody(v *GetVideoFieldUrlResponseBody) *GetVideoFieldUrlResponse {
	s.Body = v
	return s
}

type IvrCallRequest struct {
	// The end voice.
	//
	// *   If you use a voice notification file, this parameter specifies the voice ID. You can log on to the [Voice Messaging Service console](https://dyvms.console.aliyun.com/overview/home), choose **Voice Messages** > **Voice Notifications**, and then click the **Voice Notification Files** tab to view the voice ID.
	// *   If you use a TTS template, this parameter specifies the template ID. You can log on to the [Voice Messaging Service console](https://dyvms.console.aliyun.com/overview/home), choose **Voice Messages** > **Voice Notifications**, and then click the **TTS Template** tab to view the template ID.
	//
	// > The value of the ByeCode parameter must be of the same type as the value of the StartCode parameter. This means that both parameters must specify voice IDs or TTS template IDs.
	ByeCode *string `json:"ByeCode,omitempty" xml:"ByeCode,omitempty"`
	// The variables in the TTS template, in the JSON format.
	//
	// > This parameter is required when the ByeCode parameter is set to the ID of a TTS template that contains variables.
	ByeTtsParams *string `json:"ByeTtsParams,omitempty" xml:"ByeTtsParams,omitempty"`
	// The called number.
	//
	// Only phone numbers in the Chinese mainland are supported. Each request supports only one called number.
	CalledNumber *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	// The calling number.
	//
	// The value must be a number you purchased. Each request supports only one calling number. In most cases, a calling number is configured with the maximum number of concurrent requests. New requests fail if the maximum number of concurrent requests is reached. You can log on to the [Voice Messaging Service console](https://dyvms.console.aliyun.com/overview/home) and choose **Real Number Service > Real Number Management** to view the number you purchased.
	CalledShowNumber *string `json:"CalledShowNumber,omitempty" xml:"CalledShowNumber,omitempty"`
	// The information about the key pressed by the subscriber.
	MenuKeyMap []*IvrCallRequestMenuKeyMap `json:"MenuKeyMap,omitempty" xml:"MenuKeyMap,omitempty" type:"Repeated"`
	// The ID that is reserved for the caller of the operation. This ID is returned to the caller in a receipt message.
	//
	// The value is of the STRING type and must be 1 to 15 bytes in length.
	OutId   *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of replay times. Valid values: 1 to 3.
	PlayTimes            *int64  `json:"PlayTimes,omitempty" xml:"PlayTimes,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The voice that is played when the call begins.
	//
	// *   If you use a voice notification file, this parameter specifies the voice ID. You can log on to the [Voice Messaging Service console](https://dyvms.console.aliyun.com/overview/home), choose **Voice Messages** > Voice Notifications, and then click the **Voice Notification Files** tab to view the voice ID.
	// *   If you use a text-to-speech (TTS) template, this parameter specifies the template ID. You can log on to the [Voice Messaging Service console](https://dyvms.console.aliyun.com/overview/home), choose **Voice Messages** > **Voice Notifications**, and then click the **TTS Template** tab to view the template ID.
	StartCode *string `json:"StartCode,omitempty" xml:"StartCode,omitempty"`
	// The variables in the TTS template, in the JSON format.
	//
	// > This parameter is required when the StartCode parameter is set to the ID of a TTS template that contains variables.
	StartTtsParams *string `json:"StartTtsParams,omitempty" xml:"StartTtsParams,omitempty"`
	// The timeout period for the subscriber to press a key. Unit: milliseconds.
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
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
	// The voice that corresponds to the key specified by the **MenuKeyMap.N.Key** parameter.
	//
	// *   If you use a voice notification file, this parameter specifies the voice ID. You can log on to the [Voice Messaging Service console](https://dyvms.console.aliyun.com/overview/home), choose **Voice Messages** > **Voice Notifications**, and then click the **Voice Notification Files** tab to view the voice ID.
	// *   If you use a TTS template, this parameter specifies the template ID. You can log on to the [Voice Messaging Service console](https://dyvms.console.aliyun.com/overview/home), choose **Voice Messages** > **Voice Notifications**, and then click the **TTS Template** tab to view the template ID.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The key that can be pressed by the subscriber.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The variables in the TTS template, in the JSON format.
	//
	// >
	//
	// *   This parameter specifies the substitution relationship of the variables in the TTS template if the value of the **MenuKeyMap.N.Code** parameter is set to the ID of the TTS template.
	//
	// *   This parameter is required if the value of the **MenuKeyMap.N.Code** parameter is set to the ID of a TTS template that contains variables.
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
	// The unique receipt ID of the call.
	//
	// You can call the [QueryCallDetailByCallId](~~393529~~) operation to query the details of the call based on the receipt ID.
	CallId *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	// The response code.
	//
	// *   The value OK indicates that the request was successful.
	// *   For more information about other response codes, see [API error codes](~~112502~~).
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
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
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *IvrCallResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *IvrCallResponse) SetStatusCode(v int32) *IvrCallResponse {
	s.StatusCode = &v
	return s
}

func (s *IvrCallResponse) SetBody(v *IvrCallResponseBody) *IvrCallResponse {
	s.Body = v
	return s
}

type ListCallTaskRequest struct {
	// The type of the task template. Valid values:
	//
	// *   **VMS_VOICE_TTS**: the text-to-speech (TTS) notification template.
	// *   **VMS_VOICE_CODE**: the voice notification template.
	// *   **VMS_TTS**: the voice verification code template.
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Default value: **1**.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: **10**.
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The task state. Valid values:
	//
	// *   **INIT**: The task is in the initial state.
	// *   **RELEASE**: The task is being parsed.
	// *   **RUNNING**: The task is running.
	// *   **STOP**: The task is suspended.
	// *   **SYSTEM_STOP**: The task is suspended by the system.
	// *   **CANCEL**: The task is canceled.
	// *   **SYSTEM_CANCEL**: The task is canceled by the system.
	// *   **DONE**: The task is complete.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The task ID.
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The task name.
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// The template name.
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
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
	// The response code.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The task information.
	Data []*ListCallTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The page number.
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of tasks.
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
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
	// The type of the task template. Valid values:
	//
	// *   **VMS_VOICE_TTS**: the TTS notification template.
	// *   **VMS_VOICE_CODE**: the voice notification template.
	// *   **VMS_TTS**: the voice verification code template.
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// The time when the task was completed. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	CompleteTime *string `json:"CompleteTime,omitempty" xml:"CompleteTime,omitempty"`
	// The number of tasks that were complete.
	CompletedCount *int64 `json:"CompletedCount,omitempty" xml:"CompletedCount,omitempty"`
	// The task progress.
	CompletedRate *int32 `json:"CompletedRate,omitempty" xml:"CompletedRate,omitempty"`
	// This parameter is unavailable.
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The type of the called number.
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// The time when the scheduled task was started. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	FireTime *string `json:"FireTime,omitempty" xml:"FireTime,omitempty"`
	// The task ID.
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The calling number.
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	// The task state. Valid values:
	//
	// *   **INIT**: The task was in the initial state.
	// *   **RELEASE**: The task was being parsed.
	// *   **RUNNING**: The task was running.
	// *   **STOP**: The task was manually suspended.
	// *   **SYSTEM_STOP**: The task was suspended by the system.
	// *   **CANCEL**: The task was manually canceled.
	// *   **SYSTEM_CANCEL**: The task was canceled by the system.
	// *   **DONE**: The task was complete.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// This parameter is unavailable.
	StopTime *string `json:"StopTime,omitempty" xml:"StopTime,omitempty"`
	// The task name.
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// The ID of the voice template.
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	// The template name.
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The total number of called numbers.
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCallTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ListCallTaskResponse) SetStatusCode(v int32) *ListCallTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCallTaskResponse) SetBody(v *ListCallTaskResponseBody) *ListCallTaskResponse {
	s.Body = v
	return s
}

type ListCallTaskDetailRequest struct {
	// The called number.
	CalledNum *string `json:"CalledNum,omitempty" xml:"CalledNum,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Default value: **1**.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: **10**.
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The task state. Valid values:
	//
	// *   **SUCCESS**: The task is successful.
	// *   **FAIL**: The task fails.
	// *   **INIT**: The task is not started.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The task ID.
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
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
	// The response code.
	//
	// *   The value OK indicates that the request was successful.
	// *   For more information about other response codes, see [API error codes](~~112502~~).
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The information about the task.
	Data []*ListCallTaskDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The page number.
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of called numbers.
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
	// The total number of pages.
	TotalPage *int64 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
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
	// The called number.
	CalledNum *string `json:"CalledNum,omitempty" xml:"CalledNum,omitempty"`
	// The calling number.
	Caller *string `json:"Caller,omitempty" xml:"Caller,omitempty"`
	// The call duration. Unit: seconds.
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// This parameter is unavailable.
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The task state. Valid values:
	//
	// *   **SUCCESS**: The task was successful.
	// *   **FAIL**: The task failed.
	// *   **INIT**: The task was not started.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCallTaskDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ListCallTaskDetailResponse) SetStatusCode(v int32) *ListCallTaskDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCallTaskDetailResponse) SetBody(v *ListCallTaskDetailResponseBody) *ListCallTaskDetailResponse {
	s.Body = v
	return s
}

type ListHotlineTransferNumberRequest struct {
	// The China 400 number.
	HotlineNumber *string `json:"HotlineNumber,omitempty" xml:"HotlineNumber,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Default value: **1**.
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. Valid values: 1 to 10.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The qualification ID. You can call the [GetHotlineQualificationByOrder](~~393548~~) operation to obtain the qualification ID.
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
	// The response code.
	//
	// *   The value OK indicates that the request was successful.
	// *   For more information about other response codes, see [API error codes](~~112502~~).
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The information about the registered phone number.
	Data *ListHotlineTransferNumberResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// The page number.
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
	// The phone numbers.
	Values []*ListHotlineTransferNumberResponseBodyDataValues `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
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
	// The China 400 number.
	HotlineNumber *string `json:"HotlineNumber,omitempty" xml:"HotlineNumber,omitempty"`
	// The ID card number of the number owner.
	IdentityCard *string `json:"IdentityCard,omitempty" xml:"IdentityCard,omitempty"`
	// The real name of the number owner or the company name.
	NumberOwnerName *string `json:"NumberOwnerName,omitempty" xml:"NumberOwnerName,omitempty"`
	// The registered phone number.
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// The qualification ID.
	QualificationId *string `json:"QualificationId,omitempty" xml:"QualificationId,omitempty"`
	// The resource code.
	ResUniqueCode *string `json:"ResUniqueCode,omitempty" xml:"ResUniqueCode,omitempty"`
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
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListHotlineTransferNumberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ListHotlineTransferNumberResponse) SetStatusCode(v int32) *ListHotlineTransferNumberResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHotlineTransferNumberResponse) SetBody(v *ListHotlineTransferNumberResponseBody) *ListHotlineTransferNumberResponse {
	s.Body = v
	return s
}

type ListHotlineTransferRegisterFileRequest struct {
	// The China 400 number.
	HotlineNumber *string `json:"HotlineNumber,omitempty" xml:"HotlineNumber,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Default value: **1**.
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. Valid values: 1 to 10.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The qualification ID. You can call the [GetHotlineQualificationByOrder](~~393548~~) operation to obtain the qualification ID.
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
	// The response code.
	//
	// *   The value OK indicates that the request was successful.
	// *   For more information about other response codes, see [API error codes](~~112502~~).
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response parameters.
	Data *ListHotlineTransferRegisterFileResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// The page number.
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
	// The registration file.
	Values []*ListHotlineTransferRegisterFileResponseBodyDataValues `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
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
	// The authenticity of the commitment.
	Agree *string `json:"Agree,omitempty" xml:"Agree,omitempty"`
	// The enterprise name.
	CorpName *string `json:"CorpName,omitempty" xml:"CorpName,omitempty"`
	// The China 400 number.
	HotlineNumber *string `json:"HotlineNumber,omitempty" xml:"HotlineNumber,omitempty"`
	// The ID card number of the handler.
	MngOpIdentityCard *string `json:"MngOpIdentityCard,omitempty" xml:"MngOpIdentityCard,omitempty"`
	// The email address of the handler.
	MngOpMail *string `json:"MngOpMail,omitempty" xml:"MngOpMail,omitempty"`
	// The mobile phone number of the handler.
	MngOpMobile *string `json:"MngOpMobile,omitempty" xml:"MngOpMobile,omitempty"`
	// The name of the handler.
	MngOpName *string `json:"MngOpName,omitempty" xml:"MngOpName,omitempty"`
	// The qualification ID.
	QualificationId *string `json:"QualificationId,omitempty" xml:"QualificationId,omitempty"`
	// The unique code of the query operation.
	ResUniqueCode *int64 `json:"ResUniqueCode,omitempty" xml:"ResUniqueCode,omitempty"`
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
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListHotlineTransferRegisterFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ListHotlineTransferRegisterFileResponse) SetStatusCode(v int32) *ListHotlineTransferRegisterFileResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHotlineTransferRegisterFileResponse) SetBody(v *ListHotlineTransferRegisterFileResponseBody) *ListHotlineTransferRegisterFileResponse {
	s.Body = v
	return s
}

type PauseVideoFileRequest struct {
	CallId               *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	CalledNumber         *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s PauseVideoFileRequest) String() string {
	return tea.Prettify(s)
}

func (s PauseVideoFileRequest) GoString() string {
	return s.String()
}

func (s *PauseVideoFileRequest) SetCallId(v string) *PauseVideoFileRequest {
	s.CallId = &v
	return s
}

func (s *PauseVideoFileRequest) SetCalledNumber(v string) *PauseVideoFileRequest {
	s.CalledNumber = &v
	return s
}

func (s *PauseVideoFileRequest) SetOwnerId(v int64) *PauseVideoFileRequest {
	s.OwnerId = &v
	return s
}

func (s *PauseVideoFileRequest) SetResourceOwnerAccount(v string) *PauseVideoFileRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *PauseVideoFileRequest) SetResourceOwnerId(v int64) *PauseVideoFileRequest {
	s.ResourceOwnerId = &v
	return s
}

type PauseVideoFileResponseBody struct {
	AccessDeniedDetail *string                `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Code               *string                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data               map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	Message            *string                `json:"Message,omitempty" xml:"Message,omitempty"`
	Success            *bool                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PauseVideoFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PauseVideoFileResponseBody) GoString() string {
	return s.String()
}

func (s *PauseVideoFileResponseBody) SetAccessDeniedDetail(v string) *PauseVideoFileResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *PauseVideoFileResponseBody) SetCode(v string) *PauseVideoFileResponseBody {
	s.Code = &v
	return s
}

func (s *PauseVideoFileResponseBody) SetData(v map[string]interface{}) *PauseVideoFileResponseBody {
	s.Data = v
	return s
}

func (s *PauseVideoFileResponseBody) SetMessage(v string) *PauseVideoFileResponseBody {
	s.Message = &v
	return s
}

func (s *PauseVideoFileResponseBody) SetSuccess(v bool) *PauseVideoFileResponseBody {
	s.Success = &v
	return s
}

type PauseVideoFileResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PauseVideoFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PauseVideoFileResponse) String() string {
	return tea.Prettify(s)
}

func (s PauseVideoFileResponse) GoString() string {
	return s.String()
}

func (s *PauseVideoFileResponse) SetHeaders(v map[string]*string) *PauseVideoFileResponse {
	s.Headers = v
	return s
}

func (s *PauseVideoFileResponse) SetStatusCode(v int32) *PauseVideoFileResponse {
	s.StatusCode = &v
	return s
}

func (s *PauseVideoFileResponse) SetBody(v *PauseVideoFileResponseBody) *PauseVideoFileResponse {
	s.Body = v
	return s
}

type PlayVideoFileRequest struct {
	CallId               *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	CalledNumber         *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	OutId                *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	VideoId              *string `json:"VideoId,omitempty" xml:"VideoId,omitempty"`
}

func (s PlayVideoFileRequest) String() string {
	return tea.Prettify(s)
}

func (s PlayVideoFileRequest) GoString() string {
	return s.String()
}

func (s *PlayVideoFileRequest) SetCallId(v string) *PlayVideoFileRequest {
	s.CallId = &v
	return s
}

func (s *PlayVideoFileRequest) SetCalledNumber(v string) *PlayVideoFileRequest {
	s.CalledNumber = &v
	return s
}

func (s *PlayVideoFileRequest) SetOutId(v string) *PlayVideoFileRequest {
	s.OutId = &v
	return s
}

func (s *PlayVideoFileRequest) SetOwnerId(v int64) *PlayVideoFileRequest {
	s.OwnerId = &v
	return s
}

func (s *PlayVideoFileRequest) SetResourceOwnerAccount(v string) *PlayVideoFileRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *PlayVideoFileRequest) SetResourceOwnerId(v int64) *PlayVideoFileRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *PlayVideoFileRequest) SetVideoId(v string) *PlayVideoFileRequest {
	s.VideoId = &v
	return s
}

type PlayVideoFileResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Code               *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message            *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Model              *bool   `json:"Model,omitempty" xml:"Model,omitempty"`
	Success            *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PlayVideoFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PlayVideoFileResponseBody) GoString() string {
	return s.String()
}

func (s *PlayVideoFileResponseBody) SetAccessDeniedDetail(v string) *PlayVideoFileResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *PlayVideoFileResponseBody) SetCode(v string) *PlayVideoFileResponseBody {
	s.Code = &v
	return s
}

func (s *PlayVideoFileResponseBody) SetMessage(v string) *PlayVideoFileResponseBody {
	s.Message = &v
	return s
}

func (s *PlayVideoFileResponseBody) SetModel(v bool) *PlayVideoFileResponseBody {
	s.Model = &v
	return s
}

func (s *PlayVideoFileResponseBody) SetSuccess(v bool) *PlayVideoFileResponseBody {
	s.Success = &v
	return s
}

type PlayVideoFileResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PlayVideoFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PlayVideoFileResponse) String() string {
	return tea.Prettify(s)
}

func (s PlayVideoFileResponse) GoString() string {
	return s.String()
}

func (s *PlayVideoFileResponse) SetHeaders(v map[string]*string) *PlayVideoFileResponse {
	s.Headers = v
	return s
}

func (s *PlayVideoFileResponse) SetStatusCode(v int32) *PlayVideoFileResponse {
	s.StatusCode = &v
	return s
}

func (s *PlayVideoFileResponse) SetBody(v *PlayVideoFileResponseBody) *PlayVideoFileResponse {
	s.Body = v
	return s
}

type QueryCallDetailByCallIdRequest struct {
	// The unique ID of the call.
	//
	// >
	//
	// *   The CallId parameter is included in the response parameters of the outbound call operation that you call to initiate a call.
	//
	// *   The date when the call specified by CallId is started must be the same as the date specified by QueryDate.
	//
	// *   The value of CallId must match the value of ProdId.
	CallId  *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The service ID. Valid values:
	//
	// *   **11000000300006**: voice notification. You can call the [SingleCallByVoice](https://help.aliyun.com/document_detail/393517.html) operation to send a voice notification of the voice notification file type to the specified number.
	// *   **11010000138001**: voice verification code. You can call the [SingleCallByTts](https://help.aliyun.com/document_detail/393519.html) operation to send a voice verification code or a text-to-speech (TTS) voice notification to the specified number.
	// *   **11000000300005**: IVR. You can call the [IvrCall](https://help.aliyun.com/document_detail/393521.html) operation to initiate an interactive voice call to the specified number.
	// *   **11000000300009**: Session Initiation Protocol (SIP) call.
	// *   **11030000180001**: intelligent outbound call.
	ProdId *int64 `json:"ProdId,omitempty" xml:"ProdId,omitempty"`
	// The time at which call details are queried. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// > The system queries the call records that are generated within 24 hours after the specified point in time. For example, if you specify the time 20:00:01 on November 21, 2022, the system queries the call records that are generated for the specified call ID from 20:00:01 on November 21, 2022, to 20:00:01 on November 22, 2022.
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
	// The response code.
	//
	// *   The value OK indicates that the request was successful.
	// *   For more information about other response codes, see [API error codes](~~112502~~).
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The details of the call, in the JSON format.
	//
	// *   **caller**: the calling number.
	// *   **startDate**: the time when the call was started.
	// *   **stateDesc**: the description of the call state.
	// *   **duration**: the call duration. Unit: seconds. The value **0** indicates that the user was not connected.
	// *   **callerShowNumber**: the calling number displayed to the called party.
	// *   **gmtCreate**: the time when the call request was received.
	// *   **state**: the call state. The call state is returned by the Internet service provider (ISP) in real time. For more information about call states, see [ISP-returned error codes](~~55085~~).
	// *   **endDate**: the time when the call was ended.
	// *   **calleeShowNumber**: the number displayed to the called party.
	// *   **callee**: the called number.
	// *   **aRingTime**: the time when Line A started to ring, in the yyyy-MM-dd HH:mm:ss format.
	// *   **aEndTime**: the time when ringing on Line A ended, in the yyyy-MM-dd HH:mm:ss format.
	// *   **bRingTime**: the time when Line B started to ring, in the yyyy-MM-dd HH:mm:ss format.
	// *   **bEndTime**: the time when ringing on Line B ended, in the yyyy-MM-dd HH:mm:ss format.
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
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
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryCallDetailByCallIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *QueryCallDetailByCallIdResponse) SetStatusCode(v int32) *QueryCallDetailByCallIdResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryCallDetailByCallIdResponse) SetBody(v *QueryCallDetailByCallIdResponseBody) *QueryCallDetailByCallIdResponse {
	s.Body = v
	return s
}

type QueryCallDetailByTaskIdRequest struct {
	// The called number. You can view the outbound call records of only one called number.
	Callee  *string `json:"Callee,omitempty" xml:"Callee,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The start time of the outbound robocall task. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	QueryDate            *int64  `json:"QueryDate,omitempty" xml:"QueryDate,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The unique ID of the outbound robocall task. The task ID is returned after the outbound robocall task is successfully delivered. You can view the task ID on the [Task Management](https://dyvms.console.aliyun.com/job/list) page of the Voice Messaging Service console, or call the **BatchRobotSmartCall** operation to obtain the **task ID**.
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
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
	// The response code.
	//
	// *   The value OK indicates that the request was successful.
	// *   For more information about other response codes, see [API error codes](~~112502~~).
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The call details of the outbound robocall task, in the JSON format.
	//
	// *   **startDate**: the time when the call was answered.
	//
	// *   **stateDesc**: the reason why the call was hung up. If the status code of early media was returned, this parameter indicates the reason why the status code of early media was used.
	//
	// *   **statusCode**: the status code.
	//
	// *   **EndDate**: the time when the call was ended.
	//
	// *   **calleeShowNumber**: the calling number displayed to the called party.
	//
	// *   **callee**: the called number.
	//
	// *   **duration**: the billing duration.
	//
	// *   **gmtCreate**: the time when the outbound robocall task was created.
	//
	// *   **hangupDirection**: the party who hung up.
	//
	// *   **tags**: the call tags.
	//
	// *   **dialogCount**: the number of conversation rounds in the call.
	//
	// *   **sureCount**: the number of times that the robocall task was acknowledged.
	//
	// *   **denyCount**: the number of times that the robocall task was denied.
	//
	// *   **rejectCount**: the number of times that the robocall task was rejected.
	//
	// *   **customCount**: the number of times that the robocall task was customized.
	//
	// *   **knowledgeCount**: the number of times that the knowledge base was queried.
	//
	// *   **recordFile**: the download URL of the recording file. The URL is valid only for 48 hours. The recording file must be downloaded in time.
	//
	// *   **callId**: the call ID.
	//
	// *   **recordStatus**: indicates whether a recording file was available. Valid values:
	//
	//     *   1: A recording file was available.
	//     *   2: No recording file was available.
	//
	// *   **knowledgeCommonCount**: the number of call failures caused by the common issues in the knowledge base.
	//
	// *   **knowledgeBusinessCount**: the number of call failures caused by the business issues in the knowledge base.
	//
	// *   **callee**: the called number.
	//
	// *   **dialogDetail**: the conversation details. The value is a JSON array that contains the following parameters:
	//
	//     *   **role**: the role who spoke.
	//     *   **content**: the content of the speech.
	//     *   **time**: the start time of the speech.
	//
	// > The preceding parameters are for reference only. The actually returned parameters prevail.
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
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
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryCallDetailByTaskIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *QueryCallDetailByTaskIdResponse) SetStatusCode(v int32) *QueryCallDetailByTaskIdResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryCallDetailByTaskIdResponse) SetBody(v *QueryCallDetailByTaskIdResponseBody) *QueryCallDetailByTaskIdResponse {
	s.Body = v
	return s
}

type QueryCallInPoolTransferConfigRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The China 400 number used to transfer the call.
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
	// The response code.
	//
	// *   The value OK indicates that the request was successful.
	// *   For more information about other response codes, see [API error codes](~~112502~~).
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response parameters.
	Data *QueryCallInPoolTransferConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// The call mode. Valid values:
	//
	// *   **roundRobin**
	// *   **random**
	CalledRouteMode *string `json:"CalledRouteMode,omitempty" xml:"CalledRouteMode,omitempty"`
	// The details of the response parameters.
	Details []*QueryCallInPoolTransferConfigResponseBodyDataDetails `json:"Details,omitempty" xml:"Details,omitempty" type:"Repeated"`
	// The time when the call transfer task was created.
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The timeout period for transferring the call.
	TransferTimeout *string `json:"TransferTimeout,omitempty" xml:"TransferTimeout,omitempty"`
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
	// The number used to transfer the call.
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
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryCallInPoolTransferConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *QueryCallInPoolTransferConfigResponse) SetStatusCode(v int32) *QueryCallInPoolTransferConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryCallInPoolTransferConfigResponse) SetBody(v *QueryCallInPoolTransferConfigResponseBody) *QueryCallInPoolTransferConfigResponse {
	s.Body = v
	return s
}

type QueryCallInTransferRecordRequest struct {
	// The calling number of the inbound call.
	CallInCaller *string `json:"CallInCaller,omitempty" xml:"CallInCaller,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Default value: **1**.
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. Valid values: 1 to 10.
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The phone number to which a call is transferred.
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// The time at which call transfer records are queried, in the YYYY-MM-DD hh:mm:ss format.
	//
	// > The query result is all the call transfer records of the specified day.
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
	// The response code.
	//
	// *   The value OK indicates that the request was successful.
	// *   For more information about other response codes, see [API error codes](~~112502~~).
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response parameters.
	Data *QueryCallInTransferRecordResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// The page number.
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page.
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
	// The call transfer records.
	Values []*QueryCallInTransferRecordResponseBodyDataValues `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
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
	// The called number of the inbound call.
	CallInCalled *string `json:"CallInCalled,omitempty" xml:"CallInCalled,omitempty"`
	// The calling number of the inbound call.
	CallInCaller *string `json:"CallInCaller,omitempty" xml:"CallInCaller,omitempty"`
	// The time when the call was initiated.
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The recording URL.
	RecordUrl *string `json:"RecordUrl,omitempty" xml:"RecordUrl,omitempty"`
	// The phone number to which the call was transferred.
	TransferCalled *string `json:"TransferCalled,omitempty" xml:"TransferCalled,omitempty"`
	// The calling number that transferred the call.
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
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryCallInTransferRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *QueryCallInTransferRecordResponse) SetStatusCode(v int32) *QueryCallInTransferRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryCallInTransferRecordResponse) SetBody(v *QueryCallInTransferRecordResponseBody) *QueryCallInTransferRecordResponse {
	s.Body = v
	return s
}

type QueryRobotInfoListRequest struct {
	// The review state. Valid values:
	//
	// *   **CONFIGURABLE**
	// *   **AUDITING**
	// *   **AUDITPASS**
	// *   **AUDITFAIL**
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
	// The response code.
	//
	// *   The value OK indicates that the request was successful.
	// *   For more information about other response codes, see [API error codes](~~112502~~).
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The basic information about the robot, in the JSON format. The basic information contains the following parameters:
	//
	// *   **id**: the robot ID.
	// *   **robotName**: the robot name.
	// *   **robotType**: the robot type.
	// *   **auditStatus**: the review state.
	// *   **gmtCreate**: the time when the task was created.
	// *   **gmtModified**: the time when the task was modified.
	// *   **partnerId**: the partner ID.
	// *   **asrId**: the ID of the automatic speech recognition (ASR) model.
	// *   **asrType**: the ASR type. Valid values: **Public** and **Private**.
	// *   **remark**: the additional information.
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryRobotInfoListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *QueryRobotInfoListResponse) SetStatusCode(v int32) *QueryRobotInfoListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryRobotInfoListResponse) SetBody(v *QueryRobotInfoListResponseBody) *QueryRobotInfoListResponse {
	s.Body = v
	return s
}

type QueryRobotTaskCallDetailRequest struct {
	// The called number.
	Callee  *string `json:"Callee,omitempty" xml:"Callee,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The timestamp of the time at which the call details you want to query.
	QueryDate            *int64  `json:"QueryDate,omitempty" xml:"QueryDate,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The unique ID of the robocall task. You can call the [CreateRobotTask](~~393531~~) operation to obtain the task ID.
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
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
	// The response code.
	//
	// *   The value OK indicates that the request was successful.
	// *   For more information about other response codes, see [API error codes](~~112502~~).
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The call details of a robocall task, in the JSON format.
	//
	// *   **taskId**: the unique ID of the robocall task.
	// *   **caller**: the calling number.
	// *   **called**: the called number.
	// *   **duration**: the call duration. Unit: seconds.
	// *   **label**: the label of the called party.
	// *   **dialogCount**: the number of conversation rounds in the call.
	// *   **callResult**: the call result.
	// *   **hangupDirection**: the party who hung up. Valid values: **0**: the robot. **1**: the called party.
	// *   **transferResult**: the result of transferring the call to an agent. Valid values: **1**, **0**, and **3**. The value 1 indicates that the call was transferred to the agent. The value 0 indicates that the call failed to be transferred to the agent. The value 3 indicates that the call was not transferred to the agent.
	// *   **transferNumber**: the phone number of the agent to whom the call was transferred.
	// *   **transferFailReason**: the reason why the call failed to be transferred to the agent.
	// *   **callId**: the unique receipt ID of the call, in the `taskId^bizId` format.
	// *   **recallCurTimes**: the number of recalls.
	// *   **callStartTime**: the start time of the call.
	// *   **callEndTime**: the end time of the call.
	// *   **sureCount**: the number of times that the robocall task was affirmed.
	// *   **denyCount**: the number of times that the robocall task was denied.
	// *   **rejectCount**: the number of times that the robocall task was rejected.
	// *   **customCount**: the number of times that the robocall task was customized.
	// *   **knowledgeCount**: the number of times that the knowledge base was queried.
	// *   **defaultCount**: the default number of calls.
	// *   **knowledgeBusinessCount**: the number of call failures caused by the business issues in the knowledge base.
	// *   **knowledgeCommonCount**: the number of call failures caused by the common issues in the knowledge base.
	// *   **recordStatus**: Indicates whether the call has a recording file. Valid values: **1**: The call has a recording file. **2**: The call does not have a recording file.
	// *   **recordFile**: the download URL of the recording file.
	// *   **dialogDetail**: the dialog details, in a JSON-formatted array. **role**: the object of the speech. **content**: the content of the speech. **speakTime**: the time of the speech.
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
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
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryRobotTaskCallDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *QueryRobotTaskCallDetailResponse) SetStatusCode(v int32) *QueryRobotTaskCallDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryRobotTaskCallDetailResponse) SetBody(v *QueryRobotTaskCallDetailResponseBody) *QueryRobotTaskCallDetailResponse {
	s.Body = v
	return s
}

type QueryRobotTaskCallListRequest struct {
	// The call result. Valid values:
	//
	// *   **200002**: The line is busy.
	// *   **200005**: The called party cannot be connected.
	// *   **200010**: The phone of the called party is powered off.
	// *   **200011**: The called party is out of service.
	// *   **200012**: The call is lost.
	CallResult *string `json:"CallResult,omitempty" xml:"CallResult,omitempty"`
	// The called number.
	Called *string `json:"Called,omitempty" xml:"Called,omitempty"`
	// The minimum number of conversation rounds in the call.
	DialogCountFrom *string `json:"DialogCountFrom,omitempty" xml:"DialogCountFrom,omitempty"`
	// The maximum number of conversation rounds in the call.
	DialogCountTo *string `json:"DialogCountTo,omitempty" xml:"DialogCountTo,omitempty"`
	// The minimum call duration.
	DurationFrom *string `json:"DurationFrom,omitempty" xml:"DurationFrom,omitempty"`
	// The maximum call duration.
	DurationTo *string `json:"DurationTo,omitempty" xml:"DurationTo,omitempty"`
	// The party who hangs up. Valid values:
	//
	// *   **0**: the called party.
	// *   **1**: the robot.
	HangupDirection *string `json:"HangupDirection,omitempty" xml:"HangupDirection,omitempty"`
	OwnerId         *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number.
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page.
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The unique ID of the robocall task. You can call the [CreateRobotTask](~~393531~~) operation to obtain the task ID.
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
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
	// The response code.
	//
	// *   The value OK indicates that the request was successful.
	// *   For more information about other response codes, see [API error codes](~~112502~~).
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The information about the robocall task, which is a JSON-formatted array.
	//
	// *   **taskId**: the unique ID of the robocall task.
	// *   **caller**: the calling number.
	// *   **called**: the called number.
	// *   **duration**: the call duration. Unit: seconds.
	// *   **label**: the label of the called party.
	// *   **dialogCount**: the number of conversation rounds in the call.
	// *   **callResult**: the call result.
	// *   **hangupDirection**: the party who hung up. Valid values: **1** and **0**. The value 1 indicates the called party, and the value 0 indicates the robot.
	// *   **transferResult**: the result of transferring the call to an agent. Valid values: **1**, **0**, and **3**. The value 1 indicates that the call was transferred to an agent. The value 0 indicates that the call failed to be transferred to an agent. The value 3 indicates that the call was not transferred to an agent.
	// *   **transferNumber**: the phone number of the agent to whom the call was transferred.
	// *   **transferFailReason**: the reason why the call failed to be transferred to an agent.
	// *   **callId**: the unique receipt ID of the call.
	// *   **recallCurTimes**: the number of recalls.
	// *   **callStartTime**: the start time of the call.
	// *   **callEndTime**: the end time of the call.
	// *   **sureCount**: the number of times that the robocall task was acknowledged.
	// *   **denyCount**: the number of times that the robocall task was denied.
	// *   **rejectCount**: the number of times that the robocall task was rejected.
	// *   **customCount**: the number of times that the robocall task was customized.
	// *   **knowledgeCount**: the number of times that the knowledge base was queried.
	// *   **defaultCount**: the default number of calls.
	// *   **knowledgeBusinessCount**: the number of call failures caused by the business issues in the knowledge base.
	// *   **knowledgeCommonCount**: the number of call failures caused by the common issues in the knowledge base.
	// *   ****
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
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
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryRobotTaskCallListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *QueryRobotTaskCallListResponse) SetStatusCode(v int32) *QueryRobotTaskCallListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryRobotTaskCallListResponse) SetBody(v *QueryRobotTaskCallListResponseBody) *QueryRobotTaskCallListResponse {
	s.Body = v
	return s
}

type QueryRobotTaskDetailRequest struct {
	// The unique ID of the robocall task. You can call the [CreateRobotTask](~~CreateRobotTask~~) operation to obtain the task ID.
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
	// The response code.
	//
	// *   The value OK indicates that the request was successful.
	// *   For more information about other response codes, see [API error codes](~~112502~~).
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The details of the robocall task, in the JSON format.
	//
	// *   **Id**: the unique ID of the robocall task.
	// *   **taskName**: the task name.
	// *   **robotId**: the robot ID.
	// *   **robotName**: the robot name.
	// *   **corpName**: the company name.
	// *   **caller**: the number displayed to the called party.
	// *   **numberStatusIdent**: indicates whether number status identification was enabled. Valid values: **true** and **false**. The value true indicates that number status identification was enabled. The value false indicates that number status identification was not enabled.
	// *   **status**: the task state. You can call the [QueryRobotTaskList](~~QueryRobotTaskList~~) operation to obtain the task state from the `status` parameter.
	// *   **scheduleType**: the scheduling type. Valid values: **SINGLE** and **ORDER**. The value SINGLE indicates that the task was started immediately after it was created. The value ORDER indicates that the task was started at a scheduled time.
	// *   **retryType**: indicates whether auto-redial was enabled. Valid values: **1** and **0**. The value 1 indicates that auto-redial was enabled. The value 0 indicates that auto-redial was not enabled.
	// *   **recallStateCodes**: the call state in which redial is required. Valid values: **200010**, **200011**, **200002**, **200012**, and **200005**. The value 200010 indicates that the phone of the called party was powered off. The value 200011 indicates that the number of the called party was out of service. The value 200002 indicates that the line was busy. The value 200012 indicates that the call was lost. The value 200005 indicates that the called party could not be connected.
	// *   **recallTimes**: the number of redial times.
	// *   **recallInterval**: the redial interval. Unit: minutes.
	// *   **createTime**: the time when the task was created, in the yyyy-MM-dd HH:mm:ss format.
	// *   **fireTime**: the time when the task was started, in the yyyy-MM-dd HH:mm:ss format.
	// *   **completeTime**: the time when the task was completed, in the yyyy-MM-dd HH:mm:ss format.
	// *   **filename**: the name of the called number file.
	// *   **ossFilePath**: the path of the called number file.
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryRobotTaskDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *QueryRobotTaskDetailResponse) SetStatusCode(v int32) *QueryRobotTaskDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryRobotTaskDetailResponse) SetBody(v *QueryRobotTaskDetailResponseBody) *QueryRobotTaskDetailResponse {
	s.Body = v
	return s
}

type QueryRobotTaskListRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number.
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page.
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The task state. Valid values:
	//
	// *   **INIT**: The task is not started.
	// *   **READY**: The task is ready to start.
	// *   **DISPATCH**: The task is being parsed.
	// *   **EXCUTING**: The task is being executed.
	// *   **MANUAL_STOP**: The task is manually suspended.
	// *   **SYSTEM_STOP**: The task is suspended by the system.
	// *   **ARREARS_STOP**: The task is suspended due to overdue payments.
	// *   **CANCEL**: The task is manually canceled.
	// *   **SYSTEM_CANCEL**: The task is canceled by the system.
	// *   **FINISH**: The task is complete.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The task name.
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// The date when the task is created, in the yyyy-MM-dd format.
	Time *string `json:"Time,omitempty" xml:"Time,omitempty"`
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
	// The response code.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The robocall tasks, in the JSON format.
	//
	// *   **id**: the unique ID of the robocall task.
	// *   **taskName**: the task name.
	// *   **robotId**: the robot ID.
	// *   **robotName**: the robot name.
	// *   **status**: the task state.
	// *   **scheduleType**: the scheduling type. Valid values: **SINGLE** and **ORDER**. The value SINGLE indicates that the task was started immediately after it was created. The value ORDER indicates that the task was started at a scheduled time.
	// *   **createTime**: the time when the task was created, in the yyyy.MM.dd HH:mm:ss format.
	// *   **completeTime**: the time when the task was completed, in the yyyy.MM.dd HH:mm:ss format.
	// *   **fireTime**: the time when the task was started, in the yyyy.MM.dd HH:mm:ss format.
	// *   **totalCount**: the total number of calls processed.
	// *   **finishCount**: the number of calls completed.
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The page number.
	PageNo *string `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page.
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of tasks.
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryRobotTaskListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *QueryRobotTaskListResponse) SetStatusCode(v int32) *QueryRobotTaskListResponse {
	s.StatusCode = &v
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
	// The response code.
	//
	// *   The value OK indicates that the request was successful.
	// *   For more information about other response codes, see [API error codes](~~112502~~).
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The information about the robot communication script, in the JSON format.
	//
	// *   **id**: the ID of the robot communication script.
	// *   **robotName**: the name of the robot communication script.
	// *   **robotType**: the type of the robot communication script.
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryRobotv2AllListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *QueryRobotv2AllListResponse) SetStatusCode(v int32) *QueryRobotv2AllListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryRobotv2AllListResponse) SetBody(v *QueryRobotv2AllListResponseBody) *QueryRobotv2AllListResponse {
	s.Body = v
	return s
}

type QueryVideoPlayProgressRequest struct {
	CallId               *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	CalledNumber         *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s QueryVideoPlayProgressRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryVideoPlayProgressRequest) GoString() string {
	return s.String()
}

func (s *QueryVideoPlayProgressRequest) SetCallId(v string) *QueryVideoPlayProgressRequest {
	s.CallId = &v
	return s
}

func (s *QueryVideoPlayProgressRequest) SetCalledNumber(v string) *QueryVideoPlayProgressRequest {
	s.CalledNumber = &v
	return s
}

func (s *QueryVideoPlayProgressRequest) SetOwnerId(v int64) *QueryVideoPlayProgressRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryVideoPlayProgressRequest) SetResourceOwnerAccount(v string) *QueryVideoPlayProgressRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryVideoPlayProgressRequest) SetResourceOwnerId(v int64) *QueryVideoPlayProgressRequest {
	s.ResourceOwnerId = &v
	return s
}

type QueryVideoPlayProgressResponseBody struct {
	AccessDeniedDetail *string                `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Code               *string                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data               map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	Message            *string                `json:"Message,omitempty" xml:"Message,omitempty"`
	Success            *bool                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryVideoPlayProgressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryVideoPlayProgressResponseBody) GoString() string {
	return s.String()
}

func (s *QueryVideoPlayProgressResponseBody) SetAccessDeniedDetail(v string) *QueryVideoPlayProgressResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *QueryVideoPlayProgressResponseBody) SetCode(v string) *QueryVideoPlayProgressResponseBody {
	s.Code = &v
	return s
}

func (s *QueryVideoPlayProgressResponseBody) SetData(v map[string]interface{}) *QueryVideoPlayProgressResponseBody {
	s.Data = v
	return s
}

func (s *QueryVideoPlayProgressResponseBody) SetMessage(v string) *QueryVideoPlayProgressResponseBody {
	s.Message = &v
	return s
}

func (s *QueryVideoPlayProgressResponseBody) SetSuccess(v bool) *QueryVideoPlayProgressResponseBody {
	s.Success = &v
	return s
}

type QueryVideoPlayProgressResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryVideoPlayProgressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryVideoPlayProgressResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryVideoPlayProgressResponse) GoString() string {
	return s.String()
}

func (s *QueryVideoPlayProgressResponse) SetHeaders(v map[string]*string) *QueryVideoPlayProgressResponse {
	s.Headers = v
	return s
}

func (s *QueryVideoPlayProgressResponse) SetStatusCode(v int32) *QueryVideoPlayProgressResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryVideoPlayProgressResponse) SetBody(v *QueryVideoPlayProgressResponseBody) *QueryVideoPlayProgressResponse {
	s.Body = v
	return s
}

type QueryVirtualNumberRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number.
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The service name. Default value: **dyvms**.
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The route type. Valid values:
	//
	// *   **0**: number location first.
	// *   **1**: random.
	RouteType *int32 `json:"RouteType,omitempty" xml:"RouteType,omitempty"`
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
	// The response code. The value 200 indicates that the request was successful.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The details of the numbers associated with the virtual numbers. The following fields are returned:
	//
	// *   createTime: the time when the number was activated.
	// *   qualificationCount: the number of qualifications.
	// *   cityCount: the number of cities.
	// *   phoneNumCount: the number of virtual numbers.
	// *   remark: the additional information.
	// *   phoneNum: the virtual number.
	// *   routeType: the route type.
	// *   canCancel: indicates whether the number can be deactivated.
	// *   specCount: the number of Internet service providers (ISPs).
	// *   status: the number state. Valid values: **1**, **0**, and **-1**. The value 1 indicates that the number is valid. The value 0 indicates that the number is invalid. The value -1 indicates that the number was deactivated.
	// *   pageNo: the page number.
	// *   pageSize: the number of entries per page.
	// *   total: the total number of virtual numbers.
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The request ID.
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryVirtualNumberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *QueryVirtualNumberResponse) SetStatusCode(v int32) *QueryVirtualNumberResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryVirtualNumberResponse) SetBody(v *QueryVirtualNumberResponseBody) *QueryVirtualNumberResponse {
	s.Body = v
	return s
}

type QueryVirtualNumberRelationRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number.
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The virtual number.
	PhoneNum *string `json:"PhoneNum,omitempty" xml:"PhoneNum,omitempty"`
	// The service name. Default value: **dyvms**.
	ProdCode *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	// The qualification ID.
	//
	// You can log on to the [Voice Messaging Service console](https://dyvms.console.aliyun.com/overview/home), choose **Qualifications\&Communication Scripts > Qualification Management**, and then click **Details** in the Actions column to view the qualification ID.
	QualificationId *int64 `json:"QualificationId,omitempty" xml:"QualificationId,omitempty"`
	// The city to which the virtual number belongs.
	RegionNameCity *string `json:"RegionNameCity,omitempty" xml:"RegionNameCity,omitempty"`
	// The real number.
	RelatedNum           *string `json:"RelatedNum,omitempty" xml:"RelatedNum,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The route type. Valid values:
	//
	// **0**: number location first. **1**: random.
	RouteType *int32 `json:"RouteType,omitempty" xml:"RouteType,omitempty"`
	// The number type. Valid values:
	//
	// *   **1**: the number provided by a virtual network operator, in the 05710000\*\*\*\* format.
	// *   **2**: the number provided by an Internet service provider (ISP).
	// *   **3**: a 5-digit phone number that starts with 95.
	SpecId *int64 `json:"SpecId,omitempty" xml:"SpecId,omitempty"`
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
	// The response code.
	//
	// *   The value 200 indicates that the request was successful.
	// *   For more information about other response codes, see [API error codes](~~112502~~).
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The list of associations between virtual numbers and real numbers. The following fields are returned:
	//
	// *   **relatedNum**: the real number.
	// *   **createTime**: the time when the number was activated.
	// *   **pageNo**: the page number.
	// *   **pageSize**: the number of entries per page.
	// *   **total**: the total number of entries returned.
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The request ID.
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
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryVirtualNumberRelationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *QueryVirtualNumberRelationResponse) SetStatusCode(v int32) *QueryVirtualNumberRelationResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryVirtualNumberRelationResponse) SetBody(v *QueryVirtualNumberRelationResponseBody) *QueryVirtualNumberRelationResponse {
	s.Body = v
	return s
}

type QueryVoiceFileAuditInfoRequest struct {
	// The type of the voice file. Valid values:
	//
	// *   **0** (default): the voice notification file.
	// *   **2**: the recording file.
	BusinessType         *int32  `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the voice file. You can log on to the [Voice Messaging Service console](https://dyvms.console.aliyun.com/overview/home), choose **Voice Messages** > **Voice Notifications** or **Voice File Management**, and then click the **Voice Notification Files** tab to view the **voice ID**.
	//
	// > You can query up to 10 voice files each time. Separate the voice file names with commas (,).
	VoiceCodes *string `json:"VoiceCodes,omitempty" xml:"VoiceCodes,omitempty"`
}

func (s QueryVoiceFileAuditInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryVoiceFileAuditInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryVoiceFileAuditInfoRequest) SetBusinessType(v int32) *QueryVoiceFileAuditInfoRequest {
	s.BusinessType = &v
	return s
}

func (s *QueryVoiceFileAuditInfoRequest) SetOwnerId(v int64) *QueryVoiceFileAuditInfoRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryVoiceFileAuditInfoRequest) SetResourceOwnerAccount(v string) *QueryVoiceFileAuditInfoRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryVoiceFileAuditInfoRequest) SetResourceOwnerId(v int64) *QueryVoiceFileAuditInfoRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryVoiceFileAuditInfoRequest) SetVoiceCodes(v string) *QueryVoiceFileAuditInfoRequest {
	s.VoiceCodes = &v
	return s
}

type QueryVoiceFileAuditInfoResponseBody struct {
	// The response code.
	//
	// The value OK indicates that the request was successful. For more information about other response codes, see [API error codes](~~112502~~).
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response parameters.
	Data []*QueryVoiceFileAuditInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The returned message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryVoiceFileAuditInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryVoiceFileAuditInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryVoiceFileAuditInfoResponseBody) SetCode(v string) *QueryVoiceFileAuditInfoResponseBody {
	s.Code = &v
	return s
}

func (s *QueryVoiceFileAuditInfoResponseBody) SetData(v []*QueryVoiceFileAuditInfoResponseBodyData) *QueryVoiceFileAuditInfoResponseBody {
	s.Data = v
	return s
}

func (s *QueryVoiceFileAuditInfoResponseBody) SetMessage(v string) *QueryVoiceFileAuditInfoResponseBody {
	s.Message = &v
	return s
}

func (s *QueryVoiceFileAuditInfoResponseBody) SetRequestId(v string) *QueryVoiceFileAuditInfoResponseBody {
	s.RequestId = &v
	return s
}

type QueryVoiceFileAuditInfoResponseBodyData struct {
	// The review state of the voice file. Valid values:
	//
	// *   **AUDIT_STATE_INIT**: The voice file was under review.
	// *   **AUDIT_STATE_PASS**: The voice file was approved.
	// *   **AUDIT_STATE_NOT_PASS**: The voice file was rejected.
	// *   **AUDIT_STATE_UPLOADING**: The voice file was approved and is being uploaded.
	// *   **AUDIT_STATE_REDOING**: The voice file was being reprocessed.
	// *   **AUDIT_SATE_CANCEL**: The review of the voice file was canceled.
	// *   **AUDIT_PAUSE**: The review of the voice file was suspended.
	// *   **AUDIT_ORDER_FINISHED**: The voice file was approved by the ticket system and was waiting for the review of the Internet service provider (ISP).
	AuditState *string `json:"AuditState,omitempty" xml:"AuditState,omitempty"`
	// The reason why the voice file was rejected.
	RejectInfo *string `json:"RejectInfo,omitempty" xml:"RejectInfo,omitempty"`
	// The code of the voice file.
	VoiceCode *string `json:"VoiceCode,omitempty" xml:"VoiceCode,omitempty"`
}

func (s QueryVoiceFileAuditInfoResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryVoiceFileAuditInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryVoiceFileAuditInfoResponseBodyData) SetAuditState(v string) *QueryVoiceFileAuditInfoResponseBodyData {
	s.AuditState = &v
	return s
}

func (s *QueryVoiceFileAuditInfoResponseBodyData) SetRejectInfo(v string) *QueryVoiceFileAuditInfoResponseBodyData {
	s.RejectInfo = &v
	return s
}

func (s *QueryVoiceFileAuditInfoResponseBodyData) SetVoiceCode(v string) *QueryVoiceFileAuditInfoResponseBodyData {
	s.VoiceCode = &v
	return s
}

type QueryVoiceFileAuditInfoResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryVoiceFileAuditInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryVoiceFileAuditInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryVoiceFileAuditInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryVoiceFileAuditInfoResponse) SetHeaders(v map[string]*string) *QueryVoiceFileAuditInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryVoiceFileAuditInfoResponse) SetStatusCode(v int32) *QueryVoiceFileAuditInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryVoiceFileAuditInfoResponse) SetBody(v *QueryVoiceFileAuditInfoResponseBody) *QueryVoiceFileAuditInfoResponse {
	s.Body = v
	return s
}

type RecoverCallInConfigRequest struct {
	// The China 400 number that is used to transfer the inbound call.
	Number               *string `json:"Number,omitempty" xml:"Number,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s RecoverCallInConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s RecoverCallInConfigRequest) GoString() string {
	return s.String()
}

func (s *RecoverCallInConfigRequest) SetNumber(v string) *RecoverCallInConfigRequest {
	s.Number = &v
	return s
}

func (s *RecoverCallInConfigRequest) SetOwnerId(v int64) *RecoverCallInConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *RecoverCallInConfigRequest) SetResourceOwnerAccount(v string) *RecoverCallInConfigRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RecoverCallInConfigRequest) SetResourceOwnerId(v int64) *RecoverCallInConfigRequest {
	s.ResourceOwnerId = &v
	return s
}

type RecoverCallInConfigResponseBody struct {
	// The response code.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the inbound call was resumed. Valid values:
	//
	// *   true: The inbound call was resumed.
	// *   false: The inbound call failed to be resumed.
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecoverCallInConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecoverCallInConfigResponseBody) GoString() string {
	return s.String()
}

func (s *RecoverCallInConfigResponseBody) SetCode(v string) *RecoverCallInConfigResponseBody {
	s.Code = &v
	return s
}

func (s *RecoverCallInConfigResponseBody) SetData(v bool) *RecoverCallInConfigResponseBody {
	s.Data = &v
	return s
}

func (s *RecoverCallInConfigResponseBody) SetMessage(v string) *RecoverCallInConfigResponseBody {
	s.Message = &v
	return s
}

func (s *RecoverCallInConfigResponseBody) SetRequestId(v string) *RecoverCallInConfigResponseBody {
	s.RequestId = &v
	return s
}

type RecoverCallInConfigResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecoverCallInConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecoverCallInConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s RecoverCallInConfigResponse) GoString() string {
	return s.String()
}

func (s *RecoverCallInConfigResponse) SetHeaders(v map[string]*string) *RecoverCallInConfigResponse {
	s.Headers = v
	return s
}

func (s *RecoverCallInConfigResponse) SetStatusCode(v int32) *RecoverCallInConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *RecoverCallInConfigResponse) SetBody(v *RecoverCallInConfigResponseBody) *RecoverCallInConfigResponse {
	s.Body = v
	return s
}

type ResumeVideoFileRequest struct {
	CallId               *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	CalledNumber         *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ResumeVideoFileRequest) String() string {
	return tea.Prettify(s)
}

func (s ResumeVideoFileRequest) GoString() string {
	return s.String()
}

func (s *ResumeVideoFileRequest) SetCallId(v string) *ResumeVideoFileRequest {
	s.CallId = &v
	return s
}

func (s *ResumeVideoFileRequest) SetCalledNumber(v string) *ResumeVideoFileRequest {
	s.CalledNumber = &v
	return s
}

func (s *ResumeVideoFileRequest) SetOwnerId(v int64) *ResumeVideoFileRequest {
	s.OwnerId = &v
	return s
}

func (s *ResumeVideoFileRequest) SetResourceOwnerAccount(v string) *ResumeVideoFileRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ResumeVideoFileRequest) SetResourceOwnerId(v int64) *ResumeVideoFileRequest {
	s.ResourceOwnerId = &v
	return s
}

type ResumeVideoFileResponseBody struct {
	AccessDeniedDetail *string                `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Code               *string                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data               map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	Message            *string                `json:"Message,omitempty" xml:"Message,omitempty"`
	Success            *bool                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ResumeVideoFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResumeVideoFileResponseBody) GoString() string {
	return s.String()
}

func (s *ResumeVideoFileResponseBody) SetAccessDeniedDetail(v string) *ResumeVideoFileResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ResumeVideoFileResponseBody) SetCode(v string) *ResumeVideoFileResponseBody {
	s.Code = &v
	return s
}

func (s *ResumeVideoFileResponseBody) SetData(v map[string]interface{}) *ResumeVideoFileResponseBody {
	s.Data = v
	return s
}

func (s *ResumeVideoFileResponseBody) SetMessage(v string) *ResumeVideoFileResponseBody {
	s.Message = &v
	return s
}

func (s *ResumeVideoFileResponseBody) SetSuccess(v bool) *ResumeVideoFileResponseBody {
	s.Success = &v
	return s
}

type ResumeVideoFileResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResumeVideoFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResumeVideoFileResponse) String() string {
	return tea.Prettify(s)
}

func (s ResumeVideoFileResponse) GoString() string {
	return s.String()
}

func (s *ResumeVideoFileResponse) SetHeaders(v map[string]*string) *ResumeVideoFileResponse {
	s.Headers = v
	return s
}

func (s *ResumeVideoFileResponse) SetStatusCode(v int32) *ResumeVideoFileResponse {
	s.StatusCode = &v
	return s
}

func (s *ResumeVideoFileResponse) SetBody(v *ResumeVideoFileResponseBody) *ResumeVideoFileResponse {
	s.Body = v
	return s
}

type SeekVideoFileRequest struct {
	// 呼叫唯一ID
	CallId *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	// 被叫号码
	CalledNumber         *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 快进或快退值，负数为快退，单位秒
	SeekTimes *int64 `json:"SeekTimes,omitempty" xml:"SeekTimes,omitempty"`
}

func (s SeekVideoFileRequest) String() string {
	return tea.Prettify(s)
}

func (s SeekVideoFileRequest) GoString() string {
	return s.String()
}

func (s *SeekVideoFileRequest) SetCallId(v string) *SeekVideoFileRequest {
	s.CallId = &v
	return s
}

func (s *SeekVideoFileRequest) SetCalledNumber(v string) *SeekVideoFileRequest {
	s.CalledNumber = &v
	return s
}

func (s *SeekVideoFileRequest) SetOwnerId(v int64) *SeekVideoFileRequest {
	s.OwnerId = &v
	return s
}

func (s *SeekVideoFileRequest) SetResourceOwnerAccount(v string) *SeekVideoFileRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SeekVideoFileRequest) SetResourceOwnerId(v int64) *SeekVideoFileRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SeekVideoFileRequest) SetSeekTimes(v int64) *SeekVideoFileRequest {
	s.SeekTimes = &v
	return s
}

type SeekVideoFileResponseBody struct {
	AccessDeniedDetail *string                `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Code               *string                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data               map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	Message            *string                `json:"Message,omitempty" xml:"Message,omitempty"`
	Success            *bool                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SeekVideoFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SeekVideoFileResponseBody) GoString() string {
	return s.String()
}

func (s *SeekVideoFileResponseBody) SetAccessDeniedDetail(v string) *SeekVideoFileResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *SeekVideoFileResponseBody) SetCode(v string) *SeekVideoFileResponseBody {
	s.Code = &v
	return s
}

func (s *SeekVideoFileResponseBody) SetData(v map[string]interface{}) *SeekVideoFileResponseBody {
	s.Data = v
	return s
}

func (s *SeekVideoFileResponseBody) SetMessage(v string) *SeekVideoFileResponseBody {
	s.Message = &v
	return s
}

func (s *SeekVideoFileResponseBody) SetSuccess(v bool) *SeekVideoFileResponseBody {
	s.Success = &v
	return s
}

type SeekVideoFileResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SeekVideoFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SeekVideoFileResponse) String() string {
	return tea.Prettify(s)
}

func (s SeekVideoFileResponse) GoString() string {
	return s.String()
}

func (s *SeekVideoFileResponse) SetHeaders(v map[string]*string) *SeekVideoFileResponse {
	s.Headers = v
	return s
}

func (s *SeekVideoFileResponse) SetStatusCode(v int32) *SeekVideoFileResponse {
	s.StatusCode = &v
	return s
}

func (s *SeekVideoFileResponse) SetBody(v *SeekVideoFileResponseBody) *SeekVideoFileResponse {
	s.Body = v
	return s
}

type SendVerificationRequest struct {
	// The business type. Set the value to **CONTACT**.
	BizType              *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The mobile phone number that receives the SMS verification code.
	Target *string `json:"Target,omitempty" xml:"Target,omitempty"`
	// The mode of sending the SMS verification code. Set the value to **SMS**.
	VerifyType *string `json:"VerifyType,omitempty" xml:"VerifyType,omitempty"`
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
	// The response code.
	//
	// *   The value OK indicates that the request was successful.
	// *   For more information about other response codes, see [API error codes](~~112502~~).
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the verification code was sent successfully.
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendVerificationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *SendVerificationResponse) SetStatusCode(v int32) *SendVerificationResponse {
	s.StatusCode = &v
	return s
}

func (s *SendVerificationResponse) SetBody(v *SendVerificationResponseBody) *SendVerificationResponse {
	s.Body = v
	return s
}

type SetTransferCalleePoolConfigRequest struct {
	// The call mode. Valid values:
	//
	// *   **roundRobin**
	// *   **random**
	CalledRouteMode *string `json:"CalledRouteMode,omitempty" xml:"CalledRouteMode,omitempty"`
	// The information about the phone numbers for transferring the call.
	Details []*SetTransferCalleePoolConfigRequestDetails `json:"Details,omitempty" xml:"Details,omitempty" type:"Repeated"`
	OwnerId *int64                                       `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The phone number used for transferring the call.
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// The qualification ID. You can call the [GetHotlineQualificationByOrder](~~393548~~) operation to obtain the qualification ID.
	QualificationId      *string `json:"QualificationId,omitempty" xml:"QualificationId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
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
	// The called number.
	Called *string `json:"Called,omitempty" xml:"Called,omitempty"`
	// The calling number.
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
	// The response code.
	//
	// *   The value OK indicates that the request was successful.
	// *   For more information about other response codes, see [API error codes](~~112502~~).
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the phone numbers for transferring the call were configured.
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
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
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetTransferCalleePoolConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *SetTransferCalleePoolConfigResponse) SetStatusCode(v int32) *SetTransferCalleePoolConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *SetTransferCalleePoolConfigResponse) SetBody(v *SetTransferCalleePoolConfigResponseBody) *SetTransferCalleePoolConfigResponse {
	s.Body = v
	return s
}

type SingleCallByTtsRequest struct {
	// The mobile phone number that receives voice notifications.
	//
	// *   Number format in the Chinese mainland:
	//
	//     *   Mobile phone number, for example, 159\*\*\*\*0000.
	//     *   Landline number, for example, 0571\*\*\*\*5678.
	//
	// *   Number format outside the Chinese mainland: country code + phone number, for example, 85200\*\*\*\*00.
	//
	// >
	//
	// *   Each request supports only one called number. For more information, see [How to use voice notifications in the Chinese mainland](~~150016~~) or [How to use voice verification codes in regions outside the Chinese mainland](~~270044~~).
	//
	// *   Voice verification codes are sent to a called number at the following frequency: one time per minute, five times per hour, and 20 times per 24 hours.
	CalledNumber *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	// The number displayed to the called party.
	//
	// *   You do not need to specify this parameter if you use the text-to-speech (TTS) notification template or voice verification code template for outbound calls in the common mode. For more information, see [FAQ about the common outbound call mode](~~172104~~).
	// *   If you use the TTS notification template or voice verification code template for outbound calls in the dedicated mode, you must specify a number you purchased and only one number can be specified. You can log on to the [Voice Messaging Service console](https://dyvms.console.aliyun.com/overview/home) and choose **Voice Numbers** > **Real Number Management** to view the number you purchased.
	CalledShowNumber *string `json:"CalledShowNumber,omitempty" xml:"CalledShowNumber,omitempty"`
	// The custom ID that is reserved for the caller of the operation when the request is initiated. This ID is returned to the caller in a receipt message.
	//
	// The value is of the STRING type and must be 1 to 15 bytes in length.
	OutId   *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of times a voice notification is played back in a call. Valid values: 1 to 3. Default value: 3.
	PlayTimes            *int32  `json:"PlayTimes,omitempty" xml:"PlayTimes,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The playback speed. Valid value: -500 to 500.
	Speed *int32 `json:"Speed,omitempty" xml:"Speed,omitempty"`
	// The ID of the approved TTS notification template or voice verification code template.
	//
	// You can log on to the [Voice Messaging Service console](https://dyvms.console.aliyun.com/overview/home), and choose **Voice Messages** > **Voice Verification Codes** or choose **Voice Messages** > **Voice Notifications** to view the **template ID**.
	//
	// > The account to which the TTS template belongs must be the same as the account that is used to call the SingleCallByTts operation.
	TtsCode *string `json:"TtsCode,omitempty" xml:"TtsCode,omitempty"`
	// The variables in the template, in the JSON format.
	//
	// > The variables in the template must be less than 250 characters in length. The length of each single variable is not limited. These variables do not support URLs. The variables in the verification code template support only digits and letters.
	TtsParam *string `json:"TtsParam,omitempty" xml:"TtsParam,omitempty"`
	// The playback volume of the voice notification. Valid values: 0 to 100. Default value: 100.
	Volume *int32 `json:"Volume,omitempty" xml:"Volume,omitempty"`
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
	// The unique receipt ID of the call.
	//
	// You can call the [QueryCallDetailByCallId](~~393529~~) operation to query the details of the call based on the receipt ID.
	CallId *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	// The response code.
	//
	// *   The value OK indicates that the request was successful.
	// *   For more information about other response codes, see [API error codes](~~112502~~).
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SingleCallByTtsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *SingleCallByTtsResponse) SetStatusCode(v int32) *SingleCallByTtsResponse {
	s.StatusCode = &v
	return s
}

func (s *SingleCallByTtsResponse) SetBody(v *SingleCallByTtsResponseBody) *SingleCallByTtsResponse {
	s.Body = v
	return s
}

type SingleCallByVideoRequest struct {
	CalledNumber         *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	CalledShowNumber     *string `json:"CalledShowNumber,omitempty" xml:"CalledShowNumber,omitempty"`
	OutId                *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PlayTimes            *int32  `json:"PlayTimes,omitempty" xml:"PlayTimes,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Speed                *int32  `json:"Speed,omitempty" xml:"Speed,omitempty"`
	VideoCode            *string `json:"VideoCode,omitempty" xml:"VideoCode,omitempty"`
	VoiceCode            *string `json:"VoiceCode,omitempty" xml:"VoiceCode,omitempty"`
	Volume               *int32  `json:"Volume,omitempty" xml:"Volume,omitempty"`
}

func (s SingleCallByVideoRequest) String() string {
	return tea.Prettify(s)
}

func (s SingleCallByVideoRequest) GoString() string {
	return s.String()
}

func (s *SingleCallByVideoRequest) SetCalledNumber(v string) *SingleCallByVideoRequest {
	s.CalledNumber = &v
	return s
}

func (s *SingleCallByVideoRequest) SetCalledShowNumber(v string) *SingleCallByVideoRequest {
	s.CalledShowNumber = &v
	return s
}

func (s *SingleCallByVideoRequest) SetOutId(v string) *SingleCallByVideoRequest {
	s.OutId = &v
	return s
}

func (s *SingleCallByVideoRequest) SetOwnerId(v int64) *SingleCallByVideoRequest {
	s.OwnerId = &v
	return s
}

func (s *SingleCallByVideoRequest) SetPlayTimes(v int32) *SingleCallByVideoRequest {
	s.PlayTimes = &v
	return s
}

func (s *SingleCallByVideoRequest) SetResourceOwnerAccount(v string) *SingleCallByVideoRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SingleCallByVideoRequest) SetResourceOwnerId(v int64) *SingleCallByVideoRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SingleCallByVideoRequest) SetSpeed(v int32) *SingleCallByVideoRequest {
	s.Speed = &v
	return s
}

func (s *SingleCallByVideoRequest) SetVideoCode(v string) *SingleCallByVideoRequest {
	s.VideoCode = &v
	return s
}

func (s *SingleCallByVideoRequest) SetVoiceCode(v string) *SingleCallByVideoRequest {
	s.VoiceCode = &v
	return s
}

func (s *SingleCallByVideoRequest) SetVolume(v int32) *SingleCallByVideoRequest {
	s.Volume = &v
	return s
}

type SingleCallByVideoResponseBody struct {
	CallId    *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SingleCallByVideoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SingleCallByVideoResponseBody) GoString() string {
	return s.String()
}

func (s *SingleCallByVideoResponseBody) SetCallId(v string) *SingleCallByVideoResponseBody {
	s.CallId = &v
	return s
}

func (s *SingleCallByVideoResponseBody) SetCode(v string) *SingleCallByVideoResponseBody {
	s.Code = &v
	return s
}

func (s *SingleCallByVideoResponseBody) SetMessage(v string) *SingleCallByVideoResponseBody {
	s.Message = &v
	return s
}

func (s *SingleCallByVideoResponseBody) SetRequestId(v string) *SingleCallByVideoResponseBody {
	s.RequestId = &v
	return s
}

type SingleCallByVideoResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SingleCallByVideoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SingleCallByVideoResponse) String() string {
	return tea.Prettify(s)
}

func (s SingleCallByVideoResponse) GoString() string {
	return s.String()
}

func (s *SingleCallByVideoResponse) SetHeaders(v map[string]*string) *SingleCallByVideoResponse {
	s.Headers = v
	return s
}

func (s *SingleCallByVideoResponse) SetStatusCode(v int32) *SingleCallByVideoResponse {
	s.StatusCode = &v
	return s
}

func (s *SingleCallByVideoResponse) SetBody(v *SingleCallByVideoResponseBody) *SingleCallByVideoResponse {
	s.Body = v
	return s
}

type SingleCallByVoiceRequest struct {
	// The number for receiving voice notifications.
	//
	// Number format:
	//
	// *   In the Chinese mainland:
	//
	//     *   Mobile phone number, for example, 159\*\*\*\*0000.
	//     *   Landline number, for example, 0571\*\*\*\*5678.
	//
	// *   Outside the Chinese mainland: country code + phone number, for example, 85200\*\*\*\*00.
	//
	// >
	//
	// *   You can specify only one called number for a request. For more information, see [How to use voice notifications in the Chinese mainland](~~150016~~) or [How to use voice notifications in regions outside the Chinese mainland](~~268810~~).
	//
	// *   Voice notifications are sent to a called number at the following frequency: one time per minute, five times per hour, and 20 times per 24 hours.
	CalledNumber *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	// The number displayed to the called party.
	//
	// *   You do not need to specify this parameter if you use a voice notification file that uses the common outbound call mode. For more information, see [FAQ about the common outbound call mode](~~172104~~).
	// *   If you use a voice notification file that uses the dedicated outbound call mode, you must specify a number that you purchased. You can specify only one number. You can log on to the [Voice Messaging Service console](https://dyvms.console.aliyun.com/overview/home) and choose **Real Number Service** > **Real Number Management** to view the number that you purchased.
	CalledShowNumber *string `json:"CalledShowNumber,omitempty" xml:"CalledShowNumber,omitempty"`
	// The ID reserved for the caller. This ID is returned to the caller in a receipt message.
	//
	// The value must be of the STRING type and 1 to 15 bytes in length.
	OutId   *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of times the voice notification file is played. Valid values: 1 to 3.
	PlayTimes            *int32  `json:"PlayTimes,omitempty" xml:"PlayTimes,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The playback speed of the voice notification file. Valid values: -500 to 500.
	Speed *int32 `json:"Speed,omitempty" xml:"Speed,omitempty"`
	// The voice ID of the voice notification file.
	//
	// You can log on to the [Voice Messaging Service console](https://dyvms.console.aliyun.com/overview/home), choose **Voice Messages** > **Voice Notifications** or **Voice File Management**, and then click the **Voice Notification Files** tab to view the **voice ID**.
	VoiceCode *string `json:"VoiceCode,omitempty" xml:"VoiceCode,omitempty"`
	// The playback volume of the voice notification file. Valid values: 0 to 100. Default value: 100.
	Volume *int32 `json:"Volume,omitempty" xml:"Volume,omitempty"`
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
	// The unique receipt ID for the call.
	//
	// You can call the [QueryCallDetailByCallId](~~393529~~) operation to query the details of the call.
	CallId *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	// The response code.
	//
	// *   The value OK indicates that the request was successful.****
	// *   For more information about other response codes, see [API error codes](~~112502~~).
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SingleCallByVoiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *SingleCallByVoiceResponse) SetStatusCode(v int32) *SingleCallByVoiceResponse {
	s.StatusCode = &v
	return s
}

func (s *SingleCallByVoiceResponse) SetBody(v *SingleCallByVoiceResponseBody) *SingleCallByVoiceResponse {
	s.Body = v
	return s
}

type SkipVideoFileRequest struct {
	CallId               *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	CalledNumber         *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	OutId                *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SkipTimes            *int64  `json:"SkipTimes,omitempty" xml:"SkipTimes,omitempty"`
}

func (s SkipVideoFileRequest) String() string {
	return tea.Prettify(s)
}

func (s SkipVideoFileRequest) GoString() string {
	return s.String()
}

func (s *SkipVideoFileRequest) SetCallId(v string) *SkipVideoFileRequest {
	s.CallId = &v
	return s
}

func (s *SkipVideoFileRequest) SetCalledNumber(v string) *SkipVideoFileRequest {
	s.CalledNumber = &v
	return s
}

func (s *SkipVideoFileRequest) SetOutId(v string) *SkipVideoFileRequest {
	s.OutId = &v
	return s
}

func (s *SkipVideoFileRequest) SetOwnerId(v int64) *SkipVideoFileRequest {
	s.OwnerId = &v
	return s
}

func (s *SkipVideoFileRequest) SetResourceOwnerAccount(v string) *SkipVideoFileRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SkipVideoFileRequest) SetResourceOwnerId(v int64) *SkipVideoFileRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SkipVideoFileRequest) SetSkipTimes(v int64) *SkipVideoFileRequest {
	s.SkipTimes = &v
	return s
}

type SkipVideoFileResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Code               *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data               *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	Message            *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Success            *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SkipVideoFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SkipVideoFileResponseBody) GoString() string {
	return s.String()
}

func (s *SkipVideoFileResponseBody) SetAccessDeniedDetail(v string) *SkipVideoFileResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *SkipVideoFileResponseBody) SetCode(v string) *SkipVideoFileResponseBody {
	s.Code = &v
	return s
}

func (s *SkipVideoFileResponseBody) SetData(v bool) *SkipVideoFileResponseBody {
	s.Data = &v
	return s
}

func (s *SkipVideoFileResponseBody) SetMessage(v string) *SkipVideoFileResponseBody {
	s.Message = &v
	return s
}

func (s *SkipVideoFileResponseBody) SetSuccess(v bool) *SkipVideoFileResponseBody {
	s.Success = &v
	return s
}

type SkipVideoFileResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SkipVideoFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SkipVideoFileResponse) String() string {
	return tea.Prettify(s)
}

func (s SkipVideoFileResponse) GoString() string {
	return s.String()
}

func (s *SkipVideoFileResponse) SetHeaders(v map[string]*string) *SkipVideoFileResponse {
	s.Headers = v
	return s
}

func (s *SkipVideoFileResponse) SetStatusCode(v int32) *SkipVideoFileResponse {
	s.StatusCode = &v
	return s
}

func (s *SkipVideoFileResponse) SetBody(v *SkipVideoFileResponseBody) *SkipVideoFileResponse {
	s.Body = v
	return s
}

type SmartCallRequest struct {
	// Specifies whether the playback of the recording file can be interrupted. Default value: **true**. The default value indicates that the playback of the recording file can be interrupted.
	//
	// If you set the value of this parameter to false, the playback of the recording file cannot be interrupted even if the value of action_break is set to true.
	//
	// > The value of action_code_break takes precedence over the value of action_break.
	ActionCodeBreak *bool `json:"ActionCodeBreak,omitempty" xml:"ActionCodeBreak,omitempty"`
	// The duration that the user keeps speaking. The playback of the recording file is interrupted when this duration is reached. Unit: milliseconds.
	//
	// If the value of ActionCodeBreak is set to **true** for the recording file and the duration that the user keeps speaking reaches the specified duration, the playback of the recording file is interrupted. If you do not specify ActionCodeTimeBreak or set the value of ActionCodeTimeBreak to 0, the setting of ActionCodeBreak does not take effect.
	ActionCodeTimeBreak *int32 `json:"ActionCodeTimeBreak,omitempty" xml:"ActionCodeTimeBreak,omitempty"`
	// The ASR base model. Valid values:
	//
	// *   **customer_service\_8k** (default): Chinese Mandarin.
	// *   **dialect_customer_service\_8k**: a heavy accent.
	//
	// > You must specify the ASR model when you call the SmartCall operation. We recommend that you specify either of the AsrModelId and AsrBaseId parameters.
	//
	// *   If you specify only the AsrModelId parameter, the specified ASR model is used.
	//
	// *   If you specify only the AsrBaseId parameter, the ASR base model is used.
	//
	// *   If you specify neither of the two parameters, the default ASR base model is used, that is, the default value customer_service\_8k is used for the AsrBaseId parameter.
	//
	// *   If you specify both parameters, make sure that their values do not conflict with each other.
	AsrBaseId *string `json:"AsrBaseId,omitempty" xml:"AsrBaseId,omitempty"`
	// The ID of the Automatic Speech Recognition (ASR) model.
	//
	// You can log on to the [Voice Messaging Service console](https://dyvms.console.aliyun.com/overview/home) and view the ID of the ASR model on the **ASR Model Management** page.
	//
	// > You must specify the ASR model when you call the SmartCall operation. We recommend that you specify either of the AsrModelId and AsrBaseId parameters.
	//
	// *   If you specify only the AsrModelId parameter, the specified ASR model is used.
	//
	// *   If you specify only the AsrBaseId parameter, the specified ASR base model is used.
	//
	// *   If you specify neither of the two parameters, the default value customer_service\_8k is used for the AsrBaseId parameter. This means that the default Mandarin ASR base model is used.
	//
	// *   If you specify both parameters, make sure that their values do not conflict with each other.
	AsrModelId *string `json:"AsrModelId,omitempty" xml:"AsrModelId,omitempty"`
	// The ID of the background voice file that is played back when the user talks with the robot.
	//
	// You can log on to the [Voice Messaging Service console](https://dyvms.console.aliyun.com/overview/home), choose **Voice File Management**, click the **Intelligent Speech Interaction Recording File** tab, and then click **Details** in the Actions column to view the voice ID.
	BackgroundFileCode *string `json:"BackgroundFileCode,omitempty" xml:"BackgroundFileCode,omitempty"`
	// This parameter is unavailable.
	BackgroundSpeed *int32 `json:"BackgroundSpeed,omitempty" xml:"BackgroundSpeed,omitempty"`
	// This parameter is unavailable.
	BackgroundVolume *int32 `json:"BackgroundVolume,omitempty" xml:"BackgroundVolume,omitempty"`
	// The called number. Only phone numbers in the Chinese mainland are supported.
	CalledNumber *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	// The number displayed to the called party. The value must be the number you purchased.
	//
	// You can log on to the [Voice Messaging Service console](https://dyvms.console.aliyun.com/overview/home) and choose **Voice Numbers** > **Real Number Management** to view the number you purchased.
	CalledShowNumber *string `json:"CalledShowNumber,omitempty" xml:"CalledShowNumber,omitempty"`
	// The dynamic extension ID that is reserved for the caller of the operation. This ID is returned in the callback URL and is used as the development identifier of the customer.
	DynamicId *string `json:"DynamicId,omitempty" xml:"DynamicId,omitempty"`
	// Specifies whether to enable speech recognition of early media. Valid values:
	//
	// *   **false** (default): Speech recognition of early media is disabled.
	// *   **true**: Speech recognition of early media is enabled.
	//
	// > If you set the value of this parameter to **true**, the reason why the call is not answered is recorded.
	EarlyMediaAsr *bool `json:"EarlyMediaAsr,omitempty" xml:"EarlyMediaAsr,omitempty"`
	// Specifies whether to enable Inverse Text Normalization (ITN) during post-processing. Default value: **false**. If you set the value to false, ITN is not enabled during post-processing.
	//
	// If you set the value to **true**, Chinese numerals are converted into Arabic numerals for output.
	EnableITN *bool `json:"EnableITN,omitempty" xml:"EnableITN,omitempty"`
	// The silence duration. The system determines the end of the conversation based on the silence duration of the user. Unit: milliseconds. Valid values: 1000 to 20000.****
	//
	// >
	//
	// *   If you specify a value out of the valid range, the default value **10000** is used.
	//
	// *   The parameter value can be adjusted during the conversation. The last setting prevails.
	MuteTime       *int32   `json:"MuteTime,omitempty" xml:"MuteTime,omitempty"`
	NoiseThreshold *float64 `json:"NoiseThreshold,omitempty" xml:"NoiseThreshold,omitempty"`
	// The ID that is reserved for the caller of the operation. This ID is returned to the caller in a receipt message.
	//
	// The value is of the STRING type and must be 1 to 15 bytes in length.
	OutId   *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The pause duration. The system determines the end of a sentence based on the pause duration of the user in the conversation. Unit: milliseconds. Valid values: 300 to 1200.****
	//
	// >
	//
	// *   If you specify a value out of the valid range, the default value **800** is used.
	//
	// *   You cannot change the parameter value after specifying it.
	PauseTime *int32 `json:"PauseTime,omitempty" xml:"PauseTime,omitempty"`
	// Specifies whether to record the conversation. Valid values:
	//
	// *   **true**: The conversation is recorded.
	// *   **false**: The conversation is not recorded.
	RecordFlag           *bool   `json:"RecordFlag,omitempty" xml:"RecordFlag,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The maximum call duration. The call is automatically hung up when the maximum call duration is reached. Unit: seconds.
	//
	// > The maximum call duration is 3,600 seconds.
	SessionTimeout *int32 `json:"SessionTimeout,omitempty" xml:"SessionTimeout,omitempty"`
	// This parameter is unavailable.
	Speed *int32 `json:"Speed,omitempty" xml:"Speed,omitempty"`
	// Specifies whether to enable streaming ASR, which intelligently judges what the user wants to express based on the first few words spoken by the user. Valid values:
	//
	// *   **0**: Streaming ASR is disabled.
	// *   **1**: Streaming ASR is enabled.
	StreamAsr *int32 `json:"StreamAsr,omitempty" xml:"StreamAsr,omitempty"`
	// Specifies whether to set TTS sound parameters. Valid values:
	//
	// *   **true**: TTS sound parameters must be set. You must set the **TtsStyle**, **TtsColume**, and **TtsSpeed** parameters to specify a sound style.
	// *   **false**: TTS sound parameters do not need to be set. The values of TTS sound parameters do not take effect even if you set them.
	TtsConf *bool `json:"TtsConf,omitempty" xml:"TtsConf,omitempty"`
	// The speed of TTS variable playback. Valid values: -200 to 200. Default value: 0.
	TtsSpeed *int32 `json:"TtsSpeed,omitempty" xml:"TtsSpeed,omitempty"`
	// The sound style for TTS variable playback. Default value: **xiaoyun**. For more information about the sound styles, see the **Sound styles** table below.
	TtsStyle *string `json:"TtsStyle,omitempty" xml:"TtsStyle,omitempty"`
	// The volume of TTS variable playback. Valid values: 0 to 100. Default value: 0.
	TtsVolume *int32 `json:"TtsVolume,omitempty" xml:"TtsVolume,omitempty"`
	// The recording file that is played back in the intelligent outbound call.
	//
	// The file can be an online file, a voice file uploaded from the Voice Messaging Service console, or a text-to-speech (TTS) template that contains variables. You can specify multiple files and a TTS variable together. Separate them with commas (,). The values of the variables in the TTS template are specified by the **VoiceCodeParam** parameter.
	//
	// *   If you use an online file as the recording file, set the value of **VoiceCode** to the URL of the file that can be accessed over the Internet.
	// *   If you use a voice file uploaded from the Voice Messaging Service console as the recording file, set the value of **VoiceCode** to the voice ID of the file. You can log on to the [Voice Messaging Service console](https://dyvms.console.aliyun.com/overview/home), choose **Voice File Management**, click the **Intelligent Speech Interaction Recording File** tab, and then click **Details** in the Actions column to view the voice ID.
	// *   If you use a TTS template that contains variables as the recording file, set the value of **VoiceCode** to a variable name such as $name$, and also set a value for the variable in the **VoiceCodeParam** parameter.
	VoiceCode *string `json:"VoiceCode,omitempty" xml:"VoiceCode,omitempty"`
	// The value of the TTS variable, in the JSON format. This value must match the TTS variable specified by the **VoiceCode** parameter.
	VoiceCodeParam *string `json:"VoiceCodeParam,omitempty" xml:"VoiceCodeParam,omitempty"`
	// The volume at which the recording file is played. Valid values: -4 to 4. We recommend that you set the value of this parameter to **1**.
	Volume *int32 `json:"Volume,omitempty" xml:"Volume,omitempty"`
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

func (s *SmartCallRequest) SetNoiseThreshold(v float64) *SmartCallRequest {
	s.NoiseThreshold = &v
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
	// The unique receipt ID for this call.
	//
	// You can call the [QueryCallDetailByCallId](~~QueryCallDetailByCallId~~) operation to query the details of the call based on the receipt ID.
	CallId *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	// The response code.
	//
	// *   The value OK indicates that the request was successful.
	// *   For more information about other response codes, see [API error codes](~~112502~~).
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
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
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SmartCallResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *SmartCallResponse) SetStatusCode(v int32) *SmartCallResponse {
	s.StatusCode = &v
	return s
}

func (s *SmartCallResponse) SetBody(v *SmartCallResponseBody) *SmartCallResponse {
	s.Body = v
	return s
}

type SmartCallOperateRequest struct {
	// The unique receipt ID of the call. You can call the [SmartCall](~~393526~~) operation to obtain the receipt ID.
	CallId *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	// The action that is initiated to the called number of an outbound robocall.
	//
	// > Only the value **parallelBridge** is supported. This value indicates that a bridge action is initiated between a called number and an agent of the call center.
	Command *string `json:"Command,omitempty" xml:"Command,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The extended field.
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
	// The response code.
	//
	// *   The value OK indicates that the request was successful.
	// *   For more information about other response codes, see [API error codes](~~112502~~).
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The action result. Valid values:
	//
	// *   **true**: The action was successful.
	// *   **false**: The action failed.
	Status *bool `json:"Status,omitempty" xml:"Status,omitempty"`
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SmartCallOperateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *SmartCallOperateResponse) SetStatusCode(v int32) *SmartCallOperateResponse {
	s.StatusCode = &v
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
	// The time scheduled for starting the robocall task, in the yyyy-MM-dd HH:mm:ss format.
	ScheduleTime *string `json:"ScheduleTime,omitempty" xml:"ScheduleTime,omitempty"`
	// The unique ID of the robocall task. You can call the [CreateRobotTask](~~CreateRobotTask~~) operation to obtain the task ID.
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
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
	// The response code.
	//
	// *   The value OK indicates that the request was successful.
	// *   For more information about other response codes, see [API error codes](~~112502~~).
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// *   **true**: The request was successful.
	// *   **false**: The request failed.
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartRobotTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *StartRobotTaskResponse) SetStatusCode(v int32) *StartRobotTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *StartRobotTaskResponse) SetBody(v *StartRobotTaskResponseBody) *StartRobotTaskResponse {
	s.Body = v
	return s
}

type StopCallInConfigRequest struct {
	// The China 400 number from which the inbound call to be stopped is transferred.
	Number               *string `json:"Number,omitempty" xml:"Number,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s StopCallInConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s StopCallInConfigRequest) GoString() string {
	return s.String()
}

func (s *StopCallInConfigRequest) SetNumber(v string) *StopCallInConfigRequest {
	s.Number = &v
	return s
}

func (s *StopCallInConfigRequest) SetOwnerId(v int64) *StopCallInConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *StopCallInConfigRequest) SetResourceOwnerAccount(v string) *StopCallInConfigRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *StopCallInConfigRequest) SetResourceOwnerId(v int64) *StopCallInConfigRequest {
	s.ResourceOwnerId = &v
	return s
}

type StopCallInConfigResponseBody struct {
	// The response code.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the inbound call was stopped. Valid values:
	//
	// *   true: The inbound call was stopped.
	// *   false: The inbound call failed to be stopped.
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopCallInConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopCallInConfigResponseBody) GoString() string {
	return s.String()
}

func (s *StopCallInConfigResponseBody) SetCode(v string) *StopCallInConfigResponseBody {
	s.Code = &v
	return s
}

func (s *StopCallInConfigResponseBody) SetData(v bool) *StopCallInConfigResponseBody {
	s.Data = &v
	return s
}

func (s *StopCallInConfigResponseBody) SetMessage(v string) *StopCallInConfigResponseBody {
	s.Message = &v
	return s
}

func (s *StopCallInConfigResponseBody) SetRequestId(v string) *StopCallInConfigResponseBody {
	s.RequestId = &v
	return s
}

type StopCallInConfigResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopCallInConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopCallInConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s StopCallInConfigResponse) GoString() string {
	return s.String()
}

func (s *StopCallInConfigResponse) SetHeaders(v map[string]*string) *StopCallInConfigResponse {
	s.Headers = v
	return s
}

func (s *StopCallInConfigResponse) SetStatusCode(v int32) *StopCallInConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *StopCallInConfigResponse) SetBody(v *StopCallInConfigResponseBody) *StopCallInConfigResponse {
	s.Body = v
	return s
}

type StopRobotTaskRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The unique ID of the robocall task. You can call the [CreateRobotTask](~~CreateRobotTask~~) operation to obtain the task ID.
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
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
	// The response code.
	//
	// *   The value OK indicates that the request was successful.
	// *   For more information about other response codes, see [API error codes](~~112502~~).
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// *   **true**: The request was successful.
	// *   **false**: The request failed.
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
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
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopRobotTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *StopRobotTaskResponse) SetStatusCode(v int32) *StopRobotTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *StopRobotTaskResponse) SetBody(v *StopRobotTaskResponseBody) *StopRobotTaskResponse {
	s.Body = v
	return s
}

type SubmitHotlineTransferRegisterRequest struct {
	// The authenticity of the commitment. Valid values:
	//
	// *   **true**: The commitment is authentic.
	// *   **false**: The commitment is not authentic.
	Agreement *string `json:"Agreement,omitempty" xml:"Agreement,omitempty"`
	// The China 400 number.
	HotlineNumber *string `json:"HotlineNumber,omitempty" xml:"HotlineNumber,omitempty"`
	// The ID card number of the handler.
	OperatorIdentityCard *string `json:"OperatorIdentityCard,omitempty" xml:"OperatorIdentityCard,omitempty"`
	// The email address of the handler.
	OperatorMail *string `json:"OperatorMail,omitempty" xml:"OperatorMail,omitempty"`
	// The verification code that is received by the mailbox of the handler.
	OperatorMailVerifyCode *string `json:"OperatorMailVerifyCode,omitempty" xml:"OperatorMailVerifyCode,omitempty"`
	// The mobile phone number of the handler.
	OperatorMobile *string `json:"OperatorMobile,omitempty" xml:"OperatorMobile,omitempty"`
	// The verification code that is received by the mobile phone of the handler.
	OperatorMobileVerifyCode *string `json:"OperatorMobileVerifyCode,omitempty" xml:"OperatorMobileVerifyCode,omitempty"`
	// The name of the handler.
	OperatorName *string `json:"OperatorName,omitempty" xml:"OperatorName,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The qualification ID. You can call the [GetHotlineQualificationByOrder](~~393548~~) operation to obtain the qualification ID.
	QualificationId      *string `json:"QualificationId,omitempty" xml:"QualificationId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The registration information about the China 400 number.
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
	// The ID card number of the number owner.
	IdentityCard *string `json:"IdentityCard,omitempty" xml:"IdentityCard,omitempty"`
	// The China 400 number that you want to submit for registration.
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// The real name or company name of the number owner.
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
	// The response code.
	//
	// *   The value OK indicates that the request was successful.
	// *   For more information about other response codes, see [API error codes](~~112502~~).
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The registration ID.
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
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
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitHotlineTransferRegisterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *SubmitHotlineTransferRegisterResponse) SetStatusCode(v int32) *SubmitHotlineTransferRegisterResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitHotlineTransferRegisterResponse) SetBody(v *SubmitHotlineTransferRegisterResponseBody) *SubmitHotlineTransferRegisterResponse {
	s.Body = v
	return s
}

type UpgradeVideoFileRequest struct {
	CallId               *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	CalledNumber         *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	MediaType            *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	OutId                *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s UpgradeVideoFileRequest) String() string {
	return tea.Prettify(s)
}

func (s UpgradeVideoFileRequest) GoString() string {
	return s.String()
}

func (s *UpgradeVideoFileRequest) SetCallId(v string) *UpgradeVideoFileRequest {
	s.CallId = &v
	return s
}

func (s *UpgradeVideoFileRequest) SetCalledNumber(v string) *UpgradeVideoFileRequest {
	s.CalledNumber = &v
	return s
}

func (s *UpgradeVideoFileRequest) SetMediaType(v string) *UpgradeVideoFileRequest {
	s.MediaType = &v
	return s
}

func (s *UpgradeVideoFileRequest) SetOutId(v string) *UpgradeVideoFileRequest {
	s.OutId = &v
	return s
}

func (s *UpgradeVideoFileRequest) SetOwnerId(v int64) *UpgradeVideoFileRequest {
	s.OwnerId = &v
	return s
}

func (s *UpgradeVideoFileRequest) SetResourceOwnerAccount(v string) *UpgradeVideoFileRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpgradeVideoFileRequest) SetResourceOwnerId(v int64) *UpgradeVideoFileRequest {
	s.ResourceOwnerId = &v
	return s
}

type UpgradeVideoFileResponseBody struct {
	AccessDeniedDetail *string                `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Code               *string                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data               map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	Message            *string                `json:"Message,omitempty" xml:"Message,omitempty"`
	Success            *bool                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpgradeVideoFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpgradeVideoFileResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeVideoFileResponseBody) SetAccessDeniedDetail(v string) *UpgradeVideoFileResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UpgradeVideoFileResponseBody) SetCode(v string) *UpgradeVideoFileResponseBody {
	s.Code = &v
	return s
}

func (s *UpgradeVideoFileResponseBody) SetData(v map[string]interface{}) *UpgradeVideoFileResponseBody {
	s.Data = v
	return s
}

func (s *UpgradeVideoFileResponseBody) SetMessage(v string) *UpgradeVideoFileResponseBody {
	s.Message = &v
	return s
}

func (s *UpgradeVideoFileResponseBody) SetSuccess(v bool) *UpgradeVideoFileResponseBody {
	s.Success = &v
	return s
}

type UpgradeVideoFileResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpgradeVideoFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpgradeVideoFileResponse) String() string {
	return tea.Prettify(s)
}

func (s UpgradeVideoFileResponse) GoString() string {
	return s.String()
}

func (s *UpgradeVideoFileResponse) SetHeaders(v map[string]*string) *UpgradeVideoFileResponse {
	s.Headers = v
	return s
}

func (s *UpgradeVideoFileResponse) SetStatusCode(v int32) *UpgradeVideoFileResponse {
	s.StatusCode = &v
	return s
}

func (s *UpgradeVideoFileResponse) SetBody(v *UpgradeVideoFileResponseBody) *UpgradeVideoFileResponse {
	s.Body = v
	return s
}

type UploadRobotTaskCalledFileRequest struct {
	// The called numbers. Separate multiple called numbers with commas (,).
	//
	// > After you create a robocall task, you must upload called numbers in batches. You can upload up to 300,000 called numbers for each task.
	CalledNumber *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	// The unique ID of the robocall task. You can call the [CreateRobotTask](~~CreateRobotTask~~) operation to obtain the ID of the robocall task.
	Id                   *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The values of the variable in the text-to-speech (TTS) template, in the JSON format. The variable values specified by the TtsParam parameter must match the variable names specified by the TtsParamHead parameter.
	//
	// *   If all the called numbers carry the same variable values, you can set the value of the number field to **all** and upload only one copy of the variable values.
	// *   If only some of the called numbers carry the same variable values, you can set the value of the number field to **all** for these called numbers and set the value of the number field and variable values for other called numbers based on your business requirements. The system preferentially selects the values that you set for the called numbers.
	TtsParam *string `json:"TtsParam,omitempty" xml:"TtsParam,omitempty"`
	// The list of variable names carried in the robocall task, in the JSON format. The TtsParamHead parameter must be used together with the TtsParam parameter.
	TtsParamHead *string `json:"TtsParamHead,omitempty" xml:"TtsParamHead,omitempty"`
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
	// The response code.
	//
	// *   The value OK indicates that the request was successful.
	// *   For more information about other response codes, see [API error codes](~~112502~~).
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The unique ID of the robocall task.
	//
	// You can call the [QueryRobotTaskDetail](~~QueryRobotTaskDetail~~) operation to query the details of the robocall task based on the task ID.
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
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
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UploadRobotTaskCalledFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *UploadRobotTaskCalledFileResponse) SetStatusCode(v int32) *UploadRobotTaskCalledFileResponse {
	s.StatusCode = &v
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

/**
 * ### QPS limits
 * You can call this operation up to 200 times per second per account.
 *
 * @param request AddVirtualNumberRelationRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return AddVirtualNumberRelationResponse
 */
func (client *Client) AddVirtualNumberRelationWithOptions(request *AddVirtualNumberRelationRequest, runtime *util.RuntimeOptions) (_result *AddVirtualNumberRelationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CorpNameList)) {
		query["CorpNameList"] = request.CorpNameList
	}

	if !tea.BoolValue(util.IsUnset(request.NumberList)) {
		query["NumberList"] = request.NumberList
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNum)) {
		query["PhoneNum"] = request.PhoneNum
	}

	if !tea.BoolValue(util.IsUnset(request.ProdCode)) {
		query["ProdCode"] = request.ProdCode
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RouteType)) {
		query["RouteType"] = request.RouteType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddVirtualNumberRelation"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
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

/**
 * ### QPS limits
 * You can call this operation up to 200 times per second per account.
 *
 * @param request AddVirtualNumberRelationRequest
 * @return AddVirtualNumberRelationResponse
 */
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

/**
 * *   In an intelligent speech interaction task, you can use the robot communication scripts preset in the Voice Messaging Service console, or invoke the callback function to return the response mode configured by the business party in each call.
 * *   The BatchRobotSmartCall operation is used to initiate an outbound robocall task by using the robot communication scripts preset in the Voice Messaging Service console.
 * ## Prerequisites
 * *   You have passed the real-name verification for an enterprise user and passed the enterprise qualification review.
 * *   You have purchased numbers in the [Voice Messaging Service console](https://dyvms.console.aliyun.com/dyvms.htm#/number/normal).
 * *   You have added communication scripts on the [Communication script management](https://dyvms.console.aliyun.com/dyvms.htm#/smart-call/saas/robot/list) page, and the communication scripts have been approved.
 * > Before you call this operation, make sure that you are familiar with the [billing](https://www.aliyun.com/price/product#/vms/detail) of Voice Messaging Service (VMS).
 *
 * @param request BatchRobotSmartCallRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return BatchRobotSmartCallResponse
 */
func (client *Client) BatchRobotSmartCallWithOptions(request *BatchRobotSmartCallRequest, runtime *util.RuntimeOptions) (_result *BatchRobotSmartCallResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CalledNumber)) {
		query["CalledNumber"] = request.CalledNumber
	}

	if !tea.BoolValue(util.IsUnset(request.CalledShowNumber)) {
		query["CalledShowNumber"] = request.CalledShowNumber
	}

	if !tea.BoolValue(util.IsUnset(request.CorpName)) {
		query["CorpName"] = request.CorpName
	}

	if !tea.BoolValue(util.IsUnset(request.DialogId)) {
		query["DialogId"] = request.DialogId
	}

	if !tea.BoolValue(util.IsUnset(request.EarlyMediaAsr)) {
		query["EarlyMediaAsr"] = request.EarlyMediaAsr
	}

	if !tea.BoolValue(util.IsUnset(request.IsSelfLine)) {
		query["IsSelfLine"] = request.IsSelfLine
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ScheduleCall)) {
		query["ScheduleCall"] = request.ScheduleCall
	}

	if !tea.BoolValue(util.IsUnset(request.ScheduleTime)) {
		query["ScheduleTime"] = request.ScheduleTime
	}

	if !tea.BoolValue(util.IsUnset(request.TaskName)) {
		query["TaskName"] = request.TaskName
	}

	if !tea.BoolValue(util.IsUnset(request.TtsParam)) {
		query["TtsParam"] = request.TtsParam
	}

	if !tea.BoolValue(util.IsUnset(request.TtsParamHead)) {
		query["TtsParamHead"] = request.TtsParamHead
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchRobotSmartCall"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
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

/**
 * *   In an intelligent speech interaction task, you can use the robot communication scripts preset in the Voice Messaging Service console, or invoke the callback function to return the response mode configured by the business party in each call.
 * *   The BatchRobotSmartCall operation is used to initiate an outbound robocall task by using the robot communication scripts preset in the Voice Messaging Service console.
 * ## Prerequisites
 * *   You have passed the real-name verification for an enterprise user and passed the enterprise qualification review.
 * *   You have purchased numbers in the [Voice Messaging Service console](https://dyvms.console.aliyun.com/dyvms.htm#/number/normal).
 * *   You have added communication scripts on the [Communication script management](https://dyvms.console.aliyun.com/dyvms.htm#/smart-call/saas/robot/list) page, and the communication scripts have been approved.
 * > Before you call this operation, make sure that you are familiar with the [billing](https://www.aliyun.com/price/product#/vms/detail) of Voice Messaging Service (VMS).
 *
 * @param request BatchRobotSmartCallRequest
 * @return BatchRobotSmartCallResponse
 */
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

/**
 * ### QPS limits
 * You can call this operation up to 100 times per second per account.
 *
 * @param request CancelOrderRobotTaskRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CancelOrderRobotTaskResponse
 */
func (client *Client) CancelOrderRobotTaskWithOptions(request *CancelOrderRobotTaskRequest, runtime *util.RuntimeOptions) (_result *CancelOrderRobotTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CancelOrderRobotTask"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
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

/**
 * ### QPS limits
 * You can call this operation up to 100 times per second per account.
 *
 * @param request CancelOrderRobotTaskRequest
 * @return CancelOrderRobotTaskResponse
 */
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

/**
 * Only a task in progress can be terminated by calling the CancelRobotTask operation, and the task cannot be resumed after it is terminated.
 * ### QPS limits
 * You can call this operation up to 100 times per second per account.
 *
 * @param request CancelRobotTaskRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CancelRobotTaskResponse
 */
func (client *Client) CancelRobotTaskWithOptions(request *CancelRobotTaskRequest, runtime *util.RuntimeOptions) (_result *CancelRobotTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CancelRobotTask"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
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

/**
 * Only a task in progress can be terminated by calling the CancelRobotTask operation, and the task cannot be resumed after it is terminated.
 * ### QPS limits
 * You can call this operation up to 100 times per second per account.
 *
 * @param request CancelRobotTaskRequest
 * @return CancelRobotTaskResponse
 */
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

func (client *Client) ChangeMediaTypeWithOptions(request *ChangeMediaTypeRequest, runtime *util.RuntimeOptions) (_result *ChangeMediaTypeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CallId)) {
		query["CallId"] = request.CallId
	}

	if !tea.BoolValue(util.IsUnset(request.CalledNum)) {
		query["CalledNum"] = request.CalledNum
	}

	if !tea.BoolValue(util.IsUnset(request.MediaType)) {
		query["MediaType"] = request.MediaType
	}

	if !tea.BoolValue(util.IsUnset(request.OutId)) {
		query["OutId"] = request.OutId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ChangeMediaType"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ChangeMediaTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ChangeMediaType(request *ChangeMediaTypeRequest) (_result *ChangeMediaTypeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ChangeMediaTypeResponse{}
	_body, _err := client.ChangeMediaTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can create up to 1,000 voice notifications for each task.
 * ### QPS limits
 * You can call this operation up to 100 times per second per account.
 *
 * @param request CreateCallTaskRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateCallTaskResponse
 */
func (client *Client) CreateCallTaskWithOptions(request *CreateCallTaskRequest, runtime *util.RuntimeOptions) (_result *CreateCallTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		query["BizType"] = request.BizType
	}

	if !tea.BoolValue(util.IsUnset(request.Data)) {
		query["Data"] = request.Data
	}

	if !tea.BoolValue(util.IsUnset(request.DataType)) {
		query["DataType"] = request.DataType
	}

	if !tea.BoolValue(util.IsUnset(request.FireTime)) {
		query["FireTime"] = request.FireTime
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Resource)) {
		query["Resource"] = request.Resource
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.ScheduleType)) {
		query["ScheduleType"] = request.ScheduleType
	}

	if !tea.BoolValue(util.IsUnset(request.StopTime)) {
		query["StopTime"] = request.StopTime
	}

	if !tea.BoolValue(util.IsUnset(request.TaskName)) {
		query["TaskName"] = request.TaskName
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateCode)) {
		query["TemplateCode"] = request.TemplateCode
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateName)) {
		query["TemplateName"] = request.TemplateName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateCallTask"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
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

/**
 * You can create up to 1,000 voice notifications for each task.
 * ### QPS limits
 * You can call this operation up to 100 times per second per account.
 *
 * @param request CreateCallTaskRequest
 * @return CreateCallTaskResponse
 */
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

/**
 * You can call this operation to initiate an outbound robocall task by using the robot communication scripts preset in the Voice Messaging Service console. In an intelligent speech interaction task, you can use the robot communication scripts preset in the Voice Messaging Service console, or invoke the callback function to return the response mode configured by the business party in each call.
 * ### QPS limits
 * You can call this operation up to 100 times per second per account.
 *
 * @param request CreateRobotTaskRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateRobotTaskResponse
 */
func (client *Client) CreateRobotTaskWithOptions(request *CreateRobotTaskRequest, runtime *util.RuntimeOptions) (_result *CreateRobotTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Caller)) {
		query["Caller"] = request.Caller
	}

	if !tea.BoolValue(util.IsUnset(request.CorpName)) {
		query["CorpName"] = request.CorpName
	}

	if !tea.BoolValue(util.IsUnset(request.DialogId)) {
		query["DialogId"] = request.DialogId
	}

	if !tea.BoolValue(util.IsUnset(request.IsSelfLine)) {
		query["IsSelfLine"] = request.IsSelfLine
	}

	if !tea.BoolValue(util.IsUnset(request.NumberStatusIdent)) {
		query["NumberStatusIdent"] = request.NumberStatusIdent
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RecallInterval)) {
		query["RecallInterval"] = request.RecallInterval
	}

	if !tea.BoolValue(util.IsUnset(request.RecallStateCodes)) {
		query["RecallStateCodes"] = request.RecallStateCodes
	}

	if !tea.BoolValue(util.IsUnset(request.RecallTimes)) {
		query["RecallTimes"] = request.RecallTimes
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RetryType)) {
		query["RetryType"] = request.RetryType
	}

	if !tea.BoolValue(util.IsUnset(request.TaskName)) {
		query["TaskName"] = request.TaskName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateRobotTask"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
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

/**
 * You can call this operation to initiate an outbound robocall task by using the robot communication scripts preset in the Voice Messaging Service console. In an intelligent speech interaction task, you can use the robot communication scripts preset in the Voice Messaging Service console, or invoke the callback function to return the response mode configured by the business party in each call.
 * ### QPS limits
 * You can call this operation up to 100 times per second per account.
 *
 * @param request CreateRobotTaskRequest
 * @return CreateRobotTaskResponse
 */
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

func (client *Client) DegradeVideoFileWithOptions(request *DegradeVideoFileRequest, runtime *util.RuntimeOptions) (_result *DegradeVideoFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CallId)) {
		query["CallId"] = request.CallId
	}

	if !tea.BoolValue(util.IsUnset(request.CalledNumber)) {
		query["CalledNumber"] = request.CalledNumber
	}

	if !tea.BoolValue(util.IsUnset(request.MediaType)) {
		query["MediaType"] = request.MediaType
	}

	if !tea.BoolValue(util.IsUnset(request.OutId)) {
		query["OutId"] = request.OutId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DegradeVideoFile"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DegradeVideoFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DegradeVideoFile(request *DegradeVideoFileRequest) (_result *DegradeVideoFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DegradeVideoFileResponse{}
	_body, _err := client.DegradeVideoFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call this operation to delete only tasks that are not started, that are completed, and that are terminated.
 * ### QPS limits
 * You can call this operation up to 100 times per second per account.
 *
 * @param request DeleteRobotTaskRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeleteRobotTaskResponse
 */
func (client *Client) DeleteRobotTaskWithOptions(request *DeleteRobotTaskRequest, runtime *util.RuntimeOptions) (_result *DeleteRobotTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteRobotTask"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
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

/**
 * You can call this operation to delete only tasks that are not started, that are completed, and that are terminated.
 * ### QPS limits
 * You can call this operation up to 100 times per second per account.
 *
 * @param request DeleteRobotTaskRequest
 * @return DeleteRobotTaskResponse
 */
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

/**
 * ### QPS limits
 * You can call this operation up to 100 times per second per account.
 *
 * @param request ExecuteCallTaskRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ExecuteCallTaskResponse
 */
func (client *Client) ExecuteCallTaskWithOptions(request *ExecuteCallTaskRequest, runtime *util.RuntimeOptions) (_result *ExecuteCallTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FireTime)) {
		query["FireTime"] = request.FireTime
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ExecuteCallTask"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
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

/**
 * ### QPS limits
 * You can call this operation up to 100 times per second per account.
 *
 * @param request ExecuteCallTaskRequest
 * @return ExecuteCallTaskResponse
 */
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

func (client *Client) GetCallMediaTypeWithOptions(request *GetCallMediaTypeRequest, runtime *util.RuntimeOptions) (_result *GetCallMediaTypeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CallId)) {
		query["CallId"] = request.CallId
	}

	if !tea.BoolValue(util.IsUnset(request.CalledNumber)) {
		query["CalledNumber"] = request.CalledNumber
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetCallMediaType"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCallMediaTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetCallMediaType(request *GetCallMediaTypeRequest) (_result *GetCallMediaTypeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetCallMediaTypeResponse{}
	_body, _err := client.GetCallMediaTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetCallProgressWithOptions(request *GetCallProgressRequest, runtime *util.RuntimeOptions) (_result *GetCallProgressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CallId)) {
		query["CallId"] = request.CallId
	}

	if !tea.BoolValue(util.IsUnset(request.CalledNum)) {
		query["CalledNum"] = request.CalledNum
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetCallProgress"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCallProgressResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetCallProgress(request *GetCallProgressRequest) (_result *GetCallProgressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetCallProgressResponse{}
	_body, _err := client.GetCallProgressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * ### QPS limits
 * You can call this operation up to 100 times per second per account.
 *
 * @param request GetHotlineQualificationByOrderRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return GetHotlineQualificationByOrderResponse
 */
func (client *Client) GetHotlineQualificationByOrderWithOptions(request *GetHotlineQualificationByOrderRequest, runtime *util.RuntimeOptions) (_result *GetHotlineQualificationByOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrderId)) {
		query["OrderId"] = request.OrderId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetHotlineQualificationByOrder"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
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

/**
 * ### QPS limits
 * You can call this operation up to 100 times per second per account.
 *
 * @param request GetHotlineQualificationByOrderRequest
 * @return GetHotlineQualificationByOrderResponse
 */
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

func (client *Client) GetTemporaryFileUrlWithOptions(request *GetTemporaryFileUrlRequest, runtime *util.RuntimeOptions) (_result *GetTemporaryFileUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.VideoId)) {
		query["VideoId"] = request.VideoId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTemporaryFileUrl"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTemporaryFileUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTemporaryFileUrl(request *GetTemporaryFileUrlRequest) (_result *GetTemporaryFileUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTemporaryFileUrlResponse{}
	_body, _err := client.GetTemporaryFileUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * ### QPS limits
 * You can call this operation up to five times per second per account.
 *
 * @param request GetTokenRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return GetTokenResponse
 */
func (client *Client) GetTokenWithOptions(request *GetTokenRequest, runtime *util.RuntimeOptions) (_result *GetTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.TokenType)) {
		query["TokenType"] = request.TokenType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetToken"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
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

/**
 * ### QPS limits
 * You can call this operation up to five times per second per account.
 *
 * @param request GetTokenRequest
 * @return GetTokenResponse
 */
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

func (client *Client) GetVideoFieldUrlWithOptions(request *GetVideoFieldUrlRequest, runtime *util.RuntimeOptions) (_result *GetVideoFieldUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.VideoFile)) {
		query["VideoFile"] = request.VideoFile
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetVideoFieldUrl"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetVideoFieldUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetVideoFieldUrl(request *GetVideoFieldUrlRequest) (_result *GetVideoFieldUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetVideoFieldUrlResponse{}
	_body, _err := client.GetVideoFieldUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * *   Your enterprise qualification is approved. For more information, see [Submit enterprise qualifications](~~149795~~).
 * *   Voice numbers are purchased. For more information, see [Purchase numbers](~~149794~~).
 * *   When the subscriber answers the call, the subscriber hears a voice that instructs the subscriber to press a key as needed. If the [message receipt](~~112503~~) feature is enabled, the Voice Messaging Service (VMS) platform returns the information about the key pressed by the subscriber to the business system. The key information includes the order confirmation, questionnaire survey, and satisfaction survey completed by the subscriber.
 * ## QPS limits
 * You can call this operation up to 100 times per second per account.
 *
 * @param request IvrCallRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return IvrCallResponse
 */
func (client *Client) IvrCallWithOptions(request *IvrCallRequest, runtime *util.RuntimeOptions) (_result *IvrCallResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ByeCode)) {
		query["ByeCode"] = request.ByeCode
	}

	if !tea.BoolValue(util.IsUnset(request.ByeTtsParams)) {
		query["ByeTtsParams"] = request.ByeTtsParams
	}

	if !tea.BoolValue(util.IsUnset(request.CalledNumber)) {
		query["CalledNumber"] = request.CalledNumber
	}

	if !tea.BoolValue(util.IsUnset(request.CalledShowNumber)) {
		query["CalledShowNumber"] = request.CalledShowNumber
	}

	if !tea.BoolValue(util.IsUnset(request.MenuKeyMap)) {
		query["MenuKeyMap"] = request.MenuKeyMap
	}

	if !tea.BoolValue(util.IsUnset(request.OutId)) {
		query["OutId"] = request.OutId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PlayTimes)) {
		query["PlayTimes"] = request.PlayTimes
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.StartCode)) {
		query["StartCode"] = request.StartCode
	}

	if !tea.BoolValue(util.IsUnset(request.StartTtsParams)) {
		query["StartTtsParams"] = request.StartTtsParams
	}

	if !tea.BoolValue(util.IsUnset(request.Timeout)) {
		query["Timeout"] = request.Timeout
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("IvrCall"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
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

/**
 * *   Your enterprise qualification is approved. For more information, see [Submit enterprise qualifications](~~149795~~).
 * *   Voice numbers are purchased. For more information, see [Purchase numbers](~~149794~~).
 * *   When the subscriber answers the call, the subscriber hears a voice that instructs the subscriber to press a key as needed. If the [message receipt](~~112503~~) feature is enabled, the Voice Messaging Service (VMS) platform returns the information about the key pressed by the subscriber to the business system. The key information includes the order confirmation, questionnaire survey, and satisfaction survey completed by the subscriber.
 * ## QPS limits
 * You can call this operation up to 100 times per second per account.
 *
 * @param request IvrCallRequest
 * @return IvrCallResponse
 */
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

/**
 * ### QPS limits
 * You can call this operation up to 100 times per second per account.
 *
 * @param request ListCallTaskRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListCallTaskResponse
 */
func (client *Client) ListCallTaskWithOptions(request *ListCallTaskRequest, runtime *util.RuntimeOptions) (_result *ListCallTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		query["BizType"] = request.BizType
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskName)) {
		query["TaskName"] = request.TaskName
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateName)) {
		query["TemplateName"] = request.TemplateName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListCallTask"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
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

/**
 * ### QPS limits
 * You can call this operation up to 100 times per second per account.
 *
 * @param request ListCallTaskRequest
 * @return ListCallTaskResponse
 */
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

/**
 * ### QPS limits
 * You can call this operation up to 100 times per second per account.
 *
 * @param request ListCallTaskDetailRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListCallTaskDetailResponse
 */
func (client *Client) ListCallTaskDetailWithOptions(request *ListCallTaskDetailRequest, runtime *util.RuntimeOptions) (_result *ListCallTaskDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CalledNum)) {
		query["CalledNum"] = request.CalledNum
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListCallTaskDetail"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
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

/**
 * ### QPS limits
 * You can call this operation up to 100 times per second per account.
 *
 * @param request ListCallTaskDetailRequest
 * @return ListCallTaskDetailResponse
 */
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

/**
 * ### QPS limits
 * You can call this operation up to 100 times per second per account.
 *
 * @param request ListHotlineTransferNumberRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListHotlineTransferNumberResponse
 */
func (client *Client) ListHotlineTransferNumberWithOptions(request *ListHotlineTransferNumberRequest, runtime *util.RuntimeOptions) (_result *ListHotlineTransferNumberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.HotlineNumber)) {
		query["HotlineNumber"] = request.HotlineNumber
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.QualificationId)) {
		query["QualificationId"] = request.QualificationId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListHotlineTransferNumber"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
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

/**
 * ### QPS limits
 * You can call this operation up to 100 times per second per account.
 *
 * @param request ListHotlineTransferNumberRequest
 * @return ListHotlineTransferNumberResponse
 */
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

/**
 * ### QPS limits
 * You can call this operation up to 100 times per second per account.
 *
 * @param request ListHotlineTransferRegisterFileRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListHotlineTransferRegisterFileResponse
 */
func (client *Client) ListHotlineTransferRegisterFileWithOptions(request *ListHotlineTransferRegisterFileRequest, runtime *util.RuntimeOptions) (_result *ListHotlineTransferRegisterFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.HotlineNumber)) {
		query["HotlineNumber"] = request.HotlineNumber
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.QualificationId)) {
		query["QualificationId"] = request.QualificationId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListHotlineTransferRegisterFile"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
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

/**
 * ### QPS limits
 * You can call this operation up to 100 times per second per account.
 *
 * @param request ListHotlineTransferRegisterFileRequest
 * @return ListHotlineTransferRegisterFileResponse
 */
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

func (client *Client) PauseVideoFileWithOptions(request *PauseVideoFileRequest, runtime *util.RuntimeOptions) (_result *PauseVideoFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CallId)) {
		query["CallId"] = request.CallId
	}

	if !tea.BoolValue(util.IsUnset(request.CalledNumber)) {
		query["CalledNumber"] = request.CalledNumber
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PauseVideoFile"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PauseVideoFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PauseVideoFile(request *PauseVideoFileRequest) (_result *PauseVideoFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PauseVideoFileResponse{}
	_body, _err := client.PauseVideoFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PlayVideoFileWithOptions(request *PlayVideoFileRequest, runtime *util.RuntimeOptions) (_result *PlayVideoFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CallId)) {
		query["CallId"] = request.CallId
	}

	if !tea.BoolValue(util.IsUnset(request.CalledNumber)) {
		query["CalledNumber"] = request.CalledNumber
	}

	if !tea.BoolValue(util.IsUnset(request.OutId)) {
		query["OutId"] = request.OutId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.VideoId)) {
		query["VideoId"] = request.VideoId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PlayVideoFile"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PlayVideoFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PlayVideoFile(request *PlayVideoFileRequest) (_result *PlayVideoFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PlayVideoFileResponse{}
	_body, _err := client.PlayVideoFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * QueryCallDetailByCallId is a common query operation. You can call this operation to query the details of a voice notification, voice verification code, interactive voice response (IVR), intelligent inbound voice call, intelligent outbound voice call, or intelligent robocall.
 * ### QPS limits
 * You can call this operation up to 100 times per second per account.
 *
 * @param request QueryCallDetailByCallIdRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return QueryCallDetailByCallIdResponse
 */
func (client *Client) QueryCallDetailByCallIdWithOptions(request *QueryCallDetailByCallIdRequest, runtime *util.RuntimeOptions) (_result *QueryCallDetailByCallIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CallId)) {
		query["CallId"] = request.CallId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ProdId)) {
		query["ProdId"] = request.ProdId
	}

	if !tea.BoolValue(util.IsUnset(request.QueryDate)) {
		query["QueryDate"] = request.QueryDate
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryCallDetailByCallId"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
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

/**
 * QueryCallDetailByCallId is a common query operation. You can call this operation to query the details of a voice notification, voice verification code, interactive voice response (IVR), intelligent inbound voice call, intelligent outbound voice call, or intelligent robocall.
 * ### QPS limits
 * You can call this operation up to 100 times per second per account.
 *
 * @param request QueryCallDetailByCallIdRequest
 * @return QueryCallDetailByCallIdResponse
 */
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
	if !tea.BoolValue(util.IsUnset(request.Callee)) {
		query["Callee"] = request.Callee
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.QueryDate)) {
		query["QueryDate"] = request.QueryDate
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryCallDetailByTaskId"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
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

/**
 * ### QPS limits
 * You can call this operation up to 100 times per second per account.
 *
 * @param request QueryCallInPoolTransferConfigRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return QueryCallInPoolTransferConfigResponse
 */
func (client *Client) QueryCallInPoolTransferConfigWithOptions(request *QueryCallInPoolTransferConfigRequest, runtime *util.RuntimeOptions) (_result *QueryCallInPoolTransferConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNumber)) {
		query["PhoneNumber"] = request.PhoneNumber
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryCallInPoolTransferConfig"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
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

/**
 * ### QPS limits
 * You can call this operation up to 100 times per second per account.
 *
 * @param request QueryCallInPoolTransferConfigRequest
 * @return QueryCallInPoolTransferConfigResponse
 */
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

/**
 * ### QPS limits
 * You can call this operation up to 100 times per second per account.
 *
 * @param request QueryCallInTransferRecordRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return QueryCallInTransferRecordResponse
 */
func (client *Client) QueryCallInTransferRecordWithOptions(request *QueryCallInTransferRecordRequest, runtime *util.RuntimeOptions) (_result *QueryCallInTransferRecordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CallInCaller)) {
		query["CallInCaller"] = request.CallInCaller
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNumber)) {
		query["PhoneNumber"] = request.PhoneNumber
	}

	if !tea.BoolValue(util.IsUnset(request.QueryDate)) {
		query["QueryDate"] = request.QueryDate
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryCallInTransferRecord"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
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

/**
 * ### QPS limits
 * You can call this operation up to 100 times per second per account.
 *
 * @param request QueryCallInTransferRecordRequest
 * @return QueryCallInTransferRecordResponse
 */
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
	if !tea.BoolValue(util.IsUnset(request.AuditStatus)) {
		query["AuditStatus"] = request.AuditStatus
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryRobotInfoList"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
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

/**
 * ### QPS limits
 * You can call this operation up to 100 times per second per account.
 *
 * @param request QueryRobotTaskCallDetailRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return QueryRobotTaskCallDetailResponse
 */
func (client *Client) QueryRobotTaskCallDetailWithOptions(request *QueryRobotTaskCallDetailRequest, runtime *util.RuntimeOptions) (_result *QueryRobotTaskCallDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Callee)) {
		query["Callee"] = request.Callee
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.QueryDate)) {
		query["QueryDate"] = request.QueryDate
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryRobotTaskCallDetail"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
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

/**
 * ### QPS limits
 * You can call this operation up to 100 times per second per account.
 *
 * @param request QueryRobotTaskCallDetailRequest
 * @return QueryRobotTaskCallDetailResponse
 */
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

/**
 * ### QPS limits
 * You can call this operation up to 100 times per second per account.
 *
 * @param request QueryRobotTaskCallListRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return QueryRobotTaskCallListResponse
 */
func (client *Client) QueryRobotTaskCallListWithOptions(request *QueryRobotTaskCallListRequest, runtime *util.RuntimeOptions) (_result *QueryRobotTaskCallListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CallResult)) {
		query["CallResult"] = request.CallResult
	}

	if !tea.BoolValue(util.IsUnset(request.Called)) {
		query["Called"] = request.Called
	}

	if !tea.BoolValue(util.IsUnset(request.DialogCountFrom)) {
		query["DialogCountFrom"] = request.DialogCountFrom
	}

	if !tea.BoolValue(util.IsUnset(request.DialogCountTo)) {
		query["DialogCountTo"] = request.DialogCountTo
	}

	if !tea.BoolValue(util.IsUnset(request.DurationFrom)) {
		query["DurationFrom"] = request.DurationFrom
	}

	if !tea.BoolValue(util.IsUnset(request.DurationTo)) {
		query["DurationTo"] = request.DurationTo
	}

	if !tea.BoolValue(util.IsUnset(request.HangupDirection)) {
		query["HangupDirection"] = request.HangupDirection
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryRobotTaskCallList"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
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

/**
 * ### QPS limits
 * You can call this operation up to 100 times per second per account.
 *
 * @param request QueryRobotTaskCallListRequest
 * @return QueryRobotTaskCallListResponse
 */
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

/**
 * ### QPS limits
 * You can call this operation up to 100 times per second per account.
 *
 * @param request QueryRobotTaskDetailRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return QueryRobotTaskDetailResponse
 */
func (client *Client) QueryRobotTaskDetailWithOptions(request *QueryRobotTaskDetailRequest, runtime *util.RuntimeOptions) (_result *QueryRobotTaskDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryRobotTaskDetail"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
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

/**
 * ### QPS limits
 * You can call this operation up to 100 times per second per account.
 *
 * @param request QueryRobotTaskDetailRequest
 * @return QueryRobotTaskDetailResponse
 */
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

/**
 * ### QPS limits
 * You can call this operation up to 100 times per second per account.
 *
 * @param request QueryRobotTaskListRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return QueryRobotTaskListResponse
 */
func (client *Client) QueryRobotTaskListWithOptions(request *QueryRobotTaskListRequest, runtime *util.RuntimeOptions) (_result *QueryRobotTaskListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.TaskName)) {
		query["TaskName"] = request.TaskName
	}

	if !tea.BoolValue(util.IsUnset(request.Time)) {
		query["Time"] = request.Time
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryRobotTaskList"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
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

/**
 * ### QPS limits
 * You can call this operation up to 100 times per second per account.
 *
 * @param request QueryRobotTaskListRequest
 * @return QueryRobotTaskListResponse
 */
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

/**
 * ### QPS limits
 * You can call this operation up to 100 times per second per account.
 *
 * @param request QueryRobotv2AllListRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return QueryRobotv2AllListResponse
 */
func (client *Client) QueryRobotv2AllListWithOptions(request *QueryRobotv2AllListRequest, runtime *util.RuntimeOptions) (_result *QueryRobotv2AllListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryRobotv2AllList"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
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

/**
 * ### QPS limits
 * You can call this operation up to 100 times per second per account.
 *
 * @param request QueryRobotv2AllListRequest
 * @return QueryRobotv2AllListResponse
 */
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

func (client *Client) QueryVideoPlayProgressWithOptions(request *QueryVideoPlayProgressRequest, runtime *util.RuntimeOptions) (_result *QueryVideoPlayProgressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CallId)) {
		query["CallId"] = request.CallId
	}

	if !tea.BoolValue(util.IsUnset(request.CalledNumber)) {
		query["CalledNumber"] = request.CalledNumber
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryVideoPlayProgress"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryVideoPlayProgressResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryVideoPlayProgress(request *QueryVideoPlayProgressRequest) (_result *QueryVideoPlayProgressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryVideoPlayProgressResponse{}
	_body, _err := client.QueryVideoPlayProgressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * ### QPS limits
 * You can call this operation up to 100 times per second per account.
 *
 * @param request QueryVirtualNumberRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return QueryVirtualNumberResponse
 */
func (client *Client) QueryVirtualNumberWithOptions(request *QueryVirtualNumberRequest, runtime *util.RuntimeOptions) (_result *QueryVirtualNumberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProdCode)) {
		query["ProdCode"] = request.ProdCode
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RouteType)) {
		query["RouteType"] = request.RouteType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryVirtualNumber"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
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

/**
 * ### QPS limits
 * You can call this operation up to 100 times per second per account.
 *
 * @param request QueryVirtualNumberRequest
 * @return QueryVirtualNumberResponse
 */
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

/**
 * ### QPS limits
 * You can call this operation up to 200 times per second per account.
 *
 * @param request QueryVirtualNumberRelationRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return QueryVirtualNumberRelationResponse
 */
func (client *Client) QueryVirtualNumberRelationWithOptions(request *QueryVirtualNumberRelationRequest, runtime *util.RuntimeOptions) (_result *QueryVirtualNumberRelationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNum)) {
		query["PhoneNum"] = request.PhoneNum
	}

	if !tea.BoolValue(util.IsUnset(request.ProdCode)) {
		query["ProdCode"] = request.ProdCode
	}

	if !tea.BoolValue(util.IsUnset(request.QualificationId)) {
		query["QualificationId"] = request.QualificationId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionNameCity)) {
		query["RegionNameCity"] = request.RegionNameCity
	}

	if !tea.BoolValue(util.IsUnset(request.RelatedNum)) {
		query["RelatedNum"] = request.RelatedNum
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RouteType)) {
		query["RouteType"] = request.RouteType
	}

	if !tea.BoolValue(util.IsUnset(request.SpecId)) {
		query["SpecId"] = request.SpecId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryVirtualNumberRelation"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
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

/**
 * ### QPS limits
 * You can call this operation up to 200 times per second per account.
 *
 * @param request QueryVirtualNumberRelationRequest
 * @return QueryVirtualNumberRelationResponse
 */
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

func (client *Client) QueryVoiceFileAuditInfoWithOptions(request *QueryVoiceFileAuditInfoRequest, runtime *util.RuntimeOptions) (_result *QueryVoiceFileAuditInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BusinessType)) {
		query["BusinessType"] = request.BusinessType
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.VoiceCodes)) {
		query["VoiceCodes"] = request.VoiceCodes
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryVoiceFileAuditInfo"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryVoiceFileAuditInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryVoiceFileAuditInfo(request *QueryVoiceFileAuditInfoRequest) (_result *QueryVoiceFileAuditInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryVoiceFileAuditInfoResponse{}
	_body, _err := client.QueryVoiceFileAuditInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecoverCallInConfigWithOptions(request *RecoverCallInConfigRequest, runtime *util.RuntimeOptions) (_result *RecoverCallInConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Number)) {
		query["Number"] = request.Number
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RecoverCallInConfig"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecoverCallInConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecoverCallInConfig(request *RecoverCallInConfigRequest) (_result *RecoverCallInConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecoverCallInConfigResponse{}
	_body, _err := client.RecoverCallInConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ResumeVideoFileWithOptions(request *ResumeVideoFileRequest, runtime *util.RuntimeOptions) (_result *ResumeVideoFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CallId)) {
		query["CallId"] = request.CallId
	}

	if !tea.BoolValue(util.IsUnset(request.CalledNumber)) {
		query["CalledNumber"] = request.CalledNumber
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ResumeVideoFile"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ResumeVideoFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ResumeVideoFile(request *ResumeVideoFileRequest) (_result *ResumeVideoFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResumeVideoFileResponse{}
	_body, _err := client.ResumeVideoFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SeekVideoFileWithOptions(request *SeekVideoFileRequest, runtime *util.RuntimeOptions) (_result *SeekVideoFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CallId)) {
		query["CallId"] = request.CallId
	}

	if !tea.BoolValue(util.IsUnset(request.CalledNumber)) {
		query["CalledNumber"] = request.CalledNumber
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SeekTimes)) {
		query["SeekTimes"] = request.SeekTimes
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SeekVideoFile"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SeekVideoFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SeekVideoFile(request *SeekVideoFileRequest) (_result *SeekVideoFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SeekVideoFileResponse{}
	_body, _err := client.SeekVideoFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * ### QPS limits
 * You can call this operation up to 100 times per second per account.
 *
 * @param request SendVerificationRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return SendVerificationResponse
 */
func (client *Client) SendVerificationWithOptions(request *SendVerificationRequest, runtime *util.RuntimeOptions) (_result *SendVerificationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		query["BizType"] = request.BizType
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Target)) {
		query["Target"] = request.Target
	}

	if !tea.BoolValue(util.IsUnset(request.VerifyType)) {
		query["VerifyType"] = request.VerifyType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SendVerification"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
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

/**
 * ### QPS limits
 * You can call this operation up to 100 times per second per account.
 *
 * @param request SendVerificationRequest
 * @return SendVerificationResponse
 */
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

/**
 * ### QPS limits
 * You can call this operation up to 100 times per second per account.
 *
 * @param request SetTransferCalleePoolConfigRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return SetTransferCalleePoolConfigResponse
 */
func (client *Client) SetTransferCalleePoolConfigWithOptions(request *SetTransferCalleePoolConfigRequest, runtime *util.RuntimeOptions) (_result *SetTransferCalleePoolConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CalledRouteMode)) {
		query["CalledRouteMode"] = request.CalledRouteMode
	}

	if !tea.BoolValue(util.IsUnset(request.Details)) {
		query["Details"] = request.Details
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PhoneNumber)) {
		query["PhoneNumber"] = request.PhoneNumber
	}

	if !tea.BoolValue(util.IsUnset(request.QualificationId)) {
		query["QualificationId"] = request.QualificationId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetTransferCalleePoolConfig"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
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

/**
 * ### QPS limits
 * You can call this operation up to 100 times per second per account.
 *
 * @param request SetTransferCalleePoolConfigRequest
 * @return SetTransferCalleePoolConfigResponse
 */
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

/**
 * *   Due to business adjustments, the updates of the voice notification and voice verification code services have been stopped in regions outside the Chinese mainland and the services have been discontinued since March 2022. Only qualified customers can continue using the voice notification and voice verification code services.
 * *   For more information about voice plans or voice service billing, see [Pricing of VMS on China site (aliyun.com)](~~150083~~).
 * ### QPS limits
 * You can call this operation up to 1,000 times per second per account.
 *
 * @param request SingleCallByTtsRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return SingleCallByTtsResponse
 */
func (client *Client) SingleCallByTtsWithOptions(request *SingleCallByTtsRequest, runtime *util.RuntimeOptions) (_result *SingleCallByTtsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CalledNumber)) {
		query["CalledNumber"] = request.CalledNumber
	}

	if !tea.BoolValue(util.IsUnset(request.CalledShowNumber)) {
		query["CalledShowNumber"] = request.CalledShowNumber
	}

	if !tea.BoolValue(util.IsUnset(request.OutId)) {
		query["OutId"] = request.OutId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PlayTimes)) {
		query["PlayTimes"] = request.PlayTimes
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Speed)) {
		query["Speed"] = request.Speed
	}

	if !tea.BoolValue(util.IsUnset(request.TtsCode)) {
		query["TtsCode"] = request.TtsCode
	}

	if !tea.BoolValue(util.IsUnset(request.TtsParam)) {
		query["TtsParam"] = request.TtsParam
	}

	if !tea.BoolValue(util.IsUnset(request.Volume)) {
		query["Volume"] = request.Volume
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SingleCallByTts"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
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

/**
 * *   Due to business adjustments, the updates of the voice notification and voice verification code services have been stopped in regions outside the Chinese mainland and the services have been discontinued since March 2022. Only qualified customers can continue using the voice notification and voice verification code services.
 * *   For more information about voice plans or voice service billing, see [Pricing of VMS on China site (aliyun.com)](~~150083~~).
 * ### QPS limits
 * You can call this operation up to 1,000 times per second per account.
 *
 * @param request SingleCallByTtsRequest
 * @return SingleCallByTtsResponse
 */
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

func (client *Client) SingleCallByVideoWithOptions(request *SingleCallByVideoRequest, runtime *util.RuntimeOptions) (_result *SingleCallByVideoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CalledNumber)) {
		query["CalledNumber"] = request.CalledNumber
	}

	if !tea.BoolValue(util.IsUnset(request.CalledShowNumber)) {
		query["CalledShowNumber"] = request.CalledShowNumber
	}

	if !tea.BoolValue(util.IsUnset(request.OutId)) {
		query["OutId"] = request.OutId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PlayTimes)) {
		query["PlayTimes"] = request.PlayTimes
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Speed)) {
		query["Speed"] = request.Speed
	}

	if !tea.BoolValue(util.IsUnset(request.VideoCode)) {
		query["VideoCode"] = request.VideoCode
	}

	if !tea.BoolValue(util.IsUnset(request.VoiceCode)) {
		query["VoiceCode"] = request.VoiceCode
	}

	if !tea.BoolValue(util.IsUnset(request.Volume)) {
		query["Volume"] = request.Volume
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SingleCallByVideo"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SingleCallByVideoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SingleCallByVideo(request *SingleCallByVideoRequest) (_result *SingleCallByVideoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SingleCallByVideoResponse{}
	_body, _err := client.SingleCallByVideoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * > Due to business adjustments, the updates of the voice notification and voice verification code services have been stopped in regions outside the Chinese mainland and the services have been discontinued since March 2022. Only qualified customers can continue using the voice notification and voice verification code services.
 * You can call the [SingleCallByTts](~~393519~~) operation to send voice notifications with variables.
 * ### QPS limits
 * You can call this operation up to 1,200 times per second per account.
 *
 * @param request SingleCallByVoiceRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return SingleCallByVoiceResponse
 */
func (client *Client) SingleCallByVoiceWithOptions(request *SingleCallByVoiceRequest, runtime *util.RuntimeOptions) (_result *SingleCallByVoiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CalledNumber)) {
		query["CalledNumber"] = request.CalledNumber
	}

	if !tea.BoolValue(util.IsUnset(request.CalledShowNumber)) {
		query["CalledShowNumber"] = request.CalledShowNumber
	}

	if !tea.BoolValue(util.IsUnset(request.OutId)) {
		query["OutId"] = request.OutId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PlayTimes)) {
		query["PlayTimes"] = request.PlayTimes
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Speed)) {
		query["Speed"] = request.Speed
	}

	if !tea.BoolValue(util.IsUnset(request.VoiceCode)) {
		query["VoiceCode"] = request.VoiceCode
	}

	if !tea.BoolValue(util.IsUnset(request.Volume)) {
		query["Volume"] = request.Volume
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SingleCallByVoice"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
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

/**
 * > Due to business adjustments, the updates of the voice notification and voice verification code services have been stopped in regions outside the Chinese mainland and the services have been discontinued since March 2022. Only qualified customers can continue using the voice notification and voice verification code services.
 * You can call the [SingleCallByTts](~~393519~~) operation to send voice notifications with variables.
 * ### QPS limits
 * You can call this operation up to 1,200 times per second per account.
 *
 * @param request SingleCallByVoiceRequest
 * @return SingleCallByVoiceResponse
 */
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

func (client *Client) SkipVideoFileWithOptions(request *SkipVideoFileRequest, runtime *util.RuntimeOptions) (_result *SkipVideoFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CallId)) {
		query["CallId"] = request.CallId
	}

	if !tea.BoolValue(util.IsUnset(request.CalledNumber)) {
		query["CalledNumber"] = request.CalledNumber
	}

	if !tea.BoolValue(util.IsUnset(request.OutId)) {
		query["OutId"] = request.OutId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SkipTimes)) {
		query["SkipTimes"] = request.SkipTimes
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SkipVideoFile"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SkipVideoFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SkipVideoFile(request *SkipVideoFileRequest) (_result *SkipVideoFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SkipVideoFileResponse{}
	_body, _err := client.SkipVideoFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * *   The SmartCall operation must be used together with the [intelligent outbound HTTP operation](~~112703~~). After the call initiated by the Voice Messaging Service (VMS) platform is connected, the VMS platform sends the text converted from speech back to the business side, and the business side then returns the follow-up action to the VMS platform.
 * *   The SmartCall operation does not support the following characters: `@ = : "" $ { } ^ * ￥`.
 * ### QPS limits
 * You can call this operation up to 1,000 times per second per account.
 *
 * @param request SmartCallRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return SmartCallResponse
 */
func (client *Client) SmartCallWithOptions(request *SmartCallRequest, runtime *util.RuntimeOptions) (_result *SmartCallResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ActionCodeBreak)) {
		query["ActionCodeBreak"] = request.ActionCodeBreak
	}

	if !tea.BoolValue(util.IsUnset(request.ActionCodeTimeBreak)) {
		query["ActionCodeTimeBreak"] = request.ActionCodeTimeBreak
	}

	if !tea.BoolValue(util.IsUnset(request.AsrBaseId)) {
		query["AsrBaseId"] = request.AsrBaseId
	}

	if !tea.BoolValue(util.IsUnset(request.AsrModelId)) {
		query["AsrModelId"] = request.AsrModelId
	}

	if !tea.BoolValue(util.IsUnset(request.BackgroundFileCode)) {
		query["BackgroundFileCode"] = request.BackgroundFileCode
	}

	if !tea.BoolValue(util.IsUnset(request.BackgroundSpeed)) {
		query["BackgroundSpeed"] = request.BackgroundSpeed
	}

	if !tea.BoolValue(util.IsUnset(request.BackgroundVolume)) {
		query["BackgroundVolume"] = request.BackgroundVolume
	}

	if !tea.BoolValue(util.IsUnset(request.CalledNumber)) {
		query["CalledNumber"] = request.CalledNumber
	}

	if !tea.BoolValue(util.IsUnset(request.CalledShowNumber)) {
		query["CalledShowNumber"] = request.CalledShowNumber
	}

	if !tea.BoolValue(util.IsUnset(request.DynamicId)) {
		query["DynamicId"] = request.DynamicId
	}

	if !tea.BoolValue(util.IsUnset(request.EarlyMediaAsr)) {
		query["EarlyMediaAsr"] = request.EarlyMediaAsr
	}

	if !tea.BoolValue(util.IsUnset(request.EnableITN)) {
		query["EnableITN"] = request.EnableITN
	}

	if !tea.BoolValue(util.IsUnset(request.MuteTime)) {
		query["MuteTime"] = request.MuteTime
	}

	if !tea.BoolValue(util.IsUnset(request.NoiseThreshold)) {
		query["NoiseThreshold"] = request.NoiseThreshold
	}

	if !tea.BoolValue(util.IsUnset(request.OutId)) {
		query["OutId"] = request.OutId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PauseTime)) {
		query["PauseTime"] = request.PauseTime
	}

	if !tea.BoolValue(util.IsUnset(request.RecordFlag)) {
		query["RecordFlag"] = request.RecordFlag
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SessionTimeout)) {
		query["SessionTimeout"] = request.SessionTimeout
	}

	if !tea.BoolValue(util.IsUnset(request.Speed)) {
		query["Speed"] = request.Speed
	}

	if !tea.BoolValue(util.IsUnset(request.StreamAsr)) {
		query["StreamAsr"] = request.StreamAsr
	}

	if !tea.BoolValue(util.IsUnset(request.TtsConf)) {
		query["TtsConf"] = request.TtsConf
	}

	if !tea.BoolValue(util.IsUnset(request.TtsSpeed)) {
		query["TtsSpeed"] = request.TtsSpeed
	}

	if !tea.BoolValue(util.IsUnset(request.TtsStyle)) {
		query["TtsStyle"] = request.TtsStyle
	}

	if !tea.BoolValue(util.IsUnset(request.TtsVolume)) {
		query["TtsVolume"] = request.TtsVolume
	}

	if !tea.BoolValue(util.IsUnset(request.VoiceCode)) {
		query["VoiceCode"] = request.VoiceCode
	}

	if !tea.BoolValue(util.IsUnset(request.VoiceCodeParam)) {
		query["VoiceCodeParam"] = request.VoiceCodeParam
	}

	if !tea.BoolValue(util.IsUnset(request.Volume)) {
		query["Volume"] = request.Volume
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SmartCall"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
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

/**
 * *   The SmartCall operation must be used together with the [intelligent outbound HTTP operation](~~112703~~). After the call initiated by the Voice Messaging Service (VMS) platform is connected, the VMS platform sends the text converted from speech back to the business side, and the business side then returns the follow-up action to the VMS platform.
 * *   The SmartCall operation does not support the following characters: `@ = : "" $ { } ^ * ￥`.
 * ### QPS limits
 * You can call this operation up to 1,000 times per second per account.
 *
 * @param request SmartCallRequest
 * @return SmartCallResponse
 */
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

/**
 * You can call this operation to initiate a specified action on the called number of an outbound robocall when the call is transferred to an agent of the call center.
 * > You can only initiate the action of bridging a called number and an agent of the call center.
 * ### QPS limits
 * You can call this operation up to 100 times per second per account.
 *
 * @param request SmartCallOperateRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return SmartCallOperateResponse
 */
func (client *Client) SmartCallOperateWithOptions(request *SmartCallOperateRequest, runtime *util.RuntimeOptions) (_result *SmartCallOperateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CallId)) {
		query["CallId"] = request.CallId
	}

	if !tea.BoolValue(util.IsUnset(request.Command)) {
		query["Command"] = request.Command
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Param)) {
		query["Param"] = request.Param
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SmartCallOperate"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
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

/**
 * You can call this operation to initiate a specified action on the called number of an outbound robocall when the call is transferred to an agent of the call center.
 * > You can only initiate the action of bridging a called number and an agent of the call center.
 * ### QPS limits
 * You can call this operation up to 100 times per second per account.
 *
 * @param request SmartCallOperateRequest
 * @return SmartCallOperateResponse
 */
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

/**
 * ### QPS limits
 * You can call this operation up to 100 times per second per account.
 *
 * @param request StartRobotTaskRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return StartRobotTaskResponse
 */
func (client *Client) StartRobotTaskWithOptions(request *StartRobotTaskRequest, runtime *util.RuntimeOptions) (_result *StartRobotTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ScheduleTime)) {
		query["ScheduleTime"] = request.ScheduleTime
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StartRobotTask"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
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

/**
 * ### QPS limits
 * You can call this operation up to 100 times per second per account.
 *
 * @param request StartRobotTaskRequest
 * @return StartRobotTaskResponse
 */
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

func (client *Client) StopCallInConfigWithOptions(request *StopCallInConfigRequest, runtime *util.RuntimeOptions) (_result *StopCallInConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Number)) {
		query["Number"] = request.Number
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StopCallInConfig"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StopCallInConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopCallInConfig(request *StopCallInConfigRequest) (_result *StopCallInConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopCallInConfigResponse{}
	_body, _err := client.StopCallInConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * After you stop a robocall task, you can call the [StartRobotTask](~~StartRobotTask~~) operation to start it again.
 * ### QPS limits
 * You can call this operation up to 100 times per second per account.
 *
 * @param request StopRobotTaskRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return StopRobotTaskResponse
 */
func (client *Client) StopRobotTaskWithOptions(request *StopRobotTaskRequest, runtime *util.RuntimeOptions) (_result *StopRobotTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StopRobotTask"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
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

/**
 * After you stop a robocall task, you can call the [StartRobotTask](~~StartRobotTask~~) operation to start it again.
 * ### QPS limits
 * You can call this operation up to 100 times per second per account.
 *
 * @param request StopRobotTaskRequest
 * @return StopRobotTaskResponse
 */
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

/**
 * ### QPS limits
 * You can call this operation up to 100 times per second per account.
 *
 * @param request SubmitHotlineTransferRegisterRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return SubmitHotlineTransferRegisterResponse
 */
func (client *Client) SubmitHotlineTransferRegisterWithOptions(request *SubmitHotlineTransferRegisterRequest, runtime *util.RuntimeOptions) (_result *SubmitHotlineTransferRegisterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Agreement)) {
		query["Agreement"] = request.Agreement
	}

	if !tea.BoolValue(util.IsUnset(request.HotlineNumber)) {
		query["HotlineNumber"] = request.HotlineNumber
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorIdentityCard)) {
		query["OperatorIdentityCard"] = request.OperatorIdentityCard
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorMail)) {
		query["OperatorMail"] = request.OperatorMail
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorMailVerifyCode)) {
		query["OperatorMailVerifyCode"] = request.OperatorMailVerifyCode
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorMobile)) {
		query["OperatorMobile"] = request.OperatorMobile
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorMobileVerifyCode)) {
		query["OperatorMobileVerifyCode"] = request.OperatorMobileVerifyCode
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorName)) {
		query["OperatorName"] = request.OperatorName
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.QualificationId)) {
		query["QualificationId"] = request.QualificationId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.TransferPhoneNumberInfos)) {
		query["TransferPhoneNumberInfos"] = request.TransferPhoneNumberInfos
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SubmitHotlineTransferRegister"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
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

/**
 * ### QPS limits
 * You can call this operation up to 100 times per second per account.
 *
 * @param request SubmitHotlineTransferRegisterRequest
 * @return SubmitHotlineTransferRegisterResponse
 */
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

func (client *Client) UpgradeVideoFileWithOptions(request *UpgradeVideoFileRequest, runtime *util.RuntimeOptions) (_result *UpgradeVideoFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CallId)) {
		query["CallId"] = request.CallId
	}

	if !tea.BoolValue(util.IsUnset(request.CalledNumber)) {
		query["CalledNumber"] = request.CalledNumber
	}

	if !tea.BoolValue(util.IsUnset(request.MediaType)) {
		query["MediaType"] = request.MediaType
	}

	if !tea.BoolValue(util.IsUnset(request.OutId)) {
		query["OutId"] = request.OutId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpgradeVideoFile"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpgradeVideoFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpgradeVideoFile(request *UpgradeVideoFileRequest) (_result *UpgradeVideoFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpgradeVideoFileResponse{}
	_body, _err := client.UpgradeVideoFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * ### QPS limits
 * You can call this operation up to 100 times per second per account.
 *
 * @param request UploadRobotTaskCalledFileRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return UploadRobotTaskCalledFileResponse
 */
func (client *Client) UploadRobotTaskCalledFileWithOptions(request *UploadRobotTaskCalledFileRequest, runtime *util.RuntimeOptions) (_result *UploadRobotTaskCalledFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CalledNumber)) {
		query["CalledNumber"] = request.CalledNumber
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.TtsParam)) {
		query["TtsParam"] = request.TtsParam
	}

	if !tea.BoolValue(util.IsUnset(request.TtsParamHead)) {
		query["TtsParamHead"] = request.TtsParamHead
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UploadRobotTaskCalledFile"),
		Version:     tea.String("2017-05-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
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

/**
 * ### QPS limits
 * You can call this operation up to 100 times per second per account.
 *
 * @param request UploadRobotTaskCalledFileRequest
 * @return UploadRobotTaskCalledFileResponse
 */
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
