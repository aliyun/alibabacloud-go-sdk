// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSiteMonitorListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentGroup(v string) *DescribeSiteMonitorListRequest
	GetAgentGroup() *string
	SetKeyword(v string) *DescribeSiteMonitorListRequest
	GetKeyword() *string
	SetPage(v int32) *DescribeSiteMonitorListRequest
	GetPage() *int32
	SetPageSize(v int32) *DescribeSiteMonitorListRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeSiteMonitorListRequest
	GetRegionId() *string
	SetTaskId(v string) *DescribeSiteMonitorListRequest
	GetTaskId() *string
	SetTaskState(v string) *DescribeSiteMonitorListRequest
	GetTaskState() *string
	SetTaskType(v string) *DescribeSiteMonitorListRequest
	GetTaskType() *string
}

type DescribeSiteMonitorListRequest struct {
	// Task network type. Valid values:
	//
	// - PC: Cable Network
	//
	// - MOBILE: Mobile Cellular Network
	//
	// - FC: Alibaba Cloud VPC Network
	//
	// example:
	//
	// PC
	AgentGroup *string `json:"AgentGroup,omitempty" xml:"AgentGroup,omitempty"`
	// The keyword to be matched.
	//
	// >  You can search for tasks by name or address. Fuzzy search is supported.
	//
	// example:
	//
	// site
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	Page *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
	// The number of entries per page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the site monitoring task.
	//
	// example:
	//
	// a1ecd34a-8157-44d9-b060-14950837****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The task status. Valid values:
	//
	// 	- 1: The task is enabled.
	//
	// 	- 2: The task is disabled.
	//
	// example:
	//
	// 1
	TaskState *string `json:"TaskState,omitempty" xml:"TaskState,omitempty"`
	// The protocol that is used by the site monitoring task. Valid values: HTTP, PING, TCP, UDP, DNS, SMTP, POP3, and FTP.
	//
	// example:
	//
	// HTTP
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s DescribeSiteMonitorListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSiteMonitorListRequest) GoString() string {
	return s.String()
}

func (s *DescribeSiteMonitorListRequest) GetAgentGroup() *string {
	return s.AgentGroup
}

func (s *DescribeSiteMonitorListRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *DescribeSiteMonitorListRequest) GetPage() *int32 {
	return s.Page
}

func (s *DescribeSiteMonitorListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSiteMonitorListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSiteMonitorListRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeSiteMonitorListRequest) GetTaskState() *string {
	return s.TaskState
}

func (s *DescribeSiteMonitorListRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *DescribeSiteMonitorListRequest) SetAgentGroup(v string) *DescribeSiteMonitorListRequest {
	s.AgentGroup = &v
	return s
}

func (s *DescribeSiteMonitorListRequest) SetKeyword(v string) *DescribeSiteMonitorListRequest {
	s.Keyword = &v
	return s
}

func (s *DescribeSiteMonitorListRequest) SetPage(v int32) *DescribeSiteMonitorListRequest {
	s.Page = &v
	return s
}

func (s *DescribeSiteMonitorListRequest) SetPageSize(v int32) *DescribeSiteMonitorListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSiteMonitorListRequest) SetRegionId(v string) *DescribeSiteMonitorListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSiteMonitorListRequest) SetTaskId(v string) *DescribeSiteMonitorListRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeSiteMonitorListRequest) SetTaskState(v string) *DescribeSiteMonitorListRequest {
	s.TaskState = &v
	return s
}

func (s *DescribeSiteMonitorListRequest) SetTaskType(v string) *DescribeSiteMonitorListRequest {
	s.TaskType = &v
	return s
}

func (s *DescribeSiteMonitorListRequest) Validate() error {
	return dara.Validate(s)
}
