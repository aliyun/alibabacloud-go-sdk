// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOrUpdateAlertRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlertCheckType(v string) *CreateOrUpdateAlertRuleRequest
	GetAlertCheckType() *string
	SetAlertGroup(v int64) *CreateOrUpdateAlertRuleRequest
	GetAlertGroup() *int64
	SetAlertId(v int64) *CreateOrUpdateAlertRuleRequest
	GetAlertId() *int64
	SetAlertName(v string) *CreateOrUpdateAlertRuleRequest
	GetAlertName() *string
	SetAlertPiplines(v string) *CreateOrUpdateAlertRuleRequest
	GetAlertPiplines() *string
	SetAlertRuleContent(v string) *CreateOrUpdateAlertRuleRequest
	GetAlertRuleContent() *string
	SetAlertStatus(v string) *CreateOrUpdateAlertRuleRequest
	GetAlertStatus() *string
	SetAlertType(v string) *CreateOrUpdateAlertRuleRequest
	GetAlertType() *string
	SetAnnotations(v string) *CreateOrUpdateAlertRuleRequest
	GetAnnotations() *string
	SetAutoAddNewApplication(v bool) *CreateOrUpdateAlertRuleRequest
	GetAutoAddNewApplication() *bool
	SetAutoAddTargetConfig(v string) *CreateOrUpdateAlertRuleRequest
	GetAutoAddTargetConfig() *string
	SetCheckCycle(v int64) *CreateOrUpdateAlertRuleRequest
	GetCheckCycle() *int64
	SetClusterId(v string) *CreateOrUpdateAlertRuleRequest
	GetClusterId() *string
	SetDataConfig(v string) *CreateOrUpdateAlertRuleRequest
	GetDataConfig() *string
	SetDuration(v int64) *CreateOrUpdateAlertRuleRequest
	GetDuration() *int64
	SetFilters(v string) *CreateOrUpdateAlertRuleRequest
	GetFilters() *string
	SetLabels(v string) *CreateOrUpdateAlertRuleRequest
	GetLabels() *string
	SetLevel(v string) *CreateOrUpdateAlertRuleRequest
	GetLevel() *string
	SetMarkTags(v []*CreateOrUpdateAlertRuleRequestMarkTags) *CreateOrUpdateAlertRuleRequest
	GetMarkTags() []*CreateOrUpdateAlertRuleRequestMarkTags
	SetMessage(v string) *CreateOrUpdateAlertRuleRequest
	GetMessage() *string
	SetMetricsKey(v string) *CreateOrUpdateAlertRuleRequest
	GetMetricsKey() *string
	SetMetricsType(v string) *CreateOrUpdateAlertRuleRequest
	GetMetricsType() *string
	SetNotice(v string) *CreateOrUpdateAlertRuleRequest
	GetNotice() *string
	SetNotifyMode(v string) *CreateOrUpdateAlertRuleRequest
	GetNotifyMode() *string
	SetNotifyStrategy(v string) *CreateOrUpdateAlertRuleRequest
	GetNotifyStrategy() *string
	SetPids(v string) *CreateOrUpdateAlertRuleRequest
	GetPids() *string
	SetProduct(v string) *CreateOrUpdateAlertRuleRequest
	GetProduct() *string
	SetPromQL(v string) *CreateOrUpdateAlertRuleRequest
	GetPromQL() *string
	SetRegionId(v string) *CreateOrUpdateAlertRuleRequest
	GetRegionId() *string
	SetTags(v []*CreateOrUpdateAlertRuleRequestTags) *CreateOrUpdateAlertRuleRequest
	GetTags() []*CreateOrUpdateAlertRuleRequestTags
}

