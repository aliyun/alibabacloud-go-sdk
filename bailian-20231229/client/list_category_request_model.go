// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCategoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategoryName(v string) *ListCategoryRequest
	GetCategoryName() *string
	SetCategoryType(v string) *ListCategoryRequest
	GetCategoryType() *string
	SetConnectorId(v string) *ListCategoryRequest
	GetConnectorId() *string
	SetMaxResults(v int32) *ListCategoryRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListCategoryRequest
	GetNextToken() *string
	SetParentCategoryId(v string) *ListCategoryRequest
	GetParentCategoryId() *string
}

type ListCategoryRequest struct {
	// Filters the results to include only the category with this exact name. If this parameter is omitted, no filtering is applied.
	//
	// example:
	//
	// 产品清单
	CategoryName *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
	// The type of category to query. Valid value:
	//
	// - `UNSTRUCTURED`: A category for unstructured data.
	//
	// <props="china">
	//
	// > This API does not support querying structured data tables.
	//
	//
	//
	// <props="intl">
	//
	// > This API does not support querying structured data tables.
	//
	// This parameter is required.
	//
	// example:
	//
	// UNSTRUCTURED
	CategoryType *string `json:"CategoryType,omitempty" xml:"CategoryType,omitempty"`
	// The ID of the connector.
	//
	// example:
	//
	// file_conn_xxxxx
	ConnectorId *string `json:"ConnectorId,omitempty" xml:"ConnectorId,omitempty"`
	// The maximum number of categories to return per page. The valid range is 1 to 200.
	//
	// Default value: 20. If this parameter is not specified or is set to a value less than 1, the default value is used. If a value greater than 200 is specified, the maximum value of 200 is used.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token. To retrieve the next page of results, pass the `NextToken` value from the previous response.
	//
	// example:
	//
	// AAAAAdH70eOCSCKtacdomNzak4U=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the parent category.
	//
	// example:
	//
	// cate_cdd11b1b79a74e8bbd675c356a91ee3xxxxxxxx
	ParentCategoryId *string `json:"ParentCategoryId,omitempty" xml:"ParentCategoryId,omitempty"`
}

func (s ListCategoryRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCategoryRequest) GoString() string {
	return s.String()
}

func (s *ListCategoryRequest) GetCategoryName() *string {
	return s.CategoryName
}

func (s *ListCategoryRequest) GetCategoryType() *string {
	return s.CategoryType
}

func (s *ListCategoryRequest) GetConnectorId() *string {
	return s.ConnectorId
}

func (s *ListCategoryRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListCategoryRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListCategoryRequest) GetParentCategoryId() *string {
	return s.ParentCategoryId
}

func (s *ListCategoryRequest) SetCategoryName(v string) *ListCategoryRequest {
	s.CategoryName = &v
	return s
}

func (s *ListCategoryRequest) SetCategoryType(v string) *ListCategoryRequest {
	s.CategoryType = &v
	return s
}

func (s *ListCategoryRequest) SetConnectorId(v string) *ListCategoryRequest {
	s.ConnectorId = &v
	return s
}

func (s *ListCategoryRequest) SetMaxResults(v int32) *ListCategoryRequest {
	s.MaxResults = &v
	return s
}

func (s *ListCategoryRequest) SetNextToken(v string) *ListCategoryRequest {
	s.NextToken = &v
	return s
}

func (s *ListCategoryRequest) SetParentCategoryId(v string) *ListCategoryRequest {
	s.ParentCategoryId = &v
	return s
}

func (s *ListCategoryRequest) Validate() error {
	return dara.Validate(s)
}
