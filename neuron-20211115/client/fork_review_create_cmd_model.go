// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iForkReviewCreateCmd interface {
	dara.Model
	String() string
	GoString() string
	SetCompanyId(v int64) *ForkReviewCreateCmd
	GetCompanyId() *int64
	SetGitGroup(v string) *ForkReviewCreateCmd
	GetGitGroup() *string
	SetMarketId(v int64) *ForkReviewCreateCmd
	GetMarketId() *int64
	SetPbcId(v int64) *ForkReviewCreateCmd
	GetPbcId() *int64
	SetUsage(v string) *ForkReviewCreateCmd
	GetUsage() *string
}

type ForkReviewCreateCmd struct {
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
	// 商城国际版本
	Usage *string `json:"usage,omitempty" xml:"usage,omitempty"`
}

func (s ForkReviewCreateCmd) String() string {
	return dara.Prettify(s)
}

func (s ForkReviewCreateCmd) GoString() string {
	return s.String()
}

func (s *ForkReviewCreateCmd) GetCompanyId() *int64 {
	return s.CompanyId
}

func (s *ForkReviewCreateCmd) GetGitGroup() *string {
	return s.GitGroup
}

func (s *ForkReviewCreateCmd) GetMarketId() *int64 {
	return s.MarketId
}

func (s *ForkReviewCreateCmd) GetPbcId() *int64 {
	return s.PbcId
}

func (s *ForkReviewCreateCmd) GetUsage() *string {
	return s.Usage
}

func (s *ForkReviewCreateCmd) SetCompanyId(v int64) *ForkReviewCreateCmd {
	s.CompanyId = &v
	return s
}

func (s *ForkReviewCreateCmd) SetGitGroup(v string) *ForkReviewCreateCmd {
	s.GitGroup = &v
	return s
}

func (s *ForkReviewCreateCmd) SetMarketId(v int64) *ForkReviewCreateCmd {
	s.MarketId = &v
	return s
}

func (s *ForkReviewCreateCmd) SetPbcId(v int64) *ForkReviewCreateCmd {
	s.PbcId = &v
	return s
}

func (s *ForkReviewCreateCmd) SetUsage(v string) *ForkReviewCreateCmd {
	s.Usage = &v
	return s
}

func (s *ForkReviewCreateCmd) Validate() error {
	return dara.Validate(s)
}
