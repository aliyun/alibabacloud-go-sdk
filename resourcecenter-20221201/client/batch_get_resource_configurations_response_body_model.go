// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchGetResourceConfigurationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *BatchGetResourceConfigurationsResponseBody
	GetRequestId() *string
	SetResources(v []*BatchGetResourceConfigurationsResponseBodyResources) *BatchGetResourceConfigurationsResponseBody
	GetResources() []*BatchGetResourceConfigurationsResponseBodyResources
}

type BatchGetResourceConfigurationsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// F1CE0D52-32DA-531A-87A4-B9A5B68*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of resources.
	Resources []*BatchGetResourceConfigurationsResponseBodyResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
}

func (s BatchGetResourceConfigurationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchGetResourceConfigurationsResponseBody) GoString() string {
	return s.String()
}

func (s *BatchGetResourceConfigurationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchGetResourceConfigurationsResponseBody) GetResources() []*BatchGetResourceConfigurationsResponseBodyResources {
	return s.Resources
}

func (s *BatchGetResourceConfigurationsResponseBody) SetRequestId(v string) *BatchGetResourceConfigurationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchGetResourceConfigurationsResponseBody) SetResources(v []*BatchGetResourceConfigurationsResponseBodyResources) *BatchGetResourceConfigurationsResponseBody {
	s.Resources = v
	return s
}

