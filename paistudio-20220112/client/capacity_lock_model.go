// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCapacityLock interface {
	dara.Model
	String() string
	GoString() string
	SetAvailableCount(v int32) *CapacityLock
	GetAvailableCount() *int32
	SetCrsReservationId(v string) *CapacityLock
	GetCrsReservationId() *string
	SetDescription(v string) *CapacityLock
	GetDescription() *string
	SetExpireTime(v string) *CapacityLock
	GetExpireTime() *string
	SetGmtCreated(v string) *CapacityLock
	GetGmtCreated() *string
	SetGmtModified(v string) *CapacityLock
	GetGmtModified() *string
	SetId(v string) *CapacityLock
	GetId() *string
	SetInstanceType(v string) *CapacityLock
	GetInstanceType() *string
	SetLastReconcileAttemptTime(v string) *CapacityLock
	GetLastReconcileAttemptTime() *string
	SetLastSyncTime(v string) *CapacityLock
	GetLastSyncTime() *string
	SetLockProvider(v string) *CapacityLock
	GetLockProvider() *string
	SetLockedCount(v int32) *CapacityLock
	GetLockedCount() *int32
	SetOperator(v string) *CapacityLock
	GetOperator() *string
	SetPaymentType(v string) *CapacityLock
	GetPaymentType() *string
	SetPrivatePoolId(v string) *CapacityLock
	GetPrivatePoolId() *string
	SetRequestedCount(v int32) *CapacityLock
	GetRequestedCount() *int32
	SetStatus(v string) *CapacityLock
	GetStatus() *string
	SetTenantId(v string) *CapacityLock
	GetTenantId() *string
	SetUsedCount(v int32) *CapacityLock
	GetUsedCount() *int32
	SetZoneId(v string) *CapacityLock
	GetZoneId() *string
}

type CapacityLock struct {
	AvailableCount           *int32  `json:"availableCount,omitempty" xml:"availableCount,omitempty"`
	CrsReservationId         *string `json:"crsReservationId,omitempty" xml:"crsReservationId,omitempty"`
	Description              *string `json:"description,omitempty" xml:"description,omitempty"`
	ExpireTime               *string `json:"expireTime,omitempty" xml:"expireTime,omitempty"`
	GmtCreated               *string `json:"gmtCreated,omitempty" xml:"gmtCreated,omitempty"`
	GmtModified              *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	Id                       *string `json:"id,omitempty" xml:"id,omitempty"`
	InstanceType             *string `json:"instanceType,omitempty" xml:"instanceType,omitempty"`
	LastReconcileAttemptTime *string `json:"lastReconcileAttemptTime,omitempty" xml:"lastReconcileAttemptTime,omitempty"`
	LastSyncTime             *string `json:"lastSyncTime,omitempty" xml:"lastSyncTime,omitempty"`
	LockProvider             *string `json:"lockProvider,omitempty" xml:"lockProvider,omitempty"`
	LockedCount              *int32  `json:"lockedCount,omitempty" xml:"lockedCount,omitempty"`
	Operator                 *string `json:"operator,omitempty" xml:"operator,omitempty"`
	PaymentType              *string `json:"paymentType,omitempty" xml:"paymentType,omitempty"`
	PrivatePoolId            *string `json:"privatePoolId,omitempty" xml:"privatePoolId,omitempty"`
	RequestedCount           *int32  `json:"requestedCount,omitempty" xml:"requestedCount,omitempty"`
	Status                   *string `json:"status,omitempty" xml:"status,omitempty"`
	TenantId                 *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	UsedCount                *int32  `json:"usedCount,omitempty" xml:"usedCount,omitempty"`
	ZoneId                   *string `json:"zoneId,omitempty" xml:"zoneId,omitempty"`
}

func (s CapacityLock) String() string {
	return dara.Prettify(s)
}

func (s CapacityLock) GoString() string {
	return s.String()
}

func (s *CapacityLock) GetAvailableCount() *int32 {
	return s.AvailableCount
}

func (s *CapacityLock) GetCrsReservationId() *string {
	return s.CrsReservationId
}

