// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIpamsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIpamIds(v []*string) *ListIpamsRequest
	GetIpamIds() []*string
	SetIpamName(v string) *ListIpamsRequest
	GetIpamName() *string
	SetMaxResults(v int64) *ListIpamsRequest
	GetMaxResults() *int64
	SetNextToken(v string) *ListIpamsRequest
	GetNextToken() *string
	SetOwnerAccount(v string) *ListIpamsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ListIpamsRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ListIpamsRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ListIpamsRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *ListIpamsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListIpamsRequest
	GetResourceOwnerId() *int64
	SetTags(v []*ListIpamsRequestTags) *ListIpamsRequest
	GetTags() []*ListIpamsRequestTags
}

type ListIpamsRequest struct {
	// The IDs of IPAMs. Valid values of N: 1 to 100. A maximum of 100 IPAMs can be queried at a time.
	IpamIds []*string `json:"IpamIds,omitempty" xml:"IpamIds,omitempty" type:"Repeated"`
	// The name of the IPAM.
	//
	// It must be 1 to 128 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// test
	IpamName *string `json:"IpamName,omitempty" xml:"IpamName,omitempty"`
	// The number of entries per page. Valid values: **1*	- to **100**. Default value: **10**.
	//
	// example:
	//
	// 10
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- You do not need to specify this parameter for the first request.
	//
	// 	- You must specify the token that is obtained from the previous query as the value of NextToken.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the IPAM instance is hosted. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID of the IPAM.
	//
	// example:
	//
	// rg-aek2sermdd6****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The tag information.
	Tags []*ListIpamsRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListIpamsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListIpamsRequest) GoString() string {
	return s.String()
}

func (s *ListIpamsRequest) GetIpamIds() []*string {
	return s.IpamIds
}

func (s *ListIpamsRequest) GetIpamName() *string {
	return s.IpamName
}

func (s *ListIpamsRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListIpamsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListIpamsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ListIpamsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListIpamsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListIpamsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListIpamsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListIpamsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListIpamsRequest) GetTags() []*ListIpamsRequestTags {
	return s.Tags
}

func (s *ListIpamsRequest) SetIpamIds(v []*string) *ListIpamsRequest {
	s.IpamIds = v
	return s
}

func (s *ListIpamsRequest) SetIpamName(v string) *ListIpamsRequest {
	s.IpamName = &v
	return s
}

func (s *ListIpamsRequest) SetMaxResults(v int64) *ListIpamsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListIpamsRequest) SetNextToken(v string) *ListIpamsRequest {
	s.NextToken = &v
	return s
}

func (s *ListIpamsRequest) SetOwnerAccount(v string) *ListIpamsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListIpamsRequest) SetOwnerId(v int64) *ListIpamsRequest {
	s.OwnerId = &v
	return s
}

func (s *ListIpamsRequest) SetRegionId(v string) *ListIpamsRequest {
	s.RegionId = &v
	return s
}

func (s *ListIpamsRequest) SetResourceGroupId(v string) *ListIpamsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListIpamsRequest) SetResourceOwnerAccount(v string) *ListIpamsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListIpamsRequest) SetResourceOwnerId(v int64) *ListIpamsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListIpamsRequest) SetTags(v []*ListIpamsRequestTags) *ListIpamsRequest {
	s.Tags = v
	return s
}

func (s *ListIpamsRequest) Validate() error {
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

type ListIpamsRequestTags struct {
	// The tag key. You can specify at most 20 tag keys. The tag key cannot be an empty string.
	//
	// The tag key can be up to 64 characters in length and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The tag key must start with a letter but cannot start with `aliyun` or `acs:`. The tag key cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. You can specify at most 20 tag values. The tag value can be an empty string.
	//
	// The tag value can be up to 128 characters in length. It must start with a letter and can contain digits, periods (.), underscores (_), and hyphens (-). It cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceJoshua
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListIpamsRequestTags) String() string {
	return dara.Prettify(s)
}

func (s ListIpamsRequestTags) GoString() string {
	return s.String()
}

func (s *ListIpamsRequestTags) GetKey() *string {
	return s.Key
}

func (s *ListIpamsRequestTags) GetValue() *string {
	return s.Value
}

func (s *ListIpamsRequestTags) SetKey(v string) *ListIpamsRequestTags {
	s.Key = &v
	return s
}

func (s *ListIpamsRequestTags) SetValue(v string) *ListIpamsRequestTags {
	s.Value = &v
	return s
}

func (s *ListIpamsRequestTags) Validate() error {
	return dara.Validate(s)
}
