// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPbcInstructionCreateCmd interface {
	dara.Model
	String() string
	GoString() string
	SetCompanyId(v int64) *PbcInstructionCreateCmd
	GetCompanyId() *int64
	SetDocument(v string) *PbcInstructionCreateCmd
	GetDocument() *string
	SetId(v int64) *PbcInstructionCreateCmd
	GetId() *int64
	SetMarketId(v int64) *PbcInstructionCreateCmd
	GetMarketId() *int64
	SetPbcVersionId(v int64) *PbcInstructionCreateCmd
	GetPbcVersionId() *int64
}

type PbcInstructionCreateCmd struct {
	CompanyId    *int64  `json:"companyId,omitempty" xml:"companyId,omitempty"`
	Document     *string `json:"document,omitempty" xml:"document,omitempty"`
	Id           *int64  `json:"id,omitempty" xml:"id,omitempty"`
	MarketId     *int64  `json:"marketId,omitempty" xml:"marketId,omitempty"`
	PbcVersionId *int64  `json:"pbcVersionId,omitempty" xml:"pbcVersionId,omitempty"`
}

func (s PbcInstructionCreateCmd) String() string {
	return dara.Prettify(s)
}

func (s PbcInstructionCreateCmd) GoString() string {
	return s.String()
}

func (s *PbcInstructionCreateCmd) GetCompanyId() *int64 {
	return s.CompanyId
}

func (s *PbcInstructionCreateCmd) GetDocument() *string {
	return s.Document
}

func (s *PbcInstructionCreateCmd) GetId() *int64 {
	return s.Id
}

func (s *PbcInstructionCreateCmd) GetMarketId() *int64 {
	return s.MarketId
}

func (s *PbcInstructionCreateCmd) GetPbcVersionId() *int64 {
	return s.PbcVersionId
}

func (s *PbcInstructionCreateCmd) SetCompanyId(v int64) *PbcInstructionCreateCmd {
	s.CompanyId = &v
	return s
}

func (s *PbcInstructionCreateCmd) SetDocument(v string) *PbcInstructionCreateCmd {
	s.Document = &v
	return s
}

func (s *PbcInstructionCreateCmd) SetId(v int64) *PbcInstructionCreateCmd {
	s.Id = &v
	return s
}

func (s *PbcInstructionCreateCmd) SetMarketId(v int64) *PbcInstructionCreateCmd {
	s.MarketId = &v
	return s
}

func (s *PbcInstructionCreateCmd) SetPbcVersionId(v int64) *PbcInstructionCreateCmd {
	s.PbcVersionId = &v
	return s
}

func (s *PbcInstructionCreateCmd) Validate() error {
	return dara.Validate(s)
}
