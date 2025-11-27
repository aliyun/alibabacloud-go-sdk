// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMarketingActivityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliUid(v int64) *DescribeMarketingActivityRequest
	GetAliUid() *int64
	SetBid(v string) *DescribeMarketingActivityRequest
	GetBid() *string
	SetClientToken(v string) *DescribeMarketingActivityRequest
	GetClientToken() *string
	SetOwnerId(v int64) *DescribeMarketingActivityRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeMarketingActivityRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeMarketingActivityRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeMarketingActivityRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeMarketingActivityRequest
	GetResourceOwnerId() *int64
	SetUpgradeCode(v string) *DescribeMarketingActivityRequest
	GetUpgradeCode() *string
}

type DescribeMarketingActivityRequest struct {
	// The ID of the Alibaba Cloud account.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20725049
	AliUid *int64 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// 	- China site: 26842
	//
	// 	- International site: 26888
	//
	// example:
	//
	// 26842
	Bid *string `json:"Bid,omitempty" xml:"Bid,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCzxxxxxxx
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/26243.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmy*****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The service name.
	//
	// This parameter is required.
	//
	// example:
	//
	// series
	UpgradeCode *string `json:"UpgradeCode,omitempty" xml:"UpgradeCode,omitempty"`
}

func (s DescribeMarketingActivityRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeMarketingActivityRequest) GoString() string {
	return s.String()
}

func (s *DescribeMarketingActivityRequest) GetAliUid() *int64 {
	return s.AliUid
}

func (s *DescribeMarketingActivityRequest) GetBid() *string {
	return s.Bid
}

func (s *DescribeMarketingActivityRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeMarketingActivityRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeMarketingActivityRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeMarketingActivityRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeMarketingActivityRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeMarketingActivityRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeMarketingActivityRequest) GetUpgradeCode() *string {
	return s.UpgradeCode
}

func (s *DescribeMarketingActivityRequest) SetAliUid(v int64) *DescribeMarketingActivityRequest {
	s.AliUid = &v
	return s
}

func (s *DescribeMarketingActivityRequest) SetBid(v string) *DescribeMarketingActivityRequest {
	s.Bid = &v
	return s
}

func (s *DescribeMarketingActivityRequest) SetClientToken(v string) *DescribeMarketingActivityRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeMarketingActivityRequest) SetOwnerId(v int64) *DescribeMarketingActivityRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeMarketingActivityRequest) SetRegionId(v string) *DescribeMarketingActivityRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeMarketingActivityRequest) SetResourceGroupId(v string) *DescribeMarketingActivityRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeMarketingActivityRequest) SetResourceOwnerAccount(v string) *DescribeMarketingActivityRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeMarketingActivityRequest) SetResourceOwnerId(v int64) *DescribeMarketingActivityRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeMarketingActivityRequest) SetUpgradeCode(v string) *DescribeMarketingActivityRequest {
	s.UpgradeCode = &v
	return s
}

func (s *DescribeMarketingActivityRequest) Validate() error {
	return dara.Validate(s)
}
