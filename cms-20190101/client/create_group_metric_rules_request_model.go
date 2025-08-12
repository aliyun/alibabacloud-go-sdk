// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGroupMetricRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v int64) *CreateGroupMetricRulesRequest
	GetGroupId() *int64
	SetGroupMetricRules(v []*CreateGroupMetricRulesRequestGroupMetricRules) *CreateGroupMetricRulesRequest
	GetGroupMetricRules() []*CreateGroupMetricRulesRequestGroupMetricRules
	SetRegionId(v string) *CreateGroupMetricRulesRequest
	GetRegionId() *string
}

type CreateGroupMetricRulesRequest struct {
	// The ID of the application group.
	//
	// For information about how to obtain the ID of an application group, see [DescribeMonitorGroups](https://help.aliyun.com/document_detail/115032.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456
	GroupId          *int64                                           `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupMetricRules []*CreateGroupMetricRulesRequestGroupMetricRules `json:"GroupMetricRules,omitempty" xml:"GroupMetricRules,omitempty" type:"Repeated"`
	RegionId         *string                                          `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateGroupMetricRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateGroupMetricRulesRequest) GoString() string {
	return s.String()
}

func (s *CreateGroupMetricRulesRequest) GetGroupId() *int64 {
	return s.GroupId
}

func (s *CreateGroupMetricRulesRequest) GetGroupMetricRules() []*CreateGroupMetricRulesRequestGroupMetricRules {
	return s.GroupMetricRules
}

func (s *CreateGroupMetricRulesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateGroupMetricRulesRequest) SetGroupId(v int64) *CreateGroupMetricRulesRequest {
	s.GroupId = &v
	return s
}

func (s *CreateGroupMetricRulesRequest) SetGroupMetricRules(v []*CreateGroupMetricRulesRequestGroupMetricRules) *CreateGroupMetricRulesRequest {
	s.GroupMetricRules = v
	return s
}

func (s *CreateGroupMetricRulesRequest) SetRegionId(v string) *CreateGroupMetricRulesRequest {
	s.RegionId = &v
	return s
}

func (s *CreateGroupMetricRulesRequest) Validate() error {
	return dara.Validate(s)
}

