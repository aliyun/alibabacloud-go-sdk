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
	AccountId           *string                                                    `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	Configuration       map[string]interface{}                                     `json:"Configuration,omitempty" xml:"Configuration,omitempty"`
	CreateTime          *string                                                    `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ExpireTime          *string                                                    `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	IpAddressAttributes []*GetResourceConfigurationResponseBodyIpAddressAttributes `json:"IpAddressAttributes,omitempty" xml:"IpAddressAttributes,omitempty" type:"Repeated"`
	IpAddresses         []*string                                                  `json:"IpAddresses,omitempty" xml:"IpAddresses,omitempty" type:"Repeated"`
	RegionId            *string                                                    `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RequestId           *string                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceGroupId     *string                                                    `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceId          *string                                                    `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceName        *string                                                    `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	ResourceType        *string                                                    `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Tags                []*GetResourceConfigurationResponseBodyTags                `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	ZoneId              *string                                                    `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
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
	IpAddress   *string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty"`
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	Version     *string `json:"Version,omitempty" xml:"Version,omitempty"`
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
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
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
