// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBClusterAuditLogCollectorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCollectorStatus(v string) *ModifyDBClusterAuditLogCollectorRequest
	GetCollectorStatus() *string
	SetDBClusterId(v string) *ModifyDBClusterAuditLogCollectorRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *ModifyDBClusterAuditLogCollectorRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyDBClusterAuditLogCollectorRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ModifyDBClusterAuditLogCollectorRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyDBClusterAuditLogCollectorRequest
	GetResourceOwnerId() *int64
}

type ModifyDBClusterAuditLogCollectorRequest struct {
	// Specifies whether to enable or disable SQL collector. Valid values:
	//
	// 	- Enable
	//
	// 	- Disable
	//
	// This parameter is required.
	//
	// example:
	//
	// Enable
	CollectorStatus *string `json:"CollectorStatus,omitempty" xml:"CollectorStatus,omitempty"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-***************
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyDBClusterAuditLogCollectorRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBClusterAuditLogCollectorRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterAuditLogCollectorRequest) GetCollectorStatus() *string {
	return s.CollectorStatus
}

func (s *ModifyDBClusterAuditLogCollectorRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyDBClusterAuditLogCollectorRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyDBClusterAuditLogCollectorRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyDBClusterAuditLogCollectorRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyDBClusterAuditLogCollectorRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDBClusterAuditLogCollectorRequest) SetCollectorStatus(v string) *ModifyDBClusterAuditLogCollectorRequest {
	s.CollectorStatus = &v
	return s
}

func (s *ModifyDBClusterAuditLogCollectorRequest) SetDBClusterId(v string) *ModifyDBClusterAuditLogCollectorRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBClusterAuditLogCollectorRequest) SetOwnerAccount(v string) *ModifyDBClusterAuditLogCollectorRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBClusterAuditLogCollectorRequest) SetOwnerId(v int64) *ModifyDBClusterAuditLogCollectorRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBClusterAuditLogCollectorRequest) SetResourceOwnerAccount(v string) *ModifyDBClusterAuditLogCollectorRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBClusterAuditLogCollectorRequest) SetResourceOwnerId(v int64) *ModifyDBClusterAuditLogCollectorRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBClusterAuditLogCollectorRequest) Validate() error {
	return dara.Validate(s)
}
