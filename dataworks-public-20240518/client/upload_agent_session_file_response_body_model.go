// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadAgentSessionFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJsonRpcResponse(v *UploadAgentSessionFileResponseBodyJsonRpcResponse) *UploadAgentSessionFileResponseBody
	GetJsonRpcResponse() *UploadAgentSessionFileResponseBodyJsonRpcResponse
	SetRequestId(v string) *UploadAgentSessionFileResponseBody
	GetRequestId() *string
}

type UploadAgentSessionFileResponseBody struct {
	// The JSON-RPC response. Returns Result on success and Error on protocol errors.
	JsonRpcResponse *UploadAgentSessionFileResponseBodyJsonRpcResponse `json:"JsonRpcResponse,omitempty" xml:"JsonRpcResponse,omitempty" type:"Struct"`
	// The server correlation ID for the current request. You can provide this value for troubleshooting.
	//
	// example:
	//
	// 01A0B6E0-E005-1465-B5DE-5CD17677426F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UploadAgentSessionFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UploadAgentSessionFileResponseBody) GoString() string {
	return s.String()
}

func (s *UploadAgentSessionFileResponseBody) GetJsonRpcResponse() *UploadAgentSessionFileResponseBodyJsonRpcResponse {
	return s.JsonRpcResponse
}

func (s *UploadAgentSessionFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UploadAgentSessionFileResponseBody) SetJsonRpcResponse(v *UploadAgentSessionFileResponseBodyJsonRpcResponse) *UploadAgentSessionFileResponseBody {
	s.JsonRpcResponse = v
	return s
}

