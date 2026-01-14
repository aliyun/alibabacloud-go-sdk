// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEndpointGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEndpointGroups(v []*ListEndpointGroupsResponseBodyEndpointGroups) *ListEndpointGroupsResponseBody
	GetEndpointGroups() []*ListEndpointGroupsResponseBodyEndpointGroups
	SetPageNumber(v int32) *ListEndpointGroupsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListEndpointGroupsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListEndpointGroupsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListEndpointGroupsResponseBody
	GetTotalCount() *int32
}

type ListEndpointGroupsResponseBody struct {
	// The configurations of the endpoint groups.
	EndpointGroups []*ListEndpointGroupsResponseBodyEndpointGroups `json:"EndpointGroups,omitempty" xml:"EndpointGroups,omitempty" type:"Repeated"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// A052D49E-CCC2-41DB-816C-DC3381503194
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListEndpointGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListEndpointGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListEndpointGroupsResponseBody) GetEndpointGroups() []*ListEndpointGroupsResponseBodyEndpointGroups {
	return s.EndpointGroups
}

func (s *ListEndpointGroupsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListEndpointGroupsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListEndpointGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListEndpointGroupsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListEndpointGroupsResponseBody) SetEndpointGroups(v []*ListEndpointGroupsResponseBodyEndpointGroups) *ListEndpointGroupsResponseBody {
	s.EndpointGroups = v
	return s
}

func (s *ListEndpointGroupsResponseBody) SetPageNumber(v int32) *ListEndpointGroupsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListEndpointGroupsResponseBody) SetPageSize(v int32) *ListEndpointGroupsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListEndpointGroupsResponseBody) SetRequestId(v string) *ListEndpointGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEndpointGroupsResponseBody) SetTotalCount(v int32) *ListEndpointGroupsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListEndpointGroupsResponseBody) Validate() error {
	if s.EndpointGroups != nil {
		for _, item := range s.EndpointGroups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListEndpointGroupsResponseBodyEndpointGroups struct {
	// The ID of the GA instance.
	//
	// example:
	//
	// ga-bp1odcab8tmno0hdq****
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	// The description of the endpoint group.
	//
	// example:
	//
	// group1
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The configurations of the endpoints in the endpoint group.
	EndpointConfigurations []*ListEndpointGroupsResponseBodyEndpointGroupsEndpointConfigurations `json:"EndpointConfigurations,omitempty" xml:"EndpointConfigurations,omitempty" type:"Repeated"`
	// The ID of the endpoint group.
	//
	// example:
	//
	// epg-bp16jdc00bhe97sr5****
	EndpointGroupId *string `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
	// The endpoint group IP addresses.
	EndpointGroupIpList []*string `json:"EndpointGroupIpList,omitempty" xml:"EndpointGroupIpList,omitempty" type:"Repeated"`
	// The ID of the region where the endpoint group is created.
	//
	// example:
	//
	// cn-hangzhou
	EndpointGroupRegion *string `json:"EndpointGroupRegion,omitempty" xml:"EndpointGroupRegion,omitempty"`
	// The type of the endpoint group. Valid values:
	//
	// 	- **default**: a default endpoint group
	//
	// 	- **virtual:*	- a virtual endpoint group.
	//
	// example:
	//
	// default
	EndpointGroupType *string `json:"EndpointGroupType,omitempty" xml:"EndpointGroupType,omitempty"`
	// The endpoint group IP addresses to be confirmed after the GA instance is upgraded.
	EndpointGroupUnconfirmedIpList []*string                                                            `json:"EndpointGroupUnconfirmedIpList,omitempty" xml:"EndpointGroupUnconfirmedIpList,omitempty" type:"Repeated"`
	EndpointIpVersion              *string                                                              `json:"EndpointIpVersion,omitempty" xml:"EndpointIpVersion,omitempty"`
	EndpointPrivateIpList          []*ListEndpointGroupsResponseBodyEndpointGroupsEndpointPrivateIpList `json:"EndpointPrivateIpList,omitempty" xml:"EndpointPrivateIpList,omitempty" type:"Repeated"`
	// The protocol version that is used by the backend service. Valid values:
	//
	// 	- **HTTP1.1**
	//
	// 	- **HTTP2**
	EndpointProtocolVersion *string `json:"EndpointProtocolVersion,omitempty" xml:"EndpointProtocolVersion,omitempty"`
	// The protocol that is used by the backend server.
	//
	// 	- **HTTP**
	//
	// 	- **HTTPS**
	//
	// example:
	//
	// HTTP
	EndpointRequestProtocol *string `json:"EndpointRequestProtocol,omitempty" xml:"EndpointRequestProtocol,omitempty"`
	// The IDs of the forwarding rules that are associated with the endpoint group.
	ForwardingRuleIds []*string `json:"ForwardingRuleIds,omitempty" xml:"ForwardingRuleIds,omitempty" type:"Repeated"`
	// Indicates whether the health check feature is enabled.
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	HealthCheckEnabled *bool   `json:"HealthCheckEnabled,omitempty" xml:"HealthCheckEnabled,omitempty"`
	HealthCheckHost    *string `json:"HealthCheckHost,omitempty" xml:"HealthCheckHost,omitempty"`
	// The interval at which you want to perform health checks. Unit: seconds.
	//
	// example:
	//
	// 3
	HealthCheckIntervalSeconds *int32 `json:"HealthCheckIntervalSeconds,omitempty" xml:"HealthCheckIntervalSeconds,omitempty"`
	// The path that is used for health checks.
	//
	// example:
	//
	// /healthcheck
	HealthCheckPath *string `json:"HealthCheckPath,omitempty" xml:"HealthCheckPath,omitempty"`
	// The port that is used for health checks.
	//
	// example:
	//
	// 10
	HealthCheckPort *int32 `json:"HealthCheckPort,omitempty" xml:"HealthCheckPort,omitempty"`
	// The protocol over which health check requests are sent. Valid values:
	//
	// 	- **tcp*	- or **TCP**
	//
	// 	- **http*	- or **HTTP**
	//
	// 	- **https*	- or **HTTPS**
	//
	// example:
	//
	// tcp
	HealthCheckProtocol *string `json:"HealthCheckProtocol,omitempty" xml:"HealthCheckProtocol,omitempty"`
	// The ID of the listener.
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
	// The port mapping.
	PortOverrides []*ListEndpointGroupsResponseBodyEndpointGroupsPortOverrides `json:"PortOverrides,omitempty" xml:"PortOverrides,omitempty" type:"Repeated"`
	// The service that manages the instance.
	//
	// >  This parameter takes effect only if the value of **Service managed*	- is **true**.
	//
	// example:
	//
	// ALB
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// Indicates whether the GA instance is managed. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	ServiceManaged *bool `json:"ServiceManaged,omitempty" xml:"ServiceManaged,omitempty"`
	// The actions that users can perform on the managed instance.
	//
	// > -  This parameter takes effect only if the value of **ServiceManaged*	- is **true**.
	//
	// > -   Users can perform only specific actions on a managed instance.
	ServiceManagedInfos []*ListEndpointGroupsResponseBodyEndpointGroupsServiceManagedInfos `json:"ServiceManagedInfos,omitempty" xml:"ServiceManagedInfos,omitempty" type:"Repeated"`
	// The status of the endpoint group. Valid values:
	//
	// 	- **init:*	- The endpoint group is being initialized.
	//
	// 	- **active:*	- The endpoint group is running normally.
	//
	// 	- **updating:**The endpoint group is being updated.
	//
	// 	- **deleteing:*	- The endpoint group is being deleted.
	//
	// example:
	//
	// active
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The tag of the endpoint group.
	Tags []*ListEndpointGroupsResponseBodyEndpointGroupsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The number of consecutive failed health checks that must occur before an endpoint is considered unhealthy.
	//
	// example:
	//
	// 3
	ThresholdCount *int32 `json:"ThresholdCount,omitempty" xml:"ThresholdCount,omitempty"`
	// The value of the traffic distribution ratio. If a listener is associated with multiple endpoint groups, you can set this parameter to distribute different percentages of traffic to the endpoint groups.
	//
	// example:
	//
	// 20
	TrafficPercentage *int32 `json:"TrafficPercentage,omitempty" xml:"TrafficPercentage,omitempty"`
}

