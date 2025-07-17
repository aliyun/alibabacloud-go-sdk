// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopTaskFlowInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDagId(v int64) *StopTaskFlowInstanceRequest
	GetDagId() *int64
	SetDagInstanceId(v int64) *StopTaskFlowInstanceRequest
	GetDagInstanceId() *int64
	SetTid(v int64) *StopTaskFlowInstanceRequest
	GetTid() *int64
}

type StopTaskFlowInstanceRequest struct {
	// The ID of the task flow. You can call the [ListTaskFlow](https://help.aliyun.com/document_detail/424565.html) or [ListLhTaskFlowAndScenario](https://help.aliyun.com/document_detail/426672.html) operation to query the task flow ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7***
	DagId *int64 `json:"DagId,omitempty" xml:"DagId,omitempty"`
	// The ID of the execution record of the task flow. You can call the [ListTaskFlowInstance](https://help.aliyun.com/document_detail/424689.html) operation to query the execution record ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 47****
	DagInstanceId *int64 `json:"DagInstanceId,omitempty" xml:"DagInstanceId,omitempty"`
	// The ID of the tenant.
	//
	// > To view the ID of the tenant, go to the Data Management (DMS) console and move the pointer over the profile picture in the upper-right corner. For more information, see [View information about the current tenant](https://help.aliyun.com/document_detail/181330.html).
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s StopTaskFlowInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s StopTaskFlowInstanceRequest) GoString() string {
	return s.String()
}

func (s *StopTaskFlowInstanceRequest) GetDagId() *int64 {
	return s.DagId
}

func (s *StopTaskFlowInstanceRequest) GetDagInstanceId() *int64 {
	return s.DagInstanceId
}

func (s *StopTaskFlowInstanceRequest) GetTid() *int64 {
	return s.Tid
}

func (s *StopTaskFlowInstanceRequest) SetDagId(v int64) *StopTaskFlowInstanceRequest {
	s.DagId = &v
	return s
}

func (s *StopTaskFlowInstanceRequest) SetDagInstanceId(v int64) *StopTaskFlowInstanceRequest {
	s.DagInstanceId = &v
	return s
}

func (s *StopTaskFlowInstanceRequest) SetTid(v int64) *StopTaskFlowInstanceRequest {
	s.Tid = &v
	return s
}

func (s *StopTaskFlowInstanceRequest) Validate() error {
	return dara.Validate(s)
}