func (s *UploadAgentSessionFileResponseBody) SetRequestId(v string) *UploadAgentSessionFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *UploadAgentSessionFileResponseBody) Validate() error {
	if s.JsonRpcResponse != nil {
		if err := s.JsonRpcResponse.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UploadAgentSessionFileResponseBodyJsonRpcResponse struct {
	// The JSON-RPC fault message. This field is not returned on success.
	Error *UploadAgentSessionFileResponseBodyJsonRpcResponseError `json:"Error,omitempty" xml:"Error,omitempty" type:"Struct"`
	// The JSON-RPC correlation identifier for the current request, which is passed by the caller and returned as is.
	//
	// example:
	//
	// attachment-rpc-001
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The JSON-RPC protocol version.
	//
	// example:
	//
	// 2.0
	Jsonrpc *string `json:"Jsonrpc,omitempty" xml:"Jsonrpc,omitempty"`
	// Issues the staged upload URL or returns the attachment information after writing to the instance.
	Result *UploadAgentSessionFileResponseBodyJsonRpcResponseResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// The response time as a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1789549200000
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s UploadAgentSessionFileResponseBodyJsonRpcResponse) String() string {
	return dara.Prettify(s)
}

func (s UploadAgentSessionFileResponseBodyJsonRpcResponse) GoString() string {
	return s.String()
}

func (s *UploadAgentSessionFileResponseBodyJsonRpcResponse) GetError() *UploadAgentSessionFileResponseBodyJsonRpcResponseError {
	return s.Error
}

func (s *UploadAgentSessionFileResponseBodyJsonRpcResponse) GetId() *string {
	return s.Id
}

func (s *UploadAgentSessionFileResponseBodyJsonRpcResponse) GetJsonrpc() *string {
	return s.Jsonrpc
}

func (s *UploadAgentSessionFileResponseBodyJsonRpcResponse) GetResult() *UploadAgentSessionFileResponseBodyJsonRpcResponseResult {
	return s.Result
}

func (s *UploadAgentSessionFileResponseBodyJsonRpcResponse) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *UploadAgentSessionFileResponseBodyJsonRpcResponse) SetError(v *UploadAgentSessionFileResponseBodyJsonRpcResponseError) *UploadAgentSessionFileResponseBodyJsonRpcResponse {
	s.Error = v
	return s
}

func (s *UploadAgentSessionFileResponseBodyJsonRpcResponse) SetId(v string) *UploadAgentSessionFileResponseBodyJsonRpcResponse {
	s.Id = &v
	return s
}

func (s *UploadAgentSessionFileResponseBodyJsonRpcResponse) SetJsonrpc(v string) *UploadAgentSessionFileResponseBodyJsonRpcResponse {
	s.Jsonrpc = &v
	return s
}

func (s *UploadAgentSessionFileResponseBodyJsonRpcResponse) SetResult(v *UploadAgentSessionFileResponseBodyJsonRpcResponseResult) *UploadAgentSessionFileResponseBodyJsonRpcResponse {
	s.Result = v
	return s
}

func (s *UploadAgentSessionFileResponseBodyJsonRpcResponse) SetTimestamp(v int64) *UploadAgentSessionFileResponseBodyJsonRpcResponse {
	s.Timestamp = &v
	return s
}

func (s *UploadAgentSessionFileResponseBodyJsonRpcResponse) Validate() error {
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

type UploadAgentSessionFileResponseBodyJsonRpcResponseError struct {
	// The JSON-RPC error code.
	//
	// example:
	//
	// 400
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
	// SESSION_ATTACHMENT_OBJECT_MISMATCH
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// SESSION_ATTACHMENT_OBJECT_MISMATCH
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s UploadAgentSessionFileResponseBodyJsonRpcResponseError) String() string {
	return dara.Prettify(s)
}

func (s UploadAgentSessionFileResponseBodyJsonRpcResponseError) GoString() string {
	return s.String()
}

func (s *UploadAgentSessionFileResponseBodyJsonRpcResponseError) GetCode() *int32 {
	return s.Code
}

func (s *UploadAgentSessionFileResponseBodyJsonRpcResponseError) GetData() interface{} {
	return s.Data
}

func (s *UploadAgentSessionFileResponseBodyJsonRpcResponseError) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UploadAgentSessionFileResponseBodyJsonRpcResponseError) GetMessage() *string {
	return s.Message
}

func (s *UploadAgentSessionFileResponseBodyJsonRpcResponseError) SetCode(v int32) *UploadAgentSessionFileResponseBodyJsonRpcResponseError {
	s.Code = &v
	return s
}

func (s *UploadAgentSessionFileResponseBodyJsonRpcResponseError) SetData(v interface{}) *UploadAgentSessionFileResponseBodyJsonRpcResponseError {
	s.Data = v
	return s
}

func (s *UploadAgentSessionFileResponseBodyJsonRpcResponseError) SetErrorCode(v string) *UploadAgentSessionFileResponseBodyJsonRpcResponseError {
	s.ErrorCode = &v
	return s
}

func (s *UploadAgentSessionFileResponseBodyJsonRpcResponseError) SetMessage(v string) *UploadAgentSessionFileResponseBodyJsonRpcResponseError {
	s.Message = &v
	return s
}

func (s *UploadAgentSessionFileResponseBodyJsonRpcResponseError) Validate() error {
	return dara.Validate(s)
}

type UploadAgentSessionFileResponseBodyJsonRpcResponseResult struct {
	// The attachment identifier.
	//
	// example:
	//
	// bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The original file name.
	//
	// example:
	//
	// sample.csv
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The relative path in the session working directory.
	//
	// example:
	//
	// upload/session_bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb.csv
	RelativePath *string `json:"RelativePath,omitempty" xml:"RelativePath,omitempty"`
	// The time-limited object storage signature returned by prepare.
	Signature *UploadAgentSessionFileResponseBodyJsonRpcResponseResultSignature `json:"Signature,omitempty" xml:"Signature,omitempty" type:"Struct"`
	// The number of file bytes saved in the instance.
	//
	// example:
	//
	// 6
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The staged upload identifier issued by prepare.
	//
	// example:
	//
	// bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb
	UploadId *string `json:"UploadId,omitempty" xml:"UploadId,omitempty"`
}

func (s UploadAgentSessionFileResponseBodyJsonRpcResponseResult) String() string {
	return dara.Prettify(s)
}

func (s UploadAgentSessionFileResponseBodyJsonRpcResponseResult) GoString() string {
	return s.String()
}

func (s *UploadAgentSessionFileResponseBodyJsonRpcResponseResult) GetId() *string {
	return s.Id
}

func (s *UploadAgentSessionFileResponseBodyJsonRpcResponseResult) GetName() *string {
	return s.Name
}

func (s *UploadAgentSessionFileResponseBodyJsonRpcResponseResult) GetRelativePath() *string {
	return s.RelativePath
}

func (s *UploadAgentSessionFileResponseBodyJsonRpcResponseResult) GetSignature() *UploadAgentSessionFileResponseBodyJsonRpcResponseResultSignature {
	return s.Signature
}

func (s *UploadAgentSessionFileResponseBodyJsonRpcResponseResult) GetSize() *int64 {
	return s.Size
}

func (s *UploadAgentSessionFileResponseBodyJsonRpcResponseResult) GetUploadId() *string {
	return s.UploadId
}

func (s *UploadAgentSessionFileResponseBodyJsonRpcResponseResult) SetId(v string) *UploadAgentSessionFileResponseBodyJsonRpcResponseResult {
	s.Id = &v
	return s
}

func (s *UploadAgentSessionFileResponseBodyJsonRpcResponseResult) SetName(v string) *UploadAgentSessionFileResponseBodyJsonRpcResponseResult {
	s.Name = &v
	return s
}

func (s *UploadAgentSessionFileResponseBodyJsonRpcResponseResult) SetRelativePath(v string) *UploadAgentSessionFileResponseBodyJsonRpcResponseResult {
	s.RelativePath = &v
	return s
}

func (s *UploadAgentSessionFileResponseBodyJsonRpcResponseResult) SetSignature(v *UploadAgentSessionFileResponseBodyJsonRpcResponseResultSignature) *UploadAgentSessionFileResponseBodyJsonRpcResponseResult {
	s.Signature = v
	return s
}

func (s *UploadAgentSessionFileResponseBodyJsonRpcResponseResult) SetSize(v int64) *UploadAgentSessionFileResponseBodyJsonRpcResponseResult {
	s.Size = &v
	return s
}

func (s *UploadAgentSessionFileResponseBodyJsonRpcResponseResult) SetUploadId(v string) *UploadAgentSessionFileResponseBodyJsonRpcResponseResult {
	s.UploadId = &v
	return s
}

func (s *UploadAgentSessionFileResponseBodyJsonRpcResponseResult) Validate() error {
	if s.Signature != nil {
		if err := s.Signature.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UploadAgentSessionFileResponseBodyJsonRpcResponseResultSignature struct {
	// The expiration time of the current upload signature in ISO 8601 UTC string format. The trailing Z indicates UTC, and fractional seconds may be included. The signature is valid for 5 minutes. Signature expiration does not indicate that the staged object has been deleted.
	//
	// example:
	//
	// 2026-11-12T09:46:33.477365753Z
	ExpiresAt *string `json:"ExpiresAt,omitempty" xml:"ExpiresAt,omitempty"`
	// All signed request headers that must be included in the PUT request. The names and values must be used exactly as provided. This includes Content-Type, Content-MD5, Content-Length, x-oss-object-acl, and x-oss-forbid-overwrite.
	//
	// example:
	//
	// {"Content-Type":"application/octet-stream","Content-MD5":"sZRqySSS0jR8YjW00mERhA==","Content-Length":"6","x-oss-object-acl":"private","x-oss-forbid-overwrite":"true"}
	Headers map[string]*string `json:"Headers,omitempty" xml:"Headers,omitempty"`
	// The complete signed HTTPS PUT URL returned by prepare, which must be used exactly as provided for the upload. The example displayed for this field does not contain a valid signature. The actual request must use the complete signed URL returned by the current prepare call.
	//
	// example:
	//
	// https://attachments.example.test/signed
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s UploadAgentSessionFileResponseBodyJsonRpcResponseResultSignature) String() string {
	return dara.Prettify(s)
}

func (s UploadAgentSessionFileResponseBodyJsonRpcResponseResultSignature) GoString() string {
	return s.String()
}

func (s *UploadAgentSessionFileResponseBodyJsonRpcResponseResultSignature) GetExpiresAt() *string {
	return s.ExpiresAt
}

func (s *UploadAgentSessionFileResponseBodyJsonRpcResponseResultSignature) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UploadAgentSessionFileResponseBodyJsonRpcResponseResultSignature) GetUrl() *string {
	return s.Url
}

func (s *UploadAgentSessionFileResponseBodyJsonRpcResponseResultSignature) SetExpiresAt(v string) *UploadAgentSessionFileResponseBodyJsonRpcResponseResultSignature {
	s.ExpiresAt = &v
	return s
}

func (s *UploadAgentSessionFileResponseBodyJsonRpcResponseResultSignature) SetHeaders(v map[string]*string) *UploadAgentSessionFileResponseBodyJsonRpcResponseResultSignature {
	s.Headers = v
	return s
}

func (s *UploadAgentSessionFileResponseBodyJsonRpcResponseResultSignature) SetUrl(v string) *UploadAgentSessionFileResponseBodyJsonRpcResponseResultSignature {
	s.Url = &v
	return s
}

func (s *UploadAgentSessionFileResponseBodyJsonRpcResponseResultSignature) Validate() error {
	return dara.Validate(s)
}
