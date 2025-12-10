// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServiceLinkedRoleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomSuffix(v string) *CreateServiceLinkedRoleRequest
	GetCustomSuffix() *string
	SetDescription(v string) *CreateServiceLinkedRoleRequest
	GetDescription() *string
	SetServiceName(v string) *CreateServiceLinkedRoleRequest
	GetServiceName() *string
}

type CreateServiceLinkedRoleRequest struct {
	// The suffix of the role name.
	//
	// The role name (including its suffix) must be 1 to 64 characters in length and can contain letters, digits, periods (.), and hyphens (-).
	//
	// For example, if the suffix is `Example`, the role name is `ServiceLinkedRoleName_Example`.
	//
	// example:
	//
	// Example
	CustomSuffix *string `json:"CustomSuffix,omitempty" xml:"CustomSuffix,omitempty"`
	// The description of the service-linked role.
	//
	// You must configure this parameter for service-linked roles that support custom suffixes. Otherwise, the preset value is used and cannot be modified.
	//
	// The description must be 1 to 1,024 characters in length.
	//
	// example:
	//
	// Service Linked Role for PolarDB. PolarDB will use this role to access your resources in other services.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The service name.
	//
	// For more information about the service name, see [Alibaba Cloud services that support service-linked roles](https://help.aliyun.com/document_detail/461722.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// polardb.aliyuncs.com
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s CreateServiceLinkedRoleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceLinkedRoleRequest) GoString() string {
	return s.String()
}

func (s *CreateServiceLinkedRoleRequest) GetCustomSuffix() *string {
	return s.CustomSuffix
}

func (s *CreateServiceLinkedRoleRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateServiceLinkedRoleRequest) GetServiceName() *string {
	return s.ServiceName
}

func (s *CreateServiceLinkedRoleRequest) SetCustomSuffix(v string) *CreateServiceLinkedRoleRequest {
	s.CustomSuffix = &v
	return s
}

func (s *CreateServiceLinkedRoleRequest) SetDescription(v string) *CreateServiceLinkedRoleRequest {
	s.Description = &v
	return s
}

func (s *CreateServiceLinkedRoleRequest) SetServiceName(v string) *CreateServiceLinkedRoleRequest {
	s.ServiceName = &v
	return s
}

func (s *CreateServiceLinkedRoleRequest) Validate() error {
	return dara.Validate(s)
}