type CreateOrUpdateAlertRuleRequest struct {
	// The alert check type of the Prometheus alert rule. Valid values:
	//
	// 	- STATIC: a static threshold value. If you set the parameter to STATIC, you must specify the **MetricsKey*	- parameter. For more information, see the **Correspondence between AlertGroup and MetricsKey for Prometheus Service*	- table.
	//
	// 	- CUSTOM: a custom PromQL statement. If you set the parameter to CUSTOM, you must specify the **PromQL**, **Duration**, and **Message*	- parameters to create a Prometheus alert rule.
	//
	// example:
	//
	// STATIC
	AlertCheckType *string `json:"AlertCheckType,omitempty" xml:"AlertCheckType,omitempty"`
	// The alert contact group ID of the Prometheus alert rule. Valid values:
	//
	// 	- \\-1: custom PromQL
	//
	// 	- 1: Kubernetes load
	//
	// 	- 15: Kubernetes node
	//
	// example:
	//
	// -1
	AlertGroup *int64 `json:"AlertGroup,omitempty" xml:"AlertGroup,omitempty"`
	// The ID of the alert rule.
	//
	// 	- If you do not specify this parameter, a new alert rule is created.
	//
	// 	- If you specify this parameter, the specified alert rule is modified.
	//
	// example:
	//
	// 546xxx
	AlertId *int64 `json:"AlertId,omitempty" xml:"AlertId,omitempty"`
	// The name of the alert rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// Alert Rule Demo
	AlertName *string `json:"AlertName,omitempty" xml:"AlertName,omitempty"`
	// The configuration of the alert sending channel. This parameter is used to be compatible with the old version of the rule.
	//
	// example:
	//
	// -
	AlertPiplines *string `json:"AlertPiplines,omitempty" xml:"AlertPiplines,omitempty"`
	// The content of the Application Monitoring or Browser Monitoring alert rule. The following code provides an example of the **AlertRuleContent*	- parameter. For more information about the meaning of each field, see the supplementary description.
	//
	// ```json
	//
	// {
	//
	//     "Condition": "OR",
	//
	//      "AlertRuleItems": [
	//
	//              { "Operator": "CURRENT_LTE",
	//
	//                  "MetricKey": "appstat.jvm.threadcount",
	//
	//                  "Value": 1000,
	//
	//                  "Aggregate": "AVG",
	//
	//                   "N": 10,
	//
	//                   "Tolerability": 169
	//
	//             }
	//
	//        ]
	//
	//   }
	//
	// ```
	//
	// >  The filter conditions specified by the **AlertRuleItems.MetricKey*	- field depends on the value of the **MetricsType*	- parameter. For more information about the types of metrics supported by Application Monitoring and Browser Monitoring and the alert rule fields corresponding to each metric, see the supplementary description.
	//
	// example:
	//
	// { "Condition": "OR", "AlertRuleItems": [ { "Operator": "CURRENT_LTE",  "MetricKey": "appstat.jvm.threadcount",  "Value": 1000,  "Aggregate": "AVG",   "N": 1  }  ]  }
	AlertRuleContent *string `json:"AlertRuleContent,omitempty" xml:"AlertRuleContent,omitempty"`
	// The status of the alert rule. Valid values:
	//
	// 	- RUNNING (default)
	//
	// 	- STOPPED
	//
	// example:
	//
	// RUNNING
	AlertStatus *string `json:"AlertStatus,omitempty" xml:"AlertStatus,omitempty"`
	// The type of the alert rule. Valid values:
	//
	// 	- APPLICATION_MONITORING_ALERT_RULE: an alert rule for Application Monitoring.
	//
	// 	- BROWSER_MONITORING_ALERT_RULE: an alert rule for Browser Monitoring.
	//
	// 	- PROMETHEUS_MONITORING_ALERT_RULE: an alert rule for Managed Service for Prometheus.
	//
	// 	- XTRACE_MONITORING_ALERT_RULE: an alert rule for Managed Service for OpenTelemetry.
	//
	// 	- EBPF_MONITORING_ALERT_RULE: an alert rule for Application Monitoring eBPF Edition.
	//
	// 	- RUM_MONITORING_ALERT_RULE: an alert rule for Real User Monitoring.
	//
	// Valid values:
	//
	// 	- PROMETHEUS_MONITORING_ALERT_RULE
	//
	// 	- APPLICATION_MONITORING_ALERT_RULE
	//
	// 	- BROWSER_MONITORING_ALERT_RULE
	//
	// 	- prometheus monitoring alert
	//
	// 	- application monitoring alert
	//
	// 	- browser monitoring alert
	//
	// 	- XTRACE_MONITORING_ALERT_RULE
	//
	// 	- EBPF_MONITORING_ALERT_RULE
	//
	// 	- RUM_MONITORING_ALERT_RULE
	//
	// This parameter is required.
	//
	// example:
	//
	// APPLICATION_MONITORING_ALERT_RULE
	AlertType *string `json:"AlertType,omitempty" xml:"AlertType,omitempty"`
	// The annotations of the Prometheus alert rule.
	//
	// example:
	//
	// [ { "Value": "PolarDB slow queries", "Name": "_aliyun_display_name" }
	Annotations *string `json:"Annotations,omitempty" xml:"Annotations,omitempty"`
	// Specifies whether to apply the alert rule to new applications that are created in Application Monitoring or Browser Monitoring. Valid values:
	//
	// 	- `true`: enables the health check feature.
	//
	// 	- `false`: disables the automatic backup feature.
	//
	// example:
	//
	// false
	AutoAddNewApplication *bool `json:"AutoAddNewApplication,omitempty" xml:"AutoAddNewApplication,omitempty"`
	// The configurations that are automatically appended to monitor the application based on the specified alert rule.
	//
	// 	- autoAddMatchType:
	//
	//     the matching mode. Valid values: REGULAR and NOT_REGULAR.
	//
	// 	- autoAddMatchExp: the regular expression
	//
	// example:
	//
	// {\\"autoAddMatchType\\":\\"REGULAR\\",\\"autoAddMatchExp\\":\\".*cbw.*\\"}
	AutoAddTargetConfig *string `json:"AutoAddTargetConfig,omitempty" xml:"AutoAddTargetConfig,omitempty"`
	// The interval for checking the alerts in Managed Service for Prometheus.
	//
	// example:
	//
	// 1
	CheckCycle *int64 `json:"CheckCycle,omitempty" xml:"CheckCycle,omitempty"`
	// The ID of the monitored cluster.
	//
	// example:
	//
	// ceba9b9ea5b924dd0b6726d2de6******
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// Data Configuration. The dataRevision field specifies the data repair method when there is no data for the metric.
	//
	// - Fill with zero: 0
	//
	// - Fill with one: 1
	//
	// - Fill with null: 2 (default, does not trigger an alarm)
	//
	// example:
	//
	// {
	//
	//     "dataRevision": 2
	//
	// }
	DataConfig *string `json:"DataConfig,omitempty" xml:"DataConfig,omitempty"`
	// The duration of the Prometheus alert rule, in minutes, in the range of [0,1440].
	//
	// example:
	//
	// 1
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The filter conditions of the Application Monitoring or Browser Monitoring alert rule. Format:
	//
	//     "DimFilters": [
	//
	//     {
	//
	//      "FilterOpt": "ALL",
	//
	//     "FilterValues": [],         //The value of the filter condition.
	//
	//     "FilterKey": "rootIp"     //The key of the filter condition.
	//
	//     }
	//
	//     ]
	//
	// Valid values of **FilterOpt**:
	//
	// 	- STATIC: matches the value of the specified dimension.
	//
	// 	- ALL: traverses all dimension values. Dynamic thresholds do not support traversal.
	//
	// 	- DISABLE: aggregates the values of all dimensions.
	//
	// example:
	//
	// {"DimFilters": [             {               "FilterOpt": "ALL",               "FilterValues": [],               "FilterKey": "rootIp"             }           ]         }
	Filters *string `json:"Filters,omitempty" xml:"Filters,omitempty"`
	// The tags of the Prometheus alert rule.
	//
	// example:
	//
	// [  { "Value": "cms_polardb",             "Name": "_aliyun_cloud_product"           }         ]
	Labels *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// The severity level of the Prometheus alert rule.
	//
	// 	- P1: Alert notifications are sent for major issues that affect the availability of core business, have a huge impact, and may lead to serious consequences.
	//
	// 	- P2: Alert notifications are sent for service errors that affect the system availability with relatively limited impact.
	//
	// 	- P3: Alert notifications are sent for issues that may cause service errors or negative effects, or alert notifications for services that are relatively less important.
	//
	// 	- P4: Alert notifications are sent for low-priority issues that do not affect your business.
	//
	// 	- Default: Alert notifications are sent regardless of alert levels.
	//
	// example:
	//
	// P2
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// Application Tags. Used for application monitoring alert rules, to filter applications associated with alert rules.
	MarkTags []*CreateOrUpdateAlertRuleRequestMarkTags `json:"MarkTags,omitempty" xml:"MarkTags,omitempty" type:"Repeated"`
	// The alert message of the Prometheus alert rule.
	//
	// example:
	//
	// Namespace: {{$labels.namespace}} / Pod: {{$labels.pod_name}} / Container: {{$labels.container}} Memory usage exceeds 80%. Current value: {{ printf \\\\\\\\\\"%.2f\\\\\\\\\\" $value }}%
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The alert metrics. If you set the **AlertCheckType*	- parameter to **STATIC*	- when you create a Prometheus alert rule, you must specify the **MetricsKey*	- parameter.
	//
	// > Alert metrics vary depending on the value of the **AlertGroup*	- parameter. For more information about the correspondence between **AlertGroup*	- and **MetricsKey**, see the supplementary description.
	//
	// example:
	//
	// pop.status.error
	MetricsKey *string `json:"MetricsKey,omitempty" xml:"MetricsKey,omitempty"`
	// The metric type of the Application Monitoring or Browser Monitoring alert rule. For more information, see the following table.
	//
	// example:
	//
	// jvm
	MetricsType *string `json:"MetricsType,omitempty" xml:"MetricsType,omitempty"`
	// The effective time and notification time. This parameter is used to be compatible with the old version of the rule.
	//
	// example:
	//
	// -
	Notice *string `json:"Notice,omitempty" xml:"Notice,omitempty"`
	// The notification mode. You can specify normal mode or simple mode.
	//
	// 	- DIRECTED_MODE
	//
	// 	- NORMAL_MODE
	//
	// example:
	//
	// NORMAL_MODE
	NotifyMode *string `json:"NotifyMode,omitempty" xml:"NotifyMode,omitempty"`
	// The notification policy.
	//
	// 	- If you set this parameter to null, no notification policy is specified. After you create an alert rule, you can create a notification policy and specify match rules and match conditions. For example, you can specify the name of the alert rule as the match condition. When the alert rule is triggered, an alert event is generated and an alert notification is sent to the contacts or contact groups that are specified in the notification policy.
	//
	// 	- To specify a notification policy, set this parameter to the ID of the notification policy. Application Real-Time Monitoring Service (ARMS) automatically adds a match rule to the notification policy and specifies the ID of the alert rule as the match condition. The name of the alert rule is also displayed. This way, the alert events that are generated based on the alert rule can be matched by the specified notification policy.
	//
	// example:
	//
	// 569xxx
	NotifyStrategy *string `json:"NotifyStrategy,omitempty" xml:"NotifyStrategy,omitempty"`
	// The process ID (PID) that is associated with the Application Monitoring or Browser Monitoring alert rule.
	//
	// example:
	//
	// ["b590lhguqs@40d8deedfa9******"]
	Pids *string `json:"Pids,omitempty" xml:"Pids,omitempty"`
	// The product code. If you specify this parameter when you create a Prometheus alert rule, the backend checks whether the product exists.
	//
	// example:
	//
	// clickhouse
	Product *string `json:"Product,omitempty" xml:"Product,omitempty"`
	// The PromQL statement of the Prometheus alert rule.
	//
	// example:
	//
	// node_memory_MemAvailable_bytes{} / node_memory_MemTotal_bytes{} 	- 100
	PromQL *string `json:"PromQL,omitempty" xml:"PromQL,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The list of tags.
	Tags []*CreateOrUpdateAlertRuleRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s CreateOrUpdateAlertRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateOrUpdateAlertRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateAlertRuleRequest) GetAlertCheckType() *string {
	return s.AlertCheckType
}

func (s *CreateOrUpdateAlertRuleRequest) GetAlertGroup() *int64 {
	return s.AlertGroup
}

func (s *CreateOrUpdateAlertRuleRequest) GetAlertId() *int64 {
	return s.AlertId
}

func (s *CreateOrUpdateAlertRuleRequest) GetAlertName() *string {
	return s.AlertName
}

func (s *CreateOrUpdateAlertRuleRequest) GetAlertPiplines() *string {
	return s.AlertPiplines
}

func (s *CreateOrUpdateAlertRuleRequest) GetAlertRuleContent() *string {
	return s.AlertRuleContent
}

func (s *CreateOrUpdateAlertRuleRequest) GetAlertStatus() *string {
	return s.AlertStatus
}

func (s *CreateOrUpdateAlertRuleRequest) GetAlertType() *string {
	return s.AlertType
}

func (s *CreateOrUpdateAlertRuleRequest) GetAnnotations() *string {
	return s.Annotations
}

func (s *CreateOrUpdateAlertRuleRequest) GetAutoAddNewApplication() *bool {
	return s.AutoAddNewApplication
}

func (s *CreateOrUpdateAlertRuleRequest) GetAutoAddTargetConfig() *string {
	return s.AutoAddTargetConfig
}

func (s *CreateOrUpdateAlertRuleRequest) GetCheckCycle() *int64 {
	return s.CheckCycle
}

func (s *CreateOrUpdateAlertRuleRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreateOrUpdateAlertRuleRequest) GetDataConfig() *string {
	return s.DataConfig
}

func (s *CreateOrUpdateAlertRuleRequest) GetDuration() *int64 {
	return s.Duration
}

func (s *CreateOrUpdateAlertRuleRequest) GetFilters() *string {
	return s.Filters
}

func (s *CreateOrUpdateAlertRuleRequest) GetLabels() *string {
	return s.Labels
}

func (s *CreateOrUpdateAlertRuleRequest) GetLevel() *string {
	return s.Level
}

func (s *CreateOrUpdateAlertRuleRequest) GetMarkTags() []*CreateOrUpdateAlertRuleRequestMarkTags {
	return s.MarkTags
}

func (s *CreateOrUpdateAlertRuleRequest) GetMessage() *string {
	return s.Message
}

func (s *CreateOrUpdateAlertRuleRequest) GetMetricsKey() *string {
	return s.MetricsKey
}

func (s *CreateOrUpdateAlertRuleRequest) GetMetricsType() *string {
	return s.MetricsType
}

func (s *CreateOrUpdateAlertRuleRequest) GetNotice() *string {
	return s.Notice
}

func (s *CreateOrUpdateAlertRuleRequest) GetNotifyMode() *string {
	return s.NotifyMode
}

func (s *CreateOrUpdateAlertRuleRequest) GetNotifyStrategy() *string {
	return s.NotifyStrategy
}

func (s *CreateOrUpdateAlertRuleRequest) GetPids() *string {
	return s.Pids
}

func (s *CreateOrUpdateAlertRuleRequest) GetProduct() *string {
	return s.Product
}

func (s *CreateOrUpdateAlertRuleRequest) GetPromQL() *string {
	return s.PromQL
}

func (s *CreateOrUpdateAlertRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateOrUpdateAlertRuleRequest) GetTags() []*CreateOrUpdateAlertRuleRequestTags {
	return s.Tags
}

func (s *CreateOrUpdateAlertRuleRequest) SetAlertCheckType(v string) *CreateOrUpdateAlertRuleRequest {
	s.AlertCheckType = &v
	return s
}

func (s *CreateOrUpdateAlertRuleRequest) SetAlertGroup(v int64) *CreateOrUpdateAlertRuleRequest {
	s.AlertGroup = &v
	return s
}

func (s *CreateOrUpdateAlertRuleRequest) SetAlertId(v int64) *CreateOrUpdateAlertRuleRequest {
	s.AlertId = &v
	return s
}

func (s *CreateOrUpdateAlertRuleRequest) SetAlertName(v string) *CreateOrUpdateAlertRuleRequest {
	s.AlertName = &v
	return s
}

func (s *CreateOrUpdateAlertRuleRequest) SetAlertPiplines(v string) *CreateOrUpdateAlertRuleRequest {
	s.AlertPiplines = &v
	return s
}

func (s *CreateOrUpdateAlertRuleRequest) SetAlertRuleContent(v string) *CreateOrUpdateAlertRuleRequest {
	s.AlertRuleContent = &v
	return s
}

func (s *CreateOrUpdateAlertRuleRequest) SetAlertStatus(v string) *CreateOrUpdateAlertRuleRequest {
	s.AlertStatus = &v
	return s
}

func (s *CreateOrUpdateAlertRuleRequest) SetAlertType(v string) *CreateOrUpdateAlertRuleRequest {
	s.AlertType = &v
	return s
}

func (s *CreateOrUpdateAlertRuleRequest) SetAnnotations(v string) *CreateOrUpdateAlertRuleRequest {
	s.Annotations = &v
	return s
}

func (s *CreateOrUpdateAlertRuleRequest) SetAutoAddNewApplication(v bool) *CreateOrUpdateAlertRuleRequest {
	s.AutoAddNewApplication = &v
	return s
}

func (s *CreateOrUpdateAlertRuleRequest) SetAutoAddTargetConfig(v string) *CreateOrUpdateAlertRuleRequest {
	s.AutoAddTargetConfig = &v
	return s
}

func (s *CreateOrUpdateAlertRuleRequest) SetCheckCycle(v int64) *CreateOrUpdateAlertRuleRequest {
	s.CheckCycle = &v
	return s
}

func (s *CreateOrUpdateAlertRuleRequest) SetClusterId(v string) *CreateOrUpdateAlertRuleRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateOrUpdateAlertRuleRequest) SetDataConfig(v string) *CreateOrUpdateAlertRuleRequest {
	s.DataConfig = &v
	return s
}

func (s *CreateOrUpdateAlertRuleRequest) SetDuration(v int64) *CreateOrUpdateAlertRuleRequest {
	s.Duration = &v
	return s
}

func (s *CreateOrUpdateAlertRuleRequest) SetFilters(v string) *CreateOrUpdateAlertRuleRequest {
	s.Filters = &v
	return s
}

func (s *CreateOrUpdateAlertRuleRequest) SetLabels(v string) *CreateOrUpdateAlertRuleRequest {
	s.Labels = &v
	return s
}

func (s *CreateOrUpdateAlertRuleRequest) SetLevel(v string) *CreateOrUpdateAlertRuleRequest {
	s.Level = &v
	return s
}

func (s *CreateOrUpdateAlertRuleRequest) SetMarkTags(v []*CreateOrUpdateAlertRuleRequestMarkTags) *CreateOrUpdateAlertRuleRequest {
	s.MarkTags = v
	return s
}

func (s *CreateOrUpdateAlertRuleRequest) SetMessage(v string) *CreateOrUpdateAlertRuleRequest {
	s.Message = &v
	return s
}

func (s *CreateOrUpdateAlertRuleRequest) SetMetricsKey(v string) *CreateOrUpdateAlertRuleRequest {
	s.MetricsKey = &v
	return s
}

func (s *CreateOrUpdateAlertRuleRequest) SetMetricsType(v string) *CreateOrUpdateAlertRuleRequest {
	s.MetricsType = &v
	return s
}

func (s *CreateOrUpdateAlertRuleRequest) SetNotice(v string) *CreateOrUpdateAlertRuleRequest {
	s.Notice = &v
	return s
}

func (s *CreateOrUpdateAlertRuleRequest) SetNotifyMode(v string) *CreateOrUpdateAlertRuleRequest {
	s.NotifyMode = &v
	return s
}

func (s *CreateOrUpdateAlertRuleRequest) SetNotifyStrategy(v string) *CreateOrUpdateAlertRuleRequest {
	s.NotifyStrategy = &v
	return s
}

func (s *CreateOrUpdateAlertRuleRequest) SetPids(v string) *CreateOrUpdateAlertRuleRequest {
	s.Pids = &v
	return s
}

func (s *CreateOrUpdateAlertRuleRequest) SetProduct(v string) *CreateOrUpdateAlertRuleRequest {
	s.Product = &v
	return s
}

func (s *CreateOrUpdateAlertRuleRequest) SetPromQL(v string) *CreateOrUpdateAlertRuleRequest {
	s.PromQL = &v
	return s
}

func (s *CreateOrUpdateAlertRuleRequest) SetRegionId(v string) *CreateOrUpdateAlertRuleRequest {
	s.RegionId = &v
	return s
}

func (s *CreateOrUpdateAlertRuleRequest) SetTags(v []*CreateOrUpdateAlertRuleRequestTags) *CreateOrUpdateAlertRuleRequest {
	s.Tags = v
	return s
}

func (s *CreateOrUpdateAlertRuleRequest) Validate() error {
	return dara.Validate(s)
}

type CreateOrUpdateAlertRuleRequestMarkTags struct {
	// The Tag Key.
	//
	// example:
	//
	// service
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The Tag Value.
	//
	// example:
	//
	// proudct
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateOrUpdateAlertRuleRequestMarkTags) String() string {
	return dara.Prettify(s)
}

func (s CreateOrUpdateAlertRuleRequestMarkTags) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateAlertRuleRequestMarkTags) GetKey() *string {
	return s.Key
}

func (s *CreateOrUpdateAlertRuleRequestMarkTags) GetValue() *string {
	return s.Value
}

func (s *CreateOrUpdateAlertRuleRequestMarkTags) SetKey(v string) *CreateOrUpdateAlertRuleRequestMarkTags {
	s.Key = &v
	return s
}

func (s *CreateOrUpdateAlertRuleRequestMarkTags) SetValue(v string) *CreateOrUpdateAlertRuleRequestMarkTags {
	s.Value = &v
	return s
}

func (s *CreateOrUpdateAlertRuleRequestMarkTags) Validate() error {
	return dara.Validate(s)
}

type CreateOrUpdateAlertRuleRequestTags struct {
	// The tag key.
	//
	// example:
	//
	// owner
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// John
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateOrUpdateAlertRuleRequestTags) String() string {
	return dara.Prettify(s)
}

func (s CreateOrUpdateAlertRuleRequestTags) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateAlertRuleRequestTags) GetKey() *string {
	return s.Key
}

func (s *CreateOrUpdateAlertRuleRequestTags) GetValue() *string {
	return s.Value
}

func (s *CreateOrUpdateAlertRuleRequestTags) SetKey(v string) *CreateOrUpdateAlertRuleRequestTags {
	s.Key = &v
	return s
}

func (s *CreateOrUpdateAlertRuleRequestTags) SetValue(v string) *CreateOrUpdateAlertRuleRequestTags {
	s.Value = &v
	return s
}

func (s *CreateOrUpdateAlertRuleRequestTags) Validate() error {
	return dara.Validate(s)
}
