// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMetricRuleTemplateAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DescribeMetricRuleTemplateAttributeResponseBody
	GetCode() *int32
	SetMessage(v string) *DescribeMetricRuleTemplateAttributeResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeMetricRuleTemplateAttributeResponseBody
	GetRequestId() *string
	SetResource(v *DescribeMetricRuleTemplateAttributeResponseBodyResource) *DescribeMetricRuleTemplateAttributeResponseBody
	GetResource() *DescribeMetricRuleTemplateAttributeResponseBodyResource
	SetSuccess(v bool) *DescribeMetricRuleTemplateAttributeResponseBody
	GetSuccess() *bool
}

type DescribeMetricRuleTemplateAttributeResponseBody struct {
	// The HTTP status code.
	//
	// >  The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// The specified resource is not found.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 8F3A82AD-DA92-52B0-8EC6-C059D1C3839F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The details of the alert template.
	Resource *DescribeMetricRuleTemplateAttributeResponseBodyResource `json:"Resource,omitempty" xml:"Resource,omitempty" type:"Struct"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeMetricRuleTemplateAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetricRuleTemplateAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMetricRuleTemplateAttributeResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DescribeMetricRuleTemplateAttributeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeMetricRuleTemplateAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeMetricRuleTemplateAttributeResponseBody) GetResource() *DescribeMetricRuleTemplateAttributeResponseBodyResource {
	return s.Resource
}

func (s *DescribeMetricRuleTemplateAttributeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeMetricRuleTemplateAttributeResponseBody) SetCode(v int32) *DescribeMetricRuleTemplateAttributeResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeMetricRuleTemplateAttributeResponseBody) SetMessage(v string) *DescribeMetricRuleTemplateAttributeResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeMetricRuleTemplateAttributeResponseBody) SetRequestId(v string) *DescribeMetricRuleTemplateAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMetricRuleTemplateAttributeResponseBody) SetResource(v *DescribeMetricRuleTemplateAttributeResponseBodyResource) *DescribeMetricRuleTemplateAttributeResponseBody {
	s.Resource = v
	return s
}

func (s *DescribeMetricRuleTemplateAttributeResponseBody) SetSuccess(v bool) *DescribeMetricRuleTemplateAttributeResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeMetricRuleTemplateAttributeResponseBody) Validate() error {
	if s.Resource != nil {
		if err := s.Resource.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeMetricRuleTemplateAttributeResponseBodyResource struct {
	// The queried alert templates.
	AlertTemplates *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplates `json:"AlertTemplates,omitempty" xml:"AlertTemplates,omitempty" type:"Struct"`
	// The description of the alert template.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the alert template.
	//
	// example:
	//
	// ECS_Template1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The version of the alert template.
	//
	// example:
	//
	// 1
	RestVersion *string `json:"RestVersion,omitempty" xml:"RestVersion,omitempty"`
	// The ID of the alert template.
	//
	// example:
	//
	// 70****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DescribeMetricRuleTemplateAttributeResponseBodyResource) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetricRuleTemplateAttributeResponseBodyResource) GoString() string {
	return s.String()
}

func (s *DescribeMetricRuleTemplateAttributeResponseBodyResource) GetAlertTemplates() *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplates {
	return s.AlertTemplates
}

func (s *DescribeMetricRuleTemplateAttributeResponseBodyResource) GetDescription() *string {
	return s.Description
}

func (s *DescribeMetricRuleTemplateAttributeResponseBodyResource) GetName() *string {
	return s.Name
}

func (s *DescribeMetricRuleTemplateAttributeResponseBodyResource) GetRestVersion() *string {
	return s.RestVersion
}

func (s *DescribeMetricRuleTemplateAttributeResponseBodyResource) GetTemplateId() *string {
	return s.TemplateId
}

func (s *DescribeMetricRuleTemplateAttributeResponseBodyResource) SetAlertTemplates(v *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplates) *DescribeMetricRuleTemplateAttributeResponseBodyResource {
	s.AlertTemplates = v
	return s
}

func (s *DescribeMetricRuleTemplateAttributeResponseBodyResource) SetDescription(v string) *DescribeMetricRuleTemplateAttributeResponseBodyResource {
	s.Description = &v
	return s
}

func (s *DescribeMetricRuleTemplateAttributeResponseBodyResource) SetName(v string) *DescribeMetricRuleTemplateAttributeResponseBodyResource {
	s.Name = &v
	return s
}

func (s *DescribeMetricRuleTemplateAttributeResponseBodyResource) SetRestVersion(v string) *DescribeMetricRuleTemplateAttributeResponseBodyResource {
	s.RestVersion = &v
	return s
}

func (s *DescribeMetricRuleTemplateAttributeResponseBodyResource) SetTemplateId(v string) *DescribeMetricRuleTemplateAttributeResponseBodyResource {
	s.TemplateId = &v
	return s
}

func (s *DescribeMetricRuleTemplateAttributeResponseBodyResource) Validate() error {
	if s.AlertTemplates != nil {
		if err := s.AlertTemplates.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplates struct {
	AlertTemplate []*DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplate `json:"AlertTemplate,omitempty" xml:"AlertTemplate,omitempty" type:"Repeated"`
}

func (s DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplates) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplates) GoString() string {
	return s.String()
}

func (s *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplates) GetAlertTemplate() []*DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplate {
	return s.AlertTemplate
}

func (s *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplates) SetAlertTemplate(v []*DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplate) *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplates {
	s.AlertTemplate = v
	return s
}

func (s *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplates) Validate() error {
	if s.AlertTemplate != nil {
		for _, item := range s.AlertTemplate {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplate struct {
	// The abbreviation of the Alibaba Cloud service name.
	//
	// example:
	//
	// ecs
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The threshold and the alert level.
	Escalations *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalations `json:"Escalations,omitempty" xml:"Escalations,omitempty" type:"Struct"`
	// The tags of the alert template.
	Labels *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Struct"`
	// The metric name.
	//
	// example:
	//
	// cpu_total
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// The namespace of the Alibaba Cloud service.
	//
	// example:
	//
	// acs_ecs_dashboard
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The method that is used to handle alerts when no monitoring data is found. Valid values:
	//
	// 	- KEEP_LAST_STATE (default): No operation is performed.
	//
	// 	- INSUFFICIENT_DATA: An alert whose content is "Insufficient data" is triggered.
	//
	// 	- OK: The status is considered normal.
	//
	// example:
	//
	// KEEP_LAST_STATE
	NoDataPolicy *string `json:"NoDataPolicy,omitempty" xml:"NoDataPolicy,omitempty"`
	// The name of the alert rule.
	//
	// example:
	//
	// ECS_Rule
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The dimension of the alert. It is an extended field.
	//
	// example:
	//
	// {"disk":"/"}
	Selector    *string `json:"Selector,omitempty" xml:"Selector,omitempty"`
	SilenceTime *int32  `json:"SilenceTime,omitempty" xml:"SilenceTime,omitempty"`
	// The callback URL to which a request is sent when an alert is triggered.
	//
	// example:
	//
	// https://www.aliyun.com
	Webhook *string `json:"Webhook,omitempty" xml:"Webhook,omitempty"`
}

