// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *GetTaskResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetTaskResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *GetTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetTaskResponseBody
	GetSuccess() *bool
	SetTask(v *GetTaskResponseBodyTask) *GetTaskResponseBody
	GetTask() *GetTaskResponseBodyTask
}

type GetTaskResponseBody struct {
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
	// The ID of the request. You can use the ID to query logs and troubleshoot issues.
	//
	// example:
	//
	// 7838266C-E17B-58F4-B072-4DC356B58258
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
	// The task node.
	Task *GetTaskResponseBodyTask `json:"Task,omitempty" xml:"Task,omitempty" type:"Struct"`
}

func (s GetTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetTaskResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetTaskResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetTaskResponseBody) GetTask() *GetTaskResponseBodyTask {
	return s.Task
}

func (s *GetTaskResponseBody) SetErrorCode(v string) *GetTaskResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetTaskResponseBody) SetErrorMessage(v string) *GetTaskResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetTaskResponseBody) SetRequestId(v string) *GetTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTaskResponseBody) SetSuccess(v bool) *GetTaskResponseBody {
	s.Success = &v
	return s
}

func (s *GetTaskResponseBody) SetTask(v *GetTaskResponseBodyTask) *GetTaskResponseBody {
	s.Task = v
	return s
}

func (s *GetTaskResponseBody) Validate() error {
	if s.Task != nil {
		if err := s.Task.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTaskResponseBodyTask struct {
	// The ID of the task flow to which the node belongs.
	//
	// example:
	//
	// 7321
	DagId *int64 `json:"DagId,omitempty" xml:"DagId,omitempty"`
	// The position of the node on the Directed Acyclic Graph (DAG).
	//
	// example:
	//
	// {"{\\"x\\":0,\\"y\\":0,\\"layoutType\\":\\"Horizontal\\"}",  "id": 51***}
	GraphParam *string `json:"GraphParam,omitempty" xml:"GraphParam,omitempty"`
	// The advanced configuration for the node.
	NodeConfig *string `json:"NodeConfig,omitempty" xml:"NodeConfig,omitempty"`
	// The configuration for the node.
	//
	// example:
	//
	// {\\"dbList\\":[{\\"instanceId\\":177****}"   }
	NodeContent *string `json:"NodeContent,omitempty" xml:"NodeContent,omitempty"`
	// The name of the node.
	//
	// example:
	//
	// Cross-database Spark SQL-1
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// The output variables for the node. This parameter is available only for some types of nodes.
	//
	// example:
	//
	// { "outputs":[ "extractMethod":"json" , "variableName":"var",   "description":"demo desc" } ] }
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
	// {\\"variables\\":[{\\"name\\":\\"Today\\",\\"pattern\\":\\"yyyy-MM-dd|+1d\\"}]}
	TimeVariables *string `json:"TimeVariables,omitempty" xml:"TimeVariables,omitempty"`
}

func (s GetTaskResponseBodyTask) String() string {
	return dara.Prettify(s)
}

func (s GetTaskResponseBodyTask) GoString() string {
	return s.String()
}

func (s *GetTaskResponseBodyTask) GetDagId() *int64 {
	return s.DagId
}

func (s *GetTaskResponseBodyTask) GetGraphParam() *string {
	return s.GraphParam
}

func (s *GetTaskResponseBodyTask) GetNodeConfig() *string {
	return s.NodeConfig
}

func (s *GetTaskResponseBodyTask) GetNodeContent() *string {
	return s.NodeContent
}

func (s *GetTaskResponseBodyTask) GetNodeName() *string {
	return s.NodeName
}

func (s *GetTaskResponseBodyTask) GetNodeOutput() *string {
	return s.NodeOutput
}

func (s *GetTaskResponseBodyTask) GetNodeType() *string {
	return s.NodeType
}

func (s *GetTaskResponseBodyTask) GetTimeVariables() *string {
	return s.TimeVariables
}

func (s *GetTaskResponseBodyTask) SetDagId(v int64) *GetTaskResponseBodyTask {
	s.DagId = &v
	return s
}

func (s *GetTaskResponseBodyTask) SetGraphParam(v string) *GetTaskResponseBodyTask {
	s.GraphParam = &v
	return s
}

func (s *GetTaskResponseBodyTask) SetNodeConfig(v string) *GetTaskResponseBodyTask {
	s.NodeConfig = &v
	return s
}

func (s *GetTaskResponseBodyTask) SetNodeContent(v string) *GetTaskResponseBodyTask {
	s.NodeContent = &v
	return s
}

func (s *GetTaskResponseBodyTask) SetNodeName(v string) *GetTaskResponseBodyTask {
	s.NodeName = &v
	return s
}

func (s *GetTaskResponseBodyTask) SetNodeOutput(v string) *GetTaskResponseBodyTask {
	s.NodeOutput = &v
	return s
}

func (s *GetTaskResponseBodyTask) SetNodeType(v string) *GetTaskResponseBodyTask {
	s.NodeType = &v
	return s
}

func (s *GetTaskResponseBodyTask) SetTimeVariables(v string) *GetTaskResponseBodyTask {
	s.TimeVariables = &v
	return s
}

func (s *GetTaskResponseBodyTask) Validate() error {
	return dara.Validate(s)
}
