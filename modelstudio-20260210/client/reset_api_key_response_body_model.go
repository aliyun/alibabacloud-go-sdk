// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetApiKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApiKey(v *ResetApiKeyResponseBodyApiKey) *ResetApiKeyResponseBody
	GetApiKey() *ResetApiKeyResponseBodyApiKey
	SetCode(v string) *ResetApiKeyResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ResetApiKeyResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ResetApiKeyResponseBody
	GetMessage() *string
	SetRequestId(v string) *ResetApiKeyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ResetApiKeyResponseBody
	GetSuccess() *bool
}

type ResetApiKeyResponseBody struct {
	// The API key information.
	ApiKey *ResetApiKeyResponseBodyApiKey `json:"apiKey,omitempty" xml:"apiKey,omitempty" type:"Struct"`
	// The response status code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The response message.
	//
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 36045E0A-551D-592D-B1BC-4C56596CE59E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// - true: The request was successful.
	//
	// - false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ResetApiKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResetApiKeyResponseBody) GoString() string {
	return s.String()
}

func (s *ResetApiKeyResponseBody) GetApiKey() *ResetApiKeyResponseBodyApiKey {
	return s.ApiKey
}

func (s *ResetApiKeyResponseBody) GetCode() *string {
	return s.Code
}

func (s *ResetApiKeyResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ResetApiKeyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ResetApiKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResetApiKeyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ResetApiKeyResponseBody) SetApiKey(v *ResetApiKeyResponseBodyApiKey) *ResetApiKeyResponseBody {
	s.ApiKey = v
	return s
}

func (s *ResetApiKeyResponseBody) SetCode(v string) *ResetApiKeyResponseBody {
	s.Code = &v
	return s
}

func (s *ResetApiKeyResponseBody) SetHttpStatusCode(v int32) *ResetApiKeyResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ResetApiKeyResponseBody) SetMessage(v string) *ResetApiKeyResponseBody {
	s.Message = &v
	return s
}

func (s *ResetApiKeyResponseBody) SetRequestId(v string) *ResetApiKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResetApiKeyResponseBody) SetSuccess(v bool) *ResetApiKeyResponseBody {
	s.Success = &v
	return s
}

func (s *ResetApiKeyResponseBody) Validate() error {
	if s.ApiKey != nil {
		if err := s.ApiKey.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ResetApiKeyResponseBodyApiKey struct {
	// API Key ID。
	//
	// example:
	//
	// 2965964
	ApiKeyId *int64 `json:"apiKeyId,omitempty" xml:"apiKeyId,omitempty"`
	// The value of the API key.
	//
	// example:
	//
	// sk-ws-djI.8O7dkfkW2aICctnid4u4
	ApiKeyValue *string `json:"apiKeyValue,omitempty" xml:"apiKeyValue,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// ws-b2d30f148c236908
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s ResetApiKeyResponseBodyApiKey) String() string {
	return dara.Prettify(s)
}

func (s ResetApiKeyResponseBodyApiKey) GoString() string {
	return s.String()
}

func (s *ResetApiKeyResponseBodyApiKey) GetApiKeyId() *int64 {
	return s.ApiKeyId
}

func (s *ResetApiKeyResponseBodyApiKey) GetApiKeyValue() *string {
	return s.ApiKeyValue
}

func (s *ResetApiKeyResponseBodyApiKey) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ResetApiKeyResponseBodyApiKey) SetApiKeyId(v int64) *ResetApiKeyResponseBodyApiKey {
	s.ApiKeyId = &v
	return s
}

func (s *ResetApiKeyResponseBodyApiKey) SetApiKeyValue(v string) *ResetApiKeyResponseBodyApiKey {
	s.ApiKeyValue = &v
	return s
}

func (s *ResetApiKeyResponseBodyApiKey) SetWorkspaceId(v string) *ResetApiKeyResponseBodyApiKey {
	s.WorkspaceId = &v
	return s
}

func (s *ResetApiKeyResponseBodyApiKey) Validate() error {
	return dara.Validate(s)
}
