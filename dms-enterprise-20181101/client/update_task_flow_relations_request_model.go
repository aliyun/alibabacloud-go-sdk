// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTaskFlowRelationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDagId(v int64) *UpdateTaskFlowRelationsRequest
	GetDagId() *int64
	SetEdges(v []*UpdateTaskFlowRelationsRequestEdges) *UpdateTaskFlowRelationsRequest
	GetEdges() []*UpdateTaskFlowRelationsRequestEdges
	SetTid(v int64) *UpdateTaskFlowRelationsRequest
	GetTid() *int64
}

type UpdateTaskFlowRelationsRequest struct {
	// The ID of the task flow. You can call the [ListTaskFlow](https://help.aliyun.com/document_detail/424565.html) or [ListLhTaskFlowAndScenario](https://help.aliyun.com/document_detail/426672.html) operation to query the task flow ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 15***
	DagId *int64 `json:"DagId,omitempty" xml:"DagId,omitempty"`
	// The list of task flow edges to be updated.
	Edges []*UpdateTaskFlowRelationsRequestEdges `json:"Edges,omitempty" xml:"Edges,omitempty" type:"Repeated"`
	// The ID of the tenant.
	//
	// >  To view the ID of the tenant, go to the Data Management (DMS) console and move the pointer over the profile picture in the upper-right corner. For more information, see [View information about the current tenant](https://help.aliyun.com/document_detail/181330.html).
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s UpdateTaskFlowRelationsRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTaskFlowRelationsRequest) GoString() string {
	return s.String()
}

func (s *UpdateTaskFlowRelationsRequest) GetDagId() *int64 {
	return s.DagId
}

func (s *UpdateTaskFlowRelationsRequest) GetEdges() []*UpdateTaskFlowRelationsRequestEdges {
	return s.Edges
}

func (s *UpdateTaskFlowRelationsRequest) GetTid() *int64 {
	return s.Tid
}

func (s *UpdateTaskFlowRelationsRequest) SetDagId(v int64) *UpdateTaskFlowRelationsRequest {
	s.DagId = &v
	return s
}

func (s *UpdateTaskFlowRelationsRequest) SetEdges(v []*UpdateTaskFlowRelationsRequestEdges) *UpdateTaskFlowRelationsRequest {
	s.Edges = v
	return s
}

func (s *UpdateTaskFlowRelationsRequest) SetTid(v int64) *UpdateTaskFlowRelationsRequest {
	s.Tid = &v
	return s
}

func (s *UpdateTaskFlowRelationsRequest) Validate() error {
	return dara.Validate(s)
}

type UpdateTaskFlowRelationsRequestEdges struct {
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

func (s UpdateTaskFlowRelationsRequestEdges) String() string {
	return dara.Prettify(s)
}

func (s UpdateTaskFlowRelationsRequestEdges) GoString() string {
	return s.String()
}

func (s *UpdateTaskFlowRelationsRequestEdges) GetId() *int64 {
	return s.Id
}

func (s *UpdateTaskFlowRelationsRequestEdges) GetNodeEnd() *int64 {
	return s.NodeEnd
}

func (s *UpdateTaskFlowRelationsRequestEdges) GetNodeFrom() *int64 {
	return s.NodeFrom
}

func (s *UpdateTaskFlowRelationsRequestEdges) SetId(v int64) *UpdateTaskFlowRelationsRequestEdges {
	s.Id = &v
	return s
}

func (s *UpdateTaskFlowRelationsRequestEdges) SetNodeEnd(v int64) *UpdateTaskFlowRelationsRequestEdges {
	s.NodeEnd = &v
	return s
}

func (s *UpdateTaskFlowRelationsRequestEdges) SetNodeFrom(v int64) *UpdateTaskFlowRelationsRequestEdges {
	s.NodeFrom = &v
	return s
}

func (s *UpdateTaskFlowRelationsRequestEdges) Validate() error {
	return dara.Validate(s)
}
