// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIpamResourceDiscoveriesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIpamResourceDiscoveryIds(v []*string) *ListIpamResourceDiscoveriesRequest
	GetIpamResourceDiscoveryIds() []*string
	SetIpamResourceDiscoveryName(v string) *ListIpamResourceDiscoveriesRequest
	GetIpamResourceDiscoveryName() *string
	SetIsShared(v bool) *ListIpamResourceDiscoveriesRequest
	GetIsShared() *bool
	SetMaxResults(v int32) *ListIpamResourceDiscoveriesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListIpamResourceDiscoveriesRequest
	GetNextToken() *string
	SetOwnerAccount(v string) *ListIpamResourceDiscoveriesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ListIpamResourceDiscoveriesRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ListIpamResourceDiscoveriesRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ListIpamResourceDiscoveriesRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *ListIpamResourceDiscoveriesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListIpamResourceDiscoveriesRequest
	GetResourceOwnerId() *int64
	SetTags(v []*ListIpamResourceDiscoveriesRequestTags) *ListIpamResourceDiscoveriesRequest
	GetTags() []*ListIpamResourceDiscoveriesRequestTags
	SetType(v string) *ListIpamResourceDiscoveriesRequest
	GetType() *string
}

type ListIpamResourceDiscoveriesRequest struct {
	// The IDs of the resource discovery instances. You can query up to 100 instances at a time.
	IpamResourceDiscoveryIds []*string `json:"IpamResourceDiscoveryIds,omitempty" xml:"IpamResourceDiscoveryIds,omitempty" type:"Repeated"`
	// The name of the resource discovery.
	//
	// The name must be 1 to 128 characters in length and cannot start with http\\:// or https\\://.
	//
	// example:
	//
	// test
	IpamResourceDiscoveryName *string `json:"IpamResourceDiscoveryName,omitempty" xml:"IpamResourceDiscoveryName,omitempty"`
	// Specifies whether the resource discovery is shared.
	//
	// example:
	//
	// true
	IsShared *bool `json:"IsShared,omitempty" xml:"IsShared,omitempty"`
	// The maximum number of entries to return on each page. Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token for the next page of results. Valid values:
	//
	// - If **NextToken*	- is empty, no more results are available.
	//
	// - If a value is returned for **NextToken**, the value is the token that is used for the next query.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to get the region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the resource discovery belongs.
	//
	// example:
	//
	// rg-aek2sermdd6****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The list of tags.
	Tags []*ListIpamResourceDiscoveriesRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The type of the resource discovery.
	//
	// > The following types are supported:
	//
	// >
	//
	// > - system: a default resource discovery created by the system.
	//
	// >
	//
	// > - custom: a custom resource discovery created by a user.
	//
	// example:
	//
	// system
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListIpamResourceDiscoveriesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListIpamResourceDiscoveriesRequest) GoString() string {
	return s.String()
}

func (s *ListIpamResourceDiscoveriesRequest) GetIpamResourceDiscoveryIds() []*string {
	return s.IpamResourceDiscoveryIds
}

func (s *ListIpamResourceDiscoveriesRequest) GetIpamResourceDiscoveryName() *string {
	return s.IpamResourceDiscoveryName
}

func (s *ListIpamResourceDiscoveriesRequest) GetIsShared() *bool {
	return s.IsShared
}

func (s *ListIpamResourceDiscoveriesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListIpamResourceDiscoveriesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListIpamResourceDiscoveriesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ListIpamResourceDiscoveriesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListIpamResourceDiscoveriesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListIpamResourceDiscoveriesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListIpamResourceDiscoveriesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListIpamResourceDiscoveriesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListIpamResourceDiscoveriesRequest) GetTags() []*ListIpamResourceDiscoveriesRequestTags {
	return s.Tags
}

func (s *ListIpamResourceDiscoveriesRequest) GetType() *string {
	return s.Type
}

func (s *ListIpamResourceDiscoveriesRequest) SetIpamResourceDiscoveryIds(v []*string) *ListIpamResourceDiscoveriesRequest {
	s.IpamResourceDiscoveryIds = v
	return s
}

func (s *ListIpamResourceDiscoveriesRequest) SetIpamResourceDiscoveryName(v string) *ListIpamResourceDiscoveriesRequest {
	s.IpamResourceDiscoveryName = &v
	return s
}

func (s *ListIpamResourceDiscoveriesRequest) SetIsShared(v bool) *ListIpamResourceDiscoveriesRequest {
	s.IsShared = &v
	return s
}

func (s *ListIpamResourceDiscoveriesRequest) SetMaxResults(v int32) *ListIpamResourceDiscoveriesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListIpamResourceDiscoveriesRequest) SetNextToken(v string) *ListIpamResourceDiscoveriesRequest {
	s.NextToken = &v
	return s
}

func (s *ListIpamResourceDiscoveriesRequest) SetOwnerAccount(v string) *ListIpamResourceDiscoveriesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListIpamResourceDiscoveriesRequest) SetOwnerId(v int64) *ListIpamResourceDiscoveriesRequest {
	s.OwnerId = &v
	return s
}

func (s *ListIpamResourceDiscoveriesRequest) SetRegionId(v string) *ListIpamResourceDiscoveriesRequest {
	s.RegionId = &v
	return s
}

func (s *ListIpamResourceDiscoveriesRequest) SetResourceGroupId(v string) *ListIpamResourceDiscoveriesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListIpamResourceDiscoveriesRequest) SetResourceOwnerAccount(v string) *ListIpamResourceDiscoveriesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListIpamResourceDiscoveriesRequest) SetResourceOwnerId(v int64) *ListIpamResourceDiscoveriesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListIpamResourceDiscoveriesRequest) SetTags(v []*ListIpamResourceDiscoveriesRequestTags) *ListIpamResourceDiscoveriesRequest {
	s.Tags = v
	return s
}

func (s *ListIpamResourceDiscoveriesRequest) SetType(v string) *ListIpamResourceDiscoveriesRequest {
	s.Type = &v
	return s
}

func (s *ListIpamResourceDiscoveriesRequest) Validate() error {
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

type ListIpamResourceDiscoveriesRequestTags struct {
	// The tag key. You can specify up to 20 tag keys. The key cannot be an empty string.
	//
	// A tag key can be up to 64 characters in length. It must start with a letter or a Chinese character and can contain digits, periods (.), underscores (_), and hyphens (-). The key cannot start with `aliyun` or `acs:` and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. You can specify up to 20 tag values. The value can be an empty string.
	//
	// A tag value can be up to 128 characters in length and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceJoshua
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListIpamResourceDiscoveriesRequestTags) String() string {
	return dara.Prettify(s)
}

func (s ListIpamResourceDiscoveriesRequestTags) GoString() string {
	return s.String()
}

func (s *ListIpamResourceDiscoveriesRequestTags) GetKey() *string {
	return s.Key
}

func (s *ListIpamResourceDiscoveriesRequestTags) GetValue() *string {
	return s.Value
}

func (s *ListIpamResourceDiscoveriesRequestTags) SetKey(v string) *ListIpamResourceDiscoveriesRequestTags {
	s.Key = &v
	return s
}

func (s *ListIpamResourceDiscoveriesRequestTags) SetValue(v string) *ListIpamResourceDiscoveriesRequestTags {
	s.Value = &v
	return s
}

func (s *ListIpamResourceDiscoveriesRequestTags) Validate() error {
	return dara.Validate(s)
}
