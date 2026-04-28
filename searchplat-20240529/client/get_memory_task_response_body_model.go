// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMemoryTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLatency(v int32) *GetMemoryTaskResponseBody
	GetLatency() *int32
	SetRequestId(v string) *GetMemoryTaskResponseBody
	GetRequestId() *string
	SetResult(v *GetMemoryTaskResponseBodyResult) *GetMemoryTaskResponseBody
	GetResult() *GetMemoryTaskResponseBodyResult
	SetStatus(v string) *GetMemoryTaskResponseBody
	GetStatus() *string
}

type GetMemoryTaskResponseBody struct {
	Latency   *int32                           `json:"latency,omitempty" xml:"latency,omitempty"`
	RequestId *string                          `json:"request_id,omitempty" xml:"request_id,omitempty"`
	Result    *GetMemoryTaskResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Status    *string                          `json:"status,omitempty" xml:"status,omitempty"`
}

func (s GetMemoryTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMemoryTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetMemoryTaskResponseBody) GetLatency() *int32 {
	return s.Latency
}

func (s *GetMemoryTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMemoryTaskResponseBody) GetResult() *GetMemoryTaskResponseBodyResult {
	return s.Result
}

func (s *GetMemoryTaskResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetMemoryTaskResponseBody) SetLatency(v int32) *GetMemoryTaskResponseBody {
	s.Latency = &v
	return s
}

func (s *GetMemoryTaskResponseBody) SetRequestId(v string) *GetMemoryTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMemoryTaskResponseBody) SetResult(v *GetMemoryTaskResponseBodyResult) *GetMemoryTaskResponseBody {
	s.Result = v
	return s
}

func (s *GetMemoryTaskResponseBody) SetStatus(v string) *GetMemoryTaskResponseBody {
	s.Status = &v
	return s
}

func (s *GetMemoryTaskResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetMemoryTaskResponseBodyResult struct {
	ErrorMessage *string   `json:"error_message,omitempty" xml:"error_message,omitempty"`
	MemoryIds    []*string `json:"memory_ids,omitempty" xml:"memory_ids,omitempty" type:"Repeated"`
	SkillIds     []*string `json:"skill_ids,omitempty" xml:"skill_ids,omitempty" type:"Repeated"`
	Status       *string   `json:"status,omitempty" xml:"status,omitempty"`
	TaskId       *string   `json:"task_id,omitempty" xml:"task_id,omitempty"`
}

func (s GetMemoryTaskResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetMemoryTaskResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetMemoryTaskResponseBodyResult) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetMemoryTaskResponseBodyResult) GetMemoryIds() []*string {
	return s.MemoryIds
}

func (s *GetMemoryTaskResponseBodyResult) GetSkillIds() []*string {
	return s.SkillIds
}

func (s *GetMemoryTaskResponseBodyResult) GetStatus() *string {
	return s.Status
}

func (s *GetMemoryTaskResponseBodyResult) GetTaskId() *string {
	return s.TaskId
}

func (s *GetMemoryTaskResponseBodyResult) SetErrorMessage(v string) *GetMemoryTaskResponseBodyResult {
	s.ErrorMessage = &v
	return s
}

func (s *GetMemoryTaskResponseBodyResult) SetMemoryIds(v []*string) *GetMemoryTaskResponseBodyResult {
	s.MemoryIds = v
	return s
}

func (s *GetMemoryTaskResponseBodyResult) SetSkillIds(v []*string) *GetMemoryTaskResponseBodyResult {
	s.SkillIds = v
	return s
}

func (s *GetMemoryTaskResponseBodyResult) SetStatus(v string) *GetMemoryTaskResponseBodyResult {
	s.Status = &v
	return s
}

func (s *GetMemoryTaskResponseBodyResult) SetTaskId(v string) *GetMemoryTaskResponseBodyResult {
	s.TaskId = &v
	return s
}

func (s *GetMemoryTaskResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
