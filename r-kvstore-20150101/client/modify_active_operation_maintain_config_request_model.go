// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyActiveOperationMaintainConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCycleTime(v string) *ModifyActiveOperationMaintainConfigRequest
	GetCycleTime() *string
	SetCycleType(v string) *ModifyActiveOperationMaintainConfigRequest
	GetCycleType() *string
	SetMaintainEndTime(v string) *ModifyActiveOperationMaintainConfigRequest
	GetMaintainEndTime() *string
	SetMaintainStartTime(v string) *ModifyActiveOperationMaintainConfigRequest
	GetMaintainStartTime() *string
	SetOwnerAccount(v string) *ModifyActiveOperationMaintainConfigRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyActiveOperationMaintainConfigRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ModifyActiveOperationMaintainConfigRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyActiveOperationMaintainConfigRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *ModifyActiveOperationMaintainConfigRequest
	GetSecurityToken() *string
	SetStatus(v int32) *ModifyActiveOperationMaintainConfigRequest
	GetStatus() *int32
}

type ModifyActiveOperationMaintainConfigRequest struct {
	// The days of the cycle.
	//
	// - If `CycleType` is `Month`, specify the days of the month (1 to 28). Separate multiple days with a comma (,).
	//
	// - If `CycleType` is `Week`, specify the days of the week (1 to 7). Separate multiple days with a comma (,).
	//
	// example:
	//
	// 1,2,3,4,5
	CycleTime *string `json:"CycleTime,omitempty" xml:"CycleTime,omitempty"`
	// The cycle type of the maintenance window. Valid values:
	//
	// - `Month`
	//
	// - `Week`
	//
	// example:
	//
	// Week
	CycleType *string `json:"CycleType,omitempty" xml:"CycleType,omitempty"`
	// The end time of the maintenance window, specified in *HH:mm:ss*Z format (UTC time).
	//
	// example:
	//
	// 20:00:00Z
	MaintainEndTime *string `json:"MaintainEndTime,omitempty" xml:"MaintainEndTime,omitempty"`
	// The start time of the maintenance window, specified in *HH:mm:ss*Z format (UTC time).
	//
	// example:
	//
	// 16:00:00Z
	MaintainStartTime    *string `json:"MaintainStartTime,omitempty" xml:"MaintainStartTime,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// Specifies whether the configuration is enabled. Valid values:
	//
	// - 1: enabled
	//
	// - 2: disabled
	//
	// example:
	//
	// 2
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ModifyActiveOperationMaintainConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyActiveOperationMaintainConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyActiveOperationMaintainConfigRequest) GetCycleTime() *string {
	return s.CycleTime
}

func (s *ModifyActiveOperationMaintainConfigRequest) GetCycleType() *string {
	return s.CycleType
}

func (s *ModifyActiveOperationMaintainConfigRequest) GetMaintainEndTime() *string {
	return s.MaintainEndTime
}

func (s *ModifyActiveOperationMaintainConfigRequest) GetMaintainStartTime() *string {
	return s.MaintainStartTime
}

func (s *ModifyActiveOperationMaintainConfigRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyActiveOperationMaintainConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyActiveOperationMaintainConfigRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyActiveOperationMaintainConfigRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyActiveOperationMaintainConfigRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *ModifyActiveOperationMaintainConfigRequest) GetStatus() *int32 {
	return s.Status
}

func (s *ModifyActiveOperationMaintainConfigRequest) SetCycleTime(v string) *ModifyActiveOperationMaintainConfigRequest {
	s.CycleTime = &v
	return s
}

func (s *ModifyActiveOperationMaintainConfigRequest) SetCycleType(v string) *ModifyActiveOperationMaintainConfigRequest {
	s.CycleType = &v
	return s
}

func (s *ModifyActiveOperationMaintainConfigRequest) SetMaintainEndTime(v string) *ModifyActiveOperationMaintainConfigRequest {
	s.MaintainEndTime = &v
	return s
}

func (s *ModifyActiveOperationMaintainConfigRequest) SetMaintainStartTime(v string) *ModifyActiveOperationMaintainConfigRequest {
	s.MaintainStartTime = &v
	return s
}

func (s *ModifyActiveOperationMaintainConfigRequest) SetOwnerAccount(v string) *ModifyActiveOperationMaintainConfigRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyActiveOperationMaintainConfigRequest) SetOwnerId(v int64) *ModifyActiveOperationMaintainConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyActiveOperationMaintainConfigRequest) SetResourceOwnerAccount(v string) *ModifyActiveOperationMaintainConfigRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyActiveOperationMaintainConfigRequest) SetResourceOwnerId(v int64) *ModifyActiveOperationMaintainConfigRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyActiveOperationMaintainConfigRequest) SetSecurityToken(v string) *ModifyActiveOperationMaintainConfigRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyActiveOperationMaintainConfigRequest) SetStatus(v int32) *ModifyActiveOperationMaintainConfigRequest {
	s.Status = &v
	return s
}

func (s *ModifyActiveOperationMaintainConfigRequest) Validate() error {
	return dara.Validate(s)
}
