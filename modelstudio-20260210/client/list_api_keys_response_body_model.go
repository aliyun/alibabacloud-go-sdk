// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApiKeysResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApiKeys(v []*ListApiKeysResponseBodyApiKeys) *ListApiKeysResponseBody
	GetApiKeys() []*ListApiKeysResponseBodyApiKeys
	SetCode(v string) *ListApiKeysResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListApiKeysResponseBody
	GetHttpStatusCode() *int32
	SetMaxResults(v int32) *ListApiKeysResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListApiKeysResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListApiKeysResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListApiKeysResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListApiKeysResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *ListApiKeysResponseBody
	GetTotalCount() *int32
}

type ListApiKeysResponseBody struct {
	// The list of API keys.
	ApiKeys []*ListApiKeysResponseBodyApiKeys `json:"apiKeys,omitempty" xml:"apiKeys,omitempty" type:"Repeated"`
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
	// The page size.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The token used to retrieve more results. You do not need to provide this parameter for the first query. For subsequent queries, use the token obtained from the previous response.
	//
	// example:
	//
	// lwytFRtLdNk=
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// Id of the request
	//
	// example:
	//
	// E4C14AE6-E987-5C2F-9230-9960AB48F4F2
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the API call was successful. Valid values:
	//
	// - true: Successful.
	//
	// - false: Failed.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// The total number of records.
	//
	// example:
	//
	// 4
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListApiKeysResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListApiKeysResponseBody) GoString() string {
	return s.String()
}

func (s *ListApiKeysResponseBody) GetApiKeys() []*ListApiKeysResponseBodyApiKeys {
	return s.ApiKeys
}

func (s *ListApiKeysResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListApiKeysResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListApiKeysResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListApiKeysResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListApiKeysResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListApiKeysResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListApiKeysResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListApiKeysResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListApiKeysResponseBody) SetApiKeys(v []*ListApiKeysResponseBodyApiKeys) *ListApiKeysResponseBody {
	s.ApiKeys = v
	return s
}

func (s *ListApiKeysResponseBody) SetCode(v string) *ListApiKeysResponseBody {
	s.Code = &v
	return s
}

func (s *ListApiKeysResponseBody) SetHttpStatusCode(v int32) *ListApiKeysResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListApiKeysResponseBody) SetMaxResults(v int32) *ListApiKeysResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListApiKeysResponseBody) SetMessage(v string) *ListApiKeysResponseBody {
	s.Message = &v
	return s
}

func (s *ListApiKeysResponseBody) SetNextToken(v string) *ListApiKeysResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListApiKeysResponseBody) SetRequestId(v string) *ListApiKeysResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListApiKeysResponseBody) SetSuccess(v bool) *ListApiKeysResponseBody {
	s.Success = &v
	return s
}

