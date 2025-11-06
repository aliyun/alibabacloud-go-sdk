// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourceGroupMetricDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBeginTime(v int64) *ListResourceGroupMetricDataRequest
	GetBeginTime() *int64
	SetEndTime(v int64) *ListResourceGroupMetricDataRequest
	GetEndTime() *int64
	SetLength(v int32) *ListResourceGroupMetricDataRequest
	GetLength() *int32
	SetMetricName(v string) *ListResourceGroupMetricDataRequest
	GetMetricName() *string
	SetNextToken(v string) *ListResourceGroupMetricDataRequest
	GetNextToken() *string
	SetPeriod(v string) *ListResourceGroupMetricDataRequest
	GetPeriod() *string
	SetResourceGroupId(v string) *ListResourceGroupMetricDataRequest
	GetResourceGroupId() *string
}

type ListResourceGroupMetricDataRequest struct {
	// example:
	//
	// 1593950832000
	BeginTime *int64 `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// example:
	//
	// 1750176000000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 100
	Length *int32 `json:"Length,omitempty" xml:"Length,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// CUSpec
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// example:
	//
	// FFqBJBxE8I0PE0IUO6K69k7m2FfyWNNc2qQ9ReUkazhz9VA7dWZKlxBcjUwOV0imSM
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 60
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Serverless_res_group_524257424564736_6831777003XXXXX
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ListResourceGroupMetricDataRequest) String() string {
	return dara.Prettify(s)
}

func (s ListResourceGroupMetricDataRequest) GoString() string {
	return s.String()
}

func (s *ListResourceGroupMetricDataRequest) GetBeginTime() *int64 {
	return s.BeginTime
}

func (s *ListResourceGroupMetricDataRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListResourceGroupMetricDataRequest) GetLength() *int32 {
	return s.Length
}

func (s *ListResourceGroupMetricDataRequest) GetMetricName() *string {
	return s.MetricName
}

func (s *ListResourceGroupMetricDataRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListResourceGroupMetricDataRequest) GetPeriod() *string {
	return s.Period
}

func (s *ListResourceGroupMetricDataRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListResourceGroupMetricDataRequest) SetBeginTime(v int64) *ListResourceGroupMetricDataRequest {
	s.BeginTime = &v
	return s
}

func (s *ListResourceGroupMetricDataRequest) SetEndTime(v int64) *ListResourceGroupMetricDataRequest {
	s.EndTime = &v
	return s
}

func (s *ListResourceGroupMetricDataRequest) SetLength(v int32) *ListResourceGroupMetricDataRequest {
	s.Length = &v
	return s
}

func (s *ListResourceGroupMetricDataRequest) SetMetricName(v string) *ListResourceGroupMetricDataRequest {
	s.MetricName = &v
	return s
}

func (s *ListResourceGroupMetricDataRequest) SetNextToken(v string) *ListResourceGroupMetricDataRequest {
	s.NextToken = &v
	return s
}

func (s *ListResourceGroupMetricDataRequest) SetPeriod(v string) *ListResourceGroupMetricDataRequest {
	s.Period = &v
	return s
}

func (s *ListResourceGroupMetricDataRequest) SetResourceGroupId(v string) *ListResourceGroupMetricDataRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListResourceGroupMetricDataRequest) Validate() error {
	return dara.Validate(s)
}
