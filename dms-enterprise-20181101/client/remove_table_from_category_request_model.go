// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveTableFromCategoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategoryId(v int64) *RemoveTableFromCategoryRequest
	GetCategoryId() *int64
	SetDbId(v int64) *RemoveTableFromCategoryRequest
	GetDbId() *int64
	SetTableName(v string) *RemoveTableFromCategoryRequest
	GetTableName() *string
	SetTableSchemaName(v string) *RemoveTableFromCategoryRequest
	GetTableSchemaName() *string
	SetTid(v int64) *RemoveTableFromCategoryRequest
	GetTid() *int64
}

type RemoveTableFromCategoryRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30000235594
	CategoryId *int64 `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 43153
	DbId *int64 `json:"DbId,omitempty" xml:"DbId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test_table
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// example:
	//
	// dbo
	TableSchemaName *string `json:"TableSchemaName,omitempty" xml:"TableSchemaName,omitempty"`
	// example:
	//
	// 3****
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s RemoveTableFromCategoryRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveTableFromCategoryRequest) GoString() string {
	return s.String()
}

func (s *RemoveTableFromCategoryRequest) GetCategoryId() *int64 {
	return s.CategoryId
}

func (s *RemoveTableFromCategoryRequest) GetDbId() *int64 {
	return s.DbId
}

func (s *RemoveTableFromCategoryRequest) GetTableName() *string {
	return s.TableName
}

func (s *RemoveTableFromCategoryRequest) GetTableSchemaName() *string {
	return s.TableSchemaName
}

func (s *RemoveTableFromCategoryRequest) GetTid() *int64 {
	return s.Tid
}

func (s *RemoveTableFromCategoryRequest) SetCategoryId(v int64) *RemoveTableFromCategoryRequest {
	s.CategoryId = &v
	return s
}

func (s *RemoveTableFromCategoryRequest) SetDbId(v int64) *RemoveTableFromCategoryRequest {
	s.DbId = &v
	return s
}

func (s *RemoveTableFromCategoryRequest) SetTableName(v string) *RemoveTableFromCategoryRequest {
	s.TableName = &v
	return s
}

func (s *RemoveTableFromCategoryRequest) SetTableSchemaName(v string) *RemoveTableFromCategoryRequest {
	s.TableSchemaName = &v
	return s
}

func (s *RemoveTableFromCategoryRequest) SetTid(v int64) *RemoveTableFromCategoryRequest {
	s.Tid = &v
	return s
}

func (s *RemoveTableFromCategoryRequest) Validate() error {
	return dara.Validate(s)
}
