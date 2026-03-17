// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySagRouteProtocolBgpRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHoldTime(v int32) *ModifySagRouteProtocolBgpRequest
	GetHoldTime() *int32
	SetKeepAlive(v int32) *ModifySagRouteProtocolBgpRequest
	GetKeepAlive() *int32
	SetLocalAs(v int64) *ModifySagRouteProtocolBgpRequest
	GetLocalAs() *int64
	SetOwnerAccount(v string) *ModifySagRouteProtocolBgpRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifySagRouteProtocolBgpRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifySagRouteProtocolBgpRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifySagRouteProtocolBgpRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifySagRouteProtocolBgpRequest
	GetResourceOwnerId() *int64
	SetRouterId(v string) *ModifySagRouteProtocolBgpRequest
	GetRouterId() *string
	SetSmartAGId(v string) *ModifySagRouteProtocolBgpRequest
	GetSmartAGId() *string
	SetSmartAGSn(v string) *ModifySagRouteProtocolBgpRequest
	GetSmartAGSn() *string
}

type ModifySagRouteProtocolBgpRequest struct {
	// The hold time.
	//
	// Valid values: **3 to 65535**. Unit: seconds.
	//
	// >  When the SAG device establishes a peering connection with a peer device, the hold time of both devices must be the same. If the SAG device does not receive a keepalive or update message from the peer device within the hold time, the connection between the BGP peers is closed.
	//
	// This parameter is required.
	//
	// example:
	//
	// 9
	HoldTime *int32 `json:"HoldTime,omitempty" xml:"HoldTime,omitempty"`
	// The time interval at which keep-alive packets are sent.
	//
	// Valid values: **0 to 65535**. Unit: seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3
	KeepAlive *int32 `json:"KeepAlive,omitempty" xml:"KeepAlive,omitempty"`
	// The autonomous system number (ASN) to which the SAG device belongs.
	//
	// Valid values: **1 to 4294967295**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 65535
	LocalAs      *int64  `json:"LocalAs,omitempty" xml:"LocalAs,omitempty"`
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
	// The ID of the BGP router.
	//
	// Set the value to an IPv4 address.
	//
	// This parameter is required.
	//
	// example:
	//
	// 172.XX.XX.1
	RouterId *string `json:"RouterId,omitempty" xml:"RouterId,omitempty"`
	// The ID of the SAG instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// sag-whfn****
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

func (s ModifySagRouteProtocolBgpRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifySagRouteProtocolBgpRequest) GoString() string {
	return s.String()
}

func (s *ModifySagRouteProtocolBgpRequest) GetHoldTime() *int32 {
	return s.HoldTime
}

func (s *ModifySagRouteProtocolBgpRequest) GetKeepAlive() *int32 {
	return s.KeepAlive
}

func (s *ModifySagRouteProtocolBgpRequest) GetLocalAs() *int64 {
	return s.LocalAs
}

func (s *ModifySagRouteProtocolBgpRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifySagRouteProtocolBgpRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifySagRouteProtocolBgpRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifySagRouteProtocolBgpRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifySagRouteProtocolBgpRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifySagRouteProtocolBgpRequest) GetRouterId() *string {
	return s.RouterId
}

func (s *ModifySagRouteProtocolBgpRequest) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *ModifySagRouteProtocolBgpRequest) GetSmartAGSn() *string {
	return s.SmartAGSn
}

func (s *ModifySagRouteProtocolBgpRequest) SetHoldTime(v int32) *ModifySagRouteProtocolBgpRequest {
	s.HoldTime = &v
	return s
}

func (s *ModifySagRouteProtocolBgpRequest) SetKeepAlive(v int32) *ModifySagRouteProtocolBgpRequest {
	s.KeepAlive = &v
	return s
}

func (s *ModifySagRouteProtocolBgpRequest) SetLocalAs(v int64) *ModifySagRouteProtocolBgpRequest {
	s.LocalAs = &v
	return s
}

func (s *ModifySagRouteProtocolBgpRequest) SetOwnerAccount(v string) *ModifySagRouteProtocolBgpRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifySagRouteProtocolBgpRequest) SetOwnerId(v int64) *ModifySagRouteProtocolBgpRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifySagRouteProtocolBgpRequest) SetRegionId(v string) *ModifySagRouteProtocolBgpRequest {
	s.RegionId = &v
	return s
}

func (s *ModifySagRouteProtocolBgpRequest) SetResourceOwnerAccount(v string) *ModifySagRouteProtocolBgpRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifySagRouteProtocolBgpRequest) SetResourceOwnerId(v int64) *ModifySagRouteProtocolBgpRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifySagRouteProtocolBgpRequest) SetRouterId(v string) *ModifySagRouteProtocolBgpRequest {
	s.RouterId = &v
	return s
}

func (s *ModifySagRouteProtocolBgpRequest) SetSmartAGId(v string) *ModifySagRouteProtocolBgpRequest {
	s.SmartAGId = &v
	return s
}

func (s *ModifySagRouteProtocolBgpRequest) SetSmartAGSn(v string) *ModifySagRouteProtocolBgpRequest {
	s.SmartAGSn = &v
	return s
}

func (s *ModifySagRouteProtocolBgpRequest) Validate() error {
	return dara.Validate(s)
}
