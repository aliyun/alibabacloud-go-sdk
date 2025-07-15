// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApiQpsDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCallFails(v *DescribeApiQpsDataResponseBodyCallFails) *DescribeApiQpsDataResponseBody
	GetCallFails() *DescribeApiQpsDataResponseBodyCallFails
	SetCallSuccesses(v *DescribeApiQpsDataResponseBodyCallSuccesses) *DescribeApiQpsDataResponseBody
	GetCallSuccesses() *DescribeApiQpsDataResponseBodyCallSuccesses
	SetRequestId(v string) *DescribeApiQpsDataResponseBody
	GetRequestId() *string
}

type DescribeApiQpsDataResponseBody struct {
	// The returned information about failed API calls. It is an array consisting of MonitorItem data.
	CallFails *DescribeApiQpsDataResponseBodyCallFails `json:"CallFails,omitempty" xml:"CallFails,omitempty" type:"Struct"`
	// The returned information about successful API calls. It is an array consisting of MonitorItem data.
	CallSuccesses *DescribeApiQpsDataResponseBodyCallSuccesses `json:"CallSuccesses,omitempty" xml:"CallSuccesses,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BZ001
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeApiQpsDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiQpsDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApiQpsDataResponseBody) GetCallFails() *DescribeApiQpsDataResponseBodyCallFails {
	return s.CallFails
}

func (s *DescribeApiQpsDataResponseBody) GetCallSuccesses() *DescribeApiQpsDataResponseBodyCallSuccesses {
	return s.CallSuccesses
}

func (s *DescribeApiQpsDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeApiQpsDataResponseBody) SetCallFails(v *DescribeApiQpsDataResponseBodyCallFails) *DescribeApiQpsDataResponseBody {
	s.CallFails = v
	return s
}

func (s *DescribeApiQpsDataResponseBody) SetCallSuccesses(v *DescribeApiQpsDataResponseBodyCallSuccesses) *DescribeApiQpsDataResponseBody {
	s.CallSuccesses = v
	return s
}

func (s *DescribeApiQpsDataResponseBody) SetRequestId(v string) *DescribeApiQpsDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApiQpsDataResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeApiQpsDataResponseBodyCallFails struct {
	MonitorItem []*DescribeApiQpsDataResponseBodyCallFailsMonitorItem `json:"MonitorItem,omitempty" xml:"MonitorItem,omitempty" type:"Repeated"`
}

func (s DescribeApiQpsDataResponseBodyCallFails) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiQpsDataResponseBodyCallFails) GoString() string {
	return s.String()
}

func (s *DescribeApiQpsDataResponseBodyCallFails) GetMonitorItem() []*DescribeApiQpsDataResponseBodyCallFailsMonitorItem {
	return s.MonitorItem
}

func (s *DescribeApiQpsDataResponseBodyCallFails) SetMonitorItem(v []*DescribeApiQpsDataResponseBodyCallFailsMonitorItem) *DescribeApiQpsDataResponseBodyCallFails {
	s.MonitorItem = v
	return s
}

func (s *DescribeApiQpsDataResponseBodyCallFails) Validate() error {
	return dara.Validate(s)
}

type DescribeApiQpsDataResponseBodyCallFailsMonitorItem struct {
	// The time of the monitoring metric. The time format follows the ISO 8601 standard and UTC time is used. Format: YYYY-MM-DDThh:mm:ssZ
	//
	// example:
	//
	// 2016-07-28T08:20:00Z
	ItemTime *string `json:"ItemTime,omitempty" xml:"ItemTime,omitempty"`
	// The value corresponding to the monitoring metric.
	//
	// example:
	//
	// 0
	ItemValue *string `json:"ItemValue,omitempty" xml:"ItemValue,omitempty"`
}

func (s DescribeApiQpsDataResponseBodyCallFailsMonitorItem) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiQpsDataResponseBodyCallFailsMonitorItem) GoString() string {
	return s.String()
}

func (s *DescribeApiQpsDataResponseBodyCallFailsMonitorItem) GetItemTime() *string {
	return s.ItemTime
}

func (s *DescribeApiQpsDataResponseBodyCallFailsMonitorItem) GetItemValue() *string {
	return s.ItemValue
}

func (s *DescribeApiQpsDataResponseBodyCallFailsMonitorItem) SetItemTime(v string) *DescribeApiQpsDataResponseBodyCallFailsMonitorItem {
	s.ItemTime = &v
	return s
}

func (s *DescribeApiQpsDataResponseBodyCallFailsMonitorItem) SetItemValue(v string) *DescribeApiQpsDataResponseBodyCallFailsMonitorItem {
	s.ItemValue = &v
	return s
}

func (s *DescribeApiQpsDataResponseBodyCallFailsMonitorItem) Validate() error {
	return dara.Validate(s)
}

type DescribeApiQpsDataResponseBodyCallSuccesses struct {
	MonitorItem []*DescribeApiQpsDataResponseBodyCallSuccessesMonitorItem `json:"MonitorItem,omitempty" xml:"MonitorItem,omitempty" type:"Repeated"`
}

func (s DescribeApiQpsDataResponseBodyCallSuccesses) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiQpsDataResponseBodyCallSuccesses) GoString() string {
	return s.String()
}

func (s *DescribeApiQpsDataResponseBodyCallSuccesses) GetMonitorItem() []*DescribeApiQpsDataResponseBodyCallSuccessesMonitorItem {
	return s.MonitorItem
}

func (s *DescribeApiQpsDataResponseBodyCallSuccesses) SetMonitorItem(v []*DescribeApiQpsDataResponseBodyCallSuccessesMonitorItem) *DescribeApiQpsDataResponseBodyCallSuccesses {
	s.MonitorItem = v
	return s
}

func (s *DescribeApiQpsDataResponseBodyCallSuccesses) Validate() error {
	return dara.Validate(s)
}

type DescribeApiQpsDataResponseBodyCallSuccessesMonitorItem struct {
	// The time of the monitoring metric. The time format follows the ISO 8601 standard and UTC time is used. Format: YYYY-MM-DDThh:mm:ssZ
	//
	// example:
	//
	// 2016-07-28T08:24:00Z
	ItemTime *string `json:"ItemTime,omitempty" xml:"ItemTime,omitempty"`
	// The value corresponding to the monitoring metric.
	//
	// example:
	//
	// 650
	ItemValue *string `json:"ItemValue,omitempty" xml:"ItemValue,omitempty"`
}

func (s DescribeApiQpsDataResponseBodyCallSuccessesMonitorItem) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiQpsDataResponseBodyCallSuccessesMonitorItem) GoString() string {
	return s.String()
}

func (s *DescribeApiQpsDataResponseBodyCallSuccessesMonitorItem) GetItemTime() *string {
	return s.ItemTime
}

func (s *DescribeApiQpsDataResponseBodyCallSuccessesMonitorItem) GetItemValue() *string {
	return s.ItemValue
}

func (s *DescribeApiQpsDataResponseBodyCallSuccessesMonitorItem) SetItemTime(v string) *DescribeApiQpsDataResponseBodyCallSuccessesMonitorItem {
	s.ItemTime = &v
	return s
}

func (s *DescribeApiQpsDataResponseBodyCallSuccessesMonitorItem) SetItemValue(v string) *DescribeApiQpsDataResponseBodyCallSuccessesMonitorItem {
	s.ItemValue = &v
	return s
}

func (s *DescribeApiQpsDataResponseBodyCallSuccessesMonitorItem) Validate() error {
	return dara.Validate(s)
}
