// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAreaElecConstituteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetAreaElecConstituteRequest
	GetCode() *string
	SetYear(v int32) *GetAreaElecConstituteRequest
	GetYear() *int32
}

type GetAreaElecConstituteRequest struct {
	// The enterprise code.
	//
	// This parameter is required.
	//
	// example:
	//
	// Z-20240115-2
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

func (s GetAreaElecConstituteRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAreaElecConstituteRequest) GoString() string {
	return s.String()
}

func (s *GetAreaElecConstituteRequest) GetCode() *string {
	return s.Code
}

func (s *GetAreaElecConstituteRequest) GetYear() *int32 {
	return s.Year
}

func (s *GetAreaElecConstituteRequest) SetCode(v string) *GetAreaElecConstituteRequest {
	s.Code = &v
	return s
}

func (s *GetAreaElecConstituteRequest) SetYear(v int32) *GetAreaElecConstituteRequest {
	s.Year = &v
	return s
}

func (s *GetAreaElecConstituteRequest) Validate() error {
	return dara.Validate(s)
}
