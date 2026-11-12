// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCredentialRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *UpdateCredentialRequestBody) *UpdateCredentialRequest
	GetBody() *UpdateCredentialRequestBody
	SetClientToken(v string) *UpdateCredentialRequest
	GetClientToken() *string
}

type UpdateCredentialRequest struct {
	// The request body for updating the credential.
	Body *UpdateCredentialRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
	// Not supported.
	//
	// example:
	//
	// Not supported
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s UpdateCredentialRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCredentialRequest) GoString() string {
	return s.String()
}

func (s *UpdateCredentialRequest) GetBody() *UpdateCredentialRequestBody {
	return s.Body
}

func (s *UpdateCredentialRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateCredentialRequest) SetBody(v *UpdateCredentialRequestBody) *UpdateCredentialRequest {
	s.Body = v
	return s
}

func (s *UpdateCredentialRequest) SetClientToken(v string) *UpdateCredentialRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateCredentialRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateCredentialRequestBody struct {
	// The credential content, submitted as a JSON string. For the apiKey type, it only contains a non-empty apiKey. For the mcpHeader type, it contains a non-empty headers array, where each item includes a name and a value.
	//
	// example:
	//
	// {"apiKey":"sk-example-value"}
	CredentialMetadata *string `json:"credentialMetadata,omitempty" xml:"credentialMetadata,omitempty"`
	// The new credential description. The description can be up to 256 characters in length. If omitted, the original description is retained. It can be updated separately from the credential content or application scope.
	//
	// example:
	//
	// API key used to call the model service in the production environment
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The list of applicable resources for the MCP header credential. If submitted, resourceScope must also be submitted. An empty array [] is allowed when the value is SPECIFIED. Omit this field when the value is ALL.
	ResourceRefs []*UpdateCredentialRequestBodyResourceRefs `json:"resourceRefs,omitempty" xml:"resourceRefs,omitempty" type:"Repeated"`
	// The application scope of the MCP header credential. Only the mcpHeader type is supported. ALL indicates all MCP services, and SPECIFIED indicates the MCP services specified in resourceRefs. Do not submit this field for the apiKey type.
	//
	// example:
	//
	// ALL
	ResourceScope *string `json:"resourceScope,omitempty" xml:"resourceScope,omitempty"`
}

func (s UpdateCredentialRequestBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateCredentialRequestBody) GoString() string {
	return s.String()
}

func (s *UpdateCredentialRequestBody) GetCredentialMetadata() *string {
	return s.CredentialMetadata
}

func (s *UpdateCredentialRequestBody) GetDescription() *string {
	return s.Description
}

func (s *UpdateCredentialRequestBody) GetResourceRefs() []*UpdateCredentialRequestBodyResourceRefs {
	return s.ResourceRefs
}

func (s *UpdateCredentialRequestBody) GetResourceScope() *string {
	return s.ResourceScope
}

func (s *UpdateCredentialRequestBody) SetCredentialMetadata(v string) *UpdateCredentialRequestBody {
	s.CredentialMetadata = &v
	return s
}

func (s *UpdateCredentialRequestBody) SetDescription(v string) *UpdateCredentialRequestBody {
	s.Description = &v
	return s
}

func (s *UpdateCredentialRequestBody) SetResourceRefs(v []*UpdateCredentialRequestBodyResourceRefs) *UpdateCredentialRequestBody {
	s.ResourceRefs = v
	return s
}

func (s *UpdateCredentialRequestBody) SetResourceScope(v string) *UpdateCredentialRequestBody {
	s.ResourceScope = &v
	return s
}

func (s *UpdateCredentialRequestBody) Validate() error {
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

type UpdateCredentialRequestBodyResourceRefs struct {
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
	// The resource type. The current value is mcpServer, which indicates an MCP service.
	//
	// example:
	//
	// mcpServer
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
}

func (s UpdateCredentialRequestBodyResourceRefs) String() string {
	return dara.Prettify(s)
}

func (s UpdateCredentialRequestBodyResourceRefs) GoString() string {
	return s.String()
}

func (s *UpdateCredentialRequestBodyResourceRefs) GetResourceId() *string {
	return s.ResourceId
}

func (s *UpdateCredentialRequestBodyResourceRefs) GetResourceName() *string {
	return s.ResourceName
}

func (s *UpdateCredentialRequestBodyResourceRefs) GetResourceType() *string {
	return s.ResourceType
}

func (s *UpdateCredentialRequestBodyResourceRefs) SetResourceId(v string) *UpdateCredentialRequestBodyResourceRefs {
	s.ResourceId = &v
	return s
}

func (s *UpdateCredentialRequestBodyResourceRefs) SetResourceName(v string) *UpdateCredentialRequestBodyResourceRefs {
	s.ResourceName = &v
	return s
}

func (s *UpdateCredentialRequestBodyResourceRefs) SetResourceType(v string) *UpdateCredentialRequestBodyResourceRefs {
	s.ResourceType = &v
	return s
}

func (s *UpdateCredentialRequestBodyResourceRefs) Validate() error {
	return dara.Validate(s)
}
