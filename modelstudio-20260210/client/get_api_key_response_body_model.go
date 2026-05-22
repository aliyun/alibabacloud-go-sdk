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
	// DFD55E7B-0615-5343-BDA0-FBEE86F29935
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
	// 2965964
	ApiKeyId *int64 `json:"apiKeyId,omitempty" xml:"apiKeyId,omitempty"`
	// example:
	//
	// sk-ws-djI.mhU0D****testtestest
	ApiKeyValue *string                          `json:"apiKeyValue,omitempty" xml:"apiKeyValue,omitempty"`
	Auth        *GetApiKeyResponseBodyApiKeyAuth `json:"auth,omitempty" xml:"auth,omitempty" type:"Struct"`
	// example:
	//
	// 1378030599924858
	CreatedBy *string `json:"createdBy,omitempty" xml:"createdBy,omitempty"`
	// example:
	//
	// v7
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// 0
	Disabled *int32 `json:"disabled,omitempty" xml:"disabled,omitempty"`
	// example:
	//
	// 1774338222000
	GmtCreate *int64 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// ws-b2d30f148c236908
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

func (s *GetApiKeyResponseBodyApiKey) GetAuth() *GetApiKeyResponseBodyApiKeyAuth {
	return s.Auth
}

func (s *GetApiKeyResponseBodyApiKey) GetCreatedBy() *string {
	return s.CreatedBy
}

func (s *GetApiKeyResponseBodyApiKey) GetDescription() *string {
	return s.Description
}

func (s *GetApiKeyResponseBodyApiKey) GetDisabled() *int32 {
	return s.Disabled
}

func (s *GetApiKeyResponseBodyApiKey) GetGmtCreate() *int64 {
	return s.GmtCreate
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

func (s *GetApiKeyResponseBodyApiKey) SetAuth(v *GetApiKeyResponseBodyApiKeyAuth) *GetApiKeyResponseBodyApiKey {
	s.Auth = v
	return s
}

func (s *GetApiKeyResponseBodyApiKey) SetCreatedBy(v string) *GetApiKeyResponseBodyApiKey {
	s.CreatedBy = &v
	return s
}

func (s *GetApiKeyResponseBodyApiKey) SetDescription(v string) *GetApiKeyResponseBodyApiKey {
	s.Description = &v
	return s
}

func (s *GetApiKeyResponseBodyApiKey) SetDisabled(v int32) *GetApiKeyResponseBodyApiKey {
	s.Disabled = &v
	return s
}

func (s *GetApiKeyResponseBodyApiKey) SetGmtCreate(v int64) *GetApiKeyResponseBodyApiKey {
	s.GmtCreate = &v
	return s
}

func (s *GetApiKeyResponseBodyApiKey) SetWorkspaceId(v string) *GetApiKeyResponseBodyApiKey {
	s.WorkspaceId = &v
	return s
}

func (s *GetApiKeyResponseBodyApiKey) Validate() error {
	if s.Auth != nil {
		if err := s.Auth.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetApiKeyResponseBodyApiKeyAuth struct {
	AccessIps []*string `json:"accessIps,omitempty" xml:"accessIps,omitempty" type:"Repeated"`
	// example:
	//
	// All
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetApiKeyResponseBodyApiKeyAuth) String() string {
	return dara.Prettify(s)
}

func (s GetApiKeyResponseBodyApiKeyAuth) GoString() string {
	return s.String()
}

func (s *GetApiKeyResponseBodyApiKeyAuth) GetAccessIps() []*string {
	return s.AccessIps
}

func (s *GetApiKeyResponseBodyApiKeyAuth) GetType() *string {
	return s.Type
}

func (s *GetApiKeyResponseBodyApiKeyAuth) SetAccessIps(v []*string) *GetApiKeyResponseBodyApiKeyAuth {
	s.AccessIps = v
	return s
}

func (s *GetApiKeyResponseBodyApiKeyAuth) SetType(v string) *GetApiKeyResponseBodyApiKeyAuth {
	s.Type = &v
	return s
}

func (s *GetApiKeyResponseBodyApiKeyAuth) Validate() error {
	return dara.Validate(s)
}
