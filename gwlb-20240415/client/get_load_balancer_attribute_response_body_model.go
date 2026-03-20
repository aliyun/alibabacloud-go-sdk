// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLoadBalancerAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAddressIpVersion(v string) *GetLoadBalancerAttributeResponseBody
	GetAddressIpVersion() *string
	SetCreateTime(v string) *GetLoadBalancerAttributeResponseBody
	GetCreateTime() *string
	SetLoadBalancerBusinessStatus(v string) *GetLoadBalancerAttributeResponseBody
	GetLoadBalancerBusinessStatus() *string
	SetLoadBalancerId(v string) *GetLoadBalancerAttributeResponseBody
	GetLoadBalancerId() *string
	SetLoadBalancerName(v string) *GetLoadBalancerAttributeResponseBody
	GetLoadBalancerName() *string
	SetLoadBalancerStatus(v string) *GetLoadBalancerAttributeResponseBody
	GetLoadBalancerStatus() *string
	SetRequestId(v string) *GetLoadBalancerAttributeResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *GetLoadBalancerAttributeResponseBody
	GetResourceGroupId() *string
	SetTags(v []*GetLoadBalancerAttributeResponseBodyTags) *GetLoadBalancerAttributeResponseBody
	GetTags() []*GetLoadBalancerAttributeResponseBodyTags
	SetTrafficMode(v string) *GetLoadBalancerAttributeResponseBody
	GetTrafficMode() *string
	SetVpcId(v string) *GetLoadBalancerAttributeResponseBody
	GetVpcId() *string
	SetZoneMappings(v []*GetLoadBalancerAttributeResponseBodyZoneMappings) *GetLoadBalancerAttributeResponseBody
	GetZoneMappings() []*GetLoadBalancerAttributeResponseBodyZoneMappings
}

