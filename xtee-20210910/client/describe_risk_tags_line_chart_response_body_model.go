// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRiskTagsLineChartResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeRiskTagsLineChartResponseBody
	GetCode() *string
	SetHttpStatusCode(v string) *DescribeRiskTagsLineChartResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *DescribeRiskTagsLineChartResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeRiskTagsLineChartResponseBody
	GetRequestId() *string
	SetResultObject(v *DescribeRiskTagsLineChartResponseBodyResultObject) *DescribeRiskTagsLineChartResponseBody
	GetResultObject() *DescribeRiskTagsLineChartResponseBodyResultObject
	SetSuccess(v bool) *DescribeRiskTagsLineChartResponseBody
	GetSuccess() *bool
}

type DescribeRiskTagsLineChartResponseBody struct {
	// Status code. Note: 200 indicates success.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// HTTP status code
	//
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Error message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Returned result information
	ResultObject *DescribeRiskTagsLineChartResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeRiskTagsLineChartResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRiskTagsLineChartResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRiskTagsLineChartResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeRiskTagsLineChartResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *DescribeRiskTagsLineChartResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeRiskTagsLineChartResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRiskTagsLineChartResponseBody) GetResultObject() *DescribeRiskTagsLineChartResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeRiskTagsLineChartResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeRiskTagsLineChartResponseBody) SetCode(v string) *DescribeRiskTagsLineChartResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeRiskTagsLineChartResponseBody) SetHttpStatusCode(v string) *DescribeRiskTagsLineChartResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeRiskTagsLineChartResponseBody) SetMessage(v string) *DescribeRiskTagsLineChartResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeRiskTagsLineChartResponseBody) SetRequestId(v string) *DescribeRiskTagsLineChartResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRiskTagsLineChartResponseBody) SetResultObject(v *DescribeRiskTagsLineChartResponseBodyResultObject) *DescribeRiskTagsLineChartResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeRiskTagsLineChartResponseBody) SetSuccess(v bool) *DescribeRiskTagsLineChartResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeRiskTagsLineChartResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeRiskTagsLineChartResponseBodyResultObject struct {
	// Call percentage, represented as a decimal
	Percent []*float32 `json:"Percent,omitempty" xml:"Percent,omitempty" type:"Repeated"`
	// Chart data
	Series []*DescribeRiskTagsLineChartResponseBodyResultObjectSeries `json:"Series,omitempty" xml:"Series,omitempty" type:"Repeated"`
	// Total number of records.
	Total []*int64 `json:"Total,omitempty" xml:"Total,omitempty" type:"Repeated"`
	// X-axis data of the chart
	Xaxis *DescribeRiskTagsLineChartResponseBodyResultObjectXaxis `json:"Xaxis,omitempty" xml:"Xaxis,omitempty" type:"Struct"`
}

func (s DescribeRiskTagsLineChartResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeRiskTagsLineChartResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeRiskTagsLineChartResponseBodyResultObject) GetPercent() []*float32 {
	return s.Percent
}

func (s *DescribeRiskTagsLineChartResponseBodyResultObject) GetSeries() []*DescribeRiskTagsLineChartResponseBodyResultObjectSeries {
	return s.Series
}

func (s *DescribeRiskTagsLineChartResponseBodyResultObject) GetTotal() []*int64 {
	return s.Total
}

func (s *DescribeRiskTagsLineChartResponseBodyResultObject) GetXaxis() *DescribeRiskTagsLineChartResponseBodyResultObjectXaxis {
	return s.Xaxis
}

func (s *DescribeRiskTagsLineChartResponseBodyResultObject) SetPercent(v []*float32) *DescribeRiskTagsLineChartResponseBodyResultObject {
	s.Percent = v
	return s
}

func (s *DescribeRiskTagsLineChartResponseBodyResultObject) SetSeries(v []*DescribeRiskTagsLineChartResponseBodyResultObjectSeries) *DescribeRiskTagsLineChartResponseBodyResultObject {
	s.Series = v
	return s
}

func (s *DescribeRiskTagsLineChartResponseBodyResultObject) SetTotal(v []*int64) *DescribeRiskTagsLineChartResponseBodyResultObject {
	s.Total = v
	return s
}

func (s *DescribeRiskTagsLineChartResponseBodyResultObject) SetXaxis(v *DescribeRiskTagsLineChartResponseBodyResultObjectXaxis) *DescribeRiskTagsLineChartResponseBodyResultObject {
	s.Xaxis = v
	return s
}

func (s *DescribeRiskTagsLineChartResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}

type DescribeRiskTagsLineChartResponseBodyResultObjectSeries struct {
	// Data
	//
	// example:
	//
	// 10
	Data *float32 `json:"Data,omitempty" xml:"Data,omitempty"`
	// Name
	//
	// example:
	//
	// rm0102
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeRiskTagsLineChartResponseBodyResultObjectSeries) String() string {
	return dara.Prettify(s)
}

func (s DescribeRiskTagsLineChartResponseBodyResultObjectSeries) GoString() string {
	return s.String()
}

func (s *DescribeRiskTagsLineChartResponseBodyResultObjectSeries) GetData() *float32 {
	return s.Data
}

func (s *DescribeRiskTagsLineChartResponseBodyResultObjectSeries) GetName() *string {
	return s.Name
}

func (s *DescribeRiskTagsLineChartResponseBodyResultObjectSeries) SetData(v float32) *DescribeRiskTagsLineChartResponseBodyResultObjectSeries {
	s.Data = &v
	return s
}

func (s *DescribeRiskTagsLineChartResponseBodyResultObjectSeries) SetName(v string) *DescribeRiskTagsLineChartResponseBodyResultObjectSeries {
	s.Name = &v
	return s
}

func (s *DescribeRiskTagsLineChartResponseBodyResultObjectSeries) Validate() error {
	return dara.Validate(s)
}

type DescribeRiskTagsLineChartResponseBodyResultObjectXaxis struct {
	// Data returned by the chart
	Data []*string `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
}

func (s DescribeRiskTagsLineChartResponseBodyResultObjectXaxis) String() string {
	return dara.Prettify(s)
}

func (s DescribeRiskTagsLineChartResponseBodyResultObjectXaxis) GoString() string {
	return s.String()
}

func (s *DescribeRiskTagsLineChartResponseBodyResultObjectXaxis) GetData() []*string {
	return s.Data
}

func (s *DescribeRiskTagsLineChartResponseBodyResultObjectXaxis) SetData(v []*string) *DescribeRiskTagsLineChartResponseBodyResultObjectXaxis {
	s.Data = v
	return s
}

func (s *DescribeRiskTagsLineChartResponseBodyResultObjectXaxis) Validate() error {
	return dara.Validate(s)
}
