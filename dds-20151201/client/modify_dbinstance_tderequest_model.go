// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceTDERequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *ModifyDBInstanceTDERequest
	GetDBInstanceId() *string
	SetEncryptionKey(v string) *ModifyDBInstanceTDERequest
	GetEncryptionKey() *string
	SetEncryptorName(v string) *ModifyDBInstanceTDERequest
	GetEncryptorName() *string
	SetOwnerAccount(v string) *ModifyDBInstanceTDERequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyDBInstanceTDERequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ModifyDBInstanceTDERequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyDBInstanceTDERequest
	GetResourceOwnerId() *int64
	SetRoleARN(v string) *ModifyDBInstanceTDERequest
	GetRoleARN() *string
	SetSwitchMode(v string) *ModifyDBInstanceTDERequest
	GetSwitchMode() *string
	SetTDEStatus(v string) *ModifyDBInstanceTDERequest
	GetTDEStatus() *string
}

type ModifyDBInstanceTDERequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// dds-bpxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The ID of the custom key.
	//
	// example:
	//
	// 749c1df7-xxxx-xxxx-xxxx-xxxxxxxxxxxx
	EncryptionKey *string `json:"EncryptionKey,omitempty" xml:"EncryptionKey,omitempty"`
	// The encryption method. Set the value to **aes-256-cbc**.
	//
	// > This parameter is valid only when the **TEDStatus*	- parameter is set to **enabled**.
	//
	// example:
	//
	// aes-256-cbc
	EncryptorName        *string `json:"EncryptorName,omitempty" xml:"EncryptorName,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The Alibaba Cloud Resource Name (ARN) of the specified Resource Access Management (RAM) role. The ARN is displayed in the `acs:ram::$accountID:role/$roleName` format.
	//
	// > 	- `$accountID`: specifies the ID of the Alibaba Cloud account. To view the account ID, log on to the Alibaba Cloud Management Console, move your pointer over your profile picture in the upper-right corner, and then click Security Settings.
	//
	// > 	- `$roleName`: specifies the name of the RAM role. To view the RAM role name, log on to the RAM console. In the left-side navigation pane, choose Identities > Roles. On the Roles page, view the name of the RAM role.
	//
	// example:
	//
	// acs:ram::123456789012****:role/adminrole
	RoleARN    *string `json:"RoleARN,omitempty" xml:"RoleARN,omitempty"`
	SwitchMode *string `json:"SwitchMode,omitempty" xml:"SwitchMode,omitempty"`
	// The TDE status. When the value of this parameter is set to **Enabled**, TDE is enabled.
	//
	// > You cannot disable TDE after it is enabled. Proceed with caution.
	//
	// This parameter is required.
	//
	// example:
	//
	// enabled
	TDEStatus *string `json:"TDEStatus,omitempty" xml:"TDEStatus,omitempty"`
}

func (s ModifyDBInstanceTDERequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceTDERequest) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceTDERequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyDBInstanceTDERequest) GetEncryptionKey() *string {
	return s.EncryptionKey
}

func (s *ModifyDBInstanceTDERequest) GetEncryptorName() *string {
	return s.EncryptorName
}

func (s *ModifyDBInstanceTDERequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyDBInstanceTDERequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyDBInstanceTDERequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyDBInstanceTDERequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDBInstanceTDERequest) GetRoleARN() *string {
	return s.RoleARN
}

func (s *ModifyDBInstanceTDERequest) GetSwitchMode() *string {
	return s.SwitchMode
}

func (s *ModifyDBInstanceTDERequest) GetTDEStatus() *string {
	return s.TDEStatus
}

func (s *ModifyDBInstanceTDERequest) SetDBInstanceId(v string) *ModifyDBInstanceTDERequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBInstanceTDERequest) SetEncryptionKey(v string) *ModifyDBInstanceTDERequest {
	s.EncryptionKey = &v
	return s
}

func (s *ModifyDBInstanceTDERequest) SetEncryptorName(v string) *ModifyDBInstanceTDERequest {
	s.EncryptorName = &v
	return s
}

func (s *ModifyDBInstanceTDERequest) SetOwnerAccount(v string) *ModifyDBInstanceTDERequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBInstanceTDERequest) SetOwnerId(v int64) *ModifyDBInstanceTDERequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBInstanceTDERequest) SetResourceOwnerAccount(v string) *ModifyDBInstanceTDERequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBInstanceTDERequest) SetResourceOwnerId(v int64) *ModifyDBInstanceTDERequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBInstanceTDERequest) SetRoleARN(v string) *ModifyDBInstanceTDERequest {
	s.RoleARN = &v
	return s
}

func (s *ModifyDBInstanceTDERequest) SetSwitchMode(v string) *ModifyDBInstanceTDERequest {
	s.SwitchMode = &v
	return s
}

func (s *ModifyDBInstanceTDERequest) SetTDEStatus(v string) *ModifyDBInstanceTDERequest {
	s.TDEStatus = &v
	return s
}

func (s *ModifyDBInstanceTDERequest) Validate() error {
	return dara.Validate(s)
}
