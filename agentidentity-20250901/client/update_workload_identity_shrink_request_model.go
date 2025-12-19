// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWorkloadIdentityShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllowedResourceOAuth2ReturnURLsShrink(v string) *UpdateWorkloadIdentityShrinkRequest
	GetAllowedResourceOAuth2ReturnURLsShrink() *string
	SetDescription(v string) *UpdateWorkloadIdentityShrinkRequest
	GetDescription() *string
	SetIdentityProviderName(v string) *UpdateWorkloadIdentityShrinkRequest
	GetIdentityProviderName() *string
	SetRoleArn(v string) *UpdateWorkloadIdentityShrinkRequest
	GetRoleArn() *string
	SetWorkloadIdentityName(v string) *UpdateWorkloadIdentityShrinkRequest
	GetWorkloadIdentityName() *string
}

type UpdateWorkloadIdentityShrinkRequest struct {
	// if can be null:
	// false
	AllowedResourceOAuth2ReturnURLsShrink *string `json:"AllowedResourceOAuth2ReturnURLs,omitempty" xml:"AllowedResourceOAuth2ReturnURLs,omitempty"`
	// example:
	//
	// example agent
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// identity-provider-okta
	IdentityProviderName *string `json:"IdentityProviderName,omitempty" xml:"IdentityProviderName,omitempty"`
	// example:
	//
	// acs:ram::123456:role/agent-101-role
	RoleArn *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
	// example:
	//
	// agent-101
	WorkloadIdentityName *string `json:"WorkloadIdentityName,omitempty" xml:"WorkloadIdentityName,omitempty"`
}

func (s UpdateWorkloadIdentityShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkloadIdentityShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateWorkloadIdentityShrinkRequest) GetAllowedResourceOAuth2ReturnURLsShrink() *string {
	return s.AllowedResourceOAuth2ReturnURLsShrink
}

func (s *UpdateWorkloadIdentityShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateWorkloadIdentityShrinkRequest) GetIdentityProviderName() *string {
	return s.IdentityProviderName
}

func (s *UpdateWorkloadIdentityShrinkRequest) GetRoleArn() *string {
	return s.RoleArn
}

func (s *UpdateWorkloadIdentityShrinkRequest) GetWorkloadIdentityName() *string {
	return s.WorkloadIdentityName
}

func (s *UpdateWorkloadIdentityShrinkRequest) SetAllowedResourceOAuth2ReturnURLsShrink(v string) *UpdateWorkloadIdentityShrinkRequest {
	s.AllowedResourceOAuth2ReturnURLsShrink = &v
	return s
}

func (s *UpdateWorkloadIdentityShrinkRequest) SetDescription(v string) *UpdateWorkloadIdentityShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateWorkloadIdentityShrinkRequest) SetIdentityProviderName(v string) *UpdateWorkloadIdentityShrinkRequest {
	s.IdentityProviderName = &v
	return s
}

func (s *UpdateWorkloadIdentityShrinkRequest) SetRoleArn(v string) *UpdateWorkloadIdentityShrinkRequest {
	s.RoleArn = &v
	return s
}

func (s *UpdateWorkloadIdentityShrinkRequest) SetWorkloadIdentityName(v string) *UpdateWorkloadIdentityShrinkRequest {
	s.WorkloadIdentityName = &v
	return s
}

func (s *UpdateWorkloadIdentityShrinkRequest) Validate() error {
	return dara.Validate(s)
}
