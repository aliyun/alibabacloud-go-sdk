// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyActiveOperationMaintenanceConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCycleTime(v string) *ModifyActiveOperationMaintenanceConfigRequest
	GetCycleTime() *string
	SetCycleType(v string) *ModifyActiveOperationMaintenanceConfigRequest
	GetCycleType() *string
	SetMaintainEndTime(v string) *ModifyActiveOperationMaintenanceConfigRequest
	GetMaintainEndTime() *string
	SetMaintainStartTime(v string) *ModifyActiveOperationMaintenanceConfigRequest
	GetMaintainStartTime() *string
	SetOwnerAccount(v string) *ModifyActiveOperationMaintenanceConfigRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyActiveOperationMaintenanceConfigRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ModifyActiveOperationMaintenanceConfigRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyActiveOperationMaintenanceConfigRequest
	GetResourceOwnerId() *int64
	SetStatus(v int32) *ModifyActiveOperationMaintenanceConfigRequest
	GetStatus() *int32
}

type ModifyActiveOperationMaintenanceConfigRequest struct {
	// example:
	//
	// 1,2,3
	CycleTime *string `json:"CycleTime,omitempty" xml:"CycleTime,omitempty"`
	// example:
	//
	// Week
	CycleType *string `json:"CycleType,omitempty" xml:"CycleType,omitempty"`
	// example:
	//
	// 22:00Z
	MaintainEndTime *string `json:"MaintainEndTime,omitempty" xml:"MaintainEndTime,omitempty"`
	// example:
	//
	// 21:00Z
	MaintainStartTime    *string `json:"MaintainStartTime,omitempty" xml:"MaintainStartTime,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ModifyActiveOperationMaintenanceConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyActiveOperationMaintenanceConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyActiveOperationMaintenanceConfigRequest) GetCycleTime() *string {
	return s.CycleTime
}

func (s *ModifyActiveOperationMaintenanceConfigRequest) GetCycleType() *string {
	return s.CycleType
}

func (s *ModifyActiveOperationMaintenanceConfigRequest) GetMaintainEndTime() *string {
	return s.MaintainEndTime
}

func (s *ModifyActiveOperationMaintenanceConfigRequest) GetMaintainStartTime() *string {
	return s.MaintainStartTime
}

func (s *ModifyActiveOperationMaintenanceConfigRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyActiveOperationMaintenanceConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyActiveOperationMaintenanceConfigRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyActiveOperationMaintenanceConfigRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyActiveOperationMaintenanceConfigRequest) GetStatus() *int32 {
	return s.Status
}

func (s *ModifyActiveOperationMaintenanceConfigRequest) SetCycleTime(v string) *ModifyActiveOperationMaintenanceConfigRequest {
	s.CycleTime = &v
	return s
}

func (s *ModifyActiveOperationMaintenanceConfigRequest) SetCycleType(v string) *ModifyActiveOperationMaintenanceConfigRequest {
	s.CycleType = &v
	return s
}

func (s *ModifyActiveOperationMaintenanceConfigRequest) SetMaintainEndTime(v string) *ModifyActiveOperationMaintenanceConfigRequest {
	s.MaintainEndTime = &v
	return s
}

func (s *ModifyActiveOperationMaintenanceConfigRequest) SetMaintainStartTime(v string) *ModifyActiveOperationMaintenanceConfigRequest {
	s.MaintainStartTime = &v
	return s
}

func (s *ModifyActiveOperationMaintenanceConfigRequest) SetOwnerAccount(v string) *ModifyActiveOperationMaintenanceConfigRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyActiveOperationMaintenanceConfigRequest) SetOwnerId(v int64) *ModifyActiveOperationMaintenanceConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyActiveOperationMaintenanceConfigRequest) SetResourceOwnerAccount(v string) *ModifyActiveOperationMaintenanceConfigRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyActiveOperationMaintenanceConfigRequest) SetResourceOwnerId(v int64) *ModifyActiveOperationMaintenanceConfigRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyActiveOperationMaintenanceConfigRequest) SetStatus(v int32) *ModifyActiveOperationMaintenanceConfigRequest {
	s.Status = &v
	return s
}

func (s *ModifyActiveOperationMaintenanceConfigRequest) Validate() error {
	return dara.Validate(s)
}
