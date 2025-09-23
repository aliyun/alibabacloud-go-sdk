// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectVehicleIllegalParkingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageURL(v string) *DetectVehicleIllegalParkingRequest
	GetImageURL() *string
	SetRoadRegions(v []*DetectVehicleIllegalParkingRequestRoadRegions) *DetectVehicleIllegalParkingRequest
	GetRoadRegions() []*DetectVehicleIllegalParkingRequestRoadRegions
}

type DetectVehicleIllegalParkingRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/objectdet/DetectVehicleIllegalParking/DetectVehicleIllegalParking2.jpg
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	// This parameter is required.
	RoadRegions []*DetectVehicleIllegalParkingRequestRoadRegions `json:"RoadRegions,omitempty" xml:"RoadRegions,omitempty" type:"Repeated"`
}

func (s DetectVehicleIllegalParkingRequest) String() string {
	return dara.Prettify(s)
}

func (s DetectVehicleIllegalParkingRequest) GoString() string {
	return s.String()
}

func (s *DetectVehicleIllegalParkingRequest) GetImageURL() *string {
	return s.ImageURL
}

func (s *DetectVehicleIllegalParkingRequest) GetRoadRegions() []*DetectVehicleIllegalParkingRequestRoadRegions {
	return s.RoadRegions
}

func (s *DetectVehicleIllegalParkingRequest) SetImageURL(v string) *DetectVehicleIllegalParkingRequest {
	s.ImageURL = &v
	return s
}

func (s *DetectVehicleIllegalParkingRequest) SetRoadRegions(v []*DetectVehicleIllegalParkingRequestRoadRegions) *DetectVehicleIllegalParkingRequest {
	s.RoadRegions = v
	return s
}

func (s *DetectVehicleIllegalParkingRequest) Validate() error {
	return dara.Validate(s)
}

type DetectVehicleIllegalParkingRequestRoadRegions struct {
	// This parameter is required.
	RoadRegion []*DetectVehicleIllegalParkingRequestRoadRegionsRoadRegion `json:"RoadRegion,omitempty" xml:"RoadRegion,omitempty" type:"Repeated"`
}

func (s DetectVehicleIllegalParkingRequestRoadRegions) String() string {
	return dara.Prettify(s)
}

func (s DetectVehicleIllegalParkingRequestRoadRegions) GoString() string {
	return s.String()
}

func (s *DetectVehicleIllegalParkingRequestRoadRegions) GetRoadRegion() []*DetectVehicleIllegalParkingRequestRoadRegionsRoadRegion {
	return s.RoadRegion
}

func (s *DetectVehicleIllegalParkingRequestRoadRegions) SetRoadRegion(v []*DetectVehicleIllegalParkingRequestRoadRegionsRoadRegion) *DetectVehicleIllegalParkingRequestRoadRegions {
	s.RoadRegion = v
	return s
}

func (s *DetectVehicleIllegalParkingRequestRoadRegions) Validate() error {
	return dara.Validate(s)
}

type DetectVehicleIllegalParkingRequestRoadRegionsRoadRegion struct {
	// This parameter is required.
	Point *DetectVehicleIllegalParkingRequestRoadRegionsRoadRegionPoint `json:"Point,omitempty" xml:"Point,omitempty" type:"Struct"`
}

func (s DetectVehicleIllegalParkingRequestRoadRegionsRoadRegion) String() string {
	return dara.Prettify(s)
}

func (s DetectVehicleIllegalParkingRequestRoadRegionsRoadRegion) GoString() string {
	return s.String()
}

func (s *DetectVehicleIllegalParkingRequestRoadRegionsRoadRegion) GetPoint() *DetectVehicleIllegalParkingRequestRoadRegionsRoadRegionPoint {
	return s.Point
}

func (s *DetectVehicleIllegalParkingRequestRoadRegionsRoadRegion) SetPoint(v *DetectVehicleIllegalParkingRequestRoadRegionsRoadRegionPoint) *DetectVehicleIllegalParkingRequestRoadRegionsRoadRegion {
	s.Point = v
	return s
}

func (s *DetectVehicleIllegalParkingRequestRoadRegionsRoadRegion) Validate() error {
	return dara.Validate(s)
}

type DetectVehicleIllegalParkingRequestRoadRegionsRoadRegionPoint struct {
	// This parameter is required.
	//
	// example:
	//
	// 400
	X *int64 `json:"X,omitempty" xml:"X,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 400
	Y *int64 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s DetectVehicleIllegalParkingRequestRoadRegionsRoadRegionPoint) String() string {
	return dara.Prettify(s)
}

func (s DetectVehicleIllegalParkingRequestRoadRegionsRoadRegionPoint) GoString() string {
	return s.String()
}

func (s *DetectVehicleIllegalParkingRequestRoadRegionsRoadRegionPoint) GetX() *int64 {
	return s.X
}

func (s *DetectVehicleIllegalParkingRequestRoadRegionsRoadRegionPoint) GetY() *int64 {
	return s.Y
}

func (s *DetectVehicleIllegalParkingRequestRoadRegionsRoadRegionPoint) SetX(v int64) *DetectVehicleIllegalParkingRequestRoadRegionsRoadRegionPoint {
	s.X = &v
	return s
}

func (s *DetectVehicleIllegalParkingRequestRoadRegionsRoadRegionPoint) SetY(v int64) *DetectVehicleIllegalParkingRequestRoadRegionsRoadRegionPoint {
	s.Y = &v
	return s
}

func (s *DetectVehicleIllegalParkingRequestRoadRegionsRoadRegionPoint) Validate() error {
	return dara.Validate(s)
}
