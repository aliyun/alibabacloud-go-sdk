// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyActiveOperationMaintainConfRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCycleTime(v string) *ModifyActiveOperationMaintainConfRequest
	GetCycleTime() *string
	SetCycleType(v string) *ModifyActiveOperationMaintainConfRequest
	GetCycleType() *string
	SetMaintainEndTime(v string) *ModifyActiveOperationMaintainConfRequest
	GetMaintainEndTime() *string
	SetMaintainStartTime(v string) *ModifyActiveOperationMaintainConfRequest
	GetMaintainStartTime() *string
	SetOwnerAccount(v string) *ModifyActiveOperationMaintainConfRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyActiveOperationMaintainConfRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyActiveOperationMaintainConfRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyActiveOperationMaintainConfRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyActiveOperationMaintainConfRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *ModifyActiveOperationMaintainConfRequest
	GetSecurityToken() *string
	SetStatus(v int32) *ModifyActiveOperationMaintainConfRequest
	GetStatus() *int32
}

type ModifyActiveOperationMaintainConfRequest struct {
	// The days on which you want to execute O\\&M events. Valid values: 1 to 7.
	//
	// example:
	//
	// 1,2,3,4,5,6,7
	CycleTime *string `json:"CycleTime,omitempty" xml:"CycleTime,omitempty"`
	// The unit of the O\\&M cycle. Set the value to Week.
	//
	// example:
	//
	// Week
	CycleType *string `json:"CycleType,omitempty" xml:"CycleType,omitempty"`
	// The end time of the maintenance window.
	//
	// example:
	//
	// 02:00:00Z
	MaintainEndTime *string `json:"MaintainEndTime,omitempty" xml:"MaintainEndTime,omitempty"`
	// The start time of the maintenance window.
	//
	// example:
	//
	// 02:00:00Z
	MaintainStartTime *string `json:"MaintainStartTime,omitempty" xml:"MaintainStartTime,omitempty"`
	OwnerAccount      *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId           *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// Specifies whether to enable O\\&M events.
	//
	// 	- 1: yes.
	//
	// 	- 0: no.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ModifyActiveOperationMaintainConfRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyActiveOperationMaintainConfRequest) GoString() string {
	return s.String()
}

func (s *ModifyActiveOperationMaintainConfRequest) GetCycleTime() *string {
	return s.CycleTime
}

func (s *ModifyActiveOperationMaintainConfRequest) GetCycleType() *string {
	return s.CycleType
}

func (s *ModifyActiveOperationMaintainConfRequest) GetMaintainEndTime() *string {
	return s.MaintainEndTime
}

func (s *ModifyActiveOperationMaintainConfRequest) GetMaintainStartTime() *string {
	return s.MaintainStartTime
}

func (s *ModifyActiveOperationMaintainConfRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyActiveOperationMaintainConfRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyActiveOperationMaintainConfRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyActiveOperationMaintainConfRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyActiveOperationMaintainConfRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyActiveOperationMaintainConfRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *ModifyActiveOperationMaintainConfRequest) GetStatus() *int32 {
	return s.Status
}

func (s *ModifyActiveOperationMaintainConfRequest) SetCycleTime(v string) *ModifyActiveOperationMaintainConfRequest {
	s.CycleTime = &v
	return s
}

func (s *ModifyActiveOperationMaintainConfRequest) SetCycleType(v string) *ModifyActiveOperationMaintainConfRequest {
	s.CycleType = &v
	return s
}

func (s *ModifyActiveOperationMaintainConfRequest) SetMaintainEndTime(v string) *ModifyActiveOperationMaintainConfRequest {
	s.MaintainEndTime = &v
	return s
}

func (s *ModifyActiveOperationMaintainConfRequest) SetMaintainStartTime(v string) *ModifyActiveOperationMaintainConfRequest {
	s.MaintainStartTime = &v
	return s
}

func (s *ModifyActiveOperationMaintainConfRequest) SetOwnerAccount(v string) *ModifyActiveOperationMaintainConfRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyActiveOperationMaintainConfRequest) SetOwnerId(v int64) *ModifyActiveOperationMaintainConfRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyActiveOperationMaintainConfRequest) SetRegionId(v string) *ModifyActiveOperationMaintainConfRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyActiveOperationMaintainConfRequest) SetResourceOwnerAccount(v string) *ModifyActiveOperationMaintainConfRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyActiveOperationMaintainConfRequest) SetResourceOwnerId(v int64) *ModifyActiveOperationMaintainConfRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyActiveOperationMaintainConfRequest) SetSecurityToken(v string) *ModifyActiveOperationMaintainConfRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyActiveOperationMaintainConfRequest) SetStatus(v int32) *ModifyActiveOperationMaintainConfRequest {
	s.Status = &v
	return s
}

func (s *ModifyActiveOperationMaintainConfRequest) Validate() error {
	return dara.Validate(s)
}
