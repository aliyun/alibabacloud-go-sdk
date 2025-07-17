// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSchemasShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *ListSchemasShrinkRequest
	GetComment() *string
	SetName(v string) *ListSchemasShrinkRequest
	GetName() *string
	SetOrder(v string) *ListSchemasShrinkRequest
	GetOrder() *string
	SetPageNumber(v int32) *ListSchemasShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListSchemasShrinkRequest
	GetPageSize() *int32
	SetParentMetaEntityId(v string) *ListSchemasShrinkRequest
	GetParentMetaEntityId() *string
	SetSortBy(v string) *ListSchemasShrinkRequest
	GetSortBy() *string
	SetTypesShrink(v string) *ListSchemasShrinkRequest
	GetTypesShrink() *string
}

type ListSchemasShrinkRequest struct {
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
	TypesShrink *string `json:"Types,omitempty" xml:"Types,omitempty"`
}

func (s ListSchemasShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSchemasShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListSchemasShrinkRequest) GetComment() *string {
	return s.Comment
}

func (s *ListSchemasShrinkRequest) GetName() *string {
	return s.Name
}

func (s *ListSchemasShrinkRequest) GetOrder() *string {
	return s.Order
}

func (s *ListSchemasShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListSchemasShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSchemasShrinkRequest) GetParentMetaEntityId() *string {
	return s.ParentMetaEntityId
}

func (s *ListSchemasShrinkRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListSchemasShrinkRequest) GetTypesShrink() *string {
	return s.TypesShrink
}

func (s *ListSchemasShrinkRequest) SetComment(v string) *ListSchemasShrinkRequest {
	s.Comment = &v
	return s
}

func (s *ListSchemasShrinkRequest) SetName(v string) *ListSchemasShrinkRequest {
	s.Name = &v
	return s
}

func (s *ListSchemasShrinkRequest) SetOrder(v string) *ListSchemasShrinkRequest {
	s.Order = &v
	return s
}

func (s *ListSchemasShrinkRequest) SetPageNumber(v int32) *ListSchemasShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSchemasShrinkRequest) SetPageSize(v int32) *ListSchemasShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListSchemasShrinkRequest) SetParentMetaEntityId(v string) *ListSchemasShrinkRequest {
	s.ParentMetaEntityId = &v
	return s
}

func (s *ListSchemasShrinkRequest) SetSortBy(v string) *ListSchemasShrinkRequest {
	s.SortBy = &v
	return s
}

func (s *ListSchemasShrinkRequest) SetTypesShrink(v string) *ListSchemasShrinkRequest {
	s.TypesShrink = &v
	return s
}

func (s *ListSchemasShrinkRequest) Validate() error {
	return dara.Validate(s)
}
