// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWorkloadIdentityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllowedConsentCallbackURLs(v []*string) *UpdateWorkloadIdentityRequest
	GetAllowedConsentCallbackURLs() []*string
	SetAllowedResourceOAuth2ReturnURLs(v []*string) *UpdateWorkloadIdentityRequest
	GetAllowedResourceOAuth2ReturnURLs() []*string
	SetDescription(v string) *UpdateWorkloadIdentityRequest
	GetDescription() *string
	SetIdentityProviderName(v string) *UpdateWorkloadIdentityRequest
	GetIdentityProviderName() *string
	SetRoleArn(v string) *UpdateWorkloadIdentityRequest
	GetRoleArn() *string
	SetSessionBindingEnabled(v string) *UpdateWorkloadIdentityRequest
	GetSessionBindingEnabled() *string
	SetWorkloadIdentityName(v string) *UpdateWorkloadIdentityRequest
	GetWorkloadIdentityName() *string
}

type UpdateWorkloadIdentityRequest struct {
	// The whitelist of consent authorization callback URLs for the workload identity, which restricts the allowed callback URLs in the offline access delegation authorization flow.
	AllowedConsentCallbackURLs []*string `json:"AllowedConsentCallbackURLs,omitempty" xml:"AllowedConsentCallbackURLs,omitempty" type:"Repeated"`
	// The list of allowed OAuth2 return URLs.
	//
	// if can be null:
	// false
	AllowedResourceOAuth2ReturnURLs []*string `json:"AllowedResourceOAuth2ReturnURLs,omitempty" xml:"AllowedResourceOAuth2ReturnURLs,omitempty" type:"Repeated"`
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
	// Specifies whether session binding is enabled. If enabled, each session is bound to a specific workload identity instance.
	//
	// example:
	//
	// false
	SessionBindingEnabled *string `json:"SessionBindingEnabled,omitempty" xml:"SessionBindingEnabled,omitempty"`
	// The unique name of the workload identity.
	//
	// example:
	//
	// agent-101
	WorkloadIdentityName *string `json:"WorkloadIdentityName,omitempty" xml:"WorkloadIdentityName,omitempty"`
}

func (s UpdateWorkloadIdentityRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkloadIdentityRequest) GoString() string {
	return s.String()
}

func (s *UpdateWorkloadIdentityRequest) GetAllowedConsentCallbackURLs() []*string {
	return s.AllowedConsentCallbackURLs
}

func (s *UpdateWorkloadIdentityRequest) GetAllowedResourceOAuth2ReturnURLs() []*string {
	return s.AllowedResourceOAuth2ReturnURLs
}

func (s *UpdateWorkloadIdentityRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateWorkloadIdentityRequest) GetIdentityProviderName() *string {
	return s.IdentityProviderName
}

func (s *UpdateWorkloadIdentityRequest) GetRoleArn() *string {
	return s.RoleArn
}

func (s *UpdateWorkloadIdentityRequest) GetSessionBindingEnabled() *string {
	return s.SessionBindingEnabled
}

func (s *UpdateWorkloadIdentityRequest) GetWorkloadIdentityName() *string {
	return s.WorkloadIdentityName
}

func (s *UpdateWorkloadIdentityRequest) SetAllowedConsentCallbackURLs(v []*string) *UpdateWorkloadIdentityRequest {
	s.AllowedConsentCallbackURLs = v
	return s
}

func (s *UpdateWorkloadIdentityRequest) SetAllowedResourceOAuth2ReturnURLs(v []*string) *UpdateWorkloadIdentityRequest {
	s.AllowedResourceOAuth2ReturnURLs = v
	return s
}

func (s *UpdateWorkloadIdentityRequest) SetDescription(v string) *UpdateWorkloadIdentityRequest {
	s.Description = &v
	return s
}

func (s *UpdateWorkloadIdentityRequest) SetIdentityProviderName(v string) *UpdateWorkloadIdentityRequest {
	s.IdentityProviderName = &v
	return s
}

func (s *UpdateWorkloadIdentityRequest) SetRoleArn(v string) *UpdateWorkloadIdentityRequest {
	s.RoleArn = &v
	return s
}

func (s *UpdateWorkloadIdentityRequest) SetSessionBindingEnabled(v string) *UpdateWorkloadIdentityRequest {
	s.SessionBindingEnabled = &v
	return s
}

func (s *UpdateWorkloadIdentityRequest) SetWorkloadIdentityName(v string) *UpdateWorkloadIdentityRequest {
	s.WorkloadIdentityName = &v
	return s
}

func (s *UpdateWorkloadIdentityRequest) Validate() error {
	return dara.Validate(s)
}
