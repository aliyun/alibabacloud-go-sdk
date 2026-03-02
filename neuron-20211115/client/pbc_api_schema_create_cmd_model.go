// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPbcApiSchemaCreateCmd interface {
	dara.Model
	String() string
	GoString() string
	SetApiSchema(v string) *PbcApiSchemaCreateCmd
	GetApiSchema() *string
	SetCompanyId(v int64) *PbcApiSchemaCreateCmd
	GetCompanyId() *int64
	SetMarketId(v int64) *PbcApiSchemaCreateCmd
	GetMarketId() *int64
	SetPbcVersionId(v int64) *PbcApiSchemaCreateCmd
	GetPbcVersionId() *int64
}

type PbcApiSchemaCreateCmd struct {
	ApiSchema    *string `json:"apiSchema,omitempty" xml:"apiSchema,omitempty"`
	CompanyId    *int64  `json:"companyId,omitempty" xml:"companyId,omitempty"`
	MarketId     *int64  `json:"marketId,omitempty" xml:"marketId,omitempty"`
	PbcVersionId *int64  `json:"pbcVersionId,omitempty" xml:"pbcVersionId,omitempty"`
}

func (s PbcApiSchemaCreateCmd) String() string {
	return dara.Prettify(s)
}

func (s PbcApiSchemaCreateCmd) GoString() string {
	return s.String()
}

func (s *PbcApiSchemaCreateCmd) GetApiSchema() *string {
	return s.ApiSchema
}

func (s *PbcApiSchemaCreateCmd) GetCompanyId() *int64 {
	return s.CompanyId
}

func (s *PbcApiSchemaCreateCmd) GetMarketId() *int64 {
	return s.MarketId
}

func (s *PbcApiSchemaCreateCmd) GetPbcVersionId() *int64 {
	return s.PbcVersionId
}

func (s *PbcApiSchemaCreateCmd) SetApiSchema(v string) *PbcApiSchemaCreateCmd {
	s.ApiSchema = &v
	return s
}

func (s *PbcApiSchemaCreateCmd) SetCompanyId(v int64) *PbcApiSchemaCreateCmd {
	s.CompanyId = &v
	return s
}

func (s *PbcApiSchemaCreateCmd) SetMarketId(v int64) *PbcApiSchemaCreateCmd {
	s.MarketId = &v
	return s
}

func (s *PbcApiSchemaCreateCmd) SetPbcVersionId(v int64) *PbcApiSchemaCreateCmd {
	s.PbcVersionId = &v
	return s
}

func (s *PbcApiSchemaCreateCmd) Validate() error {
	return dara.Validate(s)
}
