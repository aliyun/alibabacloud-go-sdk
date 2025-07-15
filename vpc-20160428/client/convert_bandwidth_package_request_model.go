// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConvertBandwidthPackageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBandwidthPackageId(v string) *ConvertBandwidthPackageRequest
	GetBandwidthPackageId() *string
	SetClientToken(v string) *ConvertBandwidthPackageRequest
	GetClientToken() *string
	SetOwnerId(v int64) *ConvertBandwidthPackageRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ConvertBandwidthPackageRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ConvertBandwidthPackageRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ConvertBandwidthPackageRequest
	GetResourceOwnerId() *int64
}

type ConvertBandwidthPackageRequest struct {
	// The ID of the NAT service plan.
	//
	// This parameter is required.
	//
	// example:
	//
	// bwp-bp1xea10o8qxw4f****
	BandwidthPackageId *string `json:"BandwidthPackageId,omitempty" xml:"BandwidthPackageId,omitempty"`
	// The client token that is used to ensure the idempotency of the request. You can use the client to generate a client token. Make sure that a unique client token is used for each request. The **token*	- can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/36569.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region where the NAT gateway resides. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
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

func (s ConvertBandwidthPackageRequest) String() string {
	return dara.Prettify(s)
}

func (s ConvertBandwidthPackageRequest) GoString() string {
	return s.String()
}

func (s *ConvertBandwidthPackageRequest) GetBandwidthPackageId() *string {
	return s.BandwidthPackageId
}

func (s *ConvertBandwidthPackageRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ConvertBandwidthPackageRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ConvertBandwidthPackageRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ConvertBandwidthPackageRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ConvertBandwidthPackageRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ConvertBandwidthPackageRequest) SetBandwidthPackageId(v string) *ConvertBandwidthPackageRequest {
	s.BandwidthPackageId = &v
	return s
}

func (s *ConvertBandwidthPackageRequest) SetClientToken(v string) *ConvertBandwidthPackageRequest {
	s.ClientToken = &v
	return s
}

func (s *ConvertBandwidthPackageRequest) SetOwnerId(v int64) *ConvertBandwidthPackageRequest {
	s.OwnerId = &v
	return s
}

func (s *ConvertBandwidthPackageRequest) SetRegionId(v string) *ConvertBandwidthPackageRequest {
	s.RegionId = &v
	return s
}

func (s *ConvertBandwidthPackageRequest) SetResourceOwnerAccount(v string) *ConvertBandwidthPackageRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ConvertBandwidthPackageRequest) SetResourceOwnerId(v int64) *ConvertBandwidthPackageRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ConvertBandwidthPackageRequest) Validate() error {
	return dara.Validate(s)
}
