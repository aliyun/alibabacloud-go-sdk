// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCreateLogoTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *GetCreateLogoTaskResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *GetCreateLogoTaskResponseBody
	GetErrorMsg() *string
	SetRequestId(v string) *GetCreateLogoTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetCreateLogoTaskResponseBody
	GetSuccess() *bool
	SetTask(v *GetCreateLogoTaskResponseBodyTask) *GetCreateLogoTaskResponseBody
	GetTask() *GetCreateLogoTaskResponseBodyTask
}

type GetCreateLogoTaskResponseBody struct {
	// example:
	//
	// 0
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// aliuid:1998006665794443 assumeRole not exist,serviceName:realtimelogpush.dcdnservices.aliyuncs.com
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 6C6B99AC-39EC-5350-874C-204128C905E6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
	Task    *GetCreateLogoTaskResponseBodyTask `json:"Task,omitempty" xml:"Task,omitempty" type:"Struct"`
}

func (s GetCreateLogoTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCreateLogoTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetCreateLogoTaskResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetCreateLogoTaskResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *GetCreateLogoTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCreateLogoTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetCreateLogoTaskResponseBody) GetTask() *GetCreateLogoTaskResponseBodyTask {
	return s.Task
}

func (s *GetCreateLogoTaskResponseBody) SetErrorCode(v string) *GetCreateLogoTaskResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetCreateLogoTaskResponseBody) SetErrorMsg(v string) *GetCreateLogoTaskResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GetCreateLogoTaskResponseBody) SetRequestId(v string) *GetCreateLogoTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCreateLogoTaskResponseBody) SetSuccess(v bool) *GetCreateLogoTaskResponseBody {
	s.Success = &v
	return s
}

func (s *GetCreateLogoTaskResponseBody) SetTask(v *GetCreateLogoTaskResponseBodyTask) *GetCreateLogoTaskResponseBody {
	s.Task = v
	return s
}

func (s *GetCreateLogoTaskResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetCreateLogoTaskResponseBodyTask struct {
	// example:
	//
	// 604860995
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// SUCCESS
	TaskStatus *string   `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	Urls       []*string `json:"Urls,omitempty" xml:"Urls,omitempty" type:"Repeated"`
}

func (s GetCreateLogoTaskResponseBodyTask) String() string {
	return dara.Prettify(s)
}

func (s GetCreateLogoTaskResponseBodyTask) GoString() string {
	return s.String()
}

func (s *GetCreateLogoTaskResponseBodyTask) GetTaskId() *string {
	return s.TaskId
}

func (s *GetCreateLogoTaskResponseBodyTask) GetTaskStatus() *string {
	return s.TaskStatus
}

func (s *GetCreateLogoTaskResponseBodyTask) GetUrls() []*string {
	return s.Urls
}

func (s *GetCreateLogoTaskResponseBodyTask) SetTaskId(v string) *GetCreateLogoTaskResponseBodyTask {
	s.TaskId = &v
	return s
}

func (s *GetCreateLogoTaskResponseBodyTask) SetTaskStatus(v string) *GetCreateLogoTaskResponseBodyTask {
	s.TaskStatus = &v
	return s
}

func (s *GetCreateLogoTaskResponseBodyTask) SetUrls(v []*string) *GetCreateLogoTaskResponseBodyTask {
	s.Urls = v
	return s
}

func (s *GetCreateLogoTaskResponseBodyTask) Validate() error {
	return dara.Validate(s)
}
