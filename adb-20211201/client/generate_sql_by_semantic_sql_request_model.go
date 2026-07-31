// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateSqlBySemanticSqlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *GenerateSqlBySemanticSqlRequest
	GetDBClusterId() *string
	SetSchemaName(v string) *GenerateSqlBySemanticSqlRequest
	GetSchemaName() *string
	SetSql(v string) *GenerateSqlBySemanticSqlRequest
	GetSql() *string
}

type GenerateSqlBySemanticSqlRequest struct {
	// The ID of the ADB cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-uf6wjk5xxxxxxxxxx
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The schema name.
	//
	// example:
	//
	// adb_demo
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	// The SQL statement that queries the semantic view.
	//
	// This parameter is required.
	//
	// example:
	//
	// select sum(amount) from sv_sales
	Sql *string `json:"Sql,omitempty" xml:"Sql,omitempty"`
}

func (s GenerateSqlBySemanticSqlRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateSqlBySemanticSqlRequest) GoString() string {
	return s.String()
}

func (s *GenerateSqlBySemanticSqlRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *GenerateSqlBySemanticSqlRequest) GetSchemaName() *string {
	return s.SchemaName
}

func (s *GenerateSqlBySemanticSqlRequest) GetSql() *string {
	return s.Sql
}

func (s *GenerateSqlBySemanticSqlRequest) SetDBClusterId(v string) *GenerateSqlBySemanticSqlRequest {
	s.DBClusterId = &v
	return s
}

func (s *GenerateSqlBySemanticSqlRequest) SetSchemaName(v string) *GenerateSqlBySemanticSqlRequest {
	s.SchemaName = &v
	return s
}

func (s *GenerateSqlBySemanticSqlRequest) SetSql(v string) *GenerateSqlBySemanticSqlRequest {
	s.Sql = &v
	return s
}

func (s *GenerateSqlBySemanticSqlRequest) Validate() error {
	return dara.Validate(s)
}
