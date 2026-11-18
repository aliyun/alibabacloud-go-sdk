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
	// The list of alert templates.
	AlertTemplates []*ModifyMetricRuleTemplateRequestAlertTemplates `json:"AlertTemplates,omitempty" xml:"AlertTemplates,omitempty" type:"Repeated"`
	// The description of the alert template.
	//
	// example:
	//
	// ECS_template1
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the alert template.
	//
	// example:
	//
	// test123
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The version of the alert template to be modified.
	//
	// >This parameter is automatically updated based on the number of modifications.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	RestVersion *int64 `json:"RestVersion,omitempty" xml:"RestVersion,omitempty"`
	// The ID of the alert template.
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
	// The abbreviation of the Alibaba Cloud service name. Valid values of N: 1 to 200. Valid values:
	//
	// - ecs: Alibaba Cloud and non-Alibaba Cloud hosts.
	//
	// - rds: ApsaraDB RDS database.
	//
	// - ads: AnalyticDB database.
	//
	// - slb: load balancing.
	//
	// - vpc: Elastic IP Address (EIP).
	//
	// - apigateway: API Gateway.
	//
	// - cdn: Content Delivery Network (CDN).
	//
	// - cs: Container Service for Swarm.
	//
	// - dcdn: Dynamic Route for CDN (DCDN).
	//
	// - ddos: Anti-DDoS.
	//
	// - eip: Elastic IP Address (EIP).
	//
	// - elasticsearch: Elasticsearch.
	//
	// - emr: E-MapReduce (EMR).
	//
	// - ess: elastic scaling.
	//
	// - hbase: ApsaraDB for HBase database.
	//
	// - iot_edge: IoT Edge.
	//
	// - kvstore_sharding: ApsaraDB for Redis Cluster Edition.
	//
	// - kvstore_splitrw: ApsaraDB for Redis Read/Write Splitting Edition.
	//
	// - kvstore_standard: ApsaraDB for Redis Standard Edition.
	//
	// - memcache: ApsaraDB for Memcache database.
	//
	// - mns: Message Service (MNS).
	//
	// - mongodb: MongoDB replica set instance.
	//
	// - mongodb_cluster: MongoDB cluster edition.
	//
	// - mongodb_sharding: MongoDB sharded cluster.
	//
	// - mq_topic: Message Service Topic.
	//
	// - ocs: legacy cloud database Memcache.
	//
	// - opensearch: OpenSearch.
	//
	// - oss: Object Storage Service (OSS).
	//
	// - polardb: PolarDB database.
	//
	// - petadata: HybridDB for MySQL.
	//
	// - scdn: secure acceleration.
	//
	// - sharebandwidthpackages: Internet Shared Bandwidth.
	//
	// - sls: Simple Log Service (SLS).
	//
	// - vpn: VPN Gateway.
	//
	// example:
	//
	// ecs
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The name of the metric. Valid values of N: 1 to 200.
	//
	// >For more information about the valid values of the metric name, see [DescribeSystemEventMetaList](https://help.aliyun.com/document_detail/114972.html).
	//
	// example:
	//
	// cpu_total
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// The namespace of the Alibaba Cloud service, which is used to distinguish different services. Valid values of N: 1 to 200.
	//
	// Format: acs_<service name>.
	//
	// >To obtain the namespace of each Alibaba Cloud service, call the [DescribeProjectMeta](https://help.aliyun.com/document_detail/114916.html) operation.
	//
	// example:
	//
	// acs_ecs_dashboard
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The statistical period of the monitoring data. Valid values of N: 1 to 200.
	//
	// >If you set this parameter to 300, the system collects and aggregates monitoring data every 300 seconds. If monitoring data is reported every minute, the system aggregates the data of 5 minutes, such as the average, maximum, and minimum values. Do not specify this parameter separately because the statistical period may cause unexpected alerts.
	//
	// example:
	//
	// 60
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The name of the alert rule. Valid values of N: 1 to 200.
	//
	// example:
	//
	// rule1
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The options for the extended dimension fields. Valid values of N: 1 to 200.
	//
	// For example, if the alert template is applied to an application group, this parameter is set to `{"disk":"/"}`, and the MetricName parameter is set to `DiskUtilization`, the root partition (`"/"`) of the disks for all instances in the application group is associated with the corresponding alert rule.
	//
	// >For more information about the valid values of the extended dimension fields, see [DescribeSystemEventMetaList](https://help.aliyun.com/document_detail/114972.html).
	//
	// example:
	//
	// {"disk":"/"}
	Selector *string `json:"Selector,omitempty" xml:"Selector,omitempty"`
	// The webhook URL to which a POST request is sent when an alert is triggered.
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
	// The comparison operator for the critical-level alert threshold. Valid values:
	//
	// - GreaterThanOrEqualToThreshold: greater than or equal to the threshold.
	//
	// - GreaterThanThreshold: greater than the threshold.
	//
	// - LessThanOrEqualToThreshold: less than or equal to the threshold.
	//
	// - LessThanThreshold: less than the threshold.
	//
	// - NotEqualToThreshold: not equal to the threshold.
	//
	// - GreaterThanYesterday: greater than the metric value at the same time yesterday.
	//
	// - LessThanYesterday: less than the metric value at the same time yesterday.
	//
	// - GreaterThanLastWeek: greater than the metric value at the same time last week.
	//
	// - LessThanLastWeek: less than the metric value at the same time last week.
	//
	// - GreaterThanLastPeriod: greater than the metric value in the previous period.
	//
	// - LessThanLastPeriod: less than the metric value in the previous period.
	//
	// Valid values of N: 1 to 200.
	//
	// > You must specify at least one of the following alert levels: Critical, Warn, and Info. The Statistics, ComparisonOperator, Threshold, and Times parameters must be specified at the same time for the specified alert level.
	//
	// example:
	//
	// GreaterThanOrEqualToThreshold
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	N                  *int32  `json:"N,omitempty" xml:"N,omitempty"`
	// The statistical method for critical-level alerts.
	//
	// Valid values of N: 1 to 200.
	//
	// The valid values of this parameter are determined by the `Statistics` column that corresponds to the `MetricName` parameter of the specified Alibaba Cloud service, such as Maximum, Minimum, and Average. For more information about how to obtain the valid values of this parameter, see [cloud service monitoring metrics](https://help.aliyun.com/document_detail/163515.html).
	//
	// > You must specify at least one of the following alert levels: Critical, Warn, and Info. The Statistics, ComparisonOperator, Threshold, and Times parameters must be specified at the same time for the specified alert level.
	//
	// example:
	//
	// Average
	Statistics *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	// The threshold that triggers a critical-level alert notification.
	//
	// Valid values of N: 1 to 200.
	//
	// > You must specify at least one of the following alert levels: Critical, Warn, and Info. The Statistics, ComparisonOperator, Threshold, and Times parameters must be specified at the same time for the specified alert level.
	//
	// example:
	//
	// 30
	Threshold *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	// The number of times that the monitoring metrics must reach the alert threshold to send a critical-level alert notification.
	//
	// Valid values of N: 1 to 200.
	//
	// > You must specify at least one of the following alert levels: Critical, Warn, and Info. The Statistics, ComparisonOperator, Threshold, and Times parameters must be specified at the same time for the specified alert level.
	//
	// example:
	//
	// 2
	Times *int32 `json:"Times,omitempty" xml:"Times,omitempty"`
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
	// The comparison operator for the info-level alert threshold. Valid values:
	//
	// - GreaterThanOrEqualToThreshold: greater than or equal to the threshold.
	//
	// - GreaterThanThreshold: greater than the threshold.
	//
	// - LessThanOrEqualToThreshold: less than or equal to the threshold.
	//
	// - LessThanThreshold: less than the threshold.
	//
	// - NotEqualToThreshold: not equal to the threshold.
	//
	// - GreaterThanYesterday: greater than the metric value at the same time yesterday.
	//
	// - LessThanYesterday: less than the metric value at the same time yesterday.
	//
	// - GreaterThanLastWeek: greater than the metric value at the same time last week.
	//
	// - LessThanLastWeek: less than the metric value at the same time last week.
	//
	// - GreaterThanLastPeriod: greater than the metric value in the previous period.
	//
	// - LessThanLastPeriod: less than the metric value in the previous period.
	//
	// Valid values of N: 1 to 200.
	//
	// > You must specify at least one of the following alert levels: Critical, Warn, and Info. The Statistics, ComparisonOperator, Threshold, and Times parameters must be specified at the same time for the specified alert level.
	//
	// example:
	//
	// GreaterThanOrEqualToThreshold
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	N                  *int32  `json:"N,omitempty" xml:"N,omitempty"`
	// The statistical method for info-level alerts.
	//
	// Valid values of N: 1 to 200.
	//
	// The valid values of this parameter are determined by the `Statistics` column that corresponds to the `MetricName` parameter of the specified Alibaba Cloud service, such as Maximum, Minimum, and Average. For more information about how to obtain the valid values of this parameter, see [cloud service monitoring metrics](https://help.aliyun.com/document_detail/163515.html).
	//
	// > You must specify at least one of the following alert levels: Critical, Warn, and Info. The Statistics, ComparisonOperator, Threshold, and Times parameters must be specified at the same time for the specified alert level.
	//
	// example:
	//
	// Average
	Statistics *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	// The threshold for info-level alerts.
	//
	// Valid values of N: 1 to 200.
	//
	// > You must specify at least one of the following alert levels: Critical, Warn, and Info. The Statistics, ComparisonOperator, Threshold, and Times parameters must be specified at the same time for the specified alert level.
	//
	// example:
	//
	// 20
	Threshold *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	// The number of times that the monitoring metrics must reach the alert threshold to send an info-level alert notification.
	//
	// Valid values of N: 1 to 200.
	//
	// > You must specify at least one of the following alert levels: Critical, Warn, and Info. The Statistics, ComparisonOperator, Threshold, and Times parameters must be specified at the same time for the specified alert level.
	//
	// example:
	//
	// 1
	Times *int32 `json:"Times,omitempty" xml:"Times,omitempty"`
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
	// The comparison operator for the warning-level alert threshold. Valid values:
	//
	// - GreaterThanOrEqualToThreshold: greater than or equal to the threshold.
	//
	// - GreaterThanThreshold: greater than the threshold.
	//
	// - LessThanOrEqualToThreshold: less than or equal to the threshold.
	//
	// - LessThanThreshold: less than the threshold.
	//
	// - NotEqualToThreshold: not equal to the threshold.
	//
	// - GreaterThanYesterday: greater than the metric value at the same time yesterday.
	//
	// - LessThanYesterday: less than the metric value at the same time yesterday.
	//
	// - GreaterThanLastWeek: greater than the metric value at the same time last week.
	//
	// - LessThanLastWeek: less than the metric value at the same time last week.
	//
	// - GreaterThanLastPeriod: greater than the metric value in the previous period.
	//
	// - LessThanLastPeriod: less than the metric value in the previous period.
	//
	// Valid values of N: 1 to 200.
	//
	// > You must specify at least one of the following alert levels: Critical, Warn, and Info. The Statistics, ComparisonOperator, Threshold, and Times parameters must be specified at the same time for the specified alert level.
	//
	// example:
	//
	// GreaterThanOrEqualToThreshold
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	N                  *int32  `json:"N,omitempty" xml:"N,omitempty"`
	// The statistical method for warning-level alerts.
	//
	// Valid values of N: 1 to 200.
	//
	// The valid values of this parameter are determined by the `Statistics` column that corresponds to the `MetricName` parameter of the specified Alibaba Cloud service, such as Maximum, Minimum, and Average. For more information about how to obtain the valid values of this parameter, see [cloud service monitoring metrics](https://help.aliyun.com/document_detail/163515.html).
	//
	// > You must specify at least one of the following alert levels: Critical, Warn, and Info. The Statistics, ComparisonOperator, Threshold, and Times parameters must be specified at the same time for the specified alert level.
	//
	// example:
	//
	// Average
	Statistics *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	// The threshold for warning-level alerts.
	//
	// Valid values of N: 1 to 200.
	//
	// > You must specify at least one of the following alert levels: Critical, Warn, and Info. The Statistics, ComparisonOperator, Threshold, and Times parameters must be specified at the same time for the specified alert level.
	//
	// example:
	//
	// 30
	Threshold *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	// The number of times that the monitoring metrics must reach the alert threshold to send a warning-level alert notification.
	//
	// Valid values of N: 1 to 200.
	//
	// > You must specify at least one of the following alert levels: Critical, Warn, and Info. The Statistics, ComparisonOperator, Threshold, and Times parameters must be specified at the same time for the specified alert level.
	//
	// example:
	//
	// 1
	Times *int32 `json:"Times,omitempty" xml:"Times,omitempty"`
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