func (s *ListApiKeysResponseBody) SetTotalCount(v int32) *ListApiKeysResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListApiKeysResponseBody) Validate() error {
	if s.ApiKeys != nil {
		for _, item := range s.ApiKeys {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListApiKeysResponseBodyApiKeys struct {
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
	// sk-ws-djI.8O7d*****2aICctnid4u4
	ApiKeyValue *string `json:"apiKeyValue,omitempty" xml:"apiKeyValue,omitempty"`
	// The permission settings.
	Auth *ListApiKeysResponseBodyApiKeysAuth `json:"auth,omitempty" xml:"auth,omitempty" type:"Struct"`
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
	// test
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Indicates whether the API key is disabled. Valid values:
	//
	// - **0**: Active.
	//
	// - **1**: Disabled.
	//
	// example:
	//
	// 0
	Disabled *int32 `json:"disabled,omitempty" xml:"disabled,omitempty"`
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
	// ws-950f9aca7e76c816
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s ListApiKeysResponseBodyApiKeys) String() string {
	return dara.Prettify(s)
}

func (s ListApiKeysResponseBodyApiKeys) GoString() string {
	return s.String()
}

func (s *ListApiKeysResponseBodyApiKeys) GetApiKeyId() *int64 {
	return s.ApiKeyId
}

func (s *ListApiKeysResponseBodyApiKeys) GetApiKeyValue() *string {
	return s.ApiKeyValue
}

func (s *ListApiKeysResponseBodyApiKeys) GetAuth() *ListApiKeysResponseBodyApiKeysAuth {
	return s.Auth
}

func (s *ListApiKeysResponseBodyApiKeys) GetCreatedBy() *string {
	return s.CreatedBy
}

func (s *ListApiKeysResponseBodyApiKeys) GetDescription() *string {
	return s.Description
}

func (s *ListApiKeysResponseBodyApiKeys) GetDisabled() *int32 {
	return s.Disabled
}

func (s *ListApiKeysResponseBodyApiKeys) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *ListApiKeysResponseBodyApiKeys) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListApiKeysResponseBodyApiKeys) SetApiKeyId(v int64) *ListApiKeysResponseBodyApiKeys {
	s.ApiKeyId = &v
	return s
}

func (s *ListApiKeysResponseBodyApiKeys) SetApiKeyValue(v string) *ListApiKeysResponseBodyApiKeys {
	s.ApiKeyValue = &v
	return s
}

func (s *ListApiKeysResponseBodyApiKeys) SetAuth(v *ListApiKeysResponseBodyApiKeysAuth) *ListApiKeysResponseBodyApiKeys {
	s.Auth = v
	return s
}

func (s *ListApiKeysResponseBodyApiKeys) SetCreatedBy(v string) *ListApiKeysResponseBodyApiKeys {
	s.CreatedBy = &v
	return s
}

func (s *ListApiKeysResponseBodyApiKeys) SetDescription(v string) *ListApiKeysResponseBodyApiKeys {
	s.Description = &v
	return s
}

func (s *ListApiKeysResponseBodyApiKeys) SetDisabled(v int32) *ListApiKeysResponseBodyApiKeys {
	s.Disabled = &v
	return s
}

func (s *ListApiKeysResponseBodyApiKeys) SetGmtCreate(v int64) *ListApiKeysResponseBodyApiKeys {
	s.GmtCreate = &v
	return s
}

func (s *ListApiKeysResponseBodyApiKeys) SetWorkspaceId(v string) *ListApiKeysResponseBodyApiKeys {
	s.WorkspaceId = &v
	return s
}

func (s *ListApiKeysResponseBodyApiKeys) Validate() error {
	if s.Auth != nil {
		if err := s.Auth.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListApiKeysResponseBodyApiKeysAuth struct {
	// The IP address whitelist.
	AccessIps        []*string                                           `json:"accessIps,omitempty" xml:"accessIps,omitempty" type:"Repeated"`
	ModelAccessScope *ListApiKeysResponseBodyApiKeysAuthModelAccessScope `json:"modelAccessScope,omitempty" xml:"modelAccessScope,omitempty" type:"Struct"`
	// All: all permissions. Custom: custom permissions.
	//
	// example:
	//
	// Custom
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListApiKeysResponseBodyApiKeysAuth) String() string {
	return dara.Prettify(s)
}

func (s ListApiKeysResponseBodyApiKeysAuth) GoString() string {
	return s.String()
}

func (s *ListApiKeysResponseBodyApiKeysAuth) GetAccessIps() []*string {
	return s.AccessIps
}

func (s *ListApiKeysResponseBodyApiKeysAuth) GetModelAccessScope() *ListApiKeysResponseBodyApiKeysAuthModelAccessScope {
	return s.ModelAccessScope
}

func (s *ListApiKeysResponseBodyApiKeysAuth) GetType() *string {
	return s.Type
}

func (s *ListApiKeysResponseBodyApiKeysAuth) SetAccessIps(v []*string) *ListApiKeysResponseBodyApiKeysAuth {
	s.AccessIps = v
	return s
}

func (s *ListApiKeysResponseBodyApiKeysAuth) SetModelAccessScope(v *ListApiKeysResponseBodyApiKeysAuthModelAccessScope) *ListApiKeysResponseBodyApiKeysAuth {
	s.ModelAccessScope = v
	return s
}

func (s *ListApiKeysResponseBodyApiKeysAuth) SetType(v string) *ListApiKeysResponseBodyApiKeysAuth {
	s.Type = &v
	return s
}

func (s *ListApiKeysResponseBodyApiKeysAuth) Validate() error {
	if s.ModelAccessScope != nil {
		if err := s.ModelAccessScope.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListApiKeysResponseBodyApiKeysAuthModelAccessScope struct {
	AccessibleModels []*string `json:"accessibleModels,omitempty" xml:"accessibleModels,omitempty" type:"Repeated"`
	AllowAllModels   *bool     `json:"allowAllModels,omitempty" xml:"allowAllModels,omitempty"`
}

func (s ListApiKeysResponseBodyApiKeysAuthModelAccessScope) String() string {
	return dara.Prettify(s)
}

func (s ListApiKeysResponseBodyApiKeysAuthModelAccessScope) GoString() string {
	return s.String()
}

func (s *ListApiKeysResponseBodyApiKeysAuthModelAccessScope) GetAccessibleModels() []*string {
	return s.AccessibleModels
}

func (s *ListApiKeysResponseBodyApiKeysAuthModelAccessScope) GetAllowAllModels() *bool {
	return s.AllowAllModels
}

func (s *ListApiKeysResponseBodyApiKeysAuthModelAccessScope) SetAccessibleModels(v []*string) *ListApiKeysResponseBodyApiKeysAuthModelAccessScope {
	s.AccessibleModels = v
	return s
}

func (s *ListApiKeysResponseBodyApiKeysAuthModelAccessScope) SetAllowAllModels(v bool) *ListApiKeysResponseBodyApiKeysAuthModelAccessScope {
	s.AllowAllModels = &v
	return s
}

func (s *ListApiKeysResponseBodyApiKeysAuthModelAccessScope) Validate() error {
	return dara.Validate(s)
}
