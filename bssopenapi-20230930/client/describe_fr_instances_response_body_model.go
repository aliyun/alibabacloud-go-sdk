// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFrInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeFrInstancesResponseBody
	GetCurrentPage() *int32
	SetData(v []*DescribeFrInstancesResponseBodyData) *DescribeFrInstancesResponseBody
	GetData() []*DescribeFrInstancesResponseBodyData
	SetPageSize(v int32) *DescribeFrInstancesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeFrInstancesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeFrInstancesResponseBody
	GetTotalCount() *int32
}

type DescribeFrInstancesResponseBody struct {
	// The current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The data list.
	Data []*DescribeFrInstancesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6000EE23-274B-4E07-A697-FF2E999520A4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of records.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeFrInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeFrInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFrInstancesResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeFrInstancesResponseBody) GetData() []*DescribeFrInstancesResponseBodyData {
	return s.Data
}

func (s *DescribeFrInstancesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeFrInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeFrInstancesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeFrInstancesResponseBody) SetCurrentPage(v int32) *DescribeFrInstancesResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeFrInstancesResponseBody) SetData(v []*DescribeFrInstancesResponseBodyData) *DescribeFrInstancesResponseBody {
	s.Data = v
	return s
}

func (s *DescribeFrInstancesResponseBody) SetPageSize(v int32) *DescribeFrInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeFrInstancesResponseBody) SetRequestId(v string) *DescribeFrInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFrInstancesResponseBody) SetTotalCount(v int32) *DescribeFrInstancesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeFrInstancesResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeFrInstancesResponseBodyData struct {
	// The account ID.
	//
	// example:
	//
	// 1990699401005016
	AccountId *int64 `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The account name.
	//
	// example:
	//
	// icloudtest
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The capacity type name.
	//
	// example:
	//
	// 总量递减型
	CapacitiyTypeName *string `json:"CapacitiyTypeName,omitempty" xml:"CapacitiyTypeName,omitempty"`
	// The capacity type.
	CapacityType *DescribeFrInstancesResponseBodyDataCapacityType `json:"CapacityType,omitempty" xml:"CapacityType,omitempty" type:"Struct"`
	// The capacity type code.
	//
	// example:
	//
	// deadlineAcc
	CapacityTypeCode *string `json:"CapacityTypeCode,omitempty" xml:"CapacityTypeCode,omitempty"`
	// The commodity.
	Commodity *DescribeFrInstancesResponseBodyDataCommodity `json:"Commodity,omitempty" xml:"Commodity,omitempty" type:"Struct"`
	// The commodity code.
	//
	// example:
	//
	// pts
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// The commodity name.
	//
	// example:
	//
	// 性能测试
	CommodityName *string `json:"CommodityName,omitempty" xml:"CommodityName,omitempty"`
	// The current capacity baseline unit.
	//
	// example:
	//
	// CU
	CurrCapacityBaseUnit *string `json:"CurrCapacityBaseUnit,omitempty" xml:"CurrCapacityBaseUnit,omitempty"`
	// The current capacity baseline value.
	//
	// example:
	//
	// 1000.000000
	CurrCapacityBaseValue *string `json:"CurrCapacityBaseValue,omitempty" xml:"CurrCapacityBaseValue,omitempty"`
	// The current capacity display unit.
	//
	// example:
	//
	// CU
	CurrCapacityViewUnit *string `json:"CurrCapacityViewUnit,omitempty" xml:"CurrCapacityViewUnit,omitempty"`
	// The current capacity display value.
	//
	// example:
	//
	// 1000.000000
	CurrCapacityViewValue *string `json:"CurrCapacityViewValue,omitempty" xml:"CurrCapacityViewValue,omitempty"`
	// The commitment cycle.
	CycleType *DescribeFrInstancesResponseBodyDataCycleType `json:"CycleType,omitempty" xml:"CycleType,omitempty" type:"Struct"`
	// The commitment cycle code.
	//
	// example:
	//
	// hour
	CycleTypeCode *string `json:"CycleTypeCode,omitempty" xml:"CycleTypeCode,omitempty"`
	// The commitment cycle name.
	//
	// example:
	//
	// 小时
	CycleTypeName *string `json:"CycleTypeName,omitempty" xml:"CycleTypeName,omitempty"`
	// The list of deductible regions.
	DeductRegions []*DescribeFrInstancesResponseBodyDataDeductRegions `json:"DeductRegions,omitempty" xml:"DeductRegions,omitempty" type:"Repeated"`
	// Indicates whether the resource plan can be exchanged.
	//
	// example:
	//
	// false
	EnableExchange *bool `json:"EnableExchange,omitempty" xml:"EnableExchange,omitempty"`
	// Indicates whether the resource plan can be renewed.
	//
	// example:
	//
	// false
	EnableRenew *bool `json:"EnableRenew,omitempty" xml:"EnableRenew,omitempty"`
	// Indicates whether the resource plan can be upgraded.
	//
	// example:
	//
	// false
	EnableUpgrade *bool `json:"EnableUpgrade,omitempty" xml:"EnableUpgrade,omitempty"`
	// The expiration time.
	//
	// example:
	//
	// 1710604800000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The commodity code for exchange.
	//
	// example:
	//
	// null
	ExchangeCommodityCode *string `json:"ExchangeCommodityCode,omitempty" xml:"ExchangeCommodityCode,omitempty"`
	// The initial capacity baseline unit.
	//
	// example:
	//
	// CU
	InitCapacityBaseUnit *string `json:"InitCapacityBaseUnit,omitempty" xml:"InitCapacityBaseUnit,omitempty"`
	// The initial capacity baseline value.
	//
	// example:
	//
	// 1000.000000
	InitCapacityBaseValue *string `json:"InitCapacityBaseValue,omitempty" xml:"InitCapacityBaseValue,omitempty"`
	// The initial capacity display unit.
	//
	// example:
	//
	// CU
	InitCapacityViewUnit *string `json:"InitCapacityViewUnit,omitempty" xml:"InitCapacityViewUnit,omitempty"`
	// The initial capacity display value.
	//
	// example:
	//
	// 1000.000000
	InitCapacityViewValue *string `json:"InitCapacityViewValue,omitempty" xml:"InitCapacityViewValue,omitempty"`
	// The instance name.
	//
	// example:
	//
	// alb_cubag*******
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The period capacity display unit.
	//
	// example:
	//
	// CU
	PeriodCapacityViewUnit *string `json:"PeriodCapacityViewUnit,omitempty" xml:"PeriodCapacityViewUnit,omitempty"`
	// The period capacity display value.
	//
	// example:
	//
	// 1000.000000
	PeriodCapacityViewValue *string `json:"PeriodCapacityViewValue,omitempty" xml:"PeriodCapacityViewValue,omitempty"`
	// The period time.
	//
	// example:
	//
	// hour
	PeriodTime *string `json:"PeriodTime,omitempty" xml:"PeriodTime,omitempty"`
	// The product.
	Product *DescribeFrInstancesResponseBodyDataProduct `json:"Product,omitempty" xml:"Product,omitempty" type:"Struct"`
	// The product code.
	//
	// example:
	//
	// pts
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The product name.
	//
	// example:
	//
	// 性能测试
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	// The purchase time.
	//
	// example:
	//
	// 1678939036000
	PurchaseTime *int64 `json:"PurchaseTime,omitempty" xml:"PurchaseTime,omitempty"`
	// The region.
	//
	// example:
	//
	// *
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The region name.
	//
	// example:
	//
	// cn-qingdao
	RegionName *string `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
	// The specification.
	//
	// example:
	//
	// *
	Spec *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
	// The effective period.
	//
	// example:
	//
	// 1678939035000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The resource status.
	Status *DescribeFrInstancesResponseBodyDataStatus `json:"Status,omitempty" xml:"Status,omitempty" type:"Struct"`
	// The resource status code.
	//
	// example:
	//
	// valid
	StatusCode *string `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	// The resource status name.
	//
	// example:
	//
	// 有效
	StatusName *string `json:"StatusName,omitempty" xml:"StatusName,omitempty"`
	// The template.
	Template *DescribeFrInstancesResponseBodyDataTemplate `json:"Template,omitempty" xml:"Template,omitempty" type:"Struct"`
	// The template code.
	//
	// example:
	//
	// FPT_armsappbag_deadlineAcc_bj
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	// The template name.
	//
	// example:
	//
	// new_test
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The period capacity baseline unit.
	//
	// example:
	//
	// CU
	PeriodCapacityBaseUnit *string `json:"periodCapacityBaseUnit,omitempty" xml:"periodCapacityBaseUnit,omitempty"`
	// The period capacity baseline value.
	//
	// example:
	//
	// 1000.000000
	PeriodCapacityBaseValue *string `json:"periodCapacityBaseValue,omitempty" xml:"periodCapacityBaseValue,omitempty"`
}

func (s DescribeFrInstancesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeFrInstancesResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeFrInstancesResponseBodyData) GetAccountId() *int64 {
	return s.AccountId
}

func (s *DescribeFrInstancesResponseBodyData) GetAccountName() *string {
	return s.AccountName
}

func (s *DescribeFrInstancesResponseBodyData) GetCapacitiyTypeName() *string {
	return s.CapacitiyTypeName
}

func (s *DescribeFrInstancesResponseBodyData) GetCapacityType() *DescribeFrInstancesResponseBodyDataCapacityType {
	return s.CapacityType
}

func (s *DescribeFrInstancesResponseBodyData) GetCapacityTypeCode() *string {
	return s.CapacityTypeCode
}

func (s *DescribeFrInstancesResponseBodyData) GetCommodity() *DescribeFrInstancesResponseBodyDataCommodity {
	return s.Commodity
}

func (s *DescribeFrInstancesResponseBodyData) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *DescribeFrInstancesResponseBodyData) GetCommodityName() *string {
	return s.CommodityName
}

func (s *DescribeFrInstancesResponseBodyData) GetCurrCapacityBaseUnit() *string {
	return s.CurrCapacityBaseUnit
}

func (s *DescribeFrInstancesResponseBodyData) GetCurrCapacityBaseValue() *string {
	return s.CurrCapacityBaseValue
}

func (s *DescribeFrInstancesResponseBodyData) GetCurrCapacityViewUnit() *string {
	return s.CurrCapacityViewUnit
}

func (s *DescribeFrInstancesResponseBodyData) GetCurrCapacityViewValue() *string {
	return s.CurrCapacityViewValue
}

func (s *DescribeFrInstancesResponseBodyData) GetCycleType() *DescribeFrInstancesResponseBodyDataCycleType {
	return s.CycleType
}

func (s *DescribeFrInstancesResponseBodyData) GetCycleTypeCode() *string {
	return s.CycleTypeCode
}

func (s *DescribeFrInstancesResponseBodyData) GetCycleTypeName() *string {
	return s.CycleTypeName
}

func (s *DescribeFrInstancesResponseBodyData) GetDeductRegions() []*DescribeFrInstancesResponseBodyDataDeductRegions {
	return s.DeductRegions
}

func (s *DescribeFrInstancesResponseBodyData) GetEnableExchange() *bool {
	return s.EnableExchange
}

func (s *DescribeFrInstancesResponseBodyData) GetEnableRenew() *bool {
	return s.EnableRenew
}

func (s *DescribeFrInstancesResponseBodyData) GetEnableUpgrade() *bool {
	return s.EnableUpgrade
}

func (s *DescribeFrInstancesResponseBodyData) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeFrInstancesResponseBodyData) GetExchangeCommodityCode() *string {
	return s.ExchangeCommodityCode
}

func (s *DescribeFrInstancesResponseBodyData) GetInitCapacityBaseUnit() *string {
	return s.InitCapacityBaseUnit
}

func (s *DescribeFrInstancesResponseBodyData) GetInitCapacityBaseValue() *string {
	return s.InitCapacityBaseValue
}

func (s *DescribeFrInstancesResponseBodyData) GetInitCapacityViewUnit() *string {
	return s.InitCapacityViewUnit
}

func (s *DescribeFrInstancesResponseBodyData) GetInitCapacityViewValue() *string {
	return s.InitCapacityViewValue
}

func (s *DescribeFrInstancesResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeFrInstancesResponseBodyData) GetPeriodCapacityViewUnit() *string {
	return s.PeriodCapacityViewUnit
}

func (s *DescribeFrInstancesResponseBodyData) GetPeriodCapacityViewValue() *string {
	return s.PeriodCapacityViewValue
}

func (s *DescribeFrInstancesResponseBodyData) GetPeriodTime() *string {
	return s.PeriodTime
}

func (s *DescribeFrInstancesResponseBodyData) GetProduct() *DescribeFrInstancesResponseBodyDataProduct {
	return s.Product
}

func (s *DescribeFrInstancesResponseBodyData) GetProductCode() *string {
	return s.ProductCode
}

func (s *DescribeFrInstancesResponseBodyData) GetProductName() *string {
	return s.ProductName
}

func (s *DescribeFrInstancesResponseBodyData) GetPurchaseTime() *int64 {
	return s.PurchaseTime
}

func (s *DescribeFrInstancesResponseBodyData) GetRegion() *string {
	return s.Region
}

func (s *DescribeFrInstancesResponseBodyData) GetRegionName() *string {
	return s.RegionName
}

func (s *DescribeFrInstancesResponseBodyData) GetSpec() *string {
	return s.Spec
}

func (s *DescribeFrInstancesResponseBodyData) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeFrInstancesResponseBodyData) GetStatus() *DescribeFrInstancesResponseBodyDataStatus {
	return s.Status
}

func (s *DescribeFrInstancesResponseBodyData) GetStatusCode() *string {
	return s.StatusCode
}

func (s *DescribeFrInstancesResponseBodyData) GetStatusName() *string {
	return s.StatusName
}

func (s *DescribeFrInstancesResponseBodyData) GetTemplate() *DescribeFrInstancesResponseBodyDataTemplate {
	return s.Template
}

func (s *DescribeFrInstancesResponseBodyData) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *DescribeFrInstancesResponseBodyData) GetTemplateName() *string {
	return s.TemplateName
}

func (s *DescribeFrInstancesResponseBodyData) GetPeriodCapacityBaseUnit() *string {
	return s.PeriodCapacityBaseUnit
}

func (s *DescribeFrInstancesResponseBodyData) GetPeriodCapacityBaseValue() *string {
	return s.PeriodCapacityBaseValue
}

func (s *DescribeFrInstancesResponseBodyData) SetAccountId(v int64) *DescribeFrInstancesResponseBodyData {
	s.AccountId = &v
	return s
}

func (s *DescribeFrInstancesResponseBodyData) SetAccountName(v string) *DescribeFrInstancesResponseBodyData {
	s.AccountName = &v
	return s
}

func (s *DescribeFrInstancesResponseBodyData) SetCapacitiyTypeName(v string) *DescribeFrInstancesResponseBodyData {
	s.CapacitiyTypeName = &v
	return s
}

func (s *DescribeFrInstancesResponseBodyData) SetCapacityType(v *DescribeFrInstancesResponseBodyDataCapacityType) *DescribeFrInstancesResponseBodyData {
	s.CapacityType = v
	return s
}

func (s *DescribeFrInstancesResponseBodyData) SetCapacityTypeCode(v string) *DescribeFrInstancesResponseBodyData {
	s.CapacityTypeCode = &v
	return s
}

func (s *DescribeFrInstancesResponseBodyData) SetCommodity(v *DescribeFrInstancesResponseBodyDataCommodity) *DescribeFrInstancesResponseBodyData {
	s.Commodity = v
	return s
}

func (s *DescribeFrInstancesResponseBodyData) SetCommodityCode(v string) *DescribeFrInstancesResponseBodyData {
	s.CommodityCode = &v
	return s
}

func (s *DescribeFrInstancesResponseBodyData) SetCommodityName(v string) *DescribeFrInstancesResponseBodyData {
	s.CommodityName = &v
	return s
}

func (s *DescribeFrInstancesResponseBodyData) SetCurrCapacityBaseUnit(v string) *DescribeFrInstancesResponseBodyData {
	s.CurrCapacityBaseUnit = &v
	return s
}

func (s *DescribeFrInstancesResponseBodyData) SetCurrCapacityBaseValue(v string) *DescribeFrInstancesResponseBodyData {
	s.CurrCapacityBaseValue = &v
	return s
}

func (s *DescribeFrInstancesResponseBodyData) SetCurrCapacityViewUnit(v string) *DescribeFrInstancesResponseBodyData {
	s.CurrCapacityViewUnit = &v
	return s
}

func (s *DescribeFrInstancesResponseBodyData) SetCurrCapacityViewValue(v string) *DescribeFrInstancesResponseBodyData {
	s.CurrCapacityViewValue = &v
	return s
}

func (s *DescribeFrInstancesResponseBodyData) SetCycleType(v *DescribeFrInstancesResponseBodyDataCycleType) *DescribeFrInstancesResponseBodyData {
	s.CycleType = v
	return s
}

func (s *DescribeFrInstancesResponseBodyData) SetCycleTypeCode(v string) *DescribeFrInstancesResponseBodyData {
	s.CycleTypeCode = &v
	return s
}

func (s *DescribeFrInstancesResponseBodyData) SetCycleTypeName(v string) *DescribeFrInstancesResponseBodyData {
	s.CycleTypeName = &v
	return s
}

func (s *DescribeFrInstancesResponseBodyData) SetDeductRegions(v []*DescribeFrInstancesResponseBodyDataDeductRegions) *DescribeFrInstancesResponseBodyData {
	s.DeductRegions = v
	return s
}

func (s *DescribeFrInstancesResponseBodyData) SetEnableExchange(v bool) *DescribeFrInstancesResponseBodyData {
	s.EnableExchange = &v
	return s
}

func (s *DescribeFrInstancesResponseBodyData) SetEnableRenew(v bool) *DescribeFrInstancesResponseBodyData {
	s.EnableRenew = &v
	return s
}

func (s *DescribeFrInstancesResponseBodyData) SetEnableUpgrade(v bool) *DescribeFrInstancesResponseBodyData {
	s.EnableUpgrade = &v
	return s
}

func (s *DescribeFrInstancesResponseBodyData) SetEndTime(v int64) *DescribeFrInstancesResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *DescribeFrInstancesResponseBodyData) SetExchangeCommodityCode(v string) *DescribeFrInstancesResponseBodyData {
	s.ExchangeCommodityCode = &v
	return s
}

func (s *DescribeFrInstancesResponseBodyData) SetInitCapacityBaseUnit(v string) *DescribeFrInstancesResponseBodyData {
	s.InitCapacityBaseUnit = &v
	return s
}

func (s *DescribeFrInstancesResponseBodyData) SetInitCapacityBaseValue(v string) *DescribeFrInstancesResponseBodyData {
	s.InitCapacityBaseValue = &v
	return s
}

func (s *DescribeFrInstancesResponseBodyData) SetInitCapacityViewUnit(v string) *DescribeFrInstancesResponseBodyData {
	s.InitCapacityViewUnit = &v
	return s
}

func (s *DescribeFrInstancesResponseBodyData) SetInitCapacityViewValue(v string) *DescribeFrInstancesResponseBodyData {
	s.InitCapacityViewValue = &v
	return s
}

func (s *DescribeFrInstancesResponseBodyData) SetInstanceId(v string) *DescribeFrInstancesResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *DescribeFrInstancesResponseBodyData) SetPeriodCapacityViewUnit(v string) *DescribeFrInstancesResponseBodyData {
	s.PeriodCapacityViewUnit = &v
	return s
}

func (s *DescribeFrInstancesResponseBodyData) SetPeriodCapacityViewValue(v string) *DescribeFrInstancesResponseBodyData {
	s.PeriodCapacityViewValue = &v
	return s
}

func (s *DescribeFrInstancesResponseBodyData) SetPeriodTime(v string) *DescribeFrInstancesResponseBodyData {
	s.PeriodTime = &v
	return s
}

func (s *DescribeFrInstancesResponseBodyData) SetProduct(v *DescribeFrInstancesResponseBodyDataProduct) *DescribeFrInstancesResponseBodyData {
	s.Product = v
	return s
}

func (s *DescribeFrInstancesResponseBodyData) SetProductCode(v string) *DescribeFrInstancesResponseBodyData {
	s.ProductCode = &v
	return s
}

func (s *DescribeFrInstancesResponseBodyData) SetProductName(v string) *DescribeFrInstancesResponseBodyData {
	s.ProductName = &v
	return s
}

func (s *DescribeFrInstancesResponseBodyData) SetPurchaseTime(v int64) *DescribeFrInstancesResponseBodyData {
	s.PurchaseTime = &v
	return s
}

func (s *DescribeFrInstancesResponseBodyData) SetRegion(v string) *DescribeFrInstancesResponseBodyData {
	s.Region = &v
	return s
}

func (s *DescribeFrInstancesResponseBodyData) SetRegionName(v string) *DescribeFrInstancesResponseBodyData {
	s.RegionName = &v
	return s
}

func (s *DescribeFrInstancesResponseBodyData) SetSpec(v string) *DescribeFrInstancesResponseBodyData {
	s.Spec = &v
	return s
}

func (s *DescribeFrInstancesResponseBodyData) SetStartTime(v int64) *DescribeFrInstancesResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *DescribeFrInstancesResponseBodyData) SetStatus(v *DescribeFrInstancesResponseBodyDataStatus) *DescribeFrInstancesResponseBodyData {
	s.Status = v
	return s
}

func (s *DescribeFrInstancesResponseBodyData) SetStatusCode(v string) *DescribeFrInstancesResponseBodyData {
	s.StatusCode = &v
	return s
}

func (s *DescribeFrInstancesResponseBodyData) SetStatusName(v string) *DescribeFrInstancesResponseBodyData {
	s.StatusName = &v
	return s
}

func (s *DescribeFrInstancesResponseBodyData) SetTemplate(v *DescribeFrInstancesResponseBodyDataTemplate) *DescribeFrInstancesResponseBodyData {
	s.Template = v
	return s
}

func (s *DescribeFrInstancesResponseBodyData) SetTemplateCode(v string) *DescribeFrInstancesResponseBodyData {
	s.TemplateCode = &v
	return s
}

func (s *DescribeFrInstancesResponseBodyData) SetTemplateName(v string) *DescribeFrInstancesResponseBodyData {
	s.TemplateName = &v
	return s
}

func (s *DescribeFrInstancesResponseBodyData) SetPeriodCapacityBaseUnit(v string) *DescribeFrInstancesResponseBodyData {
	s.PeriodCapacityBaseUnit = &v
	return s
}

func (s *DescribeFrInstancesResponseBodyData) SetPeriodCapacityBaseValue(v string) *DescribeFrInstancesResponseBodyData {
	s.PeriodCapacityBaseValue = &v
	return s
}

func (s *DescribeFrInstancesResponseBodyData) Validate() error {
	if s.CapacityType != nil {
		if err := s.CapacityType.Validate(); err != nil {
			return err
		}
	}
	if s.Commodity != nil {
		if err := s.Commodity.Validate(); err != nil {
			return err
		}
	}
	if s.CycleType != nil {
		if err := s.CycleType.Validate(); err != nil {
			return err
		}
	}
	if s.DeductRegions != nil {
		for _, item := range s.DeductRegions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Product != nil {
		if err := s.Product.Validate(); err != nil {
			return err
		}
	}
	if s.Status != nil {
		if err := s.Status.Validate(); err != nil {
			return err
		}
	}
	if s.Template != nil {
		if err := s.Template.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeFrInstancesResponseBodyDataCapacityType struct {
	// The property code.
	//
	// example:
	//
	// deadlineAcc
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The property name.
	//
	// example:
	//
	// 总量递减型
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeFrInstancesResponseBodyDataCapacityType) String() string {
	return dara.Prettify(s)
}

func (s DescribeFrInstancesResponseBodyDataCapacityType) GoString() string {
	return s.String()
}

func (s *DescribeFrInstancesResponseBodyDataCapacityType) GetCode() *string {
	return s.Code
}

func (s *DescribeFrInstancesResponseBodyDataCapacityType) GetName() *string {
	return s.Name
}

func (s *DescribeFrInstancesResponseBodyDataCapacityType) SetCode(v string) *DescribeFrInstancesResponseBodyDataCapacityType {
	s.Code = &v
	return s
}

func (s *DescribeFrInstancesResponseBodyDataCapacityType) SetName(v string) *DescribeFrInstancesResponseBodyDataCapacityType {
	s.Name = &v
	return s
}

func (s *DescribeFrInstancesResponseBodyDataCapacityType) Validate() error {
	return dara.Validate(s)
}

type DescribeFrInstancesResponseBodyDataCommodity struct {
	// The property code.
	//
	// example:
	//
	// slb_albcubag_dp_cn
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The property name.
	//
	// example:
	//
	// ALB资源包
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeFrInstancesResponseBodyDataCommodity) String() string {
	return dara.Prettify(s)
}

func (s DescribeFrInstancesResponseBodyDataCommodity) GoString() string {
	return s.String()
}

func (s *DescribeFrInstancesResponseBodyDataCommodity) GetCode() *string {
	return s.Code
}

func (s *DescribeFrInstancesResponseBodyDataCommodity) GetName() *string {
	return s.Name
}

func (s *DescribeFrInstancesResponseBodyDataCommodity) SetCode(v string) *DescribeFrInstancesResponseBodyDataCommodity {
	s.Code = &v
	return s
}

func (s *DescribeFrInstancesResponseBodyDataCommodity) SetName(v string) *DescribeFrInstancesResponseBodyDataCommodity {
	s.Name = &v
	return s
}

func (s *DescribeFrInstancesResponseBodyDataCommodity) Validate() error {
	return dara.Validate(s)
}

type DescribeFrInstancesResponseBodyDataCycleType struct {
	// The commitment cycle code.
	//
	// example:
	//
	// hour
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The commitment cycle name.
	//
	// example:
	//
	// 小时
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeFrInstancesResponseBodyDataCycleType) String() string {
	return dara.Prettify(s)
}

func (s DescribeFrInstancesResponseBodyDataCycleType) GoString() string {
	return s.String()
}

func (s *DescribeFrInstancesResponseBodyDataCycleType) GetCode() *string {
	return s.Code
}

func (s *DescribeFrInstancesResponseBodyDataCycleType) GetName() *string {
	return s.Name
}

func (s *DescribeFrInstancesResponseBodyDataCycleType) SetCode(v string) *DescribeFrInstancesResponseBodyDataCycleType {
	s.Code = &v
	return s
}

func (s *DescribeFrInstancesResponseBodyDataCycleType) SetName(v string) *DescribeFrInstancesResponseBodyDataCycleType {
	s.Name = &v
	return s
}

func (s *DescribeFrInstancesResponseBodyDataCycleType) Validate() error {
	return dara.Validate(s)
}

type DescribeFrInstancesResponseBodyDataDeductRegions struct {
	// The deductible region code.
	//
	// example:
	//
	// cn-beijing
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The deductible region.
	//
	// example:
	//
	// 北京
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeFrInstancesResponseBodyDataDeductRegions) String() string {
	return dara.Prettify(s)
}

func (s DescribeFrInstancesResponseBodyDataDeductRegions) GoString() string {
	return s.String()
}

func (s *DescribeFrInstancesResponseBodyDataDeductRegions) GetCode() *string {
	return s.Code
}

func (s *DescribeFrInstancesResponseBodyDataDeductRegions) GetName() *string {
	return s.Name
}

func (s *DescribeFrInstancesResponseBodyDataDeductRegions) SetCode(v string) *DescribeFrInstancesResponseBodyDataDeductRegions {
	s.Code = &v
	return s
}

func (s *DescribeFrInstancesResponseBodyDataDeductRegions) SetName(v string) *DescribeFrInstancesResponseBodyDataDeductRegions {
	s.Name = &v
	return s
}

func (s *DescribeFrInstancesResponseBodyDataDeductRegions) Validate() error {
	return dara.Validate(s)
}

type DescribeFrInstancesResponseBodyDataProduct struct {
	// The property code.
	//
	// example:
	//
	// slb
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The property name.
	//
	// example:
	//
	// 负载均衡
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeFrInstancesResponseBodyDataProduct) String() string {
	return dara.Prettify(s)
}

func (s DescribeFrInstancesResponseBodyDataProduct) GoString() string {
	return s.String()
}

func (s *DescribeFrInstancesResponseBodyDataProduct) GetCode() *string {
	return s.Code
}

func (s *DescribeFrInstancesResponseBodyDataProduct) GetName() *string {
	return s.Name
}

func (s *DescribeFrInstancesResponseBodyDataProduct) SetCode(v string) *DescribeFrInstancesResponseBodyDataProduct {
	s.Code = &v
	return s
}

func (s *DescribeFrInstancesResponseBodyDataProduct) SetName(v string) *DescribeFrInstancesResponseBodyDataProduct {
	s.Name = &v
	return s
}

func (s *DescribeFrInstancesResponseBodyDataProduct) Validate() error {
	return dara.Validate(s)
}

type DescribeFrInstancesResponseBodyDataStatus struct {
	// The property code.
	//
	// example:
	//
	// valid
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The property name.
	//
	// example:
	//
	// 有效
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeFrInstancesResponseBodyDataStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeFrInstancesResponseBodyDataStatus) GoString() string {
	return s.String()
}

func (s *DescribeFrInstancesResponseBodyDataStatus) GetCode() *string {
	return s.Code
}

func (s *DescribeFrInstancesResponseBodyDataStatus) GetName() *string {
	return s.Name
}

func (s *DescribeFrInstancesResponseBodyDataStatus) SetCode(v string) *DescribeFrInstancesResponseBodyDataStatus {
	s.Code = &v
	return s
}

func (s *DescribeFrInstancesResponseBodyDataStatus) SetName(v string) *DescribeFrInstancesResponseBodyDataStatus {
	s.Name = &v
	return s
}

func (s *DescribeFrInstancesResponseBodyDataStatus) Validate() error {
	return dara.Validate(s)
}

type DescribeFrInstancesResponseBodyDataTemplate struct {
	// The property code.
	//
	// example:
	//
	// alb_cubag*******
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The property name.
	//
	// example:
	//
	// 中国内地区域
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeFrInstancesResponseBodyDataTemplate) String() string {
	return dara.Prettify(s)
}

func (s DescribeFrInstancesResponseBodyDataTemplate) GoString() string {
	return s.String()
}

func (s *DescribeFrInstancesResponseBodyDataTemplate) GetCode() *string {
	return s.Code
}

func (s *DescribeFrInstancesResponseBodyDataTemplate) GetName() *string {
	return s.Name
}

func (s *DescribeFrInstancesResponseBodyDataTemplate) SetCode(v string) *DescribeFrInstancesResponseBodyDataTemplate {
	s.Code = &v
	return s
}

func (s *DescribeFrInstancesResponseBodyDataTemplate) SetName(v string) *DescribeFrInstancesResponseBodyDataTemplate {
	s.Name = &v
	return s
}

func (s *DescribeFrInstancesResponseBodyDataTemplate) Validate() error {
	return dara.Validate(s)
}
