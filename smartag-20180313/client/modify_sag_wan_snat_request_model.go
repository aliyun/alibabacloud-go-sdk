// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySagWanSnatRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *ModifySagWanSnatRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifySagWanSnatRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifySagWanSnatRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifySagWanSnatRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifySagWanSnatRequest
	GetResourceOwnerId() *int64
	SetSmartAGId(v string) *ModifySagWanSnatRequest
	GetSmartAGId() *string
	SetSmartAGSn(v string) *ModifySagWanSnatRequest
	GetSmartAGSn() *string
	SetSnat(v string) *ModifySagWanSnatRequest
	GetSnat() *string
}

type ModifySagWanSnatRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the SAG instance is deployed.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the SAG instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// sag-whfn****
	SmartAGId *string `json:"SmartAGId,omitempty" xml:"SmartAGId,omitempty"`
	// The serial number of the SAG device associated with the SAG instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// sag32a30****
	SmartAGSn *string `json:"SmartAGSn,omitempty" xml:"SmartAGSn,omitempty"`
	// Specifies whether to enable SNAT. Valid values:
	//
	// 	- **ENABLE**: enables SNAT.
	//
	// 	- **DISABLE**: disables SNAT.
	//
	// This parameter is required.
	//
	// example:
	//
	// ENABLE
	Snat *string `json:"Snat,omitempty" xml:"Snat,omitempty"`
}

func (s ModifySagWanSnatRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifySagWanSnatRequest) GoString() string {
	return s.String()
}

func (s *ModifySagWanSnatRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifySagWanSnatRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifySagWanSnatRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifySagWanSnatRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifySagWanSnatRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifySagWanSnatRequest) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *ModifySagWanSnatRequest) GetSmartAGSn() *string {
	return s.SmartAGSn
}

func (s *ModifySagWanSnatRequest) GetSnat() *string {
	return s.Snat
}

func (s *ModifySagWanSnatRequest) SetOwnerAccount(v string) *ModifySagWanSnatRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifySagWanSnatRequest) SetOwnerId(v int64) *ModifySagWanSnatRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifySagWanSnatRequest) SetRegionId(v string) *ModifySagWanSnatRequest {
	s.RegionId = &v
	return s
}

func (s *ModifySagWanSnatRequest) SetResourceOwnerAccount(v string) *ModifySagWanSnatRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifySagWanSnatRequest) SetResourceOwnerId(v int64) *ModifySagWanSnatRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifySagWanSnatRequest) SetSmartAGId(v string) *ModifySagWanSnatRequest {
	s.SmartAGId = &v
	return s
}

func (s *ModifySagWanSnatRequest) SetSmartAGSn(v string) *ModifySagWanSnatRequest {
	s.SmartAGSn = &v
	return s
}

func (s *ModifySagWanSnatRequest) SetSnat(v string) *ModifySagWanSnatRequest {
	s.Snat = &v
	return s
}

func (s *ModifySagWanSnatRequest) Validate() error {
	return dara.Validate(s)
}
