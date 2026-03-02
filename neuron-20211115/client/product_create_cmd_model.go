// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iProductCreateCmd interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *ProductCreateCmd
	GetAccountId() *string
	SetAlias(v string) *ProductCreateCmd
	GetAlias() *string
	SetCompanyId(v int64) *ProductCreateCmd
	GetCompanyId() *int64
	SetDescription(v string) *ProductCreateCmd
	GetDescription() *string
	SetName(v string) *ProductCreateCmd
	GetName() *string
}

type ProductCreateCmd struct {
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	// example:
	//
	// 多端商城
	Alias *string `json:"alias,omitempty" xml:"alias,omitempty"`
	// example:
	//
	// 1
	CompanyId   *int64  `json:"companyId,omitempty" xml:"companyId,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// yunmall
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ProductCreateCmd) String() string {
	return dara.Prettify(s)
}

func (s ProductCreateCmd) GoString() string {
	return s.String()
}

func (s *ProductCreateCmd) GetAccountId() *string {
	return s.AccountId
}

func (s *ProductCreateCmd) GetAlias() *string {
	return s.Alias
}

func (s *ProductCreateCmd) GetCompanyId() *int64 {
	return s.CompanyId
}

func (s *ProductCreateCmd) GetDescription() *string {
	return s.Description
}

func (s *ProductCreateCmd) GetName() *string {
	return s.Name
}

func (s *ProductCreateCmd) SetAccountId(v string) *ProductCreateCmd {
	s.AccountId = &v
	return s
}

func (s *ProductCreateCmd) SetAlias(v string) *ProductCreateCmd {
	s.Alias = &v
	return s
}

func (s *ProductCreateCmd) SetCompanyId(v int64) *ProductCreateCmd {
	s.CompanyId = &v
	return s
}

func (s *ProductCreateCmd) SetDescription(v string) *ProductCreateCmd {
	s.Description = &v
	return s
}

func (s *ProductCreateCmd) SetName(v string) *ProductCreateCmd {
	s.Name = &v
	return s
}

func (s *ProductCreateCmd) Validate() error {
	return dara.Validate(s)
}
