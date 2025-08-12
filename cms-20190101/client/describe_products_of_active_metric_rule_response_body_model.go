// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeProductsOfActiveMetricRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAllProductInitMetricRuleList(v *DescribeProductsOfActiveMetricRuleResponseBodyAllProductInitMetricRuleList) *DescribeProductsOfActiveMetricRuleResponseBody
	GetAllProductInitMetricRuleList() *DescribeProductsOfActiveMetricRuleResponseBodyAllProductInitMetricRuleList
	SetCode(v int32) *DescribeProductsOfActiveMetricRuleResponseBody
	GetCode() *int32
	SetDatapoints(v string) *DescribeProductsOfActiveMetricRuleResponseBody
	GetDatapoints() *string
	SetMessage(v string) *DescribeProductsOfActiveMetricRuleResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeProductsOfActiveMetricRuleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeProductsOfActiveMetricRuleResponseBody
	GetSuccess() *bool
}

type DescribeProductsOfActiveMetricRuleResponseBody struct {
	// The information about the services for which one-click alert is enabled.
	AllProductInitMetricRuleList *DescribeProductsOfActiveMetricRuleResponseBodyAllProductInitMetricRuleList `json:"AllProductInitMetricRuleList,omitempty" xml:"AllProductInitMetricRuleList,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// >  The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The information about the services for which the initiative alert feature is enabled. Services are separated with commas (,). Valid values:
	//
	// 	- ECS: Elastic Compute Service (ECS)
	//
	// 	- rds: ApsaraDB RDS
	//
	// 	- slb: Server Load Balancer (SLB)
	//
	// 	- redis_standard: Redis Open-Source Edition (standard architecture)
	//
	// 	- redis_sharding: Redis Open-Source Edition (cluster architecture)
	//
	// 	- redis_splitrw: Redis Open-Source Edition (read/write splitting architecture)
	//
	// 	- mongodb: ApsaraDB for MongoDB of the replica set architecture
	//
	// 	- mongodb_sharding: ApsaraDB for MongoDB of the sharded cluster architecture
	//
	// 	- hbase: ApsaraDB for HBase
	//
	// 	- elasticsearch: Elasticsearch
	//
	// 	- opensearch: OpenSearch
	//
	// example:
	//
	// ecs,rds
	Datapoints *string `json:"Datapoints,omitempty" xml:"Datapoints,omitempty"`
	// The error message.
	//
	// example:
	//
	// The Request is not authorization.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F82E6667-7811-4BA0-842F-5B2DC42BBAAD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeProductsOfActiveMetricRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeProductsOfActiveMetricRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeProductsOfActiveMetricRuleResponseBody) GetAllProductInitMetricRuleList() *DescribeProductsOfActiveMetricRuleResponseBodyAllProductInitMetricRuleList {
	return s.AllProductInitMetricRuleList
}

func (s *DescribeProductsOfActiveMetricRuleResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DescribeProductsOfActiveMetricRuleResponseBody) GetDatapoints() *string {
	return s.Datapoints
}

func (s *DescribeProductsOfActiveMetricRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeProductsOfActiveMetricRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeProductsOfActiveMetricRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeProductsOfActiveMetricRuleResponseBody) SetAllProductInitMetricRuleList(v *DescribeProductsOfActiveMetricRuleResponseBodyAllProductInitMetricRuleList) *DescribeProductsOfActiveMetricRuleResponseBody {
	s.AllProductInitMetricRuleList = v
	return s
}

func (s *DescribeProductsOfActiveMetricRuleResponseBody) SetCode(v int32) *DescribeProductsOfActiveMetricRuleResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeProductsOfActiveMetricRuleResponseBody) SetDatapoints(v string) *DescribeProductsOfActiveMetricRuleResponseBody {
	s.Datapoints = &v
	return s
}

func (s *DescribeProductsOfActiveMetricRuleResponseBody) SetMessage(v string) *DescribeProductsOfActiveMetricRuleResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeProductsOfActiveMetricRuleResponseBody) SetRequestId(v string) *DescribeProductsOfActiveMetricRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeProductsOfActiveMetricRuleResponseBody) SetSuccess(v bool) *DescribeProductsOfActiveMetricRuleResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeProductsOfActiveMetricRuleResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeProductsOfActiveMetricRuleResponseBodyAllProductInitMetricRuleList struct {
	AllProductInitMetricRule []*DescribeProductsOfActiveMetricRuleResponseBodyAllProductInitMetricRuleListAllProductInitMetricRule `json:"AllProductInitMetricRule,omitempty" xml:"AllProductInitMetricRule,omitempty" type:"Repeated"`
}

func (s DescribeProductsOfActiveMetricRuleResponseBodyAllProductInitMetricRuleList) String() string {
	return dara.Prettify(s)
}

func (s DescribeProductsOfActiveMetricRuleResponseBodyAllProductInitMetricRuleList) GoString() string {
	return s.String()
}

func (s *DescribeProductsOfActiveMetricRuleResponseBodyAllProductInitMetricRuleList) GetAllProductInitMetricRule() []*DescribeProductsOfActiveMetricRuleResponseBodyAllProductInitMetricRuleListAllProductInitMetricRule {
	return s.AllProductInitMetricRule
}

func (s *DescribeProductsOfActiveMetricRuleResponseBodyAllProductInitMetricRuleList) SetAllProductInitMetricRule(v []*DescribeProductsOfActiveMetricRuleResponseBodyAllProductInitMetricRuleListAllProductInitMetricRule) *DescribeProductsOfActiveMetricRuleResponseBodyAllProductInitMetricRuleList {
	s.AllProductInitMetricRule = v
	return s
}

func (s *DescribeProductsOfActiveMetricRuleResponseBodyAllProductInitMetricRuleList) Validate() error {
	return dara.Validate(s)
}

type DescribeProductsOfActiveMetricRuleResponseBodyAllProductInitMetricRuleListAllProductInitMetricRule struct {
	// The initial alert rules that are generated after one-click alert is enabled for a service.
	AlertInitConfigList *DescribeProductsOfActiveMetricRuleResponseBodyAllProductInitMetricRuleListAllProductInitMetricRuleAlertInitConfigList `json:"AlertInitConfigList,omitempty" xml:"AlertInitConfigList,omitempty" type:"Struct"`
	// The abbreviation of the service name.
	//
	// example:
	//
	// ecs
	Product *string `json:"Product,omitempty" xml:"Product,omitempty"`
}

func (s DescribeProductsOfActiveMetricRuleResponseBodyAllProductInitMetricRuleListAllProductInitMetricRule) String() string {
	return dara.Prettify(s)
}

func (s DescribeProductsOfActiveMetricRuleResponseBodyAllProductInitMetricRuleListAllProductInitMetricRule) GoString() string {
	return s.String()
}

func (s *DescribeProductsOfActiveMetricRuleResponseBodyAllProductInitMetricRuleListAllProductInitMetricRule) GetAlertInitConfigList() *DescribeProductsOfActiveMetricRuleResponseBodyAllProductInitMetricRuleListAllProductInitMetricRuleAlertInitConfigList {
	return s.AlertInitConfigList
}

func (s *DescribeProductsOfActiveMetricRuleResponseBodyAllProductInitMetricRuleListAllProductInitMetricRule) GetProduct() *string {
	return s.Product
}

func (s *DescribeProductsOfActiveMetricRuleResponseBodyAllProductInitMetricRuleListAllProductInitMetricRule) SetAlertInitConfigList(v *DescribeProductsOfActiveMetricRuleResponseBodyAllProductInitMetricRuleListAllProductInitMetricRuleAlertInitConfigList) *DescribeProductsOfActiveMetricRuleResponseBodyAllProductInitMetricRuleListAllProductInitMetricRule {
	s.AlertInitConfigList = v
	return s
}

func (s *DescribeProductsOfActiveMetricRuleResponseBodyAllProductInitMetricRuleListAllProductInitMetricRule) SetProduct(v string) *DescribeProductsOfActiveMetricRuleResponseBodyAllProductInitMetricRuleListAllProductInitMetricRule {
	s.Product = &v
	return s
}

func (s *DescribeProductsOfActiveMetricRuleResponseBodyAllProductInitMetricRuleListAllProductInitMetricRule) Validate() error {
	return dara.Validate(s)
}

type DescribeProductsOfActiveMetricRuleResponseBodyAllProductInitMetricRuleListAllProductInitMetricRuleAlertInitConfigList struct {
	AlertInitConfig []*DescribeProductsOfActiveMetricRuleResponseBodyAllProductInitMetricRuleListAllProductInitMetricRuleAlertInitConfigListAlertInitConfig `json:"AlertInitConfig,omitempty" xml:"AlertInitConfig,omitempty" type:"Repeated"`
}

func (s DescribeProductsOfActiveMetricRuleResponseBodyAllProductInitMetricRuleListAllProductInitMetricRuleAlertInitConfigList) String() string {
	return dara.Prettify(s)
}

func (s DescribeProductsOfActiveMetricRuleResponseBodyAllProductInitMetricRuleListAllProductInitMetricRuleAlertInitConfigList) GoString() string {
	return s.String()
}

func (s *DescribeProductsOfActiveMetricRuleResponseBodyAllProductInitMetricRuleListAllProductInitMetricRuleAlertInitConfigList) GetAlertInitConfig() []*DescribeProductsOfActiveMetricRuleResponseBodyAllProductInitMetricRuleListAllProductInitMetricRuleAlertInitConfigListAlertInitConfig {
	return s.AlertInitConfig
}

func (s *DescribeProductsOfActiveMetricRuleResponseBodyAllProductInitMetricRuleListAllProductInitMetricRuleAlertInitConfigList) SetAlertInitConfig(v []*DescribeProductsOfActiveMetricRuleResponseBodyAllProductInitMetricRuleListAllProductInitMetricRuleAlertInitConfigListAlertInitConfig) *DescribeProductsOfActiveMetricRuleResponseBodyAllProductInitMetricRuleListAllProductInitMetricRuleAlertInitConfigList {
	s.AlertInitConfig = v
	return s
}

func (s *DescribeProductsOfActiveMetricRuleResponseBodyAllProductInitMetricRuleListAllProductInitMetricRuleAlertInitConfigList) Validate() error {
	return dara.Validate(s)
}

type DescribeProductsOfActiveMetricRuleResponseBodyAllProductInitMetricRuleListAllProductInitMetricRuleAlertInitConfigListAlertInitConfig struct {
	// The operator that is used to compare the metric value with the threshold for Warn-level alerts.
	//
	// Valid values:
	//
	// 	- LessThanThreshold: less than the threshold
	//
	// 	- GreaterThanLastWeek: greater than the metric value at the same time last week
	//
	// 	- LessThanOrEqualToThreshold: less than or equal to the threshold
	//
	// 	- NotEqualToThreshold: does not equal to the threshold
	//
	// 	- GreaterThanLastPeriod: greater than the metric value in the last monitoring cycle
	//
	// 	- GreaterThanYesterday: greater than the metric value at the same time yesterday
	//
	// 	- LessThanYesterday: less than the metric value at the same time yesterday
	//
	// 	- LessThanLastWeek: less than the metric value at the same time last week
	//
	// 	- GreaterThanOrEqualToThreshold: greater than or equal to the threshold
	//
	// 	- GreaterThanThreshold: greater than the threshold
	//
	// 	- LessThanLastPeriod: less than the metric value in the last monitoring cycle
	//
	// example:
	//
	// GreaterThanOrEqualToThreshold
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	// The consecutive number of times for which the metric value meets the alert condition before an alert is triggered.
	//
	// example:
	//
	// 3
	EvaluationCount *string `json:"EvaluationCount,omitempty" xml:"EvaluationCount,omitempty"`
	// The alert level.
	//
	// Valid values:
	//
	// 	- INFO
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- WARN
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- CRITICAL
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// CRITICAL
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// The metric name. For more information, see [Appendix 1: Metrics](https://help.aliyun.com/document_detail/163515.html).
	//
	// example:
	//
	// cpu_total
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// The namespace of the service. For more information, see [Appendix 1: Metrics](https://help.aliyun.com/document_detail/163515.html).
	//
	// example:
	//
	// acs_rds_dashboard
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The aggregation period of monitoring data. Unit: minutes. For more information, see [Appendix 1: Metrics](https://help.aliyun.com/document_detail/163515.html).
	//
	// example:
	//
	// 1m
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// The method used to calculate metric values that trigger alerts. For more information, see [Appendix 1: Metrics](https://help.aliyun.com/document_detail/163515.html).
	//
	// example:
	//
	// Average
	Statistics *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	// The alert threshold.
	//
	// example:
	//
	// 90
	Threshold *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s DescribeProductsOfActiveMetricRuleResponseBodyAllProductInitMetricRuleListAllProductInitMetricRuleAlertInitConfigListAlertInitConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeProductsOfActiveMetricRuleResponseBodyAllProductInitMetricRuleListAllProductInitMetricRuleAlertInitConfigListAlertInitConfig) GoString() string {
	return s.String()
}

func (s *DescribeProductsOfActiveMetricRuleResponseBodyAllProductInitMetricRuleListAllProductInitMetricRuleAlertInitConfigListAlertInitConfig) GetComparisonOperator() *string {
	return s.ComparisonOperator
}

func (s *DescribeProductsOfActiveMetricRuleResponseBodyAllProductInitMetricRuleListAllProductInitMetricRuleAlertInitConfigListAlertInitConfig) GetEvaluationCount() *string {
	return s.EvaluationCount
}

func (s *DescribeProductsOfActiveMetricRuleResponseBodyAllProductInitMetricRuleListAllProductInitMetricRuleAlertInitConfigListAlertInitConfig) GetLevel() *string {
	return s.Level
}

func (s *DescribeProductsOfActiveMetricRuleResponseBodyAllProductInitMetricRuleListAllProductInitMetricRuleAlertInitConfigListAlertInitConfig) GetMetricName() *string {
	return s.MetricName
}

func (s *DescribeProductsOfActiveMetricRuleResponseBodyAllProductInitMetricRuleListAllProductInitMetricRuleAlertInitConfigListAlertInitConfig) GetNamespace() *string {
	return s.Namespace
}

func (s *DescribeProductsOfActiveMetricRuleResponseBodyAllProductInitMetricRuleListAllProductInitMetricRuleAlertInitConfigListAlertInitConfig) GetPeriod() *string {
	return s.Period
}

func (s *DescribeProductsOfActiveMetricRuleResponseBodyAllProductInitMetricRuleListAllProductInitMetricRuleAlertInitConfigListAlertInitConfig) GetStatistics() *string {
	return s.Statistics
}

func (s *DescribeProductsOfActiveMetricRuleResponseBodyAllProductInitMetricRuleListAllProductInitMetricRuleAlertInitConfigListAlertInitConfig) GetThreshold() *string {
	return s.Threshold
}

func (s *DescribeProductsOfActiveMetricRuleResponseBodyAllProductInitMetricRuleListAllProductInitMetricRuleAlertInitConfigListAlertInitConfig) SetComparisonOperator(v string) *DescribeProductsOfActiveMetricRuleResponseBodyAllProductInitMetricRuleListAllProductInitMetricRuleAlertInitConfigListAlertInitConfig {
	s.ComparisonOperator = &v
	return s
}

func (s *DescribeProductsOfActiveMetricRuleResponseBodyAllProductInitMetricRuleListAllProductInitMetricRuleAlertInitConfigListAlertInitConfig) SetEvaluationCount(v string) *DescribeProductsOfActiveMetricRuleResponseBodyAllProductInitMetricRuleListAllProductInitMetricRuleAlertInitConfigListAlertInitConfig {
	s.EvaluationCount = &v
	return s
}

func (s *DescribeProductsOfActiveMetricRuleResponseBodyAllProductInitMetricRuleListAllProductInitMetricRuleAlertInitConfigListAlertInitConfig) SetLevel(v string) *DescribeProductsOfActiveMetricRuleResponseBodyAllProductInitMetricRuleListAllProductInitMetricRuleAlertInitConfigListAlertInitConfig {
	s.Level = &v
	return s
}

func (s *DescribeProductsOfActiveMetricRuleResponseBodyAllProductInitMetricRuleListAllProductInitMetricRuleAlertInitConfigListAlertInitConfig) SetMetricName(v string) *DescribeProductsOfActiveMetricRuleResponseBodyAllProductInitMetricRuleListAllProductInitMetricRuleAlertInitConfigListAlertInitConfig {
	s.MetricName = &v
	return s
}

func (s *DescribeProductsOfActiveMetricRuleResponseBodyAllProductInitMetricRuleListAllProductInitMetricRuleAlertInitConfigListAlertInitConfig) SetNamespace(v string) *DescribeProductsOfActiveMetricRuleResponseBodyAllProductInitMetricRuleListAllProductInitMetricRuleAlertInitConfigListAlertInitConfig {
	s.Namespace = &v
	return s
}

func (s *DescribeProductsOfActiveMetricRuleResponseBodyAllProductInitMetricRuleListAllProductInitMetricRuleAlertInitConfigListAlertInitConfig) SetPeriod(v string) *DescribeProductsOfActiveMetricRuleResponseBodyAllProductInitMetricRuleListAllProductInitMetricRuleAlertInitConfigListAlertInitConfig {
	s.Period = &v
	return s
}

func (s *DescribeProductsOfActiveMetricRuleResponseBodyAllProductInitMetricRuleListAllProductInitMetricRuleAlertInitConfigListAlertInitConfig) SetStatistics(v string) *DescribeProductsOfActiveMetricRuleResponseBodyAllProductInitMetricRuleListAllProductInitMetricRuleAlertInitConfigListAlertInitConfig {
	s.Statistics = &v
	return s
}

func (s *DescribeProductsOfActiveMetricRuleResponseBodyAllProductInitMetricRuleListAllProductInitMetricRuleAlertInitConfigListAlertInitConfig) SetThreshold(v string) *DescribeProductsOfActiveMetricRuleResponseBodyAllProductInitMetricRuleListAllProductInitMetricRuleAlertInitConfigListAlertInitConfig {
	s.Threshold = &v
	return s
}

func (s *DescribeProductsOfActiveMetricRuleResponseBodyAllProductInitMetricRuleListAllProductInitMetricRuleAlertInitConfigListAlertInitConfig) Validate() error {
	return dara.Validate(s)
}
