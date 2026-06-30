// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEndpointGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAcceleratorId(v string) *DescribeEndpointGroupResponseBody
	GetAcceleratorId() *string
	SetAccessLogRecordCustomizedHeaderList(v []*string) *DescribeEndpointGroupResponseBody
	GetAccessLogRecordCustomizedHeaderList() []*string
	SetAccessLogRecordCustomizedHeadersEnabled(v bool) *DescribeEndpointGroupResponseBody
	GetAccessLogRecordCustomizedHeadersEnabled() *bool
	SetAccessLogSwitch(v string) *DescribeEndpointGroupResponseBody
	GetAccessLogSwitch() *string
	SetDescription(v string) *DescribeEndpointGroupResponseBody
	GetDescription() *string
	SetEnableAccessLog(v bool) *DescribeEndpointGroupResponseBody
	GetEnableAccessLog() *bool
	SetEndpointConfigurations(v []*DescribeEndpointGroupResponseBodyEndpointConfigurations) *DescribeEndpointGroupResponseBody
	GetEndpointConfigurations() []*DescribeEndpointGroupResponseBodyEndpointConfigurations
	SetEndpointGroupId(v string) *DescribeEndpointGroupResponseBody
	GetEndpointGroupId() *string
	SetEndpointGroupIpList(v []*string) *DescribeEndpointGroupResponseBody
	GetEndpointGroupIpList() []*string
	SetEndpointGroupRegion(v string) *DescribeEndpointGroupResponseBody
	GetEndpointGroupRegion() *string
	SetEndpointGroupType(v string) *DescribeEndpointGroupResponseBody
	GetEndpointGroupType() *string
	SetEndpointGroupUnconfirmedIpList(v []*string) *DescribeEndpointGroupResponseBody
	GetEndpointGroupUnconfirmedIpList() []*string
	SetEndpointIpVersion(v string) *DescribeEndpointGroupResponseBody
	GetEndpointIpVersion() *string
	SetEndpointPrivateIpList(v []*DescribeEndpointGroupResponseBodyEndpointPrivateIpList) *DescribeEndpointGroupResponseBody
	GetEndpointPrivateIpList() []*DescribeEndpointGroupResponseBodyEndpointPrivateIpList
	SetEndpointProtocolVersion(v string) *DescribeEndpointGroupResponseBody
	GetEndpointProtocolVersion() *string
	SetEndpointRequestProtocol(v string) *DescribeEndpointGroupResponseBody
	GetEndpointRequestProtocol() *string
	SetForwardingRuleIds(v []*string) *DescribeEndpointGroupResponseBody
	GetForwardingRuleIds() []*string
	SetHealthCheckEnabled(v bool) *DescribeEndpointGroupResponseBody
	GetHealthCheckEnabled() *bool
	SetHealthCheckHost(v string) *DescribeEndpointGroupResponseBody
	GetHealthCheckHost() *string
	SetHealthCheckIntervalSeconds(v int32) *DescribeEndpointGroupResponseBody
	GetHealthCheckIntervalSeconds() *int32
	SetHealthCheckPath(v string) *DescribeEndpointGroupResponseBody
	GetHealthCheckPath() *string
	SetHealthCheckPort(v int32) *DescribeEndpointGroupResponseBody
	GetHealthCheckPort() *int32
	SetHealthCheckProtocol(v string) *DescribeEndpointGroupResponseBody
	GetHealthCheckProtocol() *string
	SetListenerId(v string) *DescribeEndpointGroupResponseBody
	GetListenerId() *string
	SetName(v string) *DescribeEndpointGroupResponseBody
	GetName() *string
	SetPortOverrides(v []*DescribeEndpointGroupResponseBodyPortOverrides) *DescribeEndpointGroupResponseBody
	GetPortOverrides() []*DescribeEndpointGroupResponseBodyPortOverrides
	SetRequestId(v string) *DescribeEndpointGroupResponseBody
	GetRequestId() *string
	SetServiceId(v string) *DescribeEndpointGroupResponseBody
	GetServiceId() *string
	SetServiceManaged(v bool) *DescribeEndpointGroupResponseBody
	GetServiceManaged() *bool
	SetServiceManagedInfos(v []*DescribeEndpointGroupResponseBodyServiceManagedInfos) *DescribeEndpointGroupResponseBody
	GetServiceManagedInfos() []*DescribeEndpointGroupResponseBodyServiceManagedInfos
	SetSlsLogStoreName(v string) *DescribeEndpointGroupResponseBody
	GetSlsLogStoreName() *string
	SetSlsProjectName(v string) *DescribeEndpointGroupResponseBody
	GetSlsProjectName() *string
	SetSlsRegion(v string) *DescribeEndpointGroupResponseBody
	GetSlsRegion() *string
	SetState(v string) *DescribeEndpointGroupResponseBody
	GetState() *string
	SetTags(v []*DescribeEndpointGroupResponseBodyTags) *DescribeEndpointGroupResponseBody
	GetTags() []*DescribeEndpointGroupResponseBodyTags
	SetThresholdCount(v int32) *DescribeEndpointGroupResponseBody
	GetThresholdCount() *int32
	SetTrafficPercentage(v int32) *DescribeEndpointGroupResponseBody
	GetTrafficPercentage() *int32
}

