// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySagGlobalRouteProtocolRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *ModifySagGlobalRouteProtocolRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifySagGlobalRouteProtocolRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifySagGlobalRouteProtocolRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifySagGlobalRouteProtocolRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifySagGlobalRouteProtocolRequest
	GetResourceOwnerId() *int64
	SetRouteProtocol(v string) *ModifySagGlobalRouteProtocolRequest
	GetRouteProtocol() *string
	SetSmartAGId(v string) *ModifySagGlobalRouteProtocolRequest
	GetSmartAGId() *string
	SetSmartAGSn(v string) *ModifySagGlobalRouteProtocolRequest
	GetSmartAGSn() *string
}

type ModifySagGlobalRouteProtocolRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the SAG instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The routing protocol. Valid values:
	//
	// 	- **STATIC**: static routing protocol
	//
	// 	- **OSPF**: Open Shortest Path First protocol (OSPF)
	//
	// 	- **BGP**: Border Gateway Protocol (BGP)
	//
	// This parameter is required.
	//
	// example:
	//
	// OSPF
	RouteProtocol *string `json:"RouteProtocol,omitempty" xml:"RouteProtocol,omitempty"`
	// The ID of the SAG instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// sag-whfn****
	SmartAGId *string `json:"SmartAGId,omitempty" xml:"SmartAGId,omitempty"`
	// The serial number of the SAG device that is associated with the SAG instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// sag32a30****
	SmartAGSn *string `json:"SmartAGSn,omitempty" xml:"SmartAGSn,omitempty"`
}

func (s ModifySagGlobalRouteProtocolRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifySagGlobalRouteProtocolRequest) GoString() string {
	return s.String()
}

func (s *ModifySagGlobalRouteProtocolRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifySagGlobalRouteProtocolRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifySagGlobalRouteProtocolRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifySagGlobalRouteProtocolRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifySagGlobalRouteProtocolRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifySagGlobalRouteProtocolRequest) GetRouteProtocol() *string {
	return s.RouteProtocol
}

func (s *ModifySagGlobalRouteProtocolRequest) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *ModifySagGlobalRouteProtocolRequest) GetSmartAGSn() *string {
	return s.SmartAGSn
}

func (s *ModifySagGlobalRouteProtocolRequest) SetOwnerAccount(v string) *ModifySagGlobalRouteProtocolRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifySagGlobalRouteProtocolRequest) SetOwnerId(v int64) *ModifySagGlobalRouteProtocolRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifySagGlobalRouteProtocolRequest) SetRegionId(v string) *ModifySagGlobalRouteProtocolRequest {
	s.RegionId = &v
	return s
}

func (s *ModifySagGlobalRouteProtocolRequest) SetResourceOwnerAccount(v string) *ModifySagGlobalRouteProtocolRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifySagGlobalRouteProtocolRequest) SetResourceOwnerId(v int64) *ModifySagGlobalRouteProtocolRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifySagGlobalRouteProtocolRequest) SetRouteProtocol(v string) *ModifySagGlobalRouteProtocolRequest {
	s.RouteProtocol = &v
	return s
}

func (s *ModifySagGlobalRouteProtocolRequest) SetSmartAGId(v string) *ModifySagGlobalRouteProtocolRequest {
	s.SmartAGId = &v
	return s
}

func (s *ModifySagGlobalRouteProtocolRequest) SetSmartAGSn(v string) *ModifySagGlobalRouteProtocolRequest {
	s.SmartAGSn = &v
	return s
}

func (s *ModifySagGlobalRouteProtocolRequest) Validate() error {
	return dara.Validate(s)
}
