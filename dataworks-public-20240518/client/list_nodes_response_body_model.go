// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNodesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPagingInfo(v *ListNodesResponseBodyPagingInfo) *ListNodesResponseBody
	GetPagingInfo() *ListNodesResponseBodyPagingInfo
	SetRequestId(v string) *ListNodesResponseBody
	GetRequestId() *string
}

type ListNodesResponseBody struct {
	// The pagination information.
	PagingInfo *ListNodesResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 2197B9C4-39CE-55EA-8EEA-FDBAE52DXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListNodesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListNodesResponseBody) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBody) GetPagingInfo() *ListNodesResponseBodyPagingInfo {
	return s.PagingInfo
}

func (s *ListNodesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListNodesResponseBody) SetPagingInfo(v *ListNodesResponseBodyPagingInfo) *ListNodesResponseBody {
	s.PagingInfo = v
	return s
}

func (s *ListNodesResponseBody) SetRequestId(v string) *ListNodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNodesResponseBody) Validate() error {
	if s.PagingInfo != nil {
		if err := s.PagingInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListNodesResponseBodyPagingInfo struct {
	// The list of nodes in DataStudio.
	Nodes []*ListNodesResponseBodyPagingInfoNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 42
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListNodesResponseBodyPagingInfo) String() string {
	return dara.Prettify(s)
}

func (s ListNodesResponseBodyPagingInfo) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyPagingInfo) GetNodes() []*ListNodesResponseBodyPagingInfoNodes {
	return s.Nodes
}

func (s *ListNodesResponseBodyPagingInfo) GetPageNumber() *string {
	return s.PageNumber
}

func (s *ListNodesResponseBodyPagingInfo) GetPageSize() *string {
	return s.PageSize
}

func (s *ListNodesResponseBodyPagingInfo) GetTotalCount() *string {
	return s.TotalCount
}

func (s *ListNodesResponseBodyPagingInfo) SetNodes(v []*ListNodesResponseBodyPagingInfoNodes) *ListNodesResponseBodyPagingInfo {
	s.Nodes = v
	return s
}

func (s *ListNodesResponseBodyPagingInfo) SetPageNumber(v string) *ListNodesResponseBodyPagingInfo {
	s.PageNumber = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfo) SetPageSize(v string) *ListNodesResponseBodyPagingInfo {
	s.PageSize = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfo) SetTotalCount(v string) *ListNodesResponseBodyPagingInfo {
	s.TotalCount = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfo) Validate() error {
	if s.Nodes != nil {
		for _, item := range s.Nodes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListNodesResponseBodyPagingInfoNodes struct {
	// The timestamp when the node in DataStudio was created.
	//
	// example:
	//
	// 1722910655000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The data source.
	DataSource *ListNodesResponseBodyPagingInfoNodesDataSource `json:"DataSource,omitempty" xml:"DataSource,omitempty" type:"Struct"`
	// The description of the node.
	//
	// example:
	//
	// Node description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the node.
	//
	// example:
	//
	// 860438872620113XXXX
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The node input.
	Inputs *ListNodesResponseBodyPagingInfoNodesInputs `json:"Inputs,omitempty" xml:"Inputs,omitempty" type:"Struct"`
	// The timestamp when the node in DataStudio was last modified.
	//
	// example:
	//
	// 1722910655000
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The name of the node.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The node output.
	Outputs *ListNodesResponseBodyPagingInfoNodesOutputs `json:"Outputs,omitempty" xml:"Outputs,omitempty" type:"Struct"`
	// The owner of nodes in DataStudio.
	//
	// example:
	//
	// 110755000425XXXX
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The DataWorks workspace ID. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace page to query the ID.
	//
	// example:
	//
	// 33233
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The scheduling type.
	//
	// Valid values:
	//
	// 	- Normal: The node is scheduled as expected.
	//
	// 	- Pause: The node is paused, and the running of its descendant nodes is blocked.
	//
	// 	- Skip: The node is dry run. The system does not actually run the node but directly prompts that the node is successfully run. The running duration of the node is 0 seconds. In addition, the node does not occupy resources or block the running of its descendant nodes.
	//
	// example:
	//
	// Normal
	Recurrence *string `json:"Recurrence,omitempty" xml:"Recurrence,omitempty"`
	// The information about the resource group.
	RuntimeResource *ListNodesResponseBodyPagingInfoNodesRuntimeResource `json:"RuntimeResource,omitempty" xml:"RuntimeResource,omitempty" type:"Struct"`
	// The script information.
	Script *ListNodesResponseBodyPagingInfoNodesScript `json:"Script,omitempty" xml:"Script,omitempty" type:"Struct"`
	// The scheduling policy.
	Strategy *ListNodesResponseBodyPagingInfoNodesStrategy `json:"Strategy,omitempty" xml:"Strategy,omitempty" type:"Struct"`
	// The tag information (not in use).
	Tags []*ListNodesResponseBodyPagingInfoNodesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The scheduling task ID.
	//
	// example:
	//
	// 88888888888
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The trigger.
	Trigger *ListNodesResponseBodyPagingInfoNodesTrigger `json:"Trigger,omitempty" xml:"Trigger,omitempty" type:"Struct"`
}

func (s ListNodesResponseBodyPagingInfoNodes) String() string {
	return dara.Prettify(s)
}

func (s ListNodesResponseBodyPagingInfoNodes) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyPagingInfoNodes) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListNodesResponseBodyPagingInfoNodes) GetDataSource() *ListNodesResponseBodyPagingInfoNodesDataSource {
	return s.DataSource
}

func (s *ListNodesResponseBodyPagingInfoNodes) GetDescription() *string {
	return s.Description
}

func (s *ListNodesResponseBodyPagingInfoNodes) GetId() *int64 {
	return s.Id
}

func (s *ListNodesResponseBodyPagingInfoNodes) GetInputs() *ListNodesResponseBodyPagingInfoNodesInputs {
	return s.Inputs
}

func (s *ListNodesResponseBodyPagingInfoNodes) GetModifyTime() *int64 {
	return s.ModifyTime
}

func (s *ListNodesResponseBodyPagingInfoNodes) GetName() *string {
	return s.Name
}

func (s *ListNodesResponseBodyPagingInfoNodes) GetOutputs() *ListNodesResponseBodyPagingInfoNodesOutputs {
	return s.Outputs
}

func (s *ListNodesResponseBodyPagingInfoNodes) GetOwner() *string {
	return s.Owner
}

func (s *ListNodesResponseBodyPagingInfoNodes) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListNodesResponseBodyPagingInfoNodes) GetRecurrence() *string {
	return s.Recurrence
}

