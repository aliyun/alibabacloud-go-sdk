// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListExecutionsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategories(v string) *ListExecutionsShrinkRequest
	GetCategories() *string
	SetCategory(v string) *ListExecutionsShrinkRequest
	GetCategory() *string
	SetDepth(v string) *ListExecutionsShrinkRequest
	GetDepth() *string
	SetDescription(v string) *ListExecutionsShrinkRequest
	GetDescription() *string
	SetEndDateAfter(v string) *ListExecutionsShrinkRequest
	GetEndDateAfter() *string
	SetEndDateBefore(v string) *ListExecutionsShrinkRequest
	GetEndDateBefore() *string
	SetExecutedBy(v string) *ListExecutionsShrinkRequest
	GetExecutedBy() *string
	SetExecutionId(v string) *ListExecutionsShrinkRequest
	GetExecutionId() *string
	SetIncludeChildExecution(v bool) *ListExecutionsShrinkRequest
	GetIncludeChildExecution() *bool
	SetMaxResults(v int32) *ListExecutionsShrinkRequest
	GetMaxResults() *int32
	SetMode(v string) *ListExecutionsShrinkRequest
	GetMode() *string
	SetNextToken(v string) *ListExecutionsShrinkRequest
	GetNextToken() *string
	SetParentExecutionId(v string) *ListExecutionsShrinkRequest
	GetParentExecutionId() *string
	SetRamRole(v string) *ListExecutionsShrinkRequest
	GetRamRole() *string
	SetRegionId(v string) *ListExecutionsShrinkRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ListExecutionsShrinkRequest
	GetResourceGroupId() *string
	SetResourceId(v string) *ListExecutionsShrinkRequest
	GetResourceId() *string
	SetResourceTemplateName(v string) *ListExecutionsShrinkRequest
	GetResourceTemplateName() *string
	SetSortField(v string) *ListExecutionsShrinkRequest
	GetSortField() *string
	SetSortOrder(v string) *ListExecutionsShrinkRequest
	GetSortOrder() *string
	SetStartDateAfter(v string) *ListExecutionsShrinkRequest
	GetStartDateAfter() *string
	SetStartDateBefore(v string) *ListExecutionsShrinkRequest
	GetStartDateBefore() *string
	SetStatus(v string) *ListExecutionsShrinkRequest
	GetStatus() *string
	SetTagsShrink(v string) *ListExecutionsShrinkRequest
	GetTagsShrink() *string
	SetTemplateName(v string) *ListExecutionsShrinkRequest
	GetTemplateName() *string
}

