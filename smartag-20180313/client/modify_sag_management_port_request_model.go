// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySagManagementPortRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGateway(v string) *ModifySagManagementPortRequest
	GetGateway() *string
	SetIP(v string) *ModifySagManagementPortRequest
	GetIP() *string
	SetMask(v string) *ModifySagManagementPortRequest
	GetMask() *string
	SetOwnerAccount(v string) *ModifySagManagementPortRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifySagManagementPortRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifySagManagementPortRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifySagManagementPortRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifySagManagementPortRequest
	GetResourceOwnerId() *int64
	SetSmartAGId(v string) *ModifySagManagementPortRequest
	GetSmartAGId() *string
	SetSmartAGSn(v string) *ModifySagManagementPortRequest
	GetSmartAGSn() *string
}

type ModifySagManagementPortRequest struct {
	// The IP address of the management port gateway.
	//
	// example:
	//
	// 192.XX.XX.254
	Gateway *string `json:"Gateway,omitempty" xml:"Gateway,omitempty"`
	// The IP address of the management port.
	//
	// This parameter is required.
	//
	// example:
	//
	// 192.XX.XX.1
	IP *string `json:"IP,omitempty" xml:"IP,omitempty"`
	// The subnet mask for the IP address of the management port.
	//
	// This parameter is required.
	//
	// example:
	//
	// 255.255.255.0
	Mask         *string `json:"Mask,omitempty" xml:"Mask,omitempty"`
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
}

func (s ModifySagManagementPortRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifySagManagementPortRequest) GoString() string {
	return s.String()
}

func (s *ModifySagManagementPortRequest) GetGateway() *string {
	return s.Gateway
}

func (s *ModifySagManagementPortRequest) GetIP() *string {
	return s.IP
}

func (s *ModifySagManagementPortRequest) GetMask() *string {
	return s.Mask
}

func (s *ModifySagManagementPortRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifySagManagementPortRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifySagManagementPortRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifySagManagementPortRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifySagManagementPortRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifySagManagementPortRequest) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *ModifySagManagementPortRequest) GetSmartAGSn() *string {
	return s.SmartAGSn
}

func (s *ModifySagManagementPortRequest) SetGateway(v string) *ModifySagManagementPortRequest {
	s.Gateway = &v
	return s
}

func (s *ModifySagManagementPortRequest) SetIP(v string) *ModifySagManagementPortRequest {
	s.IP = &v
	return s
}

func (s *ModifySagManagementPortRequest) SetMask(v string) *ModifySagManagementPortRequest {
	s.Mask = &v
	return s
}

func (s *ModifySagManagementPortRequest) SetOwnerAccount(v string) *ModifySagManagementPortRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifySagManagementPortRequest) SetOwnerId(v int64) *ModifySagManagementPortRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifySagManagementPortRequest) SetRegionId(v string) *ModifySagManagementPortRequest {
	s.RegionId = &v
	return s
}

func (s *ModifySagManagementPortRequest) SetResourceOwnerAccount(v string) *ModifySagManagementPortRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifySagManagementPortRequest) SetResourceOwnerId(v int64) *ModifySagManagementPortRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifySagManagementPortRequest) SetSmartAGId(v string) *ModifySagManagementPortRequest {
	s.SmartAGId = &v
	return s
}

func (s *ModifySagManagementPortRequest) SetSmartAGSn(v string) *ModifySagManagementPortRequest {
	s.SmartAGSn = &v
	return s
}

func (s *ModifySagManagementPortRequest) Validate() error {
	return dara.Validate(s)
}
