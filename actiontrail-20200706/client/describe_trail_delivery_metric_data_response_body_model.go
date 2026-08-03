// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTrailDeliveryMetricDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMetricList(v []*DescribeTrailDeliveryMetricDataResponseBodyMetricList) *DescribeTrailDeliveryMetricDataResponseBody
	GetMetricList() []*DescribeTrailDeliveryMetricDataResponseBodyMetricList
	SetRequestId(v string) *DescribeTrailDeliveryMetricDataResponseBody
	GetRequestId() *string
}

type DescribeTrailDeliveryMetricDataResponseBody struct {
	// A list of data points for the delivery monitoring metric.
	MetricList []*DescribeTrailDeliveryMetricDataResponseBodyMetricList `json:"MetricList,omitempty" xml:"MetricList,omitempty" type:"Repeated"`
	// The unique ID of the request.
	//
	// example:
	//
	// 851038F3-33AB-4C49-97D7-6AB37D35****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeTrailDeliveryMetricDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTrailDeliveryMetricDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTrailDeliveryMetricDataResponseBody) GetMetricList() []*DescribeTrailDeliveryMetricDataResponseBodyMetricList {
	return s.MetricList
}

func (s *DescribeTrailDeliveryMetricDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTrailDeliveryMetricDataResponseBody) SetMetricList(v []*DescribeTrailDeliveryMetricDataResponseBodyMetricList) *DescribeTrailDeliveryMetricDataResponseBody {
	s.MetricList = v
	return s
}

func (s *DescribeTrailDeliveryMetricDataResponseBody) SetRequestId(v string) *DescribeTrailDeliveryMetricDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTrailDeliveryMetricDataResponseBody) Validate() error {
	if s.MetricList != nil {
		for _, item := range s.MetricList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeTrailDeliveryMetricDataResponseBodyMetricList struct {
	// The value of the metric. The meaning of this parameter depends on the value of the `MetricName` parameter in the request.
	//
	// For example, if `MetricName` is set to `delivery_sls_success_count`, `Count` indicates the number of logs successfully delivered to SLS.
	//
	// example:
	//
	// 21
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The Unix timestamp, in milliseconds, that marks the start of the time window for this data point.
	//
	// example:
	//
	// 1775721600000
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s DescribeTrailDeliveryMetricDataResponseBodyMetricList) String() string {
	return dara.Prettify(s)
}

func (s DescribeTrailDeliveryMetricDataResponseBodyMetricList) GoString() string {
	return s.String()
}

func (s *DescribeTrailDeliveryMetricDataResponseBodyMetricList) GetCount() *int64 {
	return s.Count
}

func (s *DescribeTrailDeliveryMetricDataResponseBodyMetricList) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *DescribeTrailDeliveryMetricDataResponseBodyMetricList) SetCount(v int64) *DescribeTrailDeliveryMetricDataResponseBodyMetricList {
	s.Count = &v
	return s
}

func (s *DescribeTrailDeliveryMetricDataResponseBodyMetricList) SetTimestamp(v int64) *DescribeTrailDeliveryMetricDataResponseBodyMetricList {
	s.Timestamp = &v
	return s
}

func (s *DescribeTrailDeliveryMetricDataResponseBodyMetricList) Validate() error {
	return dara.Validate(s)
}
