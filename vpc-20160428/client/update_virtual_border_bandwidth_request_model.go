// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVirtualBorderBandwidthRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBandwidth(v int32) *UpdateVirtualBorderBandwidthRequest
	GetBandwidth() *int32
	SetClientToken(v string) *UpdateVirtualBorderBandwidthRequest
	GetClientToken() *string
	SetOwnerAccount(v string) *UpdateVirtualBorderBandwidthRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *UpdateVirtualBorderBandwidthRequest
	GetOwnerId() *int64
	SetRegionId(v string) *UpdateVirtualBorderBandwidthRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *UpdateVirtualBorderBandwidthRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpdateVirtualBorderBandwidthRequest
	GetResourceOwnerId() *int64
	SetVirtualBorderRouterId(v string) *UpdateVirtualBorderBandwidthRequest
	GetVirtualBorderRouterId() *string
}

type UpdateVirtualBorderBandwidthRequest struct {
	// The new maximum bandwidth value for the VBR. Unit: Mbit/s.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-0016****
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the VBR.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the VBR.
	//
	// This parameter is required.
	//
	// example:
	//
	// vbr-bp15zckdt37pq72****
	VirtualBorderRouterId *string `json:"VirtualBorderRouterId,omitempty" xml:"VirtualBorderRouterId,omitempty"`
}

func (s UpdateVirtualBorderBandwidthRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateVirtualBorderBandwidthRequest) GoString() string {
	return s.String()
}

func (s *UpdateVirtualBorderBandwidthRequest) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *UpdateVirtualBorderBandwidthRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateVirtualBorderBandwidthRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *UpdateVirtualBorderBandwidthRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateVirtualBorderBandwidthRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateVirtualBorderBandwidthRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpdateVirtualBorderBandwidthRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpdateVirtualBorderBandwidthRequest) GetVirtualBorderRouterId() *string {
	return s.VirtualBorderRouterId
}

func (s *UpdateVirtualBorderBandwidthRequest) SetBandwidth(v int32) *UpdateVirtualBorderBandwidthRequest {
	s.Bandwidth = &v
	return s
}

func (s *UpdateVirtualBorderBandwidthRequest) SetClientToken(v string) *UpdateVirtualBorderBandwidthRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateVirtualBorderBandwidthRequest) SetOwnerAccount(v string) *UpdateVirtualBorderBandwidthRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UpdateVirtualBorderBandwidthRequest) SetOwnerId(v int64) *UpdateVirtualBorderBandwidthRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateVirtualBorderBandwidthRequest) SetRegionId(v string) *UpdateVirtualBorderBandwidthRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateVirtualBorderBandwidthRequest) SetResourceOwnerAccount(v string) *UpdateVirtualBorderBandwidthRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateVirtualBorderBandwidthRequest) SetResourceOwnerId(v int64) *UpdateVirtualBorderBandwidthRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateVirtualBorderBandwidthRequest) SetVirtualBorderRouterId(v string) *UpdateVirtualBorderBandwidthRequest {
	s.VirtualBorderRouterId = &v
	return s
}

func (s *UpdateVirtualBorderBandwidthRequest) Validate() error {
	return dara.Validate(s)
}
