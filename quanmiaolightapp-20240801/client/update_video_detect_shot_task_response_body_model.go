// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVideoDetectShotTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateVideoDetectShotTaskResponseBody
	GetCode() *string
	SetData(v *UpdateVideoDetectShotTaskResponseBodyData) *UpdateVideoDetectShotTaskResponseBody
	GetData() *UpdateVideoDetectShotTaskResponseBodyData
	SetHttpStatusCode(v int32) *UpdateVideoDetectShotTaskResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateVideoDetectShotTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateVideoDetectShotTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateVideoDetectShotTaskResponseBody
	GetSuccess() *bool
}

type UpdateVideoDetectShotTaskResponseBody struct {
	// example:
	//
	// xx
	Code *string                                    `json:"code,omitempty" xml:"code,omitempty"`
	Data *UpdateVideoDetectShotTaskResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 575D5893-01DB-5C81-A899-74F67616762A
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateVideoDetectShotTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateVideoDetectShotTaskResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateVideoDetectShotTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateVideoDetectShotTaskResponseBody) GetData() *UpdateVideoDetectShotTaskResponseBodyData {
	return s.Data
}

func (s *UpdateVideoDetectShotTaskResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateVideoDetectShotTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateVideoDetectShotTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateVideoDetectShotTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateVideoDetectShotTaskResponseBody) SetCode(v string) *UpdateVideoDetectShotTaskResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateVideoDetectShotTaskResponseBody) SetData(v *UpdateVideoDetectShotTaskResponseBodyData) *UpdateVideoDetectShotTaskResponseBody {
	s.Data = v
	return s
}

func (s *UpdateVideoDetectShotTaskResponseBody) SetHttpStatusCode(v int32) *UpdateVideoDetectShotTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateVideoDetectShotTaskResponseBody) SetMessage(v string) *UpdateVideoDetectShotTaskResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateVideoDetectShotTaskResponseBody) SetRequestId(v string) *UpdateVideoDetectShotTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateVideoDetectShotTaskResponseBody) SetSuccess(v bool) *UpdateVideoDetectShotTaskResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateVideoDetectShotTaskResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateVideoDetectShotTaskResponseBodyData struct {
	// example:
	//
	// xx
	TaskErrorMessage *string `json:"taskErrorMessage,omitempty" xml:"taskErrorMessage,omitempty"`
	// example:
	//
	// 3feb69ed02d9b1a17d0f1a942675d300
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// example:
	//
	// SUCCESSED
	TaskStatus *string `json:"taskStatus,omitempty" xml:"taskStatus,omitempty"`
}

func (s UpdateVideoDetectShotTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateVideoDetectShotTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateVideoDetectShotTaskResponseBodyData) GetTaskErrorMessage() *string {
	return s.TaskErrorMessage
}

func (s *UpdateVideoDetectShotTaskResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *UpdateVideoDetectShotTaskResponseBodyData) GetTaskStatus() *string {
	return s.TaskStatus
}

func (s *UpdateVideoDetectShotTaskResponseBodyData) SetTaskErrorMessage(v string) *UpdateVideoDetectShotTaskResponseBodyData {
	s.TaskErrorMessage = &v
	return s
}

func (s *UpdateVideoDetectShotTaskResponseBodyData) SetTaskId(v string) *UpdateVideoDetectShotTaskResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *UpdateVideoDetectShotTaskResponseBodyData) SetTaskStatus(v string) *UpdateVideoDetectShotTaskResponseBodyData {
	s.TaskStatus = &v
	return s
}

func (s *UpdateVideoDetectShotTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}
