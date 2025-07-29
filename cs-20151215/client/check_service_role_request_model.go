// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckServiceRoleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRoles(v []*CheckServiceRoleRequestRoles) *CheckServiceRoleRequest
	GetRoles() []*CheckServiceRoleRequestRoles
}

type CheckServiceRoleRequest struct {
	// The list of service roles you want to check.
	//
	// This parameter is required.
	Roles []*CheckServiceRoleRequestRoles `json:"roles,omitempty" xml:"roles,omitempty" type:"Repeated"`
}

func (s CheckServiceRoleRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckServiceRoleRequest) GoString() string {
	return s.String()
}

func (s *CheckServiceRoleRequest) GetRoles() []*CheckServiceRoleRequestRoles {
	return s.Roles
}

func (s *CheckServiceRoleRequest) SetRoles(v []*CheckServiceRoleRequestRoles) *CheckServiceRoleRequest {
	s.Roles = v
	return s
}

func (s *CheckServiceRoleRequest) Validate() error {
	return dara.Validate(s)
}

type CheckServiceRoleRequestRoles struct {
	// The server role name. For more information about the service roles and their permissions in ACK, see [ACK roles](https://help.aliyun.com/document_detail/86483.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// AliyunCSManagedAutoScalerRole
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s CheckServiceRoleRequestRoles) String() string {
	return dara.Prettify(s)
}

func (s CheckServiceRoleRequestRoles) GoString() string {
	return s.String()
}

func (s *CheckServiceRoleRequestRoles) GetName() *string {
	return s.Name
}

func (s *CheckServiceRoleRequestRoles) SetName(v string) *CheckServiceRoleRequestRoles {
	s.Name = &v
	return s
}

func (s *CheckServiceRoleRequestRoles) Validate() error {
	return dara.Validate(s)
}
