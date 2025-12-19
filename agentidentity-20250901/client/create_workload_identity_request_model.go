// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWorkloadIdentityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllowedResourceOAuth2ReturnURLs(v []*string) *CreateWorkloadIdentityRequest
	GetAllowedResourceOAuth2ReturnURLs() []*string
	SetDescription(v string) *CreateWorkloadIdentityRequest
	GetDescription() *string
	SetIdentityProviderName(v string) *CreateWorkloadIdentityRequest
	GetIdentityProviderName() *string
	SetRoleArn(v string) *CreateWorkloadIdentityRequest
	GetRoleArn() *string
	SetWorkloadIdentityName(v string) *CreateWorkloadIdentityRequest
	GetWorkloadIdentityName() *string
}

type CreateWorkloadIdentityRequest struct {
	AllowedResourceOAuth2ReturnURLs []*string `json:"AllowedResourceOAuth2ReturnURLs,omitempty" xml:"AllowedResourceOAuth2ReturnURLs,omitempty" type:"Repeated"`
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

func (s CreateWorkloadIdentityRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkloadIdentityRequest) GoString() string {
	return s.String()
}

func (s *CreateWorkloadIdentityRequest) GetAllowedResourceOAuth2ReturnURLs() []*string {
	return s.AllowedResourceOAuth2ReturnURLs
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

func (s *CreateWorkloadIdentityRequest) GetWorkloadIdentityName() *string {
	return s.WorkloadIdentityName
}

func (s *CreateWorkloadIdentityRequest) SetAllowedResourceOAuth2ReturnURLs(v []*string) *CreateWorkloadIdentityRequest {
	s.AllowedResourceOAuth2ReturnURLs = v
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

func (s *CreateWorkloadIdentityRequest) SetWorkloadIdentityName(v string) *CreateWorkloadIdentityRequest {
	s.WorkloadIdentityName = &v
	return s
}

func (s *CreateWorkloadIdentityRequest) Validate() error {
	return dara.Validate(s)
}
