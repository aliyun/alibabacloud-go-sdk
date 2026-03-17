// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddSnatEntryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCidrBlock(v string) *AddSnatEntryRequest
	GetCidrBlock() *string
	SetOwnerAccount(v string) *AddSnatEntryRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *AddSnatEntryRequest
	GetOwnerId() *int64
	SetRegionId(v string) *AddSnatEntryRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *AddSnatEntryRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AddSnatEntryRequest
	GetResourceOwnerId() *int64
	SetSmartAGId(v string) *AddSnatEntryRequest
	GetSmartAGId() *string
	SetSnatIp(v string) *AddSnatEntryRequest
	GetSnatIp() *string
}

type AddSnatEntryRequest struct {
	// The destination CIDR block.
	//
	// This parameter is required.
	//
	// example:
	//
	// 192.168.1.0/24
	CidrBlock    *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the SAG instance is deployed.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the SAG instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// sag-jf5w9a8k5mhi5h****
	SmartAGId *string `json:"SmartAGId,omitempty" xml:"SmartAGId,omitempty"`
	// The public IP address.
	//
	// This parameter is required.
	//
	// example:
	//
	// 11.0.XX.XX
	SnatIp *string `json:"SnatIp,omitempty" xml:"SnatIp,omitempty"`
}

func (s AddSnatEntryRequest) String() string {
	return dara.Prettify(s)
}

func (s AddSnatEntryRequest) GoString() string {
	return s.String()
}

func (s *AddSnatEntryRequest) GetCidrBlock() *string {
	return s.CidrBlock
}

func (s *AddSnatEntryRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *AddSnatEntryRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AddSnatEntryRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AddSnatEntryRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AddSnatEntryRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AddSnatEntryRequest) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *AddSnatEntryRequest) GetSnatIp() *string {
	return s.SnatIp
}

func (s *AddSnatEntryRequest) SetCidrBlock(v string) *AddSnatEntryRequest {
	s.CidrBlock = &v
	return s
}

func (s *AddSnatEntryRequest) SetOwnerAccount(v string) *AddSnatEntryRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AddSnatEntryRequest) SetOwnerId(v int64) *AddSnatEntryRequest {
	s.OwnerId = &v
	return s
}

func (s *AddSnatEntryRequest) SetRegionId(v string) *AddSnatEntryRequest {
	s.RegionId = &v
	return s
}

func (s *AddSnatEntryRequest) SetResourceOwnerAccount(v string) *AddSnatEntryRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AddSnatEntryRequest) SetResourceOwnerId(v int64) *AddSnatEntryRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AddSnatEntryRequest) SetSmartAGId(v string) *AddSnatEntryRequest {
	s.SmartAGId = &v
	return s
}

func (s *AddSnatEntryRequest) SetSnatIp(v string) *AddSnatEntryRequest {
	s.SnatIp = &v
	return s
}

func (s *AddSnatEntryRequest) Validate() error {
	return dara.Validate(s)
}
