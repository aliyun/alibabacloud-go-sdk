// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunVideoScriptGenerateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RunVideoScriptGenerateResponseBody
	GetCode() *string
	SetHeader(v *RunVideoScriptGenerateResponseBodyHeader) *RunVideoScriptGenerateResponseBody
	GetHeader() *RunVideoScriptGenerateResponseBodyHeader
	SetHttpStatusCode(v string) *RunVideoScriptGenerateResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *RunVideoScriptGenerateResponseBody
	GetMessage() *string
	SetPayload(v *RunVideoScriptGenerateResponseBodyPayload) *RunVideoScriptGenerateResponseBody
	GetPayload() *RunVideoScriptGenerateResponseBodyPayload
	SetRequestId(v string) *RunVideoScriptGenerateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RunVideoScriptGenerateResponseBody
	GetSuccess() *bool
}

type RunVideoScriptGenerateResponseBody struct {
	// The status code. A value of 200 indicates a normal response. This field is returned when the \\`Content-Type\\` is \\`json\\`.
	//
	// example:
	//
	// NoPermission
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response header.
	Header *RunVideoScriptGenerateResponseBodyHeader `json:"Header,omitempty" xml:"Header,omitempty" type:"Struct"`
	// The HTTP status code. This field is returned when the \\`Content-Type\\` is \\`json\\`.
	//
	// example:
	//
	// 403
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error description. This field is returned when the \\`Content-Type\\` is \\`json\\`.
	//
	// example:
	//
	// You are not authorized to perform this action.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The response body.
	Payload *RunVideoScriptGenerateResponseBodyPayload `json:"Payload,omitempty" xml:"Payload,omitempty" type:"Struct"`
	// The unique ID of the request. This field is returned when the \\`Content-Type\\` is \\`json\\`.
	//
	// example:
	//
	// F2F366D6-E9FE-1006-BB70-2C650896AAB5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. \\`true\\` indicates success. \\`false\\` indicates failure. This field is returned when the \\`Content-Type\\` is \\`json\\`.
	//
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RunVideoScriptGenerateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RunVideoScriptGenerateResponseBody) GoString() string {
	return s.String()
}

func (s *RunVideoScriptGenerateResponseBody) GetCode() *string {
	return s.Code
}

func (s *RunVideoScriptGenerateResponseBody) GetHeader() *RunVideoScriptGenerateResponseBodyHeader {
	return s.Header
}

func (s *RunVideoScriptGenerateResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *RunVideoScriptGenerateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RunVideoScriptGenerateResponseBody) GetPayload() *RunVideoScriptGenerateResponseBodyPayload {
	return s.Payload
}

func (s *RunVideoScriptGenerateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RunVideoScriptGenerateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RunVideoScriptGenerateResponseBody) SetCode(v string) *RunVideoScriptGenerateResponseBody {
	s.Code = &v
	return s
}

func (s *RunVideoScriptGenerateResponseBody) SetHeader(v *RunVideoScriptGenerateResponseBodyHeader) *RunVideoScriptGenerateResponseBody {
	s.Header = v
	return s
}

func (s *RunVideoScriptGenerateResponseBody) SetHttpStatusCode(v string) *RunVideoScriptGenerateResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *RunVideoScriptGenerateResponseBody) SetMessage(v string) *RunVideoScriptGenerateResponseBody {
	s.Message = &v
	return s
}

func (s *RunVideoScriptGenerateResponseBody) SetPayload(v *RunVideoScriptGenerateResponseBodyPayload) *RunVideoScriptGenerateResponseBody {
	s.Payload = v
	return s
}

func (s *RunVideoScriptGenerateResponseBody) SetRequestId(v string) *RunVideoScriptGenerateResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunVideoScriptGenerateResponseBody) SetSuccess(v bool) *RunVideoScriptGenerateResponseBody {
	s.Success = &v
	return s
}

