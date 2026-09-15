// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCredentialResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetCredentialResponseBody
	GetCode() *string
	SetData(v *GetCredentialResponseBodyData) *GetCredentialResponseBody
	GetData() *GetCredentialResponseBodyData
	SetHttpStatusCode(v int32) *GetCredentialResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetCredentialResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetCredentialResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetCredentialResponseBody
	GetSuccess() *bool
}

type GetCredentialResponseBody struct {
	// The business status code.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The credential details.
	Data *GetCredentialResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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

func (s GetCredentialResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCredentialResponseBody) GoString() string {
	return s.String()
}

func (s *GetCredentialResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetCredentialResponseBody) GetData() *GetCredentialResponseBodyData {
	return s.Data
}

func (s *GetCredentialResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetCredentialResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetCredentialResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCredentialResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetCredentialResponseBody) SetCode(v string) *GetCredentialResponseBody {
	s.Code = &v
	return s
}

func (s *GetCredentialResponseBody) SetData(v *GetCredentialResponseBodyData) *GetCredentialResponseBody {
	s.Data = v
	return s
}

func (s *GetCredentialResponseBody) SetHttpStatusCode(v int32) *GetCredentialResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetCredentialResponseBody) SetMessage(v string) *GetCredentialResponseBody {
	s.Message = &v
	return s
}

func (s *GetCredentialResponseBody) SetRequestId(v string) *GetCredentialResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCredentialResponseBody) SetSuccess(v bool) *GetCredentialResponseBody {
	s.Success = &v
	return s
}

func (s *GetCredentialResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetCredentialResponseBodyData struct {
	// The list of agents bound to the credential.
	BoundAgents []*GetCredentialResponseBodyDataBoundAgents `json:"boundAgents,omitempty" xml:"boundAgents,omitempty" type:"Repeated"`
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
	// The masked content of the credential. When credentialType is apiKey, the value of apiKey is returned as asterisks (*) of equal length.
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
	// The credential description, up to 256 characters in length.
	//
	// example:
	//
	// API Key used to call model services in the production environment
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
	// Each item contains resourceType, resourceId, and resourceName. If the resource has been deleted, resourceName is empty.
	ResourceRefs []*GetCredentialResponseBodyDataResourceRefs `json:"resourceRefs,omitempty" xml:"resourceRefs,omitempty" type:"Repeated"`
	// The scope of resources to which the credential applies.
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

func (s GetCredentialResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetCredentialResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCredentialResponseBodyData) GetBoundAgents() []*GetCredentialResponseBodyDataBoundAgents {
	return s.BoundAgents
}

func (s *GetCredentialResponseBodyData) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *GetCredentialResponseBodyData) GetCredentialId() *string {
	return s.CredentialId
}

func (s *GetCredentialResponseBodyData) GetCredentialMetadata() *string {
	return s.CredentialMetadata
}

func (s *GetCredentialResponseBodyData) GetCredentialType() *string {
	return s.CredentialType
}

func (s *GetCredentialResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *GetCredentialResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetCredentialResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *GetCredentialResponseBodyData) GetResourceRefs() []*GetCredentialResponseBodyDataResourceRefs {
	return s.ResourceRefs
}

func (s *GetCredentialResponseBodyData) GetResourceScope() *string {
	return s.ResourceScope
}

func (s *GetCredentialResponseBodyData) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *GetCredentialResponseBodyData) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetCredentialResponseBodyData) SetBoundAgents(v []*GetCredentialResponseBodyDataBoundAgents) *GetCredentialResponseBodyData {
	s.BoundAgents = v
	return s
}

func (s *GetCredentialResponseBodyData) SetCreatedAt(v string) *GetCredentialResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *GetCredentialResponseBodyData) SetCredentialId(v string) *GetCredentialResponseBodyData {
	s.CredentialId = &v
	return s
}

