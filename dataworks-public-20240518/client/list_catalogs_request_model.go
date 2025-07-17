// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCatalogsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *ListCatalogsRequest
	GetComment() *string
	SetName(v string) *ListCatalogsRequest
	GetName() *string
	SetOrder(v string) *ListCatalogsRequest
	GetOrder() *string
	SetPageNumber(v int32) *ListCatalogsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListCatalogsRequest
	GetPageSize() *int32
	SetParentMetaEntityId(v string) *ListCatalogsRequest
	GetParentMetaEntityId() *string
	SetSortBy(v string) *ListCatalogsRequest
	GetSortBy() *string
	SetTypes(v []*string) *ListCatalogsRequest
	GetTypes() []*string
}

type ListCatalogsRequest struct {
	// example:
	//
	// this is a comment
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// example:
	//
	// abc
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
	// The parent entity ID. For more information, see [Concepts related to metadata entities](https://help.aliyun.com/document_detail/2880092.html).
	//
	// Only DLF and StarRocks data sources support this parameter.
	//
	// 	- For DLF data sources, you can call this API operation to query all catalogs. In this case, you must set the `ParentMetaEntityId` parameter to `dlf`.
	//
	// 	- For StarRocks data sources, you can call this API operation to query the catalogs in a specific instance. In this case, you can configure the `ParentMetaEntityId` parameter in the `starrocks:(instance_id|encoded_jdbc_url)` format.
	//
	// > \\
	//
	// `instance_id`: the ID of an instance. If the related data source is added to DataWorks in Alibaba Cloud instance mode, you must configure this parameter.\\
	//
	// `encoded_jdbc_url`: the JDBC connection string that is URL-encoded. If the related data source is added to DataWorks in connection string mode, you must configure this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// dlf
	//
	// starrocks:c-abc123xxx
	ParentMetaEntityId *string `json:"ParentMetaEntityId,omitempty" xml:"ParentMetaEntityId,omitempty"`
	// example:
	//
	// CreateTime
	SortBy *string   `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	Types  []*string `json:"Types,omitempty" xml:"Types,omitempty" type:"Repeated"`
}

func (s ListCatalogsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCatalogsRequest) GoString() string {
	return s.String()
}

func (s *ListCatalogsRequest) GetComment() *string {
	return s.Comment
}

func (s *ListCatalogsRequest) GetName() *string {
	return s.Name
}

func (s *ListCatalogsRequest) GetOrder() *string {
	return s.Order
}

func (s *ListCatalogsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCatalogsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCatalogsRequest) GetParentMetaEntityId() *string {
	return s.ParentMetaEntityId
}

func (s *ListCatalogsRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListCatalogsRequest) GetTypes() []*string {
	return s.Types
}

func (s *ListCatalogsRequest) SetComment(v string) *ListCatalogsRequest {
	s.Comment = &v
	return s
}

func (s *ListCatalogsRequest) SetName(v string) *ListCatalogsRequest {
	s.Name = &v
	return s
}

func (s *ListCatalogsRequest) SetOrder(v string) *ListCatalogsRequest {
	s.Order = &v
	return s
}

func (s *ListCatalogsRequest) SetPageNumber(v int32) *ListCatalogsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCatalogsRequest) SetPageSize(v int32) *ListCatalogsRequest {
	s.PageSize = &v
	return s
}

func (s *ListCatalogsRequest) SetParentMetaEntityId(v string) *ListCatalogsRequest {
	s.ParentMetaEntityId = &v
	return s
}

func (s *ListCatalogsRequest) SetSortBy(v string) *ListCatalogsRequest {
	s.SortBy = &v
	return s
}

func (s *ListCatalogsRequest) SetTypes(v []*string) *ListCatalogsRequest {
	s.Types = v
	return s
}

func (s *ListCatalogsRequest) Validate() error {
	return dara.Validate(s)
}
