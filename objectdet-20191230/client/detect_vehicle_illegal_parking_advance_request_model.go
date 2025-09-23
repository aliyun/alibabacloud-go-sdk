// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iDetectVehicleIllegalParkingAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageURLObject(v io.Reader) *DetectVehicleIllegalParkingAdvanceRequest
	GetImageURLObject() io.Reader
	SetRoadRegions(v []*DetectVehicleIllegalParkingAdvanceRequestRoadRegions) *DetectVehicleIllegalParkingAdvanceRequest
	GetRoadRegions() []*DetectVehicleIllegalParkingAdvanceRequestRoadRegions
}

type DetectVehicleIllegalParkingAdvanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/objectdet/DetectVehicleIllegalParking/DetectVehicleIllegalParking2.jpg
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	// This parameter is required.
	RoadRegions []*DetectVehicleIllegalParkingAdvanceRequestRoadRegions `json:"RoadRegions,omitempty" xml:"RoadRegions,omitempty" type:"Repeated"`
}

func (s DetectVehicleIllegalParkingAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DetectVehicleIllegalParkingAdvanceRequest) GoString() string {
	return s.String()
}

func (s *DetectVehicleIllegalParkingAdvanceRequest) GetImageURLObject() io.Reader {
	return s.ImageURLObject
}

func (s *DetectVehicleIllegalParkingAdvanceRequest) GetRoadRegions() []*DetectVehicleIllegalParkingAdvanceRequestRoadRegions {
	return s.RoadRegions
}

func (s *DetectVehicleIllegalParkingAdvanceRequest) SetImageURLObject(v io.Reader) *DetectVehicleIllegalParkingAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *DetectVehicleIllegalParkingAdvanceRequest) SetRoadRegions(v []*DetectVehicleIllegalParkingAdvanceRequestRoadRegions) *DetectVehicleIllegalParkingAdvanceRequest {
	s.RoadRegions = v
	return s
}

func (s *DetectVehicleIllegalParkingAdvanceRequest) Validate() error {
	return dara.Validate(s)
}

type DetectVehicleIllegalParkingAdvanceRequestRoadRegions struct {
	// This parameter is required.
	RoadRegion []*DetectVehicleIllegalParkingAdvanceRequestRoadRegionsRoadRegion `json:"RoadRegion,omitempty" xml:"RoadRegion,omitempty" type:"Repeated"`
}

func (s DetectVehicleIllegalParkingAdvanceRequestRoadRegions) String() string {
	return dara.Prettify(s)
}

func (s DetectVehicleIllegalParkingAdvanceRequestRoadRegions) GoString() string {
	return s.String()
}

func (s *DetectVehicleIllegalParkingAdvanceRequestRoadRegions) GetRoadRegion() []*DetectVehicleIllegalParkingAdvanceRequestRoadRegionsRoadRegion {
	return s.RoadRegion
}

func (s *DetectVehicleIllegalParkingAdvanceRequestRoadRegions) SetRoadRegion(v []*DetectVehicleIllegalParkingAdvanceRequestRoadRegionsRoadRegion) *DetectVehicleIllegalParkingAdvanceRequestRoadRegions {
	s.RoadRegion = v
	return s
}

func (s *DetectVehicleIllegalParkingAdvanceRequestRoadRegions) Validate() error {
	return dara.Validate(s)
}

type DetectVehicleIllegalParkingAdvanceRequestRoadRegionsRoadRegion struct {
	// This parameter is required.
	Point *DetectVehicleIllegalParkingAdvanceRequestRoadRegionsRoadRegionPoint `json:"Point,omitempty" xml:"Point,omitempty" type:"Struct"`
}

func (s DetectVehicleIllegalParkingAdvanceRequestRoadRegionsRoadRegion) String() string {
	return dara.Prettify(s)
}

func (s DetectVehicleIllegalParkingAdvanceRequestRoadRegionsRoadRegion) GoString() string {
	return s.String()
}

func (s *DetectVehicleIllegalParkingAdvanceRequestRoadRegionsRoadRegion) GetPoint() *DetectVehicleIllegalParkingAdvanceRequestRoadRegionsRoadRegionPoint {
	return s.Point
}

func (s *DetectVehicleIllegalParkingAdvanceRequestRoadRegionsRoadRegion) SetPoint(v *DetectVehicleIllegalParkingAdvanceRequestRoadRegionsRoadRegionPoint) *DetectVehicleIllegalParkingAdvanceRequestRoadRegionsRoadRegion {
	s.Point = v
	return s
}

func (s *DetectVehicleIllegalParkingAdvanceRequestRoadRegionsRoadRegion) Validate() error {
	return dara.Validate(s)
}

type DetectVehicleIllegalParkingAdvanceRequestRoadRegionsRoadRegionPoint struct {
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

func (s DetectVehicleIllegalParkingAdvanceRequestRoadRegionsRoadRegionPoint) String() string {
	return dara.Prettify(s)
}

func (s DetectVehicleIllegalParkingAdvanceRequestRoadRegionsRoadRegionPoint) GoString() string {
	return s.String()
}

func (s *DetectVehicleIllegalParkingAdvanceRequestRoadRegionsRoadRegionPoint) GetX() *int64 {
	return s.X
}

func (s *DetectVehicleIllegalParkingAdvanceRequestRoadRegionsRoadRegionPoint) GetY() *int64 {
	return s.Y
}

func (s *DetectVehicleIllegalParkingAdvanceRequestRoadRegionsRoadRegionPoint) SetX(v int64) *DetectVehicleIllegalParkingAdvanceRequestRoadRegionsRoadRegionPoint {
	s.X = &v
	return s
}

func (s *DetectVehicleIllegalParkingAdvanceRequestRoadRegionsRoadRegionPoint) SetY(v int64) *DetectVehicleIllegalParkingAdvanceRequestRoadRegionsRoadRegionPoint {
	s.Y = &v
	return s
}

func (s *DetectVehicleIllegalParkingAdvanceRequestRoadRegionsRoadRegionPoint) Validate() error {
	return dara.Validate(s)
}
