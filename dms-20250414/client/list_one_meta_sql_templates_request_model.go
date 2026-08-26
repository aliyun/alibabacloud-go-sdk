// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOneMetaSqlTemplatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCatalogUuid(v string) *ListOneMetaSqlTemplatesRequest
	GetCatalogUuid() *string
	SetDatabaseUuid(v string) *ListOneMetaSqlTemplatesRequest
	GetDatabaseUuid() *string
	SetEnableVectorSearch(v bool) *ListOneMetaSqlTemplatesRequest
	GetEnableVectorSearch() *bool
	SetMaxResults(v int32) *ListOneMetaSqlTemplatesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListOneMetaSqlTemplatesRequest
	GetNextToken() *string
	SetQuery(v string) *ListOneMetaSqlTemplatesRequest
	GetQuery() *string
	SetTag(v string) *ListOneMetaSqlTemplatesRequest
	GetTag() *string
	SetUuids(v string) *ListOneMetaSqlTemplatesRequest
	GetUuids() *string
}

type ListOneMetaSqlTemplatesRequest struct {
	// The UUID of the associated catalog.
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
	// The maximum number of entries to return in this response.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token for the next query.
	//
	// example:
	//
	// f056501ada12****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The search keyword.
	//
	// This parameter is required.
	//
	// example:
	//
	// sale
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// The SQL template tag.
	//
	// example:
	//
	// new_sales
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// The UUIDs of knowledge instances. Separate multiple UUIDs with commas (,).
	//
	// example:
	//
	// 86c5c290052147c***,56c5c2900dasqw***
	Uuids *string `json:"Uuids,omitempty" xml:"Uuids,omitempty"`
}

func (s ListOneMetaSqlTemplatesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListOneMetaSqlTemplatesRequest) GoString() string {
	return s.String()
}

func (s *ListOneMetaSqlTemplatesRequest) GetCatalogUuid() *string {
	return s.CatalogUuid
}

func (s *ListOneMetaSqlTemplatesRequest) GetDatabaseUuid() *string {
	return s.DatabaseUuid
}

func (s *ListOneMetaSqlTemplatesRequest) GetEnableVectorSearch() *bool {
	return s.EnableVectorSearch
}

func (s *ListOneMetaSqlTemplatesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListOneMetaSqlTemplatesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListOneMetaSqlTemplatesRequest) GetQuery() *string {
	return s.Query
}

func (s *ListOneMetaSqlTemplatesRequest) GetTag() *string {
	return s.Tag
}

func (s *ListOneMetaSqlTemplatesRequest) GetUuids() *string {
	return s.Uuids
}

func (s *ListOneMetaSqlTemplatesRequest) SetCatalogUuid(v string) *ListOneMetaSqlTemplatesRequest {
	s.CatalogUuid = &v
	return s
}

func (s *ListOneMetaSqlTemplatesRequest) SetDatabaseUuid(v string) *ListOneMetaSqlTemplatesRequest {
	s.DatabaseUuid = &v
	return s
}

func (s *ListOneMetaSqlTemplatesRequest) SetEnableVectorSearch(v bool) *ListOneMetaSqlTemplatesRequest {
	s.EnableVectorSearch = &v
	return s
}

func (s *ListOneMetaSqlTemplatesRequest) SetMaxResults(v int32) *ListOneMetaSqlTemplatesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListOneMetaSqlTemplatesRequest) SetNextToken(v string) *ListOneMetaSqlTemplatesRequest {
	s.NextToken = &v
	return s
}

func (s *ListOneMetaSqlTemplatesRequest) SetQuery(v string) *ListOneMetaSqlTemplatesRequest {
	s.Query = &v
	return s
}

func (s *ListOneMetaSqlTemplatesRequest) SetTag(v string) *ListOneMetaSqlTemplatesRequest {
	s.Tag = &v
	return s
}

func (s *ListOneMetaSqlTemplatesRequest) SetUuids(v string) *ListOneMetaSqlTemplatesRequest {
	s.Uuids = &v
	return s
}

func (s *ListOneMetaSqlTemplatesRequest) Validate() error {
	return dara.Validate(s)
}
