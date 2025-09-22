// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCarbonEmissionTrendRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetCarbonEmissionTrendRequest
	GetCode() *string
	SetModuleCode(v string) *GetCarbonEmissionTrendRequest
	GetModuleCode() *string
	SetModuleType(v int32) *GetCarbonEmissionTrendRequest
	GetModuleType() *int32
	SetTrendType(v int32) *GetCarbonEmissionTrendRequest
	GetTrendType() *int32
	SetYearList(v []*int32) *GetCarbonEmissionTrendRequest
	GetYearList() []*int32
}

type GetCarbonEmissionTrendRequest struct {
	// The enterprise code.
	//
	// This parameter is required.
	//
	// example:
	//
	// C-20240119-1
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Module code.
	//
	// example:
	//
	// carbonInventory.check.scope_1_direct_ghg_emissions
	ModuleCode *string `json:"moduleCode,omitempty" xml:"moduleCode,omitempty"`
	// Module type.
	//
	// example:
	//
	// 3
	ModuleType *int32 `json:"moduleType,omitempty" xml:"moduleType,omitempty"`
	// Trend Type.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	TrendType *int32 `json:"trendType,omitempty" xml:"trendType,omitempty"`
	// The list of inventory year.
	//
	// This parameter is required.
	YearList []*int32 `json:"yearList,omitempty" xml:"yearList,omitempty" type:"Repeated"`
}

func (s GetCarbonEmissionTrendRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCarbonEmissionTrendRequest) GoString() string {
	return s.String()
}

func (s *GetCarbonEmissionTrendRequest) GetCode() *string {
	return s.Code
}

func (s *GetCarbonEmissionTrendRequest) GetModuleCode() *string {
	return s.ModuleCode
}

func (s *GetCarbonEmissionTrendRequest) GetModuleType() *int32 {
	return s.ModuleType
}

func (s *GetCarbonEmissionTrendRequest) GetTrendType() *int32 {
	return s.TrendType
}

func (s *GetCarbonEmissionTrendRequest) GetYearList() []*int32 {
	return s.YearList
}

func (s *GetCarbonEmissionTrendRequest) SetCode(v string) *GetCarbonEmissionTrendRequest {
	s.Code = &v
	return s
}

func (s *GetCarbonEmissionTrendRequest) SetModuleCode(v string) *GetCarbonEmissionTrendRequest {
	s.ModuleCode = &v
	return s
}

func (s *GetCarbonEmissionTrendRequest) SetModuleType(v int32) *GetCarbonEmissionTrendRequest {
	s.ModuleType = &v
	return s
}

func (s *GetCarbonEmissionTrendRequest) SetTrendType(v int32) *GetCarbonEmissionTrendRequest {
	s.TrendType = &v
	return s
}

func (s *GetCarbonEmissionTrendRequest) SetYearList(v []*int32) *GetCarbonEmissionTrendRequest {
	s.YearList = v
	return s
}

func (s *GetCarbonEmissionTrendRequest) Validate() error {
	return dara.Validate(s)
}
