// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMultiModalRerankerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLatency(v int32) *GetMultiModalRerankerResponseBody
	GetLatency() *int32
	SetRequestId(v string) *GetMultiModalRerankerResponseBody
	GetRequestId() *string
	SetResult(v *GetMultiModalRerankerResponseBodyResult) *GetMultiModalRerankerResponseBody
	GetResult() *GetMultiModalRerankerResponseBodyResult
	SetUsage(v *GetMultiModalRerankerResponseBodyUsage) *GetMultiModalRerankerResponseBody
	GetUsage() *GetMultiModalRerankerResponseBodyUsage
}

type GetMultiModalRerankerResponseBody struct {
	Latency   *int32                                   `json:"latency,omitempty" xml:"latency,omitempty"`
	RequestId *string                                  `json:"request_id,omitempty" xml:"request_id,omitempty"`
	Result    *GetMultiModalRerankerResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Usage     *GetMultiModalRerankerResponseBodyUsage  `json:"usage,omitempty" xml:"usage,omitempty" type:"Struct"`
}

func (s GetMultiModalRerankerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMultiModalRerankerResponseBody) GoString() string {
	return s.String()
}

func (s *GetMultiModalRerankerResponseBody) GetLatency() *int32 {
	return s.Latency
}

func (s *GetMultiModalRerankerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMultiModalRerankerResponseBody) GetResult() *GetMultiModalRerankerResponseBodyResult {
	return s.Result
}

func (s *GetMultiModalRerankerResponseBody) GetUsage() *GetMultiModalRerankerResponseBodyUsage {
	return s.Usage
}

func (s *GetMultiModalRerankerResponseBody) SetLatency(v int32) *GetMultiModalRerankerResponseBody {
	s.Latency = &v
	return s
}

func (s *GetMultiModalRerankerResponseBody) SetRequestId(v string) *GetMultiModalRerankerResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMultiModalRerankerResponseBody) SetResult(v *GetMultiModalRerankerResponseBodyResult) *GetMultiModalRerankerResponseBody {
	s.Result = v
	return s
}

func (s *GetMultiModalRerankerResponseBody) SetUsage(v *GetMultiModalRerankerResponseBodyUsage) *GetMultiModalRerankerResponseBody {
	s.Usage = v
	return s
}

func (s *GetMultiModalRerankerResponseBody) Validate() error {
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

type GetMultiModalRerankerResponseBodyResult struct {
	Scores []*GetMultiModalRerankerResponseBodyResultScores `json:"scores,omitempty" xml:"scores,omitempty" type:"Repeated"`
}

func (s GetMultiModalRerankerResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetMultiModalRerankerResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetMultiModalRerankerResponseBodyResult) GetScores() []*GetMultiModalRerankerResponseBodyResultScores {
	return s.Scores
}

func (s *GetMultiModalRerankerResponseBodyResult) SetScores(v []*GetMultiModalRerankerResponseBodyResultScores) *GetMultiModalRerankerResponseBodyResult {
	s.Scores = v
	return s
}

func (s *GetMultiModalRerankerResponseBodyResult) Validate() error {
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

type GetMultiModalRerankerResponseBodyResultScores struct {
	Index *int32   `json:"index,omitempty" xml:"index,omitempty"`
	Score *float64 `json:"score,omitempty" xml:"score,omitempty"`
}

func (s GetMultiModalRerankerResponseBodyResultScores) String() string {
	return dara.Prettify(s)
}

func (s GetMultiModalRerankerResponseBodyResultScores) GoString() string {
	return s.String()
}

func (s *GetMultiModalRerankerResponseBodyResultScores) GetIndex() *int32 {
	return s.Index
}

func (s *GetMultiModalRerankerResponseBodyResultScores) GetScore() *float64 {
	return s.Score
}

func (s *GetMultiModalRerankerResponseBodyResultScores) SetIndex(v int32) *GetMultiModalRerankerResponseBodyResultScores {
	s.Index = &v
	return s
}

func (s *GetMultiModalRerankerResponseBodyResultScores) SetScore(v float64) *GetMultiModalRerankerResponseBodyResultScores {
	s.Score = &v
	return s
}

func (s *GetMultiModalRerankerResponseBodyResultScores) Validate() error {
	return dara.Validate(s)
}

type GetMultiModalRerankerResponseBodyUsage struct {
	ImageToken *int64 `json:"image_token,omitempty" xml:"image_token,omitempty"`
	TextToken  *int64 `json:"text_token,omitempty" xml:"text_token,omitempty"`
}

func (s GetMultiModalRerankerResponseBodyUsage) String() string {
	return dara.Prettify(s)
}

func (s GetMultiModalRerankerResponseBodyUsage) GoString() string {
	return s.String()
}

func (s *GetMultiModalRerankerResponseBodyUsage) GetImageToken() *int64 {
	return s.ImageToken
}

func (s *GetMultiModalRerankerResponseBodyUsage) GetTextToken() *int64 {
	return s.TextToken
}

func (s *GetMultiModalRerankerResponseBodyUsage) SetImageToken(v int64) *GetMultiModalRerankerResponseBodyUsage {
	s.ImageToken = &v
	return s
}

func (s *GetMultiModalRerankerResponseBodyUsage) SetTextToken(v int64) *GetMultiModalRerankerResponseBodyUsage {
	s.TextToken = &v
	return s
}

func (s *GetMultiModalRerankerResponseBodyUsage) Validate() error {
	return dara.Validate(s)
}
