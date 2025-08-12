// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchCreateInstantSiteMonitorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *BatchCreateInstantSiteMonitorResponseBody
	GetCode() *string
	SetData(v []*BatchCreateInstantSiteMonitorResponseBodyData) *BatchCreateInstantSiteMonitorResponseBody
	GetData() []*BatchCreateInstantSiteMonitorResponseBodyData
	SetMessage(v string) *BatchCreateInstantSiteMonitorResponseBody
	GetMessage() *string
	SetRequestId(v string) *BatchCreateInstantSiteMonitorResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *BatchCreateInstantSiteMonitorResponseBody
	GetSuccess() *bool
}

type BatchCreateInstantSiteMonitorResponseBody struct {
	// The HTTP status code.
	//
	// >  The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The information about the site monitoring task.
	//
	// example:
	//
	// [{"taskName": "HangZhou_ECS1", "taskId": "679fbe4f-b80b-4706-91b2-5427b43e****"}]
	Data []*BatchCreateInstantSiteMonitorResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The returned message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 7AE72720-2C96-5446-9F2B-308C7CEDFF1A
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

func (s BatchCreateInstantSiteMonitorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchCreateInstantSiteMonitorResponseBody) GoString() string {
	return s.String()
}

func (s *BatchCreateInstantSiteMonitorResponseBody) GetCode() *string {
	return s.Code
}

func (s *BatchCreateInstantSiteMonitorResponseBody) GetData() []*BatchCreateInstantSiteMonitorResponseBodyData {
	return s.Data
}

func (s *BatchCreateInstantSiteMonitorResponseBody) GetMessage() *string {
	return s.Message
}

func (s *BatchCreateInstantSiteMonitorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchCreateInstantSiteMonitorResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *BatchCreateInstantSiteMonitorResponseBody) SetCode(v string) *BatchCreateInstantSiteMonitorResponseBody {
	s.Code = &v
	return s
}

func (s *BatchCreateInstantSiteMonitorResponseBody) SetData(v []*BatchCreateInstantSiteMonitorResponseBodyData) *BatchCreateInstantSiteMonitorResponseBody {
	s.Data = v
	return s
}

func (s *BatchCreateInstantSiteMonitorResponseBody) SetMessage(v string) *BatchCreateInstantSiteMonitorResponseBody {
	s.Message = &v
	return s
}

func (s *BatchCreateInstantSiteMonitorResponseBody) SetRequestId(v string) *BatchCreateInstantSiteMonitorResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchCreateInstantSiteMonitorResponseBody) SetSuccess(v bool) *BatchCreateInstantSiteMonitorResponseBody {
	s.Success = &v
	return s
}

func (s *BatchCreateInstantSiteMonitorResponseBody) Validate() error {
	return dara.Validate(s)
}

type BatchCreateInstantSiteMonitorResponseBodyData struct {
	// The ID of the site monitoring task.
	//
	// example:
	//
	// 679fbe4f-b80b-4706-91b2-5427b43e****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The name of the site monitoring task.
	//
	// example:
	//
	// HangZhou_ECS1
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s BatchCreateInstantSiteMonitorResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s BatchCreateInstantSiteMonitorResponseBodyData) GoString() string {
	return s.String()
}

func (s *BatchCreateInstantSiteMonitorResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *BatchCreateInstantSiteMonitorResponseBodyData) GetTaskName() *string {
	return s.TaskName
}

func (s *BatchCreateInstantSiteMonitorResponseBodyData) SetTaskId(v string) *BatchCreateInstantSiteMonitorResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *BatchCreateInstantSiteMonitorResponseBodyData) SetTaskName(v string) *BatchCreateInstantSiteMonitorResponseBodyData {
	s.TaskName = &v
	return s
}

func (s *BatchCreateInstantSiteMonitorResponseBodyData) Validate() error {
	return dara.Validate(s)
}
