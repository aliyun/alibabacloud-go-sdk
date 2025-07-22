// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDemoCategory interface {
	dara.Model
	String() string
	GoString() string
	SetCategoryCode(v string) *DemoCategory
	GetCategoryCode() *string
	SetCategoryName(v string) *DemoCategory
	GetCategoryName() *string
	SetOrder(v int64) *DemoCategory
	GetOrder() *int64
	SetSubCategories(v []*DemoCategory) *DemoCategory
	GetSubCategories() []*DemoCategory
}

type DemoCategory struct {
	// example:
	//
	// sdk
	CategoryCode *string `json:"CategoryCode,omitempty" xml:"CategoryCode,omitempty"`
	// example:
	//
	// SDK Usage
	CategoryName *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
	// example:
	//
	// 12
	Order         *int64          `json:"Order,omitempty" xml:"Order,omitempty"`
	SubCategories []*DemoCategory `json:"SubCategories,omitempty" xml:"SubCategories,omitempty" type:"Repeated"`
}

func (s DemoCategory) String() string {
	return dara.Prettify(s)
}

func (s DemoCategory) GoString() string {
	return s.String()
}

func (s *DemoCategory) GetCategoryCode() *string {
	return s.CategoryCode
}

func (s *DemoCategory) GetCategoryName() *string {
	return s.CategoryName
}

func (s *DemoCategory) GetOrder() *int64 {
	return s.Order
}

func (s *DemoCategory) GetSubCategories() []*DemoCategory {
	return s.SubCategories
}

func (s *DemoCategory) SetCategoryCode(v string) *DemoCategory {
	s.CategoryCode = &v
	return s
}

func (s *DemoCategory) SetCategoryName(v string) *DemoCategory {
	s.CategoryName = &v
	return s
}

func (s *DemoCategory) SetOrder(v int64) *DemoCategory {
	s.Order = &v
	return s
}

func (s *DemoCategory) SetSubCategories(v []*DemoCategory) *DemoCategory {
	s.SubCategories = v
	return s
}

func (s *DemoCategory) Validate() error {
	return dara.Validate(s)
}
