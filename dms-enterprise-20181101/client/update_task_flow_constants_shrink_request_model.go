// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTaskFlowConstantsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDagConstantsShrink(v string) *UpdateTaskFlowConstantsShrinkRequest
	GetDagConstantsShrink() *string
	SetDagId(v int64) *UpdateTaskFlowConstantsShrinkRequest
	GetDagId() *int64
	SetTid(v int64) *UpdateTaskFlowConstantsShrinkRequest
	GetTid() *int64
}

type UpdateTaskFlowConstantsShrinkRequest struct {
	// The constants for the task flow.
	DagConstantsShrink *string `json:"DagConstants,omitempty" xml:"DagConstants,omitempty"`
	// The ID of the task flow. You can call the [ListTaskFlow](https://help.aliyun.com/document_detail/424565.html) or [ListLhTaskFlowAndScenario](https://help.aliyun.com/document_detail/426672.html) operation to query the task flow ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3****
	DagId *int64 `json:"DagId,omitempty" xml:"DagId,omitempty"`
	// The ID of the tenant. You can call the [GetUserActiveTenant](https://help.aliyun.com/document_detail/198073.html) or [ListUserTenants](https://help.aliyun.com/document_detail/198074.html) operation to query the tenant ID.
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s UpdateTaskFlowConstantsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTaskFlowConstantsShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateTaskFlowConstantsShrinkRequest) GetDagConstantsShrink() *string {
	return s.DagConstantsShrink
}

func (s *UpdateTaskFlowConstantsShrinkRequest) GetDagId() *int64 {
	return s.DagId
}

func (s *UpdateTaskFlowConstantsShrinkRequest) GetTid() *int64 {
	return s.Tid
}

func (s *UpdateTaskFlowConstantsShrinkRequest) SetDagConstantsShrink(v string) *UpdateTaskFlowConstantsShrinkRequest {
	s.DagConstantsShrink = &v
	return s
}

func (s *UpdateTaskFlowConstantsShrinkRequest) SetDagId(v int64) *UpdateTaskFlowConstantsShrinkRequest {
	s.DagId = &v
	return s
}

func (s *UpdateTaskFlowConstantsShrinkRequest) SetTid(v int64) *UpdateTaskFlowConstantsShrinkRequest {
	s.Tid = &v
	return s
}

func (s *UpdateTaskFlowConstantsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
