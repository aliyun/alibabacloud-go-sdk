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
	SetFileName(v string) *ListFileRequest
	GetFileName() *string
	SetMaxResults(v int32) *ListFileRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListFileRequest
	GetNextToken() *string
}

type ListFileRequest struct {
	// The category ID, which is the `CategoryId` returned by the [AddCategory](~~AddCategory~~) operation. To view the category ID, click the ID icon next to the category name on the Unstructured Data tab of the [Data Management](https://bailian.console.alibabacloud.com/#/data-center) page.
	//
	// This parameter is required.
	//
	// example:
	//
	// cate_cdd11b1b79a74e8bbd675c356a91ee3510024405
	CategoryId *string `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	FileName   *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
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
