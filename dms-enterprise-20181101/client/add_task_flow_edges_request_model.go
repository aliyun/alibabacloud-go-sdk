// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddTaskFlowEdgesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDagId(v int64) *AddTaskFlowEdgesRequest
	GetDagId() *int64
	SetEdges(v []*AddTaskFlowEdgesRequestEdges) *AddTaskFlowEdgesRequest
	GetEdges() []*AddTaskFlowEdgesRequestEdges
	SetTid(v int64) *AddTaskFlowEdgesRequest
	GetTid() *int64
}

type AddTaskFlowEdgesRequest struct {
	// The ID of the task flow. You can call the [ListTaskFlow](https://help.aliyun.com/document_detail/424565.html) or [ListLhTaskFlowAndScenario](https://help.aliyun.com/document_detail/426672.html) operation to query the task flow ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 15***
	DagId *int64 `json:"DagId,omitempty" xml:"DagId,omitempty"`
	// The list of edges of the task flow.
	//
	// This parameter is required.
	Edges []*AddTaskFlowEdgesRequestEdges `json:"Edges,omitempty" xml:"Edges,omitempty" type:"Repeated"`
	// The ID of the tenant.
	//
	// > : To view the ID of the tenant, go to the Data Management (DMS) console and move the pointer over the profile picture in the upper-right corner. For more information, see [View information about the current tenant](https://help.aliyun.com/document_detail/181330.html).
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s AddTaskFlowEdgesRequest) String() string {
	return dara.Prettify(s)
}

func (s AddTaskFlowEdgesRequest) GoString() string {
	return s.String()
}

func (s *AddTaskFlowEdgesRequest) GetDagId() *int64 {
	return s.DagId
}

func (s *AddTaskFlowEdgesRequest) GetEdges() []*AddTaskFlowEdgesRequestEdges {
	return s.Edges
}

func (s *AddTaskFlowEdgesRequest) GetTid() *int64 {
	return s.Tid
}

func (s *AddTaskFlowEdgesRequest) SetDagId(v int64) *AddTaskFlowEdgesRequest {
	s.DagId = &v
	return s
}

func (s *AddTaskFlowEdgesRequest) SetEdges(v []*AddTaskFlowEdgesRequestEdges) *AddTaskFlowEdgesRequest {
	s.Edges = v
	return s
}

func (s *AddTaskFlowEdgesRequest) SetTid(v int64) *AddTaskFlowEdgesRequest {
	s.Tid = &v
	return s
}

func (s *AddTaskFlowEdgesRequest) Validate() error {
	return dara.Validate(s)
}

type AddTaskFlowEdgesRequestEdges struct {
	// The ID of the node where the end node of the edge is located.
	//
	// This parameter is required.
	//
	// example:
	//
	// 44***
	NodeEnd *int64 `json:"NodeEnd,omitempty" xml:"NodeEnd,omitempty"`
	// The ID of the node where the start node of the edge is located.
	//
	// This parameter is required.
	//
	// example:
	//
	// 44***
	NodeFrom *int64 `json:"NodeFrom,omitempty" xml:"NodeFrom,omitempty"`
}

func (s AddTaskFlowEdgesRequestEdges) String() string {
	return dara.Prettify(s)
}

func (s AddTaskFlowEdgesRequestEdges) GoString() string {
	return s.String()
}

func (s *AddTaskFlowEdgesRequestEdges) GetNodeEnd() *int64 {
	return s.NodeEnd
}

func (s *AddTaskFlowEdgesRequestEdges) GetNodeFrom() *int64 {
	return s.NodeFrom
}

func (s *AddTaskFlowEdgesRequestEdges) SetNodeEnd(v int64) *AddTaskFlowEdgesRequestEdges {
	s.NodeEnd = &v
	return s
}

func (s *AddTaskFlowEdgesRequestEdges) SetNodeFrom(v int64) *AddTaskFlowEdgesRequestEdges {
	s.NodeFrom = &v
	return s
}

func (s *AddTaskFlowEdgesRequestEdges) Validate() error {
	return dara.Validate(s)
}