func (s ListEndpointGroupsResponseBodyEndpointGroups) String() string {
	return dara.Prettify(s)
}

func (s ListEndpointGroupsResponseBodyEndpointGroups) GoString() string {
	return s.String()
}

func (s *ListEndpointGroupsResponseBodyEndpointGroups) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *ListEndpointGroupsResponseBodyEndpointGroups) GetDescription() *string {
	return s.Description
}

func (s *ListEndpointGroupsResponseBodyEndpointGroups) GetEndpointConfigurations() []*ListEndpointGroupsResponseBodyEndpointGroupsEndpointConfigurations {
	return s.EndpointConfigurations
}

func (s *ListEndpointGroupsResponseBodyEndpointGroups) GetEndpointGroupId() *string {
	return s.EndpointGroupId
}

func (s *ListEndpointGroupsResponseBodyEndpointGroups) GetEndpointGroupIpList() []*string {
	return s.EndpointGroupIpList
}

func (s *ListEndpointGroupsResponseBodyEndpointGroups) GetEndpointGroupRegion() *string {
	return s.EndpointGroupRegion
}

func (s *ListEndpointGroupsResponseBodyEndpointGroups) GetEndpointGroupType() *string {
	return s.EndpointGroupType
}

func (s *ListEndpointGroupsResponseBodyEndpointGroups) GetEndpointGroupUnconfirmedIpList() []*string {
	return s.EndpointGroupUnconfirmedIpList
}

