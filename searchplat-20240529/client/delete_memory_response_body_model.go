// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMemoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLatency(v int32) *DeleteMemoryResponseBody
	GetLatency() *int32
	SetRequestId(v string) *DeleteMemoryResponseBody
	GetRequestId() *string
	SetResult(v *DeleteMemoryResponseBodyResult) *DeleteMemoryResponseBody
	GetResult() *DeleteMemoryResponseBodyResult
	SetStatus(v string) *DeleteMemoryResponseBody
	GetStatus() *string
}

type DeleteMemoryResponseBody struct {
	Latency   *int32                          `json:"latency,omitempty" xml:"latency,omitempty"`
	RequestId *string                         `json:"request_id,omitempty" xml:"request_id,omitempty"`
	Result    *DeleteMemoryResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Status    *string                         `json:"status,omitempty" xml:"status,omitempty"`
}

func (s DeleteMemoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteMemoryResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMemoryResponseBody) GetLatency() *int32 {
	return s.Latency
}

func (s *DeleteMemoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteMemoryResponseBody) GetResult() *DeleteMemoryResponseBodyResult {
	return s.Result
}

func (s *DeleteMemoryResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DeleteMemoryResponseBody) SetLatency(v int32) *DeleteMemoryResponseBody {
	s.Latency = &v
	return s
}

func (s *DeleteMemoryResponseBody) SetRequestId(v string) *DeleteMemoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMemoryResponseBody) SetResult(v *DeleteMemoryResponseBodyResult) *DeleteMemoryResponseBody {
	s.Result = v
	return s
}

func (s *DeleteMemoryResponseBody) SetStatus(v string) *DeleteMemoryResponseBody {
	s.Status = &v
	return s
}

func (s *DeleteMemoryResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteMemoryResponseBodyResult struct {
	MemoryId *string `json:"memory_id,omitempty" xml:"memory_id,omitempty"`
	Message  *string `json:"message,omitempty" xml:"message,omitempty"`
	Status   *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s DeleteMemoryResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s DeleteMemoryResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DeleteMemoryResponseBodyResult) GetMemoryId() *string {
	return s.MemoryId
}

func (s *DeleteMemoryResponseBodyResult) GetMessage() *string {
	return s.Message
}

func (s *DeleteMemoryResponseBodyResult) GetStatus() *string {
	return s.Status
}

func (s *DeleteMemoryResponseBodyResult) SetMemoryId(v string) *DeleteMemoryResponseBodyResult {
	s.MemoryId = &v
	return s
}

func (s *DeleteMemoryResponseBodyResult) SetMessage(v string) *DeleteMemoryResponseBodyResult {
	s.Message = &v
	return s
}

func (s *DeleteMemoryResponseBodyResult) SetStatus(v string) *DeleteMemoryResponseBodyResult {
	s.Status = &v
	return s
}

func (s *DeleteMemoryResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
