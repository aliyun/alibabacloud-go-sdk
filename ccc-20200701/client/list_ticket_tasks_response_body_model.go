// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTicketTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListTicketTasksResponseBody
	GetCode() *string
	SetData(v []*ListTicketTasksResponseBodyData) *ListTicketTasksResponseBody
	GetData() []*ListTicketTasksResponseBodyData
	SetHttpStatusCode(v int32) *ListTicketTasksResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListTicketTasksResponseBody
	GetMessage() *string
	SetParams(v []*string) *ListTicketTasksResponseBody
	GetParams() []*string
	SetRequestId(v string) *ListTicketTasksResponseBody
	GetRequestId() *string
}

type ListTicketTasksResponseBody struct {
	// The response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data []*ListTicketTasksResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The response message.
	//
	// example:
	//
	// 无
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// A list of error parameters.
	Params []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// DE803553-8AA9-4B9D-9E4E-A82BC69EDCEE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListTicketTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTicketTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListTicketTasksResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListTicketTasksResponseBody) GetData() []*ListTicketTasksResponseBodyData {
	return s.Data
}

func (s *ListTicketTasksResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListTicketTasksResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListTicketTasksResponseBody) GetParams() []*string {
	return s.Params
}

func (s *ListTicketTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTicketTasksResponseBody) SetCode(v string) *ListTicketTasksResponseBody {
	s.Code = &v
	return s
}

func (s *ListTicketTasksResponseBody) SetData(v []*ListTicketTasksResponseBodyData) *ListTicketTasksResponseBody {
	s.Data = v
	return s
}

func (s *ListTicketTasksResponseBody) SetHttpStatusCode(v int32) *ListTicketTasksResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListTicketTasksResponseBody) SetMessage(v string) *ListTicketTasksResponseBody {
	s.Message = &v
	return s
}

func (s *ListTicketTasksResponseBody) SetParams(v []*string) *ListTicketTasksResponseBody {
	s.Params = v
	return s
}

func (s *ListTicketTasksResponseBody) SetRequestId(v string) *ListTicketTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTicketTasksResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTicketTasksResponseBodyData struct {
	// The action performed on the node.
	//
	// example:
	//
	// CounterSignTask
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// The ID of the assignee.
	//
	// example:
	//
	// assignee@ccc-test
	Assignee *string `json:"Assignee,omitempty" xml:"Assignee,omitempty"`
	// The name of the assignee.
	//
	// example:
	//
	// assignee
	AssigneeName *string `json:"AssigneeName,omitempty" xml:"AssigneeName,omitempty"`
	// The comment.
	//
	// example:
	//
	// 同意
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The end time.
	//
	// example:
	//
	// 1620269200000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// A list of files.
	FileKeys []*string `json:"FileKeys,omitempty" xml:"FileKeys,omitempty" type:"Repeated"`
	// A list of file access URLs.
	FileUrls []*string `json:"FileUrls,omitempty" xml:"FileUrls,omitempty" type:"Repeated"`
	// The ID of the instance.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The start time.
	//
	// example:
	//
	// 1620259200000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The ID of the ticket node in the process definition.
	//
	// example:
	//
	// APPROVAL__6zu2QjAz
	TaskDefinitionNodeId *string `json:"TaskDefinitionNodeId,omitempty" xml:"TaskDefinitionNodeId,omitempty"`
	// The type of the ticket node in the process definition.
	//
	// example:
	//
	// APPROVAL
	TaskDefinitionNodeType *string `json:"TaskDefinitionNodeType,omitempty" xml:"TaskDefinitionNodeType,omitempty"`
	// The ID of the node.
	//
	// example:
	//
	// eb039a4a6a5742c6b44ccff0c1fca745
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The name of the node.
	//
	// example:
	//
	// 节点1
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// The ID of the ticket.
	//
	// example:
	//
	// 5491d3b4-14ee-4341-b5f1-db2c78beddeb
	TicketId *string `json:"TicketId,omitempty" xml:"TicketId,omitempty"`
}

func (s ListTicketTasksResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListTicketTasksResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListTicketTasksResponseBodyData) GetAction() *string {
	return s.Action
}

func (s *ListTicketTasksResponseBodyData) GetAssignee() *string {
	return s.Assignee
}

func (s *ListTicketTasksResponseBodyData) GetAssigneeName() *string {
	return s.AssigneeName
}

func (s *ListTicketTasksResponseBodyData) GetComment() *string {
	return s.Comment
}

func (s *ListTicketTasksResponseBodyData) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListTicketTasksResponseBodyData) GetFileKeys() []*string {
	return s.FileKeys
}

func (s *ListTicketTasksResponseBodyData) GetFileUrls() []*string {
	return s.FileUrls
}

func (s *ListTicketTasksResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListTicketTasksResponseBodyData) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListTicketTasksResponseBodyData) GetTaskDefinitionNodeId() *string {
	return s.TaskDefinitionNodeId
}

func (s *ListTicketTasksResponseBodyData) GetTaskDefinitionNodeType() *string {
	return s.TaskDefinitionNodeType
}

func (s *ListTicketTasksResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *ListTicketTasksResponseBodyData) GetTaskName() *string {
	return s.TaskName
}

func (s *ListTicketTasksResponseBodyData) GetTicketId() *string {
	return s.TicketId
}

func (s *ListTicketTasksResponseBodyData) SetAction(v string) *ListTicketTasksResponseBodyData {
	s.Action = &v
	return s
}

func (s *ListTicketTasksResponseBodyData) SetAssignee(v string) *ListTicketTasksResponseBodyData {
	s.Assignee = &v
	return s
}

func (s *ListTicketTasksResponseBodyData) SetAssigneeName(v string) *ListTicketTasksResponseBodyData {
	s.AssigneeName = &v
	return s
}

func (s *ListTicketTasksResponseBodyData) SetComment(v string) *ListTicketTasksResponseBodyData {
	s.Comment = &v
	return s
}

func (s *ListTicketTasksResponseBodyData) SetEndTime(v int64) *ListTicketTasksResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *ListTicketTasksResponseBodyData) SetFileKeys(v []*string) *ListTicketTasksResponseBodyData {
	s.FileKeys = v
	return s
}

func (s *ListTicketTasksResponseBodyData) SetFileUrls(v []*string) *ListTicketTasksResponseBodyData {
	s.FileUrls = v
	return s
}

func (s *ListTicketTasksResponseBodyData) SetInstanceId(v string) *ListTicketTasksResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *ListTicketTasksResponseBodyData) SetStartTime(v int64) *ListTicketTasksResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *ListTicketTasksResponseBodyData) SetTaskDefinitionNodeId(v string) *ListTicketTasksResponseBodyData {
	s.TaskDefinitionNodeId = &v
	return s
}

func (s *ListTicketTasksResponseBodyData) SetTaskDefinitionNodeType(v string) *ListTicketTasksResponseBodyData {
	s.TaskDefinitionNodeType = &v
	return s
}

func (s *ListTicketTasksResponseBodyData) SetTaskId(v string) *ListTicketTasksResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *ListTicketTasksResponseBodyData) SetTaskName(v string) *ListTicketTasksResponseBodyData {
	s.TaskName = &v
	return s
}

func (s *ListTicketTasksResponseBodyData) SetTicketId(v string) *ListTicketTasksResponseBodyData {
	s.TicketId = &v
	return s
}

func (s *ListTicketTasksResponseBodyData) Validate() error {
	return dara.Validate(s)
}
