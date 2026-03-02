// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iProductInfo interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *ProductInfo
	GetAccountId() *string
	SetCompanyId(v int64) *ProductInfo
	GetCompanyId() *int64
	SetId(v int64) *ProductInfo
	GetId() *int64
	SetIsAuthorized(v bool) *ProductInfo
	GetIsAuthorized() *bool
	SetName(v string) *ProductInfo
	GetName() *string
}

type ProductInfo struct {
	AccountId    *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	CompanyId    *int64  `json:"companyId,omitempty" xml:"companyId,omitempty"`
	Id           *int64  `json:"id,omitempty" xml:"id,omitempty"`
	IsAuthorized *bool   `json:"isAuthorized,omitempty" xml:"isAuthorized,omitempty"`
	Name         *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ProductInfo) String() string {
	return dara.Prettify(s)
}

func (s ProductInfo) GoString() string {
	return s.String()
}

func (s *ProductInfo) GetAccountId() *string {
	return s.AccountId
}

func (s *ProductInfo) GetCompanyId() *int64 {
	return s.CompanyId
}

func (s *ProductInfo) GetId() *int64 {
	return s.Id
}

func (s *ProductInfo) GetIsAuthorized() *bool {
	return s.IsAuthorized
}

func (s *ProductInfo) GetName() *string {
	return s.Name
}

func (s *ProductInfo) SetAccountId(v string) *ProductInfo {
	s.AccountId = &v
	return s
}

func (s *ProductInfo) SetCompanyId(v int64) *ProductInfo {
	s.CompanyId = &v
	return s
}

func (s *ProductInfo) SetId(v int64) *ProductInfo {
	s.Id = &v
	return s
}

func (s *ProductInfo) SetIsAuthorized(v bool) *ProductInfo {
	s.IsAuthorized = &v
	return s
}

func (s *ProductInfo) SetName(v string) *ProductInfo {
	s.Name = &v
	return s
}

func (s *ProductInfo) Validate() error {
	return dara.Validate(s)
}
