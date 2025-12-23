// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkflowInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListWorkflowInstancesResponseBodyData) *ListWorkflowInstancesResponseBody
	GetData() []*ListWorkflowInstancesResponseBodyData
	SetNextToken(v string) *ListWorkflowInstancesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListWorkflowInstancesResponseBody
	GetRequestId() *string
	SetTotalSize(v int32) *ListWorkflowInstancesResponseBody
	GetTotalSize() *int32
}

type ListWorkflowInstancesResponseBody struct {
	Data []*ListWorkflowInstancesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// 123abc***
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 10
	TotalSize *int32 `json:"totalSize,omitempty" xml:"totalSize,omitempty"`
}

func (s ListWorkflowInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListWorkflowInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListWorkflowInstancesResponseBody) GetData() []*ListWorkflowInstancesResponseBodyData {
	return s.Data
}

func (s *ListWorkflowInstancesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListWorkflowInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListWorkflowInstancesResponseBody) GetTotalSize() *int32 {
	return s.TotalSize
}

func (s *ListWorkflowInstancesResponseBody) SetData(v []*ListWorkflowInstancesResponseBodyData) *ListWorkflowInstancesResponseBody {
	s.Data = v
	return s
}

func (s *ListWorkflowInstancesResponseBody) SetNextToken(v string) *ListWorkflowInstancesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListWorkflowInstancesResponseBody) SetRequestId(v string) *ListWorkflowInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWorkflowInstancesResponseBody) SetTotalSize(v int32) *ListWorkflowInstancesResponseBody {
	s.TotalSize = &v
	return s
}

func (s *ListWorkflowInstancesResponseBody) Validate() error {
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

type ListWorkflowInstancesResponseBodyData struct {
	// example:
	//
	// 2024-01-01 00:00:00
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// workflow_instance_name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 2024-01-01 00:00:00
	ScheduleTime *string `json:"scheduleTime,omitempty" xml:"scheduleTime,omitempty"`
	// example:
	//
	// 2024-01-01 00:00:00
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// example:
	//
	// SUCCESS
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// w-3q9jo749ne5****
	WorkflowId *string `json:"workflowId,omitempty" xml:"workflowId,omitempty"`
	// example:
	//
	// wi-3q9jo749ne5****
	WorkflowInstanceId *string `json:"workflowInstanceId,omitempty" xml:"workflowInstanceId,omitempty"`
	// example:
	//
	// 1
	WorkflowVersion *int32 `json:"workflowVersion,omitempty" xml:"workflowVersion,omitempty"`
}

func (s ListWorkflowInstancesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListWorkflowInstancesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListWorkflowInstancesResponseBodyData) GetEndTime() *string {
	return s.EndTime
}

func (s *ListWorkflowInstancesResponseBodyData) GetName() *string {
	return s.Name
}

func (s *ListWorkflowInstancesResponseBodyData) GetScheduleTime() *string {
	return s.ScheduleTime
}

func (s *ListWorkflowInstancesResponseBodyData) GetStartTime() *string {
	return s.StartTime
}

func (s *ListWorkflowInstancesResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *ListWorkflowInstancesResponseBodyData) GetWorkflowId() *string {
	return s.WorkflowId
}

func (s *ListWorkflowInstancesResponseBodyData) GetWorkflowInstanceId() *string {
	return s.WorkflowInstanceId
}

func (s *ListWorkflowInstancesResponseBodyData) GetWorkflowVersion() *int32 {
	return s.WorkflowVersion
}

func (s *ListWorkflowInstancesResponseBodyData) SetEndTime(v string) *ListWorkflowInstancesResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *ListWorkflowInstancesResponseBodyData) SetName(v string) *ListWorkflowInstancesResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListWorkflowInstancesResponseBodyData) SetScheduleTime(v string) *ListWorkflowInstancesResponseBodyData {
	s.ScheduleTime = &v
	return s
}

func (s *ListWorkflowInstancesResponseBodyData) SetStartTime(v string) *ListWorkflowInstancesResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *ListWorkflowInstancesResponseBodyData) SetStatus(v string) *ListWorkflowInstancesResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListWorkflowInstancesResponseBodyData) SetWorkflowId(v string) *ListWorkflowInstancesResponseBodyData {
	s.WorkflowId = &v
	return s
}

func (s *ListWorkflowInstancesResponseBodyData) SetWorkflowInstanceId(v string) *ListWorkflowInstancesResponseBodyData {
	s.WorkflowInstanceId = &v
	return s
}

func (s *ListWorkflowInstancesResponseBodyData) SetWorkflowVersion(v int32) *ListWorkflowInstancesResponseBodyData {
	s.WorkflowVersion = &v
	return s
}

func (s *ListWorkflowInstancesResponseBodyData) Validate() error {
	return dara.Validate(s)
}
