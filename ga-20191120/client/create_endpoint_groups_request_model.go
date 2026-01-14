// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEndpointGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceleratorId(v string) *CreateEndpointGroupsRequest
	GetAcceleratorId() *string
	SetClientToken(v string) *CreateEndpointGroupsRequest
	GetClientToken() *string
	SetDryRun(v bool) *CreateEndpointGroupsRequest
	GetDryRun() *bool
	SetEndpointGroupConfigurations(v []*CreateEndpointGroupsRequestEndpointGroupConfigurations) *CreateEndpointGroupsRequest
	GetEndpointGroupConfigurations() []*CreateEndpointGroupsRequestEndpointGroupConfigurations
	SetListenerId(v string) *CreateEndpointGroupsRequest
	GetListenerId() *string
	SetRegionId(v string) *CreateEndpointGroupsRequest
	GetRegionId() *string
}

type CreateEndpointGroupsRequest struct {
	// The ID of the GA instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ga-bp1odcab8tmno0hdq****
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 1F4B6A4A-C89E-489E-BAF1-52777EE148EF
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error code is returned. If the request passes the dry run, a 2xx HTTP status code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// This parameter is required.
	EndpointGroupConfigurations []*CreateEndpointGroupsRequestEndpointGroupConfigurations `json:"EndpointGroupConfigurations,omitempty" xml:"EndpointGroupConfigurations,omitempty" type:"Repeated"`
	// The ID of the listener.
	//
	// >  If the listener protocol is **HTTP*	- or **HTTPS**, you can call the **CreateEndpointGroups*	- operation to create only one endpoint group.
	//
	// This parameter is required.
	//
	// example:
	//
	// lsr-bp1bpn0kn908w4nbw****
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// The ID of the region where the GA instance is deployed. Set the value to **cn-hangzhou**.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateEndpointGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateEndpointGroupsRequest) GoString() string {
	return s.String()
}

func (s *CreateEndpointGroupsRequest) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *CreateEndpointGroupsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateEndpointGroupsRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateEndpointGroupsRequest) GetEndpointGroupConfigurations() []*CreateEndpointGroupsRequestEndpointGroupConfigurations {
	return s.EndpointGroupConfigurations
}

func (s *CreateEndpointGroupsRequest) GetListenerId() *string {
	return s.ListenerId
}

func (s *CreateEndpointGroupsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateEndpointGroupsRequest) SetAcceleratorId(v string) *CreateEndpointGroupsRequest {
	s.AcceleratorId = &v
	return s
}

func (s *CreateEndpointGroupsRequest) SetClientToken(v string) *CreateEndpointGroupsRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateEndpointGroupsRequest) SetDryRun(v bool) *CreateEndpointGroupsRequest {
	s.DryRun = &v
	return s
}

func (s *CreateEndpointGroupsRequest) SetEndpointGroupConfigurations(v []*CreateEndpointGroupsRequestEndpointGroupConfigurations) *CreateEndpointGroupsRequest {
	s.EndpointGroupConfigurations = v
	return s
}

func (s *CreateEndpointGroupsRequest) SetListenerId(v string) *CreateEndpointGroupsRequest {
	s.ListenerId = &v
	return s
}

func (s *CreateEndpointGroupsRequest) SetRegionId(v string) *CreateEndpointGroupsRequest {
	s.RegionId = &v
	return s
}

