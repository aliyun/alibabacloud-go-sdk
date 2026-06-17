// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutResourceMetricRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRules(v []*PutResourceMetricRulesRequestRules) *PutResourceMetricRulesRequest
	GetRules() []*PutResourceMetricRulesRequestRules
}

type PutResourceMetricRulesRequest struct {
	// The list of threshold alert rules.
	//
	// Valid values of N: 1 to 50.
	//
	// This parameter is required.
	Rules []*PutResourceMetricRulesRequestRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
}

func (s PutResourceMetricRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s PutResourceMetricRulesRequest) GoString() string {
	return s.String()
}

func (s *PutResourceMetricRulesRequest) GetRules() []*PutResourceMetricRulesRequestRules {
	return s.Rules
}

func (s *PutResourceMetricRulesRequest) SetRules(v []*PutResourceMetricRulesRequestRules) *PutResourceMetricRulesRequest {
	s.Rules = v
	return s
}

func (s *PutResourceMetricRulesRequest) Validate() error {
	if s.Rules != nil {
		for _, item := range s.Rules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type PutResourceMetricRulesRequestRules struct {
	Escalations *PutResourceMetricRulesRequestRulesEscalations `json:"Escalations,omitempty" xml:"Escalations,omitempty" type:"Struct"`
	// 报警联系组。报警通知会发送给该报警联系组中的报警联系人。
	//
	// N的取值范围：1~50。
	//
	// > 报警联系组是一组报警联系人，可以包含一个或多个报警联系人。关于如何创建报警联系人和报警联系组，请参见[PutContact](https://help.aliyun.com/document_detail/114923.html)和[PutContactGroup](https://help.aliyun.com/document_detail/114929.html)。
	//
	// This parameter is required.
	//
	// example:
	//
	// ECS_Group
	ContactGroups *string `json:"ContactGroups,omitempty" xml:"ContactGroups,omitempty"`
	// 报警规则的生效时间范围。
	//
	// N的取值范围：1~50。
	//
	// example:
	//
	// 00:00-23:59
	EffectiveInterval *string `json:"EffectiveInterval,omitempty" xml:"EffectiveInterval,omitempty"`
	// 报警邮件主题。
	//
	// N的取值范围：1~50。
	//
	// example:
	//
	// ECS instance alert
	EmailSubject *string `json:"EmailSubject,omitempty" xml:"EmailSubject,omitempty"`
	// 报警规则的触发周期。
	//
	// 单位：秒。
	//
	// N的取值范围：1~50。
	//
	// >关于如何查询监控项的统计周期，请参见[云产品监控项](https://help.aliyun.com/document_detail/163515.html)。
	//
	// example:
	//
	// 60
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// 当监控项达到报警条件并进行报警时，标签同时写入监控项，在报警通知中进行展示。
	Labels []*PutResourceMetricRulesRequestRulesLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// 监控项名称。
	//
	// N的取值范围：1~50。
	//
	// 关于如何查询监控项名称，请参见[云产品监控项](https://help.aliyun.com/document_detail/163515.html)。
	//
	// example:
	//
	// cpu_total
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// 云产品的数据命名空间。
	//
	// N的取值范围：1~50。
	//
	// 关于如何查询云产品的数据命名空间，请参见[云产品监控项](https://help.aliyun.com/document_detail/163515.html)。
	//
	// This parameter is required.
	//
	// example:
	//
	// acs_ecs_dashboard
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// 无监控数据时报警的处理方式。取值：
	//
	// - KEEP_LAST_STATE（默认值）：不做任何处理。
	//
	// - INSUFFICIENT_DATA：报警内容为无数据。
	//
	// - OK：正常。
	//
	// N的取值范围：1~50。
	//
	// example:
	//
	// KEEP_LAST_STATE
	NoDataPolicy *string `json:"NoDataPolicy,omitempty" xml:"NoDataPolicy,omitempty"`
	// 报警规则的失效时间范围。
	//
	// N的取值范围：1~50。
	//
	// example:
	//
	// 00:00-06:00
	NoEffectiveInterval *string `json:"NoEffectiveInterval,omitempty" xml:"NoEffectiveInterval,omitempty"`
	// 监控项的统计周期。
	//
	// 单位：秒。默认为监控项的原始上报周期。
	//
	// N的取值范围：1~50。
	//
	// >关于如何查询监控项的统计周期，请参见[云产品监控项](https://help.aliyun.com/document_detail/163515.html)。
	//
	// example:
	//
	// 60
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// 资源信息，例如：`[{"instanceId":"i-uf6j91r34rnwawoo****"}]`、`[{"userId":"100931896542****"}]`。
	//
	// N的取值范围：1~50。
	//
	// 关于资源信息支持的维度Dimensions，请参见[云产品监控项](https://help.aliyun.com/document_detail/163515.html)。
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"instanceId":"i-uf6j91r34rnwawoo****"}]
	Resources *string `json:"Resources,omitempty" xml:"Resources,omitempty"`
	// 报警规则ID。
	//
	// N的取值范围：1~50。
	//
	// 您可以输入新的报警规则ID，也可以使用云监控已存在的报警规则ID。关于如何查询报警规则ID，请参见[DescribeMetricRuleList](https://help.aliyun.com/document_detail/114941.html)。
	//
	// > 输入新的报警规则ID，表示创建一条阈值报警规则。
	//
	// This parameter is required.
	//
	// example:
	//
	// a151cd6023eacee2f0978e03863cc1697c89508****
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// 报警规则名称。
	//
	// N的取值范围：1~50。
	//
	// 您可以输入新的报警规则名称，也可以使用云监控已存在的报警规则名称。关于如何查询报警规则名称，请参见[DescribeMetricRuleList](https://help.aliyun.com/document_detail/114941.html)。
	//
	// > 输入新的报警规则名称，表示创建一条阈值报警规则。
	//
	// This parameter is required.
	//
	// example:
	//
	// test123
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// 是否发送恢复通知
	//
	// example:
	//
	// true
	SendOK *bool `json:"SendOK,omitempty" xml:"SendOK,omitempty"`
	// 通道沉默周期。
	//
	// 单位：秒，默认值：86400。
	//
	// N的取值范围：1~50。
	//
	// > 通道沉默周期是指报警发生后未恢复正常，间隔多久重新发送一次报警通知。
	//
	// example:
	//
	// 86400
	SilenceTime *int32 `json:"SilenceTime,omitempty" xml:"SilenceTime,omitempty"`
	// 报警发生回调时指定的URL地址，向URL发送POST请求。
	//
	// N的取值范围：1~50。
	//
	// example:
	//
	// https://alert.aliyun.com.com:8080/callback
	Webhook *string `json:"Webhook,omitempty" xml:"Webhook,omitempty"`
}

