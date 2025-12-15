// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEmbeddingTuningResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLatency(v int32) *GetEmbeddingTuningResponseBody
	GetLatency() *int32
	SetRequestId(v string) *GetEmbeddingTuningResponseBody
	GetRequestId() *string
	SetResult(v *GetEmbeddingTuningResponseBodyResult) *GetEmbeddingTuningResponseBody
	GetResult() *GetEmbeddingTuningResponseBodyResult
	SetUsage(v *GetEmbeddingTuningResponseBodyUsage) *GetEmbeddingTuningResponseBody
	GetUsage() *GetEmbeddingTuningResponseBodyUsage
}

type GetEmbeddingTuningResponseBody struct {
	Latency   *int32                                `json:"latency,omitempty" xml:"latency,omitempty"`
	RequestId *string                               `json:"request_id,omitempty" xml:"request_id,omitempty"`
	Result    *GetEmbeddingTuningResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Usage     *GetEmbeddingTuningResponseBodyUsage  `json:"usage,omitempty" xml:"usage,omitempty" type:"Struct"`
}

func (s GetEmbeddingTuningResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetEmbeddingTuningResponseBody) GoString() string {
	return s.String()
}

func (s *GetEmbeddingTuningResponseBody) GetLatency() *int32 {
	return s.Latency
}

func (s *GetEmbeddingTuningResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetEmbeddingTuningResponseBody) GetResult() *GetEmbeddingTuningResponseBodyResult {
	return s.Result
}

func (s *GetEmbeddingTuningResponseBody) GetUsage() *GetEmbeddingTuningResponseBodyUsage {
	return s.Usage
}

func (s *GetEmbeddingTuningResponseBody) SetLatency(v int32) *GetEmbeddingTuningResponseBody {
	s.Latency = &v
	return s
}

func (s *GetEmbeddingTuningResponseBody) SetRequestId(v string) *GetEmbeddingTuningResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetEmbeddingTuningResponseBody) SetResult(v *GetEmbeddingTuningResponseBodyResult) *GetEmbeddingTuningResponseBody {
	s.Result = v
	return s
}

func (s *GetEmbeddingTuningResponseBody) SetUsage(v *GetEmbeddingTuningResponseBodyUsage) *GetEmbeddingTuningResponseBody {
	s.Usage = v
	return s
}

func (s *GetEmbeddingTuningResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	if s.Usage != nil {
		if err := s.Usage.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetEmbeddingTuningResponseBodyResult struct {
	Output [][]*float32 `json:"output,omitempty" xml:"output,omitempty" type:"Repeated"`
}

func (s GetEmbeddingTuningResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetEmbeddingTuningResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetEmbeddingTuningResponseBodyResult) GetOutput() [][]*float32 {
	return s.Output
}

func (s *GetEmbeddingTuningResponseBodyResult) SetOutput(v [][]*float32) *GetEmbeddingTuningResponseBodyResult {
	s.Output = v
	return s
}

func (s *GetEmbeddingTuningResponseBodyResult) Validate() error {
	return dara.Validate(s)
}

type GetEmbeddingTuningResponseBodyUsage struct {
	DocCount *int32 `json:"doc_count,omitempty" xml:"doc_count,omitempty"`
}

func (s GetEmbeddingTuningResponseBodyUsage) String() string {
	return dara.Prettify(s)
}

func (s GetEmbeddingTuningResponseBodyUsage) GoString() string {
	return s.String()
}

func (s *GetEmbeddingTuningResponseBodyUsage) GetDocCount() *int32 {
	return s.DocCount
}

func (s *GetEmbeddingTuningResponseBodyUsage) SetDocCount(v int32) *GetEmbeddingTuningResponseBodyUsage {
	s.DocCount = &v
	return s
}

func (s *GetEmbeddingTuningResponseBodyUsage) Validate() error {
	return dara.Validate(s)
}
