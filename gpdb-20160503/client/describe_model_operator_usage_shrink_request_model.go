// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeModelOperatorUsageShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiKeyIdsShrink(v string) *DescribeModelOperatorUsageShrinkRequest
	GetApiKeyIdsShrink() *string
	SetEndTime(v string) *DescribeModelOperatorUsageShrinkRequest
	GetEndTime() *string
	SetGroupBy(v string) *DescribeModelOperatorUsageShrinkRequest
	GetGroupBy() *string
	SetKeysShrink(v string) *DescribeModelOperatorUsageShrinkRequest
	GetKeysShrink() *string
	SetModelNamesShrink(v string) *DescribeModelOperatorUsageShrinkRequest
	GetModelNamesShrink() *string
	SetPeriod(v int32) *DescribeModelOperatorUsageShrinkRequest
	GetPeriod() *int32
	SetStartTime(v string) *DescribeModelOperatorUsageShrinkRequest
	GetStartTime() *string
}

type DescribeModelOperatorUsageShrinkRequest struct {
	// The list of API key IDs. Separate multiple IDs with commas (,). If this parameter is not specified, all API key IDs under the instance ID are used by default.
	//
	// > The list can contain up to 50 items.
	ApiKeyIdsShrink *string `json:"ApiKeyIds,omitempty" xml:"ApiKeyIds,omitempty"`
	// The end time of the query. Specify the time in the <i>YYYY-MM-DDThh:mmZ</i> format (UTC).
	//
	// > The end time must be later than the start time, and the interval between the start time and end time cannot exceed 7 days.
	//
	// example:
	//
	// 2026-06-02T00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The dimension by which to split the series. Separate multiple dimensions with commas (,). The order is not significant. Valid values:
	//
	// - model (default): splits by model.
	//
	// - api_key: splits by API key.
	//
	// - model,api_key: splits by model and API key.
	//
	// example:
	//
	// model
	GroupBy *string `json:"GroupBy,omitempty" xml:"GroupBy,omitempty"`
	// The list of metrics. Separate multiple metrics with commas (,). Valid values:
	//
	// - request_count: the number of requests.
	//
	// - success_count: the number of successful requests.
	//
	// - error_count: the number of failed requests.
	//
	// - success_rate: the request success rate.
	//
	// - input_token: the number of input tokens.
	//
	// - output_token: the number of output tokens.
	//
	// - total_token: the total number of tokens.
	KeysShrink *string `json:"Keys,omitempty" xml:"Keys,omitempty"`
	// The list of model names. Separate multiple names with commas (,).
	ModelNamesShrink *string `json:"ModelNames,omitempty" xml:"ModelNames,omitempty"`
	// The time bucket size in seconds. Valid values: 1, 5, 15, 60, 300, and 3600.
	//
	// >
	//
	// > - 1. If Period is not specified, the default value is determined by the following rules:
	//
	// > - - Window range ≤ 1 hour: Period = 1.
	//
	// > - - Window range ≤ 1 day: Period = 60.
	//
	// > - - Window range ≤ 7 days: Period = 60.
	//
	// > - 2. When Period is set to 1, the window must be ≤ 1 day.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The start time of the query. Specify the time in the <i>YYYY-MM-DDThh:mmZ</i> format (UTC).
	//
	// > Only metrics within the last 30 days can be queried.
	//
	// example:
	//
	// 2026-06-01T00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeModelOperatorUsageShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeModelOperatorUsageShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeModelOperatorUsageShrinkRequest) GetApiKeyIdsShrink() *string {
	return s.ApiKeyIdsShrink
}

func (s *DescribeModelOperatorUsageShrinkRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeModelOperatorUsageShrinkRequest) GetGroupBy() *string {
	return s.GroupBy
}

func (s *DescribeModelOperatorUsageShrinkRequest) GetKeysShrink() *string {
	return s.KeysShrink
}

func (s *DescribeModelOperatorUsageShrinkRequest) GetModelNamesShrink() *string {
	return s.ModelNamesShrink
}

func (s *DescribeModelOperatorUsageShrinkRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *DescribeModelOperatorUsageShrinkRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeModelOperatorUsageShrinkRequest) SetApiKeyIdsShrink(v string) *DescribeModelOperatorUsageShrinkRequest {
	s.ApiKeyIdsShrink = &v
	return s
}

func (s *DescribeModelOperatorUsageShrinkRequest) SetEndTime(v string) *DescribeModelOperatorUsageShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeModelOperatorUsageShrinkRequest) SetGroupBy(v string) *DescribeModelOperatorUsageShrinkRequest {
	s.GroupBy = &v
	return s
}

func (s *DescribeModelOperatorUsageShrinkRequest) SetKeysShrink(v string) *DescribeModelOperatorUsageShrinkRequest {
	s.KeysShrink = &v
	return s
}

func (s *DescribeModelOperatorUsageShrinkRequest) SetModelNamesShrink(v string) *DescribeModelOperatorUsageShrinkRequest {
	s.ModelNamesShrink = &v
	return s
}

func (s *DescribeModelOperatorUsageShrinkRequest) SetPeriod(v int32) *DescribeModelOperatorUsageShrinkRequest {
	s.Period = &v
	return s
}

func (s *DescribeModelOperatorUsageShrinkRequest) SetStartTime(v string) *DescribeModelOperatorUsageShrinkRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeModelOperatorUsageShrinkRequest) Validate() error {
	return dara.Validate(s)
}
