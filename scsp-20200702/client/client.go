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

type AnswerCallRequest struct {
	AccountName  *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	CallId       *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ConnectionId *string `json:"ConnectionId,omitempty" xml:"ConnectionId,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	JobId        *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s AnswerCallRequest) String() string {
	return tea.Prettify(s)
}

func (s AnswerCallRequest) GoString() string {
	return s.String()
}

func (s *AnswerCallRequest) SetAccountName(v string) *AnswerCallRequest {
	s.AccountName = &v
	return s
}

func (s *AnswerCallRequest) SetCallId(v string) *AnswerCallRequest {
	s.CallId = &v
	return s
}

func (s *AnswerCallRequest) SetClientToken(v string) *AnswerCallRequest {
	s.ClientToken = &v
	return s
}

func (s *AnswerCallRequest) SetConnectionId(v string) *AnswerCallRequest {
	s.ConnectionId = &v
	return s
}

func (s *AnswerCallRequest) SetInstanceId(v string) *AnswerCallRequest {
	s.InstanceId = &v
	return s
}

func (s *AnswerCallRequest) SetJobId(v string) *AnswerCallRequest {
	s.JobId = &v
	return s
}

type AnswerCallResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AnswerCallResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AnswerCallResponseBody) GoString() string {
	return s.String()
}

func (s *AnswerCallResponseBody) SetCode(v string) *AnswerCallResponseBody {
	s.Code = &v
	return s
}

func (s *AnswerCallResponseBody) SetMessage(v string) *AnswerCallResponseBody {
	s.Message = &v
	return s
}

func (s *AnswerCallResponseBody) SetRequestId(v string) *AnswerCallResponseBody {
	s.RequestId = &v
	return s
}

func (s *AnswerCallResponseBody) SetSuccess(v bool) *AnswerCallResponseBody {
	s.Success = &v
	return s
}

type AnswerCallResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AnswerCallResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AnswerCallResponse) String() string {
	return tea.Prettify(s)
}

func (s AnswerCallResponse) GoString() string {
	return s.String()
}

func (s *AnswerCallResponse) SetHeaders(v map[string]*string) *AnswerCallResponse {
	s.Headers = v
	return s
}

func (s *AnswerCallResponse) SetStatusCode(v int32) *AnswerCallResponse {
	s.StatusCode = &v
	return s
}

func (s *AnswerCallResponse) SetBody(v *AnswerCallResponseBody) *AnswerCallResponse {
	s.Body = v
	return s
}

type AnwserAgentMonitorRequest struct {
	AccountName       *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	CallId            *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	CallerParentId    *int64  `json:"CallerParentId,omitempty" xml:"CallerParentId,omitempty"`
	CallerType        *string `json:"CallerType,omitempty" xml:"CallerType,omitempty"`
	CallerUid         *int64  `json:"CallerUid,omitempty" xml:"CallerUid,omitempty"`
	ClientToken       *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ConnectionId      *string `json:"ConnectionId,omitempty" xml:"ConnectionId,omitempty"`
	InstanceId        *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	JobId             *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	RequestId         *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StsTokenCallerUid *int64  `json:"StsTokenCallerUid,omitempty" xml:"StsTokenCallerUid,omitempty"`
}

func (s AnwserAgentMonitorRequest) String() string {
	return tea.Prettify(s)
}

func (s AnwserAgentMonitorRequest) GoString() string {
	return s.String()
}

func (s *AnwserAgentMonitorRequest) SetAccountName(v string) *AnwserAgentMonitorRequest {
	s.AccountName = &v
	return s
}

func (s *AnwserAgentMonitorRequest) SetCallId(v string) *AnwserAgentMonitorRequest {
	s.CallId = &v
	return s
}

func (s *AnwserAgentMonitorRequest) SetCallerParentId(v int64) *AnwserAgentMonitorRequest {
	s.CallerParentId = &v
	return s
}

func (s *AnwserAgentMonitorRequest) SetCallerType(v string) *AnwserAgentMonitorRequest {
	s.CallerType = &v
	return s
}

func (s *AnwserAgentMonitorRequest) SetCallerUid(v int64) *AnwserAgentMonitorRequest {
	s.CallerUid = &v
	return s
}

func (s *AnwserAgentMonitorRequest) SetClientToken(v string) *AnwserAgentMonitorRequest {
	s.ClientToken = &v
	return s
}

func (s *AnwserAgentMonitorRequest) SetConnectionId(v string) *AnwserAgentMonitorRequest {
	s.ConnectionId = &v
	return s
}

func (s *AnwserAgentMonitorRequest) SetInstanceId(v string) *AnwserAgentMonitorRequest {
	s.InstanceId = &v
	return s
}

func (s *AnwserAgentMonitorRequest) SetJobId(v string) *AnwserAgentMonitorRequest {
	s.JobId = &v
	return s
}

func (s *AnwserAgentMonitorRequest) SetRequestId(v string) *AnwserAgentMonitorRequest {
	s.RequestId = &v
	return s
}

func (s *AnwserAgentMonitorRequest) SetStsTokenCallerUid(v int64) *AnwserAgentMonitorRequest {
	s.StsTokenCallerUid = &v
	return s
}

type AnwserAgentMonitorResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AnwserAgentMonitorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AnwserAgentMonitorResponseBody) GoString() string {
	return s.String()
}

func (s *AnwserAgentMonitorResponseBody) SetCode(v string) *AnwserAgentMonitorResponseBody {
	s.Code = &v
	return s
}

func (s *AnwserAgentMonitorResponseBody) SetMessage(v string) *AnwserAgentMonitorResponseBody {
	s.Message = &v
	return s
}

func (s *AnwserAgentMonitorResponseBody) SetRequestId(v string) *AnwserAgentMonitorResponseBody {
	s.RequestId = &v
	return s
}

func (s *AnwserAgentMonitorResponseBody) SetSuccess(v bool) *AnwserAgentMonitorResponseBody {
	s.Success = &v
	return s
}

type AnwserAgentMonitorResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AnwserAgentMonitorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AnwserAgentMonitorResponse) String() string {
	return tea.Prettify(s)
}

func (s AnwserAgentMonitorResponse) GoString() string {
	return s.String()
}

func (s *AnwserAgentMonitorResponse) SetHeaders(v map[string]*string) *AnwserAgentMonitorResponse {
	s.Headers = v
	return s
}

func (s *AnwserAgentMonitorResponse) SetStatusCode(v int32) *AnwserAgentMonitorResponse {
	s.StatusCode = &v
	return s
}

func (s *AnwserAgentMonitorResponse) SetBody(v *AnwserAgentMonitorResponseBody) *AnwserAgentMonitorResponse {
	s.Body = v
	return s
}

type AppMessagePushRequest struct {
	ExpirationTime *int64  `json:"ExpirationTime,omitempty" xml:"ExpirationTime,omitempty"`
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Status         *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	UserId         *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s AppMessagePushRequest) String() string {
	return tea.Prettify(s)
}

func (s AppMessagePushRequest) GoString() string {
	return s.String()
}

func (s *AppMessagePushRequest) SetExpirationTime(v int64) *AppMessagePushRequest {
	s.ExpirationTime = &v
	return s
}

func (s *AppMessagePushRequest) SetInstanceId(v string) *AppMessagePushRequest {
	s.InstanceId = &v
	return s
}

func (s *AppMessagePushRequest) SetStatus(v int32) *AppMessagePushRequest {
	s.Status = &v
	return s
}

func (s *AppMessagePushRequest) SetUserId(v string) *AppMessagePushRequest {
	s.UserId = &v
	return s
}

type AppMessagePushResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AppMessagePushResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AppMessagePushResponseBody) GoString() string {
	return s.String()
}

func (s *AppMessagePushResponseBody) SetCode(v string) *AppMessagePushResponseBody {
	s.Code = &v
	return s
}

func (s *AppMessagePushResponseBody) SetData(v string) *AppMessagePushResponseBody {
	s.Data = &v
	return s
}

func (s *AppMessagePushResponseBody) SetMessage(v string) *AppMessagePushResponseBody {
	s.Message = &v
	return s
}

func (s *AppMessagePushResponseBody) SetRequestId(v string) *AppMessagePushResponseBody {
	s.RequestId = &v
	return s
}

func (s *AppMessagePushResponseBody) SetSuccess(v bool) *AppMessagePushResponseBody {
	s.Success = &v
	return s
}

type AppMessagePushResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AppMessagePushResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AppMessagePushResponse) String() string {
	return tea.Prettify(s)
}

func (s AppMessagePushResponse) GoString() string {
	return s.String()
}

func (s *AppMessagePushResponse) SetHeaders(v map[string]*string) *AppMessagePushResponse {
	s.Headers = v
	return s
}

func (s *AppMessagePushResponse) SetStatusCode(v int32) *AppMessagePushResponse {
	s.StatusCode = &v
	return s
}

func (s *AppMessagePushResponse) SetBody(v *AppMessagePushResponseBody) *AppMessagePushResponse {
	s.Body = v
	return s
}

type AssignTicketRequest struct {
	AcceptorId  *int64  `json:"AcceptorId,omitempty" xml:"AcceptorId,omitempty"`
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OperatorId  *int64  `json:"OperatorId,omitempty" xml:"OperatorId,omitempty"`
	TicketId    *int64  `json:"TicketId,omitempty" xml:"TicketId,omitempty"`
}

func (s AssignTicketRequest) String() string {
	return tea.Prettify(s)
}

func (s AssignTicketRequest) GoString() string {
	return s.String()
}

func (s *AssignTicketRequest) SetAcceptorId(v int64) *AssignTicketRequest {
	s.AcceptorId = &v
	return s
}

func (s *AssignTicketRequest) SetClientToken(v string) *AssignTicketRequest {
	s.ClientToken = &v
	return s
}

func (s *AssignTicketRequest) SetInstanceId(v string) *AssignTicketRequest {
	s.InstanceId = &v
	return s
}

func (s *AssignTicketRequest) SetOperatorId(v int64) *AssignTicketRequest {
	s.OperatorId = &v
	return s
}

func (s *AssignTicketRequest) SetTicketId(v int64) *AssignTicketRequest {
	s.TicketId = &v
	return s
}

type AssignTicketResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AssignTicketResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AssignTicketResponseBody) GoString() string {
	return s.String()
}

func (s *AssignTicketResponseBody) SetCode(v string) *AssignTicketResponseBody {
	s.Code = &v
	return s
}

func (s *AssignTicketResponseBody) SetMessage(v string) *AssignTicketResponseBody {
	s.Message = &v
	return s
}

func (s *AssignTicketResponseBody) SetRequestId(v string) *AssignTicketResponseBody {
	s.RequestId = &v
	return s
}

func (s *AssignTicketResponseBody) SetSuccess(v bool) *AssignTicketResponseBody {
	s.Success = &v
	return s
}

type AssignTicketResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AssignTicketResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AssignTicketResponse) String() string {
	return tea.Prettify(s)
}

func (s AssignTicketResponse) GoString() string {
	return s.String()
}

func (s *AssignTicketResponse) SetHeaders(v map[string]*string) *AssignTicketResponse {
	s.Headers = v
	return s
}

func (s *AssignTicketResponse) SetStatusCode(v int32) *AssignTicketResponse {
	s.StatusCode = &v
	return s
}

func (s *AssignTicketResponse) SetBody(v *AssignTicketResponseBody) *AssignTicketResponse {
	s.Body = v
	return s
}

type BindAgentHotlinePhoneRequest struct {
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Phone       *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
	VerifyCode  *string `json:"VerifyCode,omitempty" xml:"VerifyCode,omitempty"`
}

func (s BindAgentHotlinePhoneRequest) String() string {
	return tea.Prettify(s)
}

func (s BindAgentHotlinePhoneRequest) GoString() string {
	return s.String()
}

func (s *BindAgentHotlinePhoneRequest) SetAccountName(v string) *BindAgentHotlinePhoneRequest {
	s.AccountName = &v
	return s
}

func (s *BindAgentHotlinePhoneRequest) SetClientToken(v string) *BindAgentHotlinePhoneRequest {
	s.ClientToken = &v
	return s
}

func (s *BindAgentHotlinePhoneRequest) SetInstanceId(v string) *BindAgentHotlinePhoneRequest {
	s.InstanceId = &v
	return s
}

func (s *BindAgentHotlinePhoneRequest) SetPhone(v string) *BindAgentHotlinePhoneRequest {
	s.Phone = &v
	return s
}

func (s *BindAgentHotlinePhoneRequest) SetVerifyCode(v string) *BindAgentHotlinePhoneRequest {
	s.VerifyCode = &v
	return s
}

type BindAgentHotlinePhoneResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BindAgentHotlinePhoneResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BindAgentHotlinePhoneResponseBody) GoString() string {
	return s.String()
}

func (s *BindAgentHotlinePhoneResponseBody) SetCode(v string) *BindAgentHotlinePhoneResponseBody {
	s.Code = &v
	return s
}

func (s *BindAgentHotlinePhoneResponseBody) SetMessage(v string) *BindAgentHotlinePhoneResponseBody {
	s.Message = &v
	return s
}

func (s *BindAgentHotlinePhoneResponseBody) SetRequestId(v string) *BindAgentHotlinePhoneResponseBody {
	s.RequestId = &v
	return s
}

func (s *BindAgentHotlinePhoneResponseBody) SetSuccess(v bool) *BindAgentHotlinePhoneResponseBody {
	s.Success = &v
	return s
}

type BindAgentHotlinePhoneResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *BindAgentHotlinePhoneResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BindAgentHotlinePhoneResponse) String() string {
	return tea.Prettify(s)
}

func (s BindAgentHotlinePhoneResponse) GoString() string {
	return s.String()
}

func (s *BindAgentHotlinePhoneResponse) SetHeaders(v map[string]*string) *BindAgentHotlinePhoneResponse {
	s.Headers = v
	return s
}

func (s *BindAgentHotlinePhoneResponse) SetStatusCode(v int32) *BindAgentHotlinePhoneResponse {
	s.StatusCode = &v
	return s
}

func (s *BindAgentHotlinePhoneResponse) SetBody(v *BindAgentHotlinePhoneResponseBody) *BindAgentHotlinePhoneResponse {
	s.Body = v
	return s
}

type ChangeChatAgentStatusRequest struct {
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Method      *string `json:"Method,omitempty" xml:"Method,omitempty"`
}

func (s ChangeChatAgentStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s ChangeChatAgentStatusRequest) GoString() string {
	return s.String()
}

func (s *ChangeChatAgentStatusRequest) SetAccountName(v string) *ChangeChatAgentStatusRequest {
	s.AccountName = &v
	return s
}

func (s *ChangeChatAgentStatusRequest) SetClientToken(v string) *ChangeChatAgentStatusRequest {
	s.ClientToken = &v
	return s
}

func (s *ChangeChatAgentStatusRequest) SetInstanceId(v string) *ChangeChatAgentStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *ChangeChatAgentStatusRequest) SetMethod(v string) *ChangeChatAgentStatusRequest {
	s.Method = &v
	return s
}

type ChangeChatAgentStatusResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *string `json:"Data,omitempty" xml:"Data,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ChangeChatAgentStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ChangeChatAgentStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeChatAgentStatusResponseBody) SetCode(v string) *ChangeChatAgentStatusResponseBody {
	s.Code = &v
	return s
}

func (s *ChangeChatAgentStatusResponseBody) SetData(v string) *ChangeChatAgentStatusResponseBody {
	s.Data = &v
	return s
}

func (s *ChangeChatAgentStatusResponseBody) SetHttpStatusCode(v int32) *ChangeChatAgentStatusResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ChangeChatAgentStatusResponseBody) SetMessage(v string) *ChangeChatAgentStatusResponseBody {
	s.Message = &v
	return s
}

func (s *ChangeChatAgentStatusResponseBody) SetRequestId(v string) *ChangeChatAgentStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChangeChatAgentStatusResponseBody) SetSuccess(v bool) *ChangeChatAgentStatusResponseBody {
	s.Success = &v
	return s
}

type ChangeChatAgentStatusResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ChangeChatAgentStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ChangeChatAgentStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s ChangeChatAgentStatusResponse) GoString() string {
	return s.String()
}

func (s *ChangeChatAgentStatusResponse) SetHeaders(v map[string]*string) *ChangeChatAgentStatusResponse {
	s.Headers = v
	return s
}

func (s *ChangeChatAgentStatusResponse) SetStatusCode(v int32) *ChangeChatAgentStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *ChangeChatAgentStatusResponse) SetBody(v *ChangeChatAgentStatusResponseBody) *ChangeChatAgentStatusResponse {
	s.Body = v
	return s
}

type CloseTicketRequest struct {
	ActionItems *string `json:"ActionItems,omitempty" xml:"ActionItems,omitempty"`
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OperatorId  *int64  `json:"OperatorId,omitempty" xml:"OperatorId,omitempty"`
	TicketId    *int64  `json:"TicketId,omitempty" xml:"TicketId,omitempty"`
}

func (s CloseTicketRequest) String() string {
	return tea.Prettify(s)
}

func (s CloseTicketRequest) GoString() string {
	return s.String()
}

func (s *CloseTicketRequest) SetActionItems(v string) *CloseTicketRequest {
	s.ActionItems = &v
	return s
}

func (s *CloseTicketRequest) SetClientToken(v string) *CloseTicketRequest {
	s.ClientToken = &v
	return s
}

func (s *CloseTicketRequest) SetInstanceId(v string) *CloseTicketRequest {
	s.InstanceId = &v
	return s
}

func (s *CloseTicketRequest) SetOperatorId(v int64) *CloseTicketRequest {
	s.OperatorId = &v
	return s
}

func (s *CloseTicketRequest) SetTicketId(v int64) *CloseTicketRequest {
	s.TicketId = &v
	return s
}

type CloseTicketResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int64  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CloseTicketResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CloseTicketResponseBody) GoString() string {
	return s.String()
}

func (s *CloseTicketResponseBody) SetCode(v string) *CloseTicketResponseBody {
	s.Code = &v
	return s
}

func (s *CloseTicketResponseBody) SetHttpStatusCode(v int64) *CloseTicketResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CloseTicketResponseBody) SetMessage(v string) *CloseTicketResponseBody {
	s.Message = &v
	return s
}

func (s *CloseTicketResponseBody) SetRequestId(v string) *CloseTicketResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloseTicketResponseBody) SetSuccess(v bool) *CloseTicketResponseBody {
	s.Success = &v
	return s
}

type CloseTicketResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CloseTicketResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CloseTicketResponse) String() string {
	return tea.Prettify(s)
}

func (s CloseTicketResponse) GoString() string {
	return s.String()
}

func (s *CloseTicketResponse) SetHeaders(v map[string]*string) *CloseTicketResponse {
	s.Headers = v
	return s
}

func (s *CloseTicketResponse) SetStatusCode(v int32) *CloseTicketResponse {
	s.StatusCode = &v
	return s
}

func (s *CloseTicketResponse) SetBody(v *CloseTicketResponseBody) *CloseTicketResponse {
	s.Body = v
	return s
}

type CreateAgentRequest struct {
	AccountName      *string  `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	ClientToken      *string  `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DisplayName      *string  `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	InstanceId       *string  `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	SkillGroupId     []*int64 `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty" type:"Repeated"`
	SkillGroupIdList []*int64 `json:"SkillGroupIdList,omitempty" xml:"SkillGroupIdList,omitempty" type:"Repeated"`
}

func (s CreateAgentRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAgentRequest) GoString() string {
	return s.String()
}

func (s *CreateAgentRequest) SetAccountName(v string) *CreateAgentRequest {
	s.AccountName = &v
	return s
}

func (s *CreateAgentRequest) SetClientToken(v string) *CreateAgentRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateAgentRequest) SetDisplayName(v string) *CreateAgentRequest {
	s.DisplayName = &v
	return s
}

func (s *CreateAgentRequest) SetInstanceId(v string) *CreateAgentRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateAgentRequest) SetSkillGroupId(v []*int64) *CreateAgentRequest {
	s.SkillGroupId = v
	return s
}

func (s *CreateAgentRequest) SetSkillGroupIdList(v []*int64) *CreateAgentRequest {
	s.SkillGroupIdList = v
	return s
}

type CreateAgentResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *int64  `json:"Data,omitempty" xml:"Data,omitempty"`
	HttpStatusCode *int64  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateAgentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAgentResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAgentResponseBody) SetCode(v string) *CreateAgentResponseBody {
	s.Code = &v
	return s
}

func (s *CreateAgentResponseBody) SetData(v int64) *CreateAgentResponseBody {
	s.Data = &v
	return s
}

func (s *CreateAgentResponseBody) SetHttpStatusCode(v int64) *CreateAgentResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateAgentResponseBody) SetMessage(v string) *CreateAgentResponseBody {
	s.Message = &v
	return s
}

func (s *CreateAgentResponseBody) SetRequestId(v string) *CreateAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAgentResponseBody) SetSuccess(v bool) *CreateAgentResponseBody {
	s.Success = &v
	return s
}

type CreateAgentResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateAgentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateAgentResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAgentResponse) GoString() string {
	return s.String()
}

func (s *CreateAgentResponse) SetHeaders(v map[string]*string) *CreateAgentResponse {
	s.Headers = v
	return s
}

func (s *CreateAgentResponse) SetStatusCode(v int32) *CreateAgentResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAgentResponse) SetBody(v *CreateAgentResponseBody) *CreateAgentResponse {
	s.Body = v
	return s
}

type CreateCustomerRequest struct {
	BizType     *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	Contacter   *string `json:"Contacter,omitempty" xml:"Contacter,omitempty"`
	Dingding    *string `json:"Dingding,omitempty" xml:"Dingding,omitempty"`
	Email       *string `json:"Email,omitempty" xml:"Email,omitempty"`
	Industry    *string `json:"Industry,omitempty" xml:"Industry,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ManagerName *string `json:"ManagerName,omitempty" xml:"ManagerName,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OuterId     *string `json:"OuterId,omitempty" xml:"OuterId,omitempty"`
	OuterIdType *int32  `json:"OuterIdType,omitempty" xml:"OuterIdType,omitempty"`
	Phone       *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
	Position    *string `json:"Position,omitempty" xml:"Position,omitempty"`
	ProdLineId  *int64  `json:"ProdLineId,omitempty" xml:"ProdLineId,omitempty"`
	TypeCode    *string `json:"TypeCode,omitempty" xml:"TypeCode,omitempty"`
}

func (s CreateCustomerRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomerRequest) GoString() string {
	return s.String()
}

func (s *CreateCustomerRequest) SetBizType(v string) *CreateCustomerRequest {
	s.BizType = &v
	return s
}

func (s *CreateCustomerRequest) SetContacter(v string) *CreateCustomerRequest {
	s.Contacter = &v
	return s
}

func (s *CreateCustomerRequest) SetDingding(v string) *CreateCustomerRequest {
	s.Dingding = &v
	return s
}

func (s *CreateCustomerRequest) SetEmail(v string) *CreateCustomerRequest {
	s.Email = &v
	return s
}

func (s *CreateCustomerRequest) SetIndustry(v string) *CreateCustomerRequest {
	s.Industry = &v
	return s
}

func (s *CreateCustomerRequest) SetInstanceId(v string) *CreateCustomerRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateCustomerRequest) SetManagerName(v string) *CreateCustomerRequest {
	s.ManagerName = &v
	return s
}

func (s *CreateCustomerRequest) SetName(v string) *CreateCustomerRequest {
	s.Name = &v
	return s
}

func (s *CreateCustomerRequest) SetOuterId(v string) *CreateCustomerRequest {
	s.OuterId = &v
	return s
}

func (s *CreateCustomerRequest) SetOuterIdType(v int32) *CreateCustomerRequest {
	s.OuterIdType = &v
	return s
}

func (s *CreateCustomerRequest) SetPhone(v string) *CreateCustomerRequest {
	s.Phone = &v
	return s
}

func (s *CreateCustomerRequest) SetPosition(v string) *CreateCustomerRequest {
	s.Position = &v
	return s
}

func (s *CreateCustomerRequest) SetProdLineId(v int64) *CreateCustomerRequest {
	s.ProdLineId = &v
	return s
}

func (s *CreateCustomerRequest) SetTypeCode(v string) *CreateCustomerRequest {
	s.TypeCode = &v
	return s
}

type CreateCustomerResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *int64  `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateCustomerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomerResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCustomerResponseBody) SetCode(v string) *CreateCustomerResponseBody {
	s.Code = &v
	return s
}

func (s *CreateCustomerResponseBody) SetData(v int64) *CreateCustomerResponseBody {
	s.Data = &v
	return s
}

func (s *CreateCustomerResponseBody) SetMessage(v string) *CreateCustomerResponseBody {
	s.Message = &v
	return s
}

func (s *CreateCustomerResponseBody) SetRequestId(v string) *CreateCustomerResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCustomerResponseBody) SetSuccess(v bool) *CreateCustomerResponseBody {
	s.Success = &v
	return s
}

type CreateCustomerResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateCustomerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateCustomerResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomerResponse) GoString() string {
	return s.String()
}

func (s *CreateCustomerResponse) SetHeaders(v map[string]*string) *CreateCustomerResponse {
	s.Headers = v
	return s
}

func (s *CreateCustomerResponse) SetStatusCode(v int32) *CreateCustomerResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCustomerResponse) SetBody(v *CreateCustomerResponseBody) *CreateCustomerResponse {
	s.Body = v
	return s
}

type CreateEntityIvrRouteRequest struct {
	DepartmentId         *string `json:"DepartmentId,omitempty" xml:"DepartmentId,omitempty"`
	EntityBizCode        *string `json:"EntityBizCode,omitempty" xml:"EntityBizCode,omitempty"`
	EntityBizCodeType    *string `json:"EntityBizCodeType,omitempty" xml:"EntityBizCodeType,omitempty"`
	EntityId             *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	EntityName           *string `json:"EntityName,omitempty" xml:"EntityName,omitempty"`
	EntityRelationNumber *string `json:"EntityRelationNumber,omitempty" xml:"EntityRelationNumber,omitempty"`
	ExtInfo              *string `json:"ExtInfo,omitempty" xml:"ExtInfo,omitempty"`
	GroupId              *int64  `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ServiceId            *int64  `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
}

func (s CreateEntityIvrRouteRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateEntityIvrRouteRequest) GoString() string {
	return s.String()
}

func (s *CreateEntityIvrRouteRequest) SetDepartmentId(v string) *CreateEntityIvrRouteRequest {
	s.DepartmentId = &v
	return s
}

func (s *CreateEntityIvrRouteRequest) SetEntityBizCode(v string) *CreateEntityIvrRouteRequest {
	s.EntityBizCode = &v
	return s
}

func (s *CreateEntityIvrRouteRequest) SetEntityBizCodeType(v string) *CreateEntityIvrRouteRequest {
	s.EntityBizCodeType = &v
	return s
}

func (s *CreateEntityIvrRouteRequest) SetEntityId(v string) *CreateEntityIvrRouteRequest {
	s.EntityId = &v
	return s
}

func (s *CreateEntityIvrRouteRequest) SetEntityName(v string) *CreateEntityIvrRouteRequest {
	s.EntityName = &v
	return s
}

func (s *CreateEntityIvrRouteRequest) SetEntityRelationNumber(v string) *CreateEntityIvrRouteRequest {
	s.EntityRelationNumber = &v
	return s
}

func (s *CreateEntityIvrRouteRequest) SetExtInfo(v string) *CreateEntityIvrRouteRequest {
	s.ExtInfo = &v
	return s
}

func (s *CreateEntityIvrRouteRequest) SetGroupId(v int64) *CreateEntityIvrRouteRequest {
	s.GroupId = &v
	return s
}

func (s *CreateEntityIvrRouteRequest) SetInstanceId(v string) *CreateEntityIvrRouteRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateEntityIvrRouteRequest) SetServiceId(v int64) *CreateEntityIvrRouteRequest {
	s.ServiceId = &v
	return s
}

type CreateEntityIvrRouteResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateEntityIvrRouteResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateEntityIvrRouteResponseBody) GoString() string {
	return s.String()
}

func (s *CreateEntityIvrRouteResponseBody) SetCode(v string) *CreateEntityIvrRouteResponseBody {
	s.Code = &v
	return s
}

func (s *CreateEntityIvrRouteResponseBody) SetMessage(v string) *CreateEntityIvrRouteResponseBody {
	s.Message = &v
	return s
}

func (s *CreateEntityIvrRouteResponseBody) SetRequestId(v string) *CreateEntityIvrRouteResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateEntityIvrRouteResponseBody) SetSuccess(v bool) *CreateEntityIvrRouteResponseBody {
	s.Success = &v
	return s
}

type CreateEntityIvrRouteResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateEntityIvrRouteResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateEntityIvrRouteResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateEntityIvrRouteResponse) GoString() string {
	return s.String()
}

func (s *CreateEntityIvrRouteResponse) SetHeaders(v map[string]*string) *CreateEntityIvrRouteResponse {
	s.Headers = v
	return s
}

func (s *CreateEntityIvrRouteResponse) SetStatusCode(v int32) *CreateEntityIvrRouteResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateEntityIvrRouteResponse) SetBody(v *CreateEntityIvrRouteResponseBody) *CreateEntityIvrRouteResponse {
	s.Body = v
	return s
}

type CreateOuterCallCenterDataRequest struct {
	BizId         *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	BizType       *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	CallType      *string `json:"CallType,omitempty" xml:"CallType,omitempty"`
	EndReason     *string `json:"EndReason,omitempty" xml:"EndReason,omitempty"`
	ExtInfo       *string `json:"ExtInfo,omitempty" xml:"ExtInfo,omitempty"`
	FromPhoneNum  *string `json:"FromPhoneNum,omitempty" xml:"FromPhoneNum,omitempty"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InterveneTime *string `json:"InterveneTime,omitempty" xml:"InterveneTime,omitempty"`
	RecordUrl     *string `json:"RecordUrl,omitempty" xml:"RecordUrl,omitempty"`
	SessionId     *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	TenantId      *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	ToPhoneNum    *string `json:"ToPhoneNum,omitempty" xml:"ToPhoneNum,omitempty"`
	UserInfo      *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s CreateOuterCallCenterDataRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateOuterCallCenterDataRequest) GoString() string {
	return s.String()
}

func (s *CreateOuterCallCenterDataRequest) SetBizId(v string) *CreateOuterCallCenterDataRequest {
	s.BizId = &v
	return s
}

func (s *CreateOuterCallCenterDataRequest) SetBizType(v string) *CreateOuterCallCenterDataRequest {
	s.BizType = &v
	return s
}

func (s *CreateOuterCallCenterDataRequest) SetCallType(v string) *CreateOuterCallCenterDataRequest {
	s.CallType = &v
	return s
}

func (s *CreateOuterCallCenterDataRequest) SetEndReason(v string) *CreateOuterCallCenterDataRequest {
	s.EndReason = &v
	return s
}

func (s *CreateOuterCallCenterDataRequest) SetExtInfo(v string) *CreateOuterCallCenterDataRequest {
	s.ExtInfo = &v
	return s
}

func (s *CreateOuterCallCenterDataRequest) SetFromPhoneNum(v string) *CreateOuterCallCenterDataRequest {
	s.FromPhoneNum = &v
	return s
}

func (s *CreateOuterCallCenterDataRequest) SetInstanceId(v string) *CreateOuterCallCenterDataRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateOuterCallCenterDataRequest) SetInterveneTime(v string) *CreateOuterCallCenterDataRequest {
	s.InterveneTime = &v
	return s
}

func (s *CreateOuterCallCenterDataRequest) SetRecordUrl(v string) *CreateOuterCallCenterDataRequest {
	s.RecordUrl = &v
	return s
}

func (s *CreateOuterCallCenterDataRequest) SetSessionId(v string) *CreateOuterCallCenterDataRequest {
	s.SessionId = &v
	return s
}

func (s *CreateOuterCallCenterDataRequest) SetTenantId(v string) *CreateOuterCallCenterDataRequest {
	s.TenantId = &v
	return s
}

func (s *CreateOuterCallCenterDataRequest) SetToPhoneNum(v string) *CreateOuterCallCenterDataRequest {
	s.ToPhoneNum = &v
	return s
}

func (s *CreateOuterCallCenterDataRequest) SetUserInfo(v string) *CreateOuterCallCenterDataRequest {
	s.UserInfo = &v
	return s
}

type CreateOuterCallCenterDataResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int64  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateOuterCallCenterDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateOuterCallCenterDataResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOuterCallCenterDataResponseBody) SetCode(v string) *CreateOuterCallCenterDataResponseBody {
	s.Code = &v
	return s
}

func (s *CreateOuterCallCenterDataResponseBody) SetHttpStatusCode(v int64) *CreateOuterCallCenterDataResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateOuterCallCenterDataResponseBody) SetMessage(v string) *CreateOuterCallCenterDataResponseBody {
	s.Message = &v
	return s
}

func (s *CreateOuterCallCenterDataResponseBody) SetRequestId(v string) *CreateOuterCallCenterDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateOuterCallCenterDataResponseBody) SetSuccess(v bool) *CreateOuterCallCenterDataResponseBody {
	s.Success = &v
	return s
}

type CreateOuterCallCenterDataResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateOuterCallCenterDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateOuterCallCenterDataResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateOuterCallCenterDataResponse) GoString() string {
	return s.String()
}

func (s *CreateOuterCallCenterDataResponse) SetHeaders(v map[string]*string) *CreateOuterCallCenterDataResponse {
	s.Headers = v
	return s
}

func (s *CreateOuterCallCenterDataResponse) SetStatusCode(v int32) *CreateOuterCallCenterDataResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateOuterCallCenterDataResponse) SetBody(v *CreateOuterCallCenterDataResponseBody) *CreateOuterCallCenterDataResponse {
	s.Body = v
	return s
}

type CreateRoleRequest struct {
	ClientToken  *string  `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId   *string  `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Operator     *string  `json:"Operator,omitempty" xml:"Operator,omitempty"`
	PermissionId []*int64 `json:"PermissionId,omitempty" xml:"PermissionId,omitempty" type:"Repeated"`
	RoleName     *string  `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
}

func (s CreateRoleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRoleRequest) GoString() string {
	return s.String()
}

func (s *CreateRoleRequest) SetClientToken(v string) *CreateRoleRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateRoleRequest) SetInstanceId(v string) *CreateRoleRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateRoleRequest) SetOperator(v string) *CreateRoleRequest {
	s.Operator = &v
	return s
}

func (s *CreateRoleRequest) SetPermissionId(v []*int64) *CreateRoleRequest {
	s.PermissionId = v
	return s
}

func (s *CreateRoleRequest) SetRoleName(v string) *CreateRoleRequest {
	s.RoleName = &v
	return s
}

type CreateRoleResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *int64  `json:"Data,omitempty" xml:"Data,omitempty"`
	HttpStatusCode *int64  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateRoleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateRoleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRoleResponseBody) SetCode(v string) *CreateRoleResponseBody {
	s.Code = &v
	return s
}

func (s *CreateRoleResponseBody) SetData(v int64) *CreateRoleResponseBody {
	s.Data = &v
	return s
}

func (s *CreateRoleResponseBody) SetHttpStatusCode(v int64) *CreateRoleResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateRoleResponseBody) SetMessage(v string) *CreateRoleResponseBody {
	s.Message = &v
	return s
}

func (s *CreateRoleResponseBody) SetRequestId(v string) *CreateRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRoleResponseBody) SetSuccess(v bool) *CreateRoleResponseBody {
	s.Success = &v
	return s
}

type CreateRoleResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateRoleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateRoleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateRoleResponse) GoString() string {
	return s.String()
}

func (s *CreateRoleResponse) SetHeaders(v map[string]*string) *CreateRoleResponse {
	s.Headers = v
	return s
}

func (s *CreateRoleResponse) SetStatusCode(v int32) *CreateRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRoleResponse) SetBody(v *CreateRoleResponseBody) *CreateRoleResponse {
	s.Body = v
	return s
}

type CreateSkillGroupRequest struct {
	ChannelType    *int32  `json:"ChannelType,omitempty" xml:"ChannelType,omitempty"`
	ClientToken    *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DisplayName    *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	SkillGroupName *string `json:"SkillGroupName,omitempty" xml:"SkillGroupName,omitempty"`
}

func (s CreateSkillGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSkillGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateSkillGroupRequest) SetChannelType(v int32) *CreateSkillGroupRequest {
	s.ChannelType = &v
	return s
}

func (s *CreateSkillGroupRequest) SetClientToken(v string) *CreateSkillGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateSkillGroupRequest) SetDescription(v string) *CreateSkillGroupRequest {
	s.Description = &v
	return s
}

func (s *CreateSkillGroupRequest) SetDisplayName(v string) *CreateSkillGroupRequest {
	s.DisplayName = &v
	return s
}

func (s *CreateSkillGroupRequest) SetInstanceId(v string) *CreateSkillGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateSkillGroupRequest) SetSkillGroupName(v string) *CreateSkillGroupRequest {
	s.SkillGroupName = &v
	return s
}

type CreateSkillGroupResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *int64  `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateSkillGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSkillGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSkillGroupResponseBody) SetCode(v string) *CreateSkillGroupResponseBody {
	s.Code = &v
	return s
}

func (s *CreateSkillGroupResponseBody) SetData(v int64) *CreateSkillGroupResponseBody {
	s.Data = &v
	return s
}

func (s *CreateSkillGroupResponseBody) SetMessage(v string) *CreateSkillGroupResponseBody {
	s.Message = &v
	return s
}

func (s *CreateSkillGroupResponseBody) SetRequestId(v string) *CreateSkillGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSkillGroupResponseBody) SetSuccess(v bool) *CreateSkillGroupResponseBody {
	s.Success = &v
	return s
}

type CreateSkillGroupResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateSkillGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateSkillGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSkillGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateSkillGroupResponse) SetHeaders(v map[string]*string) *CreateSkillGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateSkillGroupResponse) SetStatusCode(v int32) *CreateSkillGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSkillGroupResponse) SetBody(v *CreateSkillGroupResponseBody) *CreateSkillGroupResponse {
	s.Body = v
	return s
}

type CreateSubTicketRequest struct {
	AgentId      *int64  `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	BizData      *string `json:"BizData,omitempty" xml:"BizData,omitempty"`
	CreatorId    *int64  `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	CreatorName  *string `json:"CreatorName,omitempty" xml:"CreatorName,omitempty"`
	FormData     *string `json:"FormData,omitempty" xml:"FormData,omitempty"`
	FromInfo     *string `json:"FromInfo,omitempty" xml:"FromInfo,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MemberId     *int64  `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	MemberName   *string `json:"MemberName,omitempty" xml:"MemberName,omitempty"`
	ParentCaseId *int64  `json:"ParentCaseId,omitempty" xml:"ParentCaseId,omitempty"`
	Priority     *int32  `json:"Priority,omitempty" xml:"Priority,omitempty"`
	TemplateId   *int64  `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s CreateSubTicketRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSubTicketRequest) GoString() string {
	return s.String()
}

func (s *CreateSubTicketRequest) SetAgentId(v int64) *CreateSubTicketRequest {
	s.AgentId = &v
	return s
}

func (s *CreateSubTicketRequest) SetBizData(v string) *CreateSubTicketRequest {
	s.BizData = &v
	return s
}

func (s *CreateSubTicketRequest) SetCreatorId(v int64) *CreateSubTicketRequest {
	s.CreatorId = &v
	return s
}

func (s *CreateSubTicketRequest) SetCreatorName(v string) *CreateSubTicketRequest {
	s.CreatorName = &v
	return s
}

func (s *CreateSubTicketRequest) SetFormData(v string) *CreateSubTicketRequest {
	s.FormData = &v
	return s
}

func (s *CreateSubTicketRequest) SetFromInfo(v string) *CreateSubTicketRequest {
	s.FromInfo = &v
	return s
}

func (s *CreateSubTicketRequest) SetInstanceId(v string) *CreateSubTicketRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateSubTicketRequest) SetMemberId(v int64) *CreateSubTicketRequest {
	s.MemberId = &v
	return s
}

func (s *CreateSubTicketRequest) SetMemberName(v string) *CreateSubTicketRequest {
	s.MemberName = &v
	return s
}

func (s *CreateSubTicketRequest) SetParentCaseId(v int64) *CreateSubTicketRequest {
	s.ParentCaseId = &v
	return s
}

func (s *CreateSubTicketRequest) SetPriority(v int32) *CreateSubTicketRequest {
	s.Priority = &v
	return s
}

func (s *CreateSubTicketRequest) SetTemplateId(v int64) *CreateSubTicketRequest {
	s.TemplateId = &v
	return s
}

type CreateSubTicketResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *int64  `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateSubTicketResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSubTicketResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSubTicketResponseBody) SetCode(v string) *CreateSubTicketResponseBody {
	s.Code = &v
	return s
}

func (s *CreateSubTicketResponseBody) SetData(v int64) *CreateSubTicketResponseBody {
	s.Data = &v
	return s
}

func (s *CreateSubTicketResponseBody) SetMessage(v string) *CreateSubTicketResponseBody {
	s.Message = &v
	return s
}

func (s *CreateSubTicketResponseBody) SetRequestId(v string) *CreateSubTicketResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSubTicketResponseBody) SetSuccess(v bool) *CreateSubTicketResponseBody {
	s.Success = &v
	return s
}

type CreateSubTicketResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateSubTicketResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateSubTicketResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSubTicketResponse) GoString() string {
	return s.String()
}

func (s *CreateSubTicketResponse) SetHeaders(v map[string]*string) *CreateSubTicketResponse {
	s.Headers = v
	return s
}

func (s *CreateSubTicketResponse) SetStatusCode(v int32) *CreateSubTicketResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSubTicketResponse) SetBody(v *CreateSubTicketResponseBody) *CreateSubTicketResponse {
	s.Body = v
	return s
}

type CreateThirdSsoAgentRequest struct {
	AccountId     *string  `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	AccountName   *string  `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	ClientId      *string  `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	ClientToken   *string  `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DisplayName   *string  `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	InstanceId    *string  `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RoleIds       []*int64 `json:"RoleIds,omitempty" xml:"RoleIds,omitempty" type:"Repeated"`
	SkillGroupIds []*int64 `json:"SkillGroupIds,omitempty" xml:"SkillGroupIds,omitempty" type:"Repeated"`
}

func (s CreateThirdSsoAgentRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateThirdSsoAgentRequest) GoString() string {
	return s.String()
}

func (s *CreateThirdSsoAgentRequest) SetAccountId(v string) *CreateThirdSsoAgentRequest {
	s.AccountId = &v
	return s
}

func (s *CreateThirdSsoAgentRequest) SetAccountName(v string) *CreateThirdSsoAgentRequest {
	s.AccountName = &v
	return s
}

func (s *CreateThirdSsoAgentRequest) SetClientId(v string) *CreateThirdSsoAgentRequest {
	s.ClientId = &v
	return s
}

func (s *CreateThirdSsoAgentRequest) SetClientToken(v string) *CreateThirdSsoAgentRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateThirdSsoAgentRequest) SetDisplayName(v string) *CreateThirdSsoAgentRequest {
	s.DisplayName = &v
	return s
}

func (s *CreateThirdSsoAgentRequest) SetInstanceId(v string) *CreateThirdSsoAgentRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateThirdSsoAgentRequest) SetRoleIds(v []*int64) *CreateThirdSsoAgentRequest {
	s.RoleIds = v
	return s
}

func (s *CreateThirdSsoAgentRequest) SetSkillGroupIds(v []*int64) *CreateThirdSsoAgentRequest {
	s.SkillGroupIds = v
	return s
}

type CreateThirdSsoAgentResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *int64  `json:"Data,omitempty" xml:"Data,omitempty"`
	HttpStatusCode *int64  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateThirdSsoAgentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateThirdSsoAgentResponseBody) GoString() string {
	return s.String()
}

func (s *CreateThirdSsoAgentResponseBody) SetCode(v string) *CreateThirdSsoAgentResponseBody {
	s.Code = &v
	return s
}

func (s *CreateThirdSsoAgentResponseBody) SetData(v int64) *CreateThirdSsoAgentResponseBody {
	s.Data = &v
	return s
}

func (s *CreateThirdSsoAgentResponseBody) SetHttpStatusCode(v int64) *CreateThirdSsoAgentResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateThirdSsoAgentResponseBody) SetMessage(v string) *CreateThirdSsoAgentResponseBody {
	s.Message = &v
	return s
}

func (s *CreateThirdSsoAgentResponseBody) SetRequestId(v string) *CreateThirdSsoAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateThirdSsoAgentResponseBody) SetSuccess(v bool) *CreateThirdSsoAgentResponseBody {
	s.Success = &v
	return s
}

type CreateThirdSsoAgentResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateThirdSsoAgentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateThirdSsoAgentResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateThirdSsoAgentResponse) GoString() string {
	return s.String()
}

func (s *CreateThirdSsoAgentResponse) SetHeaders(v map[string]*string) *CreateThirdSsoAgentResponse {
	s.Headers = v
	return s
}

func (s *CreateThirdSsoAgentResponse) SetStatusCode(v int32) *CreateThirdSsoAgentResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateThirdSsoAgentResponse) SetBody(v *CreateThirdSsoAgentResponseBody) *CreateThirdSsoAgentResponse {
	s.Body = v
	return s
}

type CreateTicketRequest struct {
	CarbonCopy  *string `json:"CarbonCopy,omitempty" xml:"CarbonCopy,omitempty"`
	CategoryId  *int64  `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	CreatorId   *int64  `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	CreatorName *string `json:"CreatorName,omitempty" xml:"CreatorName,omitempty"`
	CreatorType *int32  `json:"CreatorType,omitempty" xml:"CreatorType,omitempty"`
	FormData    *string `json:"FormData,omitempty" xml:"FormData,omitempty"`
	FromInfo    *string `json:"FromInfo,omitempty" xml:"FromInfo,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MemberId    *int64  `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	MemberName  *string `json:"MemberName,omitempty" xml:"MemberName,omitempty"`
	Priority    *int32  `json:"Priority,omitempty" xml:"Priority,omitempty"`
	TemplateId  *int64  `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s CreateTicketRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTicketRequest) GoString() string {
	return s.String()
}

func (s *CreateTicketRequest) SetCarbonCopy(v string) *CreateTicketRequest {
	s.CarbonCopy = &v
	return s
}

func (s *CreateTicketRequest) SetCategoryId(v int64) *CreateTicketRequest {
	s.CategoryId = &v
	return s
}

func (s *CreateTicketRequest) SetClientToken(v string) *CreateTicketRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateTicketRequest) SetCreatorId(v int64) *CreateTicketRequest {
	s.CreatorId = &v
	return s
}

func (s *CreateTicketRequest) SetCreatorName(v string) *CreateTicketRequest {
	s.CreatorName = &v
	return s
}

func (s *CreateTicketRequest) SetCreatorType(v int32) *CreateTicketRequest {
	s.CreatorType = &v
	return s
}

func (s *CreateTicketRequest) SetFormData(v string) *CreateTicketRequest {
	s.FormData = &v
	return s
}

func (s *CreateTicketRequest) SetFromInfo(v string) *CreateTicketRequest {
	s.FromInfo = &v
	return s
}

func (s *CreateTicketRequest) SetInstanceId(v string) *CreateTicketRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateTicketRequest) SetMemberId(v int64) *CreateTicketRequest {
	s.MemberId = &v
	return s
}

func (s *CreateTicketRequest) SetMemberName(v string) *CreateTicketRequest {
	s.MemberName = &v
	return s
}

func (s *CreateTicketRequest) SetPriority(v int32) *CreateTicketRequest {
	s.Priority = &v
	return s
}

func (s *CreateTicketRequest) SetTemplateId(v int64) *CreateTicketRequest {
	s.TemplateId = &v
	return s
}

type CreateTicketResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *int64  `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateTicketResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTicketResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTicketResponseBody) SetCode(v string) *CreateTicketResponseBody {
	s.Code = &v
	return s
}

func (s *CreateTicketResponseBody) SetData(v int64) *CreateTicketResponseBody {
	s.Data = &v
	return s
}

func (s *CreateTicketResponseBody) SetMessage(v string) *CreateTicketResponseBody {
	s.Message = &v
	return s
}

func (s *CreateTicketResponseBody) SetRequestId(v string) *CreateTicketResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTicketResponseBody) SetSuccess(v bool) *CreateTicketResponseBody {
	s.Success = &v
	return s
}

type CreateTicketResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateTicketResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateTicketResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTicketResponse) GoString() string {
	return s.String()
}

func (s *CreateTicketResponse) SetHeaders(v map[string]*string) *CreateTicketResponse {
	s.Headers = v
	return s
}

func (s *CreateTicketResponse) SetStatusCode(v int32) *CreateTicketResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTicketResponse) SetBody(v *CreateTicketResponseBody) *CreateTicketResponse {
	s.Body = v
	return s
}

type DeleteAgentRequest struct {
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DeleteAgentRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAgentRequest) GoString() string {
	return s.String()
}

func (s *DeleteAgentRequest) SetAccountName(v string) *DeleteAgentRequest {
	s.AccountName = &v
	return s
}

func (s *DeleteAgentRequest) SetClientToken(v string) *DeleteAgentRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteAgentRequest) SetInstanceId(v string) *DeleteAgentRequest {
	s.InstanceId = &v
	return s
}

type DeleteAgentResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int64  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteAgentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAgentResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAgentResponseBody) SetCode(v string) *DeleteAgentResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteAgentResponseBody) SetHttpStatusCode(v int64) *DeleteAgentResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteAgentResponseBody) SetMessage(v string) *DeleteAgentResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteAgentResponseBody) SetRequestId(v string) *DeleteAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAgentResponseBody) SetSuccess(v bool) *DeleteAgentResponseBody {
	s.Success = &v
	return s
}

type DeleteAgentResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteAgentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteAgentResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAgentResponse) GoString() string {
	return s.String()
}

func (s *DeleteAgentResponse) SetHeaders(v map[string]*string) *DeleteAgentResponse {
	s.Headers = v
	return s
}

func (s *DeleteAgentResponse) SetStatusCode(v int32) *DeleteAgentResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAgentResponse) SetBody(v *DeleteAgentResponseBody) *DeleteAgentResponse {
	s.Body = v
	return s
}

type DeleteEntityRouteRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	UniqueId   *int64  `json:"UniqueId,omitempty" xml:"UniqueId,omitempty"`
}

func (s DeleteEntityRouteRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteEntityRouteRequest) GoString() string {
	return s.String()
}

func (s *DeleteEntityRouteRequest) SetInstanceId(v string) *DeleteEntityRouteRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteEntityRouteRequest) SetUniqueId(v int64) *DeleteEntityRouteRequest {
	s.UniqueId = &v
	return s
}

type DeleteEntityRouteResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteEntityRouteResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteEntityRouteResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteEntityRouteResponseBody) SetCode(v string) *DeleteEntityRouteResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteEntityRouteResponseBody) SetMessage(v string) *DeleteEntityRouteResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteEntityRouteResponseBody) SetRequestId(v string) *DeleteEntityRouteResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteEntityRouteResponseBody) SetSuccess(v bool) *DeleteEntityRouteResponseBody {
	s.Success = &v
	return s
}

type DeleteEntityRouteResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteEntityRouteResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteEntityRouteResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteEntityRouteResponse) GoString() string {
	return s.String()
}

func (s *DeleteEntityRouteResponse) SetHeaders(v map[string]*string) *DeleteEntityRouteResponse {
	s.Headers = v
	return s
}

func (s *DeleteEntityRouteResponse) SetStatusCode(v int32) *DeleteEntityRouteResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteEntityRouteResponse) SetBody(v *DeleteEntityRouteResponseBody) *DeleteEntityRouteResponse {
	s.Body = v
	return s
}

type DisableRoleRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RoleId      *int64  `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
}

func (s DisableRoleRequest) String() string {
	return tea.Prettify(s)
}

func (s DisableRoleRequest) GoString() string {
	return s.String()
}

func (s *DisableRoleRequest) SetClientToken(v string) *DisableRoleRequest {
	s.ClientToken = &v
	return s
}

func (s *DisableRoleRequest) SetInstanceId(v string) *DisableRoleRequest {
	s.InstanceId = &v
	return s
}

func (s *DisableRoleRequest) SetRoleId(v int64) *DisableRoleRequest {
	s.RoleId = &v
	return s
}

type DisableRoleResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DisableRoleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableRoleResponseBody) GoString() string {
	return s.String()
}

func (s *DisableRoleResponseBody) SetCode(v string) *DisableRoleResponseBody {
	s.Code = &v
	return s
}

func (s *DisableRoleResponseBody) SetMessage(v string) *DisableRoleResponseBody {
	s.Message = &v
	return s
}

func (s *DisableRoleResponseBody) SetRequestId(v string) *DisableRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableRoleResponseBody) SetSuccess(v bool) *DisableRoleResponseBody {
	s.Success = &v
	return s
}

type DisableRoleResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DisableRoleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DisableRoleResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableRoleResponse) GoString() string {
	return s.String()
}

func (s *DisableRoleResponse) SetHeaders(v map[string]*string) *DisableRoleResponse {
	s.Headers = v
	return s
}

func (s *DisableRoleResponse) SetStatusCode(v int32) *DisableRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableRoleResponse) SetBody(v *DisableRoleResponseBody) *DisableRoleResponse {
	s.Body = v
	return s
}

type EditEntityRouteRequest struct {
	DepartmentId         *string `json:"DepartmentId,omitempty" xml:"DepartmentId,omitempty"`
	EntityBizCode        *string `json:"EntityBizCode,omitempty" xml:"EntityBizCode,omitempty"`
	EntityBizCodeType    *string `json:"EntityBizCodeType,omitempty" xml:"EntityBizCodeType,omitempty"`
	EntityId             *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	EntityName           *string `json:"EntityName,omitempty" xml:"EntityName,omitempty"`
	EntityRelationNumber *string `json:"EntityRelationNumber,omitempty" xml:"EntityRelationNumber,omitempty"`
	ExtInfo              *string `json:"ExtInfo,omitempty" xml:"ExtInfo,omitempty"`
	GroupId              *int64  `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ServiceId            *int64  `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	UniqueId             *int64  `json:"UniqueId,omitempty" xml:"UniqueId,omitempty"`
}

func (s EditEntityRouteRequest) String() string {
	return tea.Prettify(s)
}

func (s EditEntityRouteRequest) GoString() string {
	return s.String()
}

func (s *EditEntityRouteRequest) SetDepartmentId(v string) *EditEntityRouteRequest {
	s.DepartmentId = &v
	return s
}

func (s *EditEntityRouteRequest) SetEntityBizCode(v string) *EditEntityRouteRequest {
	s.EntityBizCode = &v
	return s
}

func (s *EditEntityRouteRequest) SetEntityBizCodeType(v string) *EditEntityRouteRequest {
	s.EntityBizCodeType = &v
	return s
}

func (s *EditEntityRouteRequest) SetEntityId(v string) *EditEntityRouteRequest {
	s.EntityId = &v
	return s
}

func (s *EditEntityRouteRequest) SetEntityName(v string) *EditEntityRouteRequest {
	s.EntityName = &v
	return s
}

func (s *EditEntityRouteRequest) SetEntityRelationNumber(v string) *EditEntityRouteRequest {
	s.EntityRelationNumber = &v
	return s
}

func (s *EditEntityRouteRequest) SetExtInfo(v string) *EditEntityRouteRequest {
	s.ExtInfo = &v
	return s
}

func (s *EditEntityRouteRequest) SetGroupId(v int64) *EditEntityRouteRequest {
	s.GroupId = &v
	return s
}

func (s *EditEntityRouteRequest) SetInstanceId(v string) *EditEntityRouteRequest {
	s.InstanceId = &v
	return s
}

func (s *EditEntityRouteRequest) SetServiceId(v int64) *EditEntityRouteRequest {
	s.ServiceId = &v
	return s
}

func (s *EditEntityRouteRequest) SetUniqueId(v int64) *EditEntityRouteRequest {
	s.UniqueId = &v
	return s
}

type EditEntityRouteResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EditEntityRouteResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EditEntityRouteResponseBody) GoString() string {
	return s.String()
}

func (s *EditEntityRouteResponseBody) SetCode(v string) *EditEntityRouteResponseBody {
	s.Code = &v
	return s
}

func (s *EditEntityRouteResponseBody) SetMessage(v string) *EditEntityRouteResponseBody {
	s.Message = &v
	return s
}

func (s *EditEntityRouteResponseBody) SetRequestId(v string) *EditEntityRouteResponseBody {
	s.RequestId = &v
	return s
}

func (s *EditEntityRouteResponseBody) SetSuccess(v bool) *EditEntityRouteResponseBody {
	s.Success = &v
	return s
}

type EditEntityRouteResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *EditEntityRouteResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EditEntityRouteResponse) String() string {
	return tea.Prettify(s)
}

func (s EditEntityRouteResponse) GoString() string {
	return s.String()
}

func (s *EditEntityRouteResponse) SetHeaders(v map[string]*string) *EditEntityRouteResponse {
	s.Headers = v
	return s
}

func (s *EditEntityRouteResponse) SetStatusCode(v int32) *EditEntityRouteResponse {
	s.StatusCode = &v
	return s
}

func (s *EditEntityRouteResponse) SetBody(v *EditEntityRouteResponseBody) *EditEntityRouteResponse {
	s.Body = v
	return s
}

type EnableRoleRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RoleId      *int64  `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
}

func (s EnableRoleRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableRoleRequest) GoString() string {
	return s.String()
}

func (s *EnableRoleRequest) SetClientToken(v string) *EnableRoleRequest {
	s.ClientToken = &v
	return s
}

func (s *EnableRoleRequest) SetInstanceId(v string) *EnableRoleRequest {
	s.InstanceId = &v
	return s
}

func (s *EnableRoleRequest) SetRoleId(v int64) *EnableRoleRequest {
	s.RoleId = &v
	return s
}

type EnableRoleResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnableRoleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableRoleResponseBody) GoString() string {
	return s.String()
}

func (s *EnableRoleResponseBody) SetCode(v string) *EnableRoleResponseBody {
	s.Code = &v
	return s
}

func (s *EnableRoleResponseBody) SetMessage(v string) *EnableRoleResponseBody {
	s.Message = &v
	return s
}

func (s *EnableRoleResponseBody) SetRequestId(v string) *EnableRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *EnableRoleResponseBody) SetSuccess(v bool) *EnableRoleResponseBody {
	s.Success = &v
	return s
}

type EnableRoleResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *EnableRoleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EnableRoleResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableRoleResponse) GoString() string {
	return s.String()
}

func (s *EnableRoleResponse) SetHeaders(v map[string]*string) *EnableRoleResponse {
	s.Headers = v
	return s
}

func (s *EnableRoleResponse) SetStatusCode(v int32) *EnableRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *EnableRoleResponse) SetBody(v *EnableRoleResponseBody) *EnableRoleResponse {
	s.Body = v
	return s
}

type ExecuteActivityRequest struct {
	ActivityCode *string `json:"ActivityCode,omitempty" xml:"ActivityCode,omitempty"`
	ActivityForm *string `json:"ActivityForm,omitempty" xml:"ActivityForm,omitempty"`
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OperatorId   *int64  `json:"OperatorId,omitempty" xml:"OperatorId,omitempty"`
	TicketId     *int64  `json:"TicketId,omitempty" xml:"TicketId,omitempty"`
}

func (s ExecuteActivityRequest) String() string {
	return tea.Prettify(s)
}

func (s ExecuteActivityRequest) GoString() string {
	return s.String()
}

func (s *ExecuteActivityRequest) SetActivityCode(v string) *ExecuteActivityRequest {
	s.ActivityCode = &v
	return s
}

func (s *ExecuteActivityRequest) SetActivityForm(v string) *ExecuteActivityRequest {
	s.ActivityForm = &v
	return s
}

func (s *ExecuteActivityRequest) SetClientToken(v string) *ExecuteActivityRequest {
	s.ClientToken = &v
	return s
}

func (s *ExecuteActivityRequest) SetInstanceId(v string) *ExecuteActivityRequest {
	s.InstanceId = &v
	return s
}

func (s *ExecuteActivityRequest) SetOperatorId(v int64) *ExecuteActivityRequest {
	s.OperatorId = &v
	return s
}

func (s *ExecuteActivityRequest) SetTicketId(v int64) *ExecuteActivityRequest {
	s.TicketId = &v
	return s
}

type ExecuteActivityResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ExecuteActivityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExecuteActivityResponseBody) GoString() string {
	return s.String()
}

func (s *ExecuteActivityResponseBody) SetCode(v string) *ExecuteActivityResponseBody {
	s.Code = &v
	return s
}

func (s *ExecuteActivityResponseBody) SetMessage(v string) *ExecuteActivityResponseBody {
	s.Message = &v
	return s
}

func (s *ExecuteActivityResponseBody) SetRequestId(v string) *ExecuteActivityResponseBody {
	s.RequestId = &v
	return s
}

func (s *ExecuteActivityResponseBody) SetSuccess(v bool) *ExecuteActivityResponseBody {
	s.Success = &v
	return s
}

type ExecuteActivityResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ExecuteActivityResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ExecuteActivityResponse) String() string {
	return tea.Prettify(s)
}

func (s ExecuteActivityResponse) GoString() string {
	return s.String()
}

func (s *ExecuteActivityResponse) SetHeaders(v map[string]*string) *ExecuteActivityResponse {
	s.Headers = v
	return s
}

func (s *ExecuteActivityResponse) SetStatusCode(v int32) *ExecuteActivityResponse {
	s.StatusCode = &v
	return s
}

func (s *ExecuteActivityResponse) SetBody(v *ExecuteActivityResponseBody) *ExecuteActivityResponse {
	s.Body = v
	return s
}

type FetchCallRequest struct {
	AccountName      *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	CallId           *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	ClientToken      *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ConnectionId     *string `json:"ConnectionId,omitempty" xml:"ConnectionId,omitempty"`
	HoldConnectionId *string `json:"HoldConnectionId,omitempty" xml:"HoldConnectionId,omitempty"`
	InstanceId       *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	JobId            *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s FetchCallRequest) String() string {
	return tea.Prettify(s)
}

func (s FetchCallRequest) GoString() string {
	return s.String()
}

func (s *FetchCallRequest) SetAccountName(v string) *FetchCallRequest {
	s.AccountName = &v
	return s
}

func (s *FetchCallRequest) SetCallId(v string) *FetchCallRequest {
	s.CallId = &v
	return s
}

func (s *FetchCallRequest) SetClientToken(v string) *FetchCallRequest {
	s.ClientToken = &v
	return s
}

func (s *FetchCallRequest) SetConnectionId(v string) *FetchCallRequest {
	s.ConnectionId = &v
	return s
}

func (s *FetchCallRequest) SetHoldConnectionId(v string) *FetchCallRequest {
	s.HoldConnectionId = &v
	return s
}

func (s *FetchCallRequest) SetInstanceId(v string) *FetchCallRequest {
	s.InstanceId = &v
	return s
}

func (s *FetchCallRequest) SetJobId(v string) *FetchCallRequest {
	s.JobId = &v
	return s
}

type FetchCallResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s FetchCallResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FetchCallResponseBody) GoString() string {
	return s.String()
}

func (s *FetchCallResponseBody) SetCode(v string) *FetchCallResponseBody {
	s.Code = &v
	return s
}

func (s *FetchCallResponseBody) SetMessage(v string) *FetchCallResponseBody {
	s.Message = &v
	return s
}

func (s *FetchCallResponseBody) SetRequestId(v string) *FetchCallResponseBody {
	s.RequestId = &v
	return s
}

func (s *FetchCallResponseBody) SetSuccess(v bool) *FetchCallResponseBody {
	s.Success = &v
	return s
}

type FetchCallResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *FetchCallResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s FetchCallResponse) String() string {
	return tea.Prettify(s)
}

func (s FetchCallResponse) GoString() string {
	return s.String()
}

func (s *FetchCallResponse) SetHeaders(v map[string]*string) *FetchCallResponse {
	s.Headers = v
	return s
}

func (s *FetchCallResponse) SetStatusCode(v int32) *FetchCallResponse {
	s.StatusCode = &v
	return s
}

func (s *FetchCallResponse) SetBody(v *FetchCallResponseBody) *FetchCallResponse {
	s.Body = v
	return s
}

type FinishHotlineServiceRequest struct {
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s FinishHotlineServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s FinishHotlineServiceRequest) GoString() string {
	return s.String()
}

func (s *FinishHotlineServiceRequest) SetAccountName(v string) *FinishHotlineServiceRequest {
	s.AccountName = &v
	return s
}

func (s *FinishHotlineServiceRequest) SetClientToken(v string) *FinishHotlineServiceRequest {
	s.ClientToken = &v
	return s
}

func (s *FinishHotlineServiceRequest) SetInstanceId(v string) *FinishHotlineServiceRequest {
	s.InstanceId = &v
	return s
}

type FinishHotlineServiceResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s FinishHotlineServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FinishHotlineServiceResponseBody) GoString() string {
	return s.String()
}

func (s *FinishHotlineServiceResponseBody) SetCode(v string) *FinishHotlineServiceResponseBody {
	s.Code = &v
	return s
}

func (s *FinishHotlineServiceResponseBody) SetMessage(v string) *FinishHotlineServiceResponseBody {
	s.Message = &v
	return s
}

func (s *FinishHotlineServiceResponseBody) SetRequestId(v string) *FinishHotlineServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *FinishHotlineServiceResponseBody) SetSuccess(v bool) *FinishHotlineServiceResponseBody {
	s.Success = &v
	return s
}

type FinishHotlineServiceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *FinishHotlineServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s FinishHotlineServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s FinishHotlineServiceResponse) GoString() string {
	return s.String()
}

func (s *FinishHotlineServiceResponse) SetHeaders(v map[string]*string) *FinishHotlineServiceResponse {
	s.Headers = v
	return s
}

func (s *FinishHotlineServiceResponse) SetStatusCode(v int32) *FinishHotlineServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *FinishHotlineServiceResponse) SetBody(v *FinishHotlineServiceResponseBody) *FinishHotlineServiceResponse {
	s.Body = v
	return s
}

type FinishHotlineServiceNewRequest struct {
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s FinishHotlineServiceNewRequest) String() string {
	return tea.Prettify(s)
}

func (s FinishHotlineServiceNewRequest) GoString() string {
	return s.String()
}

func (s *FinishHotlineServiceNewRequest) SetAccountName(v string) *FinishHotlineServiceNewRequest {
	s.AccountName = &v
	return s
}

func (s *FinishHotlineServiceNewRequest) SetClientToken(v string) *FinishHotlineServiceNewRequest {
	s.ClientToken = &v
	return s
}

func (s *FinishHotlineServiceNewRequest) SetInstanceId(v string) *FinishHotlineServiceNewRequest {
	s.InstanceId = &v
	return s
}

type FinishHotlineServiceNewResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s FinishHotlineServiceNewResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FinishHotlineServiceNewResponseBody) GoString() string {
	return s.String()
}

func (s *FinishHotlineServiceNewResponseBody) SetCode(v string) *FinishHotlineServiceNewResponseBody {
	s.Code = &v
	return s
}

func (s *FinishHotlineServiceNewResponseBody) SetMessage(v string) *FinishHotlineServiceNewResponseBody {
	s.Message = &v
	return s
}

func (s *FinishHotlineServiceNewResponseBody) SetRequestId(v string) *FinishHotlineServiceNewResponseBody {
	s.RequestId = &v
	return s
}

func (s *FinishHotlineServiceNewResponseBody) SetSuccess(v bool) *FinishHotlineServiceNewResponseBody {
	s.Success = &v
	return s
}

type FinishHotlineServiceNewResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *FinishHotlineServiceNewResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s FinishHotlineServiceNewResponse) String() string {
	return tea.Prettify(s)
}

func (s FinishHotlineServiceNewResponse) GoString() string {
	return s.String()
}

func (s *FinishHotlineServiceNewResponse) SetHeaders(v map[string]*string) *FinishHotlineServiceNewResponse {
	s.Headers = v
	return s
}

func (s *FinishHotlineServiceNewResponse) SetStatusCode(v int32) *FinishHotlineServiceNewResponse {
	s.StatusCode = &v
	return s
}

func (s *FinishHotlineServiceNewResponse) SetBody(v *FinishHotlineServiceNewResponseBody) *FinishHotlineServiceNewResponse {
	s.Body = v
	return s
}

type GenerateWebSocketSignRequest struct {
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GenerateWebSocketSignRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateWebSocketSignRequest) GoString() string {
	return s.String()
}

func (s *GenerateWebSocketSignRequest) SetAccountName(v string) *GenerateWebSocketSignRequest {
	s.AccountName = &v
	return s
}

func (s *GenerateWebSocketSignRequest) SetClientToken(v string) *GenerateWebSocketSignRequest {
	s.ClientToken = &v
	return s
}

func (s *GenerateWebSocketSignRequest) SetInstanceId(v string) *GenerateWebSocketSignRequest {
	s.InstanceId = &v
	return s
}

type GenerateWebSocketSignResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *string `json:"Data,omitempty" xml:"Data,omitempty"`
	HttpStatusCode *int64  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GenerateWebSocketSignResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GenerateWebSocketSignResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateWebSocketSignResponseBody) SetCode(v string) *GenerateWebSocketSignResponseBody {
	s.Code = &v
	return s
}

func (s *GenerateWebSocketSignResponseBody) SetData(v string) *GenerateWebSocketSignResponseBody {
	s.Data = &v
	return s
}

func (s *GenerateWebSocketSignResponseBody) SetHttpStatusCode(v int64) *GenerateWebSocketSignResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GenerateWebSocketSignResponseBody) SetMessage(v string) *GenerateWebSocketSignResponseBody {
	s.Message = &v
	return s
}

func (s *GenerateWebSocketSignResponseBody) SetRequestId(v string) *GenerateWebSocketSignResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateWebSocketSignResponseBody) SetSuccess(v bool) *GenerateWebSocketSignResponseBody {
	s.Success = &v
	return s
}

type GenerateWebSocketSignResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GenerateWebSocketSignResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GenerateWebSocketSignResponse) String() string {
	return tea.Prettify(s)
}

func (s GenerateWebSocketSignResponse) GoString() string {
	return s.String()
}

func (s *GenerateWebSocketSignResponse) SetHeaders(v map[string]*string) *GenerateWebSocketSignResponse {
	s.Headers = v
	return s
}

func (s *GenerateWebSocketSignResponse) SetStatusCode(v int32) *GenerateWebSocketSignResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateWebSocketSignResponse) SetBody(v *GenerateWebSocketSignResponseBody) *GenerateWebSocketSignResponse {
	s.Body = v
	return s
}

type GetAgentRequest struct {
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetAgentRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAgentRequest) GoString() string {
	return s.String()
}

func (s *GetAgentRequest) SetAccountName(v string) *GetAgentRequest {
	s.AccountName = &v
	return s
}

func (s *GetAgentRequest) SetClientToken(v string) *GetAgentRequest {
	s.ClientToken = &v
	return s
}

func (s *GetAgentRequest) SetInstanceId(v string) *GetAgentRequest {
	s.InstanceId = &v
	return s
}

type GetAgentResponseBody struct {
	Code           *string                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *GetAgentResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int64                    `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                     `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAgentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAgentResponseBody) GoString() string {
	return s.String()
}

func (s *GetAgentResponseBody) SetCode(v string) *GetAgentResponseBody {
	s.Code = &v
	return s
}

func (s *GetAgentResponseBody) SetData(v *GetAgentResponseBodyData) *GetAgentResponseBody {
	s.Data = v
	return s
}

func (s *GetAgentResponseBody) SetHttpStatusCode(v int64) *GetAgentResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetAgentResponseBody) SetMessage(v string) *GetAgentResponseBody {
	s.Message = &v
	return s
}

func (s *GetAgentResponseBody) SetRequestId(v string) *GetAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAgentResponseBody) SetSuccess(v bool) *GetAgentResponseBody {
	s.Success = &v
	return s
}

type GetAgentResponseBodyData struct {
	AccountName *string                              `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	AgentId     *int64                               `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	DisplayName *string                              `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	GroupList   []*GetAgentResponseBodyDataGroupList `json:"GroupList,omitempty" xml:"GroupList,omitempty" type:"Repeated"`
	Status      *int32                               `json:"Status,omitempty" xml:"Status,omitempty"`
	TenantId    *int64                               `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s GetAgentResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetAgentResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAgentResponseBodyData) SetAccountName(v string) *GetAgentResponseBodyData {
	s.AccountName = &v
	return s
}

func (s *GetAgentResponseBodyData) SetAgentId(v int64) *GetAgentResponseBodyData {
	s.AgentId = &v
	return s
}

func (s *GetAgentResponseBodyData) SetDisplayName(v string) *GetAgentResponseBodyData {
	s.DisplayName = &v
	return s
}

func (s *GetAgentResponseBodyData) SetGroupList(v []*GetAgentResponseBodyDataGroupList) *GetAgentResponseBodyData {
	s.GroupList = v
	return s
}

func (s *GetAgentResponseBodyData) SetStatus(v int32) *GetAgentResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetAgentResponseBodyData) SetTenantId(v int64) *GetAgentResponseBodyData {
	s.TenantId = &v
	return s
}

type GetAgentResponseBodyDataGroupList struct {
	ChannelType  *int32  `json:"ChannelType,omitempty" xml:"ChannelType,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DisplayName  *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	SkillGroupId *int64  `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
}

func (s GetAgentResponseBodyDataGroupList) String() string {
	return tea.Prettify(s)
}

func (s GetAgentResponseBodyDataGroupList) GoString() string {
	return s.String()
}

func (s *GetAgentResponseBodyDataGroupList) SetChannelType(v int32) *GetAgentResponseBodyDataGroupList {
	s.ChannelType = &v
	return s
}

func (s *GetAgentResponseBodyDataGroupList) SetDescription(v string) *GetAgentResponseBodyDataGroupList {
	s.Description = &v
	return s
}

func (s *GetAgentResponseBodyDataGroupList) SetDisplayName(v string) *GetAgentResponseBodyDataGroupList {
	s.DisplayName = &v
	return s
}

func (s *GetAgentResponseBodyDataGroupList) SetName(v string) *GetAgentResponseBodyDataGroupList {
	s.Name = &v
	return s
}

func (s *GetAgentResponseBodyDataGroupList) SetSkillGroupId(v int64) *GetAgentResponseBodyDataGroupList {
	s.SkillGroupId = &v
	return s
}

type GetAgentResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetAgentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAgentResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAgentResponse) GoString() string {
	return s.String()
}

func (s *GetAgentResponse) SetHeaders(v map[string]*string) *GetAgentResponse {
	s.Headers = v
	return s
}

func (s *GetAgentResponse) SetStatusCode(v int32) *GetAgentResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAgentResponse) SetBody(v *GetAgentResponseBody) *GetAgentResponse {
	s.Body = v
	return s
}

type GetAgentHotlineRequest struct {
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetAgentHotlineRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAgentHotlineRequest) GoString() string {
	return s.String()
}

func (s *GetAgentHotlineRequest) SetAccountName(v string) *GetAgentHotlineRequest {
	s.AccountName = &v
	return s
}

func (s *GetAgentHotlineRequest) SetClientToken(v string) *GetAgentHotlineRequest {
	s.ClientToken = &v
	return s
}

func (s *GetAgentHotlineRequest) SetInstanceId(v string) *GetAgentHotlineRequest {
	s.InstanceId = &v
	return s
}

type GetAgentHotlineResponseBody struct {
	Code      *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetAgentHotlineResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                            `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAgentHotlineResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAgentHotlineResponseBody) GoString() string {
	return s.String()
}

func (s *GetAgentHotlineResponseBody) SetCode(v string) *GetAgentHotlineResponseBody {
	s.Code = &v
	return s
}

func (s *GetAgentHotlineResponseBody) SetData(v *GetAgentHotlineResponseBodyData) *GetAgentHotlineResponseBody {
	s.Data = v
	return s
}

func (s *GetAgentHotlineResponseBody) SetMessage(v string) *GetAgentHotlineResponseBody {
	s.Message = &v
	return s
}

func (s *GetAgentHotlineResponseBody) SetRequestId(v string) *GetAgentHotlineResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAgentHotlineResponseBody) SetSuccess(v bool) *GetAgentHotlineResponseBody {
	s.Success = &v
	return s
}

type GetAgentHotlineResponseBodyData struct {
	AccountName  *string                                     `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	AgentId      *int64                                      `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	AgentPhone   *string                                     `json:"AgentPhone,omitempty" xml:"AgentPhone,omitempty"`
	DepartmentId *string                                     `json:"DepartmentId,omitempty" xml:"DepartmentId,omitempty"`
	DisplayName  *string                                     `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	ExtAttr      *string                                     `json:"ExtAttr,omitempty" xml:"ExtAttr,omitempty"`
	GroupList    []*GetAgentHotlineResponseBodyDataGroupList `json:"GroupList,omitempty" xml:"GroupList,omitempty" type:"Repeated"`
	SkillGroups  *string                                     `json:"SkillGroups,omitempty" xml:"SkillGroups,omitempty"`
	TenantId     *int64                                      `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s GetAgentHotlineResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetAgentHotlineResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAgentHotlineResponseBodyData) SetAccountName(v string) *GetAgentHotlineResponseBodyData {
	s.AccountName = &v
	return s
}

func (s *GetAgentHotlineResponseBodyData) SetAgentId(v int64) *GetAgentHotlineResponseBodyData {
	s.AgentId = &v
	return s
}

func (s *GetAgentHotlineResponseBodyData) SetAgentPhone(v string) *GetAgentHotlineResponseBodyData {
	s.AgentPhone = &v
	return s
}

func (s *GetAgentHotlineResponseBodyData) SetDepartmentId(v string) *GetAgentHotlineResponseBodyData {
	s.DepartmentId = &v
	return s
}

func (s *GetAgentHotlineResponseBodyData) SetDisplayName(v string) *GetAgentHotlineResponseBodyData {
	s.DisplayName = &v
	return s
}

func (s *GetAgentHotlineResponseBodyData) SetExtAttr(v string) *GetAgentHotlineResponseBodyData {
	s.ExtAttr = &v
	return s
}

func (s *GetAgentHotlineResponseBodyData) SetGroupList(v []*GetAgentHotlineResponseBodyDataGroupList) *GetAgentHotlineResponseBodyData {
	s.GroupList = v
	return s
}

func (s *GetAgentHotlineResponseBodyData) SetSkillGroups(v string) *GetAgentHotlineResponseBodyData {
	s.SkillGroups = &v
	return s
}

func (s *GetAgentHotlineResponseBodyData) SetTenantId(v int64) *GetAgentHotlineResponseBodyData {
	s.TenantId = &v
	return s
}

type GetAgentHotlineResponseBodyDataGroupList struct {
	DisplayName  *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	SkillGroupId *int64  `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
}

func (s GetAgentHotlineResponseBodyDataGroupList) String() string {
	return tea.Prettify(s)
}

func (s GetAgentHotlineResponseBodyDataGroupList) GoString() string {
	return s.String()
}

func (s *GetAgentHotlineResponseBodyDataGroupList) SetDisplayName(v string) *GetAgentHotlineResponseBodyDataGroupList {
	s.DisplayName = &v
	return s
}

func (s *GetAgentHotlineResponseBodyDataGroupList) SetName(v string) *GetAgentHotlineResponseBodyDataGroupList {
	s.Name = &v
	return s
}

func (s *GetAgentHotlineResponseBodyDataGroupList) SetSkillGroupId(v int64) *GetAgentHotlineResponseBodyDataGroupList {
	s.SkillGroupId = &v
	return s
}

type GetAgentHotlineResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetAgentHotlineResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAgentHotlineResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAgentHotlineResponse) GoString() string {
	return s.String()
}

func (s *GetAgentHotlineResponse) SetHeaders(v map[string]*string) *GetAgentHotlineResponse {
	s.Headers = v
	return s
}

func (s *GetAgentHotlineResponse) SetStatusCode(v int32) *GetAgentHotlineResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAgentHotlineResponse) SetBody(v *GetAgentHotlineResponseBody) *GetAgentHotlineResponse {
	s.Body = v
	return s
}

type GetAgentHotlinePhoneRequest struct {
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetAgentHotlinePhoneRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAgentHotlinePhoneRequest) GoString() string {
	return s.String()
}

func (s *GetAgentHotlinePhoneRequest) SetAccountName(v string) *GetAgentHotlinePhoneRequest {
	s.AccountName = &v
	return s
}

func (s *GetAgentHotlinePhoneRequest) SetClientToken(v string) *GetAgentHotlinePhoneRequest {
	s.ClientToken = &v
	return s
}

func (s *GetAgentHotlinePhoneRequest) SetInstanceId(v string) *GetAgentHotlinePhoneRequest {
	s.InstanceId = &v
	return s
}

type GetAgentHotlinePhoneResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAgentHotlinePhoneResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAgentHotlinePhoneResponseBody) GoString() string {
	return s.String()
}

func (s *GetAgentHotlinePhoneResponseBody) SetCode(v string) *GetAgentHotlinePhoneResponseBody {
	s.Code = &v
	return s
}

func (s *GetAgentHotlinePhoneResponseBody) SetData(v string) *GetAgentHotlinePhoneResponseBody {
	s.Data = &v
	return s
}

func (s *GetAgentHotlinePhoneResponseBody) SetMessage(v string) *GetAgentHotlinePhoneResponseBody {
	s.Message = &v
	return s
}

func (s *GetAgentHotlinePhoneResponseBody) SetRequestId(v string) *GetAgentHotlinePhoneResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAgentHotlinePhoneResponseBody) SetSuccess(v bool) *GetAgentHotlinePhoneResponseBody {
	s.Success = &v
	return s
}

type GetAgentHotlinePhoneResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetAgentHotlinePhoneResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAgentHotlinePhoneResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAgentHotlinePhoneResponse) GoString() string {
	return s.String()
}

func (s *GetAgentHotlinePhoneResponse) SetHeaders(v map[string]*string) *GetAgentHotlinePhoneResponse {
	s.Headers = v
	return s
}

func (s *GetAgentHotlinePhoneResponse) SetStatusCode(v int32) *GetAgentHotlinePhoneResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAgentHotlinePhoneResponse) SetBody(v *GetAgentHotlinePhoneResponseBody) *GetAgentHotlinePhoneResponse {
	s.Body = v
	return s
}

type GetAgentWorkStatusRequest struct {
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetAgentWorkStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAgentWorkStatusRequest) GoString() string {
	return s.String()
}

func (s *GetAgentWorkStatusRequest) SetAccountName(v string) *GetAgentWorkStatusRequest {
	s.AccountName = &v
	return s
}

func (s *GetAgentWorkStatusRequest) SetClientToken(v string) *GetAgentWorkStatusRequest {
	s.ClientToken = &v
	return s
}

func (s *GetAgentWorkStatusRequest) SetInstanceId(v string) *GetAgentWorkStatusRequest {
	s.InstanceId = &v
	return s
}

type GetAgentWorkStatusResponseBody struct {
	Code      *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetAgentWorkStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAgentWorkStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAgentWorkStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetAgentWorkStatusResponseBody) SetCode(v string) *GetAgentWorkStatusResponseBody {
	s.Code = &v
	return s
}

func (s *GetAgentWorkStatusResponseBody) SetData(v *GetAgentWorkStatusResponseBodyData) *GetAgentWorkStatusResponseBody {
	s.Data = v
	return s
}

func (s *GetAgentWorkStatusResponseBody) SetMessage(v string) *GetAgentWorkStatusResponseBody {
	s.Message = &v
	return s
}

func (s *GetAgentWorkStatusResponseBody) SetRequestId(v string) *GetAgentWorkStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAgentWorkStatusResponseBody) SetSuccess(v bool) *GetAgentWorkStatusResponseBody {
	s.Success = &v
	return s
}

type GetAgentWorkStatusResponseBodyData struct {
	AccountName     *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	AgentId         *int64  `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	AgentStatusCode *string `json:"AgentStatusCode,omitempty" xml:"AgentStatusCode,omitempty"`
	AgentStatusDesc *string `json:"AgentStatusDesc,omitempty" xml:"AgentStatusDesc,omitempty"`
	ExtAttr         *string `json:"ExtAttr,omitempty" xml:"ExtAttr,omitempty"`
	TenantId        *int64  `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	WorkChannel     *string `json:"WorkChannel,omitempty" xml:"WorkChannel,omitempty"`
}

func (s GetAgentWorkStatusResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetAgentWorkStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAgentWorkStatusResponseBodyData) SetAccountName(v string) *GetAgentWorkStatusResponseBodyData {
	s.AccountName = &v
	return s
}

func (s *GetAgentWorkStatusResponseBodyData) SetAgentId(v int64) *GetAgentWorkStatusResponseBodyData {
	s.AgentId = &v
	return s
}

func (s *GetAgentWorkStatusResponseBodyData) SetAgentStatusCode(v string) *GetAgentWorkStatusResponseBodyData {
	s.AgentStatusCode = &v
	return s
}

func (s *GetAgentWorkStatusResponseBodyData) SetAgentStatusDesc(v string) *GetAgentWorkStatusResponseBodyData {
	s.AgentStatusDesc = &v
	return s
}

func (s *GetAgentWorkStatusResponseBodyData) SetExtAttr(v string) *GetAgentWorkStatusResponseBodyData {
	s.ExtAttr = &v
	return s
}

func (s *GetAgentWorkStatusResponseBodyData) SetTenantId(v int64) *GetAgentWorkStatusResponseBodyData {
	s.TenantId = &v
	return s
}

func (s *GetAgentWorkStatusResponseBodyData) SetWorkChannel(v string) *GetAgentWorkStatusResponseBodyData {
	s.WorkChannel = &v
	return s
}

type GetAgentWorkStatusResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetAgentWorkStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAgentWorkStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAgentWorkStatusResponse) GoString() string {
	return s.String()
}

func (s *GetAgentWorkStatusResponse) SetHeaders(v map[string]*string) *GetAgentWorkStatusResponse {
	s.Headers = v
	return s
}

func (s *GetAgentWorkStatusResponse) SetStatusCode(v int32) *GetAgentWorkStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAgentWorkStatusResponse) SetBody(v *GetAgentWorkStatusResponseBody) *GetAgentWorkStatusResponse {
	s.Body = v
	return s
}

type GetAllDepartmentRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetAllDepartmentRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAllDepartmentRequest) GoString() string {
	return s.String()
}

func (s *GetAllDepartmentRequest) SetInstanceId(v string) *GetAllDepartmentRequest {
	s.InstanceId = &v
	return s
}

type GetAllDepartmentResponseBody struct {
	Code      *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*GetAllDepartmentResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAllDepartmentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAllDepartmentResponseBody) GoString() string {
	return s.String()
}

func (s *GetAllDepartmentResponseBody) SetCode(v string) *GetAllDepartmentResponseBody {
	s.Code = &v
	return s
}

func (s *GetAllDepartmentResponseBody) SetData(v []*GetAllDepartmentResponseBodyData) *GetAllDepartmentResponseBody {
	s.Data = v
	return s
}

func (s *GetAllDepartmentResponseBody) SetMessage(v string) *GetAllDepartmentResponseBody {
	s.Message = &v
	return s
}

func (s *GetAllDepartmentResponseBody) SetRequestId(v string) *GetAllDepartmentResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAllDepartmentResponseBody) SetSuccess(v bool) *GetAllDepartmentResponseBody {
	s.Success = &v
	return s
}

type GetAllDepartmentResponseBodyData struct {
	DepartmentId   *int64  `json:"DepartmentId,omitempty" xml:"DepartmentId,omitempty"`
	DepartmentName *string `json:"DepartmentName,omitempty" xml:"DepartmentName,omitempty"`
	Status         *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetAllDepartmentResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetAllDepartmentResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAllDepartmentResponseBodyData) SetDepartmentId(v int64) *GetAllDepartmentResponseBodyData {
	s.DepartmentId = &v
	return s
}

func (s *GetAllDepartmentResponseBodyData) SetDepartmentName(v string) *GetAllDepartmentResponseBodyData {
	s.DepartmentName = &v
	return s
}

func (s *GetAllDepartmentResponseBodyData) SetStatus(v int32) *GetAllDepartmentResponseBodyData {
	s.Status = &v
	return s
}

type GetAllDepartmentResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetAllDepartmentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAllDepartmentResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAllDepartmentResponse) GoString() string {
	return s.String()
}

func (s *GetAllDepartmentResponse) SetHeaders(v map[string]*string) *GetAllDepartmentResponse {
	s.Headers = v
	return s
}

func (s *GetAllDepartmentResponse) SetStatusCode(v int32) *GetAllDepartmentResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAllDepartmentResponse) SetBody(v *GetAllDepartmentResponseBody) *GetAllDepartmentResponse {
	s.Body = v
	return s
}

type GetAuthInfoRequest struct {
	AppKey     *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	ForeignId  *string `json:"ForeignId,omitempty" xml:"ForeignId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetAuthInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAuthInfoRequest) GoString() string {
	return s.String()
}

func (s *GetAuthInfoRequest) SetAppKey(v string) *GetAuthInfoRequest {
	s.AppKey = &v
	return s
}

func (s *GetAuthInfoRequest) SetForeignId(v string) *GetAuthInfoRequest {
	s.ForeignId = &v
	return s
}

func (s *GetAuthInfoRequest) SetInstanceId(v string) *GetAuthInfoRequest {
	s.InstanceId = &v
	return s
}

type GetAuthInfoResponseBody struct {
	Code      *string                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetAuthInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                        `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAuthInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAuthInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetAuthInfoResponseBody) SetCode(v string) *GetAuthInfoResponseBody {
	s.Code = &v
	return s
}

func (s *GetAuthInfoResponseBody) SetData(v *GetAuthInfoResponseBodyData) *GetAuthInfoResponseBody {
	s.Data = v
	return s
}

func (s *GetAuthInfoResponseBody) SetMessage(v string) *GetAuthInfoResponseBody {
	s.Message = &v
	return s
}

func (s *GetAuthInfoResponseBody) SetRequestId(v string) *GetAuthInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAuthInfoResponseBody) SetSuccess(v bool) *GetAuthInfoResponseBody {
	s.Success = &v
	return s
}

type GetAuthInfoResponseBodyData struct {
	App       *string `json:"App,omitempty" xml:"App,omitempty"`
	AppKey    *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	AppName   *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	TenantId  *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	Time      *int64  `json:"Time,omitempty" xml:"Time,omitempty"`
	UserId    *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserName  *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s GetAuthInfoResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetAuthInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAuthInfoResponseBodyData) SetApp(v string) *GetAuthInfoResponseBodyData {
	s.App = &v
	return s
}

func (s *GetAuthInfoResponseBodyData) SetAppKey(v string) *GetAuthInfoResponseBodyData {
	s.AppKey = &v
	return s
}

func (s *GetAuthInfoResponseBodyData) SetAppName(v string) *GetAuthInfoResponseBodyData {
	s.AppName = &v
	return s
}

func (s *GetAuthInfoResponseBodyData) SetCode(v string) *GetAuthInfoResponseBodyData {
	s.Code = &v
	return s
}

func (s *GetAuthInfoResponseBodyData) SetSessionId(v string) *GetAuthInfoResponseBodyData {
	s.SessionId = &v
	return s
}

func (s *GetAuthInfoResponseBodyData) SetTenantId(v string) *GetAuthInfoResponseBodyData {
	s.TenantId = &v
	return s
}

func (s *GetAuthInfoResponseBodyData) SetTime(v int64) *GetAuthInfoResponseBodyData {
	s.Time = &v
	return s
}

func (s *GetAuthInfoResponseBodyData) SetUserId(v string) *GetAuthInfoResponseBodyData {
	s.UserId = &v
	return s
}

func (s *GetAuthInfoResponseBodyData) SetUserName(v string) *GetAuthInfoResponseBodyData {
	s.UserName = &v
	return s
}

type GetAuthInfoResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetAuthInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAuthInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAuthInfoResponse) GoString() string {
	return s.String()
}

func (s *GetAuthInfoResponse) SetHeaders(v map[string]*string) *GetAuthInfoResponse {
	s.Headers = v
	return s
}

func (s *GetAuthInfoResponse) SetStatusCode(v int32) *GetAuthInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAuthInfoResponse) SetBody(v *GetAuthInfoResponseBody) *GetAuthInfoResponse {
	s.Body = v
	return s
}

type GetByForeignIdRequest struct {
	ForeignId  *string `json:"ForeignId,omitempty" xml:"ForeignId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	SourceId   *int64  `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
}

func (s GetByForeignIdRequest) String() string {
	return tea.Prettify(s)
}

func (s GetByForeignIdRequest) GoString() string {
	return s.String()
}

func (s *GetByForeignIdRequest) SetForeignId(v string) *GetByForeignIdRequest {
	s.ForeignId = &v
	return s
}

func (s *GetByForeignIdRequest) SetInstanceId(v string) *GetByForeignIdRequest {
	s.InstanceId = &v
	return s
}

func (s *GetByForeignIdRequest) SetSourceId(v int64) *GetByForeignIdRequest {
	s.SourceId = &v
	return s
}

type GetByForeignIdResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetByForeignIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetByForeignIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetByForeignIdResponseBody) SetCode(v string) *GetByForeignIdResponseBody {
	s.Code = &v
	return s
}

func (s *GetByForeignIdResponseBody) SetData(v string) *GetByForeignIdResponseBody {
	s.Data = &v
	return s
}

func (s *GetByForeignIdResponseBody) SetMessage(v string) *GetByForeignIdResponseBody {
	s.Message = &v
	return s
}

func (s *GetByForeignIdResponseBody) SetRequestId(v string) *GetByForeignIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetByForeignIdResponseBody) SetSuccess(v bool) *GetByForeignIdResponseBody {
	s.Success = &v
	return s
}

type GetByForeignIdResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetByForeignIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetByForeignIdResponse) String() string {
	return tea.Prettify(s)
}

func (s GetByForeignIdResponse) GoString() string {
	return s.String()
}

func (s *GetByForeignIdResponse) SetHeaders(v map[string]*string) *GetByForeignIdResponse {
	s.Headers = v
	return s
}

func (s *GetByForeignIdResponse) SetStatusCode(v int32) *GetByForeignIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetByForeignIdResponse) SetBody(v *GetByForeignIdResponseBody) *GetByForeignIdResponse {
	s.Body = v
	return s
}

type GetCMSIdByForeignIdRequest struct {
	ForeignId  *string `json:"ForeignId,omitempty" xml:"ForeignId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Nick       *string `json:"Nick,omitempty" xml:"Nick,omitempty"`
	SourceId   *int64  `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
}

func (s GetCMSIdByForeignIdRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCMSIdByForeignIdRequest) GoString() string {
	return s.String()
}

func (s *GetCMSIdByForeignIdRequest) SetForeignId(v string) *GetCMSIdByForeignIdRequest {
	s.ForeignId = &v
	return s
}

func (s *GetCMSIdByForeignIdRequest) SetInstanceId(v string) *GetCMSIdByForeignIdRequest {
	s.InstanceId = &v
	return s
}

func (s *GetCMSIdByForeignIdRequest) SetNick(v string) *GetCMSIdByForeignIdRequest {
	s.Nick = &v
	return s
}

func (s *GetCMSIdByForeignIdRequest) SetSourceId(v int64) *GetCMSIdByForeignIdRequest {
	s.SourceId = &v
	return s
}

type GetCMSIdByForeignIdResponseBody struct {
	Code      *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetCMSIdByForeignIdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetCMSIdByForeignIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCMSIdByForeignIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetCMSIdByForeignIdResponseBody) SetCode(v string) *GetCMSIdByForeignIdResponseBody {
	s.Code = &v
	return s
}

func (s *GetCMSIdByForeignIdResponseBody) SetData(v *GetCMSIdByForeignIdResponseBodyData) *GetCMSIdByForeignIdResponseBody {
	s.Data = v
	return s
}

func (s *GetCMSIdByForeignIdResponseBody) SetMessage(v string) *GetCMSIdByForeignIdResponseBody {
	s.Message = &v
	return s
}

func (s *GetCMSIdByForeignIdResponseBody) SetRequestId(v string) *GetCMSIdByForeignIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCMSIdByForeignIdResponseBody) SetSuccess(v bool) *GetCMSIdByForeignIdResponseBody {
	s.Success = &v
	return s
}

type GetCMSIdByForeignIdResponseBodyData struct {
	Anonymity      *bool   `json:"Anonymity,omitempty" xml:"Anonymity,omitempty"`
	Avatar         *string `json:"Avatar,omitempty" xml:"Avatar,omitempty"`
	CustomerTypeId *int32  `json:"CustomerTypeId,omitempty" xml:"CustomerTypeId,omitempty"`
	ForeignId      *string `json:"ForeignId,omitempty" xml:"ForeignId,omitempty"`
	Gender         *string `json:"Gender,omitempty" xml:"Gender,omitempty"`
	Nick           *string `json:"Nick,omitempty" xml:"Nick,omitempty"`
	Phone          *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
	Status         *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	UserId         *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetCMSIdByForeignIdResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetCMSIdByForeignIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCMSIdByForeignIdResponseBodyData) SetAnonymity(v bool) *GetCMSIdByForeignIdResponseBodyData {
	s.Anonymity = &v
	return s
}

func (s *GetCMSIdByForeignIdResponseBodyData) SetAvatar(v string) *GetCMSIdByForeignIdResponseBodyData {
	s.Avatar = &v
	return s
}

func (s *GetCMSIdByForeignIdResponseBodyData) SetCustomerTypeId(v int32) *GetCMSIdByForeignIdResponseBodyData {
	s.CustomerTypeId = &v
	return s
}

func (s *GetCMSIdByForeignIdResponseBodyData) SetForeignId(v string) *GetCMSIdByForeignIdResponseBodyData {
	s.ForeignId = &v
	return s
}

func (s *GetCMSIdByForeignIdResponseBodyData) SetGender(v string) *GetCMSIdByForeignIdResponseBodyData {
	s.Gender = &v
	return s
}

func (s *GetCMSIdByForeignIdResponseBodyData) SetNick(v string) *GetCMSIdByForeignIdResponseBodyData {
	s.Nick = &v
	return s
}

func (s *GetCMSIdByForeignIdResponseBodyData) SetPhone(v string) *GetCMSIdByForeignIdResponseBodyData {
	s.Phone = &v
	return s
}

func (s *GetCMSIdByForeignIdResponseBodyData) SetStatus(v int32) *GetCMSIdByForeignIdResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetCMSIdByForeignIdResponseBodyData) SetUserId(v string) *GetCMSIdByForeignIdResponseBodyData {
	s.UserId = &v
	return s
}

type GetCMSIdByForeignIdResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetCMSIdByForeignIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetCMSIdByForeignIdResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCMSIdByForeignIdResponse) GoString() string {
	return s.String()
}

func (s *GetCMSIdByForeignIdResponse) SetHeaders(v map[string]*string) *GetCMSIdByForeignIdResponse {
	s.Headers = v
	return s
}

func (s *GetCMSIdByForeignIdResponse) SetStatusCode(v int32) *GetCMSIdByForeignIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCMSIdByForeignIdResponse) SetBody(v *GetCMSIdByForeignIdResponseBody) *GetCMSIdByForeignIdResponse {
	s.Body = v
	return s
}

type GetCallsPerDayRequest struct {
	DataId           *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	DataIdEnd        *string `json:"DataIdEnd,omitempty" xml:"DataIdEnd,omitempty"`
	DataIdStart      *string `json:"DataIdStart,omitempty" xml:"DataIdStart,omitempty"`
	HavePhoneNumbers *string `json:"HavePhoneNumbers,omitempty" xml:"HavePhoneNumbers,omitempty"`
	HourId           *string `json:"HourId,omitempty" xml:"HourId,omitempty"`
	InstanceId       *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MinuteId         *string `json:"MinuteId,omitempty" xml:"MinuteId,omitempty"`
	PageNo           *int64  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize         *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PhoneNumbers     *string `json:"PhoneNumbers,omitempty" xml:"PhoneNumbers,omitempty"`
}

func (s GetCallsPerDayRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCallsPerDayRequest) GoString() string {
	return s.String()
}

func (s *GetCallsPerDayRequest) SetDataId(v string) *GetCallsPerDayRequest {
	s.DataId = &v
	return s
}

func (s *GetCallsPerDayRequest) SetDataIdEnd(v string) *GetCallsPerDayRequest {
	s.DataIdEnd = &v
	return s
}

func (s *GetCallsPerDayRequest) SetDataIdStart(v string) *GetCallsPerDayRequest {
	s.DataIdStart = &v
	return s
}

func (s *GetCallsPerDayRequest) SetHavePhoneNumbers(v string) *GetCallsPerDayRequest {
	s.HavePhoneNumbers = &v
	return s
}

func (s *GetCallsPerDayRequest) SetHourId(v string) *GetCallsPerDayRequest {
	s.HourId = &v
	return s
}

func (s *GetCallsPerDayRequest) SetInstanceId(v string) *GetCallsPerDayRequest {
	s.InstanceId = &v
	return s
}

func (s *GetCallsPerDayRequest) SetMinuteId(v string) *GetCallsPerDayRequest {
	s.MinuteId = &v
	return s
}

func (s *GetCallsPerDayRequest) SetPageNo(v int64) *GetCallsPerDayRequest {
	s.PageNo = &v
	return s
}

func (s *GetCallsPerDayRequest) SetPageSize(v int64) *GetCallsPerDayRequest {
	s.PageSize = &v
	return s
}

func (s *GetCallsPerDayRequest) SetPhoneNumbers(v string) *GetCallsPerDayRequest {
	s.PhoneNumbers = &v
	return s
}

type GetCallsPerDayResponseBody struct {
	Code      *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetCallsPerDayResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *string                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetCallsPerDayResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCallsPerDayResponseBody) GoString() string {
	return s.String()
}

func (s *GetCallsPerDayResponseBody) SetCode(v string) *GetCallsPerDayResponseBody {
	s.Code = &v
	return s
}

func (s *GetCallsPerDayResponseBody) SetData(v *GetCallsPerDayResponseBodyData) *GetCallsPerDayResponseBody {
	s.Data = v
	return s
}

func (s *GetCallsPerDayResponseBody) SetMessage(v string) *GetCallsPerDayResponseBody {
	s.Message = &v
	return s
}

func (s *GetCallsPerDayResponseBody) SetRequestId(v string) *GetCallsPerDayResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCallsPerDayResponseBody) SetSuccess(v string) *GetCallsPerDayResponseBody {
	s.Success = &v
	return s
}

type GetCallsPerDayResponseBodyData struct {
	CallsPerdayResponseList []*GetCallsPerDayResponseBodyDataCallsPerdayResponseList `json:"CallsPerdayResponseList,omitempty" xml:"CallsPerdayResponseList,omitempty" type:"Repeated"`
	PageNo                  *string                                                  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize                *int64                                                   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalNum                *int64                                                   `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s GetCallsPerDayResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetCallsPerDayResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCallsPerDayResponseBodyData) SetCallsPerdayResponseList(v []*GetCallsPerDayResponseBodyDataCallsPerdayResponseList) *GetCallsPerDayResponseBodyData {
	s.CallsPerdayResponseList = v
	return s
}

func (s *GetCallsPerDayResponseBodyData) SetPageNo(v string) *GetCallsPerDayResponseBodyData {
	s.PageNo = &v
	return s
}

func (s *GetCallsPerDayResponseBodyData) SetPageSize(v int64) *GetCallsPerDayResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetCallsPerDayResponseBodyData) SetTotalNum(v int64) *GetCallsPerDayResponseBodyData {
	s.TotalNum = &v
	return s
}

type GetCallsPerDayResponseBodyDataCallsPerdayResponseList struct {
	CallInCnt  *string `json:"CallInCnt,omitempty" xml:"CallInCnt,omitempty"`
	CallOutCnt *string `json:"CallOutCnt,omitempty" xml:"CallOutCnt,omitempty"`
	DataId     *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	HourId     *string `json:"HourId,omitempty" xml:"HourId,omitempty"`
	MinuteId   *string `json:"MinuteId,omitempty" xml:"MinuteId,omitempty"`
	PhoneNum   *string `json:"PhoneNum,omitempty" xml:"PhoneNum,omitempty"`
	TenantId   *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	TenantName *string `json:"TenantName,omitempty" xml:"TenantName,omitempty"`
}

func (s GetCallsPerDayResponseBodyDataCallsPerdayResponseList) String() string {
	return tea.Prettify(s)
}

func (s GetCallsPerDayResponseBodyDataCallsPerdayResponseList) GoString() string {
	return s.String()
}

func (s *GetCallsPerDayResponseBodyDataCallsPerdayResponseList) SetCallInCnt(v string) *GetCallsPerDayResponseBodyDataCallsPerdayResponseList {
	s.CallInCnt = &v
	return s
}

func (s *GetCallsPerDayResponseBodyDataCallsPerdayResponseList) SetCallOutCnt(v string) *GetCallsPerDayResponseBodyDataCallsPerdayResponseList {
	s.CallOutCnt = &v
	return s
}

func (s *GetCallsPerDayResponseBodyDataCallsPerdayResponseList) SetDataId(v string) *GetCallsPerDayResponseBodyDataCallsPerdayResponseList {
	s.DataId = &v
	return s
}

func (s *GetCallsPerDayResponseBodyDataCallsPerdayResponseList) SetHourId(v string) *GetCallsPerDayResponseBodyDataCallsPerdayResponseList {
	s.HourId = &v
	return s
}

func (s *GetCallsPerDayResponseBodyDataCallsPerdayResponseList) SetMinuteId(v string) *GetCallsPerDayResponseBodyDataCallsPerdayResponseList {
	s.MinuteId = &v
	return s
}

func (s *GetCallsPerDayResponseBodyDataCallsPerdayResponseList) SetPhoneNum(v string) *GetCallsPerDayResponseBodyDataCallsPerdayResponseList {
	s.PhoneNum = &v
	return s
}

func (s *GetCallsPerDayResponseBodyDataCallsPerdayResponseList) SetTenantId(v string) *GetCallsPerDayResponseBodyDataCallsPerdayResponseList {
	s.TenantId = &v
	return s
}

func (s *GetCallsPerDayResponseBodyDataCallsPerdayResponseList) SetTenantName(v string) *GetCallsPerDayResponseBodyDataCallsPerdayResponseList {
	s.TenantName = &v
	return s
}

type GetCallsPerDayResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetCallsPerDayResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetCallsPerDayResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCallsPerDayResponse) GoString() string {
	return s.String()
}

func (s *GetCallsPerDayResponse) SetHeaders(v map[string]*string) *GetCallsPerDayResponse {
	s.Headers = v
	return s
}

func (s *GetCallsPerDayResponse) SetStatusCode(v int32) *GetCallsPerDayResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCallsPerDayResponse) SetBody(v *GetCallsPerDayResponseBody) *GetCallsPerDayResponse {
	s.Body = v
	return s
}

type GetEntityRouteRequest struct {
	EntityBizCode        *string `json:"EntityBizCode,omitempty" xml:"EntityBizCode,omitempty"`
	EntityId             *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	EntityName           *string `json:"EntityName,omitempty" xml:"EntityName,omitempty"`
	EntityRelationNumber *string `json:"EntityRelationNumber,omitempty" xml:"EntityRelationNumber,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	UniqueId             *int64  `json:"UniqueId,omitempty" xml:"UniqueId,omitempty"`
}

func (s GetEntityRouteRequest) String() string {
	return tea.Prettify(s)
}

func (s GetEntityRouteRequest) GoString() string {
	return s.String()
}

func (s *GetEntityRouteRequest) SetEntityBizCode(v string) *GetEntityRouteRequest {
	s.EntityBizCode = &v
	return s
}

func (s *GetEntityRouteRequest) SetEntityId(v string) *GetEntityRouteRequest {
	s.EntityId = &v
	return s
}

func (s *GetEntityRouteRequest) SetEntityName(v string) *GetEntityRouteRequest {
	s.EntityName = &v
	return s
}

func (s *GetEntityRouteRequest) SetEntityRelationNumber(v string) *GetEntityRouteRequest {
	s.EntityRelationNumber = &v
	return s
}

func (s *GetEntityRouteRequest) SetInstanceId(v string) *GetEntityRouteRequest {
	s.InstanceId = &v
	return s
}

func (s *GetEntityRouteRequest) SetUniqueId(v int64) *GetEntityRouteRequest {
	s.UniqueId = &v
	return s
}

type GetEntityRouteResponseBody struct {
	Code      *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetEntityRouteResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                           `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetEntityRouteResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetEntityRouteResponseBody) GoString() string {
	return s.String()
}

func (s *GetEntityRouteResponseBody) SetCode(v string) *GetEntityRouteResponseBody {
	s.Code = &v
	return s
}

func (s *GetEntityRouteResponseBody) SetData(v *GetEntityRouteResponseBodyData) *GetEntityRouteResponseBody {
	s.Data = v
	return s
}

func (s *GetEntityRouteResponseBody) SetMessage(v string) *GetEntityRouteResponseBody {
	s.Message = &v
	return s
}

func (s *GetEntityRouteResponseBody) SetRequestId(v string) *GetEntityRouteResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetEntityRouteResponseBody) SetSuccess(v bool) *GetEntityRouteResponseBody {
	s.Success = &v
	return s
}

type GetEntityRouteResponseBodyData struct {
	DepartmentId         *string `json:"DepartmentId,omitempty" xml:"DepartmentId,omitempty"`
	EntityBizCode        *string `json:"EntityBizCode,omitempty" xml:"EntityBizCode,omitempty"`
	EntityBizCodeType    *string `json:"EntityBizCodeType,omitempty" xml:"EntityBizCodeType,omitempty"`
	EntityId             *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	EntityName           *string `json:"EntityName,omitempty" xml:"EntityName,omitempty"`
	EntityRelationNumber *string `json:"EntityRelationNumber,omitempty" xml:"EntityRelationNumber,omitempty"`
	ExtInfo              *string `json:"ExtInfo,omitempty" xml:"ExtInfo,omitempty"`
	GroupId              *int64  `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	ServiceId            *int64  `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	UniqueId             *int64  `json:"UniqueId,omitempty" xml:"UniqueId,omitempty"`
}

func (s GetEntityRouteResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetEntityRouteResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetEntityRouteResponseBodyData) SetDepartmentId(v string) *GetEntityRouteResponseBodyData {
	s.DepartmentId = &v
	return s
}

func (s *GetEntityRouteResponseBodyData) SetEntityBizCode(v string) *GetEntityRouteResponseBodyData {
	s.EntityBizCode = &v
	return s
}

func (s *GetEntityRouteResponseBodyData) SetEntityBizCodeType(v string) *GetEntityRouteResponseBodyData {
	s.EntityBizCodeType = &v
	return s
}

func (s *GetEntityRouteResponseBodyData) SetEntityId(v string) *GetEntityRouteResponseBodyData {
	s.EntityId = &v
	return s
}

func (s *GetEntityRouteResponseBodyData) SetEntityName(v string) *GetEntityRouteResponseBodyData {
	s.EntityName = &v
	return s
}

func (s *GetEntityRouteResponseBodyData) SetEntityRelationNumber(v string) *GetEntityRouteResponseBodyData {
	s.EntityRelationNumber = &v
	return s
}

func (s *GetEntityRouteResponseBodyData) SetExtInfo(v string) *GetEntityRouteResponseBodyData {
	s.ExtInfo = &v
	return s
}

func (s *GetEntityRouteResponseBodyData) SetGroupId(v int64) *GetEntityRouteResponseBodyData {
	s.GroupId = &v
	return s
}

func (s *GetEntityRouteResponseBodyData) SetServiceId(v int64) *GetEntityRouteResponseBodyData {
	s.ServiceId = &v
	return s
}

func (s *GetEntityRouteResponseBodyData) SetUniqueId(v int64) *GetEntityRouteResponseBodyData {
	s.UniqueId = &v
	return s
}

type GetEntityRouteResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetEntityRouteResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetEntityRouteResponse) String() string {
	return tea.Prettify(s)
}

func (s GetEntityRouteResponse) GoString() string {
	return s.String()
}

func (s *GetEntityRouteResponse) SetHeaders(v map[string]*string) *GetEntityRouteResponse {
	s.Headers = v
	return s
}

func (s *GetEntityRouteResponse) SetStatusCode(v int32) *GetEntityRouteResponse {
	s.StatusCode = &v
	return s
}

func (s *GetEntityRouteResponse) SetBody(v *GetEntityRouteResponseBody) *GetEntityRouteResponse {
	s.Body = v
	return s
}

type GetEntityRouteListRequest struct {
	EntityName           *string `json:"EntityName,omitempty" xml:"EntityName,omitempty"`
	EntityRelationNumber *string `json:"EntityRelationNumber,omitempty" xml:"EntityRelationNumber,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PageNo               *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s GetEntityRouteListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetEntityRouteListRequest) GoString() string {
	return s.String()
}

func (s *GetEntityRouteListRequest) SetEntityName(v string) *GetEntityRouteListRequest {
	s.EntityName = &v
	return s
}

func (s *GetEntityRouteListRequest) SetEntityRelationNumber(v string) *GetEntityRouteListRequest {
	s.EntityRelationNumber = &v
	return s
}

func (s *GetEntityRouteListRequest) SetInstanceId(v string) *GetEntityRouteListRequest {
	s.InstanceId = &v
	return s
}

func (s *GetEntityRouteListRequest) SetPageNo(v int32) *GetEntityRouteListRequest {
	s.PageNo = &v
	return s
}

func (s *GetEntityRouteListRequest) SetPageSize(v int32) *GetEntityRouteListRequest {
	s.PageSize = &v
	return s
}

type GetEntityRouteListResponseBody struct {
	Code      *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetEntityRouteListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetEntityRouteListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetEntityRouteListResponseBody) GoString() string {
	return s.String()
}

func (s *GetEntityRouteListResponseBody) SetCode(v string) *GetEntityRouteListResponseBody {
	s.Code = &v
	return s
}

func (s *GetEntityRouteListResponseBody) SetData(v *GetEntityRouteListResponseBodyData) *GetEntityRouteListResponseBody {
	s.Data = v
	return s
}

func (s *GetEntityRouteListResponseBody) SetMessage(v string) *GetEntityRouteListResponseBody {
	s.Message = &v
	return s
}

func (s *GetEntityRouteListResponseBody) SetRequestId(v string) *GetEntityRouteListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetEntityRouteListResponseBody) SetSuccess(v bool) *GetEntityRouteListResponseBody {
	s.Success = &v
	return s
}

type GetEntityRouteListResponseBodyData struct {
	EntityRouteList []*GetEntityRouteListResponseBodyDataEntityRouteList `json:"EntityRouteList,omitempty" xml:"EntityRouteList,omitempty" type:"Repeated"`
	PageNo          *int32                                               `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize        *int32                                               `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Total           *int64                                               `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetEntityRouteListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetEntityRouteListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetEntityRouteListResponseBodyData) SetEntityRouteList(v []*GetEntityRouteListResponseBodyDataEntityRouteList) *GetEntityRouteListResponseBodyData {
	s.EntityRouteList = v
	return s
}

func (s *GetEntityRouteListResponseBodyData) SetPageNo(v int32) *GetEntityRouteListResponseBodyData {
	s.PageNo = &v
	return s
}

func (s *GetEntityRouteListResponseBodyData) SetPageSize(v int32) *GetEntityRouteListResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetEntityRouteListResponseBodyData) SetTotal(v int64) *GetEntityRouteListResponseBodyData {
	s.Total = &v
	return s
}

type GetEntityRouteListResponseBodyDataEntityRouteList struct {
	DepartmentId         *string `json:"DepartmentId,omitempty" xml:"DepartmentId,omitempty"`
	EntityBizCode        *string `json:"EntityBizCode,omitempty" xml:"EntityBizCode,omitempty"`
	EntityBizCodeType    *string `json:"EntityBizCodeType,omitempty" xml:"EntityBizCodeType,omitempty"`
	EntityId             *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	EntityName           *string `json:"EntityName,omitempty" xml:"EntityName,omitempty"`
	EntityRelationNumber *string `json:"EntityRelationNumber,omitempty" xml:"EntityRelationNumber,omitempty"`
	ExtInfo              *string `json:"ExtInfo,omitempty" xml:"ExtInfo,omitempty"`
	GroupId              *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	ServiceId            *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	UniqueId             *int64  `json:"UniqueId,omitempty" xml:"UniqueId,omitempty"`
}

func (s GetEntityRouteListResponseBodyDataEntityRouteList) String() string {
	return tea.Prettify(s)
}

func (s GetEntityRouteListResponseBodyDataEntityRouteList) GoString() string {
	return s.String()
}

func (s *GetEntityRouteListResponseBodyDataEntityRouteList) SetDepartmentId(v string) *GetEntityRouteListResponseBodyDataEntityRouteList {
	s.DepartmentId = &v
	return s
}

func (s *GetEntityRouteListResponseBodyDataEntityRouteList) SetEntityBizCode(v string) *GetEntityRouteListResponseBodyDataEntityRouteList {
	s.EntityBizCode = &v
	return s
}

func (s *GetEntityRouteListResponseBodyDataEntityRouteList) SetEntityBizCodeType(v string) *GetEntityRouteListResponseBodyDataEntityRouteList {
	s.EntityBizCodeType = &v
	return s
}

func (s *GetEntityRouteListResponseBodyDataEntityRouteList) SetEntityId(v string) *GetEntityRouteListResponseBodyDataEntityRouteList {
	s.EntityId = &v
	return s
}

func (s *GetEntityRouteListResponseBodyDataEntityRouteList) SetEntityName(v string) *GetEntityRouteListResponseBodyDataEntityRouteList {
	s.EntityName = &v
	return s
}

func (s *GetEntityRouteListResponseBodyDataEntityRouteList) SetEntityRelationNumber(v string) *GetEntityRouteListResponseBodyDataEntityRouteList {
	s.EntityRelationNumber = &v
	return s
}

func (s *GetEntityRouteListResponseBodyDataEntityRouteList) SetExtInfo(v string) *GetEntityRouteListResponseBodyDataEntityRouteList {
	s.ExtInfo = &v
	return s
}

func (s *GetEntityRouteListResponseBodyDataEntityRouteList) SetGroupId(v string) *GetEntityRouteListResponseBodyDataEntityRouteList {
	s.GroupId = &v
	return s
}

func (s *GetEntityRouteListResponseBodyDataEntityRouteList) SetServiceId(v string) *GetEntityRouteListResponseBodyDataEntityRouteList {
	s.ServiceId = &v
	return s
}

func (s *GetEntityRouteListResponseBodyDataEntityRouteList) SetUniqueId(v int64) *GetEntityRouteListResponseBodyDataEntityRouteList {
	s.UniqueId = &v
	return s
}

type GetEntityRouteListResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetEntityRouteListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetEntityRouteListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetEntityRouteListResponse) GoString() string {
	return s.String()
}

func (s *GetEntityRouteListResponse) SetHeaders(v map[string]*string) *GetEntityRouteListResponse {
	s.Headers = v
	return s
}

func (s *GetEntityRouteListResponse) SetStatusCode(v int32) *GetEntityRouteListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetEntityRouteListResponse) SetBody(v *GetEntityRouteListResponseBody) *GetEntityRouteListResponse {
	s.Body = v
	return s
}

type GetEntityTagRelationRequest struct {
	EntityId   *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetEntityTagRelationRequest) String() string {
	return tea.Prettify(s)
}

func (s GetEntityTagRelationRequest) GoString() string {
	return s.String()
}

func (s *GetEntityTagRelationRequest) SetEntityId(v string) *GetEntityTagRelationRequest {
	s.EntityId = &v
	return s
}

func (s *GetEntityTagRelationRequest) SetEntityType(v string) *GetEntityTagRelationRequest {
	s.EntityType = &v
	return s
}

func (s *GetEntityTagRelationRequest) SetInstanceId(v string) *GetEntityTagRelationRequest {
	s.InstanceId = &v
	return s
}

type GetEntityTagRelationResponseBody struct {
	Code      *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*GetEntityTagRelationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetEntityTagRelationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetEntityTagRelationResponseBody) GoString() string {
	return s.String()
}

func (s *GetEntityTagRelationResponseBody) SetCode(v string) *GetEntityTagRelationResponseBody {
	s.Code = &v
	return s
}

func (s *GetEntityTagRelationResponseBody) SetData(v []*GetEntityTagRelationResponseBodyData) *GetEntityTagRelationResponseBody {
	s.Data = v
	return s
}

func (s *GetEntityTagRelationResponseBody) SetMessage(v string) *GetEntityTagRelationResponseBody {
	s.Message = &v
	return s
}

func (s *GetEntityTagRelationResponseBody) SetRequestId(v string) *GetEntityTagRelationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetEntityTagRelationResponseBody) SetSuccess(v bool) *GetEntityTagRelationResponseBody {
	s.Success = &v
	return s
}

type GetEntityTagRelationResponseBodyData struct {
	EntityId     *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	EntityType   *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	TagCode      *string `json:"TagCode,omitempty" xml:"TagCode,omitempty"`
	TagGroupCode *string `json:"TagGroupCode,omitempty" xml:"TagGroupCode,omitempty"`
	TagGroupName *string `json:"TagGroupName,omitempty" xml:"TagGroupName,omitempty"`
	TagName      *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
}

func (s GetEntityTagRelationResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetEntityTagRelationResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetEntityTagRelationResponseBodyData) SetEntityId(v string) *GetEntityTagRelationResponseBodyData {
	s.EntityId = &v
	return s
}

func (s *GetEntityTagRelationResponseBodyData) SetEntityType(v string) *GetEntityTagRelationResponseBodyData {
	s.EntityType = &v
	return s
}

func (s *GetEntityTagRelationResponseBodyData) SetTagCode(v string) *GetEntityTagRelationResponseBodyData {
	s.TagCode = &v
	return s
}

func (s *GetEntityTagRelationResponseBodyData) SetTagGroupCode(v string) *GetEntityTagRelationResponseBodyData {
	s.TagGroupCode = &v
	return s
}

func (s *GetEntityTagRelationResponseBodyData) SetTagGroupName(v string) *GetEntityTagRelationResponseBodyData {
	s.TagGroupName = &v
	return s
}

func (s *GetEntityTagRelationResponseBodyData) SetTagName(v string) *GetEntityTagRelationResponseBodyData {
	s.TagName = &v
	return s
}

type GetEntityTagRelationResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetEntityTagRelationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetEntityTagRelationResponse) String() string {
	return tea.Prettify(s)
}

func (s GetEntityTagRelationResponse) GoString() string {
	return s.String()
}

func (s *GetEntityTagRelationResponse) SetHeaders(v map[string]*string) *GetEntityTagRelationResponse {
	s.Headers = v
	return s
}

func (s *GetEntityTagRelationResponse) SetStatusCode(v int32) *GetEntityTagRelationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetEntityTagRelationResponse) SetBody(v *GetEntityTagRelationResponseBody) *GetEntityTagRelationResponse {
	s.Body = v
	return s
}

type GetForeignIdByCMSIdRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	UserId     *int64  `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetForeignIdByCMSIdRequest) String() string {
	return tea.Prettify(s)
}

func (s GetForeignIdByCMSIdRequest) GoString() string {
	return s.String()
}

func (s *GetForeignIdByCMSIdRequest) SetInstanceId(v string) *GetForeignIdByCMSIdRequest {
	s.InstanceId = &v
	return s
}

func (s *GetForeignIdByCMSIdRequest) SetUserId(v int64) *GetForeignIdByCMSIdRequest {
	s.UserId = &v
	return s
}

type GetForeignIdByCMSIdResponseBody struct {
	Data      *GetForeignIdByCMSIdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetForeignIdByCMSIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetForeignIdByCMSIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetForeignIdByCMSIdResponseBody) SetData(v *GetForeignIdByCMSIdResponseBodyData) *GetForeignIdByCMSIdResponseBody {
	s.Data = v
	return s
}

func (s *GetForeignIdByCMSIdResponseBody) SetMessage(v string) *GetForeignIdByCMSIdResponseBody {
	s.Message = &v
	return s
}

func (s *GetForeignIdByCMSIdResponseBody) SetRequestId(v string) *GetForeignIdByCMSIdResponseBody {
	s.RequestId = &v
	return s
}

type GetForeignIdByCMSIdResponseBodyData struct {
	Anonymity      *bool   `json:"Anonymity,omitempty" xml:"Anonymity,omitempty"`
	Avatar         *string `json:"Avatar,omitempty" xml:"Avatar,omitempty"`
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	CustomerTypeId *int32  `json:"CustomerTypeId,omitempty" xml:"CustomerTypeId,omitempty"`
	ForeignId      *string `json:"ForeignId,omitempty" xml:"ForeignId,omitempty"`
	Gender         *string `json:"Gender,omitempty" xml:"Gender,omitempty"`
	Nick           *string `json:"Nick,omitempty" xml:"Nick,omitempty"`
	Phone          *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
	Status         *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	UserId         *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetForeignIdByCMSIdResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetForeignIdByCMSIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetForeignIdByCMSIdResponseBodyData) SetAnonymity(v bool) *GetForeignIdByCMSIdResponseBodyData {
	s.Anonymity = &v
	return s
}

func (s *GetForeignIdByCMSIdResponseBodyData) SetAvatar(v string) *GetForeignIdByCMSIdResponseBodyData {
	s.Avatar = &v
	return s
}

func (s *GetForeignIdByCMSIdResponseBodyData) SetCode(v string) *GetForeignIdByCMSIdResponseBodyData {
	s.Code = &v
	return s
}

func (s *GetForeignIdByCMSIdResponseBodyData) SetCustomerTypeId(v int32) *GetForeignIdByCMSIdResponseBodyData {
	s.CustomerTypeId = &v
	return s
}

func (s *GetForeignIdByCMSIdResponseBodyData) SetForeignId(v string) *GetForeignIdByCMSIdResponseBodyData {
	s.ForeignId = &v
	return s
}

func (s *GetForeignIdByCMSIdResponseBodyData) SetGender(v string) *GetForeignIdByCMSIdResponseBodyData {
	s.Gender = &v
	return s
}

func (s *GetForeignIdByCMSIdResponseBodyData) SetNick(v string) *GetForeignIdByCMSIdResponseBodyData {
	s.Nick = &v
	return s
}

func (s *GetForeignIdByCMSIdResponseBodyData) SetPhone(v string) *GetForeignIdByCMSIdResponseBodyData {
	s.Phone = &v
	return s
}

func (s *GetForeignIdByCMSIdResponseBodyData) SetStatus(v int32) *GetForeignIdByCMSIdResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetForeignIdByCMSIdResponseBodyData) SetSuccess(v bool) *GetForeignIdByCMSIdResponseBodyData {
	s.Success = &v
	return s
}

func (s *GetForeignIdByCMSIdResponseBodyData) SetUserId(v string) *GetForeignIdByCMSIdResponseBodyData {
	s.UserId = &v
	return s
}

type GetForeignIdByCMSIdResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetForeignIdByCMSIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetForeignIdByCMSIdResponse) String() string {
	return tea.Prettify(s)
}

func (s GetForeignIdByCMSIdResponse) GoString() string {
	return s.String()
}

func (s *GetForeignIdByCMSIdResponse) SetHeaders(v map[string]*string) *GetForeignIdByCMSIdResponse {
	s.Headers = v
	return s
}

func (s *GetForeignIdByCMSIdResponse) SetStatusCode(v int32) *GetForeignIdByCMSIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetForeignIdByCMSIdResponse) SetBody(v *GetForeignIdByCMSIdResponseBody) *GetForeignIdByCMSIdResponse {
	s.Body = v
	return s
}

type GetGrantedRoleIdsRequest struct {
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetGrantedRoleIdsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetGrantedRoleIdsRequest) GoString() string {
	return s.String()
}

func (s *GetGrantedRoleIdsRequest) SetAccountName(v string) *GetGrantedRoleIdsRequest {
	s.AccountName = &v
	return s
}

func (s *GetGrantedRoleIdsRequest) SetClientToken(v string) *GetGrantedRoleIdsRequest {
	s.ClientToken = &v
	return s
}

func (s *GetGrantedRoleIdsRequest) SetInstanceId(v string) *GetGrantedRoleIdsRequest {
	s.InstanceId = &v
	return s
}

type GetGrantedRoleIdsResponseBody struct {
	Code           *string  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           []*int64 `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	HttpStatusCode *int64   `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetGrantedRoleIdsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetGrantedRoleIdsResponseBody) GoString() string {
	return s.String()
}

func (s *GetGrantedRoleIdsResponseBody) SetCode(v string) *GetGrantedRoleIdsResponseBody {
	s.Code = &v
	return s
}

func (s *GetGrantedRoleIdsResponseBody) SetData(v []*int64) *GetGrantedRoleIdsResponseBody {
	s.Data = v
	return s
}

func (s *GetGrantedRoleIdsResponseBody) SetHttpStatusCode(v int64) *GetGrantedRoleIdsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetGrantedRoleIdsResponseBody) SetMessage(v string) *GetGrantedRoleIdsResponseBody {
	s.Message = &v
	return s
}

func (s *GetGrantedRoleIdsResponseBody) SetRequestId(v string) *GetGrantedRoleIdsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetGrantedRoleIdsResponseBody) SetSuccess(v bool) *GetGrantedRoleIdsResponseBody {
	s.Success = &v
	return s
}

type GetGrantedRoleIdsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetGrantedRoleIdsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetGrantedRoleIdsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetGrantedRoleIdsResponse) GoString() string {
	return s.String()
}

func (s *GetGrantedRoleIdsResponse) SetHeaders(v map[string]*string) *GetGrantedRoleIdsResponse {
	s.Headers = v
	return s
}

func (s *GetGrantedRoleIdsResponse) SetStatusCode(v int32) *GetGrantedRoleIdsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetGrantedRoleIdsResponse) SetBody(v *GetGrantedRoleIdsResponseBody) *GetGrantedRoleIdsResponse {
	s.Body = v
	return s
}

type GetHotlineAgentDetailRequest struct {
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetHotlineAgentDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s GetHotlineAgentDetailRequest) GoString() string {
	return s.String()
}

func (s *GetHotlineAgentDetailRequest) SetAccountName(v string) *GetHotlineAgentDetailRequest {
	s.AccountName = &v
	return s
}

func (s *GetHotlineAgentDetailRequest) SetClientToken(v string) *GetHotlineAgentDetailRequest {
	s.ClientToken = &v
	return s
}

func (s *GetHotlineAgentDetailRequest) SetInstanceId(v string) *GetHotlineAgentDetailRequest {
	s.InstanceId = &v
	return s
}

type GetHotlineAgentDetailResponseBody struct {
	Code           *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *GetHotlineAgentDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int64                                 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetHotlineAgentDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetHotlineAgentDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetHotlineAgentDetailResponseBody) SetCode(v string) *GetHotlineAgentDetailResponseBody {
	s.Code = &v
	return s
}

func (s *GetHotlineAgentDetailResponseBody) SetData(v *GetHotlineAgentDetailResponseBodyData) *GetHotlineAgentDetailResponseBody {
	s.Data = v
	return s
}

func (s *GetHotlineAgentDetailResponseBody) SetHttpStatusCode(v int64) *GetHotlineAgentDetailResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetHotlineAgentDetailResponseBody) SetMessage(v string) *GetHotlineAgentDetailResponseBody {
	s.Message = &v
	return s
}

func (s *GetHotlineAgentDetailResponseBody) SetRequestId(v string) *GetHotlineAgentDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHotlineAgentDetailResponseBody) SetSuccess(v bool) *GetHotlineAgentDetailResponseBody {
	s.Success = &v
	return s
}

type GetHotlineAgentDetailResponseBodyData struct {
	AgentId         *int64  `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	AgentStatus     *int32  `json:"AgentStatus,omitempty" xml:"AgentStatus,omitempty"`
	AgentStatusCode *string `json:"AgentStatusCode,omitempty" xml:"AgentStatusCode,omitempty"`
	Assigned        *bool   `json:"Assigned,omitempty" xml:"Assigned,omitempty"`
	RestType        *int32  `json:"RestType,omitempty" xml:"RestType,omitempty"`
	TenantId        *int64  `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	Token           *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s GetHotlineAgentDetailResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetHotlineAgentDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetHotlineAgentDetailResponseBodyData) SetAgentId(v int64) *GetHotlineAgentDetailResponseBodyData {
	s.AgentId = &v
	return s
}

func (s *GetHotlineAgentDetailResponseBodyData) SetAgentStatus(v int32) *GetHotlineAgentDetailResponseBodyData {
	s.AgentStatus = &v
	return s
}

func (s *GetHotlineAgentDetailResponseBodyData) SetAgentStatusCode(v string) *GetHotlineAgentDetailResponseBodyData {
	s.AgentStatusCode = &v
	return s
}

func (s *GetHotlineAgentDetailResponseBodyData) SetAssigned(v bool) *GetHotlineAgentDetailResponseBodyData {
	s.Assigned = &v
	return s
}

func (s *GetHotlineAgentDetailResponseBodyData) SetRestType(v int32) *GetHotlineAgentDetailResponseBodyData {
	s.RestType = &v
	return s
}

func (s *GetHotlineAgentDetailResponseBodyData) SetTenantId(v int64) *GetHotlineAgentDetailResponseBodyData {
	s.TenantId = &v
	return s
}

func (s *GetHotlineAgentDetailResponseBodyData) SetToken(v string) *GetHotlineAgentDetailResponseBodyData {
	s.Token = &v
	return s
}

type GetHotlineAgentDetailResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetHotlineAgentDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetHotlineAgentDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s GetHotlineAgentDetailResponse) GoString() string {
	return s.String()
}

func (s *GetHotlineAgentDetailResponse) SetHeaders(v map[string]*string) *GetHotlineAgentDetailResponse {
	s.Headers = v
	return s
}

func (s *GetHotlineAgentDetailResponse) SetStatusCode(v int32) *GetHotlineAgentDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHotlineAgentDetailResponse) SetBody(v *GetHotlineAgentDetailResponseBody) *GetHotlineAgentDetailResponse {
	s.Body = v
	return s
}

type GetHotlineAgentDetailReportRequest struct {
	CurrentPage *int32   `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	DepIds      []*int64 `json:"DepIds,omitempty" xml:"DepIds,omitempty" type:"Repeated"`
	EndDate     *int64   `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	GroupIds    []*int64 `json:"GroupIds,omitempty" xml:"GroupIds,omitempty" type:"Repeated"`
	InstanceId  *string  `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PageSize    *int32   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	StartDate   *int64   `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s GetHotlineAgentDetailReportRequest) String() string {
	return tea.Prettify(s)
}

func (s GetHotlineAgentDetailReportRequest) GoString() string {
	return s.String()
}

func (s *GetHotlineAgentDetailReportRequest) SetCurrentPage(v int32) *GetHotlineAgentDetailReportRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetHotlineAgentDetailReportRequest) SetDepIds(v []*int64) *GetHotlineAgentDetailReportRequest {
	s.DepIds = v
	return s
}

func (s *GetHotlineAgentDetailReportRequest) SetEndDate(v int64) *GetHotlineAgentDetailReportRequest {
	s.EndDate = &v
	return s
}

func (s *GetHotlineAgentDetailReportRequest) SetGroupIds(v []*int64) *GetHotlineAgentDetailReportRequest {
	s.GroupIds = v
	return s
}

func (s *GetHotlineAgentDetailReportRequest) SetInstanceId(v string) *GetHotlineAgentDetailReportRequest {
	s.InstanceId = &v
	return s
}

func (s *GetHotlineAgentDetailReportRequest) SetPageSize(v int32) *GetHotlineAgentDetailReportRequest {
	s.PageSize = &v
	return s
}

func (s *GetHotlineAgentDetailReportRequest) SetStartDate(v int64) *GetHotlineAgentDetailReportRequest {
	s.StartDate = &v
	return s
}

type GetHotlineAgentDetailReportResponseBody struct {
	Code      *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetHotlineAgentDetailReportResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *string                                      `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetHotlineAgentDetailReportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetHotlineAgentDetailReportResponseBody) GoString() string {
	return s.String()
}

func (s *GetHotlineAgentDetailReportResponseBody) SetCode(v string) *GetHotlineAgentDetailReportResponseBody {
	s.Code = &v
	return s
}

func (s *GetHotlineAgentDetailReportResponseBody) SetData(v *GetHotlineAgentDetailReportResponseBodyData) *GetHotlineAgentDetailReportResponseBody {
	s.Data = v
	return s
}

func (s *GetHotlineAgentDetailReportResponseBody) SetMessage(v string) *GetHotlineAgentDetailReportResponseBody {
	s.Message = &v
	return s
}

func (s *GetHotlineAgentDetailReportResponseBody) SetRequestId(v string) *GetHotlineAgentDetailReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHotlineAgentDetailReportResponseBody) SetSuccess(v string) *GetHotlineAgentDetailReportResponseBody {
	s.Success = &v
	return s
}

type GetHotlineAgentDetailReportResponseBodyData struct {
	Columns  []*GetHotlineAgentDetailReportResponseBodyDataColumns `json:"Columns,omitempty" xml:"Columns,omitempty" type:"Repeated"`
	Page     *int32                                                `json:"Page,omitempty" xml:"Page,omitempty"`
	PageSize *int32                                                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Rows     []map[string]interface{}                              `json:"Rows,omitempty" xml:"Rows,omitempty" type:"Repeated"`
	Total    *int32                                                `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetHotlineAgentDetailReportResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetHotlineAgentDetailReportResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetHotlineAgentDetailReportResponseBodyData) SetColumns(v []*GetHotlineAgentDetailReportResponseBodyDataColumns) *GetHotlineAgentDetailReportResponseBodyData {
	s.Columns = v
	return s
}

func (s *GetHotlineAgentDetailReportResponseBodyData) SetPage(v int32) *GetHotlineAgentDetailReportResponseBodyData {
	s.Page = &v
	return s
}

func (s *GetHotlineAgentDetailReportResponseBodyData) SetPageSize(v int32) *GetHotlineAgentDetailReportResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetHotlineAgentDetailReportResponseBodyData) SetRows(v []map[string]interface{}) *GetHotlineAgentDetailReportResponseBodyData {
	s.Rows = v
	return s
}

func (s *GetHotlineAgentDetailReportResponseBodyData) SetTotal(v int32) *GetHotlineAgentDetailReportResponseBodyData {
	s.Total = &v
	return s
}

type GetHotlineAgentDetailReportResponseBodyDataColumns struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s GetHotlineAgentDetailReportResponseBodyDataColumns) String() string {
	return tea.Prettify(s)
}

func (s GetHotlineAgentDetailReportResponseBodyDataColumns) GoString() string {
	return s.String()
}

func (s *GetHotlineAgentDetailReportResponseBodyDataColumns) SetKey(v string) *GetHotlineAgentDetailReportResponseBodyDataColumns {
	s.Key = &v
	return s
}

func (s *GetHotlineAgentDetailReportResponseBodyDataColumns) SetTitle(v string) *GetHotlineAgentDetailReportResponseBodyDataColumns {
	s.Title = &v
	return s
}

type GetHotlineAgentDetailReportResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetHotlineAgentDetailReportResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetHotlineAgentDetailReportResponse) String() string {
	return tea.Prettify(s)
}

func (s GetHotlineAgentDetailReportResponse) GoString() string {
	return s.String()
}

func (s *GetHotlineAgentDetailReportResponse) SetHeaders(v map[string]*string) *GetHotlineAgentDetailReportResponse {
	s.Headers = v
	return s
}

func (s *GetHotlineAgentDetailReportResponse) SetStatusCode(v int32) *GetHotlineAgentDetailReportResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHotlineAgentDetailReportResponse) SetBody(v *GetHotlineAgentDetailReportResponseBody) *GetHotlineAgentDetailReportResponse {
	s.Body = v
	return s
}

type GetHotlineAgentDetailWithChannelRequest struct {
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetHotlineAgentDetailWithChannelRequest) String() string {
	return tea.Prettify(s)
}

func (s GetHotlineAgentDetailWithChannelRequest) GoString() string {
	return s.String()
}

func (s *GetHotlineAgentDetailWithChannelRequest) SetAccountName(v string) *GetHotlineAgentDetailWithChannelRequest {
	s.AccountName = &v
	return s
}

func (s *GetHotlineAgentDetailWithChannelRequest) SetClientToken(v string) *GetHotlineAgentDetailWithChannelRequest {
	s.ClientToken = &v
	return s
}

func (s *GetHotlineAgentDetailWithChannelRequest) SetInstanceId(v string) *GetHotlineAgentDetailWithChannelRequest {
	s.InstanceId = &v
	return s
}

type GetHotlineAgentDetailWithChannelResponseBody struct {
	Code           *string                                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *GetHotlineAgentDetailWithChannelResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int64                                            `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetHotlineAgentDetailWithChannelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetHotlineAgentDetailWithChannelResponseBody) GoString() string {
	return s.String()
}

func (s *GetHotlineAgentDetailWithChannelResponseBody) SetCode(v string) *GetHotlineAgentDetailWithChannelResponseBody {
	s.Code = &v
	return s
}

func (s *GetHotlineAgentDetailWithChannelResponseBody) SetData(v *GetHotlineAgentDetailWithChannelResponseBodyData) *GetHotlineAgentDetailWithChannelResponseBody {
	s.Data = v
	return s
}

func (s *GetHotlineAgentDetailWithChannelResponseBody) SetHttpStatusCode(v int64) *GetHotlineAgentDetailWithChannelResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetHotlineAgentDetailWithChannelResponseBody) SetMessage(v string) *GetHotlineAgentDetailWithChannelResponseBody {
	s.Message = &v
	return s
}

func (s *GetHotlineAgentDetailWithChannelResponseBody) SetRequestId(v string) *GetHotlineAgentDetailWithChannelResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHotlineAgentDetailWithChannelResponseBody) SetSuccess(v bool) *GetHotlineAgentDetailWithChannelResponseBody {
	s.Success = &v
	return s
}

type GetHotlineAgentDetailWithChannelResponseBodyData struct {
	AgentId         *int64  `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	AgentStatus     *int32  `json:"AgentStatus,omitempty" xml:"AgentStatus,omitempty"`
	AgentStatusCode *string `json:"AgentStatusCode,omitempty" xml:"AgentStatusCode,omitempty"`
	Assigned        *bool   `json:"Assigned,omitempty" xml:"Assigned,omitempty"`
	ExtAttr         *string `json:"ExtAttr,omitempty" xml:"ExtAttr,omitempty"`
	RestType        *int32  `json:"RestType,omitempty" xml:"RestType,omitempty"`
	TenantId        *int64  `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	Token           *string `json:"Token,omitempty" xml:"Token,omitempty"`
	WorkChannel     *string `json:"WorkChannel,omitempty" xml:"WorkChannel,omitempty"`
}

func (s GetHotlineAgentDetailWithChannelResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetHotlineAgentDetailWithChannelResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetHotlineAgentDetailWithChannelResponseBodyData) SetAgentId(v int64) *GetHotlineAgentDetailWithChannelResponseBodyData {
	s.AgentId = &v
	return s
}

func (s *GetHotlineAgentDetailWithChannelResponseBodyData) SetAgentStatus(v int32) *GetHotlineAgentDetailWithChannelResponseBodyData {
	s.AgentStatus = &v
	return s
}

func (s *GetHotlineAgentDetailWithChannelResponseBodyData) SetAgentStatusCode(v string) *GetHotlineAgentDetailWithChannelResponseBodyData {
	s.AgentStatusCode = &v
	return s
}

func (s *GetHotlineAgentDetailWithChannelResponseBodyData) SetAssigned(v bool) *GetHotlineAgentDetailWithChannelResponseBodyData {
	s.Assigned = &v
	return s
}

func (s *GetHotlineAgentDetailWithChannelResponseBodyData) SetExtAttr(v string) *GetHotlineAgentDetailWithChannelResponseBodyData {
	s.ExtAttr = &v
	return s
}

func (s *GetHotlineAgentDetailWithChannelResponseBodyData) SetRestType(v int32) *GetHotlineAgentDetailWithChannelResponseBodyData {
	s.RestType = &v
	return s
}

func (s *GetHotlineAgentDetailWithChannelResponseBodyData) SetTenantId(v int64) *GetHotlineAgentDetailWithChannelResponseBodyData {
	s.TenantId = &v
	return s
}

func (s *GetHotlineAgentDetailWithChannelResponseBodyData) SetToken(v string) *GetHotlineAgentDetailWithChannelResponseBodyData {
	s.Token = &v
	return s
}

func (s *GetHotlineAgentDetailWithChannelResponseBodyData) SetWorkChannel(v string) *GetHotlineAgentDetailWithChannelResponseBodyData {
	s.WorkChannel = &v
	return s
}

type GetHotlineAgentDetailWithChannelResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetHotlineAgentDetailWithChannelResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetHotlineAgentDetailWithChannelResponse) String() string {
	return tea.Prettify(s)
}

func (s GetHotlineAgentDetailWithChannelResponse) GoString() string {
	return s.String()
}

func (s *GetHotlineAgentDetailWithChannelResponse) SetHeaders(v map[string]*string) *GetHotlineAgentDetailWithChannelResponse {
	s.Headers = v
	return s
}

func (s *GetHotlineAgentDetailWithChannelResponse) SetStatusCode(v int32) *GetHotlineAgentDetailWithChannelResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHotlineAgentDetailWithChannelResponse) SetBody(v *GetHotlineAgentDetailWithChannelResponseBody) *GetHotlineAgentDetailWithChannelResponse {
	s.Body = v
	return s
}

type GetHotlineAgentStatusRequest struct {
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetHotlineAgentStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetHotlineAgentStatusRequest) GoString() string {
	return s.String()
}

func (s *GetHotlineAgentStatusRequest) SetAccountName(v string) *GetHotlineAgentStatusRequest {
	s.AccountName = &v
	return s
}

func (s *GetHotlineAgentStatusRequest) SetInstanceId(v string) *GetHotlineAgentStatusRequest {
	s.InstanceId = &v
	return s
}

type GetHotlineAgentStatusResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *string `json:"Data,omitempty" xml:"Data,omitempty"`
	HttpStatusCode *int64  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetHotlineAgentStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetHotlineAgentStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetHotlineAgentStatusResponseBody) SetCode(v string) *GetHotlineAgentStatusResponseBody {
	s.Code = &v
	return s
}

func (s *GetHotlineAgentStatusResponseBody) SetData(v string) *GetHotlineAgentStatusResponseBody {
	s.Data = &v
	return s
}

func (s *GetHotlineAgentStatusResponseBody) SetHttpStatusCode(v int64) *GetHotlineAgentStatusResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetHotlineAgentStatusResponseBody) SetMessage(v string) *GetHotlineAgentStatusResponseBody {
	s.Message = &v
	return s
}

func (s *GetHotlineAgentStatusResponseBody) SetRequestId(v string) *GetHotlineAgentStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHotlineAgentStatusResponseBody) SetSuccess(v bool) *GetHotlineAgentStatusResponseBody {
	s.Success = &v
	return s
}

type GetHotlineAgentStatusResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetHotlineAgentStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetHotlineAgentStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetHotlineAgentStatusResponse) GoString() string {
	return s.String()
}

func (s *GetHotlineAgentStatusResponse) SetHeaders(v map[string]*string) *GetHotlineAgentStatusResponse {
	s.Headers = v
	return s
}

func (s *GetHotlineAgentStatusResponse) SetStatusCode(v int32) *GetHotlineAgentStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHotlineAgentStatusResponse) SetBody(v *GetHotlineAgentStatusResponseBody) *GetHotlineAgentStatusResponse {
	s.Body = v
	return s
}

type GetHotlineGroupDetailReportRequest struct {
	CurrentPage *int32   `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	DepIds      []*int64 `json:"DepIds,omitempty" xml:"DepIds,omitempty" type:"Repeated"`
	EndDate     *int64   `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	GroupIds    []*int64 `json:"GroupIds,omitempty" xml:"GroupIds,omitempty" type:"Repeated"`
	InstanceId  *string  `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PageSize    *int32   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	StartDate   *int64   `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s GetHotlineGroupDetailReportRequest) String() string {
	return tea.Prettify(s)
}

func (s GetHotlineGroupDetailReportRequest) GoString() string {
	return s.String()
}

func (s *GetHotlineGroupDetailReportRequest) SetCurrentPage(v int32) *GetHotlineGroupDetailReportRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetHotlineGroupDetailReportRequest) SetDepIds(v []*int64) *GetHotlineGroupDetailReportRequest {
	s.DepIds = v
	return s
}

func (s *GetHotlineGroupDetailReportRequest) SetEndDate(v int64) *GetHotlineGroupDetailReportRequest {
	s.EndDate = &v
	return s
}

func (s *GetHotlineGroupDetailReportRequest) SetGroupIds(v []*int64) *GetHotlineGroupDetailReportRequest {
	s.GroupIds = v
	return s
}

func (s *GetHotlineGroupDetailReportRequest) SetInstanceId(v string) *GetHotlineGroupDetailReportRequest {
	s.InstanceId = &v
	return s
}

func (s *GetHotlineGroupDetailReportRequest) SetPageSize(v int32) *GetHotlineGroupDetailReportRequest {
	s.PageSize = &v
	return s
}

func (s *GetHotlineGroupDetailReportRequest) SetStartDate(v int64) *GetHotlineGroupDetailReportRequest {
	s.StartDate = &v
	return s
}

type GetHotlineGroupDetailReportResponseBody struct {
	Code      *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetHotlineGroupDetailReportResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *string                                      `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetHotlineGroupDetailReportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetHotlineGroupDetailReportResponseBody) GoString() string {
	return s.String()
}

func (s *GetHotlineGroupDetailReportResponseBody) SetCode(v string) *GetHotlineGroupDetailReportResponseBody {
	s.Code = &v
	return s
}

func (s *GetHotlineGroupDetailReportResponseBody) SetData(v *GetHotlineGroupDetailReportResponseBodyData) *GetHotlineGroupDetailReportResponseBody {
	s.Data = v
	return s
}

func (s *GetHotlineGroupDetailReportResponseBody) SetMessage(v string) *GetHotlineGroupDetailReportResponseBody {
	s.Message = &v
	return s
}

func (s *GetHotlineGroupDetailReportResponseBody) SetRequestId(v string) *GetHotlineGroupDetailReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHotlineGroupDetailReportResponseBody) SetSuccess(v string) *GetHotlineGroupDetailReportResponseBody {
	s.Success = &v
	return s
}

type GetHotlineGroupDetailReportResponseBodyData struct {
	Columns  []*GetHotlineGroupDetailReportResponseBodyDataColumns `json:"Columns,omitempty" xml:"Columns,omitempty" type:"Repeated"`
	Page     *int32                                                `json:"Page,omitempty" xml:"Page,omitempty"`
	PageSize *int32                                                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Rows     []map[string]interface{}                              `json:"Rows,omitempty" xml:"Rows,omitempty" type:"Repeated"`
	Total    *int32                                                `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetHotlineGroupDetailReportResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetHotlineGroupDetailReportResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetHotlineGroupDetailReportResponseBodyData) SetColumns(v []*GetHotlineGroupDetailReportResponseBodyDataColumns) *GetHotlineGroupDetailReportResponseBodyData {
	s.Columns = v
	return s
}

func (s *GetHotlineGroupDetailReportResponseBodyData) SetPage(v int32) *GetHotlineGroupDetailReportResponseBodyData {
	s.Page = &v
	return s
}

func (s *GetHotlineGroupDetailReportResponseBodyData) SetPageSize(v int32) *GetHotlineGroupDetailReportResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetHotlineGroupDetailReportResponseBodyData) SetRows(v []map[string]interface{}) *GetHotlineGroupDetailReportResponseBodyData {
	s.Rows = v
	return s
}

func (s *GetHotlineGroupDetailReportResponseBodyData) SetTotal(v int32) *GetHotlineGroupDetailReportResponseBodyData {
	s.Total = &v
	return s
}

type GetHotlineGroupDetailReportResponseBodyDataColumns struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s GetHotlineGroupDetailReportResponseBodyDataColumns) String() string {
	return tea.Prettify(s)
}

func (s GetHotlineGroupDetailReportResponseBodyDataColumns) GoString() string {
	return s.String()
}

func (s *GetHotlineGroupDetailReportResponseBodyDataColumns) SetKey(v string) *GetHotlineGroupDetailReportResponseBodyDataColumns {
	s.Key = &v
	return s
}

func (s *GetHotlineGroupDetailReportResponseBodyDataColumns) SetTitle(v string) *GetHotlineGroupDetailReportResponseBodyDataColumns {
	s.Title = &v
	return s
}

type GetHotlineGroupDetailReportResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetHotlineGroupDetailReportResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetHotlineGroupDetailReportResponse) String() string {
	return tea.Prettify(s)
}

func (s GetHotlineGroupDetailReportResponse) GoString() string {
	return s.String()
}

func (s *GetHotlineGroupDetailReportResponse) SetHeaders(v map[string]*string) *GetHotlineGroupDetailReportResponse {
	s.Headers = v
	return s
}

func (s *GetHotlineGroupDetailReportResponse) SetStatusCode(v int32) *GetHotlineGroupDetailReportResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHotlineGroupDetailReportResponse) SetBody(v *GetHotlineGroupDetailReportResponseBody) *GetHotlineGroupDetailReportResponse {
	s.Body = v
	return s
}

type GetHotlineWaitingNumberRequest struct {
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetHotlineWaitingNumberRequest) String() string {
	return tea.Prettify(s)
}

func (s GetHotlineWaitingNumberRequest) GoString() string {
	return s.String()
}

func (s *GetHotlineWaitingNumberRequest) SetAccountName(v string) *GetHotlineWaitingNumberRequest {
	s.AccountName = &v
	return s
}

func (s *GetHotlineWaitingNumberRequest) SetClientToken(v string) *GetHotlineWaitingNumberRequest {
	s.ClientToken = &v
	return s
}

func (s *GetHotlineWaitingNumberRequest) SetInstanceId(v string) *GetHotlineWaitingNumberRequest {
	s.InstanceId = &v
	return s
}

type GetHotlineWaitingNumberResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *int64  `json:"Data,omitempty" xml:"Data,omitempty"`
	HttpStatusCode *int64  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetHotlineWaitingNumberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetHotlineWaitingNumberResponseBody) GoString() string {
	return s.String()
}

func (s *GetHotlineWaitingNumberResponseBody) SetCode(v string) *GetHotlineWaitingNumberResponseBody {
	s.Code = &v
	return s
}

func (s *GetHotlineWaitingNumberResponseBody) SetData(v int64) *GetHotlineWaitingNumberResponseBody {
	s.Data = &v
	return s
}

func (s *GetHotlineWaitingNumberResponseBody) SetHttpStatusCode(v int64) *GetHotlineWaitingNumberResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetHotlineWaitingNumberResponseBody) SetMessage(v string) *GetHotlineWaitingNumberResponseBody {
	s.Message = &v
	return s
}

func (s *GetHotlineWaitingNumberResponseBody) SetRequestId(v string) *GetHotlineWaitingNumberResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHotlineWaitingNumberResponseBody) SetSuccess(v bool) *GetHotlineWaitingNumberResponseBody {
	s.Success = &v
	return s
}

type GetHotlineWaitingNumberResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetHotlineWaitingNumberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetHotlineWaitingNumberResponse) String() string {
	return tea.Prettify(s)
}

func (s GetHotlineWaitingNumberResponse) GoString() string {
	return s.String()
}

func (s *GetHotlineWaitingNumberResponse) SetHeaders(v map[string]*string) *GetHotlineWaitingNumberResponse {
	s.Headers = v
	return s
}

func (s *GetHotlineWaitingNumberResponse) SetStatusCode(v int32) *GetHotlineWaitingNumberResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHotlineWaitingNumberResponse) SetBody(v *GetHotlineWaitingNumberResponseBody) *GetHotlineWaitingNumberResponse {
	s.Body = v
	return s
}

type GetNumLocationRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PhoneNum    *string `json:"PhoneNum,omitempty" xml:"PhoneNum,omitempty"`
}

func (s GetNumLocationRequest) String() string {
	return tea.Prettify(s)
}

func (s GetNumLocationRequest) GoString() string {
	return s.String()
}

func (s *GetNumLocationRequest) SetClientToken(v string) *GetNumLocationRequest {
	s.ClientToken = &v
	return s
}

func (s *GetNumLocationRequest) SetInstanceId(v string) *GetNumLocationRequest {
	s.InstanceId = &v
	return s
}

func (s *GetNumLocationRequest) SetPhoneNum(v string) *GetNumLocationRequest {
	s.PhoneNum = &v
	return s
}

type GetNumLocationResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *string `json:"Data,omitempty" xml:"Data,omitempty"`
	HttpStatusCode *int64  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetNumLocationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetNumLocationResponseBody) GoString() string {
	return s.String()
}

func (s *GetNumLocationResponseBody) SetCode(v string) *GetNumLocationResponseBody {
	s.Code = &v
	return s
}

func (s *GetNumLocationResponseBody) SetData(v string) *GetNumLocationResponseBody {
	s.Data = &v
	return s
}

func (s *GetNumLocationResponseBody) SetHttpStatusCode(v int64) *GetNumLocationResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetNumLocationResponseBody) SetMessage(v string) *GetNumLocationResponseBody {
	s.Message = &v
	return s
}

func (s *GetNumLocationResponseBody) SetRequestId(v string) *GetNumLocationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetNumLocationResponseBody) SetSuccess(v bool) *GetNumLocationResponseBody {
	s.Success = &v
	return s
}

type GetNumLocationResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetNumLocationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetNumLocationResponse) String() string {
	return tea.Prettify(s)
}

func (s GetNumLocationResponse) GoString() string {
	return s.String()
}

func (s *GetNumLocationResponse) SetHeaders(v map[string]*string) *GetNumLocationResponse {
	s.Headers = v
	return s
}

func (s *GetNumLocationResponse) SetStatusCode(v int32) *GetNumLocationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetNumLocationResponse) SetBody(v *GetNumLocationResponseBody) *GetNumLocationResponse {
	s.Body = v
	return s
}

type GetOutbounNumListRequest struct {
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetOutbounNumListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOutbounNumListRequest) GoString() string {
	return s.String()
}

func (s *GetOutbounNumListRequest) SetAccountName(v string) *GetOutbounNumListRequest {
	s.AccountName = &v
	return s
}

func (s *GetOutbounNumListRequest) SetClientToken(v string) *GetOutbounNumListRequest {
	s.ClientToken = &v
	return s
}

func (s *GetOutbounNumListRequest) SetInstanceId(v string) *GetOutbounNumListRequest {
	s.InstanceId = &v
	return s
}

type GetOutbounNumListResponseBody struct {
	Code           *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *GetOutbounNumListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int64                             `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetOutbounNumListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOutbounNumListResponseBody) GoString() string {
	return s.String()
}

func (s *GetOutbounNumListResponseBody) SetCode(v string) *GetOutbounNumListResponseBody {
	s.Code = &v
	return s
}

func (s *GetOutbounNumListResponseBody) SetData(v *GetOutbounNumListResponseBodyData) *GetOutbounNumListResponseBody {
	s.Data = v
	return s
}

func (s *GetOutbounNumListResponseBody) SetHttpStatusCode(v int64) *GetOutbounNumListResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetOutbounNumListResponseBody) SetMessage(v string) *GetOutbounNumListResponseBody {
	s.Message = &v
	return s
}

func (s *GetOutbounNumListResponseBody) SetRequestId(v string) *GetOutbounNumListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOutbounNumListResponseBody) SetSuccess(v bool) *GetOutbounNumListResponseBody {
	s.Success = &v
	return s
}

type GetOutbounNumListResponseBodyData struct {
	Num      []*GetOutbounNumListResponseBodyDataNum      `json:"Num,omitempty" xml:"Num,omitempty" type:"Repeated"`
	NumGroup []*GetOutbounNumListResponseBodyDataNumGroup `json:"NumGroup,omitempty" xml:"NumGroup,omitempty" type:"Repeated"`
}

func (s GetOutbounNumListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetOutbounNumListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetOutbounNumListResponseBodyData) SetNum(v []*GetOutbounNumListResponseBodyDataNum) *GetOutbounNumListResponseBodyData {
	s.Num = v
	return s
}

func (s *GetOutbounNumListResponseBodyData) SetNumGroup(v []*GetOutbounNumListResponseBodyDataNumGroup) *GetOutbounNumListResponseBodyData {
	s.NumGroup = v
	return s
}

type GetOutbounNumListResponseBodyDataNum struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Type        *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
	Value       *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetOutbounNumListResponseBodyDataNum) String() string {
	return tea.Prettify(s)
}

func (s GetOutbounNumListResponseBodyDataNum) GoString() string {
	return s.String()
}

func (s *GetOutbounNumListResponseBodyDataNum) SetDescription(v string) *GetOutbounNumListResponseBodyDataNum {
	s.Description = &v
	return s
}

func (s *GetOutbounNumListResponseBodyDataNum) SetType(v int32) *GetOutbounNumListResponseBodyDataNum {
	s.Type = &v
	return s
}

func (s *GetOutbounNumListResponseBodyDataNum) SetValue(v string) *GetOutbounNumListResponseBodyDataNum {
	s.Value = &v
	return s
}

type GetOutbounNumListResponseBodyDataNumGroup struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Type        *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
	Value       *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetOutbounNumListResponseBodyDataNumGroup) String() string {
	return tea.Prettify(s)
}

func (s GetOutbounNumListResponseBodyDataNumGroup) GoString() string {
	return s.String()
}

func (s *GetOutbounNumListResponseBodyDataNumGroup) SetDescription(v string) *GetOutbounNumListResponseBodyDataNumGroup {
	s.Description = &v
	return s
}

func (s *GetOutbounNumListResponseBodyDataNumGroup) SetType(v int32) *GetOutbounNumListResponseBodyDataNumGroup {
	s.Type = &v
	return s
}

func (s *GetOutbounNumListResponseBodyDataNumGroup) SetValue(v string) *GetOutbounNumListResponseBodyDataNumGroup {
	s.Value = &v
	return s
}

type GetOutbounNumListResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetOutbounNumListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOutbounNumListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOutbounNumListResponse) GoString() string {
	return s.String()
}

func (s *GetOutbounNumListResponse) SetHeaders(v map[string]*string) *GetOutbounNumListResponse {
	s.Headers = v
	return s
}

func (s *GetOutbounNumListResponse) SetStatusCode(v int32) *GetOutbounNumListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOutbounNumListResponse) SetBody(v *GetOutbounNumListResponseBody) *GetOutbounNumListResponse {
	s.Body = v
	return s
}

type GetOuterCallCenterDataListRequest struct {
	BizId          *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	FromPhoneNum   *string `json:"FromPhoneNum,omitempty" xml:"FromPhoneNum,omitempty"`
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	QueryEndTime   *string `json:"QueryEndTime,omitempty" xml:"QueryEndTime,omitempty"`
	QueryStartTime *string `json:"QueryStartTime,omitempty" xml:"QueryStartTime,omitempty"`
	SessionId      *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	ToPhoneNum     *string `json:"ToPhoneNum,omitempty" xml:"ToPhoneNum,omitempty"`
}

func (s GetOuterCallCenterDataListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOuterCallCenterDataListRequest) GoString() string {
	return s.String()
}

func (s *GetOuterCallCenterDataListRequest) SetBizId(v string) *GetOuterCallCenterDataListRequest {
	s.BizId = &v
	return s
}

func (s *GetOuterCallCenterDataListRequest) SetFromPhoneNum(v string) *GetOuterCallCenterDataListRequest {
	s.FromPhoneNum = &v
	return s
}

func (s *GetOuterCallCenterDataListRequest) SetInstanceId(v string) *GetOuterCallCenterDataListRequest {
	s.InstanceId = &v
	return s
}

func (s *GetOuterCallCenterDataListRequest) SetQueryEndTime(v string) *GetOuterCallCenterDataListRequest {
	s.QueryEndTime = &v
	return s
}

func (s *GetOuterCallCenterDataListRequest) SetQueryStartTime(v string) *GetOuterCallCenterDataListRequest {
	s.QueryStartTime = &v
	return s
}

func (s *GetOuterCallCenterDataListRequest) SetSessionId(v string) *GetOuterCallCenterDataListRequest {
	s.SessionId = &v
	return s
}

func (s *GetOuterCallCenterDataListRequest) SetToPhoneNum(v string) *GetOuterCallCenterDataListRequest {
	s.ToPhoneNum = &v
	return s
}

type GetOuterCallCenterDataListResponseBody struct {
	Code           *string                                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           []*GetOuterCallCenterDataListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	HttpStatusCode *int64                                        `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetOuterCallCenterDataListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOuterCallCenterDataListResponseBody) GoString() string {
	return s.String()
}

func (s *GetOuterCallCenterDataListResponseBody) SetCode(v string) *GetOuterCallCenterDataListResponseBody {
	s.Code = &v
	return s
}

func (s *GetOuterCallCenterDataListResponseBody) SetData(v []*GetOuterCallCenterDataListResponseBodyData) *GetOuterCallCenterDataListResponseBody {
	s.Data = v
	return s
}

func (s *GetOuterCallCenterDataListResponseBody) SetHttpStatusCode(v int64) *GetOuterCallCenterDataListResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetOuterCallCenterDataListResponseBody) SetMessage(v string) *GetOuterCallCenterDataListResponseBody {
	s.Message = &v
	return s
}

func (s *GetOuterCallCenterDataListResponseBody) SetRequestId(v string) *GetOuterCallCenterDataListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOuterCallCenterDataListResponseBody) SetSuccess(v bool) *GetOuterCallCenterDataListResponseBody {
	s.Success = &v
	return s
}

type GetOuterCallCenterDataListResponseBodyData struct {
	Acid          *string `json:"Acid,omitempty" xml:"Acid,omitempty"`
	BizId         *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	BizType       *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	CallType      *string `json:"CallType,omitempty" xml:"CallType,omitempty"`
	EndReason     *string `json:"EndReason,omitempty" xml:"EndReason,omitempty"`
	ExtInfo       *string `json:"ExtInfo,omitempty" xml:"ExtInfo,omitempty"`
	FromPhoneNum  *string `json:"FromPhoneNum,omitempty" xml:"FromPhoneNum,omitempty"`
	InterveneTime *string `json:"InterveneTime,omitempty" xml:"InterveneTime,omitempty"`
	SessionId     *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	ToPhoneNum    *string `json:"ToPhoneNum,omitempty" xml:"ToPhoneNum,omitempty"`
	UserInfo      *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s GetOuterCallCenterDataListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetOuterCallCenterDataListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetOuterCallCenterDataListResponseBodyData) SetAcid(v string) *GetOuterCallCenterDataListResponseBodyData {
	s.Acid = &v
	return s
}

func (s *GetOuterCallCenterDataListResponseBodyData) SetBizId(v string) *GetOuterCallCenterDataListResponseBodyData {
	s.BizId = &v
	return s
}

func (s *GetOuterCallCenterDataListResponseBodyData) SetBizType(v string) *GetOuterCallCenterDataListResponseBodyData {
	s.BizType = &v
	return s
}

func (s *GetOuterCallCenterDataListResponseBodyData) SetCallType(v string) *GetOuterCallCenterDataListResponseBodyData {
	s.CallType = &v
	return s
}

func (s *GetOuterCallCenterDataListResponseBodyData) SetEndReason(v string) *GetOuterCallCenterDataListResponseBodyData {
	s.EndReason = &v
	return s
}

func (s *GetOuterCallCenterDataListResponseBodyData) SetExtInfo(v string) *GetOuterCallCenterDataListResponseBodyData {
	s.ExtInfo = &v
	return s
}

func (s *GetOuterCallCenterDataListResponseBodyData) SetFromPhoneNum(v string) *GetOuterCallCenterDataListResponseBodyData {
	s.FromPhoneNum = &v
	return s
}

func (s *GetOuterCallCenterDataListResponseBodyData) SetInterveneTime(v string) *GetOuterCallCenterDataListResponseBodyData {
	s.InterveneTime = &v
	return s
}

func (s *GetOuterCallCenterDataListResponseBodyData) SetSessionId(v string) *GetOuterCallCenterDataListResponseBodyData {
	s.SessionId = &v
	return s
}

func (s *GetOuterCallCenterDataListResponseBodyData) SetToPhoneNum(v string) *GetOuterCallCenterDataListResponseBodyData {
	s.ToPhoneNum = &v
	return s
}

func (s *GetOuterCallCenterDataListResponseBodyData) SetUserInfo(v string) *GetOuterCallCenterDataListResponseBodyData {
	s.UserInfo = &v
	return s
}

type GetOuterCallCenterDataListResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetOuterCallCenterDataListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetOuterCallCenterDataListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOuterCallCenterDataListResponse) GoString() string {
	return s.String()
}

func (s *GetOuterCallCenterDataListResponse) SetHeaders(v map[string]*string) *GetOuterCallCenterDataListResponse {
	s.Headers = v
	return s
}

func (s *GetOuterCallCenterDataListResponse) SetStatusCode(v int32) *GetOuterCallCenterDataListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOuterCallCenterDataListResponse) SetBody(v *GetOuterCallCenterDataListResponseBody) *GetOuterCallCenterDataListResponse {
	s.Body = v
	return s
}

type GetTagListRequest struct {
	EntityId   *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetTagListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTagListRequest) GoString() string {
	return s.String()
}

func (s *GetTagListRequest) SetEntityId(v string) *GetTagListRequest {
	s.EntityId = &v
	return s
}

func (s *GetTagListRequest) SetEntityType(v string) *GetTagListRequest {
	s.EntityType = &v
	return s
}

func (s *GetTagListRequest) SetInstanceId(v string) *GetTagListRequest {
	s.InstanceId = &v
	return s
}

type GetTagListResponseBody struct {
	Code      *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*GetTagListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetTagListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTagListResponseBody) GoString() string {
	return s.String()
}

func (s *GetTagListResponseBody) SetCode(v string) *GetTagListResponseBody {
	s.Code = &v
	return s
}

func (s *GetTagListResponseBody) SetData(v []*GetTagListResponseBodyData) *GetTagListResponseBody {
	s.Data = v
	return s
}

func (s *GetTagListResponseBody) SetMessage(v string) *GetTagListResponseBody {
	s.Message = &v
	return s
}

func (s *GetTagListResponseBody) SetRequestId(v string) *GetTagListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTagListResponseBody) SetSuccess(v bool) *GetTagListResponseBody {
	s.Success = &v
	return s
}

type GetTagListResponseBodyData struct {
	ScenarioCode *string                                `json:"ScenarioCode,omitempty" xml:"ScenarioCode,omitempty"`
	TagGroupCode *string                                `json:"TagGroupCode,omitempty" xml:"TagGroupCode,omitempty"`
	TagGroupName *string                                `json:"TagGroupName,omitempty" xml:"TagGroupName,omitempty"`
	TagValues    []*GetTagListResponseBodyDataTagValues `json:"TagValues,omitempty" xml:"TagValues,omitempty" type:"Repeated"`
}

func (s GetTagListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetTagListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTagListResponseBodyData) SetScenarioCode(v string) *GetTagListResponseBodyData {
	s.ScenarioCode = &v
	return s
}

func (s *GetTagListResponseBodyData) SetTagGroupCode(v string) *GetTagListResponseBodyData {
	s.TagGroupCode = &v
	return s
}

func (s *GetTagListResponseBodyData) SetTagGroupName(v string) *GetTagListResponseBodyData {
	s.TagGroupName = &v
	return s
}

func (s *GetTagListResponseBodyData) SetTagValues(v []*GetTagListResponseBodyDataTagValues) *GetTagListResponseBodyData {
	s.TagValues = v
	return s
}

type GetTagListResponseBodyDataTagValues struct {
	Description          *string `json:"Description,omitempty" xml:"Description,omitempty"`
	EntityRelationNumber *string `json:"EntityRelationNumber,omitempty" xml:"EntityRelationNumber,omitempty"`
	Status               *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TagCode              *string `json:"TagCode,omitempty" xml:"TagCode,omitempty"`
	TagGroupCode         *string `json:"TagGroupCode,omitempty" xml:"TagGroupCode,omitempty"`
	TagGroupName         *string `json:"TagGroupName,omitempty" xml:"TagGroupName,omitempty"`
	TagName              *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
}

func (s GetTagListResponseBodyDataTagValues) String() string {
	return tea.Prettify(s)
}

func (s GetTagListResponseBodyDataTagValues) GoString() string {
	return s.String()
}

func (s *GetTagListResponseBodyDataTagValues) SetDescription(v string) *GetTagListResponseBodyDataTagValues {
	s.Description = &v
	return s
}

func (s *GetTagListResponseBodyDataTagValues) SetEntityRelationNumber(v string) *GetTagListResponseBodyDataTagValues {
	s.EntityRelationNumber = &v
	return s
}

func (s *GetTagListResponseBodyDataTagValues) SetStatus(v string) *GetTagListResponseBodyDataTagValues {
	s.Status = &v
	return s
}

func (s *GetTagListResponseBodyDataTagValues) SetTagCode(v string) *GetTagListResponseBodyDataTagValues {
	s.TagCode = &v
	return s
}

func (s *GetTagListResponseBodyDataTagValues) SetTagGroupCode(v string) *GetTagListResponseBodyDataTagValues {
	s.TagGroupCode = &v
	return s
}

func (s *GetTagListResponseBodyDataTagValues) SetTagGroupName(v string) *GetTagListResponseBodyDataTagValues {
	s.TagGroupName = &v
	return s
}

func (s *GetTagListResponseBodyDataTagValues) SetTagName(v string) *GetTagListResponseBodyDataTagValues {
	s.TagName = &v
	return s
}

type GetTagListResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetTagListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTagListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTagListResponse) GoString() string {
	return s.String()
}

func (s *GetTagListResponse) SetHeaders(v map[string]*string) *GetTagListResponse {
	s.Headers = v
	return s
}

func (s *GetTagListResponse) SetStatusCode(v int32) *GetTagListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTagListResponse) SetBody(v *GetTagListResponseBody) *GetTagListResponse {
	s.Body = v
	return s
}

type GetTicketByCaseIdRequest struct {
	CaseId     *int64  `json:"CaseId,omitempty" xml:"CaseId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetTicketByCaseIdRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTicketByCaseIdRequest) GoString() string {
	return s.String()
}

func (s *GetTicketByCaseIdRequest) SetCaseId(v int64) *GetTicketByCaseIdRequest {
	s.CaseId = &v
	return s
}

func (s *GetTicketByCaseIdRequest) SetInstanceId(v string) *GetTicketByCaseIdRequest {
	s.InstanceId = &v
	return s
}

type GetTicketByCaseIdResponseBody struct {
	Code      *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetTicketByCaseIdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetTicketByCaseIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTicketByCaseIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetTicketByCaseIdResponseBody) SetCode(v string) *GetTicketByCaseIdResponseBody {
	s.Code = &v
	return s
}

func (s *GetTicketByCaseIdResponseBody) SetData(v *GetTicketByCaseIdResponseBodyData) *GetTicketByCaseIdResponseBody {
	s.Data = v
	return s
}

func (s *GetTicketByCaseIdResponseBody) SetMessage(v string) *GetTicketByCaseIdResponseBody {
	s.Message = &v
	return s
}

func (s *GetTicketByCaseIdResponseBody) SetRequestId(v string) *GetTicketByCaseIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTicketByCaseIdResponseBody) SetSuccess(v bool) *GetTicketByCaseIdResponseBody {
	s.Success = &v
	return s
}

type GetTicketByCaseIdResponseBodyData struct {
	BuId         *int64                 `json:"BuId,omitempty" xml:"BuId,omitempty"`
	CaseId       *int64                 `json:"CaseId,omitempty" xml:"CaseId,omitempty"`
	CaseStatus   *int32                 `json:"CaseStatus,omitempty" xml:"CaseStatus,omitempty"`
	CaseType     *int32                 `json:"CaseType,omitempty" xml:"CaseType,omitempty"`
	ChannelId    *string                `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	DepartmentId *int64                 `json:"DepartmentId,omitempty" xml:"DepartmentId,omitempty"`
	EndTime      *int64                 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	ExtAttrs     map[string]interface{} `json:"ExtAttrs,omitempty" xml:"ExtAttrs,omitempty"`
	FromInfo     *string                `json:"FromInfo,omitempty" xml:"FromInfo,omitempty"`
	GmtCreate    *int64                 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified  *int64                 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	MemberId     *int64                 `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	MemberName   *string                `json:"MemberName,omitempty" xml:"MemberName,omitempty"`
	Owner        *int64                 `json:"Owner,omitempty" xml:"Owner,omitempty"`
	OwnerName    *string                `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	ParentId     *int64                 `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	Priority     *int32                 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	QuestionId   *string                `json:"QuestionId,omitempty" xml:"QuestionId,omitempty"`
	QuestionInfo *string                `json:"QuestionInfo,omitempty" xml:"QuestionInfo,omitempty"`
	SopCateId    *int64                 `json:"SopCateId,omitempty" xml:"SopCateId,omitempty"`
	SrType       *int64                 `json:"SrType,omitempty" xml:"SrType,omitempty"`
}

func (s GetTicketByCaseIdResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetTicketByCaseIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTicketByCaseIdResponseBodyData) SetBuId(v int64) *GetTicketByCaseIdResponseBodyData {
	s.BuId = &v
	return s
}

func (s *GetTicketByCaseIdResponseBodyData) SetCaseId(v int64) *GetTicketByCaseIdResponseBodyData {
	s.CaseId = &v
	return s
}

func (s *GetTicketByCaseIdResponseBodyData) SetCaseStatus(v int32) *GetTicketByCaseIdResponseBodyData {
	s.CaseStatus = &v
	return s
}

func (s *GetTicketByCaseIdResponseBodyData) SetCaseType(v int32) *GetTicketByCaseIdResponseBodyData {
	s.CaseType = &v
	return s
}

func (s *GetTicketByCaseIdResponseBodyData) SetChannelId(v string) *GetTicketByCaseIdResponseBodyData {
	s.ChannelId = &v
	return s
}

func (s *GetTicketByCaseIdResponseBodyData) SetDepartmentId(v int64) *GetTicketByCaseIdResponseBodyData {
	s.DepartmentId = &v
	return s
}

func (s *GetTicketByCaseIdResponseBodyData) SetEndTime(v int64) *GetTicketByCaseIdResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *GetTicketByCaseIdResponseBodyData) SetExtAttrs(v map[string]interface{}) *GetTicketByCaseIdResponseBodyData {
	s.ExtAttrs = v
	return s
}

func (s *GetTicketByCaseIdResponseBodyData) SetFromInfo(v string) *GetTicketByCaseIdResponseBodyData {
	s.FromInfo = &v
	return s
}

func (s *GetTicketByCaseIdResponseBodyData) SetGmtCreate(v int64) *GetTicketByCaseIdResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *GetTicketByCaseIdResponseBodyData) SetGmtModified(v int64) *GetTicketByCaseIdResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *GetTicketByCaseIdResponseBodyData) SetMemberId(v int64) *GetTicketByCaseIdResponseBodyData {
	s.MemberId = &v
	return s
}

func (s *GetTicketByCaseIdResponseBodyData) SetMemberName(v string) *GetTicketByCaseIdResponseBodyData {
	s.MemberName = &v
	return s
}

func (s *GetTicketByCaseIdResponseBodyData) SetOwner(v int64) *GetTicketByCaseIdResponseBodyData {
	s.Owner = &v
	return s
}

func (s *GetTicketByCaseIdResponseBodyData) SetOwnerName(v string) *GetTicketByCaseIdResponseBodyData {
	s.OwnerName = &v
	return s
}

func (s *GetTicketByCaseIdResponseBodyData) SetParentId(v int64) *GetTicketByCaseIdResponseBodyData {
	s.ParentId = &v
	return s
}

func (s *GetTicketByCaseIdResponseBodyData) SetPriority(v int32) *GetTicketByCaseIdResponseBodyData {
	s.Priority = &v
	return s
}

func (s *GetTicketByCaseIdResponseBodyData) SetQuestionId(v string) *GetTicketByCaseIdResponseBodyData {
	s.QuestionId = &v
	return s
}

func (s *GetTicketByCaseIdResponseBodyData) SetQuestionInfo(v string) *GetTicketByCaseIdResponseBodyData {
	s.QuestionInfo = &v
	return s
}

func (s *GetTicketByCaseIdResponseBodyData) SetSopCateId(v int64) *GetTicketByCaseIdResponseBodyData {
	s.SopCateId = &v
	return s
}

func (s *GetTicketByCaseIdResponseBodyData) SetSrType(v int64) *GetTicketByCaseIdResponseBodyData {
	s.SrType = &v
	return s
}

type GetTicketByCaseIdResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetTicketByCaseIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTicketByCaseIdResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTicketByCaseIdResponse) GoString() string {
	return s.String()
}

func (s *GetTicketByCaseIdResponse) SetHeaders(v map[string]*string) *GetTicketByCaseIdResponse {
	s.Headers = v
	return s
}

func (s *GetTicketByCaseIdResponse) SetStatusCode(v int32) *GetTicketByCaseIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTicketByCaseIdResponse) SetBody(v *GetTicketByCaseIdResponseBody) *GetTicketByCaseIdResponse {
	s.Body = v
	return s
}

type GetTicketTemplateSchemaRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	TemplateId  *int64  `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s GetTicketTemplateSchemaRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTicketTemplateSchemaRequest) GoString() string {
	return s.String()
}

func (s *GetTicketTemplateSchemaRequest) SetClientToken(v string) *GetTicketTemplateSchemaRequest {
	s.ClientToken = &v
	return s
}

func (s *GetTicketTemplateSchemaRequest) SetInstanceId(v string) *GetTicketTemplateSchemaRequest {
	s.InstanceId = &v
	return s
}

func (s *GetTicketTemplateSchemaRequest) SetTemplateId(v int64) *GetTicketTemplateSchemaRequest {
	s.TemplateId = &v
	return s
}

type GetTicketTemplateSchemaResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetTicketTemplateSchemaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTicketTemplateSchemaResponseBody) GoString() string {
	return s.String()
}

func (s *GetTicketTemplateSchemaResponseBody) SetCode(v string) *GetTicketTemplateSchemaResponseBody {
	s.Code = &v
	return s
}

func (s *GetTicketTemplateSchemaResponseBody) SetData(v string) *GetTicketTemplateSchemaResponseBody {
	s.Data = &v
	return s
}

func (s *GetTicketTemplateSchemaResponseBody) SetMessage(v string) *GetTicketTemplateSchemaResponseBody {
	s.Message = &v
	return s
}

func (s *GetTicketTemplateSchemaResponseBody) SetRequestId(v string) *GetTicketTemplateSchemaResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTicketTemplateSchemaResponseBody) SetSuccess(v bool) *GetTicketTemplateSchemaResponseBody {
	s.Success = &v
	return s
}

type GetTicketTemplateSchemaResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetTicketTemplateSchemaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTicketTemplateSchemaResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTicketTemplateSchemaResponse) GoString() string {
	return s.String()
}

func (s *GetTicketTemplateSchemaResponse) SetHeaders(v map[string]*string) *GetTicketTemplateSchemaResponse {
	s.Headers = v
	return s
}

func (s *GetTicketTemplateSchemaResponse) SetStatusCode(v int32) *GetTicketTemplateSchemaResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTicketTemplateSchemaResponse) SetBody(v *GetTicketTemplateSchemaResponseBody) *GetTicketTemplateSchemaResponse {
	s.Body = v
	return s
}

type GetUserTokenRequest struct {
	AppKey     *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Nick       *string `json:"Nick,omitempty" xml:"Nick,omitempty"`
	UserId     *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetUserTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUserTokenRequest) GoString() string {
	return s.String()
}

func (s *GetUserTokenRequest) SetAppKey(v string) *GetUserTokenRequest {
	s.AppKey = &v
	return s
}

func (s *GetUserTokenRequest) SetInstanceId(v string) *GetUserTokenRequest {
	s.InstanceId = &v
	return s
}

func (s *GetUserTokenRequest) SetNick(v string) *GetUserTokenRequest {
	s.Nick = &v
	return s
}

func (s *GetUserTokenRequest) SetUserId(v string) *GetUserTokenRequest {
	s.UserId = &v
	return s
}

type GetUserTokenResponseBody struct {
	Code      *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetUserTokenResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetUserTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserTokenResponseBody) SetCode(v string) *GetUserTokenResponseBody {
	s.Code = &v
	return s
}

func (s *GetUserTokenResponseBody) SetData(v *GetUserTokenResponseBodyData) *GetUserTokenResponseBody {
	s.Data = v
	return s
}

func (s *GetUserTokenResponseBody) SetMessage(v string) *GetUserTokenResponseBody {
	s.Message = &v
	return s
}

func (s *GetUserTokenResponseBody) SetRequestId(v string) *GetUserTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserTokenResponseBody) SetSuccess(v bool) *GetUserTokenResponseBody {
	s.Success = &v
	return s
}

type GetUserTokenResponseBodyData struct {
	Expires *int64  `json:"Expires,omitempty" xml:"Expires,omitempty"`
	Token   *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s GetUserTokenResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetUserTokenResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetUserTokenResponseBodyData) SetExpires(v int64) *GetUserTokenResponseBodyData {
	s.Expires = &v
	return s
}

func (s *GetUserTokenResponseBodyData) SetToken(v string) *GetUserTokenResponseBodyData {
	s.Token = &v
	return s
}

type GetUserTokenResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetUserTokenResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetUserTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserTokenResponse) GoString() string {
	return s.String()
}

func (s *GetUserTokenResponse) SetHeaders(v map[string]*string) *GetUserTokenResponse {
	s.Headers = v
	return s
}

func (s *GetUserTokenResponse) SetStatusCode(v int32) *GetUserTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserTokenResponse) SetBody(v *GetUserTokenResponseBody) *GetUserTokenResponse {
	s.Body = v
	return s
}

type GrantRolesRequest struct {
	AccountName *string  `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	ClientToken *string  `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId  *string  `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Operator    *string  `json:"Operator,omitempty" xml:"Operator,omitempty"`
	RoleId      []*int64 `json:"RoleId,omitempty" xml:"RoleId,omitempty" type:"Repeated"`
}

func (s GrantRolesRequest) String() string {
	return tea.Prettify(s)
}

func (s GrantRolesRequest) GoString() string {
	return s.String()
}

func (s *GrantRolesRequest) SetAccountName(v string) *GrantRolesRequest {
	s.AccountName = &v
	return s
}

func (s *GrantRolesRequest) SetClientToken(v string) *GrantRolesRequest {
	s.ClientToken = &v
	return s
}

func (s *GrantRolesRequest) SetInstanceId(v string) *GrantRolesRequest {
	s.InstanceId = &v
	return s
}

func (s *GrantRolesRequest) SetOperator(v string) *GrantRolesRequest {
	s.Operator = &v
	return s
}

func (s *GrantRolesRequest) SetRoleId(v []*int64) *GrantRolesRequest {
	s.RoleId = v
	return s
}

type GrantRolesResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int64  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GrantRolesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GrantRolesResponseBody) GoString() string {
	return s.String()
}

func (s *GrantRolesResponseBody) SetCode(v string) *GrantRolesResponseBody {
	s.Code = &v
	return s
}

func (s *GrantRolesResponseBody) SetHttpStatusCode(v int64) *GrantRolesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GrantRolesResponseBody) SetMessage(v string) *GrantRolesResponseBody {
	s.Message = &v
	return s
}

func (s *GrantRolesResponseBody) SetRequestId(v string) *GrantRolesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GrantRolesResponseBody) SetSuccess(v bool) *GrantRolesResponseBody {
	s.Success = &v
	return s
}

type GrantRolesResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GrantRolesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GrantRolesResponse) String() string {
	return tea.Prettify(s)
}

func (s GrantRolesResponse) GoString() string {
	return s.String()
}

func (s *GrantRolesResponse) SetHeaders(v map[string]*string) *GrantRolesResponse {
	s.Headers = v
	return s
}

func (s *GrantRolesResponse) SetStatusCode(v int32) *GrantRolesResponse {
	s.StatusCode = &v
	return s
}

func (s *GrantRolesResponse) SetBody(v *GrantRolesResponseBody) *GrantRolesResponse {
	s.Body = v
	return s
}

type HangUpAgentMonitorRequest struct {
	AccountName       *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	CallId            *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	CallerParentId    *int64  `json:"CallerParentId,omitempty" xml:"CallerParentId,omitempty"`
	CallerType        *string `json:"CallerType,omitempty" xml:"CallerType,omitempty"`
	CallerUid         *int64  `json:"CallerUid,omitempty" xml:"CallerUid,omitempty"`
	ClientToken       *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ConnectionId      *string `json:"ConnectionId,omitempty" xml:"ConnectionId,omitempty"`
	InstanceId        *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	JobId             *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	RequestId         *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StsTokenCallerUid *int64  `json:"StsTokenCallerUid,omitempty" xml:"StsTokenCallerUid,omitempty"`
}

func (s HangUpAgentMonitorRequest) String() string {
	return tea.Prettify(s)
}

func (s HangUpAgentMonitorRequest) GoString() string {
	return s.String()
}

func (s *HangUpAgentMonitorRequest) SetAccountName(v string) *HangUpAgentMonitorRequest {
	s.AccountName = &v
	return s
}

func (s *HangUpAgentMonitorRequest) SetCallId(v string) *HangUpAgentMonitorRequest {
	s.CallId = &v
	return s
}

func (s *HangUpAgentMonitorRequest) SetCallerParentId(v int64) *HangUpAgentMonitorRequest {
	s.CallerParentId = &v
	return s
}

func (s *HangUpAgentMonitorRequest) SetCallerType(v string) *HangUpAgentMonitorRequest {
	s.CallerType = &v
	return s
}

func (s *HangUpAgentMonitorRequest) SetCallerUid(v int64) *HangUpAgentMonitorRequest {
	s.CallerUid = &v
	return s
}

func (s *HangUpAgentMonitorRequest) SetClientToken(v string) *HangUpAgentMonitorRequest {
	s.ClientToken = &v
	return s
}

func (s *HangUpAgentMonitorRequest) SetConnectionId(v string) *HangUpAgentMonitorRequest {
	s.ConnectionId = &v
	return s
}

func (s *HangUpAgentMonitorRequest) SetInstanceId(v string) *HangUpAgentMonitorRequest {
	s.InstanceId = &v
	return s
}

func (s *HangUpAgentMonitorRequest) SetJobId(v string) *HangUpAgentMonitorRequest {
	s.JobId = &v
	return s
}

func (s *HangUpAgentMonitorRequest) SetRequestId(v string) *HangUpAgentMonitorRequest {
	s.RequestId = &v
	return s
}

func (s *HangUpAgentMonitorRequest) SetStsTokenCallerUid(v int64) *HangUpAgentMonitorRequest {
	s.StsTokenCallerUid = &v
	return s
}

type HangUpAgentMonitorResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s HangUpAgentMonitorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s HangUpAgentMonitorResponseBody) GoString() string {
	return s.String()
}

func (s *HangUpAgentMonitorResponseBody) SetCode(v string) *HangUpAgentMonitorResponseBody {
	s.Code = &v
	return s
}

func (s *HangUpAgentMonitorResponseBody) SetMessage(v string) *HangUpAgentMonitorResponseBody {
	s.Message = &v
	return s
}

func (s *HangUpAgentMonitorResponseBody) SetRequestId(v string) *HangUpAgentMonitorResponseBody {
	s.RequestId = &v
	return s
}

func (s *HangUpAgentMonitorResponseBody) SetSuccess(v bool) *HangUpAgentMonitorResponseBody {
	s.Success = &v
	return s
}

type HangUpAgentMonitorResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *HangUpAgentMonitorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s HangUpAgentMonitorResponse) String() string {
	return tea.Prettify(s)
}

func (s HangUpAgentMonitorResponse) GoString() string {
	return s.String()
}

func (s *HangUpAgentMonitorResponse) SetHeaders(v map[string]*string) *HangUpAgentMonitorResponse {
	s.Headers = v
	return s
}

func (s *HangUpAgentMonitorResponse) SetStatusCode(v int32) *HangUpAgentMonitorResponse {
	s.StatusCode = &v
	return s
}

func (s *HangUpAgentMonitorResponse) SetBody(v *HangUpAgentMonitorResponseBody) *HangUpAgentMonitorResponse {
	s.Body = v
	return s
}

type HangupCallRequest struct {
	AccountName  *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	CallId       *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ConnectionId *string `json:"ConnectionId,omitempty" xml:"ConnectionId,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	JobId        *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s HangupCallRequest) String() string {
	return tea.Prettify(s)
}

func (s HangupCallRequest) GoString() string {
	return s.String()
}

func (s *HangupCallRequest) SetAccountName(v string) *HangupCallRequest {
	s.AccountName = &v
	return s
}

func (s *HangupCallRequest) SetCallId(v string) *HangupCallRequest {
	s.CallId = &v
	return s
}

func (s *HangupCallRequest) SetClientToken(v string) *HangupCallRequest {
	s.ClientToken = &v
	return s
}

func (s *HangupCallRequest) SetConnectionId(v string) *HangupCallRequest {
	s.ConnectionId = &v
	return s
}

func (s *HangupCallRequest) SetInstanceId(v string) *HangupCallRequest {
	s.InstanceId = &v
	return s
}

func (s *HangupCallRequest) SetJobId(v string) *HangupCallRequest {
	s.JobId = &v
	return s
}

type HangupCallResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s HangupCallResponseBody) String() string {
	return tea.Prettify(s)
}

func (s HangupCallResponseBody) GoString() string {
	return s.String()
}

func (s *HangupCallResponseBody) SetCode(v string) *HangupCallResponseBody {
	s.Code = &v
	return s
}

func (s *HangupCallResponseBody) SetMessage(v string) *HangupCallResponseBody {
	s.Message = &v
	return s
}

func (s *HangupCallResponseBody) SetRequestId(v string) *HangupCallResponseBody {
	s.RequestId = &v
	return s
}

func (s *HangupCallResponseBody) SetSuccess(v bool) *HangupCallResponseBody {
	s.Success = &v
	return s
}

type HangupCallResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *HangupCallResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s HangupCallResponse) String() string {
	return tea.Prettify(s)
}

func (s HangupCallResponse) GoString() string {
	return s.String()
}

func (s *HangupCallResponse) SetHeaders(v map[string]*string) *HangupCallResponse {
	s.Headers = v
	return s
}

func (s *HangupCallResponse) SetStatusCode(v int32) *HangupCallResponse {
	s.StatusCode = &v
	return s
}

func (s *HangupCallResponse) SetBody(v *HangupCallResponseBody) *HangupCallResponse {
	s.Body = v
	return s
}

type HangupThirdCallRequest struct {
	AccountName  *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	CallId       *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ConnectionId *string `json:"ConnectionId,omitempty" xml:"ConnectionId,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	JobId        *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s HangupThirdCallRequest) String() string {
	return tea.Prettify(s)
}

func (s HangupThirdCallRequest) GoString() string {
	return s.String()
}

func (s *HangupThirdCallRequest) SetAccountName(v string) *HangupThirdCallRequest {
	s.AccountName = &v
	return s
}

func (s *HangupThirdCallRequest) SetCallId(v string) *HangupThirdCallRequest {
	s.CallId = &v
	return s
}

func (s *HangupThirdCallRequest) SetClientToken(v string) *HangupThirdCallRequest {
	s.ClientToken = &v
	return s
}

func (s *HangupThirdCallRequest) SetConnectionId(v string) *HangupThirdCallRequest {
	s.ConnectionId = &v
	return s
}

func (s *HangupThirdCallRequest) SetInstanceId(v string) *HangupThirdCallRequest {
	s.InstanceId = &v
	return s
}

func (s *HangupThirdCallRequest) SetJobId(v string) *HangupThirdCallRequest {
	s.JobId = &v
	return s
}

type HangupThirdCallResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s HangupThirdCallResponseBody) String() string {
	return tea.Prettify(s)
}

func (s HangupThirdCallResponseBody) GoString() string {
	return s.String()
}

func (s *HangupThirdCallResponseBody) SetCode(v string) *HangupThirdCallResponseBody {
	s.Code = &v
	return s
}

func (s *HangupThirdCallResponseBody) SetMessage(v string) *HangupThirdCallResponseBody {
	s.Message = &v
	return s
}

func (s *HangupThirdCallResponseBody) SetRequestId(v string) *HangupThirdCallResponseBody {
	s.RequestId = &v
	return s
}

func (s *HangupThirdCallResponseBody) SetSuccess(v bool) *HangupThirdCallResponseBody {
	s.Success = &v
	return s
}

type HangupThirdCallResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *HangupThirdCallResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s HangupThirdCallResponse) String() string {
	return tea.Prettify(s)
}

func (s HangupThirdCallResponse) GoString() string {
	return s.String()
}

func (s *HangupThirdCallResponse) SetHeaders(v map[string]*string) *HangupThirdCallResponse {
	s.Headers = v
	return s
}

func (s *HangupThirdCallResponse) SetStatusCode(v int32) *HangupThirdCallResponse {
	s.StatusCode = &v
	return s
}

func (s *HangupThirdCallResponse) SetBody(v *HangupThirdCallResponseBody) *HangupThirdCallResponse {
	s.Body = v
	return s
}

type HoldCallRequest struct {
	AccountName  *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	CallId       *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ConnectionId *string `json:"ConnectionId,omitempty" xml:"ConnectionId,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	JobId        *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s HoldCallRequest) String() string {
	return tea.Prettify(s)
}

func (s HoldCallRequest) GoString() string {
	return s.String()
}

func (s *HoldCallRequest) SetAccountName(v string) *HoldCallRequest {
	s.AccountName = &v
	return s
}

func (s *HoldCallRequest) SetCallId(v string) *HoldCallRequest {
	s.CallId = &v
	return s
}

func (s *HoldCallRequest) SetClientToken(v string) *HoldCallRequest {
	s.ClientToken = &v
	return s
}

func (s *HoldCallRequest) SetConnectionId(v string) *HoldCallRequest {
	s.ConnectionId = &v
	return s
}

func (s *HoldCallRequest) SetInstanceId(v string) *HoldCallRequest {
	s.InstanceId = &v
	return s
}

func (s *HoldCallRequest) SetJobId(v string) *HoldCallRequest {
	s.JobId = &v
	return s
}

type HoldCallResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s HoldCallResponseBody) String() string {
	return tea.Prettify(s)
}

func (s HoldCallResponseBody) GoString() string {
	return s.String()
}

func (s *HoldCallResponseBody) SetCode(v string) *HoldCallResponseBody {
	s.Code = &v
	return s
}

func (s *HoldCallResponseBody) SetMessage(v string) *HoldCallResponseBody {
	s.Message = &v
	return s
}

func (s *HoldCallResponseBody) SetRequestId(v string) *HoldCallResponseBody {
	s.RequestId = &v
	return s
}

func (s *HoldCallResponseBody) SetSuccess(v bool) *HoldCallResponseBody {
	s.Success = &v
	return s
}

type HoldCallResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *HoldCallResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s HoldCallResponse) String() string {
	return tea.Prettify(s)
}

func (s HoldCallResponse) GoString() string {
	return s.String()
}

func (s *HoldCallResponse) SetHeaders(v map[string]*string) *HoldCallResponse {
	s.Headers = v
	return s
}

func (s *HoldCallResponse) SetStatusCode(v int32) *HoldCallResponse {
	s.StatusCode = &v
	return s
}

func (s *HoldCallResponse) SetBody(v *HoldCallResponseBody) *HoldCallResponse {
	s.Body = v
	return s
}

type JoinThirdCallRequest struct {
	AccountName      *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	CallId           *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	ClientToken      *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ConnectionId     *string `json:"ConnectionId,omitempty" xml:"ConnectionId,omitempty"`
	HoldConnectionId *string `json:"HoldConnectionId,omitempty" xml:"HoldConnectionId,omitempty"`
	InstanceId       *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	JobId            *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s JoinThirdCallRequest) String() string {
	return tea.Prettify(s)
}

func (s JoinThirdCallRequest) GoString() string {
	return s.String()
}

func (s *JoinThirdCallRequest) SetAccountName(v string) *JoinThirdCallRequest {
	s.AccountName = &v
	return s
}

func (s *JoinThirdCallRequest) SetCallId(v string) *JoinThirdCallRequest {
	s.CallId = &v
	return s
}

func (s *JoinThirdCallRequest) SetClientToken(v string) *JoinThirdCallRequest {
	s.ClientToken = &v
	return s
}

func (s *JoinThirdCallRequest) SetConnectionId(v string) *JoinThirdCallRequest {
	s.ConnectionId = &v
	return s
}

func (s *JoinThirdCallRequest) SetHoldConnectionId(v string) *JoinThirdCallRequest {
	s.HoldConnectionId = &v
	return s
}

func (s *JoinThirdCallRequest) SetInstanceId(v string) *JoinThirdCallRequest {
	s.InstanceId = &v
	return s
}

func (s *JoinThirdCallRequest) SetJobId(v string) *JoinThirdCallRequest {
	s.JobId = &v
	return s
}

type JoinThirdCallResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s JoinThirdCallResponseBody) String() string {
	return tea.Prettify(s)
}

func (s JoinThirdCallResponseBody) GoString() string {
	return s.String()
}

func (s *JoinThirdCallResponseBody) SetCode(v string) *JoinThirdCallResponseBody {
	s.Code = &v
	return s
}

func (s *JoinThirdCallResponseBody) SetMessage(v string) *JoinThirdCallResponseBody {
	s.Message = &v
	return s
}

func (s *JoinThirdCallResponseBody) SetRequestId(v string) *JoinThirdCallResponseBody {
	s.RequestId = &v
	return s
}

func (s *JoinThirdCallResponseBody) SetSuccess(v bool) *JoinThirdCallResponseBody {
	s.Success = &v
	return s
}

type JoinThirdCallResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *JoinThirdCallResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s JoinThirdCallResponse) String() string {
	return tea.Prettify(s)
}

func (s JoinThirdCallResponse) GoString() string {
	return s.String()
}

func (s *JoinThirdCallResponse) SetHeaders(v map[string]*string) *JoinThirdCallResponse {
	s.Headers = v
	return s
}

func (s *JoinThirdCallResponse) SetStatusCode(v int32) *JoinThirdCallResponse {
	s.StatusCode = &v
	return s
}

func (s *JoinThirdCallResponse) SetBody(v *JoinThirdCallResponseBody) *JoinThirdCallResponse {
	s.Body = v
	return s
}

type ListAgentBySkillGroupIdRequest struct {
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	SkillGroupId *int64  `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
}

func (s ListAgentBySkillGroupIdRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAgentBySkillGroupIdRequest) GoString() string {
	return s.String()
}

func (s *ListAgentBySkillGroupIdRequest) SetClientToken(v string) *ListAgentBySkillGroupIdRequest {
	s.ClientToken = &v
	return s
}

func (s *ListAgentBySkillGroupIdRequest) SetInstanceId(v string) *ListAgentBySkillGroupIdRequest {
	s.InstanceId = &v
	return s
}

func (s *ListAgentBySkillGroupIdRequest) SetSkillGroupId(v int64) *ListAgentBySkillGroupIdRequest {
	s.SkillGroupId = &v
	return s
}

type ListAgentBySkillGroupIdResponseBody struct {
	Code      *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*ListAgentBySkillGroupIdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                      `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListAgentBySkillGroupIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAgentBySkillGroupIdResponseBody) GoString() string {
	return s.String()
}

func (s *ListAgentBySkillGroupIdResponseBody) SetCode(v string) *ListAgentBySkillGroupIdResponseBody {
	s.Code = &v
	return s
}

func (s *ListAgentBySkillGroupIdResponseBody) SetData(v []*ListAgentBySkillGroupIdResponseBodyData) *ListAgentBySkillGroupIdResponseBody {
	s.Data = v
	return s
}

func (s *ListAgentBySkillGroupIdResponseBody) SetMessage(v string) *ListAgentBySkillGroupIdResponseBody {
	s.Message = &v
	return s
}

func (s *ListAgentBySkillGroupIdResponseBody) SetRequestId(v string) *ListAgentBySkillGroupIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAgentBySkillGroupIdResponseBody) SetSuccess(v bool) *ListAgentBySkillGroupIdResponseBody {
	s.Success = &v
	return s
}

type ListAgentBySkillGroupIdResponseBodyData struct {
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	AgentId     *int64  `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	Status      *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	TenantId    *int64  `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s ListAgentBySkillGroupIdResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListAgentBySkillGroupIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAgentBySkillGroupIdResponseBodyData) SetAccountName(v string) *ListAgentBySkillGroupIdResponseBodyData {
	s.AccountName = &v
	return s
}

func (s *ListAgentBySkillGroupIdResponseBodyData) SetAgentId(v int64) *ListAgentBySkillGroupIdResponseBodyData {
	s.AgentId = &v
	return s
}

func (s *ListAgentBySkillGroupIdResponseBodyData) SetDisplayName(v string) *ListAgentBySkillGroupIdResponseBodyData {
	s.DisplayName = &v
	return s
}

func (s *ListAgentBySkillGroupIdResponseBodyData) SetStatus(v int32) *ListAgentBySkillGroupIdResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListAgentBySkillGroupIdResponseBodyData) SetTenantId(v int64) *ListAgentBySkillGroupIdResponseBodyData {
	s.TenantId = &v
	return s
}

type ListAgentBySkillGroupIdResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListAgentBySkillGroupIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAgentBySkillGroupIdResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAgentBySkillGroupIdResponse) GoString() string {
	return s.String()
}

func (s *ListAgentBySkillGroupIdResponse) SetHeaders(v map[string]*string) *ListAgentBySkillGroupIdResponse {
	s.Headers = v
	return s
}

func (s *ListAgentBySkillGroupIdResponse) SetStatusCode(v int32) *ListAgentBySkillGroupIdResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAgentBySkillGroupIdResponse) SetBody(v *ListAgentBySkillGroupIdResponseBody) *ListAgentBySkillGroupIdResponse {
	s.Body = v
	return s
}

type ListAllHotLineSkillGroupsRequest struct {
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ListAllHotLineSkillGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAllHotLineSkillGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListAllHotLineSkillGroupsRequest) SetAccountName(v string) *ListAllHotLineSkillGroupsRequest {
	s.AccountName = &v
	return s
}

func (s *ListAllHotLineSkillGroupsRequest) SetInstanceId(v string) *ListAllHotLineSkillGroupsRequest {
	s.InstanceId = &v
	return s
}

type ListAllHotLineSkillGroupsResponseBody struct {
	Code      *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*ListAllHotLineSkillGroupsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                        `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListAllHotLineSkillGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAllHotLineSkillGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAllHotLineSkillGroupsResponseBody) SetCode(v string) *ListAllHotLineSkillGroupsResponseBody {
	s.Code = &v
	return s
}

func (s *ListAllHotLineSkillGroupsResponseBody) SetData(v []*ListAllHotLineSkillGroupsResponseBodyData) *ListAllHotLineSkillGroupsResponseBody {
	s.Data = v
	return s
}

func (s *ListAllHotLineSkillGroupsResponseBody) SetMessage(v string) *ListAllHotLineSkillGroupsResponseBody {
	s.Message = &v
	return s
}

func (s *ListAllHotLineSkillGroupsResponseBody) SetRequestId(v string) *ListAllHotLineSkillGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAllHotLineSkillGroupsResponseBody) SetSuccess(v bool) *ListAllHotLineSkillGroupsResponseBody {
	s.Success = &v
	return s
}

type ListAllHotLineSkillGroupsResponseBodyData struct {
	DepGroup         *ListAllHotLineSkillGroupsResponseBodyDataDepGroup           `json:"DepGroup,omitempty" xml:"DepGroup,omitempty" type:"Struct"`
	SkillGroupAgents []*ListAllHotLineSkillGroupsResponseBodyDataSkillGroupAgents `json:"SkillGroupAgents,omitempty" xml:"SkillGroupAgents,omitempty" type:"Repeated"`
}

func (s ListAllHotLineSkillGroupsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListAllHotLineSkillGroupsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAllHotLineSkillGroupsResponseBodyData) SetDepGroup(v *ListAllHotLineSkillGroupsResponseBodyDataDepGroup) *ListAllHotLineSkillGroupsResponseBodyData {
	s.DepGroup = v
	return s
}

func (s *ListAllHotLineSkillGroupsResponseBodyData) SetSkillGroupAgents(v []*ListAllHotLineSkillGroupsResponseBodyDataSkillGroupAgents) *ListAllHotLineSkillGroupsResponseBodyData {
	s.SkillGroupAgents = v
	return s
}

type ListAllHotLineSkillGroupsResponseBodyDataDepGroup struct {
	DepGroupId   *int64  `json:"DepGroupId,omitempty" xml:"DepGroupId,omitempty"`
	DepGroupName *string `json:"DepGroupName,omitempty" xml:"DepGroupName,omitempty"`
}

func (s ListAllHotLineSkillGroupsResponseBodyDataDepGroup) String() string {
	return tea.Prettify(s)
}

func (s ListAllHotLineSkillGroupsResponseBodyDataDepGroup) GoString() string {
	return s.String()
}

func (s *ListAllHotLineSkillGroupsResponseBodyDataDepGroup) SetDepGroupId(v int64) *ListAllHotLineSkillGroupsResponseBodyDataDepGroup {
	s.DepGroupId = &v
	return s
}

func (s *ListAllHotLineSkillGroupsResponseBodyDataDepGroup) SetDepGroupName(v string) *ListAllHotLineSkillGroupsResponseBodyDataDepGroup {
	s.DepGroupName = &v
	return s
}

type ListAllHotLineSkillGroupsResponseBodyDataSkillGroupAgents struct {
	Agents     []*ListAllHotLineSkillGroupsResponseBodyDataSkillGroupAgentsAgents   `json:"Agents,omitempty" xml:"Agents,omitempty" type:"Repeated"`
	SkillGroup *ListAllHotLineSkillGroupsResponseBodyDataSkillGroupAgentsSkillGroup `json:"SkillGroup,omitempty" xml:"SkillGroup,omitempty" type:"Struct"`
}

func (s ListAllHotLineSkillGroupsResponseBodyDataSkillGroupAgents) String() string {
	return tea.Prettify(s)
}

func (s ListAllHotLineSkillGroupsResponseBodyDataSkillGroupAgents) GoString() string {
	return s.String()
}

func (s *ListAllHotLineSkillGroupsResponseBodyDataSkillGroupAgents) SetAgents(v []*ListAllHotLineSkillGroupsResponseBodyDataSkillGroupAgentsAgents) *ListAllHotLineSkillGroupsResponseBodyDataSkillGroupAgents {
	s.Agents = v
	return s
}

func (s *ListAllHotLineSkillGroupsResponseBodyDataSkillGroupAgents) SetSkillGroup(v *ListAllHotLineSkillGroupsResponseBodyDataSkillGroupAgentsSkillGroup) *ListAllHotLineSkillGroupsResponseBodyDataSkillGroupAgents {
	s.SkillGroup = v
	return s
}

type ListAllHotLineSkillGroupsResponseBodyDataSkillGroupAgentsAgents struct {
	AccountName   *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	Acid          *string `json:"Acid,omitempty" xml:"Acid,omitempty"`
	AgentId       *int64  `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	ConnId        *string `json:"ConnId,omitempty" xml:"ConnId,omitempty"`
	CustomerPhone *string `json:"CustomerPhone,omitempty" xml:"CustomerPhone,omitempty"`
	DisplayName   *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	JobId         *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	Status        *int64  `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListAllHotLineSkillGroupsResponseBodyDataSkillGroupAgentsAgents) String() string {
	return tea.Prettify(s)
}

func (s ListAllHotLineSkillGroupsResponseBodyDataSkillGroupAgentsAgents) GoString() string {
	return s.String()
}

func (s *ListAllHotLineSkillGroupsResponseBodyDataSkillGroupAgentsAgents) SetAccountName(v string) *ListAllHotLineSkillGroupsResponseBodyDataSkillGroupAgentsAgents {
	s.AccountName = &v
	return s
}

func (s *ListAllHotLineSkillGroupsResponseBodyDataSkillGroupAgentsAgents) SetAcid(v string) *ListAllHotLineSkillGroupsResponseBodyDataSkillGroupAgentsAgents {
	s.Acid = &v
	return s
}

func (s *ListAllHotLineSkillGroupsResponseBodyDataSkillGroupAgentsAgents) SetAgentId(v int64) *ListAllHotLineSkillGroupsResponseBodyDataSkillGroupAgentsAgents {
	s.AgentId = &v
	return s
}

func (s *ListAllHotLineSkillGroupsResponseBodyDataSkillGroupAgentsAgents) SetConnId(v string) *ListAllHotLineSkillGroupsResponseBodyDataSkillGroupAgentsAgents {
	s.ConnId = &v
	return s
}

func (s *ListAllHotLineSkillGroupsResponseBodyDataSkillGroupAgentsAgents) SetCustomerPhone(v string) *ListAllHotLineSkillGroupsResponseBodyDataSkillGroupAgentsAgents {
	s.CustomerPhone = &v
	return s
}

func (s *ListAllHotLineSkillGroupsResponseBodyDataSkillGroupAgentsAgents) SetDisplayName(v string) *ListAllHotLineSkillGroupsResponseBodyDataSkillGroupAgentsAgents {
	s.DisplayName = &v
	return s
}

func (s *ListAllHotLineSkillGroupsResponseBodyDataSkillGroupAgentsAgents) SetJobId(v string) *ListAllHotLineSkillGroupsResponseBodyDataSkillGroupAgentsAgents {
	s.JobId = &v
	return s
}

func (s *ListAllHotLineSkillGroupsResponseBodyDataSkillGroupAgentsAgents) SetStatus(v int64) *ListAllHotLineSkillGroupsResponseBodyDataSkillGroupAgentsAgents {
	s.Status = &v
	return s
}

type ListAllHotLineSkillGroupsResponseBodyDataSkillGroupAgentsSkillGroup struct {
	DisplayName  *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	SkillGroupId *int64  `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
}

func (s ListAllHotLineSkillGroupsResponseBodyDataSkillGroupAgentsSkillGroup) String() string {
	return tea.Prettify(s)
}

func (s ListAllHotLineSkillGroupsResponseBodyDataSkillGroupAgentsSkillGroup) GoString() string {
	return s.String()
}

func (s *ListAllHotLineSkillGroupsResponseBodyDataSkillGroupAgentsSkillGroup) SetDisplayName(v string) *ListAllHotLineSkillGroupsResponseBodyDataSkillGroupAgentsSkillGroup {
	s.DisplayName = &v
	return s
}

func (s *ListAllHotLineSkillGroupsResponseBodyDataSkillGroupAgentsSkillGroup) SetSkillGroupId(v int64) *ListAllHotLineSkillGroupsResponseBodyDataSkillGroupAgentsSkillGroup {
	s.SkillGroupId = &v
	return s
}

type ListAllHotLineSkillGroupsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListAllHotLineSkillGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAllHotLineSkillGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAllHotLineSkillGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListAllHotLineSkillGroupsResponse) SetHeaders(v map[string]*string) *ListAllHotLineSkillGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListAllHotLineSkillGroupsResponse) SetStatusCode(v int32) *ListAllHotLineSkillGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAllHotLineSkillGroupsResponse) SetBody(v *ListAllHotLineSkillGroupsResponseBody) *ListAllHotLineSkillGroupsResponse {
	s.Body = v
	return s
}

type ListHotlineRecordRequest struct {
	CallId      *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ListHotlineRecordRequest) String() string {
	return tea.Prettify(s)
}

func (s ListHotlineRecordRequest) GoString() string {
	return s.String()
}

func (s *ListHotlineRecordRequest) SetCallId(v string) *ListHotlineRecordRequest {
	s.CallId = &v
	return s
}

func (s *ListHotlineRecordRequest) SetClientToken(v string) *ListHotlineRecordRequest {
	s.ClientToken = &v
	return s
}

func (s *ListHotlineRecordRequest) SetInstanceId(v string) *ListHotlineRecordRequest {
	s.InstanceId = &v
	return s
}

type ListHotlineRecordResponseBody struct {
	Code           *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           []*ListHotlineRecordResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	HttpStatusCode *int64                               `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListHotlineRecordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListHotlineRecordResponseBody) GoString() string {
	return s.String()
}

func (s *ListHotlineRecordResponseBody) SetCode(v string) *ListHotlineRecordResponseBody {
	s.Code = &v
	return s
}

func (s *ListHotlineRecordResponseBody) SetData(v []*ListHotlineRecordResponseBodyData) *ListHotlineRecordResponseBody {
	s.Data = v
	return s
}

func (s *ListHotlineRecordResponseBody) SetHttpStatusCode(v int64) *ListHotlineRecordResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListHotlineRecordResponseBody) SetMessage(v string) *ListHotlineRecordResponseBody {
	s.Message = &v
	return s
}

func (s *ListHotlineRecordResponseBody) SetRequestId(v string) *ListHotlineRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHotlineRecordResponseBody) SetSuccess(v bool) *ListHotlineRecordResponseBody {
	s.Success = &v
	return s
}

type ListHotlineRecordResponseBodyData struct {
	CallId       *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	ConnectionId *string `json:"ConnectionId,omitempty" xml:"ConnectionId,omitempty"`
	EndTime      *bool   `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime    *bool   `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Url          *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s ListHotlineRecordResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListHotlineRecordResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListHotlineRecordResponseBodyData) SetCallId(v string) *ListHotlineRecordResponseBodyData {
	s.CallId = &v
	return s
}

func (s *ListHotlineRecordResponseBodyData) SetConnectionId(v string) *ListHotlineRecordResponseBodyData {
	s.ConnectionId = &v
	return s
}

func (s *ListHotlineRecordResponseBodyData) SetEndTime(v bool) *ListHotlineRecordResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *ListHotlineRecordResponseBodyData) SetStartTime(v bool) *ListHotlineRecordResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *ListHotlineRecordResponseBodyData) SetUrl(v string) *ListHotlineRecordResponseBodyData {
	s.Url = &v
	return s
}

type ListHotlineRecordResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListHotlineRecordResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListHotlineRecordResponse) String() string {
	return tea.Prettify(s)
}

func (s ListHotlineRecordResponse) GoString() string {
	return s.String()
}

func (s *ListHotlineRecordResponse) SetHeaders(v map[string]*string) *ListHotlineRecordResponse {
	s.Headers = v
	return s
}

func (s *ListHotlineRecordResponse) SetStatusCode(v int32) *ListHotlineRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHotlineRecordResponse) SetBody(v *ListHotlineRecordResponseBody) *ListHotlineRecordResponse {
	s.Body = v
	return s
}

type ListOutboundPhoneNumberRequest struct {
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ListOutboundPhoneNumberRequest) String() string {
	return tea.Prettify(s)
}

func (s ListOutboundPhoneNumberRequest) GoString() string {
	return s.String()
}

func (s *ListOutboundPhoneNumberRequest) SetAccountName(v string) *ListOutboundPhoneNumberRequest {
	s.AccountName = &v
	return s
}

func (s *ListOutboundPhoneNumberRequest) SetClientToken(v string) *ListOutboundPhoneNumberRequest {
	s.ClientToken = &v
	return s
}

func (s *ListOutboundPhoneNumberRequest) SetInstanceId(v string) *ListOutboundPhoneNumberRequest {
	s.InstanceId = &v
	return s
}

type ListOutboundPhoneNumberResponseBody struct {
	Code           *string   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           []*string `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	HttpStatusCode *int64    `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool     `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListOutboundPhoneNumberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListOutboundPhoneNumberResponseBody) GoString() string {
	return s.String()
}

func (s *ListOutboundPhoneNumberResponseBody) SetCode(v string) *ListOutboundPhoneNumberResponseBody {
	s.Code = &v
	return s
}

func (s *ListOutboundPhoneNumberResponseBody) SetData(v []*string) *ListOutboundPhoneNumberResponseBody {
	s.Data = v
	return s
}

func (s *ListOutboundPhoneNumberResponseBody) SetHttpStatusCode(v int64) *ListOutboundPhoneNumberResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListOutboundPhoneNumberResponseBody) SetMessage(v string) *ListOutboundPhoneNumberResponseBody {
	s.Message = &v
	return s
}

func (s *ListOutboundPhoneNumberResponseBody) SetRequestId(v string) *ListOutboundPhoneNumberResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOutboundPhoneNumberResponseBody) SetSuccess(v bool) *ListOutboundPhoneNumberResponseBody {
	s.Success = &v
	return s
}

type ListOutboundPhoneNumberResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListOutboundPhoneNumberResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListOutboundPhoneNumberResponse) String() string {
	return tea.Prettify(s)
}

func (s ListOutboundPhoneNumberResponse) GoString() string {
	return s.String()
}

func (s *ListOutboundPhoneNumberResponse) SetHeaders(v map[string]*string) *ListOutboundPhoneNumberResponse {
	s.Headers = v
	return s
}

func (s *ListOutboundPhoneNumberResponse) SetStatusCode(v int32) *ListOutboundPhoneNumberResponse {
	s.StatusCode = &v
	return s
}

func (s *ListOutboundPhoneNumberResponse) SetBody(v *ListOutboundPhoneNumberResponseBody) *ListOutboundPhoneNumberResponse {
	s.Body = v
	return s
}

type ListSkillGroupRequest struct {
	ChannelType *int32  `json:"ChannelType,omitempty" xml:"ChannelType,omitempty"`
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ListSkillGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSkillGroupRequest) GoString() string {
	return s.String()
}

func (s *ListSkillGroupRequest) SetChannelType(v int32) *ListSkillGroupRequest {
	s.ChannelType = &v
	return s
}

func (s *ListSkillGroupRequest) SetClientToken(v string) *ListSkillGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *ListSkillGroupRequest) SetInstanceId(v string) *ListSkillGroupRequest {
	s.InstanceId = &v
	return s
}

type ListSkillGroupResponseBody struct {
	Code      *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*ListSkillGroupResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListSkillGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSkillGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ListSkillGroupResponseBody) SetCode(v string) *ListSkillGroupResponseBody {
	s.Code = &v
	return s
}

func (s *ListSkillGroupResponseBody) SetData(v []*ListSkillGroupResponseBodyData) *ListSkillGroupResponseBody {
	s.Data = v
	return s
}

func (s *ListSkillGroupResponseBody) SetMessage(v string) *ListSkillGroupResponseBody {
	s.Message = &v
	return s
}

func (s *ListSkillGroupResponseBody) SetRequestId(v string) *ListSkillGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSkillGroupResponseBody) SetSuccess(v bool) *ListSkillGroupResponseBody {
	s.Success = &v
	return s
}

type ListSkillGroupResponseBodyData struct {
	ChannelType  *int32  `json:"ChannelType,omitempty" xml:"ChannelType,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DisplayName  *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	SkillGroupId *int64  `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
}

func (s ListSkillGroupResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListSkillGroupResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListSkillGroupResponseBodyData) SetChannelType(v int32) *ListSkillGroupResponseBodyData {
	s.ChannelType = &v
	return s
}

func (s *ListSkillGroupResponseBodyData) SetDescription(v string) *ListSkillGroupResponseBodyData {
	s.Description = &v
	return s
}

func (s *ListSkillGroupResponseBodyData) SetDisplayName(v string) *ListSkillGroupResponseBodyData {
	s.DisplayName = &v
	return s
}

func (s *ListSkillGroupResponseBodyData) SetName(v string) *ListSkillGroupResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListSkillGroupResponseBodyData) SetSkillGroupId(v int64) *ListSkillGroupResponseBodyData {
	s.SkillGroupId = &v
	return s
}

type ListSkillGroupResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListSkillGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSkillGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSkillGroupResponse) GoString() string {
	return s.String()
}

func (s *ListSkillGroupResponse) SetHeaders(v map[string]*string) *ListSkillGroupResponse {
	s.Headers = v
	return s
}

func (s *ListSkillGroupResponse) SetStatusCode(v int32) *ListSkillGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSkillGroupResponse) SetBody(v *ListSkillGroupResponseBody) *ListSkillGroupResponse {
	s.Body = v
	return s
}

type QueryHotlineDashboardRequest struct {
	CurrentPageNum   *int32    `json:"CurrentPageNum,omitempty" xml:"CurrentPageNum,omitempty"`
	DepartmentIdList []*int64  `json:"DepartmentIdList,omitempty" xml:"DepartmentIdList,omitempty" type:"Repeated"`
	EndDate          *int64    `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	InstanceId       *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PageSize         *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ServicerIdList   []*int64  `json:"ServicerIdList,omitempty" xml:"ServicerIdList,omitempty" type:"Repeated"`
	SortFieldList    []*string `json:"SortFieldList,omitempty" xml:"SortFieldList,omitempty" type:"Repeated"`
	StartDate        *int64    `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s QueryHotlineDashboardRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryHotlineDashboardRequest) GoString() string {
	return s.String()
}

func (s *QueryHotlineDashboardRequest) SetCurrentPageNum(v int32) *QueryHotlineDashboardRequest {
	s.CurrentPageNum = &v
	return s
}

func (s *QueryHotlineDashboardRequest) SetDepartmentIdList(v []*int64) *QueryHotlineDashboardRequest {
	s.DepartmentIdList = v
	return s
}

func (s *QueryHotlineDashboardRequest) SetEndDate(v int64) *QueryHotlineDashboardRequest {
	s.EndDate = &v
	return s
}

func (s *QueryHotlineDashboardRequest) SetInstanceId(v string) *QueryHotlineDashboardRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryHotlineDashboardRequest) SetPageSize(v int32) *QueryHotlineDashboardRequest {
	s.PageSize = &v
	return s
}

func (s *QueryHotlineDashboardRequest) SetServicerIdList(v []*int64) *QueryHotlineDashboardRequest {
	s.ServicerIdList = v
	return s
}

func (s *QueryHotlineDashboardRequest) SetSortFieldList(v []*string) *QueryHotlineDashboardRequest {
	s.SortFieldList = v
	return s
}

func (s *QueryHotlineDashboardRequest) SetStartDate(v int64) *QueryHotlineDashboardRequest {
	s.StartDate = &v
	return s
}

type QueryHotlineDashboardShrinkRequest struct {
	CurrentPageNum         *int32  `json:"CurrentPageNum,omitempty" xml:"CurrentPageNum,omitempty"`
	DepartmentIdListShrink *string `json:"DepartmentIdList,omitempty" xml:"DepartmentIdList,omitempty"`
	EndDate                *int64  `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	InstanceId             *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PageSize               *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ServicerIdListShrink   *string `json:"ServicerIdList,omitempty" xml:"ServicerIdList,omitempty"`
	SortFieldListShrink    *string `json:"SortFieldList,omitempty" xml:"SortFieldList,omitempty"`
	StartDate              *int64  `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s QueryHotlineDashboardShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryHotlineDashboardShrinkRequest) GoString() string {
	return s.String()
}

func (s *QueryHotlineDashboardShrinkRequest) SetCurrentPageNum(v int32) *QueryHotlineDashboardShrinkRequest {
	s.CurrentPageNum = &v
	return s
}

func (s *QueryHotlineDashboardShrinkRequest) SetDepartmentIdListShrink(v string) *QueryHotlineDashboardShrinkRequest {
	s.DepartmentIdListShrink = &v
	return s
}

func (s *QueryHotlineDashboardShrinkRequest) SetEndDate(v int64) *QueryHotlineDashboardShrinkRequest {
	s.EndDate = &v
	return s
}

func (s *QueryHotlineDashboardShrinkRequest) SetInstanceId(v string) *QueryHotlineDashboardShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryHotlineDashboardShrinkRequest) SetPageSize(v int32) *QueryHotlineDashboardShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *QueryHotlineDashboardShrinkRequest) SetServicerIdListShrink(v string) *QueryHotlineDashboardShrinkRequest {
	s.ServicerIdListShrink = &v
	return s
}

func (s *QueryHotlineDashboardShrinkRequest) SetSortFieldListShrink(v string) *QueryHotlineDashboardShrinkRequest {
	s.SortFieldListShrink = &v
	return s
}

func (s *QueryHotlineDashboardShrinkRequest) SetStartDate(v int64) *QueryHotlineDashboardShrinkRequest {
	s.StartDate = &v
	return s
}

type QueryHotlineDashboardResponseBody struct {
	Code           *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *QueryHotlineDashboardResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                                 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryHotlineDashboardResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryHotlineDashboardResponseBody) GoString() string {
	return s.String()
}

func (s *QueryHotlineDashboardResponseBody) SetCode(v string) *QueryHotlineDashboardResponseBody {
	s.Code = &v
	return s
}

func (s *QueryHotlineDashboardResponseBody) SetData(v *QueryHotlineDashboardResponseBodyData) *QueryHotlineDashboardResponseBody {
	s.Data = v
	return s
}

func (s *QueryHotlineDashboardResponseBody) SetHttpStatusCode(v int32) *QueryHotlineDashboardResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *QueryHotlineDashboardResponseBody) SetMessage(v string) *QueryHotlineDashboardResponseBody {
	s.Message = &v
	return s
}

func (s *QueryHotlineDashboardResponseBody) SetRequestId(v string) *QueryHotlineDashboardResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryHotlineDashboardResponseBody) SetSuccess(v bool) *QueryHotlineDashboardResponseBody {
	s.Success = &v
	return s
}

type QueryHotlineDashboardResponseBodyData struct {
	Results []*QueryHotlineDashboardResponseBodyDataResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
	Total   *int32                                          `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s QueryHotlineDashboardResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryHotlineDashboardResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryHotlineDashboardResponseBodyData) SetResults(v []*QueryHotlineDashboardResponseBodyDataResults) *QueryHotlineDashboardResponseBodyData {
	s.Results = v
	return s
}

func (s *QueryHotlineDashboardResponseBodyData) SetTotal(v int32) *QueryHotlineDashboardResponseBodyData {
	s.Total = &v
	return s
}

type QueryHotlineDashboardResponseBodyDataResults struct {
	DateId                   *string `json:"DateId,omitempty" xml:"DateId,omitempty"`
	DepartmentId             *int64  `json:"DepartmentId,omitempty" xml:"DepartmentId,omitempty"`
	DepartmentName           *string `json:"DepartmentName,omitempty" xml:"DepartmentName,omitempty"`
	EffectiveInCalls         *int32  `json:"EffectiveInCalls,omitempty" xml:"EffectiveInCalls,omitempty"`
	EffectiveInServiceNotes  *int32  `json:"EffectiveInServiceNotes,omitempty" xml:"EffectiveInServiceNotes,omitempty"`
	EffectiveOutCalls        *int32  `json:"EffectiveOutCalls,omitempty" xml:"EffectiveOutCalls,omitempty"`
	EffectiveOutServiceNotes *int32  `json:"EffectiveOutServiceNotes,omitempty" xml:"EffectiveOutServiceNotes,omitempty"`
	ServicerId               *int64  `json:"ServicerId,omitempty" xml:"ServicerId,omitempty"`
	ServicerRealName         *string `json:"ServicerRealName,omitempty" xml:"ServicerRealName,omitempty"`
	ServicerShowName         *string `json:"ServicerShowName,omitempty" xml:"ServicerShowName,omitempty"`
}

func (s QueryHotlineDashboardResponseBodyDataResults) String() string {
	return tea.Prettify(s)
}

func (s QueryHotlineDashboardResponseBodyDataResults) GoString() string {
	return s.String()
}

func (s *QueryHotlineDashboardResponseBodyDataResults) SetDateId(v string) *QueryHotlineDashboardResponseBodyDataResults {
	s.DateId = &v
	return s
}

func (s *QueryHotlineDashboardResponseBodyDataResults) SetDepartmentId(v int64) *QueryHotlineDashboardResponseBodyDataResults {
	s.DepartmentId = &v
	return s
}

func (s *QueryHotlineDashboardResponseBodyDataResults) SetDepartmentName(v string) *QueryHotlineDashboardResponseBodyDataResults {
	s.DepartmentName = &v
	return s
}

func (s *QueryHotlineDashboardResponseBodyDataResults) SetEffectiveInCalls(v int32) *QueryHotlineDashboardResponseBodyDataResults {
	s.EffectiveInCalls = &v
	return s
}

func (s *QueryHotlineDashboardResponseBodyDataResults) SetEffectiveInServiceNotes(v int32) *QueryHotlineDashboardResponseBodyDataResults {
	s.EffectiveInServiceNotes = &v
	return s
}

func (s *QueryHotlineDashboardResponseBodyDataResults) SetEffectiveOutCalls(v int32) *QueryHotlineDashboardResponseBodyDataResults {
	s.EffectiveOutCalls = &v
	return s
}

func (s *QueryHotlineDashboardResponseBodyDataResults) SetEffectiveOutServiceNotes(v int32) *QueryHotlineDashboardResponseBodyDataResults {
	s.EffectiveOutServiceNotes = &v
	return s
}

func (s *QueryHotlineDashboardResponseBodyDataResults) SetServicerId(v int64) *QueryHotlineDashboardResponseBodyDataResults {
	s.ServicerId = &v
	return s
}

func (s *QueryHotlineDashboardResponseBodyDataResults) SetServicerRealName(v string) *QueryHotlineDashboardResponseBodyDataResults {
	s.ServicerRealName = &v
	return s
}

func (s *QueryHotlineDashboardResponseBodyDataResults) SetServicerShowName(v string) *QueryHotlineDashboardResponseBodyDataResults {
	s.ServicerShowName = &v
	return s
}

type QueryHotlineDashboardResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryHotlineDashboardResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryHotlineDashboardResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryHotlineDashboardResponse) GoString() string {
	return s.String()
}

func (s *QueryHotlineDashboardResponse) SetHeaders(v map[string]*string) *QueryHotlineDashboardResponse {
	s.Headers = v
	return s
}

func (s *QueryHotlineDashboardResponse) SetStatusCode(v int32) *QueryHotlineDashboardResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryHotlineDashboardResponse) SetBody(v *QueryHotlineDashboardResponseBody) *QueryHotlineDashboardResponse {
	s.Body = v
	return s
}

type QueryHotlineSessionRequest struct {
	Acid           *string `json:"Acid,omitempty" xml:"Acid,omitempty"`
	CallResult     *string `json:"CallResult,omitempty" xml:"CallResult,omitempty"`
	CallType       *int32  `json:"CallType,omitempty" xml:"CallType,omitempty"`
	CalledNumber   *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	CallingNumber  *string `json:"CallingNumber,omitempty" xml:"CallingNumber,omitempty"`
	GroupId        *int64  `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName      *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MemberId       *string `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	MemberName     *string `json:"MemberName,omitempty" xml:"MemberName,omitempty"`
	PageNo         *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize       *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Params         *string `json:"Params,omitempty" xml:"Params,omitempty"`
	QueryEndTime   *int64  `json:"QueryEndTime,omitempty" xml:"QueryEndTime,omitempty"`
	QueryStartTime *int64  `json:"QueryStartTime,omitempty" xml:"QueryStartTime,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ServicerId     *string `json:"ServicerId,omitempty" xml:"ServicerId,omitempty"`
	ServicerName   *string `json:"ServicerName,omitempty" xml:"ServicerName,omitempty"`
}

func (s QueryHotlineSessionRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryHotlineSessionRequest) GoString() string {
	return s.String()
}

func (s *QueryHotlineSessionRequest) SetAcid(v string) *QueryHotlineSessionRequest {
	s.Acid = &v
	return s
}

func (s *QueryHotlineSessionRequest) SetCallResult(v string) *QueryHotlineSessionRequest {
	s.CallResult = &v
	return s
}

func (s *QueryHotlineSessionRequest) SetCallType(v int32) *QueryHotlineSessionRequest {
	s.CallType = &v
	return s
}

func (s *QueryHotlineSessionRequest) SetCalledNumber(v string) *QueryHotlineSessionRequest {
	s.CalledNumber = &v
	return s
}

func (s *QueryHotlineSessionRequest) SetCallingNumber(v string) *QueryHotlineSessionRequest {
	s.CallingNumber = &v
	return s
}

func (s *QueryHotlineSessionRequest) SetGroupId(v int64) *QueryHotlineSessionRequest {
	s.GroupId = &v
	return s
}

func (s *QueryHotlineSessionRequest) SetGroupName(v string) *QueryHotlineSessionRequest {
	s.GroupName = &v
	return s
}

func (s *QueryHotlineSessionRequest) SetInstanceId(v string) *QueryHotlineSessionRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryHotlineSessionRequest) SetMemberId(v string) *QueryHotlineSessionRequest {
	s.MemberId = &v
	return s
}

func (s *QueryHotlineSessionRequest) SetMemberName(v string) *QueryHotlineSessionRequest {
	s.MemberName = &v
	return s
}

func (s *QueryHotlineSessionRequest) SetPageNo(v int32) *QueryHotlineSessionRequest {
	s.PageNo = &v
	return s
}

func (s *QueryHotlineSessionRequest) SetPageSize(v int32) *QueryHotlineSessionRequest {
	s.PageSize = &v
	return s
}

func (s *QueryHotlineSessionRequest) SetParams(v string) *QueryHotlineSessionRequest {
	s.Params = &v
	return s
}

func (s *QueryHotlineSessionRequest) SetQueryEndTime(v int64) *QueryHotlineSessionRequest {
	s.QueryEndTime = &v
	return s
}

func (s *QueryHotlineSessionRequest) SetQueryStartTime(v int64) *QueryHotlineSessionRequest {
	s.QueryStartTime = &v
	return s
}

func (s *QueryHotlineSessionRequest) SetRequestId(v string) *QueryHotlineSessionRequest {
	s.RequestId = &v
	return s
}

func (s *QueryHotlineSessionRequest) SetServicerId(v string) *QueryHotlineSessionRequest {
	s.ServicerId = &v
	return s
}

func (s *QueryHotlineSessionRequest) SetServicerName(v string) *QueryHotlineSessionRequest {
	s.ServicerName = &v
	return s
}

type QueryHotlineSessionResponseBody struct {
	Code      *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *QueryHotlineSessionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryHotlineSessionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryHotlineSessionResponseBody) GoString() string {
	return s.String()
}

func (s *QueryHotlineSessionResponseBody) SetCode(v string) *QueryHotlineSessionResponseBody {
	s.Code = &v
	return s
}

func (s *QueryHotlineSessionResponseBody) SetData(v *QueryHotlineSessionResponseBodyData) *QueryHotlineSessionResponseBody {
	s.Data = v
	return s
}

func (s *QueryHotlineSessionResponseBody) SetMessage(v string) *QueryHotlineSessionResponseBody {
	s.Message = &v
	return s
}

func (s *QueryHotlineSessionResponseBody) SetRequestId(v string) *QueryHotlineSessionResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryHotlineSessionResponseBody) SetSuccess(v bool) *QueryHotlineSessionResponseBody {
	s.Success = &v
	return s
}

type QueryHotlineSessionResponseBodyData struct {
	CallDetailRecord []*QueryHotlineSessionResponseBodyDataCallDetailRecord `json:"CallDetailRecord,omitempty" xml:"CallDetailRecord,omitempty" type:"Repeated"`
	PageNumber       *int32                                                 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize         *int32                                                 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount       *int32                                                 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s QueryHotlineSessionResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryHotlineSessionResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryHotlineSessionResponseBodyData) SetCallDetailRecord(v []*QueryHotlineSessionResponseBodyDataCallDetailRecord) *QueryHotlineSessionResponseBodyData {
	s.CallDetailRecord = v
	return s
}

func (s *QueryHotlineSessionResponseBodyData) SetPageNumber(v int32) *QueryHotlineSessionResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *QueryHotlineSessionResponseBodyData) SetPageSize(v int32) *QueryHotlineSessionResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryHotlineSessionResponseBodyData) SetTotalCount(v int32) *QueryHotlineSessionResponseBodyData {
	s.TotalCount = &v
	return s
}

type QueryHotlineSessionResponseBodyDataCallDetailRecord struct {
	Acid                *string `json:"Acid,omitempty" xml:"Acid,omitempty"`
	CallContinueTime    *int32  `json:"CallContinueTime,omitempty" xml:"CallContinueTime,omitempty"`
	CallResult          *string `json:"CallResult,omitempty" xml:"CallResult,omitempty"`
	CallType            *int32  `json:"CallType,omitempty" xml:"CallType,omitempty"`
	CalledNumber        *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	CallingNumber       *string `json:"CallingNumber,omitempty" xml:"CallingNumber,omitempty"`
	CreateTime          *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	EvaluationLevel     *int32  `json:"EvaluationLevel,omitempty" xml:"EvaluationLevel,omitempty"`
	EvaluationScore     *int32  `json:"EvaluationScore,omitempty" xml:"EvaluationScore,omitempty"`
	GroupId             *int64  `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName           *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	HangUpRole          *string `json:"HangUpRole,omitempty" xml:"HangUpRole,omitempty"`
	HangUpTime          *string `json:"HangUpTime,omitempty" xml:"HangUpTime,omitempty"`
	InQueueTime         *string `json:"InQueueTime,omitempty" xml:"InQueueTime,omitempty"`
	MemberId            *string `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	MemberName          *string `json:"MemberName,omitempty" xml:"MemberName,omitempty"`
	OutQueueTime        *string `json:"OutQueueTime,omitempty" xml:"OutQueueTime,omitempty"`
	PickUpTime          *string `json:"PickUpTime,omitempty" xml:"PickUpTime,omitempty"`
	QueueUpContinueTime *int32  `json:"QueueUpContinueTime,omitempty" xml:"QueueUpContinueTime,omitempty"`
	RingContinueTime    *int32  `json:"RingContinueTime,omitempty" xml:"RingContinueTime,omitempty"`
	RingEndTime         *string `json:"RingEndTime,omitempty" xml:"RingEndTime,omitempty"`
	RingStartTime       *string `json:"RingStartTime,omitempty" xml:"RingStartTime,omitempty"`
	ServicerId          *string `json:"ServicerId,omitempty" xml:"ServicerId,omitempty"`
	ServicerName        *string `json:"ServicerName,omitempty" xml:"ServicerName,omitempty"`
}

func (s QueryHotlineSessionResponseBodyDataCallDetailRecord) String() string {
	return tea.Prettify(s)
}

func (s QueryHotlineSessionResponseBodyDataCallDetailRecord) GoString() string {
	return s.String()
}

func (s *QueryHotlineSessionResponseBodyDataCallDetailRecord) SetAcid(v string) *QueryHotlineSessionResponseBodyDataCallDetailRecord {
	s.Acid = &v
	return s
}

func (s *QueryHotlineSessionResponseBodyDataCallDetailRecord) SetCallContinueTime(v int32) *QueryHotlineSessionResponseBodyDataCallDetailRecord {
	s.CallContinueTime = &v
	return s
}

func (s *QueryHotlineSessionResponseBodyDataCallDetailRecord) SetCallResult(v string) *QueryHotlineSessionResponseBodyDataCallDetailRecord {
	s.CallResult = &v
	return s
}

func (s *QueryHotlineSessionResponseBodyDataCallDetailRecord) SetCallType(v int32) *QueryHotlineSessionResponseBodyDataCallDetailRecord {
	s.CallType = &v
	return s
}

func (s *QueryHotlineSessionResponseBodyDataCallDetailRecord) SetCalledNumber(v string) *QueryHotlineSessionResponseBodyDataCallDetailRecord {
	s.CalledNumber = &v
	return s
}

func (s *QueryHotlineSessionResponseBodyDataCallDetailRecord) SetCallingNumber(v string) *QueryHotlineSessionResponseBodyDataCallDetailRecord {
	s.CallingNumber = &v
	return s
}

func (s *QueryHotlineSessionResponseBodyDataCallDetailRecord) SetCreateTime(v string) *QueryHotlineSessionResponseBodyDataCallDetailRecord {
	s.CreateTime = &v
	return s
}

func (s *QueryHotlineSessionResponseBodyDataCallDetailRecord) SetEvaluationLevel(v int32) *QueryHotlineSessionResponseBodyDataCallDetailRecord {
	s.EvaluationLevel = &v
	return s
}

func (s *QueryHotlineSessionResponseBodyDataCallDetailRecord) SetEvaluationScore(v int32) *QueryHotlineSessionResponseBodyDataCallDetailRecord {
	s.EvaluationScore = &v
	return s
}

func (s *QueryHotlineSessionResponseBodyDataCallDetailRecord) SetGroupId(v int64) *QueryHotlineSessionResponseBodyDataCallDetailRecord {
	s.GroupId = &v
	return s
}

func (s *QueryHotlineSessionResponseBodyDataCallDetailRecord) SetGroupName(v string) *QueryHotlineSessionResponseBodyDataCallDetailRecord {
	s.GroupName = &v
	return s
}

func (s *QueryHotlineSessionResponseBodyDataCallDetailRecord) SetHangUpRole(v string) *QueryHotlineSessionResponseBodyDataCallDetailRecord {
	s.HangUpRole = &v
	return s
}

func (s *QueryHotlineSessionResponseBodyDataCallDetailRecord) SetHangUpTime(v string) *QueryHotlineSessionResponseBodyDataCallDetailRecord {
	s.HangUpTime = &v
	return s
}

func (s *QueryHotlineSessionResponseBodyDataCallDetailRecord) SetInQueueTime(v string) *QueryHotlineSessionResponseBodyDataCallDetailRecord {
	s.InQueueTime = &v
	return s
}

func (s *QueryHotlineSessionResponseBodyDataCallDetailRecord) SetMemberId(v string) *QueryHotlineSessionResponseBodyDataCallDetailRecord {
	s.MemberId = &v
	return s
}

func (s *QueryHotlineSessionResponseBodyDataCallDetailRecord) SetMemberName(v string) *QueryHotlineSessionResponseBodyDataCallDetailRecord {
	s.MemberName = &v
	return s
}

func (s *QueryHotlineSessionResponseBodyDataCallDetailRecord) SetOutQueueTime(v string) *QueryHotlineSessionResponseBodyDataCallDetailRecord {
	s.OutQueueTime = &v
	return s
}

func (s *QueryHotlineSessionResponseBodyDataCallDetailRecord) SetPickUpTime(v string) *QueryHotlineSessionResponseBodyDataCallDetailRecord {
	s.PickUpTime = &v
	return s
}

func (s *QueryHotlineSessionResponseBodyDataCallDetailRecord) SetQueueUpContinueTime(v int32) *QueryHotlineSessionResponseBodyDataCallDetailRecord {
	s.QueueUpContinueTime = &v
	return s
}

func (s *QueryHotlineSessionResponseBodyDataCallDetailRecord) SetRingContinueTime(v int32) *QueryHotlineSessionResponseBodyDataCallDetailRecord {
	s.RingContinueTime = &v
	return s
}

func (s *QueryHotlineSessionResponseBodyDataCallDetailRecord) SetRingEndTime(v string) *QueryHotlineSessionResponseBodyDataCallDetailRecord {
	s.RingEndTime = &v
	return s
}

func (s *QueryHotlineSessionResponseBodyDataCallDetailRecord) SetRingStartTime(v string) *QueryHotlineSessionResponseBodyDataCallDetailRecord {
	s.RingStartTime = &v
	return s
}

func (s *QueryHotlineSessionResponseBodyDataCallDetailRecord) SetServicerId(v string) *QueryHotlineSessionResponseBodyDataCallDetailRecord {
	s.ServicerId = &v
	return s
}

func (s *QueryHotlineSessionResponseBodyDataCallDetailRecord) SetServicerName(v string) *QueryHotlineSessionResponseBodyDataCallDetailRecord {
	s.ServicerName = &v
	return s
}

type QueryHotlineSessionResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryHotlineSessionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryHotlineSessionResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryHotlineSessionResponse) GoString() string {
	return s.String()
}

func (s *QueryHotlineSessionResponse) SetHeaders(v map[string]*string) *QueryHotlineSessionResponse {
	s.Headers = v
	return s
}

func (s *QueryHotlineSessionResponse) SetStatusCode(v int32) *QueryHotlineSessionResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryHotlineSessionResponse) SetBody(v *QueryHotlineSessionResponseBody) *QueryHotlineSessionResponse {
	s.Body = v
	return s
}

type QueryRelationTicketsRequest struct {
	CaseId             *int64                 `json:"CaseId,omitempty" xml:"CaseId,omitempty"`
	CaseType           *int32                 `json:"CaseType,omitempty" xml:"CaseType,omitempty"`
	ChannelId          *string                `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	ChannelType        *int32                 `json:"ChannelType,omitempty" xml:"ChannelType,omitempty"`
	CurrentPage        *int32                 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	DealId             *int64                 `json:"DealId,omitempty" xml:"DealId,omitempty"`
	EndCaseGmtCreate   *int64                 `json:"EndCaseGmtCreate,omitempty" xml:"EndCaseGmtCreate,omitempty"`
	Extra              map[string]interface{} `json:"Extra,omitempty" xml:"Extra,omitempty"`
	InstanceId         *string                `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MemberId           *int64                 `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	PageSize           *int32                 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SrType             *int32                 `json:"SrType,omitempty" xml:"SrType,omitempty"`
	StartCaseGmtCreate *int64                 `json:"StartCaseGmtCreate,omitempty" xml:"StartCaseGmtCreate,omitempty"`
	TaskStatus         *int32                 `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	TouchId            *int64                 `json:"TouchId,omitempty" xml:"TouchId,omitempty"`
}

func (s QueryRelationTicketsRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryRelationTicketsRequest) GoString() string {
	return s.String()
}

func (s *QueryRelationTicketsRequest) SetCaseId(v int64) *QueryRelationTicketsRequest {
	s.CaseId = &v
	return s
}

func (s *QueryRelationTicketsRequest) SetCaseType(v int32) *QueryRelationTicketsRequest {
	s.CaseType = &v
	return s
}

func (s *QueryRelationTicketsRequest) SetChannelId(v string) *QueryRelationTicketsRequest {
	s.ChannelId = &v
	return s
}

func (s *QueryRelationTicketsRequest) SetChannelType(v int32) *QueryRelationTicketsRequest {
	s.ChannelType = &v
	return s
}

func (s *QueryRelationTicketsRequest) SetCurrentPage(v int32) *QueryRelationTicketsRequest {
	s.CurrentPage = &v
	return s
}

func (s *QueryRelationTicketsRequest) SetDealId(v int64) *QueryRelationTicketsRequest {
	s.DealId = &v
	return s
}

func (s *QueryRelationTicketsRequest) SetEndCaseGmtCreate(v int64) *QueryRelationTicketsRequest {
	s.EndCaseGmtCreate = &v
	return s
}

func (s *QueryRelationTicketsRequest) SetExtra(v map[string]interface{}) *QueryRelationTicketsRequest {
	s.Extra = v
	return s
}

func (s *QueryRelationTicketsRequest) SetInstanceId(v string) *QueryRelationTicketsRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryRelationTicketsRequest) SetMemberId(v int64) *QueryRelationTicketsRequest {
	s.MemberId = &v
	return s
}

func (s *QueryRelationTicketsRequest) SetPageSize(v int32) *QueryRelationTicketsRequest {
	s.PageSize = &v
	return s
}

func (s *QueryRelationTicketsRequest) SetSrType(v int32) *QueryRelationTicketsRequest {
	s.SrType = &v
	return s
}

func (s *QueryRelationTicketsRequest) SetStartCaseGmtCreate(v int64) *QueryRelationTicketsRequest {
	s.StartCaseGmtCreate = &v
	return s
}

func (s *QueryRelationTicketsRequest) SetTaskStatus(v int32) *QueryRelationTicketsRequest {
	s.TaskStatus = &v
	return s
}

func (s *QueryRelationTicketsRequest) SetTouchId(v int64) *QueryRelationTicketsRequest {
	s.TouchId = &v
	return s
}

type QueryRelationTicketsResponseBody struct {
	CnePageSize  *int32                                  `json:"CnePageSize,omitempty" xml:"CnePageSize,omitempty"`
	Code         *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	CurrentPage  *int32                                  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Data         []*QueryRelationTicketsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	FirstResult  *int32                                  `json:"FirstResult,omitempty" xml:"FirstResult,omitempty"`
	Message      *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	NextPage     *int32                                  `json:"NextPage,omitempty" xml:"NextPage,omitempty"`
	RequestId    *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                                   `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalPage    *int32                                  `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
	TotalResults *int32                                  `json:"TotalResults,omitempty" xml:"TotalResults,omitempty"`
}

func (s QueryRelationTicketsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryRelationTicketsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryRelationTicketsResponseBody) SetCnePageSize(v int32) *QueryRelationTicketsResponseBody {
	s.CnePageSize = &v
	return s
}

func (s *QueryRelationTicketsResponseBody) SetCode(v string) *QueryRelationTicketsResponseBody {
	s.Code = &v
	return s
}

func (s *QueryRelationTicketsResponseBody) SetCurrentPage(v int32) *QueryRelationTicketsResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *QueryRelationTicketsResponseBody) SetData(v []*QueryRelationTicketsResponseBodyData) *QueryRelationTicketsResponseBody {
	s.Data = v
	return s
}

func (s *QueryRelationTicketsResponseBody) SetFirstResult(v int32) *QueryRelationTicketsResponseBody {
	s.FirstResult = &v
	return s
}

func (s *QueryRelationTicketsResponseBody) SetMessage(v string) *QueryRelationTicketsResponseBody {
	s.Message = &v
	return s
}

func (s *QueryRelationTicketsResponseBody) SetNextPage(v int32) *QueryRelationTicketsResponseBody {
	s.NextPage = &v
	return s
}

func (s *QueryRelationTicketsResponseBody) SetRequestId(v string) *QueryRelationTicketsResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryRelationTicketsResponseBody) SetSuccess(v bool) *QueryRelationTicketsResponseBody {
	s.Success = &v
	return s
}

func (s *QueryRelationTicketsResponseBody) SetTotalPage(v int32) *QueryRelationTicketsResponseBody {
	s.TotalPage = &v
	return s
}

func (s *QueryRelationTicketsResponseBody) SetTotalResults(v int32) *QueryRelationTicketsResponseBody {
	s.TotalResults = &v
	return s
}

type QueryRelationTicketsResponseBodyData struct {
	AssignTime       *int64                                              `json:"AssignTime,omitempty" xml:"AssignTime,omitempty"`
	BuId             *int64                                              `json:"BuId,omitempty" xml:"BuId,omitempty"`
	CaseBuId         *int64                                              `json:"CaseBuId,omitempty" xml:"CaseBuId,omitempty"`
	CaseDepartmentId *int64                                              `json:"CaseDepartmentId,omitempty" xml:"CaseDepartmentId,omitempty"`
	CaseGmtCreate    *int64                                              `json:"CaseGmtCreate,omitempty" xml:"CaseGmtCreate,omitempty"`
	CaseId           *int64                                              `json:"CaseId,omitempty" xml:"CaseId,omitempty"`
	CaseStatus       *int32                                              `json:"CaseStatus,omitempty" xml:"CaseStatus,omitempty"`
	CaseType         *int32                                              `json:"CaseType,omitempty" xml:"CaseType,omitempty"`
	ChannelId        *string                                             `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	ChannelType      *int32                                              `json:"ChannelType,omitempty" xml:"ChannelType,omitempty"`
	CreatorId        *int64                                              `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	DeadLine         *int64                                              `json:"DeadLine,omitempty" xml:"DeadLine,omitempty"`
	DealId           *int64                                              `json:"DealId,omitempty" xml:"DealId,omitempty"`
	DealTime         *int64                                              `json:"DealTime,omitempty" xml:"DealTime,omitempty"`
	DepartmentId     *int64                                              `json:"DepartmentId,omitempty" xml:"DepartmentId,omitempty"`
	ExtAttrs         map[string]interface{}                              `json:"ExtAttrs,omitempty" xml:"ExtAttrs,omitempty"`
	FormId           *int64                                              `json:"FormId,omitempty" xml:"FormId,omitempty"`
	FromInfo         *string                                             `json:"FromInfo,omitempty" xml:"FromInfo,omitempty"`
	GaseGmtModified  *int64                                              `json:"GaseGmtModified,omitempty" xml:"GaseGmtModified,omitempty"`
	GmtCreate        *int64                                              `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified      *int64                                              `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	GroupId          *int64                                              `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName        *string                                             `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	MemberId         *int64                                              `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	MemberName       *string                                             `json:"MemberName,omitempty" xml:"MemberName,omitempty"`
	Owner            *int64                                              `json:"Owner,omitempty" xml:"Owner,omitempty"`
	OwnerName        *string                                             `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	ParentId         *int64                                              `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	Priority         *int32                                              `json:"Priority,omitempty" xml:"Priority,omitempty"`
	QuestionId       *string                                             `json:"QuestionId,omitempty" xml:"QuestionId,omitempty"`
	QuestionInfo     *string                                             `json:"QuestionInfo,omitempty" xml:"QuestionInfo,omitempty"`
	RefCaseId        *int64                                              `json:"RefCaseId,omitempty" xml:"RefCaseId,omitempty"`
	RelationCase     []*QueryRelationTicketsResponseBodyDataRelationCase `json:"RelationCase,omitempty" xml:"RelationCase,omitempty" type:"Repeated"`
	ServiceType      *int32                                              `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	SrType           *int64                                              `json:"SrType,omitempty" xml:"SrType,omitempty"`
	TaskId           *int64                                              `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskStatus       *int32                                              `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	TaskType         *int32                                              `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	TemplateId       *int64                                              `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	Title            *string                                             `json:"Title,omitempty" xml:"Title,omitempty"`
	UserServiceId    *int64                                              `json:"UserServiceId,omitempty" xml:"UserServiceId,omitempty"`
}

func (s QueryRelationTicketsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryRelationTicketsResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryRelationTicketsResponseBodyData) SetAssignTime(v int64) *QueryRelationTicketsResponseBodyData {
	s.AssignTime = &v
	return s
}

func (s *QueryRelationTicketsResponseBodyData) SetBuId(v int64) *QueryRelationTicketsResponseBodyData {
	s.BuId = &v
	return s
}

func (s *QueryRelationTicketsResponseBodyData) SetCaseBuId(v int64) *QueryRelationTicketsResponseBodyData {
	s.CaseBuId = &v
	return s
}

func (s *QueryRelationTicketsResponseBodyData) SetCaseDepartmentId(v int64) *QueryRelationTicketsResponseBodyData {
	s.CaseDepartmentId = &v
	return s
}

func (s *QueryRelationTicketsResponseBodyData) SetCaseGmtCreate(v int64) *QueryRelationTicketsResponseBodyData {
	s.CaseGmtCreate = &v
	return s
}

func (s *QueryRelationTicketsResponseBodyData) SetCaseId(v int64) *QueryRelationTicketsResponseBodyData {
	s.CaseId = &v
	return s
}

func (s *QueryRelationTicketsResponseBodyData) SetCaseStatus(v int32) *QueryRelationTicketsResponseBodyData {
	s.CaseStatus = &v
	return s
}

func (s *QueryRelationTicketsResponseBodyData) SetCaseType(v int32) *QueryRelationTicketsResponseBodyData {
	s.CaseType = &v
	return s
}

func (s *QueryRelationTicketsResponseBodyData) SetChannelId(v string) *QueryRelationTicketsResponseBodyData {
	s.ChannelId = &v
	return s
}

func (s *QueryRelationTicketsResponseBodyData) SetChannelType(v int32) *QueryRelationTicketsResponseBodyData {
	s.ChannelType = &v
	return s
}

func (s *QueryRelationTicketsResponseBodyData) SetCreatorId(v int64) *QueryRelationTicketsResponseBodyData {
	s.CreatorId = &v
	return s
}

func (s *QueryRelationTicketsResponseBodyData) SetDeadLine(v int64) *QueryRelationTicketsResponseBodyData {
	s.DeadLine = &v
	return s
}

func (s *QueryRelationTicketsResponseBodyData) SetDealId(v int64) *QueryRelationTicketsResponseBodyData {
	s.DealId = &v
	return s
}

func (s *QueryRelationTicketsResponseBodyData) SetDealTime(v int64) *QueryRelationTicketsResponseBodyData {
	s.DealTime = &v
	return s
}

func (s *QueryRelationTicketsResponseBodyData) SetDepartmentId(v int64) *QueryRelationTicketsResponseBodyData {
	s.DepartmentId = &v
	return s
}

func (s *QueryRelationTicketsResponseBodyData) SetExtAttrs(v map[string]interface{}) *QueryRelationTicketsResponseBodyData {
	s.ExtAttrs = v
	return s
}

func (s *QueryRelationTicketsResponseBodyData) SetFormId(v int64) *QueryRelationTicketsResponseBodyData {
	s.FormId = &v
	return s
}

func (s *QueryRelationTicketsResponseBodyData) SetFromInfo(v string) *QueryRelationTicketsResponseBodyData {
	s.FromInfo = &v
	return s
}

func (s *QueryRelationTicketsResponseBodyData) SetGaseGmtModified(v int64) *QueryRelationTicketsResponseBodyData {
	s.GaseGmtModified = &v
	return s
}

func (s *QueryRelationTicketsResponseBodyData) SetGmtCreate(v int64) *QueryRelationTicketsResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *QueryRelationTicketsResponseBodyData) SetGmtModified(v int64) *QueryRelationTicketsResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *QueryRelationTicketsResponseBodyData) SetGroupId(v int64) *QueryRelationTicketsResponseBodyData {
	s.GroupId = &v
	return s
}

func (s *QueryRelationTicketsResponseBodyData) SetGroupName(v string) *QueryRelationTicketsResponseBodyData {
	s.GroupName = &v
	return s
}

func (s *QueryRelationTicketsResponseBodyData) SetMemberId(v int64) *QueryRelationTicketsResponseBodyData {
	s.MemberId = &v
	return s
}

func (s *QueryRelationTicketsResponseBodyData) SetMemberName(v string) *QueryRelationTicketsResponseBodyData {
	s.MemberName = &v
	return s
}

func (s *QueryRelationTicketsResponseBodyData) SetOwner(v int64) *QueryRelationTicketsResponseBodyData {
	s.Owner = &v
	return s
}

func (s *QueryRelationTicketsResponseBodyData) SetOwnerName(v string) *QueryRelationTicketsResponseBodyData {
	s.OwnerName = &v
	return s
}

func (s *QueryRelationTicketsResponseBodyData) SetParentId(v int64) *QueryRelationTicketsResponseBodyData {
	s.ParentId = &v
	return s
}

func (s *QueryRelationTicketsResponseBodyData) SetPriority(v int32) *QueryRelationTicketsResponseBodyData {
	s.Priority = &v
	return s
}

func (s *QueryRelationTicketsResponseBodyData) SetQuestionId(v string) *QueryRelationTicketsResponseBodyData {
	s.QuestionId = &v
	return s
}

func (s *QueryRelationTicketsResponseBodyData) SetQuestionInfo(v string) *QueryRelationTicketsResponseBodyData {
	s.QuestionInfo = &v
	return s
}

func (s *QueryRelationTicketsResponseBodyData) SetRefCaseId(v int64) *QueryRelationTicketsResponseBodyData {
	s.RefCaseId = &v
	return s
}

func (s *QueryRelationTicketsResponseBodyData) SetRelationCase(v []*QueryRelationTicketsResponseBodyDataRelationCase) *QueryRelationTicketsResponseBodyData {
	s.RelationCase = v
	return s
}

func (s *QueryRelationTicketsResponseBodyData) SetServiceType(v int32) *QueryRelationTicketsResponseBodyData {
	s.ServiceType = &v
	return s
}

func (s *QueryRelationTicketsResponseBodyData) SetSrType(v int64) *QueryRelationTicketsResponseBodyData {
	s.SrType = &v
	return s
}

func (s *QueryRelationTicketsResponseBodyData) SetTaskId(v int64) *QueryRelationTicketsResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *QueryRelationTicketsResponseBodyData) SetTaskStatus(v int32) *QueryRelationTicketsResponseBodyData {
	s.TaskStatus = &v
	return s
}

func (s *QueryRelationTicketsResponseBodyData) SetTaskType(v int32) *QueryRelationTicketsResponseBodyData {
	s.TaskType = &v
	return s
}

func (s *QueryRelationTicketsResponseBodyData) SetTemplateId(v int64) *QueryRelationTicketsResponseBodyData {
	s.TemplateId = &v
	return s
}

func (s *QueryRelationTicketsResponseBodyData) SetTitle(v string) *QueryRelationTicketsResponseBodyData {
	s.Title = &v
	return s
}

func (s *QueryRelationTicketsResponseBodyData) SetUserServiceId(v int64) *QueryRelationTicketsResponseBodyData {
	s.UserServiceId = &v
	return s
}

type QueryRelationTicketsResponseBodyDataRelationCase struct {
	AssignTime               *int64                 `json:"AssignTime,omitempty" xml:"AssignTime,omitempty"`
	BuId                     *int64                 `json:"BuId,omitempty" xml:"BuId,omitempty"`
	CaseBuId                 *int64                 `json:"CaseBuId,omitempty" xml:"CaseBuId,omitempty"`
	CaseDepartmentId         *int64                 `json:"CaseDepartmentId,omitempty" xml:"CaseDepartmentId,omitempty"`
	CaseGmtCreate            *int64                 `json:"CaseGmtCreate,omitempty" xml:"CaseGmtCreate,omitempty"`
	CaseGmtModified          *int64                 `json:"CaseGmtModified,omitempty" xml:"CaseGmtModified,omitempty"`
	CaseId                   *int64                 `json:"CaseId,omitempty" xml:"CaseId,omitempty"`
	CaseStatus               *int32                 `json:"CaseStatus,omitempty" xml:"CaseStatus,omitempty"`
	CaseType                 *int32                 `json:"CaseType,omitempty" xml:"CaseType,omitempty"`
	ChannelId                *string                `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	ChannelType              *int32                 `json:"ChannelType,omitempty" xml:"ChannelType,omitempty"`
	CreatorId                *int64                 `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	DeadLine                 *int64                 `json:"DeadLine,omitempty" xml:"DeadLine,omitempty"`
	DealId                   *int64                 `json:"DealId,omitempty" xml:"DealId,omitempty"`
	DealTime                 *int64                 `json:"DealTime,omitempty" xml:"DealTime,omitempty"`
	DepartmentId             *int64                 `json:"DepartmentId,omitempty" xml:"DepartmentId,omitempty"`
	ExtAttrs                 map[string]interface{} `json:"ExtAttrs,omitempty" xml:"ExtAttrs,omitempty"`
	FormId                   *int64                 `json:"FormId,omitempty" xml:"FormId,omitempty"`
	FromInfo                 *string                `json:"FromInfo,omitempty" xml:"FromInfo,omitempty"`
	GmtCreate                *int64                 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified              *int64                 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	GroupId                  *int64                 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName                *string                `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	MemberId                 *int64                 `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	MemberName               *string                `json:"MemberName,omitempty" xml:"MemberName,omitempty"`
	Owner                    *int64                 `json:"Owner,omitempty" xml:"Owner,omitempty"`
	OwnerName                *string                `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	ParentId                 *int64                 `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	Priority                 *int32                 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	QuestionId               *string                `json:"QuestionId,omitempty" xml:"QuestionId,omitempty"`
	QuestionInfoQuestionInfo *string                `json:"QuestionInfoQuestionInfo,omitempty" xml:"QuestionInfoQuestionInfo,omitempty"`
	RefCaseId                *int64                 `json:"RefCaseId,omitempty" xml:"RefCaseId,omitempty"`
	ServiceType              *int32                 `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	TaskId                   *int64                 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskStatus               *int32                 `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	TaskType                 *int32                 `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	TemplateId               *int64                 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	Title                    *string                `json:"Title,omitempty" xml:"Title,omitempty"`
	UserServiceId            *int64                 `json:"UserServiceId,omitempty" xml:"UserServiceId,omitempty"`
}

func (s QueryRelationTicketsResponseBodyDataRelationCase) String() string {
	return tea.Prettify(s)
}

func (s QueryRelationTicketsResponseBodyDataRelationCase) GoString() string {
	return s.String()
}

func (s *QueryRelationTicketsResponseBodyDataRelationCase) SetAssignTime(v int64) *QueryRelationTicketsResponseBodyDataRelationCase {
	s.AssignTime = &v
	return s
}

func (s *QueryRelationTicketsResponseBodyDataRelationCase) SetBuId(v int64) *QueryRelationTicketsResponseBodyDataRelationCase {
	s.BuId = &v
	return s
}

func (s *QueryRelationTicketsResponseBodyDataRelationCase) SetCaseBuId(v int64) *QueryRelationTicketsResponseBodyDataRelationCase {
	s.CaseBuId = &v
	return s
}

func (s *QueryRelationTicketsResponseBodyDataRelationCase) SetCaseDepartmentId(v int64) *QueryRelationTicketsResponseBodyDataRelationCase {
	s.CaseDepartmentId = &v
	return s
}

func (s *QueryRelationTicketsResponseBodyDataRelationCase) SetCaseGmtCreate(v int64) *QueryRelationTicketsResponseBodyDataRelationCase {
	s.CaseGmtCreate = &v
	return s
}

func (s *QueryRelationTicketsResponseBodyDataRelationCase) SetCaseGmtModified(v int64) *QueryRelationTicketsResponseBodyDataRelationCase {
	s.CaseGmtModified = &v
	return s
}

func (s *QueryRelationTicketsResponseBodyDataRelationCase) SetCaseId(v int64) *QueryRelationTicketsResponseBodyDataRelationCase {
	s.CaseId = &v
	return s
}

func (s *QueryRelationTicketsResponseBodyDataRelationCase) SetCaseStatus(v int32) *QueryRelationTicketsResponseBodyDataRelationCase {
	s.CaseStatus = &v
	return s
}

func (s *QueryRelationTicketsResponseBodyDataRelationCase) SetCaseType(v int32) *QueryRelationTicketsResponseBodyDataRelationCase {
	s.CaseType = &v
	return s
}

func (s *QueryRelationTicketsResponseBodyDataRelationCase) SetChannelId(v string) *QueryRelationTicketsResponseBodyDataRelationCase {
	s.ChannelId = &v
	return s
}

func (s *QueryRelationTicketsResponseBodyDataRelationCase) SetChannelType(v int32) *QueryRelationTicketsResponseBodyDataRelationCase {
	s.ChannelType = &v
	return s
}

func (s *QueryRelationTicketsResponseBodyDataRelationCase) SetCreatorId(v int64) *QueryRelationTicketsResponseBodyDataRelationCase {
	s.CreatorId = &v
	return s
}

func (s *QueryRelationTicketsResponseBodyDataRelationCase) SetDeadLine(v int64) *QueryRelationTicketsResponseBodyDataRelationCase {
	s.DeadLine = &v
	return s
}

func (s *QueryRelationTicketsResponseBodyDataRelationCase) SetDealId(v int64) *QueryRelationTicketsResponseBodyDataRelationCase {
	s.DealId = &v
	return s
}

func (s *QueryRelationTicketsResponseBodyDataRelationCase) SetDealTime(v int64) *QueryRelationTicketsResponseBodyDataRelationCase {
	s.DealTime = &v
	return s
}

func (s *QueryRelationTicketsResponseBodyDataRelationCase) SetDepartmentId(v int64) *QueryRelationTicketsResponseBodyDataRelationCase {
	s.DepartmentId = &v
	return s
}

func (s *QueryRelationTicketsResponseBodyDataRelationCase) SetExtAttrs(v map[string]interface{}) *QueryRelationTicketsResponseBodyDataRelationCase {
	s.ExtAttrs = v
	return s
}

func (s *QueryRelationTicketsResponseBodyDataRelationCase) SetFormId(v int64) *QueryRelationTicketsResponseBodyDataRelationCase {
	s.FormId = &v
	return s
}

func (s *QueryRelationTicketsResponseBodyDataRelationCase) SetFromInfo(v string) *QueryRelationTicketsResponseBodyDataRelationCase {
	s.FromInfo = &v
	return s
}

func (s *QueryRelationTicketsResponseBodyDataRelationCase) SetGmtCreate(v int64) *QueryRelationTicketsResponseBodyDataRelationCase {
	s.GmtCreate = &v
	return s
}

func (s *QueryRelationTicketsResponseBodyDataRelationCase) SetGmtModified(v int64) *QueryRelationTicketsResponseBodyDataRelationCase {
	s.GmtModified = &v
	return s
}

func (s *QueryRelationTicketsResponseBodyDataRelationCase) SetGroupId(v int64) *QueryRelationTicketsResponseBodyDataRelationCase {
	s.GroupId = &v
	return s
}

func (s *QueryRelationTicketsResponseBodyDataRelationCase) SetGroupName(v string) *QueryRelationTicketsResponseBodyDataRelationCase {
	s.GroupName = &v
	return s
}

func (s *QueryRelationTicketsResponseBodyDataRelationCase) SetMemberId(v int64) *QueryRelationTicketsResponseBodyDataRelationCase {
	s.MemberId = &v
	return s
}

func (s *QueryRelationTicketsResponseBodyDataRelationCase) SetMemberName(v string) *QueryRelationTicketsResponseBodyDataRelationCase {
	s.MemberName = &v
	return s
}

func (s *QueryRelationTicketsResponseBodyDataRelationCase) SetOwner(v int64) *QueryRelationTicketsResponseBodyDataRelationCase {
	s.Owner = &v
	return s
}

func (s *QueryRelationTicketsResponseBodyDataRelationCase) SetOwnerName(v string) *QueryRelationTicketsResponseBodyDataRelationCase {
	s.OwnerName = &v
	return s
}

func (s *QueryRelationTicketsResponseBodyDataRelationCase) SetParentId(v int64) *QueryRelationTicketsResponseBodyDataRelationCase {
	s.ParentId = &v
	return s
}

func (s *QueryRelationTicketsResponseBodyDataRelationCase) SetPriority(v int32) *QueryRelationTicketsResponseBodyDataRelationCase {
	s.Priority = &v
	return s
}

func (s *QueryRelationTicketsResponseBodyDataRelationCase) SetQuestionId(v string) *QueryRelationTicketsResponseBodyDataRelationCase {
	s.QuestionId = &v
	return s
}

func (s *QueryRelationTicketsResponseBodyDataRelationCase) SetQuestionInfoQuestionInfo(v string) *QueryRelationTicketsResponseBodyDataRelationCase {
	s.QuestionInfoQuestionInfo = &v
	return s
}

func (s *QueryRelationTicketsResponseBodyDataRelationCase) SetRefCaseId(v int64) *QueryRelationTicketsResponseBodyDataRelationCase {
	s.RefCaseId = &v
	return s
}

func (s *QueryRelationTicketsResponseBodyDataRelationCase) SetServiceType(v int32) *QueryRelationTicketsResponseBodyDataRelationCase {
	s.ServiceType = &v
	return s
}

func (s *QueryRelationTicketsResponseBodyDataRelationCase) SetTaskId(v int64) *QueryRelationTicketsResponseBodyDataRelationCase {
	s.TaskId = &v
	return s
}

func (s *QueryRelationTicketsResponseBodyDataRelationCase) SetTaskStatus(v int32) *QueryRelationTicketsResponseBodyDataRelationCase {
	s.TaskStatus = &v
	return s
}

func (s *QueryRelationTicketsResponseBodyDataRelationCase) SetTaskType(v int32) *QueryRelationTicketsResponseBodyDataRelationCase {
	s.TaskType = &v
	return s
}

func (s *QueryRelationTicketsResponseBodyDataRelationCase) SetTemplateId(v int64) *QueryRelationTicketsResponseBodyDataRelationCase {
	s.TemplateId = &v
	return s
}

func (s *QueryRelationTicketsResponseBodyDataRelationCase) SetTitle(v string) *QueryRelationTicketsResponseBodyDataRelationCase {
	s.Title = &v
	return s
}

func (s *QueryRelationTicketsResponseBodyDataRelationCase) SetUserServiceId(v int64) *QueryRelationTicketsResponseBodyDataRelationCase {
	s.UserServiceId = &v
	return s
}

type QueryRelationTicketsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryRelationTicketsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryRelationTicketsResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryRelationTicketsResponse) GoString() string {
	return s.String()
}

func (s *QueryRelationTicketsResponse) SetHeaders(v map[string]*string) *QueryRelationTicketsResponse {
	s.Headers = v
	return s
}

func (s *QueryRelationTicketsResponse) SetStatusCode(v int32) *QueryRelationTicketsResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryRelationTicketsResponse) SetBody(v *QueryRelationTicketsResponseBody) *QueryRelationTicketsResponse {
	s.Body = v
	return s
}

type QueryRingDetailListRequest struct {
	CallOutStatus    *string                `json:"CallOutStatus,omitempty" xml:"CallOutStatus,omitempty"`
	EndDate          *int64                 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	Extra            *string                `json:"Extra,omitempty" xml:"Extra,omitempty"`
	FromPhoneNumList map[string]interface{} `json:"FromPhoneNumList,omitempty" xml:"FromPhoneNumList,omitempty"`
	InstanceId       *string                `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PageNo           *int32                 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize         *int32                 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	StartDate        *int64                 `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	ToPhoneNumList   map[string]interface{} `json:"ToPhoneNumList,omitempty" xml:"ToPhoneNumList,omitempty"`
}

func (s QueryRingDetailListRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryRingDetailListRequest) GoString() string {
	return s.String()
}

func (s *QueryRingDetailListRequest) SetCallOutStatus(v string) *QueryRingDetailListRequest {
	s.CallOutStatus = &v
	return s
}

func (s *QueryRingDetailListRequest) SetEndDate(v int64) *QueryRingDetailListRequest {
	s.EndDate = &v
	return s
}

func (s *QueryRingDetailListRequest) SetExtra(v string) *QueryRingDetailListRequest {
	s.Extra = &v
	return s
}

func (s *QueryRingDetailListRequest) SetFromPhoneNumList(v map[string]interface{}) *QueryRingDetailListRequest {
	s.FromPhoneNumList = v
	return s
}

func (s *QueryRingDetailListRequest) SetInstanceId(v string) *QueryRingDetailListRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryRingDetailListRequest) SetPageNo(v int32) *QueryRingDetailListRequest {
	s.PageNo = &v
	return s
}

func (s *QueryRingDetailListRequest) SetPageSize(v int32) *QueryRingDetailListRequest {
	s.PageSize = &v
	return s
}

func (s *QueryRingDetailListRequest) SetStartDate(v int64) *QueryRingDetailListRequest {
	s.StartDate = &v
	return s
}

func (s *QueryRingDetailListRequest) SetToPhoneNumList(v map[string]interface{}) *QueryRingDetailListRequest {
	s.ToPhoneNumList = v
	return s
}

type QueryRingDetailListShrinkRequest struct {
	CallOutStatus          *string `json:"CallOutStatus,omitempty" xml:"CallOutStatus,omitempty"`
	EndDate                *int64  `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	Extra                  *string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	FromPhoneNumListShrink *string `json:"FromPhoneNumList,omitempty" xml:"FromPhoneNumList,omitempty"`
	InstanceId             *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PageNo                 *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize               *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	StartDate              *int64  `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	ToPhoneNumListShrink   *string `json:"ToPhoneNumList,omitempty" xml:"ToPhoneNumList,omitempty"`
}

func (s QueryRingDetailListShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryRingDetailListShrinkRequest) GoString() string {
	return s.String()
}

func (s *QueryRingDetailListShrinkRequest) SetCallOutStatus(v string) *QueryRingDetailListShrinkRequest {
	s.CallOutStatus = &v
	return s
}

func (s *QueryRingDetailListShrinkRequest) SetEndDate(v int64) *QueryRingDetailListShrinkRequest {
	s.EndDate = &v
	return s
}

func (s *QueryRingDetailListShrinkRequest) SetExtra(v string) *QueryRingDetailListShrinkRequest {
	s.Extra = &v
	return s
}

func (s *QueryRingDetailListShrinkRequest) SetFromPhoneNumListShrink(v string) *QueryRingDetailListShrinkRequest {
	s.FromPhoneNumListShrink = &v
	return s
}

func (s *QueryRingDetailListShrinkRequest) SetInstanceId(v string) *QueryRingDetailListShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryRingDetailListShrinkRequest) SetPageNo(v int32) *QueryRingDetailListShrinkRequest {
	s.PageNo = &v
	return s
}

func (s *QueryRingDetailListShrinkRequest) SetPageSize(v int32) *QueryRingDetailListShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *QueryRingDetailListShrinkRequest) SetStartDate(v int64) *QueryRingDetailListShrinkRequest {
	s.StartDate = &v
	return s
}

func (s *QueryRingDetailListShrinkRequest) SetToPhoneNumListShrink(v string) *QueryRingDetailListShrinkRequest {
	s.ToPhoneNumListShrink = &v
	return s
}

type QueryRingDetailListResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *string `json:"Data,omitempty" xml:"Data,omitempty"`
	HttpStatusCode *int64  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryRingDetailListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryRingDetailListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryRingDetailListResponseBody) SetCode(v string) *QueryRingDetailListResponseBody {
	s.Code = &v
	return s
}

func (s *QueryRingDetailListResponseBody) SetData(v string) *QueryRingDetailListResponseBody {
	s.Data = &v
	return s
}

func (s *QueryRingDetailListResponseBody) SetHttpStatusCode(v int64) *QueryRingDetailListResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *QueryRingDetailListResponseBody) SetMessage(v string) *QueryRingDetailListResponseBody {
	s.Message = &v
	return s
}

func (s *QueryRingDetailListResponseBody) SetRequestId(v string) *QueryRingDetailListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryRingDetailListResponseBody) SetSuccess(v bool) *QueryRingDetailListResponseBody {
	s.Success = &v
	return s
}

type QueryRingDetailListResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryRingDetailListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryRingDetailListResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryRingDetailListResponse) GoString() string {
	return s.String()
}

func (s *QueryRingDetailListResponse) SetHeaders(v map[string]*string) *QueryRingDetailListResponse {
	s.Headers = v
	return s
}

func (s *QueryRingDetailListResponse) SetStatusCode(v int32) *QueryRingDetailListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryRingDetailListResponse) SetBody(v *QueryRingDetailListResponseBody) *QueryRingDetailListResponse {
	s.Body = v
	return s
}

type QueryServiceConfigRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	ViewCode   *string `json:"ViewCode,omitempty" xml:"ViewCode,omitempty"`
}

func (s QueryServiceConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryServiceConfigRequest) GoString() string {
	return s.String()
}

func (s *QueryServiceConfigRequest) SetInstanceId(v string) *QueryServiceConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryServiceConfigRequest) SetParameters(v string) *QueryServiceConfigRequest {
	s.Parameters = &v
	return s
}

func (s *QueryServiceConfigRequest) SetViewCode(v string) *QueryServiceConfigRequest {
	s.ViewCode = &v
	return s
}

type QueryServiceConfigResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryServiceConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryServiceConfigResponseBody) GoString() string {
	return s.String()
}

func (s *QueryServiceConfigResponseBody) SetCode(v string) *QueryServiceConfigResponseBody {
	s.Code = &v
	return s
}

func (s *QueryServiceConfigResponseBody) SetData(v string) *QueryServiceConfigResponseBody {
	s.Data = &v
	return s
}

func (s *QueryServiceConfigResponseBody) SetMessage(v string) *QueryServiceConfigResponseBody {
	s.Message = &v
	return s
}

func (s *QueryServiceConfigResponseBody) SetRequestId(v string) *QueryServiceConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryServiceConfigResponseBody) SetSuccess(v bool) *QueryServiceConfigResponseBody {
	s.Success = &v
	return s
}

type QueryServiceConfigResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryServiceConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryServiceConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryServiceConfigResponse) GoString() string {
	return s.String()
}

func (s *QueryServiceConfigResponse) SetHeaders(v map[string]*string) *QueryServiceConfigResponse {
	s.Headers = v
	return s
}

func (s *QueryServiceConfigResponse) SetStatusCode(v int32) *QueryServiceConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryServiceConfigResponse) SetBody(v *QueryServiceConfigResponseBody) *QueryServiceConfigResponse {
	s.Body = v
	return s
}

type QueryServicerByDepartmentAndMixNameRequest struct {
	CurrentPageNum   *int64   `json:"CurrentPageNum,omitempty" xml:"CurrentPageNum,omitempty"`
	DepartmentIdList []*int64 `json:"DepartmentIdList,omitempty" xml:"DepartmentIdList,omitempty" type:"Repeated"`
	InstanceId       *string  `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	KeyWord          *string  `json:"KeyWord,omitempty" xml:"KeyWord,omitempty"`
	PageSize         *int64   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s QueryServicerByDepartmentAndMixNameRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryServicerByDepartmentAndMixNameRequest) GoString() string {
	return s.String()
}

func (s *QueryServicerByDepartmentAndMixNameRequest) SetCurrentPageNum(v int64) *QueryServicerByDepartmentAndMixNameRequest {
	s.CurrentPageNum = &v
	return s
}

func (s *QueryServicerByDepartmentAndMixNameRequest) SetDepartmentIdList(v []*int64) *QueryServicerByDepartmentAndMixNameRequest {
	s.DepartmentIdList = v
	return s
}

func (s *QueryServicerByDepartmentAndMixNameRequest) SetInstanceId(v string) *QueryServicerByDepartmentAndMixNameRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryServicerByDepartmentAndMixNameRequest) SetKeyWord(v string) *QueryServicerByDepartmentAndMixNameRequest {
	s.KeyWord = &v
	return s
}

func (s *QueryServicerByDepartmentAndMixNameRequest) SetPageSize(v int64) *QueryServicerByDepartmentAndMixNameRequest {
	s.PageSize = &v
	return s
}

type QueryServicerByDepartmentAndMixNameShrinkRequest struct {
	CurrentPageNum         *int64  `json:"CurrentPageNum,omitempty" xml:"CurrentPageNum,omitempty"`
	DepartmentIdListShrink *string `json:"DepartmentIdList,omitempty" xml:"DepartmentIdList,omitempty"`
	InstanceId             *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	KeyWord                *string `json:"KeyWord,omitempty" xml:"KeyWord,omitempty"`
	PageSize               *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s QueryServicerByDepartmentAndMixNameShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryServicerByDepartmentAndMixNameShrinkRequest) GoString() string {
	return s.String()
}

func (s *QueryServicerByDepartmentAndMixNameShrinkRequest) SetCurrentPageNum(v int64) *QueryServicerByDepartmentAndMixNameShrinkRequest {
	s.CurrentPageNum = &v
	return s
}

func (s *QueryServicerByDepartmentAndMixNameShrinkRequest) SetDepartmentIdListShrink(v string) *QueryServicerByDepartmentAndMixNameShrinkRequest {
	s.DepartmentIdListShrink = &v
	return s
}

func (s *QueryServicerByDepartmentAndMixNameShrinkRequest) SetInstanceId(v string) *QueryServicerByDepartmentAndMixNameShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryServicerByDepartmentAndMixNameShrinkRequest) SetKeyWord(v string) *QueryServicerByDepartmentAndMixNameShrinkRequest {
	s.KeyWord = &v
	return s
}

func (s *QueryServicerByDepartmentAndMixNameShrinkRequest) SetPageSize(v int64) *QueryServicerByDepartmentAndMixNameShrinkRequest {
	s.PageSize = &v
	return s
}

type QueryServicerByDepartmentAndMixNameResponseBody struct {
	Code      *string                                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *QueryServicerByDepartmentAndMixNameResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                                `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryServicerByDepartmentAndMixNameResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryServicerByDepartmentAndMixNameResponseBody) GoString() string {
	return s.String()
}

func (s *QueryServicerByDepartmentAndMixNameResponseBody) SetCode(v string) *QueryServicerByDepartmentAndMixNameResponseBody {
	s.Code = &v
	return s
}

func (s *QueryServicerByDepartmentAndMixNameResponseBody) SetData(v *QueryServicerByDepartmentAndMixNameResponseBodyData) *QueryServicerByDepartmentAndMixNameResponseBody {
	s.Data = v
	return s
}

func (s *QueryServicerByDepartmentAndMixNameResponseBody) SetMessage(v string) *QueryServicerByDepartmentAndMixNameResponseBody {
	s.Message = &v
	return s
}

func (s *QueryServicerByDepartmentAndMixNameResponseBody) SetRequestId(v string) *QueryServicerByDepartmentAndMixNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryServicerByDepartmentAndMixNameResponseBody) SetSuccess(v bool) *QueryServicerByDepartmentAndMixNameResponseBody {
	s.Success = &v
	return s
}

type QueryServicerByDepartmentAndMixNameResponseBodyData struct {
	Results []*QueryServicerByDepartmentAndMixNameResponseBodyDataResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
	Total   *int32                                                        `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s QueryServicerByDepartmentAndMixNameResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryServicerByDepartmentAndMixNameResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryServicerByDepartmentAndMixNameResponseBodyData) SetResults(v []*QueryServicerByDepartmentAndMixNameResponseBodyDataResults) *QueryServicerByDepartmentAndMixNameResponseBodyData {
	s.Results = v
	return s
}

func (s *QueryServicerByDepartmentAndMixNameResponseBodyData) SetTotal(v int32) *QueryServicerByDepartmentAndMixNameResponseBodyData {
	s.Total = &v
	return s
}

type QueryServicerByDepartmentAndMixNameResponseBodyDataResults struct {
	AccountName    *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	DepartmentId   *int64  `json:"DepartmentId,omitempty" xml:"DepartmentId,omitempty"`
	DepartmentName *string `json:"DepartmentName,omitempty" xml:"DepartmentName,omitempty"`
	RealName       *string `json:"RealName,omitempty" xml:"RealName,omitempty"`
	ServicerId     *int64  `json:"ServicerId,omitempty" xml:"ServicerId,omitempty"`
	ShowName       *string `json:"ShowName,omitempty" xml:"ShowName,omitempty"`
	UserStatus     *int32  `json:"UserStatus,omitempty" xml:"UserStatus,omitempty"`
}

func (s QueryServicerByDepartmentAndMixNameResponseBodyDataResults) String() string {
	return tea.Prettify(s)
}

func (s QueryServicerByDepartmentAndMixNameResponseBodyDataResults) GoString() string {
	return s.String()
}

func (s *QueryServicerByDepartmentAndMixNameResponseBodyDataResults) SetAccountName(v string) *QueryServicerByDepartmentAndMixNameResponseBodyDataResults {
	s.AccountName = &v
	return s
}

func (s *QueryServicerByDepartmentAndMixNameResponseBodyDataResults) SetDepartmentId(v int64) *QueryServicerByDepartmentAndMixNameResponseBodyDataResults {
	s.DepartmentId = &v
	return s
}

func (s *QueryServicerByDepartmentAndMixNameResponseBodyDataResults) SetDepartmentName(v string) *QueryServicerByDepartmentAndMixNameResponseBodyDataResults {
	s.DepartmentName = &v
	return s
}

func (s *QueryServicerByDepartmentAndMixNameResponseBodyDataResults) SetRealName(v string) *QueryServicerByDepartmentAndMixNameResponseBodyDataResults {
	s.RealName = &v
	return s
}

func (s *QueryServicerByDepartmentAndMixNameResponseBodyDataResults) SetServicerId(v int64) *QueryServicerByDepartmentAndMixNameResponseBodyDataResults {
	s.ServicerId = &v
	return s
}

func (s *QueryServicerByDepartmentAndMixNameResponseBodyDataResults) SetShowName(v string) *QueryServicerByDepartmentAndMixNameResponseBodyDataResults {
	s.ShowName = &v
	return s
}

func (s *QueryServicerByDepartmentAndMixNameResponseBodyDataResults) SetUserStatus(v int32) *QueryServicerByDepartmentAndMixNameResponseBodyDataResults {
	s.UserStatus = &v
	return s
}

type QueryServicerByDepartmentAndMixNameResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryServicerByDepartmentAndMixNameResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryServicerByDepartmentAndMixNameResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryServicerByDepartmentAndMixNameResponse) GoString() string {
	return s.String()
}

func (s *QueryServicerByDepartmentAndMixNameResponse) SetHeaders(v map[string]*string) *QueryServicerByDepartmentAndMixNameResponse {
	s.Headers = v
	return s
}

func (s *QueryServicerByDepartmentAndMixNameResponse) SetStatusCode(v int32) *QueryServicerByDepartmentAndMixNameResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryServicerByDepartmentAndMixNameResponse) SetBody(v *QueryServicerByDepartmentAndMixNameResponseBody) *QueryServicerByDepartmentAndMixNameResponse {
	s.Body = v
	return s
}

type QueryServicerByIdRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ServicerId *int64  `json:"ServicerId,omitempty" xml:"ServicerId,omitempty"`
}

func (s QueryServicerByIdRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryServicerByIdRequest) GoString() string {
	return s.String()
}

func (s *QueryServicerByIdRequest) SetInstanceId(v string) *QueryServicerByIdRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryServicerByIdRequest) SetServicerId(v int64) *QueryServicerByIdRequest {
	s.ServicerId = &v
	return s
}

type QueryServicerByIdResponseBody struct {
	Code      *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *QueryServicerByIdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryServicerByIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryServicerByIdResponseBody) GoString() string {
	return s.String()
}

func (s *QueryServicerByIdResponseBody) SetCode(v string) *QueryServicerByIdResponseBody {
	s.Code = &v
	return s
}

func (s *QueryServicerByIdResponseBody) SetData(v *QueryServicerByIdResponseBodyData) *QueryServicerByIdResponseBody {
	s.Data = v
	return s
}

func (s *QueryServicerByIdResponseBody) SetMessage(v string) *QueryServicerByIdResponseBody {
	s.Message = &v
	return s
}

func (s *QueryServicerByIdResponseBody) SetRequestId(v string) *QueryServicerByIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryServicerByIdResponseBody) SetSuccess(v bool) *QueryServicerByIdResponseBody {
	s.Success = &v
	return s
}

type QueryServicerByIdResponseBodyData struct {
	AccountName  *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	DepartmentId *int64  `json:"DepartmentId,omitempty" xml:"DepartmentId,omitempty"`
	RealName     *string `json:"RealName,omitempty" xml:"RealName,omitempty"`
	ServicerId   *int64  `json:"ServicerId,omitempty" xml:"ServicerId,omitempty"`
	ShowName     *string `json:"ShowName,omitempty" xml:"ShowName,omitempty"`
	UserStatus   *int32  `json:"UserStatus,omitempty" xml:"UserStatus,omitempty"`
}

func (s QueryServicerByIdResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryServicerByIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryServicerByIdResponseBodyData) SetAccountName(v string) *QueryServicerByIdResponseBodyData {
	s.AccountName = &v
	return s
}

func (s *QueryServicerByIdResponseBodyData) SetDepartmentId(v int64) *QueryServicerByIdResponseBodyData {
	s.DepartmentId = &v
	return s
}

func (s *QueryServicerByIdResponseBodyData) SetRealName(v string) *QueryServicerByIdResponseBodyData {
	s.RealName = &v
	return s
}

func (s *QueryServicerByIdResponseBodyData) SetServicerId(v int64) *QueryServicerByIdResponseBodyData {
	s.ServicerId = &v
	return s
}

func (s *QueryServicerByIdResponseBodyData) SetShowName(v string) *QueryServicerByIdResponseBodyData {
	s.ShowName = &v
	return s
}

func (s *QueryServicerByIdResponseBodyData) SetUserStatus(v int32) *QueryServicerByIdResponseBodyData {
	s.UserStatus = &v
	return s
}

type QueryServicerByIdResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryServicerByIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryServicerByIdResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryServicerByIdResponse) GoString() string {
	return s.String()
}

func (s *QueryServicerByIdResponse) SetHeaders(v map[string]*string) *QueryServicerByIdResponse {
	s.Headers = v
	return s
}

func (s *QueryServicerByIdResponse) SetStatusCode(v int32) *QueryServicerByIdResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryServicerByIdResponse) SetBody(v *QueryServicerByIdResponseBody) *QueryServicerByIdResponse {
	s.Body = v
	return s
}

type QuerySkillGroupsRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	GroupId     *int64  `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName   *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	GroupType   *int32  `json:"GroupType,omitempty" xml:"GroupType,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PageNo      *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s QuerySkillGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s QuerySkillGroupsRequest) GoString() string {
	return s.String()
}

func (s *QuerySkillGroupsRequest) SetClientToken(v string) *QuerySkillGroupsRequest {
	s.ClientToken = &v
	return s
}

func (s *QuerySkillGroupsRequest) SetGroupId(v int64) *QuerySkillGroupsRequest {
	s.GroupId = &v
	return s
}

func (s *QuerySkillGroupsRequest) SetGroupName(v string) *QuerySkillGroupsRequest {
	s.GroupName = &v
	return s
}

func (s *QuerySkillGroupsRequest) SetGroupType(v int32) *QuerySkillGroupsRequest {
	s.GroupType = &v
	return s
}

func (s *QuerySkillGroupsRequest) SetInstanceId(v string) *QuerySkillGroupsRequest {
	s.InstanceId = &v
	return s
}

func (s *QuerySkillGroupsRequest) SetPageNo(v int32) *QuerySkillGroupsRequest {
	s.PageNo = &v
	return s
}

func (s *QuerySkillGroupsRequest) SetPageSize(v int32) *QuerySkillGroupsRequest {
	s.PageSize = &v
	return s
}

type QuerySkillGroupsResponseBody struct {
	CurrentPage  *int32                              `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Data         []*QuerySkillGroupsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	OnePageSize  *int32                              `json:"OnePageSize,omitempty" xml:"OnePageSize,omitempty"`
	RequestId    *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalPage    *int32                              `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
	TotalResults *int32                              `json:"TotalResults,omitempty" xml:"TotalResults,omitempty"`
}

func (s QuerySkillGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QuerySkillGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySkillGroupsResponseBody) SetCurrentPage(v int32) *QuerySkillGroupsResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *QuerySkillGroupsResponseBody) SetData(v []*QuerySkillGroupsResponseBodyData) *QuerySkillGroupsResponseBody {
	s.Data = v
	return s
}

func (s *QuerySkillGroupsResponseBody) SetOnePageSize(v int32) *QuerySkillGroupsResponseBody {
	s.OnePageSize = &v
	return s
}

func (s *QuerySkillGroupsResponseBody) SetRequestId(v string) *QuerySkillGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuerySkillGroupsResponseBody) SetTotalPage(v int32) *QuerySkillGroupsResponseBody {
	s.TotalPage = &v
	return s
}

func (s *QuerySkillGroupsResponseBody) SetTotalResults(v int32) *QuerySkillGroupsResponseBody {
	s.TotalResults = &v
	return s
}

type QuerySkillGroupsResponseBodyData struct {
	ChannelType    *int32  `json:"ChannelType,omitempty" xml:"ChannelType,omitempty"`
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DisplayName    *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	SkillGroupId   *int64  `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
	SkillGroupName *string `json:"SkillGroupName,omitempty" xml:"SkillGroupName,omitempty"`
}

func (s QuerySkillGroupsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QuerySkillGroupsResponseBodyData) GoString() string {
	return s.String()
}

func (s *QuerySkillGroupsResponseBodyData) SetChannelType(v int32) *QuerySkillGroupsResponseBodyData {
	s.ChannelType = &v
	return s
}

func (s *QuerySkillGroupsResponseBodyData) SetDescription(v string) *QuerySkillGroupsResponseBodyData {
	s.Description = &v
	return s
}

func (s *QuerySkillGroupsResponseBodyData) SetDisplayName(v string) *QuerySkillGroupsResponseBodyData {
	s.DisplayName = &v
	return s
}

func (s *QuerySkillGroupsResponseBodyData) SetSkillGroupId(v int64) *QuerySkillGroupsResponseBodyData {
	s.SkillGroupId = &v
	return s
}

func (s *QuerySkillGroupsResponseBodyData) SetSkillGroupName(v string) *QuerySkillGroupsResponseBodyData {
	s.SkillGroupName = &v
	return s
}

type QuerySkillGroupsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QuerySkillGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QuerySkillGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s QuerySkillGroupsResponse) GoString() string {
	return s.String()
}

func (s *QuerySkillGroupsResponse) SetHeaders(v map[string]*string) *QuerySkillGroupsResponse {
	s.Headers = v
	return s
}

func (s *QuerySkillGroupsResponse) SetStatusCode(v int32) *QuerySkillGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *QuerySkillGroupsResponse) SetBody(v *QuerySkillGroupsResponseBody) *QuerySkillGroupsResponse {
	s.Body = v
	return s
}

type QueryTicketActionsRequest struct {
	ActionCodeList []*int32 `json:"ActionCodeList,omitempty" xml:"ActionCodeList,omitempty" type:"Repeated"`
	InstanceId     *string  `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	TicketId       *string  `json:"TicketId,omitempty" xml:"TicketId,omitempty"`
}

func (s QueryTicketActionsRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryTicketActionsRequest) GoString() string {
	return s.String()
}

func (s *QueryTicketActionsRequest) SetActionCodeList(v []*int32) *QueryTicketActionsRequest {
	s.ActionCodeList = v
	return s
}

func (s *QueryTicketActionsRequest) SetInstanceId(v string) *QueryTicketActionsRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryTicketActionsRequest) SetTicketId(v string) *QueryTicketActionsRequest {
	s.TicketId = &v
	return s
}

type QueryTicketActionsResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryTicketActionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryTicketActionsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTicketActionsResponseBody) SetCode(v string) *QueryTicketActionsResponseBody {
	s.Code = &v
	return s
}

func (s *QueryTicketActionsResponseBody) SetData(v string) *QueryTicketActionsResponseBody {
	s.Data = &v
	return s
}

func (s *QueryTicketActionsResponseBody) SetMessage(v string) *QueryTicketActionsResponseBody {
	s.Message = &v
	return s
}

func (s *QueryTicketActionsResponseBody) SetRequestId(v string) *QueryTicketActionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryTicketActionsResponseBody) SetSuccess(v bool) *QueryTicketActionsResponseBody {
	s.Success = &v
	return s
}

type QueryTicketActionsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryTicketActionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryTicketActionsResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryTicketActionsResponse) GoString() string {
	return s.String()
}

func (s *QueryTicketActionsResponse) SetHeaders(v map[string]*string) *QueryTicketActionsResponse {
	s.Headers = v
	return s
}

func (s *QueryTicketActionsResponse) SetStatusCode(v int32) *QueryTicketActionsResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryTicketActionsResponse) SetBody(v *QueryTicketActionsResponseBody) *QueryTicketActionsResponse {
	s.Body = v
	return s
}

type QueryTicketCountRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OperatorId  *int64  `json:"OperatorId,omitempty" xml:"OperatorId,omitempty"`
}

func (s QueryTicketCountRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryTicketCountRequest) GoString() string {
	return s.String()
}

func (s *QueryTicketCountRequest) SetClientToken(v string) *QueryTicketCountRequest {
	s.ClientToken = &v
	return s
}

func (s *QueryTicketCountRequest) SetInstanceId(v string) *QueryTicketCountRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryTicketCountRequest) SetOperatorId(v int64) *QueryTicketCountRequest {
	s.OperatorId = &v
	return s
}

type QueryTicketCountResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryTicketCountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryTicketCountResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTicketCountResponseBody) SetCode(v string) *QueryTicketCountResponseBody {
	s.Code = &v
	return s
}

func (s *QueryTicketCountResponseBody) SetData(v string) *QueryTicketCountResponseBody {
	s.Data = &v
	return s
}

func (s *QueryTicketCountResponseBody) SetMessage(v string) *QueryTicketCountResponseBody {
	s.Message = &v
	return s
}

func (s *QueryTicketCountResponseBody) SetRequestId(v string) *QueryTicketCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryTicketCountResponseBody) SetSuccess(v bool) *QueryTicketCountResponseBody {
	s.Success = &v
	return s
}

type QueryTicketCountResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryTicketCountResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryTicketCountResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryTicketCountResponse) GoString() string {
	return s.String()
}

func (s *QueryTicketCountResponse) SetHeaders(v map[string]*string) *QueryTicketCountResponse {
	s.Headers = v
	return s
}

func (s *QueryTicketCountResponse) SetStatusCode(v int32) *QueryTicketCountResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryTicketCountResponse) SetBody(v *QueryTicketCountResponseBody) *QueryTicketCountResponse {
	s.Body = v
	return s
}

type QueryTicketsRequest struct {
	AccountName        *string                `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	CaseId             *int64                 `json:"CaseId,omitempty" xml:"CaseId,omitempty"`
	CaseStatus         *int32                 `json:"CaseStatus,omitempty" xml:"CaseStatus,omitempty"`
	CaseType           *int32                 `json:"CaseType,omitempty" xml:"CaseType,omitempty"`
	ChannelId          *string                `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	ChannelType        *int32                 `json:"ChannelType,omitempty" xml:"ChannelType,omitempty"`
	CurrentPage        *int32                 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	DealId             *int64                 `json:"DealId,omitempty" xml:"DealId,omitempty"`
	EndCaseGmtCreate   *int64                 `json:"EndCaseGmtCreate,omitempty" xml:"EndCaseGmtCreate,omitempty"`
	Extra              map[string]interface{} `json:"Extra,omitempty" xml:"Extra,omitempty"`
	InstanceId         *string                `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MemberId           *int64                 `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	PageSize           *int32                 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ParentCaseId       *int64                 `json:"ParentCaseId,omitempty" xml:"ParentCaseId,omitempty"`
	SrType             *int64                 `json:"SrType,omitempty" xml:"SrType,omitempty"`
	StartCaseGmtCreate *int64                 `json:"StartCaseGmtCreate,omitempty" xml:"StartCaseGmtCreate,omitempty"`
	TaskStatus         *int32                 `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	TouchId            *int64                 `json:"TouchId,omitempty" xml:"TouchId,omitempty"`
}

func (s QueryTicketsRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryTicketsRequest) GoString() string {
	return s.String()
}

func (s *QueryTicketsRequest) SetAccountName(v string) *QueryTicketsRequest {
	s.AccountName = &v
	return s
}

func (s *QueryTicketsRequest) SetCaseId(v int64) *QueryTicketsRequest {
	s.CaseId = &v
	return s
}

func (s *QueryTicketsRequest) SetCaseStatus(v int32) *QueryTicketsRequest {
	s.CaseStatus = &v
	return s
}

func (s *QueryTicketsRequest) SetCaseType(v int32) *QueryTicketsRequest {
	s.CaseType = &v
	return s
}

func (s *QueryTicketsRequest) SetChannelId(v string) *QueryTicketsRequest {
	s.ChannelId = &v
	return s
}

func (s *QueryTicketsRequest) SetChannelType(v int32) *QueryTicketsRequest {
	s.ChannelType = &v
	return s
}

func (s *QueryTicketsRequest) SetCurrentPage(v int32) *QueryTicketsRequest {
	s.CurrentPage = &v
	return s
}

func (s *QueryTicketsRequest) SetDealId(v int64) *QueryTicketsRequest {
	s.DealId = &v
	return s
}

func (s *QueryTicketsRequest) SetEndCaseGmtCreate(v int64) *QueryTicketsRequest {
	s.EndCaseGmtCreate = &v
	return s
}

func (s *QueryTicketsRequest) SetExtra(v map[string]interface{}) *QueryTicketsRequest {
	s.Extra = v
	return s
}

func (s *QueryTicketsRequest) SetInstanceId(v string) *QueryTicketsRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryTicketsRequest) SetMemberId(v int64) *QueryTicketsRequest {
	s.MemberId = &v
	return s
}

func (s *QueryTicketsRequest) SetPageSize(v int32) *QueryTicketsRequest {
	s.PageSize = &v
	return s
}

func (s *QueryTicketsRequest) SetParentCaseId(v int64) *QueryTicketsRequest {
	s.ParentCaseId = &v
	return s
}

func (s *QueryTicketsRequest) SetSrType(v int64) *QueryTicketsRequest {
	s.SrType = &v
	return s
}

func (s *QueryTicketsRequest) SetStartCaseGmtCreate(v int64) *QueryTicketsRequest {
	s.StartCaseGmtCreate = &v
	return s
}

func (s *QueryTicketsRequest) SetTaskStatus(v int32) *QueryTicketsRequest {
	s.TaskStatus = &v
	return s
}

func (s *QueryTicketsRequest) SetTouchId(v int64) *QueryTicketsRequest {
	s.TouchId = &v
	return s
}

type QueryTicketsShrinkRequest struct {
	AccountName        *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	CaseId             *int64  `json:"CaseId,omitempty" xml:"CaseId,omitempty"`
	CaseStatus         *int32  `json:"CaseStatus,omitempty" xml:"CaseStatus,omitempty"`
	CaseType           *int32  `json:"CaseType,omitempty" xml:"CaseType,omitempty"`
	ChannelId          *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	ChannelType        *int32  `json:"ChannelType,omitempty" xml:"ChannelType,omitempty"`
	CurrentPage        *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	DealId             *int64  `json:"DealId,omitempty" xml:"DealId,omitempty"`
	EndCaseGmtCreate   *int64  `json:"EndCaseGmtCreate,omitempty" xml:"EndCaseGmtCreate,omitempty"`
	ExtraShrink        *string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	InstanceId         *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MemberId           *int64  `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	PageSize           *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ParentCaseId       *int64  `json:"ParentCaseId,omitempty" xml:"ParentCaseId,omitempty"`
	SrType             *int64  `json:"SrType,omitempty" xml:"SrType,omitempty"`
	StartCaseGmtCreate *int64  `json:"StartCaseGmtCreate,omitempty" xml:"StartCaseGmtCreate,omitempty"`
	TaskStatus         *int32  `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	TouchId            *int64  `json:"TouchId,omitempty" xml:"TouchId,omitempty"`
}

func (s QueryTicketsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryTicketsShrinkRequest) GoString() string {
	return s.String()
}

func (s *QueryTicketsShrinkRequest) SetAccountName(v string) *QueryTicketsShrinkRequest {
	s.AccountName = &v
	return s
}

func (s *QueryTicketsShrinkRequest) SetCaseId(v int64) *QueryTicketsShrinkRequest {
	s.CaseId = &v
	return s
}

func (s *QueryTicketsShrinkRequest) SetCaseStatus(v int32) *QueryTicketsShrinkRequest {
	s.CaseStatus = &v
	return s
}

func (s *QueryTicketsShrinkRequest) SetCaseType(v int32) *QueryTicketsShrinkRequest {
	s.CaseType = &v
	return s
}

func (s *QueryTicketsShrinkRequest) SetChannelId(v string) *QueryTicketsShrinkRequest {
	s.ChannelId = &v
	return s
}

func (s *QueryTicketsShrinkRequest) SetChannelType(v int32) *QueryTicketsShrinkRequest {
	s.ChannelType = &v
	return s
}

func (s *QueryTicketsShrinkRequest) SetCurrentPage(v int32) *QueryTicketsShrinkRequest {
	s.CurrentPage = &v
	return s
}

func (s *QueryTicketsShrinkRequest) SetDealId(v int64) *QueryTicketsShrinkRequest {
	s.DealId = &v
	return s
}

func (s *QueryTicketsShrinkRequest) SetEndCaseGmtCreate(v int64) *QueryTicketsShrinkRequest {
	s.EndCaseGmtCreate = &v
	return s
}

func (s *QueryTicketsShrinkRequest) SetExtraShrink(v string) *QueryTicketsShrinkRequest {
	s.ExtraShrink = &v
	return s
}

func (s *QueryTicketsShrinkRequest) SetInstanceId(v string) *QueryTicketsShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryTicketsShrinkRequest) SetMemberId(v int64) *QueryTicketsShrinkRequest {
	s.MemberId = &v
	return s
}

func (s *QueryTicketsShrinkRequest) SetPageSize(v int32) *QueryTicketsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *QueryTicketsShrinkRequest) SetParentCaseId(v int64) *QueryTicketsShrinkRequest {
	s.ParentCaseId = &v
	return s
}

func (s *QueryTicketsShrinkRequest) SetSrType(v int64) *QueryTicketsShrinkRequest {
	s.SrType = &v
	return s
}

func (s *QueryTicketsShrinkRequest) SetStartCaseGmtCreate(v int64) *QueryTicketsShrinkRequest {
	s.StartCaseGmtCreate = &v
	return s
}

func (s *QueryTicketsShrinkRequest) SetTaskStatus(v int32) *QueryTicketsShrinkRequest {
	s.TaskStatus = &v
	return s
}

func (s *QueryTicketsShrinkRequest) SetTouchId(v int64) *QueryTicketsShrinkRequest {
	s.TouchId = &v
	return s
}

type QueryTicketsResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryTicketsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryTicketsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTicketsResponseBody) SetCode(v string) *QueryTicketsResponseBody {
	s.Code = &v
	return s
}

func (s *QueryTicketsResponseBody) SetData(v string) *QueryTicketsResponseBody {
	s.Data = &v
	return s
}

func (s *QueryTicketsResponseBody) SetMessage(v string) *QueryTicketsResponseBody {
	s.Message = &v
	return s
}

func (s *QueryTicketsResponseBody) SetRequestId(v string) *QueryTicketsResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryTicketsResponseBody) SetSuccess(v bool) *QueryTicketsResponseBody {
	s.Success = &v
	return s
}

type QueryTicketsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *QueryTicketsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryTicketsResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryTicketsResponse) GoString() string {
	return s.String()
}

func (s *QueryTicketsResponse) SetHeaders(v map[string]*string) *QueryTicketsResponse {
	s.Headers = v
	return s
}

func (s *QueryTicketsResponse) SetStatusCode(v int32) *QueryTicketsResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryTicketsResponse) SetBody(v *QueryTicketsResponseBody) *QueryTicketsResponse {
	s.Body = v
	return s
}

type RemoveSkillGroupRequest struct {
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	SkillGroupId *string `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
}

func (s RemoveSkillGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveSkillGroupRequest) GoString() string {
	return s.String()
}

func (s *RemoveSkillGroupRequest) SetClientToken(v string) *RemoveSkillGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *RemoveSkillGroupRequest) SetInstanceId(v string) *RemoveSkillGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *RemoveSkillGroupRequest) SetSkillGroupId(v string) *RemoveSkillGroupRequest {
	s.SkillGroupId = &v
	return s
}

type RemoveSkillGroupResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RemoveSkillGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveSkillGroupResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveSkillGroupResponseBody) SetCode(v string) *RemoveSkillGroupResponseBody {
	s.Code = &v
	return s
}

func (s *RemoveSkillGroupResponseBody) SetMessage(v string) *RemoveSkillGroupResponseBody {
	s.Message = &v
	return s
}

func (s *RemoveSkillGroupResponseBody) SetRequestId(v string) *RemoveSkillGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveSkillGroupResponseBody) SetSuccess(v bool) *RemoveSkillGroupResponseBody {
	s.Success = &v
	return s
}

type RemoveSkillGroupResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RemoveSkillGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveSkillGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveSkillGroupResponse) GoString() string {
	return s.String()
}

func (s *RemoveSkillGroupResponse) SetHeaders(v map[string]*string) *RemoveSkillGroupResponse {
	s.Headers = v
	return s
}

func (s *RemoveSkillGroupResponse) SetStatusCode(v int32) *RemoveSkillGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveSkillGroupResponse) SetBody(v *RemoveSkillGroupResponseBody) *RemoveSkillGroupResponse {
	s.Body = v
	return s
}

type SearchTicketByIdRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	StatusCode  *int32  `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	TicketId    *int64  `json:"TicketId,omitempty" xml:"TicketId,omitempty"`
}

func (s SearchTicketByIdRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchTicketByIdRequest) GoString() string {
	return s.String()
}

func (s *SearchTicketByIdRequest) SetClientToken(v string) *SearchTicketByIdRequest {
	s.ClientToken = &v
	return s
}

func (s *SearchTicketByIdRequest) SetInstanceId(v string) *SearchTicketByIdRequest {
	s.InstanceId = &v
	return s
}

func (s *SearchTicketByIdRequest) SetStatusCode(v int32) *SearchTicketByIdRequest {
	s.StatusCode = &v
	return s
}

func (s *SearchTicketByIdRequest) SetTicketId(v int64) *SearchTicketByIdRequest {
	s.TicketId = &v
	return s
}

type SearchTicketByIdResponseBody struct {
	Code           *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *SearchTicketByIdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int64                            `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SearchTicketByIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchTicketByIdResponseBody) GoString() string {
	return s.String()
}

func (s *SearchTicketByIdResponseBody) SetCode(v string) *SearchTicketByIdResponseBody {
	s.Code = &v
	return s
}

func (s *SearchTicketByIdResponseBody) SetData(v *SearchTicketByIdResponseBodyData) *SearchTicketByIdResponseBody {
	s.Data = v
	return s
}

func (s *SearchTicketByIdResponseBody) SetHttpStatusCode(v int64) *SearchTicketByIdResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SearchTicketByIdResponseBody) SetMessage(v string) *SearchTicketByIdResponseBody {
	s.Message = &v
	return s
}

func (s *SearchTicketByIdResponseBody) SetRequestId(v string) *SearchTicketByIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchTicketByIdResponseBody) SetSuccess(v bool) *SearchTicketByIdResponseBody {
	s.Success = &v
	return s
}

type SearchTicketByIdResponseBodyData struct {
	Activities      []*SearchTicketByIdResponseBodyDataActivities      `json:"Activities,omitempty" xml:"Activities,omitempty" type:"Repeated"`
	ActivityRecords []*SearchTicketByIdResponseBodyDataActivityRecords `json:"ActivityRecords,omitempty" xml:"ActivityRecords,omitempty" type:"Repeated"`
	CarbonCopy      *string                                            `json:"CarbonCopy,omitempty" xml:"CarbonCopy,omitempty"`
	CaseStatus      *int32                                             `json:"CaseStatus,omitempty" xml:"CaseStatus,omitempty"`
	CategoryId      *int64                                             `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	CreateTime      *int64                                             `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CreatorId       *int64                                             `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	CreatorName     *string                                            `json:"CreatorName,omitempty" xml:"CreatorName,omitempty"`
	CreatorType     *int32                                             `json:"CreatorType,omitempty" xml:"CreatorType,omitempty"`
	FormData        *string                                            `json:"FormData,omitempty" xml:"FormData,omitempty"`
	FromInfo        *string                                            `json:"FromInfo,omitempty" xml:"FromInfo,omitempty"`
	MemberId        *int64                                             `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	MemberName      *string                                            `json:"MemberName,omitempty" xml:"MemberName,omitempty"`
	ModifiedTime    *int64                                             `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	ParentCaseId    *int64                                             `json:"ParentCaseId,omitempty" xml:"ParentCaseId,omitempty"`
	Priority        *int32                                             `json:"Priority,omitempty" xml:"Priority,omitempty"`
	ServiceId       *int64                                             `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	TemplateId      *int64                                             `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TicketId        *int64                                             `json:"TicketId,omitempty" xml:"TicketId,omitempty"`
	TicketName      *string                                            `json:"TicketName,omitempty" xml:"TicketName,omitempty"`
}

func (s SearchTicketByIdResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SearchTicketByIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *SearchTicketByIdResponseBodyData) SetActivities(v []*SearchTicketByIdResponseBodyDataActivities) *SearchTicketByIdResponseBodyData {
	s.Activities = v
	return s
}

func (s *SearchTicketByIdResponseBodyData) SetActivityRecords(v []*SearchTicketByIdResponseBodyDataActivityRecords) *SearchTicketByIdResponseBodyData {
	s.ActivityRecords = v
	return s
}

func (s *SearchTicketByIdResponseBodyData) SetCarbonCopy(v string) *SearchTicketByIdResponseBodyData {
	s.CarbonCopy = &v
	return s
}

func (s *SearchTicketByIdResponseBodyData) SetCaseStatus(v int32) *SearchTicketByIdResponseBodyData {
	s.CaseStatus = &v
	return s
}

func (s *SearchTicketByIdResponseBodyData) SetCategoryId(v int64) *SearchTicketByIdResponseBodyData {
	s.CategoryId = &v
	return s
}

func (s *SearchTicketByIdResponseBodyData) SetCreateTime(v int64) *SearchTicketByIdResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *SearchTicketByIdResponseBodyData) SetCreatorId(v int64) *SearchTicketByIdResponseBodyData {
	s.CreatorId = &v
	return s
}

func (s *SearchTicketByIdResponseBodyData) SetCreatorName(v string) *SearchTicketByIdResponseBodyData {
	s.CreatorName = &v
	return s
}

func (s *SearchTicketByIdResponseBodyData) SetCreatorType(v int32) *SearchTicketByIdResponseBodyData {
	s.CreatorType = &v
	return s
}

func (s *SearchTicketByIdResponseBodyData) SetFormData(v string) *SearchTicketByIdResponseBodyData {
	s.FormData = &v
	return s
}

func (s *SearchTicketByIdResponseBodyData) SetFromInfo(v string) *SearchTicketByIdResponseBodyData {
	s.FromInfo = &v
	return s
}

func (s *SearchTicketByIdResponseBodyData) SetMemberId(v int64) *SearchTicketByIdResponseBodyData {
	s.MemberId = &v
	return s
}

func (s *SearchTicketByIdResponseBodyData) SetMemberName(v string) *SearchTicketByIdResponseBodyData {
	s.MemberName = &v
	return s
}

func (s *SearchTicketByIdResponseBodyData) SetModifiedTime(v int64) *SearchTicketByIdResponseBodyData {
	s.ModifiedTime = &v
	return s
}

func (s *SearchTicketByIdResponseBodyData) SetParentCaseId(v int64) *SearchTicketByIdResponseBodyData {
	s.ParentCaseId = &v
	return s
}

func (s *SearchTicketByIdResponseBodyData) SetPriority(v int32) *SearchTicketByIdResponseBodyData {
	s.Priority = &v
	return s
}

func (s *SearchTicketByIdResponseBodyData) SetServiceId(v int64) *SearchTicketByIdResponseBodyData {
	s.ServiceId = &v
	return s
}

func (s *SearchTicketByIdResponseBodyData) SetTemplateId(v int64) *SearchTicketByIdResponseBodyData {
	s.TemplateId = &v
	return s
}

func (s *SearchTicketByIdResponseBodyData) SetTicketId(v int64) *SearchTicketByIdResponseBodyData {
	s.TicketId = &v
	return s
}

func (s *SearchTicketByIdResponseBodyData) SetTicketName(v string) *SearchTicketByIdResponseBodyData {
	s.TicketName = &v
	return s
}

type SearchTicketByIdResponseBodyDataActivities struct {
	ActivityCode     *string `json:"ActivityCode,omitempty" xml:"ActivityCode,omitempty"`
	ActivityFormData *string `json:"ActivityFormData,omitempty" xml:"ActivityFormData,omitempty"`
}

func (s SearchTicketByIdResponseBodyDataActivities) String() string {
	return tea.Prettify(s)
}

func (s SearchTicketByIdResponseBodyDataActivities) GoString() string {
	return s.String()
}

func (s *SearchTicketByIdResponseBodyDataActivities) SetActivityCode(v string) *SearchTicketByIdResponseBodyDataActivities {
	s.ActivityCode = &v
	return s
}

func (s *SearchTicketByIdResponseBodyDataActivities) SetActivityFormData(v string) *SearchTicketByIdResponseBodyDataActivities {
	s.ActivityFormData = &v
	return s
}

type SearchTicketByIdResponseBodyDataActivityRecords struct {
	ActionCode     *string `json:"ActionCode,omitempty" xml:"ActionCode,omitempty"`
	ActionCodeDesc *string `json:"ActionCodeDesc,omitempty" xml:"ActionCodeDesc,omitempty"`
	GmtCreate      *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	Memo           *string `json:"Memo,omitempty" xml:"Memo,omitempty"`
	Operator       *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	OperatorName   *string `json:"OperatorName,omitempty" xml:"OperatorName,omitempty"`
}

func (s SearchTicketByIdResponseBodyDataActivityRecords) String() string {
	return tea.Prettify(s)
}

func (s SearchTicketByIdResponseBodyDataActivityRecords) GoString() string {
	return s.String()
}

func (s *SearchTicketByIdResponseBodyDataActivityRecords) SetActionCode(v string) *SearchTicketByIdResponseBodyDataActivityRecords {
	s.ActionCode = &v
	return s
}

func (s *SearchTicketByIdResponseBodyDataActivityRecords) SetActionCodeDesc(v string) *SearchTicketByIdResponseBodyDataActivityRecords {
	s.ActionCodeDesc = &v
	return s
}

func (s *SearchTicketByIdResponseBodyDataActivityRecords) SetGmtCreate(v int64) *SearchTicketByIdResponseBodyDataActivityRecords {
	s.GmtCreate = &v
	return s
}

func (s *SearchTicketByIdResponseBodyDataActivityRecords) SetMemo(v string) *SearchTicketByIdResponseBodyDataActivityRecords {
	s.Memo = &v
	return s
}

func (s *SearchTicketByIdResponseBodyDataActivityRecords) SetOperator(v string) *SearchTicketByIdResponseBodyDataActivityRecords {
	s.Operator = &v
	return s
}

func (s *SearchTicketByIdResponseBodyDataActivityRecords) SetOperatorName(v string) *SearchTicketByIdResponseBodyDataActivityRecords {
	s.OperatorName = &v
	return s
}

type SearchTicketByIdResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SearchTicketByIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SearchTicketByIdResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchTicketByIdResponse) GoString() string {
	return s.String()
}

func (s *SearchTicketByIdResponse) SetHeaders(v map[string]*string) *SearchTicketByIdResponse {
	s.Headers = v
	return s
}

func (s *SearchTicketByIdResponse) SetStatusCode(v int32) *SearchTicketByIdResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchTicketByIdResponse) SetBody(v *SearchTicketByIdResponseBody) *SearchTicketByIdResponse {
	s.Body = v
	return s
}

type SearchTicketByPhoneRequest struct {
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	EndTime      *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PageNo       *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Phone        *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
	StartTime    *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TemplateId   *int64  `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TicketStatus *string `json:"TicketStatus,omitempty" xml:"TicketStatus,omitempty"`
}

func (s SearchTicketByPhoneRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchTicketByPhoneRequest) GoString() string {
	return s.String()
}

func (s *SearchTicketByPhoneRequest) SetClientToken(v string) *SearchTicketByPhoneRequest {
	s.ClientToken = &v
	return s
}

func (s *SearchTicketByPhoneRequest) SetEndTime(v int64) *SearchTicketByPhoneRequest {
	s.EndTime = &v
	return s
}

func (s *SearchTicketByPhoneRequest) SetInstanceId(v string) *SearchTicketByPhoneRequest {
	s.InstanceId = &v
	return s
}

func (s *SearchTicketByPhoneRequest) SetPageNo(v int32) *SearchTicketByPhoneRequest {
	s.PageNo = &v
	return s
}

func (s *SearchTicketByPhoneRequest) SetPageSize(v int32) *SearchTicketByPhoneRequest {
	s.PageSize = &v
	return s
}

func (s *SearchTicketByPhoneRequest) SetPhone(v string) *SearchTicketByPhoneRequest {
	s.Phone = &v
	return s
}

func (s *SearchTicketByPhoneRequest) SetStartTime(v int64) *SearchTicketByPhoneRequest {
	s.StartTime = &v
	return s
}

func (s *SearchTicketByPhoneRequest) SetTemplateId(v int64) *SearchTicketByPhoneRequest {
	s.TemplateId = &v
	return s
}

func (s *SearchTicketByPhoneRequest) SetTicketStatus(v string) *SearchTicketByPhoneRequest {
	s.TicketStatus = &v
	return s
}

type SearchTicketByPhoneResponseBody struct {
	Code         *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         []*SearchTicketByPhoneResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message      *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	OnePageSize  *int32                                 `json:"OnePageSize,omitempty" xml:"OnePageSize,omitempty"`
	PageNo       *int32                                 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	RequestId    *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalPage    *int32                                 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
	TotalResults *int32                                 `json:"TotalResults,omitempty" xml:"TotalResults,omitempty"`
}

func (s SearchTicketByPhoneResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchTicketByPhoneResponseBody) GoString() string {
	return s.String()
}

func (s *SearchTicketByPhoneResponseBody) SetCode(v string) *SearchTicketByPhoneResponseBody {
	s.Code = &v
	return s
}

func (s *SearchTicketByPhoneResponseBody) SetData(v []*SearchTicketByPhoneResponseBodyData) *SearchTicketByPhoneResponseBody {
	s.Data = v
	return s
}

func (s *SearchTicketByPhoneResponseBody) SetMessage(v string) *SearchTicketByPhoneResponseBody {
	s.Message = &v
	return s
}

func (s *SearchTicketByPhoneResponseBody) SetOnePageSize(v int32) *SearchTicketByPhoneResponseBody {
	s.OnePageSize = &v
	return s
}

func (s *SearchTicketByPhoneResponseBody) SetPageNo(v int32) *SearchTicketByPhoneResponseBody {
	s.PageNo = &v
	return s
}

func (s *SearchTicketByPhoneResponseBody) SetRequestId(v string) *SearchTicketByPhoneResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchTicketByPhoneResponseBody) SetSuccess(v bool) *SearchTicketByPhoneResponseBody {
	s.Success = &v
	return s
}

func (s *SearchTicketByPhoneResponseBody) SetTotalPage(v int32) *SearchTicketByPhoneResponseBody {
	s.TotalPage = &v
	return s
}

func (s *SearchTicketByPhoneResponseBody) SetTotalResults(v int32) *SearchTicketByPhoneResponseBody {
	s.TotalResults = &v
	return s
}

type SearchTicketByPhoneResponseBodyData struct {
	CarbonCopy   *string `json:"CarbonCopy,omitempty" xml:"CarbonCopy,omitempty"`
	CaseStatus   *int32  `json:"CaseStatus,omitempty" xml:"CaseStatus,omitempty"`
	CategoryId   *int64  `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	CreateTime   *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CreatorId    *int64  `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	CreatorName  *string `json:"CreatorName,omitempty" xml:"CreatorName,omitempty"`
	CreatorType  *int32  `json:"CreatorType,omitempty" xml:"CreatorType,omitempty"`
	FormData     *string `json:"FormData,omitempty" xml:"FormData,omitempty"`
	FromInfo     *string `json:"FromInfo,omitempty" xml:"FromInfo,omitempty"`
	MemberId     *int64  `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	MemberName   *string `json:"MemberName,omitempty" xml:"MemberName,omitempty"`
	ModifiedTime *int64  `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	Priority     *int32  `json:"Priority,omitempty" xml:"Priority,omitempty"`
	ServiceId    *int64  `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	TaskStatus   *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	TemplateId   *int64  `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TicketId     *int64  `json:"TicketId,omitempty" xml:"TicketId,omitempty"`
}

func (s SearchTicketByPhoneResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SearchTicketByPhoneResponseBodyData) GoString() string {
	return s.String()
}

func (s *SearchTicketByPhoneResponseBodyData) SetCarbonCopy(v string) *SearchTicketByPhoneResponseBodyData {
	s.CarbonCopy = &v
	return s
}

func (s *SearchTicketByPhoneResponseBodyData) SetCaseStatus(v int32) *SearchTicketByPhoneResponseBodyData {
	s.CaseStatus = &v
	return s
}

func (s *SearchTicketByPhoneResponseBodyData) SetCategoryId(v int64) *SearchTicketByPhoneResponseBodyData {
	s.CategoryId = &v
	return s
}

func (s *SearchTicketByPhoneResponseBodyData) SetCreateTime(v int64) *SearchTicketByPhoneResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *SearchTicketByPhoneResponseBodyData) SetCreatorId(v int64) *SearchTicketByPhoneResponseBodyData {
	s.CreatorId = &v
	return s
}

func (s *SearchTicketByPhoneResponseBodyData) SetCreatorName(v string) *SearchTicketByPhoneResponseBodyData {
	s.CreatorName = &v
	return s
}

func (s *SearchTicketByPhoneResponseBodyData) SetCreatorType(v int32) *SearchTicketByPhoneResponseBodyData {
	s.CreatorType = &v
	return s
}

func (s *SearchTicketByPhoneResponseBodyData) SetFormData(v string) *SearchTicketByPhoneResponseBodyData {
	s.FormData = &v
	return s
}

func (s *SearchTicketByPhoneResponseBodyData) SetFromInfo(v string) *SearchTicketByPhoneResponseBodyData {
	s.FromInfo = &v
	return s
}

func (s *SearchTicketByPhoneResponseBodyData) SetMemberId(v int64) *SearchTicketByPhoneResponseBodyData {
	s.MemberId = &v
	return s
}

func (s *SearchTicketByPhoneResponseBodyData) SetMemberName(v string) *SearchTicketByPhoneResponseBodyData {
	s.MemberName = &v
	return s
}

func (s *SearchTicketByPhoneResponseBodyData) SetModifiedTime(v int64) *SearchTicketByPhoneResponseBodyData {
	s.ModifiedTime = &v
	return s
}

func (s *SearchTicketByPhoneResponseBodyData) SetPriority(v int32) *SearchTicketByPhoneResponseBodyData {
	s.Priority = &v
	return s
}

func (s *SearchTicketByPhoneResponseBodyData) SetServiceId(v int64) *SearchTicketByPhoneResponseBodyData {
	s.ServiceId = &v
	return s
}

func (s *SearchTicketByPhoneResponseBodyData) SetTaskStatus(v string) *SearchTicketByPhoneResponseBodyData {
	s.TaskStatus = &v
	return s
}

func (s *SearchTicketByPhoneResponseBodyData) SetTemplateId(v int64) *SearchTicketByPhoneResponseBodyData {
	s.TemplateId = &v
	return s
}

func (s *SearchTicketByPhoneResponseBodyData) SetTicketId(v int64) *SearchTicketByPhoneResponseBodyData {
	s.TicketId = &v
	return s
}

type SearchTicketByPhoneResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SearchTicketByPhoneResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SearchTicketByPhoneResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchTicketByPhoneResponse) GoString() string {
	return s.String()
}

func (s *SearchTicketByPhoneResponse) SetHeaders(v map[string]*string) *SearchTicketByPhoneResponse {
	s.Headers = v
	return s
}

func (s *SearchTicketByPhoneResponse) SetStatusCode(v int32) *SearchTicketByPhoneResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchTicketByPhoneResponse) SetBody(v *SearchTicketByPhoneResponseBody) *SearchTicketByPhoneResponse {
	s.Body = v
	return s
}

type SearchTicketListRequest struct {
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	EndTime      *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OperatorId   *int64  `json:"OperatorId,omitempty" xml:"OperatorId,omitempty"`
	PageNo       *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	StartTime    *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TicketStatus *string `json:"TicketStatus,omitempty" xml:"TicketStatus,omitempty"`
}

func (s SearchTicketListRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchTicketListRequest) GoString() string {
	return s.String()
}

func (s *SearchTicketListRequest) SetClientToken(v string) *SearchTicketListRequest {
	s.ClientToken = &v
	return s
}

func (s *SearchTicketListRequest) SetEndTime(v int64) *SearchTicketListRequest {
	s.EndTime = &v
	return s
}

func (s *SearchTicketListRequest) SetInstanceId(v string) *SearchTicketListRequest {
	s.InstanceId = &v
	return s
}

func (s *SearchTicketListRequest) SetOperatorId(v int64) *SearchTicketListRequest {
	s.OperatorId = &v
	return s
}

func (s *SearchTicketListRequest) SetPageNo(v int32) *SearchTicketListRequest {
	s.PageNo = &v
	return s
}

func (s *SearchTicketListRequest) SetPageSize(v int32) *SearchTicketListRequest {
	s.PageSize = &v
	return s
}

func (s *SearchTicketListRequest) SetStartTime(v int64) *SearchTicketListRequest {
	s.StartTime = &v
	return s
}

func (s *SearchTicketListRequest) SetTicketStatus(v string) *SearchTicketListRequest {
	s.TicketStatus = &v
	return s
}

type SearchTicketListResponseBody struct {
	Code         *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data         []*SearchTicketListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message      *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	OnePageSize  *int32                              `json:"OnePageSize,omitempty" xml:"OnePageSize,omitempty"`
	PageNo       *int32                              `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	RequestId    *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalPage    *int32                              `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
	TotalResults *int32                              `json:"TotalResults,omitempty" xml:"TotalResults,omitempty"`
}

func (s SearchTicketListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchTicketListResponseBody) GoString() string {
	return s.String()
}

func (s *SearchTicketListResponseBody) SetCode(v string) *SearchTicketListResponseBody {
	s.Code = &v
	return s
}

func (s *SearchTicketListResponseBody) SetData(v []*SearchTicketListResponseBodyData) *SearchTicketListResponseBody {
	s.Data = v
	return s
}

func (s *SearchTicketListResponseBody) SetMessage(v string) *SearchTicketListResponseBody {
	s.Message = &v
	return s
}

func (s *SearchTicketListResponseBody) SetOnePageSize(v int32) *SearchTicketListResponseBody {
	s.OnePageSize = &v
	return s
}

func (s *SearchTicketListResponseBody) SetPageNo(v int32) *SearchTicketListResponseBody {
	s.PageNo = &v
	return s
}

func (s *SearchTicketListResponseBody) SetRequestId(v string) *SearchTicketListResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchTicketListResponseBody) SetSuccess(v bool) *SearchTicketListResponseBody {
	s.Success = &v
	return s
}

func (s *SearchTicketListResponseBody) SetTotalPage(v int32) *SearchTicketListResponseBody {
	s.TotalPage = &v
	return s
}

func (s *SearchTicketListResponseBody) SetTotalResults(v int32) *SearchTicketListResponseBody {
	s.TotalResults = &v
	return s
}

type SearchTicketListResponseBodyData struct {
	CarbonCopy   *string `json:"CarbonCopy,omitempty" xml:"CarbonCopy,omitempty"`
	CaseStatus   *int32  `json:"CaseStatus,omitempty" xml:"CaseStatus,omitempty"`
	CategoryId   *int64  `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	CreateTime   *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CreatorId    *int64  `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	CreatorName  *string `json:"CreatorName,omitempty" xml:"CreatorName,omitempty"`
	CreatorType  *int32  `json:"CreatorType,omitempty" xml:"CreatorType,omitempty"`
	FormData     *string `json:"FormData,omitempty" xml:"FormData,omitempty"`
	FromInfo     *string `json:"FromInfo,omitempty" xml:"FromInfo,omitempty"`
	MemberId     *int64  `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	MemberName   *string `json:"MemberName,omitempty" xml:"MemberName,omitempty"`
	ModifiedTime *int64  `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	Priority     *int32  `json:"Priority,omitempty" xml:"Priority,omitempty"`
	ServiceId    *int64  `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	TaskStatus   *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	TemplateId   *int64  `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TicketId     *int64  `json:"TicketId,omitempty" xml:"TicketId,omitempty"`
}

func (s SearchTicketListResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SearchTicketListResponseBodyData) GoString() string {
	return s.String()
}

func (s *SearchTicketListResponseBodyData) SetCarbonCopy(v string) *SearchTicketListResponseBodyData {
	s.CarbonCopy = &v
	return s
}

func (s *SearchTicketListResponseBodyData) SetCaseStatus(v int32) *SearchTicketListResponseBodyData {
	s.CaseStatus = &v
	return s
}

func (s *SearchTicketListResponseBodyData) SetCategoryId(v int64) *SearchTicketListResponseBodyData {
	s.CategoryId = &v
	return s
}

func (s *SearchTicketListResponseBodyData) SetCreateTime(v int64) *SearchTicketListResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *SearchTicketListResponseBodyData) SetCreatorId(v int64) *SearchTicketListResponseBodyData {
	s.CreatorId = &v
	return s
}

func (s *SearchTicketListResponseBodyData) SetCreatorName(v string) *SearchTicketListResponseBodyData {
	s.CreatorName = &v
	return s
}

func (s *SearchTicketListResponseBodyData) SetCreatorType(v int32) *SearchTicketListResponseBodyData {
	s.CreatorType = &v
	return s
}

func (s *SearchTicketListResponseBodyData) SetFormData(v string) *SearchTicketListResponseBodyData {
	s.FormData = &v
	return s
}

func (s *SearchTicketListResponseBodyData) SetFromInfo(v string) *SearchTicketListResponseBodyData {
	s.FromInfo = &v
	return s
}

func (s *SearchTicketListResponseBodyData) SetMemberId(v int64) *SearchTicketListResponseBodyData {
	s.MemberId = &v
	return s
}

func (s *SearchTicketListResponseBodyData) SetMemberName(v string) *SearchTicketListResponseBodyData {
	s.MemberName = &v
	return s
}

func (s *SearchTicketListResponseBodyData) SetModifiedTime(v int64) *SearchTicketListResponseBodyData {
	s.ModifiedTime = &v
	return s
}

func (s *SearchTicketListResponseBodyData) SetPriority(v int32) *SearchTicketListResponseBodyData {
	s.Priority = &v
	return s
}

func (s *SearchTicketListResponseBodyData) SetServiceId(v int64) *SearchTicketListResponseBodyData {
	s.ServiceId = &v
	return s
}

func (s *SearchTicketListResponseBodyData) SetTaskStatus(v string) *SearchTicketListResponseBodyData {
	s.TaskStatus = &v
	return s
}

func (s *SearchTicketListResponseBodyData) SetTemplateId(v int64) *SearchTicketListResponseBodyData {
	s.TemplateId = &v
	return s
}

func (s *SearchTicketListResponseBodyData) SetTicketId(v int64) *SearchTicketListResponseBodyData {
	s.TicketId = &v
	return s
}

type SearchTicketListResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SearchTicketListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SearchTicketListResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchTicketListResponse) GoString() string {
	return s.String()
}

func (s *SearchTicketListResponse) SetHeaders(v map[string]*string) *SearchTicketListResponse {
	s.Headers = v
	return s
}

func (s *SearchTicketListResponse) SetStatusCode(v int32) *SearchTicketListResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchTicketListResponse) SetBody(v *SearchTicketListResponseBody) *SearchTicketListResponse {
	s.Body = v
	return s
}

type SendHotlineHeartBeatRequest struct {
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Token       *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s SendHotlineHeartBeatRequest) String() string {
	return tea.Prettify(s)
}

func (s SendHotlineHeartBeatRequest) GoString() string {
	return s.String()
}

func (s *SendHotlineHeartBeatRequest) SetAccountName(v string) *SendHotlineHeartBeatRequest {
	s.AccountName = &v
	return s
}

func (s *SendHotlineHeartBeatRequest) SetClientToken(v string) *SendHotlineHeartBeatRequest {
	s.ClientToken = &v
	return s
}

func (s *SendHotlineHeartBeatRequest) SetInstanceId(v string) *SendHotlineHeartBeatRequest {
	s.InstanceId = &v
	return s
}

func (s *SendHotlineHeartBeatRequest) SetToken(v string) *SendHotlineHeartBeatRequest {
	s.Token = &v
	return s
}

type SendHotlineHeartBeatResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SendHotlineHeartBeatResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendHotlineHeartBeatResponseBody) GoString() string {
	return s.String()
}

func (s *SendHotlineHeartBeatResponseBody) SetCode(v string) *SendHotlineHeartBeatResponseBody {
	s.Code = &v
	return s
}

func (s *SendHotlineHeartBeatResponseBody) SetMessage(v string) *SendHotlineHeartBeatResponseBody {
	s.Message = &v
	return s
}

func (s *SendHotlineHeartBeatResponseBody) SetRequestId(v string) *SendHotlineHeartBeatResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendHotlineHeartBeatResponseBody) SetSuccess(v bool) *SendHotlineHeartBeatResponseBody {
	s.Success = &v
	return s
}

type SendHotlineHeartBeatResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SendHotlineHeartBeatResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SendHotlineHeartBeatResponse) String() string {
	return tea.Prettify(s)
}

func (s SendHotlineHeartBeatResponse) GoString() string {
	return s.String()
}

func (s *SendHotlineHeartBeatResponse) SetHeaders(v map[string]*string) *SendHotlineHeartBeatResponse {
	s.Headers = v
	return s
}

func (s *SendHotlineHeartBeatResponse) SetStatusCode(v int32) *SendHotlineHeartBeatResponse {
	s.StatusCode = &v
	return s
}

func (s *SendHotlineHeartBeatResponse) SetBody(v *SendHotlineHeartBeatResponseBody) *SendHotlineHeartBeatResponse {
	s.Body = v
	return s
}

type SendOutboundCommandRequest struct {
	AccountName   *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	CalledNumber  *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	CallingNumber *string `json:"CallingNumber,omitempty" xml:"CallingNumber,omitempty"`
	CustomerInfo  *string `json:"CustomerInfo,omitempty" xml:"CustomerInfo,omitempty"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s SendOutboundCommandRequest) String() string {
	return tea.Prettify(s)
}

func (s SendOutboundCommandRequest) GoString() string {
	return s.String()
}

func (s *SendOutboundCommandRequest) SetAccountName(v string) *SendOutboundCommandRequest {
	s.AccountName = &v
	return s
}

func (s *SendOutboundCommandRequest) SetCalledNumber(v string) *SendOutboundCommandRequest {
	s.CalledNumber = &v
	return s
}

func (s *SendOutboundCommandRequest) SetCallingNumber(v string) *SendOutboundCommandRequest {
	s.CallingNumber = &v
	return s
}

func (s *SendOutboundCommandRequest) SetCustomerInfo(v string) *SendOutboundCommandRequest {
	s.CustomerInfo = &v
	return s
}

func (s *SendOutboundCommandRequest) SetInstanceId(v string) *SendOutboundCommandRequest {
	s.InstanceId = &v
	return s
}

type SendOutboundCommandResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SendOutboundCommandResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendOutboundCommandResponseBody) GoString() string {
	return s.String()
}

func (s *SendOutboundCommandResponseBody) SetCode(v string) *SendOutboundCommandResponseBody {
	s.Code = &v
	return s
}

func (s *SendOutboundCommandResponseBody) SetData(v string) *SendOutboundCommandResponseBody {
	s.Data = &v
	return s
}

func (s *SendOutboundCommandResponseBody) SetMessage(v string) *SendOutboundCommandResponseBody {
	s.Message = &v
	return s
}

func (s *SendOutboundCommandResponseBody) SetRequestId(v string) *SendOutboundCommandResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendOutboundCommandResponseBody) SetSuccess(v bool) *SendOutboundCommandResponseBody {
	s.Success = &v
	return s
}

type SendOutboundCommandResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SendOutboundCommandResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SendOutboundCommandResponse) String() string {
	return tea.Prettify(s)
}

func (s SendOutboundCommandResponse) GoString() string {
	return s.String()
}

func (s *SendOutboundCommandResponse) SetHeaders(v map[string]*string) *SendOutboundCommandResponse {
	s.Headers = v
	return s
}

func (s *SendOutboundCommandResponse) SetStatusCode(v int32) *SendOutboundCommandResponse {
	s.StatusCode = &v
	return s
}

func (s *SendOutboundCommandResponse) SetBody(v *SendOutboundCommandResponseBody) *SendOutboundCommandResponse {
	s.Body = v
	return s
}

type SendVerificationCodeRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Phone       *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
}

func (s SendVerificationCodeRequest) String() string {
	return tea.Prettify(s)
}

func (s SendVerificationCodeRequest) GoString() string {
	return s.String()
}

func (s *SendVerificationCodeRequest) SetClientToken(v string) *SendVerificationCodeRequest {
	s.ClientToken = &v
	return s
}

func (s *SendVerificationCodeRequest) SetInstanceId(v string) *SendVerificationCodeRequest {
	s.InstanceId = &v
	return s
}

func (s *SendVerificationCodeRequest) SetPhone(v string) *SendVerificationCodeRequest {
	s.Phone = &v
	return s
}

type SendVerificationCodeResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SendVerificationCodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendVerificationCodeResponseBody) GoString() string {
	return s.String()
}

func (s *SendVerificationCodeResponseBody) SetCode(v string) *SendVerificationCodeResponseBody {
	s.Code = &v
	return s
}

func (s *SendVerificationCodeResponseBody) SetMessage(v string) *SendVerificationCodeResponseBody {
	s.Message = &v
	return s
}

func (s *SendVerificationCodeResponseBody) SetRequestId(v string) *SendVerificationCodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendVerificationCodeResponseBody) SetSuccess(v bool) *SendVerificationCodeResponseBody {
	s.Success = &v
	return s
}

type SendVerificationCodeResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SendVerificationCodeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SendVerificationCodeResponse) String() string {
	return tea.Prettify(s)
}

func (s SendVerificationCodeResponse) GoString() string {
	return s.String()
}

func (s *SendVerificationCodeResponse) SetHeaders(v map[string]*string) *SendVerificationCodeResponse {
	s.Headers = v
	return s
}

func (s *SendVerificationCodeResponse) SetStatusCode(v int32) *SendVerificationCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *SendVerificationCodeResponse) SetBody(v *SendVerificationCodeResponseBody) *SendVerificationCodeResponse {
	s.Body = v
	return s
}

type StartAgentMonitorRequest struct {
	AccountName       *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	CallId            *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	CallerParentId    *int64  `json:"CallerParentId,omitempty" xml:"CallerParentId,omitempty"`
	CallerType        *string `json:"CallerType,omitempty" xml:"CallerType,omitempty"`
	CallerUid         *int64  `json:"CallerUid,omitempty" xml:"CallerUid,omitempty"`
	ClientToken       *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ConnectionId      *string `json:"ConnectionId,omitempty" xml:"ConnectionId,omitempty"`
	InstanceId        *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	JobId             *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	RequestId         *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StsTokenCallerUid *int64  `json:"StsTokenCallerUid,omitempty" xml:"StsTokenCallerUid,omitempty"`
}

func (s StartAgentMonitorRequest) String() string {
	return tea.Prettify(s)
}

func (s StartAgentMonitorRequest) GoString() string {
	return s.String()
}

func (s *StartAgentMonitorRequest) SetAccountName(v string) *StartAgentMonitorRequest {
	s.AccountName = &v
	return s
}

func (s *StartAgentMonitorRequest) SetCallId(v string) *StartAgentMonitorRequest {
	s.CallId = &v
	return s
}

func (s *StartAgentMonitorRequest) SetCallerParentId(v int64) *StartAgentMonitorRequest {
	s.CallerParentId = &v
	return s
}

func (s *StartAgentMonitorRequest) SetCallerType(v string) *StartAgentMonitorRequest {
	s.CallerType = &v
	return s
}

func (s *StartAgentMonitorRequest) SetCallerUid(v int64) *StartAgentMonitorRequest {
	s.CallerUid = &v
	return s
}

func (s *StartAgentMonitorRequest) SetClientToken(v string) *StartAgentMonitorRequest {
	s.ClientToken = &v
	return s
}

func (s *StartAgentMonitorRequest) SetConnectionId(v string) *StartAgentMonitorRequest {
	s.ConnectionId = &v
	return s
}

func (s *StartAgentMonitorRequest) SetInstanceId(v string) *StartAgentMonitorRequest {
	s.InstanceId = &v
	return s
}

func (s *StartAgentMonitorRequest) SetJobId(v string) *StartAgentMonitorRequest {
	s.JobId = &v
	return s
}

func (s *StartAgentMonitorRequest) SetRequestId(v string) *StartAgentMonitorRequest {
	s.RequestId = &v
	return s
}

func (s *StartAgentMonitorRequest) SetStsTokenCallerUid(v int64) *StartAgentMonitorRequest {
	s.StsTokenCallerUid = &v
	return s
}

type StartAgentMonitorResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StartAgentMonitorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartAgentMonitorResponseBody) GoString() string {
	return s.String()
}

func (s *StartAgentMonitorResponseBody) SetCode(v string) *StartAgentMonitorResponseBody {
	s.Code = &v
	return s
}

func (s *StartAgentMonitorResponseBody) SetMessage(v string) *StartAgentMonitorResponseBody {
	s.Message = &v
	return s
}

func (s *StartAgentMonitorResponseBody) SetRequestId(v string) *StartAgentMonitorResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartAgentMonitorResponseBody) SetSuccess(v bool) *StartAgentMonitorResponseBody {
	s.Success = &v
	return s
}

type StartAgentMonitorResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StartAgentMonitorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartAgentMonitorResponse) String() string {
	return tea.Prettify(s)
}

func (s StartAgentMonitorResponse) GoString() string {
	return s.String()
}

func (s *StartAgentMonitorResponse) SetHeaders(v map[string]*string) *StartAgentMonitorResponse {
	s.Headers = v
	return s
}

func (s *StartAgentMonitorResponse) SetStatusCode(v int32) *StartAgentMonitorResponse {
	s.StatusCode = &v
	return s
}

func (s *StartAgentMonitorResponse) SetBody(v *StartAgentMonitorResponseBody) *StartAgentMonitorResponse {
	s.Body = v
	return s
}

type StartCallRequest struct {
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	Callee      *string `json:"Callee,omitempty" xml:"Callee,omitempty"`
	Caller      *string `json:"Caller,omitempty" xml:"Caller,omitempty"`
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s StartCallRequest) String() string {
	return tea.Prettify(s)
}

func (s StartCallRequest) GoString() string {
	return s.String()
}

func (s *StartCallRequest) SetAccountName(v string) *StartCallRequest {
	s.AccountName = &v
	return s
}

func (s *StartCallRequest) SetCallee(v string) *StartCallRequest {
	s.Callee = &v
	return s
}

func (s *StartCallRequest) SetCaller(v string) *StartCallRequest {
	s.Caller = &v
	return s
}

func (s *StartCallRequest) SetClientToken(v string) *StartCallRequest {
	s.ClientToken = &v
	return s
}

func (s *StartCallRequest) SetInstanceId(v string) *StartCallRequest {
	s.InstanceId = &v
	return s
}

type StartCallResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StartCallResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartCallResponseBody) GoString() string {
	return s.String()
}

func (s *StartCallResponseBody) SetCode(v string) *StartCallResponseBody {
	s.Code = &v
	return s
}

func (s *StartCallResponseBody) SetMessage(v string) *StartCallResponseBody {
	s.Message = &v
	return s
}

func (s *StartCallResponseBody) SetRequestId(v string) *StartCallResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartCallResponseBody) SetSuccess(v bool) *StartCallResponseBody {
	s.Success = &v
	return s
}

type StartCallResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StartCallResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartCallResponse) String() string {
	return tea.Prettify(s)
}

func (s StartCallResponse) GoString() string {
	return s.String()
}

func (s *StartCallResponse) SetHeaders(v map[string]*string) *StartCallResponse {
	s.Headers = v
	return s
}

func (s *StartCallResponse) SetStatusCode(v int32) *StartCallResponse {
	s.StatusCode = &v
	return s
}

func (s *StartCallResponse) SetBody(v *StartCallResponseBody) *StartCallResponse {
	s.Body = v
	return s
}

type StartCallV2Request struct {
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	Callee      *string `json:"Callee,omitempty" xml:"Callee,omitempty"`
	Caller      *string `json:"Caller,omitempty" xml:"Caller,omitempty"`
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	JsonMsg     *string `json:"JsonMsg,omitempty" xml:"JsonMsg,omitempty"`
}

func (s StartCallV2Request) String() string {
	return tea.Prettify(s)
}

func (s StartCallV2Request) GoString() string {
	return s.String()
}

func (s *StartCallV2Request) SetAccountName(v string) *StartCallV2Request {
	s.AccountName = &v
	return s
}

func (s *StartCallV2Request) SetCallee(v string) *StartCallV2Request {
	s.Callee = &v
	return s
}

func (s *StartCallV2Request) SetCaller(v string) *StartCallV2Request {
	s.Caller = &v
	return s
}

func (s *StartCallV2Request) SetClientToken(v string) *StartCallV2Request {
	s.ClientToken = &v
	return s
}

func (s *StartCallV2Request) SetInstanceId(v string) *StartCallV2Request {
	s.InstanceId = &v
	return s
}

func (s *StartCallV2Request) SetJsonMsg(v string) *StartCallV2Request {
	s.JsonMsg = &v
	return s
}

type StartCallV2ResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StartCallV2ResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartCallV2ResponseBody) GoString() string {
	return s.String()
}

func (s *StartCallV2ResponseBody) SetCode(v string) *StartCallV2ResponseBody {
	s.Code = &v
	return s
}

func (s *StartCallV2ResponseBody) SetMessage(v string) *StartCallV2ResponseBody {
	s.Message = &v
	return s
}

func (s *StartCallV2ResponseBody) SetRequestId(v string) *StartCallV2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartCallV2ResponseBody) SetSuccess(v bool) *StartCallV2ResponseBody {
	s.Success = &v
	return s
}

type StartCallV2Response struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StartCallV2ResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartCallV2Response) String() string {
	return tea.Prettify(s)
}

func (s StartCallV2Response) GoString() string {
	return s.String()
}

func (s *StartCallV2Response) SetHeaders(v map[string]*string) *StartCallV2Response {
	s.Headers = v
	return s
}

func (s *StartCallV2Response) SetStatusCode(v int32) *StartCallV2Response {
	s.StatusCode = &v
	return s
}

func (s *StartCallV2Response) SetBody(v *StartCallV2ResponseBody) *StartCallV2Response {
	s.Body = v
	return s
}

type StartChatWorkRequest struct {
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s StartChatWorkRequest) String() string {
	return tea.Prettify(s)
}

func (s StartChatWorkRequest) GoString() string {
	return s.String()
}

func (s *StartChatWorkRequest) SetAccountName(v string) *StartChatWorkRequest {
	s.AccountName = &v
	return s
}

func (s *StartChatWorkRequest) SetInstanceId(v string) *StartChatWorkRequest {
	s.InstanceId = &v
	return s
}

type StartChatWorkResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *string `json:"Data,omitempty" xml:"Data,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StartChatWorkResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartChatWorkResponseBody) GoString() string {
	return s.String()
}

func (s *StartChatWorkResponseBody) SetCode(v string) *StartChatWorkResponseBody {
	s.Code = &v
	return s
}

func (s *StartChatWorkResponseBody) SetData(v string) *StartChatWorkResponseBody {
	s.Data = &v
	return s
}

func (s *StartChatWorkResponseBody) SetHttpStatusCode(v int32) *StartChatWorkResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *StartChatWorkResponseBody) SetMessage(v string) *StartChatWorkResponseBody {
	s.Message = &v
	return s
}

func (s *StartChatWorkResponseBody) SetRequestId(v string) *StartChatWorkResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartChatWorkResponseBody) SetSuccess(v bool) *StartChatWorkResponseBody {
	s.Success = &v
	return s
}

type StartChatWorkResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StartChatWorkResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartChatWorkResponse) String() string {
	return tea.Prettify(s)
}

func (s StartChatWorkResponse) GoString() string {
	return s.String()
}

func (s *StartChatWorkResponse) SetHeaders(v map[string]*string) *StartChatWorkResponse {
	s.Headers = v
	return s
}

func (s *StartChatWorkResponse) SetStatusCode(v int32) *StartChatWorkResponse {
	s.StatusCode = &v
	return s
}

func (s *StartChatWorkResponse) SetBody(v *StartChatWorkResponseBody) *StartChatWorkResponse {
	s.Body = v
	return s
}

type StartHotlineServiceRequest struct {
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s StartHotlineServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s StartHotlineServiceRequest) GoString() string {
	return s.String()
}

func (s *StartHotlineServiceRequest) SetAccountName(v string) *StartHotlineServiceRequest {
	s.AccountName = &v
	return s
}

func (s *StartHotlineServiceRequest) SetClientToken(v string) *StartHotlineServiceRequest {
	s.ClientToken = &v
	return s
}

func (s *StartHotlineServiceRequest) SetInstanceId(v string) *StartHotlineServiceRequest {
	s.InstanceId = &v
	return s
}

type StartHotlineServiceResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *string `json:"Data,omitempty" xml:"Data,omitempty"`
	HttpStatusCode *int64  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StartHotlineServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartHotlineServiceResponseBody) GoString() string {
	return s.String()
}

func (s *StartHotlineServiceResponseBody) SetCode(v string) *StartHotlineServiceResponseBody {
	s.Code = &v
	return s
}

func (s *StartHotlineServiceResponseBody) SetData(v string) *StartHotlineServiceResponseBody {
	s.Data = &v
	return s
}

func (s *StartHotlineServiceResponseBody) SetHttpStatusCode(v int64) *StartHotlineServiceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *StartHotlineServiceResponseBody) SetMessage(v string) *StartHotlineServiceResponseBody {
	s.Message = &v
	return s
}

func (s *StartHotlineServiceResponseBody) SetRequestId(v string) *StartHotlineServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartHotlineServiceResponseBody) SetSuccess(v bool) *StartHotlineServiceResponseBody {
	s.Success = &v
	return s
}

type StartHotlineServiceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StartHotlineServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartHotlineServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s StartHotlineServiceResponse) GoString() string {
	return s.String()
}

func (s *StartHotlineServiceResponse) SetHeaders(v map[string]*string) *StartHotlineServiceResponse {
	s.Headers = v
	return s
}

func (s *StartHotlineServiceResponse) SetStatusCode(v int32) *StartHotlineServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *StartHotlineServiceResponse) SetBody(v *StartHotlineServiceResponseBody) *StartHotlineServiceResponse {
	s.Body = v
	return s
}

type StartHotlineWorkRequest struct {
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	SkillGroups *string `json:"SkillGroups,omitempty" xml:"SkillGroups,omitempty"`
	WorkChannel *string `json:"WorkChannel,omitempty" xml:"WorkChannel,omitempty"`
}

func (s StartHotlineWorkRequest) String() string {
	return tea.Prettify(s)
}

func (s StartHotlineWorkRequest) GoString() string {
	return s.String()
}

func (s *StartHotlineWorkRequest) SetAccountName(v string) *StartHotlineWorkRequest {
	s.AccountName = &v
	return s
}

func (s *StartHotlineWorkRequest) SetClientToken(v string) *StartHotlineWorkRequest {
	s.ClientToken = &v
	return s
}

func (s *StartHotlineWorkRequest) SetInstanceId(v string) *StartHotlineWorkRequest {
	s.InstanceId = &v
	return s
}

func (s *StartHotlineWorkRequest) SetSkillGroups(v string) *StartHotlineWorkRequest {
	s.SkillGroups = &v
	return s
}

func (s *StartHotlineWorkRequest) SetWorkChannel(v string) *StartHotlineWorkRequest {
	s.WorkChannel = &v
	return s
}

type StartHotlineWorkResponseBody struct {
	Code      *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *StartHotlineWorkResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StartHotlineWorkResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartHotlineWorkResponseBody) GoString() string {
	return s.String()
}

func (s *StartHotlineWorkResponseBody) SetCode(v string) *StartHotlineWorkResponseBody {
	s.Code = &v
	return s
}

func (s *StartHotlineWorkResponseBody) SetData(v *StartHotlineWorkResponseBodyData) *StartHotlineWorkResponseBody {
	s.Data = v
	return s
}

func (s *StartHotlineWorkResponseBody) SetMessage(v string) *StartHotlineWorkResponseBody {
	s.Message = &v
	return s
}

func (s *StartHotlineWorkResponseBody) SetRequestId(v string) *StartHotlineWorkResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartHotlineWorkResponseBody) SetSuccess(v bool) *StartHotlineWorkResponseBody {
	s.Success = &v
	return s
}

type StartHotlineWorkResponseBodyData struct {
	AccountName     *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	AgentId         *int64  `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	AgentStatusCode *string `json:"AgentStatusCode,omitempty" xml:"AgentStatusCode,omitempty"`
	AgentStatusDesc *string `json:"AgentStatusDesc,omitempty" xml:"AgentStatusDesc,omitempty"`
	AgentToken      *string `json:"AgentToken,omitempty" xml:"AgentToken,omitempty"`
	ExtAttr         *string `json:"ExtAttr,omitempty" xml:"ExtAttr,omitempty"`
	TenantId        *int64  `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	WorkChannel     *string `json:"WorkChannel,omitempty" xml:"WorkChannel,omitempty"`
}

func (s StartHotlineWorkResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s StartHotlineWorkResponseBodyData) GoString() string {
	return s.String()
}

func (s *StartHotlineWorkResponseBodyData) SetAccountName(v string) *StartHotlineWorkResponseBodyData {
	s.AccountName = &v
	return s
}

func (s *StartHotlineWorkResponseBodyData) SetAgentId(v int64) *StartHotlineWorkResponseBodyData {
	s.AgentId = &v
	return s
}

func (s *StartHotlineWorkResponseBodyData) SetAgentStatusCode(v string) *StartHotlineWorkResponseBodyData {
	s.AgentStatusCode = &v
	return s
}

func (s *StartHotlineWorkResponseBodyData) SetAgentStatusDesc(v string) *StartHotlineWorkResponseBodyData {
	s.AgentStatusDesc = &v
	return s
}

func (s *StartHotlineWorkResponseBodyData) SetAgentToken(v string) *StartHotlineWorkResponseBodyData {
	s.AgentToken = &v
	return s
}

func (s *StartHotlineWorkResponseBodyData) SetExtAttr(v string) *StartHotlineWorkResponseBodyData {
	s.ExtAttr = &v
	return s
}

func (s *StartHotlineWorkResponseBodyData) SetTenantId(v int64) *StartHotlineWorkResponseBodyData {
	s.TenantId = &v
	return s
}

func (s *StartHotlineWorkResponseBodyData) SetWorkChannel(v string) *StartHotlineWorkResponseBodyData {
	s.WorkChannel = &v
	return s
}

type StartHotlineWorkResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StartHotlineWorkResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartHotlineWorkResponse) String() string {
	return tea.Prettify(s)
}

func (s StartHotlineWorkResponse) GoString() string {
	return s.String()
}

func (s *StartHotlineWorkResponse) SetHeaders(v map[string]*string) *StartHotlineWorkResponse {
	s.Headers = v
	return s
}

func (s *StartHotlineWorkResponse) SetStatusCode(v int32) *StartHotlineWorkResponse {
	s.StatusCode = &v
	return s
}

func (s *StartHotlineWorkResponse) SetBody(v *StartHotlineWorkResponseBody) *StartHotlineWorkResponse {
	s.Body = v
	return s
}

type SuspendHotlineServiceRequest struct {
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Type        *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s SuspendHotlineServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s SuspendHotlineServiceRequest) GoString() string {
	return s.String()
}

func (s *SuspendHotlineServiceRequest) SetAccountName(v string) *SuspendHotlineServiceRequest {
	s.AccountName = &v
	return s
}

func (s *SuspendHotlineServiceRequest) SetClientToken(v string) *SuspendHotlineServiceRequest {
	s.ClientToken = &v
	return s
}

func (s *SuspendHotlineServiceRequest) SetInstanceId(v string) *SuspendHotlineServiceRequest {
	s.InstanceId = &v
	return s
}

func (s *SuspendHotlineServiceRequest) SetType(v int32) *SuspendHotlineServiceRequest {
	s.Type = &v
	return s
}

type SuspendHotlineServiceResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SuspendHotlineServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SuspendHotlineServiceResponseBody) GoString() string {
	return s.String()
}

func (s *SuspendHotlineServiceResponseBody) SetCode(v string) *SuspendHotlineServiceResponseBody {
	s.Code = &v
	return s
}

func (s *SuspendHotlineServiceResponseBody) SetMessage(v string) *SuspendHotlineServiceResponseBody {
	s.Message = &v
	return s
}

func (s *SuspendHotlineServiceResponseBody) SetRequestId(v string) *SuspendHotlineServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *SuspendHotlineServiceResponseBody) SetSuccess(v bool) *SuspendHotlineServiceResponseBody {
	s.Success = &v
	return s
}

type SuspendHotlineServiceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SuspendHotlineServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SuspendHotlineServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s SuspendHotlineServiceResponse) GoString() string {
	return s.String()
}

func (s *SuspendHotlineServiceResponse) SetHeaders(v map[string]*string) *SuspendHotlineServiceResponse {
	s.Headers = v
	return s
}

func (s *SuspendHotlineServiceResponse) SetStatusCode(v int32) *SuspendHotlineServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *SuspendHotlineServiceResponse) SetBody(v *SuspendHotlineServiceResponseBody) *SuspendHotlineServiceResponse {
	s.Body = v
	return s
}

type TransferCallToAgentRequest struct {
	AccountName       *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	CallId            *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	ClientToken       *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ConnectionId      *string `json:"ConnectionId,omitempty" xml:"ConnectionId,omitempty"`
	HoldConnectionId  *string `json:"HoldConnectionId,omitempty" xml:"HoldConnectionId,omitempty"`
	InstanceId        *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	IsSingleTransfer  *string `json:"IsSingleTransfer,omitempty" xml:"IsSingleTransfer,omitempty"`
	JobId             *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	TargetAccountName *string `json:"TargetAccountName,omitempty" xml:"TargetAccountName,omitempty"`
	Type              *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s TransferCallToAgentRequest) String() string {
	return tea.Prettify(s)
}

func (s TransferCallToAgentRequest) GoString() string {
	return s.String()
}

func (s *TransferCallToAgentRequest) SetAccountName(v string) *TransferCallToAgentRequest {
	s.AccountName = &v
	return s
}

func (s *TransferCallToAgentRequest) SetCallId(v string) *TransferCallToAgentRequest {
	s.CallId = &v
	return s
}

func (s *TransferCallToAgentRequest) SetClientToken(v string) *TransferCallToAgentRequest {
	s.ClientToken = &v
	return s
}

func (s *TransferCallToAgentRequest) SetConnectionId(v string) *TransferCallToAgentRequest {
	s.ConnectionId = &v
	return s
}

func (s *TransferCallToAgentRequest) SetHoldConnectionId(v string) *TransferCallToAgentRequest {
	s.HoldConnectionId = &v
	return s
}

func (s *TransferCallToAgentRequest) SetInstanceId(v string) *TransferCallToAgentRequest {
	s.InstanceId = &v
	return s
}

func (s *TransferCallToAgentRequest) SetIsSingleTransfer(v string) *TransferCallToAgentRequest {
	s.IsSingleTransfer = &v
	return s
}

func (s *TransferCallToAgentRequest) SetJobId(v string) *TransferCallToAgentRequest {
	s.JobId = &v
	return s
}

func (s *TransferCallToAgentRequest) SetTargetAccountName(v string) *TransferCallToAgentRequest {
	s.TargetAccountName = &v
	return s
}

func (s *TransferCallToAgentRequest) SetType(v int32) *TransferCallToAgentRequest {
	s.Type = &v
	return s
}

type TransferCallToAgentResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s TransferCallToAgentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TransferCallToAgentResponseBody) GoString() string {
	return s.String()
}

func (s *TransferCallToAgentResponseBody) SetCode(v string) *TransferCallToAgentResponseBody {
	s.Code = &v
	return s
}

func (s *TransferCallToAgentResponseBody) SetMessage(v string) *TransferCallToAgentResponseBody {
	s.Message = &v
	return s
}

func (s *TransferCallToAgentResponseBody) SetRequestId(v string) *TransferCallToAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *TransferCallToAgentResponseBody) SetSuccess(v bool) *TransferCallToAgentResponseBody {
	s.Success = &v
	return s
}

type TransferCallToAgentResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *TransferCallToAgentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TransferCallToAgentResponse) String() string {
	return tea.Prettify(s)
}

func (s TransferCallToAgentResponse) GoString() string {
	return s.String()
}

func (s *TransferCallToAgentResponse) SetHeaders(v map[string]*string) *TransferCallToAgentResponse {
	s.Headers = v
	return s
}

func (s *TransferCallToAgentResponse) SetStatusCode(v int32) *TransferCallToAgentResponse {
	s.StatusCode = &v
	return s
}

func (s *TransferCallToAgentResponse) SetBody(v *TransferCallToAgentResponseBody) *TransferCallToAgentResponse {
	s.Body = v
	return s
}

type TransferCallToPhoneRequest struct {
	AccountName      *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	CallId           *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	Callee           *string `json:"Callee,omitempty" xml:"Callee,omitempty"`
	Caller           *string `json:"Caller,omitempty" xml:"Caller,omitempty"`
	ClientToken      *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ConnectionId     *string `json:"ConnectionId,omitempty" xml:"ConnectionId,omitempty"`
	HoldConnectionId *string `json:"HoldConnectionId,omitempty" xml:"HoldConnectionId,omitempty"`
	InstanceId       *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	IsSingleTransfer *bool   `json:"IsSingleTransfer,omitempty" xml:"IsSingleTransfer,omitempty"`
	JobId            *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	Type             *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
	CalleePhone      *string `json:"calleePhone,omitempty" xml:"calleePhone,omitempty"`
	CallerPhone      *string `json:"callerPhone,omitempty" xml:"callerPhone,omitempty"`
}

func (s TransferCallToPhoneRequest) String() string {
	return tea.Prettify(s)
}

func (s TransferCallToPhoneRequest) GoString() string {
	return s.String()
}

func (s *TransferCallToPhoneRequest) SetAccountName(v string) *TransferCallToPhoneRequest {
	s.AccountName = &v
	return s
}

func (s *TransferCallToPhoneRequest) SetCallId(v string) *TransferCallToPhoneRequest {
	s.CallId = &v
	return s
}

func (s *TransferCallToPhoneRequest) SetCallee(v string) *TransferCallToPhoneRequest {
	s.Callee = &v
	return s
}

func (s *TransferCallToPhoneRequest) SetCaller(v string) *TransferCallToPhoneRequest {
	s.Caller = &v
	return s
}

func (s *TransferCallToPhoneRequest) SetClientToken(v string) *TransferCallToPhoneRequest {
	s.ClientToken = &v
	return s
}

func (s *TransferCallToPhoneRequest) SetConnectionId(v string) *TransferCallToPhoneRequest {
	s.ConnectionId = &v
	return s
}

func (s *TransferCallToPhoneRequest) SetHoldConnectionId(v string) *TransferCallToPhoneRequest {
	s.HoldConnectionId = &v
	return s
}

func (s *TransferCallToPhoneRequest) SetInstanceId(v string) *TransferCallToPhoneRequest {
	s.InstanceId = &v
	return s
}

func (s *TransferCallToPhoneRequest) SetIsSingleTransfer(v bool) *TransferCallToPhoneRequest {
	s.IsSingleTransfer = &v
	return s
}

func (s *TransferCallToPhoneRequest) SetJobId(v string) *TransferCallToPhoneRequest {
	s.JobId = &v
	return s
}

func (s *TransferCallToPhoneRequest) SetType(v int32) *TransferCallToPhoneRequest {
	s.Type = &v
	return s
}

func (s *TransferCallToPhoneRequest) SetCalleePhone(v string) *TransferCallToPhoneRequest {
	s.CalleePhone = &v
	return s
}

func (s *TransferCallToPhoneRequest) SetCallerPhone(v string) *TransferCallToPhoneRequest {
	s.CallerPhone = &v
	return s
}

type TransferCallToPhoneResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int64  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s TransferCallToPhoneResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TransferCallToPhoneResponseBody) GoString() string {
	return s.String()
}

func (s *TransferCallToPhoneResponseBody) SetCode(v string) *TransferCallToPhoneResponseBody {
	s.Code = &v
	return s
}

func (s *TransferCallToPhoneResponseBody) SetHttpStatusCode(v int64) *TransferCallToPhoneResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *TransferCallToPhoneResponseBody) SetMessage(v string) *TransferCallToPhoneResponseBody {
	s.Message = &v
	return s
}

func (s *TransferCallToPhoneResponseBody) SetRequestId(v string) *TransferCallToPhoneResponseBody {
	s.RequestId = &v
	return s
}

func (s *TransferCallToPhoneResponseBody) SetSuccess(v bool) *TransferCallToPhoneResponseBody {
	s.Success = &v
	return s
}

type TransferCallToPhoneResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *TransferCallToPhoneResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TransferCallToPhoneResponse) String() string {
	return tea.Prettify(s)
}

func (s TransferCallToPhoneResponse) GoString() string {
	return s.String()
}

func (s *TransferCallToPhoneResponse) SetHeaders(v map[string]*string) *TransferCallToPhoneResponse {
	s.Headers = v
	return s
}

func (s *TransferCallToPhoneResponse) SetStatusCode(v int32) *TransferCallToPhoneResponse {
	s.StatusCode = &v
	return s
}

func (s *TransferCallToPhoneResponse) SetBody(v *TransferCallToPhoneResponseBody) *TransferCallToPhoneResponse {
	s.Body = v
	return s
}

type TransferCallToSkillGroupRequest struct {
	AccountName      *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	CallId           *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	ClientToken      *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ConnectionId     *string `json:"ConnectionId,omitempty" xml:"ConnectionId,omitempty"`
	HoldConnectionId *string `json:"HoldConnectionId,omitempty" xml:"HoldConnectionId,omitempty"`
	InstanceId       *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	IsSingleTransfer *bool   `json:"IsSingleTransfer,omitempty" xml:"IsSingleTransfer,omitempty"`
	JobId            *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	SkillGroupId     *int64  `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
	Type             *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s TransferCallToSkillGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s TransferCallToSkillGroupRequest) GoString() string {
	return s.String()
}

func (s *TransferCallToSkillGroupRequest) SetAccountName(v string) *TransferCallToSkillGroupRequest {
	s.AccountName = &v
	return s
}

func (s *TransferCallToSkillGroupRequest) SetCallId(v string) *TransferCallToSkillGroupRequest {
	s.CallId = &v
	return s
}

func (s *TransferCallToSkillGroupRequest) SetClientToken(v string) *TransferCallToSkillGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *TransferCallToSkillGroupRequest) SetConnectionId(v string) *TransferCallToSkillGroupRequest {
	s.ConnectionId = &v
	return s
}

func (s *TransferCallToSkillGroupRequest) SetHoldConnectionId(v string) *TransferCallToSkillGroupRequest {
	s.HoldConnectionId = &v
	return s
}

func (s *TransferCallToSkillGroupRequest) SetInstanceId(v string) *TransferCallToSkillGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *TransferCallToSkillGroupRequest) SetIsSingleTransfer(v bool) *TransferCallToSkillGroupRequest {
	s.IsSingleTransfer = &v
	return s
}

func (s *TransferCallToSkillGroupRequest) SetJobId(v string) *TransferCallToSkillGroupRequest {
	s.JobId = &v
	return s
}

func (s *TransferCallToSkillGroupRequest) SetSkillGroupId(v int64) *TransferCallToSkillGroupRequest {
	s.SkillGroupId = &v
	return s
}

func (s *TransferCallToSkillGroupRequest) SetType(v int32) *TransferCallToSkillGroupRequest {
	s.Type = &v
	return s
}

type TransferCallToSkillGroupResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s TransferCallToSkillGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TransferCallToSkillGroupResponseBody) GoString() string {
	return s.String()
}

func (s *TransferCallToSkillGroupResponseBody) SetCode(v string) *TransferCallToSkillGroupResponseBody {
	s.Code = &v
	return s
}

func (s *TransferCallToSkillGroupResponseBody) SetMessage(v string) *TransferCallToSkillGroupResponseBody {
	s.Message = &v
	return s
}

func (s *TransferCallToSkillGroupResponseBody) SetRequestId(v string) *TransferCallToSkillGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *TransferCallToSkillGroupResponseBody) SetSuccess(v bool) *TransferCallToSkillGroupResponseBody {
	s.Success = &v
	return s
}

type TransferCallToSkillGroupResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *TransferCallToSkillGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TransferCallToSkillGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s TransferCallToSkillGroupResponse) GoString() string {
	return s.String()
}

func (s *TransferCallToSkillGroupResponse) SetHeaders(v map[string]*string) *TransferCallToSkillGroupResponse {
	s.Headers = v
	return s
}

func (s *TransferCallToSkillGroupResponse) SetStatusCode(v int32) *TransferCallToSkillGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *TransferCallToSkillGroupResponse) SetBody(v *TransferCallToSkillGroupResponseBody) *TransferCallToSkillGroupResponse {
	s.Body = v
	return s
}

type TransferToThirdCallRequest struct {
	AccountName      *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	CallId           *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	ClientToken      *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ConnectionId     *string `json:"ConnectionId,omitempty" xml:"ConnectionId,omitempty"`
	HoldConnectionId *string `json:"HoldConnectionId,omitempty" xml:"HoldConnectionId,omitempty"`
	InstanceId       *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	JobId            *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s TransferToThirdCallRequest) String() string {
	return tea.Prettify(s)
}

func (s TransferToThirdCallRequest) GoString() string {
	return s.String()
}

func (s *TransferToThirdCallRequest) SetAccountName(v string) *TransferToThirdCallRequest {
	s.AccountName = &v
	return s
}

func (s *TransferToThirdCallRequest) SetCallId(v string) *TransferToThirdCallRequest {
	s.CallId = &v
	return s
}

func (s *TransferToThirdCallRequest) SetClientToken(v string) *TransferToThirdCallRequest {
	s.ClientToken = &v
	return s
}

func (s *TransferToThirdCallRequest) SetConnectionId(v string) *TransferToThirdCallRequest {
	s.ConnectionId = &v
	return s
}

func (s *TransferToThirdCallRequest) SetHoldConnectionId(v string) *TransferToThirdCallRequest {
	s.HoldConnectionId = &v
	return s
}

func (s *TransferToThirdCallRequest) SetInstanceId(v string) *TransferToThirdCallRequest {
	s.InstanceId = &v
	return s
}

func (s *TransferToThirdCallRequest) SetJobId(v string) *TransferToThirdCallRequest {
	s.JobId = &v
	return s
}

type TransferToThirdCallResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s TransferToThirdCallResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TransferToThirdCallResponseBody) GoString() string {
	return s.String()
}

func (s *TransferToThirdCallResponseBody) SetCode(v string) *TransferToThirdCallResponseBody {
	s.Code = &v
	return s
}

func (s *TransferToThirdCallResponseBody) SetMessage(v string) *TransferToThirdCallResponseBody {
	s.Message = &v
	return s
}

func (s *TransferToThirdCallResponseBody) SetRequestId(v string) *TransferToThirdCallResponseBody {
	s.RequestId = &v
	return s
}

func (s *TransferToThirdCallResponseBody) SetSuccess(v bool) *TransferToThirdCallResponseBody {
	s.Success = &v
	return s
}

type TransferToThirdCallResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *TransferToThirdCallResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TransferToThirdCallResponse) String() string {
	return tea.Prettify(s)
}

func (s TransferToThirdCallResponse) GoString() string {
	return s.String()
}

func (s *TransferToThirdCallResponse) SetHeaders(v map[string]*string) *TransferToThirdCallResponse {
	s.Headers = v
	return s
}

func (s *TransferToThirdCallResponse) SetStatusCode(v int32) *TransferToThirdCallResponse {
	s.StatusCode = &v
	return s
}

func (s *TransferToThirdCallResponse) SetBody(v *TransferToThirdCallResponseBody) *TransferToThirdCallResponse {
	s.Body = v
	return s
}

type UnbindAgentHotlinePhoneRequest struct {
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s UnbindAgentHotlinePhoneRequest) String() string {
	return tea.Prettify(s)
}

func (s UnbindAgentHotlinePhoneRequest) GoString() string {
	return s.String()
}

func (s *UnbindAgentHotlinePhoneRequest) SetAccountName(v string) *UnbindAgentHotlinePhoneRequest {
	s.AccountName = &v
	return s
}

func (s *UnbindAgentHotlinePhoneRequest) SetClientToken(v string) *UnbindAgentHotlinePhoneRequest {
	s.ClientToken = &v
	return s
}

func (s *UnbindAgentHotlinePhoneRequest) SetInstanceId(v string) *UnbindAgentHotlinePhoneRequest {
	s.InstanceId = &v
	return s
}

type UnbindAgentHotlinePhoneResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UnbindAgentHotlinePhoneResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnbindAgentHotlinePhoneResponseBody) GoString() string {
	return s.String()
}

func (s *UnbindAgentHotlinePhoneResponseBody) SetCode(v string) *UnbindAgentHotlinePhoneResponseBody {
	s.Code = &v
	return s
}

func (s *UnbindAgentHotlinePhoneResponseBody) SetMessage(v string) *UnbindAgentHotlinePhoneResponseBody {
	s.Message = &v
	return s
}

func (s *UnbindAgentHotlinePhoneResponseBody) SetRequestId(v string) *UnbindAgentHotlinePhoneResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnbindAgentHotlinePhoneResponseBody) SetSuccess(v bool) *UnbindAgentHotlinePhoneResponseBody {
	s.Success = &v
	return s
}

type UnbindAgentHotlinePhoneResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UnbindAgentHotlinePhoneResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UnbindAgentHotlinePhoneResponse) String() string {
	return tea.Prettify(s)
}

func (s UnbindAgentHotlinePhoneResponse) GoString() string {
	return s.String()
}

func (s *UnbindAgentHotlinePhoneResponse) SetHeaders(v map[string]*string) *UnbindAgentHotlinePhoneResponse {
	s.Headers = v
	return s
}

func (s *UnbindAgentHotlinePhoneResponse) SetStatusCode(v int32) *UnbindAgentHotlinePhoneResponse {
	s.StatusCode = &v
	return s
}

func (s *UnbindAgentHotlinePhoneResponse) SetBody(v *UnbindAgentHotlinePhoneResponseBody) *UnbindAgentHotlinePhoneResponse {
	s.Body = v
	return s
}

type UpdateAgentRequest struct {
	AccountName      *string  `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	ClientToken      *string  `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DisplayName      *string  `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	InstanceId       *string  `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	SkillGroupId     []*int64 `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty" type:"Repeated"`
	SkillGroupIdList []*int64 `json:"SkillGroupIdList,omitempty" xml:"SkillGroupIdList,omitempty" type:"Repeated"`
}

func (s UpdateAgentRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAgentRequest) GoString() string {
	return s.String()
}

func (s *UpdateAgentRequest) SetAccountName(v string) *UpdateAgentRequest {
	s.AccountName = &v
	return s
}

func (s *UpdateAgentRequest) SetClientToken(v string) *UpdateAgentRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateAgentRequest) SetDisplayName(v string) *UpdateAgentRequest {
	s.DisplayName = &v
	return s
}

func (s *UpdateAgentRequest) SetInstanceId(v string) *UpdateAgentRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateAgentRequest) SetSkillGroupId(v []*int64) *UpdateAgentRequest {
	s.SkillGroupId = v
	return s
}

func (s *UpdateAgentRequest) SetSkillGroupIdList(v []*int64) *UpdateAgentRequest {
	s.SkillGroupIdList = v
	return s
}

type UpdateAgentResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int64  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateAgentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAgentResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAgentResponseBody) SetCode(v string) *UpdateAgentResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateAgentResponseBody) SetHttpStatusCode(v int64) *UpdateAgentResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateAgentResponseBody) SetMessage(v string) *UpdateAgentResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateAgentResponseBody) SetRequestId(v string) *UpdateAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAgentResponseBody) SetSuccess(v bool) *UpdateAgentResponseBody {
	s.Success = &v
	return s
}

type UpdateAgentResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateAgentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateAgentResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAgentResponse) GoString() string {
	return s.String()
}

func (s *UpdateAgentResponse) SetHeaders(v map[string]*string) *UpdateAgentResponse {
	s.Headers = v
	return s
}

func (s *UpdateAgentResponse) SetStatusCode(v int32) *UpdateAgentResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAgentResponse) SetBody(v *UpdateAgentResponseBody) *UpdateAgentResponse {
	s.Body = v
	return s
}

type UpdateCustomerRequest struct {
	BizType     *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	Contacter   *string `json:"Contacter,omitempty" xml:"Contacter,omitempty"`
	CustomerId  *int64  `json:"CustomerId,omitempty" xml:"CustomerId,omitempty"`
	Dingding    *string `json:"Dingding,omitempty" xml:"Dingding,omitempty"`
	Email       *string `json:"Email,omitempty" xml:"Email,omitempty"`
	Industry    *string `json:"Industry,omitempty" xml:"Industry,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ManagerName *string `json:"ManagerName,omitempty" xml:"ManagerName,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OuterId     *string `json:"OuterId,omitempty" xml:"OuterId,omitempty"`
	OuterIdType *int32  `json:"OuterIdType,omitempty" xml:"OuterIdType,omitempty"`
	Phone       *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
	Position    *string `json:"Position,omitempty" xml:"Position,omitempty"`
	ProdLineId  *int64  `json:"ProdLineId,omitempty" xml:"ProdLineId,omitempty"`
	TypeCode    *string `json:"TypeCode,omitempty" xml:"TypeCode,omitempty"`
}

func (s UpdateCustomerRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateCustomerRequest) GoString() string {
	return s.String()
}

func (s *UpdateCustomerRequest) SetBizType(v string) *UpdateCustomerRequest {
	s.BizType = &v
	return s
}

func (s *UpdateCustomerRequest) SetContacter(v string) *UpdateCustomerRequest {
	s.Contacter = &v
	return s
}

func (s *UpdateCustomerRequest) SetCustomerId(v int64) *UpdateCustomerRequest {
	s.CustomerId = &v
	return s
}

func (s *UpdateCustomerRequest) SetDingding(v string) *UpdateCustomerRequest {
	s.Dingding = &v
	return s
}

func (s *UpdateCustomerRequest) SetEmail(v string) *UpdateCustomerRequest {
	s.Email = &v
	return s
}

func (s *UpdateCustomerRequest) SetIndustry(v string) *UpdateCustomerRequest {
	s.Industry = &v
	return s
}

func (s *UpdateCustomerRequest) SetInstanceId(v string) *UpdateCustomerRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateCustomerRequest) SetManagerName(v string) *UpdateCustomerRequest {
	s.ManagerName = &v
	return s
}

func (s *UpdateCustomerRequest) SetName(v string) *UpdateCustomerRequest {
	s.Name = &v
	return s
}

func (s *UpdateCustomerRequest) SetOuterId(v string) *UpdateCustomerRequest {
	s.OuterId = &v
	return s
}

func (s *UpdateCustomerRequest) SetOuterIdType(v int32) *UpdateCustomerRequest {
	s.OuterIdType = &v
	return s
}

func (s *UpdateCustomerRequest) SetPhone(v string) *UpdateCustomerRequest {
	s.Phone = &v
	return s
}

func (s *UpdateCustomerRequest) SetPosition(v string) *UpdateCustomerRequest {
	s.Position = &v
	return s
}

func (s *UpdateCustomerRequest) SetProdLineId(v int64) *UpdateCustomerRequest {
	s.ProdLineId = &v
	return s
}

func (s *UpdateCustomerRequest) SetTypeCode(v string) *UpdateCustomerRequest {
	s.TypeCode = &v
	return s
}

type UpdateCustomerResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *int64  `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateCustomerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateCustomerResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCustomerResponseBody) SetCode(v string) *UpdateCustomerResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateCustomerResponseBody) SetData(v int64) *UpdateCustomerResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateCustomerResponseBody) SetMessage(v string) *UpdateCustomerResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateCustomerResponseBody) SetRequestId(v string) *UpdateCustomerResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCustomerResponseBody) SetSuccess(v bool) *UpdateCustomerResponseBody {
	s.Success = &v
	return s
}

type UpdateCustomerResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateCustomerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateCustomerResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateCustomerResponse) GoString() string {
	return s.String()
}

func (s *UpdateCustomerResponse) SetHeaders(v map[string]*string) *UpdateCustomerResponse {
	s.Headers = v
	return s
}

func (s *UpdateCustomerResponse) SetStatusCode(v int32) *UpdateCustomerResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCustomerResponse) SetBody(v *UpdateCustomerResponseBody) *UpdateCustomerResponse {
	s.Body = v
	return s
}

type UpdateEntityTagRelationRequest struct {
	EntityTagParam *string `json:"EntityTagParam,omitempty" xml:"EntityTagParam,omitempty"`
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s UpdateEntityTagRelationRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateEntityTagRelationRequest) GoString() string {
	return s.String()
}

func (s *UpdateEntityTagRelationRequest) SetEntityTagParam(v string) *UpdateEntityTagRelationRequest {
	s.EntityTagParam = &v
	return s
}

func (s *UpdateEntityTagRelationRequest) SetInstanceId(v string) *UpdateEntityTagRelationRequest {
	s.InstanceId = &v
	return s
}

type UpdateEntityTagRelationResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateEntityTagRelationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateEntityTagRelationResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateEntityTagRelationResponseBody) SetCode(v string) *UpdateEntityTagRelationResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateEntityTagRelationResponseBody) SetData(v string) *UpdateEntityTagRelationResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateEntityTagRelationResponseBody) SetMessage(v string) *UpdateEntityTagRelationResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateEntityTagRelationResponseBody) SetRequestId(v string) *UpdateEntityTagRelationResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateEntityTagRelationResponseBody) SetSuccess(v bool) *UpdateEntityTagRelationResponseBody {
	s.Success = &v
	return s
}

type UpdateEntityTagRelationResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateEntityTagRelationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateEntityTagRelationResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateEntityTagRelationResponse) GoString() string {
	return s.String()
}

func (s *UpdateEntityTagRelationResponse) SetHeaders(v map[string]*string) *UpdateEntityTagRelationResponse {
	s.Headers = v
	return s
}

func (s *UpdateEntityTagRelationResponse) SetStatusCode(v int32) *UpdateEntityTagRelationResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateEntityTagRelationResponse) SetBody(v *UpdateEntityTagRelationResponseBody) *UpdateEntityTagRelationResponse {
	s.Body = v
	return s
}

type UpdateRingStatusRequest struct {
	CallOutStatus *string `json:"CallOutStatus,omitempty" xml:"CallOutStatus,omitempty"`
	Extra         *string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	UniqueBizId   *string `json:"UniqueBizId,omitempty" xml:"UniqueBizId,omitempty"`
}

func (s UpdateRingStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateRingStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateRingStatusRequest) SetCallOutStatus(v string) *UpdateRingStatusRequest {
	s.CallOutStatus = &v
	return s
}

func (s *UpdateRingStatusRequest) SetExtra(v string) *UpdateRingStatusRequest {
	s.Extra = &v
	return s
}

func (s *UpdateRingStatusRequest) SetInstanceId(v string) *UpdateRingStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateRingStatusRequest) SetUniqueBizId(v string) *UpdateRingStatusRequest {
	s.UniqueBizId = &v
	return s
}

type UpdateRingStatusResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *string `json:"Data,omitempty" xml:"Data,omitempty"`
	HttpStatusCode *int64  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateRingStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateRingStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRingStatusResponseBody) SetCode(v string) *UpdateRingStatusResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateRingStatusResponseBody) SetData(v string) *UpdateRingStatusResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateRingStatusResponseBody) SetHttpStatusCode(v int64) *UpdateRingStatusResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateRingStatusResponseBody) SetMessage(v string) *UpdateRingStatusResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateRingStatusResponseBody) SetRequestId(v string) *UpdateRingStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateRingStatusResponseBody) SetSuccess(v bool) *UpdateRingStatusResponseBody {
	s.Success = &v
	return s
}

type UpdateRingStatusResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateRingStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateRingStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateRingStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateRingStatusResponse) SetHeaders(v map[string]*string) *UpdateRingStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateRingStatusResponse) SetStatusCode(v int32) *UpdateRingStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRingStatusResponse) SetBody(v *UpdateRingStatusResponseBody) *UpdateRingStatusResponse {
	s.Body = v
	return s
}

type UpdateRoleRequest struct {
	ClientToken  *string  `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId   *string  `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Operator     *string  `json:"Operator,omitempty" xml:"Operator,omitempty"`
	PermissionId []*int64 `json:"PermissionId,omitempty" xml:"PermissionId,omitempty" type:"Repeated"`
	RoleId       *int64   `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
	RoleName     *string  `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
}

func (s UpdateRoleRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateRoleRequest) GoString() string {
	return s.String()
}

func (s *UpdateRoleRequest) SetClientToken(v string) *UpdateRoleRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateRoleRequest) SetInstanceId(v string) *UpdateRoleRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateRoleRequest) SetOperator(v string) *UpdateRoleRequest {
	s.Operator = &v
	return s
}

func (s *UpdateRoleRequest) SetPermissionId(v []*int64) *UpdateRoleRequest {
	s.PermissionId = v
	return s
}

func (s *UpdateRoleRequest) SetRoleId(v int64) *UpdateRoleRequest {
	s.RoleId = &v
	return s
}

func (s *UpdateRoleRequest) SetRoleName(v string) *UpdateRoleRequest {
	s.RoleName = &v
	return s
}

type UpdateRoleResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int64  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateRoleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateRoleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRoleResponseBody) SetCode(v string) *UpdateRoleResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateRoleResponseBody) SetHttpStatusCode(v int64) *UpdateRoleResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateRoleResponseBody) SetMessage(v string) *UpdateRoleResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateRoleResponseBody) SetRequestId(v string) *UpdateRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateRoleResponseBody) SetSuccess(v bool) *UpdateRoleResponseBody {
	s.Success = &v
	return s
}

type UpdateRoleResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateRoleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateRoleResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateRoleResponse) GoString() string {
	return s.String()
}

func (s *UpdateRoleResponse) SetHeaders(v map[string]*string) *UpdateRoleResponse {
	s.Headers = v
	return s
}

func (s *UpdateRoleResponse) SetStatusCode(v int32) *UpdateRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRoleResponse) SetBody(v *UpdateRoleResponseBody) *UpdateRoleResponse {
	s.Body = v
	return s
}

type UpdateSkillGroupRequest struct {
	ChannelType    *int64  `json:"ChannelType,omitempty" xml:"ChannelType,omitempty"`
	ClientToken    *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DisplayName    *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	SkillGroupId   *int64  `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
	SkillGroupName *string `json:"SkillGroupName,omitempty" xml:"SkillGroupName,omitempty"`
}

func (s UpdateSkillGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateSkillGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateSkillGroupRequest) SetChannelType(v int64) *UpdateSkillGroupRequest {
	s.ChannelType = &v
	return s
}

func (s *UpdateSkillGroupRequest) SetClientToken(v string) *UpdateSkillGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateSkillGroupRequest) SetDescription(v string) *UpdateSkillGroupRequest {
	s.Description = &v
	return s
}

func (s *UpdateSkillGroupRequest) SetDisplayName(v string) *UpdateSkillGroupRequest {
	s.DisplayName = &v
	return s
}

func (s *UpdateSkillGroupRequest) SetInstanceId(v string) *UpdateSkillGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateSkillGroupRequest) SetSkillGroupId(v int64) *UpdateSkillGroupRequest {
	s.SkillGroupId = &v
	return s
}

func (s *UpdateSkillGroupRequest) SetSkillGroupName(v string) *UpdateSkillGroupRequest {
	s.SkillGroupName = &v
	return s
}

type UpdateSkillGroupResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int64  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateSkillGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateSkillGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSkillGroupResponseBody) SetCode(v string) *UpdateSkillGroupResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateSkillGroupResponseBody) SetHttpStatusCode(v int64) *UpdateSkillGroupResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateSkillGroupResponseBody) SetMessage(v string) *UpdateSkillGroupResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateSkillGroupResponseBody) SetRequestId(v string) *UpdateSkillGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSkillGroupResponseBody) SetSuccess(v bool) *UpdateSkillGroupResponseBody {
	s.Success = &v
	return s
}

type UpdateSkillGroupResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateSkillGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateSkillGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateSkillGroupResponse) GoString() string {
	return s.String()
}

func (s *UpdateSkillGroupResponse) SetHeaders(v map[string]*string) *UpdateSkillGroupResponse {
	s.Headers = v
	return s
}

func (s *UpdateSkillGroupResponse) SetStatusCode(v int32) *UpdateSkillGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSkillGroupResponse) SetBody(v *UpdateSkillGroupResponseBody) *UpdateSkillGroupResponse {
	s.Body = v
	return s
}

type UpdateTicketRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	FormData    *string `json:"FormData,omitempty" xml:"FormData,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OperatorId  *int64  `json:"OperatorId,omitempty" xml:"OperatorId,omitempty"`
	TicketId    *int64  `json:"TicketId,omitempty" xml:"TicketId,omitempty"`
}

func (s UpdateTicketRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateTicketRequest) GoString() string {
	return s.String()
}

func (s *UpdateTicketRequest) SetClientToken(v string) *UpdateTicketRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateTicketRequest) SetFormData(v string) *UpdateTicketRequest {
	s.FormData = &v
	return s
}

func (s *UpdateTicketRequest) SetInstanceId(v string) *UpdateTicketRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateTicketRequest) SetOperatorId(v int64) *UpdateTicketRequest {
	s.OperatorId = &v
	return s
}

func (s *UpdateTicketRequest) SetTicketId(v int64) *UpdateTicketRequest {
	s.TicketId = &v
	return s
}

type UpdateTicketResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int64  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateTicketResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateTicketResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTicketResponseBody) SetCode(v string) *UpdateTicketResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateTicketResponseBody) SetHttpStatusCode(v int64) *UpdateTicketResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateTicketResponseBody) SetMessage(v string) *UpdateTicketResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateTicketResponseBody) SetRequestId(v string) *UpdateTicketResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTicketResponseBody) SetSuccess(v bool) *UpdateTicketResponseBody {
	s.Success = &v
	return s
}

type UpdateTicketResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateTicketResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateTicketResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateTicketResponse) GoString() string {
	return s.String()
}

func (s *UpdateTicketResponse) SetHeaders(v map[string]*string) *UpdateTicketResponse {
	s.Headers = v
	return s
}

func (s *UpdateTicketResponse) SetStatusCode(v int32) *UpdateTicketResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTicketResponse) SetBody(v *UpdateTicketResponseBody) *UpdateTicketResponse {
	s.Body = v
	return s
}

type CreateTicketWithBizDataRequest struct {
	BizData     *string `json:"BizData,omitempty" xml:"BizData,omitempty"`
	CarbonCopy  *string `json:"CarbonCopy,omitempty" xml:"CarbonCopy,omitempty"`
	CategoryId  *int64  `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	CreatorId   *int64  `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	CreatorName *string `json:"CreatorName,omitempty" xml:"CreatorName,omitempty"`
	CreatorType *int32  `json:"CreatorType,omitempty" xml:"CreatorType,omitempty"`
	FormData    *string `json:"FormData,omitempty" xml:"FormData,omitempty"`
	FromInfo    *string `json:"FromInfo,omitempty" xml:"FromInfo,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MemberId    *int64  `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	MemberName  *string `json:"MemberName,omitempty" xml:"MemberName,omitempty"`
	Priority    *int32  `json:"Priority,omitempty" xml:"Priority,omitempty"`
	TemplateId  *int64  `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s CreateTicketWithBizDataRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTicketWithBizDataRequest) GoString() string {
	return s.String()
}

func (s *CreateTicketWithBizDataRequest) SetBizData(v string) *CreateTicketWithBizDataRequest {
	s.BizData = &v
	return s
}

func (s *CreateTicketWithBizDataRequest) SetCarbonCopy(v string) *CreateTicketWithBizDataRequest {
	s.CarbonCopy = &v
	return s
}

func (s *CreateTicketWithBizDataRequest) SetCategoryId(v int64) *CreateTicketWithBizDataRequest {
	s.CategoryId = &v
	return s
}

func (s *CreateTicketWithBizDataRequest) SetClientToken(v string) *CreateTicketWithBizDataRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateTicketWithBizDataRequest) SetCreatorId(v int64) *CreateTicketWithBizDataRequest {
	s.CreatorId = &v
	return s
}

func (s *CreateTicketWithBizDataRequest) SetCreatorName(v string) *CreateTicketWithBizDataRequest {
	s.CreatorName = &v
	return s
}

func (s *CreateTicketWithBizDataRequest) SetCreatorType(v int32) *CreateTicketWithBizDataRequest {
	s.CreatorType = &v
	return s
}

func (s *CreateTicketWithBizDataRequest) SetFormData(v string) *CreateTicketWithBizDataRequest {
	s.FormData = &v
	return s
}

func (s *CreateTicketWithBizDataRequest) SetFromInfo(v string) *CreateTicketWithBizDataRequest {
	s.FromInfo = &v
	return s
}

func (s *CreateTicketWithBizDataRequest) SetInstanceId(v string) *CreateTicketWithBizDataRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateTicketWithBizDataRequest) SetMemberId(v int64) *CreateTicketWithBizDataRequest {
	s.MemberId = &v
	return s
}

func (s *CreateTicketWithBizDataRequest) SetMemberName(v string) *CreateTicketWithBizDataRequest {
	s.MemberName = &v
	return s
}

func (s *CreateTicketWithBizDataRequest) SetPriority(v int32) *CreateTicketWithBizDataRequest {
	s.Priority = &v
	return s
}

func (s *CreateTicketWithBizDataRequest) SetTemplateId(v int64) *CreateTicketWithBizDataRequest {
	s.TemplateId = &v
	return s
}

type CreateTicketWithBizDataResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *int64  `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateTicketWithBizDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTicketWithBizDataResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTicketWithBizDataResponseBody) SetCode(v string) *CreateTicketWithBizDataResponseBody {
	s.Code = &v
	return s
}

func (s *CreateTicketWithBizDataResponseBody) SetData(v int64) *CreateTicketWithBizDataResponseBody {
	s.Data = &v
	return s
}

func (s *CreateTicketWithBizDataResponseBody) SetMessage(v string) *CreateTicketWithBizDataResponseBody {
	s.Message = &v
	return s
}

func (s *CreateTicketWithBizDataResponseBody) SetRequestId(v string) *CreateTicketWithBizDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTicketWithBizDataResponseBody) SetSuccess(v bool) *CreateTicketWithBizDataResponseBody {
	s.Success = &v
	return s
}

type CreateTicketWithBizDataResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateTicketWithBizDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateTicketWithBizDataResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTicketWithBizDataResponse) GoString() string {
	return s.String()
}

func (s *CreateTicketWithBizDataResponse) SetHeaders(v map[string]*string) *CreateTicketWithBizDataResponse {
	s.Headers = v
	return s
}

func (s *CreateTicketWithBizDataResponse) SetStatusCode(v int32) *CreateTicketWithBizDataResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTicketWithBizDataResponse) SetBody(v *CreateTicketWithBizDataResponseBody) *CreateTicketWithBizDataResponse {
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
	client.EndpointMap = map[string]*string{
		"cn-shanghai": tea.String("scsp-vpc.cn-shanghai.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("scsp"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) AnswerCallWithOptions(request *AnswerCallRequest, runtime *util.RuntimeOptions) (_result *AnswerCallResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountName)) {
		body["AccountName"] = request.AccountName
	}

	if !tea.BoolValue(util.IsUnset(request.CallId)) {
		body["CallId"] = request.CallId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ConnectionId)) {
		body["ConnectionId"] = request.ConnectionId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		body["JobId"] = request.JobId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AnswerCall"),
		Version:     tea.String("2020-07-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AnswerCallResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AnswerCall(request *AnswerCallRequest) (_result *AnswerCallResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AnswerCallResponse{}
	_body, _err := client.AnswerCallWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AnwserAgentMonitorWithOptions(request *AnwserAgentMonitorRequest, runtime *util.RuntimeOptions) (_result *AnwserAgentMonitorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AnwserAgentMonitor"),
		Version:     tea.String("2020-07-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AnwserAgentMonitorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AnwserAgentMonitor(request *AnwserAgentMonitorRequest) (_result *AnwserAgentMonitorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AnwserAgentMonitorResponse{}
	_body, _err := client.AnwserAgentMonitorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AppMessagePushWithOptions(request *AppMessagePushRequest, runtime *util.RuntimeOptions) (_result *AppMessagePushResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ExpirationTime)) {
		query["ExpirationTime"] = request.ExpirationTime
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AppMessagePush"),
		Version:     tea.String("2020-07-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AppMessagePushResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AppMessagePush(request *AppMessagePushRequest) (_result *AppMessagePushResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AppMessagePushResponse{}
	_body, _err := client.AppMessagePushWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AssignTicketWithOptions(request *AssignTicketRequest, runtime *util.RuntimeOptions) (_result *AssignTicketResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcceptorId)) {
		body["AcceptorId"] = request.AcceptorId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		body["OperatorId"] = request.OperatorId
	}

	if !tea.BoolValue(util.IsUnset(request.TicketId)) {
		body["TicketId"] = request.TicketId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AssignTicket"),
		Version:     tea.String("2020-07-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AssignTicketResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AssignTicket(request *AssignTicketRequest) (_result *AssignTicketResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AssignTicketResponse{}
	_body, _err := client.AssignTicketWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BindAgentHotlinePhoneWithOptions(request *BindAgentHotlinePhoneRequest, runtime *util.RuntimeOptions) (_result *BindAgentHotlinePhoneResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountName)) {
		body["AccountName"] = request.AccountName
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Phone)) {
		body["Phone"] = request.Phone
	}

	if !tea.BoolValue(util.IsUnset(request.VerifyCode)) {
		body["VerifyCode"] = request.VerifyCode
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("BindAgentHotlinePhone"),
		Version:     tea.String("2020-07-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BindAgentHotlinePhoneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BindAgentHotlinePhone(request *BindAgentHotlinePhoneRequest) (_result *BindAgentHotlinePhoneResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BindAgentHotlinePhoneResponse{}
	_body, _err := client.BindAgentHotlinePhoneWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ChangeChatAgentStatusWithOptions(request *ChangeChatAgentStatusRequest, runtime *util.RuntimeOptions) (_result *ChangeChatAgentStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountName)) {
		body["AccountName"] = request.AccountName
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Method)) {
		body["Method"] = request.Method
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ChangeChatAgentStatus"),
		Version:     tea.String("2020-07-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ChangeChatAgentStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ChangeChatAgentStatus(request *ChangeChatAgentStatusRequest) (_result *ChangeChatAgentStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ChangeChatAgentStatusResponse{}
	_body, _err := client.ChangeChatAgentStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CloseTicketWithOptions(request *CloseTicketRequest, runtime *util.RuntimeOptions) (_result *CloseTicketResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ActionItems)) {
		body["ActionItems"] = request.ActionItems
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		body["OperatorId"] = request.OperatorId
	}

	if !tea.BoolValue(util.IsUnset(request.TicketId)) {
		body["TicketId"] = request.TicketId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CloseTicket"),
		Version:     tea.String("2020-07-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CloseTicketResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CloseTicket(request *CloseTicketRequest) (_result *CloseTicketResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CloseTicketResponse{}
	_body, _err := client.CloseTicketWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateAgentWithOptions(request *CreateAgentRequest, runtime *util.RuntimeOptions) (_result *CreateAgentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountName)) {
		body["AccountName"] = request.AccountName
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DisplayName)) {
		body["DisplayName"] = request.DisplayName
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	bodyFlat := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SkillGroupId)) {
		bodyFlat["SkillGroupId"] = request.SkillGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SkillGroupIdList)) {
		bodyFlat["SkillGroupIdList"] = request.SkillGroupIdList
	}

	body = tea.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAgent"),
		Version:     tea.String("2020-07-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAgentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateAgent(request *CreateAgentRequest) (_result *CreateAgentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAgentResponse{}
	_body, _err := client.CreateAgentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateCustomerWithOptions(request *CreateCustomerRequest, runtime *util.RuntimeOptions) (_result *CreateCustomerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		query["BizType"] = request.BizType
	}

	if !tea.BoolValue(util.IsUnset(request.Contacter)) {
		query["Contacter"] = request.Contacter
	}

	if !tea.BoolValue(util.IsUnset(request.Dingding)) {
		query["Dingding"] = request.Dingding
	}

	if !tea.BoolValue(util.IsUnset(request.Email)) {
		query["Email"] = request.Email
	}

	if !tea.BoolValue(util.IsUnset(request.Industry)) {
		query["Industry"] = request.Industry
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ManagerName)) {
		query["ManagerName"] = request.ManagerName
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OuterId)) {
		query["OuterId"] = request.OuterId
	}

	if !tea.BoolValue(util.IsUnset(request.OuterIdType)) {
		query["OuterIdType"] = request.OuterIdType
	}

	if !tea.BoolValue(util.IsUnset(request.Phone)) {
		query["Phone"] = request.Phone
	}

	if !tea.BoolValue(util.IsUnset(request.Position)) {
		query["Position"] = request.Position
	}

	if !tea.BoolValue(util.IsUnset(request.ProdLineId)) {
		query["ProdLineId"] = request.ProdLineId
	}

	if !tea.BoolValue(util.IsUnset(request.TypeCode)) {
		query["TypeCode"] = request.TypeCode
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateCustomer"),
		Version:     tea.String("2020-07-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateCustomerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateCustomer(request *CreateCustomerRequest) (_result *CreateCustomerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateCustomerResponse{}
	_body, _err := client.CreateCustomerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateEntityIvrRouteWithOptions(request *CreateEntityIvrRouteRequest, runtime *util.RuntimeOptions) (_result *CreateEntityIvrRouteResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DepartmentId)) {
		body["DepartmentId"] = request.DepartmentId
	}

	if !tea.BoolValue(util.IsUnset(request.EntityBizCode)) {
		body["EntityBizCode"] = request.EntityBizCode
	}

	if !tea.BoolValue(util.IsUnset(request.EntityBizCodeType)) {
		body["EntityBizCodeType"] = request.EntityBizCodeType
	}

	if !tea.BoolValue(util.IsUnset(request.EntityId)) {
		body["EntityId"] = request.EntityId
	}

	if !tea.BoolValue(util.IsUnset(request.EntityName)) {
		body["EntityName"] = request.EntityName
	}

	if !tea.BoolValue(util.IsUnset(request.EntityRelationNumber)) {
		body["EntityRelationNumber"] = request.EntityRelationNumber
	}

	if !tea.BoolValue(util.IsUnset(request.ExtInfo)) {
		body["ExtInfo"] = request.ExtInfo
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		body["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceId)) {
		body["ServiceId"] = request.ServiceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateEntityIvrRoute"),
		Version:     tea.String("2020-07-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateEntityIvrRouteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateEntityIvrRoute(request *CreateEntityIvrRouteRequest) (_result *CreateEntityIvrRouteResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateEntityIvrRouteResponse{}
	_body, _err := client.CreateEntityIvrRouteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateOuterCallCenterDataWithOptions(request *CreateOuterCallCenterDataRequest, runtime *util.RuntimeOptions) (_result *CreateOuterCallCenterDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizId)) {
		body["BizId"] = request.BizId
	}

	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		body["BizType"] = request.BizType
	}

	if !tea.BoolValue(util.IsUnset(request.CallType)) {
		body["CallType"] = request.CallType
	}

	if !tea.BoolValue(util.IsUnset(request.EndReason)) {
		body["EndReason"] = request.EndReason
	}

	if !tea.BoolValue(util.IsUnset(request.ExtInfo)) {
		body["ExtInfo"] = request.ExtInfo
	}

	if !tea.BoolValue(util.IsUnset(request.FromPhoneNum)) {
		body["FromPhoneNum"] = request.FromPhoneNum
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.InterveneTime)) {
		body["InterveneTime"] = request.InterveneTime
	}

	if !tea.BoolValue(util.IsUnset(request.RecordUrl)) {
		body["RecordUrl"] = request.RecordUrl
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		body["SessionId"] = request.SessionId
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["TenantId"] = request.TenantId
	}

	if !tea.BoolValue(util.IsUnset(request.ToPhoneNum)) {
		body["ToPhoneNum"] = request.ToPhoneNum
	}

	if !tea.BoolValue(util.IsUnset(request.UserInfo)) {
		body["UserInfo"] = request.UserInfo
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateOuterCallCenterData"),
		Version:     tea.String("2020-07-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateOuterCallCenterDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateOuterCallCenterData(request *CreateOuterCallCenterDataRequest) (_result *CreateOuterCallCenterDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateOuterCallCenterDataResponse{}
	_body, _err := client.CreateOuterCallCenterDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateRoleWithOptions(request *CreateRoleRequest, runtime *util.RuntimeOptions) (_result *CreateRoleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Operator)) {
		body["Operator"] = request.Operator
	}

	if !tea.BoolValue(util.IsUnset(request.PermissionId)) {
		body["PermissionId"] = request.PermissionId
	}

	if !tea.BoolValue(util.IsUnset(request.RoleName)) {
		body["RoleName"] = request.RoleName
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateRole"),
		Version:     tea.String("2020-07-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateRoleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateRole(request *CreateRoleRequest) (_result *CreateRoleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateRoleResponse{}
	_body, _err := client.CreateRoleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateSkillGroupWithOptions(request *CreateSkillGroupRequest, runtime *util.RuntimeOptions) (_result *CreateSkillGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChannelType)) {
		body["ChannelType"] = request.ChannelType
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DisplayName)) {
		body["DisplayName"] = request.DisplayName
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.SkillGroupName)) {
		body["SkillGroupName"] = request.SkillGroupName
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateSkillGroup"),
		Version:     tea.String("2020-07-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateSkillGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSkillGroup(request *CreateSkillGroupRequest) (_result *CreateSkillGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateSkillGroupResponse{}
	_body, _err := client.CreateSkillGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateSubTicketWithOptions(request *CreateSubTicketRequest, runtime *util.RuntimeOptions) (_result *CreateSubTicketResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentId)) {
		query["AgentId"] = request.AgentId
	}

	if !tea.BoolValue(util.IsUnset(request.BizData)) {
		query["BizData"] = request.BizData
	}

	if !tea.BoolValue(util.IsUnset(request.CreatorId)) {
		query["CreatorId"] = request.CreatorId
	}

	if !tea.BoolValue(util.IsUnset(request.CreatorName)) {
		query["CreatorName"] = request.CreatorName
	}

	if !tea.BoolValue(util.IsUnset(request.FormData)) {
		query["FormData"] = request.FormData
	}

	if !tea.BoolValue(util.IsUnset(request.FromInfo)) {
		query["FromInfo"] = request.FromInfo
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.MemberId)) {
		query["MemberId"] = request.MemberId
	}

	if !tea.BoolValue(util.IsUnset(request.MemberName)) {
		query["MemberName"] = request.MemberName
	}

	if !tea.BoolValue(util.IsUnset(request.ParentCaseId)) {
		query["ParentCaseId"] = request.ParentCaseId
	}

	if !tea.BoolValue(util.IsUnset(request.Priority)) {
		query["Priority"] = request.Priority
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateSubTicket"),
		Version:     tea.String("2020-07-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateSubTicketResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSubTicket(request *CreateSubTicketRequest) (_result *CreateSubTicketResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateSubTicketResponse{}
	_body, _err := client.CreateSubTicketWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateThirdSsoAgentWithOptions(request *CreateThirdSsoAgentRequest, runtime *util.RuntimeOptions) (_result *CreateThirdSsoAgentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountId)) {
		body["AccountId"] = request.AccountId
	}

	if !tea.BoolValue(util.IsUnset(request.AccountName)) {
		body["AccountName"] = request.AccountName
	}

	if !tea.BoolValue(util.IsUnset(request.ClientId)) {
		body["ClientId"] = request.ClientId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DisplayName)) {
		body["DisplayName"] = request.DisplayName
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	bodyFlat := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RoleIds)) {
		bodyFlat["RoleIds"] = request.RoleIds
	}

	if !tea.BoolValue(util.IsUnset(request.SkillGroupIds)) {
		bodyFlat["SkillGroupIds"] = request.SkillGroupIds
	}

	body = tea.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateThirdSsoAgent"),
		Version:     tea.String("2020-07-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateThirdSsoAgentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateThirdSsoAgent(request *CreateThirdSsoAgentRequest) (_result *CreateThirdSsoAgentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateThirdSsoAgentResponse{}
	_body, _err := client.CreateThirdSsoAgentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateTicketWithOptions(request *CreateTicketRequest, runtime *util.RuntimeOptions) (_result *CreateTicketResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CarbonCopy)) {
		body["CarbonCopy"] = request.CarbonCopy
	}

	if !tea.BoolValue(util.IsUnset(request.CategoryId)) {
		body["CategoryId"] = request.CategoryId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.CreatorId)) {
		body["CreatorId"] = request.CreatorId
	}

	if !tea.BoolValue(util.IsUnset(request.CreatorName)) {
		body["CreatorName"] = request.CreatorName
	}

	if !tea.BoolValue(util.IsUnset(request.CreatorType)) {
		body["CreatorType"] = request.CreatorType
	}

	if !tea.BoolValue(util.IsUnset(request.FormData)) {
		body["FormData"] = request.FormData
	}

	if !tea.BoolValue(util.IsUnset(request.FromInfo)) {
		body["FromInfo"] = request.FromInfo
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.MemberId)) {
		body["MemberId"] = request.MemberId
	}

	if !tea.BoolValue(util.IsUnset(request.MemberName)) {
		body["MemberName"] = request.MemberName
	}

	if !tea.BoolValue(util.IsUnset(request.Priority)) {
		body["Priority"] = request.Priority
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		body["TemplateId"] = request.TemplateId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateTicket"),
		Version:     tea.String("2020-07-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateTicketResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateTicket(request *CreateTicketRequest) (_result *CreateTicketResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateTicketResponse{}
	_body, _err := client.CreateTicketWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteAgentWithOptions(request *DeleteAgentRequest, runtime *util.RuntimeOptions) (_result *DeleteAgentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountName)) {
		query["AccountName"] = request.AccountName
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAgent"),
		Version:     tea.String("2020-07-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteAgentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteAgent(request *DeleteAgentRequest) (_result *DeleteAgentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAgentResponse{}
	_body, _err := client.DeleteAgentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteEntityRouteWithOptions(request *DeleteEntityRouteRequest, runtime *util.RuntimeOptions) (_result *DeleteEntityRouteResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.UniqueId)) {
		body["UniqueId"] = request.UniqueId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteEntityRoute"),
		Version:     tea.String("2020-07-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteEntityRouteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteEntityRoute(request *DeleteEntityRouteRequest) (_result *DeleteEntityRouteResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteEntityRouteResponse{}
	_body, _err := client.DeleteEntityRouteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DisableRoleWithOptions(request *DisableRoleRequest, runtime *util.RuntimeOptions) (_result *DisableRoleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RoleId)) {
		body["RoleId"] = request.RoleId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DisableRole"),
		Version:     tea.String("2020-07-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DisableRoleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DisableRole(request *DisableRoleRequest) (_result *DisableRoleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisableRoleResponse{}
	_body, _err := client.DisableRoleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EditEntityRouteWithOptions(request *EditEntityRouteRequest, runtime *util.RuntimeOptions) (_result *EditEntityRouteResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DepartmentId)) {
		body["DepartmentId"] = request.DepartmentId
	}

	if !tea.BoolValue(util.IsUnset(request.EntityBizCode)) {
		body["EntityBizCode"] = request.EntityBizCode
	}

	if !tea.BoolValue(util.IsUnset(request.EntityBizCodeType)) {
		body["EntityBizCodeType"] = request.EntityBizCodeType
	}

	if !tea.BoolValue(util.IsUnset(request.EntityId)) {
		body["EntityId"] = request.EntityId
	}

	if !tea.BoolValue(util.IsUnset(request.EntityName)) {
		body["EntityName"] = request.EntityName
	}

	if !tea.BoolValue(util.IsUnset(request.EntityRelationNumber)) {
		body["EntityRelationNumber"] = request.EntityRelationNumber
	}

	if !tea.BoolValue(util.IsUnset(request.ExtInfo)) {
		body["ExtInfo"] = request.ExtInfo
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		body["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceId)) {
		body["ServiceId"] = request.ServiceId
	}

	if !tea.BoolValue(util.IsUnset(request.UniqueId)) {
		body["UniqueId"] = request.UniqueId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("EditEntityRoute"),
		Version:     tea.String("2020-07-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EditEntityRouteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EditEntityRoute(request *EditEntityRouteRequest) (_result *EditEntityRouteResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EditEntityRouteResponse{}
	_body, _err := client.EditEntityRouteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EnableRoleWithOptions(request *EnableRoleRequest, runtime *util.RuntimeOptions) (_result *EnableRoleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RoleId)) {
		body["RoleId"] = request.RoleId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("EnableRole"),
		Version:     tea.String("2020-07-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnableRoleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EnableRole(request *EnableRoleRequest) (_result *EnableRoleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableRoleResponse{}
	_body, _err := client.EnableRoleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ExecuteActivityWithOptions(request *ExecuteActivityRequest, runtime *util.RuntimeOptions) (_result *ExecuteActivityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ActivityCode)) {
		body["ActivityCode"] = request.ActivityCode
	}

	if !tea.BoolValue(util.IsUnset(request.ActivityForm)) {
		body["ActivityForm"] = request.ActivityForm
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		body["OperatorId"] = request.OperatorId
	}

	if !tea.BoolValue(util.IsUnset(request.TicketId)) {
		body["TicketId"] = request.TicketId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ExecuteActivity"),
		Version:     tea.String("2020-07-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ExecuteActivityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ExecuteActivity(request *ExecuteActivityRequest) (_result *ExecuteActivityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ExecuteActivityResponse{}
	_body, _err := client.ExecuteActivityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) FetchCallWithOptions(request *FetchCallRequest, runtime *util.RuntimeOptions) (_result *FetchCallResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountName)) {
		body["AccountName"] = request.AccountName
	}

	if !tea.BoolValue(util.IsUnset(request.CallId)) {
		body["CallId"] = request.CallId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ConnectionId)) {
		body["ConnectionId"] = request.ConnectionId
	}

	if !tea.BoolValue(util.IsUnset(request.HoldConnectionId)) {
		body["HoldConnectionId"] = request.HoldConnectionId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		body["JobId"] = request.JobId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("FetchCall"),
		Version:     tea.String("2020-07-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &FetchCallResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) FetchCall(request *FetchCallRequest) (_result *FetchCallResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &FetchCallResponse{}
	_body, _err := client.FetchCallWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) FinishHotlineServiceWithOptions(request *FinishHotlineServiceRequest, runtime *util.RuntimeOptions) (_result *FinishHotlineServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountName)) {
		body["AccountName"] = request.AccountName
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("FinishHotlineService"),
		Version:     tea.String("2020-07-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &FinishHotlineServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) FinishHotlineService(request *FinishHotlineServiceRequest) (_result *FinishHotlineServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &FinishHotlineServiceResponse{}
	_body, _err := client.FinishHotlineServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) FinishHotlineServiceNewWithOptions(request *FinishHotlineServiceNewRequest, runtime *util.RuntimeOptions) (_result *FinishHotlineServiceNewResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountName)) {
		body["AccountName"] = request.AccountName
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("FinishHotlineServiceNew"),
		Version:     tea.String("2020-07-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &FinishHotlineServiceNewResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) FinishHotlineServiceNew(request *FinishHotlineServiceNewRequest) (_result *FinishHotlineServiceNewResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &FinishHotlineServiceNewResponse{}
	_body, _err := client.FinishHotlineServiceNewWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GenerateWebSocketSignWithOptions(request *GenerateWebSocketSignRequest, runtime *util.RuntimeOptions) (_result *GenerateWebSocketSignResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountName)) {
		body["AccountName"] = request.AccountName
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GenerateWebSocketSign"),
		Version:     tea.String("2020-07-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GenerateWebSocketSignResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GenerateWebSocketSign(request *GenerateWebSocketSignRequest) (_result *GenerateWebSocketSignResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GenerateWebSocketSignResponse{}
	_body, _err := client.GenerateWebSocketSignWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAgentWithOptions(request *GetAgentRequest, runtime *util.RuntimeOptions) (_result *GetAgentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAgent"),
		Version:     tea.String("2020-07-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAgentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAgent(request *GetAgentRequest) (_result *GetAgentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAgentResponse{}
	_body, _err := client.GetAgentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAgentHotlineWithOptions(request *GetAgentHotlineRequest, runtime *util.RuntimeOptions) (_result *GetAgentHotlineResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountName)) {
		body["AccountName"] = request.AccountName
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAgentHotline"),
		Version:     tea.String("2020-07-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAgentHotlineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAgentHotline(request *GetAgentHotlineRequest) (_result *GetAgentHotlineResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAgentHotlineResponse{}
	_body, _err := client.GetAgentHotlineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAgentHotlinePhoneWithOptions(request *GetAgentHotlinePhoneRequest, runtime *util.RuntimeOptions) (_result *GetAgentHotlinePhoneResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountName)) {
		body["AccountName"] = request.AccountName
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAgentHotlinePhone"),
		Version:     tea.String("2020-07-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAgentHotlinePhoneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAgentHotlinePhone(request *GetAgentHotlinePhoneRequest) (_result *GetAgentHotlinePhoneResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAgentHotlinePhoneResponse{}
	_body, _err := client.GetAgentHotlinePhoneWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAgentWorkStatusWithOptions(request *GetAgentWorkStatusRequest, runtime *util.RuntimeOptions) (_result *GetAgentWorkStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountName)) {
		body["AccountName"] = request.AccountName
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAgentWorkStatus"),
		Version:     tea.String("2020-07-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAgentWorkStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAgentWorkStatus(request *GetAgentWorkStatusRequest) (_result *GetAgentWorkStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAgentWorkStatusResponse{}
	_body, _err := client.GetAgentWorkStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAllDepartmentWithOptions(request *GetAllDepartmentRequest, runtime *util.RuntimeOptions) (_result *GetAllDepartmentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAllDepartment"),
		Version:     tea.String("2020-07-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAllDepartmentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAllDepartment(request *GetAllDepartmentRequest) (_result *GetAllDepartmentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAllDepartmentResponse{}
	_body, _err := client.GetAllDepartmentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAuthInfoWithOptions(request *GetAuthInfoRequest, runtime *util.RuntimeOptions) (_result *GetAuthInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppKey)) {
		query["AppKey"] = request.AppKey
	}

	if !tea.BoolValue(util.IsUnset(request.ForeignId)) {
		query["ForeignId"] = request.ForeignId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAuthInfo"),
		Version:     tea.String("2020-07-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAuthInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAuthInfo(request *GetAuthInfoRequest) (_result *GetAuthInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAuthInfoResponse{}
	_body, _err := client.GetAuthInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetByForeignIdWithOptions(request *GetByForeignIdRequest, runtime *util.RuntimeOptions) (_result *GetByForeignIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ForeignId)) {
		query["ForeignId"] = request.ForeignId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceId)) {
		query["SourceId"] = request.SourceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetByForeignId"),
		Version:     tea.String("2020-07-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetByForeignIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetByForeignId(request *GetByForeignIdRequest) (_result *GetByForeignIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetByForeignIdResponse{}
	_body, _err := client.GetByForeignIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetCMSIdByForeignIdWithOptions(request *GetCMSIdByForeignIdRequest, runtime *util.RuntimeOptions) (_result *GetCMSIdByForeignIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ForeignId)) {
		query["ForeignId"] = request.ForeignId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Nick)) {
		query["Nick"] = request.Nick
	}

	if !tea.BoolValue(util.IsUnset(request.SourceId)) {
		query["SourceId"] = request.SourceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetCMSIdByForeignId"),
		Version:     tea.String("2020-07-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCMSIdByForeignIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetCMSIdByForeignId(request *GetCMSIdByForeignIdRequest) (_result *GetCMSIdByForeignIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetCMSIdByForeignIdResponse{}
	_body, _err := client.GetCMSIdByForeignIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetCallsPerDayWithOptions(request *GetCallsPerDayRequest, runtime *util.RuntimeOptions) (_result *GetCallsPerDayResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetCallsPerDay"),
		Version:     tea.String("2020-07-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCallsPerDayResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetCallsPerDay(request *GetCallsPerDayRequest) (_result *GetCallsPerDayResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetCallsPerDayResponse{}
	_body, _err := client.GetCallsPerDayWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetEntityRouteWithOptions(request *GetEntityRouteRequest, runtime *util.RuntimeOptions) (_result *GetEntityRouteResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EntityBizCode)) {
		body["EntityBizCode"] = request.EntityBizCode
	}

	if !tea.BoolValue(util.IsUnset(request.EntityId)) {
		body["EntityId"] = request.EntityId
	}

	if !tea.BoolValue(util.IsUnset(request.EntityName)) {
		body["EntityName"] = request.EntityName
	}

	if !tea.BoolValue(util.IsUnset(request.EntityRelationNumber)) {
		body["EntityRelationNumber"] = request.EntityRelationNumber
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.UniqueId)) {
		body["UniqueId"] = request.UniqueId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetEntityRoute"),
		Version:     tea.String("2020-07-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetEntityRouteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetEntityRoute(request *GetEntityRouteRequest) (_result *GetEntityRouteResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetEntityRouteResponse{}
	_body, _err := client.GetEntityRouteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetEntityRouteListWithOptions(request *GetEntityRouteListRequest, runtime *util.RuntimeOptions) (_result *GetEntityRouteListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EntityName)) {
		body["EntityName"] = request.EntityName
	}

	if !tea.BoolValue(util.IsUnset(request.EntityRelationNumber)) {
		body["EntityRelationNumber"] = request.EntityRelationNumber
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		body["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetEntityRouteList"),
		Version:     tea.String("2020-07-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetEntityRouteListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetEntityRouteList(request *GetEntityRouteListRequest) (_result *GetEntityRouteListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetEntityRouteListResponse{}
	_body, _err := client.GetEntityRouteListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetEntityTagRelationWithOptions(request *GetEntityTagRelationRequest, runtime *util.RuntimeOptions) (_result *GetEntityTagRelationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EntityId)) {
		body["EntityId"] = request.EntityId
	}

	if !tea.BoolValue(util.IsUnset(request.EntityType)) {
		body["EntityType"] = request.EntityType
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetEntityTagRelation"),
		Version:     tea.String("2020-07-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetEntityTagRelationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetEntityTagRelation(request *GetEntityTagRelationRequest) (_result *GetEntityTagRelationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetEntityTagRelationResponse{}
	_body, _err := client.GetEntityTagRelationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetForeignIdByCMSIdWithOptions(request *GetForeignIdByCMSIdRequest, runtime *util.RuntimeOptions) (_result *GetForeignIdByCMSIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetForeignIdByCMSId"),
		Version:     tea.String("2020-07-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetForeignIdByCMSIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetForeignIdByCMSId(request *GetForeignIdByCMSIdRequest) (_result *GetForeignIdByCMSIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetForeignIdByCMSIdResponse{}
	_body, _err := client.GetForeignIdByCMSIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetGrantedRoleIdsWithOptions(request *GetGrantedRoleIdsRequest, runtime *util.RuntimeOptions) (_result *GetGrantedRoleIdsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetGrantedRoleIds"),
		Version:     tea.String("2020-07-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetGrantedRoleIdsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetGrantedRoleIds(request *GetGrantedRoleIdsRequest) (_result *GetGrantedRoleIdsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetGrantedRoleIdsResponse{}
	_body, _err := client.GetGrantedRoleIdsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetHotlineAgentDetailWithOptions(request *GetHotlineAgentDetailRequest, runtime *util.RuntimeOptions) (_result *GetHotlineAgentDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetHotlineAgentDetail"),
		Version:     tea.String("2020-07-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetHotlineAgentDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetHotlineAgentDetail(request *GetHotlineAgentDetailRequest) (_result *GetHotlineAgentDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetHotlineAgentDetailResponse{}
	_body, _err := client.GetHotlineAgentDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetHotlineAgentDetailReportWithOptions(request *GetHotlineAgentDetailReportRequest, runtime *util.RuntimeOptions) (_result *GetHotlineAgentDetailReportResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.DepIds)) {
		query["DepIds"] = request.DepIds
	}

	if !tea.BoolValue(util.IsUnset(request.EndDate)) {
		query["EndDate"] = request.EndDate
	}

	if !tea.BoolValue(util.IsUnset(request.GroupIds)) {
		query["GroupIds"] = request.GroupIds
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		query["StartDate"] = request.StartDate
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetHotlineAgentDetailReport"),
		Version:     tea.String("2020-07-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetHotlineAgentDetailReportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetHotlineAgentDetailReport(request *GetHotlineAgentDetailReportRequest) (_result *GetHotlineAgentDetailReportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetHotlineAgentDetailReportResponse{}
	_body, _err := client.GetHotlineAgentDetailReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetHotlineAgentDetailWithChannelWithOptions(request *GetHotlineAgentDetailWithChannelRequest, runtime *util.RuntimeOptions) (_result *GetHotlineAgentDetailWithChannelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetHotlineAgentDetailWithChannel"),
		Version:     tea.String("2020-07-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetHotlineAgentDetailWithChannelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetHotlineAgentDetailWithChannel(request *GetHotlineAgentDetailWithChannelRequest) (_result *GetHotlineAgentDetailWithChannelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetHotlineAgentDetailWithChannelResponse{}
	_body, _err := client.GetHotlineAgentDetailWithChannelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetHotlineAgentStatusWithOptions(request *GetHotlineAgentStatusRequest, runtime *util.RuntimeOptions) (_result *GetHotlineAgentStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountName)) {
		body["AccountName"] = request.AccountName
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetHotlineAgentStatus"),
		Version:     tea.String("2020-07-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetHotlineAgentStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetHotlineAgentStatus(request *GetHotlineAgentStatusRequest) (_result *GetHotlineAgentStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetHotlineAgentStatusResponse{}
	_body, _err := client.GetHotlineAgentStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetHotlineGroupDetailReportWithOptions(request *GetHotlineGroupDetailReportRequest, runtime *util.RuntimeOptions) (_result *GetHotlineGroupDetailReportResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.DepIds)) {
		query["DepIds"] = request.DepIds
	}

	if !tea.BoolValue(util.IsUnset(request.EndDate)) {
		query["EndDate"] = request.EndDate
	}

	if !tea.BoolValue(util.IsUnset(request.GroupIds)) {
		query["GroupIds"] = request.GroupIds
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		query["StartDate"] = request.StartDate
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetHotlineGroupDetailReport"),
		Version:     tea.String("2020-07-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetHotlineGroupDetailReportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetHotlineGroupDetailReport(request *GetHotlineGroupDetailReportRequest) (_result *GetHotlineGroupDetailReportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetHotlineGroupDetailReportResponse{}
	_body, _err := client.GetHotlineGroupDetailReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetHotlineWaitingNumberWithOptions(request *GetHotlineWaitingNumberRequest, runtime *util.RuntimeOptions) (_result *GetHotlineWaitingNumberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetHotlineWaitingNumber"),
		Version:     tea.String("2020-07-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetHotlineWaitingNumberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetHotlineWaitingNumber(request *GetHotlineWaitingNumberRequest) (_result *GetHotlineWaitingNumberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetHotlineWaitingNumberResponse{}
	_body, _err := client.GetHotlineWaitingNumberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetNumLocationWithOptions(request *GetNumLocationRequest, runtime *util.RuntimeOptions) (_result *GetNumLocationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetNumLocation"),
		Version:     tea.String("2020-07-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetNumLocationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetNumLocation(request *GetNumLocationRequest) (_result *GetNumLocationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetNumLocationResponse{}
	_body, _err := client.GetNumLocationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOutbounNumListWithOptions(request *GetOutbounNumListRequest, runtime *util.RuntimeOptions) (_result *GetOutbounNumListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountName)) {
		body["AccountName"] = request.AccountName
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetOutbounNumList"),
		Version:     tea.String("2020-07-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOutbounNumListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOutbounNumList(request *GetOutbounNumListRequest) (_result *GetOutbounNumListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOutbounNumListResponse{}
	_body, _err := client.GetOutbounNumListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetOuterCallCenterDataListWithOptions(request *GetOuterCallCenterDataListRequest, runtime *util.RuntimeOptions) (_result *GetOuterCallCenterDataListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizId)) {
		body["BizId"] = request.BizId
	}

	if !tea.BoolValue(util.IsUnset(request.FromPhoneNum)) {
		body["FromPhoneNum"] = request.FromPhoneNum
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.QueryEndTime)) {
		body["QueryEndTime"] = request.QueryEndTime
	}

	if !tea.BoolValue(util.IsUnset(request.QueryStartTime)) {
		body["QueryStartTime"] = request.QueryStartTime
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		body["SessionId"] = request.SessionId
	}

	if !tea.BoolValue(util.IsUnset(request.ToPhoneNum)) {
		body["ToPhoneNum"] = request.ToPhoneNum
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetOuterCallCenterDataList"),
		Version:     tea.String("2020-07-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOuterCallCenterDataListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetOuterCallCenterDataList(request *GetOuterCallCenterDataListRequest) (_result *GetOuterCallCenterDataListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOuterCallCenterDataListResponse{}
	_body, _err := client.GetOuterCallCenterDataListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTagListWithOptions(request *GetTagListRequest, runtime *util.RuntimeOptions) (_result *GetTagListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EntityId)) {
		body["EntityId"] = request.EntityId
	}

	if !tea.BoolValue(util.IsUnset(request.EntityType)) {
		body["EntityType"] = request.EntityType
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTagList"),
		Version:     tea.String("2020-07-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTagListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTagList(request *GetTagListRequest) (_result *GetTagListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTagListResponse{}
	_body, _err := client.GetTagListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTicketByCaseIdWithOptions(request *GetTicketByCaseIdRequest, runtime *util.RuntimeOptions) (_result *GetTicketByCaseIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CaseId)) {
		body["CaseId"] = request.CaseId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTicketByCaseId"),
		Version:     tea.String("2020-07-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTicketByCaseIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTicketByCaseId(request *GetTicketByCaseIdRequest) (_result *GetTicketByCaseIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTicketByCaseIdResponse{}
	_body, _err := client.GetTicketByCaseIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTicketTemplateSchemaWithOptions(request *GetTicketTemplateSchemaRequest, runtime *util.RuntimeOptions) (_result *GetTicketTemplateSchemaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTicketTemplateSchema"),
		Version:     tea.String("2020-07-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTicketTemplateSchemaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTicketTemplateSchema(request *GetTicketTemplateSchemaRequest) (_result *GetTicketTemplateSchemaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTicketTemplateSchemaResponse{}
	_body, _err := client.GetTicketTemplateSchemaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetUserTokenWithOptions(request *GetUserTokenRequest, runtime *util.RuntimeOptions) (_result *GetUserTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppKey)) {
		body["AppKey"] = request.AppKey
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Nick)) {
		body["Nick"] = request.Nick
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetUserToken"),
		Version:     tea.String("2020-07-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUserTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetUserToken(request *GetUserTokenRequest) (_result *GetUserTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetUserTokenResponse{}
	_body, _err := client.GetUserTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GrantRolesWithOptions(request *GrantRolesRequest, runtime *util.RuntimeOptions) (_result *GrantRolesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountName)) {
		body["AccountName"] = request.AccountName
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Operator)) {
		body["Operator"] = request.Operator
	}

	if !tea.BoolValue(util.IsUnset(request.RoleId)) {
		body["RoleId"] = request.RoleId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GrantRoles"),
		Version:     tea.String("2020-07-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GrantRolesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GrantRoles(request *GrantRolesRequest) (_result *GrantRolesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GrantRolesResponse{}
	_body, _err := client.GrantRolesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) HangUpAgentMonitorWithOptions(request *HangUpAgentMonitorRequest, runtime *util.RuntimeOptions) (_result *HangUpAgentMonitorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("HangUpAgentMonitor"),
		Version:     tea.String("2020-07-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &HangUpAgentMonitorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) HangUpAgentMonitor(request *HangUpAgentMonitorRequest) (_result *HangUpAgentMonitorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &HangUpAgentMonitorResponse{}
	_body, _err := client.HangUpAgentMonitorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) HangupCallWithOptions(request *HangupCallRequest, runtime *util.RuntimeOptions) (_result *HangupCallResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountName)) {
		body["AccountName"] = request.AccountName
	}

	if !tea.BoolValue(util.IsUnset(request.CallId)) {
		body["CallId"] = request.CallId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ConnectionId)) {
		body["ConnectionId"] = request.ConnectionId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		body["JobId"] = request.JobId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("HangupCall"),
		Version:     tea.String("2020-07-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &HangupCallResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) HangupCall(request *HangupCallRequest) (_result *HangupCallResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &HangupCallResponse{}
	_body, _err := client.HangupCallWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) HangupThirdCallWithOptions(request *HangupThirdCallRequest, runtime *util.RuntimeOptions) (_result *HangupThirdCallResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountName)) {
		body["AccountName"] = request.AccountName
	}

	if !tea.BoolValue(util.IsUnset(request.CallId)) {
		body["CallId"] = request.CallId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ConnectionId)) {
		body["ConnectionId"] = request.ConnectionId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		body["JobId"] = request.JobId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("HangupThirdCall"),
		Version:     tea.String("2020-07-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &HangupThirdCallResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) HangupThirdCall(request *HangupThirdCallRequest) (_result *HangupThirdCallResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &HangupThirdCallResponse{}
	_body, _err := client.HangupThirdCallWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) HoldCallWithOptions(request *HoldCallRequest, runtime *util.RuntimeOptions) (_result *HoldCallResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountName)) {
		body["AccountName"] = request.AccountName
	}

	if !tea.BoolValue(util.IsUnset(request.CallId)) {
		body["CallId"] = request.CallId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ConnectionId)) {
		body["ConnectionId"] = request.ConnectionId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		body["JobId"] = request.JobId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("HoldCall"),
		Version:     tea.String("2020-07-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &HoldCallResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) HoldCall(request *HoldCallRequest) (_result *HoldCallResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &HoldCallResponse{}
	_body, _err := client.HoldCallWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) JoinThirdCallWithOptions(request *JoinThirdCallRequest, runtime *util.RuntimeOptions) (_result *JoinThirdCallResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountName)) {
		body["AccountName"] = request.AccountName
	}

	if !tea.BoolValue(util.IsUnset(request.CallId)) {
		body["CallId"] = request.CallId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ConnectionId)) {
		body["ConnectionId"] = request.ConnectionId
	}

	if !tea.BoolValue(util.IsUnset(request.HoldConnectionId)) {
		body["HoldConnectionId"] = request.HoldConnectionId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		body["JobId"] = request.JobId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("JoinThirdCall"),
		Version:     tea.String("2020-07-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &JoinThirdCallResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) JoinThirdCall(request *JoinThirdCallRequest) (_result *JoinThirdCallResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &JoinThirdCallResponse{}
	_body, _err := client.JoinThirdCallWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAgentBySkillGroupIdWithOptions(request *ListAgentBySkillGroupIdRequest, runtime *util.RuntimeOptions) (_result *ListAgentBySkillGroupIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAgentBySkillGroupId"),
		Version:     tea.String("2020-07-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAgentBySkillGroupIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAgentBySkillGroupId(request *ListAgentBySkillGroupIdRequest) (_result *ListAgentBySkillGroupIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAgentBySkillGroupIdResponse{}
	_body, _err := client.ListAgentBySkillGroupIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAllHotLineSkillGroupsWithOptions(request *ListAllHotLineSkillGroupsRequest, runtime *util.RuntimeOptions) (_result *ListAllHotLineSkillGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAllHotLineSkillGroups"),
		Version:     tea.String("2020-07-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAllHotLineSkillGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAllHotLineSkillGroups(request *ListAllHotLineSkillGroupsRequest) (_result *ListAllHotLineSkillGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAllHotLineSkillGroupsResponse{}
	_body, _err := client.ListAllHotLineSkillGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListHotlineRecordWithOptions(request *ListHotlineRecordRequest, runtime *util.RuntimeOptions) (_result *ListHotlineRecordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListHotlineRecord"),
		Version:     tea.String("2020-07-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListHotlineRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListHotlineRecord(request *ListHotlineRecordRequest) (_result *ListHotlineRecordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListHotlineRecordResponse{}
	_body, _err := client.ListHotlineRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListOutboundPhoneNumberWithOptions(request *ListOutboundPhoneNumberRequest, runtime *util.RuntimeOptions) (_result *ListOutboundPhoneNumberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListOutboundPhoneNumber"),
		Version:     tea.String("2020-07-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListOutboundPhoneNumberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListOutboundPhoneNumber(request *ListOutboundPhoneNumberRequest) (_result *ListOutboundPhoneNumberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListOutboundPhoneNumberResponse{}
	_body, _err := client.ListOutboundPhoneNumberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSkillGroupWithOptions(request *ListSkillGroupRequest, runtime *util.RuntimeOptions) (_result *ListSkillGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListSkillGroup"),
		Version:     tea.String("2020-07-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSkillGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSkillGroup(request *ListSkillGroupRequest) (_result *ListSkillGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSkillGroupResponse{}
	_body, _err := client.ListSkillGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryHotlineDashboardWithOptions(tmpReq *QueryHotlineDashboardRequest, runtime *util.RuntimeOptions) (_result *QueryHotlineDashboardResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &QueryHotlineDashboardShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.DepartmentIdList)) {
		request.DepartmentIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DepartmentIdList, tea.String("DepartmentIdList"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.ServicerIdList)) {
		request.ServicerIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ServicerIdList, tea.String("ServicerIdList"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.SortFieldList)) {
		request.SortFieldListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SortFieldList, tea.String("SortFieldList"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPageNum)) {
		query["CurrentPageNum"] = request.CurrentPageNum
	}

	if !tea.BoolValue(util.IsUnset(request.DepartmentIdListShrink)) {
		query["DepartmentIdList"] = request.DepartmentIdListShrink
	}

	if !tea.BoolValue(util.IsUnset(request.EndDate)) {
		query["EndDate"] = request.EndDate
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ServicerIdListShrink)) {
		query["ServicerIdList"] = request.ServicerIdListShrink
	}

	if !tea.BoolValue(util.IsUnset(request.SortFieldListShrink)) {
		query["SortFieldList"] = request.SortFieldListShrink
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		query["StartDate"] = request.StartDate
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryHotlineDashboard"),
		Version:     tea.String("2020-07-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryHotlineDashboardResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryHotlineDashboard(request *QueryHotlineDashboardRequest) (_result *QueryHotlineDashboardResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryHotlineDashboardResponse{}
	_body, _err := client.QueryHotlineDashboardWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryHotlineSessionWithOptions(request *QueryHotlineSessionRequest, runtime *util.RuntimeOptions) (_result *QueryHotlineSessionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Acid)) {
		query["Acid"] = request.Acid
	}

	if !tea.BoolValue(util.IsUnset(request.CallResult)) {
		query["CallResult"] = request.CallResult
	}

	if !tea.BoolValue(util.IsUnset(request.CallType)) {
		query["CallType"] = request.CallType
	}

	if !tea.BoolValue(util.IsUnset(request.CalledNumber)) {
		query["CalledNumber"] = request.CalledNumber
	}

	if !tea.BoolValue(util.IsUnset(request.CallingNumber)) {
		query["CallingNumber"] = request.CallingNumber
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		query["GroupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.MemberId)) {
		query["MemberId"] = request.MemberId
	}

	if !tea.BoolValue(util.IsUnset(request.MemberName)) {
		query["MemberName"] = request.MemberName
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Params)) {
		query["Params"] = request.Params
	}

	if !tea.BoolValue(util.IsUnset(request.QueryEndTime)) {
		query["QueryEndTime"] = request.QueryEndTime
	}

	if !tea.BoolValue(util.IsUnset(request.QueryStartTime)) {
		query["QueryStartTime"] = request.QueryStartTime
	}

	if !tea.BoolValue(util.IsUnset(request.RequestId)) {
		query["RequestId"] = request.RequestId
	}

	if !tea.BoolValue(util.IsUnset(request.ServicerId)) {
		query["ServicerId"] = request.ServicerId
	}

	if !tea.BoolValue(util.IsUnset(request.ServicerName)) {
		query["ServicerName"] = request.ServicerName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryHotlineSession"),
		Version:     tea.String("2020-07-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryHotlineSessionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryHotlineSession(request *QueryHotlineSessionRequest) (_result *QueryHotlineSessionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryHotlineSessionResponse{}
	_body, _err := client.QueryHotlineSessionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryRelationTicketsWithOptions(request *QueryRelationTicketsRequest, runtime *util.RuntimeOptions) (_result *QueryRelationTicketsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CaseId)) {
		body["CaseId"] = request.CaseId
	}

	if !tea.BoolValue(util.IsUnset(request.CaseType)) {
		body["CaseType"] = request.CaseType
	}

	if !tea.BoolValue(util.IsUnset(request.ChannelId)) {
		body["ChannelId"] = request.ChannelId
	}

	if !tea.BoolValue(util.IsUnset(request.ChannelType)) {
		body["ChannelType"] = request.ChannelType
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		body["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.DealId)) {
		body["DealId"] = request.DealId
	}

	if !tea.BoolValue(util.IsUnset(request.EndCaseGmtCreate)) {
		body["EndCaseGmtCreate"] = request.EndCaseGmtCreate
	}

	bodyFlat := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Extra)) {
		bodyFlat["Extra"] = request.Extra
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.MemberId)) {
		body["MemberId"] = request.MemberId
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SrType)) {
		body["SrType"] = request.SrType
	}

	if !tea.BoolValue(util.IsUnset(request.StartCaseGmtCreate)) {
		body["StartCaseGmtCreate"] = request.StartCaseGmtCreate
	}

	if !tea.BoolValue(util.IsUnset(request.TaskStatus)) {
		body["TaskStatus"] = request.TaskStatus
	}

	if !tea.BoolValue(util.IsUnset(request.TouchId)) {
		body["TouchId"] = request.TouchId
	}

	body = tea.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryRelationTickets"),
		Version:     tea.String("2020-07-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryRelationTicketsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryRelationTickets(request *QueryRelationTicketsRequest) (_result *QueryRelationTicketsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryRelationTicketsResponse{}
	_body, _err := client.QueryRelationTicketsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryRingDetailListWithOptions(tmpReq *QueryRingDetailListRequest, runtime *util.RuntimeOptions) (_result *QueryRingDetailListResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &QueryRingDetailListShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.FromPhoneNumList)) {
		request.FromPhoneNumListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FromPhoneNumList, tea.String("FromPhoneNumList"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.ToPhoneNumList)) {
		request.ToPhoneNumListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ToPhoneNumList, tea.String("ToPhoneNumList"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CallOutStatus)) {
		body["CallOutStatus"] = request.CallOutStatus
	}

	if !tea.BoolValue(util.IsUnset(request.EndDate)) {
		body["EndDate"] = request.EndDate
	}

	if !tea.BoolValue(util.IsUnset(request.Extra)) {
		body["Extra"] = request.Extra
	}

	if !tea.BoolValue(util.IsUnset(request.FromPhoneNumListShrink)) {
		body["FromPhoneNumList"] = request.FromPhoneNumListShrink
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		body["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		body["StartDate"] = request.StartDate
	}

	if !tea.BoolValue(util.IsUnset(request.ToPhoneNumListShrink)) {
		body["ToPhoneNumList"] = request.ToPhoneNumListShrink
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryRingDetailList"),
		Version:     tea.String("2020-07-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryRingDetailListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryRingDetailList(request *QueryRingDetailListRequest) (_result *QueryRingDetailListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryRingDetailListResponse{}
	_body, _err := client.QueryRingDetailListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryServiceConfigWithOptions(request *QueryServiceConfigRequest, runtime *util.RuntimeOptions) (_result *QueryServiceConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryServiceConfig"),
		Version:     tea.String("2020-07-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryServiceConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryServiceConfig(request *QueryServiceConfigRequest) (_result *QueryServiceConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryServiceConfigResponse{}
	_body, _err := client.QueryServiceConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryServicerByDepartmentAndMixNameWithOptions(tmpReq *QueryServicerByDepartmentAndMixNameRequest, runtime *util.RuntimeOptions) (_result *QueryServicerByDepartmentAndMixNameResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &QueryServicerByDepartmentAndMixNameShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.DepartmentIdList)) {
		request.DepartmentIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DepartmentIdList, tea.String("DepartmentIdList"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPageNum)) {
		query["CurrentPageNum"] = request.CurrentPageNum
	}

	if !tea.BoolValue(util.IsUnset(request.DepartmentIdListShrink)) {
		query["DepartmentIdList"] = request.DepartmentIdListShrink
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.KeyWord)) {
		query["KeyWord"] = request.KeyWord
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryServicerByDepartmentAndMixName"),
		Version:     tea.String("2020-07-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryServicerByDepartmentAndMixNameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryServicerByDepartmentAndMixName(request *QueryServicerByDepartmentAndMixNameRequest) (_result *QueryServicerByDepartmentAndMixNameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryServicerByDepartmentAndMixNameResponse{}
	_body, _err := client.QueryServicerByDepartmentAndMixNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryServicerByIdWithOptions(request *QueryServicerByIdRequest, runtime *util.RuntimeOptions) (_result *QueryServicerByIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ServicerId)) {
		query["ServicerId"] = request.ServicerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryServicerById"),
		Version:     tea.String("2020-07-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryServicerByIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryServicerById(request *QueryServicerByIdRequest) (_result *QueryServicerByIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryServicerByIdResponse{}
	_body, _err := client.QueryServicerByIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QuerySkillGroupsWithOptions(request *QuerySkillGroupsRequest, runtime *util.RuntimeOptions) (_result *QuerySkillGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		query["GroupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.GroupType)) {
		query["GroupType"] = request.GroupType
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QuerySkillGroups"),
		Version:     tea.String("2020-07-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QuerySkillGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QuerySkillGroups(request *QuerySkillGroupsRequest) (_result *QuerySkillGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QuerySkillGroupsResponse{}
	_body, _err := client.QuerySkillGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryTicketActionsWithOptions(request *QueryTicketActionsRequest, runtime *util.RuntimeOptions) (_result *QueryTicketActionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ActionCodeList)) {
		body["ActionCodeList"] = request.ActionCodeList
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.TicketId)) {
		body["TicketId"] = request.TicketId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryTicketActions"),
		Version:     tea.String("2020-07-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryTicketActionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryTicketActions(request *QueryTicketActionsRequest) (_result *QueryTicketActionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryTicketActionsResponse{}
	_body, _err := client.QueryTicketActionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryTicketCountWithOptions(request *QueryTicketCountRequest, runtime *util.RuntimeOptions) (_result *QueryTicketCountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		query["OperatorId"] = request.OperatorId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryTicketCount"),
		Version:     tea.String("2020-07-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryTicketCountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryTicketCount(request *QueryTicketCountRequest) (_result *QueryTicketCountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryTicketCountResponse{}
	_body, _err := client.QueryTicketCountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryTicketsWithOptions(tmpReq *QueryTicketsRequest, runtime *util.RuntimeOptions) (_result *QueryTicketsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &QueryTicketsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Extra)) {
		request.ExtraShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Extra, tea.String("Extra"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountName)) {
		body["AccountName"] = request.AccountName
	}

	if !tea.BoolValue(util.IsUnset(request.CaseId)) {
		body["CaseId"] = request.CaseId
	}

	if !tea.BoolValue(util.IsUnset(request.CaseStatus)) {
		body["CaseStatus"] = request.CaseStatus
	}

	if !tea.BoolValue(util.IsUnset(request.CaseType)) {
		body["CaseType"] = request.CaseType
	}

	if !tea.BoolValue(util.IsUnset(request.ChannelId)) {
		body["ChannelId"] = request.ChannelId
	}

	if !tea.BoolValue(util.IsUnset(request.ChannelType)) {
		body["ChannelType"] = request.ChannelType
	}

	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		body["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.DealId)) {
		body["DealId"] = request.DealId
	}

	if !tea.BoolValue(util.IsUnset(request.EndCaseGmtCreate)) {
		body["EndCaseGmtCreate"] = request.EndCaseGmtCreate
	}

	if !tea.BoolValue(util.IsUnset(request.ExtraShrink)) {
		body["Extra"] = request.ExtraShrink
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.MemberId)) {
		body["MemberId"] = request.MemberId
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ParentCaseId)) {
		body["ParentCaseId"] = request.ParentCaseId
	}

	if !tea.BoolValue(util.IsUnset(request.SrType)) {
		body["SrType"] = request.SrType
	}

	if !tea.BoolValue(util.IsUnset(request.StartCaseGmtCreate)) {
		body["StartCaseGmtCreate"] = request.StartCaseGmtCreate
	}

	if !tea.BoolValue(util.IsUnset(request.TaskStatus)) {
		body["TaskStatus"] = request.TaskStatus
	}

	if !tea.BoolValue(util.IsUnset(request.TouchId)) {
		body["TouchId"] = request.TouchId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryTickets"),
		Version:     tea.String("2020-07-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryTicketsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryTickets(request *QueryTicketsRequest) (_result *QueryTicketsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryTicketsResponse{}
	_body, _err := client.QueryTicketsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveSkillGroupWithOptions(request *RemoveSkillGroupRequest, runtime *util.RuntimeOptions) (_result *RemoveSkillGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.SkillGroupId)) {
		body["SkillGroupId"] = request.SkillGroupId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveSkillGroup"),
		Version:     tea.String("2020-07-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveSkillGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveSkillGroup(request *RemoveSkillGroupRequest) (_result *RemoveSkillGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveSkillGroupResponse{}
	_body, _err := client.RemoveSkillGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SearchTicketByIdWithOptions(request *SearchTicketByIdRequest, runtime *util.RuntimeOptions) (_result *SearchTicketByIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SearchTicketById"),
		Version:     tea.String("2020-07-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SearchTicketByIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SearchTicketById(request *SearchTicketByIdRequest) (_result *SearchTicketByIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SearchTicketByIdResponse{}
	_body, _err := client.SearchTicketByIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SearchTicketByPhoneWithOptions(request *SearchTicketByPhoneRequest, runtime *util.RuntimeOptions) (_result *SearchTicketByPhoneResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SearchTicketByPhone"),
		Version:     tea.String("2020-07-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SearchTicketByPhoneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SearchTicketByPhone(request *SearchTicketByPhoneRequest) (_result *SearchTicketByPhoneResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SearchTicketByPhoneResponse{}
	_body, _err := client.SearchTicketByPhoneWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SearchTicketListWithOptions(request *SearchTicketListRequest, runtime *util.RuntimeOptions) (_result *SearchTicketListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SearchTicketList"),
		Version:     tea.String("2020-07-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SearchTicketListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SearchTicketList(request *SearchTicketListRequest) (_result *SearchTicketListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SearchTicketListResponse{}
	_body, _err := client.SearchTicketListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SendHotlineHeartBeatWithOptions(request *SendHotlineHeartBeatRequest, runtime *util.RuntimeOptions) (_result *SendHotlineHeartBeatResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountName)) {
		body["AccountName"] = request.AccountName
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Token)) {
		body["Token"] = request.Token
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SendHotlineHeartBeat"),
		Version:     tea.String("2020-07-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SendHotlineHeartBeatResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SendHotlineHeartBeat(request *SendHotlineHeartBeatRequest) (_result *SendHotlineHeartBeatResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SendHotlineHeartBeatResponse{}
	_body, _err := client.SendHotlineHeartBeatWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SendOutboundCommandWithOptions(request *SendOutboundCommandRequest, runtime *util.RuntimeOptions) (_result *SendOutboundCommandResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountName)) {
		body["AccountName"] = request.AccountName
	}

	if !tea.BoolValue(util.IsUnset(request.CalledNumber)) {
		body["CalledNumber"] = request.CalledNumber
	}

	if !tea.BoolValue(util.IsUnset(request.CallingNumber)) {
		body["CallingNumber"] = request.CallingNumber
	}

	if !tea.BoolValue(util.IsUnset(request.CustomerInfo)) {
		body["CustomerInfo"] = request.CustomerInfo
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SendOutboundCommand"),
		Version:     tea.String("2020-07-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SendOutboundCommandResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SendOutboundCommand(request *SendOutboundCommandRequest) (_result *SendOutboundCommandResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SendOutboundCommandResponse{}
	_body, _err := client.SendOutboundCommandWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SendVerificationCodeWithOptions(request *SendVerificationCodeRequest, runtime *util.RuntimeOptions) (_result *SendVerificationCodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Phone)) {
		body["Phone"] = request.Phone
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SendVerificationCode"),
		Version:     tea.String("2020-07-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SendVerificationCodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SendVerificationCode(request *SendVerificationCodeRequest) (_result *SendVerificationCodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SendVerificationCodeResponse{}
	_body, _err := client.SendVerificationCodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartAgentMonitorWithOptions(request *StartAgentMonitorRequest, runtime *util.RuntimeOptions) (_result *StartAgentMonitorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StartAgentMonitor"),
		Version:     tea.String("2020-07-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StartAgentMonitorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartAgentMonitor(request *StartAgentMonitorRequest) (_result *StartAgentMonitorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartAgentMonitorResponse{}
	_body, _err := client.StartAgentMonitorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartCallWithOptions(request *StartCallRequest, runtime *util.RuntimeOptions) (_result *StartCallResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountName)) {
		body["AccountName"] = request.AccountName
	}

	if !tea.BoolValue(util.IsUnset(request.Callee)) {
		body["Callee"] = request.Callee
	}

	if !tea.BoolValue(util.IsUnset(request.Caller)) {
		body["Caller"] = request.Caller
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("StartCall"),
		Version:     tea.String("2020-07-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StartCallResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartCall(request *StartCallRequest) (_result *StartCallResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartCallResponse{}
	_body, _err := client.StartCallWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartCallV2WithOptions(request *StartCallV2Request, runtime *util.RuntimeOptions) (_result *StartCallV2Response, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountName)) {
		body["AccountName"] = request.AccountName
	}

	if !tea.BoolValue(util.IsUnset(request.Callee)) {
		body["Callee"] = request.Callee
	}

	if !tea.BoolValue(util.IsUnset(request.Caller)) {
		body["Caller"] = request.Caller
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.JsonMsg)) {
		body["JsonMsg"] = request.JsonMsg
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("StartCallV2"),
		Version:     tea.String("2020-07-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StartCallV2Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartCallV2(request *StartCallV2Request) (_result *StartCallV2Response, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartCallV2Response{}
	_body, _err := client.StartCallV2WithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartChatWorkWithOptions(request *StartChatWorkRequest, runtime *util.RuntimeOptions) (_result *StartChatWorkResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountName)) {
		body["AccountName"] = request.AccountName
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("StartChatWork"),
		Version:     tea.String("2020-07-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StartChatWorkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartChatWork(request *StartChatWorkRequest) (_result *StartChatWorkResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartChatWorkResponse{}
	_body, _err := client.StartChatWorkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartHotlineServiceWithOptions(request *StartHotlineServiceRequest, runtime *util.RuntimeOptions) (_result *StartHotlineServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountName)) {
		body["AccountName"] = request.AccountName
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("StartHotlineService"),
		Version:     tea.String("2020-07-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StartHotlineServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartHotlineService(request *StartHotlineServiceRequest) (_result *StartHotlineServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartHotlineServiceResponse{}
	_body, _err := client.StartHotlineServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartHotlineWorkWithOptions(request *StartHotlineWorkRequest, runtime *util.RuntimeOptions) (_result *StartHotlineWorkResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountName)) {
		body["AccountName"] = request.AccountName
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.SkillGroups)) {
		body["SkillGroups"] = request.SkillGroups
	}

	if !tea.BoolValue(util.IsUnset(request.WorkChannel)) {
		body["WorkChannel"] = request.WorkChannel
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("StartHotlineWork"),
		Version:     tea.String("2020-07-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StartHotlineWorkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartHotlineWork(request *StartHotlineWorkRequest) (_result *StartHotlineWorkResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartHotlineWorkResponse{}
	_body, _err := client.StartHotlineWorkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SuspendHotlineServiceWithOptions(request *SuspendHotlineServiceRequest, runtime *util.RuntimeOptions) (_result *SuspendHotlineServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountName)) {
		body["AccountName"] = request.AccountName
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SuspendHotlineService"),
		Version:     tea.String("2020-07-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SuspendHotlineServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SuspendHotlineService(request *SuspendHotlineServiceRequest) (_result *SuspendHotlineServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SuspendHotlineServiceResponse{}
	_body, _err := client.SuspendHotlineServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TransferCallToAgentWithOptions(request *TransferCallToAgentRequest, runtime *util.RuntimeOptions) (_result *TransferCallToAgentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountName)) {
		body["AccountName"] = request.AccountName
	}

	if !tea.BoolValue(util.IsUnset(request.CallId)) {
		body["CallId"] = request.CallId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ConnectionId)) {
		body["ConnectionId"] = request.ConnectionId
	}

	if !tea.BoolValue(util.IsUnset(request.HoldConnectionId)) {
		body["HoldConnectionId"] = request.HoldConnectionId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.IsSingleTransfer)) {
		body["IsSingleTransfer"] = request.IsSingleTransfer
	}

	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		body["JobId"] = request.JobId
	}

	if !tea.BoolValue(util.IsUnset(request.TargetAccountName)) {
		body["TargetAccountName"] = request.TargetAccountName
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("TransferCallToAgent"),
		Version:     tea.String("2020-07-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TransferCallToAgentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TransferCallToAgent(request *TransferCallToAgentRequest) (_result *TransferCallToAgentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TransferCallToAgentResponse{}
	_body, _err := client.TransferCallToAgentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TransferCallToPhoneWithOptions(request *TransferCallToPhoneRequest, runtime *util.RuntimeOptions) (_result *TransferCallToPhoneResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountName)) {
		body["AccountName"] = request.AccountName
	}

	if !tea.BoolValue(util.IsUnset(request.CallId)) {
		body["CallId"] = request.CallId
	}

	if !tea.BoolValue(util.IsUnset(request.Callee)) {
		body["Callee"] = request.Callee
	}

	if !tea.BoolValue(util.IsUnset(request.Caller)) {
		body["Caller"] = request.Caller
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ConnectionId)) {
		body["ConnectionId"] = request.ConnectionId
	}

	if !tea.BoolValue(util.IsUnset(request.HoldConnectionId)) {
		body["HoldConnectionId"] = request.HoldConnectionId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.IsSingleTransfer)) {
		body["IsSingleTransfer"] = request.IsSingleTransfer
	}

	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		body["JobId"] = request.JobId
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["Type"] = request.Type
	}

	if !tea.BoolValue(util.IsUnset(request.CalleePhone)) {
		body["calleePhone"] = request.CalleePhone
	}

	if !tea.BoolValue(util.IsUnset(request.CallerPhone)) {
		body["callerPhone"] = request.CallerPhone
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("TransferCallToPhone"),
		Version:     tea.String("2020-07-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TransferCallToPhoneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TransferCallToPhone(request *TransferCallToPhoneRequest) (_result *TransferCallToPhoneResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TransferCallToPhoneResponse{}
	_body, _err := client.TransferCallToPhoneWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TransferCallToSkillGroupWithOptions(request *TransferCallToSkillGroupRequest, runtime *util.RuntimeOptions) (_result *TransferCallToSkillGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountName)) {
		body["AccountName"] = request.AccountName
	}

	if !tea.BoolValue(util.IsUnset(request.CallId)) {
		body["CallId"] = request.CallId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ConnectionId)) {
		body["ConnectionId"] = request.ConnectionId
	}

	if !tea.BoolValue(util.IsUnset(request.HoldConnectionId)) {
		body["HoldConnectionId"] = request.HoldConnectionId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.IsSingleTransfer)) {
		body["IsSingleTransfer"] = request.IsSingleTransfer
	}

	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		body["JobId"] = request.JobId
	}

	if !tea.BoolValue(util.IsUnset(request.SkillGroupId)) {
		body["SkillGroupId"] = request.SkillGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("TransferCallToSkillGroup"),
		Version:     tea.String("2020-07-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TransferCallToSkillGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TransferCallToSkillGroup(request *TransferCallToSkillGroupRequest) (_result *TransferCallToSkillGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TransferCallToSkillGroupResponse{}
	_body, _err := client.TransferCallToSkillGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TransferToThirdCallWithOptions(request *TransferToThirdCallRequest, runtime *util.RuntimeOptions) (_result *TransferToThirdCallResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountName)) {
		body["AccountName"] = request.AccountName
	}

	if !tea.BoolValue(util.IsUnset(request.CallId)) {
		body["CallId"] = request.CallId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ConnectionId)) {
		body["ConnectionId"] = request.ConnectionId
	}

	if !tea.BoolValue(util.IsUnset(request.HoldConnectionId)) {
		body["HoldConnectionId"] = request.HoldConnectionId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		body["JobId"] = request.JobId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("TransferToThirdCall"),
		Version:     tea.String("2020-07-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TransferToThirdCallResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TransferToThirdCall(request *TransferToThirdCallRequest) (_result *TransferToThirdCallResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TransferToThirdCallResponse{}
	_body, _err := client.TransferToThirdCallWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UnbindAgentHotlinePhoneWithOptions(request *UnbindAgentHotlinePhoneRequest, runtime *util.RuntimeOptions) (_result *UnbindAgentHotlinePhoneResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountName)) {
		body["AccountName"] = request.AccountName
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UnbindAgentHotlinePhone"),
		Version:     tea.String("2020-07-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UnbindAgentHotlinePhoneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UnbindAgentHotlinePhone(request *UnbindAgentHotlinePhoneRequest) (_result *UnbindAgentHotlinePhoneResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UnbindAgentHotlinePhoneResponse{}
	_body, _err := client.UnbindAgentHotlinePhoneWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateAgentWithOptions(request *UpdateAgentRequest, runtime *util.RuntimeOptions) (_result *UpdateAgentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountName)) {
		body["AccountName"] = request.AccountName
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DisplayName)) {
		body["DisplayName"] = request.DisplayName
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.SkillGroupId)) {
		body["SkillGroupId"] = request.SkillGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SkillGroupIdList)) {
		body["SkillGroupIdList"] = request.SkillGroupIdList
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateAgent"),
		Version:     tea.String("2020-07-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateAgentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateAgent(request *UpdateAgentRequest) (_result *UpdateAgentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateAgentResponse{}
	_body, _err := client.UpdateAgentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateCustomerWithOptions(request *UpdateCustomerRequest, runtime *util.RuntimeOptions) (_result *UpdateCustomerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		query["BizType"] = request.BizType
	}

	if !tea.BoolValue(util.IsUnset(request.Contacter)) {
		query["Contacter"] = request.Contacter
	}

	if !tea.BoolValue(util.IsUnset(request.CustomerId)) {
		query["CustomerId"] = request.CustomerId
	}

	if !tea.BoolValue(util.IsUnset(request.Dingding)) {
		query["Dingding"] = request.Dingding
	}

	if !tea.BoolValue(util.IsUnset(request.Email)) {
		query["Email"] = request.Email
	}

	if !tea.BoolValue(util.IsUnset(request.Industry)) {
		query["Industry"] = request.Industry
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ManagerName)) {
		query["ManagerName"] = request.ManagerName
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OuterId)) {
		query["OuterId"] = request.OuterId
	}

	if !tea.BoolValue(util.IsUnset(request.OuterIdType)) {
		query["OuterIdType"] = request.OuterIdType
	}

	if !tea.BoolValue(util.IsUnset(request.Phone)) {
		query["Phone"] = request.Phone
	}

	if !tea.BoolValue(util.IsUnset(request.Position)) {
		query["Position"] = request.Position
	}

	if !tea.BoolValue(util.IsUnset(request.ProdLineId)) {
		query["ProdLineId"] = request.ProdLineId
	}

	if !tea.BoolValue(util.IsUnset(request.TypeCode)) {
		query["TypeCode"] = request.TypeCode
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateCustomer"),
		Version:     tea.String("2020-07-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateCustomerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateCustomer(request *UpdateCustomerRequest) (_result *UpdateCustomerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateCustomerResponse{}
	_body, _err := client.UpdateCustomerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateEntityTagRelationWithOptions(request *UpdateEntityTagRelationRequest, runtime *util.RuntimeOptions) (_result *UpdateEntityTagRelationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EntityTagParam)) {
		body["EntityTagParam"] = request.EntityTagParam
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateEntityTagRelation"),
		Version:     tea.String("2020-07-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateEntityTagRelationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateEntityTagRelation(request *UpdateEntityTagRelationRequest) (_result *UpdateEntityTagRelationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateEntityTagRelationResponse{}
	_body, _err := client.UpdateEntityTagRelationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateRingStatusWithOptions(request *UpdateRingStatusRequest, runtime *util.RuntimeOptions) (_result *UpdateRingStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CallOutStatus)) {
		body["CallOutStatus"] = request.CallOutStatus
	}

	if !tea.BoolValue(util.IsUnset(request.Extra)) {
		body["Extra"] = request.Extra
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.UniqueBizId)) {
		body["UniqueBizId"] = request.UniqueBizId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateRingStatus"),
		Version:     tea.String("2020-07-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateRingStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateRingStatus(request *UpdateRingStatusRequest) (_result *UpdateRingStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateRingStatusResponse{}
	_body, _err := client.UpdateRingStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateRoleWithOptions(request *UpdateRoleRequest, runtime *util.RuntimeOptions) (_result *UpdateRoleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Operator)) {
		body["Operator"] = request.Operator
	}

	if !tea.BoolValue(util.IsUnset(request.PermissionId)) {
		body["PermissionId"] = request.PermissionId
	}

	if !tea.BoolValue(util.IsUnset(request.RoleId)) {
		body["RoleId"] = request.RoleId
	}

	if !tea.BoolValue(util.IsUnset(request.RoleName)) {
		body["RoleName"] = request.RoleName
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateRole"),
		Version:     tea.String("2020-07-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateRoleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateRole(request *UpdateRoleRequest) (_result *UpdateRoleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateRoleResponse{}
	_body, _err := client.UpdateRoleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateSkillGroupWithOptions(request *UpdateSkillGroupRequest, runtime *util.RuntimeOptions) (_result *UpdateSkillGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChannelType)) {
		query["ChannelType"] = request.ChannelType
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DisplayName)) {
		query["DisplayName"] = request.DisplayName
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.SkillGroupId)) {
		query["SkillGroupId"] = request.SkillGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SkillGroupName)) {
		query["SkillGroupName"] = request.SkillGroupName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateSkillGroup"),
		Version:     tea.String("2020-07-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateSkillGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateSkillGroup(request *UpdateSkillGroupRequest) (_result *UpdateSkillGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateSkillGroupResponse{}
	_body, _err := client.UpdateSkillGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateTicketWithOptions(request *UpdateTicketRequest, runtime *util.RuntimeOptions) (_result *UpdateTicketResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.FormData)) {
		body["FormData"] = request.FormData
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OperatorId)) {
		body["OperatorId"] = request.OperatorId
	}

	if !tea.BoolValue(util.IsUnset(request.TicketId)) {
		body["TicketId"] = request.TicketId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateTicket"),
		Version:     tea.String("2020-07-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateTicketResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateTicket(request *UpdateTicketRequest) (_result *UpdateTicketResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateTicketResponse{}
	_body, _err := client.UpdateTicketWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateTicketWithBizDataWithOptions(request *CreateTicketWithBizDataRequest, runtime *util.RuntimeOptions) (_result *CreateTicketWithBizDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizData)) {
		body["BizData"] = request.BizData
	}

	if !tea.BoolValue(util.IsUnset(request.CarbonCopy)) {
		body["CarbonCopy"] = request.CarbonCopy
	}

	if !tea.BoolValue(util.IsUnset(request.CategoryId)) {
		body["CategoryId"] = request.CategoryId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.CreatorId)) {
		body["CreatorId"] = request.CreatorId
	}

	if !tea.BoolValue(util.IsUnset(request.CreatorName)) {
		body["CreatorName"] = request.CreatorName
	}

	if !tea.BoolValue(util.IsUnset(request.CreatorType)) {
		body["CreatorType"] = request.CreatorType
	}

	if !tea.BoolValue(util.IsUnset(request.FormData)) {
		body["FormData"] = request.FormData
	}

	if !tea.BoolValue(util.IsUnset(request.FromInfo)) {
		body["FromInfo"] = request.FromInfo
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.MemberId)) {
		body["MemberId"] = request.MemberId
	}

	if !tea.BoolValue(util.IsUnset(request.MemberName)) {
		body["MemberName"] = request.MemberName
	}

	if !tea.BoolValue(util.IsUnset(request.Priority)) {
		body["Priority"] = request.Priority
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		body["TemplateId"] = request.TemplateId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("createTicketWithBizData"),
		Version:     tea.String("2020-07-02"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateTicketWithBizDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateTicketWithBizData(request *CreateTicketWithBizDataRequest) (_result *CreateTicketWithBizDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateTicketWithBizDataResponse{}
	_body, _err := client.CreateTicketWithBizDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
