// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRolesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeRolesResponseBody
	GetRequestId() *string
	SetRoleList(v *DescribeRolesResponseBodyRoleList) *DescribeRolesResponseBody
	GetRoleList() *DescribeRolesResponseBodyRoleList
}

type DescribeRolesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D0521
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RoleList  *DescribeRolesResponseBodyRoleList `json:"RoleList,omitempty" xml:"RoleList,omitempty" type:"Struct"`
}

func (s DescribeRolesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRolesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRolesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRolesResponseBody) GetRoleList() *DescribeRolesResponseBodyRoleList {
	return s.RoleList
}

func (s *DescribeRolesResponseBody) SetRequestId(v string) *DescribeRolesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRolesResponseBody) SetRoleList(v *DescribeRolesResponseBodyRoleList) *DescribeRolesResponseBody {
	s.RoleList = v
	return s
}

func (s *DescribeRolesResponseBody) Validate() error {
	if s.RoleList != nil {
		if err := s.RoleList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeRolesResponseBodyRoleList struct {
	Role []*string `json:"Role,omitempty" xml:"Role,omitempty" type:"Repeated"`
}

func (s DescribeRolesResponseBodyRoleList) String() string {
	return dara.Prettify(s)
}

func (s DescribeRolesResponseBodyRoleList) GoString() string {
	return s.String()
}

func (s *DescribeRolesResponseBodyRoleList) GetRole() []*string {
	return s.Role
}

func (s *DescribeRolesResponseBodyRoleList) SetRole(v []*string) *DescribeRolesResponseBodyRoleList {
	s.Role = v
	return s
}

func (s *DescribeRolesResponseBodyRoleList) Validate() error {
	return dara.Validate(s)
}
