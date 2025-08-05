// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUsersToRoleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdd(v []*string) *UpdateUsersToRoleRequest
	GetAdd() []*string
	SetRemove(v []*string) *UpdateUsersToRoleRequest
	GetRemove() []*string
}

type UpdateUsersToRoleRequest struct {
	// The accounts.
	Add []*string `json:"add,omitempty" xml:"add,omitempty" type:"Repeated"`
	// The accounts.
	Remove []*string `json:"remove,omitempty" xml:"remove,omitempty" type:"Repeated"`
}

func (s UpdateUsersToRoleRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateUsersToRoleRequest) GoString() string {
	return s.String()
}

func (s *UpdateUsersToRoleRequest) GetAdd() []*string {
	return s.Add
}

func (s *UpdateUsersToRoleRequest) GetRemove() []*string {
	return s.Remove
}

func (s *UpdateUsersToRoleRequest) SetAdd(v []*string) *UpdateUsersToRoleRequest {
	s.Add = v
	return s
}

func (s *UpdateUsersToRoleRequest) SetRemove(v []*string) *UpdateUsersToRoleRequest {
	s.Remove = v
	return s
}

func (s *UpdateUsersToRoleRequest) Validate() error {
	return dara.Validate(s)
}
