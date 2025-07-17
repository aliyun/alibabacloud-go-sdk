// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTasksInTaskFlowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDagId(v int64) *ListTasksInTaskFlowRequest
	GetDagId() *int64
	SetTid(v int64) *ListTasksInTaskFlowRequest
	GetTid() *int64
}

type ListTasksInTaskFlowRequest struct {
	// The ID of the task flow. You can call the [ListTaskFlow](https://help.aliyun.com/document_detail/424565.html) or [ListLhTaskFlowAndScenario](https://help.aliyun.com/document_detail/426672.html) operation to query the task flow ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 32***
	DagId *int64 `json:"DagId,omitempty" xml:"DagId,omitempty"`
	// The ID of the tenant.
	//
	// >  To view the ID of the tenant, move the pointer over the profile picture in the upper-right corner of the Data Management (DMS) console. For more information, see [View information about the current tenant](https://help.aliyun.com/document_detail/181330.html).
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s ListTasksInTaskFlowRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTasksInTaskFlowRequest) GoString() string {
	return s.String()
}

func (s *ListTasksInTaskFlowRequest) GetDagId() *int64 {
	return s.DagId
}

func (s *ListTasksInTaskFlowRequest) GetTid() *int64 {
	return s.Tid
}

func (s *ListTasksInTaskFlowRequest) SetDagId(v int64) *ListTasksInTaskFlowRequest {
	s.DagId = &v
	return s
}

func (s *ListTasksInTaskFlowRequest) SetTid(v int64) *ListTasksInTaskFlowRequest {
	s.Tid = &v
	return s
}

func (s *ListTasksInTaskFlowRequest) Validate() error {
	return dara.Validate(s)
}
