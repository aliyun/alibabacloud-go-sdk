// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMultiAccountResourceConfigurationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *GetMultiAccountResourceConfigurationResponseBody
	GetAccountId() *string
	SetConfiguration(v map[string]interface{}) *GetMultiAccountResourceConfigurationResponseBody
	GetConfiguration() map[string]interface{}
	SetCreateTime(v string) *GetMultiAccountResourceConfigurationResponseBody
	GetCreateTime() *string
	SetExpireTime(v string) *GetMultiAccountResourceConfigurationResponseBody
	GetExpireTime() *string
	SetIpAddressAttributes(v []*GetMultiAccountResourceConfigurationResponseBodyIpAddressAttributes) *GetMultiAccountResourceConfigurationResponseBody
	GetIpAddressAttributes() []*GetMultiAccountResourceConfigurationResponseBodyIpAddressAttributes
	SetIpAddresses(v []*string) *GetMultiAccountResourceConfigurationResponseBody
	GetIpAddresses() []*string
	SetRegionId(v string) *GetMultiAccountResourceConfigurationResponseBody
	GetRegionId() *string
	SetRequestId(v string) *GetMultiAccountResourceConfigurationResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *GetMultiAccountResourceConfigurationResponseBody
	GetResourceGroupId() *string
	SetResourceId(v string) *GetMultiAccountResourceConfigurationResponseBody
	GetResourceId() *string
	SetResourceName(v string) *GetMultiAccountResourceConfigurationResponseBody
	GetResourceName() *string
	SetResourceType(v string) *GetMultiAccountResourceConfigurationResponseBody
	GetResourceType() *string
	SetTags(v []*GetMultiAccountResourceConfigurationResponseBodyTags) *GetMultiAccountResourceConfigurationResponseBody
	GetTags() []*GetMultiAccountResourceConfigurationResponseBodyTags
	SetZoneId(v string) *GetMultiAccountResourceConfigurationResponseBody
	GetZoneId() *string
}

type GetMultiAccountResourceConfigurationResponseBody struct {
	// The ID of the management account or member of the resource directory.
	//
	// example:
	//
	// 1619302****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The configurations of the resource.
	Configuration map[string]interface{} `json:"Configuration,omitempty" xml:"Configuration,omitempty"`
	// The time when the resource was created.
	//
	// example:
	//
	// 2023-02-14T03:12:11Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the resource expires.
	//
	// example:
	//
	// 2023-09-18T07:04:21Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The attributes of the IP address.
	IpAddressAttributes []*GetMultiAccountResourceConfigurationResponseBodyIpAddressAttributes `json:"IpAddressAttributes,omitempty" xml:"IpAddressAttributes,omitempty" type:"Repeated"`
	// The IP addresses.
	//
	// > Whether this parameter is returned is determined by the Alibaba Cloud service to which the resource belongs.
	IpAddresses []*string `json:"IpAddresses,omitempty" xml:"IpAddresses,omitempty" type:"Repeated"`
	// The region ID of the resource.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// B2DCC08B-C12A-5705-879C-5A1450016156
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the resource group to which the resource belongs.
	//
	// example:
	//
	// rg-acfmzy6d****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the resource.
	//
	// example:
	//
	// m-eb3hji****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The name of the resource.
	//
	// example:
	//
	// test_resource
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// The type of the resource.
	//
	// example:
	//
	// ACS::VPC::RouteTable
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tags of the resource.
	Tags []*GetMultiAccountResourceConfigurationResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The zone ID of the resource.
	//
	// example:
	//
	// cn-shanghai-a
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s GetMultiAccountResourceConfigurationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMultiAccountResourceConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *GetMultiAccountResourceConfigurationResponseBody) GetAccountId() *string {
	return s.AccountId
}

func (s *GetMultiAccountResourceConfigurationResponseBody) GetConfiguration() map[string]interface{} {
	return s.Configuration
}

func (s *GetMultiAccountResourceConfigurationResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetMultiAccountResourceConfigurationResponseBody) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *GetMultiAccountResourceConfigurationResponseBody) GetIpAddressAttributes() []*GetMultiAccountResourceConfigurationResponseBodyIpAddressAttributes {
	return s.IpAddressAttributes
}

func (s *GetMultiAccountResourceConfigurationResponseBody) GetIpAddresses() []*string {
	return s.IpAddresses
}

func (s *GetMultiAccountResourceConfigurationResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *GetMultiAccountResourceConfigurationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMultiAccountResourceConfigurationResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetMultiAccountResourceConfigurationResponseBody) GetResourceId() *string {
	return s.ResourceId
}

func (s *GetMultiAccountResourceConfigurationResponseBody) GetResourceName() *string {
	return s.ResourceName
}

func (s *GetMultiAccountResourceConfigurationResponseBody) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetMultiAccountResourceConfigurationResponseBody) GetTags() []*GetMultiAccountResourceConfigurationResponseBodyTags {
	return s.Tags
}

func (s *GetMultiAccountResourceConfigurationResponseBody) GetZoneId() *string {
	return s.ZoneId
}

func (s *GetMultiAccountResourceConfigurationResponseBody) SetAccountId(v string) *GetMultiAccountResourceConfigurationResponseBody {
	s.AccountId = &v
	return s
}

func (s *GetMultiAccountResourceConfigurationResponseBody) SetConfiguration(v map[string]interface{}) *GetMultiAccountResourceConfigurationResponseBody {
	s.Configuration = v
	return s
}

