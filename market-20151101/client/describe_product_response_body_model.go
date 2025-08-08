// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeProductResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAuditFailMsg(v string) *DescribeProductResponseBody
	GetAuditFailMsg() *string
	SetAuditStatus(v string) *DescribeProductResponseBody
	GetAuditStatus() *string
	SetAuditTime(v int64) *DescribeProductResponseBody
	GetAuditTime() *int64
	SetCode(v string) *DescribeProductResponseBody
	GetCode() *string
	SetDescription(v string) *DescribeProductResponseBody
	GetDescription() *string
	SetFrontCategoryId(v int64) *DescribeProductResponseBody
	GetFrontCategoryId() *int64
	SetGmtCreated(v int64) *DescribeProductResponseBody
	GetGmtCreated() *int64
	SetGmtModified(v int64) *DescribeProductResponseBody
	GetGmtModified() *int64
	SetName(v string) *DescribeProductResponseBody
	GetName() *string
	SetPicUrl(v string) *DescribeProductResponseBody
	GetPicUrl() *string
	SetProductExtras(v *DescribeProductResponseBodyProductExtras) *DescribeProductResponseBody
	GetProductExtras() *DescribeProductResponseBodyProductExtras
	SetProductSkus(v *DescribeProductResponseBodyProductSkus) *DescribeProductResponseBody
	GetProductSkus() *DescribeProductResponseBodyProductSkus
	SetRequestId(v string) *DescribeProductResponseBody
	GetRequestId() *string
	SetScore(v float32) *DescribeProductResponseBody
	GetScore() *float32
	SetShopInfo(v *DescribeProductResponseBodyShopInfo) *DescribeProductResponseBody
	GetShopInfo() *DescribeProductResponseBodyShopInfo
	SetShortDescription(v string) *DescribeProductResponseBody
	GetShortDescription() *string
	SetStatus(v string) *DescribeProductResponseBody
	GetStatus() *string
	SetSupplierPk(v int64) *DescribeProductResponseBody
	GetSupplierPk() *int64
	SetType(v string) *DescribeProductResponseBody
	GetType() *string
	SetUseCount(v int64) *DescribeProductResponseBody
	GetUseCount() *int64
}

