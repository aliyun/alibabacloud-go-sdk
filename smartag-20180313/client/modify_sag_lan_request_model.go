// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySagLanRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndIp(v string) *ModifySagLanRequest
	GetEndIp() *string
	SetIP(v string) *ModifySagLanRequest
	GetIP() *string
	SetIPType(v string) *ModifySagLanRequest
	GetIPType() *string
	SetLease(v string) *ModifySagLanRequest
	GetLease() *string
	SetMask(v string) *ModifySagLanRequest
	GetMask() *string
	SetOwnerAccount(v string) *ModifySagLanRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifySagLanRequest
	GetOwnerId() *int64
	SetPortName(v string) *ModifySagLanRequest
	GetPortName() *string
	SetRegionId(v string) *ModifySagLanRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifySagLanRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifySagLanRequest
	GetResourceOwnerId() *int64
	SetSmartAGId(v string) *ModifySagLanRequest
	GetSmartAGId() *string
	SetSmartAGSn(v string) *ModifySagLanRequest
	GetSmartAGSn() *string
	SetStartIp(v string) *ModifySagLanRequest
	GetStartIp() *string
}

type ModifySagLanRequest struct {
	// The last IP address of the DHCP pool.
	//
	// example:
	//
	// 192.XX.XX.254
	EndIp *string `json:"EndIp,omitempty" xml:"EndIp,omitempty"`
	// The IP address of the LAN port.
	//
	// This parameter is required.
	//
	// example:
	//
	// 192.XX.XX.1
	IP *string `json:"IP,omitempty" xml:"IP,omitempty"`
	// The connection type of the LAN port. Valid values:
	//
	// 	- **DHCP**: a dynamic IP address. Uses the Dynamic Host Configuration Protocol (DHCP) to dynamically assign an IP address to a connected device.
	//
	// 	- **STATIC**: a static IP address. Specifies a static IP address for the LAN port.
	//
	// This parameter is required.
	//
	// example:
	//
	// STATIC
	IPType *string `json:"IPType,omitempty" xml:"IPType,omitempty"`
	// The time duration that the IP address is retained after it is assigned through DHCP. Unit: minute.
	//
	// Valid values: **1 to 43200**.
	//
	// example:
	//
	// 7
	Lease *string `json:"Lease,omitempty" xml:"Lease,omitempty"`
	// The subnet mask of the LAN port IP address.
	//
	// This parameter is required.
	//
	// example:
	//
	// 255.255.255.0
	Mask         *string `json:"Mask,omitempty" xml:"Mask,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The name of the LAN port.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	PortName *string `json:"PortName,omitempty" xml:"PortName,omitempty"`
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
	// The serial number of the SAG device that is associated with the SAG instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// sag32a30****
	SmartAGSn *string `json:"SmartAGSn,omitempty" xml:"SmartAGSn,omitempty"`
	// The first IP address of the DHCP pool.
	//
	// example:
	//
	// 192.XX.XX.2
	StartIp *string `json:"StartIp,omitempty" xml:"StartIp,omitempty"`
}

func (s ModifySagLanRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifySagLanRequest) GoString() string {
	return s.String()
}

func (s *ModifySagLanRequest) GetEndIp() *string {
	return s.EndIp
}

func (s *ModifySagLanRequest) GetIP() *string {
	return s.IP
}

func (s *ModifySagLanRequest) GetIPType() *string {
	return s.IPType
}

func (s *ModifySagLanRequest) GetLease() *string {
	return s.Lease
}

func (s *ModifySagLanRequest) GetMask() *string {
	return s.Mask
}

func (s *ModifySagLanRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifySagLanRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifySagLanRequest) GetPortName() *string {
	return s.PortName
}

func (s *ModifySagLanRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifySagLanRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifySagLanRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifySagLanRequest) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *ModifySagLanRequest) GetSmartAGSn() *string {
	return s.SmartAGSn
}

func (s *ModifySagLanRequest) GetStartIp() *string {
	return s.StartIp
}

func (s *ModifySagLanRequest) SetEndIp(v string) *ModifySagLanRequest {
	s.EndIp = &v
	return s
}

func (s *ModifySagLanRequest) SetIP(v string) *ModifySagLanRequest {
	s.IP = &v
	return s
}

func (s *ModifySagLanRequest) SetIPType(v string) *ModifySagLanRequest {
	s.IPType = &v
	return s
}

func (s *ModifySagLanRequest) SetLease(v string) *ModifySagLanRequest {
	s.Lease = &v
	return s
}

func (s *ModifySagLanRequest) SetMask(v string) *ModifySagLanRequest {
	s.Mask = &v
	return s
}

func (s *ModifySagLanRequest) SetOwnerAccount(v string) *ModifySagLanRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifySagLanRequest) SetOwnerId(v int64) *ModifySagLanRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifySagLanRequest) SetPortName(v string) *ModifySagLanRequest {
	s.PortName = &v
	return s
}

func (s *ModifySagLanRequest) SetRegionId(v string) *ModifySagLanRequest {
	s.RegionId = &v
	return s
}

func (s *ModifySagLanRequest) SetResourceOwnerAccount(v string) *ModifySagLanRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifySagLanRequest) SetResourceOwnerId(v int64) *ModifySagLanRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifySagLanRequest) SetSmartAGId(v string) *ModifySagLanRequest {
	s.SmartAGId = &v
	return s
}

func (s *ModifySagLanRequest) SetSmartAGSn(v string) *ModifySagLanRequest {
	s.SmartAGSn = &v
	return s
}

func (s *ModifySagLanRequest) SetStartIp(v string) *ModifySagLanRequest {
	s.StartIp = &v
	return s
}

func (s *ModifySagLanRequest) Validate() error {
	return dara.Validate(s)
}
