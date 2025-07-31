// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *ModifyDBInstanceAttributeRequest
	GetDBInstanceId() *string
	SetDBInstanceReleaseProtection(v bool) *ModifyDBInstanceAttributeRequest
	GetDBInstanceReleaseProtection() *bool
	SetOwnerAccount(v string) *ModifyDBInstanceAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyDBInstanceAttributeRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ModifyDBInstanceAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyDBInstanceAttributeRequest
	GetResourceOwnerId() *int64
}

type ModifyDBInstanceAttributeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// dds-7xv0912d85924194
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// example:
	//
	// false
	DBInstanceReleaseProtection *bool   `json:"DBInstanceReleaseProtection,omitempty" xml:"DBInstanceReleaseProtection,omitempty"`
	OwnerAccount                *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount        *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId             *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyDBInstanceAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceAttributeRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyDBInstanceAttributeRequest) GetDBInstanceReleaseProtection() *bool {
	return s.DBInstanceReleaseProtection
}

func (s *ModifyDBInstanceAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyDBInstanceAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyDBInstanceAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyDBInstanceAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDBInstanceAttributeRequest) SetDBInstanceId(v string) *ModifyDBInstanceAttributeRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBInstanceAttributeRequest) SetDBInstanceReleaseProtection(v bool) *ModifyDBInstanceAttributeRequest {
	s.DBInstanceReleaseProtection = &v
	return s
}

func (s *ModifyDBInstanceAttributeRequest) SetOwnerAccount(v string) *ModifyDBInstanceAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBInstanceAttributeRequest) SetOwnerId(v int64) *ModifyDBInstanceAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBInstanceAttributeRequest) SetResourceOwnerAccount(v string) *ModifyDBInstanceAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBInstanceAttributeRequest) SetResourceOwnerId(v int64) *ModifyDBInstanceAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBInstanceAttributeRequest) Validate() error {
	return dara.Validate(s)
}