func (s *RunVideoScriptGenerateResponseBody) Validate() error {
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

type RunVideoScriptGenerateResponseBodyHeader struct {
	// The error code.
	//
	// example:
	//
	// ScriptNumberExceed
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned when the call fails.
	//
	// example:
	//
	// 脚本篇数超限
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The event name.
	//
	// example:
	//
	// result-generated
	Event *string `json:"Event,omitempty" xml:"Event,omitempty"`
	// The session ID.
	//
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// The status code.
	//
	// example:
	//
	// 400
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	// The task ID.
	//
	// example:
	//
	// b057f2fa-2277-477b-babf-cbc062307828
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The trace ID.
	//
	// example:
	//
	// 2150451a17191950923411783e2927
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s RunVideoScriptGenerateResponseBodyHeader) String() string {
	return dara.Prettify(s)
}

func (s RunVideoScriptGenerateResponseBodyHeader) GoString() string {
	return s.String()
}

func (s *RunVideoScriptGenerateResponseBodyHeader) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *RunVideoScriptGenerateResponseBodyHeader) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *RunVideoScriptGenerateResponseBodyHeader) GetEvent() *string {
	return s.Event
}

func (s *RunVideoScriptGenerateResponseBodyHeader) GetSessionId() *string {
	return s.SessionId
}

func (s *RunVideoScriptGenerateResponseBodyHeader) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunVideoScriptGenerateResponseBodyHeader) GetTaskId() *string {
	return s.TaskId
}

func (s *RunVideoScriptGenerateResponseBodyHeader) GetTraceId() *string {
	return s.TraceId
}

func (s *RunVideoScriptGenerateResponseBodyHeader) SetErrorCode(v string) *RunVideoScriptGenerateResponseBodyHeader {
	s.ErrorCode = &v
	return s
}

func (s *RunVideoScriptGenerateResponseBodyHeader) SetErrorMessage(v string) *RunVideoScriptGenerateResponseBodyHeader {
	s.ErrorMessage = &v
	return s
}

func (s *RunVideoScriptGenerateResponseBodyHeader) SetEvent(v string) *RunVideoScriptGenerateResponseBodyHeader {
	s.Event = &v
	return s
}

func (s *RunVideoScriptGenerateResponseBodyHeader) SetSessionId(v string) *RunVideoScriptGenerateResponseBodyHeader {
	s.SessionId = &v
	return s
}

func (s *RunVideoScriptGenerateResponseBodyHeader) SetStatusCode(v int32) *RunVideoScriptGenerateResponseBodyHeader {
	s.StatusCode = &v
	return s
}

func (s *RunVideoScriptGenerateResponseBodyHeader) SetTaskId(v string) *RunVideoScriptGenerateResponseBodyHeader {
	s.TaskId = &v
	return s
}

func (s *RunVideoScriptGenerateResponseBodyHeader) SetTraceId(v string) *RunVideoScriptGenerateResponseBodyHeader {
	s.TraceId = &v
	return s
}

func (s *RunVideoScriptGenerateResponseBodyHeader) Validate() error {
	return dara.Validate(s)
}

type RunVideoScriptGenerateResponseBodyPayload struct {
	// The output content object.
	Output *RunVideoScriptGenerateResponseBodyPayloadOutput `json:"Output,omitempty" xml:"Output,omitempty" type:"Struct"`
	// The token usage.
	Usage *RunVideoScriptGenerateResponseBodyPayloadUsage `json:"Usage,omitempty" xml:"Usage,omitempty" type:"Struct"`
}

func (s RunVideoScriptGenerateResponseBodyPayload) String() string {
	return dara.Prettify(s)
}

func (s RunVideoScriptGenerateResponseBodyPayload) GoString() string {
	return s.String()
}

func (s *RunVideoScriptGenerateResponseBodyPayload) GetOutput() *RunVideoScriptGenerateResponseBodyPayloadOutput {
	return s.Output
}

func (s *RunVideoScriptGenerateResponseBodyPayload) GetUsage() *RunVideoScriptGenerateResponseBodyPayloadUsage {
	return s.Usage
}

func (s *RunVideoScriptGenerateResponseBodyPayload) SetOutput(v *RunVideoScriptGenerateResponseBodyPayloadOutput) *RunVideoScriptGenerateResponseBodyPayload {
	s.Output = v
	return s
}

func (s *RunVideoScriptGenerateResponseBodyPayload) SetUsage(v *RunVideoScriptGenerateResponseBodyPayloadUsage) *RunVideoScriptGenerateResponseBodyPayload {
	s.Usage = v
	return s
}

