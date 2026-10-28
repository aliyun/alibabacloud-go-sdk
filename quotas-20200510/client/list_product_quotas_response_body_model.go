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
	// The maximum number of records returned for the current query.
	//
	// example:
	//
	// 1
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// Indicates whether a token for the next query is available.
	//
	// - If NextToken is empty, it indicates that no subsequent query is available.
	//
	// - If NextToken is returned, the value is the token used to start the next query.
	//
	// example:
	//
	// AAAAAd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The quota details.
	Quotas []*ListProductQuotasResponseBodyQuotas `json:"Quotas,omitempty" xml:"Quotas,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// D0131FD5-5397-44FE-BF5A-4B7165B813CC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of records returned for the current request.
	//
	// example:
	//
	// 1
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
	if s.Quotas != nil {
		for _, item := range s.Quotas {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListProductQuotasResponseBodyQuotas struct {
	// Indicates whether the quota is adjustable.
	//
	// example:
	//
	// true
	Adjustable *bool `json:"Adjustable,omitempty" xml:"Adjustable,omitempty"`
	// The range of quota values that can be applied for the current quota item, which serves as a reference for quota applications.
	ApplicableRange []*float32 `json:"ApplicableRange,omitempty" xml:"ApplicableRange,omitempty" type:"Repeated"`
	// The quota adjustment type.
	//
	// example:
	//
	// discontinuous
	ApplicableType *string `json:"ApplicableType,omitempty" xml:"ApplicableType,omitempty"`
	// The sample reasons for applying for a quota adjustment.
	//
	// example:
	//
	// The business xxx is expected to grow by 50%.
	ApplyReasonTips *string `json:"ApplyReasonTips,omitempty" xml:"ApplyReasonTips,omitempty"`
	// Indicates whether the used quota is displayed.
	//
	// example:
	//
	// true
	Consumable *bool `json:"Consumable,omitempty" xml:"Consumable,omitempty"`
	// The quota dimensions. Format: `{"regionId":"region"}`.
	//
	// example:
	//
	// {"regionId":"cn-hangzhou"}
	Dimensions map[string]interface{} `json:"Dimensions,omitempty" xml:"Dimensions,omitempty"`
	// The UTC time when the quota takes effect.
	//
	// example:
	//
	// 2022-09-28T06:07:00Z
	EffectiveTime *string `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	// The UTC time when the quota expires.
	//
	// example:
	//
	// 2022-09-29T06:07:00Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// Indicates whether the quota is a global quota.
	//
	// example:
	//
	// true
	GlobalQuota *bool `json:"GlobalQuota,omitempty" xml:"GlobalQuota,omitempty"`
	// The quota calculation period.
	Period *ListProductQuotasResponseBodyQuotasPeriod `json:"Period,omitempty" xml:"Period,omitempty" type:"Struct"`
	// The abbreviation of the cloud service name.
	//
	// example:
	//
	// ecs
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The ID of the quota.
	//
	// example:
	//
	// q_security-groups
	QuotaActionCode *string `json:"QuotaActionCode,omitempty" xml:"QuotaActionCode,omitempty"`
	// The ARN of the quota.
	//
	// example:
	//
	// acs:quotas:cn-hangzhou:120886317861****:quota/ecs/q_security-groups
	QuotaArn *string `json:"QuotaArn,omitempty" xml:"QuotaArn,omitempty"`
	// The quota type.
	//
	// example:
	//
	// CommonQuota
	QuotaCategory *string `json:"QuotaCategory,omitempty" xml:"QuotaCategory,omitempty"`
	// The quota description.
	//
	// example:
	//
	// The maximum number of security groups that can be owned by the current account.
	QuotaDescription *string `json:"QuotaDescription,omitempty" xml:"QuotaDescription,omitempty"`
	// The quota details.
	QuotaItems []*ListProductQuotasResponseBodyQuotasQuotaItems `json:"QuotaItems,omitempty" xml:"QuotaItems,omitempty" type:"Repeated"`
	// The quota name.
	//
	// example:
	//
	// ecs.g5.2xlarge
	QuotaName *string `json:"QuotaName,omitempty" xml:"QuotaName,omitempty"`
	// The quota type.
	//
	// example:
	//
	// privilege
	QuotaType *string `json:"QuotaType,omitempty" xml:"QuotaType,omitempty"`
	// The quota unit.
	//
	// example:
	//
	// AMOUNT
	QuotaUnit *string `json:"QuotaUnit,omitempty" xml:"QuotaUnit,omitempty"`
	// The range of quota values supported by the current quota item, which serves as a reference when you configure a quota template.
	SupportedRange []*float32 `json:"SupportedRange,omitempty" xml:"SupportedRange,omitempty" type:"Repeated"`
	// The quota value.
	//
	// example:
	//
	// 200
	TotalQuota *float32 `json:"TotalQuota,omitempty" xml:"TotalQuota,omitempty"`
	// The quota usage.
	//
	// example:
	//
	// 1009
	TotalUsage *float32 `json:"TotalUsage,omitempty" xml:"TotalUsage,omitempty"`
	// The reason why the quota is not adjustable.
	//
	// example:
	//
	// applicationProcess
	UnadjustableDetail *string `json:"UnadjustableDetail,omitempty" xml:"UnadjustableDetail,omitempty"`
	// The monitoring information of the quota usage in CloudMonitor.
	//
	// > An empty value indicates that no monitoring data is available for this quota in CloudMonitor.
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
	if s.Period != nil {
		if err := s.Period.Validate(); err != nil {
			return err
		}
	}
	if s.QuotaItems != nil {
		for _, item := range s.QuotaItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.UsageMetric != nil {
		if err := s.UsageMetric.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListProductQuotasResponseBodyQuotasPeriod struct {
	// The unit of the quota calculation period.
	//
	// example:
	//
	// day
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The value of the quota calculation period.
	//
	// example:
	//
	// {}
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
	// The quota value.
	//
	// example:
	//
	// 1009
	Quota *string `json:"Quota,omitempty" xml:"Quota,omitempty"`
	// The quota unit.
	//
	// example:
	//
	// Count
	QuotaUnit *string `json:"QuotaUnit,omitempty" xml:"QuotaUnit,omitempty"`
	// The quota category.
	//
	// example:
	//
	// BaseQuota
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The quota usage.
	//
	// example:
	//
	// 1009
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
	// The CloudMonitor monitoring dimensions.
	//
	// Format: a collection of `key:value` pairs. Example:
	//
	// `{"productCode":"***","metricKey":"***","regionId":"***","label":"***"}`.
	MetricDimensions map[string]*string `json:"MetricDimensions,omitempty" xml:"MetricDimensions,omitempty"`
	// The name of the CloudMonitor monitoring metric.
	//
	// example:
	//
	// Usage
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// The namespace of the CloudMonitor monitoring data.
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
