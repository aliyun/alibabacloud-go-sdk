// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTimingSyntheticTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAvailableAssertions(v []*CreateTimingSyntheticTaskRequestAvailableAssertions) *CreateTimingSyntheticTaskRequest
	GetAvailableAssertions() []*CreateTimingSyntheticTaskRequestAvailableAssertions
	SetCommonSetting(v *CreateTimingSyntheticTaskRequestCommonSetting) *CreateTimingSyntheticTaskRequest
	GetCommonSetting() *CreateTimingSyntheticTaskRequestCommonSetting
	SetCustomPeriod(v *CreateTimingSyntheticTaskRequestCustomPeriod) *CreateTimingSyntheticTaskRequest
	GetCustomPeriod() *CreateTimingSyntheticTaskRequestCustomPeriod
	SetFrequency(v string) *CreateTimingSyntheticTaskRequest
	GetFrequency() *string
	SetMonitorCategory(v int32) *CreateTimingSyntheticTaskRequest
	GetMonitorCategory() *int32
	SetMonitorConf(v *CreateTimingSyntheticTaskRequestMonitorConf) *CreateTimingSyntheticTaskRequest
	GetMonitorConf() *CreateTimingSyntheticTaskRequestMonitorConf
	SetMonitors(v []*CreateTimingSyntheticTaskRequestMonitors) *CreateTimingSyntheticTaskRequest
	GetMonitors() []*CreateTimingSyntheticTaskRequestMonitors
	SetName(v string) *CreateTimingSyntheticTaskRequest
	GetName() *string
	SetRegionId(v string) *CreateTimingSyntheticTaskRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateTimingSyntheticTaskRequest
	GetResourceGroupId() *string
	SetTags(v []*CreateTimingSyntheticTaskRequestTags) *CreateTimingSyntheticTaskRequest
	GetTags() []*CreateTimingSyntheticTaskRequestTags
	SetTaskType(v int32) *CreateTimingSyntheticTaskRequest
	GetTaskType() *int32
}

