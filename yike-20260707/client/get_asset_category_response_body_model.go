// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAssetCategoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCategory(v *GetAssetCategoryResponseBodyCategory) *GetAssetCategoryResponseBody
	GetCategory() *GetAssetCategoryResponseBodyCategory
	SetRequestId(v string) *GetAssetCategoryResponseBody
	GetRequestId() *string
	SetSubCategories(v []*GetAssetCategoryResponseBodySubCategories) *GetAssetCategoryResponseBody
	GetSubCategories() []*GetAssetCategoryResponseBodySubCategories
	SetSubTotal(v int64) *GetAssetCategoryResponseBody
	GetSubTotal() *int64
}

type GetAssetCategoryResponseBody struct {
	// The category details.
	Category *GetAssetCategoryResponseBodyCategory `json:"Category,omitempty" xml:"Category,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// ****63E8B7C7-4812-46AD-0FA56029AC86****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of subcategories.
	SubCategories []*GetAssetCategoryResponseBodySubCategories `json:"SubCategories,omitempty" xml:"SubCategories,omitempty" type:"Repeated"`
	// The number of subcategories.
	//
	// example:
	//
	// 5
	SubTotal *int64 `json:"SubTotal,omitempty" xml:"SubTotal,omitempty"`
}

func (s GetAssetCategoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAssetCategoryResponseBody) GoString() string {
	return s.String()
}

func (s *GetAssetCategoryResponseBody) GetCategory() *GetAssetCategoryResponseBodyCategory {
	return s.Category
}

func (s *GetAssetCategoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAssetCategoryResponseBody) GetSubCategories() []*GetAssetCategoryResponseBodySubCategories {
	return s.SubCategories
}

func (s *GetAssetCategoryResponseBody) GetSubTotal() *int64 {
	return s.SubTotal
}

func (s *GetAssetCategoryResponseBody) SetCategory(v *GetAssetCategoryResponseBodyCategory) *GetAssetCategoryResponseBody {
	s.Category = v
	return s
}

func (s *GetAssetCategoryResponseBody) SetRequestId(v string) *GetAssetCategoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAssetCategoryResponseBody) SetSubCategories(v []*GetAssetCategoryResponseBodySubCategories) *GetAssetCategoryResponseBody {
	s.SubCategories = v
	return s
}

func (s *GetAssetCategoryResponseBody) SetSubTotal(v int64) *GetAssetCategoryResponseBody {
	s.SubTotal = &v
	return s
}

func (s *GetAssetCategoryResponseBody) Validate() error {
	if s.Category != nil {
		if err := s.Category.Validate(); err != nil {
			return err
		}
	}
	if s.SubCategories != nil {
		for _, item := range s.SubCategories {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetAssetCategoryResponseBodyCategory struct {
	// The category ID.
	//
	// example:
	//
	// 50
	CategoryId *int64 `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// The category name.
	//
	// example:
	//
	// scenery
	CategoryName *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
	// The category level. Valid values:
	//
	// - **0**: level-1 category.
	//
	// - **1**: level-2 category.
	//
	// - **2**: level-3 category.
	//
	// example:
	//
	// 1
	Level *int64 `json:"Level,omitempty" xml:"Level,omitempty"`
	// The parent category ID.
	//
	// example:
	//
	// 10
	ParentId *int64 `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
}

func (s GetAssetCategoryResponseBodyCategory) String() string {
	return dara.Prettify(s)
}

func (s GetAssetCategoryResponseBodyCategory) GoString() string {
	return s.String()
}

func (s *GetAssetCategoryResponseBodyCategory) GetCategoryId() *int64 {
	return s.CategoryId
}

func (s *GetAssetCategoryResponseBodyCategory) GetCategoryName() *string {
	return s.CategoryName
}

func (s *GetAssetCategoryResponseBodyCategory) GetLevel() *int64 {
	return s.Level
}

func (s *GetAssetCategoryResponseBodyCategory) GetParentId() *int64 {
	return s.ParentId
}

func (s *GetAssetCategoryResponseBodyCategory) SetCategoryId(v int64) *GetAssetCategoryResponseBodyCategory {
	s.CategoryId = &v
	return s
}

func (s *GetAssetCategoryResponseBodyCategory) SetCategoryName(v string) *GetAssetCategoryResponseBodyCategory {
	s.CategoryName = &v
	return s
}

func (s *GetAssetCategoryResponseBodyCategory) SetLevel(v int64) *GetAssetCategoryResponseBodyCategory {
	s.Level = &v
	return s
}

func (s *GetAssetCategoryResponseBodyCategory) SetParentId(v int64) *GetAssetCategoryResponseBodyCategory {
	s.ParentId = &v
	return s
}

func (s *GetAssetCategoryResponseBodyCategory) Validate() error {
	return dara.Validate(s)
}

type GetAssetCategoryResponseBodySubCategories struct {
	// The category ID.
	//
	// example:
	//
	// 55
	CategoryId *int64 `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// The category name.
	//
	// example:
	//
	// sea
	CategoryName *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
	// The category level.
	//
	// example:
	//
	// 2
	Level *int64 `json:"Level,omitempty" xml:"Level,omitempty"`
	// The parent category ID.
	//
	// example:
	//
	// 50
	ParentId *int64 `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	// The total number of subcategories.
	//
	// example:
	//
	// 0
	SubTotal *int64 `json:"SubTotal,omitempty" xml:"SubTotal,omitempty"`
}

func (s GetAssetCategoryResponseBodySubCategories) String() string {
	return dara.Prettify(s)
}

func (s GetAssetCategoryResponseBodySubCategories) GoString() string {
	return s.String()
}

func (s *GetAssetCategoryResponseBodySubCategories) GetCategoryId() *int64 {
	return s.CategoryId
}

func (s *GetAssetCategoryResponseBodySubCategories) GetCategoryName() *string {
	return s.CategoryName
}

func (s *GetAssetCategoryResponseBodySubCategories) GetLevel() *int64 {
	return s.Level
}

func (s *GetAssetCategoryResponseBodySubCategories) GetParentId() *int64 {
	return s.ParentId
}

func (s *GetAssetCategoryResponseBodySubCategories) GetSubTotal() *int64 {
	return s.SubTotal
}

func (s *GetAssetCategoryResponseBodySubCategories) SetCategoryId(v int64) *GetAssetCategoryResponseBodySubCategories {
	s.CategoryId = &v
	return s
}

func (s *GetAssetCategoryResponseBodySubCategories) SetCategoryName(v string) *GetAssetCategoryResponseBodySubCategories {
	s.CategoryName = &v
	return s
}

func (s *GetAssetCategoryResponseBodySubCategories) SetLevel(v int64) *GetAssetCategoryResponseBodySubCategories {
	s.Level = &v
	return s
}

func (s *GetAssetCategoryResponseBodySubCategories) SetParentId(v int64) *GetAssetCategoryResponseBodySubCategories {
	s.ParentId = &v
	return s
}

func (s *GetAssetCategoryResponseBodySubCategories) SetSubTotal(v int64) *GetAssetCategoryResponseBodySubCategories {
	s.SubTotal = &v
	return s
}

func (s *GetAssetCategoryResponseBodySubCategories) Validate() error {
	return dara.Validate(s)
}
