// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryConversationDetailInfoNewRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallId(v string) *QueryConversationDetailInfoNewRequest
	GetCallId() *string
	SetDetailId(v string) *QueryConversationDetailInfoNewRequest
	GetDetailId() *string
	SetOutId(v string) *QueryConversationDetailInfoNewRequest
	GetOutId() *string
	SetOwnerId(v int64) *QueryConversationDetailInfoNewRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *QueryConversationDetailInfoNewRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QueryConversationDetailInfoNewRequest
	GetResourceOwnerId() *int64
	SetTaskId(v string) *QueryConversationDetailInfoNewRequest
	GetTaskId() *string
}

type QueryConversationDetailInfoNewRequest struct {
	// The unique ID of the call.
	//
	// example:
	//
	// 1552********^1420********
	CallId *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	// The task detail ID. View the task detail ID in the upper-left corner of the **Call Task Management*	- > **Details*	- > **Execution Records*	- > **Completed*	- > **Call Details*	- console, or call the [QueryAiCallDetailPage](https://help.aliyun.com/document_detail/2926853.html) operation to obtain the task detail ID.
	//
	// example:
	//
	// 9662*************
	DetailId *string `json:"DetailId,omitempty" xml:"DetailId,omitempty"`
	// The external business serial number reserved for external input. You can use a unique ID for business association.
	//
	// example:
	//
	// 123******
	OutId                *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The task ID. View the task ID in the **Call Task Management*	- console or call the [QueryAiCallTaskPage](https://help.aliyun.com/document_detail/2926799.html) operation to obtain the task ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 138************
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s QueryConversationDetailInfoNewRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryConversationDetailInfoNewRequest) GoString() string {
	return s.String()
}

func (s *QueryConversationDetailInfoNewRequest) GetCallId() *string {
	return s.CallId
}

func (s *QueryConversationDetailInfoNewRequest) GetDetailId() *string {
	return s.DetailId
}

func (s *QueryConversationDetailInfoNewRequest) GetOutId() *string {
	return s.OutId
}

func (s *QueryConversationDetailInfoNewRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QueryConversationDetailInfoNewRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QueryConversationDetailInfoNewRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QueryConversationDetailInfoNewRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *QueryConversationDetailInfoNewRequest) SetCallId(v string) *QueryConversationDetailInfoNewRequest {
	s.CallId = &v
	return s
}

func (s *QueryConversationDetailInfoNewRequest) SetDetailId(v string) *QueryConversationDetailInfoNewRequest {
	s.DetailId = &v
	return s
}

func (s *QueryConversationDetailInfoNewRequest) SetOutId(v string) *QueryConversationDetailInfoNewRequest {
	s.OutId = &v
	return s
}

func (s *QueryConversationDetailInfoNewRequest) SetOwnerId(v int64) *QueryConversationDetailInfoNewRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryConversationDetailInfoNewRequest) SetResourceOwnerAccount(v string) *QueryConversationDetailInfoNewRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryConversationDetailInfoNewRequest) SetResourceOwnerId(v int64) *QueryConversationDetailInfoNewRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryConversationDetailInfoNewRequest) SetTaskId(v string) *QueryConversationDetailInfoNewRequest {
	s.TaskId = &v
	return s
}

func (s *QueryConversationDetailInfoNewRequest) Validate() error {
	return dara.Validate(s)
}
