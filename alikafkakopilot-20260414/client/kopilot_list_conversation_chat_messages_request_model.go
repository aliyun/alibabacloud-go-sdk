// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iKopilotListConversationChatMessagesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBeforeTurnId(v int32) *KopilotListConversationChatMessagesRequest
	GetBeforeTurnId() *int32
	SetPageSize(v int32) *KopilotListConversationChatMessagesRequest
	GetPageSize() *int32
	SetRegionId(v string) *KopilotListConversationChatMessagesRequest
	GetRegionId() *string
	SetSessionId(v string) *KopilotListConversationChatMessagesRequest
	GetSessionId() *string
	SetTaskCursor(v string) *KopilotListConversationChatMessagesRequest
	GetTaskCursor() *string
	SetTaskPageSize(v int32) *KopilotListConversationChatMessagesRequest
	GetTaskPageSize() *int32
}

type KopilotListConversationChatMessagesRequest struct {
	// The cursor.
	//
	// > If this parameter is not specified, the last pageSize turn IDs are returned. If this parameter is specified, the turn IDs before the specified turn ID are returned.
	//
	// example:
	//
	// 2345
	BeforeTurnId *int32 `json:"BeforeTurnId,omitempty" xml:"BeforeTurnId,omitempty"`
	// The page size.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The session ID.
	//
	// example:
	//
	// e356c91c-8220-425c-9d86-********
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// The pagination cursor. Do not specify this parameter for the first query. For subsequent queries, pass in the value of Data.ScheduledTaskInfo.NextTaskCursor from the previous response.
	//
	// example:
	//
	// 123
	TaskCursor *string `json:"TaskCursor,omitempty" xml:"TaskCursor,omitempty"`
	// The number of scheduled tasks per page. Default value: 20. Valid values: 1 to 100.
	//
	// example:
	//
	// 20
	TaskPageSize *int32 `json:"TaskPageSize,omitempty" xml:"TaskPageSize,omitempty"`
}

func (s KopilotListConversationChatMessagesRequest) String() string {
	return dara.Prettify(s)
}

func (s KopilotListConversationChatMessagesRequest) GoString() string {
	return s.String()
}

func (s *KopilotListConversationChatMessagesRequest) GetBeforeTurnId() *int32 {
	return s.BeforeTurnId
}

func (s *KopilotListConversationChatMessagesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *KopilotListConversationChatMessagesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *KopilotListConversationChatMessagesRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *KopilotListConversationChatMessagesRequest) GetTaskCursor() *string {
	return s.TaskCursor
}

func (s *KopilotListConversationChatMessagesRequest) GetTaskPageSize() *int32 {
	return s.TaskPageSize
}

func (s *KopilotListConversationChatMessagesRequest) SetBeforeTurnId(v int32) *KopilotListConversationChatMessagesRequest {
	s.BeforeTurnId = &v
	return s
}

func (s *KopilotListConversationChatMessagesRequest) SetPageSize(v int32) *KopilotListConversationChatMessagesRequest {
	s.PageSize = &v
	return s
}

func (s *KopilotListConversationChatMessagesRequest) SetRegionId(v string) *KopilotListConversationChatMessagesRequest {
	s.RegionId = &v
	return s
}

func (s *KopilotListConversationChatMessagesRequest) SetSessionId(v string) *KopilotListConversationChatMessagesRequest {
	s.SessionId = &v
	return s
}

func (s *KopilotListConversationChatMessagesRequest) SetTaskCursor(v string) *KopilotListConversationChatMessagesRequest {
	s.TaskCursor = &v
	return s
}

func (s *KopilotListConversationChatMessagesRequest) SetTaskPageSize(v int32) *KopilotListConversationChatMessagesRequest {
	s.TaskPageSize = &v
	return s
}

func (s *KopilotListConversationChatMessagesRequest) Validate() error {
	return dara.Validate(s)
}
