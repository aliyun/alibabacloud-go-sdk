// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModuleItemListSubItemsShoppingItemMapValueCabinQuantityValue interface {
	dara.Model
	String() string
	GoString() string
	SetCabin(v string) *ModuleItemListSubItemsShoppingItemMapValueCabinQuantityValue
	GetCabin() *string
	SetCabinClass(v string) *ModuleItemListSubItemsShoppingItemMapValueCabinQuantityValue
	GetCabinClass() *string
	SetCabinClassName(v string) *ModuleItemListSubItemsShoppingItemMapValueCabinQuantityValue
	GetCabinClassName() *string
	SetCabinClassMemo(v string) *ModuleItemListSubItemsShoppingItemMapValueCabinQuantityValue
	GetCabinClassMemo() *string
	SetSpecification(v string) *ModuleItemListSubItemsShoppingItemMapValueCabinQuantityValue
	GetSpecification() *string
	SetQuantity(v string) *ModuleItemListSubItemsShoppingItemMapValueCabinQuantityValue
	GetQuantity() *string
	SetLinkCabins(v []*string) *ModuleItemListSubItemsShoppingItemMapValueCabinQuantityValue
	GetLinkCabins() []*string
	SetReshopChangeCabin(v bool) *ModuleItemListSubItemsShoppingItemMapValueCabinQuantityValue
	GetReshopChangeCabin() *bool
	SetChildCabinType(v int32) *ModuleItemListSubItemsShoppingItemMapValueCabinQuantityValue
	GetChildCabinType() *int32
	SetInfantBasicCabin(v string) *ModuleItemListSubItemsShoppingItemMapValueCabinQuantityValue
	GetInfantBasicCabin() *string
	SetInnerCabinClass(v int32) *ModuleItemListSubItemsShoppingItemMapValueCabinQuantityValue
	GetInnerCabinClass() *int32
}

type ModuleItemListSubItemsShoppingItemMapValueCabinQuantityValue struct {
	// example:
	//
	// R
	Cabin *string `json:"cabin,omitempty" xml:"cabin,omitempty"`
	// example:
	//
	// Y
	CabinClass     *string `json:"cabin_class,omitempty" xml:"cabin_class,omitempty"`
	CabinClassName *string `json:"cabin_class_name,omitempty" xml:"cabin_class_name,omitempty"`
	CabinClassMemo *string `json:"cabin_class_memo,omitempty" xml:"cabin_class_memo,omitempty"`
	Specification  *string `json:"specification,omitempty" xml:"specification,omitempty"`
	// example:
	//
	// A
	Quantity   *string   `json:"quantity,omitempty" xml:"quantity,omitempty"`
	LinkCabins []*string `json:"link_cabins,omitempty" xml:"link_cabins,omitempty" type:"Repeated"`
	// example:
	//
	// false
	ReshopChangeCabin *bool   `json:"reshop_change_cabin,omitempty" xml:"reshop_change_cabin,omitempty"`
	ChildCabinType    *int32  `json:"child_cabin_type,omitempty" xml:"child_cabin_type,omitempty"`
	InfantBasicCabin  *string `json:"infant_basic_cabin,omitempty" xml:"infant_basic_cabin,omitempty"`
	// example:
	//
	// 2
	InnerCabinClass *int32 `json:"inner_cabin_class,omitempty" xml:"inner_cabin_class,omitempty"`
}

func (s ModuleItemListSubItemsShoppingItemMapValueCabinQuantityValue) String() string {
	return dara.Prettify(s)
}

func (s ModuleItemListSubItemsShoppingItemMapValueCabinQuantityValue) GoString() string {
	return s.String()
}

func (s *ModuleItemListSubItemsShoppingItemMapValueCabinQuantityValue) GetCabin() *string {
	return s.Cabin
}

func (s *ModuleItemListSubItemsShoppingItemMapValueCabinQuantityValue) GetCabinClass() *string {
	return s.CabinClass
}

func (s *ModuleItemListSubItemsShoppingItemMapValueCabinQuantityValue) GetCabinClassName() *string {
	return s.CabinClassName
}

func (s *ModuleItemListSubItemsShoppingItemMapValueCabinQuantityValue) GetCabinClassMemo() *string {
	return s.CabinClassMemo
}