type DescribeEndpointGroupResponseBody struct {
	// The Global Accelerator instance ID.
	//
	// example:
	//
	// ga-bp1odcab8tmno0hdq****
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	// The custom header fields to record in access logs.
	AccessLogRecordCustomizedHeaderList []*string `json:"AccessLogRecordCustomizedHeaderList,omitempty" xml:"AccessLogRecordCustomizedHeaderList,omitempty" type:"Repeated"`
	// Specifies whether to record custom header fields in access logs. Valid values:
	//
	// - **true**: Yes.
	//
	// - **false*	- (default): No.
	//
	// > You can set this parameter to **true*	- only when **EnableAccessLog*	- is set to **true**.
	AccessLogRecordCustomizedHeadersEnabled *bool `json:"AccessLogRecordCustomizedHeadersEnabled,omitempty" xml:"AccessLogRecordCustomizedHeadersEnabled,omitempty"`
	// The status of the access log configuration. Valid values:
	//
	// - **on**: The access log is configured.
	//
	// - **off**: The access log is not configured.
	//
	// - **binding**: The access log is being configured.
	//
	// - **unbinding**: The access log configuration is being removed.
	//
	// example:
	//
	// on
	AccessLogSwitch *string `json:"AccessLogSwitch,omitempty" xml:"AccessLogSwitch,omitempty"`
	// The description of the endpoint group.
	//
	// example:
	//
	// group1
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Indicates whether access logging is enabled.
	//
	// - **true**: Access logging is enabled.
	//
	// - **false**: Access logging is disabled.
	//
	// example:
	//
	// true
	EnableAccessLog *bool `json:"EnableAccessLog,omitempty" xml:"EnableAccessLog,omitempty"`
	// The endpoint configurations.
	EndpointConfigurations []*DescribeEndpointGroupResponseBodyEndpointConfigurations `json:"EndpointConfigurations,omitempty" xml:"EndpointConfigurations,omitempty" type:"Repeated"`
	// The endpoint group ID.
	//
	// example:
	//
	// epg-bp14sz7ftcwwjgrdm****
	EndpointGroupId *string `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
	// The list of active IP addresses of the endpoints in the endpoint group.
	//
	// 	Notice: For an endpoint group configured for back-to-source from a private network, the console displays only the private back-to-source IP addresses and not the public IP addresses. If the network connection type for the backend service of the endpoint group changes (for example, from a private network to a public network, or to a mix of private and public networks), you must monitor the changes in the back-to-source IP addresses and update the access control list (ACL) of the backend service accordingly.
	EndpointGroupIpList []*string `json:"EndpointGroupIpList,omitempty" xml:"EndpointGroupIpList,omitempty" type:"Repeated"`
	// The region ID where the endpoint group is deployed.
	//
	// example:
	//
	// cn-hangzhou
	EndpointGroupRegion *string `json:"EndpointGroupRegion,omitempty" xml:"EndpointGroupRegion,omitempty"`
	// The type of the endpoint group. Valid values:
	//
	// - **default**: A default endpoint group.
	//
	// - **virtual**: A virtual endpoint group.
	//
	// example:
	//
	// default
	EndpointGroupType *string `json:"EndpointGroupType,omitempty" xml:"EndpointGroupType,omitempty"`
	// A list of endpoint IP addresses pending confirmation after a Global Accelerator instance upgrade.
	EndpointGroupUnconfirmedIpList []*string `json:"EndpointGroupUnconfirmedIpList,omitempty" xml:"EndpointGroupUnconfirmedIpList,omitempty" type:"Repeated"`
	// The IP version used to connect to the backend service. Valid values:
	//
	// - **IPv4*	- (default): Global Accelerator connects to the backend service using IPv4.
	//
	// - **IPv6**: Global Accelerator connects to the backend service using IPv6.
	//
	// - **ProtocolAffinity**: Global Accelerator connects to the backend service using the same IP version as the client request.
	//
	// example:
	//
	// IPv4
	EndpointIpVersion *string `json:"EndpointIpVersion,omitempty" xml:"EndpointIpVersion,omitempty"`
	// The private IP addresses of the endpoints.
	EndpointPrivateIpList []*DescribeEndpointGroupResponseBodyEndpointPrivateIpList `json:"EndpointPrivateIpList,omitempty" xml:"EndpointPrivateIpList,omitempty" type:"Repeated"`
	// The version of the backend service protocol. Valid values:
	//
	// - **HTTP1.1**
	//
	// - **HTTP2**
	//
	// example:
	//
	// HTTP2
	EndpointProtocolVersion *string `json:"EndpointProtocolVersion,omitempty" xml:"EndpointProtocolVersion,omitempty"`
	// The protocol used by the backend service. Valid values:
	//
	// - **HTTP**
	//
	// - **HTTPS**
	//
	// example:
	//
	// HTTP
	EndpointRequestProtocol *string `json:"EndpointRequestProtocol,omitempty" xml:"EndpointRequestProtocol,omitempty"`
	// The IDs of the associated forwarding rules.
	ForwardingRuleIds []*string `json:"ForwardingRuleIds,omitempty" xml:"ForwardingRuleIds,omitempty" type:"Repeated"`
	// Indicates whether health checks are enabled.
	//
	// - **true**: Health checks are enabled.
	//
	// - **false**: Health checks are disabled.
	//
	// example:
	//
	// true
	HealthCheckEnabled *bool `json:"HealthCheckEnabled,omitempty" xml:"HealthCheckEnabled,omitempty"`
	// The domain name used for health checks.
	//
	// example:
	//
	// www.taobao.com
	HealthCheckHost *string `json:"HealthCheckHost,omitempty" xml:"HealthCheckHost,omitempty"`
	// The health check interval, in seconds.
	//
	// example:
	//
	// 3
	HealthCheckIntervalSeconds *int32 `json:"HealthCheckIntervalSeconds,omitempty" xml:"HealthCheckIntervalSeconds,omitempty"`
	// The path for health check probes.
	//
	// example:
	//
	// /healthcheck
	HealthCheckPath *string `json:"HealthCheckPath,omitempty" xml:"HealthCheckPath,omitempty"`
	// The port used for health checks.
	//
	// example:
	//
	// 20
	HealthCheckPort *int32 `json:"HealthCheckPort,omitempty" xml:"HealthCheckPort,omitempty"`
	// The protocol used for health checks.
	//
	// - **tcp*	- or **TCP**: TCP
	//
	// - **http*	- or **HTTP**: HTTP
	//
	// - **https*	- or **HTTPS**: HTTPS
	//
	// example:
	//
	// tcp
	HealthCheckProtocol *string `json:"HealthCheckProtocol,omitempty" xml:"HealthCheckProtocol,omitempty"`
	// The listener ID.
	//
	// example:
	//
	// lsr-bp1bpn0kn908w4nbw****
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// The name of the endpoint group.
	//
	// example:
	//
	// group1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The port mapping configuration.
	PortOverrides []*DescribeEndpointGroupResponseBodyPortOverrides `json:"PortOverrides,omitempty" xml:"PortOverrides,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 6FEA0CF3-D3B9-43E5-A304-D217037876A8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the service that manages the instance.
	//
	// > This parameter is returned only if **ServiceManaged*	- is set to **True**.
	//
	// example:
	//
	// ALB
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// Indicates whether the instance is a managed instance. Valid values:
	//
	// - **true**: The instance is a managed instance.
	//
	// - **false**: The instance is not a managed instance.
	//
	// example:
	//
	// true
	ServiceManaged *bool `json:"ServiceManaged,omitempty" xml:"ServiceManaged,omitempty"`
	// A list of management states for actions that can be performed on the instance.
	//
	// > - This parameter is returned only if **ServiceManaged*	- is set to **True**.
	//
	// >
	//
	// > - When an instance is managed, some operations may be restricted.
	ServiceManagedInfos []*DescribeEndpointGroupResponseBodyServiceManagedInfos `json:"ServiceManagedInfos,omitempty" xml:"ServiceManagedInfos,omitempty" type:"Repeated"`
	// The name of the Logstore.
	//
	// example:
	//
	// lsn-01
	SlsLogStoreName *string `json:"SlsLogStoreName,omitempty" xml:"SlsLogStoreName,omitempty"`
	// The name of the Log Service project.
	//
	// example:
	//
	// pn-01
	SlsProjectName *string `json:"SlsProjectName,omitempty" xml:"SlsProjectName,omitempty"`
	// The region of the Log Service project.
	//
	// example:
	//
	// cn-hangzhou
	SlsRegion *string `json:"SlsRegion,omitempty" xml:"SlsRegion,omitempty"`
	// The status of the endpoint group.
	//
	// - **init**: The endpoint group is being initialized.
	//
	// - **active**: The endpoint group is active.
	//
	// - **updating**: The endpoint group is being updated.
	//
	// - **deleting**: The endpoint group is being deleted.
	//
	// example:
	//
	// active
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// A list of tags attached to the endpoint group.
	Tags []*DescribeEndpointGroupResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The number of consecutive failed health checks before an endpoint is marked as unhealthy.
	//
	// example:
	//
	// 3
	ThresholdCount *int32 `json:"ThresholdCount,omitempty" xml:"ThresholdCount,omitempty"`
	// The percentage of traffic that is distributed to the endpoint group. This parameter is returned only when a listener is associated with multiple endpoint groups.
	//
	// example:
	//
	// 100
	TrafficPercentage *int32 `json:"TrafficPercentage,omitempty" xml:"TrafficPercentage,omitempty"`
}

