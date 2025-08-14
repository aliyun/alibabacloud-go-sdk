// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeScoreSectionPieChartResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeScoreSectionPieChartResponseBody
	GetRequestId() *string
	SetResultObject(v *DescribeScoreSectionPieChartResponseBodyResultObject) *DescribeScoreSectionPieChartResponseBody
	GetResultObject() *DescribeScoreSectionPieChartResponseBodyResultObject
}

type DescribeScoreSectionPieChartResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// AE7E6105-7DEB-5125-9B24-DCBC139F6CD2
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Return object
	ResultObject *DescribeScoreSectionPieChartResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Struct"`
}

func (s DescribeScoreSectionPieChartResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeScoreSectionPieChartResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScoreSectionPieChartResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeScoreSectionPieChartResponseBody) GetResultObject() *DescribeScoreSectionPieChartResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeScoreSectionPieChartResponseBody) SetRequestId(v string) *DescribeScoreSectionPieChartResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeScoreSectionPieChartResponseBody) SetResultObject(v *DescribeScoreSectionPieChartResponseBodyResultObject) *DescribeScoreSectionPieChartResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeScoreSectionPieChartResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeScoreSectionPieChartResponseBodyResultObject struct {
	// Chart field, default true
	//
	// example:
	//
	// true
	Animation *bool `json:"animation,omitempty" xml:"animation,omitempty"`
	// Belongs to grid.
	Grid *DescribeScoreSectionPieChartResponseBodyResultObjectGrid `json:"grid,omitempty" xml:"grid,omitempty" type:"Struct"`
	// Data list
	Series []*DescribeScoreSectionPieChartResponseBodyResultObjectSeries `json:"series,omitempty" xml:"series,omitempty" type:"Repeated"`
}

func (s DescribeScoreSectionPieChartResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeScoreSectionPieChartResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeScoreSectionPieChartResponseBodyResultObject) GetAnimation() *bool {
	return s.Animation
}

func (s *DescribeScoreSectionPieChartResponseBodyResultObject) GetGrid() *DescribeScoreSectionPieChartResponseBodyResultObjectGrid {
	return s.Grid
}

func (s *DescribeScoreSectionPieChartResponseBodyResultObject) GetSeries() []*DescribeScoreSectionPieChartResponseBodyResultObjectSeries {
	return s.Series
}

func (s *DescribeScoreSectionPieChartResponseBodyResultObject) SetAnimation(v bool) *DescribeScoreSectionPieChartResponseBodyResultObject {
	s.Animation = &v
	return s
}

func (s *DescribeScoreSectionPieChartResponseBodyResultObject) SetGrid(v *DescribeScoreSectionPieChartResponseBodyResultObjectGrid) *DescribeScoreSectionPieChartResponseBodyResultObject {
	s.Grid = v
	return s
}

func (s *DescribeScoreSectionPieChartResponseBodyResultObject) SetSeries(v []*DescribeScoreSectionPieChartResponseBodyResultObjectSeries) *DescribeScoreSectionPieChartResponseBodyResultObject {
	s.Series = v
	return s
}

func (s *DescribeScoreSectionPieChartResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}

type DescribeScoreSectionPieChartResponseBodyResultObjectGrid struct {
	// Chart field, default false
	//
	// example:
	//
	// false
	Show *bool `json:"show,omitempty" xml:"show,omitempty"`
}

func (s DescribeScoreSectionPieChartResponseBodyResultObjectGrid) String() string {
	return dara.Prettify(s)
}

func (s DescribeScoreSectionPieChartResponseBodyResultObjectGrid) GoString() string {
	return s.String()
}

func (s *DescribeScoreSectionPieChartResponseBodyResultObjectGrid) GetShow() *bool {
	return s.Show
}

func (s *DescribeScoreSectionPieChartResponseBodyResultObjectGrid) SetShow(v bool) *DescribeScoreSectionPieChartResponseBodyResultObjectGrid {
	s.Show = &v
	return s
}

func (s *DescribeScoreSectionPieChartResponseBodyResultObjectGrid) Validate() error {
	return dara.Validate(s)
}

type DescribeScoreSectionPieChartResponseBodyResultObjectSeries struct {
	// Chart data list
	Data []*DescribeScoreSectionPieChartResponseBodyResultObjectSeriesData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// Category name.
	//
	// example:
	//
	// 分值区间占比
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Chart field, default false
	//
	// example:
	//
	// false
	RoseType *bool `json:"roseType,omitempty" xml:"roseType,omitempty"`
}

func (s DescribeScoreSectionPieChartResponseBodyResultObjectSeries) String() string {
	return dara.Prettify(s)
}

func (s DescribeScoreSectionPieChartResponseBodyResultObjectSeries) GoString() string {
	return s.String()
}

func (s *DescribeScoreSectionPieChartResponseBodyResultObjectSeries) GetData() []*DescribeScoreSectionPieChartResponseBodyResultObjectSeriesData {
	return s.Data
}

func (s *DescribeScoreSectionPieChartResponseBodyResultObjectSeries) GetName() *string {
	return s.Name
}

func (s *DescribeScoreSectionPieChartResponseBodyResultObjectSeries) GetRoseType() *bool {
	return s.RoseType
}

func (s *DescribeScoreSectionPieChartResponseBodyResultObjectSeries) SetData(v []*DescribeScoreSectionPieChartResponseBodyResultObjectSeriesData) *DescribeScoreSectionPieChartResponseBodyResultObjectSeries {
	s.Data = v
	return s
}

func (s *DescribeScoreSectionPieChartResponseBodyResultObjectSeries) SetName(v string) *DescribeScoreSectionPieChartResponseBodyResultObjectSeries {
	s.Name = &v
	return s
}

func (s *DescribeScoreSectionPieChartResponseBodyResultObjectSeries) SetRoseType(v bool) *DescribeScoreSectionPieChartResponseBodyResultObjectSeries {
	s.RoseType = &v
	return s
}

func (s *DescribeScoreSectionPieChartResponseBodyResultObjectSeries) Validate() error {
	return dara.Validate(s)
}

type DescribeScoreSectionPieChartResponseBodyResultObjectSeriesData struct {
	// Category item name.
	//
	// example:
	//
	// 名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Result value.
	//
	// example:
	//
	// 100
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s DescribeScoreSectionPieChartResponseBodyResultObjectSeriesData) String() string {
	return dara.Prettify(s)
}

func (s DescribeScoreSectionPieChartResponseBodyResultObjectSeriesData) GoString() string {
	return s.String()
}

func (s *DescribeScoreSectionPieChartResponseBodyResultObjectSeriesData) GetName() *string {
	return s.Name
}

func (s *DescribeScoreSectionPieChartResponseBodyResultObjectSeriesData) GetValue() *string {
	return s.Value
}

func (s *DescribeScoreSectionPieChartResponseBodyResultObjectSeriesData) SetName(v string) *DescribeScoreSectionPieChartResponseBodyResultObjectSeriesData {
	s.Name = &v
	return s
}

func (s *DescribeScoreSectionPieChartResponseBodyResultObjectSeriesData) SetValue(v string) *DescribeScoreSectionPieChartResponseBodyResultObjectSeriesData {
	s.Value = &v
	return s
}

func (s *DescribeScoreSectionPieChartResponseBodyResultObjectSeriesData) Validate() error {
	return dara.Validate(s)
}
