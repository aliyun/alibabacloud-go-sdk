// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMetaCollectionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdministrator(v string) *ListMetaCollectionsRequest
	GetAdministrator() *string
	SetCollectionType(v string) *ListMetaCollectionsRequest
	GetCollectionType() *string
	SetCreator(v string) *ListMetaCollectionsRequest
	GetCreator() *string
	SetFollower(v string) *ListMetaCollectionsRequest
	GetFollower() *string
	SetKeyword(v string) *ListMetaCollectionsRequest
	GetKeyword() *string
	SetNextToken(v string) *ListMetaCollectionsRequest
	GetNextToken() *string
	SetOrderBy(v string) *ListMetaCollectionsRequest
	GetOrderBy() *string
	SetPageSize(v int32) *ListMetaCollectionsRequest
	GetPageSize() *int32
	SetParentQualifiedName(v string) *ListMetaCollectionsRequest
	GetParentQualifiedName() *string
}

type ListMetaCollectionsRequest struct {
	// The ID of the collection administrator.
	//
	// example:
	//
	// 1200759642363000
	Administrator *string `json:"Administrator,omitempty" xml:"Administrator,omitempty"`
	// - ALBUM: data album
	//
	// - ALBUM_CATEGORY: category in a data album
	//
	// This parameter is required.
	//
	// example:
	//
	// ALBUM
	CollectionType *string `json:"CollectionType,omitempty" xml:"CollectionType,omitempty"`
	// The ID of the collection creator.
	//
	// example:
	//
	// 1200759642363000
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// The ID of the collection follower.
	//
	// example:
	//
	// 1200759642363000
	Follower *string `json:"Follower,omitempty" xml:"Follower,omitempty"`
	// The keyword.
	//
	// example:
	//
	// test
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// 12345
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The name of the sorting field.
	//
	// example:
	//
	// test
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// The number of entries per page. Default value: 10. Maximum value: 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The unique identifier of the parent collection.
	//
	// example:
	//
	// Album.1234
	ParentQualifiedName *string `json:"ParentQualifiedName,omitempty" xml:"ParentQualifiedName,omitempty"`
}

func (s ListMetaCollectionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMetaCollectionsRequest) GoString() string {
	return s.String()
}

func (s *ListMetaCollectionsRequest) GetAdministrator() *string {
	return s.Administrator
}

func (s *ListMetaCollectionsRequest) GetCollectionType() *string {
	return s.CollectionType
}

func (s *ListMetaCollectionsRequest) GetCreator() *string {
	return s.Creator
}

func (s *ListMetaCollectionsRequest) GetFollower() *string {
	return s.Follower
}

func (s *ListMetaCollectionsRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListMetaCollectionsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListMetaCollectionsRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListMetaCollectionsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMetaCollectionsRequest) GetParentQualifiedName() *string {
	return s.ParentQualifiedName
}

func (s *ListMetaCollectionsRequest) SetAdministrator(v string) *ListMetaCollectionsRequest {
	s.Administrator = &v
	return s
}

func (s *ListMetaCollectionsRequest) SetCollectionType(v string) *ListMetaCollectionsRequest {
	s.CollectionType = &v
	return s
}

func (s *ListMetaCollectionsRequest) SetCreator(v string) *ListMetaCollectionsRequest {
	s.Creator = &v
	return s
}

func (s *ListMetaCollectionsRequest) SetFollower(v string) *ListMetaCollectionsRequest {
	s.Follower = &v
	return s
}

func (s *ListMetaCollectionsRequest) SetKeyword(v string) *ListMetaCollectionsRequest {
	s.Keyword = &v
	return s
}

func (s *ListMetaCollectionsRequest) SetNextToken(v string) *ListMetaCollectionsRequest {
	s.NextToken = &v
	return s
}

func (s *ListMetaCollectionsRequest) SetOrderBy(v string) *ListMetaCollectionsRequest {
	s.OrderBy = &v
	return s
}

func (s *ListMetaCollectionsRequest) SetPageSize(v int32) *ListMetaCollectionsRequest {
	s.PageSize = &v
	return s
}

func (s *ListMetaCollectionsRequest) SetParentQualifiedName(v string) *ListMetaCollectionsRequest {
	s.ParentQualifiedName = &v
	return s
}

func (s *ListMetaCollectionsRequest) Validate() error {
	return dara.Validate(s)
}