func (s *ListEndpointGroupsResponseBodyEndpointGroups) GetEndpointIpVersion() *string {
	return s.EndpointIpVersion
}

func (s *ListEndpointGroupsResponseBodyEndpointGroups) GetEndpointPrivateIpList() []*ListEndpointGroupsResponseBodyEndpointGroupsEndpointPrivateIpList {
	return s.EndpointPrivateIpList
}

func (s *ListEndpointGroupsResponseBodyEndpointGroups) GetEndpointProtocolVersion() *string {
	return s.EndpointProtocolVersion
}

func (s *ListEndpointGroupsResponseBodyEndpointGroups) GetEndpointRequestProtocol() *string {
	return s.EndpointRequestProtocol
}

func (s *ListEndpointGroupsResponseBodyEndpointGroups) GetForwardingRuleIds() []*string {
	return s.ForwardingRuleIds
}

func (s *ListEndpointGroupsResponseBodyEndpointGroups) GetHealthCheckEnabled() *bool {
	return s.HealthCheckEnabled
}

func (s *ListEndpointGroupsResponseBodyEndpointGroups) GetHealthCheckHost() *string {
	return s.HealthCheckHost
}

func (s *ListEndpointGroupsResponseBodyEndpointGroups) GetHealthCheckIntervalSeconds() *int32 {
	return s.HealthCheckIntervalSeconds
}

func (s *ListEndpointGroupsResponseBodyEndpointGroups) GetHealthCheckPath() *string {
	return s.HealthCheckPath
}

func (s *ListEndpointGroupsResponseBodyEndpointGroups) GetHealthCheckPort() *int32 {
	return s.HealthCheckPort
}

func (s *ListEndpointGroupsResponseBodyEndpointGroups) GetHealthCheckProtocol() *string {
	return s.HealthCheckProtocol
}

func (s *ListEndpointGroupsResponseBodyEndpointGroups) GetListenerId() *string {
	return s.ListenerId
}

func (s *ListEndpointGroupsResponseBodyEndpointGroups) GetName() *string {
	return s.Name
}

func (s *ListEndpointGroupsResponseBodyEndpointGroups) GetPortOverrides() []*ListEndpointGroupsResponseBodyEndpointGroupsPortOverrides {
	return s.PortOverrides
}

func (s *ListEndpointGroupsResponseBodyEndpointGroups) GetServiceId() *string {
	return s.ServiceId
}

func (s *ListEndpointGroupsResponseBodyEndpointGroups) GetServiceManaged() *bool {
	return s.ServiceManaged
}

func (s *ListEndpointGroupsResponseBodyEndpointGroups) GetServiceManagedInfos() []*ListEndpointGroupsResponseBodyEndpointGroupsServiceManagedInfos {
	return s.ServiceManagedInfos
}

func (s *ListEndpointGroupsResponseBodyEndpointGroups) GetState() *string {
	return s.State
}

func (s *ListEndpointGroupsResponseBodyEndpointGroups) GetTags() []*ListEndpointGroupsResponseBodyEndpointGroupsTags {
	return s.Tags
}

func (s *ListEndpointGroupsResponseBodyEndpointGroups) GetThresholdCount() *int32 {
	return s.ThresholdCount
}

func (s *ListEndpointGroupsResponseBodyEndpointGroups) GetTrafficPercentage() *int32 {
	return s.TrafficPercentage
}

func (s *ListEndpointGroupsResponseBodyEndpointGroups) SetAcceleratorId(v string) *ListEndpointGroupsResponseBodyEndpointGroups {
	s.AcceleratorId = &v
	return s
}

