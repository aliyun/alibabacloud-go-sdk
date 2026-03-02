// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReposForkCreateCmd interface {
	dara.Model
	String() string
	GoString() string
	SetCompanyId(v int64) *ReposForkCreateCmd
	GetCompanyId() *int64
	SetGitGroup(v string) *ReposForkCreateCmd
	GetGitGroup() *string
	SetMarketId(v int64) *ReposForkCreateCmd
	GetMarketId() *int64
	SetPbcId(v int64) *ReposForkCreateCmd
	GetPbcId() *int64
	SetUsage(v string) *ReposForkCreateCmd
	GetUsage() *string
	SetVisible(v bool) *ReposForkCreateCmd
	GetVisible() *bool
}

type ReposForkCreateCmd struct {
	CompanyId *int64 `json:"companyId,omitempty" xml:"companyId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// global-mall
	GitGroup *string `json:"gitGroup,omitempty" xml:"gitGroup,omitempty"`
	MarketId *int64  `json:"marketId,omitempty" xml:"marketId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PbcId *int64 `json:"pbcId,omitempty" xml:"pbcId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 商城国际化版本
	Usage   *string `json:"usage,omitempty" xml:"usage,omitempty"`
	Visible *bool   `json:"visible,omitempty" xml:"visible,omitempty"`
}

func (s ReposForkCreateCmd) String() string {
	return dara.Prettify(s)
}

func (s ReposForkCreateCmd) GoString() string {
	return s.String()
}

func (s *ReposForkCreateCmd) GetCompanyId() *int64 {
	return s.CompanyId
}

func (s *ReposForkCreateCmd) GetGitGroup() *string {
	return s.GitGroup
}

func (s *ReposForkCreateCmd) GetMarketId() *int64 {
	return s.MarketId
}

func (s *ReposForkCreateCmd) GetPbcId() *int64 {
	return s.PbcId
}

func (s *ReposForkCreateCmd) GetUsage() *string {
	return s.Usage
}

func (s *ReposForkCreateCmd) GetVisible() *bool {
	return s.Visible
}

func (s *ReposForkCreateCmd) SetCompanyId(v int64) *ReposForkCreateCmd {
	s.CompanyId = &v
	return s
}

func (s *ReposForkCreateCmd) SetGitGroup(v string) *ReposForkCreateCmd {
	s.GitGroup = &v
	return s
}

func (s *ReposForkCreateCmd) SetMarketId(v int64) *ReposForkCreateCmd {
	s.MarketId = &v
	return s
}

func (s *ReposForkCreateCmd) SetPbcId(v int64) *ReposForkCreateCmd {
	s.PbcId = &v
	return s
}

func (s *ReposForkCreateCmd) SetUsage(v string) *ReposForkCreateCmd {
	s.Usage = &v
	return s
}

func (s *ReposForkCreateCmd) SetVisible(v bool) *ReposForkCreateCmd {
	s.Visible = &v
	return s
}

func (s *ReposForkCreateCmd) Validate() error {
	return dara.Validate(s)
}
