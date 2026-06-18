// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAiCallTaskDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *QueryAiCallTaskDetailRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *QueryAiCallTaskDetailRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QueryAiCallTaskDetailRequest
	GetResourceOwnerId() *int64
	SetTaskId(v string) *QueryAiCallTaskDetailRequest
	GetTaskId() *string
}

type QueryAiCallTaskDetailRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The task ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1187**************
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s QueryAiCallTaskDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryAiCallTaskDetailRequest) GoString() string {
	return s.String()
}

func (s *QueryAiCallTaskDetailRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QueryAiCallTaskDetailRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QueryAiCallTaskDetailRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QueryAiCallTaskDetailRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *QueryAiCallTaskDetailRequest) SetOwnerId(v int64) *QueryAiCallTaskDetailRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryAiCallTaskDetailRequest) SetResourceOwnerAccount(v string) *QueryAiCallTaskDetailRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryAiCallTaskDetailRequest) SetResourceOwnerId(v int64) *QueryAiCallTaskDetailRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryAiCallTaskDetailRequest) SetTaskId(v string) *QueryAiCallTaskDetailRequest {
	s.TaskId = &v
	return s
}

func (s *QueryAiCallTaskDetailRequest) Validate() error {
	return dara.Validate(s)
}
