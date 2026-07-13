// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstanceAsyncTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateInstanceAsyncTaskResponseBody
	GetCode() *string
	SetData(v *UpdateInstanceAsyncTaskResponseBodyData) *UpdateInstanceAsyncTaskResponseBody
	GetData() *UpdateInstanceAsyncTaskResponseBodyData
	SetHttpStatusCode(v int32) *UpdateInstanceAsyncTaskResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateInstanceAsyncTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateInstanceAsyncTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateInstanceAsyncTaskResponseBody
	GetSuccess() *bool
}

type UpdateInstanceAsyncTaskResponseBody struct {
	// example:
	//
	// SUCCESS
	Code *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *UpdateInstanceAsyncTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateInstanceAsyncTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceAsyncTaskResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateInstanceAsyncTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateInstanceAsyncTaskResponseBody) GetData() *UpdateInstanceAsyncTaskResponseBodyData {
	return s.Data
}

func (s *UpdateInstanceAsyncTaskResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateInstanceAsyncTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateInstanceAsyncTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateInstanceAsyncTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateInstanceAsyncTaskResponseBody) SetCode(v string) *UpdateInstanceAsyncTaskResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateInstanceAsyncTaskResponseBody) SetData(v *UpdateInstanceAsyncTaskResponseBodyData) *UpdateInstanceAsyncTaskResponseBody {
	s.Data = v
	return s
}

func (s *UpdateInstanceAsyncTaskResponseBody) SetHttpStatusCode(v int32) *UpdateInstanceAsyncTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateInstanceAsyncTaskResponseBody) SetMessage(v string) *UpdateInstanceAsyncTaskResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateInstanceAsyncTaskResponseBody) SetRequestId(v string) *UpdateInstanceAsyncTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateInstanceAsyncTaskResponseBody) SetSuccess(v bool) *UpdateInstanceAsyncTaskResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateInstanceAsyncTaskResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateInstanceAsyncTaskResponseBodyData struct {
	CreatedAt            *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	CurrentStep          *string `json:"CurrentStep,omitempty" xml:"CurrentStep,omitempty"`
	ModifiedAt           *string `json:"ModifiedAt,omitempty" xml:"ModifiedAt,omitempty"`
	TaskCode             *string `json:"TaskCode,omitempty" xml:"TaskCode,omitempty"`
	TaskId               *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskStatus           *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	WaitingForUserAction *bool   `json:"WaitingForUserAction,omitempty" xml:"WaitingForUserAction,omitempty"`
}

func (s UpdateInstanceAsyncTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceAsyncTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateInstanceAsyncTaskResponseBodyData) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *UpdateInstanceAsyncTaskResponseBodyData) GetCurrentStep() *string {
	return s.CurrentStep
}

func (s *UpdateInstanceAsyncTaskResponseBodyData) GetModifiedAt() *string {
	return s.ModifiedAt
}

func (s *UpdateInstanceAsyncTaskResponseBodyData) GetTaskCode() *string {
	return s.TaskCode
}

func (s *UpdateInstanceAsyncTaskResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *UpdateInstanceAsyncTaskResponseBodyData) GetTaskStatus() *string {
	return s.TaskStatus
}

func (s *UpdateInstanceAsyncTaskResponseBodyData) GetWaitingForUserAction() *bool {
	return s.WaitingForUserAction
}

func (s *UpdateInstanceAsyncTaskResponseBodyData) SetCreatedAt(v string) *UpdateInstanceAsyncTaskResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *UpdateInstanceAsyncTaskResponseBodyData) SetCurrentStep(v string) *UpdateInstanceAsyncTaskResponseBodyData {
	s.CurrentStep = &v
	return s
}

func (s *UpdateInstanceAsyncTaskResponseBodyData) SetModifiedAt(v string) *UpdateInstanceAsyncTaskResponseBodyData {
	s.ModifiedAt = &v
	return s
}

func (s *UpdateInstanceAsyncTaskResponseBodyData) SetTaskCode(v string) *UpdateInstanceAsyncTaskResponseBodyData {
	s.TaskCode = &v
	return s
}

func (s *UpdateInstanceAsyncTaskResponseBodyData) SetTaskId(v string) *UpdateInstanceAsyncTaskResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *UpdateInstanceAsyncTaskResponseBodyData) SetTaskStatus(v string) *UpdateInstanceAsyncTaskResponseBodyData {
	s.TaskStatus = &v
	return s
}

func (s *UpdateInstanceAsyncTaskResponseBodyData) SetWaitingForUserAction(v bool) *UpdateInstanceAsyncTaskResponseBodyData {
	s.WaitingForUserAction = &v
	return s
}

func (s *UpdateInstanceAsyncTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}
