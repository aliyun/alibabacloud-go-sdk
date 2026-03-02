// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCompanyRoleCreateCmd interface {
	dara.Model
	String() string
	GoString() string
	SetCompanyId(v int64) *CompanyRoleCreateCmd
	GetCompanyId() *int64
	SetName(v string) *CompanyRoleCreateCmd
	GetName() *string
}

type CompanyRoleCreateCmd struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	CompanyId *int64 `json:"companyId,omitempty" xml:"companyId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 多端商城
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s CompanyRoleCreateCmd) String() string {
	return dara.Prettify(s)
}

func (s CompanyRoleCreateCmd) GoString() string {
	return s.String()
}

func (s *CompanyRoleCreateCmd) GetCompanyId() *int64 {
	return s.CompanyId
}

func (s *CompanyRoleCreateCmd) GetName() *string {
	return s.Name
}

func (s *CompanyRoleCreateCmd) SetCompanyId(v int64) *CompanyRoleCreateCmd {
	s.CompanyId = &v
	return s
}

func (s *CompanyRoleCreateCmd) SetName(v string) *CompanyRoleCreateCmd {
	s.Name = &v
	return s
}

func (s *CompanyRoleCreateCmd) Validate() error {
	return dara.Validate(s)
}