func (s *ListEndpointGroupsResponseBodyEndpointGroups) SetDescription(v string) *ListEndpointGroupsResponseBodyEndpointGroups {
	s.Description = &v
	return s
}

func (s *ListEndpointGroupsResponseBodyEndpointGroups) SetEndpointConfigurations(v []*ListEndpointGroupsResponseBodyEndpointGroupsEndpointConfigurations) *ListEndpointGroupsResponseBodyEndpointGroups {
	s.EndpointConfigurations = v
	return s
}

func (s *ListEndpointGroupsResponseBodyEndpointGroups) SetEndpointGroupId(v string) *ListEndpointGroupsResponseBodyEndpointGroups {
	s.EndpointGroupId = &v
	return s
}

func (s *ListEndpointGroupsResponseBodyEndpointGroups) SetEndpointGroupIpList(v []*string) *ListEndpointGroupsResponseBodyEndpointGroups {
	s.EndpointGroupIpList = v
	return s
}

func (s *ListEndpointGroupsResponseBodyEndpointGroups) SetEndpointGroupRegion(v string) *ListEndpointGroupsResponseBodyEndpointGroups {
	s.EndpointGroupRegion = &v
	return s
}

func (s *ListEndpointGroupsResponseBodyEndpointGroups) SetEndpointGroupType(v string) *ListEndpointGroupsResponseBodyEndpointGroups {
	s.EndpointGroupType = &v
	return s
}

func (s *ListEndpointGroupsResponseBodyEndpointGroups) SetEndpointGroupUnconfirmedIpList(v []*string) *ListEndpointGroupsResponseBodyEndpointGroups {
	s.EndpointGroupUnconfirmedIpList = v
	return s
}

func (s *ListEndpointGroupsResponseBodyEndpointGroups) SetEndpointIpVersion(v string) *ListEndpointGroupsResponseBodyEndpointGroups {
	s.EndpointIpVersion = &v
	return s
}

func (s *ListEndpointGroupsResponseBodyEndpointGroups) SetEndpointPrivateIpList(v []*ListEndpointGroupsResponseBodyEndpointGroupsEndpointPrivateIpList) *ListEndpointGroupsResponseBodyEndpointGroups {
	s.EndpointPrivateIpList = v
	return s
}

func (s *ListEndpointGroupsResponseBodyEndpointGroups) SetEndpointProtocolVersion(v string) *ListEndpointGroupsResponseBodyEndpointGroups {
	s.EndpointProtocolVersion = &v
	return s
}

func (s *ListEndpointGroupsResponseBodyEndpointGroups) SetEndpointRequestProtocol(v string) *ListEndpointGroupsResponseBodyEndpointGroups {
	s.EndpointRequestProtocol = &v
	return s
}

func (s *ListEndpointGroupsResponseBodyEndpointGroups) SetForwardingRuleIds(v []*string) *ListEndpointGroupsResponseBodyEndpointGroups {
	s.ForwardingRuleIds = v
	return s
}

func (s *ListEndpointGroupsResponseBodyEndpointGroups) SetHealthCheckEnabled(v bool) *ListEndpointGroupsResponseBodyEndpointGroups {
	s.HealthCheckEnabled = &v
	return s
}

func (s *ListEndpointGroupsResponseBodyEndpointGroups) SetHealthCheckHost(v string) *ListEndpointGroupsResponseBodyEndpointGroups {
	s.HealthCheckHost = &v
	return s
}

func (s *ListEndpointGroupsResponseBodyEndpointGroups) SetHealthCheckIntervalSeconds(v int32) *ListEndpointGroupsResponseBodyEndpointGroups {
	s.HealthCheckIntervalSeconds = &v
	return s
}

func (s *ListEndpointGroupsResponseBodyEndpointGroups) SetHealthCheckPath(v string) *ListEndpointGroupsResponseBodyEndpointGroups {
	s.HealthCheckPath = &v
	return s
}

func (s *ListEndpointGroupsResponseBodyEndpointGroups) SetHealthCheckPort(v int32) *ListEndpointGroupsResponseBodyEndpointGroups {
	s.HealthCheckPort = &v
	return s
}

func (s *ListEndpointGroupsResponseBodyEndpointGroups) SetHealthCheckProtocol(v string) *ListEndpointGroupsResponseBodyEndpointGroups {
	s.HealthCheckProtocol = &v
	return s
}

func (s *ListEndpointGroupsResponseBodyEndpointGroups) SetListenerId(v string) *ListEndpointGroupsResponseBodyEndpointGroups {
	s.ListenerId = &v
	return s
}

