// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEgressProxyGatewayAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v string) *GetEgressProxyGatewayAttributeResponseBody
	GetCreateTime() *string
	SetDefaultEgressProxyIpPoolId(v string) *GetEgressProxyGatewayAttributeResponseBody
	GetDefaultEgressProxyIpPoolId() *string
	SetDescription(v string) *GetEgressProxyGatewayAttributeResponseBody
	GetDescription() *string
	SetDnsName(v string) *GetEgressProxyGatewayAttributeResponseBody
	GetDnsName() *string
	SetEgressProxyConfigurationId(v string) *GetEgressProxyGatewayAttributeResponseBody
	GetEgressProxyConfigurationId() *string
	SetEgressProxyGatewayId(v string) *GetEgressProxyGatewayAttributeResponseBody
	GetEgressProxyGatewayId() *string
	SetEgressProxyGatewayName(v string) *GetEgressProxyGatewayAttributeResponseBody
	GetEgressProxyGatewayName() *string
	SetListenerProperties(v []*GetEgressProxyGatewayAttributeResponseBodyListenerProperties) *GetEgressProxyGatewayAttributeResponseBody
	GetListenerProperties() []*GetEgressProxyGatewayAttributeResponseBodyListenerProperties
	SetRequestId(v string) *GetEgressProxyGatewayAttributeResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *GetEgressProxyGatewayAttributeResponseBody
	GetResourceGroupId() *string
	SetStatus(v string) *GetEgressProxyGatewayAttributeResponseBody
	GetStatus() *string
	SetTags(v []*GetEgressProxyGatewayAttributeResponseBodyTags) *GetEgressProxyGatewayAttributeResponseBody
	GetTags() []*GetEgressProxyGatewayAttributeResponseBodyTags
	SetTlsInterceptProperties(v *GetEgressProxyGatewayAttributeResponseBodyTlsInterceptProperties) *GetEgressProxyGatewayAttributeResponseBody
	GetTlsInterceptProperties() *GetEgressProxyGatewayAttributeResponseBodyTlsInterceptProperties
	SetVpcId(v string) *GetEgressProxyGatewayAttributeResponseBody
	GetVpcId() *string
	SetZoneMappings(v []*GetEgressProxyGatewayAttributeResponseBodyZoneMappings) *GetEgressProxyGatewayAttributeResponseBody
	GetZoneMappings() []*GetEgressProxyGatewayAttributeResponseBodyZoneMappings
}

