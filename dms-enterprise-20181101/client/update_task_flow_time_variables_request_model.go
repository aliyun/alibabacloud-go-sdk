// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTaskFlowTimeVariablesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDagId(v int64) *UpdateTaskFlowTimeVariablesRequest
	GetDagId() *int64
	SetTid(v int64) *UpdateTaskFlowTimeVariablesRequest
	GetTid() *int64
	SetTimeVariables(v string) *UpdateTaskFlowTimeVariablesRequest
	GetTimeVariables() *string
}

type UpdateTaskFlowTimeVariablesRequest struct {
	// The ID of the task node. You can call the [GetTaskInstanceRelation](https://help.aliyun.com/document_detail/424711.html) operation to query the node ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 43****
	DagId *int64 `json:"DagId,omitempty" xml:"DagId,omitempty"`
	// The ID of the tenant.
	//
	// > :To view the ID of the tenant, go to the Data Management (DMS) console and move the pointer over the profile picture in the upper-right corner. For more information, see [View information about the current tenant](https://help.aliyun.com/document_detail/181330.html).
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
	// The time variables for the task flow.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"variables":[ {"name":"var", "pattern":"yyyy-MM-dd|+0m+0h-2d+0w+0M+1y"} ]}
	TimeVariables *string `json:"TimeVariables,omitempty" xml:"TimeVariables,omitempty"`
}

func (s UpdateTaskFlowTimeVariablesRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTaskFlowTimeVariablesRequest) GoString() string {
	return s.String()
}

func (s *UpdateTaskFlowTimeVariablesRequest) GetDagId() *int64 {
	return s.DagId
}

func (s *UpdateTaskFlowTimeVariablesRequest) GetTid() *int64 {
	return s.Tid
}

func (s *UpdateTaskFlowTimeVariablesRequest) GetTimeVariables() *string {
	return s.TimeVariables
}

func (s *UpdateTaskFlowTimeVariablesRequest) SetDagId(v int64) *UpdateTaskFlowTimeVariablesRequest {
	s.DagId = &v
	return s
}

func (s *UpdateTaskFlowTimeVariablesRequest) SetTid(v int64) *UpdateTaskFlowTimeVariablesRequest {
	s.Tid = &v
	return s
}

func (s *UpdateTaskFlowTimeVariablesRequest) SetTimeVariables(v string) *UpdateTaskFlowTimeVariablesRequest {
	s.TimeVariables = &v
	return s
}

func (s *UpdateTaskFlowTimeVariablesRequest) Validate() error {
	return dara.Validate(s)
}
