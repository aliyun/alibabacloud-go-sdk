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
	// The details of the quota.
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
	return dara.Validate(s)
}

type GetProductQuotaResponseBodyQuota struct {
	// Indicates whether the quota is adjustable. Valid values:
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
	// The type of the adjustable value. Valid values:
	//
	// 	- continuous
	//
	// 	- discontinuous
	//
	// example:
	//
	// continuous
	ApplicableType *string `json:"ApplicableType,omitempty" xml:"ApplicableType,omitempty"`
	// The reason for submitting a quota increase request.
	//
	// example:
	//
	// The business xxx is expected to grow by 50%.
	ApplyReasonTips *string `json:"ApplyReasonTips,omitempty" xml:"ApplyReasonTips,omitempty"`
	// Indicates whether the system shows the used value of the quota. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Consumable *bool `json:"Consumable,omitempty" xml:"Consumable,omitempty"`
	// The quota dimensions. Format: `{"regionId":"Region"}`.
	//
	// example:
	//
	// {"regionId":"cn-hangzhou"}
	Dimensions map[string]interface{} `json:"Dimensions,omitempty" xml:"Dimensions,omitempty"`
	// The start time of the validity period of the quota. Specify the value in UTC.
	//
	// example:
	//
	// 2022-09-28T06:06:00Z
	EffectiveTime *string `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	// The end time of the validity period of the quota. Specify the value in UTC.
	//
	// example:
	//
	// 2022-09-29T06:06:00Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// Indicates whether the quota is a global quota. Valid values:
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
	Period *GetProductQuotaResponseBodyQuotaPeriod `json:"Period,omitempty" xml:"Period,omitempty" type:"Struct"`
	// The abbreviation of the Alibaba Cloud service name.
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
	// The Alibaba Cloud Resource Name (ARN) of the quota.
	//
	// example:
	//
	// acs:quotas:cn-hangzhou:120886317861****:quota/ecs/q_security-groups/
	QuotaArn *string `json:"QuotaArn,omitempty" xml:"QuotaArn,omitempty"`
	// The type of the quota. Valid values:
	//
	// 	- CommonQuota: general quota
	//
	// 	- FlowControl: API rate limit
	//
	// 	- WhiteListLabel: whitelist quota
	//
	// example:
	//
	// CommonQuota
	QuotaCategory *string `json:"QuotaCategory,omitempty" xml:"QuotaCategory,omitempty"`
	// The description of the quota.
	//
	// example:
	//
	// The maximum number of security groups that can be owned by the current account
	QuotaDescription *string `json:"QuotaDescription,omitempty" xml:"QuotaDescription,omitempty"`
	// The details of the quota.
	QuotaItems []*GetProductQuotaResponseBodyQuotaQuotaItems `json:"QuotaItems,omitempty" xml:"QuotaItems,omitempty" type:"Repeated"`
	// The name of the quota.
	//
	// example:
	//
	// Maximum Number of Security Groups
	QuotaName *string `json:"QuotaName,omitempty" xml:"QuotaName,omitempty"`
	// The type of the quota. Valid values:
	//
	// 	- privilege
	//
	// 	- normal
	//
	// example:
	//
	// normal
	QuotaType *string `json:"QuotaType,omitempty" xml:"QuotaType,omitempty"`
	// The unit of the quota.
	//
	// >  The unit of each quota is unique. For example, the quota whose ID is `q_cbdch3` represents the maximum number of Container Service for Kubernetes (ACK) clusters. The unit of this quota is clusters. The quota whose ID is `q_security-groups` represents the maximum number of security groups. The unit of this quota is security groups.
	//
	// example:
	//
	// Count
	QuotaUnit *string `json:"QuotaUnit,omitempty" xml:"QuotaUnit,omitempty"`
	// The range of the quota value.
	SupportedRange []*float32 `json:"SupportedRange,omitempty" xml:"SupportedRange,omitempty" type:"Repeated"`
	// The value of the quota.
	//
	// example:
	//
	// 801
	TotalQuota *float32 `json:"TotalQuota,omitempty" xml:"TotalQuota,omitempty"`
	// The used quota.
	//
	// example:
	//
	// 26
	TotalUsage *float32 `json:"TotalUsage,omitempty" xml:"TotalUsage,omitempty"`
	// The reason why the quota is not adjustable. Valid values:
	//
	// 	- nonactivated: The service is not activated.
	//
	// 	- applicationProcess: The application is being processed.
	//
	// 	- limitReached: The quota limit is reached.
	//
	// 	- supportTicketRequired: The quota can be increased only by submitting a ticket.
	//
	// example:
	//
	// limitReached
	UnadjustableDetail *string `json:"UnadjustableDetail,omitempty" xml:"UnadjustableDetail,omitempty"`
	// The monitoring information of the quota in CloudMonitor.
	//
	// >  If this parameter is empty, no monitoring data of the quota exists in CloudMonitor.
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
	return dara.Validate(s)
}

type GetProductQuotaResponseBodyQuotaPeriod struct {
	// The unit of the calculation cycle of the quota. Valid values:
	//
	// 	- second
	//
	// 	- minute
	//
	// 	- hour
	//
	// 	- day
	//
	// 	- week
	//
	// example:
	//
	// Day
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The value of the calculation cycle of the quota.
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
	// The value of the quota.
	//
	// example:
	//
	// 801
	Quota *string `json:"Quota,omitempty" xml:"Quota,omitempty"`
	// The unit of the quota.
	//
	// >  The unit of each quota is unique. For example, the quota whose ID is `q_cbdch3` represents the maximum number of ACK clusters. The unit of this quota is clusters. The quota whose ID is `q_security-groups` represents the maximum number of security groups. The unit of this quota is security groups.
	//
	// example:
	//
	// Count
	QuotaUnit *string `json:"QuotaUnit,omitempty" xml:"QuotaUnit,omitempty"`
	// The category of the quota. Valid values:
	//
	// 	- BaseQuota: base quota.
	//
	// 	- ReservedQuota: reserved quota.
	//
	// example:
	//
	// BaseQuota
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The used quota.
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
