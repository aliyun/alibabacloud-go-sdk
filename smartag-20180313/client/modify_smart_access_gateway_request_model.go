// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySmartAccessGatewayRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCidrBlock(v string) *ModifySmartAccessGatewayRequest
	GetCidrBlock() *string
	SetDescription(v string) *ModifySmartAccessGatewayRequest
	GetDescription() *string
	SetEnableSoftwareConnectionAudit(v bool) *ModifySmartAccessGatewayRequest
	GetEnableSoftwareConnectionAudit() *bool
	SetName(v string) *ModifySmartAccessGatewayRequest
	GetName() *string
	SetOwnerAccount(v string) *ModifySmartAccessGatewayRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifySmartAccessGatewayRequest
	GetOwnerId() *int64
	SetPosition(v string) *ModifySmartAccessGatewayRequest
	GetPosition() *string
	SetRegionId(v string) *ModifySmartAccessGatewayRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifySmartAccessGatewayRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifySmartAccessGatewayRequest
	GetResourceOwnerId() *int64
	SetRoutingStrategy(v string) *ModifySmartAccessGatewayRequest
	GetRoutingStrategy() *string
	SetSecurityLockThreshold(v int32) *ModifySmartAccessGatewayRequest
	GetSecurityLockThreshold() *int32
	SetSmartAGId(v string) *ModifySmartAccessGatewayRequest
	GetSmartAGId() *string
}

type ModifySmartAccessGatewayRequest struct {
	// The CIDR blocks of terminals in the private network. Make sure that the CIDR blocks do not overlap with each other.
	//
	// If the LAN ports of the terminals use dynamic routing, the IP addresses within the first private CIDR block are allocated to the terminals that have Dynamic Host Configuration Protocol (DHCP) enabled.
	//
	// example:
	//
	// 172.16.0.0/24
	CidrBlock *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	// The description of the SAG instance.
	//
	// The description must be 2 to 256 characters in length. The description must start with a letter but cannot start with `http://` or `https://`.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to audit the network connection logs of the SAS app instance.
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// true
	EnableSoftwareConnectionAudit *bool `json:"EnableSoftwareConnectionAudit,omitempty" xml:"EnableSoftwareConnectionAudit,omitempty"`
	// The name of the SAG instance.
	//
	// The name must be 2 to 128 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). It must start with a letter.
	//
	// example:
	//
	// SAG
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The location where the SAG instance is deployed.
	Position *string `json:"Position,omitempty" xml:"Position,omitempty"`
	// The ID of the region where the SAG instance is deployed.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The policy that is used to advertise routes to Alibaba Cloud. Valid values:
	//
	// 	- **static**: static routing
	//
	// 	- **dynamic**: dynamic routing
	//
	// example:
	//
	// static
	RoutingStrategy *string `json:"RoutingStrategy,omitempty" xml:"RoutingStrategy,omitempty"`
	// The time during which the disconnected SAG instance remains locked. Valid values: an integer that is greater than or equal to 0.
	//
	// Unit: seconds.
	//
	// example:
	//
	// 3
	SecurityLockThreshold *int32 `json:"SecurityLockThreshold,omitempty" xml:"SecurityLockThreshold,omitempty"`
	// The ID of the SAG instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// sag-0ovhf732a9j0******
	SmartAGId *string `json:"SmartAGId,omitempty" xml:"SmartAGId,omitempty"`
}

func (s ModifySmartAccessGatewayRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifySmartAccessGatewayRequest) GoString() string {
	return s.String()
}

func (s *ModifySmartAccessGatewayRequest) GetCidrBlock() *string {
	return s.CidrBlock
}

func (s *ModifySmartAccessGatewayRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifySmartAccessGatewayRequest) GetEnableSoftwareConnectionAudit() *bool {
	return s.EnableSoftwareConnectionAudit
}

func (s *ModifySmartAccessGatewayRequest) GetName() *string {
	return s.Name
}

func (s *ModifySmartAccessGatewayRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifySmartAccessGatewayRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifySmartAccessGatewayRequest) GetPosition() *string {
	return s.Position
}

func (s *ModifySmartAccessGatewayRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifySmartAccessGatewayRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifySmartAccessGatewayRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifySmartAccessGatewayRequest) GetRoutingStrategy() *string {
	return s.RoutingStrategy
}

func (s *ModifySmartAccessGatewayRequest) GetSecurityLockThreshold() *int32 {
	return s.SecurityLockThreshold
}

func (s *ModifySmartAccessGatewayRequest) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *ModifySmartAccessGatewayRequest) SetCidrBlock(v string) *ModifySmartAccessGatewayRequest {
	s.CidrBlock = &v
	return s
}

func (s *ModifySmartAccessGatewayRequest) SetDescription(v string) *ModifySmartAccessGatewayRequest {
	s.Description = &v
	return s
}

func (s *ModifySmartAccessGatewayRequest) SetEnableSoftwareConnectionAudit(v bool) *ModifySmartAccessGatewayRequest {
	s.EnableSoftwareConnectionAudit = &v
	return s
}

func (s *ModifySmartAccessGatewayRequest) SetName(v string) *ModifySmartAccessGatewayRequest {
	s.Name = &v
	return s
}

func (s *ModifySmartAccessGatewayRequest) SetOwnerAccount(v string) *ModifySmartAccessGatewayRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifySmartAccessGatewayRequest) SetOwnerId(v int64) *ModifySmartAccessGatewayRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifySmartAccessGatewayRequest) SetPosition(v string) *ModifySmartAccessGatewayRequest {
	s.Position = &v
	return s
}

func (s *ModifySmartAccessGatewayRequest) SetRegionId(v string) *ModifySmartAccessGatewayRequest {
	s.RegionId = &v
	return s
}

func (s *ModifySmartAccessGatewayRequest) SetResourceOwnerAccount(v string) *ModifySmartAccessGatewayRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifySmartAccessGatewayRequest) SetResourceOwnerId(v int64) *ModifySmartAccessGatewayRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifySmartAccessGatewayRequest) SetRoutingStrategy(v string) *ModifySmartAccessGatewayRequest {
	s.RoutingStrategy = &v
	return s
}

func (s *ModifySmartAccessGatewayRequest) SetSecurityLockThreshold(v int32) *ModifySmartAccessGatewayRequest {
	s.SecurityLockThreshold = &v
	return s
}

func (s *ModifySmartAccessGatewayRequest) SetSmartAGId(v string) *ModifySmartAccessGatewayRequest {
	s.SmartAGId = &v
	return s
}

func (s *ModifySmartAccessGatewayRequest) Validate() error {
	return dara.Validate(s)
}
