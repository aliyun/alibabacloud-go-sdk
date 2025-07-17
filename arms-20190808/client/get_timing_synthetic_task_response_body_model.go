// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTimingSyntheticTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *GetTimingSyntheticTaskResponseBody
	GetCode() *int64
	SetData(v *GetTimingSyntheticTaskResponseBodyData) *GetTimingSyntheticTaskResponseBody
	GetData() *GetTimingSyntheticTaskResponseBodyData
	SetMessage(v string) *GetTimingSyntheticTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetTimingSyntheticTaskResponseBody
	GetRequestId() *string
}

type GetTimingSyntheticTaskResponseBody struct {
	// The status code returned. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned struct.
	Data *GetTimingSyntheticTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// E13430A6-57A9-56E9-9D8D-28FE8DEBCA40
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetTimingSyntheticTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTimingSyntheticTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetTimingSyntheticTaskResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *GetTimingSyntheticTaskResponseBody) GetData() *GetTimingSyntheticTaskResponseBodyData {
	return s.Data
}

func (s *GetTimingSyntheticTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetTimingSyntheticTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTimingSyntheticTaskResponseBody) SetCode(v int64) *GetTimingSyntheticTaskResponseBody {
	s.Code = &v
	return s
}

func (s *GetTimingSyntheticTaskResponseBody) SetData(v *GetTimingSyntheticTaskResponseBodyData) *GetTimingSyntheticTaskResponseBody {
	s.Data = v
	return s
}

func (s *GetTimingSyntheticTaskResponseBody) SetMessage(v string) *GetTimingSyntheticTaskResponseBody {
	s.Message = &v
	return s
}