func (s DescribeEndpointGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeEndpointGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEndpointGroupResponseBody) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *DescribeEndpointGroupResponseBody) GetAccessLogRecordCustomizedHeaderList() []*string {
	return s.AccessLogRecordCustomizedHeaderList
}

func (s *DescribeEndpointGroupResponseBody) GetAccessLogRecordCustomizedHeadersEnabled() *bool {
	return s.AccessLogRecordCustomizedHeadersEnabled
}

func (s *DescribeEndpointGroupResponseBody) GetAccessLogSwitch() *string {
	return s.AccessLogSwitch
}

func (s *DescribeEndpointGroupResponseBody) GetDescription() *string {
	return s.Description
}

func (s *DescribeEndpointGroupResponseBody) GetEnableAccessLog() *bool {
	return s.EnableAccessLog
}

func (s *DescribeEndpointGroupResponseBody) GetEndpointConfigurations() []*DescribeEndpointGroupResponseBodyEndpointConfigurations {
	return s.EndpointConfigurations
}

func (s *DescribeEndpointGroupResponseBody) GetEndpointGroupId() *string {
	return s.EndpointGroupId
}

func (s *DescribeEndpointGroupResponseBody) GetEndpointGroupIpList() []*string {
	return s.EndpointGroupIpList
}

