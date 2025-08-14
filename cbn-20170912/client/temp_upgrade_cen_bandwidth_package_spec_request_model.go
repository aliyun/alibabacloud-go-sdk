// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTempUpgradeCenBandwidthPackageSpecRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBandwidth(v int32) *TempUpgradeCenBandwidthPackageSpecRequest
	GetBandwidth() *int32
	SetCenBandwidthPackageId(v string) *TempUpgradeCenBandwidthPackageSpecRequest
	GetCenBandwidthPackageId() *string
	SetEndTime(v string) *TempUpgradeCenBandwidthPackageSpecRequest
	GetEndTime() *string
	SetOwnerAccount(v string) *TempUpgradeCenBandwidthPackageSpecRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *TempUpgradeCenBandwidthPackageSpecRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *TempUpgradeCenBandwidthPackageSpecRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *TempUpgradeCenBandwidthPackageSpecRequest
	GetResourceOwnerId() *int64
}

type TempUpgradeCenBandwidthPackageSpecRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 20
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cenbwp-y08yosedeqlpua****
	CenBandwidthPackageId *string `json:"CenBandwidthPackageId,omitempty" xml:"CenBandwidthPackageId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2022-07-24T13:00:52Z
	EndTime              *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s TempUpgradeCenBandwidthPackageSpecRequest) String() string {
	return dara.Prettify(s)
}

func (s TempUpgradeCenBandwidthPackageSpecRequest) GoString() string {
	return s.String()
}

func (s *TempUpgradeCenBandwidthPackageSpecRequest) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *TempUpgradeCenBandwidthPackageSpecRequest) GetCenBandwidthPackageId() *string {
	return s.CenBandwidthPackageId
}

func (s *TempUpgradeCenBandwidthPackageSpecRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *TempUpgradeCenBandwidthPackageSpecRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *TempUpgradeCenBandwidthPackageSpecRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *TempUpgradeCenBandwidthPackageSpecRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *TempUpgradeCenBandwidthPackageSpecRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *TempUpgradeCenBandwidthPackageSpecRequest) SetBandwidth(v int32) *TempUpgradeCenBandwidthPackageSpecRequest {
	s.Bandwidth = &v
	return s
}

func (s *TempUpgradeCenBandwidthPackageSpecRequest) SetCenBandwidthPackageId(v string) *TempUpgradeCenBandwidthPackageSpecRequest {
	s.CenBandwidthPackageId = &v
	return s
}

func (s *TempUpgradeCenBandwidthPackageSpecRequest) SetEndTime(v string) *TempUpgradeCenBandwidthPackageSpecRequest {
	s.EndTime = &v
	return s
}

func (s *TempUpgradeCenBandwidthPackageSpecRequest) SetOwnerAccount(v string) *TempUpgradeCenBandwidthPackageSpecRequest {
	s.OwnerAccount = &v
	return s
}

func (s *TempUpgradeCenBandwidthPackageSpecRequest) SetOwnerId(v int64) *TempUpgradeCenBandwidthPackageSpecRequest {
	s.OwnerId = &v
	return s
}

func (s *TempUpgradeCenBandwidthPackageSpecRequest) SetResourceOwnerAccount(v string) *TempUpgradeCenBandwidthPackageSpecRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *TempUpgradeCenBandwidthPackageSpecRequest) SetResourceOwnerId(v int64) *TempUpgradeCenBandwidthPackageSpecRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *TempUpgradeCenBandwidthPackageSpecRequest) Validate() error {
	return dara.Validate(s)
}
