// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCopyDatabaseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *CopyDatabaseRequest
	GetDBInstanceName() *string
	SetDstDBName(v string) *CopyDatabaseRequest
	GetDstDBName() *string
	SetOwnerId(v int64) *CopyDatabaseRequest
	GetOwnerId() *int64
	SetReserveAccount(v int32) *CopyDatabaseRequest
	GetReserveAccount() *int32
	SetResourceGroupId(v string) *CopyDatabaseRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *CopyDatabaseRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CopyDatabaseRequest
	GetResourceOwnerId() *int64
	SetSrcDBName(v string) *CopyDatabaseRequest
	GetSrcDBName() *string
}

type CopyDatabaseRequest struct {
	// The instance name.
	//
	// example:
	//
	// rm-uf6wjk5******
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The destination database name.
	//
	// example:
	//
	// db2***
	DstDBName *string `json:"DstDBName,omitempty" xml:"DstDBName,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The reserved account.
	//
	// example:
	//
	// 1
	ReserveAccount *int32 `json:"ReserveAccount,omitempty" xml:"ReserveAccount,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmy*****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The source database name.
	//
	// example:
	//
	// db1***
	SrcDBName *string `json:"SrcDBName,omitempty" xml:"SrcDBName,omitempty"`
}

func (s CopyDatabaseRequest) String() string {
	return dara.Prettify(s)
}

func (s CopyDatabaseRequest) GoString() string {
	return s.String()
}

func (s *CopyDatabaseRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *CopyDatabaseRequest) GetDstDBName() *string {
	return s.DstDBName
}

func (s *CopyDatabaseRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CopyDatabaseRequest) GetReserveAccount() *int32 {
	return s.ReserveAccount
}

func (s *CopyDatabaseRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CopyDatabaseRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CopyDatabaseRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CopyDatabaseRequest) GetSrcDBName() *string {
	return s.SrcDBName
}

func (s *CopyDatabaseRequest) SetDBInstanceName(v string) *CopyDatabaseRequest {
	s.DBInstanceName = &v
	return s
}

func (s *CopyDatabaseRequest) SetDstDBName(v string) *CopyDatabaseRequest {
	s.DstDBName = &v
	return s
}

func (s *CopyDatabaseRequest) SetOwnerId(v int64) *CopyDatabaseRequest {
	s.OwnerId = &v
	return s
}

func (s *CopyDatabaseRequest) SetReserveAccount(v int32) *CopyDatabaseRequest {
	s.ReserveAccount = &v
	return s
}

func (s *CopyDatabaseRequest) SetResourceGroupId(v string) *CopyDatabaseRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CopyDatabaseRequest) SetResourceOwnerAccount(v string) *CopyDatabaseRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CopyDatabaseRequest) SetResourceOwnerId(v int64) *CopyDatabaseRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CopyDatabaseRequest) SetSrcDBName(v string) *CopyDatabaseRequest {
	s.SrcDBName = &v
	return s
}

func (s *CopyDatabaseRequest) Validate() error {
	return dara.Validate(s)
}
