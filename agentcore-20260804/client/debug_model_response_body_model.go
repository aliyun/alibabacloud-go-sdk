// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDebugModelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DebugModelResponseBody
	GetCode() *string
	SetData(v *DebugModelResponseBodyData) *DebugModelResponseBody
	GetData() *DebugModelResponseBodyData
	SetHttpStatusCode(v int32) *DebugModelResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DebugModelResponseBody
	GetMessage() *string
	SetRequestId(v string) *DebugModelResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DebugModelResponseBody
	GetSuccess() *bool
}

type DebugModelResponseBody struct {
	// example:
	//
	// SUCCESS
	Code *string                     `json:"code,omitempty" xml:"code,omitempty"`
	Data *DebugModelResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// request-1
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DebugModelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DebugModelResponseBody) GoString() string {
	return s.String()
}

func (s *DebugModelResponseBody) GetCode() *string {
	return s.Code
}

func (s *DebugModelResponseBody) GetData() *DebugModelResponseBodyData {
	return s.Data
}

func (s *DebugModelResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DebugModelResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DebugModelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DebugModelResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DebugModelResponseBody) SetCode(v string) *DebugModelResponseBody {
	s.Code = &v
	return s
}

func (s *DebugModelResponseBody) SetData(v *DebugModelResponseBodyData) *DebugModelResponseBody {
	s.Data = v
	return s
}

func (s *DebugModelResponseBody) SetHttpStatusCode(v int32) *DebugModelResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DebugModelResponseBody) SetMessage(v string) *DebugModelResponseBody {
	s.Message = &v
	return s
}

func (s *DebugModelResponseBody) SetRequestId(v string) *DebugModelResponseBody {
	s.RequestId = &v
	return s
}

func (s *DebugModelResponseBody) SetSuccess(v bool) *DebugModelResponseBody {
	s.Success = &v
	return s
}

func (s *DebugModelResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DebugModelResponseBodyData struct {
	// example:
	//
	// mc-1
	ConnectionId *string `json:"connectionId,omitempty" xml:"connectionId,omitempty"`
	DebugSuccess *bool   `json:"debugSuccess,omitempty" xml:"debugSuccess,omitempty"`
	// 调试失败时的错误码。取值：MODEL_CONNECTION_NOT_READY（模型连接尚未发布就绪）、MODEL_CONNECTION_TEST_FAILED（平台调用网关失败）、UPSTREAM_MODEL_NOT_FOUND（模型服务商侧不存在该模型）、UPSTREAM_UNAUTHORIZED（模型服务商拒绝所配置的凭证）、UPSTREAM_RATE_LIMITED（模型服务商限流）、UPSTREAM_SERVER_ERROR（模型服务商服务端错误）、UPSTREAM_HTTP_ERROR（模型服务商返回其它非成功状态）、UPSTREAM_EMPTY_RESPONSE（模型服务商返回空响应）、UPSTREAM_INVALID_RESPONSE（模型服务商响应格式非法）、UPSTREAM_MODEL_ERROR（模型服务商拒绝本次请求）、MODEL_RESPONSE_INVALID（响应解析失败）。
	//
	// example:
	//
	// UPSTREAM_MODEL_ERROR
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 调试失败时的错误描述，为固定脱敏文案，不透传模型服务商的原始错误详情。
	//
	// example:
	//
	// The model endpoint rejected the debug request.
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// 3
	InputTokens *int64 `json:"inputTokens,omitempty" xml:"inputTokens,omitempty"`
	// example:
	//
	// 12
	LatencyMs *int64 `json:"latencyMs,omitempty" xml:"latencyMs,omitempty"`
	// example:
	//
	// model-1
	ModelId *string `json:"modelId,omitempty" xml:"modelId,omitempty"`
	// example:
	//
	// 2
	OutputTokens *int64 `json:"outputTokens,omitempty" xml:"outputTokens,omitempty"`
	// example:
	//
	// ok
	Response *string `json:"response,omitempty" xml:"response,omitempty"`
	// 调试结果状态。取值：NORMAL（正常）、ABNORMAL（异常）。
	//
	// example:
	//
	// NORMAL
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s DebugModelResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DebugModelResponseBodyData) GoString() string {
	return s.String()
}

func (s *DebugModelResponseBodyData) GetConnectionId() *string {
	return s.ConnectionId
}

func (s *DebugModelResponseBodyData) GetDebugSuccess() *bool {
	return s.DebugSuccess
}

func (s *DebugModelResponseBodyData) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DebugModelResponseBodyData) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DebugModelResponseBodyData) GetInputTokens() *int64 {
	return s.InputTokens
}

func (s *DebugModelResponseBodyData) GetLatencyMs() *int64 {
	return s.LatencyMs
}

func (s *DebugModelResponseBodyData) GetModelId() *string {
	return s.ModelId
}

func (s *DebugModelResponseBodyData) GetOutputTokens() *int64 {
	return s.OutputTokens
}

func (s *DebugModelResponseBodyData) GetResponse() *string {
	return s.Response
}

func (s *DebugModelResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *DebugModelResponseBodyData) SetConnectionId(v string) *DebugModelResponseBodyData {
	s.ConnectionId = &v
	return s
}

func (s *DebugModelResponseBodyData) SetDebugSuccess(v bool) *DebugModelResponseBodyData {
	s.DebugSuccess = &v
	return s
}

func (s *DebugModelResponseBodyData) SetErrorCode(v string) *DebugModelResponseBodyData {
	s.ErrorCode = &v
	return s
}

func (s *DebugModelResponseBodyData) SetErrorMessage(v string) *DebugModelResponseBodyData {
	s.ErrorMessage = &v
	return s
}

func (s *DebugModelResponseBodyData) SetInputTokens(v int64) *DebugModelResponseBodyData {
	s.InputTokens = &v
	return s
}

func (s *DebugModelResponseBodyData) SetLatencyMs(v int64) *DebugModelResponseBodyData {
	s.LatencyMs = &v
	return s
}

func (s *DebugModelResponseBodyData) SetModelId(v string) *DebugModelResponseBodyData {
	s.ModelId = &v
	return s
}

func (s *DebugModelResponseBodyData) SetOutputTokens(v int64) *DebugModelResponseBodyData {
	s.OutputTokens = &v
	return s
}

func (s *DebugModelResponseBodyData) SetResponse(v string) *DebugModelResponseBodyData {
	s.Response = &v
	return s
}

func (s *DebugModelResponseBodyData) SetStatus(v string) *DebugModelResponseBodyData {
	s.Status = &v
	return s
}

func (s *DebugModelResponseBodyData) Validate() error {
	return dara.Validate(s)
}
