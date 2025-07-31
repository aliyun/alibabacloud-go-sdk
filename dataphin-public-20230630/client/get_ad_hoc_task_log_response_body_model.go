// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAdHocTaskLogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetAdHocTaskLogResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *GetAdHocTaskLogResponseBody
	GetHttpStatusCode() *int32
	SetLogInfo(v *GetAdHocTaskLogResponseBodyLogInfo) *GetAdHocTaskLogResponseBody
	GetLogInfo() *GetAdHocTaskLogResponseBodyLogInfo
	SetMessage(v string) *GetAdHocTaskLogResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetAdHocTaskLogResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetAdHocTaskLogResponseBody
	GetSuccess() *bool
}

type GetAdHocTaskLogResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32                              `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	LogInfo        *GetAdHocTaskLogResponseBodyLogInfo `json:"LogInfo,omitempty" xml:"LogInfo,omitempty" type:"Struct"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAdHocTaskLogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAdHocTaskLogResponseBody) GoString() string {
	return s.String()
}

func (s *GetAdHocTaskLogResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetAdHocTaskLogResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetAdHocTaskLogResponseBody) GetLogInfo() *GetAdHocTaskLogResponseBodyLogInfo {
	return s.LogInfo
}

func (s *GetAdHocTaskLogResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAdHocTaskLogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAdHocTaskLogResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetAdHocTaskLogResponseBody) SetCode(v string) *GetAdHocTaskLogResponseBody {
	s.Code = &v
	return s
}

func (s *GetAdHocTaskLogResponseBody) SetHttpStatusCode(v int32) *GetAdHocTaskLogResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetAdHocTaskLogResponseBody) SetLogInfo(v *GetAdHocTaskLogResponseBodyLogInfo) *GetAdHocTaskLogResponseBody {
	s.LogInfo = v
	return s
}

func (s *GetAdHocTaskLogResponseBody) SetMessage(v string) *GetAdHocTaskLogResponseBody {
	s.Message = &v
	return s
}

func (s *GetAdHocTaskLogResponseBody) SetRequestId(v string) *GetAdHocTaskLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAdHocTaskLogResponseBody) SetSuccess(v bool) *GetAdHocTaskLogResponseBody {
	s.Success = &v
	return s
}

func (s *GetAdHocTaskLogResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetAdHocTaskLogResponseBodyLogInfo struct {
	// example:
	//
	// test
	Content   *string `json:"Content,omitempty" xml:"Content,omitempty"`
	HasNext   *bool   `json:"HasNext,omitempty" xml:"HasNext,omitempty"`
	HasResult *bool   `json:"HasResult,omitempty" xml:"HasResult,omitempty"`
	// example:
	//
	// 2021
	NextOffset *int32 `json:"NextOffset,omitempty" xml:"NextOffset,omitempty"`
	// example:
	//
	// 0
	SubTaskId *int32 `json:"SubTaskId,omitempty" xml:"SubTaskId,omitempty"`
	// example:
	//
	// MaxCompute_SQL_300000843_1611548758327
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// WAIT_RESOURCE
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
}

func (s GetAdHocTaskLogResponseBodyLogInfo) String() string {
	return dara.Prettify(s)
}

func (s GetAdHocTaskLogResponseBodyLogInfo) GoString() string {
	return s.String()
}

func (s *GetAdHocTaskLogResponseBodyLogInfo) GetContent() *string {
	return s.Content
}

func (s *GetAdHocTaskLogResponseBodyLogInfo) GetHasNext() *bool {
	return s.HasNext
}

func (s *GetAdHocTaskLogResponseBodyLogInfo) GetHasResult() *bool {
	return s.HasResult
}

func (s *GetAdHocTaskLogResponseBodyLogInfo) GetNextOffset() *int32 {
	return s.NextOffset
}

func (s *GetAdHocTaskLogResponseBodyLogInfo) GetSubTaskId() *int32 {
	return s.SubTaskId
}

func (s *GetAdHocTaskLogResponseBodyLogInfo) GetTaskId() *string {
	return s.TaskId
}

func (s *GetAdHocTaskLogResponseBodyLogInfo) GetTaskStatus() *string {
	return s.TaskStatus
}

func (s *GetAdHocTaskLogResponseBodyLogInfo) SetContent(v string) *GetAdHocTaskLogResponseBodyLogInfo {
	s.Content = &v
	return s
}

func (s *GetAdHocTaskLogResponseBodyLogInfo) SetHasNext(v bool) *GetAdHocTaskLogResponseBodyLogInfo {
	s.HasNext = &v
	return s
}

func (s *GetAdHocTaskLogResponseBodyLogInfo) SetHasResult(v bool) *GetAdHocTaskLogResponseBodyLogInfo {
	s.HasResult = &v
	return s
}

func (s *GetAdHocTaskLogResponseBodyLogInfo) SetNextOffset(v int32) *GetAdHocTaskLogResponseBodyLogInfo {
	s.NextOffset = &v
	return s
}

func (s *GetAdHocTaskLogResponseBodyLogInfo) SetSubTaskId(v int32) *GetAdHocTaskLogResponseBodyLogInfo {
	s.SubTaskId = &v
	return s
}

func (s *GetAdHocTaskLogResponseBodyLogInfo) SetTaskId(v string) *GetAdHocTaskLogResponseBodyLogInfo {
	s.TaskId = &v
	return s
}

func (s *GetAdHocTaskLogResponseBodyLogInfo) SetTaskStatus(v string) *GetAdHocTaskLogResponseBodyLogInfo {
	s.TaskStatus = &v
	return s
}

func (s *GetAdHocTaskLogResponseBodyLogInfo) Validate() error {
	return dara.Validate(s)
}