func (s *GetMultiAccountResourceConfigurationResponseBody) SetCreateTime(v string) *GetMultiAccountResourceConfigurationResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetMultiAccountResourceConfigurationResponseBody) SetExpireTime(v string) *GetMultiAccountResourceConfigurationResponseBody {
	s.ExpireTime = &v
	return s
}

func (s *GetMultiAccountResourceConfigurationResponseBody) SetIpAddressAttributes(v []*GetMultiAccountResourceConfigurationResponseBodyIpAddressAttributes) *GetMultiAccountResourceConfigurationResponseBody {
	s.IpAddressAttributes = v
	return s
}

func (s *GetMultiAccountResourceConfigurationResponseBody) SetIpAddresses(v []*string) *GetMultiAccountResourceConfigurationResponseBody {
	s.IpAddresses = v
	return s
}

func (s *GetMultiAccountResourceConfigurationResponseBody) SetRegionId(v string) *GetMultiAccountResourceConfigurationResponseBody {
	s.RegionId = &v
	return s
}

func (s *GetMultiAccountResourceConfigurationResponseBody) SetRequestId(v string) *GetMultiAccountResourceConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMultiAccountResourceConfigurationResponseBody) SetResourceGroupId(v string) *GetMultiAccountResourceConfigurationResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *GetMultiAccountResourceConfigurationResponseBody) SetResourceId(v string) *GetMultiAccountResourceConfigurationResponseBody {
	s.ResourceId = &v
	return s
}

func (s *GetMultiAccountResourceConfigurationResponseBody) SetResourceName(v string) *GetMultiAccountResourceConfigurationResponseBody {
	s.ResourceName = &v
	return s
}

func (s *GetMultiAccountResourceConfigurationResponseBody) SetResourceType(v string) *GetMultiAccountResourceConfigurationResponseBody {
	s.ResourceType = &v
	return s
}

func (s *GetMultiAccountResourceConfigurationResponseBody) SetTags(v []*GetMultiAccountResourceConfigurationResponseBodyTags) *GetMultiAccountResourceConfigurationResponseBody {
	s.Tags = v
	return s
}

func (s *GetMultiAccountResourceConfigurationResponseBody) SetZoneId(v string) *GetMultiAccountResourceConfigurationResponseBody {
	s.ZoneId = &v
	return s
}

func (s *GetMultiAccountResourceConfigurationResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetMultiAccountResourceConfigurationResponseBodyIpAddressAttributes struct {
	// The IP address.
	//
	// example:
	//
	// 172.27.199.42
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

func (s GetMultiAccountResourceConfigurationResponseBodyIpAddressAttributes) String() string {
	return dara.Prettify(s)
}

func (s GetMultiAccountResourceConfigurationResponseBodyIpAddressAttributes) GoString() string {
	return s.String()
}

func (s *GetMultiAccountResourceConfigurationResponseBodyIpAddressAttributes) GetIpAddress() *string {
	return s.IpAddress
}

func (s *GetMultiAccountResourceConfigurationResponseBodyIpAddressAttributes) GetNetworkType() *string {
	return s.NetworkType
}

func (s *GetMultiAccountResourceConfigurationResponseBodyIpAddressAttributes) GetVersion() *string {
	return s.Version
}

func (s *GetMultiAccountResourceConfigurationResponseBodyIpAddressAttributes) SetIpAddress(v string) *GetMultiAccountResourceConfigurationResponseBodyIpAddressAttributes {
	s.IpAddress = &v
	return s
}

func (s *GetMultiAccountResourceConfigurationResponseBodyIpAddressAttributes) SetNetworkType(v string) *GetMultiAccountResourceConfigurationResponseBodyIpAddressAttributes {
	s.NetworkType = &v
	return s
}

func (s *GetMultiAccountResourceConfigurationResponseBodyIpAddressAttributes) SetVersion(v string) *GetMultiAccountResourceConfigurationResponseBodyIpAddressAttributes {
	s.Version = &v
	return s
}

func (s *GetMultiAccountResourceConfigurationResponseBodyIpAddressAttributes) Validate() error {
	return dara.Validate(s)
}

type GetMultiAccountResourceConfigurationResponseBodyTags struct {
	// The key of tag N.
	//
	// example:
	//
	// tag-key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N.
	//
	// example:
	//
	// tag-value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetMultiAccountResourceConfigurationResponseBodyTags) String() string {
	return dara.Prettify(s)
}

func (s GetMultiAccountResourceConfigurationResponseBodyTags) GoString() string {
	return s.String()
}

func (s *GetMultiAccountResourceConfigurationResponseBodyTags) GetKey() *string {
	return s.Key
}

func (s *GetMultiAccountResourceConfigurationResponseBodyTags) GetValue() *string {
	return s.Value
}

func (s *GetMultiAccountResourceConfigurationResponseBodyTags) SetKey(v string) *GetMultiAccountResourceConfigurationResponseBodyTags {
	s.Key = &v
	return s
}

func (s *GetMultiAccountResourceConfigurationResponseBodyTags) SetValue(v string) *GetMultiAccountResourceConfigurationResponseBodyTags {
	s.Value = &v
	return s
}

func (s *GetMultiAccountResourceConfigurationResponseBodyTags) Validate() error {
	return dara.Validate(s)
}
