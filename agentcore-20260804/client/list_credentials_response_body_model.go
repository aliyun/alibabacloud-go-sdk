// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCredentialsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListCredentialsResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListCredentialsResponseBody
	GetHttpStatusCode() *int32
	SetItems(v []*ListCredentialsResponseBodyItems) *ListCredentialsResponseBody
	GetItems() []*ListCredentialsResponseBodyItems
	SetMaxResults(v int32) *ListCredentialsResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListCredentialsResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListCredentialsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListCredentialsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListCredentialsResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *ListCredentialsResponseBody
	GetTotalCount() *int64
}

type ListCredentialsResponseBody struct {
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32                              `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	Items          []*ListCredentialsResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 10
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// request-123456
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 42
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListCredentialsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCredentialsResponseBody) GoString() string {
	return s.String()
}

func (s *ListCredentialsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListCredentialsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListCredentialsResponseBody) GetItems() []*ListCredentialsResponseBodyItems {
	return s.Items
}

func (s *ListCredentialsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListCredentialsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListCredentialsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListCredentialsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCredentialsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListCredentialsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListCredentialsResponseBody) SetCode(v string) *ListCredentialsResponseBody {
	s.Code = &v
	return s
}

func (s *ListCredentialsResponseBody) SetHttpStatusCode(v int32) *ListCredentialsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListCredentialsResponseBody) SetItems(v []*ListCredentialsResponseBodyItems) *ListCredentialsResponseBody {
	s.Items = v
	return s
}

func (s *ListCredentialsResponseBody) SetMaxResults(v int32) *ListCredentialsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListCredentialsResponseBody) SetMessage(v string) *ListCredentialsResponseBody {
	s.Message = &v
	return s
}

func (s *ListCredentialsResponseBody) SetNextToken(v string) *ListCredentialsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListCredentialsResponseBody) SetRequestId(v string) *ListCredentialsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCredentialsResponseBody) SetSuccess(v bool) *ListCredentialsResponseBody {
	s.Success = &v
	return s
}

func (s *ListCredentialsResponseBody) SetTotalCount(v int64) *ListCredentialsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListCredentialsResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListCredentialsResponseBodyItems struct {
	// example:
	//
	// 2
	BoundAgentsCounts *int32 `json:"boundAgentsCounts,omitempty" xml:"boundAgentsCounts,omitempty"`
	// example:
	//
	// 2026-08-12T03:04:05Z
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// example:
	//
	// cred-123456
	CredentialId *string `json:"credentialId,omitempty" xml:"credentialId,omitempty"`
	// example:
	//
	// {"apiKey":"****************"}
	CredentialMetadata *string `json:"credentialMetadata,omitempty" xml:"credentialMetadata,omitempty"`
	// example:
	//
	// apiKey
	CredentialType *string `json:"credentialType,omitempty" xml:"credentialType,omitempty"`
	// example:
	//
	// 线上环境调用模型服务使用的 API Key
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// model-api-key
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// example:
	//
	// 2026-08-12T03:04:05Z
	UpdatedAt *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	// example:
	//
	// ws-123456
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s ListCredentialsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListCredentialsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListCredentialsResponseBodyItems) GetBoundAgentsCounts() *int32 {
	return s.BoundAgentsCounts
}

func (s *ListCredentialsResponseBodyItems) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *ListCredentialsResponseBodyItems) GetCredentialId() *string {
	return s.CredentialId
}

func (s *ListCredentialsResponseBodyItems) GetCredentialMetadata() *string {
	return s.CredentialMetadata
}

func (s *ListCredentialsResponseBodyItems) GetCredentialType() *string {
	return s.CredentialType
}

func (s *ListCredentialsResponseBodyItems) GetDescription() *string {
	return s.Description
}

func (s *ListCredentialsResponseBodyItems) GetName() *string {
	return s.Name
}

func (s *ListCredentialsResponseBodyItems) GetRegionId() *string {
	return s.RegionId
}

func (s *ListCredentialsResponseBodyItems) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *ListCredentialsResponseBodyItems) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListCredentialsResponseBodyItems) SetBoundAgentsCounts(v int32) *ListCredentialsResponseBodyItems {
	s.BoundAgentsCounts = &v
	return s
}

func (s *ListCredentialsResponseBodyItems) SetCreatedAt(v string) *ListCredentialsResponseBodyItems {
	s.CreatedAt = &v
	return s
}

func (s *ListCredentialsResponseBodyItems) SetCredentialId(v string) *ListCredentialsResponseBodyItems {
	s.CredentialId = &v
	return s
}

func (s *ListCredentialsResponseBodyItems) SetCredentialMetadata(v string) *ListCredentialsResponseBodyItems {
	s.CredentialMetadata = &v
	return s
}

func (s *ListCredentialsResponseBodyItems) SetCredentialType(v string) *ListCredentialsResponseBodyItems {
	s.CredentialType = &v
	return s
}

func (s *ListCredentialsResponseBodyItems) SetDescription(v string) *ListCredentialsResponseBodyItems {
	s.Description = &v
	return s
}

func (s *ListCredentialsResponseBodyItems) SetName(v string) *ListCredentialsResponseBodyItems {
	s.Name = &v
	return s
}

func (s *ListCredentialsResponseBodyItems) SetRegionId(v string) *ListCredentialsResponseBodyItems {
	s.RegionId = &v
	return s
}

func (s *ListCredentialsResponseBodyItems) SetUpdatedAt(v string) *ListCredentialsResponseBodyItems {
	s.UpdatedAt = &v
	return s
}

func (s *ListCredentialsResponseBodyItems) SetWorkspaceId(v string) *ListCredentialsResponseBodyItems {
	s.WorkspaceId = &v
	return s
}

func (s *ListCredentialsResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