func (s *ListNodesResponseBodyPagingInfoNodes) GetRuntimeResource() *ListNodesResponseBodyPagingInfoNodesRuntimeResource {
	return s.RuntimeResource
}

func (s *ListNodesResponseBodyPagingInfoNodes) GetScript() *ListNodesResponseBodyPagingInfoNodesScript {
	return s.Script
}

func (s *ListNodesResponseBodyPagingInfoNodes) GetStrategy() *ListNodesResponseBodyPagingInfoNodesStrategy {
	return s.Strategy
}

func (s *ListNodesResponseBodyPagingInfoNodes) GetTags() []*ListNodesResponseBodyPagingInfoNodesTags {
	return s.Tags
}

func (s *ListNodesResponseBodyPagingInfoNodes) GetTaskId() *int64 {
	return s.TaskId
}

func (s *ListNodesResponseBodyPagingInfoNodes) GetTrigger() *ListNodesResponseBodyPagingInfoNodesTrigger {
	return s.Trigger
}

func (s *ListNodesResponseBodyPagingInfoNodes) SetCreateTime(v int64) *ListNodesResponseBodyPagingInfoNodes {
	s.CreateTime = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodes) SetDataSource(v *ListNodesResponseBodyPagingInfoNodesDataSource) *ListNodesResponseBodyPagingInfoNodes {
	s.DataSource = v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodes) SetDescription(v string) *ListNodesResponseBodyPagingInfoNodes {
	s.Description = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodes) SetId(v int64) *ListNodesResponseBodyPagingInfoNodes {
	s.Id = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodes) SetInputs(v *ListNodesResponseBodyPagingInfoNodesInputs) *ListNodesResponseBodyPagingInfoNodes {
	s.Inputs = v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodes) SetModifyTime(v int64) *ListNodesResponseBodyPagingInfoNodes {
	s.ModifyTime = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodes) SetName(v string) *ListNodesResponseBodyPagingInfoNodes {
	s.Name = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodes) SetOutputs(v *ListNodesResponseBodyPagingInfoNodesOutputs) *ListNodesResponseBodyPagingInfoNodes {
	s.Outputs = v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodes) SetOwner(v string) *ListNodesResponseBodyPagingInfoNodes {
	s.Owner = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodes) SetProjectId(v int64) *ListNodesResponseBodyPagingInfoNodes {
	s.ProjectId = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodes) SetRecurrence(v string) *ListNodesResponseBodyPagingInfoNodes {
	s.Recurrence = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodes) SetRuntimeResource(v *ListNodesResponseBodyPagingInfoNodesRuntimeResource) *ListNodesResponseBodyPagingInfoNodes {
	s.RuntimeResource = v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodes) SetScript(v *ListNodesResponseBodyPagingInfoNodesScript) *ListNodesResponseBodyPagingInfoNodes {
	s.Script = v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodes) SetStrategy(v *ListNodesResponseBodyPagingInfoNodesStrategy) *ListNodesResponseBodyPagingInfoNodes {
	s.Strategy = v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodes) SetTags(v []*ListNodesResponseBodyPagingInfoNodesTags) *ListNodesResponseBodyPagingInfoNodes {
	s.Tags = v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodes) SetTaskId(v int64) *ListNodesResponseBodyPagingInfoNodes {
	s.TaskId = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodes) SetTrigger(v *ListNodesResponseBodyPagingInfoNodesTrigger) *ListNodesResponseBodyPagingInfoNodes {
	s.Trigger = v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodes) Validate() error {
	if s.DataSource != nil {
		if err := s.DataSource.Validate(); err != nil {
			return err
		}
	}
	if s.Inputs != nil {
		if err := s.Inputs.Validate(); err != nil {
			return err
		}
	}
	if s.Outputs != nil {
		if err := s.Outputs.Validate(); err != nil {
			return err
		}
	}
	if s.RuntimeResource != nil {
		if err := s.RuntimeResource.Validate(); err != nil {
			return err
		}
	}
	if s.Script != nil {
		if err := s.Script.Validate(); err != nil {
			return err
		}
	}
	if s.Strategy != nil {
		if err := s.Strategy.Validate(); err != nil {
			return err
		}
	}
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Trigger != nil {
		if err := s.Trigger.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListNodesResponseBodyPagingInfoNodesDataSource struct {
	// The name of the data source.
	//
	// example:
	//
	// odps_first
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type of the data source.
	//
	// example:
	//
	// odps
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListNodesResponseBodyPagingInfoNodesDataSource) String() string {
	return dara.Prettify(s)
}

func (s ListNodesResponseBodyPagingInfoNodesDataSource) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyPagingInfoNodesDataSource) GetName() *string {
	return s.Name
}

func (s *ListNodesResponseBodyPagingInfoNodesDataSource) GetType() *string {
	return s.Type
}

func (s *ListNodesResponseBodyPagingInfoNodesDataSource) SetName(v string) *ListNodesResponseBodyPagingInfoNodesDataSource {
	s.Name = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesDataSource) SetType(v string) *ListNodesResponseBodyPagingInfoNodesDataSource {
	s.Type = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesDataSource) Validate() error {
	return dara.Validate(s)
}

type ListNodesResponseBodyPagingInfoNodesInputs struct {
	// The node output list.
	NodeOutputs []*ListNodesResponseBodyPagingInfoNodesInputsNodeOutputs `json:"NodeOutputs,omitempty" xml:"NodeOutputs,omitempty" type:"Repeated"`
	// The table list.
	Tables []*ListNodesResponseBodyPagingInfoNodesInputsTables `json:"Tables,omitempty" xml:"Tables,omitempty" type:"Repeated"`
	// The variable list.
	Variables []*ListNodesResponseBodyPagingInfoNodesInputsVariables `json:"Variables,omitempty" xml:"Variables,omitempty" type:"Repeated"`
}

func (s ListNodesResponseBodyPagingInfoNodesInputs) String() string {
	return dara.Prettify(s)
}

func (s ListNodesResponseBodyPagingInfoNodesInputs) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyPagingInfoNodesInputs) GetNodeOutputs() []*ListNodesResponseBodyPagingInfoNodesInputsNodeOutputs {
	return s.NodeOutputs
}

