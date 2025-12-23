// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceConfigurationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *GetResourceConfigurationResponseBody
	GetAccountId() *string
	SetConfiguration(v map[string]interface{}) *GetResourceConfigurationResponseBody
	GetConfiguration() map[string]interface{}
	SetCreateTime(v string) *GetResourceConfigurationResponseBody
	GetCreateTime() *string
	SetExpireTime(v string) *GetResourceConfigurationResponseBody
	GetExpireTime() *string
	SetIpAddressAttributes(v []*GetResourceConfigurationResponseBodyIpAddressAttributes) *GetResourceConfigurationResponseBody
	GetIpAddressAttributes() []*GetResourceConfigurationResponseBodyIpAddressAttributes
	SetIpAddresses(v []*string) *GetResourceConfigurationResponseBody
	GetIpAddresses() []*string
	SetRegionId(v string) *GetResourceConfigurationResponseBody
	GetRegionId() *string
	SetRequestId(v string) *GetResourceConfigurationResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *GetResourceConfigurationResponseBody
	GetResourceGroupId() *string
	SetResourceId(v string) *GetResourceConfigurationResponseBody
	GetResourceId() *string
	SetResourceName(v string) *GetResourceConfigurationResponseBody
	GetResourceName() *string
	SetResourceType(v string) *GetResourceConfigurationResponseBody
	GetResourceType() *string
	SetTags(v []*GetResourceConfigurationResponseBodyTags) *GetResourceConfigurationResponseBody
	GetTags() []*GetResourceConfigurationResponseBodyTags
	SetZoneId(v string) *GetResourceConfigurationResponseBody
	GetZoneId() *string
}

type GetResourceConfigurationResponseBody struct {
	// The ID of the Alibaba Cloud account to which the resource belongs.
	//
	// example:
	//
	// 151266687691****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The configurations of the resource.
	Configuration map[string]interface{} `json:"Configuration,omitempty" xml:"Configuration,omitempty"`
	// The time when the resource was created.
	//
	// example:
	//
	// 2021-06-30T09:20:08Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the resource expires.
	//
	// example:
	//
	// 2021-07-30T09:20:08Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The attributes of the IP address.
	IpAddressAttributes []*GetResourceConfigurationResponseBodyIpAddressAttributes `json:"IpAddressAttributes,omitempty" xml:"IpAddressAttributes,omitempty" type:"Repeated"`
	// The IP addresses.
	//
	// > Whether this parameter is returned is determined by the Alibaba Cloud service to which the resource belongs.
	IpAddresses []*string `json:"IpAddresses,omitempty" xml:"IpAddresses,omitempty" type:"Repeated"`
	// The region ID of the resource.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// F1CE0D52-32DA-531A-87A4-B9A5B68D5D8E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the resource group to which the resource belongs.
	//
	// example:
	//
	// rg-acfmv4k****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the resource.
	//
	// example:
	//
	// vtb-uf6978gdqbi****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The name of the resource.
	//
	// example:
	//
	// group1
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// The type of the resource.
	//
	// example:
	//
	// ACS::VPC::VSwitch
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tags of the resource.
	Tags []*GetResourceConfigurationResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The zone ID of the resource.
	//
	// example:
	//
	// cn-hangzhou-k
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s GetResourceConfigurationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetResourceConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *GetResourceConfigurationResponseBody) GetAccountId() *string {
	return s.AccountId
}

func (s *GetResourceConfigurationResponseBody) GetConfiguration() map[string]interface{} {
	return s.Configuration
}

func (s *GetResourceConfigurationResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetResourceConfigurationResponseBody) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *GetResourceConfigurationResponseBody) GetIpAddressAttributes() []*GetResourceConfigurationResponseBodyIpAddressAttributes {
	return s.IpAddressAttributes
}

func (s *GetResourceConfigurationResponseBody) GetIpAddresses() []*string {
	return s.IpAddresses
}

func (s *GetResourceConfigurationResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *GetResourceConfigurationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetResourceConfigurationResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetResourceConfigurationResponseBody) GetResourceId() *string {
	return s.ResourceId
}

func (s *GetResourceConfigurationResponseBody) GetResourceName() *string {
	return s.ResourceName
}

func (s *GetResourceConfigurationResponseBody) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetResourceConfigurationResponseBody) GetTags() []*GetResourceConfigurationResponseBodyTags {
	return s.Tags
}

func (s *GetResourceConfigurationResponseBody) GetZoneId() *string {
	return s.ZoneId
}

func (s *GetResourceConfigurationResponseBody) SetAccountId(v string) *GetResourceConfigurationResponseBody {
	s.AccountId = &v
	return s
}

func (s *GetResourceConfigurationResponseBody) SetConfiguration(v map[string]interface{}) *GetResourceConfigurationResponseBody {
	s.Configuration = v
	return s
}

