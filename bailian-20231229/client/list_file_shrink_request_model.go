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
	// The category ID, which is the `CategoryId` returned by the [AddCategory](~~AddCategory~~) operation. To view the category ID, click the ID icon next to the category name on the Unstructured Data tab of the [Data Management](https://bailian.console.alibabacloud.com/#/data-center) page.
	//
	// This parameter is required.
	//
	// example:
	//
	// cate_cdd11b1b79a74e8bbd675c356a91ee3510024405
	CategoryId    *string `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	FileIdsShrink *string `json:"FileIds,omitempty" xml:"FileIds,omitempty"`
	FileName      *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
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
