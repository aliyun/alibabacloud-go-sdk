// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunAiHelperWritingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RunAiHelperWritingResponseBody
	GetCode() *string
	SetHeader(v *RunAiHelperWritingResponseBodyHeader) *RunAiHelperWritingResponseBody
	GetHeader() *RunAiHelperWritingResponseBodyHeader
	SetHttpStatusCode(v string) *RunAiHelperWritingResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *RunAiHelperWritingResponseBody
	GetMessage() *string
	SetPayload(v *RunAiHelperWritingResponseBodyPayload) *RunAiHelperWritingResponseBody
	GetPayload() *RunAiHelperWritingResponseBodyPayload
	SetRequestId(v string) *RunAiHelperWritingResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RunAiHelperWritingResponseBody
	GetSuccess() *bool
}

type RunAiHelperWritingResponseBody struct {
	// 业务处理结果状态码
	//
	// example:
	//
	// successful
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 流式响应的头部信息，包含事件类型、状态码等元数据
	Header *RunAiHelperWritingResponseBodyHeader `json:"Header,omitempty" xml:"Header,omitempty" type:"Struct"`
	// HTTP响应状态码
	//
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// 业务处理结果描述信息
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 包含写作输出内容和Token使用量统计
	Payload *RunAiHelperWritingResponseBodyPayload `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	// 本次API请求的唯一标识
	//
	// example:
	//
	// 1813ceee-7fe5-41b4-87e5-982a4d18cca5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 请求是否处理成功
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RunAiHelperWritingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RunAiHelperWritingResponseBody) GoString() string {
	return s.String()
}

func (s *RunAiHelperWritingResponseBody) GetCode() *string {
	return s.Code
}

func (s *RunAiHelperWritingResponseBody) GetHeader() *RunAiHelperWritingResponseBodyHeader {
	return s.Header
}

func (s *RunAiHelperWritingResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *RunAiHelperWritingResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RunAiHelperWritingResponseBody) GetPayload() *RunAiHelperWritingResponseBodyPayload {
	return s.Payload
}

func (s *RunAiHelperWritingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RunAiHelperWritingResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RunAiHelperWritingResponseBody) SetCode(v string) *RunAiHelperWritingResponseBody {
	s.Code = &v
	return s
}

func (s *RunAiHelperWritingResponseBody) SetHeader(v *RunAiHelperWritingResponseBodyHeader) *RunAiHelperWritingResponseBody {
	s.Header = v
	return s
}

func (s *RunAiHelperWritingResponseBody) SetHttpStatusCode(v string) *RunAiHelperWritingResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *RunAiHelperWritingResponseBody) SetMessage(v string) *RunAiHelperWritingResponseBody {
	s.Message = &v
	return s
}

func (s *RunAiHelperWritingResponseBody) SetPayload(v *RunAiHelperWritingResponseBodyPayload) *RunAiHelperWritingResponseBody {
	s.Payload = v
	return s
}

func (s *RunAiHelperWritingResponseBody) SetRequestId(v string) *RunAiHelperWritingResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunAiHelperWritingResponseBody) SetSuccess(v bool) *RunAiHelperWritingResponseBody {
	s.Success = &v
	return s
}

func (s *RunAiHelperWritingResponseBody) Validate() error {
	if s.Header != nil {
		if err := s.Header.Validate(); err != nil {
			return err
		}
	}
	if s.Payload != nil {
		if err := s.Payload.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RunAiHelperWritingResponseBodyHeader struct {
	// 请求错误时的错误码
	//
	// example:
	//
	// InvalidParameter
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// 请求错误时的详细错误信息
	//
	// example:
	//
	// 参数错误
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// SSE事件类型，如：result-generated(生成结果)、task-finished(任务完成)
	//
	// example:
	//
	// result-generated
	Event *string `json:"Event,omitempty" xml:"Event,omitempty"`
	// 当前写作会话的唯一标识
	//
	// example:
	//
	// session-xxxxx
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// HTTP状态码
	//
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	// 写作任务的唯一标识
	//
	// example:
	//
	// task-xxxxx
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// 用于问题排查的链路追踪标识
	//
	// example:
	//
	// 1813ceee-7fe5-41b4-87e5-982a4d18cca5
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s RunAiHelperWritingResponseBodyHeader) String() string {
	return dara.Prettify(s)
}

func (s RunAiHelperWritingResponseBodyHeader) GoString() string {
	return s.String()
}

func (s *RunAiHelperWritingResponseBodyHeader) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *RunAiHelperWritingResponseBodyHeader) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *RunAiHelperWritingResponseBodyHeader) GetEvent() *string {
	return s.Event
}

func (s *RunAiHelperWritingResponseBodyHeader) GetSessionId() *string {
	return s.SessionId
}

func (s *RunAiHelperWritingResponseBodyHeader) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunAiHelperWritingResponseBodyHeader) GetTaskId() *string {
	return s.TaskId
}

func (s *RunAiHelperWritingResponseBodyHeader) GetTraceId() *string {
	return s.TraceId
}

func (s *RunAiHelperWritingResponseBodyHeader) SetErrorCode(v string) *RunAiHelperWritingResponseBodyHeader {
	s.ErrorCode = &v
	return s
}

func (s *RunAiHelperWritingResponseBodyHeader) SetErrorMessage(v string) *RunAiHelperWritingResponseBodyHeader {
	s.ErrorMessage = &v
	return s
}

func (s *RunAiHelperWritingResponseBodyHeader) SetEvent(v string) *RunAiHelperWritingResponseBodyHeader {
	s.Event = &v
	return s
}

func (s *RunAiHelperWritingResponseBodyHeader) SetSessionId(v string) *RunAiHelperWritingResponseBodyHeader {
	s.SessionId = &v
	return s
}

func (s *RunAiHelperWritingResponseBodyHeader) SetStatusCode(v int32) *RunAiHelperWritingResponseBodyHeader {
	s.StatusCode = &v
	return s
}

func (s *RunAiHelperWritingResponseBodyHeader) SetTaskId(v string) *RunAiHelperWritingResponseBodyHeader {
	s.TaskId = &v
	return s
}

func (s *RunAiHelperWritingResponseBodyHeader) SetTraceId(v string) *RunAiHelperWritingResponseBodyHeader {
	s.TraceId = &v
	return s
}

func (s *RunAiHelperWritingResponseBodyHeader) Validate() error {
	return dara.Validate(s)
}

type RunAiHelperWritingResponseBodyPayload struct {
	// AI生成的写作内容
	Output *RunAiHelperWritingResponseBodyPayloadOutput `json:"Output,omitempty" xml:"Output,omitempty" type:"Struct"`
	// 本次请求的Token消耗统计
	Usage *RunAiHelperWritingResponseBodyPayloadUsage `json:"Usage,omitempty" xml:"Usage,omitempty" type:"Struct"`
}

func (s RunAiHelperWritingResponseBodyPayload) String() string {
	return dara.Prettify(s)
}

func (s RunAiHelperWritingResponseBodyPayload) GoString() string {
	return s.String()
}

func (s *RunAiHelperWritingResponseBodyPayload) GetOutput() *RunAiHelperWritingResponseBodyPayloadOutput {
	return s.Output
}

func (s *RunAiHelperWritingResponseBodyPayload) GetUsage() *RunAiHelperWritingResponseBodyPayloadUsage {
	return s.Usage
}

func (s *RunAiHelperWritingResponseBodyPayload) SetOutput(v *RunAiHelperWritingResponseBodyPayloadOutput) *RunAiHelperWritingResponseBodyPayload {
	s.Output = v
	return s
}

func (s *RunAiHelperWritingResponseBodyPayload) SetUsage(v *RunAiHelperWritingResponseBodyPayloadUsage) *RunAiHelperWritingResponseBodyPayload {
	s.Usage = v
	return s
}

func (s *RunAiHelperWritingResponseBodyPayload) Validate() error {
	if s.Output != nil {
		if err := s.Output.Validate(); err != nil {
			return err
		}
	}
	if s.Usage != nil {
		if err := s.Usage.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RunAiHelperWritingResponseBodyPayloadOutput struct {
	// AI生成的文章内容，流式返回时为增量文本
	//
	// example:
	//
	// 人工智能正在深刻改变我们的生活...
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// 返回的写作参数键值对
	//
	// example:
	//
	// {"wordCount": "1000"}
	WritingParams map[string]*string `json:"WritingParams,omitempty" xml:"WritingParams,omitempty"`
}

func (s RunAiHelperWritingResponseBodyPayloadOutput) String() string {
	return dara.Prettify(s)
}

func (s RunAiHelperWritingResponseBodyPayloadOutput) GoString() string {
	return s.String()
}

func (s *RunAiHelperWritingResponseBodyPayloadOutput) GetText() *string {
	return s.Text
}

func (s *RunAiHelperWritingResponseBodyPayloadOutput) GetWritingParams() map[string]*string {
	return s.WritingParams
}

func (s *RunAiHelperWritingResponseBodyPayloadOutput) SetText(v string) *RunAiHelperWritingResponseBodyPayloadOutput {
	s.Text = &v
	return s
}

func (s *RunAiHelperWritingResponseBodyPayloadOutput) SetWritingParams(v map[string]*string) *RunAiHelperWritingResponseBodyPayloadOutput {
	s.WritingParams = v
	return s
}

func (s *RunAiHelperWritingResponseBodyPayloadOutput) Validate() error {
	return dara.Validate(s)
}

type RunAiHelperWritingResponseBodyPayloadUsage struct {
	// 输入内容消耗的Token数量
	//
	// example:
	//
	// 256
	InputTokens *int64 `json:"InputTokens,omitempty" xml:"InputTokens,omitempty"`
	// 生成内容消耗的Token数量
	//
	// example:
	//
	// 1024
	OutputTokens *int64 `json:"OutputTokens,omitempty" xml:"OutputTokens,omitempty"`
	// 输入和输出Token的总和
	//
	// example:
	//
	// 1280
	TotalTokens *int64 `json:"TotalTokens,omitempty" xml:"TotalTokens,omitempty"`
}

func (s RunAiHelperWritingResponseBodyPayloadUsage) String() string {
	return dara.Prettify(s)
}

func (s RunAiHelperWritingResponseBodyPayloadUsage) GoString() string {
	return s.String()
}

func (s *RunAiHelperWritingResponseBodyPayloadUsage) GetInputTokens() *int64 {
	return s.InputTokens
}

func (s *RunAiHelperWritingResponseBodyPayloadUsage) GetOutputTokens() *int64 {
	return s.OutputTokens
}

func (s *RunAiHelperWritingResponseBodyPayloadUsage) GetTotalTokens() *int64 {
	return s.TotalTokens
}

func (s *RunAiHelperWritingResponseBodyPayloadUsage) SetInputTokens(v int64) *RunAiHelperWritingResponseBodyPayloadUsage {
	s.InputTokens = &v
	return s
}

func (s *RunAiHelperWritingResponseBodyPayloadUsage) SetOutputTokens(v int64) *RunAiHelperWritingResponseBodyPayloadUsage {
	s.OutputTokens = &v
	return s
}

func (s *RunAiHelperWritingResponseBodyPayloadUsage) SetTotalTokens(v int64) *RunAiHelperWritingResponseBodyPayloadUsage {
	s.TotalTokens = &v
	return s
}

func (s *RunAiHelperWritingResponseBodyPayloadUsage) Validate() error {
	return dara.Validate(s)
}
