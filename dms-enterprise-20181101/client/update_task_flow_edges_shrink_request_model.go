// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTaskFlowEdgesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDagId(v int64) *UpdateTaskFlowEdgesShrinkRequest
	GetDagId() *int64
	SetEdgesShrink(v string) *UpdateTaskFlowEdgesShrinkRequest
	GetEdgesShrink() *string
	SetTid(v int64) *UpdateTaskFlowEdgesShrinkRequest
	GetTid() *int64
}

type UpdateTaskFlowEdgesShrinkRequest struct {
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
	EdgesShrink *string `json:"Edges,omitempty" xml:"Edges,omitempty"`
	// The tenant ID.
	//
	// > To view the tenant ID, move the pointer over the profile picture in the upper-right corner of the Data Management (DMS) console. For more information, see [Manage DMS tenants](https://help.aliyun.com/document_detail/181330.html).
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s UpdateTaskFlowEdgesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTaskFlowEdgesShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateTaskFlowEdgesShrinkRequest) GetDagId() *int64 {
	return s.DagId
}

func (s *UpdateTaskFlowEdgesShrinkRequest) GetEdgesShrink() *string {
	return s.EdgesShrink
}

func (s *UpdateTaskFlowEdgesShrinkRequest) GetTid() *int64 {
	return s.Tid
}

func (s *UpdateTaskFlowEdgesShrinkRequest) SetDagId(v int64) *UpdateTaskFlowEdgesShrinkRequest {
	s.DagId = &v
	return s
}

func (s *UpdateTaskFlowEdgesShrinkRequest) SetEdgesShrink(v string) *UpdateTaskFlowEdgesShrinkRequest {
	s.EdgesShrink = &v
	return s
}

func (s *UpdateTaskFlowEdgesShrinkRequest) SetTid(v int64) *UpdateTaskFlowEdgesShrinkRequest {
	s.Tid = &v
	return s
}

func (s *UpdateTaskFlowEdgesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
