// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTablesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTableName(v string) *GetTablesRequest
	GetTableName() *string
}

type GetTablesRequest struct {
	// example:
	//
	// item
	TableName *string `json:"tableName,omitempty" xml:"tableName,omitempty"`
}

func (s GetTablesRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTablesRequest) GoString() string {
	return s.String()
}

func (s *GetTablesRequest) GetTableName() *string {
	return s.TableName
}

func (s *GetTablesRequest) SetTableName(v string) *GetTablesRequest {
	s.TableName = &v
	return s
}

func (s *GetTablesRequest) Validate() error {
	return dara.Validate(s)
}
