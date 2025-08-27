// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModuleFlightItemListItemListShoppingItemMapValueCabinQuantityValue interface {
	dara.Model
	String() string
	GoString() string
	SetCabin(v string) *ModuleFlightItemListItemListShoppingItemMapValueCabinQuantityValue
	GetCabin() *string
	SetCabinClass(v string) *ModuleFlightItemListItemListShoppingItemMapValueCabinQuantityValue
	GetCabinClass() *string
	SetCabinClassName(v string) *ModuleFlightItemListItemListShoppingItemMapValueCabinQuantityValue
	GetCabinClassName() *string
	SetCabinClassMemo(v string) *ModuleFlightItemListItemListShoppingItemMapValueCabinQuantityValue
	GetCabinClassMemo() *string
	SetSpecification(v string) *ModuleFlightItemListItemListShoppingItemMapValueCabinQuantityValue
	GetSpecification() *string
	SetQuantity(v string) *ModuleFlightItemListItemListShoppingItemMapValueCabinQuantityValue
	GetQuantity() *string
	SetLinkCabins(v []*string) *ModuleFlightItemListItemListShoppingItemMapValueCabinQuantityValue
	GetLinkCabins() []*string
	SetReshopChangeCabin(v bool) *ModuleFlightItemListItemListShoppingItemMapValueCabinQuantityValue
	GetReshopChangeCabin() *bool
	SetChildCabinType(v int32) *ModuleFlightItemListItemListShoppingItemMapValueCabinQuantityValue
	GetChildCabinType() *int32
	SetInfantBasicCabin(v string) *ModuleFlightItemListItemListShoppingItemMapValueCabinQuantityValue
	GetInfantBasicCabin() *string
	SetInnerCabinClass(v int32) *ModuleFlightItemListItemListShoppingItemMapValueCabinQuantityValue
	GetInnerCabinClass() *int32
}

type ModuleFlightItemListItemListShoppingItemMapValueCabinQuantityValue struct {
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

func (s ModuleFlightItemListItemListShoppingItemMapValueCabinQuantityValue) String() string {
	return dara.Prettify(s)
}

func (s ModuleFlightItemListItemListShoppingItemMapValueCabinQuantityValue) GoString() string {
	return s.String()
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueCabinQuantityValue) GetCabin() *string {
	return s.Cabin
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueCabinQuantityValue) GetCabinClass() *string {
	return s.CabinClass
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueCabinQuantityValue) GetCabinClassName() *string {
	return s.CabinClassName
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueCabinQuantityValue) GetCabinClassMemo() *string {
	return s.CabinClassMemo
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueCabinQuantityValue) GetSpecification() *string {
	return s.Specification
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueCabinQuantityValue) GetQuantity() *string {
	return s.Quantity
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueCabinQuantityValue) GetLinkCabins() []*string {
	return s.LinkCabins
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueCabinQuantityValue) GetReshopChangeCabin() *bool {
	return s.ReshopChangeCabin
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueCabinQuantityValue) GetChildCabinType() *int32 {
	return s.ChildCabinType
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueCabinQuantityValue) GetInfantBasicCabin() *string {
	return s.InfantBasicCabin
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueCabinQuantityValue) GetInnerCabinClass() *int32 {
	return s.InnerCabinClass
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueCabinQuantityValue) SetCabin(v string) *ModuleFlightItemListItemListShoppingItemMapValueCabinQuantityValue {
	s.Cabin = &v
	return s
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueCabinQuantityValue) SetCabinClass(v string) *ModuleFlightItemListItemListShoppingItemMapValueCabinQuantityValue {
	s.CabinClass = &v
	return s
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueCabinQuantityValue) SetCabinClassName(v string) *ModuleFlightItemListItemListShoppingItemMapValueCabinQuantityValue {
	s.CabinClassName = &v
	return s
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueCabinQuantityValue) SetCabinClassMemo(v string) *ModuleFlightItemListItemListShoppingItemMapValueCabinQuantityValue {
	s.CabinClassMemo = &v
	return s
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueCabinQuantityValue) SetSpecification(v string) *ModuleFlightItemListItemListShoppingItemMapValueCabinQuantityValue {
	s.Specification = &v
	return s
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueCabinQuantityValue) SetQuantity(v string) *ModuleFlightItemListItemListShoppingItemMapValueCabinQuantityValue {
	s.Quantity = &v
	return s
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueCabinQuantityValue) SetLinkCabins(v []*string) *ModuleFlightItemListItemListShoppingItemMapValueCabinQuantityValue {
	s.LinkCabins = v
	return s
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueCabinQuantityValue) SetReshopChangeCabin(v bool) *ModuleFlightItemListItemListShoppingItemMapValueCabinQuantityValue {
	s.ReshopChangeCabin = &v
	return s
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueCabinQuantityValue) SetChildCabinType(v int32) *ModuleFlightItemListItemListShoppingItemMapValueCabinQuantityValue {
	s.ChildCabinType = &v
	return s
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueCabinQuantityValue) SetInfantBasicCabin(v string) *ModuleFlightItemListItemListShoppingItemMapValueCabinQuantityValue {
	s.InfantBasicCabin = &v
	return s
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueCabinQuantityValue) SetInnerCabinClass(v int32) *ModuleFlightItemListItemListShoppingItemMapValueCabinQuantityValue {
	s.InnerCabinClass = &v
	return s
}

func (s *ModuleFlightItemListItemListShoppingItemMapValueCabinQuantityValue) Validate() error {
	return dara.Validate(s)
}
