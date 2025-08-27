// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModuleFlightItemListBestPriceItemShoppingItemMapValueCabinQuantityValue interface {
	dara.Model
	String() string
	GoString() string
	SetCabin(v string) *ModuleFlightItemListBestPriceItemShoppingItemMapValueCabinQuantityValue
	GetCabin() *string
	SetCabinClass(v string) *ModuleFlightItemListBestPriceItemShoppingItemMapValueCabinQuantityValue
	GetCabinClass() *string
	SetCabinClassName(v string) *ModuleFlightItemListBestPriceItemShoppingItemMapValueCabinQuantityValue
	GetCabinClassName() *string
	SetCabinClassMemo(v string) *ModuleFlightItemListBestPriceItemShoppingItemMapValueCabinQuantityValue
	GetCabinClassMemo() *string
	SetSpecification(v string) *ModuleFlightItemListBestPriceItemShoppingItemMapValueCabinQuantityValue
	GetSpecification() *string
	SetQuantity(v string) *ModuleFlightItemListBestPriceItemShoppingItemMapValueCabinQuantityValue
	GetQuantity() *string
	SetLinkCabins(v []*string) *ModuleFlightItemListBestPriceItemShoppingItemMapValueCabinQuantityValue
	GetLinkCabins() []*string
	SetReshopChangeCabin(v bool) *ModuleFlightItemListBestPriceItemShoppingItemMapValueCabinQuantityValue
	GetReshopChangeCabin() *bool
	SetChildCabinType(v int32) *ModuleFlightItemListBestPriceItemShoppingItemMapValueCabinQuantityValue
	GetChildCabinType() *int32
	SetInfantBasicCabin(v string) *ModuleFlightItemListBestPriceItemShoppingItemMapValueCabinQuantityValue
	GetInfantBasicCabin() *string
	SetInnerCabinClass(v int32) *ModuleFlightItemListBestPriceItemShoppingItemMapValueCabinQuantityValue
	GetInnerCabinClass() *int32
}

type ModuleFlightItemListBestPriceItemShoppingItemMapValueCabinQuantityValue struct {
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

func (s ModuleFlightItemListBestPriceItemShoppingItemMapValueCabinQuantityValue) String() string {
	return dara.Prettify(s)
}

func (s ModuleFlightItemListBestPriceItemShoppingItemMapValueCabinQuantityValue) GoString() string {
	return s.String()
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueCabinQuantityValue) GetCabin() *string {
	return s.Cabin
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueCabinQuantityValue) GetCabinClass() *string {
	return s.CabinClass
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueCabinQuantityValue) GetCabinClassName() *string {
	return s.CabinClassName
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueCabinQuantityValue) GetCabinClassMemo() *string {
	return s.CabinClassMemo
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueCabinQuantityValue) GetSpecification() *string {
	return s.Specification
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueCabinQuantityValue) GetQuantity() *string {
	return s.Quantity
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueCabinQuantityValue) GetLinkCabins() []*string {
	return s.LinkCabins
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueCabinQuantityValue) GetReshopChangeCabin() *bool {
	return s.ReshopChangeCabin
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueCabinQuantityValue) GetChildCabinType() *int32 {
	return s.ChildCabinType
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueCabinQuantityValue) GetInfantBasicCabin() *string {
	return s.InfantBasicCabin
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueCabinQuantityValue) GetInnerCabinClass() *int32 {
	return s.InnerCabinClass
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueCabinQuantityValue) SetCabin(v string) *ModuleFlightItemListBestPriceItemShoppingItemMapValueCabinQuantityValue {
	s.Cabin = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueCabinQuantityValue) SetCabinClass(v string) *ModuleFlightItemListBestPriceItemShoppingItemMapValueCabinQuantityValue {
	s.CabinClass = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueCabinQuantityValue) SetCabinClassName(v string) *ModuleFlightItemListBestPriceItemShoppingItemMapValueCabinQuantityValue {
	s.CabinClassName = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueCabinQuantityValue) SetCabinClassMemo(v string) *ModuleFlightItemListBestPriceItemShoppingItemMapValueCabinQuantityValue {
	s.CabinClassMemo = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueCabinQuantityValue) SetSpecification(v string) *ModuleFlightItemListBestPriceItemShoppingItemMapValueCabinQuantityValue {
	s.Specification = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueCabinQuantityValue) SetQuantity(v string) *ModuleFlightItemListBestPriceItemShoppingItemMapValueCabinQuantityValue {
	s.Quantity = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueCabinQuantityValue) SetLinkCabins(v []*string) *ModuleFlightItemListBestPriceItemShoppingItemMapValueCabinQuantityValue {
	s.LinkCabins = v
	return s
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueCabinQuantityValue) SetReshopChangeCabin(v bool) *ModuleFlightItemListBestPriceItemShoppingItemMapValueCabinQuantityValue {
	s.ReshopChangeCabin = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueCabinQuantityValue) SetChildCabinType(v int32) *ModuleFlightItemListBestPriceItemShoppingItemMapValueCabinQuantityValue {
	s.ChildCabinType = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueCabinQuantityValue) SetInfantBasicCabin(v string) *ModuleFlightItemListBestPriceItemShoppingItemMapValueCabinQuantityValue {
	s.InfantBasicCabin = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueCabinQuantityValue) SetInnerCabinClass(v int32) *ModuleFlightItemListBestPriceItemShoppingItemMapValueCabinQuantityValue {
	s.InnerCabinClass = &v
	return s
}

func (s *ModuleFlightItemListBestPriceItemShoppingItemMapValueCabinQuantityValue) Validate() error {
	return dara.Validate(s)
}
