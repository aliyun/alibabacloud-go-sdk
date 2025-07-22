// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMetricsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDimensions(v string) *GetMetricsRequest
	GetDimensions() *string
	SetEndTime(v string) *GetMetricsRequest
	GetEndTime() *string
	SetLength(v string) *GetMetricsRequest
	GetLength() *string
	SetMetricName(v string) *GetMetricsRequest
	GetMetricName() *string
	SetNamespace(v string) *GetMetricsRequest
	GetNamespace() *string
	SetNextToken(v string) *GetMetricsRequest
	GetNextToken() *string
	SetPeriod(v string) *GetMetricsRequest
	GetPeriod() *string
	SetStartTime(v string) *GetMetricsRequest
	GetStartTime() *string
}

type GetMetricsRequest struct {
	// example:
	//
	// {"userId":"16122852825*****","jobId":"dsw-328d2bbf605*****","regionId":"cn-wulanchabu","pod":"dsw-45680-76766f8778-95gxh"}
	Dimensions *string `json:"Dimensions,omitempty" xml:"Dimensions,omitempty"`
	// example:
	//
	// 2020-11-08T16:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 10
	Length *string `json:"Length,omitempty" xml:"Length,omitempty"`
	// example:
	//
	// INSTANCE_SPEC_MEMORY_SWAP
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// example:
	//
	// acs_pai_dsw
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// example:
	//
	// 15761485350009dd70bb64cff1f0fff750b08ffff073be5fb1e785e2b020f1a949d5ea14aea7fed82f01dd8****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 60
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// example:
	//
	// 2020-11-08T15:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetMetricsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMetricsRequest) GoString() string {
	return s.String()
}

func (s *GetMetricsRequest) GetDimensions() *string {
	return s.Dimensions
}

func (s *GetMetricsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *GetMetricsRequest) GetLength() *string {
	return s.Length
}

func (s *GetMetricsRequest) GetMetricName() *string {
	return s.MetricName
}

func (s *GetMetricsRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *GetMetricsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *GetMetricsRequest) GetPeriod() *string {
	return s.Period
}

func (s *GetMetricsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *GetMetricsRequest) SetDimensions(v string) *GetMetricsRequest {
	s.Dimensions = &v
	return s
}

func (s *GetMetricsRequest) SetEndTime(v string) *GetMetricsRequest {
	s.EndTime = &v
	return s
}

func (s *GetMetricsRequest) SetLength(v string) *GetMetricsRequest {
	s.Length = &v
	return s
}

func (s *GetMetricsRequest) SetMetricName(v string) *GetMetricsRequest {
	s.MetricName = &v
	return s
}

func (s *GetMetricsRequest) SetNamespace(v string) *GetMetricsRequest {
	s.Namespace = &v
	return s
}

func (s *GetMetricsRequest) SetNextToken(v string) *GetMetricsRequest {
	s.NextToken = &v
	return s
}

func (s *GetMetricsRequest) SetPeriod(v string) *GetMetricsRequest {
	s.Period = &v
	return s
}

func (s *GetMetricsRequest) SetStartTime(v string) *GetMetricsRequest {
	s.StartTime = &v
	return s
}

func (s *GetMetricsRequest) Validate() error {
	return dara.Validate(s)
}
