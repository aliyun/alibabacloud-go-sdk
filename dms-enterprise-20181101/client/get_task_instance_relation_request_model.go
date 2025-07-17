// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTaskInstanceRelationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDagId(v int64) *GetTaskInstanceRelationRequest
	GetDagId() *int64
	SetDagInstanceId(v int64) *GetTaskInstanceRelationRequest
	GetDagInstanceId() *int64
	SetTid(v int64) *GetTaskInstanceRelationRequest
	GetTid() *int64
}

type GetTaskInstanceRelationRequest struct {
	// The ID of the task flow. You can call the [ListTaskFlow](https://help.aliyun.com/document_detail/424565.html) or [ListLhTaskFlowAndScenario](https://help.aliyun.com/document_detail/426672.html) operation to obtain the ID of the task flow.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7***
	DagId *int64 `json:"DagId,omitempty" xml:"DagId,omitempty"`
	// The ID of the execution record of the task flow. You can call the [ListTaskFlowInstance](https://help.aliyun.com/document_detail/424689.html) operation to obtain the execution record ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 47****
	DagInstanceId *int64 `json:"DagInstanceId,omitempty" xml:"DagInstanceId,omitempty"`
	// The ID of the tenant. You can call the [GetUserActiveTenant](https://help.aliyun.com/document_detail/198073.html) or [ListUserTenants](https://help.aliyun.com/document_detail/198074.html) operation to obtain the tenant ID.
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s GetTaskInstanceRelationRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTaskInstanceRelationRequest) GoString() string {
	return s.String()
}

func (s *GetTaskInstanceRelationRequest) GetDagId() *int64 {
	return s.DagId
}

func (s *GetTaskInstanceRelationRequest) GetDagInstanceId() *int64 {
	return s.DagInstanceId
}

func (s *GetTaskInstanceRelationRequest) GetTid() *int64 {
	return s.Tid
}

func (s *GetTaskInstanceRelationRequest) SetDagId(v int64) *GetTaskInstanceRelationRequest {
	s.DagId = &v
	return s
}

func (s *GetTaskInstanceRelationRequest) SetDagInstanceId(v int64) *GetTaskInstanceRelationRequest {
	s.DagInstanceId = &v
	return s
}

func (s *GetTaskInstanceRelationRequest) SetTid(v int64) *GetTaskInstanceRelationRequest {
	s.Tid = &v
	return s
}

func (s *GetTaskInstanceRelationRequest) Validate() error {
	return dara.Validate(s)
}
