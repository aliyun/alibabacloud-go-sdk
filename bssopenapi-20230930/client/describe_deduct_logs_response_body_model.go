// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDeductLogsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeDeductLogsResponseBody
	GetCurrentPage() *int32
	SetData(v []*DescribeDeductLogsResponseBodyData) *DescribeDeductLogsResponseBody
	GetData() []*DescribeDeductLogsResponseBodyData
	SetPageSize(v int32) *DescribeDeductLogsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeDeductLogsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeDeductLogsResponseBody
	GetTotalCount() *int32
}

type DescribeDeductLogsResponseBody struct {
	// The current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The data list.
	Data []*DescribeDeductLogsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
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
	// DF58589C-A06C-4224-8615-7797E6474FA3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of records.
	//
	// example:
	//
	// 8
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDeductLogsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeductLogsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDeductLogsResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeDeductLogsResponseBody) GetData() []*DescribeDeductLogsResponseBodyData {
	return s.Data
}

func (s *DescribeDeductLogsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDeductLogsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDeductLogsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeDeductLogsResponseBody) SetCurrentPage(v int32) *DescribeDeductLogsResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeDeductLogsResponseBody) SetData(v []*DescribeDeductLogsResponseBodyData) *DescribeDeductLogsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeDeductLogsResponseBody) SetPageSize(v int32) *DescribeDeductLogsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDeductLogsResponseBody) SetRequestId(v string) *DescribeDeductLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDeductLogsResponseBody) SetTotalCount(v int32) *DescribeDeductLogsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDeductLogsResponseBody) Validate() error {
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

type DescribeDeductLogsResponseBodyData struct {
	// The account ID.
	//
	// example:
	//
	// 1929817951466001
	AccountId *int64 `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The account name.
	//
	// example:
	//
	// icloudtest
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The deducted commodity.
	BillingCommodity *DescribeDeductLogsResponseBodyDataBillingCommodity `json:"BillingCommodity,omitempty" xml:"BillingCommodity,omitempty" type:"Struct"`
	// The deducted commodity code.
	//
	// example:
	//
	// snapshot
	BillingCommodityCode *string `json:"BillingCommodityCode,omitempty" xml:"BillingCommodityCode,omitempty"`
	// The deducted commodity name.
	//
	// example:
	//
	// 云服务器ECS-快照
	BillingCommodityName *string `json:"BillingCommodityName,omitempty" xml:"BillingCommodityName,omitempty"`
	// The expiration time.
	//
	// example:
	//
	// 1679036400000
	BillingEndTime *int64 `json:"BillingEndTime,omitempty" xml:"BillingEndTime,omitempty"`
	// The deduction instance.
	//
	// example:
	//
	// cn-beijing
	BillingInstanceId *string `json:"BillingInstanceId,omitempty" xml:"BillingInstanceId,omitempty"`
	// The deduction billable item.
	BillingPriceField *DescribeDeductLogsResponseBodyDataBillingPriceField `json:"BillingPriceField,omitempty" xml:"BillingPriceField,omitempty" type:"Struct"`
	// The deduction billable item code.
	//
	// example:
	//
	// Storage
	BillingPriceFieldCode *string `json:"BillingPriceFieldCode,omitempty" xml:"BillingPriceFieldCode,omitempty"`
	// The deduction billable item name.
	//
	// example:
	//
	// 零折使用容量
	BillingPriceFieldName *string `json:"BillingPriceFieldName,omitempty" xml:"BillingPriceFieldName,omitempty"`
	// The effective period.
	//
	// example:
	//
	// 1679032800000
	BillingStartTime *int64 `json:"BillingStartTime,omitempty" xml:"BillingStartTime,omitempty"`
	// The display unit of the capacity after deduction.
	//
	// example:
	//
	// GB
	CapacityAfterDeductViewUnit *string `json:"CapacityAfterDeductViewUnit,omitempty" xml:"CapacityAfterDeductViewUnit,omitempty"`
	// The display value of the capacity after deduction.
	//
	// example:
	//
	// 23.896484
	CapacityAfterDeductViewValue *string `json:"CapacityAfterDeductViewValue,omitempty" xml:"CapacityAfterDeductViewValue,omitempty"`
	// The display unit of the capacity before deduction.
	//
	// example:
	//
	// GB
	CapacityBeforeDeductViewUnit *string `json:"CapacityBeforeDeductViewUnit,omitempty" xml:"CapacityBeforeDeductViewUnit,omitempty"`
	// The display value of the capacity before deduction.
	//
	// example:
	//
	// 40.000000
	CapacityBeforeDeductViewValue *string `json:"CapacityBeforeDeductViewValue,omitempty" xml:"CapacityBeforeDeductViewValue,omitempty"`
	// The display unit of the deducted capacity.
	//
	// example:
	//
	// GB
	CapacityDeductedViewUnit *string `json:"CapacityDeductedViewUnit,omitempty" xml:"CapacityDeductedViewUnit,omitempty"`
	// The display value of the deducted capacity.
	//
	// example:
	//
	// 16.103515
	CapacityDeductedViewValue *string `json:"CapacityDeductedViewValue,omitempty" xml:"CapacityDeductedViewValue,omitempty"`
	// The capacity type.
	CapacityType *DescribeDeductLogsResponseBodyDataCapacityType `json:"CapacityType,omitempty" xml:"CapacityType,omitempty" type:"Struct"`
	// The capacity type code.
	//
	// example:
	//
	// absolute
	CapacityTypeCode *string `json:"CapacityTypeCode,omitempty" xml:"CapacityTypeCode,omitempty"`
	// The capacity type name.
	//
	// example:
	//
	// 总量恒定型
	CapacityTypeName *string `json:"CapacityTypeName,omitempty" xml:"CapacityTypeName,omitempty"`
	// The commodity.
	Commodity *DescribeDeductLogsResponseBodyDataCommodity `json:"Commodity,omitempty" xml:"Commodity,omitempty" type:"Struct"`
	// The commodity code.
	//
	// example:
	//
	// ossbag
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// The commodity name.
	//
	// example:
	//
	// 对象存储OSS资源包(包月)
	CommodityName *string `json:"CommodityName,omitempty" xml:"CommodityName,omitempty"`
	// The commitment cycle.
	CycleType *DescribeDeductLogsResponseBodyDataCycleType `json:"CycleType,omitempty" xml:"CycleType,omitempty" type:"Struct"`
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
	// The deduction time.
	//
	// example:
	//
	// 1679039572000
	DeductTime *int64 `json:"DeductTime,omitempty" xml:"DeductTime,omitempty"`
	// The deduction factor.
	//
	// example:
	//
	// 1
	Factor *string `json:"Factor,omitempty" xml:"Factor,omitempty"`
	// The ID of the account to which the instance belongs.
	//
	// example:
	//
	// 1990699401005016
	InstanceBelongAccountId *int64 `json:"InstanceBelongAccountId,omitempty" xml:"InstanceBelongAccountId,omitempty"`
	// The name of the account to which the instance belongs.
	//
	// example:
	//
	// icloudtest
	InstanceBelongAccountName *string `json:"InstanceBelongAccountName,omitempty" xml:"InstanceBelongAccountName,omitempty"`
	// The instance name.
	//
	// example:
	//
	// OSSBAG-cn-****s
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The display unit of the metering amount after deduction.
	//
	// example:
	//
	// Byte
	MeasureAfterDeductViewUnit *string `json:"MeasureAfterDeductViewUnit,omitempty" xml:"MeasureAfterDeductViewUnit,omitempty"`
	// The display value of the metering amount after deduction.
	//
	// example:
	//
	// 0
	MeasureAfterDeductViewValue *string `json:"MeasureAfterDeductViewValue,omitempty" xml:"MeasureAfterDeductViewValue,omitempty"`
	// The display unit of the metering amount before deduction.
	//
	// example:
	//
	// GB
	MeasureBeforeDeductViewUnit *string `json:"MeasureBeforeDeductViewUnit,omitempty" xml:"MeasureBeforeDeductViewUnit,omitempty"`
	// The display value of the metering amount before deduction.
	//
	// example:
	//
	// 16.103515
	MeasureBeforeDeductViewValue *string `json:"MeasureBeforeDeductViewValue,omitempty" xml:"MeasureBeforeDeductViewValue,omitempty"`
	// The display unit of the deducted metering amount.
	//
	// example:
	//
	// GB
	MeasureDeductedViewUnit *string `json:"MeasureDeductedViewUnit,omitempty" xml:"MeasureDeductedViewUnit,omitempty"`
	// The display value of the deducted metering amount.
	//
	// example:
	//
	// 16.103515
	MeasureDeductedViewValue *string `json:"MeasureDeductedViewValue,omitempty" xml:"MeasureDeductedViewValue,omitempty"`
	// The product.
	Product *DescribeDeductLogsResponseBodyDataProduct `json:"Product,omitempty" xml:"Product,omitempty" type:"Struct"`
	// The product code.
	//
	// example:
	//
	// oss
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The product name.
	//
	// example:
	//
	// 对象存储
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	// The deduction account ID.
	//
	// example:
	//
	// 1990699401005016
	RelationAccountId *int64 `json:"RelationAccountId,omitempty" xml:"RelationAccountId,omitempty"`
	// The deduction account name.
	//
	// example:
	//
	// icloudtest
	RelationAccountName *string `json:"RelationAccountName,omitempty" xml:"RelationAccountName,omitempty"`
	// The template.
	Template *DescribeDeductLogsResponseBodyDataTemplate `json:"Template,omitempty" xml:"Template,omitempty" type:"Struct"`
	// The template code.
	//
	// example:
	//
	// FPT_ossbag********
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	// The template name.
	//
	// example:
	//
	// 标准存储包
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s DescribeDeductLogsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeductLogsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDeductLogsResponseBodyData) GetAccountId() *int64 {
	return s.AccountId
}

func (s *DescribeDeductLogsResponseBodyData) GetAccountName() *string {
	return s.AccountName
}

func (s *DescribeDeductLogsResponseBodyData) GetBillingCommodity() *DescribeDeductLogsResponseBodyDataBillingCommodity {
	return s.BillingCommodity
}

func (s *DescribeDeductLogsResponseBodyData) GetBillingCommodityCode() *string {
	return s.BillingCommodityCode
}

func (s *DescribeDeductLogsResponseBodyData) GetBillingCommodityName() *string {
	return s.BillingCommodityName
}

func (s *DescribeDeductLogsResponseBodyData) GetBillingEndTime() *int64 {
	return s.BillingEndTime
}

func (s *DescribeDeductLogsResponseBodyData) GetBillingInstanceId() *string {
	return s.BillingInstanceId
}

func (s *DescribeDeductLogsResponseBodyData) GetBillingPriceField() *DescribeDeductLogsResponseBodyDataBillingPriceField {
	return s.BillingPriceField
}

func (s *DescribeDeductLogsResponseBodyData) GetBillingPriceFieldCode() *string {
	return s.BillingPriceFieldCode
}

func (s *DescribeDeductLogsResponseBodyData) GetBillingPriceFieldName() *string {
	return s.BillingPriceFieldName
}

func (s *DescribeDeductLogsResponseBodyData) GetBillingStartTime() *int64 {
	return s.BillingStartTime
}

func (s *DescribeDeductLogsResponseBodyData) GetCapacityAfterDeductViewUnit() *string {
	return s.CapacityAfterDeductViewUnit
}

func (s *DescribeDeductLogsResponseBodyData) GetCapacityAfterDeductViewValue() *string {
	return s.CapacityAfterDeductViewValue
}

func (s *DescribeDeductLogsResponseBodyData) GetCapacityBeforeDeductViewUnit() *string {
	return s.CapacityBeforeDeductViewUnit
}

func (s *DescribeDeductLogsResponseBodyData) GetCapacityBeforeDeductViewValue() *string {
	return s.CapacityBeforeDeductViewValue
}

func (s *DescribeDeductLogsResponseBodyData) GetCapacityDeductedViewUnit() *string {
	return s.CapacityDeductedViewUnit
}

func (s *DescribeDeductLogsResponseBodyData) GetCapacityDeductedViewValue() *string {
	return s.CapacityDeductedViewValue
}

func (s *DescribeDeductLogsResponseBodyData) GetCapacityType() *DescribeDeductLogsResponseBodyDataCapacityType {
	return s.CapacityType
}

func (s *DescribeDeductLogsResponseBodyData) GetCapacityTypeCode() *string {
	return s.CapacityTypeCode
}

func (s *DescribeDeductLogsResponseBodyData) GetCapacityTypeName() *string {
	return s.CapacityTypeName
}

func (s *DescribeDeductLogsResponseBodyData) GetCommodity() *DescribeDeductLogsResponseBodyDataCommodity {
	return s.Commodity
}

func (s *DescribeDeductLogsResponseBodyData) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *DescribeDeductLogsResponseBodyData) GetCommodityName() *string {
	return s.CommodityName
}

func (s *DescribeDeductLogsResponseBodyData) GetCycleType() *DescribeDeductLogsResponseBodyDataCycleType {
	return s.CycleType
}

func (s *DescribeDeductLogsResponseBodyData) GetCycleTypeCode() *string {
	return s.CycleTypeCode
}

func (s *DescribeDeductLogsResponseBodyData) GetCycleTypeName() *string {
	return s.CycleTypeName
}

func (s *DescribeDeductLogsResponseBodyData) GetDeductTime() *int64 {
	return s.DeductTime
}

func (s *DescribeDeductLogsResponseBodyData) GetFactor() *string {
	return s.Factor
}

func (s *DescribeDeductLogsResponseBodyData) GetInstanceBelongAccountId() *int64 {
	return s.InstanceBelongAccountId
}

func (s *DescribeDeductLogsResponseBodyData) GetInstanceBelongAccountName() *string {
	return s.InstanceBelongAccountName
}

func (s *DescribeDeductLogsResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeDeductLogsResponseBodyData) GetMeasureAfterDeductViewUnit() *string {
	return s.MeasureAfterDeductViewUnit
}

func (s *DescribeDeductLogsResponseBodyData) GetMeasureAfterDeductViewValue() *string {
	return s.MeasureAfterDeductViewValue
}

func (s *DescribeDeductLogsResponseBodyData) GetMeasureBeforeDeductViewUnit() *string {
	return s.MeasureBeforeDeductViewUnit
}

func (s *DescribeDeductLogsResponseBodyData) GetMeasureBeforeDeductViewValue() *string {
	return s.MeasureBeforeDeductViewValue
}

func (s *DescribeDeductLogsResponseBodyData) GetMeasureDeductedViewUnit() *string {
	return s.MeasureDeductedViewUnit
}

func (s *DescribeDeductLogsResponseBodyData) GetMeasureDeductedViewValue() *string {
	return s.MeasureDeductedViewValue
}

func (s *DescribeDeductLogsResponseBodyData) GetProduct() *DescribeDeductLogsResponseBodyDataProduct {
	return s.Product
}

func (s *DescribeDeductLogsResponseBodyData) GetProductCode() *string {
	return s.ProductCode
}

func (s *DescribeDeductLogsResponseBodyData) GetProductName() *string {
	return s.ProductName
}

func (s *DescribeDeductLogsResponseBodyData) GetRelationAccountId() *int64 {
	return s.RelationAccountId
}

func (s *DescribeDeductLogsResponseBodyData) GetRelationAccountName() *string {
	return s.RelationAccountName
}

func (s *DescribeDeductLogsResponseBodyData) GetTemplate() *DescribeDeductLogsResponseBodyDataTemplate {
	return s.Template
}

func (s *DescribeDeductLogsResponseBodyData) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *DescribeDeductLogsResponseBodyData) GetTemplateName() *string {
	return s.TemplateName
}

func (s *DescribeDeductLogsResponseBodyData) SetAccountId(v int64) *DescribeDeductLogsResponseBodyData {
	s.AccountId = &v
	return s
}

func (s *DescribeDeductLogsResponseBodyData) SetAccountName(v string) *DescribeDeductLogsResponseBodyData {
	s.AccountName = &v
	return s
}

func (s *DescribeDeductLogsResponseBodyData) SetBillingCommodity(v *DescribeDeductLogsResponseBodyDataBillingCommodity) *DescribeDeductLogsResponseBodyData {
	s.BillingCommodity = v
	return s
}

func (s *DescribeDeductLogsResponseBodyData) SetBillingCommodityCode(v string) *DescribeDeductLogsResponseBodyData {
	s.BillingCommodityCode = &v
	return s
}

func (s *DescribeDeductLogsResponseBodyData) SetBillingCommodityName(v string) *DescribeDeductLogsResponseBodyData {
	s.BillingCommodityName = &v
	return s
}

func (s *DescribeDeductLogsResponseBodyData) SetBillingEndTime(v int64) *DescribeDeductLogsResponseBodyData {
	s.BillingEndTime = &v
	return s
}

func (s *DescribeDeductLogsResponseBodyData) SetBillingInstanceId(v string) *DescribeDeductLogsResponseBodyData {
	s.BillingInstanceId = &v
	return s
}

func (s *DescribeDeductLogsResponseBodyData) SetBillingPriceField(v *DescribeDeductLogsResponseBodyDataBillingPriceField) *DescribeDeductLogsResponseBodyData {
	s.BillingPriceField = v
	return s
}

func (s *DescribeDeductLogsResponseBodyData) SetBillingPriceFieldCode(v string) *DescribeDeductLogsResponseBodyData {
	s.BillingPriceFieldCode = &v
	return s
}

func (s *DescribeDeductLogsResponseBodyData) SetBillingPriceFieldName(v string) *DescribeDeductLogsResponseBodyData {
	s.BillingPriceFieldName = &v
	return s
}

func (s *DescribeDeductLogsResponseBodyData) SetBillingStartTime(v int64) *DescribeDeductLogsResponseBodyData {
	s.BillingStartTime = &v
	return s
}

func (s *DescribeDeductLogsResponseBodyData) SetCapacityAfterDeductViewUnit(v string) *DescribeDeductLogsResponseBodyData {
	s.CapacityAfterDeductViewUnit = &v
	return s
}

func (s *DescribeDeductLogsResponseBodyData) SetCapacityAfterDeductViewValue(v string) *DescribeDeductLogsResponseBodyData {
	s.CapacityAfterDeductViewValue = &v
	return s
}

func (s *DescribeDeductLogsResponseBodyData) SetCapacityBeforeDeductViewUnit(v string) *DescribeDeductLogsResponseBodyData {
	s.CapacityBeforeDeductViewUnit = &v
	return s
}

func (s *DescribeDeductLogsResponseBodyData) SetCapacityBeforeDeductViewValue(v string) *DescribeDeductLogsResponseBodyData {
	s.CapacityBeforeDeductViewValue = &v
	return s
}

func (s *DescribeDeductLogsResponseBodyData) SetCapacityDeductedViewUnit(v string) *DescribeDeductLogsResponseBodyData {
	s.CapacityDeductedViewUnit = &v
	return s
}

func (s *DescribeDeductLogsResponseBodyData) SetCapacityDeductedViewValue(v string) *DescribeDeductLogsResponseBodyData {
	s.CapacityDeductedViewValue = &v
	return s
}

func (s *DescribeDeductLogsResponseBodyData) SetCapacityType(v *DescribeDeductLogsResponseBodyDataCapacityType) *DescribeDeductLogsResponseBodyData {
	s.CapacityType = v
	return s
}

func (s *DescribeDeductLogsResponseBodyData) SetCapacityTypeCode(v string) *DescribeDeductLogsResponseBodyData {
	s.CapacityTypeCode = &v
	return s
}

func (s *DescribeDeductLogsResponseBodyData) SetCapacityTypeName(v string) *DescribeDeductLogsResponseBodyData {
	s.CapacityTypeName = &v
	return s
}

func (s *DescribeDeductLogsResponseBodyData) SetCommodity(v *DescribeDeductLogsResponseBodyDataCommodity) *DescribeDeductLogsResponseBodyData {
	s.Commodity = v
	return s
}

func (s *DescribeDeductLogsResponseBodyData) SetCommodityCode(v string) *DescribeDeductLogsResponseBodyData {
	s.CommodityCode = &v
	return s
}

func (s *DescribeDeductLogsResponseBodyData) SetCommodityName(v string) *DescribeDeductLogsResponseBodyData {
	s.CommodityName = &v
	return s
}

func (s *DescribeDeductLogsResponseBodyData) SetCycleType(v *DescribeDeductLogsResponseBodyDataCycleType) *DescribeDeductLogsResponseBodyData {
	s.CycleType = v
	return s
}

func (s *DescribeDeductLogsResponseBodyData) SetCycleTypeCode(v string) *DescribeDeductLogsResponseBodyData {
	s.CycleTypeCode = &v
	return s
}

func (s *DescribeDeductLogsResponseBodyData) SetCycleTypeName(v string) *DescribeDeductLogsResponseBodyData {
	s.CycleTypeName = &v
	return s
}

func (s *DescribeDeductLogsResponseBodyData) SetDeductTime(v int64) *DescribeDeductLogsResponseBodyData {
	s.DeductTime = &v
	return s
}

func (s *DescribeDeductLogsResponseBodyData) SetFactor(v string) *DescribeDeductLogsResponseBodyData {
	s.Factor = &v
	return s
}

func (s *DescribeDeductLogsResponseBodyData) SetInstanceBelongAccountId(v int64) *DescribeDeductLogsResponseBodyData {
	s.InstanceBelongAccountId = &v
	return s
}

func (s *DescribeDeductLogsResponseBodyData) SetInstanceBelongAccountName(v string) *DescribeDeductLogsResponseBodyData {
	s.InstanceBelongAccountName = &v
	return s
}

func (s *DescribeDeductLogsResponseBodyData) SetInstanceId(v string) *DescribeDeductLogsResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *DescribeDeductLogsResponseBodyData) SetMeasureAfterDeductViewUnit(v string) *DescribeDeductLogsResponseBodyData {
	s.MeasureAfterDeductViewUnit = &v
	return s
}

func (s *DescribeDeductLogsResponseBodyData) SetMeasureAfterDeductViewValue(v string) *DescribeDeductLogsResponseBodyData {
	s.MeasureAfterDeductViewValue = &v
	return s
}

func (s *DescribeDeductLogsResponseBodyData) SetMeasureBeforeDeductViewUnit(v string) *DescribeDeductLogsResponseBodyData {
	s.MeasureBeforeDeductViewUnit = &v
	return s
}

func (s *DescribeDeductLogsResponseBodyData) SetMeasureBeforeDeductViewValue(v string) *DescribeDeductLogsResponseBodyData {
	s.MeasureBeforeDeductViewValue = &v
	return s
}

func (s *DescribeDeductLogsResponseBodyData) SetMeasureDeductedViewUnit(v string) *DescribeDeductLogsResponseBodyData {
	s.MeasureDeductedViewUnit = &v
	return s
}

func (s *DescribeDeductLogsResponseBodyData) SetMeasureDeductedViewValue(v string) *DescribeDeductLogsResponseBodyData {
	s.MeasureDeductedViewValue = &v
	return s
}

func (s *DescribeDeductLogsResponseBodyData) SetProduct(v *DescribeDeductLogsResponseBodyDataProduct) *DescribeDeductLogsResponseBodyData {
	s.Product = v
	return s
}

func (s *DescribeDeductLogsResponseBodyData) SetProductCode(v string) *DescribeDeductLogsResponseBodyData {
	s.ProductCode = &v
	return s
}

func (s *DescribeDeductLogsResponseBodyData) SetProductName(v string) *DescribeDeductLogsResponseBodyData {
	s.ProductName = &v
	return s
}

func (s *DescribeDeductLogsResponseBodyData) SetRelationAccountId(v int64) *DescribeDeductLogsResponseBodyData {
	s.RelationAccountId = &v
	return s
}

func (s *DescribeDeductLogsResponseBodyData) SetRelationAccountName(v string) *DescribeDeductLogsResponseBodyData {
	s.RelationAccountName = &v
	return s
}

func (s *DescribeDeductLogsResponseBodyData) SetTemplate(v *DescribeDeductLogsResponseBodyDataTemplate) *DescribeDeductLogsResponseBodyData {
	s.Template = v
	return s
}

func (s *DescribeDeductLogsResponseBodyData) SetTemplateCode(v string) *DescribeDeductLogsResponseBodyData {
	s.TemplateCode = &v
	return s
}

func (s *DescribeDeductLogsResponseBodyData) SetTemplateName(v string) *DescribeDeductLogsResponseBodyData {
	s.TemplateName = &v
	return s
}

func (s *DescribeDeductLogsResponseBodyData) Validate() error {
	if s.BillingCommodity != nil {
		if err := s.BillingCommodity.Validate(); err != nil {
			return err
		}
	}
	if s.BillingPriceField != nil {
		if err := s.BillingPriceField.Validate(); err != nil {
			return err
		}
	}
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
	if s.Product != nil {
		if err := s.Product.Validate(); err != nil {
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

type DescribeDeductLogsResponseBodyDataBillingCommodity struct {
	// The property code.
	//
	// example:
	//
	// snapshot
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The property name.
	//
	// example:
	//
	// 云服务器ECS-快照
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeDeductLogsResponseBodyDataBillingCommodity) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeductLogsResponseBodyDataBillingCommodity) GoString() string {
	return s.String()
}

func (s *DescribeDeductLogsResponseBodyDataBillingCommodity) GetCode() *string {
	return s.Code
}

func (s *DescribeDeductLogsResponseBodyDataBillingCommodity) GetName() *string {
	return s.Name
}

func (s *DescribeDeductLogsResponseBodyDataBillingCommodity) SetCode(v string) *DescribeDeductLogsResponseBodyDataBillingCommodity {
	s.Code = &v
	return s
}

func (s *DescribeDeductLogsResponseBodyDataBillingCommodity) SetName(v string) *DescribeDeductLogsResponseBodyDataBillingCommodity {
	s.Name = &v
	return s
}

func (s *DescribeDeductLogsResponseBodyDataBillingCommodity) Validate() error {
	return dara.Validate(s)
}

type DescribeDeductLogsResponseBodyDataBillingPriceField struct {
	// The property code.
	//
	// example:
	//
	// Storage
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The property name.
	//
	// example:
	//
	// 零折使用容量
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeDeductLogsResponseBodyDataBillingPriceField) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeductLogsResponseBodyDataBillingPriceField) GoString() string {
	return s.String()
}

func (s *DescribeDeductLogsResponseBodyDataBillingPriceField) GetCode() *string {
	return s.Code
}

func (s *DescribeDeductLogsResponseBodyDataBillingPriceField) GetName() *string {
	return s.Name
}

func (s *DescribeDeductLogsResponseBodyDataBillingPriceField) SetCode(v string) *DescribeDeductLogsResponseBodyDataBillingPriceField {
	s.Code = &v
	return s
}

func (s *DescribeDeductLogsResponseBodyDataBillingPriceField) SetName(v string) *DescribeDeductLogsResponseBodyDataBillingPriceField {
	s.Name = &v
	return s
}

func (s *DescribeDeductLogsResponseBodyDataBillingPriceField) Validate() error {
	return dara.Validate(s)
}

type DescribeDeductLogsResponseBodyDataCapacityType struct {
	// The property code.
	//
	// example:
	//
	// absolute
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The property name.
	//
	// example:
	//
	// 总量恒定型
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeDeductLogsResponseBodyDataCapacityType) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeductLogsResponseBodyDataCapacityType) GoString() string {
	return s.String()
}

func (s *DescribeDeductLogsResponseBodyDataCapacityType) GetCode() *string {
	return s.Code
}

func (s *DescribeDeductLogsResponseBodyDataCapacityType) GetName() *string {
	return s.Name
}

func (s *DescribeDeductLogsResponseBodyDataCapacityType) SetCode(v string) *DescribeDeductLogsResponseBodyDataCapacityType {
	s.Code = &v
	return s
}

func (s *DescribeDeductLogsResponseBodyDataCapacityType) SetName(v string) *DescribeDeductLogsResponseBodyDataCapacityType {
	s.Name = &v
	return s
}

func (s *DescribeDeductLogsResponseBodyDataCapacityType) Validate() error {
	return dara.Validate(s)
}

type DescribeDeductLogsResponseBodyDataCommodity struct {
	// The property code.
	//
	// example:
	//
	// ossbag
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The property name.
	//
	// example:
	//
	// 对象存储OSS资源包(包月)
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeDeductLogsResponseBodyDataCommodity) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeductLogsResponseBodyDataCommodity) GoString() string {
	return s.String()
}

func (s *DescribeDeductLogsResponseBodyDataCommodity) GetCode() *string {
	return s.Code
}

func (s *DescribeDeductLogsResponseBodyDataCommodity) GetName() *string {
	return s.Name
}

func (s *DescribeDeductLogsResponseBodyDataCommodity) SetCode(v string) *DescribeDeductLogsResponseBodyDataCommodity {
	s.Code = &v
	return s
}

func (s *DescribeDeductLogsResponseBodyDataCommodity) SetName(v string) *DescribeDeductLogsResponseBodyDataCommodity {
	s.Name = &v
	return s
}

func (s *DescribeDeductLogsResponseBodyDataCommodity) Validate() error {
	return dara.Validate(s)
}

type DescribeDeductLogsResponseBodyDataCycleType struct {
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

func (s DescribeDeductLogsResponseBodyDataCycleType) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeductLogsResponseBodyDataCycleType) GoString() string {
	return s.String()
}

func (s *DescribeDeductLogsResponseBodyDataCycleType) GetCode() *string {
	return s.Code
}

func (s *DescribeDeductLogsResponseBodyDataCycleType) GetName() *string {
	return s.Name
}

func (s *DescribeDeductLogsResponseBodyDataCycleType) SetCode(v string) *DescribeDeductLogsResponseBodyDataCycleType {
	s.Code = &v
	return s
}

func (s *DescribeDeductLogsResponseBodyDataCycleType) SetName(v string) *DescribeDeductLogsResponseBodyDataCycleType {
	s.Name = &v
	return s
}

func (s *DescribeDeductLogsResponseBodyDataCycleType) Validate() error {
	return dara.Validate(s)
}

type DescribeDeductLogsResponseBodyDataProduct struct {
	// The property code.
	//
	// example:
	//
	// oss
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The property name.
	//
	// example:
	//
	// 对象存储
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeDeductLogsResponseBodyDataProduct) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeductLogsResponseBodyDataProduct) GoString() string {
	return s.String()
}

func (s *DescribeDeductLogsResponseBodyDataProduct) GetCode() *string {
	return s.Code
}

func (s *DescribeDeductLogsResponseBodyDataProduct) GetName() *string {
	return s.Name
}

func (s *DescribeDeductLogsResponseBodyDataProduct) SetCode(v string) *DescribeDeductLogsResponseBodyDataProduct {
	s.Code = &v
	return s
}

func (s *DescribeDeductLogsResponseBodyDataProduct) SetName(v string) *DescribeDeductLogsResponseBodyDataProduct {
	s.Name = &v
	return s
}

func (s *DescribeDeductLogsResponseBodyDataProduct) Validate() error {
	return dara.Validate(s)
}

type DescribeDeductLogsResponseBodyDataTemplate struct {
	// The property code.
	//
	// example:
	//
	// FPT_ossbag********
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The property name.
	//
	// example:
	//
	// 标准存储包
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeDeductLogsResponseBodyDataTemplate) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeductLogsResponseBodyDataTemplate) GoString() string {
	return s.String()
}

func (s *DescribeDeductLogsResponseBodyDataTemplate) GetCode() *string {
	return s.Code
}

func (s *DescribeDeductLogsResponseBodyDataTemplate) GetName() *string {
	return s.Name
}

func (s *DescribeDeductLogsResponseBodyDataTemplate) SetCode(v string) *DescribeDeductLogsResponseBodyDataTemplate {
	s.Code = &v
	return s
}

func (s *DescribeDeductLogsResponseBodyDataTemplate) SetName(v string) *DescribeDeductLogsResponseBodyDataTemplate {
	s.Name = &v
	return s
}

func (s *DescribeDeductLogsResponseBodyDataTemplate) Validate() error {
	return dara.Validate(s)
}
