// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTodoOpsTaskApprovalsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListTodoOpsTaskApprovalsRequest
	GetInstanceId() *string
	SetKeyword(v string) *ListTodoOpsTaskApprovalsRequest
	GetKeyword() *string
	SetPageNumber(v string) *ListTodoOpsTaskApprovalsRequest
	GetPageNumber() *string
	SetPageSize(v string) *ListTodoOpsTaskApprovalsRequest
	GetPageSize() *string
	SetProjectId(v int64) *ListTodoOpsTaskApprovalsRequest
	GetProjectId() *int64
	SetRegionId(v string) *ListTodoOpsTaskApprovalsRequest
	GetRegionId() *string
	SetScheduleType(v string) *ListTodoOpsTaskApprovalsRequest
	GetScheduleType() *string
}

type ListTodoOpsTaskApprovalsRequest struct {
	// The instance ID of the Bastionhost instance.
	//
	// > You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to obtain this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-m413tuhlo03
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name or comment of the task to query. Fuzzy match is supported.
	//
	// example:
	//
	// test
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The page number of the current page in a paging query. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The maximum number of entries to return on each page in a paging query.
	//
	// The maximum value of the PageSize parameter is 100. The default number of entries per page is 20. If the PageSize parameter is left empty, 20 entries are returned by default.
	//
	// > Specify a value for the PageSize parameter.
	//
	// example:
	//
	// 20
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The project ID.
	//
	// example:
	//
	// 3
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The region ID of the Bastionhost instance.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The scheduling type of the task. Valid values:
	//
	// - **FixTime**: Scheduled execution.
	//
	// - **CycleInterval**: Periodic execution.
	//
	// - **Manual**: Manually triggered execution.
	//
	// example:
	//
	// Manual
	ScheduleType *string `json:"ScheduleType,omitempty" xml:"ScheduleType,omitempty"`
}

func (s ListTodoOpsTaskApprovalsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTodoOpsTaskApprovalsRequest) GoString() string {
	return s.String()
}

func (s *ListTodoOpsTaskApprovalsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListTodoOpsTaskApprovalsRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListTodoOpsTaskApprovalsRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *ListTodoOpsTaskApprovalsRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *ListTodoOpsTaskApprovalsRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListTodoOpsTaskApprovalsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListTodoOpsTaskApprovalsRequest) GetScheduleType() *string {
	return s.ScheduleType
}

func (s *ListTodoOpsTaskApprovalsRequest) SetInstanceId(v string) *ListTodoOpsTaskApprovalsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListTodoOpsTaskApprovalsRequest) SetKeyword(v string) *ListTodoOpsTaskApprovalsRequest {
	s.Keyword = &v
	return s
}

func (s *ListTodoOpsTaskApprovalsRequest) SetPageNumber(v string) *ListTodoOpsTaskApprovalsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTodoOpsTaskApprovalsRequest) SetPageSize(v string) *ListTodoOpsTaskApprovalsRequest {
	s.PageSize = &v
	return s
}

func (s *ListTodoOpsTaskApprovalsRequest) SetProjectId(v int64) *ListTodoOpsTaskApprovalsRequest {
	s.ProjectId = &v
	return s
}

func (s *ListTodoOpsTaskApprovalsRequest) SetRegionId(v string) *ListTodoOpsTaskApprovalsRequest {
	s.RegionId = &v
	return s
}

func (s *ListTodoOpsTaskApprovalsRequest) SetScheduleType(v string) *ListTodoOpsTaskApprovalsRequest {
	s.ScheduleType = &v
	return s
}

func (s *ListTodoOpsTaskApprovalsRequest) Validate() error {
	return dara.Validate(s)
}
