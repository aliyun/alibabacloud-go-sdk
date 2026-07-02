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
	// The monitoring data returned as a JSON-formatted string. For more information, see [Monitoring parameters](https://help.aliyun.com/document_detail/122091.html).
	//
	// > To improve data transfer efficiency, the system returns only monitoring data for metrics with non-zero values. If a metric is not returned, its value is **0**.
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
