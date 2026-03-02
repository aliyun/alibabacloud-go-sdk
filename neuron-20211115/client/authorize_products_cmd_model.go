// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthorizeProductsCmd interface {
	dara.Model
	String() string
	GoString() string
	SetCompanyId(v int64) *AuthorizeProductsCmd
	GetCompanyId() *int64
	SetProductIds(v []*int64) *AuthorizeProductsCmd
	GetProductIds() []*int64
}

type AuthorizeProductsCmd struct {
	CompanyId  *int64   `json:"companyId,omitempty" xml:"companyId,omitempty"`
	ProductIds []*int64 `json:"productIds,omitempty" xml:"productIds,omitempty" type:"Repeated"`
}

func (s AuthorizeProductsCmd) String() string {
	return dara.Prettify(s)
}

func (s AuthorizeProductsCmd) GoString() string {
	return s.String()
}

func (s *AuthorizeProductsCmd) GetCompanyId() *int64 {
	return s.CompanyId
}

func (s *AuthorizeProductsCmd) GetProductIds() []*int64 {
	return s.ProductIds
}

func (s *AuthorizeProductsCmd) SetCompanyId(v int64) *AuthorizeProductsCmd {
	s.CompanyId = &v
	return s
}

func (s *AuthorizeProductsCmd) SetProductIds(v []*int64) *AuthorizeProductsCmd {
	s.ProductIds = v
	return s
}

func (s *AuthorizeProductsCmd) Validate() error {
	return dara.Validate(s)
}
