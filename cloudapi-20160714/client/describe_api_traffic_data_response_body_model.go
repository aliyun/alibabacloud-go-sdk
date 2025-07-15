// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApiTrafficDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCallDownloads(v *DescribeApiTrafficDataResponseBodyCallDownloads) *DescribeApiTrafficDataResponseBody
	GetCallDownloads() *DescribeApiTrafficDataResponseBodyCallDownloads
	SetCallUploads(v *DescribeApiTrafficDataResponseBodyCallUploads) *DescribeApiTrafficDataResponseBody
	GetCallUploads() *DescribeApiTrafficDataResponseBodyCallUploads
	SetRequestId(v string) *DescribeApiTrafficDataResponseBody
	GetRequestId() *string
}

type DescribeApiTrafficDataResponseBody struct {
	// The returned downlink traffic data of API calls. It is an array consisting of MonitorItem data.
	CallDownloads *DescribeApiTrafficDataResponseBodyCallDownloads `json:"CallDownloads,omitempty" xml:"CallDownloads,omitempty" type:"Struct"`
	// The returned uplink traffic data of API calls. It is an array consisting of MonitorItem data.
	CallUploads *DescribeApiTrafficDataResponseBodyCallUploads `json:"CallUploads,omitempty" xml:"CallUploads,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BZ001
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeApiTrafficDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiTrafficDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApiTrafficDataResponseBody) GetCallDownloads() *DescribeApiTrafficDataResponseBodyCallDownloads {
	return s.CallDownloads
}

func (s *DescribeApiTrafficDataResponseBody) GetCallUploads() *DescribeApiTrafficDataResponseBodyCallUploads {
	return s.CallUploads
}

func (s *DescribeApiTrafficDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeApiTrafficDataResponseBody) SetCallDownloads(v *DescribeApiTrafficDataResponseBodyCallDownloads) *DescribeApiTrafficDataResponseBody {
	s.CallDownloads = v
	return s
}

func (s *DescribeApiTrafficDataResponseBody) SetCallUploads(v *DescribeApiTrafficDataResponseBodyCallUploads) *DescribeApiTrafficDataResponseBody {
	s.CallUploads = v
	return s
}

func (s *DescribeApiTrafficDataResponseBody) SetRequestId(v string) *DescribeApiTrafficDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApiTrafficDataResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeApiTrafficDataResponseBodyCallDownloads struct {
	MonitorItem []*DescribeApiTrafficDataResponseBodyCallDownloadsMonitorItem `json:"MonitorItem,omitempty" xml:"MonitorItem,omitempty" type:"Repeated"`
}

func (s DescribeApiTrafficDataResponseBodyCallDownloads) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiTrafficDataResponseBodyCallDownloads) GoString() string {
	return s.String()
}

func (s *DescribeApiTrafficDataResponseBodyCallDownloads) GetMonitorItem() []*DescribeApiTrafficDataResponseBodyCallDownloadsMonitorItem {
	return s.MonitorItem
}

func (s *DescribeApiTrafficDataResponseBodyCallDownloads) SetMonitorItem(v []*DescribeApiTrafficDataResponseBodyCallDownloadsMonitorItem) *DescribeApiTrafficDataResponseBodyCallDownloads {
	s.MonitorItem = v
	return s
}

func (s *DescribeApiTrafficDataResponseBodyCallDownloads) Validate() error {
	return dara.Validate(s)
}

type DescribeApiTrafficDataResponseBodyCallDownloadsMonitorItem struct {
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

func (s DescribeApiTrafficDataResponseBodyCallDownloadsMonitorItem) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiTrafficDataResponseBodyCallDownloadsMonitorItem) GoString() string {
	return s.String()
}

func (s *DescribeApiTrafficDataResponseBodyCallDownloadsMonitorItem) GetItemTime() *string {
	return s.ItemTime
}

func (s *DescribeApiTrafficDataResponseBodyCallDownloadsMonitorItem) GetItemValue() *string {
	return s.ItemValue
}

func (s *DescribeApiTrafficDataResponseBodyCallDownloadsMonitorItem) SetItemTime(v string) *DescribeApiTrafficDataResponseBodyCallDownloadsMonitorItem {
	s.ItemTime = &v
	return s
}

func (s *DescribeApiTrafficDataResponseBodyCallDownloadsMonitorItem) SetItemValue(v string) *DescribeApiTrafficDataResponseBodyCallDownloadsMonitorItem {
	s.ItemValue = &v
	return s
}

func (s *DescribeApiTrafficDataResponseBodyCallDownloadsMonitorItem) Validate() error {
	return dara.Validate(s)
}

type DescribeApiTrafficDataResponseBodyCallUploads struct {
	MonitorItem []*DescribeApiTrafficDataResponseBodyCallUploadsMonitorItem `json:"MonitorItem,omitempty" xml:"MonitorItem,omitempty" type:"Repeated"`
}

func (s DescribeApiTrafficDataResponseBodyCallUploads) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiTrafficDataResponseBodyCallUploads) GoString() string {
	return s.String()
}

func (s *DescribeApiTrafficDataResponseBodyCallUploads) GetMonitorItem() []*DescribeApiTrafficDataResponseBodyCallUploadsMonitorItem {
	return s.MonitorItem
}

func (s *DescribeApiTrafficDataResponseBodyCallUploads) SetMonitorItem(v []*DescribeApiTrafficDataResponseBodyCallUploadsMonitorItem) *DescribeApiTrafficDataResponseBodyCallUploads {
	s.MonitorItem = v
	return s
}

func (s *DescribeApiTrafficDataResponseBodyCallUploads) Validate() error {
	return dara.Validate(s)
}

type DescribeApiTrafficDataResponseBodyCallUploadsMonitorItem struct {
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
	// 670
	ItemValue *string `json:"ItemValue,omitempty" xml:"ItemValue,omitempty"`
}

func (s DescribeApiTrafficDataResponseBodyCallUploadsMonitorItem) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiTrafficDataResponseBodyCallUploadsMonitorItem) GoString() string {
	return s.String()
}

func (s *DescribeApiTrafficDataResponseBodyCallUploadsMonitorItem) GetItemTime() *string {
	return s.ItemTime
}

func (s *DescribeApiTrafficDataResponseBodyCallUploadsMonitorItem) GetItemValue() *string {
	return s.ItemValue
}

func (s *DescribeApiTrafficDataResponseBodyCallUploadsMonitorItem) SetItemTime(v string) *DescribeApiTrafficDataResponseBodyCallUploadsMonitorItem {
	s.ItemTime = &v
	return s
}

func (s *DescribeApiTrafficDataResponseBodyCallUploadsMonitorItem) SetItemValue(v string) *DescribeApiTrafficDataResponseBodyCallUploadsMonitorItem {
	s.ItemValue = &v
	return s
}

func (s *DescribeApiTrafficDataResponseBodyCallUploadsMonitorItem) Validate() error {
	return dara.Validate(s)
}
