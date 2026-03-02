// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPbcCreateCmd interface {
	dara.Model
	String() string
	GoString() string
	SetAlias(v string) *PbcCreateCmd
	GetAlias() *string
	SetCompany(v string) *PbcCreateCmd
	GetCompany() *string
	SetCompanyId(v int64) *PbcCreateCmd
	GetCompanyId() *int64
	SetDescription(v string) *PbcCreateCmd
	GetDescription() *string
	SetDeveloperId(v string) *PbcCreateCmd
	GetDeveloperId() *string
	SetIndustry(v string) *PbcCreateCmd
	GetIndustry() *string
	SetMarketId(v int64) *PbcCreateCmd
	GetMarketId() *int64
	SetName(v string) *PbcCreateCmd
	GetName() *string
	SetType(v string) *PbcCreateCmd
	GetType() *string
}

type PbcCreateCmd struct {
	// example:
	//
	// 基础商品
	Alias *string `json:"alias,omitempty" xml:"alias,omitempty"`
	// example:
	//
	// 企业服务
	Company *string `json:"company,omitempty" xml:"company,omitempty"`
	// example:
	//
	// 12
	CompanyId *int64 `json:"companyId,omitempty" xml:"companyId,omitempty"`
	// example:
	//
	// 基础商品PBC
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// 1
	DeveloperId *string `json:"developerId,omitempty" xml:"developerId,omitempty"`
	// example:
	//
	// common
	Industry *string `json:"industry,omitempty" xml:"industry,omitempty"`
	MarketId *int64  `json:"marketId,omitempty" xml:"marketId,omitempty"`
	// example:
	//
	// product
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// foundation
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s PbcCreateCmd) String() string {
	return dara.Prettify(s)
}

func (s PbcCreateCmd) GoString() string {
	return s.String()
}

func (s *PbcCreateCmd) GetAlias() *string {
	return s.Alias
}

func (s *PbcCreateCmd) GetCompany() *string {
	return s.Company
}

func (s *PbcCreateCmd) GetCompanyId() *int64 {
	return s.CompanyId
}

func (s *PbcCreateCmd) GetDescription() *string {
	return s.Description
}

func (s *PbcCreateCmd) GetDeveloperId() *string {
	return s.DeveloperId
}

func (s *PbcCreateCmd) GetIndustry() *string {
	return s.Industry
}

func (s *PbcCreateCmd) GetMarketId() *int64 {
	return s.MarketId
}

func (s *PbcCreateCmd) GetName() *string {
	return s.Name
}

func (s *PbcCreateCmd) GetType() *string {
	return s.Type
}

func (s *PbcCreateCmd) SetAlias(v string) *PbcCreateCmd {
	s.Alias = &v
	return s
}

func (s *PbcCreateCmd) SetCompany(v string) *PbcCreateCmd {
	s.Company = &v
	return s
}

func (s *PbcCreateCmd) SetCompanyId(v int64) *PbcCreateCmd {
	s.CompanyId = &v
	return s
}

func (s *PbcCreateCmd) SetDescription(v string) *PbcCreateCmd {
	s.Description = &v
	return s
}

func (s *PbcCreateCmd) SetDeveloperId(v string) *PbcCreateCmd {
	s.DeveloperId = &v
	return s
}

func (s *PbcCreateCmd) SetIndustry(v string) *PbcCreateCmd {
	s.Industry = &v
	return s
}

func (s *PbcCreateCmd) SetMarketId(v int64) *PbcCreateCmd {
	s.MarketId = &v
	return s
}

func (s *PbcCreateCmd) SetName(v string) *PbcCreateCmd {
	s.Name = &v
	return s
}

func (s *PbcCreateCmd) SetType(v string) *PbcCreateCmd {
	s.Type = &v
	return s
}

func (s *PbcCreateCmd) Validate() error {
	return dara.Validate(s)
}
