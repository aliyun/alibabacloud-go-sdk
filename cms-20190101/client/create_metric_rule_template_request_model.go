// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMetricRuleTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlertTemplates(v []*CreateMetricRuleTemplateRequestAlertTemplates) *CreateMetricRuleTemplateRequest
	GetAlertTemplates() []*CreateMetricRuleTemplateRequestAlertTemplates
	SetDescription(v string) *CreateMetricRuleTemplateRequest
	GetDescription() *string
	SetName(v string) *CreateMetricRuleTemplateRequest
	GetName() *string
	SetRegionId(v string) *CreateMetricRuleTemplateRequest
	GetRegionId() *string
}

type CreateMetricRuleTemplateRequest struct {
	// The details of the alert template.
	AlertTemplates []*CreateMetricRuleTemplateRequestAlertTemplates `json:"AlertTemplates,omitempty" xml:"AlertTemplates,omitempty" type:"Repeated"`
	// The description of the alert template.
	//
	// example:
	//
	// ECS_Template1
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the alert template.
	//
	// This parameter is required.
	//
	// example:
	//
	// Template1
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateMetricRuleTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMetricRuleTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreateMetricRuleTemplateRequest) GetAlertTemplates() []*CreateMetricRuleTemplateRequestAlertTemplates {
	return s.AlertTemplates
}

func (s *CreateMetricRuleTemplateRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateMetricRuleTemplateRequest) GetName() *string {
	return s.Name
}

func (s *CreateMetricRuleTemplateRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateMetricRuleTemplateRequest) SetAlertTemplates(v []*CreateMetricRuleTemplateRequestAlertTemplates) *CreateMetricRuleTemplateRequest {
	s.AlertTemplates = v
	return s
}

func (s *CreateMetricRuleTemplateRequest) SetDescription(v string) *CreateMetricRuleTemplateRequest {
	s.Description = &v
	return s
}

func (s *CreateMetricRuleTemplateRequest) SetName(v string) *CreateMetricRuleTemplateRequest {
	s.Name = &v
	return s
}

func (s *CreateMetricRuleTemplateRequest) SetRegionId(v string) *CreateMetricRuleTemplateRequest {
	s.RegionId = &v
	return s
}

