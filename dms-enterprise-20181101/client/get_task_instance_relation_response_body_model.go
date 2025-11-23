// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTaskInstanceRelationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *GetTaskInstanceRelationResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetTaskInstanceRelationResponseBody
	GetErrorMessage() *string
	SetNodeList(v *GetTaskInstanceRelationResponseBodyNodeList) *GetTaskInstanceRelationResponseBody
	GetNodeList() *GetTaskInstanceRelationResponseBodyNodeList
	SetRequestId(v string) *GetTaskInstanceRelationResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetTaskInstanceRelationResponseBody
	GetSuccess() *bool
}

type GetTaskInstanceRelationResponseBody struct {
	// The error code returned if the request fails.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the request fails.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The information about the nodes in the execution record of the task flow.
	NodeList *GetTaskInstanceRelationResponseBodyNodeList `json:"NodeList,omitempty" xml:"NodeList,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 028BF827-3801-5869-8548-F4A039256304
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- **true**: The request is successful.
	//
	// 	- **false**: The request fails.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetTaskInstanceRelationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTaskInstanceRelationResponseBody) GoString() string {
	return s.String()
}

func (s *GetTaskInstanceRelationResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetTaskInstanceRelationResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetTaskInstanceRelationResponseBody) GetNodeList() *GetTaskInstanceRelationResponseBodyNodeList {
	return s.NodeList
}

func (s *GetTaskInstanceRelationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTaskInstanceRelationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetTaskInstanceRelationResponseBody) SetErrorCode(v string) *GetTaskInstanceRelationResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetTaskInstanceRelationResponseBody) SetErrorMessage(v string) *GetTaskInstanceRelationResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetTaskInstanceRelationResponseBody) SetNodeList(v *GetTaskInstanceRelationResponseBodyNodeList) *GetTaskInstanceRelationResponseBody {
	s.NodeList = v
	return s
}

func (s *GetTaskInstanceRelationResponseBody) SetRequestId(v string) *GetTaskInstanceRelationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTaskInstanceRelationResponseBody) SetSuccess(v bool) *GetTaskInstanceRelationResponseBody {
	s.Success = &v
	return s
}

func (s *GetTaskInstanceRelationResponseBody) Validate() error {
	if s.NodeList != nil {
		if err := s.NodeList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTaskInstanceRelationResponseBodyNodeList struct {
	Node []*GetTaskInstanceRelationResponseBodyNodeListNode `json:"Node,omitempty" xml:"Node,omitempty" type:"Repeated"`
}

func (s GetTaskInstanceRelationResponseBodyNodeList) String() string {
	return dara.Prettify(s)
}

func (s GetTaskInstanceRelationResponseBodyNodeList) GoString() string {
	return s.String()
}

func (s *GetTaskInstanceRelationResponseBodyNodeList) GetNode() []*GetTaskInstanceRelationResponseBodyNodeListNode {
	return s.Node
}

func (s *GetTaskInstanceRelationResponseBodyNodeList) SetNode(v []*GetTaskInstanceRelationResponseBodyNodeListNode) *GetTaskInstanceRelationResponseBodyNodeList {
	s.Node = v
	return s
}

func (s *GetTaskInstanceRelationResponseBodyNodeList) Validate() error {
	if s.Node != nil {
		for _, item := range s.Node {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetTaskInstanceRelationResponseBodyNodeListNode struct {
	// The business time of the node.
	//
	// example:
	//
	// 2021-11-09 14:37:26
	BusinessTime *string `json:"BusinessTime,omitempty" xml:"BusinessTime,omitempty"`
	// The time when the execution of the task flow was complete. The time is displayed in the yyyy-MM-DD HH:mm:ss format.
	//
	// example:
	//
	// 2021-11-11 14:38:57
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The amount of time consumed for running the node. Unit: milliseconds.
	//
	// example:
	//
	// 170655
	ExecuteTime *int64 `json:"ExecuteTime,omitempty" xml:"ExecuteTime,omitempty"`
	// The ID of the execution record of the task flow.
	//
	// example:
	//
	// 14059
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The description of the task.
	//
	// example:
	//
	// test
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the node.
	//
	// example:
	//
	// 14059
	NodeId *int64 `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The name of the node.
	//
	// example:
	//
	// Spark SQL-1
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// The type of the node. For more information about the valid values for this parameter, see [NodeType parameter](https://help.aliyun.com/document_detail/424705.html).
	//
	// example:
	//
	// 36
	NodeType *int32 `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	// The status of the node. Valid values:
	//
	// 	- **0**: The node is waiting to be scheduled.
	//
	// 	- **1**: The node is running.
	//
	// 	- **2**: The node is suspended.
	//
	// 	- **3**: The node failed to run.
	//
	// 	- **4**: The node is run.
	//
	// 	- **5**: The node is complete.
	//
	// example:
	//
	// 4
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetTaskInstanceRelationResponseBodyNodeListNode) String() string {
	return dara.Prettify(s)
}

func (s GetTaskInstanceRelationResponseBodyNodeListNode) GoString() string {
	return s.String()
}

func (s *GetTaskInstanceRelationResponseBodyNodeListNode) GetBusinessTime() *string {
	return s.BusinessTime
}

func (s *GetTaskInstanceRelationResponseBodyNodeListNode) GetEndTime() *string {
	return s.EndTime
}

func (s *GetTaskInstanceRelationResponseBodyNodeListNode) GetExecuteTime() *int64 {
	return s.ExecuteTime
}

func (s *GetTaskInstanceRelationResponseBodyNodeListNode) GetId() *int64 {
	return s.Id
}

func (s *GetTaskInstanceRelationResponseBodyNodeListNode) GetMessage() *string {
	return s.Message
}

func (s *GetTaskInstanceRelationResponseBodyNodeListNode) GetNodeId() *int64 {
	return s.NodeId
}

func (s *GetTaskInstanceRelationResponseBodyNodeListNode) GetNodeName() *string {
	return s.NodeName
}

func (s *GetTaskInstanceRelationResponseBodyNodeListNode) GetNodeType() *int32 {
	return s.NodeType
}

func (s *GetTaskInstanceRelationResponseBodyNodeListNode) GetStatus() *int32 {
	return s.Status
}

func (s *GetTaskInstanceRelationResponseBodyNodeListNode) SetBusinessTime(v string) *GetTaskInstanceRelationResponseBodyNodeListNode {
	s.BusinessTime = &v
	return s
}

func (s *GetTaskInstanceRelationResponseBodyNodeListNode) SetEndTime(v string) *GetTaskInstanceRelationResponseBodyNodeListNode {
	s.EndTime = &v
	return s
}

func (s *GetTaskInstanceRelationResponseBodyNodeListNode) SetExecuteTime(v int64) *GetTaskInstanceRelationResponseBodyNodeListNode {
	s.ExecuteTime = &v
	return s
}

func (s *GetTaskInstanceRelationResponseBodyNodeListNode) SetId(v int64) *GetTaskInstanceRelationResponseBodyNodeListNode {
	s.Id = &v
	return s
}

func (s *GetTaskInstanceRelationResponseBodyNodeListNode) SetMessage(v string) *GetTaskInstanceRelationResponseBodyNodeListNode {
	s.Message = &v
	return s
}

func (s *GetTaskInstanceRelationResponseBodyNodeListNode) SetNodeId(v int64) *GetTaskInstanceRelationResponseBodyNodeListNode {
	s.NodeId = &v
	return s
}

func (s *GetTaskInstanceRelationResponseBodyNodeListNode) SetNodeName(v string) *GetTaskInstanceRelationResponseBodyNodeListNode {
	s.NodeName = &v
	return s
}

func (s *GetTaskInstanceRelationResponseBodyNodeListNode) SetNodeType(v int32) *GetTaskInstanceRelationResponseBodyNodeListNode {
	s.NodeType = &v
	return s
}

func (s *GetTaskInstanceRelationResponseBodyNodeListNode) SetStatus(v int32) *GetTaskInstanceRelationResponseBodyNodeListNode {
	s.Status = &v
	return s
}

func (s *GetTaskInstanceRelationResponseBodyNodeListNode) Validate() error {
	return dara.Validate(s)
}
