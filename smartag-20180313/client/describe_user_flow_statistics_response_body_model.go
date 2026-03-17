// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserFlowStatisticsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeUserFlowStatisticsResponseBody
	GetRequestId() *string
	SetSagStatistics(v *DescribeUserFlowStatisticsResponseBodySagStatistics) *DescribeUserFlowStatisticsResponseBody
	GetSagStatistics() *DescribeUserFlowStatisticsResponseBodySagStatistics
}

type DescribeUserFlowStatisticsResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 9552AD68-18EA-4074-B27D-40040FBA9683
	RequestId     *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SagStatistics *DescribeUserFlowStatisticsResponseBodySagStatistics `json:"SagStatistics,omitempty" xml:"SagStatistics,omitempty" type:"Struct"`
}

func (s DescribeUserFlowStatisticsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserFlowStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserFlowStatisticsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeUserFlowStatisticsResponseBody) GetSagStatistics() *DescribeUserFlowStatisticsResponseBodySagStatistics {
	return s.SagStatistics
}

func (s *DescribeUserFlowStatisticsResponseBody) SetRequestId(v string) *DescribeUserFlowStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUserFlowStatisticsResponseBody) SetSagStatistics(v *DescribeUserFlowStatisticsResponseBodySagStatistics) *DescribeUserFlowStatisticsResponseBody {
	s.SagStatistics = v
	return s
}

func (s *DescribeUserFlowStatisticsResponseBody) Validate() error {
	if s.SagStatistics != nil {
		if err := s.SagStatistics.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeUserFlowStatisticsResponseBodySagStatistics struct {
	Statistics []*DescribeUserFlowStatisticsResponseBodySagStatisticsStatistics `json:"Statistics,omitempty" xml:"Statistics,omitempty" type:"Repeated"`
}

func (s DescribeUserFlowStatisticsResponseBodySagStatistics) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserFlowStatisticsResponseBodySagStatistics) GoString() string {
	return s.String()
}

func (s *DescribeUserFlowStatisticsResponseBodySagStatistics) GetStatistics() []*DescribeUserFlowStatisticsResponseBodySagStatisticsStatistics {
	return s.Statistics
}

func (s *DescribeUserFlowStatisticsResponseBodySagStatistics) SetStatistics(v []*DescribeUserFlowStatisticsResponseBodySagStatisticsStatistics) *DescribeUserFlowStatisticsResponseBodySagStatistics {
	s.Statistics = v
	return s
}

func (s *DescribeUserFlowStatisticsResponseBodySagStatistics) Validate() error {
	if s.Statistics != nil {
		for _, item := range s.Statistics {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeUserFlowStatisticsResponseBodySagStatisticsStatistics struct {
	TotalBytes      *string `json:"TotalBytes,omitempty" xml:"TotalBytes,omitempty"`
	TotalLeaveBytes *string `json:"TotalLeaveBytes,omitempty" xml:"TotalLeaveBytes,omitempty"`
	UserName        *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DescribeUserFlowStatisticsResponseBodySagStatisticsStatistics) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserFlowStatisticsResponseBodySagStatisticsStatistics) GoString() string {
	return s.String()
}

func (s *DescribeUserFlowStatisticsResponseBodySagStatisticsStatistics) GetTotalBytes() *string {
	return s.TotalBytes
}

func (s *DescribeUserFlowStatisticsResponseBodySagStatisticsStatistics) GetTotalLeaveBytes() *string {
	return s.TotalLeaveBytes
}

func (s *DescribeUserFlowStatisticsResponseBodySagStatisticsStatistics) GetUserName() *string {
	return s.UserName
}

func (s *DescribeUserFlowStatisticsResponseBodySagStatisticsStatistics) SetTotalBytes(v string) *DescribeUserFlowStatisticsResponseBodySagStatisticsStatistics {
	s.TotalBytes = &v
	return s
}

func (s *DescribeUserFlowStatisticsResponseBodySagStatisticsStatistics) SetTotalLeaveBytes(v string) *DescribeUserFlowStatisticsResponseBodySagStatisticsStatistics {
	s.TotalLeaveBytes = &v
	return s
}

func (s *DescribeUserFlowStatisticsResponseBodySagStatisticsStatistics) SetUserName(v string) *DescribeUserFlowStatisticsResponseBodySagStatisticsStatistics {
	s.UserName = &v
	return s
}

func (s *DescribeUserFlowStatisticsResponseBodySagStatisticsStatistics) Validate() error {
	return dara.Validate(s)
}