func (s *ListEndpointGroupsResponseBodyEndpointGroups) SetName(v string) *ListEndpointGroupsResponseBodyEndpointGroups {
	s.Name = &v
	return s
}

func (s *ListEndpointGroupsResponseBodyEndpointGroups) SetPortOverrides(v []*ListEndpointGroupsResponseBodyEndpointGroupsPortOverrides) *ListEndpointGroupsResponseBodyEndpointGroups {
	s.PortOverrides = v
	return s
}

func (s *ListEndpointGroupsResponseBodyEndpointGroups) SetServiceId(v string) *ListEndpointGroupsResponseBodyEndpointGroups {
	s.ServiceId = &v
	return s
}

func (s *ListEndpointGroupsResponseBodyEndpointGroups) SetServiceManaged(v bool) *ListEndpointGroupsResponseBodyEndpointGroups {
	s.ServiceManaged = &v
	return s
}

func (s *ListEndpointGroupsResponseBodyEndpointGroups) SetServiceManagedInfos(v []*ListEndpointGroupsResponseBodyEndpointGroupsServiceManagedInfos) *ListEndpointGroupsResponseBodyEndpointGroups {
	s.ServiceManagedInfos = v
	return s
}

func (s *ListEndpointGroupsResponseBodyEndpointGroups) SetState(v string) *ListEndpointGroupsResponseBodyEndpointGroups {
	s.State = &v
	return s
}

func (s *ListEndpointGroupsResponseBodyEndpointGroups) SetTags(v []*ListEndpointGroupsResponseBodyEndpointGroupsTags) *ListEndpointGroupsResponseBodyEndpointGroups {
	s.Tags = v
	return s
}

func (s *ListEndpointGroupsResponseBodyEndpointGroups) SetThresholdCount(v int32) *ListEndpointGroupsResponseBodyEndpointGroups {
	s.ThresholdCount = &v
	return s
}

func (s *ListEndpointGroupsResponseBodyEndpointGroups) SetTrafficPercentage(v int32) *ListEndpointGroupsResponseBodyEndpointGroups {
	s.TrafficPercentage = &v
	return s
}

