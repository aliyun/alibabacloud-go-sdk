// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEpdInventoryConstituteItem interface {
  dara.Model
  String() string
  GoString() string
  SetCarbonEmission(v float64) *EpdInventoryConstituteItem
  GetCarbonEmission() *float64 
  SetFactor(v string) *EpdInventoryConstituteItem
  GetFactor() *string 
  SetFactorDataset(v string) *EpdInventoryConstituteItem
  GetFactorDataset() *string 
  SetFactorId(v string) *EpdInventoryConstituteItem
  GetFactorId() *string 
  SetFactorType(v int32) *EpdInventoryConstituteItem
  GetFactorType() *int32 
  SetFactorUnit(v string) *EpdInventoryConstituteItem
  GetFactorUnit() *string 
  SetInventoryId(v int64) *EpdInventoryConstituteItem
  GetInventoryId() *int64 
  SetInventoryUnit(v string) *EpdInventoryConstituteItem
  GetInventoryUnit() *string 
  SetInventoryValue(v float64) *EpdInventoryConstituteItem
  GetInventoryValue() *float64 
  SetInventoryValuePerProduct(v float64) *EpdInventoryConstituteItem
  GetInventoryValuePerProduct() *float64 
  SetInventoryValuePerProductUnit(v string) *EpdInventoryConstituteItem
  GetInventoryValuePerProductUnit() *string 
  SetItems(v []*EpdInventoryConstituteItem) *EpdInventoryConstituteItem
  GetItems() []*EpdInventoryConstituteItem 
  SetName(v string) *EpdInventoryConstituteItem
  GetName() *string 
  SetNum(v int64) *EpdInventoryConstituteItem
  GetNum() *int64 
  SetPercent(v float64) *EpdInventoryConstituteItem
  GetPercent() *float64 
  SetQuantity(v float64) *EpdInventoryConstituteItem
  GetQuantity() *float64 
  SetResourceType(v string) *EpdInventoryConstituteItem
  GetResourceType() *string 
  SetState(v int32) *EpdInventoryConstituteItem
  GetState() *int32 
  SetUnit(v string) *EpdInventoryConstituteItem
  GetUnit() *string 
}