type DescribeProductResponseBody struct {
	AuditFailMsg *string `json:"AuditFailMsg,omitempty" xml:"AuditFailMsg,omitempty"`
	// example:
	//
	// function_fail
	AuditStatus *string `json:"AuditStatus,omitempty" xml:"AuditStatus,omitempty"`
	// example:
	//
	// 1581609600000
	AuditTime *int64 `json:"AuditTime,omitempty" xml:"AuditTime,omitempty"`
	// example:
	//
	// cmjj01**45
	Code        *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 523617212
	FrontCategoryId *int64 `json:"FrontCategoryId,omitempty" xml:"FrontCategoryId,omitempty"`
	// example:
	//
	// 1578931200000
	GmtCreated *int64 `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	// example:
	//
	// 1578931200000
	GmtModified *int64  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// https://oss.aliyuncs.com/photogallery/photo/1930532890589852/6245/495d5f19-03e4-4c2e-9c4e-bef9ab6af1e1.png
	PicUrl        *string                                   `json:"PicUrl,omitempty" xml:"PicUrl,omitempty"`
	ProductExtras *DescribeProductResponseBodyProductExtras `json:"ProductExtras,omitempty" xml:"ProductExtras,omitempty" type:"Struct"`
	ProductSkus   *DescribeProductResponseBodyProductSkus   `json:"ProductSkus,omitempty" xml:"ProductSkus,omitempty" type:"Struct"`
	// example:
	//
	// 6EF60BEC-0242-43AF-BB20-270359FB54A7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 5.0
	Score            *float32                             `json:"Score,omitempty" xml:"Score,omitempty"`
	ShopInfo         *DescribeProductResponseBodyShopInfo `json:"ShopInfo,omitempty" xml:"ShopInfo,omitempty" type:"Struct"`
	ShortDescription *string                              `json:"ShortDescription,omitempty" xml:"ShortDescription,omitempty"`
	// example:
	//
	// 1
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 1526111111****
	SupplierPk *int64 `json:"SupplierPk,omitempty" xml:"SupplierPk,omitempty"`
	// example:
	//
	// MIRROR
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// 10
	UseCount *int64 `json:"UseCount,omitempty" xml:"UseCount,omitempty"`
}

func (s DescribeProductResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeProductResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeProductResponseBody) GetAuditFailMsg() *string {
	return s.AuditFailMsg
}

func (s *DescribeProductResponseBody) GetAuditStatus() *string {
	return s.AuditStatus
}

func (s *DescribeProductResponseBody) GetAuditTime() *int64 {
	return s.AuditTime
}

func (s *DescribeProductResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeProductResponseBody) GetDescription() *string {
	return s.Description
}

func (s *DescribeProductResponseBody) GetFrontCategoryId() *int64 {
	return s.FrontCategoryId
}

func (s *DescribeProductResponseBody) GetGmtCreated() *int64 {
	return s.GmtCreated
}

func (s *DescribeProductResponseBody) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *DescribeProductResponseBody) GetName() *string {
	return s.Name
}

func (s *DescribeProductResponseBody) GetPicUrl() *string {
	return s.PicUrl
}

func (s *DescribeProductResponseBody) GetProductExtras() *DescribeProductResponseBodyProductExtras {
	return s.ProductExtras
}

func (s *DescribeProductResponseBody) GetProductSkus() *DescribeProductResponseBodyProductSkus {
	return s.ProductSkus
}

func (s *DescribeProductResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeProductResponseBody) GetScore() *float32 {
	return s.Score
}

func (s *DescribeProductResponseBody) GetShopInfo() *DescribeProductResponseBodyShopInfo {
	return s.ShopInfo
}

func (s *DescribeProductResponseBody) GetShortDescription() *string {
	return s.ShortDescription
}

func (s *DescribeProductResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeProductResponseBody) GetSupplierPk() *int64 {
	return s.SupplierPk
}

func (s *DescribeProductResponseBody) GetType() *string {
	return s.Type
}

func (s *DescribeProductResponseBody) GetUseCount() *int64 {
	return s.UseCount
}

func (s *DescribeProductResponseBody) SetAuditFailMsg(v string) *DescribeProductResponseBody {
	s.AuditFailMsg = &v
	return s
}

func (s *DescribeProductResponseBody) SetAuditStatus(v string) *DescribeProductResponseBody {
	s.AuditStatus = &v
	return s
}

func (s *DescribeProductResponseBody) SetAuditTime(v int64) *DescribeProductResponseBody {
	s.AuditTime = &v
	return s
}

func (s *DescribeProductResponseBody) SetCode(v string) *DescribeProductResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeProductResponseBody) SetDescription(v string) *DescribeProductResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeProductResponseBody) SetFrontCategoryId(v int64) *DescribeProductResponseBody {
	s.FrontCategoryId = &v
	return s
}

func (s *DescribeProductResponseBody) SetGmtCreated(v int64) *DescribeProductResponseBody {
	s.GmtCreated = &v
	return s
}

func (s *DescribeProductResponseBody) SetGmtModified(v int64) *DescribeProductResponseBody {
	s.GmtModified = &v
	return s
}

func (s *DescribeProductResponseBody) SetName(v string) *DescribeProductResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeProductResponseBody) SetPicUrl(v string) *DescribeProductResponseBody {
	s.PicUrl = &v
	return s
}

func (s *DescribeProductResponseBody) SetProductExtras(v *DescribeProductResponseBodyProductExtras) *DescribeProductResponseBody {
	s.ProductExtras = v
	return s
}

func (s *DescribeProductResponseBody) SetProductSkus(v *DescribeProductResponseBodyProductSkus) *DescribeProductResponseBody {
	s.ProductSkus = v
	return s
}

func (s *DescribeProductResponseBody) SetRequestId(v string) *DescribeProductResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeProductResponseBody) SetScore(v float32) *DescribeProductResponseBody {
	s.Score = &v
	return s
}

func (s *DescribeProductResponseBody) SetShopInfo(v *DescribeProductResponseBodyShopInfo) *DescribeProductResponseBody {
	s.ShopInfo = v
	return s
}

func (s *DescribeProductResponseBody) SetShortDescription(v string) *DescribeProductResponseBody {
	s.ShortDescription = &v
	return s
}

func (s *DescribeProductResponseBody) SetStatus(v string) *DescribeProductResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeProductResponseBody) SetSupplierPk(v int64) *DescribeProductResponseBody {
	s.SupplierPk = &v
	return s
}

func (s *DescribeProductResponseBody) SetType(v string) *DescribeProductResponseBody {
	s.Type = &v
	return s
}

func (s *DescribeProductResponseBody) SetUseCount(v int64) *DescribeProductResponseBody {
	s.UseCount = &v
	return s
}

func (s *DescribeProductResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeProductResponseBodyProductExtras struct {
	ProductExtra []*DescribeProductResponseBodyProductExtrasProductExtra `json:"ProductExtra,omitempty" xml:"ProductExtra,omitempty" type:"Repeated"`
}

func (s DescribeProductResponseBodyProductExtras) String() string {
	return dara.Prettify(s)
}

func (s DescribeProductResponseBodyProductExtras) GoString() string {
	return s.String()
}

func (s *DescribeProductResponseBodyProductExtras) GetProductExtra() []*DescribeProductResponseBodyProductExtrasProductExtra {
	return s.ProductExtra
}

func (s *DescribeProductResponseBodyProductExtras) SetProductExtra(v []*DescribeProductResponseBodyProductExtrasProductExtra) *DescribeProductResponseBodyProductExtras {
	s.ProductExtra = v
	return s
}

func (s *DescribeProductResponseBodyProductExtras) Validate() error {
	return dara.Validate(s)
}

type DescribeProductResponseBodyProductExtrasProductExtra struct {
	// example:
	//
	// product_advantage
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// example:
	//
	// 0
	Order *int32 `json:"Order,omitempty" xml:"Order,omitempty"`
	// example:
	//
	// HTML
	Type   *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Values *string `json:"Values,omitempty" xml:"Values,omitempty"`
}

func (s DescribeProductResponseBodyProductExtrasProductExtra) String() string {
	return dara.Prettify(s)
}

func (s DescribeProductResponseBodyProductExtrasProductExtra) GoString() string {
	return s.String()
}

func (s *DescribeProductResponseBodyProductExtrasProductExtra) GetKey() *string {
	return s.Key
}

func (s *DescribeProductResponseBodyProductExtrasProductExtra) GetLabel() *string {
	return s.Label
}

func (s *DescribeProductResponseBodyProductExtrasProductExtra) GetOrder() *int32 {
	return s.Order
}

func (s *DescribeProductResponseBodyProductExtrasProductExtra) GetType() *string {
	return s.Type
}

func (s *DescribeProductResponseBodyProductExtrasProductExtra) GetValues() *string {
	return s.Values
}

func (s *DescribeProductResponseBodyProductExtrasProductExtra) SetKey(v string) *DescribeProductResponseBodyProductExtrasProductExtra {
	s.Key = &v
	return s
}

func (s *DescribeProductResponseBodyProductExtrasProductExtra) SetLabel(v string) *DescribeProductResponseBodyProductExtrasProductExtra {
	s.Label = &v
	return s
}

func (s *DescribeProductResponseBodyProductExtrasProductExtra) SetOrder(v int32) *DescribeProductResponseBodyProductExtrasProductExtra {
	s.Order = &v
	return s
}

func (s *DescribeProductResponseBodyProductExtrasProductExtra) SetType(v string) *DescribeProductResponseBodyProductExtrasProductExtra {
	s.Type = &v
	return s
}

func (s *DescribeProductResponseBodyProductExtrasProductExtra) SetValues(v string) *DescribeProductResponseBodyProductExtrasProductExtra {
	s.Values = &v
	return s
}

func (s *DescribeProductResponseBodyProductExtrasProductExtra) Validate() error {
	return dara.Validate(s)
}

type DescribeProductResponseBodyProductSkus struct {
	ProductSku []*DescribeProductResponseBodyProductSkusProductSku `json:"ProductSku,omitempty" xml:"ProductSku,omitempty" type:"Repeated"`
}

func (s DescribeProductResponseBodyProductSkus) String() string {
	return dara.Prettify(s)
}

func (s DescribeProductResponseBodyProductSkus) GoString() string {
	return s.String()
}

func (s *DescribeProductResponseBodyProductSkus) GetProductSku() []*DescribeProductResponseBodyProductSkusProductSku {
	return s.ProductSku
}

func (s *DescribeProductResponseBodyProductSkus) SetProductSku(v []*DescribeProductResponseBodyProductSkusProductSku) *DescribeProductResponseBodyProductSkus {
	s.ProductSku = v
	return s
}

func (s *DescribeProductResponseBodyProductSkus) Validate() error {
	return dara.Validate(s)
}

type DescribeProductResponseBodyProductSkusProductSku struct {
	// example:
	//
	// PREPAY
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// example:
	//
	// cmjj01****-Package
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// {\\"img_id\\":{\\"img_version|img_region\\":{\\"V1.7|cn-shenzhen-st3-a01\\":[\\"m-wz9ho4hmos0lpxcldqoq\\"]}}
	Constraints *string `json:"Constraints,omitempty" xml:"Constraints,omitempty"`
	// example:
	//
	// true
	Hidden  *bool                                                    `json:"Hidden,omitempty" xml:"Hidden,omitempty"`
	Modules *DescribeProductResponseBodyProductSkusProductSkuModules `json:"Modules,omitempty" xml:"Modules,omitempty" type:"Struct"`
	// example:
	//
	// 21
	Name         *string                                                       `json:"Name,omitempty" xml:"Name,omitempty"`
	OrderPeriods *DescribeProductResponseBodyProductSkusProductSkuOrderPeriods `json:"OrderPeriods,omitempty" xml:"OrderPeriods,omitempty" type:"Struct"`
}

func (s DescribeProductResponseBodyProductSkusProductSku) String() string {
	return dara.Prettify(s)
}

func (s DescribeProductResponseBodyProductSkusProductSku) GoString() string {
	return s.String()
}

func (s *DescribeProductResponseBodyProductSkusProductSku) GetChargeType() *string {
	return s.ChargeType
}

func (s *DescribeProductResponseBodyProductSkusProductSku) GetCode() *string {
	return s.Code
}

func (s *DescribeProductResponseBodyProductSkusProductSku) GetConstraints() *string {
	return s.Constraints
}

func (s *DescribeProductResponseBodyProductSkusProductSku) GetHidden() *bool {
	return s.Hidden
}

func (s *DescribeProductResponseBodyProductSkusProductSku) GetModules() *DescribeProductResponseBodyProductSkusProductSkuModules {
	return s.Modules
}

func (s *DescribeProductResponseBodyProductSkusProductSku) GetName() *string {
	return s.Name
}

func (s *DescribeProductResponseBodyProductSkusProductSku) GetOrderPeriods() *DescribeProductResponseBodyProductSkusProductSkuOrderPeriods {
	return s.OrderPeriods
}

func (s *DescribeProductResponseBodyProductSkusProductSku) SetChargeType(v string) *DescribeProductResponseBodyProductSkusProductSku {
	s.ChargeType = &v
	return s
}

func (s *DescribeProductResponseBodyProductSkusProductSku) SetCode(v string) *DescribeProductResponseBodyProductSkusProductSku {
	s.Code = &v
	return s
}

func (s *DescribeProductResponseBodyProductSkusProductSku) SetConstraints(v string) *DescribeProductResponseBodyProductSkusProductSku {
	s.Constraints = &v
	return s
}

func (s *DescribeProductResponseBodyProductSkusProductSku) SetHidden(v bool) *DescribeProductResponseBodyProductSkusProductSku {
	s.Hidden = &v
	return s
}

func (s *DescribeProductResponseBodyProductSkusProductSku) SetModules(v *DescribeProductResponseBodyProductSkusProductSkuModules) *DescribeProductResponseBodyProductSkusProductSku {
	s.Modules = v
	return s
}

func (s *DescribeProductResponseBodyProductSkusProductSku) SetName(v string) *DescribeProductResponseBodyProductSkusProductSku {
	s.Name = &v
	return s
}

func (s *DescribeProductResponseBodyProductSkusProductSku) SetOrderPeriods(v *DescribeProductResponseBodyProductSkusProductSkuOrderPeriods) *DescribeProductResponseBodyProductSkusProductSku {
	s.OrderPeriods = v
	return s
}

func (s *DescribeProductResponseBodyProductSkusProductSku) Validate() error {
	return dara.Validate(s)
}

type DescribeProductResponseBodyProductSkusProductSkuModules struct {
	Module []*DescribeProductResponseBodyProductSkusProductSkuModulesModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Repeated"`
}

