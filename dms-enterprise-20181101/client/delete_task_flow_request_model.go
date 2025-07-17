// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTaskFlowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDagId(v int64) *DeleteTaskFlowRequest
	GetDagId() *int64
	SetTid(v int64) *DeleteTaskFlowRequest
	GetTid() *int64
}

type DeleteTaskFlowRequest struct {
	// The ID of the task flow. You can call the [ListTaskFlow](https://help.aliyun.com/document_detail/424565.html) or [ListLhTaskFlowAndScenario](https://help.aliyun.com/document_detail/426672.html) operation to obtain the ID of the task flow.
	//
	// This parameter is required.
	//
	// example:
	//
	// 134137***
	DagId *int64 `json:"DagId,omitempty" xml:"DagId,omitempty"`
	// The ID of the tenant. You can call the [GetUserActiveTenant](https://help.aliyun.com/document_detail/198073.html) or [ListUserTenants](https://help.aliyun.com/document_detail/198074.html) operation to obtain the tenant ID.
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s DeleteTaskFlowRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteTaskFlowRequest) GoString() string {
	return s.String()
}

func (s *DeleteTaskFlowRequest) GetDagId() *int64 {
	return s.DagId
}

func (s *DeleteTaskFlowRequest) GetTid() *int64 {
	return s.Tid
}

func (s *DeleteTaskFlowRequest) SetDagId(v int64) *DeleteTaskFlowRequest {
	s.DagId = &v
	return s
}

func (s *DeleteTaskFlowRequest) SetTid(v int64) *DeleteTaskFlowRequest {
	s.Tid = &v
	return s
}

func (s *DeleteTaskFlowRequest) Validate() error {
	return dara.Validate(s)
}