func (s *GetCredentialResponseBodyData) SetCredentialMetadata(v string) *GetCredentialResponseBodyData {
	s.CredentialMetadata = &v
	return s
}

func (s *GetCredentialResponseBodyData) SetCredentialType(v string) *GetCredentialResponseBodyData {
	s.CredentialType = &v
	return s
}

func (s *GetCredentialResponseBodyData) SetDescription(v string) *GetCredentialResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetCredentialResponseBodyData) SetName(v string) *GetCredentialResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetCredentialResponseBodyData) SetRegionId(v string) *GetCredentialResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *GetCredentialResponseBodyData) SetResourceRefs(v []*GetCredentialResponseBodyDataResourceRefs) *GetCredentialResponseBodyData {
	s.ResourceRefs = v
	return s
}

func (s *GetCredentialResponseBodyData) SetResourceScope(v string) *GetCredentialResponseBodyData {
	s.ResourceScope = &v
	return s
}

func (s *GetCredentialResponseBodyData) SetUpdatedAt(v string) *GetCredentialResponseBodyData {
	s.UpdatedAt = &v
	return s
}

func (s *GetCredentialResponseBodyData) SetWorkspaceId(v string) *GetCredentialResponseBodyData {
	s.WorkspaceId = &v
	return s
}

func (s *GetCredentialResponseBodyData) Validate() error {
	if s.BoundAgents != nil {
		for _, item := range s.BoundAgents {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
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

type GetCredentialResponseBodyDataBoundAgents struct {
	// The agent ID.
	//
	// example:
	//
	// agent-123456
	AgentId *string `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// The agent name.
	//
	// example:
	//
	// agent-01
	AgentName *string `json:"agentName,omitempty" xml:"agentName,omitempty"`
}

func (s GetCredentialResponseBodyDataBoundAgents) String() string {
	return dara.Prettify(s)
}

func (s GetCredentialResponseBodyDataBoundAgents) GoString() string {
	return s.String()
}

func (s *GetCredentialResponseBodyDataBoundAgents) GetAgentId() *string {
	return s.AgentId
}

func (s *GetCredentialResponseBodyDataBoundAgents) GetAgentName() *string {
	return s.AgentName
}

func (s *GetCredentialResponseBodyDataBoundAgents) SetAgentId(v string) *GetCredentialResponseBodyDataBoundAgents {
	s.AgentId = &v
	return s
}

func (s *GetCredentialResponseBodyDataBoundAgents) SetAgentName(v string) *GetCredentialResponseBodyDataBoundAgents {
	s.AgentName = &v
	return s
}

func (s *GetCredentialResponseBodyDataBoundAgents) Validate() error {
	return dara.Validate(s)
}

type GetCredentialResponseBodyDataResourceRefs struct {
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

func (s GetCredentialResponseBodyDataResourceRefs) String() string {
	return dara.Prettify(s)
}

func (s GetCredentialResponseBodyDataResourceRefs) GoString() string {
	return s.String()
}

func (s *GetCredentialResponseBodyDataResourceRefs) GetResourceId() *string {
	return s.ResourceId
}

func (s *GetCredentialResponseBodyDataResourceRefs) GetResourceName() *string {
	return s.ResourceName
}

func (s *GetCredentialResponseBodyDataResourceRefs) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetCredentialResponseBodyDataResourceRefs) SetResourceId(v string) *GetCredentialResponseBodyDataResourceRefs {
	s.ResourceId = &v
	return s
}

func (s *GetCredentialResponseBodyDataResourceRefs) SetResourceName(v string) *GetCredentialResponseBodyDataResourceRefs {
	s.ResourceName = &v
	return s
}

func (s *GetCredentialResponseBodyDataResourceRefs) SetResourceType(v string) *GetCredentialResponseBodyDataResourceRefs {
	s.ResourceType = &v
	return s
}

func (s *GetCredentialResponseBodyDataResourceRefs) Validate() error {
	return dara.Validate(s)
}
