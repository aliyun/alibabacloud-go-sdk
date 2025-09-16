// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAdvanceConfigsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataSourceName(v string) *ListAdvanceConfigsRequest
	GetDataSourceName() *string
	SetIndexName(v string) *ListAdvanceConfigsRequest
	GetIndexName() *string
	SetNewMode(v bool) *ListAdvanceConfigsRequest
	GetNewMode() *bool
	SetPageNumber(v string) *ListAdvanceConfigsRequest
	GetPageNumber() *string
	SetPageSize(v string) *ListAdvanceConfigsRequest
	GetPageSize() *string
	SetType(v string) *ListAdvanceConfigsRequest
	GetType() *string
}

type ListAdvanceConfigsRequest struct {
	// The name of the data source.
	//
	// example:
	//
	// ha-cn-pl32rf0****_test_api
	DataSourceName *string `json:"dataSourceName,omitempty" xml:"dataSourceName,omitempty"`
	// The index name.
	//
	// example:
	//
	// test_api
	IndexName *string `json:"indexName,omitempty" xml:"indexName,omitempty"`
	// Specifies whether the OpenSearch Vector Search Edition instance is of the new version.
	//
	// example:
	//
	// true
	NewMode    *bool   `json:"newMode,omitempty" xml:"newMode,omitempty"`
	PageNumber *string `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *string `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The type of advanced configurations that you want to query. Valid values: - online -offline (default)
	//
	// example:
	//
	// online
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListAdvanceConfigsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAdvanceConfigsRequest) GoString() string {
	return s.String()
}

func (s *ListAdvanceConfigsRequest) GetDataSourceName() *string {
	return s.DataSourceName
}

func (s *ListAdvanceConfigsRequest) GetIndexName() *string {
	return s.IndexName
}

func (s *ListAdvanceConfigsRequest) GetNewMode() *bool {
	return s.NewMode
}

func (s *ListAdvanceConfigsRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *ListAdvanceConfigsRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *ListAdvanceConfigsRequest) GetType() *string {
	return s.Type
}

func (s *ListAdvanceConfigsRequest) SetDataSourceName(v string) *ListAdvanceConfigsRequest {
	s.DataSourceName = &v
	return s
}

func (s *ListAdvanceConfigsRequest) SetIndexName(v string) *ListAdvanceConfigsRequest {
	s.IndexName = &v
	return s
}

func (s *ListAdvanceConfigsRequest) SetNewMode(v bool) *ListAdvanceConfigsRequest {
	s.NewMode = &v
	return s
}

func (s *ListAdvanceConfigsRequest) SetPageNumber(v string) *ListAdvanceConfigsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAdvanceConfigsRequest) SetPageSize(v string) *ListAdvanceConfigsRequest {
	s.PageSize = &v
	return s
}

func (s *ListAdvanceConfigsRequest) SetType(v string) *ListAdvanceConfigsRequest {
	s.Type = &v
	return s
}

func (s *ListAdvanceConfigsRequest) Validate() error {
	return dara.Validate(s)
}
