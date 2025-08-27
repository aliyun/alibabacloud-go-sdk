// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModuleFlightItemListItemListSubItemsShoppingItemMapValueCabinQuantityValue interface {
	dara.Model
	String() string
	GoString() string
	SetCabin(v string) *ModuleFlightItemListItemListSubItemsShoppingItemMapValueCabinQuantityValue
	GetCabin() *string
	SetCabinClass(v string) *ModuleFlightItemListItemListSubItemsShoppingItemMapValueCabinQuantityValue
	GetCabinClass() *string
	SetCabinClassName(v string) *ModuleFlightItemListItemListSubItemsShoppingItemMapValueCabinQuantityValue
	GetCabinClassName() *string
	SetCabinClassMemo(v string) *ModuleFlightItemListItemListSubItemsShoppingItemMapValueCabinQuantityValue
	GetCabinClassMemo() *string
	SetSpecification(v string) *ModuleFlightItemListItemListSubItemsShoppingItemMapValueCabinQuantityValue
	GetSpecification() *string
	SetQuantity(v string) *ModuleFlightItemListItemListSubItemsShoppingItemMapValueCabinQuantityValue
	GetQuantity() *string
	SetLinkCabins(v []*string) *ModuleFlightItemListItemListSubItemsShoppingItemMapValueCabinQuantityValue
	GetLinkCabins() []*string
	SetReshopChangeCabin(v bool) *ModuleFlightItemListItemListSubItemsShoppingItemMapValueCabinQuantityValue
	GetReshopChangeCabin() *bool
	SetChildCabinType(v int32) *ModuleFlightItemListItemListSubItemsShoppingItemMapValueCabinQuantityValue
	GetChildCabinType() *int32
	SetInfantBasicCabin(v string) *ModuleFlightItemListItemListSubItemsShoppingItemMapValueCabinQuantityValue
	GetInfantBasicCabin() *string
	SetInnerCabinClass(v int32) *ModuleFlightItemListItemListSubItemsShoppingItemMapValueCabinQuantityValue
	GetInnerCabinClass() *int32
}

type ModuleFlightItemListItemListSubItemsShoppingItemMapValueCabinQuantityValue struct {
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

func (s ModuleFlightItemListItemListSubItemsShoppingItemMapValueCabinQuantityValue) String() string {
	return dara.Prettify(s)
}

func (s ModuleFlightItemListItemListSubItemsShoppingItemMapValueCabinQuantityValue) GoString() string {
	return s.String()
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueCabinQuantityValue) GetCabin() *string {
	return s.Cabin
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueCabinQuantityValue) GetCabinClass() *string {
	return s.CabinClass
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueCabinQuantityValue) GetCabinClassName() *string {
	return s.CabinClassName
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueCabinQuantityValue) GetCabinClassMemo() *string {
	return s.CabinClassMemo
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueCabinQuantityValue) GetSpecification() *string {
	return s.Specification
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueCabinQuantityValue) GetQuantity() *string {
	return s.Quantity
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueCabinQuantityValue) GetLinkCabins() []*string {
	return s.LinkCabins
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueCabinQuantityValue) GetReshopChangeCabin() *bool {
	return s.ReshopChangeCabin
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueCabinQuantityValue) GetChildCabinType() *int32 {
	return s.ChildCabinType
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueCabinQuantityValue) GetInfantBasicCabin() *string {
	return s.InfantBasicCabin
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueCabinQuantityValue) GetInnerCabinClass() *int32 {
	return s.InnerCabinClass
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueCabinQuantityValue) SetCabin(v string) *ModuleFlightItemListItemListSubItemsShoppingItemMapValueCabinQuantityValue {
	s.Cabin = &v
	return s
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueCabinQuantityValue) SetCabinClass(v string) *ModuleFlightItemListItemListSubItemsShoppingItemMapValueCabinQuantityValue {
	s.CabinClass = &v
	return s
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueCabinQuantityValue) SetCabinClassName(v string) *ModuleFlightItemListItemListSubItemsShoppingItemMapValueCabinQuantityValue {
	s.CabinClassName = &v
	return s
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueCabinQuantityValue) SetCabinClassMemo(v string) *ModuleFlightItemListItemListSubItemsShoppingItemMapValueCabinQuantityValue {
	s.CabinClassMemo = &v
	return s
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueCabinQuantityValue) SetSpecification(v string) *ModuleFlightItemListItemListSubItemsShoppingItemMapValueCabinQuantityValue {
	s.Specification = &v
	return s
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueCabinQuantityValue) SetQuantity(v string) *ModuleFlightItemListItemListSubItemsShoppingItemMapValueCabinQuantityValue {
	s.Quantity = &v
	return s
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueCabinQuantityValue) SetLinkCabins(v []*string) *ModuleFlightItemListItemListSubItemsShoppingItemMapValueCabinQuantityValue {
	s.LinkCabins = v
	return s
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueCabinQuantityValue) SetReshopChangeCabin(v bool) *ModuleFlightItemListItemListSubItemsShoppingItemMapValueCabinQuantityValue {
	s.ReshopChangeCabin = &v
	return s
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueCabinQuantityValue) SetChildCabinType(v int32) *ModuleFlightItemListItemListSubItemsShoppingItemMapValueCabinQuantityValue {
	s.ChildCabinType = &v
	return s
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueCabinQuantityValue) SetInfantBasicCabin(v string) *ModuleFlightItemListItemListSubItemsShoppingItemMapValueCabinQuantityValue {
	s.InfantBasicCabin = &v
	return s
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueCabinQuantityValue) SetInnerCabinClass(v int32) *ModuleFlightItemListItemListSubItemsShoppingItemMapValueCabinQuantityValue {
	s.InnerCabinClass = &v
	return s
}

func (s *ModuleFlightItemListItemListSubItemsShoppingItemMapValueCabinQuantityValue) Validate() error {
	return dara.Validate(s)
}
