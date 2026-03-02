// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPbcApiSchemaUpdateCmd interface {
	dara.Model
	String() string
	GoString() string
	SetApiSchema(v string) *PbcApiSchemaUpdateCmd
	GetApiSchema() *string
	SetPbcVersionId(v int64) *PbcApiSchemaUpdateCmd
	GetPbcVersionId() *int64
}

type PbcApiSchemaUpdateCmd struct {
	ApiSchema    *string `json:"apiSchema,omitempty" xml:"apiSchema,omitempty"`
	PbcVersionId *int64  `json:"pbcVersionId,omitempty" xml:"pbcVersionId,omitempty"`
}

func (s PbcApiSchemaUpdateCmd) String() string {
	return dara.Prettify(s)
}

func (s PbcApiSchemaUpdateCmd) GoString() string {
	return s.String()
}

func (s *PbcApiSchemaUpdateCmd) GetApiSchema() *string {
	return s.ApiSchema
}

func (s *PbcApiSchemaUpdateCmd) GetPbcVersionId() *int64 {
	return s.PbcVersionId
}

func (s *PbcApiSchemaUpdateCmd) SetApiSchema(v string) *PbcApiSchemaUpdateCmd {
	s.ApiSchema = &v
	return s
}

func (s *PbcApiSchemaUpdateCmd) SetPbcVersionId(v int64) *PbcApiSchemaUpdateCmd {
	s.PbcVersionId = &v
	return s
}

func (s *PbcApiSchemaUpdateCmd) Validate() error {
	return dara.Validate(s)
}
