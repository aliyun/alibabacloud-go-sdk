// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAgentSessionTitleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJsonRpcResponse(v *UpdateAgentSessionTitleResponseBodyJsonRpcResponse) *UpdateAgentSessionTitleResponseBody
	GetJsonRpcResponse() *UpdateAgentSessionTitleResponseBodyJsonRpcResponse
	SetRequestId(v string) *UpdateAgentSessionTitleResponseBody
	GetRequestId() *string
}

type UpdateAgentSessionTitleResponseBody struct {
	// The JSON-RPC response. Result is returned upon success, and Error is returned upon a protocol error.
	JsonRpcResponse *UpdateAgentSessionTitleResponseBodyJsonRpcResponse `json:"JsonRpcResponse,omitempty" xml:"JsonRpcResponse,omitempty" type:"Struct"`
	// The unique ID of the current OpenAPI request. You can use it to troubleshoot request issues.
	//
	// example:
	//
	// 01A0B6E0-E005-1465-B5DE-5CD17677426F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateAgentSessionTitleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAgentSessionTitleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAgentSessionTitleResponseBody) GetJsonRpcResponse() *UpdateAgentSessionTitleResponseBodyJsonRpcResponse {
	return s.JsonRpcResponse
}

func (s *UpdateAgentSessionTitleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAgentSessionTitleResponseBody) SetJsonRpcResponse(v *UpdateAgentSessionTitleResponseBodyJsonRpcResponse) *UpdateAgentSessionTitleResponseBody {
	s.JsonRpcResponse = v
	return s
}

