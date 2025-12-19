// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWorkloadIdentityShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllowedResourceOAuth2ReturnURLsShrink(v string) *CreateWorkloadIdentityShrinkRequest
	GetAllowedResourceOAuth2ReturnURLsShrink() *string
	SetDescription(v string) *CreateWorkloadIdentityShrinkRequest
	GetDescription() *string
	SetIdentityProviderName(v string) *CreateWorkloadIdentityShrinkRequest
	GetIdentityProviderName() *string
	SetRoleArn(v string) *CreateWorkloadIdentityShrinkRequest
	GetRoleArn() *string
	SetWorkloadIdentityName(v string) *CreateWorkloadIdentityShrinkRequest
	GetWorkloadIdentityName() *string
}

type CreateWorkloadIdentityShrinkRequest struct {
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

func (s CreateWorkloadIdentityShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkloadIdentityShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateWorkloadIdentityShrinkRequest) GetAllowedResourceOAuth2ReturnURLsShrink() *string {
	return s.AllowedResourceOAuth2ReturnURLsShrink
}

func (s *CreateWorkloadIdentityShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateWorkloadIdentityShrinkRequest) GetIdentityProviderName() *string {
	return s.IdentityProviderName
}

func (s *CreateWorkloadIdentityShrinkRequest) GetRoleArn() *string {
	return s.RoleArn
}

func (s *CreateWorkloadIdentityShrinkRequest) GetWorkloadIdentityName() *string {
	return s.WorkloadIdentityName
}

func (s *CreateWorkloadIdentityShrinkRequest) SetAllowedResourceOAuth2ReturnURLsShrink(v string) *CreateWorkloadIdentityShrinkRequest {
	s.AllowedResourceOAuth2ReturnURLsShrink = &v
	return s
}

func (s *CreateWorkloadIdentityShrinkRequest) SetDescription(v string) *CreateWorkloadIdentityShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateWorkloadIdentityShrinkRequest) SetIdentityProviderName(v string) *CreateWorkloadIdentityShrinkRequest {
	s.IdentityProviderName = &v
	return s
}

func (s *CreateWorkloadIdentityShrinkRequest) SetRoleArn(v string) *CreateWorkloadIdentityShrinkRequest {
	s.RoleArn = &v
	return s
}

func (s *CreateWorkloadIdentityShrinkRequest) SetWorkloadIdentityName(v string) *CreateWorkloadIdentityShrinkRequest {
	s.WorkloadIdentityName = &v
	return s
}

func (s *CreateWorkloadIdentityShrinkRequest) Validate() error {
	return dara.Validate(s)
}
