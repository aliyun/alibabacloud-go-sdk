// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsertRoleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActionData(v string) *InsertRoleRequest
	GetActionData() *string
	SetRoleName(v string) *InsertRoleRequest
	GetRoleName() *string
}

type InsertRoleRequest struct {
	// The set of permissions to be granted to the role. The value is in the format of `Permission group ID 1:Permission serial number 1;...;Permission group ID n:Permission serial number n`. Example: `1:1;1:2;2:1;2:2`. For more information about permission groups and permission serial numbers, see [ListAuthority](https://help.aliyun.com/document_detail/149409.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 15:1;14:2
	ActionData *string `json:"ActionData,omitempty" xml:"ActionData,omitempty"`
	// The name of the role.
	//
	// This parameter is required.
	//
	// example:
	//
	// testrole
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
}

func (s InsertRoleRequest) String() string {
	return dara.Prettify(s)
}

func (s InsertRoleRequest) GoString() string {
	return s.String()
}

func (s *InsertRoleRequest) GetActionData() *string {
	return s.ActionData
}

func (s *InsertRoleRequest) GetRoleName() *string {
	return s.RoleName
}

func (s *InsertRoleRequest) SetActionData(v string) *InsertRoleRequest {
	s.ActionData = &v
	return s
}

func (s *InsertRoleRequest) SetRoleName(v string) *InsertRoleRequest {
	s.RoleName = &v
	return s
}

func (s *InsertRoleRequest) Validate() error {
	return dara.Validate(s)
}
