// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthReportInterval interface {
	dara.Model
	String() string
	GoString() string
	SetTimeUnit(v string) *AuthReportInterval
	GetTimeUnit() *string
	SetValue(v int64) *AuthReportInterval
	GetValue() *int64
}

type AuthReportInterval struct {
	// The unit of the effective recovery duration. Valid values:
	//
	// - **Minute**: Minutes.
	//
	// - **Hour**: Hours.
	//
	// - **Day**: Days.
	//
	// example:
	//
	// Minute
	TimeUnit *string `json:"TimeUnit,omitempty" xml:"TimeUnit,omitempty"`
	// The numeric value of the effective recovery duration. The unit is specified by TimeUnit. A non-zero value must be at least 1.
	//
	// example:
	//
	// 60
	Value *int64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s AuthReportInterval) String() string {
	return dara.Prettify(s)
}

func (s AuthReportInterval) GoString() string {
	return s.String()
}

func (s *AuthReportInterval) GetTimeUnit() *string {
	return s.TimeUnit
}

func (s *AuthReportInterval) GetValue() *int64 {
	return s.Value
}

func (s *AuthReportInterval) SetTimeUnit(v string) *AuthReportInterval {
	s.TimeUnit = &v
	return s
}

func (s *AuthReportInterval) SetValue(v int64) *AuthReportInterval {
	s.Value = &v
	return s
}

func (s *AuthReportInterval) Validate() error {
	return dara.Validate(s)
}