func (s *GetTimingSyntheticTaskResponseBody) SetRequestId(v string) *GetTimingSyntheticTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTimingSyntheticTaskResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetTimingSyntheticTaskResponseBodyData struct {
	// The list of assertions.
	AvailableAssertions []*GetTimingSyntheticTaskResponseBodyDataAvailableAssertions `json:"AvailableAssertions,omitempty" xml:"AvailableAssertions,omitempty" type:"Repeated"`
	// The general settings.
	CommonSetting *GetTimingSyntheticTaskResponseBodyDataCommonSetting `json:"CommonSetting,omitempty" xml:"CommonSetting,omitempty" type:"Struct"`
	// The custom cycle.
	CustomPeriod *GetTimingSyntheticTaskResponseBodyDataCustomPeriod `json:"CustomPeriod,omitempty" xml:"CustomPeriod,omitempty" type:"Struct"`
	// The detection frequency. Valid values: 1m, 5m, 10m, 15m, 20m, 30m, 1h, 2h, 3h, 4h, 6h, 8h, 12h, and 24h.
	//
	// example:
	//
	// 5m
	Frequency *string `json:"Frequency,omitempty" xml:"Frequency,omitempty"`
	// The detection point type. 1: PC. 2: mobile device.
	//
	// example:
	//
	// 1
	MonitorCategory *int64 `json:"MonitorCategory,omitempty" xml:"MonitorCategory,omitempty"`
	// The monitoring configurations.
	MonitorConf *GetTimingSyntheticTaskResponseBodyDataMonitorConf `json:"MonitorConf,omitempty" xml:"MonitorConf,omitempty" type:"Struct"`
	// The list of monitoring points.
	Monitors []*GetTimingSyntheticTaskResponseBodyDataMonitors `json:"Monitors,omitempty" xml:"Monitors,omitempty" type:"Repeated"`
	// The name of the task.
	//
	// example:
	//
	// AlibabaCloud DNS Task
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// default
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// CREATING: The task is being created. RUNNING: The task is running. PARTIAL_RUNNING: The task is partially running. STOP: The task is stopped. LIMIT_STOP: The task is stopped due to quota insufficiency. EXCEPTION: The task is abnormal. DELETE: The task is deleted. DELETE_EXCEPTION: The task failed to be deleted.
	//
	// example:
	//
	// RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tag.
	Tags []*GetTimingSyntheticTaskResponseBodyDataTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The ID of the synthetic monitoring task.
	//
	// example:
	//
	// 5308a2691f59422c8c3b7aeccec9cd3b
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The type of the task. Valid values:
	//
	// ICMP TCP DNS HTTP Website speed measurement File download
	//
	// example:
	//
	// 5
	TaskType *int32 `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s GetTimingSyntheticTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetTimingSyntheticTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTimingSyntheticTaskResponseBodyData) GetAvailableAssertions() []*GetTimingSyntheticTaskResponseBodyDataAvailableAssertions {
	return s.AvailableAssertions
}

func (s *GetTimingSyntheticTaskResponseBodyData) GetCommonSetting() *GetTimingSyntheticTaskResponseBodyDataCommonSetting {
	return s.CommonSetting
}

func (s *GetTimingSyntheticTaskResponseBodyData) GetCustomPeriod() *GetTimingSyntheticTaskResponseBodyDataCustomPeriod {
	return s.CustomPeriod
}

func (s *GetTimingSyntheticTaskResponseBodyData) GetFrequency() *string {
	return s.Frequency
}

func (s *GetTimingSyntheticTaskResponseBodyData) GetMonitorCategory() *int64 {
	return s.MonitorCategory
}

func (s *GetTimingSyntheticTaskResponseBodyData) GetMonitorConf() *GetTimingSyntheticTaskResponseBodyDataMonitorConf {
	return s.MonitorConf
}

func (s *GetTimingSyntheticTaskResponseBodyData) GetMonitors() []*GetTimingSyntheticTaskResponseBodyDataMonitors {
	return s.Monitors
}

func (s *GetTimingSyntheticTaskResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetTimingSyntheticTaskResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *GetTimingSyntheticTaskResponseBodyData) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetTimingSyntheticTaskResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetTimingSyntheticTaskResponseBodyData) GetTags() []*GetTimingSyntheticTaskResponseBodyDataTags {
	return s.Tags
}

func (s *GetTimingSyntheticTaskResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *GetTimingSyntheticTaskResponseBodyData) GetTaskType() *int32 {
	return s.TaskType
}

func (s *GetTimingSyntheticTaskResponseBodyData) SetAvailableAssertions(v []*GetTimingSyntheticTaskResponseBodyDataAvailableAssertions) *GetTimingSyntheticTaskResponseBodyData {
	s.AvailableAssertions = v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyData) SetCommonSetting(v *GetTimingSyntheticTaskResponseBodyDataCommonSetting) *GetTimingSyntheticTaskResponseBodyData {
	s.CommonSetting = v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyData) SetCustomPeriod(v *GetTimingSyntheticTaskResponseBodyDataCustomPeriod) *GetTimingSyntheticTaskResponseBodyData {
	s.CustomPeriod = v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyData) SetFrequency(v string) *GetTimingSyntheticTaskResponseBodyData {
	s.Frequency = &v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyData) SetMonitorCategory(v int64) *GetTimingSyntheticTaskResponseBodyData {
	s.MonitorCategory = &v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyData) SetMonitorConf(v *GetTimingSyntheticTaskResponseBodyDataMonitorConf) *GetTimingSyntheticTaskResponseBodyData {
	s.MonitorConf = v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyData) SetMonitors(v []*GetTimingSyntheticTaskResponseBodyDataMonitors) *GetTimingSyntheticTaskResponseBodyData {
	s.Monitors = v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyData) SetName(v string) *GetTimingSyntheticTaskResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyData) SetRegionId(v string) *GetTimingSyntheticTaskResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyData) SetResourceGroupId(v string) *GetTimingSyntheticTaskResponseBodyData {
	s.ResourceGroupId = &v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyData) SetStatus(v string) *GetTimingSyntheticTaskResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyData) SetTags(v []*GetTimingSyntheticTaskResponseBodyDataTags) *GetTimingSyntheticTaskResponseBodyData {
	s.Tags = v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyData) SetTaskId(v string) *GetTimingSyntheticTaskResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyData) SetTaskType(v int32) *GetTimingSyntheticTaskResponseBodyData {
	s.TaskType = &v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetTimingSyntheticTaskResponseBodyDataAvailableAssertions struct {
	// The expected value.
	//
	// example:
	//
	// 100
	Expect *string `json:"Expect,omitempty" xml:"Expect,omitempty"`
	// The condition. gt: greater than. gte: greater than or equal to. lt: less than. lte: less than or equal to. eq: equal to. neq: not equal to. ctn: contain. nctn: does not contain. exist: exist. n_exist: does not exist. belong: belong to. n_belong: does not belong to. reg_match: regular expression.
	//
	// example:
	//
	// gt
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// The check target. If you set the type parameter to HttpResCode, HttpResBody, or HttpResponseTime, you do not need to set the target parameter. If you set the type parameter to HttpResHead, you must specify the key in the header. If you set the type parameter to HttpResBodyJson, use jsonPath.
	//
	// example:
	//
	// key
	Target *string `json:"Target,omitempty" xml:"Target,omitempty"`
	// The assertion type. Valid values: HttpResCode, HttpResHead, HttpResBody, HttpResBodyJson, HttpResponseTime, IcmpPackLoss (packet loss rate), IcmpPackMaxLatency (maximum packet latency), IcmpPackAvgLatency (average packet latency), TraceRouteHops (number of hops), DnsARecord (A record), DnsCName (CNAME), websiteTTFB (time to first packet), websiteTTLB (time to last packet), websiteFST (first paint time), websiteFFST (first meaningful paint), websiteOnload (full loaded time). For more information, see the following description.
	//
	// example:
	//
	// websiteTTLB
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetTimingSyntheticTaskResponseBodyDataAvailableAssertions) String() string {
	return dara.Prettify(s)
}

func (s GetTimingSyntheticTaskResponseBodyDataAvailableAssertions) GoString() string {
	return s.String()
}

func (s *GetTimingSyntheticTaskResponseBodyDataAvailableAssertions) GetExpect() *string {
	return s.Expect
}

func (s *GetTimingSyntheticTaskResponseBodyDataAvailableAssertions) GetOperator() *string {
	return s.Operator
}

func (s *GetTimingSyntheticTaskResponseBodyDataAvailableAssertions) GetTarget() *string {
	return s.Target
}

func (s *GetTimingSyntheticTaskResponseBodyDataAvailableAssertions) GetType() *string {
	return s.Type
}

func (s *GetTimingSyntheticTaskResponseBodyDataAvailableAssertions) SetExpect(v string) *GetTimingSyntheticTaskResponseBodyDataAvailableAssertions {
	s.Expect = &v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyDataAvailableAssertions) SetOperator(v string) *GetTimingSyntheticTaskResponseBodyDataAvailableAssertions {
	s.Operator = &v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyDataAvailableAssertions) SetTarget(v string) *GetTimingSyntheticTaskResponseBodyDataAvailableAssertions {
	s.Target = &v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyDataAvailableAssertions) SetType(v string) *GetTimingSyntheticTaskResponseBodyDataAvailableAssertions {
	s.Type = &v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyDataAvailableAssertions) Validate() error {
	return dara.Validate(s)
}

type GetTimingSyntheticTaskResponseBodyDataCommonSetting struct {
	// The custom host.
	CustomHost *GetTimingSyntheticTaskResponseBodyDataCommonSettingCustomHost `json:"CustomHost,omitempty" xml:"CustomHost,omitempty" type:"Struct"`
	// The reserved parameters.
	CustomPrometheusSetting *GetTimingSyntheticTaskResponseBodyDataCommonSettingCustomPrometheusSetting `json:"CustomPrometheusSetting,omitempty" xml:"CustomPrometheusSetting,omitempty" type:"Struct"`
	// User VPC information. If the dial-up is to the Alibaba Cloud intranet address, you need to configure the VPC information.
	CustomVPCSetting *GetTimingSyntheticTaskResponseBodyDataCommonSettingCustomVPCSetting `json:"CustomVPCSetting,omitempty" xml:"CustomVPCSetting,omitempty" type:"Struct"`
	// The IP version. Valid values:
	//
	// 	- 0: A version is automatically selected.
	//
	// 	- 1: IPv4
	//
	// 	- 2: IPv6
	//
	// example:
	//
	// 0
	IpType *int32 `json:"IpType,omitempty" xml:"IpType,omitempty"`
	// Whether to enable tracing.
	//
	// example:
	//
	// true
	IsOpenTrace *bool `json:"IsOpenTrace,omitempty" xml:"IsOpenTrace,omitempty"`
	// Specifies whether to evenly distribute monitoring samples. Valid values:
	//
	// 	- 0: No
	//
	// 	- 1: Yes
	//
	// example:
	//
	// 0
	MonitorSamples *int32 `json:"MonitorSamples,omitempty" xml:"MonitorSamples,omitempty"`
	// Tracing client type:
	//
	// - 0: ARMS Agent
	//
	// - 1: Open Telemetry
	//
	// - 2: Jaeger
	//
	// example:
	//
	// 1
	TraceClientType *int32 `json:"TraceClientType,omitempty" xml:"TraceClientType,omitempty"`
	// Tracing data reporting region.
	//
	// example:
	//
	// cn-hangzhou
	XtraceRegion *string `json:"XtraceRegion,omitempty" xml:"XtraceRegion,omitempty"`
}

func (s GetTimingSyntheticTaskResponseBodyDataCommonSetting) String() string {
	return dara.Prettify(s)
}

func (s GetTimingSyntheticTaskResponseBodyDataCommonSetting) GoString() string {
	return s.String()
}

func (s *GetTimingSyntheticTaskResponseBodyDataCommonSetting) GetCustomHost() *GetTimingSyntheticTaskResponseBodyDataCommonSettingCustomHost {
	return s.CustomHost
}

func (s *GetTimingSyntheticTaskResponseBodyDataCommonSetting) GetCustomPrometheusSetting() *GetTimingSyntheticTaskResponseBodyDataCommonSettingCustomPrometheusSetting {
	return s.CustomPrometheusSetting
}

func (s *GetTimingSyntheticTaskResponseBodyDataCommonSetting) GetCustomVPCSetting() *GetTimingSyntheticTaskResponseBodyDataCommonSettingCustomVPCSetting {
	return s.CustomVPCSetting
}

func (s *GetTimingSyntheticTaskResponseBodyDataCommonSetting) GetIpType() *int32 {
	return s.IpType
}

func (s *GetTimingSyntheticTaskResponseBodyDataCommonSetting) GetIsOpenTrace() *bool {
	return s.IsOpenTrace
}

func (s *GetTimingSyntheticTaskResponseBodyDataCommonSetting) GetMonitorSamples() *int32 {
	return s.MonitorSamples
}

func (s *GetTimingSyntheticTaskResponseBodyDataCommonSetting) GetTraceClientType() *int32 {
	return s.TraceClientType
}

func (s *GetTimingSyntheticTaskResponseBodyDataCommonSetting) GetXtraceRegion() *string {
	return s.XtraceRegion
}

func (s *GetTimingSyntheticTaskResponseBodyDataCommonSetting) SetCustomHost(v *GetTimingSyntheticTaskResponseBodyDataCommonSettingCustomHost) *GetTimingSyntheticTaskResponseBodyDataCommonSetting {
	s.CustomHost = v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyDataCommonSetting) SetCustomPrometheusSetting(v *GetTimingSyntheticTaskResponseBodyDataCommonSettingCustomPrometheusSetting) *GetTimingSyntheticTaskResponseBodyDataCommonSetting {
	s.CustomPrometheusSetting = v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyDataCommonSetting) SetCustomVPCSetting(v *GetTimingSyntheticTaskResponseBodyDataCommonSettingCustomVPCSetting) *GetTimingSyntheticTaskResponseBodyDataCommonSetting {
	s.CustomVPCSetting = v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyDataCommonSetting) SetIpType(v int32) *GetTimingSyntheticTaskResponseBodyDataCommonSetting {
	s.IpType = &v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyDataCommonSetting) SetIsOpenTrace(v bool) *GetTimingSyntheticTaskResponseBodyDataCommonSetting {
	s.IsOpenTrace = &v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyDataCommonSetting) SetMonitorSamples(v int32) *GetTimingSyntheticTaskResponseBodyDataCommonSetting {
	s.MonitorSamples = &v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyDataCommonSetting) SetTraceClientType(v int32) *GetTimingSyntheticTaskResponseBodyDataCommonSetting {
	s.TraceClientType = &v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyDataCommonSetting) SetXtraceRegion(v string) *GetTimingSyntheticTaskResponseBodyDataCommonSetting {
	s.XtraceRegion = &v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyDataCommonSetting) Validate() error {
	return dara.Validate(s)
}

type GetTimingSyntheticTaskResponseBodyDataCommonSettingCustomHost struct {
	// The list of hosts.
	Hosts []*GetTimingSyntheticTaskResponseBodyDataCommonSettingCustomHostHosts `json:"Hosts,omitempty" xml:"Hosts,omitempty" type:"Repeated"`
	// The selection mode. 0: Random. 1: Polling.
	//
	// example:
	//
	// 0
	SelectType *int32 `json:"SelectType,omitempty" xml:"SelectType,omitempty"`
}

func (s GetTimingSyntheticTaskResponseBodyDataCommonSettingCustomHost) String() string {
	return dara.Prettify(s)
}

func (s GetTimingSyntheticTaskResponseBodyDataCommonSettingCustomHost) GoString() string {
	return s.String()
}

func (s *GetTimingSyntheticTaskResponseBodyDataCommonSettingCustomHost) GetHosts() []*GetTimingSyntheticTaskResponseBodyDataCommonSettingCustomHostHosts {
	return s.Hosts
}

func (s *GetTimingSyntheticTaskResponseBodyDataCommonSettingCustomHost) GetSelectType() *int32 {
	return s.SelectType
}

func (s *GetTimingSyntheticTaskResponseBodyDataCommonSettingCustomHost) SetHosts(v []*GetTimingSyntheticTaskResponseBodyDataCommonSettingCustomHostHosts) *GetTimingSyntheticTaskResponseBodyDataCommonSettingCustomHost {
	s.Hosts = v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyDataCommonSettingCustomHost) SetSelectType(v int32) *GetTimingSyntheticTaskResponseBodyDataCommonSettingCustomHost {
	s.SelectType = &v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyDataCommonSettingCustomHost) Validate() error {
	return dara.Validate(s)
}

type GetTimingSyntheticTaskResponseBodyDataCommonSettingCustomHostHosts struct {
	// The domain name.
	//
	// example:
	//
	// www.aliyun.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The IP version. Valid values:
	//
	// 	- 0: A version is automatically selected.
	//
	// 	- 1: IPv4
	//
	// 	- 2: IPv6
	//
	// example:
	//
	// 0
	IpType *int32 `json:"IpType,omitempty" xml:"IpType,omitempty"`
	// The list of IP addresses.
	Ips []*string `json:"Ips,omitempty" xml:"Ips,omitempty" type:"Repeated"`
}

func (s GetTimingSyntheticTaskResponseBodyDataCommonSettingCustomHostHosts) String() string {
	return dara.Prettify(s)
}

func (s GetTimingSyntheticTaskResponseBodyDataCommonSettingCustomHostHosts) GoString() string {
	return s.String()
}

func (s *GetTimingSyntheticTaskResponseBodyDataCommonSettingCustomHostHosts) GetDomain() *string {
	return s.Domain
}

func (s *GetTimingSyntheticTaskResponseBodyDataCommonSettingCustomHostHosts) GetIpType() *int32 {
	return s.IpType
}

func (s *GetTimingSyntheticTaskResponseBodyDataCommonSettingCustomHostHosts) GetIps() []*string {
	return s.Ips
}

func (s *GetTimingSyntheticTaskResponseBodyDataCommonSettingCustomHostHosts) SetDomain(v string) *GetTimingSyntheticTaskResponseBodyDataCommonSettingCustomHostHosts {
	s.Domain = &v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyDataCommonSettingCustomHostHosts) SetIpType(v int32) *GetTimingSyntheticTaskResponseBodyDataCommonSettingCustomHostHosts {
	s.IpType = &v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyDataCommonSettingCustomHostHosts) SetIps(v []*string) *GetTimingSyntheticTaskResponseBodyDataCommonSettingCustomHostHosts {
	s.Ips = v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyDataCommonSettingCustomHostHosts) Validate() error {
	return dara.Validate(s)
}

type GetTimingSyntheticTaskResponseBodyDataCommonSettingCustomPrometheusSetting struct {
	// A reserved parameter.
	//
	// example:
	//
	// Reserved parameter
	PrometheusClusterId *string `json:"PrometheusClusterId,omitempty" xml:"PrometheusClusterId,omitempty"`
	// A reserved parameter.
	//
	// example:
	//
	// Reserved parameter
	PrometheusClusterRegion *string `json:"PrometheusClusterRegion,omitempty" xml:"PrometheusClusterRegion,omitempty"`
	// A reserved parameter.
	PrometheusLabels map[string]*string `json:"PrometheusLabels,omitempty" xml:"PrometheusLabels,omitempty"`
}

func (s GetTimingSyntheticTaskResponseBodyDataCommonSettingCustomPrometheusSetting) String() string {
	return dara.Prettify(s)
}

func (s GetTimingSyntheticTaskResponseBodyDataCommonSettingCustomPrometheusSetting) GoString() string {
	return s.String()
}

func (s *GetTimingSyntheticTaskResponseBodyDataCommonSettingCustomPrometheusSetting) GetPrometheusClusterId() *string {
	return s.PrometheusClusterId
}

func (s *GetTimingSyntheticTaskResponseBodyDataCommonSettingCustomPrometheusSetting) GetPrometheusClusterRegion() *string {
	return s.PrometheusClusterRegion
}

func (s *GetTimingSyntheticTaskResponseBodyDataCommonSettingCustomPrometheusSetting) GetPrometheusLabels() map[string]*string {
	return s.PrometheusLabels
}

func (s *GetTimingSyntheticTaskResponseBodyDataCommonSettingCustomPrometheusSetting) SetPrometheusClusterId(v string) *GetTimingSyntheticTaskResponseBodyDataCommonSettingCustomPrometheusSetting {
	s.PrometheusClusterId = &v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyDataCommonSettingCustomPrometheusSetting) SetPrometheusClusterRegion(v string) *GetTimingSyntheticTaskResponseBodyDataCommonSettingCustomPrometheusSetting {
	s.PrometheusClusterRegion = &v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyDataCommonSettingCustomPrometheusSetting) SetPrometheusLabels(v map[string]*string) *GetTimingSyntheticTaskResponseBodyDataCommonSettingCustomPrometheusSetting {
	s.PrometheusLabels = v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyDataCommonSettingCustomPrometheusSetting) Validate() error {
	return dara.Validate(s)
}

type GetTimingSyntheticTaskResponseBodyDataCommonSettingCustomVPCSetting struct {
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Security group ID. This security group is where the dial-up client is located. The security group limits the inbound and outbound rules of the dial-up client in the VPC. You need to set the inbound rules of the security group where your VPC is located to allow the security group where the dial-up client is located to access. Otherwise, the dial-up client cannot smoothly access the resources in your VPC.
	//
	// example:
	//
	// sg-xxxxxxx
	SecureGroupId *string `json:"SecureGroupId,omitempty" xml:"SecureGroupId,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-bp1i0xezblf1yrz4xxxxx
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// VPC ID.
	//
	// example:
	//
	// vpc-2zexy5nae9q2otaxxxx
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s GetTimingSyntheticTaskResponseBodyDataCommonSettingCustomVPCSetting) String() string {
	return dara.Prettify(s)
}

