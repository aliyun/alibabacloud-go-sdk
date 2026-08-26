// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOneMetaOssieModelsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCatalogUuid(v string) *ListOneMetaOssieModelsRequest
	GetCatalogUuid() *string
	SetDatabaseUuid(v string) *ListOneMetaOssieModelsRequest
	GetDatabaseUuid() *string
	SetEnableVectorSearch(v bool) *ListOneMetaOssieModelsRequest
	GetEnableVectorSearch() *bool
	SetMaxResults(v int32) *ListOneMetaOssieModelsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListOneMetaOssieModelsRequest
	GetNextToken() *string
	SetQuery(v string) *ListOneMetaOssieModelsRequest
	GetQuery() *string
	SetTag(v string) *ListOneMetaOssieModelsRequest
	GetTag() *string
}

type ListOneMetaOssieModelsRequest struct {
	// The UUID of the associated folder.
	//
	// example:
	//
	// mc-HZ-OfjcNc2z***
	CatalogUuid *string `json:"CatalogUuid,omitempty" xml:"CatalogUuid,omitempty"`
	// The UUID of the associated database.
	//
	// example:
	//
	// md-HZ-fp9K7r***
	DatabaseUuid *string `json:"DatabaseUuid,omitempty" xml:"DatabaseUuid,omitempty"`
	// Specifies whether to use semantic search.
	//
	// example:
	//
	// true
	EnableVectorSearch *bool `json:"EnableVectorSearch,omitempty" xml:"EnableVectorSearch,omitempty"`
	// The maximum number of records per page.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token for the next query.
	//
	// example:
	//
	// NesLoKLEdIZrKhDT7I2gS****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The search keyword.
	//
	// This parameter is required.
	//
	// example:
	//
	// sale
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// The semantic model tag.
	//
	// example:
	//
	// new_sales
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s ListOneMetaOssieModelsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListOneMetaOssieModelsRequest) GoString() string {
	return s.String()
}

func (s *ListOneMetaOssieModelsRequest) GetCatalogUuid() *string {
	return s.CatalogUuid
}

func (s *ListOneMetaOssieModelsRequest) GetDatabaseUuid() *string {
	return s.DatabaseUuid
}

func (s *ListOneMetaOssieModelsRequest) GetEnableVectorSearch() *bool {
	return s.EnableVectorSearch
}

func (s *ListOneMetaOssieModelsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListOneMetaOssieModelsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListOneMetaOssieModelsRequest) GetQuery() *string {
	return s.Query
}

func (s *ListOneMetaOssieModelsRequest) GetTag() *string {
	return s.Tag
}

func (s *ListOneMetaOssieModelsRequest) SetCatalogUuid(v string) *ListOneMetaOssieModelsRequest {
	s.CatalogUuid = &v
	return s
}

func (s *ListOneMetaOssieModelsRequest) SetDatabaseUuid(v string) *ListOneMetaOssieModelsRequest {
	s.DatabaseUuid = &v
	return s
}

func (s *ListOneMetaOssieModelsRequest) SetEnableVectorSearch(v bool) *ListOneMetaOssieModelsRequest {
	s.EnableVectorSearch = &v
	return s
}

func (s *ListOneMetaOssieModelsRequest) SetMaxResults(v int32) *ListOneMetaOssieModelsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListOneMetaOssieModelsRequest) SetNextToken(v string) *ListOneMetaOssieModelsRequest {
	s.NextToken = &v
	return s
}

func (s *ListOneMetaOssieModelsRequest) SetQuery(v string) *ListOneMetaOssieModelsRequest {
	s.Query = &v
	return s
}

func (s *ListOneMetaOssieModelsRequest) SetTag(v string) *ListOneMetaOssieModelsRequest {
	s.Tag = &v
	return s
}

func (s *ListOneMetaOssieModelsRequest) Validate() error {
	return dara.Validate(s)
}
