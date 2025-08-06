// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePlayFirstFrameDurationMetricDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBeginTs(v string) *DescribePlayFirstFrameDurationMetricDataRequest
	GetBeginTs() *string
	SetEndTs(v string) *DescribePlayFirstFrameDurationMetricDataRequest
	GetEndTs() *string
	SetTraceId(v string) *DescribePlayFirstFrameDurationMetricDataRequest
	GetTraceId() *string
}

type DescribePlayFirstFrameDurationMetricDataRequest struct {
	// This parameter is required.
	BeginTs *string `json:"BeginTs,omitempty" xml:"BeginTs,omitempty"`
	// This parameter is required.
	EndTs   *string `json:"EndTs,omitempty" xml:"EndTs,omitempty"`
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s DescribePlayFirstFrameDurationMetricDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePlayFirstFrameDurationMetricDataRequest) GoString() string {
	return s.String()
}

func (s *DescribePlayFirstFrameDurationMetricDataRequest) GetBeginTs() *string {
	return s.BeginTs
}

func (s *DescribePlayFirstFrameDurationMetricDataRequest) GetEndTs() *string {
	return s.EndTs
}

func (s *DescribePlayFirstFrameDurationMetricDataRequest) GetTraceId() *string {
	return s.TraceId
}

func (s *DescribePlayFirstFrameDurationMetricDataRequest) SetBeginTs(v string) *DescribePlayFirstFrameDurationMetricDataRequest {
	s.BeginTs = &v
	return s
}

func (s *DescribePlayFirstFrameDurationMetricDataRequest) SetEndTs(v string) *DescribePlayFirstFrameDurationMetricDataRequest {
	s.EndTs = &v
	return s
}

func (s *DescribePlayFirstFrameDurationMetricDataRequest) SetTraceId(v string) *DescribePlayFirstFrameDurationMetricDataRequest {
	s.TraceId = &v
	return s
}

func (s *DescribePlayFirstFrameDurationMetricDataRequest) Validate() error {
	return dara.Validate(s)
}
