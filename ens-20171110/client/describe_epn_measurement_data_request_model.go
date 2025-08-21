// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEpnMeasurementDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndDate(v string) *DescribeEpnMeasurementDataRequest
	GetEndDate() *string
	SetStartDate(v string) *DescribeEpnMeasurementDataRequest
	GetStartDate() *string
}

type DescribeEpnMeasurementDataRequest struct {
	// The end of the time range to query. Specify the time in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2022-01-30T16:00:00Z
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// The beginning of the time range to query. Specify the time in the yyyy-mm-ddthh:mm:ssz format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2021-09-30T16:00:00Z
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s DescribeEpnMeasurementDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeEpnMeasurementDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeEpnMeasurementDataRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *DescribeEpnMeasurementDataRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *DescribeEpnMeasurementDataRequest) SetEndDate(v string) *DescribeEpnMeasurementDataRequest {
	s.EndDate = &v
	return s
}

func (s *DescribeEpnMeasurementDataRequest) SetStartDate(v string) *DescribeEpnMeasurementDataRequest {
	s.StartDate = &v
	return s
}

func (s *DescribeEpnMeasurementDataRequest) Validate() error {
	return dara.Validate(s)
}
