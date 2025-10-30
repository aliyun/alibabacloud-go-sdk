// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEvaluationScoreHistoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListEvaluationScoreHistoryResponseBody
	GetRequestId() *string
	SetScoreHistory(v *ListEvaluationScoreHistoryResponseBodyScoreHistory) *ListEvaluationScoreHistoryResponseBody
	GetScoreHistory() *ListEvaluationScoreHistoryResponseBodyScoreHistory
}

type ListEvaluationScoreHistoryResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// AC9BD94C-D20C-4D27-88D4-89E8D75C051B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The historical scores.
	ScoreHistory *ListEvaluationScoreHistoryResponseBodyScoreHistory `json:"ScoreHistory,omitempty" xml:"ScoreHistory,omitempty" type:"Struct"`
}

func (s ListEvaluationScoreHistoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListEvaluationScoreHistoryResponseBody) GoString() string {
	return s.String()
}

func (s *ListEvaluationScoreHistoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListEvaluationScoreHistoryResponseBody) GetScoreHistory() *ListEvaluationScoreHistoryResponseBodyScoreHistory {
	return s.ScoreHistory
}

func (s *ListEvaluationScoreHistoryResponseBody) SetRequestId(v string) *ListEvaluationScoreHistoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEvaluationScoreHistoryResponseBody) SetScoreHistory(v *ListEvaluationScoreHistoryResponseBodyScoreHistory) *ListEvaluationScoreHistoryResponseBody {
	s.ScoreHistory = v
	return s
}

func (s *ListEvaluationScoreHistoryResponseBody) Validate() error {
	if s.ScoreHistory != nil {
		if err := s.ScoreHistory.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListEvaluationScoreHistoryResponseBodyScoreHistory struct {
	// The historical scores.
	TotalScoreHistory []*ListEvaluationScoreHistoryResponseBodyScoreHistoryTotalScoreHistory `json:"TotalScoreHistory,omitempty" xml:"TotalScoreHistory,omitempty" type:"Repeated"`
}

func (s ListEvaluationScoreHistoryResponseBodyScoreHistory) String() string {
	return dara.Prettify(s)
}

func (s ListEvaluationScoreHistoryResponseBodyScoreHistory) GoString() string {
	return s.String()
}

func (s *ListEvaluationScoreHistoryResponseBodyScoreHistory) GetTotalScoreHistory() []*ListEvaluationScoreHistoryResponseBodyScoreHistoryTotalScoreHistory {
	return s.TotalScoreHistory
}

func (s *ListEvaluationScoreHistoryResponseBodyScoreHistory) SetTotalScoreHistory(v []*ListEvaluationScoreHistoryResponseBodyScoreHistoryTotalScoreHistory) *ListEvaluationScoreHistoryResponseBodyScoreHistory {
	s.TotalScoreHistory = v
	return s
}

func (s *ListEvaluationScoreHistoryResponseBodyScoreHistory) Validate() error {
	if s.TotalScoreHistory != nil {
		for _, item := range s.TotalScoreHistory {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListEvaluationScoreHistoryResponseBodyScoreHistoryTotalScoreHistory struct {
	// The time when the score was generated. The time is in UTC.
	//
	// example:
	//
	// 2024-06-30T03:34:02Z
	EvaluationTime *string `json:"EvaluationTime,omitempty" xml:"EvaluationTime,omitempty"`
	// The score.
	//
	// Valid values: 0 to 1.
	//
	// example:
	//
	// 0.6753
	Score *float64 `json:"Score,omitempty" xml:"Score,omitempty"`
}

func (s ListEvaluationScoreHistoryResponseBodyScoreHistoryTotalScoreHistory) String() string {
	return dara.Prettify(s)
}

func (s ListEvaluationScoreHistoryResponseBodyScoreHistoryTotalScoreHistory) GoString() string {
	return s.String()
}

func (s *ListEvaluationScoreHistoryResponseBodyScoreHistoryTotalScoreHistory) GetEvaluationTime() *string {
	return s.EvaluationTime
}

func (s *ListEvaluationScoreHistoryResponseBodyScoreHistoryTotalScoreHistory) GetScore() *float64 {
	return s.Score
}

func (s *ListEvaluationScoreHistoryResponseBodyScoreHistoryTotalScoreHistory) SetEvaluationTime(v string) *ListEvaluationScoreHistoryResponseBodyScoreHistoryTotalScoreHistory {
	s.EvaluationTime = &v
	return s
}

func (s *ListEvaluationScoreHistoryResponseBodyScoreHistoryTotalScoreHistory) SetScore(v float64) *ListEvaluationScoreHistoryResponseBodyScoreHistoryTotalScoreHistory {
	s.Score = &v
	return s
}

func (s *ListEvaluationScoreHistoryResponseBodyScoreHistoryTotalScoreHistory) Validate() error {
	return dara.Validate(s)
}
