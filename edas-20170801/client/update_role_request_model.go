// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRoleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActionData(v string) *UpdateRoleRequest
	GetActionData() *string
	SetRoleId(v int32) *UpdateRoleRequest
	GetRoleId() *int32
}

type UpdateRoleRequest struct {
	// The set of permissions to be granted to the role. The value is in the format of `Permission group ID 1:Permission serial number 1;...;Permission group ID n:Permission serial number n`. Example: `1:1;1:2;2:1;2:2`. For more information about permission groups and permission serial numbers, see [ListAuthority](https://help.aliyun.com/document_detail/149409.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 15:1
	ActionData *string `json:"ActionData,omitempty" xml:"ActionData,omitempty"`
	// The ID of the role. You can call the ListRole operation to query the role IDs. For more information, see [ListRole](https://help.aliyun.com/document_detail/149410.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 32371
	RoleId *int32 `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
}

func (s UpdateRoleRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateRoleRequest) GoString() string {
	return s.String()
}

func (s *UpdateRoleRequest) GetActionData() *string {
	return s.ActionData
}

func (s *UpdateRoleRequest) GetRoleId() *int32 {
	return s.RoleId
}

func (s *UpdateRoleRequest) SetActionData(v string) *UpdateRoleRequest {
	s.ActionData = &v
	return s
}

func (s *UpdateRoleRequest) SetRoleId(v int32) *UpdateRoleRequest {
	s.RoleId = &v
	return s
}

func (s *UpdateRoleRequest) Validate() error {
	return dara.Validate(s)
}
