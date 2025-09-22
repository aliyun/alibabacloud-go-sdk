// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEmissionSourceConstituteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetEmissionSourceConstituteRequest
	GetCode() *string
	SetModuleCode(v string) *GetEmissionSourceConstituteRequest
	GetModuleCode() *string
	SetModuleType(v int32) *GetEmissionSourceConstituteRequest
	GetModuleType() *int32
	SetYear(v int32) *GetEmissionSourceConstituteRequest
	GetYear() *int32
}

type GetEmissionSourceConstituteRequest struct {
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
	// This parameter is required.
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

func (s GetEmissionSourceConstituteRequest) String() string {
	return dara.Prettify(s)
}

func (s GetEmissionSourceConstituteRequest) GoString() string {
	return s.String()
}

func (s *GetEmissionSourceConstituteRequest) GetCode() *string {
	return s.Code
}

func (s *GetEmissionSourceConstituteRequest) GetModuleCode() *string {
	return s.ModuleCode
}

func (s *GetEmissionSourceConstituteRequest) GetModuleType() *int32 {
	return s.ModuleType
}

func (s *GetEmissionSourceConstituteRequest) GetYear() *int32 {
	return s.Year
}

func (s *GetEmissionSourceConstituteRequest) SetCode(v string) *GetEmissionSourceConstituteRequest {
	s.Code = &v
	return s
}

func (s *GetEmissionSourceConstituteRequest) SetModuleCode(v string) *GetEmissionSourceConstituteRequest {
	s.ModuleCode = &v
	return s
}

func (s *GetEmissionSourceConstituteRequest) SetModuleType(v int32) *GetEmissionSourceConstituteRequest {
	s.ModuleType = &v
	return s
}

func (s *GetEmissionSourceConstituteRequest) SetYear(v int32) *GetEmissionSourceConstituteRequest {
	s.Year = &v
	return s
}

func (s *GetEmissionSourceConstituteRequest) Validate() error {
	return dara.Validate(s)
}
