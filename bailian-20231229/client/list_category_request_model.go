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
	SetMaxResults(v int32) *ListCategoryRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListCategoryRequest
	GetNextToken() *string
	SetParentCategoryId(v string) *ListCategoryRequest
	GetParentCategoryId() *string
}

type ListCategoryRequest struct {
	CategoryName *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// UNSTRUCTURED
	CategoryType *string `json:"CategoryType,omitempty" xml:"CategoryType,omitempty"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// AAAAAdH70eOCSCKtacdomNzak4U=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// cate_cdd11b1b79a74e8bbd675c356a91ee3XXXXXXXX
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