func (s GetTimingSyntheticTaskResponseBodyDataCommonSettingCustomVPCSetting) GoString() string {
	return s.String()
}

func (s *GetTimingSyntheticTaskResponseBodyDataCommonSettingCustomVPCSetting) GetRegionId() *string {
	return s.RegionId
}

func (s *GetTimingSyntheticTaskResponseBodyDataCommonSettingCustomVPCSetting) GetSecureGroupId() *string {
	return s.SecureGroupId
}

func (s *GetTimingSyntheticTaskResponseBodyDataCommonSettingCustomVPCSetting) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *GetTimingSyntheticTaskResponseBodyDataCommonSettingCustomVPCSetting) GetVpcId() *string {
	return s.VpcId
}

func (s *GetTimingSyntheticTaskResponseBodyDataCommonSettingCustomVPCSetting) SetRegionId(v string) *GetTimingSyntheticTaskResponseBodyDataCommonSettingCustomVPCSetting {
	s.RegionId = &v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyDataCommonSettingCustomVPCSetting) SetSecureGroupId(v string) *GetTimingSyntheticTaskResponseBodyDataCommonSettingCustomVPCSetting {
	s.SecureGroupId = &v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyDataCommonSettingCustomVPCSetting) SetVSwitchId(v string) *GetTimingSyntheticTaskResponseBodyDataCommonSettingCustomVPCSetting {
	s.VSwitchId = &v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyDataCommonSettingCustomVPCSetting) SetVpcId(v string) *GetTimingSyntheticTaskResponseBodyDataCommonSettingCustomVPCSetting {
	s.VpcId = &v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyDataCommonSettingCustomVPCSetting) Validate() error {
	return dara.Validate(s)
}

type GetTimingSyntheticTaskResponseBodyDataCustomPeriod struct {
	// The hour at which the test ends. Valid values: 0 to 24.
	//
	// example:
	//
	// 22
	EndHour *int64 `json:"EndHour,omitempty" xml:"EndHour,omitempty"`
	// The hour at which the test starts. Valid values: 0 to 24.
	//
	// example:
	//
	// 14
	StartHour *int64 `json:"StartHour,omitempty" xml:"StartHour,omitempty"`
}

func (s GetTimingSyntheticTaskResponseBodyDataCustomPeriod) String() string {
	return dara.Prettify(s)
}

func (s GetTimingSyntheticTaskResponseBodyDataCustomPeriod) GoString() string {
	return s.String()
}

func (s *GetTimingSyntheticTaskResponseBodyDataCustomPeriod) GetEndHour() *int64 {
	return s.EndHour
}

func (s *GetTimingSyntheticTaskResponseBodyDataCustomPeriod) GetStartHour() *int64 {
	return s.StartHour
}

func (s *GetTimingSyntheticTaskResponseBodyDataCustomPeriod) SetEndHour(v int64) *GetTimingSyntheticTaskResponseBodyDataCustomPeriod {
	s.EndHour = &v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyDataCustomPeriod) SetStartHour(v int64) *GetTimingSyntheticTaskResponseBodyDataCustomPeriod {
	s.StartHour = &v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyDataCustomPeriod) Validate() error {
	return dara.Validate(s)
}

type GetTimingSyntheticTaskResponseBodyDataMonitorConf struct {
	// The parameters of the HTTP(S) synthetic test.
	ApiHTTP *GetTimingSyntheticTaskResponseBodyDataMonitorConfApiHTTP `json:"ApiHTTP,omitempty" xml:"ApiHTTP,omitempty" type:"Struct"`
	// The file download parameters.
	FileDownload *GetTimingSyntheticTaskResponseBodyDataMonitorConfFileDownload `json:"FileDownload,omitempty" xml:"FileDownload,omitempty" type:"Struct"`
	// The DNS synthetic test parameters. This parameter is required if the TaskType parameter is set to 3.
	NetDNS *GetTimingSyntheticTaskResponseBodyDataMonitorConfNetDNS `json:"NetDNS,omitempty" xml:"NetDNS,omitempty" type:"Struct"`
	// The ICMP synthetic test parameters. This parameter is required if the TaskType parameter is set to 1.
	NetICMP *GetTimingSyntheticTaskResponseBodyDataMonitorConfNetICMP `json:"NetICMP,omitempty" xml:"NetICMP,omitempty" type:"Struct"`
	// The TCP synthetic tests parameters. This parameter is required if the TaskType parameter is set to 2.
	NetTCP *GetTimingSyntheticTaskResponseBodyDataMonitorConfNetTCP `json:"NetTCP,omitempty" xml:"NetTCP,omitempty" type:"Struct"`
	// Streaming media dial test configuration.
	Stream *GetTimingSyntheticTaskResponseBodyDataMonitorConfStream `json:"Stream,omitempty" xml:"Stream,omitempty" type:"Struct"`
	// The website-speed measurement parameters.
	Website *GetTimingSyntheticTaskResponseBodyDataMonitorConfWebsite `json:"Website,omitempty" xml:"Website,omitempty" type:"Struct"`
}

func (s GetTimingSyntheticTaskResponseBodyDataMonitorConf) String() string {
	return dara.Prettify(s)
}

