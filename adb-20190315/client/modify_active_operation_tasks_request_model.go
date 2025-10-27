// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyActiveOperationTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIds(v string) *ModifyActiveOperationTasksRequest
	GetIds() *string
	SetImmediateStart(v int32) *ModifyActiveOperationTasksRequest
	GetImmediateStart() *int32
	SetOwnerAccount(v string) *ModifyActiveOperationTasksRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyActiveOperationTasksRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyActiveOperationTasksRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyActiveOperationTasksRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyActiveOperationTasksRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *ModifyActiveOperationTasksRequest
	GetSecurityToken() *string
	SetSwitchTime(v string) *ModifyActiveOperationTasksRequest
	GetSwitchTime() *string
}

type ModifyActiveOperationTasksRequest struct {
	// The ID of the O\\&M event.
	//
	// example:
	//
	// 1482487
	Ids *string `json:"Ids,omitempty" xml:"Ids,omitempty"`
	// Specifies whether to immediately execute the O\\&M event. Valid values:
	//
	// 	- 1: immediately executes the O\\&M event.
	//
	// 	- 0: executes the O\\&M event at a specific point in time.
	//
	// example:
	//
	// 1
	ImmediateStart *int32  `json:"ImmediateStart,omitempty" xml:"ImmediateStart,omitempty"`
	OwnerAccount   *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The time from which you want to execute the O\\&M event.
	//
	// example:
	//
	// 2021-08-15T12:00:00Z
	SwitchTime *string `json:"SwitchTime,omitempty" xml:"SwitchTime,omitempty"`
}

func (s ModifyActiveOperationTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyActiveOperationTasksRequest) GoString() string {
	return s.String()
}

func (s *ModifyActiveOperationTasksRequest) GetIds() *string {
	return s.Ids
}

func (s *ModifyActiveOperationTasksRequest) GetImmediateStart() *int32 {
	return s.ImmediateStart
}

func (s *ModifyActiveOperationTasksRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyActiveOperationTasksRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyActiveOperationTasksRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyActiveOperationTasksRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyActiveOperationTasksRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyActiveOperationTasksRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *ModifyActiveOperationTasksRequest) GetSwitchTime() *string {
	return s.SwitchTime
}

func (s *ModifyActiveOperationTasksRequest) SetIds(v string) *ModifyActiveOperationTasksRequest {
	s.Ids = &v
	return s
}

func (s *ModifyActiveOperationTasksRequest) SetImmediateStart(v int32) *ModifyActiveOperationTasksRequest {
	s.ImmediateStart = &v
	return s
}

func (s *ModifyActiveOperationTasksRequest) SetOwnerAccount(v string) *ModifyActiveOperationTasksRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyActiveOperationTasksRequest) SetOwnerId(v int64) *ModifyActiveOperationTasksRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyActiveOperationTasksRequest) SetRegionId(v string) *ModifyActiveOperationTasksRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyActiveOperationTasksRequest) SetResourceOwnerAccount(v string) *ModifyActiveOperationTasksRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyActiveOperationTasksRequest) SetResourceOwnerId(v int64) *ModifyActiveOperationTasksRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyActiveOperationTasksRequest) SetSecurityToken(v string) *ModifyActiveOperationTasksRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyActiveOperationTasksRequest) SetSwitchTime(v string) *ModifyActiveOperationTasksRequest {
	s.SwitchTime = &v
	return s
}

func (s *ModifyActiveOperationTasksRequest) Validate() error {
	return dara.Validate(s)
}
