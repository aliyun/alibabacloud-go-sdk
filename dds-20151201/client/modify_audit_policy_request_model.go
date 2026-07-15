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
	SetHotStoragePeriod(v int32) *ModifyAuditPolicyRequest
	GetHotStoragePeriod() *int32
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
	// The source of the request. Set this parameter to **Console**.
	//
	// example:
	//
	// Console
	AuditLogSwitchSource *string `json:"AuditLogSwitchSource,omitempty" xml:"AuditLogSwitchSource,omitempty"`
	// The status of the audit log. Valid values:
	//
	// - **enable**: Enables the audit log feature.
	//
	// - **disabled**: Disables the audit log feature.
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
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// This parameter is effective only for the **V2_Standard*	- (DAS Enterprise Edition (NoSQL Compatible) audit log) edition. It specifies the hot storage duration for the audit log. Valid values: 0 to 7. Unit: days.
	//
	// example:
	//
	// 7
	HotStoragePeriod     *int32  `json:"HotStoragePeriod,omitempty" xml:"HotStoragePeriod,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The edition of the audit log. Valid values:
	//
	// - **Trial**: Trial Edition.
	//
	// - **Standard**: Standard Edition.
	//
	// - **V2_Standard**: DAS Enterprise Edition (NoSQL Compatible) audit log.
	//
	// > 	- The default value of this parameter is **Trial**. Starting from January 6, 2022, the Standard edition is being rolled out across regions, and new applications for the Trial edition are no longer accepted.
	//
	// >
	//
	// > 	- Starting from February 2026, the DAS Enterprise Edition (NoSQL Compatible) audit log will be rolled out across regions, and new applications for the Standard edition will no longer be accepted.
	//
	// example:
	//
	// Standard
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	// - For the **Standard*	- edition, this parameter specifies the retention period for the audit log. Valid values: 1 to 365. The default value is 30. Unit: days.
	//
	// - For the **V2_Standard*	- (DAS Enterprise Edition (NoSQL Compatible) audit log) edition, this parameter specifies the cold storage duration for the audit log. Valid values: 30, 180, 365, 1095, and 1825. Unit: days.
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

func (s *ModifyAuditPolicyRequest) GetHotStoragePeriod() *int32 {
	return s.HotStoragePeriod
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

func (s *ModifyAuditPolicyRequest) SetHotStoragePeriod(v int32) *ModifyAuditPolicyRequest {
	s.HotStoragePeriod = &v
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