type ListExecutionsShrinkRequest struct {
	// The types of the execution template. Valid values: Other, TimerTrigger, EventTrigger, and AlarmTrigger. You can specify only one of the Categories and Category parameters. We recommend that you specify Categories.
	//
	// example:
	//
	// ["TimerTrigger"、"EventTrigger"]
	Categories *string `json:"Categories,omitempty" xml:"Categories,omitempty"`
	// The type of the execution template. Valid values: Other, TimerTrigger, EventTrigger, and AlarmTrigger.
	//
	// example:
	//
	// Other
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The depth of execution. Valid values: RootDepth and FirstChildDepth. If you set this parameter to RootDepth, only the parent execution is returned. If you set this parameter to FirstChildDepth, only the child executions at the first level are returned. You can specify only one of the Depth and IncludeChildExecution parameters. We recommend that you specify Depth.
	//
	// example:
	//
	// RootDepth
	Depth *string `json:"Depth,omitempty" xml:"Depth,omitempty"`
	// The description of the execution.
	//
	// example:
	//
	// MyDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The earliest end time. The executions that stop running at or later than the specified time are queried.
	//
	// example:
	//
	// 2019-05-16T10:26:14Z
	EndDateAfter *string `json:"EndDateAfter,omitempty" xml:"EndDateAfter,omitempty"`
	// The latest end time. The executions that stop running at or earlier than the specified time are queried.
	//
	// example:
	//
	// 2019-05-16T10:26:14Z
	EndDateBefore *string `json:"EndDateBefore,omitempty" xml:"EndDateBefore,omitempty"`
	// The executor.
	//
	// example:
	//
	// vme
	ExecutedBy *string `json:"ExecutedBy,omitempty" xml:"ExecutedBy,omitempty"`
	// The ID of the execution.
	//
	// example:
	//
	// exec-xxx
	ExecutionId *string `json:"ExecutionId,omitempty" xml:"ExecutionId,omitempty"`
	// Specifies whether to include child executions. Default value: False.
	//
	// example:
	//
	// true
	IncludeChildExecution *bool `json:"IncludeChildExecution,omitempty" xml:"IncludeChildExecution,omitempty"`
	// The number of entries to return on each page. Valid values: 10 to 100. Default value: 50.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The execution mode. Valid values:
	//
	// 	- **Automatic**
	//
	// 	- **Debug**
	//
	// example:
	//
	// Automatic
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The token that is used to retrieve the next page of results.
	//
	// example:
	//
	// MTRBMDc0NjAtRUJFNy00N0NBLTk3NTctMTJDQzQ
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the parent execution.
	//
	// example:
	//
	// exec-xxx
	ParentExecutionId *string `json:"ParentExecutionId,omitempty" xml:"ParentExecutionId,omitempty"`
	// The RAM role.
	//
	// example:
	//
	// OOSServiceRole
	RamRole *string `json:"RamRole,omitempty" xml:"RamRole,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the instances you want to query belong.
	//
	// example:
	//
	// rg-acfmxsn4m4******
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the Elastic Compute Service (ECS) resource.
	//
	// example:
	//
	// i-xxx
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The name of the resource template.
	//
	// example:
	//
	// ACS-ECS-TEST
	ResourceTemplateName *string `json:"ResourceTemplateName,omitempty" xml:"ResourceTemplateName,omitempty"`
	// The field that is used to sort the executions to query. Valid values:
	//
	// 	- **StartDate**: specifies that the executions are sorted based on the time when they are created. This is the default value.
	//
	// 	- **EndDate**: specifies that the executions are sorted based on the time when they stop running.
	//
	// 	- **Status**: specifies that the executions are sorted based on their states.
	//
	// example:
	//
	// StartDate
	SortField *string `json:"SortField,omitempty" xml:"SortField,omitempty"`
	// The order in which you want to sort the results. Valid values:
	//
	// 	- **Ascending**: ascending order.
	//
	// 	- **Descending**: descending order. This is the default value.
	//
	// example:
	//
	// Ascending
	SortOrder *string `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
	// The earliest start time. The executions that start to run at or later than the specified time are queried.
	//
	// example:
	//
	// 2019-05-16T10:26:14Z
	StartDateAfter *string `json:"StartDateAfter,omitempty" xml:"StartDateAfter,omitempty"`
	// The latest start time. The executions that start to run at or earlier than the specified point in time are queried.
	//
	// example:
	//
	// 2019-05-16T10:26:14Z
	StartDateBefore *string `json:"StartDateBefore,omitempty" xml:"StartDateBefore,omitempty"`
	// The status of the execution. Valid values: Running, Started, Success, Failed, Waiting, Cancelled, Pending, and Skipped.
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags for the execution.
	//
	// example:
	//
	// {"k1":"v2","k2":"v2"}
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The name of the template. All templates whose names contain the specified template name are queried.
	//
	// example:
	//
	// MyTemplate
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s ListExecutionsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListExecutionsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListExecutionsShrinkRequest) GetCategories() *string {
	return s.Categories
}

func (s *ListExecutionsShrinkRequest) GetCategory() *string {
	return s.Category
}

func (s *ListExecutionsShrinkRequest) GetDepth() *string {
	return s.Depth
}

func (s *ListExecutionsShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *ListExecutionsShrinkRequest) GetEndDateAfter() *string {
	return s.EndDateAfter
}

func (s *ListExecutionsShrinkRequest) GetEndDateBefore() *string {
	return s.EndDateBefore
}

func (s *ListExecutionsShrinkRequest) GetExecutedBy() *string {
	return s.ExecutedBy
}

func (s *ListExecutionsShrinkRequest) GetExecutionId() *string {
	return s.ExecutionId
}

func (s *ListExecutionsShrinkRequest) GetIncludeChildExecution() *bool {
	return s.IncludeChildExecution
}

func (s *ListExecutionsShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListExecutionsShrinkRequest) GetMode() *string {
	return s.Mode
}

func (s *ListExecutionsShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListExecutionsShrinkRequest) GetParentExecutionId() *string {
	return s.ParentExecutionId
}

func (s *ListExecutionsShrinkRequest) GetRamRole() *string {
	return s.RamRole
}

func (s *ListExecutionsShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListExecutionsShrinkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListExecutionsShrinkRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListExecutionsShrinkRequest) GetResourceTemplateName() *string {
	return s.ResourceTemplateName
}

