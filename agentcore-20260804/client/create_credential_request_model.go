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
	// The credential content. The value is a JSON string. When credentialType is set to apiKey, the content can contain only the apiKey field, and the value cannot be empty. After being written, the content can only be queried in masked form.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"apiKey":"sk-example-value"}
	CredentialMetadata *string `json:"credentialMetadata,omitempty" xml:"credentialMetadata,omitempty"`
	// The credential type. Currently, only apiKey is supported.
	//
	// This parameter is required.
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
	// This parameter is required.
	//
	// example:
	//
	// model-api-key
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required and must be a non-empty array when resourceScope is set to SPECIFIED. Each item contains resourceType and resourceId. resourceName is optional.
	ResourceRefs []*CreateCredentialRequestBodyResourceRefs `json:"resourceRefs,omitempty" xml:"resourceRefs,omitempty" type:"Repeated"`
	// ALL indicates all resources. SPECIFIED indicates that the credential applies only to the resources specified in resourceRefs.
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
