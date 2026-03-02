// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPbcSchemaCreateCmd interface {
	dara.Model
	String() string
	GoString() string
	SetCompanyId(v int64) *PbcSchemaCreateCmd
	GetCompanyId() *int64
	SetId(v int64) *PbcSchemaCreateCmd
	GetId() *int64
	SetMarketId(v int64) *PbcSchemaCreateCmd
	GetMarketId() *int64
	SetPbcName(v string) *PbcSchemaCreateCmd
	GetPbcName() *string
	SetPbcVersion(v string) *PbcSchemaCreateCmd
	GetPbcVersion() *string
	SetSchema(v string) *PbcSchemaCreateCmd
	GetSchema() *string
}

type PbcSchemaCreateCmd struct {
	CompanyId  *int64  `json:"companyId,omitempty" xml:"companyId,omitempty"`
	Id         *int64  `json:"id,omitempty" xml:"id,omitempty"`
	MarketId   *int64  `json:"marketId,omitempty" xml:"marketId,omitempty"`
	PbcName    *string `json:"pbcName,omitempty" xml:"pbcName,omitempty"`
	PbcVersion *string `json:"pbcVersion,omitempty" xml:"pbcVersion,omitempty"`
	Schema     *string `json:"schema,omitempty" xml:"schema,omitempty"`
}

func (s PbcSchemaCreateCmd) String() string {
	return dara.Prettify(s)
}

func (s PbcSchemaCreateCmd) GoString() string {
	return s.String()
}

func (s *PbcSchemaCreateCmd) GetCompanyId() *int64 {
	return s.CompanyId
}

func (s *PbcSchemaCreateCmd) GetId() *int64 {
	return s.Id
}

func (s *PbcSchemaCreateCmd) GetMarketId() *int64 {
	return s.MarketId
}

func (s *PbcSchemaCreateCmd) GetPbcName() *string {
	return s.PbcName
}

func (s *PbcSchemaCreateCmd) GetPbcVersion() *string {
	return s.PbcVersion
}

func (s *PbcSchemaCreateCmd) GetSchema() *string {
	return s.Schema
}

func (s *PbcSchemaCreateCmd) SetCompanyId(v int64) *PbcSchemaCreateCmd {
	s.CompanyId = &v
	return s
}

func (s *PbcSchemaCreateCmd) SetId(v int64) *PbcSchemaCreateCmd {
	s.Id = &v
	return s
}

func (s *PbcSchemaCreateCmd) SetMarketId(v int64) *PbcSchemaCreateCmd {
	s.MarketId = &v
	return s
}

func (s *PbcSchemaCreateCmd) SetPbcName(v string) *PbcSchemaCreateCmd {
	s.PbcName = &v
	return s
}

func (s *PbcSchemaCreateCmd) SetPbcVersion(v string) *PbcSchemaCreateCmd {
	s.PbcVersion = &v
	return s
}

func (s *PbcSchemaCreateCmd) SetSchema(v string) *PbcSchemaCreateCmd {
	s.Schema = &v
	return s
}

func (s *PbcSchemaCreateCmd) Validate() error {
	return dara.Validate(s)
}
