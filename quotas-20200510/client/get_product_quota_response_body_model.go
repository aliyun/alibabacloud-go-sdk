// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProductQuotaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetQuota(v *GetProductQuotaResponseBodyQuota) *GetProductQuotaResponseBody
	GetQuota() *GetProductQuotaResponseBodyQuota
	SetRequestId(v string) *GetProductQuotaResponseBody
	GetRequestId() *string
}

type GetProductQuotaResponseBody struct {
	// The quota details.
	Quota *GetProductQuotaResponseBodyQuota `json:"Quota,omitempty" xml:"Quota,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 8FF8CAF0-29D9-4F11-B6A4-FD2CBCA016D3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetProductQuotaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetProductQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *GetProductQuotaResponseBody) GetQuota() *GetProductQuotaResponseBodyQuota {
	return s.Quota
}

func (s *GetProductQuotaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetProductQuotaResponseBody) SetQuota(v *GetProductQuotaResponseBodyQuota) *GetProductQuotaResponseBody {
	s.Quota = v
	return s
}

func (s *GetProductQuotaResponseBody) SetRequestId(v string) *GetProductQuotaResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetProductQuotaResponseBody) Validate() error {
	if s.Quota != nil {
		if err := s.Quota.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetProductQuotaResponseBodyQuota struct {
	// Indicates whether the quota is adjustable. Valid values:
	//
	// - true: Adjustable.
	//
	// - false: Not adjustable.
	//
	// example:
	//
	// true
	Adjustable *bool `json:"Adjustable,omitempty" xml:"Adjustable,omitempty"`
	// The applicable range for quota adjustment.
	ApplicableRange []*float32 `json:"ApplicableRange,omitempty" xml:"ApplicableRange,omitempty" type:"Repeated"`
	// The type of quota adjustment. Valid values:
	//
	// - continuous: Continuous.
	//
	// - discontinuous: Discontinuous.
	//
	// example:
	//
	// continuous
	ApplicableType *string `json:"ApplicableType,omitempty" xml:"ApplicableType,omitempty"`
	// The sample reasons for applying for a quota adjustment.
	//
	// example:
	//
	// The business xxx is expected to grow by 50%.
	ApplyReasonTips *string `json:"ApplyReasonTips,omitempty" xml:"ApplyReasonTips,omitempty"`
	// Indicates whether the used quota is displayed. Valid values:
	//
	// - true: The used quota is displayed.
	//
	// - false: The used quota is not displayed.
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
	// 2022-09-28T06:06:00Z
	EffectiveTime *string `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	// The UTC time when the quota expires.
	//
	// example:
	//
	// 2022-09-29T06:06:00Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// Indicates whether the quota is a global quota. Valid values:
	//
	// - true: The quota item is shared across all regions.
	//
	// - false: The quota item is independent for each region.
	//
	// example:
	//
	// true
	GlobalQuota *bool `json:"GlobalQuota,omitempty" xml:"GlobalQuota,omitempty"`
	// The quota period.
	Period *GetProductQuotaResponseBodyQuotaPeriod `json:"Period,omitempty" xml:"Period,omitempty" type:"Struct"`
	// The abbreviated name of the cloud service.
	//
	// example:
	//
	// ecs
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The quota ID.
	//
	// example:
	//
	// q_security-groups
	QuotaActionCode *string `json:"QuotaActionCode,omitempty" xml:"QuotaActionCode,omitempty"`
	// The ARN of the quota.
	//
	// example:
	//
	// acs:quotas:cn-hangzhou:120886317861****:quota/ecs/q_security-groups/
	QuotaArn *string `json:"QuotaArn,omitempty" xml:"QuotaArn,omitempty"`
	// The quota category. Valid values:
	//
	// - CommonQuota: General quota.
	//
	// - FlowControl: API rate limit.
	//
	// - WhiteListLabel: Privilege.
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
	QuotaItems []*GetProductQuotaResponseBodyQuotaQuotaItems `json:"QuotaItems,omitempty" xml:"QuotaItems,omitempty" type:"Repeated"`
	// The quota name.
	//
	// example:
	//
	// Maximum Number of Security Groups
	QuotaName *string `json:"QuotaName,omitempty" xml:"QuotaName,omitempty"`
	// The quota type. Valid values:
	//
	// - privilege: Privilege.
	//
	// - normal: Normal.
	//
	// example:
	//
	// normal
	QuotaType *string `json:"QuotaType,omitempty" xml:"QuotaType,omitempty"`
	// The quota unit.
	//
	// > The unit varies for each quota. For example, the unit of the q_cbdch3 quota (maximum number of clusters) for Container Service for Kubernetes is Cluster, and the unit of the q_security-groups quota (maximum number of security groups) for Elastic Compute Service servers is count.
	//
	// example:
	//
	// Count
	QuotaUnit *string `json:"QuotaUnit,omitempty" xml:"QuotaUnit,omitempty"`
	// The supported range of quota values for the current quota item, which can be used as a reference when configuring a quota template.
	SupportedRange []*float32 `json:"SupportedRange,omitempty" xml:"SupportedRange,omitempty" type:"Repeated"`
	// The quota value.
	//
	// example:
	//
	// 801
	TotalQuota *float32 `json:"TotalQuota,omitempty" xml:"TotalQuota,omitempty"`
	// The quota usage.
	//
	// example:
	//
	// 26
	TotalUsage *float32 `json:"TotalUsage,omitempty" xml:"TotalUsage,omitempty"`
	// The reason why the quota cannot be adjusted. Valid values:
	//
	// - nonactivated: The service is not activated.
	//
	// - applicationProcess: The quota application is being processed.
	//
	// - limitReached: The quota limit is reached.
	//
	// - supportTicketRequired: A ticket must be submitted for the application.
	//
	// example:
	//
	// limitReached
	UnadjustableDetail *string `json:"UnadjustableDetail,omitempty" xml:"UnadjustableDetail,omitempty"`
	// The monitoring information of the quota usage in CloudMonitor.
	//
	// > An empty value indicates that no monitoring data is available for this quota in CloudMonitor.
	UsageMetric *GetProductQuotaResponseBodyQuotaUsageMetric `json:"UsageMetric,omitempty" xml:"UsageMetric,omitempty" type:"Struct"`
}

