// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGasConstituteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetGasConstituteRequest
	GetCode() *string
	SetModuleCode(v string) *GetGasConstituteRequest
	GetModuleCode() *string
	SetModuleType(v int32) *GetGasConstituteRequest
	GetModuleType() *int32
	SetYear(v int32) *GetGasConstituteRequest
	GetYear() *int32
}

type GetGasConstituteRequest struct {
	// The enterprise code.
	//
	// This parameter is required.
	//
	// example:
	//
	// C-20240115-3
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
	// Year
	//
	// This parameter is required.
	//
	// example:
	//
	// 2024
	Year *int32 `json:"year,omitempty" xml:"year,omitempty"`
}

func (s GetGasConstituteRequest) String() string {
	return dara.Prettify(s)
}

func (s GetGasConstituteRequest) GoString() string {
	return s.String()
}

func (s *GetGasConstituteRequest) GetCode() *string {
	return s.Code
}

func (s *GetGasConstituteRequest) GetModuleCode() *string {
	return s.ModuleCode
}

func (s *GetGasConstituteRequest) GetModuleType() *int32 {
	return s.ModuleType
}

func (s *GetGasConstituteRequest) GetYear() *int32 {
	return s.Year
}

func (s *GetGasConstituteRequest) SetCode(v string) *GetGasConstituteRequest {
	s.Code = &v
	return s
}

func (s *GetGasConstituteRequest) SetModuleCode(v string) *GetGasConstituteRequest {
	s.ModuleCode = &v
	return s
}

func (s *GetGasConstituteRequest) SetModuleType(v int32) *GetGasConstituteRequest {
	s.ModuleType = &v
	return s
}

func (s *GetGasConstituteRequest) SetYear(v int32) *GetGasConstituteRequest {
	s.Year = &v
	return s
}

func (s *GetGasConstituteRequest) Validate() error {
	return dara.Validate(s)
}
