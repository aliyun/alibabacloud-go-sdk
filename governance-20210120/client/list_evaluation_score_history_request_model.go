// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEvaluationScoreHistoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v int64) *ListEvaluationScoreHistoryRequest
	GetAccountId() *int64
	SetEndDate(v string) *ListEvaluationScoreHistoryRequest
	GetEndDate() *string
	SetEvaluationDomain(v string) *ListEvaluationScoreHistoryRequest
	GetEvaluationDomain() *string
	SetRegionId(v string) *ListEvaluationScoreHistoryRequest
	GetRegionId() *string
	SetStartDate(v string) *ListEvaluationScoreHistoryRequest
	GetStartDate() *string
}

type ListEvaluationScoreHistoryRequest struct {
	// The ID of the member accounts. This parameter is applicable only to the multi-account detection pattern.
	//
	// example:
	//
	// 176618589410****
	AccountId *int64 `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The end date of the query. Format: YYYY-MM-DD.
	//
	// By default, the historical scores from the last 7 days are returned.
	//
	// example:
	//
	// 2024-07-11
	EndDate          *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	EvaluationDomain *string `json:"EvaluationDomain,omitempty" xml:"EvaluationDomain,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The start date of the query. Format: YYYY-MM-DD.
	//
	// You can query records from the last 180 days.
	//
	// example:
	//
	// 2024-06-11
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s ListEvaluationScoreHistoryRequest) String() string {
	return dara.Prettify(s)
}

func (s ListEvaluationScoreHistoryRequest) GoString() string {
	return s.String()
}

func (s *ListEvaluationScoreHistoryRequest) GetAccountId() *int64 {
	return s.AccountId
}

func (s *ListEvaluationScoreHistoryRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *ListEvaluationScoreHistoryRequest) GetEvaluationDomain() *string {
	return s.EvaluationDomain
}

func (s *ListEvaluationScoreHistoryRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListEvaluationScoreHistoryRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *ListEvaluationScoreHistoryRequest) SetAccountId(v int64) *ListEvaluationScoreHistoryRequest {
	s.AccountId = &v
	return s
}

func (s *ListEvaluationScoreHistoryRequest) SetEndDate(v string) *ListEvaluationScoreHistoryRequest {
	s.EndDate = &v
	return s
}

func (s *ListEvaluationScoreHistoryRequest) SetEvaluationDomain(v string) *ListEvaluationScoreHistoryRequest {
	s.EvaluationDomain = &v
	return s
}

func (s *ListEvaluationScoreHistoryRequest) SetRegionId(v string) *ListEvaluationScoreHistoryRequest {
	s.RegionId = &v
	return s
}

func (s *ListEvaluationScoreHistoryRequest) SetStartDate(v string) *ListEvaluationScoreHistoryRequest {
	s.StartDate = &v
	return s
}

func (s *ListEvaluationScoreHistoryRequest) Validate() error {
	return dara.Validate(s)
}
