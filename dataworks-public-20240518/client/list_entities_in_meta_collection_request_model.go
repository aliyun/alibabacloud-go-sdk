// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEntitiesInMetaCollectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEntityDescription(v string) *ListEntitiesInMetaCollectionRequest
	GetEntityDescription() *string
	SetEntityName(v string) *ListEntitiesInMetaCollectionRequest
	GetEntityName() *string
	SetEntityType(v string) *ListEntitiesInMetaCollectionRequest
	GetEntityType() *string
	SetId(v string) *ListEntitiesInMetaCollectionRequest
	GetId() *string
	SetOrder(v string) *ListEntitiesInMetaCollectionRequest
	GetOrder() *string
	SetPageNumber(v int32) *ListEntitiesInMetaCollectionRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListEntitiesInMetaCollectionRequest
	GetPageSize() *int32
	SetSortBy(v string) *ListEntitiesInMetaCollectionRequest
	GetSortBy() *string
}

type ListEntitiesInMetaCollectionRequest struct {
	// The description specified when the entity was added to the collection. Supports fuzzy matching. Valid only for the album type.
	//
	// example:
	//
	// test
	EntityDescription *string `json:"EntityDescription,omitempty" xml:"EntityDescription,omitempty"`
	// The entity name. Supports fuzzy matching.
	//
	// example:
	//
	// test1
	EntityName *string `json:"EntityName,omitempty" xml:"EntityName,omitempty"`
	// The entity type.
	//
	// example:
	//
	// dlf-table
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// The collection ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// category.123
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The sort order. Valid values:
	//
	// 	- Asc (default): ascending order.
	//
	// 	- Desc
	//
	// example:
	//
	// Asc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of records per page. Default value: 10. Maximum value: 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The sort field. Valid values:
	//
	// 	- Name (default)
	//
	// 	- CreateTime
	//
	// example:
	//
	// Name
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
}

func (s ListEntitiesInMetaCollectionRequest) String() string {
	return dara.Prettify(s)
}

func (s ListEntitiesInMetaCollectionRequest) GoString() string {
	return s.String()
}

func (s *ListEntitiesInMetaCollectionRequest) GetEntityDescription() *string {
	return s.EntityDescription
}

func (s *ListEntitiesInMetaCollectionRequest) GetEntityName() *string {
	return s.EntityName
}

func (s *ListEntitiesInMetaCollectionRequest) GetEntityType() *string {
	return s.EntityType
}

func (s *ListEntitiesInMetaCollectionRequest) GetId() *string {
	return s.Id
}

func (s *ListEntitiesInMetaCollectionRequest) GetOrder() *string {
	return s.Order
}

func (s *ListEntitiesInMetaCollectionRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListEntitiesInMetaCollectionRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListEntitiesInMetaCollectionRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListEntitiesInMetaCollectionRequest) SetEntityDescription(v string) *ListEntitiesInMetaCollectionRequest {
	s.EntityDescription = &v
	return s
}

func (s *ListEntitiesInMetaCollectionRequest) SetEntityName(v string) *ListEntitiesInMetaCollectionRequest {
	s.EntityName = &v
	return s
}

func (s *ListEntitiesInMetaCollectionRequest) SetEntityType(v string) *ListEntitiesInMetaCollectionRequest {
	s.EntityType = &v
	return s
}

func (s *ListEntitiesInMetaCollectionRequest) SetId(v string) *ListEntitiesInMetaCollectionRequest {
	s.Id = &v
	return s
}

func (s *ListEntitiesInMetaCollectionRequest) SetOrder(v string) *ListEntitiesInMetaCollectionRequest {
	s.Order = &v
	return s
}

func (s *ListEntitiesInMetaCollectionRequest) SetPageNumber(v int32) *ListEntitiesInMetaCollectionRequest {
	s.PageNumber = &v
	return s
}

func (s *ListEntitiesInMetaCollectionRequest) SetPageSize(v int32) *ListEntitiesInMetaCollectionRequest {
	s.PageSize = &v
	return s
}

func (s *ListEntitiesInMetaCollectionRequest) SetSortBy(v string) *ListEntitiesInMetaCollectionRequest {
	s.SortBy = &v
	return s
}

func (s *ListEntitiesInMetaCollectionRequest) Validate() error {
	return dara.Validate(s)
}