func (s DescribeProductResponseBodyProductSkusProductSkuModules) String() string {
	return dara.Prettify(s)
}

func (s DescribeProductResponseBodyProductSkusProductSkuModules) GoString() string {
	return s.String()
}

func (s *DescribeProductResponseBodyProductSkusProductSkuModules) GetModule() []*DescribeProductResponseBodyProductSkusProductSkuModulesModule {
	return s.Module
}

func (s *DescribeProductResponseBodyProductSkusProductSkuModules) SetModule(v []*DescribeProductResponseBodyProductSkusProductSkuModulesModule) *DescribeProductResponseBodyProductSkusProductSkuModules {
	s.Module = v
	return s
}

func (s *DescribeProductResponseBodyProductSkusProductSkuModules) Validate() error {
	return dara.Validate(s)
}

type DescribeProductResponseBodyProductSkusProductSkuModulesModule struct {
	// example:
	//
	// img_id
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 248220
	Id         *string                                                                  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name       *string                                                                  `json:"Name,omitempty" xml:"Name,omitempty"`
	Properties *DescribeProductResponseBodyProductSkusProductSkuModulesModuleProperties `json:"Properties,omitempty" xml:"Properties,omitempty" type:"Struct"`
}

func (s DescribeProductResponseBodyProductSkusProductSkuModulesModule) String() string {
	return dara.Prettify(s)
}

