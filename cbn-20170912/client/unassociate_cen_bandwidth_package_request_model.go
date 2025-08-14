// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnassociateCenBandwidthPackageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCenBandwidthPackageId(v string) *UnassociateCenBandwidthPackageRequest
	GetCenBandwidthPackageId() *string
	SetCenId(v string) *UnassociateCenBandwidthPackageRequest
	GetCenId() *string
	SetOwnerAccount(v string) *UnassociateCenBandwidthPackageRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *UnassociateCenBandwidthPackageRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *UnassociateCenBandwidthPackageRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UnassociateCenBandwidthPackageRequest
	GetResourceOwnerId() *int64
}

type UnassociateCenBandwidthPackageRequest struct {
	// The ID of the bandwidth plan.
	//
	// This parameter is required.
	//
	// example:
	//
	// cenbwp-4c2zaavbvh5****
	CenBandwidthPackageId *string `json:"CenBandwidthPackageId,omitempty" xml:"CenBandwidthPackageId,omitempty"`
	// The ID of the CEN instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cen-7qthudw0ll6j****
	CenId                *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s UnassociateCenBandwidthPackageRequest) String() string {
	return dara.Prettify(s)
}

func (s UnassociateCenBandwidthPackageRequest) GoString() string {
	return s.String()
}

func (s *UnassociateCenBandwidthPackageRequest) GetCenBandwidthPackageId() *string {
	return s.CenBandwidthPackageId
}

func (s *UnassociateCenBandwidthPackageRequest) GetCenId() *string {
	return s.CenId
}

func (s *UnassociateCenBandwidthPackageRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *UnassociateCenBandwidthPackageRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UnassociateCenBandwidthPackageRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UnassociateCenBandwidthPackageRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UnassociateCenBandwidthPackageRequest) SetCenBandwidthPackageId(v string) *UnassociateCenBandwidthPackageRequest {
	s.CenBandwidthPackageId = &v
	return s
}

func (s *UnassociateCenBandwidthPackageRequest) SetCenId(v string) *UnassociateCenBandwidthPackageRequest {
	s.CenId = &v
	return s
}

func (s *UnassociateCenBandwidthPackageRequest) SetOwnerAccount(v string) *UnassociateCenBandwidthPackageRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UnassociateCenBandwidthPackageRequest) SetOwnerId(v int64) *UnassociateCenBandwidthPackageRequest {
	s.OwnerId = &v
	return s
}

func (s *UnassociateCenBandwidthPackageRequest) SetResourceOwnerAccount(v string) *UnassociateCenBandwidthPackageRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UnassociateCenBandwidthPackageRequest) SetResourceOwnerId(v int64) *UnassociateCenBandwidthPackageRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UnassociateCenBandwidthPackageRequest) Validate() error {
	return dara.Validate(s)
}