func (s GetProductQuotaResponseBodyQuota) String() string {
	return dara.Prettify(s)
}

func (s GetProductQuotaResponseBodyQuota) GoString() string {
	return s.String()
}

func (s *GetProductQuotaResponseBodyQuota) GetAdjustable() *bool {
	return s.Adjustable
}

func (s *GetProductQuotaResponseBodyQuota) GetApplicableRange() []*float32 {
	return s.ApplicableRange
}

func (s *GetProductQuotaResponseBodyQuota) GetApplicableType() *string {
	return s.ApplicableType
}

func (s *GetProductQuotaResponseBodyQuota) GetApplyReasonTips() *string {
	return s.ApplyReasonTips
}

func (s *GetProductQuotaResponseBodyQuota) GetConsumable() *bool {
	return s.Consumable
}

func (s *GetProductQuotaResponseBodyQuota) GetDimensions() map[string]interface{} {
	return s.Dimensions
}

func (s *GetProductQuotaResponseBodyQuota) GetEffectiveTime() *string {
	return s.EffectiveTime
}

func (s *GetProductQuotaResponseBodyQuota) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *GetProductQuotaResponseBodyQuota) GetGlobalQuota() *bool {
	return s.GlobalQuota
}

func (s *GetProductQuotaResponseBodyQuota) GetPeriod() *GetProductQuotaResponseBodyQuotaPeriod {
	return s.Period
}

func (s *GetProductQuotaResponseBodyQuota) GetProductCode() *string {
	return s.ProductCode
}

func (s *GetProductQuotaResponseBodyQuota) GetQuotaActionCode() *string {
	return s.QuotaActionCode
}

func (s *GetProductQuotaResponseBodyQuota) GetQuotaArn() *string {
	return s.QuotaArn
}

func (s *GetProductQuotaResponseBodyQuota) GetQuotaCategory() *string {
	return s.QuotaCategory
}

