// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPbcInvokeCreateCmd interface {
	dara.Model
	String() string
	GoString() string
	SetApplicant(v string) *PbcInvokeCreateCmd
	GetApplicant() *string
	SetCompanyId(v int64) *PbcInvokeCreateCmd
	GetCompanyId() *int64
	SetInvokePbcId(v int64) *PbcInvokeCreateCmd
	GetInvokePbcId() *int64
	SetMarketId(v int64) *PbcInvokeCreateCmd
	GetMarketId() *int64
	SetPbcId(v int64) *PbcInvokeCreateCmd
	GetPbcId() *int64
	SetReviewer(v string) *PbcInvokeCreateCmd
	GetReviewer() *string
	SetUsage(v string) *PbcInvokeCreateCmd
	GetUsage() *string
	SetVisible(v bool) *PbcInvokeCreateCmd
	GetVisible() *bool
}

type PbcInvokeCreateCmd struct {
	Applicant   *string `json:"applicant,omitempty" xml:"applicant,omitempty"`
	CompanyId   *int64  `json:"companyId,omitempty" xml:"companyId,omitempty"`
	InvokePbcId *int64  `json:"invokePbcId,omitempty" xml:"invokePbcId,omitempty"`
	MarketId    *int64  `json:"marketId,omitempty" xml:"marketId,omitempty"`
	PbcId       *int64  `json:"pbcId,omitempty" xml:"pbcId,omitempty"`
	Reviewer    *string `json:"reviewer,omitempty" xml:"reviewer,omitempty"`
	Usage       *string `json:"usage,omitempty" xml:"usage,omitempty"`
	Visible     *bool   `json:"visible,omitempty" xml:"visible,omitempty"`
}

func (s PbcInvokeCreateCmd) String() string {
	return dara.Prettify(s)
}

func (s PbcInvokeCreateCmd) GoString() string {
	return s.String()
}

func (s *PbcInvokeCreateCmd) GetApplicant() *string {
	return s.Applicant
}

func (s *PbcInvokeCreateCmd) GetCompanyId() *int64 {
	return s.CompanyId
}

func (s *PbcInvokeCreateCmd) GetInvokePbcId() *int64 {
	return s.InvokePbcId
}

func (s *PbcInvokeCreateCmd) GetMarketId() *int64 {
	return s.MarketId
}

func (s *PbcInvokeCreateCmd) GetPbcId() *int64 {
	return s.PbcId
}

func (s *PbcInvokeCreateCmd) GetReviewer() *string {
	return s.Reviewer
}

func (s *PbcInvokeCreateCmd) GetUsage() *string {
	return s.Usage
}

func (s *PbcInvokeCreateCmd) GetVisible() *bool {
	return s.Visible
}

func (s *PbcInvokeCreateCmd) SetApplicant(v string) *PbcInvokeCreateCmd {
	s.Applicant = &v
	return s
}

func (s *PbcInvokeCreateCmd) SetCompanyId(v int64) *PbcInvokeCreateCmd {
	s.CompanyId = &v
	return s
}

func (s *PbcInvokeCreateCmd) SetInvokePbcId(v int64) *PbcInvokeCreateCmd {
	s.InvokePbcId = &v
	return s
}

func (s *PbcInvokeCreateCmd) SetMarketId(v int64) *PbcInvokeCreateCmd {
	s.MarketId = &v
	return s
}

func (s *PbcInvokeCreateCmd) SetPbcId(v int64) *PbcInvokeCreateCmd {
	s.PbcId = &v
	return s
}

func (s *PbcInvokeCreateCmd) SetReviewer(v string) *PbcInvokeCreateCmd {
	s.Reviewer = &v
	return s
}

func (s *PbcInvokeCreateCmd) SetUsage(v string) *PbcInvokeCreateCmd {
	s.Usage = &v
	return s
}

func (s *PbcInvokeCreateCmd) SetVisible(v bool) *PbcInvokeCreateCmd {
	s.Visible = &v
	return s
}

func (s *PbcInvokeCreateCmd) Validate() error {
	return dara.Validate(s)
}
