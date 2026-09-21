// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServiceLinkedRoleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateServiceLinkedRoleRequest
	GetClientToken() *string
	SetDryRun(v bool) *CreateServiceLinkedRoleRequest
	GetDryRun() *bool
	SetServiceLinkedRole(v string) *CreateServiceLinkedRoleRequest
	GetServiceLinkedRole() *string
}

type CreateServiceLinkedRoleRequest struct {
	// The client token that is used to ensure the idempotence of the request. Different requests must use different tokens. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run. Valid values:
	//
	// - true: performs only a dry run without executing the actual operation.
	//
	// - false: performs the actual operation.
	//
	// Default value: false.
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The service-linked role. Default value: **AliyunServiceRoleForSas**. Valid values:
	//
	// - **AliyunServiceRoleForSas**: The service-linked role for Security Center (SAS). Security Center uses this role to access your resources in other Alibaba Cloud services.
	//
	// - **AliyunServiceRoleForSasCspm**: The service-linked role for Security Center - Cloud Security Posture Management (CSPM) (sas-cspm). sas-cspm uses this role to access your resources in other Alibaba Cloud services.
	//
	// example:
	//
	// AliyunServiceRoleForSas
	ServiceLinkedRole *string `json:"ServiceLinkedRole,omitempty" xml:"ServiceLinkedRole,omitempty"`
}

func (s CreateServiceLinkedRoleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceLinkedRoleRequest) GoString() string {
	return s.String()
}

func (s *CreateServiceLinkedRoleRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateServiceLinkedRoleRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateServiceLinkedRoleRequest) GetServiceLinkedRole() *string {
	return s.ServiceLinkedRole
}

func (s *CreateServiceLinkedRoleRequest) SetClientToken(v string) *CreateServiceLinkedRoleRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateServiceLinkedRoleRequest) SetDryRun(v bool) *CreateServiceLinkedRoleRequest {
	s.DryRun = &v
	return s
}

func (s *CreateServiceLinkedRoleRequest) SetServiceLinkedRole(v string) *CreateServiceLinkedRoleRequest {
	s.ServiceLinkedRole = &v
	return s
}

func (s *CreateServiceLinkedRoleRequest) Validate() error {
	return dara.Validate(s)
}
