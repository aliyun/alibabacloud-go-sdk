// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectVehicleIllegalParkingShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageURL(v string) *DetectVehicleIllegalParkingShrinkRequest
	GetImageURL() *string
	SetRoadRegionsShrink(v string) *DetectVehicleIllegalParkingShrinkRequest
	GetRoadRegionsShrink() *string
}

type DetectVehicleIllegalParkingShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/objectdet/DetectVehicleIllegalParking/DetectVehicleIllegalParking2.jpg
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	// This parameter is required.
	RoadRegionsShrink *string `json:"RoadRegions,omitempty" xml:"RoadRegions,omitempty"`
}

func (s DetectVehicleIllegalParkingShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DetectVehicleIllegalParkingShrinkRequest) GoString() string {
	return s.String()
}

func (s *DetectVehicleIllegalParkingShrinkRequest) GetImageURL() *string {
	return s.ImageURL
}

func (s *DetectVehicleIllegalParkingShrinkRequest) GetRoadRegionsShrink() *string {
	return s.RoadRegionsShrink
}

func (s *DetectVehicleIllegalParkingShrinkRequest) SetImageURL(v string) *DetectVehicleIllegalParkingShrinkRequest {
	s.ImageURL = &v
	return s
}

func (s *DetectVehicleIllegalParkingShrinkRequest) SetRoadRegionsShrink(v string) *DetectVehicleIllegalParkingShrinkRequest {
	s.RoadRegionsShrink = &v
	return s
}

func (s *DetectVehicleIllegalParkingShrinkRequest) Validate() error {
	return dara.Validate(s)
}
