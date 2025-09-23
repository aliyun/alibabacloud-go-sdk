// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBClusterMigrationEndpointRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *ModifyDBClusterMigrationEndpointRequest
	GetDBClusterId() *string
	SetMigrationConfig(v string) *ModifyDBClusterMigrationEndpointRequest
	GetMigrationConfig() *string
	SetOwnerAccount(v string) *ModifyDBClusterMigrationEndpointRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyDBClusterMigrationEndpointRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ModifyDBClusterMigrationEndpointRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyDBClusterMigrationEndpointRequest
	GetResourceOwnerId() *int64
}

type ModifyDBClusterMigrationEndpointRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pc-bp1w68hcomx235jaz
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// [{\\"dtsJobId\\":\\"di4w4d9r12m38mb\\",\\"endpoint\\":\\"src\\",\\"endpointInstanceId\\":\\"pc-k2j96w169uhu868l8\\"}]
	MigrationConfig      *string `json:"MigrationConfig,omitempty" xml:"MigrationConfig,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyDBClusterMigrationEndpointRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBClusterMigrationEndpointRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterMigrationEndpointRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyDBClusterMigrationEndpointRequest) GetMigrationConfig() *string {
	return s.MigrationConfig
}

func (s *ModifyDBClusterMigrationEndpointRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyDBClusterMigrationEndpointRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyDBClusterMigrationEndpointRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyDBClusterMigrationEndpointRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDBClusterMigrationEndpointRequest) SetDBClusterId(v string) *ModifyDBClusterMigrationEndpointRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBClusterMigrationEndpointRequest) SetMigrationConfig(v string) *ModifyDBClusterMigrationEndpointRequest {
	s.MigrationConfig = &v
	return s
}

func (s *ModifyDBClusterMigrationEndpointRequest) SetOwnerAccount(v string) *ModifyDBClusterMigrationEndpointRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBClusterMigrationEndpointRequest) SetOwnerId(v int64) *ModifyDBClusterMigrationEndpointRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBClusterMigrationEndpointRequest) SetResourceOwnerAccount(v string) *ModifyDBClusterMigrationEndpointRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBClusterMigrationEndpointRequest) SetResourceOwnerId(v int64) *ModifyDBClusterMigrationEndpointRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBClusterMigrationEndpointRequest) Validate() error {
	return dara.Validate(s)
}
