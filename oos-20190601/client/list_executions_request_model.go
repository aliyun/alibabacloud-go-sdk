// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListExecutionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategories(v string) *ListExecutionsRequest
	GetCategories() *string
	SetCategory(v string) *ListExecutionsRequest
	GetCategory() *string
	SetDepth(v string) *ListExecutionsRequest
	GetDepth() *string
	SetDescription(v string) *ListExecutionsRequest
	GetDescription() *string
	SetEndDateAfter(v string) *ListExecutionsRequest
	GetEndDateAfter() *string
	SetEndDateBefore(v string) *ListExecutionsRequest
	GetEndDateBefore() *string
	SetExecutedBy(v string) *ListExecutionsRequest
	GetExecutedBy() *string
	SetExecutionId(v string) *ListExecutionsRequest
	GetExecutionId() *string
	SetIncludeChildExecution(v bool) *ListExecutionsRequest
	GetIncludeChildExecution() *bool
	SetMaxResults(v int32) *ListExecutionsRequest
	GetMaxResults() *int32
	SetMode(v string) *ListExecutionsRequest
	GetMode() *string
	SetNextToken(v string) *ListExecutionsRequest
	GetNextToken() *string
	SetParentExecutionId(v string) *ListExecutionsRequest
	GetParentExecutionId() *string
	SetRamRole(v string) *ListExecutionsRequest
	GetRamRole() *string
	SetRegionId(v string) *ListExecutionsRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ListExecutionsRequest
	GetResourceGroupId() *string
	SetResourceId(v string) *ListExecutionsRequest
	GetResourceId() *string
	SetResourceTemplateName(v string) *ListExecutionsRequest
	GetResourceTemplateName() *string
	SetSortField(v string) *ListExecutionsRequest
	GetSortField() *string
	SetSortOrder(v string) *ListExecutionsRequest
	GetSortOrder() *string
	SetStartDateAfter(v string) *ListExecutionsRequest
	GetStartDateAfter() *string
	SetStartDateBefore(v string) *ListExecutionsRequest
	GetStartDateBefore() *string
	SetStatus(v string) *ListExecutionsRequest
	GetStatus() *string
	SetTags(v map[string]interface{}) *ListExecutionsRequest
	GetTags() map[string]interface{}
	SetTemplateName(v string) *ListExecutionsRequest
	GetTemplateName() *string
}

type ListExecutionsRequest struct {
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
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The name of the template. All templates whose names contain the specified template name are queried.
	//
	// example:
	//
	// MyTemplate
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s ListExecutionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListExecutionsRequest) GoString() string {
	return s.String()
}

func (s *ListExecutionsRequest) GetCategories() *string {
	return s.Categories
}

func (s *ListExecutionsRequest) GetCategory() *string {
	return s.Category
}

func (s *ListExecutionsRequest) GetDepth() *string {
	return s.Depth
}

func (s *ListExecutionsRequest) GetDescription() *string {
	return s.Description
}

func (s *ListExecutionsRequest) GetEndDateAfter() *string {
	return s.EndDateAfter
}

func (s *ListExecutionsRequest) GetEndDateBefore() *string {
	return s.EndDateBefore
}

func (s *ListExecutionsRequest) GetExecutedBy() *string {
	return s.ExecutedBy
}

func (s *ListExecutionsRequest) GetExecutionId() *string {
	return s.ExecutionId
}

func (s *ListExecutionsRequest) GetIncludeChildExecution() *bool {
	return s.IncludeChildExecution
}

func (s *ListExecutionsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListExecutionsRequest) GetMode() *string {
	return s.Mode
}

func (s *ListExecutionsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListExecutionsRequest) GetParentExecutionId() *string {
	return s.ParentExecutionId
}

func (s *ListExecutionsRequest) GetRamRole() *string {
	return s.RamRole
}

func (s *ListExecutionsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListExecutionsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListExecutionsRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListExecutionsRequest) GetResourceTemplateName() *string {
	return s.ResourceTemplateName
}

func (s *ListExecutionsRequest) GetSortField() *string {
	return s.SortField
}

