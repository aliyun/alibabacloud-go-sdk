// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePlaybookNumberMetricsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMetrics(v *DescribePlaybookNumberMetricsResponseBodyMetrics) *DescribePlaybookNumberMetricsResponseBody
	GetMetrics() *DescribePlaybookNumberMetricsResponseBodyMetrics
	SetRequestId(v string) *DescribePlaybookNumberMetricsResponseBody
	GetRequestId() *string
}

type DescribePlaybookNumberMetricsResponseBody struct {
	// The statistics.
	Metrics *DescribePlaybookNumberMetricsResponseBodyMetrics `json:"Metrics,omitempty" xml:"Metrics,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// D4CC979E-3D5B-5A6A-BC87-C93C9E861C7B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePlaybookNumberMetricsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePlaybookNumberMetricsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePlaybookNumberMetricsResponseBody) GetMetrics() *DescribePlaybookNumberMetricsResponseBodyMetrics {
	return s.Metrics
}

func (s *DescribePlaybookNumberMetricsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePlaybookNumberMetricsResponseBody) SetMetrics(v *DescribePlaybookNumberMetricsResponseBodyMetrics) *DescribePlaybookNumberMetricsResponseBody {
	s.Metrics = v
	return s
}

func (s *DescribePlaybookNumberMetricsResponseBody) SetRequestId(v string) *DescribePlaybookNumberMetricsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePlaybookNumberMetricsResponseBody) Validate() error {
	if s.Metrics != nil {
		if err := s.Metrics.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribePlaybookNumberMetricsResponseBodyMetrics struct {
	// The number of enabled playbooks.
	//
	// example:
	//
	// 50
	StartUpNum *int32 `json:"StartUpNum,omitempty" xml:"StartUpNum,omitempty"`
	// The total number of playbooks.
	//
	// example:
	//
	// 100
	TotalNum *int32 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s DescribePlaybookNumberMetricsResponseBodyMetrics) String() string {
	return dara.Prettify(s)
}

func (s DescribePlaybookNumberMetricsResponseBodyMetrics) GoString() string {
	return s.String()
}

func (s *DescribePlaybookNumberMetricsResponseBodyMetrics) GetStartUpNum() *int32 {
	return s.StartUpNum
}

func (s *DescribePlaybookNumberMetricsResponseBodyMetrics) GetTotalNum() *int32 {
	return s.TotalNum
}

func (s *DescribePlaybookNumberMetricsResponseBodyMetrics) SetStartUpNum(v int32) *DescribePlaybookNumberMetricsResponseBodyMetrics {
	s.StartUpNum = &v
	return s
}

func (s *DescribePlaybookNumberMetricsResponseBodyMetrics) SetTotalNum(v int32) *DescribePlaybookNumberMetricsResponseBodyMetrics {
	s.TotalNum = &v
	return s
}

func (s *DescribePlaybookNumberMetricsResponseBodyMetrics) Validate() error {
	return dara.Validate(s)
}
