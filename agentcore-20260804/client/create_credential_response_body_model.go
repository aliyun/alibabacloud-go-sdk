// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCredentialResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateCredentialResponseBody
	GetCode() *string
	SetData(v *CreateCredentialResponseBodyData) *CreateCredentialResponseBody
	GetData() *CreateCredentialResponseBodyData
	SetHttpStatusCode(v int32) *CreateCredentialResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateCredentialResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateCredentialResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateCredentialResponseBody
	GetSuccess() *bool
}

type CreateCredentialResponseBody struct {
	// The business status code.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The created credential information.
	Data *CreateCredentialResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The response message. An error description is returned if the request fails.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// request-123456
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateCredentialResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCredentialResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCredentialResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateCredentialResponseBody) GetData() *CreateCredentialResponseBodyData {
	return s.Data
}

func (s *CreateCredentialResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateCredentialResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateCredentialResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCredentialResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateCredentialResponseBody) SetCode(v string) *CreateCredentialResponseBody {
	s.Code = &v
	return s
}

func (s *CreateCredentialResponseBody) SetData(v *CreateCredentialResponseBodyData) *CreateCredentialResponseBody {
	s.Data = v
	return s
}

func (s *CreateCredentialResponseBody) SetHttpStatusCode(v int32) *CreateCredentialResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateCredentialResponseBody) SetMessage(v string) *CreateCredentialResponseBody {
	s.Message = &v
	return s
}

func (s *CreateCredentialResponseBody) SetRequestId(v string) *CreateCredentialResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCredentialResponseBody) SetSuccess(v bool) *CreateCredentialResponseBody {
	s.Success = &v
	return s
}

func (s *CreateCredentialResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateCredentialResponseBodyData struct {
	// The creation time in UTC, formatted in RFC 3339.
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
	// The masked credential content. When credentialType is apiKey, the apiKey value is returned as asterisks (*) of equal length.
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
	// Each item contains resourceType, resourceId, and resourceName. resourceName is empty if the resource has been deleted.
	ResourceRefs []*CreateCredentialResponseBodyDataResourceRefs `json:"resourceRefs,omitempty" xml:"resourceRefs,omitempty" type:"Repeated"`
	// The credential resource scope.
	//
	// example:
	//
	// ALL
	ResourceScope *string `json:"resourceScope,omitempty" xml:"resourceScope,omitempty"`
	// The time of the last modification in UTC, formatted in RFC 3339.
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

func (s CreateCredentialResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateCredentialResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateCredentialResponseBodyData) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *CreateCredentialResponseBodyData) GetCredentialId() *string {
	return s.CredentialId
}

func (s *CreateCredentialResponseBodyData) GetCredentialMetadata() *string {
	return s.CredentialMetadata
}

func (s *CreateCredentialResponseBodyData) GetCredentialType() *string {
	return s.CredentialType
}

func (s *CreateCredentialResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *CreateCredentialResponseBodyData) GetName() *string {
	return s.Name
}

func (s *CreateCredentialResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateCredentialResponseBodyData) GetResourceRefs() []*CreateCredentialResponseBodyDataResourceRefs {
	return s.ResourceRefs
}

func (s *CreateCredentialResponseBodyData) GetResourceScope() *string {
	return s.ResourceScope
}

func (s *CreateCredentialResponseBodyData) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *CreateCredentialResponseBodyData) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateCredentialResponseBodyData) SetCreatedAt(v string) *CreateCredentialResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *CreateCredentialResponseBodyData) SetCredentialId(v string) *CreateCredentialResponseBodyData {
	s.CredentialId = &v
	return s
}

func (s *CreateCredentialResponseBodyData) SetCredentialMetadata(v string) *CreateCredentialResponseBodyData {
	s.CredentialMetadata = &v
	return s
}

func (s *CreateCredentialResponseBodyData) SetCredentialType(v string) *CreateCredentialResponseBodyData {
	s.CredentialType = &v
	return s
}

func (s *CreateCredentialResponseBodyData) SetDescription(v string) *CreateCredentialResponseBodyData {
	s.Description = &v
	return s
}

func (s *CreateCredentialResponseBodyData) SetName(v string) *CreateCredentialResponseBodyData {
	s.Name = &v
	return s
}

func (s *CreateCredentialResponseBodyData) SetRegionId(v string) *CreateCredentialResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *CreateCredentialResponseBodyData) SetResourceRefs(v []*CreateCredentialResponseBodyDataResourceRefs) *CreateCredentialResponseBodyData {
	s.ResourceRefs = v
	return s
}

func (s *CreateCredentialResponseBodyData) SetResourceScope(v string) *CreateCredentialResponseBodyData {
	s.ResourceScope = &v
	return s
}

func (s *CreateCredentialResponseBodyData) SetUpdatedAt(v string) *CreateCredentialResponseBodyData {
	s.UpdatedAt = &v
	return s
}

func (s *CreateCredentialResponseBodyData) SetWorkspaceId(v string) *CreateCredentialResponseBodyData {
	s.WorkspaceId = &v
	return s
}

func (s *CreateCredentialResponseBodyData) Validate() error {
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

type CreateCredentialResponseBodyDataResourceRefs struct {
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

func (s CreateCredentialResponseBodyDataResourceRefs) String() string {
	return dara.Prettify(s)
}

func (s CreateCredentialResponseBodyDataResourceRefs) GoString() string {
	return s.String()
}

func (s *CreateCredentialResponseBodyDataResourceRefs) GetResourceId() *string {
	return s.ResourceId
}

func (s *CreateCredentialResponseBodyDataResourceRefs) GetResourceName() *string {
	return s.ResourceName
}

func (s *CreateCredentialResponseBodyDataResourceRefs) GetResourceType() *string {
	return s.ResourceType
}

func (s *CreateCredentialResponseBodyDataResourceRefs) SetResourceId(v string) *CreateCredentialResponseBodyDataResourceRefs {
	s.ResourceId = &v
	return s
}

func (s *CreateCredentialResponseBodyDataResourceRefs) SetResourceName(v string) *CreateCredentialResponseBodyDataResourceRefs {
	s.ResourceName = &v
	return s
}

func (s *CreateCredentialResponseBodyDataResourceRefs) SetResourceType(v string) *CreateCredentialResponseBodyDataResourceRefs {
	s.ResourceType = &v
	return s
}

func (s *CreateCredentialResponseBodyDataResourceRefs) Validate() error {
	return dara.Validate(s)
}
