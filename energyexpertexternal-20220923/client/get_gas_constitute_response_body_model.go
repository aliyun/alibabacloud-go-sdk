// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGasConstituteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*GetGasConstituteResponseBodyData) *GetGasConstituteResponseBody
	GetData() []*GetGasConstituteResponseBodyData
	SetRequestId(v string) *GetGasConstituteResponseBody
	GetRequestId() *string
}

type GetGasConstituteResponseBody struct {
	// The data returned.
	Data []*GetGasConstituteResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 83A5A7DD-8974-5769-952E-590A97BEA34E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetGasConstituteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetGasConstituteResponseBody) GoString() string {
	return s.String()
}

func (s *GetGasConstituteResponseBody) GetData() []*GetGasConstituteResponseBodyData {
	return s.Data
}

func (s *GetGasConstituteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetGasConstituteResponseBody) SetData(v []*GetGasConstituteResponseBodyData) *GetGasConstituteResponseBody {
	s.Data = v
	return s
}

func (s *GetGasConstituteResponseBody) SetRequestId(v string) *GetGasConstituteResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetGasConstituteResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetGasConstituteResponseBodyData struct {
	// Carbon emissions
	//
	// example:
	//
	// 3.14
	CarbonEmissionData *float64 `json:"carbonEmissionData,omitempty" xml:"carbonEmissionData,omitempty"`
	// Gas emissions
	//
	// example:
	//
	// 3.14
	GasEmissionData *float64 `json:"gasEmissionData,omitempty" xml:"gasEmissionData,omitempty"`
	// Name of gas
	//
	// example:
	//
	// CO₂
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Proportion of carbon emissions. Example value: 0.5 (50%)
	//
	// example:
	//
	// 0.5
	Ratio *float64 `json:"ratio,omitempty" xml:"ratio,omitempty"`
	// Gas Type
	//
	// example:
	//
	// 1
	Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetGasConstituteResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetGasConstituteResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetGasConstituteResponseBodyData) GetCarbonEmissionData() *float64 {
	return s.CarbonEmissionData
}

func (s *GetGasConstituteResponseBodyData) GetGasEmissionData() *float64 {
	return s.GasEmissionData
}

func (s *GetGasConstituteResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetGasConstituteResponseBodyData) GetRatio() *float64 {
	return s.Ratio
}

func (s *GetGasConstituteResponseBodyData) GetType() *int32 {
	return s.Type
}

func (s *GetGasConstituteResponseBodyData) SetCarbonEmissionData(v float64) *GetGasConstituteResponseBodyData {
	s.CarbonEmissionData = &v
	return s
}

func (s *GetGasConstituteResponseBodyData) SetGasEmissionData(v float64) *GetGasConstituteResponseBodyData {
	s.GasEmissionData = &v
	return s
}

func (s *GetGasConstituteResponseBodyData) SetName(v string) *GetGasConstituteResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetGasConstituteResponseBodyData) SetRatio(v float64) *GetGasConstituteResponseBodyData {
	s.Ratio = &v
	return s
}

func (s *GetGasConstituteResponseBodyData) SetType(v int32) *GetGasConstituteResponseBodyData {
	s.Type = &v
	return s
}

func (s *GetGasConstituteResponseBodyData) Validate() error {
	return dara.Validate(s)
}