type GetLoadBalancerAttributeResponseBody struct {
	// The protocol version. Valid values:
	//
	// 	- **Ipv4**: IPv4.
	//
	// example:
	//
	// IPv4
	AddressIpVersion *string `json:"AddressIpVersion,omitempty" xml:"AddressIpVersion,omitempty"`
	// The time when the resource was created. The time follows the ISO 8601 standard in the **yyyy-MM-ddTHH:mm:ssZ*	- format. The time is displayed in UTC.
	//
	// example:
	//
	// 2024-07-08T10:12:58Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The business status of the GWLB instance. Valid values:
	//
	// 	- **Normal**: running as expected
	//
	// 	- **FinancialLocked**: locked due to overdue payments
	//
	// example:
	//
	// Normal
	LoadBalancerBusinessStatus *string `json:"LoadBalancerBusinessStatus,omitempty" xml:"LoadBalancerBusinessStatus,omitempty"`
	// The GWLB instance ID.
	//
	// example:
	//
	// gwlb-9njtjmqt7zfcqm****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The GWLB instance name.
	//
	// The name must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter.
	//
	// example:
	//
	// gwlb
	LoadBalancerName *string `json:"LoadBalancerName,omitempty" xml:"LoadBalancerName,omitempty"`
	// The GWLB instance status. Valid values:
	//
	// 	- **Active**: The GWLB instance is running.
	//
	// 	- **Inactive**: The GWLB instance is disabled. Listeners of GWLB instances in the Inactive state do not forward traffic.
	//
	// 	- **Provisioning**: The GWLB instance is being created.
	//
	// 	- **Configuring**: The GWLB instance is being modified.
	//
	// example:
	//
	// Active
	LoadBalancerStatus *string `json:"LoadBalancerStatus,omitempty" xml:"LoadBalancerStatus,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B6DC5DDC-9560-59BF-80FA-ED1E5CB417DF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmx7pmxcy****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tags.
	Tags []*GetLoadBalancerAttributeResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// Traffic processing mode. Valid values:
	//
	// 	- **LoadBalance**: load balancing mode. GWLB forwards traffic to backend servers.
	//
	// 	- **ByPass**: bypass mode. GWLB directly returns traffic to the GWLB endpoint without forwarding it to the backend servers.
	//
	// example:
	//
	// LoadBalance
	TrafficMode *string `json:"TrafficMode,omitempty" xml:"TrafficMode,omitempty"`
	// The VPC ID.
	//
	// example:
	//
	// vpc-k1aajsbwbaq4todet****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The mappings between zones and vSwitches. You must specify at least one zone. You can specify at most 20 zones. If the region supports two or more zones, specify at least two zones.
	ZoneMappings []*GetLoadBalancerAttributeResponseBodyZoneMappings `json:"ZoneMappings,omitempty" xml:"ZoneMappings,omitempty" type:"Repeated"`
}

func (s GetLoadBalancerAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetLoadBalancerAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *GetLoadBalancerAttributeResponseBody) GetAddressIpVersion() *string {
	return s.AddressIpVersion
}

func (s *GetLoadBalancerAttributeResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetLoadBalancerAttributeResponseBody) GetLoadBalancerBusinessStatus() *string {
	return s.LoadBalancerBusinessStatus
}

func (s *GetLoadBalancerAttributeResponseBody) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *GetLoadBalancerAttributeResponseBody) GetLoadBalancerName() *string {
	return s.LoadBalancerName
}

func (s *GetLoadBalancerAttributeResponseBody) GetLoadBalancerStatus() *string {
	return s.LoadBalancerStatus
}

func (s *GetLoadBalancerAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetLoadBalancerAttributeResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetLoadBalancerAttributeResponseBody) GetTags() []*GetLoadBalancerAttributeResponseBodyTags {
	return s.Tags
}

func (s *GetLoadBalancerAttributeResponseBody) GetTrafficMode() *string {
	return s.TrafficMode
}

func (s *GetLoadBalancerAttributeResponseBody) GetVpcId() *string {
	return s.VpcId
}

func (s *GetLoadBalancerAttributeResponseBody) GetZoneMappings() []*GetLoadBalancerAttributeResponseBodyZoneMappings {
	return s.ZoneMappings
}

func (s *GetLoadBalancerAttributeResponseBody) SetAddressIpVersion(v string) *GetLoadBalancerAttributeResponseBody {
	s.AddressIpVersion = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetCreateTime(v string) *GetLoadBalancerAttributeResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetLoadBalancerBusinessStatus(v string) *GetLoadBalancerAttributeResponseBody {
	s.LoadBalancerBusinessStatus = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetLoadBalancerId(v string) *GetLoadBalancerAttributeResponseBody {
	s.LoadBalancerId = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetLoadBalancerName(v string) *GetLoadBalancerAttributeResponseBody {
	s.LoadBalancerName = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetLoadBalancerStatus(v string) *GetLoadBalancerAttributeResponseBody {
	s.LoadBalancerStatus = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetRequestId(v string) *GetLoadBalancerAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetResourceGroupId(v string) *GetLoadBalancerAttributeResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetTags(v []*GetLoadBalancerAttributeResponseBodyTags) *GetLoadBalancerAttributeResponseBody {
	s.Tags = v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetTrafficMode(v string) *GetLoadBalancerAttributeResponseBody {
	s.TrafficMode = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetVpcId(v string) *GetLoadBalancerAttributeResponseBody {
	s.VpcId = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) SetZoneMappings(v []*GetLoadBalancerAttributeResponseBodyZoneMappings) *GetLoadBalancerAttributeResponseBody {
	s.ZoneMappings = v
	return s
}

func (s *GetLoadBalancerAttributeResponseBody) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
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

type GetLoadBalancerAttributeResponseBodyTags struct {
	// The tag key. The tag key cannot be an empty string.
	//
	// The tag key can be up to 128 characters in length. The tag key cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// testTagKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. The tag value can be up to 256 characters in length and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// testTagValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetLoadBalancerAttributeResponseBodyTags) String() string {
	return dara.Prettify(s)
}

func (s GetLoadBalancerAttributeResponseBodyTags) GoString() string {
	return s.String()
}

func (s *GetLoadBalancerAttributeResponseBodyTags) GetKey() *string {
	return s.Key
}

func (s *GetLoadBalancerAttributeResponseBodyTags) GetValue() *string {
	return s.Value
}

func (s *GetLoadBalancerAttributeResponseBodyTags) SetKey(v string) *GetLoadBalancerAttributeResponseBodyTags {
	s.Key = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBodyTags) SetValue(v string) *GetLoadBalancerAttributeResponseBodyTags {
	s.Value = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBodyTags) Validate() error {
	return dara.Validate(s)
}

type GetLoadBalancerAttributeResponseBodyZoneMappings struct {
	// The GWLB instance addresses.
	LoadBalancerAddresses []*GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses `json:"LoadBalancerAddresses,omitempty" xml:"LoadBalancerAddresses,omitempty" type:"Repeated"`
	// The vSwitch in the zone. You can specify only one vSwitch (subnet) in each zone of a GWLB instance.
	//
	// example:
	//
	// vsw-uf6v8l7d2f1k53xrl****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-hangzhou-j
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s GetLoadBalancerAttributeResponseBodyZoneMappings) String() string {
	return dara.Prettify(s)
}

func (s GetLoadBalancerAttributeResponseBodyZoneMappings) GoString() string {
	return s.String()
}

func (s *GetLoadBalancerAttributeResponseBodyZoneMappings) GetLoadBalancerAddresses() []*GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses {
	return s.LoadBalancerAddresses
}

func (s *GetLoadBalancerAttributeResponseBodyZoneMappings) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *GetLoadBalancerAttributeResponseBodyZoneMappings) GetZoneId() *string {
	return s.ZoneId
}

func (s *GetLoadBalancerAttributeResponseBodyZoneMappings) SetLoadBalancerAddresses(v []*GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses) *GetLoadBalancerAttributeResponseBodyZoneMappings {
	s.LoadBalancerAddresses = v
	return s
}

func (s *GetLoadBalancerAttributeResponseBodyZoneMappings) SetVSwitchId(v string) *GetLoadBalancerAttributeResponseBodyZoneMappings {
	s.VSwitchId = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBodyZoneMappings) SetZoneId(v string) *GetLoadBalancerAttributeResponseBodyZoneMappings {
	s.ZoneId = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBodyZoneMappings) Validate() error {
	if s.LoadBalancerAddresses != nil {
		for _, item := range s.LoadBalancerAddresses {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses struct {
	// The ID of the elastic network interface (ENI) used by the GWLB instance.
	//
	// example:
	//
	// eni-bp1iahwz3rzgvltz****
	EniId *string `json:"EniId,omitempty" xml:"EniId,omitempty"`
	// The private IPv4 address.
	//
	// example:
	//
	// 192.168.XX.XX
	PrivateIpv4Address *string `json:"PrivateIpv4Address,omitempty" xml:"PrivateIpv4Address,omitempty"`
}

func (s GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses) String() string {
	return dara.Prettify(s)
}

func (s GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses) GoString() string {
	return s.String()
}

func (s *GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses) GetEniId() *string {
	return s.EniId
}

func (s *GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses) GetPrivateIpv4Address() *string {
	return s.PrivateIpv4Address
}

func (s *GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses) SetEniId(v string) *GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses {
	s.EniId = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses) SetPrivateIpv4Address(v string) *GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses {
	s.PrivateIpv4Address = &v
	return s
}

func (s *GetLoadBalancerAttributeResponseBodyZoneMappingsLoadBalancerAddresses) Validate() error {
	return dara.Validate(s)
}