func (s GetTimingSyntheticTaskResponseBodyDataMonitorConf) GoString() string {
	return s.String()
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConf) GetApiHTTP() *GetTimingSyntheticTaskResponseBodyDataMonitorConfApiHTTP {
	return s.ApiHTTP
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConf) GetFileDownload() *GetTimingSyntheticTaskResponseBodyDataMonitorConfFileDownload {
	return s.FileDownload
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConf) GetNetDNS() *GetTimingSyntheticTaskResponseBodyDataMonitorConfNetDNS {
	return s.NetDNS
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConf) GetNetICMP() *GetTimingSyntheticTaskResponseBodyDataMonitorConfNetICMP {
	return s.NetICMP
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConf) GetNetTCP() *GetTimingSyntheticTaskResponseBodyDataMonitorConfNetTCP {
	return s.NetTCP
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConf) GetStream() *GetTimingSyntheticTaskResponseBodyDataMonitorConfStream {
	return s.Stream
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConf) GetWebsite() *GetTimingSyntheticTaskResponseBodyDataMonitorConfWebsite {
	return s.Website
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConf) SetApiHTTP(v *GetTimingSyntheticTaskResponseBodyDataMonitorConfApiHTTP) *GetTimingSyntheticTaskResponseBodyDataMonitorConf {
	s.ApiHTTP = v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConf) SetFileDownload(v *GetTimingSyntheticTaskResponseBodyDataMonitorConfFileDownload) *GetTimingSyntheticTaskResponseBodyDataMonitorConf {
	s.FileDownload = v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConf) SetNetDNS(v *GetTimingSyntheticTaskResponseBodyDataMonitorConfNetDNS) *GetTimingSyntheticTaskResponseBodyDataMonitorConf {
	s.NetDNS = v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConf) SetNetICMP(v *GetTimingSyntheticTaskResponseBodyDataMonitorConfNetICMP) *GetTimingSyntheticTaskResponseBodyDataMonitorConf {
	s.NetICMP = v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConf) SetNetTCP(v *GetTimingSyntheticTaskResponseBodyDataMonitorConfNetTCP) *GetTimingSyntheticTaskResponseBodyDataMonitorConf {
	s.NetTCP = v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConf) SetStream(v *GetTimingSyntheticTaskResponseBodyDataMonitorConfStream) *GetTimingSyntheticTaskResponseBodyDataMonitorConf {
	s.Stream = v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConf) SetWebsite(v *GetTimingSyntheticTaskResponseBodyDataMonitorConfWebsite) *GetTimingSyntheticTaskResponseBodyDataMonitorConf {
	s.Website = v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConf) Validate() error {
	return dara.Validate(s)
}

type GetTimingSyntheticTaskResponseBodyDataMonitorConfApiHTTP struct {
	// Whether to verify the certificate. The default is no.
	//
	// example:
	//
	// false
	CheckCert *bool `json:"CheckCert,omitempty" xml:"CheckCert,omitempty"`
	// The connection timeout period. Unit: milliseconds. Default value: 5000. Minimum value: 1000. Maximum value: 300000.
	//
	// example:
	//
	// 5000
	ConnectTimeout *int64 `json:"ConnectTimeout,omitempty" xml:"ConnectTimeout,omitempty"`
	// The request method.
	//
	// 	- POST
	//
	// 	- GET
	//
	// example:
	//
	// POST
	Method *string `json:"Method,omitempty" xml:"Method,omitempty"`
	// The ALPN protocol version. You can configure this parameter when you perform an HTTPS synthetic test on a WAP mobile client. Valid values:
	//
	// 0: default
	//
	// 1: HTTP/1.1
	//
	// 2: HTTP/2
	//
	// 3: disables the ALPN protocol
	//
	// example:
	//
	// 1
	ProtocolAlpnProtocol *int32 `json:"ProtocolAlpnProtocol,omitempty" xml:"ProtocolAlpnProtocol,omitempty"`
	// The HTTP request body.
	RequestBody *GetTimingSyntheticTaskResponseBodyDataMonitorConfApiHTTPRequestBody `json:"RequestBody,omitempty" xml:"RequestBody,omitempty" type:"Struct"`
	// The HTTP request header.
	RequestHeaders map[string]*string `json:"RequestHeaders,omitempty" xml:"RequestHeaders,omitempty"`
	// The URL for synthetic monitoring.
	//
	// example:
	//
	// http://127.0.0.1:8090/api/list
	TargetUrl *string `json:"TargetUrl,omitempty" xml:"TargetUrl,omitempty"`
	// The timeout period. Unit: milliseconds. Default value: 10000. Minimum value: 1000. Maximum value: 300000.
	//
	// example:
	//
	// 10000
	Timeout *int64 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s GetTimingSyntheticTaskResponseBodyDataMonitorConfApiHTTP) String() string {
	return dara.Prettify(s)
}

func (s GetTimingSyntheticTaskResponseBodyDataMonitorConfApiHTTP) GoString() string {
	return s.String()
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfApiHTTP) GetCheckCert() *bool {
	return s.CheckCert
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfApiHTTP) GetConnectTimeout() *int64 {
	return s.ConnectTimeout
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfApiHTTP) GetMethod() *string {
	return s.Method
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfApiHTTP) GetProtocolAlpnProtocol() *int32 {
	return s.ProtocolAlpnProtocol
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfApiHTTP) GetRequestBody() *GetTimingSyntheticTaskResponseBodyDataMonitorConfApiHTTPRequestBody {
	return s.RequestBody
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfApiHTTP) GetRequestHeaders() map[string]*string {
	return s.RequestHeaders
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfApiHTTP) GetTargetUrl() *string {
	return s.TargetUrl
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfApiHTTP) GetTimeout() *int64 {
	return s.Timeout
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfApiHTTP) SetCheckCert(v bool) *GetTimingSyntheticTaskResponseBodyDataMonitorConfApiHTTP {
	s.CheckCert = &v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfApiHTTP) SetConnectTimeout(v int64) *GetTimingSyntheticTaskResponseBodyDataMonitorConfApiHTTP {
	s.ConnectTimeout = &v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfApiHTTP) SetMethod(v string) *GetTimingSyntheticTaskResponseBodyDataMonitorConfApiHTTP {
	s.Method = &v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfApiHTTP) SetProtocolAlpnProtocol(v int32) *GetTimingSyntheticTaskResponseBodyDataMonitorConfApiHTTP {
	s.ProtocolAlpnProtocol = &v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfApiHTTP) SetRequestBody(v *GetTimingSyntheticTaskResponseBodyDataMonitorConfApiHTTPRequestBody) *GetTimingSyntheticTaskResponseBodyDataMonitorConfApiHTTP {
	s.RequestBody = v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfApiHTTP) SetRequestHeaders(v map[string]*string) *GetTimingSyntheticTaskResponseBodyDataMonitorConfApiHTTP {
	s.RequestHeaders = v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfApiHTTP) SetTargetUrl(v string) *GetTimingSyntheticTaskResponseBodyDataMonitorConfApiHTTP {
	s.TargetUrl = &v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfApiHTTP) SetTimeout(v int64) *GetTimingSyntheticTaskResponseBodyDataMonitorConfApiHTTP {
	s.Timeout = &v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfApiHTTP) Validate() error {
	return dara.Validate(s)
}

type GetTimingSyntheticTaskResponseBodyDataMonitorConfApiHTTPRequestBody struct {
	// The content of the request body. Format: JSON string. The parameter is required if the type parameter is set to text/plain, application/json, application/xml, or text/html. Format: JSON string.
	//
	// example:
	//
	// text/plain
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The type of the request body. Valid values: text/plain, application/json, application/x-www-form-urlencoded, multipart/form-data, application/xml, and text/html.
	//
	// example:
	//
	// multipart/form-data
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetTimingSyntheticTaskResponseBodyDataMonitorConfApiHTTPRequestBody) String() string {
	return dara.Prettify(s)
}

func (s GetTimingSyntheticTaskResponseBodyDataMonitorConfApiHTTPRequestBody) GoString() string {
	return s.String()
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfApiHTTPRequestBody) GetContent() *string {
	return s.Content
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfApiHTTPRequestBody) GetType() *string {
	return s.Type
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfApiHTTPRequestBody) SetContent(v string) *GetTimingSyntheticTaskResponseBodyDataMonitorConfApiHTTPRequestBody {
	s.Content = &v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfApiHTTPRequestBody) SetType(v string) *GetTimingSyntheticTaskResponseBodyDataMonitorConfApiHTTPRequestBody {
	s.Type = &v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfApiHTTPRequestBody) Validate() error {
	return dara.Validate(s)
}

