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
	SetAliyunLang(v string) *CreateOrUpdateAlertRuleRequest
	GetAliyunLang() *string
}

type CreateOrUpdateAlertRuleRequest struct {
	// The check type for a Prometheus monitoring alert rule.
	//
	// - `STATIC`: A static threshold. The **MetricsKey*	- parameter is required. For more information, see the description of the **MetricsKey*	- parameter below.
	//
	// - `CUSTOM`: A custom PromQL query. The **PromQL**, **Duration**, and **Message*	- parameters are required.
	//
	// example:
	//
	// STATIC
	AlertCheckType *string `json:"AlertCheckType,omitempty" xml:"AlertCheckType,omitempty"`
	// The alert group ID for the Prometheus alert rule. Valid values:
	//
	// - `-1`: Custom PromQL
	//
	// - `1`: Kubernetes workloads
	//
	// - `15`: Kubernetes nodes
	//
	// example:
	//
	// -1
	AlertGroup *int64 `json:"AlertGroup,omitempty" xml:"AlertGroup,omitempty"`
	// The ID of the alert rule.
	//
	// - Omit this parameter to create a new alert rule.
	//
	// - Specify an ID to modify an existing alert rule.
	//
	// example:
	//
	// 546xxx
	AlertId *int64 `json:"AlertId,omitempty" xml:"AlertId,omitempty"`
	// The alert rule name.
	//
	// This parameter is required.
	//
	// example:
	//
	// Alert Rule Demo
	AlertName *string `json:"AlertName,omitempty" xml:"AlertName,omitempty"`
	// The alert pipeline configuration. Used for compatibility with legacy alert rules.
	//
	// example:
	//
	// -
	AlertPiplines *string `json:"AlertPiplines,omitempty" xml:"AlertPiplines,omitempty"`
	// The content of the alert rule for application monitoring or browser monitoring. The following is a template for the **AlertRuleContent*	- parameter. For a description of the fields in the template, see the supplementary information below this table.
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
	// > The available fields for **AlertRuleItems.MetricKey*	- depend on the **MetricsType*	- value. For information about the metric types supported by application monitoring and browser monitoring and their corresponding alert rule fields, see the supplementary information below this table.
	//
	// example:
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
	AlertRuleContent *string `json:"AlertRuleContent,omitempty" xml:"AlertRuleContent,omitempty"`
	// The status of the alert rule. Valid values:
	//
	// - `RUNNING`: The alert rule is running. (Default)
	//
	// - `STOPPED`: The alert rule is stopped.
	//
	// example:
	//
	// RUNNING
	AlertStatus *string `json:"AlertStatus,omitempty" xml:"AlertStatus,omitempty"`
	// The type of the alert rule. Valid values:
	//
	// - `APPLICATION_MONITORING_ALERT_RULE`: For application monitoring.
	//
	// - `BROWSER_MONITORING_ALERT_RULE`: For browser monitoring.
	//
	// - `PROMETHEUS_MONITORING_ALERT_RULE`: For Prometheus monitoring.
	//
	// - `XTRACE_MONITORING_ALERT_RULE`: For Tracing Analysis (OpenTelemetry edition).
	//
	// - `EBPF_MONITORING_ALERT_RULE`: For eBPF monitoring.
	//
	// - `RUM_MONITORING_ALERT_RULE`: For real user monitoring (RUM).
	//
	// This parameter is required.
	//
	// example:
	//
	// APPLICATION_MONITORING_ALERT_RULE
	AlertType *string `json:"AlertType,omitempty" xml:"AlertType,omitempty"`
	// Annotations to add to the Prometheus alert rule. Specify as a JSON string representing an array of objects, each with Name and Value keys.
	//
	// example:
	//
	// [ { "Value": "PolarDB 慢查询数量",             "Name": "_aliyun_display_name"           }
	Annotations *string `json:"Annotations,omitempty" xml:"Annotations,omitempty"`
	// Determines whether to automatically apply this alert rule to new applications. This applies only to application monitoring and browser monitoring rules.
	//
	// - `true`: enables the feature.
	//
	// - `false`: disables the feature.
	//
	// example:
	//
	// false
	AutoAddNewApplication *bool `json:"AutoAddNewApplication,omitempty" xml:"AutoAddNewApplication,omitempty"`
	// The configuration for automatically adding applications to an application monitoring alert rule. Specify this parameter as a JSON string with the following fields:
	//
	// - `autoAddMatchType`: The matching method. Can be `REGULAR` (matches the regular expression) or `NOT_REGULAR` (does not match the regular expression).
	//
	//   Match type: Regular expression match (REGULAR) / Not a regular expression match (NOT_REGULAR)
	//
	// - `autoAddMatchExp`: The regular expression.
	//
	// example:
	//
	// {\\"autoAddMatchType\\":\\"REGULAR\\",\\"autoAddMatchExp\\":\\".*cbw.*\\"}
	AutoAddTargetConfig *string `json:"AutoAddTargetConfig,omitempty" xml:"AutoAddTargetConfig,omitempty"`
	// The check interval for the Prometheus alert rule.
	//
	// example:
	//
	// 1
	CheckCycle *int64 `json:"CheckCycle,omitempty" xml:"CheckCycle,omitempty"`
	// The cluster ID for the Prometheus monitoring alert rule.
	//
	// example:
	//
	// ceba9b9ea5b924dd0b6726d2de6******
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The data configuration. The dataRevision field specifies how to handle missing metric data.
	//
	// - `0`: Fills the data with 0.
	//
	// - `1`: Fills the data with 1.
	//
	// - `2`: Fills the data with null. This is the default and does not trigger an alert.
	//
	// example:
	//
	// {
	//
	//     "dataRevision": 2
	//
	// }
	DataConfig *string `json:"DataConfig,omitempty" xml:"DataConfig,omitempty"`
	// The period, in minutes, that a condition must be true before a Prometheus alert is triggered. Valid values: 0 to 1440.
	//
	// example:
	//
	// 1
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The filters for an application monitoring or browser monitoring alert rule.
	//
	// Specify this parameter as a JSON string in the following format:
	//
	// ```
	//
	// "DimFilters": [
	//
	// {
	//
	//  "FilterOpt": "ALL",
	//
	//  "FilterValues": [],         // The filter value.
	//
	//  "FilterKey": "rootIp"     // The filter key.
	//
	// }
	//
	// ]
	//
	// ```
	//
	// Valid values for **FilterOpt**:
	//
	// - `STATIC`: Matches a fixed dimension value.
	//
	// - `ALL`: Iterates over all dimension values. Note: This option is not supported for range detection.
	//
	// - `DISABLE`: Aggregates all dimension values by summing them.
	//
	// example:
	//
	// {"DimFilters": [             {               "FilterOpt": "ALL",               "FilterValues": [],               "FilterKey": "rootIp"             }           ]         }
	Filters *string `json:"Filters,omitempty" xml:"Filters,omitempty"`
	// Labels to add to the Prometheus alert rule. Specify as a JSON string representing an array of objects, each with Name and Value keys.
	//
	// example:
	//
	// [  { "Value": "cms_polardb",             "Name": "_aliyun_cloud_product"           }         ]
	Labels *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// The severity level for the Prometheus alert rule.
	//
	// - `P1`: Critical. For major issues that affect core business availability with a wide impact and severe consequences.
	//
	// - `P2`: Warning. For issues that cause partial service failures or affect system availability with a limited scope.
	//
	// - `P3`: Info. For potential issues or alerts from non-critical services.
	//
	// - `P4`: Low priority. Used for informational alerts that require attention but do not affect services.
	//
	// - `Default`: The default level, used when no specific severity is required.
	//
	// example:
	//
	// P2
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// Application tags used to filter applications in application monitoring alert rules.
	MarkTags []*CreateOrUpdateAlertRuleRequestMarkTags `json:"MarkTags,omitempty" xml:"MarkTags,omitempty" type:"Repeated"`
	// The alert message for the Prometheus alert rule.
	//
	// example:
	//
	// 命名空间: {{$labels.namespace}} / Pod: {{$labels.pod_name}} / 容器: {{$labels.container}} 内存使用率超过80%, 当前值{{ printf \\\\\\"%.2f\\\\\\" $value }}%
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The alert metric. This parameter is required for Prometheus alert rules when **AlertCheckType*	- is **STATIC**.
	//
	// > The available alert metrics vary based on the value of **AlertGroup**. For information about the mapping between **AlertGroup*	- and **MetricsKey**, see the supplementary information below this table.
	//
	// example:
	//
	// pop.status.error
	MetricsKey *string `json:"MetricsKey,omitempty" xml:"MetricsKey,omitempty"`
	// The alert metric type for application monitoring or browser monitoring alert rules. For more information, see the table below.
	//
	// example:
	//
	// jvm
	MetricsType *string `json:"MetricsType,omitempty" xml:"MetricsType,omitempty"`
	// The effective time and notification time. Used for compatibility with legacy alert rules.
	//
	// example:
	//
	// -
	Notice *string `json:"Notice,omitempty" xml:"Notice,omitempty"`
	// The notification mode. Valid values:
	//
	// - `DIRECTED_MODE`: Directed mode.
	//
	// - `NORMAL_MODE`: Normal mode.
	//
	// example:
	//
	// NORMAL_MODE
	NotifyMode *string `json:"NotifyMode,omitempty" xml:"NotifyMode,omitempty"`
	// The notification policy.
	//
	// - `null`: Does not associate the alert rule with a notification policy. You can associate them later by creating a notification policy with a matching rule, for example, based on the alert rule\\"s name. When the alert rule is triggered, alert events are sent to the contacts or contact groups specified in the matching notification policy.
	//
	// - A notification policy ID: Associates the alert rule with a specific notification policy. ARMS automatically adds a matching rule to the policy that uses the alert rule\\"s ID. This ensures that alert events from this rule are always processed by the specified policy.
	//
	// example:
	//
	// 569xxx
	NotifyStrategy *string `json:"NotifyStrategy,omitempty" xml:"NotifyStrategy,omitempty"`
	// The PIDs of applications for an application monitoring or browser monitoring alert rule. Specify as a JSON array of strings.
	//
	// example:
	//
	// ["b590lhguqs@40d8deedfa9******"]
	Pids *string `json:"Pids,omitempty" xml:"Pids,omitempty"`
	// Required for Prometheus alert rules. Used to filter by cloud service. The specified product name must be valid.
	//
	// example:
	//
	// clickhouse
	Product *string `json:"Product,omitempty" xml:"Product,omitempty"`
	// The PromQL expression to evaluate.
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
	// The tags to add to the alert rule. These are standard Alibaba Cloud resource tags.
	Tags []*CreateOrUpdateAlertRuleRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The language of the response.
	AliyunLang *string `json:"aliyunLang,omitempty" xml:"aliyunLang,omitempty"`
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

func (s *CreateOrUpdateAlertRuleRequest) GetAliyunLang() *string {
	return s.AliyunLang
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

func (s *CreateOrUpdateAlertRuleRequest) SetAliyunLang(v string) *CreateOrUpdateAlertRuleRequest {
	s.AliyunLang = &v
	return s
}

func (s *CreateOrUpdateAlertRuleRequest) Validate() error {
	if s.MarkTags != nil {
		for _, item := range s.MarkTags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateOrUpdateAlertRuleRequestMarkTags struct {
	// The tag key.
	//
	// example:
	//
	// service
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
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
