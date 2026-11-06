// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iArchiveAgentSessionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJsonRpcResponse(v *ArchiveAgentSessionResponseBodyJsonRpcResponse) *ArchiveAgentSessionResponseBody
	GetJsonRpcResponse() *ArchiveAgentSessionResponseBodyJsonRpcResponse
	SetRequestId(v string) *ArchiveAgentSessionResponseBody
	GetRequestId() *string
}

type ArchiveAgentSessionResponseBody struct {
	// The JSON-RPC response. Result is returned if the call is successful, and Error is returned if a protocol error occurs.
	JsonRpcResponse *ArchiveAgentSessionResponseBodyJsonRpcResponse `json:"JsonRpcResponse,omitempty" xml:"JsonRpcResponse,omitempty" type:"Struct"`
	// The OpenAPI request ID, which can be used to troubleshoot call failures.
	//
	// example:
	//
	// 01A0B6E0-E005-1465-B5DE-5CD17677426F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ArchiveAgentSessionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ArchiveAgentSessionResponseBody) GoString() string {
	return s.String()
}

func (s *ArchiveAgentSessionResponseBody) GetJsonRpcResponse() *ArchiveAgentSessionResponseBodyJsonRpcResponse {
	return s.JsonRpcResponse
}

func (s *ArchiveAgentSessionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ArchiveAgentSessionResponseBody) SetJsonRpcResponse(v *ArchiveAgentSessionResponseBodyJsonRpcResponse) *ArchiveAgentSessionResponseBody {
	s.JsonRpcResponse = v
	return s
}

