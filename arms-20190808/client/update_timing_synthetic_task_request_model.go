// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTimingSyntheticTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAvailableAssertions(v []*UpdateTimingSyntheticTaskRequestAvailableAssertions) *UpdateTimingSyntheticTaskRequest
	GetAvailableAssertions() []*UpdateTimingSyntheticTaskRequestAvailableAssertions
	SetCommonSetting(v *UpdateTimingSyntheticTaskRequestCommonSetting) *UpdateTimingSyntheticTaskRequest
	GetCommonSetting() *UpdateTimingSyntheticTaskRequestCommonSetting
	SetCustomPeriod(v *UpdateTimingSyntheticTaskRequestCustomPeriod) *UpdateTimingSyntheticTaskRequest
	GetCustomPeriod() *UpdateTimingSyntheticTaskRequestCustomPeriod
	SetFrequency(v string) *UpdateTimingSyntheticTaskRequest
	GetFrequency() *string
	SetMonitorConf(v *UpdateTimingSyntheticTaskRequestMonitorConf) *UpdateTimingSyntheticTaskRequest
	GetMonitorConf() *UpdateTimingSyntheticTaskRequestMonitorConf
	SetMonitors(v []*UpdateTimingSyntheticTaskRequestMonitors) *UpdateTimingSyntheticTaskRequest
	GetMonitors() []*UpdateTimingSyntheticTaskRequestMonitors
	SetName(v string) *UpdateTimingSyntheticTaskRequest
	GetName() *string
	SetRegionId(v string) *UpdateTimingSyntheticTaskRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *UpdateTimingSyntheticTaskRequest
	GetResourceGroupId() *string
	SetTags(v []*UpdateTimingSyntheticTaskRequestTags) *UpdateTimingSyntheticTaskRequest
	GetTags() []*UpdateTimingSyntheticTaskRequestTags
	SetTaskId(v string) *UpdateTimingSyntheticTaskRequest
	GetTaskId() *string
}

type UpdateTimingSyntheticTaskRequest struct {
	// The list of assertions.
	AvailableAssertions []*UpdateTimingSyntheticTaskRequestAvailableAssertions `json:"AvailableAssertions,omitempty" xml:"AvailableAssertions,omitempty" type:"Repeated"`
	// The general settings.
	CommonSetting *UpdateTimingSyntheticTaskRequestCommonSetting `json:"CommonSetting,omitempty" xml:"CommonSetting,omitempty" type:"Struct"`
	// The custom cycle.
	CustomPeriod *UpdateTimingSyntheticTaskRequestCustomPeriod `json:"CustomPeriod,omitempty" xml:"CustomPeriod,omitempty" type:"Struct"`
	// The detection frequency. Valid values: 1m, 5m, 10m, 15m, 20m, 30m, 1h, 2h, 3h, 4h, 6h, 8h, 12h, and 24h.
	//
	// example:
	//
	// 5m
	Frequency *string `json:"Frequency,omitempty" xml:"Frequency,omitempty"`
	// The monitoring configurations.
	MonitorConf *UpdateTimingSyntheticTaskRequestMonitorConf `json:"MonitorConf,omitempty" xml:"MonitorConf,omitempty" type:"Struct"`
	// The list of monitoring points.
	Monitors []*UpdateTimingSyntheticTaskRequestMonitors `json:"Monitors,omitempty" xml:"Monitors,omitempty" type:"Repeated"`
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
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmxyexli2****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The list of tags.
	Tags []*UpdateTimingSyntheticTaskRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The ID of the synthetic monitoring task.
	//
	// example:
	//
	// 5308a2691f59422c8c3b7aeccxxxxxxx
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s UpdateTimingSyntheticTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTimingSyntheticTaskRequest) GoString() string {
	return s.String()
}

func (s *UpdateTimingSyntheticTaskRequest) GetAvailableAssertions() []*UpdateTimingSyntheticTaskRequestAvailableAssertions {
	return s.AvailableAssertions
}

func (s *UpdateTimingSyntheticTaskRequest) GetCommonSetting() *UpdateTimingSyntheticTaskRequestCommonSetting {
	return s.CommonSetting
}

func (s *UpdateTimingSyntheticTaskRequest) GetCustomPeriod() *UpdateTimingSyntheticTaskRequestCustomPeriod {
	return s.CustomPeriod
}

func (s *UpdateTimingSyntheticTaskRequest) GetFrequency() *string {
	return s.Frequency
}

func (s *UpdateTimingSyntheticTaskRequest) GetMonitorConf() *UpdateTimingSyntheticTaskRequestMonitorConf {
	return s.MonitorConf
}

func (s *UpdateTimingSyntheticTaskRequest) GetMonitors() []*UpdateTimingSyntheticTaskRequestMonitors {
	return s.Monitors
}

func (s *UpdateTimingSyntheticTaskRequest) GetName() *string {
	return s.Name
}

func (s *UpdateTimingSyntheticTaskRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateTimingSyntheticTaskRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *UpdateTimingSyntheticTaskRequest) GetTags() []*UpdateTimingSyntheticTaskRequestTags {
	return s.Tags
}

func (s *UpdateTimingSyntheticTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *UpdateTimingSyntheticTaskRequest) SetAvailableAssertions(v []*UpdateTimingSyntheticTaskRequestAvailableAssertions) *UpdateTimingSyntheticTaskRequest {
	s.AvailableAssertions = v
	return s
}

func (s *UpdateTimingSyntheticTaskRequest) SetCommonSetting(v *UpdateTimingSyntheticTaskRequestCommonSetting) *UpdateTimingSyntheticTaskRequest {
	s.CommonSetting = v
	return s
}

func (s *UpdateTimingSyntheticTaskRequest) SetCustomPeriod(v *UpdateTimingSyntheticTaskRequestCustomPeriod) *UpdateTimingSyntheticTaskRequest {
	s.CustomPeriod = v
	return s
}

func (s *UpdateTimingSyntheticTaskRequest) SetFrequency(v string) *UpdateTimingSyntheticTaskRequest {
	s.Frequency = &v
	return s
}

func (s *UpdateTimingSyntheticTaskRequest) SetMonitorConf(v *UpdateTimingSyntheticTaskRequestMonitorConf) *UpdateTimingSyntheticTaskRequest {
	s.MonitorConf = v
	return s
}

func (s *UpdateTimingSyntheticTaskRequest) SetMonitors(v []*UpdateTimingSyntheticTaskRequestMonitors) *UpdateTimingSyntheticTaskRequest {
	s.Monitors = v
	return s
}

func (s *UpdateTimingSyntheticTaskRequest) SetName(v string) *UpdateTimingSyntheticTaskRequest {
	s.Name = &v
	return s
}

func (s *UpdateTimingSyntheticTaskRequest) SetRegionId(v string) *UpdateTimingSyntheticTaskRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateTimingSyntheticTaskRequest) SetResourceGroupId(v string) *UpdateTimingSyntheticTaskRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *UpdateTimingSyntheticTaskRequest) SetTags(v []*UpdateTimingSyntheticTaskRequestTags) *UpdateTimingSyntheticTaskRequest {
	s.Tags = v
	return s
}

func (s *UpdateTimingSyntheticTaskRequest) SetTaskId(v string) *UpdateTimingSyntheticTaskRequest {
	s.TaskId = &v
	return s
}

