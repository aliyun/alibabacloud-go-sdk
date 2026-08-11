// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAiAppStatsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetAiAppStatsResponseBodyData) *GetAiAppStatsResponseBody
	GetData() *GetAiAppStatsResponseBodyData
	SetRequestId(v string) *GetAiAppStatsResponseBody
	GetRequestId() *string
}

type GetAiAppStatsResponseBody struct {
	// The returned data.
	Data *GetAiAppStatsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID assigned by the backend to uniquely identify a request. Used for troubleshooting.
	//
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAiAppStatsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAiAppStatsResponseBody) GoString() string {
	return s.String()
}

func (s *GetAiAppStatsResponseBody) GetData() *GetAiAppStatsResponseBodyData {
	return s.Data
}

func (s *GetAiAppStatsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAiAppStatsResponseBody) SetData(v *GetAiAppStatsResponseBodyData) *GetAiAppStatsResponseBody {
	s.Data = v
	return s
}

func (s *GetAiAppStatsResponseBody) SetRequestId(v string) *GetAiAppStatsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAiAppStatsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAiAppStatsResponseBodyData struct {
	// The label usage chart.
	LabelStatChart []*GetAiAppStatsResponseBodyDataLabelStatChart `json:"LabelStatChart,omitempty" xml:"LabelStatChart,omitempty" type:"Repeated"`
	// The total count categorized statistics.
	TotalStat map[string]*DataTotalStatValue `json:"TotalStat,omitempty" xml:"TotalStat,omitempty"`
	// The X value of the coordinate point.
	X []*string `json:"X,omitempty" xml:"X,omitempty" type:"Repeated"`
	// The Y value of the coordinate point.
	Y []*GetAiAppStatsResponseBodyDataY `json:"Y,omitempty" xml:"Y,omitempty" type:"Repeated"`
}

func (s GetAiAppStatsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetAiAppStatsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAiAppStatsResponseBodyData) GetLabelStatChart() []*GetAiAppStatsResponseBodyDataLabelStatChart {
	return s.LabelStatChart
}

func (s *GetAiAppStatsResponseBodyData) GetTotalStat() map[string]*DataTotalStatValue {
	return s.TotalStat
}

func (s *GetAiAppStatsResponseBodyData) GetX() []*string {
	return s.X
}

func (s *GetAiAppStatsResponseBodyData) GetY() []*GetAiAppStatsResponseBodyDataY {
	return s.Y
}

func (s *GetAiAppStatsResponseBodyData) SetLabelStatChart(v []*GetAiAppStatsResponseBodyDataLabelStatChart) *GetAiAppStatsResponseBodyData {
	s.LabelStatChart = v
	return s
}

func (s *GetAiAppStatsResponseBodyData) SetTotalStat(v map[string]*DataTotalStatValue) *GetAiAppStatsResponseBodyData {
	s.TotalStat = v
	return s
}

func (s *GetAiAppStatsResponseBodyData) SetX(v []*string) *GetAiAppStatsResponseBodyData {
	s.X = v
	return s
}

func (s *GetAiAppStatsResponseBodyData) SetY(v []*GetAiAppStatsResponseBodyDataY) *GetAiAppStatsResponseBodyData {
	s.Y = v
	return s
}

func (s *GetAiAppStatsResponseBodyData) Validate() error {
	if s.LabelStatChart != nil {
		for _, item := range s.LabelStatChart {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Y != nil {
		for _, item := range s.Y {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetAiAppStatsResponseBodyDataLabelStatChart struct {
	// The tree chart.
	TreeChart []*GetAiAppStatsResponseBodyDataLabelStatChartTreeChart `json:"TreeChart,omitempty" xml:"TreeChart,omitempty" type:"Repeated"`
}

func (s GetAiAppStatsResponseBodyDataLabelStatChart) String() string {
	return dara.Prettify(s)
}

func (s GetAiAppStatsResponseBodyDataLabelStatChart) GoString() string {
	return s.String()
}

func (s *GetAiAppStatsResponseBodyDataLabelStatChart) GetTreeChart() []*GetAiAppStatsResponseBodyDataLabelStatChartTreeChart {
	return s.TreeChart
}

func (s *GetAiAppStatsResponseBodyDataLabelStatChart) SetTreeChart(v []*GetAiAppStatsResponseBodyDataLabelStatChartTreeChart) *GetAiAppStatsResponseBodyDataLabelStatChart {
	s.TreeChart = v
	return s
}

func (s *GetAiAppStatsResponseBodyDataLabelStatChart) Validate() error {
	if s.TreeChart != nil {
		for _, item := range s.TreeChart {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetAiAppStatsResponseBodyDataLabelStatChartTreeChart struct {
	// The label description.
	//
	// example:
	//
	// desc
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The label.
	//
	// example:
	//
	// example
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The score.
	//
	// example:
	//
	// 99.91
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetAiAppStatsResponseBodyDataLabelStatChartTreeChart) String() string {
	return dara.Prettify(s)
}

func (s GetAiAppStatsResponseBodyDataLabelStatChartTreeChart) GoString() string {
	return s.String()
}

func (s *GetAiAppStatsResponseBodyDataLabelStatChartTreeChart) GetDescription() *string {
	return s.Description
}

func (s *GetAiAppStatsResponseBodyDataLabelStatChartTreeChart) GetName() *string {
	return s.Name
}

func (s *GetAiAppStatsResponseBodyDataLabelStatChartTreeChart) GetValue() *string {
	return s.Value
}

func (s *GetAiAppStatsResponseBodyDataLabelStatChartTreeChart) SetDescription(v string) *GetAiAppStatsResponseBodyDataLabelStatChartTreeChart {
	s.Description = &v
	return s
}

func (s *GetAiAppStatsResponseBodyDataLabelStatChartTreeChart) SetName(v string) *GetAiAppStatsResponseBodyDataLabelStatChartTreeChart {
	s.Name = &v
	return s
}

func (s *GetAiAppStatsResponseBodyDataLabelStatChartTreeChart) SetValue(v string) *GetAiAppStatsResponseBodyDataLabelStatChartTreeChart {
	s.Value = &v
	return s
}

func (s *GetAiAppStatsResponseBodyDataLabelStatChartTreeChart) Validate() error {
	return dara.Validate(s)
}

type GetAiAppStatsResponseBodyDataY struct {
	// The returned data.
	Data []*int64 `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The name.
	//
	// example:
	//
	// example
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetAiAppStatsResponseBodyDataY) String() string {
	return dara.Prettify(s)
}

func (s GetAiAppStatsResponseBodyDataY) GoString() string {
	return s.String()
}

func (s *GetAiAppStatsResponseBodyDataY) GetData() []*int64 {
	return s.Data
}

func (s *GetAiAppStatsResponseBodyDataY) GetName() *string {
	return s.Name
}

func (s *GetAiAppStatsResponseBodyDataY) SetData(v []*int64) *GetAiAppStatsResponseBodyDataY {
	s.Data = v
	return s
}

func (s *GetAiAppStatsResponseBodyDataY) SetName(v string) *GetAiAppStatsResponseBodyDataY {
	s.Name = &v
	return s
}

func (s *GetAiAppStatsResponseBodyDataY) Validate() error {
	return dara.Validate(s)
}