func (s *CapacityLock) GetDescription() *string {
	return s.Description
}

func (s *CapacityLock) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *CapacityLock) GetGmtCreated() *string {
	return s.GmtCreated
}

func (s *CapacityLock) GetGmtModified() *string {
	return s.GmtModified
}

func (s *CapacityLock) GetId() *string {
	return s.Id
}

func (s *CapacityLock) GetInstanceType() *string {
	return s.InstanceType
}

func (s *CapacityLock) GetLastReconcileAttemptTime() *string {
	return s.LastReconcileAttemptTime
}

func (s *CapacityLock) GetLastSyncTime() *string {
	return s.LastSyncTime
}

func (s *CapacityLock) GetLockProvider() *string {
	return s.LockProvider
}

func (s *CapacityLock) GetLockedCount() *int32 {
	return s.LockedCount
}

func (s *CapacityLock) GetOperator() *string {
	return s.Operator
}

func (s *CapacityLock) GetPaymentType() *string {
	return s.PaymentType
}

func (s *CapacityLock) GetPrivatePoolId() *string {
	return s.PrivatePoolId
}

func (s *CapacityLock) GetRequestedCount() *int32 {
	return s.RequestedCount
}

func (s *CapacityLock) GetStatus() *string {
	return s.Status
}

func (s *CapacityLock) GetTenantId() *string {
	return s.TenantId
}

func (s *CapacityLock) GetUsedCount() *int32 {
	return s.UsedCount
}

func (s *CapacityLock) GetZoneId() *string {
	return s.ZoneId
}

func (s *CapacityLock) SetAvailableCount(v int32) *CapacityLock {
	s.AvailableCount = &v
	return s
}

func (s *CapacityLock) SetCrsReservationId(v string) *CapacityLock {
	s.CrsReservationId = &v
	return s
}

func (s *CapacityLock) SetDescription(v string) *CapacityLock {
	s.Description = &v
	return s
}

func (s *CapacityLock) SetExpireTime(v string) *CapacityLock {
	s.ExpireTime = &v
	return s
}

func (s *CapacityLock) SetGmtCreated(v string) *CapacityLock {
	s.GmtCreated = &v
	return s
}

func (s *CapacityLock) SetGmtModified(v string) *CapacityLock {
	s.GmtModified = &v
	return s
}

func (s *CapacityLock) SetId(v string) *CapacityLock {
	s.Id = &v
	return s
}

func (s *CapacityLock) SetInstanceType(v string) *CapacityLock {
	s.InstanceType = &v
	return s
}

func (s *CapacityLock) SetLastReconcileAttemptTime(v string) *CapacityLock {
	s.LastReconcileAttemptTime = &v
	return s
}

func (s *CapacityLock) SetLastSyncTime(v string) *CapacityLock {
	s.LastSyncTime = &v
	return s
}

func (s *CapacityLock) SetLockProvider(v string) *CapacityLock {
	s.LockProvider = &v
	return s
}

func (s *CapacityLock) SetLockedCount(v int32) *CapacityLock {
	s.LockedCount = &v
	return s
}

func (s *CapacityLock) SetOperator(v string) *CapacityLock {
	s.Operator = &v
	return s
}

func (s *CapacityLock) SetPaymentType(v string) *CapacityLock {
	s.PaymentType = &v
	return s
}

func (s *CapacityLock) SetPrivatePoolId(v string) *CapacityLock {
	s.PrivatePoolId = &v
	return s
}

func (s *CapacityLock) SetRequestedCount(v int32) *CapacityLock {
	s.RequestedCount = &v
	return s
}

func (s *CapacityLock) SetStatus(v string) *CapacityLock {
	s.Status = &v
	return s
}

func (s *CapacityLock) SetTenantId(v string) *CapacityLock {
	s.TenantId = &v
	return s
}

func (s *CapacityLock) SetUsedCount(v int32) *CapacityLock {
	s.UsedCount = &v
	return s
}

func (s *CapacityLock) SetZoneId(v string) *CapacityLock {
	s.ZoneId = &v
	return s
}

func (s *CapacityLock) Validate() error {
	return dara.Validate(s)
}
