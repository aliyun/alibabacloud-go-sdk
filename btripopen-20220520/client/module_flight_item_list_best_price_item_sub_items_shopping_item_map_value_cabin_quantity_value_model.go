// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueCabinQuantityValue interface {
	dara.Model
	String() string
	GoString() string
	SetCabin(v string) *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueCabinQuantityValue
	GetCabin() *string
	SetCabinClass(v string) *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueCabinQuantityValue
	GetCabinClass() *string
	SetCabinClassName(v string) *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueCabinQuantityValue
	GetCabinClassName() *string
	SetCabinClassMemo(v string) *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueCabinQuantityValue
	GetCabinClassMemo() *string
	SetSpecification(v string) *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueCabinQuantityValue
	GetSpecification() *string
	SetQuantity(v string) *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueCabinQuantityValue
	GetQuantity() *string
	SetLinkCabins(v []*string) *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueCabinQuantityValue
	GetLinkCabins() []*string
	SetReshopChangeCabin(v bool) *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueCabinQuantityValue
	GetReshopChangeCabin() *bool
	SetChildCabinType(v int32) *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueCabinQuantityValue
	GetChildCabinType() *int32
	SetInfantBasicCabin(v string) *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueCabinQuantityValue
	GetInfantBasicCabin() *string
	SetInnerCabinClass(v int32) *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueCabinQuantityValue
	GetInnerCabinClass() *int32
}

type ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueCabinQuantityValue struct {
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

func (s ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueCabinQuantityValue) String() string {
	return dara.Prettify(s)
}

func (s ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueCabinQuantityValue) GoString() string {
	return s.String()
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueCabinQuantityValue) GetCabin() *string {
	return s.Cabin
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueCabinQuantityValue) GetCabinClass() *string {
	return s.CabinClass
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueCabinQuantityValue) GetCabinClassName() *string {
	return s.CabinClassName
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueCabinQuantityValue) GetCabinClassMemo() *string {
	return s.CabinClassMemo
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueCabinQuantityValue) GetSpecification() *string {
	return s.Specification
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueCabinQuantityValue) GetQuantity() *string {
	return s.Quantity
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueCabinQuantityValue) GetLinkCabins() []*string {
	return s.LinkCabins
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueCabinQuantityValue) GetReshopChangeCabin() *bool {
	return s.ReshopChangeCabin
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueCabinQuantityValue) GetChildCabinType() *int32 {
	return s.ChildCabinType
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueCabinQuantityValue) GetInfantBasicCabin() *string {
	return s.InfantBasicCabin
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueCabinQuantityValue) GetInnerCabinClass() *int32 {
	return s.InnerCabinClass
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueCabinQuantityValue) SetCabin(v string) *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueCabinQuantityValue {
	s.Cabin = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueCabinQuantityValue) SetCabinClass(v string) *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueCabinQuantityValue {
	s.CabinClass = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueCabinQuantityValue) SetCabinClassName(v string) *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueCabinQuantityValue {
	s.CabinClassName = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueCabinQuantityValue) SetCabinClassMemo(v string) *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueCabinQuantityValue {
	s.CabinClassMemo = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueCabinQuantityValue) SetSpecification(v string) *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueCabinQuantityValue {
	s.Specification = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueCabinQuantityValue) SetQuantity(v string) *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueCabinQuantityValue {
	s.Quantity = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueCabinQuantityValue) SetLinkCabins(v []*string) *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueCabinQuantityValue {
	s.LinkCabins = v
	return s
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueCabinQuantityValue) SetReshopChangeCabin(v bool) *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueCabinQuantityValue {
	s.ReshopChangeCabin = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueCabinQuantityValue) SetChildCabinType(v int32) *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueCabinQuantityValue {
	s.ChildCabinType = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueCabinQuantityValue) SetInfantBasicCabin(v string) *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueCabinQuantityValue {
	s.InfantBasicCabin = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueCabinQuantityValue) SetInnerCabinClass(v int32) *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueCabinQuantityValue {
	s.InnerCabinClass = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemSubItemsShoppingItemMapValueCabinQuantityValue) Validate() error {
	return dara.Validate(s)
}
