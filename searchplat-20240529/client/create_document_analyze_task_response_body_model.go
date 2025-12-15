// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDocumentAnalyzeTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLatency(v int32) *CreateDocumentAnalyzeTaskResponseBody
	GetLatency() *int32
	SetRequestId(v string) *CreateDocumentAnalyzeTaskResponseBody
	GetRequestId() *string
	SetResult(v *CreateDocumentAnalyzeTaskResponseBodyResult) *CreateDocumentAnalyzeTaskResponseBody
	GetResult() *CreateDocumentAnalyzeTaskResponseBodyResult
}

type CreateDocumentAnalyzeTaskResponseBody struct {
	Latency   *int32                                       `json:"latency,omitempty" xml:"latency,omitempty"`
	RequestId *string                                      `json:"request_id,omitempty" xml:"request_id,omitempty"`
	Result    *CreateDocumentAnalyzeTaskResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s CreateDocumentAnalyzeTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDocumentAnalyzeTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDocumentAnalyzeTaskResponseBody) GetLatency() *int32 {
	return s.Latency
}

func (s *CreateDocumentAnalyzeTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDocumentAnalyzeTaskResponseBody) GetResult() *CreateDocumentAnalyzeTaskResponseBodyResult {
	return s.Result
}

func (s *CreateDocumentAnalyzeTaskResponseBody) SetLatency(v int32) *CreateDocumentAnalyzeTaskResponseBody {
	s.Latency = &v
	return s
}

func (s *CreateDocumentAnalyzeTaskResponseBody) SetRequestId(v string) *CreateDocumentAnalyzeTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDocumentAnalyzeTaskResponseBody) SetResult(v *CreateDocumentAnalyzeTaskResponseBodyResult) *CreateDocumentAnalyzeTaskResponseBody {
	s.Result = v
	return s
}

func (s *CreateDocumentAnalyzeTaskResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateDocumentAnalyzeTaskResponseBodyResult struct {
	TaskId *string `json:"task_id,omitempty" xml:"task_id,omitempty"`
}

func (s CreateDocumentAnalyzeTaskResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s CreateDocumentAnalyzeTaskResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateDocumentAnalyzeTaskResponseBodyResult) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateDocumentAnalyzeTaskResponseBodyResult) SetTaskId(v string) *CreateDocumentAnalyzeTaskResponseBodyResult {
	s.TaskId = &v
	return s
}

func (s *CreateDocumentAnalyzeTaskResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