type CreateTimingSyntheticTaskRequest struct {
	// The list of assertions.
	AvailableAssertions []*CreateTimingSyntheticTaskRequestAvailableAssertions `json:"AvailableAssertions,omitempty" xml:"AvailableAssertions,omitempty" type:"Repeated"`
	// The general settings.
	CommonSetting *CreateTimingSyntheticTaskRequestCommonSetting `json:"CommonSetting,omitempty" xml:"CommonSetting,omitempty" type:"Struct"`
	// The general settings.
	CustomPeriod *CreateTimingSyntheticTaskRequestCustomPeriod `json:"CustomPeriod,omitempty" xml:"CustomPeriod,omitempty" type:"Struct"`
	// The detection frequency. Valid values: 1m, 5m, 10m, 15m, 20m, 30m, 1h, 2h, 3h, 4h, 6h, 8h, 12h, and 24h.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5m
	Frequency *string `json:"Frequency,omitempty" xml:"Frequency,omitempty"`
	// The detection point type. Valid values:
	//
	// - 1: PC
	//
	// - 2: mobile device
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	MonitorCategory *int32 `json:"MonitorCategory,omitempty" xml:"MonitorCategory,omitempty"`
	// The monitoring configurations.
	//
	// This parameter is required.
	MonitorConf *CreateTimingSyntheticTaskRequestMonitorConf `json:"MonitorConf,omitempty" xml:"MonitorConf,omitempty" type:"Struct"`
	// The list of detection points.
	//
	// This parameter is required.
	Monitors []*CreateTimingSyntheticTaskRequestMonitors `json:"Monitors,omitempty" xml:"Monitors,omitempty" type:"Repeated"`
	// The name of the task.
	//
	// This parameter is required.
	//
	// example:
	//
	// demo-test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The parameter is optional.
	//
	// example:
	//
	// xxxx
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tag list.
	Tags []*CreateTimingSyntheticTaskRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The type of the task. Valid values:
	//
	// 1: ICMP. 2: TCP. 3: DNS. 4: HTTP. 5: website speed measurement. 6: file download.
	//
	// This parameter is required.
	//
	// example:
	//
	// 4
	TaskType *int32 `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s CreateTimingSyntheticTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTimingSyntheticTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateTimingSyntheticTaskRequest) GetAvailableAssertions() []*CreateTimingSyntheticTaskRequestAvailableAssertions {
	return s.AvailableAssertions
}

func (s *CreateTimingSyntheticTaskRequest) GetCommonSetting() *CreateTimingSyntheticTaskRequestCommonSetting {
	return s.CommonSetting
}

func (s *CreateTimingSyntheticTaskRequest) GetCustomPeriod() *CreateTimingSyntheticTaskRequestCustomPeriod {
	return s.CustomPeriod
}

func (s *CreateTimingSyntheticTaskRequest) GetFrequency() *string {
	return s.Frequency
}

func (s *CreateTimingSyntheticTaskRequest) GetMonitorCategory() *int32 {
	return s.MonitorCategory
}

func (s *CreateTimingSyntheticTaskRequest) GetMonitorConf() *CreateTimingSyntheticTaskRequestMonitorConf {
	return s.MonitorConf
}

func (s *CreateTimingSyntheticTaskRequest) GetMonitors() []*CreateTimingSyntheticTaskRequestMonitors {
	return s.Monitors
}

func (s *CreateTimingSyntheticTaskRequest) GetName() *string {
	return s.Name
}

func (s *CreateTimingSyntheticTaskRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateTimingSyntheticTaskRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateTimingSyntheticTaskRequest) GetTags() []*CreateTimingSyntheticTaskRequestTags {
	return s.Tags
}

func (s *CreateTimingSyntheticTaskRequest) GetTaskType() *int32 {
	return s.TaskType
}

func (s *CreateTimingSyntheticTaskRequest) SetAvailableAssertions(v []*CreateTimingSyntheticTaskRequestAvailableAssertions) *CreateTimingSyntheticTaskRequest {
	s.AvailableAssertions = v
	return s
}

func (s *CreateTimingSyntheticTaskRequest) SetCommonSetting(v *CreateTimingSyntheticTaskRequestCommonSetting) *CreateTimingSyntheticTaskRequest {
	s.CommonSetting = v
	return s
}

func (s *CreateTimingSyntheticTaskRequest) SetCustomPeriod(v *CreateTimingSyntheticTaskRequestCustomPeriod) *CreateTimingSyntheticTaskRequest {
	s.CustomPeriod = v
	return s
}

func (s *CreateTimingSyntheticTaskRequest) SetFrequency(v string) *CreateTimingSyntheticTaskRequest {
	s.Frequency = &v
	return s
}

func (s *CreateTimingSyntheticTaskRequest) SetMonitorCategory(v int32) *CreateTimingSyntheticTaskRequest {
	s.MonitorCategory = &v
	return s
}

func (s *CreateTimingSyntheticTaskRequest) SetMonitorConf(v *CreateTimingSyntheticTaskRequestMonitorConf) *CreateTimingSyntheticTaskRequest {
	s.MonitorConf = v
	return s
}

func (s *CreateTimingSyntheticTaskRequest) SetMonitors(v []*CreateTimingSyntheticTaskRequestMonitors) *CreateTimingSyntheticTaskRequest {
	s.Monitors = v
	return s
}

func (s *CreateTimingSyntheticTaskRequest) SetName(v string) *CreateTimingSyntheticTaskRequest {
	s.Name = &v
	return s
}

func (s *CreateTimingSyntheticTaskRequest) SetRegionId(v string) *CreateTimingSyntheticTaskRequest {
	s.RegionId = &v
	return s
}

func (s *CreateTimingSyntheticTaskRequest) SetResourceGroupId(v string) *CreateTimingSyntheticTaskRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateTimingSyntheticTaskRequest) SetTags(v []*CreateTimingSyntheticTaskRequestTags) *CreateTimingSyntheticTaskRequest {
	s.Tags = v
	return s
}

func (s *CreateTimingSyntheticTaskRequest) SetTaskType(v int32) *CreateTimingSyntheticTaskRequest {
	s.TaskType = &v
	return s
}

func (s *CreateTimingSyntheticTaskRequest) Validate() error {
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

type CreateTimingSyntheticTaskRequestAvailableAssertions struct {
	// The expected value.
	//
	// This parameter is required.
	//
	// example:
	//
	// 200
	Expect *string `json:"Expect,omitempty" xml:"Expect,omitempty"`
	// The condition. gt: greater than. gte: greater than or equal to. lt: less than. lte: less than or equal to. eq: equal to. neq: not equal to. ctn: contain. nctn: does not contain. exist: exist. n_exist: does not exist. belong: belong to. n_belong: does not belong to. reg_match: regular expression.
	//
	// This parameter is required.
	//
	// example:
	//
	// eq
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// The check target. If you set the type parameter to HttpResCode, HttpResBody, or HttpResponseTime, you do not need to set the target parameter. If you set the type parameter to HttpResHead, you must specify the key in the header. If you set the type parameter to HttpResBodyJson, use jsonPath.
	//
	// example:
	//
	// key
	Target *string `json:"Target,omitempty" xml:"Target,omitempty"`
	// The assertion type. Valid values: HttpResCode, HttpResHead, HttpResBody, HttpResBodyJson, HttpResponseTime, IcmpPackLoss (packet loss rate), IcmpPackMaxLatency (maximum packet latency), IcmpPackAvgLatency (average packet latency), TraceRouteHops (number of hops), DnsARecord (A record), DnsCName (CNAME), websiteTTFB (time to first packet), websiteTTLB (time to last packet), websiteFST (first paint time), websiteFFST (first meaningful paint), websiteOnload (full loaded time). For more information, see the following description.
	//
	// This parameter is required.
	//
	// example:
	//
	// DnsARecord
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateTimingSyntheticTaskRequestAvailableAssertions) String() string {
	return dara.Prettify(s)
}

func (s CreateTimingSyntheticTaskRequestAvailableAssertions) GoString() string {
	return s.String()
}

func (s *CreateTimingSyntheticTaskRequestAvailableAssertions) GetExpect() *string {
	return s.Expect
}

func (s *CreateTimingSyntheticTaskRequestAvailableAssertions) GetOperator() *string {
	return s.Operator
}

func (s *CreateTimingSyntheticTaskRequestAvailableAssertions) GetTarget() *string {
	return s.Target
}

func (s *CreateTimingSyntheticTaskRequestAvailableAssertions) GetType() *string {
	return s.Type
}

func (s *CreateTimingSyntheticTaskRequestAvailableAssertions) SetExpect(v string) *CreateTimingSyntheticTaskRequestAvailableAssertions {
	s.Expect = &v
	return s
}

func (s *CreateTimingSyntheticTaskRequestAvailableAssertions) SetOperator(v string) *CreateTimingSyntheticTaskRequestAvailableAssertions {
	s.Operator = &v
	return s
}

func (s *CreateTimingSyntheticTaskRequestAvailableAssertions) SetTarget(v string) *CreateTimingSyntheticTaskRequestAvailableAssertions {
	s.Target = &v
	return s
}

func (s *CreateTimingSyntheticTaskRequestAvailableAssertions) SetType(v string) *CreateTimingSyntheticTaskRequestAvailableAssertions {
	s.Type = &v
	return s
}

func (s *CreateTimingSyntheticTaskRequestAvailableAssertions) Validate() error {
	return dara.Validate(s)
}

type CreateTimingSyntheticTaskRequestCommonSetting struct {
	// The custom host settings.
	CustomHost *CreateTimingSyntheticTaskRequestCommonSettingCustomHost `json:"CustomHost,omitempty" xml:"CustomHost,omitempty" type:"Struct"`
	// The reserved parameters.
	CustomPrometheusSetting *CreateTimingSyntheticTaskRequestCommonSettingCustomPrometheusSetting `json:"CustomPrometheusSetting,omitempty" xml:"CustomPrometheusSetting,omitempty" type:"Struct"`
	// The information about the virtual private cloud (VPC). If the destination URL is an Alibaba Cloud internal endpoint, you need to configure a VPC.
	CustomVPCSetting *CreateTimingSyntheticTaskRequestCommonSettingCustomVPCSetting `json:"CustomVPCSetting,omitempty" xml:"CustomVPCSetting,omitempty" type:"Struct"`
	// The IP version. Valid values:
	//
	// 	- 0: A version is automatically selected.
	//
	// 	- 1: IPv4.
	//
	// 	- 2: IPv6.
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

func (s CreateTimingSyntheticTaskRequestCommonSetting) String() string {
	return dara.Prettify(s)
}

func (s CreateTimingSyntheticTaskRequestCommonSetting) GoString() string {
	return s.String()
}

func (s *CreateTimingSyntheticTaskRequestCommonSetting) GetCustomHost() *CreateTimingSyntheticTaskRequestCommonSettingCustomHost {
	return s.CustomHost
}

func (s *CreateTimingSyntheticTaskRequestCommonSetting) GetCustomPrometheusSetting() *CreateTimingSyntheticTaskRequestCommonSettingCustomPrometheusSetting {
	return s.CustomPrometheusSetting
}

func (s *CreateTimingSyntheticTaskRequestCommonSetting) GetCustomVPCSetting() *CreateTimingSyntheticTaskRequestCommonSettingCustomVPCSetting {
	return s.CustomVPCSetting
}

func (s *CreateTimingSyntheticTaskRequestCommonSetting) GetIpType() *int32 {
	return s.IpType
}

func (s *CreateTimingSyntheticTaskRequestCommonSetting) GetIsOpenTrace() *bool {
	return s.IsOpenTrace
}

func (s *CreateTimingSyntheticTaskRequestCommonSetting) GetMonitorSamples() *int32 {
	return s.MonitorSamples
}

func (s *CreateTimingSyntheticTaskRequestCommonSetting) GetTraceClientType() *int32 {
	return s.TraceClientType
}

func (s *CreateTimingSyntheticTaskRequestCommonSetting) GetXtraceRegion() *string {
	return s.XtraceRegion
}

func (s *CreateTimingSyntheticTaskRequestCommonSetting) SetCustomHost(v *CreateTimingSyntheticTaskRequestCommonSettingCustomHost) *CreateTimingSyntheticTaskRequestCommonSetting {
	s.CustomHost = v
	return s
}

func (s *CreateTimingSyntheticTaskRequestCommonSetting) SetCustomPrometheusSetting(v *CreateTimingSyntheticTaskRequestCommonSettingCustomPrometheusSetting) *CreateTimingSyntheticTaskRequestCommonSetting {
	s.CustomPrometheusSetting = v
	return s
}

func (s *CreateTimingSyntheticTaskRequestCommonSetting) SetCustomVPCSetting(v *CreateTimingSyntheticTaskRequestCommonSettingCustomVPCSetting) *CreateTimingSyntheticTaskRequestCommonSetting {
	s.CustomVPCSetting = v
	return s
}

func (s *CreateTimingSyntheticTaskRequestCommonSetting) SetIpType(v int32) *CreateTimingSyntheticTaskRequestCommonSetting {
	s.IpType = &v
	return s
}

func (s *CreateTimingSyntheticTaskRequestCommonSetting) SetIsOpenTrace(v bool) *CreateTimingSyntheticTaskRequestCommonSetting {
	s.IsOpenTrace = &v
	return s
}

func (s *CreateTimingSyntheticTaskRequestCommonSetting) SetMonitorSamples(v int32) *CreateTimingSyntheticTaskRequestCommonSetting {
	s.MonitorSamples = &v
	return s
}

func (s *CreateTimingSyntheticTaskRequestCommonSetting) SetTraceClientType(v int32) *CreateTimingSyntheticTaskRequestCommonSetting {
	s.TraceClientType = &v
	return s
}

func (s *CreateTimingSyntheticTaskRequestCommonSetting) SetXtraceRegion(v string) *CreateTimingSyntheticTaskRequestCommonSetting {
	s.XtraceRegion = &v
	return s
}

func (s *CreateTimingSyntheticTaskRequestCommonSetting) Validate() error {
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

type CreateTimingSyntheticTaskRequestCommonSettingCustomHost struct {
	// The list of hosts.
	//
	// This parameter is required.
	Hosts []*CreateTimingSyntheticTaskRequestCommonSettingCustomHostHosts `json:"Hosts,omitempty" xml:"Hosts,omitempty" type:"Repeated"`
	// The selection mode. Valid values:
	//
	// 	- 0: random
	//
	// 	- 1: polling
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	SelectType *int32 `json:"SelectType,omitempty" xml:"SelectType,omitempty"`
}

func (s CreateTimingSyntheticTaskRequestCommonSettingCustomHost) String() string {
	return dara.Prettify(s)
}

func (s CreateTimingSyntheticTaskRequestCommonSettingCustomHost) GoString() string {
	return s.String()
}

func (s *CreateTimingSyntheticTaskRequestCommonSettingCustomHost) GetHosts() []*CreateTimingSyntheticTaskRequestCommonSettingCustomHostHosts {
	return s.Hosts
}

func (s *CreateTimingSyntheticTaskRequestCommonSettingCustomHost) GetSelectType() *int32 {
	return s.SelectType
}

func (s *CreateTimingSyntheticTaskRequestCommonSettingCustomHost) SetHosts(v []*CreateTimingSyntheticTaskRequestCommonSettingCustomHostHosts) *CreateTimingSyntheticTaskRequestCommonSettingCustomHost {
	s.Hosts = v
	return s
}

func (s *CreateTimingSyntheticTaskRequestCommonSettingCustomHost) SetSelectType(v int32) *CreateTimingSyntheticTaskRequestCommonSettingCustomHost {
	s.SelectType = &v
	return s
}

func (s *CreateTimingSyntheticTaskRequestCommonSettingCustomHost) Validate() error {
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

type CreateTimingSyntheticTaskRequestCommonSettingCustomHostHosts struct {
	// The domain name.
	//
	// This parameter is required.
	//
	// example:
	//
	// www.aliyun.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The IP version. Valid values:
	//
	// 	- 0: A version is automatically selected.
	//
	// 	- 1: IPv4.
	//
	// 	- 2: IPv6.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	IpType *int32 `json:"IpType,omitempty" xml:"IpType,omitempty"`
	// The list of IP addresses.
	//
	// This parameter is required.
	Ips []*string `json:"Ips,omitempty" xml:"Ips,omitempty" type:"Repeated"`
}

func (s CreateTimingSyntheticTaskRequestCommonSettingCustomHostHosts) String() string {
	return dara.Prettify(s)
}

func (s CreateTimingSyntheticTaskRequestCommonSettingCustomHostHosts) GoString() string {
	return s.String()
}

func (s *CreateTimingSyntheticTaskRequestCommonSettingCustomHostHosts) GetDomain() *string {
	return s.Domain
}

func (s *CreateTimingSyntheticTaskRequestCommonSettingCustomHostHosts) GetIpType() *int32 {
	return s.IpType
}

func (s *CreateTimingSyntheticTaskRequestCommonSettingCustomHostHosts) GetIps() []*string {
	return s.Ips
}

func (s *CreateTimingSyntheticTaskRequestCommonSettingCustomHostHosts) SetDomain(v string) *CreateTimingSyntheticTaskRequestCommonSettingCustomHostHosts {
	s.Domain = &v
	return s
}

func (s *CreateTimingSyntheticTaskRequestCommonSettingCustomHostHosts) SetIpType(v int32) *CreateTimingSyntheticTaskRequestCommonSettingCustomHostHosts {
	s.IpType = &v
	return s
}

func (s *CreateTimingSyntheticTaskRequestCommonSettingCustomHostHosts) SetIps(v []*string) *CreateTimingSyntheticTaskRequestCommonSettingCustomHostHosts {
	s.Ips = v
	return s
}

func (s *CreateTimingSyntheticTaskRequestCommonSettingCustomHostHosts) Validate() error {
	return dara.Validate(s)
}

type CreateTimingSyntheticTaskRequestCommonSettingCustomPrometheusSetting struct {
	// A reserved parameter.
	//
	// example:
	//
	// null
	PrometheusClusterId *string `json:"PrometheusClusterId,omitempty" xml:"PrometheusClusterId,omitempty"`
	// A reserved parameter.
	//
	// example:
	//
	// null
	PrometheusClusterRegion *string `json:"PrometheusClusterRegion,omitempty" xml:"PrometheusClusterRegion,omitempty"`
	// A reserved parameter.
	PrometheusLabels map[string]*string `json:"PrometheusLabels,omitempty" xml:"PrometheusLabels,omitempty"`
}

func (s CreateTimingSyntheticTaskRequestCommonSettingCustomPrometheusSetting) String() string {
	return dara.Prettify(s)
}

func (s CreateTimingSyntheticTaskRequestCommonSettingCustomPrometheusSetting) GoString() string {
	return s.String()
}

func (s *CreateTimingSyntheticTaskRequestCommonSettingCustomPrometheusSetting) GetPrometheusClusterId() *string {
	return s.PrometheusClusterId
}

func (s *CreateTimingSyntheticTaskRequestCommonSettingCustomPrometheusSetting) GetPrometheusClusterRegion() *string {
	return s.PrometheusClusterRegion
}

func (s *CreateTimingSyntheticTaskRequestCommonSettingCustomPrometheusSetting) GetPrometheusLabels() map[string]*string {
	return s.PrometheusLabels
}

func (s *CreateTimingSyntheticTaskRequestCommonSettingCustomPrometheusSetting) SetPrometheusClusterId(v string) *CreateTimingSyntheticTaskRequestCommonSettingCustomPrometheusSetting {
	s.PrometheusClusterId = &v
	return s
}

func (s *CreateTimingSyntheticTaskRequestCommonSettingCustomPrometheusSetting) SetPrometheusClusterRegion(v string) *CreateTimingSyntheticTaskRequestCommonSettingCustomPrometheusSetting {
	s.PrometheusClusterRegion = &v
	return s
}

func (s *CreateTimingSyntheticTaskRequestCommonSettingCustomPrometheusSetting) SetPrometheusLabels(v map[string]*string) *CreateTimingSyntheticTaskRequestCommonSettingCustomPrometheusSetting {
	s.PrometheusLabels = v
	return s
}

func (s *CreateTimingSyntheticTaskRequestCommonSettingCustomPrometheusSetting) Validate() error {
	return dara.Validate(s)
}

type CreateTimingSyntheticTaskRequestCommonSettingCustomVPCSetting struct {
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
	// sg-bp13wzf9vuwegmpxxxxx
	SecureGroupId *string `json:"SecureGroupId,omitempty" xml:"SecureGroupId,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-bp14crq29vpycxp8xxxxx
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// VPC ID.
	//
	// example:
	//
	// vpc-bp1muectbr8f90vjxxxxx
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s CreateTimingSyntheticTaskRequestCommonSettingCustomVPCSetting) String() string {
	return dara.Prettify(s)
}

