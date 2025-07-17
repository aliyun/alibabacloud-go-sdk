// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTaskTimeVariablesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNodeId(v string) *UpdateTaskTimeVariablesRequest
	GetNodeId() *string
	SetTid(v int64) *UpdateTaskTimeVariablesRequest
	GetTid() *int64
	SetTimeVariables(v string) *UpdateTaskTimeVariablesRequest
	GetTimeVariables() *string
}

type UpdateTaskTimeVariablesRequest struct {
	// The ID of the task node. You can call the [GetTaskInstanceRelation](https://help.aliyun.com/document_detail/424711.html) operation to query the node ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 43****
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The ID of the tenant.
	//
	// > :To view the ID of the tenant, go to the Data Management (DMS) console and move the pointer over the profile picture in the upper-right corner. For more information, see [View information about the current tenant](https://help.aliyun.com/document_detail/181330.html).
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
	// The time variables configured for the node. The value of this parameter must be a JSON string.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"variables":[ {"name":"var", "pattern":"yyyy-MM-dd|+0m+0h-2d+0w+0M+1y"} ]}
	TimeVariables *string `json:"TimeVariables,omitempty" xml:"TimeVariables,omitempty"`
}

func (s UpdateTaskTimeVariablesRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTaskTimeVariablesRequest) GoString() string {
	return s.String()
}

func (s *UpdateTaskTimeVariablesRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *UpdateTaskTimeVariablesRequest) GetTid() *int64 {
	return s.Tid
}

func (s *UpdateTaskTimeVariablesRequest) GetTimeVariables() *string {
	return s.TimeVariables
}

func (s *UpdateTaskTimeVariablesRequest) SetNodeId(v string) *UpdateTaskTimeVariablesRequest {
	s.NodeId = &v
	return s
}

func (s *UpdateTaskTimeVariablesRequest) SetTid(v int64) *UpdateTaskTimeVariablesRequest {
	s.Tid = &v
	return s
}

func (s *UpdateTaskTimeVariablesRequest) SetTimeVariables(v string) *UpdateTaskTimeVariablesRequest {
	s.TimeVariables = &v
	return s
}

func (s *UpdateTaskTimeVariablesRequest) Validate() error {
	return dara.Validate(s)
}
