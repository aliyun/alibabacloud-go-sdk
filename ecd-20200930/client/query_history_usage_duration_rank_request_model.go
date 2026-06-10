// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryHistoryUsageDurationRankRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizType(v int32) *QueryHistoryUsageDurationRankRequest
	GetBizType() *int32
	SetEndDate(v string) *QueryHistoryUsageDurationRankRequest
	GetEndDate() *string
	SetLimit(v int32) *QueryHistoryUsageDurationRankRequest
	GetLimit() *int32
	SetNextToken(v string) *QueryHistoryUsageDurationRankRequest
	GetNextToken() *string
	SetStartDate(v string) *QueryHistoryUsageDurationRankRequest
	GetStartDate() *string
}

type QueryHistoryUsageDurationRankRequest struct {
	// The business type.
	//
	// example:
	//
	// 1
	BizType *int32 `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// The end date of the query in `YYYY-MM-DD` format. You can query data within the last 90 days.
	//
	// example:
	//
	// 2026-04-19
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// The number of entries to return. The default value is 5, and the maximum value is 200.
	//
	// example:
	//
	// 8
	Limit *int32 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// The token that is used to retrieve the next page of results. You can obtain this token from the response to the previous request.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The start date of the query in `YYYY-MM-DD` format. You can query data within the last 90 days.
	//
	// example:
	//
	// 2026-05-07
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s QueryHistoryUsageDurationRankRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryHistoryUsageDurationRankRequest) GoString() string {
	return s.String()
}

func (s *QueryHistoryUsageDurationRankRequest) GetBizType() *int32 {
	return s.BizType
}

func (s *QueryHistoryUsageDurationRankRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *QueryHistoryUsageDurationRankRequest) GetLimit() *int32 {
	return s.Limit
}

func (s *QueryHistoryUsageDurationRankRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *QueryHistoryUsageDurationRankRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *QueryHistoryUsageDurationRankRequest) SetBizType(v int32) *QueryHistoryUsageDurationRankRequest {
	s.BizType = &v
	return s
}

func (s *QueryHistoryUsageDurationRankRequest) SetEndDate(v string) *QueryHistoryUsageDurationRankRequest {
	s.EndDate = &v
	return s
}

func (s *QueryHistoryUsageDurationRankRequest) SetLimit(v int32) *QueryHistoryUsageDurationRankRequest {
	s.Limit = &v
	return s
}

func (s *QueryHistoryUsageDurationRankRequest) SetNextToken(v string) *QueryHistoryUsageDurationRankRequest {
	s.NextToken = &v
	return s
}

func (s *QueryHistoryUsageDurationRankRequest) SetStartDate(v string) *QueryHistoryUsageDurationRankRequest {
	s.StartDate = &v
	return s
}

func (s *QueryHistoryUsageDurationRankRequest) Validate() error {
	return dara.Validate(s)
}
