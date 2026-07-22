// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAssetCategoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCategory(v *CreateAssetCategoryResponseBodyCategory) *CreateAssetCategoryResponseBody
	GetCategory() *CreateAssetCategoryResponseBodyCategory
	SetRequestId(v string) *CreateAssetCategoryResponseBody
	GetRequestId() *string
}

type CreateAssetCategoryResponseBody struct {
	Category *CreateAssetCategoryResponseBodyCategory `json:"Category,omitempty" xml:"Category,omitempty" type:"Struct"`
	// example:
	//
	// ****63E8B7C7-4812-46AD-0FA56029AC86****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAssetCategoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAssetCategoryResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAssetCategoryResponseBody) GetCategory() *CreateAssetCategoryResponseBodyCategory {
	return s.Category
}

func (s *CreateAssetCategoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAssetCategoryResponseBody) SetCategory(v *CreateAssetCategoryResponseBodyCategory) *CreateAssetCategoryResponseBody {
	s.Category = v
	return s
}

func (s *CreateAssetCategoryResponseBody) SetRequestId(v string) *CreateAssetCategoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAssetCategoryResponseBody) Validate() error {
	if s.Category != nil {
		if err := s.Category.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateAssetCategoryResponseBodyCategory struct {
	// example:
	//
	// 45
	CategoryId *int64 `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// example:
	//
	// see below
	CategoryName *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
	// example:
	//
	// 0
	Level *int64 `json:"Level,omitempty" xml:"Level,omitempty"`
	// example:
	//
	// -1
	ParentId *int64 `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
}

func (s CreateAssetCategoryResponseBodyCategory) String() string {
	return dara.Prettify(s)
}

func (s CreateAssetCategoryResponseBodyCategory) GoString() string {
	return s.String()
}

func (s *CreateAssetCategoryResponseBodyCategory) GetCategoryId() *int64 {
	return s.CategoryId
}

func (s *CreateAssetCategoryResponseBodyCategory) GetCategoryName() *string {
	return s.CategoryName
}

func (s *CreateAssetCategoryResponseBodyCategory) GetLevel() *int64 {
	return s.Level
}

func (s *CreateAssetCategoryResponseBodyCategory) GetParentId() *int64 {
	return s.ParentId
}

func (s *CreateAssetCategoryResponseBodyCategory) SetCategoryId(v int64) *CreateAssetCategoryResponseBodyCategory {
	s.CategoryId = &v
	return s
}

func (s *CreateAssetCategoryResponseBodyCategory) SetCategoryName(v string) *CreateAssetCategoryResponseBodyCategory {
	s.CategoryName = &v
	return s
}

func (s *CreateAssetCategoryResponseBodyCategory) SetLevel(v int64) *CreateAssetCategoryResponseBodyCategory {
	s.Level = &v
	return s
}

func (s *CreateAssetCategoryResponseBodyCategory) SetParentId(v int64) *CreateAssetCategoryResponseBodyCategory {
	s.ParentId = &v
	return s
}

func (s *CreateAssetCategoryResponseBodyCategory) Validate() error {
	return dara.Validate(s)
}
