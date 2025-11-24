// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUsersWithPermissionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUserType(v string) *DescribeUsersWithPermissionsRequest
	GetUserType() *string
}

type DescribeUsersWithPermissionsRequest struct {
	// Specifies whether to query the IDs of all RAM users or RAM roles to which an RBAC role is assigned. Valid values:
	//
	// 	- `SubUser`: Query the IDs of all RAM users to which an RBAC role is assigned.
	//
	// 	- `RamRole`: Query the IDs of all RAM roles to which an RBAC role is assigned.
	//
	// This parameter is required.
	//
	// example:
	//
	// SubUser
	UserType *string `json:"UserType,omitempty" xml:"UserType,omitempty"`
}

func (s DescribeUsersWithPermissionsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeUsersWithPermissionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeUsersWithPermissionsRequest) GetUserType() *string {
	return s.UserType
}

func (s *DescribeUsersWithPermissionsRequest) SetUserType(v string) *DescribeUsersWithPermissionsRequest {
	s.UserType = &v
	return s
}

func (s *DescribeUsersWithPermissionsRequest) Validate() error {
	return dara.Validate(s)
}
