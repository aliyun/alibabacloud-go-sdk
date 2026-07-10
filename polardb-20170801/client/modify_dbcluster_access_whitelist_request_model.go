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
	SetPfsInstanceId(v string) *ModifyDBClusterAccessWhitelistRequest
	GetPfsInstanceId() *string
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
	// The attribute of the IP whitelist group. If this parameter is set to **hidden**, the group is not displayed in the console.
	//
	// > - IP whitelist groups that are already displayed in the console cannot be hidden.
	//
	// > - This parameter takes effect only when **WhiteListType*	- is set to **IP**.
	//
	// example:
	//
	// hidden
	DBClusterIPArrayAttribute *string `json:"DBClusterIPArrayAttribute,omitempty" xml:"DBClusterIPArrayAttribute,omitempty"`
	// The name of the IP whitelist group. The name must be 2 to 120 characters in length and can contain lowercase letters and digits. The name must start with a letter and end with a letter or digit.
	//
	// - If the specified whitelist group name does not exist, a new whitelist group is created.
	//
	// - If the specified whitelist group name already exists, the whitelist group is modified.
	//
	// - If this parameter is not specified, the default group is modified.
	//
	// > - A maximum of 50 IP whitelist groups are supported for a cluster.
	//
	// > - This parameter takes effect only when **WhiteListType*	- is set to **IP**.
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
	// - **Cover**: overwrites the original IP whitelist (default value).
	//
	// - **Append**: appends IP addresses to the whitelist.
	//
	// - **Delete**: removes IP addresses from the whitelist.
	//
	// > This parameter takes effect only when **WhiteListType*	- is set to **IP**.
	//
	// example:
	//
	// Cover
	ModifyMode   *string `json:"ModifyMode,omitempty" xml:"ModifyMode,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// pfs-xxx
	PfsInstanceId        *string `json:"PfsInstanceId,omitempty" xml:"PfsInstanceId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The security group IDs. Separate multiple security group IDs with commas (,).
	//
	// > - A maximum of 3 security groups are supported for a cluster.
	//
	// > - This parameter takes effect only when **WhiteListType*	- is set to **SecurityGroup**.
	//
	// example:
	//
	// sg-*********
	SecurityGroupIds *string `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty"`
	// The IP addresses or CIDR blocks in the IP whitelist group. A maximum of 1,000 IP addresses or CIDR blocks can be added to all IP whitelist groups. Separate multiple IP addresses with commas (,). The following two formats are supported:
	//
	// - IP address format, such as 10.23.12.24.
	//
	// - CIDR format, such as 10.23.12.24/24, where 24 indicates the prefix length of the CIDR block. The prefix length ranges from 1 to 32.
	//
	// > This parameter takes effect only when **WhiteListType*	- is set to **IP**.
	//
	// example:
	//
	// 10.23.12.24
	SecurityIps *string `json:"SecurityIps,omitempty" xml:"SecurityIps,omitempty"`
	// The type of the whitelist. Valid values:
	//
	// - **IP**: IP whitelist group.
	//
	// - **SecurityGroup**: security group.
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

func (s *ModifyDBClusterAccessWhitelistRequest) GetPfsInstanceId() *string {
	return s.PfsInstanceId
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

func (s *ModifyDBClusterAccessWhitelistRequest) SetPfsInstanceId(v string) *ModifyDBClusterAccessWhitelistRequest {
	s.PfsInstanceId = &v
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
