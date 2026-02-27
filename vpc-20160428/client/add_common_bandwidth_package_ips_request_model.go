// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddCommonBandwidthPackageIpsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBandwidthPackageId(v string) *AddCommonBandwidthPackageIpsRequest
	GetBandwidthPackageId() *string
	SetClientToken(v string) *AddCommonBandwidthPackageIpsRequest
	GetClientToken() *string
	SetIpInstanceIds(v []*string) *AddCommonBandwidthPackageIpsRequest
	GetIpInstanceIds() []*string
	SetIpType(v string) *AddCommonBandwidthPackageIpsRequest
	GetIpType() *string
	SetOwnerAccount(v string) *AddCommonBandwidthPackageIpsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *AddCommonBandwidthPackageIpsRequest
	GetOwnerId() *int64
	SetRegionId(v string) *AddCommonBandwidthPackageIpsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *AddCommonBandwidthPackageIpsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AddCommonBandwidthPackageIpsRequest
	GetResourceOwnerId() *int64
}

type AddCommonBandwidthPackageIpsRequest struct {
	// The ID of the Internet Shared Bandwidth instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cbwp-2ze2ic1xd2qeqasdf****
	BandwidthPackageId *string `json:"BandwidthPackageId,omitempty" xml:"BandwidthPackageId,omitempty"`
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
	// The list of EIPs that you want to associate with the Internet Shared Bandwidth instance.
	//
	// You can specify at most 10 EIP IDs at a time.
	//
	// This parameter is required.
	IpInstanceIds []*string `json:"IpInstanceIds,omitempty" xml:"IpInstanceIds,omitempty" type:"Repeated"`
	// The IP type. Set the value to **EIP**, which indicates that an EIP is added to the Shared Bandwidth.
	//
	// example:
	//
	// EIP
	IpType       *string `json:"IpType,omitempty" xml:"IpType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the Internet Shared Bandwidth instance.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/448570.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s AddCommonBandwidthPackageIpsRequest) String() string {
	return dara.Prettify(s)
}

func (s AddCommonBandwidthPackageIpsRequest) GoString() string {
	return s.String()
}

func (s *AddCommonBandwidthPackageIpsRequest) GetBandwidthPackageId() *string {
	return s.BandwidthPackageId
}

func (s *AddCommonBandwidthPackageIpsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *AddCommonBandwidthPackageIpsRequest) GetIpInstanceIds() []*string {
	return s.IpInstanceIds
}

func (s *AddCommonBandwidthPackageIpsRequest) GetIpType() *string {
	return s.IpType
}

func (s *AddCommonBandwidthPackageIpsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *AddCommonBandwidthPackageIpsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AddCommonBandwidthPackageIpsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AddCommonBandwidthPackageIpsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AddCommonBandwidthPackageIpsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AddCommonBandwidthPackageIpsRequest) SetBandwidthPackageId(v string) *AddCommonBandwidthPackageIpsRequest {
	s.BandwidthPackageId = &v
	return s
}

func (s *AddCommonBandwidthPackageIpsRequest) SetClientToken(v string) *AddCommonBandwidthPackageIpsRequest {
	s.ClientToken = &v
	return s
}

func (s *AddCommonBandwidthPackageIpsRequest) SetIpInstanceIds(v []*string) *AddCommonBandwidthPackageIpsRequest {
	s.IpInstanceIds = v
	return s
}

func (s *AddCommonBandwidthPackageIpsRequest) SetIpType(v string) *AddCommonBandwidthPackageIpsRequest {
	s.IpType = &v
	return s
}

func (s *AddCommonBandwidthPackageIpsRequest) SetOwnerAccount(v string) *AddCommonBandwidthPackageIpsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AddCommonBandwidthPackageIpsRequest) SetOwnerId(v int64) *AddCommonBandwidthPackageIpsRequest {
	s.OwnerId = &v
	return s
}

func (s *AddCommonBandwidthPackageIpsRequest) SetRegionId(v string) *AddCommonBandwidthPackageIpsRequest {
	s.RegionId = &v
	return s
}

func (s *AddCommonBandwidthPackageIpsRequest) SetResourceOwnerAccount(v string) *AddCommonBandwidthPackageIpsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AddCommonBandwidthPackageIpsRequest) SetResourceOwnerId(v int64) *AddCommonBandwidthPackageIpsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AddCommonBandwidthPackageIpsRequest) Validate() error {
	return dara.Validate(s)
}
