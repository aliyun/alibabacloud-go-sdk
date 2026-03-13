// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVideoSummarizationTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLatency(v int32) *CreateVideoSummarizationTaskResponseBody
	GetLatency() *int32
	SetRequestId(v string) *CreateVideoSummarizationTaskResponseBody
	GetRequestId() *string
	SetResult(v *CreateVideoSummarizationTaskResponseBodyResult) *CreateVideoSummarizationTaskResponseBody
	GetResult() *CreateVideoSummarizationTaskResponseBodyResult
}

type CreateVideoSummarizationTaskResponseBody struct {
	Latency   *int32                                          `json:"latency,omitempty" xml:"latency,omitempty"`
	RequestId *string                                         `json:"request_id,omitempty" xml:"request_id,omitempty"`
	Result    *CreateVideoSummarizationTaskResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s CreateVideoSummarizationTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateVideoSummarizationTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVideoSummarizationTaskResponseBody) GetLatency() *int32 {
	return s.Latency
}

func (s *CreateVideoSummarizationTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateVideoSummarizationTaskResponseBody) GetResult() *CreateVideoSummarizationTaskResponseBodyResult {
	return s.Result
}

func (s *CreateVideoSummarizationTaskResponseBody) SetLatency(v int32) *CreateVideoSummarizationTaskResponseBody {
	s.Latency = &v
	return s
}

func (s *CreateVideoSummarizationTaskResponseBody) SetRequestId(v string) *CreateVideoSummarizationTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVideoSummarizationTaskResponseBody) SetResult(v *CreateVideoSummarizationTaskResponseBodyResult) *CreateVideoSummarizationTaskResponseBody {
	s.Result = v
	return s
}

func (s *CreateVideoSummarizationTaskResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateVideoSummarizationTaskResponseBodyResult struct {
	TaskId *string `json:"task_id,omitempty" xml:"task_id,omitempty"`
}

func (s CreateVideoSummarizationTaskResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s CreateVideoSummarizationTaskResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateVideoSummarizationTaskResponseBodyResult) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateVideoSummarizationTaskResponseBodyResult) SetTaskId(v string) *CreateVideoSummarizationTaskResponseBodyResult {
	s.TaskId = &v
	return s
}

func (s *CreateVideoSummarizationTaskResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
