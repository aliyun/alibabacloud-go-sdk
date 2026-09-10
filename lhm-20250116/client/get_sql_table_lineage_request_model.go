// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSqlTableLineageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDefaultSchema(v string) *GetSqlTableLineageRequest
	GetDefaultSchema() *string
	SetDialect(v string) *GetSqlTableLineageRequest
	GetDialect() *string
	SetSourceSqlScriptBase64(v string) *GetSqlTableLineageRequest
	GetSourceSqlScriptBase64() *string
}

type GetSqlTableLineageRequest struct {
	// The default schema (database) name, which is used to complete table references in the SQL script that do not explicitly specify a database name.
	//
	// example:
	//
	// db_demo
	DefaultSchema *string `json:"defaultSchema,omitempty" xml:"defaultSchema,omitempty"`
	// The SQL dialect.
	//
	// example:
	//
	// hive
	Dialect *string `json:"dialect,omitempty" xml:"dialect,omitempty"`
	// The source script content, Base64-encoded.
	//
	// example:
	//
	// U0VMRUNUICogRlJPTSB0Ow==
	SourceSqlScriptBase64 *string `json:"sourceSqlScriptBase64,omitempty" xml:"sourceSqlScriptBase64,omitempty"`
}

func (s GetSqlTableLineageRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSqlTableLineageRequest) GoString() string {
	return s.String()
}

func (s *GetSqlTableLineageRequest) GetDefaultSchema() *string {
	return s.DefaultSchema
}

func (s *GetSqlTableLineageRequest) GetDialect() *string {
	return s.Dialect
}

func (s *GetSqlTableLineageRequest) GetSourceSqlScriptBase64() *string {
	return s.SourceSqlScriptBase64
}

func (s *GetSqlTableLineageRequest) SetDefaultSchema(v string) *GetSqlTableLineageRequest {
	s.DefaultSchema = &v
	return s
}

func (s *GetSqlTableLineageRequest) SetDialect(v string) *GetSqlTableLineageRequest {
	s.Dialect = &v
	return s
}

func (s *GetSqlTableLineageRequest) SetSourceSqlScriptBase64(v string) *GetSqlTableLineageRequest {
	s.SourceSqlScriptBase64 = &v
	return s
}

func (s *GetSqlTableLineageRequest) Validate() error {
	return dara.Validate(s)
}