type GetEgressProxyGatewayAttributeResponseBody struct {
	// The time when the egress proxy gateway was created. The time is displayed in UTC and follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format.
	//
	// example:
	//
	// 2022-07-02T02:49:05Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the default egress IP pool.
	//
	// example:
	//
	// eppool-gw80d96a96ca0b00****
	DefaultEgressProxyIpPoolId *string `json:"DefaultEgressProxyIpPoolId,omitempty" xml:"DefaultEgressProxyIpPoolId,omitempty"`
	// The description of the egress proxy gateway instance. The description must be 2 to 256 characters in length and can contain Chinese characters and English letters. It can also contain digits, commas (,), periods (.), semicolons (;), forward slashes (/), at signs (@), underscores (_), and hyphens (-).
	//
	// example:
	//
	// testEpg
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The DNS domain name.
	//
	// example:
	//
	// epg-gw80bb6a94f835****.cn-shanghai.epg.aliyuncs.com
	DnsName *string `json:"DnsName,omitempty" xml:"DnsName,omitempty"`
	// The ID of the associated egress proxy configuration.
	//
	// example:
	//
	// epc-uf609a6a8fdf9f4d****
	EgressProxyConfigurationId *string `json:"EgressProxyConfigurationId,omitempty" xml:"EgressProxyConfigurationId,omitempty"`
	// The ID of the egress proxy gateway instance.
	//
	// example:
	//
	// epg-gw80bb6a94f835db****
	EgressProxyGatewayId *string `json:"EgressProxyGatewayId,omitempty" xml:"EgressProxyGatewayId,omitempty"`
	// The name of the egress proxy gateway. The name must be 2 to 128 characters in length and can contain English letters and Chinese characters. It must start with a letter or a Chinese character and can contain digits, periods (.), underscores (_), and hyphens (-).
	//
	// example:
	//
	// testEpg
	EgressProxyGatewayName *string `json:"EgressProxyGatewayName,omitempty" xml:"EgressProxyGatewayName,omitempty"`
	// The list of listener properties.
	ListenerProperties []*GetEgressProxyGatewayAttributeResponseBodyListenerProperties `json:"ListenerProperties,omitempty" xml:"ListenerProperties,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 01A05AB7-6C14-5A12-94E7-F20F9F69AE3D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-aek2zye42m****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The instance status of the egress proxy gateway. Valid values:
	//
	// - **Provisioning**: The instance is being created.
	//
	// - **Active**: The instance is running.
	//
	// - **Configuring**: The instance is being reconfigured.
	//
	// - **Inactive**: The instance is stopped.
	//
	// example:
	//
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The list of tags.
	Tags []*GetEgressProxyGatewayAttributeResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The TLS interception properties.
	TlsInterceptProperties *GetEgressProxyGatewayAttributeResponseBodyTlsInterceptProperties `json:"TlsInterceptProperties,omitempty" xml:"TlsInterceptProperties,omitempty" type:"Struct"`
	// The ID of the virtual private cloud (VPC).
	//
	// example:
	//
	// vpc-n4appznmjjr362yd3****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The list of zone and vSwitch mappings. If the current region supports two or more zones, at least two zones must be returned.
	ZoneMappings []*GetEgressProxyGatewayAttributeResponseBodyZoneMappings `json:"ZoneMappings,omitempty" xml:"ZoneMappings,omitempty" type:"Repeated"`
}

func (s GetEgressProxyGatewayAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetEgressProxyGatewayAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *GetEgressProxyGatewayAttributeResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetEgressProxyGatewayAttributeResponseBody) GetDefaultEgressProxyIpPoolId() *string {
	return s.DefaultEgressProxyIpPoolId
}

func (s *GetEgressProxyGatewayAttributeResponseBody) GetDescription() *string {
	return s.Description
}

func (s *GetEgressProxyGatewayAttributeResponseBody) GetDnsName() *string {
	return s.DnsName
}

func (s *GetEgressProxyGatewayAttributeResponseBody) GetEgressProxyConfigurationId() *string {
	return s.EgressProxyConfigurationId
}

func (s *GetEgressProxyGatewayAttributeResponseBody) GetEgressProxyGatewayId() *string {
	return s.EgressProxyGatewayId
}

func (s *GetEgressProxyGatewayAttributeResponseBody) GetEgressProxyGatewayName() *string {
	return s.EgressProxyGatewayName
}

func (s *GetEgressProxyGatewayAttributeResponseBody) GetListenerProperties() []*GetEgressProxyGatewayAttributeResponseBodyListenerProperties {
	return s.ListenerProperties
}

func (s *GetEgressProxyGatewayAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetEgressProxyGatewayAttributeResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetEgressProxyGatewayAttributeResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetEgressProxyGatewayAttributeResponseBody) GetTags() []*GetEgressProxyGatewayAttributeResponseBodyTags {
	return s.Tags
}

func (s *GetEgressProxyGatewayAttributeResponseBody) GetTlsInterceptProperties() *GetEgressProxyGatewayAttributeResponseBodyTlsInterceptProperties {
	return s.TlsInterceptProperties
}

func (s *GetEgressProxyGatewayAttributeResponseBody) GetVpcId() *string {
	return s.VpcId
}

func (s *GetEgressProxyGatewayAttributeResponseBody) GetZoneMappings() []*GetEgressProxyGatewayAttributeResponseBodyZoneMappings {
	return s.ZoneMappings
}

func (s *GetEgressProxyGatewayAttributeResponseBody) SetCreateTime(v string) *GetEgressProxyGatewayAttributeResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetEgressProxyGatewayAttributeResponseBody) SetDefaultEgressProxyIpPoolId(v string) *GetEgressProxyGatewayAttributeResponseBody {
	s.DefaultEgressProxyIpPoolId = &v
	return s
}

func (s *GetEgressProxyGatewayAttributeResponseBody) SetDescription(v string) *GetEgressProxyGatewayAttributeResponseBody {
	s.Description = &v
	return s
}

func (s *GetEgressProxyGatewayAttributeResponseBody) SetDnsName(v string) *GetEgressProxyGatewayAttributeResponseBody {
	s.DnsName = &v
	return s
}

func (s *GetEgressProxyGatewayAttributeResponseBody) SetEgressProxyConfigurationId(v string) *GetEgressProxyGatewayAttributeResponseBody {
	s.EgressProxyConfigurationId = &v
	return s
}

func (s *GetEgressProxyGatewayAttributeResponseBody) SetEgressProxyGatewayId(v string) *GetEgressProxyGatewayAttributeResponseBody {
	s.EgressProxyGatewayId = &v
	return s
}

func (s *GetEgressProxyGatewayAttributeResponseBody) SetEgressProxyGatewayName(v string) *GetEgressProxyGatewayAttributeResponseBody {
	s.EgressProxyGatewayName = &v
	return s
}

func (s *GetEgressProxyGatewayAttributeResponseBody) SetListenerProperties(v []*GetEgressProxyGatewayAttributeResponseBodyListenerProperties) *GetEgressProxyGatewayAttributeResponseBody {
	s.ListenerProperties = v
	return s
}

func (s *GetEgressProxyGatewayAttributeResponseBody) SetRequestId(v string) *GetEgressProxyGatewayAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetEgressProxyGatewayAttributeResponseBody) SetResourceGroupId(v string) *GetEgressProxyGatewayAttributeResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *GetEgressProxyGatewayAttributeResponseBody) SetStatus(v string) *GetEgressProxyGatewayAttributeResponseBody {
	s.Status = &v
	return s
}

func (s *GetEgressProxyGatewayAttributeResponseBody) SetTags(v []*GetEgressProxyGatewayAttributeResponseBodyTags) *GetEgressProxyGatewayAttributeResponseBody {
	s.Tags = v
	return s
}

func (s *GetEgressProxyGatewayAttributeResponseBody) SetTlsInterceptProperties(v *GetEgressProxyGatewayAttributeResponseBodyTlsInterceptProperties) *GetEgressProxyGatewayAttributeResponseBody {
	s.TlsInterceptProperties = v
	return s
}

func (s *GetEgressProxyGatewayAttributeResponseBody) SetVpcId(v string) *GetEgressProxyGatewayAttributeResponseBody {
	s.VpcId = &v
	return s
}

func (s *GetEgressProxyGatewayAttributeResponseBody) SetZoneMappings(v []*GetEgressProxyGatewayAttributeResponseBodyZoneMappings) *GetEgressProxyGatewayAttributeResponseBody {
	s.ZoneMappings = v
	return s
}

func (s *GetEgressProxyGatewayAttributeResponseBody) Validate() error {
	if s.ListenerProperties != nil {
		for _, item := range s.ListenerProperties {
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

type GetEgressProxyGatewayAttributeResponseBodyListenerProperties struct {
	// The listener port. Valid values: 1 to 65535.
	//
	// example:
	//
	// 8080
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The listener protocol. Valid values:
	//
	// - **HTTP**
	//
	// - **HTTPS**
	//
	// example:
	//
	// HTTP
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
}

func (s GetEgressProxyGatewayAttributeResponseBodyListenerProperties) String() string {
	return dara.Prettify(s)
}

func (s GetEgressProxyGatewayAttributeResponseBodyListenerProperties) GoString() string {
	return s.String()
}

func (s *GetEgressProxyGatewayAttributeResponseBodyListenerProperties) GetPort() *int32 {
	return s.Port
}

func (s *GetEgressProxyGatewayAttributeResponseBodyListenerProperties) GetProtocol() *string {
	return s.Protocol
}

func (s *GetEgressProxyGatewayAttributeResponseBodyListenerProperties) SetPort(v int32) *GetEgressProxyGatewayAttributeResponseBodyListenerProperties {
	s.Port = &v
	return s
}

func (s *GetEgressProxyGatewayAttributeResponseBodyListenerProperties) SetProtocol(v string) *GetEgressProxyGatewayAttributeResponseBodyListenerProperties {
	s.Protocol = &v
	return s
}

func (s *GetEgressProxyGatewayAttributeResponseBodyListenerProperties) Validate() error {
	return dara.Validate(s)
}

type GetEgressProxyGatewayAttributeResponseBodyTags struct {
	// The tag key. The key can be up to 128 characters in length. It cannot start with aliyun or acs: and cannot contain http:// or https://.
	//
	// example:
	//
	// testKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. The value can be up to 256 characters in length and cannot contain http:// or https://.
	//
	// example:
	//
	// testValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetEgressProxyGatewayAttributeResponseBodyTags) String() string {
	return dara.Prettify(s)
}

func (s GetEgressProxyGatewayAttributeResponseBodyTags) GoString() string {
	return s.String()
}

func (s *GetEgressProxyGatewayAttributeResponseBodyTags) GetKey() *string {
	return s.Key
}

func (s *GetEgressProxyGatewayAttributeResponseBodyTags) GetValue() *string {
	return s.Value
}

func (s *GetEgressProxyGatewayAttributeResponseBodyTags) SetKey(v string) *GetEgressProxyGatewayAttributeResponseBodyTags {
	s.Key = &v
	return s
}

func (s *GetEgressProxyGatewayAttributeResponseBodyTags) SetValue(v string) *GetEgressProxyGatewayAttributeResponseBodyTags {
	s.Value = &v
	return s
}

func (s *GetEgressProxyGatewayAttributeResponseBodyTags) Validate() error {
	return dara.Validate(s)
}

type GetEgressProxyGatewayAttributeResponseBodyTlsInterceptProperties struct {
	// The ID of the private CA certificate.
	//
	// example:
	//
	// 1f01a6f9-a0e7-64e5-9d3c-29c056bf****
	PrivateCaCertificateId *string `json:"PrivateCaCertificateId,omitempty" xml:"PrivateCaCertificateId,omitempty"`
	// The TLS Block Mode. Valid values:
	//
	// - **enable**: Enables the TLS Block Mode.
	//
	// - **disable**: Disables the TLS Block Mode.
	//
	// example:
	//
	// enable
	TlsInterceptMode *string `json:"TlsInterceptMode,omitempty" xml:"TlsInterceptMode,omitempty"`
}

func (s GetEgressProxyGatewayAttributeResponseBodyTlsInterceptProperties) String() string {
	return dara.Prettify(s)
}

func (s GetEgressProxyGatewayAttributeResponseBodyTlsInterceptProperties) GoString() string {
	return s.String()
}

func (s *GetEgressProxyGatewayAttributeResponseBodyTlsInterceptProperties) GetPrivateCaCertificateId() *string {
	return s.PrivateCaCertificateId
}

func (s *GetEgressProxyGatewayAttributeResponseBodyTlsInterceptProperties) GetTlsInterceptMode() *string {
	return s.TlsInterceptMode
}

func (s *GetEgressProxyGatewayAttributeResponseBodyTlsInterceptProperties) SetPrivateCaCertificateId(v string) *GetEgressProxyGatewayAttributeResponseBodyTlsInterceptProperties {
	s.PrivateCaCertificateId = &v
	return s
}

func (s *GetEgressProxyGatewayAttributeResponseBodyTlsInterceptProperties) SetTlsInterceptMode(v string) *GetEgressProxyGatewayAttributeResponseBodyTlsInterceptProperties {
	s.TlsInterceptMode = &v
	return s
}

func (s *GetEgressProxyGatewayAttributeResponseBodyTlsInterceptProperties) Validate() error {
	return dara.Validate(s)
}

type GetEgressProxyGatewayAttributeResponseBodyZoneMappings struct {
	// The zone status. Valid values:
	//
	// - **Provisioning**: The zone is being created.
	//
	// - **Available**: The zone is available.
	//
	// - **Configuring**: The zone is being reconfigured.
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the vSwitch that corresponds to the zone.
	//
	// example:
	//
	// vsw-f8zvrktbjs9hh3xqz****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-shanghai-b
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s GetEgressProxyGatewayAttributeResponseBodyZoneMappings) String() string {
	return dara.Prettify(s)
}

func (s GetEgressProxyGatewayAttributeResponseBodyZoneMappings) GoString() string {
	return s.String()
}

func (s *GetEgressProxyGatewayAttributeResponseBodyZoneMappings) GetStatus() *string {
	return s.Status
}

func (s *GetEgressProxyGatewayAttributeResponseBodyZoneMappings) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *GetEgressProxyGatewayAttributeResponseBodyZoneMappings) GetZoneId() *string {
	return s.ZoneId
}

func (s *GetEgressProxyGatewayAttributeResponseBodyZoneMappings) SetStatus(v string) *GetEgressProxyGatewayAttributeResponseBodyZoneMappings {
	s.Status = &v
	return s
}

func (s *GetEgressProxyGatewayAttributeResponseBodyZoneMappings) SetVSwitchId(v string) *GetEgressProxyGatewayAttributeResponseBodyZoneMappings {
	s.VSwitchId = &v
	return s
}

func (s *GetEgressProxyGatewayAttributeResponseBodyZoneMappings) SetZoneId(v string) *GetEgressProxyGatewayAttributeResponseBodyZoneMappings {
	s.ZoneId = &v
	return s
}

func (s *GetEgressProxyGatewayAttributeResponseBodyZoneMappings) Validate() error {
	return dara.Validate(s)
}