func (s PutResourceMetricRulesRequestRules) String() string {
	return dara.Prettify(s)
}

func (s PutResourceMetricRulesRequestRules) GoString() string {
	return s.String()
}

func (s *PutResourceMetricRulesRequestRules) GetEscalations() *PutResourceMetricRulesRequestRulesEscalations {
	return s.Escalations
}

func (s *PutResourceMetricRulesRequestRules) GetContactGroups() *string {
	return s.ContactGroups
}

func (s *PutResourceMetricRulesRequestRules) GetEffectiveInterval() *string {
	return s.EffectiveInterval
}

func (s *PutResourceMetricRulesRequestRules) GetEmailSubject() *string {
	return s.EmailSubject
}

func (s *PutResourceMetricRulesRequestRules) GetInterval() *string {
	return s.Interval
}

func (s *PutResourceMetricRulesRequestRules) GetLabels() []*PutResourceMetricRulesRequestRulesLabels {
	return s.Labels
}

func (s *PutResourceMetricRulesRequestRules) GetMetricName() *string {
	return s.MetricName
}

func (s *PutResourceMetricRulesRequestRules) GetNamespace() *string {
	return s.Namespace
}

func (s *PutResourceMetricRulesRequestRules) GetNoDataPolicy() *string {
	return s.NoDataPolicy
}

