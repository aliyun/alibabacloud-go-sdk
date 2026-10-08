// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceHAConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbInstanceId(v string) *ModifyDBInstanceHAConfigRequest
	GetDbInstanceId() *string
	SetHAMode(v string) *ModifyDBInstanceHAConfigRequest
	GetHAMode() *string
	SetOwnerAccount(v string) *ModifyDBInstanceHAConfigRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyDBInstanceHAConfigRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ModifyDBInstanceHAConfigRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyDBInstanceHAConfigRequest
	GetResourceOwnerId() *int64
	SetSyncMode(v string) *ModifyDBInstanceHAConfigRequest
	GetSyncMode() *string
}

type ModifyDBInstanceHAConfigRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk543****
	DbInstanceId *string `json:"DbInstanceId,omitempty" xml:"DbInstanceId,omitempty"`
	// The High-availability Mode. Valid values:
	//
	// - RPO: Data consistency is preferred. The instance ensures data reliability to the greatest extent, which minimizes the amount of data loss. Use RPO mode if you have high requirements for data consistency.
	//
	// - RTO: Instance availability is preferred. The instance recovers services as soon as possible, which maximizes the active time. Use RTO mode if you have high requirements for database uptime.
	//
	// This parameter is required.
	//
	// example:
	//
	// RPO
	HAMode               *string `json:"HAMode,omitempty" xml:"HAMode,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The data replication method. Valid values:
	//
	// - Semi-sync: semi-synchronous replication.
	//
	// - Sync: synchronous replication.
	//
	// - Async: asynchronous replication.
	//
	// <props="china">- Mgr: MySQL Group Replication.
	//
	// This parameter is required.
	//
	// example:
	//
	// Sync
	SyncMode *string `json:"SyncMode,omitempty" xml:"SyncMode,omitempty"`
}

func (s ModifyDBInstanceHAConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceHAConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceHAConfigRequest) GetDbInstanceId() *string {
	return s.DbInstanceId
}

func (s *ModifyDBInstanceHAConfigRequest) GetHAMode() *string {
	return s.HAMode
}

func (s *ModifyDBInstanceHAConfigRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyDBInstanceHAConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyDBInstanceHAConfigRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyDBInstanceHAConfigRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDBInstanceHAConfigRequest) GetSyncMode() *string {
	return s.SyncMode
}

func (s *ModifyDBInstanceHAConfigRequest) SetDbInstanceId(v string) *ModifyDBInstanceHAConfigRequest {
	s.DbInstanceId = &v
	return s
}

func (s *ModifyDBInstanceHAConfigRequest) SetHAMode(v string) *ModifyDBInstanceHAConfigRequest {
	s.HAMode = &v
	return s
}

func (s *ModifyDBInstanceHAConfigRequest) SetOwnerAccount(v string) *ModifyDBInstanceHAConfigRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBInstanceHAConfigRequest) SetOwnerId(v int64) *ModifyDBInstanceHAConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBInstanceHAConfigRequest) SetResourceOwnerAccount(v string) *ModifyDBInstanceHAConfigRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBInstanceHAConfigRequest) SetResourceOwnerId(v int64) *ModifyDBInstanceHAConfigRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBInstanceHAConfigRequest) SetSyncMode(v string) *ModifyDBInstanceHAConfigRequest {
	s.SyncMode = &v
	return s
}

func (s *ModifyDBInstanceHAConfigRequest) Validate() error {
	return dara.Validate(s)
}
