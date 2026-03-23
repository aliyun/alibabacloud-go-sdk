// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDatabaseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCharacterSetName(v string) *CreateDatabaseRequest
	GetCharacterSetName() *string
	SetDBDescription(v string) *CreateDatabaseRequest
	GetDBDescription() *string
	SetDBInstanceId(v string) *CreateDatabaseRequest
	GetDBInstanceId() *string
	SetDBName(v string) *CreateDatabaseRequest
	GetDBName() *string
	SetOwnerAccount(v string) *CreateDatabaseRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateDatabaseRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CreateDatabaseRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateDatabaseRequest
	GetResourceOwnerId() *int64
}

type CreateDatabaseRequest struct {
	// This parameter is required.
	CharacterSetName *string `json:"CharacterSetName,omitempty" xml:"CharacterSetName,omitempty"`
	DBDescription    *string `json:"DBDescription,omitempty" xml:"DBDescription,omitempty"`
	// This parameter is required.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// This parameter is required.
	DBName               *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CreateDatabaseRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDatabaseRequest) GoString() string {
	return s.String()
}

func (s *CreateDatabaseRequest) GetCharacterSetName() *string {
	return s.CharacterSetName
}

func (s *CreateDatabaseRequest) GetDBDescription() *string {
	return s.DBDescription
}

func (s *CreateDatabaseRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *CreateDatabaseRequest) GetDBName() *string {
	return s.DBName
}

func (s *CreateDatabaseRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateDatabaseRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateDatabaseRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateDatabaseRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateDatabaseRequest) SetCharacterSetName(v string) *CreateDatabaseRequest {
	s.CharacterSetName = &v
	return s
}

func (s *CreateDatabaseRequest) SetDBDescription(v string) *CreateDatabaseRequest {
	s.DBDescription = &v
	return s
}

func (s *CreateDatabaseRequest) SetDBInstanceId(v string) *CreateDatabaseRequest {
	s.DBInstanceId = &v
	return s
}

func (s *CreateDatabaseRequest) SetDBName(v string) *CreateDatabaseRequest {
	s.DBName = &v
	return s
}

func (s *CreateDatabaseRequest) SetOwnerAccount(v string) *CreateDatabaseRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateDatabaseRequest) SetOwnerId(v int64) *CreateDatabaseRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateDatabaseRequest) SetResourceOwnerAccount(v string) *CreateDatabaseRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateDatabaseRequest) SetResourceOwnerId(v int64) *CreateDatabaseRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateDatabaseRequest) Validate() error {
	return dara.Validate(s)
}
