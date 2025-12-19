// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWorkloadIdentityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllowedResourceOAuth2ReturnURLs(v []*string) *UpdateWorkloadIdentityRequest
	GetAllowedResourceOAuth2ReturnURLs() []*string
	SetDescription(v string) *UpdateWorkloadIdentityRequest
	GetDescription() *string
	SetIdentityProviderName(v string) *UpdateWorkloadIdentityRequest
	GetIdentityProviderName() *string
	SetRoleArn(v string) *UpdateWorkloadIdentityRequest
	GetRoleArn() *string
	SetWorkloadIdentityName(v string) *UpdateWorkloadIdentityRequest
	GetWorkloadIdentityName() *string
}

type UpdateWorkloadIdentityRequest struct {
	// if can be null:
	// false
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

func (s UpdateWorkloadIdentityRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkloadIdentityRequest) GoString() string {
	return s.String()
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

func (s *UpdateWorkloadIdentityRequest) GetWorkloadIdentityName() *string {
	return s.WorkloadIdentityName
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

func (s *UpdateWorkloadIdentityRequest) SetWorkloadIdentityName(v string) *UpdateWorkloadIdentityRequest {
	s.WorkloadIdentityName = &v
	return s
}

func (s *UpdateWorkloadIdentityRequest) Validate() error {
	return dara.Validate(s)
}
