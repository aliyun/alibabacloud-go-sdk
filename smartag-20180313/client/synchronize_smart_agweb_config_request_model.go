// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSynchronizeSmartAGWebConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *SynchronizeSmartAGWebConfigRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *SynchronizeSmartAGWebConfigRequest
	GetOwnerId() *int64
	SetRegionId(v string) *SynchronizeSmartAGWebConfigRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *SynchronizeSmartAGWebConfigRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *SynchronizeSmartAGWebConfigRequest
	GetResourceOwnerId() *int64
	SetSmartAGId(v string) *SynchronizeSmartAGWebConfigRequest
	GetSmartAGId() *string
	SetSmartAGSn(v string) *SynchronizeSmartAGWebConfigRequest
	GetSmartAGSn() *string
}

type SynchronizeSmartAGWebConfigRequest struct {
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
	// sag-nylv14tghsk26c****
	SmartAGId *string `json:"SmartAGId,omitempty" xml:"SmartAGId,omitempty"`
	// The serial number of the SAG device.
	//
	// This parameter is required.
	//
	// example:
	//
	// sag32a30****
	SmartAGSn *string `json:"SmartAGSn,omitempty" xml:"SmartAGSn,omitempty"`
}

func (s SynchronizeSmartAGWebConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s SynchronizeSmartAGWebConfigRequest) GoString() string {
	return s.String()
}

func (s *SynchronizeSmartAGWebConfigRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *SynchronizeSmartAGWebConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SynchronizeSmartAGWebConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SynchronizeSmartAGWebConfigRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SynchronizeSmartAGWebConfigRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *SynchronizeSmartAGWebConfigRequest) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *SynchronizeSmartAGWebConfigRequest) GetSmartAGSn() *string {
	return s.SmartAGSn
}

func (s *SynchronizeSmartAGWebConfigRequest) SetOwnerAccount(v string) *SynchronizeSmartAGWebConfigRequest {
	s.OwnerAccount = &v
	return s
}

func (s *SynchronizeSmartAGWebConfigRequest) SetOwnerId(v int64) *SynchronizeSmartAGWebConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *SynchronizeSmartAGWebConfigRequest) SetRegionId(v string) *SynchronizeSmartAGWebConfigRequest {
	s.RegionId = &v
	return s
}

func (s *SynchronizeSmartAGWebConfigRequest) SetResourceOwnerAccount(v string) *SynchronizeSmartAGWebConfigRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SynchronizeSmartAGWebConfigRequest) SetResourceOwnerId(v int64) *SynchronizeSmartAGWebConfigRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SynchronizeSmartAGWebConfigRequest) SetSmartAGId(v string) *SynchronizeSmartAGWebConfigRequest {
	s.SmartAGId = &v
	return s
}

func (s *SynchronizeSmartAGWebConfigRequest) SetSmartAGSn(v string) *SynchronizeSmartAGWebConfigRequest {
	s.SmartAGSn = &v
	return s
}

func (s *SynchronizeSmartAGWebConfigRequest) Validate() error {
	return dara.Validate(s)
}
