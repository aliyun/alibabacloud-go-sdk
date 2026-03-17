// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySagPortRoleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *ModifySagPortRoleRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifySagPortRoleRequest
	GetOwnerId() *int64
	SetPortName(v string) *ModifySagPortRoleRequest
	GetPortName() *string
	SetRegionId(v string) *ModifySagPortRoleRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifySagPortRoleRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifySagPortRoleRequest
	GetResourceOwnerId() *int64
	SetRole(v string) *ModifySagPortRoleRequest
	GetRole() *string
	SetSmartAGId(v string) *ModifySagPortRoleRequest
	GetSmartAGId() *string
	SetSmartAGSn(v string) *ModifySagPortRoleRequest
	GetSmartAGSn() *string
}

type ModifySagPortRoleRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The name of the port.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
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
	// The role of the port. Valid values:
	//
	// 	- **NONE**: No role is assigned to the port.
	//
	// 	- **WAN**: The port is used as a WAN port. By default, port 5 of an SAG device is used as a WAN port and supports access to the Internet by using a Dynamic Host Configuration Protocol (DHCP) client, PPPoE, or a static IP address.
	//
	// 	- **LAN**: The port is used as a LAN port. The LAN port allows a DHCP server or a static IP address to connect to an on-premises terminal or switch.
	//
	// 	- **ECC**: The port is used as a port to connect to an Express Connect circuit.
	//
	// 	- **MGT**: The port is used as the management port. By default, port 2 of an SAG device is used as the management port.
	//
	// >
	//
	// 	- In exclusive mode, the management traffic is separated from the workload traffic. The management port is only used to access the SAG web console and cannot be used to transfer your workload data. You can access the SAG web console only through the management port.
	//
	// 	- A WAN port can be used when it is connected. If a port is the first one that obtains an IP address over DHCP and can access the Internet, it is set as a WAN port. In the SAG web console, you can set another port as a WAN port.
	//
	// This parameter is required.
	//
	// example:
	//
	// NONE
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
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

func (s ModifySagPortRoleRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifySagPortRoleRequest) GoString() string {
	return s.String()
}

func (s *ModifySagPortRoleRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifySagPortRoleRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifySagPortRoleRequest) GetPortName() *string {
	return s.PortName
}

func (s *ModifySagPortRoleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifySagPortRoleRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifySagPortRoleRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifySagPortRoleRequest) GetRole() *string {
	return s.Role
}

func (s *ModifySagPortRoleRequest) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *ModifySagPortRoleRequest) GetSmartAGSn() *string {
	return s.SmartAGSn
}

func (s *ModifySagPortRoleRequest) SetOwnerAccount(v string) *ModifySagPortRoleRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifySagPortRoleRequest) SetOwnerId(v int64) *ModifySagPortRoleRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifySagPortRoleRequest) SetPortName(v string) *ModifySagPortRoleRequest {
	s.PortName = &v
	return s
}

func (s *ModifySagPortRoleRequest) SetRegionId(v string) *ModifySagPortRoleRequest {
	s.RegionId = &v
	return s
}

func (s *ModifySagPortRoleRequest) SetResourceOwnerAccount(v string) *ModifySagPortRoleRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifySagPortRoleRequest) SetResourceOwnerId(v int64) *ModifySagPortRoleRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifySagPortRoleRequest) SetRole(v string) *ModifySagPortRoleRequest {
	s.Role = &v
	return s
}

func (s *ModifySagPortRoleRequest) SetSmartAGId(v string) *ModifySagPortRoleRequest {
	s.SmartAGId = &v
	return s
}

func (s *ModifySagPortRoleRequest) SetSmartAGSn(v string) *ModifySagPortRoleRequest {
	s.SmartAGSn = &v
	return s
}

func (s *ModifySagPortRoleRequest) Validate() error {
	return dara.Validate(s)
}
