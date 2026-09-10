// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSingleSqlDryRunRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasourceName(v string) *SingleSqlDryRunRequest
	GetDatasourceName() *string
	SetSql(v string) *SingleSqlDryRunRequest
	GetSql() *string
}

type SingleSqlDryRunRequest struct {
	// The data source name.
	//
	// example:
	//
	// ds_demo
	DatasourceName *string `json:"datasourceName,omitempty" xml:"datasourceName,omitempty"`
	// The SQL statement.
	//
	// example:
	//
	// SELECT id, name FROM src_table WHERE ds = \\"20260116\\"
	Sql *string `json:"sql,omitempty" xml:"sql,omitempty"`
}

func (s SingleSqlDryRunRequest) String() string {
	return dara.Prettify(s)
}

func (s SingleSqlDryRunRequest) GoString() string {
	return s.String()
}

func (s *SingleSqlDryRunRequest) GetDatasourceName() *string {
	return s.DatasourceName
}

func (s *SingleSqlDryRunRequest) GetSql() *string {
	return s.Sql
}

func (s *SingleSqlDryRunRequest) SetDatasourceName(v string) *SingleSqlDryRunRequest {
	s.DatasourceName = &v
	return s
}

func (s *SingleSqlDryRunRequest) SetSql(v string) *SingleSqlDryRunRequest {
	s.Sql = &v
	return s
}

func (s *SingleSqlDryRunRequest) Validate() error {
	return dara.Validate(s)
}
