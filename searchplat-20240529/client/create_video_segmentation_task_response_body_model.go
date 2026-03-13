// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVideoSegmentationTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLatency(v int32) *CreateVideoSegmentationTaskResponseBody
	GetLatency() *int32
	SetRequestId(v string) *CreateVideoSegmentationTaskResponseBody
	GetRequestId() *string
	SetResult(v *CreateVideoSegmentationTaskResponseBodyResult) *CreateVideoSegmentationTaskResponseBody
	GetResult() *CreateVideoSegmentationTaskResponseBodyResult
}

type CreateVideoSegmentationTaskResponseBody struct {
	Latency   *int32                                         `json:"latency,omitempty" xml:"latency,omitempty"`
	RequestId *string                                        `json:"request_id,omitempty" xml:"request_id,omitempty"`
	Result    *CreateVideoSegmentationTaskResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s CreateVideoSegmentationTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateVideoSegmentationTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVideoSegmentationTaskResponseBody) GetLatency() *int32 {
	return s.Latency
}

func (s *CreateVideoSegmentationTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateVideoSegmentationTaskResponseBody) GetResult() *CreateVideoSegmentationTaskResponseBodyResult {
	return s.Result
}

func (s *CreateVideoSegmentationTaskResponseBody) SetLatency(v int32) *CreateVideoSegmentationTaskResponseBody {
	s.Latency = &v
	return s
}

func (s *CreateVideoSegmentationTaskResponseBody) SetRequestId(v string) *CreateVideoSegmentationTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVideoSegmentationTaskResponseBody) SetResult(v *CreateVideoSegmentationTaskResponseBodyResult) *CreateVideoSegmentationTaskResponseBody {
	s.Result = v
	return s
}

func (s *CreateVideoSegmentationTaskResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateVideoSegmentationTaskResponseBodyResult struct {
	TaskId *string `json:"task_id,omitempty" xml:"task_id,omitempty"`
}

func (s CreateVideoSegmentationTaskResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s CreateVideoSegmentationTaskResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateVideoSegmentationTaskResponseBodyResult) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateVideoSegmentationTaskResponseBodyResult) SetTaskId(v string) *CreateVideoSegmentationTaskResponseBodyResult {
	s.TaskId = &v
	return s
}

func (s *CreateVideoSegmentationTaskResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
