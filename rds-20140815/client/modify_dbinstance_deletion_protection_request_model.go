// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceDeletionProtectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ModifyDBInstanceDeletionProtectionRequest
	GetClientToken() *string
	SetDBInstanceId(v string) *ModifyDBInstanceDeletionProtectionRequest
	GetDBInstanceId() *string
	SetDeletionProtection(v bool) *ModifyDBInstanceDeletionProtectionRequest
	GetDeletionProtection() *bool
	SetOwnerAccount(v string) *ModifyDBInstanceDeletionProtectionRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyDBInstanceDeletionProtectionRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ModifyDBInstanceDeletionProtectionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyDBInstanceDeletionProtectionRequest
	GetResourceOwnerId() *int64
}

type ModifyDBInstanceDeletionProtectionRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the generated token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCzxxxxxxxxxx
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The instance ID. You can call the DescribeDBInstances operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5xxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// Specifies whether to enable the release protection feature for the read-only instance. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	DeletionProtection   *bool   `json:"DeletionProtection,omitempty" xml:"DeletionProtection,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyDBInstanceDeletionProtectionRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceDeletionProtectionRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceDeletionProtectionRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyDBInstanceDeletionProtectionRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyDBInstanceDeletionProtectionRequest) GetDeletionProtection() *bool {
	return s.DeletionProtection
}

func (s *ModifyDBInstanceDeletionProtectionRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyDBInstanceDeletionProtectionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyDBInstanceDeletionProtectionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyDBInstanceDeletionProtectionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDBInstanceDeletionProtectionRequest) SetClientToken(v string) *ModifyDBInstanceDeletionProtectionRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyDBInstanceDeletionProtectionRequest) SetDBInstanceId(v string) *ModifyDBInstanceDeletionProtectionRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBInstanceDeletionProtectionRequest) SetDeletionProtection(v bool) *ModifyDBInstanceDeletionProtectionRequest {
	s.DeletionProtection = &v
	return s
}

func (s *ModifyDBInstanceDeletionProtectionRequest) SetOwnerAccount(v string) *ModifyDBInstanceDeletionProtectionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBInstanceDeletionProtectionRequest) SetOwnerId(v int64) *ModifyDBInstanceDeletionProtectionRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBInstanceDeletionProtectionRequest) SetResourceOwnerAccount(v string) *ModifyDBInstanceDeletionProtectionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBInstanceDeletionProtectionRequest) SetResourceOwnerId(v int64) *ModifyDBInstanceDeletionProtectionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBInstanceDeletionProtectionRequest) Validate() error {
	return dara.Validate(s)
}
