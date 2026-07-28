// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnTagResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAll(v bool) *UnTagResourcesRequest
	GetAll() *bool
	SetOwnerAccount(v string) *UnTagResourcesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *UnTagResourcesRequest
	GetOwnerId() *int64
	SetRegionId(v string) *UnTagResourcesRequest
	GetRegionId() *string
	SetResourceId(v []*string) *UnTagResourcesRequest
	GetResourceId() []*string
	SetResourceOwnerAccount(v string) *UnTagResourcesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UnTagResourcesRequest
	GetResourceOwnerId() *int64
	SetResourceType(v string) *UnTagResourcesRequest
	GetResourceType() *string
	SetTagKey(v []*string) *UnTagResourcesRequest
	GetTagKey() []*string
}

type UnTagResourcesRequest struct {
	// Specifies whether to unbind all tags from the resources. Valid values:
	//
	// - **true**: Unbinds all tags from the resources.
	//
	// - **false*	- (default): Does not unbind all tags from the resources.
	//
	// example:
	//
	// false
	All          *bool   `json:"All,omitempty" xml:"All,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the resources.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the region ID.
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
	// - **VSWITCH**: vSwitch instance.
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
	// - **HAVIP**: high-availability virtual IP address (HaVip) instance.
	//
	// - **DHCPOPTIONSSET**: DHCP options set instance.
	//
	// - **GATEWAYENDPOINT**: gateway endpoint instance.
	//
	// - **IPV6ADDRESS**: IPv6 address instance.
	//
	// > The resource type value is case-insensitive.
	//
	// This parameter is required.
	//
	// example:
	//
	// VPC
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tag keys to unbind. You can specify up to 20 tag keys.
	//
	// Each tag key can be up to 128 characters in length, can be an empty string, and cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceDept
	TagKey []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
}

func (s UnTagResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s UnTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *UnTagResourcesRequest) GetAll() *bool {
	return s.All
}

func (s *UnTagResourcesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *UnTagResourcesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UnTagResourcesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UnTagResourcesRequest) GetResourceId() []*string {
	return s.ResourceId
}

func (s *UnTagResourcesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UnTagResourcesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UnTagResourcesRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *UnTagResourcesRequest) GetTagKey() []*string {
	return s.TagKey
}

func (s *UnTagResourcesRequest) SetAll(v bool) *UnTagResourcesRequest {
	s.All = &v
	return s
}

func (s *UnTagResourcesRequest) SetOwnerAccount(v string) *UnTagResourcesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UnTagResourcesRequest) SetOwnerId(v int64) *UnTagResourcesRequest {
	s.OwnerId = &v
	return s
}

func (s *UnTagResourcesRequest) SetRegionId(v string) *UnTagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *UnTagResourcesRequest) SetResourceId(v []*string) *UnTagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *UnTagResourcesRequest) SetResourceOwnerAccount(v string) *UnTagResourcesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UnTagResourcesRequest) SetResourceOwnerId(v int64) *UnTagResourcesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UnTagResourcesRequest) SetResourceType(v string) *UnTagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *UnTagResourcesRequest) SetTagKey(v []*string) *UnTagResourcesRequest {
	s.TagKey = v
	return s
}

func (s *UnTagResourcesRequest) Validate() error {
	return dara.Validate(s)
}