func (s DescribeProductResponseBodyProductSkusProductSkuModulesModule) GoString() string {
	return s.String()
}

func (s *DescribeProductResponseBodyProductSkusProductSkuModulesModule) GetCode() *string {
	return s.Code
}

func (s *DescribeProductResponseBodyProductSkusProductSkuModulesModule) GetId() *string {
	return s.Id
}

func (s *DescribeProductResponseBodyProductSkusProductSkuModulesModule) GetName() *string {
	return s.Name
}

func (s *DescribeProductResponseBodyProductSkusProductSkuModulesModule) GetProperties() *DescribeProductResponseBodyProductSkusProductSkuModulesModuleProperties {
	return s.Properties
}

func (s *DescribeProductResponseBodyProductSkusProductSkuModulesModule) SetCode(v string) *DescribeProductResponseBodyProductSkusProductSkuModulesModule {
	s.Code = &v
	return s
}

func (s *DescribeProductResponseBodyProductSkusProductSkuModulesModule) SetId(v string) *DescribeProductResponseBodyProductSkusProductSkuModulesModule {
	s.Id = &v
	return s
}

func (s *DescribeProductResponseBodyProductSkusProductSkuModulesModule) SetName(v string) *DescribeProductResponseBodyProductSkusProductSkuModulesModule {
	s.Name = &v
	return s
}