func (s *GetProductQuotaResponseBodyQuota) GetQuotaDescription() *string {
	return s.QuotaDescription
}

func (s *GetProductQuotaResponseBodyQuota) GetQuotaItems() []*GetProductQuotaResponseBodyQuotaQuotaItems {
	return s.QuotaItems
}

func (s *GetProductQuotaResponseBodyQuota) GetQuotaName() *string {
	return s.QuotaName
}

func (s *GetProductQuotaResponseBodyQuota) GetQuotaType() *string {
	return s.QuotaType
}

func (s *GetProductQuotaResponseBodyQuota) GetQuotaUnit() *string {
	return s.QuotaUnit
}

func (s *GetProductQuotaResponseBodyQuota) GetSupportedRange() []*float32 {
	return s.SupportedRange
}

func (s *GetProductQuotaResponseBodyQuota) GetTotalQuota() *float32 {
	return s.TotalQuota
}

func (s *GetProductQuotaResponseBodyQuota) GetTotalUsage() *float32 {
	return s.TotalUsage
}

func (s *GetProductQuotaResponseBodyQuota) GetUnadjustableDetail() *string {
	return s.UnadjustableDetail
}

func (s *GetProductQuotaResponseBodyQuota) GetUsageMetric() *GetProductQuotaResponseBodyQuotaUsageMetric {
	return s.UsageMetric
}

func (s *GetProductQuotaResponseBodyQuota) SetAdjustable(v bool) *GetProductQuotaResponseBodyQuota {
	s.Adjustable = &v
	return s
}

func (s *GetProductQuotaResponseBodyQuota) SetApplicableRange(v []*float32) *GetProductQuotaResponseBodyQuota {
	s.ApplicableRange = v
	return s
}

func (s *GetProductQuotaResponseBodyQuota) SetApplicableType(v string) *GetProductQuotaResponseBodyQuota {
	s.ApplicableType = &v
	return s
}

func (s *GetProductQuotaResponseBodyQuota) SetApplyReasonTips(v string) *GetProductQuotaResponseBodyQuota {
	s.ApplyReasonTips = &v
	return s
}

func (s *GetProductQuotaResponseBodyQuota) SetConsumable(v bool) *GetProductQuotaResponseBodyQuota {
	s.Consumable = &v
	return s
}

func (s *GetProductQuotaResponseBodyQuota) SetDimensions(v map[string]interface{}) *GetProductQuotaResponseBodyQuota {
	s.Dimensions = v
	return s
}

func (s *GetProductQuotaResponseBodyQuota) SetEffectiveTime(v string) *GetProductQuotaResponseBodyQuota {
	s.EffectiveTime = &v
	return s
}

func (s *GetProductQuotaResponseBodyQuota) SetExpireTime(v string) *GetProductQuotaResponseBodyQuota {
	s.ExpireTime = &v
	return s
}

func (s *GetProductQuotaResponseBodyQuota) SetGlobalQuota(v bool) *GetProductQuotaResponseBodyQuota {
	s.GlobalQuota = &v
	return s
}

func (s *GetProductQuotaResponseBodyQuota) SetPeriod(v *GetProductQuotaResponseBodyQuotaPeriod) *GetProductQuotaResponseBodyQuota {
	s.Period = v
	return s
}

func (s *GetProductQuotaResponseBodyQuota) SetProductCode(v string) *GetProductQuotaResponseBodyQuota {
	s.ProductCode = &v
	return s
}

func (s *GetProductQuotaResponseBodyQuota) SetQuotaActionCode(v string) *GetProductQuotaResponseBodyQuota {
	s.QuotaActionCode = &v
	return s
}

func (s *GetProductQuotaResponseBodyQuota) SetQuotaArn(v string) *GetProductQuotaResponseBodyQuota {
	s.QuotaArn = &v
	return s
}

func (s *GetProductQuotaResponseBodyQuota) SetQuotaCategory(v string) *GetProductQuotaResponseBodyQuota {
	s.QuotaCategory = &v
	return s
}

func (s *GetProductQuotaResponseBodyQuota) SetQuotaDescription(v string) *GetProductQuotaResponseBodyQuota {
	s.QuotaDescription = &v
	return s
}

