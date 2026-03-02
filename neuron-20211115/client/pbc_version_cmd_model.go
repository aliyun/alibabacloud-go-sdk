// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPbcVersionCmd interface {
	dara.Model
	String() string
	GoString() string
	SetApiSchema(v string) *PbcVersionCmd
	GetApiSchema() *string
	SetCompanyId(v int64) *PbcVersionCmd
	GetCompanyId() *int64
	SetDeveloperId(v int64) *PbcVersionCmd
	GetDeveloperId() *int64
	SetDocument(v string) *PbcVersionCmd
	GetDocument() *string
	SetMarketId(v int64) *PbcVersionCmd
	GetMarketId() *int64
	SetName(v string) *PbcVersionCmd
	GetName() *string
	SetSchema(v string) *PbcVersionCmd
	GetSchema() *string
	SetVersion(v string) *PbcVersionCmd
	GetVersion() *string
}

type PbcVersionCmd struct {
	ApiSchema   *string `json:"apiSchema,omitempty" xml:"apiSchema,omitempty"`
	CompanyId   *int64  `json:"companyId,omitempty" xml:"companyId,omitempty"`
	DeveloperId *int64  `json:"developerId,omitempty" xml:"developerId,omitempty"`
	Document    *string `json:"document,omitempty" xml:"document,omitempty"`
	MarketId    *int64  `json:"marketId,omitempty" xml:"marketId,omitempty"`
	Name        *string `json:"name,omitempty" xml:"name,omitempty"`
	Schema      *string `json:"schema,omitempty" xml:"schema,omitempty"`
	Version     *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s PbcVersionCmd) String() string {
	return dara.Prettify(s)
}

func (s PbcVersionCmd) GoString() string {
	return s.String()
}

func (s *PbcVersionCmd) GetApiSchema() *string {
	return s.ApiSchema
}

func (s *PbcVersionCmd) GetCompanyId() *int64 {
	return s.CompanyId
}

func (s *PbcVersionCmd) GetDeveloperId() *int64 {
	return s.DeveloperId
}

func (s *PbcVersionCmd) GetDocument() *string {
	return s.Document
}

func (s *PbcVersionCmd) GetMarketId() *int64 {
	return s.MarketId
}

func (s *PbcVersionCmd) GetName() *string {
	return s.Name
}

func (s *PbcVersionCmd) GetSchema() *string {
	return s.Schema
}

func (s *PbcVersionCmd) GetVersion() *string {
	return s.Version
}

func (s *PbcVersionCmd) SetApiSchema(v string) *PbcVersionCmd {
	s.ApiSchema = &v
	return s
}

func (s *PbcVersionCmd) SetCompanyId(v int64) *PbcVersionCmd {
	s.CompanyId = &v
	return s
}

func (s *PbcVersionCmd) SetDeveloperId(v int64) *PbcVersionCmd {
	s.DeveloperId = &v
	return s
}

func (s *PbcVersionCmd) SetDocument(v string) *PbcVersionCmd {
	s.Document = &v
	return s
}

func (s *PbcVersionCmd) SetMarketId(v int64) *PbcVersionCmd {
	s.MarketId = &v
	return s
}

func (s *PbcVersionCmd) SetName(v string) *PbcVersionCmd {
	s.Name = &v
	return s
}

func (s *PbcVersionCmd) SetSchema(v string) *PbcVersionCmd {
	s.Schema = &v
	return s
}

func (s *PbcVersionCmd) SetVersion(v string) *PbcVersionCmd {
	s.Version = &v
	return s
}

func (s *PbcVersionCmd) Validate() error {
	return dara.Validate(s)
}