type GetTimingSyntheticTaskResponseBodyDataMonitorConfFileDownload struct {
	// The connection timeout period. Unit: milliseconds. Minimum value: 1000. Maximum value: 120000. Default value: 5000.
	//
	// example:
	//
	// 5000
	ConnectionTimeout *int64 `json:"ConnectionTimeout,omitempty" xml:"ConnectionTimeout,omitempty"`
	// The content of the custom request header.
	CustomHeaderContent map[string]*string `json:"CustomHeaderContent,omitempty" xml:"CustomHeaderContent,omitempty"`
	// The kernel type.
	//
	// 	- 1: curl
	//
	// 	- 0: WinInet
	//
	// example:
	//
	// 0
	DownloadKernel *int64 `json:"DownloadKernel,omitempty" xml:"DownloadKernel,omitempty"`
	// Specifies whether to ignore CA certificate authentication errors. 0: No. 1: Yes. Default value: 1.
	//
	// example:
	//
	// 0
	IgnoreCertificateAuthError *int32 `json:"IgnoreCertificateAuthError,omitempty" xml:"IgnoreCertificateAuthError,omitempty"`
	// Specifies whether to ignore certificate revocation errors. 0: No. 1: Yes. Default value: 1.
	//
	// example:
	//
	// 0
	IgnoreCertificateCanceledError *int32 `json:"IgnoreCertificateCanceledError,omitempty" xml:"IgnoreCertificateCanceledError,omitempty"`
	// Specifies whether to ignore certificate invalidity. 0: No. 1: Yes. Default value: 1.
	//
	// example:
	//
	// 0
	IgnoreCertificateOutOfDateError *int32 `json:"IgnoreCertificateOutOfDateError,omitempty" xml:"IgnoreCertificateOutOfDateError,omitempty"`
	// Specifies whether to ignore certificate status errors. 0: No. 1: Yes. Default value: 1.
	//
	// example:
	//
	// 0
	IgnoreCertificateStatusError *int32 `json:"IgnoreCertificateStatusError,omitempty" xml:"IgnoreCertificateStatusError,omitempty"`
	// Specifies whether to ignore certificate incredibility. 0: No. 1: Yes. Default value: 1.
	//
	// example:
	//
	// 0
	IgnoreCertificateUntrustworthyError *int32 `json:"IgnoreCertificateUntrustworthyError,omitempty" xml:"IgnoreCertificateUntrustworthyError,omitempty"`
	// Specifies whether to ignore certificate usage errors. 0: No. 1: Yes. Default value: 1.
	//
	// example:
	//
	// 0
	IgnoreCertificateUsingError *int32 `json:"IgnoreCertificateUsingError,omitempty" xml:"IgnoreCertificateUsingError,omitempty"`
	// Specifies whether to ignore host invalidity. 0: No. 1: Yes. Default value: 1.
	//
	// example:
	//
	// 0
	IgnoreInvalidHostError *int32 `json:"IgnoreInvalidHostError,omitempty" xml:"IgnoreInvalidHostError,omitempty"`
	// The monitoring timeout period. Unit: milliseconds. Minimum value: 1000. Maximum value: 120000. Default value: 60000.
	//
	// example:
	//
	// 6000
	MonitorTimeout *int64 `json:"MonitorTimeout,omitempty" xml:"MonitorTimeout,omitempty"`
	// The QUIC protocol type.
	//
	// 	- 1: http1
	//
	// 	- 2: http2
	//
	// 	- 3: http3
	//
	// example:
	//
	// 1
	QuickProtocol *int64 `json:"QuickProtocol,omitempty" xml:"QuickProtocol,omitempty"`
	// Specifies whether to support redirection. 0: No. 1: Yes. Default value: 1.
	//
	// example:
	//
	// 0
	Redirection *int32 `json:"Redirection,omitempty" xml:"Redirection,omitempty"`
	// The file download URL.
	//
	// example:
	//
	// https://********
	TargetUrl *string `json:"TargetUrl,omitempty" xml:"TargetUrl,omitempty"`
	// The maximum file size of a single transfer. Unit: KB. Minimum value: 1. Maximum value: 20480. Valid values: 2048.
	//
	// example:
	//
	// 2048
	TransmissionSize *int64 `json:"TransmissionSize,omitempty" xml:"TransmissionSize,omitempty"`
	// Verify keywords.
	//
	// example:
	//
	// success
	ValidateKeywords *string `json:"ValidateKeywords,omitempty" xml:"ValidateKeywords,omitempty"`
	// Verification method.
	//
	// - 0: No verification
	//
	// - 1: Verification string
	//
	// - 2: MD5 verification
	//
	// example:
	//
	// 0
	VerifyWay *int32 `json:"VerifyWay,omitempty" xml:"VerifyWay,omitempty"`
	// DNS hijacking whitelist. Matching rules support IP, IP wildcard, subnet mask and CNAME. You can fill in multiple matching rules, and multiple matching rules are separated by vertical bars (|). For example: `www.aliyun.com:203.0.3.55|203.3.44.67`, which means that all IPs except 203.0.3.55 and 203.3.44.67 under the www.aliyun.com domain name are hijacked.
	//
	// example:
	//
	// www.aliyun.com:203.0.3.55|203.3.44.67
	WhiteList *string `json:"WhiteList,omitempty" xml:"WhiteList,omitempty"`
}

func (s GetTimingSyntheticTaskResponseBodyDataMonitorConfFileDownload) String() string {
	return dara.Prettify(s)
}

func (s GetTimingSyntheticTaskResponseBodyDataMonitorConfFileDownload) GoString() string {
	return s.String()
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfFileDownload) GetConnectionTimeout() *int64 {
	return s.ConnectionTimeout
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfFileDownload) GetCustomHeaderContent() map[string]*string {
	return s.CustomHeaderContent
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfFileDownload) GetDownloadKernel() *int64 {
	return s.DownloadKernel
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfFileDownload) GetIgnoreCertificateAuthError() *int32 {
	return s.IgnoreCertificateAuthError
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfFileDownload) GetIgnoreCertificateCanceledError() *int32 {
	return s.IgnoreCertificateCanceledError
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfFileDownload) GetIgnoreCertificateOutOfDateError() *int32 {
	return s.IgnoreCertificateOutOfDateError
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfFileDownload) GetIgnoreCertificateStatusError() *int32 {
	return s.IgnoreCertificateStatusError
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfFileDownload) GetIgnoreCertificateUntrustworthyError() *int32 {
	return s.IgnoreCertificateUntrustworthyError
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfFileDownload) GetIgnoreCertificateUsingError() *int32 {
	return s.IgnoreCertificateUsingError
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfFileDownload) GetIgnoreInvalidHostError() *int32 {
	return s.IgnoreInvalidHostError
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfFileDownload) GetMonitorTimeout() *int64 {
	return s.MonitorTimeout
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfFileDownload) GetQuickProtocol() *int64 {
	return s.QuickProtocol
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfFileDownload) GetRedirection() *int32 {
	return s.Redirection
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfFileDownload) GetTargetUrl() *string {
	return s.TargetUrl
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfFileDownload) GetTransmissionSize() *int64 {
	return s.TransmissionSize
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfFileDownload) GetValidateKeywords() *string {
	return s.ValidateKeywords
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfFileDownload) GetVerifyWay() *int32 {
	return s.VerifyWay
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfFileDownload) GetWhiteList() *string {
	return s.WhiteList
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfFileDownload) SetConnectionTimeout(v int64) *GetTimingSyntheticTaskResponseBodyDataMonitorConfFileDownload {
	s.ConnectionTimeout = &v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfFileDownload) SetCustomHeaderContent(v map[string]*string) *GetTimingSyntheticTaskResponseBodyDataMonitorConfFileDownload {
	s.CustomHeaderContent = v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfFileDownload) SetDownloadKernel(v int64) *GetTimingSyntheticTaskResponseBodyDataMonitorConfFileDownload {
	s.DownloadKernel = &v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfFileDownload) SetIgnoreCertificateAuthError(v int32) *GetTimingSyntheticTaskResponseBodyDataMonitorConfFileDownload {
	s.IgnoreCertificateAuthError = &v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfFileDownload) SetIgnoreCertificateCanceledError(v int32) *GetTimingSyntheticTaskResponseBodyDataMonitorConfFileDownload {
	s.IgnoreCertificateCanceledError = &v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfFileDownload) SetIgnoreCertificateOutOfDateError(v int32) *GetTimingSyntheticTaskResponseBodyDataMonitorConfFileDownload {
	s.IgnoreCertificateOutOfDateError = &v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfFileDownload) SetIgnoreCertificateStatusError(v int32) *GetTimingSyntheticTaskResponseBodyDataMonitorConfFileDownload {
	s.IgnoreCertificateStatusError = &v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfFileDownload) SetIgnoreCertificateUntrustworthyError(v int32) *GetTimingSyntheticTaskResponseBodyDataMonitorConfFileDownload {
	s.IgnoreCertificateUntrustworthyError = &v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfFileDownload) SetIgnoreCertificateUsingError(v int32) *GetTimingSyntheticTaskResponseBodyDataMonitorConfFileDownload {
	s.IgnoreCertificateUsingError = &v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfFileDownload) SetIgnoreInvalidHostError(v int32) *GetTimingSyntheticTaskResponseBodyDataMonitorConfFileDownload {
	s.IgnoreInvalidHostError = &v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfFileDownload) SetMonitorTimeout(v int64) *GetTimingSyntheticTaskResponseBodyDataMonitorConfFileDownload {
	s.MonitorTimeout = &v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfFileDownload) SetQuickProtocol(v int64) *GetTimingSyntheticTaskResponseBodyDataMonitorConfFileDownload {
	s.QuickProtocol = &v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfFileDownload) SetRedirection(v int32) *GetTimingSyntheticTaskResponseBodyDataMonitorConfFileDownload {
	s.Redirection = &v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfFileDownload) SetTargetUrl(v string) *GetTimingSyntheticTaskResponseBodyDataMonitorConfFileDownload {
	s.TargetUrl = &v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfFileDownload) SetTransmissionSize(v int64) *GetTimingSyntheticTaskResponseBodyDataMonitorConfFileDownload {
	s.TransmissionSize = &v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfFileDownload) SetValidateKeywords(v string) *GetTimingSyntheticTaskResponseBodyDataMonitorConfFileDownload {
	s.ValidateKeywords = &v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfFileDownload) SetVerifyWay(v int32) *GetTimingSyntheticTaskResponseBodyDataMonitorConfFileDownload {
	s.VerifyWay = &v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfFileDownload) SetWhiteList(v string) *GetTimingSyntheticTaskResponseBodyDataMonitorConfFileDownload {
	s.WhiteList = &v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfFileDownload) Validate() error {
	return dara.Validate(s)
}

