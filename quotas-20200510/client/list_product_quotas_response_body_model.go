// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProductQuotasResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListProductQuotasResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListProductQuotasResponseBody
	GetNextToken() *string
	SetQuotas(v []*ListProductQuotasResponseBodyQuotas) *ListProductQuotasResponseBody
	GetQuotas() []*ListProductQuotasResponseBodyQuotas
	SetRequestId(v string) *ListProductQuotasResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListProductQuotasResponseBody
	GetTotalCount() *int32
}

type ListProductQuotasResponseBody struct {
	// The number of records that are returned for the query.
	//
	// example:
	//
	// 1
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The returned value of NextToken is a pagination token, which can be used in the next request to retrieve a new page of results.
	//
	// 	- If NextToken is empty, no next page exists.
	//
	// 	- If a value of NextToken is returned, the value indicates the token that is used for the next query.
	//
	// example:
	//
	// AAAAAd98/tlL5GF2aM7UMKQGM8LZesIPr0CbfxASQvHV/pwcmVKNfdBbW8OPld3NvG9Cy8+dNcyFzyUttQA3IONfBhRGpXFyiVoTgK+dupBsP2mX
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The details of the quotas.
	Quotas []*ListProductQuotasResponseBodyQuotas `json:"Quotas,omitempty" xml:"Quotas,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// D0131FD5-5397-44FE-BF5A-4B7165B813CC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 4
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListProductQuotasResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListProductQuotasResponseBody) GoString() string {
	return s.String()
}

func (s *ListProductQuotasResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListProductQuotasResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListProductQuotasResponseBody) GetQuotas() []*ListProductQuotasResponseBodyQuotas {
	return s.Quotas
}

func (s *ListProductQuotasResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListProductQuotasResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListProductQuotasResponseBody) SetMaxResults(v int32) *ListProductQuotasResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListProductQuotasResponseBody) SetNextToken(v string) *ListProductQuotasResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListProductQuotasResponseBody) SetQuotas(v []*ListProductQuotasResponseBodyQuotas) *ListProductQuotasResponseBody {
	s.Quotas = v
	return s
}

func (s *ListProductQuotasResponseBody) SetRequestId(v string) *ListProductQuotasResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProductQuotasResponseBody) SetTotalCount(v int32) *ListProductQuotasResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListProductQuotasResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListProductQuotasResponseBodyQuotas struct {
	// Indicates whether the quota is adjustable.
	//
	// Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Adjustable *bool `json:"Adjustable,omitempty" xml:"Adjustable,omitempty"`
	// The range of the quota value that can be requested for the quota item.
	ApplicableRange []*float32 `json:"ApplicableRange,omitempty" xml:"ApplicableRange,omitempty" type:"Repeated"`
	// The type of the adjustable value.
	//
	// Valid values:
	//
	// 	- continuous
	//
	// 	- discontinuous
	//
	// example:
	//
	// discontinuous
	ApplicableType *string `json:"ApplicableType,omitempty" xml:"ApplicableType,omitempty"`
	// The reason for submitting a quota increase request.
	//
	// example:
	//
	// The business xxx is expected to grow by 50%.
	ApplyReasonTips *string `json:"ApplyReasonTips,omitempty" xml:"ApplyReasonTips,omitempty"`
	// Indicates whether the system shows the used value of the quota.
	//
	// Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Consumable *bool `json:"Consumable,omitempty" xml:"Consumable,omitempty"`
	// The quota dimension. Format: `{"regionId":"Region"}`.
	//
	// example:
	//
	// {"regionId":"cn-hangzhou"}
	Dimensions map[string]interface{} `json:"Dimensions,omitempty" xml:"Dimensions,omitempty"`
	// The start time of the validity period of the quota. The value is displayed in UTC.
	//
	// example:
	//
	// 2022-09-28T06:07:00Z
	EffectiveTime *string `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	// The end time of the validity period of the quota. The value is displayed in UTC.
	//
	// example:
	//
	// 2022-09-29T06:07:00Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// Indicates whether the quota is a global quota.
	//
	// Valid values:
	//
	// 	- true: The quota is shared in all regions.
	//
	// 	- false: The quota is independently used in a region.
	//
	// example:
	//
	// true
	GlobalQuota *bool `json:"GlobalQuota,omitempty" xml:"GlobalQuota,omitempty"`
	// The calculation cycle of the quota.
	Period *ListProductQuotasResponseBodyQuotasPeriod `json:"Period,omitempty" xml:"Period,omitempty" type:"Struct"`
	// The abbreviation of the Alibaba Cloud service name.
	//
	// example:
	//
	// ecs-spec
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The quota ID.
	//
	// example:
	//
	// ecs.g5.2xlarge
	QuotaActionCode *string `json:"QuotaActionCode,omitempty" xml:"QuotaActionCode,omitempty"`
	// The Alibaba Cloud Resource Name (ARN) of the quota.
	//
	// example:
	//
	// acs:quotas:cn-hangzhou:107992689699****:quota/ecs/ecs.g5.2xlarge/postpaid/vpc/cn-hangzhou/instancetype/cn-hangzhou-i
	QuotaArn *string `json:"QuotaArn,omitempty" xml:"QuotaArn,omitempty"`
	// The type of the quota.
	//
	// Valid values:
	//
	// 	- FlowControl: API rate limit
	//
	// 	- WhiteListLabel: whitelist quota
	//
	// 	- CommonQuota: general quota
	//
	// example:
	//
	// CommonQuota
	QuotaCategory *string `json:"QuotaCategory,omitempty" xml:"QuotaCategory,omitempty"`
	// The description of the quota.
	//
	// example:
	//
	// ecs.g5.2xlarge
	QuotaDescription *string `json:"QuotaDescription,omitempty" xml:"QuotaDescription,omitempty"`
	// The details of the quotas.
	QuotaItems []*ListProductQuotasResponseBodyQuotasQuotaItems `json:"QuotaItems,omitempty" xml:"QuotaItems,omitempty" type:"Repeated"`
	// The name of the quota.
	//
	// example:
	//
	// ecs.g5.2xlarge
	QuotaName *string `json:"QuotaName,omitempty" xml:"QuotaName,omitempty"`
	// The type of the quota.
	//
	// Valid values:
	//
	// 	- normal
	//
	// 	- privilege
	//
	// example:
	//
	// privilege
	QuotaType *string `json:"QuotaType,omitempty" xml:"QuotaType,omitempty"`
	// The unit of the quota.
	//
	// example:
	//
	// AMOUNT
	QuotaUnit *string `json:"QuotaUnit,omitempty" xml:"QuotaUnit,omitempty"`
	// The range of the quota value that can be requested for the current quota item. When you configure a quota template, you can use the range as a reference.
	SupportedRange []*float32 `json:"SupportedRange,omitempty" xml:"SupportedRange,omitempty" type:"Repeated"`
	// The value of the quota.
	//
	// example:
	//
	// 200
	TotalQuota *float32 `json:"TotalQuota,omitempty" xml:"TotalQuota,omitempty"`
	// The quota usage.
	//
	// example:
	//
	// 1
	TotalUsage *float32 `json:"TotalUsage,omitempty" xml:"TotalUsage,omitempty"`
	// The reason why the quota is not adjustable.
	//
	// Valid values:
	//
	// 	- limitReached: The quota limit is reached.
	//
	// 	- nonactivated: The service is not activated.
	//
	// 	- applicationProcess: The application is being processed.
	//
	// example:
	//
	// applicationProcess
	UnadjustableDetail *string `json:"UnadjustableDetail,omitempty" xml:"UnadjustableDetail,omitempty"`
	// The monitoring information of the quota in CloudMonitor.
	//
	// >  If this parameter is empty, no monitoring data of the quota exists in CloudMonitor.
	UsageMetric *ListProductQuotasResponseBodyQuotasUsageMetric `json:"UsageMetric,omitempty" xml:"UsageMetric,omitempty" type:"Struct"`
}

