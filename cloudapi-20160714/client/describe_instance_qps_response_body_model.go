// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceQpsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceQps(v *DescribeInstanceQpsResponseBodyInstanceQps) *DescribeInstanceQpsResponseBody
	GetInstanceQps() *DescribeInstanceQpsResponseBodyInstanceQps
	SetRequestId(v string) *DescribeInstanceQpsResponseBody
	GetRequestId() *string
}

type DescribeInstanceQpsResponseBody struct {
	// The list of requests sent to the APIs in the instance.
	InstanceQps *DescribeInstanceQpsResponseBodyInstanceQps `json:"InstanceQps,omitempty" xml:"InstanceQps,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BZ004
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeInstanceQpsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceQpsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceQpsResponseBody) GetInstanceQps() *DescribeInstanceQpsResponseBodyInstanceQps {
	return s.InstanceQps
}

func (s *DescribeInstanceQpsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstanceQpsResponseBody) SetInstanceQps(v *DescribeInstanceQpsResponseBodyInstanceQps) *DescribeInstanceQpsResponseBody {
	s.InstanceQps = v
	return s
}

func (s *DescribeInstanceQpsResponseBody) SetRequestId(v string) *DescribeInstanceQpsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceQpsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeInstanceQpsResponseBodyInstanceQps struct {
	MonitorItem []*DescribeInstanceQpsResponseBodyInstanceQpsMonitorItem `json:"MonitorItem,omitempty" xml:"MonitorItem,omitempty" type:"Repeated"`
}

func (s DescribeInstanceQpsResponseBodyInstanceQps) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceQpsResponseBodyInstanceQps) GoString() string {
	return s.String()
}

func (s *DescribeInstanceQpsResponseBodyInstanceQps) GetMonitorItem() []*DescribeInstanceQpsResponseBodyInstanceQpsMonitorItem {
	return s.MonitorItem
}

func (s *DescribeInstanceQpsResponseBodyInstanceQps) SetMonitorItem(v []*DescribeInstanceQpsResponseBodyInstanceQpsMonitorItem) *DescribeInstanceQpsResponseBodyInstanceQps {
	s.MonitorItem = v
	return s
}

func (s *DescribeInstanceQpsResponseBodyInstanceQps) Validate() error {
	return dara.Validate(s)
}

type DescribeInstanceQpsResponseBodyInstanceQpsMonitorItem struct {
	// The monitoring time. The time follows the ISO 8601 standard. Format: YYYY-MM-DDThh:mm:ssZ
	//
	// example:
	//
	// 2022-03-29T06:25:00Z
	ItemTime *string `json:"ItemTime,omitempty" xml:"ItemTime,omitempty"`
	// The number of requests sent to the APIs in the instance.
	//
	// example:
	//
	// 500
	ItemValue *string `json:"ItemValue,omitempty" xml:"ItemValue,omitempty"`
}

func (s DescribeInstanceQpsResponseBodyInstanceQpsMonitorItem) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceQpsResponseBodyInstanceQpsMonitorItem) GoString() string {
	return s.String()
}

func (s *DescribeInstanceQpsResponseBodyInstanceQpsMonitorItem) GetItemTime() *string {
	return s.ItemTime
}

func (s *DescribeInstanceQpsResponseBodyInstanceQpsMonitorItem) GetItemValue() *string {
	return s.ItemValue
}

func (s *DescribeInstanceQpsResponseBodyInstanceQpsMonitorItem) SetItemTime(v string) *DescribeInstanceQpsResponseBodyInstanceQpsMonitorItem {
	s.ItemTime = &v
	return s
}

func (s *DescribeInstanceQpsResponseBodyInstanceQpsMonitorItem) SetItemValue(v string) *DescribeInstanceQpsResponseBodyInstanceQpsMonitorItem {
	s.ItemValue = &v
	return s
}

func (s *DescribeInstanceQpsResponseBodyInstanceQpsMonitorItem) Validate() error {
	return dara.Validate(s)
}