type GetTimingSyntheticTaskResponseBodyDataMonitorConfNetDNS struct {
	// The IP version of the DNS server. 0: IPv4. 1: IPv6. 2: A version is automatically selected. Default value: 0.
	//
	// example:
	//
	// 0
	DnsServerIpType *int32 `json:"DnsServerIpType,omitempty" xml:"DnsServerIpType,omitempty"`
	// The IP address of the DNS server. Default value: 114.114.114.114.
	//
	// example:
	//
	// 114.114.114.114
	NsServer *string `json:"NsServer,omitempty" xml:"NsServer,omitempty"`
	// The DNS query. 0: recursive, 1: iterative. Default value: 0.
	//
	// example:
	//
	// 0
	QueryMethod *int32 `json:"QueryMethod,omitempty" xml:"QueryMethod,omitempty"`
	// The destination domain name.
	//
	// example:
	//
	// www.aliyun.com
	TargetUrl *string `json:"TargetUrl,omitempty" xml:"TargetUrl,omitempty"`
	// The timeout period for the DNS synthetic test. Unit: milliseconds. The minimum value is 1000 and the maximum value is 45000. Default value: 5000.
	//
	// example:
	//
	// 5000
	Timeout *int64 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s GetTimingSyntheticTaskResponseBodyDataMonitorConfNetDNS) String() string {
	return dara.Prettify(s)
}

func (s GetTimingSyntheticTaskResponseBodyDataMonitorConfNetDNS) GoString() string {
	return s.String()
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfNetDNS) GetDnsServerIpType() *int32 {
	return s.DnsServerIpType
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfNetDNS) GetNsServer() *string {
	return s.NsServer
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfNetDNS) GetQueryMethod() *int32 {
	return s.QueryMethod
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfNetDNS) GetTargetUrl() *string {
	return s.TargetUrl
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfNetDNS) GetTimeout() *int64 {
	return s.Timeout
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfNetDNS) SetDnsServerIpType(v int32) *GetTimingSyntheticTaskResponseBodyDataMonitorConfNetDNS {
	s.DnsServerIpType = &v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfNetDNS) SetNsServer(v string) *GetTimingSyntheticTaskResponseBodyDataMonitorConfNetDNS {
	s.NsServer = &v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfNetDNS) SetQueryMethod(v int32) *GetTimingSyntheticTaskResponseBodyDataMonitorConfNetDNS {
	s.QueryMethod = &v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfNetDNS) SetTargetUrl(v string) *GetTimingSyntheticTaskResponseBodyDataMonitorConfNetDNS {
	s.TargetUrl = &v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfNetDNS) SetTimeout(v int64) *GetTimingSyntheticTaskResponseBodyDataMonitorConfNetDNS {
	s.Timeout = &v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfNetDNS) Validate() error {
	return dara.Validate(s)
}

type GetTimingSyntheticTaskResponseBodyDataMonitorConfNetICMP struct {
	// The interval at which ICMP packets are sent. Unit: milliseconds. Minimum value: 200. Maximum value: 2000. Default value: 200.
	//
	// example:
	//
	// 200
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The number of ICMP packets that are sent. Minimum value: 1. Maximum value: 50. Default value: 4.
	//
	// example:
	//
	// 4
	PackageNum *int32 `json:"PackageNum,omitempty" xml:"PackageNum,omitempty"`
	// The size of each ICMP packet. Unit: bytes. Valid values: 32, 64, 128, 256, 512, 1024.
	//
	// example:
	//
	// 1024
	PackageSize *int32 `json:"PackageSize,omitempty" xml:"PackageSize,omitempty"`
	// Specifies whether to split ICMP packets. Default value: true.
	//
	// example:
	//
	// true
	SplitPackage *bool `json:"SplitPackage,omitempty" xml:"SplitPackage,omitempty"`
	// The destination host IP address or domain name
	//
	// example:
	//
	// www.aliyun.com
	TargetUrl *string `json:"TargetUrl,omitempty" xml:"TargetUrl,omitempty"`
	// The timeout period for the TCP synthetic test. Unit: milliseconds. Minimum value: 1000. Maximum value: 300000. Default value: 20000.
	//
	// example:
	//
	// 2000
	Timeout *int64 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// Specifies whether to enable the tracert command. Default value: true.
	//
	// example:
	//
	// true
	TracertEnable *bool `json:"TracertEnable,omitempty" xml:"TracertEnable,omitempty"`
	// The maximum number of hops for tracert. Minimum value: 1. Maximum value: 128. Default value: 64.
	//
	// example:
	//
	// 64
	TracertNumMax *int32 `json:"TracertNumMax,omitempty" xml:"TracertNumMax,omitempty"`
	// The timeout period of tracert. Unit: milliseconds. Minimum value: 1000. Maximum value: 300000. Default value: 60000.
	//
	// example:
	//
	// 60000
	TracertTimeout *int64 `json:"TracertTimeout,omitempty" xml:"TracertTimeout,omitempty"`
}

func (s GetTimingSyntheticTaskResponseBodyDataMonitorConfNetICMP) String() string {
	return dara.Prettify(s)
}

func (s GetTimingSyntheticTaskResponseBodyDataMonitorConfNetICMP) GoString() string {
	return s.String()
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfNetICMP) GetInterval() *int32 {
	return s.Interval
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfNetICMP) GetPackageNum() *int32 {
	return s.PackageNum
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfNetICMP) GetPackageSize() *int32 {
	return s.PackageSize
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfNetICMP) GetSplitPackage() *bool {
	return s.SplitPackage
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfNetICMP) GetTargetUrl() *string {
	return s.TargetUrl
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfNetICMP) GetTimeout() *int64 {
	return s.Timeout
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfNetICMP) GetTracertEnable() *bool {
	return s.TracertEnable
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfNetICMP) GetTracertNumMax() *int32 {
	return s.TracertNumMax
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfNetICMP) GetTracertTimeout() *int64 {
	return s.TracertTimeout
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfNetICMP) SetInterval(v int32) *GetTimingSyntheticTaskResponseBodyDataMonitorConfNetICMP {
	s.Interval = &v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfNetICMP) SetPackageNum(v int32) *GetTimingSyntheticTaskResponseBodyDataMonitorConfNetICMP {
	s.PackageNum = &v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfNetICMP) SetPackageSize(v int32) *GetTimingSyntheticTaskResponseBodyDataMonitorConfNetICMP {
	s.PackageSize = &v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfNetICMP) SetSplitPackage(v bool) *GetTimingSyntheticTaskResponseBodyDataMonitorConfNetICMP {
	s.SplitPackage = &v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfNetICMP) SetTargetUrl(v string) *GetTimingSyntheticTaskResponseBodyDataMonitorConfNetICMP {
	s.TargetUrl = &v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfNetICMP) SetTimeout(v int64) *GetTimingSyntheticTaskResponseBodyDataMonitorConfNetICMP {
	s.Timeout = &v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfNetICMP) SetTracertEnable(v bool) *GetTimingSyntheticTaskResponseBodyDataMonitorConfNetICMP {
	s.TracertEnable = &v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfNetICMP) SetTracertNumMax(v int32) *GetTimingSyntheticTaskResponseBodyDataMonitorConfNetICMP {
	s.TracertNumMax = &v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfNetICMP) SetTracertTimeout(v int64) *GetTimingSyntheticTaskResponseBodyDataMonitorConfNetICMP {
	s.TracertTimeout = &v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfNetICMP) Validate() error {
	return dara.Validate(s)
}

type GetTimingSyntheticTaskResponseBodyDataMonitorConfNetTCP struct {
	// The number of TCP connections that are established in a test. Minimum value: 1. Maximum value: 16. Default value: 4.
	//
	// example:
	//
	// 4
	ConnectTimes *int32 `json:"ConnectTimes,omitempty" xml:"ConnectTimes,omitempty"`
	// The interval at which TCP connections are established. Unit: milliseconds. Minimum value: 200. Maximum value: 10000. Default value: 200.
	//
	// example:
	//
	// 200
	Interval *int64 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The destination host IP address.
	//
	// example:
	//
	// 127.0.0.1:8888
	TargetUrl *string `json:"TargetUrl,omitempty" xml:"TargetUrl,omitempty"`
	// The timeout period for the TCP synthetic test. Unit: milliseconds. Minimum value: 1000. Maximum value: 300000. Default value: 20000.
	//
	// example:
	//
	// 20000
	Timeout *int64 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// Specifies whether to enable the tracert command. Default value: true.
	//
	// example:
	//
	// true
	TracertEnable *bool `json:"TracertEnable,omitempty" xml:"TracertEnable,omitempty"`
	// The maximum number of hops for tracert. Minimum value: 1. Maximum value: 128. Default value: 20.
	//
	// example:
	//
	// 20
	TracertNumMax *int32 `json:"TracertNumMax,omitempty" xml:"TracertNumMax,omitempty"`
	// The timeout period of tracert. Unit: milliseconds. Minimum value: 1000. Maximum value: 300000. Default value: 60000.
	//
	// example:
	//
	// 60000
	TracertTimeout *int64 `json:"TracertTimeout,omitempty" xml:"TracertTimeout,omitempty"`
}

func (s GetTimingSyntheticTaskResponseBodyDataMonitorConfNetTCP) String() string {
	return dara.Prettify(s)
}

