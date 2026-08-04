// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterCreateMemberApiKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ModelRouterCreateMemberApiKeyResponseBodyData) *ModelRouterCreateMemberApiKeyResponseBody
	GetData() *ModelRouterCreateMemberApiKeyResponseBodyData
	SetErrCode(v string) *ModelRouterCreateMemberApiKeyResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModelRouterCreateMemberApiKeyResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModelRouterCreateMemberApiKeyResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ModelRouterCreateMemberApiKeyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModelRouterCreateMemberApiKeyResponseBody
	GetSuccess() *bool
}

type ModelRouterCreateMemberApiKeyResponseBody struct {
	// The data object.
	//
	// example:
	//
	// { "apiKeyId": 502, "apiKey": "sk-xxxxxxxxxxxxxxxx" }
	Data *ModelRouterCreateMemberApiKeyResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The fault message code.
	//
	// example:
	//
	// UNKNOWN_ERROR
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Unknown error
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// xxxx-xxxx-xxxx-xxxxxxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ModelRouterCreateMemberApiKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterCreateMemberApiKeyResponseBody) GoString() string {
	return s.String()
}

func (s *ModelRouterCreateMemberApiKeyResponseBody) GetData() *ModelRouterCreateMemberApiKeyResponseBodyData {
	return s.Data
}

func (s *ModelRouterCreateMemberApiKeyResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModelRouterCreateMemberApiKeyResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModelRouterCreateMemberApiKeyResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModelRouterCreateMemberApiKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModelRouterCreateMemberApiKeyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModelRouterCreateMemberApiKeyResponseBody) SetData(v *ModelRouterCreateMemberApiKeyResponseBodyData) *ModelRouterCreateMemberApiKeyResponseBody {
	s.Data = v
	return s
}

func (s *ModelRouterCreateMemberApiKeyResponseBody) SetErrCode(v string) *ModelRouterCreateMemberApiKeyResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModelRouterCreateMemberApiKeyResponseBody) SetErrMessage(v string) *ModelRouterCreateMemberApiKeyResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModelRouterCreateMemberApiKeyResponseBody) SetHttpStatusCode(v int32) *ModelRouterCreateMemberApiKeyResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModelRouterCreateMemberApiKeyResponseBody) SetRequestId(v string) *ModelRouterCreateMemberApiKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModelRouterCreateMemberApiKeyResponseBody) SetSuccess(v bool) *ModelRouterCreateMemberApiKeyResponseBody {
	s.Success = &v
	return s
}

func (s *ModelRouterCreateMemberApiKeyResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModelRouterCreateMemberApiKeyResponseBodyData struct {
	// The full API key in plaintext. This value is returned only once during creation. Store it securely.
	//
	// example:
	//
	// sk-xxxxxxxxxxxxxxxx
	ApiKey *string `json:"apiKey,omitempty" xml:"apiKey,omitempty"`
	// The API key ID. You can use this ID to query the bound groups by API key.
	//
	// example:
	//
	// 502
	ApiKeyId *int64 `json:"apiKeyId,omitempty" xml:"apiKeyId,omitempty"`
}

func (s ModelRouterCreateMemberApiKeyResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterCreateMemberApiKeyResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModelRouterCreateMemberApiKeyResponseBodyData) GetApiKey() *string {
	return s.ApiKey
}

func (s *ModelRouterCreateMemberApiKeyResponseBodyData) GetApiKeyId() *int64 {
	return s.ApiKeyId
}

func (s *ModelRouterCreateMemberApiKeyResponseBodyData) SetApiKey(v string) *ModelRouterCreateMemberApiKeyResponseBodyData {
	s.ApiKey = &v
	return s
}

func (s *ModelRouterCreateMemberApiKeyResponseBodyData) SetApiKeyId(v int64) *ModelRouterCreateMemberApiKeyResponseBodyData {
	s.ApiKeyId = &v
	return s
}

func (s *ModelRouterCreateMemberApiKeyResponseBodyData) Validate() error {
	return dara.Validate(s)
}
