// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOrgConstituteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetOrgConstituteRequest
	GetCode() *string
	SetModuleCode(v string) *GetOrgConstituteRequest
	GetModuleCode() *string
	SetModuleType(v int32) *GetOrgConstituteRequest
	GetModuleType() *int32
	SetYear(v int32) *GetOrgConstituteRequest
	GetYear() *int32
}

type GetOrgConstituteRequest struct {
	// The enterprise code.
	//
	// This parameter is required.
	//
	// example:
	//
	// Z-20240115-2
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
	// Year.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2024
	Year *int32 `json:"year,omitempty" xml:"year,omitempty"`
}

func (s GetOrgConstituteRequest) String() string {
	return dara.Prettify(s)
}

func (s GetOrgConstituteRequest) GoString() string {
	return s.String()
}

func (s *GetOrgConstituteRequest) GetCode() *string {
	return s.Code
}

func (s *GetOrgConstituteRequest) GetModuleCode() *string {
	return s.ModuleCode
}

func (s *GetOrgConstituteRequest) GetModuleType() *int32 {
	return s.ModuleType
}

func (s *GetOrgConstituteRequest) GetYear() *int32 {
	return s.Year
}

func (s *GetOrgConstituteRequest) SetCode(v string) *GetOrgConstituteRequest {
	s.Code = &v
	return s
}

func (s *GetOrgConstituteRequest) SetModuleCode(v string) *GetOrgConstituteRequest {
	s.ModuleCode = &v
	return s
}

func (s *GetOrgConstituteRequest) SetModuleType(v int32) *GetOrgConstituteRequest {
	s.ModuleType = &v
	return s
}

func (s *GetOrgConstituteRequest) SetYear(v int32) *GetOrgConstituteRequest {
	s.Year = &v
	return s
}

func (s *GetOrgConstituteRequest) Validate() error {
	return dara.Validate(s)
}
