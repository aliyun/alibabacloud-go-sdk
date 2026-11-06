// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAgentSessionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJsonRpcResponse(v *CreateAgentSessionResponseBodyJsonRpcResponse) *CreateAgentSessionResponseBody
	GetJsonRpcResponse() *CreateAgentSessionResponseBodyJsonRpcResponse
	SetRequestId(v string) *CreateAgentSessionResponseBody
	GetRequestId() *string
}

type CreateAgentSessionResponseBody struct {
	// The JSON-RPC response result.
	JsonRpcResponse *CreateAgentSessionResponseBodyJsonRpcResponse `json:"JsonRpcResponse,omitempty" xml:"JsonRpcResponse,omitempty" type:"Struct"`
	// The request ID for this OpenAPI call. You can use it to troubleshoot issues with the call.
	//
	// example:
	//
	// 8A9D5E6C-5817-5837-9715-6E3967EC6123
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAgentSessionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentSessionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAgentSessionResponseBody) GetJsonRpcResponse() *CreateAgentSessionResponseBodyJsonRpcResponse {
	return s.JsonRpcResponse
}

func (s *CreateAgentSessionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAgentSessionResponseBody) SetJsonRpcResponse(v *CreateAgentSessionResponseBodyJsonRpcResponse) *CreateAgentSessionResponseBody {
	s.JsonRpcResponse = v
	return s
}

func (s *CreateAgentSessionResponseBody) SetRequestId(v string) *CreateAgentSessionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAgentSessionResponseBody) Validate() error {
	if s.JsonRpcResponse != nil {
		if err := s.JsonRpcResponse.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateAgentSessionResponseBodyJsonRpcResponse struct {
	// The JSON-RPC error message. This field is not returned when the request is successful.
	Error *CreateAgentSessionResponseBodyJsonRpcResponseError `json:"Error,omitempty" xml:"Error,omitempty" type:"Struct"`
	// The ID passed by the requester. The value is returned as is.
	//
	// example:
	//
	// 1
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The JSON-RPC version. Fixed to 2.0.
	//
	// example:
	//
	// 2.0
	Jsonrpc *string `json:"Jsonrpc,omitempty" xml:"Jsonrpc,omitempty"`
	// The business data. The value is null when an error occurs.
	Result *CreateAgentSessionResponseBodyJsonRpcResponseResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// The response time. It is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1789549200000
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s CreateAgentSessionResponseBodyJsonRpcResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentSessionResponseBodyJsonRpcResponse) GoString() string {
	return s.String()
}

func (s *CreateAgentSessionResponseBodyJsonRpcResponse) GetError() *CreateAgentSessionResponseBodyJsonRpcResponseError {
	return s.Error
}

func (s *CreateAgentSessionResponseBodyJsonRpcResponse) GetId() *string {
	return s.Id
}

func (s *CreateAgentSessionResponseBodyJsonRpcResponse) GetJsonrpc() *string {
	return s.Jsonrpc
}

func (s *CreateAgentSessionResponseBodyJsonRpcResponse) GetResult() *CreateAgentSessionResponseBodyJsonRpcResponseResult {
	return s.Result
}

func (s *CreateAgentSessionResponseBodyJsonRpcResponse) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *CreateAgentSessionResponseBodyJsonRpcResponse) SetError(v *CreateAgentSessionResponseBodyJsonRpcResponseError) *CreateAgentSessionResponseBodyJsonRpcResponse {
	s.Error = v
	return s
}

func (s *CreateAgentSessionResponseBodyJsonRpcResponse) SetId(v string) *CreateAgentSessionResponseBodyJsonRpcResponse {
	s.Id = &v
	return s
}

func (s *CreateAgentSessionResponseBodyJsonRpcResponse) SetJsonrpc(v string) *CreateAgentSessionResponseBodyJsonRpcResponse {
	s.Jsonrpc = &v
	return s
}

func (s *CreateAgentSessionResponseBodyJsonRpcResponse) SetResult(v *CreateAgentSessionResponseBodyJsonRpcResponseResult) *CreateAgentSessionResponseBodyJsonRpcResponse {
	s.Result = v
	return s
}

func (s *CreateAgentSessionResponseBodyJsonRpcResponse) SetTimestamp(v int64) *CreateAgentSessionResponseBodyJsonRpcResponse {
	s.Timestamp = &v
	return s
}

func (s *CreateAgentSessionResponseBodyJsonRpcResponse) Validate() error {
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

type CreateAgentSessionResponseBodyJsonRpcResponseError struct {
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
	// To create a DataAgent instance, you must specify a resource group. Pass the resource group ID in _meta.dataAgent.resourceGroupId of the request
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s CreateAgentSessionResponseBodyJsonRpcResponseError) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentSessionResponseBodyJsonRpcResponseError) GoString() string {
	return s.String()
}

func (s *CreateAgentSessionResponseBodyJsonRpcResponseError) GetCode() *int32 {
	return s.Code
}

func (s *CreateAgentSessionResponseBodyJsonRpcResponseError) GetData() interface{} {
	return s.Data
}

func (s *CreateAgentSessionResponseBodyJsonRpcResponseError) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateAgentSessionResponseBodyJsonRpcResponseError) GetMessage() *string {
	return s.Message
}

func (s *CreateAgentSessionResponseBodyJsonRpcResponseError) SetCode(v int32) *CreateAgentSessionResponseBodyJsonRpcResponseError {
	s.Code = &v
	return s
}

func (s *CreateAgentSessionResponseBodyJsonRpcResponseError) SetData(v interface{}) *CreateAgentSessionResponseBodyJsonRpcResponseError {
	s.Data = v
	return s
}

func (s *CreateAgentSessionResponseBodyJsonRpcResponseError) SetErrorCode(v string) *CreateAgentSessionResponseBodyJsonRpcResponseError {
	s.ErrorCode = &v
	return s
}

func (s *CreateAgentSessionResponseBodyJsonRpcResponseError) SetMessage(v string) *CreateAgentSessionResponseBodyJsonRpcResponseError {
	s.Message = &v
	return s
}

func (s *CreateAgentSessionResponseBodyJsonRpcResponseError) Validate() error {
	return dara.Validate(s)
}

type CreateAgentSessionResponseBodyJsonRpcResponseResult struct {
	// The ID of the newly created session.
	//
	// example:
	//
	// sess_0f12abc34
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s CreateAgentSessionResponseBodyJsonRpcResponseResult) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentSessionResponseBodyJsonRpcResponseResult) GoString() string {
	return s.String()
}

func (s *CreateAgentSessionResponseBodyJsonRpcResponseResult) GetSessionId() *string {
	return s.SessionId
}

func (s *CreateAgentSessionResponseBodyJsonRpcResponseResult) SetSessionId(v string) *CreateAgentSessionResponseBodyJsonRpcResponseResult {
	s.SessionId = &v
	return s
}

func (s *CreateAgentSessionResponseBodyJsonRpcResponseResult) Validate() error {
	return dara.Validate(s)
}
