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
	// The batch ID. This ID is returned by the [ImportTaskNumberDatas](https://help.aliyun.com/document_detail/2926815.html) operation when you import callee data. You can view this ID on the **execution history*	- page by navigating to **call task management*	- > **details**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 139*********216
	BatchId *string `json:"BatchId,omitempty" xml:"BatchId,omitempty"`
	// The detail ID. You can find this ID in the upper-left corner of the page by navigating to **call task management*	- > **details*	- > **execution history*	- > **completed*	- > **call details**, or get it by calling the [QueryAiCallDetailPage](https://help.aliyun.com/document_detail/2926853.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 9662*************
	DetailId             *string `json:"DetailId,omitempty" xml:"DetailId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The task ID. You can find this ID on the **call task management*	- page or get it by calling the [QueryAiCallTaskPage](https://help.aliyun.com/document_detail/2926799.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 138************
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
