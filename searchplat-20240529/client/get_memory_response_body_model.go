// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMemoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLatency(v int32) *GetMemoryResponseBody
	GetLatency() *int32
	SetRequestId(v string) *GetMemoryResponseBody
	GetRequestId() *string
	SetResult(v *GetMemoryResponseBodyResult) *GetMemoryResponseBody
	GetResult() *GetMemoryResponseBodyResult
	SetStatus(v string) *GetMemoryResponseBody
	GetStatus() *string
}

type GetMemoryResponseBody struct {
	Latency   *int32                       `json:"latency,omitempty" xml:"latency,omitempty"`
	RequestId *string                      `json:"request_id,omitempty" xml:"request_id,omitempty"`
	Result    *GetMemoryResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Status    *string                      `json:"status,omitempty" xml:"status,omitempty"`
}

func (s GetMemoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMemoryResponseBody) GoString() string {
	return s.String()
}

func (s *GetMemoryResponseBody) GetLatency() *int32 {
	return s.Latency
}

func (s *GetMemoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMemoryResponseBody) GetResult() *GetMemoryResponseBodyResult {
	return s.Result
}

func (s *GetMemoryResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetMemoryResponseBody) SetLatency(v int32) *GetMemoryResponseBody {
	s.Latency = &v
	return s
}

func (s *GetMemoryResponseBody) SetRequestId(v string) *GetMemoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMemoryResponseBody) SetResult(v *GetMemoryResponseBodyResult) *GetMemoryResponseBody {
	s.Result = v
	return s
}

func (s *GetMemoryResponseBody) SetStatus(v string) *GetMemoryResponseBody {
	s.Status = &v
	return s
}

func (s *GetMemoryResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetMemoryResponseBodyResult struct {
	AgentId  *string `json:"agent_id,omitempty" xml:"agent_id,omitempty"`
	Memory   *string `json:"memory,omitempty" xml:"memory,omitempty"`
	MemoryId *string `json:"memory_id,omitempty" xml:"memory_id,omitempty"`
	UserId   *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s GetMemoryResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetMemoryResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetMemoryResponseBodyResult) GetAgentId() *string {
	return s.AgentId
}

func (s *GetMemoryResponseBodyResult) GetMemory() *string {
	return s.Memory
}

func (s *GetMemoryResponseBodyResult) GetMemoryId() *string {
	return s.MemoryId
}

func (s *GetMemoryResponseBodyResult) GetUserId() *string {
	return s.UserId
}

func (s *GetMemoryResponseBodyResult) SetAgentId(v string) *GetMemoryResponseBodyResult {
	s.AgentId = &v
	return s
}

func (s *GetMemoryResponseBodyResult) SetMemory(v string) *GetMemoryResponseBodyResult {
	s.Memory = &v
	return s
}

func (s *GetMemoryResponseBodyResult) SetMemoryId(v string) *GetMemoryResponseBodyResult {
	s.MemoryId = &v
	return s
}

func (s *GetMemoryResponseBodyResult) SetUserId(v string) *GetMemoryResponseBodyResult {
	s.UserId = &v
	return s
}

func (s *GetMemoryResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