type CreateGroupMetricRulesRequestGroupMetricRules struct {
	Escalations *CreateGroupMetricRulesRequestGroupMetricRulesEscalations `json:"Escalations,omitempty" xml:"Escalations,omitempty" type:"Struct"`
	// The name of the cloud service. Valid values of N: 1 to 200. Valid value:
	//
	// 	- PolarDB: PolarDB
	//
	// 	- NewBGPDDoS: Anti-DDoS Pro
	//
	// 	- IoTDevice: IoT Platform
	//
	// 	- DRDS: Distributed Relational Database Service (DRDS)
	//
	// 	- VS: Video Surveillance System
	//
	// 	- AMQP: Alibaba Cloud Message Queue for AMQP
	//
	// 	- ADS: AnalyticDB
	//
	// 	- APIGateway: API Gateway
	//
	// 	- InternetSharedBandwidth: EIP Bandwidth Plan
	//
	// 	- CDN: Alibaba Cloud Content Delivery Network (CDN)
	//
	// 	- CEN: Cloud Enterprise Network (CEN)
	//
	// 	- DCDN: Dynamic Route for CDN (DCDN)
	//
	// 	- DDoS: Anti-DDoS
	//
	// 	- ECS: Elastic Compute Service (ECS)
	//
	// 	- DirectMail: Direct Mail
	//
	// 	- Elasticsearch: Elasticsearch
	//
	// 	- EMR: E-MapReduce (EMR)
	//
	// 	- ESS: Auto Scaling
	//
	// 	- FunctionCompute: Function Compute
	//
	// 	- RealtimeCompute: Realtime Compute for Apache Flink
	//
	// 	- GlobalAcceleration: Global Accelerator (GA)
	//
	// 	- Hbase: ApsaraDB for HBase
	//
	// 	- TSDB: Time Series Database (TSDB)
	//
	// 	- IPv6trans: IPv6 Translation Service
	//
	// 	- Kafka: Message Queue for Apache Kafka
	//
	// 	- Kubernetes: Container Service for Kubernetes (ACK)
	//
	// 	- KVstore: ApsaraDB for Redis
	//
	// 	- MNS: Message Service (MNS)
	//
	// 	- MongoDB: ApsaraDB for MongoDB
	//
	// 	- MQ: Message Queue
	//
	// 	- NAT: NAT Gateway
	//
	// 	- OpenAd: Open Ad
	//
	// 	- OpenSearch: Open Search
	//
	// 	- OSS: Object Storage Service (OSS)
	//
	// 	- PCDN: P2P CDN
	//
	// 	- petadata: HybridDB for MySQL
	//
	// 	- RDS: ApsaraDB RDS
	//
	// 	- SCDN: Secure CDN
	//
	// 	- SLB: Server Load Balancer (SLB)
	//
	// 	- SLS: Log Service
	//
	// 	- VideoLive: ApsaraVideo Live
	//
	// 	- VOD: ApsaraVideo VOD
	//
	// 	- EIP: Elastic IP Address (EIP)
	//
	// 	- VPN: VPN Gateway
	//
	// 	- AIRec: Artificial Intelligence Recommendation
	//
	// 	- GPDB: AnalyticDB for PostgreSQL
	//
	// 	- DBS: Database Backup (DBS)
	//
	// 	- SAG: Smart Access Gateway (SAG)
	//
	// 	- Memcache: ApsaraDB for Memcache
	//
	// 	- IOT_EDGE: Link IoT Edge
	//
	// 	- OCS: ApsaraDB for Memcache (previous version)
	//
	// 	- VPC: Express Connect
	//
	// 	- EHPC: Elastic High Performance Computing (E-HPC)
	//
	// 	- MPS: ApsaraVideo Media Processing
	//
	// 	- ENS: Edge Node Service (ENS)
	//
	// 	- MaxCompute_Prepay: MaxCompute
	//
	// 	- IoT_Kubernetes: Edge Application Hosting
	//
	// 	- CMS: CloudMonitor
	//
	// 	- batchcomputenew: Batch Compute
	//
	// 	- HBaseUE: ApsaraDB for HBase Performance-enhanced Edition
	//
	// 	- UIS: Ultimate Internet Service (UIS)
	//
	// 	- nls: Intelligent Speech Interaction
	//
	// 	- ots: Tablestore
	//
	// 	- NAS: File Storage NAS
	//
	// 	- ECI: Elastic Container Instance (ECI)
	//
	// 	- OpenAPI: OpenAPI Explorer
	//
	// 	- pvtzpost: Alibaba Cloud DNS PrivateZone
	//
	// 	- blinkonk8s: Flink on Kubernetes
	//
	// 	- FunctionFlow: Serverless Workflow (SWF)
	//
	// 	- SMC: Server Migration Center (SMC)
	//
	// 	- ddosbgp: Anti-DDoS Origin
	//
	// 	- baas: Blockchain as a Service
	//
	// 	- privatelink: PrivateLink
	//
	// 	- cds: ApsaraDB for Cassandra
	//
	// 	- DDH: Dedicated Host
	//
	// 	- RocketMQ: Message Queue for Apache RocketMQ
	//
	// 	- ECC: Express Cloud Connect
	//
	// 	- hbaseserverless: ApsaraDB for HBase Serverless Edition
	//
	// 	- mns_tmp: Message Service
	//
	// 	- hdr: Hybrid Disaster Recovery (HDR)
	//
	// 	- hbr: Hybrid Backup Recovery (HBR)
	//
	// 	- ADB: AnalyticDB for MySQL V3.0
	//
	// 	- tag: Tag Service
	//
	// 	- GDB: Graph Database
	//
	// 	- WAF: Web Application Firewall (WAF)
	//
	// 	- hcs_sgw: Cloud Storage Gateway (CSG)
	//
	// 	- ipv6gateway: IPv6 Gateway
	//
	// 	- RDS_SAR: ApsaraDB Exclusive Host Group
	//
	// 	- learn: Machine Learning Platform for AI
	//
	// 	- ROS: Resource Orchestration Service (ROS)
	//
	// 	- OOS: Operation Orchestration Service (OOS)
	//
	// 	- bds: Data Synchronization for HBase
	//
	// 	- cfw: Cloud Firewall
	//
	// 	- ddosDip: Anti-DDoS Premium
	//
	// 	- datahub: DataHub
	//
	// 	- hologres: Hologres
	//
	// 	- ExpressConnect: Express Connect
	//
	// 	- dbfs: Database File System (DBFS)
	//
	// 	- clickhouse: ApsaraDB for ClickHouse
	//
	// 	- k8s: Container Service for Kubernetes (ACK)
	//
	// 	- DTS: Data Transmission Service (DTS)
	//
	// 	- AnycastEIP: Anycast Elastic IP Address
	//
	// 	- Lindorm: ApsaraDB for Lindorm
	//
	// 	- config: Cloud Config
	//
	// 	- spark: Databricks DataInsight (DDI)
	//
	// 	- serverless: Serverless App Engine (SAE)
	//
	// 	- alb: Application Load Balancer (ALB)
	//
	// 	- oceanbase: ApsaraDB for OceanBase
	//
	// 	- KMS: Key Management Service (KMS)
	//
	// 	- lvwang: Content Moderation
	//
	// 	- LinkVisual: LinkVisual
	//
	// 	- tair: ApsaraDB for Redis Enhanced Edition (Tair)
	//
	// 	- dlf: Data Lake Formation (DLF)
	//
	// 	- networkmonitor: Site Monitoring
	//
	// 	- pnc: Physical Network Change
	//
	// 	- AIS: Alibaba Cloud Infrastructure
	//
	// 	- cloudgame: Cloud Gaming Platform
	//
	// 	- RTC: Real-Time Communication
	//
	// 	- cloudbox: CloudBox
	//
	// 	- actiontrail: ActionTrail
	//
	// 	- cc: Cloud Connector
	//
	// 	- disk: Elastic Block Storage (EBS)
	//
	// 	- easygene: Genomics Computing Platform
	//
	// 	- cloudphone: Elastic Cloud Phone
	//
	// 	- BMS: Bare Metal Management Service
	//
	// 	- swas: Simple Application Server
	//
	// 	- AvailabilityMonitoring: Availability Monitoring of CloudMonitor
	//
	// This parameter is required.
	//
	// example:
	//
	// ECS
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The alert contact groups. Valid values of N: 1 to 200.
	//
	// For information about how to obtain alert contact groups, see [DescribeContactGroupList](https://help.aliyun.com/document_detail/114922.html).
	//
	// example:
	//
	// ECS_Group
	ContactGroups *string `json:"ContactGroups,omitempty" xml:"ContactGroups,omitempty"`
	// The dimension of the alert rule. Valid values of N: 1 to 200.
	//
	// Set the value to a set of key-value pairs, for example, `userId:120886317861****` or `instanceId:i-m5e1qg6uo38rztr4****`.
	//
	// example:
	//
	// [{"instanceId":"i-m5e1qg6uo38rztr4****"}]
	Dimensions *string `json:"Dimensions,omitempty" xml:"Dimensions,omitempty"`
	// The time period during which the alert rule is effective. Valid values of N: 1 to 200.
	//
	// example:
	//
	// 05:31-23:59
	EffectiveInterval *string `json:"EffectiveInterval,omitempty" xml:"EffectiveInterval,omitempty"`
	// The subject of the alert notification email. Valid values of N: 1 to 200.
	EmailSubject *string `json:"EmailSubject,omitempty" xml:"EmailSubject,omitempty"`
	// The interval at which CloudMonitor checks whether the alert rule is triggered. Valid values of N: 1 to 200.
	//
	// Unit: seconds. The default value is the lowest frequency at which the metric is polled.
	//
	// >  We recommend that you set the interval to the data aggregation period. If the interval is shorter than the data aggregation period, alerts cannot be triggered due to insufficient data.
	//
	// example:
	//
	// 60
	Interval *string                                                `json:"Interval,omitempty" xml:"Interval,omitempty"`
	Labels   []*CreateGroupMetricRulesRequestGroupMetricRulesLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// The name of the metric. Valid values of N: 1 to 200.
	//
	// For information about how to obtain the name of a metric, see [DescribeMetricMetaList](https://help.aliyun.com/document_detail/98846.html) or [Appendix 1: Metrics](https://help.aliyun.com/document_detail/163515.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cpu_total
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// The namespace of the cloud service. Valid values of N: 1 to 200.
	//
	// For information about how to obtain the namespace of a cloud service, see [DescribeMetricMetaList](https://help.aliyun.com/document_detail/98846.html) or [Appendix 1: Metrics](https://help.aliyun.com/document_detail/163515.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// acs_ecs_dashboard
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The method that is used to handle alerts when no monitoring data is found. Valid values of N: 1 to 200. Valid value:
	//
	// 	- KEEP_LAST_STATE (default value): No operation is performed.
	//
	// 	- INSUFFICIENT_DATA: An alert whose content is "Insufficient data" is triggered.
	//
	// 	- OK: The alert rule has no active alerts.
	//
	// example:
	//
	// KEEP_LAST_STATE
	NoDataPolicy *string `json:"NoDataPolicy,omitempty" xml:"NoDataPolicy,omitempty"`
	// The time period during which the alert rule is ineffective. Valid values of N: 1 to 200.
	//
	// example:
	//
	// 00:00-05:30
	NoEffectiveInterval *string `json:"NoEffectiveInterval,omitempty" xml:"NoEffectiveInterval,omitempty"`
	// example:
	//
	// {
	//
	//       "NotSendOK": true
	//
	// }
	Options *string `json:"Options,omitempty" xml:"Options,omitempty"`
	// The aggregation period of the metric data. Valid values of N: 1 to 200.
	//
	// Set the `Period` parameter to an integral multiple of 60. Unit: seconds. Default value: 300.
	//
	// example:
	//
	// 60
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// The ID of the alert rule. Valid values of N: 1 to 200.
	//
	// This parameter is required.
	//
	// example:
	//
	// 456789
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The name of the alert rule. Valid values of N: 1 to 200.
	//
	// This parameter is required.
	//
	// example:
	//
	// ECS_Rule1
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The mute period during which new alerts are not sent even if the trigger conditions are met. Valid values of N: 1 to 200.
	//
	// Unit: seconds. Default value: 86400. Minimum value: 3600.
	//
	// example:
	//
	// 86400
	SilenceTime *int32 `json:"SilenceTime,omitempty" xml:"SilenceTime,omitempty"`
	// The callback URL. Valid values of N: 1 to 200.
	//
	// The callback URL must be accessible over the Internet. CloudMonitor pushes an alert notification to the specified callback URL by sending an HTTP POST request. Only the HTTP protocol is supported.
	//
	// example:
	//
	// https://www.aliyun.com
	Webhook *string `json:"Webhook,omitempty" xml:"Webhook,omitempty"`
}

