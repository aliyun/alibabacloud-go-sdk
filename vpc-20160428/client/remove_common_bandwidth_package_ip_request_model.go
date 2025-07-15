// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveCommonBandwidthPackageIpRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBandwidthPackageId(v string) *RemoveCommonBandwidthPackageIpRequest
	GetBandwidthPackageId() *string
	SetClientToken(v string) *RemoveCommonBandwidthPackageIpRequest
	GetClientToken() *string
	SetIpInstanceId(v string) *RemoveCommonBandwidthPackageIpRequest
	GetIpInstanceId() *string
	SetOwnerAccount(v string) *RemoveCommonBandwidthPackageIpRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *RemoveCommonBandwidthPackageIpRequest
	GetOwnerId() *int64
	SetRegionId(v string) *RemoveCommonBandwidthPackageIpRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *RemoveCommonBandwidthPackageIpRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *RemoveCommonBandwidthPackageIpRequest
	GetResourceOwnerId() *int64
}

type RemoveCommonBandwidthPackageIpRequest struct {
	// The ID of the Internet Shared Bandwidth instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cbwp-2ze2ic1xd2qeqk145****
	BandwidthPackageId *string `json:"BandwidthPackageId,omitempty" xml:"BandwidthPackageId,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The EIP ID.
	//
	// You can call the [DescribeEipAddresses](https://help.aliyun.com/document_detail/36018.html) operation to query EIP IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// eip-2zeerraiwb7uj6i0d****
	IpInstanceId *string `json:"IpInstanceId,omitempty" xml:"IpInstanceId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the Internet Shared Bandwidth instance.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s RemoveCommonBandwidthPackageIpRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveCommonBandwidthPackageIpRequest) GoString() string {
	return s.String()
}

func (s *RemoveCommonBandwidthPackageIpRequest) GetBandwidthPackageId() *string {
	return s.BandwidthPackageId
}

func (s *RemoveCommonBandwidthPackageIpRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *RemoveCommonBandwidthPackageIpRequest) GetIpInstanceId() *string {
	return s.IpInstanceId
}

func (s *RemoveCommonBandwidthPackageIpRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *RemoveCommonBandwidthPackageIpRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *RemoveCommonBandwidthPackageIpRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RemoveCommonBandwidthPackageIpRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *RemoveCommonBandwidthPackageIpRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *RemoveCommonBandwidthPackageIpRequest) SetBandwidthPackageId(v string) *RemoveCommonBandwidthPackageIpRequest {
	s.BandwidthPackageId = &v
	return s
}

func (s *RemoveCommonBandwidthPackageIpRequest) SetClientToken(v string) *RemoveCommonBandwidthPackageIpRequest {
	s.ClientToken = &v
	return s
}

func (s *RemoveCommonBandwidthPackageIpRequest) SetIpInstanceId(v string) *RemoveCommonBandwidthPackageIpRequest {
	s.IpInstanceId = &v
	return s
}

func (s *RemoveCommonBandwidthPackageIpRequest) SetOwnerAccount(v string) *RemoveCommonBandwidthPackageIpRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RemoveCommonBandwidthPackageIpRequest) SetOwnerId(v int64) *RemoveCommonBandwidthPackageIpRequest {
	s.OwnerId = &v
	return s
}

func (s *RemoveCommonBandwidthPackageIpRequest) SetRegionId(v string) *RemoveCommonBandwidthPackageIpRequest {
	s.RegionId = &v
	return s
}

func (s *RemoveCommonBandwidthPackageIpRequest) SetResourceOwnerAccount(v string) *RemoveCommonBandwidthPackageIpRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RemoveCommonBandwidthPackageIpRequest) SetResourceOwnerId(v int64) *RemoveCommonBandwidthPackageIpRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RemoveCommonBandwidthPackageIpRequest) Validate() error {
	return dara.Validate(s)
}
