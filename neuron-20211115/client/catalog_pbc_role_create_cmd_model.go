// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCatalogPbcRoleCreateCmd interface {
	dara.Model
	String() string
	GoString() string
	SetCompanyId(v int64) *CatalogPbcRoleCreateCmd
	GetCompanyId() *int64
	SetMarketId(v int64) *CatalogPbcRoleCreateCmd
	GetMarketId() *int64
	SetMarketResource(v string) *CatalogPbcRoleCreateCmd
	GetMarketResource() *string
	SetName(v string) *CatalogPbcRoleCreateCmd
	GetName() *string
	SetPbcResource(v string) *CatalogPbcRoleCreateCmd
	GetPbcResource() *string
}

type CatalogPbcRoleCreateCmd struct {
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
	// neuron:catalog:market
	MarketResource *string `json:"marketResource,omitempty" xml:"marketResource,omitempty"`
	// This parameter is required.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// neuron:catalog:company
	PbcResource *string `json:"pbcResource,omitempty" xml:"pbcResource,omitempty"`
}

func (s CatalogPbcRoleCreateCmd) String() string {
	return dara.Prettify(s)
}

func (s CatalogPbcRoleCreateCmd) GoString() string {
	return s.String()
}

func (s *CatalogPbcRoleCreateCmd) GetCompanyId() *int64 {
	return s.CompanyId
}

func (s *CatalogPbcRoleCreateCmd) GetMarketId() *int64 {
	return s.MarketId
}

func (s *CatalogPbcRoleCreateCmd) GetMarketResource() *string {
	return s.MarketResource
}

func (s *CatalogPbcRoleCreateCmd) GetName() *string {
	return s.Name
}

func (s *CatalogPbcRoleCreateCmd) GetPbcResource() *string {
	return s.PbcResource
}

func (s *CatalogPbcRoleCreateCmd) SetCompanyId(v int64) *CatalogPbcRoleCreateCmd {
	s.CompanyId = &v
	return s
}

func (s *CatalogPbcRoleCreateCmd) SetMarketId(v int64) *CatalogPbcRoleCreateCmd {
	s.MarketId = &v
	return s
}

func (s *CatalogPbcRoleCreateCmd) SetMarketResource(v string) *CatalogPbcRoleCreateCmd {
	s.MarketResource = &v
	return s
}

func (s *CatalogPbcRoleCreateCmd) SetName(v string) *CatalogPbcRoleCreateCmd {
	s.Name = &v
	return s
}

func (s *CatalogPbcRoleCreateCmd) SetPbcResource(v string) *CatalogPbcRoleCreateCmd {
	s.PbcResource = &v
	return s
}

func (s *CatalogPbcRoleCreateCmd) Validate() error {
	return dara.Validate(s)
}