func (s ListProductQuotasResponseBodyQuotas) String() string {
	return dara.Prettify(s)
}

func (s ListProductQuotasResponseBodyQuotas) GoString() string {
	return s.String()
}

func (s *ListProductQuotasResponseBodyQuotas) GetAdjustable() *bool {
	return s.Adjustable
}

func (s *ListProductQuotasResponseBodyQuotas) GetApplicableRange() []*float32 {
	return s.ApplicableRange
}

func (s *ListProductQuotasResponseBodyQuotas) GetApplicableType() *string {
	return s.ApplicableType
}

func (s *ListProductQuotasResponseBodyQuotas) GetApplyReasonTips() *string {
	return s.ApplyReasonTips
}

func (s *ListProductQuotasResponseBodyQuotas) GetConsumable() *bool {
	return s.Consumable
}

func (s *ListProductQuotasResponseBodyQuotas) GetDimensions() map[string]interface{} {
	return s.Dimensions
}

func (s *ListProductQuotasResponseBodyQuotas) GetEffectiveTime() *string {
	return s.EffectiveTime
}

func (s *ListProductQuotasResponseBodyQuotas) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *ListProductQuotasResponseBodyQuotas) GetGlobalQuota() *bool {
	return s.GlobalQuota
}

func (s *ListProductQuotasResponseBodyQuotas) GetPeriod() *ListProductQuotasResponseBodyQuotasPeriod {
	return s.Period
}

