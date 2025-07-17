// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddTableToCategoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategoryId(v int64) *AddTableToCategoryRequest
	GetCategoryId() *int64
	SetDbId(v int64) *AddTableToCategoryRequest
	GetDbId() *int64
	SetTableName(v string) *AddTableToCategoryRequest
	GetTableName() *string
	SetTableSchemaName(v string) *AddTableToCategoryRequest
	GetTableSchemaName() *string
	SetTid(v int64) *AddTableToCategoryRequest
	GetTid() *int64
}

type AddTableToCategoryRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30000254257
	CategoryId *int64 `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1930****
	DbId *int64 `json:"DbId,omitempty" xml:"DbId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// table_name
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

func (s AddTableToCategoryRequest) String() string {
	return dara.Prettify(s)
}

func (s AddTableToCategoryRequest) GoString() string {
	return s.String()
}

func (s *AddTableToCategoryRequest) GetCategoryId() *int64 {
	return s.CategoryId
}

func (s *AddTableToCategoryRequest) GetDbId() *int64 {
	return s.DbId
}

func (s *AddTableToCategoryRequest) GetTableName() *string {
	return s.TableName
}

func (s *AddTableToCategoryRequest) GetTableSchemaName() *string {
	return s.TableSchemaName
}

func (s *AddTableToCategoryRequest) GetTid() *int64 {
	return s.Tid
}

func (s *AddTableToCategoryRequest) SetCategoryId(v int64) *AddTableToCategoryRequest {
	s.CategoryId = &v
	return s
}

func (s *AddTableToCategoryRequest) SetDbId(v int64) *AddTableToCategoryRequest {
	s.DbId = &v
	return s
}

func (s *AddTableToCategoryRequest) SetTableName(v string) *AddTableToCategoryRequest {
	s.TableName = &v
	return s
}

func (s *AddTableToCategoryRequest) SetTableSchemaName(v string) *AddTableToCategoryRequest {
	s.TableSchemaName = &v
	return s
}

func (s *AddTableToCategoryRequest) SetTid(v int64) *AddTableToCategoryRequest {
	s.Tid = &v
	return s
}

func (s *AddTableToCategoryRequest) Validate() error {
	return dara.Validate(s)
}