func (s GetTimingSyntheticTaskResponseBodyDataMonitorConfNetTCP) GoString() string {
	return s.String()
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfNetTCP) GetConnectTimes() *int32 {
	return s.ConnectTimes
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfNetTCP) GetInterval() *int64 {
	return s.Interval
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfNetTCP) GetTargetUrl() *string {
	return s.TargetUrl
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfNetTCP) GetTimeout() *int64 {
	return s.Timeout
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfNetTCP) GetTracertEnable() *bool {
	return s.TracertEnable
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfNetTCP) GetTracertNumMax() *int32 {
	return s.TracertNumMax
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfNetTCP) GetTracertTimeout() *int64 {
	return s.TracertTimeout
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfNetTCP) SetConnectTimes(v int32) *GetTimingSyntheticTaskResponseBodyDataMonitorConfNetTCP {
	s.ConnectTimes = &v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfNetTCP) SetInterval(v int64) *GetTimingSyntheticTaskResponseBodyDataMonitorConfNetTCP {
	s.Interval = &v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfNetTCP) SetTargetUrl(v string) *GetTimingSyntheticTaskResponseBodyDataMonitorConfNetTCP {
	s.TargetUrl = &v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfNetTCP) SetTimeout(v int64) *GetTimingSyntheticTaskResponseBodyDataMonitorConfNetTCP {
	s.Timeout = &v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfNetTCP) SetTracertEnable(v bool) *GetTimingSyntheticTaskResponseBodyDataMonitorConfNetTCP {
	s.TracertEnable = &v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfNetTCP) SetTracertNumMax(v int32) *GetTimingSyntheticTaskResponseBodyDataMonitorConfNetTCP {
	s.TracertNumMax = &v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfNetTCP) SetTracertTimeout(v int64) *GetTimingSyntheticTaskResponseBodyDataMonitorConfNetTCP {
	s.TracertTimeout = &v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfNetTCP) Validate() error {
	return dara.Validate(s)
}

type GetTimingSyntheticTaskResponseBodyDataMonitorConfStream struct {
	// Custom header, JSON Map format.
	CustomHeaderContent map[string]*string `json:"CustomHeaderContent,omitempty" xml:"CustomHeaderContent,omitempty"`
	// Player, default is 12 if not specified.
	//
	// - 12: VLC
	//
	// - 2: Flash Player
	//
	// example:
	//
	// 12
	PlayerType *int32 `json:"PlayerType,omitempty" xml:"PlayerType,omitempty"`
	// Resource address type:
	//
	// - 1: Resource address.
	//
	// - 0: Page address. If not passed, the default value is 0.
	//
	// example:
	//
	// 0
	StreamAddressType *int32 `json:"StreamAddressType,omitempty" xml:"StreamAddressType,omitempty"`
	// Monitoring duration, in seconds, supports up to 60 seconds. If not specified, the default value is 60 seconds.
	//
	// example:
	//
	// 30
	StreamMonitorTimeout *int32 `json:"StreamMonitorTimeout,omitempty" xml:"StreamMonitorTimeout,omitempty"`
	// Audio and video flag:
	//
	// - 0: video
	//
	// - 1: audio
	//
	// example:
	//
	// 0
	StreamType *int32 `json:"StreamType,omitempty" xml:"StreamType,omitempty"`
	// Streaming media resource address.
	//
	// example:
	//
	// http://www.aliyun.com/stream/test.mp4
	TargetUrl *string `json:"TargetUrl,omitempty" xml:"TargetUrl,omitempty"`
	// DNS hijacking whitelist. Matching rules support IP, IP wildcard, subnet mask and CNAME. You can fill in multiple matching rules, and multiple matching rules are separated by vertical bars (|). For example: `www.aliyun.com:203.0.3.55|203.3.44.67`, which means that all IPs except 203.0.3.55 and 203.3.44.67 under the www.aliyun.com domain name are hijacked.
	//
	// example:
	//
	// www.aliyun.com:203.0.3.55|203.3.44.67
	WhiteList *string `json:"WhiteList,omitempty" xml:"WhiteList,omitempty"`
}

func (s GetTimingSyntheticTaskResponseBodyDataMonitorConfStream) String() string {
	return dara.Prettify(s)
}

func (s GetTimingSyntheticTaskResponseBodyDataMonitorConfStream) GoString() string {
	return s.String()
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfStream) GetCustomHeaderContent() map[string]*string {
	return s.CustomHeaderContent
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfStream) GetPlayerType() *int32 {
	return s.PlayerType
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfStream) GetStreamAddressType() *int32 {
	return s.StreamAddressType
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfStream) GetStreamMonitorTimeout() *int32 {
	return s.StreamMonitorTimeout
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfStream) GetStreamType() *int32 {
	return s.StreamType
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfStream) GetTargetUrl() *string {
	return s.TargetUrl
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfStream) GetWhiteList() *string {
	return s.WhiteList
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfStream) SetCustomHeaderContent(v map[string]*string) *GetTimingSyntheticTaskResponseBodyDataMonitorConfStream {
	s.CustomHeaderContent = v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfStream) SetPlayerType(v int32) *GetTimingSyntheticTaskResponseBodyDataMonitorConfStream {
	s.PlayerType = &v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfStream) SetStreamAddressType(v int32) *GetTimingSyntheticTaskResponseBodyDataMonitorConfStream {
	s.StreamAddressType = &v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfStream) SetStreamMonitorTimeout(v int32) *GetTimingSyntheticTaskResponseBodyDataMonitorConfStream {
	s.StreamMonitorTimeout = &v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfStream) SetStreamType(v int32) *GetTimingSyntheticTaskResponseBodyDataMonitorConfStream {
	s.StreamType = &v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfStream) SetTargetUrl(v string) *GetTimingSyntheticTaskResponseBodyDataMonitorConfStream {
	s.TargetUrl = &v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfStream) SetWhiteList(v string) *GetTimingSyntheticTaskResponseBodyDataMonitorConfStream {
	s.WhiteList = &v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfStream) Validate() error {
	return dara.Validate(s)
}

type GetTimingSyntheticTaskResponseBodyDataMonitorConfWebsite struct {
	// Specifies whether to automatically scroll up and down the screen to load a page. 0: No. 1: Yes. Default value: 0.
	//
	// example:
	//
	// 0
	AutomaticScrolling *int32 `json:"AutomaticScrolling,omitempty" xml:"AutomaticScrolling,omitempty"`
	// Specifies whether to create a custom header. 0: No. 1: The first packet is modified. 2: All packets are modified. Default value: 0.
	//
	// example:
	//
	// 0
	CustomHeader *int32 `json:"CustomHeader,omitempty" xml:"CustomHeader,omitempty"`
	// The custom header. Format: JSON map.
	CustomHeaderContent map[string]*string `json:"CustomHeaderContent,omitempty" xml:"CustomHeaderContent,omitempty"`
	// When resolving a domain name (such as www.aliyun.com), if the resolved IP address or CNAME is not in the DNS hijacking whitelist, the user will fail to access or return a non-Aliyun target IP; if the IP or CNAME in the resolution result is in the DNS whitelist, it will be deemed that no DNS hijacking has occurred.
	//
	// Fill in the format: `domain name: matching rule`. Matching rules support IP, IP wildcard, subnet mask and CNAME. You can fill in multiple matching rules, and multiple matching rules are separated by vertical bars (|).
	//
	// For example: `www.aliyun.com:203.0.3.55|203.3.44.67`, which means that all IPs except 203.0.3.55 and 203.3.44.67 under the www.aliyun.com domain name are hijacked.
	//
	// example:
	//
	// www.aliyun.com:203.0.3.55|203.3.44.67
	DNSHijackWhitelist *string `json:"DNSHijackWhitelist,omitempty" xml:"DNSHijackWhitelist,omitempty"`
	// Specifies whether to disable the cache. 0: No. 1: Yes. Default value: 1.
	//
	// example:
	//
	// 1
	DisableCache *int32 `json:"DisableCache,omitempty" xml:"DisableCache,omitempty"`
	// Specifies whether to accept compressed files based on the HTTP Accept-Encoding request header. 0: No. 1: Yes. Default value: 0.
	//
	// example:
	//
	// 0
	DisableCompression *int32 `json:"DisableCompression,omitempty" xml:"DisableCompression,omitempty"`
	// If an element configured in the element blacklist appears during page loading, no request will be made to load the element.
	//
	// example:
	//
	// www.example.com/a.jpg
	ElementBlacklist *string `json:"ElementBlacklist,omitempty" xml:"ElementBlacklist,omitempty"`
	// Specifies whether to exclude invalid IP addresses.
	//
	// 	- 1: No
	//
	// 	- 0: Yes
	//
	// example:
	//
	// 0
	FilterInvalidIP *int32 `json:"FilterInvalidIP,omitempty" xml:"FilterInvalidIP,omitempty"`
	// Identify elements: Set the total number of elements to browse the page.
	//
	// example:
	//
	// 0
	FlowHijackJumpTimes *int32 `json:"FlowHijackJumpTimes,omitempty" xml:"FlowHijackJumpTimes,omitempty"`
	// Hijacking flag: Set the key information for matching. Fill in the hijacking judgment keyword or key element, and asterisks (*) are allowed.
	//
	// example:
	//
	// aliyun
	FlowHijackLogo *string `json:"FlowHijackLogo,omitempty" xml:"FlowHijackLogo,omitempty"`
	// Specifies whether to ignore SSL certificate errors during browsing. 0: No. 1: Yes. Default value: 1.
	//
	// example:
	//
	// 0
	IgnoreCertificateError *int32 `json:"IgnoreCertificateError,omitempty" xml:"IgnoreCertificateError,omitempty"`
	// The monitoring timeout period. Unit: milliseconds. Default value: 20000. Minimum value: 5000. Maximum value: 300000.
	//
	// example:
	//
	// 20000
	MonitorTimeout *int32 `json:"MonitorTimeout,omitempty" xml:"MonitorTimeout,omitempty"`
	// If any element other than the domain name setting appears on the monitoring page, it means that the page has been tampered. Common manifestations include pop-up ads, floating ads, jumps, etc.
	//
	// Fill in the format: `domain name: element`. Elements support wildcards, and multiple elements can be filled in. Multiple elements are separated by vertical bars (|). For example: `www.aliyun.com:|/cc/bb/a.gif|/vv/bb/cc.jpg`, which means that all elements except the basic document, /cc/bb/a.gif and /vv/bb/cc.jpg under the www.aliyun.com domain name are considered to be tampered with.
	//
	// example:
	//
	// www.aliyun.com:|/cc/bb/a.gif|/vv/bb/cc.jpg
	PageTamper *string `json:"PageTamper,omitempty" xml:"PageTamper,omitempty"`
	// Specifies whether to continue browsing after redirection. 0: No, 1:Yes. Default value: 1.
	//
	// example:
	//
	// 1
	Redirection *int32 `json:"Redirection,omitempty" xml:"Redirection,omitempty"`
	// The time threshold that is used to define a slow element. Unit: milliseconds. Default value: 5000. Minimum value: 1. Maximum value: 300000.
	//
	// example:
	//
	// 5000
	SlowElementThreshold *int64 `json:"SlowElementThreshold,omitempty" xml:"SlowElementThreshold,omitempty"`
	// The destination URL.
	//
	// example:
	//
	// http://www.aliyun.com
	TargetUrl *string `json:"TargetUrl,omitempty" xml:"TargetUrl,omitempty"`
	// The verification string is an arbitrary string in the source code of the monitoring page. If the source code returned by the client contains any string in the blacklist, an error 650 &quot;Verification string failed&quot; will be reported. Multiple strings are separated by vertical bars (|).
	//
	// example:
	//
	// error
	VerifyStringBlacklist *string `json:"VerifyStringBlacklist,omitempty" xml:"VerifyStringBlacklist,omitempty"`
	// The verification string is an arbitrary string in the source code of the monitoring page. The source code returned by the client must contain all the strings in the whitelist, otherwise an error 650 &quot;Verification string failed&quot; will be reported. Multiple strings are separated by a vertical bar (|).
	//
	// example:
	//
	// success
	VerifyStringWhitelist *string `json:"VerifyStringWhitelist,omitempty" xml:"VerifyStringWhitelist,omitempty"`
	// The maximum waiting time. Unit: milliseconds. Default value: 5000. Minimum value: 5000. Maximum value: 300000.
	//
	// example:
	//
	// 5000
	WaitCompletionTime *int64 `json:"WaitCompletionTime,omitempty" xml:"WaitCompletionTime,omitempty"`
}

