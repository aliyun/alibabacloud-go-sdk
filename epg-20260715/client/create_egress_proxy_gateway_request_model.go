// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEgressProxyGatewayRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateEgressProxyGatewayRequest
	GetClientToken() *string
	SetDescription(v string) *CreateEgressProxyGatewayRequest
	GetDescription() *string
	SetDryRun(v bool) *CreateEgressProxyGatewayRequest
	GetDryRun() *bool
	SetEgressProxyConfigurationId(v string) *CreateEgressProxyGatewayRequest
	GetEgressProxyConfigurationId() *string
	SetEgressProxyGatewayName(v string) *CreateEgressProxyGatewayRequest
	GetEgressProxyGatewayName() *string
	SetListenerProperties(v []*CreateEgressProxyGatewayRequestListenerProperties) *CreateEgressProxyGatewayRequest
	GetListenerProperties() []*CreateEgressProxyGatewayRequestListenerProperties
	SetResourceGroupId(v string) *CreateEgressProxyGatewayRequest
	GetResourceGroupId() *string
	SetTag(v []*CreateEgressProxyGatewayRequestTag) *CreateEgressProxyGatewayRequest
	GetTag() []*CreateEgressProxyGatewayRequestTag
	SetTlsInterceptProperties(v *CreateEgressProxyGatewayRequestTlsInterceptProperties) *CreateEgressProxyGatewayRequest
	GetTlsInterceptProperties() *CreateEgressProxyGatewayRequestTlsInterceptProperties
	SetVpcId(v string) *CreateEgressProxyGatewayRequest
	GetVpcId() *string
	SetZoneMappings(v []*CreateEgressProxyGatewayRequestZoneMappings) *CreateEgressProxyGatewayRequest
	GetZoneMappings() []*CreateEgressProxyGatewayRequestZoneMappings
}

type CreateEgressProxyGatewayRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must ensure that it is unique among different requests. The client token can contain only ASCII characters.
	//
	// > Note If you do not specify this parameter, the system automatically uses the RequestId of the API request as the client token. The RequestId of each API request is unique.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B******
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the egress proxy gateway instance.
	//
	// The description must be 2 to 256 characters in length, and can contain Chinese characters, letters, digits, commas (,), periods (.), semicolons (;), forward slashes (/), at signs (@), underscores (_), and hyphens (-).
	//
	// example:
	//
	// testEpg
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to only precheck the request. Valid values:
	//
	// - **true**: prechecks the request without creating the egress proxy gateway instance. The system checks whether required parameters are specified, the request format, and service limits. If the request fails the precheck, the corresponding error is returned. If the request passes the precheck, the DryRunOperation error code is returned.
	//
	// - **false*	- (default): sends the request. If the request passes the precheck, the HTTP 2xx status code is returned and the operation is performed.
	//
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the egress proxy configuration that is associated with the egress proxy gateway instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// epc-uf609a6a8fdf9f4****
	EgressProxyConfigurationId *string `json:"EgressProxyConfigurationId,omitempty" xml:"EgressProxyConfigurationId,omitempty"`
	// The name of the egress proxy gateway instance.
	//
	// The name must be 2 to 128 characters in length, and must start with a letter or a Chinese character. It can contain Chinese characters, letters, digits, periods (.), underscores (_), and hyphens (-).
	//
	// This parameter is required.
	//
	// example:
	//
	// testEpg
	EgressProxyGatewayName *string `json:"EgressProxyGatewayName,omitempty" xml:"EgressProxyGatewayName,omitempty"`
	// The listener configurations of the egress proxy gateway instance.
	//
	// Listeners in the same gateway instance cannot use the same port.
	//
	// This parameter is required.
	ListenerProperties []*CreateEgressProxyGatewayRequestListenerProperties `json:"ListenerProperties,omitempty" xml:"ListenerProperties,omitempty" type:"Repeated"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-gw80bb6a94f835db****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tags.
	Tag []*CreateEgressProxyGatewayRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The TLS interception configuration of the egress proxy gateway instance.
	TlsInterceptProperties *CreateEgressProxyGatewayRequestTlsInterceptProperties `json:"TlsInterceptProperties,omitempty" xml:"TlsInterceptProperties,omitempty" type:"Struct"`
	// The ID of the virtual private cloud (VPC) to which the egress proxy gateway instance belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-zm0ejg5unozkhvuhd****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The zone and vSwitch mappings of the egress proxy gateway instance.
	//
	// If the current region supports two or more zones, add at least two zones to improve availability.
	//
	// This parameter is required.
	ZoneMappings []*CreateEgressProxyGatewayRequestZoneMappings `json:"ZoneMappings,omitempty" xml:"ZoneMappings,omitempty" type:"Repeated"`
}

func (s CreateEgressProxyGatewayRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateEgressProxyGatewayRequest) GoString() string {
	return s.String()
}

func (s *CreateEgressProxyGatewayRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateEgressProxyGatewayRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateEgressProxyGatewayRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateEgressProxyGatewayRequest) GetEgressProxyConfigurationId() *string {
	return s.EgressProxyConfigurationId
}

func (s *CreateEgressProxyGatewayRequest) GetEgressProxyGatewayName() *string {
	return s.EgressProxyGatewayName
}

func (s *CreateEgressProxyGatewayRequest) GetListenerProperties() []*CreateEgressProxyGatewayRequestListenerProperties {
	return s.ListenerProperties
}

func (s *CreateEgressProxyGatewayRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateEgressProxyGatewayRequest) GetTag() []*CreateEgressProxyGatewayRequestTag {
	return s.Tag
}

func (s *CreateEgressProxyGatewayRequest) GetTlsInterceptProperties() *CreateEgressProxyGatewayRequestTlsInterceptProperties {
	return s.TlsInterceptProperties
}

func (s *CreateEgressProxyGatewayRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateEgressProxyGatewayRequest) GetZoneMappings() []*CreateEgressProxyGatewayRequestZoneMappings {
	return s.ZoneMappings
}

func (s *CreateEgressProxyGatewayRequest) SetClientToken(v string) *CreateEgressProxyGatewayRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateEgressProxyGatewayRequest) SetDescription(v string) *CreateEgressProxyGatewayRequest {
	s.Description = &v
	return s
}

func (s *CreateEgressProxyGatewayRequest) SetDryRun(v bool) *CreateEgressProxyGatewayRequest {
	s.DryRun = &v
	return s
}

func (s *CreateEgressProxyGatewayRequest) SetEgressProxyConfigurationId(v string) *CreateEgressProxyGatewayRequest {
	s.EgressProxyConfigurationId = &v
	return s
}

func (s *CreateEgressProxyGatewayRequest) SetEgressProxyGatewayName(v string) *CreateEgressProxyGatewayRequest {
	s.EgressProxyGatewayName = &v
	return s
}

func (s *CreateEgressProxyGatewayRequest) SetListenerProperties(v []*CreateEgressProxyGatewayRequestListenerProperties) *CreateEgressProxyGatewayRequest {
	s.ListenerProperties = v
	return s
}

func (s *CreateEgressProxyGatewayRequest) SetResourceGroupId(v string) *CreateEgressProxyGatewayRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateEgressProxyGatewayRequest) SetTag(v []*CreateEgressProxyGatewayRequestTag) *CreateEgressProxyGatewayRequest {
	s.Tag = v
	return s
}

func (s *CreateEgressProxyGatewayRequest) SetTlsInterceptProperties(v *CreateEgressProxyGatewayRequestTlsInterceptProperties) *CreateEgressProxyGatewayRequest {
	s.TlsInterceptProperties = v
	return s
}

func (s *CreateEgressProxyGatewayRequest) SetVpcId(v string) *CreateEgressProxyGatewayRequest {
	s.VpcId = &v
	return s
}

func (s *CreateEgressProxyGatewayRequest) SetZoneMappings(v []*CreateEgressProxyGatewayRequestZoneMappings) *CreateEgressProxyGatewayRequest {
	s.ZoneMappings = v
	return s
}

func (s *CreateEgressProxyGatewayRequest) Validate() error {
	if s.ListenerProperties != nil {
		for _, item := range s.ListenerProperties {
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
	if s.TlsInterceptProperties != nil {
		if err := s.TlsInterceptProperties.Validate(); err != nil {
			return err
		}
	}
	if s.ZoneMappings != nil {
		for _, item := range s.ZoneMappings {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateEgressProxyGatewayRequestListenerProperties struct {
	// The listener port.
	//
	// Valid values: **1*	- to **65535**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8080
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The protocol type of the listener. Valid values:
	//
	// - **HTTP**: HTTP.
	//
	// - **HTTPS**: HTTPS.
	//
	// This parameter is required.
	//
	// example:
	//
	// HTTP
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
}

func (s CreateEgressProxyGatewayRequestListenerProperties) String() string {
	return dara.Prettify(s)
}

func (s CreateEgressProxyGatewayRequestListenerProperties) GoString() string {
	return s.String()
}

func (s *CreateEgressProxyGatewayRequestListenerProperties) GetPort() *int32 {
	return s.Port
}

func (s *CreateEgressProxyGatewayRequestListenerProperties) GetProtocol() *string {
	return s.Protocol
}

func (s *CreateEgressProxyGatewayRequestListenerProperties) SetPort(v int32) *CreateEgressProxyGatewayRequestListenerProperties {
	s.Port = &v
	return s
}

func (s *CreateEgressProxyGatewayRequestListenerProperties) SetProtocol(v string) *CreateEgressProxyGatewayRequestListenerProperties {
	s.Protocol = &v
	return s
}

func (s *CreateEgressProxyGatewayRequestListenerProperties) Validate() error {
	return dara.Validate(s)
}

type CreateEgressProxyGatewayRequestTag struct {
	// The tag key. The key can be up to 128 characters in length. It cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// testKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. The value can be up to 256 characters in length, and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// testValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateEgressProxyGatewayRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateEgressProxyGatewayRequestTag) GoString() string {
	return s.String()
}

func (s *CreateEgressProxyGatewayRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateEgressProxyGatewayRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateEgressProxyGatewayRequestTag) SetKey(v string) *CreateEgressProxyGatewayRequestTag {
	s.Key = &v
	return s
}

func (s *CreateEgressProxyGatewayRequestTag) SetValue(v string) *CreateEgressProxyGatewayRequestTag {
	s.Value = &v
	return s
}

func (s *CreateEgressProxyGatewayRequestTag) Validate() error {
	return dara.Validate(s)
}

type CreateEgressProxyGatewayRequestTlsInterceptProperties struct {
	// The ID of the private CA certificate.
	//
	// - This parameter is required when **TlsInterceptMode*	- is set to **enable**.
	//
	// - Do not specify this parameter when **TlsInterceptMode*	- is set to **disable**.
	//
	// example:
	//
	// 1f01a6f9-a0e7-64e5-9d3c-29c056bf****
	PrivateCaCertificateId *string `json:"PrivateCaCertificateId,omitempty" xml:"PrivateCaCertificateId,omitempty"`
	// Specifies whether to enable TLS interception. Valid values:
	//
	// - **enable**: enables TLS interception.
	//
	// - **disable**: disables TLS interception.
	//
	// This parameter is required.
	//
	// example:
	//
	// enable
	TlsInterceptMode *string `json:"TlsInterceptMode,omitempty" xml:"TlsInterceptMode,omitempty"`
}

func (s CreateEgressProxyGatewayRequestTlsInterceptProperties) String() string {
	return dara.Prettify(s)
}

func (s CreateEgressProxyGatewayRequestTlsInterceptProperties) GoString() string {
	return s.String()
}

func (s *CreateEgressProxyGatewayRequestTlsInterceptProperties) GetPrivateCaCertificateId() *string {
	return s.PrivateCaCertificateId
}

func (s *CreateEgressProxyGatewayRequestTlsInterceptProperties) GetTlsInterceptMode() *string {
	return s.TlsInterceptMode
}

func (s *CreateEgressProxyGatewayRequestTlsInterceptProperties) SetPrivateCaCertificateId(v string) *CreateEgressProxyGatewayRequestTlsInterceptProperties {
	s.PrivateCaCertificateId = &v
	return s
}

func (s *CreateEgressProxyGatewayRequestTlsInterceptProperties) SetTlsInterceptMode(v string) *CreateEgressProxyGatewayRequestTlsInterceptProperties {
	s.TlsInterceptMode = &v
	return s
}

func (s *CreateEgressProxyGatewayRequestTlsInterceptProperties) Validate() error {
	return dara.Validate(s)
}

type CreateEgressProxyGatewayRequestZoneMappings struct {
	// The ID of the vSwitch that corresponds to the zone.
	//
	// The vSwitch must belong to the specified VPC and match the zone specified by ZoneId.
	//
	// This parameter is required.
	//
	// example:
	//
	// vsw-wz9aejevoin852mo****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The zone ID.
	//
	// You can call the DescribeZones operation to query information about zones in a region.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai-m
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateEgressProxyGatewayRequestZoneMappings) String() string {
	return dara.Prettify(s)
}

func (s CreateEgressProxyGatewayRequestZoneMappings) GoString() string {
	return s.String()
}

func (s *CreateEgressProxyGatewayRequestZoneMappings) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateEgressProxyGatewayRequestZoneMappings) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateEgressProxyGatewayRequestZoneMappings) SetVSwitchId(v string) *CreateEgressProxyGatewayRequestZoneMappings {
	s.VSwitchId = &v
	return s
}

func (s *CreateEgressProxyGatewayRequestZoneMappings) SetZoneId(v string) *CreateEgressProxyGatewayRequestZoneMappings {
	s.ZoneId = &v
	return s
}

func (s *CreateEgressProxyGatewayRequestZoneMappings) Validate() error {
	return dara.Validate(s)
}
