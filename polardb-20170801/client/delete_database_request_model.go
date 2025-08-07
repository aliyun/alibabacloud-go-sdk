// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDatabaseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DeleteDatabaseRequest
	GetDBClusterId() *string
	SetDBName(v string) *DeleteDatabaseRequest
	GetDBName() *string
	SetOwnerAccount(v string) *DeleteDatabaseRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteDatabaseRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DeleteDatabaseRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteDatabaseRequest
	GetResourceOwnerId() *int64
}

type DeleteDatabaseRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-*************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The name of the database.
	//
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

func (s DeleteDatabaseRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDatabaseRequest) GoString() string {
	return s.String()
}

func (s *DeleteDatabaseRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DeleteDatabaseRequest) GetDBName() *string {
	return s.DBName
}

func (s *DeleteDatabaseRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteDatabaseRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteDatabaseRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteDatabaseRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteDatabaseRequest) SetDBClusterId(v string) *DeleteDatabaseRequest {
	s.DBClusterId = &v
	return s
}

func (s *DeleteDatabaseRequest) SetDBName(v string) *DeleteDatabaseRequest {
	s.DBName = &v
	return s
}

func (s *DeleteDatabaseRequest) SetOwnerAccount(v string) *DeleteDatabaseRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteDatabaseRequest) SetOwnerId(v int64) *DeleteDatabaseRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteDatabaseRequest) SetResourceOwnerAccount(v string) *DeleteDatabaseRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteDatabaseRequest) SetResourceOwnerId(v int64) *DeleteDatabaseRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteDatabaseRequest) Validate() error {
	return dara.Validate(s)
}