func (s DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplate) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplate) GoString() string {
	return s.String()
}

func (s *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplate) GetCategory() *string {
	return s.Category
}

func (s *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplate) GetEscalations() *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalations {
	return s.Escalations
}

func (s *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplate) GetLabels() *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateLabels {
	return s.Labels
}

func (s *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplate) GetMetricName() *string {
	return s.MetricName
}

func (s *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplate) GetNamespace() *string {
	return s.Namespace
}

func (s *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplate) GetNoDataPolicy() *string {
	return s.NoDataPolicy
}

func (s *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplate) GetRuleName() *string {
	return s.RuleName
}

func (s *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplate) GetSelector() *string {
	return s.Selector
}

func (s *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplate) GetSilenceTime() *int32 {
	return s.SilenceTime
}

func (s *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplate) GetWebhook() *string {
	return s.Webhook
}

func (s *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplate) SetCategory(v string) *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplate {
	s.Category = &v
	return s
}

func (s *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplate) SetEscalations(v *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalations) *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplate {
	s.Escalations = v
	return s
}

func (s *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplate) SetLabels(v *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateLabels) *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplate {
	s.Labels = v
	return s
}

func (s *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplate) SetMetricName(v string) *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplate {
	s.MetricName = &v
	return s
}

func (s *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplate) SetNamespace(v string) *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplate {
	s.Namespace = &v
	return s
}

