// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMetricDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataList(v []*DescribeMetricDataResponseBodyDataList) *DescribeMetricDataResponseBody
	GetDataList() []*DescribeMetricDataResponseBodyDataList
	SetRequestId(v string) *DescribeMetricDataResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeMetricDataResponseBody
	GetTotalCount() *int32
	SetWarnings(v []*string) *DescribeMetricDataResponseBody
	GetWarnings() []*string
}

type DescribeMetricDataResponseBody struct {
	// Collection of monitoring data for the cloud disk.
	DataList []*DescribeMetricDataResponseBodyDataList `json:"DataList,omitempty" xml:"DataList,omitempty" type:"Repeated"`
	// Request ID.
	//
	// example:
	//
	// 11B55F58-D3A4-4A9B-9596-342420D0****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Total number of data points queried.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// List of warning messages.
	Warnings []*string `json:"Warnings,omitempty" xml:"Warnings,omitempty" type:"Repeated"`
}

func (s DescribeMetricDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetricDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMetricDataResponseBody) GetDataList() []*DescribeMetricDataResponseBodyDataList {
	return s.DataList
}

func (s *DescribeMetricDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeMetricDataResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeMetricDataResponseBody) GetWarnings() []*string {
	return s.Warnings
}

func (s *DescribeMetricDataResponseBody) SetDataList(v []*DescribeMetricDataResponseBodyDataList) *DescribeMetricDataResponseBody {
	s.DataList = v
	return s
}

func (s *DescribeMetricDataResponseBody) SetRequestId(v string) *DescribeMetricDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMetricDataResponseBody) SetTotalCount(v int32) *DescribeMetricDataResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeMetricDataResponseBody) SetWarnings(v []*string) *DescribeMetricDataResponseBody {
	s.Warnings = v
	return s
}

func (s *DescribeMetricDataResponseBody) Validate() error {
	if s.DataList != nil {
		for _, item := range s.DataList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeMetricDataResponseBodyDataList struct {
	// List of monitoring data, consisting of a series of consecutive second-level timestamps and the corresponding metric values at those times.
	//
	// example:
	//
	// {"1699258861": 1,"1699259461": 0}
	Datapoints interface{} `json:"Datapoints,omitempty" xml:"Datapoints,omitempty"`
	// Labels.
	//
	// example:
	//
	// {"DiskId": "d-1234"}
	Labels interface{} `json:"Labels,omitempty" xml:"Labels,omitempty"`
}

func (s DescribeMetricDataResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetricDataResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *DescribeMetricDataResponseBodyDataList) GetDatapoints() interface{} {
	return s.Datapoints
}

func (s *DescribeMetricDataResponseBodyDataList) GetLabels() interface{} {
	return s.Labels
}

func (s *DescribeMetricDataResponseBodyDataList) SetDatapoints(v interface{}) *DescribeMetricDataResponseBodyDataList {
	s.Datapoints = v
	return s
}

func (s *DescribeMetricDataResponseBodyDataList) SetLabels(v interface{}) *DescribeMetricDataResponseBodyDataList {
	s.Labels = v
	return s
}

func (s *DescribeMetricDataResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