func (s *ListEndpointGroupsResponseBodyEndpointGroups) Validate() error {
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

type ListEndpointGroupsResponseBodyEndpointGroupsEndpointConfigurations struct {
	// Indicates whether the client IP address preservation feature is enabled. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	EnableClientIPPreservation *bool `json:"EnableClientIPPreservation,omitempty" xml:"EnableClientIPPreservation,omitempty"`
	// Indicates whether the proxy protocol is used to preserve client IP addresses. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	EnableProxyProtocol *bool `json:"EnableProxyProtocol,omitempty" xml:"EnableProxyProtocol,omitempty"`
	// The IP address or domain name of the endpoint.
	//
	// example:
	//
	// 47.1.XX.XX
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// The ID of the endpoint.
	//
	// example:
	//
	// ep-bp1d2utp8qqe2a44t****
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	// The port that is used to monitor latency.
	//
	// example:
	//
	// 80
	ProbePort *int32 `json:"ProbePort,omitempty" xml:"ProbePort,omitempty"`
	// The protocol that is used to monitor latency.
	//
	// 	- **icmp**
	//
	// 	- **tcp**
	//
	// example:
	//
	// tcp
	ProbeProtocol *string `json:"ProbeProtocol,omitempty" xml:"ProbeProtocol,omitempty"`
	// The private IP address of the ENI.
	//
	// example:
	//
	// 172.168.XX.XX
	SubAddress *string `json:"SubAddress,omitempty" xml:"SubAddress,omitempty"`
	// The type of the endpoint. Valid values:
	//
	// 	- **Domain**: a custom domain name.
	//
	// 	- **Ip**: a custom IP address.
	//
	// 	- **IpTarget**: a custom private IP address.
	//
	// 	- **PublicIp**: a public IP address provided by Alibaba Cloud.
	//
	// 	- **ECS**: an Elastic Compute Service (ECS) instance.
	//
	// 	- **SLB**: a Server Load Balancer (SLB) instance.
	//
	// 	- **ALB**: an Application Load Balancer (ALB) instance.
	//
	// 	- **OSS**: an Object Storage Service (OSS) bucket.
	//
	// 	- **ENI**: an elastic network interface (ENI).
	//
	// 	- **NLB**: a Network Load Balancer (NLB) instance.
	//
	// example:
	//
	// Ip
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The IDs of vSwitches that are deployed in the VPC.
	VSwitchIds []*string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
	// The VPC ID.
	//
	// example:
	//
	// vpc-8vbhucmd5b2q2fpqqu****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The weight of the endpoint.
	//
	// example:
	//
	// 20
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s ListEndpointGroupsResponseBodyEndpointGroupsEndpointConfigurations) String() string {
	return dara.Prettify(s)
}

func (s ListEndpointGroupsResponseBodyEndpointGroupsEndpointConfigurations) GoString() string {
	return s.String()
}

func (s *ListEndpointGroupsResponseBodyEndpointGroupsEndpointConfigurations) GetEnableClientIPPreservation() *bool {
	return s.EnableClientIPPreservation
}

func (s *ListEndpointGroupsResponseBodyEndpointGroupsEndpointConfigurations) GetEnableProxyProtocol() *bool {
	return s.EnableProxyProtocol
}

func (s *ListEndpointGroupsResponseBodyEndpointGroupsEndpointConfigurations) GetEndpoint() *string {
	return s.Endpoint
}

func (s *ListEndpointGroupsResponseBodyEndpointGroupsEndpointConfigurations) GetEndpointId() *string {
	return s.EndpointId
}

func (s *ListEndpointGroupsResponseBodyEndpointGroupsEndpointConfigurations) GetProbePort() *int32 {
	return s.ProbePort
}

func (s *ListEndpointGroupsResponseBodyEndpointGroupsEndpointConfigurations) GetProbeProtocol() *string {
	return s.ProbeProtocol
}

func (s *ListEndpointGroupsResponseBodyEndpointGroupsEndpointConfigurations) GetSubAddress() *string {
	return s.SubAddress
}

func (s *ListEndpointGroupsResponseBodyEndpointGroupsEndpointConfigurations) GetType() *string {
	return s.Type
}

func (s *ListEndpointGroupsResponseBodyEndpointGroupsEndpointConfigurations) GetVSwitchIds() []*string {
	return s.VSwitchIds
}

func (s *ListEndpointGroupsResponseBodyEndpointGroupsEndpointConfigurations) GetVpcId() *string {
	return s.VpcId
}

func (s *ListEndpointGroupsResponseBodyEndpointGroupsEndpointConfigurations) GetWeight() *int32 {
	return s.Weight
}

func (s *ListEndpointGroupsResponseBodyEndpointGroupsEndpointConfigurations) SetEnableClientIPPreservation(v bool) *ListEndpointGroupsResponseBodyEndpointGroupsEndpointConfigurations {
	s.EnableClientIPPreservation = &v
	return s
}

func (s *ListEndpointGroupsResponseBodyEndpointGroupsEndpointConfigurations) SetEnableProxyProtocol(v bool) *ListEndpointGroupsResponseBodyEndpointGroupsEndpointConfigurations {
	s.EnableProxyProtocol = &v
	return s
}

func (s *ListEndpointGroupsResponseBodyEndpointGroupsEndpointConfigurations) SetEndpoint(v string) *ListEndpointGroupsResponseBodyEndpointGroupsEndpointConfigurations {
	s.Endpoint = &v
	return s
}

func (s *ListEndpointGroupsResponseBodyEndpointGroupsEndpointConfigurations) SetEndpointId(v string) *ListEndpointGroupsResponseBodyEndpointGroupsEndpointConfigurations {
	s.EndpointId = &v
	return s
}

func (s *ListEndpointGroupsResponseBodyEndpointGroupsEndpointConfigurations) SetProbePort(v int32) *ListEndpointGroupsResponseBodyEndpointGroupsEndpointConfigurations {
	s.ProbePort = &v
	return s
}

func (s *ListEndpointGroupsResponseBodyEndpointGroupsEndpointConfigurations) SetProbeProtocol(v string) *ListEndpointGroupsResponseBodyEndpointGroupsEndpointConfigurations {
	s.ProbeProtocol = &v
	return s
}

func (s *ListEndpointGroupsResponseBodyEndpointGroupsEndpointConfigurations) SetSubAddress(v string) *ListEndpointGroupsResponseBodyEndpointGroupsEndpointConfigurations {
	s.SubAddress = &v
	return s
}

func (s *ListEndpointGroupsResponseBodyEndpointGroupsEndpointConfigurations) SetType(v string) *ListEndpointGroupsResponseBodyEndpointGroupsEndpointConfigurations {
	s.Type = &v
	return s
}

func (s *ListEndpointGroupsResponseBodyEndpointGroupsEndpointConfigurations) SetVSwitchIds(v []*string) *ListEndpointGroupsResponseBodyEndpointGroupsEndpointConfigurations {
	s.VSwitchIds = v
	return s
}

func (s *ListEndpointGroupsResponseBodyEndpointGroupsEndpointConfigurations) SetVpcId(v string) *ListEndpointGroupsResponseBodyEndpointGroupsEndpointConfigurations {
	s.VpcId = &v
	return s
}

func (s *ListEndpointGroupsResponseBodyEndpointGroupsEndpointConfigurations) SetWeight(v int32) *ListEndpointGroupsResponseBodyEndpointGroupsEndpointConfigurations {
	s.Weight = &v
	return s
}

func (s *ListEndpointGroupsResponseBodyEndpointGroupsEndpointConfigurations) Validate() error {
	return dara.Validate(s)
}

type ListEndpointGroupsResponseBodyEndpointGroupsEndpointPrivateIpList struct {
	CIDR      *string `json:"CIDR,omitempty" xml:"CIDR,omitempty"`
	PrivateIp *string `json:"PrivateIp,omitempty" xml:"PrivateIp,omitempty"`
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s ListEndpointGroupsResponseBodyEndpointGroupsEndpointPrivateIpList) String() string {
	return dara.Prettify(s)
}

func (s ListEndpointGroupsResponseBodyEndpointGroupsEndpointPrivateIpList) GoString() string {
	return s.String()
}

func (s *ListEndpointGroupsResponseBodyEndpointGroupsEndpointPrivateIpList) GetCIDR() *string {
	return s.CIDR
}

func (s *ListEndpointGroupsResponseBodyEndpointGroupsEndpointPrivateIpList) GetPrivateIp() *string {
	return s.PrivateIp
}

func (s *ListEndpointGroupsResponseBodyEndpointGroupsEndpointPrivateIpList) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *ListEndpointGroupsResponseBodyEndpointGroupsEndpointPrivateIpList) SetCIDR(v string) *ListEndpointGroupsResponseBodyEndpointGroupsEndpointPrivateIpList {
	s.CIDR = &v
	return s
}