func (s *ListExecutionsShrinkRequest) GetSortField() *string {
	return s.SortField
}

func (s *ListExecutionsShrinkRequest) GetSortOrder() *string {
	return s.SortOrder
}

func (s *ListExecutionsShrinkRequest) GetStartDateAfter() *string {
	return s.StartDateAfter
}

func (s *ListExecutionsShrinkRequest) GetStartDateBefore() *string {
	return s.StartDateBefore
}

func (s *ListExecutionsShrinkRequest) GetStatus() *string {
	return s.Status
}

func (s *ListExecutionsShrinkRequest) GetTagsShrink() *string {
	return s.TagsShrink
}

func (s *ListExecutionsShrinkRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *ListExecutionsShrinkRequest) SetCategories(v string) *ListExecutionsShrinkRequest {
	s.Categories = &v
	return s
}

func (s *ListExecutionsShrinkRequest) SetCategory(v string) *ListExecutionsShrinkRequest {
	s.Category = &v
	return s
}

func (s *ListExecutionsShrinkRequest) SetDepth(v string) *ListExecutionsShrinkRequest {
	s.Depth = &v
	return s
}

func (s *ListExecutionsShrinkRequest) SetDescription(v string) *ListExecutionsShrinkRequest {
	s.Description = &v
	return s
}

func (s *ListExecutionsShrinkRequest) SetEndDateAfter(v string) *ListExecutionsShrinkRequest {
	s.EndDateAfter = &v
	return s
}

func (s *ListExecutionsShrinkRequest) SetEndDateBefore(v string) *ListExecutionsShrinkRequest {
	s.EndDateBefore = &v
	return s
}

func (s *ListExecutionsShrinkRequest) SetExecutedBy(v string) *ListExecutionsShrinkRequest {
	s.ExecutedBy = &v
	return s
}

func (s *ListExecutionsShrinkRequest) SetExecutionId(v string) *ListExecutionsShrinkRequest {
	s.ExecutionId = &v
	return s
}

func (s *ListExecutionsShrinkRequest) SetIncludeChildExecution(v bool) *ListExecutionsShrinkRequest {
	s.IncludeChildExecution = &v
	return s
}

func (s *ListExecutionsShrinkRequest) SetMaxResults(v int32) *ListExecutionsShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListExecutionsShrinkRequest) SetMode(v string) *ListExecutionsShrinkRequest {
	s.Mode = &v
	return s
}

func (s *ListExecutionsShrinkRequest) SetNextToken(v string) *ListExecutionsShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListExecutionsShrinkRequest) SetParentExecutionId(v string) *ListExecutionsShrinkRequest {
	s.ParentExecutionId = &v
	return s
}

func (s *ListExecutionsShrinkRequest) SetRamRole(v string) *ListExecutionsShrinkRequest {
	s.RamRole = &v
	return s
}

func (s *ListExecutionsShrinkRequest) SetRegionId(v string) *ListExecutionsShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ListExecutionsShrinkRequest) SetResourceGroupId(v string) *ListExecutionsShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListExecutionsShrinkRequest) SetResourceId(v string) *ListExecutionsShrinkRequest {
	s.ResourceId = &v
	return s
}

func (s *ListExecutionsShrinkRequest) SetResourceTemplateName(v string) *ListExecutionsShrinkRequest {
	s.ResourceTemplateName = &v
	return s
}

func (s *ListExecutionsShrinkRequest) SetSortField(v string) *ListExecutionsShrinkRequest {
	s.SortField = &v
	return s
}

func (s *ListExecutionsShrinkRequest) SetSortOrder(v string) *ListExecutionsShrinkRequest {
	s.SortOrder = &v
	return s
}

func (s *ListExecutionsShrinkRequest) SetStartDateAfter(v string) *ListExecutionsShrinkRequest {
	s.StartDateAfter = &v
	return s
}

func (s *ListExecutionsShrinkRequest) SetStartDateBefore(v string) *ListExecutionsShrinkRequest {
	s.StartDateBefore = &v
	return s
}

func (s *ListExecutionsShrinkRequest) SetStatus(v string) *ListExecutionsShrinkRequest {
	s.Status = &v
	return s
}

func (s *ListExecutionsShrinkRequest) SetTagsShrink(v string) *ListExecutionsShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *ListExecutionsShrinkRequest) SetTemplateName(v string) *ListExecutionsShrinkRequest {
	s.TemplateName = &v
	return s
}

func (s *ListExecutionsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
