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
	SetServiceLinkedRole(v string) *CreateServiceLinkedRoleRequest
	GetServiceLinkedRole() *string
}

type CreateServiceLinkedRoleRequest struct {
	// The client token that is used to ensure the idempotence of the request. Different requests should use different tokens. The token supports only ASCII characters and cannot exceed 64 characters in length.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The service-linked role. Default value: **AliyunServiceRoleForSas**. Valid values:
	//
	// - **AliyunServiceRoleForSas**: the service-linked role for Security Center (SAS). Security Center uses this role to access your resources in other cloud services.
	//
	// - **AliyunServiceRoleForSasCspm**: the service-linked role for Security Center - Cloud Security Posture Management (CSPM). SAS-CSPM uses this role to access your resources in other cloud services.
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

func (s *CreateServiceLinkedRoleRequest) GetServiceLinkedRole() *string {
	return s.ServiceLinkedRole
}

func (s *CreateServiceLinkedRoleRequest) SetClientToken(v string) *CreateServiceLinkedRoleRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateServiceLinkedRoleRequest) SetServiceLinkedRole(v string) *CreateServiceLinkedRoleRequest {
	s.ServiceLinkedRole = &v
	return s
}

func (s *CreateServiceLinkedRoleRequest) Validate() error {
	return dara.Validate(s)
}
