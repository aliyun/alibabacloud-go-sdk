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
	SetRegionId(v string) *ListEvaluationScoreHistoryRequest
	GetRegionId() *string
	SetStartDate(v string) *ListEvaluationScoreHistoryRequest
	GetStartDate() *string
}

type ListEvaluationScoreHistoryRequest struct {
	// The Alibaba Cloud account ID of the member. This parameter takes effect only when a multi-account governance maturity check is performed.
	//
	// example:
	//
	// 176618589410****
	AccountId *int64 `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The end of the time range to query. Specify the time in the YYYY-MM-DD format.
	//
	// By default, the historical scores that were generated in the seven days before the current date are queried.
	//
	// example:
	//
	// 2024-07-11
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The beginning of the time range to query. Specify the time in the YYYY-MM-DD format.
	//
	// You can query the historical scores within the previous 180 days.
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
