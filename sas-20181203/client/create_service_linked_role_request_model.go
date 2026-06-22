// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServiceLinkedRoleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetServiceLinkedRole(v string) *CreateServiceLinkedRoleRequest
	GetServiceLinkedRole() *string
}

type CreateServiceLinkedRoleRequest struct {
	// The service-linked role. Default value: **AliyunServiceRoleForSas**. Valid values:
	//
	// - **AliyunServiceRoleForSas**: the service-linked role for Security Center (SAS). Security Center uses this role to access your resources in other Alibaba Cloud services.
	//
	// - **AliyunServiceRoleForSasCspm**: the service-linked role for Security Center - Cloud Security Posture Management (CSPM) (sas-cspm). sas-cspm uses this role to access your resources in other Alibaba Cloud services.
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

func (s *CreateServiceLinkedRoleRequest) GetServiceLinkedRole() *string {
	return s.ServiceLinkedRole
}

func (s *CreateServiceLinkedRoleRequest) SetServiceLinkedRole(v string) *CreateServiceLinkedRoleRequest {
	s.ServiceLinkedRole = &v
	return s
}

func (s *CreateServiceLinkedRoleRequest) Validate() error {
	return dara.Validate(s)
}
