// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePageFaceVerifyDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int64) *DescribePageFaceVerifyDataRequest
	GetCurrentPage() *int64
	SetEndDate(v string) *DescribePageFaceVerifyDataRequest
	GetEndDate() *string
	SetPageSize(v int64) *DescribePageFaceVerifyDataRequest
	GetPageSize() *int64
	SetProductCode(v string) *DescribePageFaceVerifyDataRequest
	GetProductCode() *string
	SetSceneId(v int64) *DescribePageFaceVerifyDataRequest
	GetSceneId() *int64
	SetStartDate(v string) *DescribePageFaceVerifyDataRequest
	GetStartDate() *string
}

type DescribePageFaceVerifyDataRequest struct {
	// example:
	//
	// 1
	CurrentPage *int64 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// 2023-04-30
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// ID_PLUS
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// example:
	//
	// 36**01
	SceneId *int64 `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// example:
	//
	// 2023-04-10
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s DescribePageFaceVerifyDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePageFaceVerifyDataRequest) GoString() string {
	return s.String()
}

func (s *DescribePageFaceVerifyDataRequest) GetCurrentPage() *int64 {
	return s.CurrentPage
}

func (s *DescribePageFaceVerifyDataRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *DescribePageFaceVerifyDataRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribePageFaceVerifyDataRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *DescribePageFaceVerifyDataRequest) GetSceneId() *int64 {
	return s.SceneId
}

func (s *DescribePageFaceVerifyDataRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *DescribePageFaceVerifyDataRequest) SetCurrentPage(v int64) *DescribePageFaceVerifyDataRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribePageFaceVerifyDataRequest) SetEndDate(v string) *DescribePageFaceVerifyDataRequest {
	s.EndDate = &v
	return s
}

func (s *DescribePageFaceVerifyDataRequest) SetPageSize(v int64) *DescribePageFaceVerifyDataRequest {
	s.PageSize = &v
	return s
}

func (s *DescribePageFaceVerifyDataRequest) SetProductCode(v string) *DescribePageFaceVerifyDataRequest {
	s.ProductCode = &v
	return s
}

func (s *DescribePageFaceVerifyDataRequest) SetSceneId(v int64) *DescribePageFaceVerifyDataRequest {
	s.SceneId = &v
	return s
}

func (s *DescribePageFaceVerifyDataRequest) SetStartDate(v string) *DescribePageFaceVerifyDataRequest {
	s.StartDate = &v
	return s
}

func (s *DescribePageFaceVerifyDataRequest) Validate() error {
	return dara.Validate(s)
}
