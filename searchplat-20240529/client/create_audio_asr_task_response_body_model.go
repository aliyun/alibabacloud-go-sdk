// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAudioAsrTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLatency(v int32) *CreateAudioAsrTaskResponseBody
	GetLatency() *int32
	SetRequestId(v string) *CreateAudioAsrTaskResponseBody
	GetRequestId() *string
	SetResult(v *CreateAudioAsrTaskResponseBodyResult) *CreateAudioAsrTaskResponseBody
	GetResult() *CreateAudioAsrTaskResponseBodyResult
}

type CreateAudioAsrTaskResponseBody struct {
	Latency   *int32                                `json:"latency,omitempty" xml:"latency,omitempty"`
	RequestId *string                               `json:"request_id,omitempty" xml:"request_id,omitempty"`
	Result    *CreateAudioAsrTaskResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s CreateAudioAsrTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAudioAsrTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAudioAsrTaskResponseBody) GetLatency() *int32 {
	return s.Latency
}

func (s *CreateAudioAsrTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAudioAsrTaskResponseBody) GetResult() *CreateAudioAsrTaskResponseBodyResult {
	return s.Result
}

func (s *CreateAudioAsrTaskResponseBody) SetLatency(v int32) *CreateAudioAsrTaskResponseBody {
	s.Latency = &v
	return s
}

func (s *CreateAudioAsrTaskResponseBody) SetRequestId(v string) *CreateAudioAsrTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAudioAsrTaskResponseBody) SetResult(v *CreateAudioAsrTaskResponseBodyResult) *CreateAudioAsrTaskResponseBody {
	s.Result = v
	return s
}

func (s *CreateAudioAsrTaskResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateAudioAsrTaskResponseBodyResult struct {
	TaskId *string `json:"task_id,omitempty" xml:"task_id,omitempty"`
}

func (s CreateAudioAsrTaskResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s CreateAudioAsrTaskResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateAudioAsrTaskResponseBodyResult) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateAudioAsrTaskResponseBodyResult) SetTaskId(v string) *CreateAudioAsrTaskResponseBodyResult {
	s.TaskId = &v
	return s
}

func (s *CreateAudioAsrTaskResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