func (s CreateGroupMetricRulesRequestGroupMetricRules) String() string {
	return dara.Prettify(s)
}

func (s CreateGroupMetricRulesRequestGroupMetricRules) GoString() string {
	return s.String()
}

func (s *CreateGroupMetricRulesRequestGroupMetricRules) GetEscalations() *CreateGroupMetricRulesRequestGroupMetricRulesEscalations {
	return s.Escalations
}

func (s *CreateGroupMetricRulesRequestGroupMetricRules) GetCategory() *string {
	return s.Category
}

func (s *CreateGroupMetricRulesRequestGroupMetricRules) GetContactGroups() *string {
	return s.ContactGroups
}

func (s *CreateGroupMetricRulesRequestGroupMetricRules) GetDimensions() *string {
	return s.Dimensions
}

func (s *CreateGroupMetricRulesRequestGroupMetricRules) GetEffectiveInterval() *string {
	return s.EffectiveInterval
}

func (s *CreateGroupMetricRulesRequestGroupMetricRules) GetEmailSubject() *string {
	return s.EmailSubject
}

func (s *CreateGroupMetricRulesRequestGroupMetricRules) GetInterval() *string {
	return s.Interval
}

func (s *CreateGroupMetricRulesRequestGroupMetricRules) GetLabels() []*CreateGroupMetricRulesRequestGroupMetricRulesLabels {
	return s.Labels
}

