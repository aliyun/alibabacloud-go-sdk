// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCredentialResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateCredentialResponseBody
	GetCode() *string
	SetData(v *UpdateCredentialResponseBodyData) *UpdateCredentialResponseBody
	GetData() *UpdateCredentialResponseBodyData
	SetHttpStatusCode(v int32) *UpdateCredentialResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateCredentialResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateCredentialResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateCredentialResponseBody
	GetSuccess() *bool
}

type UpdateCredentialResponseBody struct {
	// The business status code.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The updated credential information.
	Data *UpdateCredentialResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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

func (s UpdateCredentialResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateCredentialResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCredentialResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateCredentialResponseBody) GetData() *UpdateCredentialResponseBodyData {
	return s.Data
}

func (s *UpdateCredentialResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateCredentialResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateCredentialResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateCredentialResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateCredentialResponseBody) SetCode(v string) *UpdateCredentialResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateCredentialResponseBody) SetData(v *UpdateCredentialResponseBodyData) *UpdateCredentialResponseBody {
	s.Data = v
	return s
}

func (s *UpdateCredentialResponseBody) SetHttpStatusCode(v int32) *UpdateCredentialResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateCredentialResponseBody) SetMessage(v string) *UpdateCredentialResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateCredentialResponseBody) SetRequestId(v string) *UpdateCredentialResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCredentialResponseBody) SetSuccess(v bool) *UpdateCredentialResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateCredentialResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateCredentialResponseBodyData struct {
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
	// The masked credential content. If credentialType is apiKey, the apiKey value is returned as asterisks (*) of equal length.
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
	// The region ID of the resource.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// Each item contains resourceType, resourceId, and resourceName. resourceName is empty if the resource has been deleted.
	ResourceRefs []*UpdateCredentialResponseBodyDataResourceRefs `json:"resourceRefs,omitempty" xml:"resourceRefs,omitempty" type:"Repeated"`
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

func (s UpdateCredentialResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateCredentialResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateCredentialResponseBodyData) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *UpdateCredentialResponseBodyData) GetCredentialId() *string {
	return s.CredentialId
}

func (s *UpdateCredentialResponseBodyData) GetCredentialMetadata() *string {
	return s.CredentialMetadata
}

func (s *UpdateCredentialResponseBodyData) GetCredentialType() *string {
	return s.CredentialType
}

func (s *UpdateCredentialResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *UpdateCredentialResponseBodyData) GetName() *string {
	return s.Name
}

func (s *UpdateCredentialResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateCredentialResponseBodyData) GetResourceRefs() []*UpdateCredentialResponseBodyDataResourceRefs {
	return s.ResourceRefs
}

func (s *UpdateCredentialResponseBodyData) GetResourceScope() *string {
	return s.ResourceScope
}

func (s *UpdateCredentialResponseBodyData) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *UpdateCredentialResponseBodyData) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *UpdateCredentialResponseBodyData) SetCreatedAt(v string) *UpdateCredentialResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *UpdateCredentialResponseBodyData) SetCredentialId(v string) *UpdateCredentialResponseBodyData {
	s.CredentialId = &v
	return s
}

func (s *UpdateCredentialResponseBodyData) SetCredentialMetadata(v string) *UpdateCredentialResponseBodyData {
	s.CredentialMetadata = &v
	return s
}

func (s *UpdateCredentialResponseBodyData) SetCredentialType(v string) *UpdateCredentialResponseBodyData {
	s.CredentialType = &v
	return s
}

func (s *UpdateCredentialResponseBodyData) SetDescription(v string) *UpdateCredentialResponseBodyData {
	s.Description = &v
	return s
}

func (s *UpdateCredentialResponseBodyData) SetName(v string) *UpdateCredentialResponseBodyData {
	s.Name = &v
	return s
}

func (s *UpdateCredentialResponseBodyData) SetRegionId(v string) *UpdateCredentialResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *UpdateCredentialResponseBodyData) SetResourceRefs(v []*UpdateCredentialResponseBodyDataResourceRefs) *UpdateCredentialResponseBodyData {
	s.ResourceRefs = v
	return s
}

func (s *UpdateCredentialResponseBodyData) SetResourceScope(v string) *UpdateCredentialResponseBodyData {
	s.ResourceScope = &v
	return s
}

func (s *UpdateCredentialResponseBodyData) SetUpdatedAt(v string) *UpdateCredentialResponseBodyData {
	s.UpdatedAt = &v
	return s
}

func (s *UpdateCredentialResponseBodyData) SetWorkspaceId(v string) *UpdateCredentialResponseBodyData {
	s.WorkspaceId = &v
	return s
}

func (s *UpdateCredentialResponseBodyData) Validate() error {
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

type UpdateCredentialResponseBodyDataResourceRefs struct {
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

func (s UpdateCredentialResponseBodyDataResourceRefs) String() string {
	return dara.Prettify(s)
}

func (s UpdateCredentialResponseBodyDataResourceRefs) GoString() string {
	return s.String()
}

func (s *UpdateCredentialResponseBodyDataResourceRefs) GetResourceId() *string {
	return s.ResourceId
}

func (s *UpdateCredentialResponseBodyDataResourceRefs) GetResourceName() *string {
	return s.ResourceName
}

func (s *UpdateCredentialResponseBodyDataResourceRefs) GetResourceType() *string {
	return s.ResourceType
}

func (s *UpdateCredentialResponseBodyDataResourceRefs) SetResourceId(v string) *UpdateCredentialResponseBodyDataResourceRefs {
	s.ResourceId = &v
	return s
}

func (s *UpdateCredentialResponseBodyDataResourceRefs) SetResourceName(v string) *UpdateCredentialResponseBodyDataResourceRefs {
	s.ResourceName = &v
	return s
}

func (s *UpdateCredentialResponseBodyDataResourceRefs) SetResourceType(v string) *UpdateCredentialResponseBodyDataResourceRefs {
	s.ResourceType = &v
	return s
}

func (s *UpdateCredentialResponseBodyDataResourceRefs) Validate() error {
	return dara.Validate(s)
}
