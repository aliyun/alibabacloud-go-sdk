// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAuditTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryAuditTaskResponseBody
	GetCode() *string
	SetData(v *QueryAuditTaskResponseBodyData) *QueryAuditTaskResponseBody
	GetData() *QueryAuditTaskResponseBodyData
	SetHttpStatusCode(v int32) *QueryAuditTaskResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *QueryAuditTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryAuditTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryAuditTaskResponseBody
	GetSuccess() *bool
}

type QueryAuditTaskResponseBody struct {
	// example:
	//
	// DataNotExists
	Code *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *QueryAuditTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 400
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 错误消息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// xxxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryAuditTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryAuditTaskResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAuditTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryAuditTaskResponseBody) GetData() *QueryAuditTaskResponseBodyData {
	return s.Data
}

func (s *QueryAuditTaskResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *QueryAuditTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryAuditTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryAuditTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryAuditTaskResponseBody) SetCode(v string) *QueryAuditTaskResponseBody {
	s.Code = &v
	return s
}

func (s *QueryAuditTaskResponseBody) SetData(v *QueryAuditTaskResponseBodyData) *QueryAuditTaskResponseBody {
	s.Data = v
	return s
}

func (s *QueryAuditTaskResponseBody) SetHttpStatusCode(v int32) *QueryAuditTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *QueryAuditTaskResponseBody) SetMessage(v string) *QueryAuditTaskResponseBody {
	s.Message = &v
	return s
}

func (s *QueryAuditTaskResponseBody) SetRequestId(v string) *QueryAuditTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryAuditTaskResponseBody) SetSuccess(v bool) *QueryAuditTaskResponseBody {
	s.Success = &v
	return s
}

func (s *QueryAuditTaskResponseBody) Validate() error {
	return dara.Validate(s)
}

type QueryAuditTaskResponseBodyData struct {
	// example:
	//
	// 2025-05-13 12:12:12
	AuditTime *string `json:"AuditTime,omitempty" xml:"AuditTime,omitempty"`
	// example:
	//
	// 审核时的原文
	Content     *string                                 `json:"Content,omitempty" xml:"Content,omitempty"`
	HtmlContent *string                                 `json:"HtmlContent,omitempty" xml:"HtmlContent,omitempty"`
	Response    *QueryAuditTaskResponseBodyDataResponse `json:"Response,omitempty" xml:"Response,omitempty" type:"Struct"`
	// example:
	//
	// RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 1
	TaskStatus *int32  `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	Title      *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s QueryAuditTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryAuditTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryAuditTaskResponseBodyData) GetAuditTime() *string {
	return s.AuditTime
}

func (s *QueryAuditTaskResponseBodyData) GetContent() *string {
	return s.Content
}

func (s *QueryAuditTaskResponseBodyData) GetHtmlContent() *string {
	return s.HtmlContent
}

func (s *QueryAuditTaskResponseBodyData) GetResponse() *QueryAuditTaskResponseBodyDataResponse {
	return s.Response
}

func (s *QueryAuditTaskResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *QueryAuditTaskResponseBodyData) GetTaskStatus() *int32 {
	return s.TaskStatus
}

func (s *QueryAuditTaskResponseBodyData) GetTitle() *string {
	return s.Title
}

func (s *QueryAuditTaskResponseBodyData) SetAuditTime(v string) *QueryAuditTaskResponseBodyData {
	s.AuditTime = &v
	return s
}

func (s *QueryAuditTaskResponseBodyData) SetContent(v string) *QueryAuditTaskResponseBodyData {
	s.Content = &v
	return s
}

func (s *QueryAuditTaskResponseBodyData) SetHtmlContent(v string) *QueryAuditTaskResponseBodyData {
	s.HtmlContent = &v
	return s
}

func (s *QueryAuditTaskResponseBodyData) SetResponse(v *QueryAuditTaskResponseBodyDataResponse) *QueryAuditTaskResponseBodyData {
	s.Response = v
	return s
}

func (s *QueryAuditTaskResponseBodyData) SetStatus(v string) *QueryAuditTaskResponseBodyData {
	s.Status = &v
	return s
}

func (s *QueryAuditTaskResponseBodyData) SetTaskStatus(v int32) *QueryAuditTaskResponseBodyData {
	s.TaskStatus = &v
	return s
}

func (s *QueryAuditTaskResponseBodyData) SetTitle(v string) *QueryAuditTaskResponseBodyData {
	s.Title = &v
	return s
}

func (s *QueryAuditTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type QueryAuditTaskResponseBodyDataResponse struct {
	Header  *QueryAuditTaskResponseBodyDataResponseHeader  `json:"Header,omitempty" xml:"Header,omitempty" type:"Struct"`
	Payload *QueryAuditTaskResponseBodyDataResponsePayload `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
}

func (s QueryAuditTaskResponseBodyDataResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryAuditTaskResponseBodyDataResponse) GoString() string {
	return s.String()
}

func (s *QueryAuditTaskResponseBodyDataResponse) GetHeader() *QueryAuditTaskResponseBodyDataResponseHeader {
	return s.Header
}

func (s *QueryAuditTaskResponseBodyDataResponse) GetPayload() *QueryAuditTaskResponseBodyDataResponsePayload {
	return s.Payload
}

func (s *QueryAuditTaskResponseBodyDataResponse) SetHeader(v *QueryAuditTaskResponseBodyDataResponseHeader) *QueryAuditTaskResponseBodyDataResponse {
	s.Header = v
	return s
}

func (s *QueryAuditTaskResponseBodyDataResponse) SetPayload(v *QueryAuditTaskResponseBodyDataResponsePayload) *QueryAuditTaskResponseBodyDataResponse {
	s.Payload = v
	return s
}

func (s *QueryAuditTaskResponseBodyDataResponse) Validate() error {
	return dara.Validate(s)
}

type QueryAuditTaskResponseBodyDataResponseHeader struct {
	// example:
	//
	// DataNotExists
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// 数据不存在
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// task-failed
	Event *string `json:"Event,omitempty" xml:"Event,omitempty"`
	// example:
	//
	// 49eab783-9172-487a-b9df-c6372c47392c
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// example:
	//
	// 896b733535274d28b1a61c78bc145217
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s QueryAuditTaskResponseBodyDataResponseHeader) String() string {
	return dara.Prettify(s)
}

func (s QueryAuditTaskResponseBodyDataResponseHeader) GoString() string {
	return s.String()
}

func (s *QueryAuditTaskResponseBodyDataResponseHeader) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *QueryAuditTaskResponseBodyDataResponseHeader) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *QueryAuditTaskResponseBodyDataResponseHeader) GetEvent() *string {
	return s.Event
}

func (s *QueryAuditTaskResponseBodyDataResponseHeader) GetSessionId() *string {
	return s.SessionId
}

func (s *QueryAuditTaskResponseBodyDataResponseHeader) GetTaskId() *string {
	return s.TaskId
}

func (s *QueryAuditTaskResponseBodyDataResponseHeader) SetErrorCode(v string) *QueryAuditTaskResponseBodyDataResponseHeader {
	s.ErrorCode = &v
	return s
}

func (s *QueryAuditTaskResponseBodyDataResponseHeader) SetErrorMessage(v string) *QueryAuditTaskResponseBodyDataResponseHeader {
	s.ErrorMessage = &v
	return s
}

func (s *QueryAuditTaskResponseBodyDataResponseHeader) SetEvent(v string) *QueryAuditTaskResponseBodyDataResponseHeader {
	s.Event = &v
	return s
}

func (s *QueryAuditTaskResponseBodyDataResponseHeader) SetSessionId(v string) *QueryAuditTaskResponseBodyDataResponseHeader {
	s.SessionId = &v
	return s
}

func (s *QueryAuditTaskResponseBodyDataResponseHeader) SetTaskId(v string) *QueryAuditTaskResponseBodyDataResponseHeader {
	s.TaskId = &v
	return s
}

func (s *QueryAuditTaskResponseBodyDataResponseHeader) Validate() error {
	return dara.Validate(s)
}

type QueryAuditTaskResponseBodyDataResponsePayload struct {
	Output *QueryAuditTaskResponseBodyDataResponsePayloadOutput `json:"Output,omitempty" xml:"Output,omitempty" type:"Struct"`
	Usage  *QueryAuditTaskResponseBodyDataResponsePayloadUsage  `json:"Usage,omitempty" xml:"Usage,omitempty" type:"Struct"`
}

