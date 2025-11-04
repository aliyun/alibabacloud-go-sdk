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
	// The information about the category.
	Category *AddCategoryResponseBodyCategory `json:"Category,omitempty" xml:"Category,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// ****63E8B7C7-4812-46AD-0FA56029AC86****
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
	// The ID of the created category.
	//
	// example:
	//
	// 45
	CateId *int64 `json:"CateId,omitempty" xml:"CateId,omitempty"`
	// The category name.
	CateName *string `json:"CateName,omitempty" xml:"CateName,omitempty"`
	// The level of the category. A value of **0*	- indicates a level-1 category, a value of **1*	- indicates a level-2 category, and a value of **2*	- indicates a level-3 category.
	//
	// example:
	//
	// 0
	Level *int64 `json:"Level,omitempty" xml:"Level,omitempty"`
	// The ID of the parent category. By default, if ParentId is left empty or less than 1, -1 is returned, which indicates that the created category is the root directory.
	//
	// example:
	//
	// -1
	ParentId *int64 `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	// The type of the category. Valid values:
	//
	// 	- **default**: audio, video, and image files. This is the default value.
	//
	// 	- **material**: short video materials.
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