func (s CreateTimingSyntheticTaskRequestCommonSettingCustomVPCSetting) GoString() string {
	return s.String()
}

func (s *CreateTimingSyntheticTaskRequestCommonSettingCustomVPCSetting) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateTimingSyntheticTaskRequestCommonSettingCustomVPCSetting) GetSecureGroupId() *string {
	return s.SecureGroupId
}

func (s *CreateTimingSyntheticTaskRequestCommonSettingCustomVPCSetting) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateTimingSyntheticTaskRequestCommonSettingCustomVPCSetting) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateTimingSyntheticTaskRequestCommonSettingCustomVPCSetting) SetRegionId(v string) *CreateTimingSyntheticTaskRequestCommonSettingCustomVPCSetting {
	s.RegionId = &v
	return s
}

func (s *CreateTimingSyntheticTaskRequestCommonSettingCustomVPCSetting) SetSecureGroupId(v string) *CreateTimingSyntheticTaskRequestCommonSettingCustomVPCSetting {
	s.SecureGroupId = &v
	return s
}

func (s *CreateTimingSyntheticTaskRequestCommonSettingCustomVPCSetting) SetVSwitchId(v string) *CreateTimingSyntheticTaskRequestCommonSettingCustomVPCSetting {
	s.VSwitchId = &v
	return s
}

func (s *CreateTimingSyntheticTaskRequestCommonSettingCustomVPCSetting) SetVpcId(v string) *CreateTimingSyntheticTaskRequestCommonSettingCustomVPCSetting {
	s.VpcId = &v
	return s
}

func (s *CreateTimingSyntheticTaskRequestCommonSettingCustomVPCSetting) Validate() error {
	return dara.Validate(s)
}

type CreateTimingSyntheticTaskRequestCustomPeriod struct {
	// The custom host settings.
	//
	// This parameter is required.
	//
	// example:
	//
	// 22
	EndHour *int32 `json:"EndHour,omitempty" xml:"EndHour,omitempty"`
	// The list of hosts.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8
	StartHour *int32 `json:"StartHour,omitempty" xml:"StartHour,omitempty"`
}

func (s CreateTimingSyntheticTaskRequestCustomPeriod) String() string {
	return dara.Prettify(s)
}

func (s CreateTimingSyntheticTaskRequestCustomPeriod) GoString() string {
	return s.String()
}