func (s *DescribeEndpointGroupResponseBody) GetEndpointGroupRegion() *string {
	return s.EndpointGroupRegion
}

func (s *DescribeEndpointGroupResponseBody) GetEndpointGroupType() *string {
	return s.EndpointGroupType
}

func (s *DescribeEndpointGroupResponseBody) GetEndpointGroupUnconfirmedIpList() []*string {
	return s.EndpointGroupUnconfirmedIpList
}

func (s *DescribeEndpointGroupResponseBody) GetEndpointIpVersion() *string {
	return s.EndpointIpVersion
}

func (s *DescribeEndpointGroupResponseBody) GetEndpointPrivateIpList() []*DescribeEndpointGroupResponseBodyEndpointPrivateIpList {
	return s.EndpointPrivateIpList
}

func (s *DescribeEndpointGroupResponseBody) GetEndpointProtocolVersion() *string {
	return s.EndpointProtocolVersion
}

func (s *DescribeEndpointGroupResponseBody) GetEndpointRequestProtocol() *string {
	return s.EndpointRequestProtocol
}

func (s *DescribeEndpointGroupResponseBody) GetForwardingRuleIds() []*string {
	return s.ForwardingRuleIds
}

func (s *DescribeEndpointGroupResponseBody) GetHealthCheckEnabled() *bool {
	return s.HealthCheckEnabled
}

func (s *DescribeEndpointGroupResponseBody) GetHealthCheckHost() *string {
	return s.HealthCheckHost
}

func (s *DescribeEndpointGroupResponseBody) GetHealthCheckIntervalSeconds() *int32 {
	return s.HealthCheckIntervalSeconds
}

func (s *DescribeEndpointGroupResponseBody) GetHealthCheckPath() *string {
	return s.HealthCheckPath
}

func (s *DescribeEndpointGroupResponseBody) GetHealthCheckPort() *int32 {
	return s.HealthCheckPort
}

func (s *DescribeEndpointGroupResponseBody) GetHealthCheckProtocol() *string {
	return s.HealthCheckProtocol
}

func (s *DescribeEndpointGroupResponseBody) GetListenerId() *string {
	return s.ListenerId
}

func (s *DescribeEndpointGroupResponseBody) GetName() *string {
	return s.Name
}

func (s *DescribeEndpointGroupResponseBody) GetPortOverrides() []*DescribeEndpointGroupResponseBodyPortOverrides {
	return s.PortOverrides
}

func (s *DescribeEndpointGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeEndpointGroupResponseBody) GetServiceId() *string {
	return s.ServiceId
}

func (s *DescribeEndpointGroupResponseBody) GetServiceManaged() *bool {
	return s.ServiceManaged
}

func (s *DescribeEndpointGroupResponseBody) GetServiceManagedInfos() []*DescribeEndpointGroupResponseBodyServiceManagedInfos {
	return s.ServiceManagedInfos
}

func (s *DescribeEndpointGroupResponseBody) GetSlsLogStoreName() *string {
	return s.SlsLogStoreName
}

func (s *DescribeEndpointGroupResponseBody) GetSlsProjectName() *string {
	return s.SlsProjectName
}

func (s *DescribeEndpointGroupResponseBody) GetSlsRegion() *string {
	return s.SlsRegion
}

func (s *DescribeEndpointGroupResponseBody) GetState() *string {
	return s.State
}

func (s *DescribeEndpointGroupResponseBody) GetTags() []*DescribeEndpointGroupResponseBodyTags {
	return s.Tags
}

func (s *DescribeEndpointGroupResponseBody) GetThresholdCount() *int32 {
	return s.ThresholdCount
}

func (s *DescribeEndpointGroupResponseBody) GetTrafficPercentage() *int32 {
	return s.TrafficPercentage
}

func (s *DescribeEndpointGroupResponseBody) SetAcceleratorId(v string) *DescribeEndpointGroupResponseBody {
	s.AcceleratorId = &v
	return s
}

func (s *DescribeEndpointGroupResponseBody) SetAccessLogRecordCustomizedHeaderList(v []*string) *DescribeEndpointGroupResponseBody {
	s.AccessLogRecordCustomizedHeaderList = v
	return s
}

func (s *DescribeEndpointGroupResponseBody) SetAccessLogRecordCustomizedHeadersEnabled(v bool) *DescribeEndpointGroupResponseBody {
	s.AccessLogRecordCustomizedHeadersEnabled = &v
	return s
}

func (s *DescribeEndpointGroupResponseBody) SetAccessLogSwitch(v string) *DescribeEndpointGroupResponseBody {
	s.AccessLogSwitch = &v
	return s
}

func (s *DescribeEndpointGroupResponseBody) SetDescription(v string) *DescribeEndpointGroupResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeEndpointGroupResponseBody) SetEnableAccessLog(v bool) *DescribeEndpointGroupResponseBody {
	s.EnableAccessLog = &v
	return s
}

