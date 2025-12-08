// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iDetectFaceAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageURLObject(v io.Reader) *DetectFaceAdvanceRequest
	GetImageURLObject() io.Reader
	SetLandmark(v bool) *DetectFaceAdvanceRequest
	GetLandmark() *bool
	SetMaxFaceNumber(v int64) *DetectFaceAdvanceRequest
	GetMaxFaceNumber() *int64
	SetPose(v bool) *DetectFaceAdvanceRequest
	GetPose() *bool
	SetQuality(v bool) *DetectFaceAdvanceRequest
	GetQuality() *bool
}

type DetectFaceAdvanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/facebody/DetectFace/DetectFace1.png
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
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

func (s DetectFaceAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DetectFaceAdvanceRequest) GoString() string {
	return s.String()
}

func (s *DetectFaceAdvanceRequest) GetImageURLObject() io.Reader {
	return s.ImageURLObject
}

func (s *DetectFaceAdvanceRequest) GetLandmark() *bool {
	return s.Landmark
}

func (s *DetectFaceAdvanceRequest) GetMaxFaceNumber() *int64 {
	return s.MaxFaceNumber
}

func (s *DetectFaceAdvanceRequest) GetPose() *bool {
	return s.Pose
}

func (s *DetectFaceAdvanceRequest) GetQuality() *bool {
	return s.Quality
}

func (s *DetectFaceAdvanceRequest) SetImageURLObject(v io.Reader) *DetectFaceAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *DetectFaceAdvanceRequest) SetLandmark(v bool) *DetectFaceAdvanceRequest {
	s.Landmark = &v
	return s
}

func (s *DetectFaceAdvanceRequest) SetMaxFaceNumber(v int64) *DetectFaceAdvanceRequest {
	s.MaxFaceNumber = &v
	return s
}

func (s *DetectFaceAdvanceRequest) SetPose(v bool) *DetectFaceAdvanceRequest {
	s.Pose = &v
	return s
}

func (s *DetectFaceAdvanceRequest) SetQuality(v bool) *DetectFaceAdvanceRequest {
	s.Quality = &v
	return s
}

func (s *DetectFaceAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
