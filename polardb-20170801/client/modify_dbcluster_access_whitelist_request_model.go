// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBClusterAccessWhitelistRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterIPArrayAttribute(v string) *ModifyDBClusterAccessWhitelistRequest
	GetDBClusterIPArrayAttribute() *string
	SetDBClusterIPArrayName(v string) *ModifyDBClusterAccessWhitelistRequest
	GetDBClusterIPArrayName() *string
	SetDBClusterId(v string) *ModifyDBClusterAccessWhitelistRequest
	GetDBClusterId() *string
	SetModifyMode(v string) *ModifyDBClusterAccessWhitelistRequest
	GetModifyMode() *string
	SetOwnerAccount(v string) *ModifyDBClusterAccessWhitelistRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyDBClusterAccessWhitelistRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ModifyDBClusterAccessWhitelistRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyDBClusterAccessWhitelistRequest
	GetResourceOwnerId() *int64
	SetSecurityGroupIds(v string) *ModifyDBClusterAccessWhitelistRequest
	GetSecurityGroupIds() *string
	SetSecurityIps(v string) *ModifyDBClusterAccessWhitelistRequest
	GetSecurityIps() *string
	SetWhiteListType(v string) *ModifyDBClusterAccessWhitelistRequest
	GetWhiteListType() *string
}

type ModifyDBClusterAccessWhitelistRequest struct {
	// The attributes of the IP whitelist group. Set this parameter to **hidden*	- to hide the IP whitelist group in the console.
	//
	// > 	- The IP whitelist group that has appeared in the console cannot be hidden.
	//
	// > 	- This parameter can be specified only when the **WhiteListType*	- parameter is set to **IP**.
	//
	// example:
	//
	// hidden
	DBClusterIPArrayAttribute *string `json:"DBClusterIPArrayAttribute,omitempty" xml:"DBClusterIPArrayAttribute,omitempty"`
	// The name of the IP whitelist group. The group name must be 2 to 120 characters in length and consists of lowercase letters and digits. It must start with a letter, and end with a letter or a digit.
	//
	// 	- If the specified whitelist group name does not exist, the whitelist group is created.
	//
	// 	- If the specified whitelist group name exists, the whitelist group is modified.
	//
	// 	- If you do not specify this parameter, the default group is modified.
	//
	// >
	//
	// 	- You can create a maximum of 50 IP whitelist groups for a cluster.
	//
	// 	- This parameter can be specified only when the **WhiteListType*	- parameter is set to **IP**.
	//
	// example:
	//
	// default
	DBClusterIPArrayName *string `json:"DBClusterIPArrayName,omitempty" xml:"DBClusterIPArrayName,omitempty"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-*************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The method used to modify the IP whitelist. Valid values:
	//
	// 	- **Cover*	- (default): overwrites the original IP whitelist.
	//
	// 	- **Append**: appends IP addresses to the original whitelist.
	//
	// 	- **Delete**: deletes IP addresses.
	//
	// >  This parameter can be specified only when the **WhiteListType*	- parameter is set to **IP**.
	//
	// example:
	//
	// Cover
	ModifyMode           *string `json:"ModifyMode,omitempty" xml:"ModifyMode,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the security group. Separate multiple security groups with commas (,).
	//
	// > 	- You can add a maximum of three security groups to a cluster.
	//
	// > 	- This parameter can be specified only when the **WhiteListType*	- parameter is set to **SecurityGroup**.
	//
	// example:
	//
	// sg-*********
	SecurityGroupIds *string `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty"`
	// The IP addresses or Classless Inter-Domain Routing (CIDR) blocks in the IP whitelist group. You can add 1,000 IP addresses or CIDR blocks to all the IP whitelist groups. Separate multiple IP addresses with commas (,). The following two formats are supported:
	//
	// 	- IP addresses. Example: 10.23.12.24.
	//
	// 	- CIDR blocks. Example: 10.23.12.24/24. 24 indicates the length of the prefix of the CIDR block. The length is the range of 1 to 32.
	//
	// >  This parameter can be specified only when the **WhiteListType*	- parameter is set to **IP**.
	//
	// example:
	//
	// 10.23.12.24
	SecurityIps *string `json:"SecurityIps,omitempty" xml:"SecurityIps,omitempty"`
	// The type of the whitelist. Valid values:
	//
	// 	- **IP**: IP whitelist group.
	//
	// 	- **SecurityGroup**: security group.
	//
	// Default value: **IP**.
	//
	// example:
	//
	// IP
	WhiteListType *string `json:"WhiteListType,omitempty" xml:"WhiteListType,omitempty"`
}

func (s ModifyDBClusterAccessWhitelistRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBClusterAccessWhitelistRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterAccessWhitelistRequest) GetDBClusterIPArrayAttribute() *string {
	return s.DBClusterIPArrayAttribute
}

func (s *ModifyDBClusterAccessWhitelistRequest) GetDBClusterIPArrayName() *string {
	return s.DBClusterIPArrayName
}

func (s *ModifyDBClusterAccessWhitelistRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyDBClusterAccessWhitelistRequest) GetModifyMode() *string {
	return s.ModifyMode
}

func (s *ModifyDBClusterAccessWhitelistRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyDBClusterAccessWhitelistRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyDBClusterAccessWhitelistRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyDBClusterAccessWhitelistRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDBClusterAccessWhitelistRequest) GetSecurityGroupIds() *string {
	return s.SecurityGroupIds
}

func (s *ModifyDBClusterAccessWhitelistRequest) GetSecurityIps() *string {
	return s.SecurityIps
}

func (s *ModifyDBClusterAccessWhitelistRequest) GetWhiteListType() *string {
	return s.WhiteListType
}

func (s *ModifyDBClusterAccessWhitelistRequest) SetDBClusterIPArrayAttribute(v string) *ModifyDBClusterAccessWhitelistRequest {
	s.DBClusterIPArrayAttribute = &v
	return s
}

func (s *ModifyDBClusterAccessWhitelistRequest) SetDBClusterIPArrayName(v string) *ModifyDBClusterAccessWhitelistRequest {
	s.DBClusterIPArrayName = &v
	return s
}

func (s *ModifyDBClusterAccessWhitelistRequest) SetDBClusterId(v string) *ModifyDBClusterAccessWhitelistRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBClusterAccessWhitelistRequest) SetModifyMode(v string) *ModifyDBClusterAccessWhitelistRequest {
	s.ModifyMode = &v
	return s
}

func (s *ModifyDBClusterAccessWhitelistRequest) SetOwnerAccount(v string) *ModifyDBClusterAccessWhitelistRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBClusterAccessWhitelistRequest) SetOwnerId(v int64) *ModifyDBClusterAccessWhitelistRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBClusterAccessWhitelistRequest) SetResourceOwnerAccount(v string) *ModifyDBClusterAccessWhitelistRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBClusterAccessWhitelistRequest) SetResourceOwnerId(v int64) *ModifyDBClusterAccessWhitelistRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBClusterAccessWhitelistRequest) SetSecurityGroupIds(v string) *ModifyDBClusterAccessWhitelistRequest {
	s.SecurityGroupIds = &v
	return s
}

func (s *ModifyDBClusterAccessWhitelistRequest) SetSecurityIps(v string) *ModifyDBClusterAccessWhitelistRequest {
	s.SecurityIps = &v
	return s
}

func (s *ModifyDBClusterAccessWhitelistRequest) SetWhiteListType(v string) *ModifyDBClusterAccessWhitelistRequest {
	s.WhiteListType = &v
	return s
}

func (s *ModifyDBClusterAccessWhitelistRequest) Validate() error {
	return dara.Validate(s)
}
