// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApiLatencyDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCallLatencys(v *DescribeApiLatencyDataResponseBodyCallLatencys) *DescribeApiLatencyDataResponseBody
	GetCallLatencys() *DescribeApiLatencyDataResponseBodyCallLatencys
	SetRequestId(v string) *DescribeApiLatencyDataResponseBody
	GetRequestId() *string
}

type DescribeApiLatencyDataResponseBody struct {
	// The returned information about API call latency. It is an array consisting of MonitorItem data.
	CallLatencys *DescribeApiLatencyDataResponseBodyCallLatencys `json:"CallLatencys,omitempty" xml:"CallLatencys,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BZ001
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeApiLatencyDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiLatencyDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApiLatencyDataResponseBody) GetCallLatencys() *DescribeApiLatencyDataResponseBodyCallLatencys {
	return s.CallLatencys
}

func (s *DescribeApiLatencyDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeApiLatencyDataResponseBody) SetCallLatencys(v *DescribeApiLatencyDataResponseBodyCallLatencys) *DescribeApiLatencyDataResponseBody {
	s.CallLatencys = v
	return s
}

func (s *DescribeApiLatencyDataResponseBody) SetRequestId(v string) *DescribeApiLatencyDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApiLatencyDataResponseBody) Validate() error {
	if s.CallLatencys != nil {
		if err := s.CallLatencys.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeApiLatencyDataResponseBodyCallLatencys struct {
	MonitorItem []*DescribeApiLatencyDataResponseBodyCallLatencysMonitorItem `json:"MonitorItem,omitempty" xml:"MonitorItem,omitempty" type:"Repeated"`
}

func (s DescribeApiLatencyDataResponseBodyCallLatencys) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiLatencyDataResponseBodyCallLatencys) GoString() string {
	return s.String()
}

func (s *DescribeApiLatencyDataResponseBodyCallLatencys) GetMonitorItem() []*DescribeApiLatencyDataResponseBodyCallLatencysMonitorItem {
	return s.MonitorItem
}

func (s *DescribeApiLatencyDataResponseBodyCallLatencys) SetMonitorItem(v []*DescribeApiLatencyDataResponseBodyCallLatencysMonitorItem) *DescribeApiLatencyDataResponseBodyCallLatencys {
	s.MonitorItem = v
	return s
}

func (s *DescribeApiLatencyDataResponseBodyCallLatencys) Validate() error {
	if s.MonitorItem != nil {
		for _, item := range s.MonitorItem {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeApiLatencyDataResponseBodyCallLatencysMonitorItem struct {
	// The time of the monitoring metric. The time format follows the ISO 8601 standard and UTC time is used. Format: YYYY-MM-DDThh:mm:ssZ
	//
	// example:
	//
	// 2016-07-28T08:20:00Z
	ItemTime *string `json:"ItemTime,omitempty" xml:"ItemTime,omitempty"`
	// The value corresponding to the monitoring metric.
	//
	// example:
	//
	// 15
	ItemValue *string `json:"ItemValue,omitempty" xml:"ItemValue,omitempty"`
}

func (s DescribeApiLatencyDataResponseBodyCallLatencysMonitorItem) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiLatencyDataResponseBodyCallLatencysMonitorItem) GoString() string {
	return s.String()
}

func (s *DescribeApiLatencyDataResponseBodyCallLatencysMonitorItem) GetItemTime() *string {
	return s.ItemTime
}

func (s *DescribeApiLatencyDataResponseBodyCallLatencysMonitorItem) GetItemValue() *string {
	return s.ItemValue
}

func (s *DescribeApiLatencyDataResponseBodyCallLatencysMonitorItem) SetItemTime(v string) *DescribeApiLatencyDataResponseBodyCallLatencysMonitorItem {
	s.ItemTime = &v
	return s
}

func (s *DescribeApiLatencyDataResponseBodyCallLatencysMonitorItem) SetItemValue(v string) *DescribeApiLatencyDataResponseBodyCallLatencysMonitorItem {
	s.ItemValue = &v
	return s
}

func (s *DescribeApiLatencyDataResponseBodyCallLatencysMonitorItem) Validate() error {
	return dara.Validate(s)
}
