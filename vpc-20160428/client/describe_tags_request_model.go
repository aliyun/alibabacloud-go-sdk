// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTagsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResult(v int32) *DescribeTagsRequest
	GetMaxResult() *int32
	SetNextToken(v string) *DescribeTagsRequest
	GetNextToken() *string
	SetOwnerAccount(v string) *DescribeTagsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeTagsRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeTagsRequest
	GetRegionId() *string
	SetResourceId(v []*string) *DescribeTagsRequest
	GetResourceId() []*string
	SetResourceOwnerAccount(v string) *DescribeTagsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeTagsRequest
	GetResourceOwnerId() *int64
	SetResourceType(v string) *DescribeTagsRequest
	GetResourceType() *string
	SetTag(v []*DescribeTagsRequestTag) *DescribeTagsRequest
	GetTag() []*DescribeTagsRequestTag
}

type DescribeTagsRequest struct {
	// The number of entries per page. Valid values: 1 to 50. Default value: 50.
	//
	// example:
	//
	// 50
	MaxResult *int32 `json:"MaxResult,omitempty" xml:"MaxResult,omitempty"`
	// The pagination token. Valid values:
	//
	// - If this is the first query or no subsequent query exists, leave this parameter empty.
	//
	// - If a subsequent query exists, set the value to the NextToken value returned by the previous API call.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the resource.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource ID. You can specify up to 50 resource IDs.
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
	// - **HAVIP**: high-availability virtual IP address instance.
	//
	// - **DHCPOPTIONSSET**: DHCP options set instance.
	//
	// - **GATEWAYENDPOINT**: gateway endpoint instance.
	//
	// > The resource type value is case-insensitive.
	//
	// example:
	//
	// VPC
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tags of the resource.
	Tag []*DescribeTagsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeTagsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTagsRequest) GoString() string {
	return s.String()
}

func (s *DescribeTagsRequest) GetMaxResult() *int32 {
	return s.MaxResult
}

func (s *DescribeTagsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeTagsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeTagsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeTagsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeTagsRequest) GetResourceId() []*string {
	return s.ResourceId
}

func (s *DescribeTagsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeTagsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeTagsRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeTagsRequest) GetTag() []*DescribeTagsRequestTag {
	return s.Tag
}

func (s *DescribeTagsRequest) SetMaxResult(v int32) *DescribeTagsRequest {
	s.MaxResult = &v
	return s
}

func (s *DescribeTagsRequest) SetNextToken(v string) *DescribeTagsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeTagsRequest) SetOwnerAccount(v string) *DescribeTagsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeTagsRequest) SetOwnerId(v int64) *DescribeTagsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeTagsRequest) SetRegionId(v string) *DescribeTagsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeTagsRequest) SetResourceId(v []*string) *DescribeTagsRequest {
	s.ResourceId = v
	return s
}

func (s *DescribeTagsRequest) SetResourceOwnerAccount(v string) *DescribeTagsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeTagsRequest) SetResourceOwnerId(v int64) *DescribeTagsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeTagsRequest) SetResourceType(v string) *DescribeTagsRequest {
	s.ResourceType = &v
	return s
}

func (s *DescribeTagsRequest) SetTag(v []*DescribeTagsRequestTag) *DescribeTagsRequest {
	s.Tag = v
	return s
}

func (s *DescribeTagsRequest) Validate() error {
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

type DescribeTagsRequestTag struct {
	// The tag key of the resource. You can specify up to 20 tag keys.
	//
	// A tag key can be up to 128 characters in length. It cannot be an empty string or start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the resource. You can specify up to 20 tag values.
	//
	// A tag value can be up to 128 characters in length. It can be an empty string but cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceJoshua
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeTagsRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeTagsRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeTagsRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeTagsRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeTagsRequestTag) SetKey(v string) *DescribeTagsRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeTagsRequestTag) SetValue(v string) *DescribeTagsRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeTagsRequestTag) Validate() error {
	return dara.Validate(s)
}
