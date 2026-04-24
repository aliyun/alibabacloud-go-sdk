// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMetricMetaListShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLabelsShrink(v string) *DescribeMetricMetaListShrinkRequest
	GetLabelsShrink() *string
	SetMetaFormat(v string) *DescribeMetricMetaListShrinkRequest
	GetMetaFormat() *string
	SetMetricName(v string) *DescribeMetricMetaListShrinkRequest
	GetMetricName() *string
	SetNamespace(v string) *DescribeMetricMetaListShrinkRequest
	GetNamespace() *string
	SetPageNumber(v int32) *DescribeMetricMetaListShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeMetricMetaListShrinkRequest
	GetPageSize() *int32
}

type DescribeMetricMetaListShrinkRequest struct {
	LabelsShrink *string `json:"labels,omitempty" xml:"labels,omitempty"`
	// example:
	//
	// CMS
	MetaFormat *string `json:"metaFormat,omitempty" xml:"metaFormat,omitempty"`
	// example:
	//
	// CPUUtilization
	MetricName *string `json:"metricName,omitempty" xml:"metricName,omitempty"`
	// example:
	//
	// acs_ecs_dashboard
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 2000
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s DescribeMetricMetaListShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetricMetaListShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeMetricMetaListShrinkRequest) GetLabelsShrink() *string {
	return s.LabelsShrink
}

func (s *DescribeMetricMetaListShrinkRequest) GetMetaFormat() *string {
	return s.MetaFormat
}

func (s *DescribeMetricMetaListShrinkRequest) GetMetricName() *string {
	return s.MetricName
}

func (s *DescribeMetricMetaListShrinkRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *DescribeMetricMetaListShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeMetricMetaListShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeMetricMetaListShrinkRequest) SetLabelsShrink(v string) *DescribeMetricMetaListShrinkRequest {
	s.LabelsShrink = &v
	return s
}

func (s *DescribeMetricMetaListShrinkRequest) SetMetaFormat(v string) *DescribeMetricMetaListShrinkRequest {
	s.MetaFormat = &v
	return s
}

func (s *DescribeMetricMetaListShrinkRequest) SetMetricName(v string) *DescribeMetricMetaListShrinkRequest {
	s.MetricName = &v
	return s
}

func (s *DescribeMetricMetaListShrinkRequest) SetNamespace(v string) *DescribeMetricMetaListShrinkRequest {
	s.Namespace = &v
	return s
}

func (s *DescribeMetricMetaListShrinkRequest) SetPageNumber(v int32) *DescribeMetricMetaListShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeMetricMetaListShrinkRequest) SetPageSize(v int32) *DescribeMetricMetaListShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeMetricMetaListShrinkRequest) Validate() error {
	return dara.Validate(s)
}
