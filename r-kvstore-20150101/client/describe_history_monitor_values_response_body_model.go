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
	SetMonitorLegend(v string) *DescribeHistoryMonitorValuesResponseBody
	GetMonitorLegend() *string
	SetRequestId(v string) *DescribeHistoryMonitorValuesResponseBody
	GetRequestId() *string
}

type DescribeHistoryMonitorValuesResponseBody struct {
	// The monitoring information returned in JSON format. For more information, see [Monitoring parameter description](https://help.aliyun.com/document_detail/122091.html).
	//
	// > To improve data transmission efficiency, only non-zero monitoring data is returned. All other monitoring data that is not displayed has a default value of **0**.
	//
	// example:
	//
	// "{\\"2022-11-06T00:00:00Z\\":{\\"memoryUsage\\":\\"6.67\\"},\\"2022-11-06T00:00:05Z\\":{\\"memoryUsage\\":\\"6.67\\"},\\"2022-11-06T00:00:10Z\\":{\\"memoryUsage\\":\\"6.67\\"},\\"2022-11-06T00:00:15Z\\":{\\"memoryUsage\\":\\"6.67\\"},\\"2022-11-06T00:00:20Z\\":{\\"memoryUsage\\":\\"6.67\\"},\\"2022-11-06T00:00:25Z\\":{\\"memoryUsage\\":\\"6.67\\"}}"
	MonitorHistory *string `json:"MonitorHistory,omitempty" xml:"MonitorHistory,omitempty"`
	MonitorLegend  *string `json:"MonitorLegend,omitempty" xml:"MonitorLegend,omitempty"`
	// The request ID.
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

func (s *DescribeHistoryMonitorValuesResponseBody) GetMonitorLegend() *string {
	return s.MonitorLegend
}

func (s *DescribeHistoryMonitorValuesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeHistoryMonitorValuesResponseBody) SetMonitorHistory(v string) *DescribeHistoryMonitorValuesResponseBody {
	s.MonitorHistory = &v
	return s
}

func (s *DescribeHistoryMonitorValuesResponseBody) SetMonitorLegend(v string) *DescribeHistoryMonitorValuesResponseBody {
	s.MonitorLegend = &v
	return s
}

func (s *DescribeHistoryMonitorValuesResponseBody) SetRequestId(v string) *DescribeHistoryMonitorValuesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHistoryMonitorValuesResponseBody) Validate() error {
	return dara.Validate(s)
}
