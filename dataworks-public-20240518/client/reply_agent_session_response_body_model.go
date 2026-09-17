// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReplyAgentSessionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJsonRpcResponse(v *ReplyAgentSessionResponseBodyJsonRpcResponse) *ReplyAgentSessionResponseBody
	GetJsonRpcResponse() *ReplyAgentSessionResponseBodyJsonRpcResponse
	SetRequestId(v string) *ReplyAgentSessionResponseBody
	GetRequestId() *string
}

type ReplyAgentSessionResponseBody struct {
	// The JSON-RPC response. Returns Result on success or Error on protocol errors.
	JsonRpcResponse *ReplyAgentSessionResponseBodyJsonRpcResponse `json:"JsonRpcResponse,omitempty" xml:"JsonRpcResponse,omitempty" type:"Struct"`
	// The request ID for this call, which can be used for troubleshooting.
	//
	// example:
	//
	// request-001
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReplyAgentSessionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReplyAgentSessionResponseBody) GoString() string {
	return s.String()
}

func (s *ReplyAgentSessionResponseBody) GetJsonRpcResponse() *ReplyAgentSessionResponseBodyJsonRpcResponse {
	return s.JsonRpcResponse
}

func (s *ReplyAgentSessionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReplyAgentSessionResponseBody) SetJsonRpcResponse(v *ReplyAgentSessionResponseBodyJsonRpcResponse) *ReplyAgentSessionResponseBody {
	s.JsonRpcResponse = v
	return s
}