func (s *CreateGroupMetricRulesRequestGroupMetricRules) GetMetricName() *string {
	return s.MetricName
}

func (s *CreateGroupMetricRulesRequestGroupMetricRules) GetNamespace() *string {
	return s.Namespace
}

func (s *CreateGroupMetricRulesRequestGroupMetricRules) GetNoDataPolicy() *string {
	return s.NoDataPolicy
}

func (s *CreateGroupMetricRulesRequestGroupMetricRules) GetNoEffectiveInterval() *string {
	return s.NoEffectiveInterval
}

func (s *CreateGroupMetricRulesRequestGroupMetricRules) GetOptions() *string {
	return s.Options
}

func (s *CreateGroupMetricRulesRequestGroupMetricRules) GetPeriod() *string {
	return s.Period
}

func (s *CreateGroupMetricRulesRequestGroupMetricRules) GetRuleId() *string {
	return s.RuleId
}

func (s *CreateGroupMetricRulesRequestGroupMetricRules) GetRuleName() *string {
	return s.RuleName
}

func (s *CreateGroupMetricRulesRequestGroupMetricRules) GetSilenceTime() *int32 {
	return s.SilenceTime
}

func (s *CreateGroupMetricRulesRequestGroupMetricRules) GetWebhook() *string {
	return s.Webhook
}

