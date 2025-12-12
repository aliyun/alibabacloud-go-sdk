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
	// example:
	//
	// 1
	CycleTime *string `json:"CycleTime,omitempty" xml:"CycleTime,omitempty"`
	// example:
	//
	// Week
	CycleType *string `json:"CycleType,omitempty" xml:"CycleType,omitempty"`
	// example:
	//
	// 20:00Z
	MaintainEndTime *string `json:"MaintainEndTime,omitempty" xml:"MaintainEndTime,omitempty"`
	// example:
	//
	// 07:00Z
	MaintainStartTime    *string `json:"MaintainStartTime,omitempty" xml:"MaintainStartTime,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
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
