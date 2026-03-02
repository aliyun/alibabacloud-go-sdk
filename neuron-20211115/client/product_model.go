// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iProduct interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *Product
	GetAccountId() *string
	SetAlias(v string) *Product
	GetAlias() *string
	SetCompanyId(v int64) *Product
	GetCompanyId() *int64
	SetCreator(v string) *Product
	GetCreator() *string
	SetDescription(v string) *Product
	GetDescription() *string
	SetGmtCreate(v string) *Product
	GetGmtCreate() *string
	SetId(v int64) *Product
	GetId() *int64
	SetName(v string) *Product
	GetName() *string
	SetRequestId(v string) *Product
	GetRequestId() *string
	SetType(v string) *Product
	GetType() *string
}

type Product struct {
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	// example:
	//
	// 多端商城
	Alias *string `json:"alias,omitempty" xml:"alias,omitempty"`
	// example:
	//
	// 1
	CompanyId   *int64  `json:"companyId,omitempty" xml:"companyId,omitempty"`
	Creator     *string `json:"creator,omitempty" xml:"creator,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	GmtCreate   *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// 1
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// yunmall
	Name      *string `json:"name,omitempty" xml:"name,omitempty"`
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Type      *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s Product) String() string {
	return dara.Prettify(s)
}

func (s Product) GoString() string {
	return s.String()
}

func (s *Product) GetAccountId() *string {
	return s.AccountId
}

func (s *Product) GetAlias() *string {
	return s.Alias
}

func (s *Product) GetCompanyId() *int64 {
	return s.CompanyId
}

func (s *Product) GetCreator() *string {
	return s.Creator
}

func (s *Product) GetDescription() *string {
	return s.Description
}

func (s *Product) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *Product) GetId() *int64 {
	return s.Id
}

func (s *Product) GetName() *string {
	return s.Name
}

func (s *Product) GetRequestId() *string {
	return s.RequestId
}

func (s *Product) GetType() *string {
	return s.Type
}

func (s *Product) SetAccountId(v string) *Product {
	s.AccountId = &v
	return s
}

func (s *Product) SetAlias(v string) *Product {
	s.Alias = &v
	return s
}

func (s *Product) SetCompanyId(v int64) *Product {
	s.CompanyId = &v
	return s
}

func (s *Product) SetCreator(v string) *Product {
	s.Creator = &v
	return s
}

func (s *Product) SetDescription(v string) *Product {
	s.Description = &v
	return s
}

func (s *Product) SetGmtCreate(v string) *Product {
	s.GmtCreate = &v
	return s
}

func (s *Product) SetId(v int64) *Product {
	s.Id = &v
	return s
}

func (s *Product) SetName(v string) *Product {
	s.Name = &v
	return s
}

func (s *Product) SetRequestId(v string) *Product {
	s.RequestId = &v
	return s
}

func (s *Product) SetType(v string) *Product {
	s.Type = &v
	return s
}

func (s *Product) Validate() error {
	return dara.Validate(s)
}
