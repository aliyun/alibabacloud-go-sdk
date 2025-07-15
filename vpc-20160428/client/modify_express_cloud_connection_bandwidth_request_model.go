// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyExpressCloudConnectionBandwidthRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBandwidth(v string) *ModifyExpressCloudConnectionBandwidthRequest
	GetBandwidth() *string
	SetEccId(v string) *ModifyExpressCloudConnectionBandwidthRequest
	GetEccId() *string
	SetOwnerAccount(v string) *ModifyExpressCloudConnectionBandwidthRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyExpressCloudConnectionBandwidthRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyExpressCloudConnectionBandwidthRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyExpressCloudConnectionBandwidthRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyExpressCloudConnectionBandwidthRequest
	GetResourceOwnerId() *int64
}

type ModifyExpressCloudConnectionBandwidthRequest struct {
	// The bandwidth of the ECC instance.
	//
	// example:
	//
	// 2
	Bandwidth *string `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The ID of the ECC instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecc-xxxxxxxxx
	EccId        *string `json:"EccId,omitempty" xml:"EccId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
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

func (s ModifyExpressCloudConnectionBandwidthRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyExpressCloudConnectionBandwidthRequest) GoString() string {
	return s.String()
}

func (s *ModifyExpressCloudConnectionBandwidthRequest) GetBandwidth() *string {
	return s.Bandwidth
}

func (s *ModifyExpressCloudConnectionBandwidthRequest) GetEccId() *string {
	return s.EccId
}

func (s *ModifyExpressCloudConnectionBandwidthRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyExpressCloudConnectionBandwidthRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyExpressCloudConnectionBandwidthRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyExpressCloudConnectionBandwidthRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyExpressCloudConnectionBandwidthRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyExpressCloudConnectionBandwidthRequest) SetBandwidth(v string) *ModifyExpressCloudConnectionBandwidthRequest {
	s.Bandwidth = &v
	return s
}

func (s *ModifyExpressCloudConnectionBandwidthRequest) SetEccId(v string) *ModifyExpressCloudConnectionBandwidthRequest {
	s.EccId = &v
	return s
}

func (s *ModifyExpressCloudConnectionBandwidthRequest) SetOwnerAccount(v string) *ModifyExpressCloudConnectionBandwidthRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyExpressCloudConnectionBandwidthRequest) SetOwnerId(v int64) *ModifyExpressCloudConnectionBandwidthRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyExpressCloudConnectionBandwidthRequest) SetRegionId(v string) *ModifyExpressCloudConnectionBandwidthRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyExpressCloudConnectionBandwidthRequest) SetResourceOwnerAccount(v string) *ModifyExpressCloudConnectionBandwidthRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyExpressCloudConnectionBandwidthRequest) SetResourceOwnerId(v int64) *ModifyExpressCloudConnectionBandwidthRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyExpressCloudConnectionBandwidthRequest) Validate() error {
	return dara.Validate(s)
}
