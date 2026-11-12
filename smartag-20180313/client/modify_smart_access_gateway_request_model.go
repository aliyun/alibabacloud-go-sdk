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
	// The private CIDR block used by clients at the on-premises site for communication. Make sure that private CIDR blocks do not conflict with each other.
	//
	// If the LAN port of the Smart Access Gateway device is configured in dynamic mode, the IP addresses used by on-premises clients with DHCP enabled are allocated from the first private CIDR block that you specify.
	//
	// example:
	//
	// 172.16.0.0/24
	CidrBlock *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	// The description of the Smart Access Gateway instance.
	//
	// The description must be 2 to 256 characters in length, and must start with a letter or a Chinese character. It cannot start with `http://` or `https://`.
	//
	// example:
	//
	// SAG description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to enable connection log auditing for the SAG app. Valid values:
	//
	// - **true**: enables connection log auditing for the SAG app.
	//
	// - **false**: disables connection log auditing for the SAG app.
	//
	// example:
	//
	// true
	EnableSoftwareConnectionAudit *bool `json:"EnableSoftwareConnectionAudit,omitempty" xml:"EnableSoftwareConnectionAudit,omitempty"`
	// The name of the Smart Access Gateway instance.
	//
	// The name must be 2 to 128 characters in length, and must start with a letter or a Chinese character. It can contain digits, periods (.), underscores (_), and hyphens (-).
	//
	// example:
	//
	// SAG
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The location where the Smart Access Gateway instance is deployed.
	//
	// example:
	//
	// [ { "adcode":"", "label":"Hangzhou, Zhejiang xx zone xx Road xx number", "value":"B0FFGxxxx", "center":"120.190xxxx, 30.182xxxx", "_new":"true" } ]
	Position *string `json:"Position,omitempty" xml:"Position,omitempty"`
	// The region ID of the Smart Access Gateway instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The route synchronization mode for the on-premises network. Valid values:
	//
	// - **static**: static routing.
	//
	// - **dynamic**: dynamic routing.
	//
	// example:
	//
	// static
	RoutingStrategy *string `json:"RoutingStrategy,omitempty" xml:"RoutingStrategy,omitempty"`
	// The threshold configured by the user for the offline lock feature. The value must be greater than or equal to 0.
	//
	// Unit: seconds.
	//
	// example:
	//
	// 3
	SecurityLockThreshold *int32 `json:"SecurityLockThreshold,omitempty" xml:"SecurityLockThreshold,omitempty"`
	// The instance ID of the Smart Access Gateway.
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
