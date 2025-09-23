// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectVehicleICongestionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageURL(v string) *DetectVehicleICongestionRequest
	GetImageURL() *string
	SetPreRegionIntersectFeatures(v []*DetectVehicleICongestionRequestPreRegionIntersectFeatures) *DetectVehicleICongestionRequest
	GetPreRegionIntersectFeatures() []*DetectVehicleICongestionRequestPreRegionIntersectFeatures
	SetRoadRegions(v []*DetectVehicleICongestionRequestRoadRegions) *DetectVehicleICongestionRequest
	GetRoadRegions() []*DetectVehicleICongestionRequestRoadRegions
}

type DetectVehicleICongestionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/objectdet/DetectVehicleICongestion/DetectVehicleICongestion1.jpg
	ImageURL                   *string                                                      `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	PreRegionIntersectFeatures []*DetectVehicleICongestionRequestPreRegionIntersectFeatures `json:"PreRegionIntersectFeatures,omitempty" xml:"PreRegionIntersectFeatures,omitempty" type:"Repeated"`
	// This parameter is required.
	RoadRegions []*DetectVehicleICongestionRequestRoadRegions `json:"RoadRegions,omitempty" xml:"RoadRegions,omitempty" type:"Repeated"`
}

func (s DetectVehicleICongestionRequest) String() string {
	return dara.Prettify(s)
}

func (s DetectVehicleICongestionRequest) GoString() string {
	return s.String()
}

func (s *DetectVehicleICongestionRequest) GetImageURL() *string {
	return s.ImageURL
}

func (s *DetectVehicleICongestionRequest) GetPreRegionIntersectFeatures() []*DetectVehicleICongestionRequestPreRegionIntersectFeatures {
	return s.PreRegionIntersectFeatures
}

func (s *DetectVehicleICongestionRequest) GetRoadRegions() []*DetectVehicleICongestionRequestRoadRegions {
	return s.RoadRegions
}

func (s *DetectVehicleICongestionRequest) SetImageURL(v string) *DetectVehicleICongestionRequest {
	s.ImageURL = &v
	return s
}

func (s *DetectVehicleICongestionRequest) SetPreRegionIntersectFeatures(v []*DetectVehicleICongestionRequestPreRegionIntersectFeatures) *DetectVehicleICongestionRequest {
	s.PreRegionIntersectFeatures = v
	return s
}

func (s *DetectVehicleICongestionRequest) SetRoadRegions(v []*DetectVehicleICongestionRequestRoadRegions) *DetectVehicleICongestionRequest {
	s.RoadRegions = v
	return s
}

func (s *DetectVehicleICongestionRequest) Validate() error {
	return dara.Validate(s)
}

type DetectVehicleICongestionRequestPreRegionIntersectFeatures struct {
	Features []*string `json:"Features,omitempty" xml:"Features,omitempty" type:"Repeated"`
}

func (s DetectVehicleICongestionRequestPreRegionIntersectFeatures) String() string {
	return dara.Prettify(s)
}

func (s DetectVehicleICongestionRequestPreRegionIntersectFeatures) GoString() string {
	return s.String()
}

func (s *DetectVehicleICongestionRequestPreRegionIntersectFeatures) GetFeatures() []*string {
	return s.Features
}

func (s *DetectVehicleICongestionRequestPreRegionIntersectFeatures) SetFeatures(v []*string) *DetectVehicleICongestionRequestPreRegionIntersectFeatures {
	s.Features = v
	return s
}

func (s *DetectVehicleICongestionRequestPreRegionIntersectFeatures) Validate() error {
	return dara.Validate(s)
}

type DetectVehicleICongestionRequestRoadRegions struct {
	// This parameter is required.
	RoadRegion []*DetectVehicleICongestionRequestRoadRegionsRoadRegion `json:"RoadRegion,omitempty" xml:"RoadRegion,omitempty" type:"Repeated"`
}

func (s DetectVehicleICongestionRequestRoadRegions) String() string {
	return dara.Prettify(s)
}

func (s DetectVehicleICongestionRequestRoadRegions) GoString() string {
	return s.String()
}

func (s *DetectVehicleICongestionRequestRoadRegions) GetRoadRegion() []*DetectVehicleICongestionRequestRoadRegionsRoadRegion {
	return s.RoadRegion
}

func (s *DetectVehicleICongestionRequestRoadRegions) SetRoadRegion(v []*DetectVehicleICongestionRequestRoadRegionsRoadRegion) *DetectVehicleICongestionRequestRoadRegions {
	s.RoadRegion = v
	return s
}

func (s *DetectVehicleICongestionRequestRoadRegions) Validate() error {
	return dara.Validate(s)
}

type DetectVehicleICongestionRequestRoadRegionsRoadRegion struct {
	// This parameter is required.
	Point *DetectVehicleICongestionRequestRoadRegionsRoadRegionPoint `json:"Point,omitempty" xml:"Point,omitempty" type:"Struct"`
}

func (s DetectVehicleICongestionRequestRoadRegionsRoadRegion) String() string {
	return dara.Prettify(s)
}

func (s DetectVehicleICongestionRequestRoadRegionsRoadRegion) GoString() string {
	return s.String()
}

func (s *DetectVehicleICongestionRequestRoadRegionsRoadRegion) GetPoint() *DetectVehicleICongestionRequestRoadRegionsRoadRegionPoint {
	return s.Point
}

func (s *DetectVehicleICongestionRequestRoadRegionsRoadRegion) SetPoint(v *DetectVehicleICongestionRequestRoadRegionsRoadRegionPoint) *DetectVehicleICongestionRequestRoadRegionsRoadRegion {
	s.Point = v
	return s
}

func (s *DetectVehicleICongestionRequestRoadRegionsRoadRegion) Validate() error {
	return dara.Validate(s)
}

type DetectVehicleICongestionRequestRoadRegionsRoadRegionPoint struct {
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

func (s DetectVehicleICongestionRequestRoadRegionsRoadRegionPoint) String() string {
	return dara.Prettify(s)
}

func (s DetectVehicleICongestionRequestRoadRegionsRoadRegionPoint) GoString() string {
	return s.String()
}

func (s *DetectVehicleICongestionRequestRoadRegionsRoadRegionPoint) GetX() *int64 {
	return s.X
}

func (s *DetectVehicleICongestionRequestRoadRegionsRoadRegionPoint) GetY() *int64 {
	return s.Y
}

func (s *DetectVehicleICongestionRequestRoadRegionsRoadRegionPoint) SetX(v int64) *DetectVehicleICongestionRequestRoadRegionsRoadRegionPoint {
	s.X = &v
	return s
}

func (s *DetectVehicleICongestionRequestRoadRegionsRoadRegionPoint) SetY(v int64) *DetectVehicleICongestionRequestRoadRegionsRoadRegionPoint {
	s.Y = &v
	return s
}

func (s *DetectVehicleICongestionRequestRoadRegionsRoadRegionPoint) Validate() error {
	return dara.Validate(s)
}