func (s *CreateTimingSyntheticTaskRequestCustomPeriod) GetEndHour() *int32 {
	return s.EndHour
}

func (s *CreateTimingSyntheticTaskRequestCustomPeriod) GetStartHour() *int32 {
	return s.StartHour
}

func (s *CreateTimingSyntheticTaskRequestCustomPeriod) SetEndHour(v int32) *CreateTimingSyntheticTaskRequestCustomPeriod {
	s.EndHour = &v
	return s
}

func (s *CreateTimingSyntheticTaskRequestCustomPeriod) SetStartHour(v int32) *CreateTimingSyntheticTaskRequestCustomPeriod {
	s.StartHour = &v
	return s
}

func (s *CreateTimingSyntheticTaskRequestCustomPeriod) Validate() error {
	return dara.Validate(s)
}

type CreateTimingSyntheticTaskRequestMonitorConf struct {
	// The parameters of the HTTP(S) synthetic test.
	ApiHTTP *CreateTimingSyntheticTaskRequestMonitorConfApiHTTP `json:"ApiHTTP,omitempty" xml:"ApiHTTP,omitempty" type:"Struct"`
	// The parameters of file downloading.
	FileDownload *CreateTimingSyntheticTaskRequestMonitorConfFileDownload `json:"FileDownload,omitempty" xml:"FileDownload,omitempty" type:"Struct"`
	// The parameters of the DNS synthetic test. This parameter is required if the TaskType parameter is set to 3.
	NetDNS *CreateTimingSyntheticTaskRequestMonitorConfNetDNS `json:"NetDNS,omitempty" xml:"NetDNS,omitempty" type:"Struct"`
	// The parameters of the ICMP synthetic test. This parameter is required if the TaskType parameter is set to 1.
	NetICMP *CreateTimingSyntheticTaskRequestMonitorConfNetICMP `json:"NetICMP,omitempty" xml:"NetICMP,omitempty" type:"Struct"`
	// The parameters of the TCP synthetic test. This parameter is required if the TaskType parameter is set to 2.
	NetTCP *CreateTimingSyntheticTaskRequestMonitorConfNetTCP `json:"NetTCP,omitempty" xml:"NetTCP,omitempty" type:"Struct"`
	// The parameters of the streaming-media synthetic test.
	Stream *CreateTimingSyntheticTaskRequestMonitorConfStream `json:"Stream,omitempty" xml:"Stream,omitempty" type:"Struct"`
	// The parameters of the website speed measurement.
	Website *CreateTimingSyntheticTaskRequestMonitorConfWebsite `json:"Website,omitempty" xml:"Website,omitempty" type:"Struct"`
}

func (s CreateTimingSyntheticTaskRequestMonitorConf) String() string {
	return dara.Prettify(s)
}

func (s CreateTimingSyntheticTaskRequestMonitorConf) GoString() string {
	return s.String()
}

func (s *CreateTimingSyntheticTaskRequestMonitorConf) GetApiHTTP() *CreateTimingSyntheticTaskRequestMonitorConfApiHTTP {
	return s.ApiHTTP
}

func (s *CreateTimingSyntheticTaskRequestMonitorConf) GetFileDownload() *CreateTimingSyntheticTaskRequestMonitorConfFileDownload {
	return s.FileDownload
}

func (s *CreateTimingSyntheticTaskRequestMonitorConf) GetNetDNS() *CreateTimingSyntheticTaskRequestMonitorConfNetDNS {
	return s.NetDNS
}

func (s *CreateTimingSyntheticTaskRequestMonitorConf) GetNetICMP() *CreateTimingSyntheticTaskRequestMonitorConfNetICMP {
	return s.NetICMP
}

func (s *CreateTimingSyntheticTaskRequestMonitorConf) GetNetTCP() *CreateTimingSyntheticTaskRequestMonitorConfNetTCP {
	return s.NetTCP
}

func (s *CreateTimingSyntheticTaskRequestMonitorConf) GetStream() *CreateTimingSyntheticTaskRequestMonitorConfStream {
	return s.Stream
}

func (s *CreateTimingSyntheticTaskRequestMonitorConf) GetWebsite() *CreateTimingSyntheticTaskRequestMonitorConfWebsite {
	return s.Website
}

func (s *CreateTimingSyntheticTaskRequestMonitorConf) SetApiHTTP(v *CreateTimingSyntheticTaskRequestMonitorConfApiHTTP) *CreateTimingSyntheticTaskRequestMonitorConf {
	s.ApiHTTP = v
	return s
}

func (s *CreateTimingSyntheticTaskRequestMonitorConf) SetFileDownload(v *CreateTimingSyntheticTaskRequestMonitorConfFileDownload) *CreateTimingSyntheticTaskRequestMonitorConf {
	s.FileDownload = v
	return s
}

func (s *CreateTimingSyntheticTaskRequestMonitorConf) SetNetDNS(v *CreateTimingSyntheticTaskRequestMonitorConfNetDNS) *CreateTimingSyntheticTaskRequestMonitorConf {
	s.NetDNS = v
	return s
}

func (s *CreateTimingSyntheticTaskRequestMonitorConf) SetNetICMP(v *CreateTimingSyntheticTaskRequestMonitorConfNetICMP) *CreateTimingSyntheticTaskRequestMonitorConf {
	s.NetICMP = v
	return s
}

func (s *CreateTimingSyntheticTaskRequestMonitorConf) SetNetTCP(v *CreateTimingSyntheticTaskRequestMonitorConfNetTCP) *CreateTimingSyntheticTaskRequestMonitorConf {
	s.NetTCP = v
	return s
}

func (s *CreateTimingSyntheticTaskRequestMonitorConf) SetStream(v *CreateTimingSyntheticTaskRequestMonitorConfStream) *CreateTimingSyntheticTaskRequestMonitorConf {
	s.Stream = v
	return s
}

func (s *CreateTimingSyntheticTaskRequestMonitorConf) SetWebsite(v *CreateTimingSyntheticTaskRequestMonitorConfWebsite) *CreateTimingSyntheticTaskRequestMonitorConf {
	s.Website = v
	return s
}