func (s *ListProductQuotasResponseBodyQuotas) GetProductCode() *string {
	return s.ProductCode
}

func (s *ListProductQuotasResponseBodyQuotas) GetQuotaActionCode() *string {
	return s.QuotaActionCode
}

func (s *ListProductQuotasResponseBodyQuotas) GetQuotaArn() *string {
	return s.QuotaArn
}

func (s *ListProductQuotasResponseBodyQuotas) GetQuotaCategory() *string {
	return s.QuotaCategory
}

func (s *ListProductQuotasResponseBodyQuotas) GetQuotaDescription() *string {
	return s.QuotaDescription
}

func (s *ListProductQuotasResponseBodyQuotas) GetQuotaItems() []*ListProductQuotasResponseBodyQuotasQuotaItems {
	return s.QuotaItems
}

func (s *ListProductQuotasResponseBodyQuotas) GetQuotaName() *string {
	return s.QuotaName
}

func (s *ListProductQuotasResponseBodyQuotas) GetQuotaType() *string {
	return s.QuotaType
}

func (s *ListProductQuotasResponseBodyQuotas) GetQuotaUnit() *string {
	return s.QuotaUnit
}

func (s *ListProductQuotasResponseBodyQuotas) GetSupportedRange() []*float32 {
	return s.SupportedRange
}

func (s *ListProductQuotasResponseBodyQuotas) GetTotalQuota() *float32 {
	return s.TotalQuota
}

func (s *ListProductQuotasResponseBodyQuotas) GetTotalUsage() *float32 {
	return s.TotalUsage
}

func (s *ListProductQuotasResponseBodyQuotas) GetUnadjustableDetail() *string {
	return s.UnadjustableDetail
}

func (s *ListProductQuotasResponseBodyQuotas) GetUsageMetric() *ListProductQuotasResponseBodyQuotasUsageMetric {
	return s.UsageMetric
}

func (s *ListProductQuotasResponseBodyQuotas) SetAdjustable(v bool) *ListProductQuotasResponseBodyQuotas {
	s.Adjustable = &v
	return s
}

func (s *ListProductQuotasResponseBodyQuotas) SetApplicableRange(v []*float32) *ListProductQuotasResponseBodyQuotas {
	s.ApplicableRange = v
	return s
}

func (s *ListProductQuotasResponseBodyQuotas) SetApplicableType(v string) *ListProductQuotasResponseBodyQuotas {
	s.ApplicableType = &v
	return s
}

func (s *ListProductQuotasResponseBodyQuotas) SetApplyReasonTips(v string) *ListProductQuotasResponseBodyQuotas {
	s.ApplyReasonTips = &v
	return s
}

func (s *ListProductQuotasResponseBodyQuotas) SetConsumable(v bool) *ListProductQuotasResponseBodyQuotas {
	s.Consumable = &v
	return s
}

func (s *ListProductQuotasResponseBodyQuotas) SetDimensions(v map[string]interface{}) *ListProductQuotasResponseBodyQuotas {
	s.Dimensions = v
	return s
}

func (s *ListProductQuotasResponseBodyQuotas) SetEffectiveTime(v string) *ListProductQuotasResponseBodyQuotas {
	s.EffectiveTime = &v
	return s
}

func (s *ListProductQuotasResponseBodyQuotas) SetExpireTime(v string) *ListProductQuotasResponseBodyQuotas {
	s.ExpireTime = &v
	return s
}

func (s *ListProductQuotasResponseBodyQuotas) SetGlobalQuota(v bool) *ListProductQuotasResponseBodyQuotas {
	s.GlobalQuota = &v
	return s
}

func (s *ListProductQuotasResponseBodyQuotas) SetPeriod(v *ListProductQuotasResponseBodyQuotasPeriod) *ListProductQuotasResponseBodyQuotas {
	s.Period = v
	return s
}

