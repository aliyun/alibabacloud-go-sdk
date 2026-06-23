// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckServiceRoleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRoles(v []*CheckServiceRoleResponseBodyRoles) *CheckServiceRoleResponseBody
	GetRoles() []*CheckServiceRoleResponseBodyRoles
}

type CheckServiceRoleResponseBody struct {
	// The service role check results.
	Roles []*CheckServiceRoleResponseBodyRoles `json:"roles,omitempty" xml:"roles,omitempty" type:"Repeated"`
}

func (s CheckServiceRoleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckServiceRoleResponseBody) GoString() string {
	return s.String()
}

func (s *CheckServiceRoleResponseBody) GetRoles() []*CheckServiceRoleResponseBodyRoles {
	return s.Roles
}

func (s *CheckServiceRoleResponseBody) SetRoles(v []*CheckServiceRoleResponseBodyRoles) *CheckServiceRoleResponseBody {
	s.Roles = v
	return s
}

func (s *CheckServiceRoleResponseBody) Validate() error {
	if s.Roles != nil {
		for _, item := range s.Roles {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CheckServiceRoleResponseBodyRoles struct {
	// Indicates whether the service role has been granted.
	//
	// example:
	//
	// true
	Granted *bool `json:"granted,omitempty" xml:"granted,omitempty"`
	// The prompt message returned when the service role is not granted.
	//
	// example:
	//
	// The role does not exist: AliyunCSManagedAutoScalerRole
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The service role name.
	//
	// example:
	//
	// AliyunCSManagedAutoScalerRole
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s CheckServiceRoleResponseBodyRoles) String() string {
	return dara.Prettify(s)
}

func (s CheckServiceRoleResponseBodyRoles) GoString() string {
	return s.String()
}

func (s *CheckServiceRoleResponseBodyRoles) GetGranted() *bool {
	return s.Granted
}

func (s *CheckServiceRoleResponseBodyRoles) GetMessage() *string {
	return s.Message
}

func (s *CheckServiceRoleResponseBodyRoles) GetName() *string {
	return s.Name
}

func (s *CheckServiceRoleResponseBodyRoles) SetGranted(v bool) *CheckServiceRoleResponseBodyRoles {
	s.Granted = &v
	return s
}

func (s *CheckServiceRoleResponseBodyRoles) SetMessage(v string) *CheckServiceRoleResponseBodyRoles {
	s.Message = &v
	return s
}

func (s *CheckServiceRoleResponseBodyRoles) SetName(v string) *CheckServiceRoleResponseBodyRoles {
	s.Name = &v
	return s
}

func (s *CheckServiceRoleResponseBodyRoles) Validate() error {
	return dara.Validate(s)
}
