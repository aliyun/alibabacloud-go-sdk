// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApiKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApiKey(v *GetApiKeyResponseBodyApiKey) *GetApiKeyResponseBody
	GetApiKey() *GetApiKeyResponseBodyApiKey
	SetCode(v string) *GetApiKeyResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *GetApiKeyResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetApiKeyResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetApiKeyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetApiKeyResponseBody
	GetSuccess() *bool
}

type GetApiKeyResponseBody struct {
	// API Key。
	ApiKey *GetApiKeyResponseBodyApiKey `json:"apiKey,omitempty" xml:"apiKey,omitempty" type:"Struct"`
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 9B9650D6-68B0-55CC-845D-E8C1E5BED38B
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetApiKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetApiKeyResponseBody) GoString() string {
	return s.String()
}

func (s *GetApiKeyResponseBody) GetApiKey() *GetApiKeyResponseBodyApiKey {
	return s.ApiKey
}

func (s *GetApiKeyResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetApiKeyResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetApiKeyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetApiKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetApiKeyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetApiKeyResponseBody) SetApiKey(v *GetApiKeyResponseBodyApiKey) *GetApiKeyResponseBody {
	s.ApiKey = v
	return s
}

func (s *GetApiKeyResponseBody) SetCode(v string) *GetApiKeyResponseBody {
	s.Code = &v
	return s
}

func (s *GetApiKeyResponseBody) SetHttpStatusCode(v int32) *GetApiKeyResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetApiKeyResponseBody) SetMessage(v string) *GetApiKeyResponseBody {
	s.Message = &v
	return s
}

func (s *GetApiKeyResponseBody) SetRequestId(v string) *GetApiKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetApiKeyResponseBody) SetSuccess(v bool) *GetApiKeyResponseBody {
	s.Success = &v
	return s
}

func (s *GetApiKeyResponseBody) Validate() error {
	if s.ApiKey != nil {
		if err := s.ApiKey.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetApiKeyResponseBodyApiKey struct {
	// API Key ID。
	//
	// example:
	//
	// 3303332
	ApiKeyId *int64 `json:"apiKeyId,omitempty" xml:"apiKeyId,omitempty"`
	// API Key Value。
	//
	// example:
	//
	// sk-f83898fqwer340049bae6209cbb0bc29
	ApiKeyValue *string `json:"apiKeyValue,omitempty" xml:"apiKeyValue,omitempty"`
	// example:
	//
	// 1378030599924858
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty"`
	// example:
	//
	// test
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// 1774338222000
	GmtCreate *int64 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// true
	IsBlocked *int32 `json:"isBlocked,omitempty" xml:"isBlocked,omitempty"`
	// example:
	//
	// llm-jj04jmmzqjz3hw8t
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s GetApiKeyResponseBodyApiKey) String() string {
	return dara.Prettify(s)
}

func (s GetApiKeyResponseBodyApiKey) GoString() string {
	return s.String()
}

func (s *GetApiKeyResponseBodyApiKey) GetApiKeyId() *int64 {
	return s.ApiKeyId
}

func (s *GetApiKeyResponseBodyApiKey) GetApiKeyValue() *string {
	return s.ApiKeyValue
}

func (s *GetApiKeyResponseBodyApiKey) GetCreator() *string {
	return s.Creator
}

func (s *GetApiKeyResponseBodyApiKey) GetDescription() *string {
	return s.Description
}

func (s *GetApiKeyResponseBodyApiKey) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *GetApiKeyResponseBodyApiKey) GetIsBlocked() *int32 {
	return s.IsBlocked
}

func (s *GetApiKeyResponseBodyApiKey) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetApiKeyResponseBodyApiKey) SetApiKeyId(v int64) *GetApiKeyResponseBodyApiKey {
	s.ApiKeyId = &v
	return s
}

func (s *GetApiKeyResponseBodyApiKey) SetApiKeyValue(v string) *GetApiKeyResponseBodyApiKey {
	s.ApiKeyValue = &v
	return s
}

func (s *GetApiKeyResponseBodyApiKey) SetCreator(v string) *GetApiKeyResponseBodyApiKey {
	s.Creator = &v
	return s
}

func (s *GetApiKeyResponseBodyApiKey) SetDescription(v string) *GetApiKeyResponseBodyApiKey {
	s.Description = &v
	return s
}

func (s *GetApiKeyResponseBodyApiKey) SetGmtCreate(v int64) *GetApiKeyResponseBodyApiKey {
	s.GmtCreate = &v
	return s
}

func (s *GetApiKeyResponseBodyApiKey) SetIsBlocked(v int32) *GetApiKeyResponseBodyApiKey {
	s.IsBlocked = &v
	return s
}

func (s *GetApiKeyResponseBodyApiKey) SetWorkspaceId(v string) *GetApiKeyResponseBodyApiKey {
	s.WorkspaceId = &v
	return s
}

func (s *GetApiKeyResponseBodyApiKey) Validate() error {
	return dara.Validate(s)
}
