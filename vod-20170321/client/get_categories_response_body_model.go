// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCategoriesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCategory(v *GetCategoriesResponseBodyCategory) *GetCategoriesResponseBody
	GetCategory() *GetCategoriesResponseBodyCategory
	SetRequestId(v string) *GetCategoriesResponseBody
	GetRequestId() *string
	SetSubCategories(v *GetCategoriesResponseBodySubCategories) *GetCategoriesResponseBody
	GetSubCategories() *GetCategoriesResponseBodySubCategories
	SetSubTotal(v int64) *GetCategoriesResponseBody
	GetSubTotal() *int64
}

type GetCategoriesResponseBody struct {
	// The information about the category.
	Category *GetCategoriesResponseBodyCategory `json:"Category,omitempty" xml:"Category,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 25818875-5F78-4AF6-D7393642CA58****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The details of the subcategory.
	SubCategories *GetCategoriesResponseBodySubCategories `json:"SubCategories,omitempty" xml:"SubCategories,omitempty" type:"Struct"`
	// The total number of subcategories.
	//
	// example:
	//
	// 3795
	SubTotal *int64 `json:"SubTotal,omitempty" xml:"SubTotal,omitempty"`
}

func (s GetCategoriesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCategoriesResponseBody) GoString() string {
	return s.String()
}

func (s *GetCategoriesResponseBody) GetCategory() *GetCategoriesResponseBodyCategory {
	return s.Category
}

func (s *GetCategoriesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCategoriesResponseBody) GetSubCategories() *GetCategoriesResponseBodySubCategories {
	return s.SubCategories
}

func (s *GetCategoriesResponseBody) GetSubTotal() *int64 {
	return s.SubTotal
}

func (s *GetCategoriesResponseBody) SetCategory(v *GetCategoriesResponseBodyCategory) *GetCategoriesResponseBody {
	s.Category = v
	return s
}

func (s *GetCategoriesResponseBody) SetRequestId(v string) *GetCategoriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCategoriesResponseBody) SetSubCategories(v *GetCategoriesResponseBodySubCategories) *GetCategoriesResponseBody {
	s.SubCategories = v
	return s
}

func (s *GetCategoriesResponseBody) SetSubTotal(v int64) *GetCategoriesResponseBody {
	s.SubTotal = &v
	return s
}

func (s *GetCategoriesResponseBody) Validate() error {
	if s.Category != nil {
		if err := s.Category.Validate(); err != nil {
			return err
		}
	}
	if s.SubCategories != nil {
		if err := s.SubCategories.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetCategoriesResponseBodyCategory struct {
	// The ID of the category.
	//
	// example:
	//
	// 100
	CateId *int64 `json:"CateId,omitempty" xml:"CateId,omitempty"`
	// The name of the category.
	//
	// example:
	//
	// film
	CateName *string `json:"CateName,omitempty" xml:"CateName,omitempty"`
	// The level of the category. Valid values:
	//
	// 	- **0**: level 1 category
	//
	// 	- **1**: level 2 category
	//
	// 	- **2**: level 3 category
	//
	// example:
	//
	// 0
	Level *int64 `json:"Level,omitempty" xml:"Level,omitempty"`
	// The ID of the parent category.
	//
	// example:
	//
	// 100012****
	ParentId *int64 `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	// The type of the category. Valid values:
	//
	// 	- **default**: audio, video, and image files
	//
	// 	- **material**: short video materials
	//
	// example:
	//
	// default
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetCategoriesResponseBodyCategory) String() string {
	return dara.Prettify(s)
}

func (s GetCategoriesResponseBodyCategory) GoString() string {
	return s.String()
}

func (s *GetCategoriesResponseBodyCategory) GetCateId() *int64 {
	return s.CateId
}

func (s *GetCategoriesResponseBodyCategory) GetCateName() *string {
	return s.CateName
}

func (s *GetCategoriesResponseBodyCategory) GetLevel() *int64 {
	return s.Level
}

func (s *GetCategoriesResponseBodyCategory) GetParentId() *int64 {
	return s.ParentId
}

func (s *GetCategoriesResponseBodyCategory) GetType() *string {
	return s.Type
}

func (s *GetCategoriesResponseBodyCategory) SetCateId(v int64) *GetCategoriesResponseBodyCategory {
	s.CateId = &v
	return s
}

func (s *GetCategoriesResponseBodyCategory) SetCateName(v string) *GetCategoriesResponseBodyCategory {
	s.CateName = &v
	return s
}

func (s *GetCategoriesResponseBodyCategory) SetLevel(v int64) *GetCategoriesResponseBodyCategory {
	s.Level = &v
	return s
}

