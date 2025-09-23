// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectVehicleICongestionShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageURL(v string) *DetectVehicleICongestionShrinkRequest
	GetImageURL() *string
	SetPreRegionIntersectFeaturesShrink(v string) *DetectVehicleICongestionShrinkRequest
	GetPreRegionIntersectFeaturesShrink() *string
	SetRoadRegionsShrink(v string) *DetectVehicleICongestionShrinkRequest
	GetRoadRegionsShrink() *string
}

type DetectVehicleICongestionShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/objectdet/DetectVehicleICongestion/DetectVehicleICongestion1.jpg
	ImageURL                         *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	PreRegionIntersectFeaturesShrink *string `json:"PreRegionIntersectFeatures,omitempty" xml:"PreRegionIntersectFeatures,omitempty"`
	// This parameter is required.
	RoadRegionsShrink *string `json:"RoadRegions,omitempty" xml:"RoadRegions,omitempty"`
}

func (s DetectVehicleICongestionShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DetectVehicleICongestionShrinkRequest) GoString() string {
	return s.String()
}

func (s *DetectVehicleICongestionShrinkRequest) GetImageURL() *string {
	return s.ImageURL
}

func (s *DetectVehicleICongestionShrinkRequest) GetPreRegionIntersectFeaturesShrink() *string {
	return s.PreRegionIntersectFeaturesShrink
}

func (s *DetectVehicleICongestionShrinkRequest) GetRoadRegionsShrink() *string {
	return s.RoadRegionsShrink
}

func (s *DetectVehicleICongestionShrinkRequest) SetImageURL(v string) *DetectVehicleICongestionShrinkRequest {
	s.ImageURL = &v
	return s
}

func (s *DetectVehicleICongestionShrinkRequest) SetPreRegionIntersectFeaturesShrink(v string) *DetectVehicleICongestionShrinkRequest {
	s.PreRegionIntersectFeaturesShrink = &v
	return s
}

func (s *DetectVehicleICongestionShrinkRequest) SetRoadRegionsShrink(v string) *DetectVehicleICongestionShrinkRequest {
	s.RoadRegionsShrink = &v
	return s
}

func (s *DetectVehicleICongestionShrinkRequest) Validate() error {
	return dara.Validate(s)
}