func (s *BatchGetResourceConfigurationsResponseBody) Validate() error {
	if s.Resources != nil {
		for _, item := range s.Resources {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type BatchGetResourceConfigurationsResponseBodyResources struct {
	// The ID of the Alibaba Cloud account.
	//
	// example:
	//
	// 151266687691****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The details of the resource configuration.
	Configuration map[string]interface{} `json:"Configuration,omitempty" xml:"Configuration,omitempty"`
	// The time when the resource was created.
	//
	// example:
	//
	// 2021-06-30T09:20:08Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The expiration time of the resource.
	//
	// example:
	//
	// 2021-07-30T09:20:08Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The properties of the IP addresses.
	IpAddressAttributes []*BatchGetResourceConfigurationsResponseBodyResourcesIpAddressAttributes `json:"IpAddressAttributes,omitempty" xml:"IpAddressAttributes,omitempty" type:"Repeated"`
	// The IP addresses.
	//
	// > Whether this parameter is returned depends on the Alibaba Cloud service to which the resource belongs.
	IpAddresses []*string `json:"IpAddresses,omitempty" xml:"IpAddresses,omitempty" type:"Repeated"`
	// The region ID of the resource.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmzawhxxc****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The resource ID.
	//
	// example:
	//
	// eip-wz9gdtce0q6h48h*****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The resource name.
	//
	// example:
	//
	// group1
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// The resource type.
	//
	// example:
	//
	// ACS::ECS::Instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tags.
	Tags []*BatchGetResourceConfigurationsResponseBodyResourcesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The zone ID.
	//
	// example:
	//
	// cn-hangzhou-k
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s BatchGetResourceConfigurationsResponseBodyResources) String() string {
	return dara.Prettify(s)
}

func (s BatchGetResourceConfigurationsResponseBodyResources) GoString() string {
	return s.String()
}

func (s *BatchGetResourceConfigurationsResponseBodyResources) GetAccountId() *string {
	return s.AccountId
}

func (s *BatchGetResourceConfigurationsResponseBodyResources) GetConfiguration() map[string]interface{} {
	return s.Configuration
}

func (s *BatchGetResourceConfigurationsResponseBodyResources) GetCreateTime() *string {
	return s.CreateTime
}

func (s *BatchGetResourceConfigurationsResponseBodyResources) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *BatchGetResourceConfigurationsResponseBodyResources) GetIpAddressAttributes() []*BatchGetResourceConfigurationsResponseBodyResourcesIpAddressAttributes {
	return s.IpAddressAttributes
}

func (s *BatchGetResourceConfigurationsResponseBodyResources) GetIpAddresses() []*string {
	return s.IpAddresses
}

func (s *BatchGetResourceConfigurationsResponseBodyResources) GetRegionId() *string {
	return s.RegionId
}

func (s *BatchGetResourceConfigurationsResponseBodyResources) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *BatchGetResourceConfigurationsResponseBodyResources) GetResourceId() *string {
	return s.ResourceId
}

func (s *BatchGetResourceConfigurationsResponseBodyResources) GetResourceName() *string {
	return s.ResourceName
}

func (s *BatchGetResourceConfigurationsResponseBodyResources) GetResourceType() *string {
	return s.ResourceType
}

func (s *BatchGetResourceConfigurationsResponseBodyResources) GetTags() []*BatchGetResourceConfigurationsResponseBodyResourcesTags {
	return s.Tags
}

func (s *BatchGetResourceConfigurationsResponseBodyResources) GetZoneId() *string {
	return s.ZoneId
}

func (s *BatchGetResourceConfigurationsResponseBodyResources) SetAccountId(v string) *BatchGetResourceConfigurationsResponseBodyResources {
	s.AccountId = &v
	return s
}

func (s *BatchGetResourceConfigurationsResponseBodyResources) SetConfiguration(v map[string]interface{}) *BatchGetResourceConfigurationsResponseBodyResources {
	s.Configuration = v
	return s
}

func (s *BatchGetResourceConfigurationsResponseBodyResources) SetCreateTime(v string) *BatchGetResourceConfigurationsResponseBodyResources {
	s.CreateTime = &v
	return s
}

func (s *BatchGetResourceConfigurationsResponseBodyResources) SetExpireTime(v string) *BatchGetResourceConfigurationsResponseBodyResources {
	s.ExpireTime = &v
	return s
}

func (s *BatchGetResourceConfigurationsResponseBodyResources) SetIpAddressAttributes(v []*BatchGetResourceConfigurationsResponseBodyResourcesIpAddressAttributes) *BatchGetResourceConfigurationsResponseBodyResources {
	s.IpAddressAttributes = v
	return s
}

func (s *BatchGetResourceConfigurationsResponseBodyResources) SetIpAddresses(v []*string) *BatchGetResourceConfigurationsResponseBodyResources {
	s.IpAddresses = v
	return s
}

func (s *BatchGetResourceConfigurationsResponseBodyResources) SetRegionId(v string) *BatchGetResourceConfigurationsResponseBodyResources {
	s.RegionId = &v
	return s
}

func (s *BatchGetResourceConfigurationsResponseBodyResources) SetResourceGroupId(v string) *BatchGetResourceConfigurationsResponseBodyResources {
	s.ResourceGroupId = &v
	return s
}

func (s *BatchGetResourceConfigurationsResponseBodyResources) SetResourceId(v string) *BatchGetResourceConfigurationsResponseBodyResources {
	s.ResourceId = &v
	return s
}

func (s *BatchGetResourceConfigurationsResponseBodyResources) SetResourceName(v string) *BatchGetResourceConfigurationsResponseBodyResources {
	s.ResourceName = &v
	return s
}

func (s *BatchGetResourceConfigurationsResponseBodyResources) SetResourceType(v string) *BatchGetResourceConfigurationsResponseBodyResources {
	s.ResourceType = &v
	return s
}

func (s *BatchGetResourceConfigurationsResponseBodyResources) SetTags(v []*BatchGetResourceConfigurationsResponseBodyResourcesTags) *BatchGetResourceConfigurationsResponseBodyResources {
	s.Tags = v
	return s
}

func (s *BatchGetResourceConfigurationsResponseBodyResources) SetZoneId(v string) *BatchGetResourceConfigurationsResponseBodyResources {
	s.ZoneId = &v
	return s
}

func (s *BatchGetResourceConfigurationsResponseBodyResources) Validate() error {
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

type BatchGetResourceConfigurationsResponseBodyResourcesIpAddressAttributes struct {
	// The IP address.
	//
	// example:
	//
	// 192.168.1.2
	IpAddress *string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty"`
	// The network type. Valid values:
	//
	// - **Public**: Internet.
	//
	// - **Private**: Private network.
	//
	// example:
	//
	// Public
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// The version of the IP address.
	//
	// example:
	//
	// Ipv4
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s BatchGetResourceConfigurationsResponseBodyResourcesIpAddressAttributes) String() string {
	return dara.Prettify(s)
}

func (s BatchGetResourceConfigurationsResponseBodyResourcesIpAddressAttributes) GoString() string {
	return s.String()
}

func (s *BatchGetResourceConfigurationsResponseBodyResourcesIpAddressAttributes) GetIpAddress() *string {
	return s.IpAddress
}

func (s *BatchGetResourceConfigurationsResponseBodyResourcesIpAddressAttributes) GetNetworkType() *string {
	return s.NetworkType
}

func (s *BatchGetResourceConfigurationsResponseBodyResourcesIpAddressAttributes) GetVersion() *string {
	return s.Version
}

func (s *BatchGetResourceConfigurationsResponseBodyResourcesIpAddressAttributes) SetIpAddress(v string) *BatchGetResourceConfigurationsResponseBodyResourcesIpAddressAttributes {
	s.IpAddress = &v
	return s
}

func (s *BatchGetResourceConfigurationsResponseBodyResourcesIpAddressAttributes) SetNetworkType(v string) *BatchGetResourceConfigurationsResponseBodyResourcesIpAddressAttributes {
	s.NetworkType = &v
	return s
}

func (s *BatchGetResourceConfigurationsResponseBodyResourcesIpAddressAttributes) SetVersion(v string) *BatchGetResourceConfigurationsResponseBodyResourcesIpAddressAttributes {
	s.Version = &v
	return s
}

func (s *BatchGetResourceConfigurationsResponseBodyResourcesIpAddressAttributes) Validate() error {
	return dara.Validate(s)
}

type BatchGetResourceConfigurationsResponseBodyResourcesTags struct {
	// The tag key.
	//
	// example:
	//
	// test_key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// test_value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s BatchGetResourceConfigurationsResponseBodyResourcesTags) String() string {
	return dara.Prettify(s)
}

func (s BatchGetResourceConfigurationsResponseBodyResourcesTags) GoString() string {
	return s.String()
}

func (s *BatchGetResourceConfigurationsResponseBodyResourcesTags) GetKey() *string {
	return s.Key
}

func (s *BatchGetResourceConfigurationsResponseBodyResourcesTags) GetValue() *string {
	return s.Value
}

func (s *BatchGetResourceConfigurationsResponseBodyResourcesTags) SetKey(v string) *BatchGetResourceConfigurationsResponseBodyResourcesTags {
	s.Key = &v
	return s
}

func (s *BatchGetResourceConfigurationsResponseBodyResourcesTags) SetValue(v string) *BatchGetResourceConfigurationsResponseBodyResourcesTags {
	s.Value = &v
	return s
}

func (s *BatchGetResourceConfigurationsResponseBodyResourcesTags) Validate() error {
	return dara.Validate(s)
}
