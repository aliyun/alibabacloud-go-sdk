// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceDropConnectionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceDropConnections(v *DescribeInstanceDropConnectionsResponseBodyInstanceDropConnections) *DescribeInstanceDropConnectionsResponseBody
	GetInstanceDropConnections() *DescribeInstanceDropConnectionsResponseBodyInstanceDropConnections
	SetRequestId(v string) *DescribeInstanceDropConnectionsResponseBody
	GetRequestId() *string
}

type DescribeInstanceDropConnectionsResponseBody struct {
	// The list of dropped connections in the instance.
	InstanceDropConnections *DescribeInstanceDropConnectionsResponseBodyInstanceDropConnections `json:"InstanceDropConnections,omitempty" xml:"InstanceDropConnections,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BA984
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeInstanceDropConnectionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceDropConnectionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceDropConnectionsResponseBody) GetInstanceDropConnections() *DescribeInstanceDropConnectionsResponseBodyInstanceDropConnections {
	return s.InstanceDropConnections
}

func (s *DescribeInstanceDropConnectionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstanceDropConnectionsResponseBody) SetInstanceDropConnections(v *DescribeInstanceDropConnectionsResponseBodyInstanceDropConnections) *DescribeInstanceDropConnectionsResponseBody {
	s.InstanceDropConnections = v
	return s
}

func (s *DescribeInstanceDropConnectionsResponseBody) SetRequestId(v string) *DescribeInstanceDropConnectionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceDropConnectionsResponseBody) Validate() error {
	if s.InstanceDropConnections != nil {
		if err := s.InstanceDropConnections.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeInstanceDropConnectionsResponseBodyInstanceDropConnections struct {
	MonitorItem []*DescribeInstanceDropConnectionsResponseBodyInstanceDropConnectionsMonitorItem `json:"MonitorItem,omitempty" xml:"MonitorItem,omitempty" type:"Repeated"`
}

func (s DescribeInstanceDropConnectionsResponseBodyInstanceDropConnections) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceDropConnectionsResponseBodyInstanceDropConnections) GoString() string {
	return s.String()
}

func (s *DescribeInstanceDropConnectionsResponseBodyInstanceDropConnections) GetMonitorItem() []*DescribeInstanceDropConnectionsResponseBodyInstanceDropConnectionsMonitorItem {
	return s.MonitorItem
}

func (s *DescribeInstanceDropConnectionsResponseBodyInstanceDropConnections) SetMonitorItem(v []*DescribeInstanceDropConnectionsResponseBodyInstanceDropConnectionsMonitorItem) *DescribeInstanceDropConnectionsResponseBodyInstanceDropConnections {
	s.MonitorItem = v
	return s
}

func (s *DescribeInstanceDropConnectionsResponseBodyInstanceDropConnections) Validate() error {
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

type DescribeInstanceDropConnectionsResponseBodyInstanceDropConnectionsMonitorItem struct {
	// The monitoring time. The time follows the ISO 8601 standard. Format: YYYY-MM-DDThh:mm:ssZ
	//
	// example:
	//
	// 2023-01-31T01:11:00Z
	ItemTime *string `json:"ItemTime,omitempty" xml:"ItemTime,omitempty"`
	// The number of dropped packets in the instance.
	//
	// example:
	//
	// 0.0
	ItemValue *string `json:"ItemValue,omitempty" xml:"ItemValue,omitempty"`
}

func (s DescribeInstanceDropConnectionsResponseBodyInstanceDropConnectionsMonitorItem) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceDropConnectionsResponseBodyInstanceDropConnectionsMonitorItem) GoString() string {
	return s.String()
}

func (s *DescribeInstanceDropConnectionsResponseBodyInstanceDropConnectionsMonitorItem) GetItemTime() *string {
	return s.ItemTime
}

func (s *DescribeInstanceDropConnectionsResponseBodyInstanceDropConnectionsMonitorItem) GetItemValue() *string {
	return s.ItemValue
}

func (s *DescribeInstanceDropConnectionsResponseBodyInstanceDropConnectionsMonitorItem) SetItemTime(v string) *DescribeInstanceDropConnectionsResponseBodyInstanceDropConnectionsMonitorItem {
	s.ItemTime = &v
	return s
}

func (s *DescribeInstanceDropConnectionsResponseBodyInstanceDropConnectionsMonitorItem) SetItemValue(v string) *DescribeInstanceDropConnectionsResponseBodyInstanceDropConnectionsMonitorItem {
	s.ItemValue = &v
	return s
}

func (s *DescribeInstanceDropConnectionsResponseBodyInstanceDropConnectionsMonitorItem) Validate() error {
	return dara.Validate(s)
}
