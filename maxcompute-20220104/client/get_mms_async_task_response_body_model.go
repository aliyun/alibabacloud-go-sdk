// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMmsAsyncTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetMmsAsyncTaskResponseBodyData) *GetMmsAsyncTaskResponseBody
	GetData() *GetMmsAsyncTaskResponseBodyData
	SetRequestId(v string) *GetMmsAsyncTaskResponseBody
	GetRequestId() *string
}

type GetMmsAsyncTaskResponseBody struct {
	Data *GetMmsAsyncTaskResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 688003E1-D1B4-5468-957E-2FFB3AC8D79B
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetMmsAsyncTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMmsAsyncTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetMmsAsyncTaskResponseBody) GetData() *GetMmsAsyncTaskResponseBodyData {
	return s.Data
}

func (s *GetMmsAsyncTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMmsAsyncTaskResponseBody) SetData(v *GetMmsAsyncTaskResponseBodyData) *GetMmsAsyncTaskResponseBody {
	s.Data = v
	return s
}

func (s *GetMmsAsyncTaskResponseBody) SetRequestId(v string) *GetMmsAsyncTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMmsAsyncTaskResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetMmsAsyncTaskResponseBodyData struct {
	// example:
	//
	// 2024-12-17 15:44:00
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// 2024-12-17 17:44:17
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// null
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// example:
	//
	// 2523
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 232
	ObjectId *int64 `json:"objectId,omitempty" xml:"objectId,omitempty"`
	// example:
	//
	// 0
	Progress *int32 `json:"progress,omitempty" xml:"progress,omitempty"`
	// example:
	//
	// null
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
	// example:
	//
	// false
	Running *bool `json:"running,omitempty" xml:"running,omitempty"`
	// example:
	//
	// 2000017
	SourceId *int64 `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// example:
	//
	// 2024-12-17 15:44:17
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// example:
	//
	// DONE
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// TASK_CREATE
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetMmsAsyncTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetMmsAsyncTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetMmsAsyncTaskResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetMmsAsyncTaskResponseBodyData) GetEndTime() *string {
	return s.EndTime
}

func (s *GetMmsAsyncTaskResponseBodyData) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *GetMmsAsyncTaskResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *GetMmsAsyncTaskResponseBodyData) GetObjectId() *int64 {
	return s.ObjectId
}

func (s *GetMmsAsyncTaskResponseBodyData) GetProgress() *int32 {
	return s.Progress
}

func (s *GetMmsAsyncTaskResponseBodyData) GetResult() *string {
	return s.Result
}

func (s *GetMmsAsyncTaskResponseBodyData) GetRunning() *bool {
	return s.Running
}

func (s *GetMmsAsyncTaskResponseBodyData) GetSourceId() *int64 {
	return s.SourceId
}

func (s *GetMmsAsyncTaskResponseBodyData) GetStartTime() *string {
	return s.StartTime
}

func (s *GetMmsAsyncTaskResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetMmsAsyncTaskResponseBodyData) GetType() *string {
	return s.Type
}

func (s *GetMmsAsyncTaskResponseBodyData) SetCreateTime(v string) *GetMmsAsyncTaskResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetMmsAsyncTaskResponseBodyData) SetEndTime(v string) *GetMmsAsyncTaskResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *GetMmsAsyncTaskResponseBodyData) SetErrorMsg(v string) *GetMmsAsyncTaskResponseBodyData {
	s.ErrorMsg = &v
	return s
}

func (s *GetMmsAsyncTaskResponseBodyData) SetId(v int64) *GetMmsAsyncTaskResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetMmsAsyncTaskResponseBodyData) SetObjectId(v int64) *GetMmsAsyncTaskResponseBodyData {
	s.ObjectId = &v
	return s
}

func (s *GetMmsAsyncTaskResponseBodyData) SetProgress(v int32) *GetMmsAsyncTaskResponseBodyData {
	s.Progress = &v
	return s
}

func (s *GetMmsAsyncTaskResponseBodyData) SetResult(v string) *GetMmsAsyncTaskResponseBodyData {
	s.Result = &v
	return s
}

func (s *GetMmsAsyncTaskResponseBodyData) SetRunning(v bool) *GetMmsAsyncTaskResponseBodyData {
	s.Running = &v
	return s
}

func (s *GetMmsAsyncTaskResponseBodyData) SetSourceId(v int64) *GetMmsAsyncTaskResponseBodyData {
	s.SourceId = &v
	return s
}

func (s *GetMmsAsyncTaskResponseBodyData) SetStartTime(v string) *GetMmsAsyncTaskResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *GetMmsAsyncTaskResponseBodyData) SetStatus(v string) *GetMmsAsyncTaskResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetMmsAsyncTaskResponseBodyData) SetType(v string) *GetMmsAsyncTaskResponseBodyData {
	s.Type = &v
	return s
}

func (s *GetMmsAsyncTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}
