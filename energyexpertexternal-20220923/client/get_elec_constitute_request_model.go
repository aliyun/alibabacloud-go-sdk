// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetElecConstituteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetElecConstituteRequest
	GetCode() *string
	SetYear(v int32) *GetElecConstituteRequest
	GetYear() *int32
}

type GetElecConstituteRequest struct {
	// The enterprise code.
	//
	// This parameter is required.
	//
	// example:
	//
	// C-20240202-01
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Year.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2024
	Year *int32 `json:"year,omitempty" xml:"year,omitempty"`
}

func (s GetElecConstituteRequest) String() string {
	return dara.Prettify(s)
}

func (s GetElecConstituteRequest) GoString() string {
	return s.String()
}

func (s *GetElecConstituteRequest) GetCode() *string {
	return s.Code
}

func (s *GetElecConstituteRequest) GetYear() *int32 {
	return s.Year
}

func (s *GetElecConstituteRequest) SetCode(v string) *GetElecConstituteRequest {
	s.Code = &v
	return s
}

func (s *GetElecConstituteRequest) SetYear(v int32) *GetElecConstituteRequest {
	s.Year = &v
	return s
}

func (s *GetElecConstituteRequest) Validate() error {
	return dara.Validate(s)
}