func (s *ListNodesResponseBodyPagingInfoNodesInputs) GetTables() []*ListNodesResponseBodyPagingInfoNodesInputsTables {
	return s.Tables
}

func (s *ListNodesResponseBodyPagingInfoNodesInputs) GetVariables() []*ListNodesResponseBodyPagingInfoNodesInputsVariables {
	return s.Variables
}

func (s *ListNodesResponseBodyPagingInfoNodesInputs) SetNodeOutputs(v []*ListNodesResponseBodyPagingInfoNodesInputsNodeOutputs) *ListNodesResponseBodyPagingInfoNodesInputs {
	s.NodeOutputs = v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesInputs) SetTables(v []*ListNodesResponseBodyPagingInfoNodesInputsTables) *ListNodesResponseBodyPagingInfoNodesInputs {
	s.Tables = v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesInputs) SetVariables(v []*ListNodesResponseBodyPagingInfoNodesInputsVariables) *ListNodesResponseBodyPagingInfoNodesInputs {
	s.Variables = v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesInputs) Validate() error {
	if s.NodeOutputs != nil {
		for _, item := range s.NodeOutputs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Tables != nil {
		for _, item := range s.Tables {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Variables != nil {
		for _, item := range s.Variables {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListNodesResponseBodyPagingInfoNodesInputsNodeOutputs struct {
	// The node output.
	//
	// example:
	//
	// 623731286945488XXXX
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s ListNodesResponseBodyPagingInfoNodesInputsNodeOutputs) String() string {
	return dara.Prettify(s)
}

func (s ListNodesResponseBodyPagingInfoNodesInputsNodeOutputs) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyPagingInfoNodesInputsNodeOutputs) GetData() *string {
	return s.Data
}

func (s *ListNodesResponseBodyPagingInfoNodesInputsNodeOutputs) SetData(v string) *ListNodesResponseBodyPagingInfoNodesInputsNodeOutputs {
	s.Data = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesInputsNodeOutputs) Validate() error {
	return dara.Validate(s)
}

type ListNodesResponseBodyPagingInfoNodesInputsTables struct {
	// The table ID.
	//
	// example:
	//
	// odps.autotest.test_output_table_1
	Guid *string `json:"Guid,omitempty" xml:"Guid,omitempty"`
}

func (s ListNodesResponseBodyPagingInfoNodesInputsTables) String() string {
	return dara.Prettify(s)
}

func (s ListNodesResponseBodyPagingInfoNodesInputsTables) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyPagingInfoNodesInputsTables) GetGuid() *string {
	return s.Guid
}

func (s *ListNodesResponseBodyPagingInfoNodesInputsTables) SetGuid(v string) *ListNodesResponseBodyPagingInfoNodesInputsTables {
	s.Guid = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesInputsTables) Validate() error {
	return dara.Validate(s)
}

type ListNodesResponseBodyPagingInfoNodesInputsVariables struct {
	// The artifact type.
	//
	// example:
	//
	// Variable
	ArtifactType *string `json:"ArtifactType,omitempty" xml:"ArtifactType,omitempty"`
	// The variable ID.
	//
	// example:
	//
	// 543211286945488XXXX
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the variable.
	//
	// example:
	//
	// input
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The node to which the variable belongs.
	Node *ListNodesResponseBodyPagingInfoNodesInputsVariablesNode `json:"Node,omitempty" xml:"Node,omitempty" type:"Struct"`
	// The scope of the variable. Valid values:
	//
	// 	- WorkSpace
	//
	// 	- NodeParameter
	//
	// 	- NodeContext
	//
	// 	- Workflow
	//
	// example:
	//
	// NodeParameter
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// The type of the variable. Valid values:
	//
	// 	- NoKvVariableExpression
	//
	// 	- Constant
	//
	// 	- PassThrough
	//
	// 	- System
	//
	// 	- NodeOutput
	//
	// example:
	//
	// Constant
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The variable value.
	//
	// example:
	//
	// 222
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListNodesResponseBodyPagingInfoNodesInputsVariables) String() string {
	return dara.Prettify(s)
}

func (s ListNodesResponseBodyPagingInfoNodesInputsVariables) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyPagingInfoNodesInputsVariables) GetArtifactType() *string {
	return s.ArtifactType
}

func (s *ListNodesResponseBodyPagingInfoNodesInputsVariables) GetId() *int64 {
	return s.Id
}

func (s *ListNodesResponseBodyPagingInfoNodesInputsVariables) GetName() *string {
	return s.Name
}

func (s *ListNodesResponseBodyPagingInfoNodesInputsVariables) GetNode() *ListNodesResponseBodyPagingInfoNodesInputsVariablesNode {
	return s.Node
}

func (s *ListNodesResponseBodyPagingInfoNodesInputsVariables) GetScope() *string {
	return s.Scope
}

func (s *ListNodesResponseBodyPagingInfoNodesInputsVariables) GetType() *string {
	return s.Type
}

func (s *ListNodesResponseBodyPagingInfoNodesInputsVariables) GetValue() *string {
	return s.Value
}

func (s *ListNodesResponseBodyPagingInfoNodesInputsVariables) SetArtifactType(v string) *ListNodesResponseBodyPagingInfoNodesInputsVariables {
	s.ArtifactType = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesInputsVariables) SetId(v int64) *ListNodesResponseBodyPagingInfoNodesInputsVariables {
	s.Id = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesInputsVariables) SetName(v string) *ListNodesResponseBodyPagingInfoNodesInputsVariables {
	s.Name = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesInputsVariables) SetNode(v *ListNodesResponseBodyPagingInfoNodesInputsVariablesNode) *ListNodesResponseBodyPagingInfoNodesInputsVariables {
	s.Node = v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesInputsVariables) SetScope(v string) *ListNodesResponseBodyPagingInfoNodesInputsVariables {
	s.Scope = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesInputsVariables) SetType(v string) *ListNodesResponseBodyPagingInfoNodesInputsVariables {
	s.Type = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesInputsVariables) SetValue(v string) *ListNodesResponseBodyPagingInfoNodesInputsVariables {
	s.Value = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesInputsVariables) Validate() error {
	if s.Node != nil {
		if err := s.Node.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListNodesResponseBodyPagingInfoNodesInputsVariablesNode struct {
	// The node output.
	//
	// example:
	//
	// 623731286945488XXXX
	Output *string `json:"Output,omitempty" xml:"Output,omitempty"`
}

func (s ListNodesResponseBodyPagingInfoNodesInputsVariablesNode) String() string {
	return dara.Prettify(s)
}

func (s ListNodesResponseBodyPagingInfoNodesInputsVariablesNode) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyPagingInfoNodesInputsVariablesNode) GetOutput() *string {
	return s.Output
}

func (s *ListNodesResponseBodyPagingInfoNodesInputsVariablesNode) SetOutput(v string) *ListNodesResponseBodyPagingInfoNodesInputsVariablesNode {
	s.Output = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesInputsVariablesNode) Validate() error {
	return dara.Validate(s)
}

type ListNodesResponseBodyPagingInfoNodesOutputs struct {
	// The node output list.
	NodeOutputs []*ListNodesResponseBodyPagingInfoNodesOutputsNodeOutputs `json:"NodeOutputs,omitempty" xml:"NodeOutputs,omitempty" type:"Repeated"`
	// The table list.
	Tables []*ListNodesResponseBodyPagingInfoNodesOutputsTables `json:"Tables,omitempty" xml:"Tables,omitempty" type:"Repeated"`
	// The variable list.
	Variables []*ListNodesResponseBodyPagingInfoNodesOutputsVariables `json:"Variables,omitempty" xml:"Variables,omitempty" type:"Repeated"`
}

func (s ListNodesResponseBodyPagingInfoNodesOutputs) String() string {
	return dara.Prettify(s)
}

func (s ListNodesResponseBodyPagingInfoNodesOutputs) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyPagingInfoNodesOutputs) GetNodeOutputs() []*ListNodesResponseBodyPagingInfoNodesOutputsNodeOutputs {
	return s.NodeOutputs
}

func (s *ListNodesResponseBodyPagingInfoNodesOutputs) GetTables() []*ListNodesResponseBodyPagingInfoNodesOutputsTables {
	return s.Tables
}

func (s *ListNodesResponseBodyPagingInfoNodesOutputs) GetVariables() []*ListNodesResponseBodyPagingInfoNodesOutputsVariables {
	return s.Variables
}

func (s *ListNodesResponseBodyPagingInfoNodesOutputs) SetNodeOutputs(v []*ListNodesResponseBodyPagingInfoNodesOutputsNodeOutputs) *ListNodesResponseBodyPagingInfoNodesOutputs {
	s.NodeOutputs = v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesOutputs) SetTables(v []*ListNodesResponseBodyPagingInfoNodesOutputsTables) *ListNodesResponseBodyPagingInfoNodesOutputs {
	s.Tables = v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesOutputs) SetVariables(v []*ListNodesResponseBodyPagingInfoNodesOutputsVariables) *ListNodesResponseBodyPagingInfoNodesOutputs {
	s.Variables = v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesOutputs) Validate() error {
	if s.NodeOutputs != nil {
		for _, item := range s.NodeOutputs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Tables != nil {
		for _, item := range s.Tables {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Variables != nil {
		for _, item := range s.Variables {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListNodesResponseBodyPagingInfoNodesOutputsNodeOutputs struct {
	// The node output.
	//
	// example:
	//
	// 860438872620113XXXX
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s ListNodesResponseBodyPagingInfoNodesOutputsNodeOutputs) String() string {
	return dara.Prettify(s)
}

func (s ListNodesResponseBodyPagingInfoNodesOutputsNodeOutputs) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyPagingInfoNodesOutputsNodeOutputs) GetData() *string {
	return s.Data
}

func (s *ListNodesResponseBodyPagingInfoNodesOutputsNodeOutputs) SetData(v string) *ListNodesResponseBodyPagingInfoNodesOutputsNodeOutputs {
	s.Data = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesOutputsNodeOutputs) Validate() error {
	return dara.Validate(s)
}

type ListNodesResponseBodyPagingInfoNodesOutputsTables struct {
	// The table ID.
	//
	// example:
	//
	// odps.autotest.test_output_table_1
	Guid *string `json:"Guid,omitempty" xml:"Guid,omitempty"`
}

func (s ListNodesResponseBodyPagingInfoNodesOutputsTables) String() string {
	return dara.Prettify(s)
}

func (s ListNodesResponseBodyPagingInfoNodesOutputsTables) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyPagingInfoNodesOutputsTables) GetGuid() *string {
	return s.Guid
}

func (s *ListNodesResponseBodyPagingInfoNodesOutputsTables) SetGuid(v string) *ListNodesResponseBodyPagingInfoNodesOutputsTables {
	s.Guid = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesOutputsTables) Validate() error {
	return dara.Validate(s)
}

type ListNodesResponseBodyPagingInfoNodesOutputsVariables struct {
	// The artifact type.
	//
	// example:
	//
	// Variable
	ArtifactType *string `json:"ArtifactType,omitempty" xml:"ArtifactType,omitempty"`
	// The variable ID.
	//
	// example:
	//
	// 623731286945488XXXX
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the variable.
	//
	// example:
	//
	// output
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The node to which the variable belongs.
	Node *ListNodesResponseBodyPagingInfoNodesOutputsVariablesNode `json:"Node,omitempty" xml:"Node,omitempty" type:"Struct"`
	// The variable scope. Valid values:
	//
	// 	- NodeParameter
	//
	// 	- NodeContext
	//
	// 	- Workflow
	//
	// 	- Workspace
	//
	// example:
	//
	// NodeParameter
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// The variable type. Valid values:
	//
	// 	- NoKvVariableExpression
	//
	// 	- Constant
	//
	// 	- PassThrough
	//
	// 	- System
	//
	// 	- NodeOutput
	//
	// example:
	//
	// Constant
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The variable value.
	//
	// example:
	//
	// 111
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListNodesResponseBodyPagingInfoNodesOutputsVariables) String() string {
	return dara.Prettify(s)
}

func (s ListNodesResponseBodyPagingInfoNodesOutputsVariables) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyPagingInfoNodesOutputsVariables) GetArtifactType() *string {
	return s.ArtifactType
}

func (s *ListNodesResponseBodyPagingInfoNodesOutputsVariables) GetId() *int64 {
	return s.Id
}

func (s *ListNodesResponseBodyPagingInfoNodesOutputsVariables) GetName() *string {
	return s.Name
}

func (s *ListNodesResponseBodyPagingInfoNodesOutputsVariables) GetNode() *ListNodesResponseBodyPagingInfoNodesOutputsVariablesNode {
	return s.Node
}

func (s *ListNodesResponseBodyPagingInfoNodesOutputsVariables) GetScope() *string {
	return s.Scope
}

func (s *ListNodesResponseBodyPagingInfoNodesOutputsVariables) GetType() *string {
	return s.Type
}

func (s *ListNodesResponseBodyPagingInfoNodesOutputsVariables) GetValue() *string {
	return s.Value
}

func (s *ListNodesResponseBodyPagingInfoNodesOutputsVariables) SetArtifactType(v string) *ListNodesResponseBodyPagingInfoNodesOutputsVariables {
	s.ArtifactType = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesOutputsVariables) SetId(v int64) *ListNodesResponseBodyPagingInfoNodesOutputsVariables {
	s.Id = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesOutputsVariables) SetName(v string) *ListNodesResponseBodyPagingInfoNodesOutputsVariables {
	s.Name = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesOutputsVariables) SetNode(v *ListNodesResponseBodyPagingInfoNodesOutputsVariablesNode) *ListNodesResponseBodyPagingInfoNodesOutputsVariables {
	s.Node = v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesOutputsVariables) SetScope(v string) *ListNodesResponseBodyPagingInfoNodesOutputsVariables {
	s.Scope = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesOutputsVariables) SetType(v string) *ListNodesResponseBodyPagingInfoNodesOutputsVariables {
	s.Type = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesOutputsVariables) SetValue(v string) *ListNodesResponseBodyPagingInfoNodesOutputsVariables {
	s.Value = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesOutputsVariables) Validate() error {
	if s.Node != nil {
		if err := s.Node.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListNodesResponseBodyPagingInfoNodesOutputsVariablesNode struct {
	// The node output.
	//
	// example:
	//
	// 860438872620113XXXX
	Output *string `json:"Output,omitempty" xml:"Output,omitempty"`
}

func (s ListNodesResponseBodyPagingInfoNodesOutputsVariablesNode) String() string {
	return dara.Prettify(s)
}

func (s ListNodesResponseBodyPagingInfoNodesOutputsVariablesNode) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyPagingInfoNodesOutputsVariablesNode) GetOutput() *string {
	return s.Output
}

func (s *ListNodesResponseBodyPagingInfoNodesOutputsVariablesNode) SetOutput(v string) *ListNodesResponseBodyPagingInfoNodesOutputsVariablesNode {
	s.Output = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesOutputsVariablesNode) Validate() error {
	return dara.Validate(s)
}

type ListNodesResponseBodyPagingInfoNodesRuntimeResource struct {
	// The identifier of the resource group. You can call the [ListResourceGroups](https://help.aliyun.com/document_detail/173913.html) operation to query the identifier of the resource group.
	//
	// example:
	//
	// S_res_group_XXXX
	ResourceGroup *string `json:"ResourceGroup,omitempty" xml:"ResourceGroup,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// S_resgrop_xxx
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ListNodesResponseBodyPagingInfoNodesRuntimeResource) String() string {
	return dara.Prettify(s)
}

func (s ListNodesResponseBodyPagingInfoNodesRuntimeResource) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyPagingInfoNodesRuntimeResource) GetResourceGroup() *string {
	return s.ResourceGroup
}

func (s *ListNodesResponseBodyPagingInfoNodesRuntimeResource) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListNodesResponseBodyPagingInfoNodesRuntimeResource) SetResourceGroup(v string) *ListNodesResponseBodyPagingInfoNodesRuntimeResource {
	s.ResourceGroup = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesRuntimeResource) SetResourceGroupId(v string) *ListNodesResponseBodyPagingInfoNodesRuntimeResource {
	s.ResourceGroupId = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesRuntimeResource) Validate() error {
	return dara.Validate(s)
}

type ListNodesResponseBodyPagingInfoNodesScript struct {
	// The script ID.
	//
	// example:
	//
	// 853573334108680XXXX
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The script path.
	//
	// example:
	//
	// root/demo
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The runtime.
	Runtime *ListNodesResponseBodyPagingInfoNodesScriptRuntime `json:"Runtime,omitempty" xml:"Runtime,omitempty" type:"Struct"`
}

func (s ListNodesResponseBodyPagingInfoNodesScript) String() string {
	return dara.Prettify(s)
}

func (s ListNodesResponseBodyPagingInfoNodesScript) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyPagingInfoNodesScript) GetId() *int64 {
	return s.Id
}

func (s *ListNodesResponseBodyPagingInfoNodesScript) GetPath() *string {
	return s.Path
}

func (s *ListNodesResponseBodyPagingInfoNodesScript) GetRuntime() *ListNodesResponseBodyPagingInfoNodesScriptRuntime {
	return s.Runtime
}

func (s *ListNodesResponseBodyPagingInfoNodesScript) SetId(v int64) *ListNodesResponseBodyPagingInfoNodesScript {
	s.Id = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesScript) SetPath(v string) *ListNodesResponseBodyPagingInfoNodesScript {
	s.Path = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesScript) SetRuntime(v *ListNodesResponseBodyPagingInfoNodesScriptRuntime) *ListNodesResponseBodyPagingInfoNodesScript {
	s.Runtime = v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesScript) Validate() error {
	if s.Runtime != nil {
		if err := s.Runtime.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListNodesResponseBodyPagingInfoNodesScriptRuntime struct {
	// The command used to distinguish node types.
	//
	// example:
	//
	// ODPS_SQL
	Command *string `json:"Command,omitempty" xml:"Command,omitempty"`
}

func (s ListNodesResponseBodyPagingInfoNodesScriptRuntime) String() string {
	return dara.Prettify(s)
}

func (s ListNodesResponseBodyPagingInfoNodesScriptRuntime) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyPagingInfoNodesScriptRuntime) GetCommand() *string {
	return s.Command
}

func (s *ListNodesResponseBodyPagingInfoNodesScriptRuntime) SetCommand(v string) *ListNodesResponseBodyPagingInfoNodesScriptRuntime {
	s.Command = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesScriptRuntime) Validate() error {
	return dara.Validate(s)
}

type ListNodesResponseBodyPagingInfoNodesStrategy struct {
	// The instance generation mode. Valid values:
	//
	// 	- T+1
	//
	// 	- Immediately
	//
	// example:
	//
	// T+1
	InstanceMode *string `json:"InstanceMode,omitempty" xml:"InstanceMode,omitempty"`
	// The rerun interval. Unit: milliseconds.
	//
	// example:
	//
	// 180000
	RerunInterval *int32 `json:"RerunInterval,omitempty" xml:"RerunInterval,omitempty"`
	// The rerun mode. Valid values:
	//
	// 	- Allowed
	//
	// 	- Denied
	//
	// 	- FailureAllowed
	//
	// example:
	//
	// Allowed
	RerunMode *string `json:"RerunMode,omitempty" xml:"RerunMode,omitempty"`
	// The number of reruns.
	//
	// example:
	//
	// 3
	RerunTimes *int32 `json:"RerunTimes,omitempty" xml:"RerunTimes,omitempty"`
	// Timeout.
	//
	// example:
	//
	// 0
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s ListNodesResponseBodyPagingInfoNodesStrategy) String() string {
	return dara.Prettify(s)
}

func (s ListNodesResponseBodyPagingInfoNodesStrategy) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyPagingInfoNodesStrategy) GetInstanceMode() *string {
	return s.InstanceMode
}

func (s *ListNodesResponseBodyPagingInfoNodesStrategy) GetRerunInterval() *int32 {
	return s.RerunInterval
}

func (s *ListNodesResponseBodyPagingInfoNodesStrategy) GetRerunMode() *string {
	return s.RerunMode
}

func (s *ListNodesResponseBodyPagingInfoNodesStrategy) GetRerunTimes() *int32 {
	return s.RerunTimes
}

func (s *ListNodesResponseBodyPagingInfoNodesStrategy) GetTimeout() *int32 {
	return s.Timeout
}

func (s *ListNodesResponseBodyPagingInfoNodesStrategy) SetInstanceMode(v string) *ListNodesResponseBodyPagingInfoNodesStrategy {
	s.InstanceMode = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesStrategy) SetRerunInterval(v int32) *ListNodesResponseBodyPagingInfoNodesStrategy {
	s.RerunInterval = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesStrategy) SetRerunMode(v string) *ListNodesResponseBodyPagingInfoNodesStrategy {
	s.RerunMode = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesStrategy) SetRerunTimes(v int32) *ListNodesResponseBodyPagingInfoNodesStrategy {
	s.RerunTimes = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesStrategy) SetTimeout(v int32) *ListNodesResponseBodyPagingInfoNodesStrategy {
	s.Timeout = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesStrategy) Validate() error {
	return dara.Validate(s)
}

type ListNodesResponseBodyPagingInfoNodesTags struct {
	// The tag key.
	//
	// example:
	//
	// null
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// null
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListNodesResponseBodyPagingInfoNodesTags) String() string {
	return dara.Prettify(s)
}

func (s ListNodesResponseBodyPagingInfoNodesTags) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyPagingInfoNodesTags) GetKey() *string {
	return s.Key
}

func (s *ListNodesResponseBodyPagingInfoNodesTags) GetValue() *string {
	return s.Value
}

func (s *ListNodesResponseBodyPagingInfoNodesTags) SetKey(v string) *ListNodesResponseBodyPagingInfoNodesTags {
	s.Key = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesTags) SetValue(v string) *ListNodesResponseBodyPagingInfoNodesTags {
	s.Value = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesTags) Validate() error {
	return dara.Validate(s)
}

type ListNodesResponseBodyPagingInfoNodesTrigger struct {
	// The cron expression for scheduling
	//
	// example:
	//
	// 00 00 00 	- 	- ?
	Cron *string `json:"Cron,omitempty" xml:"Cron,omitempty"`
	// The end time of the validity period of the trigger.
	//
	// example:
	//
	// 9999-01-01 00:00:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The trigger ID.
	//
	// example:
	//
	// 543680677872062XXXX
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The start time of the validity period of the trigger.
	//
	// example:
	//
	// 1970-01-01 00:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The time zone.
	//
	// example:
	//
	// Asia/Shanghai
	Timezone *string `json:"Timezone,omitempty" xml:"Timezone,omitempty"`
	// The trigger type.
	//
	// Valid values:
	//
	// 	- Scheduler: periodic scheduling
	//
	// 	- Manual: manual trigger
	//
	// 	- Streaming: streaming task
	//
	// example:
	//
	// Scheduler
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListNodesResponseBodyPagingInfoNodesTrigger) String() string {
	return dara.Prettify(s)
}

func (s ListNodesResponseBodyPagingInfoNodesTrigger) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyPagingInfoNodesTrigger) GetCron() *string {
	return s.Cron
}

func (s *ListNodesResponseBodyPagingInfoNodesTrigger) GetEndTime() *string {
	return s.EndTime
}

func (s *ListNodesResponseBodyPagingInfoNodesTrigger) GetId() *int64 {
	return s.Id
}

func (s *ListNodesResponseBodyPagingInfoNodesTrigger) GetStartTime() *string {
	return s.StartTime
}

func (s *ListNodesResponseBodyPagingInfoNodesTrigger) GetTimezone() *string {
	return s.Timezone
}

func (s *ListNodesResponseBodyPagingInfoNodesTrigger) GetType() *string {
	return s.Type
}

func (s *ListNodesResponseBodyPagingInfoNodesTrigger) SetCron(v string) *ListNodesResponseBodyPagingInfoNodesTrigger {
	s.Cron = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesTrigger) SetEndTime(v string) *ListNodesResponseBodyPagingInfoNodesTrigger {
	s.EndTime = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesTrigger) SetId(v int64) *ListNodesResponseBodyPagingInfoNodesTrigger {
	s.Id = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesTrigger) SetStartTime(v string) *ListNodesResponseBodyPagingInfoNodesTrigger {
	s.StartTime = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesTrigger) SetTimezone(v string) *ListNodesResponseBodyPagingInfoNodesTrigger {
	s.Timezone = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesTrigger) SetType(v string) *ListNodesResponseBodyPagingInfoNodesTrigger {
	s.Type = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesTrigger) Validate() error {
	return dara.Validate(s)
}
