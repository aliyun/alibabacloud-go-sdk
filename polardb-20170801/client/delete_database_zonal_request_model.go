// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDatabaseZonalRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DeleteDatabaseZonalRequest
	GetDBClusterId() *string
	SetDBName(v string) *DeleteDatabaseZonalRequest
	GetDBName() *string
	SetOwnerAccount(v string) *DeleteDatabaseZonalRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteDatabaseZonalRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DeleteDatabaseZonalRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteDatabaseZonalRequest
	GetResourceOwnerId() *int64
}

type DeleteDatabaseZonalRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pc-*************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// testdb
	DBName               *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteDatabaseZonalRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDatabaseZonalRequest) GoString() string {
	return s.String()
}

func (s *DeleteDatabaseZonalRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DeleteDatabaseZonalRequest) GetDBName() *string {
	return s.DBName
}

func (s *DeleteDatabaseZonalRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteDatabaseZonalRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteDatabaseZonalRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteDatabaseZonalRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteDatabaseZonalRequest) SetDBClusterId(v string) *DeleteDatabaseZonalRequest {
	s.DBClusterId = &v
	return s
}

func (s *DeleteDatabaseZonalRequest) SetDBName(v string) *DeleteDatabaseZonalRequest {
	s.DBName = &v
	return s
}

func (s *DeleteDatabaseZonalRequest) SetOwnerAccount(v string) *DeleteDatabaseZonalRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteDatabaseZonalRequest) SetOwnerId(v int64) *DeleteDatabaseZonalRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteDatabaseZonalRequest) SetResourceOwnerAccount(v string) *DeleteDatabaseZonalRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteDatabaseZonalRequest) SetResourceOwnerId(v int64) *DeleteDatabaseZonalRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteDatabaseZonalRequest) Validate() error {
	return dara.Validate(s)
}
