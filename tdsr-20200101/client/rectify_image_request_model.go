// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRectifyImageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCameraHeight(v int64) *RectifyImageRequest
	GetCameraHeight() *int64
	SetUrl(v string) *RectifyImageRequest
	GetUrl() *string
}

type RectifyImageRequest struct {
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

func (s RectifyImageRequest) String() string {
	return dara.Prettify(s)
}

func (s RectifyImageRequest) GoString() string {
	return s.String()
}

func (s *RectifyImageRequest) GetCameraHeight() *int64 {
	return s.CameraHeight
}

func (s *RectifyImageRequest) GetUrl() *string {
	return s.Url
}

func (s *RectifyImageRequest) SetCameraHeight(v int64) *RectifyImageRequest {
	s.CameraHeight = &v
	return s
}

func (s *RectifyImageRequest) SetUrl(v string) *RectifyImageRequest {
	s.Url = &v
	return s
}

func (s *RectifyImageRequest) Validate() error {
	return dara.Validate(s)
}
