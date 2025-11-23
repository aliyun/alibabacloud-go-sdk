// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTaskFlowGraphResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *GetTaskFlowGraphResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetTaskFlowGraphResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *GetTaskFlowGraphResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetTaskFlowGraphResponseBody
	GetSuccess() *bool
	SetTaskFlowGraph(v *GetTaskFlowGraphResponseBodyTaskFlowGraph) *GetTaskFlowGraphResponseBody
	GetTaskFlowGraph() *GetTaskFlowGraphResponseBodyTaskFlowGraph
}

type GetTaskFlowGraphResponseBody struct {
	// The error code returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// D9D91166-A626-5F4E-9CA6-7AB10C59DBD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The list of DAG variables of the task flow.
	TaskFlowGraph *GetTaskFlowGraphResponseBodyTaskFlowGraph `json:"TaskFlowGraph,omitempty" xml:"TaskFlowGraph,omitempty" type:"Struct"`
}

func (s GetTaskFlowGraphResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTaskFlowGraphResponseBody) GoString() string {
	return s.String()
}

func (s *GetTaskFlowGraphResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetTaskFlowGraphResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetTaskFlowGraphResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTaskFlowGraphResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetTaskFlowGraphResponseBody) GetTaskFlowGraph() *GetTaskFlowGraphResponseBodyTaskFlowGraph {
	return s.TaskFlowGraph
}

func (s *GetTaskFlowGraphResponseBody) SetErrorCode(v string) *GetTaskFlowGraphResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetTaskFlowGraphResponseBody) SetErrorMessage(v string) *GetTaskFlowGraphResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetTaskFlowGraphResponseBody) SetRequestId(v string) *GetTaskFlowGraphResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTaskFlowGraphResponseBody) SetSuccess(v bool) *GetTaskFlowGraphResponseBody {
	s.Success = &v
	return s
}

func (s *GetTaskFlowGraphResponseBody) SetTaskFlowGraph(v *GetTaskFlowGraphResponseBodyTaskFlowGraph) *GetTaskFlowGraphResponseBody {
	s.TaskFlowGraph = v
	return s
}

