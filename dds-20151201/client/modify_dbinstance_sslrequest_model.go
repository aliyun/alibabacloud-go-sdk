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
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// dds-bp2235****
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The operation on the SSL feature. Valid values:
	//
	// 	- **Open**: enables SSL encryption.
	//
	// 	- **Close**: disables SSL encryption.
	//
	// 	- **Update**: updates the SSL certificate.
	//
	// This parameter is required.
	//
	// example:
	//
	// Open
	SSLAction  *string `json:"SSLAction,omitempty" xml:"SSLAction,omitempty"`
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
