// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategoryId(v string) *ListFileRequest
	GetCategoryId() *string
	SetFileIds(v []*string) *ListFileRequest
	GetFileIds() []*string
	SetFileName(v string) *ListFileRequest
	GetFileName() *string
	SetMaxResults(v int32) *ListFileRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListFileRequest
	GetNextToken() *string
}

type ListFileRequest struct {
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
	// The list of file IDs to query. A maximum of 20 files can be queried per request.
	FileIds []*string `json:"FileIds,omitempty" xml:"FileIds,omitempty" type:"Repeated"`
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
	// If the value is not set or is less than 1, the default value is 20. If the value is greater than 200, the default value is 200.
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

func (s ListFileRequest) String() string {
	return dara.Prettify(s)
}

func (s ListFileRequest) GoString() string {
	return s.String()
}

func (s *ListFileRequest) GetCategoryId() *string {
	return s.CategoryId
}

func (s *ListFileRequest) GetFileIds() []*string {
	return s.FileIds
}

func (s *ListFileRequest) GetFileName() *string {
	return s.FileName
}

func (s *ListFileRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListFileRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListFileRequest) SetCategoryId(v string) *ListFileRequest {
	s.CategoryId = &v
	return s
}

func (s *ListFileRequest) SetFileIds(v []*string) *ListFileRequest {
	s.FileIds = v
	return s
}

func (s *ListFileRequest) SetFileName(v string) *ListFileRequest {
	s.FileName = &v
	return s
}

func (s *ListFileRequest) SetMaxResults(v int32) *ListFileRequest {
	s.MaxResults = &v
	return s
}

func (s *ListFileRequest) SetNextToken(v string) *ListFileRequest {
	s.NextToken = &v
	return s
}

func (s *ListFileRequest) Validate() error {
	return dara.Validate(s)
}
