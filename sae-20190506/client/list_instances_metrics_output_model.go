// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstancesMetricsOutput interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListInstancesMetricsOutput
	GetRequestId() *string
	SetMetricsList(v []*InstanceMetricInfo) *ListInstancesMetricsOutput
	GetMetricsList() []*InstanceMetricInfo
	SetPageNumber(v int32) *ListInstancesMetricsOutput
	GetPageNumber() *int32
	SetPageSize(v int32) *ListInstancesMetricsOutput
	GetPageSize() *int32
	SetTotalCount(v int32) *ListInstancesMetricsOutput
	GetTotalCount() *int32
}

type ListInstancesMetricsOutput struct {
	RequestId   *string               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	MetricsList []*InstanceMetricInfo `json:"metricsList,omitempty" xml:"metricsList,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 100
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// 1234
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListInstancesMetricsOutput) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesMetricsOutput) GoString() string {
	return s.String()
}

func (s *ListInstancesMetricsOutput) GetRequestId() *string {
	return s.RequestId
}

func (s *ListInstancesMetricsOutput) GetMetricsList() []*InstanceMetricInfo {
	return s.MetricsList
}

func (s *ListInstancesMetricsOutput) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListInstancesMetricsOutput) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListInstancesMetricsOutput) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListInstancesMetricsOutput) SetRequestId(v string) *ListInstancesMetricsOutput {
	s.RequestId = &v
	return s
}

func (s *ListInstancesMetricsOutput) SetMetricsList(v []*InstanceMetricInfo) *ListInstancesMetricsOutput {
	s.MetricsList = v
	return s
}

func (s *ListInstancesMetricsOutput) SetPageNumber(v int32) *ListInstancesMetricsOutput {
	s.PageNumber = &v
	return s
}

func (s *ListInstancesMetricsOutput) SetPageSize(v int32) *ListInstancesMetricsOutput {
	s.PageSize = &v
	return s
}

func (s *ListInstancesMetricsOutput) SetTotalCount(v int32) *ListInstancesMetricsOutput {
	s.TotalCount = &v
	return s
}

func (s *ListInstancesMetricsOutput) Validate() error {
	if s.MetricsList != nil {
		for _, item := range s.MetricsList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
