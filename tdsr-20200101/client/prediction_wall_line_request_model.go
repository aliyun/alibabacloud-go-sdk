// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPredictionWallLineRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCameraHeight(v int64) *PredictionWallLineRequest
	GetCameraHeight() *int64
	SetUrl(v string) *PredictionWallLineRequest
	GetUrl() *string
}

type PredictionWallLineRequest struct {
	// example:
	//
	// 160
	CameraHeight *int64 `json:"CameraHeight,omitempty" xml:"CameraHeight,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// https://www.aliyundoc.com/****.jpg
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s PredictionWallLineRequest) String() string {
	return dara.Prettify(s)
}

func (s PredictionWallLineRequest) GoString() string {
	return s.String()
}

func (s *PredictionWallLineRequest) GetCameraHeight() *int64 {
	return s.CameraHeight
}

func (s *PredictionWallLineRequest) GetUrl() *string {
	return s.Url
}

func (s *PredictionWallLineRequest) SetCameraHeight(v int64) *PredictionWallLineRequest {
	s.CameraHeight = &v
	return s
}

func (s *PredictionWallLineRequest) SetUrl(v string) *PredictionWallLineRequest {
	s.Url = &v
	return s
}

func (s *PredictionWallLineRequest) Validate() error {
	return dara.Validate(s)
}
