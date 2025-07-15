// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceHttpCodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceHttpCode(v *DescribeInstanceHttpCodeResponseBodyInstanceHttpCode) *DescribeInstanceHttpCodeResponseBody
	GetInstanceHttpCode() *DescribeInstanceHttpCodeResponseBodyInstanceHttpCode
	SetRequestId(v string) *DescribeInstanceHttpCodeResponseBody
	GetRequestId() *string
}

type DescribeInstanceHttpCodeResponseBody struct {
	// The HTTP status codes.
	InstanceHttpCode *DescribeInstanceHttpCodeResponseBodyInstanceHttpCode `json:"InstanceHttpCode,omitempty" xml:"InstanceHttpCode,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// AD00F8C0-311B-54A9-ADE2-2436771012DA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeInstanceHttpCodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceHttpCodeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceHttpCodeResponseBody) GetInstanceHttpCode() *DescribeInstanceHttpCodeResponseBodyInstanceHttpCode {
	return s.InstanceHttpCode
}

func (s *DescribeInstanceHttpCodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstanceHttpCodeResponseBody) SetInstanceHttpCode(v *DescribeInstanceHttpCodeResponseBodyInstanceHttpCode) *DescribeInstanceHttpCodeResponseBody {
	s.InstanceHttpCode = v
	return s
}

func (s *DescribeInstanceHttpCodeResponseBody) SetRequestId(v string) *DescribeInstanceHttpCodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceHttpCodeResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeInstanceHttpCodeResponseBodyInstanceHttpCode struct {
	MonitorItem []*DescribeInstanceHttpCodeResponseBodyInstanceHttpCodeMonitorItem `json:"MonitorItem,omitempty" xml:"MonitorItem,omitempty" type:"Repeated"`
}

func (s DescribeInstanceHttpCodeResponseBodyInstanceHttpCode) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceHttpCodeResponseBodyInstanceHttpCode) GoString() string {
	return s.String()
}

func (s *DescribeInstanceHttpCodeResponseBodyInstanceHttpCode) GetMonitorItem() []*DescribeInstanceHttpCodeResponseBodyInstanceHttpCodeMonitorItem {
	return s.MonitorItem
}

func (s *DescribeInstanceHttpCodeResponseBodyInstanceHttpCode) SetMonitorItem(v []*DescribeInstanceHttpCodeResponseBodyInstanceHttpCodeMonitorItem) *DescribeInstanceHttpCodeResponseBodyInstanceHttpCode {
	s.MonitorItem = v
	return s
}

func (s *DescribeInstanceHttpCodeResponseBodyInstanceHttpCode) Validate() error {
	return dara.Validate(s)
}

type DescribeInstanceHttpCodeResponseBodyInstanceHttpCodeMonitorItem struct {
	// The HTTP status code returned.
	//
	// example:
	//
	// 404
	ItemTime *string `json:"ItemTime,omitempty" xml:"ItemTime,omitempty"`
	// The corresponding value.
	//
	// example:
	//
	// 1
	ItemValue *string `json:"ItemValue,omitempty" xml:"ItemValue,omitempty"`
}

func (s DescribeInstanceHttpCodeResponseBodyInstanceHttpCodeMonitorItem) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceHttpCodeResponseBodyInstanceHttpCodeMonitorItem) GoString() string {
	return s.String()
}

func (s *DescribeInstanceHttpCodeResponseBodyInstanceHttpCodeMonitorItem) GetItemTime() *string {
	return s.ItemTime
}

func (s *DescribeInstanceHttpCodeResponseBodyInstanceHttpCodeMonitorItem) GetItemValue() *string {
	return s.ItemValue
}

func (s *DescribeInstanceHttpCodeResponseBodyInstanceHttpCodeMonitorItem) SetItemTime(v string) *DescribeInstanceHttpCodeResponseBodyInstanceHttpCodeMonitorItem {
	s.ItemTime = &v
	return s
}

func (s *DescribeInstanceHttpCodeResponseBodyInstanceHttpCodeMonitorItem) SetItemValue(v string) *DescribeInstanceHttpCodeResponseBodyInstanceHttpCodeMonitorItem {
	s.ItemValue = &v
	return s
}

func (s *DescribeInstanceHttpCodeResponseBodyInstanceHttpCodeMonitorItem) Validate() error {
	return dara.Validate(s)
}
