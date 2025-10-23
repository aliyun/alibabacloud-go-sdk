// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGwpBenchmarkListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetGwpBenchmarkListResponseBodyData) *GetGwpBenchmarkListResponseBody
	GetData() *GetGwpBenchmarkListResponseBodyData
	SetRequestId(v string) *GetGwpBenchmarkListResponseBody
	GetRequestId() *string
}

type GetGwpBenchmarkListResponseBody struct {
	// The response parameters.
	Data *GetGwpBenchmarkListResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The ID of the request. The value is unique for each request. This facilitates subsequent troubleshooting.
	//
	// example:
	//
	// A8AEC6D9-A359-5169-BD1A-BD848BA60D65
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetGwpBenchmarkListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetGwpBenchmarkListResponseBody) GoString() string {
	return s.String()
}

func (s *GetGwpBenchmarkListResponseBody) GetData() *GetGwpBenchmarkListResponseBodyData {
	return s.Data
}

func (s *GetGwpBenchmarkListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetGwpBenchmarkListResponseBody) SetData(v *GetGwpBenchmarkListResponseBodyData) *GetGwpBenchmarkListResponseBody {
	s.Data = v
	return s
}

func (s *GetGwpBenchmarkListResponseBody) SetRequestId(v string) *GetGwpBenchmarkListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetGwpBenchmarkListResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetGwpBenchmarkListResponseBodyData struct {
	// Active carbon reduction ranking list.
	Items []*GetGwpBenchmarkListResponseBodyDataItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// unit of emissions. The default value is `kgCO₂e/productUnit`.
	//
	// The `productUnit` is the unit selected for the product. The unit value is changed to `tCO₂e/productUnit` or `gCO₂e/productUnit`. For more information, see the remarks in the quantity column.
	//
	// example:
	//
	// kgCO₂e/kg
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s GetGwpBenchmarkListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetGwpBenchmarkListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetGwpBenchmarkListResponseBodyData) GetItems() []*GetGwpBenchmarkListResponseBodyDataItems {
	return s.Items
}

func (s *GetGwpBenchmarkListResponseBodyData) GetUnit() *string {
	return s.Unit
}

func (s *GetGwpBenchmarkListResponseBodyData) SetItems(v []*GetGwpBenchmarkListResponseBodyDataItems) *GetGwpBenchmarkListResponseBodyData {
	s.Items = v
	return s
}

func (s *GetGwpBenchmarkListResponseBodyData) SetUnit(v string) *GetGwpBenchmarkListResponseBodyData {
	s.Unit = &v
	return s
}

func (s *GetGwpBenchmarkListResponseBodyData) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetGwpBenchmarkListResponseBodyDataItems struct {
	// `activeReduction=benchmarkEmission-carbonEmission` Generally, baseline emissions are greater than inventory emissions. Maintain four decimal places. Unit pertains to a higher-level unit.
	//
	// example:
	//
	// 0.2169
	ActiveReduction *float64 `json:"activeReduction,omitempty" xml:"activeReduction,omitempty"`
	// Benchmark emissions. Maintain four decimal places. Unit pertains to a higher-level unit.
	//
	// example:
	//
	// 0.0108
	BenchmarkEmission *float64 `json:"benchmarkEmission,omitempty" xml:"benchmarkEmission,omitempty"`
	// Benchmark name
	//
	// example:
	//
	// old-energy
	BenchmarkName *string `json:"benchmarkName,omitempty" xml:"benchmarkName,omitempty"`
	// Inventory emissions. Maintain four decimal places. Unit pertains to a higher-level unit.
	//
	// example:
	//
	// -0.2061
	CarbonEmission *float64 `json:"carbonEmission,omitempty" xml:"carbonEmission,omitempty"`
	// name
	//
	// example:
	//
	// new-energy
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Unused temporarily.
	//
	// example:
	//
	// null
	Percent *string `json:"percent,omitempty" xml:"percent,omitempty"`
}

func (s GetGwpBenchmarkListResponseBodyDataItems) String() string {
	return dara.Prettify(s)
}

func (s GetGwpBenchmarkListResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *GetGwpBenchmarkListResponseBodyDataItems) GetActiveReduction() *float64 {
	return s.ActiveReduction
}

func (s *GetGwpBenchmarkListResponseBodyDataItems) GetBenchmarkEmission() *float64 {
	return s.BenchmarkEmission
}

func (s *GetGwpBenchmarkListResponseBodyDataItems) GetBenchmarkName() *string {
	return s.BenchmarkName
}

func (s *GetGwpBenchmarkListResponseBodyDataItems) GetCarbonEmission() *float64 {
	return s.CarbonEmission
}

func (s *GetGwpBenchmarkListResponseBodyDataItems) GetName() *string {
	return s.Name
}

func (s *GetGwpBenchmarkListResponseBodyDataItems) GetPercent() *string {
	return s.Percent
}

func (s *GetGwpBenchmarkListResponseBodyDataItems) SetActiveReduction(v float64) *GetGwpBenchmarkListResponseBodyDataItems {
	s.ActiveReduction = &v
	return s
}

func (s *GetGwpBenchmarkListResponseBodyDataItems) SetBenchmarkEmission(v float64) *GetGwpBenchmarkListResponseBodyDataItems {
	s.BenchmarkEmission = &v
	return s
}

func (s *GetGwpBenchmarkListResponseBodyDataItems) SetBenchmarkName(v string) *GetGwpBenchmarkListResponseBodyDataItems {
	s.BenchmarkName = &v
	return s
}

func (s *GetGwpBenchmarkListResponseBodyDataItems) SetCarbonEmission(v float64) *GetGwpBenchmarkListResponseBodyDataItems {
	s.CarbonEmission = &v
	return s
}

func (s *GetGwpBenchmarkListResponseBodyDataItems) SetName(v string) *GetGwpBenchmarkListResponseBodyDataItems {
	s.Name = &v
	return s
}

func (s *GetGwpBenchmarkListResponseBodyDataItems) SetPercent(v string) *GetGwpBenchmarkListResponseBodyDataItems {
	s.Percent = &v
	return s
}

func (s *GetGwpBenchmarkListResponseBodyDataItems) Validate() error {
	return dara.Validate(s)
}
