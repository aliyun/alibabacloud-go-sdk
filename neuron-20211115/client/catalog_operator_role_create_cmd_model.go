// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCatalogOperatorRoleCreateCmd interface {
	dara.Model
	String() string
	GoString() string
	SetCompanyId(v int64) *CatalogOperatorRoleCreateCmd
	GetCompanyId() *int64
	SetMarketId(v int64) *CatalogOperatorRoleCreateCmd
	GetMarketId() *int64
	SetMarketType(v string) *CatalogOperatorRoleCreateCmd
	GetMarketType() *string
}

type CatalogOperatorRoleCreateCmd struct {
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
	// 1
	MarketId *int64 `json:"marketId,omitempty" xml:"marketId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// inner
	MarketType *string `json:"marketType,omitempty" xml:"marketType,omitempty"`
}

func (s CatalogOperatorRoleCreateCmd) String() string {
	return dara.Prettify(s)
}

func (s CatalogOperatorRoleCreateCmd) GoString() string {
	return s.String()
}

func (s *CatalogOperatorRoleCreateCmd) GetCompanyId() *int64 {
	return s.CompanyId
}

func (s *CatalogOperatorRoleCreateCmd) GetMarketId() *int64 {
	return s.MarketId
}

func (s *CatalogOperatorRoleCreateCmd) GetMarketType() *string {
	return s.MarketType
}

func (s *CatalogOperatorRoleCreateCmd) SetCompanyId(v int64) *CatalogOperatorRoleCreateCmd {
	s.CompanyId = &v
	return s
}

func (s *CatalogOperatorRoleCreateCmd) SetMarketId(v int64) *CatalogOperatorRoleCreateCmd {
	s.MarketId = &v
	return s
}

func (s *CatalogOperatorRoleCreateCmd) SetMarketType(v string) *CatalogOperatorRoleCreateCmd {
	s.MarketType = &v
	return s
}

func (s *CatalogOperatorRoleCreateCmd) Validate() error {
	return dara.Validate(s)
}
