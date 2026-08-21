// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddCategoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCategory(v *AddCategoryResponseBodyCategory) *AddCategoryResponseBody
	GetCategory() *AddCategoryResponseBodyCategory
	SetRequestId(v string) *AddCategoryResponseBody
	GetRequestId() *string
}

type AddCategoryResponseBody struct {
	// The media asset category information.
	Category *AddCategoryResponseBodyCategory `json:"Category,omitempty" xml:"Category,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 25818875-5F78-4AF6-D7393642CA58****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddCategoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddCategoryResponseBody) GoString() string {
	return s.String()
}

func (s *AddCategoryResponseBody) GetCategory() *AddCategoryResponseBodyCategory {
	return s.Category
}

func (s *AddCategoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddCategoryResponseBody) SetCategory(v *AddCategoryResponseBodyCategory) *AddCategoryResponseBody {
	s.Category = v
	return s
}

func (s *AddCategoryResponseBody) SetRequestId(v string) *AddCategoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddCategoryResponseBody) Validate() error {
	if s.Category != nil {
		if err := s.Category.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddCategoryResponseBodyCategory struct {
	// The category ID. This ID can be used as a request parameter for the [UpdateCategory](~~UpdateCategory~~), [DeleteCategory](~~DeleteCategory~~), and [GetCategories](~~GetCategories~~) operations.
	//
	// example:
	//
	// 10020
	CateId *int64 `json:"CateId,omitempty" xml:"CateId,omitempty"`
	// The category name.
	//
	// example:
	//
	// Comedy
	CateName *string `json:"CateName,omitempty" xml:"CateName,omitempty"`
	// The category level. Valid values:
	//
	// - **0**: level-0 category.
	//
	// - **1**: level-1 category.
	//
	// - **2**: level-2 category.
	//
	// example:
	//
	// 1
	Level *int64 `json:"Level,omitempty" xml:"Level,omitempty"`
	// The parent category ID.
	//
	// example:
	//
	// 100012
	ParentId *int64 `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	// The category type. Valid values:
	//
	// - **default**: audio/video/image category.
	//
	// - **material**: short video material category.
	//
	// example:
	//
	// default
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s AddCategoryResponseBodyCategory) String() string {
	return dara.Prettify(s)
}

func (s AddCategoryResponseBodyCategory) GoString() string {
	return s.String()
}

func (s *AddCategoryResponseBodyCategory) GetCateId() *int64 {
	return s.CateId
}

func (s *AddCategoryResponseBodyCategory) GetCateName() *string {
	return s.CateName
}

func (s *AddCategoryResponseBodyCategory) GetLevel() *int64 {
	return s.Level
}

func (s *AddCategoryResponseBodyCategory) GetParentId() *int64 {
	return s.ParentId
}

func (s *AddCategoryResponseBodyCategory) GetType() *string {
	return s.Type
}

func (s *AddCategoryResponseBodyCategory) SetCateId(v int64) *AddCategoryResponseBodyCategory {
	s.CateId = &v
	return s
}

func (s *AddCategoryResponseBodyCategory) SetCateName(v string) *AddCategoryResponseBodyCategory {
	s.CateName = &v
	return s
}

func (s *AddCategoryResponseBodyCategory) SetLevel(v int64) *AddCategoryResponseBodyCategory {
	s.Level = &v
	return s
}

func (s *AddCategoryResponseBodyCategory) SetParentId(v int64) *AddCategoryResponseBodyCategory {
	s.ParentId = &v
	return s
}

func (s *AddCategoryResponseBodyCategory) SetType(v string) *AddCategoryResponseBodyCategory {
	s.Type = &v
	return s
}

func (s *AddCategoryResponseBodyCategory) Validate() error {
	return dara.Validate(s)
}
