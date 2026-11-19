// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAccountRolesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeAccountRolesResponseBody
	GetRequestId() *string
	SetRoles(v *DescribeAccountRolesResponseBodyRoles) *DescribeAccountRolesResponseBody
	GetRoles() *DescribeAccountRolesResponseBodyRoles
}

type DescribeAccountRolesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 7FEF1BEF-2842-5758-9126-415A297*****
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Roles     *DescribeAccountRolesResponseBodyRoles `json:"Roles,omitempty" xml:"Roles,omitempty" type:"Struct"`
}

func (s DescribeAccountRolesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccountRolesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAccountRolesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAccountRolesResponseBody) GetRoles() *DescribeAccountRolesResponseBodyRoles {
	return s.Roles
}

func (s *DescribeAccountRolesResponseBody) SetRequestId(v string) *DescribeAccountRolesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAccountRolesResponseBody) SetRoles(v *DescribeAccountRolesResponseBodyRoles) *DescribeAccountRolesResponseBody {
	s.Roles = v
	return s
}

func (s *DescribeAccountRolesResponseBody) Validate() error {
	if s.Roles != nil {
		if err := s.Roles.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAccountRolesResponseBodyRoles struct {
	Role []*DescribeAccountRolesResponseBodyRolesRole `json:"Role,omitempty" xml:"Role,omitempty" type:"Repeated"`
}

func (s DescribeAccountRolesResponseBodyRoles) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccountRolesResponseBodyRoles) GoString() string {
	return s.String()
}

func (s *DescribeAccountRolesResponseBodyRoles) GetRole() []*DescribeAccountRolesResponseBodyRolesRole {
	return s.Role
}

func (s *DescribeAccountRolesResponseBodyRoles) SetRole(v []*DescribeAccountRolesResponseBodyRolesRole) *DescribeAccountRolesResponseBodyRoles {
	s.Role = v
	return s
}

func (s *DescribeAccountRolesResponseBodyRoles) Validate() error {
	if s.Role != nil {
		for _, item := range s.Role {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAccountRolesResponseBodyRolesRole struct {
	DBName *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	Role   *string `json:"Role,omitempty" xml:"Role,omitempty"`
}

func (s DescribeAccountRolesResponseBodyRolesRole) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccountRolesResponseBodyRolesRole) GoString() string {
	return s.String()
}

func (s *DescribeAccountRolesResponseBodyRolesRole) GetDBName() *string {
	return s.DBName
}

func (s *DescribeAccountRolesResponseBodyRolesRole) GetRole() *string {
	return s.Role
}

func (s *DescribeAccountRolesResponseBodyRolesRole) SetDBName(v string) *DescribeAccountRolesResponseBodyRolesRole {
	s.DBName = &v
	return s
}

func (s *DescribeAccountRolesResponseBodyRolesRole) SetRole(v string) *DescribeAccountRolesResponseBodyRolesRole {
	s.Role = &v
	return s
}

func (s *DescribeAccountRolesResponseBodyRolesRole) Validate() error {
	return dara.Validate(s)
}