func (s *GetProductQuotaResponseBodyQuota) SetQuotaItems(v []*GetProductQuotaResponseBodyQuotaQuotaItems) *GetProductQuotaResponseBodyQuota {
	s.QuotaItems = v
	return s
}

func (s *GetProductQuotaResponseBodyQuota) SetQuotaName(v string) *GetProductQuotaResponseBodyQuota {
	s.QuotaName = &v
	return s
}

func (s *GetProductQuotaResponseBodyQuota) SetQuotaType(v string) *GetProductQuotaResponseBodyQuota {
	s.QuotaType = &v
	return s
}

func (s *GetProductQuotaResponseBodyQuota) SetQuotaUnit(v string) *GetProductQuotaResponseBodyQuota {
	s.QuotaUnit = &v
	return s
}

func (s *GetProductQuotaResponseBodyQuota) SetSupportedRange(v []*float32) *GetProductQuotaResponseBodyQuota {
	s.SupportedRange = v
	return s
}

func (s *GetProductQuotaResponseBodyQuota) SetTotalQuota(v float32) *GetProductQuotaResponseBodyQuota {
	s.TotalQuota = &v
	return s
}

func (s *GetProductQuotaResponseBodyQuota) SetTotalUsage(v float32) *GetProductQuotaResponseBodyQuota {
	s.TotalUsage = &v
	return s
}

func (s *GetProductQuotaResponseBodyQuota) SetUnadjustableDetail(v string) *GetProductQuotaResponseBodyQuota {
	s.UnadjustableDetail = &v
	return s
}

func (s *GetProductQuotaResponseBodyQuota) SetUsageMetric(v *GetProductQuotaResponseBodyQuotaUsageMetric) *GetProductQuotaResponseBodyQuota {
	s.UsageMetric = v
	return s
}

func (s *GetProductQuotaResponseBodyQuota) Validate() error {
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

type GetProductQuotaResponseBodyQuotaPeriod struct {
	// The unit of the quota period. Valid values:
	//
	// - second: Seconds.
	//
	// - minute: Minutes.
	//
	// - hour: Hours.
	//
	// - day: Days.
	//
	// - week: Weeks.
	//
	// example:
	//
	// day
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The value of the quota period.
	//
	// example:
	//
	// 1
	PeriodValue *int32 `json:"PeriodValue,omitempty" xml:"PeriodValue,omitempty"`
}

func (s GetProductQuotaResponseBodyQuotaPeriod) String() string {
	return dara.Prettify(s)
}

func (s GetProductQuotaResponseBodyQuotaPeriod) GoString() string {
	return s.String()
}

func (s *GetProductQuotaResponseBodyQuotaPeriod) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *GetProductQuotaResponseBodyQuotaPeriod) GetPeriodValue() *int32 {
	return s.PeriodValue
}

func (s *GetProductQuotaResponseBodyQuotaPeriod) SetPeriodUnit(v string) *GetProductQuotaResponseBodyQuotaPeriod {
	s.PeriodUnit = &v
	return s
}

func (s *GetProductQuotaResponseBodyQuotaPeriod) SetPeriodValue(v int32) *GetProductQuotaResponseBodyQuotaPeriod {
	s.PeriodValue = &v
	return s
}

func (s *GetProductQuotaResponseBodyQuotaPeriod) Validate() error {
	return dara.Validate(s)
}

type GetProductQuotaResponseBodyQuotaQuotaItems struct {
	// The quota value.
	//
	// example:
	//
	// 801
	Quota *string `json:"Quota,omitempty" xml:"Quota,omitempty"`
	// The quota unit.
	//
	// > The unit varies for each quota. For example, the unit of the q_cbdch3 quota (maximum number of clusters) for Container Service for Kubernetes is Cluster, and the unit of the q_security-groups quota (maximum number of security groups) for Elastic Compute Service servers is count.
	//
	// example:
	//
	// Count
	QuotaUnit *string `json:"QuotaUnit,omitempty" xml:"QuotaUnit,omitempty"`
	// The quota item type. Valid values:
	//
	// - BaseQuota: Base quota.
	//
	// - ReservedQuota: Reserved quota.
	//
	// example:
	//
	// BaseQuota
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The quota usage.
	//
	// example:
	//
	// 26
	Usage *string `json:"Usage,omitempty" xml:"Usage,omitempty"`
}

