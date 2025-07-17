// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTaskFlowCooperatorsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCooperatorIdsShrink(v string) *UpdateTaskFlowCooperatorsShrinkRequest
	GetCooperatorIdsShrink() *string
	SetDagId(v int64) *UpdateTaskFlowCooperatorsShrinkRequest
	GetDagId() *int64
	SetTid(v int64) *UpdateTaskFlowCooperatorsShrinkRequest
	GetTid() *int64
}

type UpdateTaskFlowCooperatorsShrinkRequest struct {
	// The IDs of the users who are involved in the task flow to be updated.
	CooperatorIdsShrink *string `json:"CooperatorIds,omitempty" xml:"CooperatorIds,omitempty"`
	// The ID of the task flow. You can call the [ListTaskFlow](https://help.aliyun.com/document_detail/424565.html) or [ListLhTaskFlowAndScenario](https://help.aliyun.com/document_detail/426672.html) operation to query the task flow ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 11****
	DagId *int64 `json:"DagId,omitempty" xml:"DagId,omitempty"`
	// The ID of the tenant.
	//
	// > :To view the ID of the tenant, go to the Data Management (DMS) console and move the pointer over the profile picture in the upper-right corner. For more information, see [View information about the current tenant](https://help.aliyun.com/document_detail/181330.html).
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s UpdateTaskFlowCooperatorsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTaskFlowCooperatorsShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateTaskFlowCooperatorsShrinkRequest) GetCooperatorIdsShrink() *string {
	return s.CooperatorIdsShrink
}

func (s *UpdateTaskFlowCooperatorsShrinkRequest) GetDagId() *int64 {
	return s.DagId
}

func (s *UpdateTaskFlowCooperatorsShrinkRequest) GetTid() *int64 {
	return s.Tid
}

func (s *UpdateTaskFlowCooperatorsShrinkRequest) SetCooperatorIdsShrink(v string) *UpdateTaskFlowCooperatorsShrinkRequest {
	s.CooperatorIdsShrink = &v
	return s
}

func (s *UpdateTaskFlowCooperatorsShrinkRequest) SetDagId(v int64) *UpdateTaskFlowCooperatorsShrinkRequest {
	s.DagId = &v
	return s
}

func (s *UpdateTaskFlowCooperatorsShrinkRequest) SetTid(v int64) *UpdateTaskFlowCooperatorsShrinkRequest {
	s.Tid = &v
	return s
}

func (s *UpdateTaskFlowCooperatorsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
