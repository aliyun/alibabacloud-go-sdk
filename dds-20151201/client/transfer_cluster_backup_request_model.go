// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTransferClusterBackupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *TransferClusterBackupRequest
	GetDBInstanceId() *string
	SetOwnerAccount(v string) *TransferClusterBackupRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *TransferClusterBackupRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *TransferClusterBackupRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *TransferClusterBackupRequest
	GetResourceOwnerId() *int64
}

type TransferClusterBackupRequest struct {
	// The instance ID.
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
}

func (s TransferClusterBackupRequest) String() string {
	return dara.Prettify(s)
}

func (s TransferClusterBackupRequest) GoString() string {
	return s.String()
}

func (s *TransferClusterBackupRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *TransferClusterBackupRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *TransferClusterBackupRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *TransferClusterBackupRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *TransferClusterBackupRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *TransferClusterBackupRequest) SetDBInstanceId(v string) *TransferClusterBackupRequest {
	s.DBInstanceId = &v
	return s
}

func (s *TransferClusterBackupRequest) SetOwnerAccount(v string) *TransferClusterBackupRequest {
	s.OwnerAccount = &v
	return s
}

func (s *TransferClusterBackupRequest) SetOwnerId(v int64) *TransferClusterBackupRequest {
	s.OwnerId = &v
	return s
}

func (s *TransferClusterBackupRequest) SetResourceOwnerAccount(v string) *TransferClusterBackupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *TransferClusterBackupRequest) SetResourceOwnerId(v int64) *TransferClusterBackupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *TransferClusterBackupRequest) Validate() error {
	return dara.Validate(s)
}