func (s *DescribeEndpointGroupResponseBody) SetEndpointConfigurations(v []*DescribeEndpointGroupResponseBodyEndpointConfigurations) *DescribeEndpointGroupResponseBody {
	s.EndpointConfigurations = v
	return s
}

func (s *DescribeEndpointGroupResponseBody) SetEndpointGroupId(v string) *DescribeEndpointGroupResponseBody {
	s.EndpointGroupId = &v
	return s
}

func (s *DescribeEndpointGroupResponseBody) SetEndpointGroupIpList(v []*string) *DescribeEndpointGroupResponseBody {
	s.EndpointGroupIpList = v
	return s
}

func (s *DescribeEndpointGroupResponseBody) SetEndpointGroupRegion(v string) *DescribeEndpointGroupResponseBody {
	s.EndpointGroupRegion = &v
	return s
}

func (s *DescribeEndpointGroupResponseBody) SetEndpointGroupType(v string) *DescribeEndpointGroupResponseBody {
	s.EndpointGroupType = &v
	return s
}

func (s *DescribeEndpointGroupResponseBody) SetEndpointGroupUnconfirmedIpList(v []*string) *DescribeEndpointGroupResponseBody {
	s.EndpointGroupUnconfirmedIpList = v
	return s
}

func (s *DescribeEndpointGroupResponseBody) SetEndpointIpVersion(v string) *DescribeEndpointGroupResponseBody {
	s.EndpointIpVersion = &v
	return s
}

func (s *DescribeEndpointGroupResponseBody) SetEndpointPrivateIpList(v []*DescribeEndpointGroupResponseBodyEndpointPrivateIpList) *DescribeEndpointGroupResponseBody {
	s.EndpointPrivateIpList = v
	return s
}

func (s *DescribeEndpointGroupResponseBody) SetEndpointProtocolVersion(v string) *DescribeEndpointGroupResponseBody {
	s.EndpointProtocolVersion = &v
	return s
}

func (s *DescribeEndpointGroupResponseBody) SetEndpointRequestProtocol(v string) *DescribeEndpointGroupResponseBody {
	s.EndpointRequestProtocol = &v
	return s
}

func (s *DescribeEndpointGroupResponseBody) SetForwardingRuleIds(v []*string) *DescribeEndpointGroupResponseBody {
	s.ForwardingRuleIds = v
	return s
}

func (s *DescribeEndpointGroupResponseBody) SetHealthCheckEnabled(v bool) *DescribeEndpointGroupResponseBody {
	s.HealthCheckEnabled = &v
	return s
}

func (s *DescribeEndpointGroupResponseBody) SetHealthCheckHost(v string) *DescribeEndpointGroupResponseBody {
	s.HealthCheckHost = &v
	return s
}

func (s *DescribeEndpointGroupResponseBody) SetHealthCheckIntervalSeconds(v int32) *DescribeEndpointGroupResponseBody {
	s.HealthCheckIntervalSeconds = &v
	return s
}

func (s *DescribeEndpointGroupResponseBody) SetHealthCheckPath(v string) *DescribeEndpointGroupResponseBody {
	s.HealthCheckPath = &v
	return s
}

func (s *DescribeEndpointGroupResponseBody) SetHealthCheckPort(v int32) *DescribeEndpointGroupResponseBody {
	s.HealthCheckPort = &v
	return s
}

func (s *DescribeEndpointGroupResponseBody) SetHealthCheckProtocol(v string) *DescribeEndpointGroupResponseBody {
	s.HealthCheckProtocol = &v
	return s
}

func (s *DescribeEndpointGroupResponseBody) SetListenerId(v string) *DescribeEndpointGroupResponseBody {
	s.ListenerId = &v
	return s
}

func (s *DescribeEndpointGroupResponseBody) SetName(v string) *DescribeEndpointGroupResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeEndpointGroupResponseBody) SetPortOverrides(v []*DescribeEndpointGroupResponseBodyPortOverrides) *DescribeEndpointGroupResponseBody {
	s.PortOverrides = v
	return s
}

func (s *DescribeEndpointGroupResponseBody) SetRequestId(v string) *DescribeEndpointGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEndpointGroupResponseBody) SetServiceId(v string) *DescribeEndpointGroupResponseBody {
	s.ServiceId = &v
	return s
}

func (s *DescribeEndpointGroupResponseBody) SetServiceManaged(v bool) *DescribeEndpointGroupResponseBody {
	s.ServiceManaged = &v
	return s
}

func (s *DescribeEndpointGroupResponseBody) SetServiceManagedInfos(v []*DescribeEndpointGroupResponseBodyServiceManagedInfos) *DescribeEndpointGroupResponseBody {
	s.ServiceManagedInfos = v
	return s
}

func (s *DescribeEndpointGroupResponseBody) SetSlsLogStoreName(v string) *DescribeEndpointGroupResponseBody {
	s.SlsLogStoreName = &v
	return s
}

func (s *DescribeEndpointGroupResponseBody) SetSlsProjectName(v string) *DescribeEndpointGroupResponseBody {
	s.SlsProjectName = &v
	return s
}

func (s *DescribeEndpointGroupResponseBody) SetSlsRegion(v string) *DescribeEndpointGroupResponseBody {
	s.SlsRegion = &v
	return s
}