func (s *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplate) SetNoDataPolicy(v string) *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplate {
	s.NoDataPolicy = &v
	return s
}

func (s *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplate) SetRuleName(v string) *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplate {
	s.RuleName = &v
	return s
}

func (s *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplate) SetSelector(v string) *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplate {
	s.Selector = &v
	return s
}

func (s *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplate) SetSilenceTime(v int32) *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplate {
	s.SilenceTime = &v
	return s
}

func (s *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplate) SetWebhook(v string) *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplate {
	s.Webhook = &v
	return s
}

func (s *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplate) Validate() error {
	if s.Escalations != nil {
		if err := s.Escalations.Validate(); err != nil {
			return err
		}
	}
	if s.Labels != nil {
		if err := s.Labels.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalations struct {
	// The conditions for triggering Critical-level alerts.
	Critical *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalationsCritical `json:"Critical,omitempty" xml:"Critical,omitempty" type:"Struct"`
	// The conditions for triggering Info-level alerts.
	Info *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalationsInfo `json:"Info,omitempty" xml:"Info,omitempty" type:"Struct"`
	// The conditions for triggering Warn-level alerts.
	Warn *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalationsWarn `json:"Warn,omitempty" xml:"Warn,omitempty" type:"Struct"`
}

func (s DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalations) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalations) GoString() string {
	return s.String()
}

func (s *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalations) GetCritical() *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalationsCritical {
	return s.Critical
}

func (s *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalations) GetInfo() *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalationsInfo {
	return s.Info
}

func (s *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalations) GetWarn() *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalationsWarn {
	return s.Warn
}

func (s *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalations) SetCritical(v *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalationsCritical) *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalations {
	s.Critical = v
	return s
}

func (s *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalations) SetInfo(v *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalationsInfo) *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalations {
	s.Info = v
	return s
}

func (s *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalations) SetWarn(v *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalationsWarn) *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalations {
	s.Warn = v
	return s
}

func (s *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalations) Validate() error {
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

type DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalationsCritical struct {
	// The comparison operator that is used to compare the metric value with the threshold for Critical-level alerts. Valid values:
	//
	// 	- GreaterThanOrEqualToThreshold: greater than or equal to the threshold
	//
	// 	- GreaterThanThreshold: greater than the threshold
	//
	// 	- LessThanOrEqualToThreshold: less than or equal to the threshold
	//
	// 	- LessThanThreshold: less than the threshold
	//
	// 	- NotEqualToThreshold: not equal to the threshold
	//
	// 	- GreaterThanYesterday: greater than the metric value at the same time yesterday
	//
	// 	- LessThanYesterday: less than the metric value at the same time yesterday
	//
	// 	- GreaterThanLastWeek: greater than the metric value at the same time last week
	//
	// 	- LessThanLastWeek: less than the metric value at the same time last week
	//
	// 	- GreaterThanLastPeriod: greater than the metric value in the last monitoring cycle
	//
	// 	- LessThanLastPeriod: less than the metric value in the last monitoring cycle
	//
	// example:
	//
	// GreaterThanOrEqualToThreshold
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	// The statistical method for Critical-level alerts.
	//
	// The value of the `Statistics` parameter varies with the cloud service. For more information, see [Appendix 1: Metrics](https://help.aliyun.com/document_detail/163515.html).
	//
	// example:
	//
	// Average
	Statistics *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	// The threshold for Critical-level alerts.
	//
	// example:
	//
	// 90
	Threshold *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	// The consecutive number of times for which the metric value meets the alert condition before a Critical-level alert is triggered.
	//
	// example:
	//
	// 5
	Times *int32 `json:"Times,omitempty" xml:"Times,omitempty"`
}

func (s DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalationsCritical) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalationsCritical) GoString() string {
	return s.String()
}

func (s *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalationsCritical) GetComparisonOperator() *string {
	return s.ComparisonOperator
}

func (s *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalationsCritical) GetStatistics() *string {
	return s.Statistics
}

func (s *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalationsCritical) GetThreshold() *string {
	return s.Threshold
}

func (s *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalationsCritical) GetTimes() *int32 {
	return s.Times
}

func (s *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalationsCritical) SetComparisonOperator(v string) *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalationsCritical {
	s.ComparisonOperator = &v
	return s
}

func (s *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalationsCritical) SetStatistics(v string) *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalationsCritical {
	s.Statistics = &v
	return s
}

func (s *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalationsCritical) SetThreshold(v string) *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalationsCritical {
	s.Threshold = &v
	return s
}

func (s *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalationsCritical) SetTimes(v int32) *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalationsCritical {
	s.Times = &v
	return s
}

func (s *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalationsCritical) Validate() error {
	return dara.Validate(s)
}

type DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalationsInfo struct {
	// The comparison operator that is used to compare the metric value with the threshold for Info-level alerts. Valid values:
	//
	// 	- GreaterThanOrEqualToThreshold: greater than or equal to the threshold
	//
	// 	- GreaterThanThreshold: greater than the threshold
	//
	// 	- LessThanOrEqualToThreshold: less than or equal to the threshold
	//
	// 	- LessThanThreshold: less than the threshold
	//
	// 	- NotEqualToThreshold: not equal to the threshold
	//
	// 	- GreaterThanYesterday: greater than the metric value at the same time yesterday
	//
	// 	- LessThanYesterday: less than the metric value at the same time yesterday
	//
	// 	- GreaterThanLastWeek: greater than the metric value at the same time last week
	//
	// 	- LessThanLastWeek: less than the metric value at the same time last week
	//
	// 	- GreaterThanLastPeriod: greater than the metric value in the last monitoring cycle
	//
	// 	- LessThanLastPeriod: less than the metric value in the last monitoring cycle
	//
	// example:
	//
	// GreaterThanOrEqualToThreshold
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	// The statistical method for Info-level alerts.
	//
	// The value of the `Statistics` parameter varies with the cloud service. For more information, see [Appendix 1: Metrics](https://help.aliyun.com/document_detail/163515.html).
	//
	// example:
	//
	// Average
	Statistics *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	// The threshold for Info-level alerts.
	//
	// example:
	//
	// 90
	Threshold *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	// The consecutive number of times for which the metric value meets the alert condition before an Info-level alert is triggered.
	//
	// example:
	//
	// 3
	Times *int32 `json:"Times,omitempty" xml:"Times,omitempty"`
}

func (s DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalationsInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalationsInfo) GoString() string {
	return s.String()
}

func (s *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalationsInfo) GetComparisonOperator() *string {
	return s.ComparisonOperator
}

func (s *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalationsInfo) GetStatistics() *string {
	return s.Statistics
}

func (s *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalationsInfo) GetThreshold() *string {
	return s.Threshold
}

func (s *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalationsInfo) GetTimes() *int32 {
	return s.Times
}

func (s *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalationsInfo) SetComparisonOperator(v string) *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalationsInfo {
	s.ComparisonOperator = &v
	return s
}

func (s *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalationsInfo) SetStatistics(v string) *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalationsInfo {
	s.Statistics = &v
	return s
}

func (s *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalationsInfo) SetThreshold(v string) *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalationsInfo {
	s.Threshold = &v
	return s
}

func (s *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalationsInfo) SetTimes(v int32) *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalationsInfo {
	s.Times = &v
	return s
}

func (s *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalationsInfo) Validate() error {
	return dara.Validate(s)
}

type DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalationsWarn struct {
	// The comparison operator that is used to compare the metric value with the threshold for Warn-level alerts. Valid values:
	//
	// 	- GreaterThanOrEqualToThreshold: greater than or equal to the threshold
	//
	// 	- GreaterThanThreshold: greater than the threshold
	//
	// 	- LessThanOrEqualToThreshold: less than or equal to the threshold
	//
	// 	- LessThanThreshold: less than the threshold
	//
	// 	- NotEqualToThreshold: not equal to the threshold
	//
	// 	- GreaterThanYesterday: greater than the metric value at the same time yesterday
	//
	// 	- LessThanYesterday: less than the metric value at the same time yesterday
	//
	// 	- GreaterThanLastWeek: greater than the metric value at the same time last week
	//
	// 	- LessThanLastWeek: less than the metric value at the same time last week
	//
	// 	- GreaterThanLastPeriod: greater than the metric value in the last monitoring cycle
	//
	// 	- LessThanLastPeriod: less than the metric value in the last monitoring cycle
	//
	// example:
	//
	// GreaterThanOrEqualToThreshold
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	// The statistical method for Warn-level alerts.
	//
	// The value of the `Statistics` parameter varies with the cloud service. For more information, see [Appendix 1: Metrics](https://help.aliyun.com/document_detail/163515.html).
	//
	// example:
	//
	// Average
	Statistics *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	// The threshold for Warn-level alerts.
	//
	// example:
	//
	// 90
	Threshold *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	// The consecutive number of times for which the metric value meets the alert condition before a Warn-level alert is triggered.
	//
	// example:
	//
	// 3
	Times *int32 `json:"Times,omitempty" xml:"Times,omitempty"`
}

func (s DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalationsWarn) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalationsWarn) GoString() string {
	return s.String()
}

