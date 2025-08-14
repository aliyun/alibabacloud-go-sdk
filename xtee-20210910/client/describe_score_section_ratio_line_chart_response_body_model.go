// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeScoreSectionRatioLineChartResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeScoreSectionRatioLineChartResponseBody
	GetRequestId() *string
	SetResultObject(v *DescribeScoreSectionRatioLineChartResponseBodyResultObject) *DescribeScoreSectionRatioLineChartResponseBody
	GetResultObject() *DescribeScoreSectionRatioLineChartResponseBodyResultObject
}

type DescribeScoreSectionRatioLineChartResponseBody struct {
	// Request ID
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Return object
	ResultObject *DescribeScoreSectionRatioLineChartResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Struct"`
}

func (s DescribeScoreSectionRatioLineChartResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeScoreSectionRatioLineChartResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScoreSectionRatioLineChartResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeScoreSectionRatioLineChartResponseBody) GetResultObject() *DescribeScoreSectionRatioLineChartResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeScoreSectionRatioLineChartResponseBody) SetRequestId(v string) *DescribeScoreSectionRatioLineChartResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeScoreSectionRatioLineChartResponseBody) SetResultObject(v *DescribeScoreSectionRatioLineChartResponseBodyResultObject) *DescribeScoreSectionRatioLineChartResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeScoreSectionRatioLineChartResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeScoreSectionRatioLineChartResponseBodyResultObject struct {
	// Data list
	Series []*DescribeScoreSectionRatioLineChartResponseBodyResultObjectSeries `json:"series,omitempty" xml:"series,omitempty" type:"Repeated"`
	// Details of the xaxis node.
	Xaxis *DescribeScoreSectionRatioLineChartResponseBodyResultObjectXaxis `json:"xaxis,omitempty" xml:"xaxis,omitempty" type:"Struct"`
}

func (s DescribeScoreSectionRatioLineChartResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeScoreSectionRatioLineChartResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeScoreSectionRatioLineChartResponseBodyResultObject) GetSeries() []*DescribeScoreSectionRatioLineChartResponseBodyResultObjectSeries {
	return s.Series
}

func (s *DescribeScoreSectionRatioLineChartResponseBodyResultObject) GetXaxis() *DescribeScoreSectionRatioLineChartResponseBodyResultObjectXaxis {
	return s.Xaxis
}

func (s *DescribeScoreSectionRatioLineChartResponseBodyResultObject) SetSeries(v []*DescribeScoreSectionRatioLineChartResponseBodyResultObjectSeries) *DescribeScoreSectionRatioLineChartResponseBodyResultObject {
	s.Series = v
	return s
}

func (s *DescribeScoreSectionRatioLineChartResponseBodyResultObject) SetXaxis(v *DescribeScoreSectionRatioLineChartResponseBodyResultObjectXaxis) *DescribeScoreSectionRatioLineChartResponseBodyResultObject {
	s.Xaxis = v
	return s
}

func (s *DescribeScoreSectionRatioLineChartResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}

type DescribeScoreSectionRatioLineChartResponseBodyResultObjectSeries struct {
	// Chart data list
	Data []*string `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// Statistical dimension.
	//
	// example:
	//
	// 旁路事件
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s DescribeScoreSectionRatioLineChartResponseBodyResultObjectSeries) String() string {
	return dara.Prettify(s)
}

func (s DescribeScoreSectionRatioLineChartResponseBodyResultObjectSeries) GoString() string {
	return s.String()
}

func (s *DescribeScoreSectionRatioLineChartResponseBodyResultObjectSeries) GetData() []*string {
	return s.Data
}

func (s *DescribeScoreSectionRatioLineChartResponseBodyResultObjectSeries) GetName() *string {
	return s.Name
}

func (s *DescribeScoreSectionRatioLineChartResponseBodyResultObjectSeries) SetData(v []*string) *DescribeScoreSectionRatioLineChartResponseBodyResultObjectSeries {
	s.Data = v
	return s
}

func (s *DescribeScoreSectionRatioLineChartResponseBodyResultObjectSeries) SetName(v string) *DescribeScoreSectionRatioLineChartResponseBodyResultObjectSeries {
	s.Name = &v
	return s
}

func (s *DescribeScoreSectionRatioLineChartResponseBodyResultObjectSeries) Validate() error {
	return dara.Validate(s)
}

type DescribeScoreSectionRatioLineChartResponseBodyResultObjectXaxis struct {
	// Chart data list
	Data []*string `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
}

func (s DescribeScoreSectionRatioLineChartResponseBodyResultObjectXaxis) String() string {
	return dara.Prettify(s)
}

func (s DescribeScoreSectionRatioLineChartResponseBodyResultObjectXaxis) GoString() string {
	return s.String()
}

func (s *DescribeScoreSectionRatioLineChartResponseBodyResultObjectXaxis) GetData() []*string {
	return s.Data
}

func (s *DescribeScoreSectionRatioLineChartResponseBodyResultObjectXaxis) SetData(v []*string) *DescribeScoreSectionRatioLineChartResponseBodyResultObjectXaxis {
	s.Data = v
	return s
}

func (s *DescribeScoreSectionRatioLineChartResponseBodyResultObjectXaxis) Validate() error {
	return dara.Validate(s)
}