func (s *UpdateTimingSyntheticTaskRequest) Validate() error {
	if s.AvailableAssertions != nil {
		for _, item := range s.AvailableAssertions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.CommonSetting != nil {
		if err := s.CommonSetting.Validate(); err != nil {
			return err
		}
	}
	if s.CustomPeriod != nil {
		if err := s.CustomPeriod.Validate(); err != nil {
			return err
		}
	}
	if s.MonitorConf != nil {
		if err := s.MonitorConf.Validate(); err != nil {
			return err
		}
	}
	if s.Monitors != nil {
		for _, item := range s.Monitors {
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

type UpdateTimingSyntheticTaskRequestAvailableAssertions struct {
	// The expected value.
	//
	// example:
	//
	// 200
	Expect *string `json:"Expect,omitempty" xml:"Expect,omitempty"`
	// The condition. gt: greater than. gte: greater than or equal to. lt: less than. lte: less than or equal to. eq: equal to. neq: not equal to. ctn: contain. nctn: does not contain. exist: exist. n_exist: does not exist. belong: belong to. n_belong: does not belong to. reg_match: regular expression.
	//
	// example:
	//
	// eq
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// The check target. If you set the type parameter to HttpResCode, HttpResBody, or HttpResponseTime, you do not need to set the target parameter. If you set the type parameter to HttpResHead, you must specify the key in the header. If you set the type parameter to HttpResBodyJson, use jsonPath.
	//
	// example:
	//
	// HttpResponseTime
	Target *string `json:"Target,omitempty" xml:"Target,omitempty"`
	// The assertion type. Valid values: HttpResCode, HttpResHead, HttpResBody, HttpResBodyJson, HttpResponseTime, IcmpPackLoss (packet loss rate), IcmpPackMaxLatency (maximum packet latency), IcmpPackAvgLatency (average packet latency), TraceRouteHops (number of hops), DnsARecord (A record), DnsCName (CNAME), websiteTTFB (time to first packet), websiteTTLB (time to last packet), websiteFST (first paint time), websiteFFST (first meaningful paint), websiteOnload (full loaded time). For more information, see the following description.
	//
	// example:
	//
	// TraceRouteHops
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateTimingSyntheticTaskRequestAvailableAssertions) String() string {
	return dara.Prettify(s)
}

func (s UpdateTimingSyntheticTaskRequestAvailableAssertions) GoString() string {
	return s.String()
}

func (s *UpdateTimingSyntheticTaskRequestAvailableAssertions) GetExpect() *string {
	return s.Expect
}

func (s *UpdateTimingSyntheticTaskRequestAvailableAssertions) GetOperator() *string {
	return s.Operator
}

func (s *UpdateTimingSyntheticTaskRequestAvailableAssertions) GetTarget() *string {
	return s.Target
}

func (s *UpdateTimingSyntheticTaskRequestAvailableAssertions) GetType() *string {
	return s.Type
}

func (s *UpdateTimingSyntheticTaskRequestAvailableAssertions) SetExpect(v string) *UpdateTimingSyntheticTaskRequestAvailableAssertions {
	s.Expect = &v
	return s
}

func (s *UpdateTimingSyntheticTaskRequestAvailableAssertions) SetOperator(v string) *UpdateTimingSyntheticTaskRequestAvailableAssertions {
	s.Operator = &v
	return s
}

func (s *UpdateTimingSyntheticTaskRequestAvailableAssertions) SetTarget(v string) *UpdateTimingSyntheticTaskRequestAvailableAssertions {
	s.Target = &v
	return s
}

func (s *UpdateTimingSyntheticTaskRequestAvailableAssertions) SetType(v string) *UpdateTimingSyntheticTaskRequestAvailableAssertions {
	s.Type = &v
	return s
}

func (s *UpdateTimingSyntheticTaskRequestAvailableAssertions) Validate() error {
	return dara.Validate(s)
}

type UpdateTimingSyntheticTaskRequestCommonSetting struct {
	// The custom host settings.
	CustomHost *UpdateTimingSyntheticTaskRequestCommonSettingCustomHost `json:"CustomHost,omitempty" xml:"CustomHost,omitempty" type:"Struct"`
	// The reserved parameters.
	CustomPrometheusSetting *UpdateTimingSyntheticTaskRequestCommonSettingCustomPrometheusSetting `json:"CustomPrometheusSetting,omitempty" xml:"CustomPrometheusSetting,omitempty" type:"Struct"`
	// The information about the virtual private cloud (VPC). If the destination URL is an Alibaba Cloud internal endpoint, you need to configure a VPC.
	CustomVPCSetting *UpdateTimingSyntheticTaskRequestCommonSettingCustomVPCSetting `json:"CustomVPCSetting,omitempty" xml:"CustomVPCSetting,omitempty" type:"Struct"`
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
	// Specifies whether to enable tracing.
	//
	// example:
	//
	// true
	IsOpenTrace *bool `json:"IsOpenTrace,omitempty" xml:"IsOpenTrace,omitempty"`
	// Specifies whether to evenly distribute monitoring samples. Valid values:
	//
	// 	- 0: no
	//
	// 	- 1: yes
	//
	// example:
	//
	// 0
	MonitorSamples *int32 `json:"MonitorSamples,omitempty" xml:"MonitorSamples,omitempty"`
	// The type of the client for tracing. Valid values:
	//
	// 	- 0: ARMS agent
	//
	// 	- 1: OpenTelemetry
	//
	// 	- 2: Jaeger
	//
	// example:
	//
	// 1
	TraceClientType *int32 `json:"TraceClientType,omitempty" xml:"TraceClientType,omitempty"`
	// The region to which trace data is reported.
	//
	// example:
	//
	// cn-hangzhou
	XtraceRegion *string `json:"XtraceRegion,omitempty" xml:"XtraceRegion,omitempty"`
}

func (s UpdateTimingSyntheticTaskRequestCommonSetting) String() string {
	return dara.Prettify(s)
}

func (s UpdateTimingSyntheticTaskRequestCommonSetting) GoString() string {
	return s.String()
}

func (s *UpdateTimingSyntheticTaskRequestCommonSetting) GetCustomHost() *UpdateTimingSyntheticTaskRequestCommonSettingCustomHost {
	return s.CustomHost
}

func (s *UpdateTimingSyntheticTaskRequestCommonSetting) GetCustomPrometheusSetting() *UpdateTimingSyntheticTaskRequestCommonSettingCustomPrometheusSetting {
	return s.CustomPrometheusSetting
}

func (s *UpdateTimingSyntheticTaskRequestCommonSetting) GetCustomVPCSetting() *UpdateTimingSyntheticTaskRequestCommonSettingCustomVPCSetting {
	return s.CustomVPCSetting
}

func (s *UpdateTimingSyntheticTaskRequestCommonSetting) GetIpType() *int32 {
	return s.IpType
}

func (s *UpdateTimingSyntheticTaskRequestCommonSetting) GetIsOpenTrace() *bool {
	return s.IsOpenTrace
}

func (s *UpdateTimingSyntheticTaskRequestCommonSetting) GetMonitorSamples() *int32 {
	return s.MonitorSamples
}

func (s *UpdateTimingSyntheticTaskRequestCommonSetting) GetTraceClientType() *int32 {
	return s.TraceClientType
}

func (s *UpdateTimingSyntheticTaskRequestCommonSetting) GetXtraceRegion() *string {
	return s.XtraceRegion
}

func (s *UpdateTimingSyntheticTaskRequestCommonSetting) SetCustomHost(v *UpdateTimingSyntheticTaskRequestCommonSettingCustomHost) *UpdateTimingSyntheticTaskRequestCommonSetting {
	s.CustomHost = v
	return s
}

func (s *UpdateTimingSyntheticTaskRequestCommonSetting) SetCustomPrometheusSetting(v *UpdateTimingSyntheticTaskRequestCommonSettingCustomPrometheusSetting) *UpdateTimingSyntheticTaskRequestCommonSetting {
	s.CustomPrometheusSetting = v
	return s
}

func (s *UpdateTimingSyntheticTaskRequestCommonSetting) SetCustomVPCSetting(v *UpdateTimingSyntheticTaskRequestCommonSettingCustomVPCSetting) *UpdateTimingSyntheticTaskRequestCommonSetting {
	s.CustomVPCSetting = v
	return s
}

func (s *UpdateTimingSyntheticTaskRequestCommonSetting) SetIpType(v int32) *UpdateTimingSyntheticTaskRequestCommonSetting {
	s.IpType = &v
	return s
}

func (s *UpdateTimingSyntheticTaskRequestCommonSetting) SetIsOpenTrace(v bool) *UpdateTimingSyntheticTaskRequestCommonSetting {
	s.IsOpenTrace = &v
	return s
}

func (s *UpdateTimingSyntheticTaskRequestCommonSetting) SetMonitorSamples(v int32) *UpdateTimingSyntheticTaskRequestCommonSetting {
	s.MonitorSamples = &v
	return s
}

func (s *UpdateTimingSyntheticTaskRequestCommonSetting) SetTraceClientType(v int32) *UpdateTimingSyntheticTaskRequestCommonSetting {
	s.TraceClientType = &v
	return s
}

func (s *UpdateTimingSyntheticTaskRequestCommonSetting) SetXtraceRegion(v string) *UpdateTimingSyntheticTaskRequestCommonSetting {
	s.XtraceRegion = &v
	return s
}

func (s *UpdateTimingSyntheticTaskRequestCommonSetting) Validate() error {
	if s.CustomHost != nil {
		if err := s.CustomHost.Validate(); err != nil {
			return err
		}
	}
	if s.CustomPrometheusSetting != nil {
		if err := s.CustomPrometheusSetting.Validate(); err != nil {
			return err
		}
	}
	if s.CustomVPCSetting != nil {
		if err := s.CustomVPCSetting.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateTimingSyntheticTaskRequestCommonSettingCustomHost struct {
	// The list of hosts.
	Hosts []*UpdateTimingSyntheticTaskRequestCommonSettingCustomHostHosts `json:"Hosts,omitempty" xml:"Hosts,omitempty" type:"Repeated"`
	// The selection mode. Valid values:
	//
	// 	- 0: random
	//
	// 	- 1: polling
	//
	// example:
	//
	// 0
	SelectType *int32 `json:"SelectType,omitempty" xml:"SelectType,omitempty"`
}

func (s UpdateTimingSyntheticTaskRequestCommonSettingCustomHost) String() string {
	return dara.Prettify(s)
}

func (s UpdateTimingSyntheticTaskRequestCommonSettingCustomHost) GoString() string {
	return s.String()
}

func (s *UpdateTimingSyntheticTaskRequestCommonSettingCustomHost) GetHosts() []*UpdateTimingSyntheticTaskRequestCommonSettingCustomHostHosts {
	return s.Hosts
}

func (s *UpdateTimingSyntheticTaskRequestCommonSettingCustomHost) GetSelectType() *int32 {
	return s.SelectType
}

func (s *UpdateTimingSyntheticTaskRequestCommonSettingCustomHost) SetHosts(v []*UpdateTimingSyntheticTaskRequestCommonSettingCustomHostHosts) *UpdateTimingSyntheticTaskRequestCommonSettingCustomHost {
	s.Hosts = v
	return s
}

func (s *UpdateTimingSyntheticTaskRequestCommonSettingCustomHost) SetSelectType(v int32) *UpdateTimingSyntheticTaskRequestCommonSettingCustomHost {
	s.SelectType = &v
	return s
}

func (s *UpdateTimingSyntheticTaskRequestCommonSettingCustomHost) Validate() error {
	if s.Hosts != nil {
		for _, item := range s.Hosts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateTimingSyntheticTaskRequestCommonSettingCustomHostHosts struct {
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

func (s UpdateTimingSyntheticTaskRequestCommonSettingCustomHostHosts) String() string {
	return dara.Prettify(s)
}

func (s UpdateTimingSyntheticTaskRequestCommonSettingCustomHostHosts) GoString() string {
	return s.String()
}

func (s *UpdateTimingSyntheticTaskRequestCommonSettingCustomHostHosts) GetDomain() *string {
	return s.Domain
}

func (s *UpdateTimingSyntheticTaskRequestCommonSettingCustomHostHosts) GetIpType() *int32 {
	return s.IpType
}

func (s *UpdateTimingSyntheticTaskRequestCommonSettingCustomHostHosts) GetIps() []*string {
	return s.Ips
}

func (s *UpdateTimingSyntheticTaskRequestCommonSettingCustomHostHosts) SetDomain(v string) *UpdateTimingSyntheticTaskRequestCommonSettingCustomHostHosts {
	s.Domain = &v
	return s
}

func (s *UpdateTimingSyntheticTaskRequestCommonSettingCustomHostHosts) SetIpType(v int32) *UpdateTimingSyntheticTaskRequestCommonSettingCustomHostHosts {
	s.IpType = &v
	return s
}

func (s *UpdateTimingSyntheticTaskRequestCommonSettingCustomHostHosts) SetIps(v []*string) *UpdateTimingSyntheticTaskRequestCommonSettingCustomHostHosts {
	s.Ips = v
	return s
}

func (s *UpdateTimingSyntheticTaskRequestCommonSettingCustomHostHosts) Validate() error {
	return dara.Validate(s)
}

type UpdateTimingSyntheticTaskRequestCommonSettingCustomPrometheusSetting struct {
	// A reserved parameter.
	//
	// example:
	//
	// A reserved parameter.
	PrometheusClusterId *string `json:"PrometheusClusterId,omitempty" xml:"PrometheusClusterId,omitempty"`
	// A reserved parameter.
	//
	// example:
	//
	// A reserved parameter.
	PrometheusClusterRegion *string `json:"PrometheusClusterRegion,omitempty" xml:"PrometheusClusterRegion,omitempty"`
	// The reserved parameters.
	PrometheusLabels map[string]*string `json:"PrometheusLabels,omitempty" xml:"PrometheusLabels,omitempty"`
}

func (s UpdateTimingSyntheticTaskRequestCommonSettingCustomPrometheusSetting) String() string {
	return dara.Prettify(s)
}

func (s UpdateTimingSyntheticTaskRequestCommonSettingCustomPrometheusSetting) GoString() string {
	return s.String()
}

func (s *UpdateTimingSyntheticTaskRequestCommonSettingCustomPrometheusSetting) GetPrometheusClusterId() *string {
	return s.PrometheusClusterId
}

func (s *UpdateTimingSyntheticTaskRequestCommonSettingCustomPrometheusSetting) GetPrometheusClusterRegion() *string {
	return s.PrometheusClusterRegion
}

func (s *UpdateTimingSyntheticTaskRequestCommonSettingCustomPrometheusSetting) GetPrometheusLabels() map[string]*string {
	return s.PrometheusLabels
}

func (s *UpdateTimingSyntheticTaskRequestCommonSettingCustomPrometheusSetting) SetPrometheusClusterId(v string) *UpdateTimingSyntheticTaskRequestCommonSettingCustomPrometheusSetting {
	s.PrometheusClusterId = &v
	return s
}

func (s *UpdateTimingSyntheticTaskRequestCommonSettingCustomPrometheusSetting) SetPrometheusClusterRegion(v string) *UpdateTimingSyntheticTaskRequestCommonSettingCustomPrometheusSetting {
	s.PrometheusClusterRegion = &v
	return s
}

func (s *UpdateTimingSyntheticTaskRequestCommonSettingCustomPrometheusSetting) SetPrometheusLabels(v map[string]*string) *UpdateTimingSyntheticTaskRequestCommonSettingCustomPrometheusSetting {
	s.PrometheusLabels = v
	return s
}

func (s *UpdateTimingSyntheticTaskRequestCommonSettingCustomPrometheusSetting) Validate() error {
	return dara.Validate(s)
}

type UpdateTimingSyntheticTaskRequestCommonSettingCustomVPCSetting struct {
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the security group to which the client belongs. The security group specifies the inbound and outbound rules of the client for the VPC. You need to allow the security group to which the client belongs to access the security group to which the VPC belongs. Otherwise, the client cannot access resources in the VPC.
	//
	// example:
	//
	// sg-xxxxxxxx
	SecureGroupId *string `json:"SecureGroupId,omitempty" xml:"SecureGroupId,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-2zevek6r3mpny7wxxxxxv
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// VPC ID.
	//
	// example:
	//
	// vpc-bp15bjtubjytclwxxxxxx
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s UpdateTimingSyntheticTaskRequestCommonSettingCustomVPCSetting) String() string {
	return dara.Prettify(s)
}

func (s UpdateTimingSyntheticTaskRequestCommonSettingCustomVPCSetting) GoString() string {
	return s.String()
}

func (s *UpdateTimingSyntheticTaskRequestCommonSettingCustomVPCSetting) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateTimingSyntheticTaskRequestCommonSettingCustomVPCSetting) GetSecureGroupId() *string {
	return s.SecureGroupId
}

func (s *UpdateTimingSyntheticTaskRequestCommonSettingCustomVPCSetting) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *UpdateTimingSyntheticTaskRequestCommonSettingCustomVPCSetting) GetVpcId() *string {
	return s.VpcId
}

func (s *UpdateTimingSyntheticTaskRequestCommonSettingCustomVPCSetting) SetRegionId(v string) *UpdateTimingSyntheticTaskRequestCommonSettingCustomVPCSetting {
	s.RegionId = &v
	return s
}

func (s *UpdateTimingSyntheticTaskRequestCommonSettingCustomVPCSetting) SetSecureGroupId(v string) *UpdateTimingSyntheticTaskRequestCommonSettingCustomVPCSetting {
	s.SecureGroupId = &v
	return s
}

func (s *UpdateTimingSyntheticTaskRequestCommonSettingCustomVPCSetting) SetVSwitchId(v string) *UpdateTimingSyntheticTaskRequestCommonSettingCustomVPCSetting {
	s.VSwitchId = &v
	return s
}

func (s *UpdateTimingSyntheticTaskRequestCommonSettingCustomVPCSetting) SetVpcId(v string) *UpdateTimingSyntheticTaskRequestCommonSettingCustomVPCSetting {
	s.VpcId = &v
	return s
}

func (s *UpdateTimingSyntheticTaskRequestCommonSettingCustomVPCSetting) Validate() error {
	return dara.Validate(s)
}

type UpdateTimingSyntheticTaskRequestCustomPeriod struct {
	// The hour at which the test ends. Valid values: 0 to 24.
	//
	// example:
	//
	// 22
	EndHour *int32 `json:"EndHour,omitempty" xml:"EndHour,omitempty"`
	// The hour at which the test starts. Valid values: 0 to 24.
	//
	// example:
	//
	// 14
	StartHour *int32 `json:"StartHour,omitempty" xml:"StartHour,omitempty"`
}

func (s UpdateTimingSyntheticTaskRequestCustomPeriod) String() string {
	return dara.Prettify(s)
}

func (s UpdateTimingSyntheticTaskRequestCustomPeriod) GoString() string {
	return s.String()
}

func (s *UpdateTimingSyntheticTaskRequestCustomPeriod) GetEndHour() *int32 {
	return s.EndHour
}

func (s *UpdateTimingSyntheticTaskRequestCustomPeriod) GetStartHour() *int32 {
	return s.StartHour
}

func (s *UpdateTimingSyntheticTaskRequestCustomPeriod) SetEndHour(v int32) *UpdateTimingSyntheticTaskRequestCustomPeriod {
	s.EndHour = &v
	return s
}

func (s *UpdateTimingSyntheticTaskRequestCustomPeriod) SetStartHour(v int32) *UpdateTimingSyntheticTaskRequestCustomPeriod {
	s.StartHour = &v
	return s
}

func (s *UpdateTimingSyntheticTaskRequestCustomPeriod) Validate() error {
	return dara.Validate(s)
}

type UpdateTimingSyntheticTaskRequestMonitorConf struct {
	// The parameters of the HTTP(S) synthetic test.
	ApiHTTP *UpdateTimingSyntheticTaskRequestMonitorConfApiHTTP `json:"ApiHTTP,omitempty" xml:"ApiHTTP,omitempty" type:"Struct"`
	// The parameters of file downloading.
	FileDownload *UpdateTimingSyntheticTaskRequestMonitorConfFileDownload `json:"FileDownload,omitempty" xml:"FileDownload,omitempty" type:"Struct"`
	// The parameters of the DNS synthetic test.
	NetDNS *UpdateTimingSyntheticTaskRequestMonitorConfNetDNS `json:"NetDNS,omitempty" xml:"NetDNS,omitempty" type:"Struct"`
	// The parameters of the ICMP synthetic test.
	NetICMP *UpdateTimingSyntheticTaskRequestMonitorConfNetICMP `json:"NetICMP,omitempty" xml:"NetICMP,omitempty" type:"Struct"`
	// The parameters of the TCP synthetic test.
	NetTCP *UpdateTimingSyntheticTaskRequestMonitorConfNetTCP `json:"NetTCP,omitempty" xml:"NetTCP,omitempty" type:"Struct"`
	// The parameters of the streaming-media synthetic test.
	Stream *UpdateTimingSyntheticTaskRequestMonitorConfStream `json:"Stream,omitempty" xml:"Stream,omitempty" type:"Struct"`
	// The parameters of the website speed measurement.
	Website *UpdateTimingSyntheticTaskRequestMonitorConfWebsite `json:"Website,omitempty" xml:"Website,omitempty" type:"Struct"`
}

func (s UpdateTimingSyntheticTaskRequestMonitorConf) String() string {
	return dara.Prettify(s)
}

func (s UpdateTimingSyntheticTaskRequestMonitorConf) GoString() string {
	return s.String()
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConf) GetApiHTTP() *UpdateTimingSyntheticTaskRequestMonitorConfApiHTTP {
	return s.ApiHTTP
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConf) GetFileDownload() *UpdateTimingSyntheticTaskRequestMonitorConfFileDownload {
	return s.FileDownload
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConf) GetNetDNS() *UpdateTimingSyntheticTaskRequestMonitorConfNetDNS {
	return s.NetDNS
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConf) GetNetICMP() *UpdateTimingSyntheticTaskRequestMonitorConfNetICMP {
	return s.NetICMP
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConf) GetNetTCP() *UpdateTimingSyntheticTaskRequestMonitorConfNetTCP {
	return s.NetTCP
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConf) GetStream() *UpdateTimingSyntheticTaskRequestMonitorConfStream {
	return s.Stream
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConf) GetWebsite() *UpdateTimingSyntheticTaskRequestMonitorConfWebsite {
	return s.Website
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConf) SetApiHTTP(v *UpdateTimingSyntheticTaskRequestMonitorConfApiHTTP) *UpdateTimingSyntheticTaskRequestMonitorConf {
	s.ApiHTTP = v
	return s
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConf) SetFileDownload(v *UpdateTimingSyntheticTaskRequestMonitorConfFileDownload) *UpdateTimingSyntheticTaskRequestMonitorConf {
	s.FileDownload = v
	return s
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConf) SetNetDNS(v *UpdateTimingSyntheticTaskRequestMonitorConfNetDNS) *UpdateTimingSyntheticTaskRequestMonitorConf {
	s.NetDNS = v
	return s
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConf) SetNetICMP(v *UpdateTimingSyntheticTaskRequestMonitorConfNetICMP) *UpdateTimingSyntheticTaskRequestMonitorConf {
	s.NetICMP = v
	return s
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConf) SetNetTCP(v *UpdateTimingSyntheticTaskRequestMonitorConfNetTCP) *UpdateTimingSyntheticTaskRequestMonitorConf {
	s.NetTCP = v
	return s
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConf) SetStream(v *UpdateTimingSyntheticTaskRequestMonitorConfStream) *UpdateTimingSyntheticTaskRequestMonitorConf {
	s.Stream = v
	return s
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConf) SetWebsite(v *UpdateTimingSyntheticTaskRequestMonitorConfWebsite) *UpdateTimingSyntheticTaskRequestMonitorConf {
	s.Website = v
	return s
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConf) Validate() error {
	if s.ApiHTTP != nil {
		if err := s.ApiHTTP.Validate(); err != nil {
			return err
		}
	}
	if s.FileDownload != nil {
		if err := s.FileDownload.Validate(); err != nil {
			return err
		}
	}
	if s.NetDNS != nil {
		if err := s.NetDNS.Validate(); err != nil {
			return err
		}
	}
	if s.NetICMP != nil {
		if err := s.NetICMP.Validate(); err != nil {
			return err
		}
	}
	if s.NetTCP != nil {
		if err := s.NetTCP.Validate(); err != nil {
			return err
		}
	}
	if s.Stream != nil {
		if err := s.Stream.Validate(); err != nil {
			return err
		}
	}
	if s.Website != nil {
		if err := s.Website.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateTimingSyntheticTaskRequestMonitorConfApiHTTP struct {
	// Specifies whether to verify the certificate. Default value: false.
	//
	// example:
	//
	// true
	CheckCert *bool `json:"CheckCert,omitempty" xml:"CheckCert,omitempty"`
	// The connection timeout period. Unit: milliseconds. Default value: 5000. Minimum value: 1000. Maximum value: 300000.
	//
	// example:
	//
	// 5000
	ConnectTimeout *int64 `json:"ConnectTimeout,omitempty" xml:"ConnectTimeout,omitempty"`
	// The request method. Valid values:
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
	// 1-http/1.1
	//
	// 2-h2
	//
	// 3: disables the ALPN protocol
	//
	// example:
	//
	// 0
	ProtocolAlpnProtocol *int32 `json:"ProtocolAlpnProtocol,omitempty" xml:"ProtocolAlpnProtocol,omitempty"`
	// The HTTP request body.
	RequestBody *UpdateTimingSyntheticTaskRequestMonitorConfApiHTTPRequestBody `json:"RequestBody,omitempty" xml:"RequestBody,omitempty" type:"Struct"`
	// The custom header field.
	RequestHeaders map[string]*string `json:"RequestHeaders,omitempty" xml:"RequestHeaders,omitempty"`
	// The URL or request path for synthetic monitoring.
	//
	// example:
	//
	// https://********
	TargetUrl *string `json:"TargetUrl,omitempty" xml:"TargetUrl,omitempty"`
	// The timeout period. Unit: milliseconds. Default value: 10000. Minimum value: 1000. Maximum value: 300000.
	//
	// example:
	//
	// 5000
	Timeout *int64 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s UpdateTimingSyntheticTaskRequestMonitorConfApiHTTP) String() string {
	return dara.Prettify(s)
}

func (s UpdateTimingSyntheticTaskRequestMonitorConfApiHTTP) GoString() string {
	return s.String()
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfApiHTTP) GetCheckCert() *bool {
	return s.CheckCert
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfApiHTTP) GetConnectTimeout() *int64 {
	return s.ConnectTimeout
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfApiHTTP) GetMethod() *string {
	return s.Method
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfApiHTTP) GetProtocolAlpnProtocol() *int32 {
	return s.ProtocolAlpnProtocol
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfApiHTTP) GetRequestBody() *UpdateTimingSyntheticTaskRequestMonitorConfApiHTTPRequestBody {
	return s.RequestBody
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfApiHTTP) GetRequestHeaders() map[string]*string {
	return s.RequestHeaders
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfApiHTTP) GetTargetUrl() *string {
	return s.TargetUrl
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfApiHTTP) GetTimeout() *int64 {
	return s.Timeout
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfApiHTTP) SetCheckCert(v bool) *UpdateTimingSyntheticTaskRequestMonitorConfApiHTTP {
	s.CheckCert = &v
	return s
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfApiHTTP) SetConnectTimeout(v int64) *UpdateTimingSyntheticTaskRequestMonitorConfApiHTTP {
	s.ConnectTimeout = &v
	return s
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfApiHTTP) SetMethod(v string) *UpdateTimingSyntheticTaskRequestMonitorConfApiHTTP {
	s.Method = &v
	return s
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfApiHTTP) SetProtocolAlpnProtocol(v int32) *UpdateTimingSyntheticTaskRequestMonitorConfApiHTTP {
	s.ProtocolAlpnProtocol = &v
	return s
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfApiHTTP) SetRequestBody(v *UpdateTimingSyntheticTaskRequestMonitorConfApiHTTPRequestBody) *UpdateTimingSyntheticTaskRequestMonitorConfApiHTTP {
	s.RequestBody = v
	return s
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfApiHTTP) SetRequestHeaders(v map[string]*string) *UpdateTimingSyntheticTaskRequestMonitorConfApiHTTP {
	s.RequestHeaders = v
	return s
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfApiHTTP) SetTargetUrl(v string) *UpdateTimingSyntheticTaskRequestMonitorConfApiHTTP {
	s.TargetUrl = &v
	return s
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfApiHTTP) SetTimeout(v int64) *UpdateTimingSyntheticTaskRequestMonitorConfApiHTTP {
	s.Timeout = &v
	return s
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfApiHTTP) Validate() error {
	if s.RequestBody != nil {
		if err := s.RequestBody.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateTimingSyntheticTaskRequestMonitorConfApiHTTPRequestBody struct {
	// The content of the request body. Format: JSON string. The parameter is required if the Type parameter is set to text/plain, application/json, application/xml, or text/html. Format: JSON string.
	//
	// example:
	//
	// {
	//
	//       "key1": "value1",
	//
	//       "key2": "value2"
	//
	// }
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The type of the request body. Valid values: text/plain, application/json, application/x-www-form-urlencoded, multipart/form-data, application/xml, and text/html.
	//
	// example:
	//
	// application/json
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateTimingSyntheticTaskRequestMonitorConfApiHTTPRequestBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateTimingSyntheticTaskRequestMonitorConfApiHTTPRequestBody) GoString() string {
	return s.String()
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfApiHTTPRequestBody) GetContent() *string {
	return s.Content
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfApiHTTPRequestBody) GetType() *string {
	return s.Type
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfApiHTTPRequestBody) SetContent(v string) *UpdateTimingSyntheticTaskRequestMonitorConfApiHTTPRequestBody {
	s.Content = &v
	return s
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfApiHTTPRequestBody) SetType(v string) *UpdateTimingSyntheticTaskRequestMonitorConfApiHTTPRequestBody {
	s.Type = &v
	return s
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfApiHTTPRequestBody) Validate() error {
	return dara.Validate(s)
}

type UpdateTimingSyntheticTaskRequestMonitorConfFileDownload struct {
	// Unit: milliseconds. Minimum value: 1000. Maximum value: 120000. Default value: 5000.
	//
	// example:
	//
	// 5000
	ConnectionTimeout *int64 `json:"ConnectionTimeout,omitempty" xml:"ConnectionTimeout,omitempty"`
	// The content of the custom request header. Format: JSON map.
	CustomHeaderContent map[string]*string `json:"CustomHeaderContent,omitempty" xml:"CustomHeaderContent,omitempty"`
	// The kernel type. Valid values:
	//
	// 	- 1: curl
	//
	// 	- 0: WinInet
	//
	// example:
	//
	// 0
	DownloadKernel *int32 `json:"DownloadKernel,omitempty" xml:"DownloadKernel,omitempty"`
	// Specifies whether to ignore CA certificate authentication errors. Valid values: 0: no. 1: yes. Default value: 1.
	//
	// example:
	//
	// 0
	IgnoreCertificateAuthError *int32 `json:"IgnoreCertificateAuthError,omitempty" xml:"IgnoreCertificateAuthError,omitempty"`
	// Specifies whether to ignore certificate revocation errors. Valid values: 0: no. 1: yes. Default value: 1.
	//
	// example:
	//
	// 0
	IgnoreCertificateCanceledError *int32 `json:"IgnoreCertificateCanceledError,omitempty" xml:"IgnoreCertificateCanceledError,omitempty"`
	// Specifies whether to ignore certificate invalidity. Valid values: 0: no. 1: yes. Default value: 1.
	//
	// example:
	//
	// 0
	IgnoreCertificateOutOfDateError *int32 `json:"IgnoreCertificateOutOfDateError,omitempty" xml:"IgnoreCertificateOutOfDateError,omitempty"`
	// Specifies whether to ignore certificate status errors. 0: no. 1: yes. Default value: 1.
	//
	// example:
	//
	// 0
	IgnoreCertificateStatusError *int32 `json:"IgnoreCertificateStatusError,omitempty" xml:"IgnoreCertificateStatusError,omitempty"`
	// Specifies whether to ignore certificate incredibility. Valid values: 0: no. 1: yes. Default value: 1.
	//
	// example:
	//
	// 0
	IgnoreCertificateUntrustworthyError *int32 `json:"IgnoreCertificateUntrustworthyError,omitempty" xml:"IgnoreCertificateUntrustworthyError,omitempty"`
	// Specifies whether to ignore certificate usage errors. Valid values: 0: no. 1: yes. Default value: 1.
	//
	// example:
	//
	// 0
	IgnoreCertificateUsingError *int32 `json:"IgnoreCertificateUsingError,omitempty" xml:"IgnoreCertificateUsingError,omitempty"`
	// Specifies whether to ignore host invalidity. Valid values: 0: no. 1: yes. Default value: 1.
	//
	// example:
	//
	// 0
	IgnoreInvalidHostError *int32 `json:"IgnoreInvalidHostError,omitempty" xml:"IgnoreInvalidHostError,omitempty"`
	// The monitoring timeout period. Unit: milliseconds. Minimum value: 1000. Maximum value: 120000. Default value: 60000.
	//
	// example:
	//
	// 60000
	MonitorTimeout *int64 `json:"MonitorTimeout,omitempty" xml:"MonitorTimeout,omitempty"`
	// The QUIC protocol type. Valid values:
	//
	// 	- 1: HTTP/1
	//
	// 	- 2: HTTP/2
	//
	// 	- 3: http3
	//
	// example:
	//
	// 1
	QuickProtocol *int32 `json:"QuickProtocol,omitempty" xml:"QuickProtocol,omitempty"`
	// Specifies whether to support redirection. Valid values: 0: no. 1: yes. Default value: 1.
	//
	// example:
	//
	// 0
	Redirection *int32 `json:"Redirection,omitempty" xml:"Redirection,omitempty"`
	// The URL that is used to download the file.
	//
	// example:
	//
	// https://img.alicdn.com/tfs/TB13DzOjXP7gK0jSZFjXXc5aXXa-212-48.png
	TargetUrl *string `json:"TargetUrl,omitempty" xml:"TargetUrl,omitempty"`
	// The maximum file size of a single transfer. Unit: KB. Minimum value: 1. Maximum value: 20480. Valid values: 2048.
	//
	// example:
	//
	// 2048
	TransmissionSize *int64 `json:"TransmissionSize,omitempty" xml:"TransmissionSize,omitempty"`
	// The keyword that is used in verification.
	//
	// example:
	//
	// aliyun
	ValidateKeywords *string `json:"ValidateKeywords,omitempty" xml:"ValidateKeywords,omitempty"`
	// The verification method. Valid values:
	//
	// 	- 0: no verification
	//
	// 	- 1: string verification
	//
	// 	- 2: MD5 verification
	//
	// example:
	//
	// 0
	VerifyWay *int32 `json:"VerifyWay,omitempty" xml:"VerifyWay,omitempty"`
	// The whitelisted objects that are used to avoid DNS hijacking. The objects can be IP addresses, wildcard mask, subnet mask, or CNAME records. Separate multiple objects with vertical bars (|). Example: www.aliyun.com:203.0.3.55|203.3.44.67. It indicates that all IP addresses that belong to the www.aliyun.com domain name except 203.0.3.55 and 203.3.44.67 are hijacked.
	//
	// example:
	//
	// www.aliyun.com:203.0.3.55|203.3.44.67
	WhiteList *string `json:"WhiteList,omitempty" xml:"WhiteList,omitempty"`
}

func (s UpdateTimingSyntheticTaskRequestMonitorConfFileDownload) String() string {
	return dara.Prettify(s)
}

func (s UpdateTimingSyntheticTaskRequestMonitorConfFileDownload) GoString() string {
	return s.String()
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfFileDownload) GetConnectionTimeout() *int64 {
	return s.ConnectionTimeout
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfFileDownload) GetCustomHeaderContent() map[string]*string {
	return s.CustomHeaderContent
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfFileDownload) GetDownloadKernel() *int32 {
	return s.DownloadKernel
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfFileDownload) GetIgnoreCertificateAuthError() *int32 {
	return s.IgnoreCertificateAuthError
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfFileDownload) GetIgnoreCertificateCanceledError() *int32 {
	return s.IgnoreCertificateCanceledError
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfFileDownload) GetIgnoreCertificateOutOfDateError() *int32 {
	return s.IgnoreCertificateOutOfDateError
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfFileDownload) GetIgnoreCertificateStatusError() *int32 {
	return s.IgnoreCertificateStatusError
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfFileDownload) GetIgnoreCertificateUntrustworthyError() *int32 {
	return s.IgnoreCertificateUntrustworthyError
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfFileDownload) GetIgnoreCertificateUsingError() *int32 {
	return s.IgnoreCertificateUsingError
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfFileDownload) GetIgnoreInvalidHostError() *int32 {
	return s.IgnoreInvalidHostError
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfFileDownload) GetMonitorTimeout() *int64 {
	return s.MonitorTimeout
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfFileDownload) GetQuickProtocol() *int32 {
	return s.QuickProtocol
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfFileDownload) GetRedirection() *int32 {
	return s.Redirection
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfFileDownload) GetTargetUrl() *string {
	return s.TargetUrl
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfFileDownload) GetTransmissionSize() *int64 {
	return s.TransmissionSize
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfFileDownload) GetValidateKeywords() *string {
	return s.ValidateKeywords
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfFileDownload) GetVerifyWay() *int32 {
	return s.VerifyWay
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfFileDownload) GetWhiteList() *string {
	return s.WhiteList
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfFileDownload) SetConnectionTimeout(v int64) *UpdateTimingSyntheticTaskRequestMonitorConfFileDownload {
	s.ConnectionTimeout = &v
	return s
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfFileDownload) SetCustomHeaderContent(v map[string]*string) *UpdateTimingSyntheticTaskRequestMonitorConfFileDownload {
	s.CustomHeaderContent = v
	return s
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfFileDownload) SetDownloadKernel(v int32) *UpdateTimingSyntheticTaskRequestMonitorConfFileDownload {
	s.DownloadKernel = &v
	return s
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfFileDownload) SetIgnoreCertificateAuthError(v int32) *UpdateTimingSyntheticTaskRequestMonitorConfFileDownload {
	s.IgnoreCertificateAuthError = &v
	return s
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfFileDownload) SetIgnoreCertificateCanceledError(v int32) *UpdateTimingSyntheticTaskRequestMonitorConfFileDownload {
	s.IgnoreCertificateCanceledError = &v
	return s
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfFileDownload) SetIgnoreCertificateOutOfDateError(v int32) *UpdateTimingSyntheticTaskRequestMonitorConfFileDownload {
	s.IgnoreCertificateOutOfDateError = &v
	return s
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfFileDownload) SetIgnoreCertificateStatusError(v int32) *UpdateTimingSyntheticTaskRequestMonitorConfFileDownload {
	s.IgnoreCertificateStatusError = &v
	return s
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfFileDownload) SetIgnoreCertificateUntrustworthyError(v int32) *UpdateTimingSyntheticTaskRequestMonitorConfFileDownload {
	s.IgnoreCertificateUntrustworthyError = &v
	return s
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfFileDownload) SetIgnoreCertificateUsingError(v int32) *UpdateTimingSyntheticTaskRequestMonitorConfFileDownload {
	s.IgnoreCertificateUsingError = &v
	return s
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfFileDownload) SetIgnoreInvalidHostError(v int32) *UpdateTimingSyntheticTaskRequestMonitorConfFileDownload {
	s.IgnoreInvalidHostError = &v
	return s
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfFileDownload) SetMonitorTimeout(v int64) *UpdateTimingSyntheticTaskRequestMonitorConfFileDownload {
	s.MonitorTimeout = &v
	return s
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfFileDownload) SetQuickProtocol(v int32) *UpdateTimingSyntheticTaskRequestMonitorConfFileDownload {
	s.QuickProtocol = &v
	return s
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfFileDownload) SetRedirection(v int32) *UpdateTimingSyntheticTaskRequestMonitorConfFileDownload {
	s.Redirection = &v
	return s
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfFileDownload) SetTargetUrl(v string) *UpdateTimingSyntheticTaskRequestMonitorConfFileDownload {
	s.TargetUrl = &v
	return s
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfFileDownload) SetTransmissionSize(v int64) *UpdateTimingSyntheticTaskRequestMonitorConfFileDownload {
	s.TransmissionSize = &v
	return s
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfFileDownload) SetValidateKeywords(v string) *UpdateTimingSyntheticTaskRequestMonitorConfFileDownload {
	s.ValidateKeywords = &v
	return s
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfFileDownload) SetVerifyWay(v int32) *UpdateTimingSyntheticTaskRequestMonitorConfFileDownload {
	s.VerifyWay = &v
	return s
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfFileDownload) SetWhiteList(v string) *UpdateTimingSyntheticTaskRequestMonitorConfFileDownload {
	s.WhiteList = &v
	return s
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfFileDownload) Validate() error {
	return dara.Validate(s)
}

type UpdateTimingSyntheticTaskRequestMonitorConfNetDNS struct {
	// Specifies whether to use the dig command to display the data. Valid values: 0: no. 1: yes.
	//
	// example:
	//
	// 0
	Dig *int32 `json:"Dig,omitempty" xml:"Dig,omitempty"`
	// The IP version of the DNS server. Valid values: 0: IPv4. 1: IPv6. 2: A version is automatically selected. Default value: 0.
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
	// The DNS query method. Valid values: 0: recursive. 1: iterative. Default value: 0.
	//
	// example:
	//
	// 0
	QueryMethod *int32 `json:"QueryMethod,omitempty" xml:"QueryMethod,omitempty"`
	// The domain name.
	//
	// example:
	//
	// www.aliyun.com
	TargetUrl *string `json:"TargetUrl,omitempty" xml:"TargetUrl,omitempty"`
	// The timeout period for the DNS synthetic test. Unit: milliseconds. Minimum value: 1000. Maximum value: 45000. Default value: 5000.
	//
	// example:
	//
	// 1000
	Timeout *int64 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s UpdateTimingSyntheticTaskRequestMonitorConfNetDNS) String() string {
	return dara.Prettify(s)
}

func (s UpdateTimingSyntheticTaskRequestMonitorConfNetDNS) GoString() string {
	return s.String()
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfNetDNS) GetDig() *int32 {
	return s.Dig
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfNetDNS) GetDnsServerIpType() *int32 {
	return s.DnsServerIpType
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfNetDNS) GetNsServer() *string {
	return s.NsServer
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfNetDNS) GetQueryMethod() *int32 {
	return s.QueryMethod
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfNetDNS) GetTargetUrl() *string {
	return s.TargetUrl
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfNetDNS) GetTimeout() *int64 {
	return s.Timeout
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfNetDNS) SetDig(v int32) *UpdateTimingSyntheticTaskRequestMonitorConfNetDNS {
	s.Dig = &v
	return s
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfNetDNS) SetDnsServerIpType(v int32) *UpdateTimingSyntheticTaskRequestMonitorConfNetDNS {
	s.DnsServerIpType = &v
	return s
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfNetDNS) SetNsServer(v string) *UpdateTimingSyntheticTaskRequestMonitorConfNetDNS {
	s.NsServer = &v
	return s
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfNetDNS) SetQueryMethod(v int32) *UpdateTimingSyntheticTaskRequestMonitorConfNetDNS {
	s.QueryMethod = &v
	return s
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfNetDNS) SetTargetUrl(v string) *UpdateTimingSyntheticTaskRequestMonitorConfNetDNS {
	s.TargetUrl = &v
	return s
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfNetDNS) SetTimeout(v int64) *UpdateTimingSyntheticTaskRequestMonitorConfNetDNS {
	s.Timeout = &v
	return s
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfNetDNS) Validate() error {
	return dara.Validate(s)
}

type UpdateTimingSyntheticTaskRequestMonitorConfNetICMP struct {
	// The interval at which ICMP packets are sent. Unit: milliseconds. Minimum value: 200. Maximum value: 10000.
	//
	// example:
	//
	// 300
	Interval *int64 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The number of ICMP packets that are sent. Minimum value: 1. Maximum value: 50. Default value: 4.
	//
	// example:
	//
	// 4
	PackageNum *int32 `json:"PackageNum,omitempty" xml:"PackageNum,omitempty"`
	// The size of each ICMP packet. Unit: bytes. Valid values: 32, 64, 128, 256, 512, 1024, 1080, and 1450.
	//
	// example:
	//
	// 32
	PackageSize *int32 `json:"PackageSize,omitempty" xml:"PackageSize,omitempty"`
	// Specifies whether to split ICMP packets. Default value: true.
	//
	// example:
	//
	// true
	SplitPackage *bool `json:"SplitPackage,omitempty" xml:"SplitPackage,omitempty"`
	// The IP address or domain name of the destination host. The value cannot contain port numbers, protocol headers, or request paths.
	//
	// example:
	//
	// www.aliyun.com
	TargetUrl *string `json:"TargetUrl,omitempty" xml:"TargetUrl,omitempty"`
	// The timeout period for the ICMP synthetic test. Unit: milliseconds. Minimum value: 1000. Maximum value: 300000. Default value: 20000.
	//
	// example:
	//
	// 5000
	Timeout *int64 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// Specifies whether to enable the tracert command. Default value: true.
	//
	// example:
	//
	// true
	TracertEnable *bool `json:"TracertEnable,omitempty" xml:"TracertEnable,omitempty"`
	// The maximum number of hops for the tracert command. Minimum value: 1. Maximum value: 128. Default value: 20.
	//
	// example:
	//
	// 20
	TracertNumMax *int32 `json:"TracertNumMax,omitempty" xml:"TracertNumMax,omitempty"`
	// The timeout period of the tracert command. Unit: milliseconds. Minimum value: 1000. Maximum value: 300000. Default value: 60000.
	//
	// example:
	//
	// 60000
	TracertTimeout *int64 `json:"TracertTimeout,omitempty" xml:"TracertTimeout,omitempty"`
}

func (s UpdateTimingSyntheticTaskRequestMonitorConfNetICMP) String() string {
	return dara.Prettify(s)
}

func (s UpdateTimingSyntheticTaskRequestMonitorConfNetICMP) GoString() string {
	return s.String()
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfNetICMP) GetInterval() *int64 {
	return s.Interval
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfNetICMP) GetPackageNum() *int32 {
	return s.PackageNum
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfNetICMP) GetPackageSize() *int32 {
	return s.PackageSize
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfNetICMP) GetSplitPackage() *bool {
	return s.SplitPackage
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfNetICMP) GetTargetUrl() *string {
	return s.TargetUrl
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfNetICMP) GetTimeout() *int64 {
	return s.Timeout
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfNetICMP) GetTracertEnable() *bool {
	return s.TracertEnable
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfNetICMP) GetTracertNumMax() *int32 {
	return s.TracertNumMax
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfNetICMP) GetTracertTimeout() *int64 {
	return s.TracertTimeout
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfNetICMP) SetInterval(v int64) *UpdateTimingSyntheticTaskRequestMonitorConfNetICMP {
	s.Interval = &v
	return s
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfNetICMP) SetPackageNum(v int32) *UpdateTimingSyntheticTaskRequestMonitorConfNetICMP {
	s.PackageNum = &v
	return s
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfNetICMP) SetPackageSize(v int32) *UpdateTimingSyntheticTaskRequestMonitorConfNetICMP {
	s.PackageSize = &v
	return s
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfNetICMP) SetSplitPackage(v bool) *UpdateTimingSyntheticTaskRequestMonitorConfNetICMP {
	s.SplitPackage = &v
	return s
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfNetICMP) SetTargetUrl(v string) *UpdateTimingSyntheticTaskRequestMonitorConfNetICMP {
	s.TargetUrl = &v
	return s
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfNetICMP) SetTimeout(v int64) *UpdateTimingSyntheticTaskRequestMonitorConfNetICMP {
	s.Timeout = &v
	return s
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfNetICMP) SetTracertEnable(v bool) *UpdateTimingSyntheticTaskRequestMonitorConfNetICMP {
	s.TracertEnable = &v
	return s
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfNetICMP) SetTracertNumMax(v int32) *UpdateTimingSyntheticTaskRequestMonitorConfNetICMP {
	s.TracertNumMax = &v
	return s
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfNetICMP) SetTracertTimeout(v int64) *UpdateTimingSyntheticTaskRequestMonitorConfNetICMP {
	s.TracertTimeout = &v
	return s
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfNetICMP) Validate() error {
	return dara.Validate(s)
}

type UpdateTimingSyntheticTaskRequestMonitorConfNetTCP struct {
	// The number of TCP connections that are established. Minimum value: 1. Maximum value: 16. Default value: 4.
	//
	// example:
	//
	// 4
	ConnectTimes *int32 `json:"ConnectTimes,omitempty" xml:"ConnectTimes,omitempty"`
	// The interval at which TCP connections are established. Unit: milliseconds. Minimum value: 200. Maximum value: 10000. Default value: 200.
	//
	// example:
	//
	// 300
	Interval *int64 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The IP address of the destination host.
	//
	// example:
	//
	// 127.0.0.1:8888
	TargetUrl *string `json:"TargetUrl,omitempty" xml:"TargetUrl,omitempty"`
	// The timeout period for the TCP synthetic test. Unit: milliseconds. Minimum value: 1000. Maximum value: 300000. Default value: 20000.
	//
	// example:
	//
	// 1000
	Timeout *int64 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// Specifies whether to enable the tracert command. Default value: true.
	//
	// example:
	//
	// true
	TracertEnable *bool `json:"TracertEnable,omitempty" xml:"TracertEnable,omitempty"`
	// The maximum number of hops for the tracert command. Minimum value: 1. Maximum value: 128. Default value: 20.
	//
	// example:
	//
	// 20
	TracertNumMax *int32 `json:"TracertNumMax,omitempty" xml:"TracertNumMax,omitempty"`
	// The timeout period of the tracert command. Unit: milliseconds. Minimum value: 1000. Maximum value: 300000. Default value: 60000.
	//
	// example:
	//
	// 1000
	TracertTimeout *int64 `json:"TracertTimeout,omitempty" xml:"TracertTimeout,omitempty"`
}

func (s UpdateTimingSyntheticTaskRequestMonitorConfNetTCP) String() string {
	return dara.Prettify(s)
}

func (s UpdateTimingSyntheticTaskRequestMonitorConfNetTCP) GoString() string {
	return s.String()
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfNetTCP) GetConnectTimes() *int32 {
	return s.ConnectTimes
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfNetTCP) GetInterval() *int64 {
	return s.Interval
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfNetTCP) GetTargetUrl() *string {
	return s.TargetUrl
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfNetTCP) GetTimeout() *int64 {
	return s.Timeout
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfNetTCP) GetTracertEnable() *bool {
	return s.TracertEnable
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfNetTCP) GetTracertNumMax() *int32 {
	return s.TracertNumMax
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfNetTCP) GetTracertTimeout() *int64 {
	return s.TracertTimeout
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfNetTCP) SetConnectTimes(v int32) *UpdateTimingSyntheticTaskRequestMonitorConfNetTCP {
	s.ConnectTimes = &v
	return s
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfNetTCP) SetInterval(v int64) *UpdateTimingSyntheticTaskRequestMonitorConfNetTCP {
	s.Interval = &v
	return s
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfNetTCP) SetTargetUrl(v string) *UpdateTimingSyntheticTaskRequestMonitorConfNetTCP {
	s.TargetUrl = &v
	return s
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfNetTCP) SetTimeout(v int64) *UpdateTimingSyntheticTaskRequestMonitorConfNetTCP {
	s.Timeout = &v
	return s
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfNetTCP) SetTracertEnable(v bool) *UpdateTimingSyntheticTaskRequestMonitorConfNetTCP {
	s.TracertEnable = &v
	return s
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfNetTCP) SetTracertNumMax(v int32) *UpdateTimingSyntheticTaskRequestMonitorConfNetTCP {
	s.TracertNumMax = &v
	return s
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfNetTCP) SetTracertTimeout(v int64) *UpdateTimingSyntheticTaskRequestMonitorConfNetTCP {
	s.TracertTimeout = &v
	return s
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfNetTCP) Validate() error {
	return dara.Validate(s)
}

type UpdateTimingSyntheticTaskRequestMonitorConfStream struct {
	// The custom header. Format: JSON map.
	CustomHeaderContent map[string]*string `json:"CustomHeaderContent,omitempty" xml:"CustomHeaderContent,omitempty"`
	// The player. Default value: 12. Valid values:
	//
	// 	- 12: VLC
	//
	// 	- 2: Flash Player
	//
	// example:
	//
	// 2
	PlayerType *int32 `json:"PlayerType,omitempty" xml:"PlayerType,omitempty"`
	// The address type of the resource. Valid values:
	//
	// 	- 1: resource URL.
	//
	// 	- 0: page URL. Default value: 0.
	//
	// example:
	//
	// 1
	StreamAddressType *int32 `json:"StreamAddressType,omitempty" xml:"StreamAddressType,omitempty"`
	// The monitoring duration. Unit: seconds. Maximum and default value: 60.
	//
	// example:
	//
	// 30
	StreamMonitorTimeout *int32 `json:"StreamMonitorTimeout,omitempty" xml:"StreamMonitorTimeout,omitempty"`
	// Specifies whether the resource is a video or audio. Valid values: 0: video. 1: audio.
	//
	// example:
	//
	// 0
	StreamType *int32 `json:"StreamType,omitempty" xml:"StreamType,omitempty"`
	// The resource URL of the streaming media.
	//
	// example:
	//
	// http://www.aliyun.com/stream/test.mp4
	TargetUrl *string `json:"TargetUrl,omitempty" xml:"TargetUrl,omitempty"`
	// The whitelisted objects that are used to avoid DNS hijacking. The objects can be IP addresses, wildcard mask, subnet mask, or CNAME records. Separate multiple objects with vertical bars (|). Example: www.aliyun.com:203.0.3.55|203.3.44.67. It indicates that all IP addresses that belong to the www.aliyun.com domain name except 203.0.3.55 and 203.3.44.67 are hijacked.
	//
	// example:
	//
	// www.aliyun.com:203.0.3.55|203.3.44.67
	WhiteList *string `json:"WhiteList,omitempty" xml:"WhiteList,omitempty"`
}

func (s UpdateTimingSyntheticTaskRequestMonitorConfStream) String() string {
	return dara.Prettify(s)
}

func (s UpdateTimingSyntheticTaskRequestMonitorConfStream) GoString() string {
	return s.String()
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfStream) GetCustomHeaderContent() map[string]*string {
	return s.CustomHeaderContent
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfStream) GetPlayerType() *int32 {
	return s.PlayerType
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfStream) GetStreamAddressType() *int32 {
	return s.StreamAddressType
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfStream) GetStreamMonitorTimeout() *int32 {
	return s.StreamMonitorTimeout
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfStream) GetStreamType() *int32 {
	return s.StreamType
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfStream) GetTargetUrl() *string {
	return s.TargetUrl
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfStream) GetWhiteList() *string {
	return s.WhiteList
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfStream) SetCustomHeaderContent(v map[string]*string) *UpdateTimingSyntheticTaskRequestMonitorConfStream {
	s.CustomHeaderContent = v
	return s
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfStream) SetPlayerType(v int32) *UpdateTimingSyntheticTaskRequestMonitorConfStream {
	s.PlayerType = &v
	return s
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfStream) SetStreamAddressType(v int32) *UpdateTimingSyntheticTaskRequestMonitorConfStream {
	s.StreamAddressType = &v
	return s
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfStream) SetStreamMonitorTimeout(v int32) *UpdateTimingSyntheticTaskRequestMonitorConfStream {
	s.StreamMonitorTimeout = &v
	return s
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfStream) SetStreamType(v int32) *UpdateTimingSyntheticTaskRequestMonitorConfStream {
	s.StreamType = &v
	return s
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfStream) SetTargetUrl(v string) *UpdateTimingSyntheticTaskRequestMonitorConfStream {
	s.TargetUrl = &v
	return s
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfStream) SetWhiteList(v string) *UpdateTimingSyntheticTaskRequestMonitorConfStream {
	s.WhiteList = &v
	return s
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfStream) Validate() error {
	return dara.Validate(s)
}

type UpdateTimingSyntheticTaskRequestMonitorConfWebsite struct {
	// Specifies whether to automatically scroll up and down the screen to load a page. Valid values: 0: no. 1: yes. Default value: 0.
	//
	// example:
	//
	// 0
	AutomaticScrolling *int32 `json:"AutomaticScrolling,omitempty" xml:"AutomaticScrolling,omitempty"`
	// Specifies whether to create a custom header. Valid values: 0: no. 1: The first packet is modified. 2: All packets are modified. Default value: 0.
	//
	// example:
	//
	// 0
	CustomHeader *int32 `json:"CustomHeader,omitempty" xml:"CustomHeader,omitempty"`
	// The custom header. Format: JSON map.
	CustomHeaderContent map[string]*string `json:"CustomHeaderContent,omitempty" xml:"CustomHeaderContent,omitempty"`
	// If the IP address or CNAME record resolved from a domain name is not included in the DNS whitelist, you cannot access the domain name, or an IP address that belongs to a different domain name is returned. If the IP address or CNAME record is included in the DNS whitelist, DNS hijacking does not occur.
	//
	// Format: \\<domain name>:\\<objects>. The objects can be IP addresses, wildcard mask, subnet mask, or CNAME records. Separate multiple objects with vertical bars (|). Example: www.aliyun.com:203.0.3.55|203.3.44.67. It indicates that all IP addresses that belong to the www.aliyun.com domain name except 203.0.3.55 and 203.3.44.67 are hijacked.
	//
	// example:
	//
	// www.aliyun.com:203.0.3.55|203.3.44.67
	DNSHijackWhitelist *string `json:"DNSHijackWhitelist,omitempty" xml:"DNSHijackWhitelist,omitempty"`
	// Specifies whether to disable the cache. Valid values: 0: no. 1: yes. Default value: 1.
	//
	// example:
	//
	// 0
	DisableCache *int32 `json:"DisableCache,omitempty" xml:"DisableCache,omitempty"`
	// Specifies whether to accept compressed files based on the HTTP Accept-Encoding request header. Valid values: 0: no. 1: yes. Default value: 0.
	//
	// example:
	//
	// 0
	DisableCompression *int32 `json:"DisableCompression,omitempty" xml:"DisableCompression,omitempty"`
	// The elements not to be loaded in the page loading process.
	//
	// example:
	//
	// a.jpg
	ElementBlacklist *string `json:"ElementBlacklist,omitempty" xml:"ElementBlacklist,omitempty"`
	// Specifies whether to exclude invalid IP addresses. Valid values:
	//
	// 	- 1: no
	//
	// 	- 0: yes
	//
	// example:
	//
	// 0
	FilterInvalidIP *int32 `json:"FilterInvalidIP,omitempty" xml:"FilterInvalidIP,omitempty"`
	// The total number of elements on the page.
	//
	// example:
	//
	// 1
	FlowHijackJumpTimes *int32 `json:"FlowHijackJumpTimes,omitempty" xml:"FlowHijackJumpTimes,omitempty"`
	// The keyword that is used to identify hijacking. Asterisks (\\*) are allowed.
	//
	// example:
	//
	// aliyun
	FlowHijackLogo *string `json:"FlowHijackLogo,omitempty" xml:"FlowHijackLogo,omitempty"`
	// Specifies whether to ignore SSL certificate errors during browsing. Valid values: 0: no. 1: yes. Default value: 1.
	//
	// example:
	//
	// 0
	IgnoreCertificateError *int32 `json:"IgnoreCertificateError,omitempty" xml:"IgnoreCertificateError,omitempty"`
	// The monitoring timeout period. Unit: milliseconds. Minimum value: 5000. Maximum value: 300000. Default value: 40000.
	//
	// example:
	//
	// 20000
	MonitorTimeout *int64 `json:"MonitorTimeout,omitempty" xml:"MonitorTimeout,omitempty"`
	// Elements that are not included in the whitelist and appear on the page are manipulated. These elements can be pop-up ads, floating ads, and page redirection.
	//
	// example:
	//
	// www.aliyun.com:|/cc/bb/a.gif|/vv/bb/cc.jpg
	PageTamper *string `json:"PageTamper,omitempty" xml:"PageTamper,omitempty"`
	// Specifies whether to continue browsing after redirection. Valid values: 0: no. 1: yes. Default value: 1.
	//
	// example:
	//
	// 0
	Redirection *int32 `json:"Redirection,omitempty" xml:"Redirection,omitempty"`
	// The time threshold that is used to define a slow element. Unit: milliseconds. Default value: 5000. Minimum value: 1. Maximum value: 300000.
	//
	// example:
	//
	// 5000
	SlowElementThreshold *int64 `json:"SlowElementThreshold,omitempty" xml:"SlowElementThreshold,omitempty"`
	// The URL of the website.
	//
	// example:
	//
	// https://********
	TargetUrl *string `json:"TargetUrl,omitempty" xml:"TargetUrl,omitempty"`
	// An arbitrary string in the source code of the page for verification. If the source code returned by the client contains a string that is in the blacklist, the 650 error code is reported, which indicates that the string fails to be verified. Separate multiple strings with vertical bars (|).
	//
	// example:
	//
	// error
	VerifyStringBlacklist *string `json:"VerifyStringBlacklist,omitempty" xml:"VerifyStringBlacklist,omitempty"`
	// An arbitrary string in the source code of the page for verification. If the source code returned by the client contains a string that is not in the whitelist, the 650 error code is reported, which indicates that the string fails to be verified. Separate multiple strings with vertical bars (|).
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

func (s UpdateTimingSyntheticTaskRequestMonitorConfWebsite) String() string {
	return dara.Prettify(s)
}

func (s UpdateTimingSyntheticTaskRequestMonitorConfWebsite) GoString() string {
	return s.String()
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfWebsite) GetAutomaticScrolling() *int32 {
	return s.AutomaticScrolling
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfWebsite) GetCustomHeader() *int32 {
	return s.CustomHeader
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfWebsite) GetCustomHeaderContent() map[string]*string {
	return s.CustomHeaderContent
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfWebsite) GetDNSHijackWhitelist() *string {
	return s.DNSHijackWhitelist
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfWebsite) GetDisableCache() *int32 {
	return s.DisableCache
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfWebsite) GetDisableCompression() *int32 {
	return s.DisableCompression
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfWebsite) GetElementBlacklist() *string {
	return s.ElementBlacklist
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfWebsite) GetFilterInvalidIP() *int32 {
	return s.FilterInvalidIP
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfWebsite) GetFlowHijackJumpTimes() *int32 {
	return s.FlowHijackJumpTimes
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfWebsite) GetFlowHijackLogo() *string {
	return s.FlowHijackLogo
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfWebsite) GetIgnoreCertificateError() *int32 {
	return s.IgnoreCertificateError
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfWebsite) GetMonitorTimeout() *int64 {
	return s.MonitorTimeout
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfWebsite) GetPageTamper() *string {
	return s.PageTamper
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfWebsite) GetRedirection() *int32 {
	return s.Redirection
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfWebsite) GetSlowElementThreshold() *int64 {
	return s.SlowElementThreshold
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfWebsite) GetTargetUrl() *string {
	return s.TargetUrl
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfWebsite) GetVerifyStringBlacklist() *string {
	return s.VerifyStringBlacklist
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfWebsite) GetVerifyStringWhitelist() *string {
	return s.VerifyStringWhitelist
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfWebsite) GetWaitCompletionTime() *int64 {
	return s.WaitCompletionTime
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfWebsite) SetAutomaticScrolling(v int32) *UpdateTimingSyntheticTaskRequestMonitorConfWebsite {
	s.AutomaticScrolling = &v
	return s
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfWebsite) SetCustomHeader(v int32) *UpdateTimingSyntheticTaskRequestMonitorConfWebsite {
	s.CustomHeader = &v
	return s
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfWebsite) SetCustomHeaderContent(v map[string]*string) *UpdateTimingSyntheticTaskRequestMonitorConfWebsite {
	s.CustomHeaderContent = v
	return s
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfWebsite) SetDNSHijackWhitelist(v string) *UpdateTimingSyntheticTaskRequestMonitorConfWebsite {
	s.DNSHijackWhitelist = &v
	return s
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfWebsite) SetDisableCache(v int32) *UpdateTimingSyntheticTaskRequestMonitorConfWebsite {
	s.DisableCache = &v
	return s
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfWebsite) SetDisableCompression(v int32) *UpdateTimingSyntheticTaskRequestMonitorConfWebsite {
	s.DisableCompression = &v
	return s
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfWebsite) SetElementBlacklist(v string) *UpdateTimingSyntheticTaskRequestMonitorConfWebsite {
	s.ElementBlacklist = &v
	return s
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfWebsite) SetFilterInvalidIP(v int32) *UpdateTimingSyntheticTaskRequestMonitorConfWebsite {
	s.FilterInvalidIP = &v
	return s
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfWebsite) SetFlowHijackJumpTimes(v int32) *UpdateTimingSyntheticTaskRequestMonitorConfWebsite {
	s.FlowHijackJumpTimes = &v
	return s
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfWebsite) SetFlowHijackLogo(v string) *UpdateTimingSyntheticTaskRequestMonitorConfWebsite {
	s.FlowHijackLogo = &v
	return s
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfWebsite) SetIgnoreCertificateError(v int32) *UpdateTimingSyntheticTaskRequestMonitorConfWebsite {
	s.IgnoreCertificateError = &v
	return s
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfWebsite) SetMonitorTimeout(v int64) *UpdateTimingSyntheticTaskRequestMonitorConfWebsite {
	s.MonitorTimeout = &v
	return s
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfWebsite) SetPageTamper(v string) *UpdateTimingSyntheticTaskRequestMonitorConfWebsite {
	s.PageTamper = &v
	return s
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfWebsite) SetRedirection(v int32) *UpdateTimingSyntheticTaskRequestMonitorConfWebsite {
	s.Redirection = &v
	return s
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfWebsite) SetSlowElementThreshold(v int64) *UpdateTimingSyntheticTaskRequestMonitorConfWebsite {
	s.SlowElementThreshold = &v
	return s
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfWebsite) SetTargetUrl(v string) *UpdateTimingSyntheticTaskRequestMonitorConfWebsite {
	s.TargetUrl = &v
	return s
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfWebsite) SetVerifyStringBlacklist(v string) *UpdateTimingSyntheticTaskRequestMonitorConfWebsite {
	s.VerifyStringBlacklist = &v
	return s
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfWebsite) SetVerifyStringWhitelist(v string) *UpdateTimingSyntheticTaskRequestMonitorConfWebsite {
	s.VerifyStringWhitelist = &v
	return s
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfWebsite) SetWaitCompletionTime(v int64) *UpdateTimingSyntheticTaskRequestMonitorConfWebsite {
	s.WaitCompletionTime = &v
	return s
}

func (s *UpdateTimingSyntheticTaskRequestMonitorConfWebsite) Validate() error {
	return dara.Validate(s)
}

type UpdateTimingSyntheticTaskRequestMonitors struct {
	// The city code.
	//
	// example:
	//
	// 100001
	CityCode *string `json:"CityCode,omitempty" xml:"CityCode,omitempty"`
	// The client type of the monitoring point. Valid values: 1: data center. 2: Internet. 3: mobile device. 4: ECS instance.
	//
	// example:
	//
	// 4
	ClientType *int32 `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	// The carrier code.
	//
	// example:
	//
	// 1
	OperatorCode *string `json:"OperatorCode,omitempty" xml:"OperatorCode,omitempty"`
}

func (s UpdateTimingSyntheticTaskRequestMonitors) String() string {
	return dara.Prettify(s)
}

func (s UpdateTimingSyntheticTaskRequestMonitors) GoString() string {
	return s.String()
}

func (s *UpdateTimingSyntheticTaskRequestMonitors) GetCityCode() *string {
	return s.CityCode
}

func (s *UpdateTimingSyntheticTaskRequestMonitors) GetClientType() *int32 {
	return s.ClientType
}

func (s *UpdateTimingSyntheticTaskRequestMonitors) GetOperatorCode() *string {
	return s.OperatorCode
}

func (s *UpdateTimingSyntheticTaskRequestMonitors) SetCityCode(v string) *UpdateTimingSyntheticTaskRequestMonitors {
	s.CityCode = &v
	return s
}

func (s *UpdateTimingSyntheticTaskRequestMonitors) SetClientType(v int32) *UpdateTimingSyntheticTaskRequestMonitors {
	s.ClientType = &v
	return s
}

func (s *UpdateTimingSyntheticTaskRequestMonitors) SetOperatorCode(v string) *UpdateTimingSyntheticTaskRequestMonitors {
	s.OperatorCode = &v
	return s
}

func (s *UpdateTimingSyntheticTaskRequestMonitors) Validate() error {
	return dara.Validate(s)
}

type UpdateTimingSyntheticTaskRequestTags struct {
	// The key of the tag.
	//
	// example:
	//
	// key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateTimingSyntheticTaskRequestTags) String() string {
	return dara.Prettify(s)
}

func (s UpdateTimingSyntheticTaskRequestTags) GoString() string {
	return s.String()
}

func (s *UpdateTimingSyntheticTaskRequestTags) GetKey() *string {
	return s.Key
}

func (s *UpdateTimingSyntheticTaskRequestTags) GetValue() *string {
	return s.Value
}

func (s *UpdateTimingSyntheticTaskRequestTags) SetKey(v string) *UpdateTimingSyntheticTaskRequestTags {
	s.Key = &v
	return s
}

func (s *UpdateTimingSyntheticTaskRequestTags) SetValue(v string) *UpdateTimingSyntheticTaskRequestTags {
	s.Value = &v
	return s
}

func (s *UpdateTimingSyntheticTaskRequestTags) Validate() error {
	return dara.Validate(s)
}
