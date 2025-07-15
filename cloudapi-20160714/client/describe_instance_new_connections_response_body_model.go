// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceNewConnectionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceNewConnections(v *DescribeInstanceNewConnectionsResponseBodyInstanceNewConnections) *DescribeInstanceNewConnectionsResponseBody
	GetInstanceNewConnections() *DescribeInstanceNewConnectionsResponseBodyInstanceNewConnections
	SetRequestId(v string) *DescribeInstanceNewConnectionsResponseBody
	GetRequestId() *string
}

type DescribeInstanceNewConnectionsResponseBody struct {
	// The list of new connections in the instance.
	InstanceNewConnections *DescribeInstanceNewConnectionsResponseBodyInstanceNewConnections `json:"InstanceNewConnections,omitempty" xml:"InstanceNewConnections,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BZ004
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeInstanceNewConnectionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceNewConnectionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceNewConnectionsResponseBody) GetInstanceNewConnections() *DescribeInstanceNewConnectionsResponseBodyInstanceNewConnections {
	return s.InstanceNewConnections
}

func (s *DescribeInstanceNewConnectionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstanceNewConnectionsResponseBody) SetInstanceNewConnections(v *DescribeInstanceNewConnectionsResponseBodyInstanceNewConnections) *DescribeInstanceNewConnectionsResponseBody {
	s.InstanceNewConnections = v
	return s
}

func (s *DescribeInstanceNewConnectionsResponseBody) SetRequestId(v string) *DescribeInstanceNewConnectionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceNewConnectionsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeInstanceNewConnectionsResponseBodyInstanceNewConnections struct {
	MonitorItem []*DescribeInstanceNewConnectionsResponseBodyInstanceNewConnectionsMonitorItem `json:"MonitorItem,omitempty" xml:"MonitorItem,omitempty" type:"Repeated"`
}

func (s DescribeInstanceNewConnectionsResponseBodyInstanceNewConnections) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceNewConnectionsResponseBodyInstanceNewConnections) GoString() string {
	return s.String()
}

func (s *DescribeInstanceNewConnectionsResponseBodyInstanceNewConnections) GetMonitorItem() []*DescribeInstanceNewConnectionsResponseBodyInstanceNewConnectionsMonitorItem {
	return s.MonitorItem
}

func (s *DescribeInstanceNewConnectionsResponseBodyInstanceNewConnections) SetMonitorItem(v []*DescribeInstanceNewConnectionsResponseBodyInstanceNewConnectionsMonitorItem) *DescribeInstanceNewConnectionsResponseBodyInstanceNewConnections {
	s.MonitorItem = v
	return s
}

func (s *DescribeInstanceNewConnectionsResponseBodyInstanceNewConnections) Validate() error {
	return dara.Validate(s)
}

type DescribeInstanceNewConnectionsResponseBodyInstanceNewConnectionsMonitorItem struct {
	// The monitoring time. The time follows the ISO 8601 standard and UTC time is used. Format: YYYY-MM-DDThh:mm:ssZ
	//
	// example:
	//
	// 2022-10-08T02:08:00Z
	ItemTime *string `json:"ItemTime,omitempty" xml:"ItemTime,omitempty"`
	// The number of new connections in the instance.
	//
	// example:
	//
	// 16
	ItemValue *string `json:"ItemValue,omitempty" xml:"ItemValue,omitempty"`
}

func (s DescribeInstanceNewConnectionsResponseBodyInstanceNewConnectionsMonitorItem) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceNewConnectionsResponseBodyInstanceNewConnectionsMonitorItem) GoString() string {
	return s.String()
}

func (s *DescribeInstanceNewConnectionsResponseBodyInstanceNewConnectionsMonitorItem) GetItemTime() *string {
	return s.ItemTime
}

func (s *DescribeInstanceNewConnectionsResponseBodyInstanceNewConnectionsMonitorItem) GetItemValue() *string {
	return s.ItemValue
}

func (s *DescribeInstanceNewConnectionsResponseBodyInstanceNewConnectionsMonitorItem) SetItemTime(v string) *DescribeInstanceNewConnectionsResponseBodyInstanceNewConnectionsMonitorItem {
	s.ItemTime = &v
	return s
}

func (s *DescribeInstanceNewConnectionsResponseBodyInstanceNewConnectionsMonitorItem) SetItemValue(v string) *DescribeInstanceNewConnectionsResponseBodyInstanceNewConnectionsMonitorItem {
	s.ItemValue = &v
	return s
}

func (s *DescribeInstanceNewConnectionsResponseBodyInstanceNewConnectionsMonitorItem) Validate() error {
	return dara.Validate(s)
}
