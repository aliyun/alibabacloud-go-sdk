// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWorkloadIdentityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllowedConsentCallbackURLs(v []*string) *CreateWorkloadIdentityRequest
	GetAllowedConsentCallbackURLs() []*string
	SetAllowedResourceOAuth2ReturnURLs(v []*string) *CreateWorkloadIdentityRequest
	GetAllowedResourceOAuth2ReturnURLs() []*string
	SetCreateRAMRole(v bool) *CreateWorkloadIdentityRequest
	GetCreateRAMRole() *bool
	SetDescription(v string) *CreateWorkloadIdentityRequest
	GetDescription() *string
	SetIdentityProviderName(v string) *CreateWorkloadIdentityRequest
	GetIdentityProviderName() *string
	SetRoleArn(v string) *CreateWorkloadIdentityRequest
	GetRoleArn() *string
	SetSessionBindingEnabled(v bool) *CreateWorkloadIdentityRequest
	GetSessionBindingEnabled() *bool
	SetSourceAgentArn(v string) *CreateWorkloadIdentityRequest
	GetSourceAgentArn() *string
	SetSourcePlatform(v string) *CreateWorkloadIdentityRequest
	GetSourcePlatform() *string
	SetWorkloadIdentityName(v string) *CreateWorkloadIdentityRequest
	GetWorkloadIdentityName() *string
}

type CreateWorkloadIdentityRequest struct {
	// The allowlist of Consent authorization callback URLs for the workload identity, which constrains the allowed redirect addresses during the offline delegation (OfflineAccess) authorization flow.
	AllowedConsentCallbackURLs []*string `json:"AllowedConsentCallbackURLs,omitempty" xml:"AllowedConsentCallbackURLs,omitempty" type:"Repeated"`
	// The list of allowed OAuth2 return URLs.
	AllowedResourceOAuth2ReturnURLs []*string `json:"AllowedResourceOAuth2ReturnURLs,omitempty" xml:"AllowedResourceOAuth2ReturnURLs,omitempty" type:"Repeated"`
	// Specifies whether to automatically create an associated RAM role. When set to true, the system automatically creates a RAM role with the same name as the workload identity.
	//
	// example:
	//
	// true
	CreateRAMRole *bool `json:"CreateRAMRole,omitempty" xml:"CreateRAMRole,omitempty"`
	// The description.
	//
	// example:
	//
	// example agent
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The configured identity provider.
	//
	// example:
	//
	// identity-provider-okta
	IdentityProviderName *string `json:"IdentityProviderName,omitempty" xml:"IdentityProviderName,omitempty"`
	// The ARN of the RAM role.
	//
	// example:
	//
	// acs:ram::123456:role/agent-101-role
	RoleArn *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
	// Specifies whether to enable session binding. When enabled, each session is bound to a specific workload identity instance.
	//
	// example:
	//
	// false
	SessionBindingEnabled *bool `json:"SessionBindingEnabled,omitempty" xml:"SessionBindingEnabled,omitempty"`
	// The ARN of the source agent.
	//
	// example:
	//
	// acs:agentidentity:cn-beijing:123456:workloadidentity/my-identity
	SourceAgentArn *string `json:"SourceAgentArn,omitempty" xml:"SourceAgentArn,omitempty"`
	// The source platform.
	//
	// example:
	//
	// China
	SourcePlatform *string `json:"SourcePlatform,omitempty" xml:"SourcePlatform,omitempty"`
	// The unique name of the workload identity.
	//
	// example:
	//
	// agent-101
	WorkloadIdentityName *string `json:"WorkloadIdentityName,omitempty" xml:"WorkloadIdentityName,omitempty"`
}

func (s CreateWorkloadIdentityRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkloadIdentityRequest) GoString() string {
	return s.String()
}

func (s *CreateWorkloadIdentityRequest) GetAllowedConsentCallbackURLs() []*string {
	return s.AllowedConsentCallbackURLs
}

func (s *CreateWorkloadIdentityRequest) GetAllowedResourceOAuth2ReturnURLs() []*string {
	return s.AllowedResourceOAuth2ReturnURLs
}

func (s *CreateWorkloadIdentityRequest) GetCreateRAMRole() *bool {
	return s.CreateRAMRole
}

func (s *CreateWorkloadIdentityRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateWorkloadIdentityRequest) GetIdentityProviderName() *string {
	return s.IdentityProviderName
}

func (s *CreateWorkloadIdentityRequest) GetRoleArn() *string {
	return s.RoleArn
}

func (s *CreateWorkloadIdentityRequest) GetSessionBindingEnabled() *bool {
	return s.SessionBindingEnabled
}

func (s *CreateWorkloadIdentityRequest) GetSourceAgentArn() *string {
	return s.SourceAgentArn
}

func (s *CreateWorkloadIdentityRequest) GetSourcePlatform() *string {
	return s.SourcePlatform
}

func (s *CreateWorkloadIdentityRequest) GetWorkloadIdentityName() *string {
	return s.WorkloadIdentityName
}

func (s *CreateWorkloadIdentityRequest) SetAllowedConsentCallbackURLs(v []*string) *CreateWorkloadIdentityRequest {
	s.AllowedConsentCallbackURLs = v
	return s
}

func (s *CreateWorkloadIdentityRequest) SetAllowedResourceOAuth2ReturnURLs(v []*string) *CreateWorkloadIdentityRequest {
	s.AllowedResourceOAuth2ReturnURLs = v
	return s
}

func (s *CreateWorkloadIdentityRequest) SetCreateRAMRole(v bool) *CreateWorkloadIdentityRequest {
	s.CreateRAMRole = &v
	return s
}

func (s *CreateWorkloadIdentityRequest) SetDescription(v string) *CreateWorkloadIdentityRequest {
	s.Description = &v
	return s
}

func (s *CreateWorkloadIdentityRequest) SetIdentityProviderName(v string) *CreateWorkloadIdentityRequest {
	s.IdentityProviderName = &v
	return s
}

func (s *CreateWorkloadIdentityRequest) SetRoleArn(v string) *CreateWorkloadIdentityRequest {
	s.RoleArn = &v
	return s
}

func (s *CreateWorkloadIdentityRequest) SetSessionBindingEnabled(v bool) *CreateWorkloadIdentityRequest {
	s.SessionBindingEnabled = &v
	return s
}

func (s *CreateWorkloadIdentityRequest) SetSourceAgentArn(v string) *CreateWorkloadIdentityRequest {
	s.SourceAgentArn = &v
	return s
}

func (s *CreateWorkloadIdentityRequest) SetSourcePlatform(v string) *CreateWorkloadIdentityRequest {
	s.SourcePlatform = &v
	return s
}

func (s *CreateWorkloadIdentityRequest) SetWorkloadIdentityName(v string) *CreateWorkloadIdentityRequest {
	s.WorkloadIdentityName = &v
	return s
}

func (s *CreateWorkloadIdentityRequest) Validate() error {
	return dara.Validate(s)
}
