// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTaskFlowEdgesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDagId(v int64) *UpdateTaskFlowEdgesRequest
	GetDagId() *int64
	SetEdges(v []*UpdateTaskFlowEdgesRequestEdges) *UpdateTaskFlowEdgesRequest
	GetEdges() []*UpdateTaskFlowEdgesRequestEdges
	SetTid(v int64) *UpdateTaskFlowEdgesRequest
	GetTid() *int64
}

type UpdateTaskFlowEdgesRequest struct {
	// The task flow ID. You can call the [ListTaskFlow](https://help.aliyun.com/document_detail/424565.html) or [ListLhTaskFlowAndScenario](https://help.aliyun.com/document_detail/426672.html) operation to query the task flow ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 15***
	DagId *int64 `json:"DagId,omitempty" xml:"DagId,omitempty"`
	// The list of updated task flow edges.
	//
	// This parameter is required.
	Edges []*UpdateTaskFlowEdgesRequestEdges `json:"Edges,omitempty" xml:"Edges,omitempty" type:"Repeated"`
	// The tenant ID.
	//
	// > To view the tenant ID, move the pointer over the profile picture in the upper-right corner of the Data Management (DMS) console. For more information, see [Manage DMS tenants](https://help.aliyun.com/document_detail/181330.html).
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s UpdateTaskFlowEdgesRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTaskFlowEdgesRequest) GoString() string {
	return s.String()
}

func (s *UpdateTaskFlowEdgesRequest) GetDagId() *int64 {
	return s.DagId
}

func (s *UpdateTaskFlowEdgesRequest) GetEdges() []*UpdateTaskFlowEdgesRequestEdges {
	return s.Edges
}

func (s *UpdateTaskFlowEdgesRequest) GetTid() *int64 {
	return s.Tid
}

func (s *UpdateTaskFlowEdgesRequest) SetDagId(v int64) *UpdateTaskFlowEdgesRequest {
	s.DagId = &v
	return s
}

func (s *UpdateTaskFlowEdgesRequest) SetEdges(v []*UpdateTaskFlowEdgesRequestEdges) *UpdateTaskFlowEdgesRequest {
	s.Edges = v
	return s
}

func (s *UpdateTaskFlowEdgesRequest) SetTid(v int64) *UpdateTaskFlowEdgesRequest {
	s.Tid = &v
	return s
}

func (s *UpdateTaskFlowEdgesRequest) Validate() error {
	return dara.Validate(s)
}

type UpdateTaskFlowEdgesRequestEdges struct {
	// The ID of the task flow edge.
	//
	// This parameter is required.
	//
	// example:
	//
	// 24***
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The ID of the end node of the edge.
	//
	// This parameter is required.
	//
	// example:
	//
	// 44***
	NodeEnd *int64 `json:"NodeEnd,omitempty" xml:"NodeEnd,omitempty"`
	// The ID of the start node of the edge.
	//
	// This parameter is required.
	//
	// example:
	//
	// 44***
	NodeFrom *int64 `json:"NodeFrom,omitempty" xml:"NodeFrom,omitempty"`
}

func (s UpdateTaskFlowEdgesRequestEdges) String() string {
	return dara.Prettify(s)
}

func (s UpdateTaskFlowEdgesRequestEdges) GoString() string {
	return s.String()
}

func (s *UpdateTaskFlowEdgesRequestEdges) GetId() *int64 {
	return s.Id
}

func (s *UpdateTaskFlowEdgesRequestEdges) GetNodeEnd() *int64 {
	return s.NodeEnd
}

func (s *UpdateTaskFlowEdgesRequestEdges) GetNodeFrom() *int64 {
	return s.NodeFrom
}

func (s *UpdateTaskFlowEdgesRequestEdges) SetId(v int64) *UpdateTaskFlowEdgesRequestEdges {
	s.Id = &v
	return s
}

func (s *UpdateTaskFlowEdgesRequestEdges) SetNodeEnd(v int64) *UpdateTaskFlowEdgesRequestEdges {
	s.NodeEnd = &v
	return s
}

func (s *UpdateTaskFlowEdgesRequestEdges) SetNodeFrom(v int64) *UpdateTaskFlowEdgesRequestEdges {
	s.NodeFrom = &v
	return s
}

func (s *UpdateTaskFlowEdgesRequestEdges) Validate() error {
	return dara.Validate(s)
}