func (s *DescribeProductResponseBodyProductSkusProductSkuModulesModule) SetProperties(v *DescribeProductResponseBodyProductSkusProductSkuModulesModuleProperties) *DescribeProductResponseBodyProductSkusProductSkuModulesModule {
	s.Properties = v
	return s
}

func (s *DescribeProductResponseBodyProductSkusProductSkuModulesModule) Validate() error {
	return dara.Validate(s)
}

type DescribeProductResponseBodyProductSkusProductSkuModulesModuleProperties struct {
	Property []*DescribeProductResponseBodyProductSkusProductSkuModulesModulePropertiesProperty `json:"Property,omitempty" xml:"Property,omitempty" type:"Repeated"`
}

func (s DescribeProductResponseBodyProductSkusProductSkuModulesModuleProperties) String() string {
	return dara.Prettify(s)
}

func (s DescribeProductResponseBodyProductSkusProductSkuModulesModuleProperties) GoString() string {
	return s.String()
}

func (s *DescribeProductResponseBodyProductSkusProductSkuModulesModuleProperties) GetProperty() []*DescribeProductResponseBodyProductSkusProductSkuModulesModulePropertiesProperty {
	return s.Property
}

func (s *DescribeProductResponseBodyProductSkusProductSkuModulesModuleProperties) SetProperty(v []*DescribeProductResponseBodyProductSkusProductSkuModulesModulePropertiesProperty) *DescribeProductResponseBodyProductSkusProductSkuModulesModuleProperties {
	s.Property = v
	return s
}

func (s *DescribeProductResponseBodyProductSkusProductSkuModulesModuleProperties) Validate() error {
	return dara.Validate(s)
}

type DescribeProductResponseBodyProductSkusProductSkuModulesModulePropertiesProperty struct {
	// example:
	//
	// 1
	DisplayUnit *string `json:"DisplayUnit,omitempty" xml:"DisplayUnit,omitempty"`
	// example:
	//
	// img_id
	Key            *string                                                                                        `json:"Key,omitempty" xml:"Key,omitempty"`
	Name           *string                                                                                        `json:"Name,omitempty" xml:"Name,omitempty"`
	PropertyValues *DescribeProductResponseBodyProductSkusProductSkuModulesModulePropertiesPropertyPropertyValues `json:"PropertyValues,omitempty" xml:"PropertyValues,omitempty" type:"Struct"`
	// example:
	//
	// number
	ShowType *string `json:"ShowType,omitempty" xml:"ShowType,omitempty"`
}

func (s DescribeProductResponseBodyProductSkusProductSkuModulesModulePropertiesProperty) String() string {
	return dara.Prettify(s)
}

func (s DescribeProductResponseBodyProductSkusProductSkuModulesModulePropertiesProperty) GoString() string {
	return s.String()
}

func (s *DescribeProductResponseBodyProductSkusProductSkuModulesModulePropertiesProperty) GetDisplayUnit() *string {
	return s.DisplayUnit
}

func (s *DescribeProductResponseBodyProductSkusProductSkuModulesModulePropertiesProperty) GetKey() *string {
	return s.Key
}

func (s *DescribeProductResponseBodyProductSkusProductSkuModulesModulePropertiesProperty) GetName() *string {
	return s.Name
}

func (s *DescribeProductResponseBodyProductSkusProductSkuModulesModulePropertiesProperty) GetPropertyValues() *DescribeProductResponseBodyProductSkusProductSkuModulesModulePropertiesPropertyPropertyValues {
	return s.PropertyValues
}

func (s *DescribeProductResponseBodyProductSkusProductSkuModulesModulePropertiesProperty) GetShowType() *string {
	return s.ShowType
}

func (s *DescribeProductResponseBodyProductSkusProductSkuModulesModulePropertiesProperty) SetDisplayUnit(v string) *DescribeProductResponseBodyProductSkusProductSkuModulesModulePropertiesProperty {
	s.DisplayUnit = &v
	return s
}

func (s *DescribeProductResponseBodyProductSkusProductSkuModulesModulePropertiesProperty) SetKey(v string) *DescribeProductResponseBodyProductSkusProductSkuModulesModulePropertiesProperty {
	s.Key = &v
	return s
}

func (s *DescribeProductResponseBodyProductSkusProductSkuModulesModulePropertiesProperty) SetName(v string) *DescribeProductResponseBodyProductSkusProductSkuModulesModulePropertiesProperty {
	s.Name = &v
	return s
}

func (s *DescribeProductResponseBodyProductSkusProductSkuModulesModulePropertiesProperty) SetPropertyValues(v *DescribeProductResponseBodyProductSkusProductSkuModulesModulePropertiesPropertyPropertyValues) *DescribeProductResponseBodyProductSkusProductSkuModulesModulePropertiesProperty {
	s.PropertyValues = v
	return s
}

func (s *DescribeProductResponseBodyProductSkusProductSkuModulesModulePropertiesProperty) SetShowType(v string) *DescribeProductResponseBodyProductSkusProductSkuModulesModulePropertiesProperty {
	s.ShowType = &v
	return s
}

