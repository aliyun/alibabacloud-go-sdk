// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOssCheckStatResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBarChart(v *GetOssCheckStatResponseBodyBarChart) *GetOssCheckStatResponseBody
	GetBarChart() *GetOssCheckStatResponseBodyBarChart
	SetRequestId(v string) *GetOssCheckStatResponseBody
	GetRequestId() *string
}

type GetOssCheckStatResponseBody struct {
	// The bar chart data.
	BarChart *GetOssCheckStatResponseBodyBarChart `json:"BarChart,omitempty" xml:"BarChart,omitempty" type:"Struct"`
	// The ID assigned by the backend to uniquely identify a request. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetOssCheckStatResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetOssCheckStatResponseBody) GoString() string {
	return s.String()
}

func (s *GetOssCheckStatResponseBody) GetBarChart() *GetOssCheckStatResponseBodyBarChart {
	return s.BarChart
}

func (s *GetOssCheckStatResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetOssCheckStatResponseBody) SetBarChart(v *GetOssCheckStatResponseBodyBarChart) *GetOssCheckStatResponseBody {
	s.BarChart = v
	return s
}

func (s *GetOssCheckStatResponseBody) SetRequestId(v string) *GetOssCheckStatResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOssCheckStatResponseBody) Validate() error {
	if s.BarChart != nil {
		if err := s.BarChart.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetOssCheckStatResponseBodyBarChart struct {
	// The X value of the coordinate point.
	X []*string `json:"X,omitempty" xml:"X,omitempty" type:"Repeated"`
	// The Y value of the coordinate point.
	Y []*GetOssCheckStatResponseBodyBarChartY `json:"Y,omitempty" xml:"Y,omitempty" type:"Repeated"`
}

func (s GetOssCheckStatResponseBodyBarChart) String() string {
	return dara.Prettify(s)
}

func (s GetOssCheckStatResponseBodyBarChart) GoString() string {
	return s.String()
}

func (s *GetOssCheckStatResponseBodyBarChart) GetX() []*string {
	return s.X
}

func (s *GetOssCheckStatResponseBodyBarChart) GetY() []*GetOssCheckStatResponseBodyBarChartY {
	return s.Y
}

func (s *GetOssCheckStatResponseBodyBarChart) SetX(v []*string) *GetOssCheckStatResponseBodyBarChart {
	s.X = v
	return s
}

func (s *GetOssCheckStatResponseBodyBarChart) SetY(v []*GetOssCheckStatResponseBodyBarChartY) *GetOssCheckStatResponseBodyBarChart {
	s.Y = v
	return s
}

func (s *GetOssCheckStatResponseBodyBarChart) Validate() error {
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

type GetOssCheckStatResponseBodyBarChartY struct {
	// The data.
	Data []*int64 `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The name.
	//
	// example:
	//
	// document_detection
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetOssCheckStatResponseBodyBarChartY) String() string {
	return dara.Prettify(s)
}

func (s GetOssCheckStatResponseBodyBarChartY) GoString() string {
	return s.String()
}

func (s *GetOssCheckStatResponseBodyBarChartY) GetData() []*int64 {
	return s.Data
}

func (s *GetOssCheckStatResponseBodyBarChartY) GetName() *string {
	return s.Name
}

func (s *GetOssCheckStatResponseBodyBarChartY) SetData(v []*int64) *GetOssCheckStatResponseBodyBarChartY {
	s.Data = v
	return s
}

func (s *GetOssCheckStatResponseBodyBarChartY) SetName(v string) *GetOssCheckStatResponseBodyBarChartY {
	s.Name = &v
	return s
}

func (s *GetOssCheckStatResponseBodyBarChartY) Validate() error {
	return dara.Validate(s)
}
