// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountDescription(v string) *CreateAccountRequest
	GetAccountDescription() *string
	SetAccountName(v string) *CreateAccountRequest
	GetAccountName() *string
	SetAccountPassword(v string) *CreateAccountRequest
	GetAccountPassword() *string
	SetCharacterType(v string) *CreateAccountRequest
	GetCharacterType() *string
	SetDBInstanceId(v string) *CreateAccountRequest
	GetDBInstanceId() *string
	SetDBName(v string) *CreateAccountRequest
	GetDBName() *string
	SetOwnerAccount(v string) *CreateAccountRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateAccountRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CreateAccountRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateAccountRequest
	GetResourceOwnerId() *int64
	SetRoles(v []*CreateAccountRequestRoles) *CreateAccountRequest
	GetRoles() []*CreateAccountRequestRoles
}

type CreateAccountRequest struct {
	// The account description.
	//
	// - Regular expression: `^[a-zA-Z一-龥][a-zA-Z一-龥0-9_.-]{1,255}$`. The description must start with an English letter or Chinese character, followed by English letters, Chinese characters, digits, underscores (_), periods (.), or hyphens (-). The length must be 2 to 256 characters.
	//
	// - If the description is invalid, InvalidAccountDescription.Malformed is returned.
	//
	// example:
	//
	// test
	AccountDescription *string `json:"AccountDescription,omitempty" xml:"AccountDescription,omitempty"`
	// The account name.
	//
	// - Regular expression: `^[a-z]+[a-z0-9|_]*[a-z0-9]+$`. The name must start with a lowercase letter, end with a lowercase letter or digit, and can contain lowercase letters, digits, and underscores (_). The length must be 2 to 16 characters.
	//
	// - The reserved name `root` (case-insensitive, built-in system account) cannot be used.
	//
	// - The name cannot start with `__cloud` (reserved prefix for system accounts).
	//
	// - If the name is invalid, `InvalidAccountName.Malformed` is returned.
	//
	// This parameter is required.
	//
	// example:
	//
	// admin1
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The account password.
	//
	// - The length must be 8 to 32 characters. It can contain uppercase letters, lowercase letters, digits, and the following special characters: `!@#$%^&*()_+-=`.
	//
	// - The password must contain at least **three*	- of the following four character types: uppercase letters, lowercase letters, digits, and special characters.
	//
	// - If the password is invalid, `InvalidAccountPassword.Malformed` is returned.
	//
	//
	// > For information about how to handle special characters, see [How do I resolve the connection failure caused by special characters in the account name or password of a connection string?](https://help.aliyun.com/document_detail/471568.html)
	//
	// This parameter is required.
	//
	// example:
	//
	// Test123456!
	AccountPassword *string `json:"AccountPassword,omitempty" xml:"AccountPassword,omitempty"`
	// The account type.<br>Valid values:<br>- `normal`: replica set account. For a replica set instance, any value is processed as `normal`.<br>- `cs`: sharded cluster ConfigServer account.<br>- `db`: sharded cluster Shard direct connection account. For a sharded cluster instance, any value other than `cs` is processed as `db`.<br>Default value: `db`.
	//
	// example:
	//
	// db
	CharacterType *string `json:"CharacterType,omitempty" xml:"CharacterType,omitempty"`
	// The instance ID. For a sharded cluster instance, specify the sharded cluster ID. Account operations are executed from the perspective of the ConfigServer.
	//
	// This parameter is required.
	//
	// example:
	//
	// dds-uf6e9433e955xxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The database to which the account belongs (authentication database).<br>
	//
	// Default value: `admin`.<br>Rules: If this parameter is left empty, the request is allowed and the downstream system processes it as `admin`. If this parameter is specified, the length must be 1 to 63 characters. The value cannot contain `/`, `\\`, `.`, spaces, `"`, or `$`, and cannot be `local`. Otherwise, `InvalidParameter` is returned.
	//
	// example:
	//
	// admin
	DBName               *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The list of authorization roles for the account. **Required for replica set and ConfigServer (cs) accounts**. **Prohibited for sharded cluster Shard direct connection (db) accounts**.
	//
	// - **Globally prohibited*	- (returns `InvalidRoles.NotAllowed`): `root`, `clusterAdmin`, `hostManager`, `clusterManager`, `__system`, `__queryableBackup`, and `restore`.
	//
	// - **Can be authorized to any database (including admin)**: `read`, `userAdmin`, and `enableSharding`.
	//
	// - **Can only be authorized to non-admin databases*	- (returns `InvalidRoles.InvalidDB` if authorized to admin): `readWrite`, `dbAdmin`, and `dbOwner`.
	//
	// - **Can only be authorized to the admin database**: `readAnyDatabase`, `readWriteAnyDatabase`, `userAdminAnyDatabase`, `dbAdminAnyDatabase`, `clusterMonitor`, `backup`, and `alibabaCloudAdmin`. (`alibabaCloudAdmin` is supported only for instances with kernel version 4.4 or later. For earlier versions, `InvalidRoles.NotAllowed` is returned.)
	//
	// - **Custom roles**: Roles other than the preceding ones are considered custom roles and must already exist in the instance (you can query them by calling DescribeAccountRoles). Otherwise, `InvalidRoles.NotFound` is returned. `Roles.N.DBName` must be the same as the database where the role is defined. Otherwise, `InvalidRoles.DBNotMatch` is returned.
	//
	// - If `Roles` is missing as a whole, `MissingRoles` is returned. If `Roles.N.Role` and `Roles.N.DBName` do not appear in pairs, `InvalidRoles.Malformed` is returned. If `Roles` is passed for a sharded cluster Shard direct connection account, `InvalidRoles.NotSupported` is returned.
	Roles []*CreateAccountRequestRoles `json:"Roles,omitempty" xml:"Roles,omitempty" type:"Repeated"`
}

