// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSagExpressConnectInterfaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DeleteSagExpressConnectInterfaceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteSagExpressConnectInterfaceRequest
	GetOwnerId() *int64
	SetPortName(v string) *DeleteSagExpressConnectInterfaceRequest
	GetPortName() *string
	SetRegionId(v string) *DeleteSagExpressConnectInterfaceRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteSagExpressConnectInterfaceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteSagExpressConnectInterfaceRequest
	GetResourceOwnerId() *int64
	SetSmartAGId(v string) *DeleteSagExpressConnectInterfaceRequest
	GetSmartAGId() *string
	SetSmartAGSn(v string) *DeleteSagExpressConnectInterfaceRequest
	GetSmartAGSn() *string
	SetVlan(v string) *DeleteSagExpressConnectInterfaceRequest
	GetVlan() *string
}

type DeleteSagExpressConnectInterfaceRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The name of the leased line port.
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
	// The VLAN ID of the leased line port.
	//
	// Valid values: **0 to 4094**.
	//
	// >
	//
	// 	- If the VLAN ID is 0, the leased line port is a physical port and does not support VLAN interfaces.
	//
	// 	- If the VLAN ID is 1 to 4094, the leased line port supports VLAN interfaces based on the Layer-3 protocols.
	//
	// example:
	//
	// 10
	Vlan *string `json:"Vlan,omitempty" xml:"Vlan,omitempty"`
}

func (s DeleteSagExpressConnectInterfaceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSagExpressConnectInterfaceRequest) GoString() string {
	return s.String()
}

func (s *DeleteSagExpressConnectInterfaceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteSagExpressConnectInterfaceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteSagExpressConnectInterfaceRequest) GetPortName() *string {
	return s.PortName
}

func (s *DeleteSagExpressConnectInterfaceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteSagExpressConnectInterfaceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteSagExpressConnectInterfaceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteSagExpressConnectInterfaceRequest) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *DeleteSagExpressConnectInterfaceRequest) GetSmartAGSn() *string {
	return s.SmartAGSn
}

func (s *DeleteSagExpressConnectInterfaceRequest) GetVlan() *string {
	return s.Vlan
}

func (s *DeleteSagExpressConnectInterfaceRequest) SetOwnerAccount(v string) *DeleteSagExpressConnectInterfaceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteSagExpressConnectInterfaceRequest) SetOwnerId(v int64) *DeleteSagExpressConnectInterfaceRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteSagExpressConnectInterfaceRequest) SetPortName(v string) *DeleteSagExpressConnectInterfaceRequest {
	s.PortName = &v
	return s
}

func (s *DeleteSagExpressConnectInterfaceRequest) SetRegionId(v string) *DeleteSagExpressConnectInterfaceRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteSagExpressConnectInterfaceRequest) SetResourceOwnerAccount(v string) *DeleteSagExpressConnectInterfaceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteSagExpressConnectInterfaceRequest) SetResourceOwnerId(v int64) *DeleteSagExpressConnectInterfaceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteSagExpressConnectInterfaceRequest) SetSmartAGId(v string) *DeleteSagExpressConnectInterfaceRequest {
	s.SmartAGId = &v
	return s
}

func (s *DeleteSagExpressConnectInterfaceRequest) SetSmartAGSn(v string) *DeleteSagExpressConnectInterfaceRequest {
	s.SmartAGSn = &v
	return s
}

func (s *DeleteSagExpressConnectInterfaceRequest) SetVlan(v string) *DeleteSagExpressConnectInterfaceRequest {
	s.Vlan = &v
	return s
}

func (s *DeleteSagExpressConnectInterfaceRequest) Validate() error {
	return dara.Validate(s)
}
