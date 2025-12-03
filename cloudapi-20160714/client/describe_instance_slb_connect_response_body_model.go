// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceSlbConnectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceSlbConnect(v *DescribeInstanceSlbConnectResponseBodyInstanceSlbConnect) *DescribeInstanceSlbConnectResponseBody
	GetInstanceSlbConnect() *DescribeInstanceSlbConnectResponseBodyInstanceSlbConnect
	SetRequestId(v string) *DescribeInstanceSlbConnectResponseBody
	GetRequestId() *string
}

type DescribeInstanceSlbConnectResponseBody struct {
	// The list of concurrent connections in the instance.
	InstanceSlbConnect *DescribeInstanceSlbConnectResponseBodyInstanceSlbConnect `json:"InstanceSlbConnect,omitempty" xml:"InstanceSlbConnect,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// E7FE7172-AA75-5880-B6F7-C00893E9BC06
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeInstanceSlbConnectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceSlbConnectResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSlbConnectResponseBody) GetInstanceSlbConnect() *DescribeInstanceSlbConnectResponseBodyInstanceSlbConnect {
	return s.InstanceSlbConnect
}

func (s *DescribeInstanceSlbConnectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstanceSlbConnectResponseBody) SetInstanceSlbConnect(v *DescribeInstanceSlbConnectResponseBodyInstanceSlbConnect) *DescribeInstanceSlbConnectResponseBody {
	s.InstanceSlbConnect = v
	return s
}

func (s *DescribeInstanceSlbConnectResponseBody) SetRequestId(v string) *DescribeInstanceSlbConnectResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceSlbConnectResponseBody) Validate() error {
	if s.InstanceSlbConnect != nil {
		if err := s.InstanceSlbConnect.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeInstanceSlbConnectResponseBodyInstanceSlbConnect struct {
	MonitorItem []*DescribeInstanceSlbConnectResponseBodyInstanceSlbConnectMonitorItem `json:"MonitorItem,omitempty" xml:"MonitorItem,omitempty" type:"Repeated"`
}

func (s DescribeInstanceSlbConnectResponseBodyInstanceSlbConnect) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceSlbConnectResponseBodyInstanceSlbConnect) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSlbConnectResponseBodyInstanceSlbConnect) GetMonitorItem() []*DescribeInstanceSlbConnectResponseBodyInstanceSlbConnectMonitorItem {
	return s.MonitorItem
}

func (s *DescribeInstanceSlbConnectResponseBodyInstanceSlbConnect) SetMonitorItem(v []*DescribeInstanceSlbConnectResponseBodyInstanceSlbConnectMonitorItem) *DescribeInstanceSlbConnectResponseBodyInstanceSlbConnect {
	s.MonitorItem = v
	return s
}

func (s *DescribeInstanceSlbConnectResponseBodyInstanceSlbConnect) Validate() error {
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

type DescribeInstanceSlbConnectResponseBodyInstanceSlbConnectMonitorItem struct {
	// The metric. Valid values:
	//
	// 	- InstanceMaxConnection: the maximum number of connections
	//
	// 	- InstanceInactiveConnection: the number of inactive connections
	//
	// 	- InstanceActiveConnection: the number of active connections
	//
	// example:
	//
	// InstanceActiveConnection
	Item *string `json:"Item,omitempty" xml:"Item,omitempty"`
	// The monitoring time. The time follows the ISO 8601 standard and UTC time is used. Format: YYYY-MM-DDThh:mm:ssZ
	//
	// example:
	//
	// 2022-09-15T15:07:06Z
	ItemTime *string `json:"ItemTime,omitempty" xml:"ItemTime,omitempty"`
	// The number of concurrent connections in the instance.
	//
	// example:
	//
	// 12
	ItemValue *string `json:"ItemValue,omitempty" xml:"ItemValue,omitempty"`
}

func (s DescribeInstanceSlbConnectResponseBodyInstanceSlbConnectMonitorItem) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceSlbConnectResponseBodyInstanceSlbConnectMonitorItem) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSlbConnectResponseBodyInstanceSlbConnectMonitorItem) GetItem() *string {
	return s.Item
}

func (s *DescribeInstanceSlbConnectResponseBodyInstanceSlbConnectMonitorItem) GetItemTime() *string {
	return s.ItemTime
}

func (s *DescribeInstanceSlbConnectResponseBodyInstanceSlbConnectMonitorItem) GetItemValue() *string {
	return s.ItemValue
}

func (s *DescribeInstanceSlbConnectResponseBodyInstanceSlbConnectMonitorItem) SetItem(v string) *DescribeInstanceSlbConnectResponseBodyInstanceSlbConnectMonitorItem {
	s.Item = &v
	return s
}

func (s *DescribeInstanceSlbConnectResponseBodyInstanceSlbConnectMonitorItem) SetItemTime(v string) *DescribeInstanceSlbConnectResponseBodyInstanceSlbConnectMonitorItem {
	s.ItemTime = &v
	return s
}

func (s *DescribeInstanceSlbConnectResponseBodyInstanceSlbConnectMonitorItem) SetItemValue(v string) *DescribeInstanceSlbConnectResponseBodyInstanceSlbConnectMonitorItem {
	s.ItemValue = &v
	return s
}

func (s *DescribeInstanceSlbConnectResponseBodyInstanceSlbConnectMonitorItem) Validate() error {
	return dara.Validate(s)
}
