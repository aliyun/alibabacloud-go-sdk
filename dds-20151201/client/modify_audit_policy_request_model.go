// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAuditPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuditLogSwitchSource(v string) *ModifyAuditPolicyRequest
	GetAuditLogSwitchSource() *string
	SetAuditStatus(v string) *ModifyAuditPolicyRequest
	GetAuditStatus() *string
	SetDBInstanceId(v string) *ModifyAuditPolicyRequest
	GetDBInstanceId() *string
	SetOwnerAccount(v string) *ModifyAuditPolicyRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyAuditPolicyRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ModifyAuditPolicyRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyAuditPolicyRequest
	GetResourceOwnerId() *int64
	SetServiceType(v string) *ModifyAuditPolicyRequest
	GetServiceType() *string
	SetStoragePeriod(v int32) *ModifyAuditPolicyRequest
	GetStoragePeriod() *int32
}

type ModifyAuditPolicyRequest struct {
	// The request source for the audit log feature. Set the value to **Console**.
	//
	// example:
	//
	// Console
	AuditLogSwitchSource *string `json:"AuditLogSwitchSource,omitempty" xml:"AuditLogSwitchSource,omitempty"`
	// Specifies whether to enable the audit log feature. Valid values:
	//
	// 	- **enable**
	//
	// 	- **disabled**
	//
	// This parameter is required.
	//
	// example:
	//
	// enable
	AuditStatus *string `json:"AuditStatus,omitempty" xml:"AuditStatus,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// dds-bp1785659e3f****
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The type of the audit log feature. Valid values:
	//
	// 	- **Trail**: free trial edition.
	//
	// 	- **Standard**: official edition.
	//
	// > The default value is **Trail**. Starting from January 6, 2022, the official edition of the audit log feature has been launched in all regions, and the free trial edition of the feature can no longer be applied for. We recommend that you set this parameter to **Standard**.
	//
	// example:
	//
	// Standard
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	// The log retention period. Valid values: 1 to 365 days. Default value: 30 days.
	//
	// example:
	//
	// 30
	StoragePeriod *int32 `json:"StoragePeriod,omitempty" xml:"StoragePeriod,omitempty"`
}

func (s ModifyAuditPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAuditPolicyRequest) GoString() string {
	return s.String()
}

func (s *ModifyAuditPolicyRequest) GetAuditLogSwitchSource() *string {
	return s.AuditLogSwitchSource
}

func (s *ModifyAuditPolicyRequest) GetAuditStatus() *string {
	return s.AuditStatus
}

func (s *ModifyAuditPolicyRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyAuditPolicyRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyAuditPolicyRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyAuditPolicyRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyAuditPolicyRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyAuditPolicyRequest) GetServiceType() *string {
	return s.ServiceType
}

func (s *ModifyAuditPolicyRequest) GetStoragePeriod() *int32 {
	return s.StoragePeriod
}

func (s *ModifyAuditPolicyRequest) SetAuditLogSwitchSource(v string) *ModifyAuditPolicyRequest {
	s.AuditLogSwitchSource = &v
	return s
}

func (s *ModifyAuditPolicyRequest) SetAuditStatus(v string) *ModifyAuditPolicyRequest {
	s.AuditStatus = &v
	return s
}

func (s *ModifyAuditPolicyRequest) SetDBInstanceId(v string) *ModifyAuditPolicyRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyAuditPolicyRequest) SetOwnerAccount(v string) *ModifyAuditPolicyRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyAuditPolicyRequest) SetOwnerId(v int64) *ModifyAuditPolicyRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyAuditPolicyRequest) SetResourceOwnerAccount(v string) *ModifyAuditPolicyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyAuditPolicyRequest) SetResourceOwnerId(v int64) *ModifyAuditPolicyRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyAuditPolicyRequest) SetServiceType(v string) *ModifyAuditPolicyRequest {
	s.ServiceType = &v
	return s
}

func (s *ModifyAuditPolicyRequest) SetStoragePeriod(v int32) *ModifyAuditPolicyRequest {
	s.StoragePeriod = &v
	return s
}

func (s *ModifyAuditPolicyRequest) Validate() error {
	return dara.Validate(s)
}
