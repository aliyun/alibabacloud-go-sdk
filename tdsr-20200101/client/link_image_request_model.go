// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLinkImageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCameraHeight(v int32) *LinkImageRequest
	GetCameraHeight() *int32
	SetFileName(v string) *LinkImageRequest
	GetFileName() *string
	SetPlatform(v string) *LinkImageRequest
	GetPlatform() *string
	SetSubSceneId(v string) *LinkImageRequest
	GetSubSceneId() *string
}

type LinkImageRequest struct {
	// example:
	//
	// 160
	CameraHeight *int32 `json:"CameraHeight,omitempty" xml:"CameraHeight,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ****.jpg
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// example:
	//
	// PC
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1234****
	SubSceneId *string `json:"SubSceneId,omitempty" xml:"SubSceneId,omitempty"`
}

func (s LinkImageRequest) String() string {
	return dara.Prettify(s)
}

func (s LinkImageRequest) GoString() string {
	return s.String()
}

func (s *LinkImageRequest) GetCameraHeight() *int32 {
	return s.CameraHeight
}

func (s *LinkImageRequest) GetFileName() *string {
	return s.FileName
}

func (s *LinkImageRequest) GetPlatform() *string {
	return s.Platform
}

func (s *LinkImageRequest) GetSubSceneId() *string {
	return s.SubSceneId
}

func (s *LinkImageRequest) SetCameraHeight(v int32) *LinkImageRequest {
	s.CameraHeight = &v
	return s
}

func (s *LinkImageRequest) SetFileName(v string) *LinkImageRequest {
	s.FileName = &v
	return s
}

func (s *LinkImageRequest) SetPlatform(v string) *LinkImageRequest {
	s.Platform = &v
	return s
}

func (s *LinkImageRequest) SetSubSceneId(v string) *LinkImageRequest {
	s.SubSceneId = &v
	return s
}

func (s *LinkImageRequest) Validate() error {
	return dara.Validate(s)
}
