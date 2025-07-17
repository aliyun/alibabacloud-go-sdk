// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataLakeFunctionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCatalogName(v string) *ListDataLakeFunctionRequest
	GetCatalogName() *string
	SetDataRegion(v string) *ListDataLakeFunctionRequest
	GetDataRegion() *string
	SetDbName(v string) *ListDataLakeFunctionRequest
	GetDbName() *string
	SetFunctionNamePattern(v string) *ListDataLakeFunctionRequest
	GetFunctionNamePattern() *string
	SetMaxResults(v int32) *ListDataLakeFunctionRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListDataLakeFunctionRequest
	GetNextToken() *string
	SetTid(v int64) *ListDataLakeFunctionRequest
	GetTid() *int64
	SetWorkspaceId(v int64) *ListDataLakeFunctionRequest
	GetWorkspaceId() *int64
}

type ListDataLakeFunctionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// hive
	CatalogName *string `json:"CatalogName,omitempty" xml:"CatalogName,omitempty"`
	// This parameter is required.
	//
	// if can be null:
	// false
	//
	// example:
	//
	// cn-hangzhou
	DataRegion *string `json:"DataRegion,omitempty" xml:"DataRegion,omitempty"`
	// example:
	//
	// default
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// example:
	//
	// .*
	FunctionNamePattern *string `json:"FunctionNamePattern,omitempty" xml:"FunctionNamePattern,omitempty"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// f056501ada12c1cc
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 3***
	Tid         *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
	WorkspaceId *int64 `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListDataLakeFunctionRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataLakeFunctionRequest) GoString() string {
	return s.String()
}

func (s *ListDataLakeFunctionRequest) GetCatalogName() *string {
	return s.CatalogName
}

func (s *ListDataLakeFunctionRequest) GetDataRegion() *string {
	return s.DataRegion
}

func (s *ListDataLakeFunctionRequest) GetDbName() *string {
	return s.DbName
}

func (s *ListDataLakeFunctionRequest) GetFunctionNamePattern() *string {
	return s.FunctionNamePattern
}

func (s *ListDataLakeFunctionRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDataLakeFunctionRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDataLakeFunctionRequest) GetTid() *int64 {
	return s.Tid
}

func (s *ListDataLakeFunctionRequest) GetWorkspaceId() *int64 {
	return s.WorkspaceId
}

func (s *ListDataLakeFunctionRequest) SetCatalogName(v string) *ListDataLakeFunctionRequest {
	s.CatalogName = &v
	return s
}

func (s *ListDataLakeFunctionRequest) SetDataRegion(v string) *ListDataLakeFunctionRequest {
	s.DataRegion = &v
	return s
}

func (s *ListDataLakeFunctionRequest) SetDbName(v string) *ListDataLakeFunctionRequest {
	s.DbName = &v
	return s
}

func (s *ListDataLakeFunctionRequest) SetFunctionNamePattern(v string) *ListDataLakeFunctionRequest {
	s.FunctionNamePattern = &v
	return s
}

func (s *ListDataLakeFunctionRequest) SetMaxResults(v int32) *ListDataLakeFunctionRequest {
	s.MaxResults = &v
	return s
}

func (s *ListDataLakeFunctionRequest) SetNextToken(v string) *ListDataLakeFunctionRequest {
	s.NextToken = &v
	return s
}

func (s *ListDataLakeFunctionRequest) SetTid(v int64) *ListDataLakeFunctionRequest {
	s.Tid = &v
	return s
}

func (s *ListDataLakeFunctionRequest) SetWorkspaceId(v int64) *ListDataLakeFunctionRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListDataLakeFunctionRequest) Validate() error {
	return dara.Validate(s)
}
