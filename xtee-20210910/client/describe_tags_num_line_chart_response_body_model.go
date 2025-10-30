// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTagsNumLineChartResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeTagsNumLineChartResponseBody
	GetRequestId() *string
	SetResultObject(v *DescribeTagsNumLineChartResponseBodyResultObject) *DescribeTagsNumLineChartResponseBody
	GetResultObject() *DescribeTagsNumLineChartResponseBodyResultObject
}

type DescribeTagsNumLineChartResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// AE7E6105-7DEB-5125-9B24-DCBC139F6CD2
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Returned object
	ResultObject *DescribeTagsNumLineChartResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Struct"`
}

func (s DescribeTagsNumLineChartResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTagsNumLineChartResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTagsNumLineChartResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTagsNumLineChartResponseBody) GetResultObject() *DescribeTagsNumLineChartResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeTagsNumLineChartResponseBody) SetRequestId(v string) *DescribeTagsNumLineChartResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTagsNumLineChartResponseBody) SetResultObject(v *DescribeTagsNumLineChartResponseBodyResultObject) *DescribeTagsNumLineChartResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeTagsNumLineChartResponseBody) Validate() error {
	if s.ResultObject != nil {
		if err := s.ResultObject.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeTagsNumLineChartResponseBodyResultObject struct {
	// Data list
	Series []*DescribeTagsNumLineChartResponseBodyResultObjectSeries `json:"series,omitempty" xml:"series,omitempty" type:"Repeated"`
	// xaxis node.
	Xaxis *DescribeTagsNumLineChartResponseBodyResultObjectXaxis `json:"xaxis,omitempty" xml:"xaxis,omitempty" type:"Struct"`
}

func (s DescribeTagsNumLineChartResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeTagsNumLineChartResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeTagsNumLineChartResponseBodyResultObject) GetSeries() []*DescribeTagsNumLineChartResponseBodyResultObjectSeries {
	return s.Series
}

func (s *DescribeTagsNumLineChartResponseBodyResultObject) GetXaxis() *DescribeTagsNumLineChartResponseBodyResultObjectXaxis {
	return s.Xaxis
}

func (s *DescribeTagsNumLineChartResponseBodyResultObject) SetSeries(v []*DescribeTagsNumLineChartResponseBodyResultObjectSeries) *DescribeTagsNumLineChartResponseBodyResultObject {
	s.Series = v
	return s
}

func (s *DescribeTagsNumLineChartResponseBodyResultObject) SetXaxis(v *DescribeTagsNumLineChartResponseBodyResultObjectXaxis) *DescribeTagsNumLineChartResponseBodyResultObject {
	s.Xaxis = v
	return s
}

func (s *DescribeTagsNumLineChartResponseBodyResultObject) Validate() error {
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

type DescribeTagsNumLineChartResponseBodyResultObjectSeries struct {
	// Chart data list
	Data []*string `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// Series name.
	//
	// example:
	//
	// rn101
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s DescribeTagsNumLineChartResponseBodyResultObjectSeries) String() string {
	return dara.Prettify(s)
}

func (s DescribeTagsNumLineChartResponseBodyResultObjectSeries) GoString() string {
	return s.String()
}

func (s *DescribeTagsNumLineChartResponseBodyResultObjectSeries) GetData() []*string {
	return s.Data
}

func (s *DescribeTagsNumLineChartResponseBodyResultObjectSeries) GetName() *string {
	return s.Name
}

func (s *DescribeTagsNumLineChartResponseBodyResultObjectSeries) SetData(v []*string) *DescribeTagsNumLineChartResponseBodyResultObjectSeries {
	s.Data = v
	return s
}

func (s *DescribeTagsNumLineChartResponseBodyResultObjectSeries) SetName(v string) *DescribeTagsNumLineChartResponseBodyResultObjectSeries {
	s.Name = &v
	return s
}

func (s *DescribeTagsNumLineChartResponseBodyResultObjectSeries) Validate() error {
	return dara.Validate(s)
}

type DescribeTagsNumLineChartResponseBodyResultObjectXaxis struct {
	// Chart data list
	Data []*string `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
}

func (s DescribeTagsNumLineChartResponseBodyResultObjectXaxis) String() string {
	return dara.Prettify(s)
}

func (s DescribeTagsNumLineChartResponseBodyResultObjectXaxis) GoString() string {
	return s.String()
}

func (s *DescribeTagsNumLineChartResponseBodyResultObjectXaxis) GetData() []*string {
	return s.Data
}

func (s *DescribeTagsNumLineChartResponseBodyResultObjectXaxis) SetData(v []*string) *DescribeTagsNumLineChartResponseBodyResultObjectXaxis {
	s.Data = v
	return s
}

func (s *DescribeTagsNumLineChartResponseBodyResultObjectXaxis) Validate() error {
	return dara.Validate(s)
}
