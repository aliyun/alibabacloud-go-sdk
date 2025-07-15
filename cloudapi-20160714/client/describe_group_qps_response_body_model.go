// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGroupQpsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGroupQps(v *DescribeGroupQpsResponseBodyGroupQps) *DescribeGroupQpsResponseBody
	GetGroupQps() *DescribeGroupQpsResponseBodyGroupQps
	SetRequestId(v string) *DescribeGroupQpsResponseBody
	GetRequestId() *string
}

type DescribeGroupQpsResponseBody struct {
	// The number of requests directed to the API group.
	GroupQps *DescribeGroupQpsResponseBodyGroupQps `json:"GroupQps,omitempty" xml:"GroupQps,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// D6E46F10-F26C-4AA0-BB69-FE2743D9AE62
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeGroupQpsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeGroupQpsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGroupQpsResponseBody) GetGroupQps() *DescribeGroupQpsResponseBodyGroupQps {
	return s.GroupQps
}

func (s *DescribeGroupQpsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeGroupQpsResponseBody) SetGroupQps(v *DescribeGroupQpsResponseBodyGroupQps) *DescribeGroupQpsResponseBody {
	s.GroupQps = v
	return s
}

func (s *DescribeGroupQpsResponseBody) SetRequestId(v string) *DescribeGroupQpsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGroupQpsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeGroupQpsResponseBodyGroupQps struct {
	MonitorItem []*DescribeGroupQpsResponseBodyGroupQpsMonitorItem `json:"MonitorItem,omitempty" xml:"MonitorItem,omitempty" type:"Repeated"`
}

func (s DescribeGroupQpsResponseBodyGroupQps) String() string {
	return dara.Prettify(s)
}

func (s DescribeGroupQpsResponseBodyGroupQps) GoString() string {
	return s.String()
}

func (s *DescribeGroupQpsResponseBodyGroupQps) GetMonitorItem() []*DescribeGroupQpsResponseBodyGroupQpsMonitorItem {
	return s.MonitorItem
}

func (s *DescribeGroupQpsResponseBodyGroupQps) SetMonitorItem(v []*DescribeGroupQpsResponseBodyGroupQpsMonitorItem) *DescribeGroupQpsResponseBodyGroupQps {
	s.MonitorItem = v
	return s
}

func (s *DescribeGroupQpsResponseBodyGroupQps) Validate() error {
	return dara.Validate(s)
}

type DescribeGroupQpsResponseBodyGroupQpsMonitorItem struct {
	// The point in time.
	//
	// example:
	//
	// 2023-08-02T09:15:00Z
	ItemTime *string `json:"ItemTime,omitempty" xml:"ItemTime,omitempty"`
	// The number of requests at the specified point in time.
	//
	// example:
	//
	// 17
	ItemValue *string `json:"ItemValue,omitempty" xml:"ItemValue,omitempty"`
}

func (s DescribeGroupQpsResponseBodyGroupQpsMonitorItem) String() string {
	return dara.Prettify(s)
}

func (s DescribeGroupQpsResponseBodyGroupQpsMonitorItem) GoString() string {
	return s.String()
}

func (s *DescribeGroupQpsResponseBodyGroupQpsMonitorItem) GetItemTime() *string {
	return s.ItemTime
}

func (s *DescribeGroupQpsResponseBodyGroupQpsMonitorItem) GetItemValue() *string {
	return s.ItemValue
}

func (s *DescribeGroupQpsResponseBodyGroupQpsMonitorItem) SetItemTime(v string) *DescribeGroupQpsResponseBodyGroupQpsMonitorItem {
	s.ItemTime = &v
	return s
}

func (s *DescribeGroupQpsResponseBodyGroupQpsMonitorItem) SetItemValue(v string) *DescribeGroupQpsResponseBodyGroupQpsMonitorItem {
	s.ItemValue = &v
	return s
}

func (s *DescribeGroupQpsResponseBodyGroupQpsMonitorItem) Validate() error {
	return dara.Validate(s)
}
