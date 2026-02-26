// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCar interface {
	dara.Model
	String() string
	GoString() string
	SetBoundary(v *Boundary) *Car
	GetBoundary() *Boundary
	SetCarColor(v string) *Car
	GetCarColor() *string
	SetCarColorConfidence(v float64) *Car
	GetCarColorConfidence() *float64
	SetCarType(v string) *Car
	GetCarType() *string
	SetCarTypeConfidence(v float64) *Car
	GetCarTypeConfidence() *float64
	SetConfidence(v float64) *Car
	GetConfidence() *float64
	SetLicensePlates(v []*LicensePlate) *Car
	GetLicensePlates() []*LicensePlate
}

type Car struct {
	// The boundary information.
	Boundary *Boundary `json:"Boundary,omitempty" xml:"Boundary,omitempty"`
	// The vehicle color. Valid values
	//
	// 	- white
	//
	// 	- grey
	//
	// 	- yellow
	//
	// 	- red
	//
	// 	- green
	//
	// 	- blue
	//
	// 	- black
	//
	// 	- purple
	//
	// 	- brown
	//
	// example:
	//
	// white
	CarColor *string `json:"CarColor,omitempty" xml:"CarColor,omitempty"`
	// The confidence level of the vehicle color. Valid values: 0 to 1. The value 0 indicates the lowest confidence level. The value 1 indicates the highest confidence level.
	//
	// example:
	//
	// 0.604
	CarColorConfidence *float64 `json:"CarColorConfidence,omitempty" xml:"CarColorConfidence,omitempty"`
	// The vehicle type. Valid values:
	//
	// 	- car
	//
	// 	- bus
	//
	// 	- truck
	//
	// 	- van
	//
	// example:
	//
	// van
	CarType *string `json:"CarType,omitempty" xml:"CarType,omitempty"`
	// The confidence level of the vehicle type. Valid values: 0 to 1. The value 0 indicates the lowest confidence level. The value 1 indicates the highest confidence level.
	//
	// example:
	//
	// 0.516
	CarTypeConfidence *float64 `json:"CarTypeConfidence,omitempty" xml:"CarTypeConfidence,omitempty"`
	// The confidence level of the vehicle detection result. Valid values: 0 to 1. The value 0 indicates the lowest confidence level. The value 1 indicates the highest confidence level.
	//
	// example:
	//
	// 0.999
	Confidence *float64 `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	// The license plates.
	LicensePlates []*LicensePlate `json:"LicensePlates,omitempty" xml:"LicensePlates,omitempty" type:"Repeated"`
}

func (s Car) String() string {
	return dara.Prettify(s)
}

func (s Car) GoString() string {
	return s.String()
}

func (s *Car) GetBoundary() *Boundary {
	return s.Boundary
}

func (s *Car) GetCarColor() *string {
	return s.CarColor
}

func (s *Car) GetCarColorConfidence() *float64 {
	return s.CarColorConfidence
}

func (s *Car) GetCarType() *string {
	return s.CarType
}

func (s *Car) GetCarTypeConfidence() *float64 {
	return s.CarTypeConfidence
}

func (s *Car) GetConfidence() *float64 {
	return s.Confidence
}

func (s *Car) GetLicensePlates() []*LicensePlate {
	return s.LicensePlates
}

func (s *Car) SetBoundary(v *Boundary) *Car {
	s.Boundary = v
	return s
}

func (s *Car) SetCarColor(v string) *Car {
	s.CarColor = &v
	return s
}

func (s *Car) SetCarColorConfidence(v float64) *Car {
	s.CarColorConfidence = &v
	return s
}

func (s *Car) SetCarType(v string) *Car {
	s.CarType = &v
	return s
}

func (s *Car) SetCarTypeConfidence(v float64) *Car {
	s.CarTypeConfidence = &v
	return s
}

func (s *Car) SetConfidence(v float64) *Car {
	s.Confidence = &v
	return s
}

func (s *Car) SetLicensePlates(v []*LicensePlate) *Car {
	s.LicensePlates = v
	return s
}

func (s *Car) Validate() error {
	if s.Boundary != nil {
		if err := s.Boundary.Validate(); err != nil {
			return err
		}
	}
	if s.LicensePlates != nil {
		for _, item := range s.LicensePlates {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
