// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHistoryMonitorValuesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMonitorHistory(v string) *DescribeHistoryMonitorValuesResponseBody
	GetMonitorHistory() *string
	SetRequestId(v string) *DescribeHistoryMonitorValuesResponseBody
	GetRequestId() *string
}

type DescribeHistoryMonitorValuesResponseBody struct {
	// The monitoring information returned in the JSON format. For more information, see [View performance monitoring data](https://help.aliyun.com/document_detail/122091.html).
	//
	// 	- Only metrics whose values are not 0 are returned. This improves data transmission efficiency. Metrics that are not displayed are represented by the **0*	- default value.
	//
	// 	- The query result is aligned with the data aggregation frequency. If the specified time range to query is less than or equal to 10 minutes and the data is aggregated once every 5 seconds, query results are returned at an interval of 5 seconds. If the specified StartTime value does not coincide with a point in time for data aggregation, the system returns the latest point in time for data aggregation as the first point in time. For example, if you set the StartTime parameter to 2022-01-20T12:01:48Z, the first point in time returned is 2022-01-20T12:01:45Z.
	//
	// example:
	//
	// "{\\"2022-11-06T00:00:00Z\\":{\\"memoryUsage\\":\\"6.67\\"},\\"2022-11-06T00:00:05Z\\":{\\"memoryUsage\\":\\"6.67\\"},\\"2022-11-06T00:00:10Z\\":{\\"memoryUsage\\":\\"6.67\\"},\\"2022-11-06T00:00:15Z\\":{\\"memoryUsage\\":\\"6.67\\"},\\"2022-11-06T00:00:20Z\\":{\\"memoryUsage\\":\\"6.67\\"},\\"2022-11-06T00:00:25Z\\":{\\"memoryUsage\\":\\"6.67\\"}}"
	MonitorHistory *string `json:"MonitorHistory,omitempty" xml:"MonitorHistory,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// F0997EE8-F4C2-4503-9168-85177ED7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeHistoryMonitorValuesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeHistoryMonitorValuesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHistoryMonitorValuesResponseBody) GetMonitorHistory() *string {
	return s.MonitorHistory
}

func (s *DescribeHistoryMonitorValuesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeHistoryMonitorValuesResponseBody) SetMonitorHistory(v string) *DescribeHistoryMonitorValuesResponseBody {
	s.MonitorHistory = &v
	return s
}

func (s *DescribeHistoryMonitorValuesResponseBody) SetRequestId(v string) *DescribeHistoryMonitorValuesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHistoryMonitorValuesResponseBody) Validate() error {
	return dara.Validate(s)
}
