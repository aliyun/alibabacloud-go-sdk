// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModuleGroupItemSubItemsShoppingItemMapValue interface {
	dara.Model
	String() string
	GoString() string
	SetSearchPrice(v *ModuleGroupItemSubItemsShoppingItemMapValueSearchPrice) *ModuleGroupItemSubItemsShoppingItemMapValue
	GetSearchPrice() *ModuleGroupItemSubItemsShoppingItemMapValueSearchPrice
}

type ModuleGroupItemSubItemsShoppingItemMapValue struct {
	SearchPrice *ModuleGroupItemSubItemsShoppingItemMapValueSearchPrice `json:"search_price,omitempty" xml:"search_price,omitempty" type:"Struct"`
}

func (s ModuleGroupItemSubItemsShoppingItemMapValue) String() string {
	return dara.Prettify(s)
}

func (s ModuleGroupItemSubItemsShoppingItemMapValue) GoString() string {
	return s.String()
}

func (s *ModuleGroupItemSubItemsShoppingItemMapValue) GetSearchPrice() *ModuleGroupItemSubItemsShoppingItemMapValueSearchPrice {
	return s.SearchPrice
}

func (s *ModuleGroupItemSubItemsShoppingItemMapValue) SetSearchPrice(v *ModuleGroupItemSubItemsShoppingItemMapValueSearchPrice) *ModuleGroupItemSubItemsShoppingItemMapValue {
	s.SearchPrice = v
	return s
}

func (s *ModuleGroupItemSubItemsShoppingItemMapValue) Validate() error {
	return dara.Validate(s)
}

type ModuleGroupItemSubItemsShoppingItemMapValueSearchPrice struct {
	// example:
	//
	// 120000
	TicketPrice *int32 `json:"ticket_price,omitempty" xml:"ticket_price,omitempty"`
	// example:
	//
	// 120000
	SellPrice *int32 `json:"sell_price,omitempty" xml:"sell_price,omitempty"`
	// example:
	//
	// 6000
	Tax *int32 `json:"tax,omitempty" xml:"tax,omitempty"`
}

func (s ModuleGroupItemSubItemsShoppingItemMapValueSearchPrice) String() string {
	return dara.Prettify(s)
}

func (s ModuleGroupItemSubItemsShoppingItemMapValueSearchPrice) GoString() string {
	return s.String()
}

func (s *ModuleGroupItemSubItemsShoppingItemMapValueSearchPrice) GetTicketPrice() *int32 {
	return s.TicketPrice
}

func (s *ModuleGroupItemSubItemsShoppingItemMapValueSearchPrice) GetSellPrice() *int32 {
	return s.SellPrice
}

func (s *ModuleGroupItemSubItemsShoppingItemMapValueSearchPrice) GetTax() *int32 {
	return s.Tax
}

func (s *ModuleGroupItemSubItemsShoppingItemMapValueSearchPrice) SetTicketPrice(v int32) *ModuleGroupItemSubItemsShoppingItemMapValueSearchPrice {
	s.TicketPrice = &v
	return s
}

func (s *ModuleGroupItemSubItemsShoppingItemMapValueSearchPrice) SetSellPrice(v int32) *ModuleGroupItemSubItemsShoppingItemMapValueSearchPrice {
	s.SellPrice = &v
	return s
}

func (s *ModuleGroupItemSubItemsShoppingItemMapValueSearchPrice) SetTax(v int32) *ModuleGroupItemSubItemsShoppingItemMapValueSearchPrice {
	s.Tax = &v
	return s
}

func (s *ModuleGroupItemSubItemsShoppingItemMapValueSearchPrice) Validate() error {
	return dara.Validate(s)
}
