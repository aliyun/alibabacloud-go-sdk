// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAreaElecConstituteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetAreaElecConstituteResponseBody
	GetCode() *string
	SetData(v *GetAreaElecConstituteResponseBodyData) *GetAreaElecConstituteResponseBody
	GetData() *GetAreaElecConstituteResponseBodyData
	SetRequestId(v string) *GetAreaElecConstituteResponseBody
	GetRequestId() *string
}

type GetAreaElecConstituteResponseBody struct {
	// The code returned for the request. A value of Success indicates that the request was successful. Other values indicate that the request failed. You can troubleshoot the error by viewing the error message returned.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The returned data.
	Data *GetAreaElecConstituteResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 83A5A7DD-8974-5769-952E-590A97BEA34E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetAreaElecConstituteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAreaElecConstituteResponseBody) GoString() string {
	return s.String()
}

func (s *GetAreaElecConstituteResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetAreaElecConstituteResponseBody) GetData() *GetAreaElecConstituteResponseBodyData {
	return s.Data
}

func (s *GetAreaElecConstituteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAreaElecConstituteResponseBody) SetCode(v string) *GetAreaElecConstituteResponseBody {
	s.Code = &v
	return s
}

func (s *GetAreaElecConstituteResponseBody) SetData(v *GetAreaElecConstituteResponseBodyData) *GetAreaElecConstituteResponseBody {
	s.Data = v
	return s
}

func (s *GetAreaElecConstituteResponseBody) SetRequestId(v string) *GetAreaElecConstituteResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAreaElecConstituteResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAreaElecConstituteResponseBodyData struct {
	// Photoelectric power consumption and carbon emission data of each enterprise.
	Light []*CarbonEmissionElecSummaryItem `json:"light,omitempty" xml:"light,omitempty" type:"Repeated"`
	// Data on nuclear power consumption and carbon emissions at each enterprise.
	Nuclear []*CarbonEmissionElecSummaryItem `json:"nuclear,omitempty" xml:"nuclear,omitempty" type:"Repeated"`
	// Data on renewable electricity consumption and carbon emissions at each enterprise.
	Renewing []*CarbonEmissionElecSummaryItem `json:"renewing,omitempty" xml:"renewing,omitempty" type:"Repeated"`
	// Data on mains electricity consumption and carbon emission of each enterprise.
	Urban []*CarbonEmissionElecSummaryItem `json:"urban,omitempty" xml:"urban,omitempty" type:"Repeated"`
	// Hydropower consumption and carbon emission data of each enterprise.
	Water []*CarbonEmissionElecSummaryItem `json:"water,omitempty" xml:"water,omitempty" type:"Repeated"`
	// Wind power consumption and carbon emission data of each enterprise.
	Wind []*CarbonEmissionElecSummaryItem `json:"wind,omitempty" xml:"wind,omitempty" type:"Repeated"`
	// Data of zero electricity consumption and carbon emission of each enterprise.
	Zero []*CarbonEmissionElecSummaryItem `json:"zero,omitempty" xml:"zero,omitempty" type:"Repeated"`
}

func (s GetAreaElecConstituteResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetAreaElecConstituteResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAreaElecConstituteResponseBodyData) GetLight() []*CarbonEmissionElecSummaryItem {
	return s.Light
}

func (s *GetAreaElecConstituteResponseBodyData) GetNuclear() []*CarbonEmissionElecSummaryItem {
	return s.Nuclear
}

func (s *GetAreaElecConstituteResponseBodyData) GetRenewing() []*CarbonEmissionElecSummaryItem {
	return s.Renewing
}

func (s *GetAreaElecConstituteResponseBodyData) GetUrban() []*CarbonEmissionElecSummaryItem {
	return s.Urban
}

func (s *GetAreaElecConstituteResponseBodyData) GetWater() []*CarbonEmissionElecSummaryItem {
	return s.Water
}

func (s *GetAreaElecConstituteResponseBodyData) GetWind() []*CarbonEmissionElecSummaryItem {
	return s.Wind
}

func (s *GetAreaElecConstituteResponseBodyData) GetZero() []*CarbonEmissionElecSummaryItem {
	return s.Zero
}

func (s *GetAreaElecConstituteResponseBodyData) SetLight(v []*CarbonEmissionElecSummaryItem) *GetAreaElecConstituteResponseBodyData {
	s.Light = v
	return s
}

func (s *GetAreaElecConstituteResponseBodyData) SetNuclear(v []*CarbonEmissionElecSummaryItem) *GetAreaElecConstituteResponseBodyData {
	s.Nuclear = v
	return s
}

func (s *GetAreaElecConstituteResponseBodyData) SetRenewing(v []*CarbonEmissionElecSummaryItem) *GetAreaElecConstituteResponseBodyData {
	s.Renewing = v
	return s
}

func (s *GetAreaElecConstituteResponseBodyData) SetUrban(v []*CarbonEmissionElecSummaryItem) *GetAreaElecConstituteResponseBodyData {
	s.Urban = v
	return s
}

func (s *GetAreaElecConstituteResponseBodyData) SetWater(v []*CarbonEmissionElecSummaryItem) *GetAreaElecConstituteResponseBodyData {
	s.Water = v
	return s
}

func (s *GetAreaElecConstituteResponseBodyData) SetWind(v []*CarbonEmissionElecSummaryItem) *GetAreaElecConstituteResponseBodyData {
	s.Wind = v
	return s
}

func (s *GetAreaElecConstituteResponseBodyData) SetZero(v []*CarbonEmissionElecSummaryItem) *GetAreaElecConstituteResponseBodyData {
	s.Zero = v
	return s
}

func (s *GetAreaElecConstituteResponseBodyData) Validate() error {
	if s.Light != nil {
		for _, item := range s.Light {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Nuclear != nil {
		for _, item := range s.Nuclear {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Renewing != nil {
		for _, item := range s.Renewing {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Urban != nil {
		for _, item := range s.Urban {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Water != nil {
		for _, item := range s.Water {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Wind != nil {
		for _, item := range s.Wind {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Zero != nil {
		for _, item := range s.Zero {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
