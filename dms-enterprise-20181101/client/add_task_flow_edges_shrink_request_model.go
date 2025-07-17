// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddTaskFlowEdgesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDagId(v int64) *AddTaskFlowEdgesShrinkRequest
	GetDagId() *int64
	SetEdgesShrink(v string) *AddTaskFlowEdgesShrinkRequest
	GetEdgesShrink() *string
	SetTid(v int64) *AddTaskFlowEdgesShrinkRequest
	GetTid() *int64
}

type AddTaskFlowEdgesShrinkRequest struct {
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
	EdgesShrink *string `json:"Edges,omitempty" xml:"Edges,omitempty"`
	// The ID of the tenant.
	//
	// > : To view the ID of the tenant, go to the Data Management (DMS) console and move the pointer over the profile picture in the upper-right corner. For more information, see [View information about the current tenant](https://help.aliyun.com/document_detail/181330.html).
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s AddTaskFlowEdgesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AddTaskFlowEdgesShrinkRequest) GoString() string {
	return s.String()
}

func (s *AddTaskFlowEdgesShrinkRequest) GetDagId() *int64 {
	return s.DagId
}

func (s *AddTaskFlowEdgesShrinkRequest) GetEdgesShrink() *string {
	return s.EdgesShrink
}

func (s *AddTaskFlowEdgesShrinkRequest) GetTid() *int64 {
	return s.Tid
}

func (s *AddTaskFlowEdgesShrinkRequest) SetDagId(v int64) *AddTaskFlowEdgesShrinkRequest {
	s.DagId = &v
	return s
}

func (s *AddTaskFlowEdgesShrinkRequest) SetEdgesShrink(v string) *AddTaskFlowEdgesShrinkRequest {
	s.EdgesShrink = &v
	return s
}

func (s *AddTaskFlowEdgesShrinkRequest) SetTid(v int64) *AddTaskFlowEdgesShrinkRequest {
	s.Tid = &v
	return s
}

func (s *AddTaskFlowEdgesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