func (s *ListProductQuotasResponseBodyQuotas) SetProductCode(v string) *ListProductQuotasResponseBodyQuotas {
	s.ProductCode = &v
	return s
}

func (s *ListProductQuotasResponseBodyQuotas) SetQuotaActionCode(v string) *ListProductQuotasResponseBodyQuotas {
	s.QuotaActionCode = &v
	return s
}

func (s *ListProductQuotasResponseBodyQuotas) SetQuotaArn(v string) *ListProductQuotasResponseBodyQuotas {
	s.QuotaArn = &v
	return s
}

func (s *ListProductQuotasResponseBodyQuotas) SetQuotaCategory(v string) *ListProductQuotasResponseBodyQuotas {
	s.QuotaCategory = &v
	return s
}

func (s *ListProductQuotasResponseBodyQuotas) SetQuotaDescription(v string) *ListProductQuotasResponseBodyQuotas {
	s.QuotaDescription = &v
	return s
}

func (s *ListProductQuotasResponseBodyQuotas) SetQuotaItems(v []*ListProductQuotasResponseBodyQuotasQuotaItems) *ListProductQuotasResponseBodyQuotas {
	s.QuotaItems = v
	return s
}

func (s *ListProductQuotasResponseBodyQuotas) SetQuotaName(v string) *ListProductQuotasResponseBodyQuotas {
	s.QuotaName = &v
	return s
}

func (s *ListProductQuotasResponseBodyQuotas) SetQuotaType(v string) *ListProductQuotasResponseBodyQuotas {
	s.QuotaType = &v
	return s
}

func (s *ListProductQuotasResponseBodyQuotas) SetQuotaUnit(v string) *ListProductQuotasResponseBodyQuotas {
	s.QuotaUnit = &v
	return s
}

func (s *ListProductQuotasResponseBodyQuotas) SetSupportedRange(v []*float32) *ListProductQuotasResponseBodyQuotas {
	s.SupportedRange = v
	return s
}

func (s *ListProductQuotasResponseBodyQuotas) SetTotalQuota(v float32) *ListProductQuotasResponseBodyQuotas {
	s.TotalQuota = &v
	return s
}

func (s *ListProductQuotasResponseBodyQuotas) SetTotalUsage(v float32) *ListProductQuotasResponseBodyQuotas {
	s.TotalUsage = &v
	return s
}

func (s *ListProductQuotasResponseBodyQuotas) SetUnadjustableDetail(v string) *ListProductQuotasResponseBodyQuotas {
	s.UnadjustableDetail = &v
	return s
}

func (s *ListProductQuotasResponseBodyQuotas) SetUsageMetric(v *ListProductQuotasResponseBodyQuotasUsageMetric) *ListProductQuotasResponseBodyQuotas {
	s.UsageMetric = v
	return s
}

func (s *ListProductQuotasResponseBodyQuotas) Validate() error {
	return dara.Validate(s)
}

type ListProductQuotasResponseBodyQuotasPeriod struct {
	// The unit of the calculation cycle.
	//
	// Valid values:
	//
	// 	- week
	//
	// 	- hour
	//
	// 	- day
	//
	// 	- second
	//
	// 	- minute
	//
	// example:
	//
	// day
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The value of the calculation cycle.
	//
	// example:
	//
	// 1
	PeriodValue *int32 `json:"PeriodValue,omitempty" xml:"PeriodValue,omitempty"`
}

func (s ListProductQuotasResponseBodyQuotasPeriod) String() string {
	return dara.Prettify(s)
}

func (s ListProductQuotasResponseBodyQuotasPeriod) GoString() string {
	return s.String()
}

func (s *ListProductQuotasResponseBodyQuotasPeriod) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *ListProductQuotasResponseBodyQuotasPeriod) GetPeriodValue() *int32 {
	return s.PeriodValue
}

func (s *ListProductQuotasResponseBodyQuotasPeriod) SetPeriodUnit(v string) *ListProductQuotasResponseBodyQuotasPeriod {
	s.PeriodUnit = &v
	return s
}

func (s *ListProductQuotasResponseBodyQuotasPeriod) SetPeriodValue(v int32) *ListProductQuotasResponseBodyQuotasPeriod {
	s.PeriodValue = &v
	return s
}

func (s *ListProductQuotasResponseBodyQuotasPeriod) Validate() error {
	return dara.Validate(s)
}

