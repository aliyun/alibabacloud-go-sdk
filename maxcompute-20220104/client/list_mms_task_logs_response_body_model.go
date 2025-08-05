// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMmsTaskLogsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListMmsTaskLogsResponseBodyData) *ListMmsTaskLogsResponseBody
	GetData() []*ListMmsTaskLogsResponseBodyData
	SetRequestId(v string) *ListMmsTaskLogsResponseBody
	GetRequestId() *string
}

type ListMmsTaskLogsResponseBody struct {
	Data []*ListMmsTaskLogsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// A3AE5649-EF90-54BD-86D0-C632FA950988
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListMmsTaskLogsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMmsTaskLogsResponseBody) GoString() string {
	return s.String()
}

func (s *ListMmsTaskLogsResponseBody) GetData() []*ListMmsTaskLogsResponseBodyData {
	return s.Data
}

func (s *ListMmsTaskLogsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMmsTaskLogsResponseBody) SetData(v []*ListMmsTaskLogsResponseBodyData) *ListMmsTaskLogsResponseBody {
	s.Data = v
	return s
}

func (s *ListMmsTaskLogsResponseBody) SetRequestId(v string) *ListMmsTaskLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMmsTaskLogsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListMmsTaskLogsResponseBodyData struct {
	// example:
	//
	// create schema if not exists mms_test.default;
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
	// example:
	//
	// 2024-12-17 15:44:00
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// 10000
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// ok
	Msg *string `json:"msg,omitempty" xml:"msg,omitempty"`
	// example:
	//
	// 2000015
	SourceId *int64 `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// example:
	//
	// DATA_DOING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 4023
	TaskId *int64 `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s ListMmsTaskLogsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListMmsTaskLogsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListMmsTaskLogsResponseBodyData) GetAction() *string {
	return s.Action
}

func (s *ListMmsTaskLogsResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListMmsTaskLogsResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *ListMmsTaskLogsResponseBodyData) GetMsg() *string {
	return s.Msg
}

func (s *ListMmsTaskLogsResponseBodyData) GetSourceId() *int64 {
	return s.SourceId
}

func (s *ListMmsTaskLogsResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *ListMmsTaskLogsResponseBodyData) GetTaskId() *int64 {
	return s.TaskId
}

func (s *ListMmsTaskLogsResponseBodyData) SetAction(v string) *ListMmsTaskLogsResponseBodyData {
	s.Action = &v
	return s
}

func (s *ListMmsTaskLogsResponseBodyData) SetCreateTime(v string) *ListMmsTaskLogsResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListMmsTaskLogsResponseBodyData) SetId(v int64) *ListMmsTaskLogsResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListMmsTaskLogsResponseBodyData) SetMsg(v string) *ListMmsTaskLogsResponseBodyData {
	s.Msg = &v
	return s
}

func (s *ListMmsTaskLogsResponseBodyData) SetSourceId(v int64) *ListMmsTaskLogsResponseBodyData {
	s.SourceId = &v
	return s
}

func (s *ListMmsTaskLogsResponseBodyData) SetStatus(v string) *ListMmsTaskLogsResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListMmsTaskLogsResponseBodyData) SetTaskId(v int64) *ListMmsTaskLogsResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *ListMmsTaskLogsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
