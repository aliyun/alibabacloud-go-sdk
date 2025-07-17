// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTaskFlowEdgesByConditionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDagId(v int64) *ListTaskFlowEdgesByConditionRequest
	GetDagId() *int64
	SetId(v int64) *ListTaskFlowEdgesByConditionRequest
	GetId() *int64
	SetNodeEnd(v int64) *ListTaskFlowEdgesByConditionRequest
	GetNodeEnd() *int64
	SetNodeFrom(v int64) *ListTaskFlowEdgesByConditionRequest
	GetNodeFrom() *int64
	SetTid(v int64) *ListTaskFlowEdgesByConditionRequest
	GetTid() *int64
}

type ListTaskFlowEdgesByConditionRequest struct {
	// The ID of the task flow. You can call the [ListTaskFlow](https://help.aliyun.com/document_detail/424565.html) or [ListLhTaskFlowAndScenario](https://help.aliyun.com/document_detail/426672.html) operation to query the task flow ID.
	//
	// This parameter is required.
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
	// The ID of the tenant.
	//
	// >  To view the ID of the tenant, go to the Data Management (DMS) console and move the pointer over the profile picture in the upper-right corner. For more information, see [View information about the current tenant](https://help.aliyun.com/document_detail/181330.html).
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s ListTaskFlowEdgesByConditionRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTaskFlowEdgesByConditionRequest) GoString() string {
	return s.String()
}

func (s *ListTaskFlowEdgesByConditionRequest) GetDagId() *int64 {
	return s.DagId
}

func (s *ListTaskFlowEdgesByConditionRequest) GetId() *int64 {
	return s.Id
}

func (s *ListTaskFlowEdgesByConditionRequest) GetNodeEnd() *int64 {
	return s.NodeEnd
}

func (s *ListTaskFlowEdgesByConditionRequest) GetNodeFrom() *int64 {
	return s.NodeFrom
}

func (s *ListTaskFlowEdgesByConditionRequest) GetTid() *int64 {
	return s.Tid
}

func (s *ListTaskFlowEdgesByConditionRequest) SetDagId(v int64) *ListTaskFlowEdgesByConditionRequest {
	s.DagId = &v
	return s
}

func (s *ListTaskFlowEdgesByConditionRequest) SetId(v int64) *ListTaskFlowEdgesByConditionRequest {
	s.Id = &v
	return s
}

func (s *ListTaskFlowEdgesByConditionRequest) SetNodeEnd(v int64) *ListTaskFlowEdgesByConditionRequest {
	s.NodeEnd = &v
	return s
}

func (s *ListTaskFlowEdgesByConditionRequest) SetNodeFrom(v int64) *ListTaskFlowEdgesByConditionRequest {
	s.NodeFrom = &v
	return s
}

func (s *ListTaskFlowEdgesByConditionRequest) SetTid(v int64) *ListTaskFlowEdgesByConditionRequest {
	s.Tid = &v
	return s
}

func (s *ListTaskFlowEdgesByConditionRequest) Validate() error {
	return dara.Validate(s)
}
