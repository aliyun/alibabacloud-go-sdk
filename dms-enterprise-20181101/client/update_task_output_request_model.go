// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTaskOutputRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNodeId(v string) *UpdateTaskOutputRequest
	GetNodeId() *string
	SetNodeOutput(v string) *UpdateTaskOutputRequest
	GetNodeOutput() *string
	SetTid(v int64) *UpdateTaskOutputRequest
	GetTid() *int64
}

type UpdateTaskOutputRequest struct {
	// The ID of the node. You can call the [GetTaskInstanceRelation](https://help.aliyun.com/document_detail/424711.html) operation to query the node ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 14059
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The output variables for the task.
	//
	// example:
	//
	// {“outputs”:[{"row":0, "column":-1,"combiner":","}"extractMethod":"{\\"row\\":0,\\"column\\":-1,\\"combiner\\":\\",\\"}",   "variableName":"var", "description":"For demo" }]}
	NodeOutput *string `json:"NodeOutput,omitempty" xml:"NodeOutput,omitempty"`
	// The ID of the tenant.
	//
	// >  To view the ID of the tenant, go to the Data Management (DMS) console and move the pointer over the profile picture in the upper-right corner. For more information, see [View information about the current tenant](https://help.aliyun.com/document_detail/181330.html).
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s UpdateTaskOutputRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTaskOutputRequest) GoString() string {
	return s.String()
}

func (s *UpdateTaskOutputRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *UpdateTaskOutputRequest) GetNodeOutput() *string {
	return s.NodeOutput
}

func (s *UpdateTaskOutputRequest) GetTid() *int64 {
	return s.Tid
}

func (s *UpdateTaskOutputRequest) SetNodeId(v string) *UpdateTaskOutputRequest {
	s.NodeId = &v
	return s
}

func (s *UpdateTaskOutputRequest) SetNodeOutput(v string) *UpdateTaskOutputRequest {
	s.NodeOutput = &v
	return s
}

func (s *UpdateTaskOutputRequest) SetTid(v int64) *UpdateTaskOutputRequest {
	s.Tid = &v
	return s
}

func (s *UpdateTaskOutputRequest) Validate() error {
	return dara.Validate(s)
}