func (s CreateAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAccountRequest) GoString() string {
	return s.String()
}

func (s *CreateAccountRequest) GetAccountDescription() *string {
	return s.AccountDescription
}

func (s *CreateAccountRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *CreateAccountRequest) GetAccountPassword() *string {
	return s.AccountPassword
}

func (s *CreateAccountRequest) GetCharacterType() *string {
	return s.CharacterType
}

func (s *CreateAccountRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *CreateAccountRequest) GetDBName() *string {
	return s.DBName
}

func (s *CreateAccountRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateAccountRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateAccountRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateAccountRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateAccountRequest) GetRoles() []*CreateAccountRequestRoles {
	return s.Roles
}

func (s *CreateAccountRequest) SetAccountDescription(v string) *CreateAccountRequest {
	s.AccountDescription = &v
	return s
}

func (s *CreateAccountRequest) SetAccountName(v string) *CreateAccountRequest {
	s.AccountName = &v
	return s
}

func (s *CreateAccountRequest) SetAccountPassword(v string) *CreateAccountRequest {
	s.AccountPassword = &v
	return s
}

func (s *CreateAccountRequest) SetCharacterType(v string) *CreateAccountRequest {
	s.CharacterType = &v
	return s
}

func (s *CreateAccountRequest) SetDBInstanceId(v string) *CreateAccountRequest {
	s.DBInstanceId = &v
	return s
}

func (s *CreateAccountRequest) SetDBName(v string) *CreateAccountRequest {
	s.DBName = &v
	return s
}

func (s *CreateAccountRequest) SetOwnerAccount(v string) *CreateAccountRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateAccountRequest) SetOwnerId(v int64) *CreateAccountRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateAccountRequest) SetResourceOwnerAccount(v string) *CreateAccountRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateAccountRequest) SetResourceOwnerId(v int64) *CreateAccountRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateAccountRequest) SetRoles(v []*CreateAccountRequestRoles) *CreateAccountRequest {
	s.Roles = v
	return s
}

func (s *CreateAccountRequest) Validate() error {
	if s.Roles != nil {
		for _, item := range s.Roles {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateAccountRequestRoles struct {
	// The database to which the account belongs (authentication database).<br>
	//
	// Default value: `admin`.<br>Rules: If this parameter is left empty, the request is allowed and the downstream system processes it as `admin`. If this parameter is specified, the length must be 1 to 63 characters. The value cannot contain `/`, `\\`, `.`, spaces, `"`, or `$`, and cannot be `local`. Otherwise, `InvalidParameter` is returned.
	//
	// example:
	//
	// admin
	DBName *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	// The role name. It can be a built-in role or a custom role created on the instance. For value rules, see [Roles parameter description](https://www.alibabacloud.com/help/en/mongodb/manage-database-accounts-through-the-console).
	//
	// example:
	//
	// read
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
}

func (s CreateAccountRequestRoles) String() string {
	return dara.Prettify(s)
}

func (s CreateAccountRequestRoles) GoString() string {
	return s.String()
}

func (s *CreateAccountRequestRoles) GetDBName() *string {
	return s.DBName
}

func (s *CreateAccountRequestRoles) GetRole() *string {
	return s.Role
}

func (s *CreateAccountRequestRoles) SetDBName(v string) *CreateAccountRequestRoles {
	s.DBName = &v
	return s
}

func (s *CreateAccountRequestRoles) SetRole(v string) *CreateAccountRequestRoles {
	s.Role = &v
	return s
}

func (s *CreateAccountRequestRoles) Validate() error {
	return dara.Validate(s)
}