func (s *UpdateAgentSessionTitleResponseBody) SetRequestId(v string) *UpdateAgentSessionTitleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAgentSessionTitleResponseBody) Validate() error {
	if s.JsonRpcResponse != nil {
		if err := s.JsonRpcResponse.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateAgentSessionTitleResponseBodyJsonRpcResponse struct {
	// The JSON-RPC error message. This field is not returned upon success.
	Error *UpdateAgentSessionTitleResponseBodyJsonRpcResponseError `json:"Error,omitempty" xml:"Error,omitempty" type:"Struct"`
	// The JSON-RPC correlation ID for the current request. It is passed by the caller and returned as is.
	//
	// example:
	//
	// updateagentsessiontitle-rpc-001
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The JSON-RPC protocol version.
	//
	// example:
	//
	// 2.0
	Jsonrpc *string `json:"Jsonrpc,omitempty" xml:"Jsonrpc,omitempty"`
	// The operation result.
	Result *UpdateAgentSessionTitleResponseBodyJsonRpcResponseResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// The response time. It is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1789549200000
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s UpdateAgentSessionTitleResponseBodyJsonRpcResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAgentSessionTitleResponseBodyJsonRpcResponse) GoString() string {
	return s.String()
}

func (s *UpdateAgentSessionTitleResponseBodyJsonRpcResponse) GetError() *UpdateAgentSessionTitleResponseBodyJsonRpcResponseError {
	return s.Error
}

func (s *UpdateAgentSessionTitleResponseBodyJsonRpcResponse) GetId() *string {
	return s.Id
}

func (s *UpdateAgentSessionTitleResponseBodyJsonRpcResponse) GetJsonrpc() *string {
	return s.Jsonrpc
}

func (s *UpdateAgentSessionTitleResponseBodyJsonRpcResponse) GetResult() *UpdateAgentSessionTitleResponseBodyJsonRpcResponseResult {
	return s.Result
}

func (s *UpdateAgentSessionTitleResponseBodyJsonRpcResponse) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *UpdateAgentSessionTitleResponseBodyJsonRpcResponse) SetError(v *UpdateAgentSessionTitleResponseBodyJsonRpcResponseError) *UpdateAgentSessionTitleResponseBodyJsonRpcResponse {
	s.Error = v
	return s
}

func (s *UpdateAgentSessionTitleResponseBodyJsonRpcResponse) SetId(v string) *UpdateAgentSessionTitleResponseBodyJsonRpcResponse {
	s.Id = &v
	return s
}

func (s *UpdateAgentSessionTitleResponseBodyJsonRpcResponse) SetJsonrpc(v string) *UpdateAgentSessionTitleResponseBodyJsonRpcResponse {
	s.Jsonrpc = &v
	return s
}

func (s *UpdateAgentSessionTitleResponseBodyJsonRpcResponse) SetResult(v *UpdateAgentSessionTitleResponseBodyJsonRpcResponseResult) *UpdateAgentSessionTitleResponseBodyJsonRpcResponse {
	s.Result = v
	return s
}

func (s *UpdateAgentSessionTitleResponseBodyJsonRpcResponse) SetTimestamp(v int64) *UpdateAgentSessionTitleResponseBodyJsonRpcResponse {
	s.Timestamp = &v
	return s
}

func (s *UpdateAgentSessionTitleResponseBodyJsonRpcResponse) Validate() error {
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

type UpdateAgentSessionTitleResponseBodyJsonRpcResponseError struct {
	// The JSON-RPC error code.
	//
	// example:
	//
	// 500
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
	// 0x4883300000000019
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// The session does not exist or the current account is not authorized to access it. Return to the session list and select another session
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s UpdateAgentSessionTitleResponseBodyJsonRpcResponseError) String() string {
	return dara.Prettify(s)
}

func (s UpdateAgentSessionTitleResponseBodyJsonRpcResponseError) GoString() string {
	return s.String()
}

func (s *UpdateAgentSessionTitleResponseBodyJsonRpcResponseError) GetCode() *int32 {
	return s.Code
}

func (s *UpdateAgentSessionTitleResponseBodyJsonRpcResponseError) GetData() interface{} {
	return s.Data
}

func (s *UpdateAgentSessionTitleResponseBodyJsonRpcResponseError) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateAgentSessionTitleResponseBodyJsonRpcResponseError) GetMessage() *string {
	return s.Message
}

func (s *UpdateAgentSessionTitleResponseBodyJsonRpcResponseError) SetCode(v int32) *UpdateAgentSessionTitleResponseBodyJsonRpcResponseError {
	s.Code = &v
	return s
}

func (s *UpdateAgentSessionTitleResponseBodyJsonRpcResponseError) SetData(v interface{}) *UpdateAgentSessionTitleResponseBodyJsonRpcResponseError {
	s.Data = v
	return s
}

func (s *UpdateAgentSessionTitleResponseBodyJsonRpcResponseError) SetErrorCode(v string) *UpdateAgentSessionTitleResponseBodyJsonRpcResponseError {
	s.ErrorCode = &v
	return s
}

func (s *UpdateAgentSessionTitleResponseBodyJsonRpcResponseError) SetMessage(v string) *UpdateAgentSessionTitleResponseBodyJsonRpcResponseError {
	s.Message = &v
	return s
}

func (s *UpdateAgentSessionTitleResponseBodyJsonRpcResponseError) Validate() error {
	return dara.Validate(s)
}

type UpdateAgentSessionTitleResponseBodyJsonRpcResponseResult struct {
	// The ID of the agent session whose title is modified. It corresponds to the SessionId in the request.
	//
	// example:
	//
	// 123e4567-e89b-42d3-a456-426614174000
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// The saved session title after modification. Leading and trailing whitespace characters are removed.
	//
	// example:
	//
	// New title
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s UpdateAgentSessionTitleResponseBodyJsonRpcResponseResult) String() string {
	return dara.Prettify(s)
}

func (s UpdateAgentSessionTitleResponseBodyJsonRpcResponseResult) GoString() string {
	return s.String()
}

func (s *UpdateAgentSessionTitleResponseBodyJsonRpcResponseResult) GetSessionId() *string {
	return s.SessionId
}

func (s *UpdateAgentSessionTitleResponseBodyJsonRpcResponseResult) GetTitle() *string {
	return s.Title
}

func (s *UpdateAgentSessionTitleResponseBodyJsonRpcResponseResult) SetSessionId(v string) *UpdateAgentSessionTitleResponseBodyJsonRpcResponseResult {
	s.SessionId = &v
	return s
}

func (s *UpdateAgentSessionTitleResponseBodyJsonRpcResponseResult) SetTitle(v string) *UpdateAgentSessionTitleResponseBodyJsonRpcResponseResult {
	s.Title = &v
	return s
}

func (s *UpdateAgentSessionTitleResponseBodyJsonRpcResponseResult) Validate() error {
	return dara.Validate(s)
}