func (s *ModuleItemListSubItemsShoppingItemMapValueCabinQuantityValue) GetSpecification() *string {
	return s.Specification
}

func (s *ModuleItemListSubItemsShoppingItemMapValueCabinQuantityValue) GetQuantity() *string {
	return s.Quantity
}

func (s *ModuleItemListSubItemsShoppingItemMapValueCabinQuantityValue) GetLinkCabins() []*string {
	return s.LinkCabins
}

func (s *ModuleItemListSubItemsShoppingItemMapValueCabinQuantityValue) GetReshopChangeCabin() *bool {
	return s.ReshopChangeCabin
}

func (s *ModuleItemListSubItemsShoppingItemMapValueCabinQuantityValue) GetChildCabinType() *int32 {
	return s.ChildCabinType
}

func (s *ModuleItemListSubItemsShoppingItemMapValueCabinQuantityValue) GetInfantBasicCabin() *string {
	return s.InfantBasicCabin
}

func (s *ModuleItemListSubItemsShoppingItemMapValueCabinQuantityValue) GetInnerCabinClass() *int32 {
	return s.InnerCabinClass
}

func (s *ModuleItemListSubItemsShoppingItemMapValueCabinQuantityValue) SetCabin(v string) *ModuleItemListSubItemsShoppingItemMapValueCabinQuantityValue {
	s.Cabin = &v
	return s
}

func (s *ModuleItemListSubItemsShoppingItemMapValueCabinQuantityValue) SetCabinClass(v string) *ModuleItemListSubItemsShoppingItemMapValueCabinQuantityValue {
	s.CabinClass = &v
	return s
}

func (s *ModuleItemListSubItemsShoppingItemMapValueCabinQuantityValue) SetCabinClassName(v string) *ModuleItemListSubItemsShoppingItemMapValueCabinQuantityValue {
	s.CabinClassName = &v
	return s
}

func (s *ModuleItemListSubItemsShoppingItemMapValueCabinQuantityValue) SetCabinClassMemo(v string) *ModuleItemListSubItemsShoppingItemMapValueCabinQuantityValue {
	s.CabinClassMemo = &v
	return s
}

func (s *ModuleItemListSubItemsShoppingItemMapValueCabinQuantityValue) SetSpecification(v string) *ModuleItemListSubItemsShoppingItemMapValueCabinQuantityValue {
	s.Specification = &v
	return s
}

func (s *ModuleItemListSubItemsShoppingItemMapValueCabinQuantityValue) SetQuantity(v string) *ModuleItemListSubItemsShoppingItemMapValueCabinQuantityValue {
	s.Quantity = &v
	return s
}

func (s *ModuleItemListSubItemsShoppingItemMapValueCabinQuantityValue) SetLinkCabins(v []*string) *ModuleItemListSubItemsShoppingItemMapValueCabinQuantityValue {
	s.LinkCabins = v
	return s
}

func (s *ModuleItemListSubItemsShoppingItemMapValueCabinQuantityValue) SetReshopChangeCabin(v bool) *ModuleItemListSubItemsShoppingItemMapValueCabinQuantityValue {
	s.ReshopChangeCabin = &v
	return s
}

func (s *ModuleItemListSubItemsShoppingItemMapValueCabinQuantityValue) SetChildCabinType(v int32) *ModuleItemListSubItemsShoppingItemMapValueCabinQuantityValue {
	s.ChildCabinType = &v
	return s
}

func (s *ModuleItemListSubItemsShoppingItemMapValueCabinQuantityValue) SetInfantBasicCabin(v string) *ModuleItemListSubItemsShoppingItemMapValueCabinQuantityValue {
	s.InfantBasicCabin = &v
	return s
}

func (s *ModuleItemListSubItemsShoppingItemMapValueCabinQuantityValue) SetInnerCabinClass(v int32) *ModuleItemListSubItemsShoppingItemMapValueCabinQuantityValue {
	s.InnerCabinClass = &v
	return s
}

func (s *ModuleItemListSubItemsShoppingItemMapValueCabinQuantityValue) Validate() error {
	return dara.Validate(s)
}
