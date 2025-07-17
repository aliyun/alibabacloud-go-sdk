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
	SetCreateUser(v string) *ListMetaCollectionsRequest
	GetCreateUser() *string
	SetDescription(v string) *ListMetaCollectionsRequest
	GetDescription() *string
	SetName(v string) *ListMetaCollectionsRequest
	GetName() *string
	SetOrder(v string) *ListMetaCollectionsRequest
	GetOrder() *string
	SetPageNumber(v int32) *ListMetaCollectionsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListMetaCollectionsRequest
	GetPageSize() *int32
	SetParentId(v string) *ListMetaCollectionsRequest
	GetParentId() *string
	SetSortBy(v string) *ListMetaCollectionsRequest
	GetSortBy() *string
	SetType(v string) *ListMetaCollectionsRequest
	GetType() *string
}

type ListMetaCollectionsRequest struct {
	// example:
	//
	// 12345
	Administrator *string `json:"Administrator,omitempty" xml:"Administrator,omitempty"`
	// example:
	//
	// 123456
	CreateUser  *string `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// Asc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the collection of an ancestor node.
	//
	// example:
	//
	// category.123
	ParentId *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	// example:
	//
	// Name
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The type of the collection. Valid values:
	//
	// 	- Category
	//
	// 	- Album
	//
	// 	- AlbumCategory
	//
	// This parameter is required.
	//
	// example:
	//
	// Category
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
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

func (s *ListMetaCollectionsRequest) GetCreateUser() *string {
	return s.CreateUser
}

func (s *ListMetaCollectionsRequest) GetDescription() *string {
	return s.Description
}

func (s *ListMetaCollectionsRequest) GetName() *string {
	return s.Name
}

func (s *ListMetaCollectionsRequest) GetOrder() *string {
	return s.Order
}

func (s *ListMetaCollectionsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListMetaCollectionsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMetaCollectionsRequest) GetParentId() *string {
	return s.ParentId
}

func (s *ListMetaCollectionsRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListMetaCollectionsRequest) GetType() *string {
	return s.Type
}

func (s *ListMetaCollectionsRequest) SetAdministrator(v string) *ListMetaCollectionsRequest {
	s.Administrator = &v
	return s
}

func (s *ListMetaCollectionsRequest) SetCreateUser(v string) *ListMetaCollectionsRequest {
	s.CreateUser = &v
	return s
}

func (s *ListMetaCollectionsRequest) SetDescription(v string) *ListMetaCollectionsRequest {
	s.Description = &v
	return s
}

func (s *ListMetaCollectionsRequest) SetName(v string) *ListMetaCollectionsRequest {
	s.Name = &v
	return s
}

func (s *ListMetaCollectionsRequest) SetOrder(v string) *ListMetaCollectionsRequest {
	s.Order = &v
	return s
}

func (s *ListMetaCollectionsRequest) SetPageNumber(v int32) *ListMetaCollectionsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListMetaCollectionsRequest) SetPageSize(v int32) *ListMetaCollectionsRequest {
	s.PageSize = &v
	return s
}

func (s *ListMetaCollectionsRequest) SetParentId(v string) *ListMetaCollectionsRequest {
	s.ParentId = &v
	return s
}

func (s *ListMetaCollectionsRequest) SetSortBy(v string) *ListMetaCollectionsRequest {
	s.SortBy = &v
	return s
}

func (s *ListMetaCollectionsRequest) SetType(v string) *ListMetaCollectionsRequest {
	s.Type = &v
	return s
}

func (s *ListMetaCollectionsRequest) Validate() error {
	return dara.Validate(s)
}
