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
	// Custom keys are supported only in the following regions. In other regions, the default key is used.
	//
	// - Singapore (ap-southeast-1)
	//
	// - Hangzhou (cn-hangzhou)
	//
	// - Shanghai (cn-shanghai)
	//
	// - Beijing (cn-beijing)
	//
	// - Shenzhen (cn-shenzhen)
	//
	// - Hong Kong (cn-hongkong)
	//
	// - Malaysia (ap-southeast-3)
	//
	// example:
	//
	// 749c1df7-xxxx-xxxx-xxxx-xxxxxxxxxxxx
	EncryptionKey *string `json:"EncryptionKey,omitempty" xml:"EncryptionKey,omitempty"`
	// The encryption method. Set the value to **aes-256-cbc**.
	//
	// > This parameter is available only when **TDEStatus*	- is set to **enabled**.
	//
	// example:
	//
	// aes-256-cbc
	EncryptorName        *string `json:"EncryptorName,omitempty" xml:"EncryptorName,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The Alibaba Cloud Resource Name (ARN) of the RAM role. The format is `acs:ram::$accountID:role/$roleName `.
	//
	// > - `$accountID`: The ID of your Alibaba Cloud account. To view the ID, log on to the Alibaba Cloud Management Console, move the pointer over your profile picture in the upper-right corner, and then click Security Settings.
	//
	// >
	//
	// > - `$roleName`: The name of the RAM role. To view the name, log on to the RAM console, click RAM Role Management in the navigation pane on the left, and then view the role name in the RAM Role Name list.
	//
	// example:
	//
	// acs:ram::123456789012****:role/adminrole
	RoleARN *string `json:"RoleARN,omitempty" xml:"RoleARN,omitempty"`
	// Specifies when to enable TDE. Valid values:
	//
	// - 0: Enables TDE immediately.
	//
	// - 1: Enables TDE during the maintenance window.
	//
	// example:
	//
	// 0
	SwitchMode *string `json:"SwitchMode,omitempty" xml:"SwitchMode,omitempty"`
	// The TDE status. Set the value to **enabled*	- to enable TDE.
	//
	// > You cannot disable TDE after you enable it. Enable this feature with caution.
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