func (s *ReplyAgentSessionResponseBody) SetRequestId(v string) *ReplyAgentSessionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReplyAgentSessionResponseBody) Validate() error {
	if s.JsonRpcResponse != nil {
		if err := s.JsonRpcResponse.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ReplyAgentSessionResponseBodyJsonRpcResponse struct {
	// The JSON-RPC fault information. For example, DAEMON_PERMISSION_UNAVAILABLE is returned when the daemon reply feature is not enabled.
	Error *ReplyAgentSessionResponseBodyJsonRpcResponseError `json:"Error,omitempty" xml:"Error,omitempty" type:"Struct"`
	// The JSON-RPC correlation ID for this reply request.
	//
	// example:
	//
	// reply-rpc-001
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The JSON-RPC protocol version.
	//
	// example:
	//
	// 2.0
	Jsonrpc *string `json:"Jsonrpc,omitempty" xml:"Jsonrpc,omitempty"`
	// The reply processing result. This only indicates whether the reply was accepted, not whether the original task has completed.
	Result *ReplyAgentSessionResponseBodyJsonRpcResponseResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// The response time. This is a UNIX timestamp, in milliseconds.
	//
	// example:
	//
	// 1789549200000
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s ReplyAgentSessionResponseBodyJsonRpcResponse) String() string {
	return dara.Prettify(s)
}

func (s ReplyAgentSessionResponseBodyJsonRpcResponse) GoString() string {
	return s.String()
}

func (s *ReplyAgentSessionResponseBodyJsonRpcResponse) GetError() *ReplyAgentSessionResponseBodyJsonRpcResponseError {
	return s.Error
}

func (s *ReplyAgentSessionResponseBodyJsonRpcResponse) GetId() *string {
	return s.Id
}

func (s *ReplyAgentSessionResponseBodyJsonRpcResponse) GetJsonrpc() *string {
	return s.Jsonrpc
}

func (s *ReplyAgentSessionResponseBodyJsonRpcResponse) GetResult() *ReplyAgentSessionResponseBodyJsonRpcResponseResult {
	return s.Result
}

func (s *ReplyAgentSessionResponseBodyJsonRpcResponse) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *ReplyAgentSessionResponseBodyJsonRpcResponse) SetError(v *ReplyAgentSessionResponseBodyJsonRpcResponseError) *ReplyAgentSessionResponseBodyJsonRpcResponse {
	s.Error = v
	return s
}

func (s *ReplyAgentSessionResponseBodyJsonRpcResponse) SetId(v string) *ReplyAgentSessionResponseBodyJsonRpcResponse {
	s.Id = &v
	return s
}

func (s *ReplyAgentSessionResponseBodyJsonRpcResponse) SetJsonrpc(v string) *ReplyAgentSessionResponseBodyJsonRpcResponse {
	s.Jsonrpc = &v
	return s
}

func (s *ReplyAgentSessionResponseBodyJsonRpcResponse) SetResult(v *ReplyAgentSessionResponseBodyJsonRpcResponseResult) *ReplyAgentSessionResponseBodyJsonRpcResponse {
	s.Result = v
	return s
}

func (s *ReplyAgentSessionResponseBodyJsonRpcResponse) SetTimestamp(v int64) *ReplyAgentSessionResponseBodyJsonRpcResponse {
	s.Timestamp = &v
	return s
}

func (s *ReplyAgentSessionResponseBodyJsonRpcResponse) Validate() error {
	if s.Error != nil {
		if err := s.Error.Validate(); err != nil {
			return err
		}
	}
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ReplyAgentSessionResponseBodyJsonRpcResponseError struct {
	// The JSON-RPC error code.
	//
	// example:
	//
	// -32601
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The optional additional error information. The content depends on the error type.
	//
	// example:
	//
	// {}
	Data interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// The business error code.
	//
	// example:
	//
	// DAEMON_PERMISSION_UNAVAILABLE
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// DataAgent daemon permission reply is not enabled
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s ReplyAgentSessionResponseBodyJsonRpcResponseError) String() string {
	return dara.Prettify(s)
}

func (s ReplyAgentSessionResponseBodyJsonRpcResponseError) GoString() string {
	return s.String()
}

func (s *ReplyAgentSessionResponseBodyJsonRpcResponseError) GetCode() *int32 {
	return s.Code
}

func (s *ReplyAgentSessionResponseBodyJsonRpcResponseError) GetData() interface{} {
	return s.Data
}

func (s *ReplyAgentSessionResponseBodyJsonRpcResponseError) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ReplyAgentSessionResponseBodyJsonRpcResponseError) GetMessage() *string {
	return s.Message
}

func (s *ReplyAgentSessionResponseBodyJsonRpcResponseError) SetCode(v int32) *ReplyAgentSessionResponseBodyJsonRpcResponseError {
	s.Code = &v
	return s
}

func (s *ReplyAgentSessionResponseBodyJsonRpcResponseError) SetData(v interface{}) *ReplyAgentSessionResponseBodyJsonRpcResponseError {
	s.Data = v
	return s
}

func (s *ReplyAgentSessionResponseBodyJsonRpcResponseError) SetErrorCode(v string) *ReplyAgentSessionResponseBodyJsonRpcResponseError {
	s.ErrorCode = &v
	return s
}

func (s *ReplyAgentSessionResponseBodyJsonRpcResponseError) SetMessage(v string) *ReplyAgentSessionResponseBodyJsonRpcResponseError {
	s.Message = &v
	return s
}

func (s *ReplyAgentSessionResponseBodyJsonRpcResponseError) Validate() error {
	return dara.Validate(s)
}

type ReplyAgentSessionResponseBodyJsonRpcResponseResult struct {
	// Indicates whether the daemon accepted the reply. A value of true indicates that the daemon accepted the reply. A value of false indicates that the reply was not accepted. Possible reasons include an unknown request, an already processed request, an expired request, or a nonexistent session. You cannot determine the specific reason from this value.
	//
	// example:
	//
	// true
	Accepted *bool `json:"Accepted,omitempty" xml:"Accepted,omitempty"`
}

func (s ReplyAgentSessionResponseBodyJsonRpcResponseResult) String() string {
	return dara.Prettify(s)
}

func (s ReplyAgentSessionResponseBodyJsonRpcResponseResult) GoString() string {
	return s.String()
}

func (s *ReplyAgentSessionResponseBodyJsonRpcResponseResult) GetAccepted() *bool {
	return s.Accepted
}

func (s *ReplyAgentSessionResponseBodyJsonRpcResponseResult) SetAccepted(v bool) *ReplyAgentSessionResponseBodyJsonRpcResponseResult {
	s.Accepted = &v
	return s
}

func (s *ReplyAgentSessionResponseBodyJsonRpcResponseResult) Validate() error {
	return dara.Validate(s)
}
