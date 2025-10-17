// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCategoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCategory(v *DescribeCategoryResponseBodyCategory) *DescribeCategoryResponseBody
	GetCategory() *DescribeCategoryResponseBodyCategory
	SetRequestId(v string) *DescribeCategoryResponseBody
	GetRequestId() *string
}

type DescribeCategoryResponseBody struct {
	Category *DescribeCategoryResponseBodyCategory `json:"Category,omitempty" xml:"Category,omitempty" type:"Struct"`
	// example:
	//
	// 2B0304FD-3804-5C06-9A83-77F5523664AF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCategoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCategoryResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCategoryResponseBody) GetCategory() *DescribeCategoryResponseBodyCategory {
	return s.Category
}

func (s *DescribeCategoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCategoryResponseBody) SetCategory(v *DescribeCategoryResponseBodyCategory) *DescribeCategoryResponseBody {
	s.Category = v
	return s
}

func (s *DescribeCategoryResponseBody) SetRequestId(v string) *DescribeCategoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCategoryResponseBody) Validate() error {
	if s.Category != nil {
		if err := s.Category.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCategoryResponseBodyCategory struct {
	BizCode *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
	// example:
	//
	// 30000049006
	CategoryId *int64  `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// -1
	ParentCategoryId *int64 `json:"ParentCategoryId,omitempty" xml:"ParentCategoryId,omitempty"`
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeCategoryResponseBodyCategory) String() string {
	return dara.Prettify(s)
}

func (s DescribeCategoryResponseBodyCategory) GoString() string {
	return s.String()
}

func (s *DescribeCategoryResponseBodyCategory) GetBizCode() *string {
	return s.BizCode
}

func (s *DescribeCategoryResponseBodyCategory) GetCategoryId() *int64 {
	return s.CategoryId
}

func (s *DescribeCategoryResponseBodyCategory) GetName() *string {
	return s.Name
}

func (s *DescribeCategoryResponseBodyCategory) GetParentCategoryId() *int64 {
	return s.ParentCategoryId
}

func (s *DescribeCategoryResponseBodyCategory) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeCategoryResponseBodyCategory) SetBizCode(v string) *DescribeCategoryResponseBodyCategory {
	s.BizCode = &v
	return s
}

func (s *DescribeCategoryResponseBodyCategory) SetCategoryId(v int64) *DescribeCategoryResponseBodyCategory {
	s.CategoryId = &v
	return s
}

func (s *DescribeCategoryResponseBodyCategory) SetName(v string) *DescribeCategoryResponseBodyCategory {
	s.Name = &v
	return s
}

func (s *DescribeCategoryResponseBodyCategory) SetParentCategoryId(v int64) *DescribeCategoryResponseBodyCategory {
	s.ParentCategoryId = &v
	return s
}

func (s *DescribeCategoryResponseBodyCategory) SetStatus(v int32) *DescribeCategoryResponseBodyCategory {
	s.Status = &v
	return s
}

func (s *DescribeCategoryResponseBodyCategory) Validate() error {
	return dara.Validate(s)
}
