// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCategoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCategories(v []*ListCategoryResponseBodyCategories) *ListCategoryResponseBody
	GetCategories() []*ListCategoryResponseBodyCategories
	SetRequestId(v string) *ListCategoryResponseBody
	GetRequestId() *string
}

type ListCategoryResponseBody struct {
	Categories []*ListCategoryResponseBodyCategories `json:"Categories,omitempty" xml:"Categories,omitempty" type:"Repeated"`
	// example:
	//
	// 9C5F8186-2D22-433E-9545-606D344F30B5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListCategoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCategoryResponseBody) GoString() string {
	return s.String()
}

func (s *ListCategoryResponseBody) GetCategories() []*ListCategoryResponseBodyCategories {
	return s.Categories
}

func (s *ListCategoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCategoryResponseBody) SetCategories(v []*ListCategoryResponseBodyCategories) *ListCategoryResponseBody {
	s.Categories = v
	return s
}

func (s *ListCategoryResponseBody) SetRequestId(v string) *ListCategoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCategoryResponseBody) Validate() error {
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

type ListCategoryResponseBodyCategories struct {
	BizCode *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
	// example:
	//
	// 231001028593
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

func (s ListCategoryResponseBodyCategories) String() string {
	return dara.Prettify(s)
}

func (s ListCategoryResponseBodyCategories) GoString() string {
	return s.String()
}

func (s *ListCategoryResponseBodyCategories) GetBizCode() *string {
	return s.BizCode
}

func (s *ListCategoryResponseBodyCategories) GetCategoryId() *int64 {
	return s.CategoryId
}

func (s *ListCategoryResponseBodyCategories) GetName() *string {
	return s.Name
}

func (s *ListCategoryResponseBodyCategories) GetParentCategoryId() *int64 {
	return s.ParentCategoryId
}

func (s *ListCategoryResponseBodyCategories) GetStatus() *int32 {
	return s.Status
}

func (s *ListCategoryResponseBodyCategories) SetBizCode(v string) *ListCategoryResponseBodyCategories {
	s.BizCode = &v
	return s
}

func (s *ListCategoryResponseBodyCategories) SetCategoryId(v int64) *ListCategoryResponseBodyCategories {
	s.CategoryId = &v
	return s
}

func (s *ListCategoryResponseBodyCategories) SetName(v string) *ListCategoryResponseBodyCategories {
	s.Name = &v
	return s
}

func (s *ListCategoryResponseBodyCategories) SetParentCategoryId(v int64) *ListCategoryResponseBodyCategories {
	s.ParentCategoryId = &v
	return s
}

func (s *ListCategoryResponseBodyCategories) SetStatus(v int32) *ListCategoryResponseBodyCategories {
	s.Status = &v
	return s
}

func (s *ListCategoryResponseBodyCategories) Validate() error {
	return dara.Validate(s)
}
