// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAccountDescriptionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountDescription(v string) *ModifyAccountDescriptionRequest
	GetAccountDescription() *string
	SetAccountName(v string) *ModifyAccountDescriptionRequest
	GetAccountName() *string
	SetCharacterType(v string) *ModifyAccountDescriptionRequest
	GetCharacterType() *string
	SetDBInstanceId(v string) *ModifyAccountDescriptionRequest
	GetDBInstanceId() *string
	SetOwnerAccount(v string) *ModifyAccountDescriptionRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyAccountDescriptionRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ModifyAccountDescriptionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyAccountDescriptionRequest
	GetResourceOwnerId() *int64
}

type ModifyAccountDescriptionRequest struct {
	// The description of the account.
	//
	// 	- The description must start with a letter and cannot start with http:// or https://.
	//
	// 	- The description must be 2 to 256 characters in length, and can contain letters, digits, underscores (_), and hyphens (-).
	//
	// This parameter is required.
	//
	// example:
	//
	// superadmin
	AccountDescription *string `json:"AccountDescription,omitempty" xml:"AccountDescription,omitempty"`
	// The name of the account whose description is to be modified.
	//
	// This parameter is required.
	//
	// example:
	//
	// root
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The type of the account whose description you can modify. Valid values:
	//
	// 	- **db**: shard account
	//
	// 	- **cs**: ConfigServer account
	//
	// 	- **normal*	- (default): replica set account (available)
	//
	// >  You can set this parameter only to **normal**.
	//
	// example:
	//
	// normal
	CharacterType *string `json:"CharacterType,omitempty" xml:"CharacterType,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// dds-bp2356****
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyAccountDescriptionRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAccountDescriptionRequest) GoString() string {
	return s.String()
}

func (s *ModifyAccountDescriptionRequest) GetAccountDescription() *string {
	return s.AccountDescription
}

func (s *ModifyAccountDescriptionRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *ModifyAccountDescriptionRequest) GetCharacterType() *string {
	return s.CharacterType
}

func (s *ModifyAccountDescriptionRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyAccountDescriptionRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyAccountDescriptionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyAccountDescriptionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyAccountDescriptionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyAccountDescriptionRequest) SetAccountDescription(v string) *ModifyAccountDescriptionRequest {
	s.AccountDescription = &v
	return s
}

func (s *ModifyAccountDescriptionRequest) SetAccountName(v string) *ModifyAccountDescriptionRequest {
	s.AccountName = &v
	return s
}

func (s *ModifyAccountDescriptionRequest) SetCharacterType(v string) *ModifyAccountDescriptionRequest {
	s.CharacterType = &v
	return s
}

func (s *ModifyAccountDescriptionRequest) SetDBInstanceId(v string) *ModifyAccountDescriptionRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyAccountDescriptionRequest) SetOwnerAccount(v string) *ModifyAccountDescriptionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyAccountDescriptionRequest) SetOwnerId(v int64) *ModifyAccountDescriptionRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyAccountDescriptionRequest) SetResourceOwnerAccount(v string) *ModifyAccountDescriptionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyAccountDescriptionRequest) SetResourceOwnerId(v int64) *ModifyAccountDescriptionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyAccountDescriptionRequest) Validate() error {
	return dara.Validate(s)
}
