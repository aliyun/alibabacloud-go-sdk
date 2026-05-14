// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePolarClawTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *DescribePolarClawTaskResponseBody
	GetApplicationId() *string
	SetCode(v int32) *DescribePolarClawTaskResponseBody
	GetCode() *int32
	SetMessage(v string) *DescribePolarClawTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribePolarClawTaskResponseBody
	GetRequestId() *string
	SetTask(v *DescribePolarClawTaskResponseBodyTask) *DescribePolarClawTaskResponseBody
	GetTask() *DescribePolarClawTaskResponseBodyTask
}

type DescribePolarClawTaskResponseBody struct {
	// example:
	//
	// pa-**************
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// C890995A-CF06-4F4D-8DB8-DD26C2******
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Task      *DescribePolarClawTaskResponseBodyTask `json:"Task,omitempty" xml:"Task,omitempty" type:"Struct"`
}

func (s DescribePolarClawTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarClawTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePolarClawTaskResponseBody) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *DescribePolarClawTaskResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DescribePolarClawTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribePolarClawTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePolarClawTaskResponseBody) GetTask() *DescribePolarClawTaskResponseBodyTask {
	return s.Task
}

func (s *DescribePolarClawTaskResponseBody) SetApplicationId(v string) *DescribePolarClawTaskResponseBody {
	s.ApplicationId = &v
	return s
}

func (s *DescribePolarClawTaskResponseBody) SetCode(v int32) *DescribePolarClawTaskResponseBody {
	s.Code = &v
	return s
}

func (s *DescribePolarClawTaskResponseBody) SetMessage(v string) *DescribePolarClawTaskResponseBody {
	s.Message = &v
	return s
}

func (s *DescribePolarClawTaskResponseBody) SetRequestId(v string) *DescribePolarClawTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePolarClawTaskResponseBody) SetTask(v *DescribePolarClawTaskResponseBodyTask) *DescribePolarClawTaskResponseBody {
	s.Task = v
	return s
}

func (s *DescribePolarClawTaskResponseBody) Validate() error {
	if s.Task != nil {
		if err := s.Task.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribePolarClawTaskResponseBodyTask struct {
	// example:
	//
	// 1778564698304
	CreatedAtMs *int64                                      `json:"CreatedAtMs,omitempty" xml:"CreatedAtMs,omitempty"`
	Error       *DescribePolarClawTaskResponseBodyTaskError `json:"Error,omitempty" xml:"Error,omitempty" type:"Struct"`
	// example:
	//
	// LoginPolarClawChannel
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	// example:
	//
	// {}
	Result map[string]interface{} `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// running
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// example:
	//
	// 5956e600-ce6e-4d11-9648-939ef3286e94
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// 1778564750541
	UpdatedAtMs *int64 `json:"UpdatedAtMs,omitempty" xml:"UpdatedAtMs,omitempty"`
}

func (s DescribePolarClawTaskResponseBodyTask) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarClawTaskResponseBodyTask) GoString() string {
	return s.String()
}

func (s *DescribePolarClawTaskResponseBodyTask) GetCreatedAtMs() *int64 {
	return s.CreatedAtMs
}

func (s *DescribePolarClawTaskResponseBodyTask) GetError() *DescribePolarClawTaskResponseBodyTaskError {
	return s.Error
}

func (s *DescribePolarClawTaskResponseBodyTask) GetOperation() *string {
	return s.Operation
}

func (s *DescribePolarClawTaskResponseBodyTask) GetResult() map[string]interface{} {
	return s.Result
}

func (s *DescribePolarClawTaskResponseBodyTask) GetState() *string {
	return s.State
}

func (s *DescribePolarClawTaskResponseBodyTask) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribePolarClawTaskResponseBodyTask) GetUpdatedAtMs() *int64 {
	return s.UpdatedAtMs
}

func (s *DescribePolarClawTaskResponseBodyTask) SetCreatedAtMs(v int64) *DescribePolarClawTaskResponseBodyTask {
	s.CreatedAtMs = &v
	return s
}

func (s *DescribePolarClawTaskResponseBodyTask) SetError(v *DescribePolarClawTaskResponseBodyTaskError) *DescribePolarClawTaskResponseBodyTask {
	s.Error = v
	return s
}

func (s *DescribePolarClawTaskResponseBodyTask) SetOperation(v string) *DescribePolarClawTaskResponseBodyTask {
	s.Operation = &v
	return s
}

func (s *DescribePolarClawTaskResponseBodyTask) SetResult(v map[string]interface{}) *DescribePolarClawTaskResponseBodyTask {
	s.Result = v
	return s
}

func (s *DescribePolarClawTaskResponseBodyTask) SetState(v string) *DescribePolarClawTaskResponseBodyTask {
	s.State = &v
	return s
}

func (s *DescribePolarClawTaskResponseBodyTask) SetTaskId(v string) *DescribePolarClawTaskResponseBodyTask {
	s.TaskId = &v
	return s
}

func (s *DescribePolarClawTaskResponseBodyTask) SetUpdatedAtMs(v int64) *DescribePolarClawTaskResponseBodyTask {
	s.UpdatedAtMs = &v
	return s
}

func (s *DescribePolarClawTaskResponseBodyTask) Validate() error {
	if s.Error != nil {
		if err := s.Error.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribePolarClawTaskResponseBodyTaskError struct {
	// example:
	//
	// INVALID_REQUEST
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// channelId format invalid
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s DescribePolarClawTaskResponseBodyTaskError) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarClawTaskResponseBodyTaskError) GoString() string {
	return s.String()
}

func (s *DescribePolarClawTaskResponseBodyTaskError) GetCode() *string {
	return s.Code
}

func (s *DescribePolarClawTaskResponseBodyTaskError) GetMessage() *string {
	return s.Message
}

func (s *DescribePolarClawTaskResponseBodyTaskError) SetCode(v string) *DescribePolarClawTaskResponseBodyTaskError {
	s.Code = &v
	return s
}

func (s *DescribePolarClawTaskResponseBodyTaskError) SetMessage(v string) *DescribePolarClawTaskResponseBodyTaskError {
	s.Message = &v
	return s
}

func (s *DescribePolarClawTaskResponseBodyTaskError) Validate() error {
	return dara.Validate(s)
}
