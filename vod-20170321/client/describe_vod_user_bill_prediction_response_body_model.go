// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodUserBillPredictionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBillPredictionData(v *DescribeVodUserBillPredictionResponseBodyBillPredictionData) *DescribeVodUserBillPredictionResponseBody
	GetBillPredictionData() *DescribeVodUserBillPredictionResponseBodyBillPredictionData
	SetBillType(v string) *DescribeVodUserBillPredictionResponseBody
	GetBillType() *string
	SetEndTime(v string) *DescribeVodUserBillPredictionResponseBody
	GetEndTime() *string
	SetRequestId(v string) *DescribeVodUserBillPredictionResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeVodUserBillPredictionResponseBody
	GetStartTime() *string
}

type DescribeVodUserBillPredictionResponseBody struct {
	BillPredictionData *DescribeVodUserBillPredictionResponseBodyBillPredictionData `json:"BillPredictionData,omitempty" xml:"BillPredictionData,omitempty" type:"Struct"`
	BillType           *string                                                      `json:"BillType,omitempty" xml:"BillType,omitempty"`
	EndTime            *string                                                      `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RequestId          *string                                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime          *string                                                      `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeVodUserBillPredictionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodUserBillPredictionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVodUserBillPredictionResponseBody) GetBillPredictionData() *DescribeVodUserBillPredictionResponseBodyBillPredictionData {
	return s.BillPredictionData
}

func (s *DescribeVodUserBillPredictionResponseBody) GetBillType() *string {
	return s.BillType
}

func (s *DescribeVodUserBillPredictionResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeVodUserBillPredictionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVodUserBillPredictionResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeVodUserBillPredictionResponseBody) SetBillPredictionData(v *DescribeVodUserBillPredictionResponseBodyBillPredictionData) *DescribeVodUserBillPredictionResponseBody {
	s.BillPredictionData = v
	return s
}

func (s *DescribeVodUserBillPredictionResponseBody) SetBillType(v string) *DescribeVodUserBillPredictionResponseBody {
	s.BillType = &v
	return s
}

func (s *DescribeVodUserBillPredictionResponseBody) SetEndTime(v string) *DescribeVodUserBillPredictionResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeVodUserBillPredictionResponseBody) SetRequestId(v string) *DescribeVodUserBillPredictionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVodUserBillPredictionResponseBody) SetStartTime(v string) *DescribeVodUserBillPredictionResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeVodUserBillPredictionResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeVodUserBillPredictionResponseBodyBillPredictionData struct {
	BillPredictionDataItem []*DescribeVodUserBillPredictionResponseBodyBillPredictionDataBillPredictionDataItem `json:"BillPredictionDataItem,omitempty" xml:"BillPredictionDataItem,omitempty" type:"Repeated"`
}

func (s DescribeVodUserBillPredictionResponseBodyBillPredictionData) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodUserBillPredictionResponseBodyBillPredictionData) GoString() string {
	return s.String()
}

func (s *DescribeVodUserBillPredictionResponseBodyBillPredictionData) GetBillPredictionDataItem() []*DescribeVodUserBillPredictionResponseBodyBillPredictionDataBillPredictionDataItem {
	return s.BillPredictionDataItem
}

func (s *DescribeVodUserBillPredictionResponseBodyBillPredictionData) SetBillPredictionDataItem(v []*DescribeVodUserBillPredictionResponseBodyBillPredictionDataBillPredictionDataItem) *DescribeVodUserBillPredictionResponseBodyBillPredictionData {
	s.BillPredictionDataItem = v
	return s
}

func (s *DescribeVodUserBillPredictionResponseBodyBillPredictionData) Validate() error {
	return dara.Validate(s)
}

type DescribeVodUserBillPredictionResponseBodyBillPredictionDataBillPredictionDataItem struct {
	Area    *string  `json:"Area,omitempty" xml:"Area,omitempty"`
	TimeStp *string  `json:"TimeStp,omitempty" xml:"TimeStp,omitempty"`
	Value   *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeVodUserBillPredictionResponseBodyBillPredictionDataBillPredictionDataItem) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodUserBillPredictionResponseBodyBillPredictionDataBillPredictionDataItem) GoString() string {
	return s.String()
}

func (s *DescribeVodUserBillPredictionResponseBodyBillPredictionDataBillPredictionDataItem) GetArea() *string {
	return s.Area
}

func (s *DescribeVodUserBillPredictionResponseBodyBillPredictionDataBillPredictionDataItem) GetTimeStp() *string {
	return s.TimeStp
}

func (s *DescribeVodUserBillPredictionResponseBodyBillPredictionDataBillPredictionDataItem) GetValue() *float32 {
	return s.Value
}

func (s *DescribeVodUserBillPredictionResponseBodyBillPredictionDataBillPredictionDataItem) SetArea(v string) *DescribeVodUserBillPredictionResponseBodyBillPredictionDataBillPredictionDataItem {
	s.Area = &v
	return s
}

func (s *DescribeVodUserBillPredictionResponseBodyBillPredictionDataBillPredictionDataItem) SetTimeStp(v string) *DescribeVodUserBillPredictionResponseBodyBillPredictionDataBillPredictionDataItem {
	s.TimeStp = &v
	return s
}

func (s *DescribeVodUserBillPredictionResponseBodyBillPredictionDataBillPredictionDataItem) SetValue(v float32) *DescribeVodUserBillPredictionResponseBodyBillPredictionDataBillPredictionDataItem {
	s.Value = &v
	return s
}

func (s *DescribeVodUserBillPredictionResponseBodyBillPredictionDataBillPredictionDataItem) Validate() error {
	return dara.Validate(s)
}
