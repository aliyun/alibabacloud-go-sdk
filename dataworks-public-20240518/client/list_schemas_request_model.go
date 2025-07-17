// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSchemasRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *ListSchemasRequest
	GetComment() *string
	SetName(v string) *ListSchemasRequest
	GetName() *string
	SetOrder(v string) *ListSchemasRequest
	GetOrder() *string
	SetPageNumber(v int32) *ListSchemasRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListSchemasRequest
	GetPageSize() *int32
	SetParentMetaEntityId(v string) *ListSchemasRequest
	GetParentMetaEntityId() *string
	SetSortBy(v string) *ListSchemasRequest
	GetSortBy() *string
	SetTypes(v []*string) *ListSchemasRequest
	GetTypes() []*string
}

type ListSchemasRequest struct {
	// The comment. Fuzzy match is supported.
	//
	// example:
	//
	// test comment
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The name. Fuzzy match is supported.
	//
	// example:
	//
	// abc
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The order in which schemas are sorted. Default value: Asc. Valid values:
	//
	// 	- Asc: ascending order
	//
	// 	- Desc: descending order
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
	// The number of entries per page. Default value: 10. Maximum value: 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The parent entity ID. For more information, see [Concepts related to metadata entities](https://help.aliyun.com/document_detail/2880092.html). For the Hologres metadata crawler type, you can call the ListDatabases operation to query the settings of the `ParentMetaEntityId` parameter.
	//
	// Configure the `ParentMetaEntityId` parameter in the `${EntityType}:${Instance ID or escaped URL}:${Catalog identifier}:${Database name}` format. If a level does not exist, leave the level empty.
	//
	// >  If you want to query the information about a MaxCompute schema, specify an empty string at the Instance ID level as a placeholder and a MaxCompute project name at the Database name level. Make sure that the schema feature is enabled for the MaxCompute project.
	//
	// This parameter is required.
	//
	// example:
	//
	// maxcompute-project:123456XXX::test_project
	//
	// holo-database:h-abc123xxx::test_db
	ParentMetaEntityId *string `json:"ParentMetaEntityId,omitempty" xml:"ParentMetaEntityId,omitempty"`
	// The field used for sorting. Default value: CreateTime. Valid values:
	//
	// 	- CreateTime
	//
	// 	- ModifyTime
	//
	// 	- Name
	//
	// 	- Type
	//
	// example:
	//
	// CreateTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The types. Exact match is supported. If this parameter is left empty, all types are queried.
	Types []*string `json:"Types,omitempty" xml:"Types,omitempty" type:"Repeated"`
}

func (s ListSchemasRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSchemasRequest) GoString() string {
	return s.String()
}

func (s *ListSchemasRequest) GetComment() *string {
	return s.Comment
}

func (s *ListSchemasRequest) GetName() *string {
	return s.Name
}

func (s *ListSchemasRequest) GetOrder() *string {
	return s.Order
}

func (s *ListSchemasRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListSchemasRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSchemasRequest) GetParentMetaEntityId() *string {
	return s.ParentMetaEntityId
}

func (s *ListSchemasRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListSchemasRequest) GetTypes() []*string {
	return s.Types
}

func (s *ListSchemasRequest) SetComment(v string) *ListSchemasRequest {
	s.Comment = &v
	return s
}

func (s *ListSchemasRequest) SetName(v string) *ListSchemasRequest {
	s.Name = &v
	return s
}

func (s *ListSchemasRequest) SetOrder(v string) *ListSchemasRequest {
	s.Order = &v
	return s
}

func (s *ListSchemasRequest) SetPageNumber(v int32) *ListSchemasRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSchemasRequest) SetPageSize(v int32) *ListSchemasRequest {
	s.PageSize = &v
	return s
}

func (s *ListSchemasRequest) SetParentMetaEntityId(v string) *ListSchemasRequest {
	s.ParentMetaEntityId = &v
	return s
}

func (s *ListSchemasRequest) SetSortBy(v string) *ListSchemasRequest {
	s.SortBy = &v
	return s
}

func (s *ListSchemasRequest) SetTypes(v []*string) *ListSchemasRequest {
	s.Types = v
	return s
}

func (s *ListSchemasRequest) Validate() error {
	return dara.Validate(s)
}
