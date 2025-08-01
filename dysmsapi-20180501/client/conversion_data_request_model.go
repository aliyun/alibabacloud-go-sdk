// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConversionDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConversionRate(v string) *ConversionDataRequest
	GetConversionRate() *string
	SetReportTime(v int64) *ConversionDataRequest
	GetReportTime() *int64
}

type ConversionDataRequest struct {
	// Conversion rate monitoring return value.
	//
	// >  The value of this parameter is of type double, and the value is between [0,1].
	//
	// This parameter is required.
	//
	// example:
	//
	// 0.53
	ConversionRate *string `json:"ConversionRate,omitempty" xml:"ConversionRate,omitempty"`
	// Timestamp of the conversion rate observation should be a Unix timestamp, a millisecond-level long integer.
	//
	// >  If this field is not specified: the current timestamp is the default.
	//
	// example:
	//
	// 1349055900000
	ReportTime *int64 `json:"ReportTime,omitempty" xml:"ReportTime,omitempty"`
}

func (s ConversionDataRequest) String() string {
	return dara.Prettify(s)
}

func (s ConversionDataRequest) GoString() string {
	return s.String()
}

func (s *ConversionDataRequest) GetConversionRate() *string {
	return s.ConversionRate
}

func (s *ConversionDataRequest) GetReportTime() *int64 {
	return s.ReportTime
}

func (s *ConversionDataRequest) SetConversionRate(v string) *ConversionDataRequest {
	s.ConversionRate = &v
	return s
}

func (s *ConversionDataRequest) SetReportTime(v int64) *ConversionDataRequest {
	s.ReportTime = &v
	return s
}

func (s *ConversionDataRequest) Validate() error {
	return dara.Validate(s)
}