func (s *DescribeEndpointGroupResponseBody) SetState(v string) *DescribeEndpointGroupResponseBody {
	s.State = &v
	return s
}

func (s *DescribeEndpointGroupResponseBody) SetTags(v []*DescribeEndpointGroupResponseBodyTags) *DescribeEndpointGroupResponseBody {
	s.Tags = v
	return s
}

func (s *DescribeEndpointGroupResponseBody) SetThresholdCount(v int32) *DescribeEndpointGroupResponseBody {
	s.ThresholdCount = &v
	return s
}

func (s *DescribeEndpointGroupResponseBody) SetTrafficPercentage(v int32) *DescribeEndpointGroupResponseBody {
	s.TrafficPercentage = &v
	return s
}

func (s *DescribeEndpointGroupResponseBody) Validate() error {
	if s.EndpointConfigurations != nil {
		for _, item := range s.EndpointConfigurations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.EndpointPrivateIpList != nil {
		for _, item := range s.EndpointPrivateIpList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PortOverrides != nil {
		for _, item := range s.PortOverrides {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ServiceManagedInfos != nil {
		for _, item := range s.ServiceManagedInfos {
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

type DescribeEndpointGroupResponseBodyEndpointConfigurations struct {
	// The API keys for the endpoint configuration.
	ApiKeys []*string `json:"ApiKeys,omitempty" xml:"ApiKeys,omitempty" type:"Repeated"`
	// Indicates whether client IP preservation is enabled by using the automatic method.
	//
	// - **true**: Enabled.
	//
	// - **false**: Disabled.
	//
	// example:
	//
	// false
	EnableClientIPPreservation *bool `json:"EnableClientIPPreservation,omitempty" xml:"EnableClientIPPreservation,omitempty"`
	// Indicates whether client IP preservation is enabled using the Proxy Protocol.
	//
	// - **true**: Enabled.
	//
	// - **false**: Disabled.
	//
	// example:
	//
	// false
	EnableProxyProtocol *bool `json:"EnableProxyProtocol,omitempty" xml:"EnableProxyProtocol,omitempty"`
	// The IP address, domain name, or instance ID of the endpoint.
	//
	// example:
	//
	// 120.XX.XX.21
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// The port used for latency probing.
	//
	// example:
	//
	// 80
	ProbePort *int32 `json:"ProbePort,omitempty" xml:"ProbePort,omitempty"`
	// The protocol for latency probing. Valid values:
	//
	// - **tcp**: TCP
	//
	// - **icmp**: ICMP
	//
	// example:
	//
	// tcp
	ProbeProtocol *string `json:"ProbeProtocol,omitempty" xml:"ProbeProtocol,omitempty"`
	// The provider of the endpoint configuration.
	//
	// example:
	//
	// BAILIAN
	Provider *string `json:"Provider,omitempty" xml:"Provider,omitempty"`
	// The private IP address of the elastic network interface.
	//
	// example:
	//
	// 172.168.XX.XX
	SubAddress *string `json:"SubAddress,omitempty" xml:"SubAddress,omitempty"`
	// The type of the endpoint. Valid values:
	//
	// - **Domain**: A custom domain name.
	//
	// - **Ip**: A custom IP address.
	//
	// - **IpTarget**: A custom private IP address.
	//
	// - **PublicIp**: An Alibaba Cloud public IP address.
	//
	// - **ECS**: An ECS instance.
	//
	// - **SLB**: An SLB instance.
	//
	// - **ALB**: An ALB instance.
	//
	// - **OSS**: An OSS instance.
	//
	// - **ENI**: An elastic network interface.
	//
	// - **NLB**: An NLB instance.
	//
	// example:
	//
	// Ip
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// A list of VSwitch IDs.
	VSwitchIds []*string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
	// The VPC ID.
	//
	// example:
	//
	// vpc-hp30gtnxkfxj1l2xt****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The weight of the endpoint.
	//
	// example:
	//
	// 255
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s DescribeEndpointGroupResponseBodyEndpointConfigurations) String() string {
	return dara.Prettify(s)
}

func (s DescribeEndpointGroupResponseBodyEndpointConfigurations) GoString() string {
	return s.String()
}

func (s *DescribeEndpointGroupResponseBodyEndpointConfigurations) GetApiKeys() []*string {
	return s.ApiKeys
}

func (s *DescribeEndpointGroupResponseBodyEndpointConfigurations) GetEnableClientIPPreservation() *bool {
	return s.EnableClientIPPreservation
}

func (s *DescribeEndpointGroupResponseBodyEndpointConfigurations) GetEnableProxyProtocol() *bool {
	return s.EnableProxyProtocol
}

func (s *DescribeEndpointGroupResponseBodyEndpointConfigurations) GetEndpoint() *string {
	return s.Endpoint
}

func (s *DescribeEndpointGroupResponseBodyEndpointConfigurations) GetProbePort() *int32 {
	return s.ProbePort
}

func (s *DescribeEndpointGroupResponseBodyEndpointConfigurations) GetProbeProtocol() *string {
	return s.ProbeProtocol
}

func (s *DescribeEndpointGroupResponseBodyEndpointConfigurations) GetProvider() *string {
	return s.Provider
}

func (s *DescribeEndpointGroupResponseBodyEndpointConfigurations) GetSubAddress() *string {
	return s.SubAddress
}

func (s *DescribeEndpointGroupResponseBodyEndpointConfigurations) GetType() *string {
	return s.Type
}

func (s *DescribeEndpointGroupResponseBodyEndpointConfigurations) GetVSwitchIds() []*string {
	return s.VSwitchIds
}

func (s *DescribeEndpointGroupResponseBodyEndpointConfigurations) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeEndpointGroupResponseBodyEndpointConfigurations) GetWeight() *int32 {
	return s.Weight
}

func (s *DescribeEndpointGroupResponseBodyEndpointConfigurations) SetApiKeys(v []*string) *DescribeEndpointGroupResponseBodyEndpointConfigurations {
	s.ApiKeys = v
	return s
}

func (s *DescribeEndpointGroupResponseBodyEndpointConfigurations) SetEnableClientIPPreservation(v bool) *DescribeEndpointGroupResponseBodyEndpointConfigurations {
	s.EnableClientIPPreservation = &v
	return s
}

func (s *DescribeEndpointGroupResponseBodyEndpointConfigurations) SetEnableProxyProtocol(v bool) *DescribeEndpointGroupResponseBodyEndpointConfigurations {
	s.EnableProxyProtocol = &v
	return s
}

func (s *DescribeEndpointGroupResponseBodyEndpointConfigurations) SetEndpoint(v string) *DescribeEndpointGroupResponseBodyEndpointConfigurations {
	s.Endpoint = &v
	return s
}

func (s *DescribeEndpointGroupResponseBodyEndpointConfigurations) SetProbePort(v int32) *DescribeEndpointGroupResponseBodyEndpointConfigurations {
	s.ProbePort = &v
	return s
}

func (s *DescribeEndpointGroupResponseBodyEndpointConfigurations) SetProbeProtocol(v string) *DescribeEndpointGroupResponseBodyEndpointConfigurations {
	s.ProbeProtocol = &v
	return s
}

func (s *DescribeEndpointGroupResponseBodyEndpointConfigurations) SetProvider(v string) *DescribeEndpointGroupResponseBodyEndpointConfigurations {
	s.Provider = &v
	return s
}

func (s *DescribeEndpointGroupResponseBodyEndpointConfigurations) SetSubAddress(v string) *DescribeEndpointGroupResponseBodyEndpointConfigurations {
	s.SubAddress = &v
	return s
}

func (s *DescribeEndpointGroupResponseBodyEndpointConfigurations) SetType(v string) *DescribeEndpointGroupResponseBodyEndpointConfigurations {
	s.Type = &v
	return s
}

func (s *DescribeEndpointGroupResponseBodyEndpointConfigurations) SetVSwitchIds(v []*string) *DescribeEndpointGroupResponseBodyEndpointConfigurations {
	s.VSwitchIds = v
	return s
}

func (s *DescribeEndpointGroupResponseBodyEndpointConfigurations) SetVpcId(v string) *DescribeEndpointGroupResponseBodyEndpointConfigurations {
	s.VpcId = &v
	return s
}

func (s *DescribeEndpointGroupResponseBodyEndpointConfigurations) SetWeight(v int32) *DescribeEndpointGroupResponseBodyEndpointConfigurations {
	s.Weight = &v
	return s
}

func (s *DescribeEndpointGroupResponseBodyEndpointConfigurations) Validate() error {
	return dara.Validate(s)
}

type DescribeEndpointGroupResponseBodyEndpointPrivateIpList struct {
	CIDR *string `json:"CIDR,omitempty" xml:"CIDR,omitempty"`
	// The private IP address.
	//
	// example:
	//
	// 172.16.49.***
	PrivateIp *string `json:"PrivateIp,omitempty" xml:"PrivateIp,omitempty"`
	// The VSwitch ID in the VPC.
	//
	// example:
	//
	// gsw-bp1rsfd2frym**
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s DescribeEndpointGroupResponseBodyEndpointPrivateIpList) String() string {
	return dara.Prettify(s)
}

func (s DescribeEndpointGroupResponseBodyEndpointPrivateIpList) GoString() string {
	return s.String()
}

func (s *DescribeEndpointGroupResponseBodyEndpointPrivateIpList) GetCIDR() *string {
	return s.CIDR
}

func (s *DescribeEndpointGroupResponseBodyEndpointPrivateIpList) GetPrivateIp() *string {
	return s.PrivateIp
}

func (s *DescribeEndpointGroupResponseBodyEndpointPrivateIpList) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeEndpointGroupResponseBodyEndpointPrivateIpList) SetCIDR(v string) *DescribeEndpointGroupResponseBodyEndpointPrivateIpList {
	s.CIDR = &v
	return s
}

func (s *DescribeEndpointGroupResponseBodyEndpointPrivateIpList) SetPrivateIp(v string) *DescribeEndpointGroupResponseBodyEndpointPrivateIpList {
	s.PrivateIp = &v
	return s
}

func (s *DescribeEndpointGroupResponseBodyEndpointPrivateIpList) SetVSwitchId(v string) *DescribeEndpointGroupResponseBodyEndpointPrivateIpList {
	s.VSwitchId = &v
	return s
}

func (s *DescribeEndpointGroupResponseBodyEndpointPrivateIpList) Validate() error {
	return dara.Validate(s)
}

type DescribeEndpointGroupResponseBodyPortOverrides struct {
	// The endpoint port.
	//
	// example:
	//
	// 80
	EndpointPort *int32 `json:"EndpointPort,omitempty" xml:"EndpointPort,omitempty"`
	// The listener port.
	//
	// example:
	//
	// 443
	ListenerPort *int32 `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
}

func (s DescribeEndpointGroupResponseBodyPortOverrides) String() string {
	return dara.Prettify(s)
}

func (s DescribeEndpointGroupResponseBodyPortOverrides) GoString() string {
	return s.String()
}

func (s *DescribeEndpointGroupResponseBodyPortOverrides) GetEndpointPort() *int32 {
	return s.EndpointPort
}

func (s *DescribeEndpointGroupResponseBodyPortOverrides) GetListenerPort() *int32 {
	return s.ListenerPort
}

func (s *DescribeEndpointGroupResponseBodyPortOverrides) SetEndpointPort(v int32) *DescribeEndpointGroupResponseBodyPortOverrides {
	s.EndpointPort = &v
	return s
}

func (s *DescribeEndpointGroupResponseBodyPortOverrides) SetListenerPort(v int32) *DescribeEndpointGroupResponseBodyPortOverrides {
	s.ListenerPort = &v
	return s
}

func (s *DescribeEndpointGroupResponseBodyPortOverrides) Validate() error {
	return dara.Validate(s)
}

type DescribeEndpointGroupResponseBodyServiceManagedInfos struct {
	// The name of the action on the managed instance. Valid values:
	//
	// - **Create**: Create an instance.
	//
	// - **Update**: Update the instance.
	//
	// - **Delete**: Delete the instance.
	//
	// - **Associate**: Associate the instance.
	//
	// - **UserUnmanaged**: Releases the instance from service management.
	//
	// - **CreateChild**: Create a child resource.
	//
	// example:
	//
	// Update
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// The type of the child resource. Valid values:
	//
	// - **Listener**: A listener.
	//
	// - **IpSet**: An acceleration region.
	//
	// - **EndpointGroup**: An endpoint group.
	//
	// - **ForwardingRule**: A forwarding rule.
	//
	// - **Endpoint**: An endpoint.
	//
	// - **EndpointGroupDestination**: A protocol mapping for an endpoint group of a custom routing listener.
	//
	// - **EndpointPolicy**: A traffic policy for an endpoint of a custom routing listener.
	//
	// > This parameter is valid only when **Action*	- is set to **CreateChild**.
	//
	// example:
	//
	// Listener
	ChildType *string `json:"ChildType,omitempty" xml:"ChildType,omitempty"`
	// Indicates whether the action is managed. Valid values:
	//
	// - **true**: The action is managed. You cannot perform this action on the instance.
	//
	// - **false**: The action is not managed. You can perform this action on the instance.
	//
	// example:
	//
	// false
	IsManaged *bool `json:"IsManaged,omitempty" xml:"IsManaged,omitempty"`
}

func (s DescribeEndpointGroupResponseBodyServiceManagedInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeEndpointGroupResponseBodyServiceManagedInfos) GoString() string {
	return s.String()
}

func (s *DescribeEndpointGroupResponseBodyServiceManagedInfos) GetAction() *string {
	return s.Action
}

func (s *DescribeEndpointGroupResponseBodyServiceManagedInfos) GetChildType() *string {
	return s.ChildType
}

func (s *DescribeEndpointGroupResponseBodyServiceManagedInfos) GetIsManaged() *bool {
	return s.IsManaged
}

func (s *DescribeEndpointGroupResponseBodyServiceManagedInfos) SetAction(v string) *DescribeEndpointGroupResponseBodyServiceManagedInfos {
	s.Action = &v
	return s
}

func (s *DescribeEndpointGroupResponseBodyServiceManagedInfos) SetChildType(v string) *DescribeEndpointGroupResponseBodyServiceManagedInfos {
	s.ChildType = &v
	return s
}

func (s *DescribeEndpointGroupResponseBodyServiceManagedInfos) SetIsManaged(v bool) *DescribeEndpointGroupResponseBodyServiceManagedInfos {
	s.IsManaged = &v
	return s
}

func (s *DescribeEndpointGroupResponseBodyServiceManagedInfos) Validate() error {
	return dara.Validate(s)
}

type DescribeEndpointGroupResponseBodyTags struct {
	// The tag key.
	//
	// example:
	//
	// test-key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// test-value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeEndpointGroupResponseBodyTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeEndpointGroupResponseBodyTags) GoString() string {
	return s.String()
}

func (s *DescribeEndpointGroupResponseBodyTags) GetKey() *string {
	return s.Key
}

func (s *DescribeEndpointGroupResponseBodyTags) GetValue() *string {
	return s.Value
}

func (s *DescribeEndpointGroupResponseBodyTags) SetKey(v string) *DescribeEndpointGroupResponseBodyTags {
	s.Key = &v
	return s
}

func (s *DescribeEndpointGroupResponseBodyTags) SetValue(v string) *DescribeEndpointGroupResponseBodyTags {
	s.Value = &v
	return s
}

func (s *DescribeEndpointGroupResponseBodyTags) Validate() error {
	return dara.Validate(s)
}
