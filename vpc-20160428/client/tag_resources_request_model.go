// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTagResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *TagResourcesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *TagResourcesRequest
	GetOwnerId() *int64
	SetRegionId(v string) *TagResourcesRequest
	GetRegionId() *string
	SetResourceId(v []*string) *TagResourcesRequest
	GetResourceId() []*string
	SetResourceOwnerAccount(v string) *TagResourcesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *TagResourcesRequest
	GetResourceOwnerId() *int64
	SetResourceType(v string) *TagResourcesRequest
	GetResourceType() *string
	SetTag(v []*TagResourcesRequestTag) *TagResourcesRequest
	GetTag() []*TagResourcesRequestTag
}

type TagResourcesRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the resource to which you want to create and bind tags.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource IDs. You can specify up to 50 resource IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-bp16qjewdsunr41m1****
	ResourceId           []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The resource type. Valid values:
	//
	// - **VPC**: virtual private cloud (VPC) instance.
	//
	// - **VSWITCH**: virtual switch instance.
	//
	// - **ROUTETABLE**: route table instance.
	//
	// - **EIP**: elastic IP address (EIP) instance.
	//
	// - **VPNGATEWAY**: VPN gateway instance.
	//
	// - **NATGATEWAY**: NAT gateway instance.
	//
	// - **COMMONBANDWIDTHPACKAGE**: Internet Shared Bandwidth instance.
	//
	// - **PREFIXLIST**: prefix list instance.
	//
	// - **PUBLICIPADDRESSPOOL**: IP address pool instance.
	//
	// - **IPV4GATEWAY**: IPv4 gateway instance.
	//
	// - **IPV6GATEWAY**: IPv6 gateway instance.
	//
	// - **NETWORKACL**: network ACL instance.
	//
	// - **TRAFFICMIRRORFILTER**: traffic mirror filter instance.
	//
	// - **TRAFFICMIRRORSESSION**: traffic mirror session instance.
	//
	// - **FLOWLOG**: flow log instance.
	//
	// - **HAVIP**: high-availability (HA) virtual IP address instance.
	//
	// - **DHCPOPTIONSSET**: DHCP options set instance.
	//
	// - **GATEWAYENDPOINT**: gateway endpoint instance.
	//
	// - **IPV6ADDRESS**: IPv6 address instance.
	//
	//
	// > The resource type value is case-insensitive.
	//
	// This parameter is required.
	//
	// example:
	//
	// VPC
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tag information.
	//
	// This parameter is required.
	//
	// example:
	//
	// TagResources
	Tag []*TagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s TagResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s TagResourcesRequest) GoString() string {
	return s.String()
}

func (s *TagResourcesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *TagResourcesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *TagResourcesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *TagResourcesRequest) GetResourceId() []*string {
	return s.ResourceId
}

func (s *TagResourcesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *TagResourcesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *TagResourcesRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *TagResourcesRequest) GetTag() []*TagResourcesRequestTag {
	return s.Tag
}

func (s *TagResourcesRequest) SetOwnerAccount(v string) *TagResourcesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *TagResourcesRequest) SetOwnerId(v int64) *TagResourcesRequest {
	s.OwnerId = &v
	return s
}

func (s *TagResourcesRequest) SetRegionId(v string) *TagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *TagResourcesRequest) SetResourceId(v []*string) *TagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *TagResourcesRequest) SetResourceOwnerAccount(v string) *TagResourcesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *TagResourcesRequest) SetResourceOwnerId(v int64) *TagResourcesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *TagResourcesRequest) SetResourceType(v string) *TagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *TagResourcesRequest) SetTag(v []*TagResourcesRequestTag) *TagResourcesRequest {
	s.Tag = v
	return s
}

func (s *TagResourcesRequest) Validate() error {
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

type TagResourcesRequestTag struct {
	// The tag key of the resource. You must specify at least 1 and can specify up to 20 tag keys.
	//
	// The tag key can be up to 128 characters in length, and cannot start with `aliyun` or `acs:`, or contain `http://` or `https://`.
	//
	// > The **Tag.N.Key*	- parameter is required when you call this operation, and cannot be an empty string.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the resource. You must specify at least 1 and can specify up to 20 tag values.
	//
	// The tag value can be up to 128 characters in length, and cannot start with `aliyun` or `acs:`, or contain `http://` or `https://`.
	//
	// > The **Tag.N.Value*	- parameter is required when you call this operation, and can be an empty string.
	//
	// example:
	//
	// FinanceJoshua
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s TagResourcesRequestTag) String() string {
	return dara.Prettify(s)
}

func (s TagResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *TagResourcesRequestTag) GetKey() *string {
	return s.Key
}

func (s *TagResourcesRequestTag) GetValue() *string {
	return s.Value
}

func (s *TagResourcesRequestTag) SetKey(v string) *TagResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *TagResourcesRequestTag) SetValue(v string) *TagResourcesRequestTag {
	s.Value = &v
	return s
}

func (s *TagResourcesRequestTag) Validate() error {
	return dara.Validate(s)
}