func (s *DescribeProductResponseBodyProductSkusProductSkuModulesModulePropertiesProperty) Validate() error {
	return dara.Validate(s)
}

type DescribeProductResponseBodyProductSkusProductSkuModulesModulePropertiesPropertyPropertyValues struct {
	PropertyValue []*DescribeProductResponseBodyProductSkusProductSkuModulesModulePropertiesPropertyPropertyValuesPropertyValue `json:"PropertyValue,omitempty" xml:"PropertyValue,omitempty" type:"Repeated"`
}

func (s DescribeProductResponseBodyProductSkusProductSkuModulesModulePropertiesPropertyPropertyValues) String() string {
	return dara.Prettify(s)
}

func (s DescribeProductResponseBodyProductSkusProductSkuModulesModulePropertiesPropertyPropertyValues) GoString() string {
	return s.String()
}

func (s *DescribeProductResponseBodyProductSkusProductSkuModulesModulePropertiesPropertyPropertyValues) GetPropertyValue() []*DescribeProductResponseBodyProductSkusProductSkuModulesModulePropertiesPropertyPropertyValuesPropertyValue {
	return s.PropertyValue
}

func (s *DescribeProductResponseBodyProductSkusProductSkuModulesModulePropertiesPropertyPropertyValues) SetPropertyValue(v []*DescribeProductResponseBodyProductSkusProductSkuModulesModulePropertiesPropertyPropertyValuesPropertyValue) *DescribeProductResponseBodyProductSkusProductSkuModulesModulePropertiesPropertyPropertyValues {
	s.PropertyValue = v
	return s
}

func (s *DescribeProductResponseBodyProductSkusProductSkuModulesModulePropertiesPropertyPropertyValues) Validate() error {
	return dara.Validate(s)
}

type DescribeProductResponseBodyProductSkusProductSkuModulesModulePropertiesPropertyPropertyValuesPropertyValue struct {
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// 123
	Max *string `json:"Max,omitempty" xml:"Max,omitempty"`
	// example:
	//
	// 11
	Min *string `json:"Min,omitempty" xml:"Min,omitempty"`
	// example:
	//
	// abc
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// example:
	//
	// 1
	Step *string `json:"Step,omitempty" xml:"Step,omitempty"`
	// example:
	//
	// single_string
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// m-28e213e7t
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeProductResponseBodyProductSkusProductSkuModulesModulePropertiesPropertyPropertyValuesPropertyValue) String() string {
	return dara.Prettify(s)
}

func (s DescribeProductResponseBodyProductSkusProductSkuModulesModulePropertiesPropertyPropertyValuesPropertyValue) GoString() string {
	return s.String()
}

func (s *DescribeProductResponseBodyProductSkusProductSkuModulesModulePropertiesPropertyPropertyValuesPropertyValue) GetDisplayName() *string {
	return s.DisplayName
}

func (s *DescribeProductResponseBodyProductSkusProductSkuModulesModulePropertiesPropertyPropertyValuesPropertyValue) GetMax() *string {
	return s.Max
}

func (s *DescribeProductResponseBodyProductSkusProductSkuModulesModulePropertiesPropertyPropertyValuesPropertyValue) GetMin() *string {
	return s.Min
}

func (s *DescribeProductResponseBodyProductSkusProductSkuModulesModulePropertiesPropertyPropertyValuesPropertyValue) GetRemark() *string {
	return s.Remark
}

func (s *DescribeProductResponseBodyProductSkusProductSkuModulesModulePropertiesPropertyPropertyValuesPropertyValue) GetStep() *string {
	return s.Step
}

func (s *DescribeProductResponseBodyProductSkusProductSkuModulesModulePropertiesPropertyPropertyValuesPropertyValue) GetType() *string {
	return s.Type
}

func (s *DescribeProductResponseBodyProductSkusProductSkuModulesModulePropertiesPropertyPropertyValuesPropertyValue) GetValue() *string {
	return s.Value
}

func (s *DescribeProductResponseBodyProductSkusProductSkuModulesModulePropertiesPropertyPropertyValuesPropertyValue) SetDisplayName(v string) *DescribeProductResponseBodyProductSkusProductSkuModulesModulePropertiesPropertyPropertyValuesPropertyValue {
	s.DisplayName = &v
	return s
}

func (s *DescribeProductResponseBodyProductSkusProductSkuModulesModulePropertiesPropertyPropertyValuesPropertyValue) SetMax(v string) *DescribeProductResponseBodyProductSkusProductSkuModulesModulePropertiesPropertyPropertyValuesPropertyValue {
	s.Max = &v
	return s
}

func (s *DescribeProductResponseBodyProductSkusProductSkuModulesModulePropertiesPropertyPropertyValuesPropertyValue) SetMin(v string) *DescribeProductResponseBodyProductSkusProductSkuModulesModulePropertiesPropertyPropertyValuesPropertyValue {
	s.Min = &v
	return s
}

