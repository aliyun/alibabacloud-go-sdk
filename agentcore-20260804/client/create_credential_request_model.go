// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCredentialRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *CreateCredentialRequestBody) *CreateCredentialRequest
	GetBody() *CreateCredentialRequestBody
	SetClientToken(v string) *CreateCredentialRequest
	GetClientToken() *string
}

type CreateCredentialRequest struct {
	// The request body for creating a credential.
	Body *CreateCredentialRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
	// Not supported.
	//
	// example:
	//
	// Not supported
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s CreateCredentialRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCredentialRequest) GoString() string {
	return s.String()
}

func (s *CreateCredentialRequest) GetBody() *CreateCredentialRequestBody {
	return s.Body
}

func (s *CreateCredentialRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateCredentialRequest) SetBody(v *CreateCredentialRequestBody) *CreateCredentialRequest {
	s.Body = v
	return s
}

func (s *CreateCredentialRequest) SetClientToken(v string) *CreateCredentialRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateCredentialRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateCredentialRequestBody struct {
	// The credential content, submitted as a JSON string. For the apiKey type, it contains only a non-empty apiKey. For the mcpHeader type, it contains a non-empty headers array, where each item contains a name and a value.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"apiKey":"sk-example-value"}
	CredentialMetadata *string `json:"credentialMetadata,omitempty" xml:"credentialMetadata,omitempty"`
	// The credential type. Valid values: apiKey (API key) and mcpHeader (MCP request header).
	//
	// This parameter is required.
	//
	// example:
	//
	// apiKey
	CredentialType *string `json:"credentialType,omitempty" xml:"credentialType,omitempty"`
	// The credential description. It can be up to 256 characters in length.
	//
	// example:
	//
	// API key used to call the model service in the production environment
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The credential name. It must be unique within the workspace, contain only letters, digits, periods, underscores, and hyphens, and be 3 to 99 characters in length. Runtime reserved names cannot be used.
	//
	// This parameter is required.
	//
	// example:
	//
	// model-api-key
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The list of applicable resources for the MCP header credential. This field must be non-empty when resourceScope is set to SPECIFIED, and must be omitted when resourceScope is set to ALL.
	ResourceRefs []*CreateCredentialRequestBodyResourceRefs `json:"resourceRefs,omitempty" xml:"resourceRefs,omitempty" type:"Repeated"`
	// The applicable scope of the MCP header credential. This field is supported only for the mcpHeader type. ALL indicates all MCP services, and SPECIFIED indicates the MCP services specified in resourceRefs. Do not submit this field for the apiKey type.
	//
	// example:
	//
	// ALL
	ResourceScope *string `json:"resourceScope,omitempty" xml:"resourceScope,omitempty"`
}

func (s CreateCredentialRequestBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCredentialRequestBody) GoString() string {
	return s.String()
}

func (s *CreateCredentialRequestBody) GetCredentialMetadata() *string {
	return s.CredentialMetadata
}

func (s *CreateCredentialRequestBody) GetCredentialType() *string {
	return s.CredentialType
}

func (s *CreateCredentialRequestBody) GetDescription() *string {
	return s.Description
}

func (s *CreateCredentialRequestBody) GetName() *string {
	return s.Name
}

func (s *CreateCredentialRequestBody) GetResourceRefs() []*CreateCredentialRequestBodyResourceRefs {
	return s.ResourceRefs
}

func (s *CreateCredentialRequestBody) GetResourceScope() *string {
	return s.ResourceScope
}

func (s *CreateCredentialRequestBody) SetCredentialMetadata(v string) *CreateCredentialRequestBody {
	s.CredentialMetadata = &v
	return s
}

func (s *CreateCredentialRequestBody) SetCredentialType(v string) *CreateCredentialRequestBody {
	s.CredentialType = &v
	return s
}

func (s *CreateCredentialRequestBody) SetDescription(v string) *CreateCredentialRequestBody {
	s.Description = &v
	return s
}

func (s *CreateCredentialRequestBody) SetName(v string) *CreateCredentialRequestBody {
	s.Name = &v
	return s
}

func (s *CreateCredentialRequestBody) SetResourceRefs(v []*CreateCredentialRequestBodyResourceRefs) *CreateCredentialRequestBody {
	s.ResourceRefs = v
	return s
}

func (s *CreateCredentialRequestBody) SetResourceScope(v string) *CreateCredentialRequestBody {
	s.ResourceScope = &v
	return s
}

func (s *CreateCredentialRequestBody) Validate() error {
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

type CreateCredentialRequestBodyResourceRefs struct {
	// The ID of the applicable MCP service.
	//
	// example:
	//
	// mcp-example
	ResourceId *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	// The resource name. This field is empty if the resource is deleted.
	//
	// example:
	//
	// my-mcp-service
	ResourceName *string `json:"resourceName,omitempty" xml:"resourceName,omitempty"`
	// The resource type. Currently, only mcpServer is supported, which indicates an MCP service.
	//
	// example:
	//
	// mcpServer
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
}

func (s CreateCredentialRequestBodyResourceRefs) String() string {
	return dara.Prettify(s)
}

func (s CreateCredentialRequestBodyResourceRefs) GoString() string {
	return s.String()
}

func (s *CreateCredentialRequestBodyResourceRefs) GetResourceId() *string {
	return s.ResourceId
}

func (s *CreateCredentialRequestBodyResourceRefs) GetResourceName() *string {
	return s.ResourceName
}

func (s *CreateCredentialRequestBodyResourceRefs) GetResourceType() *string {
	return s.ResourceType
}

func (s *CreateCredentialRequestBodyResourceRefs) SetResourceId(v string) *CreateCredentialRequestBodyResourceRefs {
	s.ResourceId = &v
	return s
}

func (s *CreateCredentialRequestBodyResourceRefs) SetResourceName(v string) *CreateCredentialRequestBodyResourceRefs {
	s.ResourceName = &v
	return s
}

func (s *CreateCredentialRequestBodyResourceRefs) SetResourceType(v string) *CreateCredentialRequestBodyResourceRefs {
	s.ResourceType = &v
	return s
}

func (s *CreateCredentialRequestBodyResourceRefs) Validate() error {
	return dara.Validate(s)
}
