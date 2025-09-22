// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecalculateCarbonEmissionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RecalculateCarbonEmissionRequest
	GetCode() *string
	SetYear(v string) *RecalculateCarbonEmissionRequest
	GetYear() *string
}

type RecalculateCarbonEmissionRequest struct {
	// The enterprise code.
	//
	// This parameter is required.
	//
	// example:
	//
	// C-20240202-01
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Year of inventory.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2024
	Year *string `json:"year,omitempty" xml:"year,omitempty"`
}

func (s RecalculateCarbonEmissionRequest) String() string {
	return dara.Prettify(s)
}

func (s RecalculateCarbonEmissionRequest) GoString() string {
	return s.String()
}

func (s *RecalculateCarbonEmissionRequest) GetCode() *string {
	return s.Code
}

func (s *RecalculateCarbonEmissionRequest) GetYear() *string {
	return s.Year
}

func (s *RecalculateCarbonEmissionRequest) SetCode(v string) *RecalculateCarbonEmissionRequest {
	s.Code = &v
	return s
}

func (s *RecalculateCarbonEmissionRequest) SetYear(v string) *RecalculateCarbonEmissionRequest {
	s.Year = &v
	return s
}

func (s *RecalculateCarbonEmissionRequest) Validate() error {
	return dara.Validate(s)
}
