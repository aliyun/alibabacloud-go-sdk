// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMemoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLatency(v int32) *CreateMemoryResponseBody
	GetLatency() *int32
	SetRequestId(v string) *CreateMemoryResponseBody
	GetRequestId() *string
	SetResult(v *CreateMemoryResponseBodyResult) *CreateMemoryResponseBody
	GetResult() *CreateMemoryResponseBodyResult
	SetStatus(v string) *CreateMemoryResponseBody
	GetStatus() *string
}

type CreateMemoryResponseBody struct {
	Latency   *int32                          `json:"latency,omitempty" xml:"latency,omitempty"`
	RequestId *string                         `json:"request_id,omitempty" xml:"request_id,omitempty"`
	Result    *CreateMemoryResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Status    *string                         `json:"status,omitempty" xml:"status,omitempty"`
}

func (s CreateMemoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMemoryResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMemoryResponseBody) GetLatency() *int32 {
	return s.Latency
}

func (s *CreateMemoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMemoryResponseBody) GetResult() *CreateMemoryResponseBodyResult {
	return s.Result
}

func (s *CreateMemoryResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CreateMemoryResponseBody) SetLatency(v int32) *CreateMemoryResponseBody {
	s.Latency = &v
	return s
}

func (s *CreateMemoryResponseBody) SetRequestId(v string) *CreateMemoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMemoryResponseBody) SetResult(v *CreateMemoryResponseBodyResult) *CreateMemoryResponseBody {
	s.Result = v
	return s
}

func (s *CreateMemoryResponseBody) SetStatus(v string) *CreateMemoryResponseBody {
	s.Status = &v
	return s
}

func (s *CreateMemoryResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateMemoryResponseBodyResult struct {
	Memory *CreateMemoryResponseBodyResultMemory `json:"memory,omitempty" xml:"memory,omitempty" type:"Struct"`
	Skill  *CreateMemoryResponseBodyResultSkill  `json:"skill,omitempty" xml:"skill,omitempty" type:"Struct"`
}

func (s CreateMemoryResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s CreateMemoryResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateMemoryResponseBodyResult) GetMemory() *CreateMemoryResponseBodyResultMemory {
	return s.Memory
}

func (s *CreateMemoryResponseBodyResult) GetSkill() *CreateMemoryResponseBodyResultSkill {
	return s.Skill
}

func (s *CreateMemoryResponseBodyResult) SetMemory(v *CreateMemoryResponseBodyResultMemory) *CreateMemoryResponseBodyResult {
	s.Memory = v
	return s
}

func (s *CreateMemoryResponseBodyResult) SetSkill(v *CreateMemoryResponseBodyResultSkill) *CreateMemoryResponseBodyResult {
	s.Skill = v
	return s
}

func (s *CreateMemoryResponseBodyResult) Validate() error {
	if s.Memory != nil {
		if err := s.Memory.Validate(); err != nil {
			return err
		}
	}
	if s.Skill != nil {
		if err := s.Skill.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateMemoryResponseBodyResultMemory struct {
	TaskId *string `json:"task_id,omitempty" xml:"task_id,omitempty"`
}

func (s CreateMemoryResponseBodyResultMemory) String() string {
	return dara.Prettify(s)
}

func (s CreateMemoryResponseBodyResultMemory) GoString() string {
	return s.String()
}

func (s *CreateMemoryResponseBodyResultMemory) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateMemoryResponseBodyResultMemory) SetTaskId(v string) *CreateMemoryResponseBodyResultMemory {
	s.TaskId = &v
	return s
}

func (s *CreateMemoryResponseBodyResultMemory) Validate() error {
	return dara.Validate(s)
}

type CreateMemoryResponseBodyResultSkill struct {
	TaskId *string `json:"task_id,omitempty" xml:"task_id,omitempty"`
}

func (s CreateMemoryResponseBodyResultSkill) String() string {
	return dara.Prettify(s)
}

func (s CreateMemoryResponseBodyResultSkill) GoString() string {
	return s.String()
}

func (s *CreateMemoryResponseBodyResultSkill) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateMemoryResponseBodyResultSkill) SetTaskId(v string) *CreateMemoryResponseBodyResultSkill {
	s.TaskId = &v
	return s
}

func (s *CreateMemoryResponseBodyResultSkill) Validate() error {
	return dara.Validate(s)
}
