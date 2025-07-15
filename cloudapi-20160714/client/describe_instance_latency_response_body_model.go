// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceLatencyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceLatency(v *DescribeInstanceLatencyResponseBodyInstanceLatency) *DescribeInstanceLatencyResponseBody
	GetInstanceLatency() *DescribeInstanceLatencyResponseBodyInstanceLatency
	SetRequestId(v string) *DescribeInstanceLatencyResponseBody
	GetRequestId() *string
}

type DescribeInstanceLatencyResponseBody struct {
	// The list of average latencies in the instance.
	InstanceLatency *DescribeInstanceLatencyResponseBodyInstanceLatency `json:"InstanceLatency,omitempty" xml:"InstanceLatency,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BZ004
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeInstanceLatencyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceLatencyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceLatencyResponseBody) GetInstanceLatency() *DescribeInstanceLatencyResponseBodyInstanceLatency {
	return s.InstanceLatency
}

func (s *DescribeInstanceLatencyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstanceLatencyResponseBody) SetInstanceLatency(v *DescribeInstanceLatencyResponseBodyInstanceLatency) *DescribeInstanceLatencyResponseBody {
	s.InstanceLatency = v
	return s
}

func (s *DescribeInstanceLatencyResponseBody) SetRequestId(v string) *DescribeInstanceLatencyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceLatencyResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeInstanceLatencyResponseBodyInstanceLatency struct {
	MonitorItem []*DescribeInstanceLatencyResponseBodyInstanceLatencyMonitorItem `json:"MonitorItem,omitempty" xml:"MonitorItem,omitempty" type:"Repeated"`
}

func (s DescribeInstanceLatencyResponseBodyInstanceLatency) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceLatencyResponseBodyInstanceLatency) GoString() string {
	return s.String()
}

func (s *DescribeInstanceLatencyResponseBodyInstanceLatency) GetMonitorItem() []*DescribeInstanceLatencyResponseBodyInstanceLatencyMonitorItem {
	return s.MonitorItem
}

func (s *DescribeInstanceLatencyResponseBodyInstanceLatency) SetMonitorItem(v []*DescribeInstanceLatencyResponseBodyInstanceLatencyMonitorItem) *DescribeInstanceLatencyResponseBodyInstanceLatency {
	s.MonitorItem = v
	return s
}

func (s *DescribeInstanceLatencyResponseBodyInstanceLatency) Validate() error {
	return dara.Validate(s)
}

type DescribeInstanceLatencyResponseBodyInstanceLatencyMonitorItem struct {
	// The metric. Valid values:
	//
	// 	- gatewayLatency API: the processing latency of API Gateway
	//
	// 	- latency: the processing latency of the backend service.
	//
	// example:
	//
	// latency
	Item *string `json:"Item,omitempty" xml:"Item,omitempty"`
	// The monitoring time. The time follows the ISO 8601 standard and UTC time is used. Format: YYYY-MM-DDThh:mm:ssZ
	//
	// example:
	//
	// 2022-09-06T02:05:13Z
	ItemTime *string `json:"ItemTime,omitempty" xml:"ItemTime,omitempty"`
	// The value of the average latency.
	//
	// example:
	//
	// 10
	ItemValue *string `json:"ItemValue,omitempty" xml:"ItemValue,omitempty"`
}

func (s DescribeInstanceLatencyResponseBodyInstanceLatencyMonitorItem) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceLatencyResponseBodyInstanceLatencyMonitorItem) GoString() string {
	return s.String()
}

func (s *DescribeInstanceLatencyResponseBodyInstanceLatencyMonitorItem) GetItem() *string {
	return s.Item
}

func (s *DescribeInstanceLatencyResponseBodyInstanceLatencyMonitorItem) GetItemTime() *string {
	return s.ItemTime
}

func (s *DescribeInstanceLatencyResponseBodyInstanceLatencyMonitorItem) GetItemValue() *string {
	return s.ItemValue
}

func (s *DescribeInstanceLatencyResponseBodyInstanceLatencyMonitorItem) SetItem(v string) *DescribeInstanceLatencyResponseBodyInstanceLatencyMonitorItem {
	s.Item = &v
	return s
}

func (s *DescribeInstanceLatencyResponseBodyInstanceLatencyMonitorItem) SetItemTime(v string) *DescribeInstanceLatencyResponseBodyInstanceLatencyMonitorItem {
	s.ItemTime = &v
	return s
}

func (s *DescribeInstanceLatencyResponseBodyInstanceLatencyMonitorItem) SetItemValue(v string) *DescribeInstanceLatencyResponseBodyInstanceLatencyMonitorItem {
	s.ItemValue = &v
	return s
}

func (s *DescribeInstanceLatencyResponseBodyInstanceLatencyMonitorItem) Validate() error {
	return dara.Validate(s)
}