type EpdInventoryConstituteItem struct {
  CarbonEmission *float64 `json:"carbonEmission,omitempty" xml:"carbonEmission,omitempty"`
  Factor *string `json:"factor,omitempty" xml:"factor,omitempty"`
  FactorDataset *string `json:"factorDataset,omitempty" xml:"factorDataset,omitempty"`
  FactorId *string `json:"factorId,omitempty" xml:"factorId,omitempty"`
  FactorType *int32 `json:"factorType,omitempty" xml:"factorType,omitempty"`
  FactorUnit *string `json:"factorUnit,omitempty" xml:"factorUnit,omitempty"`
  InventoryId *int64 `json:"inventoryId,omitempty" xml:"inventoryId,omitempty"`
  InventoryUnit *string `json:"inventoryUnit,omitempty" xml:"inventoryUnit,omitempty"`
  InventoryValue *float64 `json:"inventoryValue,omitempty" xml:"inventoryValue,omitempty"`
  InventoryValuePerProduct *float64 `json:"inventoryValuePerProduct,omitempty" xml:"inventoryValuePerProduct,omitempty"`
  InventoryValuePerProductUnit *string `json:"inventoryValuePerProductUnit,omitempty" xml:"inventoryValuePerProductUnit,omitempty"`
  Items []*EpdInventoryConstituteItem `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  Num *int64 `json:"num,omitempty" xml:"num,omitempty"`
  Percent *float64 `json:"percent,omitempty" xml:"percent,omitempty"`
  Quantity *float64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
  ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
  State *int32 `json:"state,omitempty" xml:"state,omitempty"`
  Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s EpdInventoryConstituteItem) String() string {
  return dara.Prettify(s)
}

func (s EpdInventoryConstituteItem) GoString() string {
  return s.String()
}

func (s *EpdInventoryConstituteItem) GetCarbonEmission() *float64  {
  return s.CarbonEmission
}

func (s *EpdInventoryConstituteItem) GetFactor() *string  {
  return s.Factor
}

func (s *EpdInventoryConstituteItem) GetFactorDataset() *string  {
  return s.FactorDataset
}

func (s *EpdInventoryConstituteItem) GetFactorId() *string  {
  return s.FactorId
}

func (s *EpdInventoryConstituteItem) GetFactorType() *int32  {
  return s.FactorType
}

func (s *EpdInventoryConstituteItem) GetFactorUnit() *string  {
  return s.FactorUnit
}

func (s *EpdInventoryConstituteItem) GetInventoryId() *int64  {
  return s.InventoryId
}

func (s *EpdInventoryConstituteItem) GetInventoryUnit() *string  {
  return s.InventoryUnit
}

func (s *EpdInventoryConstituteItem) GetInventoryValue() *float64  {
  return s.InventoryValue
}

func (s *EpdInventoryConstituteItem) GetInventoryValuePerProduct() *float64  {
  return s.InventoryValuePerProduct
}

func (s *EpdInventoryConstituteItem) GetInventoryValuePerProductUnit() *string  {
  return s.InventoryValuePerProductUnit
}

func (s *EpdInventoryConstituteItem) GetItems() []*EpdInventoryConstituteItem  {
  return s.Items
}

func (s *EpdInventoryConstituteItem) GetName() *string  {
  return s.Name
}

func (s *EpdInventoryConstituteItem) GetNum() *int64  {
  return s.Num
}

func (s *EpdInventoryConstituteItem) GetPercent() *float64  {
  return s.Percent
}

func (s *EpdInventoryConstituteItem) GetQuantity() *float64  {
  return s.Quantity
}

func (s *EpdInventoryConstituteItem) GetResourceType() *string  {
  return s.ResourceType
}

func (s *EpdInventoryConstituteItem) GetState() *int32  {
  return s.State
}

func (s *EpdInventoryConstituteItem) GetUnit() *string  {
  return s.Unit
}

func (s *EpdInventoryConstituteItem) SetCarbonEmission(v float64) *EpdInventoryConstituteItem {
  s.CarbonEmission = &v
  return s
}

func (s *EpdInventoryConstituteItem) SetFactor(v string) *EpdInventoryConstituteItem {
  s.Factor = &v
  return s
}

func (s *EpdInventoryConstituteItem) SetFactorDataset(v string) *EpdInventoryConstituteItem {
  s.FactorDataset = &v
  return s
}

func (s *EpdInventoryConstituteItem) SetFactorId(v string) *EpdInventoryConstituteItem {
  s.FactorId = &v
  return s
}

func (s *EpdInventoryConstituteItem) SetFactorType(v int32) *EpdInventoryConstituteItem {
  s.FactorType = &v
  return s
}

func (s *EpdInventoryConstituteItem) SetFactorUnit(v string) *EpdInventoryConstituteItem {
  s.FactorUnit = &v
  return s
}

func (s *EpdInventoryConstituteItem) SetInventoryId(v int64) *EpdInventoryConstituteItem {
  s.InventoryId = &v
  return s
}

func (s *EpdInventoryConstituteItem) SetInventoryUnit(v string) *EpdInventoryConstituteItem {
  s.InventoryUnit = &v
  return s
}

func (s *EpdInventoryConstituteItem) SetInventoryValue(v float64) *EpdInventoryConstituteItem {
  s.InventoryValue = &v
  return s
}

func (s *EpdInventoryConstituteItem) SetInventoryValuePerProduct(v float64) *EpdInventoryConstituteItem {
  s.InventoryValuePerProduct = &v
  return s
}

func (s *EpdInventoryConstituteItem) SetInventoryValuePerProductUnit(v string) *EpdInventoryConstituteItem {
  s.InventoryValuePerProductUnit = &v
  return s
}

func (s *EpdInventoryConstituteItem) SetItems(v []*EpdInventoryConstituteItem) *EpdInventoryConstituteItem {
  s.Items = v
  return s
}

func (s *EpdInventoryConstituteItem) SetName(v string) *EpdInventoryConstituteItem {
  s.Name = &v
  return s
}

func (s *EpdInventoryConstituteItem) SetNum(v int64) *EpdInventoryConstituteItem {
  s.Num = &v
  return s
}

func (s *EpdInventoryConstituteItem) SetPercent(v float64) *EpdInventoryConstituteItem {
  s.Percent = &v
  return s
}

func (s *EpdInventoryConstituteItem) SetQuantity(v float64) *EpdInventoryConstituteItem {
  s.Quantity = &v
  return s
}

func (s *EpdInventoryConstituteItem) SetResourceType(v string) *EpdInventoryConstituteItem {
  s.ResourceType = &v
  return s
}

func (s *EpdInventoryConstituteItem) SetState(v int32) *EpdInventoryConstituteItem {
  s.State = &v
  return s
}

func (s *EpdInventoryConstituteItem) SetUnit(v string) *EpdInventoryConstituteItem {
  s.Unit = &v
  return s
}

func (s *EpdInventoryConstituteItem) Validate() error {
  return dara.Validate(s)
}

