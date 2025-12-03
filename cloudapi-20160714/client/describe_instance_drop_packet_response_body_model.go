// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceDropPacketResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceDropPacket(v *DescribeInstanceDropPacketResponseBodyInstanceDropPacket) *DescribeInstanceDropPacketResponseBody
	GetInstanceDropPacket() *DescribeInstanceDropPacketResponseBodyInstanceDropPacket
	SetRequestId(v string) *DescribeInstanceDropPacketResponseBody
	GetRequestId() *string
}

type DescribeInstanceDropPacketResponseBody struct {
	// The list of dropped packets in the instance.
	InstanceDropPacket *DescribeInstanceDropPacketResponseBodyInstanceDropPacket `json:"InstanceDropPacket,omitempty" xml:"InstanceDropPacket,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BZ001
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeInstanceDropPacketResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceDropPacketResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceDropPacketResponseBody) GetInstanceDropPacket() *DescribeInstanceDropPacketResponseBodyInstanceDropPacket {
	return s.InstanceDropPacket
}

func (s *DescribeInstanceDropPacketResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstanceDropPacketResponseBody) SetInstanceDropPacket(v *DescribeInstanceDropPacketResponseBodyInstanceDropPacket) *DescribeInstanceDropPacketResponseBody {
	s.InstanceDropPacket = v
	return s
}

func (s *DescribeInstanceDropPacketResponseBody) SetRequestId(v string) *DescribeInstanceDropPacketResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceDropPacketResponseBody) Validate() error {
	if s.InstanceDropPacket != nil {
		if err := s.InstanceDropPacket.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeInstanceDropPacketResponseBodyInstanceDropPacket struct {
	MonitorItem []*DescribeInstanceDropPacketResponseBodyInstanceDropPacketMonitorItem `json:"MonitorItem,omitempty" xml:"MonitorItem,omitempty" type:"Repeated"`
}

func (s DescribeInstanceDropPacketResponseBodyInstanceDropPacket) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceDropPacketResponseBodyInstanceDropPacket) GoString() string {
	return s.String()
}

func (s *DescribeInstanceDropPacketResponseBodyInstanceDropPacket) GetMonitorItem() []*DescribeInstanceDropPacketResponseBodyInstanceDropPacketMonitorItem {
	return s.MonitorItem
}

func (s *DescribeInstanceDropPacketResponseBodyInstanceDropPacket) SetMonitorItem(v []*DescribeInstanceDropPacketResponseBodyInstanceDropPacketMonitorItem) *DescribeInstanceDropPacketResponseBodyInstanceDropPacket {
	s.MonitorItem = v
	return s
}

func (s *DescribeInstanceDropPacketResponseBodyInstanceDropPacket) Validate() error {
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

type DescribeInstanceDropPacketResponseBodyInstanceDropPacketMonitorItem struct {
	// The metric. Valid values:
	//
	// 	- InstanceDropPacketRX: the number of inbound packets dropped in the instance per second.
	//
	// 	- InstanceDropPacketTX: the number of outbound packets dropped in the instance per second.
	//
	// example:
	//
	// InstanceDropPacketRX
	Item *string `json:"Item,omitempty" xml:"Item,omitempty"`
	// The monitoring time. The time follows the ISO 8601 standard. Format: YYYY-MM-DDThh:mm:ssZ
	//
	// example:
	//
	// 2022-09-06T04:00:36Z
	ItemTime *string `json:"ItemTime,omitempty" xml:"ItemTime,omitempty"`
	// The number of dropped packets in the instance.
	//
	// example:
	//
	// 0.0
	ItemValue *string `json:"ItemValue,omitempty" xml:"ItemValue,omitempty"`
}

func (s DescribeInstanceDropPacketResponseBodyInstanceDropPacketMonitorItem) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceDropPacketResponseBodyInstanceDropPacketMonitorItem) GoString() string {
	return s.String()
}

func (s *DescribeInstanceDropPacketResponseBodyInstanceDropPacketMonitorItem) GetItem() *string {
	return s.Item
}

func (s *DescribeInstanceDropPacketResponseBodyInstanceDropPacketMonitorItem) GetItemTime() *string {
	return s.ItemTime
}

func (s *DescribeInstanceDropPacketResponseBodyInstanceDropPacketMonitorItem) GetItemValue() *string {
	return s.ItemValue
}

func (s *DescribeInstanceDropPacketResponseBodyInstanceDropPacketMonitorItem) SetItem(v string) *DescribeInstanceDropPacketResponseBodyInstanceDropPacketMonitorItem {
	s.Item = &v
	return s
}

func (s *DescribeInstanceDropPacketResponseBodyInstanceDropPacketMonitorItem) SetItemTime(v string) *DescribeInstanceDropPacketResponseBodyInstanceDropPacketMonitorItem {
	s.ItemTime = &v
	return s
}

func (s *DescribeInstanceDropPacketResponseBodyInstanceDropPacketMonitorItem) SetItemValue(v string) *DescribeInstanceDropPacketResponseBodyInstanceDropPacketMonitorItem {
	s.ItemValue = &v
	return s
}

func (s *DescribeInstanceDropPacketResponseBodyInstanceDropPacketMonitorItem) Validate() error {
	return dara.Validate(s)
}