func (s GetTimingSyntheticTaskResponseBodyDataMonitorConfWebsite) String() string {
	return dara.Prettify(s)
}

func (s GetTimingSyntheticTaskResponseBodyDataMonitorConfWebsite) GoString() string {
	return s.String()
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfWebsite) GetAutomaticScrolling() *int32 {
	return s.AutomaticScrolling
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfWebsite) GetCustomHeader() *int32 {
	return s.CustomHeader
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfWebsite) GetCustomHeaderContent() map[string]*string {
	return s.CustomHeaderContent
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfWebsite) GetDNSHijackWhitelist() *string {
	return s.DNSHijackWhitelist
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfWebsite) GetDisableCache() *int32 {
	return s.DisableCache
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfWebsite) GetDisableCompression() *int32 {
	return s.DisableCompression
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfWebsite) GetElementBlacklist() *string {
	return s.ElementBlacklist
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfWebsite) GetFilterInvalidIP() *int32 {
	return s.FilterInvalidIP
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfWebsite) GetFlowHijackJumpTimes() *int32 {
	return s.FlowHijackJumpTimes
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfWebsite) GetFlowHijackLogo() *string {
	return s.FlowHijackLogo
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfWebsite) GetIgnoreCertificateError() *int32 {
	return s.IgnoreCertificateError
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfWebsite) GetMonitorTimeout() *int32 {
	return s.MonitorTimeout
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfWebsite) GetPageTamper() *string {
	return s.PageTamper
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfWebsite) GetRedirection() *int32 {
	return s.Redirection
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfWebsite) GetSlowElementThreshold() *int64 {
	return s.SlowElementThreshold
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfWebsite) GetTargetUrl() *string {
	return s.TargetUrl
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfWebsite) GetVerifyStringBlacklist() *string {
	return s.VerifyStringBlacklist
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfWebsite) GetVerifyStringWhitelist() *string {
	return s.VerifyStringWhitelist
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfWebsite) GetWaitCompletionTime() *int64 {
	return s.WaitCompletionTime
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfWebsite) SetAutomaticScrolling(v int32) *GetTimingSyntheticTaskResponseBodyDataMonitorConfWebsite {
	s.AutomaticScrolling = &v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfWebsite) SetCustomHeader(v int32) *GetTimingSyntheticTaskResponseBodyDataMonitorConfWebsite {
	s.CustomHeader = &v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfWebsite) SetCustomHeaderContent(v map[string]*string) *GetTimingSyntheticTaskResponseBodyDataMonitorConfWebsite {
	s.CustomHeaderContent = v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfWebsite) SetDNSHijackWhitelist(v string) *GetTimingSyntheticTaskResponseBodyDataMonitorConfWebsite {
	s.DNSHijackWhitelist = &v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfWebsite) SetDisableCache(v int32) *GetTimingSyntheticTaskResponseBodyDataMonitorConfWebsite {
	s.DisableCache = &v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfWebsite) SetDisableCompression(v int32) *GetTimingSyntheticTaskResponseBodyDataMonitorConfWebsite {
	s.DisableCompression = &v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfWebsite) SetElementBlacklist(v string) *GetTimingSyntheticTaskResponseBodyDataMonitorConfWebsite {
	s.ElementBlacklist = &v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfWebsite) SetFilterInvalidIP(v int32) *GetTimingSyntheticTaskResponseBodyDataMonitorConfWebsite {
	s.FilterInvalidIP = &v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfWebsite) SetFlowHijackJumpTimes(v int32) *GetTimingSyntheticTaskResponseBodyDataMonitorConfWebsite {
	s.FlowHijackJumpTimes = &v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfWebsite) SetFlowHijackLogo(v string) *GetTimingSyntheticTaskResponseBodyDataMonitorConfWebsite {
	s.FlowHijackLogo = &v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfWebsite) SetIgnoreCertificateError(v int32) *GetTimingSyntheticTaskResponseBodyDataMonitorConfWebsite {
	s.IgnoreCertificateError = &v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfWebsite) SetMonitorTimeout(v int32) *GetTimingSyntheticTaskResponseBodyDataMonitorConfWebsite {
	s.MonitorTimeout = &v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfWebsite) SetPageTamper(v string) *GetTimingSyntheticTaskResponseBodyDataMonitorConfWebsite {
	s.PageTamper = &v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfWebsite) SetRedirection(v int32) *GetTimingSyntheticTaskResponseBodyDataMonitorConfWebsite {
	s.Redirection = &v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfWebsite) SetSlowElementThreshold(v int64) *GetTimingSyntheticTaskResponseBodyDataMonitorConfWebsite {
	s.SlowElementThreshold = &v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfWebsite) SetTargetUrl(v string) *GetTimingSyntheticTaskResponseBodyDataMonitorConfWebsite {
	s.TargetUrl = &v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfWebsite) SetVerifyStringBlacklist(v string) *GetTimingSyntheticTaskResponseBodyDataMonitorConfWebsite {
	s.VerifyStringBlacklist = &v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfWebsite) SetVerifyStringWhitelist(v string) *GetTimingSyntheticTaskResponseBodyDataMonitorConfWebsite {
	s.VerifyStringWhitelist = &v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfWebsite) SetWaitCompletionTime(v int64) *GetTimingSyntheticTaskResponseBodyDataMonitorConfWebsite {
	s.WaitCompletionTime = &v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitorConfWebsite) Validate() error {
	return dara.Validate(s)
}

type GetTimingSyntheticTaskResponseBodyDataMonitors struct {
	// The city code.
	//
	// example:
	//
	// 110100
	CityCode *string `json:"CityCode,omitempty" xml:"CityCode,omitempty"`
	// The client type of the monitoring point. Valid values: 1: data center. 2: Internet. 3: mobile device. 4: ECS instance.
	//
	// example:
	//
	// 1
	ClientType *int32 `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	// The carrier code.
	//
	// example:
	//
	// 1
	OperatorCode *string `json:"OperatorCode,omitempty" xml:"OperatorCode,omitempty"`
}

func (s GetTimingSyntheticTaskResponseBodyDataMonitors) String() string {
	return dara.Prettify(s)
}

func (s GetTimingSyntheticTaskResponseBodyDataMonitors) GoString() string {
	return s.String()
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitors) GetCityCode() *string {
	return s.CityCode
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitors) GetClientType() *int32 {
	return s.ClientType
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitors) GetOperatorCode() *string {
	return s.OperatorCode
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitors) SetCityCode(v string) *GetTimingSyntheticTaskResponseBodyDataMonitors {
	s.CityCode = &v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitors) SetClientType(v int32) *GetTimingSyntheticTaskResponseBodyDataMonitors {
	s.ClientType = &v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitors) SetOperatorCode(v string) *GetTimingSyntheticTaskResponseBodyDataMonitors {
	s.OperatorCode = &v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyDataMonitors) Validate() error {
	return dara.Validate(s)
}

type GetTimingSyntheticTaskResponseBodyDataTags struct {
	// The key of the tag.
	//
	// example:
	//
	// user1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// myweb
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetTimingSyntheticTaskResponseBodyDataTags) String() string {
	return dara.Prettify(s)
}

func (s GetTimingSyntheticTaskResponseBodyDataTags) GoString() string {
	return s.String()
}

func (s *GetTimingSyntheticTaskResponseBodyDataTags) GetKey() *string {
	return s.Key
}

func (s *GetTimingSyntheticTaskResponseBodyDataTags) GetValue() *string {
	return s.Value
}

func (s *GetTimingSyntheticTaskResponseBodyDataTags) SetKey(v string) *GetTimingSyntheticTaskResponseBodyDataTags {
	s.Key = &v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyDataTags) SetValue(v string) *GetTimingSyntheticTaskResponseBodyDataTags {
	s.Value = &v
	return s
}

func (s *GetTimingSyntheticTaskResponseBodyDataTags) Validate() error {
	return dara.Validate(s)
}
