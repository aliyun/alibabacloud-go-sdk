// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModuleItemListShoppingItemMapValueCabinQuantityValue interface {
	dara.Model
	String() string
	GoString() string
	SetCabin(v string) *ModuleItemListShoppingItemMapValueCabinQuantityValue
	GetCabin() *string
	SetCabinClass(v string) *ModuleItemListShoppingItemMapValueCabinQuantityValue
	GetCabinClass() *string
	SetCabinClassName(v string) *ModuleItemListShoppingItemMapValueCabinQuantityValue
	GetCabinClassName() *string
	SetCabinClassMemo(v string) *ModuleItemListShoppingItemMapValueCabinQuantityValue
	GetCabinClassMemo() *string
	SetSpecification(v string) *ModuleItemListShoppingItemMapValueCabinQuantityValue
	GetSpecification() *string
	SetQuantity(v string) *ModuleItemListShoppingItemMapValueCabinQuantityValue
	GetQuantity() *string
	SetLinkCabins(v []*string) *ModuleItemListShoppingItemMapValueCabinQuantityValue
	GetLinkCabins() []*string
	SetReshopChangeCabin(v bool) *ModuleItemListShoppingItemMapValueCabinQuantityValue
	GetReshopChangeCabin() *bool
	SetChildCabinType(v int32) *ModuleItemListShoppingItemMapValueCabinQuantityValue
	GetChildCabinType() *int32
	SetInfantBasicCabin(v string) *ModuleItemListShoppingItemMapValueCabinQuantityValue
	GetInfantBasicCabin() *string
	SetInnerCabinClass(v int32) *ModuleItemListShoppingItemMapValueCabinQuantityValue
	GetInnerCabinClass() *int32
}

type ModuleItemListShoppingItemMapValueCabinQuantityValue struct {
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

func (s ModuleItemListShoppingItemMapValueCabinQuantityValue) String() string {
	return dara.Prettify(s)
}

func (s ModuleItemListShoppingItemMapValueCabinQuantityValue) GoString() string {
	return s.String()
}

func (s *ModuleItemListShoppingItemMapValueCabinQuantityValue) GetCabin() *string {
	return s.Cabin
}

func (s *ModuleItemListShoppingItemMapValueCabinQuantityValue) GetCabinClass() *string {
	return s.CabinClass
}

func (s *ModuleItemListShoppingItemMapValueCabinQuantityValue) GetCabinClassName() *string {
	return s.CabinClassName
}

func (s *ModuleItemListShoppingItemMapValueCabinQuantityValue) GetCabinClassMemo() *string {
	return s.CabinClassMemo
}

func (s *ModuleItemListShoppingItemMapValueCabinQuantityValue) GetSpecification() *string {
	return s.Specification
}

func (s *ModuleItemListShoppingItemMapValueCabinQuantityValue) GetQuantity() *string {
	return s.Quantity
}

func (s *ModuleItemListShoppingItemMapValueCabinQuantityValue) GetLinkCabins() []*string {
	return s.LinkCabins
}

func (s *ModuleItemListShoppingItemMapValueCabinQuantityValue) GetReshopChangeCabin() *bool {
	return s.ReshopChangeCabin
}

func (s *ModuleItemListShoppingItemMapValueCabinQuantityValue) GetChildCabinType() *int32 {
	return s.ChildCabinType
}

func (s *ModuleItemListShoppingItemMapValueCabinQuantityValue) GetInfantBasicCabin() *string {
	return s.InfantBasicCabin
}

func (s *ModuleItemListShoppingItemMapValueCabinQuantityValue) GetInnerCabinClass() *int32 {
	return s.InnerCabinClass
}

func (s *ModuleItemListShoppingItemMapValueCabinQuantityValue) SetCabin(v string) *ModuleItemListShoppingItemMapValueCabinQuantityValue {
	s.Cabin = &v
	return s
}

func (s *ModuleItemListShoppingItemMapValueCabinQuantityValue) SetCabinClass(v string) *ModuleItemListShoppingItemMapValueCabinQuantityValue {
	s.CabinClass = &v
	return s
}

func (s *ModuleItemListShoppingItemMapValueCabinQuantityValue) SetCabinClassName(v string) *ModuleItemListShoppingItemMapValueCabinQuantityValue {
	s.CabinClassName = &v
	return s
}

func (s *ModuleItemListShoppingItemMapValueCabinQuantityValue) SetCabinClassMemo(v string) *ModuleItemListShoppingItemMapValueCabinQuantityValue {
	s.CabinClassMemo = &v
	return s
}

func (s *ModuleItemListShoppingItemMapValueCabinQuantityValue) SetSpecification(v string) *ModuleItemListShoppingItemMapValueCabinQuantityValue {
	s.Specification = &v
	return s
}

func (s *ModuleItemListShoppingItemMapValueCabinQuantityValue) SetQuantity(v string) *ModuleItemListShoppingItemMapValueCabinQuantityValue {
	s.Quantity = &v
	return s
}

func (s *ModuleItemListShoppingItemMapValueCabinQuantityValue) SetLinkCabins(v []*string) *ModuleItemListShoppingItemMapValueCabinQuantityValue {
	s.LinkCabins = v
	return s
}

func (s *ModuleItemListShoppingItemMapValueCabinQuantityValue) SetReshopChangeCabin(v bool) *ModuleItemListShoppingItemMapValueCabinQuantityValue {
	s.ReshopChangeCabin = &v
	return s
}

func (s *ModuleItemListShoppingItemMapValueCabinQuantityValue) SetChildCabinType(v int32) *ModuleItemListShoppingItemMapValueCabinQuantityValue {
	s.ChildCabinType = &v
	return s
}

func (s *ModuleItemListShoppingItemMapValueCabinQuantityValue) SetInfantBasicCabin(v string) *ModuleItemListShoppingItemMapValueCabinQuantityValue {
	s.InfantBasicCabin = &v
	return s
}

func (s *ModuleItemListShoppingItemMapValueCabinQuantityValue) SetInnerCabinClass(v int32) *ModuleItemListShoppingItemMapValueCabinQuantityValue {
	s.InnerCabinClass = &v
	return s
}

func (s *ModuleItemListShoppingItemMapValueCabinQuantityValue) Validate() error {
	return dara.Validate(s)
}
