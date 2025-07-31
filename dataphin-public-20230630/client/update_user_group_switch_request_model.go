// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUserGroupSwitchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActive(v bool) *UpdateUserGroupSwitchRequest
	GetActive() *bool
	SetOpTenantId(v int64) *UpdateUserGroupSwitchRequest
	GetOpTenantId() *int64
	SetUserGroupId(v string) *UpdateUserGroupSwitchRequest
	GetUserGroupId() *string
}

type UpdateUserGroupSwitchRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// true
	Active *bool `json:"Active,omitempty" xml:"Active,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 31323
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
}

func (s UpdateUserGroupSwitchRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserGroupSwitchRequest) GoString() string {
	return s.String()
}

func (s *UpdateUserGroupSwitchRequest) GetActive() *bool {
	return s.Active
}

func (s *UpdateUserGroupSwitchRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *UpdateUserGroupSwitchRequest) GetUserGroupId() *string {
	return s.UserGroupId
}

func (s *UpdateUserGroupSwitchRequest) SetActive(v bool) *UpdateUserGroupSwitchRequest {
	s.Active = &v
	return s
}

func (s *UpdateUserGroupSwitchRequest) SetOpTenantId(v int64) *UpdateUserGroupSwitchRequest {
	s.OpTenantId = &v
	return s
}

func (s *UpdateUserGroupSwitchRequest) SetUserGroupId(v string) *UpdateUserGroupSwitchRequest {
	s.UserGroupId = &v
	return s
}

func (s *UpdateUserGroupSwitchRequest) Validate() error {
	return dara.Validate(s)
}
