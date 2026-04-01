// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBDescriptionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBDescription(v string) *ModifyDBDescriptionRequest
	GetDBDescription() *string
	SetDBInstanceId(v string) *ModifyDBDescriptionRequest
	GetDBInstanceId() *string
	SetDBName(v string) *ModifyDBDescriptionRequest
	GetDBName() *string
	SetOwnerAccount(v string) *ModifyDBDescriptionRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyDBDescriptionRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ModifyDBDescriptionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyDBDescriptionRequest
	GetResourceOwnerId() *int64
}

type ModifyDBDescriptionRequest struct {
	// The description of the database.
	//
	// This parameter is required.
	//
	// example:
	//
	// Test database A
	DBDescription *string `json:"DBDescription,omitempty" xml:"DBDescription,omitempty"`
	// The instance ID. You can call the DescribeDBInstances operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5xxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The name of the database.
	//
	// This parameter is required.
	//
	// example:
	//
	// testDB01
	DBName               *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyDBDescriptionRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBDescriptionRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBDescriptionRequest) GetDBDescription() *string {
	return s.DBDescription
}

func (s *ModifyDBDescriptionRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyDBDescriptionRequest) GetDBName() *string {
	return s.DBName
}

func (s *ModifyDBDescriptionRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyDBDescriptionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyDBDescriptionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyDBDescriptionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDBDescriptionRequest) SetDBDescription(v string) *ModifyDBDescriptionRequest {
	s.DBDescription = &v
	return s
}

func (s *ModifyDBDescriptionRequest) SetDBInstanceId(v string) *ModifyDBDescriptionRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBDescriptionRequest) SetDBName(v string) *ModifyDBDescriptionRequest {
	s.DBName = &v
	return s
}

func (s *ModifyDBDescriptionRequest) SetOwnerAccount(v string) *ModifyDBDescriptionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBDescriptionRequest) SetOwnerId(v int64) *ModifyDBDescriptionRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBDescriptionRequest) SetResourceOwnerAccount(v string) *ModifyDBDescriptionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBDescriptionRequest) SetResourceOwnerId(v int64) *ModifyDBDescriptionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBDescriptionRequest) Validate() error {
	return dara.Validate(s)
}
