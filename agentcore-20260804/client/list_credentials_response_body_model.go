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
	// The business status code.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The list of credentials.
	Items []*ListCredentialsResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// The maximum number of records per page that takes effect for this query.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The response message. An error description is returned if the request fails.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The pagination token for the next page. This value is empty if there is no next page.
	//
	// example:
	//
	// 10
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// request-123456
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// The total number of credentials that match the query conditions.
	//
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
	// The number of agents bound to this credential.
	//
	// example:
	//
	// 2
	BoundAgentsCounts *int32 `json:"boundAgentsCounts,omitempty" xml:"boundAgentsCounts,omitempty"`
	// The creation time in UTC, formatted according to RFC 3339.
	//
	// example:
	//
	// 2026-08-12T03:04:05Z
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// The credential ID.
	//
	// example:
	//
	// cred-123456
	CredentialId *string `json:"credentialId,omitempty" xml:"credentialId,omitempty"`
	// The masked content of the credential. When credentialType is apiKey, the apiKey value is returned as asterisks (*) of equal length.
	//
	// example:
	//
	// {"apiKey":"****************"}
	CredentialMetadata *string `json:"credentialMetadata,omitempty" xml:"credentialMetadata,omitempty"`
	// The credential type. Currently, only apiKey is supported.
	//
	// example:
	//
	// apiKey
	CredentialType *string `json:"credentialType,omitempty" xml:"credentialType,omitempty"`
	// The credential description. The description can be up to 256 characters in length.
	//
	// example:
	//
	// API Key used for calling model services in the production environment
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The credential name. The name must be unique within the workspace and can contain only letters, digits, periods (.), underscores (_), and hyphens (-). The name must be 3 to 128 characters in length and cannot use runtime reserved names.
	//
	// example:
	//
	// model-api-key
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The region ID where the resource resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The list of resources to which the credential can be applied.
	ResourceRefs []*ListCredentialsResponseBodyItemsResourceRefs `json:"resourceRefs,omitempty" xml:"resourceRefs,omitempty" type:"Repeated"`
	// The resource scope of the credential.
	//
	// example:
	//
	// ALL
	ResourceScope *string `json:"resourceScope,omitempty" xml:"resourceScope,omitempty"`
	// The time of the last modification in UTC, formatted according to RFC 3339.
	//
	// example:
	//
	// 2026-08-12T03:04:05Z
	UpdatedAt *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	// The workspace ID.
	//
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

func (s *ListCredentialsResponseBodyItems) GetResourceRefs() []*ListCredentialsResponseBodyItemsResourceRefs {
	return s.ResourceRefs
}

func (s *ListCredentialsResponseBodyItems) GetResourceScope() *string {
	return s.ResourceScope
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

func (s *ListCredentialsResponseBodyItems) SetResourceRefs(v []*ListCredentialsResponseBodyItemsResourceRefs) *ListCredentialsResponseBodyItems {
	s.ResourceRefs = v
	return s
}

func (s *ListCredentialsResponseBodyItems) SetResourceScope(v string) *ListCredentialsResponseBodyItems {
	s.ResourceScope = &v
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
	if s.ResourceRefs != nil {
		for _, item := range s.ResourceRefs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListCredentialsResponseBodyItemsResourceRefs struct {
	// The unique identifier of the resource.
	//
	// example:
	//
	// agent-xxxx
	ResourceId *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	// The resource name. This value is empty if the resource has been deleted.
	//
	// example:
	//
	// my-agent
	ResourceName *string `json:"resourceName,omitempty" xml:"resourceName,omitempty"`
	// The resource type, such as agent.
	//
	// example:
	//
	// agent
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
}

func (s ListCredentialsResponseBodyItemsResourceRefs) String() string {
	return dara.Prettify(s)
}

func (s ListCredentialsResponseBodyItemsResourceRefs) GoString() string {
	return s.String()
}

func (s *ListCredentialsResponseBodyItemsResourceRefs) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListCredentialsResponseBodyItemsResourceRefs) GetResourceName() *string {
	return s.ResourceName
}

func (s *ListCredentialsResponseBodyItemsResourceRefs) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListCredentialsResponseBodyItemsResourceRefs) SetResourceId(v string) *ListCredentialsResponseBodyItemsResourceRefs {
	s.ResourceId = &v
	return s
}

func (s *ListCredentialsResponseBodyItemsResourceRefs) SetResourceName(v string) *ListCredentialsResponseBodyItemsResourceRefs {
	s.ResourceName = &v
	return s
}

func (s *ListCredentialsResponseBodyItemsResourceRefs) SetResourceType(v string) *ListCredentialsResponseBodyItemsResourceRefs {
	s.ResourceType = &v
	return s
}

func (s *ListCredentialsResponseBodyItemsResourceRefs) Validate() error {
	return dara.Validate(s)
}
