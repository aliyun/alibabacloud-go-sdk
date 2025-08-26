// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataLakeFunctionNameRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCatalogName(v string) *ListDataLakeFunctionNameRequest
	GetCatalogName() *string
	SetDbName(v string) *ListDataLakeFunctionNameRequest
	GetDbName() *string
	SetFunctionNamePattern(v string) *ListDataLakeFunctionNameRequest
	GetFunctionNamePattern() *string
	SetMaxResults(v int32) *ListDataLakeFunctionNameRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListDataLakeFunctionNameRequest
	GetNextToken() *string
	SetTid(v int64) *ListDataLakeFunctionNameRequest
	GetTid() *int64
	SetWorkspaceId(v int64) *ListDataLakeFunctionNameRequest
	GetWorkspaceId() *int64
}

type ListDataLakeFunctionNameRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// hive
	CatalogName *string `json:"CatalogName,omitempty" xml:"CatalogName,omitempty"`
	// This parameter is required.
	//
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
	// f056501ada12****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 3****
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
	// example:
	//
	// 12****
	WorkspaceId *int64 `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListDataLakeFunctionNameRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataLakeFunctionNameRequest) GoString() string {
	return s.String()
}

func (s *ListDataLakeFunctionNameRequest) GetCatalogName() *string {
	return s.CatalogName
}

func (s *ListDataLakeFunctionNameRequest) GetDbName() *string {
	return s.DbName
}

func (s *ListDataLakeFunctionNameRequest) GetFunctionNamePattern() *string {
	return s.FunctionNamePattern
}

func (s *ListDataLakeFunctionNameRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDataLakeFunctionNameRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDataLakeFunctionNameRequest) GetTid() *int64 {
	return s.Tid
}

func (s *ListDataLakeFunctionNameRequest) GetWorkspaceId() *int64 {
	return s.WorkspaceId
}

func (s *ListDataLakeFunctionNameRequest) SetCatalogName(v string) *ListDataLakeFunctionNameRequest {
	s.CatalogName = &v
	return s
}

func (s *ListDataLakeFunctionNameRequest) SetDbName(v string) *ListDataLakeFunctionNameRequest {
	s.DbName = &v
	return s
}

func (s *ListDataLakeFunctionNameRequest) SetFunctionNamePattern(v string) *ListDataLakeFunctionNameRequest {
	s.FunctionNamePattern = &v
	return s
}

func (s *ListDataLakeFunctionNameRequest) SetMaxResults(v int32) *ListDataLakeFunctionNameRequest {
	s.MaxResults = &v
	return s
}

func (s *ListDataLakeFunctionNameRequest) SetNextToken(v string) *ListDataLakeFunctionNameRequest {
	s.NextToken = &v
	return s
}

func (s *ListDataLakeFunctionNameRequest) SetTid(v int64) *ListDataLakeFunctionNameRequest {
	s.Tid = &v
	return s
}

func (s *ListDataLakeFunctionNameRequest) SetWorkspaceId(v int64) *ListDataLakeFunctionNameRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListDataLakeFunctionNameRequest) Validate() error {
	return dara.Validate(s)
}
