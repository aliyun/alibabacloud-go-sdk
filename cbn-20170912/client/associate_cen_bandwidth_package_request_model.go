// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateCenBandwidthPackageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCenBandwidthPackageId(v string) *AssociateCenBandwidthPackageRequest
	GetCenBandwidthPackageId() *string
	SetCenId(v string) *AssociateCenBandwidthPackageRequest
	GetCenId() *string
	SetOwnerAccount(v string) *AssociateCenBandwidthPackageRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *AssociateCenBandwidthPackageRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *AssociateCenBandwidthPackageRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AssociateCenBandwidthPackageRequest
	GetResourceOwnerId() *int64
}

type AssociateCenBandwidthPackageRequest struct {
	// The ID of the bandwidth plan.
	//
	// This parameter is required.
	//
	// example:
	//
	// cenbwp-4c2zaavbvh5fx****
	CenBandwidthPackageId *string `json:"CenBandwidthPackageId,omitempty" xml:"CenBandwidthPackageId,omitempty"`
	// The ID of the CEN instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cen-7qthudw0ll6jmc****
	CenId                *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s AssociateCenBandwidthPackageRequest) String() string {
	return dara.Prettify(s)
}

func (s AssociateCenBandwidthPackageRequest) GoString() string {
	return s.String()
}

func (s *AssociateCenBandwidthPackageRequest) GetCenBandwidthPackageId() *string {
	return s.CenBandwidthPackageId
}

func (s *AssociateCenBandwidthPackageRequest) GetCenId() *string {
	return s.CenId
}

func (s *AssociateCenBandwidthPackageRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *AssociateCenBandwidthPackageRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AssociateCenBandwidthPackageRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AssociateCenBandwidthPackageRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AssociateCenBandwidthPackageRequest) SetCenBandwidthPackageId(v string) *AssociateCenBandwidthPackageRequest {
	s.CenBandwidthPackageId = &v
	return s
}

func (s *AssociateCenBandwidthPackageRequest) SetCenId(v string) *AssociateCenBandwidthPackageRequest {
	s.CenId = &v
	return s
}

func (s *AssociateCenBandwidthPackageRequest) SetOwnerAccount(v string) *AssociateCenBandwidthPackageRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AssociateCenBandwidthPackageRequest) SetOwnerId(v int64) *AssociateCenBandwidthPackageRequest {
	s.OwnerId = &v
	return s
}

func (s *AssociateCenBandwidthPackageRequest) SetResourceOwnerAccount(v string) *AssociateCenBandwidthPackageRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AssociateCenBandwidthPackageRequest) SetResourceOwnerId(v int64) *AssociateCenBandwidthPackageRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AssociateCenBandwidthPackageRequest) Validate() error {
	return dara.Validate(s)
}
