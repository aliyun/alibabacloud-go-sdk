// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstancePacketsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstancePackets(v *DescribeInstancePacketsResponseBodyInstancePackets) *DescribeInstancePacketsResponseBody
	GetInstancePackets() *DescribeInstancePacketsResponseBodyInstancePackets
	SetRequestId(v string) *DescribeInstancePacketsResponseBody
	GetRequestId() *string
}

type DescribeInstancePacketsResponseBody struct {
	// The list of inbound and outbound data packets in the instance.
	InstancePackets *DescribeInstancePacketsResponseBodyInstancePackets `json:"InstancePackets,omitempty" xml:"InstancePackets,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BZ004
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeInstancePacketsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancePacketsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstancePacketsResponseBody) GetInstancePackets() *DescribeInstancePacketsResponseBodyInstancePackets {
	return s.InstancePackets
}

func (s *DescribeInstancePacketsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstancePacketsResponseBody) SetInstancePackets(v *DescribeInstancePacketsResponseBodyInstancePackets) *DescribeInstancePacketsResponseBody {
	s.InstancePackets = v
	return s
}

func (s *DescribeInstancePacketsResponseBody) SetRequestId(v string) *DescribeInstancePacketsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstancePacketsResponseBody) Validate() error {
	if s.InstancePackets != nil {
		if err := s.InstancePackets.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeInstancePacketsResponseBodyInstancePackets struct {
	MonitorItem []*DescribeInstancePacketsResponseBodyInstancePacketsMonitorItem `json:"MonitorItem,omitempty" xml:"MonitorItem,omitempty" type:"Repeated"`
}

func (s DescribeInstancePacketsResponseBodyInstancePackets) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancePacketsResponseBodyInstancePackets) GoString() string {
	return s.String()
}

func (s *DescribeInstancePacketsResponseBodyInstancePackets) GetMonitorItem() []*DescribeInstancePacketsResponseBodyInstancePacketsMonitorItem {
	return s.MonitorItem
}

func (s *DescribeInstancePacketsResponseBodyInstancePackets) SetMonitorItem(v []*DescribeInstancePacketsResponseBodyInstancePacketsMonitorItem) *DescribeInstancePacketsResponseBodyInstancePackets {
	s.MonitorItem = v
	return s
}

func (s *DescribeInstancePacketsResponseBodyInstancePackets) Validate() error {
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

type DescribeInstancePacketsResponseBodyInstancePacketsMonitorItem struct {
	// The metric. Valid values:
	//
	// 	- InstancePacketRX: inbound data packets
	//
	// 	- InstancePacketTX: outbound data packets
	//
	// example:
	//
	// InstancePacketRX
	Item *string `json:"Item,omitempty" xml:"Item,omitempty"`
	// The monitoring time. The time follows the ISO 8601 standard and UTC time is used. Format: YYYY-MM-DDThh:mm:ssZ
	//
	// example:
	//
	// 2022-05-24T10:14:53Z
	ItemTime *string `json:"ItemTime,omitempty" xml:"ItemTime,omitempty"`
	// The number of inbound and outbound data packets in the instance.
	//
	// example:
	//
	// 0
	ItemValue *string `json:"ItemValue,omitempty" xml:"ItemValue,omitempty"`
}

func (s DescribeInstancePacketsResponseBodyInstancePacketsMonitorItem) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancePacketsResponseBodyInstancePacketsMonitorItem) GoString() string {
	return s.String()
}

func (s *DescribeInstancePacketsResponseBodyInstancePacketsMonitorItem) GetItem() *string {
	return s.Item
}

func (s *DescribeInstancePacketsResponseBodyInstancePacketsMonitorItem) GetItemTime() *string {
	return s.ItemTime
}

func (s *DescribeInstancePacketsResponseBodyInstancePacketsMonitorItem) GetItemValue() *string {
	return s.ItemValue
}

func (s *DescribeInstancePacketsResponseBodyInstancePacketsMonitorItem) SetItem(v string) *DescribeInstancePacketsResponseBodyInstancePacketsMonitorItem {
	s.Item = &v
	return s
}

func (s *DescribeInstancePacketsResponseBodyInstancePacketsMonitorItem) SetItemTime(v string) *DescribeInstancePacketsResponseBodyInstancePacketsMonitorItem {
	s.ItemTime = &v
	return s
}

func (s *DescribeInstancePacketsResponseBodyInstancePacketsMonitorItem) SetItemValue(v string) *DescribeInstancePacketsResponseBodyInstancePacketsMonitorItem {
	s.ItemValue = &v
	return s
}

func (s *DescribeInstancePacketsResponseBodyInstancePacketsMonitorItem) Validate() error {
	return dara.Validate(s)
}