func (s *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalationsWarn) GetComparisonOperator() *string {
	return s.ComparisonOperator
}

func (s *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalationsWarn) GetStatistics() *string {
	return s.Statistics
}

func (s *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalationsWarn) GetThreshold() *string {
	return s.Threshold
}

func (s *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalationsWarn) GetTimes() *int32 {
	return s.Times
}

func (s *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalationsWarn) SetComparisonOperator(v string) *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalationsWarn {
	s.ComparisonOperator = &v
	return s
}

func (s *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalationsWarn) SetStatistics(v string) *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalationsWarn {
	s.Statistics = &v
	return s
}

func (s *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalationsWarn) SetThreshold(v string) *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalationsWarn {
	s.Threshold = &v
	return s
}

func (s *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalationsWarn) SetTimes(v int32) *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalationsWarn {
	s.Times = &v
	return s
}

func (s *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateEscalationsWarn) Validate() error {
	return dara.Validate(s)
}

type DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateLabels struct {
	Labels []*DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateLabelsLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
}

func (s DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateLabels) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateLabels) GoString() string {
	return s.String()
}

func (s *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateLabels) GetLabels() []*DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateLabelsLabels {
	return s.Labels
}

func (s *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateLabels) SetLabels(v []*DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateLabelsLabels) *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateLabels {
	s.Labels = v
	return s
}

func (s *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateLabels) Validate() error {
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

type DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateLabelsLabels struct {
	// The tag key of the alert template.
	//
	// example:
	//
	// label1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the alert template.
	//
	// example:
	//
	// value1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateLabelsLabels) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateLabelsLabels) GoString() string {
	return s.String()
}

func (s *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateLabelsLabels) GetKey() *string {
	return s.Key
}

func (s *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateLabelsLabels) GetValue() *string {
	return s.Value
}

func (s *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateLabelsLabels) SetKey(v string) *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateLabelsLabels {
	s.Key = &v
	return s
}

func (s *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateLabelsLabels) SetValue(v string) *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateLabelsLabels {
	s.Value = &v
	return s
}

func (s *DescribeMetricRuleTemplateAttributeResponseBodyResourceAlertTemplatesAlertTemplateLabelsLabels) Validate() error {
	return dara.Validate(s)
}
