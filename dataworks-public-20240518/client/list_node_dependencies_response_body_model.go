// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNodeDependenciesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPagingInfo(v *ListNodeDependenciesResponseBodyPagingInfo) *ListNodeDependenciesResponseBody
	GetPagingInfo() *ListNodeDependenciesResponseBodyPagingInfo
	SetRequestId(v string) *ListNodeDependenciesResponseBody
	GetRequestId() *string
}

type ListNodeDependenciesResponseBody struct {
	// The pagination information.
	PagingInfo *ListNodeDependenciesResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 204EAF68-CCE3-5112-8DA0-E7A60F02XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListNodeDependenciesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListNodeDependenciesResponseBody) GoString() string {
	return s.String()
}

func (s *ListNodeDependenciesResponseBody) GetPagingInfo() *ListNodeDependenciesResponseBodyPagingInfo {
	return s.PagingInfo
}

func (s *ListNodeDependenciesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListNodeDependenciesResponseBody) SetPagingInfo(v *ListNodeDependenciesResponseBodyPagingInfo) *ListNodeDependenciesResponseBody {
	s.PagingInfo = v
	return s
}

func (s *ListNodeDependenciesResponseBody) SetRequestId(v string) *ListNodeDependenciesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNodeDependenciesResponseBody) Validate() error {
	if s.PagingInfo != nil {
		if err := s.PagingInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListNodeDependenciesResponseBodyPagingInfo struct {
	// The list of dependent nodes.
	Nodes []*ListNodeDependenciesResponseBodyPagingInfoNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
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
	// 90
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListNodeDependenciesResponseBodyPagingInfo) String() string {
	return dara.Prettify(s)
}

func (s ListNodeDependenciesResponseBodyPagingInfo) GoString() string {
	return s.String()
}

func (s *ListNodeDependenciesResponseBodyPagingInfo) GetNodes() []*ListNodeDependenciesResponseBodyPagingInfoNodes {
	return s.Nodes
}

func (s *ListNodeDependenciesResponseBodyPagingInfo) GetPageNumber() *string {
	return s.PageNumber
}

func (s *ListNodeDependenciesResponseBodyPagingInfo) GetPageSize() *string {
	return s.PageSize
}

func (s *ListNodeDependenciesResponseBodyPagingInfo) GetTotalCount() *string {
	return s.TotalCount
}

func (s *ListNodeDependenciesResponseBodyPagingInfo) SetNodes(v []*ListNodeDependenciesResponseBodyPagingInfoNodes) *ListNodeDependenciesResponseBodyPagingInfo {
	s.Nodes = v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfo) SetPageNumber(v string) *ListNodeDependenciesResponseBodyPagingInfo {
	s.PageNumber = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfo) SetPageSize(v string) *ListNodeDependenciesResponseBodyPagingInfo {
	s.PageSize = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfo) SetTotalCount(v string) *ListNodeDependenciesResponseBodyPagingInfo {
	s.TotalCount = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfo) Validate() error {
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

type ListNodeDependenciesResponseBodyPagingInfoNodes struct {
	// The timestamp when the node was created.
	//
	// example:
	//
	// 1724505917000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The data source.
	DataSource *ListNodeDependenciesResponseBodyPagingInfoNodesDataSource `json:"DataSource,omitempty" xml:"DataSource,omitempty" type:"Struct"`
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
	// 723932906364267XXXX
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The input of the node.
	Inputs *ListNodeDependenciesResponseBodyPagingInfoNodesInputs `json:"Inputs,omitempty" xml:"Inputs,omitempty" type:"Struct"`
	// The timestamp when the node was last modified.
	//
	// example:
	//
	// 1724505917000
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The name of the node.
	//
	// example:
	//
	// Node name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The output of the node.
	Outputs *ListNodeDependenciesResponseBodyPagingInfoNodesOutputs `json:"Outputs,omitempty" xml:"Outputs,omitempty" type:"Struct"`
	// The owner of the node.
	//
	// example:
	//
	// 110755000425XXXX
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The ID of the workspace to which the node belongs.
	//
	// example:
	//
	// 65133
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The scheduling type.
	//
	// Valid values:
	//
	// 	- Normal: Nodes are scheduled as expected.
	//
	// 	- Pause: Nodes are paused, and the running of their descendant nodes is blocked.
	//
	// 	- Skip: Nodes are dry run. The system does not actually run the nodes but directly prompts that the nodes are successfully run. The running duration of the nodes is 0 seconds. In addition, the nodes do not occupy resources or block the running of their descendant nodes.
	//
	// example:
	//
	// Normal
	Recurrence *string `json:"Recurrence,omitempty" xml:"Recurrence,omitempty"`
	// The information about the resource group.
	RuntimeResource *ListNodeDependenciesResponseBodyPagingInfoNodesRuntimeResource `json:"RuntimeResource,omitempty" xml:"RuntimeResource,omitempty" type:"Struct"`
	// The script information.
	Script *ListNodeDependenciesResponseBodyPagingInfoNodesScript `json:"Script,omitempty" xml:"Script,omitempty" type:"Struct"`
	// The scheduling policy.
	Strategy *ListNodeDependenciesResponseBodyPagingInfoNodesStrategy `json:"Strategy,omitempty" xml:"Strategy,omitempty" type:"Struct"`
	// The tags. This parameter is not in use.
	Tags []*ListNodeDependenciesResponseBodyPagingInfoNodesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The scheduling task ID.
	//
	// example:
	//
	// 580667964888595XXXX
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The trigger.
	Trigger *ListNodeDependenciesResponseBodyPagingInfoNodesTrigger `json:"Trigger,omitempty" xml:"Trigger,omitempty" type:"Struct"`
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodes) String() string {
	return dara.Prettify(s)
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodes) GoString() string {
	return s.String()
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodes) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodes) GetDataSource() *ListNodeDependenciesResponseBodyPagingInfoNodesDataSource {
	return s.DataSource
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodes) GetDescription() *string {
	return s.Description
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodes) GetId() *int64 {
	return s.Id
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodes) GetInputs() *ListNodeDependenciesResponseBodyPagingInfoNodesInputs {
	return s.Inputs
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodes) GetModifyTime() *int64 {
	return s.ModifyTime
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodes) GetName() *string {
	return s.Name
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodes) GetOutputs() *ListNodeDependenciesResponseBodyPagingInfoNodesOutputs {
	return s.Outputs
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodes) GetOwner() *string {
	return s.Owner
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodes) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodes) GetRecurrence() *string {
	return s.Recurrence
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodes) GetRuntimeResource() *ListNodeDependenciesResponseBodyPagingInfoNodesRuntimeResource {
	return s.RuntimeResource
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodes) GetScript() *ListNodeDependenciesResponseBodyPagingInfoNodesScript {
	return s.Script
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodes) GetStrategy() *ListNodeDependenciesResponseBodyPagingInfoNodesStrategy {
	return s.Strategy
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodes) GetTags() []*ListNodeDependenciesResponseBodyPagingInfoNodesTags {
	return s.Tags
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodes) GetTaskId() *int64 {
	return s.TaskId
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodes) GetTrigger() *ListNodeDependenciesResponseBodyPagingInfoNodesTrigger {
	return s.Trigger
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodes) SetCreateTime(v int64) *ListNodeDependenciesResponseBodyPagingInfoNodes {
	s.CreateTime = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodes) SetDataSource(v *ListNodeDependenciesResponseBodyPagingInfoNodesDataSource) *ListNodeDependenciesResponseBodyPagingInfoNodes {
	s.DataSource = v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodes) SetDescription(v string) *ListNodeDependenciesResponseBodyPagingInfoNodes {
	s.Description = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodes) SetId(v int64) *ListNodeDependenciesResponseBodyPagingInfoNodes {
	s.Id = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodes) SetInputs(v *ListNodeDependenciesResponseBodyPagingInfoNodesInputs) *ListNodeDependenciesResponseBodyPagingInfoNodes {
	s.Inputs = v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodes) SetModifyTime(v int64) *ListNodeDependenciesResponseBodyPagingInfoNodes {
	s.ModifyTime = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodes) SetName(v string) *ListNodeDependenciesResponseBodyPagingInfoNodes {
	s.Name = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodes) SetOutputs(v *ListNodeDependenciesResponseBodyPagingInfoNodesOutputs) *ListNodeDependenciesResponseBodyPagingInfoNodes {
	s.Outputs = v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodes) SetOwner(v string) *ListNodeDependenciesResponseBodyPagingInfoNodes {
	s.Owner = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodes) SetProjectId(v int64) *ListNodeDependenciesResponseBodyPagingInfoNodes {
	s.ProjectId = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodes) SetRecurrence(v string) *ListNodeDependenciesResponseBodyPagingInfoNodes {
	s.Recurrence = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodes) SetRuntimeResource(v *ListNodeDependenciesResponseBodyPagingInfoNodesRuntimeResource) *ListNodeDependenciesResponseBodyPagingInfoNodes {
	s.RuntimeResource = v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodes) SetScript(v *ListNodeDependenciesResponseBodyPagingInfoNodesScript) *ListNodeDependenciesResponseBodyPagingInfoNodes {
	s.Script = v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodes) SetStrategy(v *ListNodeDependenciesResponseBodyPagingInfoNodesStrategy) *ListNodeDependenciesResponseBodyPagingInfoNodes {
	s.Strategy = v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodes) SetTags(v []*ListNodeDependenciesResponseBodyPagingInfoNodesTags) *ListNodeDependenciesResponseBodyPagingInfoNodes {
	s.Tags = v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodes) SetTaskId(v int64) *ListNodeDependenciesResponseBodyPagingInfoNodes {
	s.TaskId = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodes) SetTrigger(v *ListNodeDependenciesResponseBodyPagingInfoNodesTrigger) *ListNodeDependenciesResponseBodyPagingInfoNodes {
	s.Trigger = v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodes) Validate() error {
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

type ListNodeDependenciesResponseBodyPagingInfoNodesDataSource struct {
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

func (s ListNodeDependenciesResponseBodyPagingInfoNodesDataSource) String() string {
	return dara.Prettify(s)
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesDataSource) GoString() string {
	return s.String()
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesDataSource) GetName() *string {
	return s.Name
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesDataSource) GetType() *string {
	return s.Type
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesDataSource) SetName(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesDataSource {
	s.Name = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesDataSource) SetType(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesDataSource {
	s.Type = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesDataSource) Validate() error {
	return dara.Validate(s)
}

type ListNodeDependenciesResponseBodyPagingInfoNodesInputs struct {
	// The list of node outputs.
	NodeOutputs []*ListNodeDependenciesResponseBodyPagingInfoNodesInputsNodeOutputs `json:"NodeOutputs,omitempty" xml:"NodeOutputs,omitempty" type:"Repeated"`
	// The table list.
	Tables []*ListNodeDependenciesResponseBodyPagingInfoNodesInputsTables `json:"Tables,omitempty" xml:"Tables,omitempty" type:"Repeated"`
	// The variable list.
	Variables []*ListNodeDependenciesResponseBodyPagingInfoNodesInputsVariables `json:"Variables,omitempty" xml:"Variables,omitempty" type:"Repeated"`
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesInputs) String() string {
	return dara.Prettify(s)
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesInputs) GoString() string {
	return s.String()
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesInputs) GetNodeOutputs() []*ListNodeDependenciesResponseBodyPagingInfoNodesInputsNodeOutputs {
	return s.NodeOutputs
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesInputs) GetTables() []*ListNodeDependenciesResponseBodyPagingInfoNodesInputsTables {
	return s.Tables
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesInputs) GetVariables() []*ListNodeDependenciesResponseBodyPagingInfoNodesInputsVariables {
	return s.Variables
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesInputs) SetNodeOutputs(v []*ListNodeDependenciesResponseBodyPagingInfoNodesInputsNodeOutputs) *ListNodeDependenciesResponseBodyPagingInfoNodesInputs {
	s.NodeOutputs = v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesInputs) SetTables(v []*ListNodeDependenciesResponseBodyPagingInfoNodesInputsTables) *ListNodeDependenciesResponseBodyPagingInfoNodesInputs {
	s.Tables = v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesInputs) SetVariables(v []*ListNodeDependenciesResponseBodyPagingInfoNodesInputsVariables) *ListNodeDependenciesResponseBodyPagingInfoNodesInputs {
	s.Variables = v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesInputs) Validate() error {
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

type ListNodeDependenciesResponseBodyPagingInfoNodesInputsNodeOutputs struct {
	// The node output.
	//
	// example:
	//
	// 860438872620113XXXX
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesInputsNodeOutputs) String() string {
	return dara.Prettify(s)
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesInputsNodeOutputs) GoString() string {
	return s.String()
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesInputsNodeOutputs) GetData() *string {
	return s.Data
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesInputsNodeOutputs) SetData(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesInputsNodeOutputs {
	s.Data = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesInputsNodeOutputs) Validate() error {
	return dara.Validate(s)
}

type ListNodeDependenciesResponseBodyPagingInfoNodesInputsTables struct {
	// The table ID.
	//
	// example:
	//
	// odps.autotest.test_output_table_1
	Guid *string `json:"Guid,omitempty" xml:"Guid,omitempty"`
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesInputsTables) String() string {
	return dara.Prettify(s)
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesInputsTables) GoString() string {
	return s.String()
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesInputsTables) GetGuid() *string {
	return s.Guid
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesInputsTables) SetGuid(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesInputsTables {
	s.Guid = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesInputsTables) Validate() error {
	return dara.Validate(s)
}

type ListNodeDependenciesResponseBodyPagingInfoNodesInputsVariables struct {
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
	// 543218872620113XXXX
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The variable name.
	//
	// example:
	//
	// input
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The node to which the variable belongs.
	Node *ListNodeDependenciesResponseBodyPagingInfoNodesInputsVariablesNode `json:"Node,omitempty" xml:"Node,omitempty" type:"Struct"`
	// The scope of the variable. Valid values:
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
	// The variable name.
	//
	// example:
	//
	// 111
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesInputsVariables) String() string {
	return dara.Prettify(s)
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesInputsVariables) GoString() string {
	return s.String()
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesInputsVariables) GetArtifactType() *string {
	return s.ArtifactType
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesInputsVariables) GetId() *int64 {
	return s.Id
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesInputsVariables) GetName() *string {
	return s.Name
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesInputsVariables) GetNode() *ListNodeDependenciesResponseBodyPagingInfoNodesInputsVariablesNode {
	return s.Node
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesInputsVariables) GetScope() *string {
	return s.Scope
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesInputsVariables) GetType() *string {
	return s.Type
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesInputsVariables) GetValue() *string {
	return s.Value
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesInputsVariables) SetArtifactType(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesInputsVariables {
	s.ArtifactType = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesInputsVariables) SetId(v int64) *ListNodeDependenciesResponseBodyPagingInfoNodesInputsVariables {
	s.Id = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesInputsVariables) SetName(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesInputsVariables {
	s.Name = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesInputsVariables) SetNode(v *ListNodeDependenciesResponseBodyPagingInfoNodesInputsVariablesNode) *ListNodeDependenciesResponseBodyPagingInfoNodesInputsVariables {
	s.Node = v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesInputsVariables) SetScope(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesInputsVariables {
	s.Scope = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesInputsVariables) SetType(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesInputsVariables {
	s.Type = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesInputsVariables) SetValue(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesInputsVariables {
	s.Value = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesInputsVariables) Validate() error {
	if s.Node != nil {
		if err := s.Node.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListNodeDependenciesResponseBodyPagingInfoNodesInputsVariablesNode struct {
	// The output of the node.
	//
	// example:
	//
	// 860438872620113XXXX
	Output *string `json:"Output,omitempty" xml:"Output,omitempty"`
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesInputsVariablesNode) String() string {
	return dara.Prettify(s)
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesInputsVariablesNode) GoString() string {
	return s.String()
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesInputsVariablesNode) GetOutput() *string {
	return s.Output
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesInputsVariablesNode) SetOutput(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesInputsVariablesNode {
	s.Output = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesInputsVariablesNode) Validate() error {
	return dara.Validate(s)
}

type ListNodeDependenciesResponseBodyPagingInfoNodesOutputs struct {
	// The list of node outputs.
	NodeOutputs []*ListNodeDependenciesResponseBodyPagingInfoNodesOutputsNodeOutputs `json:"NodeOutputs,omitempty" xml:"NodeOutputs,omitempty" type:"Repeated"`
	// The table list.
	Tables []*ListNodeDependenciesResponseBodyPagingInfoNodesOutputsTables `json:"Tables,omitempty" xml:"Tables,omitempty" type:"Repeated"`
	// The variable list.
	Variables []*ListNodeDependenciesResponseBodyPagingInfoNodesOutputsVariables `json:"Variables,omitempty" xml:"Variables,omitempty" type:"Repeated"`
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesOutputs) String() string {
	return dara.Prettify(s)
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesOutputs) GoString() string {
	return s.String()
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesOutputs) GetNodeOutputs() []*ListNodeDependenciesResponseBodyPagingInfoNodesOutputsNodeOutputs {
	return s.NodeOutputs
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesOutputs) GetTables() []*ListNodeDependenciesResponseBodyPagingInfoNodesOutputsTables {
	return s.Tables
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesOutputs) GetVariables() []*ListNodeDependenciesResponseBodyPagingInfoNodesOutputsVariables {
	return s.Variables
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesOutputs) SetNodeOutputs(v []*ListNodeDependenciesResponseBodyPagingInfoNodesOutputsNodeOutputs) *ListNodeDependenciesResponseBodyPagingInfoNodesOutputs {
	s.NodeOutputs = v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesOutputs) SetTables(v []*ListNodeDependenciesResponseBodyPagingInfoNodesOutputsTables) *ListNodeDependenciesResponseBodyPagingInfoNodesOutputs {
	s.Tables = v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesOutputs) SetVariables(v []*ListNodeDependenciesResponseBodyPagingInfoNodesOutputsVariables) *ListNodeDependenciesResponseBodyPagingInfoNodesOutputs {
	s.Variables = v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesOutputs) Validate() error {
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

type ListNodeDependenciesResponseBodyPagingInfoNodesOutputsNodeOutputs struct {
	// The output of the node.
	//
	// example:
	//
	// 463497880880954XXXX
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesOutputsNodeOutputs) String() string {
	return dara.Prettify(s)
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesOutputsNodeOutputs) GoString() string {
	return s.String()
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesOutputsNodeOutputs) GetData() *string {
	return s.Data
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesOutputsNodeOutputs) SetData(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesOutputsNodeOutputs {
	s.Data = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesOutputsNodeOutputs) Validate() error {
	return dara.Validate(s)
}

type ListNodeDependenciesResponseBodyPagingInfoNodesOutputsTables struct {
	// The table ID.
	//
	// example:
	//
	// odps.autotest.test_output_table_1
	Guid *string `json:"Guid,omitempty" xml:"Guid,omitempty"`
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesOutputsTables) String() string {
	return dara.Prettify(s)
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesOutputsTables) GoString() string {
	return s.String()
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesOutputsTables) GetGuid() *string {
	return s.Guid
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesOutputsTables) SetGuid(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesOutputsTables {
	s.Guid = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesOutputsTables) Validate() error {
	return dara.Validate(s)
}

type ListNodeDependenciesResponseBodyPagingInfoNodesOutputsVariables struct {
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
	// 543217824470354XXXX
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The variable name.
	//
	// example:
	//
	// output
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The node to which the variable belongs.
	Node *ListNodeDependenciesResponseBodyPagingInfoNodesOutputsVariablesNode `json:"Node,omitempty" xml:"Node,omitempty" type:"Struct"`
	// The scope of the variable. Valid values:
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
	// The value of the variable.
	//
	// example:
	//
	// 111
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesOutputsVariables) String() string {
	return dara.Prettify(s)
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesOutputsVariables) GoString() string {
	return s.String()
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesOutputsVariables) GetArtifactType() *string {
	return s.ArtifactType
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesOutputsVariables) GetId() *int64 {
	return s.Id
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesOutputsVariables) GetName() *string {
	return s.Name
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesOutputsVariables) GetNode() *ListNodeDependenciesResponseBodyPagingInfoNodesOutputsVariablesNode {
	return s.Node
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesOutputsVariables) GetScope() *string {
	return s.Scope
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesOutputsVariables) GetType() *string {
	return s.Type
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesOutputsVariables) GetValue() *string {
	return s.Value
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesOutputsVariables) SetArtifactType(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesOutputsVariables {
	s.ArtifactType = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesOutputsVariables) SetId(v int64) *ListNodeDependenciesResponseBodyPagingInfoNodesOutputsVariables {
	s.Id = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesOutputsVariables) SetName(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesOutputsVariables {
	s.Name = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesOutputsVariables) SetNode(v *ListNodeDependenciesResponseBodyPagingInfoNodesOutputsVariablesNode) *ListNodeDependenciesResponseBodyPagingInfoNodesOutputsVariables {
	s.Node = v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesOutputsVariables) SetScope(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesOutputsVariables {
	s.Scope = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesOutputsVariables) SetType(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesOutputsVariables {
	s.Type = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesOutputsVariables) SetValue(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesOutputsVariables {
	s.Value = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesOutputsVariables) Validate() error {
	if s.Node != nil {
		if err := s.Node.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListNodeDependenciesResponseBodyPagingInfoNodesOutputsVariablesNode struct {
	// The node output corresponding to the variable.
	//
	// example:
	//
	// 463497880880954XXXX
	Output *string `json:"Output,omitempty" xml:"Output,omitempty"`
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesOutputsVariablesNode) String() string {
	return dara.Prettify(s)
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesOutputsVariablesNode) GoString() string {
	return s.String()
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesOutputsVariablesNode) GetOutput() *string {
	return s.Output
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesOutputsVariablesNode) SetOutput(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesOutputsVariablesNode {
	s.Output = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesOutputsVariablesNode) Validate() error {
	return dara.Validate(s)
}

type ListNodeDependenciesResponseBodyPagingInfoNodesRuntimeResource struct {
	// The resource group ID.
	//
	// example:
	//
	// S_res_group_XXXX_XXXX
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesRuntimeResource) String() string {
	return dara.Prettify(s)
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesRuntimeResource) GoString() string {
	return s.String()
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesRuntimeResource) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesRuntimeResource) SetResourceGroupId(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesRuntimeResource {
	s.ResourceGroupId = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesRuntimeResource) Validate() error {
	return dara.Validate(s)
}

type ListNodeDependenciesResponseBodyPagingInfoNodesScript struct {
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
	Runtime *ListNodeDependenciesResponseBodyPagingInfoNodesScriptRuntime `json:"Runtime,omitempty" xml:"Runtime,omitempty" type:"Struct"`
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesScript) String() string {
	return dara.Prettify(s)
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesScript) GoString() string {
	return s.String()
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesScript) GetId() *int64 {
	return s.Id
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesScript) GetPath() *string {
	return s.Path
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesScript) GetRuntime() *ListNodeDependenciesResponseBodyPagingInfoNodesScriptRuntime {
	return s.Runtime
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesScript) SetId(v int64) *ListNodeDependenciesResponseBodyPagingInfoNodesScript {
	s.Id = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesScript) SetPath(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesScript {
	s.Path = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesScript) SetRuntime(v *ListNodeDependenciesResponseBodyPagingInfoNodesScriptRuntime) *ListNodeDependenciesResponseBodyPagingInfoNodesScript {
	s.Runtime = v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesScript) Validate() error {
	if s.Runtime != nil {
		if err := s.Runtime.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListNodeDependenciesResponseBodyPagingInfoNodesScriptRuntime struct {
	// The command used to distinguish node types.
	//
	// example:
	//
	// ODPS_SQL
	Command *string `json:"Command,omitempty" xml:"Command,omitempty"`
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesScriptRuntime) String() string {
	return dara.Prettify(s)
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesScriptRuntime) GoString() string {
	return s.String()
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesScriptRuntime) GetCommand() *string {
	return s.Command
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesScriptRuntime) SetCommand(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesScriptRuntime {
	s.Command = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesScriptRuntime) Validate() error {
	return dara.Validate(s)
}

type ListNodeDependenciesResponseBodyPagingInfoNodesStrategy struct {
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
	// The interval between retries after failure. Unit: milliseconds.
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
	// The number of retries after failure.
	//
	// example:
	//
	// 3
	RerunTimes *int32 `json:"RerunTimes,omitempty" xml:"RerunTimes,omitempty"`
	// The timeout period. Unit: milliseconds.
	//
	// example:
	//
	// 0
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesStrategy) String() string {
	return dara.Prettify(s)
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesStrategy) GoString() string {
	return s.String()
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesStrategy) GetInstanceMode() *string {
	return s.InstanceMode
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesStrategy) GetRerunInterval() *int32 {
	return s.RerunInterval
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesStrategy) GetRerunMode() *string {
	return s.RerunMode
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesStrategy) GetRerunTimes() *int32 {
	return s.RerunTimes
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesStrategy) GetTimeout() *int32 {
	return s.Timeout
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesStrategy) SetInstanceMode(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesStrategy {
	s.InstanceMode = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesStrategy) SetRerunInterval(v int32) *ListNodeDependenciesResponseBodyPagingInfoNodesStrategy {
	s.RerunInterval = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesStrategy) SetRerunMode(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesStrategy {
	s.RerunMode = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesStrategy) SetRerunTimes(v int32) *ListNodeDependenciesResponseBodyPagingInfoNodesStrategy {
	s.RerunTimes = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesStrategy) SetTimeout(v int32) *ListNodeDependenciesResponseBodyPagingInfoNodesStrategy {
	s.Timeout = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesStrategy) Validate() error {
	return dara.Validate(s)
}

type ListNodeDependenciesResponseBodyPagingInfoNodesTags struct {
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

func (s ListNodeDependenciesResponseBodyPagingInfoNodesTags) String() string {
	return dara.Prettify(s)
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesTags) GoString() string {
	return s.String()
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesTags) GetKey() *string {
	return s.Key
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesTags) GetValue() *string {
	return s.Value
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesTags) SetKey(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesTags {
	s.Key = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesTags) SetValue(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesTags {
	s.Value = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesTags) Validate() error {
	return dara.Validate(s)
}

type ListNodeDependenciesResponseBodyPagingInfoNodesTrigger struct {
	// The cron expression for scheduling.
	//
	// example:
	//
	// 00 00 00 	- 	- ?
	Cron *string `json:"Cron,omitempty" xml:"Cron,omitempty"`
	// The effective end time of the schedule, in the format yyyy-MM-dd HH:mm:ss.
	//
	// example:
	//
	// 9999-01-01 00:00:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The unique identifier of the trigger.
	//
	// example:
	//
	// 543680677872062XXXX
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The effective start time of the schedule, in the format yyyy-MM-dd HH:mm:ss.
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
	// 	- Scheduler: periodic scheduling.
	//
	// 	- Manual: manual scheduling.
	//
	// 	- Streaming: streaming scheduler.
	//
	// example:
	//
	// Scheduler
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesTrigger) String() string {
	return dara.Prettify(s)
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesTrigger) GoString() string {
	return s.String()
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesTrigger) GetCron() *string {
	return s.Cron
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesTrigger) GetEndTime() *string {
	return s.EndTime
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesTrigger) GetId() *int64 {
	return s.Id
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesTrigger) GetStartTime() *string {
	return s.StartTime
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesTrigger) GetTimezone() *string {
	return s.Timezone
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesTrigger) GetType() *string {
	return s.Type
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesTrigger) SetCron(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesTrigger {
	s.Cron = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesTrigger) SetEndTime(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesTrigger {
	s.EndTime = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesTrigger) SetId(v int64) *ListNodeDependenciesResponseBodyPagingInfoNodesTrigger {
	s.Id = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesTrigger) SetStartTime(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesTrigger {
	s.StartTime = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesTrigger) SetTimezone(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesTrigger {
	s.Timezone = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesTrigger) SetType(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesTrigger {
	s.Type = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesTrigger) Validate() error {
	return dara.Validate(s)
}
