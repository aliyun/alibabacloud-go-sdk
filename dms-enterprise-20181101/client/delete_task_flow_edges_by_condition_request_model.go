// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTaskFlowEdgesByConditionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDagId(v int64) *DeleteTaskFlowEdgesByConditionRequest
	GetDagId() *int64
	SetId(v int64) *DeleteTaskFlowEdgesByConditionRequest
	GetId() *int64
	SetNodeEnd(v int64) *DeleteTaskFlowEdgesByConditionRequest
	GetNodeEnd() *int64
	SetNodeFrom(v int64) *DeleteTaskFlowEdgesByConditionRequest
	GetNodeFrom() *int64
	SetTid(v int64) *DeleteTaskFlowEdgesByConditionRequest
	GetTid() *int64
}

type DeleteTaskFlowEdgesByConditionRequest struct {
	// The ID of the task flow. You can call the [ListTaskFlow](https://help.aliyun.com/document_detail/424565.html) or [ListLhTaskFlowAndScenario](https://help.aliyun.com/document_detail/426672.html) operation to query the task flow ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 15****
	DagId *int64 `json:"DagId,omitempty" xml:"DagId,omitempty"`
	// The ID of the task flow edge to delete.
	//
	// example:
	//
	// 24***
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The ID of the end node of the edge to delete.
	//
	// example:
	//
	// 44***
	NodeEnd *int64 `json:"NodeEnd,omitempty" xml:"NodeEnd,omitempty"`
	// The ID of the start node on the edge to delete.
	//
	// example:
	//
	// 44***
	NodeFrom *int64 `json:"NodeFrom,omitempty" xml:"NodeFrom,omitempty"`
	// The ID of the tenant.
	//
	// >  To view the ID of the tenant, go to the Data Management (DMS) console and move the pointer over the profile picture in the upper-right corner. For more information, see [View information about the current tenant](https://help.aliyun.com/document_detail/181330.html).
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s DeleteTaskFlowEdgesByConditionRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteTaskFlowEdgesByConditionRequest) GoString() string {
	return s.String()
}

func (s *DeleteTaskFlowEdgesByConditionRequest) GetDagId() *int64 {
	return s.DagId
}

func (s *DeleteTaskFlowEdgesByConditionRequest) GetId() *int64 {
	return s.Id
}

func (s *DeleteTaskFlowEdgesByConditionRequest) GetNodeEnd() *int64 {
	return s.NodeEnd
}

func (s *DeleteTaskFlowEdgesByConditionRequest) GetNodeFrom() *int64 {
	return s.NodeFrom
}

func (s *DeleteTaskFlowEdgesByConditionRequest) GetTid() *int64 {
	return s.Tid
}

func (s *DeleteTaskFlowEdgesByConditionRequest) SetDagId(v int64) *DeleteTaskFlowEdgesByConditionRequest {
	s.DagId = &v
	return s
}

func (s *DeleteTaskFlowEdgesByConditionRequest) SetId(v int64) *DeleteTaskFlowEdgesByConditionRequest {
	s.Id = &v
	return s
}

func (s *DeleteTaskFlowEdgesByConditionRequest) SetNodeEnd(v int64) *DeleteTaskFlowEdgesByConditionRequest {
	s.NodeEnd = &v
	return s
}

func (s *DeleteTaskFlowEdgesByConditionRequest) SetNodeFrom(v int64) *DeleteTaskFlowEdgesByConditionRequest {
	s.NodeFrom = &v
	return s
}

func (s *DeleteTaskFlowEdgesByConditionRequest) SetTid(v int64) *DeleteTaskFlowEdgesByConditionRequest {
	s.Tid = &v
	return s
}

func (s *DeleteTaskFlowEdgesByConditionRequest) Validate() error {
	return dara.Validate(s)
}