func (s *ListExecutionsRequest) GetSortOrder() *string {
	return s.SortOrder
}

func (s *ListExecutionsRequest) GetStartDateAfter() *string {
	return s.StartDateAfter
}

func (s *ListExecutionsRequest) GetStartDateBefore() *string {
	return s.StartDateBefore
}

func (s *ListExecutionsRequest) GetStatus() *string {
	return s.Status
}

func (s *ListExecutionsRequest) GetTags() map[string]interface{} {
	return s.Tags
}

func (s *ListExecutionsRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *ListExecutionsRequest) SetCategories(v string) *ListExecutionsRequest {
	s.Categories = &v
	return s
}

func (s *ListExecutionsRequest) SetCategory(v string) *ListExecutionsRequest {
	s.Category = &v
	return s
}

func (s *ListExecutionsRequest) SetDepth(v string) *ListExecutionsRequest {
	s.Depth = &v
	return s
}

func (s *ListExecutionsRequest) SetDescription(v string) *ListExecutionsRequest {
	s.Description = &v
	return s
}

func (s *ListExecutionsRequest) SetEndDateAfter(v string) *ListExecutionsRequest {
	s.EndDateAfter = &v
	return s
}

func (s *ListExecutionsRequest) SetEndDateBefore(v string) *ListExecutionsRequest {
	s.EndDateBefore = &v
	return s
}

func (s *ListExecutionsRequest) SetExecutedBy(v string) *ListExecutionsRequest {
	s.ExecutedBy = &v
	return s
}

func (s *ListExecutionsRequest) SetExecutionId(v string) *ListExecutionsRequest {
	s.ExecutionId = &v
	return s
}

func (s *ListExecutionsRequest) SetIncludeChildExecution(v bool) *ListExecutionsRequest {
	s.IncludeChildExecution = &v
	return s
}

func (s *ListExecutionsRequest) SetMaxResults(v int32) *ListExecutionsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListExecutionsRequest) SetMode(v string) *ListExecutionsRequest {
	s.Mode = &v
	return s
}

func (s *ListExecutionsRequest) SetNextToken(v string) *ListExecutionsRequest {
	s.NextToken = &v
	return s
}

func (s *ListExecutionsRequest) SetParentExecutionId(v string) *ListExecutionsRequest {
	s.ParentExecutionId = &v
	return s
}

func (s *ListExecutionsRequest) SetRamRole(v string) *ListExecutionsRequest {
	s.RamRole = &v
	return s
}

func (s *ListExecutionsRequest) SetRegionId(v string) *ListExecutionsRequest {
	s.RegionId = &v
	return s
}

func (s *ListExecutionsRequest) SetResourceGroupId(v string) *ListExecutionsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListExecutionsRequest) SetResourceId(v string) *ListExecutionsRequest {
	s.ResourceId = &v
	return s
}

func (s *ListExecutionsRequest) SetResourceTemplateName(v string) *ListExecutionsRequest {
	s.ResourceTemplateName = &v
	return s
}

func (s *ListExecutionsRequest) SetSortField(v string) *ListExecutionsRequest {
	s.SortField = &v
	return s
}

func (s *ListExecutionsRequest) SetSortOrder(v string) *ListExecutionsRequest {
	s.SortOrder = &v
	return s
}

func (s *ListExecutionsRequest) SetStartDateAfter(v string) *ListExecutionsRequest {
	s.StartDateAfter = &v
	return s
}

func (s *ListExecutionsRequest) SetStartDateBefore(v string) *ListExecutionsRequest {
	s.StartDateBefore = &v
	return s
}

func (s *ListExecutionsRequest) SetStatus(v string) *ListExecutionsRequest {
	s.Status = &v
	return s
}

func (s *ListExecutionsRequest) SetTags(v map[string]interface{}) *ListExecutionsRequest {
	s.Tags = v
	return s
}

func (s *ListExecutionsRequest) SetTemplateName(v string) *ListExecutionsRequest {
	s.TemplateName = &v
	return s
}

func (s *ListExecutionsRequest) Validate() error {
	return dara.Validate(s)
}
