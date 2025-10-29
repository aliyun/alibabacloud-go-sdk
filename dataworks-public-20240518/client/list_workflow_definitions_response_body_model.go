// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkflowDefinitionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPagingInfo(v *ListWorkflowDefinitionsResponseBodyPagingInfo) *ListWorkflowDefinitionsResponseBody
	GetPagingInfo() *ListWorkflowDefinitionsResponseBodyPagingInfo
	SetRequestId(v string) *ListWorkflowDefinitionsResponseBody
	GetRequestId() *string
}

type ListWorkflowDefinitionsResponseBody struct {
	// The pagination information.
	PagingInfo *ListWorkflowDefinitionsResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 8C3ED0C5-ABAB-55E1-854B-DAC02B11XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListWorkflowDefinitionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListWorkflowDefinitionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListWorkflowDefinitionsResponseBody) GetPagingInfo() *ListWorkflowDefinitionsResponseBodyPagingInfo {
	return s.PagingInfo
}

func (s *ListWorkflowDefinitionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListWorkflowDefinitionsResponseBody) SetPagingInfo(v *ListWorkflowDefinitionsResponseBodyPagingInfo) *ListWorkflowDefinitionsResponseBody {
	s.PagingInfo = v
	return s
}

func (s *ListWorkflowDefinitionsResponseBody) SetRequestId(v string) *ListWorkflowDefinitionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWorkflowDefinitionsResponseBody) Validate() error {
	if s.PagingInfo != nil {
		if err := s.PagingInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListWorkflowDefinitionsResponseBodyPagingInfo struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 227
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The workflows.
	WorkflowDefinitions []*ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitions `json:"WorkflowDefinitions,omitempty" xml:"WorkflowDefinitions,omitempty" type:"Repeated"`
}

func (s ListWorkflowDefinitionsResponseBodyPagingInfo) String() string {
	return dara.Prettify(s)
}

func (s ListWorkflowDefinitionsResponseBodyPagingInfo) GoString() string {
	return s.String()
}

func (s *ListWorkflowDefinitionsResponseBodyPagingInfo) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListWorkflowDefinitionsResponseBodyPagingInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListWorkflowDefinitionsResponseBodyPagingInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListWorkflowDefinitionsResponseBodyPagingInfo) GetWorkflowDefinitions() []*ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitions {
	return s.WorkflowDefinitions
}

func (s *ListWorkflowDefinitionsResponseBodyPagingInfo) SetPageNumber(v int32) *ListWorkflowDefinitionsResponseBodyPagingInfo {
	s.PageNumber = &v
	return s
}

func (s *ListWorkflowDefinitionsResponseBodyPagingInfo) SetPageSize(v int32) *ListWorkflowDefinitionsResponseBodyPagingInfo {
	s.PageSize = &v
	return s
}

func (s *ListWorkflowDefinitionsResponseBodyPagingInfo) SetTotalCount(v int32) *ListWorkflowDefinitionsResponseBodyPagingInfo {
	s.TotalCount = &v
	return s
}

func (s *ListWorkflowDefinitionsResponseBodyPagingInfo) SetWorkflowDefinitions(v []*ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitions) *ListWorkflowDefinitionsResponseBodyPagingInfo {
	s.WorkflowDefinitions = v
	return s
}