func (s *CreateTimingSyntheticTaskRequestMonitorConf) Validate() error {
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

type CreateTimingSyntheticTaskRequestMonitorConfApiHTTP struct {
	// Specifies whether to verify the certificate. Default value: no.
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
	// The request method. Valid values: GET and POST.
	//
	// example:
	//
	// GET
	Method *string `json:"Method,omitempty" xml:"Method,omitempty"`
	// The ALPN protocol version. You can configure this parameter when you perform an HTTPS synthetic test on a WAP mobile client. Valid values:
	//
	// 0: default
	//
	// 1: http/1.1
	//
	// 2: h2
	//
	// 3: disables the ALPN protocol
	//
	// example:
	//
	// 1
	ProtocolAlpnProtocol *int32 `json:"ProtocolAlpnProtocol,omitempty" xml:"ProtocolAlpnProtocol,omitempty"`
	// The HTTP request body.
	RequestBody *CreateTimingSyntheticTaskRequestMonitorConfApiHTTPRequestBody `json:"RequestBody,omitempty" xml:"RequestBody,omitempty" type:"Struct"`
	// The HTTP request header.
	RequestHeaders map[string]*string `json:"RequestHeaders,omitempty" xml:"RequestHeaders,omitempty"`
	// The URL or request path for synthetic monitoring.
	//
	// This parameter is required.
	//
	// example:
	//
	// http://www.demo.com/api/list
	TargetUrl *string `json:"TargetUrl,omitempty" xml:"TargetUrl,omitempty"`
	// The timeout period. Unit: milliseconds. Default value: 10000. Minimum value: 1000. Maximum value: 300000.
	//
	// example:
	//
	// 10000
	Timeout *int64 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s CreateTimingSyntheticTaskRequestMonitorConfApiHTTP) String() string {
	return dara.Prettify(s)
}

func (s CreateTimingSyntheticTaskRequestMonitorConfApiHTTP) GoString() string {
	return s.String()
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfApiHTTP) GetCheckCert() *bool {
	return s.CheckCert
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfApiHTTP) GetConnectTimeout() *int64 {
	return s.ConnectTimeout
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfApiHTTP) GetMethod() *string {
	return s.Method
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfApiHTTP) GetProtocolAlpnProtocol() *int32 {
	return s.ProtocolAlpnProtocol
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfApiHTTP) GetRequestBody() *CreateTimingSyntheticTaskRequestMonitorConfApiHTTPRequestBody {
	return s.RequestBody
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfApiHTTP) GetRequestHeaders() map[string]*string {
	return s.RequestHeaders
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfApiHTTP) GetTargetUrl() *string {
	return s.TargetUrl
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfApiHTTP) GetTimeout() *int64 {
	return s.Timeout
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfApiHTTP) SetCheckCert(v bool) *CreateTimingSyntheticTaskRequestMonitorConfApiHTTP {
	s.CheckCert = &v
	return s
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfApiHTTP) SetConnectTimeout(v int64) *CreateTimingSyntheticTaskRequestMonitorConfApiHTTP {
	s.ConnectTimeout = &v
	return s
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfApiHTTP) SetMethod(v string) *CreateTimingSyntheticTaskRequestMonitorConfApiHTTP {
	s.Method = &v
	return s
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfApiHTTP) SetProtocolAlpnProtocol(v int32) *CreateTimingSyntheticTaskRequestMonitorConfApiHTTP {
	s.ProtocolAlpnProtocol = &v
	return s
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfApiHTTP) SetRequestBody(v *CreateTimingSyntheticTaskRequestMonitorConfApiHTTPRequestBody) *CreateTimingSyntheticTaskRequestMonitorConfApiHTTP {
	s.RequestBody = v
	return s
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfApiHTTP) SetRequestHeaders(v map[string]*string) *CreateTimingSyntheticTaskRequestMonitorConfApiHTTP {
	s.RequestHeaders = v
	return s
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfApiHTTP) SetTargetUrl(v string) *CreateTimingSyntheticTaskRequestMonitorConfApiHTTP {
	s.TargetUrl = &v
	return s
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfApiHTTP) SetTimeout(v int64) *CreateTimingSyntheticTaskRequestMonitorConfApiHTTP {
	s.Timeout = &v
	return s
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfApiHTTP) Validate() error {
	if s.RequestBody != nil {
		if err := s.RequestBody.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateTimingSyntheticTaskRequestMonitorConfApiHTTPRequestBody struct {
	// The content of the request body. Format: JSON string. The parameter is required if the Type parameter is set to text/plain, application/json, application/xml, or text/html. Format: JSON string.
	//
	// example:
	//
	// {
	//
	//   "key1": "value1",
	//
	//   "key2": "value2"
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

func (s CreateTimingSyntheticTaskRequestMonitorConfApiHTTPRequestBody) String() string {
	return dara.Prettify(s)
}

func (s CreateTimingSyntheticTaskRequestMonitorConfApiHTTPRequestBody) GoString() string {
	return s.String()
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfApiHTTPRequestBody) GetContent() *string {
	return s.Content
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfApiHTTPRequestBody) GetType() *string {
	return s.Type
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfApiHTTPRequestBody) SetContent(v string) *CreateTimingSyntheticTaskRequestMonitorConfApiHTTPRequestBody {
	s.Content = &v
	return s
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfApiHTTPRequestBody) SetType(v string) *CreateTimingSyntheticTaskRequestMonitorConfApiHTTPRequestBody {
	s.Type = &v
	return s
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfApiHTTPRequestBody) Validate() error {
	return dara.Validate(s)
}

type CreateTimingSyntheticTaskRequestMonitorConfFileDownload struct {
	// Unit: milliseconds. Minimum value: 1000. Maximum value: 120000. Default value: 5000.
	//
	// example:
	//
	// 5000
	ConnectionTimeout *int64 `json:"ConnectionTimeout,omitempty" xml:"ConnectionTimeout,omitempty"`
	// The content of the custom request header.
	CustomHeaderContent map[string]*string `json:"CustomHeaderContent,omitempty" xml:"CustomHeaderContent,omitempty"`
	// The kernel type. Valid values:
	//
	// 	- 1: curl
	//
	// 	- 0: WinInet
	//
	// Default value: 1
	//
	// example:
	//
	// 1
	DownloadKernel *int32 `json:"DownloadKernel,omitempty" xml:"DownloadKernel,omitempty"`
	// Specifies whether to ignore CA certificate authentication errors. Valid values: 0: no. 1: yes. Default value: 1.
	//
	// example:
	//
	// 1
	IgnoreCertificateAuthError *int32 `json:"IgnoreCertificateAuthError,omitempty" xml:"IgnoreCertificateAuthError,omitempty"`
	// Specifies whether to ignore certificate revocation errors. Valid values: 0: no. 1: yes. Default value: 1.
	//
	// example:
	//
	// 1
	IgnoreCertificateCanceledError *int32 `json:"IgnoreCertificateCanceledError,omitempty" xml:"IgnoreCertificateCanceledError,omitempty"`
	// Specifies whether to ignore certificate invalidity. Valid values: 0: no. 1: yes. Default value: 1.
	//
	// example:
	//
	// 1
	IgnoreCertificateOutOfDateError *int32 `json:"IgnoreCertificateOutOfDateError,omitempty" xml:"IgnoreCertificateOutOfDateError,omitempty"`
	// Specifies whether to ignore certificate status errors. Valid values: 0: no. 1: yes. Default value: 1.
	//
	// example:
	//
	// 1
	IgnoreCertificateStatusError *int32 `json:"IgnoreCertificateStatusError,omitempty" xml:"IgnoreCertificateStatusError,omitempty"`
	// Specifies whether to ignore certificate incredibility. Valid values: 0: no. 1: yes. Default value: 1.
	//
	// example:
	//
	// 1
	IgnoreCertificateUntrustworthyError *int32 `json:"IgnoreCertificateUntrustworthyError,omitempty" xml:"IgnoreCertificateUntrustworthyError,omitempty"`
	// Specifies whether to ignore certificate usage errors. Valid values: 0: no. 1: yes. Default value: 1.
	//
	// example:
	//
	// 1
	IgnoreCertificateUsingError *int32 `json:"IgnoreCertificateUsingError,omitempty" xml:"IgnoreCertificateUsingError,omitempty"`
	// Specifies whether to ignore host invalidity. Valid values: 0: no. 1: yes. Default value: 1.
	//
	// example:
	//
	// 1
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
	// 	- 3: HTTP/3
	//
	// Default value: 1
	//
	// example:
	//
	// 1
	QuickProtocol *int32 `json:"QuickProtocol,omitempty" xml:"QuickProtocol,omitempty"`
	// Specifies whether to support redirection. Valid values: 0: no. 1: yes. Default value: 1.
	//
	// example:
	//
	// 1
	Redirection *int32 `json:"Redirection,omitempty" xml:"Redirection,omitempty"`
	// The URL that is used to download the file.
	//
	// This parameter is required.
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

func (s CreateTimingSyntheticTaskRequestMonitorConfFileDownload) String() string {
	return dara.Prettify(s)
}

func (s CreateTimingSyntheticTaskRequestMonitorConfFileDownload) GoString() string {
	return s.String()
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfFileDownload) GetConnectionTimeout() *int64 {
	return s.ConnectionTimeout
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfFileDownload) GetCustomHeaderContent() map[string]*string {
	return s.CustomHeaderContent
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfFileDownload) GetDownloadKernel() *int32 {
	return s.DownloadKernel
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfFileDownload) GetIgnoreCertificateAuthError() *int32 {
	return s.IgnoreCertificateAuthError
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfFileDownload) GetIgnoreCertificateCanceledError() *int32 {
	return s.IgnoreCertificateCanceledError
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfFileDownload) GetIgnoreCertificateOutOfDateError() *int32 {
	return s.IgnoreCertificateOutOfDateError
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfFileDownload) GetIgnoreCertificateStatusError() *int32 {
	return s.IgnoreCertificateStatusError
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfFileDownload) GetIgnoreCertificateUntrustworthyError() *int32 {
	return s.IgnoreCertificateUntrustworthyError
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfFileDownload) GetIgnoreCertificateUsingError() *int32 {
	return s.IgnoreCertificateUsingError
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfFileDownload) GetIgnoreInvalidHostError() *int32 {
	return s.IgnoreInvalidHostError
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfFileDownload) GetMonitorTimeout() *int64 {
	return s.MonitorTimeout
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfFileDownload) GetQuickProtocol() *int32 {
	return s.QuickProtocol
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfFileDownload) GetRedirection() *int32 {
	return s.Redirection
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfFileDownload) GetTargetUrl() *string {
	return s.TargetUrl
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfFileDownload) GetTransmissionSize() *int64 {
	return s.TransmissionSize
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfFileDownload) GetValidateKeywords() *string {
	return s.ValidateKeywords
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfFileDownload) GetVerifyWay() *int32 {
	return s.VerifyWay
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfFileDownload) GetWhiteList() *string {
	return s.WhiteList
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfFileDownload) SetConnectionTimeout(v int64) *CreateTimingSyntheticTaskRequestMonitorConfFileDownload {
	s.ConnectionTimeout = &v
	return s
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfFileDownload) SetCustomHeaderContent(v map[string]*string) *CreateTimingSyntheticTaskRequestMonitorConfFileDownload {
	s.CustomHeaderContent = v
	return s
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfFileDownload) SetDownloadKernel(v int32) *CreateTimingSyntheticTaskRequestMonitorConfFileDownload {
	s.DownloadKernel = &v
	return s
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfFileDownload) SetIgnoreCertificateAuthError(v int32) *CreateTimingSyntheticTaskRequestMonitorConfFileDownload {
	s.IgnoreCertificateAuthError = &v
	return s
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfFileDownload) SetIgnoreCertificateCanceledError(v int32) *CreateTimingSyntheticTaskRequestMonitorConfFileDownload {
	s.IgnoreCertificateCanceledError = &v
	return s
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfFileDownload) SetIgnoreCertificateOutOfDateError(v int32) *CreateTimingSyntheticTaskRequestMonitorConfFileDownload {
	s.IgnoreCertificateOutOfDateError = &v
	return s
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfFileDownload) SetIgnoreCertificateStatusError(v int32) *CreateTimingSyntheticTaskRequestMonitorConfFileDownload {
	s.IgnoreCertificateStatusError = &v
	return s
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfFileDownload) SetIgnoreCertificateUntrustworthyError(v int32) *CreateTimingSyntheticTaskRequestMonitorConfFileDownload {
	s.IgnoreCertificateUntrustworthyError = &v
	return s
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfFileDownload) SetIgnoreCertificateUsingError(v int32) *CreateTimingSyntheticTaskRequestMonitorConfFileDownload {
	s.IgnoreCertificateUsingError = &v
	return s
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfFileDownload) SetIgnoreInvalidHostError(v int32) *CreateTimingSyntheticTaskRequestMonitorConfFileDownload {
	s.IgnoreInvalidHostError = &v
	return s
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfFileDownload) SetMonitorTimeout(v int64) *CreateTimingSyntheticTaskRequestMonitorConfFileDownload {
	s.MonitorTimeout = &v
	return s
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfFileDownload) SetQuickProtocol(v int32) *CreateTimingSyntheticTaskRequestMonitorConfFileDownload {
	s.QuickProtocol = &v
	return s
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfFileDownload) SetRedirection(v int32) *CreateTimingSyntheticTaskRequestMonitorConfFileDownload {
	s.Redirection = &v
	return s
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfFileDownload) SetTargetUrl(v string) *CreateTimingSyntheticTaskRequestMonitorConfFileDownload {
	s.TargetUrl = &v
	return s
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfFileDownload) SetTransmissionSize(v int64) *CreateTimingSyntheticTaskRequestMonitorConfFileDownload {
	s.TransmissionSize = &v
	return s
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfFileDownload) SetValidateKeywords(v string) *CreateTimingSyntheticTaskRequestMonitorConfFileDownload {
	s.ValidateKeywords = &v
	return s
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfFileDownload) SetVerifyWay(v int32) *CreateTimingSyntheticTaskRequestMonitorConfFileDownload {
	s.VerifyWay = &v
	return s
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfFileDownload) SetWhiteList(v string) *CreateTimingSyntheticTaskRequestMonitorConfFileDownload {
	s.WhiteList = &v
	return s
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfFileDownload) Validate() error {
	return dara.Validate(s)
}

type CreateTimingSyntheticTaskRequestMonitorConfNetDNS struct {
	// The IP version of the DNS server.
	//
	// 	- 0 (default): IPv4.
	//
	// 	- 1: IPv6.
	//
	// 	- 2: A version is automatically selected.
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
	// The DNS query method. Valid values:
	//
	// 	- 0 (default): recursive
	//
	// 	- 1: iterative
	//
	// example:
	//
	// 0
	QueryMethod *int32 `json:"QueryMethod,omitempty" xml:"QueryMethod,omitempty"`
	// The destination domain name.
	//
	// This parameter is required.
	//
	// example:
	//
	// www.aliyun.com
	TargetUrl *string `json:"TargetUrl,omitempty" xml:"TargetUrl,omitempty"`
	// The timeout period for the DNS synthetic test. Unit: milliseconds. Minimum value: 1000. Maximum value: 45000. Default value: 5000.
	//
	// example:
	//
	// 5000
	Timeout *int64 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s CreateTimingSyntheticTaskRequestMonitorConfNetDNS) String() string {
	return dara.Prettify(s)
}

func (s CreateTimingSyntheticTaskRequestMonitorConfNetDNS) GoString() string {
	return s.String()
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfNetDNS) GetDnsServerIpType() *int32 {
	return s.DnsServerIpType
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfNetDNS) GetNsServer() *string {
	return s.NsServer
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfNetDNS) GetQueryMethod() *int32 {
	return s.QueryMethod
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfNetDNS) GetTargetUrl() *string {
	return s.TargetUrl
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfNetDNS) GetTimeout() *int64 {
	return s.Timeout
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfNetDNS) SetDnsServerIpType(v int32) *CreateTimingSyntheticTaskRequestMonitorConfNetDNS {
	s.DnsServerIpType = &v
	return s
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfNetDNS) SetNsServer(v string) *CreateTimingSyntheticTaskRequestMonitorConfNetDNS {
	s.NsServer = &v
	return s
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfNetDNS) SetQueryMethod(v int32) *CreateTimingSyntheticTaskRequestMonitorConfNetDNS {
	s.QueryMethod = &v
	return s
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfNetDNS) SetTargetUrl(v string) *CreateTimingSyntheticTaskRequestMonitorConfNetDNS {
	s.TargetUrl = &v
	return s
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfNetDNS) SetTimeout(v int64) *CreateTimingSyntheticTaskRequestMonitorConfNetDNS {
	s.Timeout = &v
	return s
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfNetDNS) Validate() error {
	return dara.Validate(s)
}

type CreateTimingSyntheticTaskRequestMonitorConfNetICMP struct {
	// The interval at which ICMP packets are sent. Unit: milliseconds. Minimum value: 200. Maximum value: 2000. Default value: 200.
	//
	// example:
	//
	// 200
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
	// The destination IP address or domain name.
	//
	// This parameter is required.
	//
	// example:
	//
	// www.example.com
	TargetUrl *string `json:"TargetUrl,omitempty" xml:"TargetUrl,omitempty"`
	// The timeout period for the ICMP synthetic test. Unit: milliseconds. Minimum value: 1000. Maximum value: 300000. Default value: 20000.
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

func (s CreateTimingSyntheticTaskRequestMonitorConfNetICMP) String() string {
	return dara.Prettify(s)
}

func (s CreateTimingSyntheticTaskRequestMonitorConfNetICMP) GoString() string {
	return s.String()
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfNetICMP) GetInterval() *int64 {
	return s.Interval
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfNetICMP) GetPackageNum() *int32 {
	return s.PackageNum
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfNetICMP) GetPackageSize() *int32 {
	return s.PackageSize
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfNetICMP) GetSplitPackage() *bool {
	return s.SplitPackage
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfNetICMP) GetTargetUrl() *string {
	return s.TargetUrl
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfNetICMP) GetTimeout() *int64 {
	return s.Timeout
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfNetICMP) GetTracertEnable() *bool {
	return s.TracertEnable
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfNetICMP) GetTracertNumMax() *int32 {
	return s.TracertNumMax
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfNetICMP) GetTracertTimeout() *int64 {
	return s.TracertTimeout
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfNetICMP) SetInterval(v int64) *CreateTimingSyntheticTaskRequestMonitorConfNetICMP {
	s.Interval = &v
	return s
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfNetICMP) SetPackageNum(v int32) *CreateTimingSyntheticTaskRequestMonitorConfNetICMP {
	s.PackageNum = &v
	return s
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfNetICMP) SetPackageSize(v int32) *CreateTimingSyntheticTaskRequestMonitorConfNetICMP {
	s.PackageSize = &v
	return s
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfNetICMP) SetSplitPackage(v bool) *CreateTimingSyntheticTaskRequestMonitorConfNetICMP {
	s.SplitPackage = &v
	return s
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfNetICMP) SetTargetUrl(v string) *CreateTimingSyntheticTaskRequestMonitorConfNetICMP {
	s.TargetUrl = &v
	return s
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfNetICMP) SetTimeout(v int64) *CreateTimingSyntheticTaskRequestMonitorConfNetICMP {
	s.Timeout = &v
	return s
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfNetICMP) SetTracertEnable(v bool) *CreateTimingSyntheticTaskRequestMonitorConfNetICMP {
	s.TracertEnable = &v
	return s
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfNetICMP) SetTracertNumMax(v int32) *CreateTimingSyntheticTaskRequestMonitorConfNetICMP {
	s.TracertNumMax = &v
	return s
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfNetICMP) SetTracertTimeout(v int64) *CreateTimingSyntheticTaskRequestMonitorConfNetICMP {
	s.TracertTimeout = &v
	return s
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfNetICMP) Validate() error {
	return dara.Validate(s)
}

type CreateTimingSyntheticTaskRequestMonitorConfNetTCP struct {
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
	// 200
	Interval *int64 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The IP address of the destination host.
	//
	// This parameter is required.
	//
	// example:
	//
	// www.example.com
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

func (s CreateTimingSyntheticTaskRequestMonitorConfNetTCP) String() string {
	return dara.Prettify(s)
}

func (s CreateTimingSyntheticTaskRequestMonitorConfNetTCP) GoString() string {
	return s.String()
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfNetTCP) GetConnectTimes() *int32 {
	return s.ConnectTimes
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfNetTCP) GetInterval() *int64 {
	return s.Interval
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfNetTCP) GetTargetUrl() *string {
	return s.TargetUrl
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfNetTCP) GetTimeout() *int64 {
	return s.Timeout
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfNetTCP) GetTracertEnable() *bool {
	return s.TracertEnable
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfNetTCP) GetTracertNumMax() *int32 {
	return s.TracertNumMax
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfNetTCP) GetTracertTimeout() *int64 {
	return s.TracertTimeout
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfNetTCP) SetConnectTimes(v int32) *CreateTimingSyntheticTaskRequestMonitorConfNetTCP {
	s.ConnectTimes = &v
	return s
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfNetTCP) SetInterval(v int64) *CreateTimingSyntheticTaskRequestMonitorConfNetTCP {
	s.Interval = &v
	return s
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfNetTCP) SetTargetUrl(v string) *CreateTimingSyntheticTaskRequestMonitorConfNetTCP {
	s.TargetUrl = &v
	return s
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfNetTCP) SetTimeout(v int64) *CreateTimingSyntheticTaskRequestMonitorConfNetTCP {
	s.Timeout = &v
	return s
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfNetTCP) SetTracertEnable(v bool) *CreateTimingSyntheticTaskRequestMonitorConfNetTCP {
	s.TracertEnable = &v
	return s
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfNetTCP) SetTracertNumMax(v int32) *CreateTimingSyntheticTaskRequestMonitorConfNetTCP {
	s.TracertNumMax = &v
	return s
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfNetTCP) SetTracertTimeout(v int64) *CreateTimingSyntheticTaskRequestMonitorConfNetTCP {
	s.TracertTimeout = &v
	return s
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfNetTCP) Validate() error {
	return dara.Validate(s)
}

type CreateTimingSyntheticTaskRequestMonitorConfStream struct {
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
	// 12
	PlayerType *int32 `json:"PlayerType,omitempty" xml:"PlayerType,omitempty"`
	// The address type of the resource. Valid values:
	//
	// 	- 1: resource URL
	//
	// 	- 0 (default): page URL
	//
	// example:
	//
	// 0
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

func (s CreateTimingSyntheticTaskRequestMonitorConfStream) String() string {
	return dara.Prettify(s)
}

func (s CreateTimingSyntheticTaskRequestMonitorConfStream) GoString() string {
	return s.String()
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfStream) GetCustomHeaderContent() map[string]*string {
	return s.CustomHeaderContent
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfStream) GetPlayerType() *int32 {
	return s.PlayerType
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfStream) GetStreamAddressType() *int32 {
	return s.StreamAddressType
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfStream) GetStreamMonitorTimeout() *int32 {
	return s.StreamMonitorTimeout
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfStream) GetStreamType() *int32 {
	return s.StreamType
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfStream) GetTargetUrl() *string {
	return s.TargetUrl
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfStream) GetWhiteList() *string {
	return s.WhiteList
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfStream) SetCustomHeaderContent(v map[string]*string) *CreateTimingSyntheticTaskRequestMonitorConfStream {
	s.CustomHeaderContent = v
	return s
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfStream) SetPlayerType(v int32) *CreateTimingSyntheticTaskRequestMonitorConfStream {
	s.PlayerType = &v
	return s
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfStream) SetStreamAddressType(v int32) *CreateTimingSyntheticTaskRequestMonitorConfStream {
	s.StreamAddressType = &v
	return s
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfStream) SetStreamMonitorTimeout(v int32) *CreateTimingSyntheticTaskRequestMonitorConfStream {
	s.StreamMonitorTimeout = &v
	return s
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfStream) SetStreamType(v int32) *CreateTimingSyntheticTaskRequestMonitorConfStream {
	s.StreamType = &v
	return s
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfStream) SetTargetUrl(v string) *CreateTimingSyntheticTaskRequestMonitorConfStream {
	s.TargetUrl = &v
	return s
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfStream) SetWhiteList(v string) *CreateTimingSyntheticTaskRequestMonitorConfStream {
	s.WhiteList = &v
	return s
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfStream) Validate() error {
	return dara.Validate(s)
}

type CreateTimingSyntheticTaskRequestMonitorConfWebsite struct {
	// Specifies whether to automatically scroll up and down the screen to load a page.
	//
	// 	- 0 (default): no
	//
	// 	- 1: yes
	//
	// example:
	//
	// 0
	AutomaticScrolling *int32 `json:"AutomaticScrolling,omitempty" xml:"AutomaticScrolling,omitempty"`
	// Specifies whether to create a custom header.
	//
	// 	- 0 (default): No custom header is created.
	//
	// 	- 1: A custom header is created for the first packet.
	//
	// 	- 2: A custom header is created for all packets.
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
	// Specifies whether to disable caching.
	//
	// 	- 0: no
	//
	// 	- 1 (default): yes
	//
	// example:
	//
	// 1
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
	// www.example.com/a.jpg
	ElementBlacklist *string `json:"ElementBlacklist,omitempty" xml:"ElementBlacklist,omitempty"`
	// Specifies whether to exclude invalid IP addresses. Valid values: 0: yes. 1: no. Default value: 0.
	//
	// example:
	//
	// 0
	FilterInvalidIP *int32 `json:"FilterInvalidIP,omitempty" xml:"FilterInvalidIP,omitempty"`
	// The total number of elements on the page.
	//
	// example:
	//
	// 10
	FlowHijackJumpTimes *int32 `json:"FlowHijackJumpTimes,omitempty" xml:"FlowHijackJumpTimes,omitempty"`
	// The keyword that is used to identify hijacking. Asterisks (\\*) are allowed.
	//
	// example:
	//
	// aliyun
	FlowHijackLogo *string `json:"FlowHijackLogo,omitempty" xml:"FlowHijackLogo,omitempty"`
	// Specifies whether to ignore certificate errors during certificate verification in the SSL handshake process and continue browsing. Valid values: 0: no. 1: yes. Default value: 1.
	//
	// example:
	//
	// 1
	IgnoreCertificateError *int32 `json:"IgnoreCertificateError,omitempty" xml:"IgnoreCertificateError,omitempty"`
	// The monitoring timeout period. Unit: milliseconds. This parameter is optional. Default value: 20000.
	//
	// example:
	//
	// 20000
	MonitorTimeout *int64 `json:"MonitorTimeout,omitempty" xml:"MonitorTimeout,omitempty"`
	// Elements that are not included in the whitelist and appear on the page are tampered with. These elements can be pop-up ads, floating ads, and page redirection.
	//
	// Format: \\<domain name>:\\<elements>. The elements can be wildcard masks. Separate multiple elements with vertical bars (|). Example: www.aliyun.com:|/cc/bb/a.gif|/vv/bb/cc.jpg. It indicates that all elements that belong to the www.aliyun.com domain name except the basic documents, /cc/bb/a.gif, and /vv/bb/cc.jpg are tampered with.
	//
	// example:
	//
	// www.aliyun.com:|/cc/bb/a.gif|/vv/bb/cc.jpg
	PageTamper *string `json:"PageTamper,omitempty" xml:"PageTamper,omitempty"`
	// Specifies whether to continue browsing after redirection. Valid values: 0: no. 1: yes. Default value: 1.
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
	// The URL of the website.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://www.aliyun.com
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

func (s CreateTimingSyntheticTaskRequestMonitorConfWebsite) String() string {
	return dara.Prettify(s)
}

func (s CreateTimingSyntheticTaskRequestMonitorConfWebsite) GoString() string {
	return s.String()
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfWebsite) GetAutomaticScrolling() *int32 {
	return s.AutomaticScrolling
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfWebsite) GetCustomHeader() *int32 {
	return s.CustomHeader
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfWebsite) GetCustomHeaderContent() map[string]*string {
	return s.CustomHeaderContent
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfWebsite) GetDNSHijackWhitelist() *string {
	return s.DNSHijackWhitelist
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfWebsite) GetDisableCache() *int32 {
	return s.DisableCache
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfWebsite) GetDisableCompression() *int32 {
	return s.DisableCompression
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfWebsite) GetElementBlacklist() *string {
	return s.ElementBlacklist
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfWebsite) GetFilterInvalidIP() *int32 {
	return s.FilterInvalidIP
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfWebsite) GetFlowHijackJumpTimes() *int32 {
	return s.FlowHijackJumpTimes
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfWebsite) GetFlowHijackLogo() *string {
	return s.FlowHijackLogo
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfWebsite) GetIgnoreCertificateError() *int32 {
	return s.IgnoreCertificateError
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfWebsite) GetMonitorTimeout() *int64 {
	return s.MonitorTimeout
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfWebsite) GetPageTamper() *string {
	return s.PageTamper
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfWebsite) GetRedirection() *int32 {
	return s.Redirection
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfWebsite) GetSlowElementThreshold() *int64 {
	return s.SlowElementThreshold
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfWebsite) GetTargetUrl() *string {
	return s.TargetUrl
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfWebsite) GetVerifyStringBlacklist() *string {
	return s.VerifyStringBlacklist
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfWebsite) GetVerifyStringWhitelist() *string {
	return s.VerifyStringWhitelist
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfWebsite) GetWaitCompletionTime() *int64 {
	return s.WaitCompletionTime
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfWebsite) SetAutomaticScrolling(v int32) *CreateTimingSyntheticTaskRequestMonitorConfWebsite {
	s.AutomaticScrolling = &v
	return s
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfWebsite) SetCustomHeader(v int32) *CreateTimingSyntheticTaskRequestMonitorConfWebsite {
	s.CustomHeader = &v
	return s
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfWebsite) SetCustomHeaderContent(v map[string]*string) *CreateTimingSyntheticTaskRequestMonitorConfWebsite {
	s.CustomHeaderContent = v
	return s
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfWebsite) SetDNSHijackWhitelist(v string) *CreateTimingSyntheticTaskRequestMonitorConfWebsite {
	s.DNSHijackWhitelist = &v
	return s
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfWebsite) SetDisableCache(v int32) *CreateTimingSyntheticTaskRequestMonitorConfWebsite {
	s.DisableCache = &v
	return s
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfWebsite) SetDisableCompression(v int32) *CreateTimingSyntheticTaskRequestMonitorConfWebsite {
	s.DisableCompression = &v
	return s
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfWebsite) SetElementBlacklist(v string) *CreateTimingSyntheticTaskRequestMonitorConfWebsite {
	s.ElementBlacklist = &v
	return s
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfWebsite) SetFilterInvalidIP(v int32) *CreateTimingSyntheticTaskRequestMonitorConfWebsite {
	s.FilterInvalidIP = &v
	return s
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfWebsite) SetFlowHijackJumpTimes(v int32) *CreateTimingSyntheticTaskRequestMonitorConfWebsite {
	s.FlowHijackJumpTimes = &v
	return s
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfWebsite) SetFlowHijackLogo(v string) *CreateTimingSyntheticTaskRequestMonitorConfWebsite {
	s.FlowHijackLogo = &v
	return s
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfWebsite) SetIgnoreCertificateError(v int32) *CreateTimingSyntheticTaskRequestMonitorConfWebsite {
	s.IgnoreCertificateError = &v
	return s
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfWebsite) SetMonitorTimeout(v int64) *CreateTimingSyntheticTaskRequestMonitorConfWebsite {
	s.MonitorTimeout = &v
	return s
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfWebsite) SetPageTamper(v string) *CreateTimingSyntheticTaskRequestMonitorConfWebsite {
	s.PageTamper = &v
	return s
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfWebsite) SetRedirection(v int32) *CreateTimingSyntheticTaskRequestMonitorConfWebsite {
	s.Redirection = &v
	return s
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfWebsite) SetSlowElementThreshold(v int64) *CreateTimingSyntheticTaskRequestMonitorConfWebsite {
	s.SlowElementThreshold = &v
	return s
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfWebsite) SetTargetUrl(v string) *CreateTimingSyntheticTaskRequestMonitorConfWebsite {
	s.TargetUrl = &v
	return s
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfWebsite) SetVerifyStringBlacklist(v string) *CreateTimingSyntheticTaskRequestMonitorConfWebsite {
	s.VerifyStringBlacklist = &v
	return s
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfWebsite) SetVerifyStringWhitelist(v string) *CreateTimingSyntheticTaskRequestMonitorConfWebsite {
	s.VerifyStringWhitelist = &v
	return s
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfWebsite) SetWaitCompletionTime(v int64) *CreateTimingSyntheticTaskRequestMonitorConfWebsite {
	s.WaitCompletionTime = &v
	return s
}

func (s *CreateTimingSyntheticTaskRequestMonitorConfWebsite) Validate() error {
	return dara.Validate(s)
}

type CreateTimingSyntheticTaskRequestMonitors struct {
	// The city code.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100023
	CityCode *string `json:"CityCode,omitempty" xml:"CityCode,omitempty"`
	// The client type of the detection point. Valid values:
	//
	// - 1: data center
	//
	// - 2: Internet
	//
	// - 3: mobile device
	//
	// - 4: ECS instance
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ClientType *int32 `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	// The carrier code.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	OperatorCode *string `json:"OperatorCode,omitempty" xml:"OperatorCode,omitempty"`
}

func (s CreateTimingSyntheticTaskRequestMonitors) String() string {
	return dara.Prettify(s)
}

func (s CreateTimingSyntheticTaskRequestMonitors) GoString() string {
	return s.String()
}

func (s *CreateTimingSyntheticTaskRequestMonitors) GetCityCode() *string {
	return s.CityCode
}

func (s *CreateTimingSyntheticTaskRequestMonitors) GetClientType() *int32 {
	return s.ClientType
}

func (s *CreateTimingSyntheticTaskRequestMonitors) GetOperatorCode() *string {
	return s.OperatorCode
}

func (s *CreateTimingSyntheticTaskRequestMonitors) SetCityCode(v string) *CreateTimingSyntheticTaskRequestMonitors {
	s.CityCode = &v
	return s
}

func (s *CreateTimingSyntheticTaskRequestMonitors) SetClientType(v int32) *CreateTimingSyntheticTaskRequestMonitors {
	s.ClientType = &v
	return s
}

func (s *CreateTimingSyntheticTaskRequestMonitors) SetOperatorCode(v string) *CreateTimingSyntheticTaskRequestMonitors {
	s.OperatorCode = &v
	return s
}

func (s *CreateTimingSyntheticTaskRequestMonitors) Validate() error {
	return dara.Validate(s)
}

type CreateTimingSyntheticTaskRequestTags struct {
	// The key of the tag.
	//
	// example:
	//
	// Key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// 500
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateTimingSyntheticTaskRequestTags) String() string {
	return dara.Prettify(s)
}

func (s CreateTimingSyntheticTaskRequestTags) GoString() string {
	return s.String()
}

func (s *CreateTimingSyntheticTaskRequestTags) GetKey() *string {
	return s.Key
}

func (s *CreateTimingSyntheticTaskRequestTags) GetValue() *string {
	return s.Value
}

func (s *CreateTimingSyntheticTaskRequestTags) SetKey(v string) *CreateTimingSyntheticTaskRequestTags {
	s.Key = &v
	return s
}

func (s *CreateTimingSyntheticTaskRequestTags) SetValue(v string) *CreateTimingSyntheticTaskRequestTags {
	s.Value = &v
	return s
}

func (s *CreateTimingSyntheticTaskRequestTags) Validate() error {
	return dara.Validate(s)
}
