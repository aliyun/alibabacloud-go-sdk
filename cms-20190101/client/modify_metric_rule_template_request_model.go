// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyMetricRuleTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlertTemplates(v []*ModifyMetricRuleTemplateRequestAlertTemplates) *ModifyMetricRuleTemplateRequest
	GetAlertTemplates() []*ModifyMetricRuleTemplateRequestAlertTemplates
	SetDescription(v string) *ModifyMetricRuleTemplateRequest
	GetDescription() *string
	SetName(v string) *ModifyMetricRuleTemplateRequest
	GetName() *string
	SetRegionId(v string) *ModifyMetricRuleTemplateRequest
	GetRegionId() *string
	SetRestVersion(v int64) *ModifyMetricRuleTemplateRequest
	GetRestVersion() *int64
	SetTemplateId(v int64) *ModifyMetricRuleTemplateRequest
	GetTemplateId() *int64
}

type ModifyMetricRuleTemplateRequest struct {
	// The details of the alert template.
	AlertTemplates []*ModifyMetricRuleTemplateRequestAlertTemplates `json:"AlertTemplates,omitempty" xml:"AlertTemplates,omitempty" type:"Repeated"`
	// The description of the alert template.
	//
	// example:
	//
	// ECS_template1
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the alert template.
	//
	// For information about how to obtain the name of an alert template, see [DescribeMetricRuleTemplateList](https://help.aliyun.com/document_detail/114982.html).
	//
	// example:
	//
	// test123
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The version of the alert template. The version changes with the number of times that the alert template is modified.
	//
	// For information about how to obtain the version of an alert template, see [DescribeMetricRuleTemplateList](https://help.aliyun.com/document_detail/114982.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	RestVersion *int64 `json:"RestVersion,omitempty" xml:"RestVersion,omitempty"`
	// The ID of the alert template.
	//
	// For information about how to obtain the ID of an alert template, see [DescribeMetricRuleTemplateList](https://help.aliyun.com/document_detail/114982.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s ModifyMetricRuleTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyMetricRuleTemplateRequest) GoString() string {
	return s.String()
}

func (s *ModifyMetricRuleTemplateRequest) GetAlertTemplates() []*ModifyMetricRuleTemplateRequestAlertTemplates {
	return s.AlertTemplates
}

func (s *ModifyMetricRuleTemplateRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyMetricRuleTemplateRequest) GetName() *string {
	return s.Name
}

func (s *ModifyMetricRuleTemplateRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyMetricRuleTemplateRequest) GetRestVersion() *int64 {
	return s.RestVersion
}

func (s *ModifyMetricRuleTemplateRequest) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *ModifyMetricRuleTemplateRequest) SetAlertTemplates(v []*ModifyMetricRuleTemplateRequestAlertTemplates) *ModifyMetricRuleTemplateRequest {
	s.AlertTemplates = v
	return s
}

func (s *ModifyMetricRuleTemplateRequest) SetDescription(v string) *ModifyMetricRuleTemplateRequest {
	s.Description = &v
	return s
}

func (s *ModifyMetricRuleTemplateRequest) SetName(v string) *ModifyMetricRuleTemplateRequest {
	s.Name = &v
	return s
}

func (s *ModifyMetricRuleTemplateRequest) SetRegionId(v string) *ModifyMetricRuleTemplateRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyMetricRuleTemplateRequest) SetRestVersion(v int64) *ModifyMetricRuleTemplateRequest {
	s.RestVersion = &v
	return s
}

