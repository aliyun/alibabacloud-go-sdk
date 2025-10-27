// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkflowInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListWorkflowInstanceResponseBody
	GetCode() *int32
	SetData(v *ListWorkflowInstanceResponseBodyData) *ListWorkflowInstanceResponseBody
	GetData() *ListWorkflowInstanceResponseBodyData
	SetMessage(v string) *ListWorkflowInstanceResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListWorkflowInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListWorkflowInstanceResponseBody
	GetSuccess() *bool
}

type ListWorkflowInstanceResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The information about workflow instances.
	Data *ListWorkflowInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned error message.
	//
	// example:
	//
	// workflowId=xxx is not existed
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 39090022-1F3B-4797-8518-6B61095F1AF0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListWorkflowInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListWorkflowInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ListWorkflowInstanceResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListWorkflowInstanceResponseBody) GetData() *ListWorkflowInstanceResponseBodyData {
	return s.Data
}

func (s *ListWorkflowInstanceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListWorkflowInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListWorkflowInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListWorkflowInstanceResponseBody) SetCode(v int32) *ListWorkflowInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *ListWorkflowInstanceResponseBody) SetData(v *ListWorkflowInstanceResponseBodyData) *ListWorkflowInstanceResponseBody {
	s.Data = v
	return s
}

func (s *ListWorkflowInstanceResponseBody) SetMessage(v string) *ListWorkflowInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *ListWorkflowInstanceResponseBody) SetRequestId(v string) *ListWorkflowInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWorkflowInstanceResponseBody) SetSuccess(v bool) *ListWorkflowInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *ListWorkflowInstanceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListWorkflowInstanceResponseBodyData struct {
	// The workflow instances.
	WfInstanceInfos []*ListWorkflowInstanceResponseBodyDataWfInstanceInfos `json:"WfInstanceInfos,omitempty" xml:"WfInstanceInfos,omitempty" type:"Repeated"`
}

func (s ListWorkflowInstanceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListWorkflowInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListWorkflowInstanceResponseBodyData) GetWfInstanceInfos() []*ListWorkflowInstanceResponseBodyDataWfInstanceInfos {
	return s.WfInstanceInfos
}

func (s *ListWorkflowInstanceResponseBodyData) SetWfInstanceInfos(v []*ListWorkflowInstanceResponseBodyDataWfInstanceInfos) *ListWorkflowInstanceResponseBodyData {
	s.WfInstanceInfos = v
	return s
}

func (s *ListWorkflowInstanceResponseBodyData) Validate() error {
	if s.WfInstanceInfos != nil {
		for _, item := range s.WfInstanceInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListWorkflowInstanceResponseBodyDataWfInstanceInfos struct {
	// The data timestamp of the workflow instance.
	//
	// example:
	//
	// 2023-01-03 18:00:00
	DataTime *string `json:"DataTime,omitempty" xml:"DataTime,omitempty"`
	// The time when the workflow instance stopped running.
	//
	// example:
	//
	// 2023-01-03 18:00:21
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The time when the workflow instance was scheduled to run.
	//
	// example:
	//
	// 2023-01-03 18:00:00
	ScheduleTime *string `json:"ScheduleTime,omitempty" xml:"ScheduleTime,omitempty"`
	// The time when the workflow instance started to run.
	//
	// example:
	//
	// 2023-01-03 18:00:01
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The state of the workflow instance. Valid values:
	//
	// 	- 1: pending
	//
	// 	- 2: preparing
	//
	// 	- 3: running
	//
	// 	- 4: successful
	//
	// 	- 5: failed
	//
	// example:
	//
	// 5
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The workflow instance ID.
	//
	// example:
	//
	// 123456
	WfInstanceId *int64 `json:"WfInstanceId,omitempty" xml:"WfInstanceId,omitempty"`
	// The workflow ID.
	//
	// example:
	//
	// 123
	WorkflowId *int64 `json:"WorkflowId,omitempty" xml:"WorkflowId,omitempty"`
}

func (s ListWorkflowInstanceResponseBodyDataWfInstanceInfos) String() string {
	return dara.Prettify(s)
}

func (s ListWorkflowInstanceResponseBodyDataWfInstanceInfos) GoString() string {
	return s.String()
}

func (s *ListWorkflowInstanceResponseBodyDataWfInstanceInfos) GetDataTime() *string {
	return s.DataTime
}

func (s *ListWorkflowInstanceResponseBodyDataWfInstanceInfos) GetEndTime() *string {
	return s.EndTime
}

func (s *ListWorkflowInstanceResponseBodyDataWfInstanceInfos) GetScheduleTime() *string {
	return s.ScheduleTime
}

func (s *ListWorkflowInstanceResponseBodyDataWfInstanceInfos) GetStartTime() *string {
	return s.StartTime
}

func (s *ListWorkflowInstanceResponseBodyDataWfInstanceInfos) GetStatus() *int32 {
	return s.Status
}

func (s *ListWorkflowInstanceResponseBodyDataWfInstanceInfos) GetWfInstanceId() *int64 {
	return s.WfInstanceId
}

func (s *ListWorkflowInstanceResponseBodyDataWfInstanceInfos) GetWorkflowId() *int64 {
	return s.WorkflowId
}

func (s *ListWorkflowInstanceResponseBodyDataWfInstanceInfos) SetDataTime(v string) *ListWorkflowInstanceResponseBodyDataWfInstanceInfos {
	s.DataTime = &v
	return s
}

func (s *ListWorkflowInstanceResponseBodyDataWfInstanceInfos) SetEndTime(v string) *ListWorkflowInstanceResponseBodyDataWfInstanceInfos {
	s.EndTime = &v
	return s
}

func (s *ListWorkflowInstanceResponseBodyDataWfInstanceInfos) SetScheduleTime(v string) *ListWorkflowInstanceResponseBodyDataWfInstanceInfos {
	s.ScheduleTime = &v
	return s
}

func (s *ListWorkflowInstanceResponseBodyDataWfInstanceInfos) SetStartTime(v string) *ListWorkflowInstanceResponseBodyDataWfInstanceInfos {
	s.StartTime = &v
	return s
}

func (s *ListWorkflowInstanceResponseBodyDataWfInstanceInfos) SetStatus(v int32) *ListWorkflowInstanceResponseBodyDataWfInstanceInfos {
	s.Status = &v
	return s
}

func (s *ListWorkflowInstanceResponseBodyDataWfInstanceInfos) SetWfInstanceId(v int64) *ListWorkflowInstanceResponseBodyDataWfInstanceInfos {
	s.WfInstanceId = &v
	return s
}

func (s *ListWorkflowInstanceResponseBodyDataWfInstanceInfos) SetWorkflowId(v int64) *ListWorkflowInstanceResponseBodyDataWfInstanceInfos {
	s.WorkflowId = &v
	return s
}

func (s *ListWorkflowInstanceResponseBodyDataWfInstanceInfos) Validate() error {
	return dara.Validate(s)
}
