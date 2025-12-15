// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDocumentRankResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLatency(v int32) *GetDocumentRankResponseBody
	GetLatency() *int32
	SetRequestId(v string) *GetDocumentRankResponseBody
	GetRequestId() *string
	SetResult(v *GetDocumentRankResponseBodyResult) *GetDocumentRankResponseBody
	GetResult() *GetDocumentRankResponseBodyResult
	SetUsage(v *GetDocumentRankResponseBodyUsage) *GetDocumentRankResponseBody
	GetUsage() *GetDocumentRankResponseBodyUsage
}

type GetDocumentRankResponseBody struct {
	Latency   *int32                             `json:"latency,omitempty" xml:"latency,omitempty"`
	RequestId *string                            `json:"request_id,omitempty" xml:"request_id,omitempty"`
	Result    *GetDocumentRankResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Usage     *GetDocumentRankResponseBodyUsage  `json:"usage,omitempty" xml:"usage,omitempty" type:"Struct"`
}

func (s GetDocumentRankResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDocumentRankResponseBody) GoString() string {
	return s.String()
}

func (s *GetDocumentRankResponseBody) GetLatency() *int32 {
	return s.Latency
}

func (s *GetDocumentRankResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDocumentRankResponseBody) GetResult() *GetDocumentRankResponseBodyResult {
	return s.Result
}

func (s *GetDocumentRankResponseBody) GetUsage() *GetDocumentRankResponseBodyUsage {
	return s.Usage
}

func (s *GetDocumentRankResponseBody) SetLatency(v int32) *GetDocumentRankResponseBody {
	s.Latency = &v
	return s
}

func (s *GetDocumentRankResponseBody) SetRequestId(v string) *GetDocumentRankResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDocumentRankResponseBody) SetResult(v *GetDocumentRankResponseBodyResult) *GetDocumentRankResponseBody {
	s.Result = v
	return s
}

func (s *GetDocumentRankResponseBody) SetUsage(v *GetDocumentRankResponseBodyUsage) *GetDocumentRankResponseBody {
	s.Usage = v
	return s
}

func (s *GetDocumentRankResponseBody) Validate() error {
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

type GetDocumentRankResponseBodyResult struct {
	Scores []*GetDocumentRankResponseBodyResultScores `json:"scores,omitempty" xml:"scores,omitempty" type:"Repeated"`
}

func (s GetDocumentRankResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetDocumentRankResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetDocumentRankResponseBodyResult) GetScores() []*GetDocumentRankResponseBodyResultScores {
	return s.Scores
}

func (s *GetDocumentRankResponseBodyResult) SetScores(v []*GetDocumentRankResponseBodyResultScores) *GetDocumentRankResponseBodyResult {
	s.Scores = v
	return s
}

func (s *GetDocumentRankResponseBodyResult) Validate() error {
	if s.Scores != nil {
		for _, item := range s.Scores {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetDocumentRankResponseBodyResultScores struct {
	Index *int32   `json:"index,omitempty" xml:"index,omitempty"`
	Score *float64 `json:"score,omitempty" xml:"score,omitempty"`
}

func (s GetDocumentRankResponseBodyResultScores) String() string {
	return dara.Prettify(s)
}

func (s GetDocumentRankResponseBodyResultScores) GoString() string {
	return s.String()
}

func (s *GetDocumentRankResponseBodyResultScores) GetIndex() *int32 {
	return s.Index
}

func (s *GetDocumentRankResponseBodyResultScores) GetScore() *float64 {
	return s.Score
}

func (s *GetDocumentRankResponseBodyResultScores) SetIndex(v int32) *GetDocumentRankResponseBodyResultScores {
	s.Index = &v
	return s
}

func (s *GetDocumentRankResponseBodyResultScores) SetScore(v float64) *GetDocumentRankResponseBodyResultScores {
	s.Score = &v
	return s
}

func (s *GetDocumentRankResponseBodyResultScores) Validate() error {
	return dara.Validate(s)
}

type GetDocumentRankResponseBodyUsage struct {
	DocCount *int64 `json:"doc_count,omitempty" xml:"doc_count,omitempty"`
}

func (s GetDocumentRankResponseBodyUsage) String() string {
	return dara.Prettify(s)
}

func (s GetDocumentRankResponseBodyUsage) GoString() string {
	return s.String()
}

func (s *GetDocumentRankResponseBodyUsage) GetDocCount() *int64 {
	return s.DocCount
}

func (s *GetDocumentRankResponseBodyUsage) SetDocCount(v int64) *GetDocumentRankResponseBodyUsage {
	s.DocCount = &v
	return s
}

func (s *GetDocumentRankResponseBodyUsage) Validate() error {
	return dara.Validate(s)
}
