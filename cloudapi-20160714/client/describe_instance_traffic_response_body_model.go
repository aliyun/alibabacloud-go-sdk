// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceTrafficResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceTraffic(v *DescribeInstanceTrafficResponseBodyInstanceTraffic) *DescribeInstanceTrafficResponseBody
	GetInstanceTraffic() *DescribeInstanceTrafficResponseBodyInstanceTraffic
	SetRequestId(v string) *DescribeInstanceTrafficResponseBody
	GetRequestId() *string
}

type DescribeInstanceTrafficResponseBody struct {
	InstanceTraffic *DescribeInstanceTrafficResponseBodyInstanceTraffic `json:"InstanceTraffic,omitempty" xml:"InstanceTraffic,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BZ004
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeInstanceTrafficResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceTrafficResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceTrafficResponseBody) GetInstanceTraffic() *DescribeInstanceTrafficResponseBodyInstanceTraffic {
	return s.InstanceTraffic
}

func (s *DescribeInstanceTrafficResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstanceTrafficResponseBody) SetInstanceTraffic(v *DescribeInstanceTrafficResponseBodyInstanceTraffic) *DescribeInstanceTrafficResponseBody {
	s.InstanceTraffic = v
	return s
}

func (s *DescribeInstanceTrafficResponseBody) SetRequestId(v string) *DescribeInstanceTrafficResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceTrafficResponseBody) Validate() error {
	if s.InstanceTraffic != nil {
		if err := s.InstanceTraffic.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeInstanceTrafficResponseBodyInstanceTraffic struct {
	MonitorItem []*DescribeInstanceTrafficResponseBodyInstanceTrafficMonitorItem `json:"MonitorItem,omitempty" xml:"MonitorItem,omitempty" type:"Repeated"`
}

func (s DescribeInstanceTrafficResponseBodyInstanceTraffic) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceTrafficResponseBodyInstanceTraffic) GoString() string {
	return s.String()
}

func (s *DescribeInstanceTrafficResponseBodyInstanceTraffic) GetMonitorItem() []*DescribeInstanceTrafficResponseBodyInstanceTrafficMonitorItem {
	return s.MonitorItem
}

func (s *DescribeInstanceTrafficResponseBodyInstanceTraffic) SetMonitorItem(v []*DescribeInstanceTrafficResponseBodyInstanceTrafficMonitorItem) *DescribeInstanceTrafficResponseBodyInstanceTraffic {
	s.MonitorItem = v
	return s
}

func (s *DescribeInstanceTrafficResponseBodyInstanceTraffic) Validate() error {
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

type DescribeInstanceTrafficResponseBodyInstanceTrafficMonitorItem struct {
	Item      *string `json:"Item,omitempty" xml:"Item,omitempty"`
	ItemTime  *string `json:"ItemTime,omitempty" xml:"ItemTime,omitempty"`
	ItemValue *string `json:"ItemValue,omitempty" xml:"ItemValue,omitempty"`
}

func (s DescribeInstanceTrafficResponseBodyInstanceTrafficMonitorItem) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceTrafficResponseBodyInstanceTrafficMonitorItem) GoString() string {
	return s.String()
}

func (s *DescribeInstanceTrafficResponseBodyInstanceTrafficMonitorItem) GetItem() *string {
	return s.Item
}

func (s *DescribeInstanceTrafficResponseBodyInstanceTrafficMonitorItem) GetItemTime() *string {
	return s.ItemTime
}

func (s *DescribeInstanceTrafficResponseBodyInstanceTrafficMonitorItem) GetItemValue() *string {
	return s.ItemValue
}

func (s *DescribeInstanceTrafficResponseBodyInstanceTrafficMonitorItem) SetItem(v string) *DescribeInstanceTrafficResponseBodyInstanceTrafficMonitorItem {
	s.Item = &v
	return s
}

func (s *DescribeInstanceTrafficResponseBodyInstanceTrafficMonitorItem) SetItemTime(v string) *DescribeInstanceTrafficResponseBodyInstanceTrafficMonitorItem {
	s.ItemTime = &v
	return s
}

func (s *DescribeInstanceTrafficResponseBodyInstanceTrafficMonitorItem) SetItemValue(v string) *DescribeInstanceTrafficResponseBodyInstanceTrafficMonitorItem {
	s.ItemValue = &v
	return s
}

func (s *DescribeInstanceTrafficResponseBodyInstanceTrafficMonitorItem) Validate() error {
	return dara.Validate(s)
}
