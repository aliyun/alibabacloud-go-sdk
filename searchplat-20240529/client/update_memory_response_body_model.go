// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMemoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLatency(v int32) *UpdateMemoryResponseBody
	GetLatency() *int32
	SetRequestId(v string) *UpdateMemoryResponseBody
	GetRequestId() *string
	SetResult(v *UpdateMemoryResponseBodyResult) *UpdateMemoryResponseBody
	GetResult() *UpdateMemoryResponseBodyResult
	SetStatus(v string) *UpdateMemoryResponseBody
	GetStatus() *string
}

type UpdateMemoryResponseBody struct {
	Latency   *int32                          `json:"latency,omitempty" xml:"latency,omitempty"`
	RequestId *string                         `json:"request_id,omitempty" xml:"request_id,omitempty"`
	Result    *UpdateMemoryResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Status    *string                         `json:"status,omitempty" xml:"status,omitempty"`
}

func (s UpdateMemoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateMemoryResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMemoryResponseBody) GetLatency() *int32 {
	return s.Latency
}

func (s *UpdateMemoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateMemoryResponseBody) GetResult() *UpdateMemoryResponseBodyResult {
	return s.Result
}

func (s *UpdateMemoryResponseBody) GetStatus() *string {
	return s.Status
}

func (s *UpdateMemoryResponseBody) SetLatency(v int32) *UpdateMemoryResponseBody {
	s.Latency = &v
	return s
}

func (s *UpdateMemoryResponseBody) SetRequestId(v string) *UpdateMemoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMemoryResponseBody) SetResult(v *UpdateMemoryResponseBodyResult) *UpdateMemoryResponseBody {
	s.Result = v
	return s
}

func (s *UpdateMemoryResponseBody) SetStatus(v string) *UpdateMemoryResponseBody {
	s.Status = &v
	return s
}

func (s *UpdateMemoryResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateMemoryResponseBodyResult struct {
	Memory   *string `json:"memory,omitempty" xml:"memory,omitempty"`
	MemoryId *string `json:"memory_id,omitempty" xml:"memory_id,omitempty"`
	Message  *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s UpdateMemoryResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s UpdateMemoryResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateMemoryResponseBodyResult) GetMemory() *string {
	return s.Memory
}

func (s *UpdateMemoryResponseBodyResult) GetMemoryId() *string {
	return s.MemoryId
}

func (s *UpdateMemoryResponseBodyResult) GetMessage() *string {
	return s.Message
}

func (s *UpdateMemoryResponseBodyResult) SetMemory(v string) *UpdateMemoryResponseBodyResult {
	s.Memory = &v
	return s
}

func (s *UpdateMemoryResponseBodyResult) SetMemoryId(v string) *UpdateMemoryResponseBodyResult {
	s.MemoryId = &v
	return s
}

func (s *UpdateMemoryResponseBodyResult) SetMessage(v string) *UpdateMemoryResponseBodyResult {
	s.Message = &v
	return s
}

func (s *UpdateMemoryResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
