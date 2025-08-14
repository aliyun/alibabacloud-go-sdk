// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCenBandwidthPackageSpecRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBandwidth(v int32) *ModifyCenBandwidthPackageSpecRequest
	GetBandwidth() *int32
	SetCenBandwidthPackageId(v string) *ModifyCenBandwidthPackageSpecRequest
	GetCenBandwidthPackageId() *string
	SetOwnerAccount(v string) *ModifyCenBandwidthPackageSpecRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyCenBandwidthPackageSpecRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ModifyCenBandwidthPackageSpecRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyCenBandwidthPackageSpecRequest
	GetResourceOwnerId() *int64
}

type ModifyCenBandwidthPackageSpecRequest struct {
	// The new maximum bandwidth value of the bandwidth plan. Unit: Mbit/s.
	//
	// Valid values: **2*	- to **10000**.
	//
	// example:
	//
	// 2
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The ID of the bandwidth plan.
	//
	// This parameter is required.
	//
	// example:
	//
	// cenbwp-4c2zaavbvh5x****
	CenBandwidthPackageId *string `json:"CenBandwidthPackageId,omitempty" xml:"CenBandwidthPackageId,omitempty"`
	OwnerAccount          *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId               *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount  *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId       *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyCenBandwidthPackageSpecRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyCenBandwidthPackageSpecRequest) GoString() string {
	return s.String()
}

func (s *ModifyCenBandwidthPackageSpecRequest) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *ModifyCenBandwidthPackageSpecRequest) GetCenBandwidthPackageId() *string {
	return s.CenBandwidthPackageId
}

func (s *ModifyCenBandwidthPackageSpecRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyCenBandwidthPackageSpecRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyCenBandwidthPackageSpecRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyCenBandwidthPackageSpecRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyCenBandwidthPackageSpecRequest) SetBandwidth(v int32) *ModifyCenBandwidthPackageSpecRequest {
	s.Bandwidth = &v
	return s
}

func (s *ModifyCenBandwidthPackageSpecRequest) SetCenBandwidthPackageId(v string) *ModifyCenBandwidthPackageSpecRequest {
	s.CenBandwidthPackageId = &v
	return s
}

func (s *ModifyCenBandwidthPackageSpecRequest) SetOwnerAccount(v string) *ModifyCenBandwidthPackageSpecRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyCenBandwidthPackageSpecRequest) SetOwnerId(v int64) *ModifyCenBandwidthPackageSpecRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyCenBandwidthPackageSpecRequest) SetResourceOwnerAccount(v string) *ModifyCenBandwidthPackageSpecRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyCenBandwidthPackageSpecRequest) SetResourceOwnerId(v int64) *ModifyCenBandwidthPackageSpecRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyCenBandwidthPackageSpecRequest) Validate() error {
	return dara.Validate(s)
}
