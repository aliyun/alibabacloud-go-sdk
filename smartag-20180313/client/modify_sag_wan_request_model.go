// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySagWanRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBandwidth(v int32) *ModifySagWanRequest
	GetBandwidth() *int32
	SetGateway(v string) *ModifySagWanRequest
	GetGateway() *string
	SetIP(v string) *ModifySagWanRequest
	GetIP() *string
	SetIPType(v string) *ModifySagWanRequest
	GetIPType() *string
	SetISP(v string) *ModifySagWanRequest
	GetISP() *string
	SetMask(v string) *ModifySagWanRequest
	GetMask() *string
	SetOwnerAccount(v string) *ModifySagWanRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifySagWanRequest
	GetOwnerId() *int64
	SetPassword(v string) *ModifySagWanRequest
	GetPassword() *string
	SetPortName(v string) *ModifySagWanRequest
	GetPortName() *string
	SetPriority(v int32) *ModifySagWanRequest
	GetPriority() *int32
	SetRegionId(v string) *ModifySagWanRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifySagWanRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifySagWanRequest
	GetResourceOwnerId() *int64
	SetSmartAGId(v string) *ModifySagWanRequest
	GetSmartAGId() *string
	SetSmartAGSn(v string) *ModifySagWanRequest
	GetSmartAGSn() *string
	SetUsername(v string) *ModifySagWanRequest
	GetUsername() *string
	SetWeight(v int32) *ModifySagWanRequest
	GetWeight() *int32
}