func (s *PutResourceMetricRulesRequestRules) GetNoEffectiveInterval() *string {
	return s.NoEffectiveInterval
}

func (s *PutResourceMetricRulesRequestRules) GetPeriod() *string {
	return s.Period
}

func (s *PutResourceMetricRulesRequestRules) GetResources() *string {
	return s.Resources
}

func (s *PutResourceMetricRulesRequestRules) GetRuleId() *string {
	return s.RuleId
}

func (s *PutResourceMetricRulesRequestRules) GetRuleName() *string {
	return s.RuleName
}

func (s *PutResourceMetricRulesRequestRules) GetSendOK() *bool {
	return s.SendOK
}

func (s *PutResourceMetricRulesRequestRules) GetSilenceTime() *int32 {
	return s.SilenceTime
}

func (s *PutResourceMetricRulesRequestRules) GetWebhook() *string {
	return s.Webhook
}

func (s *PutResourceMetricRulesRequestRules) SetEscalations(v *PutResourceMetricRulesRequestRulesEscalations) *PutResourceMetricRulesRequestRules {
	s.Escalations = v
	return s
}

func (s *PutResourceMetricRulesRequestRules) SetContactGroups(v string) *PutResourceMetricRulesRequestRules {
	s.ContactGroups = &v
	return s
}

func (s *PutResourceMetricRulesRequestRules) SetEffectiveInterval(v string) *PutResourceMetricRulesRequestRules {
	s.EffectiveInterval = &v
	return s
}

func (s *PutResourceMetricRulesRequestRules) SetEmailSubject(v string) *PutResourceMetricRulesRequestRules {
	s.EmailSubject = &v
	return s
}

func (s *PutResourceMetricRulesRequestRules) SetInterval(v string) *PutResourceMetricRulesRequestRules {
	s.Interval = &v
	return s
}

func (s *PutResourceMetricRulesRequestRules) SetLabels(v []*PutResourceMetricRulesRequestRulesLabels) *PutResourceMetricRulesRequestRules {
	s.Labels = v
	return s
}

func (s *PutResourceMetricRulesRequestRules) SetMetricName(v string) *PutResourceMetricRulesRequestRules {
	s.MetricName = &v
	return s
}

func (s *PutResourceMetricRulesRequestRules) SetNamespace(v string) *PutResourceMetricRulesRequestRules {
	s.Namespace = &v
	return s
}

func (s *PutResourceMetricRulesRequestRules) SetNoDataPolicy(v string) *PutResourceMetricRulesRequestRules {
	s.NoDataPolicy = &v
	return s
}

func (s *PutResourceMetricRulesRequestRules) SetNoEffectiveInterval(v string) *PutResourceMetricRulesRequestRules {
	s.NoEffectiveInterval = &v
	return s
}

func (s *PutResourceMetricRulesRequestRules) SetPeriod(v string) *PutResourceMetricRulesRequestRules {
	s.Period = &v
	return s
}

func (s *PutResourceMetricRulesRequestRules) SetResources(v string) *PutResourceMetricRulesRequestRules {
	s.Resources = &v
	return s
}

func (s *PutResourceMetricRulesRequestRules) SetRuleId(v string) *PutResourceMetricRulesRequestRules {
	s.RuleId = &v
	return s
}

func (s *PutResourceMetricRulesRequestRules) SetRuleName(v string) *PutResourceMetricRulesRequestRules {
	s.RuleName = &v
	return s
}

func (s *PutResourceMetricRulesRequestRules) SetSendOK(v bool) *PutResourceMetricRulesRequestRules {
	s.SendOK = &v
	return s
}

