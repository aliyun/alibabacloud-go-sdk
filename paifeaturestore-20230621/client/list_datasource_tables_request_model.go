// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDatasourceTablesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTableName(v string) *ListDatasourceTablesRequest
	GetTableName() *string
}

type ListDatasourceTablesRequest struct {
	// example:
	//
	// table1
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s ListDatasourceTablesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDatasourceTablesRequest) GoString() string {
	return s.String()
}

func (s *ListDatasourceTablesRequest) GetTableName() *string {
	return s.TableName
}

func (s *ListDatasourceTablesRequest) SetTableName(v string) *ListDatasourceTablesRequest {
	s.TableName = &v
	return s
}

func (s *ListDatasourceTablesRequest) Validate() error {
	return dara.Validate(s)
}
