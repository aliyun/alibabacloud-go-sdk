// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySagExpressConnectInterfaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIP(v string) *ModifySagExpressConnectInterfaceRequest
	GetIP() *string
	SetMask(v string) *ModifySagExpressConnectInterfaceRequest
	GetMask() *string
	SetOwnerAccount(v string) *ModifySagExpressConnectInterfaceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifySagExpressConnectInterfaceRequest
	GetOwnerId() *int64
	SetPortName(v string) *ModifySagExpressConnectInterfaceRequest
	GetPortName() *string
	SetRegionId(v string) *ModifySagExpressConnectInterfaceRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifySagExpressConnectInterfaceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifySagExpressConnectInterfaceRequest
	GetResourceOwnerId() *int64
	SetSmartAGId(v string) *ModifySagExpressConnectInterfaceRequest
	GetSmartAGId() *string
	SetSmartAGSn(v string) *ModifySagExpressConnectInterfaceRequest
	GetSmartAGSn() *string
	SetVlan(v string) *ModifySagExpressConnectInterfaceRequest
	GetVlan() *string
}

type ModifySagExpressConnectInterfaceRequest struct {
	// The IP address.
	//
	// This parameter is required.
	//
	// example:
	//
	// 192.XX.XX.1
	IP *string `json:"IP,omitempty" xml:"IP,omitempty"`
	// The subnet mask of the IP address.
	//
	// This parameter is required.
	//
	// example:
	//
	// 255.255.255.252
	Mask         *string `json:"Mask,omitempty" xml:"Mask,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The name of the Express Connect port.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5
	PortName *string `json:"PortName,omitempty" xml:"PortName,omitempty"`
	// The ID of the region where the Smart Access Gateway (SAG) instance is deployed.
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
	// The serial number of the SAG device associated with the SAG instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// sag32a30****
	SmartAGSn *string `json:"SmartAGSn,omitempty" xml:"SmartAGSn,omitempty"`
	// The VLAN ID of the subinterface.
	//
	// Valid values: **0 to 4094**.
	//
	// >
	//
	// 	- If the VLAN ID is 0, this port is a physical port and does not support VLAN subinterfaces.
	//
	// 	- If the VLAN ID is 1 to 4094, this port supports VLAN subinterfaces based on the Layer 3 protocols.
	//
	// example:
	//
	// 10
	Vlan *string `json:"Vlan,omitempty" xml:"Vlan,omitempty"`
}

func (s ModifySagExpressConnectInterfaceRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifySagExpressConnectInterfaceRequest) GoString() string {
	return s.String()
}

func (s *ModifySagExpressConnectInterfaceRequest) GetIP() *string {
	return s.IP
}

func (s *ModifySagExpressConnectInterfaceRequest) GetMask() *string {
	return s.Mask
}

func (s *ModifySagExpressConnectInterfaceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifySagExpressConnectInterfaceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifySagExpressConnectInterfaceRequest) GetPortName() *string {
	return s.PortName
}

func (s *ModifySagExpressConnectInterfaceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifySagExpressConnectInterfaceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifySagExpressConnectInterfaceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifySagExpressConnectInterfaceRequest) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *ModifySagExpressConnectInterfaceRequest) GetSmartAGSn() *string {
	return s.SmartAGSn
}

func (s *ModifySagExpressConnectInterfaceRequest) GetVlan() *string {
	return s.Vlan
}

func (s *ModifySagExpressConnectInterfaceRequest) SetIP(v string) *ModifySagExpressConnectInterfaceRequest {
	s.IP = &v
	return s
}

func (s *ModifySagExpressConnectInterfaceRequest) SetMask(v string) *ModifySagExpressConnectInterfaceRequest {
	s.Mask = &v
	return s
}

func (s *ModifySagExpressConnectInterfaceRequest) SetOwnerAccount(v string) *ModifySagExpressConnectInterfaceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifySagExpressConnectInterfaceRequest) SetOwnerId(v int64) *ModifySagExpressConnectInterfaceRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifySagExpressConnectInterfaceRequest) SetPortName(v string) *ModifySagExpressConnectInterfaceRequest {
	s.PortName = &v
	return s
}

func (s *ModifySagExpressConnectInterfaceRequest) SetRegionId(v string) *ModifySagExpressConnectInterfaceRequest {
	s.RegionId = &v
	return s
}

func (s *ModifySagExpressConnectInterfaceRequest) SetResourceOwnerAccount(v string) *ModifySagExpressConnectInterfaceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifySagExpressConnectInterfaceRequest) SetResourceOwnerId(v int64) *ModifySagExpressConnectInterfaceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifySagExpressConnectInterfaceRequest) SetSmartAGId(v string) *ModifySagExpressConnectInterfaceRequest {
	s.SmartAGId = &v
	return s
}

func (s *ModifySagExpressConnectInterfaceRequest) SetSmartAGSn(v string) *ModifySagExpressConnectInterfaceRequest {
	s.SmartAGSn = &v
	return s
}

func (s *ModifySagExpressConnectInterfaceRequest) SetVlan(v string) *ModifySagExpressConnectInterfaceRequest {
	s.Vlan = &v
	return s
}

func (s *ModifySagExpressConnectInterfaceRequest) Validate() error {
	return dara.Validate(s)
}
