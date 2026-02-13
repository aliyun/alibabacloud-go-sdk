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
  // Emissions are presented with 24 decimal places. It\\"s advisable to utilize the truncated version. These emissions encompass all tiers: At the first level, they reflect the product\\"s total emissions under the current environmental impact; at the second level, they show the process\\"s total emissions; and at the third level, they represent the inventory\\"s emissions within the same environmental impact context.
  // 
  // example:
  // 
  // 1009.976265540000000000000000000000
  CarbonEmission *float64 `json:"carbonEmission,omitempty" xml:"carbonEmission,omitempty"`
  // The value of the factor. Only the third level is not empty. The factor in the inventory information indicates the factor value of the selected factor. The factor value is kept to four decimal places after the decimal point, and "1.0000" indicates that the factor value is 1. The factor value should be used in combination with factorUnit. If factorUnit is "kg CO2-Eq mg/kg", it means that the carbon emission per 1kg of the list is 1kg CO2-Eq.
  // 
  // example:
  // 
  // 1.0000
  Factor *string `json:"factor,omitempty" xml:"factor,omitempty"`
  // The database to which the factorDataset factor belongs. This field is used in conjunction with factorType. If factorType is 1, the data name of the factor is displayed. If factorType is 2,factorId indicates the ID of the referenced product. If factorType is null,factorId is meaningless. This field is a new field. Old data may not have data in this field and is displayed as "/".
  // 
  // example:
  // 
  // /
  FactorDataset *string `json:"factorDataset,omitempty" xml:"factorDataset,omitempty"`
  // The id of the factor.
  // 
  // example:
  // 
  // 1234
  FactorId *string `json:"factorId,omitempty" xml:"factorId,omitempty"`
  // The key of the factor type. Only the third level is not empty. The factorType in the inventory information indicates that the factor source type is selected; the optional values are 1, 2, or null:1: factor library, 2: product library. null indicates that the factor is not selected from the factor library or product library, and the factor is constructed by the user.
  // 
  // example:
  // 
  // 1
  FactorType *int32 `json:"factorType,omitempty" xml:"factorType,omitempty"`
  // Factor Unit
  // 
  // example:
  // 
  // kg CO2-Eq/kg
  FactorUnit *string `json:"factorUnit,omitempty" xml:"factorUnit,omitempty"`
  // manifest id
  // 
  // example:
  // 
  // 1
  InventoryId *int64 `json:"inventoryId,omitempty" xml:"inventoryId,omitempty"`
  // Inventory Unit
  // 
  // example:
  // 
  // t
  InventoryUnit *string `json:"inventoryUnit,omitempty" xml:"inventoryUnit,omitempty"`
  // The quantity of the inventory. It is not empty only at the third level. The third level is the inventory details. This field indicates the id of the inventory. It should be used in conjunction with the InventoryUnit.
  // 
  // example:
  // 
  // 2.000000000000000000000000
  InventoryValue *float64 `json:"inventoryValue,omitempty" xml:"inventoryValue,omitempty"`
  // Activity data per functional unit: only the third level is not empty; the value retains 24 decimal places, indicating the number of activities per functional unit; should be used in conjunction with the inventoryValuePerProductUnit.
  // 
  // example:
  // 
  // 1.000000000000000000000000
  InventoryValuePerProduct *float64 `json:"inventoryValuePerProduct,omitempty" xml:"inventoryValuePerProduct,omitempty"`
  // Unit of activity data per functional unit. Only the third level is not empty. in the inventory information indicates the unit of activity data per functional unit.
  // 
  // example:
  // 
  // kg
  InventoryValuePerProductUnit *string `json:"inventoryValuePerProductUnit,omitempty" xml:"inventoryValuePerProductUnit,omitempty"`
  // List of children
  Items []*EpdInventoryConstituteItem `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
  // Returns the name of the current level. The names of different levels have different meanings. The first level returns the environmental impact type. The second level returns the current process name. The third level returns the current inventory name
  // 
  // example:
  // 
  // climate change
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
  // Category key, please refer to [Carbon Footprint Appendices](https://carbon-doc.oss-cn-hangzhou.aliyuncs.com/CarbonFootprintAppendices-en.pdf).
  // 
  // example:
  // 
  // 1
  Num *int64 `json:"num,omitempty" xml:"num,omitempty"`
  // The percentage. Four decimal places are retained and 31.4000 is returned to indicate the percentage 31.4%. The first level returns null; The second level returns the current process as a percentage of total emissions; the third level returns the current inventory as a percentage of total emissions.
  // 
  // example:
  // 
  // 31.4000
  Percent *float64 `json:"percent,omitempty" xml:"percent,omitempty"`
  // Raw activity data. Only the third level returns a quantity that is not null, indicating a manifest.
  // 
  // example:
  // 
  // 1.000000000000
  Quantity *float64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
  // The type of the inventory. Only the third level returns non-null, and the third level is the inventory details. This field indicates the resource type name of the inventory. You may refer to [Carbon Footprint Appendices](https://carbon-doc.oss-cn-hangzhou.aliyuncs.com/CarbonFootprintAppendices-en.pdf).
  // 
  // example:
  // 
  // Energy
  ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
  // The data status. 1 indicates accurate calculation, 2 indicates default data, 3 indicates missing factor, and 0 value is used in other cases.
  // 
  // example:
  // 
  // 1
  State *int32 `json:"state,omitempty" xml:"state,omitempty"`
  // The unit of environmental impact result value. This unit is the unit of the environmental impact result value of the corresponding environmental impact category. For example, the unit kg CO2-Eq represent the emission values shown here is kg CO2-Eq.
  // 
  // example:
  // 
  // kg CO2-Eq
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
  if s.Items != nil {
    for _, item := range s.Items {
      if item != nil {
        if err := item.Validate(); err != nil {
          return err
        }
      }
    }
  }
  return nil
}

