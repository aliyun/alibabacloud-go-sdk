// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTagsRatioLineChartResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeTagsRatioLineChartResponseBody
	GetRequestId() *string
	SetResultObject(v *DescribeTagsRatioLineChartResponseBodyResultObject) *DescribeTagsRatioLineChartResponseBody
	GetResultObject() *DescribeTagsRatioLineChartResponseBodyResultObject
}

type DescribeTagsRatioLineChartResponseBody struct {
	// Request ID
	//
	// example:
	//
	// AE7E6105-7DEB-5125-9B24-DCBC139F6CD2
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Return object
	ResultObject *DescribeTagsRatioLineChartResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Struct"`
}

func (s DescribeTagsRatioLineChartResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTagsRatioLineChartResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTagsRatioLineChartResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTagsRatioLineChartResponseBody) GetResultObject() *DescribeTagsRatioLineChartResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeTagsRatioLineChartResponseBody) SetRequestId(v string) *DescribeTagsRatioLineChartResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTagsRatioLineChartResponseBody) SetResultObject(v *DescribeTagsRatioLineChartResponseBodyResultObject) *DescribeTagsRatioLineChartResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeTagsRatioLineChartResponseBody) Validate() error {
	if s.ResultObject != nil {
		if err := s.ResultObject.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeTagsRatioLineChartResponseBodyResultObject struct {
	// Data list
	Series []*DescribeTagsRatioLineChartResponseBodyResultObjectSeries `json:"series,omitempty" xml:"series,omitempty" type:"Repeated"`
	// xaxis node.
	Xaxis *DescribeTagsRatioLineChartResponseBodyResultObjectXaxis `json:"xaxis,omitempty" xml:"xaxis,omitempty" type:"Struct"`
}

func (s DescribeTagsRatioLineChartResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeTagsRatioLineChartResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeTagsRatioLineChartResponseBodyResultObject) GetSeries() []*DescribeTagsRatioLineChartResponseBodyResultObjectSeries {
	return s.Series
}

func (s *DescribeTagsRatioLineChartResponseBodyResultObject) GetXaxis() *DescribeTagsRatioLineChartResponseBodyResultObjectXaxis {
	return s.Xaxis
}

func (s *DescribeTagsRatioLineChartResponseBodyResultObject) SetSeries(v []*DescribeTagsRatioLineChartResponseBodyResultObjectSeries) *DescribeTagsRatioLineChartResponseBodyResultObject {
	s.Series = v
	return s
}

func (s *DescribeTagsRatioLineChartResponseBodyResultObject) SetXaxis(v *DescribeTagsRatioLineChartResponseBodyResultObjectXaxis) *DescribeTagsRatioLineChartResponseBodyResultObject {
	s.Xaxis = v
	return s
}

func (s *DescribeTagsRatioLineChartResponseBodyResultObject) Validate() error {
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

type DescribeTagsRatioLineChartResponseBodyResultObjectSeries struct {
	// Result data.
	Data []*string `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// Series name.
	//
	// example:
	//
	// rn101
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s DescribeTagsRatioLineChartResponseBodyResultObjectSeries) String() string {
	return dara.Prettify(s)
}

func (s DescribeTagsRatioLineChartResponseBodyResultObjectSeries) GoString() string {
	return s.String()
}

func (s *DescribeTagsRatioLineChartResponseBodyResultObjectSeries) GetData() []*string {
	return s.Data
}

func (s *DescribeTagsRatioLineChartResponseBodyResultObjectSeries) GetName() *string {
	return s.Name
}

func (s *DescribeTagsRatioLineChartResponseBodyResultObjectSeries) SetData(v []*string) *DescribeTagsRatioLineChartResponseBodyResultObjectSeries {
	s.Data = v
	return s
}

func (s *DescribeTagsRatioLineChartResponseBodyResultObjectSeries) SetName(v string) *DescribeTagsRatioLineChartResponseBodyResultObjectSeries {
	s.Name = &v
	return s
}

func (s *DescribeTagsRatioLineChartResponseBodyResultObjectSeries) Validate() error {
	return dara.Validate(s)
}

type DescribeTagsRatioLineChartResponseBodyResultObjectXaxis struct {
	// X-axis data
	Data []*string `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
}

func (s DescribeTagsRatioLineChartResponseBodyResultObjectXaxis) String() string {
	return dara.Prettify(s)
}

func (s DescribeTagsRatioLineChartResponseBodyResultObjectXaxis) GoString() string {
	return s.String()
}

func (s *DescribeTagsRatioLineChartResponseBodyResultObjectXaxis) GetData() []*string {
	return s.Data
}

func (s *DescribeTagsRatioLineChartResponseBodyResultObjectXaxis) SetData(v []*string) *DescribeTagsRatioLineChartResponseBodyResultObjectXaxis {
	s.Data = v
	return s
}

func (s *DescribeTagsRatioLineChartResponseBodyResultObjectXaxis) Validate() error {
	return dara.Validate(s)
}
