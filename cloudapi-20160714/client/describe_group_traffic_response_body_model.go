// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGroupTrafficResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeGroupTrafficResponseBody
	GetRequestId() *string
	SetTrafficPerSecond(v *DescribeGroupTrafficResponseBodyTrafficPerSecond) *DescribeGroupTrafficResponseBody
	GetTrafficPerSecond() *DescribeGroupTrafficResponseBodyTrafficPerSecond
}

type DescribeGroupTrafficResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// EF924FE4-2EDD-4CD3-89EC-34E4708574E7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The traffic information per second.
	TrafficPerSecond *DescribeGroupTrafficResponseBodyTrafficPerSecond `json:"TrafficPerSecond,omitempty" xml:"TrafficPerSecond,omitempty" type:"Struct"`
}

func (s DescribeGroupTrafficResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeGroupTrafficResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGroupTrafficResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeGroupTrafficResponseBody) GetTrafficPerSecond() *DescribeGroupTrafficResponseBodyTrafficPerSecond {
	return s.TrafficPerSecond
}

func (s *DescribeGroupTrafficResponseBody) SetRequestId(v string) *DescribeGroupTrafficResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGroupTrafficResponseBody) SetTrafficPerSecond(v *DescribeGroupTrafficResponseBodyTrafficPerSecond) *DescribeGroupTrafficResponseBody {
	s.TrafficPerSecond = v
	return s
}

func (s *DescribeGroupTrafficResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeGroupTrafficResponseBodyTrafficPerSecond struct {
	MonitorItem []*DescribeGroupTrafficResponseBodyTrafficPerSecondMonitorItem `json:"MonitorItem,omitempty" xml:"MonitorItem,omitempty" type:"Repeated"`
}

func (s DescribeGroupTrafficResponseBodyTrafficPerSecond) String() string {
	return dara.Prettify(s)
}

func (s DescribeGroupTrafficResponseBodyTrafficPerSecond) GoString() string {
	return s.String()
}

func (s *DescribeGroupTrafficResponseBodyTrafficPerSecond) GetMonitorItem() []*DescribeGroupTrafficResponseBodyTrafficPerSecondMonitorItem {
	return s.MonitorItem
}

func (s *DescribeGroupTrafficResponseBodyTrafficPerSecond) SetMonitorItem(v []*DescribeGroupTrafficResponseBodyTrafficPerSecondMonitorItem) *DescribeGroupTrafficResponseBodyTrafficPerSecond {
	s.MonitorItem = v
	return s
}

func (s *DescribeGroupTrafficResponseBodyTrafficPerSecond) Validate() error {
	return dara.Validate(s)
}

type DescribeGroupTrafficResponseBodyTrafficPerSecondMonitorItem struct {
	// The metric. Valid values:
	//
	// 	- inbound: traffic consumed by requests
	//
	// 	- outbound: traffic consumed by responses
	//
	// example:
	//
	// inbound
	Item *string `json:"Item,omitempty" xml:"Item,omitempty"`
	// The corresponding time. The time follows the ISO 8601 standard and UTC time is used. Format: YYYY-MM-DDThh:mm:ssZ.
	//
	// example:
	//
	// 2023-01-29T01:30:00Z
	ItemTime *string `json:"ItemTime,omitempty" xml:"ItemTime,omitempty"`
	// The traffic volume per second.
	//
	// example:
	//
	// 100.0
	ItemValue *string `json:"ItemValue,omitempty" xml:"ItemValue,omitempty"`
}

func (s DescribeGroupTrafficResponseBodyTrafficPerSecondMonitorItem) String() string {
	return dara.Prettify(s)
}

func (s DescribeGroupTrafficResponseBodyTrafficPerSecondMonitorItem) GoString() string {
	return s.String()
}

func (s *DescribeGroupTrafficResponseBodyTrafficPerSecondMonitorItem) GetItem() *string {
	return s.Item
}

func (s *DescribeGroupTrafficResponseBodyTrafficPerSecondMonitorItem) GetItemTime() *string {
	return s.ItemTime
}

func (s *DescribeGroupTrafficResponseBodyTrafficPerSecondMonitorItem) GetItemValue() *string {
	return s.ItemValue
}

func (s *DescribeGroupTrafficResponseBodyTrafficPerSecondMonitorItem) SetItem(v string) *DescribeGroupTrafficResponseBodyTrafficPerSecondMonitorItem {
	s.Item = &v
	return s
}

func (s *DescribeGroupTrafficResponseBodyTrafficPerSecondMonitorItem) SetItemTime(v string) *DescribeGroupTrafficResponseBodyTrafficPerSecondMonitorItem {
	s.ItemTime = &v
	return s
}

func (s *DescribeGroupTrafficResponseBodyTrafficPerSecondMonitorItem) SetItemValue(v string) *DescribeGroupTrafficResponseBodyTrafficPerSecondMonitorItem {
	s.ItemValue = &v
	return s
}

func (s *DescribeGroupTrafficResponseBodyTrafficPerSecondMonitorItem) Validate() error {
	return dara.Validate(s)
}