func (s *ArchiveAgentSessionResponseBody) SetRequestId(v string) *ArchiveAgentSessionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ArchiveAgentSessionResponseBody) Validate() error {
	if s.JsonRpcResponse != nil {
		if err := s.JsonRpcResponse.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ArchiveAgentSessionResponseBodyJsonRpcResponse struct {
	// The JSON-RPC fault message. This field is not returned if the call is successful.
	Error *ArchiveAgentSessionResponseBodyJsonRpcResponseError `json:"Error,omitempty" xml:"Error,omitempty" type:"Struct"`
	// The JSON-RPC correlation identifier of the current request, which is passed by the caller and returned as-is.
	//
	// example:
	//
	// archive-rpc-001
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The JSON-RPC protocol version.
	//
	// example:
	//
	// 2.0
	Jsonrpc *string `json:"Jsonrpc,omitempty" xml:"Jsonrpc,omitempty"`
	// The session archive status.
	Result *ArchiveAgentSessionResponseBodyJsonRpcResponseResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// The response time, which is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1789549200000
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s ArchiveAgentSessionResponseBodyJsonRpcResponse) String() string {
	return dara.Prettify(s)
}

func (s ArchiveAgentSessionResponseBodyJsonRpcResponse) GoString() string {
	return s.String()
}

func (s *ArchiveAgentSessionResponseBodyJsonRpcResponse) GetError() *ArchiveAgentSessionResponseBodyJsonRpcResponseError {
	return s.Error
}

func (s *ArchiveAgentSessionResponseBodyJsonRpcResponse) GetId() *string {
	return s.Id
}

func (s *ArchiveAgentSessionResponseBodyJsonRpcResponse) GetJsonrpc() *string {
	return s.Jsonrpc
}

func (s *ArchiveAgentSessionResponseBodyJsonRpcResponse) GetResult() *ArchiveAgentSessionResponseBodyJsonRpcResponseResult {
	return s.Result
}

func (s *ArchiveAgentSessionResponseBodyJsonRpcResponse) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *ArchiveAgentSessionResponseBodyJsonRpcResponse) SetError(v *ArchiveAgentSessionResponseBodyJsonRpcResponseError) *ArchiveAgentSessionResponseBodyJsonRpcResponse {
	s.Error = v
	return s
}

func (s *ArchiveAgentSessionResponseBodyJsonRpcResponse) SetId(v string) *ArchiveAgentSessionResponseBodyJsonRpcResponse {
	s.Id = &v
	return s
}

func (s *ArchiveAgentSessionResponseBodyJsonRpcResponse) SetJsonrpc(v string) *ArchiveAgentSessionResponseBodyJsonRpcResponse {
	s.Jsonrpc = &v
	return s
}

func (s *ArchiveAgentSessionResponseBodyJsonRpcResponse) SetResult(v *ArchiveAgentSessionResponseBodyJsonRpcResponseResult) *ArchiveAgentSessionResponseBodyJsonRpcResponse {
	s.Result = v
	return s
}

func (s *ArchiveAgentSessionResponseBodyJsonRpcResponse) SetTimestamp(v int64) *ArchiveAgentSessionResponseBodyJsonRpcResponse {
	s.Timestamp = &v
	return s
}

func (s *ArchiveAgentSessionResponseBodyJsonRpcResponse) Validate() error {
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

type ArchiveAgentSessionResponseBodyJsonRpcResponseError struct {
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

func (s ArchiveAgentSessionResponseBodyJsonRpcResponseError) String() string {
	return dara.Prettify(s)
}

func (s ArchiveAgentSessionResponseBodyJsonRpcResponseError) GoString() string {
	return s.String()
}

func (s *ArchiveAgentSessionResponseBodyJsonRpcResponseError) GetCode() *int32 {
	return s.Code
}

func (s *ArchiveAgentSessionResponseBodyJsonRpcResponseError) GetData() interface{} {
	return s.Data
}

func (s *ArchiveAgentSessionResponseBodyJsonRpcResponseError) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ArchiveAgentSessionResponseBodyJsonRpcResponseError) GetMessage() *string {
	return s.Message
}

func (s *ArchiveAgentSessionResponseBodyJsonRpcResponseError) SetCode(v int32) *ArchiveAgentSessionResponseBodyJsonRpcResponseError {
	s.Code = &v
	return s
}

func (s *ArchiveAgentSessionResponseBodyJsonRpcResponseError) SetData(v interface{}) *ArchiveAgentSessionResponseBodyJsonRpcResponseError {
	s.Data = v
	return s
}

func (s *ArchiveAgentSessionResponseBodyJsonRpcResponseError) SetErrorCode(v string) *ArchiveAgentSessionResponseBodyJsonRpcResponseError {
	s.ErrorCode = &v
	return s
}

func (s *ArchiveAgentSessionResponseBodyJsonRpcResponseError) SetMessage(v string) *ArchiveAgentSessionResponseBodyJsonRpcResponseError {
	s.Message = &v
	return s
}

func (s *ArchiveAgentSessionResponseBodyJsonRpcResponseError) Validate() error {
	return dara.Validate(s)
}

type ArchiveAgentSessionResponseBodyJsonRpcResponseResult struct {
	// The archive status after the operation is completed. A value of true indicates that the session is archived, and a value of false indicates that the session is unarchived.
	Archived *bool `json:"Archived,omitempty" xml:"Archived,omitempty"`
	// The ID of the session for which the archive status is set.
	//
	// example:
	//
	// session_example
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s ArchiveAgentSessionResponseBodyJsonRpcResponseResult) String() string {
	return dara.Prettify(s)
}

func (s ArchiveAgentSessionResponseBodyJsonRpcResponseResult) GoString() string {
	return s.String()
}

func (s *ArchiveAgentSessionResponseBodyJsonRpcResponseResult) GetArchived() *bool {
	return s.Archived
}

func (s *ArchiveAgentSessionResponseBodyJsonRpcResponseResult) GetSessionId() *string {
	return s.SessionId
}

func (s *ArchiveAgentSessionResponseBodyJsonRpcResponseResult) SetArchived(v bool) *ArchiveAgentSessionResponseBodyJsonRpcResponseResult {
	s.Archived = &v
	return s
}

func (s *ArchiveAgentSessionResponseBodyJsonRpcResponseResult) SetSessionId(v string) *ArchiveAgentSessionResponseBodyJsonRpcResponseResult {
	s.SessionId = &v
	return s
}

func (s *ArchiveAgentSessionResponseBodyJsonRpcResponseResult) Validate() error {
	return dara.Validate(s)
}