func (s *ModifyMetricRuleTemplateRequest) SetTemplateId(v int64) *ModifyMetricRuleTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *ModifyMetricRuleTemplateRequest) Validate() error {
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

type ModifyMetricRuleTemplateRequestAlertTemplates struct {
	Escalations *ModifyMetricRuleTemplateRequestAlertTemplatesEscalations `json:"Escalations,omitempty" xml:"Escalations,omitempty" type:"Struct"`
	// The abbreviation of the cloud service name.
	//
	// Valid values of N: 1 to 200.
	//
	// For more information about how to obtain the abbreviation of a cloud service name, see `metricCategory` in the response parameter `Labels` of the [DescribeProjectMeta](https://help.aliyun.com/document_detail/114916.html) operation.
	//
	// example:
	//
	// ecs
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The metric name.
	//
	// Valid values of N: 1 to 200.
	//
	// For information about how to obtain metrics, see [DescribeMetricMetaList](https://help.aliyun.com/document_detail/98846.html) or [Appendix 1: Metrics](https://help.aliyun.com/document_detail/163515.html).
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
	// example:
	//
	// acs_ecs_dashboard
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The statistical period of the monitoring data.
	//
	// Valid values of N: 1 to 200.
	//
	// > If the value is set to 300 seconds, the monitoring data is collected every 300 seconds. If the monitoring data is reported every 1 minute, the alert system calculates the average, maximum, and minimum values of the monitoring data of 5 minutes and checks whether the aggregated values exceed the threshold. To prevent unexpected alerts, we recommend that you set this parameter together with other parameters.
	//
	// example:
	//
	// 60
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The name of the alert rule.
	//
	// Valid values of N: 1 to 200.
	//
	// example:
	//
	// rule1
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The dimension of the alert. It is an extended field.
	//
	// Valid values of N: 1 to 200.
	//
	// For example, an alert template is applied to an application group, this parameter is set to `{"disk":"/"}`, and the MetricName parameter is set to `DiskUtilization`. In this case, the generated alert rule is applied to the root disk partition (`"/"`) of all instances in the application group to which the alert template is applied.
	//
	// > For more information about the values of extended fields, see [DescribeMetricRuleTemplateAttribute](https://help.aliyun.com/document_detail/114979.html).
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
	// https://apiwebhook.hipac.cn/api/v1/alarm/aly/eregfeeferrtbnmkdszp
	Webhook *string `json:"Webhook,omitempty" xml:"Webhook,omitempty"`
}

func (s ModifyMetricRuleTemplateRequestAlertTemplates) String() string {
	return dara.Prettify(s)
}

func (s ModifyMetricRuleTemplateRequestAlertTemplates) GoString() string {
	return s.String()
}

func (s *ModifyMetricRuleTemplateRequestAlertTemplates) GetEscalations() *ModifyMetricRuleTemplateRequestAlertTemplatesEscalations {
	return s.Escalations
}

func (s *ModifyMetricRuleTemplateRequestAlertTemplates) GetCategory() *string {
	return s.Category
}

func (s *ModifyMetricRuleTemplateRequestAlertTemplates) GetMetricName() *string {
	return s.MetricName
}

func (s *ModifyMetricRuleTemplateRequestAlertTemplates) GetNamespace() *string {
	return s.Namespace
}

func (s *ModifyMetricRuleTemplateRequestAlertTemplates) GetPeriod() *int32 {
	return s.Period
}

func (s *ModifyMetricRuleTemplateRequestAlertTemplates) GetRuleName() *string {
	return s.RuleName
}

func (s *ModifyMetricRuleTemplateRequestAlertTemplates) GetSelector() *string {
	return s.Selector
}

func (s *ModifyMetricRuleTemplateRequestAlertTemplates) GetWebhook() *string {
	return s.Webhook
}

func (s *ModifyMetricRuleTemplateRequestAlertTemplates) SetEscalations(v *ModifyMetricRuleTemplateRequestAlertTemplatesEscalations) *ModifyMetricRuleTemplateRequestAlertTemplates {
	s.Escalations = v
	return s
}

func (s *ModifyMetricRuleTemplateRequestAlertTemplates) SetCategory(v string) *ModifyMetricRuleTemplateRequestAlertTemplates {
	s.Category = &v
	return s
}

func (s *ModifyMetricRuleTemplateRequestAlertTemplates) SetMetricName(v string) *ModifyMetricRuleTemplateRequestAlertTemplates {
	s.MetricName = &v
	return s
}

func (s *ModifyMetricRuleTemplateRequestAlertTemplates) SetNamespace(v string) *ModifyMetricRuleTemplateRequestAlertTemplates {
	s.Namespace = &v
	return s
}

func (s *ModifyMetricRuleTemplateRequestAlertTemplates) SetPeriod(v int32) *ModifyMetricRuleTemplateRequestAlertTemplates {
	s.Period = &v
	return s
}

func (s *ModifyMetricRuleTemplateRequestAlertTemplates) SetRuleName(v string) *ModifyMetricRuleTemplateRequestAlertTemplates {
	s.RuleName = &v
	return s
}

func (s *ModifyMetricRuleTemplateRequestAlertTemplates) SetSelector(v string) *ModifyMetricRuleTemplateRequestAlertTemplates {
	s.Selector = &v
	return s
}

func (s *ModifyMetricRuleTemplateRequestAlertTemplates) SetWebhook(v string) *ModifyMetricRuleTemplateRequestAlertTemplates {
	s.Webhook = &v
	return s
}

func (s *ModifyMetricRuleTemplateRequestAlertTemplates) Validate() error {
	if s.Escalations != nil {
		if err := s.Escalations.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyMetricRuleTemplateRequestAlertTemplatesEscalations struct {
	Critical *ModifyMetricRuleTemplateRequestAlertTemplatesEscalationsCritical `json:"Critical,omitempty" xml:"Critical,omitempty" type:"Struct"`
	Info     *ModifyMetricRuleTemplateRequestAlertTemplatesEscalationsInfo     `json:"Info,omitempty" xml:"Info,omitempty" type:"Struct"`
	Warn     *ModifyMetricRuleTemplateRequestAlertTemplatesEscalationsWarn     `json:"Warn,omitempty" xml:"Warn,omitempty" type:"Struct"`
}

func (s ModifyMetricRuleTemplateRequestAlertTemplatesEscalations) String() string {
	return dara.Prettify(s)
}

func (s ModifyMetricRuleTemplateRequestAlertTemplatesEscalations) GoString() string {
	return s.String()
}

func (s *ModifyMetricRuleTemplateRequestAlertTemplatesEscalations) GetCritical() *ModifyMetricRuleTemplateRequestAlertTemplatesEscalationsCritical {
	return s.Critical
}

func (s *ModifyMetricRuleTemplateRequestAlertTemplatesEscalations) GetInfo() *ModifyMetricRuleTemplateRequestAlertTemplatesEscalationsInfo {
	return s.Info
}

func (s *ModifyMetricRuleTemplateRequestAlertTemplatesEscalations) GetWarn() *ModifyMetricRuleTemplateRequestAlertTemplatesEscalationsWarn {
	return s.Warn
}

func (s *ModifyMetricRuleTemplateRequestAlertTemplatesEscalations) SetCritical(v *ModifyMetricRuleTemplateRequestAlertTemplatesEscalationsCritical) *ModifyMetricRuleTemplateRequestAlertTemplatesEscalations {
	s.Critical = v
	return s
}

func (s *ModifyMetricRuleTemplateRequestAlertTemplatesEscalations) SetInfo(v *ModifyMetricRuleTemplateRequestAlertTemplatesEscalationsInfo) *ModifyMetricRuleTemplateRequestAlertTemplatesEscalations {
	s.Info = v
	return s
}

func (s *ModifyMetricRuleTemplateRequestAlertTemplatesEscalations) SetWarn(v *ModifyMetricRuleTemplateRequestAlertTemplatesEscalationsWarn) *ModifyMetricRuleTemplateRequestAlertTemplatesEscalations {
	s.Warn = v
	return s
}

func (s *ModifyMetricRuleTemplateRequestAlertTemplatesEscalations) Validate() error {
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

type ModifyMetricRuleTemplateRequestAlertTemplatesEscalationsCritical struct {
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	N                  *int32  `json:"N,omitempty" xml:"N,omitempty"`
	Statistics         *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	Threshold          *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	Times              *int32  `json:"Times,omitempty" xml:"Times,omitempty"`
}

func (s ModifyMetricRuleTemplateRequestAlertTemplatesEscalationsCritical) String() string {
	return dara.Prettify(s)
}

func (s ModifyMetricRuleTemplateRequestAlertTemplatesEscalationsCritical) GoString() string {
	return s.String()
}

func (s *ModifyMetricRuleTemplateRequestAlertTemplatesEscalationsCritical) GetComparisonOperator() *string {
	return s.ComparisonOperator
}

func (s *ModifyMetricRuleTemplateRequestAlertTemplatesEscalationsCritical) GetN() *int32 {
	return s.N
}

func (s *ModifyMetricRuleTemplateRequestAlertTemplatesEscalationsCritical) GetStatistics() *string {
	return s.Statistics
}

func (s *ModifyMetricRuleTemplateRequestAlertTemplatesEscalationsCritical) GetThreshold() *string {
	return s.Threshold
}

func (s *ModifyMetricRuleTemplateRequestAlertTemplatesEscalationsCritical) GetTimes() *int32 {
	return s.Times
}

func (s *ModifyMetricRuleTemplateRequestAlertTemplatesEscalationsCritical) SetComparisonOperator(v string) *ModifyMetricRuleTemplateRequestAlertTemplatesEscalationsCritical {
	s.ComparisonOperator = &v
	return s
}

func (s *ModifyMetricRuleTemplateRequestAlertTemplatesEscalationsCritical) SetN(v int32) *ModifyMetricRuleTemplateRequestAlertTemplatesEscalationsCritical {
	s.N = &v
	return s
}

func (s *ModifyMetricRuleTemplateRequestAlertTemplatesEscalationsCritical) SetStatistics(v string) *ModifyMetricRuleTemplateRequestAlertTemplatesEscalationsCritical {
	s.Statistics = &v
	return s
}

func (s *ModifyMetricRuleTemplateRequestAlertTemplatesEscalationsCritical) SetThreshold(v string) *ModifyMetricRuleTemplateRequestAlertTemplatesEscalationsCritical {
	s.Threshold = &v
	return s
}

func (s *ModifyMetricRuleTemplateRequestAlertTemplatesEscalationsCritical) SetTimes(v int32) *ModifyMetricRuleTemplateRequestAlertTemplatesEscalationsCritical {
	s.Times = &v
	return s
}

func (s *ModifyMetricRuleTemplateRequestAlertTemplatesEscalationsCritical) Validate() error {
	return dara.Validate(s)
}

type ModifyMetricRuleTemplateRequestAlertTemplatesEscalationsInfo struct {
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	N                  *int32  `json:"N,omitempty" xml:"N,omitempty"`
	Statistics         *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	Threshold          *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	Times              *int32  `json:"Times,omitempty" xml:"Times,omitempty"`
}

func (s ModifyMetricRuleTemplateRequestAlertTemplatesEscalationsInfo) String() string {
	return dara.Prettify(s)
}

func (s ModifyMetricRuleTemplateRequestAlertTemplatesEscalationsInfo) GoString() string {
	return s.String()
}

func (s *ModifyMetricRuleTemplateRequestAlertTemplatesEscalationsInfo) GetComparisonOperator() *string {
	return s.ComparisonOperator
}

func (s *ModifyMetricRuleTemplateRequestAlertTemplatesEscalationsInfo) GetN() *int32 {
	return s.N
}

func (s *ModifyMetricRuleTemplateRequestAlertTemplatesEscalationsInfo) GetStatistics() *string {
	return s.Statistics
}

func (s *ModifyMetricRuleTemplateRequestAlertTemplatesEscalationsInfo) GetThreshold() *string {
	return s.Threshold
}

func (s *ModifyMetricRuleTemplateRequestAlertTemplatesEscalationsInfo) GetTimes() *int32 {
	return s.Times
}

func (s *ModifyMetricRuleTemplateRequestAlertTemplatesEscalationsInfo) SetComparisonOperator(v string) *ModifyMetricRuleTemplateRequestAlertTemplatesEscalationsInfo {
	s.ComparisonOperator = &v
	return s
}

func (s *ModifyMetricRuleTemplateRequestAlertTemplatesEscalationsInfo) SetN(v int32) *ModifyMetricRuleTemplateRequestAlertTemplatesEscalationsInfo {
	s.N = &v
	return s
}

func (s *ModifyMetricRuleTemplateRequestAlertTemplatesEscalationsInfo) SetStatistics(v string) *ModifyMetricRuleTemplateRequestAlertTemplatesEscalationsInfo {
	s.Statistics = &v
	return s
}

func (s *ModifyMetricRuleTemplateRequestAlertTemplatesEscalationsInfo) SetThreshold(v string) *ModifyMetricRuleTemplateRequestAlertTemplatesEscalationsInfo {
	s.Threshold = &v
	return s
}

func (s *ModifyMetricRuleTemplateRequestAlertTemplatesEscalationsInfo) SetTimes(v int32) *ModifyMetricRuleTemplateRequestAlertTemplatesEscalationsInfo {
	s.Times = &v
	return s
}

func (s *ModifyMetricRuleTemplateRequestAlertTemplatesEscalationsInfo) Validate() error {
	return dara.Validate(s)
}

type ModifyMetricRuleTemplateRequestAlertTemplatesEscalationsWarn struct {
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	N                  *int32  `json:"N,omitempty" xml:"N,omitempty"`
	Statistics         *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	Threshold          *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	Times              *int32  `json:"Times,omitempty" xml:"Times,omitempty"`
}

func (s ModifyMetricRuleTemplateRequestAlertTemplatesEscalationsWarn) String() string {
	return dara.Prettify(s)
}

func (s ModifyMetricRuleTemplateRequestAlertTemplatesEscalationsWarn) GoString() string {
	return s.String()
}

func (s *ModifyMetricRuleTemplateRequestAlertTemplatesEscalationsWarn) GetComparisonOperator() *string {
	return s.ComparisonOperator
}

func (s *ModifyMetricRuleTemplateRequestAlertTemplatesEscalationsWarn) GetN() *int32 {
	return s.N
}

func (s *ModifyMetricRuleTemplateRequestAlertTemplatesEscalationsWarn) GetStatistics() *string {
	return s.Statistics
}

func (s *ModifyMetricRuleTemplateRequestAlertTemplatesEscalationsWarn) GetThreshold() *string {
	return s.Threshold
}

func (s *ModifyMetricRuleTemplateRequestAlertTemplatesEscalationsWarn) GetTimes() *int32 {
	return s.Times
}

func (s *ModifyMetricRuleTemplateRequestAlertTemplatesEscalationsWarn) SetComparisonOperator(v string) *ModifyMetricRuleTemplateRequestAlertTemplatesEscalationsWarn {
	s.ComparisonOperator = &v
	return s
}

func (s *ModifyMetricRuleTemplateRequestAlertTemplatesEscalationsWarn) SetN(v int32) *ModifyMetricRuleTemplateRequestAlertTemplatesEscalationsWarn {
	s.N = &v
	return s
}

func (s *ModifyMetricRuleTemplateRequestAlertTemplatesEscalationsWarn) SetStatistics(v string) *ModifyMetricRuleTemplateRequestAlertTemplatesEscalationsWarn {
	s.Statistics = &v
	return s
}

func (s *ModifyMetricRuleTemplateRequestAlertTemplatesEscalationsWarn) SetThreshold(v string) *ModifyMetricRuleTemplateRequestAlertTemplatesEscalationsWarn {
	s.Threshold = &v
	return s
}

func (s *ModifyMetricRuleTemplateRequestAlertTemplatesEscalationsWarn) SetTimes(v int32) *ModifyMetricRuleTemplateRequestAlertTemplatesEscalationsWarn {
	s.Times = &v
	return s
}

func (s *ModifyMetricRuleTemplateRequestAlertTemplatesEscalationsWarn) Validate() error {
	return dara.Validate(s)
}