func (s *CreateGroupMetricRulesRequestGroupMetricRules) SetEscalations(v *CreateGroupMetricRulesRequestGroupMetricRulesEscalations) *CreateGroupMetricRulesRequestGroupMetricRules {
	s.Escalations = v
	return s
}

func (s *CreateGroupMetricRulesRequestGroupMetricRules) SetCategory(v string) *CreateGroupMetricRulesRequestGroupMetricRules {
	s.Category = &v
	return s
}

func (s *CreateGroupMetricRulesRequestGroupMetricRules) SetContactGroups(v string) *CreateGroupMetricRulesRequestGroupMetricRules {
	s.ContactGroups = &v
	return s
}

func (s *CreateGroupMetricRulesRequestGroupMetricRules) SetDimensions(v string) *CreateGroupMetricRulesRequestGroupMetricRules {
	s.Dimensions = &v
	return s
}

func (s *CreateGroupMetricRulesRequestGroupMetricRules) SetEffectiveInterval(v string) *CreateGroupMetricRulesRequestGroupMetricRules {
	s.EffectiveInterval = &v
	return s
}

func (s *CreateGroupMetricRulesRequestGroupMetricRules) SetEmailSubject(v string) *CreateGroupMetricRulesRequestGroupMetricRules {
	s.EmailSubject = &v
	return s
}

func (s *CreateGroupMetricRulesRequestGroupMetricRules) SetInterval(v string) *CreateGroupMetricRulesRequestGroupMetricRules {
	s.Interval = &v
	return s
}