type ModifySagWanRequest struct {
	// The maximum bandwidth of the WAN port. Unit: Mbit/s.
	//
	// example:
	//
	// 50
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The IP address of the gateway.
	//
	// >
	//
	// 	- If the WAN port uses a static IP address, you must set this parameter.
	//
	// 	- After this parameter is set, the SAG device generates a default route.
	//
	// example:
	//
	// 192.XX.XX.254
	Gateway *string `json:"Gateway,omitempty" xml:"Gateway,omitempty"`
	// The IP address of the WAN port.
	//
	// >  If the WAN port uses a static IP address, you must set this parameter.
	//
	// example:
	//
	// 192.XX.XX.1
	IP *string `json:"IP,omitempty" xml:"IP,omitempty"`
	// The connection type of the WAN port: Valid values:
	//
	// 	- **DHCP**: The WAN port connects to the Internet through an IP address that is dynamically assigned by the Dynamic Host Configuration Protocol (DHCP) server.
	//
	// 	- **STATIC**: The WAN port connects to the Internet through a static IP address. If you set this value, you must configure a static IP address, a subnet mask, and a gateway IP address for the WAN port.
	//
	// 	- **PPPOE**: The WAN port connects to the Internet through dial-up connections. If you set this value, you must enter the username and password of the PPPoE account provided by the Internet service provider (ISP).
	//
	// This parameter is required.
	//
	// example:
	//
	// DHCP
	IPType *string `json:"IPType,omitempty" xml:"IPType,omitempty"`
	// The ISP used by the WAN port. Valid values:
	//
	// 	- **CT**: China Telecom
	//
	// 	- **CM**: China Mobile
	//
	// 	- **CU**: China Unicom
	//
	// 	- **Other**: Other ISPs.
	//
	// example:
	//
	// CT
	ISP *string `json:"ISP,omitempty" xml:"ISP,omitempty"`
	// The subnet mask of the WAN port IP address.
	//
	// >  If the WAN port uses a static IP address, you must set this parameter.
	//
	// example:
	//
	// 255.255.255.0
	Mask         *string `json:"Mask,omitempty" xml:"Mask,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The password of the PPPoE account.
	//
	// The username must be 6 to 30 characters in length, and can contain digits and letters.
	//
	// >  If the connection type of the WAN port is PPPoE, you must set this parameter. If you do not need to modify the password, you can ignore this parameter.
	//
	// example:
	//
	// P12ppq***
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The number of the WAN port.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	PortName *string `json:"PortName,omitempty" xml:"PortName,omitempty"`
	// The priority of the WAN port.
	//
	// Valid values: **-1*	- and **1 to 50**.
	//
	// A smaller value indicates a higher priority. A value of **-1*	- indicates that traffic forwarding is disabled on the WAN port.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The ID of the region where the SAG instance is deployed.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/69813.html) operation to query the most recent region list.
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
	// The serial number of the SAG device.
	//
	// This parameter is required.
	//
	// example:
	//
	// sag32a30****
	SmartAGSn *string `json:"SmartAGSn,omitempty" xml:"SmartAGSn,omitempty"`
	// The username of the PPPoE account.
	//
	// The username must be 6 to 30 characters in length, and can contain digits and letters.
	//
	// >  If the connection type of the WAN port is PPPoE, you must set this parameter.
	//
	// example:
	//
	// P12ppp***
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
	// The weight of the WAN port.
	//
	// Valid values: **1 to 100**. Default value: **100**.
	//
	// example:
	//
	// 100
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s ModifySagWanRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifySagWanRequest) GoString() string {
	return s.String()
}

func (s *ModifySagWanRequest) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *ModifySagWanRequest) GetGateway() *string {
	return s.Gateway
}

func (s *ModifySagWanRequest) GetIP() *string {
	return s.IP
}

func (s *ModifySagWanRequest) GetIPType() *string {
	return s.IPType
}

func (s *ModifySagWanRequest) GetISP() *string {
	return s.ISP
}

func (s *ModifySagWanRequest) GetMask() *string {
	return s.Mask
}

func (s *ModifySagWanRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifySagWanRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifySagWanRequest) GetPassword() *string {
	return s.Password
}

func (s *ModifySagWanRequest) GetPortName() *string {
	return s.PortName
}

func (s *ModifySagWanRequest) GetPriority() *int32 {
	return s.Priority
}

func (s *ModifySagWanRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifySagWanRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifySagWanRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifySagWanRequest) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *ModifySagWanRequest) GetSmartAGSn() *string {
	return s.SmartAGSn
}

func (s *ModifySagWanRequest) GetUsername() *string {
	return s.Username
}

func (s *ModifySagWanRequest) GetWeight() *int32 {
	return s.Weight
}

func (s *ModifySagWanRequest) SetBandwidth(v int32) *ModifySagWanRequest {
	s.Bandwidth = &v
	return s
}

func (s *ModifySagWanRequest) SetGateway(v string) *ModifySagWanRequest {
	s.Gateway = &v
	return s
}

func (s *ModifySagWanRequest) SetIP(v string) *ModifySagWanRequest {
	s.IP = &v
	return s
}

func (s *ModifySagWanRequest) SetIPType(v string) *ModifySagWanRequest {
	s.IPType = &v
	return s
}

func (s *ModifySagWanRequest) SetISP(v string) *ModifySagWanRequest {
	s.ISP = &v
	return s
}

func (s *ModifySagWanRequest) SetMask(v string) *ModifySagWanRequest {
	s.Mask = &v
	return s
}

func (s *ModifySagWanRequest) SetOwnerAccount(v string) *ModifySagWanRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifySagWanRequest) SetOwnerId(v int64) *ModifySagWanRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifySagWanRequest) SetPassword(v string) *ModifySagWanRequest {
	s.Password = &v
	return s
}

func (s *ModifySagWanRequest) SetPortName(v string) *ModifySagWanRequest {
	s.PortName = &v
	return s
}

func (s *ModifySagWanRequest) SetPriority(v int32) *ModifySagWanRequest {
	s.Priority = &v
	return s
}

func (s *ModifySagWanRequest) SetRegionId(v string) *ModifySagWanRequest {
	s.RegionId = &v
	return s
}

func (s *ModifySagWanRequest) SetResourceOwnerAccount(v string) *ModifySagWanRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifySagWanRequest) SetResourceOwnerId(v int64) *ModifySagWanRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifySagWanRequest) SetSmartAGId(v string) *ModifySagWanRequest {
	s.SmartAGId = &v
	return s
}

func (s *ModifySagWanRequest) SetSmartAGSn(v string) *ModifySagWanRequest {
	s.SmartAGSn = &v
	return s
}

func (s *ModifySagWanRequest) SetUsername(v string) *ModifySagWanRequest {
	s.Username = &v
	return s
}

func (s *ModifySagWanRequest) SetWeight(v int32) *ModifySagWanRequest {
	s.Weight = &v
	return s
}

func (s *ModifySagWanRequest) Validate() error {
	return dara.Validate(s)
}
