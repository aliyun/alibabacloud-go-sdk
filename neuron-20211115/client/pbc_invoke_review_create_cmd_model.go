// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPbcInvokeReviewCreateCmd interface {
	dara.Model
	String() string
	GoString() string
	SetCompanyId(v int64) *PbcInvokeReviewCreateCmd
	GetCompanyId() *int64
	SetInvokePbcId(v int64) *PbcInvokeReviewCreateCmd
	GetInvokePbcId() *int64
	SetMarketId(v int64) *PbcInvokeReviewCreateCmd
	GetMarketId() *int64
	SetPbcId(v int64) *PbcInvokeReviewCreateCmd
	GetPbcId() *int64
	SetUsage(v string) *PbcInvokeReviewCreateCmd
	GetUsage() *string
}

type PbcInvokeReviewCreateCmd struct {
	CompanyId   *int64  `json:"companyId,omitempty" xml:"companyId,omitempty"`
	InvokePbcId *int64  `json:"invokePbcId,omitempty" xml:"invokePbcId,omitempty"`
	MarketId    *int64  `json:"marketId,omitempty" xml:"marketId,omitempty"`
	PbcId       *int64  `json:"pbcId,omitempty" xml:"pbcId,omitempty"`
	Usage       *string `json:"usage,omitempty" xml:"usage,omitempty"`
}

func (s PbcInvokeReviewCreateCmd) String() string {
	return dara.Prettify(s)
}

func (s PbcInvokeReviewCreateCmd) GoString() string {
	return s.String()
}

func (s *PbcInvokeReviewCreateCmd) GetCompanyId() *int64 {
	return s.CompanyId
}

func (s *PbcInvokeReviewCreateCmd) GetInvokePbcId() *int64 {
	return s.InvokePbcId
}

func (s *PbcInvokeReviewCreateCmd) GetMarketId() *int64 {
	return s.MarketId
}

func (s *PbcInvokeReviewCreateCmd) GetPbcId() *int64 {
	return s.PbcId
}

func (s *PbcInvokeReviewCreateCmd) GetUsage() *string {
	return s.Usage
}

func (s *PbcInvokeReviewCreateCmd) SetCompanyId(v int64) *PbcInvokeReviewCreateCmd {
	s.CompanyId = &v
	return s
}

func (s *PbcInvokeReviewCreateCmd) SetInvokePbcId(v int64) *PbcInvokeReviewCreateCmd {
	s.InvokePbcId = &v
	return s
}

func (s *PbcInvokeReviewCreateCmd) SetMarketId(v int64) *PbcInvokeReviewCreateCmd {
	s.MarketId = &v
	return s
}

func (s *PbcInvokeReviewCreateCmd) SetPbcId(v int64) *PbcInvokeReviewCreateCmd {
	s.PbcId = &v
	return s
}

func (s *PbcInvokeReviewCreateCmd) SetUsage(v string) *PbcInvokeReviewCreateCmd {
	s.Usage = &v
	return s
}

func (s *PbcInvokeReviewCreateCmd) Validate() error {
	return dara.Validate(s)
}
