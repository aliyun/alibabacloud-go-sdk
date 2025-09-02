// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMetaCollectionEntitiesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCollectionQualifiedName(v string) *ListMetaCollectionEntitiesRequest
	GetCollectionQualifiedName() *string
	SetEntityType(v string) *ListMetaCollectionEntitiesRequest
	GetEntityType() *string
	SetKeyword(v string) *ListMetaCollectionEntitiesRequest
	GetKeyword() *string
	SetNextToken(v string) *ListMetaCollectionEntitiesRequest
	GetNextToken() *string
	SetPageSize(v int32) *ListMetaCollectionEntitiesRequest
	GetPageSize() *int32
}

type ListMetaCollectionEntitiesRequest struct {
	// The unique identifier of the collection.
	//
	// This parameter is required.
	//
	// example:
	//
	// album.12345
	CollectionQualifiedName *string `json:"CollectionQualifiedName,omitempty" xml:"CollectionQualifiedName,omitempty"`
	// The type of the entities.
	//
	// For example, if this parameter is set to maxcompute-table, the entity is a MaxCompute table.
	//
	// example:
	//
	// maxcompute-table
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// The search keyword.
	//
	// example:
	//
	// test
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// 12222
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListMetaCollectionEntitiesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMetaCollectionEntitiesRequest) GoString() string {
	return s.String()
}

func (s *ListMetaCollectionEntitiesRequest) GetCollectionQualifiedName() *string {
	return s.CollectionQualifiedName
}

func (s *ListMetaCollectionEntitiesRequest) GetEntityType() *string {
	return s.EntityType
}

func (s *ListMetaCollectionEntitiesRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListMetaCollectionEntitiesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListMetaCollectionEntitiesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMetaCollectionEntitiesRequest) SetCollectionQualifiedName(v string) *ListMetaCollectionEntitiesRequest {
	s.CollectionQualifiedName = &v
	return s
}

func (s *ListMetaCollectionEntitiesRequest) SetEntityType(v string) *ListMetaCollectionEntitiesRequest {
	s.EntityType = &v
	return s
}

func (s *ListMetaCollectionEntitiesRequest) SetKeyword(v string) *ListMetaCollectionEntitiesRequest {
	s.Keyword = &v
	return s
}

func (s *ListMetaCollectionEntitiesRequest) SetNextToken(v string) *ListMetaCollectionEntitiesRequest {
	s.NextToken = &v
	return s
}

func (s *ListMetaCollectionEntitiesRequest) SetPageSize(v int32) *ListMetaCollectionEntitiesRequest {
	s.PageSize = &v
	return s
}

func (s *ListMetaCollectionEntitiesRequest) Validate() error {
	return dara.Validate(s)
}
