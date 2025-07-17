// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTaskFlowRelationsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDagId(v int64) *UpdateTaskFlowRelationsShrinkRequest
	GetDagId() *int64
	SetEdgesShrink(v string) *UpdateTaskFlowRelationsShrinkRequest
	GetEdgesShrink() *string
	SetTid(v int64) *UpdateTaskFlowRelationsShrinkRequest
	GetTid() *int64
}

type UpdateTaskFlowRelationsShrinkRequest struct {
	// The ID of the task flow. You can call the [ListTaskFlow](https://help.aliyun.com/document_detail/424565.html) or [ListLhTaskFlowAndScenario](https://help.aliyun.com/document_detail/426672.html) operation to query the task flow ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 15***
	DagId *int64 `json:"DagId,omitempty" xml:"DagId,omitempty"`
	// The list of task flow edges to be updated.
	EdgesShrink *string `json:"Edges,omitempty" xml:"Edges,omitempty"`
	// The ID of the tenant.
	//
	// >  To view the ID of the tenant, go to the Data Management (DMS) console and move the pointer over the profile picture in the upper-right corner. For more information, see [View information about the current tenant](https://help.aliyun.com/document_detail/181330.html).
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s UpdateTaskFlowRelationsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTaskFlowRelationsShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateTaskFlowRelationsShrinkRequest) GetDagId() *int64 {
	return s.DagId
}

func (s *UpdateTaskFlowRelationsShrinkRequest) GetEdgesShrink() *string {
	return s.EdgesShrink
}

func (s *UpdateTaskFlowRelationsShrinkRequest) GetTid() *int64 {
	return s.Tid
}

func (s *UpdateTaskFlowRelationsShrinkRequest) SetDagId(v int64) *UpdateTaskFlowRelationsShrinkRequest {
	s.DagId = &v
	return s
}

func (s *UpdateTaskFlowRelationsShrinkRequest) SetEdgesShrink(v string) *UpdateTaskFlowRelationsShrinkRequest {
	s.EdgesShrink = &v
	return s
}

func (s *UpdateTaskFlowRelationsShrinkRequest) SetTid(v int64) *UpdateTaskFlowRelationsShrinkRequest {
	s.Tid = &v
	return s
}

func (s *UpdateTaskFlowRelationsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
