// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSagExpressConnectInterfaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIP(v string) *CreateSagExpressConnectInterfaceRequest
	GetIP() *string
	SetMask(v string) *CreateSagExpressConnectInterfaceRequest
	GetMask() *string
	SetOwnerAccount(v string) *CreateSagExpressConnectInterfaceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateSagExpressConnectInterfaceRequest
	GetOwnerId() *int64
	SetPortName(v string) *CreateSagExpressConnectInterfaceRequest
	GetPortName() *string
	SetRegionId(v string) *CreateSagExpressConnectInterfaceRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CreateSagExpressConnectInterfaceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateSagExpressConnectInterfaceRequest
	GetResourceOwnerId() *int64
	SetSmartAGId(v string) *CreateSagExpressConnectInterfaceRequest
	GetSmartAGId() *string
	SetSmartAGSn(v string) *CreateSagExpressConnectInterfaceRequest
	GetSmartAGSn() *string
	SetVlan(v string) *CreateSagExpressConnectInterfaceRequest
	GetVlan() *string
}

type CreateSagExpressConnectInterfaceRequest struct {
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
	// Value: **0 to 4094**.
	//
	// >
	//
	// 	- If the VLAN ID is 0, this port is a physical port and does not support VLAN subinterfaces.
	//
	// 	- If the VLAN ID is 1 to 4094, this port supports VLAN subinterfaces based on the Layer 3 protocols.
	//
	// 	- You can create a maximum of 5 VLAN subinterfaces.
	//
	// example:
	//
	// 10
	Vlan *string `json:"Vlan,omitempty" xml:"Vlan,omitempty"`
}

func (s CreateSagExpressConnectInterfaceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSagExpressConnectInterfaceRequest) GoString() string {
	return s.String()
}

func (s *CreateSagExpressConnectInterfaceRequest) GetIP() *string {
	return s.IP
}

func (s *CreateSagExpressConnectInterfaceRequest) GetMask() *string {
	return s.Mask
}

func (s *CreateSagExpressConnectInterfaceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateSagExpressConnectInterfaceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateSagExpressConnectInterfaceRequest) GetPortName() *string {
	return s.PortName
}

func (s *CreateSagExpressConnectInterfaceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateSagExpressConnectInterfaceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateSagExpressConnectInterfaceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateSagExpressConnectInterfaceRequest) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *CreateSagExpressConnectInterfaceRequest) GetSmartAGSn() *string {
	return s.SmartAGSn
}

func (s *CreateSagExpressConnectInterfaceRequest) GetVlan() *string {
	return s.Vlan
}

func (s *CreateSagExpressConnectInterfaceRequest) SetIP(v string) *CreateSagExpressConnectInterfaceRequest {
	s.IP = &v
	return s
}

func (s *CreateSagExpressConnectInterfaceRequest) SetMask(v string) *CreateSagExpressConnectInterfaceRequest {
	s.Mask = &v
	return s
}

func (s *CreateSagExpressConnectInterfaceRequest) SetOwnerAccount(v string) *CreateSagExpressConnectInterfaceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateSagExpressConnectInterfaceRequest) SetOwnerId(v int64) *CreateSagExpressConnectInterfaceRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateSagExpressConnectInterfaceRequest) SetPortName(v string) *CreateSagExpressConnectInterfaceRequest {
	s.PortName = &v
	return s
}

func (s *CreateSagExpressConnectInterfaceRequest) SetRegionId(v string) *CreateSagExpressConnectInterfaceRequest {
	s.RegionId = &v
	return s
}

func (s *CreateSagExpressConnectInterfaceRequest) SetResourceOwnerAccount(v string) *CreateSagExpressConnectInterfaceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateSagExpressConnectInterfaceRequest) SetResourceOwnerId(v int64) *CreateSagExpressConnectInterfaceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateSagExpressConnectInterfaceRequest) SetSmartAGId(v string) *CreateSagExpressConnectInterfaceRequest {
	s.SmartAGId = &v
	return s
}

func (s *CreateSagExpressConnectInterfaceRequest) SetSmartAGSn(v string) *CreateSagExpressConnectInterfaceRequest {
	s.SmartAGSn = &v
	return s
}

func (s *CreateSagExpressConnectInterfaceRequest) SetVlan(v string) *CreateSagExpressConnectInterfaceRequest {
	s.Vlan = &v
	return s
}

func (s *CreateSagExpressConnectInterfaceRequest) Validate() error {
	return dara.Validate(s)
}
