// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAssetCategoriesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCategories(v []*ListAssetCategoriesResponseBodyCategories) *ListAssetCategoriesResponseBody
	GetCategories() []*ListAssetCategoriesResponseBodyCategories
	SetRequestId(v string) *ListAssetCategoriesResponseBody
	GetRequestId() *string
	SetTotal(v int64) *ListAssetCategoriesResponseBody
	GetTotal() *int64
}

type ListAssetCategoriesResponseBody struct {
	// The list of categories on the current page.
	Categories []*ListAssetCategoriesResponseBodyCategories `json:"Categories,omitempty" xml:"Categories,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// ****63E8B7C7-4812-46AD-0FA56029AC86****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of categories.
	//
	// example:
	//
	// 50
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListAssetCategoriesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAssetCategoriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAssetCategoriesResponseBody) GetCategories() []*ListAssetCategoriesResponseBodyCategories {
	return s.Categories
}

func (s *ListAssetCategoriesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAssetCategoriesResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *ListAssetCategoriesResponseBody) SetCategories(v []*ListAssetCategoriesResponseBodyCategories) *ListAssetCategoriesResponseBody {
	s.Categories = v
	return s
}

func (s *ListAssetCategoriesResponseBody) SetRequestId(v string) *ListAssetCategoriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAssetCategoriesResponseBody) SetTotal(v int64) *ListAssetCategoriesResponseBody {
	s.Total = &v
	return s
}

func (s *ListAssetCategoriesResponseBody) Validate() error {
	if s.Categories != nil {
		for _, item := range s.Categories {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAssetCategoriesResponseBodyCategories struct {
	// The category ID.
	//
	// example:
	//
	// 45
	CategoryId *int64 `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// The category name.
	//
	// example:
	//
	// scenery
	CategoryName *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
	// The category level. A level-1 category has a value of 0, a level-2 category has a value of 1, and a level-3 category has a value of 2.
	//
	// example:
	//
	// 1
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// The parent category ID.
	//
	// example:
	//
	// 10
	ParentId *int64 `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
}

func (s ListAssetCategoriesResponseBodyCategories) String() string {
	return dara.Prettify(s)
}

func (s ListAssetCategoriesResponseBodyCategories) GoString() string {
	return s.String()
}

func (s *ListAssetCategoriesResponseBodyCategories) GetCategoryId() *int64 {
	return s.CategoryId
}

func (s *ListAssetCategoriesResponseBodyCategories) GetCategoryName() *string {
	return s.CategoryName
}

func (s *ListAssetCategoriesResponseBodyCategories) GetLevel() *string {
	return s.Level
}

func (s *ListAssetCategoriesResponseBodyCategories) GetParentId() *int64 {
	return s.ParentId
}

func (s *ListAssetCategoriesResponseBodyCategories) SetCategoryId(v int64) *ListAssetCategoriesResponseBodyCategories {
	s.CategoryId = &v
	return s
}

func (s *ListAssetCategoriesResponseBodyCategories) SetCategoryName(v string) *ListAssetCategoriesResponseBodyCategories {
	s.CategoryName = &v
	return s
}

func (s *ListAssetCategoriesResponseBodyCategories) SetLevel(v string) *ListAssetCategoriesResponseBodyCategories {
	s.Level = &v
	return s
}

func (s *ListAssetCategoriesResponseBodyCategories) SetParentId(v int64) *ListAssetCategoriesResponseBodyCategories {
	s.ParentId = &v
	return s
}

func (s *ListAssetCategoriesResponseBodyCategories) Validate() error {
	return dara.Validate(s)
}
