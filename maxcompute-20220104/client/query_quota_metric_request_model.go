// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryQuotaMetricRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInterval(v int64) *QueryQuotaMetricRequest
	GetInterval() *int64
	SetNickname(v string) *QueryQuotaMetricRequest
	GetNickname() *string
	SetSubQuotaNickname(v string) *QueryQuotaMetricRequest
	GetSubQuotaNickname() *string
	SetEndTime(v int64) *QueryQuotaMetricRequest
	GetEndTime() *int64
	SetStartTime(v int64) *QueryQuotaMetricRequest
	GetStartTime() *int64
	SetStrategy(v string) *QueryQuotaMetricRequest
	GetStrategy() *string
}

type QueryQuotaMetricRequest struct {
	// example:
	//
	// 60
	Interval *int64 `json:"interval,omitempty" xml:"interval,omitempty"`
	// example:
	//
	// os_sns_p
	Nickname *string `json:"nickname,omitempty" xml:"nickname,omitempty"`
	// example:
	//
	// os_sns
	SubQuotaNickname *string `json:"subQuotaNickname,omitempty" xml:"subQuotaNickname,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1735536322
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1735534322
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// example:
	//
	// max
	Strategy *string `json:"strategy,omitempty" xml:"strategy,omitempty"`
}

func (s QueryQuotaMetricRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryQuotaMetricRequest) GoString() string {
	return s.String()
}

func (s *QueryQuotaMetricRequest) GetInterval() *int64 {
	return s.Interval
}

func (s *QueryQuotaMetricRequest) GetNickname() *string {
	return s.Nickname
}

func (s *QueryQuotaMetricRequest) GetSubQuotaNickname() *string {
	return s.SubQuotaNickname
}

func (s *QueryQuotaMetricRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *QueryQuotaMetricRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *QueryQuotaMetricRequest) GetStrategy() *string {
	return s.Strategy
}

func (s *QueryQuotaMetricRequest) SetInterval(v int64) *QueryQuotaMetricRequest {
	s.Interval = &v
	return s
}

func (s *QueryQuotaMetricRequest) SetNickname(v string) *QueryQuotaMetricRequest {
	s.Nickname = &v
	return s
}

func (s *QueryQuotaMetricRequest) SetSubQuotaNickname(v string) *QueryQuotaMetricRequest {
	s.SubQuotaNickname = &v
	return s
}

func (s *QueryQuotaMetricRequest) SetEndTime(v int64) *QueryQuotaMetricRequest {
	s.EndTime = &v
	return s
}

func (s *QueryQuotaMetricRequest) SetStartTime(v int64) *QueryQuotaMetricRequest {
	s.StartTime = &v
	return s
}

func (s *QueryQuotaMetricRequest) SetStrategy(v string) *QueryQuotaMetricRequest {
	s.Strategy = &v
	return s
}

func (s *QueryQuotaMetricRequest) Validate() error {
	return dara.Validate(s)
}
