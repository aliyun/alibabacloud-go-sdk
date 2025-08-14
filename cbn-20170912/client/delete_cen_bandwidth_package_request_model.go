// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCenBandwidthPackageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCenBandwidthPackageId(v string) *DeleteCenBandwidthPackageRequest
	GetCenBandwidthPackageId() *string
	SetOwnerAccount(v string) *DeleteCenBandwidthPackageRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteCenBandwidthPackageRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DeleteCenBandwidthPackageRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteCenBandwidthPackageRequest
	GetResourceOwnerId() *int64
}

type DeleteCenBandwidthPackageRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cenbwp-4c2zaavbvh5f42****
	CenBandwidthPackageId *string `json:"CenBandwidthPackageId,omitempty" xml:"CenBandwidthPackageId,omitempty"`
	OwnerAccount          *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId               *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount  *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId       *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteCenBandwidthPackageRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCenBandwidthPackageRequest) GoString() string {
	return s.String()
}

func (s *DeleteCenBandwidthPackageRequest) GetCenBandwidthPackageId() *string {
	return s.CenBandwidthPackageId
}

func (s *DeleteCenBandwidthPackageRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteCenBandwidthPackageRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteCenBandwidthPackageRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteCenBandwidthPackageRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteCenBandwidthPackageRequest) SetCenBandwidthPackageId(v string) *DeleteCenBandwidthPackageRequest {
	s.CenBandwidthPackageId = &v
	return s
}

func (s *DeleteCenBandwidthPackageRequest) SetOwnerAccount(v string) *DeleteCenBandwidthPackageRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteCenBandwidthPackageRequest) SetOwnerId(v int64) *DeleteCenBandwidthPackageRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteCenBandwidthPackageRequest) SetResourceOwnerAccount(v string) *DeleteCenBandwidthPackageRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteCenBandwidthPackageRequest) SetResourceOwnerId(v int64) *DeleteCenBandwidthPackageRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteCenBandwidthPackageRequest) Validate() error {
	return dara.Validate(s)
}
