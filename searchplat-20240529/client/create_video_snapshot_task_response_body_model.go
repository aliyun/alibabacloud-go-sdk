// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVideoSnapshotTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLatency(v int32) *CreateVideoSnapshotTaskResponseBody
	GetLatency() *int32
	SetRequestId(v string) *CreateVideoSnapshotTaskResponseBody
	GetRequestId() *string
	SetResult(v *CreateVideoSnapshotTaskResponseBodyResult) *CreateVideoSnapshotTaskResponseBody
	GetResult() *CreateVideoSnapshotTaskResponseBodyResult
}

type CreateVideoSnapshotTaskResponseBody struct {
	Latency   *int32                                     `json:"latency,omitempty" xml:"latency,omitempty"`
	RequestId *string                                    `json:"request_id,omitempty" xml:"request_id,omitempty"`
	Result    *CreateVideoSnapshotTaskResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s CreateVideoSnapshotTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateVideoSnapshotTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVideoSnapshotTaskResponseBody) GetLatency() *int32 {
	return s.Latency
}

func (s *CreateVideoSnapshotTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateVideoSnapshotTaskResponseBody) GetResult() *CreateVideoSnapshotTaskResponseBodyResult {
	return s.Result
}

func (s *CreateVideoSnapshotTaskResponseBody) SetLatency(v int32) *CreateVideoSnapshotTaskResponseBody {
	s.Latency = &v
	return s
}

func (s *CreateVideoSnapshotTaskResponseBody) SetRequestId(v string) *CreateVideoSnapshotTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVideoSnapshotTaskResponseBody) SetResult(v *CreateVideoSnapshotTaskResponseBodyResult) *CreateVideoSnapshotTaskResponseBody {
	s.Result = v
	return s
}

func (s *CreateVideoSnapshotTaskResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateVideoSnapshotTaskResponseBodyResult struct {
	TaskId *string `json:"task_id,omitempty" xml:"task_id,omitempty"`
}

func (s CreateVideoSnapshotTaskResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s CreateVideoSnapshotTaskResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateVideoSnapshotTaskResponseBodyResult) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateVideoSnapshotTaskResponseBodyResult) SetTaskId(v string) *CreateVideoSnapshotTaskResponseBodyResult {
	s.TaskId = &v
	return s
}

func (s *CreateVideoSnapshotTaskResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
