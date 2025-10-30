// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeScoreSectionNumLineChartResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeScoreSectionNumLineChartResponseBody
	GetRequestId() *string
	SetResultObject(v *DescribeScoreSectionNumLineChartResponseBodyResultObject) *DescribeScoreSectionNumLineChartResponseBody
	GetResultObject() *DescribeScoreSectionNumLineChartResponseBodyResultObject
}

type DescribeScoreSectionNumLineChartResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Returned object
	ResultObject *DescribeScoreSectionNumLineChartResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Struct"`
}

func (s DescribeScoreSectionNumLineChartResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeScoreSectionNumLineChartResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScoreSectionNumLineChartResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeScoreSectionNumLineChartResponseBody) GetResultObject() *DescribeScoreSectionNumLineChartResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeScoreSectionNumLineChartResponseBody) SetRequestId(v string) *DescribeScoreSectionNumLineChartResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeScoreSectionNumLineChartResponseBody) SetResultObject(v *DescribeScoreSectionNumLineChartResponseBodyResultObject) *DescribeScoreSectionNumLineChartResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeScoreSectionNumLineChartResponseBody) Validate() error {
	if s.ResultObject != nil {
		if err := s.ResultObject.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeScoreSectionNumLineChartResponseBodyResultObject struct {
	// Data list
	Series []*DescribeScoreSectionNumLineChartResponseBodyResultObjectSeries `json:"series,omitempty" xml:"series,omitempty" type:"Repeated"`
	// Details of xaxis node.
	Xaxis *DescribeScoreSectionNumLineChartResponseBodyResultObjectXaxis `json:"xaxis,omitempty" xml:"xaxis,omitempty" type:"Struct"`
}

func (s DescribeScoreSectionNumLineChartResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeScoreSectionNumLineChartResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeScoreSectionNumLineChartResponseBodyResultObject) GetSeries() []*DescribeScoreSectionNumLineChartResponseBodyResultObjectSeries {
	return s.Series
}

func (s *DescribeScoreSectionNumLineChartResponseBodyResultObject) GetXaxis() *DescribeScoreSectionNumLineChartResponseBodyResultObjectXaxis {
	return s.Xaxis
}

func (s *DescribeScoreSectionNumLineChartResponseBodyResultObject) SetSeries(v []*DescribeScoreSectionNumLineChartResponseBodyResultObjectSeries) *DescribeScoreSectionNumLineChartResponseBodyResultObject {
	s.Series = v
	return s
}

func (s *DescribeScoreSectionNumLineChartResponseBodyResultObject) SetXaxis(v *DescribeScoreSectionNumLineChartResponseBodyResultObjectXaxis) *DescribeScoreSectionNumLineChartResponseBodyResultObject {
	s.Xaxis = v
	return s
}

func (s *DescribeScoreSectionNumLineChartResponseBodyResultObject) Validate() error {
	if s.Series != nil {
		for _, item := range s.Series {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Xaxis != nil {
		if err := s.Xaxis.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeScoreSectionNumLineChartResponseBodyResultObjectSeries struct {
	// List of current category results.
	Data []*string `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// Category name.
	//
	// example:
	//
	// 旁路事件
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s DescribeScoreSectionNumLineChartResponseBodyResultObjectSeries) String() string {
	return dara.Prettify(s)
}

func (s DescribeScoreSectionNumLineChartResponseBodyResultObjectSeries) GoString() string {
	return s.String()
}

func (s *DescribeScoreSectionNumLineChartResponseBodyResultObjectSeries) GetData() []*string {
	return s.Data
}

func (s *DescribeScoreSectionNumLineChartResponseBodyResultObjectSeries) GetName() *string {
	return s.Name
}

func (s *DescribeScoreSectionNumLineChartResponseBodyResultObjectSeries) SetData(v []*string) *DescribeScoreSectionNumLineChartResponseBodyResultObjectSeries {
	s.Data = v
	return s
}

func (s *DescribeScoreSectionNumLineChartResponseBodyResultObjectSeries) SetName(v string) *DescribeScoreSectionNumLineChartResponseBodyResultObjectSeries {
	s.Name = &v
	return s
}

func (s *DescribeScoreSectionNumLineChartResponseBodyResultObjectSeries) Validate() error {
	return dara.Validate(s)
}

type DescribeScoreSectionNumLineChartResponseBodyResultObjectXaxis struct {
	// Data structure.
	Data []*string `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
}

func (s DescribeScoreSectionNumLineChartResponseBodyResultObjectXaxis) String() string {
	return dara.Prettify(s)
}

func (s DescribeScoreSectionNumLineChartResponseBodyResultObjectXaxis) GoString() string {
	return s.String()
}

func (s *DescribeScoreSectionNumLineChartResponseBodyResultObjectXaxis) GetData() []*string {
	return s.Data
}

func (s *DescribeScoreSectionNumLineChartResponseBodyResultObjectXaxis) SetData(v []*string) *DescribeScoreSectionNumLineChartResponseBodyResultObjectXaxis {
	s.Data = v
	return s
}

func (s *DescribeScoreSectionNumLineChartResponseBodyResultObjectXaxis) Validate() error {
	return dara.Validate(s)
}
