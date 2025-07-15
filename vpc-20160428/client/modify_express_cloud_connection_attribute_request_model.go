// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyExpressCloudConnectionAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBgpAs(v string) *ModifyExpressCloudConnectionAttributeRequest
	GetBgpAs() *string
	SetCeIp(v string) *ModifyExpressCloudConnectionAttributeRequest
	GetCeIp() *string
	SetDescription(v string) *ModifyExpressCloudConnectionAttributeRequest
	GetDescription() *string
	SetEccId(v string) *ModifyExpressCloudConnectionAttributeRequest
	GetEccId() *string
	SetName(v string) *ModifyExpressCloudConnectionAttributeRequest
	GetName() *string
	SetOwnerAccount(v string) *ModifyExpressCloudConnectionAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyExpressCloudConnectionAttributeRequest
	GetOwnerId() *int64
	SetPeIp(v string) *ModifyExpressCloudConnectionAttributeRequest
	GetPeIp() *string
	SetRegionId(v string) *ModifyExpressCloudConnectionAttributeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyExpressCloudConnectionAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyExpressCloudConnectionAttributeRequest
	GetResourceOwnerId() *int64
}

type ModifyExpressCloudConnectionAttributeRequest struct {
	// The BGP autonomous system number (ASN) to be configured for the Smart Access Gateway (SAG) device.
	//
	// example:
	//
	// sag-ejfge***
	BgpAs *string `json:"BgpAs,omitempty" xml:"BgpAs,omitempty"`
	// The peer IP address when the SAG device is connected to the cloud.
	//
	// example:
	//
	// ``172.16.**.**``
	CeIp *string `json:"CeIp,omitempty" xml:"CeIp,omitempty"`
	// Descriptions of ECC.
	//
	// example:
	//
	// ECC
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the ECC instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecc-bp1t9osmuln*******
	EccId *string `json:"EccId,omitempty" xml:"EccId,omitempty"`
	// The name of the ECC instance.
	//
	// example:
	//
	// doctest
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The on-premises IP address when the SAG device is connected to the cloud.
	//
	// example:
	//
	// ``10.10.**.**``
	PeIp *string `json:"PeIp,omitempty" xml:"PeIp,omitempty"`
	// The region ID of the ECC instance.
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

func (s ModifyExpressCloudConnectionAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyExpressCloudConnectionAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyExpressCloudConnectionAttributeRequest) GetBgpAs() *string {
	return s.BgpAs
}

func (s *ModifyExpressCloudConnectionAttributeRequest) GetCeIp() *string {
	return s.CeIp
}

func (s *ModifyExpressCloudConnectionAttributeRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyExpressCloudConnectionAttributeRequest) GetEccId() *string {
	return s.EccId
}

func (s *ModifyExpressCloudConnectionAttributeRequest) GetName() *string {
	return s.Name
}

func (s *ModifyExpressCloudConnectionAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyExpressCloudConnectionAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyExpressCloudConnectionAttributeRequest) GetPeIp() *string {
	return s.PeIp
}

func (s *ModifyExpressCloudConnectionAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyExpressCloudConnectionAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyExpressCloudConnectionAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyExpressCloudConnectionAttributeRequest) SetBgpAs(v string) *ModifyExpressCloudConnectionAttributeRequest {
	s.BgpAs = &v
	return s
}

func (s *ModifyExpressCloudConnectionAttributeRequest) SetCeIp(v string) *ModifyExpressCloudConnectionAttributeRequest {
	s.CeIp = &v
	return s
}

func (s *ModifyExpressCloudConnectionAttributeRequest) SetDescription(v string) *ModifyExpressCloudConnectionAttributeRequest {
	s.Description = &v
	return s
}

func (s *ModifyExpressCloudConnectionAttributeRequest) SetEccId(v string) *ModifyExpressCloudConnectionAttributeRequest {
	s.EccId = &v
	return s
}

func (s *ModifyExpressCloudConnectionAttributeRequest) SetName(v string) *ModifyExpressCloudConnectionAttributeRequest {
	s.Name = &v
	return s
}

func (s *ModifyExpressCloudConnectionAttributeRequest) SetOwnerAccount(v string) *ModifyExpressCloudConnectionAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyExpressCloudConnectionAttributeRequest) SetOwnerId(v int64) *ModifyExpressCloudConnectionAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyExpressCloudConnectionAttributeRequest) SetPeIp(v string) *ModifyExpressCloudConnectionAttributeRequest {
	s.PeIp = &v
	return s
}

func (s *ModifyExpressCloudConnectionAttributeRequest) SetRegionId(v string) *ModifyExpressCloudConnectionAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyExpressCloudConnectionAttributeRequest) SetResourceOwnerAccount(v string) *ModifyExpressCloudConnectionAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyExpressCloudConnectionAttributeRequest) SetResourceOwnerId(v int64) *ModifyExpressCloudConnectionAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyExpressCloudConnectionAttributeRequest) Validate() error {
	return dara.Validate(s)
}
