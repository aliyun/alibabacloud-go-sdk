// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceSSLRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *ModifyDBInstanceSSLRequest
	GetDBInstanceId() *string
	SetForceEncryption(v string) *ModifyDBInstanceSSLRequest
	GetForceEncryption() *string
	SetOwnerAccount(v string) *ModifyDBInstanceSSLRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyDBInstanceSSLRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ModifyDBInstanceSSLRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyDBInstanceSSLRequest
	GetResourceOwnerId() *int64
	SetSSLAction(v string) *ModifyDBInstanceSSLRequest
	GetSSLAction() *string
	SetSwitchMode(v string) *ModifyDBInstanceSSLRequest
	GetSwitchMode() *string
}

type ModifyDBInstanceSSLRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// dds-bp2235****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// Specifies whether to forcibly enable SSL encryption for connections. Valid values:
	//
	// - **1**: Forcibly enable SSL encryption.
	//
	// - **0**: Do not forcibly enable SSL encryption.
	//
	// > 	- Forced SSL encryption is supported only for MongoDB 7.0 and 8.0 instances that use cloud disks and meet the following minor engine version requirements:
	//
	// >
	//
	// > 	- - For version 7.0, the minor engine version must be 8.0.13 or later.
	//
	// >
	//
	// > 	- - For version 8.0, the minor engine version must be 9.0.5 or later.
	//
	// 	Warning:
	//
	// After you enable forced SSL encryption, only SSL connections to the instance are allowed.
	//
	// example:
	//
	// 0
	ForceEncryption      *string `json:"ForceEncryption,omitempty" xml:"ForceEncryption,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The operation to perform on the SSL feature. Valid values:
	//
	// - **Open**: Enable SSL encryption.
	//
	// - **Close**: Disable SSL encryption.
	//
	// - **Update**: Update the SSL certificate.
	//
	// This parameter is required.
	//
	// example:
	//
	// Open
	SSLAction *string `json:"SSLAction,omitempty" xml:"SSLAction,omitempty"`
	// The time to modify the SSL configuration of the MongoDB instance. Valid values:
	//
	// - 0: Modify immediately.
	//
	// - 1: Modify within the maintenance window.
	//
	// example:
	//
	// 0
	SwitchMode *string `json:"SwitchMode,omitempty" xml:"SwitchMode,omitempty"`
}

func (s ModifyDBInstanceSSLRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceSSLRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceSSLRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyDBInstanceSSLRequest) GetForceEncryption() *string {
	return s.ForceEncryption
}

func (s *ModifyDBInstanceSSLRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyDBInstanceSSLRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyDBInstanceSSLRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyDBInstanceSSLRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDBInstanceSSLRequest) GetSSLAction() *string {
	return s.SSLAction
}

func (s *ModifyDBInstanceSSLRequest) GetSwitchMode() *string {
	return s.SwitchMode
}

func (s *ModifyDBInstanceSSLRequest) SetDBInstanceId(v string) *ModifyDBInstanceSSLRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBInstanceSSLRequest) SetForceEncryption(v string) *ModifyDBInstanceSSLRequest {
	s.ForceEncryption = &v
	return s
}

func (s *ModifyDBInstanceSSLRequest) SetOwnerAccount(v string) *ModifyDBInstanceSSLRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBInstanceSSLRequest) SetOwnerId(v int64) *ModifyDBInstanceSSLRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBInstanceSSLRequest) SetResourceOwnerAccount(v string) *ModifyDBInstanceSSLRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBInstanceSSLRequest) SetResourceOwnerId(v int64) *ModifyDBInstanceSSLRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBInstanceSSLRequest) SetSSLAction(v string) *ModifyDBInstanceSSLRequest {
	s.SSLAction = &v
	return s
}

func (s *ModifyDBInstanceSSLRequest) SetSwitchMode(v string) *ModifyDBInstanceSSLRequest {
	s.SwitchMode = &v
	return s
}

func (s *ModifyDBInstanceSSLRequest) Validate() error {
	return dara.Validate(s)
}
