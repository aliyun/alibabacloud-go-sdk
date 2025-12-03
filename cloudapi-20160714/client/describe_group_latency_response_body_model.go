// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGroupLatencyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLatencyPacket(v *DescribeGroupLatencyResponseBodyLatencyPacket) *DescribeGroupLatencyResponseBody
	GetLatencyPacket() *DescribeGroupLatencyResponseBodyLatencyPacket
	SetRequestId(v string) *DescribeGroupLatencyResponseBody
	GetRequestId() *string
}

type DescribeGroupLatencyResponseBody struct {
	// The latency information.
	LatencyPacket *DescribeGroupLatencyResponseBodyLatencyPacket `json:"LatencyPacket,omitempty" xml:"LatencyPacket,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 75DC3AB0-421C-5371-8170-86AEABF77AD0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeGroupLatencyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeGroupLatencyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGroupLatencyResponseBody) GetLatencyPacket() *DescribeGroupLatencyResponseBodyLatencyPacket {
	return s.LatencyPacket
}

func (s *DescribeGroupLatencyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeGroupLatencyResponseBody) SetLatencyPacket(v *DescribeGroupLatencyResponseBodyLatencyPacket) *DescribeGroupLatencyResponseBody {
	s.LatencyPacket = v
	return s
}

func (s *DescribeGroupLatencyResponseBody) SetRequestId(v string) *DescribeGroupLatencyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGroupLatencyResponseBody) Validate() error {
	if s.LatencyPacket != nil {
		if err := s.LatencyPacket.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeGroupLatencyResponseBodyLatencyPacket struct {
	MonitorItem []*DescribeGroupLatencyResponseBodyLatencyPacketMonitorItem `json:"MonitorItem,omitempty" xml:"MonitorItem,omitempty" type:"Repeated"`
}

func (s DescribeGroupLatencyResponseBodyLatencyPacket) String() string {
	return dara.Prettify(s)
}

func (s DescribeGroupLatencyResponseBodyLatencyPacket) GoString() string {
	return s.String()
}

func (s *DescribeGroupLatencyResponseBodyLatencyPacket) GetMonitorItem() []*DescribeGroupLatencyResponseBodyLatencyPacketMonitorItem {
	return s.MonitorItem
}

func (s *DescribeGroupLatencyResponseBodyLatencyPacket) SetMonitorItem(v []*DescribeGroupLatencyResponseBodyLatencyPacketMonitorItem) *DescribeGroupLatencyResponseBodyLatencyPacket {
	s.MonitorItem = v
	return s
}

func (s *DescribeGroupLatencyResponseBodyLatencyPacket) Validate() error {
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

type DescribeGroupLatencyResponseBodyLatencyPacketMonitorItem struct {
	// The metric. Valid values:
	//
	// 	- latency: the backend processing latency
	//
	// 	- gatewayLatency: the API Gateway processing latency
	//
	// example:
	//
	// latency
	Item *string `json:"Item,omitempty" xml:"Item,omitempty"`
	// The point in time when the latency data was collected. The format is YYYY-MM-DDThh:mm:ssZ.
	//
	// example:
	//
	// 2023-03-30T16:10:00Z
	ItemTime *string `json:"ItemTime,omitempty" xml:"ItemTime,omitempty"`
	// The latency. Unit: ms.
	//
	// example:
	//
	// 100.0
	ItemValue *string `json:"ItemValue,omitempty" xml:"ItemValue,omitempty"`
}

func (s DescribeGroupLatencyResponseBodyLatencyPacketMonitorItem) String() string {
	return dara.Prettify(s)
}

func (s DescribeGroupLatencyResponseBodyLatencyPacketMonitorItem) GoString() string {
	return s.String()
}

func (s *DescribeGroupLatencyResponseBodyLatencyPacketMonitorItem) GetItem() *string {
	return s.Item
}

func (s *DescribeGroupLatencyResponseBodyLatencyPacketMonitorItem) GetItemTime() *string {
	return s.ItemTime
}

func (s *DescribeGroupLatencyResponseBodyLatencyPacketMonitorItem) GetItemValue() *string {
	return s.ItemValue
}

func (s *DescribeGroupLatencyResponseBodyLatencyPacketMonitorItem) SetItem(v string) *DescribeGroupLatencyResponseBodyLatencyPacketMonitorItem {
	s.Item = &v
	return s
}

func (s *DescribeGroupLatencyResponseBodyLatencyPacketMonitorItem) SetItemTime(v string) *DescribeGroupLatencyResponseBodyLatencyPacketMonitorItem {
	s.ItemTime = &v
	return s
}

func (s *DescribeGroupLatencyResponseBodyLatencyPacketMonitorItem) SetItemValue(v string) *DescribeGroupLatencyResponseBodyLatencyPacketMonitorItem {
	s.ItemValue = &v
	return s
}

func (s *DescribeGroupLatencyResponseBodyLatencyPacketMonitorItem) Validate() error {
	return dara.Validate(s)
}
