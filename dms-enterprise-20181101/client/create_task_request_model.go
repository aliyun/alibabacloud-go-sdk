// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDagId(v int64) *CreateTaskRequest
	GetDagId() *int64
	SetGraphParam(v string) *CreateTaskRequest
	GetGraphParam() *string
	SetNodeContent(v string) *CreateTaskRequest
	GetNodeContent() *string
	SetNodeName(v string) *CreateTaskRequest
	GetNodeName() *string
	SetNodeOutput(v string) *CreateTaskRequest
	GetNodeOutput() *string
	SetNodeType(v string) *CreateTaskRequest
	GetNodeType() *string
	SetTid(v int64) *CreateTaskRequest
	GetTid() *int64
	SetTimeVariables(v string) *CreateTaskRequest
	GetTimeVariables() *string
}

type CreateTaskRequest struct {
	// The ID of the task flow. You can call the [ListTaskFlow](https://help.aliyun.com/document_detail/424565.html) or [ListLhTaskFlowAndScenario](https://help.aliyun.com/document_detail/426672.html) operation to query the task flow ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7***
	DagId *int64 `json:"DagId,omitempty" xml:"DagId,omitempty"`
	// The position of the node on the Directed Acyclic Graph (DAG).
	//
	// example:
	//
	// test
	GraphParam *string `json:"GraphParam,omitempty" xml:"GraphParam,omitempty"`
	// The configuration of the node.
	//
	// example:
	//
	// test
	NodeContent *string `json:"NodeContent,omitempty" xml:"NodeContent,omitempty"`
	// The name of the node that you want to create.
	//
	// This parameter is required.
	//
	// example:
	//
	// zhttest
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// The output variables configured for the task.
	//
	// example:
	//
	// test
	NodeOutput *string `json:"NodeOutput,omitempty" xml:"NodeOutput,omitempty"`
	// The type of the node that you want to create. For more information about the valid values for this parameter, see [NodeType parameter](https://help.aliyun.com/document_detail/424705.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 36
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	// The tenant ID.
	//
	// >  To view the tenant ID, move the pointer over the profile picture in the upper-right corner of the Data Management (DMS) console. For more information, see the "View information about the current tenant" section of the [Manage DMS tenants](https://help.aliyun.com/document_detail/181330.html) topic.
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
	// The time variables configured for the node.
	//
	// example:
	//
	// test
	TimeVariables *string `json:"TimeVariables,omitempty" xml:"TimeVariables,omitempty"`
}

func (s CreateTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateTaskRequest) GetDagId() *int64 {
	return s.DagId
}

func (s *CreateTaskRequest) GetGraphParam() *string {
	return s.GraphParam
}

func (s *CreateTaskRequest) GetNodeContent() *string {
	return s.NodeContent
}

func (s *CreateTaskRequest) GetNodeName() *string {
	return s.NodeName
}

func (s *CreateTaskRequest) GetNodeOutput() *string {
	return s.NodeOutput
}

func (s *CreateTaskRequest) GetNodeType() *string {
	return s.NodeType
}

func (s *CreateTaskRequest) GetTid() *int64 {
	return s.Tid
}

func (s *CreateTaskRequest) GetTimeVariables() *string {
	return s.TimeVariables
}

func (s *CreateTaskRequest) SetDagId(v int64) *CreateTaskRequest {
	s.DagId = &v
	return s
}

func (s *CreateTaskRequest) SetGraphParam(v string) *CreateTaskRequest {
	s.GraphParam = &v
	return s
}

func (s *CreateTaskRequest) SetNodeContent(v string) *CreateTaskRequest {
	s.NodeContent = &v
	return s
}

func (s *CreateTaskRequest) SetNodeName(v string) *CreateTaskRequest {
	s.NodeName = &v
	return s
}

func (s *CreateTaskRequest) SetNodeOutput(v string) *CreateTaskRequest {
	s.NodeOutput = &v
	return s
}

func (s *CreateTaskRequest) SetNodeType(v string) *CreateTaskRequest {
	s.NodeType = &v
	return s
}

func (s *CreateTaskRequest) SetTid(v int64) *CreateTaskRequest {
	s.Tid = &v
	return s
}

func (s *CreateTaskRequest) SetTimeVariables(v string) *CreateTaskRequest {
	s.TimeVariables = &v
	return s
}

func (s *CreateTaskRequest) Validate() error {
	return dara.Validate(s)
}
