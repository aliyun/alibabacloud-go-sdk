// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeModelOperatorUsageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiKeyIds(v []*int32) *DescribeModelOperatorUsageRequest
	GetApiKeyIds() []*int32
	SetEndTime(v string) *DescribeModelOperatorUsageRequest
	GetEndTime() *string
	SetGroupBy(v string) *DescribeModelOperatorUsageRequest
	GetGroupBy() *string
	SetKeys(v []*string) *DescribeModelOperatorUsageRequest
	GetKeys() []*string
	SetModelNames(v []*string) *DescribeModelOperatorUsageRequest
	GetModelNames() []*string
	SetPeriod(v int32) *DescribeModelOperatorUsageRequest
	GetPeriod() *int32
	SetStartTime(v string) *DescribeModelOperatorUsageRequest
	GetStartTime() *string
}

type DescribeModelOperatorUsageRequest struct {
	// The list of API key IDs. Separate multiple IDs with commas (,). If this parameter is not specified, all API key IDs under the instance ID are used by default.
	//
	// > The list can contain up to 50 items.
	ApiKeyIds []*int32 `json:"ApiKeyIds,omitempty" xml:"ApiKeyIds,omitempty" type:"Repeated"`
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
	Keys []*string `json:"Keys,omitempty" xml:"Keys,omitempty" type:"Repeated"`
	// The list of model names. Separate multiple names with commas (,).
	ModelNames []*string `json:"ModelNames,omitempty" xml:"ModelNames,omitempty" type:"Repeated"`
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

func (s DescribeModelOperatorUsageRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeModelOperatorUsageRequest) GoString() string {
	return s.String()
}

func (s *DescribeModelOperatorUsageRequest) GetApiKeyIds() []*int32 {
	return s.ApiKeyIds
}

func (s *DescribeModelOperatorUsageRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeModelOperatorUsageRequest) GetGroupBy() *string {
	return s.GroupBy
}

func (s *DescribeModelOperatorUsageRequest) GetKeys() []*string {
	return s.Keys
}

func (s *DescribeModelOperatorUsageRequest) GetModelNames() []*string {
	return s.ModelNames
}

func (s *DescribeModelOperatorUsageRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *DescribeModelOperatorUsageRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeModelOperatorUsageRequest) SetApiKeyIds(v []*int32) *DescribeModelOperatorUsageRequest {
	s.ApiKeyIds = v
	return s
}

func (s *DescribeModelOperatorUsageRequest) SetEndTime(v string) *DescribeModelOperatorUsageRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeModelOperatorUsageRequest) SetGroupBy(v string) *DescribeModelOperatorUsageRequest {
	s.GroupBy = &v
	return s
}

func (s *DescribeModelOperatorUsageRequest) SetKeys(v []*string) *DescribeModelOperatorUsageRequest {
	s.Keys = v
	return s
}

func (s *DescribeModelOperatorUsageRequest) SetModelNames(v []*string) *DescribeModelOperatorUsageRequest {
	s.ModelNames = v
	return s
}

func (s *DescribeModelOperatorUsageRequest) SetPeriod(v int32) *DescribeModelOperatorUsageRequest {
	s.Period = &v
	return s
}

func (s *DescribeModelOperatorUsageRequest) SetStartTime(v string) *DescribeModelOperatorUsageRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeModelOperatorUsageRequest) Validate() error {
	return dara.Validate(s)
}