func (s *RunVideoScriptGenerateResponseBodyPayload) Validate() error {
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

type RunVideoScriptGenerateResponseBodyPayloadOutput struct {
	// The text generation result.
	//
	// example:
	//
	// 大家好，我是[xxx]。今天带大家走进黄山，感受奇松、怪石、云海、温泉的绝美风光。首站迎客松，800年历史，枝干如臂，热情迎接每一位游客。接着登光明顶，360度全景尽收眼底。再探秘西海大峡谷，体验原始自然的震撼。最后，在温泉中放松身心，享受旅途的美好。希望这次旅行能给你留下难忘的记忆。我是[你的名字]，感谢观看，我们下次再见！","91522b25a4f440189320c9ede8ae6c85":"大家好，我是[您的名字]，今天带大家探索黄山的奇妙之旅。首先，我们将见到黄山的象征——迎客松，感受它800年的历史与欢迎。随后攀登光明顶，迎接壮丽的日出；漫步西海大峡谷，体验险峻之美；最后，在温泉中放松身心。希望这次旅行能让你爱上黄山。谢谢观看！","1c23af4a899e4b908bdcffa7d8d0ddc9":"大家好，欢迎来到黄山！这里以奇松、怪石、云海、温泉四绝闻名。从云谷寺开始，感受古朴氛围；挑战百步云梯，体验攀登乐趣；漫步西海大峡谷，领略壮丽景色；最后在玉屏楼迎接日出，享受心灵的宁静。希望这次旅行给你留下美好回忆！
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RunVideoScriptGenerateResponseBodyPayloadOutput) String() string {
	return dara.Prettify(s)
}

func (s RunVideoScriptGenerateResponseBodyPayloadOutput) GoString() string {
	return s.String()
}

func (s *RunVideoScriptGenerateResponseBodyPayloadOutput) GetText() *string {
	return s.Text
}

func (s *RunVideoScriptGenerateResponseBodyPayloadOutput) SetText(v string) *RunVideoScriptGenerateResponseBodyPayloadOutput {
	s.Text = &v
	return s
}

func (s *RunVideoScriptGenerateResponseBodyPayloadOutput) Validate() error {
	return dara.Validate(s)
}

type RunVideoScriptGenerateResponseBodyPayloadUsage struct {
	// The number of tokens used for the input.
	//
	// example:
	//
	// 100
	InputTokens *int64 `json:"InputTokens,omitempty" xml:"InputTokens,omitempty"`
	// The number of tokens for the output.
	//
	// example:
	//
	// 100
	OutputTokens *int64 `json:"OutputTokens,omitempty" xml:"OutputTokens,omitempty"`
	// The total number of tokens.
	//
	// example:
	//
	// 200
	TotalTokens *int64 `json:"TotalTokens,omitempty" xml:"TotalTokens,omitempty"`
}

func (s RunVideoScriptGenerateResponseBodyPayloadUsage) String() string {
	return dara.Prettify(s)
}

func (s RunVideoScriptGenerateResponseBodyPayloadUsage) GoString() string {
	return s.String()
}

func (s *RunVideoScriptGenerateResponseBodyPayloadUsage) GetInputTokens() *int64 {
	return s.InputTokens
}

func (s *RunVideoScriptGenerateResponseBodyPayloadUsage) GetOutputTokens() *int64 {
	return s.OutputTokens
}

func (s *RunVideoScriptGenerateResponseBodyPayloadUsage) GetTotalTokens() *int64 {
	return s.TotalTokens
}

func (s *RunVideoScriptGenerateResponseBodyPayloadUsage) SetInputTokens(v int64) *RunVideoScriptGenerateResponseBodyPayloadUsage {
	s.InputTokens = &v
	return s
}

func (s *RunVideoScriptGenerateResponseBodyPayloadUsage) SetOutputTokens(v int64) *RunVideoScriptGenerateResponseBodyPayloadUsage {
	s.OutputTokens = &v
	return s
}

func (s *RunVideoScriptGenerateResponseBodyPayloadUsage) SetTotalTokens(v int64) *RunVideoScriptGenerateResponseBodyPayloadUsage {
	s.TotalTokens = &v
	return s
}

func (s *RunVideoScriptGenerateResponseBodyPayloadUsage) Validate() error {
	return dara.Validate(s)
}
