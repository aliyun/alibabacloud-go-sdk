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
	// The new credential content. The value is a JSON string. If credentialType is set to apiKey, only the apiKey field can be included, and the value cannot be empty. At least one of credentialMetadata and description must be specified.
	//
	// example:
	//
	// {"apiKey":"sk-example-value"}
	CredentialMetadata *string `json:"credentialMetadata,omitempty" xml:"credentialMetadata,omitempty"`
	// The new credential description. The description can be up to 256 characters in length. At least one of description and credentialMetadata must be specified.
	//
	// example:
	//
	// API Key used for calling model services in the production environment
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// This parameter is required and must be a non-empty array when resourceScope is set to SPECIFIED. Each item contains resourceType and resourceId. resourceName is optional.
	ResourceRefs []*UpdateCredentialRequestBodyResourceRefs `json:"resourceRefs,omitempty" xml:"resourceRefs,omitempty" type:"Repeated"`
	// ALL indicates all resources. SPECIFIED indicates that the credential applies only to the resources specified in resourceRefs.
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
