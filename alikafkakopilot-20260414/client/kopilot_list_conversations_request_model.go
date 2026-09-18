// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iKopilotListConversationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDestinationCursor(v string) *KopilotListConversationsRequest
	GetDestinationCursor() *string
	SetDestinationPageSize(v int32) *KopilotListConversationsRequest
	GetDestinationPageSize() *int32
	SetIncludeAutomationOverview(v bool) *KopilotListConversationsRequest
	GetIncludeAutomationOverview() *bool
	SetPage(v int32) *KopilotListConversationsRequest
	GetPage() *int32
	SetRegionId(v string) *KopilotListConversationsRequest
	GetRegionId() *string
	SetSize(v int32) *KopilotListConversationsRequest
	GetSize() *int32
	SetTaskCursor(v string) *KopilotListConversationsRequest
	GetTaskCursor() *string
	SetTaskPageSize(v int32) *KopilotListConversationsRequest
	GetTaskPageSize() *int32
}

type KopilotListConversationsRequest struct {
	// The pagination cursor for notification channels. Do not specify this parameter for the first query. For subsequent queries, pass in the value of Data.AutomationOverview.Destinations.NextCursor from the previous response.
	//
	// example:
	//
	// 123
	DestinationCursor *string `json:"DestinationCursor,omitempty" xml:"DestinationCursor,omitempty"`
	// The number of entries per page for the notification channel list. Default value: 20. Valid values: 1 to 100.
	//
	// example:
	//
	// 20
	DestinationPageSize *int32 `json:"DestinationPageSize,omitempty" xml:"DestinationPageSize,omitempty"`
	// Specifies whether to return the overview of scheduled tasks and notification channels for the account. A value of true indicates that the overview is returned. If this parameter is not specified, the overview is not returned.
	//
	// example:
	//
	// true
	IncludeAutomationOverview *bool `json:"IncludeAutomationOverview,omitempty" xml:"IncludeAutomationOverview,omitempty"`
	// The current page number.
	//
	// example:
	//
	// 1
	Page *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The page size.
	//
	// example:
	//
	// 10
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The pagination cursor. Do not specify this parameter for the first query. For subsequent queries, pass in the value of Data.AutomationOverview.Tasks.NextCursor from the previous response.
	//
	// example:
	//
	// 123
	TaskCursor *string `json:"TaskCursor,omitempty" xml:"TaskCursor,omitempty"`
	// The number of entries per page for the scheduled task list. Default value: 20. Valid values: 1 to 100.
	//
	// example:
	//
	// 20
	TaskPageSize *int32 `json:"TaskPageSize,omitempty" xml:"TaskPageSize,omitempty"`
}

func (s KopilotListConversationsRequest) String() string {
	return dara.Prettify(s)
}

func (s KopilotListConversationsRequest) GoString() string {
	return s.String()
}

func (s *KopilotListConversationsRequest) GetDestinationCursor() *string {
	return s.DestinationCursor
}

func (s *KopilotListConversationsRequest) GetDestinationPageSize() *int32 {
	return s.DestinationPageSize
}

func (s *KopilotListConversationsRequest) GetIncludeAutomationOverview() *bool {
	return s.IncludeAutomationOverview
}

func (s *KopilotListConversationsRequest) GetPage() *int32 {
	return s.Page
}

func (s *KopilotListConversationsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *KopilotListConversationsRequest) GetSize() *int32 {
	return s.Size
}

func (s *KopilotListConversationsRequest) GetTaskCursor() *string {
	return s.TaskCursor
}

func (s *KopilotListConversationsRequest) GetTaskPageSize() *int32 {
	return s.TaskPageSize
}

func (s *KopilotListConversationsRequest) SetDestinationCursor(v string) *KopilotListConversationsRequest {
	s.DestinationCursor = &v
	return s
}

func (s *KopilotListConversationsRequest) SetDestinationPageSize(v int32) *KopilotListConversationsRequest {
	s.DestinationPageSize = &v
	return s
}

func (s *KopilotListConversationsRequest) SetIncludeAutomationOverview(v bool) *KopilotListConversationsRequest {
	s.IncludeAutomationOverview = &v
	return s
}

func (s *KopilotListConversationsRequest) SetPage(v int32) *KopilotListConversationsRequest {
	s.Page = &v
	return s
}

func (s *KopilotListConversationsRequest) SetRegionId(v string) *KopilotListConversationsRequest {
	s.RegionId = &v
	return s
}

func (s *KopilotListConversationsRequest) SetSize(v int32) *KopilotListConversationsRequest {
	s.Size = &v
	return s
}

func (s *KopilotListConversationsRequest) SetTaskCursor(v string) *KopilotListConversationsRequest {
	s.TaskCursor = &v
	return s
}

func (s *KopilotListConversationsRequest) SetTaskPageSize(v int32) *KopilotListConversationsRequest {
	s.TaskPageSize = &v
	return s
}

func (s *KopilotListConversationsRequest) Validate() error {
	return dara.Validate(s)
}
