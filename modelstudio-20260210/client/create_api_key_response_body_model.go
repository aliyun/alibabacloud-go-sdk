// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApiKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApiKey(v *CreateApiKeyResponseBodyApiKey) *CreateApiKeyResponseBody
	GetApiKey() *CreateApiKeyResponseBodyApiKey
	SetCode(v string) *CreateApiKeyResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *CreateApiKeyResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateApiKeyResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateApiKeyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateApiKeyResponseBody
	GetSuccess() *bool
}

type CreateApiKeyResponseBody struct {
	// The API key information.
	ApiKey *CreateApiKeyResponseBodyApiKey `json:"apiKey,omitempty" xml:"apiKey,omitempty" type:"Struct"`
	// The response code.
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
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// C0CDC72E-52D7-5BC8-9396-9276B4FDF6B3
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// - true: Succeeded.
	//
	// - false: Failed.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateApiKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateApiKeyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateApiKeyResponseBody) GetApiKey() *CreateApiKeyResponseBodyApiKey {
	return s.ApiKey
}

func (s *CreateApiKeyResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateApiKeyResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateApiKeyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateApiKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateApiKeyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateApiKeyResponseBody) SetApiKey(v *CreateApiKeyResponseBodyApiKey) *CreateApiKeyResponseBody {
	s.ApiKey = v
	return s
}

func (s *CreateApiKeyResponseBody) SetCode(v string) *CreateApiKeyResponseBody {
	s.Code = &v
	return s
}

func (s *CreateApiKeyResponseBody) SetHttpStatusCode(v int32) *CreateApiKeyResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateApiKeyResponseBody) SetMessage(v string) *CreateApiKeyResponseBody {
	s.Message = &v
	return s
}

func (s *CreateApiKeyResponseBody) SetRequestId(v string) *CreateApiKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateApiKeyResponseBody) SetSuccess(v bool) *CreateApiKeyResponseBody {
	s.Success = &v
	return s
}

func (s *CreateApiKeyResponseBody) Validate() error {
	if s.ApiKey != nil {
		if err := s.ApiKey.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateApiKeyResponseBodyApiKey struct {
	// API Key ID。
	//
	// example:
	//
	// 1858636
	ApiKeyId *int64 `json:"apiKeyId,omitempty" xml:"apiKeyId,omitempty"`
	// The value of the API key.
	//
	// 	Notice: Copy your API key immediately. Store it securely. Anyone who obtains this key can initiate service requests on your behalf and incur charges. If you lose the key, you can reset it or create a new one.
	//
	// example:
	//
	// sk-ws-djI.8O7dkfkW2aICctnid4u4
	ApiKeyValue *string `json:"apiKeyValue,omitempty" xml:"apiKeyValue,omitempty"`
	// The creator.
	//
	// example:
	//
	// 1378030599924858
	CreatedBy *string `json:"createdBy,omitempty" xml:"createdBy,omitempty"`
	// The description.
	//
	// example:
	//
	// desc
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 1774338222000
	GmtCreate *int64 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// ws-3fa048e86117d91f
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s CreateApiKeyResponseBodyApiKey) String() string {
	return dara.Prettify(s)
}

func (s CreateApiKeyResponseBodyApiKey) GoString() string {
	return s.String()
}

func (s *CreateApiKeyResponseBodyApiKey) GetApiKeyId() *int64 {
	return s.ApiKeyId
}

func (s *CreateApiKeyResponseBodyApiKey) GetApiKeyValue() *string {
	return s.ApiKeyValue
}

func (s *CreateApiKeyResponseBodyApiKey) GetCreatedBy() *string {
	return s.CreatedBy
}

func (s *CreateApiKeyResponseBodyApiKey) GetDescription() *string {
	return s.Description
}

func (s *CreateApiKeyResponseBodyApiKey) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *CreateApiKeyResponseBodyApiKey) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateApiKeyResponseBodyApiKey) SetApiKeyId(v int64) *CreateApiKeyResponseBodyApiKey {
	s.ApiKeyId = &v
	return s
}

func (s *CreateApiKeyResponseBodyApiKey) SetApiKeyValue(v string) *CreateApiKeyResponseBodyApiKey {
	s.ApiKeyValue = &v
	return s
}

func (s *CreateApiKeyResponseBodyApiKey) SetCreatedBy(v string) *CreateApiKeyResponseBodyApiKey {
	s.CreatedBy = &v
	return s
}

func (s *CreateApiKeyResponseBodyApiKey) SetDescription(v string) *CreateApiKeyResponseBodyApiKey {
	s.Description = &v
	return s
}

func (s *CreateApiKeyResponseBodyApiKey) SetGmtCreate(v int64) *CreateApiKeyResponseBodyApiKey {
	s.GmtCreate = &v
	return s
}

func (s *CreateApiKeyResponseBodyApiKey) SetWorkspaceId(v string) *CreateApiKeyResponseBodyApiKey {
	s.WorkspaceId = &v
	return s
}

func (s *CreateApiKeyResponseBodyApiKey) Validate() error {
	return dara.Validate(s)
}