func (s *GetCategoriesResponseBodyCategory) SetParentId(v int64) *GetCategoriesResponseBodyCategory {
	s.ParentId = &v
	return s
}

func (s *GetCategoriesResponseBodyCategory) SetType(v string) *GetCategoriesResponseBodyCategory {
	s.Type = &v
	return s
}

func (s *GetCategoriesResponseBodyCategory) Validate() error {
	return dara.Validate(s)
}

type GetCategoriesResponseBodySubCategories struct {
	Category []*GetCategoriesResponseBodySubCategoriesCategory `json:"Category,omitempty" xml:"Category,omitempty" type:"Repeated"`
}

func (s GetCategoriesResponseBodySubCategories) String() string {
	return dara.Prettify(s)
}

func (s GetCategoriesResponseBodySubCategories) GoString() string {
	return s.String()
}

func (s *GetCategoriesResponseBodySubCategories) GetCategory() []*GetCategoriesResponseBodySubCategoriesCategory {
	return s.Category
}

func (s *GetCategoriesResponseBodySubCategories) SetCategory(v []*GetCategoriesResponseBodySubCategoriesCategory) *GetCategoriesResponseBodySubCategories {
	s.Category = v
	return s
}

func (s *GetCategoriesResponseBodySubCategories) Validate() error {
	if s.Category != nil {
		for _, item := range s.Category {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetCategoriesResponseBodySubCategoriesCategory struct {
	// The ID of the category.
	//
	// example:
	//
	// 100
	CateId *int64 `json:"CateId,omitempty" xml:"CateId,omitempty"`
	// The name of the category.
	//
	// example:
	//
	// film
	CateName *string `json:"CateName,omitempty" xml:"CateName,omitempty"`
	// The level of the category. Valid values:
	//
	// 	- **0**: level 1 category
	//
	// 	- **1**: level 2 category
	//
	// 	- **2**: level 3 category
	//
	// example:
	//
	// 1
	Level *int64 `json:"Level,omitempty" xml:"Level,omitempty"`
	// The ID of the parent category.
	//
	// example:
	//
	// 10020****
	ParentId *int64 `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	// The total number of subcategories.
	//
	// example:
	//
	// 1
	SubTotal *int64 `json:"SubTotal,omitempty" xml:"SubTotal,omitempty"`
	// The type of the subcategory. Valid values:
	//
	// 	- **default**: audio, video, and image files
	//
	// 	- **material**: short video materials
	//
	// example:
	//
	// default
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetCategoriesResponseBodySubCategoriesCategory) String() string {
	return dara.Prettify(s)
}

func (s GetCategoriesResponseBodySubCategoriesCategory) GoString() string {
	return s.String()
}

func (s *GetCategoriesResponseBodySubCategoriesCategory) GetCateId() *int64 {
	return s.CateId
}

func (s *GetCategoriesResponseBodySubCategoriesCategory) GetCateName() *string {
	return s.CateName
}

func (s *GetCategoriesResponseBodySubCategoriesCategory) GetLevel() *int64 {
	return s.Level
}

func (s *GetCategoriesResponseBodySubCategoriesCategory) GetParentId() *int64 {
	return s.ParentId
}

func (s *GetCategoriesResponseBodySubCategoriesCategory) GetSubTotal() *int64 {
	return s.SubTotal
}

func (s *GetCategoriesResponseBodySubCategoriesCategory) GetType() *string {
	return s.Type
}

func (s *GetCategoriesResponseBodySubCategoriesCategory) SetCateId(v int64) *GetCategoriesResponseBodySubCategoriesCategory {
	s.CateId = &v
	return s
}

func (s *GetCategoriesResponseBodySubCategoriesCategory) SetCateName(v string) *GetCategoriesResponseBodySubCategoriesCategory {
	s.CateName = &v
	return s
}

func (s *GetCategoriesResponseBodySubCategoriesCategory) SetLevel(v int64) *GetCategoriesResponseBodySubCategoriesCategory {
	s.Level = &v
	return s
}

func (s *GetCategoriesResponseBodySubCategoriesCategory) SetParentId(v int64) *GetCategoriesResponseBodySubCategoriesCategory {
	s.ParentId = &v
	return s
}

func (s *GetCategoriesResponseBodySubCategoriesCategory) SetSubTotal(v int64) *GetCategoriesResponseBodySubCategoriesCategory {
	s.SubTotal = &v
	return s
}

func (s *GetCategoriesResponseBodySubCategoriesCategory) SetType(v string) *GetCategoriesResponseBodySubCategoriesCategory {
	s.Type = &v
	return s
}

func (s *GetCategoriesResponseBodySubCategoriesCategory) Validate() error {
	return dara.Validate(s)
}
