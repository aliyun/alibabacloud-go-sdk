// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTableColumnsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbId(v int32) *ListTableColumnsRequest
	GetDbId() *int32
	SetTableName(v string) *ListTableColumnsRequest
	GetTableName() *string
	SetTableSchemaName(v string) *ListTableColumnsRequest
	GetTableSchemaName() *string
	SetTid(v int64) *ListTableColumnsRequest
	GetTid() *int64
}

type ListTableColumnsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123
	DbId *int32 `json:"DbId,omitempty" xml:"DbId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 100g_customer
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// example:
	//
	// dbo
	TableSchemaName *string `json:"TableSchemaName,omitempty" xml:"TableSchemaName,omitempty"`
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s ListTableColumnsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTableColumnsRequest) GoString() string {
	return s.String()
}

func (s *ListTableColumnsRequest) GetDbId() *int32 {
	return s.DbId
}

func (s *ListTableColumnsRequest) GetTableName() *string {
	return s.TableName
}

func (s *ListTableColumnsRequest) GetTableSchemaName() *string {
	return s.TableSchemaName
}

func (s *ListTableColumnsRequest) GetTid() *int64 {
	return s.Tid
}

func (s *ListTableColumnsRequest) SetDbId(v int32) *ListTableColumnsRequest {
	s.DbId = &v
	return s
}

func (s *ListTableColumnsRequest) SetTableName(v string) *ListTableColumnsRequest {
	s.TableName = &v
	return s
}

func (s *ListTableColumnsRequest) SetTableSchemaName(v string) *ListTableColumnsRequest {
	s.TableSchemaName = &v
	return s
}

func (s *ListTableColumnsRequest) SetTid(v int64) *ListTableColumnsRequest {
	s.Tid = &v
	return s
}

func (s *ListTableColumnsRequest) Validate() error {
	return dara.Validate(s)
}
