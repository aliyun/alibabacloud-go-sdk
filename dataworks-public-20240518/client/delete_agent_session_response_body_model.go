// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAgentSessionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJsonRpcResponse(v *DeleteAgentSessionResponseBodyJsonRpcResponse) *DeleteAgentSessionResponseBody
	GetJsonRpcResponse() *DeleteAgentSessionResponseBodyJsonRpcResponse
	SetRequestId(v string) *DeleteAgentSessionResponseBody
	GetRequestId() *string
}

type DeleteAgentSessionResponseBody struct {
	// The JSON-RPC response. It returns Result upon success and Error upon protocol errors.
	JsonRpcResponse *DeleteAgentSessionResponseBodyJsonRpcResponse `json:"JsonRpcResponse,omitempty" xml:"JsonRpcResponse,omitempty" type:"Struct"`
	// The OpenAPI request ID, which can be used to troubleshoot call failures.
	//
	// example:
	//
	// 01A0B6E0-E005-1465-B5DE-5CD17677426F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAgentSessionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAgentSessionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAgentSessionResponseBody) GetJsonRpcResponse() *DeleteAgentSessionResponseBodyJsonRpcResponse {
	return s.JsonRpcResponse
}

func (s *DeleteAgentSessionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAgentSessionResponseBody) SetJsonRpcResponse(v *DeleteAgentSessionResponseBodyJsonRpcResponse) *DeleteAgentSessionResponseBody {
	s.JsonRpcResponse = v
	return s
}

func (s *DeleteAgentSessionResponseBody) SetRequestId(v string) *DeleteAgentSessionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAgentSessionResponseBody) Validate() error {
	if s.JsonRpcResponse != nil {
		if err := s.JsonRpcResponse.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteAgentSessionResponseBodyJsonRpcResponse struct {
	// The JSON-RPC fault message. This field is not returned upon success.
	Error *DeleteAgentSessionResponseBodyJsonRpcResponseError `json:"Error,omitempty" xml:"Error,omitempty" type:"Struct"`
	// The JSON-RPC correlation identifier for the current request. It is passed by the caller and returned as-is.
	//
	// example:
	//
	// delete-rpc-001
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The JSON-RPC protocol version.
	//
	// example:
	//
	// 2.0
	Jsonrpc *string `json:"Jsonrpc,omitempty" xml:"Jsonrpc,omitempty"`
	// An empty object is returned upon successful deletion. If the deletion fails, check the sibling Error field.
	//
	// example:
	//
	// {}
	Result interface{} `json:"Result,omitempty" xml:"Result,omitempty"`
	// The response time. It is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1789549200000
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s DeleteAgentSessionResponseBodyJsonRpcResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAgentSessionResponseBodyJsonRpcResponse) GoString() string {
	return s.String()
}

func (s *DeleteAgentSessionResponseBodyJsonRpcResponse) GetError() *DeleteAgentSessionResponseBodyJsonRpcResponseError {
	return s.Error
}

func (s *DeleteAgentSessionResponseBodyJsonRpcResponse) GetId() *string {
	return s.Id
}

func (s *DeleteAgentSessionResponseBodyJsonRpcResponse) GetJsonrpc() *string {
	return s.Jsonrpc
}

func (s *DeleteAgentSessionResponseBodyJsonRpcResponse) GetResult() interface{} {
	return s.Result
}

func (s *DeleteAgentSessionResponseBodyJsonRpcResponse) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *DeleteAgentSessionResponseBodyJsonRpcResponse) SetError(v *DeleteAgentSessionResponseBodyJsonRpcResponseError) *DeleteAgentSessionResponseBodyJsonRpcResponse {
	s.Error = v
	return s
}

func (s *DeleteAgentSessionResponseBodyJsonRpcResponse) SetId(v string) *DeleteAgentSessionResponseBodyJsonRpcResponse {
	s.Id = &v
	return s
}

func (s *DeleteAgentSessionResponseBodyJsonRpcResponse) SetJsonrpc(v string) *DeleteAgentSessionResponseBodyJsonRpcResponse {
	s.Jsonrpc = &v
	return s
}

func (s *DeleteAgentSessionResponseBodyJsonRpcResponse) SetResult(v interface{}) *DeleteAgentSessionResponseBodyJsonRpcResponse {
	s.Result = v
	return s
}

func (s *DeleteAgentSessionResponseBodyJsonRpcResponse) SetTimestamp(v int64) *DeleteAgentSessionResponseBodyJsonRpcResponse {
	s.Timestamp = &v
	return s
}

func (s *DeleteAgentSessionResponseBodyJsonRpcResponse) Validate() error {
	if s.Error != nil {
		if err := s.Error.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteAgentSessionResponseBodyJsonRpcResponseError struct {
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
	// The session does not exist or the current account does not have access permissions. Return to the session list and select another session
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s DeleteAgentSessionResponseBodyJsonRpcResponseError) String() string {
	return dara.Prettify(s)
}

func (s DeleteAgentSessionResponseBodyJsonRpcResponseError) GoString() string {
	return s.String()
}

func (s *DeleteAgentSessionResponseBodyJsonRpcResponseError) GetCode() *int32 {
	return s.Code
}

func (s *DeleteAgentSessionResponseBodyJsonRpcResponseError) GetData() interface{} {
	return s.Data
}

func (s *DeleteAgentSessionResponseBodyJsonRpcResponseError) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteAgentSessionResponseBodyJsonRpcResponseError) GetMessage() *string {
	return s.Message
}

func (s *DeleteAgentSessionResponseBodyJsonRpcResponseError) SetCode(v int32) *DeleteAgentSessionResponseBodyJsonRpcResponseError {
	s.Code = &v
	return s
}

func (s *DeleteAgentSessionResponseBodyJsonRpcResponseError) SetData(v interface{}) *DeleteAgentSessionResponseBodyJsonRpcResponseError {
	s.Data = v
	return s
}

func (s *DeleteAgentSessionResponseBodyJsonRpcResponseError) SetErrorCode(v string) *DeleteAgentSessionResponseBodyJsonRpcResponseError {
	s.ErrorCode = &v
	return s
}

func (s *DeleteAgentSessionResponseBodyJsonRpcResponseError) SetMessage(v string) *DeleteAgentSessionResponseBodyJsonRpcResponseError {
	s.Message = &v
	return s
}

func (s *DeleteAgentSessionResponseBodyJsonRpcResponseError) Validate() error {
	return dara.Validate(s)
}