func (s *CreateEndpointGroupsRequest) Validate() error {
	if s.EndpointGroupConfigurations != nil {
		for _, item := range s.EndpointGroupConfigurations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateEndpointGroupsRequestEndpointGroupConfigurations struct {
	EndpointConfigurations   []*CreateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations `json:"EndpointConfigurations,omitempty" xml:"EndpointConfigurations,omitempty" type:"Repeated"`
	EndpointGroupDescription *string                                                                         `json:"EndpointGroupDescription,omitempty" xml:"EndpointGroupDescription,omitempty"`
	EndpointGroupName        *string                                                                         `json:"EndpointGroupName,omitempty" xml:"EndpointGroupName,omitempty"`
	// This parameter is required.
	EndpointGroupRegion        *string                                                                `json:"EndpointGroupRegion,omitempty" xml:"EndpointGroupRegion,omitempty"`
	EndpointGroupType          *string                                                                `json:"EndpointGroupType,omitempty" xml:"EndpointGroupType,omitempty"`
	EndpointIpVersion          *string                                                                `json:"EndpointIpVersion,omitempty" xml:"EndpointIpVersion,omitempty"`
	EndpointProtocolVersion    *string                                                                `json:"EndpointProtocolVersion,omitempty" xml:"EndpointProtocolVersion,omitempty"`
	EndpointRequestProtocol    *string                                                                `json:"EndpointRequestProtocol,omitempty" xml:"EndpointRequestProtocol,omitempty"`
	HealthCheckEnabled         *bool                                                                  `json:"HealthCheckEnabled,omitempty" xml:"HealthCheckEnabled,omitempty"`
	HealthCheckHost            *string                                                                `json:"HealthCheckHost,omitempty" xml:"HealthCheckHost,omitempty"`
	HealthCheckIntervalSeconds *int64                                                                 `json:"HealthCheckIntervalSeconds,omitempty" xml:"HealthCheckIntervalSeconds,omitempty"`
	HealthCheckPath            *string                                                                `json:"HealthCheckPath,omitempty" xml:"HealthCheckPath,omitempty"`
	HealthCheckPort            *int64                                                                 `json:"HealthCheckPort,omitempty" xml:"HealthCheckPort,omitempty"`
	HealthCheckProtocol        *string                                                                `json:"HealthCheckProtocol,omitempty" xml:"HealthCheckProtocol,omitempty"`
	PortOverrides              []*CreateEndpointGroupsRequestEndpointGroupConfigurationsPortOverrides `json:"PortOverrides,omitempty" xml:"PortOverrides,omitempty" type:"Repeated"`
	SystemTag                  []*CreateEndpointGroupsRequestEndpointGroupConfigurationsSystemTag     `json:"SystemTag,omitempty" xml:"SystemTag,omitempty" type:"Repeated"`
	Tag                        []*CreateEndpointGroupsRequestEndpointGroupConfigurationsTag           `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	ThresholdCount             *int64                                                                 `json:"ThresholdCount,omitempty" xml:"ThresholdCount,omitempty"`
	TrafficPercentage          *int64                                                                 `json:"TrafficPercentage,omitempty" xml:"TrafficPercentage,omitempty"`
}

func (s CreateEndpointGroupsRequestEndpointGroupConfigurations) String() string {
	return dara.Prettify(s)
}

func (s CreateEndpointGroupsRequestEndpointGroupConfigurations) GoString() string {
	return s.String()
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurations) GetEndpointConfigurations() []*CreateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations {
	return s.EndpointConfigurations
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurations) GetEndpointGroupDescription() *string {
	return s.EndpointGroupDescription
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurations) GetEndpointGroupName() *string {
	return s.EndpointGroupName
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurations) GetEndpointGroupRegion() *string {
	return s.EndpointGroupRegion
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurations) GetEndpointGroupType() *string {
	return s.EndpointGroupType
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurations) GetEndpointIpVersion() *string {
	return s.EndpointIpVersion
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurations) GetEndpointProtocolVersion() *string {
	return s.EndpointProtocolVersion
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurations) GetEndpointRequestProtocol() *string {
	return s.EndpointRequestProtocol
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurations) GetHealthCheckEnabled() *bool {
	return s.HealthCheckEnabled
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurations) GetHealthCheckHost() *string {
	return s.HealthCheckHost
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurations) GetHealthCheckIntervalSeconds() *int64 {
	return s.HealthCheckIntervalSeconds
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurations) GetHealthCheckPath() *string {
	return s.HealthCheckPath
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurations) GetHealthCheckPort() *int64 {
	return s.HealthCheckPort
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurations) GetHealthCheckProtocol() *string {
	return s.HealthCheckProtocol
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurations) GetPortOverrides() []*CreateEndpointGroupsRequestEndpointGroupConfigurationsPortOverrides {
	return s.PortOverrides
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurations) GetSystemTag() []*CreateEndpointGroupsRequestEndpointGroupConfigurationsSystemTag {
	return s.SystemTag
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurations) GetTag() []*CreateEndpointGroupsRequestEndpointGroupConfigurationsTag {
	return s.Tag
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurations) GetThresholdCount() *int64 {
	return s.ThresholdCount
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurations) GetTrafficPercentage() *int64 {
	return s.TrafficPercentage
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurations) SetEndpointConfigurations(v []*CreateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations) *CreateEndpointGroupsRequestEndpointGroupConfigurations {
	s.EndpointConfigurations = v
	return s
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurations) SetEndpointGroupDescription(v string) *CreateEndpointGroupsRequestEndpointGroupConfigurations {
	s.EndpointGroupDescription = &v
	return s
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurations) SetEndpointGroupName(v string) *CreateEndpointGroupsRequestEndpointGroupConfigurations {
	s.EndpointGroupName = &v
	return s
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurations) SetEndpointGroupRegion(v string) *CreateEndpointGroupsRequestEndpointGroupConfigurations {
	s.EndpointGroupRegion = &v
	return s
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurations) SetEndpointGroupType(v string) *CreateEndpointGroupsRequestEndpointGroupConfigurations {
	s.EndpointGroupType = &v
	return s
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurations) SetEndpointIpVersion(v string) *CreateEndpointGroupsRequestEndpointGroupConfigurations {
	s.EndpointIpVersion = &v
	return s
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurations) SetEndpointProtocolVersion(v string) *CreateEndpointGroupsRequestEndpointGroupConfigurations {
	s.EndpointProtocolVersion = &v
	return s
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurations) SetEndpointRequestProtocol(v string) *CreateEndpointGroupsRequestEndpointGroupConfigurations {
	s.EndpointRequestProtocol = &v
	return s
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurations) SetHealthCheckEnabled(v bool) *CreateEndpointGroupsRequestEndpointGroupConfigurations {
	s.HealthCheckEnabled = &v
	return s
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurations) SetHealthCheckHost(v string) *CreateEndpointGroupsRequestEndpointGroupConfigurations {
	s.HealthCheckHost = &v
	return s
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurations) SetHealthCheckIntervalSeconds(v int64) *CreateEndpointGroupsRequestEndpointGroupConfigurations {
	s.HealthCheckIntervalSeconds = &v
	return s
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurations) SetHealthCheckPath(v string) *CreateEndpointGroupsRequestEndpointGroupConfigurations {
	s.HealthCheckPath = &v
	return s
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurations) SetHealthCheckPort(v int64) *CreateEndpointGroupsRequestEndpointGroupConfigurations {
	s.HealthCheckPort = &v
	return s
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurations) SetHealthCheckProtocol(v string) *CreateEndpointGroupsRequestEndpointGroupConfigurations {
	s.HealthCheckProtocol = &v
	return s
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurations) SetPortOverrides(v []*CreateEndpointGroupsRequestEndpointGroupConfigurationsPortOverrides) *CreateEndpointGroupsRequestEndpointGroupConfigurations {
	s.PortOverrides = v
	return s
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurations) SetSystemTag(v []*CreateEndpointGroupsRequestEndpointGroupConfigurationsSystemTag) *CreateEndpointGroupsRequestEndpointGroupConfigurations {
	s.SystemTag = v
	return s
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurations) SetTag(v []*CreateEndpointGroupsRequestEndpointGroupConfigurationsTag) *CreateEndpointGroupsRequestEndpointGroupConfigurations {
	s.Tag = v
	return s
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurations) SetThresholdCount(v int64) *CreateEndpointGroupsRequestEndpointGroupConfigurations {
	s.ThresholdCount = &v
	return s
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurations) SetTrafficPercentage(v int64) *CreateEndpointGroupsRequestEndpointGroupConfigurations {
	s.TrafficPercentage = &v
	return s
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurations) Validate() error {
	if s.EndpointConfigurations != nil {
		for _, item := range s.EndpointConfigurations {
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
	if s.SystemTag != nil {
		for _, item := range s.SystemTag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations struct {
	EnableClientIPPreservation *bool     `json:"EnableClientIPPreservation,omitempty" xml:"EnableClientIPPreservation,omitempty"`
	EnableProxyProtocol        *bool     `json:"EnableProxyProtocol,omitempty" xml:"EnableProxyProtocol,omitempty"`
	Endpoint                   *string   `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	SubAddress                 *string   `json:"SubAddress,omitempty" xml:"SubAddress,omitempty"`
	Type                       *string   `json:"Type,omitempty" xml:"Type,omitempty"`
	VSwitchIds                 []*string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
	VpcId                      *string   `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	Weight                     *int64    `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s CreateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations) String() string {
	return dara.Prettify(s)
}

func (s CreateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations) GoString() string {
	return s.String()
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations) GetEnableClientIPPreservation() *bool {
	return s.EnableClientIPPreservation
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations) GetEnableProxyProtocol() *bool {
	return s.EnableProxyProtocol
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations) GetEndpoint() *string {
	return s.Endpoint
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations) GetSubAddress() *string {
	return s.SubAddress
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations) GetType() *string {
	return s.Type
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations) GetVSwitchIds() []*string {
	return s.VSwitchIds
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations) GetWeight() *int64 {
	return s.Weight
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations) SetEnableClientIPPreservation(v bool) *CreateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations {
	s.EnableClientIPPreservation = &v
	return s
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations) SetEnableProxyProtocol(v bool) *CreateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations {
	s.EnableProxyProtocol = &v
	return s
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations) SetEndpoint(v string) *CreateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations {
	s.Endpoint = &v
	return s
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations) SetSubAddress(v string) *CreateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations {
	s.SubAddress = &v
	return s
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations) SetType(v string) *CreateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations {
	s.Type = &v
	return s
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations) SetVSwitchIds(v []*string) *CreateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations {
	s.VSwitchIds = v
	return s
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations) SetVpcId(v string) *CreateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations {
	s.VpcId = &v
	return s
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations) SetWeight(v int64) *CreateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations {
	s.Weight = &v
	return s
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurationsEndpointConfigurations) Validate() error {
	return dara.Validate(s)
}

type CreateEndpointGroupsRequestEndpointGroupConfigurationsPortOverrides struct {
	EndpointPort *int64 `json:"EndpointPort,omitempty" xml:"EndpointPort,omitempty"`
	ListenerPort *int64 `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
}

func (s CreateEndpointGroupsRequestEndpointGroupConfigurationsPortOverrides) String() string {
	return dara.Prettify(s)
}

func (s CreateEndpointGroupsRequestEndpointGroupConfigurationsPortOverrides) GoString() string {
	return s.String()
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurationsPortOverrides) GetEndpointPort() *int64 {
	return s.EndpointPort
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurationsPortOverrides) GetListenerPort() *int64 {
	return s.ListenerPort
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurationsPortOverrides) SetEndpointPort(v int64) *CreateEndpointGroupsRequestEndpointGroupConfigurationsPortOverrides {
	s.EndpointPort = &v
	return s
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurationsPortOverrides) SetListenerPort(v int64) *CreateEndpointGroupsRequestEndpointGroupConfigurationsPortOverrides {
	s.ListenerPort = &v
	return s
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurationsPortOverrides) Validate() error {
	return dara.Validate(s)
}

type CreateEndpointGroupsRequestEndpointGroupConfigurationsSystemTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEndpointGroupsRequestEndpointGroupConfigurationsSystemTag) String() string {
	return dara.Prettify(s)
}

func (s CreateEndpointGroupsRequestEndpointGroupConfigurationsSystemTag) GoString() string {
	return s.String()
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurationsSystemTag) GetKey() *string {
	return s.Key
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurationsSystemTag) GetScope() *string {
	return s.Scope
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurationsSystemTag) GetValue() *string {
	return s.Value
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurationsSystemTag) SetKey(v string) *CreateEndpointGroupsRequestEndpointGroupConfigurationsSystemTag {
	s.Key = &v
	return s
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurationsSystemTag) SetScope(v string) *CreateEndpointGroupsRequestEndpointGroupConfigurationsSystemTag {
	s.Scope = &v
	return s
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurationsSystemTag) SetValue(v string) *CreateEndpointGroupsRequestEndpointGroupConfigurationsSystemTag {
	s.Value = &v
	return s
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurationsSystemTag) Validate() error {
	return dara.Validate(s)
}

type CreateEndpointGroupsRequestEndpointGroupConfigurationsTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEndpointGroupsRequestEndpointGroupConfigurationsTag) String() string {
	return dara.Prettify(s)
}

func (s CreateEndpointGroupsRequestEndpointGroupConfigurationsTag) GoString() string {
	return s.String()
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurationsTag) GetKey() *string {
	return s.Key
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurationsTag) GetValue() *string {
	return s.Value
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurationsTag) SetKey(v string) *CreateEndpointGroupsRequestEndpointGroupConfigurationsTag {
	s.Key = &v
	return s
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurationsTag) SetValue(v string) *CreateEndpointGroupsRequestEndpointGroupConfigurationsTag {
	s.Value = &v
	return s
}

func (s *CreateEndpointGroupsRequestEndpointGroupConfigurationsTag) Validate() error {
	return dara.Validate(s)
}