func (s *CreateGroupMetricRulesRequestGroupMetricRules) SetLabels(v []*CreateGroupMetricRulesRequestGroupMetricRulesLabels) *CreateGroupMetricRulesRequestGroupMetricRules {
	s.Labels = v
	return s
}

func (s *CreateGroupMetricRulesRequestGroupMetricRules) SetMetricName(v string) *CreateGroupMetricRulesRequestGroupMetricRules {
	s.MetricName = &v
	return s
}

func (s *CreateGroupMetricRulesRequestGroupMetricRules) SetNamespace(v string) *CreateGroupMetricRulesRequestGroupMetricRules {
	s.Namespace = &v
	return s
}

func (s *CreateGroupMetricRulesRequestGroupMetricRules) SetNoDataPolicy(v string) *CreateGroupMetricRulesRequestGroupMetricRules {
	s.NoDataPolicy = &v
	return s
}

func (s *CreateGroupMetricRulesRequestGroupMetricRules) SetNoEffectiveInterval(v string) *CreateGroupMetricRulesRequestGroupMetricRules {
	s.NoEffectiveInterval = &v
	return s
}

func (s *CreateGroupMetricRulesRequestGroupMetricRules) SetOptions(v string) *CreateGroupMetricRulesRequestGroupMetricRules {
	s.Options = &v
	return s
}

func (s *CreateGroupMetricRulesRequestGroupMetricRules) SetPeriod(v string) *CreateGroupMetricRulesRequestGroupMetricRules {
	s.Period = &v
	return s
}

func (s *CreateGroupMetricRulesRequestGroupMetricRules) SetRuleId(v string) *CreateGroupMetricRulesRequestGroupMetricRules {
	s.RuleId = &v
	return s
}

func (s *CreateGroupMetricRulesRequestGroupMetricRules) SetRuleName(v string) *CreateGroupMetricRulesRequestGroupMetricRules {
	s.RuleName = &v
	return s
}

func (s *CreateGroupMetricRulesRequestGroupMetricRules) SetSilenceTime(v int32) *CreateGroupMetricRulesRequestGroupMetricRules {
	s.SilenceTime = &v
	return s
}

func (s *CreateGroupMetricRulesRequestGroupMetricRules) SetWebhook(v string) *CreateGroupMetricRulesRequestGroupMetricRules {
	s.Webhook = &v
	return s
}

func (s *CreateGroupMetricRulesRequestGroupMetricRules) Validate() error {
	return dara.Validate(s)
}

type CreateGroupMetricRulesRequestGroupMetricRulesEscalations struct {
	Critical *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsCritical `json:"Critical,omitempty" xml:"Critical,omitempty" type:"Struct"`
	Info     *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsInfo     `json:"Info,omitempty" xml:"Info,omitempty" type:"Struct"`
	Warn     *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsWarn     `json:"Warn,omitempty" xml:"Warn,omitempty" type:"Struct"`
}

func (s CreateGroupMetricRulesRequestGroupMetricRulesEscalations) String() string {
	return dara.Prettify(s)
}

func (s CreateGroupMetricRulesRequestGroupMetricRulesEscalations) GoString() string {
	return s.String()
}

func (s *CreateGroupMetricRulesRequestGroupMetricRulesEscalations) GetCritical() *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsCritical {
	return s.Critical
}

func (s *CreateGroupMetricRulesRequestGroupMetricRulesEscalations) GetInfo() *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsInfo {
	return s.Info
}

func (s *CreateGroupMetricRulesRequestGroupMetricRulesEscalations) GetWarn() *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsWarn {
	return s.Warn
}

func (s *CreateGroupMetricRulesRequestGroupMetricRulesEscalations) SetCritical(v *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsCritical) *CreateGroupMetricRulesRequestGroupMetricRulesEscalations {
	s.Critical = v
	return s
}