func (s QueryAuditTaskResponseBodyDataResponsePayload) String() string {
	return dara.Prettify(s)
}

func (s QueryAuditTaskResponseBodyDataResponsePayload) GoString() string {
	return s.String()
}

func (s *QueryAuditTaskResponseBodyDataResponsePayload) GetOutput() *QueryAuditTaskResponseBodyDataResponsePayloadOutput {
	return s.Output
}

func (s *QueryAuditTaskResponseBodyDataResponsePayload) GetUsage() *QueryAuditTaskResponseBodyDataResponsePayloadUsage {
	return s.Usage
}

func (s *QueryAuditTaskResponseBodyDataResponsePayload) SetOutput(v *QueryAuditTaskResponseBodyDataResponsePayloadOutput) *QueryAuditTaskResponseBodyDataResponsePayload {
	s.Output = v
	return s
}

func (s *QueryAuditTaskResponseBodyDataResponsePayload) SetUsage(v *QueryAuditTaskResponseBodyDataResponsePayloadUsage) *QueryAuditTaskResponseBodyDataResponsePayload {
	s.Usage = v
	return s
}

func (s *QueryAuditTaskResponseBodyDataResponsePayload) Validate() error {
	return dara.Validate(s)
}

type QueryAuditTaskResponseBodyDataResponsePayloadOutput struct {
	// example:
	//
	// x\\"x\\"x
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s QueryAuditTaskResponseBodyDataResponsePayloadOutput) String() string {
	return dara.Prettify(s)
}

func (s QueryAuditTaskResponseBodyDataResponsePayloadOutput) GoString() string {
	return s.String()
}

func (s *QueryAuditTaskResponseBodyDataResponsePayloadOutput) GetText() *string {
	return s.Text
}

func (s *QueryAuditTaskResponseBodyDataResponsePayloadOutput) SetText(v string) *QueryAuditTaskResponseBodyDataResponsePayloadOutput {
	s.Text = &v
	return s
}

func (s *QueryAuditTaskResponseBodyDataResponsePayloadOutput) Validate() error {
	return dara.Validate(s)
}

type QueryAuditTaskResponseBodyDataResponsePayloadUsage struct {
	// example:
	//
	// 200
	InputTokens *int32 `json:"InputTokens,omitempty" xml:"InputTokens,omitempty"`
	// example:
	//
	// 100
	OutputTokens *int32 `json:"OutputTokens,omitempty" xml:"OutputTokens,omitempty"`
	// example:
	//
	// 300
	TotalTokens *int32 `json:"TotalTokens,omitempty" xml:"TotalTokens,omitempty"`
}

func (s QueryAuditTaskResponseBodyDataResponsePayloadUsage) String() string {
	return dara.Prettify(s)
}

func (s QueryAuditTaskResponseBodyDataResponsePayloadUsage) GoString() string {
	return s.String()
}

func (s *QueryAuditTaskResponseBodyDataResponsePayloadUsage) GetInputTokens() *int32 {
	return s.InputTokens
}

func (s *QueryAuditTaskResponseBodyDataResponsePayloadUsage) GetOutputTokens() *int32 {
	return s.OutputTokens
}

func (s *QueryAuditTaskResponseBodyDataResponsePayloadUsage) GetTotalTokens() *int32 {
	return s.TotalTokens
}

func (s *QueryAuditTaskResponseBodyDataResponsePayloadUsage) SetInputTokens(v int32) *QueryAuditTaskResponseBodyDataResponsePayloadUsage {
	s.InputTokens = &v
	return s
}

func (s *QueryAuditTaskResponseBodyDataResponsePayloadUsage) SetOutputTokens(v int32) *QueryAuditTaskResponseBodyDataResponsePayloadUsage {
	s.OutputTokens = &v
	return s
}

func (s *QueryAuditTaskResponseBodyDataResponsePayloadUsage) SetTotalTokens(v int32) *QueryAuditTaskResponseBodyDataResponsePayloadUsage {
	s.TotalTokens = &v
	return s
}

func (s *QueryAuditTaskResponseBodyDataResponsePayloadUsage) Validate() error {
	return dara.Validate(s)
}
