// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectFaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageURL(v string) *DetectFaceRequest
	GetImageURL() *string
	SetLandmark(v bool) *DetectFaceRequest
	GetLandmark() *bool
	SetMaxFaceNumber(v int64) *DetectFaceRequest
	GetMaxFaceNumber() *int64
	SetPose(v bool) *DetectFaceRequest
	GetPose() *bool
	SetQuality(v bool) *DetectFaceRequest
	GetQuality() *bool
}

type DetectFaceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/facebody/DetectFace/DetectFace1.png
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	// example:
	//
	// true
	Landmark *bool `json:"Landmark,omitempty" xml:"Landmark,omitempty"`
	// example:
	//
	// 1
	MaxFaceNumber *int64 `json:"MaxFaceNumber,omitempty" xml:"MaxFaceNumber,omitempty"`
	// example:
	//
	// true
	Pose *bool `json:"Pose,omitempty" xml:"Pose,omitempty"`
	// example:
	//
	// true
	Quality *bool `json:"Quality,omitempty" xml:"Quality,omitempty"`
}

func (s DetectFaceRequest) String() string {
	return dara.Prettify(s)
}

func (s DetectFaceRequest) GoString() string {
	return s.String()
}

func (s *DetectFaceRequest) GetImageURL() *string {
	return s.ImageURL
}

func (s *DetectFaceRequest) GetLandmark() *bool {
	return s.Landmark
}

func (s *DetectFaceRequest) GetMaxFaceNumber() *int64 {
	return s.MaxFaceNumber
}

func (s *DetectFaceRequest) GetPose() *bool {
	return s.Pose
}

func (s *DetectFaceRequest) GetQuality() *bool {
	return s.Quality
}

func (s *DetectFaceRequest) SetImageURL(v string) *DetectFaceRequest {
	s.ImageURL = &v
	return s
}

func (s *DetectFaceRequest) SetLandmark(v bool) *DetectFaceRequest {
	s.Landmark = &v
	return s
}

func (s *DetectFaceRequest) SetMaxFaceNumber(v int64) *DetectFaceRequest {
	s.MaxFaceNumber = &v
	return s
}

func (s *DetectFaceRequest) SetPose(v bool) *DetectFaceRequest {
	s.Pose = &v
	return s
}

func (s *DetectFaceRequest) SetQuality(v bool) *DetectFaceRequest {
	s.Quality = &v
	return s
}

func (s *DetectFaceRequest) Validate() error {
	return dara.Validate(s)
}