func (s *CreateGroupMetricRulesRequestGroupMetricRulesEscalations) SetInfo(v *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsInfo) *CreateGroupMetricRulesRequestGroupMetricRulesEscalations {
	s.Info = v
	return s
}

func (s *CreateGroupMetricRulesRequestGroupMetricRulesEscalations) SetWarn(v *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsWarn) *CreateGroupMetricRulesRequestGroupMetricRulesEscalations {
	s.Warn = v
	return s
}

func (s *CreateGroupMetricRulesRequestGroupMetricRulesEscalations) Validate() error {
	return dara.Validate(s)
}

type CreateGroupMetricRulesRequestGroupMetricRulesEscalationsCritical struct {
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	N                  *string `json:"N,omitempty" xml:"N,omitempty"`
	PreCondition       *string `json:"PreCondition,omitempty" xml:"PreCondition,omitempty"`
	Statistics         *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	Threshold          *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	Times              *int32  `json:"Times,omitempty" xml:"Times,omitempty"`
}

func (s CreateGroupMetricRulesRequestGroupMetricRulesEscalationsCritical) String() string {
	return dara.Prettify(s)
}

func (s CreateGroupMetricRulesRequestGroupMetricRulesEscalationsCritical) GoString() string {
	return s.String()
}

func (s *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsCritical) GetComparisonOperator() *string {
	return s.ComparisonOperator
}

func (s *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsCritical) GetN() *string {
	return s.N
}

func (s *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsCritical) GetPreCondition() *string {
	return s.PreCondition
}

func (s *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsCritical) GetStatistics() *string {
	return s.Statistics
}

func (s *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsCritical) GetThreshold() *string {
	return s.Threshold
}

func (s *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsCritical) GetTimes() *int32 {
	return s.Times
}

func (s *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsCritical) SetComparisonOperator(v string) *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsCritical {
	s.ComparisonOperator = &v
	return s
}

func (s *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsCritical) SetN(v string) *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsCritical {
	s.N = &v
	return s
}

func (s *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsCritical) SetPreCondition(v string) *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsCritical {
	s.PreCondition = &v
	return s
}

func (s *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsCritical) SetStatistics(v string) *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsCritical {
	s.Statistics = &v
	return s
}

func (s *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsCritical) SetThreshold(v string) *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsCritical {
	s.Threshold = &v
	return s
}

func (s *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsCritical) SetTimes(v int32) *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsCritical {
	s.Times = &v
	return s
}

func (s *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsCritical) Validate() error {
	return dara.Validate(s)
}

type CreateGroupMetricRulesRequestGroupMetricRulesEscalationsInfo struct {
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	N                  *string `json:"N,omitempty" xml:"N,omitempty"`
	PreCondition       *string `json:"PreCondition,omitempty" xml:"PreCondition,omitempty"`
	Statistics         *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	Threshold          *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	Times              *int32  `json:"Times,omitempty" xml:"Times,omitempty"`
}

func (s CreateGroupMetricRulesRequestGroupMetricRulesEscalationsInfo) String() string {
	return dara.Prettify(s)
}

func (s CreateGroupMetricRulesRequestGroupMetricRulesEscalationsInfo) GoString() string {
	return s.String()
}

func (s *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsInfo) GetComparisonOperator() *string {
	return s.ComparisonOperator
}

func (s *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsInfo) GetN() *string {
	return s.N
}

func (s *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsInfo) GetPreCondition() *string {
	return s.PreCondition
}

func (s *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsInfo) GetStatistics() *string {
	return s.Statistics
}

func (s *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsInfo) GetThreshold() *string {
	return s.Threshold
}

func (s *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsInfo) GetTimes() *int32 {
	return s.Times
}

func (s *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsInfo) SetComparisonOperator(v string) *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsInfo {
	s.ComparisonOperator = &v
	return s
}

func (s *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsInfo) SetN(v string) *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsInfo {
	s.N = &v
	return s
}

func (s *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsInfo) SetPreCondition(v string) *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsInfo {
	s.PreCondition = &v
	return s
}