func (s *GetTaskFlowGraphResponseBody) Validate() error {
	if s.TaskFlowGraph != nil {
		if err := s.TaskFlowGraph.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTaskFlowGraphResponseBodyTaskFlowGraph struct {
	// Indicates whether the task flow is editable. Valid values:
	//
	// - **true**: editable
	//
	// - **false**: non-editable
	//
	// example:
	//
	// true
	CanEdit *bool `json:"CanEdit,omitempty" xml:"CanEdit,omitempty"`
	// The name of the task flow.
	//
	// example:
	//
	// test
	DagName *string `json:"DagName,omitempty" xml:"DagName,omitempty"`
	// The list of task flow edges.
	Edges *GetTaskFlowGraphResponseBodyTaskFlowGraphEdges `json:"Edges,omitempty" xml:"Edges,omitempty" type:"Struct"`
	// The node list of the task flow.
	Nodes *GetTaskFlowGraphResponseBodyTaskFlowGraphNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Struct"`
	// The status of the task flow. Valid values:
	//
	// - **0**: invalid
	//
	// - **1**: not scheduled
	//
	// - **2**: to be scheduled
	//
	// example:
	//
	// 2
	Status *int64 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetTaskFlowGraphResponseBodyTaskFlowGraph) String() string {
	return dara.Prettify(s)
}

func (s GetTaskFlowGraphResponseBodyTaskFlowGraph) GoString() string {
	return s.String()
}

func (s *GetTaskFlowGraphResponseBodyTaskFlowGraph) GetCanEdit() *bool {
	return s.CanEdit
}

func (s *GetTaskFlowGraphResponseBodyTaskFlowGraph) GetDagName() *string {
	return s.DagName
}

func (s *GetTaskFlowGraphResponseBodyTaskFlowGraph) GetEdges() *GetTaskFlowGraphResponseBodyTaskFlowGraphEdges {
	return s.Edges
}

func (s *GetTaskFlowGraphResponseBodyTaskFlowGraph) GetNodes() *GetTaskFlowGraphResponseBodyTaskFlowGraphNodes {
	return s.Nodes
}

func (s *GetTaskFlowGraphResponseBodyTaskFlowGraph) GetStatus() *int64 {
	return s.Status
}

func (s *GetTaskFlowGraphResponseBodyTaskFlowGraph) SetCanEdit(v bool) *GetTaskFlowGraphResponseBodyTaskFlowGraph {
	s.CanEdit = &v
	return s
}

func (s *GetTaskFlowGraphResponseBodyTaskFlowGraph) SetDagName(v string) *GetTaskFlowGraphResponseBodyTaskFlowGraph {
	s.DagName = &v
	return s
}

func (s *GetTaskFlowGraphResponseBodyTaskFlowGraph) SetEdges(v *GetTaskFlowGraphResponseBodyTaskFlowGraphEdges) *GetTaskFlowGraphResponseBodyTaskFlowGraph {
	s.Edges = v
	return s
}

func (s *GetTaskFlowGraphResponseBodyTaskFlowGraph) SetNodes(v *GetTaskFlowGraphResponseBodyTaskFlowGraphNodes) *GetTaskFlowGraphResponseBodyTaskFlowGraph {
	s.Nodes = v
	return s
}

func (s *GetTaskFlowGraphResponseBodyTaskFlowGraph) SetStatus(v int64) *GetTaskFlowGraphResponseBodyTaskFlowGraph {
	s.Status = &v
	return s
}

func (s *GetTaskFlowGraphResponseBodyTaskFlowGraph) Validate() error {
	if s.Edges != nil {
		if err := s.Edges.Validate(); err != nil {
			return err
		}
	}
	if s.Nodes != nil {
		if err := s.Nodes.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTaskFlowGraphResponseBodyTaskFlowGraphEdges struct {
	Edge []*GetTaskFlowGraphResponseBodyTaskFlowGraphEdgesEdge `json:"Edge,omitempty" xml:"Edge,omitempty" type:"Repeated"`
}

func (s GetTaskFlowGraphResponseBodyTaskFlowGraphEdges) String() string {
	return dara.Prettify(s)
}

func (s GetTaskFlowGraphResponseBodyTaskFlowGraphEdges) GoString() string {
	return s.String()
}

func (s *GetTaskFlowGraphResponseBodyTaskFlowGraphEdges) GetEdge() []*GetTaskFlowGraphResponseBodyTaskFlowGraphEdgesEdge {
	return s.Edge
}

func (s *GetTaskFlowGraphResponseBodyTaskFlowGraphEdges) SetEdge(v []*GetTaskFlowGraphResponseBodyTaskFlowGraphEdgesEdge) *GetTaskFlowGraphResponseBodyTaskFlowGraphEdges {
	s.Edge = v
	return s
}

func (s *GetTaskFlowGraphResponseBodyTaskFlowGraphEdges) Validate() error {
	if s.Edge != nil {
		for _, item := range s.Edge {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetTaskFlowGraphResponseBodyTaskFlowGraphEdgesEdge struct {
	// The ID of the task flow.
	//
	// example:
	//
	// 15***
	DagId *int64 `json:"DagId,omitempty" xml:"DagId,omitempty"`
	// The ID of the task flow edge.
	//
	// example:
	//
	// 24***
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The ID of the end node on the edge.
	//
	// example:
	//
	// 44***
	NodeEnd *int64 `json:"NodeEnd,omitempty" xml:"NodeEnd,omitempty"`
	// The ID of the start node on the edge.
	//
	// example:
	//
	// 44***
	NodeFrom *int64 `json:"NodeFrom,omitempty" xml:"NodeFrom,omitempty"`
}

func (s GetTaskFlowGraphResponseBodyTaskFlowGraphEdgesEdge) String() string {
	return dara.Prettify(s)
}

func (s GetTaskFlowGraphResponseBodyTaskFlowGraphEdgesEdge) GoString() string {
	return s.String()
}

func (s *GetTaskFlowGraphResponseBodyTaskFlowGraphEdgesEdge) GetDagId() *int64 {
	return s.DagId
}

func (s *GetTaskFlowGraphResponseBodyTaskFlowGraphEdgesEdge) GetId() *int64 {
	return s.Id
}

func (s *GetTaskFlowGraphResponseBodyTaskFlowGraphEdgesEdge) GetNodeEnd() *int64 {
	return s.NodeEnd
}

func (s *GetTaskFlowGraphResponseBodyTaskFlowGraphEdgesEdge) GetNodeFrom() *int64 {
	return s.NodeFrom
}

func (s *GetTaskFlowGraphResponseBodyTaskFlowGraphEdgesEdge) SetDagId(v int64) *GetTaskFlowGraphResponseBodyTaskFlowGraphEdgesEdge {
	s.DagId = &v
	return s
}

func (s *GetTaskFlowGraphResponseBodyTaskFlowGraphEdgesEdge) SetId(v int64) *GetTaskFlowGraphResponseBodyTaskFlowGraphEdgesEdge {
	s.Id = &v
	return s
}

func (s *GetTaskFlowGraphResponseBodyTaskFlowGraphEdgesEdge) SetNodeEnd(v int64) *GetTaskFlowGraphResponseBodyTaskFlowGraphEdgesEdge {
	s.NodeEnd = &v
	return s
}

func (s *GetTaskFlowGraphResponseBodyTaskFlowGraphEdgesEdge) SetNodeFrom(v int64) *GetTaskFlowGraphResponseBodyTaskFlowGraphEdgesEdge {
	s.NodeFrom = &v
	return s
}

func (s *GetTaskFlowGraphResponseBodyTaskFlowGraphEdgesEdge) Validate() error {
	return dara.Validate(s)
}

type GetTaskFlowGraphResponseBodyTaskFlowGraphNodes struct {
	Node []*GetTaskFlowGraphResponseBodyTaskFlowGraphNodesNode `json:"Node,omitempty" xml:"Node,omitempty" type:"Repeated"`
}

func (s GetTaskFlowGraphResponseBodyTaskFlowGraphNodes) String() string {
	return dara.Prettify(s)
}

func (s GetTaskFlowGraphResponseBodyTaskFlowGraphNodes) GoString() string {
	return s.String()
}

func (s *GetTaskFlowGraphResponseBodyTaskFlowGraphNodes) GetNode() []*GetTaskFlowGraphResponseBodyTaskFlowGraphNodesNode {
	return s.Node
}

func (s *GetTaskFlowGraphResponseBodyTaskFlowGraphNodes) SetNode(v []*GetTaskFlowGraphResponseBodyTaskFlowGraphNodesNode) *GetTaskFlowGraphResponseBodyTaskFlowGraphNodes {
	s.Node = v
	return s
}

func (s *GetTaskFlowGraphResponseBodyTaskFlowGraphNodes) Validate() error {
	if s.Node != nil {
		for _, item := range s.Node {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetTaskFlowGraphResponseBodyTaskFlowGraphNodesNode struct {
	// The ID of the task flow.
	//
	// example:
	//
	// 15***
	DagId *int64 `json:"DagId,omitempty" xml:"DagId,omitempty"`
	// The position of the node in the DAG.
	//
	// example:
	//
	// {\\"x\\":541,\\"y\\":322,\\"layoutType\\":\\"Horizontal\\"}
	GraphParam *string `json:"GraphParam,omitempty" xml:"GraphParam,omitempty"`
	// The advanced configuration of the node.
	NodeConfig *string `json:"NodeConfig,omitempty" xml:"NodeConfig,omitempty"`
	// The configuration of the node.
	//
	// example:
	//
	// {ODI3OTNRVC****UHVFT29"}
	NodeContent *string `json:"NodeContent,omitempty" xml:"NodeContent,omitempty"`
	// The ID of the node.
	//
	// example:
	//
	// 44***
	NodeId *int64 `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The name of the node.
	//
	// example:
	//
	// test
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// The type of the node. For more information about the valid values for this parameter, see [NodeType parameter](https://help.aliyun.com/document_detail/424705.html).
	//
	// example:
	//
	// 36
	NodeType *int64 `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	// The time variables for the node.
	//
	// example:
	//
	// {\\"variables\\":[{\\"name\\":\\"Today\\",\\"pattern\\":\\"yyyy-MM-dd|+1d\\"}]}
	TimeVariables *string `json:"TimeVariables,omitempty" xml:"TimeVariables,omitempty"`
}

func (s GetTaskFlowGraphResponseBodyTaskFlowGraphNodesNode) String() string {
	return dara.Prettify(s)
}

func (s GetTaskFlowGraphResponseBodyTaskFlowGraphNodesNode) GoString() string {
	return s.String()
}

func (s *GetTaskFlowGraphResponseBodyTaskFlowGraphNodesNode) GetDagId() *int64 {
	return s.DagId
}

func (s *GetTaskFlowGraphResponseBodyTaskFlowGraphNodesNode) GetGraphParam() *string {
	return s.GraphParam
}

func (s *GetTaskFlowGraphResponseBodyTaskFlowGraphNodesNode) GetNodeConfig() *string {
	return s.NodeConfig
}

func (s *GetTaskFlowGraphResponseBodyTaskFlowGraphNodesNode) GetNodeContent() *string {
	return s.NodeContent
}

func (s *GetTaskFlowGraphResponseBodyTaskFlowGraphNodesNode) GetNodeId() *int64 {
	return s.NodeId
}

func (s *GetTaskFlowGraphResponseBodyTaskFlowGraphNodesNode) GetNodeName() *string {
	return s.NodeName
}

func (s *GetTaskFlowGraphResponseBodyTaskFlowGraphNodesNode) GetNodeType() *int64 {
	return s.NodeType
}

func (s *GetTaskFlowGraphResponseBodyTaskFlowGraphNodesNode) GetTimeVariables() *string {
	return s.TimeVariables
}

func (s *GetTaskFlowGraphResponseBodyTaskFlowGraphNodesNode) SetDagId(v int64) *GetTaskFlowGraphResponseBodyTaskFlowGraphNodesNode {
	s.DagId = &v
	return s
}

func (s *GetTaskFlowGraphResponseBodyTaskFlowGraphNodesNode) SetGraphParam(v string) *GetTaskFlowGraphResponseBodyTaskFlowGraphNodesNode {
	s.GraphParam = &v
	return s
}

func (s *GetTaskFlowGraphResponseBodyTaskFlowGraphNodesNode) SetNodeConfig(v string) *GetTaskFlowGraphResponseBodyTaskFlowGraphNodesNode {
	s.NodeConfig = &v
	return s
}

func (s *GetTaskFlowGraphResponseBodyTaskFlowGraphNodesNode) SetNodeContent(v string) *GetTaskFlowGraphResponseBodyTaskFlowGraphNodesNode {
	s.NodeContent = &v
	return s
}

func (s *GetTaskFlowGraphResponseBodyTaskFlowGraphNodesNode) SetNodeId(v int64) *GetTaskFlowGraphResponseBodyTaskFlowGraphNodesNode {
	s.NodeId = &v
	return s
}

func (s *GetTaskFlowGraphResponseBodyTaskFlowGraphNodesNode) SetNodeName(v string) *GetTaskFlowGraphResponseBodyTaskFlowGraphNodesNode {
	s.NodeName = &v
	return s
}

func (s *GetTaskFlowGraphResponseBodyTaskFlowGraphNodesNode) SetNodeType(v int64) *GetTaskFlowGraphResponseBodyTaskFlowGraphNodesNode {
	s.NodeType = &v
	return s
}

func (s *GetTaskFlowGraphResponseBodyTaskFlowGraphNodesNode) SetTimeVariables(v string) *GetTaskFlowGraphResponseBodyTaskFlowGraphNodesNode {
	s.TimeVariables = &v
	return s
}

func (s *GetTaskFlowGraphResponseBodyTaskFlowGraphNodesNode) Validate() error {
	return dara.Validate(s)
}