func (s GetProductQuotaResponseBodyQuotaQuotaItems) String() string {
	return dara.Prettify(s)
}

func (s GetProductQuotaResponseBodyQuotaQuotaItems) GoString() string {
	return s.String()
}

func (s *GetProductQuotaResponseBodyQuotaQuotaItems) GetQuota() *string {
	return s.Quota
}

func (s *GetProductQuotaResponseBodyQuotaQuotaItems) GetQuotaUnit() *string {
	return s.QuotaUnit
}

func (s *GetProductQuotaResponseBodyQuotaQuotaItems) GetType() *string {
	return s.Type
}

func (s *GetProductQuotaResponseBodyQuotaQuotaItems) GetUsage() *string {
	return s.Usage
}

func (s *GetProductQuotaResponseBodyQuotaQuotaItems) SetQuota(v string) *GetProductQuotaResponseBodyQuotaQuotaItems {
	s.Quota = &v
	return s
}

func (s *GetProductQuotaResponseBodyQuotaQuotaItems) SetQuotaUnit(v string) *GetProductQuotaResponseBodyQuotaQuotaItems {
	s.QuotaUnit = &v
	return s
}

func (s *GetProductQuotaResponseBodyQuotaQuotaItems) SetType(v string) *GetProductQuotaResponseBodyQuotaQuotaItems {
	s.Type = &v
	return s
}

func (s *GetProductQuotaResponseBodyQuotaQuotaItems) SetUsage(v string) *GetProductQuotaResponseBodyQuotaQuotaItems {
	s.Usage = &v
	return s
}

func (s *GetProductQuotaResponseBodyQuotaQuotaItems) Validate() error {
	return dara.Validate(s)
}

type GetProductQuotaResponseBodyQuotaUsageMetric struct {
	// The CloudMonitor monitoring dimensions.
	//
	// Format: a collection of `key:value` pairs. Example:
	//
	// `{"productCode":"***","metricKey":"***","regionId":"***","label":"***"}`.
	MetricDimensions map[string]*string `json:"MetricDimensions,omitempty" xml:"MetricDimensions,omitempty"`
	// The CloudMonitor metric name.
	//
	// example:
	//
	// Usage
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// The CloudMonitor metric namespace.
	//
	// example:
	//
	// acs_quotas_flowcontrol
	MetricNamespace *string `json:"MetricNamespace,omitempty" xml:"MetricNamespace,omitempty"`
}

func (s GetProductQuotaResponseBodyQuotaUsageMetric) String() string {
	return dara.Prettify(s)
}

func (s GetProductQuotaResponseBodyQuotaUsageMetric) GoString() string {
	return s.String()
}

func (s *GetProductQuotaResponseBodyQuotaUsageMetric) GetMetricDimensions() map[string]*string {
	return s.MetricDimensions
}

func (s *GetProductQuotaResponseBodyQuotaUsageMetric) GetMetricName() *string {
	return s.MetricName
}

func (s *GetProductQuotaResponseBodyQuotaUsageMetric) GetMetricNamespace() *string {
	return s.MetricNamespace
}

func (s *GetProductQuotaResponseBodyQuotaUsageMetric) SetMetricDimensions(v map[string]*string) *GetProductQuotaResponseBodyQuotaUsageMetric {
	s.MetricDimensions = v
	return s
}

func (s *GetProductQuotaResponseBodyQuotaUsageMetric) SetMetricName(v string) *GetProductQuotaResponseBodyQuotaUsageMetric {
	s.MetricName = &v
	return s
}

func (s *GetProductQuotaResponseBodyQuotaUsageMetric) SetMetricNamespace(v string) *GetProductQuotaResponseBodyQuotaUsageMetric {
	s.MetricNamespace = &v
	return s
}

func (s *GetProductQuotaResponseBodyQuotaUsageMetric) Validate() error {
	return dara.Validate(s)
}