func (s *CreateMetricRuleTemplateRequest) Validate() error {
	if s.AlertTemplates != nil {
		for _, item := range s.AlertTemplates {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateMetricRuleTemplateRequestAlertTemplates struct {
	Escalations *CreateMetricRuleTemplateRequestAlertTemplatesEscalations `json:"Escalations,omitempty" xml:"Escalations,omitempty" type:"Struct"`
	// The abbreviation of the cloud service name.
	//
	// Valid values of N: 1 to 200.
	//
	// For more information about how to obtain the abbreviation of a cloud service name, see `metricCategory` in the response parameter `Labels` of the [DescribeProjectMeta](https://help.aliyun.com/document_detail/114916.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecs
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The metric name.
	//
	// Valid values of N: 1 to 200.
	//
	// For information about how to obtain the name of a metric, see [DescribeMetricMetaList](https://help.aliyun.com/document_detail/98846.html) or [Metrics](https://help.aliyun.com/document_detail/163515.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cpu_total
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// The namespace of the cloud service.
	//
	// Valid values of N: 1 to 200.
	//
	// For information about how to obtain the namespace of a cloud service, see [DescribeMetricMetaList](https://help.aliyun.com/document_detail/98846.html) or [Appendix 1: Metrics](https://help.aliyun.com/document_detail/163515.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// acs_ecs_dashboard
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The aggregation period of monitoring data. Unit: seconds.
	//
	// The default value is the minimum aggregation period. Generally, you do not need to specify the minimum aggregation period.
	//
	// Valid values of N: 1 to 200.
	//
	// example:
	//
	// 60
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The name of the alert rule.
	//
	// Valid values of N: 1 to 200.
	//
	// This parameter is required.
	//
	// example:
	//
	// ECS_Rule1
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The dimension of the alert. It is an extended field.
	//
	// Valid values of N: 1 to 200.
	//
	// For example, an alert template is applied to an application group, this parameter is set to `{"disk":"/"}`, and the MetricName parameter is set to `DiskUtilization`. In this case, the generated alert rule is applied to the root disk partition (`"/"`) of all instances in the application group to which the alert template is applied.
	//
	// >  For more information about the values of extended fields, see [DescribeMetricRuleTemplateAttribute](https://help.aliyun.com/document_detail/114979.html).
	//
	// example:
	//
	// {"disk":"/"}
	Selector *string `json:"Selector,omitempty" xml:"Selector,omitempty"`
	// The callback URL.
	//
	// Valid values of N: 1 to 200.
	//
	// The callback URL must be accessible over the Internet. CloudMonitor pushes an alert notification to the specified callback URL by sending an HTTP POST request. Only the HTTP protocol is supported.
	//
	// example:
	//
	// http://ww.aliyun.com
	Webhook *string `json:"Webhook,omitempty" xml:"Webhook,omitempty"`
}

func (s CreateMetricRuleTemplateRequestAlertTemplates) String() string {
	return dara.Prettify(s)
}

func (s CreateMetricRuleTemplateRequestAlertTemplates) GoString() string {
	return s.String()
}

func (s *CreateMetricRuleTemplateRequestAlertTemplates) GetEscalations() *CreateMetricRuleTemplateRequestAlertTemplatesEscalations {
	return s.Escalations
}

func (s *CreateMetricRuleTemplateRequestAlertTemplates) GetCategory() *string {
	return s.Category
}

func (s *CreateMetricRuleTemplateRequestAlertTemplates) GetMetricName() *string {
	return s.MetricName
}

func (s *CreateMetricRuleTemplateRequestAlertTemplates) GetNamespace() *string {
	return s.Namespace
}

func (s *CreateMetricRuleTemplateRequestAlertTemplates) GetPeriod() *int32 {
	return s.Period
}

func (s *CreateMetricRuleTemplateRequestAlertTemplates) GetRuleName() *string {
	return s.RuleName
}

func (s *CreateMetricRuleTemplateRequestAlertTemplates) GetSelector() *string {
	return s.Selector
}

func (s *CreateMetricRuleTemplateRequestAlertTemplates) GetWebhook() *string {
	return s.Webhook
}

func (s *CreateMetricRuleTemplateRequestAlertTemplates) SetEscalations(v *CreateMetricRuleTemplateRequestAlertTemplatesEscalations) *CreateMetricRuleTemplateRequestAlertTemplates {
	s.Escalations = v
	return s
}

func (s *CreateMetricRuleTemplateRequestAlertTemplates) SetCategory(v string) *CreateMetricRuleTemplateRequestAlertTemplates {
	s.Category = &v
	return s
}

func (s *CreateMetricRuleTemplateRequestAlertTemplates) SetMetricName(v string) *CreateMetricRuleTemplateRequestAlertTemplates {
	s.MetricName = &v
	return s
}

func (s *CreateMetricRuleTemplateRequestAlertTemplates) SetNamespace(v string) *CreateMetricRuleTemplateRequestAlertTemplates {
	s.Namespace = &v
	return s
}

func (s *CreateMetricRuleTemplateRequestAlertTemplates) SetPeriod(v int32) *CreateMetricRuleTemplateRequestAlertTemplates {
	s.Period = &v
	return s
}

func (s *CreateMetricRuleTemplateRequestAlertTemplates) SetRuleName(v string) *CreateMetricRuleTemplateRequestAlertTemplates {
	s.RuleName = &v
	return s
}

func (s *CreateMetricRuleTemplateRequestAlertTemplates) SetSelector(v string) *CreateMetricRuleTemplateRequestAlertTemplates {
	s.Selector = &v
	return s
}

func (s *CreateMetricRuleTemplateRequestAlertTemplates) SetWebhook(v string) *CreateMetricRuleTemplateRequestAlertTemplates {
	s.Webhook = &v
	return s
}

func (s *CreateMetricRuleTemplateRequestAlertTemplates) Validate() error {
	if s.Escalations != nil {
		if err := s.Escalations.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateMetricRuleTemplateRequestAlertTemplatesEscalations struct {
	Critical *CreateMetricRuleTemplateRequestAlertTemplatesEscalationsCritical `json:"Critical,omitempty" xml:"Critical,omitempty" type:"Struct"`
	Info     *CreateMetricRuleTemplateRequestAlertTemplatesEscalationsInfo     `json:"Info,omitempty" xml:"Info,omitempty" type:"Struct"`
	Warn     *CreateMetricRuleTemplateRequestAlertTemplatesEscalationsWarn     `json:"Warn,omitempty" xml:"Warn,omitempty" type:"Struct"`
}

func (s CreateMetricRuleTemplateRequestAlertTemplatesEscalations) String() string {
	return dara.Prettify(s)
}

func (s CreateMetricRuleTemplateRequestAlertTemplatesEscalations) GoString() string {
	return s.String()
}

func (s *CreateMetricRuleTemplateRequestAlertTemplatesEscalations) GetCritical() *CreateMetricRuleTemplateRequestAlertTemplatesEscalationsCritical {
	return s.Critical
}

func (s *CreateMetricRuleTemplateRequestAlertTemplatesEscalations) GetInfo() *CreateMetricRuleTemplateRequestAlertTemplatesEscalationsInfo {
	return s.Info
}

func (s *CreateMetricRuleTemplateRequestAlertTemplatesEscalations) GetWarn() *CreateMetricRuleTemplateRequestAlertTemplatesEscalationsWarn {
	return s.Warn
}

func (s *CreateMetricRuleTemplateRequestAlertTemplatesEscalations) SetCritical(v *CreateMetricRuleTemplateRequestAlertTemplatesEscalationsCritical) *CreateMetricRuleTemplateRequestAlertTemplatesEscalations {
	s.Critical = v
	return s
}

func (s *CreateMetricRuleTemplateRequestAlertTemplatesEscalations) SetInfo(v *CreateMetricRuleTemplateRequestAlertTemplatesEscalationsInfo) *CreateMetricRuleTemplateRequestAlertTemplatesEscalations {
	s.Info = v
	return s
}

func (s *CreateMetricRuleTemplateRequestAlertTemplatesEscalations) SetWarn(v *CreateMetricRuleTemplateRequestAlertTemplatesEscalationsWarn) *CreateMetricRuleTemplateRequestAlertTemplatesEscalations {
	s.Warn = v
	return s
}

func (s *CreateMetricRuleTemplateRequestAlertTemplatesEscalations) Validate() error {
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

type CreateMetricRuleTemplateRequestAlertTemplatesEscalationsCritical struct {
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	N                  *int32  `json:"N,omitempty" xml:"N,omitempty"`
	Statistics         *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	Threshold          *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	Times              *int32  `json:"Times,omitempty" xml:"Times,omitempty"`
}

func (s CreateMetricRuleTemplateRequestAlertTemplatesEscalationsCritical) String() string {
	return dara.Prettify(s)
}

func (s CreateMetricRuleTemplateRequestAlertTemplatesEscalationsCritical) GoString() string {
	return s.String()
}

func (s *CreateMetricRuleTemplateRequestAlertTemplatesEscalationsCritical) GetComparisonOperator() *string {
	return s.ComparisonOperator
}

func (s *CreateMetricRuleTemplateRequestAlertTemplatesEscalationsCritical) GetN() *int32 {
	return s.N
}

func (s *CreateMetricRuleTemplateRequestAlertTemplatesEscalationsCritical) GetStatistics() *string {
	return s.Statistics
}

func (s *CreateMetricRuleTemplateRequestAlertTemplatesEscalationsCritical) GetThreshold() *string {
	return s.Threshold
}

func (s *CreateMetricRuleTemplateRequestAlertTemplatesEscalationsCritical) GetTimes() *int32 {
	return s.Times
}

func (s *CreateMetricRuleTemplateRequestAlertTemplatesEscalationsCritical) SetComparisonOperator(v string) *CreateMetricRuleTemplateRequestAlertTemplatesEscalationsCritical {
	s.ComparisonOperator = &v
	return s
}

func (s *CreateMetricRuleTemplateRequestAlertTemplatesEscalationsCritical) SetN(v int32) *CreateMetricRuleTemplateRequestAlertTemplatesEscalationsCritical {
	s.N = &v
	return s
}

func (s *CreateMetricRuleTemplateRequestAlertTemplatesEscalationsCritical) SetStatistics(v string) *CreateMetricRuleTemplateRequestAlertTemplatesEscalationsCritical {
	s.Statistics = &v
	return s
}

func (s *CreateMetricRuleTemplateRequestAlertTemplatesEscalationsCritical) SetThreshold(v string) *CreateMetricRuleTemplateRequestAlertTemplatesEscalationsCritical {
	s.Threshold = &v
	return s
}

func (s *CreateMetricRuleTemplateRequestAlertTemplatesEscalationsCritical) SetTimes(v int32) *CreateMetricRuleTemplateRequestAlertTemplatesEscalationsCritical {
	s.Times = &v
	return s
}

func (s *CreateMetricRuleTemplateRequestAlertTemplatesEscalationsCritical) Validate() error {
	return dara.Validate(s)
}

type CreateMetricRuleTemplateRequestAlertTemplatesEscalationsInfo struct {
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	N                  *int32  `json:"N,omitempty" xml:"N,omitempty"`
	Statistics         *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	Threshold          *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	Times              *int32  `json:"Times,omitempty" xml:"Times,omitempty"`
}

func (s CreateMetricRuleTemplateRequestAlertTemplatesEscalationsInfo) String() string {
	return dara.Prettify(s)
}

func (s CreateMetricRuleTemplateRequestAlertTemplatesEscalationsInfo) GoString() string {
	return s.String()
}

func (s *CreateMetricRuleTemplateRequestAlertTemplatesEscalationsInfo) GetComparisonOperator() *string {
	return s.ComparisonOperator
}

func (s *CreateMetricRuleTemplateRequestAlertTemplatesEscalationsInfo) GetN() *int32 {
	return s.N
}

func (s *CreateMetricRuleTemplateRequestAlertTemplatesEscalationsInfo) GetStatistics() *string {
	return s.Statistics
}

func (s *CreateMetricRuleTemplateRequestAlertTemplatesEscalationsInfo) GetThreshold() *string {
	return s.Threshold
}

func (s *CreateMetricRuleTemplateRequestAlertTemplatesEscalationsInfo) GetTimes() *int32 {
	return s.Times
}

func (s *CreateMetricRuleTemplateRequestAlertTemplatesEscalationsInfo) SetComparisonOperator(v string) *CreateMetricRuleTemplateRequestAlertTemplatesEscalationsInfo {
	s.ComparisonOperator = &v
	return s
}

func (s *CreateMetricRuleTemplateRequestAlertTemplatesEscalationsInfo) SetN(v int32) *CreateMetricRuleTemplateRequestAlertTemplatesEscalationsInfo {
	s.N = &v
	return s
}

func (s *CreateMetricRuleTemplateRequestAlertTemplatesEscalationsInfo) SetStatistics(v string) *CreateMetricRuleTemplateRequestAlertTemplatesEscalationsInfo {
	s.Statistics = &v
	return s
}

func (s *CreateMetricRuleTemplateRequestAlertTemplatesEscalationsInfo) SetThreshold(v string) *CreateMetricRuleTemplateRequestAlertTemplatesEscalationsInfo {
	s.Threshold = &v
	return s
}

func (s *CreateMetricRuleTemplateRequestAlertTemplatesEscalationsInfo) SetTimes(v int32) *CreateMetricRuleTemplateRequestAlertTemplatesEscalationsInfo {
	s.Times = &v
	return s
}

func (s *CreateMetricRuleTemplateRequestAlertTemplatesEscalationsInfo) Validate() error {
	return dara.Validate(s)
}

type CreateMetricRuleTemplateRequestAlertTemplatesEscalationsWarn struct {
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	N                  *int32  `json:"N,omitempty" xml:"N,omitempty"`
	Statistics         *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	Threshold          *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	Times              *int32  `json:"Times,omitempty" xml:"Times,omitempty"`
}

func (s CreateMetricRuleTemplateRequestAlertTemplatesEscalationsWarn) String() string {
	return dara.Prettify(s)
}

func (s CreateMetricRuleTemplateRequestAlertTemplatesEscalationsWarn) GoString() string {
	return s.String()
}

func (s *CreateMetricRuleTemplateRequestAlertTemplatesEscalationsWarn) GetComparisonOperator() *string {
	return s.ComparisonOperator
}

func (s *CreateMetricRuleTemplateRequestAlertTemplatesEscalationsWarn) GetN() *int32 {
	return s.N
}

func (s *CreateMetricRuleTemplateRequestAlertTemplatesEscalationsWarn) GetStatistics() *string {
	return s.Statistics
}

func (s *CreateMetricRuleTemplateRequestAlertTemplatesEscalationsWarn) GetThreshold() *string {
	return s.Threshold
}

func (s *CreateMetricRuleTemplateRequestAlertTemplatesEscalationsWarn) GetTimes() *int32 {
	return s.Times
}

func (s *CreateMetricRuleTemplateRequestAlertTemplatesEscalationsWarn) SetComparisonOperator(v string) *CreateMetricRuleTemplateRequestAlertTemplatesEscalationsWarn {
	s.ComparisonOperator = &v
	return s
}

func (s *CreateMetricRuleTemplateRequestAlertTemplatesEscalationsWarn) SetN(v int32) *CreateMetricRuleTemplateRequestAlertTemplatesEscalationsWarn {
	s.N = &v
	return s
}

func (s *CreateMetricRuleTemplateRequestAlertTemplatesEscalationsWarn) SetStatistics(v string) *CreateMetricRuleTemplateRequestAlertTemplatesEscalationsWarn {
	s.Statistics = &v
	return s
}

func (s *CreateMetricRuleTemplateRequestAlertTemplatesEscalationsWarn) SetThreshold(v string) *CreateMetricRuleTemplateRequestAlertTemplatesEscalationsWarn {
	s.Threshold = &v
	return s
}

func (s *CreateMetricRuleTemplateRequestAlertTemplatesEscalationsWarn) SetTimes(v int32) *CreateMetricRuleTemplateRequestAlertTemplatesEscalationsWarn {
	s.Times = &v
	return s
}

func (s *CreateMetricRuleTemplateRequestAlertTemplatesEscalationsWarn) Validate() error {
	return dara.Validate(s)
}
