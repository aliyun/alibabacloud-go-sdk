// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryHistoryActiveUserCountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetActiveUserCount(v *QueryHistoryActiveUserCountResponseBodyActiveUserCount) *QueryHistoryActiveUserCountResponseBody
	GetActiveUserCount() *QueryHistoryActiveUserCountResponseBodyActiveUserCount
	SetRequestId(v string) *QueryHistoryActiveUserCountResponseBody
	GetRequestId() *string
}

type QueryHistoryActiveUserCountResponseBody struct {
	ActiveUserCount *QueryHistoryActiveUserCountResponseBodyActiveUserCount `json:"ActiveUserCount,omitempty" xml:"ActiveUserCount,omitempty" type:"Struct"`
	// example:
	//
	// 1234567890abcdefg
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryHistoryActiveUserCountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryHistoryActiveUserCountResponseBody) GoString() string {
	return s.String()
}

func (s *QueryHistoryActiveUserCountResponseBody) GetActiveUserCount() *QueryHistoryActiveUserCountResponseBodyActiveUserCount {
	return s.ActiveUserCount
}

func (s *QueryHistoryActiveUserCountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryHistoryActiveUserCountResponseBody) SetActiveUserCount(v *QueryHistoryActiveUserCountResponseBodyActiveUserCount) *QueryHistoryActiveUserCountResponseBody {
	s.ActiveUserCount = v
	return s
}

func (s *QueryHistoryActiveUserCountResponseBody) SetRequestId(v string) *QueryHistoryActiveUserCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryHistoryActiveUserCountResponseBody) Validate() error {
	if s.ActiveUserCount != nil {
		if err := s.ActiveUserCount.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryHistoryActiveUserCountResponseBodyActiveUserCount struct {
	// example:
	//
	// 20
	DailyActiveUserCount *int32 `json:"DailyActiveUserCount,omitempty" xml:"DailyActiveUserCount,omitempty"`
	// example:
	//
	// 300
	MonthlyActiveUserCount *int32 `json:"MonthlyActiveUserCount,omitempty" xml:"MonthlyActiveUserCount,omitempty"`
}

func (s QueryHistoryActiveUserCountResponseBodyActiveUserCount) String() string {
	return dara.Prettify(s)
}

func (s QueryHistoryActiveUserCountResponseBodyActiveUserCount) GoString() string {
	return s.String()
}

func (s *QueryHistoryActiveUserCountResponseBodyActiveUserCount) GetDailyActiveUserCount() *int32 {
	return s.DailyActiveUserCount
}

func (s *QueryHistoryActiveUserCountResponseBodyActiveUserCount) GetMonthlyActiveUserCount() *int32 {
	return s.MonthlyActiveUserCount
}

func (s *QueryHistoryActiveUserCountResponseBodyActiveUserCount) SetDailyActiveUserCount(v int32) *QueryHistoryActiveUserCountResponseBodyActiveUserCount {
	s.DailyActiveUserCount = &v
	return s
}

func (s *QueryHistoryActiveUserCountResponseBodyActiveUserCount) SetMonthlyActiveUserCount(v int32) *QueryHistoryActiveUserCountResponseBodyActiveUserCount {
	s.MonthlyActiveUserCount = &v
	return s
}

func (s *QueryHistoryActiveUserCountResponseBodyActiveUserCount) Validate() error {
	return dara.Validate(s)
}
