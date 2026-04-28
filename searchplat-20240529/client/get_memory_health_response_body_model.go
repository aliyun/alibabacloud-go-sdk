// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMemoryHealthResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLatency(v int32) *GetMemoryHealthResponseBody
	GetLatency() *int32
	SetRequestId(v string) *GetMemoryHealthResponseBody
	GetRequestId() *string
	SetResult(v *GetMemoryHealthResponseBodyResult) *GetMemoryHealthResponseBody
	GetResult() *GetMemoryHealthResponseBodyResult
	SetStatus(v string) *GetMemoryHealthResponseBody
	GetStatus() *string
}

type GetMemoryHealthResponseBody struct {
	Latency   *int32                             `json:"latency,omitempty" xml:"latency,omitempty"`
	RequestId *string                            `json:"request_id,omitempty" xml:"request_id,omitempty"`
	Result    *GetMemoryHealthResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Status    *string                            `json:"status,omitempty" xml:"status,omitempty"`
}

func (s GetMemoryHealthResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMemoryHealthResponseBody) GoString() string {
	return s.String()
}

func (s *GetMemoryHealthResponseBody) GetLatency() *int32 {
	return s.Latency
}

func (s *GetMemoryHealthResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMemoryHealthResponseBody) GetResult() *GetMemoryHealthResponseBodyResult {
	return s.Result
}

func (s *GetMemoryHealthResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetMemoryHealthResponseBody) SetLatency(v int32) *GetMemoryHealthResponseBody {
	s.Latency = &v
	return s
}

func (s *GetMemoryHealthResponseBody) SetRequestId(v string) *GetMemoryHealthResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMemoryHealthResponseBody) SetResult(v *GetMemoryHealthResponseBodyResult) *GetMemoryHealthResponseBody {
	s.Result = v
	return s
}

func (s *GetMemoryHealthResponseBody) SetStatus(v string) *GetMemoryHealthResponseBody {
	s.Status = &v
	return s
}

func (s *GetMemoryHealthResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetMemoryHealthResponseBodyResult struct {
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s GetMemoryHealthResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetMemoryHealthResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetMemoryHealthResponseBodyResult) GetStatus() *string {
	return s.Status
}

func (s *GetMemoryHealthResponseBodyResult) SetStatus(v string) *GetMemoryHealthResponseBodyResult {
	s.Status = &v
	return s
}

func (s *GetMemoryHealthResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
