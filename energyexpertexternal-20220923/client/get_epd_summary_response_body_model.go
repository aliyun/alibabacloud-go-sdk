// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEpdSummaryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*GetEpdSummaryResponseBodyData) *GetEpdSummaryResponseBody
	GetData() []*GetEpdSummaryResponseBodyData
	SetRequestId(v string) *GetEpdSummaryResponseBody
	GetRequestId() *string
}

type GetEpdSummaryResponseBody struct {
	// Response parameters
	Data []*GetEpdSummaryResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// The ID of the request. The value is unique for each request. This facilitates subsequent troubleshooting.
	//
	// example:
	//
	// B91B5559-065A-55C3-8D75-DA218EBFD1DC
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetEpdSummaryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetEpdSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *GetEpdSummaryResponseBody) GetData() []*GetEpdSummaryResponseBodyData {
	return s.Data
}

func (s *GetEpdSummaryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetEpdSummaryResponseBody) SetData(v []*GetEpdSummaryResponseBodyData) *GetEpdSummaryResponseBody {
	s.Data = v
	return s
}

func (s *GetEpdSummaryResponseBody) SetRequestId(v string) *GetEpdSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetEpdSummaryResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetEpdSummaryResponseBodyData struct {
	// Emissions. The result is maintained up to 31 decimal places for precise intermediate calculation and subsequently truncated for display. It is advised to pair the emissions figure with the unit of the environmental impact result for a comprehensive understanding.
	//
	// example:
	//
	// 1009.976265540000000000000000000000
	CarbonEmission *float64 `json:"carbonEmission,omitempty" xml:"carbonEmission,omitempty"`
	// The evaluation index adopted for the environmental impact
	//
	// example:
	//
	// GWP100a
	Indicator *string `json:"indicator,omitempty" xml:"indicator,omitempty"`
	// The category key. The environmental impact category. Currently, a maximum of 19 categories are supported. Enumeration refers to [Carbon Footprint Appendices](https://carbon-doc.oss-cn-hangzhou.aliyuncs.com/CarbonFootprintAppendices-en.pdf).
	//
	// example:
	//
	// gwp
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// Calculation method standard
	//
	// example:
	//
	// CML v4.8 2016
	Method *string `json:"method,omitempty" xml:"method,omitempty"`
	// The name of the category.
	//
	// example:
	//
	// climate change
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Category num: the unique serial number of the environmental impact category. A maximum of 19 categories are supported. Enumeration refers to [Carbon Footprint Appendices](https://carbon-doc.oss-cn-hangzhou.aliyuncs.com/CarbonFootprintAppendices-en.pdf).
	//
	// example:
	//
	// 1
	Num *int64 `json:"num,omitempty" xml:"num,omitempty"`
	// Environmental impact result Value Unit.
	//
	// example:
	//
	// kg CO2-Eq
	PreUnit *string `json:"preUnit,omitempty" xml:"preUnit,omitempty"`
	// The data status. 1 indicates that the calculation is accurate, 2 indicates that the default data is used, and 3 indicates that the missing factor uses the value of 0.
	//
	// example:
	//
	// 1
	State *int64 `json:"state,omitempty" xml:"state,omitempty"`
}

func (s GetEpdSummaryResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetEpdSummaryResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetEpdSummaryResponseBodyData) GetCarbonEmission() *float64 {
	return s.CarbonEmission
}

func (s *GetEpdSummaryResponseBodyData) GetIndicator() *string {
	return s.Indicator
}

func (s *GetEpdSummaryResponseBodyData) GetKey() *string {
	return s.Key
}

func (s *GetEpdSummaryResponseBodyData) GetMethod() *string {
	return s.Method
}

func (s *GetEpdSummaryResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetEpdSummaryResponseBodyData) GetNum() *int64 {
	return s.Num
}

func (s *GetEpdSummaryResponseBodyData) GetPreUnit() *string {
	return s.PreUnit
}

func (s *GetEpdSummaryResponseBodyData) GetState() *int64 {
	return s.State
}

func (s *GetEpdSummaryResponseBodyData) SetCarbonEmission(v float64) *GetEpdSummaryResponseBodyData {
	s.CarbonEmission = &v
	return s
}

func (s *GetEpdSummaryResponseBodyData) SetIndicator(v string) *GetEpdSummaryResponseBodyData {
	s.Indicator = &v
	return s
}

func (s *GetEpdSummaryResponseBodyData) SetKey(v string) *GetEpdSummaryResponseBodyData {
	s.Key = &v
	return s
}

func (s *GetEpdSummaryResponseBodyData) SetMethod(v string) *GetEpdSummaryResponseBodyData {
	s.Method = &v
	return s
}

func (s *GetEpdSummaryResponseBodyData) SetName(v string) *GetEpdSummaryResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetEpdSummaryResponseBodyData) SetNum(v int64) *GetEpdSummaryResponseBodyData {
	s.Num = &v
	return s
}

func (s *GetEpdSummaryResponseBodyData) SetPreUnit(v string) *GetEpdSummaryResponseBodyData {
	s.PreUnit = &v
	return s
}

func (s *GetEpdSummaryResponseBodyData) SetState(v int64) *GetEpdSummaryResponseBodyData {
	s.State = &v
	return s
}

func (s *GetEpdSummaryResponseBodyData) Validate() error {
	return dara.Validate(s)
}