func (s *GetResourceConfigurationResponseBody) SetCreateTime(v string) *GetResourceConfigurationResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetResourceConfigurationResponseBody) SetExpireTime(v string) *GetResourceConfigurationResponseBody {
	s.ExpireTime = &v
	return s
}

func (s *GetResourceConfigurationResponseBody) SetIpAddressAttributes(v []*GetResourceConfigurationResponseBodyIpAddressAttributes) *GetResourceConfigurationResponseBody {
	s.IpAddressAttributes = v
	return s
}

func (s *GetResourceConfigurationResponseBody) SetIpAddresses(v []*string) *GetResourceConfigurationResponseBody {
	s.IpAddresses = v
	return s
}

func (s *GetResourceConfigurationResponseBody) SetRegionId(v string) *GetResourceConfigurationResponseBody {
	s.RegionId = &v
	return s
}

func (s *GetResourceConfigurationResponseBody) SetRequestId(v string) *GetResourceConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetResourceConfigurationResponseBody) SetResourceGroupId(v string) *GetResourceConfigurationResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *GetResourceConfigurationResponseBody) SetResourceId(v string) *GetResourceConfigurationResponseBody {
	s.ResourceId = &v
	return s
}

func (s *GetResourceConfigurationResponseBody) SetResourceName(v string) *GetResourceConfigurationResponseBody {
	s.ResourceName = &v
	return s
}

func (s *GetResourceConfigurationResponseBody) SetResourceType(v string) *GetResourceConfigurationResponseBody {
	s.ResourceType = &v
	return s
}

func (s *GetResourceConfigurationResponseBody) SetTags(v []*GetResourceConfigurationResponseBodyTags) *GetResourceConfigurationResponseBody {
	s.Tags = v
	return s
}

func (s *GetResourceConfigurationResponseBody) SetZoneId(v string) *GetResourceConfigurationResponseBody {
	s.ZoneId = &v
	return s
}

func (s *GetResourceConfigurationResponseBody) Validate() error {
	if s.IpAddressAttributes != nil {
		for _, item := range s.IpAddressAttributes {
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

type GetResourceConfigurationResponseBodyIpAddressAttributes struct {
	// The IP address.
	//
	// example:
	//
	// 192.168.1.2
	IpAddress *string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty"`
	// The network type. Valid values:
	//
	// 	- **Public**: the Internet
	//
	// 	- **Private**: internal network
	//
	// example:
	//
	// Public
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// The version.
	//
	// example:
	//
	// Ipv4
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetResourceConfigurationResponseBodyIpAddressAttributes) String() string {
	return dara.Prettify(s)
}

func (s GetResourceConfigurationResponseBodyIpAddressAttributes) GoString() string {
	return s.String()
}

func (s *GetResourceConfigurationResponseBodyIpAddressAttributes) GetIpAddress() *string {
	return s.IpAddress
}

func (s *GetResourceConfigurationResponseBodyIpAddressAttributes) GetNetworkType() *string {
	return s.NetworkType
}

func (s *GetResourceConfigurationResponseBodyIpAddressAttributes) GetVersion() *string {
	return s.Version
}

func (s *GetResourceConfigurationResponseBodyIpAddressAttributes) SetIpAddress(v string) *GetResourceConfigurationResponseBodyIpAddressAttributes {
	s.IpAddress = &v
	return s
}

func (s *GetResourceConfigurationResponseBodyIpAddressAttributes) SetNetworkType(v string) *GetResourceConfigurationResponseBodyIpAddressAttributes {
	s.NetworkType = &v
	return s
}

func (s *GetResourceConfigurationResponseBodyIpAddressAttributes) SetVersion(v string) *GetResourceConfigurationResponseBodyIpAddressAttributes {
	s.Version = &v
	return s
}

func (s *GetResourceConfigurationResponseBodyIpAddressAttributes) Validate() error {
	return dara.Validate(s)
}

type GetResourceConfigurationResponseBodyTags struct {
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

func (s GetResourceConfigurationResponseBodyTags) String() string {
	return dara.Prettify(s)
}

func (s GetResourceConfigurationResponseBodyTags) GoString() string {
	return s.String()
}

func (s *GetResourceConfigurationResponseBodyTags) GetKey() *string {
	return s.Key
}

func (s *GetResourceConfigurationResponseBodyTags) GetValue() *string {
	return s.Value
}

func (s *GetResourceConfigurationResponseBodyTags) SetKey(v string) *GetResourceConfigurationResponseBodyTags {
	s.Key = &v
	return s
}

func (s *GetResourceConfigurationResponseBodyTags) SetValue(v string) *GetResourceConfigurationResponseBodyTags {
	s.Value = &v
	return s
}

func (s *GetResourceConfigurationResponseBodyTags) Validate() error {
	return dara.Validate(s)
}
