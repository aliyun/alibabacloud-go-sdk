// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryConversationDetailInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBatchId(v string) *QueryConversationDetailInfoRequest
	GetBatchId() *string
	SetDetailId(v string) *QueryConversationDetailInfoRequest
	GetDetailId() *string
	SetOwnerId(v int64) *QueryConversationDetailInfoRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *QueryConversationDetailInfoRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QueryConversationDetailInfoRequest
	GetResourceOwnerId() *int64
	SetTaskId(v string) *QueryConversationDetailInfoRequest
	GetTaskId() *string
}

type QueryConversationDetailInfoRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1231321231****
	BatchId *string `json:"BatchId,omitempty" xml:"BatchId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1231321231****
	DetailId             *string `json:"DetailId,omitempty" xml:"DetailId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1231321231****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s QueryConversationDetailInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryConversationDetailInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryConversationDetailInfoRequest) GetBatchId() *string {
	return s.BatchId
}

func (s *QueryConversationDetailInfoRequest) GetDetailId() *string {
	return s.DetailId
}

func (s *QueryConversationDetailInfoRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QueryConversationDetailInfoRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QueryConversationDetailInfoRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QueryConversationDetailInfoRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *QueryConversationDetailInfoRequest) SetBatchId(v string) *QueryConversationDetailInfoRequest {
	s.BatchId = &v
	return s
}

func (s *QueryConversationDetailInfoRequest) SetDetailId(v string) *QueryConversationDetailInfoRequest {
	s.DetailId = &v
	return s
}

func (s *QueryConversationDetailInfoRequest) SetOwnerId(v int64) *QueryConversationDetailInfoRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryConversationDetailInfoRequest) SetResourceOwnerAccount(v string) *QueryConversationDetailInfoRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryConversationDetailInfoRequest) SetResourceOwnerId(v int64) *QueryConversationDetailInfoRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryConversationDetailInfoRequest) SetTaskId(v string) *QueryConversationDetailInfoRequest {
	s.TaskId = &v
	return s
}

func (s *QueryConversationDetailInfoRequest) Validate() error {
	return dara.Validate(s)
}