func (s *PutResourceMetricRulesRequestRules) SetSilenceTime(v int32) *PutResourceMetricRulesRequestRules {
	s.SilenceTime = &v
	return s
}

func (s *PutResourceMetricRulesRequestRules) SetWebhook(v string) *PutResourceMetricRulesRequestRules {
	s.Webhook = &v
	return s
}

func (s *PutResourceMetricRulesRequestRules) Validate() error {
	if s.Escalations != nil {
		if err := s.Escalations.Validate(); err != nil {
			return err
		}
	}
	if s.Labels != nil {
		for _, item := range s.Labels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type PutResourceMetricRulesRequestRulesEscalations struct {
	Critical *PutResourceMetricRulesRequestRulesEscalationsCritical `json:"Critical,omitempty" xml:"Critical,omitempty" type:"Struct"`
	Info     *PutResourceMetricRulesRequestRulesEscalationsInfo     `json:"Info,omitempty" xml:"Info,omitempty" type:"Struct"`
	Warn     *PutResourceMetricRulesRequestRulesEscalationsWarn     `json:"Warn,omitempty" xml:"Warn,omitempty" type:"Struct"`
}

func (s PutResourceMetricRulesRequestRulesEscalations) String() string {
	return dara.Prettify(s)
}

func (s PutResourceMetricRulesRequestRulesEscalations) GoString() string {
	return s.String()
}

func (s *PutResourceMetricRulesRequestRulesEscalations) GetCritical() *PutResourceMetricRulesRequestRulesEscalationsCritical {
	return s.Critical
}

func (s *PutResourceMetricRulesRequestRulesEscalations) GetInfo() *PutResourceMetricRulesRequestRulesEscalationsInfo {
	return s.Info
}

func (s *PutResourceMetricRulesRequestRulesEscalations) GetWarn() *PutResourceMetricRulesRequestRulesEscalationsWarn {
	return s.Warn
}

func (s *PutResourceMetricRulesRequestRulesEscalations) SetCritical(v *PutResourceMetricRulesRequestRulesEscalationsCritical) *PutResourceMetricRulesRequestRulesEscalations {
	s.Critical = v
	return s
}

func (s *PutResourceMetricRulesRequestRulesEscalations) SetInfo(v *PutResourceMetricRulesRequestRulesEscalationsInfo) *PutResourceMetricRulesRequestRulesEscalations {
	s.Info = v
	return s
}

func (s *PutResourceMetricRulesRequestRulesEscalations) SetWarn(v *PutResourceMetricRulesRequestRulesEscalationsWarn) *PutResourceMetricRulesRequestRulesEscalations {
	s.Warn = v
	return s
}

func (s *PutResourceMetricRulesRequestRulesEscalations) Validate() error {
	if s.Critical != nil {
		if err := s.Critical.Validate(); err != nil {
			return err
		}
	}
	if s.Info != nil {
		if err := s.Info.Validate(); err != nil {
			return err
		}
	}
	if s.Warn != nil {
		if err := s.Warn.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PutResourceMetricRulesRequestRulesEscalationsCritical struct {
	// Critical级别阈值比较符。取值：
	//
	// - GreaterThanOrEqualToThreshold：大于等于。
	//
	// - GreaterThanThreshold：大于。
	//
	// - LessThanOrEqualToThreshold：小于等于。
	//
	// - LessThanThreshold：小于。
	//
	// - NotEqualToThreshold：不等于。
	//
	// - EqualToThreshold：等于。
	//
	// - GreaterThanYesterday：同比昨天时间上涨。
	//
	// - LessThanYesterday：同比昨天时间下降。
	//
	// - GreaterThanLastWeek：同比上周同一时间上涨。
	//
	// - LessThanLastWeek：同比上周同一时间下降。
	//
	// - GreaterThanLastPeriod：环比上周期上涨。
	//
	// - LessThanLastPeriod：环比上周期下降。
	//
	// N的取值范围：1~50。
	//
	// > 报警级别Critical（严重）、Warn（警告）或Info（信息）至少设置一个，且该报警级别中的参数Statistics、ComparisonOperator、Threshold和Times必须同时设置。
	//
	// example:
	//
	// GreaterThanOrEqualToThreshold
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	N                  *int32  `json:"N,omitempty" xml:"N,omitempty"`
	PreCondition       *string `json:"PreCondition,omitempty" xml:"PreCondition,omitempty"`
	// Critical级别报警统计方法。
	//
	// N的取值范围：1~50。
	//
	// 该参数的取值由指定云产品的`MetricName`对应的`Statistics`列决定，例如：Maximum、Minimum和Average。关于如何获取该参数的取值，请参见[云产品监控项](https://help.aliyun.com/document_detail/163515.html)。
	//
	// > 报警级别Critical（严重）、Warn（警告）或Info（信息）至少设置一个，且该报警级别中的参数Statistics、ComparisonOperator、Threshold和Times必须同时设置。
	//
	// example:
	//
	// Average
	Statistics *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	// Critical级别报警阈值。
	//
	// N的取值范围：1~50。
	//
	//
	// > 报警级别Critical（严重）、Warn（警告）或Info（信息）至少设置一个，且该报警级别中的参数Statistics、ComparisonOperator、Threshold和Times必须同时设置。
	//
	// example:
	//
	// 90
	Threshold *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	// Critical级别报警重试次数。
	//
	// N的取值范围：1~50。
	//
	// > 报警级别Critical（严重）、Warn（警告）或Info（信息）至少设置一个，且该报警级别中的参数Statistics、ComparisonOperator、Threshold和Times必须同时设置。
	//
	// example:
	//
	// 3
	Times *int32 `json:"Times,omitempty" xml:"Times,omitempty"`
}

func (s PutResourceMetricRulesRequestRulesEscalationsCritical) String() string {
	return dara.Prettify(s)
}

func (s PutResourceMetricRulesRequestRulesEscalationsCritical) GoString() string {
	return s.String()
}

func (s *PutResourceMetricRulesRequestRulesEscalationsCritical) GetComparisonOperator() *string {
	return s.ComparisonOperator
}

func (s *PutResourceMetricRulesRequestRulesEscalationsCritical) GetN() *int32 {
	return s.N
}

func (s *PutResourceMetricRulesRequestRulesEscalationsCritical) GetPreCondition() *string {
	return s.PreCondition
}

func (s *PutResourceMetricRulesRequestRulesEscalationsCritical) GetStatistics() *string {
	return s.Statistics
}

func (s *PutResourceMetricRulesRequestRulesEscalationsCritical) GetThreshold() *string {
	return s.Threshold
}

func (s *PutResourceMetricRulesRequestRulesEscalationsCritical) GetTimes() *int32 {
	return s.Times
}

func (s *PutResourceMetricRulesRequestRulesEscalationsCritical) SetComparisonOperator(v string) *PutResourceMetricRulesRequestRulesEscalationsCritical {
	s.ComparisonOperator = &v
	return s
}

func (s *PutResourceMetricRulesRequestRulesEscalationsCritical) SetN(v int32) *PutResourceMetricRulesRequestRulesEscalationsCritical {
	s.N = &v
	return s
}

func (s *PutResourceMetricRulesRequestRulesEscalationsCritical) SetPreCondition(v string) *PutResourceMetricRulesRequestRulesEscalationsCritical {
	s.PreCondition = &v
	return s
}

func (s *PutResourceMetricRulesRequestRulesEscalationsCritical) SetStatistics(v string) *PutResourceMetricRulesRequestRulesEscalationsCritical {
	s.Statistics = &v
	return s
}

func (s *PutResourceMetricRulesRequestRulesEscalationsCritical) SetThreshold(v string) *PutResourceMetricRulesRequestRulesEscalationsCritical {
	s.Threshold = &v
	return s
}

func (s *PutResourceMetricRulesRequestRulesEscalationsCritical) SetTimes(v int32) *PutResourceMetricRulesRequestRulesEscalationsCritical {
	s.Times = &v
	return s
}

func (s *PutResourceMetricRulesRequestRulesEscalationsCritical) Validate() error {
	return dara.Validate(s)
}

type PutResourceMetricRulesRequestRulesEscalationsInfo struct {
	// Info级别阈值比较符。取值：
	//
	// - GreaterThanOrEqualToThreshold：大于等于。
	//
	// - GreaterThanThreshold：大于。
	//
	// - LessThanOrEqualToThreshold：小于等于。
	//
	// - LessThanThreshold：小于。
	//
	// - NotEqualToThreshold：不等于。
	//
	// - EqualToThreshold：等于。
	//
	// - GreaterThanYesterday：同比昨天时间上涨。
	//
	// - LessThanYesterday：同比昨天时间下降。
	//
	// - GreaterThanLastWeek：同比上周同一时间上涨。
	//
	// - LessThanLastWeek：同比上周同一时间下降。
	//
	// - GreaterThanLastPeriod：环比上周期上涨。
	//
	// - LessThanLastPeriod：环比上周期下降。
	//
	// N的取值范围：1~50。
	//
	// > 报警级别Critical（严重）、Warn（警告）或Info（信息）至少设置一个，且该报警级别中的参数Statistics、ComparisonOperator、Threshold和Times必须同时设置。
	//
	// example:
	//
	// GreaterThanOrEqualToThreshold
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	N                  *int32  `json:"N,omitempty" xml:"N,omitempty"`
	PreCondition       *string `json:"PreCondition,omitempty" xml:"PreCondition,omitempty"`
	// Info级别报警统计方法。
	//
	// N的取值范围：1~50。
	//
	// 该参数的取值由指定云产品的`MetricName`对应的`Statistics`列决定，例如：Maximum、Minimum和Average。关于如何获取该参数的取值，请参见[云产品监控项](https://help.aliyun.com/document_detail/163515.html)。
	//
	// > 报警级别Critical（严重）、Warn（警告）或Info（信息）至少设置一个，且该报警级别中的参数Statistics、ComparisonOperator、Threshold和Times必须同时设置。
	//
	// example:
	//
	// Average
	Statistics *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	// Info级别报警阈值。
	//
	// > 报警级别Critical（严重）、Warn（警告）或Info（信息）至少设置一个，且该报警级别中的参数Statistics、ComparisonOperator、Threshold和Times必须同时设置。
	//
	// example:
	//
	// 90
	Threshold *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	// Info级别报警重试次数。
	//
	// N的取值范围：1~50。
	//
	// > 报警级别Critical（严重）、Warn（警告）或Info（信息）至少设置一个，且该报警级别中的参数Statistics、ComparisonOperator、Threshold和Times必须同时设置。
	//
	// example:
	//
	// 3
	Times *int32 `json:"Times,omitempty" xml:"Times,omitempty"`
}

func (s PutResourceMetricRulesRequestRulesEscalationsInfo) String() string {
	return dara.Prettify(s)
}

func (s PutResourceMetricRulesRequestRulesEscalationsInfo) GoString() string {
	return s.String()
}

func (s *PutResourceMetricRulesRequestRulesEscalationsInfo) GetComparisonOperator() *string {
	return s.ComparisonOperator
}

func (s *PutResourceMetricRulesRequestRulesEscalationsInfo) GetN() *int32 {
	return s.N
}

func (s *PutResourceMetricRulesRequestRulesEscalationsInfo) GetPreCondition() *string {
	return s.PreCondition
}

func (s *PutResourceMetricRulesRequestRulesEscalationsInfo) GetStatistics() *string {
	return s.Statistics
}

func (s *PutResourceMetricRulesRequestRulesEscalationsInfo) GetThreshold() *string {
	return s.Threshold
}

func (s *PutResourceMetricRulesRequestRulesEscalationsInfo) GetTimes() *int32 {
	return s.Times
}

func (s *PutResourceMetricRulesRequestRulesEscalationsInfo) SetComparisonOperator(v string) *PutResourceMetricRulesRequestRulesEscalationsInfo {
	s.ComparisonOperator = &v
	return s
}

func (s *PutResourceMetricRulesRequestRulesEscalationsInfo) SetN(v int32) *PutResourceMetricRulesRequestRulesEscalationsInfo {
	s.N = &v
	return s
}

func (s *PutResourceMetricRulesRequestRulesEscalationsInfo) SetPreCondition(v string) *PutResourceMetricRulesRequestRulesEscalationsInfo {
	s.PreCondition = &v
	return s
}

func (s *PutResourceMetricRulesRequestRulesEscalationsInfo) SetStatistics(v string) *PutResourceMetricRulesRequestRulesEscalationsInfo {
	s.Statistics = &v
	return s
}

func (s *PutResourceMetricRulesRequestRulesEscalationsInfo) SetThreshold(v string) *PutResourceMetricRulesRequestRulesEscalationsInfo {
	s.Threshold = &v
	return s
}

func (s *PutResourceMetricRulesRequestRulesEscalationsInfo) SetTimes(v int32) *PutResourceMetricRulesRequestRulesEscalationsInfo {
	s.Times = &v
	return s
}

func (s *PutResourceMetricRulesRequestRulesEscalationsInfo) Validate() error {
	return dara.Validate(s)
}

type PutResourceMetricRulesRequestRulesEscalationsWarn struct {
	// Warn级别阈值比较符。取值：
	//
	// - GreaterThanOrEqualToThreshold：大于等于。
	//
	// - GreaterThanThreshold：大于。
	//
	// - LessThanOrEqualToThreshold：小于等于。
	//
	// - LessThanThreshold：小于。
	//
	// - NotEqualToThreshold：不等于。
	//
	// - EqualToThreshold：等于。
	//
	// - GreaterThanYesterday：同比昨天时间上涨。
	//
	// - LessThanYesterday：同比昨天时间下降。
	//
	// - GreaterThanLastWeek：同比上周同一时间上涨。
	//
	// - LessThanLastWeek：同比上周同一时间下降。
	//
	// - GreaterThanLastPeriod：环比上周期上涨。
	//
	// - LessThanLastPeriod：环比上周期下降。
	//
	// N的取值范围：1~50。
	//
	// > 报警级别Critical（严重）、Warn（警告）或Info（信息）至少设置一个，且该报警级别中的参数Statistics、ComparisonOperator、Threshold和Times必须同时设置。
	//
	// example:
	//
	// GreaterThanOrEqualToThreshold
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	N                  *int32  `json:"N,omitempty" xml:"N,omitempty"`
	PreCondition       *string `json:"PreCondition,omitempty" xml:"PreCondition,omitempty"`
	// Warn级别报警统计方法。
	//
	// N的取值范围：1~50。
	//
	// 该参数的取值由指定云产品的`MetricName`对应的`Statistics`列决定，例如：Maximum、Minimum和Average。关于如何获取该参数的取值，请参见[云产品监控项](https://help.aliyun.com/document_detail/163515.html)。
	//
	// > 报警级别Critical（严重）、Warn（警告）或Info（信息）至少设置一个，且该报警级别中的参数Statistics、ComparisonOperator、Threshold和Times必须同时设置。
	//
	// example:
	//
	// Average
	Statistics *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	// Warn级别报警阈值。
	//
	// N的取值范围：1~50。
	//
	// > 报警级别Critical（严重）、Warn（警告）或Info（信息）至少设置一个，且该报警级别中的参数Statistics、ComparisonOperator、Threshold和Times必须同时设置。
	//
	// example:
	//
	// 90
	Threshold *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	// Warn级别报警重试次数。
	//
	// N的取值范围：1~50。
	//
	// > 报警级别Critical（严重）、Warn（警告）或Info（信息）至少设置一个，且该报警级别中的参数Statistics、ComparisonOperator、Threshold和Times必须同时设置。
	//
	// example:
	//
	// 3
	Times *int32 `json:"Times,omitempty" xml:"Times,omitempty"`
}

func (s PutResourceMetricRulesRequestRulesEscalationsWarn) String() string {
	return dara.Prettify(s)
}

func (s PutResourceMetricRulesRequestRulesEscalationsWarn) GoString() string {
	return s.String()
}

func (s *PutResourceMetricRulesRequestRulesEscalationsWarn) GetComparisonOperator() *string {
	return s.ComparisonOperator
}

func (s *PutResourceMetricRulesRequestRulesEscalationsWarn) GetN() *int32 {
	return s.N
}

func (s *PutResourceMetricRulesRequestRulesEscalationsWarn) GetPreCondition() *string {
	return s.PreCondition
}

func (s *PutResourceMetricRulesRequestRulesEscalationsWarn) GetStatistics() *string {
	return s.Statistics
}

func (s *PutResourceMetricRulesRequestRulesEscalationsWarn) GetThreshold() *string {
	return s.Threshold
}

func (s *PutResourceMetricRulesRequestRulesEscalationsWarn) GetTimes() *int32 {
	return s.Times
}

func (s *PutResourceMetricRulesRequestRulesEscalationsWarn) SetComparisonOperator(v string) *PutResourceMetricRulesRequestRulesEscalationsWarn {
	s.ComparisonOperator = &v
	return s
}

func (s *PutResourceMetricRulesRequestRulesEscalationsWarn) SetN(v int32) *PutResourceMetricRulesRequestRulesEscalationsWarn {
	s.N = &v
	return s
}

func (s *PutResourceMetricRulesRequestRulesEscalationsWarn) SetPreCondition(v string) *PutResourceMetricRulesRequestRulesEscalationsWarn {
	s.PreCondition = &v
	return s
}

func (s *PutResourceMetricRulesRequestRulesEscalationsWarn) SetStatistics(v string) *PutResourceMetricRulesRequestRulesEscalationsWarn {
	s.Statistics = &v
	return s
}

func (s *PutResourceMetricRulesRequestRulesEscalationsWarn) SetThreshold(v string) *PutResourceMetricRulesRequestRulesEscalationsWarn {
	s.Threshold = &v
	return s
}

func (s *PutResourceMetricRulesRequestRulesEscalationsWarn) SetTimes(v int32) *PutResourceMetricRulesRequestRulesEscalationsWarn {
	s.Times = &v
	return s
}

func (s *PutResourceMetricRulesRequestRulesEscalationsWarn) Validate() error {
	return dara.Validate(s)
}

type PutResourceMetricRulesRequestRulesLabels struct {
	// 标签键。
	//
	// example:
	//
	// tagKey1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// 标签值。
	//
	// > 标签值支持模板参数，将模板参数替换为实际标签值。
	//
	// example:
	//
	// ECS
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s PutResourceMetricRulesRequestRulesLabels) String() string {
	return dara.Prettify(s)
}

func (s PutResourceMetricRulesRequestRulesLabels) GoString() string {
	return s.String()
}

func (s *PutResourceMetricRulesRequestRulesLabels) GetKey() *string {
	return s.Key
}

func (s *PutResourceMetricRulesRequestRulesLabels) GetValue() *string {
	return s.Value
}

func (s *PutResourceMetricRulesRequestRulesLabels) SetKey(v string) *PutResourceMetricRulesRequestRulesLabels {
	s.Key = &v
	return s
}

func (s *PutResourceMetricRulesRequestRulesLabels) SetValue(v string) *PutResourceMetricRulesRequestRulesLabels {
	s.Value = &v
	return s
}

func (s *PutResourceMetricRulesRequestRulesLabels) Validate() error {
	return dara.Validate(s)
}
