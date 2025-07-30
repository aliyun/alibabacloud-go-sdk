// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceMaintainTimeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ModifyInstanceMaintainTimeRequest
	GetInstanceId() *string
	SetMaintainEndTime(v string) *ModifyInstanceMaintainTimeRequest
	GetMaintainEndTime() *string
	SetMaintainStartTime(v string) *ModifyInstanceMaintainTimeRequest
	GetMaintainStartTime() *string
	SetOwnerAccount(v string) *ModifyInstanceMaintainTimeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyInstanceMaintainTimeRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ModifyInstanceMaintainTimeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyInstanceMaintainTimeRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *ModifyInstanceMaintainTimeRequest
	GetSecurityToken() *string
}

type ModifyInstanceMaintainTimeRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The end time of the maintenance window. The time is in the *HH:mm*Z format. The time is displayed in UTC. For example, if you want the maintenance window to end at 2:00 (UTC+8), set this parameter to `18:00Z`.
	//
	// >  The interval between the start time and the end time cannot be less than 1 hour.
	//
	// This parameter is required.
	//
	// example:
	//
	// 04:00Z
	MaintainEndTime *string `json:"MaintainEndTime,omitempty" xml:"MaintainEndTime,omitempty"`
	// The start time of the maintenance window. The time is in the *HH:mm*Z format. The time is displayed in UTC. For example, if you want the maintenance to start at 1:00 (UTC+8), set this parameter to `17:00Z`. After you call the API operation, you can view the actual time in the Tair (Redis OSS-compatible) console. For more information, see [Set a maintenance window](https://help.aliyun.com/document_detail/55252.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 03:00Z
	MaintainStartTime    *string `json:"MaintainStartTime,omitempty" xml:"MaintainStartTime,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s ModifyInstanceMaintainTimeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceMaintainTimeRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceMaintainTimeRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyInstanceMaintainTimeRequest) GetMaintainEndTime() *string {
	return s.MaintainEndTime
}

func (s *ModifyInstanceMaintainTimeRequest) GetMaintainStartTime() *string {
	return s.MaintainStartTime
}

func (s *ModifyInstanceMaintainTimeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyInstanceMaintainTimeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyInstanceMaintainTimeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyInstanceMaintainTimeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyInstanceMaintainTimeRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *ModifyInstanceMaintainTimeRequest) SetInstanceId(v string) *ModifyInstanceMaintainTimeRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyInstanceMaintainTimeRequest) SetMaintainEndTime(v string) *ModifyInstanceMaintainTimeRequest {
	s.MaintainEndTime = &v
	return s
}

func (s *ModifyInstanceMaintainTimeRequest) SetMaintainStartTime(v string) *ModifyInstanceMaintainTimeRequest {
	s.MaintainStartTime = &v
	return s
}

func (s *ModifyInstanceMaintainTimeRequest) SetOwnerAccount(v string) *ModifyInstanceMaintainTimeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyInstanceMaintainTimeRequest) SetOwnerId(v int64) *ModifyInstanceMaintainTimeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyInstanceMaintainTimeRequest) SetResourceOwnerAccount(v string) *ModifyInstanceMaintainTimeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyInstanceMaintainTimeRequest) SetResourceOwnerId(v int64) *ModifyInstanceMaintainTimeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyInstanceMaintainTimeRequest) SetSecurityToken(v string) *ModifyInstanceMaintainTimeRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyInstanceMaintainTimeRequest) Validate() error {
	return dara.Validate(s)
}
