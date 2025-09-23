// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iDetectVehicleICongestionAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageURLObject(v io.Reader) *DetectVehicleICongestionAdvanceRequest
	GetImageURLObject() io.Reader
	SetPreRegionIntersectFeatures(v []*DetectVehicleICongestionAdvanceRequestPreRegionIntersectFeatures) *DetectVehicleICongestionAdvanceRequest
	GetPreRegionIntersectFeatures() []*DetectVehicleICongestionAdvanceRequestPreRegionIntersectFeatures
	SetRoadRegions(v []*DetectVehicleICongestionAdvanceRequestRoadRegions) *DetectVehicleICongestionAdvanceRequest
	GetRoadRegions() []*DetectVehicleICongestionAdvanceRequestRoadRegions
}

type DetectVehicleICongestionAdvanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/objectdet/DetectVehicleICongestion/DetectVehicleICongestion1.jpg
	ImageURLObject             io.Reader                                                           `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	PreRegionIntersectFeatures []*DetectVehicleICongestionAdvanceRequestPreRegionIntersectFeatures `json:"PreRegionIntersectFeatures,omitempty" xml:"PreRegionIntersectFeatures,omitempty" type:"Repeated"`
	// This parameter is required.
	RoadRegions []*DetectVehicleICongestionAdvanceRequestRoadRegions `json:"RoadRegions,omitempty" xml:"RoadRegions,omitempty" type:"Repeated"`
}

func (s DetectVehicleICongestionAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DetectVehicleICongestionAdvanceRequest) GoString() string {
	return s.String()
}

func (s *DetectVehicleICongestionAdvanceRequest) GetImageURLObject() io.Reader {
	return s.ImageURLObject
}

func (s *DetectVehicleICongestionAdvanceRequest) GetPreRegionIntersectFeatures() []*DetectVehicleICongestionAdvanceRequestPreRegionIntersectFeatures {
	return s.PreRegionIntersectFeatures
}

func (s *DetectVehicleICongestionAdvanceRequest) GetRoadRegions() []*DetectVehicleICongestionAdvanceRequestRoadRegions {
	return s.RoadRegions
}

func (s *DetectVehicleICongestionAdvanceRequest) SetImageURLObject(v io.Reader) *DetectVehicleICongestionAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *DetectVehicleICongestionAdvanceRequest) SetPreRegionIntersectFeatures(v []*DetectVehicleICongestionAdvanceRequestPreRegionIntersectFeatures) *DetectVehicleICongestionAdvanceRequest {
	s.PreRegionIntersectFeatures = v
	return s
}

func (s *DetectVehicleICongestionAdvanceRequest) SetRoadRegions(v []*DetectVehicleICongestionAdvanceRequestRoadRegions) *DetectVehicleICongestionAdvanceRequest {
	s.RoadRegions = v
	return s
}

func (s *DetectVehicleICongestionAdvanceRequest) Validate() error {
	return dara.Validate(s)
}

type DetectVehicleICongestionAdvanceRequestPreRegionIntersectFeatures struct {
	Features []*string `json:"Features,omitempty" xml:"Features,omitempty" type:"Repeated"`
}

func (s DetectVehicleICongestionAdvanceRequestPreRegionIntersectFeatures) String() string {
	return dara.Prettify(s)
}

func (s DetectVehicleICongestionAdvanceRequestPreRegionIntersectFeatures) GoString() string {
	return s.String()
}

func (s *DetectVehicleICongestionAdvanceRequestPreRegionIntersectFeatures) GetFeatures() []*string {
	return s.Features
}

func (s *DetectVehicleICongestionAdvanceRequestPreRegionIntersectFeatures) SetFeatures(v []*string) *DetectVehicleICongestionAdvanceRequestPreRegionIntersectFeatures {
	s.Features = v
	return s
}

func (s *DetectVehicleICongestionAdvanceRequestPreRegionIntersectFeatures) Validate() error {
	return dara.Validate(s)
}

type DetectVehicleICongestionAdvanceRequestRoadRegions struct {
	// This parameter is required.
	RoadRegion []*DetectVehicleICongestionAdvanceRequestRoadRegionsRoadRegion `json:"RoadRegion,omitempty" xml:"RoadRegion,omitempty" type:"Repeated"`
}

func (s DetectVehicleICongestionAdvanceRequestRoadRegions) String() string {
	return dara.Prettify(s)
}

func (s DetectVehicleICongestionAdvanceRequestRoadRegions) GoString() string {
	return s.String()
}

func (s *DetectVehicleICongestionAdvanceRequestRoadRegions) GetRoadRegion() []*DetectVehicleICongestionAdvanceRequestRoadRegionsRoadRegion {
	return s.RoadRegion
}

func (s *DetectVehicleICongestionAdvanceRequestRoadRegions) SetRoadRegion(v []*DetectVehicleICongestionAdvanceRequestRoadRegionsRoadRegion) *DetectVehicleICongestionAdvanceRequestRoadRegions {
	s.RoadRegion = v
	return s
}

func (s *DetectVehicleICongestionAdvanceRequestRoadRegions) Validate() error {
	return dara.Validate(s)
}

type DetectVehicleICongestionAdvanceRequestRoadRegionsRoadRegion struct {
	// This parameter is required.
	Point *DetectVehicleICongestionAdvanceRequestRoadRegionsRoadRegionPoint `json:"Point,omitempty" xml:"Point,omitempty" type:"Struct"`
}

func (s DetectVehicleICongestionAdvanceRequestRoadRegionsRoadRegion) String() string {
	return dara.Prettify(s)
}

func (s DetectVehicleICongestionAdvanceRequestRoadRegionsRoadRegion) GoString() string {
	return s.String()
}

func (s *DetectVehicleICongestionAdvanceRequestRoadRegionsRoadRegion) GetPoint() *DetectVehicleICongestionAdvanceRequestRoadRegionsRoadRegionPoint {
	return s.Point
}

func (s *DetectVehicleICongestionAdvanceRequestRoadRegionsRoadRegion) SetPoint(v *DetectVehicleICongestionAdvanceRequestRoadRegionsRoadRegionPoint) *DetectVehicleICongestionAdvanceRequestRoadRegionsRoadRegion {
	s.Point = v
	return s
}

func (s *DetectVehicleICongestionAdvanceRequestRoadRegionsRoadRegion) Validate() error {
	return dara.Validate(s)
}

type DetectVehicleICongestionAdvanceRequestRoadRegionsRoadRegionPoint struct {
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

func (s DetectVehicleICongestionAdvanceRequestRoadRegionsRoadRegionPoint) String() string {
	return dara.Prettify(s)
}

func (s DetectVehicleICongestionAdvanceRequestRoadRegionsRoadRegionPoint) GoString() string {
	return s.String()
}

func (s *DetectVehicleICongestionAdvanceRequestRoadRegionsRoadRegionPoint) GetX() *int64 {
	return s.X
}

func (s *DetectVehicleICongestionAdvanceRequestRoadRegionsRoadRegionPoint) GetY() *int64 {
	return s.Y
}

func (s *DetectVehicleICongestionAdvanceRequestRoadRegionsRoadRegionPoint) SetX(v int64) *DetectVehicleICongestionAdvanceRequestRoadRegionsRoadRegionPoint {
	s.X = &v
	return s
}

func (s *DetectVehicleICongestionAdvanceRequestRoadRegionsRoadRegionPoint) SetY(v int64) *DetectVehicleICongestionAdvanceRequestRoadRegionsRoadRegionPoint {
	s.Y = &v
	return s
}

func (s *DetectVehicleICongestionAdvanceRequestRoadRegionsRoadRegionPoint) Validate() error {
	return dara.Validate(s)
}
