// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMeasurementDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndDate(v string) *DescribeMeasurementDataRequest
	GetEndDate() *string
	SetStartDate(v string) *DescribeMeasurementDataRequest
	GetStartDate() *string
}

type DescribeMeasurementDataRequest struct {
	// The end time in UTC. The format is yyyy-MM-ddTHH:mm:ssZ.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2019-08-30T00:00:00Z
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// The start time in UTC. The format is yyyy-MM-ddTHH:mm:ssZ.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2019-06-01T00:00:00Z
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s DescribeMeasurementDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeMeasurementDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeMeasurementDataRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *DescribeMeasurementDataRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *DescribeMeasurementDataRequest) SetEndDate(v string) *DescribeMeasurementDataRequest {
	s.EndDate = &v
	return s
}

func (s *DescribeMeasurementDataRequest) SetStartDate(v string) *DescribeMeasurementDataRequest {
	s.StartDate = &v
	return s
}

func (s *DescribeMeasurementDataRequest) Validate() error {
	return dara.Validate(s)
}