func (s *ListEndpointGroupsResponseBodyEndpointGroupsEndpointPrivateIpList) SetPrivateIp(v string) *ListEndpointGroupsResponseBodyEndpointGroupsEndpointPrivateIpList {
	s.PrivateIp = &v
	return s
}

func (s *ListEndpointGroupsResponseBodyEndpointGroupsEndpointPrivateIpList) SetVSwitchId(v string) *ListEndpointGroupsResponseBodyEndpointGroupsEndpointPrivateIpList {
	s.VSwitchId = &v
	return s
}

func (s *ListEndpointGroupsResponseBodyEndpointGroupsEndpointPrivateIpList) Validate() error {
	return dara.Validate(s)
}

type ListEndpointGroupsResponseBodyEndpointGroupsPortOverrides struct {
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

func (s ListEndpointGroupsResponseBodyEndpointGroupsPortOverrides) String() string {
	return dara.Prettify(s)
}

func (s ListEndpointGroupsResponseBodyEndpointGroupsPortOverrides) GoString() string {
	return s.String()
}

func (s *ListEndpointGroupsResponseBodyEndpointGroupsPortOverrides) GetEndpointPort() *int32 {
	return s.EndpointPort
}

func (s *ListEndpointGroupsResponseBodyEndpointGroupsPortOverrides) GetListenerPort() *int32 {
	return s.ListenerPort
}

func (s *ListEndpointGroupsResponseBodyEndpointGroupsPortOverrides) SetEndpointPort(v int32) *ListEndpointGroupsResponseBodyEndpointGroupsPortOverrides {
	s.EndpointPort = &v
	return s
}

func (s *ListEndpointGroupsResponseBodyEndpointGroupsPortOverrides) SetListenerPort(v int32) *ListEndpointGroupsResponseBodyEndpointGroupsPortOverrides {
	s.ListenerPort = &v
	return s
}

func (s *ListEndpointGroupsResponseBodyEndpointGroupsPortOverrides) Validate() error {
	return dara.Validate(s)
}

type ListEndpointGroupsResponseBodyEndpointGroupsServiceManagedInfos struct {
	// The name of the action that was performed on the managed instance. Valid values:
	//
	// 	- **Create:*	- Create an instance.
	//
	// 	- **Update:*	- Update the current instance.
	//
	// 	- **Delete:*	- Delete the current instance.
	//
	// 	- **Associate:*	- Reference the current instance.
	//
	// 	- **UserUnmanaged:*	- Unmanage the instance.
	//
	// 	- **CreateChild:*	- Create a child resource in the current instance.
	//
	// example:
	//
	// Update
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// The type of the child resource. Valid values:
	//
	// 	- **Listener:*	- listener.
	//
	// 	- **IpSet:*	- acceleration region.
	//
	// 	- **EndpointGroup:*	- endpoint group.
	//
	// 	- **ForwardingRule:*	- forwarding rule.
	//
	// 	- **Endpoint:*	- endpoint.
	//
	// 	- **EndpointGroupDestination:*	- the protocol mapping of an endpoint group associated with a custom routing listener.
	//
	// 	- **EndpointPolicy:*	- the traffic policy of an endpoint associated with a custom routing listener.
	//
	// >  This parameter takes effect only if you set **Action*	- to **CreateChild**.
	//
	// example:
	//
	// Listener
	ChildType *string `json:"ChildType,omitempty" xml:"ChildType,omitempty"`
	// Indicates whether the specified actions are managed. Valid values:
	//
	// 	- **true**: The specified actions are managed, and users cannot perform the specified actions on the managed instance.
	//
	// 	- **false**: The specified actions are not managed, and users can perform the specified actions on the managed instance.
	//
	// example:
	//
	// false
	IsManaged *bool `json:"IsManaged,omitempty" xml:"IsManaged,omitempty"`
}

func (s ListEndpointGroupsResponseBodyEndpointGroupsServiceManagedInfos) String() string {
	return dara.Prettify(s)
}

func (s ListEndpointGroupsResponseBodyEndpointGroupsServiceManagedInfos) GoString() string {
	return s.String()
}

func (s *ListEndpointGroupsResponseBodyEndpointGroupsServiceManagedInfos) GetAction() *string {
	return s.Action
}

func (s *ListEndpointGroupsResponseBodyEndpointGroupsServiceManagedInfos) GetChildType() *string {
	return s.ChildType
}

func (s *ListEndpointGroupsResponseBodyEndpointGroupsServiceManagedInfos) GetIsManaged() *bool {
	return s.IsManaged
}

func (s *ListEndpointGroupsResponseBodyEndpointGroupsServiceManagedInfos) SetAction(v string) *ListEndpointGroupsResponseBodyEndpointGroupsServiceManagedInfos {
	s.Action = &v
	return s
}

func (s *ListEndpointGroupsResponseBodyEndpointGroupsServiceManagedInfos) SetChildType(v string) *ListEndpointGroupsResponseBodyEndpointGroupsServiceManagedInfos {
	s.ChildType = &v
	return s
}

func (s *ListEndpointGroupsResponseBodyEndpointGroupsServiceManagedInfos) SetIsManaged(v bool) *ListEndpointGroupsResponseBodyEndpointGroupsServiceManagedInfos {
	s.IsManaged = &v
	return s
}

func (s *ListEndpointGroupsResponseBodyEndpointGroupsServiceManagedInfos) Validate() error {
	return dara.Validate(s)
}

type ListEndpointGroupsResponseBodyEndpointGroupsTags struct {
	// The tag key of the endpoint group.
	//
	// example:
	//
	// test-key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the endpoint group.
	//
	// example:
	//
	// test-value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListEndpointGroupsResponseBodyEndpointGroupsTags) String() string {
	return dara.Prettify(s)
}

func (s ListEndpointGroupsResponseBodyEndpointGroupsTags) GoString() string {
	return s.String()
}

func (s *ListEndpointGroupsResponseBodyEndpointGroupsTags) GetKey() *string {
	return s.Key
}

func (s *ListEndpointGroupsResponseBodyEndpointGroupsTags) GetValue() *string {
	return s.Value
}

func (s *ListEndpointGroupsResponseBodyEndpointGroupsTags) SetKey(v string) *ListEndpointGroupsResponseBodyEndpointGroupsTags {
	s.Key = &v
	return s
}

func (s *ListEndpointGroupsResponseBodyEndpointGroupsTags) SetValue(v string) *ListEndpointGroupsResponseBodyEndpointGroupsTags {
	s.Value = &v
	return s
}

func (s *ListEndpointGroupsResponseBodyEndpointGroupsTags) Validate() error {
	return dara.Validate(s)
}
