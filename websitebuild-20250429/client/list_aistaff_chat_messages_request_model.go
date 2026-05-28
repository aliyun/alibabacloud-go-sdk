// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAIStaffChatMessagesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *ListAIStaffChatMessagesRequest
	GetBizId() *string
	SetConversationId(v string) *ListAIStaffChatMessagesRequest
	GetConversationId() *string
	SetPageSize(v int32) *ListAIStaffChatMessagesRequest
	GetPageSize() *int32
	SetStartCreateTime(v string) *ListAIStaffChatMessagesRequest
	GetStartCreateTime() *string
}

type ListAIStaffChatMessagesRequest struct {
	// example:
	//
	// WD20250703155602000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// example:
	//
	// 81bc5a34-1d8d-4ef7-a208-7401c51b054b
	ConversationId *string `json:"ConversationId,omitempty" xml:"ConversationId,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// null
	StartCreateTime *string `json:"StartCreateTime,omitempty" xml:"StartCreateTime,omitempty"`
}

func (s ListAIStaffChatMessagesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAIStaffChatMessagesRequest) GoString() string {
	return s.String()
}

func (s *ListAIStaffChatMessagesRequest) GetBizId() *string {
	return s.BizId
}

func (s *ListAIStaffChatMessagesRequest) GetConversationId() *string {
	return s.ConversationId
}

func (s *ListAIStaffChatMessagesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAIStaffChatMessagesRequest) GetStartCreateTime() *string {
	return s.StartCreateTime
}

func (s *ListAIStaffChatMessagesRequest) SetBizId(v string) *ListAIStaffChatMessagesRequest {
	s.BizId = &v
	return s
}

func (s *ListAIStaffChatMessagesRequest) SetConversationId(v string) *ListAIStaffChatMessagesRequest {
	s.ConversationId = &v
	return s
}

func (s *ListAIStaffChatMessagesRequest) SetPageSize(v int32) *ListAIStaffChatMessagesRequest {
	s.PageSize = &v
	return s
}

func (s *ListAIStaffChatMessagesRequest) SetStartCreateTime(v string) *ListAIStaffChatMessagesRequest {
	s.StartCreateTime = &v
	return s
}

func (s *ListAIStaffChatMessagesRequest) Validate() error {
	return dara.Validate(s)
}
