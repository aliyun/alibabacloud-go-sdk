// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTasksInTaskFlowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ListTasksInTaskFlowResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListTasksInTaskFlowResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ListTasksInTaskFlowResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListTasksInTaskFlowResponseBody
	GetSuccess() *bool
	SetTasks(v *ListTasksInTaskFlowResponseBodyTasks) *ListTasksInTaskFlowResponseBody
	GetTasks() *ListTasksInTaskFlowResponseBodyTasks
}

type ListTasksInTaskFlowResponseBody struct {
	// The error code returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the request. You can use the ID to locate logs and troubleshoot issues.
	//
	// example:
	//
	// 8B36B063-6B7D-5595-9FCF-3844B7B7ACD4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The tasks in the task flow.
	Tasks *ListTasksInTaskFlowResponseBodyTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Struct"`
}

func (s ListTasksInTaskFlowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTasksInTaskFlowResponseBody) GoString() string {
	return s.String()
}

func (s *ListTasksInTaskFlowResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListTasksInTaskFlowResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListTasksInTaskFlowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTasksInTaskFlowResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListTasksInTaskFlowResponseBody) GetTasks() *ListTasksInTaskFlowResponseBodyTasks {
	return s.Tasks
}

func (s *ListTasksInTaskFlowResponseBody) SetErrorCode(v string) *ListTasksInTaskFlowResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListTasksInTaskFlowResponseBody) SetErrorMessage(v string) *ListTasksInTaskFlowResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListTasksInTaskFlowResponseBody) SetRequestId(v string) *ListTasksInTaskFlowResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTasksInTaskFlowResponseBody) SetSuccess(v bool) *ListTasksInTaskFlowResponseBody {
	s.Success = &v
	return s
}

func (s *ListTasksInTaskFlowResponseBody) SetTasks(v *ListTasksInTaskFlowResponseBodyTasks) *ListTasksInTaskFlowResponseBody {
	s.Tasks = v
	return s
}

func (s *ListTasksInTaskFlowResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListTasksInTaskFlowResponseBodyTasks struct {
	Task []*ListTasksInTaskFlowResponseBodyTasksTask `json:"Task,omitempty" xml:"Task,omitempty" type:"Repeated"`
}

func (s ListTasksInTaskFlowResponseBodyTasks) String() string {
	return dara.Prettify(s)
}

func (s ListTasksInTaskFlowResponseBodyTasks) GoString() string {
	return s.String()
}

func (s *ListTasksInTaskFlowResponseBodyTasks) GetTask() []*ListTasksInTaskFlowResponseBodyTasksTask {
	return s.Task
}

func (s *ListTasksInTaskFlowResponseBodyTasks) SetTask(v []*ListTasksInTaskFlowResponseBodyTasksTask) *ListTasksInTaskFlowResponseBodyTasks {
	s.Task = v
	return s
}

func (s *ListTasksInTaskFlowResponseBodyTasks) Validate() error {
	return dara.Validate(s)
}

type ListTasksInTaskFlowResponseBodyTasksTask struct {
	// The position of the node on the Directed Acyclic Graph (DAG).
	//
	// example:
	//
	// {\\"x\\":435,\\"y\\":192,\\"layoutType\\":\\"Horizontal\\"}
	GraphParam *string `json:"GraphParam,omitempty" xml:"GraphParam,omitempty"`
	// The advanced configuration for the node.
	//
	// example:
	//
	// {     "rerun":{    "rerunEnable":true,      "rerunCount":1,   "rerunInterval":10 //  }}
	NodeConfig *string `json:"NodeConfig,omitempty" xml:"NodeConfig,omitempty"`
	// The configuration for the node.
	//
	// example:
	//
	// {     "dbId":123***,  "sql":"",    "dbType":"polardb" }
	NodeContent *string `json:"NodeContent,omitempty" xml:"NodeContent,omitempty"`
	// The ID of the node.
	//
	// example:
	//
	// 92***
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The name of the node.
	//
	// example:
	//
	// Cross-Database Spark SQL-1
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// The output variables for the task.
	//
	// example:
	//
	// {  "outputs":[{ { "row":0, "column":-1,                  "combiner":"," }            "extractMethod":"{\\"row\\":0,\\"column\\":-1,\\"combiner\\":\\",\\"}",         "variableName":"var",    "description":"For demo"} ] }
	NodeOutput *string `json:"NodeOutput,omitempty" xml:"NodeOutput,omitempty"`
	// The type of the node. For more information about the valid values for this parameter, see [NodeType parameter](https://help.aliyun.com/document_detail/424705.html).
	//
	// example:
	//
	// SPARK_SQL
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	// The time variables configured for the node.
	//
	// example:
	//
	// {\\"variables\\":[{\\"name\\":\\"test1\\",\\"pattern\\":\\"yyyy-MM-dd
	TimeVariables *string `json:"TimeVariables,omitempty" xml:"TimeVariables,omitempty"`
}

func (s ListTasksInTaskFlowResponseBodyTasksTask) String() string {
	return dara.Prettify(s)
}

func (s ListTasksInTaskFlowResponseBodyTasksTask) GoString() string {
	return s.String()
}

func (s *ListTasksInTaskFlowResponseBodyTasksTask) GetGraphParam() *string {
	return s.GraphParam
}

func (s *ListTasksInTaskFlowResponseBodyTasksTask) GetNodeConfig() *string {
	return s.NodeConfig
}

func (s *ListTasksInTaskFlowResponseBodyTasksTask) GetNodeContent() *string {
	return s.NodeContent
}

func (s *ListTasksInTaskFlowResponseBodyTasksTask) GetNodeId() *string {
	return s.NodeId
}

func (s *ListTasksInTaskFlowResponseBodyTasksTask) GetNodeName() *string {
	return s.NodeName
}

func (s *ListTasksInTaskFlowResponseBodyTasksTask) GetNodeOutput() *string {
	return s.NodeOutput
}

func (s *ListTasksInTaskFlowResponseBodyTasksTask) GetNodeType() *string {
	return s.NodeType
}

func (s *ListTasksInTaskFlowResponseBodyTasksTask) GetTimeVariables() *string {
	return s.TimeVariables
}

func (s *ListTasksInTaskFlowResponseBodyTasksTask) SetGraphParam(v string) *ListTasksInTaskFlowResponseBodyTasksTask {
	s.GraphParam = &v
	return s
}

func (s *ListTasksInTaskFlowResponseBodyTasksTask) SetNodeConfig(v string) *ListTasksInTaskFlowResponseBodyTasksTask {
	s.NodeConfig = &v
	return s
}

func (s *ListTasksInTaskFlowResponseBodyTasksTask) SetNodeContent(v string) *ListTasksInTaskFlowResponseBodyTasksTask {
	s.NodeContent = &v
	return s
}

func (s *ListTasksInTaskFlowResponseBodyTasksTask) SetNodeId(v string) *ListTasksInTaskFlowResponseBodyTasksTask {
	s.NodeId = &v
	return s
}

func (s *ListTasksInTaskFlowResponseBodyTasksTask) SetNodeName(v string) *ListTasksInTaskFlowResponseBodyTasksTask {
	s.NodeName = &v
	return s
}

func (s *ListTasksInTaskFlowResponseBodyTasksTask) SetNodeOutput(v string) *ListTasksInTaskFlowResponseBodyTasksTask {
	s.NodeOutput = &v
	return s
}

func (s *ListTasksInTaskFlowResponseBodyTasksTask) SetNodeType(v string) *ListTasksInTaskFlowResponseBodyTasksTask {
	s.NodeType = &v
	return s
}

func (s *ListTasksInTaskFlowResponseBodyTasksTask) SetTimeVariables(v string) *ListTasksInTaskFlowResponseBodyTasksTask {
	s.TimeVariables = &v
	return s
}

func (s *ListTasksInTaskFlowResponseBodyTasksTask) Validate() error {
	return dara.Validate(s)
}