func (s *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsInfo) SetStatistics(v string) *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsInfo {
	s.Statistics = &v
	return s
}

func (s *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsInfo) SetThreshold(v string) *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsInfo {
	s.Threshold = &v
	return s
}

func (s *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsInfo) SetTimes(v int32) *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsInfo {
	s.Times = &v
	return s
}

func (s *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsInfo) Validate() error {
	return dara.Validate(s)
}

type CreateGroupMetricRulesRequestGroupMetricRulesEscalationsWarn struct {
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	N                  *string `json:"N,omitempty" xml:"N,omitempty"`
	PreCondition       *string `json:"PreCondition,omitempty" xml:"PreCondition,omitempty"`
	Statistics         *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	Threshold          *string `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	Times              *int32  `json:"Times,omitempty" xml:"Times,omitempty"`
}

func (s CreateGroupMetricRulesRequestGroupMetricRulesEscalationsWarn) String() string {
	return dara.Prettify(s)
}

func (s CreateGroupMetricRulesRequestGroupMetricRulesEscalationsWarn) GoString() string {
	return s.String()
}

func (s *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsWarn) GetComparisonOperator() *string {
	return s.ComparisonOperator
}

func (s *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsWarn) GetN() *string {
	return s.N
}

func (s *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsWarn) GetPreCondition() *string {
	return s.PreCondition
}

func (s *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsWarn) GetStatistics() *string {
	return s.Statistics
}

func (s *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsWarn) GetThreshold() *string {
	return s.Threshold
}

func (s *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsWarn) GetTimes() *int32 {
	return s.Times
}

func (s *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsWarn) SetComparisonOperator(v string) *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsWarn {
	s.ComparisonOperator = &v
	return s
}

func (s *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsWarn) SetN(v string) *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsWarn {
	s.N = &v
	return s
}

func (s *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsWarn) SetPreCondition(v string) *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsWarn {
	s.PreCondition = &v
	return s
}

func (s *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsWarn) SetStatistics(v string) *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsWarn {
	s.Statistics = &v
	return s
}

func (s *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsWarn) SetThreshold(v string) *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsWarn {
	s.Threshold = &v
	return s
}

func (s *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsWarn) SetTimes(v int32) *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsWarn {
	s.Times = &v
	return s
}

func (s *CreateGroupMetricRulesRequestGroupMetricRulesEscalationsWarn) Validate() error {
	return dara.Validate(s)
}

type CreateGroupMetricRulesRequestGroupMetricRulesLabels struct {
	// The tag key of the alert rule. The specified tag is contained in alert notifications.
	//
	// Valid values of N: 1 to 200.
	//
	// example:
	//
	// key1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the alert rule. The specified tag is contained in alert notifications.
	//
	// Valid values of N: 1 to 200.
	//
	// example:
	//
	// value1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateGroupMetricRulesRequestGroupMetricRulesLabels) String() string {
	return dara.Prettify(s)
}

func (s CreateGroupMetricRulesRequestGroupMetricRulesLabels) GoString() string {
	return s.String()
}

func (s *CreateGroupMetricRulesRequestGroupMetricRulesLabels) GetKey() *string {
	return s.Key
}

func (s *CreateGroupMetricRulesRequestGroupMetricRulesLabels) GetValue() *string {
	return s.Value
}

func (s *CreateGroupMetricRulesRequestGroupMetricRulesLabels) SetKey(v string) *CreateGroupMetricRulesRequestGroupMetricRulesLabels {
	s.Key = &v
	return s
}

func (s *CreateGroupMetricRulesRequestGroupMetricRulesLabels) SetValue(v string) *CreateGroupMetricRulesRequestGroupMetricRulesLabels {
	s.Value = &v
	return s
}

func (s *CreateGroupMetricRulesRequestGroupMetricRulesLabels) Validate() error {
	return dara.Validate(s)
}
