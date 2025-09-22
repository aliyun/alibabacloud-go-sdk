// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEmissionSummaryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetEmissionSummaryRequest
	GetCode() *string
	SetModuleCode(v string) *GetEmissionSummaryRequest
	GetModuleCode() *string
	SetModuleType(v int32) *GetEmissionSummaryRequest
	GetModuleType() *int32
	SetYear(v int32) *GetEmissionSummaryRequest
	GetYear() *int32
}

type GetEmissionSummaryRequest struct {
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
	// Year of inventory.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2024
	Year *int32 `json:"year,omitempty" xml:"year,omitempty"`
}

func (s GetEmissionSummaryRequest) String() string {
	return dara.Prettify(s)
}

func (s GetEmissionSummaryRequest) GoString() string {
	return s.String()
}

func (s *GetEmissionSummaryRequest) GetCode() *string {
	return s.Code
}

func (s *GetEmissionSummaryRequest) GetModuleCode() *string {
	return s.ModuleCode
}

func (s *GetEmissionSummaryRequest) GetModuleType() *int32 {
	return s.ModuleType
}

func (s *GetEmissionSummaryRequest) GetYear() *int32 {
	return s.Year
}

func (s *GetEmissionSummaryRequest) SetCode(v string) *GetEmissionSummaryRequest {
	s.Code = &v
	return s
}

func (s *GetEmissionSummaryRequest) SetModuleCode(v string) *GetEmissionSummaryRequest {
	s.ModuleCode = &v
	return s
}

func (s *GetEmissionSummaryRequest) SetModuleType(v int32) *GetEmissionSummaryRequest {
	s.ModuleType = &v
	return s
}

func (s *GetEmissionSummaryRequest) SetYear(v int32) *GetEmissionSummaryRequest {
	s.Year = &v
	return s
}

func (s *GetEmissionSummaryRequest) Validate() error {
	return dara.Validate(s)
}
