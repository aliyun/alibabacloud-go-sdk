// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyActiveOperationTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIds(v string) *ModifyActiveOperationTaskRequest
	GetIds() *string
	SetOwnerAccount(v string) *ModifyActiveOperationTaskRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyActiveOperationTaskRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ModifyActiveOperationTaskRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyActiveOperationTaskRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *ModifyActiveOperationTaskRequest
	GetSecurityToken() *string
	SetSwitchTime(v string) *ModifyActiveOperationTaskRequest
	GetSwitchTime() *string
}

type ModifyActiveOperationTaskRequest struct {
	// The ID of the O\\&M task. Separate multiple IDs with commas (,).
	//
	// > You can call the [DescribeActiveOperationTask](https://help.aliyun.com/document_detail/473865.html) operation to query the ID of an O\\&M task.
	//
	// This parameter is required.
	//
	// example:
	//
	// 11111,22222
	Ids                  *string `json:"Ids,omitempty" xml:"Ids,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The scheduled switchover time to be specified. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC.
	//
	// > The time cannot be later than the latest operation time. You can call the [DescribeActiveOperationTask](https://help.aliyun.com/document_detail/473865.html) operation to obtain the latest operation time, which is the value of the **Deadline*	- parameter in the response.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2019-10-17T18:50:00Z
	SwitchTime *string `json:"SwitchTime,omitempty" xml:"SwitchTime,omitempty"`
}

func (s ModifyActiveOperationTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyActiveOperationTaskRequest) GoString() string {
	return s.String()
}

func (s *ModifyActiveOperationTaskRequest) GetIds() *string {
	return s.Ids
}

func (s *ModifyActiveOperationTaskRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyActiveOperationTaskRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyActiveOperationTaskRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyActiveOperationTaskRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyActiveOperationTaskRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *ModifyActiveOperationTaskRequest) GetSwitchTime() *string {
	return s.SwitchTime
}

func (s *ModifyActiveOperationTaskRequest) SetIds(v string) *ModifyActiveOperationTaskRequest {
	s.Ids = &v
	return s
}

func (s *ModifyActiveOperationTaskRequest) SetOwnerAccount(v string) *ModifyActiveOperationTaskRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyActiveOperationTaskRequest) SetOwnerId(v int64) *ModifyActiveOperationTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyActiveOperationTaskRequest) SetResourceOwnerAccount(v string) *ModifyActiveOperationTaskRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyActiveOperationTaskRequest) SetResourceOwnerId(v int64) *ModifyActiveOperationTaskRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyActiveOperationTaskRequest) SetSecurityToken(v string) *ModifyActiveOperationTaskRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyActiveOperationTaskRequest) SetSwitchTime(v string) *ModifyActiveOperationTaskRequest {
	s.SwitchTime = &v
	return s
}

func (s *ModifyActiveOperationTaskRequest) Validate() error {
	return dara.Validate(s)
}
