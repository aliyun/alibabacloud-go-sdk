// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateImageAnalyzeTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLatency(v int32) *CreateImageAnalyzeTaskResponseBody
	GetLatency() *int32
	SetRequestId(v string) *CreateImageAnalyzeTaskResponseBody
	GetRequestId() *string
	SetResult(v *CreateImageAnalyzeTaskResponseBodyResult) *CreateImageAnalyzeTaskResponseBody
	GetResult() *CreateImageAnalyzeTaskResponseBodyResult
}

type CreateImageAnalyzeTaskResponseBody struct {
	Latency   *int32                                    `json:"latency,omitempty" xml:"latency,omitempty"`
	RequestId *string                                   `json:"request_id,omitempty" xml:"request_id,omitempty"`
	Result    *CreateImageAnalyzeTaskResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s CreateImageAnalyzeTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateImageAnalyzeTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateImageAnalyzeTaskResponseBody) GetLatency() *int32 {
	return s.Latency
}

func (s *CreateImageAnalyzeTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateImageAnalyzeTaskResponseBody) GetResult() *CreateImageAnalyzeTaskResponseBodyResult {
	return s.Result
}

func (s *CreateImageAnalyzeTaskResponseBody) SetLatency(v int32) *CreateImageAnalyzeTaskResponseBody {
	s.Latency = &v
	return s
}

func (s *CreateImageAnalyzeTaskResponseBody) SetRequestId(v string) *CreateImageAnalyzeTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateImageAnalyzeTaskResponseBody) SetResult(v *CreateImageAnalyzeTaskResponseBodyResult) *CreateImageAnalyzeTaskResponseBody {
	s.Result = v
	return s
}

func (s *CreateImageAnalyzeTaskResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateImageAnalyzeTaskResponseBodyResult struct {
	TaskId *string `json:"task_id,omitempty" xml:"task_id,omitempty"`
}

func (s CreateImageAnalyzeTaskResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s CreateImageAnalyzeTaskResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateImageAnalyzeTaskResponseBodyResult) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateImageAnalyzeTaskResponseBodyResult) SetTaskId(v string) *CreateImageAnalyzeTaskResponseBodyResult {
	s.TaskId = &v
	return s
}

func (s *CreateImageAnalyzeTaskResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