func (s *ListWorkflowDefinitionsResponseBodyPagingInfo) Validate() error {
	if s.WorkflowDefinitions != nil {
		for _, item := range s.WorkflowDefinitions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitions struct {
	// The time when the workflow was created. This value is a UNIX timestamp.
	//
	// example:
	//
	// 1698057323000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the workflow.
	//
	// example:
	//
	// Workflow description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the workflow.
	//
	// example:
	//
	// 463497880880954XXXX
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The time when the workflow was last modified. This value is a UNIX timestamp.
	//
	// example:
	//
	// 1698057323000
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The name of the workflow.
	//
	// example:
	//
	// OpenAPI test workflow Demo
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The owner.
	//
	// example:
	//
	// 110755000425XXXX
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The ID of the DataWorks workspace to which the workflow belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// 4710
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The script information.
	Script *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitionsScript `json:"Script,omitempty" xml:"Script,omitempty" type:"Struct"`
	// The type of the workflow.
	//
	// Valid values:
	//
	// 	- CycleWorkflow
	//
	// 	- ManualWorkflow
	//
	// example:
	//
	// CycleWorkflow
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitions) String() string {
	return dara.Prettify(s)
}

func (s ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitions) GoString() string {
	return s.String()
}

func (s *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitions) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitions) GetDescription() *string {
	return s.Description
}

func (s *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitions) GetId() *int64 {
	return s.Id
}

func (s *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitions) GetModifyTime() *int64 {
	return s.ModifyTime
}

func (s *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitions) GetName() *string {
	return s.Name
}

func (s *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitions) GetOwner() *string {
	return s.Owner
}

func (s *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitions) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitions) GetScript() *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitionsScript {
	return s.Script
}

func (s *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitions) GetType() *string {
	return s.Type
}

func (s *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitions) SetCreateTime(v int64) *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitions {
	s.CreateTime = &v
	return s
}

func (s *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitions) SetDescription(v string) *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitions {
	s.Description = &v
	return s
}

func (s *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitions) SetId(v int64) *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitions {
	s.Id = &v
	return s
}

func (s *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitions) SetModifyTime(v int64) *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitions {
	s.ModifyTime = &v
	return s
}

func (s *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitions) SetName(v string) *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitions {
	s.Name = &v
	return s
}

func (s *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitions) SetOwner(v string) *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitions {
	s.Owner = &v
	return s
}

func (s *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitions) SetProjectId(v int64) *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitions {
	s.ProjectId = &v
	return s
}

func (s *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitions) SetScript(v *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitionsScript) *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitions {
	s.Script = v
	return s
}

func (s *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitions) SetType(v string) *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitions {
	s.Type = &v
	return s
}

func (s *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitions) Validate() error {
	if s.Script != nil {
		if err := s.Script.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitionsScript struct {
	// The script ID.
	//
	// example:
	//
	// 698002781368644XXXX
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The script path.
	//
	// example:
	//
	// XX/OpenAPI_test/workflow_test/OpenAPI_test_workflow_Demo
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The runtime.
	Runtime *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitionsScriptRuntime `json:"Runtime,omitempty" xml:"Runtime,omitempty" type:"Struct"`
}

func (s ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitionsScript) String() string {
	return dara.Prettify(s)
}

func (s ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitionsScript) GoString() string {
	return s.String()
}

func (s *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitionsScript) GetId() *int64 {
	return s.Id
}

func (s *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitionsScript) GetPath() *string {
	return s.Path
}

func (s *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitionsScript) GetRuntime() *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitionsScriptRuntime {
	return s.Runtime
}

func (s *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitionsScript) SetId(v int64) *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitionsScript {
	s.Id = &v
	return s
}

func (s *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitionsScript) SetPath(v string) *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitionsScript {
	s.Path = &v
	return s
}

func (s *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitionsScript) SetRuntime(v *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitionsScriptRuntime) *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitionsScript {
	s.Runtime = v
	return s
}

func (s *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitionsScript) Validate() error {
	if s.Runtime != nil {
		if err := s.Runtime.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitionsScriptRuntime struct {
	// The command.
	//
	// example:
	//
	// WORKFLOW
	Command *string `json:"Command,omitempty" xml:"Command,omitempty"`
}

func (s ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitionsScriptRuntime) String() string {
	return dara.Prettify(s)
}

func (s ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitionsScriptRuntime) GoString() string {
	return s.String()
}

func (s *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitionsScriptRuntime) GetCommand() *string {
	return s.Command
}

func (s *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitionsScriptRuntime) SetCommand(v string) *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitionsScriptRuntime {
	s.Command = &v
	return s
}

func (s *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitionsScriptRuntime) Validate() error {
	return dara.Validate(s)
}
