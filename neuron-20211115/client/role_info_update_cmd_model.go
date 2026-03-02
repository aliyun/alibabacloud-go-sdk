// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRoleInfoUpdateCmd interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *RoleInfoUpdateCmd
	GetId() *int64
	SetName(v string) *RoleInfoUpdateCmd
	GetName() *string
}

type RoleInfoUpdateCmd struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 开发者
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s RoleInfoUpdateCmd) String() string {
	return dara.Prettify(s)
}

func (s RoleInfoUpdateCmd) GoString() string {
	return s.String()
}

func (s *RoleInfoUpdateCmd) GetId() *int64 {
	return s.Id
}

func (s *RoleInfoUpdateCmd) GetName() *string {
	return s.Name
}

func (s *RoleInfoUpdateCmd) SetId(v int64) *RoleInfoUpdateCmd {
	s.Id = &v
	return s
}

func (s *RoleInfoUpdateCmd) SetName(v string) *RoleInfoUpdateCmd {
	s.Name = &v
	return s
}

func (s *RoleInfoUpdateCmd) Validate() error {
	return dara.Validate(s)
}
