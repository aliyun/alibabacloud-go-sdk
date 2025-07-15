// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNetworkAclsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DescribeNetworkAclsRequest
	GetClientToken() *string
	SetNetworkAclId(v string) *DescribeNetworkAclsRequest
	GetNetworkAclId() *string
	SetNetworkAclName(v string) *DescribeNetworkAclsRequest
	GetNetworkAclName() *string
	SetOwnerAccount(v string) *DescribeNetworkAclsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeNetworkAclsRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeNetworkAclsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeNetworkAclsRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeNetworkAclsRequest
	GetRegionId() *string
	SetResourceId(v string) *DescribeNetworkAclsRequest
	GetResourceId() *string
	SetResourceOwnerAccount(v string) *DescribeNetworkAclsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeNetworkAclsRequest
	GetResourceOwnerId() *int64
	SetResourceType(v string) *DescribeNetworkAclsRequest
	GetResourceType() *string
	SetTags(v []*DescribeNetworkAclsRequestTags) *DescribeNetworkAclsRequest
	GetTags() []*DescribeNetworkAclsRequestTags
	SetVpcId(v string) *DescribeNetworkAclsRequest
	GetVpcId() *string
}

type DescribeNetworkAclsRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the network ACL.
	//
	// example:
	//
	// nacl-bp1lhl0taikrbgnh****
	NetworkAclId *string `json:"NetworkAclId,omitempty" xml:"NetworkAclId,omitempty"`
	// The name of the network ACL.
	//
	// The name must be 1 to 128 characters in length, and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// acl-1
	NetworkAclName *string `json:"NetworkAclName,omitempty" xml:"NetworkAclName,omitempty"`
	OwnerAccount   *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Maximum value: **50**. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the network ACL.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the associated instance.
	//
	// example:
	//
	// vsw-bp1de348lntdwnhbg****
	ResourceId           *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The type of the associated instance. Set the value to **VSwitch**.
	//
	// This parameter is valid only if **ResourceType*	- and **ResourceId*	- are both specified.
	//
	// example:
	//
	// VSwitch
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tag list.
	Tags []*DescribeNetworkAclsRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The ID of the VPC to which the network ACL belongs.
	//
	// example:
	//
	// vpc-m5ebpc2xh64mqm27e****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeNetworkAclsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkAclsRequest) GoString() string {
	return s.String()
}

func (s *DescribeNetworkAclsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeNetworkAclsRequest) GetNetworkAclId() *string {
	return s.NetworkAclId
}

func (s *DescribeNetworkAclsRequest) GetNetworkAclName() *string {
	return s.NetworkAclName
}

func (s *DescribeNetworkAclsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeNetworkAclsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeNetworkAclsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeNetworkAclsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeNetworkAclsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeNetworkAclsRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *DescribeNetworkAclsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeNetworkAclsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeNetworkAclsRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeNetworkAclsRequest) GetTags() []*DescribeNetworkAclsRequestTags {
	return s.Tags
}

func (s *DescribeNetworkAclsRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeNetworkAclsRequest) SetClientToken(v string) *DescribeNetworkAclsRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeNetworkAclsRequest) SetNetworkAclId(v string) *DescribeNetworkAclsRequest {
	s.NetworkAclId = &v
	return s
}

func (s *DescribeNetworkAclsRequest) SetNetworkAclName(v string) *DescribeNetworkAclsRequest {
	s.NetworkAclName = &v
	return s
}

func (s *DescribeNetworkAclsRequest) SetOwnerAccount(v string) *DescribeNetworkAclsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeNetworkAclsRequest) SetOwnerId(v int64) *DescribeNetworkAclsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeNetworkAclsRequest) SetPageNumber(v int32) *DescribeNetworkAclsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeNetworkAclsRequest) SetPageSize(v int32) *DescribeNetworkAclsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeNetworkAclsRequest) SetRegionId(v string) *DescribeNetworkAclsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeNetworkAclsRequest) SetResourceId(v string) *DescribeNetworkAclsRequest {
	s.ResourceId = &v
	return s
}

func (s *DescribeNetworkAclsRequest) SetResourceOwnerAccount(v string) *DescribeNetworkAclsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeNetworkAclsRequest) SetResourceOwnerId(v int64) *DescribeNetworkAclsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeNetworkAclsRequest) SetResourceType(v string) *DescribeNetworkAclsRequest {
	s.ResourceType = &v
	return s
}

func (s *DescribeNetworkAclsRequest) SetTags(v []*DescribeNetworkAclsRequestTags) *DescribeNetworkAclsRequest {
	s.Tags = v
	return s
}

func (s *DescribeNetworkAclsRequest) SetVpcId(v string) *DescribeNetworkAclsRequest {
	s.VpcId = &v
	return s
}

func (s *DescribeNetworkAclsRequest) Validate() error {
	return dara.Validate(s)
}

type DescribeNetworkAclsRequestTags struct {
	// The key of tag N to add to the resource. You can specify up to 20 tag keys. The tag key cannot be an empty string.
	//
	// The tag key can be up to 128 characters in length. It cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N to add to the resource. You can specify at most 20 tag values. The tag value can be an empty string.
	//
	// The tag value can be up to 128 characters in length and cannot contain `http://` or `https://`. The tag value cannot start with `aliyun` or `acs:`.
	//
	// example:
	//
	// FinanceJoshua
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeNetworkAclsRequestTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkAclsRequestTags) GoString() string {
	return s.String()
}

func (s *DescribeNetworkAclsRequestTags) GetKey() *string {
	return s.Key
}

func (s *DescribeNetworkAclsRequestTags) GetValue() *string {
	return s.Value
}

func (s *DescribeNetworkAclsRequestTags) SetKey(v string) *DescribeNetworkAclsRequestTags {
	s.Key = &v
	return s
}

func (s *DescribeNetworkAclsRequestTags) SetValue(v string) *DescribeNetworkAclsRequestTags {
	s.Value = &v
	return s
}

func (s *DescribeNetworkAclsRequestTags) Validate() error {
	return dara.Validate(s)
}