type ListProductQuotasResponseBodyQuotasQuotaItems struct {
	// The value of the quota.
	//
	// example:
	//
	// 10
	Quota *string `json:"Quota,omitempty" xml:"Quota,omitempty"`
	// The unit of the quota.
	//
	// example:
	//
	// AMOUNT
	QuotaUnit *string `json:"QuotaUnit,omitempty" xml:"QuotaUnit,omitempty"`
	// The category of the quota.
	//
	// Valid values:
	//
	// 	- BaseQuota: base quota
	//
	// 	- ReservedQuota: reserved quota
	//
	// example:
	//
	// BaseQuota
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The quota usage.
	//
	// example:
	//
	// 1
	Usage *string `json:"Usage,omitempty" xml:"Usage,omitempty"`
}

func (s ListProductQuotasResponseBodyQuotasQuotaItems) String() string {
	return dara.Prettify(s)
}

func (s ListProductQuotasResponseBodyQuotasQuotaItems) GoString() string {
	return s.String()
}

func (s *ListProductQuotasResponseBodyQuotasQuotaItems) GetQuota() *string {
	return s.Quota
}

func (s *ListProductQuotasResponseBodyQuotasQuotaItems) GetQuotaUnit() *string {
	return s.QuotaUnit
}

func (s *ListProductQuotasResponseBodyQuotasQuotaItems) GetType() *string {
	return s.Type
}

func (s *ListProductQuotasResponseBodyQuotasQuotaItems) GetUsage() *string {
	return s.Usage
}

func (s *ListProductQuotasResponseBodyQuotasQuotaItems) SetQuota(v string) *ListProductQuotasResponseBodyQuotasQuotaItems {
	s.Quota = &v
	return s
}

func (s *ListProductQuotasResponseBodyQuotasQuotaItems) SetQuotaUnit(v string) *ListProductQuotasResponseBodyQuotasQuotaItems {
	s.QuotaUnit = &v
	return s
}

func (s *ListProductQuotasResponseBodyQuotasQuotaItems) SetType(v string) *ListProductQuotasResponseBodyQuotasQuotaItems {
	s.Type = &v
	return s
}

func (s *ListProductQuotasResponseBodyQuotasQuotaItems) SetUsage(v string) *ListProductQuotasResponseBodyQuotasQuotaItems {
	s.Usage = &v
	return s
}

func (s *ListProductQuotasResponseBodyQuotasQuotaItems) Validate() error {
	return dara.Validate(s)
}

type ListProductQuotasResponseBodyQuotasUsageMetric struct {
	// The monitoring dimensions.
	//
	// The value is a collection of `key:value` pairs. Example: `{"productCode":"***","metricKey":"***","regionId":"***","label":"***"}`.
	MetricDimensions map[string]*string `json:"MetricDimensions,omitempty" xml:"MetricDimensions,omitempty"`
	// The monitoring metric.
	//
	// example:
	//
	// Usage
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// The monitoring namespace.
	//
	// example:
	//
	// acs_quotas_flowcontrol
	MetricNamespace *string `json:"MetricNamespace,omitempty" xml:"MetricNamespace,omitempty"`
}

func (s ListProductQuotasResponseBodyQuotasUsageMetric) String() string {
	return dara.Prettify(s)
}

func (s ListProductQuotasResponseBodyQuotasUsageMetric) GoString() string {
	return s.String()
}

func (s *ListProductQuotasResponseBodyQuotasUsageMetric) GetMetricDimensions() map[string]*string {
	return s.MetricDimensions
}

func (s *ListProductQuotasResponseBodyQuotasUsageMetric) GetMetricName() *string {
	return s.MetricName
}

func (s *ListProductQuotasResponseBodyQuotasUsageMetric) GetMetricNamespace() *string {
	return s.MetricNamespace
}

func (s *ListProductQuotasResponseBodyQuotasUsageMetric) SetMetricDimensions(v map[string]*string) *ListProductQuotasResponseBodyQuotasUsageMetric {
	s.MetricDimensions = v
	return s
}

func (s *ListProductQuotasResponseBodyQuotasUsageMetric) SetMetricName(v string) *ListProductQuotasResponseBodyQuotasUsageMetric {
	s.MetricName = &v
	return s
}

func (s *ListProductQuotasResponseBodyQuotasUsageMetric) SetMetricNamespace(v string) *ListProductQuotasResponseBodyQuotasUsageMetric {
	s.MetricNamespace = &v
	return s
}

func (s *ListProductQuotasResponseBodyQuotasUsageMetric) Validate() error {
	return dara.Validate(s)
}
