// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFileShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategoryId(v string) *ListFileShrinkRequest
	GetCategoryId() *string
	SetFileIdsShrink(v string) *ListFileShrinkRequest
	GetFileIdsShrink() *string
	SetFileName(v string) *ListFileShrinkRequest
	GetFileName() *string
	SetMaxResults(v int32) *ListFileShrinkRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListFileShrinkRequest
	GetNextToken() *string
}

type ListFileShrinkRequest struct {
	// <props="china">
	//
	// The category ID, which is the `CategoryId` returned by the **AddCategory*	- operation. You can also obtain it on the [Application Data](https://bailian.console.aliyun.com/?tab=app#/data-center) - Files tab by clicking the ID icon next to the category name.
	//
	//
	//
	// <props="intl">
	//
	// The category ID, which is the `CategoryId` returned by the **AddCategory*	- operation. You can also obtain it on the [Application Data](https://modelstudio.console.alibabacloud.com/?tab=app#/data-center) - Files tab by clicking the ID icon next to the category name.
	//
	// .
	//
	// This parameter is required.
	//
	// example:
	//
	// cate_cdd11b1b79a74e8bbd675c356a91ee35xxxxxxxx
	CategoryId *string `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// The list of file IDs to query. A maximum of 20 files can be queried at a time.
	FileIdsShrink *string `json:"FileIds,omitempty" xml:"FileIds,omitempty"`
	// The file name (without extension). Only exact match is supported. Fuzzy search is not supported.
	//
	// example:
	//
	// product-overview
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The number of entries per page for paging. Valid values: 1 to 200.
	//
	// Default value:
	//
	// If no value is set or the value is less than 1, the default value is 20. If the value is set to greater than 200, the default value is 200.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token. Set this to the NextToken value returned by the previous API call.
	//
	// example:
	//
	// AAAAAdH70eOCSCKtacdomNzak4U=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListFileShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListFileShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListFileShrinkRequest) GetCategoryId() *string {
	return s.CategoryId
}

func (s *ListFileShrinkRequest) GetFileIdsShrink() *string {
	return s.FileIdsShrink
}

func (s *ListFileShrinkRequest) GetFileName() *string {
	return s.FileName
}

func (s *ListFileShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListFileShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListFileShrinkRequest) SetCategoryId(v string) *ListFileShrinkRequest {
	s.CategoryId = &v
	return s
}

func (s *ListFileShrinkRequest) SetFileIdsShrink(v string) *ListFileShrinkRequest {
	s.FileIdsShrink = &v
	return s
}

func (s *ListFileShrinkRequest) SetFileName(v string) *ListFileShrinkRequest {
	s.FileName = &v
	return s
}

func (s *ListFileShrinkRequest) SetMaxResults(v int32) *ListFileShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListFileShrinkRequest) SetNextToken(v string) *ListFileShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListFileShrinkRequest) Validate() error {
	return dara.Validate(s)
}
