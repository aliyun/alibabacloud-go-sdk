// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCallNumberDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *CallNumberDetailRequest
	GetId() *int64
	SetOwnerId(v int64) *CallNumberDetailRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CallNumberDetailRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CallNumberDetailRequest
	GetResourceOwnerId() *int64
	SetTaskId(v int64) *CallNumberDetailRequest
	GetTaskId() *int64
}

type CallNumberDetailRequest struct {
	// 外呼编号ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 95
	Id                   *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 任务ID
	//
	// example:
	//
	// 33
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CallNumberDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s CallNumberDetailRequest) GoString() string {
	return s.String()
}

func (s *CallNumberDetailRequest) GetId() *int64 {
	return s.Id
}

func (s *CallNumberDetailRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CallNumberDetailRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CallNumberDetailRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CallNumberDetailRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *CallNumberDetailRequest) SetId(v int64) *CallNumberDetailRequest {
	s.Id = &v
	return s
}

func (s *CallNumberDetailRequest) SetOwnerId(v int64) *CallNumberDetailRequest {
	s.OwnerId = &v
	return s
}

func (s *CallNumberDetailRequest) SetResourceOwnerAccount(v string) *CallNumberDetailRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CallNumberDetailRequest) SetResourceOwnerId(v int64) *CallNumberDetailRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CallNumberDetailRequest) SetTaskId(v int64) *CallNumberDetailRequest {
	s.TaskId = &v
	return s
}

func (s *CallNumberDetailRequest) Validate() error {
	return dara.Validate(s)
}