func (s *DescribeProductResponseBodyProductSkusProductSkuModulesModulePropertiesPropertyPropertyValuesPropertyValue) SetRemark(v string) *DescribeProductResponseBodyProductSkusProductSkuModulesModulePropertiesPropertyPropertyValuesPropertyValue {
	s.Remark = &v
	return s
}

func (s *DescribeProductResponseBodyProductSkusProductSkuModulesModulePropertiesPropertyPropertyValuesPropertyValue) SetStep(v string) *DescribeProductResponseBodyProductSkusProductSkuModulesModulePropertiesPropertyPropertyValuesPropertyValue {
	s.Step = &v
	return s
}

func (s *DescribeProductResponseBodyProductSkusProductSkuModulesModulePropertiesPropertyPropertyValuesPropertyValue) SetType(v string) *DescribeProductResponseBodyProductSkusProductSkuModulesModulePropertiesPropertyPropertyValuesPropertyValue {
	s.Type = &v
	return s
}

func (s *DescribeProductResponseBodyProductSkusProductSkuModulesModulePropertiesPropertyPropertyValuesPropertyValue) SetValue(v string) *DescribeProductResponseBodyProductSkusProductSkuModulesModulePropertiesPropertyPropertyValuesPropertyValue {
	s.Value = &v
	return s
}

func (s *DescribeProductResponseBodyProductSkusProductSkuModulesModulePropertiesPropertyPropertyValuesPropertyValue) Validate() error {
	return dara.Validate(s)
}

type DescribeProductResponseBodyProductSkusProductSkuOrderPeriods struct {
	OrderPeriod []*DescribeProductResponseBodyProductSkusProductSkuOrderPeriodsOrderPeriod `json:"OrderPeriod,omitempty" xml:"OrderPeriod,omitempty" type:"Repeated"`
}

func (s DescribeProductResponseBodyProductSkusProductSkuOrderPeriods) String() string {
	return dara.Prettify(s)
}

func (s DescribeProductResponseBodyProductSkusProductSkuOrderPeriods) GoString() string {
	return s.String()
}

func (s *DescribeProductResponseBodyProductSkusProductSkuOrderPeriods) GetOrderPeriod() []*DescribeProductResponseBodyProductSkusProductSkuOrderPeriodsOrderPeriod {
	return s.OrderPeriod
}

func (s *DescribeProductResponseBodyProductSkusProductSkuOrderPeriods) SetOrderPeriod(v []*DescribeProductResponseBodyProductSkusProductSkuOrderPeriodsOrderPeriod) *DescribeProductResponseBodyProductSkusProductSkuOrderPeriods {
	s.OrderPeriod = v
	return s
}

func (s *DescribeProductResponseBodyProductSkusProductSkuOrderPeriods) Validate() error {
	return dara.Validate(s)
}

type DescribeProductResponseBodyProductSkusProductSkuOrderPeriodsOrderPeriod struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// HOUR
	PeriodType *string `json:"PeriodType,omitempty" xml:"PeriodType,omitempty"`
}

func (s DescribeProductResponseBodyProductSkusProductSkuOrderPeriodsOrderPeriod) String() string {
	return dara.Prettify(s)
}

func (s DescribeProductResponseBodyProductSkusProductSkuOrderPeriodsOrderPeriod) GoString() string {
	return s.String()
}

func (s *DescribeProductResponseBodyProductSkusProductSkuOrderPeriodsOrderPeriod) GetName() *string {
	return s.Name
}

func (s *DescribeProductResponseBodyProductSkusProductSkuOrderPeriodsOrderPeriod) GetPeriodType() *string {
	return s.PeriodType
}

func (s *DescribeProductResponseBodyProductSkusProductSkuOrderPeriodsOrderPeriod) SetName(v string) *DescribeProductResponseBodyProductSkusProductSkuOrderPeriodsOrderPeriod {
	s.Name = &v
	return s
}

func (s *DescribeProductResponseBodyProductSkusProductSkuOrderPeriodsOrderPeriod) SetPeriodType(v string) *DescribeProductResponseBodyProductSkusProductSkuOrderPeriodsOrderPeriod {
	s.PeriodType = &v
	return s
}

func (s *DescribeProductResponseBodyProductSkusProductSkuOrderPeriodsOrderPeriod) Validate() error {
	return dara.Validate(s)
}

type DescribeProductResponseBodyShopInfo struct {
	// example:
	//
	// 46**41@example.com
	Emails *string `json:"Emails,omitempty" xml:"Emails,omitempty"`
	// example:
	//
	// 123
	Id         *int64                                         `json:"Id,omitempty" xml:"Id,omitempty"`
	Name       *string                                        `json:"Name,omitempty" xml:"Name,omitempty"`
	Telephones *DescribeProductResponseBodyShopInfoTelephones `json:"Telephones,omitempty" xml:"Telephones,omitempty" type:"Struct"`
	WangWangs  *DescribeProductResponseBodyShopInfoWangWangs  `json:"WangWangs,omitempty" xml:"WangWangs,omitempty" type:"Struct"`
}

func (s DescribeProductResponseBodyShopInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeProductResponseBodyShopInfo) GoString() string {
	return s.String()
}

func (s *DescribeProductResponseBodyShopInfo) GetEmails() *string {
	return s.Emails
}

func (s *DescribeProductResponseBodyShopInfo) GetId() *int64 {
	return s.Id
}

func (s *DescribeProductResponseBodyShopInfo) GetName() *string {
	return s.Name
}

func (s *DescribeProductResponseBodyShopInfo) GetTelephones() *DescribeProductResponseBodyShopInfoTelephones {
	return s.Telephones
}

func (s *DescribeProductResponseBodyShopInfo) GetWangWangs() *DescribeProductResponseBodyShopInfoWangWangs {
	return s.WangWangs
}

func (s *DescribeProductResponseBodyShopInfo) SetEmails(v string) *DescribeProductResponseBodyShopInfo {
	s.Emails = &v
	return s
}

func (s *DescribeProductResponseBodyShopInfo) SetId(v int64) *DescribeProductResponseBodyShopInfo {
	s.Id = &v
	return s
}

func (s *DescribeProductResponseBodyShopInfo) SetName(v string) *DescribeProductResponseBodyShopInfo {
	s.Name = &v
	return s
}

func (s *DescribeProductResponseBodyShopInfo) SetTelephones(v *DescribeProductResponseBodyShopInfoTelephones) *DescribeProductResponseBodyShopInfo {
	s.Telephones = v
	return s
}

func (s *DescribeProductResponseBodyShopInfo) SetWangWangs(v *DescribeProductResponseBodyShopInfoWangWangs) *DescribeProductResponseBodyShopInfo {
	s.WangWangs = v
	return s
}

func (s *DescribeProductResponseBodyShopInfo) Validate() error {
	return dara.Validate(s)
}

type DescribeProductResponseBodyShopInfoTelephones struct {
	Telephone []*string `json:"Telephone,omitempty" xml:"Telephone,omitempty" type:"Repeated"`
}

func (s DescribeProductResponseBodyShopInfoTelephones) String() string {
	return dara.Prettify(s)
}

func (s DescribeProductResponseBodyShopInfoTelephones) GoString() string {
	return s.String()
}

func (s *DescribeProductResponseBodyShopInfoTelephones) GetTelephone() []*string {
	return s.Telephone
}

func (s *DescribeProductResponseBodyShopInfoTelephones) SetTelephone(v []*string) *DescribeProductResponseBodyShopInfoTelephones {
	s.Telephone = v
	return s
}

func (s *DescribeProductResponseBodyShopInfoTelephones) Validate() error {
	return dara.Validate(s)
}

type DescribeProductResponseBodyShopInfoWangWangs struct {
	WangWang []*DescribeProductResponseBodyShopInfoWangWangsWangWang `json:"WangWang,omitempty" xml:"WangWang,omitempty" type:"Repeated"`
}

func (s DescribeProductResponseBodyShopInfoWangWangs) String() string {
	return dara.Prettify(s)
}

func (s DescribeProductResponseBodyShopInfoWangWangs) GoString() string {
	return s.String()
}

func (s *DescribeProductResponseBodyShopInfoWangWangs) GetWangWang() []*DescribeProductResponseBodyShopInfoWangWangsWangWang {
	return s.WangWang
}

func (s *DescribeProductResponseBodyShopInfoWangWangs) SetWangWang(v []*DescribeProductResponseBodyShopInfoWangWangsWangWang) *DescribeProductResponseBodyShopInfoWangWangs {
	s.WangWang = v
	return s
}

func (s *DescribeProductResponseBodyShopInfoWangWangs) Validate() error {
	return dara.Validate(s)
}

type DescribeProductResponseBodyShopInfoWangWangsWangWang struct {
	// example:
	//
	// 123
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// example:
	//
	// ABC
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DescribeProductResponseBodyShopInfoWangWangsWangWang) String() string {
	return dara.Prettify(s)
}

func (s DescribeProductResponseBodyShopInfoWangWangsWangWang) GoString() string {
	return s.String()
}

func (s *DescribeProductResponseBodyShopInfoWangWangsWangWang) GetRemark() *string {
	return s.Remark
}

func (s *DescribeProductResponseBodyShopInfoWangWangsWangWang) GetUserName() *string {
	return s.UserName
}

func (s *DescribeProductResponseBodyShopInfoWangWangsWangWang) SetRemark(v string) *DescribeProductResponseBodyShopInfoWangWangsWangWang {
	s.Remark = &v
	return s
}

func (s *DescribeProductResponseBodyShopInfoWangWangsWangWang) SetUserName(v string) *DescribeProductResponseBodyShopInfoWangWangsWangWang {
	s.UserName = &v
	return s
}

func (s *DescribeProductResponseBodyShopInfoWangWangsWangWang) Validate() error {
	return dara.Validate(s)
}
