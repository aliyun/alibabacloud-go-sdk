// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceDescriptionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceDescription(v string) *ModifyDBInstanceDescriptionRequest
	GetDBInstanceDescription() *string
	SetDBInstanceId(v string) *ModifyDBInstanceDescriptionRequest
	GetDBInstanceId() *string
	SetOwnerAccount(v string) *ModifyDBInstanceDescriptionRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyDBInstanceDescriptionRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ModifyDBInstanceDescriptionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyDBInstanceDescriptionRequest
	GetResourceOwnerId() *int64
}

type ModifyDBInstanceDescriptionRequest struct {
	// The name of the instance.
	//
	// > The name must be 2 to 64 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// Instance in Alibaba Cloud test environment
	DBInstanceDescription *string `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5****
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyDBInstanceDescriptionRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceDescriptionRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceDescriptionRequest) GetDBInstanceDescription() *string {
	return s.DBInstanceDescription
}

func (s *ModifyDBInstanceDescriptionRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyDBInstanceDescriptionRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyDBInstanceDescriptionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyDBInstanceDescriptionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyDBInstanceDescriptionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDBInstanceDescriptionRequest) SetDBInstanceDescription(v string) *ModifyDBInstanceDescriptionRequest {
	s.DBInstanceDescription = &v
	return s
}

func (s *ModifyDBInstanceDescriptionRequest) SetDBInstanceId(v string) *ModifyDBInstanceDescriptionRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBInstanceDescriptionRequest) SetOwnerAccount(v string) *ModifyDBInstanceDescriptionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBInstanceDescriptionRequest) SetOwnerId(v int64) *ModifyDBInstanceDescriptionRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBInstanceDescriptionRequest) SetResourceOwnerAccount(v string) *ModifyDBInstanceDescriptionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBInstanceDescriptionRequest) SetResourceOwnerId(v int64) *ModifyDBInstanceDescriptionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBInstanceDescriptionRequest) Validate() error {
	return dara.Validate(s)
}
