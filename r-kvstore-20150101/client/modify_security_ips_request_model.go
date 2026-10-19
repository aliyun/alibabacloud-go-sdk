// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySecurityIpsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ModifySecurityIpsRequest
	GetInstanceId() *string
	SetModifyMode(v string) *ModifySecurityIpsRequest
	GetModifyMode() *string
	SetOwnerAccount(v string) *ModifySecurityIpsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifySecurityIpsRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ModifySecurityIpsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifySecurityIpsRequest
	GetResourceOwnerId() *int64
	SetSecurityIpGroupAttribute(v string) *ModifySecurityIpsRequest
	GetSecurityIpGroupAttribute() *string
	SetSecurityIpGroupName(v string) *ModifySecurityIpsRequest
	GetSecurityIpGroupName() *string
	SetSecurityIps(v string) *ModifySecurityIpsRequest
	GetSecurityIps() *string
	SetSecurityToken(v string) *ModifySecurityIpsRequest
	GetSecurityToken() *string
}

type ModifySecurityIpsRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The modification method. Valid values:
	//
	// 	- **Cover*	- (default): overwrites the original whitelist.
	//
	// 	- **Append**: appends IP addresses to the whitelist.
	//
	// 	- **Delete**: deletes the whitelist.
	//
	// example:
	//
	// Append
	ModifyMode           *string `json:"ModifyMode,omitempty" xml:"ModifyMode,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The attribute of the IP whitelist group. This parameter is empty by default. The console does not display the whitelist group if the value is set to **hidden**.
	//
	// example:
	//
	// hidden
	SecurityIpGroupAttribute *string `json:"SecurityIpGroupAttribute,omitempty" xml:"SecurityIpGroupAttribute,omitempty"`
	// The name of the IP whitelist group.
	//
	// <warning>You cannot modify a system-generated whitelist group. If you do not specify this parameter, the whitelist of the default group is modified.</warning>
	//
	// example:
	//
	// default
	SecurityIpGroupName *string `json:"SecurityIpGroupName,omitempty" xml:"SecurityIpGroupName,omitempty"`
	// The IP addresses in the IP whitelist group. You can add up to 1,000 IP addresses. Separate multiple IP addresses with commas (,). The following formats are supported: 0.0.0.0/0,10.23.12.24 or 10.23.12.24/24 (CIDR notation, where /24 specifies the prefix length, which ranges from 1 to 32).
	//
	// This parameter is required.
	//
	// example:
	//
	// 100.64.***.0/24,10.101.11.***
	SecurityIps   *string `json:"SecurityIps,omitempty" xml:"SecurityIps,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s ModifySecurityIpsRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifySecurityIpsRequest) GoString() string {
	return s.String()
}

func (s *ModifySecurityIpsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifySecurityIpsRequest) GetModifyMode() *string {
	return s.ModifyMode
}

func (s *ModifySecurityIpsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifySecurityIpsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifySecurityIpsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifySecurityIpsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifySecurityIpsRequest) GetSecurityIpGroupAttribute() *string {
	return s.SecurityIpGroupAttribute
}

func (s *ModifySecurityIpsRequest) GetSecurityIpGroupName() *string {
	return s.SecurityIpGroupName
}

func (s *ModifySecurityIpsRequest) GetSecurityIps() *string {
	return s.SecurityIps
}

func (s *ModifySecurityIpsRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *ModifySecurityIpsRequest) SetInstanceId(v string) *ModifySecurityIpsRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifySecurityIpsRequest) SetModifyMode(v string) *ModifySecurityIpsRequest {
	s.ModifyMode = &v
	return s
}

func (s *ModifySecurityIpsRequest) SetOwnerAccount(v string) *ModifySecurityIpsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifySecurityIpsRequest) SetOwnerId(v int64) *ModifySecurityIpsRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifySecurityIpsRequest) SetResourceOwnerAccount(v string) *ModifySecurityIpsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifySecurityIpsRequest) SetResourceOwnerId(v int64) *ModifySecurityIpsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifySecurityIpsRequest) SetSecurityIpGroupAttribute(v string) *ModifySecurityIpsRequest {
	s.SecurityIpGroupAttribute = &v
	return s
}

func (s *ModifySecurityIpsRequest) SetSecurityIpGroupName(v string) *ModifySecurityIpsRequest {
	s.SecurityIpGroupName = &v
	return s
}

func (s *ModifySecurityIpsRequest) SetSecurityIps(v string) *ModifySecurityIpsRequest {
	s.SecurityIps = &v
	return s
}

func (s *ModifySecurityIpsRequest) SetSecurityToken(v string) *ModifySecurityIpsRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifySecurityIpsRequest) Validate() error {
	return dara.Validate(s)
}
