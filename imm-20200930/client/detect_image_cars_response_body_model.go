// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectImageCarsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCars(v []*Car) *DetectImageCarsResponseBody
	GetCars() []*Car
	SetRequestId(v string) *DetectImageCarsResponseBody
	GetRequestId() *string
}

type DetectImageCarsResponseBody struct {
	// The vehicles.
	//
	// This parameter is required.
	Cars []*Car `json:"Cars,omitempty" xml:"Cars,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// A8745209-DD0E-027E-8ABA-085E0C******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetectImageCarsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DetectImageCarsResponseBody) GoString() string {
	return s.String()
}

func (s *DetectImageCarsResponseBody) GetCars() []*Car {
	return s.Cars
}

func (s *DetectImageCarsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DetectImageCarsResponseBody) SetCars(v []*Car) *DetectImageCarsResponseBody {
	s.Cars = v
	return s
}

func (s *DetectImageCarsResponseBody) SetRequestId(v string) *DetectImageCarsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetectImageCarsResponseBody) Validate() error {
	return dara.Validate(s)
}
