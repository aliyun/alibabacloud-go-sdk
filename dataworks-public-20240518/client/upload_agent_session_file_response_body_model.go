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
	JsonRpcResponse *UploadAgentSessionFileResponseBodyJsonRpcResponse `json:"JsonRpcResponse,omitempty" xml:"JsonRpcResponse,omitempty" type:"Struct"`
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
	Error *UploadAgentSessionFileResponseBodyJsonRpcResponseError `json:"Error,omitempty" xml:"Error,omitempty" type:"Struct"`
	// example:
	//
	// attachment-rpc-001
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 2.0
	Jsonrpc *string                                                  `json:"Jsonrpc,omitempty" xml:"Jsonrpc,omitempty"`
	Result  *UploadAgentSessionFileResponseBodyJsonRpcResponseResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
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
	// example:
	//
	// 500
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// {}
	Data interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 0x4883300000000019
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// 会话不存在或当前账号无权访问，请返回会话列表重新选择
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
	// example:
	//
	// bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// sample.csv
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// upload/session_bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb.csv
	RelativePath *string                                                           `json:"RelativePath,omitempty" xml:"RelativePath,omitempty"`
	Signature    *UploadAgentSessionFileResponseBodyJsonRpcResponseResultSignature `json:"Signature,omitempty" xml:"Signature,omitempty" type:"Struct"`
	// example:
	//
	// 6
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
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
	// example:
	//
	// 2026-11-12T09:46:33.477365753Z
	ExpiresAt *string `json:"ExpiresAt,omitempty" xml:"ExpiresAt,omitempty"`
	// example:
	//
	// {"Content-Type":"application/octet-stream","Content-MD5":"sZRqySSS0jR8YjW00mERhA==","Content-Length":"6","x-oss-object-acl":"private","x-oss-forbid-overwrite":"true"}
	Headers map[string]*string `json:"Headers,omitempty" xml:"Headers,omitempty"`
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
